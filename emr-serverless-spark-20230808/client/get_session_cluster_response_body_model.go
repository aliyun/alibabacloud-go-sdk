// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSessionClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSessionClusterResponseBody
	GetRequestId() *string
	SetSessionCluster(v *GetSessionClusterResponseBodySessionCluster) *GetSessionClusterResponseBody
	GetSessionCluster() *GetSessionClusterResponseBodySessionCluster
}

type GetSessionClusterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The session object.
	SessionCluster *GetSessionClusterResponseBodySessionCluster `json:"sessionCluster,omitempty" xml:"sessionCluster,omitempty" type:"Struct"`
}

func (s GetSessionClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSessionClusterResponseBody) GoString() string {
	return s.String()
}

func (s *GetSessionClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSessionClusterResponseBody) GetSessionCluster() *GetSessionClusterResponseBodySessionCluster {
	return s.SessionCluster
}

func (s *GetSessionClusterResponseBody) SetRequestId(v string) *GetSessionClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSessionClusterResponseBody) SetSessionCluster(v *GetSessionClusterResponseBodySessionCluster) *GetSessionClusterResponseBody {
	s.SessionCluster = v
	return s
}

func (s *GetSessionClusterResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetSessionClusterResponseBodySessionCluster struct {
	// The Spark configurations.
	ApplicationConfigs []*GetSessionClusterResponseBodySessionClusterApplicationConfigs `json:"applicationConfigs,omitempty" xml:"applicationConfigs,omitempty" type:"Repeated"`
	// Indicates whether automatic startup is enabled.
	AutoStartConfiguration *GetSessionClusterResponseBodySessionClusterAutoStartConfiguration `json:"autoStartConfiguration,omitempty" xml:"autoStartConfiguration,omitempty" type:"Struct"`
	// Indicates whether automatic termination is enabled.
	AutoStopConfiguration *GetSessionClusterResponseBodySessionClusterAutoStopConfiguration `json:"autoStopConfiguration,omitempty" xml:"autoStopConfiguration,omitempty" type:"Struct"`
	ConnectionToken       *string                                                           `json:"connectionToken,omitempty" xml:"connectionToken,omitempty"`
	// The version of the Spark engine.
	//
	// example:
	//
	// esr-2.2(Java Runtime)
	DisplayReleaseVersion *string `json:"displayReleaseVersion,omitempty" xml:"displayReleaseVersion,omitempty"`
	// The domain name to which the Spark UI of the session belongs.
	//
	// example:
	//
	// your.domain.com
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The internal endpoint.
	//
	// example:
	//
	// emr-spark-gateway-cn-hangzhou-internal.data.aliyuncs.com
	DomainInner *string `json:"domainInner,omitempty" xml:"domainInner,omitempty"`
	// The ID of the job that is associated with the session.
	//
	// example:
	//
	// TSK-xxxxxxxx
	DraftId *string `json:"draftId,omitempty" xml:"draftId,omitempty"`
	// The environment ID.
	//
	// example:
	//
	// env-cpv569tlhtgndjl86t40
	EnvId *string `json:"envId,omitempty" xml:"envId,omitempty"`
	// The additional metadata of the session.
	//
	// example:
	//
	// {"extraInfoKey":"extraInfoValue"}
	Extra *string `json:"extra,omitempty" xml:"extra,omitempty"`
	// Indicates whether the Fusion engine is used for acceleration.
	//
	// example:
	//
	// false
	Fusion *bool `json:"fusion,omitempty" xml:"fusion,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2024-09-01 06:23:01
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The type of the job. This parameter is required and cannot be modified after the deployment is created. Valid values:
	//
	// 	- SQLSCRIPT
	//
	// 	- JAR
	//
	// 	- PYTHON
	//
	// example:
	//
	// SQL
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// The name of the session.
	//
	// example:
	//
	// test
	Name                  *string `json:"name,omitempty" xml:"name,omitempty"`
	PublicEndpointEnabled *bool   `json:"publicEndpointEnabled,omitempty" xml:"publicEndpointEnabled,omitempty"`
	// The queue name.
	//
	// example:
	//
	// jobName
	QueueName *string `json:"queueName,omitempty" xml:"queueName,omitempty"`
	// The version of Serverless Spark.
	//
	// example:
	//
	// esr-2.2(Java Runtime)
	ReleaseVersion *string `json:"releaseVersion,omitempty" xml:"releaseVersion,omitempty"`
	// The session ID.
	//
	// example:
	//
	// 1234abcd-12ab-34cd-56ef-1234567890ab
	SessionClusterId *string `json:"sessionClusterId,omitempty" xml:"sessionClusterId,omitempty"`
	// The start time.
	//
	// example:
	//
	// 2024-09-01 06:23:01
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The job status.
	//
	// 	- Starting
	//
	// 	- Running
	//
	// 	- Stopping
	//
	// 	- Stopped
	//
	// 	- Error
	//
	// example:
	//
	// Running
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// The reason of the job status change.
	StateChangeReason *GetSessionClusterResponseBodySessionClusterStateChangeReason `json:"stateChangeReason,omitempty" xml:"stateChangeReason,omitempty" type:"Struct"`
	// The user ID.
	//
	// example:
	//
	// jr-231231
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// The name of the account that is used to create the session.
	//
	// example:
	//
	// user1
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
	// The Spark UI of the session.
	//
	// example:
	//
	// https://spark-ui/link
	WebUI *string `json:"webUI,omitempty" xml:"webUI,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// w-1234abcd
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetSessionClusterResponseBodySessionCluster) String() string {
	return dara.Prettify(s)
}

func (s GetSessionClusterResponseBodySessionCluster) GoString() string {
	return s.String()
}

func (s *GetSessionClusterResponseBodySessionCluster) GetApplicationConfigs() []*GetSessionClusterResponseBodySessionClusterApplicationConfigs {
	return s.ApplicationConfigs
}

func (s *GetSessionClusterResponseBodySessionCluster) GetAutoStartConfiguration() *GetSessionClusterResponseBodySessionClusterAutoStartConfiguration {
	return s.AutoStartConfiguration
}

func (s *GetSessionClusterResponseBodySessionCluster) GetAutoStopConfiguration() *GetSessionClusterResponseBodySessionClusterAutoStopConfiguration {
	return s.AutoStopConfiguration
}

func (s *GetSessionClusterResponseBodySessionCluster) GetConnectionToken() *string {
	return s.ConnectionToken
}

func (s *GetSessionClusterResponseBodySessionCluster) GetDisplayReleaseVersion() *string {
	return s.DisplayReleaseVersion
}

func (s *GetSessionClusterResponseBodySessionCluster) GetDomain() *string {
	return s.Domain
}

func (s *GetSessionClusterResponseBodySessionCluster) GetDomainInner() *string {
	return s.DomainInner
}

func (s *GetSessionClusterResponseBodySessionCluster) GetDraftId() *string {
	return s.DraftId
}

func (s *GetSessionClusterResponseBodySessionCluster) GetEnvId() *string {
	return s.EnvId
}

func (s *GetSessionClusterResponseBodySessionCluster) GetExtra() *string {
	return s.Extra
}

func (s *GetSessionClusterResponseBodySessionCluster) GetFusion() *bool {
	return s.Fusion
}

func (s *GetSessionClusterResponseBodySessionCluster) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *GetSessionClusterResponseBodySessionCluster) GetKind() *string {
	return s.Kind
}

