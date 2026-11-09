// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceAutoRenewalAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenewal(v bool) *DescribeInstanceAutoRenewalAttributeRequest
	GetAutoRenewal() *bool
	SetInstanceIdList(v []*string) *DescribeInstanceAutoRenewalAttributeRequest
	GetInstanceIdList() []*string
	SetMaxResults(v int32) *DescribeInstanceAutoRenewalAttributeRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeInstanceAutoRenewalAttributeRequest
	GetNextToken() *string
	SetPageNo(v int32) *DescribeInstanceAutoRenewalAttributeRequest
	GetPageNo() *int32
	SetPageSize(v int32) *DescribeInstanceAutoRenewalAttributeRequest
	GetPageSize() *int32
	SetResourceType(v string) *DescribeInstanceAutoRenewalAttributeRequest
	GetResourceType() *string
	SetResourceTypeList(v []*string) *DescribeInstanceAutoRenewalAttributeRequest
	GetResourceTypeList() []*string
	SetSearchRegionId(v string) *DescribeInstanceAutoRenewalAttributeRequest
	GetSearchRegionId() *string
}

type DescribeInstanceAutoRenewalAttributeRequest struct {
	// The auto-renewal status. Valid values:
	//
	// - true: Auto-renewal is enabled.
	//
	// - false: Auto-renewal is disabled.
	AutoRenewal *bool `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	// The list of instance IDs for which to query the auto-renewal status.
	InstanceIdList []*string `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty" type:"Repeated"`
	// The maximum number of entries to return per query.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The mark for paging. Leave this parameter empty for the first query. For subsequent paging, pass the NextToken value returned in the previous response as is. Do not construct or decode the value.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6k=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The maximum number of entries per page. Maximum value: 500. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The resource type. Example: CorePackage indicates a core-hour package.
	//
	// example:
	//
	// CorePackage
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The list of resource types.
	ResourceTypeList []*string `json:"ResourceTypeList,omitempty" xml:"ResourceTypeList,omitempty" type:"Repeated"`
	// The region ID used to filter instances.
	//
	// example:
	//
	// cn-hangzhou
	SearchRegionId *string `json:"SearchRegionId,omitempty" xml:"SearchRegionId,omitempty"`
}

func (s DescribeInstanceAutoRenewalAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAutoRenewalAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) GetAutoRenewal() *bool {
	return s.AutoRenewal
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) GetInstanceIdList() []*string {
	return s.InstanceIdList
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) GetResourceTypeList() []*string {
	return s.ResourceTypeList
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) GetSearchRegionId() *string {
	return s.SearchRegionId
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) SetAutoRenewal(v bool) *DescribeInstanceAutoRenewalAttributeRequest {
	s.AutoRenewal = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) SetInstanceIdList(v []*string) *DescribeInstanceAutoRenewalAttributeRequest {
	s.InstanceIdList = v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) SetMaxResults(v int32) *DescribeInstanceAutoRenewalAttributeRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) SetNextToken(v string) *DescribeInstanceAutoRenewalAttributeRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) SetPageNo(v int32) *DescribeInstanceAutoRenewalAttributeRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) SetPageSize(v int32) *DescribeInstanceAutoRenewalAttributeRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) SetResourceType(v string) *DescribeInstanceAutoRenewalAttributeRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) SetResourceTypeList(v []*string) *DescribeInstanceAutoRenewalAttributeRequest {
	s.ResourceTypeList = v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) SetSearchRegionId(v string) *DescribeInstanceAutoRenewalAttributeRequest {
	s.SearchRegionId = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) Validate() error {
	return dara.Validate(s)
}
