// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLogDeliverySLRResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateLogDeliverySLRResponseBody
	GetCode() *int32
	SetData(v bool) *CreateLogDeliverySLRResponseBody
	GetData() *bool
	SetMessage(v string) *CreateLogDeliverySLRResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateLogDeliverySLRResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateLogDeliverySLRResponseBody
	GetSuccess() *bool
}

type CreateLogDeliverySLRResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateLogDeliverySLRResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLogDeliverySLRResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLogDeliverySLRResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateLogDeliverySLRResponseBody) GetData() *bool {
	return s.Data
}

func (s *CreateLogDeliverySLRResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateLogDeliverySLRResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLogDeliverySLRResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateLogDeliverySLRResponseBody) SetCode(v int32) *CreateLogDeliverySLRResponseBody {
	s.Code = &v
	return s
}

func (s *CreateLogDeliverySLRResponseBody) SetData(v bool) *CreateLogDeliverySLRResponseBody {
	s.Data = &v
	return s
}

func (s *CreateLogDeliverySLRResponseBody) SetMessage(v string) *CreateLogDeliverySLRResponseBody {
	s.Message = &v
	return s
}

func (s *CreateLogDeliverySLRResponseBody) SetRequestId(v string) *CreateLogDeliverySLRResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLogDeliverySLRResponseBody) SetSuccess(v bool) *CreateLogDeliverySLRResponseBody {
	s.Success = &v
	return s
}

func (s *CreateLogDeliverySLRResponseBody) Validate() error {
	return dara.Validate(s)
}
