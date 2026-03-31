// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudResourceListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloudResourceId(v string) *DescribeCloudResourceListRequest
	GetCloudResourceId() *string
	SetInstanceId(v string) *DescribeCloudResourceListRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *DescribeCloudResourceListRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeCloudResourceListRequest
	GetNextToken() *string
	SetOwnerUserId(v string) *DescribeCloudResourceListRequest
	GetOwnerUserId() *string
	SetPort(v string) *DescribeCloudResourceListRequest
	GetPort() *string
	SetRegionId(v string) *DescribeCloudResourceListRequest
	GetRegionId() *string
	SetResourceInstanceId(v string) *DescribeCloudResourceListRequest
	GetResourceInstanceId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeCloudResourceListRequest
	GetResourceManagerResourceGroupId() *string
	SetResourceProduct(v string) *DescribeCloudResourceListRequest
	GetResourceProduct() *string
}

type DescribeCloudResourceListRequest struct {
	// example:
	//
	// i-8vbdlsd********81e22-80-ecs
	CloudResourceId *string `json:"CloudResourceId,omitempty" xml:"CloudResourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf_elasticity-cn-0xldbqt****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAGBgV9tolsLfijC4wam2htS*****D/46H3X2wIS
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1111111111
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// example:
	//
	// 80
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// i-8vbdlsd********81e22
	ResourceInstanceId *string `json:"ResourceInstanceId,omitempty" xml:"ResourceInstanceId,omitempty"`
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// example:
	//
	// ecs
	ResourceProduct *string `json:"ResourceProduct,omitempty" xml:"ResourceProduct,omitempty"`
}

func (s DescribeCloudResourceListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudResourceListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudResourceListRequest) GetCloudResourceId() *string {
	return s.CloudResourceId
}

func (s *DescribeCloudResourceListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCloudResourceListRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeCloudResourceListRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeCloudResourceListRequest) GetOwnerUserId() *string {
	return s.OwnerUserId
}

func (s *DescribeCloudResourceListRequest) GetPort() *string {
	return s.Port
}

func (s *DescribeCloudResourceListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCloudResourceListRequest) GetResourceInstanceId() *string {
	return s.ResourceInstanceId
}

func (s *DescribeCloudResourceListRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeCloudResourceListRequest) GetResourceProduct() *string {
	return s.ResourceProduct
}

func (s *DescribeCloudResourceListRequest) SetCloudResourceId(v string) *DescribeCloudResourceListRequest {
	s.CloudResourceId = &v
	return s
}

func (s *DescribeCloudResourceListRequest) SetInstanceId(v string) *DescribeCloudResourceListRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeCloudResourceListRequest) SetMaxResults(v int32) *DescribeCloudResourceListRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeCloudResourceListRequest) SetNextToken(v string) *DescribeCloudResourceListRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeCloudResourceListRequest) SetOwnerUserId(v string) *DescribeCloudResourceListRequest {
	s.OwnerUserId = &v
	return s
}

func (s *DescribeCloudResourceListRequest) SetPort(v string) *DescribeCloudResourceListRequest {
	s.Port = &v
	return s
}

func (s *DescribeCloudResourceListRequest) SetRegionId(v string) *DescribeCloudResourceListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCloudResourceListRequest) SetResourceInstanceId(v string) *DescribeCloudResourceListRequest {
	s.ResourceInstanceId = &v
	return s
}

func (s *DescribeCloudResourceListRequest) SetResourceManagerResourceGroupId(v string) *DescribeCloudResourceListRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeCloudResourceListRequest) SetResourceProduct(v string) *DescribeCloudResourceListRequest {
	s.ResourceProduct = &v
	return s
}

func (s *DescribeCloudResourceListRequest) Validate() error {
	return dara.Validate(s)
}
