// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiServiceConfig interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyGenerateMode(v string) *AiServiceConfig
	GetApiKeyGenerateMode() *string
	SetAddress(v string) *AiServiceConfig
	GetAddress() *string
	SetApiKeys(v []*string) *AiServiceConfig
	GetApiKeys() []*string
	SetBedrockServiceConfig(v *AiServiceConfigBedrockServiceConfig) *AiServiceConfig
	GetBedrockServiceConfig() *AiServiceConfigBedrockServiceConfig
	SetCompatibleProtocols(v []*string) *AiServiceConfig
	GetCompatibleProtocols() []*string
	SetDefaultModelName(v string) *AiServiceConfig
	GetDefaultModelName() *string
	SetEnableHealthCheck(v bool) *AiServiceConfig
	GetEnableHealthCheck() *bool
	SetEnableOutlierDetection(v bool) *AiServiceConfig
	GetEnableOutlierDetection() *bool
	SetPaiEASServiceConfig(v *AiServiceConfigPaiEASServiceConfig) *AiServiceConfig
	GetPaiEASServiceConfig() *AiServiceConfigPaiEASServiceConfig
	SetProtocols(v []*string) *AiServiceConfig
	GetProtocols() []*string
	SetProvider(v string) *AiServiceConfig
	GetProvider() *string
	SetVertexServiceConfig(v *AiServiceConfigVertexServiceConfig) *AiServiceConfig
	GetVertexServiceConfig() *AiServiceConfigVertexServiceConfig
}

type AiServiceConfig struct {
	// The generation mode of the API key.
	//
	// example:
	//
	// Custom
	ApiKeyGenerateMode *string `json:"ApiKeyGenerateMode,omitempty" xml:"ApiKeyGenerateMode,omitempty"`
	// The service address.
	//
	// example:
	//
	// https://dashscope.aliyun.com
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// The list of API keys.
	ApiKeys []*string `json:"apiKeys,omitempty" xml:"apiKeys,omitempty" type:"Repeated"`
	// The configuration specific to the Bedrock service.
	BedrockServiceConfig *AiServiceConfigBedrockServiceConfig `json:"bedrockServiceConfig,omitempty" xml:"bedrockServiceConfig,omitempty" type:"Struct"`
	// The list of compatible protocols supported by the service.
	CompatibleProtocols []*string `json:"compatibleProtocols,omitempty" xml:"compatibleProtocols,omitempty" type:"Repeated"`
	// The name of the default model.
	//
	// example:
	//
	// gpt-4
	DefaultModelName *string `json:"defaultModelName,omitempty" xml:"defaultModelName,omitempty"`
	// Indicates whether health check is enabled.
	//
	// example:
	//
	// true
	EnableHealthCheck *bool `json:"enableHealthCheck,omitempty" xml:"enableHealthCheck,omitempty"`
	// Indicates whether outlier detection is enabled.
	//
	// example:
	//
	// true
	EnableOutlierDetection *bool `json:"enableOutlierDetection,omitempty" xml:"enableOutlierDetection,omitempty"`
	// The configuration specific to the PAI-EAS service.
	PaiEASServiceConfig *AiServiceConfigPaiEASServiceConfig `json:"paiEASServiceConfig,omitempty" xml:"paiEASServiceConfig,omitempty" type:"Struct"`
	// The list of model protocols.
	Protocols []*string `json:"protocols,omitempty" xml:"protocols,omitempty" type:"Repeated"`
	// The large language model provider.
	//
	// example:
	//
	// qwen
	//
	// moonshot
	//
	// baichuan
	//
	// yi
	//
	// zhipuai
	//
	// 360ai
	//
	// hunyuan
	//
	// stepfun
	//
	// spark
	//
	// openai
	//
	// claude
	//
	// anthropic
	//
	// doubao
	//
	// minimax
	//
	// gemini
	//
	// azure
	//
	// deepseek
	//
	// custom
	//
	// pai-eas
	//
	// bedrock
	//
	// vertex
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
	// The configuration of the Vertex AI service.
	VertexServiceConfig *AiServiceConfigVertexServiceConfig `json:"vertexServiceConfig,omitempty" xml:"vertexServiceConfig,omitempty" type:"Struct"`
}

