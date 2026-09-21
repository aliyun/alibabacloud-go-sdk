// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateManagedAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateManagedAgentRequestBody) *CreateManagedAgentRequest
	GetBody() *CreateManagedAgentRequestBody
	SetClientToken(v string) *CreateManagedAgentRequest
	GetClientToken() *string
}

type CreateManagedAgentRequest struct {
	// The request body.
	Body *CreateManagedAgentRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// The reserved idempotency token. The backend does not provide idempotency guarantees in the current version.
	//
	// example:
	//
	// client-token-1
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateManagedAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentRequest) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentRequest) GetBody() *CreateManagedAgentRequestBody {
	return s.Body
}

func (s *CreateManagedAgentRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateManagedAgentRequest) SetBody(v *CreateManagedAgentRequestBody) *CreateManagedAgentRequest {
	s.Body = v
	return s
}

func (s *CreateManagedAgentRequest) SetClientToken(v string) *CreateManagedAgentRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateManagedAgentRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateManagedAgentRequestBody struct {
	// Omit or set to [] during creation to indicate no AFS mounts. Set to null to reject. The total number of AFS and OSS mounts cannot exceed 10.
	AgenticFsMounts []*CreateManagedAgentRequestBodyAgenticFsMounts `json:"agenticFsMounts,omitempty" xml:"agenticFsMounts,omitempty" type:"Repeated"`
	// The description of the managed agent.
	//
	// example:
	//
	// An agent for code review
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The environment configuration.
	Environment *CreateManagedAgentRequestBodyEnvironment `json:"environment,omitempty" xml:"environment,omitempty" type:"Struct"`
	// The agent runtime harness.
	Harness *CreateManagedAgentRequestBodyHarness `json:"harness,omitempty" xml:"harness,omitempty" type:"Struct"`
	// The agent instruction that guides the behavior of the agent.
	//
	// example:
	//
	// You are a code review assistant
	Instruction *string `json:"instruction,omitempty" xml:"instruction,omitempty"`
	// The model configuration.
	//
	// This parameter is required.
	Model *CreateManagedAgentRequestBodyModel `json:"model,omitempty" xml:"model,omitempty" type:"Struct"`
	// The name of the managed agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-agent
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The network configuration.
	Network *CreateManagedAgentRequestBodyNetwork `json:"network,omitempty" xml:"network,omitempty" type:"Struct"`
	// The OSS mount list. A maximum of 10 entries are allowed.
	OssMounts []*CreateManagedAgentRequestBodyOssMounts `json:"ossMounts,omitempty" xml:"ossMounts,omitempty" type:"Repeated"`
	// The runtime configuration.
	//
	// This parameter is required.
	Runtime *CreateManagedAgentRequestBodyRuntime `json:"runtime,omitempty" xml:"runtime,omitempty" type:"Struct"`
	// The list of skill configurations.
	Skills []*CreateManagedAgentRequestBodySkills `json:"skills,omitempty" xml:"skills,omitempty" type:"Repeated"`
	// The list of sub-agent configurations.
	SubAgents []*CreateManagedAgentRequestBodySubAgents `json:"subAgents,omitempty" xml:"subAgents,omitempty" type:"Repeated"`
	// The agent template configuration.
	Template *CreateManagedAgentRequestBodyTemplate `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
	// The list of tool configurations.
	Tools []*CreateManagedAgentRequestBodyTools `json:"tools,omitempty" xml:"tools,omitempty" type:"Repeated"`
}

func (s CreateManagedAgentRequestBody) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentRequestBody) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentRequestBody) GetAgenticFsMounts() []*CreateManagedAgentRequestBodyAgenticFsMounts {
	return s.AgenticFsMounts
}

func (s *CreateManagedAgentRequestBody) GetDescription() *string {
	return s.Description
}

func (s *CreateManagedAgentRequestBody) GetEnvironment() *CreateManagedAgentRequestBodyEnvironment {
	return s.Environment
}

func (s *CreateManagedAgentRequestBody) GetHarness() *CreateManagedAgentRequestBodyHarness {
	return s.Harness
}

func (s *CreateManagedAgentRequestBody) GetInstruction() *string {
	return s.Instruction
}

func (s *CreateManagedAgentRequestBody) GetModel() *CreateManagedAgentRequestBodyModel {
	return s.Model
}

func (s *CreateManagedAgentRequestBody) GetName() *string {
	return s.Name
}

func (s *CreateManagedAgentRequestBody) GetNetwork() *CreateManagedAgentRequestBodyNetwork {
	return s.Network
}

func (s *CreateManagedAgentRequestBody) GetOssMounts() []*CreateManagedAgentRequestBodyOssMounts {
	return s.OssMounts
}

func (s *CreateManagedAgentRequestBody) GetRuntime() *CreateManagedAgentRequestBodyRuntime {
	return s.Runtime
}

func (s *CreateManagedAgentRequestBody) GetSkills() []*CreateManagedAgentRequestBodySkills {
	return s.Skills
}

func (s *CreateManagedAgentRequestBody) GetSubAgents() []*CreateManagedAgentRequestBodySubAgents {
	return s.SubAgents
}

func (s *CreateManagedAgentRequestBody) GetTemplate() *CreateManagedAgentRequestBodyTemplate {
	return s.Template
}

func (s *CreateManagedAgentRequestBody) GetTools() []*CreateManagedAgentRequestBodyTools {
	return s.Tools
}

func (s *CreateManagedAgentRequestBody) SetAgenticFsMounts(v []*CreateManagedAgentRequestBodyAgenticFsMounts) *CreateManagedAgentRequestBody {
	s.AgenticFsMounts = v
	return s
}

func (s *CreateManagedAgentRequestBody) SetDescription(v string) *CreateManagedAgentRequestBody {
	s.Description = &v
	return s
}

func (s *CreateManagedAgentRequestBody) SetEnvironment(v *CreateManagedAgentRequestBodyEnvironment) *CreateManagedAgentRequestBody {
	s.Environment = v
	return s
}

func (s *CreateManagedAgentRequestBody) SetHarness(v *CreateManagedAgentRequestBodyHarness) *CreateManagedAgentRequestBody {
	s.Harness = v
	return s
}

func (s *CreateManagedAgentRequestBody) SetInstruction(v string) *CreateManagedAgentRequestBody {
	s.Instruction = &v
	return s
}

func (s *CreateManagedAgentRequestBody) SetModel(v *CreateManagedAgentRequestBodyModel) *CreateManagedAgentRequestBody {
	s.Model = v
	return s
}

func (s *CreateManagedAgentRequestBody) SetName(v string) *CreateManagedAgentRequestBody {
	s.Name = &v
	return s
}

func (s *CreateManagedAgentRequestBody) SetNetwork(v *CreateManagedAgentRequestBodyNetwork) *CreateManagedAgentRequestBody {
	s.Network = v
	return s
}

func (s *CreateManagedAgentRequestBody) SetOssMounts(v []*CreateManagedAgentRequestBodyOssMounts) *CreateManagedAgentRequestBody {
	s.OssMounts = v
	return s
}

func (s *CreateManagedAgentRequestBody) SetRuntime(v *CreateManagedAgentRequestBodyRuntime) *CreateManagedAgentRequestBody {
	s.Runtime = v
	return s
}

func (s *CreateManagedAgentRequestBody) SetSkills(v []*CreateManagedAgentRequestBodySkills) *CreateManagedAgentRequestBody {
	s.Skills = v
	return s
}

func (s *CreateManagedAgentRequestBody) SetSubAgents(v []*CreateManagedAgentRequestBodySubAgents) *CreateManagedAgentRequestBody {
	s.SubAgents = v
	return s
}

func (s *CreateManagedAgentRequestBody) SetTemplate(v *CreateManagedAgentRequestBodyTemplate) *CreateManagedAgentRequestBody {
	s.Template = v
	return s
}

func (s *CreateManagedAgentRequestBody) SetTools(v []*CreateManagedAgentRequestBodyTools) *CreateManagedAgentRequestBody {
	s.Tools = v
	return s
}

func (s *CreateManagedAgentRequestBody) Validate() error {
	if s.AgenticFsMounts != nil {
		for _, item := range s.AgenticFsMounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Environment != nil {
		if err := s.Environment.Validate(); err != nil {
			return err
		}
	}
	if s.Harness != nil {
		if err := s.Harness.Validate(); err != nil {
			return err
		}
	}
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
		}
	}
	if s.Network != nil {
		if err := s.Network.Validate(); err != nil {
			return err
		}
	}
	if s.OssMounts != nil {
		for _, item := range s.OssMounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Runtime != nil {
		if err := s.Runtime.Validate(); err != nil {
			return err
		}
	}
	if s.Skills != nil {
		for _, item := range s.Skills {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SubAgents != nil {
		for _, item := range s.SubAgents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Template != nil {
		if err := s.Template.Validate(); err != nil {
			return err
		}
	}
	if s.Tools != nil {
		for _, item := range s.Tools {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateManagedAgentRequestBodyAgenticFsMounts struct {
	// The subdirectory under /mnt/agenticfs/ in the container. This field is validated as required by the backend for each mount entry. Mount targets must not be duplicated or have parent-child overlaps.
	//
	// example:
	//
	// /mnt/agenticfs/data
	MountPath *string `json:"mountPath,omitempty" xml:"mountPath,omitempty"`
	// A non-empty relative directory that exists under the AccessPoint. This field is validated as required by the backend for each mount entry. Root directories, absolute paths, and parent directory segments are not allowed.
	//
	// example:
	//
	// workspace/data
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// Specifies whether to mount in read-only mode. Default value: false. This is not a RAM role read-only policy.
	//
	// example:
	//
	// false
	ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
	// The AccessPoint domain name. This field is validated as required by the backend for each mount entry. Do not include the protocol, port, or path. Use the DomainName value from the NAS ListAccessPoints response.
	//
	// example:
	//
	// ap-0123456789abcdef0.0123456789-vlm36.cn-hangzhou.nas.aliyuncs.com
	Server *string `json:"server,omitempty" xml:"server,omitempty"`
}

func (s CreateManagedAgentRequestBodyAgenticFsMounts) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentRequestBodyAgenticFsMounts) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentRequestBodyAgenticFsMounts) GetMountPath() *string {
	return s.MountPath
}

func (s *CreateManagedAgentRequestBodyAgenticFsMounts) GetPath() *string {
	return s.Path
}

func (s *CreateManagedAgentRequestBodyAgenticFsMounts) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *CreateManagedAgentRequestBodyAgenticFsMounts) GetServer() *string {
	return s.Server
}

func (s *CreateManagedAgentRequestBodyAgenticFsMounts) SetMountPath(v string) *CreateManagedAgentRequestBodyAgenticFsMounts {
	s.MountPath = &v
	return s
}

func (s *CreateManagedAgentRequestBodyAgenticFsMounts) SetPath(v string) *CreateManagedAgentRequestBodyAgenticFsMounts {
	s.Path = &v
	return s
}

func (s *CreateManagedAgentRequestBodyAgenticFsMounts) SetReadOnly(v bool) *CreateManagedAgentRequestBodyAgenticFsMounts {
	s.ReadOnly = &v
	return s
}

func (s *CreateManagedAgentRequestBodyAgenticFsMounts) SetServer(v string) *CreateManagedAgentRequestBodyAgenticFsMounts {
	s.Server = &v
	return s
}

func (s *CreateManagedAgentRequestBodyAgenticFsMounts) Validate() error {
	return dara.Validate(s)
}

type CreateManagedAgentRequestBodyEnvironment struct {
	// The list of credential references.
	CredentialReferences []*CreateManagedAgentRequestBodyEnvironmentCredentialReferences `json:"credentialReferences,omitempty" xml:"credentialReferences,omitempty" type:"Repeated"`
	// The list of environment variables.
	Variables []*CreateManagedAgentRequestBodyEnvironmentVariables `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
}

