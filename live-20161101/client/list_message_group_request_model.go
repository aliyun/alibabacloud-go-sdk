// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMessageGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListMessageGroupRequest
	GetAppId() *string
	SetPageNum(v int32) *ListMessageGroupRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListMessageGroupRequest
	GetPageSize() *int32
	SetSortType(v int32) *ListMessageGroupRequest
	GetSortType() *int32
	SetUserId(v string) *ListMessageGroupRequest
	GetUserId() *string
}

type ListMessageGroupRequest struct {
	// The ID of the interactive messaging application.
	//
	// example:
	//
	// VKL3***
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The number of the page to return. Default value: 1. Valid values: 1 to 100000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of message groups to return on each page. Default value: 20.
	//
	// Valid values: 1 to 50.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The sort order. Valid values:
	//
	// 	- 0: ascending order by time
	//
	// 	- 1: descending order by time
	//
	// example:
	//
	// 1
	SortType *int32 `json:"SortType,omitempty" xml:"SortType,omitempty"`
	// The ID of the user. Each user has a unique ID in the application. You can specify multiple user IDs.
	//
	// example:
	//
	// de1**a0
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListMessageGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMessageGroupRequest) GoString() string {
	return s.String()
}

func (s *ListMessageGroupRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListMessageGroupRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListMessageGroupRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMessageGroupRequest) GetSortType() *int32 {
	return s.SortType
}

func (s *ListMessageGroupRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListMessageGroupRequest) SetAppId(v string) *ListMessageGroupRequest {
	s.AppId = &v
	return s
}

func (s *ListMessageGroupRequest) SetPageNum(v int32) *ListMessageGroupRequest {
	s.PageNum = &v
	return s
}

func (s *ListMessageGroupRequest) SetPageSize(v int32) *ListMessageGroupRequest {
	s.PageSize = &v
	return s
}

func (s *ListMessageGroupRequest) SetSortType(v int32) *ListMessageGroupRequest {
	s.SortType = &v
	return s
}

func (s *ListMessageGroupRequest) SetUserId(v string) *ListMessageGroupRequest {
	s.UserId = &v
	return s
}

func (s *ListMessageGroupRequest) Validate() error {
	return dara.Validate(s)
}
