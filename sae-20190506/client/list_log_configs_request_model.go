// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLogConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListLogConfigsRequest
	GetAppId() *string
	SetCurrentPage(v int32) *ListLogConfigsRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *ListLogConfigsRequest
	GetPageSize() *int32
}

type ListLogConfigsRequest struct {
	// 10
	//
	// This parameter is required.
	//
	// example:
	//
	// 56f77b65-788d-442a-9885-7f20d91f****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 1
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListLogConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLogConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListLogConfigsRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListLogConfigsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListLogConfigsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLogConfigsRequest) SetAppId(v string) *ListLogConfigsRequest {
	s.AppId = &v
	return s
}

func (s *ListLogConfigsRequest) SetCurrentPage(v int32) *ListLogConfigsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListLogConfigsRequest) SetPageSize(v int32) *ListLogConfigsRequest {
	s.PageSize = &v
	return s
}

func (s *ListLogConfigsRequest) Validate() error {
	return dara.Validate(s)
}
