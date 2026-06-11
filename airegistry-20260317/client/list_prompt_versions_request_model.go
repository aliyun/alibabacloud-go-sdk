// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPromptVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *ListPromptVersionsRequest
	GetNamespaceId() *string
	SetPageNo(v int32) *ListPromptVersionsRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListPromptVersionsRequest
	GetPageSize() *int32
	SetPromptKey(v string) *ListPromptVersionsRequest
	GetPromptKey() *string
}

type ListPromptVersionsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// customer-service-qa
	PromptKey *string `json:"PromptKey,omitempty" xml:"PromptKey,omitempty"`
}

func (s ListPromptVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPromptVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListPromptVersionsRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListPromptVersionsRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListPromptVersionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPromptVersionsRequest) GetPromptKey() *string {
	return s.PromptKey
}

func (s *ListPromptVersionsRequest) SetNamespaceId(v string) *ListPromptVersionsRequest {
	s.NamespaceId = &v
	return s
}

func (s *ListPromptVersionsRequest) SetPageNo(v int32) *ListPromptVersionsRequest {
	s.PageNo = &v
	return s
}

func (s *ListPromptVersionsRequest) SetPageSize(v int32) *ListPromptVersionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListPromptVersionsRequest) SetPromptKey(v string) *ListPromptVersionsRequest {
	s.PromptKey = &v
	return s
}

func (s *ListPromptVersionsRequest) Validate() error {
	return dara.Validate(s)
}