func (s CreateManagedAgentRequestBodyEnvironment) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentRequestBodyEnvironment) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentRequestBodyEnvironment) GetCredentialReferences() []*CreateManagedAgentRequestBodyEnvironmentCredentialReferences {
	return s.CredentialReferences
}

func (s *CreateManagedAgentRequestBodyEnvironment) GetVariables() []*CreateManagedAgentRequestBodyEnvironmentVariables {
	return s.Variables
}

func (s *CreateManagedAgentRequestBodyEnvironment) SetCredentialReferences(v []*CreateManagedAgentRequestBodyEnvironmentCredentialReferences) *CreateManagedAgentRequestBodyEnvironment {
	s.CredentialReferences = v
	return s
}

func (s *CreateManagedAgentRequestBodyEnvironment) SetVariables(v []*CreateManagedAgentRequestBodyEnvironmentVariables) *CreateManagedAgentRequestBodyEnvironment {
	s.Variables = v
	return s
}

func (s *CreateManagedAgentRequestBodyEnvironment) Validate() error {
	if s.CredentialReferences != nil {
		for _, item := range s.CredentialReferences {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Variables != nil {
		for _, item := range s.Variables {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateManagedAgentRequestBodyEnvironmentCredentialReferences struct {
	// The credential ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cred-1
	CredentialId *string `json:"credentialId,omitempty" xml:"credentialId,omitempty"`
}

func (s CreateManagedAgentRequestBodyEnvironmentCredentialReferences) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentRequestBodyEnvironmentCredentialReferences) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentRequestBodyEnvironmentCredentialReferences) GetCredentialId() *string {
	return s.CredentialId
}

func (s *CreateManagedAgentRequestBodyEnvironmentCredentialReferences) SetCredentialId(v string) *CreateManagedAgentRequestBodyEnvironmentCredentialReferences {
	s.CredentialId = &v
	return s
}

func (s *CreateManagedAgentRequestBodyEnvironmentCredentialReferences) Validate() error {
	return dara.Validate(s)
}

type CreateManagedAgentRequestBodyEnvironmentVariables struct {
	// The name of the environment variable.
	//
	// This parameter is required.
	//
	// example:
	//
	// API_KEY
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The value of the environment variable.
	//
	// This parameter is required.
	//
	// example:
	//
	// sk-xxxx
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateManagedAgentRequestBodyEnvironmentVariables) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentRequestBodyEnvironmentVariables) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentRequestBodyEnvironmentVariables) GetName() *string {
	return s.Name
}

func (s *CreateManagedAgentRequestBodyEnvironmentVariables) GetValue() *string {
	return s.Value
}

func (s *CreateManagedAgentRequestBodyEnvironmentVariables) SetName(v string) *CreateManagedAgentRequestBodyEnvironmentVariables {
	s.Name = &v
	return s
}

func (s *CreateManagedAgentRequestBodyEnvironmentVariables) SetValue(v string) *CreateManagedAgentRequestBodyEnvironmentVariables {
	s.Value = &v
	return s
}

func (s *CreateManagedAgentRequestBodyEnvironmentVariables) Validate() error {
	return dara.Validate(s)
}

type CreateManagedAgentRequestBodyHarness struct {
	// The runtime harness configuration.
	Configuration *CreateManagedAgentRequestBodyHarnessConfiguration `json:"configuration,omitempty" xml:"configuration,omitempty" type:"Struct"`
	// The runtime harness type.
	//
	// example:
	//
	// qodercli
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateManagedAgentRequestBodyHarness) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentRequestBodyHarness) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentRequestBodyHarness) GetConfiguration() *CreateManagedAgentRequestBodyHarnessConfiguration {
	return s.Configuration
}

