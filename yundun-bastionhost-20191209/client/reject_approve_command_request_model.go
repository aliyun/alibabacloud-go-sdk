// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRejectApproveCommandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommandId(v string) *RejectApproveCommandRequest
	GetCommandId() *string
	SetInstanceId(v string) *RejectApproveCommandRequest
	GetInstanceId() *string
	SetRegionId(v string) *RejectApproveCommandRequest
	GetRegionId() *string
}

type RejectApproveCommandRequest struct {
	// The ID of the command that you want to reject.
	//
	// >  You can call the [ListApproveCommands](https://help.aliyun.com/document_detail/2584310.html) operation to query the IDs of all commands that need to be reviewed.
	//
	// This parameter is required.
	//
	// example:
	//
	// 574
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// The ID of the bastion host.
	//
	// >  You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RejectApproveCommandRequest) String() string {
	return dara.Prettify(s)
}

func (s RejectApproveCommandRequest) GoString() string {
	return s.String()
}

func (s *RejectApproveCommandRequest) GetCommandId() *string {
	return s.CommandId
}

func (s *RejectApproveCommandRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RejectApproveCommandRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RejectApproveCommandRequest) SetCommandId(v string) *RejectApproveCommandRequest {
	s.CommandId = &v
	return s
}

func (s *RejectApproveCommandRequest) SetInstanceId(v string) *RejectApproveCommandRequest {
	s.InstanceId = &v
	return s
}

func (s *RejectApproveCommandRequest) SetRegionId(v string) *RejectApproveCommandRequest {
	s.RegionId = &v
	return s
}

func (s *RejectApproveCommandRequest) Validate() error {
	return dara.Validate(s)
}
