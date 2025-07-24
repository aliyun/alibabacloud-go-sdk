// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValueAddedCredits interface {
	dara.Model
	String() string
	GoString() string
	SetAdvanced(v int32) *ValueAddedCredits
	GetAdvanced() *int32
	SetSummary(v int32) *ValueAddedCredits
	GetSummary() *int32
}

type ValueAddedCredits struct {
	Advanced *int32 `json:"advanced,omitempty" xml:"advanced,omitempty"`
	Summary  *int32 `json:"summary,omitempty" xml:"summary,omitempty"`
}

func (s ValueAddedCredits) String() string {
	return dara.Prettify(s)
}

func (s ValueAddedCredits) GoString() string {
	return s.String()
}

func (s *ValueAddedCredits) GetAdvanced() *int32 {
	return s.Advanced
}

func (s *ValueAddedCredits) GetSummary() *int32 {
	return s.Summary
}

func (s *ValueAddedCredits) SetAdvanced(v int32) *ValueAddedCredits {
	s.Advanced = &v
	return s
}

func (s *ValueAddedCredits) SetSummary(v int32) *ValueAddedCredits {
	s.Summary = &v
	return s
}

func (s *ValueAddedCredits) Validate() error {
	return dara.Validate(s)
}