func (s *CreateManagedAgentRequestBodyHarness) GetType() *string {
	return s.Type
}

func (s *CreateManagedAgentRequestBodyHarness) SetConfiguration(v *CreateManagedAgentRequestBodyHarnessConfiguration) *CreateManagedAgentRequestBodyHarness {
	s.Configuration = v
	return s
}

func (s *CreateManagedAgentRequestBodyHarness) SetType(v string) *CreateManagedAgentRequestBodyHarness {
	s.Type = &v
	return s
}

func (s *CreateManagedAgentRequestBodyHarness) Validate() error {
	if s.Configuration != nil {
		if err := s.Configuration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateManagedAgentRequestBodyHarnessConfiguration struct {
	// The connector service account key.
	//
	// example:
	//
	// key-xxxx
	ConnectorServiceAccountKey *string `json:"connectorServiceAccountKey,omitempty" xml:"connectorServiceAccountKey,omitempty"`
	// The connector service account name.
	//
	// example:
	//
	// my-connector-key
	ConnectorServiceAccountName *string `json:"connectorServiceAccountName,omitempty" xml:"connectorServiceAccountName,omitempty"`
}

func (s CreateManagedAgentRequestBodyHarnessConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentRequestBodyHarnessConfiguration) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentRequestBodyHarnessConfiguration) GetConnectorServiceAccountKey() *string {
	return s.ConnectorServiceAccountKey
}

func (s *CreateManagedAgentRequestBodyHarnessConfiguration) GetConnectorServiceAccountName() *string {
	return s.ConnectorServiceAccountName
}

func (s *CreateManagedAgentRequestBodyHarnessConfiguration) SetConnectorServiceAccountKey(v string) *CreateManagedAgentRequestBodyHarnessConfiguration {
	s.ConnectorServiceAccountKey = &v
	return s
}

func (s *CreateManagedAgentRequestBodyHarnessConfiguration) SetConnectorServiceAccountName(v string) *CreateManagedAgentRequestBodyHarnessConfiguration {
	s.ConnectorServiceAccountName = &v
	return s
}

func (s *CreateManagedAgentRequestBodyHarnessConfiguration) Validate() error {
	return dara.Validate(s)
}

type CreateManagedAgentRequestBodyModel struct {
	// The model connection ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// mc-1
	ModelConnectionId *string `json:"modelConnectionId,omitempty" xml:"modelConnectionId,omitempty"`
	// The upstream model name.
	//
	// example:
	//
	// qwen-max
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	// The model token quota configuration. If not specified, no quota is configured.
	Quota *CreateManagedAgentRequestBodyModelQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
}

func (s CreateManagedAgentRequestBodyModel) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentRequestBodyModel) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentRequestBodyModel) GetModelConnectionId() *string {
	return s.ModelConnectionId
}

