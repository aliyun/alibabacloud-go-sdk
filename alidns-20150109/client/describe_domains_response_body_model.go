// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomains(v *DescribeDomainsResponseBodyDomains) *DescribeDomainsResponseBody
	GetDomains() *DescribeDomainsResponseBodyDomains
	SetPageNumber(v int64) *DescribeDomainsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeDomainsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeDomainsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeDomainsResponseBody
	GetTotalCount() *int64
}

type DescribeDomainsResponseBody struct {
	// The domain names.
	Domains *DescribeDomainsResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 2
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 68386699-8B9E-4D5B-BC4C-75A28F6C2A00
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of domain names.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDomainsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBody) GetDomains() *DescribeDomainsResponseBodyDomains {
	return s.Domains
}

func (s *DescribeDomainsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeDomainsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDomainsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDomainsResponseBody) SetDomains(v *DescribeDomainsResponseBodyDomains) *DescribeDomainsResponseBody {
	s.Domains = v
	return s
}

func (s *DescribeDomainsResponseBody) SetPageNumber(v int64) *DescribeDomainsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDomainsResponseBody) SetPageSize(v int64) *DescribeDomainsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainsResponseBody) SetRequestId(v string) *DescribeDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainsResponseBody) SetTotalCount(v int64) *DescribeDomainsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDomainsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainsResponseBodyDomains struct {
	Domain []*DescribeDomainsResponseBodyDomainsDomain `json:"Domain,omitempty" xml:"Domain,omitempty" type:"Repeated"`
}

func (s DescribeDomainsResponseBodyDomains) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainsResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBodyDomains) GetDomain() []*DescribeDomainsResponseBodyDomainsDomain {
	return s.Domain
}

func (s *DescribeDomainsResponseBodyDomains) SetDomain(v []*DescribeDomainsResponseBodyDomainsDomain) *DescribeDomainsResponseBodyDomains {
	s.Domain = v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainsResponseBodyDomainsDomain struct {
	// Indicates whether the domain name was registered with Alibaba Cloud.
	//
	// example:
	//
	// true
	AliDomain *bool `json:"AliDomain,omitempty" xml:"AliDomain,omitempty"`
	// The time when the domain name was added. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-01-30T05:25Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the domain name was added. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1660546144000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The names of the DNS servers configured for the domain name assigned by Alibaba Cloud DNS.
	DnsServers *DescribeDomainsResponseBodyDomainsDomainDnsServers `json:"DnsServers,omitempty" xml:"DnsServers,omitempty" type:"Struct"`
	// The ID of the domain name.
	//
	// example:
	//
	// 00efd71a-770e-4255-b54e-6fe5659baffe
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// Indicates whether the DNS traffic analysis feature is enabled for the domain name. Valid values:
	//
	// 	- OPEN
	//
	// 	- CLOSE
	//
	// example:
	//
	// OPEN
	DomainLoggingSwitchStatus *string `json:"DomainLoggingSwitchStatus,omitempty" xml:"DomainLoggingSwitchStatus,omitempty"`
	// The domain name.
	//
	// example:
	//
	// abc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The ID of the domain name group.
	//
	// example:
	//
	// 2223
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the domain name group.
	//
	// example:
	//
	// mygroup
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The time when the Alibaba Cloud DNS instance expires. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-03-14T16:00Z
	InstanceEndTime *string `json:"InstanceEndTime,omitempty" xml:"InstanceEndTime,omitempty"`
	// Indicates whether the Alibaba Cloud DNS instance expires.
	//
	// example:
	//
	// false
	InstanceExpired *bool `json:"InstanceExpired,omitempty" xml:"InstanceExpired,omitempty"`
	// The ID of the Alibaba Cloud DNS instance.
	//
	// example:
	//
	// i-7bg
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The Punycode for the domain name. This parameter is returned only for Chinese domain names.
	//
	// example:
	//
	// abc.com
	PunyCode *string `json:"PunyCode,omitempty" xml:"PunyCode,omitempty"`
	// The number of Domain Name System (DNS) records added for the domain name.
	//
	// example:
	//
	// 100
	RecordCount *int64 `json:"RecordCount,omitempty" xml:"RecordCount,omitempty"`
	// The email address of the registrant.
	//
	// example:
	//
	// test@example.com
	RegistrantEmail *string `json:"RegistrantEmail,omitempty" xml:"RegistrantEmail,omitempty"`
	// The description of the domain name.
	//
	// example:
	//
	// remark
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The ID of the resource group to which the domain name belongs.
	//
	// example:
	//
	// rg-acf
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SlaveDnsStatus  *string `json:"SlaveDnsStatus,omitempty" xml:"SlaveDnsStatus,omitempty"`
	// Indicates whether the domain name was added to favorites.
	//
	// example:
	//
	// true
	Starmark *bool `json:"Starmark,omitempty" xml:"Starmark,omitempty"`
	// The tags added to the resource.
	Tags *DescribeDomainsResponseBodyDomainsDomainTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The edition code of Alibaba Cloud DNS.
	//
	// example:
	//
	// version_enterprise_basic
	VersionCode *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	// The edition of Alibaba Cloud DNS.
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s DescribeDomainsResponseBodyDomainsDomain) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainsResponseBodyDomainsDomain) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBodyDomainsDomain) GetAliDomain() *bool {
	return s.AliDomain
}

