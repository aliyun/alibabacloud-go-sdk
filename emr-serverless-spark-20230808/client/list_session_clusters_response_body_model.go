// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSessionClustersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListSessionClustersResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListSessionClustersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListSessionClustersResponseBody
	GetRequestId() *string
	SetSessionClusters(v []*ListSessionClustersResponseBodySessionClusters) *ListSessionClustersResponseBody
	GetSessionClusters() []*ListSessionClustersResponseBodySessionClusters
	SetTotalCount(v int32) *ListSessionClustersResponseBody
	GetTotalCount() *int32
}

type ListSessionClustersResponseBody struct {
	// The maximum number of entries returned.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// A pagination token.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C89568980
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The sessions.
	SessionClusters []*ListSessionClustersResponseBodySessionClusters `json:"sessionClusters,omitempty" xml:"sessionClusters,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 200
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListSessionClustersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSessionClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListSessionClustersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSessionClustersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSessionClustersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSessionClustersResponseBody) GetSessionClusters() []*ListSessionClustersResponseBodySessionClusters {
	return s.SessionClusters
}

func (s *ListSessionClustersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSessionClustersResponseBody) SetMaxResults(v int32) *ListSessionClustersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSessionClustersResponseBody) SetNextToken(v string) *ListSessionClustersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSessionClustersResponseBody) SetRequestId(v string) *ListSessionClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSessionClustersResponseBody) SetSessionClusters(v []*ListSessionClustersResponseBodySessionClusters) *ListSessionClustersResponseBody {
	s.SessionClusters = v
	return s
}

func (s *ListSessionClustersResponseBody) SetTotalCount(v int32) *ListSessionClustersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSessionClustersResponseBody) Validate() error {
	if s.SessionClusters != nil {
		for _, item := range s.SessionClusters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSessionClustersResponseBodySessionClusters struct {
	// The session configurations, which are equivalent to the configurations of the Spark job.
	ApplicationConfigs []*ListSessionClustersResponseBodySessionClustersApplicationConfigs `json:"applicationConfigs,omitempty" xml:"applicationConfigs,omitempty" type:"Repeated"`
	// The automatic startup configurations.
	AutoStartConfiguration *ListSessionClustersResponseBodySessionClustersAutoStartConfiguration `json:"autoStartConfiguration,omitempty" xml:"autoStartConfiguration,omitempty" type:"Struct"`
	// The configurations of automatic termination.
	AutoStopConfiguration *ListSessionClustersResponseBodySessionClustersAutoStopConfiguration `json:"autoStopConfiguration,omitempty" xml:"autoStopConfiguration,omitempty" type:"Struct"`
	ConnectionToken       *string                                                              `json:"connectionToken,omitempty" xml:"connectionToken,omitempty"`
	// The version of the Spark engine.
	//
	// example:
	//
	// esr-4.0.0 (Spark 3.5.2, Scala 2.12)
	DisplayReleaseVersion *string `json:"displayReleaseVersion,omitempty" xml:"displayReleaseVersion,omitempty"`
	// The public endpoint of the Thrift server.
	//
	// example:
	//
	// emr-spark-gateway-cn-hangzhou.data.aliyun.com
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The internal endpoint of the Thrift server.
	//
	// example:
	//
	// emr-spark-gateway-cn-hangzhou-internal.data.aliyuncs.com
	DomainInner *string `json:"domainInner,omitempty" xml:"domainInner,omitempty"`
	// The ID of the job that is associated with the session.
	//
	// example:
	//
	// TSK-xxxxxxxxx
	DraftId *string `json:"draftId,omitempty" xml:"draftId,omitempty"`
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
	// 1732267598000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The session type.
	//
	// Valid values:
	//
	// 	- NOTEBOOK
	//
	// 	- THRIFT
	//
	// 	- SQL
	//
	// example:
	//
	// SQL
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// The name of the session.
	//
	// example:
	//
	// adhoc_query
	Name                  *string `json:"name,omitempty" xml:"name,omitempty"`
	PublicEndpointEnabled *bool   `json:"publicEndpointEnabled,omitempty" xml:"publicEndpointEnabled,omitempty"`
	// The name of the queue that is used to run the session.
	//
	// example:
	//
	// dev_queue
	QueueName *string `json:"queueName,omitempty" xml:"queueName,omitempty"`
	// The version of EMR Serverless Spark.
	//
	// example:
	//
	// esr-2.1
	ReleaseVersion *string `json:"releaseVersion,omitempty" xml:"releaseVersion,omitempty"`
	// The session ID.
	//
	// example:
	//
	// sc-123131
	SessionClusterId *string `json:"sessionClusterId,omitempty" xml:"sessionClusterId,omitempty"`
	// The start time.
	//
	// example:
	//
	// 1732267598000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The status of the session.
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
	// The details of the most recent status change of the session.
	StateChangeReason *ListSessionClustersResponseBodySessionClustersStateChangeReason `json:"stateChangeReason,omitempty" xml:"stateChangeReason,omitempty" type:"Struct"`
	// The user ID.
	//
	// example:
	//
	// 123131
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// The username.
	//
	// example:
	//
	// test_user
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
	// The Spark UI of the session.
	//
	// example:
	//
	// http://spark-ui-xxxx
	WebUI *string `json:"webUI,omitempty" xml:"webUI,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// w-1234abcd
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListSessionClustersResponseBodySessionClusters) String() string {
	return dara.Prettify(s)
}

