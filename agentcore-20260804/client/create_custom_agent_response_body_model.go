// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateCustomAgentResponseBody
	GetCode() *string
	SetData(v *CreateCustomAgentResponseBodyData) *CreateCustomAgentResponseBody
	GetData() *CreateCustomAgentResponseBodyData
	SetHttpStatusCode(v int32) *CreateCustomAgentResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateCustomAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateCustomAgentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateCustomAgentResponseBody
	GetSuccess() *bool
}

type CreateCustomAgentResponseBody struct {
	// The request result code. The value is SUCCESS when the request is successful.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The latest submitted configuration and status of the custom agent. A successful response indicates that the request is accepted, but the configuration may not have taken effect yet.
	Data *CreateCustomAgentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code of the request. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The description of the processing result for the request. A value of success indicates that the request was processed successfully.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The unique identifier of the request. You can use this ID to locate and troubleshoot issues.
	//
	// example:
	//
	// 550E8400-E29B-41D4-A716-446655440000
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was processed successfully. A value of true indicates that the request was processed successfully. Note that a successful creation, update, or deletion does not mean that the background operation is complete.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateCustomAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateCustomAgentResponseBody) GetData() *CreateCustomAgentResponseBodyData {
	return s.Data
}

func (s *CreateCustomAgentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateCustomAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateCustomAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCustomAgentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateCustomAgentResponseBody) SetCode(v string) *CreateCustomAgentResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCustomAgentResponseBody) SetData(v *CreateCustomAgentResponseBodyData) *CreateCustomAgentResponseBody {
	s.Data = v
	return s
}

func (s *CreateCustomAgentResponseBody) SetHttpStatusCode(v int32) *CreateCustomAgentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateCustomAgentResponseBody) SetMessage(v string) *CreateCustomAgentResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCustomAgentResponseBody) SetRequestId(v string) *CreateCustomAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCustomAgentResponseBody) SetSuccess(v bool) *CreateCustomAgentResponseBody {
	s.Success = &v
	return s
}