func (s *DescribeDomainsResponseBodyDomainsDomain) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDomainsResponseBodyDomainsDomain) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeDomainsResponseBodyDomainsDomain) GetDnsServers() *DescribeDomainsResponseBodyDomainsDomainDnsServers {
	return s.DnsServers
}

func (s *DescribeDomainsResponseBodyDomainsDomain) GetDomainId() *string {
	return s.DomainId
}

func (s *DescribeDomainsResponseBodyDomainsDomain) GetDomainLoggingSwitchStatus() *string {
	return s.DomainLoggingSwitchStatus
}

func (s *DescribeDomainsResponseBodyDomainsDomain) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainsResponseBodyDomainsDomain) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeDomainsResponseBodyDomainsDomain) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeDomainsResponseBodyDomainsDomain) GetInstanceEndTime() *string {
	return s.InstanceEndTime
}

func (s *DescribeDomainsResponseBodyDomainsDomain) GetInstanceExpired() *bool {
	return s.InstanceExpired
}

func (s *DescribeDomainsResponseBodyDomainsDomain) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDomainsResponseBodyDomainsDomain) GetPunyCode() *string {
	return s.PunyCode
}

func (s *DescribeDomainsResponseBodyDomainsDomain) GetRecordCount() *int64 {
	return s.RecordCount
}

func (s *DescribeDomainsResponseBodyDomainsDomain) GetRegistrantEmail() *string {
	return s.RegistrantEmail
}

func (s *DescribeDomainsResponseBodyDomainsDomain) GetRemark() *string {
	return s.Remark
}

func (s *DescribeDomainsResponseBodyDomainsDomain) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDomainsResponseBodyDomainsDomain) GetSlaveDnsStatus() *string {
	return s.SlaveDnsStatus
}

func (s *DescribeDomainsResponseBodyDomainsDomain) GetStarmark() *bool {
	return s.Starmark
}

func (s *DescribeDomainsResponseBodyDomainsDomain) GetTags() *DescribeDomainsResponseBodyDomainsDomainTags {
	return s.Tags
}

func (s *DescribeDomainsResponseBodyDomainsDomain) GetVersionCode() *string {
	return s.VersionCode
}

func (s *DescribeDomainsResponseBodyDomainsDomain) GetVersionName() *string {
	return s.VersionName
}

func (s *DescribeDomainsResponseBodyDomainsDomain) SetAliDomain(v bool) *DescribeDomainsResponseBodyDomainsDomain {
	s.AliDomain = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomainsDomain) SetCreateTime(v string) *DescribeDomainsResponseBodyDomainsDomain {
	s.CreateTime = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomainsDomain) SetCreateTimestamp(v int64) *DescribeDomainsResponseBodyDomainsDomain {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomainsDomain) SetDnsServers(v *DescribeDomainsResponseBodyDomainsDomainDnsServers) *DescribeDomainsResponseBodyDomainsDomain {
	s.DnsServers = v
	return s
}

