// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebLockProcessListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeWebLockProcessListRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeWebLockProcessListRequest
	GetPageSize() *int32
	SetProcessName(v string) *DescribeWebLockProcessListRequest
	GetProcessName() *string
	SetStatus(v int32) *DescribeWebLockProcessListRequest
	GetStatus() *int32
}

type DescribeWebLockProcessListRequest struct {
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries to return on each page. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// > We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the process.
	//
	// example:
	//
	// cron
	ProcessName *string `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	// Specifies whether the process is added to the process whitelist. Valid values:
	//
	// 	- **1**: The process is added to the process whitelist.
	//
	// 	- **0**: The process is not added to the process whitelist.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeWebLockProcessListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebLockProcessListRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebLockProcessListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeWebLockProcessListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeWebLockProcessListRequest) GetProcessName() *string {
	return s.ProcessName
}

func (s *DescribeWebLockProcessListRequest) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeWebLockProcessListRequest) SetCurrentPage(v int32) *DescribeWebLockProcessListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeWebLockProcessListRequest) SetPageSize(v int32) *DescribeWebLockProcessListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeWebLockProcessListRequest) SetProcessName(v string) *DescribeWebLockProcessListRequest {
	s.ProcessName = &v
	return s
}

func (s *DescribeWebLockProcessListRequest) SetStatus(v int32) *DescribeWebLockProcessListRequest {
	s.Status = &v
	return s
}

func (s *DescribeWebLockProcessListRequest) Validate() error {
	return dara.Validate(s)
}