func (s *CreateCustomAgentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCustomAgentResponseBodyData struct {
	// The unique identifier of the custom agent. You can obtain this ID from the response of the CreateCustomAgent or ListCustomAgents operation.
	//
	// example:
	//
	// agent-550e8400e29b41d4a716446655440000
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The latest submitted list of additional AgenticFS mounts.
	AgenticFsMounts []*CreateCustomAgentResponseBodyDataAgenticFsMounts `json:"agenticFsMounts,omitempty" xml:"agenticFsMounts,omitempty" type:"Repeated"`
	// The latest submitted deployment artifact configuration, including the service listening port and the ACR container image. To check whether the configuration has taken effect, refer to the latestVersionStatus field.
	Artifact *CreateCustomAgentResponseBodyDataArtifact `json:"artifact,omitempty" xml:"artifact,omitempty" type:"Struct"`
	// The creation method of the agent. The value is CUSTOM for custom agents.
	//
	// example:
	//
	// CUSTOM
	CreateMode *string `json:"createMode,omitempty" xml:"createMode,omitempty"`
	// The creation time of the agent. The time is in ISO 8601 format. The trailing Z indicates the UTC time zone. Fractional seconds may be included.
	//
	// example:
	//
	// 2026-10-10T06:00:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The deployment method of the agent. Custom agents are deployed and managed by the platform. The value is MANAGED for custom agents.
	//
	// example:
	//
	// MANAGED
	DeployType *string `json:"deployType,omitempty" xml:"deployType,omitempty"`
	// The description of the agent.
	//
	// example:
	//
	// Custom agent that provides conversation services
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The environment variable configuration for the agent process.
	Environment *CreateCustomAgentResponseBodyDataEnvironment `json:"environment,omitempty" xml:"environment,omitempty" type:"Struct"`
	// The HTTP health check configuration of the agent. This field is empty if no health check is configured.
	HealthCheck *CreateCustomAgentResponseBodyDataHealthCheck `json:"healthCheck,omitempty" xml:"healthCheck,omitempty" type:"Struct"`
	// The waiting or execution error of the latest configuration version. This field is returned only when the version is waiting for a retry or fails to execute.
	LatestExecutionError *CreateCustomAgentResponseBodyDataLatestExecutionError `json:"latestExecutionError,omitempty" xml:"latestExecutionError,omitempty" type:"Struct"`
	// The latest configuration version number. Updates with the same configuration may retain the original version number. This field does not indicate that the configuration has taken effect.
	//
	// example:
	//
	// 1
	LatestSpecVersion *int64 `json:"latestSpecVersion,omitempty" xml:"latestSpecVersion,omitempty"`
	// The execution status corresponding to latestSpecVersion. Valid values:
	//
	// - pending: Waiting for execution.
	//
	// - processing: Executing.
	//
	// - waiting_retry: Waiting for a retry.
	//
	// - succeeded: Execution succeeded.
	//
	// - failed: Execution failed.
	//
	// - superseded: Superseded by a subsequent version.
	//
	// example:
	//
	// succeeded
	LatestVersionStatus *string `json:"latestVersionStatus,omitempty" xml:"latestVersionStatus,omitempty"`
	// The latest submitted log delivery configuration.
	Logging *CreateCustomAgentResponseBodyDataLogging `json:"logging,omitempty" xml:"logging,omitempty" type:"Struct"`
	// The name of the custom agent.
	//
	// example:
	//
	// my-custom-agent
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The network access configuration of the agent. The actual access scope is also restricted by the network configuration of the workspace.
	Network *CreateCustomAgentResponseBodyDataNetwork `json:"network,omitempty" xml:"network,omitempty" type:"Struct"`
	// The latest submitted list of OSS directory mounts.
	OssMounts []*CreateCustomAgentResponseBodyDataOssMounts `json:"ossMounts,omitempty" xml:"ossMounts,omitempty" type:"Repeated"`
	// The list of invocation protocol configurations for the agent.
	Protocol []*CreateCustomAgentResponseBodyDataProtocol `json:"protocol,omitempty" xml:"protocol,omitempty" type:"Repeated"`
	// The region ID of the workspace to which the agent belongs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The ARN of the RAM role used by the agent to access user cloud resources.
	//
	// example:
	//
	// acs:ram::1234567890123456:role/agentcore-runtime-role
	RoleArn *string `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	// The runtime configurations of the latest commit, including compute specifications, session policy, and elastic policy.
	Runtime *CreateCustomAgentResponseBodyDataRuntime `json:"runtime,omitempty" xml:"runtime,omitempty" type:"Struct"`
	// The instance counts grouped by runtime instance phase. Keys include PENDING (creating or initializing), RUNNING (running), HIBERNATING (hibernating), HIBERNATED (hibernated), RESUMING (resuming), TERMINATING (terminating), and FAILED (failed). Only phases that actually occur are returned. Missing keys are treated as 0.
	SandboxPhaseCounts map[string]*int64 `json:"sandboxPhaseCounts,omitempty" xml:"sandboxPhaseCounts,omitempty"`
	// The overall status of the agent. Valid values:
	//
	// - Creating: The agent is being created.
	//
	// - Running: The agent is running.
	//
	// - Failed: The agent failed.
	//
	// - Updating: The agent is being updated.
	//
	// - Deleting: The agent is being deleted.
	//
	// - Deleted: The agent is deleted.
	//
	// To query the execution progress of the latest configuration, check latestVersionStatus.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the agent was last updated. The time is in ISO 8601 format. A trailing Z indicates the UTC time zone. Fractional seconds may be included.
	//
	// example:
	//
	// 2026-10-10T06:05:00Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// The ID of the workspace to which the custom agent belongs.
	//
	// example:
	//
	// ws-a1b2c3d4e5f6
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s CreateCustomAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentResponseBodyData) GetAgentId() *string {
	return s.AgentId
}

func (s *CreateCustomAgentResponseBodyData) GetAgenticFsMounts() []*CreateCustomAgentResponseBodyDataAgenticFsMounts {
	return s.AgenticFsMounts
}

func (s *CreateCustomAgentResponseBodyData) GetArtifact() *CreateCustomAgentResponseBodyDataArtifact {
	return s.Artifact
}

func (s *CreateCustomAgentResponseBodyData) GetCreateMode() *string {
	return s.CreateMode
}

func (s *CreateCustomAgentResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateCustomAgentResponseBodyData) GetDeployType() *string {
	return s.DeployType
}

func (s *CreateCustomAgentResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *CreateCustomAgentResponseBodyData) GetEnvironment() *CreateCustomAgentResponseBodyDataEnvironment {
	return s.Environment
}

func (s *CreateCustomAgentResponseBodyData) GetHealthCheck() *CreateCustomAgentResponseBodyDataHealthCheck {
	return s.HealthCheck
}

func (s *CreateCustomAgentResponseBodyData) GetLatestExecutionError() *CreateCustomAgentResponseBodyDataLatestExecutionError {
	return s.LatestExecutionError
}

func (s *CreateCustomAgentResponseBodyData) GetLatestSpecVersion() *int64 {
	return s.LatestSpecVersion
}

func (s *CreateCustomAgentResponseBodyData) GetLatestVersionStatus() *string {
	return s.LatestVersionStatus
}

func (s *CreateCustomAgentResponseBodyData) GetLogging() *CreateCustomAgentResponseBodyDataLogging {
	return s.Logging
}

func (s *CreateCustomAgentResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CreateCustomAgentResponseBodyData) GetNetwork() *CreateCustomAgentResponseBodyDataNetwork {
	return s.Network
}

func (s *CreateCustomAgentResponseBodyData) GetOssMounts() []*CreateCustomAgentResponseBodyDataOssMounts {
	return s.OssMounts
}

func (s *CreateCustomAgentResponseBodyData) GetProtocol() []*CreateCustomAgentResponseBodyDataProtocol {
	return s.Protocol
}

func (s *CreateCustomAgentResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateCustomAgentResponseBodyData) GetRoleArn() *string {
	return s.RoleArn
}

func (s *CreateCustomAgentResponseBodyData) GetRuntime() *CreateCustomAgentResponseBodyDataRuntime {
	return s.Runtime
}

func (s *CreateCustomAgentResponseBodyData) GetSandboxPhaseCounts() map[string]*int64 {
	return s.SandboxPhaseCounts
}

func (s *CreateCustomAgentResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreateCustomAgentResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *CreateCustomAgentResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateCustomAgentResponseBodyData) SetAgentId(v string) *CreateCustomAgentResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetAgenticFsMounts(v []*CreateCustomAgentResponseBodyDataAgenticFsMounts) *CreateCustomAgentResponseBodyData {
	s.AgenticFsMounts = v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetArtifact(v *CreateCustomAgentResponseBodyDataArtifact) *CreateCustomAgentResponseBodyData {
	s.Artifact = v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetCreateMode(v string) *CreateCustomAgentResponseBodyData {
	s.CreateMode = &v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetCreatedAt(v string) *CreateCustomAgentResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetDeployType(v string) *CreateCustomAgentResponseBodyData {
	s.DeployType = &v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetDescription(v string) *CreateCustomAgentResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetEnvironment(v *CreateCustomAgentResponseBodyDataEnvironment) *CreateCustomAgentResponseBodyData {
	s.Environment = v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetHealthCheck(v *CreateCustomAgentResponseBodyDataHealthCheck) *CreateCustomAgentResponseBodyData {
	s.HealthCheck = v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetLatestExecutionError(v *CreateCustomAgentResponseBodyDataLatestExecutionError) *CreateCustomAgentResponseBodyData {
	s.LatestExecutionError = v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetLatestSpecVersion(v int64) *CreateCustomAgentResponseBodyData {
	s.LatestSpecVersion = &v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetLatestVersionStatus(v string) *CreateCustomAgentResponseBodyData {
	s.LatestVersionStatus = &v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetLogging(v *CreateCustomAgentResponseBodyDataLogging) *CreateCustomAgentResponseBodyData {
	s.Logging = v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetName(v string) *CreateCustomAgentResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetNetwork(v *CreateCustomAgentResponseBodyDataNetwork) *CreateCustomAgentResponseBodyData {
	s.Network = v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetOssMounts(v []*CreateCustomAgentResponseBodyDataOssMounts) *CreateCustomAgentResponseBodyData {
	s.OssMounts = v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetProtocol(v []*CreateCustomAgentResponseBodyDataProtocol) *CreateCustomAgentResponseBodyData {
	s.Protocol = v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetRegionId(v string) *CreateCustomAgentResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetRoleArn(v string) *CreateCustomAgentResponseBodyData {
	s.RoleArn = &v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetRuntime(v *CreateCustomAgentResponseBodyDataRuntime) *CreateCustomAgentResponseBodyData {
	s.Runtime = v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetSandboxPhaseCounts(v map[string]*int64) *CreateCustomAgentResponseBodyData {
	s.SandboxPhaseCounts = v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetStatus(v string) *CreateCustomAgentResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetUpdatedAt(v string) *CreateCustomAgentResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetWorkspaceId(v string) *CreateCustomAgentResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *CreateCustomAgentResponseBodyData) Validate() error {
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
	if s.LatestExecutionError != nil {
		if err := s.LatestExecutionError.Validate(); err != nil {
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

type CreateCustomAgentResponseBodyDataAgenticFsMounts struct {
	// The AgenticFS mount directory in the container, located under /mnt/agenticfs/.
	//
	// example:
	//
	// /mnt/agenticfs/data
	MountPath *string `json:"mountPath,omitempty" xml:"mountPath,omitempty"`
	// The relative directory within the mounted access point.
	//
	// example:
	//
	// workspace/data
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// Specifies whether the mount is read-only. A value of true indicates that the mount is read-only. A value of false indicates that the mount is readable and writable.
	//
	// example:
	//
	// false
	ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
	// The domain name of the AgenticFS access point used for the mount.
	//
	// example:
	//
	// ap-0123456789abcdef0.0123456789-vlm36.cn-hangzhou.nas.aliyuncs.com
	Server *string `json:"server,omitempty" xml:"server,omitempty"`
}

func (s CreateCustomAgentResponseBodyDataAgenticFsMounts) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentResponseBodyDataAgenticFsMounts) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentResponseBodyDataAgenticFsMounts) GetMountPath() *string {
	return s.MountPath
}

func (s *CreateCustomAgentResponseBodyDataAgenticFsMounts) GetPath() *string {
	return s.Path
}

func (s *CreateCustomAgentResponseBodyDataAgenticFsMounts) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *CreateCustomAgentResponseBodyDataAgenticFsMounts) GetServer() *string {
	return s.Server
}

func (s *CreateCustomAgentResponseBodyDataAgenticFsMounts) SetMountPath(v string) *CreateCustomAgentResponseBodyDataAgenticFsMounts {
	s.MountPath = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataAgenticFsMounts) SetPath(v string) *CreateCustomAgentResponseBodyDataAgenticFsMounts {
	s.Path = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataAgenticFsMounts) SetReadOnly(v bool) *CreateCustomAgentResponseBodyDataAgenticFsMounts {
	s.ReadOnly = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataAgenticFsMounts) SetServer(v string) *CreateCustomAgentResponseBodyDataAgenticFsMounts {
	s.Server = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataAgenticFsMounts) Validate() error {
	return dara.Validate(s)
}

type CreateCustomAgentResponseBodyDataArtifact struct {
	// The code package configuration field. Code package deployment is not supported.
	Code *CreateCustomAgentResponseBodyDataArtifactCode `json:"code,omitempty" xml:"code,omitempty" type:"Struct"`
	// The latest submitted ACR container image configuration.
	Container *CreateCustomAgentResponseBodyDataArtifactContainer `json:"container,omitempty" xml:"container,omitempty" type:"Struct"`
	// The TCP port listened on by the agent service.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8080
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
}

func (s CreateCustomAgentResponseBodyDataArtifact) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentResponseBodyDataArtifact) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentResponseBodyDataArtifact) GetCode() *CreateCustomAgentResponseBodyDataArtifactCode {
	return s.Code
}

func (s *CreateCustomAgentResponseBodyDataArtifact) GetContainer() *CreateCustomAgentResponseBodyDataArtifactContainer {
	return s.Container
}

func (s *CreateCustomAgentResponseBodyDataArtifact) GetPort() *int32 {
	return s.Port
}

func (s *CreateCustomAgentResponseBodyDataArtifact) SetCode(v *CreateCustomAgentResponseBodyDataArtifactCode) *CreateCustomAgentResponseBodyDataArtifact {
	s.Code = v
	return s
}

func (s *CreateCustomAgentResponseBodyDataArtifact) SetContainer(v *CreateCustomAgentResponseBodyDataArtifactContainer) *CreateCustomAgentResponseBodyDataArtifact {
	s.Container = v
	return s
}

func (s *CreateCustomAgentResponseBodyDataArtifact) SetPort(v int32) *CreateCustomAgentResponseBodyDataArtifact {
	s.Port = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataArtifact) Validate() error {
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

type CreateCustomAgentResponseBodyDataArtifactCode struct {
	// The startup entry point of the code package. Code package deployment is not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// python app.py
	Entrypoint *string `json:"entrypoint,omitempty" xml:"entrypoint,omitempty"`
	// The language runtime identifier of the code package. Code package deployment is not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// python
	LanguageRuntime *string `json:"languageRuntime,omitempty" xml:"languageRuntime,omitempty"`
	// The source URI of the code package. Code package deployment is not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://my-agent-bucket/packages/agent.zip
	SourceUri *string `json:"sourceUri,omitempty" xml:"sourceUri,omitempty"`
}

func (s CreateCustomAgentResponseBodyDataArtifactCode) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentResponseBodyDataArtifactCode) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentResponseBodyDataArtifactCode) GetEntrypoint() *string {
	return s.Entrypoint
}

func (s *CreateCustomAgentResponseBodyDataArtifactCode) GetLanguageRuntime() *string {
	return s.LanguageRuntime
}

func (s *CreateCustomAgentResponseBodyDataArtifactCode) GetSourceUri() *string {
	return s.SourceUri
}

func (s *CreateCustomAgentResponseBodyDataArtifactCode) SetEntrypoint(v string) *CreateCustomAgentResponseBodyDataArtifactCode {
	s.Entrypoint = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataArtifactCode) SetLanguageRuntime(v string) *CreateCustomAgentResponseBodyDataArtifactCode {
	s.LanguageRuntime = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataArtifactCode) SetSourceUri(v string) *CreateCustomAgentResponseBodyDataArtifactCode {
	s.SourceUri = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataArtifactCode) Validate() error {
	return dara.Validate(s)
}

type CreateCustomAgentResponseBodyDataArtifactContainer struct {
	// The custom executable program and startup parameters. If not configured, the default startup method of the image is used.
	//
	// example:
	//
	// python app.py
	Entrypoint *string `json:"entrypoint,omitempty" xml:"entrypoint,omitempty"`
	// The ID of the ACR instance to which the image repository belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-a1b2c3d4e5f6
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The namespace to which the ACR image repository belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-agent
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The name of the image repository in the ACR namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// chat-agent
	Repo *string `json:"repo,omitempty" xml:"repo,omitempty"`
	// The latest submitted image tag.
	//
	// This parameter is required.
	//
	// example:
	//
	// v1.0.0
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// The image source type. A value of acr indicates Alibaba Cloud Container Registry.
	//
	// This parameter is required.
	//
	// example:
	//
	// acr
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateCustomAgentResponseBodyDataArtifactContainer) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentResponseBodyDataArtifactContainer) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentResponseBodyDataArtifactContainer) GetEntrypoint() *string {
	return s.Entrypoint
}

func (s *CreateCustomAgentResponseBodyDataArtifactContainer) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateCustomAgentResponseBodyDataArtifactContainer) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateCustomAgentResponseBodyDataArtifactContainer) GetRepo() *string {
	return s.Repo
}

func (s *CreateCustomAgentResponseBodyDataArtifactContainer) GetTag() *string {
	return s.Tag
}

func (s *CreateCustomAgentResponseBodyDataArtifactContainer) GetType() *string {
	return s.Type
}

func (s *CreateCustomAgentResponseBodyDataArtifactContainer) SetEntrypoint(v string) *CreateCustomAgentResponseBodyDataArtifactContainer {
	s.Entrypoint = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataArtifactContainer) SetInstanceId(v string) *CreateCustomAgentResponseBodyDataArtifactContainer {
	s.InstanceId = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataArtifactContainer) SetNamespace(v string) *CreateCustomAgentResponseBodyDataArtifactContainer {
	s.Namespace = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataArtifactContainer) SetRepo(v string) *CreateCustomAgentResponseBodyDataArtifactContainer {
	s.Repo = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataArtifactContainer) SetTag(v string) *CreateCustomAgentResponseBodyDataArtifactContainer {
	s.Tag = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataArtifactContainer) SetType(v string) *CreateCustomAgentResponseBodyDataArtifactContainer {
	s.Type = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataArtifactContainer) Validate() error {
	return dara.Validate(s)
}

type CreateCustomAgentResponseBodyDataEnvironment struct {
	// The latest submitted list of environment variables.
	Variables []*CreateCustomAgentResponseBodyDataEnvironmentVariables `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
}

func (s CreateCustomAgentResponseBodyDataEnvironment) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentResponseBodyDataEnvironment) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentResponseBodyDataEnvironment) GetVariables() []*CreateCustomAgentResponseBodyDataEnvironmentVariables {
	return s.Variables
}

