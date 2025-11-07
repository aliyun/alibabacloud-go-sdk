// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExecutionTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *GetExecutionTemplateResponseBody
	GetContent() *string
	SetRequestId(v string) *GetExecutionTemplateResponseBody
	GetRequestId() *string
	SetTemplate(v *GetExecutionTemplateResponseBodyTemplate) *GetExecutionTemplateResponseBody
	GetTemplate() *GetExecutionTemplateResponseBodyTemplate
}

type GetExecutionTemplateResponseBody struct {
	// The content of the template.
	//
	// example:
	//
	// "{\\n \\"FormatVersion\\": \\"OOS-2019-06-01\\",\\n \\"Parameters\\": {\\n \\"Status\\": {\\n \\"Type\\": \\"String\\",\\n \\"Description\\": \\"(Required) The ID of the ECS instance.\\"\\n }\\n },\\n \\"Tasks\\": [\\n {\\n \\"Name\\": \\"bar\\",\\n \\"Properties\\": {\\n \\"Parameters\\": {\\n \\"Status\\": \\"{{ Status }}\\"\\n },\\n \\"API\\": \\"DescribeInstances\\",\\n \\"Service\\": \\"Ecs\\"\\n },\\n \\"Action\\": \\"acs::ExecuteAPI\\",\\n \\"Outputs\\": {\\n \\"InstanceIds\\", {\\n \\"ValueSelector\\": \\".Instances.Instance[].InstanceId\\",\\n \\"Type\\": \\"List\\"\\n }\\n }\\n }\\n ],\\n \\"Outputs\\": {\\n \\"InstanceIds\\": {\\n \\"Value\\": \\" {{ bar.InstanceIds }} \\",\\n \\"Type\\": \\"List\\"\\n }\\n }\\n}\\n"
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 14A60-EBE7-47CA-9757-12C1D47A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The metadata of the template.
	Template *GetExecutionTemplateResponseBodyTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
}

func (s GetExecutionTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetExecutionTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetExecutionTemplateResponseBody) GetContent() *string {
	return s.Content
}

func (s *GetExecutionTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetExecutionTemplateResponseBody) GetTemplate() *GetExecutionTemplateResponseBodyTemplate {
	return s.Template
}

func (s *GetExecutionTemplateResponseBody) SetContent(v string) *GetExecutionTemplateResponseBody {
	s.Content = &v
	return s
}

func (s *GetExecutionTemplateResponseBody) SetRequestId(v string) *GetExecutionTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExecutionTemplateResponseBody) SetTemplate(v *GetExecutionTemplateResponseBodyTemplate) *GetExecutionTemplateResponseBody {
	s.Template = v
	return s
}

func (s *GetExecutionTemplateResponseBody) Validate() error {
	if s.Template != nil {
		if err := s.Template.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetExecutionTemplateResponseBodyTemplate struct {
	// The creator of the template.
	//
	// example:
	//
	// root(13090000)
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The time when the template was created.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The description of the template.
	//
	// example:
	//
	// Get status of instances
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The SHA-256 value of the template content.
	//
	// example:
	//
	// 4bc7d7a21b3e003434b9c223f6e6d2578b5ebfeb5be28c1fcf8a8a1b11907bb4
	Hash *string `json:"Hash,omitempty" xml:"Hash,omitempty"`
	// The share type of the template. The share type of a user-created template is **Private**.
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The tag keys and values. The number of key-value pairs ranges from 1 to 20.
	//
	// example:
	//
	// {"k1":"k2","k2":"v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The format of the template. The system automatically determines whether the format is JSON or YAML.
	//
	// example:
	//
	// JSON
	TemplateFormat *string `json:"TemplateFormat,omitempty" xml:"TemplateFormat,omitempty"`
	// The ID of the template.
	//
	// example:
	//
	// t-94753d4d828d38
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// MyTemplate
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The version of the template. The name of the version consists of the letter v and a number. The number starts from 1.
	//
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// The user who last updated the template.
	//
	// example:
	//
	// root(13090000)
	UpdatedBy *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	// The time when the template was last updated.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s GetExecutionTemplateResponseBodyTemplate) String() string {
	return dara.Prettify(s)
}

func (s GetExecutionTemplateResponseBodyTemplate) GoString() string {
	return s.String()
}

func (s *GetExecutionTemplateResponseBodyTemplate) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *GetExecutionTemplateResponseBodyTemplate) GetCreatedDate() *string {
	return s.CreatedDate
}

func (s *GetExecutionTemplateResponseBodyTemplate) GetDescription() *string {
	return s.Description
}

func (s *GetExecutionTemplateResponseBodyTemplate) GetHash() *string {
	return s.Hash
}

func (s *GetExecutionTemplateResponseBodyTemplate) GetShareType() *string {
	return s.ShareType
}

func (s *GetExecutionTemplateResponseBodyTemplate) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *GetExecutionTemplateResponseBodyTemplate) GetTemplateFormat() *string {
	return s.TemplateFormat
}

