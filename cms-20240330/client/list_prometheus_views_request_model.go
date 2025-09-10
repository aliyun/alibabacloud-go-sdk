// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrometheusViewsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilterRegionIds(v string) *ListPrometheusViewsRequest
	GetFilterRegionIds() *string
	SetMaxResults(v int32) *ListPrometheusViewsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListPrometheusViewsRequest
	GetNextToken() *string
	SetPrometheusViewIds(v string) *ListPrometheusViewsRequest
	GetPrometheusViewIds() *string
	SetPrometheusViewName(v string) *ListPrometheusViewsRequest
	GetPrometheusViewName() *string
	SetResourceGroupId(v string) *ListPrometheusViewsRequest
	GetResourceGroupId() *string
	SetResourceType(v string) *ListPrometheusViewsRequest
	GetResourceType() *string
	SetTag(v []*ListPrometheusViewsRequestTag) *ListPrometheusViewsRequest
	GetTag() []*ListPrometheusViewsRequestTag
	SetVersion(v string) *ListPrometheusViewsRequest
	GetVersion() *string
	SetWorkspace(v string) *ListPrometheusViewsRequest
	GetWorkspace() *string
}

type ListPrometheusViewsRequest struct {
	// example:
	//
	// cn-zhangjiakou,cn-beijing
	FilterRegionIds *string `json:"filterRegionIds,omitempty" xml:"filterRegionIds,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// 44ANBjKZmQeKnaB1fXRq06w7sFYK3MUcCALMD9qQbmEiE
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// view-xxx
	PrometheusViewIds *string `json:"prometheusViewIds,omitempty" xml:"prometheusViewIds,omitempty"`
	// example:
	//
	// view1
	PrometheusViewName *string `json:"prometheusViewName,omitempty" xml:"prometheusViewName,omitempty"`
	// example:
	//
	// rg-acfm3gn5i6bigbi
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// example:
	//
	// PROMETHEUSVIEW
	ResourceType *string                          `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	Tag          []*ListPrometheusViewsRequestTag `json:"tag,omitempty" xml:"tag,omitempty" type:"Repeated"`
	// example:
	//
	// V2
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// example:
	//
	// workspace-test
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListPrometheusViewsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusViewsRequest) GoString() string {
	return s.String()
}

func (s *ListPrometheusViewsRequest) GetFilterRegionIds() *string {
	return s.FilterRegionIds
}

func (s *ListPrometheusViewsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPrometheusViewsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPrometheusViewsRequest) GetPrometheusViewIds() *string {
	return s.PrometheusViewIds
}

func (s *ListPrometheusViewsRequest) GetPrometheusViewName() *string {
	return s.PrometheusViewName
}

func (s *ListPrometheusViewsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListPrometheusViewsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListPrometheusViewsRequest) GetTag() []*ListPrometheusViewsRequestTag {
	return s.Tag
}

func (s *ListPrometheusViewsRequest) GetVersion() *string {
	return s.Version
}

func (s *ListPrometheusViewsRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListPrometheusViewsRequest) SetFilterRegionIds(v string) *ListPrometheusViewsRequest {
	s.FilterRegionIds = &v
	return s
}

func (s *ListPrometheusViewsRequest) SetMaxResults(v int32) *ListPrometheusViewsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPrometheusViewsRequest) SetNextToken(v string) *ListPrometheusViewsRequest {
	s.NextToken = &v
	return s
}

func (s *ListPrometheusViewsRequest) SetPrometheusViewIds(v string) *ListPrometheusViewsRequest {
	s.PrometheusViewIds = &v
	return s
}

func (s *ListPrometheusViewsRequest) SetPrometheusViewName(v string) *ListPrometheusViewsRequest {
	s.PrometheusViewName = &v
	return s
}

func (s *ListPrometheusViewsRequest) SetResourceGroupId(v string) *ListPrometheusViewsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListPrometheusViewsRequest) SetResourceType(v string) *ListPrometheusViewsRequest {
	s.ResourceType = &v
	return s
}

func (s *ListPrometheusViewsRequest) SetTag(v []*ListPrometheusViewsRequestTag) *ListPrometheusViewsRequest {
	s.Tag = v
	return s
}

func (s *ListPrometheusViewsRequest) SetVersion(v string) *ListPrometheusViewsRequest {
	s.Version = &v
	return s
}

func (s *ListPrometheusViewsRequest) SetWorkspace(v string) *ListPrometheusViewsRequest {
	s.Workspace = &v
	return s
}

func (s *ListPrometheusViewsRequest) Validate() error {
	return dara.Validate(s)
}

type ListPrometheusViewsRequestTag struct {
	// example:
	//
	// key1
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// value1
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListPrometheusViewsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusViewsRequestTag) GoString() string {
	return s.String()
}

func (s *ListPrometheusViewsRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListPrometheusViewsRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListPrometheusViewsRequestTag) SetKey(v string) *ListPrometheusViewsRequestTag {
	s.Key = &v
	return s
}

func (s *ListPrometheusViewsRequestTag) SetValue(v string) *ListPrometheusViewsRequestTag {
	s.Value = &v
	return s
}

func (s *ListPrometheusViewsRequestTag) Validate() error {
	return dara.Validate(s)
}
