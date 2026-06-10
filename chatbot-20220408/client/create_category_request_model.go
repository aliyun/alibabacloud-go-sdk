// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *CreateCategoryRequest
	GetAgentKey() *string
	SetBizCode(v string) *CreateCategoryRequest
	GetBizCode() *string
	SetKnowledgeType(v int32) *CreateCategoryRequest
	GetKnowledgeType() *int32
	SetName(v string) *CreateCategoryRequest
	GetName() *string
	SetParentCategoryId(v int64) *CreateCategoryRequest
	GetParentCategoryId() *int64
}

type CreateCategoryRequest struct {
	// The key of the business space. If this parameter is not specified, the default business space is used. You can obtain the key from the Business Management page of your main account.
	//
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The business code.
	//
	// example:
	//
	// bizcode123
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// The knowledge type. Valid values are 1 for an FAQ category and 3 for a document category. The default value is 1.
	//
	// example:
	//
	// 1
	KnowledgeType *int32 `json:"KnowledgeType,omitempty" xml:"KnowledgeType,omitempty"`
	// The category name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 测试1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the parent category. Defaults to -1, which indicates the root category.
	//
	// example:
	//
	// -1
	ParentCategoryId *int64 `json:"ParentCategoryId,omitempty" xml:"ParentCategoryId,omitempty"`
}

func (s CreateCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCategoryRequest) GoString() string {
	return s.String()
}

func (s *CreateCategoryRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *CreateCategoryRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *CreateCategoryRequest) GetKnowledgeType() *int32 {
	return s.KnowledgeType
}

func (s *CreateCategoryRequest) GetName() *string {
	return s.Name
}

func (s *CreateCategoryRequest) GetParentCategoryId() *int64 {
	return s.ParentCategoryId
}

func (s *CreateCategoryRequest) SetAgentKey(v string) *CreateCategoryRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateCategoryRequest) SetBizCode(v string) *CreateCategoryRequest {
	s.BizCode = &v
	return s
}

func (s *CreateCategoryRequest) SetKnowledgeType(v int32) *CreateCategoryRequest {
	s.KnowledgeType = &v
	return s
}

func (s *CreateCategoryRequest) SetName(v string) *CreateCategoryRequest {
	s.Name = &v
	return s
}

func (s *CreateCategoryRequest) SetParentCategoryId(v int64) *CreateCategoryRequest {
	s.ParentCategoryId = &v
	return s
}

func (s *CreateCategoryRequest) Validate() error {
	return dara.Validate(s)
}
