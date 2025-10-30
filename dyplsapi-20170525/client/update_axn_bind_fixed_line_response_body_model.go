// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAxnBindFixedLineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateAxnBindFixedLineResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *UpdateAxnBindFixedLineResponseBody
	GetCode() *string
	SetData(v *UpdateAxnBindFixedLineResponseBodyData) *UpdateAxnBindFixedLineResponseBody
	GetData() *UpdateAxnBindFixedLineResponseBodyData
	SetMessage(v string) *UpdateAxnBindFixedLineResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAxnBindFixedLineResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAxnBindFixedLineResponseBody
	GetSuccess() *bool
}

type UpdateAxnBindFixedLineResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UpdateAxnBindFixedLineResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3FDD0A8F-34F1-5BD4-AF9F-CD90B3DE7C06
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAxnBindFixedLineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAxnBindFixedLineResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAxnBindFixedLineResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateAxnBindFixedLineResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateAxnBindFixedLineResponseBody) GetData() *UpdateAxnBindFixedLineResponseBodyData {
	return s.Data
}

func (s *UpdateAxnBindFixedLineResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAxnBindFixedLineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAxnBindFixedLineResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAxnBindFixedLineResponseBody) SetAccessDeniedDetail(v string) *UpdateAxnBindFixedLineResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateAxnBindFixedLineResponseBody) SetCode(v string) *UpdateAxnBindFixedLineResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAxnBindFixedLineResponseBody) SetData(v *UpdateAxnBindFixedLineResponseBodyData) *UpdateAxnBindFixedLineResponseBody {
	s.Data = v
	return s
}

func (s *UpdateAxnBindFixedLineResponseBody) SetMessage(v string) *UpdateAxnBindFixedLineResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAxnBindFixedLineResponseBody) SetRequestId(v string) *UpdateAxnBindFixedLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAxnBindFixedLineResponseBody) SetSuccess(v bool) *UpdateAxnBindFixedLineResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAxnBindFixedLineResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAxnBindFixedLineResponseBodyData struct {
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
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAxnBindFixedLineResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateAxnBindFixedLineResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateAxnBindFixedLineResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *UpdateAxnBindFixedLineResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *UpdateAxnBindFixedLineResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAxnBindFixedLineResponseBodyData) SetCode(v string) *UpdateAxnBindFixedLineResponseBodyData {
	s.Code = &v
	return s
}

func (s *UpdateAxnBindFixedLineResponseBodyData) SetMessage(v string) *UpdateAxnBindFixedLineResponseBodyData {
	s.Message = &v
	return s
}

func (s *UpdateAxnBindFixedLineResponseBodyData) SetSuccess(v bool) *UpdateAxnBindFixedLineResponseBodyData {
	s.Success = &v
	return s
}

func (s *UpdateAxnBindFixedLineResponseBodyData) Validate() error {
	return dara.Validate(s)
}