func (s *GetSessionClusterResponseBodySessionCluster) GetName() *string {
	return s.Name
}

func (s *GetSessionClusterResponseBodySessionCluster) GetPublicEndpointEnabled() *bool {
	return s.PublicEndpointEnabled
}

func (s *GetSessionClusterResponseBodySessionCluster) GetQueueName() *string {
	return s.QueueName
}

func (s *GetSessionClusterResponseBodySessionCluster) GetReleaseVersion() *string {
	return s.ReleaseVersion
}

func (s *GetSessionClusterResponseBodySessionCluster) GetSessionClusterId() *string {
	return s.SessionClusterId
}

func (s *GetSessionClusterResponseBodySessionCluster) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetSessionClusterResponseBodySessionCluster) GetState() *string {
	return s.State
}

func (s *GetSessionClusterResponseBodySessionCluster) GetStateChangeReason() *GetSessionClusterResponseBodySessionClusterStateChangeReason {
	return s.StateChangeReason
}

func (s *GetSessionClusterResponseBodySessionCluster) GetUserId() *string {
	return s.UserId
}

func (s *GetSessionClusterResponseBodySessionCluster) GetUserName() *string {
	return s.UserName
}

func (s *GetSessionClusterResponseBodySessionCluster) GetWebUI() *string {
	return s.WebUI
}

