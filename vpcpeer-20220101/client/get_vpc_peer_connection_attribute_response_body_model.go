// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVpcPeerConnectionAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptingOwnerUid(v int64) *GetVpcPeerConnectionAttributeResponseBody
	GetAcceptingOwnerUid() *int64
	SetAcceptingRegionId(v string) *GetVpcPeerConnectionAttributeResponseBody
	GetAcceptingRegionId() *string
	SetAcceptingVpc(v *GetVpcPeerConnectionAttributeResponseBodyAcceptingVpc) *GetVpcPeerConnectionAttributeResponseBody
	GetAcceptingVpc() *GetVpcPeerConnectionAttributeResponseBodyAcceptingVpc
	SetBandwidth(v int32) *GetVpcPeerConnectionAttributeResponseBody
	GetBandwidth() *int32
	SetBizStatus(v string) *GetVpcPeerConnectionAttributeResponseBody
	GetBizStatus() *string
	SetDescription(v string) *GetVpcPeerConnectionAttributeResponseBody
	GetDescription() *string
	SetGmtCreate(v string) *GetVpcPeerConnectionAttributeResponseBody
	GetGmtCreate() *string
	SetGmtExpired(v string) *GetVpcPeerConnectionAttributeResponseBody
	GetGmtExpired() *string
	SetGmtModified(v string) *GetVpcPeerConnectionAttributeResponseBody
	GetGmtModified() *string
	SetInstanceId(v string) *GetVpcPeerConnectionAttributeResponseBody
	GetInstanceId() *string
	SetLinkType(v string) *GetVpcPeerConnectionAttributeResponseBody
	GetLinkType() *string
	SetManagedService(v string) *GetVpcPeerConnectionAttributeResponseBody
	GetManagedService() *string
	SetName(v string) *GetVpcPeerConnectionAttributeResponseBody
	GetName() *string
	SetOwnerId(v int64) *GetVpcPeerConnectionAttributeResponseBody
	GetOwnerId() *int64
	SetRegionId(v string) *GetVpcPeerConnectionAttributeResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetVpcPeerConnectionAttributeResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *GetVpcPeerConnectionAttributeResponseBody
	GetResourceGroupId() *string
	SetStatus(v string) *GetVpcPeerConnectionAttributeResponseBody
	GetStatus() *string
	SetTags(v []*GetVpcPeerConnectionAttributeResponseBodyTags) *GetVpcPeerConnectionAttributeResponseBody
	GetTags() []*GetVpcPeerConnectionAttributeResponseBodyTags
	SetVpc(v *GetVpcPeerConnectionAttributeResponseBodyVpc) *GetVpcPeerConnectionAttributeResponseBody
	GetVpc() *GetVpcPeerConnectionAttributeResponseBodyVpc
}

