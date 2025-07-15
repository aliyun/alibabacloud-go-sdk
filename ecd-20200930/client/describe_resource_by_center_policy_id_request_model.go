// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceByCenterPolicyIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeResourceByCenterPolicyIdRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeResourceByCenterPolicyIdRequest
	GetNextToken() *string
	SetPolicyGroupId(v string) *DescribeResourceByCenterPolicyIdRequest
	GetPolicyGroupId() *string
	SetProductType(v string) *DescribeResourceByCenterPolicyIdRequest
	GetProductType() *string
	SetResourceId(v string) *DescribeResourceByCenterPolicyIdRequest
	GetResourceId() *string
}

type DescribeResourceByCenterPolicyIdRequest struct {
	// The number of entries per page.
	//
	// 	- Maximum value: 100.
	//
	// 	- Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6l69tQX7yFxx6/4dbooBAOc
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The policy ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pg-53iyi2aar0nd6****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// The service type.
	//
	// Valid values:
	//
	// 	- app: cloud applications.
	//
	// 	- resourceGroup: resource groups.
	//
	// 	- desktop: cloud computers.
	//
	// 	- desktopGroup: cloud computer shares.
	//
	// example:
	//
	// desktop
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// ecd-ia2zw38bi6cm7****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s DescribeResourceByCenterPolicyIdRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceByCenterPolicyIdRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourceByCenterPolicyIdRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeResourceByCenterPolicyIdRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeResourceByCenterPolicyIdRequest) GetPolicyGroupId() *string {
	return s.PolicyGroupId
}

func (s *DescribeResourceByCenterPolicyIdRequest) GetProductType() *string {
	return s.ProductType
}

func (s *DescribeResourceByCenterPolicyIdRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeResourceByCenterPolicyIdRequest) SetMaxResults(v int32) *DescribeResourceByCenterPolicyIdRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeResourceByCenterPolicyIdRequest) SetNextToken(v string) *DescribeResourceByCenterPolicyIdRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeResourceByCenterPolicyIdRequest) SetPolicyGroupId(v string) *DescribeResourceByCenterPolicyIdRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *DescribeResourceByCenterPolicyIdRequest) SetProductType(v string) *DescribeResourceByCenterPolicyIdRequest {
	s.ProductType = &v
	return s
}

func (s *DescribeResourceByCenterPolicyIdRequest) SetResourceId(v string) *DescribeResourceByCenterPolicyIdRequest {
	s.ResourceId = &v
	return s
}

func (s *DescribeResourceByCenterPolicyIdRequest) Validate() error {
	return dara.Validate(s)
}
