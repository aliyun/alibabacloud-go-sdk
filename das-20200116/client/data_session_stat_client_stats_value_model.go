// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataSessionStatClientStatsValue interface {
	dara.Model
	String() string
	GoString() string
	SetActiveCount(v int64) *DataSessionStatClientStatsValue
	GetActiveCount() *int64
	SetTotalCount(v int64) *DataSessionStatClientStatsValue
	GetTotalCount() *int64
}

type DataSessionStatClientStatsValue struct {
	// The number of clients whose IP addresses are active.
	//
	// example:
	//
	// 0
	ActiveCount *int64 `json:"ActiveCount,omitempty" xml:"ActiveCount,omitempty"`
	// The total number of IP addresses of clients.
	//
	// example:
	//
	// 11
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DataSessionStatClientStatsValue) String() string {
	return dara.Prettify(s)
}

func (s DataSessionStatClientStatsValue) GoString() string {
	return s.String()
}

func (s *DataSessionStatClientStatsValue) GetActiveCount() *int64 {
	return s.ActiveCount
}

func (s *DataSessionStatClientStatsValue) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DataSessionStatClientStatsValue) SetActiveCount(v int64) *DataSessionStatClientStatsValue {
	s.ActiveCount = &v
	return s
}

func (s *DataSessionStatClientStatsValue) SetTotalCount(v int64) *DataSessionStatClientStatsValue {
	s.TotalCount = &v
	return s
}

func (s *DataSessionStatClientStatsValue) Validate() error {
	return dara.Validate(s)
}
