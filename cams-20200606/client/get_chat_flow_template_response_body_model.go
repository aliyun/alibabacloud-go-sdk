// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatFlowTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetChatFlowTemplateResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetChatFlowTemplateResponseBody
	GetCode() *string
	SetData(v *GetChatFlowTemplateResponseBodyData) *GetChatFlowTemplateResponseBody
	GetData() *GetChatFlowTemplateResponseBodyData
	SetMessage(v string) *GetChatFlowTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetChatFlowTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetChatFlowTemplateResponseBody
	GetSuccess() *bool
}

type GetChatFlowTemplateResponseBody struct {
	AccessDeniedDetail *string                              `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               *GetChatFlowTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message            *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetChatFlowTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetChatFlowTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetChatFlowTemplateResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetChatFlowTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetChatFlowTemplateResponseBody) GetData() *GetChatFlowTemplateResponseBodyData {
	return s.Data
}

func (s *GetChatFlowTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetChatFlowTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetChatFlowTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetChatFlowTemplateResponseBody) SetAccessDeniedDetail(v string) *GetChatFlowTemplateResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetChatFlowTemplateResponseBody) SetCode(v string) *GetChatFlowTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *GetChatFlowTemplateResponseBody) SetData(v *GetChatFlowTemplateResponseBodyData) *GetChatFlowTemplateResponseBody {
	s.Data = v
	return s
}

func (s *GetChatFlowTemplateResponseBody) SetMessage(v string) *GetChatFlowTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *GetChatFlowTemplateResponseBody) SetRequestId(v string) *GetChatFlowTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetChatFlowTemplateResponseBody) SetSuccess(v bool) *GetChatFlowTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *GetChatFlowTemplateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetChatFlowTemplateResponseBodyData struct {
	Response map[string]interface{} `json:"Response,omitempty" xml:"Response,omitempty"`
}

func (s GetChatFlowTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetChatFlowTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetChatFlowTemplateResponseBodyData) GetResponse() map[string]interface{} {
	return s.Response
}

func (s *GetChatFlowTemplateResponseBodyData) SetResponse(v map[string]interface{}) *GetChatFlowTemplateResponseBodyData {
	s.Response = v
	return s
}

func (s *GetChatFlowTemplateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
