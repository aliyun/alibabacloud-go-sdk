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
	// The reserved query parameter. The backend does not currently provide idempotency guarantees for deduplication or result recovery based on this token.
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
	// The NAS storage mount configuration. Currently, the AgenticFS type is supported, which is used to mount file directories in access points to agent runtime instances. The AgenticFS directory mount list can contain up to 10 items in total with ossMounts. Omitting this field or passing [] during creation indicates no mounts. During an update, omitting this field retains the original configuration, and passing [] clears it. Null is not accepted.
	AgenticFsMounts []*UpdateManagedAgentRequestBodyAgenticFsMounts `json:"agenticFsMounts,omitempty" xml:"agenticFsMounts,omitempty" type:"Repeated"`
	// An optional JSON object. Null, stringified JSON, or the legacy contextSpaceId field is not accepted. Omitting the entire object or passing {} retains the existing configuration. Omitting enabled or bindings within the object retains their original values respectively. Setting enabled to true can reuse existing non-empty selections, while setting enabled to false only disables and retains the selections. If bindings is a non-empty array, it is entirely replaced. If bindings is [], it is cleared only when the final state is disabled. To disable and clear, pass {enabled:false,bindings:[]}. Explicit null at any level is rejected, and default false or empty arrays must not be automatically populated.
	//
	// if can be null:
	// false
	ContextConfiguration *UpdateManagedAgentRequestBodyContextConfiguration `json:"contextConfiguration,omitempty" xml:"contextConfiguration,omitempty" type:"Struct"`
	// The description of the managed Harness agent.
	//
	// example:
	//
	// Agent for code review
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The environment configuration.
	Environment *UpdateManagedAgentRequestBodyEnvironment `json:"environment,omitempty" xml:"environment,omitempty" type:"Struct"`
	// The runtime framework object. If omitted, the original framework is retained. If provided, the type must be the same as that specified during creation. You can update the account binding for the same type of framework.
	Harness *UpdateManagedAgentRequestBodyHarness `json:"harness,omitempty" xml:"harness,omitempty" type:"Struct"`
	// The agent instructions used to guide the behavior of the agent.
	//
	// example:
	//
	// You are a code review assistant
	Instruction *string `json:"instruction,omitempty" xml:"instruction,omitempty"`
	// The model configuration.
	Model *UpdateManagedAgentRequestBodyModel `json:"model,omitempty" xml:"model,omitempty" type:"Struct"`
	// The name of the managed Harness agent.
	//
	// example:
	//
	// my-agent
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The runtime network configuration. If provided, both accessInternet.enabled and accessVpc.enabled must be included. Actual access is subject to workspace network restrictions.
	Network *UpdateManagedAgentRequestBodyNetwork `json:"network,omitempty" xml:"network,omitempty" type:"Struct"`
	// The list of OSS directory mounts. A maximum of 10 items in total with agenticFsMounts is allowed. The entire list is replaced during an update. Pass an empty array [] to clear the list.
	OssMounts []*UpdateManagedAgentRequestBodyOssMounts `json:"ossMounts,omitempty" xml:"ossMounts,omitempty" type:"Repeated"`
	// The prompt variable declarations.
	PromptVariables []*UpdateManagedAgentRequestBodyPromptVariables `json:"promptVariables,omitempty" xml:"promptVariables,omitempty" type:"Repeated"`
	// The runtime configuration. If the entire object is omitted, the original configuration is retained. If provided, both compute and sessionPolicy must be specified. Omitted fields in hpa use default values.
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

func (s *UpdateManagedAgentRequestBody) GetAccessMode() *string {
	return s.AccessMode
}

func (s *UpdateManagedAgentRequestBody) GetAccessScope() *string {
	return s.AccessScope
}

func (s *UpdateManagedAgentRequestBody) GetAgenticFsMounts() []*UpdateManagedAgentRequestBodyAgenticFsMounts {
	return s.AgenticFsMounts
}

