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
	// A reserved query parameter. The backend currently does not provide idempotency guarantees, such as deduplication or recovery of the original processing result based on this token.
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
	// The inbound access mode.
	//
	// example:
	//
	// API_KEY
	AccessMode *string `json:"accessMode,omitempty" xml:"accessMode,omitempty"`
	// The user access scope.
	//
	// example:
	//
	// ALL
	AccessScope *string `json:"accessScope,omitempty" xml:"accessScope,omitempty"`
	// The NAS storage mount configuration. Currently, only the AgenticFS type is supported, which is used to mount file directories in access points to agent runtime instances. The list of AgenticFS directory mounts. The total number of AgenticFS mounts and `ossMounts` cannot exceed 10. If omitted or set to `[]` during creation, no mount is configured. If omitted during an update, the original mounts are retained. If set to `[]` during an update, the mounts are cleared. Null is not accepted.
	AgenticFsMounts []*CreateManagedAgentRequestBodyAgenticFsMounts `json:"agenticFsMounts,omitempty" xml:"agenticFsMounts,omitempty" type:"Repeated"`
	// The optional JSON object. Only `enabled` and `bindings` are allowed. Null, JSON character strings, or the legacy `contextSpaceId` field is not accepted. If the entire object is omitted, the server-side default configurations are used, which is typically disabled. Pass `{}` or `{enabled:false}` to explicitly keep it disabled. If `enabled` is set to true, you must select at least one resource in the current workspace. Non-empty `bindings` are validated for resource ownership and readiness even when disabled. Platform-managed context channels are not configured through regular tools.
	//
	// if can be null:
	// false
	ContextConfiguration *CreateManagedAgentRequestBodyContextConfiguration `json:"contextConfiguration,omitempty" xml:"contextConfiguration,omitempty" type:"Struct"`
	// The description of the managed Harness agent.
	//
	// example:
	//
	// Agent for code review
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The environment configuration.
	Environment *CreateManagedAgentRequestBodyEnvironment `json:"environment,omitempty" xml:"environment,omitempty" type:"Struct"`
	// The runtime framework object. If omitted, the default value is qwenpaw. Supported values are qwenpaw and qodercli. The type cannot be changed after creation.
	Harness *CreateManagedAgentRequestBodyHarness `json:"harness,omitempty" xml:"harness,omitempty" type:"Struct"`
	// The agent instruction, which is used to guide the behavior of the agent.
	//
	// example:
	//
	// You are a code review assistant
	Instruction *string `json:"instruction,omitempty" xml:"instruction,omitempty"`
	// The model configuration.
	//
	// This parameter is required.
	Model *CreateManagedAgentRequestBodyModel `json:"model,omitempty" xml:"model,omitempty" type:"Struct"`
	// The name of the managed Harness agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-agent
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The runtime network configuration. If provided, you must include both accessInternet.enabled and accessVpc.enabled. Actual access is subject to workspace network restrictions.
	Network *CreateManagedAgentRequestBodyNetwork `json:"network,omitempty" xml:"network,omitempty" type:"Struct"`
	// The list of OSS directory mounts. The total number of items in this list and agenticFsMounts cannot exceed 10. Updates replace the entire list. Pass an empty array [] to clear the list.
	OssMounts []*CreateManagedAgentRequestBodyOssMounts `json:"ossMounts,omitempty" xml:"ossMounts,omitempty" type:"Repeated"`
	// The prompt variable declaration.
	PromptVariables []*CreateManagedAgentRequestBodyPromptVariables `json:"promptVariables,omitempty" xml:"promptVariables,omitempty" type:"Repeated"`
	// The runtime configuration. You must provide both compute and sessionPolicy. Omitted fields in hpa use default values.
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

func (s *CreateManagedAgentRequestBody) GetAccessMode() *string {
	return s.AccessMode
}

func (s *CreateManagedAgentRequestBody) GetAccessScope() *string {
	return s.AccessScope
}

