// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBrowsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBrowserName(v string) *ListBrowsersRequest
	GetBrowserName() *string
	SetPageNumber(v int32) *ListBrowsersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListBrowsersRequest
	GetPageSize() *int32
	SetStatus(v string) *ListBrowsersRequest
	GetStatus() *string
}

type ListBrowsersRequest struct {
	// 根据浏览器实例名称进行模糊匹配过滤
	//
	// example:
	//
	// browser
	BrowserName *string `json:"browserName,omitempty" xml:"browserName,omitempty"`
	// 当前页码，从1开始计数
	//
	// example:
	//
	// 10
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 每页返回的记录数量
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 根据浏览器实例的运行状态进行过滤，可选值：CREATING、READY、DELETING等
	//
	// if can be null:
	// true
	//
	// example:
	//
	// CREATING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListBrowsersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBrowsersRequest) GoString() string {
	return s.String()
}

func (s *ListBrowsersRequest) GetBrowserName() *string {
	return s.BrowserName
}

func (s *ListBrowsersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListBrowsersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBrowsersRequest) GetStatus() *string {
	return s.Status
}

func (s *ListBrowsersRequest) SetBrowserName(v string) *ListBrowsersRequest {
	s.BrowserName = &v
	return s
}

func (s *ListBrowsersRequest) SetPageNumber(v int32) *ListBrowsersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListBrowsersRequest) SetPageSize(v int32) *ListBrowsersRequest {
	s.PageSize = &v
	return s
}

func (s *ListBrowsersRequest) SetStatus(v string) *ListBrowsersRequest {
	s.Status = &v
	return s
}

func (s *ListBrowsersRequest) Validate() error {
	return dara.Validate(s)
}
