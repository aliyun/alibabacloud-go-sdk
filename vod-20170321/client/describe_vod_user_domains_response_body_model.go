// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodUserDomainsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomains(v *DescribeVodUserDomainsResponseBodyDomains) *DescribeVodUserDomainsResponseBody
	GetDomains() *DescribeVodUserDomainsResponseBodyDomains
	SetPageNumber(v int64) *DescribeVodUserDomainsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeVodUserDomainsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeVodUserDomainsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeVodUserDomainsResponseBody
	GetTotalCount() *int64
}

type DescribeVodUserDomainsResponseBody struct {
	// The detailed information about each domain name for CDN. The returned information is displayed in the format that is specified by the PageData parameter.
	Domains *DescribeVodUserDomainsResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E4EBD2BF-5EB0-4476-****-9D94E1B15267
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeVodUserDomainsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodUserDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodUserDomainsResponseBody) GetDomains() *DescribeVodUserDomainsResponseBodyDomains {
	return s.Domains
}

func (s *DescribeVodUserDomainsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeVodUserDomainsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeVodUserDomainsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodUserDomainsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeVodUserDomainsResponseBody) SetDomains(v *DescribeVodUserDomainsResponseBodyDomains) *DescribeVodUserDomainsResponseBody {
	s.Domains = v
	return s
}

func (s *DescribeVodUserDomainsResponseBody) SetPageNumber(v int64) *DescribeVodUserDomainsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeVodUserDomainsResponseBody) SetPageSize(v int64) *DescribeVodUserDomainsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVodUserDomainsResponseBody) SetRequestId(v string) *DescribeVodUserDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodUserDomainsResponseBody) SetTotalCount(v int64) *DescribeVodUserDomainsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVodUserDomainsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVodUserDomainsResponseBodyDomains struct {
	PageData []*DescribeVodUserDomainsResponseBodyDomainsPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
}

func (s DescribeVodUserDomainsResponseBodyDomains) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodUserDomainsResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *DescribeVodUserDomainsResponseBodyDomains) GetPageData() []*DescribeVodUserDomainsResponseBodyDomainsPageData {
	return s.PageData
}

func (s *DescribeVodUserDomainsResponseBodyDomains) SetPageData(v []*DescribeVodUserDomainsResponseBodyDomainsPageData) *DescribeVodUserDomainsResponseBodyDomains {
	s.PageData = v
	return s
}

func (s *DescribeVodUserDomainsResponseBodyDomains) Validate() error {
	return dara.Validate(s)
}

type DescribeVodUserDomainsResponseBodyDomainsPageData struct {
	// The CNAME that is assigned to the domain name for CDN.
	//
	// example:
	//
	// learn.developer.aliyundoc.com
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// The remarks.
	//
	// example:
	//
	// Zhejiang ICP Filing No. ****
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The domain name for CDN.
	//
	// example:
	//
	// example.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The status of the domain name for CDN. Valid values:
	//
	// 	- **online**: indicates that the domain name is enabled.
	//
	// 	- **offline**: indicates that the domain name is disabled.
	//
	// 	- **configuring**: indicates that the domain name is being configured.
	//
	// 	- **configure_failed**: indicates that the domain name failed to be configured.
	//
	// 	- **checking**: indicates that the domain name is under review.
	//
	// 	- **check_failed**: indicates that the domain name failed the review.
	//
	// example:
	//
	// online
	DomainStatus *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	// The time when the domain name for CDN was added. The time is in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-08-29T08:40:53Z
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The last time when the domain name for CDN was modified. The time is in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-12-29T09:24:12Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Indicates whether the accelerated domain name was in a sandbox.
	//
	// example:
	//
	// normal
	Sandbox *string `json:"Sandbox,omitempty" xml:"Sandbox,omitempty"`
	// The information about the origin server.
	Sources *DescribeVodUserDomainsResponseBodyDomainsPageDataSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Struct"`
	// Indicates whether HTTPS is enabled.
	//
	// 	- **on**: HTTPS is enabled.
	//
	// 	- **off**: HTTPS is not eabled.
	//
	// example:
	//
	// on
	SslProtocol *string `json:"SslProtocol,omitempty" xml:"SslProtocol,omitempty"`
}

func (s DescribeVodUserDomainsResponseBodyDomainsPageData) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodUserDomainsResponseBodyDomainsPageData) GoString() string {
	return s.String()
}

func (s *DescribeVodUserDomainsResponseBodyDomainsPageData) GetCname() *string {
	return s.Cname
}

func (s *DescribeVodUserDomainsResponseBodyDomainsPageData) GetDescription() *string {
	return s.Description
}

func (s *DescribeVodUserDomainsResponseBodyDomainsPageData) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodUserDomainsResponseBodyDomainsPageData) GetDomainStatus() *string {
	return s.DomainStatus
}

func (s *DescribeVodUserDomainsResponseBodyDomainsPageData) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeVodUserDomainsResponseBodyDomainsPageData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeVodUserDomainsResponseBodyDomainsPageData) GetSandbox() *string {
	return s.Sandbox
}

func (s *DescribeVodUserDomainsResponseBodyDomainsPageData) GetSources() *DescribeVodUserDomainsResponseBodyDomainsPageDataSources {
	return s.Sources
}

func (s *DescribeVodUserDomainsResponseBodyDomainsPageData) GetSslProtocol() *string {
	return s.SslProtocol
}

