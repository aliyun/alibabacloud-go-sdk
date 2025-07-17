// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIntervalLimitOfSLAResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetIntervalLimitOfSLAResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetIntervalLimitOfSLAResponseBody
	GetErrorMessage() *string
	SetIntervalLimit(v int32) *GetIntervalLimitOfSLAResponseBody
	GetIntervalLimit() *int32
	SetRequestId(v string) *GetIntervalLimitOfSLAResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetIntervalLimitOfSLAResponseBody
	GetSuccess() *bool
}

type GetIntervalLimitOfSLAResponseBody struct {
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
	// The minimum scheduling cycle. Unit: minutes.
	//
	// example:
	//
	// 59
	IntervalLimit *int32 `json:"IntervalLimit,omitempty" xml:"IntervalLimit,omitempty"`
	// The ID of the request. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// 5B96E35F-A58E-5399-9041-09CF9A1E46EA
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

func (s GetIntervalLimitOfSLAResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIntervalLimitOfSLAResponseBody) GoString() string {
	return s.String()
}

func (s *GetIntervalLimitOfSLAResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetIntervalLimitOfSLAResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetIntervalLimitOfSLAResponseBody) GetIntervalLimit() *int32 {
	return s.IntervalLimit
}

func (s *GetIntervalLimitOfSLAResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIntervalLimitOfSLAResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetIntervalLimitOfSLAResponseBody) SetErrorCode(v string) *GetIntervalLimitOfSLAResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetIntervalLimitOfSLAResponseBody) SetErrorMessage(v string) *GetIntervalLimitOfSLAResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetIntervalLimitOfSLAResponseBody) SetIntervalLimit(v int32) *GetIntervalLimitOfSLAResponseBody {
	s.IntervalLimit = &v
	return s
}

func (s *GetIntervalLimitOfSLAResponseBody) SetRequestId(v string) *GetIntervalLimitOfSLAResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIntervalLimitOfSLAResponseBody) SetSuccess(v bool) *GetIntervalLimitOfSLAResponseBody {
	s.Success = &v
	return s
}

func (s *GetIntervalLimitOfSLAResponseBody) Validate() error {
	return dara.Validate(s)
}
