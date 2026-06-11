// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *DeleteSkillRequest
	GetNamespaceId() *string
	SetSkillName(v string) *DeleteSkillRequest
	GetSkillName() *string
}

type DeleteSkillRequest struct {
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

func (s DeleteSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSkillRequest) GoString() string {
	return s.String()
}

func (s *DeleteSkillRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *DeleteSkillRequest) GetSkillName() *string {
	return s.SkillName
}

func (s *DeleteSkillRequest) SetNamespaceId(v string) *DeleteSkillRequest {
	s.NamespaceId = &v
	return s
}

func (s *DeleteSkillRequest) SetSkillName(v string) *DeleteSkillRequest {
	s.SkillName = &v
	return s
}

func (s *DeleteSkillRequest) Validate() error {
	return dara.Validate(s)
}
