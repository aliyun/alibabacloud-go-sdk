// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v bool) *CreateDBInstanceShrinkRequest
	GetAutoRenew() *bool
	SetCNNodeCount(v int32) *CreateDBInstanceShrinkRequest
	GetCNNodeCount() *int32
	SetClientToken(v string) *CreateDBInstanceShrinkRequest
	GetClientToken() *string
	SetCnClass(v string) *CreateDBInstanceShrinkRequest
	GetCnClass() *string
	SetDBNodeClass(v string) *CreateDBInstanceShrinkRequest
	GetDBNodeClass() *string
	SetDBNodeCount(v int32) *CreateDBInstanceShrinkRequest
	GetDBNodeCount() *int32
	SetDNNodeCount(v int32) *CreateDBInstanceShrinkRequest
	GetDNNodeCount() *int32
	SetDnClass(v string) *CreateDBInstanceShrinkRequest
	GetDnClass() *string
	SetDnStorageSpace(v string) *CreateDBInstanceShrinkRequest
	GetDnStorageSpace() *string
	SetEngineVersion(v string) *CreateDBInstanceShrinkRequest
	GetEngineVersion() *string
	SetExtraParamsShrink(v string) *CreateDBInstanceShrinkRequest
	GetExtraParamsShrink() *string
	SetIsColumnarReadDBInstance(v bool) *CreateDBInstanceShrinkRequest
	GetIsColumnarReadDBInstance() *bool
	SetIsReadDBInstance(v bool) *CreateDBInstanceShrinkRequest
	GetIsReadDBInstance() *bool
	SetNetworkType(v string) *CreateDBInstanceShrinkRequest
	GetNetworkType() *string
	SetPayType(v string) *CreateDBInstanceShrinkRequest
	GetPayType() *string
	SetPeriod(v string) *CreateDBInstanceShrinkRequest
	GetPeriod() *string
	SetPrimaryDBInstanceName(v string) *CreateDBInstanceShrinkRequest
	GetPrimaryDBInstanceName() *string
	SetPrimaryZone(v string) *CreateDBInstanceShrinkRequest
	GetPrimaryZone() *string
	SetRegionId(v string) *CreateDBInstanceShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateDBInstanceShrinkRequest
	GetResourceGroupId() *string
	SetSecondaryZone(v string) *CreateDBInstanceShrinkRequest
	GetSecondaryZone() *string
	SetSeries(v string) *CreateDBInstanceShrinkRequest
	GetSeries() *string
	SetTertiaryZone(v string) *CreateDBInstanceShrinkRequest
	GetTertiaryZone() *string
	SetTopologyType(v string) *CreateDBInstanceShrinkRequest
	GetTopologyType() *string
	SetUsedTime(v int32) *CreateDBInstanceShrinkRequest
	GetUsedTime() *int32
	SetVPCId(v string) *CreateDBInstanceShrinkRequest
	GetVPCId() *string
	SetVSwitchId(v string) *CreateDBInstanceShrinkRequest
	GetVSwitchId() *string
	SetZoneId(v string) *CreateDBInstanceShrinkRequest
	GetZoneId() *string
}

