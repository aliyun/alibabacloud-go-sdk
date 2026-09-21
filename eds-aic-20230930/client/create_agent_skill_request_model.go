// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileList(v []*string) *CreateAgentSkillRequest
	GetFileList() []*string
	SetIconKey(v string) *CreateAgentSkillRequest
	GetIconKey() *string
	SetPackageOssKey(v string) *CreateAgentSkillRequest
	GetPackageOssKey() *string
	SetSkillDescription(v string) *CreateAgentSkillRequest
	GetSkillDescription() *string
	SetSkillName(v string) *CreateAgentSkillRequest
	GetSkillName() *string
	SetSkillPackageUrl(v string) *CreateAgentSkillRequest
	GetSkillPackageUrl() *string
}

type CreateAgentSkillRequest struct {
	// The list of files in the skill package.
	FileList []*string `json:"FileList,omitempty" xml:"FileList,omitempty" type:"Repeated"`
	// The icon of the custom skill.
	//
	// example:
	//
	// icon01
	IconKey *string `json:"IconKey,omitempty" xml:"IconKey,omitempty"`
	// The OSS path of the skill package. This parameter is reserved by the system and does not need to be specified.
	//
	// example:
	//
	// test/sk-test/current/skill.zip
	PackageOssKey *string `json:"PackageOssKey,omitempty" xml:"PackageOssKey,omitempty"`
	// The skill description.
	//
	// example:
	//
	// Current weather and forecasts with wttr.in via curl for locations, rain, temperature, travel planning.
	SkillDescription *string `json:"SkillDescription,omitempty" xml:"SkillDescription,omitempty"`
	// The skill name.
	//
	// example:
	//
	// weather-enhanced
	SkillName *string `json:"SkillName,omitempty" xml:"SkillName,omitempty"`
	// The OSS download URL of the skill package. This parameter is required for API calls.
	//
	// example:
	//
	// https://test.oss-cn-hangzhou.aliyuncs.com/weather_skill.zip?Expires=1788168088&OSSAccessKeyId=****
	SkillPackageUrl *string `json:"SkillPackageUrl,omitempty" xml:"SkillPackageUrl,omitempty"`
}

func (s CreateAgentSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentSkillRequest) GoString() string {
	return s.String()
}

func (s *CreateAgentSkillRequest) GetFileList() []*string {
	return s.FileList
}

func (s *CreateAgentSkillRequest) GetIconKey() *string {
	return s.IconKey
}

func (s *CreateAgentSkillRequest) GetPackageOssKey() *string {
	return s.PackageOssKey
}

func (s *CreateAgentSkillRequest) GetSkillDescription() *string {
	return s.SkillDescription
}

func (s *CreateAgentSkillRequest) GetSkillName() *string {
	return s.SkillName
}

func (s *CreateAgentSkillRequest) GetSkillPackageUrl() *string {
	return s.SkillPackageUrl
}

func (s *CreateAgentSkillRequest) SetFileList(v []*string) *CreateAgentSkillRequest {
	s.FileList = v
	return s
}

func (s *CreateAgentSkillRequest) SetIconKey(v string) *CreateAgentSkillRequest {
	s.IconKey = &v
	return s
}

func (s *CreateAgentSkillRequest) SetPackageOssKey(v string) *CreateAgentSkillRequest {
	s.PackageOssKey = &v
	return s
}

func (s *CreateAgentSkillRequest) SetSkillDescription(v string) *CreateAgentSkillRequest {
	s.SkillDescription = &v
	return s
}

func (s *CreateAgentSkillRequest) SetSkillName(v string) *CreateAgentSkillRequest {
	s.SkillName = &v
	return s
}

func (s *CreateAgentSkillRequest) SetSkillPackageUrl(v string) *CreateAgentSkillRequest {
	s.SkillPackageUrl = &v
	return s
}

func (s *CreateAgentSkillRequest) Validate() error {
	return dara.Validate(s)
}
