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
	SetStatus(v string) *ListDatabaseDetailsRequest
	GetStatus() *string
}

type ListDatabaseDetailsRequest struct {
	DatabaseNamePattern *string `json:"databaseNamePattern,omitempty" xml:"databaseNamePattern,omitempty"`
	MaxResults          *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	PageToken           *string `json:"pageToken,omitempty" xml:"pageToken,omitempty"`
	Status              *string `json:"status,omitempty" xml:"status,omitempty"`
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

func (s *ListDatabaseDetailsRequest) GetStatus() *string {
	return s.Status
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

func (s *ListDatabaseDetailsRequest) SetStatus(v string) *ListDatabaseDetailsRequest {
	s.Status = &v
	return s
}

func (s *ListDatabaseDetailsRequest) Validate() error {
	return dara.Validate(s)
}
