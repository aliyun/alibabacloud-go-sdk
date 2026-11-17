// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNatRateLimitRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteNatRateLimitRuleResponseBody
	GetRequestId() *string
}

type DeleteNatRateLimitRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2315DEB7-5E92-423A-91F7-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNatRateLimitRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNatRateLimitRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNatRateLimitRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNatRateLimitRuleResponseBody) SetRequestId(v string) *DeleteNatRateLimitRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNatRateLimitRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
