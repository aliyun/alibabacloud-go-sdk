// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComputeSubQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *UpdateComputeSubQuotaResponseBody
	GetData() *string
	SetErrorCode(v string) *UpdateComputeSubQuotaResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *UpdateComputeSubQuotaResponseBody
	GetErrorMsg() *string
	SetHttpCode(v int32) *UpdateComputeSubQuotaResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *UpdateComputeSubQuotaResponseBody
	GetRequestId() *string
}

type UpdateComputeSubQuotaResponseBody struct {
	// The returned result.
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
	// this quota is not exist.
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
	// The ID of the request.
	//
	// example:
	//
	// 0b57ff7616612271051086500ea3ce
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateComputeSubQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeSubQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateComputeSubQuotaResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateComputeSubQuotaResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateComputeSubQuotaResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *UpdateComputeSubQuotaResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *UpdateComputeSubQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateComputeSubQuotaResponseBody) SetData(v string) *UpdateComputeSubQuotaResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateComputeSubQuotaResponseBody) SetErrorCode(v string) *UpdateComputeSubQuotaResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateComputeSubQuotaResponseBody) SetErrorMsg(v string) *UpdateComputeSubQuotaResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateComputeSubQuotaResponseBody) SetHttpCode(v int32) *UpdateComputeSubQuotaResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateComputeSubQuotaResponseBody) SetRequestId(v string) *UpdateComputeSubQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateComputeSubQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}
