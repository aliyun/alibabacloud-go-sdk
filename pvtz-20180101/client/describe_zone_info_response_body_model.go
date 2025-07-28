// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeZoneInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBindEdgeDnsClusters(v *DescribeZoneInfoResponseBodyBindEdgeDnsClusters) *DescribeZoneInfoResponseBody
	GetBindEdgeDnsClusters() *DescribeZoneInfoResponseBodyBindEdgeDnsClusters
	SetBindVpcs(v *DescribeZoneInfoResponseBodyBindVpcs) *DescribeZoneInfoResponseBody
	GetBindVpcs() *DescribeZoneInfoResponseBodyBindVpcs
	SetCreateTime(v string) *DescribeZoneInfoResponseBody
	GetCreateTime() *string
	SetCreateTimestamp(v int64) *DescribeZoneInfoResponseBody
	GetCreateTimestamp() *int64
	SetCreator(v string) *DescribeZoneInfoResponseBody
	GetCreator() *string
	SetCreatorType(v string) *DescribeZoneInfoResponseBody
	GetCreatorType() *string
	SetDnsGroup(v string) *DescribeZoneInfoResponseBody
	GetDnsGroup() *string
	SetDnsGroupChanging(v bool) *DescribeZoneInfoResponseBody
	GetDnsGroupChanging() *bool
	SetIsPtr(v bool) *DescribeZoneInfoResponseBody
	GetIsPtr() *bool
	SetProxyPattern(v string) *DescribeZoneInfoResponseBody
	GetProxyPattern() *string
	SetRecordCount(v int32) *DescribeZoneInfoResponseBody
	GetRecordCount() *int32
	SetRemark(v string) *DescribeZoneInfoResponseBody
	GetRemark() *string
	SetRequestId(v string) *DescribeZoneInfoResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *DescribeZoneInfoResponseBody
	GetResourceGroupId() *string
	SetSlaveDns(v bool) *DescribeZoneInfoResponseBody
	GetSlaveDns() *bool
	SetUpdateTime(v string) *DescribeZoneInfoResponseBody
	GetUpdateTime() *string
	SetUpdateTimestamp(v int64) *DescribeZoneInfoResponseBody
	GetUpdateTimestamp() *int64
	SetZoneId(v string) *DescribeZoneInfoResponseBody
	GetZoneId() *string
	SetZoneName(v string) *DescribeZoneInfoResponseBody
	GetZoneName() *string
	SetZoneTag(v string) *DescribeZoneInfoResponseBody
	GetZoneTag() *string
	SetZoneType(v string) *DescribeZoneInfoResponseBody
	GetZoneType() *string
}

type DescribeZoneInfoResponseBody struct {
	BindEdgeDnsClusters *DescribeZoneInfoResponseBodyBindEdgeDnsClusters `json:"BindEdgeDnsClusters,omitempty" xml:"BindEdgeDnsClusters,omitempty" type:"Struct"`
	// The VPCs associated with the zone.
	BindVpcs *DescribeZoneInfoResponseBodyBindVpcs `json:"BindVpcs,omitempty" xml:"BindVpcs,omitempty" type:"Struct"`
	// The time when the zone was created. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-01-23T03:15Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the zone was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1516775741000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The creator of the zone.
	//
	// example:
	//
	// 141339776561****
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The type of the creator.
	//
	// example:
	//
	// USER
	CreatorType *string `json:"CreatorType,omitempty" xml:"CreatorType,omitempty"`
	// The logical location type of the built-in authoritative module in which the zone is added. Valid values:
	//
	// 	- **NORMAL_ZONE**: regular module
	//
	// 	- **FAST_ZONE**: acceleration module
	//
	// example:
	//
	// FAST_ZONE
	DnsGroup *string `json:"DnsGroup,omitempty" xml:"DnsGroup,omitempty"`
	// Indicates whether the zone is being removed to another logical location. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	DnsGroupChanging *bool `json:"DnsGroupChanging,omitempty" xml:"DnsGroupChanging,omitempty"`
	// Indicates whether the zone is a reverse lookup zone. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	IsPtr *bool `json:"IsPtr,omitempty" xml:"IsPtr,omitempty"`
	// Indicates whether the recursive resolution proxy for subdomain names is enabled. Valid values:
	//
	// 	- ZONE: The recursive resolution proxy for subdomain names is disabled. In this case, NXDOMAIN is returned if the queried domain name does not exist in the zone.
	//
	// 	- RECORD: The recursive resolution proxy for subdomain names is enabled. In this case, if the queried domain name does not exist in the zone, DNS requests are recursively forwarded to the forward module and then to the recursion module until DNS results are returned.
	//
	// example:
	//
	// ZONE
	ProxyPattern *string `json:"ProxyPattern,omitempty" xml:"ProxyPattern,omitempty"`
	// The total number of DNS records added in the zone.
	//
	// example:
	//
	// 2
	RecordCount *int32 `json:"RecordCount,omitempty" xml:"RecordCount,omitempty"`
	// The description of the zone.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F73F41A3-B6DD-42CA-A793-FFF93277835D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the zone belongs.
	//
	// example:
	//
	// rg-acfmykd63gt****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Indicates whether the secondary Domain Name System (DNS) feature is enabled for the zone. Valid values:
	//
	// 	- **true**: The secondary DNS feature is enabled.
	//
	// 	- **false**: The secondary DNS feature is disabled.
	//
	// example:
	//
	// true
	SlaveDns *bool `json:"SlaveDns,omitempty" xml:"SlaveDns,omitempty"`
	// The time when the zone was last updated. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-01-24T06:35Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The time when the zone was last updated. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1516775741000
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	// The zone ID. This ID uniquely identifies the zone.
	//
	// example:
	//
	// df2d03865266bd9842306db586d3****
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The zone name.
	//
	// example:
	//
	// zone-test.cn
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
	// The tag added to the zone.
	//
	// example:
	//
	// pvtz
	ZoneTag *string `json:"ZoneTag,omitempty" xml:"ZoneTag,omitempty"`
	// The zone type. Valid values:
	//
	// 	- **AUTH_ZONE**: authoritative zone
	//
	// 	- **CLOUD_PRODUCT_ZONE**: authoritative zone for cloud services
	//
	// example:
	//
	// CLOUD_PRODUCT_ZONE
	ZoneType *string `json:"ZoneType,omitempty" xml:"ZoneType,omitempty"`
}

