// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcPeerConnectionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListVpcPeerConnectionsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListVpcPeerConnectionsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListVpcPeerConnectionsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListVpcPeerConnectionsResponseBody
	GetTotalCount() *int32
	SetVpcPeerConnects(v []*ListVpcPeerConnectionsResponseBodyVpcPeerConnects) *ListVpcPeerConnectionsResponseBody
	GetVpcPeerConnects() []*ListVpcPeerConnectionsResponseBodyVpcPeerConnects
}

type ListVpcPeerConnectionsResponseBody struct {
	// The number of entries per page. Valid values: **1*	- to **100**. Default value: **20**.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used for the next query. Valid values:
	//
	// 	- If no value is returned for **NextToken**, no next queries are sent.
	//
	// 	- If the value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED39DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The details of the VPC peering connections.
	VpcPeerConnects []*ListVpcPeerConnectionsResponseBodyVpcPeerConnects `json:"VpcPeerConnects,omitempty" xml:"VpcPeerConnects,omitempty" type:"Repeated"`
}

func (s ListVpcPeerConnectionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVpcPeerConnectionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpcPeerConnectionsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListVpcPeerConnectionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVpcPeerConnectionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVpcPeerConnectionsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListVpcPeerConnectionsResponseBody) GetVpcPeerConnects() []*ListVpcPeerConnectionsResponseBodyVpcPeerConnects {
	return s.VpcPeerConnects
}

func (s *ListVpcPeerConnectionsResponseBody) SetMaxResults(v int32) *ListVpcPeerConnectionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBody) SetNextToken(v string) *ListVpcPeerConnectionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBody) SetRequestId(v string) *ListVpcPeerConnectionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBody) SetTotalCount(v int32) *ListVpcPeerConnectionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBody) SetVpcPeerConnects(v []*ListVpcPeerConnectionsResponseBodyVpcPeerConnects) *ListVpcPeerConnectionsResponseBody {
	s.VpcPeerConnects = v
	return s
}

