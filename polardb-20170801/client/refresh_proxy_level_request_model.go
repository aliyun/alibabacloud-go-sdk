// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshProxyLevelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *RefreshProxyLevelRequest
	GetDBClusterId() *string
	SetFromTimeService(v bool) *RefreshProxyLevelRequest
	GetFromTimeService() *bool
	SetOwnerAccount(v string) *RefreshProxyLevelRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RefreshProxyLevelRequest
	GetOwnerId() *int64
	SetPlannedEndTime(v string) *RefreshProxyLevelRequest
	GetPlannedEndTime() *string
	SetPlannedStartTime(v string) *RefreshProxyLevelRequest
	GetPlannedStartTime() *string
	SetProxyTargetClass(v string) *RefreshProxyLevelRequest
	GetProxyTargetClass() *string
	SetResourceGroupId(v string) *RefreshProxyLevelRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *RefreshProxyLevelRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RefreshProxyLevelRequest
	GetResourceOwnerId() *int64
}

type RefreshProxyLevelRequest struct {
	// The ID of the PolarDB cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-******************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies whether the execution is initiated by the scheduled task service. Default value: false. If this parameter is set to false and PlannedStartTime is specified, a scheduled task is created. If this parameter is set to true, the task is executed immediately.
	//
	// example:
	//
	// false
	FromTimeService *bool   `json:"FromTimeService,omitempty" xml:"FromTimeService,omitempty"`
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The latest time to execute the proxy specification upgrade. Specify the time in the YYYY-MM-DDThh:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2022-04-28T14:30:00Z
	PlannedEndTime *string `json:"PlannedEndTime,omitempty" xml:"PlannedEndTime,omitempty"`
	// The earliest time to execute the proxy specification upgrade. Specify the time in the YYYY-MM-DDThh:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2022-04-28T14:00:00Z
	PlannedStartTime *string `json:"PlannedStartTime,omitempty" xml:"PlannedStartTime,omitempty"`
	// The target proxy performance level. Valid values: ProxyStandard, ProxyStandardx2, and ProxyStandardx4.
	//
	// This parameter is required.
	//
	// example:
	//
	// ProxyStandard
	ProxyTargetClass *string `json:"ProxyTargetClass,omitempty" xml:"ProxyTargetClass,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-************
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RefreshProxyLevelRequest) String() string {
	return dara.Prettify(s)
}

func (s RefreshProxyLevelRequest) GoString() string {
	return s.String()
}

func (s *RefreshProxyLevelRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *RefreshProxyLevelRequest) GetFromTimeService() *bool {
	return s.FromTimeService
}

func (s *RefreshProxyLevelRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RefreshProxyLevelRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RefreshProxyLevelRequest) GetPlannedEndTime() *string {
	return s.PlannedEndTime
}

func (s *RefreshProxyLevelRequest) GetPlannedStartTime() *string {
	return s.PlannedStartTime
}

func (s *RefreshProxyLevelRequest) GetProxyTargetClass() *string {
	return s.ProxyTargetClass
}

func (s *RefreshProxyLevelRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *RefreshProxyLevelRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RefreshProxyLevelRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RefreshProxyLevelRequest) SetDBClusterId(v string) *RefreshProxyLevelRequest {
	s.DBClusterId = &v
	return s
}

func (s *RefreshProxyLevelRequest) SetFromTimeService(v bool) *RefreshProxyLevelRequest {
	s.FromTimeService = &v
	return s
}

func (s *RefreshProxyLevelRequest) SetOwnerAccount(v string) *RefreshProxyLevelRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RefreshProxyLevelRequest) SetOwnerId(v int64) *RefreshProxyLevelRequest {
	s.OwnerId = &v
	return s
}

func (s *RefreshProxyLevelRequest) SetPlannedEndTime(v string) *RefreshProxyLevelRequest {
	s.PlannedEndTime = &v
	return s
}

func (s *RefreshProxyLevelRequest) SetPlannedStartTime(v string) *RefreshProxyLevelRequest {
	s.PlannedStartTime = &v
	return s
}

func (s *RefreshProxyLevelRequest) SetProxyTargetClass(v string) *RefreshProxyLevelRequest {
	s.ProxyTargetClass = &v
	return s
}

func (s *RefreshProxyLevelRequest) SetResourceGroupId(v string) *RefreshProxyLevelRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *RefreshProxyLevelRequest) SetResourceOwnerAccount(v string) *RefreshProxyLevelRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RefreshProxyLevelRequest) SetResourceOwnerId(v int64) *RefreshProxyLevelRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RefreshProxyLevelRequest) Validate() error {
	return dara.Validate(s)
}
