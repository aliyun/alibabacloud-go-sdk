// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDisasterRecoveryItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *StopDisasterRecoveryItemResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *StopDisasterRecoveryItemResponseBody
	GetCode() *string
	SetData(v bool) *StopDisasterRecoveryItemResponseBody
	GetData() *bool
	SetDynamicCode(v string) *StopDisasterRecoveryItemResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *StopDisasterRecoveryItemResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *StopDisasterRecoveryItemResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *StopDisasterRecoveryItemResponseBody
	GetMessage() *string
	SetRequestId(v string) *StopDisasterRecoveryItemResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StopDisasterRecoveryItemResponseBody
	GetSuccess() *bool
}

type StopDisasterRecoveryItemResponseBody struct {
	// The details about the access denial. This parameter is returned only if the access is denied because the Resource Access Management (RAM) user does not have the required permissions.
	//
	// example:
	//
	// xxx
	AccessDeniedDetail *string `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty"`
	// Error code
	//
	// example:
	//
	// DisasterRecoveryItemStatus.Error
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Return result
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// Dynamic error code
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// Dynamic error message
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Error message
	//
	// example:
	//
	// The current status of the disaster recovery item does not support this operation.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID
	//
	// example:
	//
	// A07B41BD-6DD3-5349-9E76-00303xxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Whether the operation was successful
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s StopDisasterRecoveryItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopDisasterRecoveryItemResponseBody) GoString() string {
	return s.String()
}

func (s *StopDisasterRecoveryItemResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *StopDisasterRecoveryItemResponseBody) GetCode() *string {
	return s.Code
}

func (s *StopDisasterRecoveryItemResponseBody) GetData() *bool {
	return s.Data
}

func (s *StopDisasterRecoveryItemResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *StopDisasterRecoveryItemResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *StopDisasterRecoveryItemResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *StopDisasterRecoveryItemResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StopDisasterRecoveryItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopDisasterRecoveryItemResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StopDisasterRecoveryItemResponseBody) SetAccessDeniedDetail(v string) *StopDisasterRecoveryItemResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *StopDisasterRecoveryItemResponseBody) SetCode(v string) *StopDisasterRecoveryItemResponseBody {
	s.Code = &v
	return s
}

func (s *StopDisasterRecoveryItemResponseBody) SetData(v bool) *StopDisasterRecoveryItemResponseBody {
	s.Data = &v
	return s
}

func (s *StopDisasterRecoveryItemResponseBody) SetDynamicCode(v string) *StopDisasterRecoveryItemResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *StopDisasterRecoveryItemResponseBody) SetDynamicMessage(v string) *StopDisasterRecoveryItemResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *StopDisasterRecoveryItemResponseBody) SetHttpStatusCode(v int32) *StopDisasterRecoveryItemResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StopDisasterRecoveryItemResponseBody) SetMessage(v string) *StopDisasterRecoveryItemResponseBody {
	s.Message = &v
	return s
}

func (s *StopDisasterRecoveryItemResponseBody) SetRequestId(v string) *StopDisasterRecoveryItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopDisasterRecoveryItemResponseBody) SetSuccess(v bool) *StopDisasterRecoveryItemResponseBody {
	s.Success = &v
	return s
}

func (s *StopDisasterRecoveryItemResponseBody) Validate() error {
	return dara.Validate(s)
}
