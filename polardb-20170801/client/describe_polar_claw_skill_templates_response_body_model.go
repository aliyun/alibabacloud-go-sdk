// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarClawSkillTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v string) *DescribePolarClawSkillTemplatesResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *DescribePolarClawSkillTemplatesResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DescribePolarClawSkillTemplatesResponseBody
	GetRequestId() *string
	SetSkillTemplates(v []*DescribePolarClawSkillTemplatesResponseBodySkillTemplates) *DescribePolarClawSkillTemplatesResponseBody
	GetSkillTemplates() []*DescribePolarClawSkillTemplatesResponseBodySkillTemplates
	SetTotalCount(v string) *DescribePolarClawSkillTemplatesResponseBody
	GetTotalCount() *string
}

type DescribePolarClawSkillTemplatesResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// - **30**
	//
	// - **50**
	//
	// - **100**
	//
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 925B84D9-CA72-432C-95CF-738C22******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The skill template group.
	SkillTemplates []*DescribePolarClawSkillTemplatesResponseBodySkillTemplates `json:"SkillTemplates,omitempty" xml:"SkillTemplates,omitempty" type:"Repeated"`
	// The total number of entries. This parameter is returned only when you query a list.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePolarClawSkillTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawSkillTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolarClawSkillTemplatesResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribePolarClawSkillTemplatesResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribePolarClawSkillTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePolarClawSkillTemplatesResponseBody) GetSkillTemplates() []*DescribePolarClawSkillTemplatesResponseBodySkillTemplates {
	return s.SkillTemplates
}

func (s *DescribePolarClawSkillTemplatesResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribePolarClawSkillTemplatesResponseBody) SetPageNumber(v string) *DescribePolarClawSkillTemplatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePolarClawSkillTemplatesResponseBody) SetPageSize(v string) *DescribePolarClawSkillTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePolarClawSkillTemplatesResponseBody) SetRequestId(v string) *DescribePolarClawSkillTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePolarClawSkillTemplatesResponseBody) SetSkillTemplates(v []*DescribePolarClawSkillTemplatesResponseBodySkillTemplates) *DescribePolarClawSkillTemplatesResponseBody {
	s.SkillTemplates = v
	return s
}

