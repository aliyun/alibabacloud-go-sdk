// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTablesRequest
	GetMaxResults() *int32
	SetPageToken(v string) *ListTablesRequest
	GetPageToken() *string
	SetTableNamePattern(v string) *ListTablesRequest
	GetTableNamePattern() *string
}

type ListTablesRequest struct {
	// example:
	//
	// 1000
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// ""
	PageToken        *string `json:"pageToken,omitempty" xml:"pageToken,omitempty"`
	TableNamePattern *string `json:"tableNamePattern,omitempty" xml:"tableNamePattern,omitempty"`
}

func (s ListTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTablesRequest) GoString() string {
	return s.String()
}

func (s *ListTablesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTablesRequest) GetPageToken() *string {
	return s.PageToken
}

func (s *ListTablesRequest) GetTableNamePattern() *string {
	return s.TableNamePattern
}

func (s *ListTablesRequest) SetMaxResults(v int32) *ListTablesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTablesRequest) SetPageToken(v string) *ListTablesRequest {
	s.PageToken = &v
	return s
}

func (s *ListTablesRequest) SetTableNamePattern(v string) *ListTablesRequest {
	s.TableNamePattern = &v
	return s
}

func (s *ListTablesRequest) Validate() error {
	return dara.Validate(s)
}
