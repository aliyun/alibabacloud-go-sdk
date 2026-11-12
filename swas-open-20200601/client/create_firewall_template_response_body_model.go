// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFirewallTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFirewallTemplateId(v string) *CreateFirewallTemplateResponseBody
	GetFirewallTemplateId() *string
	SetRequestId(v string) *CreateFirewallTemplateResponseBody
	GetRequestId() *string
}

type CreateFirewallTemplateResponseBody struct {
	// The firewall template ID.
	//
	// example:
	//
	// ft-bcf1a7hrdq717****
	FirewallTemplateId *string `json:"FirewallTemplateId,omitempty" xml:"FirewallTemplateId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFirewallTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFirewallTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFirewallTemplateResponseBody) GetFirewallTemplateId() *string {
	return s.FirewallTemplateId
}

func (s *CreateFirewallTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFirewallTemplateResponseBody) SetFirewallTemplateId(v string) *CreateFirewallTemplateResponseBody {
	s.FirewallTemplateId = &v
	return s
}

func (s *CreateFirewallTemplateResponseBody) SetRequestId(v string) *CreateFirewallTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFirewallTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
