// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTypedAIAsset interface {
	dara.Model
	String() string
	GoString() string
	SetAccessType(v string) *TypedAIAsset
	GetAccessType() *string
	SetAgentId(v string) *TypedAIAsset
	GetAgentId() *string
	SetConfigPath(v string) *TypedAIAsset
	GetConfigPath() *string
	SetDescription(v string) *TypedAIAsset
	GetDescription() *string
	SetEndpoint(v string) *TypedAIAsset
	GetEndpoint() *string
	SetEndpointType(v string) *TypedAIAsset
	GetEndpointType() *string
	SetFirstSeen(v string) *TypedAIAsset
	GetFirstSeen() *string
	SetForm(v string) *TypedAIAsset
	GetForm() *string
	SetId(v string) *TypedAIAsset
	GetId() *string
	SetIdentityIdentifiers(v *TypedAIAssetIdentityIdentifiers) *TypedAIAsset
	GetIdentityIdentifiers() *TypedAIAssetIdentityIdentifiers
	SetInstallPath(v string) *TypedAIAsset
	GetInstallPath() *string
	SetInstanceLocalId(v string) *TypedAIAsset
	GetInstanceLocalId() *string
	SetLastCheckTime(v string) *TypedAIAsset
	GetLastCheckTime() *string
	SetLastDetectTaskId(v string) *TypedAIAsset
	GetLastDetectTaskId() *string
	SetLastSeen(v string) *TypedAIAsset
	GetLastSeen() *string
	SetMatchedRules(v []*string) *TypedAIAsset
	GetMatchedRules() []*string
	SetMcpId(v string) *TypedAIAsset
	GetMcpId() *string
	SetProductCode(v string) *TypedAIAsset
	GetProductCode() *string
	SetProviderId(v string) *TypedAIAsset
	GetProviderId() *string
	SetProviderName(v string) *TypedAIAsset
	GetProviderName() *string
	SetProviderNameCn(v string) *TypedAIAsset
	GetProviderNameCn() *string
	SetProviderNameEn(v string) *TypedAIAsset
	GetProviderNameEn() *string
	SetProviderType(v string) *TypedAIAsset
	GetProviderType() *string
	SetRecognitionEvidence(v *TypedAIAssetRecognitionEvidence) *TypedAIAsset
	GetRecognitionEvidence() *TypedAIAssetRecognitionEvidence
	SetRecognitionStatus(v string) *TypedAIAsset
	GetRecognitionStatus() *string
	SetRiskInfo(v *TypedAIAssetRiskInfo) *TypedAIAsset
	GetRiskInfo() *TypedAIAssetRiskInfo
	SetRiskLevel(v string) *TypedAIAsset
	GetRiskLevel() *string
	SetScanId(v string) *TypedAIAsset
	GetScanId() *string
	SetSourceType(v string) *TypedAIAsset
	GetSourceType() *string
	SetStandardSha256(v string) *TypedAIAsset
	GetStandardSha256() *string
	SetTransport(v string) *TypedAIAsset
	GetTransport() *string
	SetVendor(v int32) *TypedAIAsset
	GetVendor() *int32
	SetVersion(v string) *TypedAIAsset
	GetVersion() *string
}

