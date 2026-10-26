// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRejectOperationTaskApprovalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApproveId(v string) *RejectOperationTaskApprovalRequest
	GetApproveId() *string
	SetInstanceId(v string) *RejectOperationTaskApprovalRequest
	GetInstanceId() *string
	SetProjectId(v int64) *RejectOperationTaskApprovalRequest
	GetProjectId() *int64
	SetRegionId(v string) *RejectOperationTaskApprovalRequest
	GetRegionId() *string
}

type RejectOperationTaskApprovalRequest struct {
	// The O&M task approval ID.
	//
	// > You can call the ListTodoOpsTaskApprovals operation to query the value of this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ApproveId *string `json:"ApproveId,omitempty" xml:"ApproveId,omitempty"`
	// The instance ID of the Bastionhost instance.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the value of this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 1
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the Bastionhost instance.
	//
	// > For the mappings between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RejectOperationTaskApprovalRequest) String() string {
	return dara.Prettify(s)
}

func (s RejectOperationTaskApprovalRequest) GoString() string {
	return s.String()
}

func (s *RejectOperationTaskApprovalRequest) GetApproveId() *string {
	return s.ApproveId
}

func (s *RejectOperationTaskApprovalRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RejectOperationTaskApprovalRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *RejectOperationTaskApprovalRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RejectOperationTaskApprovalRequest) SetApproveId(v string) *RejectOperationTaskApprovalRequest {
	s.ApproveId = &v
	return s
}

func (s *RejectOperationTaskApprovalRequest) SetInstanceId(v string) *RejectOperationTaskApprovalRequest {
	s.InstanceId = &v
	return s
}

func (s *RejectOperationTaskApprovalRequest) SetProjectId(v int64) *RejectOperationTaskApprovalRequest {
	s.ProjectId = &v
	return s
}

func (s *RejectOperationTaskApprovalRequest) SetRegionId(v string) *RejectOperationTaskApprovalRequest {
	s.RegionId = &v
	return s
}

func (s *RejectOperationTaskApprovalRequest) Validate() error {
	return dara.Validate(s)
}
