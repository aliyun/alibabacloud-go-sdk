// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarClawSkillTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DescribePolarClawSkillTemplateRequest
	GetApplicationId() *string
	SetApplicationType(v string) *DescribePolarClawSkillTemplateRequest
	GetApplicationType() *string
	SetRegionId(v string) *DescribePolarClawSkillTemplateRequest
	GetRegionId() *string
	SetSkillTemplateId(v string) *DescribePolarClawSkillTemplateRequest
	GetSkillTemplateId() *string
}

type DescribePolarClawSkillTemplateRequest struct {
	// The application ID. This parameter is required when you call this operation. If this parameter is not specified, InvalidApplicationId.Malformed is returned.
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
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query available regions.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The skill template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pcpg-skill-**************
	SkillTemplateId *string `json:"SkillTemplateId,omitempty" xml:"SkillTemplateId,omitempty"`
}

func (s DescribePolarClawSkillTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawSkillTemplateRequest) GoString() string {
	return s.String()
}

func (s *DescribePolarClawSkillTemplateRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribePolarClawSkillTemplateRequest) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *DescribePolarClawSkillTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePolarClawSkillTemplateRequest) GetSkillTemplateId() *string {
	return s.SkillTemplateId
}

func (s *DescribePolarClawSkillTemplateRequest) SetApplicationId(v string) *DescribePolarClawSkillTemplateRequest {
	s.ApplicationId = &v
	return s
}

func (s *DescribePolarClawSkillTemplateRequest) SetApplicationType(v string) *DescribePolarClawSkillTemplateRequest {
	s.ApplicationType = &v
	return s
}

func (s *DescribePolarClawSkillTemplateRequest) SetRegionId(v string) *DescribePolarClawSkillTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePolarClawSkillTemplateRequest) SetSkillTemplateId(v string) *DescribePolarClawSkillTemplateRequest {
	s.SkillTemplateId = &v
	return s
}

func (s *DescribePolarClawSkillTemplateRequest) Validate() error {
	return dara.Validate(s)
}
