// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateCustomAgentRequestBody) *CreateCustomAgentRequest
	GetBody() *CreateCustomAgentRequestBody
	SetClientToken(v string) *CreateCustomAgentRequest
	GetClientToken() *string
}

type CreateCustomAgentRequest struct {
	// The JSON configuration of the custom agent. You must specify name, artifact, roleArn, and runtime.
	Body *CreateCustomAgentRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// Idempotency is not currently guaranteed. Do not rely on this parameter to prevent duplicate requests.
	//
	// example:
	//
	// custom-agent-request-20261010-001
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateCustomAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentRequest) GetBody() *CreateCustomAgentRequestBody {
	return s.Body
}

func (s *CreateCustomAgentRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateCustomAgentRequest) SetBody(v *CreateCustomAgentRequestBody) *CreateCustomAgentRequest {
	s.Body = v
	return s
}

func (s *CreateCustomAgentRequest) SetClientToken(v string) *CreateCustomAgentRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCustomAgentRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCustomAgentRequestBody struct {
	// The list of additional AgenticFS mounts. The total number of AgenticFS and OSS mounts cannot exceed 10. If you omit this parameter or pass an empty array [] during creation, no additional mounts are configured. If you omit this parameter during an update, the original value is retained. If you pass an empty array [], the mounts are cleared. You cannot pass null.
	AgenticFsMounts []*CreateCustomAgentRequestBodyAgenticFsMounts `json:"agenticFsMounts,omitempty" xml:"agenticFsMounts,omitempty" type:"Repeated"`
	// The deployment artifact configuration. Currently, only ACR container images are supported. Specify container and the service listening port port. Omit code.
	//
	// This parameter is required.
	Artifact *CreateCustomAgentRequestBodyArtifact `json:"artifact,omitempty" xml:"artifact,omitempty" type:"Struct"`
	// The description of the agent\\"s purpose. The maximum length is 1,024 characters after leading and trailing whitespace is removed.
	//
	// example:
	//
	// Custom agent that provides conversation services
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The environment variable configuration passed to the agent runtime process.
	Environment *CreateCustomAgentRequestBodyEnvironment `json:"environment,omitempty" xml:"environment,omitempty" type:"Struct"`
	// The optional HTTP health check configuration. If this object is provided, port is required, and unspecified items use default values. If omitted during creation, no health check is configured. If this top-level field is omitted during an update, the original configuration is retained. If null is passed, the health check is removed.
	HealthCheck *CreateCustomAgentRequestBodyHealthCheck `json:"healthCheck,omitempty" xml:"healthCheck,omitempty" type:"Struct"`
	// The agent log delivery configuration. You can specify the Simple Log Service project and Logstore. After a successful creation or update, verify that logs are delivered properly.
	Logging *CreateCustomAgentRequestBodyLogging `json:"logging,omitempty" xml:"logging,omitempty" type:"Struct"`
	// The name of the custom agent. The length must be 1 to 128 characters after leading and trailing whitespace is removed. The name must be unique within the same workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-custom-agent
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The network access configuration of the agent. If this object is provided, omitted access items are treated as allowed. The actual accessible range is still limited by the workspace network configuration. To retain the original restrictions during an update, specify both accessInternet and accessVpc.
	Network *CreateCustomAgentRequestBodyNetwork `json:"network,omitempty" xml:"network,omitempty" type:"Struct"`
	// The list of OSS mounts. The total number of OSS and AgenticFS mounts cannot exceed 10. For each item, you must specify bucketName, path, and mountPath. Mount targets must not be duplicated or overlap. The entire list is replaced during an update. Pass an empty array [] to clear the list.
	OssMounts []*CreateCustomAgentRequestBodyOssMounts `json:"ossMounts,omitempty" xml:"ossMounts,omitempty" type:"Repeated"`
	// The list of invocation protocols supported by the agent. Each item includes the protocol type and optional paths and request headers. When updating, replace the entire list.
	Protocol []*CreateCustomAgentRequestBodyProtocol `json:"protocol,omitempty" xml:"protocol,omitempty" type:"Repeated"`
	// The ARN of the RAM role used when the agent accesses user cloud resources. The role must have permissions to access the cloud resources.
	//
	// This parameter is required.
	//
	// example:
	//
	// acs:ram::1234567890123456:role/agentcore-runtime-role
	RoleArn *string `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	// The runtime configuration, including compute specifications, session policies, and an optional elastic policy. compute and sessionPolicy must be specified together.
	//
	// This parameter is required.
	Runtime *CreateCustomAgentRequestBodyRuntime `json:"runtime,omitempty" xml:"runtime,omitempty" type:"Struct"`
}

func (s CreateCustomAgentRequestBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentRequestBody) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentRequestBody) GetAgenticFsMounts() []*CreateCustomAgentRequestBodyAgenticFsMounts {
	return s.AgenticFsMounts
}

func (s *CreateCustomAgentRequestBody) GetArtifact() *CreateCustomAgentRequestBodyArtifact {
	return s.Artifact
}

func (s *CreateCustomAgentRequestBody) GetDescription() *string {
	return s.Description
}

func (s *CreateCustomAgentRequestBody) GetEnvironment() *CreateCustomAgentRequestBodyEnvironment {
	return s.Environment
}

func (s *CreateCustomAgentRequestBody) GetHealthCheck() *CreateCustomAgentRequestBodyHealthCheck {
	return s.HealthCheck
}

func (s *CreateCustomAgentRequestBody) GetLogging() *CreateCustomAgentRequestBodyLogging {
	return s.Logging
}

func (s *CreateCustomAgentRequestBody) GetName() *string {
	return s.Name
}

func (s *CreateCustomAgentRequestBody) GetNetwork() *CreateCustomAgentRequestBodyNetwork {
	return s.Network
}

func (s *CreateCustomAgentRequestBody) GetOssMounts() []*CreateCustomAgentRequestBodyOssMounts {
	return s.OssMounts
}

func (s *CreateCustomAgentRequestBody) GetProtocol() []*CreateCustomAgentRequestBodyProtocol {
	return s.Protocol
}

func (s *CreateCustomAgentRequestBody) GetRoleArn() *string {
	return s.RoleArn
}

func (s *CreateCustomAgentRequestBody) GetRuntime() *CreateCustomAgentRequestBodyRuntime {
	return s.Runtime
}

func (s *CreateCustomAgentRequestBody) SetAgenticFsMounts(v []*CreateCustomAgentRequestBodyAgenticFsMounts) *CreateCustomAgentRequestBody {
	s.AgenticFsMounts = v
	return s
}

func (s *CreateCustomAgentRequestBody) SetArtifact(v *CreateCustomAgentRequestBodyArtifact) *CreateCustomAgentRequestBody {
	s.Artifact = v
	return s
}

func (s *CreateCustomAgentRequestBody) SetDescription(v string) *CreateCustomAgentRequestBody {
	s.Description = &v
	return s
}

func (s *CreateCustomAgentRequestBody) SetEnvironment(v *CreateCustomAgentRequestBodyEnvironment) *CreateCustomAgentRequestBody {
	s.Environment = v
	return s
}

func (s *CreateCustomAgentRequestBody) SetHealthCheck(v *CreateCustomAgentRequestBodyHealthCheck) *CreateCustomAgentRequestBody {
	s.HealthCheck = v
	return s
}

func (s *CreateCustomAgentRequestBody) SetLogging(v *CreateCustomAgentRequestBodyLogging) *CreateCustomAgentRequestBody {
	s.Logging = v
	return s
}

func (s *CreateCustomAgentRequestBody) SetName(v string) *CreateCustomAgentRequestBody {
	s.Name = &v
	return s
}

func (s *CreateCustomAgentRequestBody) SetNetwork(v *CreateCustomAgentRequestBodyNetwork) *CreateCustomAgentRequestBody {
	s.Network = v
	return s
}

func (s *CreateCustomAgentRequestBody) SetOssMounts(v []*CreateCustomAgentRequestBodyOssMounts) *CreateCustomAgentRequestBody {
	s.OssMounts = v
	return s
}

func (s *CreateCustomAgentRequestBody) SetProtocol(v []*CreateCustomAgentRequestBodyProtocol) *CreateCustomAgentRequestBody {
	s.Protocol = v
	return s
}

func (s *CreateCustomAgentRequestBody) SetRoleArn(v string) *CreateCustomAgentRequestBody {
	s.RoleArn = &v
	return s
}

func (s *CreateCustomAgentRequestBody) SetRuntime(v *CreateCustomAgentRequestBodyRuntime) *CreateCustomAgentRequestBody {
	s.Runtime = v
	return s
}

func (s *CreateCustomAgentRequestBody) Validate() error {
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

type CreateCustomAgentRequestBodyAgenticFsMounts struct {
	// The subdirectory under /mnt/agenticfs/ in the container. This parameter is required for each mount item. Mount targets must not be duplicated or overlap in a parent-child relationship.
	//
	// example:
	//
	// /mnt/agenticfs/data
	MountPath *string `json:"mountPath,omitempty" xml:"mountPath,omitempty"`
	// The existing non-empty relative directory under the access point. This parameter is required for each mount item. Root directories, absolute paths, and parent directory references are not allowed.
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
	// The access point domain name. This parameter is required for each mount item. Use the DomainName from the NAS ListAccessPoints response. Do not include the protocol, port, or path.
	//
	// example:
	//
	// ap-0123456789abcdef0.0123456789-vlm36.cn-hangzhou.nas.aliyuncs.com
	Server *string `json:"server,omitempty" xml:"server,omitempty"`
}

func (s CreateCustomAgentRequestBodyAgenticFsMounts) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentRequestBodyAgenticFsMounts) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentRequestBodyAgenticFsMounts) GetMountPath() *string {
	return s.MountPath
}

func (s *CreateCustomAgentRequestBodyAgenticFsMounts) GetPath() *string {
	return s.Path
}

func (s *CreateCustomAgentRequestBodyAgenticFsMounts) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *CreateCustomAgentRequestBodyAgenticFsMounts) GetServer() *string {
	return s.Server
}

func (s *CreateCustomAgentRequestBodyAgenticFsMounts) SetMountPath(v string) *CreateCustomAgentRequestBodyAgenticFsMounts {
	s.MountPath = &v
	return s
}

func (s *CreateCustomAgentRequestBodyAgenticFsMounts) SetPath(v string) *CreateCustomAgentRequestBodyAgenticFsMounts {
	s.Path = &v
	return s
}

func (s *CreateCustomAgentRequestBodyAgenticFsMounts) SetReadOnly(v bool) *CreateCustomAgentRequestBodyAgenticFsMounts {
	s.ReadOnly = &v
	return s
}

func (s *CreateCustomAgentRequestBodyAgenticFsMounts) SetServer(v string) *CreateCustomAgentRequestBodyAgenticFsMounts {
	s.Server = &v
	return s
}

func (s *CreateCustomAgentRequestBodyAgenticFsMounts) Validate() error {
	return dara.Validate(s)
}

type CreateCustomAgentRequestBodyArtifact struct {
	// The code package deployment configuration. This is not currently supported. Omit this parameter and use artifact.container.
	Code *CreateCustomAgentRequestBodyArtifactCode `json:"code,omitempty" xml:"code,omitempty" type:"Struct"`
	// The ACR container image configuration. You must specify type, instanceId, namespace, repo, and tag. Do not configure code at the same time.
	Container *CreateCustomAgentRequestBodyArtifactContainer `json:"container,omitempty" xml:"container,omitempty" type:"Struct"`
	// The TCP port on which the agent service listens. Valid values: 1 to 65535.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8080
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
}

func (s CreateCustomAgentRequestBodyArtifact) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentRequestBodyArtifact) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentRequestBodyArtifact) GetCode() *CreateCustomAgentRequestBodyArtifactCode {
	return s.Code
}

func (s *CreateCustomAgentRequestBodyArtifact) GetContainer() *CreateCustomAgentRequestBodyArtifactContainer {
	return s.Container
}

func (s *CreateCustomAgentRequestBodyArtifact) GetPort() *int32 {
	return s.Port
}

func (s *CreateCustomAgentRequestBodyArtifact) SetCode(v *CreateCustomAgentRequestBodyArtifactCode) *CreateCustomAgentRequestBodyArtifact {
	s.Code = v
	return s
}

func (s *CreateCustomAgentRequestBodyArtifact) SetContainer(v *CreateCustomAgentRequestBodyArtifactContainer) *CreateCustomAgentRequestBodyArtifact {
	s.Container = v
	return s
}

func (s *CreateCustomAgentRequestBodyArtifact) SetPort(v int32) *CreateCustomAgentRequestBodyArtifact {
	s.Port = &v
	return s
}

func (s *CreateCustomAgentRequestBodyArtifact) Validate() error {
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

type CreateCustomAgentRequestBodyArtifactCode struct {
	// The startup entry point of the code package. Code package deployment is not currently supported. Omit this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// python app.py
	Entrypoint *string `json:"entrypoint,omitempty" xml:"entrypoint,omitempty"`
	// The language runtime identifier of the code package. Code package deployment is not currently supported. Omit this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// python
	LanguageRuntime *string `json:"languageRuntime,omitempty" xml:"languageRuntime,omitempty"`
	// The source URI of the code package. Code package deployment is not currently supported. Omit this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://my-agent-bucket/packages/agent.zip
	SourceUri *string `json:"sourceUri,omitempty" xml:"sourceUri,omitempty"`
}

func (s CreateCustomAgentRequestBodyArtifactCode) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentRequestBodyArtifactCode) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentRequestBodyArtifactCode) GetEntrypoint() *string {
	return s.Entrypoint
}

func (s *CreateCustomAgentRequestBodyArtifactCode) GetLanguageRuntime() *string {
	return s.LanguageRuntime
}

func (s *CreateCustomAgentRequestBodyArtifactCode) GetSourceUri() *string {
	return s.SourceUri
}

func (s *CreateCustomAgentRequestBodyArtifactCode) SetEntrypoint(v string) *CreateCustomAgentRequestBodyArtifactCode {
	s.Entrypoint = &v
	return s
}

func (s *CreateCustomAgentRequestBodyArtifactCode) SetLanguageRuntime(v string) *CreateCustomAgentRequestBodyArtifactCode {
	s.LanguageRuntime = &v
	return s
}

func (s *CreateCustomAgentRequestBodyArtifactCode) SetSourceUri(v string) *CreateCustomAgentRequestBodyArtifactCode {
	s.SourceUri = &v
	return s
}

func (s *CreateCustomAgentRequestBodyArtifactCode) Validate() error {
	return dara.Validate(s)
}

type CreateCustomAgentRequestBodyArtifactContainer struct {
	// The optional image startup command. If specified, it must be a non-empty string. If omitted, the default startup method of the image is used.
	//
	// example:
	//
	// python app.py
	Entrypoint *string `json:"entrypoint,omitempty" xml:"entrypoint,omitempty"`
	// The ID of the ACR instance to which the image repository belongs. This parameter cannot be empty.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-a1b2c3d4e5f6
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The namespace to which the ACR image repository belongs. This parameter cannot be empty.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-agent
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The name of the image repository in the ACR namespace. This parameter cannot be empty.
	//
	// This parameter is required.
	//
	// example:
	//
	// chat-agent
	Repo *string `json:"repo,omitempty" xml:"repo,omitempty"`
	// The image tag used for deployment. This parameter cannot be empty. When updating the tag, you must also specify artifact.port and other required image fields.
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

func (s CreateCustomAgentRequestBodyArtifactContainer) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentRequestBodyArtifactContainer) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentRequestBodyArtifactContainer) GetEntrypoint() *string {
	return s.Entrypoint
}

func (s *CreateCustomAgentRequestBodyArtifactContainer) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateCustomAgentRequestBodyArtifactContainer) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateCustomAgentRequestBodyArtifactContainer) GetRepo() *string {
	return s.Repo
}

func (s *CreateCustomAgentRequestBodyArtifactContainer) GetTag() *string {
	return s.Tag
}

func (s *CreateCustomAgentRequestBodyArtifactContainer) GetType() *string {
	return s.Type
}

func (s *CreateCustomAgentRequestBodyArtifactContainer) SetEntrypoint(v string) *CreateCustomAgentRequestBodyArtifactContainer {
	s.Entrypoint = &v
	return s
}

func (s *CreateCustomAgentRequestBodyArtifactContainer) SetInstanceId(v string) *CreateCustomAgentRequestBodyArtifactContainer {
	s.InstanceId = &v
	return s
}

func (s *CreateCustomAgentRequestBodyArtifactContainer) SetNamespace(v string) *CreateCustomAgentRequestBodyArtifactContainer {
	s.Namespace = &v
	return s
}

func (s *CreateCustomAgentRequestBodyArtifactContainer) SetRepo(v string) *CreateCustomAgentRequestBodyArtifactContainer {
	s.Repo = &v
	return s
}

func (s *CreateCustomAgentRequestBodyArtifactContainer) SetTag(v string) *CreateCustomAgentRequestBodyArtifactContainer {
	s.Tag = &v
	return s
}

func (s *CreateCustomAgentRequestBodyArtifactContainer) SetType(v string) *CreateCustomAgentRequestBodyArtifactContainer {
	s.Type = &v
	return s
}

func (s *CreateCustomAgentRequestBodyArtifactContainer) Validate() error {
	return dara.Validate(s)
}

type CreateCustomAgentRequestBodyEnvironment struct {
	// The list of environment variables. Each item includes a name and a value. The entire list is replaced during an update.
	Variables []*CreateCustomAgentRequestBodyEnvironmentVariables `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
}

