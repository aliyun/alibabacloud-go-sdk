// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDisasterRecoveryItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteDisasterRecoveryItemResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DeleteDisasterRecoveryItemResponseBody
	GetCode() *string
	SetData(v bool) *DeleteDisasterRecoveryItemResponseBody
	GetData() *bool
	SetDynamicCode(v string) *DeleteDisasterRecoveryItemResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DeleteDisasterRecoveryItemResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *DeleteDisasterRecoveryItemResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteDisasterRecoveryItemResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteDisasterRecoveryItemResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDisasterRecoveryItemResponseBody
	GetSuccess() *bool
}

type DeleteDisasterRecoveryItemResponseBody struct {
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
	// DisasterRecoveryItemStatus.Error
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The return data
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
	// 0B962390-D84B-5D44-8C11-79DFxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Whether the operation was successful
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteDisasterRecoveryItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDisasterRecoveryItemResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDisasterRecoveryItemResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteDisasterRecoveryItemResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteDisasterRecoveryItemResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteDisasterRecoveryItemResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DeleteDisasterRecoveryItemResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DeleteDisasterRecoveryItemResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteDisasterRecoveryItemResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteDisasterRecoveryItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDisasterRecoveryItemResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDisasterRecoveryItemResponseBody) SetAccessDeniedDetail(v string) *DeleteDisasterRecoveryItemResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteDisasterRecoveryItemResponseBody) SetCode(v string) *DeleteDisasterRecoveryItemResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDisasterRecoveryItemResponseBody) SetData(v bool) *DeleteDisasterRecoveryItemResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteDisasterRecoveryItemResponseBody) SetDynamicCode(v string) *DeleteDisasterRecoveryItemResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteDisasterRecoveryItemResponseBody) SetDynamicMessage(v string) *DeleteDisasterRecoveryItemResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteDisasterRecoveryItemResponseBody) SetHttpStatusCode(v int32) *DeleteDisasterRecoveryItemResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteDisasterRecoveryItemResponseBody) SetMessage(v string) *DeleteDisasterRecoveryItemResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDisasterRecoveryItemResponseBody) SetRequestId(v string) *DeleteDisasterRecoveryItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDisasterRecoveryItemResponseBody) SetSuccess(v bool) *DeleteDisasterRecoveryItemResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDisasterRecoveryItemResponseBody) Validate() error {
	return dara.Validate(s)
}
