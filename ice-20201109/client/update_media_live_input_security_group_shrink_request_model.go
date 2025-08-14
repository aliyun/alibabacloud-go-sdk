// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaLiveInputSecurityGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *UpdateMediaLiveInputSecurityGroupShrinkRequest
	GetName() *string
	SetSecurityGroupId(v string) *UpdateMediaLiveInputSecurityGroupShrinkRequest
	GetSecurityGroupId() *string
	SetWhitelistRulesShrink(v string) *UpdateMediaLiveInputSecurityGroupShrinkRequest
	GetWhitelistRulesShrink() *string
}

type UpdateMediaLiveInputSecurityGroupShrinkRequest struct {
	// The name of the security group. Letters, digits, hyphens (-), and underscores (_) are supported. It can be up to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// mysg
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the security group.
	//
	// This parameter is required.
	//
	// example:
	//
	// SEGK5KA6KYKAWQQH
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The security group rules.
	//
	// This parameter is required.
	WhitelistRulesShrink *string `json:"WhitelistRules,omitempty" xml:"WhitelistRules,omitempty"`
}

func (s UpdateMediaLiveInputSecurityGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaLiveInputSecurityGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateMediaLiveInputSecurityGroupShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateMediaLiveInputSecurityGroupShrinkRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *UpdateMediaLiveInputSecurityGroupShrinkRequest) GetWhitelistRulesShrink() *string {
	return s.WhitelistRulesShrink
}

func (s *UpdateMediaLiveInputSecurityGroupShrinkRequest) SetName(v string) *UpdateMediaLiveInputSecurityGroupShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateMediaLiveInputSecurityGroupShrinkRequest) SetSecurityGroupId(v string) *UpdateMediaLiveInputSecurityGroupShrinkRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *UpdateMediaLiveInputSecurityGroupShrinkRequest) SetWhitelistRulesShrink(v string) *UpdateMediaLiveInputSecurityGroupShrinkRequest {
	s.WhitelistRulesShrink = &v
	return s
}

func (s *UpdateMediaLiveInputSecurityGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
