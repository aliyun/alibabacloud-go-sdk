// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaLiveInputSecurityGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *UpdateMediaLiveInputSecurityGroupRequest
	GetName() *string
	SetSecurityGroupId(v string) *UpdateMediaLiveInputSecurityGroupRequest
	GetSecurityGroupId() *string
	SetWhitelistRules(v []*string) *UpdateMediaLiveInputSecurityGroupRequest
	GetWhitelistRules() []*string
}

type UpdateMediaLiveInputSecurityGroupRequest struct {
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
	WhitelistRules []*string `json:"WhitelistRules,omitempty" xml:"WhitelistRules,omitempty" type:"Repeated"`
}

func (s UpdateMediaLiveInputSecurityGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaLiveInputSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateMediaLiveInputSecurityGroupRequest) GetName() *string {
	return s.Name
}

func (s *UpdateMediaLiveInputSecurityGroupRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *UpdateMediaLiveInputSecurityGroupRequest) GetWhitelistRules() []*string {
	return s.WhitelistRules
}

func (s *UpdateMediaLiveInputSecurityGroupRequest) SetName(v string) *UpdateMediaLiveInputSecurityGroupRequest {
	s.Name = &v
	return s
}

func (s *UpdateMediaLiveInputSecurityGroupRequest) SetSecurityGroupId(v string) *UpdateMediaLiveInputSecurityGroupRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *UpdateMediaLiveInputSecurityGroupRequest) SetWhitelistRules(v []*string) *UpdateMediaLiveInputSecurityGroupRequest {
	s.WhitelistRules = v
	return s
}

func (s *UpdateMediaLiveInputSecurityGroupRequest) Validate() error {
	return dara.Validate(s)
}
