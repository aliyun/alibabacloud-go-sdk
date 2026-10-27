// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateCustomAgentResponseBody
	GetCode() *string
	SetData(v *UpdateCustomAgentResponseBodyData) *UpdateCustomAgentResponseBody
	GetData() *UpdateCustomAgentResponseBodyData
	SetHttpStatusCode(v int32) *UpdateCustomAgentResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateCustomAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateCustomAgentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateCustomAgentResponseBody
	GetSuccess() *bool
}

type UpdateCustomAgentResponseBody struct {
	// The request result code. The value is SUCCESS if the request is successful.
	//
	// This parameter is required.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The latest submitted configuration and status of the custom agent. A successful response indicates that the request is accepted, but the configuration may not have taken effect yet.
	//
	// This parameter is required.
	Data *UpdateCustomAgentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code of the request. A value of 200 indicates success.
	//
	// This parameter is required.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The message that describes the processing result of the request. The value is success if the request is successful.
	//
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The unique identifier of the request, which is used to locate and troubleshoot issues.
	//
	// This parameter is required.
	//
	// example:
	//
	// 550E8400-E29B-41D4-A716-446655440000
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful. A value of true indicates success. A successful creation, update, or deletion does not mean that the background operation is complete.
	//
	// This parameter is required.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateCustomAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateCustomAgentResponseBody) GetData() *UpdateCustomAgentResponseBodyData {
	return s.Data
}

func (s *UpdateCustomAgentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateCustomAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateCustomAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCustomAgentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateCustomAgentResponseBody) SetCode(v string) *UpdateCustomAgentResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateCustomAgentResponseBody) SetData(v *UpdateCustomAgentResponseBodyData) *UpdateCustomAgentResponseBody {
	s.Data = v
	return s
}

