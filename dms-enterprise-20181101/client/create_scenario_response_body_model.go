// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScenarioResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateScenarioResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateScenarioResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateScenarioResponseBody
	GetRequestId() *string
	SetScenarioId(v int64) *CreateScenarioResponseBody
	GetScenarioId() *int64
	SetSuccess(v bool) *CreateScenarioResponseBody
	GetSuccess() *bool
}

type CreateScenarioResponseBody struct {
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
	// 4E1D2B4D-3E53-4ABC-999D-1D2520B3471A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the business scenario.
	//
	// example:
	//
	// 36***
	ScenarioId *int64 `json:"ScenarioId,omitempty" xml:"ScenarioId,omitempty"`
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

func (s CreateScenarioResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateScenarioResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScenarioResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateScenarioResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateScenarioResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateScenarioResponseBody) GetScenarioId() *int64 {
	return s.ScenarioId
}

func (s *CreateScenarioResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateScenarioResponseBody) SetErrorCode(v string) *CreateScenarioResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateScenarioResponseBody) SetErrorMessage(v string) *CreateScenarioResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateScenarioResponseBody) SetRequestId(v string) *CreateScenarioResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScenarioResponseBody) SetScenarioId(v int64) *CreateScenarioResponseBody {
	s.ScenarioId = &v
	return s
}

func (s *CreateScenarioResponseBody) SetSuccess(v bool) *CreateScenarioResponseBody {
	s.Success = &v
	return s
}

func (s *CreateScenarioResponseBody) Validate() error {
	return dara.Validate(s)
}
