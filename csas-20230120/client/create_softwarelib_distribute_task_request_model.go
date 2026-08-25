// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSoftwarelibDistributeTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDevTags(v []*string) *CreateSoftwarelibDistributeTaskRequest
	GetDevTags() []*string
	SetDeviceGroupIds(v []*string) *CreateSoftwarelibDistributeTaskRequest
	GetDeviceGroupIds() []*string
	SetExecuteMode(v string) *CreateSoftwarelibDistributeTaskRequest
	GetExecuteMode() *string
	SetExecuteParameters(v string) *CreateSoftwarelibDistributeTaskRequest
	GetExecuteParameters() *string
	SetExecutePeriod(v string) *CreateSoftwarelibDistributeTaskRequest
	GetExecutePeriod() *string
	SetExpireMode(v string) *CreateSoftwarelibDistributeTaskRequest
	GetExpireMode() *string
	SetGmtExpired(v string) *CreateSoftwarelibDistributeTaskRequest
	GetGmtExpired() *string
	SetMatchMode(v string) *CreateSoftwarelibDistributeTaskRequest
	GetMatchMode() *string
	SetName(v string) *CreateSoftwarelibDistributeTaskRequest
	GetName() *string
	SetRetryTimes(v string) *CreateSoftwarelibDistributeTaskRequest
	GetRetryTimes() *string
	SetRunAsAccount(v string) *CreateSoftwarelibDistributeTaskRequest
	GetRunAsAccount() *string
	SetSoftwareId(v string) *CreateSoftwarelibDistributeTaskRequest
	GetSoftwareId() *string
	SetSoftwareName(v string) *CreateSoftwarelibDistributeTaskRequest
	GetSoftwareName() *string
	SetSupportOs(v string) *CreateSoftwarelibDistributeTaskRequest
	GetSupportOs() *string
	SetTaskType(v string) *CreateSoftwarelibDistributeTaskRequest
	GetTaskType() *string
	SetTimeout(v string) *CreateSoftwarelibDistributeTaskRequest
	GetTimeout() *string
	SetUserGroupIds(v []*string) *CreateSoftwarelibDistributeTaskRequest
	GetUserGroupIds() []*string
	SetVersionId(v string) *CreateSoftwarelibDistributeTaskRequest
	GetVersionId() *string
}

