// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolarClawSkillTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *CreatePolarClawSkillTemplateRequest
	GetApplicationId() *string
	SetApplicationType(v string) *CreatePolarClawSkillTemplateRequest
	GetApplicationType() *string
	SetRegionId(v string) *CreatePolarClawSkillTemplateRequest
	GetRegionId() *string
	SetSkillTemplateDesc(v string) *CreatePolarClawSkillTemplateRequest
	GetSkillTemplateDesc() *string
	SetSkillTemplateName(v string) *CreatePolarClawSkillTemplateRequest
	GetSkillTemplateName() *string
	SetSkillTemplateParam(v string) *CreatePolarClawSkillTemplateRequest
	GetSkillTemplateParam() *string
}

type CreatePolarClawSkillTemplateRequest struct {
	// The application ID.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The application type.
	//
	// > If you are calling this operation for the first time, you do not need to specify this parameter.
	//
	// example:
	//
	// polardb_polarclaw
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query all available regions for the current account, including region IDs.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The description of the skill template.
	//
	// example:
	//
	// skill template for test.
	SkillTemplateDesc *string `json:"SkillTemplateDesc,omitempty" xml:"SkillTemplateDesc,omitempty"`
	// The name of the skill template.
	//
	// This parameter is required.
	//
	// example:
	//
	// skill_template_test
	SkillTemplateName *string `json:"SkillTemplateName,omitempty" xml:"SkillTemplateName,omitempty"`
	// The skill template parameters.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{\\"skillName\\":\\"polarclawhub\\",\\"version\\":\\"0.0.0\\"},{\\"skillName\\":\\"github\\",\\"version\\":\\"1.0.0\\"}]
	SkillTemplateParam *string `json:"SkillTemplateParam,omitempty" xml:"SkillTemplateParam,omitempty"`
}

func (s CreatePolarClawSkillTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePolarClawSkillTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreatePolarClawSkillTemplateRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *CreatePolarClawSkillTemplateRequest) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *CreatePolarClawSkillTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePolarClawSkillTemplateRequest) GetSkillTemplateDesc() *string {
	return s.SkillTemplateDesc
}

func (s *CreatePolarClawSkillTemplateRequest) GetSkillTemplateName() *string {
	return s.SkillTemplateName
}

func (s *CreatePolarClawSkillTemplateRequest) GetSkillTemplateParam() *string {
	return s.SkillTemplateParam
}

func (s *CreatePolarClawSkillTemplateRequest) SetApplicationId(v string) *CreatePolarClawSkillTemplateRequest {
	s.ApplicationId = &v
	return s
}

func (s *CreatePolarClawSkillTemplateRequest) SetApplicationType(v string) *CreatePolarClawSkillTemplateRequest {
	s.ApplicationType = &v
	return s
}

func (s *CreatePolarClawSkillTemplateRequest) SetRegionId(v string) *CreatePolarClawSkillTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePolarClawSkillTemplateRequest) SetSkillTemplateDesc(v string) *CreatePolarClawSkillTemplateRequest {
	s.SkillTemplateDesc = &v
	return s
}

func (s *CreatePolarClawSkillTemplateRequest) SetSkillTemplateName(v string) *CreatePolarClawSkillTemplateRequest {
	s.SkillTemplateName = &v
	return s
}

func (s *CreatePolarClawSkillTemplateRequest) SetSkillTemplateParam(v string) *CreatePolarClawSkillTemplateRequest {
	s.SkillTemplateParam = &v
	return s
}

func (s *CreatePolarClawSkillTemplateRequest) Validate() error {
	return dara.Validate(s)
}
