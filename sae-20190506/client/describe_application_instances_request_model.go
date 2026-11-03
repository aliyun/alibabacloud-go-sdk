// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeApplicationInstancesRequest
	GetAppId() *string
	SetCurrentPage(v int32) *DescribeApplicationInstancesRequest
	GetCurrentPage() *int32
	SetGroupId(v string) *DescribeApplicationInstancesRequest
	GetGroupId() *string
	SetInstanceId(v string) *DescribeApplicationInstancesRequest
	GetInstanceId() *string
	SetPageSize(v int32) *DescribeApplicationInstancesRequest
	GetPageSize() *int32
	SetPipelineId(v string) *DescribeApplicationInstancesRequest
	GetPipelineId() *string
	SetReverse(v bool) *DescribeApplicationInstancesRequest
	GetReverse() *bool
}

type DescribeApplicationInstancesRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// d700e680-aa4d-4ec1-afc2-6566b5ff****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The application group ID. You can call the [DescribeApplicationGroups](https://help.aliyun.com/document_detail/126249.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// b2a8a925-477a-4ed7-b825-d5e22500****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The application instance ID.
	//
	// example:
	//
	// demo-faaca66c-5959-45cc-b3bf-d26093b2e9c0******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The paging Settings for the number of entries per page in a paged query. Default value: **10**. Valid values: (0, 1000000000).
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The batch ID. You can call the [DescribeChangeOrder](https://help.aliyun.com/document_detail/126617.html) operation to obtain the ID.
	//
	// example:
	//
	// 85750d48-6cfc-4dbf-8ea0-840397******
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// Specifies the sorting order of application instances based on their running status. If the status is the same, the instances are sorted by instance ID. Valid values:
	//
	// - **true**: Descending order.
	//
	// - **false**: Ascending order.
	//
	// The following list shows the instance statuses in ascending order:
	//
	// 1. **Error**: An error occurred during instance startup.
	//
	// 2. **CrashLoopBackOff**: The container failed to start. An error occurred during startup, and the error persisted after a restart.
	//
	// 3. **ErrImagePull**: An error occurred while pulling the container image for the instance.
	//
	// 4. **ImagePullBackOff**: The container image cannot be retrieved.
	//
	// 5. **Pending**: The instance is waiting to be scheduled.
	//
	// 6. **Unknown**: An unknown exception occurred.
	//
	// 7. **Terminating**: The instance is being terminated.
	//
	// 8. **NotFound**: The instance cannot be found.
	//
	// 9. **PodInitializing**: The instance is being initialized.
	//
	// 10. **Init:0/1**: The instance is initializing.
	//
	// 11. **Running**: The instance is running.
	//
	// example:
	//
	// true
	Reverse *bool `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
}

func (s DescribeApplicationInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationInstancesRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeApplicationInstancesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeApplicationInstancesRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeApplicationInstancesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeApplicationInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApplicationInstancesRequest) GetPipelineId() *string {
	return s.PipelineId
}

func (s *DescribeApplicationInstancesRequest) GetReverse() *bool {
	return s.Reverse
}

func (s *DescribeApplicationInstancesRequest) SetAppId(v string) *DescribeApplicationInstancesRequest {
	s.AppId = &v
	return s
}

func (s *DescribeApplicationInstancesRequest) SetCurrentPage(v int32) *DescribeApplicationInstancesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeApplicationInstancesRequest) SetGroupId(v string) *DescribeApplicationInstancesRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeApplicationInstancesRequest) SetInstanceId(v string) *DescribeApplicationInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeApplicationInstancesRequest) SetPageSize(v int32) *DescribeApplicationInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApplicationInstancesRequest) SetPipelineId(v string) *DescribeApplicationInstancesRequest {
	s.PipelineId = &v
	return s
}

func (s *DescribeApplicationInstancesRequest) SetReverse(v bool) *DescribeApplicationInstancesRequest {
	s.Reverse = &v
	return s
}

func (s *DescribeApplicationInstancesRequest) Validate() error {
	return dara.Validate(s)
}
