// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateManagedAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateManagedAgentRequestBody) *UpdateManagedAgentRequest
	GetBody() *UpdateManagedAgentRequestBody
	SetClientToken(v string) *UpdateManagedAgentRequest
	GetClientToken() *string
}

type UpdateManagedAgentRequest struct {
	// The request body.
	Body *UpdateManagedAgentRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// The reserved idempotency token. The backend does not provide idempotency guarantees in the current release.
	//
	// example:
	//
	// client-token-1
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateManagedAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentRequest) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentRequest) GetBody() *UpdateManagedAgentRequestBody {
	return s.Body
}

func (s *UpdateManagedAgentRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateManagedAgentRequest) SetBody(v *UpdateManagedAgentRequestBody) *UpdateManagedAgentRequest {
	s.Body = v
	return s
}

func (s *UpdateManagedAgentRequest) SetClientToken(v string) *UpdateManagedAgentRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateManagedAgentRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateManagedAgentRequestBody struct {
	// Omit to retain existing values, pass [] to clear, or pass a non-empty array for full replacement. null is rejected. Combined with OSS mounts, a maximum of 10 items are allowed.
	AgenticFsMounts []*UpdateManagedAgentRequestBodyAgenticFsMounts `json:"agenticFsMounts,omitempty" xml:"agenticFsMounts,omitempty" type:"Repeated"`
	// The description of the managed agent.
	//
	// example:
	//
	// An agent for code review
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The environment configuration.
	Environment *UpdateManagedAgentRequestBodyEnvironment `json:"environment,omitempty" xml:"environment,omitempty" type:"Struct"`
	// The agent runtime harness.
	Harness *UpdateManagedAgentRequestBodyHarness `json:"harness,omitempty" xml:"harness,omitempty" type:"Struct"`
	// The agent instruction that guides the behavior of the agent.
	//
	// example:
	//
	// You are a code review assistant
	Instruction *string `json:"instruction,omitempty" xml:"instruction,omitempty"`
	// The model configuration.
	Model *UpdateManagedAgentRequestBodyModel `json:"model,omitempty" xml:"model,omitempty" type:"Struct"`
	// The name of the managed agent.
	//
	// example:
	//
	// my-agent
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The network configuration.
	Network *UpdateManagedAgentRequestBodyNetwork `json:"network,omitempty" xml:"network,omitempty" type:"Struct"`
	// The list of OSS mounts. A maximum of 10 items are allowed. Pass an empty array to clear existing mounts.
	OssMounts []*UpdateManagedAgentRequestBodyOssMounts `json:"ossMounts,omitempty" xml:"ossMounts,omitempty" type:"Repeated"`
	// The runtime configuration.
	Runtime *UpdateManagedAgentRequestBodyRuntime `json:"runtime,omitempty" xml:"runtime,omitempty" type:"Struct"`
	// The list of skill configurations.
	Skills []*UpdateManagedAgentRequestBodySkills `json:"skills,omitempty" xml:"skills,omitempty" type:"Repeated"`
	// The list of sub-agent configurations.
	SubAgents []*UpdateManagedAgentRequestBodySubAgents `json:"subAgents,omitempty" xml:"subAgents,omitempty" type:"Repeated"`
	// The agent template configuration.
	Template *UpdateManagedAgentRequestBodyTemplate `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
	// The list of tool configurations.
	Tools []*UpdateManagedAgentRequestBodyTools `json:"tools,omitempty" xml:"tools,omitempty" type:"Repeated"`
}

func (s UpdateManagedAgentRequestBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentRequestBody) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentRequestBody) GetAgenticFsMounts() []*UpdateManagedAgentRequestBodyAgenticFsMounts {
	return s.AgenticFsMounts
}

func (s *UpdateManagedAgentRequestBody) GetDescription() *string {
	return s.Description
}

func (s *UpdateManagedAgentRequestBody) GetEnvironment() *UpdateManagedAgentRequestBodyEnvironment {
	return s.Environment
}

func (s *UpdateManagedAgentRequestBody) GetHarness() *UpdateManagedAgentRequestBodyHarness {
	return s.Harness
}

func (s *UpdateManagedAgentRequestBody) GetInstruction() *string {
	return s.Instruction
}

func (s *UpdateManagedAgentRequestBody) GetModel() *UpdateManagedAgentRequestBodyModel {
	return s.Model
}

func (s *UpdateManagedAgentRequestBody) GetName() *string {
	return s.Name
}

func (s *UpdateManagedAgentRequestBody) GetNetwork() *UpdateManagedAgentRequestBodyNetwork {
	return s.Network
}

func (s *UpdateManagedAgentRequestBody) GetOssMounts() []*UpdateManagedAgentRequestBodyOssMounts {
	return s.OssMounts
}

func (s *UpdateManagedAgentRequestBody) GetRuntime() *UpdateManagedAgentRequestBodyRuntime {
	return s.Runtime
}

func (s *UpdateManagedAgentRequestBody) GetSkills() []*UpdateManagedAgentRequestBodySkills {
	return s.Skills
}

func (s *UpdateManagedAgentRequestBody) GetSubAgents() []*UpdateManagedAgentRequestBodySubAgents {
	return s.SubAgents
}

func (s *UpdateManagedAgentRequestBody) GetTemplate() *UpdateManagedAgentRequestBodyTemplate {
	return s.Template
}

func (s *UpdateManagedAgentRequestBody) GetTools() []*UpdateManagedAgentRequestBodyTools {
	return s.Tools
}

func (s *UpdateManagedAgentRequestBody) SetAgenticFsMounts(v []*UpdateManagedAgentRequestBodyAgenticFsMounts) *UpdateManagedAgentRequestBody {
	s.AgenticFsMounts = v
	return s
}

func (s *UpdateManagedAgentRequestBody) SetDescription(v string) *UpdateManagedAgentRequestBody {
	s.Description = &v
	return s
}

func (s *UpdateManagedAgentRequestBody) SetEnvironment(v *UpdateManagedAgentRequestBodyEnvironment) *UpdateManagedAgentRequestBody {
	s.Environment = v
	return s
}

func (s *UpdateManagedAgentRequestBody) SetHarness(v *UpdateManagedAgentRequestBodyHarness) *UpdateManagedAgentRequestBody {
	s.Harness = v
	return s
}

func (s *UpdateManagedAgentRequestBody) SetInstruction(v string) *UpdateManagedAgentRequestBody {
	s.Instruction = &v
	return s
}

func (s *UpdateManagedAgentRequestBody) SetModel(v *UpdateManagedAgentRequestBodyModel) *UpdateManagedAgentRequestBody {
	s.Model = v
	return s
}

func (s *UpdateManagedAgentRequestBody) SetName(v string) *UpdateManagedAgentRequestBody {
	s.Name = &v
	return s
}

func (s *UpdateManagedAgentRequestBody) SetNetwork(v *UpdateManagedAgentRequestBodyNetwork) *UpdateManagedAgentRequestBody {
	s.Network = v
	return s
}

func (s *UpdateManagedAgentRequestBody) SetOssMounts(v []*UpdateManagedAgentRequestBodyOssMounts) *UpdateManagedAgentRequestBody {
	s.OssMounts = v
	return s
}

func (s *UpdateManagedAgentRequestBody) SetRuntime(v *UpdateManagedAgentRequestBodyRuntime) *UpdateManagedAgentRequestBody {
	s.Runtime = v
	return s
}

func (s *UpdateManagedAgentRequestBody) SetSkills(v []*UpdateManagedAgentRequestBodySkills) *UpdateManagedAgentRequestBody {
	s.Skills = v
	return s
}

func (s *UpdateManagedAgentRequestBody) SetSubAgents(v []*UpdateManagedAgentRequestBodySubAgents) *UpdateManagedAgentRequestBody {
	s.SubAgents = v
	return s
}

func (s *UpdateManagedAgentRequestBody) SetTemplate(v *UpdateManagedAgentRequestBodyTemplate) *UpdateManagedAgentRequestBody {
	s.Template = v
	return s
}

func (s *UpdateManagedAgentRequestBody) SetTools(v []*UpdateManagedAgentRequestBodyTools) *UpdateManagedAgentRequestBody {
	s.Tools = v
	return s
}

func (s *UpdateManagedAgentRequestBody) Validate() error {
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

type UpdateManagedAgentRequestBodyAgenticFsMounts struct {
	// The subdirectory under /mnt/agenticfs/ in the container. Required for each mount item as validated by the backend. Mount targets must not be duplicated or have parent-child overlaps.
	//
	// example:
	//
	// /mnt/agenticfs/data
	MountPath *string `json:"mountPath,omitempty" xml:"mountPath,omitempty"`
	// A non-empty relative directory that exists under the AccessPoint. Required for each mount item as validated by the backend. Root directories, absolute paths, and parent directory segments are not allowed.
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
	// The AccessPoint domain name. Required for each mount item as validated by the backend. Do not include the protocol, port, or path. Use the DomainName value from the NAS ListAccessPoints response.
	//
	// example:
	//
	// ap-0123456789abcdef0.0123456789-vlm36.cn-hangzhou.nas.aliyuncs.com
	Server *string `json:"server,omitempty" xml:"server,omitempty"`
}

func (s UpdateManagedAgentRequestBodyAgenticFsMounts) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentRequestBodyAgenticFsMounts) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentRequestBodyAgenticFsMounts) GetMountPath() *string {
	return s.MountPath
}

func (s *UpdateManagedAgentRequestBodyAgenticFsMounts) GetPath() *string {
	return s.Path
}

func (s *UpdateManagedAgentRequestBodyAgenticFsMounts) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *UpdateManagedAgentRequestBodyAgenticFsMounts) GetServer() *string {
	return s.Server
}

func (s *UpdateManagedAgentRequestBodyAgenticFsMounts) SetMountPath(v string) *UpdateManagedAgentRequestBodyAgenticFsMounts {
	s.MountPath = &v
	return s
}

func (s *UpdateManagedAgentRequestBodyAgenticFsMounts) SetPath(v string) *UpdateManagedAgentRequestBodyAgenticFsMounts {
	s.Path = &v
	return s
}

func (s *UpdateManagedAgentRequestBodyAgenticFsMounts) SetReadOnly(v bool) *UpdateManagedAgentRequestBodyAgenticFsMounts {
	s.ReadOnly = &v
	return s
}

func (s *UpdateManagedAgentRequestBodyAgenticFsMounts) SetServer(v string) *UpdateManagedAgentRequestBodyAgenticFsMounts {
	s.Server = &v
	return s
}

func (s *UpdateManagedAgentRequestBodyAgenticFsMounts) Validate() error {
	return dara.Validate(s)
}

type UpdateManagedAgentRequestBodyEnvironment struct {
	// The list of credential references.
	CredentialReferences []*UpdateManagedAgentRequestBodyEnvironmentCredentialReferences `json:"credentialReferences,omitempty" xml:"credentialReferences,omitempty" type:"Repeated"`
	// The list of environment variables.
	Variables []*UpdateManagedAgentRequestBodyEnvironmentVariables `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
}

