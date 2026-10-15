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
	// The name of the API key.
	//
	// example:
	//
	// dev
	ApiKeyName *string `json:"apiKeyName,omitempty" xml:"apiKeyName,omitempty"`
	// The value of the API key to import.
	ApiKeyValue *string `json:"apiKeyValue,omitempty" xml:"apiKeyValue,omitempty"`
	// The expiration time of the API key.
	ExpireTime *string `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// The unique identifier of the team to which the API key belongs.
	//
	// example:
	//
	// 70d1c834-0383-58d8-97ac-5336eb91abcd
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
