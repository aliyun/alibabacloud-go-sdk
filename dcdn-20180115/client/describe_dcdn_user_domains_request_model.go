// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnUserDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChangeEndTime(v string) *DescribeDcdnUserDomainsRequest
	GetChangeEndTime() *string
	SetChangeStartTime(v string) *DescribeDcdnUserDomainsRequest
	GetChangeStartTime() *string
	SetCheckDomainShow(v bool) *DescribeDcdnUserDomainsRequest
	GetCheckDomainShow() *bool
	SetCoverage(v string) *DescribeDcdnUserDomainsRequest
	GetCoverage() *string
	SetDomainName(v string) *DescribeDcdnUserDomainsRequest
	GetDomainName() *string
	SetDomainSearchType(v string) *DescribeDcdnUserDomainsRequest
	GetDomainSearchType() *string
	SetDomainStatus(v string) *DescribeDcdnUserDomainsRequest
	GetDomainStatus() *string
	SetOwnerId(v int64) *DescribeDcdnUserDomainsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeDcdnUserDomainsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDcdnUserDomainsRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *DescribeDcdnUserDomainsRequest
	GetResourceGroupId() *string
	SetSecurityToken(v string) *DescribeDcdnUserDomainsRequest
	GetSecurityToken() *string
	SetTag(v []*DescribeDcdnUserDomainsRequestTag) *DescribeDcdnUserDomainsRequest
	GetTag() []*DescribeDcdnUserDomainsRequestTag
	SetWebSiteType(v string) *DescribeDcdnUserDomainsRequest
	GetWebSiteType() *string
}