func (s UpdateManagedAgentRequestBodyEnvironment) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentRequestBodyEnvironment) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentRequestBodyEnvironment) GetCredentialReferences() []*UpdateManagedAgentRequestBodyEnvironmentCredentialReferences {
	return s.CredentialReferences
}

func (s *UpdateManagedAgentRequestBodyEnvironment) GetVariables() []*UpdateManagedAgentRequestBodyEnvironmentVariables {
	return s.Variables
}

func (s *UpdateManagedAgentRequestBodyEnvironment) SetCredentialReferences(v []*UpdateManagedAgentRequestBodyEnvironmentCredentialReferences) *UpdateManagedAgentRequestBodyEnvironment {
	s.CredentialReferences = v
	return s
}

func (s *UpdateManagedAgentRequestBodyEnvironment) SetVariables(v []*UpdateManagedAgentRequestBodyEnvironmentVariables) *UpdateManagedAgentRequestBodyEnvironment {
	s.Variables = v
	return s
}

func (s *UpdateManagedAgentRequestBodyEnvironment) Validate() error {
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

type UpdateManagedAgentRequestBodyEnvironmentCredentialReferences struct {
	// The credential ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cred-1
	CredentialId *string `json:"credentialId,omitempty" xml:"credentialId,omitempty"`
}

func (s UpdateManagedAgentRequestBodyEnvironmentCredentialReferences) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentRequestBodyEnvironmentCredentialReferences) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentRequestBodyEnvironmentCredentialReferences) GetCredentialId() *string {
	return s.CredentialId
}

