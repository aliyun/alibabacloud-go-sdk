// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnUserDomainsByFuncResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomains(v *DescribeCdnUserDomainsByFuncResponseBodyDomains) *DescribeCdnUserDomainsByFuncResponseBody
	GetDomains() *DescribeCdnUserDomainsByFuncResponseBodyDomains
	SetPageNumber(v int64) *DescribeCdnUserDomainsByFuncResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeCdnUserDomainsByFuncResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeCdnUserDomainsByFuncResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeCdnUserDomainsByFuncResponseBody
	GetTotalCount() *int64
}

type DescribeCdnUserDomainsByFuncResponseBody struct {
	// The configurations of the accelerated domain name.
	Domains *DescribeCdnUserDomainsByFuncResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Struct"`
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
	// 5
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// AA75AADB-5E25-4970-B480-EAA1F5658483
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 16
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCdnUserDomainsByFuncResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnUserDomainsByFuncResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserDomainsByFuncResponseBody) GetDomains() *DescribeCdnUserDomainsByFuncResponseBodyDomains {
	return s.Domains
}

func (s *DescribeCdnUserDomainsByFuncResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeCdnUserDomainsByFuncResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeCdnUserDomainsByFuncResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdnUserDomainsByFuncResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeCdnUserDomainsByFuncResponseBody) SetDomains(v *DescribeCdnUserDomainsByFuncResponseBodyDomains) *DescribeCdnUserDomainsByFuncResponseBody {
	s.Domains = v
	return s
}

func (s *DescribeCdnUserDomainsByFuncResponseBody) SetPageNumber(v int64) *DescribeCdnUserDomainsByFuncResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCdnUserDomainsByFuncResponseBody) SetPageSize(v int64) *DescribeCdnUserDomainsByFuncResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCdnUserDomainsByFuncResponseBody) SetRequestId(v string) *DescribeCdnUserDomainsByFuncResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnUserDomainsByFuncResponseBody) SetTotalCount(v int64) *DescribeCdnUserDomainsByFuncResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCdnUserDomainsByFuncResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCdnUserDomainsByFuncResponseBodyDomains struct {
	PageData []*DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
}

func (s DescribeCdnUserDomainsByFuncResponseBodyDomains) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnUserDomainsByFuncResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomains) GetPageData() []*DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData {
	return s.PageData
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomains) SetPageData(v []*DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData) *DescribeCdnUserDomainsByFuncResponseBodyDomains {
	s.PageData = v
	return s
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomains) Validate() error {
	return dara.Validate(s)
}

type DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData struct {
	// The type of workload accelerated by Alibaba Cloud CDN. Valid values:
	//
	// 	- **web**: image and small file distribution
	//
	// 	- **download**: large file distribution
	//
	// 	- **video**: on-demand video and audio streaming
	//
	// 	- **liveStream**: live streaming
	//
	// example:
	//
	// web
	CdnType *string `json:"CdnType,omitempty" xml:"CdnType,omitempty"`
	// The CNAME assigned to the accelerated domain name.
	//
	// example:
	//
	// example.com.w.alikunlun.com
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// The description of the status.
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
	// 	- **online**: The domain name is enabled.
	//
	// 	- **offline**: The domain is disabled.
	//
	// 	- **configuring**: The endpoint group is being configured.
	//
	// 	- **configure_failed**: The domain failed to be configured.
	//
	// 	- **checking**: The domain name is under review.
	//
	// 	- **check_failed**: The domain name failed the review.
	//
	// 	- **stopping**: The domain name is be disabled.
	//
	// 	- **deleting**: being deleted
	//
	// example:
	//
	// configure_failed
	DomainStatus *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	// The time when the accelerated domain name was added to Alibaba Cloud CDN.
	//
	// example:
	//
	// 2015-10-28T11:05:52Z
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The last time when the accelerated domain was modified.
	//
	// example:
	//
	// 2015-10-29T10:15:31Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-xxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The information about the origin server.
	Sources *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Struct"`
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

func (s DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData) GetCdnType() *string {
	return s.CdnType
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData) GetCname() *string {
	return s.Cname
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData) GetDescription() *string {
	return s.Description
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData) GetDomainStatus() *string {
	return s.DomainStatus
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData) GetSources() *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSources {
	return s.Sources
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData) GetSslProtocol() *string {
	return s.SslProtocol
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData) SetCdnType(v string) *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData {
	s.CdnType = &v
	return s
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData) SetCname(v string) *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData {
	s.Cname = &v
	return s
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData) SetDescription(v string) *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData {
	s.Description = &v
	return s
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData) SetDomainName(v string) *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData {
	s.DomainName = &v
	return s
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData) SetDomainStatus(v string) *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData {
	s.DomainStatus = &v
	return s
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData) SetGmtCreated(v string) *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData {
	s.GmtCreated = &v
	return s
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData) SetGmtModified(v string) *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData {
	s.GmtModified = &v
	return s
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData) SetResourceGroupId(v string) *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData) SetSources(v *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSources) *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData {
	s.Sources = v
	return s
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData) SetSslProtocol(v string) *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData {
	s.SslProtocol = &v
	return s
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData) Validate() error {
	return dara.Validate(s)
}

type DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSources struct {
	Source []*DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Repeated"`
}

func (s DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSources) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSources) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSources) GetSource() []*DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource {
	return s.Source
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSources) SetSource(v []*DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource) *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSources {
	s.Source = v
	return s
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSources) Validate() error {
	return dara.Validate(s)
}

type DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource struct {
	// The address of the origin server.
	//
	// example:
	//
	// 1.1.1.1
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The port of the origin server.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The priority of the configuration item.
	//
	// example:
	//
	// 20
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The type of the origin server.
	//
	// example:
	//
	// ipaddr
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The weight of the origin server if multiple origin servers have been specified.
	//
	// example:
	//
	// 10
	Weight *string `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource) GetContent() *string {
	return s.Content
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource) GetPort() *int32 {
	return s.Port
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource) GetPriority() *string {
	return s.Priority
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource) GetType() *string {
	return s.Type
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource) GetWeight() *string {
	return s.Weight
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource) SetContent(v string) *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource {
	s.Content = &v
	return s
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource) SetPort(v int32) *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource {
	s.Port = &v
	return s
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource) SetPriority(v string) *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource {
	s.Priority = &v
	return s
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource) SetType(v string) *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource {
	s.Type = &v
	return s
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource) SetWeight(v string) *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource {
	s.Weight = &v
	return s
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource) Validate() error {
	return dara.Validate(s)
}
