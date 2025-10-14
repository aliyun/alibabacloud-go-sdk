// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPricingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSolutionId(v string) *PricingRequest
	GetSolutionId() *string
}

type PricingRequest struct {
	// solution_id returned by Enrich
	//
	// This parameter is required.
	//
	// example:
	//
	// eJwz8DeySEo0NjQ01TU3TU7TNTFINNO1SE5O0jVKM0hKNjEwTElLNYwz0A32cNT1dfPVNTIwMjYwNjRQ8/A3NLI01Q0Ic0cRBwBVFxJJ
	SolutionId *string `json:"solution_id,omitempty" xml:"solution_id,omitempty"`
}

func (s PricingRequest) String() string {
	return dara.Prettify(s)
}

func (s PricingRequest) GoString() string {
	return s.String()
}

func (s *PricingRequest) GetSolutionId() *string {
	return s.SolutionId
}

func (s *PricingRequest) SetSolutionId(v string) *PricingRequest {
	s.SolutionId = &v
	return s
}

func (s *PricingRequest) Validate() error {
	return dara.Validate(s)
}