func (s CreateCustomAgentRequestBodyEnvironment) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentRequestBodyEnvironment) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentRequestBodyEnvironment) GetVariables() []*CreateCustomAgentRequestBodyEnvironmentVariables {
	return s.Variables
}

func (s *CreateCustomAgentRequestBodyEnvironment) SetVariables(v []*CreateCustomAgentRequestBodyEnvironmentVariables) *CreateCustomAgentRequestBodyEnvironment {
	s.Variables = v
	return s
}

func (s *CreateCustomAgentRequestBodyEnvironment) Validate() error {
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

type CreateCustomAgentRequestBodyEnvironmentVariables struct {
	// The name of the environment variable. This parameter cannot be empty and must be unique within the current environment variable list.
	//
	// This parameter is required.
	//
	// example:
	//
	// LOG_LEVEL
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The value of the environment variable. It must be a string. Empty strings are allowed.
	//
	// This parameter is required.
	//
	// example:
	//
	// info
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateCustomAgentRequestBodyEnvironmentVariables) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentRequestBodyEnvironmentVariables) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentRequestBodyEnvironmentVariables) GetName() *string {
	return s.Name
}

func (s *CreateCustomAgentRequestBodyEnvironmentVariables) GetValue() *string {
	return s.Value
}

func (s *CreateCustomAgentRequestBodyEnvironmentVariables) SetName(v string) *CreateCustomAgentRequestBodyEnvironmentVariables {
	s.Name = &v
	return s
}

