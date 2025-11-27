// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAllWhitelistTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFuzzySearch(v bool) *DescribeAllWhitelistTemplateRequest
	GetFuzzySearch() *bool
	SetMaxRecordsPerPage(v int32) *DescribeAllWhitelistTemplateRequest
	GetMaxRecordsPerPage() *int32
	SetPageNumbers(v int32) *DescribeAllWhitelistTemplateRequest
	GetPageNumbers() *int32
	SetRegionId(v string) *DescribeAllWhitelistTemplateRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeAllWhitelistTemplateRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeAllWhitelistTemplateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAllWhitelistTemplateRequest
	GetResourceOwnerId() *int64
	SetTemplateName(v string) *DescribeAllWhitelistTemplateRequest
	GetTemplateName() *string
}

type DescribeAllWhitelistTemplateRequest struct {
	// Specifies whether to enable fuzzy search. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	FuzzySearch *bool `json:"FuzzySearch,omitempty" xml:"FuzzySearch,omitempty"`
	// The number of entries to return on each page. Enumerated valid values: 10, 30, and 50.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	MaxRecordsPerPage *int32 `json:"MaxRecordsPerPage,omitempty" xml:"MaxRecordsPerPage,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumbers *int32 `json:"PageNumbers,omitempty" xml:"PageNumbers,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID. For more information about resource groups, see related documentation.
	//
	// example:
	//
	// rg-acfmyhigxskzysy
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The name of the IP whitelist template. If you specify this parameter when you perform a fuzzy search, you can call the DescribeWhitelistTemplate operation to query the name of the whitelist template during the fuzzy search.
	//
	// example:
	//
	// template
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s DescribeAllWhitelistTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllWhitelistTemplateRequest) GoString() string {
	return s.String()
}

func (s *DescribeAllWhitelistTemplateRequest) GetFuzzySearch() *bool {
	return s.FuzzySearch
}

func (s *DescribeAllWhitelistTemplateRequest) GetMaxRecordsPerPage() *int32 {
	return s.MaxRecordsPerPage
}

func (s *DescribeAllWhitelistTemplateRequest) GetPageNumbers() *int32 {
	return s.PageNumbers
}

func (s *DescribeAllWhitelistTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAllWhitelistTemplateRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeAllWhitelistTemplateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAllWhitelistTemplateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAllWhitelistTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *DescribeAllWhitelistTemplateRequest) SetFuzzySearch(v bool) *DescribeAllWhitelistTemplateRequest {
	s.FuzzySearch = &v
	return s
}

func (s *DescribeAllWhitelistTemplateRequest) SetMaxRecordsPerPage(v int32) *DescribeAllWhitelistTemplateRequest {
	s.MaxRecordsPerPage = &v
	return s
}

func (s *DescribeAllWhitelistTemplateRequest) SetPageNumbers(v int32) *DescribeAllWhitelistTemplateRequest {
	s.PageNumbers = &v
	return s
}

func (s *DescribeAllWhitelistTemplateRequest) SetRegionId(v string) *DescribeAllWhitelistTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAllWhitelistTemplateRequest) SetResourceGroupId(v string) *DescribeAllWhitelistTemplateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeAllWhitelistTemplateRequest) SetResourceOwnerAccount(v string) *DescribeAllWhitelistTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAllWhitelistTemplateRequest) SetResourceOwnerId(v int64) *DescribeAllWhitelistTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAllWhitelistTemplateRequest) SetTemplateName(v string) *DescribeAllWhitelistTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *DescribeAllWhitelistTemplateRequest) Validate() error {
	return dara.Validate(s)
}
