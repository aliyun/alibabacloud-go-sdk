// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlashSmsTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListFlashSmsTemplatesResponseBody
	GetCode() *string
	SetData(v *ListFlashSmsTemplatesResponseBodyData) *ListFlashSmsTemplatesResponseBody
	GetData() *ListFlashSmsTemplatesResponseBodyData
	SetHttpStatusCode(v int32) *ListFlashSmsTemplatesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListFlashSmsTemplatesResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListFlashSmsTemplatesResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListFlashSmsTemplatesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListFlashSmsTemplatesResponseBody
	GetSuccess() *bool
}

type ListFlashSmsTemplatesResponseBody struct {
	// 返回码
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 返回数据
	Data *ListFlashSmsTemplatesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP状态码
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// 错误信息
	//
	// example:
	//
	// Instance does not exist. Instance=out001
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 错误信息中的变量值列表
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// 请求ID
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 是否调用成功
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListFlashSmsTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFlashSmsTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlashSmsTemplatesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListFlashSmsTemplatesResponseBody) GetData() *ListFlashSmsTemplatesResponseBodyData {
	return s.Data
}

func (s *ListFlashSmsTemplatesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListFlashSmsTemplatesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListFlashSmsTemplatesResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListFlashSmsTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFlashSmsTemplatesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListFlashSmsTemplatesResponseBody) SetCode(v string) *ListFlashSmsTemplatesResponseBody {
	s.Code = &v
	return s
}

func (s *ListFlashSmsTemplatesResponseBody) SetData(v *ListFlashSmsTemplatesResponseBodyData) *ListFlashSmsTemplatesResponseBody {
	s.Data = v
	return s
}

func (s *ListFlashSmsTemplatesResponseBody) SetHttpStatusCode(v int32) *ListFlashSmsTemplatesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListFlashSmsTemplatesResponseBody) SetMessage(v string) *ListFlashSmsTemplatesResponseBody {
	s.Message = &v
	return s
}

func (s *ListFlashSmsTemplatesResponseBody) SetParams(v []*string) *ListFlashSmsTemplatesResponseBody {
	s.Params = v
	return s
}

func (s *ListFlashSmsTemplatesResponseBody) SetRequestId(v string) *ListFlashSmsTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlashSmsTemplatesResponseBody) SetSuccess(v bool) *ListFlashSmsTemplatesResponseBody {
	s.Success = &v
	return s
}

func (s *ListFlashSmsTemplatesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListFlashSmsTemplatesResponseBodyData struct {
	// 数据列表
	FlashSmsTemplates []*ListFlashSmsTemplatesResponseBodyDataFlashSmsTemplates `json:"FlashSmsTemplates,omitempty" xml:"FlashSmsTemplates,omitempty" type:"Repeated"`
	// 页码，从1开始
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页记录数
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 符合条件的记录总数
	//
	// example:
	//
	// 0
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFlashSmsTemplatesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListFlashSmsTemplatesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFlashSmsTemplatesResponseBodyData) GetFlashSmsTemplates() []*ListFlashSmsTemplatesResponseBodyDataFlashSmsTemplates {
	return s.FlashSmsTemplates
}

func (s *ListFlashSmsTemplatesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListFlashSmsTemplatesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFlashSmsTemplatesResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListFlashSmsTemplatesResponseBodyData) SetFlashSmsTemplates(v []*ListFlashSmsTemplatesResponseBodyDataFlashSmsTemplates) *ListFlashSmsTemplatesResponseBodyData {
	s.FlashSmsTemplates = v
	return s
}

func (s *ListFlashSmsTemplatesResponseBodyData) SetPageNumber(v int32) *ListFlashSmsTemplatesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListFlashSmsTemplatesResponseBodyData) SetPageSize(v int32) *ListFlashSmsTemplatesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListFlashSmsTemplatesResponseBodyData) SetTotalCount(v int32) *ListFlashSmsTemplatesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListFlashSmsTemplatesResponseBodyData) Validate() error {
	if s.FlashSmsTemplates != nil {
		for _, item := range s.FlashSmsTemplates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListFlashSmsTemplatesResponseBodyDataFlashSmsTemplates struct {
	// 模板ID
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 模板名称
	//
	// example:
	//
	// 测试模版
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ListFlashSmsTemplatesResponseBodyDataFlashSmsTemplates) String() string {
	return dara.Prettify(s)
}

func (s ListFlashSmsTemplatesResponseBodyDataFlashSmsTemplates) GoString() string {
	return s.String()
}

func (s *ListFlashSmsTemplatesResponseBodyDataFlashSmsTemplates) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListFlashSmsTemplatesResponseBodyDataFlashSmsTemplates) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListFlashSmsTemplatesResponseBodyDataFlashSmsTemplates) SetTemplateId(v string) *ListFlashSmsTemplatesResponseBodyDataFlashSmsTemplates {
	s.TemplateId = &v
	return s
}

func (s *ListFlashSmsTemplatesResponseBodyDataFlashSmsTemplates) SetTemplateName(v string) *ListFlashSmsTemplatesResponseBodyDataFlashSmsTemplates {
	s.TemplateName = &v
	return s
}

func (s *ListFlashSmsTemplatesResponseBodyDataFlashSmsTemplates) Validate() error {
	return dara.Validate(s)
}
