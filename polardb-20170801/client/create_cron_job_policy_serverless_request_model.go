// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCronJobPolicyServerlessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowShutDown(v string) *CreateCronJobPolicyServerlessRequest
	GetAllowShutDown() *string
	SetCronExpression(v string) *CreateCronJobPolicyServerlessRequest
	GetCronExpression() *string
	SetDBClusterId(v string) *CreateCronJobPolicyServerlessRequest
	GetDBClusterId() *string
	SetEndTime(v string) *CreateCronJobPolicyServerlessRequest
	GetEndTime() *string
	SetOwnerAccount(v string) *CreateCronJobPolicyServerlessRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateCronJobPolicyServerlessRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateCronJobPolicyServerlessRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateCronJobPolicyServerlessRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateCronJobPolicyServerlessRequest
	GetResourceOwnerId() *int64
	SetScaleApRoNumMax(v string) *CreateCronJobPolicyServerlessRequest
	GetScaleApRoNumMax() *string
	SetScaleApRoNumMin(v string) *CreateCronJobPolicyServerlessRequest
	GetScaleApRoNumMin() *string
	SetScaleMax(v string) *CreateCronJobPolicyServerlessRequest
	GetScaleMax() *string
	SetScaleMin(v string) *CreateCronJobPolicyServerlessRequest
	GetScaleMin() *string
	SetScaleRoNumMax(v string) *CreateCronJobPolicyServerlessRequest
	GetScaleRoNumMax() *string
	SetScaleRoNumMin(v string) *CreateCronJobPolicyServerlessRequest
	GetScaleRoNumMin() *string
	SetSecondsUntilAutoPause(v string) *CreateCronJobPolicyServerlessRequest
	GetSecondsUntilAutoPause() *string
	SetServerlessRuleCpuEnlargeThreshold(v string) *CreateCronJobPolicyServerlessRequest
	GetServerlessRuleCpuEnlargeThreshold() *string
	SetServerlessRuleCpuShrinkThreshold(v string) *CreateCronJobPolicyServerlessRequest
	GetServerlessRuleCpuShrinkThreshold() *string
	SetServerlessRuleMode(v string) *CreateCronJobPolicyServerlessRequest
	GetServerlessRuleMode() *string
	SetStartTime(v string) *CreateCronJobPolicyServerlessRequest
	GetStartTime() *string
}

type CreateCronJobPolicyServerlessRequest struct {
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
	// pc-xxxxxxxxxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-09-23T01:01:00Z
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
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
	// 3
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
	// 60
	ServerlessRuleCpuEnlargeThreshold *string `json:"ServerlessRuleCpuEnlargeThreshold,omitempty" xml:"ServerlessRuleCpuEnlargeThreshold,omitempty"`
	// example:
	//
	// 30
	ServerlessRuleCpuShrinkThreshold *string `json:"ServerlessRuleCpuShrinkThreshold,omitempty" xml:"ServerlessRuleCpuShrinkThreshold,omitempty"`
	// example:
	//
	// normal
	ServerlessRuleMode *string `json:"ServerlessRuleMode,omitempty" xml:"ServerlessRuleMode,omitempty"`
	// example:
	//
	// 2020-11-14T00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CreateCronJobPolicyServerlessRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCronJobPolicyServerlessRequest) GoString() string {
	return s.String()
}

func (s *CreateCronJobPolicyServerlessRequest) GetAllowShutDown() *string {
	return s.AllowShutDown
}

func (s *CreateCronJobPolicyServerlessRequest) GetCronExpression() *string {
	return s.CronExpression
}

func (s *CreateCronJobPolicyServerlessRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateCronJobPolicyServerlessRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateCronJobPolicyServerlessRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateCronJobPolicyServerlessRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateCronJobPolicyServerlessRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateCronJobPolicyServerlessRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateCronJobPolicyServerlessRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateCronJobPolicyServerlessRequest) GetScaleApRoNumMax() *string {
	return s.ScaleApRoNumMax
}

func (s *CreateCronJobPolicyServerlessRequest) GetScaleApRoNumMin() *string {
	return s.ScaleApRoNumMin
}

func (s *CreateCronJobPolicyServerlessRequest) GetScaleMax() *string {
	return s.ScaleMax
}

func (s *CreateCronJobPolicyServerlessRequest) GetScaleMin() *string {
	return s.ScaleMin
}

func (s *CreateCronJobPolicyServerlessRequest) GetScaleRoNumMax() *string {
	return s.ScaleRoNumMax
}

