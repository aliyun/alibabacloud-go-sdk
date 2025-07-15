// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationType(v string) *ListApplicationsRequest
	GetApplicationType() *string
	SetMaxResults(v int32) *ListApplicationsRequest
	GetMaxResults() *int32
	SetName(v string) *ListApplicationsRequest
	GetName() *string
	SetNames(v string) *ListApplicationsRequest
	GetNames() *string
	SetNextToken(v string) *ListApplicationsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListApplicationsRequest
	GetRegionId() *string
	SetTags(v map[string]interface{}) *ListApplicationsRequest
	GetTags() map[string]interface{}
}

type ListApplicationsRequest struct {
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
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListApplicationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationsRequest) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *ListApplicationsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListApplicationsRequest) GetName() *string {
	return s.Name
}

func (s *ListApplicationsRequest) GetNames() *string {
	return s.Names
}

func (s *ListApplicationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListApplicationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListApplicationsRequest) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *ListApplicationsRequest) SetApplicationType(v string) *ListApplicationsRequest {
	s.ApplicationType = &v
	return s
}

func (s *ListApplicationsRequest) SetMaxResults(v int32) *ListApplicationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListApplicationsRequest) SetName(v string) *ListApplicationsRequest {
	s.Name = &v
	return s
}

func (s *ListApplicationsRequest) SetNames(v string) *ListApplicationsRequest {
	s.Names = &v
	return s
}

func (s *ListApplicationsRequest) SetNextToken(v string) *ListApplicationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListApplicationsRequest) SetRegionId(v string) *ListApplicationsRequest {
	s.RegionId = &v
	return s
}

func (s *ListApplicationsRequest) SetTags(v map[string]interface{}) *ListApplicationsRequest {
	s.Tags = v
	return s
}

func (s *ListApplicationsRequest) Validate() error {
	return dara.Validate(s)
}