func (s *CreateManagedAgentRequestBodyModel) GetModelName() *string {
	return s.ModelName
}

func (s *CreateManagedAgentRequestBodyModel) GetQuota() *CreateManagedAgentRequestBodyModelQuota {
	return s.Quota
}

func (s *CreateManagedAgentRequestBodyModel) SetModelConnectionId(v string) *CreateManagedAgentRequestBodyModel {
	s.ModelConnectionId = &v
	return s
}

func (s *CreateManagedAgentRequestBodyModel) SetModelName(v string) *CreateManagedAgentRequestBodyModel {
	s.ModelName = &v
	return s
}

func (s *CreateManagedAgentRequestBodyModel) SetQuota(v *CreateManagedAgentRequestBodyModelQuota) *CreateManagedAgentRequestBodyModel {
	s.Quota = v
	return s
}

func (s *CreateManagedAgentRequestBodyModel) Validate() error {
	if s.Quota != nil {
		if err := s.Quota.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateManagedAgentRequestBodyModelQuota struct {
	// Specifies whether to enable the token quota. Default value: true. Set to false to disable and delete existing quota rules.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The quota limit type. This field is validated as required by the backend when the quota is enabled. Fixed value: token.
	//
	// example:
	//
	// token
	LimitType *string `json:"limitType,omitempty" xml:"limitType,omitempty"`
	// The statistical period of the quota. This field is validated as required by the backend when the quota is enabled. Valid values:
	//
	// - day: daily.
	//
	// - month: monthly.
	//
	// example:
	//
	// day
	PeriodType *string `json:"periodType,omitempty" xml:"periodType,omitempty"`
	// The maximum number of tokens that can be consumed within a single period. This field is validated as required by the backend when the quota is enabled. The value must be greater than 0.
	//
	// example:
	//
	// 1000000
	UsageLimit *int64 `json:"usageLimit,omitempty" xml:"usageLimit,omitempty"`
}

func (s CreateManagedAgentRequestBodyModelQuota) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentRequestBodyModelQuota) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentRequestBodyModelQuota) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateManagedAgentRequestBodyModelQuota) GetLimitType() *string {
	return s.LimitType
}

func (s *CreateManagedAgentRequestBodyModelQuota) GetPeriodType() *string {
	return s.PeriodType
}

