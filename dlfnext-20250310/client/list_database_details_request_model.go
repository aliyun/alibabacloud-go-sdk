// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatabaseDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseNamePattern(v string) *ListDatabaseDetailsRequest
	GetDatabaseNamePattern() *string
	SetMaxResults(v int32) *ListDatabaseDetailsRequest
	GetMaxResults() *int32
	SetPageToken(v string) *ListDatabaseDetailsRequest
	GetPageToken() *string
}

type ListDatabaseDetailsRequest struct {
	// The database name pattern for fuzzy matching. Supports the percent sign (%).
	//
	// example:
	//
	// database%
	DatabaseNamePattern *string `json:"databaseNamePattern,omitempty" xml:"databaseNamePattern,omitempty"`
	// The number of entries to return on each page.
	//
	// Default value: 1000.
	//
	// Maximum value: 1000.
	//
	// example:
	//
	// 1000
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The token to retrieve the next page of results. If the response does not include this token, pass an empty string ("").
	//
	// example:
	//
	// ""
	PageToken *string `json:"pageToken,omitempty" xml:"pageToken,omitempty"`
}

func (s ListDatabaseDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDatabaseDetailsRequest) GoString() string {
	return s.String()
}

func (s *ListDatabaseDetailsRequest) GetDatabaseNamePattern() *string {
	return s.DatabaseNamePattern
}

func (s *ListDatabaseDetailsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDatabaseDetailsRequest) GetPageToken() *string {
	return s.PageToken
}

func (s *ListDatabaseDetailsRequest) SetDatabaseNamePattern(v string) *ListDatabaseDetailsRequest {
	s.DatabaseNamePattern = &v
	return s
}

func (s *ListDatabaseDetailsRequest) SetMaxResults(v int32) *ListDatabaseDetailsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDatabaseDetailsRequest) SetPageToken(v string) *ListDatabaseDetailsRequest {
	s.PageToken = &v
	return s
}

func (s *ListDatabaseDetailsRequest) Validate() error {
	return dara.Validate(s)
}
