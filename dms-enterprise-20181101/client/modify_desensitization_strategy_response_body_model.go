// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDesensitizationStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ModifyDesensitizationStrategyResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ModifyDesensitizationStrategyResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ModifyDesensitizationStrategyResponseBody
	GetRequestId() *string
	SetResult(v bool) *ModifyDesensitizationStrategyResponseBody
	GetResult() *bool
	SetSuccess(v bool) *ModifyDesensitizationStrategyResponseBody
	GetSuccess() *bool
}

type ModifyDesensitizationStrategyResponseBody struct {
	// The status code.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 283C461F-11D8-48AA-B695-DF092DA32AF3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	//
	// example:
	//
	// True
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyDesensitizationStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDesensitizationStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDesensitizationStrategyResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ModifyDesensitizationStrategyResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ModifyDesensitizationStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDesensitizationStrategyResponseBody) GetResult() *bool {
	return s.Result
}

func (s *ModifyDesensitizationStrategyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyDesensitizationStrategyResponseBody) SetErrorCode(v string) *ModifyDesensitizationStrategyResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ModifyDesensitizationStrategyResponseBody) SetErrorMessage(v string) *ModifyDesensitizationStrategyResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ModifyDesensitizationStrategyResponseBody) SetRequestId(v string) *ModifyDesensitizationStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDesensitizationStrategyResponseBody) SetResult(v bool) *ModifyDesensitizationStrategyResponseBody {
	s.Result = &v
	return s
}

func (s *ModifyDesensitizationStrategyResponseBody) SetSuccess(v bool) *ModifyDesensitizationStrategyResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyDesensitizationStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
