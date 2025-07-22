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
}

type ListTableDetailsRequest struct {
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

func (s *ListTableDetailsRequest) Validate() error {
	return dara.Validate(s)
}
