// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDisasterRecoveryPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateDisasterRecoveryPlanResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *UpdateDisasterRecoveryPlanResponseBody
	GetCode() *string
	SetData(v bool) *UpdateDisasterRecoveryPlanResponseBody
	GetData() *bool
	SetDynamicCode(v string) *UpdateDisasterRecoveryPlanResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *UpdateDisasterRecoveryPlanResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *UpdateDisasterRecoveryPlanResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateDisasterRecoveryPlanResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateDisasterRecoveryPlanResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDisasterRecoveryPlanResponseBody
	GetSuccess() *bool
}

type UpdateDisasterRecoveryPlanResponseBody struct {
	// The details about the access denial. This parameter is returned only if the access is denied because the Resource Access Management (RAM) user does not have the required permissions.
	//
	// example:
	//
	// xxx
	AccessDeniedDetail *string `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty"`
	// The error code.
	//
	// example:
	//
	// DisasterRecoveryPlanStatus.Error
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The data returned.
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
	// The current status of the disaster recovery plan does not support this operation.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A07B41BD-6DD3-5349-9E76-00303Dxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateDisasterRecoveryPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDisasterRecoveryPlanResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDisasterRecoveryPlanResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateDisasterRecoveryPlanResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateDisasterRecoveryPlanResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateDisasterRecoveryPlanResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *UpdateDisasterRecoveryPlanResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *UpdateDisasterRecoveryPlanResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateDisasterRecoveryPlanResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateDisasterRecoveryPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDisasterRecoveryPlanResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDisasterRecoveryPlanResponseBody) SetAccessDeniedDetail(v string) *UpdateDisasterRecoveryPlanResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateDisasterRecoveryPlanResponseBody) SetCode(v string) *UpdateDisasterRecoveryPlanResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateDisasterRecoveryPlanResponseBody) SetData(v bool) *UpdateDisasterRecoveryPlanResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateDisasterRecoveryPlanResponseBody) SetDynamicCode(v string) *UpdateDisasterRecoveryPlanResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UpdateDisasterRecoveryPlanResponseBody) SetDynamicMessage(v string) *UpdateDisasterRecoveryPlanResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpdateDisasterRecoveryPlanResponseBody) SetHttpStatusCode(v int32) *UpdateDisasterRecoveryPlanResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateDisasterRecoveryPlanResponseBody) SetMessage(v string) *UpdateDisasterRecoveryPlanResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateDisasterRecoveryPlanResponseBody) SetRequestId(v string) *UpdateDisasterRecoveryPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDisasterRecoveryPlanResponseBody) SetSuccess(v bool) *UpdateDisasterRecoveryPlanResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDisasterRecoveryPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