func (s *CreateCustomAgentRequestBodyEnvironmentVariables) SetValue(v string) *CreateCustomAgentRequestBodyEnvironmentVariables {
	s.Value = &v
	return s
}

func (s *CreateCustomAgentRequestBodyEnvironmentVariables) Validate() error {
	return dara.Validate(s)
}

type CreateCustomAgentRequestBodyHealthCheck struct {
	// This parameter does not currently take effect. Omit this parameter.
	//
	// example:
	//
	// 1
	HealthyThreshold *int32 `json:"healthyThreshold,omitempty" xml:"healthyThreshold,omitempty"`
	// The interval between two health checks, in seconds. It must be a positive integer. If omitted, the default value is 10.
	//
	// example:
	//
	// 10
	IntervalSeconds *int32 `json:"intervalSeconds,omitempty" xml:"intervalSeconds,omitempty"`
	// The HTTP health check path. It must start with a forward slash (/). If omitted, the default value is /healthz.
	//
	// example:
	//
	// /healthz
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// The port accessed for health checks. This parameter is required when configuring a health check. Valid values: 1 to 65535.
	//
	// example:
	//
	// 8080
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// The timeout period for a single health check, in seconds. It must be a positive integer and less than intervalSeconds. If omitted, the default value is 2.
	//
	// example:
	//
	// 2
	TimeoutSeconds *int32 `json:"timeoutSeconds,omitempty" xml:"timeoutSeconds,omitempty"`
	// The number of consecutive health check failures required to mark the instance as unhealthy. It must be a positive integer. If omitted, the default value is 3.
	//
	// example:
	//
	// 3
	UnhealthyThreshold *int32 `json:"unhealthyThreshold,omitempty" xml:"unhealthyThreshold,omitempty"`
}

