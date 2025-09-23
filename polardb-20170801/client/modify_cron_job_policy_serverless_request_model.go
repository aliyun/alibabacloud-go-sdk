// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCronJobPolicyServerlessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowShutDown(v string) *ModifyCronJobPolicyServerlessRequest
	GetAllowShutDown() *string
	SetCronExpression(v string) *ModifyCronJobPolicyServerlessRequest
	GetCronExpression() *string
	SetDBClusterId(v string) *ModifyCronJobPolicyServerlessRequest
	GetDBClusterId() *string
	SetEndTime(v string) *ModifyCronJobPolicyServerlessRequest
	GetEndTime() *string
	SetJobId(v string) *ModifyCronJobPolicyServerlessRequest
	GetJobId() *string
	SetOwnerAccount(v string) *ModifyCronJobPolicyServerlessRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyCronJobPolicyServerlessRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyCronJobPolicyServerlessRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyCronJobPolicyServerlessRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyCronJobPolicyServerlessRequest
	GetResourceOwnerId() *int64
	SetScaleApRoNumMax(v string) *ModifyCronJobPolicyServerlessRequest
	GetScaleApRoNumMax() *string
	SetScaleApRoNumMin(v string) *ModifyCronJobPolicyServerlessRequest
	GetScaleApRoNumMin() *string
	SetScaleMax(v string) *ModifyCronJobPolicyServerlessRequest
	GetScaleMax() *string
	SetScaleMin(v string) *ModifyCronJobPolicyServerlessRequest
	GetScaleMin() *string
	SetScaleRoNumMax(v string) *ModifyCronJobPolicyServerlessRequest
	GetScaleRoNumMax() *string
	SetScaleRoNumMin(v string) *ModifyCronJobPolicyServerlessRequest
	GetScaleRoNumMin() *string
	SetSecondsUntilAutoPause(v string) *ModifyCronJobPolicyServerlessRequest
	GetSecondsUntilAutoPause() *string
	SetServerlessRuleCpuEnlargeThreshold(v string) *ModifyCronJobPolicyServerlessRequest
	GetServerlessRuleCpuEnlargeThreshold() *string
	SetServerlessRuleCpuShrinkThreshold(v string) *ModifyCronJobPolicyServerlessRequest
	GetServerlessRuleCpuShrinkThreshold() *string
	SetServerlessRuleMode(v string) *ModifyCronJobPolicyServerlessRequest
	GetServerlessRuleMode() *string
	SetStartTime(v string) *ModifyCronJobPolicyServerlessRequest
	GetStartTime() *string
}

