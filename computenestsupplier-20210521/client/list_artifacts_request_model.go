// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListArtifactsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*ListArtifactsRequestFilter) *ListArtifactsRequest
	GetFilter() []*ListArtifactsRequestFilter
	SetMaxResults(v int32) *ListArtifactsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListArtifactsRequest
	GetNextToken() *string
	SetResourceGroupId(v string) *ListArtifactsRequest
	GetResourceGroupId() *string
	SetTag(v []*ListArtifactsRequestTag) *ListArtifactsRequest
	GetTag() []*ListArtifactsRequestTag
}

type ListArtifactsRequest struct {
	// The filter.
	Filter []*ListArtifactsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// AAAAAc3HCuYhJi/wvpk4xOr0VLbfVwapgMwCN1wYzPVzLbItEdB0uWSY7AGnM3qCgm/YnjuEfwSnMwiMkcUoI0hRQzE=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek25yfxxxxxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	Tag []*ListArtifactsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListArtifactsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListArtifactsRequest) GoString() string {
	return s.String()
}

func (s *ListArtifactsRequest) GetFilter() []*ListArtifactsRequestFilter {
	return s.Filter
}

func (s *ListArtifactsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListArtifactsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListArtifactsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListArtifactsRequest) GetTag() []*ListArtifactsRequestTag {
	return s.Tag
}

func (s *ListArtifactsRequest) SetFilter(v []*ListArtifactsRequestFilter) *ListArtifactsRequest {
	s.Filter = v
	return s
}

func (s *ListArtifactsRequest) SetMaxResults(v int32) *ListArtifactsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListArtifactsRequest) SetNextToken(v string) *ListArtifactsRequest {
	s.NextToken = &v
	return s
}

func (s *ListArtifactsRequest) SetResourceGroupId(v string) *ListArtifactsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListArtifactsRequest) SetTag(v []*ListArtifactsRequestTag) *ListArtifactsRequest {
	s.Tag = v
	return s
}

func (s *ListArtifactsRequest) Validate() error {
	return dara.Validate(s)
}

type ListArtifactsRequestFilter struct {
	// The parameter name of the filter. You can specify one or more filters. Valid values:
	//
	// 	- *Name*: The name of the deployment package. Fuzzy match is used.
	//
	// 	- ArtifactId: The ID of the deployment package.
	//
	// 	- ArtifactType: The type of the deployment package.
	//
	// example:
	//
	// ArtifactType
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameter values of the filter.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListArtifactsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListArtifactsRequestFilter) GoString() string {
	return s.String()
}

func (s *ListArtifactsRequestFilter) GetName() *string {
	return s.Name
}

func (s *ListArtifactsRequestFilter) GetValues() []*string {
	return s.Values
}

func (s *ListArtifactsRequestFilter) SetName(v string) *ListArtifactsRequestFilter {
	s.Name = &v
	return s
}

func (s *ListArtifactsRequestFilter) SetValues(v []*string) *ListArtifactsRequestFilter {
	s.Values = v
	return s
}

func (s *ListArtifactsRequestFilter) Validate() error {
	return dara.Validate(s)
}

type ListArtifactsRequestTag struct {
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

func (s ListArtifactsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListArtifactsRequestTag) GoString() string {
	return s.String()
}

func (s *ListArtifactsRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListArtifactsRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListArtifactsRequestTag) SetKey(v string) *ListArtifactsRequestTag {
	s.Key = &v
	return s
}

func (s *ListArtifactsRequestTag) SetValue(v string) *ListArtifactsRequestTag {
	s.Value = &v
	return s
}

func (s *ListArtifactsRequestTag) Validate() error {
	return dara.Validate(s)
}
