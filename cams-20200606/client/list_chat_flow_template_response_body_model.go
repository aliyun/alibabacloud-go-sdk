// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatFlowTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListChatFlowTemplateResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ListChatFlowTemplateResponseBody
	GetCode() *string
	SetData(v *ListChatFlowTemplateResponseBodyData) *ListChatFlowTemplateResponseBody
	GetData() *ListChatFlowTemplateResponseBodyData
	SetMessage(v string) *ListChatFlowTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListChatFlowTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListChatFlowTemplateResponseBody
	GetSuccess() *bool
}

type ListChatFlowTemplateResponseBody struct {
	AccessDeniedDetail *string                               `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               *ListChatFlowTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message            *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListChatFlowTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListChatFlowTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ListChatFlowTemplateResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListChatFlowTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListChatFlowTemplateResponseBody) GetData() *ListChatFlowTemplateResponseBodyData {
	return s.Data
}

func (s *ListChatFlowTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListChatFlowTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListChatFlowTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListChatFlowTemplateResponseBody) SetAccessDeniedDetail(v string) *ListChatFlowTemplateResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListChatFlowTemplateResponseBody) SetCode(v string) *ListChatFlowTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *ListChatFlowTemplateResponseBody) SetData(v *ListChatFlowTemplateResponseBodyData) *ListChatFlowTemplateResponseBody {
	s.Data = v
	return s
}

func (s *ListChatFlowTemplateResponseBody) SetMessage(v string) *ListChatFlowTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *ListChatFlowTemplateResponseBody) SetRequestId(v string) *ListChatFlowTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListChatFlowTemplateResponseBody) SetSuccess(v bool) *ListChatFlowTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *ListChatFlowTemplateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListChatFlowTemplateResponseBodyData struct {
	Response map[string]interface{} `json:"Response,omitempty" xml:"Response,omitempty"`
}

func (s ListChatFlowTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListChatFlowTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListChatFlowTemplateResponseBodyData) GetResponse() map[string]interface{} {
	return s.Response
}

func (s *ListChatFlowTemplateResponseBodyData) SetResponse(v map[string]interface{}) *ListChatFlowTemplateResponseBodyData {
	s.Response = v
	return s
}

func (s *ListChatFlowTemplateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
