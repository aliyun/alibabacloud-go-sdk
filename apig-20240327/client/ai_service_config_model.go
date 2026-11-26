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
	SetApiKeyFailover(v *AiServiceConfigApiKeyFailover) *AiServiceConfig
	GetApiKeyFailover() *AiServiceConfigApiKeyFailover
	SetApiKeyReferences(v []*string) *AiServiceConfig
	GetApiKeyReferences() []*string
	SetApiKeys(v []*string) *AiServiceConfig
	GetApiKeys() []*string
	SetAuthHeaderKey(v string) *AiServiceConfig
	GetAuthHeaderKey() *string
	SetBedrockServiceConfig(v *AiServiceConfigBedrockServiceConfig) *AiServiceConfig
	GetBedrockServiceConfig() *AiServiceConfigBedrockServiceConfig
	SetClaudeCodeMode(v bool) *AiServiceConfig
	GetClaudeCodeMode() *bool
	SetCompatibleProtocols(v []*string) *AiServiceConfig
	GetCompatibleProtocols() []*string
	SetConfigRevision(v string) *AiServiceConfig
	GetConfigRevision() *string
	SetDefaultModelName(v string) *AiServiceConfig
	GetDefaultModelName() *string
	SetEnableHealthCheck(v bool) *AiServiceConfig
	GetEnableHealthCheck() *bool
	SetEnableOutlierDetection(v bool) *AiServiceConfig
	GetEnableOutlierDetection() *bool
	SetKlingServiceConfig(v *AiServiceConfigKlingServiceConfig) *AiServiceConfig
	GetKlingServiceConfig() *AiServiceConfigKlingServiceConfig
	SetPaiEASServiceConfig(v *AiServiceConfigPaiEASServiceConfig) *AiServiceConfig
	GetPaiEASServiceConfig() *AiServiceConfigPaiEASServiceConfig
	SetProtocols(v []*string) *AiServiceConfig
	GetProtocols() []*string
	SetProvider(v string) *AiServiceConfig
	GetProvider() *string
	SetRetryOnFailure(v *AiServiceConfigRetryOnFailure) *AiServiceConfig
	GetRetryOnFailure() *AiServiceConfigRetryOnFailure
	SetUpstreamAuthMode(v string) *AiServiceConfig
	GetUpstreamAuthMode() *string
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
	// The configuration for API key failure ejection and recovery.
	ApiKeyFailover *AiServiceConfigApiKeyFailover `json:"apiKeyFailover,omitempty" xml:"apiKeyFailover,omitempty" type:"Struct"`
	// The read-only summary of KMS reference names, excluding credential values.
	ApiKeyReferences []*string `json:"apiKeyReferences,omitempty" xml:"apiKeyReferences,omitempty" type:"Repeated"`
	// The list of API keys.
	ApiKeys []*string `json:"apiKeys,omitempty" xml:"apiKeys,omitempty" type:"Repeated"`
	// The name of the request header used to pass authentication credentials to the upstream.
	//
	// example:
	//
	// Authorization
	AuthHeaderKey *string `json:"authHeaderKey,omitempty" xml:"authHeaderKey,omitempty"`
	// The Bedrock service-specific configuration.
	BedrockServiceConfig *AiServiceConfigBedrockServiceConfig `json:"bedrockServiceConfig,omitempty" xml:"bedrockServiceConfig,omitempty" type:"Struct"`
	// Specifies whether to enable Claude Code compatibility mode.
	//
	// example:
	//
	// false
	ClaudeCodeMode *bool `json:"claudeCodeMode,omitempty" xml:"claudeCodeMode,omitempty"`
	// The list of compatible protocols supported by the service.
	CompatibleProtocols []*string `json:"compatibleProtocols,omitempty" xml:"compatibleProtocols,omitempty" type:"Repeated"`
	// The read-only service configuration version. Use expectedConfigRevision for conditional updates.
	//
	// example:
	//
	// 7e4f52a9b0c34d16a268e913f5c2087b
	ConfigRevision *string `json:"configRevision,omitempty" xml:"configRevision,omitempty"`
	// The name of the default model.
	//
	// example:
	//
	// gpt-4
	DefaultModelName *string `json:"defaultModelName,omitempty" xml:"defaultModelName,omitempty"`
	// Specifies whether to enable health checks.
	//
	// example:
	//
	// true
	EnableHealthCheck *bool `json:"enableHealthCheck,omitempty" xml:"enableHealthCheck,omitempty"`
	// Specifies whether to enable the outlier detection feature.
	//
	// example:
	//
	// true
	EnableOutlierDetection *bool `json:"enableOutlierDetection,omitempty" xml:"enableOutlierDetection,omitempty"`
	// The authentication and token configuration specific to the KlingAI service.
	KlingServiceConfig *AiServiceConfigKlingServiceConfig `json:"klingServiceConfig,omitempty" xml:"klingServiceConfig,omitempty" type:"Struct"`
	// The PAI-EAS service-specific configuration.
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
	// The configuration for retrying with a different key after the current request fails.
	RetryOnFailure *AiServiceConfigRetryOnFailure `json:"retryOnFailure,omitempty" xml:"retryOnFailure,omitempty" type:"Struct"`
	// The upstream authentication mode, which distinguishes between API key authentication, no authentication, and pass-through downstream authentication.
	//
	// example:
	//
	// api-key
	UpstreamAuthMode *string `json:"upstreamAuthMode,omitempty" xml:"upstreamAuthMode,omitempty"`
	// The configuration information for the Vertex AI service.
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

