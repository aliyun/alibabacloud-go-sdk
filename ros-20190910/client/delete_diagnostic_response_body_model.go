// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDiagnosticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteDiagnosticResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteDiagnosticResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteDiagnosticResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteDiagnosticResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DeleteDiagnosticResponseBody
	GetSuccess() *string
}

type DeleteDiagnosticResponseBody struct {
	// The error code returned.
	//
	// example:
	//
	// Forbidden
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// You are not authorized to complete this action.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 91EF0DC1-3BB5-555E-AAA1-4F5C640D15DA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDiagnosticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDiagnosticResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDiagnosticResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteDiagnosticResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteDiagnosticResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteDiagnosticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDiagnosticResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DeleteDiagnosticResponseBody) SetCode(v string) *DeleteDiagnosticResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDiagnosticResponseBody) SetHttpStatusCode(v int32) *DeleteDiagnosticResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteDiagnosticResponseBody) SetMessage(v string) *DeleteDiagnosticResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDiagnosticResponseBody) SetRequestId(v string) *DeleteDiagnosticResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDiagnosticResponseBody) SetSuccess(v string) *DeleteDiagnosticResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDiagnosticResponseBody) Validate() error {
	return dara.Validate(s)
}
