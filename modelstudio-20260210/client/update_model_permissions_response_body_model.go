// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelPermissionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateModelPermissionsResponseBody
	GetCode() *string
	SetErrorMessage(v string) *UpdateModelPermissionsResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int64) *UpdateModelPermissionsResponseBody
	GetHttpStatusCode() *int64
	SetRequestId(v string) *UpdateModelPermissionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateModelPermissionsResponseBody
	GetSuccess() *bool
}

type UpdateModelPermissionsResponseBody struct {
	// The error code. This parameter is empty when the call is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The error message. This parameter is empty when the call is successful.
	//
	// example:
	//
	// The specified parameter is invalid
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The unique request ID.
	//
	// example:
	//
	// 36045E0A-551D-592D-B1BC-4C56596CE59E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the API call is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateModelPermissionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateModelPermissionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateModelPermissionsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateModelPermissionsResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *UpdateModelPermissionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateModelPermissionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateModelPermissionsResponseBody) SetCode(v string) *UpdateModelPermissionsResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateModelPermissionsResponseBody) SetErrorMessage(v string) *UpdateModelPermissionsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateModelPermissionsResponseBody) SetHttpStatusCode(v int64) *UpdateModelPermissionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateModelPermissionsResponseBody) SetRequestId(v string) *UpdateModelPermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateModelPermissionsResponseBody) SetSuccess(v bool) *UpdateModelPermissionsResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateModelPermissionsResponseBody) Validate() error {
	return dara.Validate(s)
}
