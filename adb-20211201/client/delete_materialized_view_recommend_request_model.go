// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMaterializedViewRecommendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteMaterializedViewRecommendRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DeleteMaterializedViewRecommendRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteMaterializedViewRecommendRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteMaterializedViewRecommendRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteMaterializedViewRecommendRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteMaterializedViewRecommendRequest
	GetResourceOwnerId() *int64
	SetTaskName(v string) *DeleteMaterializedViewRecommendRequest
	GetTaskName() *string
}

type DeleteMaterializedViewRecommendRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-uf66*****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The name of the recommendation task.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_task_1
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s DeleteMaterializedViewRecommendRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMaterializedViewRecommendRequest) GoString() string {
	return s.String()
}

func (s *DeleteMaterializedViewRecommendRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteMaterializedViewRecommendRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteMaterializedViewRecommendRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteMaterializedViewRecommendRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteMaterializedViewRecommendRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteMaterializedViewRecommendRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteMaterializedViewRecommendRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *DeleteMaterializedViewRecommendRequest) SetDBClusterId(v string) *DeleteMaterializedViewRecommendRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteMaterializedViewRecommendRequest) SetOwnerAccount(v string) *DeleteMaterializedViewRecommendRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteMaterializedViewRecommendRequest) SetOwnerId(v int64) *DeleteMaterializedViewRecommendRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteMaterializedViewRecommendRequest) SetRegionId(v string) *DeleteMaterializedViewRecommendRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteMaterializedViewRecommendRequest) SetResourceOwnerAccount(v string) *DeleteMaterializedViewRecommendRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteMaterializedViewRecommendRequest) SetResourceOwnerId(v int64) *DeleteMaterializedViewRecommendRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteMaterializedViewRecommendRequest) SetTaskName(v string) *DeleteMaterializedViewRecommendRequest {
	s.TaskName = &v
	return s
}

func (s *DeleteMaterializedViewRecommendRequest) Validate() error {
	return dara.Validate(s)
}
