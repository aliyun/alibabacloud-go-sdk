// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnUserDomainsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomains(v *DescribeDcdnUserDomainsResponseBodyDomains) *DescribeDcdnUserDomainsResponseBody
	GetDomains() *DescribeDcdnUserDomainsResponseBodyDomains
	SetPageNumber(v int64) *DescribeDcdnUserDomainsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeDcdnUserDomainsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeDcdnUserDomainsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeDcdnUserDomainsResponseBody
	GetTotalCount() *int64
}

type DescribeDcdnUserDomainsResponseBody struct {
	// The information about the queried domains.
	Domains *DescribeDcdnUserDomainsResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Struct"`
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
	// 12
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
	// 12
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDcdnUserDomainsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserDomainsResponseBody) GetDomains() *DescribeDcdnUserDomainsResponseBodyDomains {
	return s.Domains
}

func (s *DescribeDcdnUserDomainsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeDcdnUserDomainsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDcdnUserDomainsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnUserDomainsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDcdnUserDomainsResponseBody) SetDomains(v *DescribeDcdnUserDomainsResponseBodyDomains) *DescribeDcdnUserDomainsResponseBody {
	s.Domains = v
	return s
}

func (s *DescribeDcdnUserDomainsResponseBody) SetPageNumber(v int64) *DescribeDcdnUserDomainsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnUserDomainsResponseBody) SetPageSize(v int64) *DescribeDcdnUserDomainsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnUserDomainsResponseBody) SetRequestId(v string) *DescribeDcdnUserDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnUserDomainsResponseBody) SetTotalCount(v int64) *DescribeDcdnUserDomainsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDcdnUserDomainsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnUserDomainsResponseBodyDomains struct {
	PageData []*DescribeDcdnUserDomainsResponseBodyDomainsPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
}

func (s DescribeDcdnUserDomainsResponseBodyDomains) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserDomainsResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserDomainsResponseBodyDomains) GetPageData() []*DescribeDcdnUserDomainsResponseBodyDomainsPageData {
	return s.PageData
}

func (s *DescribeDcdnUserDomainsResponseBodyDomains) SetPageData(v []*DescribeDcdnUserDomainsResponseBodyDomainsPageData) *DescribeDcdnUserDomainsResponseBodyDomains {
	s.PageData = v
	return s
}

func (s *DescribeDcdnUserDomainsResponseBodyDomains) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnUserDomainsResponseBodyDomainsPageData struct {
	// The CNAME of the domain.
	//
	// example:
	//
	// test.aliyun.com.w.alikunlun.com
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// The reason why the domain failed the review.
	//
	// example:
	//
	// audit failed
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The domain ID.
	//
	// example:
	//
	// 11223344
	DomainId *int64 `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// The domain name.
	//
	// example:
	//
	// www.example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The domain status.
	//
	// 	- **online**: The domain is active.
	//
	// 	- **offline**: The domain is suspended.
	//
	// 	- **configuring**: The domain is being configured.
	//
	// 	- **configure_failed**: The domain failed to be configured.
	//
	// 	- **checking**: The domain is under review.
	//
	// 	- **check_failed**: The domain failed the review.
	//
	// example:
	//
	// online
	DomainStatus *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	// The computing service type. Valid values:
	//
	// 	- **routine**
	//
	// 	- **image**
	//
	// 	- **cloudFunction**
	//
	// example:
	//
	// routine
	FunctionType *string `json:"FunctionType,omitempty" xml:"FunctionType,omitempty"`
	// The time when the domain was added to DCDN.
	//
	// example:
	//
	// 2015-10-28T11:05:50Z
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The time when the domain was modified.
	//
	// example:
	//
	// 2015-10-28T09:31:59Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmv6jutt**
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Indicates whether HTTPS was enabled.
	//
	// 	- **on**: enabled.
	//
	// 	- **off**: disabled.
	//
	// example:
	//
	// on
	SSLProtocol *string `json:"SSLProtocol,omitempty" xml:"SSLProtocol,omitempty"`
	// The sandbox status.
	//
	// example:
	//
	// normal
	Sandbox *string `json:"Sandbox,omitempty" xml:"Sandbox,omitempty"`
	// The acceleration scenario. Valid values:
	//
	// 	- **apiscene**: accelerates response to API calls.
	//
	// 	- **webservicescene**: accelerates content delivery for websites.
	//
	// 	- **staticscene**: accelerates the delivery of videos, images, and text.
	//
	// 	- **If you leave this parameter empty, no scenarios are supported.
	//
	// example:
	//
	// apiscene
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The information about the origin servers.
	Sources *DescribeDcdnUserDomainsResponseBodyDomainsPageDataSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Struct"`
}

