// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDiskReplicaPairRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v int64) *CreateDiskReplicaPairRequest
	GetBandwidth() *int64
	SetChargeType(v string) *CreateDiskReplicaPairRequest
	GetChargeType() *string
	SetClientToken(v string) *CreateDiskReplicaPairRequest
	GetClientToken() *string
	SetDescription(v string) *CreateDiskReplicaPairRequest
	GetDescription() *string
	SetDestinationDiskId(v string) *CreateDiskReplicaPairRequest
	GetDestinationDiskId() *string
	SetDestinationRegionId(v string) *CreateDiskReplicaPairRequest
	GetDestinationRegionId() *string
	SetDestinationZoneId(v string) *CreateDiskReplicaPairRequest
	GetDestinationZoneId() *string
	SetDiskId(v string) *CreateDiskReplicaPairRequest
	GetDiskId() *string
	SetEnableRtc(v bool) *CreateDiskReplicaPairRequest
	GetEnableRtc() *bool
	SetPairName(v string) *CreateDiskReplicaPairRequest
	GetPairName() *string
	SetPeriod(v int64) *CreateDiskReplicaPairRequest
	GetPeriod() *int64
	SetPeriodUnit(v string) *CreateDiskReplicaPairRequest
	GetPeriodUnit() *string
	SetRPO(v int64) *CreateDiskReplicaPairRequest
	GetRPO() *int64
	SetRegionId(v string) *CreateDiskReplicaPairRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateDiskReplicaPairRequest
	GetResourceGroupId() *string
	SetSourceZoneId(v string) *CreateDiskReplicaPairRequest
	GetSourceZoneId() *string
	SetTag(v []*CreateDiskReplicaPairRequestTag) *CreateDiskReplicaPairRequest
	GetTag() []*CreateDiskReplicaPairRequestTag
}

type CreateDiskReplicaPairRequest struct {
	// The bandwidth to use to asynchronously replicate data from the primary disk to the secondary disk. Unit: Kbit/s. Valid values:
	//
	// 	- 10240
	//
	// 	- 20480
	//
	// 	- 51200
	//
	// 	- 102400
	//
	// Default value: 10240. When you set the ChargeType parameter to POSTPAY, the Bandwidth parameter is automatically set to 0 and cannot be modified. The value 0 indicates that bandwidth is dynamically allocated based on the volume of data that is asynchronously replicated from the primary disk to the secondary disk.
	//
	// example:
	//
	// 10240
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The billing method of the replication pair. Valid values:
	//
	// 	- PREPAY: subscription
	//
	// 	- POSTPAY: pay-as-you-go
	//
	// Default value: POSTPAY.
	//
	// example:
	//
	// PREPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The client token to ensure the idempotency of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the replication pair. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// This is description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the secondary disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-sa1f82p58p1tdw9g****
	DestinationDiskId *string `json:"DestinationDiskId,omitempty" xml:"DestinationDiskId,omitempty"`
	// The region ID of the secondary disk. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/354276.html) operation to query the most recent list of regions in which async replication is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	DestinationRegionId *string `json:"DestinationRegionId,omitempty" xml:"DestinationRegionId,omitempty"`
	// The zone ID of the secondary disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai-e
	DestinationZoneId *string `json:"DestinationZoneId,omitempty" xml:"DestinationZoneId,omitempty"`
	// The ID of the primary disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-iq80sgp4d0xbk24q****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// Whether to enable replication time control. By default, this parameter is disabled.
	//
	// example:
	//
	// true
	EnableRtc *bool `json:"EnableRtc,omitempty" xml:"EnableRtc,omitempty"`
	// The name of the replication pair. The name must be 2 to 128 characters in length. The name must start with a letter and cannot start with `http://` or `https://`. The name can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// example:
	//
	// TestReplicaPair
	PairName *string `json:"PairName,omitempty" xml:"PairName,omitempty"`
	// The subscription duration of the replication pair. When `ChargeType` is set to PREPAY, this parameter must be specified. Valid values: 1, 2, 3, 6, 12, 24, 36, and 60. The subscription duration unit is specified by `PeriodUnit`.
	//
	// example:
	//
	// 1
	Period *int64 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration of the replication pair. Set the value to Month. Valid value: Month
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The recovery point objective (RPO) of the replication pair. Unit: seconds. Valid value: 900.
	//
	// example:
	//
	// 900
	RPO *int64 `json:"RPO,omitempty" xml:"RPO,omitempty"`
	// The ID of the region in which to create the replication pair.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the replication pair belongs.
	//
	// example:
	//
	// rg-acfmvs****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The zone ID of the primary disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing-f
	SourceZoneId *string `json:"SourceZoneId,omitempty" xml:"SourceZoneId,omitempty"`
	// The tags to add to the replication pair-consistent group. You can specify up to 20 tags.
	Tag []*CreateDiskReplicaPairRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateDiskReplicaPairRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDiskReplicaPairRequest) GoString() string {
	return s.String()
}