func (s *GetExecutionTemplateResponseBodyTemplate) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetExecutionTemplateResponseBodyTemplate) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetExecutionTemplateResponseBodyTemplate) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *GetExecutionTemplateResponseBodyTemplate) GetUpdatedBy() *string {
	return s.UpdatedBy
}

func (s *GetExecutionTemplateResponseBodyTemplate) GetUpdatedDate() *string {
	return s.UpdatedDate
}

func (s *GetExecutionTemplateResponseBodyTemplate) SetCreatedBy(v string) *GetExecutionTemplateResponseBodyTemplate {
	s.CreatedBy = &v
	return s
}

func (s *GetExecutionTemplateResponseBodyTemplate) SetCreatedDate(v string) *GetExecutionTemplateResponseBodyTemplate {
	s.CreatedDate = &v
	return s
}

func (s *GetExecutionTemplateResponseBodyTemplate) SetDescription(v string) *GetExecutionTemplateResponseBodyTemplate {
	s.Description = &v
	return s
}

func (s *GetExecutionTemplateResponseBodyTemplate) SetHash(v string) *GetExecutionTemplateResponseBodyTemplate {
	s.Hash = &v
	return s
}

func (s *GetExecutionTemplateResponseBodyTemplate) SetShareType(v string) *GetExecutionTemplateResponseBodyTemplate {
	s.ShareType = &v
	return s
}

func (s *GetExecutionTemplateResponseBodyTemplate) SetTags(v map[string]interface{}) *GetExecutionTemplateResponseBodyTemplate {
	s.Tags = v
	return s
}

func (s *GetExecutionTemplateResponseBodyTemplate) SetTemplateFormat(v string) *GetExecutionTemplateResponseBodyTemplate {
	s.TemplateFormat = &v
	return s
}

func (s *GetExecutionTemplateResponseBodyTemplate) SetTemplateId(v string) *GetExecutionTemplateResponseBodyTemplate {
	s.TemplateId = &v
	return s
}

func (s *GetExecutionTemplateResponseBodyTemplate) SetTemplateName(v string) *GetExecutionTemplateResponseBodyTemplate {
	s.TemplateName = &v
	return s
}

func (s *GetExecutionTemplateResponseBodyTemplate) SetTemplateVersion(v string) *GetExecutionTemplateResponseBodyTemplate {
	s.TemplateVersion = &v
	return s
}

func (s *GetExecutionTemplateResponseBodyTemplate) SetUpdatedBy(v string) *GetExecutionTemplateResponseBodyTemplate {
	s.UpdatedBy = &v
	return s
}

func (s *GetExecutionTemplateResponseBodyTemplate) SetUpdatedDate(v string) *GetExecutionTemplateResponseBodyTemplate {
	s.UpdatedDate = &v
	return s
}

func (s *GetExecutionTemplateResponseBodyTemplate) Validate() error {
	return dara.Validate(s)
}