type TypedAIAsset struct {
	// The access method for the MCP. Valid values: local_process for local processes, remote_service for remote services, and unknown for undetermined methods.
	//
	// example:
	//
	// local_process
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// The database primary key of the agent instance associated with the leaf asset, returned as a decimal numeric string. The value can be 0 when no agent is associated, and it can also be 0 in the details of the agent itself. This field is not returned in aggregated lists.
	//
	// example:
	//
	// 1234567890123456789
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The configuration path where the LLM, MCP, or tool declaration is located. This is not a skill package identifier. It is returned in instance details but not in aggregated lists.
	//
	// example:
	//
	// /home/example/.config/example-agent/config.json
	ConfigPath *string `json:"ConfigPath,omitempty" xml:"ConfigPath,omitempty"`
	// The description text collected for an MCP, tool, or skill. This is not returned if missing. Aggregated lists only return descriptions that are consistent within the group.
	//
	// example:
	//
	// 读取并整理指定目录中的文档。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The configured service endpoint for an LLM model service or remote MCP. Local stdio MCPs may not have this field. Its presence does not indicate a successful connection or actual invocation.
	//
	// example:
	//
	// https://api.example.com/v1
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The classification of the LLM service endpoint. Valid values: official for official services, third_party for third-party services, and unknown for undetermined services. This classification does not represent a security detection conclusion.
	//
	// example:
	//
	// official
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The time when the asset instance was first observed. The value is a database datetime string that may include six decimal places for seconds, without a unified time zone indicator. This field is not returned in aggregated lists.
	//
	// example:
	//
	// 2026-10-16 10:30:00.000000
	FirstSeen *string `json:"FirstSeen,omitempty" xml:"FirstSeen,omitempty"`
	// The installation form of the agent, such as desktop_app, ide_extension, cli, or sdk. This is not the endpoint operating system such as Windows or macOS. The specific value is subject to the instance recognition result.
	//
	// example:
	//
	// cli
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The database primary key of the table where the asset instance is located, returned as a decimal numeric string to prevent precision loss of large integers. This is not the asset aggregation key AssetKey, and this instance field is not returned in aggregated lists.
	//
	// example:
	//
	// 1234567890123456789
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The collected product identity clues and their categorization by identifier type. These are used for auxiliary identification and do not indicate that signatures, publishers, or file contents have been authenticated through security authentication.
	//
	// example:
	//
	// {"macos_bundle_id":["com.example.agent"],"npm_package":["@example/agent-cli"]}
	IdentityIdentifiers *TypedAIAssetIdentityIdentifiers `json:"IdentityIdentifiers,omitempty" xml:"IdentityIdentifiers,omitempty" type:"Struct"`
	// The installation path of the agent instance. This field is used only for instance details and is not returned in aggregated lists. It is not applicable to other asset types.
	//
	// example:
	//
	// /opt/example-agent
	InstallPath *string `json:"InstallPath,omitempty" xml:"InstallPath,omitempty"`
	// The instance identity installed or declared within the endpoint, calculated by the script based on the asset type, ownership, and location. The identity scope must still include the tenant and endpoint, and it cannot be used as a globally unique ID across endpoints. This field is not returned in aggregation lists.
	//
	// example:
	//
	// sha256:70e28ce51d55863f2ac5bb81efcee54026ff9fdde9b52346c659f4e24017e7f6
	InstanceLocalId *string `json:"InstanceLocalId,omitempty" xml:"InstanceLocalId,omitempty"`
	// The time of the last successful detection for the current skill content. This field is empty or not returned if no saved successful detection record exists. The value is a database datetime string without a unified time zone indicator. This field is not returned in aggregated lists.
	//
	// example:
	//
	// 2026-10-16 10:30:00.000000
	LastCheckTime *string `json:"LastCheckTime,omitempty" xml:"LastCheckTime,omitempty"`
	// The database ID of the latest detection task associated with the current skill content, returned as a decimal numeric string. The value can be 0 if no task is associated. The presence of a task ID does not indicate that the detection is complete. This field is not returned in aggregated lists.
	//
	// example:
	//
	// 1234567890123456791
	LastDetectTaskId *string `json:"LastDetectTaskId,omitempty" xml:"LastDetectTaskId,omitempty"`
	// The time when the asset instance was last actually observed. The value is a database datetime string that may include six decimal places for seconds. This is not the page query time. This field is not returned in aggregated lists.
	//
	// example:
	//
	// 2026-10-16 11:30:00.000000
	LastSeen *string `json:"LastSeen,omitempty" xml:"LastSeen,omitempty"`
	// The array of static recognition basis tags matched by the agent, which can include model_call, tool_call, autonomous_execution, state_management, and agent_declaration. This does not indicate that the agent was actually executed or passed security certification.
	//
	// example:
	//
	// ["model_call","tool_call"]
	MatchedRules []*string `json:"MatchedRules,omitempty" xml:"MatchedRules,omitempty" type:"Repeated"`
	// The database primary key of the MCP instance associated with the tool, returned as a decimal numeric string. The value can be 0 when no MCP is associated, and this field is not applicable to other asset types. This field is not returned in aggregated lists.
	//
	// example:
	//
	// 1234567890123456790
	McpId *string `json:"McpId,omitempty" xml:"McpId,omitempty"`
	// The product code of the agent. The specific value is subject to the product catalog in the recognition rules. This field is not returned when it is not applicable, unidentified, or inconsistent across aggregated attributes.
	//
	// example:
	//
	// example_agent
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The provider identifier matched based on the service endpoint and provider directory. This may not be returned if no match is found and is not the product code of the associated agent.
	//
	// example:
	//
	// example_provider
	ProviderId *string `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
	// The provider display name of the tool. This is generated from the associated MCP, agent, or declaration source and is not a proof of security or trust.
	//
	// example:
	//
	// Example MCP
	ProviderName *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
	// The Chinese display name of the service provider. The directory name is prioritized. The name parsed from the service endpoint may be used if no match is found. The presence of this field does not indicate that the service provider is identified.
	//
	// example:
	//
	// 示例模型服务商
	ProviderNameCn *string `json:"ProviderNameCn,omitempty" xml:"ProviderNameCn,omitempty"`
	// The English display name of the service provider. The directory name is prioritized. The name parsed from the service endpoint may be used if no match is found. The presence of this field does not indicate that the service provider is identified.
	//
	// example:
	//
	// Example Model Provider
	ProviderNameEn *string `json:"ProviderNameEn,omitempty" xml:"ProviderNameEn,omitempty"`
	// The declared provider type of the tool. Valid values: mcp, agent, software, file, or unknown. This describes the declaration source and does not represent the actual caller, authorization status, or vendor identity.
	//
	// example:
	//
	// mcp
	ProviderType *string `json:"ProviderType,omitempty" xml:"ProviderType,omitempty"`
	// The static recognition basis of the agent, which lists the source paths and evidence text grouped by recognition tags. This is not a runtime audit record, and instance-level evidence is not returned in aggregated lists.
	//
	// example:
	//
	// {"model_call":[{"source_path":"/opt/example-agent/model.py"}]}
	RecognitionEvidence *TypedAIAssetRecognitionEvidence `json:"RecognitionEvidence,omitempty" xml:"RecognitionEvidence,omitempty" type:"Struct"`
	// The product recognition status of the agent. identified indicates recognized, unidentified indicates pending recognition, and conflict indicates a recognition conflict. This status does not indicate whether the asset is secure.
	//
	// example:
	//
	// identified
	RecognitionStatus *string `json:"RecognitionStatus,omitempty" xml:"RecognitionStatus,omitempty"`
	// The counts of high, medium, and low-risk items saved for the current skill content. This may be empty or not returned when no detection results are available. A missing value cannot be interpreted as a zero risk count or an indication of security.
	//
	// example:
	//
	// {"highRiskCnt":1,"mediumRiskCnt":0,"lowRiskCnt":0}
	RiskInfo *TypedAIAssetRiskInfo `json:"RiskInfo,omitempty" xml:"RiskInfo,omitempty" type:"Struct"`
	// The final risk level saved for the current skill content. This is not returned when no detection conclusion is available. The value high is provided only as a format example and does not serve as a new enumeration constraint.
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The scan batch identifier of the most recent scan where the instance was actually observed. This is not a database primary key or asset content hash. This field is not returned in aggregated lists.
	//
	// example:
	//
	// scan-0123456789abcdef0123456789abcdef
	ScanId *string `json:"ScanId,omitempty" xml:"ScanId,omitempty"`
	// The source channel classification of the skill. The default value for current collection and storage is unknown. The unknown value cannot be inferred as an official or community source.
	//
	// example:
	//
	// unknown
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The SHA-256 hash of the complete skill package bytes. The format is a 64-character lowercase hexadecimal string without the sha256: prefix. The current script uses the skill-canonical-zip-v1 specification for compression. This is not equivalent to the installation instance ID and does not indicate that the package has been inspected or is secure.
	//
	// example:
	//
	// 60d3292dd23fe7d2760d6a228123d1e4eb3d91bdb511089015e854d68ea9e460
	StandardSha256 *string `json:"StandardSha256,omitempty" xml:"StandardSha256,omitempty"`
	// The transport protocol declared in the MCP configuration. The collected value is retained as-is, such as stdio, sse, or streamable_http. This cannot be used to determine whether the service is running or has been invoked.
	//
	// example:
	//
	// stdio
	Transport *string `json:"Transport,omitempty" xml:"Transport,omitempty"`
	// The vendor code for the collected and stored skill. The current default value is 1000. This cannot be used to infer a specific vendor or security status.
	//
	// example:
	//
	// 1000
	Vendor *int32 `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// The collected installation version of the agent. This field is not returned when the version is not collected. In aggregated lists, this field is returned only when the versions of all instances in the group are identical and not empty.
	//
	// example:
	//
	// 1.2.3
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s TypedAIAsset) String() string {
	return dara.Prettify(s)
}

