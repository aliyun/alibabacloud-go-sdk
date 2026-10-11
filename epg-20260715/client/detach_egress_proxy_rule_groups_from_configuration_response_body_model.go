// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachEgressProxyRuleGroupsFromConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachEgressProxyRuleGroupsFromConfigurationResponseBody
	GetRequestId() *string
}

type DetachEgressProxyRuleGroupsFromConfigurationResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 01A05AC2-567F-5BB5-8F2D-436C4BE402E3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachEgressProxyRuleGroupsFromConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachEgressProxyRuleGroupsFromConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *DetachEgressProxyRuleGroupsFromConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachEgressProxyRuleGroupsFromConfigurationResponseBody) SetRequestId(v string) *DetachEgressProxyRuleGroupsFromConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachEgressProxyRuleGroupsFromConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}
