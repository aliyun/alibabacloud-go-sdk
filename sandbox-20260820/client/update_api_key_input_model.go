// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApiKeyInput interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyName(v string) *UpdateApiKeyInput
	GetApiKeyName() *string
	SetExpireTime(v string) *UpdateApiKeyInput
	GetExpireTime() *string
	SetIpBlacklist(v []*IPConfig) *UpdateApiKeyInput
	GetIpBlacklist() []*IPConfig
	SetIpWhitelist(v []*IPConfig) *UpdateApiKeyInput
	GetIpWhitelist() []*IPConfig
	SetStatus(v string) *UpdateApiKeyInput
	GetStatus() *string
	SetVpcBindings(v []*string) *UpdateApiKeyInput
	GetVpcBindings() []*string
}

type UpdateApiKeyInput struct {
	// The new name of the API key. The name can be up to 128 characters in length and can contain only letters, digits, spaces, hyphens (-), underscores (_), and periods (.).
	//
	// example:
	//
	// ci-pipeline-key-v2
	ApiKeyName *string `json:"apiKeyName,omitempty" xml:"apiKeyName,omitempty"`
	// The expiration time of the API key in UTC. The time must be in RFC 3339 format. An empty value indicates that the API key never expires.
	//
	// example:
	//
	// 2099-12-31T23:59:59Z
	ExpireTime *string `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// The IP address blacklist. If you specify this parameter, the existing blacklist is completely overwritten. This parameter is mutually exclusive with ipWhitelist.
	//
	// example:
	//
	// [{"ipAddress":"192.0.2.1","description":"example client"}]
	IpBlacklist []*IPConfig `json:"ipBlacklist" xml:"ipBlacklist" type:"Repeated"`
	// The IP address whitelist. If you specify this parameter, the existing whitelist is completely overwritten. This parameter is mutually exclusive with ipBlacklist.
	//
	// example:
	//
	// [{"ipAddress":"192.0.2.1","description":"example client"}]
	IpWhitelist []*IPConfig `json:"ipWhitelist" xml:"ipWhitelist" type:"Repeated"`
	// The new status of the API key. Valid values:
	//
	// - active: Enabled.
	//
	// - inactive: Disabled.
	//
	// example:
	//
	// inactive
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The set of allowed source VPC IDs. This set is updated independently of the IP blacklists and whitelists. If you omit this parameter or set it to null, the original set is retained. A non-empty array completely replaces the original set. An explicit empty array [] purges the VPC restrictions. Updating only this parameter does not affect the IP address lists, and updating the IP address lists does not affect the VPC restrictions. An empty request body {} purges all IP address and VPC source restrictions. You can update this parameter along with other parameters such as the name.
	//
	// example:
	//
	// ["vpc-examplea","vpc-exampleb"]
	VpcBindings []*string `json:"vpcBindings" xml:"vpcBindings" type:"Repeated"`
}

func (s UpdateApiKeyInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateApiKeyInput) GoString() string {
	return s.String()
}

func (s *UpdateApiKeyInput) GetApiKeyName() *string {
	return s.ApiKeyName
}

func (s *UpdateApiKeyInput) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *UpdateApiKeyInput) GetIpBlacklist() []*IPConfig {
	return s.IpBlacklist
}

func (s *UpdateApiKeyInput) GetIpWhitelist() []*IPConfig {
	return s.IpWhitelist
}

func (s *UpdateApiKeyInput) GetStatus() *string {
	return s.Status
}

func (s *UpdateApiKeyInput) GetVpcBindings() []*string {
	return s.VpcBindings
}

func (s *UpdateApiKeyInput) SetApiKeyName(v string) *UpdateApiKeyInput {
	s.ApiKeyName = &v
	return s
}

func (s *UpdateApiKeyInput) SetExpireTime(v string) *UpdateApiKeyInput {
	s.ExpireTime = &v
	return s
}

func (s *UpdateApiKeyInput) SetIpBlacklist(v []*IPConfig) *UpdateApiKeyInput {
	s.IpBlacklist = v
	return s
}

func (s *UpdateApiKeyInput) SetIpWhitelist(v []*IPConfig) *UpdateApiKeyInput {
	s.IpWhitelist = v
	return s
}

func (s *UpdateApiKeyInput) SetStatus(v string) *UpdateApiKeyInput {
	s.Status = &v
	return s
}

func (s *UpdateApiKeyInput) SetVpcBindings(v []*string) *UpdateApiKeyInput {
	s.VpcBindings = v
	return s
}

func (s *UpdateApiKeyInput) Validate() error {
	if s.IpBlacklist != nil {
		for _, item := range s.IpBlacklist {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.IpWhitelist != nil {
		for _, item := range s.IpWhitelist {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
