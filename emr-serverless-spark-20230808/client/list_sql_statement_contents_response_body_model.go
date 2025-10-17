// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSqlStatementContentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListSqlStatementContentsResponseBody
	GetRequestId() *string
	SetSqlStatementContents(v *ListSqlStatementContentsResponseBodySqlStatementContents) *ListSqlStatementContentsResponseBody
	GetSqlStatementContents() *ListSqlStatementContentsResponseBodySqlStatementContents
}

type ListSqlStatementContentsResponseBody struct {
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId            *string                                                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	SqlStatementContents *ListSqlStatementContentsResponseBodySqlStatementContents `json:"sqlStatementContents,omitempty" xml:"sqlStatementContents,omitempty" type:"Struct"`
}

func (s ListSqlStatementContentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSqlStatementContentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSqlStatementContentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSqlStatementContentsResponseBody) GetSqlStatementContents() *ListSqlStatementContentsResponseBodySqlStatementContents {
	return s.SqlStatementContents
}

func (s *ListSqlStatementContentsResponseBody) SetRequestId(v string) *ListSqlStatementContentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSqlStatementContentsResponseBody) SetSqlStatementContents(v *ListSqlStatementContentsResponseBodySqlStatementContents) *ListSqlStatementContentsResponseBody {
	s.SqlStatementContents = v
	return s
}

func (s *ListSqlStatementContentsResponseBody) Validate() error {
	if s.SqlStatementContents != nil {
		if err := s.SqlStatementContents.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSqlStatementContentsResponseBodySqlStatementContents struct {
	// example:
	//
	// [{\\"values\\":[\\"wj*****\\",\\"test\\",\\"2025-04-15\\"]}]
	Contents *string `json:"contents,omitempty" xml:"contents,omitempty"`
	// example:
	//
	// 2000
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 1
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 40000
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListSqlStatementContentsResponseBodySqlStatementContents) String() string {
	return dara.Prettify(s)
}

func (s ListSqlStatementContentsResponseBodySqlStatementContents) GoString() string {
	return s.String()
}

func (s *ListSqlStatementContentsResponseBodySqlStatementContents) GetContents() *string {
	return s.Contents
}

func (s *ListSqlStatementContentsResponseBodySqlStatementContents) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSqlStatementContentsResponseBodySqlStatementContents) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSqlStatementContentsResponseBodySqlStatementContents) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSqlStatementContentsResponseBodySqlStatementContents) SetContents(v string) *ListSqlStatementContentsResponseBodySqlStatementContents {
	s.Contents = &v
	return s
}

func (s *ListSqlStatementContentsResponseBodySqlStatementContents) SetMaxResults(v int32) *ListSqlStatementContentsResponseBodySqlStatementContents {
	s.MaxResults = &v
	return s
}

func (s *ListSqlStatementContentsResponseBodySqlStatementContents) SetNextToken(v string) *ListSqlStatementContentsResponseBodySqlStatementContents {
	s.NextToken = &v
	return s
}

func (s *ListSqlStatementContentsResponseBodySqlStatementContents) SetTotalCount(v int32) *ListSqlStatementContentsResponseBodySqlStatementContents {
	s.TotalCount = &v
	return s
}

func (s *ListSqlStatementContentsResponseBodySqlStatementContents) Validate() error {
	return dara.Validate(s)
}