func (s TypedAIAsset) GoString() string {
	return s.String()
}

func (s *TypedAIAsset) GetAccessType() *string {
	return s.AccessType
}

func (s *TypedAIAsset) GetAgentId() *string {
	return s.AgentId
}

func (s *TypedAIAsset) GetConfigPath() *string {
	return s.ConfigPath
}

func (s *TypedAIAsset) GetDescription() *string {
	return s.Description
}

func (s *TypedAIAsset) GetEndpoint() *string {
	return s.Endpoint
}

func (s *TypedAIAsset) GetEndpointType() *string {
	return s.EndpointType
}

func (s *TypedAIAsset) GetFirstSeen() *string {
	return s.FirstSeen
}

func (s *TypedAIAsset) GetForm() *string {
	return s.Form
}

func (s *TypedAIAsset) GetId() *string {
	return s.Id
}

func (s *TypedAIAsset) GetIdentityIdentifiers() *TypedAIAssetIdentityIdentifiers {
	return s.IdentityIdentifiers
}

func (s *TypedAIAsset) GetInstallPath() *string {
	return s.InstallPath
}

func (s *TypedAIAsset) GetInstanceLocalId() *string {
	return s.InstanceLocalId
}

func (s *TypedAIAsset) GetLastCheckTime() *string {
	return s.LastCheckTime
}

func (s *TypedAIAsset) GetLastDetectTaskId() *string {
	return s.LastDetectTaskId
}

func (s *TypedAIAsset) GetLastSeen() *string {
	return s.LastSeen
}

func (s *TypedAIAsset) GetMatchedRules() []*string {
	return s.MatchedRules
}

func (s *TypedAIAsset) GetMcpId() *string {
	return s.McpId
}

func (s *TypedAIAsset) GetProductCode() *string {
	return s.ProductCode
}

func (s *TypedAIAsset) GetProviderId() *string {
	return s.ProviderId
}

func (s *TypedAIAsset) GetProviderName() *string {
	return s.ProviderName
}

func (s *TypedAIAsset) GetProviderNameCn() *string {
	return s.ProviderNameCn
}

func (s *TypedAIAsset) GetProviderNameEn() *string {
	return s.ProviderNameEn
}

func (s *TypedAIAsset) GetProviderType() *string {
	return s.ProviderType
}

func (s *TypedAIAsset) GetRecognitionEvidence() *TypedAIAssetRecognitionEvidence {
	return s.RecognitionEvidence
}

func (s *TypedAIAsset) GetRecognitionStatus() *string {
	return s.RecognitionStatus
}

func (s *TypedAIAsset) GetRiskInfo() *TypedAIAssetRiskInfo {
	return s.RiskInfo
}

func (s *TypedAIAsset) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *TypedAIAsset) GetScanId() *string {
	return s.ScanId
}

func (s *TypedAIAsset) GetSourceType() *string {
	return s.SourceType
}

func (s *TypedAIAsset) GetStandardSha256() *string {
	return s.StandardSha256
}

func (s *TypedAIAsset) GetTransport() *string {
	return s.Transport
}

func (s *TypedAIAsset) GetVendor() *int32 {
	return s.Vendor
}

func (s *TypedAIAsset) GetVersion() *string {
	return s.Version
}

func (s *TypedAIAsset) SetAccessType(v string) *TypedAIAsset {
	s.AccessType = &v
	return s
}

func (s *TypedAIAsset) SetAgentId(v string) *TypedAIAsset {
	s.AgentId = &v
	return s
}

func (s *TypedAIAsset) SetConfigPath(v string) *TypedAIAsset {
	s.ConfigPath = &v
	return s
}

func (s *TypedAIAsset) SetDescription(v string) *TypedAIAsset {
	s.Description = &v
	return s
}

func (s *TypedAIAsset) SetEndpoint(v string) *TypedAIAsset {
	s.Endpoint = &v
	return s
}

func (s *TypedAIAsset) SetEndpointType(v string) *TypedAIAsset {
	s.EndpointType = &v
	return s
}

func (s *TypedAIAsset) SetFirstSeen(v string) *TypedAIAsset {
	s.FirstSeen = &v
	return s
}