func (s AiServiceConfig) String() string {
	return dara.Prettify(s)
}

func (s AiServiceConfig) GoString() string {
	return s.String()
}

func (s *AiServiceConfig) GetApiKeyGenerateMode() *string {
	return s.ApiKeyGenerateMode
}

func (s *AiServiceConfig) GetAddress() *string {
	return s.Address
}

func (s *AiServiceConfig) GetApiKeys() []*string {
	return s.ApiKeys
}

func (s *AiServiceConfig) GetBedrockServiceConfig() *AiServiceConfigBedrockServiceConfig {
	return s.BedrockServiceConfig
}

func (s *AiServiceConfig) GetCompatibleProtocols() []*string {
	return s.CompatibleProtocols
}

func (s *AiServiceConfig) GetDefaultModelName() *string {
	return s.DefaultModelName
}

func (s *AiServiceConfig) GetEnableHealthCheck() *bool {
	return s.EnableHealthCheck
}

func (s *AiServiceConfig) GetEnableOutlierDetection() *bool {
	return s.EnableOutlierDetection
}

func (s *AiServiceConfig) GetPaiEASServiceConfig() *AiServiceConfigPaiEASServiceConfig {
	return s.PaiEASServiceConfig
}

func (s *AiServiceConfig) GetProtocols() []*string {
	return s.Protocols
}

func (s *AiServiceConfig) GetProvider() *string {
	return s.Provider
}

func (s *AiServiceConfig) GetVertexServiceConfig() *AiServiceConfigVertexServiceConfig {
	return s.VertexServiceConfig
}

func (s *AiServiceConfig) SetApiKeyGenerateMode(v string) *AiServiceConfig {
	s.ApiKeyGenerateMode = &v
	return s
}

func (s *AiServiceConfig) SetAddress(v string) *AiServiceConfig {
	s.Address = &v
	return s
}

func (s *AiServiceConfig) SetApiKeys(v []*string) *AiServiceConfig {
	s.ApiKeys = v
	return s
}

func (s *AiServiceConfig) SetBedrockServiceConfig(v *AiServiceConfigBedrockServiceConfig) *AiServiceConfig {
	s.BedrockServiceConfig = v
	return s
}

func (s *AiServiceConfig) SetCompatibleProtocols(v []*string) *AiServiceConfig {
	s.CompatibleProtocols = v
	return s
}

func (s *AiServiceConfig) SetDefaultModelName(v string) *AiServiceConfig {
	s.DefaultModelName = &v
	return s
}

func (s *AiServiceConfig) SetEnableHealthCheck(v bool) *AiServiceConfig {
	s.EnableHealthCheck = &v
	return s
}

func (s *AiServiceConfig) SetEnableOutlierDetection(v bool) *AiServiceConfig {
	s.EnableOutlierDetection = &v
	return s
}

func (s *AiServiceConfig) SetPaiEASServiceConfig(v *AiServiceConfigPaiEASServiceConfig) *AiServiceConfig {
	s.PaiEASServiceConfig = v
	return s
}

func (s *AiServiceConfig) SetProtocols(v []*string) *AiServiceConfig {
	s.Protocols = v
	return s
}

func (s *AiServiceConfig) SetProvider(v string) *AiServiceConfig {
	s.Provider = &v
	return s
}

func (s *AiServiceConfig) SetVertexServiceConfig(v *AiServiceConfigVertexServiceConfig) *AiServiceConfig {
	s.VertexServiceConfig = v
	return s
}

