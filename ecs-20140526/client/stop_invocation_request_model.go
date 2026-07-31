// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopInvocationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v []*string) *StopInvocationRequest
	GetInstanceId() []*string
	SetInvokeId(v string) *StopInvocationRequest
	GetInvokeId() *string
	SetOwnerAccount(v string) *StopInvocationRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *StopInvocationRequest
	GetOwnerId() *int64
	SetRegionId(v string) *StopInvocationRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *StopInvocationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *StopInvocationRequest
	GetResourceOwnerId() *int64
}

type StopInvocationRequest struct {
	// The IDs of instances on which you want to stop the command. You can specify up to 50 instance IDs. Valid values of N: 1 to 50.
	//
	// example:
	//
	// i-bp67acfmxazb4p****
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	// The command execution ID. You can call [DescribeInvocations](https://help.aliyun.com/document_detail/64840.html) to query all execution IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// t-7d2a745b412b4601b2d47f6a768d****
	InvokeId     *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s StopInvocationRequest) String() string {
	return dara.Prettify(s)
}

func (s StopInvocationRequest) GoString() string {
	return s.String()
}

func (s *StopInvocationRequest) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *StopInvocationRequest) GetInvokeId() *string {
	return s.InvokeId
}

func (s *StopInvocationRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *StopInvocationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StopInvocationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopInvocationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *StopInvocationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *StopInvocationRequest) SetInstanceId(v []*string) *StopInvocationRequest {
	s.InstanceId = v
	return s
}

func (s *StopInvocationRequest) SetInvokeId(v string) *StopInvocationRequest {
	s.InvokeId = &v
	return s
}

func (s *StopInvocationRequest) SetOwnerAccount(v string) *StopInvocationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *StopInvocationRequest) SetOwnerId(v int64) *StopInvocationRequest {
	s.OwnerId = &v
	return s
}

func (s *StopInvocationRequest) SetRegionId(v string) *StopInvocationRequest {
	s.RegionId = &v
	return s
}

func (s *StopInvocationRequest) SetResourceOwnerAccount(v string) *StopInvocationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StopInvocationRequest) SetResourceOwnerId(v int64) *StopInvocationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *StopInvocationRequest) Validate() error {
	return dara.Validate(s)
}
