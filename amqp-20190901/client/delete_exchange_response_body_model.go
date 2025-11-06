// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExchangeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteExchangeResponseBody
	GetCode() *int32
	SetData(v string) *DeleteExchangeResponseBody
	GetData() *string
	SetMessage(v string) *DeleteExchangeResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteExchangeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteExchangeResponseBody
	GetSuccess() *bool
}

type DeleteExchangeResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteExchangeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteExchangeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExchangeResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteExchangeResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteExchangeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteExchangeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteExchangeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteExchangeResponseBody) SetCode(v int32) *DeleteExchangeResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteExchangeResponseBody) SetData(v string) *DeleteExchangeResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteExchangeResponseBody) SetMessage(v string) *DeleteExchangeResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteExchangeResponseBody) SetRequestId(v string) *DeleteExchangeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteExchangeResponseBody) SetSuccess(v bool) *DeleteExchangeResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteExchangeResponseBody) Validate() error {
	return dara.Validate(s)
}
