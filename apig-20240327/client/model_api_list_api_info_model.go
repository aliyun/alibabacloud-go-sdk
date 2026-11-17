// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelApiListApiInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAgentProtocols(v []*string) *ModelApiListApiInfo
	GetAgentProtocols() []*string
	SetAiProtocols(v []*string) *ModelApiListApiInfo
	GetAiProtocols() []*string
	SetAuthConfig(v *AuthConfig) *ModelApiListApiInfo
	GetAuthConfig() *AuthConfig
	SetBasePath(v string) *ModelApiListApiInfo
	GetBasePath() *string
	SetCreationClientToken(v string) *ModelApiListApiInfo
	GetCreationClientToken() *string
	SetDeleteStatus(v string) *ModelApiListApiInfo
	GetDeleteStatus() *string
	SetDeployCntMap(v map[string]*ModelApiListApiInfoDeployCntMapValue) *ModelApiListApiInfo
	GetDeployCntMap() map[string]*ModelApiListApiInfoDeployCntMapValue
	SetDeployConfigs(v []*HttpApiDeployConfig) *ModelApiListApiInfo
	GetDeployConfigs() []*HttpApiDeployConfig
	SetDeployErrorMessage(v string) *ModelApiListApiInfo
	GetDeployErrorMessage() *string
	SetDescription(v string) *ModelApiListApiInfo
	GetDescription() *string
	SetDomainIds(v []*string) *ModelApiListApiInfo
	GetDomainIds() []*string
	SetDomainInfos(v []*ModelApiDomainInfo) *ModelApiListApiInfo
	GetDomainInfos() []*ModelApiDomainInfo
	SetEnableAuth(v bool) *ModelApiListApiInfo
	GetEnableAuth() *bool
	SetEnvironmentId(v string) *ModelApiListApiInfo
	GetEnvironmentId() *string
	SetEnvironments(v []*ModelApiListApiInfoEnvironments) *ModelApiListApiInfo
	GetEnvironments() []*ModelApiListApiInfoEnvironments
	SetGatewayId(v string) *ModelApiListApiInfo
	GetGatewayId() *string
	SetGatewayInfo(v *GatewayInfo) *ModelApiListApiInfo
	GetGatewayInfo() *GatewayInfo
	SetHttpApiId(v string) *ModelApiListApiInfo
	GetHttpApiId() *string
	SetIngressInfo(v *ModelApiListApiInfoIngressInfo) *ModelApiListApiInfo
	GetIngressInfo() *ModelApiListApiInfoIngressInfo
	SetModelApiInfo(v *ModelApiListApiInfoModelApiInfo) *ModelApiListApiInfo
	GetModelApiInfo() *ModelApiListApiInfoModelApiInfo
	SetModelCardRuntimeDiagnostics(v []*ModelApiListApiInfoModelCardRuntimeDiagnostics) *ModelApiListApiInfo
	GetModelCardRuntimeDiagnostics() []*ModelApiListApiInfoModelCardRuntimeDiagnostics
	SetModelCategory(v string) *ModelApiListApiInfo
	GetModelCategory() *string
	SetName(v string) *ModelApiListApiInfo
	GetName() *string
	SetPolicyConfigs(v []*ModelApiPolicyConfig) *ModelApiListApiInfo
	GetPolicyConfigs() []*ModelApiPolicyConfig
	SetProtocols(v []*string) *ModelApiListApiInfo
	GetProtocols() []*string
	SetResourceGroupId(v string) *ModelApiListApiInfo
	GetResourceGroupId() *string
	SetStatus(v string) *ModelApiListApiInfo
	GetStatus() *string
	SetType(v string) *ModelApiListApiInfo
	GetType() *string
	SetVersionInfo(v *HttpApiVersionInfo) *ModelApiListApiInfo
	GetVersionInfo() *HttpApiVersionInfo
}

