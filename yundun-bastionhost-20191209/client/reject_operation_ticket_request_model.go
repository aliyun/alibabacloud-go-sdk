// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRejectOperationTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *RejectOperationTicketRequest
	GetComment() *string
	SetInstanceId(v string) *RejectOperationTicketRequest
	GetInstanceId() *string
	SetOperationTicketId(v string) *RejectOperationTicketRequest
	GetOperationTicketId() *string
	SetRegionId(v string) *RejectOperationTicketRequest
	GetRegionId() *string
}

type RejectOperationTicketRequest struct {
	// The review remarks.
	//
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
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
	// The ID of the O\\&M application that you want to reject.
	//
	// >  You can call the [ListOperationTickets](https://help.aliyun.com/document_detail/2584313.html) operation to query the IDs of all O\\&M applications that require review.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	OperationTicketId *string `json:"OperationTicketId,omitempty" xml:"OperationTicketId,omitempty"`
	// The region ID of the bastion host.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RejectOperationTicketRequest) String() string {
	return dara.Prettify(s)
}

func (s RejectOperationTicketRequest) GoString() string {
	return s.String()
}

func (s *RejectOperationTicketRequest) GetComment() *string {
	return s.Comment
}

func (s *RejectOperationTicketRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RejectOperationTicketRequest) GetOperationTicketId() *string {
	return s.OperationTicketId
}

func (s *RejectOperationTicketRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RejectOperationTicketRequest) SetComment(v string) *RejectOperationTicketRequest {
	s.Comment = &v
	return s
}

func (s *RejectOperationTicketRequest) SetInstanceId(v string) *RejectOperationTicketRequest {
	s.InstanceId = &v
	return s
}

func (s *RejectOperationTicketRequest) SetOperationTicketId(v string) *RejectOperationTicketRequest {
	s.OperationTicketId = &v
	return s
}

func (s *RejectOperationTicketRequest) SetRegionId(v string) *RejectOperationTicketRequest {
	s.RegionId = &v
	return s
}

func (s *RejectOperationTicketRequest) Validate() error {
	return dara.Validate(s)
}
