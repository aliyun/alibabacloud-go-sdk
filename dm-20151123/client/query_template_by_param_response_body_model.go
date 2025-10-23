// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTemplateByParamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *QueryTemplateByParamResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *QueryTemplateByParamResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *QueryTemplateByParamResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *QueryTemplateByParamResponseBody
	GetTotalCount() *int32
	SetData(v *QueryTemplateByParamResponseBodyData) *QueryTemplateByParamResponseBody
	GetData() *QueryTemplateByParamResponseBodyData
}

type QueryTemplateByParamResponseBody struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 10Axxxxxxxxxxxx37
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 21
	TotalCount *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Data       *QueryTemplateByParamResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s QueryTemplateByParamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryTemplateByParamResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTemplateByParamResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryTemplateByParamResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryTemplateByParamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryTemplateByParamResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QueryTemplateByParamResponseBody) GetData() *QueryTemplateByParamResponseBodyData {
	return s.Data
}

func (s *QueryTemplateByParamResponseBody) SetPageNumber(v int32) *QueryTemplateByParamResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryTemplateByParamResponseBody) SetPageSize(v int32) *QueryTemplateByParamResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryTemplateByParamResponseBody) SetRequestId(v string) *QueryTemplateByParamResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTemplateByParamResponseBody) SetTotalCount(v int32) *QueryTemplateByParamResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryTemplateByParamResponseBody) SetData(v *QueryTemplateByParamResponseBodyData) *QueryTemplateByParamResponseBody {
	s.Data = v
	return s
}

func (s *QueryTemplateByParamResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryTemplateByParamResponseBodyData struct {
	Template []*QueryTemplateByParamResponseBodyDataTemplate `json:"template,omitempty" xml:"template,omitempty" type:"Repeated"`
}

func (s QueryTemplateByParamResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryTemplateByParamResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTemplateByParamResponseBodyData) GetTemplate() []*QueryTemplateByParamResponseBodyDataTemplate {
	return s.Template
}

func (s *QueryTemplateByParamResponseBodyData) SetTemplate(v []*QueryTemplateByParamResponseBodyDataTemplate) *QueryTemplateByParamResponseBodyData {
	s.Template = v
	return s
}

func (s *QueryTemplateByParamResponseBodyData) Validate() error {
	if s.Template != nil {
		for _, item := range s.Template {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryTemplateByParamResponseBodyDataTemplate struct {
	// example:
	//
	// 2019-09-29T13:28Z
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	SmsStatus       *int32  `json:"SmsStatus,omitempty" xml:"SmsStatus,omitempty"`
	SmsTemplateCode *int32  `json:"SmsTemplateCode,omitempty" xml:"SmsTemplateCode,omitempty"`
	Smsrejectinfo   *int32  `json:"Smsrejectinfo,omitempty" xml:"Smsrejectinfo,omitempty"`
	TemplateComment *string `json:"TemplateComment,omitempty" xml:"TemplateComment,omitempty"`
	// example:
	//
	// 3xxxx8
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// test
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// 2
	TemplateStatus *string `json:"TemplateStatus,omitempty" xml:"TemplateStatus,omitempty"`
	TemplateType   *int32  `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// example:
	//
	// 1661830382
	UtcCreatetime *int64 `json:"UtcCreatetime,omitempty" xml:"UtcCreatetime,omitempty"`
}

func (s QueryTemplateByParamResponseBodyDataTemplate) String() string {
	return dara.Prettify(s)
}

func (s QueryTemplateByParamResponseBodyDataTemplate) GoString() string {
	return s.String()
}

func (s *QueryTemplateByParamResponseBodyDataTemplate) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryTemplateByParamResponseBodyDataTemplate) GetSmsStatus() *int32 {
	return s.SmsStatus
}

func (s *QueryTemplateByParamResponseBodyDataTemplate) GetSmsTemplateCode() *int32 {
	return s.SmsTemplateCode
}

func (s *QueryTemplateByParamResponseBodyDataTemplate) GetSmsrejectinfo() *int32 {
	return s.Smsrejectinfo
}

func (s *QueryTemplateByParamResponseBodyDataTemplate) GetTemplateComment() *string {
	return s.TemplateComment
}

func (s *QueryTemplateByParamResponseBodyDataTemplate) GetTemplateId() *string {
	return s.TemplateId
}

func (s *QueryTemplateByParamResponseBodyDataTemplate) GetTemplateName() *string {
	return s.TemplateName
}

func (s *QueryTemplateByParamResponseBodyDataTemplate) GetTemplateStatus() *string {
	return s.TemplateStatus
}

func (s *QueryTemplateByParamResponseBodyDataTemplate) GetTemplateType() *int32 {
	return s.TemplateType
}

func (s *QueryTemplateByParamResponseBodyDataTemplate) GetUtcCreatetime() *int64 {
	return s.UtcCreatetime
}

func (s *QueryTemplateByParamResponseBodyDataTemplate) SetCreateTime(v string) *QueryTemplateByParamResponseBodyDataTemplate {
	s.CreateTime = &v
	return s
}

func (s *QueryTemplateByParamResponseBodyDataTemplate) SetSmsStatus(v int32) *QueryTemplateByParamResponseBodyDataTemplate {
	s.SmsStatus = &v
	return s
}

func (s *QueryTemplateByParamResponseBodyDataTemplate) SetSmsTemplateCode(v int32) *QueryTemplateByParamResponseBodyDataTemplate {
	s.SmsTemplateCode = &v
	return s
}

func (s *QueryTemplateByParamResponseBodyDataTemplate) SetSmsrejectinfo(v int32) *QueryTemplateByParamResponseBodyDataTemplate {
	s.Smsrejectinfo = &v
	return s
}

func (s *QueryTemplateByParamResponseBodyDataTemplate) SetTemplateComment(v string) *QueryTemplateByParamResponseBodyDataTemplate {
	s.TemplateComment = &v
	return s
}

func (s *QueryTemplateByParamResponseBodyDataTemplate) SetTemplateId(v string) *QueryTemplateByParamResponseBodyDataTemplate {
	s.TemplateId = &v
	return s
}

func (s *QueryTemplateByParamResponseBodyDataTemplate) SetTemplateName(v string) *QueryTemplateByParamResponseBodyDataTemplate {
	s.TemplateName = &v
	return s
}

func (s *QueryTemplateByParamResponseBodyDataTemplate) SetTemplateStatus(v string) *QueryTemplateByParamResponseBodyDataTemplate {
	s.TemplateStatus = &v
	return s
}

func (s *QueryTemplateByParamResponseBodyDataTemplate) SetTemplateType(v int32) *QueryTemplateByParamResponseBodyDataTemplate {
	s.TemplateType = &v
	return s
}

func (s *QueryTemplateByParamResponseBodyDataTemplate) SetUtcCreatetime(v int64) *QueryTemplateByParamResponseBodyDataTemplate {
	s.UtcCreatetime = &v
	return s
}

func (s *QueryTemplateByParamResponseBodyDataTemplate) Validate() error {
	return dara.Validate(s)
}
