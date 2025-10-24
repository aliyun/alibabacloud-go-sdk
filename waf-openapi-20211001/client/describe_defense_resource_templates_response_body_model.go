// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseResourceTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDefenseResourceTemplatesResponseBody
	GetRequestId() *string
	SetTemplates(v []*DescribeDefenseResourceTemplatesResponseBodyTemplates) *DescribeDefenseResourceTemplatesResponseBody
	GetTemplates() []*DescribeDefenseResourceTemplatesResponseBodyTemplates
}

type DescribeDefenseResourceTemplatesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2305CEB0-BA5A-5543-A1D3-3F1D0891****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The protection templates.
	Templates []*DescribeDefenseResourceTemplatesResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
}

func (s DescribeDefenseResourceTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseResourceTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDefenseResourceTemplatesResponseBody) GetTemplates() []*DescribeDefenseResourceTemplatesResponseBodyTemplates {
	return s.Templates
}

func (s *DescribeDefenseResourceTemplatesResponseBody) SetRequestId(v string) *DescribeDefenseResourceTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDefenseResourceTemplatesResponseBody) SetTemplates(v []*DescribeDefenseResourceTemplatesResponseBodyTemplates) *DescribeDefenseResourceTemplatesResponseBody {
	s.Templates = v
	return s
}

func (s *DescribeDefenseResourceTemplatesResponseBody) Validate() error {
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

type DescribeDefenseResourceTemplatesResponseBodyTemplates struct {
	// The scenario in which the protection template is used.
	//
	// 	- **waf_group**: basic protection.
	//
	// 	- **antiscan**: scan protection.
	//
	// 	- **ip_blacklist**: IP address blacklist.
	//
	// 	- **custom_acl**: custom rule.
	//
	// 	- **whitelist**: whitelist.
	//
	// 	- **region_block**: region blacklist.
	//
	// 	- **custom_response**: custom response.
	//
	// 	- **cc**: HTTP flood protection.
	//
	// 	- **tamperproof**: website tamper-proofing.
	//
	// 	- **dlp**: data leakage prevention.
	//
	// example:
	//
	// whitelist
	DefenseScene *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	// The sub-scenario in which the template is used. Valid values:
	//
	// 	- **web**: bot management for website protection.
	//
	// 	- **app**: bot management for app protection.
	//
	// 	- **basic**: bot management for basic protection.
	//
	// example:
	//
	// basic
	DefenseSubScene *string `json:"DefenseSubScene,omitempty" xml:"DefenseSubScene,omitempty"`
	// The description of the protection template.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the protection template was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1692930539000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the protection template.
	//
	// example:
	//
	// 12345
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the protection template.
	//
	// example:
	//
	// TestTemplateName
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The origin of the protection template. The value custom indicates that the template is a custom template created by the user.
	//
	// example:
	//
	// custom
	TemplateOrigin *string `json:"TemplateOrigin,omitempty" xml:"TemplateOrigin,omitempty"`
	// The status of the protection template. Valid values:
	//
	// 	- **0**: disabled.
	//
	// 	- **1**: enabled.
	//
	// example:
	//
	// 1
	TemplateStatus *int32 `json:"TemplateStatus,omitempty" xml:"TemplateStatus,omitempty"`
	// The type of the protection template. Valid values:
	//
	// 	- **user_default**: default template.
	//
	// 	- **user_custom**: custom template.
	//
	// example:
	//
	// user_custom
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s DescribeDefenseResourceTemplatesResponseBodyTemplates) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseResourceTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceTemplatesResponseBodyTemplates) GetDefenseScene() *string {
	return s.DefenseScene
}

func (s *DescribeDefenseResourceTemplatesResponseBodyTemplates) GetDefenseSubScene() *string {
	return s.DefenseSubScene
}

func (s *DescribeDefenseResourceTemplatesResponseBodyTemplates) GetDescription() *string {
	return s.Description
}

func (s *DescribeDefenseResourceTemplatesResponseBodyTemplates) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeDefenseResourceTemplatesResponseBodyTemplates) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *DescribeDefenseResourceTemplatesResponseBodyTemplates) GetTemplateName() *string {
	return s.TemplateName
}

func (s *DescribeDefenseResourceTemplatesResponseBodyTemplates) GetTemplateOrigin() *string {
	return s.TemplateOrigin
}

func (s *DescribeDefenseResourceTemplatesResponseBodyTemplates) GetTemplateStatus() *int32 {
	return s.TemplateStatus
}

func (s *DescribeDefenseResourceTemplatesResponseBodyTemplates) GetTemplateType() *string {
	return s.TemplateType
}

func (s *DescribeDefenseResourceTemplatesResponseBodyTemplates) SetDefenseScene(v string) *DescribeDefenseResourceTemplatesResponseBodyTemplates {
	s.DefenseScene = &v
	return s
}

func (s *DescribeDefenseResourceTemplatesResponseBodyTemplates) SetDefenseSubScene(v string) *DescribeDefenseResourceTemplatesResponseBodyTemplates {
	s.DefenseSubScene = &v
	return s
}

func (s *DescribeDefenseResourceTemplatesResponseBodyTemplates) SetDescription(v string) *DescribeDefenseResourceTemplatesResponseBodyTemplates {
	s.Description = &v
	return s
}

func (s *DescribeDefenseResourceTemplatesResponseBodyTemplates) SetGmtModified(v int64) *DescribeDefenseResourceTemplatesResponseBodyTemplates {
	s.GmtModified = &v
	return s
}

func (s *DescribeDefenseResourceTemplatesResponseBodyTemplates) SetTemplateId(v int64) *DescribeDefenseResourceTemplatesResponseBodyTemplates {
	s.TemplateId = &v
	return s
}

func (s *DescribeDefenseResourceTemplatesResponseBodyTemplates) SetTemplateName(v string) *DescribeDefenseResourceTemplatesResponseBodyTemplates {
	s.TemplateName = &v
	return s
}

func (s *DescribeDefenseResourceTemplatesResponseBodyTemplates) SetTemplateOrigin(v string) *DescribeDefenseResourceTemplatesResponseBodyTemplates {
	s.TemplateOrigin = &v
	return s
}

func (s *DescribeDefenseResourceTemplatesResponseBodyTemplates) SetTemplateStatus(v int32) *DescribeDefenseResourceTemplatesResponseBodyTemplates {
	s.TemplateStatus = &v
	return s
}

func (s *DescribeDefenseResourceTemplatesResponseBodyTemplates) SetTemplateType(v string) *DescribeDefenseResourceTemplatesResponseBodyTemplates {
	s.TemplateType = &v
	return s
}

func (s *DescribeDefenseResourceTemplatesResponseBodyTemplates) Validate() error {
	return dara.Validate(s)
}
