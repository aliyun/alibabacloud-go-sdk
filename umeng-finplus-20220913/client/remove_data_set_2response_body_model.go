// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDataSet2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RemoveDataSet2ResponseBody
	GetCode() *string
	SetData(v bool) *RemoveDataSet2ResponseBody
	GetData() *bool
	SetMsg(v string) *RemoveDataSet2ResponseBody
	GetMsg() *string
	SetRequestId(v string) *RemoveDataSet2ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RemoveDataSet2ResponseBody
	GetSuccess() *bool
}

type RemoveDataSet2ResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Msg       *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveDataSet2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveDataSet2ResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveDataSet2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *RemoveDataSet2ResponseBody) GetData() *bool {
	return s.Data
}

func (s *RemoveDataSet2ResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *RemoveDataSet2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveDataSet2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveDataSet2ResponseBody) SetCode(v string) *RemoveDataSet2ResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveDataSet2ResponseBody) SetData(v bool) *RemoveDataSet2ResponseBody {
	s.Data = &v
	return s
}

func (s *RemoveDataSet2ResponseBody) SetMsg(v string) *RemoveDataSet2ResponseBody {
	s.Msg = &v
	return s
}

func (s *RemoveDataSet2ResponseBody) SetRequestId(v string) *RemoveDataSet2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveDataSet2ResponseBody) SetSuccess(v bool) *RemoveDataSet2ResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveDataSet2ResponseBody) Validate() error {
	return dara.Validate(s)
}
