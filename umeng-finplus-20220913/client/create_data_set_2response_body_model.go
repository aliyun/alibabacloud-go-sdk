// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataSet2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateDataSet2ResponseBody
	GetCode() *string
	SetData(v int64) *CreateDataSet2ResponseBody
	GetData() *int64
	SetMsg(v string) *CreateDataSet2ResponseBody
	GetMsg() *string
	SetRequestId(v string) *CreateDataSet2ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDataSet2ResponseBody
	GetSuccess() *bool
}

type CreateDataSet2ResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Msg       *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDataSet2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataSet2ResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataSet2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateDataSet2ResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateDataSet2ResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *CreateDataSet2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataSet2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDataSet2ResponseBody) SetCode(v string) *CreateDataSet2ResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDataSet2ResponseBody) SetData(v int64) *CreateDataSet2ResponseBody {
	s.Data = &v
	return s
}

func (s *CreateDataSet2ResponseBody) SetMsg(v string) *CreateDataSet2ResponseBody {
	s.Msg = &v
	return s
}

func (s *CreateDataSet2ResponseBody) SetRequestId(v string) *CreateDataSet2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataSet2ResponseBody) SetSuccess(v bool) *CreateDataSet2ResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDataSet2ResponseBody) Validate() error {
	return dara.Validate(s)
}
