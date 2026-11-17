// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelApiApiInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAgentProtocols(v []*string) *ModelApiApiInfo
	GetAgentProtocols() []*string
	SetBasePath(v string) *ModelApiApiInfo
	GetBasePath() *string
	SetCreationClientToken(v string) *ModelApiApiInfo
	GetCreationClientToken() *string
	SetDeleteStatus(v string) *ModelApiApiInfo
	GetDeleteStatus() *string
	SetDeployCntMap(v map[string]*ModelApiApiInfoDeployCntMapValue) *ModelApiApiInfo
	GetDeployCntMap() map[string]*ModelApiApiInfoDeployCntMapValue
	SetDeployErrorMessage(v string) *ModelApiApiInfo
	GetDeployErrorMessage() *string
	SetDescription(v string) *ModelApiApiInfo
	GetDescription() *string
	SetDomainIds(v []*string) *ModelApiApiInfo
	GetDomainIds() []*string
	SetDomainInfos(v []*ModelApiDomainInfo) *ModelApiApiInfo
	GetDomainInfos() []*ModelApiDomainInfo
	SetEnableAuth(v bool) *ModelApiApiInfo
	GetEnableAuth() *bool
	SetEnvironmentId(v string) *ModelApiApiInfo
	GetEnvironmentId() *string
	SetGatewayId(v string) *ModelApiApiInfo
	GetGatewayId() *string
	SetGatewayInfo(v *GatewayInfo) *ModelApiApiInfo
	GetGatewayInfo() *GatewayInfo
	SetHttpApiId(v string) *ModelApiApiInfo
	GetHttpApiId() *string
	SetIngressInfo(v *ModelApiApiInfoIngressInfo) *ModelApiApiInfo
	GetIngressInfo() *ModelApiApiInfoIngressInfo
	SetModelApiInfo(v *ModelApiApiInfoModelApiInfo) *ModelApiApiInfo
	GetModelApiInfo() *ModelApiApiInfoModelApiInfo
	SetModelCardRuntimeDiagnostics(v []*ModelApiApiInfoModelCardRuntimeDiagnostics) *ModelApiApiInfo
	GetModelCardRuntimeDiagnostics() []*ModelApiApiInfoModelCardRuntimeDiagnostics
	SetModelCategory(v string) *ModelApiApiInfo
	GetModelCategory() *string
	SetName(v string) *ModelApiApiInfo
	GetName() *string
	SetPolicyConfigs(v []*ModelApiPolicyConfig) *ModelApiApiInfo
	GetPolicyConfigs() []*ModelApiPolicyConfig
	SetProtocols(v []*string) *ModelApiApiInfo
	GetProtocols() []*string
	SetPublishedEndpoints(v []*ModelApiApiInfoPublishedEndpoints) *ModelApiApiInfo
	GetPublishedEndpoints() []*ModelApiApiInfoPublishedEndpoints
	SetResourceGroupId(v string) *ModelApiApiInfo
	GetResourceGroupId() *string
	SetStatus(v string) *ModelApiApiInfo
	GetStatus() *string
	SetType(v string) *ModelApiApiInfo
	GetType() *string
	SetVersionInfo(v *HttpApiVersionInfo) *ModelApiApiInfo
	GetVersionInfo() *HttpApiVersionInfo
}