func (s *UpdateCustomAgentResponseBody) SetHttpStatusCode(v int32) *UpdateCustomAgentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateCustomAgentResponseBody) SetMessage(v string) *UpdateCustomAgentResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateCustomAgentResponseBody) SetRequestId(v string) *UpdateCustomAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCustomAgentResponseBody) SetSuccess(v bool) *UpdateCustomAgentResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateCustomAgentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateCustomAgentResponseBodyData struct {
	// The unique identifier of the custom agent. You can obtain this identifier from the response of the CreateCustomAgent or ListCustomAgents operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// agent-550e8400e29b41d4a716446655440000
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The latest submitted list of additional AgenticFS mounts.
	AgenticFsMounts []*UpdateCustomAgentResponseBodyDataAgenticFsMounts `json:"agenticFsMounts,omitempty" xml:"agenticFsMounts,omitempty" type:"Repeated"`
	// The latest submitted deployment artifact configuration, including the service listening port and the ACR container image. To check whether the configuration has taken effect, refer to latestVersionStatus.
	Artifact *UpdateCustomAgentResponseBodyDataArtifact `json:"artifact,omitempty" xml:"artifact,omitempty" type:"Struct"`
	// The currently saved code package information.
	CodePackageInfo *UpdateCustomAgentResponseBodyDataCodePackageInfo `json:"codePackageInfo,omitempty" xml:"codePackageInfo,omitempty" type:"Struct"`
	// The creation method of the agent. A value of CUSTOM is returned for custom agents.
	//
	// example:
	//
	// CUSTOM
	CreateMode *string `json:"createMode,omitempty" xml:"createMode,omitempty"`
	// The time when the agent was created. The time is in ISO 8601 format. A trailing Z indicates the UTC time zone. Fractional seconds may be included.
	//
	// example:
	//
	// 2026-10-10T06:00:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The deployment method of the agent. Custom agents are deployed and managed by the platform. A value of MANAGED is returned.
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
	// The effective configuration version.
	//
	// This parameter is required.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 1
	EffectiveSpecVersion *int64 `json:"effectiveSpecVersion,omitempty" xml:"effectiveSpecVersion,omitempty"`
	// The environment variable configuration for the agent process.
	Environment *UpdateCustomAgentResponseBodyDataEnvironment `json:"environment,omitempty" xml:"environment,omitempty" type:"Struct"`
	// The HTTP health check configuration of the agent. This parameter is empty if no health check is configured.
	HealthCheck *UpdateCustomAgentResponseBodyDataHealthCheck `json:"healthCheck,omitempty" xml:"healthCheck,omitempty" type:"Struct"`
	// The ID of the currently associated image cache.
	//
	// This parameter is required.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// img-example
	ImageCacheId *string `json:"imageCacheId,omitempty" xml:"imageCacheId,omitempty"`
	// The cache preparation acceptance summary for this request.
	ImageCachePreparation *UpdateCustomAgentResponseBodyDataImageCachePreparation `json:"imageCachePreparation,omitempty" xml:"imageCachePreparation,omitempty" type:"Struct"`
	// The waiting or execution error of the latest configuration version. This field is returned only when the version is waiting for a retry or fails to execute.
	LatestExecutionError *UpdateCustomAgentResponseBodyDataLatestExecutionError `json:"latestExecutionError,omitempty" xml:"latestExecutionError,omitempty" type:"Struct"`
	// The current latest configuration version number. Updates to the same configuration may retain the original version number. This field does not indicate that the configuration has taken effect.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	LatestSpecVersion *int64 `json:"latestSpecVersion,omitempty" xml:"latestSpecVersion,omitempty"`
	// The execution status corresponding to latestSpecVersion. Valid values:
	//
	// - pending: waiting for execution
	//
	// - processing: executing
	//
	// - waiting_retry: waiting for retry
	//
	// - succeeded: executed successfully
	//
	// - failed: execution failed
	//
	// - superseded: superseded by a subsequent version
	//
	// if can be null:
	// true
	//
	// example:
	//
	// succeeded
	LatestVersionStatus *string `json:"latestVersionStatus,omitempty" xml:"latestVersionStatus,omitempty"`
	// The latest submitted log delivery configuration.
	Logging *UpdateCustomAgentResponseBodyDataLogging `json:"logging,omitempty" xml:"logging,omitempty" type:"Struct"`
	// The custom agent name.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-custom-agent
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The network access configuration of the agent. The actual access scope is also limited by the network configuration of the workspace.
	Network *UpdateCustomAgentResponseBodyDataNetwork `json:"network,omitempty" xml:"network,omitempty" type:"Struct"`
	// The high-code agent observability configuration.
	Observability *UpdateCustomAgentResponseBodyDataObservability `json:"observability,omitempty" xml:"observability,omitempty" type:"Struct"`
	// The latest submitted list of OSS directory mounts.
	OssMounts []*UpdateCustomAgentResponseBodyDataOssMounts `json:"ossMounts,omitempty" xml:"ossMounts,omitempty" type:"Repeated"`
	// The list of invocation protocol configurations for the agent.
	Protocol []*UpdateCustomAgentResponseBodyDataProtocol `json:"protocol,omitempty" xml:"protocol,omitempty" type:"Repeated"`
	// The region ID of the workspace to which the agent belongs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The ARN of the RAM role used when the agent accesses user cloud resources.
	//
	// example:
	//
	// acs:ram::1234567890123456:role/agentcore-runtime-role
	RoleArn *string `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	// The latest committed runtime configuration, including compute specifications, session policy, and elastic policy.
	Runtime *UpdateCustomAgentResponseBodyDataRuntime `json:"runtime,omitempty" xml:"runtime,omitempty" type:"Struct"`
	// The instance counts grouped by sandbox phase. Keys include PENDING (creating or initializing), RUNNING (running), HIBERNATING (hibernating), HIBERNATED (hibernated), RESUMING (resuming), TERMINATING (terminating), and FAILED (failed). Only phases that actually occur are returned. Missing keys are treated as 0.
	SandboxPhaseCounts map[string]*int64 `json:"sandboxPhaseCounts,omitempty" xml:"sandboxPhaseCounts,omitempty"`
	// The overall status of the agent. Valid values: Creating: The agent is being created. Running: The agent is running. Failed: The agent failed. Updating: The agent is being updated. Deleting: The agent is being deleted. Deleted: The agent is deleted. To check the execution progress of the latest configuration, query latestVersionStatus.
	//
	// This parameter is required.
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
	// Specifies whether the image cache is enabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	UseImageCache *bool `json:"useImageCache,omitempty" xml:"useImageCache,omitempty"`
	// The ID of the workspace to which the custom agent belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ws-a1b2c3d4e5f6
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s UpdateCustomAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentResponseBodyData) GetAgentId() *string {
	return s.AgentId
}

func (s *UpdateCustomAgentResponseBodyData) GetAgenticFsMounts() []*UpdateCustomAgentResponseBodyDataAgenticFsMounts {
	return s.AgenticFsMounts
}

func (s *UpdateCustomAgentResponseBodyData) GetArtifact() *UpdateCustomAgentResponseBodyDataArtifact {
	return s.Artifact
}

func (s *UpdateCustomAgentResponseBodyData) GetCodePackageInfo() *UpdateCustomAgentResponseBodyDataCodePackageInfo {
	return s.CodePackageInfo
}

func (s *UpdateCustomAgentResponseBodyData) GetCreateMode() *string {
	return s.CreateMode
}

func (s *UpdateCustomAgentResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *UpdateCustomAgentResponseBodyData) GetDeployType() *string {
	return s.DeployType
}

func (s *UpdateCustomAgentResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *UpdateCustomAgentResponseBodyData) GetEffectiveSpecVersion() *int64 {
	return s.EffectiveSpecVersion
}

func (s *UpdateCustomAgentResponseBodyData) GetEnvironment() *UpdateCustomAgentResponseBodyDataEnvironment {
	return s.Environment
}

func (s *UpdateCustomAgentResponseBodyData) GetHealthCheck() *UpdateCustomAgentResponseBodyDataHealthCheck {
	return s.HealthCheck
}

func (s *UpdateCustomAgentResponseBodyData) GetImageCacheId() *string {
	return s.ImageCacheId
}

func (s *UpdateCustomAgentResponseBodyData) GetImageCachePreparation() *UpdateCustomAgentResponseBodyDataImageCachePreparation {
	return s.ImageCachePreparation
}

func (s *UpdateCustomAgentResponseBodyData) GetLatestExecutionError() *UpdateCustomAgentResponseBodyDataLatestExecutionError {
	return s.LatestExecutionError
}

func (s *UpdateCustomAgentResponseBodyData) GetLatestSpecVersion() *int64 {
	return s.LatestSpecVersion
}

func (s *UpdateCustomAgentResponseBodyData) GetLatestVersionStatus() *string {
	return s.LatestVersionStatus
}

func (s *UpdateCustomAgentResponseBodyData) GetLogging() *UpdateCustomAgentResponseBodyDataLogging {
	return s.Logging
}

func (s *UpdateCustomAgentResponseBodyData) GetName() *string {
	return s.Name
}

func (s *UpdateCustomAgentResponseBodyData) GetNetwork() *UpdateCustomAgentResponseBodyDataNetwork {
	return s.Network
}

func (s *UpdateCustomAgentResponseBodyData) GetObservability() *UpdateCustomAgentResponseBodyDataObservability {
	return s.Observability
}

func (s *UpdateCustomAgentResponseBodyData) GetOssMounts() []*UpdateCustomAgentResponseBodyDataOssMounts {
	return s.OssMounts
}

func (s *UpdateCustomAgentResponseBodyData) GetProtocol() []*UpdateCustomAgentResponseBodyDataProtocol {
	return s.Protocol
}

func (s *UpdateCustomAgentResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateCustomAgentResponseBodyData) GetRoleArn() *string {
	return s.RoleArn
}

func (s *UpdateCustomAgentResponseBodyData) GetRuntime() *UpdateCustomAgentResponseBodyDataRuntime {
	return s.Runtime
}

func (s *UpdateCustomAgentResponseBodyData) GetSandboxPhaseCounts() map[string]*int64 {
	return s.SandboxPhaseCounts
}

func (s *UpdateCustomAgentResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *UpdateCustomAgentResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *UpdateCustomAgentResponseBodyData) GetUseImageCache() *bool {
	return s.UseImageCache
}

func (s *UpdateCustomAgentResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateCustomAgentResponseBodyData) SetAgentId(v string) *UpdateCustomAgentResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyData) SetAgenticFsMounts(v []*UpdateCustomAgentResponseBodyDataAgenticFsMounts) *UpdateCustomAgentResponseBodyData {
	s.AgenticFsMounts = v
	return s
}

func (s *UpdateCustomAgentResponseBodyData) SetArtifact(v *UpdateCustomAgentResponseBodyDataArtifact) *UpdateCustomAgentResponseBodyData {
	s.Artifact = v
	return s
}

func (s *UpdateCustomAgentResponseBodyData) SetCodePackageInfo(v *UpdateCustomAgentResponseBodyDataCodePackageInfo) *UpdateCustomAgentResponseBodyData {
	s.CodePackageInfo = v
	return s
}

func (s *UpdateCustomAgentResponseBodyData) SetCreateMode(v string) *UpdateCustomAgentResponseBodyData {
	s.CreateMode = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyData) SetCreatedAt(v string) *UpdateCustomAgentResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyData) SetDeployType(v string) *UpdateCustomAgentResponseBodyData {
	s.DeployType = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyData) SetDescription(v string) *UpdateCustomAgentResponseBodyData {
	s.Description = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyData) SetEffectiveSpecVersion(v int64) *UpdateCustomAgentResponseBodyData {
	s.EffectiveSpecVersion = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyData) SetEnvironment(v *UpdateCustomAgentResponseBodyDataEnvironment) *UpdateCustomAgentResponseBodyData {
	s.Environment = v
	return s
}

func (s *UpdateCustomAgentResponseBodyData) SetHealthCheck(v *UpdateCustomAgentResponseBodyDataHealthCheck) *UpdateCustomAgentResponseBodyData {
	s.HealthCheck = v
	return s
}

func (s *UpdateCustomAgentResponseBodyData) SetImageCacheId(v string) *UpdateCustomAgentResponseBodyData {
	s.ImageCacheId = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyData) SetImageCachePreparation(v *UpdateCustomAgentResponseBodyDataImageCachePreparation) *UpdateCustomAgentResponseBodyData {
	s.ImageCachePreparation = v
	return s
}

func (s *UpdateCustomAgentResponseBodyData) SetLatestExecutionError(v *UpdateCustomAgentResponseBodyDataLatestExecutionError) *UpdateCustomAgentResponseBodyData {
	s.LatestExecutionError = v
	return s
}

func (s *UpdateCustomAgentResponseBodyData) SetLatestSpecVersion(v int64) *UpdateCustomAgentResponseBodyData {
	s.LatestSpecVersion = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyData) SetLatestVersionStatus(v string) *UpdateCustomAgentResponseBodyData {
	s.LatestVersionStatus = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyData) SetLogging(v *UpdateCustomAgentResponseBodyDataLogging) *UpdateCustomAgentResponseBodyData {
	s.Logging = v
	return s
}

func (s *UpdateCustomAgentResponseBodyData) SetName(v string) *UpdateCustomAgentResponseBodyData {
	s.Name = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyData) SetNetwork(v *UpdateCustomAgentResponseBodyDataNetwork) *UpdateCustomAgentResponseBodyData {
	s.Network = v
	return s
}

func (s *UpdateCustomAgentResponseBodyData) SetObservability(v *UpdateCustomAgentResponseBodyDataObservability) *UpdateCustomAgentResponseBodyData {
	s.Observability = v
	return s
}

func (s *UpdateCustomAgentResponseBodyData) SetOssMounts(v []*UpdateCustomAgentResponseBodyDataOssMounts) *UpdateCustomAgentResponseBodyData {
	s.OssMounts = v
	return s
}

func (s *UpdateCustomAgentResponseBodyData) SetProtocol(v []*UpdateCustomAgentResponseBodyDataProtocol) *UpdateCustomAgentResponseBodyData {
	s.Protocol = v
	return s
}

func (s *UpdateCustomAgentResponseBodyData) SetRegionId(v string) *UpdateCustomAgentResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyData) SetRoleArn(v string) *UpdateCustomAgentResponseBodyData {
	s.RoleArn = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyData) SetRuntime(v *UpdateCustomAgentResponseBodyDataRuntime) *UpdateCustomAgentResponseBodyData {
	s.Runtime = v
	return s
}

func (s *UpdateCustomAgentResponseBodyData) SetSandboxPhaseCounts(v map[string]*int64) *UpdateCustomAgentResponseBodyData {
	s.SandboxPhaseCounts = v
	return s
}

func (s *UpdateCustomAgentResponseBodyData) SetStatus(v string) *UpdateCustomAgentResponseBodyData {
	s.Status = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyData) SetUpdatedAt(v string) *UpdateCustomAgentResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyData) SetUseImageCache(v bool) *UpdateCustomAgentResponseBodyData {
	s.UseImageCache = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyData) SetWorkspaceId(v string) *UpdateCustomAgentResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyData) Validate() error {
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
	if s.CodePackageInfo != nil {
		if err := s.CodePackageInfo.Validate(); err != nil {
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
	if s.ImageCachePreparation != nil {
		if err := s.ImageCachePreparation.Validate(); err != nil {
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
	if s.Observability != nil {
		if err := s.Observability.Validate(); err != nil {
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

type UpdateCustomAgentResponseBodyDataAgenticFsMounts struct {
	// The AgenticFS mount directory in the container, located under /mnt/agenticfs/.
	//
	// example:
	//
	// /mnt/agenticfs/data
	MountPath *string `json:"mountPath,omitempty" xml:"mountPath,omitempty"`
	// The relative directory in the mounted access point.
	//
	// example:
	//
	// workspace/data
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// Specifies whether the mount is read-only. A value of true indicates a read-only mount. A value of false indicates a read-write mount.
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

func (s UpdateCustomAgentResponseBodyDataAgenticFsMounts) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentResponseBodyDataAgenticFsMounts) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentResponseBodyDataAgenticFsMounts) GetMountPath() *string {
	return s.MountPath
}

func (s *UpdateCustomAgentResponseBodyDataAgenticFsMounts) GetPath() *string {
	return s.Path
}

func (s *UpdateCustomAgentResponseBodyDataAgenticFsMounts) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *UpdateCustomAgentResponseBodyDataAgenticFsMounts) GetServer() *string {
	return s.Server
}

func (s *UpdateCustomAgentResponseBodyDataAgenticFsMounts) SetMountPath(v string) *UpdateCustomAgentResponseBodyDataAgenticFsMounts {
	s.MountPath = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataAgenticFsMounts) SetPath(v string) *UpdateCustomAgentResponseBodyDataAgenticFsMounts {
	s.Path = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataAgenticFsMounts) SetReadOnly(v bool) *UpdateCustomAgentResponseBodyDataAgenticFsMounts {
	s.ReadOnly = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataAgenticFsMounts) SetServer(v string) *UpdateCustomAgentResponseBodyDataAgenticFsMounts {
	s.Server = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataAgenticFsMounts) Validate() error {
	return dara.Validate(s)
}

type UpdateCustomAgentResponseBodyDataArtifact struct {
	// The code package deployment configuration. This field is mutually exclusive with the container field. Code packages use the ZIP or TAR.GZ format. Dependencies must be packaged by the user. The platform does not install dependencies or perform online builds.
	Code *UpdateCustomAgentResponseBodyDataArtifactCode `json:"code,omitempty" xml:"code,omitempty" type:"Struct"`
	// The latest submitted ACR container image configuration.
	Container *UpdateCustomAgentResponseBodyDataArtifactContainer `json:"container,omitempty" xml:"container,omitempty" type:"Struct"`
	// The TCP port listened on by the agent service.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8080
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
}

func (s UpdateCustomAgentResponseBodyDataArtifact) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentResponseBodyDataArtifact) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentResponseBodyDataArtifact) GetCode() *UpdateCustomAgentResponseBodyDataArtifactCode {
	return s.Code
}

func (s *UpdateCustomAgentResponseBodyDataArtifact) GetContainer() *UpdateCustomAgentResponseBodyDataArtifactContainer {
	return s.Container
}

func (s *UpdateCustomAgentResponseBodyDataArtifact) GetPort() *int32 {
	return s.Port
}

func (s *UpdateCustomAgentResponseBodyDataArtifact) SetCode(v *UpdateCustomAgentResponseBodyDataArtifactCode) *UpdateCustomAgentResponseBodyDataArtifact {
	s.Code = v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataArtifact) SetContainer(v *UpdateCustomAgentResponseBodyDataArtifactContainer) *UpdateCustomAgentResponseBodyDataArtifact {
	s.Container = v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataArtifact) SetPort(v int32) *UpdateCustomAgentResponseBodyDataArtifact {
	s.Port = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataArtifact) Validate() error {
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

type UpdateCustomAgentResponseBodyDataArtifactCode struct {
	// The startup command for the code package. This field is required in code package mode. The working directory is the extracted code root directory /code, such as python main.py. Quotes and escapes are supported. Environment variables, pipes, and logical operations are not automatically expanded. For complex commands, use scripts within the package.
	//
	// This parameter is required.
	//
	// example:
	//
	// python main.py
	Entrypoint *string `json:"entrypoint,omitempty" xml:"entrypoint,omitempty"`
	// The identifier of the code runtime environment in the currently saved configuration. The runtime image uses the saved snapshot. Adding a new version to the directory does not automatically change the existing configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// python3.12
	LanguageRuntime *string `json:"languageRuntime,omitempty" xml:"languageRuntime,omitempty"`
	// The oss:// address of the platform-fixed code package referenced by the currently saved configuration. This address does not include the signature parameters of the source HTTPS link. You can use this address to reuse existing code packages. To download the content, call GetAgentCodePackageDownloadUrl.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://agentcore-platform-example/code-packages/workspaces/ws-a1b2c3d4e5f6/550e8400-e29b-41d4-a716-446655440000/package.zip
	SourceUri *string `json:"sourceUri,omitempty" xml:"sourceUri,omitempty"`
}

func (s UpdateCustomAgentResponseBodyDataArtifactCode) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentResponseBodyDataArtifactCode) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentResponseBodyDataArtifactCode) GetEntrypoint() *string {
	return s.Entrypoint
}

func (s *UpdateCustomAgentResponseBodyDataArtifactCode) GetLanguageRuntime() *string {
	return s.LanguageRuntime
}

func (s *UpdateCustomAgentResponseBodyDataArtifactCode) GetSourceUri() *string {
	return s.SourceUri
}

func (s *UpdateCustomAgentResponseBodyDataArtifactCode) SetEntrypoint(v string) *UpdateCustomAgentResponseBodyDataArtifactCode {
	s.Entrypoint = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataArtifactCode) SetLanguageRuntime(v string) *UpdateCustomAgentResponseBodyDataArtifactCode {
	s.LanguageRuntime = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataArtifactCode) SetSourceUri(v string) *UpdateCustomAgentResponseBodyDataArtifactCode {
	s.SourceUri = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataArtifactCode) Validate() error {
	return dara.Validate(s)
}

type UpdateCustomAgentResponseBodyDataArtifactContainer struct {
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
	// The source type of the image. A value of acr indicates Alibaba Cloud Container Registry.
	//
	// This parameter is required.
	//
	// example:
	//
	// acr
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateCustomAgentResponseBodyDataArtifactContainer) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentResponseBodyDataArtifactContainer) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentResponseBodyDataArtifactContainer) GetEntrypoint() *string {
	return s.Entrypoint
}

func (s *UpdateCustomAgentResponseBodyDataArtifactContainer) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateCustomAgentResponseBodyDataArtifactContainer) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateCustomAgentResponseBodyDataArtifactContainer) GetRepo() *string {
	return s.Repo
}

func (s *UpdateCustomAgentResponseBodyDataArtifactContainer) GetTag() *string {
	return s.Tag
}

func (s *UpdateCustomAgentResponseBodyDataArtifactContainer) GetType() *string {
	return s.Type
}

func (s *UpdateCustomAgentResponseBodyDataArtifactContainer) SetEntrypoint(v string) *UpdateCustomAgentResponseBodyDataArtifactContainer {
	s.Entrypoint = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataArtifactContainer) SetInstanceId(v string) *UpdateCustomAgentResponseBodyDataArtifactContainer {
	s.InstanceId = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataArtifactContainer) SetNamespace(v string) *UpdateCustomAgentResponseBodyDataArtifactContainer {
	s.Namespace = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataArtifactContainer) SetRepo(v string) *UpdateCustomAgentResponseBodyDataArtifactContainer {
	s.Repo = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataArtifactContainer) SetTag(v string) *UpdateCustomAgentResponseBodyDataArtifactContainer {
	s.Tag = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataArtifactContainer) SetType(v string) *UpdateCustomAgentResponseBodyDataArtifactContainer {
	s.Type = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataArtifactContainer) Validate() error {
	return dara.Validate(s)
}

type UpdateCustomAgentResponseBodyDataCodePackageInfo struct {
	// The CRC64 value of the code package object.
	//
	// example:
	//
	// 456
	Crc64 *string `json:"crc64,omitempty" xml:"crc64,omitempty"`
	// The ETag of the code package object.
	//
	// example:
	//
	// fixed-etag
	Etag *string `json:"etag,omitempty" xml:"etag,omitempty"`
	// The original file name of the code package.
	//
	// example:
	//
	// main.zip
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// The compressed size of the code package.
	//
	// example:
	//
	// 123
	FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// The format of the code package.
	//
	// example:
	//
	// zip
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
}

func (s UpdateCustomAgentResponseBodyDataCodePackageInfo) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentResponseBodyDataCodePackageInfo) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentResponseBodyDataCodePackageInfo) GetCrc64() *string {
	return s.Crc64
}

func (s *UpdateCustomAgentResponseBodyDataCodePackageInfo) GetEtag() *string {
	return s.Etag
}

func (s *UpdateCustomAgentResponseBodyDataCodePackageInfo) GetFileName() *string {
	return s.FileName
}

func (s *UpdateCustomAgentResponseBodyDataCodePackageInfo) GetFileSize() *int64 {
	return s.FileSize
}

func (s *UpdateCustomAgentResponseBodyDataCodePackageInfo) GetFormat() *string {
	return s.Format
}

func (s *UpdateCustomAgentResponseBodyDataCodePackageInfo) SetCrc64(v string) *UpdateCustomAgentResponseBodyDataCodePackageInfo {
	s.Crc64 = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataCodePackageInfo) SetEtag(v string) *UpdateCustomAgentResponseBodyDataCodePackageInfo {
	s.Etag = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataCodePackageInfo) SetFileName(v string) *UpdateCustomAgentResponseBodyDataCodePackageInfo {
	s.FileName = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataCodePackageInfo) SetFileSize(v int64) *UpdateCustomAgentResponseBodyDataCodePackageInfo {
	s.FileSize = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataCodePackageInfo) SetFormat(v string) *UpdateCustomAgentResponseBodyDataCodePackageInfo {
	s.Format = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataCodePackageInfo) Validate() error {
	return dara.Validate(s)
}

type UpdateCustomAgentResponseBodyDataEnvironment struct {
	// The latest submitted list of environment variables.
	Variables []*UpdateCustomAgentResponseBodyDataEnvironmentVariables `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
}

