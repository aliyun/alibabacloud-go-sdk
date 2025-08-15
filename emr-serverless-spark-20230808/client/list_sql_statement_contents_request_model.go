// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSqlStatementContentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *ListSqlStatementContentsRequest
	GetFileName() *string
	SetMaxResults(v int32) *ListSqlStatementContentsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListSqlStatementContentsRequest
	GetNextToken() *string
}

type ListSqlStatementContentsRequest struct {
	// example:
	//
	// oss://oss-****.cn-hangzhou.oss-dls.aliyuncs.com/w-86a9a4da*****a1/spark/logs/jr-b737b****6164d/driver/st-afde7******bb3f
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// example:
	//
	// 2000
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 1
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListSqlStatementContentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSqlStatementContentsRequest) GoString() string {
	return s.String()
}

func (s *ListSqlStatementContentsRequest) GetFileName() *string {
	return s.FileName
}

func (s *ListSqlStatementContentsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSqlStatementContentsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSqlStatementContentsRequest) SetFileName(v string) *ListSqlStatementContentsRequest {
	s.FileName = &v
	return s
}

func (s *ListSqlStatementContentsRequest) SetMaxResults(v int32) *ListSqlStatementContentsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSqlStatementContentsRequest) SetNextToken(v string) *ListSqlStatementContentsRequest {
	s.NextToken = &v
	return s
}

func (s *ListSqlStatementContentsRequest) Validate() error {
	return dara.Validate(s)
}