type ModelApiListApiInfo struct {
	// The collection of agent protocols for the model API.
	AgentProtocols []*string `json:"agentProtocols,omitempty" xml:"agentProtocols,omitempty" type:"Repeated"`
	// The API protocols created by scenario. APIs created by selecting models provide open paths through modelApiInfo.
	AiProtocols []*string `json:"aiProtocols,omitempty" xml:"aiProtocols,omitempty" type:"Repeated"`
	// The API authentication configurations created by scenario.
	AuthConfig *AuthConfig `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	// The API base path.
	//
	// example:
	//
	// /v1
	BasePath *string `json:"basePath,omitempty" xml:"basePath,omitempty"`
	// The read-only association identifier for the creation request. It is returned only within approximately 24 hours after the first persistence and when the association with the same instance is unique after normal authentication. It is omitted when the identifier is historical, expired, or ambiguous. You cannot use this identifier to determine whether publishing or authorization is successful, nor does it replace the creation idempotency record.
	//
	// example:
	//
	// b1907ca4-952c-4c5f-8d1f-0aa6ae339c96
	CreationClientToken *string `json:"creationClientToken,omitempty" xml:"creationClientToken,omitempty"`
	// The optional read-only deletion status of the new model API. Deleting indicates that the deletion request is accepted but not completed. DeleteFailed indicates that the background deletion failed, and you can use the original clientToken to retry DeleteModelApi. This field is not returned for normal deployment and historical undeployment. The absence of this field does not indicate that the deletion is complete.
	//
	// example:
	//
	// Deleting
	DeleteStatus *string `json:"deleteStatus,omitempty" xml:"deleteStatus,omitempty"`
	// The publish status of the API operation.
	//
	// example:
	//
	// {"example-scope":{"deployedCnt":1,"cnt":1}}
	DeployCntMap map[string]*ModelApiListApiInfoDeployCntMapValue `json:"deployCntMap,omitempty" xml:"deployCntMap,omitempty"`
	// The API deployment configurations and associated services created by scenario, provided as read-only list information.
	DeployConfigs []*HttpApiDeployConfig `json:"deployConfigs,omitempty" xml:"deployConfigs,omitempty" type:"Repeated"`
	// The read-only reason for deployment or deletion failure.
	//
	// example:
	//
	// The model configuration is unavailable
	DeployErrorMessage *string `json:"deployErrorMessage,omitempty" xml:"deployErrorMessage,omitempty"`
	// The API description.
	//
	// example:
	//
	// Model API for business integration
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The selected domain IDs.
	DomainIds []*string `json:"domainIds,omitempty" xml:"domainIds,omitempty" type:"Repeated"`
	// The read-only domain details and default domain flag.
	DomainInfos []*ModelApiDomainInfo `json:"domainInfos,omitempty" xml:"domainInfos,omitempty" type:"Repeated"`
	// Specifies whether to enable authentication.
	//
	// example:
	//
	// true
	EnableAuth *bool `json:"enableAuth,omitempty" xml:"enableAuth,omitempty"`
	// The read-only default environment ID of the bound instance.
	//
	// example:
	//
	// env-d00000000000000000000
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// The publish environment context of the API.
	Environments []*ModelApiListApiInfoEnvironments `json:"environments,omitempty" xml:"environments,omitempty" type:"Repeated"`
	// The ID of the gateway instance to which the API belongs.
	//
	// example:
	//
	// gw-d00000000000000000000
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The read-only information of the bound instance.
	GatewayInfo *GatewayInfo `json:"gatewayInfo,omitempty" xml:"gatewayInfo,omitempty"`
	// The model API ID.
	//
	// example:
	//
	// api-xxx
	HttpApiId *string `json:"httpApiId,omitempty" xml:"httpApiId,omitempty"`
	// The HTTP ingress API information.
	IngressInfo *ModelApiListApiInfoIngressInfo `json:"ingressInfo,omitempty" xml:"ingressInfo,omitempty" type:"Struct"`
	// The public selection information of the model API.
	ModelApiInfo *ModelApiListApiInfoModelApiInfo `json:"modelApiInfo,omitempty" xml:"modelApiInfo,omitempty" type:"Struct"`
	// The read-only runtime diagnostics of the model card, derived from the original markers of the deployment configuration. This is used only for the readability and diagnosability of abnormal drafts.
	ModelCardRuntimeDiagnostics []*ModelApiListApiInfoModelCardRuntimeDiagnostics `json:"modelCardRuntimeDiagnostics,omitempty" xml:"modelCardRuntimeDiagnostics,omitempty" type:"Repeated"`
	// The model category of the model API.
	//
	// example:
	//
	// Text
	ModelCategory *string `json:"modelCategory,omitempty" xml:"modelCategory,omitempty"`
	// The API name.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The inline configurations for AI request logs and AI security protection.
	PolicyConfigs []*ModelApiPolicyConfig `json:"policyConfigs,omitempty" xml:"policyConfigs,omitempty" type:"Repeated"`
	// The list of API protocols.
	Protocols []*string `json:"protocols,omitempty" xml:"protocols,omitempty" type:"Repeated"`
	// The resource group ID.
	//
	// example:
	//
	// rg-xxx
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The deployment status on the bound instance. In the new version, Undeploying indicates that the deletion is in progress, and UndeployFailed indicates that the deletion failed. Refer to deleteStatus to determine whether the operation is a deletion.
	//
	// example:
	//
	// Deployed
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The API type.
	//
	// example:
	//
	// LLM
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The API versioning information.
	VersionInfo *HttpApiVersionInfo `json:"versionInfo,omitempty" xml:"versionInfo,omitempty"`
}

func (s ModelApiListApiInfo) String() string {
	return dara.Prettify(s)
}

func (s ModelApiListApiInfo) GoString() string {
	return s.String()
}

func (s *ModelApiListApiInfo) GetAgentProtocols() []*string {
	return s.AgentProtocols
}

func (s *ModelApiListApiInfo) GetAiProtocols() []*string {
	return s.AiProtocols
}

func (s *ModelApiListApiInfo) GetAuthConfig() *AuthConfig {
	return s.AuthConfig
}

func (s *ModelApiListApiInfo) GetBasePath() *string {
	return s.BasePath
}

func (s *ModelApiListApiInfo) GetCreationClientToken() *string {
	return s.CreationClientToken
}

func (s *ModelApiListApiInfo) GetDeleteStatus() *string {
	return s.DeleteStatus
}

func (s *ModelApiListApiInfo) GetDeployCntMap() map[string]*ModelApiListApiInfoDeployCntMapValue {
	return s.DeployCntMap
}

func (s *ModelApiListApiInfo) GetDeployConfigs() []*HttpApiDeployConfig {
	return s.DeployConfigs
}

func (s *ModelApiListApiInfo) GetDeployErrorMessage() *string {
	return s.DeployErrorMessage
}

func (s *ModelApiListApiInfo) GetDescription() *string {
	return s.Description
}

func (s *ModelApiListApiInfo) GetDomainIds() []*string {
	return s.DomainIds
}

func (s *ModelApiListApiInfo) GetDomainInfos() []*ModelApiDomainInfo {
	return s.DomainInfos
}

func (s *ModelApiListApiInfo) GetEnableAuth() *bool {
	return s.EnableAuth
}

func (s *ModelApiListApiInfo) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *ModelApiListApiInfo) GetEnvironments() []*ModelApiListApiInfoEnvironments {
	return s.Environments
}

func (s *ModelApiListApiInfo) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ModelApiListApiInfo) GetGatewayInfo() *GatewayInfo {
	return s.GatewayInfo
}

func (s *ModelApiListApiInfo) GetHttpApiId() *string {
	return s.HttpApiId
}

func (s *ModelApiListApiInfo) GetIngressInfo() *ModelApiListApiInfoIngressInfo {
	return s.IngressInfo
}

func (s *ModelApiListApiInfo) GetModelApiInfo() *ModelApiListApiInfoModelApiInfo {
	return s.ModelApiInfo
}

func (s *ModelApiListApiInfo) GetModelCardRuntimeDiagnostics() []*ModelApiListApiInfoModelCardRuntimeDiagnostics {
	return s.ModelCardRuntimeDiagnostics
}

func (s *ModelApiListApiInfo) GetModelCategory() *string {
	return s.ModelCategory
}

func (s *ModelApiListApiInfo) GetName() *string {
	return s.Name
}

func (s *ModelApiListApiInfo) GetPolicyConfigs() []*ModelApiPolicyConfig {
	return s.PolicyConfigs
}

func (s *ModelApiListApiInfo) GetProtocols() []*string {
	return s.Protocols
}

func (s *ModelApiListApiInfo) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModelApiListApiInfo) GetStatus() *string {
	return s.Status
}

func (s *ModelApiListApiInfo) GetType() *string {
	return s.Type
}

func (s *ModelApiListApiInfo) GetVersionInfo() *HttpApiVersionInfo {
	return s.VersionInfo
}

func (s *ModelApiListApiInfo) SetAgentProtocols(v []*string) *ModelApiListApiInfo {
	s.AgentProtocols = v
	return s
}

func (s *ModelApiListApiInfo) SetAiProtocols(v []*string) *ModelApiListApiInfo {
	s.AiProtocols = v
	return s
}

func (s *ModelApiListApiInfo) SetAuthConfig(v *AuthConfig) *ModelApiListApiInfo {
	s.AuthConfig = v
	return s
}

func (s *ModelApiListApiInfo) SetBasePath(v string) *ModelApiListApiInfo {
	s.BasePath = &v
	return s
}

func (s *ModelApiListApiInfo) SetCreationClientToken(v string) *ModelApiListApiInfo {
	s.CreationClientToken = &v
	return s
}

func (s *ModelApiListApiInfo) SetDeleteStatus(v string) *ModelApiListApiInfo {
	s.DeleteStatus = &v
	return s
}

func (s *ModelApiListApiInfo) SetDeployCntMap(v map[string]*ModelApiListApiInfoDeployCntMapValue) *ModelApiListApiInfo {
	s.DeployCntMap = v
	return s
}

func (s *ModelApiListApiInfo) SetDeployConfigs(v []*HttpApiDeployConfig) *ModelApiListApiInfo {
	s.DeployConfigs = v
	return s
}

func (s *ModelApiListApiInfo) SetDeployErrorMessage(v string) *ModelApiListApiInfo {
	s.DeployErrorMessage = &v
	return s
}

func (s *ModelApiListApiInfo) SetDescription(v string) *ModelApiListApiInfo {
	s.Description = &v
	return s
}

func (s *ModelApiListApiInfo) SetDomainIds(v []*string) *ModelApiListApiInfo {
	s.DomainIds = v
	return s
}

func (s *ModelApiListApiInfo) SetDomainInfos(v []*ModelApiDomainInfo) *ModelApiListApiInfo {
	s.DomainInfos = v
	return s
}

func (s *ModelApiListApiInfo) SetEnableAuth(v bool) *ModelApiListApiInfo {
	s.EnableAuth = &v
	return s
}

func (s *ModelApiListApiInfo) SetEnvironmentId(v string) *ModelApiListApiInfo {
	s.EnvironmentId = &v
	return s
}

func (s *ModelApiListApiInfo) SetEnvironments(v []*ModelApiListApiInfoEnvironments) *ModelApiListApiInfo {
	s.Environments = v
	return s
}

func (s *ModelApiListApiInfo) SetGatewayId(v string) *ModelApiListApiInfo {
	s.GatewayId = &v
	return s
}

func (s *ModelApiListApiInfo) SetGatewayInfo(v *GatewayInfo) *ModelApiListApiInfo {
	s.GatewayInfo = v
	return s
}

func (s *ModelApiListApiInfo) SetHttpApiId(v string) *ModelApiListApiInfo {
	s.HttpApiId = &v
	return s
}

func (s *ModelApiListApiInfo) SetIngressInfo(v *ModelApiListApiInfoIngressInfo) *ModelApiListApiInfo {
	s.IngressInfo = v
	return s
}

func (s *ModelApiListApiInfo) SetModelApiInfo(v *ModelApiListApiInfoModelApiInfo) *ModelApiListApiInfo {
	s.ModelApiInfo = v
	return s
}

func (s *ModelApiListApiInfo) SetModelCardRuntimeDiagnostics(v []*ModelApiListApiInfoModelCardRuntimeDiagnostics) *ModelApiListApiInfo {
	s.ModelCardRuntimeDiagnostics = v
	return s
}

func (s *ModelApiListApiInfo) SetModelCategory(v string) *ModelApiListApiInfo {
	s.ModelCategory = &v
	return s
}

func (s *ModelApiListApiInfo) SetName(v string) *ModelApiListApiInfo {
	s.Name = &v
	return s
}

func (s *ModelApiListApiInfo) SetPolicyConfigs(v []*ModelApiPolicyConfig) *ModelApiListApiInfo {
	s.PolicyConfigs = v
	return s
}

func (s *ModelApiListApiInfo) SetProtocols(v []*string) *ModelApiListApiInfo {
	s.Protocols = v
	return s
}

func (s *ModelApiListApiInfo) SetResourceGroupId(v string) *ModelApiListApiInfo {
	s.ResourceGroupId = &v
	return s
}

func (s *ModelApiListApiInfo) SetStatus(v string) *ModelApiListApiInfo {
	s.Status = &v
	return s
}

func (s *ModelApiListApiInfo) SetType(v string) *ModelApiListApiInfo {
	s.Type = &v
	return s
}

func (s *ModelApiListApiInfo) SetVersionInfo(v *HttpApiVersionInfo) *ModelApiListApiInfo {
	s.VersionInfo = v
	return s
}

func (s *ModelApiListApiInfo) Validate() error {
	if s.AuthConfig != nil {
		if err := s.AuthConfig.Validate(); err != nil {
			return err
		}
	}
	if s.DeployConfigs != nil {
		for _, item := range s.DeployConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DomainInfos != nil {
		for _, item := range s.DomainInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Environments != nil {
		for _, item := range s.Environments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.GatewayInfo != nil {
		if err := s.GatewayInfo.Validate(); err != nil {
			return err
		}
	}
	if s.IngressInfo != nil {
		if err := s.IngressInfo.Validate(); err != nil {
			return err
		}
	}
	if s.ModelApiInfo != nil {
		if err := s.ModelApiInfo.Validate(); err != nil {
			return err
		}
	}
	if s.ModelCardRuntimeDiagnostics != nil {
		for _, item := range s.ModelCardRuntimeDiagnostics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PolicyConfigs != nil {
		for _, item := range s.PolicyConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VersionInfo != nil {
		if err := s.VersionInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelApiListApiInfoEnvironments struct {
	// The alias of the environment.
	//
	// example:
	//
	// test
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// The release scenario.
	//
	// example:
	//
	// SingleService
	BackendScene *string `json:"backendScene,omitempty" xml:"backendScene,omitempty"`
	// The backend type.
	//
	// example:
	//
	// Service
	BackendType *string `json:"backendType,omitempty" xml:"backendType,omitempty"`
	// The custom domain name information published by the user.
	CustomDomains []*HttpApiDomainInfo `json:"customDomains,omitempty" xml:"customDomains,omitempty" type:"Repeated"`
	// The read-only reason for the release failure.
	//
	// example:
	//
	// The model configuration is unavailable
	DeployErrorMessage *string `json:"deployErrorMessage,omitempty" xml:"deployErrorMessage,omitempty"`
	// The release status of the API in the current environment.
	//
	// example:
	//
	// Deployed
	DeployStatus *string `json:"deployStatus,omitempty" xml:"deployStatus,omitempty"`
	// The ID of the environment.
	//
	// example:
	//
	// env-xxx
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// The gateway instance information to which the current environment belongs.
	GatewayInfo *ModelApiListApiInfoEnvironmentsGatewayInfo `json:"gatewayInfo,omitempty" xml:"gatewayInfo,omitempty" type:"Struct"`
	// The name of the environment.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The existing service configurations.
	ServiceConfigs []*ModelApiListApiInfoEnvironmentsServiceConfigs `json:"serviceConfigs,omitempty" xml:"serviceConfigs,omitempty" type:"Repeated"`
	// The list of default domain names for the environment.
	SubDomains []*ModelApiListApiInfoEnvironmentsSubDomains `json:"subDomains,omitempty" xml:"subDomains,omitempty" type:"Repeated"`
}

func (s ModelApiListApiInfoEnvironments) String() string {
	return dara.Prettify(s)
}

func (s ModelApiListApiInfoEnvironments) GoString() string {
	return s.String()
}

func (s *ModelApiListApiInfoEnvironments) GetAlias() *string {
	return s.Alias
}

func (s *ModelApiListApiInfoEnvironments) GetBackendScene() *string {
	return s.BackendScene
}

func (s *ModelApiListApiInfoEnvironments) GetBackendType() *string {
	return s.BackendType
}

func (s *ModelApiListApiInfoEnvironments) GetCustomDomains() []*HttpApiDomainInfo {
	return s.CustomDomains
}

func (s *ModelApiListApiInfoEnvironments) GetDeployErrorMessage() *string {
	return s.DeployErrorMessage
}

func (s *ModelApiListApiInfoEnvironments) GetDeployStatus() *string {
	return s.DeployStatus
}

func (s *ModelApiListApiInfoEnvironments) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *ModelApiListApiInfoEnvironments) GetGatewayInfo() *ModelApiListApiInfoEnvironmentsGatewayInfo {
	return s.GatewayInfo
}

func (s *ModelApiListApiInfoEnvironments) GetName() *string {
	return s.Name
}

func (s *ModelApiListApiInfoEnvironments) GetServiceConfigs() []*ModelApiListApiInfoEnvironmentsServiceConfigs {
	return s.ServiceConfigs
}

func (s *ModelApiListApiInfoEnvironments) GetSubDomains() []*ModelApiListApiInfoEnvironmentsSubDomains {
	return s.SubDomains
}

func (s *ModelApiListApiInfoEnvironments) SetAlias(v string) *ModelApiListApiInfoEnvironments {
	s.Alias = &v
	return s
}

func (s *ModelApiListApiInfoEnvironments) SetBackendScene(v string) *ModelApiListApiInfoEnvironments {
	s.BackendScene = &v
	return s
}

func (s *ModelApiListApiInfoEnvironments) SetBackendType(v string) *ModelApiListApiInfoEnvironments {
	s.BackendType = &v
	return s
}

func (s *ModelApiListApiInfoEnvironments) SetCustomDomains(v []*HttpApiDomainInfo) *ModelApiListApiInfoEnvironments {
	s.CustomDomains = v
	return s
}

func (s *ModelApiListApiInfoEnvironments) SetDeployErrorMessage(v string) *ModelApiListApiInfoEnvironments {
	s.DeployErrorMessage = &v
	return s
}

func (s *ModelApiListApiInfoEnvironments) SetDeployStatus(v string) *ModelApiListApiInfoEnvironments {
	s.DeployStatus = &v
	return s
}

func (s *ModelApiListApiInfoEnvironments) SetEnvironmentId(v string) *ModelApiListApiInfoEnvironments {
	s.EnvironmentId = &v
	return s
}

func (s *ModelApiListApiInfoEnvironments) SetGatewayInfo(v *ModelApiListApiInfoEnvironmentsGatewayInfo) *ModelApiListApiInfoEnvironments {
	s.GatewayInfo = v
	return s
}

func (s *ModelApiListApiInfoEnvironments) SetName(v string) *ModelApiListApiInfoEnvironments {
	s.Name = &v
	return s
}

func (s *ModelApiListApiInfoEnvironments) SetServiceConfigs(v []*ModelApiListApiInfoEnvironmentsServiceConfigs) *ModelApiListApiInfoEnvironments {
	s.ServiceConfigs = v
	return s
}

func (s *ModelApiListApiInfoEnvironments) SetSubDomains(v []*ModelApiListApiInfoEnvironmentsSubDomains) *ModelApiListApiInfoEnvironments {
	s.SubDomains = v
	return s
}

func (s *ModelApiListApiInfoEnvironments) Validate() error {
	if s.CustomDomains != nil {
		for _, item := range s.CustomDomains {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.GatewayInfo != nil {
		if err := s.GatewayInfo.Validate(); err != nil {
			return err
		}
	}
	if s.ServiceConfigs != nil {
		for _, item := range s.ServiceConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SubDomains != nil {
		for _, item := range s.SubDomains {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModelApiListApiInfoEnvironmentsGatewayInfo struct {
	// The edition of the gateway instance. Valid values: Professional, Serverless, MultiTenantServerless, and Unknown.
	//
	// example:
	//
	// MultiTenantServerless
	GatewayEdition *string `json:"gatewayEdition,omitempty" xml:"gatewayEdition,omitempty"`
	// The gateway instance ID.
	//
	// example:
	//
	// gw-xxx
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The name of the gateway instance.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ModelApiListApiInfoEnvironmentsGatewayInfo) String() string {
	return dara.Prettify(s)
}

func (s ModelApiListApiInfoEnvironmentsGatewayInfo) GoString() string {
	return s.String()
}

func (s *ModelApiListApiInfoEnvironmentsGatewayInfo) GetGatewayEdition() *string {
	return s.GatewayEdition
}

func (s *ModelApiListApiInfoEnvironmentsGatewayInfo) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ModelApiListApiInfoEnvironmentsGatewayInfo) GetName() *string {
	return s.Name
}

func (s *ModelApiListApiInfoEnvironmentsGatewayInfo) SetGatewayEdition(v string) *ModelApiListApiInfoEnvironmentsGatewayInfo {
	s.GatewayEdition = &v
	return s
}

func (s *ModelApiListApiInfoEnvironmentsGatewayInfo) SetGatewayId(v string) *ModelApiListApiInfoEnvironmentsGatewayInfo {
	s.GatewayId = &v
	return s
}

func (s *ModelApiListApiInfoEnvironmentsGatewayInfo) SetName(v string) *ModelApiListApiInfoEnvironmentsGatewayInfo {
	s.Name = &v
	return s
}

func (s *ModelApiListApiInfoEnvironmentsGatewayInfo) Validate() error {
	return dara.Validate(s)
}

type ModelApiListApiInfoEnvironmentsServiceConfigs struct {
	// The ID of the service.
	//
	// example:
	//
	// gs-xxx
	GatewayServiceId *string `json:"gatewayServiceId,omitempty" xml:"gatewayServiceId,omitempty"`
	// The match conditions.
	Match *HttpApiBackendMatchConditions `json:"match,omitempty" xml:"match,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// demo-service
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The service port.
	//
	// example:
	//
	// 8080
	Port *string `json:"port,omitempty" xml:"port,omitempty"`
	// The protocol.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The ID of the service.
	//
	// example:
	//
	// svc-xxx
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// The service version.
	//
	// example:
	//
	// v1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The weight.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s ModelApiListApiInfoEnvironmentsServiceConfigs) String() string {
	return dara.Prettify(s)
}

