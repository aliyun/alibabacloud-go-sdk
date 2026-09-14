// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDiskReplicaPairRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v int64) *ModifyDiskReplicaPairRequest
	GetBandwidth() *int64
	SetClientToken(v string) *ModifyDiskReplicaPairRequest
	GetClientToken() *string
	SetDescription(v string) *ModifyDiskReplicaPairRequest
	GetDescription() *string
	SetEnableRtc(v bool) *ModifyDiskReplicaPairRequest
	GetEnableRtc() *bool
	SetPairName(v string) *ModifyDiskReplicaPairRequest
	GetPairName() *string
	SetRPO(v int64) *ModifyDiskReplicaPairRequest
	GetRPO() *int64
	SetRegionId(v string) *ModifyDiskReplicaPairRequest
	GetRegionId() *string
	SetReplicaPairId(v string) *ModifyDiskReplicaPairRequest
	GetReplicaPairId() *string
}

type ModifyDiskReplicaPairRequest struct {
	// The bandwidth for async replication, in Kbps.
	//
	// > This parameter is not yet available.
	//
	// example:
	//
	// 10240
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// A client token to ensure the idempotence of the request. Generate a value for this parameter from your client. Make sure that the value is unique among different requests. The ClientToken value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the replication pair.
	//
	// example:
	//
	// This is description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable replication time control (RTC). Valid values:
	//
	// - false: Disables RTC.
	//
	// - true: Enables RTC.
	//
	// Default value: false.
	//
	// > If a replication pair is part of a replication group, its RTC setting is the same as the setting of the group.
	//
	// example:
	//
	// true
	EnableRtc *bool `json:"EnableRtc,omitempty" xml:"EnableRtc,omitempty"`
	// The name of the replication pair.
	//
	// example:
	//
	// TestReplicaPair
	PairName *string `json:"PairName,omitempty" xml:"PairName,omitempty"`
	// The recovery point objective (RPO) of the replication pair-consistent group. Unit: seconds. Currently, only a value of 900 is supported.
	//
	// example:
	//
	// 900
	RPO *int64 `json:"RPO,omitempty" xml:"RPO,omitempty"`
	// The region ID of the primary or secondary disk in the replication pair. You can call [DescribeRegions](https://help.aliyun.com/document_detail/354276.html) to query the regions that support async replication.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the replication pair.
	//
	// This parameter is required.
	//
	// example:
	//
	// pair-cn-dsa****
	ReplicaPairId *string `json:"ReplicaPairId,omitempty" xml:"ReplicaPairId,omitempty"`
}

func (s ModifyDiskReplicaPairRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDiskReplicaPairRequest) GoString() string {
	return s.String()
}

func (s *ModifyDiskReplicaPairRequest) GetBandwidth() *int64 {
	return s.Bandwidth
}

func (s *ModifyDiskReplicaPairRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyDiskReplicaPairRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyDiskReplicaPairRequest) GetEnableRtc() *bool {
	return s.EnableRtc
}

func (s *ModifyDiskReplicaPairRequest) GetPairName() *string {
	return s.PairName
}

func (s *ModifyDiskReplicaPairRequest) GetRPO() *int64 {
	return s.RPO
}

func (s *ModifyDiskReplicaPairRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDiskReplicaPairRequest) GetReplicaPairId() *string {
	return s.ReplicaPairId
}

func (s *ModifyDiskReplicaPairRequest) SetBandwidth(v int64) *ModifyDiskReplicaPairRequest {
	s.Bandwidth = &v
	return s
}

func (s *ModifyDiskReplicaPairRequest) SetClientToken(v string) *ModifyDiskReplicaPairRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDiskReplicaPairRequest) SetDescription(v string) *ModifyDiskReplicaPairRequest {
	s.Description = &v
	return s
}

func (s *ModifyDiskReplicaPairRequest) SetEnableRtc(v bool) *ModifyDiskReplicaPairRequest {
	s.EnableRtc = &v
	return s
}

func (s *ModifyDiskReplicaPairRequest) SetPairName(v string) *ModifyDiskReplicaPairRequest {
	s.PairName = &v
	return s
}

func (s *ModifyDiskReplicaPairRequest) SetRPO(v int64) *ModifyDiskReplicaPairRequest {
	s.RPO = &v
	return s
}

func (s *ModifyDiskReplicaPairRequest) SetRegionId(v string) *ModifyDiskReplicaPairRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDiskReplicaPairRequest) SetReplicaPairId(v string) *ModifyDiskReplicaPairRequest {
	s.ReplicaPairId = &v
	return s
}

func (s *ModifyDiskReplicaPairRequest) Validate() error {
	return dara.Validate(s)
}
