// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDisasterRecoveryItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *StartDisasterRecoveryItemResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *StartDisasterRecoveryItemResponseBody
	GetCode() *string
	SetData(v bool) *StartDisasterRecoveryItemResponseBody
	GetData() *bool
	SetDynamicCode(v string) *StartDisasterRecoveryItemResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *StartDisasterRecoveryItemResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *StartDisasterRecoveryItemResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *StartDisasterRecoveryItemResponseBody
	GetMessage() *string
	SetRequestId(v string) *StartDisasterRecoveryItemResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartDisasterRecoveryItemResponseBody
	GetSuccess() *bool
}

type StartDisasterRecoveryItemResponseBody struct {
	// The details about the access denial. This parameter is returned only if the access is denied due to the reason that the Resource Access Management (RAM) user does not have the required permissions.
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
	// InstanceId
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
	// C7E8AE3A-219B-52EE-BE32-4036Fxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Whether the operation was successful
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s StartDisasterRecoveryItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartDisasterRecoveryItemResponseBody) GoString() string {
	return s.String()
}

func (s *StartDisasterRecoveryItemResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *StartDisasterRecoveryItemResponseBody) GetCode() *string {
	return s.Code
}

func (s *StartDisasterRecoveryItemResponseBody) GetData() *bool {
	return s.Data
}

func (s *StartDisasterRecoveryItemResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *StartDisasterRecoveryItemResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *StartDisasterRecoveryItemResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *StartDisasterRecoveryItemResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartDisasterRecoveryItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartDisasterRecoveryItemResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartDisasterRecoveryItemResponseBody) SetAccessDeniedDetail(v string) *StartDisasterRecoveryItemResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *StartDisasterRecoveryItemResponseBody) SetCode(v string) *StartDisasterRecoveryItemResponseBody {
	s.Code = &v
	return s
}

func (s *StartDisasterRecoveryItemResponseBody) SetData(v bool) *StartDisasterRecoveryItemResponseBody {
	s.Data = &v
	return s
}

func (s *StartDisasterRecoveryItemResponseBody) SetDynamicCode(v string) *StartDisasterRecoveryItemResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *StartDisasterRecoveryItemResponseBody) SetDynamicMessage(v string) *StartDisasterRecoveryItemResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *StartDisasterRecoveryItemResponseBody) SetHttpStatusCode(v int32) *StartDisasterRecoveryItemResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StartDisasterRecoveryItemResponseBody) SetMessage(v string) *StartDisasterRecoveryItemResponseBody {
	s.Message = &v
	return s
}

func (s *StartDisasterRecoveryItemResponseBody) SetRequestId(v string) *StartDisasterRecoveryItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartDisasterRecoveryItemResponseBody) SetSuccess(v bool) *StartDisasterRecoveryItemResponseBody {
	s.Success = &v
	return s
}

func (s *StartDisasterRecoveryItemResponseBody) Validate() error {
	return dara.Validate(s)
}
