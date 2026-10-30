// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCustomAgentResponseBody
	GetCode() *string
	SetData(v *GetCustomAgentResponseBodyData) *GetCustomAgentResponseBody
	GetData() *GetCustomAgentResponseBodyData
	SetHttpStatusCode(v int32) *GetCustomAgentResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetCustomAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCustomAgentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCustomAgentResponseBody
	GetSuccess() *bool
}

type GetCustomAgentResponseBody struct {
	// The request result code. The value is SUCCESS if the request is successful.
	//
	// This parameter is required.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The latest submitted configuration and status of the custom agent. To determine whether the configuration has taken effect, check the latestVersionStatus parameter.
	//
	// This parameter is required.
	Data *GetCustomAgentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code of the request. The value is 200 if the request is successful.
	//
	// This parameter is required.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The description of the processing result for the request. The value is success if the request is successful.
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
	// Indicates whether the request was processed successfully. A value of true indicates that the request was processed successfully. A successful creation, update, or deletion does not mean that the background operation is complete.
	//
	// This parameter is required.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetCustomAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCustomAgentResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCustomAgentResponseBody) GetData() *GetCustomAgentResponseBodyData {
	return s.Data
}

func (s *GetCustomAgentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetCustomAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCustomAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCustomAgentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCustomAgentResponseBody) SetCode(v string) *GetCustomAgentResponseBody {
	s.Code = &v
	return s
}

func (s *GetCustomAgentResponseBody) SetData(v *GetCustomAgentResponseBodyData) *GetCustomAgentResponseBody {
	s.Data = v
	return s
}

func (s *GetCustomAgentResponseBody) SetHttpStatusCode(v int32) *GetCustomAgentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetCustomAgentResponseBody) SetMessage(v string) *GetCustomAgentResponseBody {
	s.Message = &v
	return s
}

func (s *GetCustomAgentResponseBody) SetRequestId(v string) *GetCustomAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCustomAgentResponseBody) SetSuccess(v bool) *GetCustomAgentResponseBody {
	s.Success = &v
	return s
}

func (s *GetCustomAgentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCustomAgentResponseBodyData struct {
	// The unique identifier of the custom agent. You can obtain this ID from the response of CreateCustomAgent or ListCustomAgents.
	//
	// This parameter is required.
	//
	// example:
	//
	// agent-550e8400e29b41d4a716446655440000
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The latest submitted list of additional AgenticFS mounts.
	AgenticFsMounts []*GetCustomAgentResponseBodyDataAgenticFsMounts `json:"agenticFsMounts,omitempty" xml:"agenticFsMounts,omitempty" type:"Repeated"`
	// The latest submitted deployment artifact configuration, including the service listening port and the ACR container image. To determine whether the configuration has taken effect, check the latestVersionStatus parameter.
	Artifact *GetCustomAgentResponseBodyDataArtifact `json:"artifact,omitempty" xml:"artifact,omitempty" type:"Struct"`
	// The currently saved code package information.
	CodePackageInfo *GetCustomAgentResponseBodyDataCodePackageInfo `json:"codePackageInfo,omitempty" xml:"codePackageInfo,omitempty" type:"Struct"`
	// The creation method of the agent. A value of CUSTOM is returned for custom agents.
	//
	// example:
	//
	// CUSTOM
	CreateMode *string `json:"createMode,omitempty" xml:"createMode,omitempty"`
	// The time when the agent was created. The time is in ISO 8601 format. The trailing Z indicates the UTC time zone. Fractional seconds may be included.
	//
	// example:
	//
	// 2026-10-10T06:00:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The image source snapshot of the current active configuration version. This value is null before the initial configuration takes effect and does not represent the health status.
	CurrentImage *GetCustomAgentResponseBodyDataCurrentImage `json:"currentImage,omitempty" xml:"currentImage,omitempty" type:"Struct"`
	// The recently saved source and its verified username. The source may be in preparation or failed. No password is returned.
	CustomImage *GetCustomAgentResponseBodyDataCustomImage `json:"customImage,omitempty" xml:"customImage,omitempty" type:"Struct"`
	// The deployment method of the agent. Custom agents are deployed and managed by the platform, and MANAGED is returned.
	//
	// example:
	//
	// MANAGED
	DeployType *string `json:"deployType,omitempty" xml:"deployType,omitempty"`
	// The description of the agent\\"s purpose.
	//
	// example:
	//
	// Custom agent for providing conversational services
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The active configuration version.
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
	Environment *GetCustomAgentResponseBodyDataEnvironment `json:"environment,omitempty" xml:"environment,omitempty" type:"Struct"`
	// The HTTP health check configuration for the agent. Empty if not configured.
	HealthCheck *GetCustomAgentResponseBodyDataHealthCheck `json:"healthCheck,omitempty" xml:"healthCheck,omitempty" type:"Struct"`
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
	// The currently associated image preparation task. The phase indicates acceptance. Query the preparation API for real-time progress. A standard save operation does not create an image preparation task.
	ImagePreparation *GetCustomAgentResponseBodyDataImagePreparation `json:"imagePreparation,omitempty" xml:"imagePreparation,omitempty" type:"Struct"`
	// The waiting or execution error of the latest configuration version. Returned only when this version is waiting for a retry or has failed.
	LatestExecutionError *GetCustomAgentResponseBodyDataLatestExecutionError `json:"latestExecutionError,omitempty" xml:"latestExecutionError,omitempty" type:"Struct"`
	// The latest configuration version number. Updates with the same configuration may retain the original version number. This field does not indicate that the configuration has taken effect.
	//
	// This parameter is required.
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
	// - succeeded: Executed successfully.
	//
	// - failed: Execution failed.
	//
	// - superseded: Superseded by a subsequent version.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// succeeded
	LatestVersionStatus *string `json:"latestVersionStatus,omitempty" xml:"latestVersionStatus,omitempty"`
	// The latest submitted log delivery configuration.
	Logging *GetCustomAgentResponseBodyDataLogging `json:"logging,omitempty" xml:"logging,omitempty" type:"Struct"`
	// The name of the custom agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-custom-agent
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The network access configuration of the agent. The actual access scope is also restricted by the workspace network configuration.
	Network *GetCustomAgentResponseBodyDataNetwork `json:"network,omitempty" xml:"network,omitempty" type:"Struct"`
	// The observability configuration for the high-code agent.
	Observability *GetCustomAgentResponseBodyDataObservability `json:"observability,omitempty" xml:"observability,omitempty" type:"Struct"`
	// The list of recently submitted OSS directory mounts.
	OssMounts []*GetCustomAgentResponseBodyDataOssMounts `json:"ossMounts,omitempty" xml:"ossMounts,omitempty" type:"Repeated"`
	// The list of invocation protocol configurations for the agent.
	Protocol []*GetCustomAgentResponseBodyDataProtocol `json:"protocol,omitempty" xml:"protocol,omitempty" type:"Repeated"`
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
	// The latest committed runtime configuration, including compute specifications, session policy, and elastic policy.
	Runtime *GetCustomAgentResponseBodyDataRuntime `json:"runtime,omitempty" xml:"runtime,omitempty" type:"Struct"`
	// Counts the number of instances by runtime instance phase. Keys include PENDING (creating or initializing), RUNNING (running), HIBERNATING (hibernating), HIBERNATED (hibernated), RESUMING (resuming), TERMINATING (terminating), and FAILED (failed). Only phases that actually occur are returned, and missing keys are treated as 0.
	SandboxPhaseCounts map[string]*int64 `json:"sandboxPhaseCounts,omitempty" xml:"sandboxPhaseCounts,omitempty"`
	// The overall status of the agent. Valid values: Creating: Creating. Running: Running. Failed: Failed. Updating: Updating. Deleting: Deleting. Deleted: Deleted. To check the execution progress of the latest configuration, query latestVersionStatus.
	//
	// This parameter is required.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The most recent update time of the agent, in ISO 8601 format. The trailing Z indicates the UTC time zone, and fractional seconds may be included.
	//
	// example:
	//
	// 2026-10-10T06:05:00Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// The image cache switch.
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

func (s GetCustomAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCustomAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCustomAgentResponseBodyData) GetAgentId() *string {
	return s.AgentId
}

func (s *GetCustomAgentResponseBodyData) GetAgenticFsMounts() []*GetCustomAgentResponseBodyDataAgenticFsMounts {
	return s.AgenticFsMounts
}

func (s *GetCustomAgentResponseBodyData) GetArtifact() *GetCustomAgentResponseBodyDataArtifact {
	return s.Artifact
}

func (s *GetCustomAgentResponseBodyData) GetCodePackageInfo() *GetCustomAgentResponseBodyDataCodePackageInfo {
	return s.CodePackageInfo
}

func (s *GetCustomAgentResponseBodyData) GetCreateMode() *string {
	return s.CreateMode
}

func (s *GetCustomAgentResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetCustomAgentResponseBodyData) GetCurrentImage() *GetCustomAgentResponseBodyDataCurrentImage {
	return s.CurrentImage
}

func (s *GetCustomAgentResponseBodyData) GetCustomImage() *GetCustomAgentResponseBodyDataCustomImage {
	return s.CustomImage
}

func (s *GetCustomAgentResponseBodyData) GetDeployType() *string {
	return s.DeployType
}

func (s *GetCustomAgentResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetCustomAgentResponseBodyData) GetEffectiveSpecVersion() *int64 {
	return s.EffectiveSpecVersion
}

func (s *GetCustomAgentResponseBodyData) GetEnvironment() *GetCustomAgentResponseBodyDataEnvironment {
	return s.Environment
}

func (s *GetCustomAgentResponseBodyData) GetHealthCheck() *GetCustomAgentResponseBodyDataHealthCheck {
	return s.HealthCheck
}

func (s *GetCustomAgentResponseBodyData) GetImageCacheId() *string {
	return s.ImageCacheId
}

func (s *GetCustomAgentResponseBodyData) GetImagePreparation() *GetCustomAgentResponseBodyDataImagePreparation {
	return s.ImagePreparation
}

func (s *GetCustomAgentResponseBodyData) GetLatestExecutionError() *GetCustomAgentResponseBodyDataLatestExecutionError {
	return s.LatestExecutionError
}

func (s *GetCustomAgentResponseBodyData) GetLatestSpecVersion() *int64 {
	return s.LatestSpecVersion
}

func (s *GetCustomAgentResponseBodyData) GetLatestVersionStatus() *string {
	return s.LatestVersionStatus
}

func (s *GetCustomAgentResponseBodyData) GetLogging() *GetCustomAgentResponseBodyDataLogging {
	return s.Logging
}

func (s *GetCustomAgentResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetCustomAgentResponseBodyData) GetNetwork() *GetCustomAgentResponseBodyDataNetwork {
	return s.Network
}

func (s *GetCustomAgentResponseBodyData) GetObservability() *GetCustomAgentResponseBodyDataObservability {
	return s.Observability
}

func (s *GetCustomAgentResponseBodyData) GetOssMounts() []*GetCustomAgentResponseBodyDataOssMounts {
	return s.OssMounts
}

func (s *GetCustomAgentResponseBodyData) GetProtocol() []*GetCustomAgentResponseBodyDataProtocol {
	return s.Protocol
}

func (s *GetCustomAgentResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetCustomAgentResponseBodyData) GetRoleArn() *string {
	return s.RoleArn
}

func (s *GetCustomAgentResponseBodyData) GetRuntime() *GetCustomAgentResponseBodyDataRuntime {
	return s.Runtime
}

func (s *GetCustomAgentResponseBodyData) GetSandboxPhaseCounts() map[string]*int64 {
	return s.SandboxPhaseCounts
}

func (s *GetCustomAgentResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetCustomAgentResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetCustomAgentResponseBodyData) GetUseImageCache() *bool {
	return s.UseImageCache
}

func (s *GetCustomAgentResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetCustomAgentResponseBodyData) SetAgentId(v string) *GetCustomAgentResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *GetCustomAgentResponseBodyData) SetAgenticFsMounts(v []*GetCustomAgentResponseBodyDataAgenticFsMounts) *GetCustomAgentResponseBodyData {
	s.AgenticFsMounts = v
	return s
}

func (s *GetCustomAgentResponseBodyData) SetArtifact(v *GetCustomAgentResponseBodyDataArtifact) *GetCustomAgentResponseBodyData {
	s.Artifact = v
	return s
}

func (s *GetCustomAgentResponseBodyData) SetCodePackageInfo(v *GetCustomAgentResponseBodyDataCodePackageInfo) *GetCustomAgentResponseBodyData {
	s.CodePackageInfo = v
	return s
}

func (s *GetCustomAgentResponseBodyData) SetCreateMode(v string) *GetCustomAgentResponseBodyData {
	s.CreateMode = &v
	return s
}

func (s *GetCustomAgentResponseBodyData) SetCreatedAt(v string) *GetCustomAgentResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetCustomAgentResponseBodyData) SetCurrentImage(v *GetCustomAgentResponseBodyDataCurrentImage) *GetCustomAgentResponseBodyData {
	s.CurrentImage = v
	return s
}

