// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasetVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListDatasetVersionsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDatasetVersionsRequest
	GetNextToken() *string
}

type ListDatasetVersionsRequest struct {
	// The maximum number of versions to return per page. Valid values: 1 to 100. Default value: 100. The actual number of returned versions may be less than this value.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token for the next page. Do not specify this parameter in the first request. To query subsequent pages, pass the nextToken value from the previous response as is, and keep the agent workspace and dataset name unchanged. The example is for demonstration purposes only. Use the actual token returned by the API operation.
	//
	// example:
	//
	// umaQfI7x758Ns4TgWrj8yA3fYlnk7dJgsfhMrSViRY8=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListDatasetVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListDatasetVersionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDatasetVersionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDatasetVersionsRequest) SetMaxResults(v int32) *ListDatasetVersionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDatasetVersionsRequest) SetNextToken(v string) *ListDatasetVersionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListDatasetVersionsRequest) Validate() error {
	return dara.Validate(s)
}
