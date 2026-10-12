// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchExportHttpApisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiIds(v []*string) *BatchExportHttpApisRequest
	GetApiIds() []*string
	SetApiType(v string) *BatchExportHttpApisRequest
	GetApiType() *string
	SetExtensionConfig(v *BatchExportHttpApisRequestExtensionConfig) *BatchExportHttpApisRequest
	GetExtensionConfig() *BatchExportHttpApisRequestExtensionConfig
	SetFormat(v string) *BatchExportHttpApisRequest
	GetFormat() *string
	SetGatewayId(v string) *BatchExportHttpApisRequest
	GetGatewayId() *string
}

type BatchExportHttpApisRequest struct {
	// The list of API IDs to export.
	//
	// This parameter is required.
	ApiIds []*string `json:"apiIds,omitempty" xml:"apiIds,omitempty" type:"Repeated"`
	// The API type.
	//
	// This parameter is required.
	//
	// example:
	//
	// Http
	ApiType *string `json:"apiType,omitempty" xml:"apiType,omitempty"`
	// The export extension configuration. When this configuration is specified for a REST API, gatewayId is required. This configuration is required for batch export of HTTP APIs.
	ExtensionConfig *BatchExportHttpApisRequestExtensionConfig `json:"extensionConfig,omitempty" xml:"extensionConfig,omitempty" type:"Struct"`
	// The export format. Valid values: json and yaml. Default value: yaml.
	//
	// example:
	//
	// yaml
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// The gateway instance ID. This parameter is required when extensionConfig is specified for a REST API. This parameter is required for batch export of HTTP APIs.
	//
	// example:
	//
	// gw-xxx
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
}

func (s BatchExportHttpApisRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchExportHttpApisRequest) GoString() string {
	return s.String()
}

func (s *BatchExportHttpApisRequest) GetApiIds() []*string {
	return s.ApiIds
}

func (s *BatchExportHttpApisRequest) GetApiType() *string {
	return s.ApiType
}

func (s *BatchExportHttpApisRequest) GetExtensionConfig() *BatchExportHttpApisRequestExtensionConfig {
	return s.ExtensionConfig
}

func (s *BatchExportHttpApisRequest) GetFormat() *string {
	return s.Format
}

func (s *BatchExportHttpApisRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *BatchExportHttpApisRequest) SetApiIds(v []*string) *BatchExportHttpApisRequest {
	s.ApiIds = v
	return s
}

func (s *BatchExportHttpApisRequest) SetApiType(v string) *BatchExportHttpApisRequest {
	s.ApiType = &v
	return s
}

func (s *BatchExportHttpApisRequest) SetExtensionConfig(v *BatchExportHttpApisRequestExtensionConfig) *BatchExportHttpApisRequest {
	s.ExtensionConfig = v
	return s
}

func (s *BatchExportHttpApisRequest) SetFormat(v string) *BatchExportHttpApisRequest {
	s.Format = &v
	return s
}

func (s *BatchExportHttpApisRequest) SetGatewayId(v string) *BatchExportHttpApisRequest {
	s.GatewayId = &v
	return s
}

func (s *BatchExportHttpApisRequest) Validate() error {
	if s.ExtensionConfig != nil {
		if err := s.ExtensionConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchExportHttpApisRequestExtensionConfig struct {
	// Specifies whether to export the consumer authorization configuration.
	//
	// example:
	//
	// false
	WithAuthConfig *bool `json:"withAuthConfig,omitempty" xml:"withAuthConfig,omitempty"`
	// Specifies whether to export the authorized consumers.
	//
	// example:
	//
	// false
	WithAuthConsumer *bool `json:"withAuthConsumer,omitempty" xml:"withAuthConsumer,omitempty"`
	// Specifies whether to export the plugin configuration.
	//
	// example:
	//
	// false
	WithPlugin *bool `json:"withPlugin,omitempty" xml:"withPlugin,omitempty"`
	// Specifies whether to export the policy configuration.
	//
	// example:
	//
	// false
	WithPolicy *bool `json:"withPolicy,omitempty" xml:"withPolicy,omitempty"`
	// Specifies whether to export the backend service.
	//
	// example:
	//
	// false
	WithService *bool `json:"withService,omitempty" xml:"withService,omitempty"`
}

func (s BatchExportHttpApisRequestExtensionConfig) String() string {
	return dara.Prettify(s)
}

func (s BatchExportHttpApisRequestExtensionConfig) GoString() string {
	return s.String()
}

func (s *BatchExportHttpApisRequestExtensionConfig) GetWithAuthConfig() *bool {
	return s.WithAuthConfig
}

func (s *BatchExportHttpApisRequestExtensionConfig) GetWithAuthConsumer() *bool {
	return s.WithAuthConsumer
}

func (s *BatchExportHttpApisRequestExtensionConfig) GetWithPlugin() *bool {
	return s.WithPlugin
}

func (s *BatchExportHttpApisRequestExtensionConfig) GetWithPolicy() *bool {
	return s.WithPolicy
}

func (s *BatchExportHttpApisRequestExtensionConfig) GetWithService() *bool {
	return s.WithService
}

func (s *BatchExportHttpApisRequestExtensionConfig) SetWithAuthConfig(v bool) *BatchExportHttpApisRequestExtensionConfig {
	s.WithAuthConfig = &v
	return s
}

func (s *BatchExportHttpApisRequestExtensionConfig) SetWithAuthConsumer(v bool) *BatchExportHttpApisRequestExtensionConfig {
	s.WithAuthConsumer = &v
	return s
}

func (s *BatchExportHttpApisRequestExtensionConfig) SetWithPlugin(v bool) *BatchExportHttpApisRequestExtensionConfig {
	s.WithPlugin = &v
	return s
}

func (s *BatchExportHttpApisRequestExtensionConfig) SetWithPolicy(v bool) *BatchExportHttpApisRequestExtensionConfig {
	s.WithPolicy = &v
	return s
}

func (s *BatchExportHttpApisRequestExtensionConfig) SetWithService(v bool) *BatchExportHttpApisRequestExtensionConfig {
	s.WithService = &v
	return s
}

func (s *BatchExportHttpApisRequestExtensionConfig) Validate() error {
	return dara.Validate(s)
}
