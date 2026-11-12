// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarClawSkillAssessmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DescribePolarClawSkillAssessmentRequest
	GetApplicationId() *string
	SetRegionId(v string) *DescribePolarClawSkillAssessmentRequest
	GetRegionId() *string
	SetSkillName(v string) *DescribePolarClawSkillAssessmentRequest
	GetSkillName() *string
	SetSkillVersion(v string) *DescribePolarClawSkillAssessmentRequest
	GetSkillVersion() *string
	SetSource(v string) *DescribePolarClawSkillAssessmentRequest
	GetSource() *string
}

type DescribePolarClawSkillAssessmentRequest struct {
	// The application ID.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query all available regions for the current account, including region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the skill. This parameter is required. If it is not provided, the API returns an InvalidParameter.SkillName error (HTTP 400).
	//
	// example:
	//
	// polarclawhub
	SkillName *string `json:"SkillName,omitempty" xml:"SkillName,omitempty"`
	// The version of the skill. This parameter is required. If it is not provided, the API returns an InvalidParameter.SkillVersion error (HTTP 400).
	//
	// example:
	//
	// 1.0.0
	SkillVersion *string `json:"SkillVersion,omitempty" xml:"SkillVersion,omitempty"`
	// The source of the skill. Valid values: PUBLIC, SELF, and OTHER. This parameter is required. If it is not provided, the API returns an InvalidParameter.Source error (HTTP 400).
	//
	// example:
	//
	// PUBLIC
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s DescribePolarClawSkillAssessmentRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawSkillAssessmentRequest) GoString() string {
	return s.String()
}

func (s *DescribePolarClawSkillAssessmentRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribePolarClawSkillAssessmentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePolarClawSkillAssessmentRequest) GetSkillName() *string {
	return s.SkillName
}

func (s *DescribePolarClawSkillAssessmentRequest) GetSkillVersion() *string {
	return s.SkillVersion
}

func (s *DescribePolarClawSkillAssessmentRequest) GetSource() *string {
	return s.Source
}

func (s *DescribePolarClawSkillAssessmentRequest) SetApplicationId(v string) *DescribePolarClawSkillAssessmentRequest {
	s.ApplicationId = &v
	return s
}

func (s *DescribePolarClawSkillAssessmentRequest) SetRegionId(v string) *DescribePolarClawSkillAssessmentRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePolarClawSkillAssessmentRequest) SetSkillName(v string) *DescribePolarClawSkillAssessmentRequest {
	s.SkillName = &v
	return s
}

func (s *DescribePolarClawSkillAssessmentRequest) SetSkillVersion(v string) *DescribePolarClawSkillAssessmentRequest {
	s.SkillVersion = &v
	return s
}

func (s *DescribePolarClawSkillAssessmentRequest) SetSource(v string) *DescribePolarClawSkillAssessmentRequest {
	s.Source = &v
	return s
}

func (s *DescribePolarClawSkillAssessmentRequest) Validate() error {
	return dara.Validate(s)
}
