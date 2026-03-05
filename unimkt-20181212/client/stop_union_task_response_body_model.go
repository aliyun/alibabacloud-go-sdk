// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopUnionTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *StopUnionTaskResponseBody
	GetCode() *int64
	SetData(v bool) *StopUnionTaskResponseBody
	GetData() *bool
	SetErrorMsg(v string) *StopUnionTaskResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *StopUnionTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StopUnionTaskResponseBody
	GetSuccess() *bool
}

type StopUnionTaskResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopUnionTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopUnionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StopUnionTaskResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *StopUnionTaskResponseBody) GetData() *bool {
	return s.Data
}

func (s *StopUnionTaskResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *StopUnionTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopUnionTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StopUnionTaskResponseBody) SetCode(v int64) *StopUnionTaskResponseBody {
	s.Code = &v
	return s
}

func (s *StopUnionTaskResponseBody) SetData(v bool) *StopUnionTaskResponseBody {
	s.Data = &v
	return s
}

func (s *StopUnionTaskResponseBody) SetErrorMsg(v string) *StopUnionTaskResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *StopUnionTaskResponseBody) SetRequestId(v string) *StopUnionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopUnionTaskResponseBody) SetSuccess(v bool) *StopUnionTaskResponseBody {
	s.Success = &v
	return s
}

func (s *StopUnionTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
