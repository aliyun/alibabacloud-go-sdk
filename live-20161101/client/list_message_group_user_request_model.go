// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMessageGroupUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListMessageGroupUserRequest
	GetAppId() *string
	SetGroupId(v string) *ListMessageGroupUserRequest
	GetGroupId() *string
	SetPageNum(v int32) *ListMessageGroupUserRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListMessageGroupUserRequest
	GetPageSize() *int32
	SetSortType(v int32) *ListMessageGroupUserRequest
	GetSortType() *int32
}

type ListMessageGroupUserRequest struct {
	// The ID of the interactive messaging application.
	//
	// This parameter is required.
	//
	// example:
	//
	// VKL3***
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the message group.
	//
	// This parameter is required.
	//
	// example:
	//
	// AE35-****-T95F
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The number of the page to return. Default value: 1.
	//
	// Valid values: 1 to 100000.
	//
	// example:
	//
	// 10
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of members to return on each page. Default value: 20.
	//
	// Valid values: 1 to 50.
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
}

func (s ListMessageGroupUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMessageGroupUserRequest) GoString() string {
	return s.String()
}

func (s *ListMessageGroupUserRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListMessageGroupUserRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ListMessageGroupUserRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListMessageGroupUserRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMessageGroupUserRequest) GetSortType() *int32 {
	return s.SortType
}

func (s *ListMessageGroupUserRequest) SetAppId(v string) *ListMessageGroupUserRequest {
	s.AppId = &v
	return s
}

func (s *ListMessageGroupUserRequest) SetGroupId(v string) *ListMessageGroupUserRequest {
	s.GroupId = &v
	return s
}

func (s *ListMessageGroupUserRequest) SetPageNum(v int32) *ListMessageGroupUserRequest {
	s.PageNum = &v
	return s
}

func (s *ListMessageGroupUserRequest) SetPageSize(v int32) *ListMessageGroupUserRequest {
	s.PageSize = &v
	return s
}

func (s *ListMessageGroupUserRequest) SetSortType(v int32) *ListMessageGroupUserRequest {
	s.SortType = &v
	return s
}

func (s *ListMessageGroupUserRequest) Validate() error {
	return dara.Validate(s)
}
