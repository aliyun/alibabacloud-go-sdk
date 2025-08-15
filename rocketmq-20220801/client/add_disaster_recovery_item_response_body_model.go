// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDisasterRecoveryItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *AddDisasterRecoveryItemResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *AddDisasterRecoveryItemResponseBody
	GetCode() *string
	SetData(v int64) *AddDisasterRecoveryItemResponseBody
	GetData() *int64
	SetDynamicCode(v string) *AddDisasterRecoveryItemResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *AddDisasterRecoveryItemResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *AddDisasterRecoveryItemResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddDisasterRecoveryItemResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddDisasterRecoveryItemResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddDisasterRecoveryItemResponseBody
	GetSuccess() *bool
}

type AddDisasterRecoveryItemResponseBody struct {
	// Access denied details, only in the scenario where the user is denied access due to RAM not having permission
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
	// Return result, mapping task ID
	//
	// example:
	//
	// 1300000016
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
	// AF9A8B10-C426-530F-A0DD-96320B39****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Whether the operation was successful
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddDisasterRecoveryItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddDisasterRecoveryItemResponseBody) GoString() string {
	return s.String()
}

func (s *AddDisasterRecoveryItemResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *AddDisasterRecoveryItemResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddDisasterRecoveryItemResponseBody) GetData() *int64 {
	return s.Data
}

func (s *AddDisasterRecoveryItemResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *AddDisasterRecoveryItemResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *AddDisasterRecoveryItemResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddDisasterRecoveryItemResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddDisasterRecoveryItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddDisasterRecoveryItemResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddDisasterRecoveryItemResponseBody) SetAccessDeniedDetail(v string) *AddDisasterRecoveryItemResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *AddDisasterRecoveryItemResponseBody) SetCode(v string) *AddDisasterRecoveryItemResponseBody {
	s.Code = &v
	return s
}

func (s *AddDisasterRecoveryItemResponseBody) SetData(v int64) *AddDisasterRecoveryItemResponseBody {
	s.Data = &v
	return s
}

func (s *AddDisasterRecoveryItemResponseBody) SetDynamicCode(v string) *AddDisasterRecoveryItemResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *AddDisasterRecoveryItemResponseBody) SetDynamicMessage(v string) *AddDisasterRecoveryItemResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *AddDisasterRecoveryItemResponseBody) SetHttpStatusCode(v int32) *AddDisasterRecoveryItemResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddDisasterRecoveryItemResponseBody) SetMessage(v string) *AddDisasterRecoveryItemResponseBody {
	s.Message = &v
	return s
}

func (s *AddDisasterRecoveryItemResponseBody) SetRequestId(v string) *AddDisasterRecoveryItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDisasterRecoveryItemResponseBody) SetSuccess(v bool) *AddDisasterRecoveryItemResponseBody {
	s.Success = &v
	return s
}

func (s *AddDisasterRecoveryItemResponseBody) Validate() error {
	return dara.Validate(s)
}
