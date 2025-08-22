// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnIpaUserDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckDomainShow(v bool) *DescribeDcdnIpaUserDomainsRequest
	GetCheckDomainShow() *bool
	SetDomainName(v string) *DescribeDcdnIpaUserDomainsRequest
	GetDomainName() *string
	SetDomainSearchType(v string) *DescribeDcdnIpaUserDomainsRequest
	GetDomainSearchType() *string
	SetDomainStatus(v string) *DescribeDcdnIpaUserDomainsRequest
	GetDomainStatus() *string
	SetFuncFilter(v string) *DescribeDcdnIpaUserDomainsRequest
	GetFuncFilter() *string
	SetFuncId(v string) *DescribeDcdnIpaUserDomainsRequest
	GetFuncId() *string
	SetOwnerId(v int64) *DescribeDcdnIpaUserDomainsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeDcdnIpaUserDomainsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDcdnIpaUserDomainsRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *DescribeDcdnIpaUserDomainsRequest
	GetResourceGroupId() *string
	SetTag(v []*DescribeDcdnIpaUserDomainsRequestTag) *DescribeDcdnIpaUserDomainsRequest
	GetTag() []*DescribeDcdnIpaUserDomainsRequestTag
}

type DescribeDcdnIpaUserDomainsRequest struct {
	// Specifies whether to display domain names that are under review, failed the review, or failed to be configured. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	CheckDomainShow *bool `json:"CheckDomainShow,omitempty" xml:"CheckDomainShow,omitempty"`
	// The domain name that is used as a keyword to filter domain names. Fuzzy match is supported.
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
	// 	- **full_match**: exact match
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
	// The status of the feature.
	//
	// 	- config: The feature is enabled.
	//
	// 	- unconfig: The feature is not enabled.
	//
	// example:
	//
	// config
	FuncFilter *string `json:"FuncFilter,omitempty" xml:"FuncFilter,omitempty"`
	// The ID of the feature. For example, a value of 7 specifies the feature of configuring an expiration rule for a specific directory. For more information about feature IDs, see [Parameters for configuring features for domain names](https://help.aliyun.com/document_detail/410622.html).
	//
	// example:
	//
	// 7
	FuncId  *string `json:"FuncId,omitempty" xml:"FuncId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Valid values: **1*	- to **100000**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of domain names per page. Default value: **20**.***	- Valid values: **1*	- to **500**.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// abcd1234abcd****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags of the command.
	Tag []*DescribeDcdnIpaUserDomainsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDcdnIpaUserDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnIpaUserDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaUserDomainsRequest) GetCheckDomainShow() *bool {
	return s.CheckDomainShow
}

func (s *DescribeDcdnIpaUserDomainsRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnIpaUserDomainsRequest) GetDomainSearchType() *string {
	return s.DomainSearchType
}

func (s *DescribeDcdnIpaUserDomainsRequest) GetDomainStatus() *string {
	return s.DomainStatus
}

func (s *DescribeDcdnIpaUserDomainsRequest) GetFuncFilter() *string {
	return s.FuncFilter
}

func (s *DescribeDcdnIpaUserDomainsRequest) GetFuncId() *string {
	return s.FuncId
}

func (s *DescribeDcdnIpaUserDomainsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDcdnIpaUserDomainsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDcdnIpaUserDomainsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDcdnIpaUserDomainsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDcdnIpaUserDomainsRequest) GetTag() []*DescribeDcdnIpaUserDomainsRequestTag {
	return s.Tag
}

func (s *DescribeDcdnIpaUserDomainsRequest) SetCheckDomainShow(v bool) *DescribeDcdnIpaUserDomainsRequest {
	s.CheckDomainShow = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsRequest) SetDomainName(v string) *DescribeDcdnIpaUserDomainsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsRequest) SetDomainSearchType(v string) *DescribeDcdnIpaUserDomainsRequest {
	s.DomainSearchType = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsRequest) SetDomainStatus(v string) *DescribeDcdnIpaUserDomainsRequest {
	s.DomainStatus = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsRequest) SetFuncFilter(v string) *DescribeDcdnIpaUserDomainsRequest {
	s.FuncFilter = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsRequest) SetFuncId(v string) *DescribeDcdnIpaUserDomainsRequest {
	s.FuncId = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsRequest) SetOwnerId(v int64) *DescribeDcdnIpaUserDomainsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsRequest) SetPageNumber(v int32) *DescribeDcdnIpaUserDomainsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsRequest) SetPageSize(v int32) *DescribeDcdnIpaUserDomainsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsRequest) SetResourceGroupId(v string) *DescribeDcdnIpaUserDomainsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsRequest) SetTag(v []*DescribeDcdnIpaUserDomainsRequestTag) *DescribeDcdnIpaUserDomainsRequest {
	s.Tag = v
	return s
}

func (s *DescribeDcdnIpaUserDomainsRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnIpaUserDomainsRequestTag struct {
	// The tag key. Valid values of N: 1 to 20. You can call the TagDcdnResources operation to set a tag for a domain name.
	//
	// example:
	//
	// 1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. Valid values of N: 1 to 20.
	//
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDcdnIpaUserDomainsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnIpaUserDomainsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaUserDomainsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeDcdnIpaUserDomainsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeDcdnIpaUserDomainsRequestTag) SetKey(v string) *DescribeDcdnIpaUserDomainsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsRequestTag) SetValue(v string) *DescribeDcdnIpaUserDomainsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsRequestTag) Validate() error {
	return dara.Validate(s)
}
