// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConsumerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetConsumerResponseBody
	GetCode() *string
	SetData(v *GetConsumerResponseBodyData) *GetConsumerResponseBody
	GetData() *GetConsumerResponseBodyData
	SetMessage(v string) *GetConsumerResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetConsumerResponseBody
	GetRequestId() *string
}

type GetConsumerResponseBody struct {
	// O código de resposta.
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Os dados de resposta.
	Data *GetConsumerResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// A mensagem de resposta.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// O ID da solicitação.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetConsumerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerResponseBody) GoString() string {
	return s.String()
}

func (s *GetConsumerResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetConsumerResponseBody) GetData() *GetConsumerResponseBodyData {
	return s.Data
}

func (s *GetConsumerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetConsumerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetConsumerResponseBody) SetCode(v string) *GetConsumerResponseBody {
	s.Code = &v
	return s
}

func (s *GetConsumerResponseBody) SetData(v *GetConsumerResponseBodyData) *GetConsumerResponseBody {
	s.Data = v
	return s
}

func (s *GetConsumerResponseBody) SetMessage(v string) *GetConsumerResponseBody {
	s.Message = &v
	return s
}

func (s *GetConsumerResponseBody) SetRequestId(v string) *GetConsumerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConsumerResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetConsumerResponseBodyData struct {
	// As configurações de autenticação por par de AccessKey.
	AkSkIdentityConfigs []*AkSkIdentityConfig `json:"akSkIdentityConfigs,omitempty" xml:"akSkIdentityConfigs,omitempty" type:"Repeated"`
	// A configuração de autenticação de chave de API.
	ApiKeyIdentityConfig *ApiKeyIdentityConfig `json:"apiKeyIdentityConfig,omitempty" xml:"apiKeyIdentityConfig,omitempty"`
	// O ID do consumidor de API.
	ConsumerId *string `json:"consumerId,omitempty" xml:"consumerId,omitempty"`
	// O carimbo de data/hora de criação.
	CreateTimestamp *int64 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	// O status de publicação da API no ambiente atual.
	DeployStatus *string `json:"deployStatus,omitempty" xml:"deployStatus,omitempty"`
	// A descrição.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Indica se o consumidor de API está habilitado.
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// A configuração de autenticação JWT.
	JwtIdentityConfig *JwtIdentityConfig `json:"jwtIdentityConfig,omitempty" xml:"jwtIdentityConfig,omitempty"`
	// O nome do consumidor de API.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// O carimbo de data/hora de atualização.
	UpdateTimestamp *int64 `json:"updateTimestamp,omitempty" xml:"updateTimestamp,omitempty"`
}

func (s GetConsumerResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetConsumerResponseBodyData) GetAkSkIdentityConfigs() []*AkSkIdentityConfig {
	return s.AkSkIdentityConfigs
}

func (s *GetConsumerResponseBodyData) GetApiKeyIdentityConfig() *ApiKeyIdentityConfig {
	return s.ApiKeyIdentityConfig
}

func (s *GetConsumerResponseBodyData) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *GetConsumerResponseBodyData) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *GetConsumerResponseBodyData) GetDeployStatus() *string {
	return s.DeployStatus
}

func (s *GetConsumerResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetConsumerResponseBodyData) GetEnable() *bool {
	return s.Enable
}

func (s *GetConsumerResponseBodyData) GetJwtIdentityConfig() *JwtIdentityConfig {
	return s.JwtIdentityConfig
}

func (s *GetConsumerResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetConsumerResponseBodyData) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *GetConsumerResponseBodyData) SetAkSkIdentityConfigs(v []*AkSkIdentityConfig) *GetConsumerResponseBodyData {
	s.AkSkIdentityConfigs = v
	return s
}

func (s *GetConsumerResponseBodyData) SetApiKeyIdentityConfig(v *ApiKeyIdentityConfig) *GetConsumerResponseBodyData {
	s.ApiKeyIdentityConfig = v
	return s
}

func (s *GetConsumerResponseBodyData) SetConsumerId(v string) *GetConsumerResponseBodyData {
	s.ConsumerId = &v
	return s
}

func (s *GetConsumerResponseBodyData) SetCreateTimestamp(v int64) *GetConsumerResponseBodyData {
	s.CreateTimestamp = &v
	return s
}

func (s *GetConsumerResponseBodyData) SetDeployStatus(v string) *GetConsumerResponseBodyData {
	s.DeployStatus = &v
	return s
}

func (s *GetConsumerResponseBodyData) SetDescription(v string) *GetConsumerResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetConsumerResponseBodyData) SetEnable(v bool) *GetConsumerResponseBodyData {
	s.Enable = &v
	return s
}

func (s *GetConsumerResponseBodyData) SetJwtIdentityConfig(v *JwtIdentityConfig) *GetConsumerResponseBodyData {
	s.JwtIdentityConfig = v
	return s
}

func (s *GetConsumerResponseBodyData) SetName(v string) *GetConsumerResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetConsumerResponseBodyData) SetUpdateTimestamp(v int64) *GetConsumerResponseBodyData {
	s.UpdateTimestamp = &v
	return s
}

func (s *GetConsumerResponseBodyData) Validate() error {
	if s.AkSkIdentityConfigs != nil {
		for _, item := range s.AkSkIdentityConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ApiKeyIdentityConfig != nil {
		if err := s.ApiKeyIdentityConfig.Validate(); err != nil {
			return err
		}
	}
	if s.JwtIdentityConfig != nil {
		if err := s.JwtIdentityConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
