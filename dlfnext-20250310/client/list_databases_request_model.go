// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatabasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseNamePattern(v string) *ListDatabasesRequest
	GetDatabaseNamePattern() *string
	SetMaxResults(v int32) *ListDatabasesRequest
	GetMaxResults() *int32
	SetPageToken(v string) *ListDatabasesRequest
	GetPageToken() *string
}

type ListDatabasesRequest struct {
	// The database name pattern for fuzzy matching.
	//
	// example:
	//
	// database%
	DatabaseNamePattern *string `json:"databaseNamePattern,omitempty" xml:"databaseNamePattern,omitempty"`
	// The maximum number of records to return in a single request.
	//
	// example:
	//
	// 1000
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The token for the next page of results. If the response does not provide a token, pass an empty string ("").
	//
	// example:
	//
	// ""
	PageToken *string `json:"pageToken,omitempty" xml:"pageToken,omitempty"`
}

func (s ListDatabasesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDatabasesRequest) GoString() string {
	return s.String()
}

func (s *ListDatabasesRequest) GetDatabaseNamePattern() *string {
	return s.DatabaseNamePattern
}

func (s *ListDatabasesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDatabasesRequest) GetPageToken() *string {
	return s.PageToken
}

func (s *ListDatabasesRequest) SetDatabaseNamePattern(v string) *ListDatabasesRequest {
	s.DatabaseNamePattern = &v
	return s
}

func (s *ListDatabasesRequest) SetMaxResults(v int32) *ListDatabasesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDatabasesRequest) SetPageToken(v string) *ListDatabasesRequest {
	s.PageToken = &v
	return s
}

func (s *ListDatabasesRequest) Validate() error {
	return dara.Validate(s)
}