func (s *CreateManagedAgentRequestBodyModelQuota) GetUsageLimit() *int64 {
	return s.UsageLimit
}

func (s *CreateManagedAgentRequestBodyModelQuota) SetEnabled(v bool) *CreateManagedAgentRequestBodyModelQuota {
	s.Enabled = &v
	return s
}

func (s *CreateManagedAgentRequestBodyModelQuota) SetLimitType(v string) *CreateManagedAgentRequestBodyModelQuota {
	s.LimitType = &v
	return s
}

func (s *CreateManagedAgentRequestBodyModelQuota) SetPeriodType(v string) *CreateManagedAgentRequestBodyModelQuota {
	s.PeriodType = &v
	return s
}

func (s *CreateManagedAgentRequestBodyModelQuota) SetUsageLimit(v int64) *CreateManagedAgentRequestBodyModelQuota {
	s.UsageLimit = &v
	return s
}

func (s *CreateManagedAgentRequestBodyModelQuota) Validate() error {
	return dara.Validate(s)
}

type CreateManagedAgentRequestBodyNetwork struct {
	// The public network access configuration.
	AccessInternet *CreateManagedAgentRequestBodyNetworkAccessInternet `json:"accessInternet,omitempty" xml:"accessInternet,omitempty" type:"Struct"`
	// The VPC access configuration.
	AccessVpc *CreateManagedAgentRequestBodyNetworkAccessVpc `json:"accessVpc,omitempty" xml:"accessVpc,omitempty" type:"Struct"`
}

func (s CreateManagedAgentRequestBodyNetwork) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentRequestBodyNetwork) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentRequestBodyNetwork) GetAccessInternet() *CreateManagedAgentRequestBodyNetworkAccessInternet {
	return s.AccessInternet
}

func (s *CreateManagedAgentRequestBodyNetwork) GetAccessVpc() *CreateManagedAgentRequestBodyNetworkAccessVpc {
	return s.AccessVpc
}

func (s *CreateManagedAgentRequestBodyNetwork) SetAccessInternet(v *CreateManagedAgentRequestBodyNetworkAccessInternet) *CreateManagedAgentRequestBodyNetwork {
	s.AccessInternet = v
	return s
}

func (s *CreateManagedAgentRequestBodyNetwork) SetAccessVpc(v *CreateManagedAgentRequestBodyNetworkAccessVpc) *CreateManagedAgentRequestBodyNetwork {
	s.AccessVpc = v
	return s
}

func (s *CreateManagedAgentRequestBodyNetwork) Validate() error {
	if s.AccessInternet != nil {
		if err := s.AccessInternet.Validate(); err != nil {
			return err
		}
	}
	if s.AccessVpc != nil {
		if err := s.AccessVpc.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateManagedAgentRequestBodyNetworkAccessInternet struct {
	// Specifies whether to allow public network access.
	//
	// example:
	//
	// false
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s CreateManagedAgentRequestBodyNetworkAccessInternet) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentRequestBodyNetworkAccessInternet) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentRequestBodyNetworkAccessInternet) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateManagedAgentRequestBodyNetworkAccessInternet) SetEnabled(v bool) *CreateManagedAgentRequestBodyNetworkAccessInternet {
	s.Enabled = &v
	return s
}

func (s *CreateManagedAgentRequestBodyNetworkAccessInternet) Validate() error {
	return dara.Validate(s)
}

type CreateManagedAgentRequestBodyNetworkAccessVpc struct {
	// Specifies whether to allow VPC access.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s CreateManagedAgentRequestBodyNetworkAccessVpc) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentRequestBodyNetworkAccessVpc) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentRequestBodyNetworkAccessVpc) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateManagedAgentRequestBodyNetworkAccessVpc) SetEnabled(v bool) *CreateManagedAgentRequestBodyNetworkAccessVpc {
	s.Enabled = &v
	return s
}

func (s *CreateManagedAgentRequestBodyNetworkAccessVpc) Validate() error {
	return dara.Validate(s)
}

type CreateManagedAgentRequestBodyOssMounts struct {
	// The OSS bucket name. This field is validated as required by the backend for each mount entry.
	//
	// example:
	//
	// bucket-001
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// The absolute mount path in the container. This field is validated as required by the backend for each mount entry.
	//
	// example:
	//
	// /mnt/oss/datasets
	MountPath *string `json:"mountPath,omitempty" xml:"mountPath,omitempty"`
	// The relative object prefix within the bucket. If not specified, the entire bucket is mounted.
	//
	// example:
	//
	// datasets
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// Specifies whether to mount in read-only mode. Default value: false.
	ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
}

func (s CreateManagedAgentRequestBodyOssMounts) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentRequestBodyOssMounts) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentRequestBodyOssMounts) GetBucketName() *string {
	return s.BucketName
}

func (s *CreateManagedAgentRequestBodyOssMounts) GetMountPath() *string {
	return s.MountPath
}

func (s *CreateManagedAgentRequestBodyOssMounts) GetPath() *string {
	return s.Path
}

func (s *CreateManagedAgentRequestBodyOssMounts) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *CreateManagedAgentRequestBodyOssMounts) SetBucketName(v string) *CreateManagedAgentRequestBodyOssMounts {
	s.BucketName = &v
	return s
}

