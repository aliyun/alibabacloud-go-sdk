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
	// The request ID.
	//
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The template information.
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
	if s.Template != nil {
		if err := s.Template.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDefenseTemplateResponseBodyTemplate struct {
	// The protection scenario. For more information, see the **DefenseScene*	- parameter in [CreateDefenseRule](~~CreateDefenseRule~~).
	//
	// example:
	//
	// waf_group
	DefenseScene *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	// The protection template sub-scenario. Valid values:
	//
	// - **web**: bot management web protection scenario template.
	//
	// - **app**: bot management app protection scenario template.
	//
	// - **basic**: bot management basic protection template.
	//
	// - **bot_custom_acl**: bot management advanced custom rule protection template.
	//
	// example:
	//
	// app
	DefenseSubScene *string `json:"DefenseSubScene,omitempty" xml:"DefenseSubScene,omitempty"`
	// The description of the protection template.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The detailed template information. For more information, see the Detail parameter in [CreateDefenseTemplate](https://help.aliyun.com/document_detail/461613.html).
	//
	// example:
	//
	// {"trafficFeature":"{\\"global\\":0,\\"excludeStatus\\":1,\\"conditions\\":[{\\"key\\":\\"URL\\",\\"opValue\\":\\"not-contain\\",\\"values\\":\\"test\\"}]}"}
	Detail map[string]interface{} `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The time when the protection template was last modified.
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
	// The template name.
	//
	// example:
	//
	// test0621
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The source of the protection template. The value custom indicates that the template is user-defined.
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
	// The protection templatetype. Valid values:
	//
	// - **user_default**: user default protection.
	//
	// - **user_custom**: user custom protection.
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

func (s *DescribeDefenseTemplateResponseBodyTemplate) GetDetail() map[string]interface{} {
	return s.Detail
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

func (s *DescribeDefenseTemplateResponseBodyTemplate) SetDetail(v map[string]interface{}) *DescribeDefenseTemplateResponseBodyTemplate {
	s.Detail = v
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