func (s *ListVpcPeerConnectionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListVpcPeerConnectionsResponseBodyVpcPeerConnects struct {
	// The ID of the Alibaba Cloud account to which the accepter VPC belongs.
	//
	// example:
	//
	// 253460731706911258
	AcceptingOwnerUid *int64 `json:"AcceptingOwnerUid,omitempty" xml:"AcceptingOwnerUid,omitempty"`
	// The region ID of the accepter VPC.
	//
	// example:
	//
	// cn-hangzhou
	AcceptingRegionId *string `json:"AcceptingRegionId,omitempty" xml:"AcceptingRegionId,omitempty"`
	// The details of the accepter VPC.
	AcceptingVpc *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsAcceptingVpc `json:"AcceptingVpc,omitempty" xml:"AcceptingVpc,omitempty" type:"Struct"`
	// The bandwidth of the VPC peering connection. Unit: Mbit/s. The value is an integer greater than 0.
	//
	// >  If the value is set to -1, it indicates that no limit is imposed on the bandwidth.
	//
	// Default value:
	//
	// 	- The default bandwidth of an inter-region VPC peering connection is **1024*	- Mbit/s.
	//
	// 	- The default bandwidth of an intra-region VPC peering connection is **-1*	- Mbit/s.
	//
	// example:
	//
	// 1024
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The business status of the VPC peering connection. Valid values:
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
	// The time when the VPC peering connection was created. The time is displayed in the `YYYY-MM-DDThh:mm:ssZ` format in UTC.
	//
	// example:
	//
	// 2022-04-24T09:02:36Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The expiration time of the VPC peering connection. The time is displayed in the `YYYY-MM-DDThh:mm:ssZ` format in UTC.
	//
	// The expiration time is returned only when the **Status*	- of the VPC peering connection is **Accepting*	- or **Expired**. Otherwise, **null*	- is returned.
	//
	// example:
	//
	// 2022-05-01T09:02:36Z
	GmtExpired *string `json:"GmtExpired,omitempty" xml:"GmtExpired,omitempty"`
	// The time when the VPC peering connection was modified. The time is displayed in the `YYYY-MM-DDThh:mm:ssZ` format in UTC.
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
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek2gvbs746gt4q
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
	// For more information about the status of VPC peering connections, see [Overview of VPC peering connections](https://help.aliyun.com/document_detail/418507.html).
	//
	// example:
	//
	// Activated
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag list.
	Tags []*ListVpcPeerConnectionsResponseBodyVpcPeerConnectsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The details of the requester VPC.
	Vpc *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsVpc `json:"Vpc,omitempty" xml:"Vpc,omitempty" type:"Struct"`
}

func (s ListVpcPeerConnectionsResponseBodyVpcPeerConnects) String() string {
	return dara.Prettify(s)
}

func (s ListVpcPeerConnectionsResponseBodyVpcPeerConnects) GoString() string {
	return s.String()
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) GetAcceptingOwnerUid() *int64 {
	return s.AcceptingOwnerUid
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) GetAcceptingRegionId() *string {
	return s.AcceptingRegionId
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) GetAcceptingVpc() *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsAcceptingVpc {
	return s.AcceptingVpc
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) GetBizStatus() *string {
	return s.BizStatus
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) GetDescription() *string {
	return s.Description
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) GetGmtExpired() *string {
	return s.GmtExpired
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) GetLinkType() *string {
	return s.LinkType
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) GetManagedService() *string {
	return s.ManagedService
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) GetName() *string {
	return s.Name
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) GetRegionId() *string {
	return s.RegionId
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) GetStatus() *string {
	return s.Status
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) GetTags() []*ListVpcPeerConnectionsResponseBodyVpcPeerConnectsTags {
	return s.Tags
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) GetVpc() *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsVpc {
	return s.Vpc
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) SetAcceptingOwnerUid(v int64) *ListVpcPeerConnectionsResponseBodyVpcPeerConnects {
	s.AcceptingOwnerUid = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) SetAcceptingRegionId(v string) *ListVpcPeerConnectionsResponseBodyVpcPeerConnects {
	s.AcceptingRegionId = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) SetAcceptingVpc(v *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsAcceptingVpc) *ListVpcPeerConnectionsResponseBodyVpcPeerConnects {
	s.AcceptingVpc = v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) SetBandwidth(v int32) *ListVpcPeerConnectionsResponseBodyVpcPeerConnects {
	s.Bandwidth = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) SetBizStatus(v string) *ListVpcPeerConnectionsResponseBodyVpcPeerConnects {
	s.BizStatus = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) SetDescription(v string) *ListVpcPeerConnectionsResponseBodyVpcPeerConnects {
	s.Description = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) SetGmtCreate(v string) *ListVpcPeerConnectionsResponseBodyVpcPeerConnects {
	s.GmtCreate = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) SetGmtExpired(v string) *ListVpcPeerConnectionsResponseBodyVpcPeerConnects {
	s.GmtExpired = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) SetGmtModified(v string) *ListVpcPeerConnectionsResponseBodyVpcPeerConnects {
	s.GmtModified = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) SetInstanceId(v string) *ListVpcPeerConnectionsResponseBodyVpcPeerConnects {
	s.InstanceId = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) SetLinkType(v string) *ListVpcPeerConnectionsResponseBodyVpcPeerConnects {
	s.LinkType = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) SetManagedService(v string) *ListVpcPeerConnectionsResponseBodyVpcPeerConnects {
	s.ManagedService = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) SetName(v string) *ListVpcPeerConnectionsResponseBodyVpcPeerConnects {
	s.Name = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) SetOwnerId(v int64) *ListVpcPeerConnectionsResponseBodyVpcPeerConnects {
	s.OwnerId = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) SetRegionId(v string) *ListVpcPeerConnectionsResponseBodyVpcPeerConnects {
	s.RegionId = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) SetResourceGroupId(v string) *ListVpcPeerConnectionsResponseBodyVpcPeerConnects {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) SetStatus(v string) *ListVpcPeerConnectionsResponseBodyVpcPeerConnects {
	s.Status = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) SetTags(v []*ListVpcPeerConnectionsResponseBodyVpcPeerConnectsTags) *ListVpcPeerConnectionsResponseBodyVpcPeerConnects {
	s.Tags = v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) SetVpc(v *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsVpc) *ListVpcPeerConnectionsResponseBodyVpcPeerConnects {
	s.Vpc = v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) Validate() error {
	return dara.Validate(s)
}

type ListVpcPeerConnectionsResponseBodyVpcPeerConnectsAcceptingVpc struct {
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

func (s ListVpcPeerConnectionsResponseBodyVpcPeerConnectsAcceptingVpc) String() string {
	return dara.Prettify(s)
}

func (s ListVpcPeerConnectionsResponseBodyVpcPeerConnectsAcceptingVpc) GoString() string {
	return s.String()
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsAcceptingVpc) GetIpv4Cidrs() []*string {
	return s.Ipv4Cidrs
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsAcceptingVpc) GetIpv6Cidrs() []*string {
	return s.Ipv6Cidrs
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsAcceptingVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsAcceptingVpc) SetIpv4Cidrs(v []*string) *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsAcceptingVpc {
	s.Ipv4Cidrs = v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsAcceptingVpc) SetIpv6Cidrs(v []*string) *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsAcceptingVpc {
	s.Ipv6Cidrs = v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsAcceptingVpc) SetVpcId(v string) *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsAcceptingVpc {
	s.VpcId = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsAcceptingVpc) Validate() error {
	return dara.Validate(s)
}

type ListVpcPeerConnectionsResponseBodyVpcPeerConnectsTags struct {
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

func (s ListVpcPeerConnectionsResponseBodyVpcPeerConnectsTags) String() string {
	return dara.Prettify(s)
}

func (s ListVpcPeerConnectionsResponseBodyVpcPeerConnectsTags) GoString() string {
	return s.String()
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsTags) GetKey() *string {
	return s.Key
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsTags) GetValue() *string {
	return s.Value
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsTags) SetKey(v string) *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsTags {
	s.Key = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsTags) SetValue(v string) *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsTags {
	s.Value = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsTags) Validate() error {
	return dara.Validate(s)
}

type ListVpcPeerConnectionsResponseBodyVpcPeerConnectsVpc struct {
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

func (s ListVpcPeerConnectionsResponseBodyVpcPeerConnectsVpc) String() string {
	return dara.Prettify(s)
}

func (s ListVpcPeerConnectionsResponseBodyVpcPeerConnectsVpc) GoString() string {
	return s.String()
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsVpc) GetIpv4Cidrs() []*string {
	return s.Ipv4Cidrs
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsVpc) GetIpv6Cidrs() []*string {
	return s.Ipv6Cidrs
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsVpc) SetIpv4Cidrs(v []*string) *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsVpc {
	s.Ipv4Cidrs = v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsVpc) SetIpv6Cidrs(v []*string) *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsVpc {
	s.Ipv6Cidrs = v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsVpc) SetVpcId(v string) *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsVpc {
	s.VpcId = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsVpc) Validate() error {
	return dara.Validate(s)
}
