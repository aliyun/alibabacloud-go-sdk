// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataAgentApplication interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *DataAgentApplication
	GetAgentId() *string
	SetAppId(v string) *DataAgentApplication
	GetAppId() *string
	SetAppName(v string) *DataAgentApplication
	GetAppName() *string
	SetAppUrl(v string) *DataAgentApplication
	GetAppUrl() *string
	SetApplicationExtraInfo(v string) *DataAgentApplication
	GetApplicationExtraInfo() *string
	SetApplicationRevision(v int64) *DataAgentApplication
	GetApplicationRevision() *int64
	SetAvailabilityErrorCode(v string) *DataAgentApplication
	GetAvailabilityErrorCode() *string
	SetAvailabilityErrorMessage(v string) *DataAgentApplication
	GetAvailabilityErrorMessage() *string
	SetCreatorName(v string) *DataAgentApplication
	GetCreatorName() *string
	SetCreatorUid(v string) *DataAgentApplication
	GetCreatorUid() *string
	SetDescription(v string) *DataAgentApplication
	GetDescription() *string
	SetGmtCreated(v string) *DataAgentApplication
	GetGmtCreated() *string
	SetGmtModified(v string) *DataAgentApplication
	GetGmtModified() *string
	SetLastRequest(v *DataAgentApplicationLastRequest) *DataAgentApplication
	GetLastRequest() *DataAgentApplicationLastRequest
	SetMainUid(v string) *DataAgentApplication
	GetMainUid() *string
	SetObservationStale(v bool) *DataAgentApplication
	GetObservationStale() *bool
	SetReady(v bool) *DataAgentApplication
	GetReady() *bool
	SetRegion(v string) *DataAgentApplication
	GetRegion() *string
	SetSessionId(v string) *DataAgentApplication
	GetSessionId() *string
	SetStatus(v string) *DataAgentApplication
	GetStatus() *string
	SetVSwitchId(v string) *DataAgentApplication
	GetVSwitchId() *string
	SetVpcId(v string) *DataAgentApplication
	GetVpcId() *string
	SetWorkspaceId(v string) *DataAgentApplication
	GetWorkspaceId() *string
}

type DataAgentApplication struct {
	// The ID of the currently associated Data Agent.
	//
	// example:
	//
	// avgwuxxxxxxxxhldt7el9
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The stable identifier of the application.
	//
	// example:
	//
	// ac6izw6xxxxxxxxxxx3ulya0d
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// test-app-name
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The access URL of the application. This URL is generated and returned as-is by RM when the deployment snapshot ready=true, and is not filtered again based on Status or ObservationStale. If not provided, this parameter is empty or omitted.
	//
	// example:
	//
	// https://abc123.app.pre.datapaw-app.cloud
	AppUrl *string `json:"AppUrl,omitempty" xml:"AppUrl,omitempty"`
	// The extended information of the application.
	//
	// example:
	//
	// {}
	ApplicationExtraInfo *string `json:"ApplicationExtraInfo,omitempty" xml:"ApplicationExtraInfo,omitempty"`
	// The control revision number of the application. The initial value is 0. This number increments after a control request is accepted and can be passed as ExpectedRevision.
	//
	// example:
	//
	// 6
	ApplicationRevision *int64 `json:"ApplicationRevision,omitempty" xml:"ApplicationRevision,omitempty"`
	// The error code for the current service unavailability, which is expressed separately from the error code of the most recent control request. If no error occurs, this parameter is empty.
	//
	// example:
	//
	// SCRIPT_EXIT_FAILED
	AvailabilityErrorCode *string `json:"AvailabilityErrorCode,omitempty" xml:"AvailabilityErrorCode,omitempty"`
	// The summary of the cause for the current service unavailability. If no error occurs, this parameter is empty.
	//
	// example:
	//
	// stage=script, exitCode=2
	AvailabilityErrorMessage *string `json:"AvailabilityErrorMessage,omitempty" xml:"AvailabilityErrorMessage,omitempty"`
	// The name of the application creator.
	//
	// example:
	//
	// test-name
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	// The UID of the application owner.
	//
	// example:
	//
	// 30031588888885853
	CreatorUid *string `json:"CreatorUid,omitempty" xml:"CreatorUid,omitempty"`
	// The description of the application, which can be up to 250 characters in length.
	//
	// example:
	//
	// this is a test application
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the application was created.
	//
	// example:
	//
	// 2026-09-12T10:40:12.000+00:00
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The time when the application was last modified.
	//
	// example:
	//
	// 2026-09-12T10:40:12.000+00:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The most recent user control request. If no control request has been made, this parameter is empty.
	LastRequest *DataAgentApplicationLastRequest `json:"LastRequest,omitempty" xml:"LastRequest,omitempty" type:"Struct"`
	// The UID of the Alibaba Cloud account.
	//
	// example:
	//
	// 1673828888852166
	MainUid *string `json:"MainUid,omitempty" xml:"MainUid,omitempty"`
	// Specifies whether the runtime or service observation is missing or stale. If this parameter is set to true, the availability information may be outdated.
	//
	// example:
	//
	// false
	ObservationStale *bool `json:"ObservationStale,omitempty" xml:"ObservationStale,omitempty"`
	// Specifies whether the current platform service is available. This does not indicate whether the most recent control request succeeded or whether the business components are healthy.
	//
	// example:
	//
	// true
	Ready *bool `json:"Ready,omitempty" xml:"Ready,omitempty"`
	// The region.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the current or most recently associated session.
	//
	// example:
	//
	// axc3lsxxxxxxxxxdapwe
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The status of the application. Valid values:
	//
	// - REGISTERED
	//
	// - PREPARING
	//
	// - WAITING_AGREEMENT: Indicates that the system is waiting for the owner of the target workspace to sign a valid agreement. The original deployment or restart request is still being executed. After the agreement is signed, RM automatically resumes the original request.
	//
	// - SEALING
	//
	// - REVIEWING
	//
	// - DEPLOYING
	//
	// - DEPLOYED
	//
	// - DEPLOY_FAILED
	//
	// - STOPPING
	//
	// - STOPPED
	//
	// - STOP_FAILED
	//
	// - RECOVERING
	//
	// - UNAVAILABLE
	//
	// - DELETING
	//
	// - DELETE_FAILED
	//
	// - DELETED: The logical deletion terminal state. The application no longer appears in List or Describe operations.
	//
	// example:
	//
	// REGISTERED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the vSwitch used for the actual deployment of the currently running application. If no resources are currently running, this parameter is empty.
	//
	// example:
	//
	// vsw-example
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC used for the actual deployment of the currently running application. If no resources are currently running, this parameter is empty.
	//
	// example:
	//
	// vpc-example
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the workspace.
	//
	// example:
	//
	// 3700inkdc2y7zs0r37m5pika6
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DataAgentApplication) String() string {
	return dara.Prettify(s)
}