func (s *AiServiceConfig) Validate() error {
	if s.BedrockServiceConfig != nil {
		if err := s.BedrockServiceConfig.Validate(); err != nil {
			return err
		}
	}
	if s.PaiEASServiceConfig != nil {
		if err := s.PaiEASServiceConfig.Validate(); err != nil {
			return err
		}
	}
	if s.VertexServiceConfig != nil {
		if err := s.VertexServiceConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AiServiceConfigBedrockServiceConfig struct {
	// The AWS access key.
	//
	// example:
	//
	// AKIAIOSFODNN7EXAMPLE
	AwsAccessKey *string `json:"awsAccessKey,omitempty" xml:"awsAccessKey,omitempty"`
	// The AWS region.
	//
	// example:
	//
	// us-east-1
	AwsRegion *string `json:"awsRegion,omitempty" xml:"awsRegion,omitempty"`
	// The AWS secret key.
	//
	// example:
	//
	// wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
	AwsSecretKey *string `json:"awsSecretKey,omitempty" xml:"awsSecretKey,omitempty"`
}

func (s AiServiceConfigBedrockServiceConfig) String() string {
	return dara.Prettify(s)
}

func (s AiServiceConfigBedrockServiceConfig) GoString() string {
	return s.String()
}

func (s *AiServiceConfigBedrockServiceConfig) GetAwsAccessKey() *string {
	return s.AwsAccessKey
}

func (s *AiServiceConfigBedrockServiceConfig) GetAwsRegion() *string {
	return s.AwsRegion
}

func (s *AiServiceConfigBedrockServiceConfig) GetAwsSecretKey() *string {
	return s.AwsSecretKey
}

func (s *AiServiceConfigBedrockServiceConfig) SetAwsAccessKey(v string) *AiServiceConfigBedrockServiceConfig {
	s.AwsAccessKey = &v
	return s
}

func (s *AiServiceConfigBedrockServiceConfig) SetAwsRegion(v string) *AiServiceConfigBedrockServiceConfig {
	s.AwsRegion = &v
	return s
}

func (s *AiServiceConfigBedrockServiceConfig) SetAwsSecretKey(v string) *AiServiceConfigBedrockServiceConfig {
	s.AwsSecretKey = &v
	return s
}

func (s *AiServiceConfigBedrockServiceConfig) Validate() error {
	return dara.Validate(s)
}

type AiServiceConfigPaiEASServiceConfig struct {
	// The connection type: private network (vpc) or public network (public).
	//
	// example:
	//
	// "public"，"private"
	EndpointType *string `json:"endpointType,omitempty" xml:"endpointType,omitempty"`
	// The service ID.
	//
	// example:
	//
	// svc-cvaeb0um1hkqp79k72v0
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// The service name.
	//
	// example:
	//
	// auto-ai-service
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// cb89d4de30042df4459442589664e00d
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s AiServiceConfigPaiEASServiceConfig) String() string {
	return dara.Prettify(s)
}

func (s AiServiceConfigPaiEASServiceConfig) GoString() string {
	return s.String()
}

func (s *AiServiceConfigPaiEASServiceConfig) GetEndpointType() *string {
	return s.EndpointType
}

func (s *AiServiceConfigPaiEASServiceConfig) GetServiceId() *string {
	return s.ServiceId
}

func (s *AiServiceConfigPaiEASServiceConfig) GetServiceName() *string {
	return s.ServiceName
}

func (s *AiServiceConfigPaiEASServiceConfig) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *AiServiceConfigPaiEASServiceConfig) SetEndpointType(v string) *AiServiceConfigPaiEASServiceConfig {
	s.EndpointType = &v
	return s
}

func (s *AiServiceConfigPaiEASServiceConfig) SetServiceId(v string) *AiServiceConfigPaiEASServiceConfig {
	s.ServiceId = &v
	return s
}

func (s *AiServiceConfigPaiEASServiceConfig) SetServiceName(v string) *AiServiceConfigPaiEASServiceConfig {
	s.ServiceName = &v
	return s
}

func (s *AiServiceConfigPaiEASServiceConfig) SetWorkspaceId(v string) *AiServiceConfigPaiEASServiceConfig {
	s.WorkspaceId = &v
	return s
}

func (s *AiServiceConfigPaiEASServiceConfig) Validate() error {
	return dara.Validate(s)
}

type AiServiceConfigVertexServiceConfig struct {
	// The safety filtering level settings for the Gemini model.
	GeminiSafetySetting map[string]*string `json:"geminiSafetySetting,omitempty" xml:"geminiSafetySetting,omitempty"`
	// The authentication key for Vertex AI.
	//
	// example:
	//
	// {\\"type\\":\\"service_account\\",\\"project_id\\":\\"my-project\\",\\"private_key\\":\\"-----BEGIN PRIVATE KEY-----...\\",\\"client_email\\":\\"xxx@xxx.iam.gserviceaccount.com\\",\\"token_uri\\":\\"https://oauth2.googleapis.com/token\\"}
	VertexAuthKey *string `json:"vertexAuthKey,omitempty" xml:"vertexAuthKey,omitempty"`
	// The service name of the authentication service used by Vertex AI.
	//
	// example:
	//
	// google-oauth2
	VertexAuthServiceName *string `json:"vertexAuthServiceName,omitempty" xml:"vertexAuthServiceName,omitempty"`
	// The Google Cloud project ID.
	//
	// example:
	//
	// my-project
	VertexProjectId *string `json:"vertexProjectId,omitempty" xml:"vertexProjectId,omitempty"`
	// The Google Cloud region where the Vertex AI service is located.
	//
	// example:
	//
	// us-central1
	VertexRegion *string `json:"vertexRegion,omitempty" xml:"vertexRegion,omitempty"`
	// The number of seconds before the access token expires to refresh it in advance.
	//
	// example:
	//
	// 300
	VertexTokenRefreshAhead *int32 `json:"vertexTokenRefreshAhead,omitempty" xml:"vertexTokenRefreshAhead,omitempty"`
}

func (s AiServiceConfigVertexServiceConfig) String() string {
	return dara.Prettify(s)
}

func (s AiServiceConfigVertexServiceConfig) GoString() string {
	return s.String()
}

func (s *AiServiceConfigVertexServiceConfig) GetGeminiSafetySetting() map[string]*string {
	return s.GeminiSafetySetting
}

func (s *AiServiceConfigVertexServiceConfig) GetVertexAuthKey() *string {
	return s.VertexAuthKey
}

func (s *AiServiceConfigVertexServiceConfig) GetVertexAuthServiceName() *string {
	return s.VertexAuthServiceName
}

func (s *AiServiceConfigVertexServiceConfig) GetVertexProjectId() *string {
	return s.VertexProjectId
}

func (s *AiServiceConfigVertexServiceConfig) GetVertexRegion() *string {
	return s.VertexRegion
}

func (s *AiServiceConfigVertexServiceConfig) GetVertexTokenRefreshAhead() *int32 {
	return s.VertexTokenRefreshAhead
}

func (s *AiServiceConfigVertexServiceConfig) SetGeminiSafetySetting(v map[string]*string) *AiServiceConfigVertexServiceConfig {
	s.GeminiSafetySetting = v
	return s
}

func (s *AiServiceConfigVertexServiceConfig) SetVertexAuthKey(v string) *AiServiceConfigVertexServiceConfig {
	s.VertexAuthKey = &v
	return s
}

func (s *AiServiceConfigVertexServiceConfig) SetVertexAuthServiceName(v string) *AiServiceConfigVertexServiceConfig {
	s.VertexAuthServiceName = &v
	return s
}

func (s *AiServiceConfigVertexServiceConfig) SetVertexProjectId(v string) *AiServiceConfigVertexServiceConfig {
	s.VertexProjectId = &v
	return s
}

func (s *AiServiceConfigVertexServiceConfig) SetVertexRegion(v string) *AiServiceConfigVertexServiceConfig {
	s.VertexRegion = &v
	return s
}

func (s *AiServiceConfigVertexServiceConfig) SetVertexTokenRefreshAhead(v int32) *AiServiceConfigVertexServiceConfig {
	s.VertexTokenRefreshAhead = &v
	return s
}

func (s *AiServiceConfigVertexServiceConfig) Validate() error {
	return dara.Validate(s)
}
