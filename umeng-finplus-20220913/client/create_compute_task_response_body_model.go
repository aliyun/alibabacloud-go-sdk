// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComputeTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateComputeTaskResponseBody
	GetCode() *string
	SetData(v int64) *CreateComputeTaskResponseBody
	GetData() *int64
	SetMsg(v string) *CreateComputeTaskResponseBody
	GetMsg() *string
	SetRequestId(v string) *CreateComputeTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateComputeTaskResponseBody
	GetSuccess() *bool
}

type CreateComputeTaskResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Msg       *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateComputeTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateComputeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateComputeTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateComputeTaskResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateComputeTaskResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *CreateComputeTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateComputeTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateComputeTaskResponseBody) SetCode(v string) *CreateComputeTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateComputeTaskResponseBody) SetData(v int64) *CreateComputeTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CreateComputeTaskResponseBody) SetMsg(v string) *CreateComputeTaskResponseBody {
	s.Msg = &v
	return s
}

func (s *CreateComputeTaskResponseBody) SetRequestId(v string) *CreateComputeTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateComputeTaskResponseBody) SetSuccess(v bool) *CreateComputeTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateComputeTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
