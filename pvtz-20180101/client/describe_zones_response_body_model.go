// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeZonesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeZonesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeZonesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeZonesResponseBody
	GetRequestId() *string
	SetTotalItems(v int32) *DescribeZonesResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *DescribeZonesResponseBody
	GetTotalPages() *int32
	SetZones(v *DescribeZonesResponseBodyZones) *DescribeZonesResponseBody
	GetZones() *DescribeZonesResponseBodyZones
}

type DescribeZonesResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 3
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// The total number of returned pages.
	//
	// example:
	//
	// 3
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	// The zones.
	Zones *DescribeZonesResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
}

func (s DescribeZonesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeZonesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeZonesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeZonesResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *DescribeZonesResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *DescribeZonesResponseBody) GetZones() *DescribeZonesResponseBodyZones {
	return s.Zones
}

func (s *DescribeZonesResponseBody) SetPageNumber(v int32) *DescribeZonesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeZonesResponseBody) SetPageSize(v int32) *DescribeZonesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeZonesResponseBody) SetRequestId(v string) *DescribeZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeZonesResponseBody) SetTotalItems(v int32) *DescribeZonesResponseBody {
	s.TotalItems = &v
	return s
}

func (s *DescribeZonesResponseBody) SetTotalPages(v int32) *DescribeZonesResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeZonesResponseBody) SetZones(v *DescribeZonesResponseBodyZones) *DescribeZonesResponseBody {
	s.Zones = v
	return s
}

