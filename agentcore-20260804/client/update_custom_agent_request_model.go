// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateCustomAgentRequestBody) *UpdateCustomAgentRequest
	GetBody() *UpdateCustomAgentRequestBody
	SetClientToken(v string) *UpdateCustomAgentRequest
	GetClientToken() *string
}

type UpdateCustomAgentRequest struct {
	// The update configuration for the custom agent. Provide at least one top-level field. Omitted top-level fields retain their original values. When providing a configuration object, specify its required subfields and include any optional configurations that you want to retain. Arrays are replaced in their entirety.
	Body *UpdateCustomAgentRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// Idempotency is not currently guaranteed. Do not rely on this parameter to prevent duplicate requests.
	//
	// example:
	//
	// custom-agent-request-20261010-001
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateCustomAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentRequest) GetBody() *UpdateCustomAgentRequestBody {
	return s.Body
}

func (s *UpdateCustomAgentRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateCustomAgentRequest) SetBody(v *UpdateCustomAgentRequestBody) *UpdateCustomAgentRequest {
	s.Body = v
	return s
}

func (s *UpdateCustomAgentRequest) SetClientToken(v string) *UpdateCustomAgentRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCustomAgentRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateCustomAgentRequestBody struct {
	// The list of additional AgenticFS mounts. The total number of AgenticFS and OSS mounts cannot exceed 10. During creation, omitting this field or passing an empty array [] indicates no additional mounts. During an update, omitting this field retains the original value, and passing an empty array [] clears the list. You cannot pass null.
	AgenticFsMounts []*UpdateCustomAgentRequestBodyAgenticFsMounts `json:"agenticFsMounts,omitempty" xml:"agenticFsMounts,omitempty" type:"Repeated"`
	// The deployment artifact configuration. Currently, only ACR container images are supported. Specify the container field and the service listening port field, and omit the code field. During an update, you must specify the port field and all required fields of the container object.
	Artifact *UpdateCustomAgentRequestBodyArtifact `json:"artifact,omitempty" xml:"artifact,omitempty" type:"Struct"`
	// The description of the agent\\"s purpose. The maximum length is 1,024 characters after leading and trailing whitespaces are removed. If omitted, the original value is retained. Pass null to clear the description.
	//
	// example:
	//
	// Custom agent that provides conversation services
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The environment variable configuration for the agent process. Omitting this top-level field retains the original value. Providing this object but omitting the variables field clears the custom environment variables. To retain the variables, provide the complete list.
	Environment *UpdateCustomAgentRequestBodyEnvironment `json:"environment,omitempty" xml:"environment,omitempty" type:"Struct"`
	// The optional HTTP health check configuration. When providing this object, you must specify the port field. Other unspecified items use default values. Omitting this field during creation indicates no configuration. During an update, omitting this top-level field retains the original configuration, and passing null removes the health check.
	HealthCheck *UpdateCustomAgentRequestBodyHealthCheck `json:"healthCheck,omitempty" xml:"healthCheck,omitempty" type:"Struct"`
	// The log delivery configuration for the agent. You can specify the Simple Log Service project and logstore. After a successful creation or update, verify that logs are delivered properly.
	Logging *UpdateCustomAgentRequestBodyLogging `json:"logging,omitempty" xml:"logging,omitempty" type:"Struct"`
	// The name of the custom agent. The length must be 1 to 128 characters after leading and trailing whitespaces are removed. The name must be unique within the same workspace. If omitted, the original value is retained. You cannot pass null or a blank name.
	//
	// example:
	//
	// my-custom-agent
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The network access configuration for the agent. When providing this object, omitted access items are treated as allowed. The actual accessible scope is still subject to the workspace network configuration. To retain the original restrictions during an update, specify both accessInternet and accessVpc.
	Network *UpdateCustomAgentRequestBodyNetwork `json:"network,omitempty" xml:"network,omitempty" type:"Struct"`
	// The list of OSS mounts. The total number of OSS and AgenticFS mounts cannot exceed 10. For each item, you must specify the bucketName, path, and mountPath fields. Mount targets must not be duplicated or overlap with each other. This list is replaced in its entirety during an update. Pass an empty array [] to clear the list.
	OssMounts []*UpdateCustomAgentRequestBodyOssMounts `json:"ossMounts,omitempty" xml:"ossMounts,omitempty" type:"Repeated"`
	// The list of invocation protocols supported by the agent. Each item includes the protocol type, along with optional paths and request headers. The entire list is replaced during an update.
	Protocol []*UpdateCustomAgentRequestBodyProtocol `json:"protocol,omitempty" xml:"protocol,omitempty" type:"Repeated"`
	// The ARN of the RAM role used by the agent to access user cloud resources. The role must have permissions to access the cloud resources.
	//
	// example:
	//
	// acs:ram::1234567890123456:role/agentcore-runtime-role
	RoleArn *string `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	// The runtime configuration. When this object is provided, both compute and sessionPolicy must be present. To retain the original elastic policy, provide the complete hpa configuration. Omitting the top-level runtime field retains the original configuration.
	Runtime *UpdateCustomAgentRequestBodyRuntime `json:"runtime,omitempty" xml:"runtime,omitempty" type:"Struct"`
}

func (s UpdateCustomAgentRequestBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentRequestBody) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentRequestBody) GetAgenticFsMounts() []*UpdateCustomAgentRequestBodyAgenticFsMounts {
	return s.AgenticFsMounts
}

func (s *UpdateCustomAgentRequestBody) GetArtifact() *UpdateCustomAgentRequestBodyArtifact {
	return s.Artifact
}

func (s *UpdateCustomAgentRequestBody) GetDescription() *string {
	return s.Description
}

func (s *UpdateCustomAgentRequestBody) GetEnvironment() *UpdateCustomAgentRequestBodyEnvironment {
	return s.Environment
}

func (s *UpdateCustomAgentRequestBody) GetHealthCheck() *UpdateCustomAgentRequestBodyHealthCheck {
	return s.HealthCheck
}

func (s *UpdateCustomAgentRequestBody) GetLogging() *UpdateCustomAgentRequestBodyLogging {
	return s.Logging
}

func (s *UpdateCustomAgentRequestBody) GetName() *string {
	return s.Name
}

func (s *UpdateCustomAgentRequestBody) GetNetwork() *UpdateCustomAgentRequestBodyNetwork {
	return s.Network
}

func (s *UpdateCustomAgentRequestBody) GetOssMounts() []*UpdateCustomAgentRequestBodyOssMounts {
	return s.OssMounts
}

func (s *UpdateCustomAgentRequestBody) GetProtocol() []*UpdateCustomAgentRequestBodyProtocol {
	return s.Protocol
}

func (s *UpdateCustomAgentRequestBody) GetRoleArn() *string {
	return s.RoleArn
}

func (s *UpdateCustomAgentRequestBody) GetRuntime() *UpdateCustomAgentRequestBodyRuntime {
	return s.Runtime
}

func (s *UpdateCustomAgentRequestBody) SetAgenticFsMounts(v []*UpdateCustomAgentRequestBodyAgenticFsMounts) *UpdateCustomAgentRequestBody {
	s.AgenticFsMounts = v
	return s
}

func (s *UpdateCustomAgentRequestBody) SetArtifact(v *UpdateCustomAgentRequestBodyArtifact) *UpdateCustomAgentRequestBody {
	s.Artifact = v
	return s
}

func (s *UpdateCustomAgentRequestBody) SetDescription(v string) *UpdateCustomAgentRequestBody {
	s.Description = &v
	return s
}

func (s *UpdateCustomAgentRequestBody) SetEnvironment(v *UpdateCustomAgentRequestBodyEnvironment) *UpdateCustomAgentRequestBody {
	s.Environment = v
	return s
}

func (s *UpdateCustomAgentRequestBody) SetHealthCheck(v *UpdateCustomAgentRequestBodyHealthCheck) *UpdateCustomAgentRequestBody {
	s.HealthCheck = v
	return s
}

func (s *UpdateCustomAgentRequestBody) SetLogging(v *UpdateCustomAgentRequestBodyLogging) *UpdateCustomAgentRequestBody {
	s.Logging = v
	return s
}

func (s *UpdateCustomAgentRequestBody) SetName(v string) *UpdateCustomAgentRequestBody {
	s.Name = &v
	return s
}

func (s *UpdateCustomAgentRequestBody) SetNetwork(v *UpdateCustomAgentRequestBodyNetwork) *UpdateCustomAgentRequestBody {
	s.Network = v
	return s
}

func (s *UpdateCustomAgentRequestBody) SetOssMounts(v []*UpdateCustomAgentRequestBodyOssMounts) *UpdateCustomAgentRequestBody {
	s.OssMounts = v
	return s
}

func (s *UpdateCustomAgentRequestBody) SetProtocol(v []*UpdateCustomAgentRequestBodyProtocol) *UpdateCustomAgentRequestBody {
	s.Protocol = v
	return s
}

func (s *UpdateCustomAgentRequestBody) SetRoleArn(v string) *UpdateCustomAgentRequestBody {
	s.RoleArn = &v
	return s
}

func (s *UpdateCustomAgentRequestBody) SetRuntime(v *UpdateCustomAgentRequestBodyRuntime) *UpdateCustomAgentRequestBody {
	s.Runtime = v
	return s
}

func (s *UpdateCustomAgentRequestBody) Validate() error {
	if s.AgenticFsMounts != nil {
		for _, item := range s.AgenticFsMounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Artifact != nil {
		if err := s.Artifact.Validate(); err != nil {
			return err
		}
	}
	if s.Environment != nil {
		if err := s.Environment.Validate(); err != nil {
			return err
		}
	}
	if s.HealthCheck != nil {
		if err := s.HealthCheck.Validate(); err != nil {
			return err
		}
	}
	if s.Logging != nil {
		if err := s.Logging.Validate(); err != nil {
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
	if s.Protocol != nil {
		for _, item := range s.Protocol {
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
	return nil
}

type UpdateCustomAgentRequestBodyAgenticFsMounts struct {
	// The subdirectory under /mnt/agenticfs/ in the container. This field is required for each mount item. Mount targets must not be duplicated or form parent-child overrides.
	//
	// example:
	//
	// /mnt/agenticfs/data
	MountPath *string `json:"mountPath,omitempty" xml:"mountPath,omitempty"`
	// The existing non-empty relative directory under the access point. This field is required for each mount item. Root directories, absolute paths, and parent directory traversals are not allowed.
	//
	// example:
	//
	// workspace/data
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// Specifies whether the mount is read-only. A value of true indicates read-only, and false indicates read and write. If omitted, the default value is false.
	//
	// example:
	//
	// false
	ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
	// The domain name of the access point. This field is required for each mount item. Use the DomainName value from the NAS ListAccessPoints response. Do not include the protocol, port, or path.
	//
	// example:
	//
	// ap-0123456789abcdef0.0123456789-vlm36.cn-hangzhou.nas.aliyuncs.com
	Server *string `json:"server,omitempty" xml:"server,omitempty"`
}

func (s UpdateCustomAgentRequestBodyAgenticFsMounts) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentRequestBodyAgenticFsMounts) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentRequestBodyAgenticFsMounts) GetMountPath() *string {
	return s.MountPath
}

func (s *UpdateCustomAgentRequestBodyAgenticFsMounts) GetPath() *string {
	return s.Path
}

func (s *UpdateCustomAgentRequestBodyAgenticFsMounts) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *UpdateCustomAgentRequestBodyAgenticFsMounts) GetServer() *string {
	return s.Server
}

func (s *UpdateCustomAgentRequestBodyAgenticFsMounts) SetMountPath(v string) *UpdateCustomAgentRequestBodyAgenticFsMounts {
	s.MountPath = &v
	return s
}

func (s *UpdateCustomAgentRequestBodyAgenticFsMounts) SetPath(v string) *UpdateCustomAgentRequestBodyAgenticFsMounts {
	s.Path = &v
	return s
}

func (s *UpdateCustomAgentRequestBodyAgenticFsMounts) SetReadOnly(v bool) *UpdateCustomAgentRequestBodyAgenticFsMounts {
	s.ReadOnly = &v
	return s
}

func (s *UpdateCustomAgentRequestBodyAgenticFsMounts) SetServer(v string) *UpdateCustomAgentRequestBodyAgenticFsMounts {
	s.Server = &v
	return s
}

func (s *UpdateCustomAgentRequestBodyAgenticFsMounts) Validate() error {
	return dara.Validate(s)
}

type UpdateCustomAgentRequestBodyArtifact struct {
	// The code package deployment configuration. This is currently not supported. Omit this field and use artifact.container instead.
	Code *UpdateCustomAgentRequestBodyArtifactCode `json:"code,omitempty" xml:"code,omitempty" type:"Struct"`
	// The ACR container image configuration. You must specify the type, instanceId, namespace, repo, and tag fields. Do not configure the code field at the same time.
	Container *UpdateCustomAgentRequestBodyArtifactContainer `json:"container,omitempty" xml:"container,omitempty" type:"Struct"`
	// The TCP port listened to by the agent service. Valid values: 1 to 65535.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8080
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
}

func (s UpdateCustomAgentRequestBodyArtifact) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentRequestBodyArtifact) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentRequestBodyArtifact) GetCode() *UpdateCustomAgentRequestBodyArtifactCode {
	return s.Code
}

func (s *UpdateCustomAgentRequestBodyArtifact) GetContainer() *UpdateCustomAgentRequestBodyArtifactContainer {
	return s.Container
}

func (s *UpdateCustomAgentRequestBodyArtifact) GetPort() *int32 {
	return s.Port
}

func (s *UpdateCustomAgentRequestBodyArtifact) SetCode(v *UpdateCustomAgentRequestBodyArtifactCode) *UpdateCustomAgentRequestBodyArtifact {
	s.Code = v
	return s
}

func (s *UpdateCustomAgentRequestBodyArtifact) SetContainer(v *UpdateCustomAgentRequestBodyArtifactContainer) *UpdateCustomAgentRequestBodyArtifact {
	s.Container = v
	return s
}

func (s *UpdateCustomAgentRequestBodyArtifact) SetPort(v int32) *UpdateCustomAgentRequestBodyArtifact {
	s.Port = &v
	return s
}

func (s *UpdateCustomAgentRequestBodyArtifact) Validate() error {
	if s.Code != nil {
		if err := s.Code.Validate(); err != nil {
			return err
		}
	}
	if s.Container != nil {
		if err := s.Container.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateCustomAgentRequestBodyArtifactCode struct {
	// The startup entry point for the code package. Code package deployment is currently not supported. Omit this field.
	//
	// This parameter is required.
	//
	// example:
	//
	// python app.py
	Entrypoint *string `json:"entrypoint,omitempty" xml:"entrypoint,omitempty"`
	// The language runtime identifier for the code package. Code package deployment is currently not supported. Omit this field.
	//
	// This parameter is required.
	//
	// example:
	//
	// python
	LanguageRuntime *string `json:"languageRuntime,omitempty" xml:"languageRuntime,omitempty"`
	// The source URI of the code package. Code package deployment is currently not supported. Omit this field.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://my-agent-bucket/packages/agent.zip
	SourceUri *string `json:"sourceUri,omitempty" xml:"sourceUri,omitempty"`
}

func (s UpdateCustomAgentRequestBodyArtifactCode) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentRequestBodyArtifactCode) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentRequestBodyArtifactCode) GetEntrypoint() *string {
	return s.Entrypoint
}

func (s *UpdateCustomAgentRequestBodyArtifactCode) GetLanguageRuntime() *string {
	return s.LanguageRuntime
}

func (s *UpdateCustomAgentRequestBodyArtifactCode) GetSourceUri() *string {
	return s.SourceUri
}

func (s *UpdateCustomAgentRequestBodyArtifactCode) SetEntrypoint(v string) *UpdateCustomAgentRequestBodyArtifactCode {
	s.Entrypoint = &v
	return s
}

func (s *UpdateCustomAgentRequestBodyArtifactCode) SetLanguageRuntime(v string) *UpdateCustomAgentRequestBodyArtifactCode {
	s.LanguageRuntime = &v
	return s
}

func (s *UpdateCustomAgentRequestBodyArtifactCode) SetSourceUri(v string) *UpdateCustomAgentRequestBodyArtifactCode {
	s.SourceUri = &v
	return s
}

func (s *UpdateCustomAgentRequestBodyArtifactCode) Validate() error {
	return dara.Validate(s)
}

type UpdateCustomAgentRequestBodyArtifactContainer struct {
	// The optional executable program and startup parameters. If specified, this must be a non-empty string. When updating the artifact, omitting this field uses the default startup method of the image. To retain the original custom startup command, provide this field together.
	//
	// example:
	//
	// python app.py
	Entrypoint *string `json:"entrypoint,omitempty" xml:"entrypoint,omitempty"`
	// The ID of the ACR instance to which the image repository belongs. This field cannot be empty.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-a1b2c3d4e5f6
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The namespace to which the ACR image repository belongs. This field cannot be empty.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-agent
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The name of the image repository in the ACR namespace. This field cannot be empty.
	//
	// This parameter is required.
	//
	// example:
	//
	// chat-agent
	Repo *string `json:"repo,omitempty" xml:"repo,omitempty"`
	// The image tag used for deployment. This field cannot be empty. When updating the tag, you must also provide artifact.port and other required image fields.
	//
	// This parameter is required.
	//
	// example:
	//
	// v1.0.0
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// The image source type. Currently, only acr is supported, which indicates Alibaba Cloud Container Registry (ACR).
	//
	// This parameter is required.
	//
	// example:
	//
	// acr
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateCustomAgentRequestBodyArtifactContainer) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentRequestBodyArtifactContainer) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentRequestBodyArtifactContainer) GetEntrypoint() *string {
	return s.Entrypoint
}

func (s *UpdateCustomAgentRequestBodyArtifactContainer) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateCustomAgentRequestBodyArtifactContainer) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateCustomAgentRequestBodyArtifactContainer) GetRepo() *string {
	return s.Repo
}

func (s *UpdateCustomAgentRequestBodyArtifactContainer) GetTag() *string {
	return s.Tag
}

func (s *UpdateCustomAgentRequestBodyArtifactContainer) GetType() *string {
	return s.Type
}

func (s *UpdateCustomAgentRequestBodyArtifactContainer) SetEntrypoint(v string) *UpdateCustomAgentRequestBodyArtifactContainer {
	s.Entrypoint = &v
	return s
}

func (s *UpdateCustomAgentRequestBodyArtifactContainer) SetInstanceId(v string) *UpdateCustomAgentRequestBodyArtifactContainer {
	s.InstanceId = &v
	return s
}

func (s *UpdateCustomAgentRequestBodyArtifactContainer) SetNamespace(v string) *UpdateCustomAgentRequestBodyArtifactContainer {
	s.Namespace = &v
	return s
}

func (s *UpdateCustomAgentRequestBodyArtifactContainer) SetRepo(v string) *UpdateCustomAgentRequestBodyArtifactContainer {
	s.Repo = &v
	return s
}

func (s *UpdateCustomAgentRequestBodyArtifactContainer) SetTag(v string) *UpdateCustomAgentRequestBodyArtifactContainer {
	s.Tag = &v
	return s
}

func (s *UpdateCustomAgentRequestBodyArtifactContainer) SetType(v string) *UpdateCustomAgentRequestBodyArtifactContainer {
	s.Type = &v
	return s
}

func (s *UpdateCustomAgentRequestBodyArtifactContainer) Validate() error {
	return dara.Validate(s)
}

type UpdateCustomAgentRequestBodyEnvironment struct {
	// The list of environment variables. Each item includes a name and a value. This list is replaced in its entirety during an update.
	Variables []*UpdateCustomAgentRequestBodyEnvironmentVariables `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
}

