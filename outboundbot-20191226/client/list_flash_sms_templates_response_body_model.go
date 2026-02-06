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
	SetData(v []*ListFlashSmsTemplatesResponseBodyData) *ListFlashSmsTemplatesResponseBody
	GetData() []*ListFlashSmsTemplatesResponseBodyData
	SetHttpStatusCode(v int32) *ListFlashSmsTemplatesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListFlashSmsTemplatesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListFlashSmsTemplatesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListFlashSmsTemplatesResponseBody
	GetSuccess() *bool
}

type ListFlashSmsTemplatesResponseBody struct {
	// example:
	//
	// OK
	Code *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListFlashSmsTemplatesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EC08CC41-6870-5594-939A-F758F057898F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *ListFlashSmsTemplatesResponseBody) GetData() []*ListFlashSmsTemplatesResponseBodyData {
	return s.Data
}

func (s *ListFlashSmsTemplatesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListFlashSmsTemplatesResponseBody) GetMessage() *string {
	return s.Message
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

func (s *ListFlashSmsTemplatesResponseBody) SetData(v []*ListFlashSmsTemplatesResponseBodyData) *ListFlashSmsTemplatesResponseBody {
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

type ListFlashSmsTemplatesResponseBodyData struct {
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	// example:
	//
	// 17*******************01
	TemplateId   *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ListFlashSmsTemplatesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListFlashSmsTemplatesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFlashSmsTemplatesResponseBodyData) GetTemplateContent() *string {
	return s.TemplateContent
}

func (s *ListFlashSmsTemplatesResponseBodyData) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListFlashSmsTemplatesResponseBodyData) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListFlashSmsTemplatesResponseBodyData) SetTemplateContent(v string) *ListFlashSmsTemplatesResponseBodyData {
	s.TemplateContent = &v
	return s
}

func (s *ListFlashSmsTemplatesResponseBodyData) SetTemplateId(v string) *ListFlashSmsTemplatesResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *ListFlashSmsTemplatesResponseBodyData) SetTemplateName(v string) *ListFlashSmsTemplatesResponseBodyData {
	s.TemplateName = &v
	return s
}

func (s *ListFlashSmsTemplatesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
