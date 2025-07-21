// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNetworkRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateNetworkRuleRequest
	GetDescription() *string
	SetName(v string) *UpdateNetworkRuleRequest
	GetName() *string
	SetSourcePrivateIp(v string) *UpdateNetworkRuleRequest
	GetSourcePrivateIp() *string
}

type UpdateNetworkRuleRequest struct {
	// The description after the update.
	//
	// example:
	//
	// Creat by kst-hzz62ee817bvyyr5****
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the access control rule that you want to update.
	//
	// This parameter is required.
	//
	// example:
	//
	// networkrule_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The private IP address or CIDR block after the update. Separate multiple items with commas (,).
	//
	// example:
	//
	// ["192.10.XX.XX","192.168.XX.XX/24"]
	SourcePrivateIp *string `json:"SourcePrivateIp,omitempty" xml:"SourcePrivateIp,omitempty"`
}

func (s UpdateNetworkRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNetworkRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateNetworkRuleRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateNetworkRuleRequest) GetName() *string {
	return s.Name
}

func (s *UpdateNetworkRuleRequest) GetSourcePrivateIp() *string {
	return s.SourcePrivateIp
}

func (s *UpdateNetworkRuleRequest) SetDescription(v string) *UpdateNetworkRuleRequest {
	s.Description = &v
	return s
}

func (s *UpdateNetworkRuleRequest) SetName(v string) *UpdateNetworkRuleRequest {
	s.Name = &v
	return s
}

func (s *UpdateNetworkRuleRequest) SetSourcePrivateIp(v string) *UpdateNetworkRuleRequest {
	s.SourcePrivateIp = &v
	return s
}

func (s *UpdateNetworkRuleRequest) Validate() error {
	return dara.Validate(s)
}