func (s *GetSessionClusterResponseBodySessionCluster) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetSessionClusterResponseBodySessionCluster) SetApplicationConfigs(v []*GetSessionClusterResponseBodySessionClusterApplicationConfigs) *GetSessionClusterResponseBodySessionCluster {
	s.ApplicationConfigs = v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetAutoStartConfiguration(v *GetSessionClusterResponseBodySessionClusterAutoStartConfiguration) *GetSessionClusterResponseBodySessionCluster {
	s.AutoStartConfiguration = v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetAutoStopConfiguration(v *GetSessionClusterResponseBodySessionClusterAutoStopConfiguration) *GetSessionClusterResponseBodySessionCluster {
	s.AutoStopConfiguration = v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetConnectionToken(v string) *GetSessionClusterResponseBodySessionCluster {
	s.ConnectionToken = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetDisplayReleaseVersion(v string) *GetSessionClusterResponseBodySessionCluster {
	s.DisplayReleaseVersion = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetDomain(v string) *GetSessionClusterResponseBodySessionCluster {
	s.Domain = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetDomainInner(v string) *GetSessionClusterResponseBodySessionCluster {
	s.DomainInner = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetDraftId(v string) *GetSessionClusterResponseBodySessionCluster {
	s.DraftId = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetEnvId(v string) *GetSessionClusterResponseBodySessionCluster {
	s.EnvId = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetExtra(v string) *GetSessionClusterResponseBodySessionCluster {
	s.Extra = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetFusion(v bool) *GetSessionClusterResponseBodySessionCluster {
	s.Fusion = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetGmtCreate(v int64) *GetSessionClusterResponseBodySessionCluster {
	s.GmtCreate = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetKind(v string) *GetSessionClusterResponseBodySessionCluster {
	s.Kind = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetName(v string) *GetSessionClusterResponseBodySessionCluster {
	s.Name = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetPublicEndpointEnabled(v bool) *GetSessionClusterResponseBodySessionCluster {
	s.PublicEndpointEnabled = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetQueueName(v string) *GetSessionClusterResponseBodySessionCluster {
	s.QueueName = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetReleaseVersion(v string) *GetSessionClusterResponseBodySessionCluster {
	s.ReleaseVersion = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetSessionClusterId(v string) *GetSessionClusterResponseBodySessionCluster {
	s.SessionClusterId = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetStartTime(v int64) *GetSessionClusterResponseBodySessionCluster {
	s.StartTime = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetState(v string) *GetSessionClusterResponseBodySessionCluster {
	s.State = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetStateChangeReason(v *GetSessionClusterResponseBodySessionClusterStateChangeReason) *GetSessionClusterResponseBodySessionCluster {
	s.StateChangeReason = v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetUserId(v string) *GetSessionClusterResponseBodySessionCluster {
	s.UserId = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetUserName(v string) *GetSessionClusterResponseBodySessionCluster {
	s.UserName = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetWebUI(v string) *GetSessionClusterResponseBodySessionCluster {
	s.WebUI = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetWorkspaceId(v string) *GetSessionClusterResponseBodySessionCluster {
	s.WorkspaceId = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) Validate() error {
	return dara.Validate(s)
}

type GetSessionClusterResponseBodySessionClusterApplicationConfigs struct {
	// The name of the configuration file.
	//
	// example:
	//
	// spark-defaults.conf
	ConfigFileName *string `json:"configFileName,omitempty" xml:"configFileName,omitempty"`
	// The key of the configuration.
	//
	// example:
	//
	// spark.app.name
	ConfigItemKey *string `json:"configItemKey,omitempty" xml:"configItemKey,omitempty"`
	// The configuration value.
	//
	// example:
	//
	// test
	ConfigItemValue *string `json:"configItemValue,omitempty" xml:"configItemValue,omitempty"`
}

func (s GetSessionClusterResponseBodySessionClusterApplicationConfigs) String() string {
	return dara.Prettify(s)
}

func (s GetSessionClusterResponseBodySessionClusterApplicationConfigs) GoString() string {
	return s.String()
}

func (s *GetSessionClusterResponseBodySessionClusterApplicationConfigs) GetConfigFileName() *string {
	return s.ConfigFileName
}

func (s *GetSessionClusterResponseBodySessionClusterApplicationConfigs) GetConfigItemKey() *string {
	return s.ConfigItemKey
}

func (s *GetSessionClusterResponseBodySessionClusterApplicationConfigs) GetConfigItemValue() *string {
	return s.ConfigItemValue
}

func (s *GetSessionClusterResponseBodySessionClusterApplicationConfigs) SetConfigFileName(v string) *GetSessionClusterResponseBodySessionClusterApplicationConfigs {
	s.ConfigFileName = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionClusterApplicationConfigs) SetConfigItemKey(v string) *GetSessionClusterResponseBodySessionClusterApplicationConfigs {
	s.ConfigItemKey = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionClusterApplicationConfigs) SetConfigItemValue(v string) *GetSessionClusterResponseBodySessionClusterApplicationConfigs {
	s.ConfigItemValue = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionClusterApplicationConfigs) Validate() error {
	return dara.Validate(s)
}

type GetSessionClusterResponseBodySessionClusterAutoStartConfiguration struct {
	// Indicates whether automatic startup is enabled.
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
}

func (s GetSessionClusterResponseBodySessionClusterAutoStartConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetSessionClusterResponseBodySessionClusterAutoStartConfiguration) GoString() string {
	return s.String()
}

func (s *GetSessionClusterResponseBodySessionClusterAutoStartConfiguration) GetEnable() *bool {
	return s.Enable
}

func (s *GetSessionClusterResponseBodySessionClusterAutoStartConfiguration) SetEnable(v bool) *GetSessionClusterResponseBodySessionClusterAutoStartConfiguration {
	s.Enable = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionClusterAutoStartConfiguration) Validate() error {
	return dara.Validate(s)
}

type GetSessionClusterResponseBodySessionClusterAutoStopConfiguration struct {
	// Indicates whether automatic termination is enabled.
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The idle timeout period. The session is automatically terminated when the idle timeout period is exceeded.
	//
	// example:
	//
	// 60
	IdleTimeoutMinutes *int32 `json:"idleTimeoutMinutes,omitempty" xml:"idleTimeoutMinutes,omitempty"`
}

func (s GetSessionClusterResponseBodySessionClusterAutoStopConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetSessionClusterResponseBodySessionClusterAutoStopConfiguration) GoString() string {
	return s.String()
}

func (s *GetSessionClusterResponseBodySessionClusterAutoStopConfiguration) GetEnable() *bool {
	return s.Enable
}

func (s *GetSessionClusterResponseBodySessionClusterAutoStopConfiguration) GetIdleTimeoutMinutes() *int32 {
	return s.IdleTimeoutMinutes
}

func (s *GetSessionClusterResponseBodySessionClusterAutoStopConfiguration) SetEnable(v bool) *GetSessionClusterResponseBodySessionClusterAutoStopConfiguration {
	s.Enable = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionClusterAutoStopConfiguration) SetIdleTimeoutMinutes(v int32) *GetSessionClusterResponseBodySessionClusterAutoStopConfiguration {
	s.IdleTimeoutMinutes = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionClusterAutoStopConfiguration) Validate() error {
	return dara.Validate(s)
}

type GetSessionClusterResponseBodySessionClusterStateChangeReason struct {
	// The status change code.
	//
	// example:
	//
	// 1000000
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The status change message.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetSessionClusterResponseBodySessionClusterStateChangeReason) String() string {
	return dara.Prettify(s)
}

func (s GetSessionClusterResponseBodySessionClusterStateChangeReason) GoString() string {
	return s.String()
}

func (s *GetSessionClusterResponseBodySessionClusterStateChangeReason) GetCode() *string {
	return s.Code
}

func (s *GetSessionClusterResponseBodySessionClusterStateChangeReason) GetMessage() *string {
	return s.Message
}

func (s *GetSessionClusterResponseBodySessionClusterStateChangeReason) SetCode(v string) *GetSessionClusterResponseBodySessionClusterStateChangeReason {
	s.Code = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionClusterStateChangeReason) SetMessage(v string) *GetSessionClusterResponseBodySessionClusterStateChangeReason {
	s.Message = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionClusterStateChangeReason) Validate() error {
	return dara.Validate(s)
}