func (s UpdateCustomAgentRequestBodyEnvironment) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentRequestBodyEnvironment) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentRequestBodyEnvironment) GetVariables() []*UpdateCustomAgentRequestBodyEnvironmentVariables {
	return s.Variables
}

func (s *UpdateCustomAgentRequestBodyEnvironment) SetVariables(v []*UpdateCustomAgentRequestBodyEnvironmentVariables) *UpdateCustomAgentRequestBodyEnvironment {
	s.Variables = v
	return s
}

func (s *UpdateCustomAgentRequestBodyEnvironment) Validate() error {
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

type UpdateCustomAgentRequestBodyEnvironmentVariables struct {
	// The name of the environment variable. This field cannot be empty and must be unique within the current environment variable list.
	//
	// This parameter is required.
	//
	// example:
	//
	// LOG_LEVEL
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The value of the environment variable. This must be a string. Empty strings are allowed.
	//
	// This parameter is required.
	//
	// example:
	//
	// info
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UpdateCustomAgentRequestBodyEnvironmentVariables) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentRequestBodyEnvironmentVariables) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentRequestBodyEnvironmentVariables) GetName() *string {
	return s.Name
}

func (s *UpdateCustomAgentRequestBodyEnvironmentVariables) GetValue() *string {
	return s.Value
}

