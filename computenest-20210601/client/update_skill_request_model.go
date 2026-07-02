// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateSkillRequest
	GetClientToken() *string
	SetOssUrl(v string) *UpdateSkillRequest
	GetOssUrl() *string
	SetSkillDescription(v string) *UpdateSkillRequest
	GetSkillDescription() *string
	SetSkillDisplayName(v string) *UpdateSkillRequest
	GetSkillDisplayName() *string
	SetSkillId(v string) *UpdateSkillRequest
	GetSkillId() *string
	SetSkillLabels(v []*string) *UpdateSkillRequest
	GetSkillLabels() []*string
	SetSkillName(v string) *UpdateSkillRequest
	GetSkillName() *string
	SetSourceSkillId(v string) *UpdateSkillRequest
	GetSourceSkillId() *string
	SetSourceType(v string) *UpdateSkillRequest
	GetSourceType() *string
}

type UpdateSkillRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Required when SourceType is set to UPLOAD. The OSS URL of the skill package to upload.
	//
	// example:
	//
	// https://embedding-pic.oss-cn-beijing-internal.aliyuncs.com/30516570
	OssUrl *string `json:"OssUrl,omitempty" xml:"OssUrl,omitempty"`
	// The description of the skill.
	//
	// example:
	//
	// 11111
	SkillDescription *string `json:"SkillDescription,omitempty" xml:"SkillDescription,omitempty"`
	SkillDisplayName *string `json:"SkillDisplayName,omitempty" xml:"SkillDisplayName,omitempty"`
	// The ID of the skill to update.
	//
	// This parameter is required.
	//
	// example:
	//
	// 06e9dca2-0ac9-4d2e-a965-e9db9c057e00
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	// The labels of the skill.
	SkillLabels []*string `json:"SkillLabels,omitempty" xml:"SkillLabels,omitempty" type:"Repeated"`
	// The name of the skill.
	//
	// example:
	//
	// 111111
	SkillName *string `json:"SkillName,omitempty" xml:"SkillName,omitempty"`
	// Required when SourceType is set to COPY. The ID of the public skill.
	//
	// example:
	//
	// s-111
	SourceSkillId *string `json:"SourceSkillId,omitempty" xml:"SourceSkillId,omitempty"`
	// The source type for updating the skill.
	//
	// example:
	//
	// COPY
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s UpdateSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSkillRequest) GoString() string {
	return s.String()
}

func (s *UpdateSkillRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateSkillRequest) GetOssUrl() *string {
	return s.OssUrl
}

func (s *UpdateSkillRequest) GetSkillDescription() *string {
	return s.SkillDescription
}

func (s *UpdateSkillRequest) GetSkillDisplayName() *string {
	return s.SkillDisplayName
}

func (s *UpdateSkillRequest) GetSkillId() *string {
	return s.SkillId
}

func (s *UpdateSkillRequest) GetSkillLabels() []*string {
	return s.SkillLabels
}

func (s *UpdateSkillRequest) GetSkillName() *string {
	return s.SkillName
}

func (s *UpdateSkillRequest) GetSourceSkillId() *string {
	return s.SourceSkillId
}

func (s *UpdateSkillRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateSkillRequest) SetClientToken(v string) *UpdateSkillRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateSkillRequest) SetOssUrl(v string) *UpdateSkillRequest {
	s.OssUrl = &v
	return s
}

func (s *UpdateSkillRequest) SetSkillDescription(v string) *UpdateSkillRequest {
	s.SkillDescription = &v
	return s
}

func (s *UpdateSkillRequest) SetSkillDisplayName(v string) *UpdateSkillRequest {
	s.SkillDisplayName = &v
	return s
}

func (s *UpdateSkillRequest) SetSkillId(v string) *UpdateSkillRequest {
	s.SkillId = &v
	return s
}

func (s *UpdateSkillRequest) SetSkillLabels(v []*string) *UpdateSkillRequest {
	s.SkillLabels = v
	return s
}

func (s *UpdateSkillRequest) SetSkillName(v string) *UpdateSkillRequest {
	s.SkillName = &v
	return s
}

func (s *UpdateSkillRequest) SetSourceSkillId(v string) *UpdateSkillRequest {
	s.SourceSkillId = &v
	return s
}

func (s *UpdateSkillRequest) SetSourceType(v string) *UpdateSkillRequest {
	s.SourceType = &v
	return s
}

func (s *UpdateSkillRequest) Validate() error {
	return dara.Validate(s)
}
