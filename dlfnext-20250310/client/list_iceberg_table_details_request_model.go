// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIcebergTableDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListIcebergTableDetailsRequest
	GetMaxResults() *int32
	SetPageToken(v string) *ListIcebergTableDetailsRequest
	GetPageToken() *string
	SetTableNamePattern(v string) *ListIcebergTableDetailsRequest
	GetTableNamePattern() *string
}

type ListIcebergTableDetailsRequest struct {
	// example:
	//
	// 1000
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// ""
	PageToken *string `json:"pageToken,omitempty" xml:"pageToken,omitempty"`
	// example:
	//
	// table%
	TableNamePattern *string `json:"tableNamePattern,omitempty" xml:"tableNamePattern,omitempty"`
}

func (s ListIcebergTableDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIcebergTableDetailsRequest) GoString() string {
	return s.String()
}

func (s *ListIcebergTableDetailsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListIcebergTableDetailsRequest) GetPageToken() *string {
	return s.PageToken
}

func (s *ListIcebergTableDetailsRequest) GetTableNamePattern() *string {
	return s.TableNamePattern
}

func (s *ListIcebergTableDetailsRequest) SetMaxResults(v int32) *ListIcebergTableDetailsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListIcebergTableDetailsRequest) SetPageToken(v string) *ListIcebergTableDetailsRequest {
	s.PageToken = &v
	return s
}

func (s *ListIcebergTableDetailsRequest) SetTableNamePattern(v string) *ListIcebergTableDetailsRequest {
	s.TableNamePattern = &v
	return s
}

func (s *ListIcebergTableDetailsRequest) Validate() error {
	return dara.Validate(s)
}
