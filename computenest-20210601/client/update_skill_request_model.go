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
	// A unique, client-generated token to ensure request idempotence. **ClientToken*	- can contain only ASCII characters and must not exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required if `SourceType` is set to `UPLOAD`. It specifies the Object Storage Service (OSS) URL of the compressed skill package to upload.
	//
	// example:
	//
	// https://embedding-pic.oss-cn-beijing-internal.aliyuncs.com/30516570
	OssUrl *string `json:"OssUrl,omitempty" xml:"OssUrl,omitempty"`
	// The skill description.
	//
	// example:
	//
	// 11111
	SkillDescription *string `json:"SkillDescription,omitempty" xml:"SkillDescription,omitempty"`
	// The ID of the skill to update.
	//
	// This parameter is required.
	//
	// example:
	//
	// 06e9dca2-0ac9-4d2e-a965-e9db9c057e00
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	// An array of skill labels.
	SkillLabels []*string `json:"SkillLabels,omitempty" xml:"SkillLabels,omitempty" type:"Repeated"`
	// The skill name.
	//
	// example:
	//
	// 111111
	SkillName *string `json:"SkillName,omitempty" xml:"SkillName,omitempty"`
	// This parameter is required if `SourceType` is set to `COPY`. It specifies the ID of the public skill.
	//
	// example:
	//
	// s-111
	SourceSkillId *string `json:"SourceSkillId,omitempty" xml:"SourceSkillId,omitempty"`
	// The source type for the skill update.
	//
	// This parameter is required.
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
