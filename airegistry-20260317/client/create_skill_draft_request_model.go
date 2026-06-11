// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSkillDraftRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBasedOnVersion(v string) *CreateSkillDraftRequest
	GetBasedOnVersion() *string
	SetCommitMsg(v string) *CreateSkillDraftRequest
	GetCommitMsg() *string
	SetNamespaceId(v string) *CreateSkillDraftRequest
	GetNamespaceId() *string
	SetSkillCard(v string) *CreateSkillDraftRequest
	GetSkillCard() *string
	SetSkillName(v string) *CreateSkillDraftRequest
	GetSkillName() *string
	SetTargetVersion(v string) *CreateSkillDraftRequest
	GetTargetVersion() *string
}

type CreateSkillDraftRequest struct {
	// example:
	//
	// 0.0.1
	BasedOnVersion *string `json:"BasedOnVersion,omitempty" xml:"BasedOnVersion,omitempty"`
	// example:
	//
	// 初始版本
	CommitMsg *string `json:"CommitMsg,omitempty" xml:"CommitMsg,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
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
	// example:
	//
	// 0.0.2
	TargetVersion *string `json:"TargetVersion,omitempty" xml:"TargetVersion,omitempty"`
}

func (s CreateSkillDraftRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillDraftRequest) GoString() string {
	return s.String()
}

func (s *CreateSkillDraftRequest) GetBasedOnVersion() *string {
	return s.BasedOnVersion
}

func (s *CreateSkillDraftRequest) GetCommitMsg() *string {
	return s.CommitMsg
}

func (s *CreateSkillDraftRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *CreateSkillDraftRequest) GetSkillCard() *string {
	return s.SkillCard
}

func (s *CreateSkillDraftRequest) GetSkillName() *string {
	return s.SkillName
}

func (s *CreateSkillDraftRequest) GetTargetVersion() *string {
	return s.TargetVersion
}

func (s *CreateSkillDraftRequest) SetBasedOnVersion(v string) *CreateSkillDraftRequest {
	s.BasedOnVersion = &v
	return s
}

func (s *CreateSkillDraftRequest) SetCommitMsg(v string) *CreateSkillDraftRequest {
	s.CommitMsg = &v
	return s
}

func (s *CreateSkillDraftRequest) SetNamespaceId(v string) *CreateSkillDraftRequest {
	s.NamespaceId = &v
	return s
}

func (s *CreateSkillDraftRequest) SetSkillCard(v string) *CreateSkillDraftRequest {
	s.SkillCard = &v
	return s
}

func (s *CreateSkillDraftRequest) SetSkillName(v string) *CreateSkillDraftRequest {
	s.SkillName = &v
	return s
}

func (s *CreateSkillDraftRequest) SetTargetVersion(v string) *CreateSkillDraftRequest {
	s.TargetVersion = &v
	return s
}

func (s *CreateSkillDraftRequest) Validate() error {
	return dara.Validate(s)
}
