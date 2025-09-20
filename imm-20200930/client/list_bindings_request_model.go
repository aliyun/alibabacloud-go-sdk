// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBindingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *ListBindingsRequest
	GetDatasetName() *string
	SetMaxResults(v int64) *ListBindingsRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListBindingsRequest
	GetNextToken() *string
	SetProjectName(v string) *ListBindingsRequest
	GetProjectName() *string
}

type ListBindingsRequest struct {
	// The name of the dataset.[](~~478160~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// 	- The maximum number of bindings to return. Valid values: 0 to 200.
	//
	// 	- If you do not specify this parameter or set the parameter to 0, the default value of 100 is used.
	//
	// example:
	//
	// 1
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 	- The pagination token that is used in the next request to retrieve a new page of results if the total number of results exceeds the value of the MaxResults parameter.
	//
	// 	- The next call to the operation returns results lexicographically after the NextToken parameter value.
	//
	// 	- You do not need to specify this parameter in your initial request.
	//
	// example:
	//
	// immtest:dataset001:examplebucket01
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The name of the project.[](~~478153~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s ListBindingsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBindingsRequest) GoString() string {
	return s.String()
}

func (s *ListBindingsRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *ListBindingsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListBindingsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListBindingsRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListBindingsRequest) SetDatasetName(v string) *ListBindingsRequest {
	s.DatasetName = &v
	return s
}

func (s *ListBindingsRequest) SetMaxResults(v int64) *ListBindingsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListBindingsRequest) SetNextToken(v string) *ListBindingsRequest {
	s.NextToken = &v
	return s
}

func (s *ListBindingsRequest) SetProjectName(v string) *ListBindingsRequest {
	s.ProjectName = &v
	return s
}

func (s *ListBindingsRequest) Validate() error {
	return dara.Validate(s)
}