func (s *UpdateManagedAgentRequestBodyEnvironmentCredentialReferences) SetCredentialId(v string) *UpdateManagedAgentRequestBodyEnvironmentCredentialReferences {
	s.CredentialId = &v
	return s
}

func (s *UpdateManagedAgentRequestBodyEnvironmentCredentialReferences) Validate() error {
	return dara.Validate(s)
}

type UpdateManagedAgentRequestBodyEnvironmentVariables struct {
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

func (s UpdateManagedAgentRequestBodyEnvironmentVariables) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentRequestBodyEnvironmentVariables) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentRequestBodyEnvironmentVariables) GetName() *string {
	return s.Name
}

func (s *UpdateManagedAgentRequestBodyEnvironmentVariables) GetValue() *string {
	return s.Value
}

func (s *UpdateManagedAgentRequestBodyEnvironmentVariables) SetName(v string) *UpdateManagedAgentRequestBodyEnvironmentVariables {
	s.Name = &v
	return s
}

func (s *UpdateManagedAgentRequestBodyEnvironmentVariables) SetValue(v string) *UpdateManagedAgentRequestBodyEnvironmentVariables {
	s.Value = &v
	return s
}

func (s *UpdateManagedAgentRequestBodyEnvironmentVariables) Validate() error {
	return dara.Validate(s)
}

type UpdateManagedAgentRequestBodyHarness struct {
	// The runtime harness configuration.
	Configuration *UpdateManagedAgentRequestBodyHarnessConfiguration `json:"configuration,omitempty" xml:"configuration,omitempty" type:"Struct"`
	// The runtime harness type.
	//
	// example:
	//
	// qodercli
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateManagedAgentRequestBodyHarness) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentRequestBodyHarness) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentRequestBodyHarness) GetConfiguration() *UpdateManagedAgentRequestBodyHarnessConfiguration {
	return s.Configuration
}

func (s *UpdateManagedAgentRequestBodyHarness) GetType() *string {
	return s.Type
}

func (s *UpdateManagedAgentRequestBodyHarness) SetConfiguration(v *UpdateManagedAgentRequestBodyHarnessConfiguration) *UpdateManagedAgentRequestBodyHarness {
	s.Configuration = v
	return s
}

func (s *UpdateManagedAgentRequestBodyHarness) SetType(v string) *UpdateManagedAgentRequestBodyHarness {
	s.Type = &v
	return s
}