func (s *CreateCustomAgentResponseBodyDataEnvironment) SetVariables(v []*CreateCustomAgentResponseBodyDataEnvironmentVariables) *CreateCustomAgentResponseBodyDataEnvironment {
	s.Variables = v
	return s
}

func (s *CreateCustomAgentResponseBodyDataEnvironment) Validate() error {
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

type CreateCustomAgentResponseBodyDataEnvironmentVariables struct {
	// The name of the environment variable.
	//
	// This parameter is required.
	//
	// example:
	//
	// LOG_LEVEL
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The string value of the environment variable. Empty strings are allowed.
	//
	// This parameter is required.
	//
	// example:
	//
	// info
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateCustomAgentResponseBodyDataEnvironmentVariables) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentResponseBodyDataEnvironmentVariables) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentResponseBodyDataEnvironmentVariables) GetName() *string {
	return s.Name
}

func (s *CreateCustomAgentResponseBodyDataEnvironmentVariables) GetValue() *string {
	return s.Value
}

func (s *CreateCustomAgentResponseBodyDataEnvironmentVariables) SetName(v string) *CreateCustomAgentResponseBodyDataEnvironmentVariables {
	s.Name = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataEnvironmentVariables) SetValue(v string) *CreateCustomAgentResponseBodyDataEnvironmentVariables {
	s.Value = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataEnvironmentVariables) Validate() error {
	return dara.Validate(s)
}

