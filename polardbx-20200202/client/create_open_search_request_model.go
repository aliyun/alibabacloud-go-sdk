// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOpenSearchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v bool) *CreateOpenSearchRequest
	GetAutoRenew() *bool
	SetClientToken(v string) *CreateOpenSearchRequest
	GetClientToken() *string
	SetDBInstanceDescription(v string) *CreateOpenSearchRequest
	GetDBInstanceDescription() *string
	SetDBNodeClass(v string) *CreateOpenSearchRequest
	GetDBNodeClass() *string
	SetEngineVersion(v string) *CreateOpenSearchRequest
	GetEngineVersion() *string
	SetInstanceSpec(v string) *CreateOpenSearchRequest
	GetInstanceSpec() *string
	SetNodeCount(v int32) *CreateOpenSearchRequest
	GetNodeCount() *int32
	SetPayType(v string) *CreateOpenSearchRequest
	GetPayType() *string
	SetPeriod(v string) *CreateOpenSearchRequest
	GetPeriod() *string
	SetRegionId(v string) *CreateOpenSearchRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateOpenSearchRequest
	GetResourceGroupId() *string
	SetStorageSpace(v int32) *CreateOpenSearchRequest
	GetStorageSpace() *int32
	SetStorageType(v string) *CreateOpenSearchRequest
	GetStorageType() *string
	SetTopologyType(v string) *CreateOpenSearchRequest
	GetTopologyType() *string
	SetUsedTime(v int32) *CreateOpenSearchRequest
	GetUsedTime() *int32
	SetVPCId(v string) *CreateOpenSearchRequest
	GetVPCId() *string
	SetVSwitchId(v string) *CreateOpenSearchRequest
	GetVSwitchId() *string
	SetZone2(v string) *CreateOpenSearchRequest
	GetZone2() *string
	SetZone3(v string) *CreateOpenSearchRequest
	GetZone3() *string
	SetZoneId(v string) *CreateOpenSearchRequest
	GetZoneId() *string
}

