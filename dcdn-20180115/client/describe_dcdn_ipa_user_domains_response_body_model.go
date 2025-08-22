// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnIpaUserDomainsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomains(v *DescribeDcdnIpaUserDomainsResponseBodyDomains) *DescribeDcdnIpaUserDomainsResponseBody
	GetDomains() *DescribeDcdnIpaUserDomainsResponseBodyDomains
	SetPageNumber(v int64) *DescribeDcdnIpaUserDomainsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeDcdnIpaUserDomainsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeDcdnIpaUserDomainsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeDcdnIpaUserDomainsResponseBody
	GetTotalCount() *int64
}

type DescribeDcdnIpaUserDomainsResponseBody struct {
	// The array that consists of multiple PageData parameters. The details about each accelerated domain name are included in a separate PageData parameter.
	Domains *DescribeDcdnIpaUserDomainsResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of domain names returned per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// AA75AADB-5E25-4970-B480-EAA1F5658483
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of domain names returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDcdnIpaUserDomainsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnIpaUserDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaUserDomainsResponseBody) GetDomains() *DescribeDcdnIpaUserDomainsResponseBodyDomains {
	return s.Domains
}

func (s *DescribeDcdnIpaUserDomainsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeDcdnIpaUserDomainsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDcdnIpaUserDomainsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnIpaUserDomainsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDcdnIpaUserDomainsResponseBody) SetDomains(v *DescribeDcdnIpaUserDomainsResponseBodyDomains) *DescribeDcdnIpaUserDomainsResponseBody {
	s.Domains = v
	return s
}

func (s *DescribeDcdnIpaUserDomainsResponseBody) SetPageNumber(v int64) *DescribeDcdnIpaUserDomainsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsResponseBody) SetPageSize(v int64) *DescribeDcdnIpaUserDomainsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsResponseBody) SetRequestId(v string) *DescribeDcdnIpaUserDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsResponseBody) SetTotalCount(v int64) *DescribeDcdnIpaUserDomainsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnIpaUserDomainsResponseBodyDomains struct {
	PageData []*DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
}

func (s DescribeDcdnIpaUserDomainsResponseBodyDomains) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnIpaUserDomainsResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomains) GetPageData() []*DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData {
	return s.PageData
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomains) SetPageData(v []*DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData) *DescribeDcdnIpaUserDomainsResponseBodyDomains {
	s.PageData = v
	return s
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomains) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData struct {
	// The CNAME assigned to the accelerated domain name.
	//
	// example:
	//
	// example.com.*.com
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// The reason why the accelerated domain name failed the review.
	//
	// example:
	//
	// audit failed
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The accelerated domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The status of the accelerated domain name. Valid values:
	//
	// 	- **online**: enabled
	//
	// 	- **offline**: disabled
	//
	// 	- **configuring**: configuring
	//
	// 	- **configure_failed**: configuration failed
	//
	// 	- **checking**: reviewing
	//
	// 	- **check_failed:*	- review failed
	//
	// example:
	//
	// online
	DomainStatus *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	// The time when the accelerated domain name was added to Alibaba Cloud CDN.
	//
	// example:
	//
	// 2015-10-28T09:32:51Z
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The time when the accelerated domain name was modified.
	//
	// example:
	//
	// 2015-10-28T11:05:52Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// abcd1234abcd****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of HTTPS.
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// on
	SSLProtocol *string `json:"SSLProtocol,omitempty" xml:"SSLProtocol,omitempty"`
	// Indicates whether the accelerated domain name was in a sandbox.
	//
	// example:
	//
	// normal
	Sandbox *string `json:"Sandbox,omitempty" xml:"Sandbox,omitempty"`
	// The information about the origin server.
	Sources *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Struct"`
}

func (s DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData) GetCname() *string {
	return s.Cname
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData) GetDescription() *string {
	return s.Description
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData) GetDomainStatus() *string {
	return s.DomainStatus
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData) GetSSLProtocol() *string {
	return s.SSLProtocol
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData) GetSandbox() *string {
	return s.Sandbox
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData) GetSources() *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSources {
	return s.Sources
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData) SetCname(v string) *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData {
	s.Cname = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData) SetDescription(v string) *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData {
	s.Description = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData) SetDomainName(v string) *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData) SetDomainStatus(v string) *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData {
	s.DomainStatus = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData) SetGmtCreated(v string) *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData {
	s.GmtCreated = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData) SetGmtModified(v string) *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData {
	s.GmtModified = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData) SetResourceGroupId(v string) *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData) SetSSLProtocol(v string) *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData {
	s.SSLProtocol = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData) SetSandbox(v string) *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData {
	s.Sandbox = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData) SetSources(v *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSources) *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData {
	s.Sources = v
	return s
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSources struct {
	Source []*DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSourcesSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Repeated"`
}

func (s DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSources) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSources) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSources) GetSource() []*DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSourcesSource {
	return s.Source
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSources) SetSource(v []*DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSourcesSource) *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSources {
	s.Source = v
	return s
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSources) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSourcesSource struct {
	// The address of the origin server.
	//
	// example:
	//
	// *.aliyuncs.com
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The port of the origin server.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The priority.
	//
	// example:
	//
	// 20
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The type of the origin server.
	//
	// example:
	//
	// OSS Domain
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The weight of the origin server if multiple origin servers have been specified.
	//
	// example:
	//
	// 20
	Weight *string `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSourcesSource) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSourcesSource) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSourcesSource) GetContent() *string {
	return s.Content
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSourcesSource) GetPort() *int32 {
	return s.Port
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSourcesSource) GetPriority() *string {
	return s.Priority
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSourcesSource) GetType() *string {
	return s.Type
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSourcesSource) GetWeight() *string {
	return s.Weight
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSourcesSource) SetContent(v string) *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Content = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSourcesSource) SetPort(v int32) *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Port = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSourcesSource) SetPriority(v string) *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Priority = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSourcesSource) SetType(v string) *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Type = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSourcesSource) SetWeight(v string) *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Weight = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSourcesSource) Validate() error {
	return dara.Validate(s)
}
