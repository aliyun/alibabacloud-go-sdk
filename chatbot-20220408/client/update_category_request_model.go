// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *UpdateCategoryRequest
	GetAgentKey() *string
	SetBizCode(v string) *UpdateCategoryRequest
	GetBizCode() *string
	SetCategoryId(v int64) *UpdateCategoryRequest
	GetCategoryId() *int64
	SetName(v string) *UpdateCategoryRequest
	GetName() *string
}

type UpdateCategoryRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	BizCode  *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 231001028593
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCategoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateCategoryRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *UpdateCategoryRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *UpdateCategoryRequest) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *UpdateCategoryRequest) GetName() *string {
	return s.Name
}

func (s *UpdateCategoryRequest) SetAgentKey(v string) *UpdateCategoryRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateCategoryRequest) SetBizCode(v string) *UpdateCategoryRequest {
	s.BizCode = &v
	return s
}

func (s *UpdateCategoryRequest) SetCategoryId(v int64) *UpdateCategoryRequest {
	s.CategoryId = &v
	return s
}

func (s *UpdateCategoryRequest) SetName(v string) *UpdateCategoryRequest {
	s.Name = &v
	return s
}

func (s *UpdateCategoryRequest) Validate() error {
	return dara.Validate(s)
}