func (s *UpdateManagedAgentRequestBodyHarness) Validate() error {
	if s.Configuration != nil {
		if err := s.Configuration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateManagedAgentRequestBodyHarnessConfiguration struct {
	// The Connector Service Account Key.
	//
	// example:
	//
	// key-xxxx
	ConnectorServiceAccountKey *string `json:"connectorServiceAccountKey,omitempty" xml:"connectorServiceAccountKey,omitempty"`
	// The Connector Service Account Name.
	//
	// example:
	//
	// my-connector-key
	ConnectorServiceAccountName *string `json:"connectorServiceAccountName,omitempty" xml:"connectorServiceAccountName,omitempty"`
}

func (s UpdateManagedAgentRequestBodyHarnessConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentRequestBodyHarnessConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentRequestBodyHarnessConfiguration) GetConnectorServiceAccountKey() *string {
	return s.ConnectorServiceAccountKey
}

func (s *UpdateManagedAgentRequestBodyHarnessConfiguration) GetConnectorServiceAccountName() *string {
	return s.ConnectorServiceAccountName
}

func (s *UpdateManagedAgentRequestBodyHarnessConfiguration) SetConnectorServiceAccountKey(v string) *UpdateManagedAgentRequestBodyHarnessConfiguration {
	s.ConnectorServiceAccountKey = &v
	return s
}

func (s *UpdateManagedAgentRequestBodyHarnessConfiguration) SetConnectorServiceAccountName(v string) *UpdateManagedAgentRequestBodyHarnessConfiguration {
	s.ConnectorServiceAccountName = &v
	return s
}

func (s *UpdateManagedAgentRequestBodyHarnessConfiguration) Validate() error {
	return dara.Validate(s)
}

type UpdateManagedAgentRequestBodyModel struct {
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
	Quota *UpdateManagedAgentRequestBodyModelQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
}

func (s UpdateManagedAgentRequestBodyModel) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentRequestBodyModel) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentRequestBodyModel) GetModelConnectionId() *string {
	return s.ModelConnectionId
}

func (s *UpdateManagedAgentRequestBodyModel) GetModelName() *string {
	return s.ModelName
}

func (s *UpdateManagedAgentRequestBodyModel) GetQuota() *UpdateManagedAgentRequestBodyModelQuota {
	return s.Quota
}

func (s *UpdateManagedAgentRequestBodyModel) SetModelConnectionId(v string) *UpdateManagedAgentRequestBodyModel {
	s.ModelConnectionId = &v
	return s
}

func (s *UpdateManagedAgentRequestBodyModel) SetModelName(v string) *UpdateManagedAgentRequestBodyModel {
	s.ModelName = &v
	return s
}

func (s *UpdateManagedAgentRequestBodyModel) SetQuota(v *UpdateManagedAgentRequestBodyModelQuota) *UpdateManagedAgentRequestBodyModel {
	s.Quota = v
	return s
}

func (s *UpdateManagedAgentRequestBodyModel) Validate() error {
	if s.Quota != nil {
		if err := s.Quota.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateManagedAgentRequestBodyModelQuota struct {
	// Specifies whether to enable the token quota. Default value: true. Set to false to disable and delete existing quota rules.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The quota limit type. Required when the quota is enabled, as validated by the backend. Fixed value: token.
	//
	// example:
	//
	// token
	LimitType *string `json:"limitType,omitempty" xml:"limitType,omitempty"`
	// The statistical period for the quota. Required when the quota is enabled, as validated by the backend. Valid values: day (daily) and month (monthly).
	//
	// example:
	//
	// day
	PeriodType *string `json:"periodType,omitempty" xml:"periodType,omitempty"`
	// The maximum number of tokens that can be consumed within a single period. Required when the quota is enabled, as validated by the backend. The value must be greater than 0.
	//
	// example:
	//
	// 1000000
	UsageLimit *int64 `json:"usageLimit,omitempty" xml:"usageLimit,omitempty"`
}

func (s UpdateManagedAgentRequestBodyModelQuota) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentRequestBodyModelQuota) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentRequestBodyModelQuota) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateManagedAgentRequestBodyModelQuota) GetLimitType() *string {
	return s.LimitType
}

func (s *UpdateManagedAgentRequestBodyModelQuota) GetPeriodType() *string {
	return s.PeriodType
}

func (s *UpdateManagedAgentRequestBodyModelQuota) GetUsageLimit() *int64 {
	return s.UsageLimit
}

func (s *UpdateManagedAgentRequestBodyModelQuota) SetEnabled(v bool) *UpdateManagedAgentRequestBodyModelQuota {
	s.Enabled = &v
	return s
}

func (s *UpdateManagedAgentRequestBodyModelQuota) SetLimitType(v string) *UpdateManagedAgentRequestBodyModelQuota {
	s.LimitType = &v
	return s
}

func (s *UpdateManagedAgentRequestBodyModelQuota) SetPeriodType(v string) *UpdateManagedAgentRequestBodyModelQuota {
	s.PeriodType = &v
	return s
}

func (s *UpdateManagedAgentRequestBodyModelQuota) SetUsageLimit(v int64) *UpdateManagedAgentRequestBodyModelQuota {
	s.UsageLimit = &v
	return s
}

func (s *UpdateManagedAgentRequestBodyModelQuota) Validate() error {
	return dara.Validate(s)
}

type UpdateManagedAgentRequestBodyNetwork struct {
	// The public network access configuration.
	AccessInternet *UpdateManagedAgentRequestBodyNetworkAccessInternet `json:"accessInternet,omitempty" xml:"accessInternet,omitempty" type:"Struct"`
	// The VPC access configuration.
	AccessVpc *UpdateManagedAgentRequestBodyNetworkAccessVpc `json:"accessVpc,omitempty" xml:"accessVpc,omitempty" type:"Struct"`
}

