// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentInstanceMemoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListAgentInstanceMemoryResponseBodyData) *ListAgentInstanceMemoryResponseBody
	GetData() []*ListAgentInstanceMemoryResponseBodyData
	SetPageNumber(v int32) *ListAgentInstanceMemoryResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAgentInstanceMemoryResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListAgentInstanceMemoryResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListAgentInstanceMemoryResponseBody
	GetTotalCount() *int32
}

type ListAgentInstanceMemoryResponseBody struct {
	// The memory configurations and statistics information of agent instances on the current page.
	Data []*ListAgentInstanceMemoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The current page number. Pages start from 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page used for this query.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The unique identifier of the request, which can be used for troubleshooting.
	//
	// example:
	//
	// 01A1CEEB-CDFF-18F9-B2BF-9997600D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of instances that meet the filter conditions.
	//
	// example:
	//
	// 14
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAgentInstanceMemoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAgentInstanceMemoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgentInstanceMemoryResponseBody) GetData() []*ListAgentInstanceMemoryResponseBodyData {
	return s.Data
}

func (s *ListAgentInstanceMemoryResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAgentInstanceMemoryResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAgentInstanceMemoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgentInstanceMemoryResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAgentInstanceMemoryResponseBody) SetData(v []*ListAgentInstanceMemoryResponseBodyData) *ListAgentInstanceMemoryResponseBody {
	s.Data = v
	return s
}