func (s *UpdateCustomAgentRequestBodyEnvironmentVariables) SetName(v string) *UpdateCustomAgentRequestBodyEnvironmentVariables {
	s.Name = &v
	return s
}

func (s *UpdateCustomAgentRequestBodyEnvironmentVariables) SetValue(v string) *UpdateCustomAgentRequestBodyEnvironmentVariables {
	s.Value = &v
	return s
}

func (s *UpdateCustomAgentRequestBodyEnvironmentVariables) Validate() error {
	return dara.Validate(s)
}

type UpdateCustomAgentRequestBodyHealthCheck struct {
	// This parameter currently does not take effect. Omit this field.
	//
	// example:
	//
	// 1
	HealthyThreshold *int32 `json:"healthyThreshold,omitempty" xml:"healthyThreshold,omitempty"`
	// The interval between two health checks, in seconds. This must be a positive integer. If omitted, the default value is 10.
	//
	// example:
	//
	// 10
	IntervalSeconds *int32 `json:"intervalSeconds,omitempty" xml:"intervalSeconds,omitempty"`
	// The HTTP health check path. This must start with a forward slash (/). If omitted, the default value is /healthz.
	//
	// example:
	//
	// /healthz
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// The port accessed for health checks. This field is required when configuring a health check. Valid values: 1 to 65535.
	//
	// example:
	//
	// 8080
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// The timeout period for a single health check, in seconds. This must be a positive integer and less than intervalSeconds. If omitted, the default value is 2.
	//
	// example:
	//
	// 2
	TimeoutSeconds *int32 `json:"timeoutSeconds,omitempty" xml:"timeoutSeconds,omitempty"`
	// The number of consecutive health check failures required to mark the agent as unhealthy. This must be a positive integer. If omitted, the default value is 3.
	//
	// example:
	//
	// 3
	UnhealthyThreshold *int32 `json:"unhealthyThreshold,omitempty" xml:"unhealthyThreshold,omitempty"`
}

