// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRLProgressSync interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v float64) *RLProgressSync
	GetCost() *float64
	SetState(v string) *RLProgressSync
	GetState() *string
}

type RLProgressSync struct {
	// The parameter synchronization duration in seconds. This property has a value only when State is end.
	//
	// example:
	//
	// 1.5
	Cost *float64 `json:"Cost,omitempty" xml:"Cost,omitempty"`
	// begin / end
	//
	// example:
	//
	// end
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s RLProgressSync) String() string {
	return dara.Prettify(s)
}

func (s RLProgressSync) GoString() string {
	return s.String()
}

func (s *RLProgressSync) GetCost() *float64 {
	return s.Cost
}

func (s *RLProgressSync) GetState() *string {
	return s.State
}

func (s *RLProgressSync) SetCost(v float64) *RLProgressSync {
	s.Cost = &v
	return s
}

func (s *RLProgressSync) SetState(v string) *RLProgressSync {
	s.State = &v
	return s
}

func (s *RLProgressSync) Validate() error {
	return dara.Validate(s)
}
