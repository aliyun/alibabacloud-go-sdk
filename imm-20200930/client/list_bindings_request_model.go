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
	SetName(v string) *ListBindingsRequest
	GetName() *string
	SetNextToken(v string) *ListBindingsRequest
	GetNextToken() *string
	SetProjectName(v string) *ListBindingsRequest
	GetProjectName() *string
}

type ListBindingsRequest struct {
	// The dataset name. For information about how to obtain the dataset name, see [CreateDataset](https://help.aliyun.com/document_detail/478160.html).
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// - The maximum number of bindings to return. Valid values: 0 to 200.
	//
	// - If this parameter is not set or is set to 0, the default value 100 is used.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The name of the binding task.
	//
	// example:
	//
	// imm
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// - The pagination token that is used when the total number of bindings exceeds the MaxResults value.
	//
	// - Binding information is returned in alphabetical order starting from the NextToken value.
	//
	// - Leave this parameter empty for the first request.
	//
	// example:
	//
	// immtest:dataset001:examplebucket01
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The project name. For information about how to obtain the project name, see [CreateProject](https://help.aliyun.com/document_detail/478153.html).
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

func (s *ListBindingsRequest) GetName() *string {
	return s.Name
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

func (s *ListBindingsRequest) SetName(v string) *ListBindingsRequest {
	s.Name = &v
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
