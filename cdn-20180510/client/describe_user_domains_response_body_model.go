// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserDomainsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomains(v *DescribeUserDomainsResponseBodyDomains) *DescribeUserDomainsResponseBody
	GetDomains() *DescribeUserDomainsResponseBodyDomains
	SetPageNumber(v int64) *DescribeUserDomainsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeUserDomainsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeUserDomainsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeUserDomainsResponseBody
	GetTotalCount() *int64
}

type DescribeUserDomainsResponseBody struct {
	// The list of the accelerated domain names returned.
	Domains *DescribeUserDomainsResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Struct"`
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
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// BDA62CE4-3477-439A-B52E-D2D7C829D7C1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeUserDomainsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserDomainsResponseBody) GetDomains() *DescribeUserDomainsResponseBodyDomains {
	return s.Domains
}

func (s *DescribeUserDomainsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeUserDomainsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeUserDomainsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserDomainsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeUserDomainsResponseBody) SetDomains(v *DescribeUserDomainsResponseBodyDomains) *DescribeUserDomainsResponseBody {
	s.Domains = v
	return s
}

func (s *DescribeUserDomainsResponseBody) SetPageNumber(v int64) *DescribeUserDomainsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeUserDomainsResponseBody) SetPageSize(v int64) *DescribeUserDomainsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeUserDomainsResponseBody) SetRequestId(v string) *DescribeUserDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserDomainsResponseBody) SetTotalCount(v int64) *DescribeUserDomainsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeUserDomainsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeUserDomainsResponseBodyDomains struct {
	PageData []*DescribeUserDomainsResponseBodyDomainsPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
}

func (s DescribeUserDomainsResponseBodyDomains) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserDomainsResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *DescribeUserDomainsResponseBodyDomains) GetPageData() []*DescribeUserDomainsResponseBodyDomainsPageData {
	return s.PageData
}

func (s *DescribeUserDomainsResponseBodyDomains) SetPageData(v []*DescribeUserDomainsResponseBodyDomainsPageData) *DescribeUserDomainsResponseBodyDomains {
	s.PageData = v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomains) Validate() error {
	return dara.Validate(s)
}

type DescribeUserDomainsResponseBodyDomainsPageData struct {
	// The type of the workload accelerated by Alibaba Cloud CDN. Valid values:
	//
	// 	- **web**: images and small files
	//
	// 	- **download**: large files
	//
	// 	- **video**: on-demand video and audio streaming
	//
	// example:
	//
	// download
	CdnType *string `json:"CdnType,omitempty" xml:"CdnType,omitempty"`
	// The CNAME assigned to the accelerated domain name.
	//
	// example:
	//
	// example.com.w.alikunlun.net
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// The acceleration region. Valid values:
	//
	// 	- **domestic**: Chinese mainland
	//
	// 	- **global**: global
	//
	// 	- **overseas**: outside the Chinese mainland
	//
	// example:
	//
	// domestic
	Coverage *string `json:"Coverage,omitempty" xml:"Coverage,omitempty"`
	// The information about Internet Content Provider (ICP) filing.
	//
	// example:
	//
	// filing description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the accelerated domain name.
	//
	// example:
	//
	// 11223344
	DomainId *int64 `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// The accelerated domain.
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
	// 	- **stopping**
	//
	// 	- **deleting**
	//
	// example:
	//
	// configure_failed
	DomainStatus *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	// The time when the accelerated domain name was added.
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
	// abcd1234abcd1234
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Indicates whether the accelerated domain name is in a sandbox.
	//
	// example:
	//
	// true
	Sandbox *string `json:"Sandbox,omitempty" xml:"Sandbox,omitempty"`
	// The information about the origin server.
	Sources *DescribeUserDomainsResponseBodyDomainsPageDataSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Struct"`
	// Indicates whether HTTPS is enabled. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// on
	SslProtocol *string `json:"SslProtocol,omitempty" xml:"SslProtocol,omitempty"`
}

