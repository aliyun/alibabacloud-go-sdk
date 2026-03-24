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
	// The ID of the request.
	//
	// example:
	//
	// 4F26D2F1-E288-5104-8518-05E240E337A4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array of protection templates.
	Templates []*DescribeDefenseTemplatesResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
	// The total number of protection templates returned.
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
	// The protection scenario. Valid values:
	//
	// - **waf_group**: Basic Protection.
	//
	// - **antiscan**: Scan Protection.
	//
	// - **ip_blacklist**: IP Blocklist.
	//
	// - **custom_acl**: Custom Rule.
	//
	// - **whitelist**: Allowlist.
	//
	// - **region_block**: Geographic Blocking.
	//
	// - **custom_response**: Custom Response.
	//
	// - **cc**: HTTP Flood Protection.
	//
	// - **tamperproof**: Webpage Tamper Protection.
	//
	// - **dlp**: Data Loss Prevention.
	//
	// - **bot_manager**: Bot Management.
	//
	// example:
	//
	// whitelist
	DefenseScene *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	// The sub-scenario for the Bot Management template. This parameter is returned only when `DefenseScene` is set to `bot_manager`. Valid values:
	//
	// - **web**: web protection
	//
	// - **app**: app protection
	//
	// - **basic**: basic protection
	//
	// - **bot_custom_acl**: The protection template for advanced Custom Rules in Bot Management.
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
	// The last modification time of the protection template. This value is a UNIX timestamp in milliseconds.
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
	// The origin of the protection template to be created. The value is custom, which indicates a user-defined template.
	//
	// example:
	//
	// custom
	TemplateOrigin *string `json:"TemplateOrigin,omitempty" xml:"TemplateOrigin,omitempty"`
	// The status of the protection template. Valid values:
	//
	// - **0**: Disabled.
	//
	// - **1**: Enabled.
	//
	// example:
	//
	// 1
	TemplateStatus *int32 `json:"TemplateStatus,omitempty" xml:"TemplateStatus,omitempty"`
	// The type of the protection template. Valid values:
	//
	// - **user_default**: The user\\"s default protection template.
	//
	// - **user_custom**: A custom protection template defined by the user.
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
