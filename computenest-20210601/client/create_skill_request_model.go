// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateSkillRequest
	GetClientToken() *string
	SetOssUrl(v string) *CreateSkillRequest
	GetOssUrl() *string
	SetSkillDescription(v string) *CreateSkillRequest
	GetSkillDescription() *string
	SetSkillDisplayName(v string) *CreateSkillRequest
	GetSkillDisplayName() *string
	SetSkillLabels(v []*string) *CreateSkillRequest
	GetSkillLabels() []*string
	SetSkillName(v string) *CreateSkillRequest
	GetSkillName() *string
	SetSkillSpaceId(v string) *CreateSkillRequest
	GetSkillSpaceId() *string
	SetSourceSkillId(v string) *CreateSkillRequest
	GetSourceSkillId() *string
	SetSourceType(v string) *CreateSkillRequest
	GetSourceType() *string
}

type CreateSkillRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The OSS URL of the Skill package to upload. This parameter is required when SourceType is set to UPLOAD.
	//
	// example:
	//
	// https://embedding-pic.oss-cn-beijing-internal.aliyuncs.com/30516570
	OssUrl *string `json:"OssUrl,omitempty" xml:"OssUrl,omitempty"`
	// The Skill description.
	//
	// example:
	//
	// 11111
	SkillDescription *string `json:"SkillDescription,omitempty" xml:"SkillDescription,omitempty"`
	SkillDisplayName *string `json:"SkillDisplayName,omitempty" xml:"SkillDisplayName,omitempty"`
	// The Skill labels.
	//
	// example:
	//
	// ["category:frontend-development"]
	SkillLabels []*string `json:"SkillLabels,omitempty" xml:"SkillLabels,omitempty" type:"Repeated"`
	// The Skill name.
	//
	// example:
	//
	// 11111
	SkillName *string `json:"SkillName,omitempty" xml:"SkillName,omitempty"`
	// The ID of the SkillSpace to which the Skill belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ss-111111
	SkillSpaceId *string `json:"SkillSpaceId,omitempty" xml:"SkillSpaceId,omitempty"`
	// The public Skill ID. This parameter is required when SourceType is set to COPY.
	//
	// example:
	//
	// s-11111
	SourceSkillId *string `json:"SourceSkillId,omitempty" xml:"SourceSkillId,omitempty"`
	// The source type used when creating the Skill.
	//
	// This parameter is required.
	//
	// example:
	//
	// COPY
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s CreateSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillRequest) GoString() string {
	return s.String()
}

func (s *CreateSkillRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateSkillRequest) GetOssUrl() *string {
	return s.OssUrl
}

func (s *CreateSkillRequest) GetSkillDescription() *string {
	return s.SkillDescription
}

func (s *CreateSkillRequest) GetSkillDisplayName() *string {
	return s.SkillDisplayName
}

func (s *CreateSkillRequest) GetSkillLabels() []*string {
	return s.SkillLabels
}

func (s *CreateSkillRequest) GetSkillName() *string {
	return s.SkillName
}

func (s *CreateSkillRequest) GetSkillSpaceId() *string {
	return s.SkillSpaceId
}

func (s *CreateSkillRequest) GetSourceSkillId() *string {
	return s.SourceSkillId
}

func (s *CreateSkillRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateSkillRequest) SetClientToken(v string) *CreateSkillRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSkillRequest) SetOssUrl(v string) *CreateSkillRequest {
	s.OssUrl = &v
	return s
}

func (s *CreateSkillRequest) SetSkillDescription(v string) *CreateSkillRequest {
	s.SkillDescription = &v
	return s
}

func (s *CreateSkillRequest) SetSkillDisplayName(v string) *CreateSkillRequest {
	s.SkillDisplayName = &v
	return s
}

func (s *CreateSkillRequest) SetSkillLabels(v []*string) *CreateSkillRequest {
	s.SkillLabels = v
	return s
}

func (s *CreateSkillRequest) SetSkillName(v string) *CreateSkillRequest {
	s.SkillName = &v
	return s
}

func (s *CreateSkillRequest) SetSkillSpaceId(v string) *CreateSkillRequest {
	s.SkillSpaceId = &v
	return s
}

func (s *CreateSkillRequest) SetSourceSkillId(v string) *CreateSkillRequest {
	s.SourceSkillId = &v
	return s
}

func (s *CreateSkillRequest) SetSourceType(v string) *CreateSkillRequest {
	s.SourceType = &v
	return s
}

func (s *CreateSkillRequest) Validate() error {
	return dara.Validate(s)
}
