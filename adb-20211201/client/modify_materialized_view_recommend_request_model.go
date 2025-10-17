// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMaterializedViewRecommendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyMaterializedViewRecommendRequest
	GetDBClusterId() *string
	SetDescription(v string) *ModifyMaterializedViewRecommendRequest
	GetDescription() *string
	SetMinRewriteQueryCount(v int32) *ModifyMaterializedViewRecommendRequest
	GetMinRewriteQueryCount() *int32
	SetMinRewriteQueryPattern(v int32) *ModifyMaterializedViewRecommendRequest
	GetMinRewriteQueryPattern() *int32
	SetOwnerAccount(v string) *ModifyMaterializedViewRecommendRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyMaterializedViewRecommendRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyMaterializedViewRecommendRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyMaterializedViewRecommendRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyMaterializedViewRecommendRequest
	GetResourceOwnerId() *int64
	SetScanQueriesRange(v int32) *ModifyMaterializedViewRecommendRequest
	GetScanQueriesRange() *int32
	SetSchedulingDay(v string) *ModifyMaterializedViewRecommendRequest
	GetSchedulingDay() *string
	SetSchedulingPolicy(v string) *ModifyMaterializedViewRecommendRequest
	GetSchedulingPolicy() *string
	SetSlowQueryThreshold(v int32) *ModifyMaterializedViewRecommendRequest
	GetSlowQueryThreshold() *int32
	SetSpecifiedTime(v string) *ModifyMaterializedViewRecommendRequest
	GetSpecifiedTime() *string
	SetTaskName(v string) *ModifyMaterializedViewRecommendRequest
	GetTaskName() *string
}

type ModifyMaterializedViewRecommendRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// amv-uf6o*****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// task desc
	Description            *string `json:"Description,omitempty" xml:"Description,omitempty"`
	MinRewriteQueryCount   *int32  `json:"MinRewriteQueryCount,omitempty" xml:"MinRewriteQueryCount,omitempty"`
	MinRewriteQueryPattern *int32  `json:"MinRewriteQueryPattern,omitempty" xml:"MinRewriteQueryPattern,omitempty"`
	OwnerAccount           *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 3
	ScanQueriesRange *int32 `json:"ScanQueriesRange,omitempty" xml:"ScanQueriesRange,omitempty"`
	// example:
	//
	// Monday;Wednesday
	SchedulingDay *string `json:"SchedulingDay,omitempty" xml:"SchedulingDay,omitempty"`
	// example:
	//
	// weekly
	SchedulingPolicy   *string `json:"SchedulingPolicy,omitempty" xml:"SchedulingPolicy,omitempty"`
	SlowQueryThreshold *int32  `json:"SlowQueryThreshold,omitempty" xml:"SlowQueryThreshold,omitempty"`
	// example:
	//
	// 10:00:00
	SpecifiedTime *string `json:"SpecifiedTime,omitempty" xml:"SpecifiedTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// task_n1
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s ModifyMaterializedViewRecommendRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyMaterializedViewRecommendRequest) GoString() string {
	return s.String()
}

func (s *ModifyMaterializedViewRecommendRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyMaterializedViewRecommendRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyMaterializedViewRecommendRequest) GetMinRewriteQueryCount() *int32 {
	return s.MinRewriteQueryCount
}

func (s *ModifyMaterializedViewRecommendRequest) GetMinRewriteQueryPattern() *int32 {
	return s.MinRewriteQueryPattern
}

func (s *ModifyMaterializedViewRecommendRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyMaterializedViewRecommendRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyMaterializedViewRecommendRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyMaterializedViewRecommendRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyMaterializedViewRecommendRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyMaterializedViewRecommendRequest) GetScanQueriesRange() *int32 {
	return s.ScanQueriesRange
}

func (s *ModifyMaterializedViewRecommendRequest) GetSchedulingDay() *string {
	return s.SchedulingDay
}

func (s *ModifyMaterializedViewRecommendRequest) GetSchedulingPolicy() *string {
	return s.SchedulingPolicy
}

func (s *ModifyMaterializedViewRecommendRequest) GetSlowQueryThreshold() *int32 {
	return s.SlowQueryThreshold
}

func (s *ModifyMaterializedViewRecommendRequest) GetSpecifiedTime() *string {
	return s.SpecifiedTime
}

func (s *ModifyMaterializedViewRecommendRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *ModifyMaterializedViewRecommendRequest) SetDBClusterId(v string) *ModifyMaterializedViewRecommendRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyMaterializedViewRecommendRequest) SetDescription(v string) *ModifyMaterializedViewRecommendRequest {
	s.Description = &v
	return s
}

func (s *ModifyMaterializedViewRecommendRequest) SetMinRewriteQueryCount(v int32) *ModifyMaterializedViewRecommendRequest {
	s.MinRewriteQueryCount = &v
	return s
}

func (s *ModifyMaterializedViewRecommendRequest) SetMinRewriteQueryPattern(v int32) *ModifyMaterializedViewRecommendRequest {
	s.MinRewriteQueryPattern = &v
	return s
}

func (s *ModifyMaterializedViewRecommendRequest) SetOwnerAccount(v string) *ModifyMaterializedViewRecommendRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyMaterializedViewRecommendRequest) SetOwnerId(v int64) *ModifyMaterializedViewRecommendRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyMaterializedViewRecommendRequest) SetRegionId(v string) *ModifyMaterializedViewRecommendRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyMaterializedViewRecommendRequest) SetResourceOwnerAccount(v string) *ModifyMaterializedViewRecommendRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyMaterializedViewRecommendRequest) SetResourceOwnerId(v int64) *ModifyMaterializedViewRecommendRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyMaterializedViewRecommendRequest) SetScanQueriesRange(v int32) *ModifyMaterializedViewRecommendRequest {
	s.ScanQueriesRange = &v
	return s
}

func (s *ModifyMaterializedViewRecommendRequest) SetSchedulingDay(v string) *ModifyMaterializedViewRecommendRequest {
	s.SchedulingDay = &v
	return s
}

func (s *ModifyMaterializedViewRecommendRequest) SetSchedulingPolicy(v string) *ModifyMaterializedViewRecommendRequest {
	s.SchedulingPolicy = &v
	return s
}

func (s *ModifyMaterializedViewRecommendRequest) SetSlowQueryThreshold(v int32) *ModifyMaterializedViewRecommendRequest {
	s.SlowQueryThreshold = &v
	return s
}

func (s *ModifyMaterializedViewRecommendRequest) SetSpecifiedTime(v string) *ModifyMaterializedViewRecommendRequest {
	s.SpecifiedTime = &v
	return s
}

func (s *ModifyMaterializedViewRecommendRequest) SetTaskName(v string) *ModifyMaterializedViewRecommendRequest {
	s.TaskName = &v
	return s
}

func (s *ModifyMaterializedViewRecommendRequest) Validate() error {
	return dara.Validate(s)
}
