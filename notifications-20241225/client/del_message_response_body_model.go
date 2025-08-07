// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDelMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DelMessageResponseBody
	GetCode() *string
	SetData(v bool) *DelMessageResponseBody
	GetData() *bool
	SetMessage(v string) *DelMessageResponseBody
	GetMessage() *string
	SetRequestId(v string) *DelMessageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DelMessageResponseBody
	GetSuccess() *bool
}

type DelMessageResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DelMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DelMessageResponseBody) GoString() string {
	return s.String()
}

func (s *DelMessageResponseBody) GetCode() *string {
	return s.Code
}

func (s *DelMessageResponseBody) GetData() *bool {
	return s.Data
}

func (s *DelMessageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DelMessageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DelMessageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DelMessageResponseBody) SetCode(v string) *DelMessageResponseBody {
	s.Code = &v
	return s
}

func (s *DelMessageResponseBody) SetData(v bool) *DelMessageResponseBody {
	s.Data = &v
	return s
}

func (s *DelMessageResponseBody) SetMessage(v string) *DelMessageResponseBody {
	s.Message = &v
	return s
}

func (s *DelMessageResponseBody) SetRequestId(v string) *DelMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DelMessageResponseBody) SetSuccess(v bool) *DelMessageResponseBody {
	s.Success = &v
	return s
}

func (s *DelMessageResponseBody) Validate() error {
	return dara.Validate(s)
}
