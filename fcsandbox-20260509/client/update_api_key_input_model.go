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
	SetStatus(v string) *UpdateApiKeyInput
	GetStatus() *string
}

type UpdateApiKeyInput struct {
	ApiKeyName *string `json:"apiKeyName,omitempty" xml:"apiKeyName,omitempty"`
	ExpireTime *string `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	Status     *string `json:"status,omitempty" xml:"status,omitempty"`
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

func (s *UpdateApiKeyInput) SetStatus(v string) *UpdateApiKeyInput {
	s.Status = &v
	return s
}

func (s *UpdateApiKeyInput) Validate() error {
	return dara.Validate(s)
}