func (s UpdateCustomAgentRequestBodyHealthCheck) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentRequestBodyHealthCheck) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentRequestBodyHealthCheck) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *UpdateCustomAgentRequestBodyHealthCheck) GetIntervalSeconds() *int32 {
	return s.IntervalSeconds
}

func (s *UpdateCustomAgentRequestBodyHealthCheck) GetPath() *string {
	return s.Path
}

func (s *UpdateCustomAgentRequestBodyHealthCheck) GetPort() *int32 {
	return s.Port
}

func (s *UpdateCustomAgentRequestBodyHealthCheck) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *UpdateCustomAgentRequestBodyHealthCheck) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *UpdateCustomAgentRequestBodyHealthCheck) SetHealthyThreshold(v int32) *UpdateCustomAgentRequestBodyHealthCheck {
	s.HealthyThreshold = &v
	return s
}

func (s *UpdateCustomAgentRequestBodyHealthCheck) SetIntervalSeconds(v int32) *UpdateCustomAgentRequestBodyHealthCheck {
	s.IntervalSeconds = &v
	return s
}

func (s *UpdateCustomAgentRequestBodyHealthCheck) SetPath(v string) *UpdateCustomAgentRequestBodyHealthCheck {
	s.Path = &v
	return s
}

func (s *UpdateCustomAgentRequestBodyHealthCheck) SetPort(v int32) *UpdateCustomAgentRequestBodyHealthCheck {
	s.Port = &v
	return s
}