func (s CreateCustomAgentRequestBodyHealthCheck) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentRequestBodyHealthCheck) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentRequestBodyHealthCheck) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *CreateCustomAgentRequestBodyHealthCheck) GetIntervalSeconds() *int32 {
	return s.IntervalSeconds
}

func (s *CreateCustomAgentRequestBodyHealthCheck) GetPath() *string {
	return s.Path
}

func (s *CreateCustomAgentRequestBodyHealthCheck) GetPort() *int32 {
	return s.Port
}

func (s *CreateCustomAgentRequestBodyHealthCheck) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *CreateCustomAgentRequestBodyHealthCheck) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *CreateCustomAgentRequestBodyHealthCheck) SetHealthyThreshold(v int32) *CreateCustomAgentRequestBodyHealthCheck {
	s.HealthyThreshold = &v
	return s
}

func (s *CreateCustomAgentRequestBodyHealthCheck) SetIntervalSeconds(v int32) *CreateCustomAgentRequestBodyHealthCheck {
	s.IntervalSeconds = &v
	return s
}

func (s *CreateCustomAgentRequestBodyHealthCheck) SetPath(v string) *CreateCustomAgentRequestBodyHealthCheck {
	s.Path = &v
	return s
}

func (s *CreateCustomAgentRequestBodyHealthCheck) SetPort(v int32) *CreateCustomAgentRequestBodyHealthCheck {
	s.Port = &v
	return s
}

