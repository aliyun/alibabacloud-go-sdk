// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDataSetTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitDataSetTaskResponseBody
	GetCode() *string
	SetData(v bool) *SubmitDataSetTaskResponseBody
	GetData() *bool
	SetMsg(v string) *SubmitDataSetTaskResponseBody
	GetMsg() *string
	SetRequestId(v string) *SubmitDataSetTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitDataSetTaskResponseBody
	GetSuccess() *bool
}

type SubmitDataSetTaskResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Msg       *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitDataSetTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitDataSetTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitDataSetTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitDataSetTaskResponseBody) GetData() *bool {
	return s.Data
}

func (s *SubmitDataSetTaskResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *SubmitDataSetTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitDataSetTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitDataSetTaskResponseBody) SetCode(v string) *SubmitDataSetTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitDataSetTaskResponseBody) SetData(v bool) *SubmitDataSetTaskResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitDataSetTaskResponseBody) SetMsg(v string) *SubmitDataSetTaskResponseBody {
	s.Msg = &v
	return s
}

func (s *SubmitDataSetTaskResponseBody) SetRequestId(v string) *SubmitDataSetTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitDataSetTaskResponseBody) SetSuccess(v bool) *SubmitDataSetTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitDataSetTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
