// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *OfflineSkillRequest
	GetNamespaceId() *string
	SetScope(v string) *OfflineSkillRequest
	GetScope() *string
	SetSkillName(v string) *OfflineSkillRequest
	GetSkillName() *string
	SetSkillVersion(v string) *OfflineSkillRequest
	GetSkillVersion() *string
}

type OfflineSkillRequest struct {
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

func (s OfflineSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s OfflineSkillRequest) GoString() string {
	return s.String()
}

func (s *OfflineSkillRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *OfflineSkillRequest) GetScope() *string {
	return s.Scope
}

func (s *OfflineSkillRequest) GetSkillName() *string {
	return s.SkillName
}

func (s *OfflineSkillRequest) GetSkillVersion() *string {
	return s.SkillVersion
}

func (s *OfflineSkillRequest) SetNamespaceId(v string) *OfflineSkillRequest {
	s.NamespaceId = &v
	return s
}

func (s *OfflineSkillRequest) SetScope(v string) *OfflineSkillRequest {
	s.Scope = &v
	return s
}

func (s *OfflineSkillRequest) SetSkillName(v string) *OfflineSkillRequest {
	s.SkillName = &v
	return s
}

func (s *OfflineSkillRequest) SetSkillVersion(v string) *OfflineSkillRequest {
	s.SkillVersion = &v
	return s
}

func (s *OfflineSkillRequest) Validate() error {
	return dara.Validate(s)
}