type CreateCustomAgentResponseBodyDataHealthCheck struct {
	// This parameter does not take effect.
	//
	// example:
	//
	// 1
	HealthyThreshold *int32 `json:"healthyThreshold,omitempty" xml:"healthyThreshold,omitempty"`
	// The interval between two consecutive health checks, in seconds.
	//
	// example:
	//
	// 10
	IntervalSeconds *int32 `json:"intervalSeconds,omitempty" xml:"intervalSeconds,omitempty"`
	// The HTTP Health Check Path.
	//
	// example:
	//
	// /healthz
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// The port accessed for the HTTP health check.
	//
	// example:
	//
	// 8080
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// The timeout period for a single health check, in seconds.
	//
	// example:
	//
	// 2
	TimeoutSeconds *int32 `json:"timeoutSeconds,omitempty" xml:"timeoutSeconds,omitempty"`
	// The number of consecutive health check failures required to mark the agent as unhealthy.
	//
	// example:
	//
	// 3
	UnhealthyThreshold *int32 `json:"unhealthyThreshold,omitempty" xml:"unhealthyThreshold,omitempty"`
}

func (s CreateCustomAgentResponseBodyDataHealthCheck) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentResponseBodyDataHealthCheck) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentResponseBodyDataHealthCheck) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *CreateCustomAgentResponseBodyDataHealthCheck) GetIntervalSeconds() *int32 {
	return s.IntervalSeconds
}

