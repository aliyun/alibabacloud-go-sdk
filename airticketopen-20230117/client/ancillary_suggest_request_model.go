// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAncillarySuggestRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSolutionId(v string) *AncillarySuggestRequest
	GetSolutionId() *string
}

type AncillarySuggestRequest struct {
	// solution_id returned by enrich
	//
	// This parameter is required.
	//
	// example:
	//
	// eJwz8DeySEo0NjQ01TU3TU7TNTFINNO1SE5O0jVKM0hKNjEwTElLNYwz0A32cNT1dfPVNTIwMjYwNjRQ8/A3NLI01Q0Ic0cRBwBVFxJJ
	SolutionId *string `json:"solution_id,omitempty" xml:"solution_id,omitempty"`
}

func (s AncillarySuggestRequest) String() string {
	return dara.Prettify(s)
}

func (s AncillarySuggestRequest) GoString() string {
	return s.String()
}

func (s *AncillarySuggestRequest) GetSolutionId() *string {
	return s.SolutionId
}

func (s *AncillarySuggestRequest) SetSolutionId(v string) *AncillarySuggestRequest {
	s.SolutionId = &v
	return s
}

func (s *AncillarySuggestRequest) Validate() error {
	return dara.Validate(s)
}
