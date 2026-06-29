// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportHttpApiRequest interface {
  dara.Model
  String() string
  GoString() string
  SetExtensionConfig(v *ExportHttpApiRequestExtensionConfig) *ExportHttpApiRequest
  GetExtensionConfig() *ExportHttpApiRequestExtensionConfig 
  SetGatewayId(v string) *ExportHttpApiRequest
  GetGatewayId() *string 
  SetOperationIds(v []*string) *ExportHttpApiRequest
  GetOperationIds() []*string 
}

type ExportHttpApiRequest struct {
  // Specifies whether to export plug-in configurations.
  ExtensionConfig *ExportHttpApiRequestExtensionConfig `json:"extensionConfig,omitempty" xml:"extensionConfig,omitempty" type:"Struct"`
  // The gateway instance ID. This parameter is required when you export gateway extension information.
  // 
  // example:
  // 
  // gw-xxx
  GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
  // The IDs of specific operations to export when you export a REST API.
  OperationIds []*string `json:"operationIds,omitempty" xml:"operationIds,omitempty" type:"Repeated"`
}

func (s ExportHttpApiRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportHttpApiRequest) GoString() string {
  return s.String()
}

func (s *ExportHttpApiRequest) GetExtensionConfig() *ExportHttpApiRequestExtensionConfig  {
  return s.ExtensionConfig
}

func (s *ExportHttpApiRequest) GetGatewayId() *string  {
  return s.GatewayId
}

func (s *ExportHttpApiRequest) GetOperationIds() []*string  {
  return s.OperationIds
}

func (s *ExportHttpApiRequest) SetExtensionConfig(v *ExportHttpApiRequestExtensionConfig) *ExportHttpApiRequest {
  s.ExtensionConfig = v
  return s
}

func (s *ExportHttpApiRequest) SetGatewayId(v string) *ExportHttpApiRequest {
  s.GatewayId = &v
  return s
}

func (s *ExportHttpApiRequest) SetOperationIds(v []*string) *ExportHttpApiRequest {
  s.OperationIds = v
  return s
}

func (s *ExportHttpApiRequest) Validate() error {
  if s.ExtensionConfig != nil {
    if err := s.ExtensionConfig.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExportHttpApiRequestExtensionConfig struct {
  // Specifies whether to export consumer authorization configurations for operations or routes.
  WithAuthConfig *bool `json:"withAuthConfig,omitempty" xml:"withAuthConfig,omitempty"`
  // Specifies whether to export authorized consumers.
  WithAuthConsumer *bool `json:"withAuthConsumer,omitempty" xml:"withAuthConsumer,omitempty"`
  // Specifies whether to export plug-in configurations.
  WithPlugin *bool `json:"withPlugin,omitempty" xml:"withPlugin,omitempty"`
  // Specifies whether to export policy configurations.
  WithPolicy *bool `json:"withPolicy,omitempty" xml:"withPolicy,omitempty"`
  // Specifies whether to export backend services.
  WithService *bool `json:"withService,omitempty" xml:"withService,omitempty"`
}

func (s ExportHttpApiRequestExtensionConfig) String() string {
  return dara.Prettify(s)
}

func (s ExportHttpApiRequestExtensionConfig) GoString() string {
  return s.String()
}

func (s *ExportHttpApiRequestExtensionConfig) GetWithAuthConfig() *bool  {
  return s.WithAuthConfig
}

func (s *ExportHttpApiRequestExtensionConfig) GetWithAuthConsumer() *bool  {
  return s.WithAuthConsumer
}

func (s *ExportHttpApiRequestExtensionConfig) GetWithPlugin() *bool  {
  return s.WithPlugin
}

func (s *ExportHttpApiRequestExtensionConfig) GetWithPolicy() *bool  {
  return s.WithPolicy
}

func (s *ExportHttpApiRequestExtensionConfig) GetWithService() *bool  {
  return s.WithService
}

func (s *ExportHttpApiRequestExtensionConfig) SetWithAuthConfig(v bool) *ExportHttpApiRequestExtensionConfig {
  s.WithAuthConfig = &v
  return s
}

func (s *ExportHttpApiRequestExtensionConfig) SetWithAuthConsumer(v bool) *ExportHttpApiRequestExtensionConfig {
  s.WithAuthConsumer = &v
  return s
}

func (s *ExportHttpApiRequestExtensionConfig) SetWithPlugin(v bool) *ExportHttpApiRequestExtensionConfig {
  s.WithPlugin = &v
  return s
}

func (s *ExportHttpApiRequestExtensionConfig) SetWithPolicy(v bool) *ExportHttpApiRequestExtensionConfig {
  s.WithPolicy = &v
  return s
}

func (s *ExportHttpApiRequestExtensionConfig) SetWithService(v bool) *ExportHttpApiRequestExtensionConfig {
  s.WithService = &v
  return s
}

func (s *ExportHttpApiRequestExtensionConfig) Validate() error {
  return dara.Validate(s)
}

