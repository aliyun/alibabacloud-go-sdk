// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLlmSmartCallEncryptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *LlmSmartCallEncryptResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *LlmSmartCallEncryptResponseBody
	GetCode() *string
	SetData(v *LlmSmartCallEncryptResponseBodyData) *LlmSmartCallEncryptResponseBody
	GetData() *LlmSmartCallEncryptResponseBodyData
	SetMessage(v string) *LlmSmartCallEncryptResponseBody
	GetMessage() *string
	SetRequestId(v string) *LlmSmartCallEncryptResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *LlmSmartCallEncryptResponseBody
	GetSuccess() *bool
}

type LlmSmartCallEncryptResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// ok
	Code    *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *LlmSmartCallEncryptResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// F92F9749-105E-518F-8B08-CF16EF36A0E2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s LlmSmartCallEncryptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LlmSmartCallEncryptResponseBody) GoString() string {
	return s.String()
}

func (s *LlmSmartCallEncryptResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *LlmSmartCallEncryptResponseBody) GetCode() *string {
	return s.Code
}

func (s *LlmSmartCallEncryptResponseBody) GetData() *LlmSmartCallEncryptResponseBodyData {
	return s.Data
}

func (s *LlmSmartCallEncryptResponseBody) GetMessage() *string {
	return s.Message
}

func (s *LlmSmartCallEncryptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LlmSmartCallEncryptResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *LlmSmartCallEncryptResponseBody) SetAccessDeniedDetail(v string) *LlmSmartCallEncryptResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *LlmSmartCallEncryptResponseBody) SetCode(v string) *LlmSmartCallEncryptResponseBody {
	s.Code = &v
	return s
}

func (s *LlmSmartCallEncryptResponseBody) SetData(v *LlmSmartCallEncryptResponseBodyData) *LlmSmartCallEncryptResponseBody {
	s.Data = v
	return s
}

func (s *LlmSmartCallEncryptResponseBody) SetMessage(v string) *LlmSmartCallEncryptResponseBody {
	s.Message = &v
	return s
}

func (s *LlmSmartCallEncryptResponseBody) SetRequestId(v string) *LlmSmartCallEncryptResponseBody {
	s.RequestId = &v
	return s
}

func (s *LlmSmartCallEncryptResponseBody) SetSuccess(v bool) *LlmSmartCallEncryptResponseBody {
	s.Success = &v
	return s
}

func (s *LlmSmartCallEncryptResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type LlmSmartCallEncryptResponseBodyData struct {
	// example:
	//
	// 149922088206^136666368206
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
}

func (s LlmSmartCallEncryptResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s LlmSmartCallEncryptResponseBodyData) GoString() string {
	return s.String()
}

func (s *LlmSmartCallEncryptResponseBodyData) GetCallId() *string {
	return s.CallId
}

func (s *LlmSmartCallEncryptResponseBodyData) SetCallId(v string) *LlmSmartCallEncryptResponseBodyData {
	s.CallId = &v
	return s
}

func (s *LlmSmartCallEncryptResponseBodyData) Validate() error {
	return dara.Validate(s)
}
