// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAliyunPKByAliyunIDResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAliyunPKByAliyunIDResponseBody
	GetCode() *string
	SetData(v string) *GetAliyunPKByAliyunIDResponseBody
	GetData() *string
	SetMsg(v string) *GetAliyunPKByAliyunIDResponseBody
	GetMsg() *string
	SetRequestId(v string) *GetAliyunPKByAliyunIDResponseBody
	GetRequestId() *string
}

type GetAliyunPKByAliyunIDResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Msg       *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAliyunPKByAliyunIDResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAliyunPKByAliyunIDResponseBody) GoString() string {
	return s.String()
}

func (s *GetAliyunPKByAliyunIDResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAliyunPKByAliyunIDResponseBody) GetData() *string {
	return s.Data
}

func (s *GetAliyunPKByAliyunIDResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *GetAliyunPKByAliyunIDResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAliyunPKByAliyunIDResponseBody) SetCode(v string) *GetAliyunPKByAliyunIDResponseBody {
	s.Code = &v
	return s
}

func (s *GetAliyunPKByAliyunIDResponseBody) SetData(v string) *GetAliyunPKByAliyunIDResponseBody {
	s.Data = &v
	return s
}

func (s *GetAliyunPKByAliyunIDResponseBody) SetMsg(v string) *GetAliyunPKByAliyunIDResponseBody {
	s.Msg = &v
	return s
}

func (s *GetAliyunPKByAliyunIDResponseBody) SetRequestId(v string) *GetAliyunPKByAliyunIDResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAliyunPKByAliyunIDResponseBody) Validate() error {
	return dara.Validate(s)
}
