package participation

import (
	"fmt"

	"github.com/iotaledger/hive.go/serializer"
)

// NewBallotBuilder creates a new BallotBuilder.
func NewBallotBuilder() *BallotBuilder {
	return &BallotBuilder{
		ballot: &Ballot{},
	}
}

// BallotBuilder is used to easily build up a Ballot.
type BallotBuilder struct {
	ballot *Ballot
}

// AddQuestion adds the given question to the Ballot.
func (qb *BallotBuilder) AddQuestion(entry *Question) *BallotBuilder {
	qb.ballot.Questions = append(qb.ballot.Questions, entry)
	return qb
}

// Build builds the Ballot.
func (qb *BallotBuilder) Build() (*Ballot, error) {
	if _, err := qb.ballot.Serialize(serializer.DeSeriModePerformValidation); err != nil {
		return nil, fmt.Errorf("unable to build ballot: %w", err)
	}
	return qb.ballot, nil
}
