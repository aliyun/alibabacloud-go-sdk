// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataSessionStatDbStatsValue interface {
	dara.Model
	String() string
	GoString() string
	SetActiveCount(v int64) *DataSessionStatDbStatsValue
	GetActiveCount() *int64
	SetTotalCount(v int64) *DataSessionStatDbStatsValue
	GetTotalCount() *int64
}

type DataSessionStatDbStatsValue struct {
	// The number of active namespaces.
	//
	// example:
	//
	// 0
	ActiveCount *int64 `json:"ActiveCount,omitempty" xml:"ActiveCount,omitempty"`
	// The total number of namespaces.
	//
	// example:
	//
	// 11
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DataSessionStatDbStatsValue) String() string {
	return dara.Prettify(s)
}

func (s DataSessionStatDbStatsValue) GoString() string {
	return s.String()
}

func (s *DataSessionStatDbStatsValue) GetActiveCount() *int64 {
	return s.ActiveCount
}

func (s *DataSessionStatDbStatsValue) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DataSessionStatDbStatsValue) SetActiveCount(v int64) *DataSessionStatDbStatsValue {
	s.ActiveCount = &v
	return s
}

func (s *DataSessionStatDbStatsValue) SetTotalCount(v int64) *DataSessionStatDbStatsValue {
	s.TotalCount = &v
	return s
}

func (s *DataSessionStatDbStatsValue) Validate() error {
	return dara.Validate(s)
}