func (s UpdateManagedAgentRequestBodyNetwork) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentRequestBodyNetwork) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentRequestBodyNetwork) GetAccessInternet() *UpdateManagedAgentRequestBodyNetworkAccessInternet {
	return s.AccessInternet
}

func (s *UpdateManagedAgentRequestBodyNetwork) GetAccessVpc() *UpdateManagedAgentRequestBodyNetworkAccessVpc {
	return s.AccessVpc
}

func (s *UpdateManagedAgentRequestBodyNetwork) SetAccessInternet(v *UpdateManagedAgentRequestBodyNetworkAccessInternet) *UpdateManagedAgentRequestBodyNetwork {
	s.AccessInternet = v
	return s
}

func (s *UpdateManagedAgentRequestBodyNetwork) SetAccessVpc(v *UpdateManagedAgentRequestBodyNetworkAccessVpc) *UpdateManagedAgentRequestBodyNetwork {
	s.AccessVpc = v
	return s
}

func (s *UpdateManagedAgentRequestBodyNetwork) Validate() error {
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

type UpdateManagedAgentRequestBodyNetworkAccessInternet struct {
	// Specifies whether to allow public network access.
	//
	// example:
	//
	// false
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s UpdateManagedAgentRequestBodyNetworkAccessInternet) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentRequestBodyNetworkAccessInternet) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentRequestBodyNetworkAccessInternet) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateManagedAgentRequestBodyNetworkAccessInternet) SetEnabled(v bool) *UpdateManagedAgentRequestBodyNetworkAccessInternet {
	s.Enabled = &v
	return s
}

func (s *UpdateManagedAgentRequestBodyNetworkAccessInternet) Validate() error {
	return dara.Validate(s)
}

type UpdateManagedAgentRequestBodyNetworkAccessVpc struct {
	// Specifies whether to allow VPC access.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s UpdateManagedAgentRequestBodyNetworkAccessVpc) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentRequestBodyNetworkAccessVpc) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentRequestBodyNetworkAccessVpc) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateManagedAgentRequestBodyNetworkAccessVpc) SetEnabled(v bool) *UpdateManagedAgentRequestBodyNetworkAccessVpc {
	s.Enabled = &v
	return s
}

func (s *UpdateManagedAgentRequestBodyNetworkAccessVpc) Validate() error {
	return dara.Validate(s)
}

type UpdateManagedAgentRequestBodyOssMounts struct {
	// The OSS bucket name. Required for each mount item as validated by the backend.
	//
	// example:
	//
	// bucket-001
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// The absolute mount path in the container. Required for each mount item as validated by the backend.
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

func (s UpdateManagedAgentRequestBodyOssMounts) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentRequestBodyOssMounts) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentRequestBodyOssMounts) GetBucketName() *string {
	return s.BucketName
}

func (s *UpdateManagedAgentRequestBodyOssMounts) GetMountPath() *string {
	return s.MountPath
}

func (s *UpdateManagedAgentRequestBodyOssMounts) GetPath() *string {
	return s.Path
}

func (s *UpdateManagedAgentRequestBodyOssMounts) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *UpdateManagedAgentRequestBodyOssMounts) SetBucketName(v string) *UpdateManagedAgentRequestBodyOssMounts {
	s.BucketName = &v
	return s
}

func (s *UpdateManagedAgentRequestBodyOssMounts) SetMountPath(v string) *UpdateManagedAgentRequestBodyOssMounts {
	s.MountPath = &v
	return s
}

func (s *UpdateManagedAgentRequestBodyOssMounts) SetPath(v string) *UpdateManagedAgentRequestBodyOssMounts {
	s.Path = &v
	return s
}

func (s *UpdateManagedAgentRequestBodyOssMounts) SetReadOnly(v bool) *UpdateManagedAgentRequestBodyOssMounts {
	s.ReadOnly = &v
	return s
}

func (s *UpdateManagedAgentRequestBodyOssMounts) Validate() error {
	return dara.Validate(s)
}

type UpdateManagedAgentRequestBodyRuntime struct {
	// The compute configuration.
	//
	// This parameter is required.
	Compute *UpdateManagedAgentRequestBodyRuntimeCompute `json:"compute,omitempty" xml:"compute,omitempty" type:"Struct"`
	// The Sandbox auto scaling and session configuration.
	Hpa *UpdateManagedAgentRequestBodyRuntimeHpa `json:"hpa,omitempty" xml:"hpa,omitempty" type:"Struct"`
	// The session policy configuration.
	//
	// This parameter is required.
	SessionPolicy *UpdateManagedAgentRequestBodyRuntimeSessionPolicy `json:"sessionPolicy,omitempty" xml:"sessionPolicy,omitempty" type:"Struct"`
}

func (s UpdateManagedAgentRequestBodyRuntime) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentRequestBodyRuntime) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentRequestBodyRuntime) GetCompute() *UpdateManagedAgentRequestBodyRuntimeCompute {
	return s.Compute
}

func (s *UpdateManagedAgentRequestBodyRuntime) GetHpa() *UpdateManagedAgentRequestBodyRuntimeHpa {
	return s.Hpa
}

