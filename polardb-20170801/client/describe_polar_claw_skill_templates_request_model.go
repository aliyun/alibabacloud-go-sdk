// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarClawSkillTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DescribePolarClawSkillTemplatesRequest
	GetApplicationId() *string
	SetApplicationType(v string) *DescribePolarClawSkillTemplatesRequest
	GetApplicationType() *string
	SetEndTime(v string) *DescribePolarClawSkillTemplatesRequest
	GetEndTime() *string
	SetPageNumber(v int32) *DescribePolarClawSkillTemplatesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribePolarClawSkillTemplatesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribePolarClawSkillTemplatesRequest
	GetRegionId() *string
	SetSkillTemplateId(v string) *DescribePolarClawSkillTemplatesRequest
	GetSkillTemplateId() *string
	SetStartTime(v string) *DescribePolarClawSkillTemplatesRequest
	GetStartTime() *string
}

type DescribePolarClawSkillTemplatesRequest struct {
	// The application ID.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The application type.
	//
	// > If this is your first time calling this operation, you do not need to specify this parameter.
	//
	// example:
	//
	// supabase
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// The end time of the query in UTC.
	//
	// example:
	//
	// 2022-11-16T16:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to view all available regions for the current account, including region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The skill template ID.
	//
	// example:
	//
	// pcpg-skill-2ze5v8n27o7x92a
	SkillTemplateId *string `json:"SkillTemplateId,omitempty" xml:"SkillTemplateId,omitempty"`
	// The start time of the query in UTC.
	//
	// example:
	//
	// 2022-11-15T16:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribePolarClawSkillTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawSkillTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribePolarClawSkillTemplatesRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribePolarClawSkillTemplatesRequest) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *DescribePolarClawSkillTemplatesRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribePolarClawSkillTemplatesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePolarClawSkillTemplatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePolarClawSkillTemplatesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePolarClawSkillTemplatesRequest) GetSkillTemplateId() *string {
	return s.SkillTemplateId
}

func (s *DescribePolarClawSkillTemplatesRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribePolarClawSkillTemplatesRequest) SetApplicationId(v string) *DescribePolarClawSkillTemplatesRequest {
	s.ApplicationId = &v
	return s
}

func (s *DescribePolarClawSkillTemplatesRequest) SetApplicationType(v string) *DescribePolarClawSkillTemplatesRequest {
	s.ApplicationType = &v
	return s
}

func (s *DescribePolarClawSkillTemplatesRequest) SetEndTime(v string) *DescribePolarClawSkillTemplatesRequest {
	s.EndTime = &v
	return s
}

func (s *DescribePolarClawSkillTemplatesRequest) SetPageNumber(v int32) *DescribePolarClawSkillTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePolarClawSkillTemplatesRequest) SetPageSize(v int32) *DescribePolarClawSkillTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePolarClawSkillTemplatesRequest) SetRegionId(v string) *DescribePolarClawSkillTemplatesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePolarClawSkillTemplatesRequest) SetSkillTemplateId(v string) *DescribePolarClawSkillTemplatesRequest {
	s.SkillTemplateId = &v
	return s
}

func (s *DescribePolarClawSkillTemplatesRequest) SetStartTime(v string) *DescribePolarClawSkillTemplatesRequest {
	s.StartTime = &v
	return s
}

func (s *DescribePolarClawSkillTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
