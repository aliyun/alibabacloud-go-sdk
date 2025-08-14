// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCenChildInstanceRouteEntryToCenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *CreateCenChildInstanceRouteEntryToCenRequest
	GetCenId() *string
	SetChildInstanceAliUid(v int64) *CreateCenChildInstanceRouteEntryToCenRequest
	GetChildInstanceAliUid() *int64
	SetChildInstanceId(v string) *CreateCenChildInstanceRouteEntryToCenRequest
	GetChildInstanceId() *string
	SetChildInstanceRegionId(v string) *CreateCenChildInstanceRouteEntryToCenRequest
	GetChildInstanceRegionId() *string
	SetChildInstanceType(v string) *CreateCenChildInstanceRouteEntryToCenRequest
	GetChildInstanceType() *string
	SetDestinationCidrBlock(v string) *CreateCenChildInstanceRouteEntryToCenRequest
	GetDestinationCidrBlock() *string
	SetOwnerAccount(v string) *CreateCenChildInstanceRouteEntryToCenRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateCenChildInstanceRouteEntryToCenRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateCenChildInstanceRouteEntryToCenRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateCenChildInstanceRouteEntryToCenRequest
	GetResourceOwnerId() *int64
	SetRouteTableId(v string) *CreateCenChildInstanceRouteEntryToCenRequest
	GetRouteTableId() *string
}

type CreateCenChildInstanceRouteEntryToCenRequest struct {
	// The ID of the CEN instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-7febra5nqj7jjh****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The ID of the Alibaba Cloud account to which the network instance belongs.
	//
	// > If the network instance belongs to another Alibaba Cloud account, this parameter is required.
	//
	// example:
	//
	// 1787100000000000
	ChildInstanceAliUid *int64 `json:"ChildInstanceAliUid,omitempty" xml:"ChildInstanceAliUid,omitempty"`
	// The ID of the network instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-k1alm2jbuwibhxtx2****
	ChildInstanceId *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	// The ID of the region where the network instance is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// ap-southeast-5
	ChildInstanceRegionId *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	// The type of the network instance. Valid values:
	//
	// 	- **VPC**: a virtual private cloud (VPC)
	//
	// 	- **VBR**: a virtual border router (VBR)
	//
	// This parameter is required.
	//
	// example:
	//
	// VPC
	ChildInstanceType *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	// The destination CIDR block of the route.
	//
	// Specify the value of this parameter in CIDR notation. Example: 192.168.10.0/24.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.22.0/24
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the route table configured on the network instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vtb-k1aa8ulqs39f86op8****
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s CreateCenChildInstanceRouteEntryToCenRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCenChildInstanceRouteEntryToCenRequest) GoString() string {
	return s.String()
}

func (s *CreateCenChildInstanceRouteEntryToCenRequest) GetCenId() *string {
	return s.CenId
}

func (s *CreateCenChildInstanceRouteEntryToCenRequest) GetChildInstanceAliUid() *int64 {
	return s.ChildInstanceAliUid
}

func (s *CreateCenChildInstanceRouteEntryToCenRequest) GetChildInstanceId() *string {
	return s.ChildInstanceId
}

func (s *CreateCenChildInstanceRouteEntryToCenRequest) GetChildInstanceRegionId() *string {
	return s.ChildInstanceRegionId
}

func (s *CreateCenChildInstanceRouteEntryToCenRequest) GetChildInstanceType() *string {
	return s.ChildInstanceType
}

func (s *CreateCenChildInstanceRouteEntryToCenRequest) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *CreateCenChildInstanceRouteEntryToCenRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateCenChildInstanceRouteEntryToCenRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateCenChildInstanceRouteEntryToCenRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateCenChildInstanceRouteEntryToCenRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateCenChildInstanceRouteEntryToCenRequest) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *CreateCenChildInstanceRouteEntryToCenRequest) SetCenId(v string) *CreateCenChildInstanceRouteEntryToCenRequest {
	s.CenId = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToCenRequest) SetChildInstanceAliUid(v int64) *CreateCenChildInstanceRouteEntryToCenRequest {
	s.ChildInstanceAliUid = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToCenRequest) SetChildInstanceId(v string) *CreateCenChildInstanceRouteEntryToCenRequest {
	s.ChildInstanceId = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToCenRequest) SetChildInstanceRegionId(v string) *CreateCenChildInstanceRouteEntryToCenRequest {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToCenRequest) SetChildInstanceType(v string) *CreateCenChildInstanceRouteEntryToCenRequest {
	s.ChildInstanceType = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToCenRequest) SetDestinationCidrBlock(v string) *CreateCenChildInstanceRouteEntryToCenRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToCenRequest) SetOwnerAccount(v string) *CreateCenChildInstanceRouteEntryToCenRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToCenRequest) SetOwnerId(v int64) *CreateCenChildInstanceRouteEntryToCenRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToCenRequest) SetResourceOwnerAccount(v string) *CreateCenChildInstanceRouteEntryToCenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToCenRequest) SetResourceOwnerId(v int64) *CreateCenChildInstanceRouteEntryToCenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToCenRequest) SetRouteTableId(v string) *CreateCenChildInstanceRouteEntryToCenRequest {
	s.RouteTableId = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToCenRequest) Validate() error {
	return dara.Validate(s)
}