func (s DescribeZoneInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeZoneInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeZoneInfoResponseBody) GetBindEdgeDnsClusters() *DescribeZoneInfoResponseBodyBindEdgeDnsClusters {
	return s.BindEdgeDnsClusters
}

func (s *DescribeZoneInfoResponseBody) GetBindVpcs() *DescribeZoneInfoResponseBodyBindVpcs {
	return s.BindVpcs
}

func (s *DescribeZoneInfoResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeZoneInfoResponseBody) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeZoneInfoResponseBody) GetCreator() *string {
	return s.Creator
}

func (s *DescribeZoneInfoResponseBody) GetCreatorType() *string {
	return s.CreatorType
}

func (s *DescribeZoneInfoResponseBody) GetDnsGroup() *string {
	return s.DnsGroup
}

func (s *DescribeZoneInfoResponseBody) GetDnsGroupChanging() *bool {
	return s.DnsGroupChanging
}

func (s *DescribeZoneInfoResponseBody) GetIsPtr() *bool {
	return s.IsPtr
}

func (s *DescribeZoneInfoResponseBody) GetProxyPattern() *string {
	return s.ProxyPattern
}

func (s *DescribeZoneInfoResponseBody) GetRecordCount() *int32 {
	return s.RecordCount
}

func (s *DescribeZoneInfoResponseBody) GetRemark() *string {
	return s.Remark
}

func (s *DescribeZoneInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeZoneInfoResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeZoneInfoResponseBody) GetSlaveDns() *bool {
	return s.SlaveDns
}

func (s *DescribeZoneInfoResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeZoneInfoResponseBody) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *DescribeZoneInfoResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeZoneInfoResponseBody) GetZoneName() *string {
	return s.ZoneName
}

func (s *DescribeZoneInfoResponseBody) GetZoneTag() *string {
	return s.ZoneTag
}

func (s *DescribeZoneInfoResponseBody) GetZoneType() *string {
	return s.ZoneType
}