func (s *CreateCustomAgentRequestBodyHealthCheck) SetTimeoutSeconds(v int32) *CreateCustomAgentRequestBodyHealthCheck {
	s.TimeoutSeconds = &v
	return s
}

func (s *CreateCustomAgentRequestBodyHealthCheck) SetUnhealthyThreshold(v int32) *CreateCustomAgentRequestBodyHealthCheck {
	s.UnhealthyThreshold = &v
	return s
}

func (s *CreateCustomAgentRequestBodyHealthCheck) Validate() error {
	return dara.Validate(s)
}

type CreateCustomAgentRequestBodyLogging struct {
	// The Simple Log Service log delivery destination. When configuring this object, neither project nor logstore can be empty.
	Sls *CreateCustomAgentRequestBodyLoggingSls `json:"sls,omitempty" xml:"sls,omitempty" type:"Struct"`
}

func (s CreateCustomAgentRequestBodyLogging) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentRequestBodyLogging) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentRequestBodyLogging) GetSls() *CreateCustomAgentRequestBodyLoggingSls {
	return s.Sls
}

func (s *CreateCustomAgentRequestBodyLogging) SetSls(v *CreateCustomAgentRequestBodyLoggingSls) *CreateCustomAgentRequestBodyLogging {
	s.Sls = v
	return s
}

func (s *CreateCustomAgentRequestBodyLogging) Validate() error {
	if s.Sls != nil {
		if err := s.Sls.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCustomAgentRequestBodyLoggingSls struct {
	// The name of the Simple Log Service Logstore that receives the agent logs.
	//
	// example:
	//
	// agent-runtime
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// The name of the Simple Log Service project that receives the agent logs.
	//
	// example:
	//
	// my-agent-logs
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
}

func (s CreateCustomAgentRequestBodyLoggingSls) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentRequestBodyLoggingSls) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentRequestBodyLoggingSls) GetLogstore() *string {
	return s.Logstore
}

func (s *CreateCustomAgentRequestBodyLoggingSls) GetProject() *string {
	return s.Project
}

func (s *CreateCustomAgentRequestBodyLoggingSls) SetLogstore(v string) *CreateCustomAgentRequestBodyLoggingSls {
	s.Logstore = &v
	return s
}

func (s *CreateCustomAgentRequestBodyLoggingSls) SetProject(v string) *CreateCustomAgentRequestBodyLoggingSls {
	s.Project = &v
	return s
}

func (s *CreateCustomAgentRequestBodyLoggingSls) Validate() error {
	return dara.Validate(s)
}

type CreateCustomAgentRequestBodyNetwork struct {
	// The Internet access configuration. If this access item is omitted, Internet access is not restricted on the agent side. However, the workspace must still allow Internet access.
	AccessInternet *CreateCustomAgentRequestBodyNetworkAccessInternet `json:"accessInternet,omitempty" xml:"accessInternet,omitempty" type:"Struct"`
	// The user VPC access configuration. If this access item is omitted, VPC access is not restricted on the agent side. However, the workspace must still be configured to allow access to the user VPC.
	AccessVpc *CreateCustomAgentRequestBodyNetworkAccessVpc `json:"accessVpc,omitempty" xml:"accessVpc,omitempty" type:"Struct"`
}

func (s CreateCustomAgentRequestBodyNetwork) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentRequestBodyNetwork) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentRequestBodyNetwork) GetAccessInternet() *CreateCustomAgentRequestBodyNetworkAccessInternet {
	return s.AccessInternet
}

func (s *CreateCustomAgentRequestBodyNetwork) GetAccessVpc() *CreateCustomAgentRequestBodyNetworkAccessVpc {
	return s.AccessVpc
}