func (s UpdateCustomAgentResponseBodyDataEnvironment) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentResponseBodyDataEnvironment) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentResponseBodyDataEnvironment) GetVariables() []*UpdateCustomAgentResponseBodyDataEnvironmentVariables {
	return s.Variables
}

func (s *UpdateCustomAgentResponseBodyDataEnvironment) SetVariables(v []*UpdateCustomAgentResponseBodyDataEnvironmentVariables) *UpdateCustomAgentResponseBodyDataEnvironment {
	s.Variables = v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataEnvironment) Validate() error {
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

type UpdateCustomAgentResponseBodyDataEnvironmentVariables struct {
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

func (s UpdateCustomAgentResponseBodyDataEnvironmentVariables) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentResponseBodyDataEnvironmentVariables) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentResponseBodyDataEnvironmentVariables) GetName() *string {
	return s.Name
}

func (s *UpdateCustomAgentResponseBodyDataEnvironmentVariables) GetValue() *string {
	return s.Value
}

func (s *UpdateCustomAgentResponseBodyDataEnvironmentVariables) SetName(v string) *UpdateCustomAgentResponseBodyDataEnvironmentVariables {
	s.Name = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataEnvironmentVariables) SetValue(v string) *UpdateCustomAgentResponseBodyDataEnvironmentVariables {
	s.Value = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataEnvironmentVariables) Validate() error {
	return dara.Validate(s)
}