func (s DataAgentApplication) GoString() string {
	return s.String()
}

func (s *DataAgentApplication) GetAgentId() *string {
	return s.AgentId
}

func (s *DataAgentApplication) GetAppId() *string {
	return s.AppId
}

func (s *DataAgentApplication) GetAppName() *string {
	return s.AppName
}

func (s *DataAgentApplication) GetAppUrl() *string {
	return s.AppUrl
}

func (s *DataAgentApplication) GetApplicationExtraInfo() *string {
	return s.ApplicationExtraInfo
}

func (s *DataAgentApplication) GetApplicationRevision() *int64 {
	return s.ApplicationRevision
}

func (s *DataAgentApplication) GetAvailabilityErrorCode() *string {
	return s.AvailabilityErrorCode
}

func (s *DataAgentApplication) GetAvailabilityErrorMessage() *string {
	return s.AvailabilityErrorMessage
}

func (s *DataAgentApplication) GetCreatorName() *string {
	return s.CreatorName
}

func (s *DataAgentApplication) GetCreatorUid() *string {
	return s.CreatorUid
}

func (s *DataAgentApplication) GetDescription() *string {
	return s.Description
}

func (s *DataAgentApplication) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DataAgentApplication) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DataAgentApplication) GetLastRequest() *DataAgentApplicationLastRequest {
	return s.LastRequest
}

func (s *DataAgentApplication) GetMainUid() *string {
	return s.MainUid
}

func (s *DataAgentApplication) GetObservationStale() *bool {
	return s.ObservationStale
}

func (s *DataAgentApplication) GetReady() *bool {
	return s.Ready
}

func (s *DataAgentApplication) GetRegion() *string {
	return s.Region
}

func (s *DataAgentApplication) GetSessionId() *string {
	return s.SessionId
}

func (s *DataAgentApplication) GetStatus() *string {
	return s.Status
}

func (s *DataAgentApplication) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DataAgentApplication) GetVpcId() *string {
	return s.VpcId
}

func (s *DataAgentApplication) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DataAgentApplication) SetAgentId(v string) *DataAgentApplication {
	s.AgentId = &v
	return s
}

