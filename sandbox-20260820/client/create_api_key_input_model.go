// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApiKeyInput interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyName(v string) *CreateApiKeyInput
	GetApiKeyName() *string
	SetExpireTime(v string) *CreateApiKeyInput
	GetExpireTime() *string
	SetIpBlacklist(v []*IPConfig) *CreateApiKeyInput
	GetIpBlacklist() []*IPConfig
	SetIpWhitelist(v []*IPConfig) *CreateApiKeyInput
	GetIpWhitelist() []*IPConfig
	SetTeamID(v string) *CreateApiKeyInput
	GetTeamID() *string
}

type CreateApiKeyInput struct {
	// The name of the API key. The name can be up to 128 characters in length and can contain only letters, digits, spaces, hyphens (-), underscores (_), and periods (.).
	//
	// example:
	//
	// ci-pipeline-key
	ApiKeyName *string `json:"apiKeyName,omitempty" xml:"apiKeyName,omitempty"`
	// The expiration time of the API key. The time is in UTC and follows the RFC 3339 format. If this parameter is left empty, the API key never expires.
	//
	// example:
	//
	// 2099-12-31T23:59:59Z
	ExpireTime *string `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// The IP address blacklist. If this parameter is configured, the IP addresses in the list cannot use the API key. This parameter is mutually exclusive with ipWhitelist.
	IpBlacklist []*IPConfig `json:"ipBlacklist" xml:"ipBlacklist" type:"Repeated"`
	// The IP address whitelist. If this parameter is configured, only the IP addresses in the list can use the API key. This parameter is mutually exclusive with ipBlacklist.
	IpWhitelist []*IPConfig `json:"ipWhitelist" xml:"ipWhitelist" type:"Repeated"`
	// The ID of the team to which the API key belongs. The ID is in UUID format. If this parameter is not specified, the default team of the current account is used.
	//
	// example:
	//
	// 5f4a2c18-****
	TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
}

func (s CreateApiKeyInput) String() string {
	return dara.Prettify(s)
}

func (s CreateApiKeyInput) GoString() string {
	return s.String()
}

func (s *CreateApiKeyInput) GetApiKeyName() *string {
	return s.ApiKeyName
}

func (s *CreateApiKeyInput) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *CreateApiKeyInput) GetIpBlacklist() []*IPConfig {
	return s.IpBlacklist
}

func (s *CreateApiKeyInput) GetIpWhitelist() []*IPConfig {
	return s.IpWhitelist
}

func (s *CreateApiKeyInput) GetTeamID() *string {
	return s.TeamID
}

func (s *CreateApiKeyInput) SetApiKeyName(v string) *CreateApiKeyInput {
	s.ApiKeyName = &v
	return s
}

func (s *CreateApiKeyInput) SetExpireTime(v string) *CreateApiKeyInput {
	s.ExpireTime = &v
	return s
}

func (s *CreateApiKeyInput) SetIpBlacklist(v []*IPConfig) *CreateApiKeyInput {
	s.IpBlacklist = v
	return s
}

func (s *CreateApiKeyInput) SetIpWhitelist(v []*IPConfig) *CreateApiKeyInput {
	s.IpWhitelist = v
	return s
}

func (s *CreateApiKeyInput) SetTeamID(v string) *CreateApiKeyInput {
	s.TeamID = &v
	return s
}

func (s *CreateApiKeyInput) Validate() error {
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