type UpdateCustomAgentResponseBodyDataHealthCheck struct {
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

func (s UpdateCustomAgentResponseBodyDataHealthCheck) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentResponseBodyDataHealthCheck) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentResponseBodyDataHealthCheck) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *UpdateCustomAgentResponseBodyDataHealthCheck) GetIntervalSeconds() *int32 {
	return s.IntervalSeconds
}

func (s *UpdateCustomAgentResponseBodyDataHealthCheck) GetPath() *string {
	return s.Path
}

func (s *UpdateCustomAgentResponseBodyDataHealthCheck) GetPort() *int32 {
	return s.Port
}

func (s *UpdateCustomAgentResponseBodyDataHealthCheck) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *UpdateCustomAgentResponseBodyDataHealthCheck) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *UpdateCustomAgentResponseBodyDataHealthCheck) SetHealthyThreshold(v int32) *UpdateCustomAgentResponseBodyDataHealthCheck {
	s.HealthyThreshold = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataHealthCheck) SetIntervalSeconds(v int32) *UpdateCustomAgentResponseBodyDataHealthCheck {
	s.IntervalSeconds = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataHealthCheck) SetPath(v string) *UpdateCustomAgentResponseBodyDataHealthCheck {
	s.Path = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataHealthCheck) SetPort(v int32) *UpdateCustomAgentResponseBodyDataHealthCheck {
	s.Port = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataHealthCheck) SetTimeoutSeconds(v int32) *UpdateCustomAgentResponseBodyDataHealthCheck {
	s.TimeoutSeconds = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataHealthCheck) SetUnhealthyThreshold(v int32) *UpdateCustomAgentResponseBodyDataHealthCheck {
	s.UnhealthyThreshold = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataHealthCheck) Validate() error {
	return dara.Validate(s)
}