type CreateOpenSearchRequest struct {
	// Specifies whether to enable auto-renewal. Default value: true.
	//
	// - **true**: enabled.
	//
	// - **false**: disabled.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The client token used to ensure the idempotence of the request. Use a different value for each creation request.
	//
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// TestInstance
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The node specifications code of PolarDBX Search data nodes. Available specifications depend on the region and sales configuration. Use a PolarDBX Search specification code that is available for purchase in the current region.
	//
	// This parameter is required.
	//
	// example:
	//
	// opensearch.sn2ne.large.1
	DBNodeClass *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	// The PolarDBX Search DPI engine version. The value is fixed to 3.0. If this parameter is not specified, the default value 3.0 is used.
	//
	// example:
	//
	// 3.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// A compatible parameter that does not take effect. Use DBNodeClass to specify the PolarDBX Search data node specifications.
	//
	// example:
	//
	// opensearch.sn2ne.large.1
	InstanceSpec *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	// The number of PolarDBX Search data nodes. The value must be a positive integer and a multiple of the number of selected zones.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	NodeCount *int32 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// The billing method of the instance.
	//
	// - **PREPAY**: subscription.
	//
	// - **POSTPAY**: pay-as-you-go.
	//
	// This parameter is required.
	//
	// example:
	//
	// PREPAY
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The billing cycle. Valid values for subscription: Year and Month. Default value for pay-as-you-go: Hour.
	//
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The region in which the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID. This parameter can be left empty. This parameter is not supported.
	//
	// example:
	//
	// rg-xxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The storage space per node, in GB. The value must be a positive integer.
	//
	// example:
	//
	// 20
	StorageSpace *int32 `json:"StorageSpace,omitempty" xml:"StorageSpace,omitempty"`
	// The storage type. Default value: cloud_auto.
	//
	// example:
	//
	// cloud_auto
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The topology type. Valid values:
	//
	// - **1azone**: single active zone.
	//
	// - **3azones**: three active zones.
	//
	// example:
	//
	// 3azones
	TopologyType *string `json:"TopologyType,omitempty" xml:"TopologyType,omitempty"`
	// The subscription duration. Specify the number of months or years for prepaid instances.
	//
	// > When Period is set to Year, valid values for this parameter are 1, 2, and 3.
	//
	// example:
	//
	// 1
	UsedTime *int32 `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// VPC ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-*****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-*********
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The second zone. This parameter is required when TopologyType is set to 3azones. The value cannot be the same as other zones.
	//
	// example:
	//
	// cn-hangzhou-i
	Zone2 *string `json:"Zone2,omitempty" xml:"Zone2,omitempty"`
	// The third zone. This parameter is required when TopologyType is set to 3azones. The value cannot be the same as other zones.
	//
	// example:
	//
	// cn-hangzhou-j
	Zone3 *string `json:"Zone3,omitempty" xml:"Zone3,omitempty"`
	// The zone of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateOpenSearchRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOpenSearchRequest) GoString() string {
	return s.String()
}

func (s *CreateOpenSearchRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateOpenSearchRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateOpenSearchRequest) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *CreateOpenSearchRequest) GetDBNodeClass() *string {
	return s.DBNodeClass
}

func (s *CreateOpenSearchRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *CreateOpenSearchRequest) GetInstanceSpec() *string {
	return s.InstanceSpec
}

func (s *CreateOpenSearchRequest) GetNodeCount() *int32 {
	return s.NodeCount
}

func (s *CreateOpenSearchRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateOpenSearchRequest) GetPeriod() *string {
	return s.Period
}

func (s *CreateOpenSearchRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateOpenSearchRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateOpenSearchRequest) GetStorageSpace() *int32 {
	return s.StorageSpace
}

func (s *CreateOpenSearchRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *CreateOpenSearchRequest) GetTopologyType() *string {
	return s.TopologyType
}

func (s *CreateOpenSearchRequest) GetUsedTime() *int32 {
	return s.UsedTime
}

func (s *CreateOpenSearchRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *CreateOpenSearchRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateOpenSearchRequest) GetZone2() *string {
	return s.Zone2
}

func (s *CreateOpenSearchRequest) GetZone3() *string {
	return s.Zone3
}

func (s *CreateOpenSearchRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateOpenSearchRequest) SetAutoRenew(v bool) *CreateOpenSearchRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateOpenSearchRequest) SetClientToken(v string) *CreateOpenSearchRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateOpenSearchRequest) SetDBInstanceDescription(v string) *CreateOpenSearchRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *CreateOpenSearchRequest) SetDBNodeClass(v string) *CreateOpenSearchRequest {
	s.DBNodeClass = &v
	return s
}

func (s *CreateOpenSearchRequest) SetEngineVersion(v string) *CreateOpenSearchRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateOpenSearchRequest) SetInstanceSpec(v string) *CreateOpenSearchRequest {
	s.InstanceSpec = &v
	return s
}

func (s *CreateOpenSearchRequest) SetNodeCount(v int32) *CreateOpenSearchRequest {
	s.NodeCount = &v
	return s
}

func (s *CreateOpenSearchRequest) SetPayType(v string) *CreateOpenSearchRequest {
	s.PayType = &v
	return s
}

func (s *CreateOpenSearchRequest) SetPeriod(v string) *CreateOpenSearchRequest {
	s.Period = &v
	return s
}

func (s *CreateOpenSearchRequest) SetRegionId(v string) *CreateOpenSearchRequest {
	s.RegionId = &v
	return s
}

func (s *CreateOpenSearchRequest) SetResourceGroupId(v string) *CreateOpenSearchRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateOpenSearchRequest) SetStorageSpace(v int32) *CreateOpenSearchRequest {
	s.StorageSpace = &v
	return s
}

func (s *CreateOpenSearchRequest) SetStorageType(v string) *CreateOpenSearchRequest {
	s.StorageType = &v
	return s
}

func (s *CreateOpenSearchRequest) SetTopologyType(v string) *CreateOpenSearchRequest {
	s.TopologyType = &v
	return s
}

func (s *CreateOpenSearchRequest) SetUsedTime(v int32) *CreateOpenSearchRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateOpenSearchRequest) SetVPCId(v string) *CreateOpenSearchRequest {
	s.VPCId = &v
	return s
}

func (s *CreateOpenSearchRequest) SetVSwitchId(v string) *CreateOpenSearchRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateOpenSearchRequest) SetZone2(v string) *CreateOpenSearchRequest {
	s.Zone2 = &v
	return s
}

func (s *CreateOpenSearchRequest) SetZone3(v string) *CreateOpenSearchRequest {
	s.Zone3 = &v
	return s
}

func (s *CreateOpenSearchRequest) SetZoneId(v string) *CreateOpenSearchRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateOpenSearchRequest) Validate() error {
	return dara.Validate(s)
}
