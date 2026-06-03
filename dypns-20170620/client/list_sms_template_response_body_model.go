// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSmsTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListSmsTemplateResponseBody
	GetCode() *string
	SetData(v []*ListSmsTemplateResponseBodyData) *ListSmsTemplateResponseBody
	GetData() []*ListSmsTemplateResponseBodyData
	SetMessage(v string) *ListSmsTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListSmsTemplateResponseBody
	GetRequestId() *string
}

type ListSmsTemplateResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*ListSmsTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSmsTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSmsTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ListSmsTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListSmsTemplateResponseBody) GetData() []*ListSmsTemplateResponseBodyData {
	return s.Data
}

func (s *ListSmsTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSmsTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSmsTemplateResponseBody) SetCode(v string) *ListSmsTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *ListSmsTemplateResponseBody) SetData(v []*ListSmsTemplateResponseBodyData) *ListSmsTemplateResponseBody {
	s.Data = v
	return s
}

func (s *ListSmsTemplateResponseBody) SetMessage(v string) *ListSmsTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *ListSmsTemplateResponseBody) SetRequestId(v string) *ListSmsTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSmsTemplateResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSmsTemplateResponseBodyData struct {
	AuditResult        *string `json:"AuditResult,omitempty" xml:"AuditResult,omitempty"`
	BizUrl             *string `json:"BizUrl,omitempty" xml:"BizUrl,omitempty"`
	BusinessType       *int32  `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	CreateDate         *int64  `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	DefaultFlag        *bool   `json:"DefaultFlag,omitempty" xml:"DefaultFlag,omitempty"`
	Remark             *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	SmsTemplateCode    *string `json:"SmsTemplateCode,omitempty" xml:"SmsTemplateCode,omitempty"`
	SmsTemplateContent *string `json:"SmsTemplateContent,omitempty" xml:"SmsTemplateContent,omitempty"`
	SmsTemplateName    *string `json:"SmsTemplateName,omitempty" xml:"SmsTemplateName,omitempty"`
	Status             *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListSmsTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSmsTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSmsTemplateResponseBodyData) GetAuditResult() *string {
	return s.AuditResult
}

func (s *ListSmsTemplateResponseBodyData) GetBizUrl() *string {
	return s.BizUrl
}

func (s *ListSmsTemplateResponseBodyData) GetBusinessType() *int32 {
	return s.BusinessType
}

func (s *ListSmsTemplateResponseBodyData) GetCreateDate() *int64 {
	return s.CreateDate
}

func (s *ListSmsTemplateResponseBodyData) GetDefaultFlag() *bool {
	return s.DefaultFlag
}

func (s *ListSmsTemplateResponseBodyData) GetRemark() *string {
	return s.Remark
}

func (s *ListSmsTemplateResponseBodyData) GetSmsTemplateCode() *string {
	return s.SmsTemplateCode
}

func (s *ListSmsTemplateResponseBodyData) GetSmsTemplateContent() *string {
	return s.SmsTemplateContent
}

func (s *ListSmsTemplateResponseBodyData) GetSmsTemplateName() *string {
	return s.SmsTemplateName
}

func (s *ListSmsTemplateResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListSmsTemplateResponseBodyData) SetAuditResult(v string) *ListSmsTemplateResponseBodyData {
	s.AuditResult = &v
	return s
}

func (s *ListSmsTemplateResponseBodyData) SetBizUrl(v string) *ListSmsTemplateResponseBodyData {
	s.BizUrl = &v
	return s
}

func (s *ListSmsTemplateResponseBodyData) SetBusinessType(v int32) *ListSmsTemplateResponseBodyData {
	s.BusinessType = &v
	return s
}

func (s *ListSmsTemplateResponseBodyData) SetCreateDate(v int64) *ListSmsTemplateResponseBodyData {
	s.CreateDate = &v
	return s
}

func (s *ListSmsTemplateResponseBodyData) SetDefaultFlag(v bool) *ListSmsTemplateResponseBodyData {
	s.DefaultFlag = &v
	return s
}

func (s *ListSmsTemplateResponseBodyData) SetRemark(v string) *ListSmsTemplateResponseBodyData {
	s.Remark = &v
	return s
}

func (s *ListSmsTemplateResponseBodyData) SetSmsTemplateCode(v string) *ListSmsTemplateResponseBodyData {
	s.SmsTemplateCode = &v
	return s
}

func (s *ListSmsTemplateResponseBodyData) SetSmsTemplateContent(v string) *ListSmsTemplateResponseBodyData {
	s.SmsTemplateContent = &v
	return s
}

func (s *ListSmsTemplateResponseBodyData) SetSmsTemplateName(v string) *ListSmsTemplateResponseBodyData {
	s.SmsTemplateName = &v
	return s
}

func (s *ListSmsTemplateResponseBodyData) SetStatus(v string) *ListSmsTemplateResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListSmsTemplateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
