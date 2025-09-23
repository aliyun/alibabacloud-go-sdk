// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSQLRateLimitingRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DeleteSQLRateLimitingRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteSQLRateLimitingRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteSQLRateLimitingRulesResponseBody
	GetSuccess() *bool
}

type DeleteSQLRateLimitingRulesResponseBody struct {
	// example:
	//
	// Message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 5E71541A-6007-4DCC-A38A-F872C31FEB45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSQLRateLimitingRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSQLRateLimitingRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSQLRateLimitingRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteSQLRateLimitingRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSQLRateLimitingRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteSQLRateLimitingRulesResponseBody) SetMessage(v string) *DeleteSQLRateLimitingRulesResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSQLRateLimitingRulesResponseBody) SetRequestId(v string) *DeleteSQLRateLimitingRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSQLRateLimitingRulesResponseBody) SetSuccess(v bool) *DeleteSQLRateLimitingRulesResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteSQLRateLimitingRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