func (s *UpdateCustomAgentRequestBodyHealthCheck) SetTimeoutSeconds(v int32) *UpdateCustomAgentRequestBodyHealthCheck {
	s.TimeoutSeconds = &v
	return s
}

func (s *UpdateCustomAgentRequestBodyHealthCheck) SetUnhealthyThreshold(v int32) *UpdateCustomAgentRequestBodyHealthCheck {
	s.UnhealthyThreshold = &v
	return s
}

func (s *UpdateCustomAgentRequestBodyHealthCheck) Validate() error {
	return dara.Validate(s)
}

type UpdateCustomAgentRequestBodyLogging struct {
	// The Simple Log Service log delivery destination. When updating this object, you must provide both the project and logstore fields, and they cannot be empty.
	Sls *UpdateCustomAgentRequestBodyLoggingSls `json:"sls,omitempty" xml:"sls,omitempty" type:"Struct"`
}

func (s UpdateCustomAgentRequestBodyLogging) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentRequestBodyLogging) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentRequestBodyLogging) GetSls() *UpdateCustomAgentRequestBodyLoggingSls {
	return s.Sls
}

func (s *UpdateCustomAgentRequestBodyLogging) SetSls(v *UpdateCustomAgentRequestBodyLoggingSls) *UpdateCustomAgentRequestBodyLogging {
	s.Sls = v
	return s
}

func (s *UpdateCustomAgentRequestBodyLogging) Validate() error {
	if s.Sls != nil {
		if err := s.Sls.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateCustomAgentRequestBodyLoggingSls struct {
	// The name of the Simple Log Service logstore that receives agent logs.
	//
	// example:
	//
	// agent-runtime
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// The name of the Simple Log Service project that receives agent logs.
	//
	// example:
	//
	// my-agent-logs
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
}

func (s UpdateCustomAgentRequestBodyLoggingSls) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentRequestBodyLoggingSls) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentRequestBodyLoggingSls) GetLogstore() *string {
	return s.Logstore
}

func (s *UpdateCustomAgentRequestBodyLoggingSls) GetProject() *string {
	return s.Project
}

func (s *UpdateCustomAgentRequestBodyLoggingSls) SetLogstore(v string) *UpdateCustomAgentRequestBodyLoggingSls {
	s.Logstore = &v
	return s
}

func (s *UpdateCustomAgentRequestBodyLoggingSls) SetProject(v string) *UpdateCustomAgentRequestBodyLoggingSls {
	s.Project = &v
	return s
}

func (s *UpdateCustomAgentRequestBodyLoggingSls) Validate() error {
	return dara.Validate(s)
}

type UpdateCustomAgentRequestBodyNetwork struct {
	// The Internet access configuration. If this access item is omitted, Internet access is not restricted on the agent side. However, the workspace must still allow Internet access.
	AccessInternet *UpdateCustomAgentRequestBodyNetworkAccessInternet `json:"accessInternet,omitempty" xml:"accessInternet,omitempty" type:"Struct"`
	// The user VPC access configuration. If this access item is omitted, VPC access is not restricted on the agent side. However, the workspace must still be configured to allow access to the user VPC.
	AccessVpc *UpdateCustomAgentRequestBodyNetworkAccessVpc `json:"accessVpc,omitempty" xml:"accessVpc,omitempty" type:"Struct"`
}

