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
	SetLastUsedTime(v string) *ApiKey
	GetLastUsedTime() *string
	SetUserID(v string) *ApiKey
	GetUserID() *string
	SetUsername(v string) *ApiKey
	GetUsername() *string
}

type ApiKey struct {
	ApiKeyID     *string `json:"apiKeyID,omitempty" xml:"apiKeyID,omitempty"`
	ApiKeyMask   *string `json:"apiKeyMask,omitempty" xml:"apiKeyMask,omitempty"`
	ApiKeyName   *string `json:"apiKeyName,omitempty" xml:"apiKeyName,omitempty"`
	ApiKeyValue  *string `json:"apiKeyValue,omitempty" xml:"apiKeyValue,omitempty"`
	CreatedTime  *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	ExpireTime   *string `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	LastUsedTime *string `json:"lastUsedTime,omitempty" xml:"lastUsedTime,omitempty"`
	UserID       *string `json:"userID,omitempty" xml:"userID,omitempty"`
	Username     *string `json:"username,omitempty" xml:"username,omitempty"`
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

func (s *ApiKey) GetLastUsedTime() *string {
	return s.LastUsedTime
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

func (s *ApiKey) SetLastUsedTime(v string) *ApiKey {
	s.LastUsedTime = &v
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
	return dara.Validate(s)
}
