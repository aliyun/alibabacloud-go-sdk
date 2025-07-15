// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotNewsWithTypeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListHotNewsWithTypeShrinkRequest
	GetAgentKey() *string
	SetCurrent(v int32) *ListHotNewsWithTypeShrinkRequest
	GetCurrent() *int32
	SetNewsType(v string) *ListHotNewsWithTypeShrinkRequest
	GetNewsType() *string
	SetNewsTypesShrink(v string) *ListHotNewsWithTypeShrinkRequest
	GetNewsTypesShrink() *string
	SetSize(v int32) *ListHotNewsWithTypeShrinkRequest
	GetSize() *int32
}

type ListHotNewsWithTypeShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c160c841c8e54295bf2f441432785944_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 1
	Current *int32 `json:"Current,omitempty" xml:"Current,omitempty"`
	// example:
	//
	// society
	NewsType        *string `json:"NewsType,omitempty" xml:"NewsType,omitempty"`
	NewsTypesShrink *string `json:"NewsTypes,omitempty" xml:"NewsTypes,omitempty"`
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListHotNewsWithTypeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHotNewsWithTypeShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListHotNewsWithTypeShrinkRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListHotNewsWithTypeShrinkRequest) GetCurrent() *int32 {
	return s.Current
}

func (s *ListHotNewsWithTypeShrinkRequest) GetNewsType() *string {
	return s.NewsType
}

func (s *ListHotNewsWithTypeShrinkRequest) GetNewsTypesShrink() *string {
	return s.NewsTypesShrink
}

func (s *ListHotNewsWithTypeShrinkRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListHotNewsWithTypeShrinkRequest) SetAgentKey(v string) *ListHotNewsWithTypeShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *ListHotNewsWithTypeShrinkRequest) SetCurrent(v int32) *ListHotNewsWithTypeShrinkRequest {
	s.Current = &v
	return s
}

func (s *ListHotNewsWithTypeShrinkRequest) SetNewsType(v string) *ListHotNewsWithTypeShrinkRequest {
	s.NewsType = &v
	return s
}

func (s *ListHotNewsWithTypeShrinkRequest) SetNewsTypesShrink(v string) *ListHotNewsWithTypeShrinkRequest {
	s.NewsTypesShrink = &v
	return s
}

func (s *ListHotNewsWithTypeShrinkRequest) SetSize(v int32) *ListHotNewsWithTypeShrinkRequest {
	s.Size = &v
	return s
}

func (s *ListHotNewsWithTypeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
