// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAxbBindFixedLineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateAxbBindFixedLineResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *UpdateAxbBindFixedLineResponseBody
	GetCode() *string
	SetData(v *UpdateAxbBindFixedLineResponseBodyData) *UpdateAxbBindFixedLineResponseBody
	GetData() *UpdateAxbBindFixedLineResponseBodyData
	SetMessage(v string) *UpdateAxbBindFixedLineResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAxbBindFixedLineResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAxbBindFixedLineResponseBody
	GetSuccess() *bool
}

type UpdateAxbBindFixedLineResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UpdateAxbBindFixedLineResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 73A5C73A-1D97-54B6-B47C-541BE59F84D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAxbBindFixedLineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAxbBindFixedLineResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAxbBindFixedLineResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateAxbBindFixedLineResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateAxbBindFixedLineResponseBody) GetData() *UpdateAxbBindFixedLineResponseBodyData {
	return s.Data
}

func (s *UpdateAxbBindFixedLineResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAxbBindFixedLineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAxbBindFixedLineResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAxbBindFixedLineResponseBody) SetAccessDeniedDetail(v string) *UpdateAxbBindFixedLineResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateAxbBindFixedLineResponseBody) SetCode(v string) *UpdateAxbBindFixedLineResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAxbBindFixedLineResponseBody) SetData(v *UpdateAxbBindFixedLineResponseBodyData) *UpdateAxbBindFixedLineResponseBody {
	s.Data = v
	return s
}

func (s *UpdateAxbBindFixedLineResponseBody) SetMessage(v string) *UpdateAxbBindFixedLineResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAxbBindFixedLineResponseBody) SetRequestId(v string) *UpdateAxbBindFixedLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAxbBindFixedLineResponseBody) SetSuccess(v bool) *UpdateAxbBindFixedLineResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAxbBindFixedLineResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAxbBindFixedLineResponseBodyData struct {
	// 响应码 0-成功
	//
	// example:
	//
	// 0
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 响应消息
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 是否处理成功  true-成功
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAxbBindFixedLineResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateAxbBindFixedLineResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateAxbBindFixedLineResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *UpdateAxbBindFixedLineResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *UpdateAxbBindFixedLineResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAxbBindFixedLineResponseBodyData) SetCode(v string) *UpdateAxbBindFixedLineResponseBodyData {
	s.Code = &v
	return s
}

func (s *UpdateAxbBindFixedLineResponseBodyData) SetMessage(v string) *UpdateAxbBindFixedLineResponseBodyData {
	s.Message = &v
	return s
}

func (s *UpdateAxbBindFixedLineResponseBodyData) SetSuccess(v bool) *UpdateAxbBindFixedLineResponseBodyData {
	s.Success = &v
	return s
}

func (s *UpdateAxbBindFixedLineResponseBodyData) Validate() error {
	return dara.Validate(s)
}