type CreateDBInstanceShrinkRequest struct {
	// example:
	//
	// true
	AutoRenew   *bool  `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	CNNodeCount *int32 `json:"CNNodeCount,omitempty" xml:"CNNodeCount,omitempty"`
	// example:
	//
	// xxxxxx-xxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CnClass     *string `json:"CnClass,omitempty" xml:"CnClass,omitempty"`
	// example:
	//
	// polarx.x4.2xlarge.2d
	DBNodeClass *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	// example:
	//
	// 2
	DBNodeCount    *int32  `json:"DBNodeCount,omitempty" xml:"DBNodeCount,omitempty"`
	DNNodeCount    *int32  `json:"DNNodeCount,omitempty" xml:"DNNodeCount,omitempty"`
	DnClass        *string `json:"DnClass,omitempty" xml:"DnClass,omitempty"`
	DnStorageSpace *string `json:"DnStorageSpace,omitempty" xml:"DnStorageSpace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2.0
	EngineVersion            *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	ExtraParamsShrink        *string `json:"ExtraParams,omitempty" xml:"ExtraParams,omitempty"`
	IsColumnarReadDBInstance *bool   `json:"IsColumnarReadDBInstance,omitempty" xml:"IsColumnarReadDBInstance,omitempty"`
	// example:
	//
	// false
	IsReadDBInstance *bool `json:"IsReadDBInstance,omitempty" xml:"IsReadDBInstance,omitempty"`
	// example:
	//
	// vpc
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PREPAY
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// pxc-*********
	PrimaryDBInstanceName *string `json:"PrimaryDBInstanceName,omitempty" xml:"PrimaryDBInstanceName,omitempty"`
	// example:
	//
	// cn-shenzhen-e
	PrimaryZone *string `json:"PrimaryZone,omitempty" xml:"PrimaryZone,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// null
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// cn-shenzhen-a
	SecondaryZone *string `json:"SecondaryZone,omitempty" xml:"SecondaryZone,omitempty"`
	Series        *string `json:"Series,omitempty" xml:"Series,omitempty"`
	// example:
	//
	// cn-shenzhen-e
	TertiaryZone *string `json:"TertiaryZone,omitempty" xml:"TertiaryZone,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3azones
	TopologyType *string `json:"TopologyType,omitempty" xml:"TopologyType,omitempty"`
	// example:
	//
	// 1
	UsedTime *int32 `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// VPC ID。
	//
	// example:
	//
	// vpc-*****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// example:
	//
	// vsw-*********
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// cn-hangzhou-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateDBInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceShrinkRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateDBInstanceShrinkRequest) GetCNNodeCount() *int32 {
	return s.CNNodeCount
}

func (s *CreateDBInstanceShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDBInstanceShrinkRequest) GetCnClass() *string {
	return s.CnClass
}

func (s *CreateDBInstanceShrinkRequest) GetDBNodeClass() *string {
	return s.DBNodeClass
}

func (s *CreateDBInstanceShrinkRequest) GetDBNodeCount() *int32 {
	return s.DBNodeCount
}

func (s *CreateDBInstanceShrinkRequest) GetDNNodeCount() *int32 {
	return s.DNNodeCount
}

func (s *CreateDBInstanceShrinkRequest) GetDnClass() *string {
	return s.DnClass
}

func (s *CreateDBInstanceShrinkRequest) GetDnStorageSpace() *string {
	return s.DnStorageSpace
}

func (s *CreateDBInstanceShrinkRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *CreateDBInstanceShrinkRequest) GetExtraParamsShrink() *string {
	return s.ExtraParamsShrink
}

func (s *CreateDBInstanceShrinkRequest) GetIsColumnarReadDBInstance() *bool {
	return s.IsColumnarReadDBInstance
}

func (s *CreateDBInstanceShrinkRequest) GetIsReadDBInstance() *bool {
	return s.IsReadDBInstance
}

func (s *CreateDBInstanceShrinkRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *CreateDBInstanceShrinkRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateDBInstanceShrinkRequest) GetPeriod() *string {
	return s.Period
}

func (s *CreateDBInstanceShrinkRequest) GetPrimaryDBInstanceName() *string {
	return s.PrimaryDBInstanceName
}

func (s *CreateDBInstanceShrinkRequest) GetPrimaryZone() *string {
	return s.PrimaryZone
}

func (s *CreateDBInstanceShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDBInstanceShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDBInstanceShrinkRequest) GetSecondaryZone() *string {
	return s.SecondaryZone
}

func (s *CreateDBInstanceShrinkRequest) GetSeries() *string {
	return s.Series
}

func (s *CreateDBInstanceShrinkRequest) GetTertiaryZone() *string {
	return s.TertiaryZone
}

func (s *CreateDBInstanceShrinkRequest) GetTopologyType() *string {
	return s.TopologyType
}

func (s *CreateDBInstanceShrinkRequest) GetUsedTime() *int32 {
	return s.UsedTime
}

func (s *CreateDBInstanceShrinkRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *CreateDBInstanceShrinkRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateDBInstanceShrinkRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateDBInstanceShrinkRequest) SetAutoRenew(v bool) *CreateDBInstanceShrinkRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetCNNodeCount(v int32) *CreateDBInstanceShrinkRequest {
	s.CNNodeCount = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetClientToken(v string) *CreateDBInstanceShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetCnClass(v string) *CreateDBInstanceShrinkRequest {
	s.CnClass = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetDBNodeClass(v string) *CreateDBInstanceShrinkRequest {
	s.DBNodeClass = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetDBNodeCount(v int32) *CreateDBInstanceShrinkRequest {
	s.DBNodeCount = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetDNNodeCount(v int32) *CreateDBInstanceShrinkRequest {
	s.DNNodeCount = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetDnClass(v string) *CreateDBInstanceShrinkRequest {
	s.DnClass = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetDnStorageSpace(v string) *CreateDBInstanceShrinkRequest {
	s.DnStorageSpace = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetEngineVersion(v string) *CreateDBInstanceShrinkRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetExtraParamsShrink(v string) *CreateDBInstanceShrinkRequest {
	s.ExtraParamsShrink = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetIsColumnarReadDBInstance(v bool) *CreateDBInstanceShrinkRequest {
	s.IsColumnarReadDBInstance = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetIsReadDBInstance(v bool) *CreateDBInstanceShrinkRequest {
	s.IsReadDBInstance = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetNetworkType(v string) *CreateDBInstanceShrinkRequest {
	s.NetworkType = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetPayType(v string) *CreateDBInstanceShrinkRequest {
	s.PayType = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetPeriod(v string) *CreateDBInstanceShrinkRequest {
	s.Period = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetPrimaryDBInstanceName(v string) *CreateDBInstanceShrinkRequest {
	s.PrimaryDBInstanceName = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetPrimaryZone(v string) *CreateDBInstanceShrinkRequest {
	s.PrimaryZone = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetRegionId(v string) *CreateDBInstanceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetResourceGroupId(v string) *CreateDBInstanceShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetSecondaryZone(v string) *CreateDBInstanceShrinkRequest {
	s.SecondaryZone = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetSeries(v string) *CreateDBInstanceShrinkRequest {
	s.Series = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetTertiaryZone(v string) *CreateDBInstanceShrinkRequest {
	s.TertiaryZone = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetTopologyType(v string) *CreateDBInstanceShrinkRequest {
	s.TopologyType = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetUsedTime(v int32) *CreateDBInstanceShrinkRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetVPCId(v string) *CreateDBInstanceShrinkRequest {
	s.VPCId = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetVSwitchId(v string) *CreateDBInstanceShrinkRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetZoneId(v string) *CreateDBInstanceShrinkRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
