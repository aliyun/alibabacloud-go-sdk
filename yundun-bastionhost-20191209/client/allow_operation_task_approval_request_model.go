// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllowOperationTaskApprovalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApproveId(v string) *AllowOperationTaskApprovalRequest
	GetApproveId() *string
	SetInstanceId(v string) *AllowOperationTaskApprovalRequest
	GetInstanceId() *string
	SetProjectId(v int64) *AllowOperationTaskApprovalRequest
	GetProjectId() *int64
	SetRegionId(v string) *AllowOperationTaskApprovalRequest
	GetRegionId() *string
}

type AllowOperationTaskApprovalRequest struct {
	// The approval ID of the O&M task.
	//
	// > You can call the ListTodoOpsTaskApprovals operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ApproveId *string `json:"ApproveId,omitempty" xml:"ApproveId,omitempty"`
	// The instance ID of the Bastionhost.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-lbj3bw4ma02
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the Bastionhost instance.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AllowOperationTaskApprovalRequest) String() string {
	return dara.Prettify(s)
}

func (s AllowOperationTaskApprovalRequest) GoString() string {
	return s.String()
}

func (s *AllowOperationTaskApprovalRequest) GetApproveId() *string {
	return s.ApproveId
}

func (s *AllowOperationTaskApprovalRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AllowOperationTaskApprovalRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *AllowOperationTaskApprovalRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AllowOperationTaskApprovalRequest) SetApproveId(v string) *AllowOperationTaskApprovalRequest {
	s.ApproveId = &v
	return s
}

func (s *AllowOperationTaskApprovalRequest) SetInstanceId(v string) *AllowOperationTaskApprovalRequest {
	s.InstanceId = &v
	return s
}

func (s *AllowOperationTaskApprovalRequest) SetProjectId(v int64) *AllowOperationTaskApprovalRequest {
	s.ProjectId = &v
	return s
}

func (s *AllowOperationTaskApprovalRequest) SetRegionId(v string) *AllowOperationTaskApprovalRequest {
	s.RegionId = &v
	return s
}

func (s *AllowOperationTaskApprovalRequest) Validate() error {
	return dara.Validate(s)
}
