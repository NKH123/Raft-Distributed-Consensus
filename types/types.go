package types

import (
	"net/http"
	"sync"
)

// ServerData is the per server data that contains
// the most updated data from each server. This is
// maintained for its own needs. Maps from the
// servers IP to its state.
var ServerData map[string]*State = make(map[string]*State)

// Server describes a single server instance in the cluster
type Server struct {
	IP   string `json:"ip"`
	Port string `json:"port"`
}

// Configuration is the entire config file description
type Configuration struct {
	Servers map[int]Server `json:"servers"`
}

// LogData is an instance of a single log
type LogData struct {
	Term    int    `json:"term"`
	Command string `json:"command"`
}

// State has all the data on the
// servers state.
type State struct {
	// can be follower, leader or candidate
	// all servers start as a follower, if they
	// dont hear from a leader, they can become
	// candidates. Leaders are elected from the
	// leader election process.
	Name        string `json:"name"`
	ID          int    `json:"ID"`
	CurrentTerm int    `json:"currentTerm"`
	// VotedFor maintains the ID of the voted
	// server; -1 if its leader, -2 at init
	VotedFor int `json:"votedFor"`
	// Log is the command received by the leader.
	// each entry contains the term and the command.
	Log []LogData `json:"log"`
	// above 4 variables are persistent in the server
	// CommitIndex maintains the highest log entry
	// that is known to be committed.
	CommitIndex int `json:"commitIndex"`
	// LastApplied is the highest log entry applied
	// to the state machine
	LastApplied int `json:"lastApplied"`
	// above 2 variables are volatile on all servers
	// NextIndex maintains a list of the next log
	// entry to be sent to the followers.
	NextIndex []int `json:"nextIndex"`
	// MatchIndex maintains the highest log entry
	// that is known to be replicated on the server
	MatchIndex []int `json:"matchIndex"`
	// above 2 variables are volatile only int the
	// leader and for each follower. Its also
	// re-init after each election.
	Lock sync.Mutex
}

// RaftServer describes a single raft server
type RaftServer struct {
	ServerState State         `json:"serverState"`
	Config      Configuration `json:"config"`
}

// URLResponse facilitates responses for ConcurrentReqRes
type URLResponse struct {
	URL string         `json:"URL"`
	Res *http.Response `json:"res"`
}
