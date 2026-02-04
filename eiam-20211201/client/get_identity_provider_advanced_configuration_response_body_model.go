// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentityProviderAdvancedConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAdvancedConfiguration(v *GetIdentityProviderAdvancedConfigurationResponseBodyAdvancedConfiguration) *GetIdentityProviderAdvancedConfigurationResponseBody
	GetAdvancedConfiguration() *GetIdentityProviderAdvancedConfigurationResponseBodyAdvancedConfiguration
	SetRequestId(v string) *GetIdentityProviderAdvancedConfigurationResponseBody
	GetRequestId() *string
}

type GetIdentityProviderAdvancedConfigurationResponseBody struct {
	AdvancedConfiguration *GetIdentityProviderAdvancedConfigurationResponseBodyAdvancedConfiguration `json:"AdvancedConfiguration,omitempty" xml:"AdvancedConfiguration,omitempty" type:"Struct"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetIdentityProviderAdvancedConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderAdvancedConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderAdvancedConfigurationResponseBody) GetAdvancedConfiguration() *GetIdentityProviderAdvancedConfigurationResponseBodyAdvancedConfiguration {
	return s.AdvancedConfiguration
}

func (s *GetIdentityProviderAdvancedConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIdentityProviderAdvancedConfigurationResponseBody) SetAdvancedConfiguration(v *GetIdentityProviderAdvancedConfigurationResponseBodyAdvancedConfiguration) *GetIdentityProviderAdvancedConfigurationResponseBody {
	s.AdvancedConfiguration = v
	return s
}

func (s *GetIdentityProviderAdvancedConfigurationResponseBody) SetRequestId(v string) *GetIdentityProviderAdvancedConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIdentityProviderAdvancedConfigurationResponseBody) Validate() error {
	if s.AdvancedConfiguration != nil {
		if err := s.AdvancedConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetIdentityProviderAdvancedConfigurationResponseBodyAdvancedConfiguration struct {
	// 钉钉高阶配置
	DingtalkAdvancedConfig *GetIdentityProviderAdvancedConfigurationResponseBodyAdvancedConfigurationDingtalkAdvancedConfig `json:"DingtalkAdvancedConfig,omitempty" xml:"DingtalkAdvancedConfig,omitempty" type:"Struct"`
	// IDaaS EIAM 身份提供方ID
	//
	// example:
	//
	// idp_na2rzpyc67zr7ixdfy35zgrxxx
	IdentityProviderId *string `json:"IdentityProviderId,omitempty" xml:"IdentityProviderId,omitempty"`
	// IDaaS EIAM 实例Id
	//
	// example:
	//
	// idaas_kpvmhktvun6u66dgpjh3l4wxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetIdentityProviderAdvancedConfigurationResponseBodyAdvancedConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderAdvancedConfigurationResponseBodyAdvancedConfiguration) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderAdvancedConfigurationResponseBodyAdvancedConfiguration) GetDingtalkAdvancedConfig() *GetIdentityProviderAdvancedConfigurationResponseBodyAdvancedConfigurationDingtalkAdvancedConfig {
	return s.DingtalkAdvancedConfig
}

func (s *GetIdentityProviderAdvancedConfigurationResponseBodyAdvancedConfiguration) GetIdentityProviderId() *string {
	return s.IdentityProviderId
}

func (s *GetIdentityProviderAdvancedConfigurationResponseBodyAdvancedConfiguration) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetIdentityProviderAdvancedConfigurationResponseBodyAdvancedConfiguration) SetDingtalkAdvancedConfig(v *GetIdentityProviderAdvancedConfigurationResponseBodyAdvancedConfigurationDingtalkAdvancedConfig) *GetIdentityProviderAdvancedConfigurationResponseBodyAdvancedConfiguration {
	s.DingtalkAdvancedConfig = v
	return s
}

func (s *GetIdentityProviderAdvancedConfigurationResponseBodyAdvancedConfiguration) SetIdentityProviderId(v string) *GetIdentityProviderAdvancedConfigurationResponseBodyAdvancedConfiguration {
	s.IdentityProviderId = &v
	return s
}

func (s *GetIdentityProviderAdvancedConfigurationResponseBodyAdvancedConfiguration) SetInstanceId(v string) *GetIdentityProviderAdvancedConfigurationResponseBodyAdvancedConfiguration {
	s.InstanceId = &v
	return s
}

func (s *GetIdentityProviderAdvancedConfigurationResponseBodyAdvancedConfiguration) Validate() error {
	if s.DingtalkAdvancedConfig != nil {
		if err := s.DingtalkAdvancedConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetIdentityProviderAdvancedConfigurationResponseBodyAdvancedConfigurationDingtalkAdvancedConfig struct {
	// IDaaS EIAM 钉钉一方应用的AppKey
	//
	// example:
	//
	// ding5xo9rg0csw3f6xxx
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// IDaaS EIAM 钉钉一方应用的AppSecret
	//
	// example:
	//
	// ***
	AppSecret *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
}

func (s GetIdentityProviderAdvancedConfigurationResponseBodyAdvancedConfigurationDingtalkAdvancedConfig) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderAdvancedConfigurationResponseBodyAdvancedConfigurationDingtalkAdvancedConfig) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderAdvancedConfigurationResponseBodyAdvancedConfigurationDingtalkAdvancedConfig) GetAppKey() *string {
	return s.AppKey
}

func (s *GetIdentityProviderAdvancedConfigurationResponseBodyAdvancedConfigurationDingtalkAdvancedConfig) GetAppSecret() *string {
	return s.AppSecret
}

func (s *GetIdentityProviderAdvancedConfigurationResponseBodyAdvancedConfigurationDingtalkAdvancedConfig) SetAppKey(v string) *GetIdentityProviderAdvancedConfigurationResponseBodyAdvancedConfigurationDingtalkAdvancedConfig {
	s.AppKey = &v
	return s
}

func (s *GetIdentityProviderAdvancedConfigurationResponseBodyAdvancedConfigurationDingtalkAdvancedConfig) SetAppSecret(v string) *GetIdentityProviderAdvancedConfigurationResponseBodyAdvancedConfigurationDingtalkAdvancedConfig {
	s.AppSecret = &v
	return s
}

func (s *GetIdentityProviderAdvancedConfigurationResponseBodyAdvancedConfigurationDingtalkAdvancedConfig) Validate() error {
	return dara.Validate(s)
}