func (s *CreateManagedAgentRequestBody) GetAgenticFsMounts() []*CreateManagedAgentRequestBodyAgenticFsMounts {
	return s.AgenticFsMounts
}

func (s *CreateManagedAgentRequestBody) GetContextConfiguration() *CreateManagedAgentRequestBodyContextConfiguration {
	return s.ContextConfiguration
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

func (s *CreateManagedAgentRequestBody) GetPromptVariables() []*CreateManagedAgentRequestBodyPromptVariables {
	return s.PromptVariables
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

func (s *CreateManagedAgentRequestBody) SetAccessMode(v string) *CreateManagedAgentRequestBody {
	s.AccessMode = &v
	return s
}

func (s *CreateManagedAgentRequestBody) SetAccessScope(v string) *CreateManagedAgentRequestBody {
	s.AccessScope = &v
	return s
}

func (s *CreateManagedAgentRequestBody) SetAgenticFsMounts(v []*CreateManagedAgentRequestBodyAgenticFsMounts) *CreateManagedAgentRequestBody {
	s.AgenticFsMounts = v
	return s
}

func (s *CreateManagedAgentRequestBody) SetContextConfiguration(v *CreateManagedAgentRequestBodyContextConfiguration) *CreateManagedAgentRequestBody {
	s.ContextConfiguration = v
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

func (s *CreateManagedAgentRequestBody) SetPromptVariables(v []*CreateManagedAgentRequestBodyPromptVariables) *CreateManagedAgentRequestBody {
	s.PromptVariables = v
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
	if s.ContextConfiguration != nil {
		if err := s.ContextConfiguration.Validate(); err != nil {
			return err
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
	if s.PromptVariables != nil {
		for _, item := range s.PromptVariables {
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
	// The absolute mount path in the container, such as `/data/train`. Mount targets cannot be duplicated or form parent-child overwrites, and must avoid platform-reserved directories.
	//
	// example:
	//
	// /mnt/agenticfs/data
	MountPath *string `json:"mountPath,omitempty" xml:"mountPath,omitempty"`
	// The existing non-empty relative directory under the access point. This is a required field validated by the backend for each mount item. Root directories, absolute paths, or parent directory segments are not allowed.
	//
	// example:
	//
	// workspace/data
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// Specifies whether the mount is read-only. If omitted, the default value is false. This is not the read-only policy of a RAM role.
	//
	// example:
	//
	// false
	ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
	// The access point domain name. This is a required field validated by the backend for each mount item. It does not contain a protocol, port, or path. Use the `DomainName` from the NAS `ListAccessPoints` response.
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

type CreateManagedAgentRequestBodyContextConfiguration struct {
	// The optional array. Null or null items are not accepted. Each item only allows `type` and `resourceId`. The combination of the same type and normalized resource ID cannot be duplicated. If omitted during creation, no resource is selected. If omitted during an update, the original selection is retained. A non-empty array replaces the entire selection. An empty array can only clear the selection when the merged `enabled` value is false. To disable and clear the selection at the same time, pass `enabled:false` and `bindings:[]`. Explicitly submitted non-empty selections are validated for resource ownership and readiness, including selections saved when disabled.
	//
	// if can be null:
	// false
	Bindings []*CreateManagedAgentRequestBodyContextConfigurationBindings `json:"bindings,omitempty" xml:"bindings,omitempty" type:"Repeated"`
	// The optional Boolean value. Null is not accepted. If this field is omitted during creation, the value is false. If it is omitted during an update, the current value is retained and cannot be automatically filled with false by the client. When updated to true, saved `bindings` can be reused. The merged enabled configuration must contain at least one resource. When updated to false, the resources are only disabled, and the original resource selection is not cleared.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s CreateManagedAgentRequestBodyContextConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentRequestBodyContextConfiguration) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentRequestBodyContextConfiguration) GetBindings() []*CreateManagedAgentRequestBodyContextConfigurationBindings {
	return s.Bindings
}

func (s *CreateManagedAgentRequestBodyContextConfiguration) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateManagedAgentRequestBodyContextConfiguration) SetBindings(v []*CreateManagedAgentRequestBodyContextConfigurationBindings) *CreateManagedAgentRequestBodyContextConfiguration {
	s.Bindings = v
	return s
}

func (s *CreateManagedAgentRequestBodyContextConfiguration) SetEnabled(v bool) *CreateManagedAgentRequestBodyContextConfiguration {
	s.Enabled = &v
	return s
}

func (s *CreateManagedAgentRequestBodyContextConfiguration) Validate() error {
	if s.Bindings != nil {
		for _, item := range s.Bindings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateManagedAgentRequestBodyContextConfigurationBindings struct {
	// The required field validated by the backend for each binding item. Null, blank, or control characters are not accepted. The backend trims leading and trailing spaces and then validates that the length is 1 to 128 characters. The first character must be a letter or digit, and the remaining characters can contain letters, digits, periods, hyphens, or underscores. Only resources in the current workspace can be referenced.
	//
	// example:
	//
	// kb-example
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The required field validated by the backend for each binding item. Only `KNOWLEDGE_BASE` (knowledge base) or `DATA_CONNECTOR` (data connector) is accepted. The value is case-sensitive. Null or blank values are not accepted.
	//
	// example:
	//
	// KNOWLEDGE_BASE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateManagedAgentRequestBodyContextConfigurationBindings) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentRequestBodyContextConfigurationBindings) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentRequestBodyContextConfigurationBindings) GetResourceId() *string {
	return s.ResourceId
}

func (s *CreateManagedAgentRequestBodyContextConfigurationBindings) GetType() *string {
	return s.Type
}

func (s *CreateManagedAgentRequestBodyContextConfigurationBindings) SetResourceId(v string) *CreateManagedAgentRequestBodyContextConfigurationBindings {
	s.ResourceId = &v
	return s
}

func (s *CreateManagedAgentRequestBodyContextConfigurationBindings) SetType(v string) *CreateManagedAgentRequestBodyContextConfigurationBindings {
	s.Type = &v
	return s
}

func (s *CreateManagedAgentRequestBodyContextConfigurationBindings) Validate() error {
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
	// The ID of the credential.
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
	// The runtime framework configuration.
	Configuration *CreateManagedAgentRequestBodyHarnessConfiguration `json:"configuration,omitempty" xml:"configuration,omitempty" type:"Struct"`
	// The runtime framework type. This field is required when `harness` is provided. Valid values: qwenpaw and qodercli.
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
	// The account key record ID in the qodercli connector. You can query it by using `ListConnectors`. It is not the original key. This parameter can be omitted if there is only one record, but must be specified if there are multiple records. Do not specify this parameter for qwenpaw.
	//
	// example:
	//
	// key-example
	ConnectorServiceAccountKey *string `json:"connectorServiceAccountKey,omitempty" xml:"connectorServiceAccountKey,omitempty"`
	// The connector account name bound to qodercli. Do not specify this parameter for qwenpaw.
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
	// For qwenpaw, use the workspace model connection ID. For qodercli, use the available model identifier returned by `ListConnectorModels`.
	//
	// This parameter is required.
	//
	// example:
	//
	// mc-1
	ModelConnectionId *string `json:"modelConnectionId,omitempty" xml:"modelConnectionId,omitempty"`
	// For qwenpaw, use the name of the upstream model created under the model connection. This parameter is required. For qodercli, this parameter can be omitted.
	//
	// example:
	//
	// qwen-max
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	// The model token quota configuration. If this parameter is not passed, no quota is configured.
	Quota *CreateManagedAgentRequestBodyModelQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
	// Valid values: off, low, medium, high, xhigh, and max. If not passed or set to null during creation, the default value is used. When updating and submitting the model configuration, if the thinking level is not passed or is null, it is restored to the default value. Except for off, which is converted to none, the value is passed to the runtime framework as is. The system does not validate model support or automatically downgrade the level. If the model does not support the specified level, adjust the level.
	//
	// example:
	//
	// high
	ThinkingLevel *string `json:"thinkingLevel,omitempty" xml:"thinkingLevel,omitempty"`
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

func (s *CreateManagedAgentRequestBodyModel) GetThinkingLevel() *string {
	return s.ThinkingLevel
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

func (s *CreateManagedAgentRequestBodyModel) SetThinkingLevel(v string) *CreateManagedAgentRequestBodyModel {
	s.ThinkingLevel = &v
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
	// Specifies whether to enable the token quota. If not specified, the default value is true. If set to false, the quota is disabled and existing quota rules are deleted.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The quota limit type. This field is required and validated by the backend when the quota is enabled. Fixed value: token.
	//
	// example:
	//
	// token
	LimitType *string `json:"limitType,omitempty" xml:"limitType,omitempty"`
	// The statistical period of the quota. This field is required and validated by the backend when the quota is enabled. Valid values: day (daily) and month (monthly).
	//
	// example:
	//
	// day
	PeriodType *string `json:"periodType,omitempty" xml:"periodType,omitempty"`
	// The maximum number of tokens allowed to be consumed in a single period. This field is required and validated by the backend when the quota is enabled. The value must be greater than 0.
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
	// Specifies whether to allow Internet access.
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
	// The name of the OSS bucket. This field is required for each mount item and is validated by the backend.
	//
	// example:
	//
	// bucket-001
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// The absolute mount path in the container, such as `/data/train`. Mount targets cannot be duplicated or form parent-child overwrites, and must avoid platform-reserved directories.
	//
	// example:
	//
	// /mnt/oss/datasets
	MountPath *string `json:"mountPath,omitempty" xml:"mountPath,omitempty"`
	// The non-empty relative object directory prefix in the bucket. This field is required for each mount item.
	//
	// example:
	//
	// datasets
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// Specifies whether the mount is read-only. Default value: false.
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

type CreateManagedAgentRequestBodyPromptVariables struct {
	// The default value.
	//
	// example:
	//
	// default
	DefaultValue *string `json:"defaultValue,omitempty" xml:"defaultValue,omitempty"`
	// The variable description.
	//
	// example:
	//
	// desc
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The variable name.
	//
	// example:
	//
	// brand
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateManagedAgentRequestBodyPromptVariables) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentRequestBodyPromptVariables) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentRequestBodyPromptVariables) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *CreateManagedAgentRequestBodyPromptVariables) GetDescription() *string {
	return s.Description
}

func (s *CreateManagedAgentRequestBodyPromptVariables) GetName() *string {
	return s.Name
}

func (s *CreateManagedAgentRequestBodyPromptVariables) SetDefaultValue(v string) *CreateManagedAgentRequestBodyPromptVariables {
	s.DefaultValue = &v
	return s
}

func (s *CreateManagedAgentRequestBodyPromptVariables) SetDescription(v string) *CreateManagedAgentRequestBodyPromptVariables {
	s.Description = &v
	return s
}

func (s *CreateManagedAgentRequestBodyPromptVariables) SetName(v string) *CreateManagedAgentRequestBodyPromptVariables {
	s.Name = &v
	return s
}

func (s *CreateManagedAgentRequestBodyPromptVariables) Validate() error {
	return dara.Validate(s)
}

type CreateManagedAgentRequestBodyRuntime struct {
	// The compute configuration.
	//
	// This parameter is required.
	Compute *CreateManagedAgentRequestBodyRuntimeCompute `json:"compute,omitempty" xml:"compute,omitempty" type:"Struct"`
	// The elastic scaling and session configuration. Omitted fields use default values. To retain the original policy during an update, submit the complete hpa object or omit the entire runtime object.
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
	// The compute specification.
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
	// Specifies whether to enable elastic scaling. Default value: false.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The maximum number of concurrent sessions allowed per instance. Valid values: 1 to 20. Default value: 20.
	//
	// example:
	//
	// 20
	MaxConcurrentSessionsPerSandbox *int32 `json:"maxConcurrentSessionsPerSandbox,omitempty" xml:"maxConcurrentSessionsPerSandbox,omitempty"`
	// The maximum number of instances. Valid values: 0 to 100. Default value: 10. The value cannot be less than minSandboxCount.
	//
	// example:
	//
	// 3
	MaxSandboxCount *int32 `json:"maxSandboxCount,omitempty" xml:"maxSandboxCount,omitempty"`
	// The minimum number of instances. Valid values: 0 to 100. Default value: 1. The value cannot be greater than maxSandboxCount.
	//
	// example:
	//
	// 1
	MinSandboxCount *int32 `json:"minSandboxCount,omitempty" xml:"minSandboxCount,omitempty"`
	// The maximum idle time for a session, in seconds. Valid values: 60 to 86400. Default value: 600.
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
	// The name of the request header used for session identification. Required when type is set to ISOLATED_HEADER_FIELD. Must be a valid HTTP request header name. Do not specify this parameter for other policy types.
	//
	// example:
	//
	// X-Session-Id
	HeaderName *string `json:"headerName,omitempty" xml:"headerName,omitempty"`
	// The type of the session policy.
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
	// The name of the skill.
	//
	// This parameter is required.
	//
	// example:
	//
	// code-analysis
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The source type of the skill.
	//
	// example:
	//
	// REFERENCE
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// The skill version field for backward compatibility with legacy requests. For new requests, use sourceType and versionSelector. Do not submit this field in combination with sourceType.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The version selector for the skill reference. Supports LABEL for tags and VERSION for specific versions. Defaults to LABEL/latest when omitted. Not applicable for STATIC sources.
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
	// The type of the version selector.
	//
	// example:
	//
	// LABEL
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The value of the selector. For LABEL, specify the tag name, such as latest. For VERSION, specify the exact version number.
	//
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
	// The instruction for the sub-agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// Review the code
	Instruction *string `json:"instruction,omitempty" xml:"instruction,omitempty"`
	// The name of the sub-agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// reviewer-agent
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The list of skills for the sub-agent.
	Skills []*CreateManagedAgentRequestBodySubAgentsSkills `json:"skills,omitempty" xml:"skills,omitempty" type:"Repeated"`
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

func (s *CreateManagedAgentRequestBodySubAgents) GetSkills() []*CreateManagedAgentRequestBodySubAgentsSkills {
	return s.Skills
}

func (s *CreateManagedAgentRequestBodySubAgents) SetInstruction(v string) *CreateManagedAgentRequestBodySubAgents {
	s.Instruction = &v
	return s
}

func (s *CreateManagedAgentRequestBodySubAgents) SetName(v string) *CreateManagedAgentRequestBodySubAgents {
	s.Name = &v
	return s
}

func (s *CreateManagedAgentRequestBodySubAgents) SetSkills(v []*CreateManagedAgentRequestBodySubAgentsSkills) *CreateManagedAgentRequestBodySubAgents {
	s.Skills = v
	return s
}

func (s *CreateManagedAgentRequestBodySubAgents) Validate() error {
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

type CreateManagedAgentRequestBodySubAgentsSkills struct {
	// The name of the skill.
	//
	// example:
	//
	// web-search
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The version of the skill.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s CreateManagedAgentRequestBodySubAgentsSkills) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentRequestBodySubAgentsSkills) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentRequestBodySubAgentsSkills) GetName() *string {
	return s.Name
}

func (s *CreateManagedAgentRequestBodySubAgentsSkills) GetVersion() *string {
	return s.Version
}

func (s *CreateManagedAgentRequestBodySubAgentsSkills) SetName(v string) *CreateManagedAgentRequestBodySubAgentsSkills {
	s.Name = &v
	return s
}

func (s *CreateManagedAgentRequestBodySubAgentsSkills) SetVersion(v string) *CreateManagedAgentRequestBodySubAgentsSkills {
	s.Version = &v
	return s
}

func (s *CreateManagedAgentRequestBodySubAgentsSkills) Validate() error {
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
	// The name of the tool.
	//
	// This parameter is required.
	//
	// example:
	//
	// code-reviewer
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The type of the tool.
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
