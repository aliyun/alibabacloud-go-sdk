// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelOperationTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CancelOperationTicketRequest
	GetInstanceId() *string
	SetOperationTicketId(v string) *CancelOperationTicketRequest
	GetOperationTicketId() *string
	SetRegionId(v string) *CancelOperationTicketRequest
	GetRegionId() *string
}

type CancelOperationTicketRequest struct {
	// The instance ID of the Bastionhost instance.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-78v1gh****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the O&M request to be withdrawn.
	//
	// > You can only withdraw O&M requests that have not been fully approved.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	OperationTicketId *string `json:"OperationTicketId,omitempty" xml:"OperationTicketId,omitempty"`
	// The region ID of the Bastionhost instance.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CancelOperationTicketRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelOperationTicketRequest) GoString() string {
	return s.String()
}

func (s *CancelOperationTicketRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CancelOperationTicketRequest) GetOperationTicketId() *string {
	return s.OperationTicketId
}

func (s *CancelOperationTicketRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CancelOperationTicketRequest) SetInstanceId(v string) *CancelOperationTicketRequest {
	s.InstanceId = &v
	return s
}

func (s *CancelOperationTicketRequest) SetOperationTicketId(v string) *CancelOperationTicketRequest {
	s.OperationTicketId = &v
	return s
}

func (s *CancelOperationTicketRequest) SetRegionId(v string) *CancelOperationTicketRequest {
	s.RegionId = &v
	return s
}

func (s *CancelOperationTicketRequest) Validate() error {
	return dara.Validate(s)
}