func (s ListSessionClustersResponseBodySessionClusters) GoString() string {
	return s.String()
}

func (s *ListSessionClustersResponseBodySessionClusters) GetApplicationConfigs() []*ListSessionClustersResponseBodySessionClustersApplicationConfigs {
	return s.ApplicationConfigs
}

func (s *ListSessionClustersResponseBodySessionClusters) GetAutoStartConfiguration() *ListSessionClustersResponseBodySessionClustersAutoStartConfiguration {
	return s.AutoStartConfiguration
}

func (s *ListSessionClustersResponseBodySessionClusters) GetAutoStopConfiguration() *ListSessionClustersResponseBodySessionClustersAutoStopConfiguration {
	return s.AutoStopConfiguration
}

func (s *ListSessionClustersResponseBodySessionClusters) GetConnectionToken() *string {
	return s.ConnectionToken
}

func (s *ListSessionClustersResponseBodySessionClusters) GetDisplayReleaseVersion() *string {
	return s.DisplayReleaseVersion
}

func (s *ListSessionClustersResponseBodySessionClusters) GetDomain() *string {
	return s.Domain
}

func (s *ListSessionClustersResponseBodySessionClusters) GetDomainInner() *string {
	return s.DomainInner
}

func (s *ListSessionClustersResponseBodySessionClusters) GetDraftId() *string {
	return s.DraftId
}

func (s *ListSessionClustersResponseBodySessionClusters) GetExtra() *string {
	return s.Extra
}

func (s *ListSessionClustersResponseBodySessionClusters) GetFusion() *bool {
	return s.Fusion
}

func (s *ListSessionClustersResponseBodySessionClusters) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListSessionClustersResponseBodySessionClusters) GetKind() *string {
	return s.Kind
}

func (s *ListSessionClustersResponseBodySessionClusters) GetName() *string {
	return s.Name
}

func (s *ListSessionClustersResponseBodySessionClusters) GetPublicEndpointEnabled() *bool {
	return s.PublicEndpointEnabled
}

func (s *ListSessionClustersResponseBodySessionClusters) GetQueueName() *string {
	return s.QueueName
}

func (s *ListSessionClustersResponseBodySessionClusters) GetReleaseVersion() *string {
	return s.ReleaseVersion
}

func (s *ListSessionClustersResponseBodySessionClusters) GetSessionClusterId() *string {
	return s.SessionClusterId
}

func (s *ListSessionClustersResponseBodySessionClusters) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListSessionClustersResponseBodySessionClusters) GetState() *string {
	return s.State
}

func (s *ListSessionClustersResponseBodySessionClusters) GetStateChangeReason() *ListSessionClustersResponseBodySessionClustersStateChangeReason {
	return s.StateChangeReason
}

func (s *ListSessionClustersResponseBodySessionClusters) GetUserId() *string {
	return s.UserId
}

