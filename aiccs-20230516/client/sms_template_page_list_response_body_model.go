// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSmsTemplatePageListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *SmsTemplatePageListResponseBody
	GetCode() *int64
	SetMessage(v string) *SmsTemplatePageListResponseBody
	GetMessage() *string
	SetModel(v *SmsTemplatePageListResponseBodyModel) *SmsTemplatePageListResponseBody
	GetModel() *SmsTemplatePageListResponseBodyModel
	SetRequestId(v string) *SmsTemplatePageListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SmsTemplatePageListResponseBody
	GetSuccess() *bool
	SetTimestamp(v int64) *SmsTemplatePageListResponseBody
	GetTimestamp() *int64
}

type SmsTemplatePageListResponseBody struct {
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值
	Message *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *SmsTemplatePageListResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// 8EFC6D10-307B-1ECA-A8C6-7CBDF776AAD2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1683440860035
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s SmsTemplatePageListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SmsTemplatePageListResponseBody) GoString() string {
	return s.String()
}

func (s *SmsTemplatePageListResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *SmsTemplatePageListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SmsTemplatePageListResponseBody) GetModel() *SmsTemplatePageListResponseBodyModel {
	return s.Model
}

func (s *SmsTemplatePageListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SmsTemplatePageListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SmsTemplatePageListResponseBody) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *SmsTemplatePageListResponseBody) SetCode(v int64) *SmsTemplatePageListResponseBody {
	s.Code = &v
	return s
}

func (s *SmsTemplatePageListResponseBody) SetMessage(v string) *SmsTemplatePageListResponseBody {
	s.Message = &v
	return s
}

func (s *SmsTemplatePageListResponseBody) SetModel(v *SmsTemplatePageListResponseBodyModel) *SmsTemplatePageListResponseBody {
	s.Model = v
	return s
}

func (s *SmsTemplatePageListResponseBody) SetRequestId(v string) *SmsTemplatePageListResponseBody {
	s.RequestId = &v
	return s
}

func (s *SmsTemplatePageListResponseBody) SetSuccess(v bool) *SmsTemplatePageListResponseBody {
	s.Success = &v
	return s
}

func (s *SmsTemplatePageListResponseBody) SetTimestamp(v int64) *SmsTemplatePageListResponseBody {
	s.Timestamp = &v
	return s
}

func (s *SmsTemplatePageListResponseBody) Validate() error {
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SmsTemplatePageListResponseBodyModel struct {
	List []*SmsTemplatePageListResponseBodyModelList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 53
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 42
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 31
	TotalPage *int64 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s SmsTemplatePageListResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s SmsTemplatePageListResponseBodyModel) GoString() string {
	return s.String()
}

func (s *SmsTemplatePageListResponseBodyModel) GetList() []*SmsTemplatePageListResponseBodyModelList {
	return s.List
}

func (s *SmsTemplatePageListResponseBodyModel) GetPageNo() *int64 {
	return s.PageNo
}

func (s *SmsTemplatePageListResponseBodyModel) GetPageSize() *int64 {
	return s.PageSize
}

func (s *SmsTemplatePageListResponseBodyModel) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *SmsTemplatePageListResponseBodyModel) GetTotalPage() *int64 {
	return s.TotalPage
}

func (s *SmsTemplatePageListResponseBodyModel) SetList(v []*SmsTemplatePageListResponseBodyModelList) *SmsTemplatePageListResponseBodyModel {
	s.List = v
	return s
}

func (s *SmsTemplatePageListResponseBodyModel) SetPageNo(v int64) *SmsTemplatePageListResponseBodyModel {
	s.PageNo = &v
	return s
}

func (s *SmsTemplatePageListResponseBodyModel) SetPageSize(v int64) *SmsTemplatePageListResponseBodyModel {
	s.PageSize = &v
	return s
}

func (s *SmsTemplatePageListResponseBodyModel) SetTotalCount(v int64) *SmsTemplatePageListResponseBodyModel {
	s.TotalCount = &v
	return s
}

func (s *SmsTemplatePageListResponseBodyModel) SetTotalPage(v int64) *SmsTemplatePageListResponseBodyModel {
	s.TotalPage = &v
	return s
}

