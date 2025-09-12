// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowLedgeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateKnowLedgeResponseBody
	GetCode() *string
	SetData(v bool) *CreateKnowLedgeResponseBody
	GetData() *bool
	SetMsg(v string) *CreateKnowLedgeResponseBody
	GetMsg() *string
	SetRequestId(v string) *CreateKnowLedgeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateKnowLedgeResponseBody
	GetSuccess() *bool
}

type CreateKnowLedgeResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Msg       *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateKnowLedgeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowLedgeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKnowLedgeResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateKnowLedgeResponseBody) GetData() *bool {
	return s.Data
}

func (s *CreateKnowLedgeResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *CreateKnowLedgeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateKnowLedgeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateKnowLedgeResponseBody) SetCode(v string) *CreateKnowLedgeResponseBody {
	s.Code = &v
	return s
}

func (s *CreateKnowLedgeResponseBody) SetData(v bool) *CreateKnowLedgeResponseBody {
	s.Data = &v
	return s
}

func (s *CreateKnowLedgeResponseBody) SetMsg(v string) *CreateKnowLedgeResponseBody {
	s.Msg = &v
	return s
}

func (s *CreateKnowLedgeResponseBody) SetRequestId(v string) *CreateKnowLedgeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateKnowLedgeResponseBody) SetSuccess(v bool) *CreateKnowLedgeResponseBody {
	s.Success = &v
	return s
}

func (s *CreateKnowLedgeResponseBody) Validate() error {
	return dara.Validate(s)
}