func (s *CreateCustomAgentRequestBodyNetwork) SetAccessInternet(v *CreateCustomAgentRequestBodyNetworkAccessInternet) *CreateCustomAgentRequestBodyNetwork {
	s.AccessInternet = v
	return s
}

func (s *CreateCustomAgentRequestBodyNetwork) SetAccessVpc(v *CreateCustomAgentRequestBodyNetworkAccessVpc) *CreateCustomAgentRequestBodyNetwork {
	s.AccessVpc = v
	return s
}

func (s *CreateCustomAgentRequestBodyNetwork) Validate() error {
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

type CreateCustomAgentRequestBodyNetworkAccessInternet struct {
	// Specifies whether to allow the agent to access the Internet. A Boolean value is required when configuring accessInternet. For this to take effect, the workspace must also allow Internet access.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s CreateCustomAgentRequestBodyNetworkAccessInternet) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentRequestBodyNetworkAccessInternet) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentRequestBodyNetworkAccessInternet) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateCustomAgentRequestBodyNetworkAccessInternet) SetEnabled(v bool) *CreateCustomAgentRequestBodyNetworkAccessInternet {
	s.Enabled = &v
	return s
}

func (s *CreateCustomAgentRequestBodyNetworkAccessInternet) Validate() error {
	return dara.Validate(s)
}

type CreateCustomAgentRequestBodyNetworkAccessVpc struct {
	// Specifies whether to allow the agent to access the user VPC configured for the workspace. A Boolean value is required when configuring accessVpc. For this to take effect, the workspace must also allow VPC access.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s CreateCustomAgentRequestBodyNetworkAccessVpc) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentRequestBodyNetworkAccessVpc) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentRequestBodyNetworkAccessVpc) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateCustomAgentRequestBodyNetworkAccessVpc) SetEnabled(v bool) *CreateCustomAgentRequestBodyNetworkAccessVpc {
	s.Enabled = &v
	return s
}

func (s *CreateCustomAgentRequestBodyNetworkAccessVpc) Validate() error {
	return dara.Validate(s)
}

type CreateCustomAgentRequestBodyOssMounts struct {
	// The name of the OSS bucket. This parameter is required for each mount item. The name must be 3 to 63 characters in length and can contain only lowercase letters, digits, and hyphens. It cannot start or end with a hyphen.
	//
	// example:
	//
	// my-agent-data
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// The subdirectory under /mnt/oss/ in the container. This parameter is required for each mount item. Mount targets must not be duplicated or overlap in a parent-child relationship.
	//
	// example:
	//
	// /mnt/oss/train
	MountPath *string `json:"mountPath,omitempty" xml:"mountPath,omitempty"`
	// The non-empty object directory prefix in the bucket, such as datasets/train. Leading and trailing forward slashes (/) are removed. Root directories are not supported. Do not use parent directory references.
	//
	// example:
	//
	// datasets/train
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// Specifies whether the mount is read-only. A value of true indicates read-only, and false indicates read and write. If omitted, the default value is false.
	ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
}

func (s CreateCustomAgentRequestBodyOssMounts) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentRequestBodyOssMounts) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentRequestBodyOssMounts) GetBucketName() *string {
	return s.BucketName
}

func (s *CreateCustomAgentRequestBodyOssMounts) GetMountPath() *string {
	return s.MountPath
}

func (s *CreateCustomAgentRequestBodyOssMounts) GetPath() *string {
	return s.Path
}

func (s *CreateCustomAgentRequestBodyOssMounts) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *CreateCustomAgentRequestBodyOssMounts) SetBucketName(v string) *CreateCustomAgentRequestBodyOssMounts {
	s.BucketName = &v
	return s
}

func (s *CreateCustomAgentRequestBodyOssMounts) SetMountPath(v string) *CreateCustomAgentRequestBodyOssMounts {
	s.MountPath = &v
	return s
}

func (s *CreateCustomAgentRequestBodyOssMounts) SetPath(v string) *CreateCustomAgentRequestBodyOssMounts {
	s.Path = &v
	return s
}

func (s *CreateCustomAgentRequestBodyOssMounts) SetReadOnly(v bool) *CreateCustomAgentRequestBodyOssMounts {
	s.ReadOnly = &v
	return s
}

func (s *CreateCustomAgentRequestBodyOssMounts) Validate() error {
	return dara.Validate(s)
}

type CreateCustomAgentRequestBodyProtocol struct {
	// The request header configuration for the current protocol. The key is a non-empty request header name, and the value must be a string.
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	// The list of API paths corresponding to the current protocol.
	Paths []*CreateCustomAgentRequestBodyProtocolPaths `json:"paths,omitempty" xml:"paths,omitempty" type:"Repeated"`
	// The invocation protocol type. Valid values: agui: AG-UI protocol; openai-chat-completion: OpenAI Chat Completions compatible protocol; custom: custom protocol.
	//
	// This parameter is required.
	//
	// example:
	//
	// agui
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateCustomAgentRequestBodyProtocol) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentRequestBodyProtocol) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentRequestBodyProtocol) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCustomAgentRequestBodyProtocol) GetPaths() []*CreateCustomAgentRequestBodyProtocolPaths {
	return s.Paths
}

func (s *CreateCustomAgentRequestBodyProtocol) GetType() *string {
	return s.Type
}

