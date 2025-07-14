// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebApplicationTrafficConfigBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *WebApplicationTrafficConfigBody
	GetCode() *int32
	SetData(v *WebTrafficConfig) *WebApplicationTrafficConfigBody
	GetData() *WebTrafficConfig
	SetMessage(v string) *WebApplicationTrafficConfigBody
	GetMessage() *string
	SetRequestId(v string) *WebApplicationTrafficConfigBody
	GetRequestId() *string
	SetSuccess(v bool) *WebApplicationTrafficConfigBody
	GetSuccess() *bool
}

type WebApplicationTrafficConfigBody struct {
	Code      *int32            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *WebTrafficConfig `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s WebApplicationTrafficConfigBody) String() string {
	return dara.Prettify(s)
}

func (s WebApplicationTrafficConfigBody) GoString() string {
	return s.String()
}

func (s *WebApplicationTrafficConfigBody) GetCode() *int32 {
	return s.Code
}

func (s *WebApplicationTrafficConfigBody) GetData() *WebTrafficConfig {
	return s.Data
}

func (s *WebApplicationTrafficConfigBody) GetMessage() *string {
	return s.Message
}

func (s *WebApplicationTrafficConfigBody) GetRequestId() *string {
	return s.RequestId
}

func (s *WebApplicationTrafficConfigBody) GetSuccess() *bool {
	return s.Success
}

func (s *WebApplicationTrafficConfigBody) SetCode(v int32) *WebApplicationTrafficConfigBody {
	s.Code = &v
	return s
}

func (s *WebApplicationTrafficConfigBody) SetData(v *WebTrafficConfig) *WebApplicationTrafficConfigBody {
	s.Data = v
	return s
}

func (s *WebApplicationTrafficConfigBody) SetMessage(v string) *WebApplicationTrafficConfigBody {
	s.Message = &v
	return s
}

func (s *WebApplicationTrafficConfigBody) SetRequestId(v string) *WebApplicationTrafficConfigBody {
	s.RequestId = &v
	return s
}

func (s *WebApplicationTrafficConfigBody) SetSuccess(v bool) *WebApplicationTrafficConfigBody {
	s.Success = &v
	return s
}

func (s *WebApplicationTrafficConfigBody) Validate() error {
	return dara.Validate(s)
}
