// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEndpointSwitchStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *DescribeEndpointSwitchStatusRequest
	GetAccountId() *string
	SetClientToken(v string) *DescribeEndpointSwitchStatusRequest
	GetClientToken() *string
	SetOwnerId(v string) *DescribeEndpointSwitchStatusRequest
	GetOwnerId() *string
	SetRegionId(v string) *DescribeEndpointSwitchStatusRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeEndpointSwitchStatusRequest
	GetResourceGroupId() *string
	SetTaskId(v string) *DescribeEndpointSwitchStatusRequest
	GetTaskId() *string
}

type DescribeEndpointSwitchStatusRequest struct {
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter because this parameter will be removed in the future.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The **ClientToken*	- parameter can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe63****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The task ID, which is returned after you call the [SwitchSynchronizationEndpoint](https://help.aliyun.com/document_detail/201858.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeEndpointSwitchStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEndpointSwitchStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeEndpointSwitchStatusRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *DescribeEndpointSwitchStatusRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeEndpointSwitchStatusRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeEndpointSwitchStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEndpointSwitchStatusRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeEndpointSwitchStatusRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeEndpointSwitchStatusRequest) SetAccountId(v string) *DescribeEndpointSwitchStatusRequest {
	s.AccountId = &v
	return s
}

func (s *DescribeEndpointSwitchStatusRequest) SetClientToken(v string) *DescribeEndpointSwitchStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeEndpointSwitchStatusRequest) SetOwnerId(v string) *DescribeEndpointSwitchStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeEndpointSwitchStatusRequest) SetRegionId(v string) *DescribeEndpointSwitchStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEndpointSwitchStatusRequest) SetResourceGroupId(v string) *DescribeEndpointSwitchStatusRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeEndpointSwitchStatusRequest) SetTaskId(v string) *DescribeEndpointSwitchStatusRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeEndpointSwitchStatusRequest) Validate() error {
	return dara.Validate(s)
}
