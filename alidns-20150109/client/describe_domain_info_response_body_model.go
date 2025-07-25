// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAliDomain(v bool) *DescribeDomainInfoResponseBody
	GetAliDomain() *bool
	SetAvailableTtls(v *DescribeDomainInfoResponseBodyAvailableTtls) *DescribeDomainInfoResponseBody
	GetAvailableTtls() *DescribeDomainInfoResponseBodyAvailableTtls
	SetCreateTime(v string) *DescribeDomainInfoResponseBody
	GetCreateTime() *string
	SetDnsServers(v *DescribeDomainInfoResponseBodyDnsServers) *DescribeDomainInfoResponseBody
	GetDnsServers() *DescribeDomainInfoResponseBodyDnsServers
	SetDomainId(v string) *DescribeDomainInfoResponseBody
	GetDomainId() *string
	SetDomainLoggingSwitchStatus(v string) *DescribeDomainInfoResponseBody
	GetDomainLoggingSwitchStatus() *string
	SetDomainName(v string) *DescribeDomainInfoResponseBody
	GetDomainName() *string
	SetGroupId(v string) *DescribeDomainInfoResponseBody
	GetGroupId() *string
	SetGroupName(v string) *DescribeDomainInfoResponseBody
	GetGroupName() *string
	SetInBlackHole(v bool) *DescribeDomainInfoResponseBody
	GetInBlackHole() *bool
	SetInClean(v bool) *DescribeDomainInfoResponseBody
	GetInClean() *bool
	SetInstanceId(v string) *DescribeDomainInfoResponseBody
	GetInstanceId() *string
	SetLineType(v string) *DescribeDomainInfoResponseBody
	GetLineType() *string
	SetMinTtl(v int64) *DescribeDomainInfoResponseBody
	GetMinTtl() *int64
	SetPunyCode(v string) *DescribeDomainInfoResponseBody
	GetPunyCode() *string
	SetRecordLineTreeJson(v string) *DescribeDomainInfoResponseBody
	GetRecordLineTreeJson() *string
	SetRecordLines(v *DescribeDomainInfoResponseBodyRecordLines) *DescribeDomainInfoResponseBody
	GetRecordLines() *DescribeDomainInfoResponseBodyRecordLines
	SetRegionLines(v bool) *DescribeDomainInfoResponseBody
	GetRegionLines() *bool
	SetRemark(v string) *DescribeDomainInfoResponseBody
	GetRemark() *string
	SetRequestId(v string) *DescribeDomainInfoResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *DescribeDomainInfoResponseBody
	GetResourceGroupId() *string
	SetSlaveDns(v bool) *DescribeDomainInfoResponseBody
	GetSlaveDns() *bool
	SetSubDomain(v bool) *DescribeDomainInfoResponseBody
	GetSubDomain() *bool
	SetVersionCode(v string) *DescribeDomainInfoResponseBody
	GetVersionCode() *string
	SetVersionName(v string) *DescribeDomainInfoResponseBody
	GetVersionName() *string
}

