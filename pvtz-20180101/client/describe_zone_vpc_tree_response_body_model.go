// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeZoneVpcTreeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeZoneVpcTreeResponseBody
	GetRequestId() *string
	SetZones(v *DescribeZoneVpcTreeResponseBodyZones) *DescribeZoneVpcTreeResponseBody
	GetZones() *DescribeZoneVpcTreeResponseBodyZones
}

type DescribeZoneVpcTreeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7B07FBC3-3A53-4939-A3C6-2BDFE407BAB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The zones.
	Zones *DescribeZoneVpcTreeResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
}

func (s DescribeZoneVpcTreeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeZoneVpcTreeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeZoneVpcTreeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeZoneVpcTreeResponseBody) GetZones() *DescribeZoneVpcTreeResponseBodyZones {
	return s.Zones
}

func (s *DescribeZoneVpcTreeResponseBody) SetRequestId(v string) *DescribeZoneVpcTreeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBody) SetZones(v *DescribeZoneVpcTreeResponseBodyZones) *DescribeZoneVpcTreeResponseBody {
	s.Zones = v
	return s
}

func (s *DescribeZoneVpcTreeResponseBody) Validate() error {
	if s.Zones != nil {
		if err := s.Zones.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeZoneVpcTreeResponseBodyZones struct {
	Zone []*DescribeZoneVpcTreeResponseBodyZonesZone `json:"Zone,omitempty" xml:"Zone,omitempty" type:"Repeated"`
}

func (s DescribeZoneVpcTreeResponseBodyZones) String() string {
	return dara.Prettify(s)
}

func (s DescribeZoneVpcTreeResponseBodyZones) GoString() string {
	return s.String()
}

func (s *DescribeZoneVpcTreeResponseBodyZones) GetZone() []*DescribeZoneVpcTreeResponseBodyZonesZone {
	return s.Zone
}

func (s *DescribeZoneVpcTreeResponseBodyZones) SetZone(v []*DescribeZoneVpcTreeResponseBodyZonesZone) *DescribeZoneVpcTreeResponseBodyZones {
	s.Zone = v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZones) Validate() error {
	if s.Zone != nil {
		for _, item := range s.Zone {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeZoneVpcTreeResponseBodyZonesZone struct {
	// The time when the zone was created. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-09-18T08:20Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the zone was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1568794812000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The creator of the zone.
	//
	// example:
	//
	// 141339776561****
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The operator type.
	//
	// example:
	//
	// USER
	CreatorType *string `json:"CreatorType,omitempty" xml:"CreatorType,omitempty"`
	// The logical location of the built-in authoritative module in which the zone is added. Valid values:
	//
	// 	- NORMAL_ZONE: regular module
	//
	// 	- FAST_ZONE: acceleration module
	//
	// example:
	//
	// NORMAL_ZONE
	DnsGroup *string `json:"DnsGroup,omitempty" xml:"DnsGroup,omitempty"`
	// Indicates whether the zone is being removed to another logical location. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
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
	// The number of Domain Name System (DNS) records added for the zone.
	//
	// example:
	//
	// 1
	RecordCount *int32 `json:"RecordCount,omitempty" xml:"RecordCount,omitempty"`
	// The description of the zone.
	//
	// example:
	//
	// demo
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The time when the zone was last modified. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-09-18T08:20Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The time when the zone was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1568794834000
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	// The VPCs associated with the zone.
	Vpcs *DescribeZoneVpcTreeResponseBodyZonesZoneVpcs `json:"Vpcs,omitempty" xml:"Vpcs,omitempty" type:"Struct"`
	// The zone ID. This ID uniquely identifies the zone.
	//
	// example:
	//
	// 6d83e3b31aa60ca4aaa7161f1b6baa95
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The zone name.
	//
	// example:
	//
	// example.com
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
	// The type of the cloud service.
	//
	//
	// **Valid values:**
	//
	// 	- If ZoneType is set to AUTH_ZONE, no value is returned for this parameter.
	//
	// 	- If ZoneType is set to CLOUD_PRODUCT_ZONE, the type of the cloud service is returned.
	//
	// example:
	//
	// BLINK
	ZoneTag *string `json:"ZoneTag,omitempty" xml:"ZoneTag,omitempty"`
	// The zone type. Valid values:
	//
	// 	- AUTH_ZONE: authoritative zone
	//
	// 	- CLOUD_PRODUCT_ZONE: authoritative zone for cloud services
	//
	// example:
	//
	// AUTH_ZONE
	ZoneType *string `json:"ZoneType,omitempty" xml:"ZoneType,omitempty"`
}

func (s DescribeZoneVpcTreeResponseBodyZonesZone) String() string {
	return dara.Prettify(s)
}

func (s DescribeZoneVpcTreeResponseBodyZonesZone) GoString() string {
	return s.String()
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) GetCreator() *string {
	return s.Creator
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) GetCreatorType() *string {
	return s.CreatorType
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) GetDnsGroup() *string {
	return s.DnsGroup
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) GetDnsGroupChanging() *bool {
	return s.DnsGroupChanging
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) GetIsPtr() *bool {
	return s.IsPtr
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) GetRecordCount() *int32 {
	return s.RecordCount
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) GetRemark() *string {
	return s.Remark
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) GetVpcs() *DescribeZoneVpcTreeResponseBodyZonesZoneVpcs {
	return s.Vpcs
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) GetZoneName() *string {
	return s.ZoneName
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) GetZoneTag() *string {
	return s.ZoneTag
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) GetZoneType() *string {
	return s.ZoneType
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) SetCreateTime(v string) *DescribeZoneVpcTreeResponseBodyZonesZone {
	s.CreateTime = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) SetCreateTimestamp(v int64) *DescribeZoneVpcTreeResponseBodyZonesZone {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) SetCreator(v string) *DescribeZoneVpcTreeResponseBodyZonesZone {
	s.Creator = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) SetCreatorType(v string) *DescribeZoneVpcTreeResponseBodyZonesZone {
	s.CreatorType = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) SetDnsGroup(v string) *DescribeZoneVpcTreeResponseBodyZonesZone {
	s.DnsGroup = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) SetDnsGroupChanging(v bool) *DescribeZoneVpcTreeResponseBodyZonesZone {
	s.DnsGroupChanging = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) SetIsPtr(v bool) *DescribeZoneVpcTreeResponseBodyZonesZone {
	s.IsPtr = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) SetRecordCount(v int32) *DescribeZoneVpcTreeResponseBodyZonesZone {
	s.RecordCount = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) SetRemark(v string) *DescribeZoneVpcTreeResponseBodyZonesZone {
	s.Remark = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) SetUpdateTime(v string) *DescribeZoneVpcTreeResponseBodyZonesZone {
	s.UpdateTime = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) SetUpdateTimestamp(v int64) *DescribeZoneVpcTreeResponseBodyZonesZone {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) SetVpcs(v *DescribeZoneVpcTreeResponseBodyZonesZoneVpcs) *DescribeZoneVpcTreeResponseBodyZonesZone {
	s.Vpcs = v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) SetZoneId(v string) *DescribeZoneVpcTreeResponseBodyZonesZone {
	s.ZoneId = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) SetZoneName(v string) *DescribeZoneVpcTreeResponseBodyZonesZone {
	s.ZoneName = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) SetZoneTag(v string) *DescribeZoneVpcTreeResponseBodyZonesZone {
	s.ZoneTag = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) SetZoneType(v string) *DescribeZoneVpcTreeResponseBodyZonesZone {
	s.ZoneType = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) Validate() error {
	if s.Vpcs != nil {
		if err := s.Vpcs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeZoneVpcTreeResponseBodyZonesZoneVpcs struct {
	Vpc []*DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc `json:"Vpc,omitempty" xml:"Vpc,omitempty" type:"Repeated"`
}

func (s DescribeZoneVpcTreeResponseBodyZonesZoneVpcs) String() string {
	return dara.Prettify(s)
}

func (s DescribeZoneVpcTreeResponseBodyZonesZoneVpcs) GoString() string {
	return s.String()
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZoneVpcs) GetVpc() []*DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc {
	return s.Vpc
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZoneVpcs) SetVpc(v []*DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc) *DescribeZoneVpcTreeResponseBodyZonesZoneVpcs {
	s.Vpc = v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZoneVpcs) Validate() error {
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

type DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc struct {
	// The region ID of the VPC.
	//
	// example:
	//
	// cn-heyuan
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the region to which the VPC belongs.
	//
	// example:
	//
	// China (Heyuan)
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	// The VPC ID. The unique ID of the VPC.
	//
	// example:
	//
	// vpc-f8zvrvr1payllgz38****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The VPC name.
	//
	// example:
	//
	// demo-vpc
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
}

func (s DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc) String() string {
	return dara.Prettify(s)
}

func (s DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc) GoString() string {
	return s.String()
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc) GetRegionName() *string {
	return s.RegionName
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc) GetVpcName() *string {
	return s.VpcName
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc) GetVpcType() *string {
	return s.VpcType
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc) SetRegionId(v string) *DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc {
	s.RegionId = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc) SetRegionName(v string) *DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc {
	s.RegionName = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc) SetVpcId(v string) *DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc {
	s.VpcId = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc) SetVpcName(v string) *DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc {
	s.VpcName = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc) SetVpcType(v string) *DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc {
	s.VpcType = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc) Validate() error {
	return dara.Validate(s)
}
