// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScenarioResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateScenarioResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateScenarioResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateScenarioResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateScenarioResponseBody
	GetSuccess() *bool
}

type UpdateScenarioResponseBody struct {
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
	// The ID of the request.
	//
	// example:
	//
	// 7FAD400F-7A5C-4193-8F9A-39D86C4F0231
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

func (s UpdateScenarioResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateScenarioResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateScenarioResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateScenarioResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateScenarioResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateScenarioResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateScenarioResponseBody) SetErrorCode(v string) *UpdateScenarioResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateScenarioResponseBody) SetErrorMessage(v string) *UpdateScenarioResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateScenarioResponseBody) SetRequestId(v string) *UpdateScenarioResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateScenarioResponseBody) SetSuccess(v bool) *UpdateScenarioResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateScenarioResponseBody) Validate() error {
	return dara.Validate(s)
}
