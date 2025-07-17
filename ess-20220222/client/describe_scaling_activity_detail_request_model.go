// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScalingActivityDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeScalingActivityDetailRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeScalingActivityDetailRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeScalingActivityDetailRequest
	GetResourceOwnerId() *int64
	SetScalingActivityId(v string) *DescribeScalingActivityDetailRequest
	GetScalingActivityId() *string
}

type DescribeScalingActivityDetailRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the scaling activity that you want to query. You can call the DescribeScalingActivities operation to query the IDs of scaling activities.
	//
	// This parameter is required.
	//
	// example:
	//
	// asa-bp1c9djwrgxjyk31****
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
}

func (s DescribeScalingActivityDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingActivityDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeScalingActivityDetailRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeScalingActivityDetailRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeScalingActivityDetailRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeScalingActivityDetailRequest) GetScalingActivityId() *string {
	return s.ScalingActivityId
}

func (s *DescribeScalingActivityDetailRequest) SetOwnerId(v int64) *DescribeScalingActivityDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeScalingActivityDetailRequest) SetResourceOwnerAccount(v string) *DescribeScalingActivityDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeScalingActivityDetailRequest) SetResourceOwnerId(v int64) *DescribeScalingActivityDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeScalingActivityDetailRequest) SetScalingActivityId(v string) *DescribeScalingActivityDetailRequest {
	s.ScalingActivityId = &v
	return s
}

func (s *DescribeScalingActivityDetailRequest) Validate() error {
	return dara.Validate(s)
}
