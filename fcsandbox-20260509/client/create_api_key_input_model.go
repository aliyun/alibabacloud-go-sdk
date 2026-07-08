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
	SetTeamID(v string) *CreateApiKeyInput
	GetTeamID() *string
}

type CreateApiKeyInput struct {
	ApiKeyName *string `json:"apiKeyName,omitempty" xml:"apiKeyName,omitempty"`
	ExpireTime *string `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	TeamID     *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
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

func (s *CreateApiKeyInput) SetTeamID(v string) *CreateApiKeyInput {
	s.TeamID = &v
	return s
}

func (s *CreateApiKeyInput) Validate() error {
	return dara.Validate(s)
}
