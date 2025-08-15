// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateInstanceResponseBody
	GetCode() *string
	SetData(v string) *CreateInstanceResponseBody
	GetData() *string
	SetDynamicCode(v string) *CreateInstanceResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *CreateInstanceResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *CreateInstanceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateInstanceResponseBody
	GetSuccess() *bool
}

type CreateInstanceResponseBody struct {
	// The error code returned if the call failed.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The ID of the created instance.
	//
	// example:
	//
	// rmq-cn-7e22ody****
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The ID of the request. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// AF9A8B10-C426-530F-A0DD-96320B39****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateInstanceResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateInstanceResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *CreateInstanceResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *CreateInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
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

func (s *CreateInstanceResponseBody) SetCode(v string) *CreateInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateInstanceResponseBody) SetData(v string) *CreateInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *CreateInstanceResponseBody) SetDynamicCode(v string) *CreateInstanceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateInstanceResponseBody) SetDynamicMessage(v string) *CreateInstanceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateInstanceResponseBody) SetHttpStatusCode(v int32) *CreateInstanceResponseBody {
	s.HttpStatusCode = &v
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
