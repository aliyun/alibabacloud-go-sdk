// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQosCarRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateQosCarRequest
	GetDescription() *string
	SetLimitType(v string) *CreateQosCarRequest
	GetLimitType() *string
	SetMaxBandwidthAbs(v int32) *CreateQosCarRequest
	GetMaxBandwidthAbs() *int32
	SetMaxBandwidthPercent(v int32) *CreateQosCarRequest
	GetMaxBandwidthPercent() *int32
	SetMinBandwidthAbs(v int32) *CreateQosCarRequest
	GetMinBandwidthAbs() *int32
	SetMinBandwidthPercent(v int32) *CreateQosCarRequest
	GetMinBandwidthPercent() *int32
	SetName(v string) *CreateQosCarRequest
	GetName() *string
	SetOwnerAccount(v string) *CreateQosCarRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateQosCarRequest
	GetOwnerId() *int64
	SetPercentSourceType(v string) *CreateQosCarRequest
	GetPercentSourceType() *string
	SetPriority(v int32) *CreateQosCarRequest
	GetPriority() *int32
	SetQosId(v string) *CreateQosCarRequest
	GetQosId() *string
	SetRegionId(v string) *CreateQosCarRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateQosCarRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateQosCarRequest
	GetResourceOwnerId() *int64
}

type CreateQosCarRequest struct {
	// The description of the traffic throttling rule.
	//
	// example:
	//
	// Qosdesc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The type of the traffic throttling rule. Valid values:
	//
	// 	- **Absolute**: throttles traffic based on a specific range of bandwidth values.
	//
	// 	- **Percent**: throttles traffic based on a specific range of bandwidth percentage.
	//
	// This parameter is required.
	//
	// example:
	//
	// Absolute
	LimitType *string `json:"LimitType,omitempty" xml:"LimitType,omitempty"`
	// The maximum bandwidth value. The value must be an integer. Unit: Mbit/s.
	//
	// This parameter is returned when **LimitType*	- is set to **Absolute**.
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
	// This parameter is returned when **LimitType*	- is set to **Absolute**.
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
	// The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// nametest
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The type of bandwidth when traffic is throttled based on bandwidth percentage. Valid values:
	//
	// 	- **CcnBandwidth**: CCN bandwidth
	//
	// 	- **InternetUpBandwidth**: total Internet bandwidth
	//
	// example:
	//
	// CcnBandwidth
	PercentSourceType *string `json:"PercentSourceType,omitempty" xml:"PercentSourceType,omitempty"`
	// The priority of the traffic throttling rule.
	//
	// Valid values: **1*	- to **3**. A smaller value indicates a higher priority. If rules have the same priority, the one created the earliest is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the QoS policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// qos-xitd8690ucu8ro****
	QosId *string `json:"QosId,omitempty" xml:"QosId,omitempty"`
	// The ID of the region to which the QoS policy belongs.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/69813.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateQosCarRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateQosCarRequest) GoString() string {
	return s.String()
}

func (s *CreateQosCarRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateQosCarRequest) GetLimitType() *string {
	return s.LimitType
}

func (s *CreateQosCarRequest) GetMaxBandwidthAbs() *int32 {
	return s.MaxBandwidthAbs
}

func (s *CreateQosCarRequest) GetMaxBandwidthPercent() *int32 {
	return s.MaxBandwidthPercent
}

func (s *CreateQosCarRequest) GetMinBandwidthAbs() *int32 {
	return s.MinBandwidthAbs
}

func (s *CreateQosCarRequest) GetMinBandwidthPercent() *int32 {
	return s.MinBandwidthPercent
}

func (s *CreateQosCarRequest) GetName() *string {
	return s.Name
}

func (s *CreateQosCarRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateQosCarRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateQosCarRequest) GetPercentSourceType() *string {
	return s.PercentSourceType
}

func (s *CreateQosCarRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateQosCarRequest) GetQosId() *string {
	return s.QosId
}

func (s *CreateQosCarRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateQosCarRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateQosCarRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateQosCarRequest) SetDescription(v string) *CreateQosCarRequest {
	s.Description = &v
	return s
}

func (s *CreateQosCarRequest) SetLimitType(v string) *CreateQosCarRequest {
	s.LimitType = &v
	return s
}

func (s *CreateQosCarRequest) SetMaxBandwidthAbs(v int32) *CreateQosCarRequest {
	s.MaxBandwidthAbs = &v
	return s
}

func (s *CreateQosCarRequest) SetMaxBandwidthPercent(v int32) *CreateQosCarRequest {
	s.MaxBandwidthPercent = &v
	return s
}

func (s *CreateQosCarRequest) SetMinBandwidthAbs(v int32) *CreateQosCarRequest {
	s.MinBandwidthAbs = &v
	return s
}

func (s *CreateQosCarRequest) SetMinBandwidthPercent(v int32) *CreateQosCarRequest {
	s.MinBandwidthPercent = &v
	return s
}

func (s *CreateQosCarRequest) SetName(v string) *CreateQosCarRequest {
	s.Name = &v
	return s
}

func (s *CreateQosCarRequest) SetOwnerAccount(v string) *CreateQosCarRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateQosCarRequest) SetOwnerId(v int64) *CreateQosCarRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateQosCarRequest) SetPercentSourceType(v string) *CreateQosCarRequest {
	s.PercentSourceType = &v
	return s
}

func (s *CreateQosCarRequest) SetPriority(v int32) *CreateQosCarRequest {
	s.Priority = &v
	return s
}

func (s *CreateQosCarRequest) SetQosId(v string) *CreateQosCarRequest {
	s.QosId = &v
	return s
}

func (s *CreateQosCarRequest) SetRegionId(v string) *CreateQosCarRequest {
	s.RegionId = &v
	return s
}

func (s *CreateQosCarRequest) SetResourceOwnerAccount(v string) *CreateQosCarRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateQosCarRequest) SetResourceOwnerId(v int64) *CreateQosCarRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateQosCarRequest) Validate() error {
	return dara.Validate(s)
}