func (s *ListAgentInstanceMemoryResponseBody) SetPageNumber(v int32) *ListAgentInstanceMemoryResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAgentInstanceMemoryResponseBody) SetPageSize(v int32) *ListAgentInstanceMemoryResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAgentInstanceMemoryResponseBody) SetRequestId(v string) *ListAgentInstanceMemoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgentInstanceMemoryResponseBody) SetTotalCount(v int32) *ListAgentInstanceMemoryResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAgentInstanceMemoryResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAgentInstanceMemoryResponseBodyData struct {
	// The Agent instance ID, which can be used to set the instance memory switch or clear the instance memory.
	//
	// example:
	//
	// agi-08vm****
	AgentInstanceId *string `json:"AgentInstanceId,omitempty" xml:"AgentInstanceId,omitempty"`
	// The lifecycle status of the agent instance.
	//
	// - `UNDEPLOYED`: Undeployed.
	//
	// - `DEPLOYING`: Deploying.
	//
	// - `DEPLOYED`: Deployed.
	//
	// - `STARTING`: Starting.
	//
	// - `RUNNING`: Running.
	//
	// - `STOPPING`: Stopping.
	//
	// - `STOPPED`: Stopped.
	//
	// - `UNAVAILABLE`: Unavailable.
	//
	// - `FAILED`: Failed.
	//
	// - `UPGRADING`: Upgrading.
	//
	// - `REBUILDING`: Resetting.
	//
	// - `DELETING`: Deleting.
	//
	// - `DELETED`: Deleted.
	//
	// example:
	//
	// RUNNING
	AgentInstanceStatus *string `json:"AgentInstanceStatus,omitempty" xml:"AgentInstanceStatus,omitempty"`
	// The business platform to which the Agent belongs.
	//
	// example:
	//
	// ENTERPRISE_AGENTIC_COMPUTER
	AgentPlatform *string `json:"AgentPlatform,omitempty" xml:"AgentPlatform,omitempty"`
	// The Agent provider.
	//
	// example:
	//
	// AgenticComputer
	AgentProvider *string `json:"AgentProvider,omitempty" xml:"AgentProvider,omitempty"`
	// The number of undeleted memory backups for the instance.
	//
	// example:
	//
	// 0
	BackupCount *int32 `json:"BackupCount,omitempty" xml:"BackupCount,omitempty"`
	// The ID of the end user to which the instance belongs.
	//
	// example:
	//
	// user01
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The time of the last memory update. The time is returned in the YYYY-MM-DDThh:mm:ss format, which may include fractional seconds. The time value uses the default time zone of the server, and the string does not contain a time zone offset. This field is omitted if the last update time is unavailable.
	//
	// example:
	//
	// 2026-10-26T21:14:36.693
	LastMemoryUpdatedAt *string `json:"LastMemoryUpdatedAt,omitempty" xml:"LastMemoryUpdatedAt,omitempty"`
	// The time of the latest backup. The time is returned in the YYYY-MM-DDThh:mm:ss format, which may include fractional seconds. The time value uses the default time zone of the server, and the string does not contain a time zone offset. This field is omitted if the backup time is unavailable.
	//
	// example:
	//
	// 2026-10-26T20:58:00.536
	LatestBackupAt *string `json:"LatestBackupAt,omitempty" xml:"LatestBackupAt,omitempty"`
	// The number of memory entries for the instance. This field may be omitted if statistics are unavailable.
	//
	// example:
	//
	// 3
	MemoryCount *int64 `json:"MemoryCount,omitempty" xml:"MemoryCount,omitempty"`
	// The memory switch configured for the instance, returned as a string. A value of `true` indicates that the switch is turned on. A value of `false` indicates that the switch is turned off.
	//
	// example:
	//
	// true
	MemoryEnabled *string `json:"MemoryEnabled,omitempty" xml:"MemoryEnabled,omitempty"`
	// The ID of the region where the runtime is located.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource associated with the instance.
	//
	// example:
	//
	// ecd-d8ut****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The runtime ID of the instance.
	//
	// example:
	//
	// ecd-d8ut****
	RuntimeId *string `json:"RuntimeId,omitempty" xml:"RuntimeId,omitempty"`
	// The runtime name. The associated resource name is returned in preference. If the resource name is unavailable, the instance name is returned.
	//
	// example:
	//
	// Office Assistant
	RuntimeName *string `json:"RuntimeName,omitempty" xml:"RuntimeName,omitempty"`
	// The list of user groups to which the end user belongs.
	UserGroups []*ListAgentInstanceMemoryResponseBodyDataUserGroups `json:"UserGroups,omitempty" xml:"UserGroups,omitempty" type:"Repeated"`
	// The username of the end user to which the instance belongs.
	//
	// example:
	//
	// user01
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The list of organizations to which the end user belongs.
	UserOrgs []*ListAgentInstanceMemoryResponseBodyDataUserOrgs `json:"UserOrgs,omitempty" xml:"UserOrgs,omitempty" type:"Repeated"`
	// The Wuying user ID of the end user.
	//
	// example:
	//
	// 0c614d26****
	WyId *string `json:"WyId,omitempty" xml:"WyId,omitempty"`
}

func (s ListAgentInstanceMemoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAgentInstanceMemoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAgentInstanceMemoryResponseBodyData) GetAgentInstanceId() *string {
	return s.AgentInstanceId
}

func (s *ListAgentInstanceMemoryResponseBodyData) GetAgentInstanceStatus() *string {
	return s.AgentInstanceStatus
}

func (s *ListAgentInstanceMemoryResponseBodyData) GetAgentPlatform() *string {
	return s.AgentPlatform
}

func (s *ListAgentInstanceMemoryResponseBodyData) GetAgentProvider() *string {
	return s.AgentProvider
}

func (s *ListAgentInstanceMemoryResponseBodyData) GetBackupCount() *int32 {
	return s.BackupCount
}

func (s *ListAgentInstanceMemoryResponseBodyData) GetEndUserId() *string {
	return s.EndUserId
}

func (s *ListAgentInstanceMemoryResponseBodyData) GetLastMemoryUpdatedAt() *string {
	return s.LastMemoryUpdatedAt
}

func (s *ListAgentInstanceMemoryResponseBodyData) GetLatestBackupAt() *string {
	return s.LatestBackupAt
}

