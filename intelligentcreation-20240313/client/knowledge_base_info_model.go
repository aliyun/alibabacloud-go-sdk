// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKnowledgeBaseInfo interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *KnowledgeBaseInfo
	GetDescription() *string
	SetId(v string) *KnowledgeBaseInfo
	GetId() *string
	SetIndustry(v string) *KnowledgeBaseInfo
	GetIndustry() *string
	SetKnowledgeBaseType(v string) *KnowledgeBaseInfo
	GetKnowledgeBaseType() *string
	SetName(v string) *KnowledgeBaseInfo
	GetName() *string
}

type KnowledgeBaseInfo struct {
	Description       *string `json:"description,omitempty" xml:"description,omitempty"`
	Id                *string `json:"id,omitempty" xml:"id,omitempty"`
	Industry          *string `json:"industry,omitempty" xml:"industry,omitempty"`
	KnowledgeBaseType *string `json:"knowledgeBaseType,omitempty" xml:"knowledgeBaseType,omitempty"`
	Name              *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s KnowledgeBaseInfo) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeBaseInfo) GoString() string {
	return s.String()
}

func (s *KnowledgeBaseInfo) GetDescription() *string {
	return s.Description
}

func (s *KnowledgeBaseInfo) GetId() *string {
	return s.Id
}

func (s *KnowledgeBaseInfo) GetIndustry() *string {
	return s.Industry
}

func (s *KnowledgeBaseInfo) GetKnowledgeBaseType() *string {
	return s.KnowledgeBaseType
}

func (s *KnowledgeBaseInfo) GetName() *string {
	return s.Name
}

func (s *KnowledgeBaseInfo) SetDescription(v string) *KnowledgeBaseInfo {
	s.Description = &v
	return s
}

func (s *KnowledgeBaseInfo) SetId(v string) *KnowledgeBaseInfo {
	s.Id = &v
	return s
}

func (s *KnowledgeBaseInfo) SetIndustry(v string) *KnowledgeBaseInfo {
	s.Industry = &v
	return s
}

func (s *KnowledgeBaseInfo) SetKnowledgeBaseType(v string) *KnowledgeBaseInfo {
	s.KnowledgeBaseType = &v
	return s
}

func (s *KnowledgeBaseInfo) SetName(v string) *KnowledgeBaseInfo {
	s.Name = &v
	return s
}

func (s *KnowledgeBaseInfo) Validate() error {
	return dara.Validate(s)
}
