// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAxbBindFixedLineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteAxbBindFixedLineResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DeleteAxbBindFixedLineResponseBody
	GetCode() *string
	SetData(v *DeleteAxbBindFixedLineResponseBodyData) *DeleteAxbBindFixedLineResponseBody
	GetData() *DeleteAxbBindFixedLineResponseBodyData
	SetMessage(v string) *DeleteAxbBindFixedLineResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAxbBindFixedLineResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteAxbBindFixedLineResponseBody
	GetSuccess() *bool
}

type DeleteAxbBindFixedLineResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// isp.SYSTEM_ERROR
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DeleteAxbBindFixedLineResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// F036366A-0182-5066-A686-19C4C82F2D51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAxbBindFixedLineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAxbBindFixedLineResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAxbBindFixedLineResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteAxbBindFixedLineResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteAxbBindFixedLineResponseBody) GetData() *DeleteAxbBindFixedLineResponseBodyData {
	return s.Data
}

func (s *DeleteAxbBindFixedLineResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAxbBindFixedLineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAxbBindFixedLineResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAxbBindFixedLineResponseBody) SetAccessDeniedDetail(v string) *DeleteAxbBindFixedLineResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteAxbBindFixedLineResponseBody) SetCode(v string) *DeleteAxbBindFixedLineResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAxbBindFixedLineResponseBody) SetData(v *DeleteAxbBindFixedLineResponseBodyData) *DeleteAxbBindFixedLineResponseBody {
	s.Data = v
	return s
}

func (s *DeleteAxbBindFixedLineResponseBody) SetMessage(v string) *DeleteAxbBindFixedLineResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAxbBindFixedLineResponseBody) SetRequestId(v string) *DeleteAxbBindFixedLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAxbBindFixedLineResponseBody) SetSuccess(v bool) *DeleteAxbBindFixedLineResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAxbBindFixedLineResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteAxbBindFixedLineResponseBodyData struct {
	// 响应码  0-成功
	//
	// example:
	//
	// 0
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 响应消息
	//
	// example:
	//
	// 示例值示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 是否处理成功  true-成功  false-失败
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAxbBindFixedLineResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteAxbBindFixedLineResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteAxbBindFixedLineResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *DeleteAxbBindFixedLineResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *DeleteAxbBindFixedLineResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAxbBindFixedLineResponseBodyData) SetCode(v string) *DeleteAxbBindFixedLineResponseBodyData {
	s.Code = &v
	return s
}

func (s *DeleteAxbBindFixedLineResponseBodyData) SetMessage(v string) *DeleteAxbBindFixedLineResponseBodyData {
	s.Message = &v
	return s
}

func (s *DeleteAxbBindFixedLineResponseBodyData) SetSuccess(v bool) *DeleteAxbBindFixedLineResponseBodyData {
	s.Success = &v
	return s
}

func (s *DeleteAxbBindFixedLineResponseBodyData) Validate() error {
	return dara.Validate(s)
}