type ModelApiApiInfo struct {
	// The collection of agent protocols for the model API.
	AgentProtocols []*string `json:"agentProtocols,omitempty" xml:"agentProtocols,omitempty" type:"Repeated"`
	// The API base path.
	//
	// example:
	//
	// /v1
	BasePath *string `json:"basePath,omitempty" xml:"basePath,omitempty"`
	// The read-only correlation identifier for creation requests. It is returned only when the association with the same instance is unique within approximately 24 hours after the first persistence following normal authentication. It is omitted for historical, expired, or ambiguous cases. This identifier cannot be used to determine the success of publishing or authorization, nor does it replace the creation of idempotent records.
	//
	// example:
	//
	// b1907ca4-952c-4c5f-8d1f-0aa6ae339c96
	CreationClientToken *string `json:"creationClientToken,omitempty" xml:"creationClientToken,omitempty"`
	// The deletion status of the new version model API. This field is read-only and optional. Deleting indicates that the deletion request is accepted but not yet completed. DeleteFailed indicates that the background deletion failed, and you can use the original clientToken to retry the DeleteModelApi operation. This field is not returned for normal publishing and historical undeployment. The absence of this field does not indicate that the deletion is complete.
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
	DeployCntMap map[string]*ModelApiApiInfoDeployCntMapValue `json:"deployCntMap,omitempty" xml:"deployCntMap,omitempty"`
	// The reason for the deployment or deletion failure. This field is read-only.
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
	// The read-only domain name details and the default domain name flag.
	DomainInfos []*ModelApiDomainInfo `json:"domainInfos,omitempty" xml:"domainInfos,omitempty" type:"Repeated"`
	// Specifies whether to enable authentication.
	//
	// example:
	//
	// true
	EnableAuth *bool `json:"enableAuth,omitempty" xml:"enableAuth,omitempty"`
	// The default environment ID of the bound instance. This field is read-only.
	//
	// example:
	//
	// env-d00000000000000000000
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
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
	IngressInfo *ModelApiApiInfoIngressInfo `json:"ingressInfo,omitempty" xml:"ingressInfo,omitempty" type:"Struct"`
	// The model API public selection information.
	ModelApiInfo *ModelApiApiInfoModelApiInfo `json:"modelApiInfo,omitempty" xml:"modelApiInfo,omitempty" type:"Struct"`
	// The read-only runtime diagnostic information for model cards, derived from the original markers of the deployment configuration. It is used only for the readability and diagnostics of abnormal drafts.
	ModelCardRuntimeDiagnostics []*ModelApiApiInfoModelCardRuntimeDiagnostics `json:"modelCardRuntimeDiagnostics,omitempty" xml:"modelCardRuntimeDiagnostics,omitempty" type:"Repeated"`
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
	// The collection of successfully published per-path executable models. This field is read-only. It is empty when a snapshot is missing, which does not represent consumer permission.
	PublishedEndpoints []*ModelApiApiInfoPublishedEndpoints `json:"publishedEndpoints,omitempty" xml:"publishedEndpoints,omitempty" type:"Repeated"`
	// The resource group ID.
	//
	// example:
	//
	// rg-xxx
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The publishing status on the bound instance. In the new version, the status is Undeploying when a deletion is in progress, and UndeployFailed when a deletion fails. To determine whether an operation is a deletion, refer to the deleteStatus field.
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

func (s ModelApiApiInfo) String() string {
	return dara.Prettify(s)
}

func (s ModelApiApiInfo) GoString() string {
	return s.String()
}

func (s *ModelApiApiInfo) GetAgentProtocols() []*string {
	return s.AgentProtocols
}

func (s *ModelApiApiInfo) GetBasePath() *string {
	return s.BasePath
}

func (s *ModelApiApiInfo) GetCreationClientToken() *string {
	return s.CreationClientToken
}

func (s *ModelApiApiInfo) GetDeleteStatus() *string {
	return s.DeleteStatus
}

func (s *ModelApiApiInfo) GetDeployCntMap() map[string]*ModelApiApiInfoDeployCntMapValue {
	return s.DeployCntMap
}

func (s *ModelApiApiInfo) GetDeployErrorMessage() *string {
	return s.DeployErrorMessage
}

func (s *ModelApiApiInfo) GetDescription() *string {
	return s.Description
}

func (s *ModelApiApiInfo) GetDomainIds() []*string {
	return s.DomainIds
}

func (s *ModelApiApiInfo) GetDomainInfos() []*ModelApiDomainInfo {
	return s.DomainInfos
}

func (s *ModelApiApiInfo) GetEnableAuth() *bool {
	return s.EnableAuth
}

func (s *ModelApiApiInfo) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *ModelApiApiInfo) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ModelApiApiInfo) GetGatewayInfo() *GatewayInfo {
	return s.GatewayInfo
}

