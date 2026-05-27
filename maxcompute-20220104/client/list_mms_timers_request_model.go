// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMmsTimersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListMmsTimersRequest
	GetName() *string
	SetPageNum(v int32) *ListMmsTimersRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListMmsTimersRequest
	GetPageSize() *int32
	SetSrcDbName(v string) *ListMmsTimersRequest
	GetSrcDbName() *string
	SetSrcTableName(v string) *ListMmsTimersRequest
	GetSrcTableName() *string
	SetStopped(v bool) *ListMmsTimersRequest
	GetStopped() *bool
}

type ListMmsTimersRequest struct {
	// example:
	//
	// demo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// test_db_1
	SrcDbName *string `json:"srcDbName,omitempty" xml:"srcDbName,omitempty"`
	// example:
	//
	// test_table_1
	SrcTableName *string `json:"srcTableName,omitempty" xml:"srcTableName,omitempty"`
	// example:
	//
	// false
	Stopped *bool `json:"stopped,omitempty" xml:"stopped,omitempty"`
}

func (s ListMmsTimersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMmsTimersRequest) GoString() string {
	return s.String()
}

func (s *ListMmsTimersRequest) GetName() *string {
	return s.Name
}

func (s *ListMmsTimersRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListMmsTimersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMmsTimersRequest) GetSrcDbName() *string {
	return s.SrcDbName
}

func (s *ListMmsTimersRequest) GetSrcTableName() *string {
	return s.SrcTableName
}

func (s *ListMmsTimersRequest) GetStopped() *bool {
	return s.Stopped
}

func (s *ListMmsTimersRequest) SetName(v string) *ListMmsTimersRequest {
	s.Name = &v
	return s
}

func (s *ListMmsTimersRequest) SetPageNum(v int32) *ListMmsTimersRequest {
	s.PageNum = &v
	return s
}

func (s *ListMmsTimersRequest) SetPageSize(v int32) *ListMmsTimersRequest {
	s.PageSize = &v
	return s
}

func (s *ListMmsTimersRequest) SetSrcDbName(v string) *ListMmsTimersRequest {
	s.SrcDbName = &v
	return s
}

func (s *ListMmsTimersRequest) SetSrcTableName(v string) *ListMmsTimersRequest {
	s.SrcTableName = &v
	return s
}

func (s *ListMmsTimersRequest) SetStopped(v bool) *ListMmsTimersRequest {
	s.Stopped = &v
	return s
}

func (s *ListMmsTimersRequest) Validate() error {
	return dara.Validate(s)
}
