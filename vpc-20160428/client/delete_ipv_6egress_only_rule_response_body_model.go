// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpv6EgressOnlyRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteIpv6EgressOnlyRuleResponseBody
	GetRequestId() *string
}

type DeleteIpv6EgressOnlyRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9DFEDBEE-E5AB-49E8-A2DC-CC114C67AF75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIpv6EgressOnlyRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpv6EgressOnlyRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIpv6EgressOnlyRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIpv6EgressOnlyRuleResponseBody) SetRequestId(v string) *DeleteIpv6EgressOnlyRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIpv6EgressOnlyRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
