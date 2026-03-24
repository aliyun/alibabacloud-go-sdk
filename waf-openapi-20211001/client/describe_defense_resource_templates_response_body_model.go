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
	// 2305CEB0-BA5A-5543-A1D3-3F1D08911B1C
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
	// The protection scenario. For more information, see the **DefenseScene*	- parameter in [CreateDefenseRule](https://help.aliyun.com/document_detail/461421.html).
	//
	// example:
	//
	// whitelist
	DefenseScene *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	// The sub-scenario of the template. Valid values:
	//
	// - **web**: The bot management template for web protection.
	//
	// - **app**: The bot management template for app protection.
	//
	// - **basic**: The basic bot management template.
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
	// The time when the protection template was created, in UNIX timestamp format. Unit: milliseconds.
	//
	// example:
	//
	// 1692930539000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the protection template.
	//
	// example:
	//
	// 34328
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the protection template.
	//
	// example:
	//
	// antiscanTemplate
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The origin of the protection template. The value custom indicates a user-created template.
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
	// The type of the template. Valid values:
	//
	// - **user_default**: The default template for the user.
	//
	// - **user_custom**: A custom template created by the user.
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
