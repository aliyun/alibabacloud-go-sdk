// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMediaLiveInputSecurityGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreateMediaLiveInputSecurityGroupRequest
	GetName() *string
	SetWhitelistRules(v []*string) *CreateMediaLiveInputSecurityGroupRequest
	GetWhitelistRules() []*string
}

type CreateMediaLiveInputSecurityGroupRequest struct {
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
	WhitelistRules []*string `json:"WhitelistRules,omitempty" xml:"WhitelistRules,omitempty" type:"Repeated"`
}

func (s CreateMediaLiveInputSecurityGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaLiveInputSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateMediaLiveInputSecurityGroupRequest) GetName() *string {
	return s.Name
}

func (s *CreateMediaLiveInputSecurityGroupRequest) GetWhitelistRules() []*string {
	return s.WhitelistRules
}

func (s *CreateMediaLiveInputSecurityGroupRequest) SetName(v string) *CreateMediaLiveInputSecurityGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateMediaLiveInputSecurityGroupRequest) SetWhitelistRules(v []*string) *CreateMediaLiveInputSecurityGroupRequest {
	s.WhitelistRules = v
	return s
}

func (s *CreateMediaLiveInputSecurityGroupRequest) Validate() error {
	return dara.Validate(s)
}
