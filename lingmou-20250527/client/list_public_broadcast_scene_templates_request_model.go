// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublicBroadcastSceneTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListPublicBroadcastSceneTemplatesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListPublicBroadcastSceneTemplatesRequest
	GetNextToken() *string
	SetPage(v int64) *ListPublicBroadcastSceneTemplatesRequest
	GetPage() *int64
	SetSize(v int64) *ListPublicBroadcastSceneTemplatesRequest
	GetSize() *int64
	SetTags(v string) *ListPublicBroadcastSceneTemplatesRequest
	GetTags() *string
}

type ListPublicBroadcastSceneTemplatesRequest struct {
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 20
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// AI,BROADCAST
	Tags *string `json:"tags,omitempty" xml:"tags,omitempty"`
}

func (s ListPublicBroadcastSceneTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPublicBroadcastSceneTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListPublicBroadcastSceneTemplatesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPublicBroadcastSceneTemplatesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPublicBroadcastSceneTemplatesRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListPublicBroadcastSceneTemplatesRequest) GetSize() *int64 {
	return s.Size
}

func (s *ListPublicBroadcastSceneTemplatesRequest) GetTags() *string {
	return s.Tags
}

func (s *ListPublicBroadcastSceneTemplatesRequest) SetMaxResults(v int32) *ListPublicBroadcastSceneTemplatesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPublicBroadcastSceneTemplatesRequest) SetNextToken(v string) *ListPublicBroadcastSceneTemplatesRequest {
	s.NextToken = &v
	return s
}

func (s *ListPublicBroadcastSceneTemplatesRequest) SetPage(v int64) *ListPublicBroadcastSceneTemplatesRequest {
	s.Page = &v
	return s
}

func (s *ListPublicBroadcastSceneTemplatesRequest) SetSize(v int64) *ListPublicBroadcastSceneTemplatesRequest {
	s.Size = &v
	return s
}

func (s *ListPublicBroadcastSceneTemplatesRequest) SetTags(v string) *ListPublicBroadcastSceneTemplatesRequest {
	s.Tags = &v
	return s
}

func (s *ListPublicBroadcastSceneTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
