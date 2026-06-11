// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnlineSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *OnlineSkillRequest
	GetNamespaceId() *string
	SetScope(v string) *OnlineSkillRequest
	GetScope() *string
	SetSkillName(v string) *OnlineSkillRequest
	GetSkillName() *string
	SetSkillVersion(v string) *OnlineSkillRequest
	GetSkillVersion() *string
}

type OnlineSkillRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// example:
	//
	// version
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// customer-service-skill
	SkillName *string `json:"SkillName,omitempty" xml:"SkillName,omitempty"`
	// example:
	//
	// 0.0.2
	SkillVersion *string `json:"SkillVersion,omitempty" xml:"SkillVersion,omitempty"`
}

func (s OnlineSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s OnlineSkillRequest) GoString() string {
	return s.String()
}

func (s *OnlineSkillRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *OnlineSkillRequest) GetScope() *string {
	return s.Scope
}

func (s *OnlineSkillRequest) GetSkillName() *string {
	return s.SkillName
}

func (s *OnlineSkillRequest) GetSkillVersion() *string {
	return s.SkillVersion
}

func (s *OnlineSkillRequest) SetNamespaceId(v string) *OnlineSkillRequest {
	s.NamespaceId = &v
	return s
}

func (s *OnlineSkillRequest) SetScope(v string) *OnlineSkillRequest {
	s.Scope = &v
	return s
}

func (s *OnlineSkillRequest) SetSkillName(v string) *OnlineSkillRequest {
	s.SkillName = &v
	return s
}

func (s *OnlineSkillRequest) SetSkillVersion(v string) *OnlineSkillRequest {
	s.SkillVersion = &v
	return s
}

func (s *OnlineSkillRequest) Validate() error {
	return dara.Validate(s)
}