type UpdateCustomAgentResponseBodyDataImageCachePreparation struct {
	// The ID of the associated cache for this request.
	//
	// This parameter is required.
	//
	// example:
	//
	// img-example
	ImageId *string `json:"imageId,omitempty" xml:"imageId,omitempty"`
	// The cache status at the time of acceptance.
	//
	// This parameter is required.
	//
	// example:
	//
	// PENDING
	StatusAtAcceptance *string `json:"statusAtAcceptance,omitempty" xml:"statusAtAcceptance,omitempty"`
	// The acceptance status of the cache preparation.
	//
	// This parameter is required.
	//
	// example:
	//
	// ACCEPTED
	SubmissionStatus *string `json:"submissionStatus,omitempty" xml:"submissionStatus,omitempty"`
}

func (s UpdateCustomAgentResponseBodyDataImageCachePreparation) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentResponseBodyDataImageCachePreparation) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentResponseBodyDataImageCachePreparation) GetImageId() *string {
	return s.ImageId
}

func (s *UpdateCustomAgentResponseBodyDataImageCachePreparation) GetStatusAtAcceptance() *string {
	return s.StatusAtAcceptance
}

func (s *UpdateCustomAgentResponseBodyDataImageCachePreparation) GetSubmissionStatus() *string {
	return s.SubmissionStatus
}