func (s *CreateDiskReplicaPairRequest) GetBandwidth() *int64 {
	return s.Bandwidth
}

func (s *CreateDiskReplicaPairRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateDiskReplicaPairRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDiskReplicaPairRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDiskReplicaPairRequest) GetDestinationDiskId() *string {
	return s.DestinationDiskId
}

func (s *CreateDiskReplicaPairRequest) GetDestinationRegionId() *string {
	return s.DestinationRegionId
}

func (s *CreateDiskReplicaPairRequest) GetDestinationZoneId() *string {
	return s.DestinationZoneId
}

func (s *CreateDiskReplicaPairRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *CreateDiskReplicaPairRequest) GetEnableRtc() *bool {
	return s.EnableRtc
}

func (s *CreateDiskReplicaPairRequest) GetPairName() *string {
	return s.PairName
}

func (s *CreateDiskReplicaPairRequest) GetPeriod() *int64 {
	return s.Period
}

func (s *CreateDiskReplicaPairRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateDiskReplicaPairRequest) GetRPO() *int64 {
	return s.RPO
}

func (s *CreateDiskReplicaPairRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDiskReplicaPairRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDiskReplicaPairRequest) GetSourceZoneId() *string {
	return s.SourceZoneId
}

func (s *CreateDiskReplicaPairRequest) GetTag() []*CreateDiskReplicaPairRequestTag {
	return s.Tag
}

func (s *CreateDiskReplicaPairRequest) SetBandwidth(v int64) *CreateDiskReplicaPairRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetChargeType(v string) *CreateDiskReplicaPairRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetClientToken(v string) *CreateDiskReplicaPairRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetDescription(v string) *CreateDiskReplicaPairRequest {
	s.Description = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetDestinationDiskId(v string) *CreateDiskReplicaPairRequest {
	s.DestinationDiskId = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetDestinationRegionId(v string) *CreateDiskReplicaPairRequest {
	s.DestinationRegionId = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetDestinationZoneId(v string) *CreateDiskReplicaPairRequest {
	s.DestinationZoneId = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetDiskId(v string) *CreateDiskReplicaPairRequest {
	s.DiskId = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetEnableRtc(v bool) *CreateDiskReplicaPairRequest {
	s.EnableRtc = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetPairName(v string) *CreateDiskReplicaPairRequest {
	s.PairName = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetPeriod(v int64) *CreateDiskReplicaPairRequest {
	s.Period = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetPeriodUnit(v string) *CreateDiskReplicaPairRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetRPO(v int64) *CreateDiskReplicaPairRequest {
	s.RPO = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetRegionId(v string) *CreateDiskReplicaPairRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetResourceGroupId(v string) *CreateDiskReplicaPairRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetSourceZoneId(v string) *CreateDiskReplicaPairRequest {
	s.SourceZoneId = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetTag(v []*CreateDiskReplicaPairRequestTag) *CreateDiskReplicaPairRequest {
	s.Tag = v
	return s
}

func (s *CreateDiskReplicaPairRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateDiskReplicaPairRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDiskReplicaPairRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateDiskReplicaPairRequestTag) GoString() string {
	return s.String()
}

func (s *CreateDiskReplicaPairRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateDiskReplicaPairRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateDiskReplicaPairRequestTag) SetKey(v string) *CreateDiskReplicaPairRequestTag {
	s.Key = &v
	return s
}

func (s *CreateDiskReplicaPairRequestTag) SetValue(v string) *CreateDiskReplicaPairRequestTag {
	s.Value = &v
	return s
}

func (s *CreateDiskReplicaPairRequestTag) Validate() error {
	return dara.Validate(s)
}
