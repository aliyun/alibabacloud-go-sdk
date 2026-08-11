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
	// The unique identifier of the API key.
	//
	// example:
	//
	// 34f6a4c4-499f-4bbd-baa0-0e699f53abcd
	ApiKeyID *string `json:"apiKeyID,omitempty" xml:"apiKeyID,omitempty"`
	// The masked display value of the API key.
	//
	// example:
	//
	// e2b_xxxx****xxxx
	ApiKeyMask *string `json:"apiKeyMask,omitempty" xml:"apiKeyMask,omitempty"`
	// The name of the API key.
	//
	// example:
	//
	// dev
	ApiKeyName *string `json:"apiKeyName,omitempty" xml:"apiKeyName,omitempty"`
	// The value of the API key.
	//
	// example:
	//
	// e2b_xxxxxx79cd777ef8exxxxxx4ad6f1b567cxxxxxx
	ApiKeyValue *string `json:"apiKeyValue,omitempty" xml:"apiKeyValue,omitempty"`
	// The time when the API key was created.
	//
	// example:
	//
	// 2023-09-13T08:27:20Z
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// The expiration time.
	//
	// example:
	//
	// 2023-10-13T08:27:20Z
	ExpireTime  *string     `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	IpBlacklist []*IPConfig `json:"ipBlacklist,omitempty" xml:"ipBlacklist,omitempty" type:"Repeated"`
	IpWhitelist []*IPConfig `json:"ipWhitelist,omitempty" xml:"ipWhitelist,omitempty" type:"Repeated"`
	// The time when the API key was last used.
	//
	// example:
	//
	// 2023-09-14T08:27:20Z
	LastUsedTime *string `json:"lastUsedTime,omitempty" xml:"lastUsedTime,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmwxqyrgwabcd
	ResourceGroupID *string `json:"resourceGroupID,omitempty" xml:"resourceGroupID,omitempty"`
	Source          *string `json:"source,omitempty" xml:"source,omitempty"`
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
	// The unique identifier of the team.
	//
	// example:
	//
	// 70d1c834-0383-58d8-97ac-5336eb91abcd
	TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
	// The name of the team.
	//
	// example:
	//
	// Development Team
	TeamName *string `json:"teamName,omitempty" xml:"teamName,omitempty"`
	TeamPlan *string `json:"teamPlan,omitempty" xml:"teamPlan,omitempty"`
	// The UID of the creator.
	//
	// example:
	//
	// 12345
	UserID *string `json:"userID,omitempty" xml:"userID,omitempty"`
	// The creator.
	//
	// example:
	//
	// user1
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
