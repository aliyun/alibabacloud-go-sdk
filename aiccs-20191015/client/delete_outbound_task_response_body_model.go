// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOutboundTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteOutboundTaskResponseBody
	GetCode() *string
	SetData(v string) *DeleteOutboundTaskResponseBody
	GetData() *string
	SetMessage(v string) *DeleteOutboundTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteOutboundTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteOutboundTaskResponseBody
	GetSuccess() *bool
}

type DeleteOutboundTaskResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteOutboundTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteOutboundTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOutboundTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteOutboundTaskResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteOutboundTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteOutboundTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteOutboundTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteOutboundTaskResponseBody) SetCode(v string) *DeleteOutboundTaskResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteOutboundTaskResponseBody) SetData(v string) *DeleteOutboundTaskResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteOutboundTaskResponseBody) SetMessage(v string) *DeleteOutboundTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteOutboundTaskResponseBody) SetRequestId(v string) *DeleteOutboundTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteOutboundTaskResponseBody) SetSuccess(v bool) *DeleteOutboundTaskResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteOutboundTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
