// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSkillDraftRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommitMsg(v string) *UpdateSkillDraftRequest
	GetCommitMsg() *string
	SetNamespaceId(v string) *UpdateSkillDraftRequest
	GetNamespaceId() *string
	SetSkillCard(v string) *UpdateSkillDraftRequest
	GetSkillCard() *string
	SetSkillName(v string) *UpdateSkillDraftRequest
	GetSkillName() *string
}

type UpdateSkillDraftRequest struct {
	// example:
	//
	// 更新说明
	CommitMsg *string `json:"CommitMsg,omitempty" xml:"CommitMsg,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"name":"customer-service-skill","description":"..."}
	SkillCard *string `json:"SkillCard,omitempty" xml:"SkillCard,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// customer-service-skill
	SkillName *string `json:"SkillName,omitempty" xml:"SkillName,omitempty"`
}

func (s UpdateSkillDraftRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSkillDraftRequest) GoString() string {
	return s.String()
}

func (s *UpdateSkillDraftRequest) GetCommitMsg() *string {
	return s.CommitMsg
}

func (s *UpdateSkillDraftRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *UpdateSkillDraftRequest) GetSkillCard() *string {
	return s.SkillCard
}

func (s *UpdateSkillDraftRequest) GetSkillName() *string {
	return s.SkillName
}

func (s *UpdateSkillDraftRequest) SetCommitMsg(v string) *UpdateSkillDraftRequest {
	s.CommitMsg = &v
	return s
}

func (s *UpdateSkillDraftRequest) SetNamespaceId(v string) *UpdateSkillDraftRequest {
	s.NamespaceId = &v
	return s
}

func (s *UpdateSkillDraftRequest) SetSkillCard(v string) *UpdateSkillDraftRequest {
	s.SkillCard = &v
	return s
}

func (s *UpdateSkillDraftRequest) SetSkillName(v string) *UpdateSkillDraftRequest {
	s.SkillName = &v
	return s
}

func (s *UpdateSkillDraftRequest) Validate() error {
	return dara.Validate(s)
}
