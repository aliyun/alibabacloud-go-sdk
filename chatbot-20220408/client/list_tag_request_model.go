// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListTagRequest
	GetAgentKey() *string
	SetClientToken(v string) *ListTagRequest
	GetClientToken() *string
	SetGroupId(v int64) *ListTagRequest
	GetGroupId() *int64
	SetPageNumber(v int32) *ListTagRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTagRequest
	GetPageSize() *int32
	SetTagName(v string) *ListTagRequest
	GetTagName() *string
}

type ListTagRequest struct {
	// example:
	//
	// xxx
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// xxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// 43445343
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// xx
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s ListTagRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagRequest) GoString() string {
	return s.String()
}

func (s *ListTagRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListTagRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListTagRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *ListTagRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTagRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTagRequest) GetTagName() *string {
	return s.TagName
}

func (s *ListTagRequest) SetAgentKey(v string) *ListTagRequest {
	s.AgentKey = &v
	return s
}

func (s *ListTagRequest) SetClientToken(v string) *ListTagRequest {
	s.ClientToken = &v
	return s
}

func (s *ListTagRequest) SetGroupId(v int64) *ListTagRequest {
	s.GroupId = &v
	return s
}

func (s *ListTagRequest) SetPageNumber(v int32) *ListTagRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTagRequest) SetPageSize(v int32) *ListTagRequest {
	s.PageSize = &v
	return s
}

func (s *ListTagRequest) SetTagName(v string) *ListTagRequest {
	s.TagName = &v
	return s
}

func (s *ListTagRequest) Validate() error {
	return dara.Validate(s)
}
