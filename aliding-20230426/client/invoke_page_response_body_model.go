// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokePageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *InvokePageResponseBody
	GetData() interface{}
	SetErrorCode(v string) *InvokePageResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *InvokePageResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *InvokePageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InvokePageResponseBody
	GetSuccess() *bool
}

type InvokePageResponseBody struct {
	Data      interface{} `json:"data,omitempty" xml:"data,omitempty"`
	ErrorCode *string     `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string     `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	RequestId *string     `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool       `json:"success,omitempty" xml:"success,omitempty"`
}

func (s InvokePageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InvokePageResponseBody) GoString() string {
	return s.String()
}

func (s *InvokePageResponseBody) GetData() interface{} {
	return s.Data
}

func (s *InvokePageResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *InvokePageResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *InvokePageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InvokePageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InvokePageResponseBody) SetData(v interface{}) *InvokePageResponseBody {
	s.Data = v
	return s
}

func (s *InvokePageResponseBody) SetErrorCode(v string) *InvokePageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *InvokePageResponseBody) SetErrorMsg(v string) *InvokePageResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *InvokePageResponseBody) SetRequestId(v string) *InvokePageResponseBody {
	s.RequestId = &v
	return s
}

func (s *InvokePageResponseBody) SetSuccess(v bool) *InvokePageResponseBody {
	s.Success = &v
	return s
}

func (s *InvokePageResponseBody) Validate() error {
	return dara.Validate(s)
}
