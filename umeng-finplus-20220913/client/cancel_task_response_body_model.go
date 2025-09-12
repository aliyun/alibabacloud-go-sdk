// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CancelTaskResponseBody
	GetCode() *string
	SetData(v bool) *CancelTaskResponseBody
	GetData() *bool
	SetMsg(v string) *CancelTaskResponseBody
	GetMsg() *string
	SetRequestId(v string) *CancelTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CancelTaskResponseBody
	GetSuccess() *bool
}

type CancelTaskResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Msg       *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CancelTaskResponseBody) GetData() *bool {
	return s.Data
}

func (s *CancelTaskResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *CancelTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CancelTaskResponseBody) SetCode(v string) *CancelTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CancelTaskResponseBody) SetData(v bool) *CancelTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CancelTaskResponseBody) SetMsg(v string) *CancelTaskResponseBody {
	s.Msg = &v
	return s
}

func (s *CancelTaskResponseBody) SetRequestId(v string) *CancelTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelTaskResponseBody) SetSuccess(v bool) *CancelTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CancelTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
