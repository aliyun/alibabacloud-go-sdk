// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSkillVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *SubmitSkillVersionRequest
	GetNamespaceId() *string
	SetSkillName(v string) *SubmitSkillVersionRequest
	GetSkillName() *string
	SetSkillVersion(v string) *SubmitSkillVersionRequest
	GetSkillVersion() *string
}

type SubmitSkillVersionRequest struct {
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

func (s SubmitSkillVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitSkillVersionRequest) GoString() string {
	return s.String()
}

func (s *SubmitSkillVersionRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *SubmitSkillVersionRequest) GetSkillName() *string {
	return s.SkillName
}

func (s *SubmitSkillVersionRequest) GetSkillVersion() *string {
	return s.SkillVersion
}

func (s *SubmitSkillVersionRequest) SetNamespaceId(v string) *SubmitSkillVersionRequest {
	s.NamespaceId = &v
	return s
}

func (s *SubmitSkillVersionRequest) SetSkillName(v string) *SubmitSkillVersionRequest {
	s.SkillName = &v
	return s
}

func (s *SubmitSkillVersionRequest) SetSkillVersion(v string) *SubmitSkillVersionRequest {
	s.SkillVersion = &v
	return s
}

func (s *SubmitSkillVersionRequest) Validate() error {
	return dara.Validate(s)
}
