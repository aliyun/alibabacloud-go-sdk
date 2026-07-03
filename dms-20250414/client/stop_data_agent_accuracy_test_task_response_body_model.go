// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDataAgentAccuracyTestTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *StopDataAgentAccuracyTestTaskResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *StopDataAgentAccuracyTestTaskResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *StopDataAgentAccuracyTestTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StopDataAgentAccuracyTestTaskResponseBody
	GetSuccess() *bool
}

type StopDataAgentAccuracyTestTaskResponseBody struct {
	// The error code.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the call failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E0D21075-xxx-FD8AD04A63B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopDataAgentAccuracyTestTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopDataAgentAccuracyTestTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StopDataAgentAccuracyTestTaskResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *StopDataAgentAccuracyTestTaskResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *StopDataAgentAccuracyTestTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopDataAgentAccuracyTestTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StopDataAgentAccuracyTestTaskResponseBody) SetErrorCode(v string) *StopDataAgentAccuracyTestTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *StopDataAgentAccuracyTestTaskResponseBody) SetErrorMessage(v string) *StopDataAgentAccuracyTestTaskResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *StopDataAgentAccuracyTestTaskResponseBody) SetRequestId(v string) *StopDataAgentAccuracyTestTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopDataAgentAccuracyTestTaskResponseBody) SetSuccess(v bool) *StopDataAgentAccuracyTestTaskResponseBody {
	s.Success = &v
	return s
}

func (s *StopDataAgentAccuracyTestTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
