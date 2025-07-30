// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCredentialConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunEnvRoleKey(v string) *CredentialConfig
	GetAliyunEnvRoleKey() *string
	SetCredentialConfigItems(v []*CredentialConfigItem) *CredentialConfig
	GetCredentialConfigItems() []*CredentialConfigItem
	SetEnableCredentialInject(v bool) *CredentialConfig
	GetEnableCredentialInject() *bool
}

type CredentialConfig struct {
	AliyunEnvRoleKey       *string                 `json:"AliyunEnvRoleKey,omitempty" xml:"AliyunEnvRoleKey,omitempty"`
	CredentialConfigItems  []*CredentialConfigItem `json:"CredentialConfigItems,omitempty" xml:"CredentialConfigItems,omitempty" type:"Repeated"`
	EnableCredentialInject *bool                   `json:"EnableCredentialInject,omitempty" xml:"EnableCredentialInject,omitempty"`
}

func (s CredentialConfig) String() string {
	return dara.Prettify(s)
}

func (s CredentialConfig) GoString() string {
	return s.String()
}

func (s *CredentialConfig) GetAliyunEnvRoleKey() *string {
	return s.AliyunEnvRoleKey
}

func (s *CredentialConfig) GetCredentialConfigItems() []*CredentialConfigItem {
	return s.CredentialConfigItems
}

func (s *CredentialConfig) GetEnableCredentialInject() *bool {
	return s.EnableCredentialInject
}

func (s *CredentialConfig) SetAliyunEnvRoleKey(v string) *CredentialConfig {
	s.AliyunEnvRoleKey = &v
	return s
}

func (s *CredentialConfig) SetCredentialConfigItems(v []*CredentialConfigItem) *CredentialConfig {
	s.CredentialConfigItems = v
	return s
}

func (s *CredentialConfig) SetEnableCredentialInject(v bool) *CredentialConfig {
	s.EnableCredentialInject = &v
	return s
}

func (s *CredentialConfig) Validate() error {
	return dara.Validate(s)
}
