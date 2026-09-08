// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadWebhookContactSendTemplateListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReadWebhookContactSendTemplateListResponseBody
	GetCode() *string
	SetData(v []*ReadWebhookContactSendTemplateListResponseBodyData) *ReadWebhookContactSendTemplateListResponseBody
	GetData() []*ReadWebhookContactSendTemplateListResponseBodyData
	SetHttpCode(v int32) *ReadWebhookContactSendTemplateListResponseBody
	GetHttpCode() *int32
	SetMessage(v string) *ReadWebhookContactSendTemplateListResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReadWebhookContactSendTemplateListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReadWebhookContactSendTemplateListResponseBody
	GetSuccess() *bool
}

type ReadWebhookContactSendTemplateListResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result.
	Data []*ReadWebhookContactSendTemplateListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// 2xx
	HttpCode *int32 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The business message.
	//
	// example:
	//
	// Succeeded
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A5F62766-1C2F-1F56-A39D-63E3D30F0633
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReadWebhookContactSendTemplateListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReadWebhookContactSendTemplateListResponseBody) GoString() string {
	return s.String()
}

func (s *ReadWebhookContactSendTemplateListResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReadWebhookContactSendTemplateListResponseBody) GetData() []*ReadWebhookContactSendTemplateListResponseBodyData {
	return s.Data
}

func (s *ReadWebhookContactSendTemplateListResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *ReadWebhookContactSendTemplateListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReadWebhookContactSendTemplateListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReadWebhookContactSendTemplateListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReadWebhookContactSendTemplateListResponseBody) SetCode(v string) *ReadWebhookContactSendTemplateListResponseBody {
	s.Code = &v
	return s
}

func (s *ReadWebhookContactSendTemplateListResponseBody) SetData(v []*ReadWebhookContactSendTemplateListResponseBodyData) *ReadWebhookContactSendTemplateListResponseBody {
	s.Data = v
	return s
}

func (s *ReadWebhookContactSendTemplateListResponseBody) SetHttpCode(v int32) *ReadWebhookContactSendTemplateListResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ReadWebhookContactSendTemplateListResponseBody) SetMessage(v string) *ReadWebhookContactSendTemplateListResponseBody {
	s.Message = &v
	return s
}

func (s *ReadWebhookContactSendTemplateListResponseBody) SetRequestId(v string) *ReadWebhookContactSendTemplateListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReadWebhookContactSendTemplateListResponseBody) SetSuccess(v bool) *ReadWebhookContactSendTemplateListResponseBody {
	s.Success = &v
	return s
}

func (s *ReadWebhookContactSendTemplateListResponseBody) Validate() error {
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

type ReadWebhookContactSendTemplateListResponseBodyData struct {
	// The template code.
	//
	// example:
	//
	// lark
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The template.
	//
	// example:
	//
	// /
	Template interface{} `json:"Template,omitempty" xml:"Template,omitempty"`
}

func (s ReadWebhookContactSendTemplateListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ReadWebhookContactSendTemplateListResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReadWebhookContactSendTemplateListResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *ReadWebhookContactSendTemplateListResponseBodyData) GetTemplate() interface{} {
	return s.Template
}

func (s *ReadWebhookContactSendTemplateListResponseBodyData) SetCode(v string) *ReadWebhookContactSendTemplateListResponseBodyData {
	s.Code = &v
	return s
}

func (s *ReadWebhookContactSendTemplateListResponseBodyData) SetTemplate(v interface{}) *ReadWebhookContactSendTemplateListResponseBodyData {
	s.Template = v
	return s
}

func (s *ReadWebhookContactSendTemplateListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