type CreateSoftwarelibDistributeTaskRequest struct {
	// The collection of terminal device IDs. Duplicate values are not allowed. Each ID must not exceed 1000 characters in length. This parameter is required when MatchMode is set to DevTagNormal. This parameter is not allowed when MatchMode is set to other values. Otherwise, the request is rejected.
	DevTags []*string `json:"DevTags,omitempty" xml:"DevTags,omitempty" type:"Repeated"`
	// The collection of device group IDs. Duplicate values are not allowed. This parameter is required when MatchMode is set to DeviceGroupNormal. This parameter is not allowed when MatchMode is set to other values. Otherwise, the request is rejected. You can call [ListDeviceGroups](~~ListDeviceGroups~~) to obtain the values.
	DeviceGroupIds []*string `json:"DeviceGroupIds,omitempty" xml:"DeviceGroupIds,omitempty" type:"Repeated"`
	// The execution mode. Valid values:
	//
	// - **Once**: immediate execution.
	//
	// - **Schedule**: scheduled execution.
	//
	// example:
	//
	// Once
	ExecuteMode *string `json:"ExecuteMode,omitempty" xml:"ExecuteMode,omitempty"`
	// The scheduling execution parameters in JSON format.
	//
	// example:
	//
	// {
	//
	//     "template_code": "software_install",
	//
	//     "software_url": "https://****",
	//
	//     "md5": "87ccfcce1956e9f2385137f690d7fb64",
	//
	//     "install_check_switch": false,
	//
	//     "software_name": "",
	//
	//     "software_name_en": "",
	//
	//     "speed_limit": 0,
	//
	//     "software_id": "softwarelib-software-18a2417da2779e65"
	//
	// }
	ExecuteParameters *string `json:"ExecuteParameters,omitempty" xml:"ExecuteParameters,omitempty"`
	// The task execution cycle in JSON format. The validType field specifies the cycle type. Valid values:
	//
	// - **Once**: one-time execution.
	//
	// - **Interval**: execution at intervals.
	//
	// - **Weekly**: weekly execution.
	//
	// example:
	//
	// {"validType":"Once"}
	ExecutePeriod *string `json:"ExecutePeriod,omitempty" xml:"ExecutePeriod,omitempty"`
	// The expiration type. Valid values:
	//
	// - **Expire**: expires at the time specified by GmtExpired.
	//
	// - **Never**: never expires.
	//
	// example:
	//
	// Expire
	ExpireMode *string `json:"ExpireMode,omitempty" xml:"ExpireMode,omitempty"`
	// The task expiration time as a millisecond-level UNIX timestamp. This parameter takes effect only when ExpireMode is set to Expire.
	//
	// example:
	//
	// 1786945543000
	GmtExpired *string `json:"GmtExpired,omitempty" xml:"GmtExpired,omitempty"`
	// The policy matching target type. Valid values:
	//
	// - **UserGroupAll**: all users.
	//
	// - **UserGroupNormal**: specified user groups.
	//
	// - **DevTagNormal**: specified devices.
	//
	// - **DeviceGroupNormal**: specified device groups.
	//
	// - **DevTagAll**: all devices.
	//
	// - **None**: not configured.
	//
	// example:
	//
	// UserGroupAll
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// The task name. The name must be 1 to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_task
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of retries after a task failure.
	//
	// example:
	//
	// 5
	RetryTimes *string `json:"RetryTimes,omitempty" xml:"RetryTimes,omitempty"`
	// The administrator account name used to run the task on Windows. The name must not exceed 128 characters in length.
	//
	// example:
	//
	// admin
	RunAsAccount *string `json:"RunAsAccount,omitempty" xml:"RunAsAccount,omitempty"`
	// The software ID in the software library. You can call [ListSoftwarelibSoftware](~~ListSoftwarelibSoftware~~) to obtain the value.
	//
	// example:
	//
	// softwarelib-software-9f9de7b5a16f****
	SoftwareId *string `json:"SoftwareId,omitempty" xml:"SoftwareId,omitempty"`
	// The software name. The name must not exceed 128 characters in length.
	//
	// example:
	//
	// test software
	SoftwareName *string `json:"SoftwareName,omitempty" xml:"SoftwareName,omitempty"`
	// The operating system to which the task applies. Only a single value is supported. Valid values:
	//
	// - **Windows**: Windows.
	//
	// - **Mac(Apple)**: macOS with Apple silicon.
	//
	// - **Mac(Intel)**: macOS with Intel processors.
	//
	// example:
	//
	// Mac(Apple)
	SupportOs *string `json:"SupportOs,omitempty" xml:"SupportOs,omitempty"`
	// The task type. Valid values:
	//
	// - **server**: a task delivered from the console.
	//
	// - **client**: a task initiated from the client.
	//
	// example:
	//
	// server
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The task execution timeout period. Unit: seconds. For example, a value of 3600 indicates 1 hour.
	//
	// example:
	//
	// 3600
	Timeout *string `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The collection of user group IDs. Duplicate values are not allowed. This parameter is required and must contain at least one value when MatchMode is set to UserGroupNormal. This parameter is not allowed when MatchMode is set to other values. Otherwise, the request is rejected. You can call [ListUserGroups](~~ListUserGroups~~) to obtain the values.
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	// The ID of the software version to distribute. You can call [ListSoftwarelibVersion](~~ListSoftwarelibVersion~~) to obtain the value.
	//
	// example:
	//
	// softwarelib-version-30925615d2e4****
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s CreateSoftwarelibDistributeTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSoftwarelibDistributeTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateSoftwarelibDistributeTaskRequest) GetDevTags() []*string {
	return s.DevTags
}

func (s *CreateSoftwarelibDistributeTaskRequest) GetDeviceGroupIds() []*string {
	return s.DeviceGroupIds
}

func (s *CreateSoftwarelibDistributeTaskRequest) GetExecuteMode() *string {
	return s.ExecuteMode
}

func (s *CreateSoftwarelibDistributeTaskRequest) GetExecuteParameters() *string {
	return s.ExecuteParameters
}

func (s *CreateSoftwarelibDistributeTaskRequest) GetExecutePeriod() *string {
	return s.ExecutePeriod
}

func (s *CreateSoftwarelibDistributeTaskRequest) GetExpireMode() *string {
	return s.ExpireMode
}

func (s *CreateSoftwarelibDistributeTaskRequest) GetGmtExpired() *string {
	return s.GmtExpired
}

func (s *CreateSoftwarelibDistributeTaskRequest) GetMatchMode() *string {
	return s.MatchMode
}

func (s *CreateSoftwarelibDistributeTaskRequest) GetName() *string {
	return s.Name
}

func (s *CreateSoftwarelibDistributeTaskRequest) GetRetryTimes() *string {
	return s.RetryTimes
}

func (s *CreateSoftwarelibDistributeTaskRequest) GetRunAsAccount() *string {
	return s.RunAsAccount
}

func (s *CreateSoftwarelibDistributeTaskRequest) GetSoftwareId() *string {
	return s.SoftwareId
}

func (s *CreateSoftwarelibDistributeTaskRequest) GetSoftwareName() *string {
	return s.SoftwareName
}

func (s *CreateSoftwarelibDistributeTaskRequest) GetSupportOs() *string {
	return s.SupportOs
}

func (s *CreateSoftwarelibDistributeTaskRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *CreateSoftwarelibDistributeTaskRequest) GetTimeout() *string {
	return s.Timeout
}

func (s *CreateSoftwarelibDistributeTaskRequest) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *CreateSoftwarelibDistributeTaskRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *CreateSoftwarelibDistributeTaskRequest) SetDevTags(v []*string) *CreateSoftwarelibDistributeTaskRequest {
	s.DevTags = v
	return s
}

func (s *CreateSoftwarelibDistributeTaskRequest) SetDeviceGroupIds(v []*string) *CreateSoftwarelibDistributeTaskRequest {
	s.DeviceGroupIds = v
	return s
}

func (s *CreateSoftwarelibDistributeTaskRequest) SetExecuteMode(v string) *CreateSoftwarelibDistributeTaskRequest {
	s.ExecuteMode = &v
	return s
}

func (s *CreateSoftwarelibDistributeTaskRequest) SetExecuteParameters(v string) *CreateSoftwarelibDistributeTaskRequest {
	s.ExecuteParameters = &v
	return s
}

func (s *CreateSoftwarelibDistributeTaskRequest) SetExecutePeriod(v string) *CreateSoftwarelibDistributeTaskRequest {
	s.ExecutePeriod = &v
	return s
}

func (s *CreateSoftwarelibDistributeTaskRequest) SetExpireMode(v string) *CreateSoftwarelibDistributeTaskRequest {
	s.ExpireMode = &v
	return s
}

func (s *CreateSoftwarelibDistributeTaskRequest) SetGmtExpired(v string) *CreateSoftwarelibDistributeTaskRequest {
	s.GmtExpired = &v
	return s
}

func (s *CreateSoftwarelibDistributeTaskRequest) SetMatchMode(v string) *CreateSoftwarelibDistributeTaskRequest {
	s.MatchMode = &v
	return s
}

func (s *CreateSoftwarelibDistributeTaskRequest) SetName(v string) *CreateSoftwarelibDistributeTaskRequest {
	s.Name = &v
	return s
}

func (s *CreateSoftwarelibDistributeTaskRequest) SetRetryTimes(v string) *CreateSoftwarelibDistributeTaskRequest {
	s.RetryTimes = &v
	return s
}

func (s *CreateSoftwarelibDistributeTaskRequest) SetRunAsAccount(v string) *CreateSoftwarelibDistributeTaskRequest {
	s.RunAsAccount = &v
	return s
}

func (s *CreateSoftwarelibDistributeTaskRequest) SetSoftwareId(v string) *CreateSoftwarelibDistributeTaskRequest {
	s.SoftwareId = &v
	return s
}

func (s *CreateSoftwarelibDistributeTaskRequest) SetSoftwareName(v string) *CreateSoftwarelibDistributeTaskRequest {
	s.SoftwareName = &v
	return s
}

func (s *CreateSoftwarelibDistributeTaskRequest) SetSupportOs(v string) *CreateSoftwarelibDistributeTaskRequest {
	s.SupportOs = &v
	return s
}

func (s *CreateSoftwarelibDistributeTaskRequest) SetTaskType(v string) *CreateSoftwarelibDistributeTaskRequest {
	s.TaskType = &v
	return s
}

func (s *CreateSoftwarelibDistributeTaskRequest) SetTimeout(v string) *CreateSoftwarelibDistributeTaskRequest {
	s.Timeout = &v
	return s
}

func (s *CreateSoftwarelibDistributeTaskRequest) SetUserGroupIds(v []*string) *CreateSoftwarelibDistributeTaskRequest {
	s.UserGroupIds = v
	return s
}

func (s *CreateSoftwarelibDistributeTaskRequest) SetVersionId(v string) *CreateSoftwarelibDistributeTaskRequest {
	s.VersionId = &v
	return s
}

func (s *CreateSoftwarelibDistributeTaskRequest) Validate() error {
	return dara.Validate(s)
}
