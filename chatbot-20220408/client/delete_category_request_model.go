// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *DeleteCategoryRequest
	GetAgentKey() *string
	SetCategoryId(v int64) *DeleteCategoryRequest
	GetCategoryId() *int64
}

type DeleteCategoryRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30000049006
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
}

func (s DeleteCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCategoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteCategoryRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *DeleteCategoryRequest) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *DeleteCategoryRequest) SetAgentKey(v string) *DeleteCategoryRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteCategoryRequest) SetCategoryId(v int64) *DeleteCategoryRequest {
	s.CategoryId = &v
	return s
}

func (s *DeleteCategoryRequest) Validate() error {
	return dara.Validate(s)
}
