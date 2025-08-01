// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateZnodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateZnodeResponseBody
	GetErrorCode() *string
	SetMessage(v string) *UpdateZnodeResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateZnodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateZnodeResponseBody
	GetSuccess() *bool
}

type UpdateZnodeResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
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
	// 316F5F64-F73D-42DC-8632-01E308B6****
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

func (s UpdateZnodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateZnodeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateZnodeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateZnodeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateZnodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateZnodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateZnodeResponseBody) SetErrorCode(v string) *UpdateZnodeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateZnodeResponseBody) SetMessage(v string) *UpdateZnodeResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateZnodeResponseBody) SetRequestId(v string) *UpdateZnodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateZnodeResponseBody) SetSuccess(v bool) *UpdateZnodeResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateZnodeResponseBody) Validate() error {
	return dara.Validate(s)
}
