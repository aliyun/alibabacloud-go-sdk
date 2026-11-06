// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarClawSkillTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribePolarClawSkillTemplateResponseBody
	GetRequestId() *string
	SetSkillTemplates(v []*DescribePolarClawSkillTemplateResponseBodySkillTemplates) *DescribePolarClawSkillTemplateResponseBody
	GetSkillTemplates() []*DescribePolarClawSkillTemplateResponseBodySkillTemplates
}

type DescribePolarClawSkillTemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CDB3258F-B5DE-43C4-8935-CBA0CA******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The skill template group.
	SkillTemplates []*DescribePolarClawSkillTemplateResponseBodySkillTemplates `json:"SkillTemplates,omitempty" xml:"SkillTemplates,omitempty" type:"Repeated"`
}

func (s DescribePolarClawSkillTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawSkillTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolarClawSkillTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePolarClawSkillTemplateResponseBody) GetSkillTemplates() []*DescribePolarClawSkillTemplateResponseBodySkillTemplates {
	return s.SkillTemplates
}

func (s *DescribePolarClawSkillTemplateResponseBody) SetRequestId(v string) *DescribePolarClawSkillTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePolarClawSkillTemplateResponseBody) SetSkillTemplates(v []*DescribePolarClawSkillTemplateResponseBodySkillTemplates) *DescribePolarClawSkillTemplateResponseBody {
	s.SkillTemplates = v
	return s
}

func (s *DescribePolarClawSkillTemplateResponseBody) Validate() error {
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

type DescribePolarClawSkillTemplateResponseBodySkillTemplates struct {
	// Indicates whether the template is actually created.
	//
	// example:
	//
	// 2026-03-21 12:00:00
	Created *string `json:"Created,omitempty" xml:"Created,omitempty"`
	// The last modification time of the parameter template.
	//
	// example:
	//
	// 2026-04-23 23:00:00
	Modified *string `json:"Modified,omitempty" xml:"Modified,omitempty"`
	// The description of the skill template.
	//
	// example:
	//
	// skill template for test
	SkillTemplateDesc *string `json:"SkillTemplateDesc,omitempty" xml:"SkillTemplateDesc,omitempty"`
	// The skill template ID.
	//
	// example:
	//
	// pcpg-skill-**************
	SkillTemplateId *string `json:"SkillTemplateId,omitempty" xml:"SkillTemplateId,omitempty"`
	// The name of the skill template.
	//
	// example:
	//
	// skill-template-test
	SkillTemplateName *string `json:"SkillTemplateName,omitempty" xml:"SkillTemplateName,omitempty"`
	// The map of associated skill details. The key is the skill identifier, and the value is a JSON string that can be parsed into an array of skill details. Each array element contains skillName and version.
	//
	// example:
	//
	// {"skill1":"[{\\"skillName\\":\\"example-skill\\",\\"version\\":\\"1.0.0\\"}]"}
	Skills map[string]*string `json:"Skills,omitempty" xml:"Skills,omitempty"`
	// The source application ID.
	//
	// example:
	//
	// pa-**************
	SrcApplicationId *string `json:"SrcApplicationId,omitempty" xml:"SrcApplicationId,omitempty"`
}

func (s DescribePolarClawSkillTemplateResponseBodySkillTemplates) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawSkillTemplateResponseBodySkillTemplates) GoString() string {
	return s.String()
}

func (s *DescribePolarClawSkillTemplateResponseBodySkillTemplates) GetCreated() *string {
	return s.Created
}

func (s *DescribePolarClawSkillTemplateResponseBodySkillTemplates) GetModified() *string {
	return s.Modified
}

func (s *DescribePolarClawSkillTemplateResponseBodySkillTemplates) GetSkillTemplateDesc() *string {
	return s.SkillTemplateDesc
}

func (s *DescribePolarClawSkillTemplateResponseBodySkillTemplates) GetSkillTemplateId() *string {
	return s.SkillTemplateId
}

func (s *DescribePolarClawSkillTemplateResponseBodySkillTemplates) GetSkillTemplateName() *string {
	return s.SkillTemplateName
}

func (s *DescribePolarClawSkillTemplateResponseBodySkillTemplates) GetSkills() map[string]*string {
	return s.Skills
}

func (s *DescribePolarClawSkillTemplateResponseBodySkillTemplates) GetSrcApplicationId() *string {
	return s.SrcApplicationId
}

func (s *DescribePolarClawSkillTemplateResponseBodySkillTemplates) SetCreated(v string) *DescribePolarClawSkillTemplateResponseBodySkillTemplates {
	s.Created = &v
	return s
}

func (s *DescribePolarClawSkillTemplateResponseBodySkillTemplates) SetModified(v string) *DescribePolarClawSkillTemplateResponseBodySkillTemplates {
	s.Modified = &v
	return s
}

func (s *DescribePolarClawSkillTemplateResponseBodySkillTemplates) SetSkillTemplateDesc(v string) *DescribePolarClawSkillTemplateResponseBodySkillTemplates {
	s.SkillTemplateDesc = &v
	return s
}

func (s *DescribePolarClawSkillTemplateResponseBodySkillTemplates) SetSkillTemplateId(v string) *DescribePolarClawSkillTemplateResponseBodySkillTemplates {
	s.SkillTemplateId = &v
	return s
}

func (s *DescribePolarClawSkillTemplateResponseBodySkillTemplates) SetSkillTemplateName(v string) *DescribePolarClawSkillTemplateResponseBodySkillTemplates {
	s.SkillTemplateName = &v
	return s
}

func (s *DescribePolarClawSkillTemplateResponseBodySkillTemplates) SetSkills(v map[string]*string) *DescribePolarClawSkillTemplateResponseBodySkillTemplates {
	s.Skills = v
	return s
}

func (s *DescribePolarClawSkillTemplateResponseBodySkillTemplates) SetSrcApplicationId(v string) *DescribePolarClawSkillTemplateResponseBodySkillTemplates {
	s.SrcApplicationId = &v
	return s
}

func (s *DescribePolarClawSkillTemplateResponseBodySkillTemplates) Validate() error {
	return dara.Validate(s)
}
