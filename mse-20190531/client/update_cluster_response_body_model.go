// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateClusterResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v string) *UpdateClusterResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *UpdateClusterResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateClusterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateClusterResponseBody
	GetSuccess() *bool
}

type UpdateClusterResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// mse-100-100
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// The request is successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 5B170A0D-2C5D-4CF8-B808-03966B86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateClusterResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateClusterResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateClusterResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *UpdateClusterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateClusterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateClusterResponseBody) SetErrorCode(v string) *UpdateClusterResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateClusterResponseBody) SetHttpStatusCode(v string) *UpdateClusterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateClusterResponseBody) SetMessage(v string) *UpdateClusterResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateClusterResponseBody) SetRequestId(v string) *UpdateClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateClusterResponseBody) SetSuccess(v bool) *UpdateClusterResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
