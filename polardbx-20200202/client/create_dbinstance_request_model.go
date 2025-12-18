// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v bool) *CreateDBInstanceRequest
	GetAutoRenew() *bool
	SetCNNodeCount(v int32) *CreateDBInstanceRequest
	GetCNNodeCount() *int32
	SetClientToken(v string) *CreateDBInstanceRequest
	GetClientToken() *string
	SetCnClass(v string) *CreateDBInstanceRequest
	GetCnClass() *string
	SetDBNodeClass(v string) *CreateDBInstanceRequest
	GetDBNodeClass() *string
	SetDBNodeCount(v int32) *CreateDBInstanceRequest
	GetDBNodeCount() *int32
	SetDNNodeCount(v int32) *CreateDBInstanceRequest
	GetDNNodeCount() *int32
	SetDescription(v string) *CreateDBInstanceRequest
	GetDescription() *string
	SetDnClass(v string) *CreateDBInstanceRequest
	GetDnClass() *string
	SetDnStorageSpace(v string) *CreateDBInstanceRequest
	GetDnStorageSpace() *string
	SetEngineVersion(v string) *CreateDBInstanceRequest
	GetEngineVersion() *string
	SetExtraParams(v map[string]*string) *CreateDBInstanceRequest
	GetExtraParams() map[string]*string
	SetIsColumnarReadDBInstance(v bool) *CreateDBInstanceRequest
	GetIsColumnarReadDBInstance() *bool
	SetIsReadDBInstance(v bool) *CreateDBInstanceRequest
	GetIsReadDBInstance() *bool
	SetNetworkType(v string) *CreateDBInstanceRequest
	GetNetworkType() *string
	SetPayType(v string) *CreateDBInstanceRequest
	GetPayType() *string
	SetPeriod(v string) *CreateDBInstanceRequest
	GetPeriod() *string
	SetPrimaryDBInstanceName(v string) *CreateDBInstanceRequest
	GetPrimaryDBInstanceName() *string
	SetPrimaryZone(v string) *CreateDBInstanceRequest
	GetPrimaryZone() *string
	SetRegionId(v string) *CreateDBInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateDBInstanceRequest
	GetResourceGroupId() *string
	SetSecondaryZone(v string) *CreateDBInstanceRequest
	GetSecondaryZone() *string
	SetSeries(v string) *CreateDBInstanceRequest
	GetSeries() *string
	SetTertiaryZone(v string) *CreateDBInstanceRequest
	GetTertiaryZone() *string
	SetTopologyType(v string) *CreateDBInstanceRequest
	GetTopologyType() *string
	SetUsedTime(v int32) *CreateDBInstanceRequest
	GetUsedTime() *int32
	SetVPCId(v string) *CreateDBInstanceRequest
	GetVPCId() *string
	SetVSwitchId(v string) *CreateDBInstanceRequest
	GetVSwitchId() *string
	SetZoneId(v string) *CreateDBInstanceRequest
	GetZoneId() *string
}

type CreateDBInstanceRequest struct {
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
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DnClass        *string `json:"DnClass,omitempty" xml:"DnClass,omitempty"`
	DnStorageSpace *string `json:"DnStorageSpace,omitempty" xml:"DnStorageSpace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2.0
	EngineVersion            *string            `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	ExtraParams              map[string]*string `json:"ExtraParams,omitempty" xml:"ExtraParams,omitempty"`
	IsColumnarReadDBInstance *bool              `json:"IsColumnarReadDBInstance,omitempty" xml:"IsColumnarReadDBInstance,omitempty"`
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

func (s CreateDBInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateDBInstanceRequest) GetCNNodeCount() *int32 {
	return s.CNNodeCount
}

func (s *CreateDBInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDBInstanceRequest) GetCnClass() *string {
	return s.CnClass
}

func (s *CreateDBInstanceRequest) GetDBNodeClass() *string {
	return s.DBNodeClass
}

func (s *CreateDBInstanceRequest) GetDBNodeCount() *int32 {
	return s.DBNodeCount
}

func (s *CreateDBInstanceRequest) GetDNNodeCount() *int32 {
	return s.DNNodeCount
}

func (s *CreateDBInstanceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDBInstanceRequest) GetDnClass() *string {
	return s.DnClass
}

func (s *CreateDBInstanceRequest) GetDnStorageSpace() *string {
	return s.DnStorageSpace
}

func (s *CreateDBInstanceRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *CreateDBInstanceRequest) GetExtraParams() map[string]*string {
	return s.ExtraParams
}

func (s *CreateDBInstanceRequest) GetIsColumnarReadDBInstance() *bool {
	return s.IsColumnarReadDBInstance
}

func (s *CreateDBInstanceRequest) GetIsReadDBInstance() *bool {
	return s.IsReadDBInstance
}

func (s *CreateDBInstanceRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *CreateDBInstanceRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateDBInstanceRequest) GetPeriod() *string {
	return s.Period
}

func (s *CreateDBInstanceRequest) GetPrimaryDBInstanceName() *string {
	return s.PrimaryDBInstanceName
}

func (s *CreateDBInstanceRequest) GetPrimaryZone() *string {
	return s.PrimaryZone
}

func (s *CreateDBInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDBInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDBInstanceRequest) GetSecondaryZone() *string {
	return s.SecondaryZone
}

func (s *CreateDBInstanceRequest) GetSeries() *string {
	return s.Series
}

func (s *CreateDBInstanceRequest) GetTertiaryZone() *string {
	return s.TertiaryZone
}

func (s *CreateDBInstanceRequest) GetTopologyType() *string {
	return s.TopologyType
}

func (s *CreateDBInstanceRequest) GetUsedTime() *int32 {
	return s.UsedTime
}

func (s *CreateDBInstanceRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *CreateDBInstanceRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateDBInstanceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateDBInstanceRequest) SetAutoRenew(v bool) *CreateDBInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateDBInstanceRequest) SetCNNodeCount(v int32) *CreateDBInstanceRequest {
	s.CNNodeCount = &v
	return s
}