func (s *TypedAIAsset) SetForm(v string) *TypedAIAsset {
	s.Form = &v
	return s
}

func (s *TypedAIAsset) SetId(v string) *TypedAIAsset {
	s.Id = &v
	return s
}

func (s *TypedAIAsset) SetIdentityIdentifiers(v *TypedAIAssetIdentityIdentifiers) *TypedAIAsset {
	s.IdentityIdentifiers = v
	return s
}

func (s *TypedAIAsset) SetInstallPath(v string) *TypedAIAsset {
	s.InstallPath = &v
	return s
}

func (s *TypedAIAsset) SetInstanceLocalId(v string) *TypedAIAsset {
	s.InstanceLocalId = &v
	return s
}

func (s *TypedAIAsset) SetLastCheckTime(v string) *TypedAIAsset {
	s.LastCheckTime = &v
	return s
}

func (s *TypedAIAsset) SetLastDetectTaskId(v string) *TypedAIAsset {
	s.LastDetectTaskId = &v
	return s
}

func (s *TypedAIAsset) SetLastSeen(v string) *TypedAIAsset {
	s.LastSeen = &v
	return s
}

func (s *TypedAIAsset) SetMatchedRules(v []*string) *TypedAIAsset {
	s.MatchedRules = v
	return s
}

func (s *TypedAIAsset) SetMcpId(v string) *TypedAIAsset {
	s.McpId = &v
	return s
}

func (s *TypedAIAsset) SetProductCode(v string) *TypedAIAsset {
	s.ProductCode = &v
	return s
}

func (s *TypedAIAsset) SetProviderId(v string) *TypedAIAsset {
	s.ProviderId = &v
	return s
}

func (s *TypedAIAsset) SetProviderName(v string) *TypedAIAsset {
	s.ProviderName = &v
	return s
}

func (s *TypedAIAsset) SetProviderNameCn(v string) *TypedAIAsset {
	s.ProviderNameCn = &v
	return s
}

func (s *TypedAIAsset) SetProviderNameEn(v string) *TypedAIAsset {
	s.ProviderNameEn = &v
	return s
}

func (s *TypedAIAsset) SetProviderType(v string) *TypedAIAsset {
	s.ProviderType = &v
	return s
}

func (s *TypedAIAsset) SetRecognitionEvidence(v *TypedAIAssetRecognitionEvidence) *TypedAIAsset {
	s.RecognitionEvidence = v
	return s
}

func (s *TypedAIAsset) SetRecognitionStatus(v string) *TypedAIAsset {
	s.RecognitionStatus = &v
	return s
}

func (s *TypedAIAsset) SetRiskInfo(v *TypedAIAssetRiskInfo) *TypedAIAsset {
	s.RiskInfo = v
	return s
}

func (s *TypedAIAsset) SetRiskLevel(v string) *TypedAIAsset {
	s.RiskLevel = &v
	return s
}

func (s *TypedAIAsset) SetScanId(v string) *TypedAIAsset {
	s.ScanId = &v
	return s
}

func (s *TypedAIAsset) SetSourceType(v string) *TypedAIAsset {
	s.SourceType = &v
	return s
}

func (s *TypedAIAsset) SetStandardSha256(v string) *TypedAIAsset {
	s.StandardSha256 = &v
	return s
}

func (s *TypedAIAsset) SetTransport(v string) *TypedAIAsset {
	s.Transport = &v
	return s
}

func (s *TypedAIAsset) SetVendor(v int32) *TypedAIAsset {
	s.Vendor = &v
	return s
}

func (s *TypedAIAsset) SetVersion(v string) *TypedAIAsset {
	s.Version = &v
	return s
}

