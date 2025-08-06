// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDefaultAITemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetDefaultAITemplateResponseBody
	GetRequestId() *string
	SetTemplateInfo(v *GetDefaultAITemplateResponseBodyTemplateInfo) *GetDefaultAITemplateResponseBody
	GetTemplateInfo() *GetDefaultAITemplateResponseBodyTemplateInfo
}

type GetDefaultAITemplateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// A017F1DE-3DC3-4441-6755-37E81113****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the AI template.
	TemplateInfo *GetDefaultAITemplateResponseBodyTemplateInfo `json:"TemplateInfo,omitempty" xml:"TemplateInfo,omitempty" type:"Struct"`
}

func (s GetDefaultAITemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDefaultAITemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetDefaultAITemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDefaultAITemplateResponseBody) GetTemplateInfo() *GetDefaultAITemplateResponseBodyTemplateInfo {
	return s.TemplateInfo
}

func (s *GetDefaultAITemplateResponseBody) SetRequestId(v string) *GetDefaultAITemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDefaultAITemplateResponseBody) SetTemplateInfo(v *GetDefaultAITemplateResponseBodyTemplateInfo) *GetDefaultAITemplateResponseBody {
	s.TemplateInfo = v
	return s
}

func (s *GetDefaultAITemplateResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDefaultAITemplateResponseBodyTemplateInfo struct {
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
	// Default
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
	// The type of the AI template. The value is **AIMediaAudit**, which indicates automated review.
	//
	// example:
	//
	// AIMediaAudit
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s GetDefaultAITemplateResponseBodyTemplateInfo) String() string {
	return dara.Prettify(s)
}

func (s GetDefaultAITemplateResponseBodyTemplateInfo) GoString() string {
	return s.String()
}

func (s *GetDefaultAITemplateResponseBodyTemplateInfo) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetDefaultAITemplateResponseBodyTemplateInfo) GetIsDefault() *string {
	return s.IsDefault
}

func (s *GetDefaultAITemplateResponseBodyTemplateInfo) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetDefaultAITemplateResponseBodyTemplateInfo) GetSource() *string {
	return s.Source
}

func (s *GetDefaultAITemplateResponseBodyTemplateInfo) GetTemplateConfig() *string {
	return s.TemplateConfig
}

func (s *GetDefaultAITemplateResponseBodyTemplateInfo) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetDefaultAITemplateResponseBodyTemplateInfo) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetDefaultAITemplateResponseBodyTemplateInfo) GetTemplateType() *string {
	return s.TemplateType
}

func (s *GetDefaultAITemplateResponseBodyTemplateInfo) SetCreationTime(v string) *GetDefaultAITemplateResponseBodyTemplateInfo {
	s.CreationTime = &v
	return s
}

func (s *GetDefaultAITemplateResponseBodyTemplateInfo) SetIsDefault(v string) *GetDefaultAITemplateResponseBodyTemplateInfo {
	s.IsDefault = &v
	return s
}

func (s *GetDefaultAITemplateResponseBodyTemplateInfo) SetModifyTime(v string) *GetDefaultAITemplateResponseBodyTemplateInfo {
	s.ModifyTime = &v
	return s
}

func (s *GetDefaultAITemplateResponseBodyTemplateInfo) SetSource(v string) *GetDefaultAITemplateResponseBodyTemplateInfo {
	s.Source = &v
	return s
}

func (s *GetDefaultAITemplateResponseBodyTemplateInfo) SetTemplateConfig(v string) *GetDefaultAITemplateResponseBodyTemplateInfo {
	s.TemplateConfig = &v
	return s
}

func (s *GetDefaultAITemplateResponseBodyTemplateInfo) SetTemplateId(v string) *GetDefaultAITemplateResponseBodyTemplateInfo {
	s.TemplateId = &v
	return s
}

func (s *GetDefaultAITemplateResponseBodyTemplateInfo) SetTemplateName(v string) *GetDefaultAITemplateResponseBodyTemplateInfo {
	s.TemplateName = &v
	return s
}

func (s *GetDefaultAITemplateResponseBodyTemplateInfo) SetTemplateType(v string) *GetDefaultAITemplateResponseBodyTemplateInfo {
	s.TemplateType = &v
	return s
}

func (s *GetDefaultAITemplateResponseBodyTemplateInfo) Validate() error {
	return dara.Validate(s)
}