func (s *DescribeZoneInfoResponseBody) SetBindEdgeDnsClusters(v *DescribeZoneInfoResponseBodyBindEdgeDnsClusters) *DescribeZoneInfoResponseBody {
	s.BindEdgeDnsClusters = v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetBindVpcs(v *DescribeZoneInfoResponseBodyBindVpcs) *DescribeZoneInfoResponseBody {
	s.BindVpcs = v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetCreateTime(v string) *DescribeZoneInfoResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetCreateTimestamp(v int64) *DescribeZoneInfoResponseBody {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetCreator(v string) *DescribeZoneInfoResponseBody {
	s.Creator = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetCreatorType(v string) *DescribeZoneInfoResponseBody {
	s.CreatorType = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetDnsGroup(v string) *DescribeZoneInfoResponseBody {
	s.DnsGroup = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetDnsGroupChanging(v bool) *DescribeZoneInfoResponseBody {
	s.DnsGroupChanging = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetIsPtr(v bool) *DescribeZoneInfoResponseBody {
	s.IsPtr = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetProxyPattern(v string) *DescribeZoneInfoResponseBody {
	s.ProxyPattern = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetRecordCount(v int32) *DescribeZoneInfoResponseBody {
	s.RecordCount = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetRemark(v string) *DescribeZoneInfoResponseBody {
	s.Remark = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetRequestId(v string) *DescribeZoneInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetResourceGroupId(v string) *DescribeZoneInfoResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetSlaveDns(v bool) *DescribeZoneInfoResponseBody {
	s.SlaveDns = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetUpdateTime(v string) *DescribeZoneInfoResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetUpdateTimestamp(v int64) *DescribeZoneInfoResponseBody {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetZoneId(v string) *DescribeZoneInfoResponseBody {
	s.ZoneId = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetZoneName(v string) *DescribeZoneInfoResponseBody {
	s.ZoneName = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetZoneTag(v string) *DescribeZoneInfoResponseBody {
	s.ZoneTag = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetZoneType(v string) *DescribeZoneInfoResponseBody {
	s.ZoneType = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeZoneInfoResponseBodyBindEdgeDnsClusters struct {
	EdgeDnsCluster []*DescribeZoneInfoResponseBodyBindEdgeDnsClustersEdgeDnsCluster `json:"EdgeDnsCluster,omitempty" xml:"EdgeDnsCluster,omitempty" type:"Repeated"`
}

func (s DescribeZoneInfoResponseBodyBindEdgeDnsClusters) String() string {
	return dara.Prettify(s)
}

func (s DescribeZoneInfoResponseBodyBindEdgeDnsClusters) GoString() string {
	return s.String()
}

func (s *DescribeZoneInfoResponseBodyBindEdgeDnsClusters) GetEdgeDnsCluster() []*DescribeZoneInfoResponseBodyBindEdgeDnsClustersEdgeDnsCluster {
	return s.EdgeDnsCluster
}

func (s *DescribeZoneInfoResponseBodyBindEdgeDnsClusters) SetEdgeDnsCluster(v []*DescribeZoneInfoResponseBodyBindEdgeDnsClustersEdgeDnsCluster) *DescribeZoneInfoResponseBodyBindEdgeDnsClusters {
	s.EdgeDnsCluster = v
	return s
}

func (s *DescribeZoneInfoResponseBodyBindEdgeDnsClusters) Validate() error {
	return dara.Validate(s)
}

type DescribeZoneInfoResponseBodyBindEdgeDnsClustersEdgeDnsCluster struct {
	ClusterId     *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName   *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	ClusterUserId *int64  `json:"ClusterUserId,omitempty" xml:"ClusterUserId,omitempty"`
}

func (s DescribeZoneInfoResponseBodyBindEdgeDnsClustersEdgeDnsCluster) String() string {
	return dara.Prettify(s)
}

func (s DescribeZoneInfoResponseBodyBindEdgeDnsClustersEdgeDnsCluster) GoString() string {
	return s.String()
}

func (s *DescribeZoneInfoResponseBodyBindEdgeDnsClustersEdgeDnsCluster) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeZoneInfoResponseBodyBindEdgeDnsClustersEdgeDnsCluster) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeZoneInfoResponseBodyBindEdgeDnsClustersEdgeDnsCluster) GetClusterUserId() *int64 {
	return s.ClusterUserId
}

func (s *DescribeZoneInfoResponseBodyBindEdgeDnsClustersEdgeDnsCluster) SetClusterId(v string) *DescribeZoneInfoResponseBodyBindEdgeDnsClustersEdgeDnsCluster {
	s.ClusterId = &v
	return s
}

func (s *DescribeZoneInfoResponseBodyBindEdgeDnsClustersEdgeDnsCluster) SetClusterName(v string) *DescribeZoneInfoResponseBodyBindEdgeDnsClustersEdgeDnsCluster {
	s.ClusterName = &v
	return s
}

func (s *DescribeZoneInfoResponseBodyBindEdgeDnsClustersEdgeDnsCluster) SetClusterUserId(v int64) *DescribeZoneInfoResponseBodyBindEdgeDnsClustersEdgeDnsCluster {
	s.ClusterUserId = &v
	return s
}

func (s *DescribeZoneInfoResponseBodyBindEdgeDnsClustersEdgeDnsCluster) Validate() error {
	return dara.Validate(s)
}

type DescribeZoneInfoResponseBodyBindVpcs struct {
	Vpc []*DescribeZoneInfoResponseBodyBindVpcsVpc `json:"Vpc,omitempty" xml:"Vpc,omitempty" type:"Repeated"`
}

func (s DescribeZoneInfoResponseBodyBindVpcs) String() string {
	return dara.Prettify(s)
}

func (s DescribeZoneInfoResponseBodyBindVpcs) GoString() string {
	return s.String()
}

func (s *DescribeZoneInfoResponseBodyBindVpcs) GetVpc() []*DescribeZoneInfoResponseBodyBindVpcsVpc {
	return s.Vpc
}

func (s *DescribeZoneInfoResponseBodyBindVpcs) SetVpc(v []*DescribeZoneInfoResponseBodyBindVpcsVpc) *DescribeZoneInfoResponseBodyBindVpcs {
	s.Vpc = v
	return s
}

func (s *DescribeZoneInfoResponseBodyBindVpcs) Validate() error {
	return dara.Validate(s)
}

type DescribeZoneInfoResponseBodyBindVpcsVpc struct {
	// The region ID of the VPC.
	//
	// example:
	//
	// cn-heyuan
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the region where the VPC resides.
	//
	// example:
	//
	// China (Heyuan)
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	// The VPC ID. This ID uniquely identifies the VPC.
	//
	// example:
	//
	// vpc-f8zvrvr1payllgz38****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The VPC name.
	//
	// example:
	//
	// vpc_test
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
	// The VPC type. Valid values:
	//
	// 	- STANDARD: standard VPC
	//
	// 	- EDS: Elastic Desktop Service (EDS) workspace VPC
	//
	// example:
	//
	// STANDARD
	VpcType *string `json:"VpcType,omitempty" xml:"VpcType,omitempty"`
	// The user ID to which the VPC belongs. If null is returned, the VPC belongs to the current user.
	//
	// example:
	//
	// 141339776561****
	VpcUserId *int64 `json:"VpcUserId,omitempty" xml:"VpcUserId,omitempty"`
}

func (s DescribeZoneInfoResponseBodyBindVpcsVpc) String() string {
	return dara.Prettify(s)
}

func (s DescribeZoneInfoResponseBodyBindVpcsVpc) GoString() string {
	return s.String()
}

func (s *DescribeZoneInfoResponseBodyBindVpcsVpc) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeZoneInfoResponseBodyBindVpcsVpc) GetRegionName() *string {
	return s.RegionName
}

func (s *DescribeZoneInfoResponseBodyBindVpcsVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeZoneInfoResponseBodyBindVpcsVpc) GetVpcName() *string {
	return s.VpcName
}

func (s *DescribeZoneInfoResponseBodyBindVpcsVpc) GetVpcType() *string {
	return s.VpcType
}

func (s *DescribeZoneInfoResponseBodyBindVpcsVpc) GetVpcUserId() *int64 {
	return s.VpcUserId
}

func (s *DescribeZoneInfoResponseBodyBindVpcsVpc) SetRegionId(v string) *DescribeZoneInfoResponseBodyBindVpcsVpc {
	s.RegionId = &v
	return s
}

func (s *DescribeZoneInfoResponseBodyBindVpcsVpc) SetRegionName(v string) *DescribeZoneInfoResponseBodyBindVpcsVpc {
	s.RegionName = &v
	return s
}

func (s *DescribeZoneInfoResponseBodyBindVpcsVpc) SetVpcId(v string) *DescribeZoneInfoResponseBodyBindVpcsVpc {
	s.VpcId = &v
	return s
}

func (s *DescribeZoneInfoResponseBodyBindVpcsVpc) SetVpcName(v string) *DescribeZoneInfoResponseBodyBindVpcsVpc {
	s.VpcName = &v
	return s
}

func (s *DescribeZoneInfoResponseBodyBindVpcsVpc) SetVpcType(v string) *DescribeZoneInfoResponseBodyBindVpcsVpc {
	s.VpcType = &v
	return s
}

func (s *DescribeZoneInfoResponseBodyBindVpcsVpc) SetVpcUserId(v int64) *DescribeZoneInfoResponseBodyBindVpcsVpc {
	s.VpcUserId = &v
	return s
}

func (s *DescribeZoneInfoResponseBodyBindVpcsVpc) Validate() error {
	return dara.Validate(s)
}
