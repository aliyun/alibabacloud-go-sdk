// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*ListServiceInstancesRequestFilter) *ListServiceInstancesRequest
	GetFilter() []*ListServiceInstancesRequestFilter
	SetMaxResults(v int32) *ListServiceInstancesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListServiceInstancesRequest
	GetNextToken() *string
	SetRegionId(v string) *ListServiceInstancesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListServiceInstancesRequest
	GetResourceGroupId() *string
	SetTag(v []*ListServiceInstancesRequestTag) *ListServiceInstancesRequest
	GetTag() []*ListServiceInstancesRequestTag
}

type ListServiceInstancesRequest struct {
	// The filter.
	Filter []*ListServiceInstancesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The number of entries page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// BBBAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tag key and value.
	Tag []*ListServiceInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListServiceInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesRequest) GetFilter() []*ListServiceInstancesRequestFilter {
	return s.Filter
}

func (s *ListServiceInstancesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServiceInstancesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServiceInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListServiceInstancesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListServiceInstancesRequest) GetTag() []*ListServiceInstancesRequestTag {
	return s.Tag
}

func (s *ListServiceInstancesRequest) SetFilter(v []*ListServiceInstancesRequestFilter) *ListServiceInstancesRequest {
	s.Filter = v
	return s
}

func (s *ListServiceInstancesRequest) SetMaxResults(v int32) *ListServiceInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstancesRequest) SetNextToken(v string) *ListServiceInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstancesRequest) SetRegionId(v string) *ListServiceInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *ListServiceInstancesRequest) SetResourceGroupId(v string) *ListServiceInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListServiceInstancesRequest) SetTag(v []*ListServiceInstancesRequestTag) *ListServiceInstancesRequest {
	s.Tag = v
	return s
}

func (s *ListServiceInstancesRequest) Validate() error {
	return dara.Validate(s)
}

type ListServiceInstancesRequestFilter struct {
	// The parameter name of the filter. You can specify one or more parameter names to query services. Valid values:
	//
	// - Name：Query by service name.
	//
	// - ServiceInstanceName：Query by service  instance name.
	//
	// - ServiceInstanceId：Query by service  instance ID.
	//
	// - ServiceId：Query by service ID.
	//
	// - Version：Query by service version.
	//
	// - Status：Query by service status.
	//
	// - DeployType: Query by service deployType.
	//
	// - ServiceType：Query by service deployType.
	//
	// example:
	//
	// ServiceInstanceId
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameter values of the filter.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListServiceInstancesRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstancesRequestFilter) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesRequestFilter) GetName() *string {
	return s.Name
}

func (s *ListServiceInstancesRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *ListServiceInstancesRequestFilter) SetName(v string) *ListServiceInstancesRequestFilter {
	s.Name = &v
	return s
}

func (s *ListServiceInstancesRequestFilter) SetValue(v []*string) *ListServiceInstancesRequestFilter {
	s.Value = v
	return s
}

func (s *ListServiceInstancesRequestFilter) Validate() error {
	return dara.Validate(s)
}

type ListServiceInstancesRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListServiceInstancesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListServiceInstancesRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListServiceInstancesRequestTag) SetKey(v string) *ListServiceInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *ListServiceInstancesRequestTag) SetValue(v string) *ListServiceInstancesRequestTag {
	s.Value = &v
	return s
}

func (s *ListServiceInstancesRequestTag) Validate() error {
	return dara.Validate(s)
}
