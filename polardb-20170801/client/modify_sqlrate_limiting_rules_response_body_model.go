// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySQLRateLimitingRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *ModifySQLRateLimitingRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifySQLRateLimitingRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifySQLRateLimitingRulesResponseBody
	GetSuccess() *bool
}

type ModifySQLRateLimitingRulesResponseBody struct {
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 35D3E3DA-4650-407A-BFF5-59BFF1******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifySQLRateLimitingRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySQLRateLimitingRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySQLRateLimitingRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifySQLRateLimitingRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySQLRateLimitingRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifySQLRateLimitingRulesResponseBody) SetMessage(v string) *ModifySQLRateLimitingRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ModifySQLRateLimitingRulesResponseBody) SetRequestId(v string) *ModifySQLRateLimitingRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySQLRateLimitingRulesResponseBody) SetSuccess(v bool) *ModifySQLRateLimitingRulesResponseBody {
	s.Success = &v
	return s
}

func (s *ModifySQLRateLimitingRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
