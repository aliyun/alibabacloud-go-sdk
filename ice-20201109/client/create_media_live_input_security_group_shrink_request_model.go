// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMediaLiveInputSecurityGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreateMediaLiveInputSecurityGroupShrinkRequest
	GetName() *string
	SetWhitelistRulesShrink(v string) *CreateMediaLiveInputSecurityGroupShrinkRequest
	GetWhitelistRulesShrink() *string
}

type CreateMediaLiveInputSecurityGroupShrinkRequest struct {
	// The name of the security group. Letters, digits, hyphens (-), and underscores (_) are supported. The maximum length is 64 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// mysg
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The security group rules.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["10.1.1.0/24", "11.11.11.11/0"]
	WhitelistRulesShrink *string `json:"WhitelistRules,omitempty" xml:"WhitelistRules,omitempty"`
}

func (s CreateMediaLiveInputSecurityGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaLiveInputSecurityGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateMediaLiveInputSecurityGroupShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateMediaLiveInputSecurityGroupShrinkRequest) GetWhitelistRulesShrink() *string {
	return s.WhitelistRulesShrink
}

func (s *CreateMediaLiveInputSecurityGroupShrinkRequest) SetName(v string) *CreateMediaLiveInputSecurityGroupShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateMediaLiveInputSecurityGroupShrinkRequest) SetWhitelistRulesShrink(v string) *CreateMediaLiveInputSecurityGroupShrinkRequest {
	s.WhitelistRulesShrink = &v
	return s
}

func (s *CreateMediaLiveInputSecurityGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
