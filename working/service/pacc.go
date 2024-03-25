package service

import (
	"encoding/json"
	"io"
	"net/http"
	"sync"
)

type scrv struct {
	mu    *sync.RWMutex
	stats map[uint]uint
}

func (s *scrv) Vote(w http.ResponseWriter, r *http.Request) {
	// check request method
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	req := struct {
		CandidateID uint   `json:"Candidate_id"`
		Passport    string `json:"passport"`
	}{}
	// get request body
	raw, err := io.ReadAll(r.Body) //интерфейс потокового чтения
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(raw, &req); err != nil {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}
	//walidate
	if len(req.Passport) == 0 || req.CandidateID == 0 {
		w.WriteHeader(http.StatusBadRequest)

		return
	}
	s.mu.Lock()
	s.stats[req.CandidateID]++
	s.mu.Unlock()

	w.WriteHeader(http.StatusOK)

}

func (s *scrv) Poll(w http.ResponseWriter, r *http.Request) {}

func New() scrv {
	return scrv{
		mu:    &sync.RWMutex{},
		stats: make(map[uint]uint),
	}
}
