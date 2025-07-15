// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationGroupBillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationName(v string) *DescribeApplicationGroupBillRequest
	GetApplicationName() *string
	SetBillingCycle(v string) *DescribeApplicationGroupBillRequest
	GetBillingCycle() *string
	SetMaxResults(v int32) *DescribeApplicationGroupBillRequest
	GetMaxResults() *int32
	SetName(v string) *DescribeApplicationGroupBillRequest
	GetName() *string
	SetNextToken(v string) *DescribeApplicationGroupBillRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeApplicationGroupBillRequest
	GetRegionId() *string
	SetResourceType(v string) *DescribeApplicationGroupBillRequest
	GetResourceType() *string
}

type DescribeApplicationGroupBillRequest struct {
	// The application name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_application
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The billing cycle, in the YYYY-MM format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-06
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The application group name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_application_group
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// -
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the cloud resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// ALIYUN::ECS::INSTANCE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeApplicationGroupBillRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationGroupBillRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationGroupBillRequest) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *DescribeApplicationGroupBillRequest) GetBillingCycle() *string {
	return s.BillingCycle
}

func (s *DescribeApplicationGroupBillRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeApplicationGroupBillRequest) GetName() *string {
	return s.Name
}

func (s *DescribeApplicationGroupBillRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeApplicationGroupBillRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApplicationGroupBillRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeApplicationGroupBillRequest) SetApplicationName(v string) *DescribeApplicationGroupBillRequest {
	s.ApplicationName = &v
	return s
}

func (s *DescribeApplicationGroupBillRequest) SetBillingCycle(v string) *DescribeApplicationGroupBillRequest {
	s.BillingCycle = &v
	return s
}

func (s *DescribeApplicationGroupBillRequest) SetMaxResults(v int32) *DescribeApplicationGroupBillRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeApplicationGroupBillRequest) SetName(v string) *DescribeApplicationGroupBillRequest {
	s.Name = &v
	return s
}

func (s *DescribeApplicationGroupBillRequest) SetNextToken(v string) *DescribeApplicationGroupBillRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeApplicationGroupBillRequest) SetRegionId(v string) *DescribeApplicationGroupBillRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApplicationGroupBillRequest) SetResourceType(v string) *DescribeApplicationGroupBillRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeApplicationGroupBillRequest) Validate() error {
	return dara.Validate(s)
}
