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
	if s.Domains != nil {
		if err := s.Domains.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.PageData != nil {
		for _, item := range s.PageData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeUserDomainsResponseBodyDomainsPageData struct {
	CdnType         *string                                                `json:"CdnType,omitempty" xml:"CdnType,omitempty"`
	Cname           *string                                                `json:"Cname,omitempty" xml:"Cname,omitempty"`
	Coverage        *string                                                `json:"Coverage,omitempty" xml:"Coverage,omitempty"`
	Description     *string                                                `json:"Description,omitempty" xml:"Description,omitempty"`
	DomainId        *int64                                                 `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	DomainName      *string                                                `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainStatus    *string                                                `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	GmtCreated      *string                                                `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtModified     *string                                                `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	ResourceGroupId *string                                                `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Sandbox         *string                                                `json:"Sandbox,omitempty" xml:"Sandbox,omitempty"`
	Sources         *DescribeUserDomainsResponseBodyDomainsPageDataSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Struct"`
	SslProtocol     *string                                                `json:"SslProtocol,omitempty" xml:"SslProtocol,omitempty"`
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
	if s.Sources != nil {
		if err := s.Sources.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.Source != nil {
		for _, item := range s.Source {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeUserDomainsResponseBodyDomainsPageDataSourcesSource struct {
	Content  *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Weight   *string `json:"Weight,omitempty" xml:"Weight,omitempty"`
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
