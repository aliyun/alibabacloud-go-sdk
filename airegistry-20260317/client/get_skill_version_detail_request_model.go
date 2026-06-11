// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillVersionDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *GetSkillVersionDetailRequest
	GetNamespaceId() *string
	SetSkillName(v string) *GetSkillVersionDetailRequest
	GetSkillName() *string
	SetSkillVersion(v string) *GetSkillVersionDetailRequest
	GetSkillVersion() *string
}

type GetSkillVersionDetailRequest struct {
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
	// 0.0.1
	SkillVersion *string `json:"SkillVersion,omitempty" xml:"SkillVersion,omitempty"`
}

func (s GetSkillVersionDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSkillVersionDetailRequest) GoString() string {
	return s.String()
}

func (s *GetSkillVersionDetailRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *GetSkillVersionDetailRequest) GetSkillName() *string {
	return s.SkillName
}

func (s *GetSkillVersionDetailRequest) GetSkillVersion() *string {
	return s.SkillVersion
}

func (s *GetSkillVersionDetailRequest) SetNamespaceId(v string) *GetSkillVersionDetailRequest {
	s.NamespaceId = &v
	return s
}

func (s *GetSkillVersionDetailRequest) SetSkillName(v string) *GetSkillVersionDetailRequest {
	s.SkillName = &v
	return s
}

func (s *GetSkillVersionDetailRequest) SetSkillVersion(v string) *GetSkillVersionDetailRequest {
	s.SkillVersion = &v
	return s
}

func (s *GetSkillVersionDetailRequest) Validate() error {
	return dara.Validate(s)
}
