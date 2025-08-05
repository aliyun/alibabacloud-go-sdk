// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDownloadTaskTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v string) *DescribeDownloadTaskTypeRequest
	GetCurrentPage() *string
	SetLang(v string) *DescribeDownloadTaskTypeRequest
	GetLang() *string
	SetPageSize(v string) *DescribeDownloadTaskTypeRequest
	GetPageSize() *string
	SetTaskType(v string) *DescribeDownloadTaskTypeRequest
	GetTaskType() *string
}

type DescribeDownloadTaskTypeRequest struct {
	// The page number. Pages start from page 1. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// 	- **zh*	- (default): Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 50.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the task.
	//
	// example:
	//
	// InternetFirewallAsset
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeDownloadTaskTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDownloadTaskTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDownloadTaskTypeRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeDownloadTaskTypeRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDownloadTaskTypeRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeDownloadTaskTypeRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeDownloadTaskTypeRequest) SetCurrentPage(v string) *DescribeDownloadTaskTypeRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDownloadTaskTypeRequest) SetLang(v string) *DescribeDownloadTaskTypeRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDownloadTaskTypeRequest) SetPageSize(v string) *DescribeDownloadTaskTypeRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDownloadTaskTypeRequest) SetTaskType(v string) *DescribeDownloadTaskTypeRequest {
	s.TaskType = &v
	return s
}

func (s *DescribeDownloadTaskTypeRequest) Validate() error {
	return dara.Validate(s)
}
