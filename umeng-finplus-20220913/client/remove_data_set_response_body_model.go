// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDataSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RemoveDataSetResponseBody
	GetCode() *string
	SetData(v bool) *RemoveDataSetResponseBody
	GetData() *bool
	SetMsg(v string) *RemoveDataSetResponseBody
	GetMsg() *string
	SetRequestId(v string) *RemoveDataSetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RemoveDataSetResponseBody
	GetSuccess() *bool
}

type RemoveDataSetResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Msg       *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveDataSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveDataSetResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveDataSetResponseBody) GetCode() *string {
	return s.Code
}

func (s *RemoveDataSetResponseBody) GetData() *bool {
	return s.Data
}

func (s *RemoveDataSetResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *RemoveDataSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveDataSetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveDataSetResponseBody) SetCode(v string) *RemoveDataSetResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveDataSetResponseBody) SetData(v bool) *RemoveDataSetResponseBody {
	s.Data = &v
	return s
}

func (s *RemoveDataSetResponseBody) SetMsg(v string) *RemoveDataSetResponseBody {
	s.Msg = &v
	return s
}

func (s *RemoveDataSetResponseBody) SetRequestId(v string) *RemoveDataSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveDataSetResponseBody) SetSuccess(v bool) *RemoveDataSetResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveDataSetResponseBody) Validate() error {
	return dara.Validate(s)
}
