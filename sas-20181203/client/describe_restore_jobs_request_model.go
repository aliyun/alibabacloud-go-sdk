// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRestoreJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeRestoreJobsRequest
	GetCurrentPage() *int32
	SetMachineRemark(v string) *DescribeRestoreJobsRequest
	GetMachineRemark() *string
	SetPageSize(v int32) *DescribeRestoreJobsRequest
	GetPageSize() *int32
	SetStatus(v string) *DescribeRestoreJobsRequest
	GetStatus() *string
}

type DescribeRestoreJobsRequest struct {
	// The number of the page to return. Default value: **1**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The unique identifier of the server on which the restoration task is run. For example, you can use the IP address or the name of the server.
	//
	// example:
	//
	// 1.1.XX.XX
	MachineRemark *string `json:"MachineRemark,omitempty" xml:"MachineRemark,omitempty"`
	// The number of entries to return on each page. Default value: **10**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The status of the restoration task. Valid values:
	//
	// 	- **RUNNING**: The task is running.
	//
	// 	- **COMPLETE**: The task is complete.
	//
	// 	- **FAILED**: The task fails.
	//
	// 	- **CANCELING**: The task is being canceled.
	//
	// 	- **CANCELED**: The task is canceled.
	//
	// 	- **PARTIAL_COMPLETE**: The task is partially successful.
	//
	// 	- **CREATED**: The task is created but is not run.
	//
	// 	- **EXPIRED**: The task is not updated.
	//
	// 	- **QUEUED**: The task is waiting to be run.
	//
	// 	- **CLIENT_DELETED**: The task fails because the anti-ransomware agent is uninstalled.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeRestoreJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreJobsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRestoreJobsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeRestoreJobsRequest) GetMachineRemark() *string {
	return s.MachineRemark
}

func (s *DescribeRestoreJobsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRestoreJobsRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeRestoreJobsRequest) SetCurrentPage(v int32) *DescribeRestoreJobsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeRestoreJobsRequest) SetMachineRemark(v string) *DescribeRestoreJobsRequest {
	s.MachineRemark = &v
	return s
}

func (s *DescribeRestoreJobsRequest) SetPageSize(v int32) *DescribeRestoreJobsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRestoreJobsRequest) SetStatus(v string) *DescribeRestoreJobsRequest {
	s.Status = &v
	return s
}

func (s *DescribeRestoreJobsRequest) Validate() error {
	return dara.Validate(s)
}
