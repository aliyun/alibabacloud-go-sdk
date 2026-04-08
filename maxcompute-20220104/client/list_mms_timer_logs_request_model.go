// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMmsTimerLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNum(v int32) *ListMmsTimerLogsRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListMmsTimerLogsRequest
	GetPageSize() *int32
}

type ListMmsTimerLogsRequest struct {
	// example:
	//
	// 1
	PageNum *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListMmsTimerLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMmsTimerLogsRequest) GoString() string {
	return s.String()
}

func (s *ListMmsTimerLogsRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListMmsTimerLogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMmsTimerLogsRequest) SetPageNum(v int32) *ListMmsTimerLogsRequest {
	s.PageNum = &v
	return s
}

func (s *ListMmsTimerLogsRequest) SetPageSize(v int32) *ListMmsTimerLogsRequest {
	s.PageSize = &v
	return s
}

func (s *ListMmsTimerLogsRequest) Validate() error {
	return dara.Validate(s)
}
