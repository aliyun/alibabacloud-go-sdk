// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIdpConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAttrMapConfig(v *IdpAttrMapConfig) *IdpConfig
	GetAttrMapConfig() *IdpAttrMapConfig
	SetConnectConfig(v *IdpConnectConfig) *IdpConfig
	GetConnectConfig() *IdpConnectConfig
	SetDescription(v string) *IdpConfig
	GetDescription() *string
	SetDingtalkConfig(v *IdpDingtalkSubConfig) *IdpConfig
	GetDingtalkConfig() *IdpDingtalkSubConfig
	SetEnabled(v bool) *IdpConfig
	GetEnabled() *bool
	SetFeishuConfig(v *IdpFeishuSubConfig) *IdpConfig
	GetFeishuConfig() *IdpFeishuSubConfig
	SetIdaasConfig(v *IdpIdaas2SubConfig) *IdpConfig
	GetIdaasConfig() *IdpIdaas2SubConfig
	SetIdpConfigId(v string) *IdpConfig
	GetIdpConfigId() *string
	SetLastSyncTimeUnix(v int64) *IdpConfig
	GetLastSyncTimeUnix() *int64
	SetLdapConfig(v *IdpLdapSubConfig) *IdpConfig
	GetLdapConfig() *IdpLdapSubConfig
	SetLoginConfig(v *IdpLoginConfig) *IdpConfig
	GetLoginConfig() *IdpLoginConfig
	SetName(v string) *IdpConfig
	GetName() *string
	SetSyncConfig(v *IdpSyncConfig) *IdpConfig
	GetSyncConfig() *IdpSyncConfig
	SetSyncStatus(v string) *IdpConfig
	GetSyncStatus() *string
	SetType(v string) *IdpConfig
	GetType() *string
	SetWeixinConfig(v *IdpWeixin2SubConfig) *IdpConfig
	GetWeixinConfig() *IdpWeixin2SubConfig
	SetWuyingConfig(v *OpenStructIdpWuyingSubConfig) *IdpConfig
	GetWuyingConfig() *OpenStructIdpWuyingSubConfig
}