func (s *GetCustomAgentResponseBodyData) SetCustomImage(v *GetCustomAgentResponseBodyDataCustomImage) *GetCustomAgentResponseBodyData {
	s.CustomImage = v
	return s
}

func (s *GetCustomAgentResponseBodyData) SetDeployType(v string) *GetCustomAgentResponseBodyData {
	s.DeployType = &v
	return s
}

func (s *GetCustomAgentResponseBodyData) SetDescription(v string) *GetCustomAgentResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetCustomAgentResponseBodyData) SetEffectiveSpecVersion(v int64) *GetCustomAgentResponseBodyData {
	s.EffectiveSpecVersion = &v
	return s
}

func (s *GetCustomAgentResponseBodyData) SetEnvironment(v *GetCustomAgentResponseBodyDataEnvironment) *GetCustomAgentResponseBodyData {
	s.Environment = v
	return s
}

func (s *GetCustomAgentResponseBodyData) SetHealthCheck(v *GetCustomAgentResponseBodyDataHealthCheck) *GetCustomAgentResponseBodyData {
	s.HealthCheck = v
	return s
}

func (s *GetCustomAgentResponseBodyData) SetImageCacheId(v string) *GetCustomAgentResponseBodyData {
	s.ImageCacheId = &v
	return s
}

func (s *GetCustomAgentResponseBodyData) SetImagePreparation(v *GetCustomAgentResponseBodyDataImagePreparation) *GetCustomAgentResponseBodyData {
	s.ImagePreparation = v
	return s
}

func (s *GetCustomAgentResponseBodyData) SetLatestExecutionError(v *GetCustomAgentResponseBodyDataLatestExecutionError) *GetCustomAgentResponseBodyData {
	s.LatestExecutionError = v
	return s
}

func (s *GetCustomAgentResponseBodyData) SetLatestSpecVersion(v int64) *GetCustomAgentResponseBodyData {
	s.LatestSpecVersion = &v
	return s
}

func (s *GetCustomAgentResponseBodyData) SetLatestVersionStatus(v string) *GetCustomAgentResponseBodyData {
	s.LatestVersionStatus = &v
	return s
}

func (s *GetCustomAgentResponseBodyData) SetLogging(v *GetCustomAgentResponseBodyDataLogging) *GetCustomAgentResponseBodyData {
	s.Logging = v
	return s
}

func (s *GetCustomAgentResponseBodyData) SetName(v string) *GetCustomAgentResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetCustomAgentResponseBodyData) SetNetwork(v *GetCustomAgentResponseBodyDataNetwork) *GetCustomAgentResponseBodyData {
	s.Network = v
	return s
}

func (s *GetCustomAgentResponseBodyData) SetObservability(v *GetCustomAgentResponseBodyDataObservability) *GetCustomAgentResponseBodyData {
	s.Observability = v
	return s
}

func (s *GetCustomAgentResponseBodyData) SetOssMounts(v []*GetCustomAgentResponseBodyDataOssMounts) *GetCustomAgentResponseBodyData {
	s.OssMounts = v
	return s
}

