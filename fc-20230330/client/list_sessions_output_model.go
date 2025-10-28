// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSessionsOutput interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListSessionsOutput
	GetNextToken() *string
	SetSessions(v []*Session) *ListSessionsOutput
	GetSessions() []*Session
}

type ListSessionsOutput struct {
	// example:
	//
	// MTIzNCNhYmM=
	NextToken *string    `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Sessions  []*Session `json:"sessions" xml:"sessions" type:"Repeated"`
}

func (s ListSessionsOutput) String() string {
	return dara.Prettify(s)
}

func (s ListSessionsOutput) GoString() string {
	return s.String()
}

func (s *ListSessionsOutput) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSessionsOutput) GetSessions() []*Session {
	return s.Sessions
}

func (s *ListSessionsOutput) SetNextToken(v string) *ListSessionsOutput {
	s.NextToken = &v
	return s
}

func (s *ListSessionsOutput) SetSessions(v []*Session) *ListSessionsOutput {
	s.Sessions = v
	return s
}

func (s *ListSessionsOutput) Validate() error {
	if s.Sessions != nil {
		for _, item := range s.Sessions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
