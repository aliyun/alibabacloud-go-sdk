// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHealthSummary interface {
	dara.Model
	String() string
	GoString() string
	SetBadCount(v int64) *HealthSummary
	GetBadCount() *int64
	SetGoodCount(v int64) *HealthSummary
	GetGoodCount() *int64
	SetNoneCount(v int64) *HealthSummary
	GetNoneCount() *int64
	SetStoppedCount(v int64) *HealthSummary
	GetStoppedCount() *int64
	SetTotalCount(v int64) *HealthSummary
	GetTotalCount() *int64
	SetUnknownCount(v int64) *HealthSummary
	GetUnknownCount() *int64
	SetWarningCount(v int64) *HealthSummary
	GetWarningCount() *int64
}

type HealthSummary struct {
	// example:
	//
	// 2
	BadCount *int64 `json:"BadCount,omitempty" xml:"BadCount,omitempty"`
	// example:
	//
	// 2
	GoodCount *int64 `json:"GoodCount,omitempty" xml:"GoodCount,omitempty"`
	// example:
	//
	// 0
	NoneCount *int64 `json:"NoneCount,omitempty" xml:"NoneCount,omitempty"`
	// example:
	//
	// 2
	StoppedCount *int64 `json:"StoppedCount,omitempty" xml:"StoppedCount,omitempty"`
	// example:
	//
	// 8
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 0
	UnknownCount *int64 `json:"UnknownCount,omitempty" xml:"UnknownCount,omitempty"`
	// example:
	//
	// 2
	WarningCount *int64 `json:"WarningCount,omitempty" xml:"WarningCount,omitempty"`
}

func (s HealthSummary) String() string {
	return dara.Prettify(s)
}

func (s HealthSummary) GoString() string {
	return s.String()
}

func (s *HealthSummary) GetBadCount() *int64 {
	return s.BadCount
}

func (s *HealthSummary) GetGoodCount() *int64 {
	return s.GoodCount
}

func (s *HealthSummary) GetNoneCount() *int64 {
	return s.NoneCount
}

func (s *HealthSummary) GetStoppedCount() *int64 {
	return s.StoppedCount
}

func (s *HealthSummary) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *HealthSummary) GetUnknownCount() *int64 {
	return s.UnknownCount
}

func (s *HealthSummary) GetWarningCount() *int64 {
	return s.WarningCount
}

func (s *HealthSummary) SetBadCount(v int64) *HealthSummary {
	s.BadCount = &v
	return s
}

func (s *HealthSummary) SetGoodCount(v int64) *HealthSummary {
	s.GoodCount = &v
	return s
}

func (s *HealthSummary) SetNoneCount(v int64) *HealthSummary {
	s.NoneCount = &v
	return s
}

func (s *HealthSummary) SetStoppedCount(v int64) *HealthSummary {
	s.StoppedCount = &v
	return s
}

func (s *HealthSummary) SetTotalCount(v int64) *HealthSummary {
	s.TotalCount = &v
	return s
}

func (s *HealthSummary) SetUnknownCount(v int64) *HealthSummary {
	s.UnknownCount = &v
	return s
}

func (s *HealthSummary) SetWarningCount(v int64) *HealthSummary {
	s.WarningCount = &v
	return s
}

func (s *HealthSummary) Validate() error {
	return dara.Validate(s)
}
