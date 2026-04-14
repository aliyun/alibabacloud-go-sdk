// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *DescTemplateResponseBody
	GetCreateTime() *string
	SetRemark(v string) *DescTemplateResponseBody
	GetRemark() *string
	SetRequestId(v string) *DescTemplateResponseBody
	GetRequestId() *string
	SetSmsContent(v string) *DescTemplateResponseBody
	GetSmsContent() *string
	SetSmsType(v string) *DescTemplateResponseBody
	GetSmsType() *string
	SetTemplateName(v string) *DescTemplateResponseBody
	GetTemplateName() *string
	SetTemplateNickName(v string) *DescTemplateResponseBody
	GetTemplateNickName() *string
	SetTemplateStatus(v string) *DescTemplateResponseBody
	GetTemplateStatus() *string
	SetTemplateSubject(v string) *DescTemplateResponseBody
	GetTemplateSubject() *string
	SetTemplateText(v string) *DescTemplateResponseBody
	GetTemplateText() *string
	SetTemplateType(v string) *DescTemplateResponseBody
	GetTemplateType() *string
}

type DescTemplateResponseBody struct {
	// example:
	//
	// 2019-09-29T13:28Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Remark     *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// 95xxx5F
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SmsContent *string `json:"SmsContent,omitempty" xml:"SmsContent,omitempty"`
	SmsType    *string `json:"SmsType,omitempty" xml:"SmsType,omitempty"`
	// example:
	//
	// test1
	TemplateName     *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateNickName *string `json:"TemplateNickName,omitempty" xml:"TemplateNickName,omitempty"`
	// example:
	//
	// 2
	TemplateStatus *string `json:"TemplateStatus,omitempty" xml:"TemplateStatus,omitempty"`
	// example:
	//
	// test
	TemplateSubject *string `json:"TemplateSubject,omitempty" xml:"TemplateSubject,omitempty"`
	// example:
	//
	// <p>1</p>
	TemplateText *string `json:"TemplateText,omitempty" xml:"TemplateText,omitempty"`
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s DescTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DescTemplateResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescTemplateResponseBody) GetRemark() *string {
	return s.Remark
}

func (s *DescTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescTemplateResponseBody) GetSmsContent() *string {
	return s.SmsContent
}

func (s *DescTemplateResponseBody) GetSmsType() *string {
	return s.SmsType
}

func (s *DescTemplateResponseBody) GetTemplateName() *string {
	return s.TemplateName
}

func (s *DescTemplateResponseBody) GetTemplateNickName() *string {
	return s.TemplateNickName
}

func (s *DescTemplateResponseBody) GetTemplateStatus() *string {
	return s.TemplateStatus
}

func (s *DescTemplateResponseBody) GetTemplateSubject() *string {
	return s.TemplateSubject
}

func (s *DescTemplateResponseBody) GetTemplateText() *string {
	return s.TemplateText
}

func (s *DescTemplateResponseBody) GetTemplateType() *string {
	return s.TemplateType
}

func (s *DescTemplateResponseBody) SetCreateTime(v string) *DescTemplateResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescTemplateResponseBody) SetRemark(v string) *DescTemplateResponseBody {
	s.Remark = &v
	return s
}

func (s *DescTemplateResponseBody) SetRequestId(v string) *DescTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescTemplateResponseBody) SetSmsContent(v string) *DescTemplateResponseBody {
	s.SmsContent = &v
	return s
}

func (s *DescTemplateResponseBody) SetSmsType(v string) *DescTemplateResponseBody {
	s.SmsType = &v
	return s
}

func (s *DescTemplateResponseBody) SetTemplateName(v string) *DescTemplateResponseBody {
	s.TemplateName = &v
	return s
}

func (s *DescTemplateResponseBody) SetTemplateNickName(v string) *DescTemplateResponseBody {
	s.TemplateNickName = &v
	return s
}

func (s *DescTemplateResponseBody) SetTemplateStatus(v string) *DescTemplateResponseBody {
	s.TemplateStatus = &v
	return s
}

func (s *DescTemplateResponseBody) SetTemplateSubject(v string) *DescTemplateResponseBody {
	s.TemplateSubject = &v
	return s
}

func (s *DescTemplateResponseBody) SetTemplateText(v string) *DescTemplateResponseBody {
	s.TemplateText = &v
	return s
}

func (s *DescTemplateResponseBody) SetTemplateType(v string) *DescTemplateResponseBody {
	s.TemplateType = &v
	return s
}

func (s *DescTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
