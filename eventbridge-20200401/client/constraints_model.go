// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConstraints interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v int64) *Constraints
	GetBeginTime() *int64
	SetEndTime(v int64) *Constraints
	GetEndTime() *int64
}

type Constraints struct {
	// example:
	//
	// 1717200000000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// example:
	//
	// 1717200000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s Constraints) String() string {
	return dara.Prettify(s)
}

func (s Constraints) GoString() string {
	return s.String()
}

func (s *Constraints) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *Constraints) GetEndTime() *int64 {
	return s.EndTime
}

func (s *Constraints) SetBeginTime(v int64) *Constraints {
	s.BeginTime = &v
	return s
}

func (s *Constraints) SetEndTime(v int64) *Constraints {
	s.EndTime = &v
	return s
}

func (s *Constraints) Validate() error {
	return dara.Validate(s)
}