func (s *DescribeZonesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeZonesResponseBodyZones struct {
	Zone []*DescribeZonesResponseBodyZonesZone `json:"Zone,omitempty" xml:"Zone,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBodyZones) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesResponseBodyZones) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZones) GetZone() []*DescribeZonesResponseBodyZonesZone {
	return s.Zone
}

func (s *DescribeZonesResponseBodyZones) SetZone(v []*DescribeZonesResponseBodyZonesZone) *DescribeZonesResponseBodyZones {
	s.Zone = v
	return s
}

func (s *DescribeZonesResponseBodyZones) Validate() error {
	return dara.Validate(s)
}

type DescribeZonesResponseBodyZonesZone struct {
	// The time when the zone was created. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-12-28T13:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the zone was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1514466483000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The creator of the zone.
	//
	// example:
	//
	// 546356****
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The account type. Valid values:
	//
	// 	- **CUSTOMER**: Alibaba Cloud account
	//
	// 	- **SUB**: RAM user
	//
	// 	- **STS**: assumed role that obtains the Security Token Service (STS) token of a RAM role
	//
	// 	- **OTHER**: other types
	//
	// example:
	//
	// SUB
	CreatorSubType *string `json:"CreatorSubType,omitempty" xml:"CreatorSubType,omitempty"`
	// The logical location type of the built-in authoritative module in which the zone is added. Valid values:
	//
	// 	- **NORMAL_ZONE**: regular module
	//
	// 	- **FAST_ZONE**: acceleration module
	//
	// example:
	//
	// NORMAL_ZONE
	DnsGroup *string `json:"DnsGroup,omitempty" xml:"DnsGroup,omitempty"`
	// Indicates whether the zone is being removed to another logical location. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	DnsGroupChanging *bool `json:"DnsGroupChanging,omitempty" xml:"DnsGroupChanging,omitempty"`
	// Indicates whether the zone is a reverse lookup zone. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsPtr *bool `json:"IsPtr,omitempty" xml:"IsPtr,omitempty"`
	// Indicates whether the recursive resolution proxy for subdomain names is enabled. Valid values:
	//
	// 	- **ZONE**: The recursive resolution proxy for subdomain names is disabled. In this case, NXDOMAIN is returned if the queried domain name does not exist in the zone.
	//
	// 	- **RECORD**: The recursive resolution proxy for subdomain names is enabled. In this case, if the queried domain name does not exist in the zone, DNS requests are recursively forwarded to the forward module and then to the recursion module until DNS results are returned.
	//
	// example:
	//
	// ZONE
	ProxyPattern *string `json:"ProxyPattern,omitempty" xml:"ProxyPattern,omitempty"`
	// The number of Domain Name System (DNS) records added in the zone.
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
	// The ID of the resource group to which the zone belongs.
	//
	// example:
	//
	// rg-aekz2qj7awz****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags added to the zone.
	ResourceTags   *DescribeZonesResponseBodyZonesZoneResourceTags `json:"ResourceTags,omitempty" xml:"ResourceTags,omitempty" type:"Struct"`
	SlaveDnsStatus *string                                         `json:"SlaveDnsStatus,omitempty" xml:"SlaveDnsStatus,omitempty"`
	// The time when the zone was last modified. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-01-03T08:57Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The time when the DNS record was updated. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since 00:00:00 UTC on January 1, 1970.
	//
	// example:
	//
	// 1514969843000
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	// The zone ID. This ID uniquely identifies the zone.
	//
	// example:
	//
	// 6d83e3b31aa60ca4aaa7161f1b6b**95
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The name of the zone.
	//
	// example:
	//
	// test.com
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
	// The type of the cloud service. Valid values:
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
	// 	- **AUTH_ZONE**: authoritative zone
	//
	// 	- **CLOUD_PRODUCT_ZONE**: authoritative zone for cloud services
	//
	// example:
	//
	// CLOUD_PRODUCT_ZONE
	ZoneType *string `json:"ZoneType,omitempty" xml:"ZoneType,omitempty"`
}

func (s DescribeZonesResponseBodyZonesZone) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZone) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZone) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeZonesResponseBodyZonesZone) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeZonesResponseBodyZonesZone) GetCreator() *string {
	return s.Creator
}

func (s *DescribeZonesResponseBodyZonesZone) GetCreatorSubType() *string {
	return s.CreatorSubType
}

func (s *DescribeZonesResponseBodyZonesZone) GetDnsGroup() *string {
	return s.DnsGroup
}

func (s *DescribeZonesResponseBodyZonesZone) GetDnsGroupChanging() *bool {
	return s.DnsGroupChanging
}

func (s *DescribeZonesResponseBodyZonesZone) GetIsPtr() *bool {
	return s.IsPtr
}

func (s *DescribeZonesResponseBodyZonesZone) GetProxyPattern() *string {
	return s.ProxyPattern
}

func (s *DescribeZonesResponseBodyZonesZone) GetRecordCount() *int32 {
	return s.RecordCount
}

func (s *DescribeZonesResponseBodyZonesZone) GetRemark() *string {
	return s.Remark
}

func (s *DescribeZonesResponseBodyZonesZone) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeZonesResponseBodyZonesZone) GetResourceTags() *DescribeZonesResponseBodyZonesZoneResourceTags {
	return s.ResourceTags
}

func (s *DescribeZonesResponseBodyZonesZone) GetSlaveDnsStatus() *string {
	return s.SlaveDnsStatus
}

func (s *DescribeZonesResponseBodyZonesZone) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeZonesResponseBodyZonesZone) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *DescribeZonesResponseBodyZonesZone) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeZonesResponseBodyZonesZone) GetZoneName() *string {
	return s.ZoneName
}

func (s *DescribeZonesResponseBodyZonesZone) GetZoneTag() *string {
	return s.ZoneTag
}

func (s *DescribeZonesResponseBodyZonesZone) GetZoneType() *string {
	return s.ZoneType
}

func (s *DescribeZonesResponseBodyZonesZone) SetCreateTime(v string) *DescribeZonesResponseBodyZonesZone {
	s.CreateTime = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetCreateTimestamp(v int64) *DescribeZonesResponseBodyZonesZone {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetCreator(v string) *DescribeZonesResponseBodyZonesZone {
	s.Creator = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetCreatorSubType(v string) *DescribeZonesResponseBodyZonesZone {
	s.CreatorSubType = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetDnsGroup(v string) *DescribeZonesResponseBodyZonesZone {
	s.DnsGroup = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetDnsGroupChanging(v bool) *DescribeZonesResponseBodyZonesZone {
	s.DnsGroupChanging = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetIsPtr(v bool) *DescribeZonesResponseBodyZonesZone {
	s.IsPtr = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetProxyPattern(v string) *DescribeZonesResponseBodyZonesZone {
	s.ProxyPattern = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetRecordCount(v int32) *DescribeZonesResponseBodyZonesZone {
	s.RecordCount = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetRemark(v string) *DescribeZonesResponseBodyZonesZone {
	s.Remark = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetResourceGroupId(v string) *DescribeZonesResponseBodyZonesZone {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetResourceTags(v *DescribeZonesResponseBodyZonesZoneResourceTags) *DescribeZonesResponseBodyZonesZone {
	s.ResourceTags = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetSlaveDnsStatus(v string) *DescribeZonesResponseBodyZonesZone {
	s.SlaveDnsStatus = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetUpdateTime(v string) *DescribeZonesResponseBodyZonesZone {
	s.UpdateTime = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetUpdateTimestamp(v int64) *DescribeZonesResponseBodyZonesZone {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetZoneId(v string) *DescribeZonesResponseBodyZonesZone {
	s.ZoneId = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetZoneName(v string) *DescribeZonesResponseBodyZonesZone {
	s.ZoneName = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetZoneTag(v string) *DescribeZonesResponseBodyZonesZone {
	s.ZoneTag = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetZoneType(v string) *DescribeZonesResponseBodyZonesZone {
	s.ZoneType = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) Validate() error {
	return dara.Validate(s)
}

type DescribeZonesResponseBodyZonesZoneResourceTags struct {
	ResourceTag []*DescribeZonesResponseBodyZonesZoneResourceTagsResourceTag `json:"ResourceTag,omitempty" xml:"ResourceTag,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBodyZonesZoneResourceTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZoneResourceTags) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZoneResourceTags) GetResourceTag() []*DescribeZonesResponseBodyZonesZoneResourceTagsResourceTag {
	return s.ResourceTag
}

func (s *DescribeZonesResponseBodyZonesZoneResourceTags) SetResourceTag(v []*DescribeZonesResponseBodyZonesZoneResourceTagsResourceTag) *DescribeZonesResponseBodyZonesZoneResourceTags {
	s.ResourceTag = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZoneResourceTags) Validate() error {
	return dara.Validate(s)
}

type DescribeZonesResponseBodyZonesZoneResourceTagsResourceTag struct {
	// The key of tag N added to the zone.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N added to the zone.
	//
	// example:
	//
	// daily
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeZonesResponseBodyZonesZoneResourceTagsResourceTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZoneResourceTagsResourceTag) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZoneResourceTagsResourceTag) GetKey() *string {
	return s.Key
}

func (s *DescribeZonesResponseBodyZonesZoneResourceTagsResourceTag) GetValue() *string {
	return s.Value
}

func (s *DescribeZonesResponseBodyZonesZoneResourceTagsResourceTag) SetKey(v string) *DescribeZonesResponseBodyZonesZoneResourceTagsResourceTag {
	s.Key = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZoneResourceTagsResourceTag) SetValue(v string) *DescribeZonesResponseBodyZonesZoneResourceTagsResourceTag {
	s.Value = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZoneResourceTagsResourceTag) Validate() error {
	return dara.Validate(s)
}