func (s *SmsTemplatePageListResponseBodyModel) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SmsTemplatePageListResponseBodyModelList struct {
	// 短信内容
	//
	// example:
	//
	// 示例值示例值示例值
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 创建时间
	//
	// example:
	//
	// 2021-09-26 11:34:59
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 模板所需参数
	//
	// example:
	//
	// 示例值示例值
	Properties *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// 智能短链ID
	//
	// example:
	//
	// 46
	ShortUrlTaskId *int64 `json:"ShortUrlTaskId,omitempty" xml:"ShortUrlTaskId,omitempty"`
	// 短信签名
	//
	// example:
	//
	// a234n
	Sign *string `json:"Sign,omitempty" xml:"Sign,omitempty"`
	// 短信类型
	//
	// example:
	//
	// 示例值示例值
	SmsType *string `json:"SmsType,omitempty" xml:"SmsType,omitempty"`
	// 模板状态
	//
	// example:
	//
	// 18
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 模板ID
	//
	// example:
	//
	// 67
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 模板名称
	//
	// example:
	//
	// 示例值示例值示例值
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// 模板类型
	//
	// example:
	//
	// 56
	TemplateType *int64 `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s SmsTemplatePageListResponseBodyModelList) String() string {
	return dara.Prettify(s)
}

func (s SmsTemplatePageListResponseBodyModelList) GoString() string {
	return s.String()
}

func (s *SmsTemplatePageListResponseBodyModelList) GetContent() *string {
	return s.Content
}

func (s *SmsTemplatePageListResponseBodyModelList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *SmsTemplatePageListResponseBodyModelList) GetProperties() *string {
	return s.Properties
}

func (s *SmsTemplatePageListResponseBodyModelList) GetShortUrlTaskId() *int64 {
	return s.ShortUrlTaskId
}

func (s *SmsTemplatePageListResponseBodyModelList) GetSign() *string {
	return s.Sign
}

func (s *SmsTemplatePageListResponseBodyModelList) GetSmsType() *string {
	return s.SmsType
}

func (s *SmsTemplatePageListResponseBodyModelList) GetStatus() *int64 {
	return s.Status
}

func (s *SmsTemplatePageListResponseBodyModelList) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *SmsTemplatePageListResponseBodyModelList) GetTemplateName() *string {
	return s.TemplateName
}

func (s *SmsTemplatePageListResponseBodyModelList) GetTemplateType() *int64 {
	return s.TemplateType
}

func (s *SmsTemplatePageListResponseBodyModelList) SetContent(v string) *SmsTemplatePageListResponseBodyModelList {
	s.Content = &v
	return s
}

func (s *SmsTemplatePageListResponseBodyModelList) SetCreateTime(v string) *SmsTemplatePageListResponseBodyModelList {
	s.CreateTime = &v
	return s
}

func (s *SmsTemplatePageListResponseBodyModelList) SetProperties(v string) *SmsTemplatePageListResponseBodyModelList {
	s.Properties = &v
	return s
}

func (s *SmsTemplatePageListResponseBodyModelList) SetShortUrlTaskId(v int64) *SmsTemplatePageListResponseBodyModelList {
	s.ShortUrlTaskId = &v
	return s
}

func (s *SmsTemplatePageListResponseBodyModelList) SetSign(v string) *SmsTemplatePageListResponseBodyModelList {
	s.Sign = &v
	return s
}

func (s *SmsTemplatePageListResponseBodyModelList) SetSmsType(v string) *SmsTemplatePageListResponseBodyModelList {
	s.SmsType = &v
	return s
}

func (s *SmsTemplatePageListResponseBodyModelList) SetStatus(v int64) *SmsTemplatePageListResponseBodyModelList {
	s.Status = &v
	return s
}

func (s *SmsTemplatePageListResponseBodyModelList) SetTemplateId(v int64) *SmsTemplatePageListResponseBodyModelList {
	s.TemplateId = &v
	return s
}

func (s *SmsTemplatePageListResponseBodyModelList) SetTemplateName(v string) *SmsTemplatePageListResponseBodyModelList {
	s.TemplateName = &v
	return s
}

func (s *SmsTemplatePageListResponseBodyModelList) SetTemplateType(v int64) *SmsTemplatePageListResponseBodyModelList {
	s.TemplateType = &v
	return s
}

func (s *SmsTemplatePageListResponseBodyModelList) Validate() error {
	return dara.Validate(s)
}