func (s *CreateManagedAgentRequestBodyOssMounts) SetMountPath(v string) *CreateManagedAgentRequestBodyOssMounts {
	s.MountPath = &v
	return s
}

func (s *CreateManagedAgentRequestBodyOssMounts) SetPath(v string) *CreateManagedAgentRequestBodyOssMounts {
	s.Path = &v
	return s
}

func (s *CreateManagedAgentRequestBodyOssMounts) SetReadOnly(v bool) *CreateManagedAgentRequestBodyOssMounts {
	s.ReadOnly = &v
	return s
}

func (s *CreateManagedAgentRequestBodyOssMounts) Validate() error {
	return dara.Validate(s)
}

type CreateManagedAgentRequestBodyRuntime struct {
	// The compute configuration.
	//
	// This parameter is required.
	Compute *CreateManagedAgentRequestBodyRuntimeCompute `json:"compute,omitempty" xml:"compute,omitempty" type:"Struct"`
	// The sandbox auto-scaling and session configuration.
	Hpa *CreateManagedAgentRequestBodyRuntimeHpa `json:"hpa,omitempty" xml:"hpa,omitempty" type:"Struct"`
	// The session policy configuration.
	//
	// This parameter is required.
	SessionPolicy *CreateManagedAgentRequestBodyRuntimeSessionPolicy `json:"sessionPolicy,omitempty" xml:"sessionPolicy,omitempty" type:"Struct"`
}

func (s CreateManagedAgentRequestBodyRuntime) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentRequestBodyRuntime) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentRequestBodyRuntime) GetCompute() *CreateManagedAgentRequestBodyRuntimeCompute {
	return s.Compute
}

func (s *CreateManagedAgentRequestBodyRuntime) GetHpa() *CreateManagedAgentRequestBodyRuntimeHpa {
	return s.Hpa
}

func (s *CreateManagedAgentRequestBodyRuntime) GetSessionPolicy() *CreateManagedAgentRequestBodyRuntimeSessionPolicy {
	return s.SessionPolicy
}

func (s *CreateManagedAgentRequestBodyRuntime) SetCompute(v *CreateManagedAgentRequestBodyRuntimeCompute) *CreateManagedAgentRequestBodyRuntime {
	s.Compute = v
	return s
}

func (s *CreateManagedAgentRequestBodyRuntime) SetHpa(v *CreateManagedAgentRequestBodyRuntimeHpa) *CreateManagedAgentRequestBodyRuntime {
	s.Hpa = v
	return s
}

func (s *CreateManagedAgentRequestBodyRuntime) SetSessionPolicy(v *CreateManagedAgentRequestBodyRuntimeSessionPolicy) *CreateManagedAgentRequestBodyRuntime {
	s.SessionPolicy = v
	return s
}

func (s *CreateManagedAgentRequestBodyRuntime) Validate() error {
	if s.Compute != nil {
		if err := s.Compute.Validate(); err != nil {
			return err
		}
	}
	if s.Hpa != nil {
		if err := s.Hpa.Validate(); err != nil {
			return err
		}
	}
	if s.SessionPolicy != nil {
		if err := s.SessionPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateManagedAgentRequestBodyRuntimeCompute struct {
	// The compute class.
	//
	// This parameter is required.
	//
	// example:
	//
	// STANDARD
	ComputeClass *string `json:"computeClass,omitempty" xml:"computeClass,omitempty"`
}

func (s CreateManagedAgentRequestBodyRuntimeCompute) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentRequestBodyRuntimeCompute) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentRequestBodyRuntimeCompute) GetComputeClass() *string {
	return s.ComputeClass
}

func (s *CreateManagedAgentRequestBodyRuntimeCompute) SetComputeClass(v string) *CreateManagedAgentRequestBodyRuntimeCompute {
	s.ComputeClass = &v
	return s
}

func (s *CreateManagedAgentRequestBodyRuntimeCompute) Validate() error {
	return dara.Validate(s)
}

type CreateManagedAgentRequestBodyRuntimeHpa struct {
	// Specifies whether to enable auto-scaling. This field is validated as required by the backend when hpa is present.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The maximum number of active sessions per sandbox. This field is validated as required by the backend when hpa is present.
	//
	// example:
	//
	// 5
	MaxConcurrentSessionsPerSandbox *int32 `json:"maxConcurrentSessionsPerSandbox,omitempty" xml:"maxConcurrentSessionsPerSandbox,omitempty"`
	// The maximum number of sandboxes. Required when HPA is enabled. The value must be greater than or equal to the minimum value.
	//
	// example:
	//
	// 3
	MaxSandboxCount *int32 `json:"maxSandboxCount,omitempty" xml:"maxSandboxCount,omitempty"`
	// The minimum number of sandboxes. Required when HPA is enabled.
	//
	// example:
	//
	// 1
	MinSandboxCount *int32 `json:"minSandboxCount,omitempty" xml:"minSandboxCount,omitempty"`
	// The time-to-live (TTL) for a session after inactivity, in seconds. This field is validated as required by the backend when hpa is present.
	//
	// example:
	//
	// 3600
	SessionTtlSeconds *int32 `json:"sessionTtlSeconds,omitempty" xml:"sessionTtlSeconds,omitempty"`
}

