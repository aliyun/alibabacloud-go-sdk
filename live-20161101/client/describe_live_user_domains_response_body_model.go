// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveUserDomainsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomains(v *DescribeLiveUserDomainsResponseBodyDomains) *DescribeLiveUserDomainsResponseBody
	GetDomains() *DescribeLiveUserDomainsResponseBodyDomains
	SetPageNumber(v int64) *DescribeLiveUserDomainsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeLiveUserDomainsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeLiveUserDomainsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeLiveUserDomainsResponseBody
	GetTotalCount() *int64
}

type DescribeLiveUserDomainsResponseBody struct {
	// The domain names.
	Domains *DescribeLiveUserDomainsResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Struct"`
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
	// 100
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E4EBD2BF-5EB0-4476-8829-9D94E1B15267
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeLiveUserDomainsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUserDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveUserDomainsResponseBody) GetDomains() *DescribeLiveUserDomainsResponseBodyDomains {
	return s.Domains
}

func (s *DescribeLiveUserDomainsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeLiveUserDomainsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeLiveUserDomainsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveUserDomainsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeLiveUserDomainsResponseBody) SetDomains(v *DescribeLiveUserDomainsResponseBodyDomains) *DescribeLiveUserDomainsResponseBody {
	s.Domains = v
	return s
}

func (s *DescribeLiveUserDomainsResponseBody) SetPageNumber(v int64) *DescribeLiveUserDomainsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeLiveUserDomainsResponseBody) SetPageSize(v int64) *DescribeLiveUserDomainsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveUserDomainsResponseBody) SetRequestId(v string) *DescribeLiveUserDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveUserDomainsResponseBody) SetTotalCount(v int64) *DescribeLiveUserDomainsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeLiveUserDomainsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveUserDomainsResponseBodyDomains struct {
	PageData []*DescribeLiveUserDomainsResponseBodyDomainsPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
}

func (s DescribeLiveUserDomainsResponseBodyDomains) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUserDomainsResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *DescribeLiveUserDomainsResponseBodyDomains) GetPageData() []*DescribeLiveUserDomainsResponseBodyDomainsPageData {
	return s.PageData
}

func (s *DescribeLiveUserDomainsResponseBodyDomains) SetPageData(v []*DescribeLiveUserDomainsResponseBodyDomainsPageData) *DescribeLiveUserDomainsResponseBodyDomains {
	s.PageData = v
	return s
}

func (s *DescribeLiveUserDomainsResponseBodyDomains) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveUserDomainsResponseBodyDomainsPageData struct {
	// The CNAME generated for the domain name.
	//
	// example:
	//
	// learn.developer.aliyundoc.com
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// The description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The domain name.
	//
	// example:
	//
	// demo.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The time when the domain name was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-08-29T12:15:36Z
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The time when the domain name was last modified. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-12-29T12:15:36Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The status of the domain name. Valid values:
	//
	// 	- **online**
	//
	// 	- **offline**
	//
	// 	- **configuring**
	//
	// example:
	//
	// online
	LiveDomainStatus *string `json:"LiveDomainStatus,omitempty" xml:"LiveDomainStatus,omitempty"`
	// The type of the domain name. Valid values:
	//
	// 	- **liveVideo**: streaming domain
	//
	// 	- **liveEdge**: ingest domain
	//
	// example:
	//
	// liveVideo
	LiveDomainType *string `json:"LiveDomainType,omitempty" xml:"LiveDomainType,omitempty"`
	// The ID of the region in which the domain name resides.
	//
	// example:
	//
	// cn-beijing
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekzvyjetyjjmla
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeLiveUserDomainsResponseBodyDomainsPageData) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUserDomainsResponseBodyDomainsPageData) GoString() string {
	return s.String()
}

func (s *DescribeLiveUserDomainsResponseBodyDomainsPageData) GetCname() *string {
	return s.Cname
}

func (s *DescribeLiveUserDomainsResponseBodyDomainsPageData) GetDescription() *string {
	return s.Description
}

func (s *DescribeLiveUserDomainsResponseBodyDomainsPageData) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveUserDomainsResponseBodyDomainsPageData) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeLiveUserDomainsResponseBodyDomainsPageData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeLiveUserDomainsResponseBodyDomainsPageData) GetLiveDomainStatus() *string {
	return s.LiveDomainStatus
}

func (s *DescribeLiveUserDomainsResponseBodyDomainsPageData) GetLiveDomainType() *string {
	return s.LiveDomainType
}

func (s *DescribeLiveUserDomainsResponseBodyDomainsPageData) GetRegionName() *string {
	return s.RegionName
}

func (s *DescribeLiveUserDomainsResponseBodyDomainsPageData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeLiveUserDomainsResponseBodyDomainsPageData) SetCname(v string) *DescribeLiveUserDomainsResponseBodyDomainsPageData {
	s.Cname = &v
	return s
}

func (s *DescribeLiveUserDomainsResponseBodyDomainsPageData) SetDescription(v string) *DescribeLiveUserDomainsResponseBodyDomainsPageData {
	s.Description = &v
	return s
}

func (s *DescribeLiveUserDomainsResponseBodyDomainsPageData) SetDomainName(v string) *DescribeLiveUserDomainsResponseBodyDomainsPageData {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveUserDomainsResponseBodyDomainsPageData) SetGmtCreated(v string) *DescribeLiveUserDomainsResponseBodyDomainsPageData {
	s.GmtCreated = &v
	return s
}

func (s *DescribeLiveUserDomainsResponseBodyDomainsPageData) SetGmtModified(v string) *DescribeLiveUserDomainsResponseBodyDomainsPageData {
	s.GmtModified = &v
	return s
}

func (s *DescribeLiveUserDomainsResponseBodyDomainsPageData) SetLiveDomainStatus(v string) *DescribeLiveUserDomainsResponseBodyDomainsPageData {
	s.LiveDomainStatus = &v
	return s
}

func (s *DescribeLiveUserDomainsResponseBodyDomainsPageData) SetLiveDomainType(v string) *DescribeLiveUserDomainsResponseBodyDomainsPageData {
	s.LiveDomainType = &v
	return s
}

func (s *DescribeLiveUserDomainsResponseBodyDomainsPageData) SetRegionName(v string) *DescribeLiveUserDomainsResponseBodyDomainsPageData {
	s.RegionName = &v
	return s
}

func (s *DescribeLiveUserDomainsResponseBodyDomainsPageData) SetResourceGroupId(v string) *DescribeLiveUserDomainsResponseBodyDomainsPageData {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeLiveUserDomainsResponseBodyDomainsPageData) Validate() error {
	return dara.Validate(s)
}
