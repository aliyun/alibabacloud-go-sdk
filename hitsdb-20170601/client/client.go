// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateHiTSDBInstanceRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AppKey               *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	InstanceName         *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceAlias        *string `json:"InstanceAlias,omitempty" xml:"InstanceAlias,omitempty"`
	InstanceClass        *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	InstanceStorage      *string `json:"InstanceStorage,omitempty" xml:"InstanceStorage,omitempty"`
	PayType              *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	VPCId                *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId            *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	MaxTimelineLimit     *string `json:"MaxTimelineLimit,omitempty" xml:"MaxTimelineLimit,omitempty"`
	InstanceTps          *string `json:"InstanceTps,omitempty" xml:"InstanceTps,omitempty"`
	EngineType           *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	MaxSeriesPerDatabase *string `json:"MaxSeriesPerDatabase,omitempty" xml:"MaxSeriesPerDatabase,omitempty"`
	MaxDatabaseLimit     *string `json:"MaxDatabaseLimit,omitempty" xml:"MaxDatabaseLimit,omitempty"`
	PricingCycle         *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	Duration             *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	TSDBVersion          *string `json:"TSDBVersion,omitempty" xml:"TSDBVersion,omitempty"`
	DiskCategory         *string `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateHiTSDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateHiTSDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateHiTSDBInstanceRequest) SetSecurityToken(v string) *CreateHiTSDBInstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateHiTSDBInstanceRequest) SetOwnerId(v int64) *CreateHiTSDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateHiTSDBInstanceRequest) SetResourceOwnerAccount(v string) *CreateHiTSDBInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateHiTSDBInstanceRequest) SetResourceOwnerId(v int64) *CreateHiTSDBInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateHiTSDBInstanceRequest) SetOwnerAccount(v string) *CreateHiTSDBInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateHiTSDBInstanceRequest) SetRegionId(v string) *CreateHiTSDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateHiTSDBInstanceRequest) SetAppKey(v string) *CreateHiTSDBInstanceRequest {
	s.AppKey = &v
	return s
}

func (s *CreateHiTSDBInstanceRequest) SetZoneId(v string) *CreateHiTSDBInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateHiTSDBInstanceRequest) SetInstanceName(v string) *CreateHiTSDBInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateHiTSDBInstanceRequest) SetInstanceAlias(v string) *CreateHiTSDBInstanceRequest {
	s.InstanceAlias = &v
	return s
}

func (s *CreateHiTSDBInstanceRequest) SetInstanceClass(v string) *CreateHiTSDBInstanceRequest {
	s.InstanceClass = &v
	return s
}

func (s *CreateHiTSDBInstanceRequest) SetInstanceStorage(v string) *CreateHiTSDBInstanceRequest {
	s.InstanceStorage = &v
	return s
}

func (s *CreateHiTSDBInstanceRequest) SetPayType(v string) *CreateHiTSDBInstanceRequest {
	s.PayType = &v
	return s
}

func (s *CreateHiTSDBInstanceRequest) SetVPCId(v string) *CreateHiTSDBInstanceRequest {
	s.VPCId = &v
	return s
}

func (s *CreateHiTSDBInstanceRequest) SetVSwitchId(v string) *CreateHiTSDBInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateHiTSDBInstanceRequest) SetMaxTimelineLimit(v string) *CreateHiTSDBInstanceRequest {
	s.MaxTimelineLimit = &v
	return s
}

func (s *CreateHiTSDBInstanceRequest) SetInstanceTps(v string) *CreateHiTSDBInstanceRequest {
	s.InstanceTps = &v
	return s
}

func (s *CreateHiTSDBInstanceRequest) SetEngineType(v string) *CreateHiTSDBInstanceRequest {
	s.EngineType = &v
	return s
}

func (s *CreateHiTSDBInstanceRequest) SetMaxSeriesPerDatabase(v string) *CreateHiTSDBInstanceRequest {
	s.MaxSeriesPerDatabase = &v
	return s
}

func (s *CreateHiTSDBInstanceRequest) SetMaxDatabaseLimit(v string) *CreateHiTSDBInstanceRequest {
	s.MaxDatabaseLimit = &v
	return s
}

func (s *CreateHiTSDBInstanceRequest) SetPricingCycle(v string) *CreateHiTSDBInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateHiTSDBInstanceRequest) SetDuration(v string) *CreateHiTSDBInstanceRequest {
	s.Duration = &v
	return s
}

func (s *CreateHiTSDBInstanceRequest) SetTSDBVersion(v string) *CreateHiTSDBInstanceRequest {
	s.TSDBVersion = &v
	return s
}

func (s *CreateHiTSDBInstanceRequest) SetDiskCategory(v string) *CreateHiTSDBInstanceRequest {
	s.DiskCategory = &v
	return s
}

func (s *CreateHiTSDBInstanceRequest) SetClientToken(v string) *CreateHiTSDBInstanceRequest {
	s.ClientToken = &v
	return s
}

type CreateHiTSDBInstanceResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OrderId    *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateHiTSDBInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateHiTSDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHiTSDBInstanceResponseBody) SetRequestId(v string) *CreateHiTSDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHiTSDBInstanceResponseBody) SetInstanceId(v string) *CreateHiTSDBInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateHiTSDBInstanceResponseBody) SetOrderId(v int64) *CreateHiTSDBInstanceResponseBody {
	s.OrderId = &v
	return s
}

type CreateHiTSDBInstanceResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateHiTSDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateHiTSDBInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateHiTSDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateHiTSDBInstanceResponse) SetHeaders(v map[string]*string) *CreateHiTSDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateHiTSDBInstanceResponse) SetBody(v *CreateHiTSDBInstanceResponseBody) *CreateHiTSDBInstanceResponse {
	s.Body = v
	return s
}

type DeleteHiTSDBInstanceRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	AppKey               *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteHiTSDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteHiTSDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteHiTSDBInstanceRequest) SetSecurityToken(v string) *DeleteHiTSDBInstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteHiTSDBInstanceRequest) SetOwnerId(v int64) *DeleteHiTSDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteHiTSDBInstanceRequest) SetResourceOwnerAccount(v string) *DeleteHiTSDBInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteHiTSDBInstanceRequest) SetResourceOwnerId(v int64) *DeleteHiTSDBInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteHiTSDBInstanceRequest) SetOwnerAccount(v string) *DeleteHiTSDBInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteHiTSDBInstanceRequest) SetAppKey(v string) *DeleteHiTSDBInstanceRequest {
	s.AppKey = &v
	return s
}

func (s *DeleteHiTSDBInstanceRequest) SetInstanceId(v string) *DeleteHiTSDBInstanceRequest {
	s.InstanceId = &v
	return s
}

type DeleteHiTSDBInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteHiTSDBInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteHiTSDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHiTSDBInstanceResponseBody) SetRequestId(v string) *DeleteHiTSDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteHiTSDBInstanceResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteHiTSDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteHiTSDBInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteHiTSDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteHiTSDBInstanceResponse) SetHeaders(v map[string]*string) *DeleteHiTSDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteHiTSDBInstanceResponse) SetBody(v *DeleteHiTSDBInstanceResponseBody) *DeleteHiTSDBInstanceResponse {
	s.Body = v
	return s
}

type DescribeHiTSDBInstanceRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	AppKey               *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeHiTSDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHiTSDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeHiTSDBInstanceRequest) SetSecurityToken(v string) *DescribeHiTSDBInstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeHiTSDBInstanceRequest) SetOwnerId(v int64) *DescribeHiTSDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeHiTSDBInstanceRequest) SetResourceOwnerAccount(v string) *DescribeHiTSDBInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeHiTSDBInstanceRequest) SetResourceOwnerId(v int64) *DescribeHiTSDBInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeHiTSDBInstanceRequest) SetOwnerAccount(v string) *DescribeHiTSDBInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeHiTSDBInstanceRequest) SetAppKey(v string) *DescribeHiTSDBInstanceRequest {
	s.AppKey = &v
	return s
}

func (s *DescribeHiTSDBInstanceRequest) SetInstanceId(v string) *DescribeHiTSDBInstanceRequest {
	s.InstanceId = &v
	return s
}

type DescribeHiTSDBInstanceResponseBody struct {
	AutoRenew              *string                                             `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	GmtCreated             *string                                             `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	CpuNumber              *string                                             `json:"CpuNumber,omitempty" xml:"CpuNumber,omitempty"`
	MemSize                *string                                             `json:"MemSize,omitempty" xml:"MemSize,omitempty"`
	NetworkType            *string                                             `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	GmtExpire              *string                                             `json:"GmtExpire,omitempty" xml:"GmtExpire,omitempty"`
	InstanceAlias          *string                                             `json:"InstanceAlias,omitempty" xml:"InstanceAlias,omitempty"`
	InstanceStatus         *string                                             `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	ExpiredTime            *int64                                              `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	PaymentType            *string                                             `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	MaxTimelineLimit       *string                                             `json:"MaxTimelineLimit,omitempty" xml:"MaxTimelineLimit,omitempty"`
	PublicConnectionString *string                                             `json:"PublicConnectionString,omitempty" xml:"PublicConnectionString,omitempty"`
	EngineType             *string                                             `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	InstanceTps            *string                                             `json:"InstanceTps,omitempty" xml:"InstanceTps,omitempty"`
	Status                 *string                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	InstanceStorage        *string                                             `json:"InstanceStorage,omitempty" xml:"InstanceStorage,omitempty"`
	RequestId              *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ZoneId                 *string                                             `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	InstanceId             *string                                             `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	CreateTime             *int64                                              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DiskCategory           *string                                             `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	InstanceClass          *string                                             `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	VswitchId              *string                                             `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	Series                 *int32                                              `json:"Series,omitempty" xml:"Series,omitempty"`
	VpcId                  *string                                             `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ChargeType             *string                                             `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	SecurityIpList         []*DescribeHiTSDBInstanceResponseBodySecurityIpList `json:"SecurityIpList,omitempty" xml:"SecurityIpList,omitempty" type:"Repeated"`
	InstanceDescription    *string                                             `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	RegionId               *string                                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ConnectionString       *string                                             `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
}

func (s DescribeHiTSDBInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHiTSDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHiTSDBInstanceResponseBody) SetAutoRenew(v string) *DescribeHiTSDBInstanceResponseBody {
	s.AutoRenew = &v
	return s
}

func (s *DescribeHiTSDBInstanceResponseBody) SetGmtCreated(v string) *DescribeHiTSDBInstanceResponseBody {
	s.GmtCreated = &v
	return s
}

func (s *DescribeHiTSDBInstanceResponseBody) SetCpuNumber(v string) *DescribeHiTSDBInstanceResponseBody {
	s.CpuNumber = &v
	return s
}

func (s *DescribeHiTSDBInstanceResponseBody) SetMemSize(v string) *DescribeHiTSDBInstanceResponseBody {
	s.MemSize = &v
	return s
}

func (s *DescribeHiTSDBInstanceResponseBody) SetNetworkType(v string) *DescribeHiTSDBInstanceResponseBody {
	s.NetworkType = &v
	return s
}

func (s *DescribeHiTSDBInstanceResponseBody) SetGmtExpire(v string) *DescribeHiTSDBInstanceResponseBody {
	s.GmtExpire = &v
	return s
}

func (s *DescribeHiTSDBInstanceResponseBody) SetInstanceAlias(v string) *DescribeHiTSDBInstanceResponseBody {
	s.InstanceAlias = &v
	return s
}

func (s *DescribeHiTSDBInstanceResponseBody) SetInstanceStatus(v string) *DescribeHiTSDBInstanceResponseBody {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeHiTSDBInstanceResponseBody) SetExpiredTime(v int64) *DescribeHiTSDBInstanceResponseBody {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeHiTSDBInstanceResponseBody) SetPaymentType(v string) *DescribeHiTSDBInstanceResponseBody {
	s.PaymentType = &v
	return s
}

func (s *DescribeHiTSDBInstanceResponseBody) SetMaxTimelineLimit(v string) *DescribeHiTSDBInstanceResponseBody {
	s.MaxTimelineLimit = &v
	return s
}

func (s *DescribeHiTSDBInstanceResponseBody) SetPublicConnectionString(v string) *DescribeHiTSDBInstanceResponseBody {
	s.PublicConnectionString = &v
	return s
}

func (s *DescribeHiTSDBInstanceResponseBody) SetEngineType(v string) *DescribeHiTSDBInstanceResponseBody {
	s.EngineType = &v
	return s
}

func (s *DescribeHiTSDBInstanceResponseBody) SetInstanceTps(v string) *DescribeHiTSDBInstanceResponseBody {
	s.InstanceTps = &v
	return s
}

func (s *DescribeHiTSDBInstanceResponseBody) SetStatus(v string) *DescribeHiTSDBInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeHiTSDBInstanceResponseBody) SetInstanceStorage(v string) *DescribeHiTSDBInstanceResponseBody {
	s.InstanceStorage = &v
	return s
}

func (s *DescribeHiTSDBInstanceResponseBody) SetRequestId(v string) *DescribeHiTSDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHiTSDBInstanceResponseBody) SetZoneId(v string) *DescribeHiTSDBInstanceResponseBody {
	s.ZoneId = &v
	return s
}

func (s *DescribeHiTSDBInstanceResponseBody) SetInstanceId(v string) *DescribeHiTSDBInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeHiTSDBInstanceResponseBody) SetCreateTime(v int64) *DescribeHiTSDBInstanceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeHiTSDBInstanceResponseBody) SetDiskCategory(v string) *DescribeHiTSDBInstanceResponseBody {
	s.DiskCategory = &v
	return s
}

func (s *DescribeHiTSDBInstanceResponseBody) SetInstanceClass(v string) *DescribeHiTSDBInstanceResponseBody {
	s.InstanceClass = &v
	return s
}

func (s *DescribeHiTSDBInstanceResponseBody) SetVswitchId(v string) *DescribeHiTSDBInstanceResponseBody {
	s.VswitchId = &v
	return s
}

func (s *DescribeHiTSDBInstanceResponseBody) SetSeries(v int32) *DescribeHiTSDBInstanceResponseBody {
	s.Series = &v
	return s
}

func (s *DescribeHiTSDBInstanceResponseBody) SetVpcId(v string) *DescribeHiTSDBInstanceResponseBody {
	s.VpcId = &v
	return s
}

func (s *DescribeHiTSDBInstanceResponseBody) SetChargeType(v string) *DescribeHiTSDBInstanceResponseBody {
	s.ChargeType = &v
	return s
}

func (s *DescribeHiTSDBInstanceResponseBody) SetSecurityIpList(v []*DescribeHiTSDBInstanceResponseBodySecurityIpList) *DescribeHiTSDBInstanceResponseBody {
	s.SecurityIpList = v
	return s
}

func (s *DescribeHiTSDBInstanceResponseBody) SetInstanceDescription(v string) *DescribeHiTSDBInstanceResponseBody {
	s.InstanceDescription = &v
	return s
}

func (s *DescribeHiTSDBInstanceResponseBody) SetRegionId(v string) *DescribeHiTSDBInstanceResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeHiTSDBInstanceResponseBody) SetConnectionString(v string) *DescribeHiTSDBInstanceResponseBody {
	s.ConnectionString = &v
	return s
}

type DescribeHiTSDBInstanceResponseBodySecurityIpList struct {
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s DescribeHiTSDBInstanceResponseBodySecurityIpList) String() string {
	return tea.Prettify(s)
}

func (s DescribeHiTSDBInstanceResponseBodySecurityIpList) GoString() string {
	return s.String()
}

func (s *DescribeHiTSDBInstanceResponseBodySecurityIpList) SetIp(v string) *DescribeHiTSDBInstanceResponseBodySecurityIpList {
	s.Ip = &v
	return s
}

type DescribeHiTSDBInstanceResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeHiTSDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeHiTSDBInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHiTSDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeHiTSDBInstanceResponse) SetHeaders(v map[string]*string) *DescribeHiTSDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeHiTSDBInstanceResponse) SetBody(v *DescribeHiTSDBInstanceResponseBody) *DescribeHiTSDBInstanceResponse {
	s.Body = v
	return s
}

type DescribeHiTSDBInstanceListRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	AppKey               *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	QueryStr             *string `json:"QueryStr,omitempty" xml:"QueryStr,omitempty"`
	StatusList           *string `json:"StatusList,omitempty" xml:"StatusList,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	EngineType           *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
}

func (s DescribeHiTSDBInstanceListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHiTSDBInstanceListRequest) GoString() string {
	return s.String()
}

func (s *DescribeHiTSDBInstanceListRequest) SetSecurityToken(v string) *DescribeHiTSDBInstanceListRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeHiTSDBInstanceListRequest) SetOwnerId(v int64) *DescribeHiTSDBInstanceListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeHiTSDBInstanceListRequest) SetResourceOwnerAccount(v string) *DescribeHiTSDBInstanceListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeHiTSDBInstanceListRequest) SetResourceOwnerId(v int64) *DescribeHiTSDBInstanceListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeHiTSDBInstanceListRequest) SetOwnerAccount(v string) *DescribeHiTSDBInstanceListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeHiTSDBInstanceListRequest) SetAppKey(v string) *DescribeHiTSDBInstanceListRequest {
	s.AppKey = &v
	return s
}

func (s *DescribeHiTSDBInstanceListRequest) SetQueryStr(v string) *DescribeHiTSDBInstanceListRequest {
	s.QueryStr = &v
	return s
}

func (s *DescribeHiTSDBInstanceListRequest) SetStatusList(v string) *DescribeHiTSDBInstanceListRequest {
	s.StatusList = &v
	return s
}

func (s *DescribeHiTSDBInstanceListRequest) SetPageNumber(v int32) *DescribeHiTSDBInstanceListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeHiTSDBInstanceListRequest) SetPageSize(v int32) *DescribeHiTSDBInstanceListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeHiTSDBInstanceListRequest) SetEngineType(v string) *DescribeHiTSDBInstanceListRequest {
	s.EngineType = &v
	return s
}

