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
}

type UpdateApiKeyInput struct {
	// The new name of the API key. The name can be up to 128 characters in length and can contain letters, digits, spaces, hyphens (-), underscores (_), and periods (.).
	//
	// example:
	//
	// ci-pipeline-key-v2
	ApiKeyName *string `json:"apiKeyName,omitempty" xml:"apiKeyName,omitempty"`
	// The expiration time of the API key, in UTC, in RFC 3339 format. Leave this parameter empty to indicate that the API key never expires.
	//
	// example:
	//
	// 2099-12-31T23:59:59Z
	ExpireTime *string `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// The IP blacklist. If specified, the existing blacklist is entirely overwritten. This parameter is mutually exclusive with ipWhitelist.
	IpBlacklist []*IPConfig `json:"ipBlacklist" xml:"ipBlacklist" type:"Repeated"`
	// The IP whitelist. If specified, the existing whitelist is entirely overwritten. This parameter is mutually exclusive with ipBlacklist.
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
