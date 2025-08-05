// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComputeQuotaPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *UpdateComputeQuotaPlanResponseBody
	GetData() *string
	SetErrorCode(v string) *UpdateComputeQuotaPlanResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *UpdateComputeQuotaPlanResponseBody
	GetErrorMsg() *string
	SetHttpCode(v int32) *UpdateComputeQuotaPlanResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *UpdateComputeQuotaPlanResponseBody
	GetRequestId() *string
}

type UpdateComputeQuotaPlanResponseBody struct {
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
	// The HTTP status code.
	//
	// - 1xx: informational response. The request is received and is being processed.
	//
	// - 2xx: success. The request is successfully received, understood, and accepted by the server.
	//
	// - 3xx: redirection. The request is redirected, and further actions are required to complete the request.
	//
	// - 4xx: client error. The request contains invalid request parameters or syntaxes, or specific request conditions cannot be met.
	//
	// - 5xx: server error. The server cannot meet requirements due to other reasons.
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

func (s UpdateComputeQuotaPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeQuotaPlanResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateComputeQuotaPlanResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateComputeQuotaPlanResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateComputeQuotaPlanResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *UpdateComputeQuotaPlanResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *UpdateComputeQuotaPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateComputeQuotaPlanResponseBody) SetData(v string) *UpdateComputeQuotaPlanResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateComputeQuotaPlanResponseBody) SetErrorCode(v string) *UpdateComputeQuotaPlanResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateComputeQuotaPlanResponseBody) SetErrorMsg(v string) *UpdateComputeQuotaPlanResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateComputeQuotaPlanResponseBody) SetHttpCode(v int32) *UpdateComputeQuotaPlanResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateComputeQuotaPlanResponseBody) SetRequestId(v string) *UpdateComputeQuotaPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateComputeQuotaPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
