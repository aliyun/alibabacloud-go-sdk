// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListArtifactBuildLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactId(v string) *ListArtifactBuildLogsRequest
	GetArtifactId() *string
	SetArtifactVersion(v string) *ListArtifactBuildLogsRequest
	GetArtifactVersion() *string
	SetFilter(v []*ListArtifactBuildLogsRequestFilter) *ListArtifactBuildLogsRequest
	GetFilter() []*ListArtifactBuildLogsRequestFilter
	SetMaxResults(v int32) *ListArtifactBuildLogsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListArtifactBuildLogsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListArtifactBuildLogsRequest
	GetRegionId() *string
	SetSortOrder(v string) *ListArtifactBuildLogsRequest
	GetSortOrder() *string
}

type ListArtifactBuildLogsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// artifact-fbad2ca276194d019714
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// example:
	//
	// draft
	ArtifactVersion *string                               `json:"ArtifactVersion,omitempty" xml:"ArtifactVersion,omitempty"`
	Filter          []*ListArtifactBuildLogsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAbL3H6CZmy6oocwGDqzQ+Gc=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// Ascending
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s ListArtifactBuildLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListArtifactBuildLogsRequest) GoString() string {
	return s.String()
}

func (s *ListArtifactBuildLogsRequest) GetArtifactId() *string {
	return s.ArtifactId
}

func (s *ListArtifactBuildLogsRequest) GetArtifactVersion() *string {
	return s.ArtifactVersion
}

func (s *ListArtifactBuildLogsRequest) GetFilter() []*ListArtifactBuildLogsRequestFilter {
	return s.Filter
}

func (s *ListArtifactBuildLogsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListArtifactBuildLogsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListArtifactBuildLogsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListArtifactBuildLogsRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListArtifactBuildLogsRequest) SetArtifactId(v string) *ListArtifactBuildLogsRequest {
	s.ArtifactId = &v
	return s
}

func (s *ListArtifactBuildLogsRequest) SetArtifactVersion(v string) *ListArtifactBuildLogsRequest {
	s.ArtifactVersion = &v
	return s
}

func (s *ListArtifactBuildLogsRequest) SetFilter(v []*ListArtifactBuildLogsRequestFilter) *ListArtifactBuildLogsRequest {
	s.Filter = v
	return s
}

func (s *ListArtifactBuildLogsRequest) SetMaxResults(v int32) *ListArtifactBuildLogsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListArtifactBuildLogsRequest) SetNextToken(v string) *ListArtifactBuildLogsRequest {
	s.NextToken = &v
	return s
}

func (s *ListArtifactBuildLogsRequest) SetRegionId(v string) *ListArtifactBuildLogsRequest {
	s.RegionId = &v
	return s
}

func (s *ListArtifactBuildLogsRequest) SetSortOrder(v string) *ListArtifactBuildLogsRequest {
	s.SortOrder = &v
	return s
}

func (s *ListArtifactBuildLogsRequest) Validate() error {
	return dara.Validate(s)
}

type ListArtifactBuildLogsRequestFilter struct {
	// example:
	//
	// BuildStartTime
	Name  *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListArtifactBuildLogsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListArtifactBuildLogsRequestFilter) GoString() string {
	return s.String()
}

func (s *ListArtifactBuildLogsRequestFilter) GetName() *string {
	return s.Name
}

func (s *ListArtifactBuildLogsRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *ListArtifactBuildLogsRequestFilter) SetName(v string) *ListArtifactBuildLogsRequestFilter {
	s.Name = &v
	return s
}

func (s *ListArtifactBuildLogsRequestFilter) SetValue(v []*string) *ListArtifactBuildLogsRequestFilter {
	s.Value = v
	return s
}

func (s *ListArtifactBuildLogsRequestFilter) Validate() error {
	return dara.Validate(s)
}
