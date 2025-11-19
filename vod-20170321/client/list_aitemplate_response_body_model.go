// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAITemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListAITemplateResponseBody
	GetRequestId() *string
	SetTemplateInfoList(v []*ListAITemplateResponseBodyTemplateInfoList) *ListAITemplateResponseBody
	GetTemplateInfoList() []*ListAITemplateResponseBodyTemplateInfoList
}

type ListAITemplateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 75B7BC67-FB8C-4653-4788-F4B01ED2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the AI templates.
	TemplateInfoList []*ListAITemplateResponseBodyTemplateInfoList `json:"TemplateInfoList,omitempty" xml:"TemplateInfoList,omitempty" type:"Repeated"`
}

func (s ListAITemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAITemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ListAITemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAITemplateResponseBody) GetTemplateInfoList() []*ListAITemplateResponseBodyTemplateInfoList {
	return s.TemplateInfoList
}

func (s *ListAITemplateResponseBody) SetRequestId(v string) *ListAITemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAITemplateResponseBody) SetTemplateInfoList(v []*ListAITemplateResponseBodyTemplateInfoList) *ListAITemplateResponseBody {
	s.TemplateInfoList = v
	return s
}

func (s *ListAITemplateResponseBody) Validate() error {
	if s.TemplateInfoList != nil {
		for _, item := range s.TemplateInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAITemplateResponseBodyTemplateInfoList struct {
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
	// NoDefault
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

func (s ListAITemplateResponseBodyTemplateInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListAITemplateResponseBodyTemplateInfoList) GoString() string {
	return s.String()
}

func (s *ListAITemplateResponseBodyTemplateInfoList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListAITemplateResponseBodyTemplateInfoList) GetIsDefault() *string {
	return s.IsDefault
}

func (s *ListAITemplateResponseBodyTemplateInfoList) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListAITemplateResponseBodyTemplateInfoList) GetSource() *string {
	return s.Source
}

func (s *ListAITemplateResponseBodyTemplateInfoList) GetTemplateConfig() *string {
	return s.TemplateConfig
}

func (s *ListAITemplateResponseBodyTemplateInfoList) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListAITemplateResponseBodyTemplateInfoList) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListAITemplateResponseBodyTemplateInfoList) GetTemplateType() *string {
	return s.TemplateType
}

func (s *ListAITemplateResponseBodyTemplateInfoList) SetCreationTime(v string) *ListAITemplateResponseBodyTemplateInfoList {
	s.CreationTime = &v
	return s
}

func (s *ListAITemplateResponseBodyTemplateInfoList) SetIsDefault(v string) *ListAITemplateResponseBodyTemplateInfoList {
	s.IsDefault = &v
	return s
}

func (s *ListAITemplateResponseBodyTemplateInfoList) SetModifyTime(v string) *ListAITemplateResponseBodyTemplateInfoList {
	s.ModifyTime = &v
	return s
}

func (s *ListAITemplateResponseBodyTemplateInfoList) SetSource(v string) *ListAITemplateResponseBodyTemplateInfoList {
	s.Source = &v
	return s
}

func (s *ListAITemplateResponseBodyTemplateInfoList) SetTemplateConfig(v string) *ListAITemplateResponseBodyTemplateInfoList {
	s.TemplateConfig = &v
	return s
}

func (s *ListAITemplateResponseBodyTemplateInfoList) SetTemplateId(v string) *ListAITemplateResponseBodyTemplateInfoList {
	s.TemplateId = &v
	return s
}

func (s *ListAITemplateResponseBodyTemplateInfoList) SetTemplateName(v string) *ListAITemplateResponseBodyTemplateInfoList {
	s.TemplateName = &v
	return s
}

func (s *ListAITemplateResponseBodyTemplateInfoList) SetTemplateType(v string) *ListAITemplateResponseBodyTemplateInfoList {
	s.TemplateType = &v
	return s
}

func (s *ListAITemplateResponseBodyTemplateInfoList) Validate() error {
	return dara.Validate(s)
}
