// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTaskLogListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v string) *DescribeTaskLogListRequest
	GetCurrentPage() *string
	SetIsPage(v bool) *DescribeTaskLogListRequest
	GetIsPage() *bool
	SetLang(v string) *DescribeTaskLogListRequest
	GetLang() *string
	SetPageSize(v string) *DescribeTaskLogListRequest
	GetPageSize() *string
	SetTaskId(v string) *DescribeTaskLogListRequest
	GetTaskId() *string
	SetTaskLogId(v string) *DescribeTaskLogListRequest
	GetTaskLogId() *string
	SetRegId(v string) *DescribeTaskLogListRequest
	GetRegId() *string
}

type DescribeTaskLogListRequest struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Specifies whether to enable pagination.
	//
	// example:
	//
	// true
	IsPage *bool `json:"IsPage,omitempty" xml:"IsPage,omitempty"`
	// The language type for the request and response messages. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 18044
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task log ID.
	//
	// example:
	//
	// 107
	TaskLogId *string `json:"TaskLogId,omitempty" xml:"TaskLogId,omitempty"`
	// The region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeTaskLogListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTaskLogListRequest) GoString() string {
	return s.String()
}

func (s *DescribeTaskLogListRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeTaskLogListRequest) GetIsPage() *bool {
	return s.IsPage
}

func (s *DescribeTaskLogListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeTaskLogListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeTaskLogListRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeTaskLogListRequest) GetTaskLogId() *string {
	return s.TaskLogId
}

func (s *DescribeTaskLogListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeTaskLogListRequest) SetCurrentPage(v string) *DescribeTaskLogListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeTaskLogListRequest) SetIsPage(v bool) *DescribeTaskLogListRequest {
	s.IsPage = &v
	return s
}

func (s *DescribeTaskLogListRequest) SetLang(v string) *DescribeTaskLogListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeTaskLogListRequest) SetPageSize(v string) *DescribeTaskLogListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTaskLogListRequest) SetTaskId(v string) *DescribeTaskLogListRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeTaskLogListRequest) SetTaskLogId(v string) *DescribeTaskLogListRequest {
	s.TaskLogId = &v
	return s
}

func (s *DescribeTaskLogListRequest) SetRegId(v string) *DescribeTaskLogListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeTaskLogListRequest) Validate() error {
	return dara.Validate(s)
}
