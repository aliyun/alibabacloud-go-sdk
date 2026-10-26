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
	SetProjectId(v int64) *RejectOperationTicketRequest
	GetProjectId() *int64
	SetRegionId(v string) *RejectOperationTicketRequest
	GetRegionId() *string
}

type RejectOperationTicketRequest struct {
	// The approval comment.
	//
	// example:
	//
	// Reject
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The instance ID of the Bastionhost instance.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain the parameter value.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the O&M application to be rejected.
	//
	// > You can call the [ListOperationTickets](https://help.aliyun.com/document_detail/2584313.html) operation to query the IDs of all O&M applications that require approval.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	OperationTicketId *string `json:"OperationTicketId,omitempty" xml:"OperationTicketId,omitempty"`
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

func (s *RejectOperationTicketRequest) GetProjectId() *int64 {
	return s.ProjectId
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

func (s *RejectOperationTicketRequest) SetProjectId(v int64) *RejectOperationTicketRequest {
	s.ProjectId = &v
	return s
}

func (s *RejectOperationTicketRequest) SetRegionId(v string) *RejectOperationTicketRequest {
	s.RegionId = &v
	return s
}

func (s *RejectOperationTicketRequest) Validate() error {
	return dara.Validate(s)
}
