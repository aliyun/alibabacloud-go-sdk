// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccountFactoryBaselineItemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListAccountFactoryBaselineItemsRequest
	GetMaxResults() *int32
	SetNames(v []*string) *ListAccountFactoryBaselineItemsRequest
	GetNames() []*string
	SetNextToken(v string) *ListAccountFactoryBaselineItemsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListAccountFactoryBaselineItemsRequest
	GetRegionId() *string
	SetType(v string) *ListAccountFactoryBaselineItemsRequest
	GetType() *string
	SetVersions(v []*string) *ListAccountFactoryBaselineItemsRequest
	GetVersions() []*string
}

type ListAccountFactoryBaselineItemsRequest struct {
	// The maximum number of entries per page.
	//
	// Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The names of the baseline items.
	Names []*string `json:"Names,omitempty" xml:"Names,omitempty" type:"Repeated"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request.
	//
	// example:
	//
	// AAAAACDGQdAEX3m42z3sQ+f3VTK2Xr2DzYbz/SAfc/zJRqod
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the baseline items.
	//
	// example:
	//
	// AccountFactory
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The versions of the baseline items.
	Versions []*string `json:"Versions,omitempty" xml:"Versions,omitempty" type:"Repeated"`
}

func (s ListAccountFactoryBaselineItemsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAccountFactoryBaselineItemsRequest) GoString() string {
	return s.String()
}

func (s *ListAccountFactoryBaselineItemsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAccountFactoryBaselineItemsRequest) GetNames() []*string {
	return s.Names
}

func (s *ListAccountFactoryBaselineItemsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAccountFactoryBaselineItemsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAccountFactoryBaselineItemsRequest) GetType() *string {
	return s.Type
}

func (s *ListAccountFactoryBaselineItemsRequest) GetVersions() []*string {
	return s.Versions
}

func (s *ListAccountFactoryBaselineItemsRequest) SetMaxResults(v int32) *ListAccountFactoryBaselineItemsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAccountFactoryBaselineItemsRequest) SetNames(v []*string) *ListAccountFactoryBaselineItemsRequest {
	s.Names = v
	return s
}

func (s *ListAccountFactoryBaselineItemsRequest) SetNextToken(v string) *ListAccountFactoryBaselineItemsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAccountFactoryBaselineItemsRequest) SetRegionId(v string) *ListAccountFactoryBaselineItemsRequest {
	s.RegionId = &v
	return s
}

func (s *ListAccountFactoryBaselineItemsRequest) SetType(v string) *ListAccountFactoryBaselineItemsRequest {
	s.Type = &v
	return s
}

func (s *ListAccountFactoryBaselineItemsRequest) SetVersions(v []*string) *ListAccountFactoryBaselineItemsRequest {
	s.Versions = v
	return s
}

func (s *ListAccountFactoryBaselineItemsRequest) Validate() error {
	return dara.Validate(s)
}
