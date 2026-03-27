// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrometheusInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilterRegionIds(v string) *ListPrometheusInstancesShrinkRequest
	GetFilterRegionIds() *string
	SetMaxResults(v int32) *ListPrometheusInstancesShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListPrometheusInstancesShrinkRequest
	GetNextToken() *string
	SetPrometheusInstanceIds(v string) *ListPrometheusInstancesShrinkRequest
	GetPrometheusInstanceIds() *string
	SetPrometheusInstanceName(v string) *ListPrometheusInstancesShrinkRequest
	GetPrometheusInstanceName() *string
	SetResourceGroupId(v string) *ListPrometheusInstancesShrinkRequest
	GetResourceGroupId() *string
	SetResourceType(v string) *ListPrometheusInstancesShrinkRequest
	GetResourceType() *string
	SetTagShrink(v string) *ListPrometheusInstancesShrinkRequest
	GetTagShrink() *string
	SetVersion(v string) *ListPrometheusInstancesShrinkRequest
	GetVersion() *string
	SetWorkspace(v string) *ListPrometheusInstancesShrinkRequest
	GetWorkspace() *string
}

type ListPrometheusInstancesShrinkRequest struct {
	// Specified list of regionIds to filter (comma-separated).
	//
	// if can be null:
	// true
	//
	// example:
	//
	// cn-shenzhen
	FilterRegionIds *string `json:"filterRegionIds,omitempty" xml:"filterRegionIds,omitempty"`
	// Maximum number of records to return.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// Query token.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// xxxxxxxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// List of instance IDs (comma-separated)
	//
	// if can be null:
	// true
	//
	// example:
	//
	// rw-00001,rw-00002,rw-00003
	PrometheusInstanceIds *string `json:"prometheusInstanceIds,omitempty" xml:"prometheusInstanceIds,omitempty"`
	// Instance name (partial match supported)
	//
	// if can be null:
	// true
	//
	// example:
	//
	// test
	PrometheusInstanceName *string `json:"prometheusInstanceName,omitempty" xml:"prometheusInstanceName,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-aek2bhocin5e2na
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// Resource type of the instance.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// Prometheus
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// List of tags.
	TagShrink *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// Instance version: V1 or V2
	//
	// if can be null:
	// true
	//
	// example:
	//
	// V2
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// if can be null:
	// true
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListPrometheusInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListPrometheusInstancesShrinkRequest) GetFilterRegionIds() *string {
	return s.FilterRegionIds
}

func (s *ListPrometheusInstancesShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPrometheusInstancesShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPrometheusInstancesShrinkRequest) GetPrometheusInstanceIds() *string {
	return s.PrometheusInstanceIds
}

func (s *ListPrometheusInstancesShrinkRequest) GetPrometheusInstanceName() *string {
	return s.PrometheusInstanceName
}

func (s *ListPrometheusInstancesShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListPrometheusInstancesShrinkRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListPrometheusInstancesShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *ListPrometheusInstancesShrinkRequest) GetVersion() *string {
	return s.Version
}

func (s *ListPrometheusInstancesShrinkRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListPrometheusInstancesShrinkRequest) SetFilterRegionIds(v string) *ListPrometheusInstancesShrinkRequest {
	s.FilterRegionIds = &v
	return s
}

func (s *ListPrometheusInstancesShrinkRequest) SetMaxResults(v int32) *ListPrometheusInstancesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPrometheusInstancesShrinkRequest) SetNextToken(v string) *ListPrometheusInstancesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListPrometheusInstancesShrinkRequest) SetPrometheusInstanceIds(v string) *ListPrometheusInstancesShrinkRequest {
	s.PrometheusInstanceIds = &v
	return s
}

func (s *ListPrometheusInstancesShrinkRequest) SetPrometheusInstanceName(v string) *ListPrometheusInstancesShrinkRequest {
	s.PrometheusInstanceName = &v
	return s
}

func (s *ListPrometheusInstancesShrinkRequest) SetResourceGroupId(v string) *ListPrometheusInstancesShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListPrometheusInstancesShrinkRequest) SetResourceType(v string) *ListPrometheusInstancesShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *ListPrometheusInstancesShrinkRequest) SetTagShrink(v string) *ListPrometheusInstancesShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *ListPrometheusInstancesShrinkRequest) SetVersion(v string) *ListPrometheusInstancesShrinkRequest {
	s.Version = &v
	return s
}

func (s *ListPrometheusInstancesShrinkRequest) SetWorkspace(v string) *ListPrometheusInstancesShrinkRequest {
	s.Workspace = &v
	return s
}

func (s *ListPrometheusInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