func (s DescribeDcdnUserDomainsResponseBodyDomainsPageData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserDomainsResponseBodyDomainsPageData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageData) GetCname() *string {
	return s.Cname
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageData) GetDescription() *string {
	return s.Description
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageData) GetDomainId() *int64 {
	return s.DomainId
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageData) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageData) GetDomainStatus() *string {
	return s.DomainStatus
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageData) GetFunctionType() *string {
	return s.FunctionType
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageData) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageData) GetSSLProtocol() *string {
	return s.SSLProtocol
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageData) GetSandbox() *string {
	return s.Sandbox
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageData) GetScene() *string {
	return s.Scene
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageData) GetSources() *DescribeDcdnUserDomainsResponseBodyDomainsPageDataSources {
	return s.Sources
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageData) SetCname(v string) *DescribeDcdnUserDomainsResponseBodyDomainsPageData {
	s.Cname = &v
	return s
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageData) SetDescription(v string) *DescribeDcdnUserDomainsResponseBodyDomainsPageData {
	s.Description = &v
	return s
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageData) SetDomainId(v int64) *DescribeDcdnUserDomainsResponseBodyDomainsPageData {
	s.DomainId = &v
	return s
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageData) SetDomainName(v string) *DescribeDcdnUserDomainsResponseBodyDomainsPageData {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageData) SetDomainStatus(v string) *DescribeDcdnUserDomainsResponseBodyDomainsPageData {
	s.DomainStatus = &v
	return s
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageData) SetFunctionType(v string) *DescribeDcdnUserDomainsResponseBodyDomainsPageData {
	s.FunctionType = &v
	return s
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageData) SetGmtCreated(v string) *DescribeDcdnUserDomainsResponseBodyDomainsPageData {
	s.GmtCreated = &v
	return s
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageData) SetGmtModified(v string) *DescribeDcdnUserDomainsResponseBodyDomainsPageData {
	s.GmtModified = &v
	return s
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageData) SetResourceGroupId(v string) *DescribeDcdnUserDomainsResponseBodyDomainsPageData {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageData) SetSSLProtocol(v string) *DescribeDcdnUserDomainsResponseBodyDomainsPageData {
	s.SSLProtocol = &v
	return s
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageData) SetSandbox(v string) *DescribeDcdnUserDomainsResponseBodyDomainsPageData {
	s.Sandbox = &v
	return s
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageData) SetScene(v string) *DescribeDcdnUserDomainsResponseBodyDomainsPageData {
	s.Scene = &v
	return s
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageData) SetSources(v *DescribeDcdnUserDomainsResponseBodyDomainsPageDataSources) *DescribeDcdnUserDomainsResponseBodyDomainsPageData {
	s.Sources = v
	return s
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageData) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnUserDomainsResponseBodyDomainsPageDataSources struct {
	Source []*DescribeDcdnUserDomainsResponseBodyDomainsPageDataSourcesSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Repeated"`
}

func (s DescribeDcdnUserDomainsResponseBodyDomainsPageDataSources) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserDomainsResponseBodyDomainsPageDataSources) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageDataSources) GetSource() []*DescribeDcdnUserDomainsResponseBodyDomainsPageDataSourcesSource {
	return s.Source
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageDataSources) SetSource(v []*DescribeDcdnUserDomainsResponseBodyDomainsPageDataSourcesSource) *DescribeDcdnUserDomainsResponseBodyDomainsPageDataSources {
	s.Source = v
	return s
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageDataSources) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnUserDomainsResponseBodyDomainsPageDataSourcesSource struct {
	// The address of the origin server.
	//
	// example:
	//
	// example.com
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The port of the origin server.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The priority of the origin server.
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
	// The weight of the origin server.
	//
	// example:
	//
	// 20
	Weight *string `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeDcdnUserDomainsResponseBodyDomainsPageDataSourcesSource) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserDomainsResponseBodyDomainsPageDataSourcesSource) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageDataSourcesSource) GetContent() *string {
	return s.Content
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageDataSourcesSource) GetPort() *int32 {
	return s.Port
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageDataSourcesSource) GetPriority() *string {
	return s.Priority
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageDataSourcesSource) GetType() *string {
	return s.Type
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageDataSourcesSource) GetWeight() *string {
	return s.Weight
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageDataSourcesSource) SetContent(v string) *DescribeDcdnUserDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Content = &v
	return s
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageDataSourcesSource) SetPort(v int32) *DescribeDcdnUserDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Port = &v
	return s
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageDataSourcesSource) SetPriority(v string) *DescribeDcdnUserDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Priority = &v
	return s
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageDataSourcesSource) SetType(v string) *DescribeDcdnUserDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Type = &v
	return s
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageDataSourcesSource) SetWeight(v string) *DescribeDcdnUserDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Weight = &v
	return s
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageDataSourcesSource) Validate() error {
	return dara.Validate(s)
}