func (s *TypedAIAsset) Validate() error {
	if s.IdentityIdentifiers != nil {
		if err := s.IdentityIdentifiers.Validate(); err != nil {
			return err
		}
	}
	if s.RecognitionEvidence != nil {
		if err := s.RecognitionEvidence.Validate(); err != nil {
			return err
		}
	}
	if s.RiskInfo != nil {
		if err := s.RiskInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TypedAIAssetIdentityIdentifiers struct {
	// The array of application names declared in desktop application manifests.
	//
	// example:
	//
	// ["Example Agent"]
	DesktopManifestName []*string `json:"desktop_manifest_name,omitempty" xml:"desktop_manifest_name,omitempty" type:"Repeated"`
	// The array of plugin IDs declared in JetBrains plugin manifests.
	//
	// example:
	//
	// ["com.example.agent.plugin"]
	JetbrainsPluginId []*string `json:"jetbrains_plugin_id,omitempty" xml:"jetbrains_plugin_id,omitempty" type:"Repeated"`
	// The array of macOS application bundle identifiers. These describe collected product identity clues and do not indicate that signatures are verified.
	//
	// example:
	//
	// ["com.example.agent"]
	MacosBundleId []*string `json:"macos_bundle_id,omitempty" xml:"macos_bundle_id,omitempty" type:"Repeated"`
	// The array of npm package names. Scoped packages retain the @scope/name format.
	//
	// example:
	//
	// ["@example/agent-cli"]
	NpmPackage []*string `json:"npm_package,omitempty" xml:"npm_package,omitempty" type:"Repeated"`
	// The array of product name and company name combinations in Windows PE version resources. These are not file hashes or verified signature identities.
	//
	// example:
	//
	// [{"product_name":"Example Agent","company_name":"Example Corp"}]
	PeProductIdentity []*TypedAIAssetIdentityIdentifiersPeProductIdentity `json:"pe_product_identity,omitempty" xml:"pe_product_identity,omitempty" type:"Repeated"`
	// The array of Python package distribution names. These are used for product identity identification and are not a list of import module names.
	//
	// example:
	//
	// ["example-agent-sdk"]
	PypiPackage []*string `json:"pypi_package,omitempty" xml:"pypi_package,omitempty" type:"Repeated"`
	// The array of VS Code extension identifiers. Each is typically the publisher name and extension name connected by a dot.
	//
	// example:
	//
	// ["example.agent-assistant"]
	VscodeExtensionId []*string `json:"vscode_extension_id,omitempty" xml:"vscode_extension_id,omitempty" type:"Repeated"`
	// The array of Windows AppX/MSIX manifest identities. Each item provides both the package name and publisher. This does not indicate that the publisher signature is verified.
	//
	// example:
	//
	// [{"name":"Example.Agent","publisher":"CN=Example Corp, O=Example Corp, C=US"}]
	WindowsAppxIdentity []*TypedAIAssetIdentityIdentifiersWindowsAppxIdentity `json:"windows_appx_identity,omitempty" xml:"windows_appx_identity,omitempty" type:"Repeated"`
	// The array of Windows application package family names. Each includes the package name and publisher identifier suffix and is not the application installation path.
	//
	// example:
	//
	// ["Example.Agent_123456789abcd"]
	WindowsPackageFamilyName []*string `json:"windows_package_family_name,omitempty" xml:"windows_package_family_name,omitempty" type:"Repeated"`
}

func (s TypedAIAssetIdentityIdentifiers) String() string {
	return dara.Prettify(s)
}

func (s TypedAIAssetIdentityIdentifiers) GoString() string {
	return s.String()
}

func (s *TypedAIAssetIdentityIdentifiers) GetDesktopManifestName() []*string {
	return s.DesktopManifestName
}

func (s *TypedAIAssetIdentityIdentifiers) GetJetbrainsPluginId() []*string {
	return s.JetbrainsPluginId
}

func (s *TypedAIAssetIdentityIdentifiers) GetMacosBundleId() []*string {
	return s.MacosBundleId
}

func (s *TypedAIAssetIdentityIdentifiers) GetNpmPackage() []*string {
	return s.NpmPackage
}

func (s *TypedAIAssetIdentityIdentifiers) GetPeProductIdentity() []*TypedAIAssetIdentityIdentifiersPeProductIdentity {
	return s.PeProductIdentity
}

func (s *TypedAIAssetIdentityIdentifiers) GetPypiPackage() []*string {
	return s.PypiPackage
}

func (s *TypedAIAssetIdentityIdentifiers) GetVscodeExtensionId() []*string {
	return s.VscodeExtensionId
}

func (s *TypedAIAssetIdentityIdentifiers) GetWindowsAppxIdentity() []*TypedAIAssetIdentityIdentifiersWindowsAppxIdentity {
	return s.WindowsAppxIdentity
}

func (s *TypedAIAssetIdentityIdentifiers) GetWindowsPackageFamilyName() []*string {
	return s.WindowsPackageFamilyName
}

func (s *TypedAIAssetIdentityIdentifiers) SetDesktopManifestName(v []*string) *TypedAIAssetIdentityIdentifiers {
	s.DesktopManifestName = v
	return s
}

func (s *TypedAIAssetIdentityIdentifiers) SetJetbrainsPluginId(v []*string) *TypedAIAssetIdentityIdentifiers {
	s.JetbrainsPluginId = v
	return s
}

func (s *TypedAIAssetIdentityIdentifiers) SetMacosBundleId(v []*string) *TypedAIAssetIdentityIdentifiers {
	s.MacosBundleId = v
	return s
}

func (s *TypedAIAssetIdentityIdentifiers) SetNpmPackage(v []*string) *TypedAIAssetIdentityIdentifiers {
	s.NpmPackage = v
	return s
}

func (s *TypedAIAssetIdentityIdentifiers) SetPeProductIdentity(v []*TypedAIAssetIdentityIdentifiersPeProductIdentity) *TypedAIAssetIdentityIdentifiers {
	s.PeProductIdentity = v
	return s
}

func (s *TypedAIAssetIdentityIdentifiers) SetPypiPackage(v []*string) *TypedAIAssetIdentityIdentifiers {
	s.PypiPackage = v
	return s
}

func (s *TypedAIAssetIdentityIdentifiers) SetVscodeExtensionId(v []*string) *TypedAIAssetIdentityIdentifiers {
	s.VscodeExtensionId = v
	return s
}

func (s *TypedAIAssetIdentityIdentifiers) SetWindowsAppxIdentity(v []*TypedAIAssetIdentityIdentifiersWindowsAppxIdentity) *TypedAIAssetIdentityIdentifiers {
	s.WindowsAppxIdentity = v
	return s
}

func (s *TypedAIAssetIdentityIdentifiers) SetWindowsPackageFamilyName(v []*string) *TypedAIAssetIdentityIdentifiers {
	s.WindowsPackageFamilyName = v
	return s
}

func (s *TypedAIAssetIdentityIdentifiers) Validate() error {
	if s.PeProductIdentity != nil {
		for _, item := range s.PeProductIdentity {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.WindowsAppxIdentity != nil {
		for _, item := range s.WindowsAppxIdentity {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TypedAIAssetIdentityIdentifiersPeProductIdentity struct {
	// The CompanyName field in a Windows PE version resource. This is not a certified vendor identity.
	//
	// example:
	//
	// Example Corp
	CompanyName *string `json:"company_name,omitempty" xml:"company_name,omitempty"`
	// The ProductName field in a Windows PE version resource.
	//
	// example:
	//
	// Example Agent
	ProductName *string `json:"product_name,omitempty" xml:"product_name,omitempty"`
}

func (s TypedAIAssetIdentityIdentifiersPeProductIdentity) String() string {
	return dara.Prettify(s)
}

func (s TypedAIAssetIdentityIdentifiersPeProductIdentity) GoString() string {
	return s.String()
}

func (s *TypedAIAssetIdentityIdentifiersPeProductIdentity) GetCompanyName() *string {
	return s.CompanyName
}

func (s *TypedAIAssetIdentityIdentifiersPeProductIdentity) GetProductName() *string {
	return s.ProductName
}

func (s *TypedAIAssetIdentityIdentifiersPeProductIdentity) SetCompanyName(v string) *TypedAIAssetIdentityIdentifiersPeProductIdentity {
	s.CompanyName = &v
	return s
}

func (s *TypedAIAssetIdentityIdentifiersPeProductIdentity) SetProductName(v string) *TypedAIAssetIdentityIdentifiersPeProductIdentity {
	s.ProductName = &v
	return s
}

func (s *TypedAIAssetIdentityIdentifiersPeProductIdentity) Validate() error {
	return dara.Validate(s)
}

type TypedAIAssetIdentityIdentifiersWindowsAppxIdentity struct {
	// The original package name of the Identity.Name in a Windows AppX/MSIX manifest.
	//
	// example:
	//
	// Example.Agent
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The publisher string of the Identity.Publisher in a Windows AppX/MSIX manifest. This is not equivalent to the signature verification result.
	//
	// example:
	//
	// CN=Example Corp, O=Example Corp, C=US
	Publisher *string `json:"publisher,omitempty" xml:"publisher,omitempty"`
}

func (s TypedAIAssetIdentityIdentifiersWindowsAppxIdentity) String() string {
	return dara.Prettify(s)
}

func (s TypedAIAssetIdentityIdentifiersWindowsAppxIdentity) GoString() string {
	return s.String()
}

func (s *TypedAIAssetIdentityIdentifiersWindowsAppxIdentity) GetName() *string {
	return s.Name
}

func (s *TypedAIAssetIdentityIdentifiersWindowsAppxIdentity) GetPublisher() *string {
	return s.Publisher
}

func (s *TypedAIAssetIdentityIdentifiersWindowsAppxIdentity) SetName(v string) *TypedAIAssetIdentityIdentifiersWindowsAppxIdentity {
	s.Name = &v
	return s
}

func (s *TypedAIAssetIdentityIdentifiersWindowsAppxIdentity) SetPublisher(v string) *TypedAIAssetIdentityIdentifiersWindowsAppxIdentity {
	s.Publisher = &v
	return s
}

func (s *TypedAIAssetIdentityIdentifiersWindowsAppxIdentity) Validate() error {
	return dara.Validate(s)
}

type TypedAIAssetRecognitionEvidence struct {
	// The array of evidence for explicit declarations of agent roles, capabilities, or usage methods. The declaration content itself does not constitute security certification.
	//
	// example:
	//
	// [{"source_path":"/opt/example-agent/package.json","content":"An AI agent that plans tasks and uses configured tools."}]
	AgentDeclaration []*TypedAIAssetRecognitionEvidenceAgentDeclaration `json:"agent_declaration,omitempty" xml:"agent_declaration,omitempty" type:"Repeated"`
	// The array of static evidence related to autonomous task execution or agent loops. This does not indicate that actual execution was observed during this scan.
	//
	// example:
	//
	// [{"source_path":"/opt/example-agent/agent.py"}]
	AutonomousExecution []*TypedAIAssetRecognitionEvidenceAutonomousExecution `json:"autonomous_execution,omitempty" xml:"autonomous_execution,omitempty" type:"Repeated"`
	// The array of static evidence related to model clients, requests, or configurations. This does not indicate that model requests were actually sent.
	//
	// example:
	//
	// [{"source_path":"/opt/example-agent/model.py"}]
	ModelCall []*TypedAIAssetRecognitionEvidenceModelCall `json:"model_call,omitempty" xml:"model_call,omitempty" type:"Repeated"`
	// The array of static evidence related to sessions, memory, or state management. This does not prove actual data persistence or runtime behavior.
	//
	// example:
	//
	// [{"source_path":"/opt/example-agent/session.py"}]
	StateManagement []*TypedAIAssetRecognitionEvidenceStateManagement `json:"state_management,omitempty" xml:"state_management,omitempty" type:"Repeated"`
	// The array of static evidence related to tool declarations or code to invoke tools. This does not indicate that the tools were actually executed.
	//
	// example:
	//
	// [{"source_path":"/opt/example-agent/tools.py"}]
	ToolCall []*TypedAIAssetRecognitionEvidenceToolCall `json:"tool_call,omitempty" xml:"tool_call,omitempty" type:"Repeated"`
}

func (s TypedAIAssetRecognitionEvidence) String() string {
	return dara.Prettify(s)
}

func (s TypedAIAssetRecognitionEvidence) GoString() string {
	return s.String()
}

func (s *TypedAIAssetRecognitionEvidence) GetAgentDeclaration() []*TypedAIAssetRecognitionEvidenceAgentDeclaration {
	return s.AgentDeclaration
}

func (s *TypedAIAssetRecognitionEvidence) GetAutonomousExecution() []*TypedAIAssetRecognitionEvidenceAutonomousExecution {
	return s.AutonomousExecution
}

func (s *TypedAIAssetRecognitionEvidence) GetModelCall() []*TypedAIAssetRecognitionEvidenceModelCall {
	return s.ModelCall
}

func (s *TypedAIAssetRecognitionEvidence) GetStateManagement() []*TypedAIAssetRecognitionEvidenceStateManagement {
	return s.StateManagement
}

func (s *TypedAIAssetRecognitionEvidence) GetToolCall() []*TypedAIAssetRecognitionEvidenceToolCall {
	return s.ToolCall
}

func (s *TypedAIAssetRecognitionEvidence) SetAgentDeclaration(v []*TypedAIAssetRecognitionEvidenceAgentDeclaration) *TypedAIAssetRecognitionEvidence {
	s.AgentDeclaration = v
	return s
}

func (s *TypedAIAssetRecognitionEvidence) SetAutonomousExecution(v []*TypedAIAssetRecognitionEvidenceAutonomousExecution) *TypedAIAssetRecognitionEvidence {
	s.AutonomousExecution = v
	return s
}

func (s *TypedAIAssetRecognitionEvidence) SetModelCall(v []*TypedAIAssetRecognitionEvidenceModelCall) *TypedAIAssetRecognitionEvidence {
	s.ModelCall = v
	return s
}

func (s *TypedAIAssetRecognitionEvidence) SetStateManagement(v []*TypedAIAssetRecognitionEvidenceStateManagement) *TypedAIAssetRecognitionEvidence {
	s.StateManagement = v
	return s
}

func (s *TypedAIAssetRecognitionEvidence) SetToolCall(v []*TypedAIAssetRecognitionEvidenceToolCall) *TypedAIAssetRecognitionEvidence {
	s.ToolCall = v
	return s
}

func (s *TypedAIAssetRecognitionEvidence) Validate() error {
	if s.AgentDeclaration != nil {
		for _, item := range s.AgentDeclaration {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.AutonomousExecution != nil {
		for _, item := range s.AutonomousExecution {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ModelCall != nil {
		for _, item := range s.ModelCall {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.StateManagement != nil {
		for _, item := range s.StateManagement {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ToolCall != nil {
		for _, item := range s.ToolCall {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TypedAIAssetRecognitionEvidenceAgentDeclaration struct {
	// The collected declaration text related to the recognition basis. This is currently used mainly for agent_declaration, and other recognition evidence may not return content. This should not be interpreted as an actual execution log.
	//
	// example:
	//
	// An AI agent that plans tasks and uses configured tools.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The local file or configuration path where the recognition evidence is located. This is used to display and locate static clues, and is not a server-side downloadable URL.
	//
	// example:
	//
	// /opt/example-agent/package.json
	SourcePath *string `json:"source_path,omitempty" xml:"source_path,omitempty"`
}

func (s TypedAIAssetRecognitionEvidenceAgentDeclaration) String() string {
	return dara.Prettify(s)
}

func (s TypedAIAssetRecognitionEvidenceAgentDeclaration) GoString() string {
	return s.String()
}

func (s *TypedAIAssetRecognitionEvidenceAgentDeclaration) GetContent() *string {
	return s.Content
}

func (s *TypedAIAssetRecognitionEvidenceAgentDeclaration) GetSourcePath() *string {
	return s.SourcePath
}

func (s *TypedAIAssetRecognitionEvidenceAgentDeclaration) SetContent(v string) *TypedAIAssetRecognitionEvidenceAgentDeclaration {
	s.Content = &v
	return s
}

func (s *TypedAIAssetRecognitionEvidenceAgentDeclaration) SetSourcePath(v string) *TypedAIAssetRecognitionEvidenceAgentDeclaration {
	s.SourcePath = &v
	return s
}

func (s *TypedAIAssetRecognitionEvidenceAgentDeclaration) Validate() error {
	return dara.Validate(s)
}

type TypedAIAssetRecognitionEvidenceAutonomousExecution struct {
	// The collected declaration text related to the recognition basis. This is currently used mainly for agent_declaration, and other recognition evidence may not return content. This should not be interpreted as an actual execution log.
	//
	// example:
	//
	// An AI agent that plans tasks and uses configured tools.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The local file or configuration path where the recognition evidence is located. This is used to display and locate static clues, and is not a server-side downloadable URL.
	//
	// example:
	//
	// /opt/example-agent/package.json
	SourcePath *string `json:"source_path,omitempty" xml:"source_path,omitempty"`
}

func (s TypedAIAssetRecognitionEvidenceAutonomousExecution) String() string {
	return dara.Prettify(s)
}

func (s TypedAIAssetRecognitionEvidenceAutonomousExecution) GoString() string {
	return s.String()
}

func (s *TypedAIAssetRecognitionEvidenceAutonomousExecution) GetContent() *string {
	return s.Content
}

func (s *TypedAIAssetRecognitionEvidenceAutonomousExecution) GetSourcePath() *string {
	return s.SourcePath
}

func (s *TypedAIAssetRecognitionEvidenceAutonomousExecution) SetContent(v string) *TypedAIAssetRecognitionEvidenceAutonomousExecution {
	s.Content = &v
	return s
}

func (s *TypedAIAssetRecognitionEvidenceAutonomousExecution) SetSourcePath(v string) *TypedAIAssetRecognitionEvidenceAutonomousExecution {
	s.SourcePath = &v
	return s
}

func (s *TypedAIAssetRecognitionEvidenceAutonomousExecution) Validate() error {
	return dara.Validate(s)
}

type TypedAIAssetRecognitionEvidenceModelCall struct {
	// The collected declaration text related to the recognition basis. This is currently used mainly for agent_declaration, and other recognition evidence may not return content. This should not be interpreted as an actual execution log.
	//
	// example:
	//
	// An AI agent that plans tasks and uses configured tools.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The local file or configuration path where the recognition evidence is located. This is used to display and locate static clues, and is not a server-side downloadable URL.
	//
	// example:
	//
	// /opt/example-agent/package.json
	SourcePath *string `json:"source_path,omitempty" xml:"source_path,omitempty"`
}

func (s TypedAIAssetRecognitionEvidenceModelCall) String() string {
	return dara.Prettify(s)
}

func (s TypedAIAssetRecognitionEvidenceModelCall) GoString() string {
	return s.String()
}

func (s *TypedAIAssetRecognitionEvidenceModelCall) GetContent() *string {
	return s.Content
}

func (s *TypedAIAssetRecognitionEvidenceModelCall) GetSourcePath() *string {
	return s.SourcePath
}

func (s *TypedAIAssetRecognitionEvidenceModelCall) SetContent(v string) *TypedAIAssetRecognitionEvidenceModelCall {
	s.Content = &v
	return s
}

func (s *TypedAIAssetRecognitionEvidenceModelCall) SetSourcePath(v string) *TypedAIAssetRecognitionEvidenceModelCall {
	s.SourcePath = &v
	return s
}

func (s *TypedAIAssetRecognitionEvidenceModelCall) Validate() error {
	return dara.Validate(s)
}

type TypedAIAssetRecognitionEvidenceStateManagement struct {
	// The collected declaration text related to the recognition basis. This is currently used mainly for agent_declaration, and other recognition evidence may not return content. This should not be interpreted as an actual execution log.
	//
	// example:
	//
	// An AI agent that plans tasks and uses configured tools.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The local file or configuration path where the recognition evidence is located. This is used to display and locate static clues, and is not a server-side downloadable URL.
	//
	// example:
	//
	// /opt/example-agent/package.json
	SourcePath *string `json:"source_path,omitempty" xml:"source_path,omitempty"`
}

func (s TypedAIAssetRecognitionEvidenceStateManagement) String() string {
	return dara.Prettify(s)
}

func (s TypedAIAssetRecognitionEvidenceStateManagement) GoString() string {
	return s.String()
}

func (s *TypedAIAssetRecognitionEvidenceStateManagement) GetContent() *string {
	return s.Content
}

func (s *TypedAIAssetRecognitionEvidenceStateManagement) GetSourcePath() *string {
	return s.SourcePath
}

func (s *TypedAIAssetRecognitionEvidenceStateManagement) SetContent(v string) *TypedAIAssetRecognitionEvidenceStateManagement {
	s.Content = &v
	return s
}

func (s *TypedAIAssetRecognitionEvidenceStateManagement) SetSourcePath(v string) *TypedAIAssetRecognitionEvidenceStateManagement {
	s.SourcePath = &v
	return s
}

func (s *TypedAIAssetRecognitionEvidenceStateManagement) Validate() error {
	return dara.Validate(s)
}

type TypedAIAssetRecognitionEvidenceToolCall struct {
	// The collected declaration text related to the recognition basis. This is currently used mainly for agent_declaration, and other recognition evidence may not return content. This should not be interpreted as an actual execution log.
	//
	// example:
	//
	// An AI agent that plans tasks and uses configured tools.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The local file or configuration path where the recognition evidence is located. This is used to display and locate static clues, and is not a server-side downloadable URL.
	//
	// example:
	//
	// /opt/example-agent/package.json
	SourcePath *string `json:"source_path,omitempty" xml:"source_path,omitempty"`
}

func (s TypedAIAssetRecognitionEvidenceToolCall) String() string {
	return dara.Prettify(s)
}

func (s TypedAIAssetRecognitionEvidenceToolCall) GoString() string {
	return s.String()
}

func (s *TypedAIAssetRecognitionEvidenceToolCall) GetContent() *string {
	return s.Content
}

func (s *TypedAIAssetRecognitionEvidenceToolCall) GetSourcePath() *string {
	return s.SourcePath
}

func (s *TypedAIAssetRecognitionEvidenceToolCall) SetContent(v string) *TypedAIAssetRecognitionEvidenceToolCall {
	s.Content = &v
	return s
}

func (s *TypedAIAssetRecognitionEvidenceToolCall) SetSourcePath(v string) *TypedAIAssetRecognitionEvidenceToolCall {
	s.SourcePath = &v
	return s
}

func (s *TypedAIAssetRecognitionEvidenceToolCall) Validate() error {
	return dara.Validate(s)
}

type TypedAIAssetRiskInfo struct {
	// The number of high-risk items in the currently saved detection results. This is interpreted as an actual count only when detection results are available.
	//
	// example:
	//
	// 1
	HighRiskCnt *int64 `json:"highRiskCnt,omitempty" xml:"highRiskCnt,omitempty"`
	// The number of low-risk items in the currently saved detection results. A value of zero indicates no low-risk items in the results but does not indicate overall security.
	//
	// example:
	//
	// 0
	LowRiskCnt *int64 `json:"lowRiskCnt,omitempty" xml:"lowRiskCnt,omitempty"`
	// The number of medium-risk items in the currently saved detection results. A value of zero indicates no medium-risk items in the results but does not indicate the absence of other risks.
	//
	// example:
	//
	// 0
	MediumRiskCnt *int64 `json:"mediumRiskCnt,omitempty" xml:"mediumRiskCnt,omitempty"`
}

func (s TypedAIAssetRiskInfo) String() string {
	return dara.Prettify(s)
}

func (s TypedAIAssetRiskInfo) GoString() string {
	return s.String()
}

func (s *TypedAIAssetRiskInfo) GetHighRiskCnt() *int64 {
	return s.HighRiskCnt
}

func (s *TypedAIAssetRiskInfo) GetLowRiskCnt() *int64 {
	return s.LowRiskCnt
}

func (s *TypedAIAssetRiskInfo) GetMediumRiskCnt() *int64 {
	return s.MediumRiskCnt
}

func (s *TypedAIAssetRiskInfo) SetHighRiskCnt(v int64) *TypedAIAssetRiskInfo {
	s.HighRiskCnt = &v
	return s
}

func (s *TypedAIAssetRiskInfo) SetLowRiskCnt(v int64) *TypedAIAssetRiskInfo {
	s.LowRiskCnt = &v
	return s
}

func (s *TypedAIAssetRiskInfo) SetMediumRiskCnt(v int64) *TypedAIAssetRiskInfo {
	s.MediumRiskCnt = &v
	return s
}

func (s *TypedAIAssetRiskInfo) Validate() error {
	return dara.Validate(s)
}