func (s ModelApiListApiInfoEnvironmentsServiceConfigs) GoString() string {
	return s.String()
}

func (s *ModelApiListApiInfoEnvironmentsServiceConfigs) GetGatewayServiceId() *string {
	return s.GatewayServiceId
}

func (s *ModelApiListApiInfoEnvironmentsServiceConfigs) GetMatch() *HttpApiBackendMatchConditions {
	return s.Match
}

func (s *ModelApiListApiInfoEnvironmentsServiceConfigs) GetName() *string {
	return s.Name
}

func (s *ModelApiListApiInfoEnvironmentsServiceConfigs) GetPort() *string {
	return s.Port
}

func (s *ModelApiListApiInfoEnvironmentsServiceConfigs) GetProtocol() *string {
	return s.Protocol
}

func (s *ModelApiListApiInfoEnvironmentsServiceConfigs) GetServiceId() *string {
	return s.ServiceId
}

func (s *ModelApiListApiInfoEnvironmentsServiceConfigs) GetVersion() *string {
	return s.Version
}

func (s *ModelApiListApiInfoEnvironmentsServiceConfigs) GetWeight() *int32 {
	return s.Weight
}

func (s *ModelApiListApiInfoEnvironmentsServiceConfigs) SetGatewayServiceId(v string) *ModelApiListApiInfoEnvironmentsServiceConfigs {
	s.GatewayServiceId = &v
	return s
}