func (s *UpdateCustomAgentResponseBodyDataImageCachePreparation) SetImageId(v string) *UpdateCustomAgentResponseBodyDataImageCachePreparation {
	s.ImageId = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataImageCachePreparation) SetStatusAtAcceptance(v string) *UpdateCustomAgentResponseBodyDataImageCachePreparation {
	s.StatusAtAcceptance = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataImageCachePreparation) SetSubmissionStatus(v string) *UpdateCustomAgentResponseBodyDataImageCachePreparation {
	s.SubmissionStatus = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataImageCachePreparation) Validate() error {
	return dara.Validate(s)
}

type UpdateCustomAgentResponseBodyDataLatestExecutionError struct {
	// The reason code for the waiting or execution failure of the latest configuration. This code is used to determine the subsequent processing method.
	//
	// This parameter is required.
	//
	// example:
	//
	// AgentExecutionFailed
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The description of the reason why the latest configuration is waiting or failed to execute. Combine this with the code and retryable fields to determine what to do next.
	//
	// This parameter is required.
	//
	// example:
	//
	// Agent execution failed. Check the configuration and resubmit
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Specifies whether automatic retry is enabled. If the value is true, you can query again later. If the value is false, check and fix the configuration based on the reason, and then resubmit.
	//
	// This parameter is required.
	Retryable *bool `json:"retryable,omitempty" xml:"retryable,omitempty"`
}

func (s UpdateCustomAgentResponseBodyDataLatestExecutionError) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentResponseBodyDataLatestExecutionError) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentResponseBodyDataLatestExecutionError) GetCode() *string {
	return s.Code
}

func (s *UpdateCustomAgentResponseBodyDataLatestExecutionError) GetMessage() *string {
	return s.Message
}

func (s *UpdateCustomAgentResponseBodyDataLatestExecutionError) GetRetryable() *bool {
	return s.Retryable
}

func (s *UpdateCustomAgentResponseBodyDataLatestExecutionError) SetCode(v string) *UpdateCustomAgentResponseBodyDataLatestExecutionError {
	s.Code = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataLatestExecutionError) SetMessage(v string) *UpdateCustomAgentResponseBodyDataLatestExecutionError {
	s.Message = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataLatestExecutionError) SetRetryable(v bool) *UpdateCustomAgentResponseBodyDataLatestExecutionError {
	s.Retryable = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataLatestExecutionError) Validate() error {
	return dara.Validate(s)
}

type UpdateCustomAgentResponseBodyDataLogging struct {
	// The Simple Log Service (SLS) log delivery destination.
	Sls *UpdateCustomAgentResponseBodyDataLoggingSls `json:"sls,omitempty" xml:"sls,omitempty" type:"Struct"`
}

func (s UpdateCustomAgentResponseBodyDataLogging) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentResponseBodyDataLogging) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentResponseBodyDataLogging) GetSls() *UpdateCustomAgentResponseBodyDataLoggingSls {
	return s.Sls
}

func (s *UpdateCustomAgentResponseBodyDataLogging) SetSls(v *UpdateCustomAgentResponseBodyDataLoggingSls) *UpdateCustomAgentResponseBodyDataLogging {
	s.Sls = v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataLogging) Validate() error {
	if s.Sls != nil {
		if err := s.Sls.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateCustomAgentResponseBodyDataLoggingSls struct {
	// The name of the SLS logstore that receives agent logs.
	//
	// example:
	//
	// agent-runtime
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// The name of the SLS project that receives agent logs.
	//
	// example:
	//
	// my-agent-logs
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
}

func (s UpdateCustomAgentResponseBodyDataLoggingSls) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentResponseBodyDataLoggingSls) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentResponseBodyDataLoggingSls) GetLogstore() *string {
	return s.Logstore
}

func (s *UpdateCustomAgentResponseBodyDataLoggingSls) GetProject() *string {
	return s.Project
}

func (s *UpdateCustomAgentResponseBodyDataLoggingSls) SetLogstore(v string) *UpdateCustomAgentResponseBodyDataLoggingSls {
	s.Logstore = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataLoggingSls) SetProject(v string) *UpdateCustomAgentResponseBodyDataLoggingSls {
	s.Project = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataLoggingSls) Validate() error {
	return dara.Validate(s)
}

type UpdateCustomAgentResponseBodyDataNetwork struct {
	// The Internet access configuration of the agent. If this field is not configured, Internet access is allowed on the agent side but is still limited by the workspace.
	AccessInternet *UpdateCustomAgentResponseBodyDataNetworkAccessInternet `json:"accessInternet,omitempty" xml:"accessInternet,omitempty" type:"Struct"`
	// The user VPC access configuration of the agent. If this field is not configured, user VPC access is allowed on the agent side but is still limited by the workspace.
	AccessVpc *UpdateCustomAgentResponseBodyDataNetworkAccessVpc `json:"accessVpc,omitempty" xml:"accessVpc,omitempty" type:"Struct"`
}

func (s UpdateCustomAgentResponseBodyDataNetwork) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentResponseBodyDataNetwork) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentResponseBodyDataNetwork) GetAccessInternet() *UpdateCustomAgentResponseBodyDataNetworkAccessInternet {
	return s.AccessInternet
}

func (s *UpdateCustomAgentResponseBodyDataNetwork) GetAccessVpc() *UpdateCustomAgentResponseBodyDataNetworkAccessVpc {
	return s.AccessVpc
}