type DescribeHiTSDBInstanceListResponseBody struct {
	RequestId    *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize     *int32                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber   *int32                                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Total        *int32                                                `json:"Total,omitempty" xml:"Total,omitempty"`
	InstanceList []*DescribeHiTSDBInstanceListResponseBodyInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
}

func (s DescribeHiTSDBInstanceListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHiTSDBInstanceListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHiTSDBInstanceListResponseBody) SetRequestId(v string) *DescribeHiTSDBInstanceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHiTSDBInstanceListResponseBody) SetPageSize(v int32) *DescribeHiTSDBInstanceListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeHiTSDBInstanceListResponseBody) SetPageNumber(v int32) *DescribeHiTSDBInstanceListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeHiTSDBInstanceListResponseBody) SetTotal(v int32) *DescribeHiTSDBInstanceListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeHiTSDBInstanceListResponseBody) SetInstanceList(v []*DescribeHiTSDBInstanceListResponseBodyInstanceList) *DescribeHiTSDBInstanceListResponseBody {
	s.InstanceList = v
	return s
}

type DescribeHiTSDBInstanceListResponseBodyInstanceList struct {
	VpcId                *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	MaxSeriesPerDatabase *string `json:"MaxSeriesPerDatabase,omitempty" xml:"MaxSeriesPerDatabase,omitempty"`
	PaymentType          *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	EngineType           *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	VswitchId            *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	InstanceClass        *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	CreateTime           *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	UserId               *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	ChargeType           *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	InstanceStorage      *string `json:"InstanceStorage,omitempty" xml:"InstanceStorage,omitempty"`
	NetworkType          *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LockMode             *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	InstanceDescription  *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	GmtCreated           *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	InstanceAlias        *string `json:"InstanceAlias,omitempty" xml:"InstanceAlias,omitempty"`
	InstanceTps          *string `json:"InstanceTps,omitempty" xml:"InstanceTps,omitempty"`
	ExpiredTime          *int64  `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	InstanceStatus       *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	GmtExpire            *string `json:"GmtExpire,omitempty" xml:"GmtExpire,omitempty"`
}

func (s DescribeHiTSDBInstanceListResponseBodyInstanceList) String() string {
	return tea.Prettify(s)
}

func (s DescribeHiTSDBInstanceListResponseBodyInstanceList) GoString() string {
	return s.String()
}

func (s *DescribeHiTSDBInstanceListResponseBodyInstanceList) SetVpcId(v string) *DescribeHiTSDBInstanceListResponseBodyInstanceList {
	s.VpcId = &v
	return s
}

func (s *DescribeHiTSDBInstanceListResponseBodyInstanceList) SetStatus(v string) *DescribeHiTSDBInstanceListResponseBodyInstanceList {
	s.Status = &v
	return s
}

func (s *DescribeHiTSDBInstanceListResponseBodyInstanceList) SetMaxSeriesPerDatabase(v string) *DescribeHiTSDBInstanceListResponseBodyInstanceList {
	s.MaxSeriesPerDatabase = &v
	return s
}

func (s *DescribeHiTSDBInstanceListResponseBodyInstanceList) SetPaymentType(v string) *DescribeHiTSDBInstanceListResponseBodyInstanceList {
	s.PaymentType = &v
	return s
}

func (s *DescribeHiTSDBInstanceListResponseBodyInstanceList) SetEngineType(v string) *DescribeHiTSDBInstanceListResponseBodyInstanceList {
	s.EngineType = &v
	return s
}

func (s *DescribeHiTSDBInstanceListResponseBodyInstanceList) SetVswitchId(v string) *DescribeHiTSDBInstanceListResponseBodyInstanceList {
	s.VswitchId = &v
	return s
}

func (s *DescribeHiTSDBInstanceListResponseBodyInstanceList) SetInstanceClass(v string) *DescribeHiTSDBInstanceListResponseBodyInstanceList {
	s.InstanceClass = &v
	return s
}

func (s *DescribeHiTSDBInstanceListResponseBodyInstanceList) SetCreateTime(v int64) *DescribeHiTSDBInstanceListResponseBodyInstanceList {
	s.CreateTime = &v
	return s
}

func (s *DescribeHiTSDBInstanceListResponseBodyInstanceList) SetUserId(v string) *DescribeHiTSDBInstanceListResponseBodyInstanceList {
	s.UserId = &v
	return s
}

func (s *DescribeHiTSDBInstanceListResponseBodyInstanceList) SetChargeType(v string) *DescribeHiTSDBInstanceListResponseBodyInstanceList {
	s.ChargeType = &v
	return s
}

func (s *DescribeHiTSDBInstanceListResponseBodyInstanceList) SetInstanceStorage(v string) *DescribeHiTSDBInstanceListResponseBodyInstanceList {
	s.InstanceStorage = &v
	return s
}

func (s *DescribeHiTSDBInstanceListResponseBodyInstanceList) SetNetworkType(v string) *DescribeHiTSDBInstanceListResponseBodyInstanceList {
	s.NetworkType = &v
	return s
}

func (s *DescribeHiTSDBInstanceListResponseBodyInstanceList) SetInstanceId(v string) *DescribeHiTSDBInstanceListResponseBodyInstanceList {
	s.InstanceId = &v
	return s
}

func (s *DescribeHiTSDBInstanceListResponseBodyInstanceList) SetLockMode(v string) *DescribeHiTSDBInstanceListResponseBodyInstanceList {
	s.LockMode = &v
	return s
}

func (s *DescribeHiTSDBInstanceListResponseBodyInstanceList) SetInstanceDescription(v string) *DescribeHiTSDBInstanceListResponseBodyInstanceList {
	s.InstanceDescription = &v
	return s
}

func (s *DescribeHiTSDBInstanceListResponseBodyInstanceList) SetRegionId(v string) *DescribeHiTSDBInstanceListResponseBodyInstanceList {
	s.RegionId = &v
	return s
}

func (s *DescribeHiTSDBInstanceListResponseBodyInstanceList) SetGmtCreated(v string) *DescribeHiTSDBInstanceListResponseBodyInstanceList {
	s.GmtCreated = &v
	return s
}

func (s *DescribeHiTSDBInstanceListResponseBodyInstanceList) SetInstanceAlias(v string) *DescribeHiTSDBInstanceListResponseBodyInstanceList {
	s.InstanceAlias = &v
	return s
}

func (s *DescribeHiTSDBInstanceListResponseBodyInstanceList) SetInstanceTps(v string) *DescribeHiTSDBInstanceListResponseBodyInstanceList {
	s.InstanceTps = &v
	return s
}

func (s *DescribeHiTSDBInstanceListResponseBodyInstanceList) SetExpiredTime(v int64) *DescribeHiTSDBInstanceListResponseBodyInstanceList {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeHiTSDBInstanceListResponseBodyInstanceList) SetZoneId(v string) *DescribeHiTSDBInstanceListResponseBodyInstanceList {
	s.ZoneId = &v
	return s
}

func (s *DescribeHiTSDBInstanceListResponseBodyInstanceList) SetInstanceStatus(v string) *DescribeHiTSDBInstanceListResponseBodyInstanceList {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeHiTSDBInstanceListResponseBodyInstanceList) SetGmtExpire(v string) *DescribeHiTSDBInstanceListResponseBodyInstanceList {
	s.GmtExpire = &v
	return s
}

type DescribeHiTSDBInstanceListResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeHiTSDBInstanceListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeHiTSDBInstanceListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHiTSDBInstanceListResponse) GoString() string {
	return s.String()
}

func (s *DescribeHiTSDBInstanceListResponse) SetHeaders(v map[string]*string) *DescribeHiTSDBInstanceListResponse {
	s.Headers = v
	return s
}

func (s *DescribeHiTSDBInstanceListResponse) SetBody(v *DescribeHiTSDBInstanceListResponseBody) *DescribeHiTSDBInstanceListResponse {
	s.Body = v
	return s
}

type DescribeHiTSDBInstanceSecurityIpListRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	GroupName            *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s DescribeHiTSDBInstanceSecurityIpListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHiTSDBInstanceSecurityIpListRequest) GoString() string {
	return s.String()
}

func (s *DescribeHiTSDBInstanceSecurityIpListRequest) SetSecurityToken(v string) *DescribeHiTSDBInstanceSecurityIpListRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeHiTSDBInstanceSecurityIpListRequest) SetOwnerId(v int64) *DescribeHiTSDBInstanceSecurityIpListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeHiTSDBInstanceSecurityIpListRequest) SetResourceOwnerAccount(v string) *DescribeHiTSDBInstanceSecurityIpListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeHiTSDBInstanceSecurityIpListRequest) SetResourceOwnerId(v int64) *DescribeHiTSDBInstanceSecurityIpListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeHiTSDBInstanceSecurityIpListRequest) SetOwnerAccount(v string) *DescribeHiTSDBInstanceSecurityIpListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeHiTSDBInstanceSecurityIpListRequest) SetInstanceId(v string) *DescribeHiTSDBInstanceSecurityIpListRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeHiTSDBInstanceSecurityIpListRequest) SetGroupName(v string) *DescribeHiTSDBInstanceSecurityIpListRequest {
	s.GroupName = &v
	return s
}

type DescribeHiTSDBInstanceSecurityIpListResponseBody struct {
	RequestId      *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecurityIpList []*DescribeHiTSDBInstanceSecurityIpListResponseBodySecurityIpList `json:"SecurityIpList,omitempty" xml:"SecurityIpList,omitempty" type:"Repeated"`
}

func (s DescribeHiTSDBInstanceSecurityIpListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHiTSDBInstanceSecurityIpListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHiTSDBInstanceSecurityIpListResponseBody) SetRequestId(v string) *DescribeHiTSDBInstanceSecurityIpListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHiTSDBInstanceSecurityIpListResponseBody) SetSecurityIpList(v []*DescribeHiTSDBInstanceSecurityIpListResponseBodySecurityIpList) *DescribeHiTSDBInstanceSecurityIpListResponseBody {
	s.SecurityIpList = v
	return s
}

type DescribeHiTSDBInstanceSecurityIpListResponseBodySecurityIpList struct {
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s DescribeHiTSDBInstanceSecurityIpListResponseBodySecurityIpList) String() string {
	return tea.Prettify(s)
}

func (s DescribeHiTSDBInstanceSecurityIpListResponseBodySecurityIpList) GoString() string {
	return s.String()
}

func (s *DescribeHiTSDBInstanceSecurityIpListResponseBodySecurityIpList) SetIp(v string) *DescribeHiTSDBInstanceSecurityIpListResponseBodySecurityIpList {
	s.Ip = &v
	return s
}

type DescribeHiTSDBInstanceSecurityIpListResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeHiTSDBInstanceSecurityIpListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeHiTSDBInstanceSecurityIpListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHiTSDBInstanceSecurityIpListResponse) GoString() string {
	return s.String()
}

func (s *DescribeHiTSDBInstanceSecurityIpListResponse) SetHeaders(v map[string]*string) *DescribeHiTSDBInstanceSecurityIpListResponse {
	s.Headers = v
	return s
}

func (s *DescribeHiTSDBInstanceSecurityIpListResponse) SetBody(v *DescribeHiTSDBInstanceSecurityIpListResponseBody) *DescribeHiTSDBInstanceSecurityIpListResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	AcceptLanguage       *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetSecurityToken(v string) *DescribeRegionsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeRegionsRequest) SetOwnerId(v int64) *DescribeRegionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRegionsRequest) SetResourceOwnerAccount(v string) *DescribeRegionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRegionsRequest) SetResourceOwnerId(v int64) *DescribeRegionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRegionsRequest) SetOwnerAccount(v string) *DescribeRegionsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRegionsRequest) SetAcceptLanguage(v string) *DescribeRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

type DescribeRegionsResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Regions   *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	Region []*DescribeRegionsResponseBodyRegionsRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetRegion(v []*DescribeRegionsResponseBodyRegionsRegion) *DescribeRegionsResponseBodyRegions {
	s.Region = v
	return s
}

type DescribeRegionsResponseBodyRegionsRegion struct {
	LocalName      *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetLocalName(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionId(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeZonesRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	Language             *string `json:"Language,omitempty" xml:"Language,omitempty"`
}

func (s DescribeZonesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeZonesRequest) SetSecurityToken(v string) *DescribeZonesRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeZonesRequest) SetOwnerId(v int64) *DescribeZonesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeZonesRequest) SetResourceOwnerAccount(v string) *DescribeZonesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeZonesRequest) SetResourceOwnerId(v int64) *DescribeZonesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeZonesRequest) SetOwnerAccount(v string) *DescribeZonesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeZonesRequest) SetLanguage(v string) *DescribeZonesRequest {
	s.Language = &v
	return s
}

type DescribeZonesResponseBody struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ZoneList  *DescribeZonesResponseBodyZoneList `json:"ZoneList,omitempty" xml:"ZoneList,omitempty" type:"Struct"`
}

func (s DescribeZonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBody) SetRequestId(v string) *DescribeZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeZonesResponseBody) SetZoneList(v *DescribeZonesResponseBodyZoneList) *DescribeZonesResponseBody {
	s.ZoneList = v
	return s
}

type DescribeZonesResponseBodyZoneList struct {
	ZoneModel []*DescribeZonesResponseBodyZoneListZoneModel `json:"ZoneModel,omitempty" xml:"ZoneModel,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBodyZoneList) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBodyZoneList) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZoneList) SetZoneModel(v []*DescribeZonesResponseBodyZoneListZoneModel) *DescribeZonesResponseBodyZoneList {
	s.ZoneModel = v
	return s
}

