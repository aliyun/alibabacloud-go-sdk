// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFirewallRulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateFirewallRulesShrinkRequest
	GetClientToken() *string
	SetFirewallRulesShrink(v string) *CreateFirewallRulesShrinkRequest
	GetFirewallRulesShrink() *string
	SetInstanceId(v string) *CreateFirewallRulesShrinkRequest
	GetInstanceId() *string
	SetRegionId(v string) *CreateFirewallRulesShrinkRequest
	GetRegionId() *string
	SetTag(v []*CreateFirewallRulesShrinkRequestTag) *CreateFirewallRulesShrinkRequest
	GetTag() []*CreateFirewallRulesShrinkRequestTag
}

type CreateFirewallRulesShrinkRequest struct {
	// The client token.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The array of firewall rules.
	FirewallRulesShrink *string `json:"FirewallRules,omitempty" xml:"FirewallRules,omitempty"`
	// The instance ID of the specified simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the regions where simple application servers are available.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of tags. The list can contain up to 20 tags.
	Tag []*CreateFirewallRulesShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateFirewallRulesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFirewallRulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateFirewallRulesShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateFirewallRulesShrinkRequest) GetFirewallRulesShrink() *string {
	return s.FirewallRulesShrink
}

func (s *CreateFirewallRulesShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateFirewallRulesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateFirewallRulesShrinkRequest) GetTag() []*CreateFirewallRulesShrinkRequestTag {
	return s.Tag
}

func (s *CreateFirewallRulesShrinkRequest) SetClientToken(v string) *CreateFirewallRulesShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateFirewallRulesShrinkRequest) SetFirewallRulesShrink(v string) *CreateFirewallRulesShrinkRequest {
	s.FirewallRulesShrink = &v
	return s
}

func (s *CreateFirewallRulesShrinkRequest) SetInstanceId(v string) *CreateFirewallRulesShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateFirewallRulesShrinkRequest) SetRegionId(v string) *CreateFirewallRulesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateFirewallRulesShrinkRequest) SetTag(v []*CreateFirewallRulesShrinkRequestTag) *CreateFirewallRulesShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreateFirewallRulesShrinkRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateFirewallRulesShrinkRequestTag struct {
	// The tag key that is specified for the firewall rule. Valid values of N: 1 to 20.
	//
	// The tag key cannot be an empty string. It can be up to 64 characters in length, cannot start with aliyun or acs:, and cannot contain http:// or https://.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value that is specified for the firewall rule. Valid values of N: 1 to 20.
	//
	// The tag value can be an empty string. It can be up to 64 characters in length and cannot contain http:// or https://.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateFirewallRulesShrinkRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateFirewallRulesShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateFirewallRulesShrinkRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateFirewallRulesShrinkRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateFirewallRulesShrinkRequestTag) SetKey(v string) *CreateFirewallRulesShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateFirewallRulesShrinkRequestTag) SetValue(v string) *CreateFirewallRulesShrinkRequestTag {
	s.Value = &v
	return s
}

func (s *CreateFirewallRulesShrinkRequestTag) Validate() error {
	return dara.Validate(s)
}