type ModifyCronJobPolicyServerlessRequest struct {
	// example:
	//
	// true
	AllowShutDown *string `json:"AllowShutDown,omitempty" xml:"AllowShutDown,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0 0 13 	- 	- ?
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pc-************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2020-02-12T15:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 8006e51c-dab3-4602-bc69-4f728002c6ce
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 1
	ScaleApRoNumMax *string `json:"ScaleApRoNumMax,omitempty" xml:"ScaleApRoNumMax,omitempty"`
	// example:
	//
	// 1
	ScaleApRoNumMin *string `json:"ScaleApRoNumMin,omitempty" xml:"ScaleApRoNumMin,omitempty"`
	// example:
	//
	// 12
	ScaleMax *string `json:"ScaleMax,omitempty" xml:"ScaleMax,omitempty"`
	// example:
	//
	// 1
	ScaleMin *string `json:"ScaleMin,omitempty" xml:"ScaleMin,omitempty"`
	// example:
	//
	// 4
	ScaleRoNumMax *string `json:"ScaleRoNumMax,omitempty" xml:"ScaleRoNumMax,omitempty"`
	// example:
	//
	// 2
	ScaleRoNumMin *string `json:"ScaleRoNumMin,omitempty" xml:"ScaleRoNumMin,omitempty"`
	// example:
	//
	// 10
	SecondsUntilAutoPause *string `json:"SecondsUntilAutoPause,omitempty" xml:"SecondsUntilAutoPause,omitempty"`
	// example:
	//
	// 80
	ServerlessRuleCpuEnlargeThreshold *string `json:"ServerlessRuleCpuEnlargeThreshold,omitempty" xml:"ServerlessRuleCpuEnlargeThreshold,omitempty"`
	// example:
	//
	// 25
	ServerlessRuleCpuShrinkThreshold *string `json:"ServerlessRuleCpuShrinkThreshold,omitempty" xml:"ServerlessRuleCpuShrinkThreshold,omitempty"`
	// example:
	//
	// normal
	ServerlessRuleMode *string `json:"ServerlessRuleMode,omitempty" xml:"ServerlessRuleMode,omitempty"`
	// example:
	//
	// 2020-09-23T01:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ModifyCronJobPolicyServerlessRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCronJobPolicyServerlessRequest) GoString() string {
	return s.String()
}

func (s *ModifyCronJobPolicyServerlessRequest) GetAllowShutDown() *string {
	return s.AllowShutDown
}

func (s *ModifyCronJobPolicyServerlessRequest) GetCronExpression() *string {
	return s.CronExpression
}

func (s *ModifyCronJobPolicyServerlessRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyCronJobPolicyServerlessRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ModifyCronJobPolicyServerlessRequest) GetJobId() *string {
	return s.JobId
}

func (s *ModifyCronJobPolicyServerlessRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyCronJobPolicyServerlessRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyCronJobPolicyServerlessRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyCronJobPolicyServerlessRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyCronJobPolicyServerlessRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyCronJobPolicyServerlessRequest) GetScaleApRoNumMax() *string {
	return s.ScaleApRoNumMax
}

func (s *ModifyCronJobPolicyServerlessRequest) GetScaleApRoNumMin() *string {
	return s.ScaleApRoNumMin
}

func (s *ModifyCronJobPolicyServerlessRequest) GetScaleMax() *string {
	return s.ScaleMax
}

func (s *ModifyCronJobPolicyServerlessRequest) GetScaleMin() *string {
	return s.ScaleMin
}

func (s *ModifyCronJobPolicyServerlessRequest) GetScaleRoNumMax() *string {
	return s.ScaleRoNumMax
}

func (s *ModifyCronJobPolicyServerlessRequest) GetScaleRoNumMin() *string {
	return s.ScaleRoNumMin
}

func (s *ModifyCronJobPolicyServerlessRequest) GetSecondsUntilAutoPause() *string {
	return s.SecondsUntilAutoPause
}

func (s *ModifyCronJobPolicyServerlessRequest) GetServerlessRuleCpuEnlargeThreshold() *string {
	return s.ServerlessRuleCpuEnlargeThreshold
}

func (s *ModifyCronJobPolicyServerlessRequest) GetServerlessRuleCpuShrinkThreshold() *string {
	return s.ServerlessRuleCpuShrinkThreshold
}

func (s *ModifyCronJobPolicyServerlessRequest) GetServerlessRuleMode() *string {
	return s.ServerlessRuleMode
}

func (s *ModifyCronJobPolicyServerlessRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ModifyCronJobPolicyServerlessRequest) SetAllowShutDown(v string) *ModifyCronJobPolicyServerlessRequest {
	s.AllowShutDown = &v
	return s
}

func (s *ModifyCronJobPolicyServerlessRequest) SetCronExpression(v string) *ModifyCronJobPolicyServerlessRequest {
	s.CronExpression = &v
	return s
}

func (s *ModifyCronJobPolicyServerlessRequest) SetDBClusterId(v string) *ModifyCronJobPolicyServerlessRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyCronJobPolicyServerlessRequest) SetEndTime(v string) *ModifyCronJobPolicyServerlessRequest {
	s.EndTime = &v
	return s
}

func (s *ModifyCronJobPolicyServerlessRequest) SetJobId(v string) *ModifyCronJobPolicyServerlessRequest {
	s.JobId = &v
	return s
}

func (s *ModifyCronJobPolicyServerlessRequest) SetOwnerAccount(v string) *ModifyCronJobPolicyServerlessRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyCronJobPolicyServerlessRequest) SetOwnerId(v int64) *ModifyCronJobPolicyServerlessRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyCronJobPolicyServerlessRequest) SetRegionId(v string) *ModifyCronJobPolicyServerlessRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyCronJobPolicyServerlessRequest) SetResourceOwnerAccount(v string) *ModifyCronJobPolicyServerlessRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyCronJobPolicyServerlessRequest) SetResourceOwnerId(v int64) *ModifyCronJobPolicyServerlessRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyCronJobPolicyServerlessRequest) SetScaleApRoNumMax(v string) *ModifyCronJobPolicyServerlessRequest {
	s.ScaleApRoNumMax = &v
	return s
}

func (s *ModifyCronJobPolicyServerlessRequest) SetScaleApRoNumMin(v string) *ModifyCronJobPolicyServerlessRequest {
	s.ScaleApRoNumMin = &v
	return s
}

func (s *ModifyCronJobPolicyServerlessRequest) SetScaleMax(v string) *ModifyCronJobPolicyServerlessRequest {
	s.ScaleMax = &v
	return s
}

func (s *ModifyCronJobPolicyServerlessRequest) SetScaleMin(v string) *ModifyCronJobPolicyServerlessRequest {
	s.ScaleMin = &v
	return s
}

func (s *ModifyCronJobPolicyServerlessRequest) SetScaleRoNumMax(v string) *ModifyCronJobPolicyServerlessRequest {
	s.ScaleRoNumMax = &v
	return s
}

func (s *ModifyCronJobPolicyServerlessRequest) SetScaleRoNumMin(v string) *ModifyCronJobPolicyServerlessRequest {
	s.ScaleRoNumMin = &v
	return s
}

func (s *ModifyCronJobPolicyServerlessRequest) SetSecondsUntilAutoPause(v string) *ModifyCronJobPolicyServerlessRequest {
	s.SecondsUntilAutoPause = &v
	return s
}

func (s *ModifyCronJobPolicyServerlessRequest) SetServerlessRuleCpuEnlargeThreshold(v string) *ModifyCronJobPolicyServerlessRequest {
	s.ServerlessRuleCpuEnlargeThreshold = &v
	return s
}

func (s *ModifyCronJobPolicyServerlessRequest) SetServerlessRuleCpuShrinkThreshold(v string) *ModifyCronJobPolicyServerlessRequest {
	s.ServerlessRuleCpuShrinkThreshold = &v
	return s
}

func (s *ModifyCronJobPolicyServerlessRequest) SetServerlessRuleMode(v string) *ModifyCronJobPolicyServerlessRequest {
	s.ServerlessRuleMode = &v
	return s
}

func (s *ModifyCronJobPolicyServerlessRequest) SetStartTime(v string) *ModifyCronJobPolicyServerlessRequest {
	s.StartTime = &v
	return s
}

func (s *ModifyCronJobPolicyServerlessRequest) Validate() error {
	return dara.Validate(s)
}