type DescribeZonesResponseBodyZoneListZoneModel struct {
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	ZoneId    *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeZonesResponseBodyZoneListZoneModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBodyZoneListZoneModel) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZoneListZoneModel) SetLocalName(v string) *DescribeZonesResponseBodyZoneListZoneModel {
	s.LocalName = &v
	return s
}

func (s *DescribeZonesResponseBodyZoneListZoneModel) SetZoneId(v string) *DescribeZonesResponseBodyZoneListZoneModel {
	s.ZoneId = &v
	return s
}

type DescribeZonesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeZonesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeZonesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponse) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponse) SetHeaders(v map[string]*string) *DescribeZonesResponse {
	s.Headers = v
	return s
}

func (s *DescribeZonesResponse) SetBody(v *DescribeZonesResponseBody) *DescribeZonesResponse {
	s.Body = v
	return s
}

type ModifyHiTSDBInstanceClassRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	AppKey               *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceClass        *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	InstanceStorage      *string `json:"InstanceStorage,omitempty" xml:"InstanceStorage,omitempty"`
}

func (s ModifyHiTSDBInstanceClassRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyHiTSDBInstanceClassRequest) GoString() string {
	return s.String()
}

func (s *ModifyHiTSDBInstanceClassRequest) SetSecurityToken(v string) *ModifyHiTSDBInstanceClassRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyHiTSDBInstanceClassRequest) SetOwnerId(v int64) *ModifyHiTSDBInstanceClassRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyHiTSDBInstanceClassRequest) SetResourceOwnerAccount(v string) *ModifyHiTSDBInstanceClassRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyHiTSDBInstanceClassRequest) SetResourceOwnerId(v int64) *ModifyHiTSDBInstanceClassRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyHiTSDBInstanceClassRequest) SetOwnerAccount(v string) *ModifyHiTSDBInstanceClassRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyHiTSDBInstanceClassRequest) SetAppKey(v string) *ModifyHiTSDBInstanceClassRequest {
	s.AppKey = &v
	return s
}

