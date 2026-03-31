// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComputeQuotaScheduleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *UpdateComputeQuotaScheduleResponseBody
	GetData() *string
	SetErrorCode(v string) *UpdateComputeQuotaScheduleResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *UpdateComputeQuotaScheduleResponseBody
	GetErrorMsg() *string
	SetHttpCode(v int32) *UpdateComputeQuotaScheduleResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *UpdateComputeQuotaScheduleResponseBody
	GetRequestId() *string
}

type UpdateComputeQuotaScheduleResponseBody struct {
	// The data returned.
	//
	// example:
	//
	// success
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The error code.
	//
	// example:
	//
	// QUOTA_PLAN_NOT_FOUND
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// plan \\"***\\" does not exist
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// HTTP status code.
	//
	// - 1xx: Informational - The request has been received and is being processed.
	//
	// - 2xx: Success - The request action was successfully received, understood, and accepted by the server.
	//
	// - 3xx: Redirection - Further action must be taken to complete the request.
	//
	// - 4xx: Client Error - The request contains an error in the request parameters, syntax, or specific request conditions cannot be met.
	//
	// - 5xx: Server Error - The server could not fulfill the request due to other reasons.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0be3e0aa16667684362147582e038f
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateComputeQuotaScheduleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeQuotaScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateComputeQuotaScheduleResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateComputeQuotaScheduleResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateComputeQuotaScheduleResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *UpdateComputeQuotaScheduleResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *UpdateComputeQuotaScheduleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateComputeQuotaScheduleResponseBody) SetData(v string) *UpdateComputeQuotaScheduleResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateComputeQuotaScheduleResponseBody) SetErrorCode(v string) *UpdateComputeQuotaScheduleResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateComputeQuotaScheduleResponseBody) SetErrorMsg(v string) *UpdateComputeQuotaScheduleResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateComputeQuotaScheduleResponseBody) SetHttpCode(v int32) *UpdateComputeQuotaScheduleResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateComputeQuotaScheduleResponseBody) SetRequestId(v string) *UpdateComputeQuotaScheduleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateComputeQuotaScheduleResponseBody) Validate() error {
	return dara.Validate(s)
}