func (s *GetCustomAgentResponseBodyData) SetProtocol(v []*GetCustomAgentResponseBodyDataProtocol) *GetCustomAgentResponseBodyData {
	s.Protocol = v
	return s
}

func (s *GetCustomAgentResponseBodyData) SetRegionId(v string) *GetCustomAgentResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetCustomAgentResponseBodyData) SetRoleArn(v string) *GetCustomAgentResponseBodyData {
	s.RoleArn = &v
	return s
}

func (s *GetCustomAgentResponseBodyData) SetRuntime(v *GetCustomAgentResponseBodyDataRuntime) *GetCustomAgentResponseBodyData {
	s.Runtime = v
	return s
}

func (s *GetCustomAgentResponseBodyData) SetSandboxPhaseCounts(v map[string]*int64) *GetCustomAgentResponseBodyData {
	s.SandboxPhaseCounts = v
	return s
}

func (s *GetCustomAgentResponseBodyData) SetStatus(v string) *GetCustomAgentResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetCustomAgentResponseBodyData) SetUpdatedAt(v string) *GetCustomAgentResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *GetCustomAgentResponseBodyData) SetUseImageCache(v bool) *GetCustomAgentResponseBodyData {
	s.UseImageCache = &v
	return s
}

func (s *GetCustomAgentResponseBodyData) SetWorkspaceId(v string) *GetCustomAgentResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *GetCustomAgentResponseBodyData) Validate() error {
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
	if s.CurrentImage != nil {
		if err := s.CurrentImage.Validate(); err != nil {
			return err
		}
	}
	if s.CustomImage != nil {
		if err := s.CustomImage.Validate(); err != nil {
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
	if s.ImagePreparation != nil {
		if err := s.ImagePreparation.Validate(); err != nil {
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

type GetCustomAgentResponseBodyDataAgenticFsMounts struct {
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
	// Specifies whether the mount is read-only. A value of true indicates a read-only mount. A value of false indicates a read-write mount.
	//
	// example:
	//
	// false
	ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
	// The AgenticFS access point domain name used for the mount.
	//
	// example:
	//
	// ap-0123456789abcdef0.0123456789-vlm36.cn-hangzhou.nas.aliyuncs.com
	Server *string `json:"server,omitempty" xml:"server,omitempty"`
}

func (s GetCustomAgentResponseBodyDataAgenticFsMounts) String() string {
	return dara.Prettify(s)
}

func (s GetCustomAgentResponseBodyDataAgenticFsMounts) GoString() string {
	return s.String()
}

func (s *GetCustomAgentResponseBodyDataAgenticFsMounts) GetMountPath() *string {
	return s.MountPath
}

func (s *GetCustomAgentResponseBodyDataAgenticFsMounts) GetPath() *string {
	return s.Path
}

func (s *GetCustomAgentResponseBodyDataAgenticFsMounts) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *GetCustomAgentResponseBodyDataAgenticFsMounts) GetServer() *string {
	return s.Server
}

func (s *GetCustomAgentResponseBodyDataAgenticFsMounts) SetMountPath(v string) *GetCustomAgentResponseBodyDataAgenticFsMounts {
	s.MountPath = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataAgenticFsMounts) SetPath(v string) *GetCustomAgentResponseBodyDataAgenticFsMounts {
	s.Path = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataAgenticFsMounts) SetReadOnly(v bool) *GetCustomAgentResponseBodyDataAgenticFsMounts {
	s.ReadOnly = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataAgenticFsMounts) SetServer(v string) *GetCustomAgentResponseBodyDataAgenticFsMounts {
	s.Server = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataAgenticFsMounts) Validate() error {
	return dara.Validate(s)
}

type GetCustomAgentResponseBodyDataArtifact struct {
	// The code package deployment configuration. This parameter is mutually exclusive with the container parameter. The code package uses the ZIP or TAR.GZ format. Dependencies must be packaged by the user. The platform does not install dependencies or perform online builds.
	Code *GetCustomAgentResponseBodyDataArtifactCode `json:"code,omitempty" xml:"code,omitempty" type:"Struct"`
	// The latest submitted ACR container image configuration.
	Container *GetCustomAgentResponseBodyDataArtifactContainer `json:"container,omitempty" xml:"container,omitempty" type:"Struct"`
	// The TCP port listened on by the agent service.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8080
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
}

func (s GetCustomAgentResponseBodyDataArtifact) String() string {
	return dara.Prettify(s)
}

func (s GetCustomAgentResponseBodyDataArtifact) GoString() string {
	return s.String()
}

func (s *GetCustomAgentResponseBodyDataArtifact) GetCode() *GetCustomAgentResponseBodyDataArtifactCode {
	return s.Code
}

func (s *GetCustomAgentResponseBodyDataArtifact) GetContainer() *GetCustomAgentResponseBodyDataArtifactContainer {
	return s.Container
}

func (s *GetCustomAgentResponseBodyDataArtifact) GetPort() *int32 {
	return s.Port
}

func (s *GetCustomAgentResponseBodyDataArtifact) SetCode(v *GetCustomAgentResponseBodyDataArtifactCode) *GetCustomAgentResponseBodyDataArtifact {
	s.Code = v
	return s
}

func (s *GetCustomAgentResponseBodyDataArtifact) SetContainer(v *GetCustomAgentResponseBodyDataArtifactContainer) *GetCustomAgentResponseBodyDataArtifact {
	s.Container = v
	return s
}

func (s *GetCustomAgentResponseBodyDataArtifact) SetPort(v int32) *GetCustomAgentResponseBodyDataArtifact {
	s.Port = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataArtifact) Validate() error {
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

type GetCustomAgentResponseBodyDataArtifactCode struct {
	// The startup command for the code package. This parameter is required in code package mode. The working directory is the extracted code root directory /code, such as python main.py. Quotes and escapes are supported. Environment variables, pipes, and logical operations are not automatically expanded. For complex commands, use a script within the package.
	//
	// This parameter is required.
	//
	// example:
	//
	// python main.py
	Entrypoint *string `json:"entrypoint,omitempty" xml:"entrypoint,omitempty"`
	// The identifier of the currently saved code runtime environment configuration. The runtime image uses the saved snapshot. Adding a new version to the directory does not automatically change the existing configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// python3.12
	LanguageRuntime *string `json:"languageRuntime,omitempty" xml:"languageRuntime,omitempty"`
	// The fixed oss:// address of the platform code package referenced by the currently saved configuration. This address does not include the signature parameters of the source HTTPS link. You can use this address to reuse an existing code package. To download the content, call the GetAgentCodePackageDownloadUrl operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://agentcore-platform-example/code-packages/workspaces/ws-a1b2c3d4e5f6/550e8400-e29b-41d4-a716-446655440000/package.zip
	SourceUri *string `json:"sourceUri,omitempty" xml:"sourceUri,omitempty"`
}

func (s GetCustomAgentResponseBodyDataArtifactCode) String() string {
	return dara.Prettify(s)
}

func (s GetCustomAgentResponseBodyDataArtifactCode) GoString() string {
	return s.String()
}

func (s *GetCustomAgentResponseBodyDataArtifactCode) GetEntrypoint() *string {
	return s.Entrypoint
}

func (s *GetCustomAgentResponseBodyDataArtifactCode) GetLanguageRuntime() *string {
	return s.LanguageRuntime
}

func (s *GetCustomAgentResponseBodyDataArtifactCode) GetSourceUri() *string {
	return s.SourceUri
}

func (s *GetCustomAgentResponseBodyDataArtifactCode) SetEntrypoint(v string) *GetCustomAgentResponseBodyDataArtifactCode {
	s.Entrypoint = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataArtifactCode) SetLanguageRuntime(v string) *GetCustomAgentResponseBodyDataArtifactCode {
	s.LanguageRuntime = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataArtifactCode) SetSourceUri(v string) *GetCustomAgentResponseBodyDataArtifactCode {
	s.SourceUri = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataArtifactCode) Validate() error {
	return dara.Validate(s)
}

type GetCustomAgentResponseBodyDataArtifactContainer struct {
	// The custom executable program and startup parameters. If this parameter is not configured, the default startup method of the image is used.
	//
	// example:
	//
	// python app.py
	Entrypoint *string `json:"entrypoint,omitempty" xml:"entrypoint,omitempty"`
	// The ID of the ACR instance to which the image repository belongs.
	//
	// example:
	//
	// cri-a1b2c3d4e5f6
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The namespace to which the ACR image repository belongs.
	//
	// example:
	//
	// my-agent
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The name of the image repository in the ACR namespace.
	//
	// example:
	//
	// chat-agent
	Repo *string `json:"repo,omitempty" xml:"repo,omitempty"`
	// The latest submitted image tag.
	//
	// example:
	//
	// v1.0.0
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// The image source type. A value of acr indicates Alibaba Cloud Container Registry (ACR).
	//
	// This parameter is required.
	//
	// example:
	//
	// acr
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetCustomAgentResponseBodyDataArtifactContainer) String() string {
	return dara.Prettify(s)
}

func (s GetCustomAgentResponseBodyDataArtifactContainer) GoString() string {
	return s.String()
}

func (s *GetCustomAgentResponseBodyDataArtifactContainer) GetEntrypoint() *string {
	return s.Entrypoint
}

func (s *GetCustomAgentResponseBodyDataArtifactContainer) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetCustomAgentResponseBodyDataArtifactContainer) GetNamespace() *string {
	return s.Namespace
}

func (s *GetCustomAgentResponseBodyDataArtifactContainer) GetRepo() *string {
	return s.Repo
}

func (s *GetCustomAgentResponseBodyDataArtifactContainer) GetTag() *string {
	return s.Tag
}

func (s *GetCustomAgentResponseBodyDataArtifactContainer) GetType() *string {
	return s.Type
}

func (s *GetCustomAgentResponseBodyDataArtifactContainer) SetEntrypoint(v string) *GetCustomAgentResponseBodyDataArtifactContainer {
	s.Entrypoint = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataArtifactContainer) SetInstanceId(v string) *GetCustomAgentResponseBodyDataArtifactContainer {
	s.InstanceId = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataArtifactContainer) SetNamespace(v string) *GetCustomAgentResponseBodyDataArtifactContainer {
	s.Namespace = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataArtifactContainer) SetRepo(v string) *GetCustomAgentResponseBodyDataArtifactContainer {
	s.Repo = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataArtifactContainer) SetTag(v string) *GetCustomAgentResponseBodyDataArtifactContainer {
	s.Tag = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataArtifactContainer) SetType(v string) *GetCustomAgentResponseBodyDataArtifactContainer {
	s.Type = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataArtifactContainer) Validate() error {
	return dara.Validate(s)
}

type GetCustomAgentResponseBodyDataCodePackageInfo struct {
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

func (s GetCustomAgentResponseBodyDataCodePackageInfo) String() string {
	return dara.Prettify(s)
}

func (s GetCustomAgentResponseBodyDataCodePackageInfo) GoString() string {
	return s.String()
}

func (s *GetCustomAgentResponseBodyDataCodePackageInfo) GetCrc64() *string {
	return s.Crc64
}

func (s *GetCustomAgentResponseBodyDataCodePackageInfo) GetEtag() *string {
	return s.Etag
}

func (s *GetCustomAgentResponseBodyDataCodePackageInfo) GetFileName() *string {
	return s.FileName
}

func (s *GetCustomAgentResponseBodyDataCodePackageInfo) GetFileSize() *int64 {
	return s.FileSize
}

func (s *GetCustomAgentResponseBodyDataCodePackageInfo) GetFormat() *string {
	return s.Format
}

func (s *GetCustomAgentResponseBodyDataCodePackageInfo) SetCrc64(v string) *GetCustomAgentResponseBodyDataCodePackageInfo {
	s.Crc64 = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataCodePackageInfo) SetEtag(v string) *GetCustomAgentResponseBodyDataCodePackageInfo {
	s.Etag = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataCodePackageInfo) SetFileName(v string) *GetCustomAgentResponseBodyDataCodePackageInfo {
	s.FileName = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataCodePackageInfo) SetFileSize(v int64) *GetCustomAgentResponseBodyDataCodePackageInfo {
	s.FileSize = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataCodePackageInfo) SetFormat(v string) *GetCustomAgentResponseBodyDataCodePackageInfo {
	s.Format = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataCodePackageInfo) Validate() error {
	return dara.Validate(s)
}

type GetCustomAgentResponseBodyDataCurrentImage struct {
	// The original image address for the custom source.
	//
	// example:
	//
	// registry.example.com/team/agent:v1
	Image *string `json:"image,omitempty" xml:"image,omitempty"`
	// The instance ID for the ACR source.
	//
	// example:
	//
	// cri-example
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The namespace for the ACR source.
	//
	// example:
	//
	// team
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The independent pull network for the custom source.
	PullNetwork *GetCustomAgentResponseBodyDataCurrentImagePullNetwork `json:"pullNetwork,omitempty" xml:"pullNetwork,omitempty" type:"Struct"`
	// The repository name for the ACR source.
	//
	// example:
	//
	// agent
	Repo *string `json:"repo,omitempty" xml:"repo,omitempty"`
	// The certificate exception option for the custom source.
	//
	// example:
	//
	// false
	SkipTlsVerify *bool `json:"skipTlsVerify,omitempty" xml:"skipTlsVerify,omitempty"`
	// The image source type of the current active version.
	//
	// example:
	//
	// custom
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// The tag of the ACR source.
	//
	// example:
	//
	// v1
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// The source snapshot from the current active configuration version. Omitted for anonymous sources. A new source failure does not overwrite this username, and no password is returned.
	//
	// example:
	//
	// reader
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s GetCustomAgentResponseBodyDataCurrentImage) String() string {
	return dara.Prettify(s)
}

func (s GetCustomAgentResponseBodyDataCurrentImage) GoString() string {
	return s.String()
}

func (s *GetCustomAgentResponseBodyDataCurrentImage) GetImage() *string {
	return s.Image
}

func (s *GetCustomAgentResponseBodyDataCurrentImage) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetCustomAgentResponseBodyDataCurrentImage) GetNamespace() *string {
	return s.Namespace
}

func (s *GetCustomAgentResponseBodyDataCurrentImage) GetPullNetwork() *GetCustomAgentResponseBodyDataCurrentImagePullNetwork {
	return s.PullNetwork
}

func (s *GetCustomAgentResponseBodyDataCurrentImage) GetRepo() *string {
	return s.Repo
}

func (s *GetCustomAgentResponseBodyDataCurrentImage) GetSkipTlsVerify() *bool {
	return s.SkipTlsVerify
}

func (s *GetCustomAgentResponseBodyDataCurrentImage) GetSourceType() *string {
	return s.SourceType
}

func (s *GetCustomAgentResponseBodyDataCurrentImage) GetTag() *string {
	return s.Tag
}

func (s *GetCustomAgentResponseBodyDataCurrentImage) GetUsername() *string {
	return s.Username
}

func (s *GetCustomAgentResponseBodyDataCurrentImage) SetImage(v string) *GetCustomAgentResponseBodyDataCurrentImage {
	s.Image = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataCurrentImage) SetInstanceId(v string) *GetCustomAgentResponseBodyDataCurrentImage {
	s.InstanceId = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataCurrentImage) SetNamespace(v string) *GetCustomAgentResponseBodyDataCurrentImage {
	s.Namespace = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataCurrentImage) SetPullNetwork(v *GetCustomAgentResponseBodyDataCurrentImagePullNetwork) *GetCustomAgentResponseBodyDataCurrentImage {
	s.PullNetwork = v
	return s
}

func (s *GetCustomAgentResponseBodyDataCurrentImage) SetRepo(v string) *GetCustomAgentResponseBodyDataCurrentImage {
	s.Repo = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataCurrentImage) SetSkipTlsVerify(v bool) *GetCustomAgentResponseBodyDataCurrentImage {
	s.SkipTlsVerify = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataCurrentImage) SetSourceType(v string) *GetCustomAgentResponseBodyDataCurrentImage {
	s.SourceType = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataCurrentImage) SetTag(v string) *GetCustomAgentResponseBodyDataCurrentImage {
	s.Tag = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataCurrentImage) SetUsername(v string) *GetCustomAgentResponseBodyDataCurrentImage {
	s.Username = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataCurrentImage) Validate() error {
	if s.PullNetwork != nil {
		if err := s.PullNetwork.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCustomAgentResponseBodyDataCurrentImagePullNetwork struct {
	// The security group ID. This parameter is required when type is set to vpc and must belong to the selected VPC.
	//
	// example:
	//
	// sg-example
	SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	// The pull network type. This parameter is required when pullNetwork is submitted and is verified by the backend. The default value for a new creation is public network, which is independent of the agent runtime network.
	//
	// example:
	//
	// public
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The security group ID. This parameter is required when type is set to vpc and must belong to the selected VPC.
	//
	// example:
	//
	// vsw-example
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	// The VPC ID. This parameter is required when type is set to vpc. The associated account and region are verified by the backend.
	//
	// example:
	//
	// vpc-example
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s GetCustomAgentResponseBodyDataCurrentImagePullNetwork) String() string {
	return dara.Prettify(s)
}

func (s GetCustomAgentResponseBodyDataCurrentImagePullNetwork) GoString() string {
	return s.String()
}

func (s *GetCustomAgentResponseBodyDataCurrentImagePullNetwork) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *GetCustomAgentResponseBodyDataCurrentImagePullNetwork) GetType() *string {
	return s.Type
}

func (s *GetCustomAgentResponseBodyDataCurrentImagePullNetwork) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GetCustomAgentResponseBodyDataCurrentImagePullNetwork) GetVpcId() *string {
	return s.VpcId
}

func (s *GetCustomAgentResponseBodyDataCurrentImagePullNetwork) SetSecurityGroupId(v string) *GetCustomAgentResponseBodyDataCurrentImagePullNetwork {
	s.SecurityGroupId = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataCurrentImagePullNetwork) SetType(v string) *GetCustomAgentResponseBodyDataCurrentImagePullNetwork {
	s.Type = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataCurrentImagePullNetwork) SetVSwitchId(v string) *GetCustomAgentResponseBodyDataCurrentImagePullNetwork {
	s.VSwitchId = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataCurrentImagePullNetwork) SetVpcId(v string) *GetCustomAgentResponseBodyDataCurrentImagePullNetwork {
	s.VpcId = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataCurrentImagePullNetwork) Validate() error {
	return dara.Validate(s)
}

type GetCustomAgentResponseBodyDataCustomImage struct {
	// The latest saved source address, which may still be in preparation.
	//
	// example:
	//
	// registry.example.com/team/agent:v1
	Image *string `json:"image,omitempty" xml:"image,omitempty"`
	// The saved dedicated pull network.
	PullNetwork *GetCustomAgentResponseBodyDataCustomImagePullNetwork `json:"pullNetwork,omitempty" xml:"pullNetwork,omitempty" type:"Struct"`
	// The saved HTTPS certificate exception option.
	//
	// example:
	//
	// false
	SkipTlsVerify *bool `json:"skipTlsVerify,omitempty" xml:"skipTlsVerify,omitempty"`
	// Returns the saved username only after the source is successfully prepared. Omitted for anonymous or unverified sources. This does not indicate the presence of a reusable password.
	//
	// example:
	//
	// reader
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s GetCustomAgentResponseBodyDataCustomImage) String() string {
	return dara.Prettify(s)
}

func (s GetCustomAgentResponseBodyDataCustomImage) GoString() string {
	return s.String()
}

func (s *GetCustomAgentResponseBodyDataCustomImage) GetImage() *string {
	return s.Image
}

func (s *GetCustomAgentResponseBodyDataCustomImage) GetPullNetwork() *GetCustomAgentResponseBodyDataCustomImagePullNetwork {
	return s.PullNetwork
}

func (s *GetCustomAgentResponseBodyDataCustomImage) GetSkipTlsVerify() *bool {
	return s.SkipTlsVerify
}

func (s *GetCustomAgentResponseBodyDataCustomImage) GetUsername() *string {
	return s.Username
}

func (s *GetCustomAgentResponseBodyDataCustomImage) SetImage(v string) *GetCustomAgentResponseBodyDataCustomImage {
	s.Image = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataCustomImage) SetPullNetwork(v *GetCustomAgentResponseBodyDataCustomImagePullNetwork) *GetCustomAgentResponseBodyDataCustomImage {
	s.PullNetwork = v
	return s
}

func (s *GetCustomAgentResponseBodyDataCustomImage) SetSkipTlsVerify(v bool) *GetCustomAgentResponseBodyDataCustomImage {
	s.SkipTlsVerify = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataCustomImage) SetUsername(v string) *GetCustomAgentResponseBodyDataCustomImage {
	s.Username = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataCustomImage) Validate() error {
	if s.PullNetwork != nil {
		if err := s.PullNetwork.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCustomAgentResponseBodyDataCustomImagePullNetwork struct {
	// The security group ID. This parameter is required when type is set to vpc and must belong to the selected VPC.
	//
	// example:
	//
	// sg-example
	SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	// The pull network type. This parameter is required when pullNetwork is submitted and is verified by the backend. The default value for a new creation is public network, which is independent of the agent runtime network.
	//
	// example:
	//
	// public
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The security group ID. This parameter is required when type is set to vpc and must belong to the selected VPC.
	//
	// example:
	//
	// vsw-example
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	// The VPC ID. This parameter is required when type is set to vpc. The associated account and region are verified by the backend.
	//
	// example:
	//
	// vpc-example
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s GetCustomAgentResponseBodyDataCustomImagePullNetwork) String() string {
	return dara.Prettify(s)
}

func (s GetCustomAgentResponseBodyDataCustomImagePullNetwork) GoString() string {
	return s.String()
}

func (s *GetCustomAgentResponseBodyDataCustomImagePullNetwork) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *GetCustomAgentResponseBodyDataCustomImagePullNetwork) GetType() *string {
	return s.Type
}

func (s *GetCustomAgentResponseBodyDataCustomImagePullNetwork) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GetCustomAgentResponseBodyDataCustomImagePullNetwork) GetVpcId() *string {
	return s.VpcId
}

func (s *GetCustomAgentResponseBodyDataCustomImagePullNetwork) SetSecurityGroupId(v string) *GetCustomAgentResponseBodyDataCustomImagePullNetwork {
	s.SecurityGroupId = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataCustomImagePullNetwork) SetType(v string) *GetCustomAgentResponseBodyDataCustomImagePullNetwork {
	s.Type = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataCustomImagePullNetwork) SetVSwitchId(v string) *GetCustomAgentResponseBodyDataCustomImagePullNetwork {
	s.VSwitchId = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataCustomImagePullNetwork) SetVpcId(v string) *GetCustomAgentResponseBodyDataCustomImagePullNetwork {
	s.VpcId = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataCustomImagePullNetwork) Validate() error {
	return dara.Validate(s)
}

type GetCustomAgentResponseBodyDataEnvironment struct {
	// The list of recently submitted environment variables.
	Variables []*GetCustomAgentResponseBodyDataEnvironmentVariables `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
}

func (s GetCustomAgentResponseBodyDataEnvironment) String() string {
	return dara.Prettify(s)
}

func (s GetCustomAgentResponseBodyDataEnvironment) GoString() string {
	return s.String()
}

func (s *GetCustomAgentResponseBodyDataEnvironment) GetVariables() []*GetCustomAgentResponseBodyDataEnvironmentVariables {
	return s.Variables
}

func (s *GetCustomAgentResponseBodyDataEnvironment) SetVariables(v []*GetCustomAgentResponseBodyDataEnvironmentVariables) *GetCustomAgentResponseBodyDataEnvironment {
	s.Variables = v
	return s
}

func (s *GetCustomAgentResponseBodyDataEnvironment) Validate() error {
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

type GetCustomAgentResponseBodyDataEnvironmentVariables struct {
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

func (s GetCustomAgentResponseBodyDataEnvironmentVariables) String() string {
	return dara.Prettify(s)
}

func (s GetCustomAgentResponseBodyDataEnvironmentVariables) GoString() string {
	return s.String()
}

func (s *GetCustomAgentResponseBodyDataEnvironmentVariables) GetName() *string {
	return s.Name
}

func (s *GetCustomAgentResponseBodyDataEnvironmentVariables) GetValue() *string {
	return s.Value
}

func (s *GetCustomAgentResponseBodyDataEnvironmentVariables) SetName(v string) *GetCustomAgentResponseBodyDataEnvironmentVariables {
	s.Name = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataEnvironmentVariables) SetValue(v string) *GetCustomAgentResponseBodyDataEnvironmentVariables {
	s.Value = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataEnvironmentVariables) Validate() error {
	return dara.Validate(s)
}

type GetCustomAgentResponseBodyDataHealthCheck struct {
	// Deprecated
	//
	// This parameter does not take effect.
	//
	// example:
	//
	// 1
	HealthyThreshold *int32 `json:"healthyThreshold,omitempty" xml:"healthyThreshold,omitempty"`
	// The interval between two health checks, in seconds.
	//
	// example:
	//
	// 10
	IntervalSeconds *int32 `json:"intervalSeconds,omitempty" xml:"intervalSeconds,omitempty"`
	// The HTTP health check path.
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

func (s GetCustomAgentResponseBodyDataHealthCheck) String() string {
	return dara.Prettify(s)
}

func (s GetCustomAgentResponseBodyDataHealthCheck) GoString() string {
	return s.String()
}

func (s *GetCustomAgentResponseBodyDataHealthCheck) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *GetCustomAgentResponseBodyDataHealthCheck) GetIntervalSeconds() *int32 {
	return s.IntervalSeconds
}

func (s *GetCustomAgentResponseBodyDataHealthCheck) GetPath() *string {
	return s.Path
}

func (s *GetCustomAgentResponseBodyDataHealthCheck) GetPort() *int32 {
	return s.Port
}

func (s *GetCustomAgentResponseBodyDataHealthCheck) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *GetCustomAgentResponseBodyDataHealthCheck) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *GetCustomAgentResponseBodyDataHealthCheck) SetHealthyThreshold(v int32) *GetCustomAgentResponseBodyDataHealthCheck {
	s.HealthyThreshold = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataHealthCheck) SetIntervalSeconds(v int32) *GetCustomAgentResponseBodyDataHealthCheck {
	s.IntervalSeconds = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataHealthCheck) SetPath(v string) *GetCustomAgentResponseBodyDataHealthCheck {
	s.Path = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataHealthCheck) SetPort(v int32) *GetCustomAgentResponseBodyDataHealthCheck {
	s.Port = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataHealthCheck) SetTimeoutSeconds(v int32) *GetCustomAgentResponseBodyDataHealthCheck {
	s.TimeoutSeconds = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataHealthCheck) SetUnhealthyThreshold(v int32) *GetCustomAgentResponseBodyDataHealthCheck {
	s.UnhealthyThreshold = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataHealthCheck) Validate() error {
	return dara.Validate(s)
}

type GetCustomAgentResponseBodyDataImagePreparation struct {
	// The current image preparation phase. Configuration acceptance does not mean successful deployment.
	//
	// example:
	//
	// QUEUED
	Phase *string `json:"phase,omitempty" xml:"phase,omitempty"`
	// The task ID submitted from the console. This is not the Resource Management (RM) operation ID.
	//
	// example:
	//
	// agent-example:1
	PreparationId *string `json:"preparationId,omitempty" xml:"preparationId,omitempty"`
}

func (s GetCustomAgentResponseBodyDataImagePreparation) String() string {
	return dara.Prettify(s)
}

func (s GetCustomAgentResponseBodyDataImagePreparation) GoString() string {
	return s.String()
}

func (s *GetCustomAgentResponseBodyDataImagePreparation) GetPhase() *string {
	return s.Phase
}

func (s *GetCustomAgentResponseBodyDataImagePreparation) GetPreparationId() *string {
	return s.PreparationId
}

func (s *GetCustomAgentResponseBodyDataImagePreparation) SetPhase(v string) *GetCustomAgentResponseBodyDataImagePreparation {
	s.Phase = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataImagePreparation) SetPreparationId(v string) *GetCustomAgentResponseBodyDataImagePreparation {
	s.PreparationId = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataImagePreparation) Validate() error {
	return dara.Validate(s)
}

type GetCustomAgentResponseBodyDataLatestExecutionError struct {
	// The reason code for the waiting or execution failure of the latest configuration. Used to determine subsequent handling methods.
	//
	// This parameter is required.
	//
	// example:
	//
	// AgentExecutionFailed
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The reason why the latest configuration is pending or failed to execute. Determine what to do next based on the code and retryable fields.
	//
	// This parameter is required.
	//
	// example:
	//
	// Agent execution failed. Check the configuration and resubmit
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Indicates whether the task is automatically retried. If true, you can query again later. If false, check and fix the configuration based on the reason, and then resubmit.
	//
	// This parameter is required.
	Retryable *bool `json:"retryable,omitempty" xml:"retryable,omitempty"`
}

func (s GetCustomAgentResponseBodyDataLatestExecutionError) String() string {
	return dara.Prettify(s)
}

func (s GetCustomAgentResponseBodyDataLatestExecutionError) GoString() string {
	return s.String()
}

func (s *GetCustomAgentResponseBodyDataLatestExecutionError) GetCode() *string {
	return s.Code
}

func (s *GetCustomAgentResponseBodyDataLatestExecutionError) GetMessage() *string {
	return s.Message
}

func (s *GetCustomAgentResponseBodyDataLatestExecutionError) GetRetryable() *bool {
	return s.Retryable
}

func (s *GetCustomAgentResponseBodyDataLatestExecutionError) SetCode(v string) *GetCustomAgentResponseBodyDataLatestExecutionError {
	s.Code = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataLatestExecutionError) SetMessage(v string) *GetCustomAgentResponseBodyDataLatestExecutionError {
	s.Message = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataLatestExecutionError) SetRetryable(v bool) *GetCustomAgentResponseBodyDataLatestExecutionError {
	s.Retryable = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataLatestExecutionError) Validate() error {
	return dara.Validate(s)
}

type GetCustomAgentResponseBodyDataLogging struct {
	// The Simple Log Service (SLS) log delivery destination.
	Sls *GetCustomAgentResponseBodyDataLoggingSls `json:"sls,omitempty" xml:"sls,omitempty" type:"Struct"`
}

func (s GetCustomAgentResponseBodyDataLogging) String() string {
	return dara.Prettify(s)
}

func (s GetCustomAgentResponseBodyDataLogging) GoString() string {
	return s.String()
}

func (s *GetCustomAgentResponseBodyDataLogging) GetSls() *GetCustomAgentResponseBodyDataLoggingSls {
	return s.Sls
}

func (s *GetCustomAgentResponseBodyDataLogging) SetSls(v *GetCustomAgentResponseBodyDataLoggingSls) *GetCustomAgentResponseBodyDataLogging {
	s.Sls = v
	return s
}

func (s *GetCustomAgentResponseBodyDataLogging) Validate() error {
	if s.Sls != nil {
		if err := s.Sls.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCustomAgentResponseBodyDataLoggingSls struct {
	// The name of the SLS Logstore that receives agent logs.
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

func (s GetCustomAgentResponseBodyDataLoggingSls) String() string {
	return dara.Prettify(s)
}

func (s GetCustomAgentResponseBodyDataLoggingSls) GoString() string {
	return s.String()
}

func (s *GetCustomAgentResponseBodyDataLoggingSls) GetLogstore() *string {
	return s.Logstore
}

func (s *GetCustomAgentResponseBodyDataLoggingSls) GetProject() *string {
	return s.Project
}

func (s *GetCustomAgentResponseBodyDataLoggingSls) SetLogstore(v string) *GetCustomAgentResponseBodyDataLoggingSls {
	s.Logstore = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataLoggingSls) SetProject(v string) *GetCustomAgentResponseBodyDataLoggingSls {
	s.Project = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataLoggingSls) Validate() error {
	return dara.Validate(s)
}

type GetCustomAgentResponseBodyDataNetwork struct {
	// The Internet access configuration of the agent. If not configured, Internet access is allowed on the agent side but remains restricted by the workspace.
	AccessInternet *GetCustomAgentResponseBodyDataNetworkAccessInternet `json:"accessInternet,omitempty" xml:"accessInternet,omitempty" type:"Struct"`
	// The user VPC access configuration of the agent. If not configured, user VPC access is allowed on the agent side but remains restricted by the workspace.
	AccessVpc *GetCustomAgentResponseBodyDataNetworkAccessVpc `json:"accessVpc,omitempty" xml:"accessVpc,omitempty" type:"Struct"`
}

func (s GetCustomAgentResponseBodyDataNetwork) String() string {
	return dara.Prettify(s)
}

func (s GetCustomAgentResponseBodyDataNetwork) GoString() string {
	return s.String()
}

func (s *GetCustomAgentResponseBodyDataNetwork) GetAccessInternet() *GetCustomAgentResponseBodyDataNetworkAccessInternet {
	return s.AccessInternet
}

func (s *GetCustomAgentResponseBodyDataNetwork) GetAccessVpc() *GetCustomAgentResponseBodyDataNetworkAccessVpc {
	return s.AccessVpc
}

func (s *GetCustomAgentResponseBodyDataNetwork) SetAccessInternet(v *GetCustomAgentResponseBodyDataNetworkAccessInternet) *GetCustomAgentResponseBodyDataNetwork {
	s.AccessInternet = v
	return s
}

func (s *GetCustomAgentResponseBodyDataNetwork) SetAccessVpc(v *GetCustomAgentResponseBodyDataNetworkAccessVpc) *GetCustomAgentResponseBodyDataNetwork {
	s.AccessVpc = v
	return s
}

func (s *GetCustomAgentResponseBodyDataNetwork) Validate() error {
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

type GetCustomAgentResponseBodyDataNetworkAccessInternet struct {
	// Specifies whether Internet access is allowed on the agent side. The actual access remains restricted by the workspace.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s GetCustomAgentResponseBodyDataNetworkAccessInternet) String() string {
	return dara.Prettify(s)
}

func (s GetCustomAgentResponseBodyDataNetworkAccessInternet) GoString() string {
	return s.String()
}

func (s *GetCustomAgentResponseBodyDataNetworkAccessInternet) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetCustomAgentResponseBodyDataNetworkAccessInternet) SetEnabled(v bool) *GetCustomAgentResponseBodyDataNetworkAccessInternet {
	s.Enabled = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataNetworkAccessInternet) Validate() error {
	return dara.Validate(s)
}

type GetCustomAgentResponseBodyDataNetworkAccessVpc struct {
	// Specifies whether access to the user VPC configured for the workspace is allowed on the agent side. The actual access remains restricted by the workspace.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s GetCustomAgentResponseBodyDataNetworkAccessVpc) String() string {
	return dara.Prettify(s)
}

func (s GetCustomAgentResponseBodyDataNetworkAccessVpc) GoString() string {
	return s.String()
}

func (s *GetCustomAgentResponseBodyDataNetworkAccessVpc) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetCustomAgentResponseBodyDataNetworkAccessVpc) SetEnabled(v bool) *GetCustomAgentResponseBodyDataNetworkAccessVpc {
	s.Enabled = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataNetworkAccessVpc) Validate() error {
	return dara.Validate(s)
}

type GetCustomAgentResponseBodyDataObservability struct {
	// Specifies whether agent observability is enabled.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s GetCustomAgentResponseBodyDataObservability) String() string {
	return dara.Prettify(s)
}

func (s GetCustomAgentResponseBodyDataObservability) GoString() string {
	return s.String()
}

func (s *GetCustomAgentResponseBodyDataObservability) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetCustomAgentResponseBodyDataObservability) SetEnabled(v bool) *GetCustomAgentResponseBodyDataObservability {
	s.Enabled = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataObservability) Validate() error {
	return dara.Validate(s)
}

type GetCustomAgentResponseBodyDataOssMounts struct {
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
	// The object directory prefix in the mounted bucket.
	//
	// example:
	//
	// datasets/train
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// Specifies whether the mount is read-only. A value of true indicates a read-only mount. A value of false indicates a read-write mount.
	ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
}

func (s GetCustomAgentResponseBodyDataOssMounts) String() string {
	return dara.Prettify(s)
}

func (s GetCustomAgentResponseBodyDataOssMounts) GoString() string {
	return s.String()
}

func (s *GetCustomAgentResponseBodyDataOssMounts) GetBucketName() *string {
	return s.BucketName
}

func (s *GetCustomAgentResponseBodyDataOssMounts) GetMountPath() *string {
	return s.MountPath
}

func (s *GetCustomAgentResponseBodyDataOssMounts) GetPath() *string {
	return s.Path
}

func (s *GetCustomAgentResponseBodyDataOssMounts) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *GetCustomAgentResponseBodyDataOssMounts) SetBucketName(v string) *GetCustomAgentResponseBodyDataOssMounts {
	s.BucketName = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataOssMounts) SetMountPath(v string) *GetCustomAgentResponseBodyDataOssMounts {
	s.MountPath = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataOssMounts) SetPath(v string) *GetCustomAgentResponseBodyDataOssMounts {
	s.Path = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataOssMounts) SetReadOnly(v bool) *GetCustomAgentResponseBodyDataOssMounts {
	s.ReadOnly = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataOssMounts) Validate() error {
	return dara.Validate(s)
}

type GetCustomAgentResponseBodyDataProtocol struct {
	// The request header configuration for the current protocol, where the key is the name and the value is a string.
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	// The list of API paths corresponding to the current protocol.
	Paths []*GetCustomAgentResponseBodyDataProtocolPaths `json:"paths,omitempty" xml:"paths,omitempty" type:"Repeated"`
	// The invocation protocol type. Valid values: agui (AG-UI protocol), openai-chat-completion (OpenAI Chat Completions compatible protocol), and custom (custom protocol).
	//
	// This parameter is required.
	//
	// example:
	//
	// agui
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetCustomAgentResponseBodyDataProtocol) String() string {
	return dara.Prettify(s)
}

func (s GetCustomAgentResponseBodyDataProtocol) GoString() string {
	return s.String()
}

func (s *GetCustomAgentResponseBodyDataProtocol) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCustomAgentResponseBodyDataProtocol) GetPaths() []*GetCustomAgentResponseBodyDataProtocolPaths {
	return s.Paths
}

func (s *GetCustomAgentResponseBodyDataProtocol) GetType() *string {
	return s.Type
}

func (s *GetCustomAgentResponseBodyDataProtocol) SetHeaders(v map[string]*string) *GetCustomAgentResponseBodyDataProtocol {
	s.Headers = v
	return s
}

func (s *GetCustomAgentResponseBodyDataProtocol) SetPaths(v []*GetCustomAgentResponseBodyDataProtocolPaths) *GetCustomAgentResponseBodyDataProtocol {
	s.Paths = v
	return s
}

func (s *GetCustomAgentResponseBodyDataProtocol) SetType(v string) *GetCustomAgentResponseBodyDataProtocol {
	s.Type = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataProtocol) Validate() error {
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

type GetCustomAgentResponseBodyDataProtocolPaths struct {
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

func (s GetCustomAgentResponseBodyDataProtocolPaths) String() string {
	return dara.Prettify(s)
}

func (s GetCustomAgentResponseBodyDataProtocolPaths) GoString() string {
	return s.String()
}

func (s *GetCustomAgentResponseBodyDataProtocolPaths) GetDescription() *string {
	return s.Description
}

func (s *GetCustomAgentResponseBodyDataProtocolPaths) GetName() *string {
	return s.Name
}

func (s *GetCustomAgentResponseBodyDataProtocolPaths) SetDescription(v string) *GetCustomAgentResponseBodyDataProtocolPaths {
	s.Description = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataProtocolPaths) SetName(v string) *GetCustomAgentResponseBodyDataProtocolPaths {
	s.Name = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataProtocolPaths) Validate() error {
	return dara.Validate(s)
}

type GetCustomAgentResponseBodyDataRuntime struct {
	// The compute resource configuration for the agent runtime instance.
	//
	// This parameter is required.
	Compute *GetCustomAgentResponseBodyDataRuntimeCompute `json:"compute,omitempty" xml:"compute,omitempty" type:"Struct"`
	// The elastic policy configuration for the agent.
	Hpa *GetCustomAgentResponseBodyDataRuntimeHpa `json:"hpa,omitempty" xml:"hpa,omitempty" type:"Struct"`
	// The session policy configuration for the agent.
	//
	// This parameter is required.
	SessionPolicy *GetCustomAgentResponseBodyDataRuntimeSessionPolicy `json:"sessionPolicy,omitempty" xml:"sessionPolicy,omitempty" type:"Struct"`
}

func (s GetCustomAgentResponseBodyDataRuntime) String() string {
	return dara.Prettify(s)
}

func (s GetCustomAgentResponseBodyDataRuntime) GoString() string {
	return s.String()
}

func (s *GetCustomAgentResponseBodyDataRuntime) GetCompute() *GetCustomAgentResponseBodyDataRuntimeCompute {
	return s.Compute
}

func (s *GetCustomAgentResponseBodyDataRuntime) GetHpa() *GetCustomAgentResponseBodyDataRuntimeHpa {
	return s.Hpa
}

func (s *GetCustomAgentResponseBodyDataRuntime) GetSessionPolicy() *GetCustomAgentResponseBodyDataRuntimeSessionPolicy {
	return s.SessionPolicy
}

func (s *GetCustomAgentResponseBodyDataRuntime) SetCompute(v *GetCustomAgentResponseBodyDataRuntimeCompute) *GetCustomAgentResponseBodyDataRuntime {
	s.Compute = v
	return s
}

func (s *GetCustomAgentResponseBodyDataRuntime) SetHpa(v *GetCustomAgentResponseBodyDataRuntimeHpa) *GetCustomAgentResponseBodyDataRuntime {
	s.Hpa = v
	return s
}

func (s *GetCustomAgentResponseBodyDataRuntime) SetSessionPolicy(v *GetCustomAgentResponseBodyDataRuntimeSessionPolicy) *GetCustomAgentResponseBodyDataRuntime {
	s.SessionPolicy = v
	return s
}

func (s *GetCustomAgentResponseBodyDataRuntime) Validate() error {
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

type GetCustomAgentResponseBodyDataRuntimeCompute struct {
	// The compute specification. Valid values: INTERACTIVE (1 vCPU and 2 GiB of memory), STANDARD (2 vCPUs and 4 GiB of memory), and HEAVY (4 vCPUs and 8 GiB of memory).
	//
	// This parameter is required.
	//
	// example:
	//
	// STANDARD
	ComputeClass *string `json:"computeClass,omitempty" xml:"computeClass,omitempty"`
}

func (s GetCustomAgentResponseBodyDataRuntimeCompute) String() string {
	return dara.Prettify(s)
}

func (s GetCustomAgentResponseBodyDataRuntimeCompute) GoString() string {
	return s.String()
}

func (s *GetCustomAgentResponseBodyDataRuntimeCompute) GetComputeClass() *string {
	return s.ComputeClass
}

func (s *GetCustomAgentResponseBodyDataRuntimeCompute) SetComputeClass(v string) *GetCustomAgentResponseBodyDataRuntimeCompute {
	s.ComputeClass = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataRuntimeCompute) Validate() error {
	return dara.Validate(s)
}

type GetCustomAgentResponseBodyDataRuntimeHpa struct {
	// Specifies whether to enable elastic scaling. Valid values: true: Enabled. false: Disabled.
	//
	// This parameter is required.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The configured maximum number of concurrent sessions for a single runtime instance.
	//
	// example:
	//
	// 100
	MaxConcurrentSessionsPerSandbox *int32 `json:"maxConcurrentSessionsPerSandbox,omitempty" xml:"maxConcurrentSessionsPerSandbox,omitempty"`
	// The configured maximum number of runtime instances.
	//
	// example:
	//
	// 10
	MaxSandboxCount *int32 `json:"maxSandboxCount,omitempty" xml:"maxSandboxCount,omitempty"`
	// The configured minimum number of runtime instances.
	//
	// example:
	//
	// 1
	MinSandboxCount *int32 `json:"minSandboxCount,omitempty" xml:"minSandboxCount,omitempty"`
	// The configured session idle timeout limit, in seconds.
	//
	// example:
	//
	// 600
	SessionTtlSeconds *int32 `json:"sessionTtlSeconds,omitempty" xml:"sessionTtlSeconds,omitempty"`
}

func (s GetCustomAgentResponseBodyDataRuntimeHpa) String() string {
	return dara.Prettify(s)
}

func (s GetCustomAgentResponseBodyDataRuntimeHpa) GoString() string {
	return s.String()
}

func (s *GetCustomAgentResponseBodyDataRuntimeHpa) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetCustomAgentResponseBodyDataRuntimeHpa) GetMaxConcurrentSessionsPerSandbox() *int32 {
	return s.MaxConcurrentSessionsPerSandbox
}

func (s *GetCustomAgentResponseBodyDataRuntimeHpa) GetMaxSandboxCount() *int32 {
	return s.MaxSandboxCount
}

func (s *GetCustomAgentResponseBodyDataRuntimeHpa) GetMinSandboxCount() *int32 {
	return s.MinSandboxCount
}

func (s *GetCustomAgentResponseBodyDataRuntimeHpa) GetSessionTtlSeconds() *int32 {
	return s.SessionTtlSeconds
}

func (s *GetCustomAgentResponseBodyDataRuntimeHpa) SetEnabled(v bool) *GetCustomAgentResponseBodyDataRuntimeHpa {
	s.Enabled = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataRuntimeHpa) SetMaxConcurrentSessionsPerSandbox(v int32) *GetCustomAgentResponseBodyDataRuntimeHpa {
	s.MaxConcurrentSessionsPerSandbox = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataRuntimeHpa) SetMaxSandboxCount(v int32) *GetCustomAgentResponseBodyDataRuntimeHpa {
	s.MaxSandboxCount = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataRuntimeHpa) SetMinSandboxCount(v int32) *GetCustomAgentResponseBodyDataRuntimeHpa {
	s.MinSandboxCount = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataRuntimeHpa) SetSessionTtlSeconds(v int32) *GetCustomAgentResponseBodyDataRuntimeHpa {
	s.SessionTtlSeconds = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataRuntimeHpa) Validate() error {
	return dara.Validate(s)
}

type GetCustomAgentResponseBodyDataRuntimeSessionPolicy struct {
	// The name of the HTTP request header used by the session policy, applicable to ISOLATED_HEADER_FIELD.
	//
	// example:
	//
	// X-Session-Id
	HeaderName *string `json:"headerName,omitempty" xml:"headerName,omitempty"`
	// The session policy type. Valid values: AFFINITY: Session affinity. ISOLATED_HEADER_FIELD: Associates sessions by using the request header specified by headerName. ISOLATED_COOKIE is not recommended. The compatible value DISABLED in requests is returned as AFFINITY.
	//
	// This parameter is required.
	//
	// example:
	//
	// AFFINITY
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetCustomAgentResponseBodyDataRuntimeSessionPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetCustomAgentResponseBodyDataRuntimeSessionPolicy) GoString() string {
	return s.String()
}

func (s *GetCustomAgentResponseBodyDataRuntimeSessionPolicy) GetHeaderName() *string {
	return s.HeaderName
}

func (s *GetCustomAgentResponseBodyDataRuntimeSessionPolicy) GetType() *string {
	return s.Type
}

func (s *GetCustomAgentResponseBodyDataRuntimeSessionPolicy) SetHeaderName(v string) *GetCustomAgentResponseBodyDataRuntimeSessionPolicy {
	s.HeaderName = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataRuntimeSessionPolicy) SetType(v string) *GetCustomAgentResponseBodyDataRuntimeSessionPolicy {
	s.Type = &v
	return s
}

func (s *GetCustomAgentResponseBodyDataRuntimeSessionPolicy) Validate() error {
	return dara.Validate(s)
}
