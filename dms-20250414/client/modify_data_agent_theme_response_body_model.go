// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDataAgentThemeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ModifyDataAgentThemeResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ModifyDataAgentThemeResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ModifyDataAgentThemeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyDataAgentThemeResponseBody
	GetSuccess() *bool
}

type ModifyDataAgentThemeResponseBody struct {
	// The error code returned when the request is abnormal.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the call fails.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID, which is used to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyDataAgentThemeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDataAgentThemeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDataAgentThemeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ModifyDataAgentThemeResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ModifyDataAgentThemeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDataAgentThemeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyDataAgentThemeResponseBody) SetErrorCode(v string) *ModifyDataAgentThemeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ModifyDataAgentThemeResponseBody) SetErrorMessage(v string) *ModifyDataAgentThemeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ModifyDataAgentThemeResponseBody) SetRequestId(v string) *ModifyDataAgentThemeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDataAgentThemeResponseBody) SetSuccess(v bool) *ModifyDataAgentThemeResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyDataAgentThemeResponseBody) Validate() error {
	return dara.Validate(s)
}
