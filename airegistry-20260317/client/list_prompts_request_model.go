// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPromptsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizTags(v string) *ListPromptsRequest
	GetBizTags() *string
	SetNamespaceId(v string) *ListPromptsRequest
	GetNamespaceId() *string
	SetPageNo(v int32) *ListPromptsRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListPromptsRequest
	GetPageSize() *int32
	SetPromptKey(v string) *ListPromptsRequest
	GetPromptKey() *string
	SetSearch(v string) *ListPromptsRequest
	GetSearch() *string
}

type ListPromptsRequest struct {
	// The business tags used to filter prompts. Separate multiple tags with commas (,).
	//
	// example:
	//
	// cs,qa
	BizTags *string `json:"BizTags,omitempty" xml:"BizTags,omitempty"`
	// The namespace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 50.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The identifier of the prompt to search for. This parameter is optional.
	//
	// example:
	//
	// customer
	PromptKey *string `json:"PromptKey,omitempty" xml:"PromptKey,omitempty"`
	// The search mode. Valid values:
	//
	// - accurate: exact match.
	//
	// - blur: fuzzy match.
	//
	// example:
	//
	// blur
	Search *string `json:"Search,omitempty" xml:"Search,omitempty"`
}

func (s ListPromptsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPromptsRequest) GoString() string {
	return s.String()
}

func (s *ListPromptsRequest) GetBizTags() *string {
	return s.BizTags
}

func (s *ListPromptsRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListPromptsRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListPromptsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPromptsRequest) GetPromptKey() *string {
	return s.PromptKey
}

func (s *ListPromptsRequest) GetSearch() *string {
	return s.Search
}

func (s *ListPromptsRequest) SetBizTags(v string) *ListPromptsRequest {
	s.BizTags = &v
	return s
}

func (s *ListPromptsRequest) SetNamespaceId(v string) *ListPromptsRequest {
	s.NamespaceId = &v
	return s
}

func (s *ListPromptsRequest) SetPageNo(v int32) *ListPromptsRequest {
	s.PageNo = &v
	return s
}

func (s *ListPromptsRequest) SetPageSize(v int32) *ListPromptsRequest {
	s.PageSize = &v
	return s
}

func (s *ListPromptsRequest) SetPromptKey(v string) *ListPromptsRequest {
	s.PromptKey = &v
	return s
}

func (s *ListPromptsRequest) SetSearch(v string) *ListPromptsRequest {
	s.Search = &v
	return s
}

func (s *ListPromptsRequest) Validate() error {
	return dara.Validate(s)
}
