// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCheckItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListCheckItemRequest
	GetCurrentPage() *int32
	SetLang(v string) *ListCheckItemRequest
	GetLang() *string
	SetPageSize(v int32) *ListCheckItemRequest
	GetPageSize() *int32
	SetTaskSources(v []*string) *ListCheckItemRequest
	GetTaskSources() []*string
}

type ListCheckItemRequest struct {
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid value:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries per page. Default value: **20**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// List of task sources.
	TaskSources []*string `json:"TaskSources,omitempty" xml:"TaskSources,omitempty" type:"Repeated"`
}

func (s ListCheckItemRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCheckItemRequest) GoString() string {
	return s.String()
}

func (s *ListCheckItemRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListCheckItemRequest) GetLang() *string {
	return s.Lang
}

func (s *ListCheckItemRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCheckItemRequest) GetTaskSources() []*string {
	return s.TaskSources
}

func (s *ListCheckItemRequest) SetCurrentPage(v int32) *ListCheckItemRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListCheckItemRequest) SetLang(v string) *ListCheckItemRequest {
	s.Lang = &v
	return s
}

func (s *ListCheckItemRequest) SetPageSize(v int32) *ListCheckItemRequest {
	s.PageSize = &v
	return s
}

func (s *ListCheckItemRequest) SetTaskSources(v []*string) *ListCheckItemRequest {
	s.TaskSources = v
	return s
}

func (s *ListCheckItemRequest) Validate() error {
	return dara.Validate(s)
}
