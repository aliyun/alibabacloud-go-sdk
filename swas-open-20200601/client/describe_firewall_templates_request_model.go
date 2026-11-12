// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFirewallTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFirewallTemplateId(v []*string) *DescribeFirewallTemplatesRequest
	GetFirewallTemplateId() []*string
	SetName(v string) *DescribeFirewallTemplatesRequest
	GetName() *string
	SetPageNumber(v int32) *DescribeFirewallTemplatesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeFirewallTemplatesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeFirewallTemplatesRequest
	GetRegionId() *string
}

type DescribeFirewallTemplatesRequest struct {
	// The IDs of the firewall templates.
	FirewallTemplateId []*string `json:"FirewallTemplateId,omitempty" xml:"FirewallTemplateId,omitempty" type:"Repeated"`
	// The name of the firewall template.
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The page number.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeFirewallTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirewallTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeFirewallTemplatesRequest) GetFirewallTemplateId() []*string {
	return s.FirewallTemplateId
}

func (s *DescribeFirewallTemplatesRequest) GetName() *string {
	return s.Name
}

func (s *DescribeFirewallTemplatesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeFirewallTemplatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeFirewallTemplatesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeFirewallTemplatesRequest) SetFirewallTemplateId(v []*string) *DescribeFirewallTemplatesRequest {
	s.FirewallTemplateId = v
	return s
}

func (s *DescribeFirewallTemplatesRequest) SetName(v string) *DescribeFirewallTemplatesRequest {
	s.Name = &v
	return s
}

func (s *DescribeFirewallTemplatesRequest) SetPageNumber(v int32) *DescribeFirewallTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeFirewallTemplatesRequest) SetPageSize(v int32) *DescribeFirewallTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeFirewallTemplatesRequest) SetRegionId(v string) *DescribeFirewallTemplatesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFirewallTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
