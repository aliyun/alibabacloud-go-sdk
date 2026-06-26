// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTableDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTableDetailsRequest
	GetMaxResults() *int32
	SetPageToken(v string) *ListTableDetailsRequest
	GetPageToken() *string
	SetTableNamePattern(v string) *ListTableDetailsRequest
	GetTableNamePattern() *string
	SetType(v string) *ListTableDetailsRequest
	GetType() *string
}

type ListTableDetailsRequest struct {
	// The maximum number of records to return in a single request.
	//
	// example:
	//
	// 1000
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The token to retrieve the next page of results. Pass the token that was returned by the previous request. For the first request, pass an empty string ("").
	//
	// example:
	//
	// ""
	PageToken *string `json:"pageToken,omitempty" xml:"pageToken,omitempty"`
	// The pattern used to filter table names.
	//
	// example:
	//
	// table%
	TableNamePattern *string `json:"tableNamePattern,omitempty" xml:"tableNamePattern,omitempty"`
	// The type.
	//
	// example:
	//
	// table
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListTableDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTableDetailsRequest) GoString() string {
	return s.String()
}

func (s *ListTableDetailsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTableDetailsRequest) GetPageToken() *string {
	return s.PageToken
}

func (s *ListTableDetailsRequest) GetTableNamePattern() *string {
	return s.TableNamePattern
}

func (s *ListTableDetailsRequest) GetType() *string {
	return s.Type
}

func (s *ListTableDetailsRequest) SetMaxResults(v int32) *ListTableDetailsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTableDetailsRequest) SetPageToken(v string) *ListTableDetailsRequest {
	s.PageToken = &v
	return s
}

func (s *ListTableDetailsRequest) SetTableNamePattern(v string) *ListTableDetailsRequest {
	s.TableNamePattern = &v
	return s
}

func (s *ListTableDetailsRequest) SetType(v string) *ListTableDetailsRequest {
	s.Type = &v
	return s
}

func (s *ListTableDetailsRequest) Validate() error {
	return dara.Validate(s)
}
