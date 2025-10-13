// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceNetworkTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *UpdateInstanceNetworkTypeResponseBody
	GetData() *string
	SetErrorCode(v string) *UpdateInstanceNetworkTypeResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateInstanceNetworkTypeResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *UpdateInstanceNetworkTypeResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *UpdateInstanceNetworkTypeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateInstanceNetworkTypeResponseBody
	GetSuccess() *bool
}

type UpdateInstanceNetworkTypeResponseBody struct {
	// The returned result, which indicates whether the operation was successful.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// example:
	//
	// 404
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Internal server error.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9CC37B9F-F4B4-5FF1-939B-AEE78DC70130
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateInstanceNetworkTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceNetworkTypeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceNetworkTypeResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateInstanceNetworkTypeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateInstanceNetworkTypeResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateInstanceNetworkTypeResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *UpdateInstanceNetworkTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateInstanceNetworkTypeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateInstanceNetworkTypeResponseBody) SetData(v string) *UpdateInstanceNetworkTypeResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateInstanceNetworkTypeResponseBody) SetErrorCode(v string) *UpdateInstanceNetworkTypeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateInstanceNetworkTypeResponseBody) SetErrorMessage(v string) *UpdateInstanceNetworkTypeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateInstanceNetworkTypeResponseBody) SetHttpStatusCode(v string) *UpdateInstanceNetworkTypeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateInstanceNetworkTypeResponseBody) SetRequestId(v string) *UpdateInstanceNetworkTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceNetworkTypeResponseBody) SetSuccess(v bool) *UpdateInstanceNetworkTypeResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateInstanceNetworkTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
