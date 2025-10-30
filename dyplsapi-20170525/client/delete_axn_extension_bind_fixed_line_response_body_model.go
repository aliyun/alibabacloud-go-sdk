// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAxnExtensionBindFixedLineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteAxnExtensionBindFixedLineResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DeleteAxnExtensionBindFixedLineResponseBody
	GetCode() *string
	SetData(v *DeleteAxnExtensionBindFixedLineResponseBodyData) *DeleteAxnExtensionBindFixedLineResponseBody
	GetData() *DeleteAxnExtensionBindFixedLineResponseBodyData
	SetMessage(v string) *DeleteAxnExtensionBindFixedLineResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAxnExtensionBindFixedLineResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteAxnExtensionBindFixedLineResponseBody
	GetSuccess() *bool
}

type DeleteAxnExtensionBindFixedLineResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DeleteAxnExtensionBindFixedLineResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// conflict with subs id=1000203635098305, phoneA conflict
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3DA9D6DF-C5FA-5A0D-B6C2-547B1FD1F9B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAxnExtensionBindFixedLineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAxnExtensionBindFixedLineResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAxnExtensionBindFixedLineResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteAxnExtensionBindFixedLineResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteAxnExtensionBindFixedLineResponseBody) GetData() *DeleteAxnExtensionBindFixedLineResponseBodyData {
	return s.Data
}

func (s *DeleteAxnExtensionBindFixedLineResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAxnExtensionBindFixedLineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAxnExtensionBindFixedLineResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAxnExtensionBindFixedLineResponseBody) SetAccessDeniedDetail(v string) *DeleteAxnExtensionBindFixedLineResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteAxnExtensionBindFixedLineResponseBody) SetCode(v string) *DeleteAxnExtensionBindFixedLineResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAxnExtensionBindFixedLineResponseBody) SetData(v *DeleteAxnExtensionBindFixedLineResponseBodyData) *DeleteAxnExtensionBindFixedLineResponseBody {
	s.Data = v
	return s
}

func (s *DeleteAxnExtensionBindFixedLineResponseBody) SetMessage(v string) *DeleteAxnExtensionBindFixedLineResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAxnExtensionBindFixedLineResponseBody) SetRequestId(v string) *DeleteAxnExtensionBindFixedLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAxnExtensionBindFixedLineResponseBody) SetSuccess(v bool) *DeleteAxnExtensionBindFixedLineResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAxnExtensionBindFixedLineResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteAxnExtensionBindFixedLineResponseBodyData struct {
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

func (s DeleteAxnExtensionBindFixedLineResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteAxnExtensionBindFixedLineResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteAxnExtensionBindFixedLineResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *DeleteAxnExtensionBindFixedLineResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *DeleteAxnExtensionBindFixedLineResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAxnExtensionBindFixedLineResponseBodyData) SetCode(v string) *DeleteAxnExtensionBindFixedLineResponseBodyData {
	s.Code = &v
	return s
}

func (s *DeleteAxnExtensionBindFixedLineResponseBodyData) SetMessage(v string) *DeleteAxnExtensionBindFixedLineResponseBodyData {
	s.Message = &v
	return s
}

func (s *DeleteAxnExtensionBindFixedLineResponseBodyData) SetSuccess(v bool) *DeleteAxnExtensionBindFixedLineResponseBodyData {
	s.Success = &v
	return s
}

func (s *DeleteAxnExtensionBindFixedLineResponseBodyData) Validate() error {
	return dara.Validate(s)
}