func (s *UpdateManagedAgentRequestBody) GetContextConfiguration() *UpdateManagedAgentRequestBodyContextConfiguration {
	return s.ContextConfiguration
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

func (s *UpdateManagedAgentRequestBody) GetPromptVariables() []*UpdateManagedAgentRequestBodyPromptVariables {
	return s.PromptVariables
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

func (s *UpdateManagedAgentRequestBody) SetAccessMode(v string) *UpdateManagedAgentRequestBody {
	s.AccessMode = &v
	return s
}

func (s *UpdateManagedAgentRequestBody) SetAccessScope(v string) *UpdateManagedAgentRequestBody {
	s.AccessScope = &v
	return s
}

func (s *UpdateManagedAgentRequestBody) SetAgenticFsMounts(v []*UpdateManagedAgentRequestBodyAgenticFsMounts) *UpdateManagedAgentRequestBody {
	s.AgenticFsMounts = v
	return s
}

func (s *UpdateManagedAgentRequestBody) SetContextConfiguration(v *UpdateManagedAgentRequestBodyContextConfiguration) *UpdateManagedAgentRequestBody {
	s.ContextConfiguration = v
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

func (s *UpdateManagedAgentRequestBody) SetPromptVariables(v []*UpdateManagedAgentRequestBodyPromptVariables) *UpdateManagedAgentRequestBody {
	s.PromptVariables = v
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

type UpdateManagedAgentRequestBodyAgenticFsMounts struct {
	// The absolute mount path within the container, such as /data/train. Mount targets must not be duplicated or form parent-child overrides, and must avoid platform reserved directories.
	//
	// example:
	//
	// /mnt/agenticfs/data
	MountPath *string `json:"mountPath,omitempty" xml:"mountPath,omitempty"`
	// An existing non-empty relative directory under the AccessPoint. Required and validated by the backend for each mount item. Root directories, absolute paths, or parent directory segments are not allowed.
	//
	// example:
	//
	// workspace/data
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// Specifies whether to mount as read-only. The default value is false if omitted. This is not a RAM role read-only policy.
	//
	// example:
	//
	// false
	ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
	// The AccessPoint domain name. Required and validated by the backend for each mount item. It does not contain a protocol, port, or path, and uses the DomainName from the NAS ListAccessPoints response.
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

type UpdateManagedAgentRequestBodyContextConfiguration struct {
	// An optional array. Null or null items are not accepted. Each item only allows type and resourceId. The combination of the same type and normalized resource ID must not be duplicated. Omitting this field during creation indicates no selection. Omitting it during an update retains the original selection, and a non-empty array entirely replaces it. An empty array can only clear the selections when the merged enabled is false. To disable and clear, pass enabled:false and bindings:[] simultaneously. Explicitly submitted non-empty selections validate resource ownership and readiness status, including selections saved when disabled.
	//
	// if can be null:
	// false
	Bindings []*UpdateManagedAgentRequestBodyContextConfigurationBindings `json:"bindings,omitempty" xml:"bindings,omitempty" type:"Repeated"`
	// An optional Boolean value. Null is not accepted. If this field is omitted during creation, the default value is false. If omitted during an update, the current value is retained, and the client must not automatically populate it with false. When updated to true, saved bindings can be reused. The merged enabled configuration must contain at least one resource. When updated to false, it only disables and does not clear the original resource selections.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s UpdateManagedAgentRequestBodyContextConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentRequestBodyContextConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentRequestBodyContextConfiguration) GetBindings() []*UpdateManagedAgentRequestBodyContextConfigurationBindings {
	return s.Bindings
}

func (s *UpdateManagedAgentRequestBodyContextConfiguration) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateManagedAgentRequestBodyContextConfiguration) SetBindings(v []*UpdateManagedAgentRequestBodyContextConfigurationBindings) *UpdateManagedAgentRequestBodyContextConfiguration {
	s.Bindings = v
	return s
}

func (s *UpdateManagedAgentRequestBodyContextConfiguration) SetEnabled(v bool) *UpdateManagedAgentRequestBodyContextConfiguration {
	s.Enabled = &v
	return s
}

func (s *UpdateManagedAgentRequestBodyContextConfiguration) Validate() error {
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

type UpdateManagedAgentRequestBodyContextConfigurationBindings struct {
	// Required and validated by the backend for each binding item. Null, blank, or control characters are not accepted. The backend trims leading and trailing spaces first, and then validates that the length is 1 to 128 characters. The first character must be a letter or digit, and the remaining characters can contain letters, digits, periods, hyphens, or underscores. Only resources in the current workspace can be referenced.
	//
	// example:
	//
	// kb-example
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// Required and validated by the backend for each binding item. Only KNOWLEDGE_BASE (knowledge base) or DATA_CONNECTOR (data connector) is accepted. Case-sensitive. Null or blank values are not accepted.
	//
	// example:
	//
	// KNOWLEDGE_BASE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateManagedAgentRequestBodyContextConfigurationBindings) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentRequestBodyContextConfigurationBindings) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentRequestBodyContextConfigurationBindings) GetResourceId() *string {
	return s.ResourceId
}

func (s *UpdateManagedAgentRequestBodyContextConfigurationBindings) GetType() *string {
	return s.Type
}

func (s *UpdateManagedAgentRequestBodyContextConfigurationBindings) SetResourceId(v string) *UpdateManagedAgentRequestBodyContextConfigurationBindings {
	s.ResourceId = &v
	return s
}

func (s *UpdateManagedAgentRequestBodyContextConfigurationBindings) SetType(v string) *UpdateManagedAgentRequestBodyContextConfigurationBindings {
	s.Type = &v
	return s
}

func (s *UpdateManagedAgentRequestBodyContextConfigurationBindings) Validate() error {
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
	// The environment variable name.
	//
	// This parameter is required.
	//
	// example:
	//
	// API_KEY
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The environment variable value.
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
	// The runtime framework configuration.
	Configuration *UpdateManagedAgentRequestBodyHarnessConfiguration `json:"configuration,omitempty" xml:"configuration,omitempty" type:"Struct"`
	// The runtime framework type. Required when harness is provided, and must be consistent with that specified during creation.
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
	// The account key record ID in the qodercli connector. You can query it through ListConnectors. It is not the original key. This field can be omitted if there is only one record, but must be specified if there are multiple records. Do not specify this field for qwenpaw.
	//
	// example:
	//
	// key-example
	ConnectorServiceAccountKey *string `json:"connectorServiceAccountKey,omitempty" xml:"connectorServiceAccountKey,omitempty"`
	// The connector account name bound to qodercli. Do not specify this field for qwenpaw.
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
	// For qwenpaw, use the workspace model connection ID. For qodercli, use the available model identifier returned by ListConnectorModels.
	//
	// This parameter is required.
	//
	// example:
	//
	// mc-1
	ModelConnectionId *string `json:"modelConnectionId,omitempty" xml:"modelConnectionId,omitempty"`
	// For qwenpaw, use the name of the upstream model created under the model connection. This field is required. For qodercli, this field can be omitted.
	//
	// example:
	//
	// qwen-max
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	// The model token quota configuration. If not passed, no quota is configured.
	Quota *UpdateManagedAgentRequestBodyModelQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
	// Valid values: off, low, medium, high, xhigh, and max. If not passed or set to null during creation, default is used. When updating and submitting the model configuration, if the thinking level is not passed or is null, it reverts to default. Except for off which is converted to none, the value is passed to the runtime framework as is. The system does not validate model support or automatically downgrade the level. If the model does not support the specified level, adjust the level accordingly.
	//
	// example:
	//
	// high
	ThinkingLevel *string `json:"thinkingLevel,omitempty" xml:"thinkingLevel,omitempty"`
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

func (s *UpdateManagedAgentRequestBodyModel) GetThinkingLevel() *string {
	return s.ThinkingLevel
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

func (s *UpdateManagedAgentRequestBodyModel) SetThinkingLevel(v string) *UpdateManagedAgentRequestBodyModel {
	s.ThinkingLevel = &v
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
	// Specifies whether to enable the token quota. The default value is true if not specified. Passing false disables the quota and deletes existing quota rules.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The quota limit type. Required and validated by the backend when the quota is enabled. Fixed value: token.
	//
	// example:
	//
	// token
	LimitType *string `json:"limitType,omitempty" xml:"limitType,omitempty"`
	// The quota statistical period. Required and validated by the backend when the quota is enabled. Valid values: day (daily) and month (monthly).
	//
	// example:
	//
	// day
	PeriodType *string `json:"periodType,omitempty" xml:"periodType,omitempty"`
	// The maximum number of tokens allowed to be consumed in a single period. Required and validated by the backend when the quota is enabled. The value must be greater than 0.
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
	// Specifies whether to allow Internet access.
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
	// The name of the OSS bucket. This field is required and validated by the backend for each mount item.
	//
	// example:
	//
	// bucket-001
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// The absolute mount path within the container, such as /data/train. Mount targets must not be duplicated or form parent-child overrides, and must avoid platform reserved directories.
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

type UpdateManagedAgentRequestBodyPromptVariables struct {
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

func (s UpdateManagedAgentRequestBodyPromptVariables) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentRequestBodyPromptVariables) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentRequestBodyPromptVariables) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *UpdateManagedAgentRequestBodyPromptVariables) GetDescription() *string {
	return s.Description
}

func (s *UpdateManagedAgentRequestBodyPromptVariables) GetName() *string {
	return s.Name
}

func (s *UpdateManagedAgentRequestBodyPromptVariables) SetDefaultValue(v string) *UpdateManagedAgentRequestBodyPromptVariables {
	s.DefaultValue = &v
	return s
}

func (s *UpdateManagedAgentRequestBodyPromptVariables) SetDescription(v string) *UpdateManagedAgentRequestBodyPromptVariables {
	s.Description = &v
	return s
}

func (s *UpdateManagedAgentRequestBodyPromptVariables) SetName(v string) *UpdateManagedAgentRequestBodyPromptVariables {
	s.Name = &v
	return s
}

func (s *UpdateManagedAgentRequestBodyPromptVariables) Validate() error {
	return dara.Validate(s)
}

type UpdateManagedAgentRequestBodyRuntime struct {
	// The compute configuration.
	//
	// This parameter is required.
	Compute *UpdateManagedAgentRequestBodyRuntimeCompute `json:"compute,omitempty" xml:"compute,omitempty" type:"Struct"`
	// The elastic scaling and session configuration. Omitted fields use default values. To retain the original policy during an update, submit the complete hpa object or omit the entire runtime object.
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
	// The compute specification.
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
	// The name of the request header for session identification. This parameter is required when type is set to ISOLATED_HEADER_FIELD and must be a valid HTTP request header name. Do not specify this parameter for other policies.
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
	// The skill source type.
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
	// The skill reference version selector. Supports LABEL for tags and VERSION for specific versions. If omitted, LABEL/latest is used. Selectors are not provided for STATIC sources.
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
	// The version selector type.
	//
	// example:
	//
	// LABEL
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The selector value. When the type is LABEL, this is the tag name, such as latest. When the type is VERSION, this is the specific version number.
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
	// The list of sub-agent skills.
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
	// The skill name.
	//
	// example:
	//
	// web-search
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The skill version.
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
	// The name of the template in the AI Registry.
	//
	// This parameter is required.
	//
	// example:
	//
	// code-review-template
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The version of the template in the AI Registry.
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