func (s *ModifyHiTSDBInstanceClassRequest) SetInstanceId(v string) *ModifyHiTSDBInstanceClassRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyHiTSDBInstanceClassRequest) SetInstanceClass(v string) *ModifyHiTSDBInstanceClassRequest {
	s.InstanceClass = &v
	return s
}

func (s *ModifyHiTSDBInstanceClassRequest) SetInstanceStorage(v string) *ModifyHiTSDBInstanceClassRequest {
	s.InstanceStorage = &v
	return s
}

type ModifyHiTSDBInstanceClassResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyHiTSDBInstanceClassResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyHiTSDBInstanceClassResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHiTSDBInstanceClassResponseBody) SetRequestId(v string) *ModifyHiTSDBInstanceClassResponseBody {
	s.RequestId = &v
	return s
}

type ModifyHiTSDBInstanceClassResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyHiTSDBInstanceClassResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyHiTSDBInstanceClassResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyHiTSDBInstanceClassResponse) GoString() string {
	return s.String()
}

func (s *ModifyHiTSDBInstanceClassResponse) SetHeaders(v map[string]*string) *ModifyHiTSDBInstanceClassResponse {
	s.Headers = v
	return s
}

func (s *ModifyHiTSDBInstanceClassResponse) SetBody(v *ModifyHiTSDBInstanceClassResponseBody) *ModifyHiTSDBInstanceClassResponse {
	s.Body = v
	return s
}