func (s *ModelApiApiInfo) GetHttpApiId() *string {
	return s.HttpApiId
}

func (s *ModelApiApiInfo) GetIngressInfo() *ModelApiApiInfoIngressInfo {
	return s.IngressInfo
}

func (s *ModelApiApiInfo) GetModelApiInfo() *ModelApiApiInfoModelApiInfo {
	return s.ModelApiInfo
}

func (s *ModelApiApiInfo) GetModelCardRuntimeDiagnostics() []*ModelApiApiInfoModelCardRuntimeDiagnostics {
	return s.ModelCardRuntimeDiagnostics
}

func (s *ModelApiApiInfo) GetModelCategory() *string {
	return s.ModelCategory
}

func (s *ModelApiApiInfo) GetName() *string {
	return s.Name
}

func (s *ModelApiApiInfo) GetPolicyConfigs() []*ModelApiPolicyConfig {
	return s.PolicyConfigs
}

func (s *ModelApiApiInfo) GetProtocols() []*string {
	return s.Protocols
}

func (s *ModelApiApiInfo) GetPublishedEndpoints() []*ModelApiApiInfoPublishedEndpoints {
	return s.PublishedEndpoints
}

func (s *ModelApiApiInfo) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModelApiApiInfo) GetStatus() *string {
	return s.Status
}

func (s *ModelApiApiInfo) GetType() *string {
	return s.Type
}

func (s *ModelApiApiInfo) GetVersionInfo() *HttpApiVersionInfo {
	return s.VersionInfo
}

func (s *ModelApiApiInfo) SetAgentProtocols(v []*string) *ModelApiApiInfo {
	s.AgentProtocols = v
	return s
}

func (s *ModelApiApiInfo) SetBasePath(v string) *ModelApiApiInfo {
	s.BasePath = &v
	return s
}

func (s *ModelApiApiInfo) SetCreationClientToken(v string) *ModelApiApiInfo {
	s.CreationClientToken = &v
	return s
}

func (s *ModelApiApiInfo) SetDeleteStatus(v string) *ModelApiApiInfo {
	s.DeleteStatus = &v
	return s
}

func (s *ModelApiApiInfo) SetDeployCntMap(v map[string]*ModelApiApiInfoDeployCntMapValue) *ModelApiApiInfo {
	s.DeployCntMap = v
	return s
}

func (s *ModelApiApiInfo) SetDeployErrorMessage(v string) *ModelApiApiInfo {
	s.DeployErrorMessage = &v
	return s
}

func (s *ModelApiApiInfo) SetDescription(v string) *ModelApiApiInfo {
	s.Description = &v
	return s
}

func (s *ModelApiApiInfo) SetDomainIds(v []*string) *ModelApiApiInfo {
	s.DomainIds = v
	return s
}

func (s *ModelApiApiInfo) SetDomainInfos(v []*ModelApiDomainInfo) *ModelApiApiInfo {
	s.DomainInfos = v
	return s
}

func (s *ModelApiApiInfo) SetEnableAuth(v bool) *ModelApiApiInfo {
	s.EnableAuth = &v
	return s
}

func (s *ModelApiApiInfo) SetEnvironmentId(v string) *ModelApiApiInfo {
	s.EnvironmentId = &v
	return s
}

func (s *ModelApiApiInfo) SetGatewayId(v string) *ModelApiApiInfo {
	s.GatewayId = &v
	return s
}

func (s *ModelApiApiInfo) SetGatewayInfo(v *GatewayInfo) *ModelApiApiInfo {
	s.GatewayInfo = v
	return s
}

