// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKnowledgeBase interface {
	dara.Model
	String() string
	GoString() string
	SetCoverUri(v string) *KnowledgeBase
	GetCoverUri() *string
	SetCreatedAt(v int64) *KnowledgeBase
	GetCreatedAt() *int64
	SetDescription(v string) *KnowledgeBase
	GetDescription() *string
	SetFileFilter(v string) *KnowledgeBase
	GetFileFilter() *string
	SetKnowledgeBaseId(v string) *KnowledgeBase
	GetKnowledgeBaseId() *string
	SetLinkRuleList(v []*LinkRule) *KnowledgeBase
	GetLinkRuleList() []*LinkRule
	SetName(v string) *KnowledgeBase
	GetName() *string
	SetOwnerId(v string) *KnowledgeBase
	GetOwnerId() *string
	SetOwnerName(v string) *KnowledgeBase
	GetOwnerName() *string
	SetOwnerType(v string) *KnowledgeBase
	GetOwnerType() *string
	SetUpdatedAt(v int64) *KnowledgeBase
	GetUpdatedAt() *int64
}

type KnowledgeBase struct {
	CoverUri        *string     `json:"cover_uri,omitempty" xml:"cover_uri,omitempty"`
	CreatedAt       *int64      `json:"created_at,omitempty" xml:"created_at,omitempty"`
	Description     *string     `json:"description,omitempty" xml:"description,omitempty"`
	FileFilter      *string     `json:"file_filter,omitempty" xml:"file_filter,omitempty"`
	KnowledgeBaseId *string     `json:"knowledge_base_id,omitempty" xml:"knowledge_base_id,omitempty"`
	LinkRuleList    []*LinkRule `json:"link_rule_list,omitempty" xml:"link_rule_list,omitempty" type:"Repeated"`
	Name            *string     `json:"name,omitempty" xml:"name,omitempty"`
	OwnerId         *string     `json:"owner_id,omitempty" xml:"owner_id,omitempty"`
	OwnerName       *string     `json:"owner_name,omitempty" xml:"owner_name,omitempty"`
	OwnerType       *string     `json:"owner_type,omitempty" xml:"owner_type,omitempty"`
	UpdatedAt       *int64      `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

func (s KnowledgeBase) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeBase) GoString() string {
	return s.String()
}

func (s *KnowledgeBase) GetCoverUri() *string {
	return s.CoverUri
}

func (s *KnowledgeBase) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *KnowledgeBase) GetDescription() *string {
	return s.Description
}

func (s *KnowledgeBase) GetFileFilter() *string {
	return s.FileFilter
}

func (s *KnowledgeBase) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *KnowledgeBase) GetLinkRuleList() []*LinkRule {
	return s.LinkRuleList
}

func (s *KnowledgeBase) GetName() *string {
	return s.Name
}

func (s *KnowledgeBase) GetOwnerId() *string {
	return s.OwnerId
}

func (s *KnowledgeBase) GetOwnerName() *string {
	return s.OwnerName
}

func (s *KnowledgeBase) GetOwnerType() *string {
	return s.OwnerType
}

func (s *KnowledgeBase) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *KnowledgeBase) SetCoverUri(v string) *KnowledgeBase {
	s.CoverUri = &v
	return s
}

func (s *KnowledgeBase) SetCreatedAt(v int64) *KnowledgeBase {
	s.CreatedAt = &v
	return s
}

func (s *KnowledgeBase) SetDescription(v string) *KnowledgeBase {
	s.Description = &v
	return s
}

func (s *KnowledgeBase) SetFileFilter(v string) *KnowledgeBase {
	s.FileFilter = &v
	return s
}

func (s *KnowledgeBase) SetKnowledgeBaseId(v string) *KnowledgeBase {
	s.KnowledgeBaseId = &v
	return s
}

func (s *KnowledgeBase) SetLinkRuleList(v []*LinkRule) *KnowledgeBase {
	s.LinkRuleList = v
	return s
}

func (s *KnowledgeBase) SetName(v string) *KnowledgeBase {
	s.Name = &v
	return s
}

func (s *KnowledgeBase) SetOwnerId(v string) *KnowledgeBase {
	s.OwnerId = &v
	return s
}

func (s *KnowledgeBase) SetOwnerName(v string) *KnowledgeBase {
	s.OwnerName = &v
	return s
}

func (s *KnowledgeBase) SetOwnerType(v string) *KnowledgeBase {
	s.OwnerType = &v
	return s
}

func (s *KnowledgeBase) SetUpdatedAt(v int64) *KnowledgeBase {
	s.UpdatedAt = &v
	return s
}

func (s *KnowledgeBase) Validate() error {
	if s.LinkRuleList != nil {
		for _, item := range s.LinkRuleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