type DescribeDcdnUserDomainsRequest struct {
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC+0.
	//
	// > The end time must be later than the start time.
	//
	// example:
	//
	// 2018-05-10T14:00:00Z
	ChangeEndTime *string `json:"ChangeEndTime,omitempty" xml:"ChangeEndTime,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC+0.
	//
	// example:
	//
	// 2018-05-10T12:00:00Z
	ChangeStartTime *string `json:"ChangeStartTime,omitempty" xml:"ChangeStartTime,omitempty"`
	// Specifies whether to display domain names that are under review, failed the review, or failed to be configured. Valid values:
	//
	// 	- true: displays domain names.
	//
	// 	- false: does not display detailed information.
	//
	// example:
	//
	// false
	CheckDomainShow *bool `json:"CheckDomainShow,omitempty" xml:"CheckDomainShow,omitempty"`
	// The acceleration region. By default, all acceleration regions are queried.
	//
	// 	- **domestic**: Chinese mainland
	//
	// 	- **overseas**: outside the Chinese mainland
	//
	// 	- **global**: global
	//
	// example:
	//
	// domestic
	Coverage *string `json:"Coverage,omitempty" xml:"Coverage,omitempty"`
	// The accelerated domain names. If you do not set this parameter, configurations of all domain names that match the conditions are returned.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The search method. Default value: full_match. Valid values:
	//
	// 	- **fuzzy_match**: fuzzy match
	//
	// 	- **pre_match**: prefix match
	//
	// 	- **suf_match**: suffix match
	//
	// 	- **full_match*	- (default): exact match
	//
	// > If you specify the domain names to query but do not set the DomainSearchType parameter, the exact match mode is used.
	//
	// example:
	//
	// fuzzy_match
	DomainSearchType *string `json:"DomainSearchType,omitempty" xml:"DomainSearchType,omitempty"`
	// The status of the domain name. Valid values:
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
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of returned pages. Valid values: **1*	- to **100000**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: **20**. Valid values: **1*	- to **500**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmv6jutt**
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The list of tags.
	Tag []*DescribeDcdnUserDomainsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The business type of the domain. Separate multiple values with commas (,). Default value: **dynamic**. To query common domains, keep the default value. To query domains of the computing business type, enter **computing_routine*	- or **computing_image**.
	//
	// example:
	//
	// computing_routine
	WebSiteType *string `json:"WebSiteType,omitempty" xml:"WebSiteType,omitempty"`
}

func (s DescribeDcdnUserDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserDomainsRequest) GetChangeEndTime() *string {
	return s.ChangeEndTime
}

func (s *DescribeDcdnUserDomainsRequest) GetChangeStartTime() *string {
	return s.ChangeStartTime
}

func (s *DescribeDcdnUserDomainsRequest) GetCheckDomainShow() *bool {
	return s.CheckDomainShow
}

func (s *DescribeDcdnUserDomainsRequest) GetCoverage() *string {
	return s.Coverage
}

func (s *DescribeDcdnUserDomainsRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnUserDomainsRequest) GetDomainSearchType() *string {
	return s.DomainSearchType
}

func (s *DescribeDcdnUserDomainsRequest) GetDomainStatus() *string {
	return s.DomainStatus
}

func (s *DescribeDcdnUserDomainsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDcdnUserDomainsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDcdnUserDomainsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDcdnUserDomainsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDcdnUserDomainsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeDcdnUserDomainsRequest) GetTag() []*DescribeDcdnUserDomainsRequestTag {
	return s.Tag
}

func (s *DescribeDcdnUserDomainsRequest) GetWebSiteType() *string {
	return s.WebSiteType
}

func (s *DescribeDcdnUserDomainsRequest) SetChangeEndTime(v string) *DescribeDcdnUserDomainsRequest {
	s.ChangeEndTime = &v
	return s
}

func (s *DescribeDcdnUserDomainsRequest) SetChangeStartTime(v string) *DescribeDcdnUserDomainsRequest {
	s.ChangeStartTime = &v
	return s
}

func (s *DescribeDcdnUserDomainsRequest) SetCheckDomainShow(v bool) *DescribeDcdnUserDomainsRequest {
	s.CheckDomainShow = &v
	return s
}

func (s *DescribeDcdnUserDomainsRequest) SetCoverage(v string) *DescribeDcdnUserDomainsRequest {
	s.Coverage = &v
	return s
}

func (s *DescribeDcdnUserDomainsRequest) SetDomainName(v string) *DescribeDcdnUserDomainsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnUserDomainsRequest) SetDomainSearchType(v string) *DescribeDcdnUserDomainsRequest {
	s.DomainSearchType = &v
	return s
}

func (s *DescribeDcdnUserDomainsRequest) SetDomainStatus(v string) *DescribeDcdnUserDomainsRequest {
	s.DomainStatus = &v
	return s
}

func (s *DescribeDcdnUserDomainsRequest) SetOwnerId(v int64) *DescribeDcdnUserDomainsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnUserDomainsRequest) SetPageNumber(v int32) *DescribeDcdnUserDomainsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnUserDomainsRequest) SetPageSize(v int32) *DescribeDcdnUserDomainsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnUserDomainsRequest) SetResourceGroupId(v string) *DescribeDcdnUserDomainsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDcdnUserDomainsRequest) SetSecurityToken(v string) *DescribeDcdnUserDomainsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDcdnUserDomainsRequest) SetTag(v []*DescribeDcdnUserDomainsRequestTag) *DescribeDcdnUserDomainsRequest {
	s.Tag = v
	return s
}

func (s *DescribeDcdnUserDomainsRequest) SetWebSiteType(v string) *DescribeDcdnUserDomainsRequest {
	s.WebSiteType = &v
	return s
}

func (s *DescribeDcdnUserDomainsRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnUserDomainsRequestTag struct {
	// The tag key. Valid values of N: **1*	- to **20**. You can call the TagDcdnResources operation to set a tag for a domain name.
	//
	// example:
	//
	// 1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. Valid values of N: **1*	- to **20**.
	//
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDcdnUserDomainsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserDomainsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserDomainsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeDcdnUserDomainsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeDcdnUserDomainsRequestTag) SetKey(v string) *DescribeDcdnUserDomainsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDcdnUserDomainsRequestTag) SetValue(v string) *DescribeDcdnUserDomainsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeDcdnUserDomainsRequestTag) Validate() error {
	return dara.Validate(s)
}
