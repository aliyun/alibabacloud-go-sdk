// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOneMetaKnowledgeBase interface {
	dara.Model
	String() string
	GoString() string
	SetCreator(v string) *OneMetaKnowledgeBase
	GetCreator() *string
	SetDescription(v string) *OneMetaKnowledgeBase
	GetDescription() *string
	SetGmtCreate(v string) *OneMetaKnowledgeBase
	GetGmtCreate() *string
	SetGmtModified(v string) *OneMetaKnowledgeBase
	GetGmtModified() *string
	SetKbUuid(v string) *OneMetaKnowledgeBase
	GetKbUuid() *string
	SetName(v string) *OneMetaKnowledgeBase
	GetName() *string
	SetState(v int32) *OneMetaKnowledgeBase
	GetState() *int32
	SetTag(v string) *OneMetaKnowledgeBase
	GetTag() *string
}

type OneMetaKnowledgeBase struct {
	Creator     *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate   *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	KbUuid      *string `json:"KbUuid,omitempty" xml:"KbUuid,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	State       *int32  `json:"State,omitempty" xml:"State,omitempty"`
	Tag         *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s OneMetaKnowledgeBase) String() string {
	return dara.Prettify(s)
}

func (s OneMetaKnowledgeBase) GoString() string {
	return s.String()
}

func (s *OneMetaKnowledgeBase) GetCreator() *string {
	return s.Creator
}

func (s *OneMetaKnowledgeBase) GetDescription() *string {
	return s.Description
}

func (s *OneMetaKnowledgeBase) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *OneMetaKnowledgeBase) GetGmtModified() *string {
	return s.GmtModified
}

func (s *OneMetaKnowledgeBase) GetKbUuid() *string {
	return s.KbUuid
}

func (s *OneMetaKnowledgeBase) GetName() *string {
	return s.Name
}

func (s *OneMetaKnowledgeBase) GetState() *int32 {
	return s.State
}

func (s *OneMetaKnowledgeBase) GetTag() *string {
	return s.Tag
}

func (s *OneMetaKnowledgeBase) SetCreator(v string) *OneMetaKnowledgeBase {
	s.Creator = &v
	return s
}

func (s *OneMetaKnowledgeBase) SetDescription(v string) *OneMetaKnowledgeBase {
	s.Description = &v
	return s
}

func (s *OneMetaKnowledgeBase) SetGmtCreate(v string) *OneMetaKnowledgeBase {
	s.GmtCreate = &v
	return s
}

func (s *OneMetaKnowledgeBase) SetGmtModified(v string) *OneMetaKnowledgeBase {
	s.GmtModified = &v
	return s
}

func (s *OneMetaKnowledgeBase) SetKbUuid(v string) *OneMetaKnowledgeBase {
	s.KbUuid = &v
	return s
}

func (s *OneMetaKnowledgeBase) SetName(v string) *OneMetaKnowledgeBase {
	s.Name = &v
	return s
}

func (s *OneMetaKnowledgeBase) SetState(v int32) *OneMetaKnowledgeBase {
	s.State = &v
	return s
}

func (s *OneMetaKnowledgeBase) SetTag(v string) *OneMetaKnowledgeBase {
	s.Tag = &v
	return s
}

func (s *OneMetaKnowledgeBase) Validate() error {
	return dara.Validate(s)
}
