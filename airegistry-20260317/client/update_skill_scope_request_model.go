// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSkillScopeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *UpdateSkillScopeRequest
	GetNamespaceId() *string
	SetScope(v string) *UpdateSkillScopeRequest
	GetScope() *string
	SetSkillName(v string) *UpdateSkillScopeRequest
	GetSkillName() *string
}

type UpdateSkillScopeRequest struct {
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
	// PUBLIC
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// customer-service-skill
	SkillName *string `json:"SkillName,omitempty" xml:"SkillName,omitempty"`
}

func (s UpdateSkillScopeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSkillScopeRequest) GoString() string {
	return s.String()
}

func (s *UpdateSkillScopeRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *UpdateSkillScopeRequest) GetScope() *string {
	return s.Scope
}

func (s *UpdateSkillScopeRequest) GetSkillName() *string {
	return s.SkillName
}

func (s *UpdateSkillScopeRequest) SetNamespaceId(v string) *UpdateSkillScopeRequest {
	s.NamespaceId = &v
	return s
}

func (s *UpdateSkillScopeRequest) SetScope(v string) *UpdateSkillScopeRequest {
	s.Scope = &v
	return s
}

func (s *UpdateSkillScopeRequest) SetSkillName(v string) *UpdateSkillScopeRequest {
	s.SkillName = &v
	return s
}

func (s *UpdateSkillScopeRequest) Validate() error {
	return dara.Validate(s)
}
