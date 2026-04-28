// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRateLimitPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRateLimitPolicyResponseBody
	GetRequestId() *string
}

type DeleteRateLimitPolicyResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 6BD9CDE4-5E7B-4BF3-9BB8-83C73E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRateLimitPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRateLimitPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRateLimitPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRateLimitPolicyResponseBody) SetRequestId(v string) *DeleteRateLimitPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRateLimitPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
