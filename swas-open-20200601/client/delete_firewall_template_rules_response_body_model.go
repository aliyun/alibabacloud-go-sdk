// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFirewallTemplateRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteFirewallTemplateRulesResponseBody
	GetRequestId() *string
}

type DeleteFirewallTemplateRulesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFirewallTemplateRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFirewallTemplateRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFirewallTemplateRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFirewallTemplateRulesResponseBody) SetRequestId(v string) *DeleteFirewallTemplateRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFirewallTemplateRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
