// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRuleNumLimitOfSLAResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetRuleNumLimitOfSLAResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetRuleNumLimitOfSLAResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetRuleNumLimitOfSLAResponseBody
	GetRequestId() *string
	SetRuleNumLimit(v int32) *GetRuleNumLimitOfSLAResponseBody
	GetRuleNumLimit() *int32
	SetSuccess(v bool) *GetRuleNumLimitOfSLAResponseBody
	GetSuccess() *bool
}

type GetRuleNumLimitOfSLAResponseBody struct {
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
	// D05B3EE1-B6D3-5B17-8CA6-A8054828E5B2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The maximum number of SLA rules.
	//
	// example:
	//
	// 12
	RuleNumLimit *int32 `json:"RuleNumLimit,omitempty" xml:"RuleNumLimit,omitempty"`
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

func (s GetRuleNumLimitOfSLAResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRuleNumLimitOfSLAResponseBody) GoString() string {
	return s.String()
}

func (s *GetRuleNumLimitOfSLAResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetRuleNumLimitOfSLAResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetRuleNumLimitOfSLAResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRuleNumLimitOfSLAResponseBody) GetRuleNumLimit() *int32 {
	return s.RuleNumLimit
}

func (s *GetRuleNumLimitOfSLAResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetRuleNumLimitOfSLAResponseBody) SetErrorCode(v string) *GetRuleNumLimitOfSLAResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetRuleNumLimitOfSLAResponseBody) SetErrorMessage(v string) *GetRuleNumLimitOfSLAResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetRuleNumLimitOfSLAResponseBody) SetRequestId(v string) *GetRuleNumLimitOfSLAResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRuleNumLimitOfSLAResponseBody) SetRuleNumLimit(v int32) *GetRuleNumLimitOfSLAResponseBody {
	s.RuleNumLimit = &v
	return s
}

func (s *GetRuleNumLimitOfSLAResponseBody) SetSuccess(v bool) *GetRuleNumLimitOfSLAResponseBody {
	s.Success = &v
	return s
}

func (s *GetRuleNumLimitOfSLAResponseBody) Validate() error {
	return dara.Validate(s)
}
