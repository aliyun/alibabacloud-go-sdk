// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateKnowLedgeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ValidateKnowLedgeResponseBody
	GetCode() *string
	SetData(v bool) *ValidateKnowLedgeResponseBody
	GetData() *bool
	SetMsg(v string) *ValidateKnowLedgeResponseBody
	GetMsg() *string
	SetRequestId(v string) *ValidateKnowLedgeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ValidateKnowLedgeResponseBody
	GetSuccess() *bool
}

type ValidateKnowLedgeResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Msg       *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ValidateKnowLedgeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ValidateKnowLedgeResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateKnowLedgeResponseBody) GetCode() *string {
	return s.Code
}

func (s *ValidateKnowLedgeResponseBody) GetData() *bool {
	return s.Data
}

func (s *ValidateKnowLedgeResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *ValidateKnowLedgeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ValidateKnowLedgeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ValidateKnowLedgeResponseBody) SetCode(v string) *ValidateKnowLedgeResponseBody {
	s.Code = &v
	return s
}

func (s *ValidateKnowLedgeResponseBody) SetData(v bool) *ValidateKnowLedgeResponseBody {
	s.Data = &v
	return s
}

func (s *ValidateKnowLedgeResponseBody) SetMsg(v string) *ValidateKnowLedgeResponseBody {
	s.Msg = &v
	return s
}

func (s *ValidateKnowLedgeResponseBody) SetRequestId(v string) *ValidateKnowLedgeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateKnowLedgeResponseBody) SetSuccess(v bool) *ValidateKnowLedgeResponseBody {
	s.Success = &v
	return s
}

func (s *ValidateKnowLedgeResponseBody) Validate() error {
	return dara.Validate(s)
}
