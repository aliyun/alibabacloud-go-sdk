// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDesensitizationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *AddDesensitizationRuleResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *AddDesensitizationRuleResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *AddDesensitizationRuleResponseBody
	GetRequestId() *string
	SetRuleId(v int32) *AddDesensitizationRuleResponseBody
	GetRuleId() *int32
	SetSuccess(v bool) *AddDesensitizationRuleResponseBody
	GetSuccess() *bool
}

type AddDesensitizationRuleResponseBody struct {
	// The error code that is returned.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message that is returned.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID generated for the data masking rule.
	//
	// example:
	//
	// 35***
	RuleId *int32 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddDesensitizationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddDesensitizationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *AddDesensitizationRuleResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AddDesensitizationRuleResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *AddDesensitizationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddDesensitizationRuleResponseBody) GetRuleId() *int32 {
	return s.RuleId
}

func (s *AddDesensitizationRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddDesensitizationRuleResponseBody) SetErrorCode(v string) *AddDesensitizationRuleResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AddDesensitizationRuleResponseBody) SetErrorMessage(v string) *AddDesensitizationRuleResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddDesensitizationRuleResponseBody) SetRequestId(v string) *AddDesensitizationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDesensitizationRuleResponseBody) SetRuleId(v int32) *AddDesensitizationRuleResponseBody {
	s.RuleId = &v
	return s
}

func (s *AddDesensitizationRuleResponseBody) SetSuccess(v bool) *AddDesensitizationRuleResponseBody {
	s.Success = &v
	return s
}

func (s *AddDesensitizationRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
