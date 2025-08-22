// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnUserDomainsByFuncResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomains(v *DescribeDcdnUserDomainsByFuncResponseBodyDomains) *DescribeDcdnUserDomainsByFuncResponseBody
	GetDomains() *DescribeDcdnUserDomainsByFuncResponseBodyDomains
	SetPageNumber(v int64) *DescribeDcdnUserDomainsByFuncResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeDcdnUserDomainsByFuncResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeDcdnUserDomainsByFuncResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeDcdnUserDomainsByFuncResponseBody
	GetTotalCount() *int64
}

type DescribeDcdnUserDomainsByFuncResponseBody struct {
	// The array that consists of multiple PageData parameters. The details about each accelerated domain name are included in a separate PageData parameter.
	Domains *DescribeDcdnUserDomainsByFuncResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Struct"`
	// The page number of the page returned.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 5
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AA75AADB-5E25-4970-B480-EAA1F5658483
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of domain names returned.
	//
	// example:
	//
	// 16
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDcdnUserDomainsByFuncResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserDomainsByFuncResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserDomainsByFuncResponseBody) GetDomains() *DescribeDcdnUserDomainsByFuncResponseBodyDomains {
	return s.Domains
}

func (s *DescribeDcdnUserDomainsByFuncResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeDcdnUserDomainsByFuncResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDcdnUserDomainsByFuncResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnUserDomainsByFuncResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDcdnUserDomainsByFuncResponseBody) SetDomains(v *DescribeDcdnUserDomainsByFuncResponseBodyDomains) *DescribeDcdnUserDomainsByFuncResponseBody {
	s.Domains = v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncResponseBody) SetPageNumber(v int64) *DescribeDcdnUserDomainsByFuncResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncResponseBody) SetPageSize(v int64) *DescribeDcdnUserDomainsByFuncResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncResponseBody) SetRequestId(v string) *DescribeDcdnUserDomainsByFuncResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncResponseBody) SetTotalCount(v int64) *DescribeDcdnUserDomainsByFuncResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnUserDomainsByFuncResponseBodyDomains struct {
	PageData []*DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
}

func (s DescribeDcdnUserDomainsByFuncResponseBodyDomains) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserDomainsByFuncResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomains) GetPageData() []*DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData {
	return s.PageData
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomains) SetPageData(v []*DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData) *DescribeDcdnUserDomainsByFuncResponseBodyDomains {
	s.PageData = v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomains) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData struct {
	// The CNAME assigned to the accelerated domain name.
	//
	// example:
	//
	// example.com.w.alikunlun.net
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
	// 	- **online**
	//
	// 	- **offline**
	//
	// 	- **configuring**
	//
	// 	- **configure_failed**
	//
	// 	- **checking**
	//
	// 	- **check_failed**
	//
	// example:
	//
	// configure_failed
	DomainStatus *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	// The time when the accelerated domain name was added to Dynamic Content Delivery Network (DCDN).
	//
	// example:
	//
	// 2015-10-28T09:32:51Z
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The time when the accelerated domain name was modified.
	//
	// example:
	//
	// 2015-10-28T11:05:50Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmyuji4b6r4**
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Indicates whether the accelerated domain name was in a sandbox.
	//
	// example:
	//
	// normal
	Sandbox *string `json:"Sandbox,omitempty" xml:"Sandbox,omitempty"`
	// The information about the origin servers.
	Sources *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Struct"`
	// Indicates whether HTTPS is enabled. Valid values:
	//
	// 	- **on**: HTTPS is enabled.
	//
	// 	- **off**: HTTPS is disabled.
	//
	// example:
	//
	// on
	SslProtocol *string `json:"SslProtocol,omitempty" xml:"SslProtocol,omitempty"`
}

func (s DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData) GetCname() *string {
	return s.Cname
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData) GetDescription() *string {
	return s.Description
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData) GetDomainStatus() *string {
	return s.DomainStatus
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData) GetSandbox() *string {
	return s.Sandbox
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData) GetSources() *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSources {
	return s.Sources
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData) GetSslProtocol() *string {
	return s.SslProtocol
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData) SetCname(v string) *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData {
	s.Cname = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData) SetDescription(v string) *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData {
	s.Description = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData) SetDomainName(v string) *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData) SetDomainStatus(v string) *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData {
	s.DomainStatus = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData) SetGmtCreated(v string) *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData {
	s.GmtCreated = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData) SetGmtModified(v string) *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData {
	s.GmtModified = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData) SetResourceGroupId(v string) *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData) SetSandbox(v string) *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData {
	s.Sandbox = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData) SetSources(v *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSources) *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData {
	s.Sources = v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData) SetSslProtocol(v string) *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData {
	s.SslProtocol = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSources struct {
	Source []*DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Repeated"`
}

func (s DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSources) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSources) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSources) GetSource() []*DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource {
	return s.Source
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSources) SetSource(v []*DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource) *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSources {
	s.Source = v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSources) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource struct {
	// The origin server address.
	//
	// example:
	//
	// image.developer.aliyundoc.com
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
	// oss
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The weight of the origin server if multiple origin servers have been specified.
	//
	// example:
	//
	// 20
	Weight *string `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource) GetContent() *string {
	return s.Content
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource) GetPort() *int32 {
	return s.Port
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource) GetPriority() *string {
	return s.Priority
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource) GetType() *string {
	return s.Type
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource) GetWeight() *string {
	return s.Weight
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource) SetContent(v string) *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource {
	s.Content = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource) SetPort(v int32) *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource {
	s.Port = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource) SetPriority(v string) *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource {
	s.Priority = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource) SetType(v string) *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource {
	s.Type = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource) SetWeight(v string) *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource {
	s.Weight = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource) Validate() error {
	return dara.Validate(s)
}
