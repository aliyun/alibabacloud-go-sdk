// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSkillBizTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizTags(v string) *UpdateSkillBizTagsRequest
	GetBizTags() *string
	SetNamespaceId(v string) *UpdateSkillBizTagsRequest
	GetNamespaceId() *string
	SetSkillName(v string) *UpdateSkillBizTagsRequest
	GetSkillName() *string
}

type UpdateSkillBizTagsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ["cs","qa","support"]
	BizTags *string `json:"BizTags,omitempty" xml:"BizTags,omitempty"`
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

func (s UpdateSkillBizTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSkillBizTagsRequest) GoString() string {
	return s.String()
}

func (s *UpdateSkillBizTagsRequest) GetBizTags() *string {
	return s.BizTags
}

func (s *UpdateSkillBizTagsRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *UpdateSkillBizTagsRequest) GetSkillName() *string {
	return s.SkillName
}

func (s *UpdateSkillBizTagsRequest) SetBizTags(v string) *UpdateSkillBizTagsRequest {
	s.BizTags = &v
	return s
}

func (s *UpdateSkillBizTagsRequest) SetNamespaceId(v string) *UpdateSkillBizTagsRequest {
	s.NamespaceId = &v
	return s
}

func (s *UpdateSkillBizTagsRequest) SetSkillName(v string) *UpdateSkillBizTagsRequest {
	s.SkillName = &v
	return s
}

func (s *UpdateSkillBizTagsRequest) Validate() error {
	return dara.Validate(s)
}