func (s *AiServiceConfig) GetApiKeyFailover() *AiServiceConfigApiKeyFailover {
	return s.ApiKeyFailover
}

func (s *AiServiceConfig) GetApiKeyReferences() []*string {
	return s.ApiKeyReferences
}

func (s *AiServiceConfig) GetApiKeys() []*string {
	return s.ApiKeys
}

func (s *AiServiceConfig) GetAuthHeaderKey() *string {
	return s.AuthHeaderKey
}

func (s *AiServiceConfig) GetBedrockServiceConfig() *AiServiceConfigBedrockServiceConfig {
	return s.BedrockServiceConfig
}

func (s *AiServiceConfig) GetClaudeCodeMode() *bool {
	return s.ClaudeCodeMode
}

func (s *AiServiceConfig) GetCompatibleProtocols() []*string {
	return s.CompatibleProtocols
}

func (s *AiServiceConfig) GetConfigRevision() *string {
	return s.ConfigRevision
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

func (s *AiServiceConfig) GetKlingServiceConfig() *AiServiceConfigKlingServiceConfig {
	return s.KlingServiceConfig
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

func (s *AiServiceConfig) GetRetryOnFailure() *AiServiceConfigRetryOnFailure {
	return s.RetryOnFailure
}

func (s *AiServiceConfig) GetUpstreamAuthMode() *string {
	return s.UpstreamAuthMode
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

func (s *AiServiceConfig) SetApiKeyFailover(v *AiServiceConfigApiKeyFailover) *AiServiceConfig {
	s.ApiKeyFailover = v
	return s
}

func (s *AiServiceConfig) SetApiKeyReferences(v []*string) *AiServiceConfig {
	s.ApiKeyReferences = v
	return s
}

func (s *AiServiceConfig) SetApiKeys(v []*string) *AiServiceConfig {
	s.ApiKeys = v
	return s
}

func (s *AiServiceConfig) SetAuthHeaderKey(v string) *AiServiceConfig {
	s.AuthHeaderKey = &v
	return s
}

func (s *AiServiceConfig) SetBedrockServiceConfig(v *AiServiceConfigBedrockServiceConfig) *AiServiceConfig {
	s.BedrockServiceConfig = v
	return s
}

func (s *AiServiceConfig) SetClaudeCodeMode(v bool) *AiServiceConfig {
	s.ClaudeCodeMode = &v
	return s
}

func (s *AiServiceConfig) SetCompatibleProtocols(v []*string) *AiServiceConfig {
	s.CompatibleProtocols = v
	return s
}

func (s *AiServiceConfig) SetConfigRevision(v string) *AiServiceConfig {
	s.ConfigRevision = &v
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

func (s *AiServiceConfig) SetKlingServiceConfig(v *AiServiceConfigKlingServiceConfig) *AiServiceConfig {
	s.KlingServiceConfig = v
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

func (s *AiServiceConfig) SetRetryOnFailure(v *AiServiceConfigRetryOnFailure) *AiServiceConfig {
	s.RetryOnFailure = v
	return s
}

func (s *AiServiceConfig) SetUpstreamAuthMode(v string) *AiServiceConfig {
	s.UpstreamAuthMode = &v
	return s
}

func (s *AiServiceConfig) SetVertexServiceConfig(v *AiServiceConfigVertexServiceConfig) *AiServiceConfig {
	s.VertexServiceConfig = v
	return s
}

func (s *AiServiceConfig) Validate() error {
	if s.ApiKeyFailover != nil {
		if err := s.ApiKeyFailover.Validate(); err != nil {
			return err
		}
	}
	if s.BedrockServiceConfig != nil {
		if err := s.BedrockServiceConfig.Validate(); err != nil {
			return err
		}
	}
	if s.KlingServiceConfig != nil {
		if err := s.KlingServiceConfig.Validate(); err != nil {
			return err
		}
	}
	if s.PaiEASServiceConfig != nil {
		if err := s.PaiEASServiceConfig.Validate(); err != nil {
			return err
		}
	}
	if s.RetryOnFailure != nil {
		if err := s.RetryOnFailure.Validate(); err != nil {
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

type AiServiceConfigApiKeyFailover struct {
	// The cooldown duration after an API key is ejected.
	//
	// example:
	//
	// 30000
	CooldownDuration *int64 `json:"cooldownDuration,omitempty" xml:"cooldownDuration,omitempty"`
	// Specifies whether to enable API key failure ejection.
	//
	// example:
	//
	// false
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The set of response status matching rules that trigger API key failure ejection.
	FailoverOnStatus []*string `json:"failoverOnStatus,omitempty" xml:"failoverOnStatus,omitempty" type:"Repeated"`
	// The failure count threshold that triggers API key ejection.
	//
	// example:
	//
	// 3
	FailureThreshold *int64 `json:"failureThreshold,omitempty" xml:"failureThreshold,omitempty"`
}

func (s AiServiceConfigApiKeyFailover) String() string {
	return dara.Prettify(s)
}

func (s AiServiceConfigApiKeyFailover) GoString() string {
	return s.String()
}

func (s *AiServiceConfigApiKeyFailover) GetCooldownDuration() *int64 {
	return s.CooldownDuration
}

func (s *AiServiceConfigApiKeyFailover) GetEnabled() *bool {
	return s.Enabled
}

func (s *AiServiceConfigApiKeyFailover) GetFailoverOnStatus() []*string {
	return s.FailoverOnStatus
}

func (s *AiServiceConfigApiKeyFailover) GetFailureThreshold() *int64 {
	return s.FailureThreshold
}

func (s *AiServiceConfigApiKeyFailover) SetCooldownDuration(v int64) *AiServiceConfigApiKeyFailover {
	s.CooldownDuration = &v
	return s
}

func (s *AiServiceConfigApiKeyFailover) SetEnabled(v bool) *AiServiceConfigApiKeyFailover {
	s.Enabled = &v
	return s
}

func (s *AiServiceConfigApiKeyFailover) SetFailoverOnStatus(v []*string) *AiServiceConfigApiKeyFailover {
	s.FailoverOnStatus = v
	return s
}

func (s *AiServiceConfigApiKeyFailover) SetFailureThreshold(v int64) *AiServiceConfigApiKeyFailover {
	s.FailureThreshold = &v
	return s
}

func (s *AiServiceConfigApiKeyFailover) Validate() error {
	return dara.Validate(s)
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
	// Configures whether to insert Bedrock cache markers based on prompt positions.
	//
	// example:
	//
	// {"systemPrompt":true}
	BedrockPromptCachePointPositions map[string]*bool `json:"bedrockPromptCachePointPositions,omitempty" xml:"bedrockPromptCachePointPositions,omitempty"`
	// The Bedrock service endpoint mode, which supports runtime and mantle. If omitted, runtime is used.
	//
	// example:
	//
	// runtime
	EndpointMode *string `json:"endpointMode,omitempty" xml:"endpointMode,omitempty"`
	// Specifies whether to enable prompt caching for the Bedrock upstream.
	//
	// example:
	//
	// true
	PromptCacheEnable *bool `json:"promptCacheEnable,omitempty" xml:"promptCacheEnable,omitempty"`
	// The data retention policy for the prompt cache.
	//
	// example:
	//
	// in_memory
	PromptCacheRetention *string `json:"promptCacheRetention,omitempty" xml:"promptCacheRetention,omitempty"`
	// The applicable model strategy for Bedrock prompt caching.
	//
	// example:
	//
	// supported_models_only
	PromptCacheStrategy *string `json:"promptCacheStrategy,omitempty" xml:"promptCacheStrategy,omitempty"`
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

func (s *AiServiceConfigBedrockServiceConfig) GetBedrockPromptCachePointPositions() map[string]*bool {
	return s.BedrockPromptCachePointPositions
}

func (s *AiServiceConfigBedrockServiceConfig) GetEndpointMode() *string {
	return s.EndpointMode
}

func (s *AiServiceConfigBedrockServiceConfig) GetPromptCacheEnable() *bool {
	return s.PromptCacheEnable
}

func (s *AiServiceConfigBedrockServiceConfig) GetPromptCacheRetention() *string {
	return s.PromptCacheRetention
}

func (s *AiServiceConfigBedrockServiceConfig) GetPromptCacheStrategy() *string {
	return s.PromptCacheStrategy
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

func (s *AiServiceConfigBedrockServiceConfig) SetBedrockPromptCachePointPositions(v map[string]*bool) *AiServiceConfigBedrockServiceConfig {
	s.BedrockPromptCachePointPositions = v
	return s
}

func (s *AiServiceConfigBedrockServiceConfig) SetEndpointMode(v string) *AiServiceConfigBedrockServiceConfig {
	s.EndpointMode = &v
	return s
}

func (s *AiServiceConfigBedrockServiceConfig) SetPromptCacheEnable(v bool) *AiServiceConfigBedrockServiceConfig {
	s.PromptCacheEnable = &v
	return s
}

func (s *AiServiceConfigBedrockServiceConfig) SetPromptCacheRetention(v string) *AiServiceConfigBedrockServiceConfig {
	s.PromptCacheRetention = &v
	return s
}

func (s *AiServiceConfigBedrockServiceConfig) SetPromptCacheStrategy(v string) *AiServiceConfigBedrockServiceConfig {
	s.PromptCacheStrategy = &v
	return s
}

func (s *AiServiceConfigBedrockServiceConfig) Validate() error {
	return dara.Validate(s)
}

type AiServiceConfigKlingServiceConfig struct {
	// The KlingAI access key identifier. Examples use desensitized values.
	//
	// example:
	//
	// kling-access-key-****
	KlingAccessKey *string `json:"klingAccessKey,omitempty" xml:"klingAccessKey,omitempty"`
	// The KlingAI secret key. Examples use desensitized values.
	//
	// example:
	//
	// kling-secret-key-****
	KlingSecretKey *string `json:"klingSecretKey,omitempty" xml:"klingSecretKey,omitempty"`
	// The advance refresh configuration for KlingAI tokens.
	//
	// example:
	//
	// 120
	KlingTokenRefreshAhead *int64 `json:"klingTokenRefreshAhead,omitempty" xml:"klingTokenRefreshAhead,omitempty"`
}

func (s AiServiceConfigKlingServiceConfig) String() string {
	return dara.Prettify(s)
}

func (s AiServiceConfigKlingServiceConfig) GoString() string {
	return s.String()
}

func (s *AiServiceConfigKlingServiceConfig) GetKlingAccessKey() *string {
	return s.KlingAccessKey
}

func (s *AiServiceConfigKlingServiceConfig) GetKlingSecretKey() *string {
	return s.KlingSecretKey
}

func (s *AiServiceConfigKlingServiceConfig) GetKlingTokenRefreshAhead() *int64 {
	return s.KlingTokenRefreshAhead
}

func (s *AiServiceConfigKlingServiceConfig) SetKlingAccessKey(v string) *AiServiceConfigKlingServiceConfig {
	s.KlingAccessKey = &v
	return s
}

func (s *AiServiceConfigKlingServiceConfig) SetKlingSecretKey(v string) *AiServiceConfigKlingServiceConfig {
	s.KlingSecretKey = &v
	return s
}

func (s *AiServiceConfigKlingServiceConfig) SetKlingTokenRefreshAhead(v int64) *AiServiceConfigKlingServiceConfig {
	s.KlingTokenRefreshAhead = &v
	return s
}

func (s *AiServiceConfigKlingServiceConfig) Validate() error {
	return dara.Validate(s)
}

type AiServiceConfigPaiEASServiceConfig struct {
	// The connection type, which can be a private network (VPC) or a public network.
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

type AiServiceConfigRetryOnFailure struct {
	// Specifies whether to enable retries upon current request failure.
	//
	// example:
	//
	// false
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The maximum number of retries for the current request.
	//
	// example:
	//
	// 2
	MaxRetries *int64 `json:"maxRetries,omitempty" xml:"maxRetries,omitempty"`
	// The set of response status matching rules that trigger retries for the current request.
	RetryOnStatus []*string `json:"retryOnStatus,omitempty" xml:"retryOnStatus,omitempty" type:"Repeated"`
	// The retry timeout configuration for the current request.
	//
	// example:
	//
	// 30000
	RetryTimeout *int64 `json:"retryTimeout,omitempty" xml:"retryTimeout,omitempty"`
}

func (s AiServiceConfigRetryOnFailure) String() string {
	return dara.Prettify(s)
}

func (s AiServiceConfigRetryOnFailure) GoString() string {
	return s.String()
}

func (s *AiServiceConfigRetryOnFailure) GetEnabled() *bool {
	return s.Enabled
}

func (s *AiServiceConfigRetryOnFailure) GetMaxRetries() *int64 {
	return s.MaxRetries
}

func (s *AiServiceConfigRetryOnFailure) GetRetryOnStatus() []*string {
	return s.RetryOnStatus
}

func (s *AiServiceConfigRetryOnFailure) GetRetryTimeout() *int64 {
	return s.RetryTimeout
}

func (s *AiServiceConfigRetryOnFailure) SetEnabled(v bool) *AiServiceConfigRetryOnFailure {
	s.Enabled = &v
	return s
}

func (s *AiServiceConfigRetryOnFailure) SetMaxRetries(v int64) *AiServiceConfigRetryOnFailure {
	s.MaxRetries = &v
	return s
}

func (s *AiServiceConfigRetryOnFailure) SetRetryOnStatus(v []*string) *AiServiceConfigRetryOnFailure {
	s.RetryOnStatus = v
	return s
}

func (s *AiServiceConfigRetryOnFailure) SetRetryTimeout(v int64) *AiServiceConfigRetryOnFailure {
	s.RetryTimeout = &v
	return s
}

func (s *AiServiceConfigRetryOnFailure) Validate() error {
	return dara.Validate(s)
}

type AiServiceConfigVertexServiceConfig struct {
	// The safety filtering level settings for the Gemini model.
	GeminiSafetySetting map[string]*string `json:"geminiSafetySetting,omitempty" xml:"geminiSafetySetting,omitempty"`
	// The authentication key for Vertex AI.
	//
	// example:
	//
	// {\"type\":\"service_account\",\"project_id\":\"my-project\",\"private_key\":\"-----BEGIN PRIVATE KEY-----...\",\"client_email\":\"xxx@xxx.iam.gserviceaccount.com\",\"token_uri\":\"https://oauth2.googleapis.com/token\"}
	VertexAuthKey *string `json:"vertexAuthKey,omitempty" xml:"vertexAuthKey,omitempty"`
	// The authentication service name used by Vertex AI.
	//
	// example:
	//
	// google-oauth2
	VertexAuthServiceName *string `json:"vertexAuthServiceName,omitempty" xml:"vertexAuthServiceName,omitempty"`
	// Specifies whether to use the OpenAI-compatible interface of Vertex AI.
	//
	// example:
	//
	// false
	VertexOpenAICompatible *bool `json:"vertexOpenAICompatible,omitempty" xml:"vertexOpenAICompatible,omitempty"`
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
	// The number of seconds before access token expiration to trigger an advance refresh.
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

func (s *AiServiceConfigVertexServiceConfig) GetVertexOpenAICompatible() *bool {
	return s.VertexOpenAICompatible
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

func (s *AiServiceConfigVertexServiceConfig) SetVertexOpenAICompatible(v bool) *AiServiceConfigVertexServiceConfig {
	s.VertexOpenAICompatible = &v
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
