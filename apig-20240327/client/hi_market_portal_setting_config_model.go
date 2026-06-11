// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHiMarketPortalSettingConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAutoApproveDevelopers(v bool) *HiMarketPortalSettingConfig
	GetAutoApproveDevelopers() *bool
	SetAutoApproveSubscriptions(v bool) *HiMarketPortalSettingConfig
	GetAutoApproveSubscriptions() *bool
	SetBuiltinAuthEnabled(v bool) *HiMarketPortalSettingConfig
	GetBuiltinAuthEnabled() *bool
}

type HiMarketPortalSettingConfig struct {
	// Specifies whether to automatically approve new developer registrations. If set to `false`, you must manually approve each new developer.\\
	//
	// \\
	//
	// **Default**: `false`.\\
	//
	// \\
	AutoApproveDevelopers *bool `json:"autoApproveDevelopers,omitempty" xml:"autoApproveDevelopers,omitempty"`
	// Specifies whether to automatically approve new API subscriptions. If set to `false`, you must manually approve each new subscription.\\
	//
	// \\
	//
	// **Default**: `false`.\\
	//
	// \\
	AutoApproveSubscriptions *bool `json:"autoApproveSubscriptions,omitempty" xml:"autoApproveSubscriptions,omitempty"`
	// Specifies whether to enable built-in authentication. If set to `true`, users must sign in to access the portal.\\
	//
	// \\
	//
	// **Default**: `false`.\\
	//
	// \\
	BuiltinAuthEnabled *bool `json:"builtinAuthEnabled,omitempty" xml:"builtinAuthEnabled,omitempty"`
}

func (s HiMarketPortalSettingConfig) String() string {
	return dara.Prettify(s)
}

func (s HiMarketPortalSettingConfig) GoString() string {
	return s.String()
}

func (s *HiMarketPortalSettingConfig) GetAutoApproveDevelopers() *bool {
	return s.AutoApproveDevelopers
}

func (s *HiMarketPortalSettingConfig) GetAutoApproveSubscriptions() *bool {
	return s.AutoApproveSubscriptions
}

func (s *HiMarketPortalSettingConfig) GetBuiltinAuthEnabled() *bool {
	return s.BuiltinAuthEnabled
}

func (s *HiMarketPortalSettingConfig) SetAutoApproveDevelopers(v bool) *HiMarketPortalSettingConfig {
	s.AutoApproveDevelopers = &v
	return s
}

func (s *HiMarketPortalSettingConfig) SetAutoApproveSubscriptions(v bool) *HiMarketPortalSettingConfig {
	s.AutoApproveSubscriptions = &v
	return s
}

func (s *HiMarketPortalSettingConfig) SetBuiltinAuthEnabled(v bool) *HiMarketPortalSettingConfig {
	s.BuiltinAuthEnabled = &v
	return s
}

func (s *HiMarketPortalSettingConfig) Validate() error {
	return dara.Validate(s)
}
