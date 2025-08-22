// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRDDomainsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomains(v *DescribeRDDomainsResponseBodyDomains) *DescribeRDDomainsResponseBody
	GetDomains() *DescribeRDDomainsResponseBodyDomains
	SetPageNumber(v int64) *DescribeRDDomainsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeRDDomainsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeRDDomainsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeRDDomainsResponseBody
	GetTotalCount() *int64
}

type DescribeRDDomainsResponseBody struct {
	// The status information about the accelerated domain name.
	Domains *DescribeRDDomainsResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Struct"`
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
	// The request ID.
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

func (s DescribeRDDomainsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRDDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRDDomainsResponseBody) GetDomains() *DescribeRDDomainsResponseBodyDomains {
	return s.Domains
}

func (s *DescribeRDDomainsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeRDDomainsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeRDDomainsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRDDomainsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeRDDomainsResponseBody) SetDomains(v *DescribeRDDomainsResponseBodyDomains) *DescribeRDDomainsResponseBody {
	s.Domains = v
	return s
}

func (s *DescribeRDDomainsResponseBody) SetPageNumber(v int64) *DescribeRDDomainsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeRDDomainsResponseBody) SetPageSize(v int64) *DescribeRDDomainsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRDDomainsResponseBody) SetRequestId(v string) *DescribeRDDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRDDomainsResponseBody) SetTotalCount(v int64) *DescribeRDDomainsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRDDomainsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRDDomainsResponseBodyDomains struct {
	PageData []*DescribeRDDomainsResponseBodyDomainsPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
}

func (s DescribeRDDomainsResponseBodyDomains) String() string {
	return dara.Prettify(s)
}

func (s DescribeRDDomainsResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *DescribeRDDomainsResponseBodyDomains) GetPageData() []*DescribeRDDomainsResponseBodyDomainsPageData {
	return s.PageData
}

func (s *DescribeRDDomainsResponseBodyDomains) SetPageData(v []*DescribeRDDomainsResponseBodyDomainsPageData) *DescribeRDDomainsResponseBodyDomains {
	s.PageData = v
	return s
}

func (s *DescribeRDDomainsResponseBodyDomains) Validate() error {
	return dara.Validate(s)
}

type DescribeRDDomainsResponseBodyDomainsPageData struct {
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// The CNAME assigned to the accelerated domain name.
	//
	// example:
	//
	// image.developer.aliyundoc.com
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
	// 	- online: The domain name is enabled.
	//
	// 	- offline: The domain name is disabled.
	//
	// 	- configuring: The domain name is being configured.
	//
	// 	- configure_failed: The domain name failed to be configured.
	//
	// 	- checking: The domain name is being reviewed.
	//
	// 	- check_failed: The domain name failed the review.
	//
	// example:
	//
	// online
	DomainStatus *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	// The time when the accelerated domain name was added to DCDN.
	//
	// example:
	//
	// 2015-10-27T06:26:34Z
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The time when the accelerated domain name was modified.
	//
	// example:
	//
	// 2015-10-23T09:30:00Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// abcd1234abcd1234
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Indicates whether the accelerated domain name was in a sandbox.
	//
	// example:
	//
	// normal
	Sandbox     *string `json:"Sandbox,omitempty" xml:"Sandbox,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// The information about the origin server.
	Sources *DescribeRDDomainsResponseBodyDomainsPageDataSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Struct"`
	// Indicates whether HTTPS is enabled.
	//
	// 	- on
	//
	// 	- off
	//
	// example:
	//
	// on
	SslProtocol *string `json:"SslProtocol,omitempty" xml:"SslProtocol,omitempty"`
}

func (s DescribeRDDomainsResponseBodyDomainsPageData) String() string {
	return dara.Prettify(s)
}

func (s DescribeRDDomainsResponseBodyDomainsPageData) GoString() string {
	return s.String()
}

func (s *DescribeRDDomainsResponseBodyDomainsPageData) GetBizName() *string {
	return s.BizName
}

func (s *DescribeRDDomainsResponseBodyDomainsPageData) GetCname() *string {
	return s.Cname
}

func (s *DescribeRDDomainsResponseBodyDomainsPageData) GetDescription() *string {
	return s.Description
}

func (s *DescribeRDDomainsResponseBodyDomainsPageData) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeRDDomainsResponseBodyDomainsPageData) GetDomainStatus() *string {
	return s.DomainStatus
}

func (s *DescribeRDDomainsResponseBodyDomainsPageData) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeRDDomainsResponseBodyDomainsPageData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeRDDomainsResponseBodyDomainsPageData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeRDDomainsResponseBodyDomainsPageData) GetSandbox() *string {
	return s.Sandbox
}

func (s *DescribeRDDomainsResponseBodyDomainsPageData) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *DescribeRDDomainsResponseBodyDomainsPageData) GetSources() *DescribeRDDomainsResponseBodyDomainsPageDataSources {
	return s.Sources
}

func (s *DescribeRDDomainsResponseBodyDomainsPageData) GetSslProtocol() *string {
	return s.SslProtocol
}

func (s *DescribeRDDomainsResponseBodyDomainsPageData) SetBizName(v string) *DescribeRDDomainsResponseBodyDomainsPageData {
	s.BizName = &v
	return s
}

func (s *DescribeRDDomainsResponseBodyDomainsPageData) SetCname(v string) *DescribeRDDomainsResponseBodyDomainsPageData {
	s.Cname = &v
	return s
}