func (s *UpdateCustomAgentResponseBodyDataNetwork) SetAccessInternet(v *UpdateCustomAgentResponseBodyDataNetworkAccessInternet) *UpdateCustomAgentResponseBodyDataNetwork {
	s.AccessInternet = v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataNetwork) SetAccessVpc(v *UpdateCustomAgentResponseBodyDataNetworkAccessVpc) *UpdateCustomAgentResponseBodyDataNetwork {
	s.AccessVpc = v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataNetwork) Validate() error {
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

type UpdateCustomAgentResponseBodyDataNetworkAccessInternet struct {
	// Specifies whether Internet access is allowed on the agent side. The actual access is still limited by the workspace.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s UpdateCustomAgentResponseBodyDataNetworkAccessInternet) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentResponseBodyDataNetworkAccessInternet) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentResponseBodyDataNetworkAccessInternet) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateCustomAgentResponseBodyDataNetworkAccessInternet) SetEnabled(v bool) *UpdateCustomAgentResponseBodyDataNetworkAccessInternet {
	s.Enabled = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataNetworkAccessInternet) Validate() error {
	return dara.Validate(s)
}

type UpdateCustomAgentResponseBodyDataNetworkAccessVpc struct {
	// Specifies whether access to the user VPC configured for the workspace is allowed on the agent side. The actual access is still limited by the workspace.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s UpdateCustomAgentResponseBodyDataNetworkAccessVpc) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentResponseBodyDataNetworkAccessVpc) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentResponseBodyDataNetworkAccessVpc) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateCustomAgentResponseBodyDataNetworkAccessVpc) SetEnabled(v bool) *UpdateCustomAgentResponseBodyDataNetworkAccessVpc {
	s.Enabled = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataNetworkAccessVpc) Validate() error {
	return dara.Validate(s)
}

type UpdateCustomAgentResponseBodyDataObservability struct {
	// Specifies whether to enable agent observability.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s UpdateCustomAgentResponseBodyDataObservability) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentResponseBodyDataObservability) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentResponseBodyDataObservability) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateCustomAgentResponseBodyDataObservability) SetEnabled(v bool) *UpdateCustomAgentResponseBodyDataObservability {
	s.Enabled = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataObservability) Validate() error {
	return dara.Validate(s)
}

type UpdateCustomAgentResponseBodyDataOssMounts struct {
	// The name of the mounted OSS bucket.
	//
	// example:
	//
	// my-agent-data
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// The OSS mount directory in the container, which is located under /mnt/oss/.
	//
	// example:
	//
	// /mnt/oss/train
	MountPath *string `json:"mountPath,omitempty" xml:"mountPath,omitempty"`
	// The object directory prefix in the mounted bucket.
	//
	// example:
	//
	// datasets/train
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// Specifies whether the mount is read-only. A value of true indicates a read-only mount. A value of false indicates a read-write mount.
	ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
}

func (s UpdateCustomAgentResponseBodyDataOssMounts) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentResponseBodyDataOssMounts) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentResponseBodyDataOssMounts) GetBucketName() *string {
	return s.BucketName
}

func (s *UpdateCustomAgentResponseBodyDataOssMounts) GetMountPath() *string {
	return s.MountPath
}

func (s *UpdateCustomAgentResponseBodyDataOssMounts) GetPath() *string {
	return s.Path
}

func (s *UpdateCustomAgentResponseBodyDataOssMounts) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *UpdateCustomAgentResponseBodyDataOssMounts) SetBucketName(v string) *UpdateCustomAgentResponseBodyDataOssMounts {
	s.BucketName = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataOssMounts) SetMountPath(v string) *UpdateCustomAgentResponseBodyDataOssMounts {
	s.MountPath = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataOssMounts) SetPath(v string) *UpdateCustomAgentResponseBodyDataOssMounts {
	s.Path = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataOssMounts) SetReadOnly(v bool) *UpdateCustomAgentResponseBodyDataOssMounts {
	s.ReadOnly = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataOssMounts) Validate() error {
	return dara.Validate(s)
}

type UpdateCustomAgentResponseBodyDataProtocol struct {
	// The request header configuration for the current protocol. The key is the name and the value is a string.
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	// The list of API paths for the current protocol.
	Paths []*UpdateCustomAgentResponseBodyDataProtocolPaths `json:"paths,omitempty" xml:"paths,omitempty" type:"Repeated"`
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

func (s UpdateCustomAgentResponseBodyDataProtocol) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentResponseBodyDataProtocol) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentResponseBodyDataProtocol) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCustomAgentResponseBodyDataProtocol) GetPaths() []*UpdateCustomAgentResponseBodyDataProtocolPaths {
	return s.Paths
}

func (s *UpdateCustomAgentResponseBodyDataProtocol) GetType() *string {
	return s.Type
}

func (s *UpdateCustomAgentResponseBodyDataProtocol) SetHeaders(v map[string]*string) *UpdateCustomAgentResponseBodyDataProtocol {
	s.Headers = v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataProtocol) SetPaths(v []*UpdateCustomAgentResponseBodyDataProtocolPaths) *UpdateCustomAgentResponseBodyDataProtocol {
	s.Paths = v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataProtocol) SetType(v string) *UpdateCustomAgentResponseBodyDataProtocol {
	s.Type = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataProtocol) Validate() error {
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

type UpdateCustomAgentResponseBodyDataProtocolPaths struct {
	// The usage description of the current API path.
	//
	// example:
	//
	// AG-UI conversation API
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The interface path corresponding to the current protocol.
	//
	// This parameter is required.
	//
	// example:
	//
	// /agui
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s UpdateCustomAgentResponseBodyDataProtocolPaths) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentResponseBodyDataProtocolPaths) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentResponseBodyDataProtocolPaths) GetDescription() *string {
	return s.Description
}

func (s *UpdateCustomAgentResponseBodyDataProtocolPaths) GetName() *string {
	return s.Name
}

func (s *UpdateCustomAgentResponseBodyDataProtocolPaths) SetDescription(v string) *UpdateCustomAgentResponseBodyDataProtocolPaths {
	s.Description = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataProtocolPaths) SetName(v string) *UpdateCustomAgentResponseBodyDataProtocolPaths {
	s.Name = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataProtocolPaths) Validate() error {
	return dara.Validate(s)
}

type UpdateCustomAgentResponseBodyDataRuntime struct {
	// The compute resource configuration for the agent runtime instance.
	//
	// This parameter is required.
	Compute *UpdateCustomAgentResponseBodyDataRuntimeCompute `json:"compute,omitempty" xml:"compute,omitempty" type:"Struct"`
	// The elastic policy configuration of the agent.
	Hpa *UpdateCustomAgentResponseBodyDataRuntimeHpa `json:"hpa,omitempty" xml:"hpa,omitempty" type:"Struct"`
	// The session policy configuration of the agent.
	//
	// This parameter is required.
	SessionPolicy *UpdateCustomAgentResponseBodyDataRuntimeSessionPolicy `json:"sessionPolicy,omitempty" xml:"sessionPolicy,omitempty" type:"Struct"`
}

