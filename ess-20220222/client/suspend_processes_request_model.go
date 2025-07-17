// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendProcessesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *SuspendProcessesRequest
	GetClientToken() *string
	SetOwnerId(v int64) *SuspendProcessesRequest
	GetOwnerId() *int64
	SetProcesses(v []*string) *SuspendProcessesRequest
	GetProcesses() []*string
	SetRegionId(v string) *SuspendProcessesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *SuspendProcessesRequest
	GetResourceOwnerAccount() *string
	SetScalingGroupId(v string) *SuspendProcessesRequest
	GetScalingGroupId() *string
}

type SuspendProcessesRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [Ensure idempotence](https://help.aliyun.com/document_detail/25965.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The types of processes that you want to suspend. Valid values:
	//
	// 	- ScaleIn: the scale-in process.
	//
	// 	- ScaleOut: the scale-out process.
	//
	// 	- HealthCheck: the health check process.
	//
	// 	- AlarmNotification: the process of executing an event-triggered task.
	//
	// 	- ScheduledAction: the process of executing a scheduled task.
	//
	// Presently, Auto Scaling supports suspending the five mentioned process types. In cases where more than five types are specified, Auto Scaling will automatically disregard duplicates and proceed with suspending the unique process types.
	//
	// This parameter is required.
	Processes []*string `json:"Processes,omitempty" xml:"Processes,omitempty" type:"Repeated"`
	// The region ID of the scaling group.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The ID of the scaling group.
	//
	// This parameter is required.
	//
	// example:
	//
	// asg-bp15oubotmrq11xe****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s SuspendProcessesRequest) String() string {
	return dara.Prettify(s)
}

func (s SuspendProcessesRequest) GoString() string {
	return s.String()
}

func (s *SuspendProcessesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SuspendProcessesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SuspendProcessesRequest) GetProcesses() []*string {
	return s.Processes
}

func (s *SuspendProcessesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SuspendProcessesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SuspendProcessesRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *SuspendProcessesRequest) SetClientToken(v string) *SuspendProcessesRequest {
	s.ClientToken = &v
	return s
}

func (s *SuspendProcessesRequest) SetOwnerId(v int64) *SuspendProcessesRequest {
	s.OwnerId = &v
	return s
}

func (s *SuspendProcessesRequest) SetProcesses(v []*string) *SuspendProcessesRequest {
	s.Processes = v
	return s
}

func (s *SuspendProcessesRequest) SetRegionId(v string) *SuspendProcessesRequest {
	s.RegionId = &v
	return s
}

func (s *SuspendProcessesRequest) SetResourceOwnerAccount(v string) *SuspendProcessesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SuspendProcessesRequest) SetScalingGroupId(v string) *SuspendProcessesRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *SuspendProcessesRequest) Validate() error {
	return dara.Validate(s)
}
