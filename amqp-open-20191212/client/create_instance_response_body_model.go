// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateInstanceResponseBody
	GetCode() *int32
	SetData(v interface{}) *CreateInstanceResponseBody
	GetData() interface{}
	SetMessage(v string) *CreateInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateInstanceResponseBody
	GetSuccess() *bool
}

type CreateInstanceResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the created instance.
	//
	// example:
	//
	// amqp-cn-xxxxx
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// xxx failed,xxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID. You can use the ID to troubleshoot issues. This parameter is a common parameter.
	//
	// example:
	//
	// CCBB1225-C392-480E-8C7F-D09AB2CD2***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateInstanceResponseBody) GetData() interface{} {
	return s.Data
}

func (s *CreateInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateInstanceResponseBody) SetCode(v int32) *CreateInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateInstanceResponseBody) SetData(v interface{}) *CreateInstanceResponseBody {
	s.Data = v
	return s
}

func (s *CreateInstanceResponseBody) SetMessage(v string) *CreateInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateInstanceResponseBody) SetRequestId(v string) *CreateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetSuccess(v bool) *CreateInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
