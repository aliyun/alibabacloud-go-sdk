// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrometheusViewsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilterRegionIds(v string) *ListPrometheusViewsShrinkRequest
	GetFilterRegionIds() *string
	SetMaxResults(v int32) *ListPrometheusViewsShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListPrometheusViewsShrinkRequest
	GetNextToken() *string
	SetPrometheusViewIds(v string) *ListPrometheusViewsShrinkRequest
	GetPrometheusViewIds() *string
	SetPrometheusViewName(v string) *ListPrometheusViewsShrinkRequest
	GetPrometheusViewName() *string
	SetResourceGroupId(v string) *ListPrometheusViewsShrinkRequest
	GetResourceGroupId() *string
	SetResourceType(v string) *ListPrometheusViewsShrinkRequest
	GetResourceType() *string
	SetTagShrink(v string) *ListPrometheusViewsShrinkRequest
	GetTagShrink() *string
	SetVersion(v string) *ListPrometheusViewsShrinkRequest
	GetVersion() *string
	SetWorkspace(v string) *ListPrometheusViewsShrinkRequest
	GetWorkspace() *string
}

type ListPrometheusViewsShrinkRequest struct {
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
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	TagShrink    *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// example:
	//
	// V2
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// example:
	//
	// workspace-test
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListPrometheusViewsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusViewsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListPrometheusViewsShrinkRequest) GetFilterRegionIds() *string {
	return s.FilterRegionIds
}

func (s *ListPrometheusViewsShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPrometheusViewsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPrometheusViewsShrinkRequest) GetPrometheusViewIds() *string {
	return s.PrometheusViewIds
}

func (s *ListPrometheusViewsShrinkRequest) GetPrometheusViewName() *string {
	return s.PrometheusViewName
}

func (s *ListPrometheusViewsShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListPrometheusViewsShrinkRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListPrometheusViewsShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *ListPrometheusViewsShrinkRequest) GetVersion() *string {
	return s.Version
}

func (s *ListPrometheusViewsShrinkRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListPrometheusViewsShrinkRequest) SetFilterRegionIds(v string) *ListPrometheusViewsShrinkRequest {
	s.FilterRegionIds = &v
	return s
}

func (s *ListPrometheusViewsShrinkRequest) SetMaxResults(v int32) *ListPrometheusViewsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPrometheusViewsShrinkRequest) SetNextToken(v string) *ListPrometheusViewsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListPrometheusViewsShrinkRequest) SetPrometheusViewIds(v string) *ListPrometheusViewsShrinkRequest {
	s.PrometheusViewIds = &v
	return s
}

func (s *ListPrometheusViewsShrinkRequest) SetPrometheusViewName(v string) *ListPrometheusViewsShrinkRequest {
	s.PrometheusViewName = &v
	return s
}

func (s *ListPrometheusViewsShrinkRequest) SetResourceGroupId(v string) *ListPrometheusViewsShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListPrometheusViewsShrinkRequest) SetResourceType(v string) *ListPrometheusViewsShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *ListPrometheusViewsShrinkRequest) SetTagShrink(v string) *ListPrometheusViewsShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *ListPrometheusViewsShrinkRequest) SetVersion(v string) *ListPrometheusViewsShrinkRequest {
	s.Version = &v
	return s
}

func (s *ListPrometheusViewsShrinkRequest) SetWorkspace(v string) *ListPrometheusViewsShrinkRequest {
	s.Workspace = &v
	return s
}

func (s *ListPrometheusViewsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
