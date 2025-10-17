// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMaterializedViewRecommendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CreateMaterializedViewRecommendRequest
	GetDBClusterId() *string
	SetDescription(v string) *CreateMaterializedViewRecommendRequest
	GetDescription() *string
	SetMinRewriteQueryCount(v int32) *CreateMaterializedViewRecommendRequest
	GetMinRewriteQueryCount() *int32
	SetMinRewriteQueryPattern(v int32) *CreateMaterializedViewRecommendRequest
	GetMinRewriteQueryPattern() *int32
	SetOwnerAccount(v string) *CreateMaterializedViewRecommendRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateMaterializedViewRecommendRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateMaterializedViewRecommendRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateMaterializedViewRecommendRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateMaterializedViewRecommendRequest
	GetResourceOwnerId() *int64
	SetScanQueriesRange(v int32) *CreateMaterializedViewRecommendRequest
	GetScanQueriesRange() *int32
	SetSchedulingDay(v string) *CreateMaterializedViewRecommendRequest
	GetSchedulingDay() *string
	SetSchedulingPolicy(v string) *CreateMaterializedViewRecommendRequest
	GetSchedulingPolicy() *string
	SetSlowQueryThreshold(v int32) *CreateMaterializedViewRecommendRequest
	GetSlowQueryThreshold() *int32
	SetSpecifiedTime(v string) *CreateMaterializedViewRecommendRequest
	GetSpecifiedTime() *string
	SetTaskName(v string) *CreateMaterializedViewRecommendRequest
	GetTaskName() *string
}

type CreateMaterializedViewRecommendRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// amv-8vbwm***
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
	// This parameter is required.
	//
	// example:
	//
	// weekly
	SchedulingPolicy   *string `json:"SchedulingPolicy,omitempty" xml:"SchedulingPolicy,omitempty"`
	SlowQueryThreshold *int32  `json:"SlowQueryThreshold,omitempty" xml:"SlowQueryThreshold,omitempty"`
	// This parameter is required.
	//
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

func (s CreateMaterializedViewRecommendRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMaterializedViewRecommendRequest) GoString() string {
	return s.String()
}

func (s *CreateMaterializedViewRecommendRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateMaterializedViewRecommendRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateMaterializedViewRecommendRequest) GetMinRewriteQueryCount() *int32 {
	return s.MinRewriteQueryCount
}

func (s *CreateMaterializedViewRecommendRequest) GetMinRewriteQueryPattern() *int32 {
	return s.MinRewriteQueryPattern
}

func (s *CreateMaterializedViewRecommendRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateMaterializedViewRecommendRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateMaterializedViewRecommendRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateMaterializedViewRecommendRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateMaterializedViewRecommendRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateMaterializedViewRecommendRequest) GetScanQueriesRange() *int32 {
	return s.ScanQueriesRange
}

func (s *CreateMaterializedViewRecommendRequest) GetSchedulingDay() *string {
	return s.SchedulingDay
}

func (s *CreateMaterializedViewRecommendRequest) GetSchedulingPolicy() *string {
	return s.SchedulingPolicy
}

func (s *CreateMaterializedViewRecommendRequest) GetSlowQueryThreshold() *int32 {
	return s.SlowQueryThreshold
}

func (s *CreateMaterializedViewRecommendRequest) GetSpecifiedTime() *string {
	return s.SpecifiedTime
}

func (s *CreateMaterializedViewRecommendRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateMaterializedViewRecommendRequest) SetDBClusterId(v string) *CreateMaterializedViewRecommendRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateMaterializedViewRecommendRequest) SetDescription(v string) *CreateMaterializedViewRecommendRequest {
	s.Description = &v
	return s
}

func (s *CreateMaterializedViewRecommendRequest) SetMinRewriteQueryCount(v int32) *CreateMaterializedViewRecommendRequest {
	s.MinRewriteQueryCount = &v
	return s
}

func (s *CreateMaterializedViewRecommendRequest) SetMinRewriteQueryPattern(v int32) *CreateMaterializedViewRecommendRequest {
	s.MinRewriteQueryPattern = &v
	return s
}

func (s *CreateMaterializedViewRecommendRequest) SetOwnerAccount(v string) *CreateMaterializedViewRecommendRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateMaterializedViewRecommendRequest) SetOwnerId(v int64) *CreateMaterializedViewRecommendRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateMaterializedViewRecommendRequest) SetRegionId(v string) *CreateMaterializedViewRecommendRequest {
	s.RegionId = &v
	return s
}

func (s *CreateMaterializedViewRecommendRequest) SetResourceOwnerAccount(v string) *CreateMaterializedViewRecommendRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateMaterializedViewRecommendRequest) SetResourceOwnerId(v int64) *CreateMaterializedViewRecommendRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateMaterializedViewRecommendRequest) SetScanQueriesRange(v int32) *CreateMaterializedViewRecommendRequest {
	s.ScanQueriesRange = &v
	return s
}

func (s *CreateMaterializedViewRecommendRequest) SetSchedulingDay(v string) *CreateMaterializedViewRecommendRequest {
	s.SchedulingDay = &v
	return s
}

func (s *CreateMaterializedViewRecommendRequest) SetSchedulingPolicy(v string) *CreateMaterializedViewRecommendRequest {
	s.SchedulingPolicy = &v
	return s
}

func (s *CreateMaterializedViewRecommendRequest) SetSlowQueryThreshold(v int32) *CreateMaterializedViewRecommendRequest {
	s.SlowQueryThreshold = &v
	return s
}

func (s *CreateMaterializedViewRecommendRequest) SetSpecifiedTime(v string) *CreateMaterializedViewRecommendRequest {
	s.SpecifiedTime = &v
	return s
}

func (s *CreateMaterializedViewRecommendRequest) SetTaskName(v string) *CreateMaterializedViewRecommendRequest {
	s.TaskName = &v
	return s
}

func (s *CreateMaterializedViewRecommendRequest) Validate() error {
	return dara.Validate(s)
}
