// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebApplicationBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *WebApplicationBody
	GetCode() *int32
	SetData(v *WebApplication) *WebApplicationBody
	GetData() *WebApplication
	SetMessage(v string) *WebApplicationBody
	GetMessage() *string
	SetRequestId(v string) *WebApplicationBody
	GetRequestId() *string
	SetSuccess(v bool) *WebApplicationBody
	GetSuccess() *bool
}

type WebApplicationBody struct {
	Code      *int32          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *WebApplication `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s WebApplicationBody) String() string {
	return dara.Prettify(s)
}

func (s WebApplicationBody) GoString() string {
	return s.String()
}

func (s *WebApplicationBody) GetCode() *int32 {
	return s.Code
}

func (s *WebApplicationBody) GetData() *WebApplication {
	return s.Data
}

func (s *WebApplicationBody) GetMessage() *string {
	return s.Message
}

func (s *WebApplicationBody) GetRequestId() *string {
	return s.RequestId
}

func (s *WebApplicationBody) GetSuccess() *bool {
	return s.Success
}

func (s *WebApplicationBody) SetCode(v int32) *WebApplicationBody {
	s.Code = &v
	return s
}

func (s *WebApplicationBody) SetData(v *WebApplication) *WebApplicationBody {
	s.Data = v
	return s
}

func (s *WebApplicationBody) SetMessage(v string) *WebApplicationBody {
	s.Message = &v
	return s
}

func (s *WebApplicationBody) SetRequestId(v string) *WebApplicationBody {
	s.RequestId = &v
	return s
}

func (s *WebApplicationBody) SetSuccess(v bool) *WebApplicationBody {
	s.Success = &v
	return s
}

func (s *WebApplicationBody) Validate() error {
	return dara.Validate(s)
}