func (s *ModelApiListApiInfoEnvironmentsServiceConfigs) SetMatch(v *HttpApiBackendMatchConditions) *ModelApiListApiInfoEnvironmentsServiceConfigs {
	s.Match = v
	return s
}

func (s *ModelApiListApiInfoEnvironmentsServiceConfigs) SetName(v string) *ModelApiListApiInfoEnvironmentsServiceConfigs {
	s.Name = &v
	return s
}

func (s *ModelApiListApiInfoEnvironmentsServiceConfigs) SetPort(v string) *ModelApiListApiInfoEnvironmentsServiceConfigs {
	s.Port = &v
	return s
}

func (s *ModelApiListApiInfoEnvironmentsServiceConfigs) SetProtocol(v string) *ModelApiListApiInfoEnvironmentsServiceConfigs {
	s.Protocol = &v
	return s
}

func (s *ModelApiListApiInfoEnvironmentsServiceConfigs) SetServiceId(v string) *ModelApiListApiInfoEnvironmentsServiceConfigs {
	s.ServiceId = &v
	return s
}

func (s *ModelApiListApiInfoEnvironmentsServiceConfigs) SetVersion(v string) *ModelApiListApiInfoEnvironmentsServiceConfigs {
	s.Version = &v
	return s
}

func (s *ModelApiListApiInfoEnvironmentsServiceConfigs) SetWeight(v int32) *ModelApiListApiInfoEnvironmentsServiceConfigs {
	s.Weight = &v
	return s
}

