// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyQosCarRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyQosCarRequest
	GetDescription() *string
	SetLimitType(v string) *ModifyQosCarRequest
	GetLimitType() *string
	SetMaxBandwidthAbs(v int32) *ModifyQosCarRequest
	GetMaxBandwidthAbs() *int32
	SetMaxBandwidthPercent(v int32) *ModifyQosCarRequest
	GetMaxBandwidthPercent() *int32
	SetMinBandwidthAbs(v int32) *ModifyQosCarRequest
	GetMinBandwidthAbs() *int32
	SetMinBandwidthPercent(v int32) *ModifyQosCarRequest
	GetMinBandwidthPercent() *int32
	SetName(v string) *ModifyQosCarRequest
	GetName() *string
	SetOwnerAccount(v string) *ModifyQosCarRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyQosCarRequest
	GetOwnerId() *int64
	SetPercentSourceType(v string) *ModifyQosCarRequest
	GetPercentSourceType() *string
	SetPriority(v int32) *ModifyQosCarRequest
	GetPriority() *int32
	SetQosCarId(v string) *ModifyQosCarRequest
	GetQosCarId() *string
	SetQosId(v string) *ModifyQosCarRequest
	GetQosId() *string
	SetRegionId(v string) *ModifyQosCarRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyQosCarRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyQosCarRequest
	GetResourceOwnerId() *int64
}

type ModifyQosCarRequest struct {
	// The description of the traffic throttling rule.
	//
	// example:
	//
	// Qostest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The type of the traffic throttling rule. Valid values:
	//
	// 	- **Absolute**: throttles traffic based on a specific range of bandwidth.
	//
	// 	- **Percent**: throttles traffic based on a specific range of bandwidth percentage.
	//
	// example:
	//
	// Absolute
	LimitType *string `json:"LimitType,omitempty" xml:"LimitType,omitempty"`
	// The maximum bandwidth value. The value must be an integer. Unit: Mbit /s.
	//
	// This parameter is required when you set **LimitType*	- to **Absolute**.
	//
	// >  The maximum bandwidth value must be greater than the minimum bandwidth value.
	//
	// example:
	//
	// 6
	MaxBandwidthAbs *int32 `json:"MaxBandwidthAbs,omitempty" xml:"MaxBandwidthAbs,omitempty"`
	// The maximum bandwidth percentage. Unit: percent (%). Valid values: **1 to 100**.
	//
	// This parameter is required when you set **LimitType*	- to **Percent**.
	//
	// >  The maximum bandwidth percentage must be greater than the minimum bandwidth percentage.
	//
	// example:
	//
	// 90
	MaxBandwidthPercent *int32 `json:"MaxBandwidthPercent,omitempty" xml:"MaxBandwidthPercent,omitempty"`
	// The minimum bandwidth value. The value must be an integer. Unit: Mbit/s.
	//
	// This parameter is required when you set **LimitType*	- to **Absolute**.
	//
	// example:
	//
	// 2
	MinBandwidthAbs *int32 `json:"MinBandwidthAbs,omitempty" xml:"MinBandwidthAbs,omitempty"`
	// The minimum bandwidth percentage. Unit: percent (%). Valid values: **1 to 100**.
	//
	// This parameter is required when you set **LimitType*	- to **Percent**.
	//
	// example:
	//
	// 20
	MinBandwidthPercent *int32 `json:"MinBandwidthPercent,omitempty" xml:"MinBandwidthPercent,omitempty"`
	// The name of the traffic throttling rule.
	//
	// The name must be 2 to 128 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// nametest
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The type of bandwidth when traffic is throttled based on bandwidth percentage. Valid values:
	//
	// 	- **CcnBandwidth**: Cloud Connect Network (CCN) bandwidth
	//
	// 	- **InternetUpBandwidth**: total Internet bandwidth
	//
	// example:
	//
	// CcnBandwidth
	PercentSourceType *string `json:"PercentSourceType,omitempty" xml:"PercentSourceType,omitempty"`
	// The priority value of the traffic throttling rule. A smaller value specifies a higher priority. If multiple rules have the same priority, the rule that is applied first takes effect. Valid values: **1 to 7**.
	//
	// example:
	//
	// 2
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the traffic throttling rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// qoscar-n5k8g97lihlph****
	QosCarId *string `json:"QosCarId,omitempty" xml:"QosCarId,omitempty"`
	// The ID of the QoS policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// qos-awfxl1adxeqyk****
	QosId *string `json:"QosId,omitempty" xml:"QosId,omitempty"`
	// The ID of the region to which the QoS policy belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyQosCarRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyQosCarRequest) GoString() string {
	return s.String()
}