func (s *UpdateManagedAgentRequestBodyRuntime) GetSessionPolicy() *UpdateManagedAgentRequestBodyRuntimeSessionPolicy {
	return s.SessionPolicy
}

func (s *UpdateManagedAgentRequestBodyRuntime) SetCompute(v *UpdateManagedAgentRequestBodyRuntimeCompute) *UpdateManagedAgentRequestBodyRuntime {
	s.Compute = v
	return s
}

func (s *UpdateManagedAgentRequestBodyRuntime) SetHpa(v *UpdateManagedAgentRequestBodyRuntimeHpa) *UpdateManagedAgentRequestBodyRuntime {
	s.Hpa = v
	return s
}

func (s *UpdateManagedAgentRequestBodyRuntime) SetSessionPolicy(v *UpdateManagedAgentRequestBodyRuntimeSessionPolicy) *UpdateManagedAgentRequestBodyRuntime {
	s.SessionPolicy = v
	return s
}

func (s *UpdateManagedAgentRequestBodyRuntime) Validate() error {
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

type UpdateManagedAgentRequestBodyRuntimeCompute struct {
	// The compute class.
	//
	// This parameter is required.
	//
	// example:
	//
	// STANDARD
	ComputeClass *string `json:"computeClass,omitempty" xml:"computeClass,omitempty"`
}

func (s UpdateManagedAgentRequestBodyRuntimeCompute) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentRequestBodyRuntimeCompute) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentRequestBodyRuntimeCompute) GetComputeClass() *string {
	return s.ComputeClass
}

func (s *UpdateManagedAgentRequestBodyRuntimeCompute) SetComputeClass(v string) *UpdateManagedAgentRequestBodyRuntimeCompute {
	s.ComputeClass = &v
	return s
}

func (s *UpdateManagedAgentRequestBodyRuntimeCompute) Validate() error {
	return dara.Validate(s)
}

type UpdateManagedAgentRequestBodyRuntimeHpa struct {
	// Specifies whether to enable auto scaling. Required when hpa is present, as validated by the backend.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The maximum number of active sessions per Sandbox. Required when hpa is present, as validated by the backend.
	//
	// example:
	//
	// 5
	MaxConcurrentSessionsPerSandbox *int32 `json:"maxConcurrentSessionsPerSandbox,omitempty" xml:"maxConcurrentSessionsPerSandbox,omitempty"`
	// The maximum number of Sandboxes. Required when HPA is enabled and must be greater than or equal to the minimum value.
	//
	// example:
	//
	// 3
	MaxSandboxCount *int32 `json:"maxSandboxCount,omitempty" xml:"maxSandboxCount,omitempty"`
	// The minimum number of Sandboxes. Required when HPA is enabled.
	//
	// example:
	//
	// 1
	MinSandboxCount *int32 `json:"minSandboxCount,omitempty" xml:"minSandboxCount,omitempty"`
	// The time-to-live (TTL) for a session after inactivity, in seconds. Required when hpa is present, as validated by the backend.
	//
	// example:
	//
	// 3600
	SessionTtlSeconds *int32 `json:"sessionTtlSeconds,omitempty" xml:"sessionTtlSeconds,omitempty"`
}

func (s UpdateManagedAgentRequestBodyRuntimeHpa) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentRequestBodyRuntimeHpa) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentRequestBodyRuntimeHpa) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateManagedAgentRequestBodyRuntimeHpa) GetMaxConcurrentSessionsPerSandbox() *int32 {
	return s.MaxConcurrentSessionsPerSandbox
}

func (s *UpdateManagedAgentRequestBodyRuntimeHpa) GetMaxSandboxCount() *int32 {
	return s.MaxSandboxCount
}

func (s *UpdateManagedAgentRequestBodyRuntimeHpa) GetMinSandboxCount() *int32 {
	return s.MinSandboxCount
}

func (s *UpdateManagedAgentRequestBodyRuntimeHpa) GetSessionTtlSeconds() *int32 {
	return s.SessionTtlSeconds
}

func (s *UpdateManagedAgentRequestBodyRuntimeHpa) SetEnabled(v bool) *UpdateManagedAgentRequestBodyRuntimeHpa {
	s.Enabled = &v
	return s
}

func (s *UpdateManagedAgentRequestBodyRuntimeHpa) SetMaxConcurrentSessionsPerSandbox(v int32) *UpdateManagedAgentRequestBodyRuntimeHpa {
	s.MaxConcurrentSessionsPerSandbox = &v
	return s
}

func (s *UpdateManagedAgentRequestBodyRuntimeHpa) SetMaxSandboxCount(v int32) *UpdateManagedAgentRequestBodyRuntimeHpa {
	s.MaxSandboxCount = &v
	return s
}

func (s *UpdateManagedAgentRequestBodyRuntimeHpa) SetMinSandboxCount(v int32) *UpdateManagedAgentRequestBodyRuntimeHpa {
	s.MinSandboxCount = &v
	return s
}

func (s *UpdateManagedAgentRequestBodyRuntimeHpa) SetSessionTtlSeconds(v int32) *UpdateManagedAgentRequestBodyRuntimeHpa {
	s.SessionTtlSeconds = &v
	return s
}