func (s UpdateCustomAgentRequestBodyNetwork) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentRequestBodyNetwork) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentRequestBodyNetwork) GetAccessInternet() *UpdateCustomAgentRequestBodyNetworkAccessInternet {
	return s.AccessInternet
}

func (s *UpdateCustomAgentRequestBodyNetwork) GetAccessVpc() *UpdateCustomAgentRequestBodyNetworkAccessVpc {
	return s.AccessVpc
}

func (s *UpdateCustomAgentRequestBodyNetwork) SetAccessInternet(v *UpdateCustomAgentRequestBodyNetworkAccessInternet) *UpdateCustomAgentRequestBodyNetwork {
	s.AccessInternet = v
	return s
}

func (s *UpdateCustomAgentRequestBodyNetwork) SetAccessVpc(v *UpdateCustomAgentRequestBodyNetworkAccessVpc) *UpdateCustomAgentRequestBodyNetwork {
	s.AccessVpc = v
	return s
}

func (s *UpdateCustomAgentRequestBodyNetwork) Validate() error {
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

type UpdateCustomAgentRequestBodyNetworkAccessInternet struct {
	// Specifies whether to allow the agent to access the Internet. A Boolean value is required when configuring accessInternet. For this to take effect, the workspace must also allow Internet access.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s UpdateCustomAgentRequestBodyNetworkAccessInternet) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentRequestBodyNetworkAccessInternet) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentRequestBodyNetworkAccessInternet) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateCustomAgentRequestBodyNetworkAccessInternet) SetEnabled(v bool) *UpdateCustomAgentRequestBodyNetworkAccessInternet {
	s.Enabled = &v
	return s
}

func (s *UpdateCustomAgentRequestBodyNetworkAccessInternet) Validate() error {
	return dara.Validate(s)
}

type UpdateCustomAgentRequestBodyNetworkAccessVpc struct {
	// Specifies whether to allow the agent to access the user VPC configured for the workspace. A Boolean value is required when configuring accessVpc. For this to take effect, the workspace must also allow VPC access.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s UpdateCustomAgentRequestBodyNetworkAccessVpc) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentRequestBodyNetworkAccessVpc) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentRequestBodyNetworkAccessVpc) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateCustomAgentRequestBodyNetworkAccessVpc) SetEnabled(v bool) *UpdateCustomAgentRequestBodyNetworkAccessVpc {
	s.Enabled = &v
	return s
}

func (s *UpdateCustomAgentRequestBodyNetworkAccessVpc) Validate() error {
	return dara.Validate(s)
}

type UpdateCustomAgentRequestBodyOssMounts struct {
	// The name of the OSS bucket. This field is required for each mount item. The name must be 3 to 63 characters in length and can contain only lowercase letters, digits, and hyphens. It cannot start or end with a hyphen.
	//
	// example:
	//
	// my-agent-data
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// The subdirectory under /mnt/oss/ in the container. This field is required for each mount item. Mount targets must not be duplicated or form parent-child overrides.
	//
	// example:
	//
	// /mnt/oss/train
	MountPath *string `json:"mountPath,omitempty" xml:"mountPath,omitempty"`
	// The non-empty object directory prefix in the bucket, such as datasets/train. Leading and trailing forward slashes (/) are removed. Root directories are not supported. Do not use parent directory traversals.
	//
	// example:
	//
	// datasets/train
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// Specifies whether the mount is read-only. A value of true indicates read-only, and false indicates read and write. If omitted, the default value is false.
	ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
}

func (s UpdateCustomAgentRequestBodyOssMounts) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentRequestBodyOssMounts) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentRequestBodyOssMounts) GetBucketName() *string {
	return s.BucketName
}

func (s *UpdateCustomAgentRequestBodyOssMounts) GetMountPath() *string {
	return s.MountPath
}

func (s *UpdateCustomAgentRequestBodyOssMounts) GetPath() *string {
	return s.Path
}

func (s *UpdateCustomAgentRequestBodyOssMounts) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *UpdateCustomAgentRequestBodyOssMounts) SetBucketName(v string) *UpdateCustomAgentRequestBodyOssMounts {
	s.BucketName = &v
	return s
}

func (s *UpdateCustomAgentRequestBodyOssMounts) SetMountPath(v string) *UpdateCustomAgentRequestBodyOssMounts {
	s.MountPath = &v
	return s
}

func (s *UpdateCustomAgentRequestBodyOssMounts) SetPath(v string) *UpdateCustomAgentRequestBodyOssMounts {
	s.Path = &v
	return s
}

func (s *UpdateCustomAgentRequestBodyOssMounts) SetReadOnly(v bool) *UpdateCustomAgentRequestBodyOssMounts {
	s.ReadOnly = &v
	return s
}

func (s *UpdateCustomAgentRequestBodyOssMounts) Validate() error {
	return dara.Validate(s)
}

type UpdateCustomAgentRequestBodyProtocol struct {
	// The request header configuration for the current protocol. The key is a non-empty request header name, and the value must be a string.
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	// The list of API paths corresponding to the current protocol.
	Paths []*UpdateCustomAgentRequestBodyProtocolPaths `json:"paths,omitempty" xml:"paths,omitempty" type:"Repeated"`
	// The invocation protocol type. Valid values:
	//
	// - agui: AG-UI protocol.
	//
	// - openai-chat-completion: OpenAI Chat Completions compatible protocol.
	//
	// - custom: Custom protocol.
	//
	// This parameter is required.
	//
	// example:
	//
	// agui
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateCustomAgentRequestBodyProtocol) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentRequestBodyProtocol) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentRequestBodyProtocol) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCustomAgentRequestBodyProtocol) GetPaths() []*UpdateCustomAgentRequestBodyProtocolPaths {
	return s.Paths
}

