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
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// https://embedding-pic.oss-cn-beijing-internal.aliyuncs.com/30516570
	OssUrl *string `json:"OssUrl,omitempty" xml:"OssUrl,omitempty"`
	// example:
	//
	// 11111
	SkillDescription *string `json:"SkillDescription,omitempty" xml:"SkillDescription,omitempty"`
	// example:
	//
	// ["category:frontend-development"]
	SkillLabels []*string `json:"SkillLabels,omitempty" xml:"SkillLabels,omitempty" type:"Repeated"`
	// example:
	//
	// 11111
	SkillName *string `json:"SkillName,omitempty" xml:"SkillName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ss-111111
	SkillSpaceId *string `json:"SkillSpaceId,omitempty" xml:"SkillSpaceId,omitempty"`
	// example:
	//
	// s-11111
	SourceSkillId *string `json:"SourceSkillId,omitempty" xml:"SourceSkillId,omitempty"`
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
