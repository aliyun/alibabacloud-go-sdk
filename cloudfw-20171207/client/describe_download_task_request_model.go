// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDownloadTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v string) *DescribeDownloadTaskRequest
	GetCurrentPage() *string
	SetLang(v string) *DescribeDownloadTaskRequest
	GetLang() *string
	SetPageSize(v string) *DescribeDownloadTaskRequest
	GetPageSize() *string
	SetTaskType(v string) *DescribeDownloadTaskRequest
	GetTaskType() *string
}

type DescribeDownloadTaskRequest struct {
	// The page number for a paged query.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The language of the response message. Valid values:
	//
	// - **zh*	- (default): Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of entries to return on each page for a paged query. The default value is 10. The maximum value is 50.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the task. This is an enumeration. For a list of valid values, see the API for querying file download task types. If you do not set this parameter, tasks for all file types are queried.
	//
	// example:
	//
	// InternetFirewallAsset
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeDownloadTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDownloadTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeDownloadTaskRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeDownloadTaskRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDownloadTaskRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeDownloadTaskRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeDownloadTaskRequest) SetCurrentPage(v string) *DescribeDownloadTaskRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDownloadTaskRequest) SetLang(v string) *DescribeDownloadTaskRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDownloadTaskRequest) SetPageSize(v string) *DescribeDownloadTaskRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDownloadTaskRequest) SetTaskType(v string) *DescribeDownloadTaskRequest {
	s.TaskType = &v
	return s
}

func (s *DescribeDownloadTaskRequest) Validate() error {
	return dara.Validate(s)
}