type IdpConfig struct {
	AttrMapConfig    *IdpAttrMapConfig             `json:"AttrMapConfig,omitempty" xml:"AttrMapConfig,omitempty"`
	ConnectConfig    *IdpConnectConfig             `json:"ConnectConfig,omitempty" xml:"ConnectConfig,omitempty"`
	Description      *string                       `json:"Description,omitempty" xml:"Description,omitempty"`
	DingtalkConfig   *IdpDingtalkSubConfig         `json:"DingtalkConfig,omitempty" xml:"DingtalkConfig,omitempty"`
	Enabled          *bool                         `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	FeishuConfig     *IdpFeishuSubConfig           `json:"FeishuConfig,omitempty" xml:"FeishuConfig,omitempty"`
	IdaasConfig      *IdpIdaas2SubConfig           `json:"IdaasConfig,omitempty" xml:"IdaasConfig,omitempty"`
	IdpConfigId      *string                       `json:"IdpConfigId,omitempty" xml:"IdpConfigId,omitempty"`
	LastSyncTimeUnix *int64                        `json:"LastSyncTimeUnix,omitempty" xml:"LastSyncTimeUnix,omitempty"`
	LdapConfig       *IdpLdapSubConfig             `json:"LdapConfig,omitempty" xml:"LdapConfig,omitempty"`
	LoginConfig      *IdpLoginConfig               `json:"LoginConfig,omitempty" xml:"LoginConfig,omitempty"`
	Name             *string                       `json:"Name,omitempty" xml:"Name,omitempty"`
	SyncConfig       *IdpSyncConfig                `json:"SyncConfig,omitempty" xml:"SyncConfig,omitempty"`
	SyncStatus       *string                       `json:"SyncStatus,omitempty" xml:"SyncStatus,omitempty"`
	Type             *string                       `json:"Type,omitempty" xml:"Type,omitempty"`
	WeixinConfig     *IdpWeixin2SubConfig          `json:"WeixinConfig,omitempty" xml:"WeixinConfig,omitempty"`
	WuyingConfig     *OpenStructIdpWuyingSubConfig `json:"WuyingConfig,omitempty" xml:"WuyingConfig,omitempty"`
}

func (s IdpConfig) String() string {
	return dara.Prettify(s)
}

func (s IdpConfig) GoString() string {
	return s.String()
}

func (s *IdpConfig) GetAttrMapConfig() *IdpAttrMapConfig {
	return s.AttrMapConfig
}

func (s *IdpConfig) GetConnectConfig() *IdpConnectConfig {
	return s.ConnectConfig
}

func (s *IdpConfig) GetDescription() *string {
	return s.Description
}

func (s *IdpConfig) GetDingtalkConfig() *IdpDingtalkSubConfig {
	return s.DingtalkConfig
}

func (s *IdpConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *IdpConfig) GetFeishuConfig() *IdpFeishuSubConfig {
	return s.FeishuConfig
}

func (s *IdpConfig) GetIdaasConfig() *IdpIdaas2SubConfig {
	return s.IdaasConfig
}

func (s *IdpConfig) GetIdpConfigId() *string {
	return s.IdpConfigId
}

func (s *IdpConfig) GetLastSyncTimeUnix() *int64 {
	return s.LastSyncTimeUnix
}

func (s *IdpConfig) GetLdapConfig() *IdpLdapSubConfig {
	return s.LdapConfig
}

func (s *IdpConfig) GetLoginConfig() *IdpLoginConfig {
	return s.LoginConfig
}

func (s *IdpConfig) GetName() *string {
	return s.Name
}

func (s *IdpConfig) GetSyncConfig() *IdpSyncConfig {
	return s.SyncConfig
}

func (s *IdpConfig) GetSyncStatus() *string {
	return s.SyncStatus
}

func (s *IdpConfig) GetType() *string {
	return s.Type
}

func (s *IdpConfig) GetWeixinConfig() *IdpWeixin2SubConfig {
	return s.WeixinConfig
}

func (s *IdpConfig) GetWuyingConfig() *OpenStructIdpWuyingSubConfig {
	return s.WuyingConfig
}

func (s *IdpConfig) SetAttrMapConfig(v *IdpAttrMapConfig) *IdpConfig {
	s.AttrMapConfig = v
	return s
}

func (s *IdpConfig) SetConnectConfig(v *IdpConnectConfig) *IdpConfig {
	s.ConnectConfig = v
	return s
}

func (s *IdpConfig) SetDescription(v string) *IdpConfig {
	s.Description = &v
	return s
}

func (s *IdpConfig) SetDingtalkConfig(v *IdpDingtalkSubConfig) *IdpConfig {
	s.DingtalkConfig = v
	return s
}

func (s *IdpConfig) SetEnabled(v bool) *IdpConfig {
	s.Enabled = &v
	return s
}

func (s *IdpConfig) SetFeishuConfig(v *IdpFeishuSubConfig) *IdpConfig {
	s.FeishuConfig = v
	return s
}

func (s *IdpConfig) SetIdaasConfig(v *IdpIdaas2SubConfig) *IdpConfig {
	s.IdaasConfig = v
	return s
}

func (s *IdpConfig) SetIdpConfigId(v string) *IdpConfig {
	s.IdpConfigId = &v
	return s
}

func (s *IdpConfig) SetLastSyncTimeUnix(v int64) *IdpConfig {
	s.LastSyncTimeUnix = &v
	return s
}

func (s *IdpConfig) SetLdapConfig(v *IdpLdapSubConfig) *IdpConfig {
	s.LdapConfig = v
	return s
}

func (s *IdpConfig) SetLoginConfig(v *IdpLoginConfig) *IdpConfig {
	s.LoginConfig = v
	return s
}

func (s *IdpConfig) SetName(v string) *IdpConfig {
	s.Name = &v
	return s
}

func (s *IdpConfig) SetSyncConfig(v *IdpSyncConfig) *IdpConfig {
	s.SyncConfig = v
	return s
}

func (s *IdpConfig) SetSyncStatus(v string) *IdpConfig {
	s.SyncStatus = &v
	return s
}

func (s *IdpConfig) SetType(v string) *IdpConfig {
	s.Type = &v
	return s
}

func (s *IdpConfig) SetWeixinConfig(v *IdpWeixin2SubConfig) *IdpConfig {
	s.WeixinConfig = v
	return s
}

func (s *IdpConfig) SetWuyingConfig(v *OpenStructIdpWuyingSubConfig) *IdpConfig {
	s.WuyingConfig = v
	return s
}

func (s *IdpConfig) Validate() error {
	if s.AttrMapConfig != nil {
		if err := s.AttrMapConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ConnectConfig != nil {
		if err := s.ConnectConfig.Validate(); err != nil {
			return err
		}
	}
	if s.DingtalkConfig != nil {
		if err := s.DingtalkConfig.Validate(); err != nil {
			return err
		}
	}
	if s.FeishuConfig != nil {
		if err := s.FeishuConfig.Validate(); err != nil {
			return err
		}
	}
	if s.IdaasConfig != nil {
		if err := s.IdaasConfig.Validate(); err != nil {
			return err
		}
	}
	if s.LdapConfig != nil {
		if err := s.LdapConfig.Validate(); err != nil {
			return err
		}
	}
	if s.LoginConfig != nil {
		if err := s.LoginConfig.Validate(); err != nil {
			return err
		}
	}
	if s.SyncConfig != nil {
		if err := s.SyncConfig.Validate(); err != nil {
			return err
		}
	}
	if s.WeixinConfig != nil {
		if err := s.WeixinConfig.Validate(); err != nil {
			return err
		}
	}
	if s.WuyingConfig != nil {
		if err := s.WuyingConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
