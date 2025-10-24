// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListTagGroupRequest
	GetAgentKey() *string
	SetClientToken(v string) *ListTagGroupRequest
	GetClientToken() *string
	SetGroupName(v string) *ListTagGroupRequest
	GetGroupName() *string
	SetPageNumber(v int32) *ListTagGroupRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTagGroupRequest
	GetPageSize() *int32
}

type ListTagGroupRequest struct {
	// example:
	//
	// xxxx
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// xxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// xx
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListTagGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagGroupRequest) GoString() string {
	return s.String()
}

func (s *ListTagGroupRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListTagGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListTagGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ListTagGroupRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTagGroupRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTagGroupRequest) SetAgentKey(v string) *ListTagGroupRequest {
	s.AgentKey = &v
	return s
}

func (s *ListTagGroupRequest) SetClientToken(v string) *ListTagGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *ListTagGroupRequest) SetGroupName(v string) *ListTagGroupRequest {
	s.GroupName = &v
	return s
}

func (s *ListTagGroupRequest) SetPageNumber(v int32) *ListTagGroupRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTagGroupRequest) SetPageSize(v int32) *ListTagGroupRequest {
	s.PageSize = &v
	return s
}

func (s *ListTagGroupRequest) Validate() error {
	return dara.Validate(s)
}
