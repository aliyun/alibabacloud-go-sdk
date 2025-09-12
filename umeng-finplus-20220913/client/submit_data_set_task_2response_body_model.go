// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDataSetTask2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitDataSetTask2ResponseBody
	GetCode() *string
	SetData(v bool) *SubmitDataSetTask2ResponseBody
	GetData() *bool
	SetMsg(v string) *SubmitDataSetTask2ResponseBody
	GetMsg() *string
	SetRequestId(v string) *SubmitDataSetTask2ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitDataSetTask2ResponseBody
	GetSuccess() *bool
}

type SubmitDataSetTask2ResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Msg       *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitDataSetTask2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitDataSetTask2ResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitDataSetTask2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitDataSetTask2ResponseBody) GetData() *bool {
	return s.Data
}

func (s *SubmitDataSetTask2ResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *SubmitDataSetTask2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitDataSetTask2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitDataSetTask2ResponseBody) SetCode(v string) *SubmitDataSetTask2ResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitDataSetTask2ResponseBody) SetData(v bool) *SubmitDataSetTask2ResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitDataSetTask2ResponseBody) SetMsg(v string) *SubmitDataSetTask2ResponseBody {
	s.Msg = &v
	return s
}

func (s *SubmitDataSetTask2ResponseBody) SetRequestId(v string) *SubmitDataSetTask2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitDataSetTask2ResponseBody) SetSuccess(v bool) *SubmitDataSetTask2ResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitDataSetTask2ResponseBody) Validate() error {
	return dara.Validate(s)
}
