// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMessageAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *QueryMessageAppRequest
	GetAppId() *string
	SetAppName(v string) *QueryMessageAppRequest
	GetAppName() *string
	SetPageNum(v int32) *QueryMessageAppRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryMessageAppRequest
	GetPageSize() *int32
	SetSortType(v int32) *QueryMessageAppRequest
	GetSortType() *int32
}

type QueryMessageAppRequest struct {
	// The ID of the interactive messaging application.
	//
	// example:
	//
	// VKL3***
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the interactive messaging application.
	//
	// example:
	//
	// testApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The number of the page to return. Default value: 1. Valid values: 1 to 100000.
	//
	// example:
	//
	// 10
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of applications to return on each page. Default value: 20. Valid values: 1 to 50.
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

func (s QueryMessageAppRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMessageAppRequest) GoString() string {
	return s.String()
}

func (s *QueryMessageAppRequest) GetAppId() *string {
	return s.AppId
}

func (s *QueryMessageAppRequest) GetAppName() *string {
	return s.AppName
}

func (s *QueryMessageAppRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryMessageAppRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryMessageAppRequest) GetSortType() *int32 {
	return s.SortType
}

func (s *QueryMessageAppRequest) SetAppId(v string) *QueryMessageAppRequest {
	s.AppId = &v
	return s
}

func (s *QueryMessageAppRequest) SetAppName(v string) *QueryMessageAppRequest {
	s.AppName = &v
	return s
}

func (s *QueryMessageAppRequest) SetPageNum(v int32) *QueryMessageAppRequest {
	s.PageNum = &v
	return s
}

func (s *QueryMessageAppRequest) SetPageSize(v int32) *QueryMessageAppRequest {
	s.PageSize = &v
	return s
}

func (s *QueryMessageAppRequest) SetSortType(v int32) *QueryMessageAppRequest {
	s.SortType = &v
	return s
}

func (s *QueryMessageAppRequest) Validate() error {
	return dara.Validate(s)
}