func (s *DescribePolarClawSkillTemplatesResponseBody) SetTotalCount(v string) *DescribePolarClawSkillTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePolarClawSkillTemplatesResponseBody) Validate() error {
	if s.SkillTemplates != nil {
		for _, item := range s.SkillTemplates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePolarClawSkillTemplatesResponseBodySkillTemplates struct {
	// The time when the skill template was created in UTC.
	//
	// example:
	//
	// 2022-11-15T16:00Z
	Created *string `json:"Created,omitempty" xml:"Created,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// 423
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The time when the skill template was last modified in UTC.
	//
	// example:
	//
	// 2022-11-16T16:00Z
	Modified *string `json:"Modified,omitempty" xml:"Modified,omitempty"`
	// The number of skills.
	//
	// example:
	//
	// 10
	SkillCounts *string `json:"SkillCounts,omitempty" xml:"SkillCounts,omitempty"`
	// The skill template description.
	//
	// example:
	//
	// skill template for test
	SkillTemplateDesc *string `json:"SkillTemplateDesc,omitempty" xml:"SkillTemplateDesc,omitempty"`
	// The skill template ID.
	//
	// example:
	//
	// pcpg-skill-2ze5v8n27o7x92a
	SkillTemplateId *string `json:"SkillTemplateId,omitempty" xml:"SkillTemplateId,omitempty"`
	// The skill template name.
	//
	// example:
	//
	// skill-template-test
	SkillTemplateName *string `json:"SkillTemplateName,omitempty" xml:"SkillTemplateName,omitempty"`
	// The skill templatetype. Valid values: 0: system template. 1: user template.
	//
	// example:
	//
	// 1
	SkillTemplateType *string `json:"SkillTemplateType,omitempty" xml:"SkillTemplateType,omitempty"`
	// The source application ID.
	//
	// example:
	//
	// pa-**************
	SrcApplicationId *string `json:"SrcApplicationId,omitempty" xml:"SrcApplicationId,omitempty"`
}

func (s DescribePolarClawSkillTemplatesResponseBodySkillTemplates) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawSkillTemplatesResponseBodySkillTemplates) GoString() string {
	return s.String()
}

func (s *DescribePolarClawSkillTemplatesResponseBodySkillTemplates) GetCreated() *string {
	return s.Created
}

func (s *DescribePolarClawSkillTemplatesResponseBodySkillTemplates) GetId() *string {
	return s.Id
}

func (s *DescribePolarClawSkillTemplatesResponseBodySkillTemplates) GetModified() *string {
	return s.Modified
}

func (s *DescribePolarClawSkillTemplatesResponseBodySkillTemplates) GetSkillCounts() *string {
	return s.SkillCounts
}

func (s *DescribePolarClawSkillTemplatesResponseBodySkillTemplates) GetSkillTemplateDesc() *string {
	return s.SkillTemplateDesc
}

func (s *DescribePolarClawSkillTemplatesResponseBodySkillTemplates) GetSkillTemplateId() *string {
	return s.SkillTemplateId
}

func (s *DescribePolarClawSkillTemplatesResponseBodySkillTemplates) GetSkillTemplateName() *string {
	return s.SkillTemplateName
}

func (s *DescribePolarClawSkillTemplatesResponseBodySkillTemplates) GetSkillTemplateType() *string {
	return s.SkillTemplateType
}

func (s *DescribePolarClawSkillTemplatesResponseBodySkillTemplates) GetSrcApplicationId() *string {
	return s.SrcApplicationId
}

func (s *DescribePolarClawSkillTemplatesResponseBodySkillTemplates) SetCreated(v string) *DescribePolarClawSkillTemplatesResponseBodySkillTemplates {
	s.Created = &v
	return s
}

func (s *DescribePolarClawSkillTemplatesResponseBodySkillTemplates) SetId(v string) *DescribePolarClawSkillTemplatesResponseBodySkillTemplates {
	s.Id = &v
	return s
}

func (s *DescribePolarClawSkillTemplatesResponseBodySkillTemplates) SetModified(v string) *DescribePolarClawSkillTemplatesResponseBodySkillTemplates {
	s.Modified = &v
	return s
}

func (s *DescribePolarClawSkillTemplatesResponseBodySkillTemplates) SetSkillCounts(v string) *DescribePolarClawSkillTemplatesResponseBodySkillTemplates {
	s.SkillCounts = &v
	return s
}

func (s *DescribePolarClawSkillTemplatesResponseBodySkillTemplates) SetSkillTemplateDesc(v string) *DescribePolarClawSkillTemplatesResponseBodySkillTemplates {
	s.SkillTemplateDesc = &v
	return s
}

func (s *DescribePolarClawSkillTemplatesResponseBodySkillTemplates) SetSkillTemplateId(v string) *DescribePolarClawSkillTemplatesResponseBodySkillTemplates {
	s.SkillTemplateId = &v
	return s
}

func (s *DescribePolarClawSkillTemplatesResponseBodySkillTemplates) SetSkillTemplateName(v string) *DescribePolarClawSkillTemplatesResponseBodySkillTemplates {
	s.SkillTemplateName = &v
	return s
}

func (s *DescribePolarClawSkillTemplatesResponseBodySkillTemplates) SetSkillTemplateType(v string) *DescribePolarClawSkillTemplatesResponseBodySkillTemplates {
	s.SkillTemplateType = &v
	return s
}

func (s *DescribePolarClawSkillTemplatesResponseBodySkillTemplates) SetSrcApplicationId(v string) *DescribePolarClawSkillTemplatesResponseBodySkillTemplates {
	s.SrcApplicationId = &v
	return s
}

func (s *DescribePolarClawSkillTemplatesResponseBodySkillTemplates) Validate() error {
	return dara.Validate(s)
}
