// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIcebergNamespaceDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListIcebergNamespaceDetailsRequest
	GetMaxResults() *int32
	SetNamespaceNamePattern(v string) *ListIcebergNamespaceDetailsRequest
	GetNamespaceNamePattern() *string
	SetPageToken(v string) *ListIcebergNamespaceDetailsRequest
	GetPageToken() *string
}

type ListIcebergNamespaceDetailsRequest struct {
	// example:
	//
	// 1000
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// namespace%
	NamespaceNamePattern *string `json:"namespaceNamePattern,omitempty" xml:"namespaceNamePattern,omitempty"`
	// example:
	//
	// ""
	PageToken *string `json:"pageToken,omitempty" xml:"pageToken,omitempty"`
}

func (s ListIcebergNamespaceDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIcebergNamespaceDetailsRequest) GoString() string {
	return s.String()
}

func (s *ListIcebergNamespaceDetailsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListIcebergNamespaceDetailsRequest) GetNamespaceNamePattern() *string {
	return s.NamespaceNamePattern
}

func (s *ListIcebergNamespaceDetailsRequest) GetPageToken() *string {
	return s.PageToken
}

func (s *ListIcebergNamespaceDetailsRequest) SetMaxResults(v int32) *ListIcebergNamespaceDetailsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListIcebergNamespaceDetailsRequest) SetNamespaceNamePattern(v string) *ListIcebergNamespaceDetailsRequest {
	s.NamespaceNamePattern = &v
	return s
}

func (s *ListIcebergNamespaceDetailsRequest) SetPageToken(v string) *ListIcebergNamespaceDetailsRequest {
	s.PageToken = &v
	return s
}

func (s *ListIcebergNamespaceDetailsRequest) Validate() error {
	return dara.Validate(s)
}