func (s *ListAgentInstanceMemoryResponseBodyData) GetMemoryCount() *int64 {
	return s.MemoryCount
}

func (s *ListAgentInstanceMemoryResponseBodyData) GetMemoryEnabled() *string {
	return s.MemoryEnabled
}

func (s *ListAgentInstanceMemoryResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAgentInstanceMemoryResponseBodyData) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListAgentInstanceMemoryResponseBodyData) GetRuntimeId() *string {
	return s.RuntimeId
}

func (s *ListAgentInstanceMemoryResponseBodyData) GetRuntimeName() *string {
	return s.RuntimeName
}

func (s *ListAgentInstanceMemoryResponseBodyData) GetUserGroups() []*ListAgentInstanceMemoryResponseBodyDataUserGroups {
	return s.UserGroups
}

func (s *ListAgentInstanceMemoryResponseBodyData) GetUserName() *string {
	return s.UserName
}

func (s *ListAgentInstanceMemoryResponseBodyData) GetUserOrgs() []*ListAgentInstanceMemoryResponseBodyDataUserOrgs {
	return s.UserOrgs
}

func (s *ListAgentInstanceMemoryResponseBodyData) GetWyId() *string {
	return s.WyId
}

func (s *ListAgentInstanceMemoryResponseBodyData) SetAgentInstanceId(v string) *ListAgentInstanceMemoryResponseBodyData {
	s.AgentInstanceId = &v
	return s
}

func (s *ListAgentInstanceMemoryResponseBodyData) SetAgentInstanceStatus(v string) *ListAgentInstanceMemoryResponseBodyData {
	s.AgentInstanceStatus = &v
	return s
}

func (s *ListAgentInstanceMemoryResponseBodyData) SetAgentPlatform(v string) *ListAgentInstanceMemoryResponseBodyData {
	s.AgentPlatform = &v
	return s
}

func (s *ListAgentInstanceMemoryResponseBodyData) SetAgentProvider(v string) *ListAgentInstanceMemoryResponseBodyData {
	s.AgentProvider = &v
	return s
}

func (s *ListAgentInstanceMemoryResponseBodyData) SetBackupCount(v int32) *ListAgentInstanceMemoryResponseBodyData {
	s.BackupCount = &v
	return s
}

func (s *ListAgentInstanceMemoryResponseBodyData) SetEndUserId(v string) *ListAgentInstanceMemoryResponseBodyData {
	s.EndUserId = &v
	return s
}

func (s *ListAgentInstanceMemoryResponseBodyData) SetLastMemoryUpdatedAt(v string) *ListAgentInstanceMemoryResponseBodyData {
	s.LastMemoryUpdatedAt = &v
	return s
}

func (s *ListAgentInstanceMemoryResponseBodyData) SetLatestBackupAt(v string) *ListAgentInstanceMemoryResponseBodyData {
	s.LatestBackupAt = &v
	return s
}

func (s *ListAgentInstanceMemoryResponseBodyData) SetMemoryCount(v int64) *ListAgentInstanceMemoryResponseBodyData {
	s.MemoryCount = &v
	return s
}

func (s *ListAgentInstanceMemoryResponseBodyData) SetMemoryEnabled(v string) *ListAgentInstanceMemoryResponseBodyData {
	s.MemoryEnabled = &v
	return s
}

func (s *ListAgentInstanceMemoryResponseBodyData) SetRegionId(v string) *ListAgentInstanceMemoryResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *ListAgentInstanceMemoryResponseBodyData) SetResourceId(v string) *ListAgentInstanceMemoryResponseBodyData {
	s.ResourceId = &v
	return s
}

func (s *ListAgentInstanceMemoryResponseBodyData) SetRuntimeId(v string) *ListAgentInstanceMemoryResponseBodyData {
	s.RuntimeId = &v
	return s
}

