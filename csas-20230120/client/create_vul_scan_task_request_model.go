// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVulScanTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTimestamp(v int64) *CreateVulScanTaskRequest
	GetEndTimestamp() *int64
	SetMatchMode(v string) *CreateVulScanTaskRequest
	GetMatchMode() *string
	SetTaskDescription(v string) *CreateVulScanTaskRequest
	GetTaskDescription() *string
	SetTaskName(v string) *CreateVulScanTaskRequest
	GetTaskName() *string
	SetUserGroupIds(v []*string) *CreateVulScanTaskRequest
	GetUserGroupIds() []*string
	SetWhitelist(v []*string) *CreateVulScanTaskRequest
	GetWhitelist() []*string
}

type CreateVulScanTaskRequest struct {
	// The task expiration time, in seconds-level UNIX timestamp. After this time is reached, endpoints no longer pull and execute this task, and incomplete scans are not continued.
	//
	// example:
	//
	// 1786291200
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The matching mode for the effective scope. Valid values:
	//
	// - **UserGroupAll**: Takes effect for all users under the current Alibaba Cloud account.
	//
	// - **UserGroupNormal**: Takes effect only for users in specified user groups. In this case, UserGroupIds is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// UserGroupNormal
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// The task description.
	//
	// example:
	//
	// Execute a vulnerability scanning on R&D department endpoints
	TaskDescription *string `json:"TaskDescription,omitempty" xml:"TaskDescription,omitempty"`
	// The task name. The name can be up to 128 characters in length and can contain Chinese characters, uppercase and lowercase letters, digits, periods (.), underscores (_), and hyphens (-). Spaces are not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// R&D Department Vulnerability Scanning
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The collection of user group IDs for which the task takes effect. This parameter is required when MatchMode is set to UserGroupNormal and cannot be specified when MatchMode is set to UserGroupAll. The collection must contain at least 1 and at most 100 entries. Duplicate values are not allowed.
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	// The list of exempt usernames. Users in this list are excluded from this scan. The list can contain up to 1000 entries. Duplicate values are not allowed.
	Whitelist []*string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
}

func (s CreateVulScanTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVulScanTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateVulScanTaskRequest) GetEndTimestamp() *int64 {
	return s.EndTimestamp
}

func (s *CreateVulScanTaskRequest) GetMatchMode() *string {
	return s.MatchMode
}

func (s *CreateVulScanTaskRequest) GetTaskDescription() *string {
	return s.TaskDescription
}

func (s *CreateVulScanTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateVulScanTaskRequest) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *CreateVulScanTaskRequest) GetWhitelist() []*string {
	return s.Whitelist
}

func (s *CreateVulScanTaskRequest) SetEndTimestamp(v int64) *CreateVulScanTaskRequest {
	s.EndTimestamp = &v
	return s
}

func (s *CreateVulScanTaskRequest) SetMatchMode(v string) *CreateVulScanTaskRequest {
	s.MatchMode = &v
	return s
}

func (s *CreateVulScanTaskRequest) SetTaskDescription(v string) *CreateVulScanTaskRequest {
	s.TaskDescription = &v
	return s
}

func (s *CreateVulScanTaskRequest) SetTaskName(v string) *CreateVulScanTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateVulScanTaskRequest) SetUserGroupIds(v []*string) *CreateVulScanTaskRequest {
	s.UserGroupIds = v
	return s
}

func (s *CreateVulScanTaskRequest) SetWhitelist(v []*string) *CreateVulScanTaskRequest {
	s.Whitelist = v
	return s
}

func (s *CreateVulScanTaskRequest) Validate() error {
	return dara.Validate(s)
}
