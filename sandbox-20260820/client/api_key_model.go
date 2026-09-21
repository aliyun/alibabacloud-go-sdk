// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApiKey interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyID(v string) *ApiKey
	GetApiKeyID() *string
	SetApiKeyMask(v string) *ApiKey
	GetApiKeyMask() *string
	SetApiKeyName(v string) *ApiKey
	GetApiKeyName() *string
	SetApiKeyValue(v string) *ApiKey
	GetApiKeyValue() *string
	SetCreatedTime(v string) *ApiKey
	GetCreatedTime() *string
	SetExpireTime(v string) *ApiKey
	GetExpireTime() *string
	SetIpBlacklist(v []*IPConfig) *ApiKey
	GetIpBlacklist() []*IPConfig
	SetIpWhitelist(v []*IPConfig) *ApiKey
	GetIpWhitelist() []*IPConfig
	SetLastUsedTime(v string) *ApiKey
	GetLastUsedTime() *string
	SetResourceGroupID(v string) *ApiKey
	GetResourceGroupID() *string
	SetSource(v string) *ApiKey
	GetSource() *string
	SetStatus(v string) *ApiKey
	GetStatus() *string
	SetTeamID(v string) *ApiKey
	GetTeamID() *string
	SetTeamName(v string) *ApiKey
	GetTeamName() *string
	SetTeamPlan(v string) *ApiKey
	GetTeamPlan() *string
	SetUserID(v string) *ApiKey
	GetUserID() *string
	SetUsername(v string) *ApiKey
	GetUsername() *string
}

type ApiKey struct {
	// example:
	//
	// a1f8c3d6-****
	ApiKeyID *string `json:"apiKeyID,omitempty" xml:"apiKeyID,omitempty"`
	// example:
	//
	// e2b_3f9a****b915
	ApiKeyMask *string `json:"apiKeyMask,omitempty" xml:"apiKeyMask,omitempty"`
	// example:
	//
	// ci-pipeline-key
	ApiKeyName *string `json:"apiKeyName,omitempty" xml:"apiKeyName,omitempty"`
	// example:
	//
	// e2b_****
	ApiKeyValue *string `json:"apiKeyValue,omitempty" xml:"apiKeyValue,omitempty"`
	// example:
	//
	// 2026-08-20T08:30:00Z
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// example:
	//
	// 2099-12-31T23:59:59Z
	ExpireTime  *string     `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	IpBlacklist []*IPConfig `json:"ipBlacklist,omitempty" xml:"ipBlacklist,omitempty" type:"Repeated"`
	IpWhitelist []*IPConfig `json:"ipWhitelist,omitempty" xml:"ipWhitelist,omitempty" type:"Repeated"`
	// example:
	//
	// 2026-09-10T14:22:07Z
	LastUsedTime *string `json:"lastUsedTime,omitempty" xml:"lastUsedTime,omitempty"`
	// example:
	//
	// rg-****
	ResourceGroupID *string `json:"resourceGroupID,omitempty" xml:"resourceGroupID,omitempty"`
	// example:
	//
	// generated
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 5f4a2c18-****
	TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
	// example:
	//
	// sandbox-dev
	TeamName *string `json:"teamName,omitempty" xml:"teamName,omitempty"`
	// example:
	//
	// std
	TeamPlan *string `json:"teamPlan,omitempty" xml:"teamPlan,omitempty"`
	// example:
	//
	// 9c1d4e72-****
	UserID *string `json:"userID,omitempty" xml:"userID,omitempty"`
	// example:
	//
	// sandbox-admin
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s ApiKey) String() string {
	return dara.Prettify(s)
}

func (s ApiKey) GoString() string {
	return s.String()
}

func (s *ApiKey) GetApiKeyID() *string {
	return s.ApiKeyID
}

func (s *ApiKey) GetApiKeyMask() *string {
	return s.ApiKeyMask
}

func (s *ApiKey) GetApiKeyName() *string {
	return s.ApiKeyName
}

func (s *ApiKey) GetApiKeyValue() *string {
	return s.ApiKeyValue
}

func (s *ApiKey) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *ApiKey) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ApiKey) GetIpBlacklist() []*IPConfig {
	return s.IpBlacklist
}

func (s *ApiKey) GetIpWhitelist() []*IPConfig {
	return s.IpWhitelist
}

func (s *ApiKey) GetLastUsedTime() *string {
	return s.LastUsedTime
}

func (s *ApiKey) GetResourceGroupID() *string {
	return s.ResourceGroupID
}

func (s *ApiKey) GetSource() *string {
	return s.Source
}

func (s *ApiKey) GetStatus() *string {
	return s.Status
}

func (s *ApiKey) GetTeamID() *string {
	return s.TeamID
}

func (s *ApiKey) GetTeamName() *string {
	return s.TeamName
}

func (s *ApiKey) GetTeamPlan() *string {
	return s.TeamPlan
}

func (s *ApiKey) GetUserID() *string {
	return s.UserID
}

func (s *ApiKey) GetUsername() *string {
	return s.Username
}

func (s *ApiKey) SetApiKeyID(v string) *ApiKey {
	s.ApiKeyID = &v
	return s
}

func (s *ApiKey) SetApiKeyMask(v string) *ApiKey {
	s.ApiKeyMask = &v
	return s
}

func (s *ApiKey) SetApiKeyName(v string) *ApiKey {
	s.ApiKeyName = &v
	return s
}

func (s *ApiKey) SetApiKeyValue(v string) *ApiKey {
	s.ApiKeyValue = &v
	return s
}

func (s *ApiKey) SetCreatedTime(v string) *ApiKey {
	s.CreatedTime = &v
	return s
}

func (s *ApiKey) SetExpireTime(v string) *ApiKey {
	s.ExpireTime = &v
	return s
}

func (s *ApiKey) SetIpBlacklist(v []*IPConfig) *ApiKey {
	s.IpBlacklist = v
	return s
}

func (s *ApiKey) SetIpWhitelist(v []*IPConfig) *ApiKey {
	s.IpWhitelist = v
	return s
}

func (s *ApiKey) SetLastUsedTime(v string) *ApiKey {
	s.LastUsedTime = &v
	return s
}

func (s *ApiKey) SetResourceGroupID(v string) *ApiKey {
	s.ResourceGroupID = &v
	return s
}

func (s *ApiKey) SetSource(v string) *ApiKey {
	s.Source = &v
	return s
}

func (s *ApiKey) SetStatus(v string) *ApiKey {
	s.Status = &v
	return s
}

func (s *ApiKey) SetTeamID(v string) *ApiKey {
	s.TeamID = &v
	return s
}

func (s *ApiKey) SetTeamName(v string) *ApiKey {
	s.TeamName = &v
	return s
}

func (s *ApiKey) SetTeamPlan(v string) *ApiKey {
	s.TeamPlan = &v
	return s
}

func (s *ApiKey) SetUserID(v string) *ApiKey {
	s.UserID = &v
	return s
}

func (s *ApiKey) SetUsername(v string) *ApiKey {
	s.Username = &v
	return s
}

func (s *ApiKey) Validate() error {
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
