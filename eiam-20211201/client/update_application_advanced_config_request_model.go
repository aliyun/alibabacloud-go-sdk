// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationAdvancedConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *UpdateApplicationAdvancedConfigRequest
	GetApplicationId() *string
	SetInstanceId(v string) *UpdateApplicationAdvancedConfigRequest
	GetInstanceId() *string
	SetScimServerAdvancedConfig(v *UpdateApplicationAdvancedConfigRequestScimServerAdvancedConfig) *UpdateApplicationAdvancedConfigRequest
	GetScimServerAdvancedConfig() *UpdateApplicationAdvancedConfigRequestScimServerAdvancedConfig
}

type UpdateApplicationAdvancedConfigRequest struct {
	// IDaaS的应用资源ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Scim Server 高阶配置
	ScimServerAdvancedConfig *UpdateApplicationAdvancedConfigRequestScimServerAdvancedConfig `json:"ScimServerAdvancedConfig,omitempty" xml:"ScimServerAdvancedConfig,omitempty" type:"Struct"`
}

func (s UpdateApplicationAdvancedConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationAdvancedConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationAdvancedConfigRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpdateApplicationAdvancedConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateApplicationAdvancedConfigRequest) GetScimServerAdvancedConfig() *UpdateApplicationAdvancedConfigRequestScimServerAdvancedConfig {
	return s.ScimServerAdvancedConfig
}

func (s *UpdateApplicationAdvancedConfigRequest) SetApplicationId(v string) *UpdateApplicationAdvancedConfigRequest {
	s.ApplicationId = &v
	return s
}

func (s *UpdateApplicationAdvancedConfigRequest) SetInstanceId(v string) *UpdateApplicationAdvancedConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateApplicationAdvancedConfigRequest) SetScimServerAdvancedConfig(v *UpdateApplicationAdvancedConfigRequestScimServerAdvancedConfig) *UpdateApplicationAdvancedConfigRequest {
	s.ScimServerAdvancedConfig = v
	return s
}

func (s *UpdateApplicationAdvancedConfigRequest) Validate() error {
	if s.ScimServerAdvancedConfig != nil {
		if err := s.ScimServerAdvancedConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateApplicationAdvancedConfigRequestScimServerAdvancedConfig struct {
	// 支持的用户自定义字段ID列表。
	SupportedUserCustomFieldIds []*string `json:"SupportedUserCustomFieldIds,omitempty" xml:"SupportedUserCustomFieldIds,omitempty" type:"Repeated"`
	// 用户扩展字段的命名空间。
	//
	// example:
	//
	// urn:ietf:params:scim:schemas:extension:customfield:2.0:User
	UserCustomFieldNamespace *string `json:"UserCustomFieldNamespace,omitempty" xml:"UserCustomFieldNamespace,omitempty"`
}

func (s UpdateApplicationAdvancedConfigRequestScimServerAdvancedConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationAdvancedConfigRequestScimServerAdvancedConfig) GoString() string {
	return s.String()
}

func (s *UpdateApplicationAdvancedConfigRequestScimServerAdvancedConfig) GetSupportedUserCustomFieldIds() []*string {
	return s.SupportedUserCustomFieldIds
}

func (s *UpdateApplicationAdvancedConfigRequestScimServerAdvancedConfig) GetUserCustomFieldNamespace() *string {
	return s.UserCustomFieldNamespace
}

func (s *UpdateApplicationAdvancedConfigRequestScimServerAdvancedConfig) SetSupportedUserCustomFieldIds(v []*string) *UpdateApplicationAdvancedConfigRequestScimServerAdvancedConfig {
	s.SupportedUserCustomFieldIds = v
	return s
}

func (s *UpdateApplicationAdvancedConfigRequestScimServerAdvancedConfig) SetUserCustomFieldNamespace(v string) *UpdateApplicationAdvancedConfigRequestScimServerAdvancedConfig {
	s.UserCustomFieldNamespace = &v
	return s
}

func (s *UpdateApplicationAdvancedConfigRequestScimServerAdvancedConfig) Validate() error {
	return dara.Validate(s)
}
