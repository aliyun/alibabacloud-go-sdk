// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebCCRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeWebCCRulesRequest
	GetDomain() *string
	SetPageNumber(v int32) *DescribeWebCCRulesRequest
	GetPageNumber() *int32
	SetPageSize(v string) *DescribeWebCCRulesRequest
	GetPageSize() *string
	SetResourceGroupId(v string) *DescribeWebCCRulesRequest
	GetResourceGroupId() *string
}

type DescribeWebCCRulesRequest struct {
	// The domain name of the website.
	//
	// > A forwarding rule must be configured for the domain name. You can call the [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) operation to query all domain names.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The number of the page to return. For example, to query the returned results on the first page, set the value to **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the resource group to which the instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeWebCCRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebCCRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebCCRulesRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeWebCCRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeWebCCRulesRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeWebCCRulesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeWebCCRulesRequest) SetDomain(v string) *DescribeWebCCRulesRequest {
	s.Domain = &v
	return s
}

func (s *DescribeWebCCRulesRequest) SetPageNumber(v int32) *DescribeWebCCRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeWebCCRulesRequest) SetPageSize(v string) *DescribeWebCCRulesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeWebCCRulesRequest) SetResourceGroupId(v string) *DescribeWebCCRulesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeWebCCRulesRequest) Validate() error {
	return dara.Validate(s)
}