func (s *CreateCustomAgentResponseBodyDataHealthCheck) GetPath() *string {
	return s.Path
}

func (s *CreateCustomAgentResponseBodyDataHealthCheck) GetPort() *int32 {
	return s.Port
}

func (s *CreateCustomAgentResponseBodyDataHealthCheck) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *CreateCustomAgentResponseBodyDataHealthCheck) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *CreateCustomAgentResponseBodyDataHealthCheck) SetHealthyThreshold(v int32) *CreateCustomAgentResponseBodyDataHealthCheck {
	s.HealthyThreshold = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataHealthCheck) SetIntervalSeconds(v int32) *CreateCustomAgentResponseBodyDataHealthCheck {
	s.IntervalSeconds = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataHealthCheck) SetPath(v string) *CreateCustomAgentResponseBodyDataHealthCheck {
	s.Path = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataHealthCheck) SetPort(v int32) *CreateCustomAgentResponseBodyDataHealthCheck {
	s.Port = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataHealthCheck) SetTimeoutSeconds(v int32) *CreateCustomAgentResponseBodyDataHealthCheck {
	s.TimeoutSeconds = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataHealthCheck) SetUnhealthyThreshold(v int32) *CreateCustomAgentResponseBodyDataHealthCheck {
	s.UnhealthyThreshold = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataHealthCheck) Validate() error {
	return dara.Validate(s)
}

type CreateCustomAgentResponseBodyDataLatestExecutionError struct {
	// The reason code for the waiting or execution failure of the latest configuration. This code is used to determine the subsequent handling method.
	//
	// example:
	//
	// AgentExecutionFailed
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The description of the reason why the latest configuration is waiting or has failed to execute. Refer to the code and retryable fields to determine what to do next.
	//
	// example:
	//
	// The agent execution failed. Check the configuration and resubmit
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Specifies whether the operation is automatically retried. A value of true indicates that you can query again later. A value of false indicates that you must check and fix the configuration based on the reason, and then resubmit the request.
	Retryable *bool `json:"retryable,omitempty" xml:"retryable,omitempty"`
}

func (s CreateCustomAgentResponseBodyDataLatestExecutionError) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentResponseBodyDataLatestExecutionError) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentResponseBodyDataLatestExecutionError) GetCode() *string {
	return s.Code
}

func (s *CreateCustomAgentResponseBodyDataLatestExecutionError) GetMessage() *string {
	return s.Message
}

func (s *CreateCustomAgentResponseBodyDataLatestExecutionError) GetRetryable() *bool {
	return s.Retryable
}

