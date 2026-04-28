// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRateLimitPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyRateLimitPolicyResponseBody
	GetRequestId() *string
}

type ModifyRateLimitPolicyResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 24A1990B-4F6E-482B-B8CB-75C612******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyRateLimitPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyRateLimitPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRateLimitPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyRateLimitPolicyResponseBody) SetRequestId(v string) *ModifyRateLimitPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyRateLimitPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
