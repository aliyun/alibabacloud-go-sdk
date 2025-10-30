// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAxnExtensionBindFixedLineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateAxnExtensionBindFixedLineResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *UpdateAxnExtensionBindFixedLineResponseBody
	GetCode() *string
	SetData(v *UpdateAxnExtensionBindFixedLineResponseBodyData) *UpdateAxnExtensionBindFixedLineResponseBody
	GetData() *UpdateAxnExtensionBindFixedLineResponseBodyData
	SetMessage(v string) *UpdateAxnExtensionBindFixedLineResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAxnExtensionBindFixedLineResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAxnExtensionBindFixedLineResponseBody
	GetSuccess() *bool
}

type UpdateAxnExtensionBindFixedLineResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UpdateAxnExtensionBindFixedLineResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 31031C54-7727-5057-9ED1-FA276B64205E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAxnExtensionBindFixedLineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAxnExtensionBindFixedLineResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAxnExtensionBindFixedLineResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateAxnExtensionBindFixedLineResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateAxnExtensionBindFixedLineResponseBody) GetData() *UpdateAxnExtensionBindFixedLineResponseBodyData {
	return s.Data
}

func (s *UpdateAxnExtensionBindFixedLineResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAxnExtensionBindFixedLineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAxnExtensionBindFixedLineResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAxnExtensionBindFixedLineResponseBody) SetAccessDeniedDetail(v string) *UpdateAxnExtensionBindFixedLineResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineResponseBody) SetCode(v string) *UpdateAxnExtensionBindFixedLineResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineResponseBody) SetData(v *UpdateAxnExtensionBindFixedLineResponseBodyData) *UpdateAxnExtensionBindFixedLineResponseBody {
	s.Data = v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineResponseBody) SetMessage(v string) *UpdateAxnExtensionBindFixedLineResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineResponseBody) SetRequestId(v string) *UpdateAxnExtensionBindFixedLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineResponseBody) SetSuccess(v bool) *UpdateAxnExtensionBindFixedLineResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAxnExtensionBindFixedLineResponseBodyData struct {
	// 响应码 0：成功，其它失败，具体见文档
	//
	// example:
	//
	// 0
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 描述信息
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAxnExtensionBindFixedLineResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateAxnExtensionBindFixedLineResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateAxnExtensionBindFixedLineResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *UpdateAxnExtensionBindFixedLineResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *UpdateAxnExtensionBindFixedLineResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAxnExtensionBindFixedLineResponseBodyData) SetCode(v string) *UpdateAxnExtensionBindFixedLineResponseBodyData {
	s.Code = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineResponseBodyData) SetMessage(v string) *UpdateAxnExtensionBindFixedLineResponseBodyData {
	s.Message = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineResponseBodyData) SetSuccess(v bool) *UpdateAxnExtensionBindFixedLineResponseBodyData {
	s.Success = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineResponseBodyData) Validate() error {
	return dara.Validate(s)
}
