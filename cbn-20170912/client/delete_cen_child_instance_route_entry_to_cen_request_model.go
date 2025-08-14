// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCenChildInstanceRouteEntryToCenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *DeleteCenChildInstanceRouteEntryToCenRequest
	GetCenId() *string
	SetChildInstanceAliUid(v int64) *DeleteCenChildInstanceRouteEntryToCenRequest
	GetChildInstanceAliUid() *int64
	SetChildInstanceId(v string) *DeleteCenChildInstanceRouteEntryToCenRequest
	GetChildInstanceId() *string
	SetChildInstanceRegionId(v string) *DeleteCenChildInstanceRouteEntryToCenRequest
	GetChildInstanceRegionId() *string
	SetChildInstanceType(v string) *DeleteCenChildInstanceRouteEntryToCenRequest
	GetChildInstanceType() *string
	SetDestinationCidrBlock(v string) *DeleteCenChildInstanceRouteEntryToCenRequest
	GetDestinationCidrBlock() *string
	SetOwnerAccount(v string) *DeleteCenChildInstanceRouteEntryToCenRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteCenChildInstanceRouteEntryToCenRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteCenChildInstanceRouteEntryToCenRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteCenChildInstanceRouteEntryToCenRequest
	GetResourceOwnerId() *int64
	SetRouteTableId(v string) *DeleteCenChildInstanceRouteEntryToCenRequest
	GetRouteTableId() *string
}

type DeleteCenChildInstanceRouteEntryToCenRequest struct {
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
	// > If the network instance and the CEN instance belong to different Alibaba Cloud accounts, this parameter is required.
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
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to obtain the region ID.
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

func (s DeleteCenChildInstanceRouteEntryToCenRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCenChildInstanceRouteEntryToCenRequest) GoString() string {
	return s.String()
}

func (s *DeleteCenChildInstanceRouteEntryToCenRequest) GetCenId() *string {
	return s.CenId
}

func (s *DeleteCenChildInstanceRouteEntryToCenRequest) GetChildInstanceAliUid() *int64 {
	return s.ChildInstanceAliUid
}

func (s *DeleteCenChildInstanceRouteEntryToCenRequest) GetChildInstanceId() *string {
	return s.ChildInstanceId
}

func (s *DeleteCenChildInstanceRouteEntryToCenRequest) GetChildInstanceRegionId() *string {
	return s.ChildInstanceRegionId
}

func (s *DeleteCenChildInstanceRouteEntryToCenRequest) GetChildInstanceType() *string {
	return s.ChildInstanceType
}

func (s *DeleteCenChildInstanceRouteEntryToCenRequest) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *DeleteCenChildInstanceRouteEntryToCenRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteCenChildInstanceRouteEntryToCenRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteCenChildInstanceRouteEntryToCenRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteCenChildInstanceRouteEntryToCenRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteCenChildInstanceRouteEntryToCenRequest) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *DeleteCenChildInstanceRouteEntryToCenRequest) SetCenId(v string) *DeleteCenChildInstanceRouteEntryToCenRequest {
	s.CenId = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToCenRequest) SetChildInstanceAliUid(v int64) *DeleteCenChildInstanceRouteEntryToCenRequest {
	s.ChildInstanceAliUid = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToCenRequest) SetChildInstanceId(v string) *DeleteCenChildInstanceRouteEntryToCenRequest {
	s.ChildInstanceId = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToCenRequest) SetChildInstanceRegionId(v string) *DeleteCenChildInstanceRouteEntryToCenRequest {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToCenRequest) SetChildInstanceType(v string) *DeleteCenChildInstanceRouteEntryToCenRequest {
	s.ChildInstanceType = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToCenRequest) SetDestinationCidrBlock(v string) *DeleteCenChildInstanceRouteEntryToCenRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToCenRequest) SetOwnerAccount(v string) *DeleteCenChildInstanceRouteEntryToCenRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToCenRequest) SetOwnerId(v int64) *DeleteCenChildInstanceRouteEntryToCenRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToCenRequest) SetResourceOwnerAccount(v string) *DeleteCenChildInstanceRouteEntryToCenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToCenRequest) SetResourceOwnerId(v int64) *DeleteCenChildInstanceRouteEntryToCenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToCenRequest) SetRouteTableId(v string) *DeleteCenChildInstanceRouteEntryToCenRequest {
	s.RouteTableId = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToCenRequest) Validate() error {
	return dara.Validate(s)
}