func (s *ModelApiApiInfo) SetHttpApiId(v string) *ModelApiApiInfo {
	s.HttpApiId = &v
	return s
}

func (s *ModelApiApiInfo) SetIngressInfo(v *ModelApiApiInfoIngressInfo) *ModelApiApiInfo {
	s.IngressInfo = v
	return s
}

func (s *ModelApiApiInfo) SetModelApiInfo(v *ModelApiApiInfoModelApiInfo) *ModelApiApiInfo {
	s.ModelApiInfo = v
	return s
}

func (s *ModelApiApiInfo) SetModelCardRuntimeDiagnostics(v []*ModelApiApiInfoModelCardRuntimeDiagnostics) *ModelApiApiInfo {
	s.ModelCardRuntimeDiagnostics = v
	return s
}

func (s *ModelApiApiInfo) SetModelCategory(v string) *ModelApiApiInfo {
	s.ModelCategory = &v
	return s
}

func (s *ModelApiApiInfo) SetName(v string) *ModelApiApiInfo {
	s.Name = &v
	return s
}

func (s *ModelApiApiInfo) SetPolicyConfigs(v []*ModelApiPolicyConfig) *ModelApiApiInfo {
	s.PolicyConfigs = v
	return s
}

func (s *ModelApiApiInfo) SetProtocols(v []*string) *ModelApiApiInfo {
	s.Protocols = v
	return s
}

func (s *ModelApiApiInfo) SetPublishedEndpoints(v []*ModelApiApiInfoPublishedEndpoints) *ModelApiApiInfo {
	s.PublishedEndpoints = v
	return s
}

func (s *ModelApiApiInfo) SetResourceGroupId(v string) *ModelApiApiInfo {
	s.ResourceGroupId = &v
	return s
}

func (s *ModelApiApiInfo) SetStatus(v string) *ModelApiApiInfo {
	s.Status = &v
	return s
}

func (s *ModelApiApiInfo) SetType(v string) *ModelApiApiInfo {
	s.Type = &v
	return s
}

func (s *ModelApiApiInfo) SetVersionInfo(v *HttpApiVersionInfo) *ModelApiApiInfo {
	s.VersionInfo = v
	return s
}

