// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSQLRateLimitingRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *AddSQLRateLimitingRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddSQLRateLimitingRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddSQLRateLimitingRulesResponseBody
	GetSuccess() *bool
}

type AddSQLRateLimitingRulesResponseBody struct {
	// example:
	//
	// Message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 69A85BAF-1089-4CDF-A82F-0A140F******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddSQLRateLimitingRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddSQLRateLimitingRulesResponseBody) GoString() string {
	return s.String()
}

func (s *AddSQLRateLimitingRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddSQLRateLimitingRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddSQLRateLimitingRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddSQLRateLimitingRulesResponseBody) SetMessage(v string) *AddSQLRateLimitingRulesResponseBody {
	s.Message = &v
	return s
}

func (s *AddSQLRateLimitingRulesResponseBody) SetRequestId(v string) *AddSQLRateLimitingRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddSQLRateLimitingRulesResponseBody) SetSuccess(v bool) *AddSQLRateLimitingRulesResponseBody {
	s.Success = &v
	return s
}

func (s *AddSQLRateLimitingRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