func (s CreateManagedAgentRequestBodyRuntimeHpa) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentRequestBodyRuntimeHpa) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentRequestBodyRuntimeHpa) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateManagedAgentRequestBodyRuntimeHpa) GetMaxConcurrentSessionsPerSandbox() *int32 {
	return s.MaxConcurrentSessionsPerSandbox
}

func (s *CreateManagedAgentRequestBodyRuntimeHpa) GetMaxSandboxCount() *int32 {
	return s.MaxSandboxCount
}

func (s *CreateManagedAgentRequestBodyRuntimeHpa) GetMinSandboxCount() *int32 {
	return s.MinSandboxCount
}

func (s *CreateManagedAgentRequestBodyRuntimeHpa) GetSessionTtlSeconds() *int32 {
	return s.SessionTtlSeconds
}

func (s *CreateManagedAgentRequestBodyRuntimeHpa) SetEnabled(v bool) *CreateManagedAgentRequestBodyRuntimeHpa {
	s.Enabled = &v
	return s
}

func (s *CreateManagedAgentRequestBodyRuntimeHpa) SetMaxConcurrentSessionsPerSandbox(v int32) *CreateManagedAgentRequestBodyRuntimeHpa {
	s.MaxConcurrentSessionsPerSandbox = &v
	return s
}

func (s *CreateManagedAgentRequestBodyRuntimeHpa) SetMaxSandboxCount(v int32) *CreateManagedAgentRequestBodyRuntimeHpa {
	s.MaxSandboxCount = &v
	return s
}

func (s *CreateManagedAgentRequestBodyRuntimeHpa) SetMinSandboxCount(v int32) *CreateManagedAgentRequestBodyRuntimeHpa {
	s.MinSandboxCount = &v
	return s
}

func (s *CreateManagedAgentRequestBodyRuntimeHpa) SetSessionTtlSeconds(v int32) *CreateManagedAgentRequestBodyRuntimeHpa {
	s.SessionTtlSeconds = &v
	return s
}

func (s *CreateManagedAgentRequestBodyRuntimeHpa) Validate() error {
	return dara.Validate(s)
}

type CreateManagedAgentRequestBodyRuntimeSessionPolicy struct {
	// The name of the HTTP header used for session affinity. This parameter takes effect when sessionPolicy.type is set to ISOLATED_HEADER_FIELD.
	//
	// example:
	//
	// X-Session-Id
	HeaderName *string `json:"headerName,omitempty" xml:"headerName,omitempty"`
	// The session policy type.
	//
	// This parameter is required.
	//
	// example:
	//
	// DISABLED
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateManagedAgentRequestBodyRuntimeSessionPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentRequestBodyRuntimeSessionPolicy) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentRequestBodyRuntimeSessionPolicy) GetHeaderName() *string {
	return s.HeaderName
}

func (s *CreateManagedAgentRequestBodyRuntimeSessionPolicy) GetType() *string {
	return s.Type
}

func (s *CreateManagedAgentRequestBodyRuntimeSessionPolicy) SetHeaderName(v string) *CreateManagedAgentRequestBodyRuntimeSessionPolicy {
	s.HeaderName = &v
	return s
}

func (s *CreateManagedAgentRequestBodyRuntimeSessionPolicy) SetType(v string) *CreateManagedAgentRequestBodyRuntimeSessionPolicy {
	s.Type = &v
	return s
}

func (s *CreateManagedAgentRequestBodyRuntimeSessionPolicy) Validate() error {
	return dara.Validate(s)
}

type CreateManagedAgentRequestBodySkills struct {
	// The skill name.
	//
	// This parameter is required.
	//
	// example:
	//
	// code-analysis
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// REFERENCE
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// The skill version.
	//
	// example:
	//
	// 1.0.0
	Version         *string                                             `json:"version,omitempty" xml:"version,omitempty"`
	VersionSelector *CreateManagedAgentRequestBodySkillsVersionSelector `json:"versionSelector,omitempty" xml:"versionSelector,omitempty" type:"Struct"`
}

func (s CreateManagedAgentRequestBodySkills) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentRequestBodySkills) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentRequestBodySkills) GetName() *string {
	return s.Name
}

func (s *CreateManagedAgentRequestBodySkills) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateManagedAgentRequestBodySkills) GetVersion() *string {
	return s.Version
}

func (s *CreateManagedAgentRequestBodySkills) GetVersionSelector() *CreateManagedAgentRequestBodySkillsVersionSelector {
	return s.VersionSelector
}

func (s *CreateManagedAgentRequestBodySkills) SetName(v string) *CreateManagedAgentRequestBodySkills {
	s.Name = &v
	return s
}

func (s *CreateManagedAgentRequestBodySkills) SetSourceType(v string) *CreateManagedAgentRequestBodySkills {
	s.SourceType = &v
	return s
}

func (s *CreateManagedAgentRequestBodySkills) SetVersion(v string) *CreateManagedAgentRequestBodySkills {
	s.Version = &v
	return s
}

func (s *CreateManagedAgentRequestBodySkills) SetVersionSelector(v *CreateManagedAgentRequestBodySkillsVersionSelector) *CreateManagedAgentRequestBodySkills {
	s.VersionSelector = v
	return s
}

