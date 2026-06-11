// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *GetSkillDetailRequest
	GetNamespaceId() *string
	SetSkillName(v string) *GetSkillDetailRequest
	GetSkillName() *string
}

type GetSkillDetailRequest struct {
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
}

func (s GetSkillDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSkillDetailRequest) GoString() string {
	return s.String()
}

func (s *GetSkillDetailRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *GetSkillDetailRequest) GetSkillName() *string {
	return s.SkillName
}

func (s *GetSkillDetailRequest) SetNamespaceId(v string) *GetSkillDetailRequest {
	s.NamespaceId = &v
	return s
}

func (s *GetSkillDetailRequest) SetSkillName(v string) *GetSkillDetailRequest {
	s.SkillName = &v
	return s
}

func (s *GetSkillDetailRequest) Validate() error {
	return dara.Validate(s)
}
