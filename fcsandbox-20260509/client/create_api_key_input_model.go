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
	// The API key name.
	//
	// example:
	//
	// dev
	ApiKeyName *string `json:"apiKeyName,omitempty" xml:"apiKeyName,omitempty"`
	// The expiration time.
	//
	// example:
	//
	// 2026-07-24T16:00:00.000Z
	ExpireTime  *string     `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	IpBlacklist []*IPConfig `json:"ipBlacklist,omitempty" xml:"ipBlacklist,omitempty" type:"Repeated"`
	IpWhitelist []*IPConfig `json:"ipWhitelist,omitempty" xml:"ipWhitelist,omitempty" type:"Repeated"`
	// The unique identifier of the team.
	//
	// example:
	//
	// 70d1c834-0383-58d8-97ac-5336eb91abcd
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
