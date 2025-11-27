// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAttachedOSSBucketsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int64) *ListAttachedOSSBucketsRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListAttachedOSSBucketsRequest
	GetNextToken() *string
	SetProjectName(v string) *ListAttachedOSSBucketsRequest
	GetProjectName() *string
}

type ListAttachedOSSBucketsRequest struct {
	// The maximum number of tasks in the returned result list. The value range is (0, 200], with a default value of 100.
	//
	// example:
	//
	// 1
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Pagination token.
	//
	// When the total number of files exceeds the set MaxResults, this token is used for pagination. It returns the list of file information in lexicographical order starting from NextToken.
	//
	// > When calling this interface for the first time in a single query, set this value to empty.
	//
	// example:
	//
	// MTIzNDU2Nzg6aW1tdGVzdDpleGFtcGxlYnVja2V0OmRhdGFzZXQwMDE6b3NzOi8vZXhhbXBsZWJ1Y2tldC9zYW1wbGVvYmplY3QxLmpwZw==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Project name, for more information on how to obtain it, see [Create Project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s ListAttachedOSSBucketsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAttachedOSSBucketsRequest) GoString() string {
	return s.String()
}

func (s *ListAttachedOSSBucketsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListAttachedOSSBucketsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAttachedOSSBucketsRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListAttachedOSSBucketsRequest) SetMaxResults(v int64) *ListAttachedOSSBucketsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAttachedOSSBucketsRequest) SetNextToken(v string) *ListAttachedOSSBucketsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAttachedOSSBucketsRequest) SetProjectName(v string) *ListAttachedOSSBucketsRequest {
	s.ProjectName = &v
	return s
}

func (s *ListAttachedOSSBucketsRequest) Validate() error {
	return dara.Validate(s)
}