func (s *CreateManagedAgentRequestBodySkills) Validate() error {
	if s.VersionSelector != nil {
		if err := s.VersionSelector.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateManagedAgentRequestBodySkillsVersionSelector struct {
	// example:
	//
	// LABEL
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// latest
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateManagedAgentRequestBodySkillsVersionSelector) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentRequestBodySkillsVersionSelector) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentRequestBodySkillsVersionSelector) GetType() *string {
	return s.Type
}

func (s *CreateManagedAgentRequestBodySkillsVersionSelector) GetValue() *string {
	return s.Value
}

func (s *CreateManagedAgentRequestBodySkillsVersionSelector) SetType(v string) *CreateManagedAgentRequestBodySkillsVersionSelector {
	s.Type = &v
	return s
}

func (s *CreateManagedAgentRequestBodySkillsVersionSelector) SetValue(v string) *CreateManagedAgentRequestBodySkillsVersionSelector {
	s.Value = &v
	return s
}

func (s *CreateManagedAgentRequestBodySkillsVersionSelector) Validate() error {
	return dara.Validate(s)
}

type CreateManagedAgentRequestBodySubAgents struct {
	// The sub-agent instruction.
	//
	// This parameter is required.
	//
	// example:
	//
	// Review the code
	Instruction *string `json:"instruction,omitempty" xml:"instruction,omitempty"`
	// The sub-agent name.
	//
	// This parameter is required.
	//
	// example:
	//
	// reviewer-agent
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateManagedAgentRequestBodySubAgents) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentRequestBodySubAgents) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentRequestBodySubAgents) GetInstruction() *string {
	return s.Instruction
}

func (s *CreateManagedAgentRequestBodySubAgents) GetName() *string {
	return s.Name
}

func (s *CreateManagedAgentRequestBodySubAgents) SetInstruction(v string) *CreateManagedAgentRequestBodySubAgents {
	s.Instruction = &v
	return s
}

func (s *CreateManagedAgentRequestBodySubAgents) SetName(v string) *CreateManagedAgentRequestBodySubAgents {
	s.Name = &v
	return s
}

func (s *CreateManagedAgentRequestBodySubAgents) Validate() error {
	return dara.Validate(s)
}

type CreateManagedAgentRequestBodyTemplate struct {
	// The AI registry template configuration.
	AiRegistry *CreateManagedAgentRequestBodyTemplateAiRegistry `json:"aiRegistry,omitempty" xml:"aiRegistry,omitempty" type:"Struct"`
}

func (s CreateManagedAgentRequestBodyTemplate) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentRequestBodyTemplate) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentRequestBodyTemplate) GetAiRegistry() *CreateManagedAgentRequestBodyTemplateAiRegistry {
	return s.AiRegistry
}

func (s *CreateManagedAgentRequestBodyTemplate) SetAiRegistry(v *CreateManagedAgentRequestBodyTemplateAiRegistry) *CreateManagedAgentRequestBodyTemplate {
	s.AiRegistry = v
	return s
}

func (s *CreateManagedAgentRequestBodyTemplate) Validate() error {
	if s.AiRegistry != nil {
		if err := s.AiRegistry.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateManagedAgentRequestBodyTemplateAiRegistry struct {
	// The name of the template in the AI registry.
	//
	// This parameter is required.
	//
	// example:
	//
	// code-review-template
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The version of the template in the AI registry.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s CreateManagedAgentRequestBodyTemplateAiRegistry) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentRequestBodyTemplateAiRegistry) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentRequestBodyTemplateAiRegistry) GetName() *string {
	return s.Name
}

func (s *CreateManagedAgentRequestBodyTemplateAiRegistry) GetVersion() *string {
	return s.Version
}

func (s *CreateManagedAgentRequestBodyTemplateAiRegistry) SetName(v string) *CreateManagedAgentRequestBodyTemplateAiRegistry {
	s.Name = &v
	return s
}

func (s *CreateManagedAgentRequestBodyTemplateAiRegistry) SetVersion(v string) *CreateManagedAgentRequestBodyTemplateAiRegistry {
	s.Version = &v
	return s
}

func (s *CreateManagedAgentRequestBodyTemplateAiRegistry) Validate() error {
	return dara.Validate(s)
}

type CreateManagedAgentRequestBodyTools struct {
	// The tool name.
	//
	// This parameter is required.
	//
	// example:
	//
	// code-reviewer
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The tool type.
	//
	// This parameter is required.
	//
	// example:
	//
	// MCP
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateManagedAgentRequestBodyTools) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentRequestBodyTools) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentRequestBodyTools) GetName() *string {
	return s.Name
}

func (s *CreateManagedAgentRequestBodyTools) GetType() *string {
	return s.Type
}

func (s *CreateManagedAgentRequestBodyTools) SetName(v string) *CreateManagedAgentRequestBodyTools {
	s.Name = &v
	return s
}

func (s *CreateManagedAgentRequestBodyTools) SetType(v string) *CreateManagedAgentRequestBodyTools {
	s.Type = &v
	return s
}

func (s *CreateManagedAgentRequestBodyTools) Validate() error {
	return dara.Validate(s)
}
