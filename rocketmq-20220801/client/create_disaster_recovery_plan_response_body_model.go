// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDisasterRecoveryPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateDisasterRecoveryPlanResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CreateDisasterRecoveryPlanResponseBody
	GetCode() *string
	SetData(v int64) *CreateDisasterRecoveryPlanResponseBody
	GetData() *int64
	SetDynamicCode(v string) *CreateDisasterRecoveryPlanResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *CreateDisasterRecoveryPlanResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *CreateDisasterRecoveryPlanResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateDisasterRecoveryPlanResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateDisasterRecoveryPlanResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDisasterRecoveryPlanResponseBody
	GetSuccess() *bool
}

type CreateDisasterRecoveryPlanResponseBody struct {
	// Access denied details, provided only in scenarios where access is denied due to lack of RAM permissions
	//
	// example:
	//
	// xxx
	AccessDeniedDetail *string `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty"`
	// Error code
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The result, which is the backup plan ID
	//
	// example:
	//
	// 1234
	Data *int64 `json:"data,omitempty" xml:"data,omitempty"`
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
	// xxx
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID
	//
	// example:
	//
	// C7E8AE3A-219B-52EE-BE32-4036F5xxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the operation was successful
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateDisasterRecoveryPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDisasterRecoveryPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDisasterRecoveryPlanResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateDisasterRecoveryPlanResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateDisasterRecoveryPlanResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateDisasterRecoveryPlanResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *CreateDisasterRecoveryPlanResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *CreateDisasterRecoveryPlanResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateDisasterRecoveryPlanResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateDisasterRecoveryPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDisasterRecoveryPlanResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDisasterRecoveryPlanResponseBody) SetAccessDeniedDetail(v string) *CreateDisasterRecoveryPlanResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateDisasterRecoveryPlanResponseBody) SetCode(v string) *CreateDisasterRecoveryPlanResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDisasterRecoveryPlanResponseBody) SetData(v int64) *CreateDisasterRecoveryPlanResponseBody {
	s.Data = &v
	return s
}

func (s *CreateDisasterRecoveryPlanResponseBody) SetDynamicCode(v string) *CreateDisasterRecoveryPlanResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateDisasterRecoveryPlanResponseBody) SetDynamicMessage(v string) *CreateDisasterRecoveryPlanResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateDisasterRecoveryPlanResponseBody) SetHttpStatusCode(v int32) *CreateDisasterRecoveryPlanResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateDisasterRecoveryPlanResponseBody) SetMessage(v string) *CreateDisasterRecoveryPlanResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDisasterRecoveryPlanResponseBody) SetRequestId(v string) *CreateDisasterRecoveryPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDisasterRecoveryPlanResponseBody) SetSuccess(v bool) *CreateDisasterRecoveryPlanResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDisasterRecoveryPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
