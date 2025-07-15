// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationType(v string) *ListApplicationsShrinkRequest
	GetApplicationType() *string
	SetMaxResults(v int32) *ListApplicationsShrinkRequest
	GetMaxResults() *int32
	SetName(v string) *ListApplicationsShrinkRequest
	GetName() *string
	SetNames(v string) *ListApplicationsShrinkRequest
	GetNames() *string
	SetNextToken(v string) *ListApplicationsShrinkRequest
	GetNextToken() *string
	SetRegionId(v string) *ListApplicationsShrinkRequest
	GetRegionId() *string
	SetTagsShrink(v string) *ListApplicationsShrinkRequest
	GetTagsShrink() *string
}

type ListApplicationsShrinkRequest struct {
	// The type of the application.
	//
	// Valid values:
	//
	// 	- ComputeNest
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Custom
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- DingTalk
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// DingTalk
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// The number of entries to return on each page. Valid values: 10 to 100. Default value: 50.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// "MyApplications"
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The names of the applications.
	//
	// example:
	//
	// ["MyApplication"]
	Names *string `json:"Names,omitempty" xml:"Names,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// -
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID. Set the value to cn-hangzhou.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tags.
	//
	// example:
	//
	// {"k1": "v1","k2": "v2"}
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListApplicationsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationsShrinkRequest) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *ListApplicationsShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListApplicationsShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListApplicationsShrinkRequest) GetNames() *string {
	return s.Names
}

func (s *ListApplicationsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListApplicationsShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListApplicationsShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *ListApplicationsShrinkRequest) SetApplicationType(v string) *ListApplicationsShrinkRequest {
	s.ApplicationType = &v
	return s
}

func (s *ListApplicationsShrinkRequest) SetMaxResults(v int32) *ListApplicationsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListApplicationsShrinkRequest) SetName(v string) *ListApplicationsShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListApplicationsShrinkRequest) SetNames(v string) *ListApplicationsShrinkRequest {
	s.Names = &v
	return s
}

func (s *ListApplicationsShrinkRequest) SetNextToken(v string) *ListApplicationsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListApplicationsShrinkRequest) SetRegionId(v string) *ListApplicationsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListApplicationsShrinkRequest) SetTagsShrink(v string) *ListApplicationsShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListApplicationsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