func (s *DescribeDomainsResponseBodyDomainsDomain) SetDomainId(v string) *DescribeDomainsResponseBodyDomainsDomain {
	s.DomainId = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomainsDomain) SetDomainLoggingSwitchStatus(v string) *DescribeDomainsResponseBodyDomainsDomain {
	s.DomainLoggingSwitchStatus = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomainsDomain) SetDomainName(v string) *DescribeDomainsResponseBodyDomainsDomain {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomainsDomain) SetGroupId(v string) *DescribeDomainsResponseBodyDomainsDomain {
	s.GroupId = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomainsDomain) SetGroupName(v string) *DescribeDomainsResponseBodyDomainsDomain {
	s.GroupName = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomainsDomain) SetInstanceEndTime(v string) *DescribeDomainsResponseBodyDomainsDomain {
	s.InstanceEndTime = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomainsDomain) SetInstanceExpired(v bool) *DescribeDomainsResponseBodyDomainsDomain {
	s.InstanceExpired = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomainsDomain) SetInstanceId(v string) *DescribeDomainsResponseBodyDomainsDomain {
	s.InstanceId = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomainsDomain) SetPunyCode(v string) *DescribeDomainsResponseBodyDomainsDomain {
	s.PunyCode = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomainsDomain) SetRecordCount(v int64) *DescribeDomainsResponseBodyDomainsDomain {
	s.RecordCount = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomainsDomain) SetRegistrantEmail(v string) *DescribeDomainsResponseBodyDomainsDomain {
	s.RegistrantEmail = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomainsDomain) SetRemark(v string) *DescribeDomainsResponseBodyDomainsDomain {
	s.Remark = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomainsDomain) SetResourceGroupId(v string) *DescribeDomainsResponseBodyDomainsDomain {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomainsDomain) SetSlaveDnsStatus(v string) *DescribeDomainsResponseBodyDomainsDomain {
	s.SlaveDnsStatus = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomainsDomain) SetStarmark(v bool) *DescribeDomainsResponseBodyDomainsDomain {
	s.Starmark = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomainsDomain) SetTags(v *DescribeDomainsResponseBodyDomainsDomainTags) *DescribeDomainsResponseBodyDomainsDomain {
	s.Tags = v
	return s
}

func (s *DescribeDomainsResponseBodyDomainsDomain) SetVersionCode(v string) *DescribeDomainsResponseBodyDomainsDomain {
	s.VersionCode = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomainsDomain) SetVersionName(v string) *DescribeDomainsResponseBodyDomainsDomain {
	s.VersionName = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomainsDomain) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainsResponseBodyDomainsDomainDnsServers struct {
	DnsServer []*string `json:"DnsServer,omitempty" xml:"DnsServer,omitempty" type:"Repeated"`
}

func (s DescribeDomainsResponseBodyDomainsDomainDnsServers) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainsResponseBodyDomainsDomainDnsServers) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBodyDomainsDomainDnsServers) GetDnsServer() []*string {
	return s.DnsServer
}

func (s *DescribeDomainsResponseBodyDomainsDomainDnsServers) SetDnsServer(v []*string) *DescribeDomainsResponseBodyDomainsDomainDnsServers {
	s.DnsServer = v
	return s
}

func (s *DescribeDomainsResponseBodyDomainsDomainDnsServers) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainsResponseBodyDomainsDomainTags struct {
	Tag []*DescribeDomainsResponseBodyDomainsDomainTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDomainsResponseBodyDomainsDomainTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainsResponseBodyDomainsDomainTags) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBodyDomainsDomainTags) GetTag() []*DescribeDomainsResponseBodyDomainsDomainTagsTag {
	return s.Tag
}

func (s *DescribeDomainsResponseBodyDomainsDomainTags) SetTag(v []*DescribeDomainsResponseBodyDomainsDomainTagsTag) *DescribeDomainsResponseBodyDomainsDomainTags {
	s.Tag = v
	return s
}

func (s *DescribeDomainsResponseBodyDomainsDomainTags) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainsResponseBodyDomainsDomainTagsTag struct {
	// The key of tag N added to the resource.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N added to the resource.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDomainsResponseBodyDomainsDomainTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainsResponseBodyDomainsDomainTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBodyDomainsDomainTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeDomainsResponseBodyDomainsDomainTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeDomainsResponseBodyDomainsDomainTagsTag) SetKey(v string) *DescribeDomainsResponseBodyDomainsDomainTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomainsDomainTagsTag) SetValue(v string) *DescribeDomainsResponseBodyDomainsDomainTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomainsDomainTagsTag) Validate() error {
	return dara.Validate(s)
}
