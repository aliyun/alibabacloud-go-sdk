// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSLARulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateSLARulesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateSLARulesResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateSLARulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateSLARulesResponseBody
	GetSuccess() *bool
}

type UpdateSLARulesResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request. You can use the ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 88E053F7-347B-52DD-A186-1F340EEC0C27
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateSLARulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSLARulesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSLARulesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateSLARulesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateSLARulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSLARulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateSLARulesResponseBody) SetErrorCode(v string) *UpdateSLARulesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateSLARulesResponseBody) SetErrorMessage(v string) *UpdateSLARulesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateSLARulesResponseBody) SetRequestId(v string) *UpdateSLARulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSLARulesResponseBody) SetSuccess(v bool) *UpdateSLARulesResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateSLARulesResponseBody) Validate() error {
	return dara.Validate(s)
}