func (s *UpdateCustomAgentRequestBodyProtocol) GetType() *string {
	return s.Type
}

func (s *UpdateCustomAgentRequestBodyProtocol) SetHeaders(v map[string]*string) *UpdateCustomAgentRequestBodyProtocol {
	s.Headers = v
	return s
}

func (s *UpdateCustomAgentRequestBodyProtocol) SetPaths(v []*UpdateCustomAgentRequestBodyProtocolPaths) *UpdateCustomAgentRequestBodyProtocol {
	s.Paths = v
	return s
}

func (s *UpdateCustomAgentRequestBodyProtocol) SetType(v string) *UpdateCustomAgentRequestBodyProtocol {
	s.Type = &v
	return s
}

func (s *UpdateCustomAgentRequestBodyProtocol) Validate() error {
	if s.Paths != nil {
		for _, item := range s.Paths {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateCustomAgentRequestBodyProtocolPaths struct {
	// The description of the current API path.
	//
	// example:
	//
	// AG-UI conversation API
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The non-empty API path corresponding to the current protocol.
	//
	// This parameter is required.
	//
	// example:
	//
	// /agui
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s UpdateCustomAgentRequestBodyProtocolPaths) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentRequestBodyProtocolPaths) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentRequestBodyProtocolPaths) GetDescription() *string {
	return s.Description
}

func (s *UpdateCustomAgentRequestBodyProtocolPaths) GetName() *string {
	return s.Name
}

func (s *UpdateCustomAgentRequestBodyProtocolPaths) SetDescription(v string) *UpdateCustomAgentRequestBodyProtocolPaths {
	s.Description = &v
	return s
}

func (s *UpdateCustomAgentRequestBodyProtocolPaths) SetName(v string) *UpdateCustomAgentRequestBodyProtocolPaths {
	s.Name = &v
	return s
}

func (s *UpdateCustomAgentRequestBodyProtocolPaths) Validate() error {
	return dara.Validate(s)
}

type UpdateCustomAgentRequestBodyRuntime struct {
	// The compute resource configuration for the agent runtime instance.
	//
	// This parameter is required.
	Compute *UpdateCustomAgentRequestBodyRuntimeCompute `json:"compute,omitempty" xml:"compute,omitempty" type:"Struct"`
	// The optional elastic policy. When runtime is provided, omitting hpa or its fields uses the default values: enabled is false, the minimum number of instances is 1, the maximum number of instances is 10, the maximum number of concurrent sessions per instance is 100, and the session idle timeout is 600 seconds. To retain the original policy during an update, provide the complete hpa configuration.
	Hpa *UpdateCustomAgentRequestBodyRuntimeHpa `json:"hpa,omitempty" xml:"hpa,omitempty" type:"Struct"`
	// The session policy configuration, which is used to set the session association method for requests.
	//
	// This parameter is required.
	SessionPolicy *UpdateCustomAgentRequestBodyRuntimeSessionPolicy `json:"sessionPolicy,omitempty" xml:"sessionPolicy,omitempty" type:"Struct"`
}

func (s UpdateCustomAgentRequestBodyRuntime) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentRequestBodyRuntime) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentRequestBodyRuntime) GetCompute() *UpdateCustomAgentRequestBodyRuntimeCompute {
	return s.Compute
}

func (s *UpdateCustomAgentRequestBodyRuntime) GetHpa() *UpdateCustomAgentRequestBodyRuntimeHpa {
	return s.Hpa
}

func (s *UpdateCustomAgentRequestBodyRuntime) GetSessionPolicy() *UpdateCustomAgentRequestBodyRuntimeSessionPolicy {
	return s.SessionPolicy
}

func (s *UpdateCustomAgentRequestBodyRuntime) SetCompute(v *UpdateCustomAgentRequestBodyRuntimeCompute) *UpdateCustomAgentRequestBodyRuntime {
	s.Compute = v
	return s
}

func (s *UpdateCustomAgentRequestBodyRuntime) SetHpa(v *UpdateCustomAgentRequestBodyRuntimeHpa) *UpdateCustomAgentRequestBodyRuntime {
	s.Hpa = v
	return s
}

func (s *UpdateCustomAgentRequestBodyRuntime) SetSessionPolicy(v *UpdateCustomAgentRequestBodyRuntimeSessionPolicy) *UpdateCustomAgentRequestBodyRuntime {
	s.SessionPolicy = v
	return s
}

func (s *UpdateCustomAgentRequestBodyRuntime) Validate() error {
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

type UpdateCustomAgentRequestBodyRuntimeCompute struct {
	// The compute specification. Valid values:
	//
	// - INTERACTIVE: 1 vCPU and 2 GiB of memory.
	//
	// - STANDARD: 2 vCPUs and 4 GiB of memory.
	//
	// - HEAVY: 4 vCPUs and 8 GiB of memory.
	//
	// This parameter is required.
	//
	// example:
	//
	// STANDARD
	ComputeClass *string `json:"computeClass,omitempty" xml:"computeClass,omitempty"`
}

func (s UpdateCustomAgentRequestBodyRuntimeCompute) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentRequestBodyRuntimeCompute) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentRequestBodyRuntimeCompute) GetComputeClass() *string {
	return s.ComputeClass
}

func (s *UpdateCustomAgentRequestBodyRuntimeCompute) SetComputeClass(v string) *UpdateCustomAgentRequestBodyRuntimeCompute {
	s.ComputeClass = &v
	return s
}

