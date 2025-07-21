// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateNetworkRuleRequest
	GetDescription() *string
	SetName(v string) *CreateNetworkRuleRequest
	GetName() *string
	SetSourcePrivateIp(v string) *CreateNetworkRuleRequest
	GetSourcePrivateIp() *string
	SetType(v string) *CreateNetworkRuleRequest
	GetType() *string
}

type CreateNetworkRuleRequest struct {
	// The description.
	//
	// example:
	//
	// networkrule description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the access control rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// networkrule_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The private IP address or private CIDR block. Separate multiple items with commas (,).
	//
	// example:
	//
	// ["192.10.XX.XX","192.168.XX.XX/24"]
	SourcePrivateIp *string `json:"SourcePrivateIp,omitempty" xml:"SourcePrivateIp,omitempty"`
	// The network type.
	//
	// Only private IP addresses are supported. Set the value to Private.
	//
	// This parameter is required.
	//
	// example:
	//
	// Private
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateNetworkRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateNetworkRuleRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateNetworkRuleRequest) GetName() *string {
	return s.Name
}

func (s *CreateNetworkRuleRequest) GetSourcePrivateIp() *string {
	return s.SourcePrivateIp
}

func (s *CreateNetworkRuleRequest) GetType() *string {
	return s.Type
}

func (s *CreateNetworkRuleRequest) SetDescription(v string) *CreateNetworkRuleRequest {
	s.Description = &v
	return s
}

func (s *CreateNetworkRuleRequest) SetName(v string) *CreateNetworkRuleRequest {
	s.Name = &v
	return s
}

func (s *CreateNetworkRuleRequest) SetSourcePrivateIp(v string) *CreateNetworkRuleRequest {
	s.SourcePrivateIp = &v
	return s
}

func (s *CreateNetworkRuleRequest) SetType(v string) *CreateNetworkRuleRequest {
	s.Type = &v
	return s
}

func (s *CreateNetworkRuleRequest) Validate() error {
	return dara.Validate(s)
}