func (s *ListSessionClustersResponseBodySessionClusters) GetUserName() *string {
	return s.UserName
}

func (s *ListSessionClustersResponseBodySessionClusters) GetWebUI() *string {
	return s.WebUI
}

func (s *ListSessionClustersResponseBodySessionClusters) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListSessionClustersResponseBodySessionClusters) SetApplicationConfigs(v []*ListSessionClustersResponseBodySessionClustersApplicationConfigs) *ListSessionClustersResponseBodySessionClusters {
	s.ApplicationConfigs = v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetAutoStartConfiguration(v *ListSessionClustersResponseBodySessionClustersAutoStartConfiguration) *ListSessionClustersResponseBodySessionClusters {
	s.AutoStartConfiguration = v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetAutoStopConfiguration(v *ListSessionClustersResponseBodySessionClustersAutoStopConfiguration) *ListSessionClustersResponseBodySessionClusters {
	s.AutoStopConfiguration = v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetConnectionToken(v string) *ListSessionClustersResponseBodySessionClusters {
	s.ConnectionToken = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetDisplayReleaseVersion(v string) *ListSessionClustersResponseBodySessionClusters {
	s.DisplayReleaseVersion = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetDomain(v string) *ListSessionClustersResponseBodySessionClusters {
	s.Domain = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetDomainInner(v string) *ListSessionClustersResponseBodySessionClusters {
	s.DomainInner = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetDraftId(v string) *ListSessionClustersResponseBodySessionClusters {
	s.DraftId = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetExtra(v string) *ListSessionClustersResponseBodySessionClusters {
	s.Extra = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetFusion(v bool) *ListSessionClustersResponseBodySessionClusters {
	s.Fusion = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetGmtCreate(v int64) *ListSessionClustersResponseBodySessionClusters {
	s.GmtCreate = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetKind(v string) *ListSessionClustersResponseBodySessionClusters {
	s.Kind = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetName(v string) *ListSessionClustersResponseBodySessionClusters {
	s.Name = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetPublicEndpointEnabled(v bool) *ListSessionClustersResponseBodySessionClusters {
	s.PublicEndpointEnabled = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetQueueName(v string) *ListSessionClustersResponseBodySessionClusters {
	s.QueueName = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetReleaseVersion(v string) *ListSessionClustersResponseBodySessionClusters {
	s.ReleaseVersion = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetSessionClusterId(v string) *ListSessionClustersResponseBodySessionClusters {
	s.SessionClusterId = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetStartTime(v int64) *ListSessionClustersResponseBodySessionClusters {
	s.StartTime = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetState(v string) *ListSessionClustersResponseBodySessionClusters {
	s.State = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetStateChangeReason(v *ListSessionClustersResponseBodySessionClustersStateChangeReason) *ListSessionClustersResponseBodySessionClusters {
	s.StateChangeReason = v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetUserId(v string) *ListSessionClustersResponseBodySessionClusters {
	s.UserId = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetUserName(v string) *ListSessionClustersResponseBodySessionClusters {
	s.UserName = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetWebUI(v string) *ListSessionClustersResponseBodySessionClusters {
	s.WebUI = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetWorkspaceId(v string) *ListSessionClustersResponseBodySessionClusters {
	s.WorkspaceId = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) Validate() error {
	if s.ApplicationConfigs != nil {
		for _, item := range s.ApplicationConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.AutoStartConfiguration != nil {
		if err := s.AutoStartConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.AutoStopConfiguration != nil {
		if err := s.AutoStopConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.StateChangeReason != nil {
		if err := s.StateChangeReason.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSessionClustersResponseBodySessionClustersApplicationConfigs struct {
	// The name of the configuration file.
	//
	// example:
	//
	// spark-default.conf
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
	// test_application
	ConfigItemValue *string `json:"configItemValue,omitempty" xml:"configItemValue,omitempty"`
}

func (s ListSessionClustersResponseBodySessionClustersApplicationConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListSessionClustersResponseBodySessionClustersApplicationConfigs) GoString() string {
	return s.String()
}

func (s *ListSessionClustersResponseBodySessionClustersApplicationConfigs) GetConfigFileName() *string {
	return s.ConfigFileName
}

func (s *ListSessionClustersResponseBodySessionClustersApplicationConfigs) GetConfigItemKey() *string {
	return s.ConfigItemKey
}

func (s *ListSessionClustersResponseBodySessionClustersApplicationConfigs) GetConfigItemValue() *string {
	return s.ConfigItemValue
}

func (s *ListSessionClustersResponseBodySessionClustersApplicationConfigs) SetConfigFileName(v string) *ListSessionClustersResponseBodySessionClustersApplicationConfigs {
	s.ConfigFileName = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClustersApplicationConfigs) SetConfigItemKey(v string) *ListSessionClustersResponseBodySessionClustersApplicationConfigs {
	s.ConfigItemKey = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClustersApplicationConfigs) SetConfigItemValue(v string) *ListSessionClustersResponseBodySessionClustersApplicationConfigs {
	s.ConfigItemValue = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClustersApplicationConfigs) Validate() error {
	return dara.Validate(s)
}

type ListSessionClustersResponseBodySessionClustersAutoStartConfiguration struct {
	// Indicates whether automatic startup is enabled.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
}

func (s ListSessionClustersResponseBodySessionClustersAutoStartConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ListSessionClustersResponseBodySessionClustersAutoStartConfiguration) GoString() string {
	return s.String()
}

func (s *ListSessionClustersResponseBodySessionClustersAutoStartConfiguration) GetEnable() *bool {
	return s.Enable
}

func (s *ListSessionClustersResponseBodySessionClustersAutoStartConfiguration) SetEnable(v bool) *ListSessionClustersResponseBodySessionClustersAutoStartConfiguration {
	s.Enable = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClustersAutoStartConfiguration) Validate() error {
	return dara.Validate(s)
}

type ListSessionClustersResponseBodySessionClustersAutoStopConfiguration struct {
	// Indicates whether automatic termination is enabled.
	//
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The idle timeout period. The session is automatically terminated when the idle timeout period is exceeded.
	//
	// example:
	//
	// 45
	IdleTimeoutMinutes *int32 `json:"idleTimeoutMinutes,omitempty" xml:"idleTimeoutMinutes,omitempty"`
}

func (s ListSessionClustersResponseBodySessionClustersAutoStopConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ListSessionClustersResponseBodySessionClustersAutoStopConfiguration) GoString() string {
	return s.String()
}

func (s *ListSessionClustersResponseBodySessionClustersAutoStopConfiguration) GetEnable() *bool {
	return s.Enable
}

func (s *ListSessionClustersResponseBodySessionClustersAutoStopConfiguration) GetIdleTimeoutMinutes() *int32 {
	return s.IdleTimeoutMinutes
}

func (s *ListSessionClustersResponseBodySessionClustersAutoStopConfiguration) SetEnable(v bool) *ListSessionClustersResponseBodySessionClustersAutoStopConfiguration {
	s.Enable = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClustersAutoStopConfiguration) SetIdleTimeoutMinutes(v int32) *ListSessionClustersResponseBodySessionClustersAutoStopConfiguration {
	s.IdleTimeoutMinutes = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClustersAutoStopConfiguration) Validate() error {
	return dara.Validate(s)
}

type ListSessionClustersResponseBodySessionClustersStateChangeReason struct {
	// The status change code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The status change message.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s ListSessionClustersResponseBodySessionClustersStateChangeReason) String() string {
	return dara.Prettify(s)
}

func (s ListSessionClustersResponseBodySessionClustersStateChangeReason) GoString() string {
	return s.String()
}

func (s *ListSessionClustersResponseBodySessionClustersStateChangeReason) GetCode() *string {
	return s.Code
}

func (s *ListSessionClustersResponseBodySessionClustersStateChangeReason) GetMessage() *string {
	return s.Message
}

func (s *ListSessionClustersResponseBodySessionClustersStateChangeReason) SetCode(v string) *ListSessionClustersResponseBodySessionClustersStateChangeReason {
	s.Code = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClustersStateChangeReason) SetMessage(v string) *ListSessionClustersResponseBodySessionClustersStateChangeReason {
	s.Message = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClustersStateChangeReason) Validate() error {
	return dara.Validate(s)
}
