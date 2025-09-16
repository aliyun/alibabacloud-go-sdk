// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListLogsRequest
	GetEndTime() *string
	SetPageNum(v string) *ListLogsRequest
	GetPageNum() *string
	SetPageSize(v string) *ListLogsRequest
	GetPageSize() *string
	SetQuery(v string) *ListLogsRequest
	GetQuery() *string
	SetStartTime(v string) *ListLogsRequest
	GetStartTime() *string
	SetType(v string) *ListLogsRequest
	GetType() *string
}

type ListLogsRequest struct {
	// The end tim. The value is a timestamp in seconds.
	//
	// example:
	//
	// 1710432000
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The number of entries per num. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *string `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The query statement
	//
	// example:
	//
	// status: 200 AND totalTime > 0.01
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// The start time. The value is a timestamp in seconds.
	//
	// example:
	//
	// 1706340600
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// -push   -select
	//
	// example:
	//
	// push
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLogsRequest) GoString() string {
	return s.String()
}

func (s *ListLogsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListLogsRequest) GetPageNum() *string {
	return s.PageNum
}

func (s *ListLogsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListLogsRequest) GetQuery() *string {
	return s.Query
}

func (s *ListLogsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListLogsRequest) GetType() *string {
	return s.Type
}

func (s *ListLogsRequest) SetEndTime(v string) *ListLogsRequest {
	s.EndTime = &v
	return s
}

func (s *ListLogsRequest) SetPageNum(v string) *ListLogsRequest {
	s.PageNum = &v
	return s
}

func (s *ListLogsRequest) SetPageSize(v string) *ListLogsRequest {
	s.PageSize = &v
	return s
}

func (s *ListLogsRequest) SetQuery(v string) *ListLogsRequest {
	s.Query = &v
	return s
}

func (s *ListLogsRequest) SetStartTime(v string) *ListLogsRequest {
	s.StartTime = &v
	return s
}

func (s *ListLogsRequest) SetType(v string) *ListLogsRequest {
	s.Type = &v
	return s
}

func (s *ListLogsRequest) Validate() error {
	return dara.Validate(s)
}
