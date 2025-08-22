// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *ListAppsRequest
	GetAppName() *string
	SetClusterId(v string) *ListAppsRequest
	GetClusterId() *string
	SetPageNum(v int32) *ListAppsRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListAppsRequest
	GetPageSize() *int32
	SetTitle(v string) *ListAppsRequest
	GetTitle() *string
}

type ListAppsRequest struct {
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Title    *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListAppsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppsRequest) GoString() string {
	return s.String()
}

func (s *ListAppsRequest) GetAppName() *string {
	return s.AppName
}

func (s *ListAppsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListAppsRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListAppsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAppsRequest) GetTitle() *string {
	return s.Title
}

func (s *ListAppsRequest) SetAppName(v string) *ListAppsRequest {
	s.AppName = &v
	return s
}

func (s *ListAppsRequest) SetClusterId(v string) *ListAppsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListAppsRequest) SetPageNum(v int32) *ListAppsRequest {
	s.PageNum = &v
	return s
}

func (s *ListAppsRequest) SetPageSize(v int32) *ListAppsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppsRequest) SetTitle(v string) *ListAppsRequest {
	s.Title = &v
	return s
}

func (s *ListAppsRequest) Validate() error {
	return dara.Validate(s)
}
