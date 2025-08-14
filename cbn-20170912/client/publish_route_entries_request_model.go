// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishRouteEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *PublishRouteEntriesRequest
	GetCenId() *string
	SetChildInstanceId(v string) *PublishRouteEntriesRequest
	GetChildInstanceId() *string
	SetChildInstanceRegionId(v string) *PublishRouteEntriesRequest
	GetChildInstanceRegionId() *string
	SetChildInstanceRouteTableId(v string) *PublishRouteEntriesRequest
	GetChildInstanceRouteTableId() *string
	SetChildInstanceType(v string) *PublishRouteEntriesRequest
	GetChildInstanceType() *string
	SetDestinationCidrBlock(v string) *PublishRouteEntriesRequest
	GetDestinationCidrBlock() *string
	SetResourceOwnerAccount(v string) *PublishRouteEntriesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *PublishRouteEntriesRequest
	GetResourceOwnerId() *int64
}

type PublishRouteEntriesRequest struct {
	// The ID of the CEN instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-7qthudw0ll6jmc****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The ID of the network instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-rj9gt5nll27onu****
	ChildInstanceId *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	// The ID of the region where the network instance is deployed.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// us-west-1
	ChildInstanceRegionId *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	// The ID of the route table configured on the network instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vtb-bp174d1gje7****
	ChildInstanceRouteTableId *string `json:"ChildInstanceRouteTableId,omitempty" xml:"ChildInstanceRouteTableId,omitempty"`
	// The type of the network instance. Valid values:
	//
	// 	- **VPC**: VPC
	//
	// 	- **VBR**: VBR
	//
	// This parameter is required.
	//
	// example:
	//
	// VPC
	ChildInstanceType *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	// The destination CIDR block of the route that you want to advertise.
	//
	// This parameter is required.
	//
	// example:
	//
	// 172.16.1.0/24
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s PublishRouteEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishRouteEntriesRequest) GoString() string {
	return s.String()
}

func (s *PublishRouteEntriesRequest) GetCenId() *string {
	return s.CenId
}

func (s *PublishRouteEntriesRequest) GetChildInstanceId() *string {
	return s.ChildInstanceId
}

func (s *PublishRouteEntriesRequest) GetChildInstanceRegionId() *string {
	return s.ChildInstanceRegionId
}

func (s *PublishRouteEntriesRequest) GetChildInstanceRouteTableId() *string {
	return s.ChildInstanceRouteTableId
}

func (s *PublishRouteEntriesRequest) GetChildInstanceType() *string {
	return s.ChildInstanceType
}

func (s *PublishRouteEntriesRequest) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *PublishRouteEntriesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *PublishRouteEntriesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *PublishRouteEntriesRequest) SetCenId(v string) *PublishRouteEntriesRequest {
	s.CenId = &v
	return s
}

func (s *PublishRouteEntriesRequest) SetChildInstanceId(v string) *PublishRouteEntriesRequest {
	s.ChildInstanceId = &v
	return s
}

func (s *PublishRouteEntriesRequest) SetChildInstanceRegionId(v string) *PublishRouteEntriesRequest {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *PublishRouteEntriesRequest) SetChildInstanceRouteTableId(v string) *PublishRouteEntriesRequest {
	s.ChildInstanceRouteTableId = &v
	return s
}

func (s *PublishRouteEntriesRequest) SetChildInstanceType(v string) *PublishRouteEntriesRequest {
	s.ChildInstanceType = &v
	return s
}

func (s *PublishRouteEntriesRequest) SetDestinationCidrBlock(v string) *PublishRouteEntriesRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *PublishRouteEntriesRequest) SetResourceOwnerAccount(v string) *PublishRouteEntriesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PublishRouteEntriesRequest) SetResourceOwnerId(v int64) *PublishRouteEntriesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *PublishRouteEntriesRequest) Validate() error {
	return dara.Validate(s)
}
