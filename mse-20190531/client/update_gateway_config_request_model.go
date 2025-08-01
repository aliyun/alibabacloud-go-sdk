// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateGatewayConfigRequest
	GetAcceptLanguage() *string
	SetConfigName(v string) *UpdateGatewayConfigRequest
	GetConfigName() *string
	SetConfigValue(v string) *UpdateGatewayConfigRequest
	GetConfigValue() *string
	SetGatewayUniqueId(v string) *UpdateGatewayConfigRequest
	GetGatewayUniqueId() *string
}

type UpdateGatewayConfigRequest struct {
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// EnableK8sSourceWorkloadFilter
	ConfigName *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"enable":true,"filterOpt":"EQUAL","labelKey":"key","labelValue":"value"}
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// gw-b525dc1adf3c486ab96224a6346*****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
}

func (s UpdateGatewayConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayConfigRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateGatewayConfigRequest) GetConfigName() *string {
	return s.ConfigName
}

func (s *UpdateGatewayConfigRequest) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *UpdateGatewayConfigRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdateGatewayConfigRequest) SetAcceptLanguage(v string) *UpdateGatewayConfigRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateGatewayConfigRequest) SetConfigName(v string) *UpdateGatewayConfigRequest {
	s.ConfigName = &v
	return s
}

func (s *UpdateGatewayConfigRequest) SetConfigValue(v string) *UpdateGatewayConfigRequest {
	s.ConfigValue = &v
	return s
}

func (s *UpdateGatewayConfigRequest) SetGatewayUniqueId(v string) *UpdateGatewayConfigRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateGatewayConfigRequest) Validate() error {
	return dara.Validate(s)
}
