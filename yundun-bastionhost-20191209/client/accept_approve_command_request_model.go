// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAcceptApproveCommandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommandId(v string) *AcceptApproveCommandRequest
	GetCommandId() *string
	SetInstanceId(v string) *AcceptApproveCommandRequest
	GetInstanceId() *string
	SetProjectId(v int64) *AcceptApproveCommandRequest
	GetProjectId() *int64
	SetRegionId(v string) *AcceptApproveCommandRequest
	GetRegionId() *string
}

type AcceptApproveCommandRequest struct {
	// The ID of the command to be approved for execution.
	//
	// > You can call the [ListApproveCommands](https://help.aliyun.com/document_detail/2584310.html) operation to query the IDs of all commands that require approval.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// The instance ID of the bastion host.
	//
	// > You can call the DescribeInstances operation to query the bastion host instance ID.
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
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the bastion host.
	//
	// > For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AcceptApproveCommandRequest) String() string {
	return dara.Prettify(s)
}

func (s AcceptApproveCommandRequest) GoString() string {
	return s.String()
}

func (s *AcceptApproveCommandRequest) GetCommandId() *string {
	return s.CommandId
}

func (s *AcceptApproveCommandRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AcceptApproveCommandRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *AcceptApproveCommandRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AcceptApproveCommandRequest) SetCommandId(v string) *AcceptApproveCommandRequest {
	s.CommandId = &v
	return s
}

func (s *AcceptApproveCommandRequest) SetInstanceId(v string) *AcceptApproveCommandRequest {
	s.InstanceId = &v
	return s
}

func (s *AcceptApproveCommandRequest) SetProjectId(v int64) *AcceptApproveCommandRequest {
	s.ProjectId = &v
	return s
}

func (s *AcceptApproveCommandRequest) SetRegionId(v string) *AcceptApproveCommandRequest {
	s.RegionId = &v
	return s
}

func (s *AcceptApproveCommandRequest) Validate() error {
	return dara.Validate(s)
}
