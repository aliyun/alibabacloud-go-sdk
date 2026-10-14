// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyActiveOperationTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImmediateStart(v int32) *ModifyActiveOperationTasksRequest
	GetImmediateStart() *int32
	SetOwnerAccount(v string) *ModifyActiveOperationTasksRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyActiveOperationTasksRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyActiveOperationTasksRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyActiveOperationTasksRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyActiveOperationTasksRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ModifyActiveOperationTasksRequest
	GetSecurityToken() *string
	SetSwitchTime(v string) *ModifyActiveOperationTasksRequest
	GetSwitchTime() *string
	SetTaskIds(v string) *ModifyActiveOperationTasksRequest
	GetTaskIds() *string
}

type ModifyActiveOperationTasksRequest struct {
	// Specifies whether to immediately start execution scheduling.
	//
	// - 0: No (default).
	//
	// - 1: Yes.
	//
	// >- If the value is 0, the SwitchTime parameter is required.
	//
	// >- If the value is 1, the SwitchTime parameter does not take effect. The event start time is set to the current time, and the switchover time is automatically calculated based on the new start time. The event immediately enters execution scheduling (not an immediate switchover) and begins pre-switchover preparations. The switchover is performed after the preparations are complete. You can call the DescribeActiveOperationTasks operation and obtain the preparation time from the PrepareInterval response parameter.
	//
	// example:
	//
	// 0
	ImmediateStart *int32  `json:"ImmediateStart,omitempty" xml:"ImmediateStart,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query the region information of all clusters under the target account.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The planned switchover time to set. Specify the time in the yyyy-MM-ddTHH:mm:ssZ format (UTC).
	//
	// >
	//
	// > - The time cannot be later than the deadline.
	//
	// > - You can call the DescribeActiveOperationTasks operation and obtain the deadline from the Deadline response parameter.
	//
	// example:
	//
	// 2023-04-25T06:00:00Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// The list of task IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11111,22222
	TaskIds *string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty"`
}

func (s ModifyActiveOperationTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyActiveOperationTasksRequest) GoString() string {
	return s.String()
}

func (s *ModifyActiveOperationTasksRequest) GetImmediateStart() *int32 {
	return s.ImmediateStart
}

func (s *ModifyActiveOperationTasksRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyActiveOperationTasksRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyActiveOperationTasksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyActiveOperationTasksRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyActiveOperationTasksRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyActiveOperationTasksRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyActiveOperationTasksRequest) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *ModifyActiveOperationTasksRequest) GetTaskIds() *string {
	return s.TaskIds
}

func (s *ModifyActiveOperationTasksRequest) SetImmediateStart(v int32) *ModifyActiveOperationTasksRequest {
	s.ImmediateStart = &v
	return s
}

func (s *ModifyActiveOperationTasksRequest) SetOwnerAccount(v string) *ModifyActiveOperationTasksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyActiveOperationTasksRequest) SetOwnerId(v int64) *ModifyActiveOperationTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyActiveOperationTasksRequest) SetRegionId(v string) *ModifyActiveOperationTasksRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyActiveOperationTasksRequest) SetResourceOwnerAccount(v string) *ModifyActiveOperationTasksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyActiveOperationTasksRequest) SetResourceOwnerId(v int64) *ModifyActiveOperationTasksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyActiveOperationTasksRequest) SetSecurityToken(v string) *ModifyActiveOperationTasksRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyActiveOperationTasksRequest) SetSwitchTime(v string) *ModifyActiveOperationTasksRequest {
	s.SwitchTime = &v
	return s
}

func (s *ModifyActiveOperationTasksRequest) SetTaskIds(v string) *ModifyActiveOperationTasksRequest {
	s.TaskIds = &v
	return s
}

func (s *ModifyActiveOperationTasksRequest) Validate() error {
	return dara.Validate(s)
}