type ModifyHiTSDBInstanceSecurityIpListRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SecurityIpList       *string `json:"SecurityIpList,omitempty" xml:"SecurityIpList,omitempty"`
	GroupName            *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s ModifyHiTSDBInstanceSecurityIpListRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyHiTSDBInstanceSecurityIpListRequest) GoString() string {
	return s.String()
}

func (s *ModifyHiTSDBInstanceSecurityIpListRequest) SetSecurityToken(v string) *ModifyHiTSDBInstanceSecurityIpListRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyHiTSDBInstanceSecurityIpListRequest) SetOwnerId(v int64) *ModifyHiTSDBInstanceSecurityIpListRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyHiTSDBInstanceSecurityIpListRequest) SetResourceOwnerAccount(v string) *ModifyHiTSDBInstanceSecurityIpListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyHiTSDBInstanceSecurityIpListRequest) SetResourceOwnerId(v int64) *ModifyHiTSDBInstanceSecurityIpListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyHiTSDBInstanceSecurityIpListRequest) SetOwnerAccount(v string) *ModifyHiTSDBInstanceSecurityIpListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyHiTSDBInstanceSecurityIpListRequest) SetInstanceId(v string) *ModifyHiTSDBInstanceSecurityIpListRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyHiTSDBInstanceSecurityIpListRequest) SetSecurityIpList(v string) *ModifyHiTSDBInstanceSecurityIpListRequest {
	s.SecurityIpList = &v
	return s
}

func (s *ModifyHiTSDBInstanceSecurityIpListRequest) SetGroupName(v string) *ModifyHiTSDBInstanceSecurityIpListRequest {
	s.GroupName = &v
	return s
}

type ModifyHiTSDBInstanceSecurityIpListResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyHiTSDBInstanceSecurityIpListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyHiTSDBInstanceSecurityIpListResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHiTSDBInstanceSecurityIpListResponseBody) SetRequestId(v string) *ModifyHiTSDBInstanceSecurityIpListResponseBody {
	s.RequestId = &v
	return s
}

type ModifyHiTSDBInstanceSecurityIpListResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyHiTSDBInstanceSecurityIpListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyHiTSDBInstanceSecurityIpListResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyHiTSDBInstanceSecurityIpListResponse) GoString() string {
	return s.String()
}

func (s *ModifyHiTSDBInstanceSecurityIpListResponse) SetHeaders(v map[string]*string) *ModifyHiTSDBInstanceSecurityIpListResponse {
	s.Headers = v
	return s
}

func (s *ModifyHiTSDBInstanceSecurityIpListResponse) SetBody(v *ModifyHiTSDBInstanceSecurityIpListResponseBody) *ModifyHiTSDBInstanceSecurityIpListResponse {
	s.Body = v
	return s
}

type RenameHiTSDBInstanceAliasRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	AppKey               *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceAlias        *string `json:"InstanceAlias,omitempty" xml:"InstanceAlias,omitempty"`
}

func (s RenameHiTSDBInstanceAliasRequest) String() string {
	return tea.Prettify(s)
}

func (s RenameHiTSDBInstanceAliasRequest) GoString() string {
	return s.String()
}

func (s *RenameHiTSDBInstanceAliasRequest) SetSecurityToken(v string) *RenameHiTSDBInstanceAliasRequest {
	s.SecurityToken = &v
	return s
}

func (s *RenameHiTSDBInstanceAliasRequest) SetOwnerId(v int64) *RenameHiTSDBInstanceAliasRequest {
	s.OwnerId = &v
	return s
}

func (s *RenameHiTSDBInstanceAliasRequest) SetResourceOwnerAccount(v string) *RenameHiTSDBInstanceAliasRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RenameHiTSDBInstanceAliasRequest) SetResourceOwnerId(v int64) *RenameHiTSDBInstanceAliasRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RenameHiTSDBInstanceAliasRequest) SetOwnerAccount(v string) *RenameHiTSDBInstanceAliasRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RenameHiTSDBInstanceAliasRequest) SetAppKey(v string) *RenameHiTSDBInstanceAliasRequest {
	s.AppKey = &v
	return s
}

func (s *RenameHiTSDBInstanceAliasRequest) SetInstanceId(v string) *RenameHiTSDBInstanceAliasRequest {
	s.InstanceId = &v
	return s
}

func (s *RenameHiTSDBInstanceAliasRequest) SetInstanceAlias(v string) *RenameHiTSDBInstanceAliasRequest {
	s.InstanceAlias = &v
	return s
}

type RenameHiTSDBInstanceAliasResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenameHiTSDBInstanceAliasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenameHiTSDBInstanceAliasResponseBody) GoString() string {
	return s.String()
}

func (s *RenameHiTSDBInstanceAliasResponseBody) SetRequestId(v string) *RenameHiTSDBInstanceAliasResponseBody {
	s.RequestId = &v
	return s
}

type RenameHiTSDBInstanceAliasResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RenameHiTSDBInstanceAliasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RenameHiTSDBInstanceAliasResponse) String() string {
	return tea.Prettify(s)
}

func (s RenameHiTSDBInstanceAliasResponse) GoString() string {
	return s.String()
}

func (s *RenameHiTSDBInstanceAliasResponse) SetHeaders(v map[string]*string) *RenameHiTSDBInstanceAliasResponse {
	s.Headers = v
	return s
}

func (s *RenameHiTSDBInstanceAliasResponse) SetBody(v *RenameHiTSDBInstanceAliasResponseBody) *RenameHiTSDBInstanceAliasResponse {
	s.Body = v
	return s
}

type RenewTSDBInstanceRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PricingCycle         *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	Duration             *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
}

