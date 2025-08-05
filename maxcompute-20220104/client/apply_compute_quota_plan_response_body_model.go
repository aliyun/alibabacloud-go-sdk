// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyComputeQuotaPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *ApplyComputeQuotaPlanResponseBody
	GetData() *string
	SetErrorCode(v string) *ApplyComputeQuotaPlanResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *ApplyComputeQuotaPlanResponseBody
	GetErrorMsg() *string
	SetHttpCode(v int32) *ApplyComputeQuotaPlanResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *ApplyComputeQuotaPlanResponseBody
	GetRequestId() *string
}

type ApplyComputeQuotaPlanResponseBody struct {
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
	// OBJECT_NOT_EXIST
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// this quota plan is not exist.
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
	// 0b87b7a316654730544735643e9200
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ApplyComputeQuotaPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyComputeQuotaPlanResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyComputeQuotaPlanResponseBody) GetData() *string {
	return s.Data
}

func (s *ApplyComputeQuotaPlanResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ApplyComputeQuotaPlanResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ApplyComputeQuotaPlanResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *ApplyComputeQuotaPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyComputeQuotaPlanResponseBody) SetData(v string) *ApplyComputeQuotaPlanResponseBody {
	s.Data = &v
	return s
}

func (s *ApplyComputeQuotaPlanResponseBody) SetErrorCode(v string) *ApplyComputeQuotaPlanResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ApplyComputeQuotaPlanResponseBody) SetErrorMsg(v string) *ApplyComputeQuotaPlanResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ApplyComputeQuotaPlanResponseBody) SetHttpCode(v int32) *ApplyComputeQuotaPlanResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ApplyComputeQuotaPlanResponseBody) SetRequestId(v string) *ApplyComputeQuotaPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyComputeQuotaPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