func (s DescribeUserDomainsResponseBodyDomainsPageData) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserDomainsResponseBodyDomainsPageData) GoString() string {
	return s.String()
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) GetCdnType() *string {
	return s.CdnType
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) GetCname() *string {
	return s.Cname
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) GetCoverage() *string {
	return s.Coverage
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) GetDescription() *string {
	return s.Description
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) GetDomainId() *int64 {
	return s.DomainId
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) GetDomainStatus() *string {
	return s.DomainStatus
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) GetSandbox() *string {
	return s.Sandbox
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) GetSources() *DescribeUserDomainsResponseBodyDomainsPageDataSources {
	return s.Sources
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) GetSslProtocol() *string {
	return s.SslProtocol
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) SetCdnType(v string) *DescribeUserDomainsResponseBodyDomainsPageData {
	s.CdnType = &v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) SetCname(v string) *DescribeUserDomainsResponseBodyDomainsPageData {
	s.Cname = &v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) SetCoverage(v string) *DescribeUserDomainsResponseBodyDomainsPageData {
	s.Coverage = &v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) SetDescription(v string) *DescribeUserDomainsResponseBodyDomainsPageData {
	s.Description = &v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) SetDomainId(v int64) *DescribeUserDomainsResponseBodyDomainsPageData {
	s.DomainId = &v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) SetDomainName(v string) *DescribeUserDomainsResponseBodyDomainsPageData {
	s.DomainName = &v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) SetDomainStatus(v string) *DescribeUserDomainsResponseBodyDomainsPageData {
	s.DomainStatus = &v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) SetGmtCreated(v string) *DescribeUserDomainsResponseBodyDomainsPageData {
	s.GmtCreated = &v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) SetGmtModified(v string) *DescribeUserDomainsResponseBodyDomainsPageData {
	s.GmtModified = &v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) SetResourceGroupId(v string) *DescribeUserDomainsResponseBodyDomainsPageData {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) SetSandbox(v string) *DescribeUserDomainsResponseBodyDomainsPageData {
	s.Sandbox = &v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) SetSources(v *DescribeUserDomainsResponseBodyDomainsPageDataSources) *DescribeUserDomainsResponseBodyDomainsPageData {
	s.Sources = v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) SetSslProtocol(v string) *DescribeUserDomainsResponseBodyDomainsPageData {
	s.SslProtocol = &v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) Validate() error {
	return dara.Validate(s)
}

type DescribeUserDomainsResponseBodyDomainsPageDataSources struct {
	Source []*DescribeUserDomainsResponseBodyDomainsPageDataSourcesSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Repeated"`
}

func (s DescribeUserDomainsResponseBodyDomainsPageDataSources) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserDomainsResponseBodyDomainsPageDataSources) GoString() string {
	return s.String()
}

func (s *DescribeUserDomainsResponseBodyDomainsPageDataSources) GetSource() []*DescribeUserDomainsResponseBodyDomainsPageDataSourcesSource {
	return s.Source
}

func (s *DescribeUserDomainsResponseBodyDomainsPageDataSources) SetSource(v []*DescribeUserDomainsResponseBodyDomainsPageDataSourcesSource) *DescribeUserDomainsResponseBodyDomainsPageDataSources {
	s.Source = v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomainsPageDataSources) Validate() error {
	return dara.Validate(s)
}

type DescribeUserDomainsResponseBodyDomainsPageDataSourcesSource struct {
	// The address of the origin server.
	//
	// example:
	//
	// ***.oss-cn-hangzhou.aliyuncs.com
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
	// 15
	Weight *string `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeUserDomainsResponseBodyDomainsPageDataSourcesSource) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserDomainsResponseBodyDomainsPageDataSourcesSource) GoString() string {
	return s.String()
}

func (s *DescribeUserDomainsResponseBodyDomainsPageDataSourcesSource) GetContent() *string {
	return s.Content
}

func (s *DescribeUserDomainsResponseBodyDomainsPageDataSourcesSource) GetPort() *int32 {
	return s.Port
}

func (s *DescribeUserDomainsResponseBodyDomainsPageDataSourcesSource) GetPriority() *string {
	return s.Priority
}

func (s *DescribeUserDomainsResponseBodyDomainsPageDataSourcesSource) GetType() *string {
	return s.Type
}

func (s *DescribeUserDomainsResponseBodyDomainsPageDataSourcesSource) GetWeight() *string {
	return s.Weight
}

func (s *DescribeUserDomainsResponseBodyDomainsPageDataSourcesSource) SetContent(v string) *DescribeUserDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Content = &v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomainsPageDataSourcesSource) SetPort(v int32) *DescribeUserDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Port = &v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomainsPageDataSourcesSource) SetPriority(v string) *DescribeUserDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Priority = &v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomainsPageDataSourcesSource) SetType(v string) *DescribeUserDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Type = &v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomainsPageDataSourcesSource) SetWeight(v string) *DescribeUserDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Weight = &v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomainsPageDataSourcesSource) Validate() error {
	return dara.Validate(s)
}
