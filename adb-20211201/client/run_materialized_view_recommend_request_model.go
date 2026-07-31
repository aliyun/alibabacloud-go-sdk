// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunMaterializedViewRecommendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *RunMaterializedViewRecommendRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *RunMaterializedViewRecommendRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RunMaterializedViewRecommendRequest
	GetOwnerId() *int64
	SetRegionId(v string) *RunMaterializedViewRecommendRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *RunMaterializedViewRecommendRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RunMaterializedViewRecommendRequest
	GetResourceOwnerId() *int64
	SetTaskName(v string) *RunMaterializedViewRecommendRequest
	GetTaskName() *string
}

type RunMaterializedViewRecommendRequest struct {
	// <props="china">The ID of an Enterprise Edition, Basic Edition, or Data Lakehouse Edition cluster.
	//
	// <props="intl">The ID of a Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1u8c0mgfg58****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The Region ID of the cluster.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The name of the recommendation task. To run tasks in a batch, separate their names with a comma.
	//
	// This parameter is required.
	//
	// example:
	//
	// task1
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s RunMaterializedViewRecommendRequest) String() string {
	return dara.Prettify(s)
}

func (s RunMaterializedViewRecommendRequest) GoString() string {
	return s.String()
}

func (s *RunMaterializedViewRecommendRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *RunMaterializedViewRecommendRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RunMaterializedViewRecommendRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RunMaterializedViewRecommendRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RunMaterializedViewRecommendRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RunMaterializedViewRecommendRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RunMaterializedViewRecommendRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *RunMaterializedViewRecommendRequest) SetDBClusterId(v string) *RunMaterializedViewRecommendRequest {
	s.DBClusterId = &v
	return s
}

func (s *RunMaterializedViewRecommendRequest) SetOwnerAccount(v string) *RunMaterializedViewRecommendRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RunMaterializedViewRecommendRequest) SetOwnerId(v int64) *RunMaterializedViewRecommendRequest {
	s.OwnerId = &v
	return s
}

func (s *RunMaterializedViewRecommendRequest) SetRegionId(v string) *RunMaterializedViewRecommendRequest {
	s.RegionId = &v
	return s
}

func (s *RunMaterializedViewRecommendRequest) SetResourceOwnerAccount(v string) *RunMaterializedViewRecommendRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RunMaterializedViewRecommendRequest) SetResourceOwnerId(v int64) *RunMaterializedViewRecommendRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RunMaterializedViewRecommendRequest) SetTaskName(v string) *RunMaterializedViewRecommendRequest {
	s.TaskName = &v
	return s
}

func (s *RunMaterializedViewRecommendRequest) Validate() error {
	return dara.Validate(s)
}
