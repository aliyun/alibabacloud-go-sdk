// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRegistryNamespacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListRegistryNamespacesRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ListRegistryNamespacesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListRegistryNamespacesRequest
	GetNextToken() *string
	SetType(v string) *ListRegistryNamespacesRequest
	GetType() *string
}

type ListRegistryNamespacesRequest struct {
	// example:
	//
	// key
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// hg7nXVngyM6tQlfXYzM1uI/7dKNGp1JMgsKtvCagmtY=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// self
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListRegistryNamespacesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRegistryNamespacesRequest) GoString() string {
	return s.String()
}

func (s *ListRegistryNamespacesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListRegistryNamespacesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRegistryNamespacesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRegistryNamespacesRequest) GetType() *string {
	return s.Type
}

func (s *ListRegistryNamespacesRequest) SetKeyword(v string) *ListRegistryNamespacesRequest {
	s.Keyword = &v
	return s
}

func (s *ListRegistryNamespacesRequest) SetMaxResults(v int32) *ListRegistryNamespacesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRegistryNamespacesRequest) SetNextToken(v string) *ListRegistryNamespacesRequest {
	s.NextToken = &v
	return s
}

func (s *ListRegistryNamespacesRequest) SetType(v string) *ListRegistryNamespacesRequest {
	s.Type = &v
	return s
}

func (s *ListRegistryNamespacesRequest) Validate() error {
	return dara.Validate(s)
}
