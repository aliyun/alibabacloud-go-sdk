// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateDataSetResponseBody
	GetCode() *string
	SetData(v int64) *CreateDataSetResponseBody
	GetData() *int64
	SetMsg(v string) *CreateDataSetResponseBody
	GetMsg() *string
	SetRequestId(v string) *CreateDataSetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDataSetResponseBody
	GetSuccess() *bool
}

type CreateDataSetResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Msg       *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDataSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataSetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataSetResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateDataSetResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateDataSetResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *CreateDataSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataSetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDataSetResponseBody) SetCode(v string) *CreateDataSetResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDataSetResponseBody) SetData(v int64) *CreateDataSetResponseBody {
	s.Data = &v
	return s
}

func (s *CreateDataSetResponseBody) SetMsg(v string) *CreateDataSetResponseBody {
	s.Msg = &v
	return s
}

func (s *CreateDataSetResponseBody) SetRequestId(v string) *CreateDataSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataSetResponseBody) SetSuccess(v bool) *CreateDataSetResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDataSetResponseBody) Validate() error {
	return dara.Validate(s)
}
