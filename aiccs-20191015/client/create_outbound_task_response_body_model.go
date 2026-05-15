// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOutboundTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateOutboundTaskResponseBody
	GetCode() *string
	SetData(v string) *CreateOutboundTaskResponseBody
	GetData() *string
	SetMessage(v string) *CreateOutboundTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateOutboundTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateOutboundTaskResponseBody
	GetSuccess() *bool
}

type CreateOutboundTaskResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateOutboundTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOutboundTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOutboundTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateOutboundTaskResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateOutboundTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateOutboundTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOutboundTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateOutboundTaskResponseBody) SetCode(v string) *CreateOutboundTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateOutboundTaskResponseBody) SetData(v string) *CreateOutboundTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CreateOutboundTaskResponseBody) SetMessage(v string) *CreateOutboundTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateOutboundTaskResponseBody) SetRequestId(v string) *CreateOutboundTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOutboundTaskResponseBody) SetSuccess(v bool) *CreateOutboundTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateOutboundTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