func (s *UpdateManagedAgentRequestBodyRuntimeHpa) Validate() error {
	return dara.Validate(s)
}

type UpdateManagedAgentRequestBodyRuntimeSessionPolicy struct {
	// The name of the HTTP header used for session affinity. This parameter takes effect only when sessionPolicy.type is set to ISOLATED_HEADER_FIELD.
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

func (s UpdateManagedAgentRequestBodyRuntimeSessionPolicy) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentRequestBodyRuntimeSessionPolicy) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentRequestBodyRuntimeSessionPolicy) GetHeaderName() *string {
	return s.HeaderName
}

func (s *UpdateManagedAgentRequestBodyRuntimeSessionPolicy) GetType() *string {
	return s.Type
}

func (s *UpdateManagedAgentRequestBodyRuntimeSessionPolicy) SetHeaderName(v string) *UpdateManagedAgentRequestBodyRuntimeSessionPolicy {
	s.HeaderName = &v
	return s
}

func (s *UpdateManagedAgentRequestBodyRuntimeSessionPolicy) SetType(v string) *UpdateManagedAgentRequestBodyRuntimeSessionPolicy {
	s.Type = &v
	return s
}

func (s *UpdateManagedAgentRequestBodyRuntimeSessionPolicy) Validate() error {
	return dara.Validate(s)
}

type UpdateManagedAgentRequestBodySkills struct {
	// The skill name.
	//
	// This parameter is required.
	//
	// example:
	//
	// code-analysis
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The skill source type. Valid values:
	//
	// - REFERENCE: referenced from AI Registry.
	//
	// - STATIC: statically bundled with the package.
	//
	// example:
	//
	// REFERENCE
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// The skill version.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The version selector for the reference. Defaults to LABEL/latest if omitted. Currently supports LABEL/latest.
	VersionSelector *UpdateManagedAgentRequestBodySkillsVersionSelector `json:"versionSelector,omitempty" xml:"versionSelector,omitempty" type:"Struct"`
}

func (s UpdateManagedAgentRequestBodySkills) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentRequestBodySkills) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentRequestBodySkills) GetName() *string {
	return s.Name
}

func (s *UpdateManagedAgentRequestBodySkills) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateManagedAgentRequestBodySkills) GetVersion() *string {
	return s.Version
}

func (s *UpdateManagedAgentRequestBodySkills) GetVersionSelector() *UpdateManagedAgentRequestBodySkillsVersionSelector {
	return s.VersionSelector
}

func (s *UpdateManagedAgentRequestBodySkills) SetName(v string) *UpdateManagedAgentRequestBodySkills {
	s.Name = &v
	return s
}

func (s *UpdateManagedAgentRequestBodySkills) SetSourceType(v string) *UpdateManagedAgentRequestBodySkills {
	s.SourceType = &v
	return s
}

func (s *UpdateManagedAgentRequestBodySkills) SetVersion(v string) *UpdateManagedAgentRequestBodySkills {
	s.Version = &v
	return s
}

func (s *UpdateManagedAgentRequestBodySkills) SetVersionSelector(v *UpdateManagedAgentRequestBodySkillsVersionSelector) *UpdateManagedAgentRequestBodySkills {
	s.VersionSelector = v
	return s
}

func (s *UpdateManagedAgentRequestBodySkills) Validate() error {
	if s.VersionSelector != nil {
		if err := s.VersionSelector.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateManagedAgentRequestBodySkillsVersionSelector struct {
	// The version selector type. Valid values:
	//
	// - LABEL: select by label.
	//
	// - VERSION: select by specific version.
	//
	// example:
	//
	// LABEL
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The selector value. If the type is LABEL, specify a label name such as latest. If the type is VERSION, specify a specific version number.
	//
	// example:
	//
	// latest
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UpdateManagedAgentRequestBodySkillsVersionSelector) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentRequestBodySkillsVersionSelector) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentRequestBodySkillsVersionSelector) GetType() *string {
	return s.Type
}

func (s *UpdateManagedAgentRequestBodySkillsVersionSelector) GetValue() *string {
	return s.Value
}

func (s *UpdateManagedAgentRequestBodySkillsVersionSelector) SetType(v string) *UpdateManagedAgentRequestBodySkillsVersionSelector {
	s.Type = &v
	return s
}

func (s *UpdateManagedAgentRequestBodySkillsVersionSelector) SetValue(v string) *UpdateManagedAgentRequestBodySkillsVersionSelector {
	s.Value = &v
	return s
}

func (s *UpdateManagedAgentRequestBodySkillsVersionSelector) Validate() error {
	return dara.Validate(s)
}

type UpdateManagedAgentRequestBodySubAgents struct {
	// The sub-agent instruction.
	//
	// This parameter is required.
	//
	// example:
	//
	// Please review the code
	Instruction *string `json:"instruction,omitempty" xml:"instruction,omitempty"`
	// The sub-agent name.
	//
	// This parameter is required.
	//
	// example:
	//
	// reviewer-agent
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The skills exclusively used by this sub-agent. Skill names must be unique within the same sub-agent. If this parameter is not specified or an empty array is passed, no skills are configured.
	Skills []*UpdateManagedAgentRequestBodySubAgentsSkills `json:"skills,omitempty" xml:"skills,omitempty" type:"Repeated"`
}