type GetVpcPeerConnectionAttributeResponseBody struct {
	// The ID of the Alibaba Cloud account to which the accepter VPC belongs.
	//
	// example:
	//
	// 283117732402483989
	AcceptingOwnerUid *int64 `json:"AcceptingOwnerUid,omitempty" xml:"AcceptingOwnerUid,omitempty"`
	// The region ID of the accepter VPC.
	//
	// example:
	//
	// cn-hangzhou
	AcceptingRegionId *string `json:"AcceptingRegionId,omitempty" xml:"AcceptingRegionId,omitempty"`
	// The details of the accepter VPC.
	AcceptingVpc *GetVpcPeerConnectionAttributeResponseBodyAcceptingVpc `json:"AcceptingVpc,omitempty" xml:"AcceptingVpc,omitempty" type:"Struct"`
	// The bandwidth of the VPC peering connection. Unit: Mbit /s. The value is an integer greater than 0.
	//
	// >  A value of -1 indicates that the bandwidth is unlimited.
	//
	// Default value:
	//
	// 	- The default bandwidth value of an inter-region VPC peering connection is 1,024 Mbit/s.
	//
	// 	- The default bandwidth value of an intra-region VPC peering connection is -1 Mbit/s, which indicates that the bandwidth is unlimited.
	//
	// example:
	//
	// 1024
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The status of the VPC peering connection. Valid values:
	//
	// 	- **Normal**
	//
	// 	- **FinancialLocked**
	//
	// example:
	//
	// Normal
	BizStatus *string `json:"BizStatus,omitempty" xml:"BizStatus,omitempty"`
	// The description of the VPC peering connection.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the VPC peering connection was created. The time follows the ISO 8601 standard in the `YYYY-MM-DDThh:mm:ss` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-04-24T09:02:36Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The expiration time of the VPC peering connection.
	//
	// A valid expiration time is returned only when the **Status*	- of the VPC peering connection is **Accepting*	- or **Expired**. Otherwise, **null*	- is returned.
	//
	// example:
	//
	// 2022-05-01T09:02:36Z
	GmtExpired *string `json:"GmtExpired,omitempty" xml:"GmtExpired,omitempty"`
	// The time when the VPC peering connection was modified. The time follows the ISO 8601 standard in the `YYYY-MM-DDThh:mm:ss` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-04-24T19:20:45Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the VPC peering connection.
	//
	// example:
	//
	// pcc-lnk0m24khwvtkm****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The link type of the VPC peering connection.
	//
	// example:
	//
	// Gold
	LinkType *string `json:"LinkType,omitempty" xml:"LinkType,omitempty"`
	// example:
	//
	// SWAS
	ManagedService *string `json:"ManagedService,omitempty" xml:"ManagedService,omitempty"`
	// The name of the VPC peering connection.
	//
	// example:
	//
	// vpcpeer
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the Alibaba Cloud account to which the requester VPC belongs.
	//
	// example:
	//
	// 253460731706911258
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the requester VPC.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3AC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmz7hy5z267ni
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the VPC peering connection. Valid values:
	//
	// 	- **Creating**
	//
	// 	- **Accepting**
	//
	// 	- **Updating**
	//
	// 	- **Rejected**
	//
	// 	- **Expired**
	//
	// 	- **Activated**
	//
	// 	- **Deleting**
	//
	// 	- **Deleted**
	//
	// For more information about the status of VPC peering connections, see [Overview](https://help.aliyun.com/document_detail/418507.html).
	//
	// example:
	//
	// Activated
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag list.
	Tags []*GetVpcPeerConnectionAttributeResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The details of the requester VPC.
	Vpc *GetVpcPeerConnectionAttributeResponseBodyVpc `json:"Vpc,omitempty" xml:"Vpc,omitempty" type:"Struct"`
}

func (s GetVpcPeerConnectionAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVpcPeerConnectionAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetVpcPeerConnectionAttributeResponseBody) GetAcceptingOwnerUid() *int64 {
	return s.AcceptingOwnerUid
}

func (s *GetVpcPeerConnectionAttributeResponseBody) GetAcceptingRegionId() *string {
	return s.AcceptingRegionId
}

func (s *GetVpcPeerConnectionAttributeResponseBody) GetAcceptingVpc() *GetVpcPeerConnectionAttributeResponseBodyAcceptingVpc {
	return s.AcceptingVpc
}

func (s *GetVpcPeerConnectionAttributeResponseBody) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *GetVpcPeerConnectionAttributeResponseBody) GetBizStatus() *string {
	return s.BizStatus
}

func (s *GetVpcPeerConnectionAttributeResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetVpcPeerConnectionAttributeResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetVpcPeerConnectionAttributeResponseBody) GetGmtExpired() *string {
	return s.GmtExpired
}

func (s *GetVpcPeerConnectionAttributeResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetVpcPeerConnectionAttributeResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetVpcPeerConnectionAttributeResponseBody) GetLinkType() *string {
	return s.LinkType
}