func (s *CreateCustomAgentResponseBodyDataLatestExecutionError) SetCode(v string) *CreateCustomAgentResponseBodyDataLatestExecutionError {
	s.Code = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataLatestExecutionError) SetMessage(v string) *CreateCustomAgentResponseBodyDataLatestExecutionError {
	s.Message = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataLatestExecutionError) SetRetryable(v bool) *CreateCustomAgentResponseBodyDataLatestExecutionError {
	s.Retryable = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataLatestExecutionError) Validate() error {
	return dara.Validate(s)
}

type CreateCustomAgentResponseBodyDataLogging struct {
	// The Simple Log Service (SLS) log delivery destination.
	Sls *CreateCustomAgentResponseBodyDataLoggingSls `json:"sls,omitempty" xml:"sls,omitempty" type:"Struct"`
}

func (s CreateCustomAgentResponseBodyDataLogging) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentResponseBodyDataLogging) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentResponseBodyDataLogging) GetSls() *CreateCustomAgentResponseBodyDataLoggingSls {
	return s.Sls
}

func (s *CreateCustomAgentResponseBodyDataLogging) SetSls(v *CreateCustomAgentResponseBodyDataLoggingSls) *CreateCustomAgentResponseBodyDataLogging {
	s.Sls = v
	return s
}

func (s *CreateCustomAgentResponseBodyDataLogging) Validate() error {
	if s.Sls != nil {
		if err := s.Sls.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCustomAgentResponseBodyDataLoggingSls struct {
	// The name of the Simple Log Service Logstore that receives agent logs.
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

func (s CreateCustomAgentResponseBodyDataLoggingSls) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentResponseBodyDataLoggingSls) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentResponseBodyDataLoggingSls) GetLogstore() *string {
	return s.Logstore
}

func (s *CreateCustomAgentResponseBodyDataLoggingSls) GetProject() *string {
	return s.Project
}

func (s *CreateCustomAgentResponseBodyDataLoggingSls) SetLogstore(v string) *CreateCustomAgentResponseBodyDataLoggingSls {
	s.Logstore = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataLoggingSls) SetProject(v string) *CreateCustomAgentResponseBodyDataLoggingSls {
	s.Project = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataLoggingSls) Validate() error {
	return dara.Validate(s)
}

type CreateCustomAgentResponseBodyDataNetwork struct {
	// The Internet access configuration of the agent. If this field is not configured, Internet access is allowed on the agent side, but it is still restricted by the workspace.
	AccessInternet *CreateCustomAgentResponseBodyDataNetworkAccessInternet `json:"accessInternet,omitempty" xml:"accessInternet,omitempty" type:"Struct"`
	// The user VPC access configuration of the agent. If this field is not configured, user VPC access is allowed on the agent side, but it is still restricted by the workspace.
	AccessVpc *CreateCustomAgentResponseBodyDataNetworkAccessVpc `json:"accessVpc,omitempty" xml:"accessVpc,omitempty" type:"Struct"`
}

func (s CreateCustomAgentResponseBodyDataNetwork) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentResponseBodyDataNetwork) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentResponseBodyDataNetwork) GetAccessInternet() *CreateCustomAgentResponseBodyDataNetworkAccessInternet {
	return s.AccessInternet
}

func (s *CreateCustomAgentResponseBodyDataNetwork) GetAccessVpc() *CreateCustomAgentResponseBodyDataNetworkAccessVpc {
	return s.AccessVpc
}

func (s *CreateCustomAgentResponseBodyDataNetwork) SetAccessInternet(v *CreateCustomAgentResponseBodyDataNetworkAccessInternet) *CreateCustomAgentResponseBodyDataNetwork {
	s.AccessInternet = v
	return s
}

func (s *CreateCustomAgentResponseBodyDataNetwork) SetAccessVpc(v *CreateCustomAgentResponseBodyDataNetworkAccessVpc) *CreateCustomAgentResponseBodyDataNetwork {
	s.AccessVpc = v
	return s
}

func (s *CreateCustomAgentResponseBodyDataNetwork) Validate() error {
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

type CreateCustomAgentResponseBodyDataNetworkAccessInternet struct {
	// Specifies whether Internet access is allowed on the agent side. The actual access is still restricted by the workspace.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s CreateCustomAgentResponseBodyDataNetworkAccessInternet) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentResponseBodyDataNetworkAccessInternet) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentResponseBodyDataNetworkAccessInternet) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateCustomAgentResponseBodyDataNetworkAccessInternet) SetEnabled(v bool) *CreateCustomAgentResponseBodyDataNetworkAccessInternet {
	s.Enabled = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataNetworkAccessInternet) Validate() error {
	return dara.Validate(s)
}

type CreateCustomAgentResponseBodyDataNetworkAccessVpc struct {
	// Specifies whether access to the user VPC configured for the workspace is allowed on the agent side. The actual access is still restricted by the workspace.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s CreateCustomAgentResponseBodyDataNetworkAccessVpc) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentResponseBodyDataNetworkAccessVpc) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentResponseBodyDataNetworkAccessVpc) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateCustomAgentResponseBodyDataNetworkAccessVpc) SetEnabled(v bool) *CreateCustomAgentResponseBodyDataNetworkAccessVpc {
	s.Enabled = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataNetworkAccessVpc) Validate() error {
	return dara.Validate(s)
}

type CreateCustomAgentResponseBodyDataOssMounts struct {
	// The name of the mounted OSS bucket.
	//
	// example:
	//
	// my-agent-data
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// The OSS mount directory in the container, located under /mnt/oss/.
	//
	// example:
	//
	// /mnt/oss/train
	MountPath *string `json:"mountPath,omitempty" xml:"mountPath,omitempty"`
	// The prefix of the object directory in the mounted bucket.
	//
	// example:
	//
	// datasets/train
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// Specifies whether the mount is read-only. A value of true indicates that the mount is read-only. A value of false indicates that the mount is readable and writable.
	ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
}

func (s CreateCustomAgentResponseBodyDataOssMounts) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentResponseBodyDataOssMounts) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentResponseBodyDataOssMounts) GetBucketName() *string {
	return s.BucketName
}

