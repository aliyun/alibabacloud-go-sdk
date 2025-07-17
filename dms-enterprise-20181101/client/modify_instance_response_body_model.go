// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ModifyInstanceResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ModifyInstanceResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ModifyInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyInstanceResponseBody
	GetSuccess() *bool
}

type ModifyInstanceResponseBody struct {
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ModifyInstanceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ModifyInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyInstanceResponseBody) SetErrorCode(v string) *ModifyInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ModifyInstanceResponseBody) SetErrorMessage(v string) *ModifyInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ModifyInstanceResponseBody) SetRequestId(v string) *ModifyInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceResponseBody) SetSuccess(v bool) *ModifyInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
