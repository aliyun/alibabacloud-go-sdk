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
	SetRegionId(v string) *RejectOperationTaskApprovalRequest
	GetRegionId() *string
}

type RejectOperationTaskApprovalRequest struct {
	// The O&M task approval ID.
	//
	// > You can call the ListTodoOpsTaskApprovals operation to query this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ApproveId *string `json:"ApproveId,omitempty" xml:"ApproveId,omitempty"`
	// The instance ID of the bastion host.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host.
	//
	// > For information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
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

func (s *RejectOperationTaskApprovalRequest) SetRegionId(v string) *RejectOperationTaskApprovalRequest {
	s.RegionId = &v
	return s
}

func (s *RejectOperationTaskApprovalRequest) Validate() error {
	return dara.Validate(s)
}