type DescribeDomainInfoResponseBody struct {
	// Indicates whether the domain name was registered in Alibaba Cloud.
	//
	// example:
	//
	// true
	AliDomain *bool `json:"AliDomain,omitempty" xml:"AliDomain,omitempty"`
	// The available time to live (TTL) values that can be configured for the domain name. Available TTL values are not returned by default. If you want to query such information, set NeedDetailAttributes to true.
	AvailableTtls *DescribeDomainInfoResponseBodyAvailableTtls `json:"AvailableTtls,omitempty" xml:"AvailableTtls,omitempty" type:"Struct"`
	// The time when the domain name was created.
	//
	// example:
	//
	// 2015-12-12T09:23Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The DNS servers that are used to resolve the domain name.
	DnsServers *DescribeDomainInfoResponseBodyDnsServers `json:"DnsServers,omitempty" xml:"DnsServers,omitempty" type:"Struct"`
	// The ID of the domain name.
	//
	// example:
	//
	// 00efd71a-770e-4255-b54e-6fe5659baffe
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// Indicates whether the DNS traffic analysis feature is enabled. Valid values:
	DomainLoggingSwitchStatus *string `json:"DomainLoggingSwitchStatus,omitempty" xml:"DomainLoggingSwitchStatus,omitempty"`
	// The domain name.
	//
	// example:
	//
	// example.com
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
	// Indicates whether blackhole filtering was triggered.
	//
	// example:
	//
	// false
	InBlackHole *bool `json:"InBlackHole,omitempty" xml:"InBlackHole,omitempty"`
	// Indicates whether traffic scrubbing was in progress.
	//
	// example:
	//
	// false
	InClean *bool `json:"InClean,omitempty" xml:"InClean,omitempty"`
	// The ID of the Alibaba Cloud DNS instance.
	//
	// example:
	//
	// i-7bg
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the DNS request line.
	//
	// example:
	//
	// region_province
	LineType *string `json:"LineType,omitempty" xml:"LineType,omitempty"`
	// The minimum TTL value.
	//
	// example:
	//
	// 1
	MinTtl *int64 `json:"MinTtl,omitempty" xml:"MinTtl,omitempty"`
	// The Punycode for the domain name. This parameter is returned only for Chinese domain names.
	//
	// example:
	//
	// example.com
	PunyCode *string `json:"PunyCode,omitempty" xml:"PunyCode,omitempty"`
	// The tree-structure DNS request lines.
	//
	// example:
	//
	// {"default":{},"unicom":{},"telecom":{},"mobile":{},"edu":{},"oversea":{},"baidu":{},"biying":{},"google":{}}
	RecordLineTreeJson *string `json:"RecordLineTreeJson,omitempty" xml:"RecordLineTreeJson,omitempty"`
	// The DNS request lines.
	RecordLines *DescribeDomainInfoResponseBodyRecordLines `json:"RecordLines,omitempty" xml:"RecordLines,omitempty" type:"Struct"`
	// Indicates whether the DNS request lines are regional lines.
	//
	// example:
	//
	// false
	RegionLines *bool `json:"RegionLines,omitempty" xml:"RegionLines,omitempty"`
	// The description.
	//
	// example:
	//
	// remark
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek3dj3wvclgcxo
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Indicates whether secondary DNS is supported.
	//
	// example:
	//
	// true
	SlaveDns *bool `json:"SlaveDns,omitempty" xml:"SlaveDns,omitempty"`
	// Indicates whether the queried domain name is a hosted subdomain name. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	SubDomain *bool `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
	// The version ID of Alibaba Cloud DNS.
	//
	// example:
	//
	// version1
	VersionCode *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	// The edition of Alibaba Cloud DNS.
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s DescribeDomainInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainInfoResponseBody) GetAliDomain() *bool {
	return s.AliDomain
}

func (s *DescribeDomainInfoResponseBody) GetAvailableTtls() *DescribeDomainInfoResponseBodyAvailableTtls {
	return s.AvailableTtls
}

func (s *DescribeDomainInfoResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDomainInfoResponseBody) GetDnsServers() *DescribeDomainInfoResponseBodyDnsServers {
	return s.DnsServers
}

func (s *DescribeDomainInfoResponseBody) GetDomainId() *string {
	return s.DomainId
}

func (s *DescribeDomainInfoResponseBody) GetDomainLoggingSwitchStatus() *string {
	return s.DomainLoggingSwitchStatus
}

func (s *DescribeDomainInfoResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainInfoResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeDomainInfoResponseBody) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeDomainInfoResponseBody) GetInBlackHole() *bool {
	return s.InBlackHole
}

func (s *DescribeDomainInfoResponseBody) GetInClean() *bool {
	return s.InClean
}

func (s *DescribeDomainInfoResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDomainInfoResponseBody) GetLineType() *string {
	return s.LineType
}

func (s *DescribeDomainInfoResponseBody) GetMinTtl() *int64 {
	return s.MinTtl
}

func (s *DescribeDomainInfoResponseBody) GetPunyCode() *string {
	return s.PunyCode
}

func (s *DescribeDomainInfoResponseBody) GetRecordLineTreeJson() *string {
	return s.RecordLineTreeJson
}

func (s *DescribeDomainInfoResponseBody) GetRecordLines() *DescribeDomainInfoResponseBodyRecordLines {
	return s.RecordLines
}