func (s *DataAgentApplication) SetAppId(v string) *DataAgentApplication {
	s.AppId = &v
	return s
}

func (s *DataAgentApplication) SetAppName(v string) *DataAgentApplication {
	s.AppName = &v
	return s
}

func (s *DataAgentApplication) SetAppUrl(v string) *DataAgentApplication {
	s.AppUrl = &v
	return s
}

func (s *DataAgentApplication) SetApplicationExtraInfo(v string) *DataAgentApplication {
	s.ApplicationExtraInfo = &v
	return s
}

func (s *DataAgentApplication) SetApplicationRevision(v int64) *DataAgentApplication {
	s.ApplicationRevision = &v
	return s
}

func (s *DataAgentApplication) SetAvailabilityErrorCode(v string) *DataAgentApplication {
	s.AvailabilityErrorCode = &v
	return s
}

func (s *DataAgentApplication) SetAvailabilityErrorMessage(v string) *DataAgentApplication {
	s.AvailabilityErrorMessage = &v
	return s
}

func (s *DataAgentApplication) SetCreatorName(v string) *DataAgentApplication {
	s.CreatorName = &v
	return s
}

func (s *DataAgentApplication) SetCreatorUid(v string) *DataAgentApplication {
	s.CreatorUid = &v
	return s
}

func (s *DataAgentApplication) SetDescription(v string) *DataAgentApplication {
	s.Description = &v
	return s
}

func (s *DataAgentApplication) SetGmtCreated(v string) *DataAgentApplication {
	s.GmtCreated = &v
	return s
}

func (s *DataAgentApplication) SetGmtModified(v string) *DataAgentApplication {
	s.GmtModified = &v
	return s
}

func (s *DataAgentApplication) SetLastRequest(v *DataAgentApplicationLastRequest) *DataAgentApplication {
	s.LastRequest = v
	return s
}

func (s *DataAgentApplication) SetMainUid(v string) *DataAgentApplication {
	s.MainUid = &v
	return s
}

func (s *DataAgentApplication) SetObservationStale(v bool) *DataAgentApplication {
	s.ObservationStale = &v
	return s
}

func (s *DataAgentApplication) SetReady(v bool) *DataAgentApplication {
	s.Ready = &v
	return s
}

func (s *DataAgentApplication) SetRegion(v string) *DataAgentApplication {
	s.Region = &v
	return s
}

func (s *DataAgentApplication) SetSessionId(v string) *DataAgentApplication {
	s.SessionId = &v
	return s
}

func (s *DataAgentApplication) SetStatus(v string) *DataAgentApplication {
	s.Status = &v
	return s
}

func (s *DataAgentApplication) SetVSwitchId(v string) *DataAgentApplication {
	s.VSwitchId = &v
	return s
}

func (s *DataAgentApplication) SetVpcId(v string) *DataAgentApplication {
	s.VpcId = &v
	return s
}

func (s *DataAgentApplication) SetWorkspaceId(v string) *DataAgentApplication {
	s.WorkspaceId = &v
	return s
}

func (s *DataAgentApplication) Validate() error {
	if s.LastRequest != nil {
		if err := s.LastRequest.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DataAgentApplicationLastRequest struct {
	// The control action. Valid values:
	//
	// - DEPLOY
	//
	// - STOP
	//
	// - RESTART
	//
	// - DELETE
	//
	// example:
	//
	// RESTART
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The error code for the failure of the most recent user control request. If no error occurs, this parameter is empty.
	//
	// example:
	//
	// SCRIPT_EXIT_FAILED
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The summary of the failure cause for the most recent user control request. If no error occurs, this parameter is empty.
	//
	// example:
	//
	// stage=script, exitCode=2
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s DataAgentApplicationLastRequest) String() string {
	return dara.Prettify(s)
}

func (s DataAgentApplicationLastRequest) GoString() string {
	return s.String()
}

func (s *DataAgentApplicationLastRequest) GetAction() *string {
	return s.Action
}

func (s *DataAgentApplicationLastRequest) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DataAgentApplicationLastRequest) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DataAgentApplicationLastRequest) SetAction(v string) *DataAgentApplicationLastRequest {
	s.Action = &v
	return s
}

func (s *DataAgentApplicationLastRequest) SetErrorCode(v string) *DataAgentApplicationLastRequest {
	s.ErrorCode = &v
	return s
}

func (s *DataAgentApplicationLastRequest) SetErrorMessage(v string) *DataAgentApplicationLastRequest {
	s.ErrorMessage = &v
	return s
}

func (s *DataAgentApplicationLastRequest) Validate() error {
	return dara.Validate(s)
}
