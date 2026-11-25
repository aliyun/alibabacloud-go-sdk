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
	// The IDs of the terminal devices. Duplicate IDs are not allowed. A single ID can be up to 1,000 characters in length. This parameter is required when MatchMode is set to DevTagNormal. If MatchMode is set to other values, this parameter must not be specified. Otherwise, the request is rejected.
	DevTags []*string `json:"DevTags,omitempty" xml:"DevTags,omitempty" type:"Repeated"`
	// The IDs of the device groups. Duplicate IDs are not allowed. This parameter is required when MatchMode is set to DeviceGroupNormal. If MatchMode is set to other values, this parameter must not be specified. Otherwise, the request is rejected. You can obtain the values by calling the [ListDeviceGroups](~~ListDeviceGroups~~) operation.
	DeviceGroupIds []*string `json:"DeviceGroupIds,omitempty" xml:"DeviceGroupIds,omitempty" type:"Repeated"`
	// The execution mode. Valid values:
	//
	// - **Once**: Immediate execution.
	//
	// - **Schedule**: Scheduled execution.
	//
	// example:
	//
	// Once
	ExecuteMode *string `json:"ExecuteMode,omitempty" xml:"ExecuteMode,omitempty"`
	// The scheduled execution parameters. The value is a JSON-formatted string.
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
	// The execution cycle of the task. The value is a JSON-formatted string. The validType field specifies the cycle type. Valid values:
	//
	// - **Once**: Single execution.
	//
	// - **Interval**: Interval-based execution.
	//
	// - **Weekly**: Weekly execution.
	//
	// example:
	//
	// {"validType":"Once"}
	ExecutePeriod *string `json:"ExecutePeriod,omitempty" xml:"ExecutePeriod,omitempty"`
	// The expiration type. Valid values:
	//
	// - **Expire**: The task expires at the time specified by GmtExpired.
	//
	// - **Never**: The task never expires.
	//
	// example:
	//
	// Expire
	ExpireMode *string `json:"ExpireMode,omitempty" xml:"ExpireMode,omitempty"`
	// The expiration time of the task. The value is a UNIX timestamp in milliseconds. This parameter takes effect only when ExpireMode is set to Expire.
	//
	// example:
	//
	// 1786945543000
	GmtExpired *string `json:"GmtExpired,omitempty" xml:"GmtExpired,omitempty"`
	// The target type for policy matching. Valid values:
	//
	// - **UserGroupAll**: All users.
	//
	// - **UserGroupNormal**: Specified user groups.
	//
	// - **DevTagNormal**: Specified devices.
	//
	// - **DeviceGroupNormal**: Specified device groups.
	//
	// - **DevTagAll**: All devices.
	//
	// - **None**: Not configured.
	//
	// example:
	//
	// UserGroupAll
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// The name of the task. The name must be 1 to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_task
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of retries after the task fails.
	//
	// example:
	//
	// 5
	RetryTimes *string `json:"RetryTimes,omitempty" xml:"RetryTimes,omitempty"`
	// The name of the administrator account used to execute the task in Windows. The name can be up to 128 characters in length.
	//
	// example:
	//
	// admin
	RunAsAccount *string `json:"RunAsAccount,omitempty" xml:"RunAsAccount,omitempty"`
	// The ID of the software in the software library. You can obtain the value by calling the [ListSoftwarelibSoftware](~~ListSoftwarelibSoftware~~) operation.
	//
	// example:
	//
	// softwarelib-software-9f9de7b5a16f****
	SoftwareId *string `json:"SoftwareId,omitempty" xml:"SoftwareId,omitempty"`
	// The name of the software. The name can be up to 128 characters in length.
	//
	// example:
	//
	// test software
	SoftwareName *string `json:"SoftwareName,omitempty" xml:"SoftwareName,omitempty"`
	// The operating system to which the task applies. Only a single value is supported. Valid values:
	//
	// - **Windows**: Windows.
	//
	// - **Mac(Apple)**: macOS on Apple silicon.
	//
	// - **Mac(Intel)**: macOS on Intel processors.
	//
	// example:
	//
	// Mac(Apple)
	SupportOs *string `json:"SupportOs,omitempty" xml:"SupportOs,omitempty"`
	// The type of the task. Valid values:
	//
	// - **server**: Tasks delivered from the console.
	//
	// - **client**: Tasks initiated by the client.
	//
	// example:
	//
	// server
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The timeout period for task execution. Unit: seconds. For example, a value of 3600 indicates 1 hour.
	//
	// example:
	//
	// 3600
	Timeout *string `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The IDs of the user groups. Duplicate IDs are not allowed. This parameter is required and must contain at least one ID when MatchMode is set to UserGroupNormal. If MatchMode is set to other values, this parameter must not be specified. Otherwise, the request is rejected. You can obtain the values by calling the [ListUserGroups](~~ListUserGroups~~) operation.
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	// The ID of the software version to be distributed. When you create a software version, you can obtain the value from the VersionId response parameter of the [CreateSoftwarelibVersion](~~CreateSoftwarelibVersion~~) operation. For more information about how to view and manage software versions, see [Software management](https://www.alibabacloud.com/help/en/sase/user-guide/software-management).
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
