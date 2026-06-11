// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iForcePublishSkillVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *ForcePublishSkillVersionRequest
	GetNamespaceId() *string
	SetSkillName(v string) *ForcePublishSkillVersionRequest
	GetSkillName() *string
	SetSkillVersion(v string) *ForcePublishSkillVersionRequest
	GetSkillVersion() *string
	SetUpdateLatestLabel(v bool) *ForcePublishSkillVersionRequest
	GetUpdateLatestLabel() *bool
}

type ForcePublishSkillVersionRequest struct {
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

func (s ForcePublishSkillVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s ForcePublishSkillVersionRequest) GoString() string {
	return s.String()
}

func (s *ForcePublishSkillVersionRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ForcePublishSkillVersionRequest) GetSkillName() *string {
	return s.SkillName
}

func (s *ForcePublishSkillVersionRequest) GetSkillVersion() *string {
	return s.SkillVersion
}

func (s *ForcePublishSkillVersionRequest) GetUpdateLatestLabel() *bool {
	return s.UpdateLatestLabel
}

func (s *ForcePublishSkillVersionRequest) SetNamespaceId(v string) *ForcePublishSkillVersionRequest {
	s.NamespaceId = &v
	return s
}

func (s *ForcePublishSkillVersionRequest) SetSkillName(v string) *ForcePublishSkillVersionRequest {
	s.SkillName = &v
	return s
}

func (s *ForcePublishSkillVersionRequest) SetSkillVersion(v string) *ForcePublishSkillVersionRequest {
	s.SkillVersion = &v
	return s
}

func (s *ForcePublishSkillVersionRequest) SetUpdateLatestLabel(v bool) *ForcePublishSkillVersionRequest {
	s.UpdateLatestLabel = &v
	return s
}

func (s *ForcePublishSkillVersionRequest) Validate() error {
	return dara.Validate(s)
}
