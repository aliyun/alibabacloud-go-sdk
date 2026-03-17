// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUnbindFlowLogSagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeUnbindFlowLogSagsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeUnbindFlowLogSagsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeUnbindFlowLogSagsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeUnbindFlowLogSagsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeUnbindFlowLogSagsRequest
	GetResourceOwnerId() *int64
}

type DescribeUnbindFlowLogSagsRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the SAG instances are deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeUnbindFlowLogSagsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUnbindFlowLogSagsRequest) GoString() string {
	return s.String()
}

func (s *DescribeUnbindFlowLogSagsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeUnbindFlowLogSagsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeUnbindFlowLogSagsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeUnbindFlowLogSagsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeUnbindFlowLogSagsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeUnbindFlowLogSagsRequest) SetOwnerAccount(v string) *DescribeUnbindFlowLogSagsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeUnbindFlowLogSagsRequest) SetOwnerId(v int64) *DescribeUnbindFlowLogSagsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeUnbindFlowLogSagsRequest) SetRegionId(v string) *DescribeUnbindFlowLogSagsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUnbindFlowLogSagsRequest) SetResourceOwnerAccount(v string) *DescribeUnbindFlowLogSagsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeUnbindFlowLogSagsRequest) SetResourceOwnerId(v int64) *DescribeUnbindFlowLogSagsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeUnbindFlowLogSagsRequest) Validate() error {
	return dara.Validate(s)
}
