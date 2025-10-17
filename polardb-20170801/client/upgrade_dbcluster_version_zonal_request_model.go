// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeDBClusterVersionZonalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpgradeDBClusterVersionZonalRequest
	GetClientToken() *string
	SetDBClusterId(v string) *UpgradeDBClusterVersionZonalRequest
	GetDBClusterId() *string
	SetFromTimeService(v bool) *UpgradeDBClusterVersionZonalRequest
	GetFromTimeService() *bool
	SetOwnerAccount(v string) *UpgradeDBClusterVersionZonalRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpgradeDBClusterVersionZonalRequest
	GetOwnerId() *int64
	SetPlannedEndTime(v string) *UpgradeDBClusterVersionZonalRequest
	GetPlannedEndTime() *string
	SetPlannedStartTime(v string) *UpgradeDBClusterVersionZonalRequest
	GetPlannedStartTime() *string
	SetResourceOwnerAccount(v string) *UpgradeDBClusterVersionZonalRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpgradeDBClusterVersionZonalRequest
	GetResourceOwnerId() *int64
	SetTargetDBRevisionVersionCode(v string) *UpgradeDBClusterVersionZonalRequest
	GetTargetDBRevisionVersionCode() *string
	SetTargetProxyRevisionVersionCode(v string) *UpgradeDBClusterVersionZonalRequest
	GetTargetProxyRevisionVersionCode() *string
	SetUpgradeLabel(v string) *UpgradeDBClusterVersionZonalRequest
	GetUpgradeLabel() *string
	SetUpgradePolicy(v string) *UpgradeDBClusterVersionZonalRequest
	GetUpgradePolicy() *string
	SetUpgradeType(v string) *UpgradeDBClusterVersionZonalRequest
	GetUpgradeType() *string
}

type UpgradeDBClusterVersionZonalRequest struct {
	// example:
	//
	// 6000170000591aed949d0f54a343f1a4233c1e7d1c5c******
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pc-a************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// false
	FromTimeService *bool   `json:"FromTimeService,omitempty" xml:"FromTimeService,omitempty"`
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 2021-01-14T09:30:00Z
	PlannedEndTime *string `json:"PlannedEndTime,omitempty" xml:"PlannedEndTime,omitempty"`
	// example:
	//
	// 2022-04-28T14:00:00Z
	PlannedStartTime     *string `json:"PlannedStartTime,omitempty" xml:"PlannedStartTime,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 20230707
	TargetDBRevisionVersionCode *string `json:"TargetDBRevisionVersionCode,omitempty" xml:"TargetDBRevisionVersionCode,omitempty"`
	// example:
	//
	// 20240702
	TargetProxyRevisionVersionCode *string `json:"TargetProxyRevisionVersionCode,omitempty" xml:"TargetProxyRevisionVersionCode,omitempty"`
	// example:
	//
	// INNOVATE
	UpgradeLabel *string `json:"UpgradeLabel,omitempty" xml:"UpgradeLabel,omitempty"`
	// example:
	//
	// HOT
	UpgradePolicy *string `json:"UpgradePolicy,omitempty" xml:"UpgradePolicy,omitempty"`
	// example:
	//
	// PROXY
	UpgradeType *string `json:"UpgradeType,omitempty" xml:"UpgradeType,omitempty"`
}

func (s UpgradeDBClusterVersionZonalRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeDBClusterVersionZonalRequest) GoString() string {
	return s.String()
}

func (s *UpgradeDBClusterVersionZonalRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpgradeDBClusterVersionZonalRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *UpgradeDBClusterVersionZonalRequest) GetFromTimeService() *bool {
	return s.FromTimeService
}

func (s *UpgradeDBClusterVersionZonalRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpgradeDBClusterVersionZonalRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpgradeDBClusterVersionZonalRequest) GetPlannedEndTime() *string {
	return s.PlannedEndTime
}

func (s *UpgradeDBClusterVersionZonalRequest) GetPlannedStartTime() *string {
	return s.PlannedStartTime
}

func (s *UpgradeDBClusterVersionZonalRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpgradeDBClusterVersionZonalRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpgradeDBClusterVersionZonalRequest) GetTargetDBRevisionVersionCode() *string {
	return s.TargetDBRevisionVersionCode
}

func (s *UpgradeDBClusterVersionZonalRequest) GetTargetProxyRevisionVersionCode() *string {
	return s.TargetProxyRevisionVersionCode
}

func (s *UpgradeDBClusterVersionZonalRequest) GetUpgradeLabel() *string {
	return s.UpgradeLabel
}

func (s *UpgradeDBClusterVersionZonalRequest) GetUpgradePolicy() *string {
	return s.UpgradePolicy
}

func (s *UpgradeDBClusterVersionZonalRequest) GetUpgradeType() *string {
	return s.UpgradeType
}

func (s *UpgradeDBClusterVersionZonalRequest) SetClientToken(v string) *UpgradeDBClusterVersionZonalRequest {
	s.ClientToken = &v
	return s
}

func (s *UpgradeDBClusterVersionZonalRequest) SetDBClusterId(v string) *UpgradeDBClusterVersionZonalRequest {
	s.DBClusterId = &v
	return s
}

func (s *UpgradeDBClusterVersionZonalRequest) SetFromTimeService(v bool) *UpgradeDBClusterVersionZonalRequest {
	s.FromTimeService = &v
	return s
}

func (s *UpgradeDBClusterVersionZonalRequest) SetOwnerAccount(v string) *UpgradeDBClusterVersionZonalRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpgradeDBClusterVersionZonalRequest) SetOwnerId(v int64) *UpgradeDBClusterVersionZonalRequest {
	s.OwnerId = &v
	return s
}

func (s *UpgradeDBClusterVersionZonalRequest) SetPlannedEndTime(v string) *UpgradeDBClusterVersionZonalRequest {
	s.PlannedEndTime = &v
	return s
}

func (s *UpgradeDBClusterVersionZonalRequest) SetPlannedStartTime(v string) *UpgradeDBClusterVersionZonalRequest {
	s.PlannedStartTime = &v
	return s
}

func (s *UpgradeDBClusterVersionZonalRequest) SetResourceOwnerAccount(v string) *UpgradeDBClusterVersionZonalRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpgradeDBClusterVersionZonalRequest) SetResourceOwnerId(v int64) *UpgradeDBClusterVersionZonalRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpgradeDBClusterVersionZonalRequest) SetTargetDBRevisionVersionCode(v string) *UpgradeDBClusterVersionZonalRequest {
	s.TargetDBRevisionVersionCode = &v
	return s
}

func (s *UpgradeDBClusterVersionZonalRequest) SetTargetProxyRevisionVersionCode(v string) *UpgradeDBClusterVersionZonalRequest {
	s.TargetProxyRevisionVersionCode = &v
	return s
}

func (s *UpgradeDBClusterVersionZonalRequest) SetUpgradeLabel(v string) *UpgradeDBClusterVersionZonalRequest {
	s.UpgradeLabel = &v
	return s
}

func (s *UpgradeDBClusterVersionZonalRequest) SetUpgradePolicy(v string) *UpgradeDBClusterVersionZonalRequest {
	s.UpgradePolicy = &v
	return s
}

func (s *UpgradeDBClusterVersionZonalRequest) SetUpgradeType(v string) *UpgradeDBClusterVersionZonalRequest {
	s.UpgradeType = &v
	return s
}

func (s *UpgradeDBClusterVersionZonalRequest) Validate() error {
	return dara.Validate(s)
}