func (s *DescribeRDDomainsResponseBodyDomainsPageData) SetDescription(v string) *DescribeRDDomainsResponseBodyDomainsPageData {
	s.Description = &v
	return s
}

func (s *DescribeRDDomainsResponseBodyDomainsPageData) SetDomainName(v string) *DescribeRDDomainsResponseBodyDomainsPageData {
	s.DomainName = &v
	return s
}

func (s *DescribeRDDomainsResponseBodyDomainsPageData) SetDomainStatus(v string) *DescribeRDDomainsResponseBodyDomainsPageData {
	s.DomainStatus = &v
	return s
}

func (s *DescribeRDDomainsResponseBodyDomainsPageData) SetGmtCreated(v string) *DescribeRDDomainsResponseBodyDomainsPageData {
	s.GmtCreated = &v
	return s
}

func (s *DescribeRDDomainsResponseBodyDomainsPageData) SetGmtModified(v string) *DescribeRDDomainsResponseBodyDomainsPageData {
	s.GmtModified = &v
	return s
}

func (s *DescribeRDDomainsResponseBodyDomainsPageData) SetResourceGroupId(v string) *DescribeRDDomainsResponseBodyDomainsPageData {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeRDDomainsResponseBodyDomainsPageData) SetSandbox(v string) *DescribeRDDomainsResponseBodyDomainsPageData {
	s.Sandbox = &v
	return s
}

func (s *DescribeRDDomainsResponseBodyDomainsPageData) SetServiceCode(v string) *DescribeRDDomainsResponseBodyDomainsPageData {
	s.ServiceCode = &v
	return s
}

func (s *DescribeRDDomainsResponseBodyDomainsPageData) SetSources(v *DescribeRDDomainsResponseBodyDomainsPageDataSources) *DescribeRDDomainsResponseBodyDomainsPageData {
	s.Sources = v
	return s
}

func (s *DescribeRDDomainsResponseBodyDomainsPageData) SetSslProtocol(v string) *DescribeRDDomainsResponseBodyDomainsPageData {
	s.SslProtocol = &v
	return s
}

func (s *DescribeRDDomainsResponseBodyDomainsPageData) Validate() error {
	return dara.Validate(s)
}

type DescribeRDDomainsResponseBodyDomainsPageDataSources struct {
	Source []*DescribeRDDomainsResponseBodyDomainsPageDataSourcesSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Repeated"`
}

func (s DescribeRDDomainsResponseBodyDomainsPageDataSources) String() string {
	return dara.Prettify(s)
}

func (s DescribeRDDomainsResponseBodyDomainsPageDataSources) GoString() string {
	return s.String()
}

func (s *DescribeRDDomainsResponseBodyDomainsPageDataSources) GetSource() []*DescribeRDDomainsResponseBodyDomainsPageDataSourcesSource {
	return s.Source
}

func (s *DescribeRDDomainsResponseBodyDomainsPageDataSources) SetSource(v []*DescribeRDDomainsResponseBodyDomainsPageDataSourcesSource) *DescribeRDDomainsResponseBodyDomainsPageDataSources {
	s.Source = v
	return s
}

func (s *DescribeRDDomainsResponseBodyDomainsPageDataSources) Validate() error {
	return dara.Validate(s)
}

type DescribeRDDomainsResponseBodyDomainsPageDataSourcesSource struct {
	// The address of the origin server.
	//
	// example:
	//
	// example.aliyundoc.com
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
	// The origin server weight if multiple origin servers have been specified.
	//
	// example:
	//
	// 20
	Weight *string `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeRDDomainsResponseBodyDomainsPageDataSourcesSource) String() string {
	return dara.Prettify(s)
}

func (s DescribeRDDomainsResponseBodyDomainsPageDataSourcesSource) GoString() string {
	return s.String()
}

func (s *DescribeRDDomainsResponseBodyDomainsPageDataSourcesSource) GetContent() *string {
	return s.Content
}

func (s *DescribeRDDomainsResponseBodyDomainsPageDataSourcesSource) GetPort() *int32 {
	return s.Port
}

func (s *DescribeRDDomainsResponseBodyDomainsPageDataSourcesSource) GetPriority() *string {
	return s.Priority
}

func (s *DescribeRDDomainsResponseBodyDomainsPageDataSourcesSource) GetType() *string {
	return s.Type
}

func (s *DescribeRDDomainsResponseBodyDomainsPageDataSourcesSource) GetWeight() *string {
	return s.Weight
}

func (s *DescribeRDDomainsResponseBodyDomainsPageDataSourcesSource) SetContent(v string) *DescribeRDDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Content = &v
	return s
}

func (s *DescribeRDDomainsResponseBodyDomainsPageDataSourcesSource) SetPort(v int32) *DescribeRDDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Port = &v
	return s
}

func (s *DescribeRDDomainsResponseBodyDomainsPageDataSourcesSource) SetPriority(v string) *DescribeRDDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Priority = &v
	return s
}

func (s *DescribeRDDomainsResponseBodyDomainsPageDataSourcesSource) SetType(v string) *DescribeRDDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Type = &v
	return s
}

func (s *DescribeRDDomainsResponseBodyDomainsPageDataSourcesSource) SetWeight(v string) *DescribeRDDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Weight = &v
	return s
}

func (s *DescribeRDDomainsResponseBodyDomainsPageDataSourcesSource) Validate() error {
	return dara.Validate(s)
}
