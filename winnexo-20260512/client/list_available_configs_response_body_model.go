// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAvailableConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAvailableConfigsResponseBody
	GetCode() *string
	SetConfigs(v []*ListAvailableConfigsResponseBodyConfigs) *ListAvailableConfigsResponseBody
	GetConfigs() []*ListAvailableConfigsResponseBodyConfigs
	SetMessage(v string) *ListAvailableConfigsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAvailableConfigsResponseBody
	GetRequestId() *string
}

type ListAvailableConfigsResponseBody struct {
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 可用的组织同步配置列表
	Configs []*ListAvailableConfigsResponseBodyConfigs `json:"configs,omitempty" xml:"configs,omitempty" type:"Repeated"`
	// 错误描述，成功时为空
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListAvailableConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAvailableConfigsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAvailableConfigsResponseBody) GetConfigs() []*ListAvailableConfigsResponseBodyConfigs {
	return s.Configs
}

func (s *ListAvailableConfigsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAvailableConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAvailableConfigsResponseBody) SetCode(v string) *ListAvailableConfigsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAvailableConfigsResponseBody) SetConfigs(v []*ListAvailableConfigsResponseBodyConfigs) *ListAvailableConfigsResponseBody {
	s.Configs = v
	return s
}

func (s *ListAvailableConfigsResponseBody) SetMessage(v string) *ListAvailableConfigsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAvailableConfigsResponseBody) SetRequestId(v string) *ListAvailableConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAvailableConfigsResponseBody) Validate() error {
	if s.Configs != nil {
		for _, item := range s.Configs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAvailableConfigsResponseBodyConfigs struct {
	// 企业标识（wecom=corpId, saml=idpEntityId, oauth2=clientId, custom=客户自定义）。注意：OAuth2 多 IdP 配置使用相同 clientId 时，需在 syncOrgStructure 中显式传 ssoSettingsId
	//
	// example:
	//
	// exampleCorpId
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 企业展示名称
	//
	// example:
	//
	// string_value
	CorpName *string `json:"corpName,omitempty" xml:"corpName,omitempty"`
	// 平台类型: wecom / saml / oauth2 / custom
	//
	// example:
	//
	// string_value
	PlatformType *string `json:"platformType,omitempty" xml:"platformType,omitempty"`
	// SSO 配置 ID（仅 SAML/OAuth2/WeCom 有值，custom 为 null）
	//
	// example:
	//
	// exampleSsoSettingsId
	SsoSettingsId *string `json:"ssoSettingsId,omitempty" xml:"ssoSettingsId,omitempty"`
	// SSO 配置名称（仅 SAML/OAuth2/WeCom 有值，custom 为 null）
	//
	// example:
	//
	// string_value
	SsoSettingsName *string `json:"ssoSettingsName,omitempty" xml:"ssoSettingsName,omitempty"`
}

func (s ListAvailableConfigsResponseBodyConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableConfigsResponseBodyConfigs) GoString() string {
	return s.String()
}

func (s *ListAvailableConfigsResponseBodyConfigs) GetCorpId() *string {
	return s.CorpId
}

func (s *ListAvailableConfigsResponseBodyConfigs) GetCorpName() *string {
	return s.CorpName
}

func (s *ListAvailableConfigsResponseBodyConfigs) GetPlatformType() *string {
	return s.PlatformType
}

func (s *ListAvailableConfigsResponseBodyConfigs) GetSsoSettingsId() *string {
	return s.SsoSettingsId
}

func (s *ListAvailableConfigsResponseBodyConfigs) GetSsoSettingsName() *string {
	return s.SsoSettingsName
}

func (s *ListAvailableConfigsResponseBodyConfigs) SetCorpId(v string) *ListAvailableConfigsResponseBodyConfigs {
	s.CorpId = &v
	return s
}

func (s *ListAvailableConfigsResponseBodyConfigs) SetCorpName(v string) *ListAvailableConfigsResponseBodyConfigs {
	s.CorpName = &v
	return s
}

func (s *ListAvailableConfigsResponseBodyConfigs) SetPlatformType(v string) *ListAvailableConfigsResponseBodyConfigs {
	s.PlatformType = &v
	return s
}

func (s *ListAvailableConfigsResponseBodyConfigs) SetSsoSettingsId(v string) *ListAvailableConfigsResponseBodyConfigs {
	s.SsoSettingsId = &v
	return s
}

func (s *ListAvailableConfigsResponseBodyConfigs) SetSsoSettingsName(v string) *ListAvailableConfigsResponseBodyConfigs {
	s.SsoSettingsName = &v
	return s
}

func (s *ListAvailableConfigsResponseBodyConfigs) Validate() error {
	return dara.Validate(s)
}
