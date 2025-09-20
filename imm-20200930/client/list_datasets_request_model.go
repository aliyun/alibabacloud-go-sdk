// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int64) *ListDatasetsRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListDatasetsRequest
	GetNextToken() *string
	SetPrefix(v string) *ListDatasetsRequest
	GetPrefix() *string
	SetProjectName(v string) *ListDatasetsRequest
	GetProjectName() *string
}

type ListDatasetsRequest struct {
	// The maximum number of datasets to return. Valid values: 0 to 200.
	//
	// If this parameter is left empty or set to 0, 100 datasets are returned.
	//
	// example:
	//
	// 1
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token.
	//
	// If the total number of datasets is greater than the value of MaxResults, you must specify this parameter. The list is returned in lexicographic order starting from the value of NextToken.
	//
	// >  The first time you call this operation in a query, set this parameter to null.
	//
	// example:
	//
	// 12345678:immtest:dataset002
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The dataset prefix.
	//
	// example:
	//
	// dataset
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	// The name of the project. For more information, see [CreateProject](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s ListDatasetsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetsRequest) GoString() string {
	return s.String()
}

func (s *ListDatasetsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListDatasetsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDatasetsRequest) GetPrefix() *string {
	return s.Prefix
}

func (s *ListDatasetsRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListDatasetsRequest) SetMaxResults(v int64) *ListDatasetsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDatasetsRequest) SetNextToken(v string) *ListDatasetsRequest {
	s.NextToken = &v
	return s
}

func (s *ListDatasetsRequest) SetPrefix(v string) *ListDatasetsRequest {
	s.Prefix = &v
	return s
}

func (s *ListDatasetsRequest) SetProjectName(v string) *ListDatasetsRequest {
	s.ProjectName = &v
	return s
}

func (s *ListDatasetsRequest) Validate() error {
	return dara.Validate(s)
}