func (s UpdateManagedAgentRequestBodySubAgents) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentRequestBodySubAgents) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentRequestBodySubAgents) GetInstruction() *string {
	return s.Instruction
}

func (s *UpdateManagedAgentRequestBodySubAgents) GetName() *string {
	return s.Name
}

func (s *UpdateManagedAgentRequestBodySubAgents) GetSkills() []*UpdateManagedAgentRequestBodySubAgentsSkills {
	return s.Skills
}

func (s *UpdateManagedAgentRequestBodySubAgents) SetInstruction(v string) *UpdateManagedAgentRequestBodySubAgents {
	s.Instruction = &v
	return s
}

func (s *UpdateManagedAgentRequestBodySubAgents) SetName(v string) *UpdateManagedAgentRequestBodySubAgents {
	s.Name = &v
	return s
}

func (s *UpdateManagedAgentRequestBodySubAgents) SetSkills(v []*UpdateManagedAgentRequestBodySubAgentsSkills) *UpdateManagedAgentRequestBodySubAgents {
	s.Skills = v
	return s
}

func (s *UpdateManagedAgentRequestBodySubAgents) Validate() error {
	if s.Skills != nil {
		for _, item := range s.Skills {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateManagedAgentRequestBodySubAgentsSkills struct {
	// The skill name used by the sub-agent. Declared as optional for compatibility, but the backend validates that each entry is required.
	//
	// example:
	//
	// web-search
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The optional version number. If omitted, set to null, or left blank, the latest version is resolved.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s UpdateManagedAgentRequestBodySubAgentsSkills) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentRequestBodySubAgentsSkills) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentRequestBodySubAgentsSkills) GetName() *string {
	return s.Name
}

func (s *UpdateManagedAgentRequestBodySubAgentsSkills) GetVersion() *string {
	return s.Version
}

func (s *UpdateManagedAgentRequestBodySubAgentsSkills) SetName(v string) *UpdateManagedAgentRequestBodySubAgentsSkills {
	s.Name = &v
	return s
}

func (s *UpdateManagedAgentRequestBodySubAgentsSkills) SetVersion(v string) *UpdateManagedAgentRequestBodySubAgentsSkills {
	s.Version = &v
	return s
}

func (s *UpdateManagedAgentRequestBodySubAgentsSkills) Validate() error {
	return dara.Validate(s)
}

type UpdateManagedAgentRequestBodyTemplate struct {
	// The AI Registry template configuration.
	AiRegistry *UpdateManagedAgentRequestBodyTemplateAiRegistry `json:"aiRegistry,omitempty" xml:"aiRegistry,omitempty" type:"Struct"`
}

func (s UpdateManagedAgentRequestBodyTemplate) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentRequestBodyTemplate) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentRequestBodyTemplate) GetAiRegistry() *UpdateManagedAgentRequestBodyTemplateAiRegistry {
	return s.AiRegistry
}

func (s *UpdateManagedAgentRequestBodyTemplate) SetAiRegistry(v *UpdateManagedAgentRequestBodyTemplateAiRegistry) *UpdateManagedAgentRequestBodyTemplate {
	s.AiRegistry = v
	return s
}

func (s *UpdateManagedAgentRequestBodyTemplate) Validate() error {
	if s.AiRegistry != nil {
		if err := s.AiRegistry.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateManagedAgentRequestBodyTemplateAiRegistry struct {
	// The name of the template in AI Registry.
	//
	// This parameter is required.
	//
	// example:
	//
	// code-review-template
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The version of the template in AI Registry.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s UpdateManagedAgentRequestBodyTemplateAiRegistry) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentRequestBodyTemplateAiRegistry) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentRequestBodyTemplateAiRegistry) GetName() *string {
	return s.Name
}

func (s *UpdateManagedAgentRequestBodyTemplateAiRegistry) GetVersion() *string {
	return s.Version
}

func (s *UpdateManagedAgentRequestBodyTemplateAiRegistry) SetName(v string) *UpdateManagedAgentRequestBodyTemplateAiRegistry {
	s.Name = &v
	return s
}

func (s *UpdateManagedAgentRequestBodyTemplateAiRegistry) SetVersion(v string) *UpdateManagedAgentRequestBodyTemplateAiRegistry {
	s.Version = &v
	return s
}

func (s *UpdateManagedAgentRequestBodyTemplateAiRegistry) Validate() error {
	return dara.Validate(s)
}

type UpdateManagedAgentRequestBodyTools struct {
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

func (s UpdateManagedAgentRequestBodyTools) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentRequestBodyTools) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentRequestBodyTools) GetName() *string {
	return s.Name
}

func (s *UpdateManagedAgentRequestBodyTools) GetType() *string {
	return s.Type
}

func (s *UpdateManagedAgentRequestBodyTools) SetName(v string) *UpdateManagedAgentRequestBodyTools {
	s.Name = &v
	return s
}

func (s *UpdateManagedAgentRequestBodyTools) SetType(v string) *UpdateManagedAgentRequestBodyTools {
	s.Type = &v
	return s
}

func (s *UpdateManagedAgentRequestBodyTools) Validate() error {
	return dara.Validate(s)
}