func (s UpdateCustomAgentResponseBodyDataRuntime) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentResponseBodyDataRuntime) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentResponseBodyDataRuntime) GetCompute() *UpdateCustomAgentResponseBodyDataRuntimeCompute {
	return s.Compute
}

func (s *UpdateCustomAgentResponseBodyDataRuntime) GetHpa() *UpdateCustomAgentResponseBodyDataRuntimeHpa {
	return s.Hpa
}

func (s *UpdateCustomAgentResponseBodyDataRuntime) GetSessionPolicy() *UpdateCustomAgentResponseBodyDataRuntimeSessionPolicy {
	return s.SessionPolicy
}

func (s *UpdateCustomAgentResponseBodyDataRuntime) SetCompute(v *UpdateCustomAgentResponseBodyDataRuntimeCompute) *UpdateCustomAgentResponseBodyDataRuntime {
	s.Compute = v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataRuntime) SetHpa(v *UpdateCustomAgentResponseBodyDataRuntimeHpa) *UpdateCustomAgentResponseBodyDataRuntime {
	s.Hpa = v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataRuntime) SetSessionPolicy(v *UpdateCustomAgentResponseBodyDataRuntimeSessionPolicy) *UpdateCustomAgentResponseBodyDataRuntime {
	s.SessionPolicy = v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataRuntime) Validate() error {
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

type UpdateCustomAgentResponseBodyDataRuntimeCompute struct {
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

func (s UpdateCustomAgentResponseBodyDataRuntimeCompute) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentResponseBodyDataRuntimeCompute) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentResponseBodyDataRuntimeCompute) GetComputeClass() *string {
	return s.ComputeClass
}

func (s *UpdateCustomAgentResponseBodyDataRuntimeCompute) SetComputeClass(v string) *UpdateCustomAgentResponseBodyDataRuntimeCompute {
	s.ComputeClass = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataRuntimeCompute) Validate() error {
	return dara.Validate(s)
}

type UpdateCustomAgentResponseBodyDataRuntimeHpa struct {
	// Specifies whether elastic scaling is enabled. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// This parameter is required.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The maximum number of concurrent sessions configured for a single running instance.
	//
	// example:
	//
	// 100
	MaxConcurrentSessionsPerSandbox *int32 `json:"maxConcurrentSessionsPerSandbox,omitempty" xml:"maxConcurrentSessionsPerSandbox,omitempty"`
	// The maximum number of running instances configured.
	//
	// example:
	//
	// 10
	MaxSandboxCount *int32 `json:"maxSandboxCount,omitempty" xml:"maxSandboxCount,omitempty"`
	// The minimum number of running instances configured.
	//
	// example:
	//
	// 1
	MinSandboxCount *int32 `json:"minSandboxCount,omitempty" xml:"minSandboxCount,omitempty"`
	// The maximum idle time for a session, in seconds.
	//
	// example:
	//
	// 600
	SessionTtlSeconds *int32 `json:"sessionTtlSeconds,omitempty" xml:"sessionTtlSeconds,omitempty"`
}

func (s UpdateCustomAgentResponseBodyDataRuntimeHpa) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentResponseBodyDataRuntimeHpa) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentResponseBodyDataRuntimeHpa) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateCustomAgentResponseBodyDataRuntimeHpa) GetMaxConcurrentSessionsPerSandbox() *int32 {
	return s.MaxConcurrentSessionsPerSandbox
}

func (s *UpdateCustomAgentResponseBodyDataRuntimeHpa) GetMaxSandboxCount() *int32 {
	return s.MaxSandboxCount
}

func (s *UpdateCustomAgentResponseBodyDataRuntimeHpa) GetMinSandboxCount() *int32 {
	return s.MinSandboxCount
}

func (s *UpdateCustomAgentResponseBodyDataRuntimeHpa) GetSessionTtlSeconds() *int32 {
	return s.SessionTtlSeconds
}

func (s *UpdateCustomAgentResponseBodyDataRuntimeHpa) SetEnabled(v bool) *UpdateCustomAgentResponseBodyDataRuntimeHpa {
	s.Enabled = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataRuntimeHpa) SetMaxConcurrentSessionsPerSandbox(v int32) *UpdateCustomAgentResponseBodyDataRuntimeHpa {
	s.MaxConcurrentSessionsPerSandbox = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataRuntimeHpa) SetMaxSandboxCount(v int32) *UpdateCustomAgentResponseBodyDataRuntimeHpa {
	s.MaxSandboxCount = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataRuntimeHpa) SetMinSandboxCount(v int32) *UpdateCustomAgentResponseBodyDataRuntimeHpa {
	s.MinSandboxCount = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataRuntimeHpa) SetSessionTtlSeconds(v int32) *UpdateCustomAgentResponseBodyDataRuntimeHpa {
	s.SessionTtlSeconds = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataRuntimeHpa) Validate() error {
	return dara.Validate(s)
}

type UpdateCustomAgentResponseBodyDataRuntimeSessionPolicy struct {
	// The name of the HTTP request header used by the session policy. This applies to ISOLATED_HEADER_FIELD.
	//
	// example:
	//
	// X-Session-Id
	HeaderName *string `json:"headerName,omitempty" xml:"headerName,omitempty"`
	// The session policy type. Valid values: AFFINITY: Session affinity. ISOLATED_HEADER_FIELD: Uses the request header specified by headerName to associate sessions. ISOLATED_COOKIE is not recommended. The compatible value DISABLED in requests is returned as AFFINITY.
	//
	// This parameter is required.
	//
	// example:
	//
	// AFFINITY
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateCustomAgentResponseBodyDataRuntimeSessionPolicy) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentResponseBodyDataRuntimeSessionPolicy) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentResponseBodyDataRuntimeSessionPolicy) GetHeaderName() *string {
	return s.HeaderName
}

func (s *UpdateCustomAgentResponseBodyDataRuntimeSessionPolicy) GetType() *string {
	return s.Type
}

func (s *UpdateCustomAgentResponseBodyDataRuntimeSessionPolicy) SetHeaderName(v string) *UpdateCustomAgentResponseBodyDataRuntimeSessionPolicy {
	s.HeaderName = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataRuntimeSessionPolicy) SetType(v string) *UpdateCustomAgentResponseBodyDataRuntimeSessionPolicy {
	s.Type = &v
	return s
}

func (s *UpdateCustomAgentResponseBodyDataRuntimeSessionPolicy) Validate() error {
	return dara.Validate(s)
}
