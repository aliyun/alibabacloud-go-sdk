// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComputeTask2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateComputeTask2ResponseBody
	GetCode() *string
	SetData(v int64) *CreateComputeTask2ResponseBody
	GetData() *int64
	SetMsg(v string) *CreateComputeTask2ResponseBody
	GetMsg() *string
	SetRequestId(v string) *CreateComputeTask2ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateComputeTask2ResponseBody
	GetSuccess() *bool
}

type CreateComputeTask2ResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Msg       *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateComputeTask2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateComputeTask2ResponseBody) GoString() string {
	return s.String()
}

func (s *CreateComputeTask2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateComputeTask2ResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateComputeTask2ResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *CreateComputeTask2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateComputeTask2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateComputeTask2ResponseBody) SetCode(v string) *CreateComputeTask2ResponseBody {
	s.Code = &v
	return s
}

func (s *CreateComputeTask2ResponseBody) SetData(v int64) *CreateComputeTask2ResponseBody {
	s.Data = &v
	return s
}

func (s *CreateComputeTask2ResponseBody) SetMsg(v string) *CreateComputeTask2ResponseBody {
	s.Msg = &v
	return s
}

func (s *CreateComputeTask2ResponseBody) SetRequestId(v string) *CreateComputeTask2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateComputeTask2ResponseBody) SetSuccess(v bool) *CreateComputeTask2ResponseBody {
	s.Success = &v
	return s
}

func (s *CreateComputeTask2ResponseBody) Validate() error {
	return dara.Validate(s)
}
