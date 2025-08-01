// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAclResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateAclResponseBody
	GetErrorCode() *string
	SetMessage(v string) *UpdateAclResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAclResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAclResponseBody
	GetSuccess() *bool
}

type UpdateAclResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// mse-100-100
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
	// 7466566F-F30F-4A29-965D-3E0AF21D****
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

func (s UpdateAclResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAclResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAclResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateAclResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAclResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAclResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAclResponseBody) SetErrorCode(v string) *UpdateAclResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateAclResponseBody) SetMessage(v string) *UpdateAclResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAclResponseBody) SetRequestId(v string) *UpdateAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAclResponseBody) SetSuccess(v bool) *UpdateAclResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAclResponseBody) Validate() error {
	return dara.Validate(s)
}