func (s *CreateCustomAgentResponseBodyDataOssMounts) GetMountPath() *string {
	return s.MountPath
}

func (s *CreateCustomAgentResponseBodyDataOssMounts) GetPath() *string {
	return s.Path
}

func (s *CreateCustomAgentResponseBodyDataOssMounts) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *CreateCustomAgentResponseBodyDataOssMounts) SetBucketName(v string) *CreateCustomAgentResponseBodyDataOssMounts {
	s.BucketName = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataOssMounts) SetMountPath(v string) *CreateCustomAgentResponseBodyDataOssMounts {
	s.MountPath = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataOssMounts) SetPath(v string) *CreateCustomAgentResponseBodyDataOssMounts {
	s.Path = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataOssMounts) SetReadOnly(v bool) *CreateCustomAgentResponseBodyDataOssMounts {
	s.ReadOnly = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataOssMounts) Validate() error {
	return dara.Validate(s)
}

type CreateCustomAgentResponseBodyDataProtocol struct {
	// The request header configurations for the current protocol. The key is the header name and the value is a string.
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	// The list of API paths corresponding to the current protocol.
	Paths []*CreateCustomAgentResponseBodyDataProtocolPaths `json:"paths,omitempty" xml:"paths,omitempty" type:"Repeated"`
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

func (s CreateCustomAgentResponseBodyDataProtocol) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentResponseBodyDataProtocol) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentResponseBodyDataProtocol) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCustomAgentResponseBodyDataProtocol) GetPaths() []*CreateCustomAgentResponseBodyDataProtocolPaths {
	return s.Paths
}

func (s *CreateCustomAgentResponseBodyDataProtocol) GetType() *string {
	return s.Type
}

func (s *CreateCustomAgentResponseBodyDataProtocol) SetHeaders(v map[string]*string) *CreateCustomAgentResponseBodyDataProtocol {
	s.Headers = v
	return s
}

func (s *CreateCustomAgentResponseBodyDataProtocol) SetPaths(v []*CreateCustomAgentResponseBodyDataProtocolPaths) *CreateCustomAgentResponseBodyDataProtocol {
	s.Paths = v
	return s
}

func (s *CreateCustomAgentResponseBodyDataProtocol) SetType(v string) *CreateCustomAgentResponseBodyDataProtocol {
	s.Type = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataProtocol) Validate() error {
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

type CreateCustomAgentResponseBodyDataProtocolPaths struct {
	// The usage description of the current API path.
	//
	// example:
	//
	// AG-UI conversation API
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The API path corresponding to the current protocol.
	//
	// This parameter is required.
	//
	// example:
	//
	// /agui
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateCustomAgentResponseBodyDataProtocolPaths) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentResponseBodyDataProtocolPaths) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentResponseBodyDataProtocolPaths) GetDescription() *string {
	return s.Description
}

func (s *CreateCustomAgentResponseBodyDataProtocolPaths) GetName() *string {
	return s.Name
}

func (s *CreateCustomAgentResponseBodyDataProtocolPaths) SetDescription(v string) *CreateCustomAgentResponseBodyDataProtocolPaths {
	s.Description = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataProtocolPaths) SetName(v string) *CreateCustomAgentResponseBodyDataProtocolPaths {
	s.Name = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataProtocolPaths) Validate() error {
	return dara.Validate(s)
}

type CreateCustomAgentResponseBodyDataRuntime struct {
	// The compute resource configuration for the agent runtime instance.
	//
	// This parameter is required.
	Compute *CreateCustomAgentResponseBodyDataRuntimeCompute `json:"compute,omitempty" xml:"compute,omitempty" type:"Struct"`
	// The elastic policy configurations for the agent.
	Hpa *CreateCustomAgentResponseBodyDataRuntimeHpa `json:"hpa,omitempty" xml:"hpa,omitempty" type:"Struct"`
	// The session policy configurations for the agent.
	//
	// This parameter is required.
	SessionPolicy *CreateCustomAgentResponseBodyDataRuntimeSessionPolicy `json:"sessionPolicy,omitempty" xml:"sessionPolicy,omitempty" type:"Struct"`
}

func (s CreateCustomAgentResponseBodyDataRuntime) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentResponseBodyDataRuntime) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentResponseBodyDataRuntime) GetCompute() *CreateCustomAgentResponseBodyDataRuntimeCompute {
	return s.Compute
}

func (s *CreateCustomAgentResponseBodyDataRuntime) GetHpa() *CreateCustomAgentResponseBodyDataRuntimeHpa {
	return s.Hpa
}

func (s *CreateCustomAgentResponseBodyDataRuntime) GetSessionPolicy() *CreateCustomAgentResponseBodyDataRuntimeSessionPolicy {
	return s.SessionPolicy
}

func (s *CreateCustomAgentResponseBodyDataRuntime) SetCompute(v *CreateCustomAgentResponseBodyDataRuntimeCompute) *CreateCustomAgentResponseBodyDataRuntime {
	s.Compute = v
	return s
}

