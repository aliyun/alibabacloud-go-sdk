// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListMessageRequest
	GetAppId() *string
	SetGroupId(v string) *ListMessageRequest
	GetGroupId() *string
	SetPageNum(v int32) *ListMessageRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListMessageRequest
	GetPageSize() *int32
	SetSortType(v int32) *ListMessageRequest
	GetSortType() *int32
	SetType(v int32) *ListMessageRequest
	GetType() *int32
}

type ListMessageRequest struct {
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
	// The number of the page to return. Default value: 1. Valid values: 1 to 100000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries to return on each page. Default value: 20. Valid values: 1 to 50.
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
	// The type of the custom message. Valid values: integers greater than 10000.
	//
	// example:
	//
	// 10002
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMessageRequest) GoString() string {
	return s.String()
}

func (s *ListMessageRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListMessageRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ListMessageRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListMessageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMessageRequest) GetSortType() *int32 {
	return s.SortType
}

func (s *ListMessageRequest) GetType() *int32 {
	return s.Type
}

func (s *ListMessageRequest) SetAppId(v string) *ListMessageRequest {
	s.AppId = &v
	return s
}

func (s *ListMessageRequest) SetGroupId(v string) *ListMessageRequest {
	s.GroupId = &v
	return s
}

func (s *ListMessageRequest) SetPageNum(v int32) *ListMessageRequest {
	s.PageNum = &v
	return s
}

func (s *ListMessageRequest) SetPageSize(v int32) *ListMessageRequest {
	s.PageSize = &v
	return s
}

func (s *ListMessageRequest) SetSortType(v int32) *ListMessageRequest {
	s.SortType = &v
	return s
}

func (s *ListMessageRequest) SetType(v int32) *ListMessageRequest {
	s.Type = &v
	return s
}

func (s *ListMessageRequest) Validate() error {
	return dara.Validate(s)
}