func (s *DescribeVodUserDomainsResponseBodyDomainsPageData) SetCname(v string) *DescribeVodUserDomainsResponseBodyDomainsPageData {
	s.Cname = &v
	return s
}

func (s *DescribeVodUserDomainsResponseBodyDomainsPageData) SetDescription(v string) *DescribeVodUserDomainsResponseBodyDomainsPageData {
	s.Description = &v
	return s
}

func (s *DescribeVodUserDomainsResponseBodyDomainsPageData) SetDomainName(v string) *DescribeVodUserDomainsResponseBodyDomainsPageData {
	s.DomainName = &v
	return s
}

func (s *DescribeVodUserDomainsResponseBodyDomainsPageData) SetDomainStatus(v string) *DescribeVodUserDomainsResponseBodyDomainsPageData {
	s.DomainStatus = &v
	return s
}

func (s *DescribeVodUserDomainsResponseBodyDomainsPageData) SetGmtCreated(v string) *DescribeVodUserDomainsResponseBodyDomainsPageData {
	s.GmtCreated = &v
	return s
}

func (s *DescribeVodUserDomainsResponseBodyDomainsPageData) SetGmtModified(v string) *DescribeVodUserDomainsResponseBodyDomainsPageData {
	s.GmtModified = &v
	return s
}

func (s *DescribeVodUserDomainsResponseBodyDomainsPageData) SetSandbox(v string) *DescribeVodUserDomainsResponseBodyDomainsPageData {
	s.Sandbox = &v
	return s
}

func (s *DescribeVodUserDomainsResponseBodyDomainsPageData) SetSources(v *DescribeVodUserDomainsResponseBodyDomainsPageDataSources) *DescribeVodUserDomainsResponseBodyDomainsPageData {
	s.Sources = v
	return s
}

func (s *DescribeVodUserDomainsResponseBodyDomainsPageData) SetSslProtocol(v string) *DescribeVodUserDomainsResponseBodyDomainsPageData {
	s.SslProtocol = &v
	return s
}

func (s *DescribeVodUserDomainsResponseBodyDomainsPageData) Validate() error {
	return dara.Validate(s)
}

type DescribeVodUserDomainsResponseBodyDomainsPageDataSources struct {
	Source []*DescribeVodUserDomainsResponseBodyDomainsPageDataSourcesSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Repeated"`
}

func (s DescribeVodUserDomainsResponseBodyDomainsPageDataSources) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodUserDomainsResponseBodyDomainsPageDataSources) GoString() string {
	return s.String()
}

func (s *DescribeVodUserDomainsResponseBodyDomainsPageDataSources) GetSource() []*DescribeVodUserDomainsResponseBodyDomainsPageDataSourcesSource {
	return s.Source
}

func (s *DescribeVodUserDomainsResponseBodyDomainsPageDataSources) SetSource(v []*DescribeVodUserDomainsResponseBodyDomainsPageDataSourcesSource) *DescribeVodUserDomainsResponseBodyDomainsPageDataSources {
	s.Source = v
	return s
}

func (s *DescribeVodUserDomainsResponseBodyDomainsPageDataSources) Validate() error {
	return dara.Validate(s)
}

type DescribeVodUserDomainsResponseBodyDomainsPageDataSourcesSource struct {
	// The address of the origin server.
	//
	// example:
	//
	// 192.168.0.1
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The port number. Valid values: **443*	- and **80**.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The priority of the origin server.
	//
	// example:
	//
	// 5
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The type of the origin server. Valid values:
	//
	// 	- **ipaddr**: an IP address.
	//
	// 	- **domain**: an origin domain name
	//
	// 	- **oss**: the OSS domain of an Object Storage Service (OSS) bucket
	//
	// example:
	//
	// oss
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeVodUserDomainsResponseBodyDomainsPageDataSourcesSource) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodUserDomainsResponseBodyDomainsPageDataSourcesSource) GoString() string {
	return s.String()
}

func (s *DescribeVodUserDomainsResponseBodyDomainsPageDataSourcesSource) GetContent() *string {
	return s.Content
}

func (s *DescribeVodUserDomainsResponseBodyDomainsPageDataSourcesSource) GetPort() *int32 {
	return s.Port
}

func (s *DescribeVodUserDomainsResponseBodyDomainsPageDataSourcesSource) GetPriority() *string {
	return s.Priority
}

func (s *DescribeVodUserDomainsResponseBodyDomainsPageDataSourcesSource) GetType() *string {
	return s.Type
}

func (s *DescribeVodUserDomainsResponseBodyDomainsPageDataSourcesSource) SetContent(v string) *DescribeVodUserDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Content = &v
	return s
}

func (s *DescribeVodUserDomainsResponseBodyDomainsPageDataSourcesSource) SetPort(v int32) *DescribeVodUserDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Port = &v
	return s
}

func (s *DescribeVodUserDomainsResponseBodyDomainsPageDataSourcesSource) SetPriority(v string) *DescribeVodUserDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Priority = &v
	return s
}

func (s *DescribeVodUserDomainsResponseBodyDomainsPageDataSourcesSource) SetType(v string) *DescribeVodUserDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Type = &v
	return s
}

func (s *DescribeVodUserDomainsResponseBodyDomainsPageDataSourcesSource) Validate() error {
	return dara.Validate(s)
}
