// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComputeQuotaPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *CreateComputeQuotaPlanResponseBody
	GetData() *string
	SetErrorCode(v string) *CreateComputeQuotaPlanResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *CreateComputeQuotaPlanResponseBody
	GetErrorMsg() *string
	SetHttpCode(v int32) *CreateComputeQuotaPlanResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *CreateComputeQuotaPlanResponseBody
	GetRequestId() *string
}

type CreateComputeQuotaPlanResponseBody struct {
	// The returned data.
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
	// 0b87b7e716665825896565060e87a4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateComputeQuotaPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateComputeQuotaPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateComputeQuotaPlanResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateComputeQuotaPlanResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateComputeQuotaPlanResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *CreateComputeQuotaPlanResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *CreateComputeQuotaPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateComputeQuotaPlanResponseBody) SetData(v string) *CreateComputeQuotaPlanResponseBody {
	s.Data = &v
	return s
}

func (s *CreateComputeQuotaPlanResponseBody) SetErrorCode(v string) *CreateComputeQuotaPlanResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateComputeQuotaPlanResponseBody) SetErrorMsg(v string) *CreateComputeQuotaPlanResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateComputeQuotaPlanResponseBody) SetHttpCode(v int32) *CreateComputeQuotaPlanResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateComputeQuotaPlanResponseBody) SetRequestId(v string) *CreateComputeQuotaPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateComputeQuotaPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