func (s *UpdateCustomAgentRequestBodyRuntimeCompute) Validate() error {
	return dara.Validate(s)
}

type UpdateCustomAgentRequestBodyRuntimeHpa struct {
	// Specifies whether to enable elastic scaling. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// When runtime is provided, omitting this field defaults to false.
	//
	// This parameter is required.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The maximum number of concurrent sessions for a single runtime instance. Valid values: 1 to 100. When runtime is provided, omitting this field uses the default value of 100.
	//
	// example:
	//
	// 100
	MaxConcurrentSessionsPerSandbox *int32 `json:"maxConcurrentSessionsPerSandbox,omitempty" xml:"maxConcurrentSessionsPerSandbox,omitempty"`
	// The maximum number of runtime instances. Valid values: 1 to 100. The value cannot be less than minSandboxCount. When runtime is provided, omitting this field uses the default value of 10.
	//
	// example:
	//
	// 10
	MaxSandboxCount *int32 `json:"maxSandboxCount,omitempty" xml:"maxSandboxCount,omitempty"`
	// The minimum number of runtime instances. Valid values: 1 to 100. The value cannot be greater than maxSandboxCount. When runtime is provided, omitting this field uses the default value of 1.
	//
	// example:
	//
	// 1
	MinSandboxCount *int32 `json:"minSandboxCount,omitempty" xml:"minSandboxCount,omitempty"`
	// The session idle timeout in seconds. Valid values: 60 to 86400. When runtime is provided, omitting this field uses the default value of 600.
	//
	// example:
	//
	// 600
	SessionTtlSeconds *int32 `json:"sessionTtlSeconds,omitempty" xml:"sessionTtlSeconds,omitempty"`
}

func (s UpdateCustomAgentRequestBodyRuntimeHpa) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentRequestBodyRuntimeHpa) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentRequestBodyRuntimeHpa) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateCustomAgentRequestBodyRuntimeHpa) GetMaxConcurrentSessionsPerSandbox() *int32 {
	return s.MaxConcurrentSessionsPerSandbox
}

func (s *UpdateCustomAgentRequestBodyRuntimeHpa) GetMaxSandboxCount() *int32 {
	return s.MaxSandboxCount
}

func (s *UpdateCustomAgentRequestBodyRuntimeHpa) GetMinSandboxCount() *int32 {
	return s.MinSandboxCount
}

func (s *UpdateCustomAgentRequestBodyRuntimeHpa) GetSessionTtlSeconds() *int32 {
	return s.SessionTtlSeconds
}

func (s *UpdateCustomAgentRequestBodyRuntimeHpa) SetEnabled(v bool) *UpdateCustomAgentRequestBodyRuntimeHpa {
	s.Enabled = &v
	return s
}

func (s *UpdateCustomAgentRequestBodyRuntimeHpa) SetMaxConcurrentSessionsPerSandbox(v int32) *UpdateCustomAgentRequestBodyRuntimeHpa {
	s.MaxConcurrentSessionsPerSandbox = &v
	return s
}

func (s *UpdateCustomAgentRequestBodyRuntimeHpa) SetMaxSandboxCount(v int32) *UpdateCustomAgentRequestBodyRuntimeHpa {
	s.MaxSandboxCount = &v
	return s
}

func (s *UpdateCustomAgentRequestBodyRuntimeHpa) SetMinSandboxCount(v int32) *UpdateCustomAgentRequestBodyRuntimeHpa {
	s.MinSandboxCount = &v
	return s
}

func (s *UpdateCustomAgentRequestBodyRuntimeHpa) SetSessionTtlSeconds(v int32) *UpdateCustomAgentRequestBodyRuntimeHpa {
	s.SessionTtlSeconds = &v
	return s
}

func (s *UpdateCustomAgentRequestBodyRuntimeHpa) Validate() error {
	return dara.Validate(s)
}

type UpdateCustomAgentRequestBodyRuntimeSessionPolicy struct {
	// The name of the HTTP request header used by the session policy. This parameter is required when type is ISOLATED_HEADER_FIELD and must be a valid HTTP request header name. Omit this parameter for other policies.
	//
	// example:
	//
	// X-Session-Id
	HeaderName *string `json:"headerName,omitempty" xml:"headerName,omitempty"`
	// The session policy type. Valid values:
	//
	// - AFFINITY: Session affinity.
	//
	// - ISOLATED_HEADER_FIELD: Associates sessions by using the request header specified by headerName.
	//
	// - ISOLATED_COOKIE: Not recommended.
	//
	// - DISABLED: A compatibility option. The response returns AFFINITY.
	//
	// This parameter is required.
	//
	// example:
	//
	// AFFINITY
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateCustomAgentRequestBodyRuntimeSessionPolicy) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentRequestBodyRuntimeSessionPolicy) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentRequestBodyRuntimeSessionPolicy) GetHeaderName() *string {
	return s.HeaderName
}

func (s *UpdateCustomAgentRequestBodyRuntimeSessionPolicy) GetType() *string {
	return s.Type
}

func (s *UpdateCustomAgentRequestBodyRuntimeSessionPolicy) SetHeaderName(v string) *UpdateCustomAgentRequestBodyRuntimeSessionPolicy {
	s.HeaderName = &v
	return s
}

func (s *UpdateCustomAgentRequestBodyRuntimeSessionPolicy) SetType(v string) *UpdateCustomAgentRequestBodyRuntimeSessionPolicy {
	s.Type = &v
	return s
}

func (s *UpdateCustomAgentRequestBodyRuntimeSessionPolicy) Validate() error {
	return dara.Validate(s)
}
