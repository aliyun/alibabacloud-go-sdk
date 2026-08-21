// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSkillFileDetectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateSkillFileDetectRequest
	GetClientToken() *string
	SetOssUrl(v string) *CreateSkillFileDetectRequest
	GetOssUrl() *string
	SetRegionId(v string) *CreateSkillFileDetectRequest
	GetRegionId() *string
}

type CreateSkillFileDetectRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests. The value of **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The OSS URL of the Skill compressed file to be detected.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://embedding-pic.oss-cn-beijing-internal.aliyuncs.com/skill-creator.zip
	OssUrl *string `json:"OssUrl,omitempty" xml:"OssUrl,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateSkillFileDetectRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillFileDetectRequest) GoString() string {
	return s.String()
}

func (s *CreateSkillFileDetectRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateSkillFileDetectRequest) GetOssUrl() *string {
	return s.OssUrl
}

func (s *CreateSkillFileDetectRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSkillFileDetectRequest) SetClientToken(v string) *CreateSkillFileDetectRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSkillFileDetectRequest) SetOssUrl(v string) *CreateSkillFileDetectRequest {
	s.OssUrl = &v
	return s
}

func (s *CreateSkillFileDetectRequest) SetRegionId(v string) *CreateSkillFileDetectRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSkillFileDetectRequest) Validate() error {
	return dara.Validate(s)
}