func (s *CreateDBInstanceRequest) SetClientToken(v string) *CreateDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBInstanceRequest) SetCnClass(v string) *CreateDBInstanceRequest {
	s.CnClass = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBNodeClass(v string) *CreateDBInstanceRequest {
	s.DBNodeClass = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBNodeCount(v int32) *CreateDBInstanceRequest {
	s.DBNodeCount = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDNNodeCount(v int32) *CreateDBInstanceRequest {
	s.DNNodeCount = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDescription(v string) *CreateDBInstanceRequest {
	s.Description = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDnClass(v string) *CreateDBInstanceRequest {
	s.DnClass = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDnStorageSpace(v string) *CreateDBInstanceRequest {
	s.DnStorageSpace = &v
	return s
}

func (s *CreateDBInstanceRequest) SetEngineVersion(v string) *CreateDBInstanceRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateDBInstanceRequest) SetExtraParams(v map[string]*string) *CreateDBInstanceRequest {
	s.ExtraParams = v
	return s
}

func (s *CreateDBInstanceRequest) SetIsColumnarReadDBInstance(v bool) *CreateDBInstanceRequest {
	s.IsColumnarReadDBInstance = &v
	return s
}

func (s *CreateDBInstanceRequest) SetIsReadDBInstance(v bool) *CreateDBInstanceRequest {
	s.IsReadDBInstance = &v
	return s
}

func (s *CreateDBInstanceRequest) SetNetworkType(v string) *CreateDBInstanceRequest {
	s.NetworkType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetPayType(v string) *CreateDBInstanceRequest {
	s.PayType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetPeriod(v string) *CreateDBInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateDBInstanceRequest) SetPrimaryDBInstanceName(v string) *CreateDBInstanceRequest {
	s.PrimaryDBInstanceName = &v
	return s
}

func (s *CreateDBInstanceRequest) SetPrimaryZone(v string) *CreateDBInstanceRequest {
	s.PrimaryZone = &v
	return s
}

func (s *CreateDBInstanceRequest) SetRegionId(v string) *CreateDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetResourceGroupId(v string) *CreateDBInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetSecondaryZone(v string) *CreateDBInstanceRequest {
	s.SecondaryZone = &v
	return s
}

func (s *CreateDBInstanceRequest) SetSeries(v string) *CreateDBInstanceRequest {
	s.Series = &v
	return s
}

func (s *CreateDBInstanceRequest) SetTertiaryZone(v string) *CreateDBInstanceRequest {
	s.TertiaryZone = &v
	return s
}

func (s *CreateDBInstanceRequest) SetTopologyType(v string) *CreateDBInstanceRequest {
	s.TopologyType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetUsedTime(v int32) *CreateDBInstanceRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateDBInstanceRequest) SetVPCId(v string) *CreateDBInstanceRequest {
	s.VPCId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetVSwitchId(v string) *CreateDBInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetZoneId(v string) *CreateDBInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateDBInstanceRequest) Validate() error {
	return dara.Validate(s)
}