func (s *CreateCustomAgentResponseBodyDataRuntime) SetHpa(v *CreateCustomAgentResponseBodyDataRuntimeHpa) *CreateCustomAgentResponseBodyDataRuntime {
	s.Hpa = v
	return s
}

func (s *CreateCustomAgentResponseBodyDataRuntime) SetSessionPolicy(v *CreateCustomAgentResponseBodyDataRuntimeSessionPolicy) *CreateCustomAgentResponseBodyDataRuntime {
	s.SessionPolicy = v
	return s
}

func (s *CreateCustomAgentResponseBodyDataRuntime) Validate() error {
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

type CreateCustomAgentResponseBodyDataRuntimeCompute struct {
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

func (s CreateCustomAgentResponseBodyDataRuntimeCompute) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentResponseBodyDataRuntimeCompute) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentResponseBodyDataRuntimeCompute) GetComputeClass() *string {
	return s.ComputeClass
}

func (s *CreateCustomAgentResponseBodyDataRuntimeCompute) SetComputeClass(v string) *CreateCustomAgentResponseBodyDataRuntimeCompute {
	s.ComputeClass = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataRuntimeCompute) Validate() error {
	return dara.Validate(s)
}

type CreateCustomAgentResponseBodyDataRuntimeHpa struct {
	// Specifies whether elastic scaling is enabled. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// This parameter is required.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The maximum number of concurrent sessions configured for a single runtime instance.
	//
	// example:
	//
	// 100
	MaxConcurrentSessionsPerSandbox *int32 `json:"maxConcurrentSessionsPerSandbox,omitempty" xml:"maxConcurrentSessionsPerSandbox,omitempty"`
	// The maximum number of runtime instances configured.
	//
	// example:
	//
	// 10
	MaxSandboxCount *int32 `json:"maxSandboxCount,omitempty" xml:"maxSandboxCount,omitempty"`
	// The minimum number of runtime instances configured.
	//
	// example:
	//
	// 1
	MinSandboxCount *int32 `json:"minSandboxCount,omitempty" xml:"minSandboxCount,omitempty"`
	// The maximum idle time configured for a session, in seconds.
	//
	// example:
	//
	// 600
	SessionTtlSeconds *int32 `json:"sessionTtlSeconds,omitempty" xml:"sessionTtlSeconds,omitempty"`
}

func (s CreateCustomAgentResponseBodyDataRuntimeHpa) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentResponseBodyDataRuntimeHpa) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentResponseBodyDataRuntimeHpa) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateCustomAgentResponseBodyDataRuntimeHpa) GetMaxConcurrentSessionsPerSandbox() *int32 {
	return s.MaxConcurrentSessionsPerSandbox
}

func (s *CreateCustomAgentResponseBodyDataRuntimeHpa) GetMaxSandboxCount() *int32 {
	return s.MaxSandboxCount
}

func (s *CreateCustomAgentResponseBodyDataRuntimeHpa) GetMinSandboxCount() *int32 {
	return s.MinSandboxCount
}

func (s *CreateCustomAgentResponseBodyDataRuntimeHpa) GetSessionTtlSeconds() *int32 {
	return s.SessionTtlSeconds
}

func (s *CreateCustomAgentResponseBodyDataRuntimeHpa) SetEnabled(v bool) *CreateCustomAgentResponseBodyDataRuntimeHpa {
	s.Enabled = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataRuntimeHpa) SetMaxConcurrentSessionsPerSandbox(v int32) *CreateCustomAgentResponseBodyDataRuntimeHpa {
	s.MaxConcurrentSessionsPerSandbox = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataRuntimeHpa) SetMaxSandboxCount(v int32) *CreateCustomAgentResponseBodyDataRuntimeHpa {
	s.MaxSandboxCount = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataRuntimeHpa) SetMinSandboxCount(v int32) *CreateCustomAgentResponseBodyDataRuntimeHpa {
	s.MinSandboxCount = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataRuntimeHpa) SetSessionTtlSeconds(v int32) *CreateCustomAgentResponseBodyDataRuntimeHpa {
	s.SessionTtlSeconds = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataRuntimeHpa) Validate() error {
	return dara.Validate(s)
}

type CreateCustomAgentResponseBodyDataRuntimeSessionPolicy struct {
	// The name of the HTTP request header used by the session policy. This parameter applies to ISOLATED_HEADER_FIELD.
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
	// Note: The compatibility value DISABLED in a request is returned as AFFINITY.
	//
	// This parameter is required.
	//
	// example:
	//
	// AFFINITY
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateCustomAgentResponseBodyDataRuntimeSessionPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentResponseBodyDataRuntimeSessionPolicy) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentResponseBodyDataRuntimeSessionPolicy) GetHeaderName() *string {
	return s.HeaderName
}

func (s *CreateCustomAgentResponseBodyDataRuntimeSessionPolicy) GetType() *string {
	return s.Type
}

func (s *CreateCustomAgentResponseBodyDataRuntimeSessionPolicy) SetHeaderName(v string) *CreateCustomAgentResponseBodyDataRuntimeSessionPolicy {
	s.HeaderName = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataRuntimeSessionPolicy) SetType(v string) *CreateCustomAgentResponseBodyDataRuntimeSessionPolicy {
	s.Type = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataRuntimeSessionPolicy) Validate() error {
	return dara.Validate(s)
}