func (s *DescribeDomainInfoResponseBody) GetRegionLines() *bool {
	return s.RegionLines
}

func (s *DescribeDomainInfoResponseBody) GetRemark() *string {
	return s.Remark
}

func (s *DescribeDomainInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainInfoResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDomainInfoResponseBody) GetSlaveDns() *bool {
	return s.SlaveDns
}

func (s *DescribeDomainInfoResponseBody) GetSubDomain() *bool {
	return s.SubDomain
}

func (s *DescribeDomainInfoResponseBody) GetVersionCode() *string {
	return s.VersionCode
}

func (s *DescribeDomainInfoResponseBody) GetVersionName() *string {
	return s.VersionName
}

func (s *DescribeDomainInfoResponseBody) SetAliDomain(v bool) *DescribeDomainInfoResponseBody {
	s.AliDomain = &v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetAvailableTtls(v *DescribeDomainInfoResponseBodyAvailableTtls) *DescribeDomainInfoResponseBody {
	s.AvailableTtls = v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetCreateTime(v string) *DescribeDomainInfoResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetDnsServers(v *DescribeDomainInfoResponseBodyDnsServers) *DescribeDomainInfoResponseBody {
	s.DnsServers = v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetDomainId(v string) *DescribeDomainInfoResponseBody {
	s.DomainId = &v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetDomainLoggingSwitchStatus(v string) *DescribeDomainInfoResponseBody {
	s.DomainLoggingSwitchStatus = &v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetDomainName(v string) *DescribeDomainInfoResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetGroupId(v string) *DescribeDomainInfoResponseBody {
	s.GroupId = &v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetGroupName(v string) *DescribeDomainInfoResponseBody {
	s.GroupName = &v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetInBlackHole(v bool) *DescribeDomainInfoResponseBody {
	s.InBlackHole = &v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetInClean(v bool) *DescribeDomainInfoResponseBody {
	s.InClean = &v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetInstanceId(v string) *DescribeDomainInfoResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetLineType(v string) *DescribeDomainInfoResponseBody {
	s.LineType = &v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetMinTtl(v int64) *DescribeDomainInfoResponseBody {
	s.MinTtl = &v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetPunyCode(v string) *DescribeDomainInfoResponseBody {
	s.PunyCode = &v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetRecordLineTreeJson(v string) *DescribeDomainInfoResponseBody {
	s.RecordLineTreeJson = &v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetRecordLines(v *DescribeDomainInfoResponseBodyRecordLines) *DescribeDomainInfoResponseBody {
	s.RecordLines = v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetRegionLines(v bool) *DescribeDomainInfoResponseBody {
	s.RegionLines = &v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetRemark(v string) *DescribeDomainInfoResponseBody {
	s.Remark = &v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetRequestId(v string) *DescribeDomainInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetResourceGroupId(v string) *DescribeDomainInfoResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetSlaveDns(v bool) *DescribeDomainInfoResponseBody {
	s.SlaveDns = &v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetSubDomain(v bool) *DescribeDomainInfoResponseBody {
	s.SubDomain = &v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetVersionCode(v string) *DescribeDomainInfoResponseBody {
	s.VersionCode = &v
	return s
}

func (s *DescribeDomainInfoResponseBody) SetVersionName(v string) *DescribeDomainInfoResponseBody {
	s.VersionName = &v
	return s
}

func (s *DescribeDomainInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainInfoResponseBodyAvailableTtls struct {
	AvailableTtl []*string `json:"AvailableTtl,omitempty" xml:"AvailableTtl,omitempty" type:"Repeated"`
}

func (s DescribeDomainInfoResponseBodyAvailableTtls) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainInfoResponseBodyAvailableTtls) GoString() string {
	return s.String()
}

func (s *DescribeDomainInfoResponseBodyAvailableTtls) GetAvailableTtl() []*string {
	return s.AvailableTtl
}

func (s *DescribeDomainInfoResponseBodyAvailableTtls) SetAvailableTtl(v []*string) *DescribeDomainInfoResponseBodyAvailableTtls {
	s.AvailableTtl = v
	return s
}

func (s *DescribeDomainInfoResponseBodyAvailableTtls) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainInfoResponseBodyDnsServers struct {
	DnsServer []*string `json:"DnsServer,omitempty" xml:"DnsServer,omitempty" type:"Repeated"`
}

func (s DescribeDomainInfoResponseBodyDnsServers) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainInfoResponseBodyDnsServers) GoString() string {
	return s.String()
}

func (s *DescribeDomainInfoResponseBodyDnsServers) GetDnsServer() []*string {
	return s.DnsServer
}

func (s *DescribeDomainInfoResponseBodyDnsServers) SetDnsServer(v []*string) *DescribeDomainInfoResponseBodyDnsServers {
	s.DnsServer = v
	return s
}

func (s *DescribeDomainInfoResponseBodyDnsServers) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainInfoResponseBodyRecordLines struct {
	RecordLine []*DescribeDomainInfoResponseBodyRecordLinesRecordLine `json:"RecordLine,omitempty" xml:"RecordLine,omitempty" type:"Repeated"`
}

func (s DescribeDomainInfoResponseBodyRecordLines) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainInfoResponseBodyRecordLines) GoString() string {
	return s.String()
}

func (s *DescribeDomainInfoResponseBodyRecordLines) GetRecordLine() []*DescribeDomainInfoResponseBodyRecordLinesRecordLine {
	return s.RecordLine
}

func (s *DescribeDomainInfoResponseBodyRecordLines) SetRecordLine(v []*DescribeDomainInfoResponseBodyRecordLinesRecordLine) *DescribeDomainInfoResponseBodyRecordLines {
	s.RecordLine = v
	return s
}

func (s *DescribeDomainInfoResponseBodyRecordLines) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainInfoResponseBodyRecordLinesRecordLine struct {
	// The code of the parent line. This parameter is not returned if the line has no parent line.
	//
	// example:
	//
	// internal
	FatherCode *string `json:"FatherCode,omitempty" xml:"FatherCode,omitempty"`
	// The code of the line.
	//
	// example:
	//
	// cn_region_xibei
	LineCode *string `json:"LineCode,omitempty" xml:"LineCode,omitempty"`
	// The name of the parent line.
	LineDisplayName *string `json:"LineDisplayName,omitempty" xml:"LineDisplayName,omitempty"`
	// The name of the line.
	LineName *string `json:"LineName,omitempty" xml:"LineName,omitempty"`
}

func (s DescribeDomainInfoResponseBodyRecordLinesRecordLine) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainInfoResponseBodyRecordLinesRecordLine) GoString() string {
	return s.String()
}

func (s *DescribeDomainInfoResponseBodyRecordLinesRecordLine) GetFatherCode() *string {
	return s.FatherCode
}

func (s *DescribeDomainInfoResponseBodyRecordLinesRecordLine) GetLineCode() *string {
	return s.LineCode
}

func (s *DescribeDomainInfoResponseBodyRecordLinesRecordLine) GetLineDisplayName() *string {
	return s.LineDisplayName
}

func (s *DescribeDomainInfoResponseBodyRecordLinesRecordLine) GetLineName() *string {
	return s.LineName
}

func (s *DescribeDomainInfoResponseBodyRecordLinesRecordLine) SetFatherCode(v string) *DescribeDomainInfoResponseBodyRecordLinesRecordLine {
	s.FatherCode = &v
	return s
}

func (s *DescribeDomainInfoResponseBodyRecordLinesRecordLine) SetLineCode(v string) *DescribeDomainInfoResponseBodyRecordLinesRecordLine {
	s.LineCode = &v
	return s
}

func (s *DescribeDomainInfoResponseBodyRecordLinesRecordLine) SetLineDisplayName(v string) *DescribeDomainInfoResponseBodyRecordLinesRecordLine {
	s.LineDisplayName = &v
	return s
}

func (s *DescribeDomainInfoResponseBodyRecordLinesRecordLine) SetLineName(v string) *DescribeDomainInfoResponseBodyRecordLinesRecordLine {
	s.LineName = &v
	return s
}

func (s *DescribeDomainInfoResponseBodyRecordLinesRecordLine) Validate() error {
	return dara.Validate(s)
}
