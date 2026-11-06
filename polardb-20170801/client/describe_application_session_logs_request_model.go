// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationSessionLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgent(v string) *DescribeApplicationSessionLogsRequest
	GetAgent() *string
	SetApplicationId(v string) *DescribeApplicationSessionLogsRequest
	GetApplicationId() *string
	SetEndTime(v string) *DescribeApplicationSessionLogsRequest
	GetEndTime() *string
	SetOwnerAccount(v string) *DescribeApplicationSessionLogsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeApplicationSessionLogsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeApplicationSessionLogsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeApplicationSessionLogsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeApplicationSessionLogsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeApplicationSessionLogsRequest
	GetResourceOwnerAccount() *string
	SetRole(v string) *DescribeApplicationSessionLogsRequest
	GetRole() *string
	SetSessionId(v string) *DescribeApplicationSessionLogsRequest
	GetSessionId() *string
	SetStartTime(v string) *DescribeApplicationSessionLogsRequest
	GetStartTime() *string
	SetType(v string) *DescribeApplicationSessionLogsRequest
	GetType() *string
}

type DescribeApplicationSessionLogsRequest struct {
	// The agent name. This parameter must be specified together with the SessionId parameter, or neither is specified.
	//
	// example:
	//
	// test-agent
	Agent *string `json:"Agent,omitempty" xml:"Agent,omitempty"`
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The end time of the query. Specify the time in the ISO 8601 standard in the YYYY-MM-DDTHH:mmZ format. The time must be in UTC. The end time cannot be earlier than the start time, and the query time span cannot exceed 30 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-11-04T02:00Z
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Pages start from page 1. Maximum value: 2147483647. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: 30 to 100. Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The role filter for logs. A value of user indicates a user, assistant indicates an assistant, and toolResult indicates a tool execution result.
	//
	// example:
	//
	// assistant
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The session ID. This parameter must be specified together with the Agent parameter, or neither is specified.
	//
	// example:
	//
	// session-123
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The start time of the query. Specify the time in the ISO 8601 standard in the YYYY-MM-DDTHH:mmZ format. The time must be in UTC. The start time cannot be later than the current time, and the interval between the start time and the current time cannot exceed 31 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-11-04T01:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The log type. This operation only queries session logs.
	//
	// example:
	//
	// session
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeApplicationSessionLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationSessionLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationSessionLogsRequest) GetAgent() *string {
	return s.Agent
}

func (s *DescribeApplicationSessionLogsRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribeApplicationSessionLogsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeApplicationSessionLogsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeApplicationSessionLogsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeApplicationSessionLogsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeApplicationSessionLogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApplicationSessionLogsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApplicationSessionLogsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeApplicationSessionLogsRequest) GetRole() *string {
	return s.Role
}

func (s *DescribeApplicationSessionLogsRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *DescribeApplicationSessionLogsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeApplicationSessionLogsRequest) GetType() *string {
	return s.Type
}

func (s *DescribeApplicationSessionLogsRequest) SetAgent(v string) *DescribeApplicationSessionLogsRequest {
	s.Agent = &v
	return s
}

func (s *DescribeApplicationSessionLogsRequest) SetApplicationId(v string) *DescribeApplicationSessionLogsRequest {
	s.ApplicationId = &v
	return s
}

func (s *DescribeApplicationSessionLogsRequest) SetEndTime(v string) *DescribeApplicationSessionLogsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeApplicationSessionLogsRequest) SetOwnerAccount(v string) *DescribeApplicationSessionLogsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeApplicationSessionLogsRequest) SetOwnerId(v int64) *DescribeApplicationSessionLogsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeApplicationSessionLogsRequest) SetPageNumber(v int32) *DescribeApplicationSessionLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApplicationSessionLogsRequest) SetPageSize(v int32) *DescribeApplicationSessionLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApplicationSessionLogsRequest) SetRegionId(v string) *DescribeApplicationSessionLogsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApplicationSessionLogsRequest) SetResourceOwnerAccount(v string) *DescribeApplicationSessionLogsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeApplicationSessionLogsRequest) SetRole(v string) *DescribeApplicationSessionLogsRequest {
	s.Role = &v
	return s
}

func (s *DescribeApplicationSessionLogsRequest) SetSessionId(v string) *DescribeApplicationSessionLogsRequest {
	s.SessionId = &v
	return s
}

func (s *DescribeApplicationSessionLogsRequest) SetStartTime(v string) *DescribeApplicationSessionLogsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeApplicationSessionLogsRequest) SetType(v string) *DescribeApplicationSessionLogsRequest {
	s.Type = &v
	return s
}

func (s *DescribeApplicationSessionLogsRequest) Validate() error {
	return dara.Validate(s)
}
