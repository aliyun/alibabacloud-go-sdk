// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDisasterRecoveryPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteDisasterRecoveryPlanResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DeleteDisasterRecoveryPlanResponseBody
	GetCode() *string
	SetData(v bool) *DeleteDisasterRecoveryPlanResponseBody
	GetData() *bool
	SetDynamicCode(v string) *DeleteDisasterRecoveryPlanResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DeleteDisasterRecoveryPlanResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *DeleteDisasterRecoveryPlanResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteDisasterRecoveryPlanResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteDisasterRecoveryPlanResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDisasterRecoveryPlanResponseBody
	GetSuccess() *bool
}

type DeleteDisasterRecoveryPlanResponseBody struct {
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
	// AF9A8B10-C426-530F-A0DD-96320B39****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteDisasterRecoveryPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDisasterRecoveryPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDisasterRecoveryPlanResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteDisasterRecoveryPlanResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteDisasterRecoveryPlanResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteDisasterRecoveryPlanResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DeleteDisasterRecoveryPlanResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DeleteDisasterRecoveryPlanResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteDisasterRecoveryPlanResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteDisasterRecoveryPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDisasterRecoveryPlanResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDisasterRecoveryPlanResponseBody) SetAccessDeniedDetail(v string) *DeleteDisasterRecoveryPlanResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteDisasterRecoveryPlanResponseBody) SetCode(v string) *DeleteDisasterRecoveryPlanResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDisasterRecoveryPlanResponseBody) SetData(v bool) *DeleteDisasterRecoveryPlanResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteDisasterRecoveryPlanResponseBody) SetDynamicCode(v string) *DeleteDisasterRecoveryPlanResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteDisasterRecoveryPlanResponseBody) SetDynamicMessage(v string) *DeleteDisasterRecoveryPlanResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteDisasterRecoveryPlanResponseBody) SetHttpStatusCode(v int32) *DeleteDisasterRecoveryPlanResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteDisasterRecoveryPlanResponseBody) SetMessage(v string) *DeleteDisasterRecoveryPlanResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDisasterRecoveryPlanResponseBody) SetRequestId(v string) *DeleteDisasterRecoveryPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDisasterRecoveryPlanResponseBody) SetSuccess(v bool) *DeleteDisasterRecoveryPlanResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDisasterRecoveryPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
