// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNatRateLimitRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyNatRateLimitRuleResponseBody
	GetRequestId() *string
}

type ModifyNatRateLimitRuleResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 2315DEB7-5E92-423A-91F7-4C1EC9AD97C3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyNatRateLimitRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyNatRateLimitRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNatRateLimitRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyNatRateLimitRuleResponseBody) SetRequestId(v string) *ModifyNatRateLimitRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyNatRateLimitRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