func (s *ModelApiListApiInfoEnvironmentsServiceConfigs) Validate() error {
	if s.Match != nil {
		if err := s.Match.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelApiListApiInfoEnvironmentsSubDomains struct {
	// The ID of the domain name.
	//
	// example:
	//
	// d-xxx
	DomainId *string `json:"domainId,omitempty" xml:"domainId,omitempty"`
	// The domain name.
	//
	// example:
	//
	// www.example.com
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The network type.
	//
	// example:
	//
	// Internet
	NetworkType *string `json:"networkType,omitempty" xml:"networkType,omitempty"`
	// The access protocol.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s ModelApiListApiInfoEnvironmentsSubDomains) String() string {
	return dara.Prettify(s)
}

func (s ModelApiListApiInfoEnvironmentsSubDomains) GoString() string {
	return s.String()
}

func (s *ModelApiListApiInfoEnvironmentsSubDomains) GetDomainId() *string {
	return s.DomainId
}

func (s *ModelApiListApiInfoEnvironmentsSubDomains) GetName() *string {
	return s.Name
}

func (s *ModelApiListApiInfoEnvironmentsSubDomains) GetNetworkType() *string {
	return s.NetworkType
}

func (s *ModelApiListApiInfoEnvironmentsSubDomains) GetProtocol() *string {
	return s.Protocol
}

func (s *ModelApiListApiInfoEnvironmentsSubDomains) SetDomainId(v string) *ModelApiListApiInfoEnvironmentsSubDomains {
	s.DomainId = &v
	return s
}

func (s *ModelApiListApiInfoEnvironmentsSubDomains) SetName(v string) *ModelApiListApiInfoEnvironmentsSubDomains {
	s.Name = &v
	return s
}

func (s *ModelApiListApiInfoEnvironmentsSubDomains) SetNetworkType(v string) *ModelApiListApiInfoEnvironmentsSubDomains {
	s.NetworkType = &v
	return s
}

func (s *ModelApiListApiInfoEnvironmentsSubDomains) SetProtocol(v string) *ModelApiListApiInfoEnvironmentsSubDomains {
	s.Protocol = &v
	return s
}

func (s *ModelApiListApiInfoEnvironmentsSubDomains) Validate() error {
	return dara.Validate(s)
}

type ModelApiListApiInfoIngressInfo struct {
	// The environment context.
	EnvironmentInfo *ModelApiListApiInfoIngressInfoEnvironmentInfo `json:"environmentInfo,omitempty" xml:"environmentInfo,omitempty" type:"Struct"`
	// The ingress class being listened to.
	//
	// example:
	//
	// apig
	IngressClass *string `json:"ingressClass,omitempty" xml:"ingressClass,omitempty"`
	// The Kubernetes cluster information.
	K8sClusterInfo *ModelApiListApiInfoIngressInfoK8sClusterInfo `json:"k8sClusterInfo,omitempty" xml:"k8sClusterInfo,omitempty" type:"Struct"`
	// Specifies whether to update the address in the ingress status.
	//
	// example:
	//
	// false
	OverrideIngressIp *bool `json:"overrideIngressIp,omitempty" xml:"overrideIngressIp,omitempty"`
	// The source ID.
	//
	// example:
	//
	// src-d00000000000000000000
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The watched namespace.
	//
	// example:
	//
	// default
	WatchNamespace *string `json:"watchNamespace,omitempty" xml:"watchNamespace,omitempty"`
}

func (s ModelApiListApiInfoIngressInfo) String() string {
	return dara.Prettify(s)
}

func (s ModelApiListApiInfoIngressInfo) GoString() string {
	return s.String()
}

func (s *ModelApiListApiInfoIngressInfo) GetEnvironmentInfo() *ModelApiListApiInfoIngressInfoEnvironmentInfo {
	return s.EnvironmentInfo
}

func (s *ModelApiListApiInfoIngressInfo) GetIngressClass() *string {
	return s.IngressClass
}

func (s *ModelApiListApiInfoIngressInfo) GetK8sClusterInfo() *ModelApiListApiInfoIngressInfoK8sClusterInfo {
	return s.K8sClusterInfo
}

func (s *ModelApiListApiInfoIngressInfo) GetOverrideIngressIp() *bool {
	return s.OverrideIngressIp
}

func (s *ModelApiListApiInfoIngressInfo) GetSourceId() *string {
	return s.SourceId
}

func (s *ModelApiListApiInfoIngressInfo) GetWatchNamespace() *string {
	return s.WatchNamespace
}

func (s *ModelApiListApiInfoIngressInfo) SetEnvironmentInfo(v *ModelApiListApiInfoIngressInfoEnvironmentInfo) *ModelApiListApiInfoIngressInfo {
	s.EnvironmentInfo = v
	return s
}

func (s *ModelApiListApiInfoIngressInfo) SetIngressClass(v string) *ModelApiListApiInfoIngressInfo {
	s.IngressClass = &v
	return s
}

func (s *ModelApiListApiInfoIngressInfo) SetK8sClusterInfo(v *ModelApiListApiInfoIngressInfoK8sClusterInfo) *ModelApiListApiInfoIngressInfo {
	s.K8sClusterInfo = v
	return s
}

func (s *ModelApiListApiInfoIngressInfo) SetOverrideIngressIp(v bool) *ModelApiListApiInfoIngressInfo {
	s.OverrideIngressIp = &v
	return s
}

func (s *ModelApiListApiInfoIngressInfo) SetSourceId(v string) *ModelApiListApiInfoIngressInfo {
	s.SourceId = &v
	return s
}

func (s *ModelApiListApiInfoIngressInfo) SetWatchNamespace(v string) *ModelApiListApiInfoIngressInfo {
	s.WatchNamespace = &v
	return s
}

func (s *ModelApiListApiInfoIngressInfo) Validate() error {
	if s.EnvironmentInfo != nil {
		if err := s.EnvironmentInfo.Validate(); err != nil {
			return err
		}
	}
	if s.K8sClusterInfo != nil {
		if err := s.K8sClusterInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelApiListApiInfoIngressInfoEnvironmentInfo struct {
	// The environment ID.
	//
	// example:
	//
	// env-d00000000000000000000
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
}

func (s ModelApiListApiInfoIngressInfoEnvironmentInfo) String() string {
	return dara.Prettify(s)
}

func (s ModelApiListApiInfoIngressInfoEnvironmentInfo) GoString() string {
	return s.String()
}

func (s *ModelApiListApiInfoIngressInfoEnvironmentInfo) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *ModelApiListApiInfoIngressInfoEnvironmentInfo) SetEnvironmentId(v string) *ModelApiListApiInfoIngressInfoEnvironmentInfo {
	s.EnvironmentId = &v
	return s
}

func (s *ModelApiListApiInfoIngressInfoEnvironmentInfo) Validate() error {
	return dara.Validate(s)
}

type ModelApiListApiInfoIngressInfoK8sClusterInfo struct {
	// The cluster ID.
	//
	// example:
	//
	// c00000000000000000000000000000001
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
}

func (s ModelApiListApiInfoIngressInfoK8sClusterInfo) String() string {
	return dara.Prettify(s)
}

func (s ModelApiListApiInfoIngressInfoK8sClusterInfo) GoString() string {
	return s.String()
}

func (s *ModelApiListApiInfoIngressInfoK8sClusterInfo) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModelApiListApiInfoIngressInfoK8sClusterInfo) SetClusterId(v string) *ModelApiListApiInfoIngressInfoK8sClusterInfo {
	s.ClusterId = &v
	return s
}

func (s *ModelApiListApiInfoIngressInfoK8sClusterInfo) Validate() error {
	return dara.Validate(s)
}

type ModelApiListApiInfoModelApiInfo struct {
	// The save state of the direct supply fallback. A value of false indicates that it is disabled. This only reflects the saved configuration and does not indicate successful release or activation.
	//
	// example:
	//
	// false
	DirectSupplyFallbackEnabled *bool `json:"directSupplyFallbackEnabled,omitempty" xml:"directSupplyFallbackEnabled,omitempty"`
	// The model exposure mode. Only SELECTED is supported in the first phase.
	//
	// example:
	//
	// SELECTED
	ExposureMode *string `json:"exposureMode,omitempty" xml:"exposureMode,omitempty"`
	// The model selection configuration.
	ModelSelection *ModelSelection `json:"modelSelection,omitempty" xml:"modelSelection,omitempty"`
	// The read-only public model collection for the current configuration. If complete is false, the collection is unknown, and partial members must not be used as the full count.
	ResolvedModelSelection *ModelApiListApiInfoModelApiInfoResolvedModelSelection `json:"resolvedModelSelection,omitempty" xml:"resolvedModelSelection,omitempty" type:"Struct"`
	// The routing mode. Only Direct is supported in the first phase.
	//
	// example:
	//
	// Direct
	RoutingMode *string `json:"routingMode,omitempty" xml:"routingMode,omitempty"`
}

func (s ModelApiListApiInfoModelApiInfo) String() string {
	return dara.Prettify(s)
}

func (s ModelApiListApiInfoModelApiInfo) GoString() string {
	return s.String()
}

func (s *ModelApiListApiInfoModelApiInfo) GetDirectSupplyFallbackEnabled() *bool {
	return s.DirectSupplyFallbackEnabled
}

func (s *ModelApiListApiInfoModelApiInfo) GetExposureMode() *string {
	return s.ExposureMode
}

func (s *ModelApiListApiInfoModelApiInfo) GetModelSelection() *ModelSelection {
	return s.ModelSelection
}

func (s *ModelApiListApiInfoModelApiInfo) GetResolvedModelSelection() *ModelApiListApiInfoModelApiInfoResolvedModelSelection {
	return s.ResolvedModelSelection
}

func (s *ModelApiListApiInfoModelApiInfo) GetRoutingMode() *string {
	return s.RoutingMode
}

func (s *ModelApiListApiInfoModelApiInfo) SetDirectSupplyFallbackEnabled(v bool) *ModelApiListApiInfoModelApiInfo {
	s.DirectSupplyFallbackEnabled = &v
	return s
}

func (s *ModelApiListApiInfoModelApiInfo) SetExposureMode(v string) *ModelApiListApiInfoModelApiInfo {
	s.ExposureMode = &v
	return s
}

func (s *ModelApiListApiInfoModelApiInfo) SetModelSelection(v *ModelSelection) *ModelApiListApiInfoModelApiInfo {
	s.ModelSelection = v
	return s
}

func (s *ModelApiListApiInfoModelApiInfo) SetResolvedModelSelection(v *ModelApiListApiInfoModelApiInfoResolvedModelSelection) *ModelApiListApiInfoModelApiInfo {
	s.ResolvedModelSelection = v
	return s
}

func (s *ModelApiListApiInfoModelApiInfo) SetRoutingMode(v string) *ModelApiListApiInfoModelApiInfo {
	s.RoutingMode = &v
	return s
}

func (s *ModelApiListApiInfoModelApiInfo) Validate() error {
	if s.ModelSelection != nil {
		if err := s.ModelSelection.Validate(); err != nil {
			return err
		}
	}
	if s.ResolvedModelSelection != nil {
		if err := s.ResolvedModelSelection.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelApiListApiInfoModelApiInfoResolvedModelSelection struct {
	// Indicates whether the group expansion collection for the current configuration is completely read. A complete empty group returns true and an empty array. A missing, invalid group, or read failure returns false. This does not represent the online release or consumer permission status.
	//
	// example:
	//
	// true
	Complete *bool `json:"complete,omitempty" xml:"complete,omitempty"`
	// The model card IDs of the publicly selected models for the current configuration, stably sorted and deduplicated. Logical models retain their own IDs, excluding fallback-only targets or underlying execution cards of logical models. If complete is false, this is an empty array and must not be interpreted as zero models.
	ModelCardIds []*string `json:"modelCardIds,omitempty" xml:"modelCardIds,omitempty" type:"Repeated"`
}

func (s ModelApiListApiInfoModelApiInfoResolvedModelSelection) String() string {
	return dara.Prettify(s)
}

func (s ModelApiListApiInfoModelApiInfoResolvedModelSelection) GoString() string {
	return s.String()
}

func (s *ModelApiListApiInfoModelApiInfoResolvedModelSelection) GetComplete() *bool {
	return s.Complete
}

func (s *ModelApiListApiInfoModelApiInfoResolvedModelSelection) GetModelCardIds() []*string {
	return s.ModelCardIds
}

func (s *ModelApiListApiInfoModelApiInfoResolvedModelSelection) SetComplete(v bool) *ModelApiListApiInfoModelApiInfoResolvedModelSelection {
	s.Complete = &v
	return s
}

func (s *ModelApiListApiInfoModelApiInfoResolvedModelSelection) SetModelCardIds(v []*string) *ModelApiListApiInfoModelApiInfoResolvedModelSelection {
	s.ModelCardIds = v
	return s
}

func (s *ModelApiListApiInfoModelApiInfoResolvedModelSelection) Validate() error {
	return dara.Validate(s)
}

type ModelApiListApiInfoModelCardRuntimeDiagnostics struct {
	// The index of the corresponding deployment configuration.
	//
	// example:
	//
	// 0
	DeployConfigIndex *int32 `json:"deployConfigIndex,omitempty" xml:"deployConfigIndex,omitempty"`
	// The gateway instance ID.
	//
	// example:
	//
	// gw-d00000000000000000000
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// Indicates whether the marker of the deployment configuration is invalid.
	//
	// example:
	//
	// true
	InvalidConfiguration *bool `json:"invalidConfiguration,omitempty" xml:"invalidConfiguration,omitempty"`
	// The list of reasons for invalidity.
	Reasons []*string `json:"reasons,omitempty" xml:"reasons,omitempty" type:"Repeated"`
}

func (s ModelApiListApiInfoModelCardRuntimeDiagnostics) String() string {
	return dara.Prettify(s)
}

func (s ModelApiListApiInfoModelCardRuntimeDiagnostics) GoString() string {
	return s.String()
}

func (s *ModelApiListApiInfoModelCardRuntimeDiagnostics) GetDeployConfigIndex() *int32 {
	return s.DeployConfigIndex
}

func (s *ModelApiListApiInfoModelCardRuntimeDiagnostics) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ModelApiListApiInfoModelCardRuntimeDiagnostics) GetInvalidConfiguration() *bool {
	return s.InvalidConfiguration
}

func (s *ModelApiListApiInfoModelCardRuntimeDiagnostics) GetReasons() []*string {
	return s.Reasons
}

func (s *ModelApiListApiInfoModelCardRuntimeDiagnostics) SetDeployConfigIndex(v int32) *ModelApiListApiInfoModelCardRuntimeDiagnostics {
	s.DeployConfigIndex = &v
	return s
}

func (s *ModelApiListApiInfoModelCardRuntimeDiagnostics) SetGatewayId(v string) *ModelApiListApiInfoModelCardRuntimeDiagnostics {
	s.GatewayId = &v
	return s
}

func (s *ModelApiListApiInfoModelCardRuntimeDiagnostics) SetInvalidConfiguration(v bool) *ModelApiListApiInfoModelCardRuntimeDiagnostics {
	s.InvalidConfiguration = &v
	return s
}

func (s *ModelApiListApiInfoModelCardRuntimeDiagnostics) SetReasons(v []*string) *ModelApiListApiInfoModelCardRuntimeDiagnostics {
	s.Reasons = v
	return s
}

func (s *ModelApiListApiInfoModelCardRuntimeDiagnostics) Validate() error {
	return dara.Validate(s)
}
