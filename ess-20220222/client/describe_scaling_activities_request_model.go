// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScalingActivitiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceRefreshTaskId(v string) *DescribeScalingActivitiesRequest
	GetInstanceRefreshTaskId() *string
	SetOwnerAccount(v string) *DescribeScalingActivitiesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeScalingActivitiesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeScalingActivitiesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeScalingActivitiesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeScalingActivitiesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeScalingActivitiesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeScalingActivitiesRequest
	GetResourceOwnerId() *int64
	SetScalingActivityIds(v []*string) *DescribeScalingActivitiesRequest
	GetScalingActivityIds() []*string
	SetScalingGroupId(v string) *DescribeScalingActivitiesRequest
	GetScalingGroupId() *string
	SetStatusCode(v string) *DescribeScalingActivitiesRequest
	GetStatusCode() *string
}

type DescribeScalingActivitiesRequest struct {
	// The ID of the instance refresh task. If you specify this parameter, this operation returns the list of scaling activities associated with the instance refresh task.
	//
	// example:
	//
	// ir-a12ds234fasd*****
	InstanceRefreshTaskId *string `json:"InstanceRefreshTaskId,omitempty" xml:"InstanceRefreshTaskId,omitempty"`
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: 50.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the scaling group to which the scaling activity that you want to query belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The IDs of the scaling activities that you want to query.
	//
	// >  When you call this operation, you must specify one of the following parameters: `ScalingGroupId` and `ScalingActivityIds`. You cannot specify both of them at the same time. If you specify neither of them, an error is reported.
	ScalingActivityIds []*string `json:"ScalingActivityIds,omitempty" xml:"ScalingActivityIds,omitempty" type:"Repeated"`
	// The ID of the scaling group.
	//
	// >  When you call this operation, you must specify one of the following parameters: `ScalingGroupId` and `ScalingActivityIds`. You cannot specify both of them at the same time. If you specify neither of them, an error is reported.
	//
	// example:
	//
	// asg-bp18p2yfxow2dloq****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// The status of the scaling activity. Valid values:
	//
	// 	- Successful: The scaling activity is successful.
	//
	// 	- Warning: The scaling activity is partially successful.
	//
	// 	- Failed: The scaling activity failed.
	//
	// 	- InProgress: The scaling activity is in progress.
	//
	// 	- Rejected: The request to trigger the scaling activity is rejected.
	//
	// example:
	//
	// Successful
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s DescribeScalingActivitiesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingActivitiesRequest) GoString() string {
	return s.String()
}

func (s *DescribeScalingActivitiesRequest) GetInstanceRefreshTaskId() *string {
	return s.InstanceRefreshTaskId
}

func (s *DescribeScalingActivitiesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeScalingActivitiesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeScalingActivitiesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeScalingActivitiesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeScalingActivitiesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeScalingActivitiesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeScalingActivitiesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeScalingActivitiesRequest) GetScalingActivityIds() []*string {
	return s.ScalingActivityIds
}

func (s *DescribeScalingActivitiesRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DescribeScalingActivitiesRequest) GetStatusCode() *string {
	return s.StatusCode
}

func (s *DescribeScalingActivitiesRequest) SetInstanceRefreshTaskId(v string) *DescribeScalingActivitiesRequest {
	s.InstanceRefreshTaskId = &v
	return s
}

func (s *DescribeScalingActivitiesRequest) SetOwnerAccount(v string) *DescribeScalingActivitiesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeScalingActivitiesRequest) SetOwnerId(v int64) *DescribeScalingActivitiesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeScalingActivitiesRequest) SetPageNumber(v int32) *DescribeScalingActivitiesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeScalingActivitiesRequest) SetPageSize(v int32) *DescribeScalingActivitiesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeScalingActivitiesRequest) SetRegionId(v string) *DescribeScalingActivitiesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeScalingActivitiesRequest) SetResourceOwnerAccount(v string) *DescribeScalingActivitiesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeScalingActivitiesRequest) SetResourceOwnerId(v int64) *DescribeScalingActivitiesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeScalingActivitiesRequest) SetScalingActivityIds(v []*string) *DescribeScalingActivitiesRequest {
	s.ScalingActivityIds = v
	return s
}

func (s *DescribeScalingActivitiesRequest) SetScalingGroupId(v string) *DescribeScalingActivitiesRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeScalingActivitiesRequest) SetStatusCode(v string) *DescribeScalingActivitiesRequest {
	s.StatusCode = &v
	return s
}

func (s *DescribeScalingActivitiesRequest) Validate() error {
	return dara.Validate(s)
}