func (s *CreateCronJobPolicyServerlessRequest) GetScaleRoNumMin() *string {
	return s.ScaleRoNumMin
}

func (s *CreateCronJobPolicyServerlessRequest) GetSecondsUntilAutoPause() *string {
	return s.SecondsUntilAutoPause
}

func (s *CreateCronJobPolicyServerlessRequest) GetServerlessRuleCpuEnlargeThreshold() *string {
	return s.ServerlessRuleCpuEnlargeThreshold
}

func (s *CreateCronJobPolicyServerlessRequest) GetServerlessRuleCpuShrinkThreshold() *string {
	return s.ServerlessRuleCpuShrinkThreshold
}

func (s *CreateCronJobPolicyServerlessRequest) GetServerlessRuleMode() *string {
	return s.ServerlessRuleMode
}

func (s *CreateCronJobPolicyServerlessRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateCronJobPolicyServerlessRequest) SetAllowShutDown(v string) *CreateCronJobPolicyServerlessRequest {
	s.AllowShutDown = &v
	return s
}

func (s *CreateCronJobPolicyServerlessRequest) SetCronExpression(v string) *CreateCronJobPolicyServerlessRequest {
	s.CronExpression = &v
	return s
}

func (s *CreateCronJobPolicyServerlessRequest) SetDBClusterId(v string) *CreateCronJobPolicyServerlessRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateCronJobPolicyServerlessRequest) SetEndTime(v string) *CreateCronJobPolicyServerlessRequest {
	s.EndTime = &v
	return s
}

func (s *CreateCronJobPolicyServerlessRequest) SetOwnerAccount(v string) *CreateCronJobPolicyServerlessRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateCronJobPolicyServerlessRequest) SetOwnerId(v int64) *CreateCronJobPolicyServerlessRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateCronJobPolicyServerlessRequest) SetRegionId(v string) *CreateCronJobPolicyServerlessRequest {
	s.RegionId = &v
	return s
}

func (s *CreateCronJobPolicyServerlessRequest) SetResourceOwnerAccount(v string) *CreateCronJobPolicyServerlessRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateCronJobPolicyServerlessRequest) SetResourceOwnerId(v int64) *CreateCronJobPolicyServerlessRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateCronJobPolicyServerlessRequest) SetScaleApRoNumMax(v string) *CreateCronJobPolicyServerlessRequest {
	s.ScaleApRoNumMax = &v
	return s
}

func (s *CreateCronJobPolicyServerlessRequest) SetScaleApRoNumMin(v string) *CreateCronJobPolicyServerlessRequest {
	s.ScaleApRoNumMin = &v
	return s
}

func (s *CreateCronJobPolicyServerlessRequest) SetScaleMax(v string) *CreateCronJobPolicyServerlessRequest {
	s.ScaleMax = &v
	return s
}

func (s *CreateCronJobPolicyServerlessRequest) SetScaleMin(v string) *CreateCronJobPolicyServerlessRequest {
	s.ScaleMin = &v
	return s
}

func (s *CreateCronJobPolicyServerlessRequest) SetScaleRoNumMax(v string) *CreateCronJobPolicyServerlessRequest {
	s.ScaleRoNumMax = &v
	return s
}

func (s *CreateCronJobPolicyServerlessRequest) SetScaleRoNumMin(v string) *CreateCronJobPolicyServerlessRequest {
	s.ScaleRoNumMin = &v
	return s
}

func (s *CreateCronJobPolicyServerlessRequest) SetSecondsUntilAutoPause(v string) *CreateCronJobPolicyServerlessRequest {
	s.SecondsUntilAutoPause = &v
	return s
}

func (s *CreateCronJobPolicyServerlessRequest) SetServerlessRuleCpuEnlargeThreshold(v string) *CreateCronJobPolicyServerlessRequest {
	s.ServerlessRuleCpuEnlargeThreshold = &v
	return s
}

func (s *CreateCronJobPolicyServerlessRequest) SetServerlessRuleCpuShrinkThreshold(v string) *CreateCronJobPolicyServerlessRequest {
	s.ServerlessRuleCpuShrinkThreshold = &v
	return s
}

func (s *CreateCronJobPolicyServerlessRequest) SetServerlessRuleMode(v string) *CreateCronJobPolicyServerlessRequest {
	s.ServerlessRuleMode = &v
	return s
}

func (s *CreateCronJobPolicyServerlessRequest) SetStartTime(v string) *CreateCronJobPolicyServerlessRequest {
	s.StartTime = &v
	return s
}

func (s *CreateCronJobPolicyServerlessRequest) Validate() error {
	return dara.Validate(s)
}