func (s *CreateCustomAgentRequestBodyProtocol) SetHeaders(v map[string]*string) *CreateCustomAgentRequestBodyProtocol {
	s.Headers = v
	return s
}

func (s *CreateCustomAgentRequestBodyProtocol) SetPaths(v []*CreateCustomAgentRequestBodyProtocolPaths) *CreateCustomAgentRequestBodyProtocol {
	s.Paths = v
	return s
}

func (s *CreateCustomAgentRequestBodyProtocol) SetType(v string) *CreateCustomAgentRequestBodyProtocol {
	s.Type = &v
	return s
}

func (s *CreateCustomAgentRequestBodyProtocol) Validate() error {
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

type CreateCustomAgentRequestBodyProtocolPaths struct {
	// The usage description of the current API path.
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

func (s CreateCustomAgentRequestBodyProtocolPaths) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentRequestBodyProtocolPaths) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentRequestBodyProtocolPaths) GetDescription() *string {
	return s.Description
}

func (s *CreateCustomAgentRequestBodyProtocolPaths) GetName() *string {
	return s.Name
}

func (s *CreateCustomAgentRequestBodyProtocolPaths) SetDescription(v string) *CreateCustomAgentRequestBodyProtocolPaths {
	s.Description = &v
	return s
}

func (s *CreateCustomAgentRequestBodyProtocolPaths) SetName(v string) *CreateCustomAgentRequestBodyProtocolPaths {
	s.Name = &v
	return s
}

func (s *CreateCustomAgentRequestBodyProtocolPaths) Validate() error {
	return dara.Validate(s)
}

type CreateCustomAgentRequestBodyRuntime struct {
	// The compute resource configuration for the agent runtime instance.
	//
	// This parameter is required.
	Compute *CreateCustomAgentRequestBodyRuntimeCompute `json:"compute,omitempty" xml:"compute,omitempty" type:"Struct"`
	// The optional elastic policy. When runtime is provided, if hpa or its fields are omitted, the default values are used: enabled is false, the minimum number of instances is 1, the maximum number of instances is 10, the maximum number of concurrent sessions per instance is 100, and the session idle timeout is 600 seconds. To retain the original policy during an update, provide the complete hpa configuration.
	Hpa *CreateCustomAgentRequestBodyRuntimeHpa `json:"hpa,omitempty" xml:"hpa,omitempty" type:"Struct"`
	// The session policy configuration, which is used to set the session association method for requests.
	//
	// This parameter is required.
	SessionPolicy *CreateCustomAgentRequestBodyRuntimeSessionPolicy `json:"sessionPolicy,omitempty" xml:"sessionPolicy,omitempty" type:"Struct"`
}

func (s CreateCustomAgentRequestBodyRuntime) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentRequestBodyRuntime) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentRequestBodyRuntime) GetCompute() *CreateCustomAgentRequestBodyRuntimeCompute {
	return s.Compute
}

func (s *CreateCustomAgentRequestBodyRuntime) GetHpa() *CreateCustomAgentRequestBodyRuntimeHpa {
	return s.Hpa
}

func (s *CreateCustomAgentRequestBodyRuntime) GetSessionPolicy() *CreateCustomAgentRequestBodyRuntimeSessionPolicy {
	return s.SessionPolicy
}

func (s *CreateCustomAgentRequestBodyRuntime) SetCompute(v *CreateCustomAgentRequestBodyRuntimeCompute) *CreateCustomAgentRequestBodyRuntime {
	s.Compute = v
	return s
}

func (s *CreateCustomAgentRequestBodyRuntime) SetHpa(v *CreateCustomAgentRequestBodyRuntimeHpa) *CreateCustomAgentRequestBodyRuntime {
	s.Hpa = v
	return s
}

func (s *CreateCustomAgentRequestBodyRuntime) SetSessionPolicy(v *CreateCustomAgentRequestBodyRuntimeSessionPolicy) *CreateCustomAgentRequestBodyRuntime {
	s.SessionPolicy = v
	return s
}

func (s *CreateCustomAgentRequestBodyRuntime) Validate() error {
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

type CreateCustomAgentRequestBodyRuntimeCompute struct {
	// The compute specification. Valid values: INTERACTIVE: 1 vCPU and 2 GiB of memory; STANDARD: 2 vCPUs and 4 GiB of memory; HEAVY: 4 vCPUs and 8 GiB of memory.
	//
	// This parameter is required.
	//
	// example:
	//
	// STANDARD
	ComputeClass *string `json:"computeClass,omitempty" xml:"computeClass,omitempty"`
}

func (s CreateCustomAgentRequestBodyRuntimeCompute) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentRequestBodyRuntimeCompute) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentRequestBodyRuntimeCompute) GetComputeClass() *string {
	return s.ComputeClass
}

func (s *CreateCustomAgentRequestBodyRuntimeCompute) SetComputeClass(v string) *CreateCustomAgentRequestBodyRuntimeCompute {
	s.ComputeClass = &v
	return s
}

func (s *CreateCustomAgentRequestBodyRuntimeCompute) Validate() error {
	return dara.Validate(s)
}