func (s *GetVpcPeerConnectionAttributeResponseBody) GetManagedService() *string {
	return s.ManagedService
}

func (s *GetVpcPeerConnectionAttributeResponseBody) GetName() *string {
	return s.Name
}

func (s *GetVpcPeerConnectionAttributeResponseBody) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetVpcPeerConnectionAttributeResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetVpcPeerConnectionAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVpcPeerConnectionAttributeResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetVpcPeerConnectionAttributeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetVpcPeerConnectionAttributeResponseBody) GetTags() []*GetVpcPeerConnectionAttributeResponseBodyTags {
	return s.Tags
}

func (s *GetVpcPeerConnectionAttributeResponseBody) GetVpc() *GetVpcPeerConnectionAttributeResponseBodyVpc {
	return s.Vpc
}

func (s *GetVpcPeerConnectionAttributeResponseBody) SetAcceptingOwnerUid(v int64) *GetVpcPeerConnectionAttributeResponseBody {
	s.AcceptingOwnerUid = &v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBody) SetAcceptingRegionId(v string) *GetVpcPeerConnectionAttributeResponseBody {
	s.AcceptingRegionId = &v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBody) SetAcceptingVpc(v *GetVpcPeerConnectionAttributeResponseBodyAcceptingVpc) *GetVpcPeerConnectionAttributeResponseBody {
	s.AcceptingVpc = v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBody) SetBandwidth(v int32) *GetVpcPeerConnectionAttributeResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBody) SetBizStatus(v string) *GetVpcPeerConnectionAttributeResponseBody {
	s.BizStatus = &v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBody) SetDescription(v string) *GetVpcPeerConnectionAttributeResponseBody {
	s.Description = &v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBody) SetGmtCreate(v string) *GetVpcPeerConnectionAttributeResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBody) SetGmtExpired(v string) *GetVpcPeerConnectionAttributeResponseBody {
	s.GmtExpired = &v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBody) SetGmtModified(v string) *GetVpcPeerConnectionAttributeResponseBody {
	s.GmtModified = &v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBody) SetInstanceId(v string) *GetVpcPeerConnectionAttributeResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBody) SetLinkType(v string) *GetVpcPeerConnectionAttributeResponseBody {
	s.LinkType = &v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBody) SetManagedService(v string) *GetVpcPeerConnectionAttributeResponseBody {
	s.ManagedService = &v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBody) SetName(v string) *GetVpcPeerConnectionAttributeResponseBody {
	s.Name = &v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBody) SetOwnerId(v int64) *GetVpcPeerConnectionAttributeResponseBody {
	s.OwnerId = &v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBody) SetRegionId(v string) *GetVpcPeerConnectionAttributeResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBody) SetRequestId(v string) *GetVpcPeerConnectionAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBody) SetResourceGroupId(v string) *GetVpcPeerConnectionAttributeResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBody) SetStatus(v string) *GetVpcPeerConnectionAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBody) SetTags(v []*GetVpcPeerConnectionAttributeResponseBodyTags) *GetVpcPeerConnectionAttributeResponseBody {
	s.Tags = v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBody) SetVpc(v *GetVpcPeerConnectionAttributeResponseBodyVpc) *GetVpcPeerConnectionAttributeResponseBody {
	s.Vpc = v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetVpcPeerConnectionAttributeResponseBodyAcceptingVpc struct {
	// The CIDR block of the accepter VPC.
	Ipv4Cidrs []*string `json:"Ipv4Cidrs,omitempty" xml:"Ipv4Cidrs,omitempty" type:"Repeated"`
	// The IPv6 CIDR block of the accepter VPC.
	Ipv6Cidrs []*string `json:"Ipv6Cidrs,omitempty" xml:"Ipv6Cidrs,omitempty" type:"Repeated"`
	// The ID of the accepter VPC.
	//
	// example:
	//
	// vpc-bp1vzjkp2q1xgnind****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetVpcPeerConnectionAttributeResponseBodyAcceptingVpc) String() string {
	return dara.Prettify(s)
}

func (s GetVpcPeerConnectionAttributeResponseBodyAcceptingVpc) GoString() string {
	return s.String()
}

func (s *GetVpcPeerConnectionAttributeResponseBodyAcceptingVpc) GetIpv4Cidrs() []*string {
	return s.Ipv4Cidrs
}

func (s *GetVpcPeerConnectionAttributeResponseBodyAcceptingVpc) GetIpv6Cidrs() []*string {
	return s.Ipv6Cidrs
}

func (s *GetVpcPeerConnectionAttributeResponseBodyAcceptingVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *GetVpcPeerConnectionAttributeResponseBodyAcceptingVpc) SetIpv4Cidrs(v []*string) *GetVpcPeerConnectionAttributeResponseBodyAcceptingVpc {
	s.Ipv4Cidrs = v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBodyAcceptingVpc) SetIpv6Cidrs(v []*string) *GetVpcPeerConnectionAttributeResponseBodyAcceptingVpc {
	s.Ipv6Cidrs = v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBodyAcceptingVpc) SetVpcId(v string) *GetVpcPeerConnectionAttributeResponseBodyAcceptingVpc {
	s.VpcId = &v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBodyAcceptingVpc) Validate() error {
	return dara.Validate(s)
}

type GetVpcPeerConnectionAttributeResponseBodyTags struct {
	// The tag key.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetVpcPeerConnectionAttributeResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s GetVpcPeerConnectionAttributeResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetVpcPeerConnectionAttributeResponseBodyTags) GetKey() *string {
	return s.Key
}

func (s *GetVpcPeerConnectionAttributeResponseBodyTags) GetValue() *string {
	return s.Value
}

func (s *GetVpcPeerConnectionAttributeResponseBodyTags) SetKey(v string) *GetVpcPeerConnectionAttributeResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBodyTags) SetValue(v string) *GetVpcPeerConnectionAttributeResponseBodyTags {
	s.Value = &v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBodyTags) Validate() error {
	return dara.Validate(s)
}

