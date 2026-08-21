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
	BindVpcs            *DescribeZoneInfoResponseBodyBindVpcs            `json:"BindVpcs,omitempty" xml:"BindVpcs,omitempty" type:"Struct"`
	// The time when the zone was created.
	//
	// example:
	//
	// 2024-07-15T09:40Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The timestamp that indicates when the zone was created.
	//
	// example:
	//
	// 1721036404000
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
	// The DNS group. Valid values:
	//
	// - **NORMAL_ZONE**: Regular group
	//
	// - **FAST_ZONE**: Fast group
	//
	// example:
	//
	// FAST_ZONE
	DnsGroup *string `json:"DnsGroup,omitempty" xml:"DnsGroup,omitempty"`
	// The status of switching the DNS group for the zone. You can switch a zone between the regular and fast groups. Valid values:
	//
	// - true: The DNS group is being switched.
	//
	// - false: The DNS group is not being switched.
	//
	// example:
	//
	// false
	DnsGroupChanging *bool `json:"DnsGroupChanging,omitempty" xml:"DnsGroupChanging,omitempty"`
	// Indicates whether the zone is a reverse lookup zone. Valid values:
	//
	// - true: The zone is a reverse lookup zone.
	//
	// - false: The zone is not a reverse lookup zone.
	//
	// example:
	//
	// false
	IsPtr *bool `json:"IsPtr,omitempty" xml:"IsPtr,omitempty"`
	// Indicates whether the recursive proxy for subdomains is enabled.
	//
	// - ZONE: Disabled. If a non-existent subdomain is resolved, an NXDOMAIN response is returned to indicate that the subdomain does not exist.
	//
	// - RECORD: Enabled. If a non-existent subdomain is resolved, the system queries the forwarding and recursion modules in sequence and uses the final result to respond to the DNS query.
	//
	// example:
	//
	// ZONE
	ProxyPattern *string `json:"ProxyPattern,omitempty" xml:"ProxyPattern,omitempty"`
	// The total number of DNS records in the zone.
	//
	// example:
	//
	// 5
	RecordCount *int32 `json:"RecordCount,omitempty" xml:"RecordCount,omitempty"`
	// The remarks.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The unique ID of the request.
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
	// Indicates whether secondary DNS is enabled. Valid values:
	//
	// - **true**: Enabled
	//
	// - **false**: Disabled
	//
	// example:
	//
	// false
	SlaveDns *bool `json:"SlaveDns,omitempty" xml:"SlaveDns,omitempty"`
	// The time when the zone was last updated.
	//
	// example:
	//
	// 2024-07-22T09:39Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The timestamp that indicates when the zone was last updated.
	//
	// example:
	//
	// 1721641148000
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	// The unique ID of the zone.
	//
	// example:
	//
	// df2d03865266bd9842306db586d3****
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The name of the zone.
	//
	// example:
	//
	// zone-test.cn
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
	// The tag of the zone.
	//
	// example:
	//
	// pvtz
	ZoneTag *string `json:"ZoneTag,omitempty" xml:"ZoneTag,omitempty"`
	// The type of the zone. Valid values:
	//
	// - **AUTH_ZONE**: Authoritative zone.
	//
	// - **CLOUD_PRODUCT_ZONE**: Authoritative zone for a cloud product.
	//
	// example:
	//
	// AUTH_ZONE
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
	if s.BindEdgeDnsClusters != nil {
		if err := s.BindEdgeDnsClusters.Validate(); err != nil {
			return err
		}
	}
	if s.BindVpcs != nil {
		if err := s.BindVpcs.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.EdgeDnsCluster != nil {
		for _, item := range s.EdgeDnsCluster {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
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
	if s.Vpc != nil {
		for _, item := range s.Vpc {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeZoneInfoResponseBodyBindVpcsVpc struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	VpcId      *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcName    *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
	VpcType    *string `json:"VpcType,omitempty" xml:"VpcType,omitempty"`
	VpcUserId  *int64  `json:"VpcUserId,omitempty" xml:"VpcUserId,omitempty"`
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
