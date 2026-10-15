// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportApiKeyInput interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyName(v string) *ImportApiKeyInput
	GetApiKeyName() *string
	SetApiKeyValue(v string) *ImportApiKeyInput
	GetApiKeyValue() *string
	SetExpireTime(v string) *ImportApiKeyInput
	GetExpireTime() *string
	SetTeamID(v string) *ImportApiKeyInput
	GetTeamID() *string
}

type ImportApiKeyInput struct {
	// The API key name. The name can be up to 128 characters in length and can contain letters, digits, spaces, hyphens (-), underscores (_), and periods (.).
	//
	// example:
	//
	// migrated-key
	ApiKeyName *string `json:"apiKeyName,omitempty" xml:"apiKeyName,omitempty"`
	// The plaintext value of the API key to import. After the import, the source of the key is set to custom.
	//
	// example:
	//
	// e2b_****
	ApiKeyValue *string `json:"apiKeyValue,omitempty" xml:"apiKeyValue,omitempty"`
	// The expiration time of the API key in UTC. The time is in RFC 3339 format. If this parameter is left empty, the API key never expires.
	//
	// example:
	//
	// 2099-12-31T23:59:59Z
	ExpireTime *string `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// The ID of the team to which the API key belongs, in UUID format. If this parameter is not specified, the default team of the current account is used.
	//
	// example:
	//
	// 5f4a2c18-****
	TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
}

func (s ImportApiKeyInput) String() string {
	return dara.Prettify(s)
}

func (s ImportApiKeyInput) GoString() string {
	return s.String()
}

func (s *ImportApiKeyInput) GetApiKeyName() *string {
	return s.ApiKeyName
}

func (s *ImportApiKeyInput) GetApiKeyValue() *string {
	return s.ApiKeyValue
}

func (s *ImportApiKeyInput) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ImportApiKeyInput) GetTeamID() *string {
	return s.TeamID
}

func (s *ImportApiKeyInput) SetApiKeyName(v string) *ImportApiKeyInput {
	s.ApiKeyName = &v
	return s
}

func (s *ImportApiKeyInput) SetApiKeyValue(v string) *ImportApiKeyInput {
	s.ApiKeyValue = &v
	return s
}

func (s *ImportApiKeyInput) SetExpireTime(v string) *ImportApiKeyInput {
	s.ExpireTime = &v
	return s
}

func (s *ImportApiKeyInput) SetTeamID(v string) *ImportApiKeyInput {
	s.TeamID = &v
	return s
}

func (s *ImportApiKeyInput) Validate() error {
	return dara.Validate(s)
}