type GetVpcPeerConnectionAttributeResponseBodyVpc struct {
	// The CIDR block of the requester VPC.
	Ipv4Cidrs []*string `json:"Ipv4Cidrs,omitempty" xml:"Ipv4Cidrs,omitempty" type:"Repeated"`
	// The IPv6 CIDR block of the requester VPC.
	Ipv6Cidrs []*string `json:"Ipv6Cidrs,omitempty" xml:"Ipv6Cidrs,omitempty" type:"Repeated"`
	// The ID of the requester VPC.
	//
	// example:
	//
	// vpc-bp1gsk7h12ew7oegk****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetVpcPeerConnectionAttributeResponseBodyVpc) String() string {
	return dara.Prettify(s)
}

func (s GetVpcPeerConnectionAttributeResponseBodyVpc) GoString() string {
	return s.String()
}

func (s *GetVpcPeerConnectionAttributeResponseBodyVpc) GetIpv4Cidrs() []*string {
	return s.Ipv4Cidrs
}

func (s *GetVpcPeerConnectionAttributeResponseBodyVpc) GetIpv6Cidrs() []*string {
	return s.Ipv6Cidrs
}

func (s *GetVpcPeerConnectionAttributeResponseBodyVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *GetVpcPeerConnectionAttributeResponseBodyVpc) SetIpv4Cidrs(v []*string) *GetVpcPeerConnectionAttributeResponseBodyVpc {
	s.Ipv4Cidrs = v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBodyVpc) SetIpv6Cidrs(v []*string) *GetVpcPeerConnectionAttributeResponseBodyVpc {
	s.Ipv6Cidrs = v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBodyVpc) SetVpcId(v string) *GetVpcPeerConnectionAttributeResponseBodyVpc {
	s.VpcId = &v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBodyVpc) Validate() error {
	return dara.Validate(s)
}
