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
	// The status. Valid values:
	//
	// - active
	//
	// - inactive
	//
	// example:
	//
	// active
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
