// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScenarioResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteScenarioResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteScenarioResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteScenarioResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteScenarioResponseBody
	GetSuccess() *bool
}

type DeleteScenarioResponseBody struct {
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
	// 12***
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

func (s DeleteScenarioResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteScenarioResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScenarioResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteScenarioResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteScenarioResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteScenarioResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteScenarioResponseBody) SetErrorCode(v string) *DeleteScenarioResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteScenarioResponseBody) SetErrorMessage(v string) *DeleteScenarioResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteScenarioResponseBody) SetRequestId(v string) *DeleteScenarioResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteScenarioResponseBody) SetSuccess(v bool) *DeleteScenarioResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteScenarioResponseBody) Validate() error {
	return dara.Validate(s)
}
