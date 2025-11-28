// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBroadcastTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListBroadcastTemplatesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListBroadcastTemplatesRequest
	GetNextToken() *string
	SetPage(v int32) *ListBroadcastTemplatesRequest
	GetPage() *int32
	SetSize(v int32) *ListBroadcastTemplatesRequest
	GetSize() *int32
}

type ListBroadcastTemplatesRequest struct {
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 20
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListBroadcastTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBroadcastTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListBroadcastTemplatesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListBroadcastTemplatesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListBroadcastTemplatesRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListBroadcastTemplatesRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListBroadcastTemplatesRequest) SetMaxResults(v int32) *ListBroadcastTemplatesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListBroadcastTemplatesRequest) SetNextToken(v string) *ListBroadcastTemplatesRequest {
	s.NextToken = &v
	return s
}

func (s *ListBroadcastTemplatesRequest) SetPage(v int32) *ListBroadcastTemplatesRequest {
	s.Page = &v
	return s
}

func (s *ListBroadcastTemplatesRequest) SetSize(v int32) *ListBroadcastTemplatesRequest {
	s.Size = &v
	return s
}

func (s *ListBroadcastTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
