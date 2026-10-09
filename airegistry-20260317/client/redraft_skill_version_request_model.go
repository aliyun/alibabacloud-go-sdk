// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRedraftSkillVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *RedraftSkillVersionRequest
	GetNamespaceId() *string
	SetSkillName(v string) *RedraftSkillVersionRequest
	GetSkillName() *string
	SetSkillVersion(v string) *RedraftSkillVersionRequest
	GetSkillVersion() *string
}

type RedraftSkillVersionRequest struct {
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
	// customer-service-skill
	SkillName *string `json:"SkillName,omitempty" xml:"SkillName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.0.2
	SkillVersion *string `json:"SkillVersion,omitempty" xml:"SkillVersion,omitempty"`
}

func (s RedraftSkillVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s RedraftSkillVersionRequest) GoString() string {
	return s.String()
}

func (s *RedraftSkillVersionRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *RedraftSkillVersionRequest) GetSkillName() *string {
	return s.SkillName
}

func (s *RedraftSkillVersionRequest) GetSkillVersion() *string {
	return s.SkillVersion
}

func (s *RedraftSkillVersionRequest) SetNamespaceId(v string) *RedraftSkillVersionRequest {
	s.NamespaceId = &v
	return s
}

func (s *RedraftSkillVersionRequest) SetSkillName(v string) *RedraftSkillVersionRequest {
	s.SkillName = &v
	return s
}

func (s *RedraftSkillVersionRequest) SetSkillVersion(v string) *RedraftSkillVersionRequest {
	s.SkillVersion = &v
	return s
}

func (s *RedraftSkillVersionRequest) Validate() error {
	return dara.Validate(s)
}