func (s *ModelApiApiInfo) Validate() error {
	if s.DomainInfos != nil {
		for _, item := range s.DomainInfos {
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
	if s.PublishedEndpoints != nil {
		for _, item := range s.PublishedEndpoints {
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

type ModelApiApiInfoIngressInfo struct {
	// The environment context.
	EnvironmentInfo *ModelApiApiInfoIngressInfoEnvironmentInfo `json:"environmentInfo,omitempty" xml:"environmentInfo,omitempty" type:"Struct"`
	// The ingress class.
	//
	// example:
	//
	// apig
	IngressClass *string `json:"ingressClass,omitempty" xml:"ingressClass,omitempty"`
	// The Kubernetes cluster information.
	K8sClusterInfo *ModelApiApiInfoIngressInfoK8sClusterInfo `json:"k8sClusterInfo,omitempty" xml:"k8sClusterInfo,omitempty" type:"Struct"`
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

func (s ModelApiApiInfoIngressInfo) String() string {
	return dara.Prettify(s)
}

func (s ModelApiApiInfoIngressInfo) GoString() string {
	return s.String()
}

func (s *ModelApiApiInfoIngressInfo) GetEnvironmentInfo() *ModelApiApiInfoIngressInfoEnvironmentInfo {
	return s.EnvironmentInfo
}

func (s *ModelApiApiInfoIngressInfo) GetIngressClass() *string {
	return s.IngressClass
}

func (s *ModelApiApiInfoIngressInfo) GetK8sClusterInfo() *ModelApiApiInfoIngressInfoK8sClusterInfo {
	return s.K8sClusterInfo
}

func (s *ModelApiApiInfoIngressInfo) GetOverrideIngressIp() *bool {
	return s.OverrideIngressIp
}

func (s *ModelApiApiInfoIngressInfo) GetSourceId() *string {
	return s.SourceId
}

func (s *ModelApiApiInfoIngressInfo) GetWatchNamespace() *string {
	return s.WatchNamespace
}

func (s *ModelApiApiInfoIngressInfo) SetEnvironmentInfo(v *ModelApiApiInfoIngressInfoEnvironmentInfo) *ModelApiApiInfoIngressInfo {
	s.EnvironmentInfo = v
	return s
}

func (s *ModelApiApiInfoIngressInfo) SetIngressClass(v string) *ModelApiApiInfoIngressInfo {
	s.IngressClass = &v
	return s
}

func (s *ModelApiApiInfoIngressInfo) SetK8sClusterInfo(v *ModelApiApiInfoIngressInfoK8sClusterInfo) *ModelApiApiInfoIngressInfo {
	s.K8sClusterInfo = v
	return s
}

func (s *ModelApiApiInfoIngressInfo) SetOverrideIngressIp(v bool) *ModelApiApiInfoIngressInfo {
	s.OverrideIngressIp = &v
	return s
}

func (s *ModelApiApiInfoIngressInfo) SetSourceId(v string) *ModelApiApiInfoIngressInfo {
	s.SourceId = &v
	return s
}

func (s *ModelApiApiInfoIngressInfo) SetWatchNamespace(v string) *ModelApiApiInfoIngressInfo {
	s.WatchNamespace = &v
	return s
}

func (s *ModelApiApiInfoIngressInfo) Validate() error {
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

type ModelApiApiInfoIngressInfoEnvironmentInfo struct {
	// The environment ID.
	//
	// example:
	//
	// env-d00000000000000000000
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
}

func (s ModelApiApiInfoIngressInfoEnvironmentInfo) String() string {
	return dara.Prettify(s)
}

func (s ModelApiApiInfoIngressInfoEnvironmentInfo) GoString() string {
	return s.String()
}

func (s *ModelApiApiInfoIngressInfoEnvironmentInfo) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *ModelApiApiInfoIngressInfoEnvironmentInfo) SetEnvironmentId(v string) *ModelApiApiInfoIngressInfoEnvironmentInfo {
	s.EnvironmentId = &v
	return s
}

func (s *ModelApiApiInfoIngressInfoEnvironmentInfo) Validate() error {
	return dara.Validate(s)
}

type ModelApiApiInfoIngressInfoK8sClusterInfo struct {
	// The cluster ID.
	//
	// example:
	//
	// c00000000000000000000000000000001
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
}

func (s ModelApiApiInfoIngressInfoK8sClusterInfo) String() string {
	return dara.Prettify(s)
}

func (s ModelApiApiInfoIngressInfoK8sClusterInfo) GoString() string {
	return s.String()
}

func (s *ModelApiApiInfoIngressInfoK8sClusterInfo) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModelApiApiInfoIngressInfoK8sClusterInfo) SetClusterId(v string) *ModelApiApiInfoIngressInfoK8sClusterInfo {
	s.ClusterId = &v
	return s
}

func (s *ModelApiApiInfoIngressInfoK8sClusterInfo) Validate() error {
	return dara.Validate(s)
}

type ModelApiApiInfoModelApiInfo struct {
	// The saved state of the direct supply fallback. A value of false indicates that it is disabled. This only reflects the saved configuration and does not indicate successful publishing or activation.
	//
	// example:
	//
	// false
	DirectSupplyFallbackEnabled *bool `json:"directSupplyFallbackEnabled,omitempty" xml:"directSupplyFallbackEnabled,omitempty"`
	// The model exposure mode. Only SELECTED is supported in Phase 1.
	//
	// example:
	//
	// SELECTED
	ExposureMode *string `json:"exposureMode,omitempty" xml:"exposureMode,omitempty"`
	// The model selection configuration.
	ModelSelection *ModelSelection `json:"modelSelection,omitempty" xml:"modelSelection,omitempty"`
	// The read-only collection of publicly exposed models for the current configuration. It retains the modelSelection edit source, but the successful online collection is still subject to publishedEndpoints.
	ResolvedModelSelection *ModelApiApiInfoModelApiInfoResolvedModelSelection `json:"resolvedModelSelection,omitempty" xml:"resolvedModelSelection,omitempty" type:"Struct"`
	// The routing mode. Only Direct is supported in Phase 1.
	//
	// example:
	//
	// Direct
	RoutingMode *string `json:"routingMode,omitempty" xml:"routingMode,omitempty"`
}

func (s ModelApiApiInfoModelApiInfo) String() string {
	return dara.Prettify(s)
}

func (s ModelApiApiInfoModelApiInfo) GoString() string {
	return s.String()
}

func (s *ModelApiApiInfoModelApiInfo) GetDirectSupplyFallbackEnabled() *bool {
	return s.DirectSupplyFallbackEnabled
}

func (s *ModelApiApiInfoModelApiInfo) GetExposureMode() *string {
	return s.ExposureMode
}

func (s *ModelApiApiInfoModelApiInfo) GetModelSelection() *ModelSelection {
	return s.ModelSelection
}

func (s *ModelApiApiInfoModelApiInfo) GetResolvedModelSelection() *ModelApiApiInfoModelApiInfoResolvedModelSelection {
	return s.ResolvedModelSelection
}

func (s *ModelApiApiInfoModelApiInfo) GetRoutingMode() *string {
	return s.RoutingMode
}

func (s *ModelApiApiInfoModelApiInfo) SetDirectSupplyFallbackEnabled(v bool) *ModelApiApiInfoModelApiInfo {
	s.DirectSupplyFallbackEnabled = &v
	return s
}

func (s *ModelApiApiInfoModelApiInfo) SetExposureMode(v string) *ModelApiApiInfoModelApiInfo {
	s.ExposureMode = &v
	return s
}

func (s *ModelApiApiInfoModelApiInfo) SetModelSelection(v *ModelSelection) *ModelApiApiInfoModelApiInfo {
	s.ModelSelection = v
	return s
}

func (s *ModelApiApiInfoModelApiInfo) SetResolvedModelSelection(v *ModelApiApiInfoModelApiInfoResolvedModelSelection) *ModelApiApiInfoModelApiInfo {
	s.ResolvedModelSelection = v
	return s
}

func (s *ModelApiApiInfoModelApiInfo) SetRoutingMode(v string) *ModelApiApiInfoModelApiInfo {
	s.RoutingMode = &v
	return s
}

func (s *ModelApiApiInfoModelApiInfo) Validate() error {
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

type ModelApiApiInfoModelApiInfoResolvedModelSelection struct {
	// Indicates whether the expanded collection of the current configuration is fully read. A fully empty group returns true and an empty array. Missing, invalid groups, or read failures return false. This does not reflect the online publishing or consumer permission status.
	//
	// example:
	//
	// true
	Complete *bool `json:"complete,omitempty" xml:"complete,omitempty"`
	// The identities of the models selected for public exposure in the current configuration, stably sorted and deduplicated. Logical models retain their own IDs and do not include fallback-only targets or underlying execution cards of logical models. If complete is false, this is an empty array and must not be interpreted as zero models.
	ModelCardIds []*string `json:"modelCardIds,omitempty" xml:"modelCardIds,omitempty" type:"Repeated"`
}

func (s ModelApiApiInfoModelApiInfoResolvedModelSelection) String() string {
	return dara.Prettify(s)
}

func (s ModelApiApiInfoModelApiInfoResolvedModelSelection) GoString() string {
	return s.String()
}

func (s *ModelApiApiInfoModelApiInfoResolvedModelSelection) GetComplete() *bool {
	return s.Complete
}

func (s *ModelApiApiInfoModelApiInfoResolvedModelSelection) GetModelCardIds() []*string {
	return s.ModelCardIds
}

func (s *ModelApiApiInfoModelApiInfoResolvedModelSelection) SetComplete(v bool) *ModelApiApiInfoModelApiInfoResolvedModelSelection {
	s.Complete = &v
	return s
}

func (s *ModelApiApiInfoModelApiInfoResolvedModelSelection) SetModelCardIds(v []*string) *ModelApiApiInfoModelApiInfoResolvedModelSelection {
	s.ModelCardIds = v
	return s
}

func (s *ModelApiApiInfoModelApiInfoResolvedModelSelection) Validate() error {
	return dara.Validate(s)
}

type ModelApiApiInfoModelCardRuntimeDiagnostics struct {
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

func (s ModelApiApiInfoModelCardRuntimeDiagnostics) String() string {
	return dara.Prettify(s)
}

func (s ModelApiApiInfoModelCardRuntimeDiagnostics) GoString() string {
	return s.String()
}

func (s *ModelApiApiInfoModelCardRuntimeDiagnostics) GetDeployConfigIndex() *int32 {
	return s.DeployConfigIndex
}

func (s *ModelApiApiInfoModelCardRuntimeDiagnostics) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ModelApiApiInfoModelCardRuntimeDiagnostics) GetInvalidConfiguration() *bool {
	return s.InvalidConfiguration
}

func (s *ModelApiApiInfoModelCardRuntimeDiagnostics) GetReasons() []*string {
	return s.Reasons
}

func (s *ModelApiApiInfoModelCardRuntimeDiagnostics) SetDeployConfigIndex(v int32) *ModelApiApiInfoModelCardRuntimeDiagnostics {
	s.DeployConfigIndex = &v
	return s
}

func (s *ModelApiApiInfoModelCardRuntimeDiagnostics) SetGatewayId(v string) *ModelApiApiInfoModelCardRuntimeDiagnostics {
	s.GatewayId = &v
	return s
}

func (s *ModelApiApiInfoModelCardRuntimeDiagnostics) SetInvalidConfiguration(v bool) *ModelApiApiInfoModelCardRuntimeDiagnostics {
	s.InvalidConfiguration = &v
	return s
}

func (s *ModelApiApiInfoModelCardRuntimeDiagnostics) SetReasons(v []*string) *ModelApiApiInfoModelCardRuntimeDiagnostics {
	s.Reasons = v
	return s
}

func (s *ModelApiApiInfoModelCardRuntimeDiagnostics) Validate() error {
	return dara.Validate(s)
}

type ModelApiApiInfoPublishedEndpoints struct {
	// The name of the successfully published built-in endpoint.
	//
	// example:
	//
	// model-responses
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// The collection of executable ModelCard IDs successfully published for this path.
	ModelCardIds []*string `json:"modelCardIds,omitempty" xml:"modelCardIds,omitempty" type:"Repeated"`
}

func (s ModelApiApiInfoPublishedEndpoints) String() string {
	return dara.Prettify(s)
}

func (s ModelApiApiInfoPublishedEndpoints) GoString() string {
	return s.String()
}

func (s *ModelApiApiInfoPublishedEndpoints) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ModelApiApiInfoPublishedEndpoints) GetModelCardIds() []*string {
	return s.ModelCardIds
}

func (s *ModelApiApiInfoPublishedEndpoints) SetEndpoint(v string) *ModelApiApiInfoPublishedEndpoints {
	s.Endpoint = &v
	return s
}

func (s *ModelApiApiInfoPublishedEndpoints) SetModelCardIds(v []*string) *ModelApiApiInfoPublishedEndpoints {
	s.ModelCardIds = v
	return s
}

func (s *ModelApiApiInfoPublishedEndpoints) Validate() error {
	return dara.Validate(s)
}
