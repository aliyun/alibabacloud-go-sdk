// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWhatsappCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *WhatsappCallResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *WhatsappCallResponseBody
	GetCode() *string
	SetMessage(v string) *WhatsappCallResponseBody
	GetMessage() *string
	SetModel(v *WhatsappCallResponseBodyModel) *WhatsappCallResponseBody
	GetModel() *WhatsappCallResponseBodyModel
	SetRequestId(v string) *WhatsappCallResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *WhatsappCallResponseBody
	GetSuccess() *bool
}

type WhatsappCallResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Message *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *WhatsappCallResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// 示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s WhatsappCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s WhatsappCallResponseBody) GoString() string {
	return s.String()
}

func (s *WhatsappCallResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *WhatsappCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *WhatsappCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *WhatsappCallResponseBody) GetModel() *WhatsappCallResponseBodyModel {
	return s.Model
}

func (s *WhatsappCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *WhatsappCallResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *WhatsappCallResponseBody) SetAccessDeniedDetail(v string) *WhatsappCallResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *WhatsappCallResponseBody) SetCode(v string) *WhatsappCallResponseBody {
	s.Code = &v
	return s
}

func (s *WhatsappCallResponseBody) SetMessage(v string) *WhatsappCallResponseBody {
	s.Message = &v
	return s
}

func (s *WhatsappCallResponseBody) SetModel(v *WhatsappCallResponseBodyModel) *WhatsappCallResponseBody {
	s.Model = v
	return s
}

func (s *WhatsappCallResponseBody) SetRequestId(v string) *WhatsappCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *WhatsappCallResponseBody) SetSuccess(v bool) *WhatsappCallResponseBody {
	s.Success = &v
	return s
}

func (s *WhatsappCallResponseBody) Validate() error {
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type WhatsappCallResponseBodyModel struct {
	// example:
	//
	// 示例值示例值
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
}

func (s WhatsappCallResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s WhatsappCallResponseBodyModel) GoString() string {
	return s.String()
}

func (s *WhatsappCallResponseBodyModel) GetCallId() *string {
	return s.CallId
}

func (s *WhatsappCallResponseBodyModel) SetCallId(v string) *WhatsappCallResponseBodyModel {
	s.CallId = &v
	return s
}

func (s *WhatsappCallResponseBodyModel) Validate() error {
	return dara.Validate(s)
}
