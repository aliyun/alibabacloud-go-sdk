// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExchangeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateExchangeResponseBody
	GetCode() *int32
	SetData(v bool) *CreateExchangeResponseBody
	GetData() *bool
	SetMessage(v string) *CreateExchangeResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateExchangeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateExchangeResponseBody
	GetSuccess() *bool
}

type CreateExchangeResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateExchangeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateExchangeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExchangeResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateExchangeResponseBody) GetData() *bool {
	return s.Data
}

func (s *CreateExchangeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateExchangeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateExchangeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateExchangeResponseBody) SetCode(v int32) *CreateExchangeResponseBody {
	s.Code = &v
	return s
}

func (s *CreateExchangeResponseBody) SetData(v bool) *CreateExchangeResponseBody {
	s.Data = &v
	return s
}

func (s *CreateExchangeResponseBody) SetMessage(v string) *CreateExchangeResponseBody {
	s.Message = &v
	return s
}

func (s *CreateExchangeResponseBody) SetRequestId(v string) *CreateExchangeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateExchangeResponseBody) SetSuccess(v bool) *CreateExchangeResponseBody {
	s.Success = &v
	return s
}

func (s *CreateExchangeResponseBody) Validate() error {
	return dara.Validate(s)
}
