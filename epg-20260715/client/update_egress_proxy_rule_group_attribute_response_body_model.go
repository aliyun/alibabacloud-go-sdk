// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEgressProxyRuleGroupAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateEgressProxyRuleGroupAttributeResponseBody
	GetRequestId() *string
}

type UpdateEgressProxyRuleGroupAttributeResponseBody struct {
	// example:
	//
	// 01A05AB5-5BED-56ED-8EE9-304FBC2F8537
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateEgressProxyRuleGroupAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEgressProxyRuleGroupAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEgressProxyRuleGroupAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEgressProxyRuleGroupAttributeResponseBody) SetRequestId(v string) *UpdateEgressProxyRuleGroupAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEgressProxyRuleGroupAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
