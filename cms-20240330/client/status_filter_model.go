// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStatusFilter interface {
	dara.Model
	String() string
	GoString() string
	SetEq(v string) *StatusFilter
	GetEq() *string
}

type StatusFilter struct {
	// The exact match condition for the alert status. Only alert rules whose status equals the specified value are returned. Valid values:
	//
	// - Alarm: The alert rule is in the alerting state.
	//
	// - Ok: The alert rule is in the normal state.
	//
	// - InsufficientData: Insufficient data is available.
	//
	// example:
	//
	// Alarm
	Eq *string `json:"eq,omitempty" xml:"eq,omitempty"`
}

func (s StatusFilter) String() string {
	return dara.Prettify(s)
}

func (s StatusFilter) GoString() string {
	return s.String()
}

func (s *StatusFilter) GetEq() *string {
	return s.Eq
}

func (s *StatusFilter) SetEq(v string) *StatusFilter {
	s.Eq = &v
	return s
}

func (s *StatusFilter) Validate() error {
	return dara.Validate(s)
}