func (s RenewTSDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RenewTSDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *RenewTSDBInstanceRequest) SetSecurityToken(v string) *RenewTSDBInstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *RenewTSDBInstanceRequest) SetOwnerId(v int64) *RenewTSDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *RenewTSDBInstanceRequest) SetResourceOwnerAccount(v string) *RenewTSDBInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RenewTSDBInstanceRequest) SetResourceOwnerId(v int64) *RenewTSDBInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RenewTSDBInstanceRequest) SetOwnerAccount(v string) *RenewTSDBInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RenewTSDBInstanceRequest) SetRegionId(v string) *RenewTSDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *RenewTSDBInstanceRequest) SetInstanceId(v string) *RenewTSDBInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RenewTSDBInstanceRequest) SetPricingCycle(v string) *RenewTSDBInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *RenewTSDBInstanceRequest) SetDuration(v int32) *RenewTSDBInstanceRequest {
	s.Duration = &v
	return s
}

type RenewTSDBInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId   *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s RenewTSDBInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenewTSDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RenewTSDBInstanceResponseBody) SetRequestId(v string) *RenewTSDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewTSDBInstanceResponseBody) SetOrderId(v int64) *RenewTSDBInstanceResponseBody {
	s.OrderId = &v
	return s
}

type RenewTSDBInstanceResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RenewTSDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RenewTSDBInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RenewTSDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *RenewTSDBInstanceResponse) SetHeaders(v map[string]*string) *RenewTSDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *RenewTSDBInstanceResponse) SetBody(v *RenewTSDBInstanceResponseBody) *RenewTSDBInstanceResponse {
	s.Body = v
	return s
}

type RestartHiTSDBInstanceRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s RestartHiTSDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartHiTSDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestartHiTSDBInstanceRequest) SetSecurityToken(v string) *RestartHiTSDBInstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *RestartHiTSDBInstanceRequest) SetOwnerId(v int64) *RestartHiTSDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *RestartHiTSDBInstanceRequest) SetResourceOwnerAccount(v string) *RestartHiTSDBInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RestartHiTSDBInstanceRequest) SetResourceOwnerId(v int64) *RestartHiTSDBInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RestartHiTSDBInstanceRequest) SetOwnerAccount(v string) *RestartHiTSDBInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RestartHiTSDBInstanceRequest) SetInstanceId(v string) *RestartHiTSDBInstanceRequest {
	s.InstanceId = &v
	return s
}

type RestartHiTSDBInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartHiTSDBInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestartHiTSDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RestartHiTSDBInstanceResponseBody) SetRequestId(v string) *RestartHiTSDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

type RestartHiTSDBInstanceResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RestartHiTSDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RestartHiTSDBInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartHiTSDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *RestartHiTSDBInstanceResponse) SetHeaders(v map[string]*string) *RestartHiTSDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *RestartHiTSDBInstanceResponse) SetBody(v *RestartHiTSDBInstanceResponseBody) *RestartHiTSDBInstanceResponse {
	s.Body = v
	return s
}

type SwitchHiTSDBInstancePublicNetRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SwitchAction         *int64  `json:"SwitchAction,omitempty" xml:"SwitchAction,omitempty"`
}

func (s SwitchHiTSDBInstancePublicNetRequest) String() string {
	return tea.Prettify(s)
}

func (s SwitchHiTSDBInstancePublicNetRequest) GoString() string {
	return s.String()
}

func (s *SwitchHiTSDBInstancePublicNetRequest) SetSecurityToken(v string) *SwitchHiTSDBInstancePublicNetRequest {
	s.SecurityToken = &v
	return s
}

func (s *SwitchHiTSDBInstancePublicNetRequest) SetOwnerId(v int64) *SwitchHiTSDBInstancePublicNetRequest {
	s.OwnerId = &v
	return s
}

func (s *SwitchHiTSDBInstancePublicNetRequest) SetResourceOwnerAccount(v string) *SwitchHiTSDBInstancePublicNetRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SwitchHiTSDBInstancePublicNetRequest) SetResourceOwnerId(v int64) *SwitchHiTSDBInstancePublicNetRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SwitchHiTSDBInstancePublicNetRequest) SetOwnerAccount(v string) *SwitchHiTSDBInstancePublicNetRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SwitchHiTSDBInstancePublicNetRequest) SetInstanceId(v string) *SwitchHiTSDBInstancePublicNetRequest {
	s.InstanceId = &v
	return s
}

func (s *SwitchHiTSDBInstancePublicNetRequest) SetSwitchAction(v int64) *SwitchHiTSDBInstancePublicNetRequest {
	s.SwitchAction = &v
	return s
}

type SwitchHiTSDBInstancePublicNetResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SwitchHiTSDBInstancePublicNetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SwitchHiTSDBInstancePublicNetResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchHiTSDBInstancePublicNetResponseBody) SetRequestId(v string) *SwitchHiTSDBInstancePublicNetResponseBody {
	s.RequestId = &v
	return s
}

type SwitchHiTSDBInstancePublicNetResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SwitchHiTSDBInstancePublicNetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SwitchHiTSDBInstancePublicNetResponse) String() string {
	return tea.Prettify(s)
}

func (s SwitchHiTSDBInstancePublicNetResponse) GoString() string {
	return s.String()
}

func (s *SwitchHiTSDBInstancePublicNetResponse) SetHeaders(v map[string]*string) *SwitchHiTSDBInstancePublicNetResponse {
	s.Headers = v
	return s
}

