// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStatisticsDetails interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int64) *StatisticsDetails
	GetCount() *int64
	SetIdleNum(v int64) *StatisticsDetails
	GetIdleNum() *int64
}

type StatisticsDetails struct {
	Count   *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	IdleNum *int64 `json:"IdleNum,omitempty" xml:"IdleNum,omitempty"`
}

func (s StatisticsDetails) String() string {
	return dara.Prettify(s)
}

func (s StatisticsDetails) GoString() string {
	return s.String()
}

func (s *StatisticsDetails) GetCount() *int64 {
	return s.Count
}

func (s *StatisticsDetails) GetIdleNum() *int64 {
	return s.IdleNum
}

func (s *StatisticsDetails) SetCount(v int64) *StatisticsDetails {
	s.Count = &v
	return s
}

func (s *StatisticsDetails) SetIdleNum(v int64) *StatisticsDetails {
	s.IdleNum = &v
	return s
}

func (s *StatisticsDetails) Validate() error {
	return dara.Validate(s)
}
