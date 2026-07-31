// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDefenseTemplatesResponseBody
	GetRequestId() *string
	SetTemplates(v []*DescribeDefenseTemplatesResponseBodyTemplates) *DescribeDefenseTemplatesResponseBody
	GetTemplates() []*DescribeDefenseTemplatesResponseBodyTemplates
	SetTotalCount(v int64) *DescribeDefenseTemplatesResponseBody
	GetTotalCount() *int64
}

type DescribeDefenseTemplatesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4F26D2F1-E288-5104-8518-05E240E337A4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of protection templates.
	Templates []*DescribeDefenseTemplatesResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDefenseTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDefenseTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDefenseTemplatesResponseBody) GetTemplates() []*DescribeDefenseTemplatesResponseBodyTemplates {
	return s.Templates
}

func (s *DescribeDefenseTemplatesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDefenseTemplatesResponseBody) SetRequestId(v string) *DescribeDefenseTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDefenseTemplatesResponseBody) SetTemplates(v []*DescribeDefenseTemplatesResponseBodyTemplates) *DescribeDefenseTemplatesResponseBody {
	s.Templates = v
	return s
}

func (s *DescribeDefenseTemplatesResponseBody) SetTotalCount(v int64) *DescribeDefenseTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDefenseTemplatesResponseBody) Validate() error {
	if s.Templates != nil {
		for _, item := range s.Templates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDefenseTemplatesResponseBodyTemplates struct {
	// The WAF protection scenario. Valid values:
	//
	// - **waf_group**: basic protection.
	//
	// - **antiscan**: scan protection.
	//
	// - **ip_blacklist**: IP blacklist.
	//
	// - **custom_acl**: custom rule.
	//
	// - **whitelist**: whitelist.
	//
	// - **region_block**: Location Blacklist.
	//
	// - **custom_response**: custom response.
	//
	// - **cc**: HTTP flood protection.
	//
	// - **tamperproof**: web tamper proofing.
	//
	// - **dlp**: data leak prevention.
	//
	// - **bot_manager**: new BOT management.
	//
	// example:
	//
	// whitelist
	DefenseScene *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	// The sub-scenario of the protection template. Valid values:
	//
	// - **web**: BOT management web protection scenario template.
	//
	// - **app**: BOT management app protection scenario template.
	//
	// - **basic**: BOT management basic protection template.
	//
	// - **bot_custom_acl**: BOT management advanced custom rule protection template.
	//
	// example:
	//
	// basic
	DefenseSubScene *string `json:"DefenseSubScene,omitempty" xml:"DefenseSubScene,omitempty"`
	// The description of the template.
	//
	// example:
	//
	// testTemplate
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The creation time of the protection template. The value is a timestamp in milliseconds.
	//
	// example:
	//
	// 1683776070000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the protection template.
	//
	// example:
	//
	// 56477
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the protection template.
	//
	// example:
	//
	// template-blockarea1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The source of the protection template. The value is custom, which indicates user-defined.
	//
	// example:
	//
	// custom
	TemplateOrigin *string `json:"TemplateOrigin,omitempty" xml:"TemplateOrigin,omitempty"`
	// The status of the protection template. Valid values:
	//
	// - **0**: disabled.
	//
	// - **1**: enabled.
	//
	// example:
	//
	// 1
	TemplateStatus *int32 `json:"TemplateStatus,omitempty" xml:"TemplateStatus,omitempty"`
	// The templatetype of the protection template. Valid values:
	//
	// - **user_default**: user default protection.
	//
	// - **user_custom**: user custom protection.
	//
	// example:
	//
	// user_custom
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s DescribeDefenseTemplatesResponseBodyTemplates) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *DescribeDefenseTemplatesResponseBodyTemplates) GetDefenseScene() *string {
	return s.DefenseScene
}

func (s *DescribeDefenseTemplatesResponseBodyTemplates) GetDefenseSubScene() *string {
	return s.DefenseSubScene
}

func (s *DescribeDefenseTemplatesResponseBodyTemplates) GetDescription() *string {
	return s.Description
}

func (s *DescribeDefenseTemplatesResponseBodyTemplates) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeDefenseTemplatesResponseBodyTemplates) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *DescribeDefenseTemplatesResponseBodyTemplates) GetTemplateName() *string {
	return s.TemplateName
}

func (s *DescribeDefenseTemplatesResponseBodyTemplates) GetTemplateOrigin() *string {
	return s.TemplateOrigin
}

func (s *DescribeDefenseTemplatesResponseBodyTemplates) GetTemplateStatus() *int32 {
	return s.TemplateStatus
}

func (s *DescribeDefenseTemplatesResponseBodyTemplates) GetTemplateType() *string {
	return s.TemplateType
}

func (s *DescribeDefenseTemplatesResponseBodyTemplates) SetDefenseScene(v string) *DescribeDefenseTemplatesResponseBodyTemplates {
	s.DefenseScene = &v
	return s
}

func (s *DescribeDefenseTemplatesResponseBodyTemplates) SetDefenseSubScene(v string) *DescribeDefenseTemplatesResponseBodyTemplates {
	s.DefenseSubScene = &v
	return s
}

func (s *DescribeDefenseTemplatesResponseBodyTemplates) SetDescription(v string) *DescribeDefenseTemplatesResponseBodyTemplates {
	s.Description = &v
	return s
}

func (s *DescribeDefenseTemplatesResponseBodyTemplates) SetGmtModified(v int64) *DescribeDefenseTemplatesResponseBodyTemplates {
	s.GmtModified = &v
	return s
}

func (s *DescribeDefenseTemplatesResponseBodyTemplates) SetTemplateId(v int64) *DescribeDefenseTemplatesResponseBodyTemplates {
	s.TemplateId = &v
	return s
}

func (s *DescribeDefenseTemplatesResponseBodyTemplates) SetTemplateName(v string) *DescribeDefenseTemplatesResponseBodyTemplates {
	s.TemplateName = &v
	return s
}

func (s *DescribeDefenseTemplatesResponseBodyTemplates) SetTemplateOrigin(v string) *DescribeDefenseTemplatesResponseBodyTemplates {
	s.TemplateOrigin = &v
	return s
}

func (s *DescribeDefenseTemplatesResponseBodyTemplates) SetTemplateStatus(v int32) *DescribeDefenseTemplatesResponseBodyTemplates {
	s.TemplateStatus = &v
	return s
}

func (s *DescribeDefenseTemplatesResponseBodyTemplates) SetTemplateType(v string) *DescribeDefenseTemplatesResponseBodyTemplates {
	s.TemplateType = &v
	return s
}

func (s *DescribeDefenseTemplatesResponseBodyTemplates) Validate() error {
	return dara.Validate(s)
}
