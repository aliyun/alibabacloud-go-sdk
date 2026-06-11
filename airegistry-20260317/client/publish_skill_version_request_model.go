// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishSkillVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *PublishSkillVersionRequest
	GetNamespaceId() *string
	SetSkillName(v string) *PublishSkillVersionRequest
	GetSkillName() *string
	SetSkillVersion(v string) *PublishSkillVersionRequest
	GetSkillVersion() *string
	SetUpdateLatestLabel(v bool) *PublishSkillVersionRequest
	GetUpdateLatestLabel() *bool
}

type PublishSkillVersionRequest struct {
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
	// example:
	//
	// true
	UpdateLatestLabel *bool `json:"UpdateLatestLabel,omitempty" xml:"UpdateLatestLabel,omitempty"`
}

func (s PublishSkillVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishSkillVersionRequest) GoString() string {
	return s.String()
}

func (s *PublishSkillVersionRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *PublishSkillVersionRequest) GetSkillName() *string {
	return s.SkillName
}

func (s *PublishSkillVersionRequest) GetSkillVersion() *string {
	return s.SkillVersion
}

func (s *PublishSkillVersionRequest) GetUpdateLatestLabel() *bool {
	return s.UpdateLatestLabel
}

func (s *PublishSkillVersionRequest) SetNamespaceId(v string) *PublishSkillVersionRequest {
	s.NamespaceId = &v
	return s
}

func (s *PublishSkillVersionRequest) SetSkillName(v string) *PublishSkillVersionRequest {
	s.SkillName = &v
	return s
}

func (s *PublishSkillVersionRequest) SetSkillVersion(v string) *PublishSkillVersionRequest {
	s.SkillVersion = &v
	return s
}

func (s *PublishSkillVersionRequest) SetUpdateLatestLabel(v bool) *PublishSkillVersionRequest {
	s.UpdateLatestLabel = &v
	return s
}

func (s *PublishSkillVersionRequest) Validate() error {
	return dara.Validate(s)
}