type CreateCustomAgentRequestBodyRuntimeHpa struct {
	// Specifies whether to enable elastic scaling. A value of true indicates that elastic scaling is enabled. A value of false indicates that elastic scaling is disabled. When runtime is provided, if this field is omitted, the default value is false.
	//
	// This parameter is required.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The maximum number of concurrent sessions for a single runtime instance. Valid values: 1 to 100. When runtime is provided, if this field is omitted, the default value is 100.
	//
	// example:
	//
	// 100
	MaxConcurrentSessionsPerSandbox *int32 `json:"maxConcurrentSessionsPerSandbox,omitempty" xml:"maxConcurrentSessionsPerSandbox,omitempty"`
	// The maximum number of runtime instances. Valid values: 1 to 100. The value cannot be less than minSandboxCount. When runtime is provided, if this field is omitted, the default value is 10.
	//
	// example:
	//
	// 10
	MaxSandboxCount *int32 `json:"maxSandboxCount,omitempty" xml:"maxSandboxCount,omitempty"`
	// The minimum number of runtime instances. Valid values: 1 to 100. The value cannot be greater than maxSandboxCount. When runtime is provided, if this field is omitted, the default value is 1.
	//
	// example:
	//
	// 1
	MinSandboxCount *int32 `json:"minSandboxCount,omitempty" xml:"minSandboxCount,omitempty"`
	// The session idle timeout in seconds. Valid values: 60 to 86400. When runtime is provided, if this field is omitted, the default value is 600.
	//
	// example:
	//
	// 600
	SessionTtlSeconds *int32 `json:"sessionTtlSeconds,omitempty" xml:"sessionTtlSeconds,omitempty"`
}

func (s CreateCustomAgentRequestBodyRuntimeHpa) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentRequestBodyRuntimeHpa) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentRequestBodyRuntimeHpa) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateCustomAgentRequestBodyRuntimeHpa) GetMaxConcurrentSessionsPerSandbox() *int32 {
	return s.MaxConcurrentSessionsPerSandbox
}

func (s *CreateCustomAgentRequestBodyRuntimeHpa) GetMaxSandboxCount() *int32 {
	return s.MaxSandboxCount
}

func (s *CreateCustomAgentRequestBodyRuntimeHpa) GetMinSandboxCount() *int32 {
	return s.MinSandboxCount
}

func (s *CreateCustomAgentRequestBodyRuntimeHpa) GetSessionTtlSeconds() *int32 {
	return s.SessionTtlSeconds
}

func (s *CreateCustomAgentRequestBodyRuntimeHpa) SetEnabled(v bool) *CreateCustomAgentRequestBodyRuntimeHpa {
	s.Enabled = &v
	return s
}

func (s *CreateCustomAgentRequestBodyRuntimeHpa) SetMaxConcurrentSessionsPerSandbox(v int32) *CreateCustomAgentRequestBodyRuntimeHpa {
	s.MaxConcurrentSessionsPerSandbox = &v
	return s
}

func (s *CreateCustomAgentRequestBodyRuntimeHpa) SetMaxSandboxCount(v int32) *CreateCustomAgentRequestBodyRuntimeHpa {
	s.MaxSandboxCount = &v
	return s
}

func (s *CreateCustomAgentRequestBodyRuntimeHpa) SetMinSandboxCount(v int32) *CreateCustomAgentRequestBodyRuntimeHpa {
	s.MinSandboxCount = &v
	return s
}

func (s *CreateCustomAgentRequestBodyRuntimeHpa) SetSessionTtlSeconds(v int32) *CreateCustomAgentRequestBodyRuntimeHpa {
	s.SessionTtlSeconds = &v
	return s
}

func (s *CreateCustomAgentRequestBodyRuntimeHpa) Validate() error {
	return dara.Validate(s)
}

type CreateCustomAgentRequestBodyRuntimeSessionPolicy struct {
	// The name of the HTTP request header used by the session policy. This parameter is required when type is set to ISOLATED_HEADER_FIELD and must be a valid HTTP request header name. Omit this parameter for other policies.
	//
	// example:
	//
	// X-Session-Id
	HeaderName *string `json:"headerName,omitempty" xml:"headerName,omitempty"`
	// The session policy type. Valid values: AFFINITY: session affinity; ISOLATED_HEADER_FIELD: associates sessions by using the request header specified by headerName. ISOLATED_COOKIE is not recommended. DISABLED is a compatibility option, and the response returns AFFINITY.
	//
	// This parameter is required.
	//
	// example:
	//
	// AFFINITY
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateCustomAgentRequestBodyRuntimeSessionPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentRequestBodyRuntimeSessionPolicy) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentRequestBodyRuntimeSessionPolicy) GetHeaderName() *string {
	return s.HeaderName
}

func (s *CreateCustomAgentRequestBodyRuntimeSessionPolicy) GetType() *string {
	return s.Type
}

func (s *CreateCustomAgentRequestBodyRuntimeSessionPolicy) SetHeaderName(v string) *CreateCustomAgentRequestBodyRuntimeSessionPolicy {
	s.HeaderName = &v
	return s
}

func (s *CreateCustomAgentRequestBodyRuntimeSessionPolicy) SetType(v string) *CreateCustomAgentRequestBodyRuntimeSessionPolicy {
	s.Type = &v
	return s
}

func (s *CreateCustomAgentRequestBodyRuntimeSessionPolicy) Validate() error {
	return dara.Validate(s)
}
