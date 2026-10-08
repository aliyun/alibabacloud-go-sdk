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
	// The maximum number of datasets to return. Valid values: 0 to 200. If you do not specify this parameter or set it to 0, the default value 100 is used.
	//
	// example:
	//
	// 1
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token.
	//
	// If the total number of datasets exceeds the value of MaxResults, this token is used for pagination. The list of dataset information is returned in lexicographical order starting from NextToken.
	//
	// > When you call this operation for the first time in a query, leave this parameter empty.
	//
	// example:
	//
	// 12345678:immtest:dataset002
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The prefix of the dataset name.
	//
	// example:
	//
	// dataset
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	// The name of the project. For more information about how to obtain the project name, see [Create a project](https://help.aliyun.com/document_detail/478153.html).
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
