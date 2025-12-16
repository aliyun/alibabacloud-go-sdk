// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationAdvancedConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationAdvancedConfig(v *GetApplicationAdvancedConfigResponseBodyApplicationAdvancedConfig) *GetApplicationAdvancedConfigResponseBody
	GetApplicationAdvancedConfig() *GetApplicationAdvancedConfigResponseBodyApplicationAdvancedConfig
	SetRequestId(v string) *GetApplicationAdvancedConfigResponseBody
	GetRequestId() *string
}

type GetApplicationAdvancedConfigResponseBody struct {
	ApplicationAdvancedConfig *GetApplicationAdvancedConfigResponseBodyApplicationAdvancedConfig `json:"ApplicationAdvancedConfig,omitempty" xml:"ApplicationAdvancedConfig,omitempty" type:"Struct"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApplicationAdvancedConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationAdvancedConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationAdvancedConfigResponseBody) GetApplicationAdvancedConfig() *GetApplicationAdvancedConfigResponseBodyApplicationAdvancedConfig {
	return s.ApplicationAdvancedConfig
}

func (s *GetApplicationAdvancedConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApplicationAdvancedConfigResponseBody) SetApplicationAdvancedConfig(v *GetApplicationAdvancedConfigResponseBodyApplicationAdvancedConfig) *GetApplicationAdvancedConfigResponseBody {
	s.ApplicationAdvancedConfig = v
	return s
}

func (s *GetApplicationAdvancedConfigResponseBody) SetRequestId(v string) *GetApplicationAdvancedConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApplicationAdvancedConfigResponseBody) Validate() error {
	if s.ApplicationAdvancedConfig != nil {
		if err := s.ApplicationAdvancedConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApplicationAdvancedConfigResponseBodyApplicationAdvancedConfig struct {
	// IDaaS EIAM 应用Id
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// IDaaS EIAM 实例Id
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Scim Server 高阶配置
	ScimServerAdvancedConfig *GetApplicationAdvancedConfigResponseBodyApplicationAdvancedConfigScimServerAdvancedConfig `json:"ScimServerAdvancedConfig,omitempty" xml:"ScimServerAdvancedConfig,omitempty" type:"Struct"`
}

func (s GetApplicationAdvancedConfigResponseBodyApplicationAdvancedConfig) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationAdvancedConfigResponseBodyApplicationAdvancedConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationAdvancedConfigResponseBodyApplicationAdvancedConfig) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *GetApplicationAdvancedConfigResponseBodyApplicationAdvancedConfig) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetApplicationAdvancedConfigResponseBodyApplicationAdvancedConfig) GetScimServerAdvancedConfig() *GetApplicationAdvancedConfigResponseBodyApplicationAdvancedConfigScimServerAdvancedConfig {
	return s.ScimServerAdvancedConfig
}

func (s *GetApplicationAdvancedConfigResponseBodyApplicationAdvancedConfig) SetApplicationId(v string) *GetApplicationAdvancedConfigResponseBodyApplicationAdvancedConfig {
	s.ApplicationId = &v
	return s
}

func (s *GetApplicationAdvancedConfigResponseBodyApplicationAdvancedConfig) SetInstanceId(v string) *GetApplicationAdvancedConfigResponseBodyApplicationAdvancedConfig {
	s.InstanceId = &v
	return s
}

func (s *GetApplicationAdvancedConfigResponseBodyApplicationAdvancedConfig) SetScimServerAdvancedConfig(v *GetApplicationAdvancedConfigResponseBodyApplicationAdvancedConfigScimServerAdvancedConfig) *GetApplicationAdvancedConfigResponseBodyApplicationAdvancedConfig {
	s.ScimServerAdvancedConfig = v
	return s
}

func (s *GetApplicationAdvancedConfigResponseBodyApplicationAdvancedConfig) Validate() error {
	if s.ScimServerAdvancedConfig != nil {
		if err := s.ScimServerAdvancedConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApplicationAdvancedConfigResponseBodyApplicationAdvancedConfigScimServerAdvancedConfig struct {
	// 支持的用户自定义字段ID列表。
	SupportedUserCustomFieldIds []*string `json:"SupportedUserCustomFieldIds,omitempty" xml:"SupportedUserCustomFieldIds,omitempty" type:"Repeated"`
	// 用户扩展字段的命名空间。
	//
	// example:
	//
	// urn:ietf:params:scim:schemas:extension:customfield:2.0:User
	UserCustomFieldNamespace *string `json:"UserCustomFieldNamespace,omitempty" xml:"UserCustomFieldNamespace,omitempty"`
}

func (s GetApplicationAdvancedConfigResponseBodyApplicationAdvancedConfigScimServerAdvancedConfig) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationAdvancedConfigResponseBodyApplicationAdvancedConfigScimServerAdvancedConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationAdvancedConfigResponseBodyApplicationAdvancedConfigScimServerAdvancedConfig) GetSupportedUserCustomFieldIds() []*string {
	return s.SupportedUserCustomFieldIds
}

func (s *GetApplicationAdvancedConfigResponseBodyApplicationAdvancedConfigScimServerAdvancedConfig) GetUserCustomFieldNamespace() *string {
	return s.UserCustomFieldNamespace
}

func (s *GetApplicationAdvancedConfigResponseBodyApplicationAdvancedConfigScimServerAdvancedConfig) SetSupportedUserCustomFieldIds(v []*string) *GetApplicationAdvancedConfigResponseBodyApplicationAdvancedConfigScimServerAdvancedConfig {
	s.SupportedUserCustomFieldIds = v
	return s
}

func (s *GetApplicationAdvancedConfigResponseBodyApplicationAdvancedConfigScimServerAdvancedConfig) SetUserCustomFieldNamespace(v string) *GetApplicationAdvancedConfigResponseBodyApplicationAdvancedConfigScimServerAdvancedConfig {
	s.UserCustomFieldNamespace = &v
	return s
}

func (s *GetApplicationAdvancedConfigResponseBodyApplicationAdvancedConfigScimServerAdvancedConfig) Validate() error {
	return dara.Validate(s)
}
