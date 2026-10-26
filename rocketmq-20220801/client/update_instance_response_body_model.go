// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateInstanceResponseBody
	GetCode() *string
	SetData(v bool) *UpdateInstanceResponseBody
	GetData() *bool
	SetDynamicCode(v string) *UpdateInstanceResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *UpdateInstanceResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *UpdateInstanceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateInstanceResponseBody
	GetSuccess() *bool
}

type UpdateInstanceResponseBody struct {
	// The fault code.
	//
	// example:
	//
	// MissingInstanceId
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned result.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// The dynamic fault code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic fault message.
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The fault message.
	//
	// example:
	//
	// Parameter instanceId is mandatory for this action .
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID. Each request ID is unique and active for troubleshooting and locating issues.
	//
	// example:
	//
	// AA87DE09-DA44-52F4-9515-78B1B607****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the operation is executed successfully.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateInstanceResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateInstanceResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *UpdateInstanceResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *UpdateInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateInstanceResponseBody) SetCode(v string) *UpdateInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetData(v bool) *UpdateInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetDynamicCode(v string) *UpdateInstanceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetDynamicMessage(v string) *UpdateInstanceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetHttpStatusCode(v int32) *UpdateInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetMessage(v string) *UpdateInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetRequestId(v string) *UpdateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetSuccess(v bool) *UpdateInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