func (s *ListAgentInstanceMemoryResponseBodyData) SetRuntimeName(v string) *ListAgentInstanceMemoryResponseBodyData {
	s.RuntimeName = &v
	return s
}

func (s *ListAgentInstanceMemoryResponseBodyData) SetUserGroups(v []*ListAgentInstanceMemoryResponseBodyDataUserGroups) *ListAgentInstanceMemoryResponseBodyData {
	s.UserGroups = v
	return s
}

func (s *ListAgentInstanceMemoryResponseBodyData) SetUserName(v string) *ListAgentInstanceMemoryResponseBodyData {
	s.UserName = &v
	return s
}

func (s *ListAgentInstanceMemoryResponseBodyData) SetUserOrgs(v []*ListAgentInstanceMemoryResponseBodyDataUserOrgs) *ListAgentInstanceMemoryResponseBodyData {
	s.UserOrgs = v
	return s
}

func (s *ListAgentInstanceMemoryResponseBodyData) SetWyId(v string) *ListAgentInstanceMemoryResponseBodyData {
	s.WyId = &v
	return s
}

func (s *ListAgentInstanceMemoryResponseBodyData) Validate() error {
	if s.UserGroups != nil {
		for _, item := range s.UserGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UserOrgs != nil {
		for _, item := range s.UserOrgs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAgentInstanceMemoryResponseBodyDataUserGroups struct {
	// The user group ID.
	//
	// example:
	//
	// group-aliy****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The user group name.
	//
	// example:
	//
	// Office user group
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s ListAgentInstanceMemoryResponseBodyDataUserGroups) String() string {
	return dara.Prettify(s)
}

func (s ListAgentInstanceMemoryResponseBodyDataUserGroups) GoString() string {
	return s.String()
}

func (s *ListAgentInstanceMemoryResponseBodyDataUserGroups) GetGroupId() *string {
	return s.GroupId
}

func (s *ListAgentInstanceMemoryResponseBodyDataUserGroups) GetGroupName() *string {
	return s.GroupName
}

func (s *ListAgentInstanceMemoryResponseBodyDataUserGroups) SetGroupId(v string) *ListAgentInstanceMemoryResponseBodyDataUserGroups {
	s.GroupId = &v
	return s
}

func (s *ListAgentInstanceMemoryResponseBodyDataUserGroups) SetGroupName(v string) *ListAgentInstanceMemoryResponseBodyDataUserGroups {
	s.GroupName = &v
	return s
}

func (s *ListAgentInstanceMemoryResponseBodyDataUserGroups) Validate() error {
	return dara.Validate(s)
}

type ListAgentInstanceMemoryResponseBodyDataUserOrgs struct {
	// The organization ID.
	//
	// example:
	//
	// org-aliy****
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// The organization name.
	//
	// example:
	//
	// Enterprise office
	OrgName *string `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
}

func (s ListAgentInstanceMemoryResponseBodyDataUserOrgs) String() string {
	return dara.Prettify(s)
}

func (s ListAgentInstanceMemoryResponseBodyDataUserOrgs) GoString() string {
	return s.String()
}

func (s *ListAgentInstanceMemoryResponseBodyDataUserOrgs) GetOrgId() *string {
	return s.OrgId
}

func (s *ListAgentInstanceMemoryResponseBodyDataUserOrgs) GetOrgName() *string {
	return s.OrgName
}

func (s *ListAgentInstanceMemoryResponseBodyDataUserOrgs) SetOrgId(v string) *ListAgentInstanceMemoryResponseBodyDataUserOrgs {
	s.OrgId = &v
	return s
}

func (s *ListAgentInstanceMemoryResponseBodyDataUserOrgs) SetOrgName(v string) *ListAgentInstanceMemoryResponseBodyDataUserOrgs {
	s.OrgName = &v
	return s
}

func (s *ListAgentInstanceMemoryResponseBodyDataUserOrgs) Validate() error {
	return dara.Validate(s)
}
