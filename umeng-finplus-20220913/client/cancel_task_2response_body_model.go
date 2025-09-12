// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelTask2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CancelTask2ResponseBody
	GetCode() *string
	SetData(v bool) *CancelTask2ResponseBody
	GetData() *bool
	SetMsg(v string) *CancelTask2ResponseBody
	GetMsg() *string
	SetRequestId(v string) *CancelTask2ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CancelTask2ResponseBody
	GetSuccess() *bool
}

type CancelTask2ResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Msg       *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelTask2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelTask2ResponseBody) GoString() string {
	return s.String()
}

func (s *CancelTask2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *CancelTask2ResponseBody) GetData() *bool {
	return s.Data
}

func (s *CancelTask2ResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *CancelTask2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelTask2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CancelTask2ResponseBody) SetCode(v string) *CancelTask2ResponseBody {
	s.Code = &v
	return s
}

func (s *CancelTask2ResponseBody) SetData(v bool) *CancelTask2ResponseBody {
	s.Data = &v
	return s
}

func (s *CancelTask2ResponseBody) SetMsg(v string) *CancelTask2ResponseBody {
	s.Msg = &v
	return s
}

func (s *CancelTask2ResponseBody) SetRequestId(v string) *CancelTask2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelTask2ResponseBody) SetSuccess(v bool) *CancelTask2ResponseBody {
	s.Success = &v
	return s
}

func (s *CancelTask2ResponseBody) Validate() error {
	return dara.Validate(s)
}
