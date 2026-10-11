// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachEgressProxyRuleGroupsToConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachEgressProxyRuleGroupsToConfigurationResponseBody
	GetRequestId() *string
}

type AttachEgressProxyRuleGroupsToConfigurationResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 11E669F9-B301-575C-AB02-C54E41ADFA4B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachEgressProxyRuleGroupsToConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachEgressProxyRuleGroupsToConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *AttachEgressProxyRuleGroupsToConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachEgressProxyRuleGroupsToConfigurationResponseBody) SetRequestId(v string) *AttachEgressProxyRuleGroupsToConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachEgressProxyRuleGroupsToConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}