func (s *ModifyQosCarRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyQosCarRequest) GetLimitType() *string {
	return s.LimitType
}

func (s *ModifyQosCarRequest) GetMaxBandwidthAbs() *int32 {
	return s.MaxBandwidthAbs
}

func (s *ModifyQosCarRequest) GetMaxBandwidthPercent() *int32 {
	return s.MaxBandwidthPercent
}

func (s *ModifyQosCarRequest) GetMinBandwidthAbs() *int32 {
	return s.MinBandwidthAbs
}

func (s *ModifyQosCarRequest) GetMinBandwidthPercent() *int32 {
	return s.MinBandwidthPercent
}

func (s *ModifyQosCarRequest) GetName() *string {
	return s.Name
}

func (s *ModifyQosCarRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyQosCarRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyQosCarRequest) GetPercentSourceType() *string {
	return s.PercentSourceType
}

func (s *ModifyQosCarRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *ModifyQosCarRequest) GetQosCarId() *string {
	return s.QosCarId
}

func (s *ModifyQosCarRequest) GetQosId() *string {
	return s.QosId
}

func (s *ModifyQosCarRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyQosCarRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyQosCarRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyQosCarRequest) SetDescription(v string) *ModifyQosCarRequest {
	s.Description = &v
	return s
}

func (s *ModifyQosCarRequest) SetLimitType(v string) *ModifyQosCarRequest {
	s.LimitType = &v
	return s
}

func (s *ModifyQosCarRequest) SetMaxBandwidthAbs(v int32) *ModifyQosCarRequest {
	s.MaxBandwidthAbs = &v
	return s
}

func (s *ModifyQosCarRequest) SetMaxBandwidthPercent(v int32) *ModifyQosCarRequest {
	s.MaxBandwidthPercent = &v
	return s
}

func (s *ModifyQosCarRequest) SetMinBandwidthAbs(v int32) *ModifyQosCarRequest {
	s.MinBandwidthAbs = &v
	return s
}

func (s *ModifyQosCarRequest) SetMinBandwidthPercent(v int32) *ModifyQosCarRequest {
	s.MinBandwidthPercent = &v
	return s
}

func (s *ModifyQosCarRequest) SetName(v string) *ModifyQosCarRequest {
	s.Name = &v
	return s
}

func (s *ModifyQosCarRequest) SetOwnerAccount(v string) *ModifyQosCarRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyQosCarRequest) SetOwnerId(v int64) *ModifyQosCarRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyQosCarRequest) SetPercentSourceType(v string) *ModifyQosCarRequest {
	s.PercentSourceType = &v
	return s
}

func (s *ModifyQosCarRequest) SetPriority(v int32) *ModifyQosCarRequest {
	s.Priority = &v
	return s
}

func (s *ModifyQosCarRequest) SetQosCarId(v string) *ModifyQosCarRequest {
	s.QosCarId = &v
	return s
}

func (s *ModifyQosCarRequest) SetQosId(v string) *ModifyQosCarRequest {
	s.QosId = &v
	return s
}

func (s *ModifyQosCarRequest) SetRegionId(v string) *ModifyQosCarRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyQosCarRequest) SetResourceOwnerAccount(v string) *ModifyQosCarRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyQosCarRequest) SetResourceOwnerId(v int64) *ModifyQosCarRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyQosCarRequest) Validate() error {
	return dara.Validate(s)
}
