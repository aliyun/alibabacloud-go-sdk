// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAccountFlowListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDayNum(v int32) *AccountFlowListRequest
	GetDayNum() *int32
	SetPageIndex(v int32) *AccountFlowListRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *AccountFlowListRequest
	GetPageSize() *int32
	SetUtcBeginTime(v int64) *AccountFlowListRequest
	GetUtcBeginTime() *int64
}

type AccountFlowListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 0
	DayNum *int32 `json:"day_num,omitempty" xml:"day_num,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1677427200000
	UtcBeginTime *int64 `json:"utc_begin_time,omitempty" xml:"utc_begin_time,omitempty"`
}

func (s AccountFlowListRequest) String() string {
	return dara.Prettify(s)
}

func (s AccountFlowListRequest) GoString() string {
	return s.String()
}

func (s *AccountFlowListRequest) GetDayNum() *int32 {
	return s.DayNum
}

func (s *AccountFlowListRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *AccountFlowListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *AccountFlowListRequest) GetUtcBeginTime() *int64 {
	return s.UtcBeginTime
}

func (s *AccountFlowListRequest) SetDayNum(v int32) *AccountFlowListRequest {
	s.DayNum = &v
	return s
}

func (s *AccountFlowListRequest) SetPageIndex(v int32) *AccountFlowListRequest {
	s.PageIndex = &v
	return s
}

func (s *AccountFlowListRequest) SetPageSize(v int32) *AccountFlowListRequest {
	s.PageSize = &v
	return s
}

func (s *AccountFlowListRequest) SetUtcBeginTime(v int64) *AccountFlowListRequest {
	s.UtcBeginTime = &v
	return s
}

func (s *AccountFlowListRequest) Validate() error {
	return dara.Validate(s)
}
