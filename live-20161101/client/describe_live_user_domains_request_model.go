// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveUserDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveUserDomainsRequest
	GetDomainName() *string
	SetDomainSearchType(v string) *DescribeLiveUserDomainsRequest
	GetDomainSearchType() *string
	SetDomainStatus(v string) *DescribeLiveUserDomainsRequest
	GetDomainStatus() *string
	SetLiveDomainType(v string) *DescribeLiveUserDomainsRequest
	GetLiveDomainType() *string
	SetOwnerId(v int64) *DescribeLiveUserDomainsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeLiveUserDomainsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeLiveUserDomainsRequest
	GetPageSize() *int32
	SetRegionName(v string) *DescribeLiveUserDomainsRequest
	GetRegionName() *string
	SetResourceGroupId(v string) *DescribeLiveUserDomainsRequest
	GetResourceGroupId() *string
	SetSecurityToken(v string) *DescribeLiveUserDomainsRequest
	GetSecurityToken() *string
	SetTag(v []*DescribeLiveUserDomainsRequestTag) *DescribeLiveUserDomainsRequest
	GetTag() []*DescribeLiveUserDomainsRequestTag
}

type DescribeLiveUserDomainsRequest struct {
	// The domain name that is used as a keyword to filter domain names. Fuzzy match is supported.
	//
	// >
	//
	// 	- If you set LiveDomainType to liveVideo and leave this parameter empty, the streaming domains are queried. - If you set LiveDomainType to liveEdge and leave this parameter empty, the ingest domains are queried.
	//
	// example:
	//
	// *.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The search mode. Valid values:
	//
	// 	- **fuzzy_match*	- (default): fuzzy match
	//
	// 	- **pre_match**: prefix match
	//
	// 	- **suf_match**: suffix match
	//
	// 	- **full_match**: exact match
	//
	// example:
	//
	// fuzzy_match
	DomainSearchType *string `json:"DomainSearchType,omitempty" xml:"DomainSearchType,omitempty"`
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
	DomainStatus *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	// The type of the domain name. Valid values:
	//
	// 	- **liveVideo**: streaming domain
	//
	// 	- **liveEdge**: ingest domain
	//
	// >  If you leave this parameter empty, all ingest domains and streaming domains within your Alibaba Cloud account are queried by default.
	//
	// example:
	//
	// liveVideo
	LiveDomainType *string `json:"LiveDomainType,omitempty" xml:"LiveDomainType,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Valid values: **1 to 100000**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **20**. Maximum value: **50**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	// rg-aek2egyoep3jp7a
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The tags.
	Tag []*DescribeLiveUserDomainsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeLiveUserDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUserDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveUserDomainsRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveUserDomainsRequest) GetDomainSearchType() *string {
	return s.DomainSearchType
}

func (s *DescribeLiveUserDomainsRequest) GetDomainStatus() *string {
	return s.DomainStatus
}

func (s *DescribeLiveUserDomainsRequest) GetLiveDomainType() *string {
	return s.LiveDomainType
}

func (s *DescribeLiveUserDomainsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveUserDomainsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeLiveUserDomainsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLiveUserDomainsRequest) GetRegionName() *string {
	return s.RegionName
}

func (s *DescribeLiveUserDomainsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeLiveUserDomainsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeLiveUserDomainsRequest) GetTag() []*DescribeLiveUserDomainsRequestTag {
	return s.Tag
}

func (s *DescribeLiveUserDomainsRequest) SetDomainName(v string) *DescribeLiveUserDomainsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveUserDomainsRequest) SetDomainSearchType(v string) *DescribeLiveUserDomainsRequest {
	s.DomainSearchType = &v
	return s
}

func (s *DescribeLiveUserDomainsRequest) SetDomainStatus(v string) *DescribeLiveUserDomainsRequest {
	s.DomainStatus = &v
	return s
}

func (s *DescribeLiveUserDomainsRequest) SetLiveDomainType(v string) *DescribeLiveUserDomainsRequest {
	s.LiveDomainType = &v
	return s
}

func (s *DescribeLiveUserDomainsRequest) SetOwnerId(v int64) *DescribeLiveUserDomainsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveUserDomainsRequest) SetPageNumber(v int32) *DescribeLiveUserDomainsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLiveUserDomainsRequest) SetPageSize(v int32) *DescribeLiveUserDomainsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveUserDomainsRequest) SetRegionName(v string) *DescribeLiveUserDomainsRequest {
	s.RegionName = &v
	return s
}

func (s *DescribeLiveUserDomainsRequest) SetResourceGroupId(v string) *DescribeLiveUserDomainsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeLiveUserDomainsRequest) SetSecurityToken(v string) *DescribeLiveUserDomainsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeLiveUserDomainsRequest) SetTag(v []*DescribeLiveUserDomainsRequestTag) *DescribeLiveUserDomainsRequest {
	s.Tag = v
	return s
}

func (s *DescribeLiveUserDomainsRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveUserDomainsRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeLiveUserDomainsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUserDomainsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeLiveUserDomainsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeLiveUserDomainsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeLiveUserDomainsRequestTag) SetKey(v string) *DescribeLiveUserDomainsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeLiveUserDomainsRequestTag) SetValue(v string) *DescribeLiveUserDomainsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeLiveUserDomainsRequestTag) Validate() error {
	return dara.Validate(s)
}
