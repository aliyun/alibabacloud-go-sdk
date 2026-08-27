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
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The list of queried Logtail configurations.
	Configs []*ListAvailableConfigsResponseBodyConfigs `json:"configs,omitempty" xml:"configs,omitempty" type:"Repeated"`
	// The status code description.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
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
	// The enterprise ID.
	//
	// example:
	//
	// exampleCorpId
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// The organization name.
	//
	// example:
	//
	// string_value
	CorpName *string `json:"corpName,omitempty" xml:"corpName,omitempty"`
	// The platform type.
	//
	// example:
	//
	// string_value
	PlatformType *string `json:"platformType,omitempty" xml:"platformType,omitempty"`
	// The SSO configuration ID. This field has a value only for SAML, OAuth2, or WeCom types. The value is null for custom types.
	//
	// example:
	//
	// exampleSsoSettingsId
	SsoSettingsId *string `json:"ssoSettingsId,omitempty" xml:"ssoSettingsId,omitempty"`
	// The SSO configuration name. This field has a value only for SAML, OAuth2, or WeCom types. The value is null for custom types.
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
