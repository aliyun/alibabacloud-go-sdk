// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWithdrawPublishedRouteEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *WithdrawPublishedRouteEntriesRequest
	GetCenId() *string
	SetChildInstanceId(v string) *WithdrawPublishedRouteEntriesRequest
	GetChildInstanceId() *string
	SetChildInstanceRegionId(v string) *WithdrawPublishedRouteEntriesRequest
	GetChildInstanceRegionId() *string
	SetChildInstanceRouteTableId(v string) *WithdrawPublishedRouteEntriesRequest
	GetChildInstanceRouteTableId() *string
	SetChildInstanceType(v string) *WithdrawPublishedRouteEntriesRequest
	GetChildInstanceType() *string
	SetDestinationCidrBlock(v string) *WithdrawPublishedRouteEntriesRequest
	GetDestinationCidrBlock() *string
	SetResourceOwnerAccount(v string) *WithdrawPublishedRouteEntriesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *WithdrawPublishedRouteEntriesRequest
	GetResourceOwnerId() *int64
}

type WithdrawPublishedRouteEntriesRequest struct {
	// The ID of the CEN instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-sxjfjkjfkjfiein****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The ID of the attached network instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-rj9gt5nll27onu7****
	ChildInstanceId *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	// The ID of the region where the attached network instance is created.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	ChildInstanceRegionId *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	// The ID of the route table of the attached network instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vtb-bp174d1gje79u1g4t****
	ChildInstanceRouteTableId *string `json:"ChildInstanceRouteTableId,omitempty" xml:"ChildInstanceRouteTableId,omitempty"`
	// The type of the attached network instance.
	//
	// Set the value to **VPC**, which indicates a virtual private cloud (VPC).
	//
	// This parameter is required.
	//
	// example:
	//
	// VPC
	ChildInstanceType *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	// The destination CIDR block of the route that you want to withdraw.
	//
	// This parameter is required.
	//
	// example:
	//
	// 172.XX.XX.0/24
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s WithdrawPublishedRouteEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s WithdrawPublishedRouteEntriesRequest) GoString() string {
	return s.String()
}

func (s *WithdrawPublishedRouteEntriesRequest) GetCenId() *string {
	return s.CenId
}

func (s *WithdrawPublishedRouteEntriesRequest) GetChildInstanceId() *string {
	return s.ChildInstanceId
}

func (s *WithdrawPublishedRouteEntriesRequest) GetChildInstanceRegionId() *string {
	return s.ChildInstanceRegionId
}

func (s *WithdrawPublishedRouteEntriesRequest) GetChildInstanceRouteTableId() *string {
	return s.ChildInstanceRouteTableId
}

func (s *WithdrawPublishedRouteEntriesRequest) GetChildInstanceType() *string {
	return s.ChildInstanceType
}

func (s *WithdrawPublishedRouteEntriesRequest) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *WithdrawPublishedRouteEntriesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *WithdrawPublishedRouteEntriesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *WithdrawPublishedRouteEntriesRequest) SetCenId(v string) *WithdrawPublishedRouteEntriesRequest {
	s.CenId = &v
	return s
}

func (s *WithdrawPublishedRouteEntriesRequest) SetChildInstanceId(v string) *WithdrawPublishedRouteEntriesRequest {
	s.ChildInstanceId = &v
	return s
}

func (s *WithdrawPublishedRouteEntriesRequest) SetChildInstanceRegionId(v string) *WithdrawPublishedRouteEntriesRequest {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *WithdrawPublishedRouteEntriesRequest) SetChildInstanceRouteTableId(v string) *WithdrawPublishedRouteEntriesRequest {
	s.ChildInstanceRouteTableId = &v
	return s
}

func (s *WithdrawPublishedRouteEntriesRequest) SetChildInstanceType(v string) *WithdrawPublishedRouteEntriesRequest {
	s.ChildInstanceType = &v
	return s
}

func (s *WithdrawPublishedRouteEntriesRequest) SetDestinationCidrBlock(v string) *WithdrawPublishedRouteEntriesRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *WithdrawPublishedRouteEntriesRequest) SetResourceOwnerAccount(v string) *WithdrawPublishedRouteEntriesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *WithdrawPublishedRouteEntriesRequest) SetResourceOwnerId(v int64) *WithdrawPublishedRouteEntriesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *WithdrawPublishedRouteEntriesRequest) Validate() error {
	return dara.Validate(s)
}
