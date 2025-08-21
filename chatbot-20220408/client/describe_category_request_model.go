// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *DescribeCategoryRequest
	GetAgentKey() *string
	SetCategoryId(v int64) *DescribeCategoryRequest
	GetCategoryId() *int64
}

type DescribeCategoryRequest struct {
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

func (s DescribeCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCategoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeCategoryRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *DescribeCategoryRequest) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *DescribeCategoryRequest) SetAgentKey(v string) *DescribeCategoryRequest {
	s.AgentKey = &v
	return s
}

func (s *DescribeCategoryRequest) SetCategoryId(v int64) *DescribeCategoryRequest {
	s.CategoryId = &v
	return s
}

func (s *DescribeCategoryRequest) Validate() error {
	return dara.Validate(s)
}
