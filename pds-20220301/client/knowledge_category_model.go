// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKnowledgeCategory interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v int64) *KnowledgeCategory
	GetCreatedAt() *int64
	SetDescription(v string) *KnowledgeCategory
	GetDescription() *string
	SetKeywords(v []*string) *KnowledgeCategory
	GetKeywords() []*string
	SetKnowledgeBaseId(v string) *KnowledgeCategory
	GetKnowledgeBaseId() *string
	SetKnowledgeBaseName(v string) *KnowledgeCategory
	GetKnowledgeBaseName() *string
	SetKnowledgeCategoryId(v string) *KnowledgeCategory
	GetKnowledgeCategoryId() *string
	SetName(v string) *KnowledgeCategory
	GetName() *string
	SetOwner(v string) *KnowledgeCategory
	GetOwner() *string
	SetOwnerType(v string) *KnowledgeCategory
	GetOwnerType() *string
	SetParentKnowledgeCategoryId(v string) *KnowledgeCategory
	GetParentKnowledgeCategoryId() *string
	SetStatus(v string) *KnowledgeCategory
	GetStatus() *string
	SetUpdatedAt(v int64) *KnowledgeCategory
	GetUpdatedAt() *int64
}

type KnowledgeCategory struct {
	CreatedAt                 *int64    `json:"created_at,omitempty" xml:"created_at,omitempty"`
	Description               *string   `json:"description,omitempty" xml:"description,omitempty"`
	Keywords                  []*string `json:"keywords,omitempty" xml:"keywords,omitempty" type:"Repeated"`
	KnowledgeBaseId           *string   `json:"knowledge_base_id,omitempty" xml:"knowledge_base_id,omitempty"`
	KnowledgeBaseName         *string   `json:"knowledge_base_name,omitempty" xml:"knowledge_base_name,omitempty"`
	KnowledgeCategoryId       *string   `json:"knowledge_category_id,omitempty" xml:"knowledge_category_id,omitempty"`
	Name                      *string   `json:"name,omitempty" xml:"name,omitempty"`
	Owner                     *string   `json:"owner,omitempty" xml:"owner,omitempty"`
	OwnerType                 *string   `json:"owner_type,omitempty" xml:"owner_type,omitempty"`
	ParentKnowledgeCategoryId *string   `json:"parent_knowledge_category_id,omitempty" xml:"parent_knowledge_category_id,omitempty"`
	Status                    *string   `json:"status,omitempty" xml:"status,omitempty"`
	UpdatedAt                 *int64    `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

func (s KnowledgeCategory) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeCategory) GoString() string {
	return s.String()
}

func (s *KnowledgeCategory) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *KnowledgeCategory) GetDescription() *string {
	return s.Description
}

func (s *KnowledgeCategory) GetKeywords() []*string {
	return s.Keywords
}

func (s *KnowledgeCategory) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *KnowledgeCategory) GetKnowledgeBaseName() *string {
	return s.KnowledgeBaseName
}

func (s *KnowledgeCategory) GetKnowledgeCategoryId() *string {
	return s.KnowledgeCategoryId
}

func (s *KnowledgeCategory) GetName() *string {
	return s.Name
}

func (s *KnowledgeCategory) GetOwner() *string {
	return s.Owner
}

func (s *KnowledgeCategory) GetOwnerType() *string {
	return s.OwnerType
}

func (s *KnowledgeCategory) GetParentKnowledgeCategoryId() *string {
	return s.ParentKnowledgeCategoryId
}

func (s *KnowledgeCategory) GetStatus() *string {
	return s.Status
}

func (s *KnowledgeCategory) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *KnowledgeCategory) SetCreatedAt(v int64) *KnowledgeCategory {
	s.CreatedAt = &v
	return s
}

func (s *KnowledgeCategory) SetDescription(v string) *KnowledgeCategory {
	s.Description = &v
	return s
}

func (s *KnowledgeCategory) SetKeywords(v []*string) *KnowledgeCategory {
	s.Keywords = v
	return s
}

func (s *KnowledgeCategory) SetKnowledgeBaseId(v string) *KnowledgeCategory {
	s.KnowledgeBaseId = &v
	return s
}

func (s *KnowledgeCategory) SetKnowledgeBaseName(v string) *KnowledgeCategory {
	s.KnowledgeBaseName = &v
	return s
}

func (s *KnowledgeCategory) SetKnowledgeCategoryId(v string) *KnowledgeCategory {
	s.KnowledgeCategoryId = &v
	return s
}

func (s *KnowledgeCategory) SetName(v string) *KnowledgeCategory {
	s.Name = &v
	return s
}

func (s *KnowledgeCategory) SetOwner(v string) *KnowledgeCategory {
	s.Owner = &v
	return s
}

func (s *KnowledgeCategory) SetOwnerType(v string) *KnowledgeCategory {
	s.OwnerType = &v
	return s
}

func (s *KnowledgeCategory) SetParentKnowledgeCategoryId(v string) *KnowledgeCategory {
	s.ParentKnowledgeCategoryId = &v
	return s
}

func (s *KnowledgeCategory) SetStatus(v string) *KnowledgeCategory {
	s.Status = &v
	return s
}

func (s *KnowledgeCategory) SetUpdatedAt(v int64) *KnowledgeCategory {
	s.UpdatedAt = &v
	return s
}

func (s *KnowledgeCategory) Validate() error {
	return dara.Validate(s)
}
