// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAITemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetAITemplateResponseBody
	GetRequestId() *string
	SetTemplateInfo(v *GetAITemplateResponseBodyTemplateInfo) *GetAITemplateResponseBody
	GetTemplateInfo() *GetAITemplateResponseBodyTemplateInfo
}

type GetAITemplateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 080DA371-8AC0-4CD4-4476-33E64282****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the AI template.
	TemplateInfo *GetAITemplateResponseBodyTemplateInfo `json:"TemplateInfo,omitempty" xml:"TemplateInfo,omitempty" type:"Struct"`
}

func (s GetAITemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAITemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetAITemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAITemplateResponseBody) GetTemplateInfo() *GetAITemplateResponseBodyTemplateInfo {
	return s.TemplateInfo
}

func (s *GetAITemplateResponseBody) SetRequestId(v string) *GetAITemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAITemplateResponseBody) SetTemplateInfo(v *GetAITemplateResponseBodyTemplateInfo) *GetAITemplateResponseBody {
	s.TemplateInfo = v
	return s
}

func (s *GetAITemplateResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAITemplateResponseBodyTemplateInfo struct {
	// The time when the AI template was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-07-08T06:50:45Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// Indicates whether the template is the default AI template. Valid values:
	//
	// 	- **Default**
	//
	// 	- **NotDefault**
	//
	// example:
	//
	// NotDefault
	IsDefault *string `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The time when the AI template was modified. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-07-08T06:58:45Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The source of the AI template. Valid values:
	//
	// 	- **System**
	//
	// 	- **Custom**
	//
	// example:
	//
	// Custom
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The detailed configurations of the AI template. The value is a JSON string. For more information, see [AITemplateConfig](~~89863#title-vd3-499-o36~~).
	//
	// example:
	//
	// {"AuditRange":["text-title","video"],"AuditContent":["screen"],"AuditItem":["terrorism","porn"],"AuditAutoBlock":"yes"}
	TemplateConfig *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	// The ID of the AI template.
	//
	// example:
	//
	// 1706a0063dd733f6a823ef32e0a5****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the AI template.
	//
	// example:
	//
	// DemoAITemplate
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The type of the AI template. Valid values:
	//
	// 	- **AIMediaAudit**: automated review
	//
	// 	- **AIImage**: smart thumbnail
	//
	// example:
	//
	// AIMediaAudit
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s GetAITemplateResponseBodyTemplateInfo) String() string {
	return dara.Prettify(s)
}

func (s GetAITemplateResponseBodyTemplateInfo) GoString() string {
	return s.String()
}

func (s *GetAITemplateResponseBodyTemplateInfo) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetAITemplateResponseBodyTemplateInfo) GetIsDefault() *string {
	return s.IsDefault
}

func (s *GetAITemplateResponseBodyTemplateInfo) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetAITemplateResponseBodyTemplateInfo) GetSource() *string {
	return s.Source
}

func (s *GetAITemplateResponseBodyTemplateInfo) GetTemplateConfig() *string {
	return s.TemplateConfig
}

func (s *GetAITemplateResponseBodyTemplateInfo) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetAITemplateResponseBodyTemplateInfo) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetAITemplateResponseBodyTemplateInfo) GetTemplateType() *string {
	return s.TemplateType
}

func (s *GetAITemplateResponseBodyTemplateInfo) SetCreationTime(v string) *GetAITemplateResponseBodyTemplateInfo {
	s.CreationTime = &v
	return s
}

func (s *GetAITemplateResponseBodyTemplateInfo) SetIsDefault(v string) *GetAITemplateResponseBodyTemplateInfo {
	s.IsDefault = &v
	return s
}

func (s *GetAITemplateResponseBodyTemplateInfo) SetModifyTime(v string) *GetAITemplateResponseBodyTemplateInfo {
	s.ModifyTime = &v
	return s
}

func (s *GetAITemplateResponseBodyTemplateInfo) SetSource(v string) *GetAITemplateResponseBodyTemplateInfo {
	s.Source = &v
	return s
}

func (s *GetAITemplateResponseBodyTemplateInfo) SetTemplateConfig(v string) *GetAITemplateResponseBodyTemplateInfo {
	s.TemplateConfig = &v
	return s
}

func (s *GetAITemplateResponseBodyTemplateInfo) SetTemplateId(v string) *GetAITemplateResponseBodyTemplateInfo {
	s.TemplateId = &v
	return s
}

func (s *GetAITemplateResponseBodyTemplateInfo) SetTemplateName(v string) *GetAITemplateResponseBodyTemplateInfo {
	s.TemplateName = &v
	return s
}

func (s *GetAITemplateResponseBodyTemplateInfo) SetTemplateType(v string) *GetAITemplateResponseBodyTemplateInfo {
	s.TemplateType = &v
	return s
}

func (s *GetAITemplateResponseBodyTemplateInfo) Validate() error {
	return dara.Validate(s)
}
