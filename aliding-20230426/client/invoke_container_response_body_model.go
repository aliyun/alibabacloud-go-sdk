// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeContainerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *InvokeContainerResponseBody
	GetData() interface{}
	SetErrorCode(v string) *InvokeContainerResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *InvokeContainerResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *InvokeContainerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InvokeContainerResponseBody
	GetSuccess() *bool
}

type InvokeContainerResponseBody struct {
	Data      interface{} `json:"data,omitempty" xml:"data,omitempty"`
	ErrorCode *string     `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string     `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	RequestId *string     `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool       `json:"success,omitempty" xml:"success,omitempty"`
}

func (s InvokeContainerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InvokeContainerResponseBody) GoString() string {
	return s.String()
}

func (s *InvokeContainerResponseBody) GetData() interface{} {
	return s.Data
}

func (s *InvokeContainerResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *InvokeContainerResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *InvokeContainerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InvokeContainerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InvokeContainerResponseBody) SetData(v interface{}) *InvokeContainerResponseBody {
	s.Data = v
	return s
}

func (s *InvokeContainerResponseBody) SetErrorCode(v string) *InvokeContainerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *InvokeContainerResponseBody) SetErrorMsg(v string) *InvokeContainerResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *InvokeContainerResponseBody) SetRequestId(v string) *InvokeContainerResponseBody {
	s.RequestId = &v
	return s
}

func (s *InvokeContainerResponseBody) SetSuccess(v bool) *InvokeContainerResponseBody {
	s.Success = &v
	return s
}

func (s *InvokeContainerResponseBody) Validate() error {
	return dara.Validate(s)
}
