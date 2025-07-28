// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDefenseTemplateResponseBody
	GetRequestId() *string
	SetTemplate(v *DescribeDefenseTemplateResponseBodyTemplate) *DescribeDefenseTemplateResponseBody
	GetTemplate() *DescribeDefenseTemplateResponseBodyTemplate
}

type DescribeDefenseTemplateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the template.
	Template *DescribeDefenseTemplateResponseBodyTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
}

func (s DescribeDefenseTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDefenseTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDefenseTemplateResponseBody) GetTemplate() *DescribeDefenseTemplateResponseBodyTemplate {
	return s.Template
}

func (s *DescribeDefenseTemplateResponseBody) SetRequestId(v string) *DescribeDefenseTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDefenseTemplateResponseBody) SetTemplate(v *DescribeDefenseTemplateResponseBodyTemplate) *DescribeDefenseTemplateResponseBody {
	s.Template = v
	return s
}

func (s *DescribeDefenseTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDefenseTemplateResponseBodyTemplate struct {
	// The scenario in which the template is used. For more information, see the description of the **DefenseScene*	- parameter in the [CreateDefenseRule](~~CreateDefenseRule~~) topic.
	//
	// example:
	//
	// waf_group
	DefenseScene *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	// The sub-scenario in which the template is used. Valid values:
	//
	// 	- **web**: The template is a bot management template that is used for website protection.
	//
	// 	- **app**: The template is a bot management template that is used for app protection.
	//
	// 	- **basic**: The template is a bot management template that is used for basic protection.
	//
	// example:
	//
	// app
	DefenseSubScene *string `json:"DefenseSubScene,omitempty" xml:"DefenseSubScene,omitempty"`
	// The description of the protection rule template.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The most recent time when the protection rule template was modified.
	//
	// example:
	//
	// 1665283642000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the protection rule template.
	//
	// example:
	//
	// 10097
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the protection rule template.
	//
	// example:
	//
	// test0621
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The origin of the protection rule template. If the value of this parameter is custom, the protection rule template is created by the user.
	//
	// example:
	//
	// custom
	TemplateOrigin *string `json:"TemplateOrigin,omitempty" xml:"TemplateOrigin,omitempty"`
	// The status of the protection rule template. Valid values:
	//
	// 	- **0:*	- disabled.
	//
	// 	- **1:*	- enabled.
	//
	// example:
	//
	// 1
	TemplateStatus *int32 `json:"TemplateStatus,omitempty" xml:"TemplateStatus,omitempty"`
	// The type of the protection rule template. Valid values:
	//
	// 	- **user_default:*	- default template.
	//
	// 	- **user_custom:*	- custom template.
	//
	// example:
	//
	// user_default
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s DescribeDefenseTemplateResponseBodyTemplate) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseTemplateResponseBodyTemplate) GoString() string {
	return s.String()
}

func (s *DescribeDefenseTemplateResponseBodyTemplate) GetDefenseScene() *string {
	return s.DefenseScene
}

func (s *DescribeDefenseTemplateResponseBodyTemplate) GetDefenseSubScene() *string {
	return s.DefenseSubScene
}

func (s *DescribeDefenseTemplateResponseBodyTemplate) GetDescription() *string {
	return s.Description
}

func (s *DescribeDefenseTemplateResponseBodyTemplate) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeDefenseTemplateResponseBodyTemplate) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *DescribeDefenseTemplateResponseBodyTemplate) GetTemplateName() *string {
	return s.TemplateName
}

func (s *DescribeDefenseTemplateResponseBodyTemplate) GetTemplateOrigin() *string {
	return s.TemplateOrigin
}

func (s *DescribeDefenseTemplateResponseBodyTemplate) GetTemplateStatus() *int32 {
	return s.TemplateStatus
}

func (s *DescribeDefenseTemplateResponseBodyTemplate) GetTemplateType() *string {
	return s.TemplateType
}

func (s *DescribeDefenseTemplateResponseBodyTemplate) SetDefenseScene(v string) *DescribeDefenseTemplateResponseBodyTemplate {
	s.DefenseScene = &v
	return s
}

func (s *DescribeDefenseTemplateResponseBodyTemplate) SetDefenseSubScene(v string) *DescribeDefenseTemplateResponseBodyTemplate {
	s.DefenseSubScene = &v
	return s
}

func (s *DescribeDefenseTemplateResponseBodyTemplate) SetDescription(v string) *DescribeDefenseTemplateResponseBodyTemplate {
	s.Description = &v
	return s
}

func (s *DescribeDefenseTemplateResponseBodyTemplate) SetGmtModified(v int64) *DescribeDefenseTemplateResponseBodyTemplate {
	s.GmtModified = &v
	return s
}

func (s *DescribeDefenseTemplateResponseBodyTemplate) SetTemplateId(v int64) *DescribeDefenseTemplateResponseBodyTemplate {
	s.TemplateId = &v
	return s
}

func (s *DescribeDefenseTemplateResponseBodyTemplate) SetTemplateName(v string) *DescribeDefenseTemplateResponseBodyTemplate {
	s.TemplateName = &v
	return s
}

func (s *DescribeDefenseTemplateResponseBodyTemplate) SetTemplateOrigin(v string) *DescribeDefenseTemplateResponseBodyTemplate {
	s.TemplateOrigin = &v
	return s
}

func (s *DescribeDefenseTemplateResponseBodyTemplate) SetTemplateStatus(v int32) *DescribeDefenseTemplateResponseBodyTemplate {
	s.TemplateStatus = &v
	return s
}

func (s *DescribeDefenseTemplateResponseBodyTemplate) SetTemplateType(v string) *DescribeDefenseTemplateResponseBodyTemplate {
	s.TemplateType = &v
	return s
}

func (s *DescribeDefenseTemplateResponseBodyTemplate) Validate() error {
	return dara.Validate(s)
}