func (s *SwitchHiTSDBInstancePublicNetResponse) SetBody(v *SwitchHiTSDBInstancePublicNetResponseBody) *SwitchHiTSDBInstancePublicNetResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-qingdao":                  tea.String("hitsdb.aliyuncs.com"),
		"cn-beijing":                  tea.String("hitsdb.aliyuncs.com"),
		"cn-hangzhou":                 tea.String("hitsdb.aliyuncs.com"),
		"cn-shanghai":                 tea.String("hitsdb.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("hitsdb.aliyuncs.com"),
		"cn-hongkong":                 tea.String("hitsdb.aliyuncs.com"),
		"ap-southeast-1":              tea.String("hitsdb.aliyuncs.com"),
		"us-west-1":                   tea.String("hitsdb.aliyuncs.com"),
		"us-east-1":                   tea.String("hitsdb.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("hitsdb.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("hitsdb.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("hitsdb.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("hitsdb.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("hitsdb.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("hitsdb.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("hitsdb.aliyuncs.com"),
		"cn-chengdu":                  tea.String("hitsdb.aliyuncs.com"),
		"cn-edge-1":                   tea.String("hitsdb.aliyuncs.com"),
		"cn-fujian":                   tea.String("hitsdb.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("hitsdb.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("hitsdb.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("hitsdb.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("hitsdb.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("hitsdb.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("hitsdb.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("hitsdb.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("hitsdb.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("hitsdb.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("hitsdb.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("hitsdb.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("hitsdb.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("hitsdb.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("hitsdb.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("hitsdb.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("hitsdb.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("hitsdb.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("hitsdb.aliyuncs.com"),
		"cn-wuhan":                    tea.String("hitsdb.aliyuncs.com"),
		"cn-wulanchabu":               tea.String("hitsdb.aliyuncs.com"),
		"cn-yushanfang":               tea.String("hitsdb.aliyuncs.com"),
		"cn-zhangbei":                 tea.String("hitsdb.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("hitsdb.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("hitsdb.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("hitsdb.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("hitsdb.aliyuncs.com"),
		"me-east-1":                   tea.String("hitsdb.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("hitsdb.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("hitsdb"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateHiTSDBInstanceWithOptions(request *CreateHiTSDBInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateHiTSDBInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateHiTSDBInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateHiTSDBInstance"), tea.String("2017-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateHiTSDBInstance(request *CreateHiTSDBInstanceRequest) (_result *CreateHiTSDBInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateHiTSDBInstanceResponse{}
	_body, _err := client.CreateHiTSDBInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteHiTSDBInstanceWithOptions(request *DeleteHiTSDBInstanceRequest, runtime *util.RuntimeOptions) (_result *DeleteHiTSDBInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteHiTSDBInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteHiTSDBInstance"), tea.String("2017-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteHiTSDBInstance(request *DeleteHiTSDBInstanceRequest) (_result *DeleteHiTSDBInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteHiTSDBInstanceResponse{}
	_body, _err := client.DeleteHiTSDBInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHiTSDBInstanceWithOptions(request *DescribeHiTSDBInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeHiTSDBInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeHiTSDBInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeHiTSDBInstance"), tea.String("2017-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHiTSDBInstance(request *DescribeHiTSDBInstanceRequest) (_result *DescribeHiTSDBInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHiTSDBInstanceResponse{}
	_body, _err := client.DescribeHiTSDBInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHiTSDBInstanceListWithOptions(request *DescribeHiTSDBInstanceListRequest, runtime *util.RuntimeOptions) (_result *DescribeHiTSDBInstanceListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeHiTSDBInstanceListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeHiTSDBInstanceList"), tea.String("2017-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHiTSDBInstanceList(request *DescribeHiTSDBInstanceListRequest) (_result *DescribeHiTSDBInstanceListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHiTSDBInstanceListResponse{}
	_body, _err := client.DescribeHiTSDBInstanceListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHiTSDBInstanceSecurityIpListWithOptions(request *DescribeHiTSDBInstanceSecurityIpListRequest, runtime *util.RuntimeOptions) (_result *DescribeHiTSDBInstanceSecurityIpListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeHiTSDBInstanceSecurityIpListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeHiTSDBInstanceSecurityIpList"), tea.String("2017-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHiTSDBInstanceSecurityIpList(request *DescribeHiTSDBInstanceSecurityIpListRequest) (_result *DescribeHiTSDBInstanceSecurityIpListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHiTSDBInstanceSecurityIpListResponse{}
	_body, _err := client.DescribeHiTSDBInstanceSecurityIpListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRegions"), tea.String("2017-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeZonesWithOptions(request *DescribeZonesRequest, runtime *util.RuntimeOptions) (_result *DescribeZonesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeZonesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeZones"), tea.String("2017-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeZones(request *DescribeZonesRequest) (_result *DescribeZonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeZonesResponse{}
	_body, _err := client.DescribeZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyHiTSDBInstanceClassWithOptions(request *ModifyHiTSDBInstanceClassRequest, runtime *util.RuntimeOptions) (_result *ModifyHiTSDBInstanceClassResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyHiTSDBInstanceClassResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyHiTSDBInstanceClass"), tea.String("2017-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyHiTSDBInstanceClass(request *ModifyHiTSDBInstanceClassRequest) (_result *ModifyHiTSDBInstanceClassResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyHiTSDBInstanceClassResponse{}
	_body, _err := client.ModifyHiTSDBInstanceClassWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyHiTSDBInstanceSecurityIpListWithOptions(request *ModifyHiTSDBInstanceSecurityIpListRequest, runtime *util.RuntimeOptions) (_result *ModifyHiTSDBInstanceSecurityIpListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyHiTSDBInstanceSecurityIpListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyHiTSDBInstanceSecurityIpList"), tea.String("2017-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyHiTSDBInstanceSecurityIpList(request *ModifyHiTSDBInstanceSecurityIpListRequest) (_result *ModifyHiTSDBInstanceSecurityIpListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyHiTSDBInstanceSecurityIpListResponse{}
	_body, _err := client.ModifyHiTSDBInstanceSecurityIpListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenameHiTSDBInstanceAliasWithOptions(request *RenameHiTSDBInstanceAliasRequest, runtime *util.RuntimeOptions) (_result *RenameHiTSDBInstanceAliasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RenameHiTSDBInstanceAliasResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RenameHiTSDBInstanceAlias"), tea.String("2017-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenameHiTSDBInstanceAlias(request *RenameHiTSDBInstanceAliasRequest) (_result *RenameHiTSDBInstanceAliasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenameHiTSDBInstanceAliasResponse{}
	_body, _err := client.RenameHiTSDBInstanceAliasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenewTSDBInstanceWithOptions(request *RenewTSDBInstanceRequest, runtime *util.RuntimeOptions) (_result *RenewTSDBInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RenewTSDBInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RenewTSDBInstance"), tea.String("2017-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenewTSDBInstance(request *RenewTSDBInstanceRequest) (_result *RenewTSDBInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenewTSDBInstanceResponse{}
	_body, _err := client.RenewTSDBInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RestartHiTSDBInstanceWithOptions(request *RestartHiTSDBInstanceRequest, runtime *util.RuntimeOptions) (_result *RestartHiTSDBInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RestartHiTSDBInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RestartHiTSDBInstance"), tea.String("2017-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RestartHiTSDBInstance(request *RestartHiTSDBInstanceRequest) (_result *RestartHiTSDBInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestartHiTSDBInstanceResponse{}
	_body, _err := client.RestartHiTSDBInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SwitchHiTSDBInstancePublicNetWithOptions(request *SwitchHiTSDBInstancePublicNetRequest, runtime *util.RuntimeOptions) (_result *SwitchHiTSDBInstancePublicNetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SwitchHiTSDBInstancePublicNetResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SwitchHiTSDBInstancePublicNet"), tea.String("2017-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SwitchHiTSDBInstancePublicNet(request *SwitchHiTSDBInstancePublicNetRequest) (_result *SwitchHiTSDBInstancePublicNetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SwitchHiTSDBInstancePublicNetResponse{}
	_body, _err := client.SwitchHiTSDBInstancePublicNetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
