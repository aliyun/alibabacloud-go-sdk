// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDisasterRecoveryItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateDisasterRecoveryItemResponseBody
	GetCode() *string
	SetData(v bool) *UpdateDisasterRecoveryItemResponseBody
	GetData() *bool
	SetDynamicCode(v string) *UpdateDisasterRecoveryItemResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *UpdateDisasterRecoveryItemResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *UpdateDisasterRecoveryItemResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateDisasterRecoveryItemResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateDisasterRecoveryItemResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDisasterRecoveryItemResponseBody
	GetSuccess() *bool
}

type UpdateDisasterRecoveryItemResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Topic.NotFound
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
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
	// The response code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The instance cannot be found.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0C32BED2-FA9F-50AD-9DA7-8B70E26C9D0D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateDisasterRecoveryItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDisasterRecoveryItemResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDisasterRecoveryItemResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateDisasterRecoveryItemResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateDisasterRecoveryItemResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *UpdateDisasterRecoveryItemResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *UpdateDisasterRecoveryItemResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateDisasterRecoveryItemResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateDisasterRecoveryItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDisasterRecoveryItemResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDisasterRecoveryItemResponseBody) SetCode(v string) *UpdateDisasterRecoveryItemResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateDisasterRecoveryItemResponseBody) SetData(v bool) *UpdateDisasterRecoveryItemResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateDisasterRecoveryItemResponseBody) SetDynamicCode(v string) *UpdateDisasterRecoveryItemResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UpdateDisasterRecoveryItemResponseBody) SetDynamicMessage(v string) *UpdateDisasterRecoveryItemResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpdateDisasterRecoveryItemResponseBody) SetHttpStatusCode(v int32) *UpdateDisasterRecoveryItemResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateDisasterRecoveryItemResponseBody) SetMessage(v string) *UpdateDisasterRecoveryItemResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateDisasterRecoveryItemResponseBody) SetRequestId(v string) *UpdateDisasterRecoveryItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDisasterRecoveryItemResponseBody) SetSuccess(v bool) *UpdateDisasterRecoveryItemResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDisasterRecoveryItemResponseBody) Validate() error {
	return dara.Validate(s)
}
