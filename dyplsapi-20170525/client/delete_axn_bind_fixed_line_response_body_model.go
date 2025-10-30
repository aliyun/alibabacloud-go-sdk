// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAxnBindFixedLineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteAxnBindFixedLineResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DeleteAxnBindFixedLineResponseBody
	GetCode() *string
	SetData(v *DeleteAxnBindFixedLineResponseBodyData) *DeleteAxnBindFixedLineResponseBody
	GetData() *DeleteAxnBindFixedLineResponseBodyData
	SetMessage(v string) *DeleteAxnBindFixedLineResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAxnBindFixedLineResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteAxnBindFixedLineResponseBody
	GetSuccess() *bool
}

type DeleteAxnBindFixedLineResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DeleteAxnBindFixedLineResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// AE2D6997-643A-59CB-9B3C-918572E5CEAA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAxnBindFixedLineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAxnBindFixedLineResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAxnBindFixedLineResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteAxnBindFixedLineResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteAxnBindFixedLineResponseBody) GetData() *DeleteAxnBindFixedLineResponseBodyData {
	return s.Data
}

func (s *DeleteAxnBindFixedLineResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAxnBindFixedLineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAxnBindFixedLineResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAxnBindFixedLineResponseBody) SetAccessDeniedDetail(v string) *DeleteAxnBindFixedLineResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteAxnBindFixedLineResponseBody) SetCode(v string) *DeleteAxnBindFixedLineResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAxnBindFixedLineResponseBody) SetData(v *DeleteAxnBindFixedLineResponseBodyData) *DeleteAxnBindFixedLineResponseBody {
	s.Data = v
	return s
}

func (s *DeleteAxnBindFixedLineResponseBody) SetMessage(v string) *DeleteAxnBindFixedLineResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAxnBindFixedLineResponseBody) SetRequestId(v string) *DeleteAxnBindFixedLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAxnBindFixedLineResponseBody) SetSuccess(v bool) *DeleteAxnBindFixedLineResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAxnBindFixedLineResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteAxnBindFixedLineResponseBodyData struct {
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

func (s DeleteAxnBindFixedLineResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteAxnBindFixedLineResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteAxnBindFixedLineResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *DeleteAxnBindFixedLineResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *DeleteAxnBindFixedLineResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAxnBindFixedLineResponseBodyData) SetCode(v string) *DeleteAxnBindFixedLineResponseBodyData {
	s.Code = &v
	return s
}

func (s *DeleteAxnBindFixedLineResponseBodyData) SetMessage(v string) *DeleteAxnBindFixedLineResponseBodyData {
	s.Message = &v
	return s
}

func (s *DeleteAxnBindFixedLineResponseBodyData) SetSuccess(v bool) *DeleteAxnBindFixedLineResponseBodyData {
	s.Success = &v
	return s
}

func (s *DeleteAxnBindFixedLineResponseBodyData) Validate() error {
	return dara.Validate(s)
}
