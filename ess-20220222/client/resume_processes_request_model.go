// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeProcessesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ResumeProcessesRequest
	GetClientToken() *string
	SetOwnerId(v int64) *ResumeProcessesRequest
	GetOwnerId() *int64
	SetProcesses(v []*string) *ResumeProcessesRequest
	GetProcesses() []*string
	SetRegionId(v string) *ResumeProcessesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ResumeProcessesRequest
	GetResourceOwnerAccount() *string
	SetScalingGroupId(v string) *ResumeProcessesRequest
	GetScalingGroupId() *string
}

type ResumeProcessesRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that the value is unique among different requests.
	//
	// The token can only contain ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25965.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Details of the processes that you want to resume.
	//
	// This parameter is required.
	Processes []*string `json:"Processes,omitempty" xml:"Processes,omitempty" type:"Repeated"`
	// The region ID of the scaling group.
	//
	// example:
	//
	// cn-qingdao
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

func (s ResumeProcessesRequest) String() string {
	return dara.Prettify(s)
}

func (s ResumeProcessesRequest) GoString() string {
	return s.String()
}

func (s *ResumeProcessesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ResumeProcessesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ResumeProcessesRequest) GetProcesses() []*string {
	return s.Processes
}

func (s *ResumeProcessesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ResumeProcessesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ResumeProcessesRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *ResumeProcessesRequest) SetClientToken(v string) *ResumeProcessesRequest {
	s.ClientToken = &v
	return s
}

func (s *ResumeProcessesRequest) SetOwnerId(v int64) *ResumeProcessesRequest {
	s.OwnerId = &v
	return s
}

func (s *ResumeProcessesRequest) SetProcesses(v []*string) *ResumeProcessesRequest {
	s.Processes = v
	return s
}

func (s *ResumeProcessesRequest) SetRegionId(v string) *ResumeProcessesRequest {
	s.RegionId = &v
	return s
}

func (s *ResumeProcessesRequest) SetResourceOwnerAccount(v string) *ResumeProcessesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ResumeProcessesRequest) SetScalingGroupId(v string) *ResumeProcessesRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *ResumeProcessesRequest) Validate() error {
	return dara.Validate(s)
}
