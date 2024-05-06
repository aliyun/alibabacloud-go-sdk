// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddUserHdfsInfoRequest struct {
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ExtInfo              *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s AddUserHdfsInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s AddUserHdfsInfoRequest) GoString() string {
	return s.String()
}

func (s *AddUserHdfsInfoRequest) SetClusterId(v string) *AddUserHdfsInfoRequest {
	s.ClusterId = &v
	return s
}

func (s *AddUserHdfsInfoRequest) SetExtInfo(v string) *AddUserHdfsInfoRequest {
	s.ExtInfo = &v
	return s
}

func (s *AddUserHdfsInfoRequest) SetOwnerId(v int64) *AddUserHdfsInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *AddUserHdfsInfoRequest) SetRegionId(v string) *AddUserHdfsInfoRequest {
	s.RegionId = &v
	return s
}

func (s *AddUserHdfsInfoRequest) SetResourceOwnerAccount(v string) *AddUserHdfsInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddUserHdfsInfoRequest) SetResourceOwnerId(v int64) *AddUserHdfsInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddUserHdfsInfoRequest) SetZoneId(v string) *AddUserHdfsInfoRequest {
	s.ZoneId = &v
	return s
}

type AddUserHdfsInfoResponseBody struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddUserHdfsInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddUserHdfsInfoResponseBody) GoString() string {
	return s.String()
}

func (s *AddUserHdfsInfoResponseBody) SetClusterId(v string) *AddUserHdfsInfoResponseBody {
	s.ClusterId = &v
	return s
}

func (s *AddUserHdfsInfoResponseBody) SetRequestId(v string) *AddUserHdfsInfoResponseBody {
	s.RequestId = &v
	return s
}

type AddUserHdfsInfoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddUserHdfsInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddUserHdfsInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s AddUserHdfsInfoResponse) GoString() string {
	return s.String()
}

func (s *AddUserHdfsInfoResponse) SetHeaders(v map[string]*string) *AddUserHdfsInfoResponse {
	s.Headers = v
	return s
}

func (s *AddUserHdfsInfoResponse) SetStatusCode(v int32) *AddUserHdfsInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *AddUserHdfsInfoResponse) SetBody(v *AddUserHdfsInfoResponseBody) *AddUserHdfsInfoResponse {
	s.Body = v
	return s
}

type AllocatePublicNetworkAddressRequest struct {
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AllocatePublicNetworkAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s AllocatePublicNetworkAddressRequest) GoString() string {
	return s.String()
}

func (s *AllocatePublicNetworkAddressRequest) SetClusterId(v string) *AllocatePublicNetworkAddressRequest {
	s.ClusterId = &v
	return s
}

func (s *AllocatePublicNetworkAddressRequest) SetOwnerId(v int64) *AllocatePublicNetworkAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *AllocatePublicNetworkAddressRequest) SetResourceOwnerAccount(v string) *AllocatePublicNetworkAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AllocatePublicNetworkAddressRequest) SetResourceOwnerId(v int64) *AllocatePublicNetworkAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

type AllocatePublicNetworkAddressResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AllocatePublicNetworkAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AllocatePublicNetworkAddressResponseBody) GoString() string {
	return s.String()
}

func (s *AllocatePublicNetworkAddressResponseBody) SetRequestId(v string) *AllocatePublicNetworkAddressResponseBody {
	s.RequestId = &v
	return s
}

type AllocatePublicNetworkAddressResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AllocatePublicNetworkAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AllocatePublicNetworkAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s AllocatePublicNetworkAddressResponse) GoString() string {
	return s.String()
}

func (s *AllocatePublicNetworkAddressResponse) SetHeaders(v map[string]*string) *AllocatePublicNetworkAddressResponse {
	s.Headers = v
	return s
}

func (s *AllocatePublicNetworkAddressResponse) SetStatusCode(v int32) *AllocatePublicNetworkAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *AllocatePublicNetworkAddressResponse) SetBody(v *AllocatePublicNetworkAddressResponseBody) *AllocatePublicNetworkAddressResponse {
	s.Body = v
	return s
}

type CheckVersionsOfComponentsRequest struct {
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Components           *string `json:"Components,omitempty" xml:"Components,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CheckVersionsOfComponentsRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckVersionsOfComponentsRequest) GoString() string {
	return s.String()
}

func (s *CheckVersionsOfComponentsRequest) SetClusterId(v string) *CheckVersionsOfComponentsRequest {
	s.ClusterId = &v
	return s
}

func (s *CheckVersionsOfComponentsRequest) SetComponents(v string) *CheckVersionsOfComponentsRequest {
	s.Components = &v
	return s
}

func (s *CheckVersionsOfComponentsRequest) SetOwnerId(v int64) *CheckVersionsOfComponentsRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckVersionsOfComponentsRequest) SetRegionId(v string) *CheckVersionsOfComponentsRequest {
	s.RegionId = &v
	return s
}

func (s *CheckVersionsOfComponentsRequest) SetResourceOwnerAccount(v string) *CheckVersionsOfComponentsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckVersionsOfComponentsRequest) SetResourceOwnerId(v int64) *CheckVersionsOfComponentsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckVersionsOfComponentsRequest) SetZoneId(v string) *CheckVersionsOfComponentsRequest {
	s.ZoneId = &v
	return s
}

type CheckVersionsOfComponentsResponseBody struct {
	Components *CheckVersionsOfComponentsResponseBodyComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Struct"`
	RequestId  *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckVersionsOfComponentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckVersionsOfComponentsResponseBody) GoString() string {
	return s.String()
}

func (s *CheckVersionsOfComponentsResponseBody) SetComponents(v *CheckVersionsOfComponentsResponseBodyComponents) *CheckVersionsOfComponentsResponseBody {
	s.Components = v
	return s
}

func (s *CheckVersionsOfComponentsResponseBody) SetRequestId(v string) *CheckVersionsOfComponentsResponseBody {
	s.RequestId = &v
	return s
}

type CheckVersionsOfComponentsResponseBodyComponents struct {
	Components []*CheckVersionsOfComponentsResponseBodyComponentsComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
}

func (s CheckVersionsOfComponentsResponseBodyComponents) String() string {
	return tea.Prettify(s)
}

func (s CheckVersionsOfComponentsResponseBodyComponents) GoString() string {
	return s.String()
}

func (s *CheckVersionsOfComponentsResponseBodyComponents) SetComponents(v []*CheckVersionsOfComponentsResponseBodyComponentsComponents) *CheckVersionsOfComponentsResponseBodyComponents {
	s.Components = v
	return s
}

type CheckVersionsOfComponentsResponseBodyComponentsComponents struct {
	Component       *string `json:"Component,omitempty" xml:"Component,omitempty"`
	IsLatestVersion *string `json:"IsLatestVersion,omitempty" xml:"IsLatestVersion,omitempty"`
}

func (s CheckVersionsOfComponentsResponseBodyComponentsComponents) String() string {
	return tea.Prettify(s)
}

func (s CheckVersionsOfComponentsResponseBodyComponentsComponents) GoString() string {
	return s.String()
}

func (s *CheckVersionsOfComponentsResponseBodyComponentsComponents) SetComponent(v string) *CheckVersionsOfComponentsResponseBodyComponentsComponents {
	s.Component = &v
	return s
}

func (s *CheckVersionsOfComponentsResponseBodyComponentsComponents) SetIsLatestVersion(v string) *CheckVersionsOfComponentsResponseBodyComponentsComponents {
	s.IsLatestVersion = &v
	return s
}

type CheckVersionsOfComponentsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckVersionsOfComponentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckVersionsOfComponentsResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckVersionsOfComponentsResponse) GoString() string {
	return s.String()
}

func (s *CheckVersionsOfComponentsResponse) SetHeaders(v map[string]*string) *CheckVersionsOfComponentsResponse {
	s.Headers = v
	return s
}

func (s *CheckVersionsOfComponentsResponse) SetStatusCode(v int32) *CheckVersionsOfComponentsResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckVersionsOfComponentsResponse) SetBody(v *CheckVersionsOfComponentsResponseBody) *CheckVersionsOfComponentsResponse {
	s.Body = v
	return s
}

type CloseBackupRequest struct {
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CloseBackupRequest) String() string {
	return tea.Prettify(s)
}

func (s CloseBackupRequest) GoString() string {
	return s.String()
}

func (s *CloseBackupRequest) SetClusterId(v string) *CloseBackupRequest {
	s.ClusterId = &v
	return s
}

func (s *CloseBackupRequest) SetOwnerId(v int64) *CloseBackupRequest {
	s.OwnerId = &v
	return s
}

func (s *CloseBackupRequest) SetRegionId(v string) *CloseBackupRequest {
	s.RegionId = &v
	return s
}

func (s *CloseBackupRequest) SetResourceOwnerAccount(v string) *CloseBackupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloseBackupRequest) SetResourceOwnerId(v int64) *CloseBackupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloseBackupRequest) SetZoneId(v string) *CloseBackupRequest {
	s.ZoneId = &v
	return s
}

type CloseBackupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloseBackupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloseBackupResponseBody) GoString() string {
	return s.String()
}

func (s *CloseBackupResponseBody) SetRequestId(v string) *CloseBackupResponseBody {
	s.RequestId = &v
	return s
}

type CloseBackupResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseBackupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloseBackupResponse) String() string {
	return tea.Prettify(s)
}

func (s CloseBackupResponse) GoString() string {
	return s.String()
}

func (s *CloseBackupResponse) SetHeaders(v map[string]*string) *CloseBackupResponse {
	s.Headers = v
	return s
}

func (s *CloseBackupResponse) SetStatusCode(v int32) *CloseBackupResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseBackupResponse) SetBody(v *CloseBackupResponseBody) *CloseBackupResponse {
	s.Body = v
	return s
}

type ConvertClusterRequest struct {
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Duration             *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PricingCycle         *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ConvertClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s ConvertClusterRequest) GoString() string {
	return s.String()
}

func (s *ConvertClusterRequest) SetClusterId(v string) *ConvertClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *ConvertClusterRequest) SetDuration(v int32) *ConvertClusterRequest {
	s.Duration = &v
	return s
}

func (s *ConvertClusterRequest) SetOwnerId(v int64) *ConvertClusterRequest {
	s.OwnerId = &v
	return s
}

func (s *ConvertClusterRequest) SetPricingCycle(v string) *ConvertClusterRequest {
	s.PricingCycle = &v
	return s
}

func (s *ConvertClusterRequest) SetResourceOwnerAccount(v string) *ConvertClusterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ConvertClusterRequest) SetResourceOwnerId(v int64) *ConvertClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

type ConvertClusterResponseBody struct {
	OrderId   *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConvertClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConvertClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ConvertClusterResponseBody) SetOrderId(v int64) *ConvertClusterResponseBody {
	s.OrderId = &v
	return s
}

func (s *ConvertClusterResponseBody) SetRequestId(v string) *ConvertClusterResponseBody {
	s.RequestId = &v
	return s
}

type ConvertClusterResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConvertClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConvertClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s ConvertClusterResponse) GoString() string {
	return s.String()
}

func (s *ConvertClusterResponse) SetHeaders(v map[string]*string) *ConvertClusterResponse {
	s.Headers = v
	return s
}

func (s *ConvertClusterResponse) SetStatusCode(v int32) *ConvertClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *ConvertClusterResponse) SetBody(v *ConvertClusterResponseBody) *ConvertClusterResponse {
	s.Body = v
	return s
}

type CreateClusterRequest struct {
	AutoRenew            *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	BackupId             *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CloudType            *string `json:"CloudType,omitempty" xml:"CloudType,omitempty"`
	ClusterName          *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	ColdStorageSize      *string `json:"ColdStorageSize,omitempty" xml:"ColdStorageSize,omitempty"`
	CoreDiskQuantity     *string `json:"CoreDiskQuantity,omitempty" xml:"CoreDiskQuantity,omitempty"`
	CoreDiskSize         *string `json:"CoreDiskSize,omitempty" xml:"CoreDiskSize,omitempty"`
	CoreDiskType         *string `json:"CoreDiskType,omitempty" xml:"CoreDiskType,omitempty"`
	CoreInstanceQuantity *string `json:"CoreInstanceQuantity,omitempty" xml:"CoreInstanceQuantity,omitempty"`
	CoreInstanceType     *string `json:"CoreInstanceType,omitempty" xml:"CoreInstanceType,omitempty"`
	DbInstanceConnType   *string `json:"DbInstanceConnType,omitempty" xml:"DbInstanceConnType,omitempty"`
	DbInstanceType       *string `json:"DbInstanceType,omitempty" xml:"DbInstanceType,omitempty"`
	DbType               *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	DepMode              *string `json:"DepMode,omitempty" xml:"DepMode,omitempty"`
	Duration             *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Engine               *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion        *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	IsColdStorage        *string `json:"IsColdStorage,omitempty" xml:"IsColdStorage,omitempty"`
	MasterInstanceType   *string `json:"MasterInstanceType,omitempty" xml:"MasterInstanceType,omitempty"`
	NetType              *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	PayType              *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	PricingCycle         *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RestoreTime          *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	SecurityIPList       *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	SrcDBInstanceId      *string `json:"SrcDBInstanceId,omitempty" xml:"SrcDBInstanceId,omitempty"`
	VSwitchId            *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId                *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterRequest) SetAutoRenew(v string) *CreateClusterRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateClusterRequest) SetBackupId(v string) *CreateClusterRequest {
	s.BackupId = &v
	return s
}

func (s *CreateClusterRequest) SetClientToken(v string) *CreateClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateClusterRequest) SetCloudType(v string) *CreateClusterRequest {
	s.CloudType = &v
	return s
}

func (s *CreateClusterRequest) SetClusterName(v string) *CreateClusterRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateClusterRequest) SetColdStorageSize(v string) *CreateClusterRequest {
	s.ColdStorageSize = &v
	return s
}

func (s *CreateClusterRequest) SetCoreDiskQuantity(v string) *CreateClusterRequest {
	s.CoreDiskQuantity = &v
	return s
}

func (s *CreateClusterRequest) SetCoreDiskSize(v string) *CreateClusterRequest {
	s.CoreDiskSize = &v
	return s
}

func (s *CreateClusterRequest) SetCoreDiskType(v string) *CreateClusterRequest {
	s.CoreDiskType = &v
	return s
}

func (s *CreateClusterRequest) SetCoreInstanceQuantity(v string) *CreateClusterRequest {
	s.CoreInstanceQuantity = &v
	return s
}

func (s *CreateClusterRequest) SetCoreInstanceType(v string) *CreateClusterRequest {
	s.CoreInstanceType = &v
	return s
}

func (s *CreateClusterRequest) SetDbInstanceConnType(v string) *CreateClusterRequest {
	s.DbInstanceConnType = &v
	return s
}

func (s *CreateClusterRequest) SetDbInstanceType(v string) *CreateClusterRequest {
	s.DbInstanceType = &v
	return s
}

func (s *CreateClusterRequest) SetDbType(v string) *CreateClusterRequest {
	s.DbType = &v
	return s
}

func (s *CreateClusterRequest) SetDepMode(v string) *CreateClusterRequest {
	s.DepMode = &v
	return s
}

func (s *CreateClusterRequest) SetDuration(v string) *CreateClusterRequest {
	s.Duration = &v
	return s
}

func (s *CreateClusterRequest) SetEngine(v string) *CreateClusterRequest {
	s.Engine = &v
	return s
}

func (s *CreateClusterRequest) SetEngineVersion(v string) *CreateClusterRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateClusterRequest) SetIsColdStorage(v string) *CreateClusterRequest {
	s.IsColdStorage = &v
	return s
}

func (s *CreateClusterRequest) SetMasterInstanceType(v string) *CreateClusterRequest {
	s.MasterInstanceType = &v
	return s
}

func (s *CreateClusterRequest) SetNetType(v string) *CreateClusterRequest {
	s.NetType = &v
	return s
}

func (s *CreateClusterRequest) SetPayType(v string) *CreateClusterRequest {
	s.PayType = &v
	return s
}

func (s *CreateClusterRequest) SetPricingCycle(v string) *CreateClusterRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateClusterRequest) SetRegionId(v string) *CreateClusterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateClusterRequest) SetRestoreTime(v string) *CreateClusterRequest {
	s.RestoreTime = &v
	return s
}

func (s *CreateClusterRequest) SetSecurityIPList(v string) *CreateClusterRequest {
	s.SecurityIPList = &v
	return s
}

func (s *CreateClusterRequest) SetSrcDBInstanceId(v string) *CreateClusterRequest {
	s.SrcDBInstanceId = &v
	return s
}

func (s *CreateClusterRequest) SetVSwitchId(v string) *CreateClusterRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateClusterRequest) SetVpcId(v string) *CreateClusterRequest {
	s.VpcId = &v
	return s
}

func (s *CreateClusterRequest) SetZoneId(v string) *CreateClusterRequest {
	s.ZoneId = &v
	return s
}

type CreateClusterResponseBody struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClusterResponseBody) SetClusterId(v string) *CreateClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateClusterResponseBody) SetOrderId(v string) *CreateClusterResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateClusterResponseBody) SetRequestId(v string) *CreateClusterResponseBody {
	s.RequestId = &v
	return s
}

type CreateClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateClusterResponse) SetHeaders(v map[string]*string) *CreateClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateClusterResponse) SetStatusCode(v int32) *CreateClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateClusterResponse) SetBody(v *CreateClusterResponseBody) *CreateClusterResponse {
	s.Body = v
	return s
}

type CreateGlobalResourceRequest struct {
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceName         *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceType         *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateGlobalResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGlobalResourceRequest) GoString() string {
	return s.String()
}

func (s *CreateGlobalResourceRequest) SetClusterId(v string) *CreateGlobalResourceRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateGlobalResourceRequest) SetOwnerId(v int64) *CreateGlobalResourceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateGlobalResourceRequest) SetRegionId(v string) *CreateGlobalResourceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateGlobalResourceRequest) SetResourceName(v string) *CreateGlobalResourceRequest {
	s.ResourceName = &v
	return s
}

func (s *CreateGlobalResourceRequest) SetResourceOwnerAccount(v string) *CreateGlobalResourceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateGlobalResourceRequest) SetResourceOwnerId(v int64) *CreateGlobalResourceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateGlobalResourceRequest) SetResourceType(v string) *CreateGlobalResourceRequest {
	s.ResourceType = &v
	return s
}

func (s *CreateGlobalResourceRequest) SetZoneId(v string) *CreateGlobalResourceRequest {
	s.ZoneId = &v
	return s
}

type CreateGlobalResourceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateGlobalResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGlobalResourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGlobalResourceResponseBody) SetRequestId(v string) *CreateGlobalResourceResponseBody {
	s.RequestId = &v
	return s
}

type CreateGlobalResourceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGlobalResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGlobalResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGlobalResourceResponse) GoString() string {
	return s.String()
}

func (s *CreateGlobalResourceResponse) SetHeaders(v map[string]*string) *CreateGlobalResourceResponse {
	s.Headers = v
	return s
}

func (s *CreateGlobalResourceResponse) SetStatusCode(v int32) *CreateGlobalResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGlobalResourceResponse) SetBody(v *CreateGlobalResourceResponseBody) *CreateGlobalResourceResponse {
	s.Body = v
	return s
}

type CreateHbaseSlbServerRequest struct {
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SlbServer            *string `json:"SlbServer,omitempty" xml:"SlbServer,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateHbaseSlbServerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateHbaseSlbServerRequest) GoString() string {
	return s.String()
}

func (s *CreateHbaseSlbServerRequest) SetClusterId(v string) *CreateHbaseSlbServerRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateHbaseSlbServerRequest) SetOwnerId(v int64) *CreateHbaseSlbServerRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateHbaseSlbServerRequest) SetRegionId(v string) *CreateHbaseSlbServerRequest {
	s.RegionId = &v
	return s
}

func (s *CreateHbaseSlbServerRequest) SetResourceOwnerAccount(v string) *CreateHbaseSlbServerRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateHbaseSlbServerRequest) SetResourceOwnerId(v int64) *CreateHbaseSlbServerRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateHbaseSlbServerRequest) SetSlbServer(v string) *CreateHbaseSlbServerRequest {
	s.SlbServer = &v
	return s
}

func (s *CreateHbaseSlbServerRequest) SetZoneId(v string) *CreateHbaseSlbServerRequest {
	s.ZoneId = &v
	return s
}

type CreateHbaseSlbServerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateHbaseSlbServerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateHbaseSlbServerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHbaseSlbServerResponseBody) SetRequestId(v string) *CreateHbaseSlbServerResponseBody {
	s.RequestId = &v
	return s
}

type CreateHbaseSlbServerResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHbaseSlbServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHbaseSlbServerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateHbaseSlbServerResponse) GoString() string {
	return s.String()
}

func (s *CreateHbaseSlbServerResponse) SetHeaders(v map[string]*string) *CreateHbaseSlbServerResponse {
	s.Headers = v
	return s
}

func (s *CreateHbaseSlbServerResponse) SetStatusCode(v int32) *CreateHbaseSlbServerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHbaseSlbServerResponse) SetBody(v *CreateHbaseSlbServerResponseBody) *CreateHbaseSlbServerResponse {
	s.Body = v
	return s
}

type CreateSubscriptionRequest struct {
	DestinationInstanceId       *string `json:"DestinationInstanceId,omitempty" xml:"DestinationInstanceId,omitempty"`
	DestinationInstanceRegionId *string `json:"DestinationInstanceRegionId,omitempty" xml:"DestinationInstanceRegionId,omitempty"`
	ExtraContext                *string `json:"ExtraContext,omitempty" xml:"ExtraContext,omitempty"`
	Mapping                     *string `json:"Mapping,omitempty" xml:"Mapping,omitempty"`
	OwnerId                     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount        *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId             *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SlbServer                   *string `json:"SlbServer,omitempty" xml:"SlbServer,omitempty"`
	SourceInstanceId            *string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
	SourceInstanceRegionId      *string `json:"SourceInstanceRegionId,omitempty" xml:"SourceInstanceRegionId,omitempty"`
	SubscriptionDescription     *string `json:"SubscriptionDescription,omitempty" xml:"SubscriptionDescription,omitempty"`
	SubscriptionType            *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	ZoneId                      *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateSubscriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSubscriptionRequest) GoString() string {
	return s.String()
}

func (s *CreateSubscriptionRequest) SetDestinationInstanceId(v string) *CreateSubscriptionRequest {
	s.DestinationInstanceId = &v
	return s
}

func (s *CreateSubscriptionRequest) SetDestinationInstanceRegionId(v string) *CreateSubscriptionRequest {
	s.DestinationInstanceRegionId = &v
	return s
}

func (s *CreateSubscriptionRequest) SetExtraContext(v string) *CreateSubscriptionRequest {
	s.ExtraContext = &v
	return s
}

func (s *CreateSubscriptionRequest) SetMapping(v string) *CreateSubscriptionRequest {
	s.Mapping = &v
	return s
}

func (s *CreateSubscriptionRequest) SetOwnerId(v int64) *CreateSubscriptionRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSubscriptionRequest) SetResourceOwnerAccount(v string) *CreateSubscriptionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateSubscriptionRequest) SetResourceOwnerId(v int64) *CreateSubscriptionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateSubscriptionRequest) SetSlbServer(v string) *CreateSubscriptionRequest {
	s.SlbServer = &v
	return s
}

func (s *CreateSubscriptionRequest) SetSourceInstanceId(v string) *CreateSubscriptionRequest {
	s.SourceInstanceId = &v
	return s
}

func (s *CreateSubscriptionRequest) SetSourceInstanceRegionId(v string) *CreateSubscriptionRequest {
	s.SourceInstanceRegionId = &v
	return s
}

func (s *CreateSubscriptionRequest) SetSubscriptionDescription(v string) *CreateSubscriptionRequest {
	s.SubscriptionDescription = &v
	return s
}

func (s *CreateSubscriptionRequest) SetSubscriptionType(v string) *CreateSubscriptionRequest {
	s.SubscriptionType = &v
	return s
}

func (s *CreateSubscriptionRequest) SetZoneId(v string) *CreateSubscriptionRequest {
	s.ZoneId = &v
	return s
}

type CreateSubscriptionResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubscriptionId *string `json:"SubscriptionId,omitempty" xml:"SubscriptionId,omitempty"`
}

func (s CreateSubscriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSubscriptionResponseBody) SetRequestId(v string) *CreateSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSubscriptionResponseBody) SetSubscriptionId(v string) *CreateSubscriptionResponseBody {
	s.SubscriptionId = &v
	return s
}

type CreateSubscriptionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSubscriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSubscriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSubscriptionResponse) GoString() string {
	return s.String()
}

func (s *CreateSubscriptionResponse) SetHeaders(v map[string]*string) *CreateSubscriptionResponse {
	s.Headers = v
	return s
}

func (s *CreateSubscriptionResponse) SetStatusCode(v int32) *CreateSubscriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSubscriptionResponse) SetBody(v *CreateSubscriptionResponseBody) *CreateSubscriptionResponse {
	s.Body = v
	return s
}

type DeleteClusterRequest struct {
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DeleteClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteClusterRequest) SetClientToken(v string) *DeleteClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteClusterRequest) SetClusterId(v string) *DeleteClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteClusterRequest) SetOwnerId(v int64) *DeleteClusterRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteClusterRequest) SetRegionId(v string) *DeleteClusterRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteClusterRequest) SetResourceOwnerAccount(v string) *DeleteClusterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteClusterRequest) SetResourceOwnerId(v int64) *DeleteClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteClusterRequest) SetZoneId(v string) *DeleteClusterRequest {
	s.ZoneId = &v
	return s
}

type DeleteClusterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponseBody) SetRequestId(v string) *DeleteClusterResponseBody {
	s.RequestId = &v
	return s
}

type DeleteClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponse) SetHeaders(v map[string]*string) *DeleteClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteClusterResponse) SetStatusCode(v int32) *DeleteClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteClusterResponse) SetBody(v *DeleteClusterResponseBody) *DeleteClusterResponse {
	s.Body = v
	return s
}

type DeleteGlobalResourceRequest struct {
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceName         *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceType         *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DeleteGlobalResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGlobalResourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteGlobalResourceRequest) SetClusterId(v string) *DeleteGlobalResourceRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteGlobalResourceRequest) SetOwnerId(v int64) *DeleteGlobalResourceRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteGlobalResourceRequest) SetRegionId(v string) *DeleteGlobalResourceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteGlobalResourceRequest) SetResourceName(v string) *DeleteGlobalResourceRequest {
	s.ResourceName = &v
	return s
}

func (s *DeleteGlobalResourceRequest) SetResourceOwnerAccount(v string) *DeleteGlobalResourceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteGlobalResourceRequest) SetResourceOwnerId(v int64) *DeleteGlobalResourceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteGlobalResourceRequest) SetResourceType(v string) *DeleteGlobalResourceRequest {
	s.ResourceType = &v
	return s
}

func (s *DeleteGlobalResourceRequest) SetZoneId(v string) *DeleteGlobalResourceRequest {
	s.ZoneId = &v
	return s
}

type DeleteGlobalResourceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGlobalResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGlobalResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGlobalResourceResponseBody) SetRequestId(v string) *DeleteGlobalResourceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteGlobalResourceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGlobalResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGlobalResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGlobalResourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteGlobalResourceResponse) SetHeaders(v map[string]*string) *DeleteGlobalResourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteGlobalResourceResponse) SetStatusCode(v int32) *DeleteGlobalResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGlobalResourceResponse) SetBody(v *DeleteGlobalResourceResponseBody) *DeleteGlobalResourceResponse {
	s.Body = v
	return s
}

type DeleteHbaseSlbServerRequest struct {
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SlbServer            *string `json:"SlbServer,omitempty" xml:"SlbServer,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DeleteHbaseSlbServerRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteHbaseSlbServerRequest) GoString() string {
	return s.String()
}

func (s *DeleteHbaseSlbServerRequest) SetClusterId(v string) *DeleteHbaseSlbServerRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteHbaseSlbServerRequest) SetOwnerId(v int64) *DeleteHbaseSlbServerRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteHbaseSlbServerRequest) SetRegionId(v string) *DeleteHbaseSlbServerRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteHbaseSlbServerRequest) SetResourceOwnerAccount(v string) *DeleteHbaseSlbServerRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteHbaseSlbServerRequest) SetResourceOwnerId(v int64) *DeleteHbaseSlbServerRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteHbaseSlbServerRequest) SetSlbServer(v string) *DeleteHbaseSlbServerRequest {
	s.SlbServer = &v
	return s
}

func (s *DeleteHbaseSlbServerRequest) SetZoneId(v string) *DeleteHbaseSlbServerRequest {
	s.ZoneId = &v
	return s
}

type DeleteHbaseSlbServerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteHbaseSlbServerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteHbaseSlbServerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHbaseSlbServerResponseBody) SetRequestId(v string) *DeleteHbaseSlbServerResponseBody {
	s.RequestId = &v
	return s
}

type DeleteHbaseSlbServerResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHbaseSlbServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHbaseSlbServerResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteHbaseSlbServerResponse) GoString() string {
	return s.String()
}

func (s *DeleteHbaseSlbServerResponse) SetHeaders(v map[string]*string) *DeleteHbaseSlbServerResponse {
	s.Headers = v
	return s
}

func (s *DeleteHbaseSlbServerResponse) SetStatusCode(v int32) *DeleteHbaseSlbServerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHbaseSlbServerResponse) SetBody(v *DeleteHbaseSlbServerResponseBody) *DeleteHbaseSlbServerResponse {
	s.Body = v
	return s
}

type DeleteServerlessInstanceRequest struct {
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DeleteServerlessInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteServerlessInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteServerlessInstanceRequest) SetClientToken(v string) *DeleteServerlessInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteServerlessInstanceRequest) SetInstanceId(v string) *DeleteServerlessInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteServerlessInstanceRequest) SetOwnerId(v int64) *DeleteServerlessInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteServerlessInstanceRequest) SetRegionId(v string) *DeleteServerlessInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteServerlessInstanceRequest) SetResourceOwnerAccount(v string) *DeleteServerlessInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteServerlessInstanceRequest) SetResourceOwnerId(v int64) *DeleteServerlessInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteServerlessInstanceRequest) SetZoneId(v string) *DeleteServerlessInstanceRequest {
	s.ZoneId = &v
	return s
}

type DeleteServerlessInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServerlessInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteServerlessInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServerlessInstanceResponseBody) SetRequestId(v string) *DeleteServerlessInstanceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteServerlessInstanceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteServerlessInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteServerlessInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServerlessInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteServerlessInstanceResponse) SetHeaders(v map[string]*string) *DeleteServerlessInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteServerlessInstanceResponse) SetStatusCode(v int32) *DeleteServerlessInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServerlessInstanceResponse) SetBody(v *DeleteServerlessInstanceResponseBody) *DeleteServerlessInstanceResponse {
	s.Body = v
	return s
}

type DeleteUserHdfsInfoRequest struct {
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	NameService          *string `json:"NameService,omitempty" xml:"NameService,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DeleteUserHdfsInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserHdfsInfoRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserHdfsInfoRequest) SetClusterId(v string) *DeleteUserHdfsInfoRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteUserHdfsInfoRequest) SetNameService(v string) *DeleteUserHdfsInfoRequest {
	s.NameService = &v
	return s
}

func (s *DeleteUserHdfsInfoRequest) SetOwnerId(v int64) *DeleteUserHdfsInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteUserHdfsInfoRequest) SetRegionId(v string) *DeleteUserHdfsInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteUserHdfsInfoRequest) SetResourceOwnerAccount(v string) *DeleteUserHdfsInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteUserHdfsInfoRequest) SetResourceOwnerId(v int64) *DeleteUserHdfsInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteUserHdfsInfoRequest) SetZoneId(v string) *DeleteUserHdfsInfoRequest {
	s.ZoneId = &v
	return s
}

type DeleteUserHdfsInfoResponseBody struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserHdfsInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserHdfsInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserHdfsInfoResponseBody) SetClusterId(v string) *DeleteUserHdfsInfoResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DeleteUserHdfsInfoResponseBody) SetRequestId(v string) *DeleteUserHdfsInfoResponseBody {
	s.RequestId = &v
	return s
}

type DeleteUserHdfsInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserHdfsInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserHdfsInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserHdfsInfoResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserHdfsInfoResponse) SetHeaders(v map[string]*string) *DeleteUserHdfsInfoResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserHdfsInfoResponse) SetStatusCode(v int32) *DeleteUserHdfsInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserHdfsInfoResponse) SetBody(v *DeleteUserHdfsInfoResponseBody) *DeleteUserHdfsInfoResponse {
	s.Body = v
	return s
}

type DescribeBackupPolicyRequest struct {
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeBackupPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyRequest) SetClusterId(v string) *DescribeBackupPolicyRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetOwnerId(v int64) *DescribeBackupPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetRegionId(v string) *DescribeBackupPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetResourceOwnerAccount(v string) *DescribeBackupPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetResourceOwnerId(v int64) *DescribeBackupPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetZoneId(v string) *DescribeBackupPolicyRequest {
	s.ZoneId = &v
	return s
}

type DescribeBackupPolicyResponseBody struct {
	BackupRetentionPeriod       *string `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	PreferredBackupEndTimeUTC   *string `json:"PreferredBackupEndTimeUTC,omitempty" xml:"PreferredBackupEndTimeUTC,omitempty"`
	PreferredBackupPeriod       *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	PreferredBackupStartTimeUTC *string `json:"PreferredBackupStartTimeUTC,omitempty" xml:"PreferredBackupStartTimeUTC,omitempty"`
	PreferredBackupTime         *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	RequestId                   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBackupPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBody) SetBackupRetentionPeriod(v string) *DescribeBackupPolicyResponseBody {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetPreferredBackupEndTimeUTC(v string) *DescribeBackupPolicyResponseBody {
	s.PreferredBackupEndTimeUTC = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetPreferredBackupPeriod(v string) *DescribeBackupPolicyResponseBody {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetPreferredBackupStartTimeUTC(v string) *DescribeBackupPolicyResponseBody {
	s.PreferredBackupStartTimeUTC = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetPreferredBackupTime(v string) *DescribeBackupPolicyResponseBody {
	s.PreferredBackupTime = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetRequestId(v string) *DescribeBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

type DescribeBackupPolicyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponse) SetHeaders(v map[string]*string) *DescribeBackupPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupPolicyResponse) SetStatusCode(v int32) *DescribeBackupPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupPolicyResponse) SetBody(v *DescribeBackupPolicyResponseBody) *DescribeBackupPolicyResponse {
	s.Body = v
	return s
}

type DescribeBackupsRequest struct {
	BackupId             *int32  `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EndTimeUTC           *string `json:"EndTimeUTC,omitempty" xml:"EndTimeUTC,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StartTimeUTC         *string `json:"StartTimeUTC,omitempty" xml:"StartTimeUTC,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeBackupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupsRequest) SetBackupId(v int32) *DescribeBackupsRequest {
	s.BackupId = &v
	return s
}

func (s *DescribeBackupsRequest) SetClusterId(v string) *DescribeBackupsRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeBackupsRequest) SetEndTime(v string) *DescribeBackupsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeBackupsRequest) SetEndTimeUTC(v string) *DescribeBackupsRequest {
	s.EndTimeUTC = &v
	return s
}

func (s *DescribeBackupsRequest) SetOwnerId(v int64) *DescribeBackupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBackupsRequest) SetPageNumber(v int32) *DescribeBackupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupsRequest) SetPageSize(v int32) *DescribeBackupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupsRequest) SetRegionId(v string) *DescribeBackupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeBackupsRequest) SetResourceOwnerAccount(v string) *DescribeBackupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeBackupsRequest) SetResourceOwnerId(v int64) *DescribeBackupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeBackupsRequest) SetStartTime(v string) *DescribeBackupsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeBackupsRequest) SetStartTimeUTC(v string) *DescribeBackupsRequest {
	s.StartTimeUTC = &v
	return s
}

func (s *DescribeBackupsRequest) SetZoneId(v string) *DescribeBackupsRequest {
	s.ZoneId = &v
	return s
}

type DescribeBackupsResponseBody struct {
	Backups      *DescribeBackupsResponseBodyBackups `json:"Backups,omitempty" xml:"Backups,omitempty" type:"Struct"`
	EnableStatus *string                             `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	PageNumber   *int32                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId    *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount   *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBackupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBody) SetBackups(v *DescribeBackupsResponseBodyBackups) *DescribeBackupsResponseBody {
	s.Backups = v
	return s
}

func (s *DescribeBackupsResponseBody) SetEnableStatus(v string) *DescribeBackupsResponseBody {
	s.EnableStatus = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetPageNumber(v int32) *DescribeBackupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetPageSize(v int32) *DescribeBackupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetRequestId(v string) *DescribeBackupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetTotalCount(v int32) *DescribeBackupsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeBackupsResponseBodyBackups struct {
	Backup []*DescribeBackupsResponseBodyBackupsBackup `json:"Backup,omitempty" xml:"Backup,omitempty" type:"Repeated"`
}

func (s DescribeBackupsResponseBodyBackups) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsResponseBodyBackups) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBodyBackups) SetBackup(v []*DescribeBackupsResponseBodyBackupsBackup) *DescribeBackupsResponseBodyBackups {
	s.Backup = v
	return s
}

type DescribeBackupsResponseBodyBackupsBackup struct {
	BackupDBNames      *string `json:"BackupDBNames,omitempty" xml:"BackupDBNames,omitempty"`
	BackupDownloadURL  *string `json:"BackupDownloadURL,omitempty" xml:"BackupDownloadURL,omitempty"`
	BackupEndTime      *string `json:"BackupEndTime,omitempty" xml:"BackupEndTime,omitempty"`
	BackupEndTimeUTC   *string `json:"BackupEndTimeUTC,omitempty" xml:"BackupEndTimeUTC,omitempty"`
	BackupId           *int32  `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	BackupMethod       *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	BackupMode         *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	BackupSize         *string `json:"BackupSize,omitempty" xml:"BackupSize,omitempty"`
	BackupStartTime    *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	BackupStartTimeUTC *string `json:"BackupStartTimeUTC,omitempty" xml:"BackupStartTimeUTC,omitempty"`
	BackupStatus       *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	BackupType         *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
}

func (s DescribeBackupsResponseBodyBackupsBackup) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsResponseBodyBackupsBackup) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupDBNames(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupDBNames = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupDownloadURL(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupDownloadURL = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupEndTime(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupEndTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupEndTimeUTC(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupEndTimeUTC = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupId(v int32) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupId = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupMethod(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupMethod = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupMode(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupMode = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupSize(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupSize = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupStartTime(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupStartTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupStartTimeUTC(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupStartTimeUTC = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupStatus(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupStatus = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupType(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupType = &v
	return s
}

type DescribeBackupsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponse) SetHeaders(v map[string]*string) *DescribeBackupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupsResponse) SetStatusCode(v int32) *DescribeBackupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupsResponse) SetBody(v *DescribeBackupsResponseBody) *DescribeBackupsResponse {
	s.Body = v
	return s
}

type DescribeClusterAttributeRequest struct {
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeClusterAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterAttributeRequest) SetClusterId(v string) *DescribeClusterAttributeRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterAttributeRequest) SetOwnerId(v int64) *DescribeClusterAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeClusterAttributeRequest) SetRegionId(v string) *DescribeClusterAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeClusterAttributeRequest) SetResourceOwnerAccount(v string) *DescribeClusterAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeClusterAttributeRequest) SetResourceOwnerId(v int64) *DescribeClusterAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeClusterAttributeRequest) SetZoneId(v string) *DescribeClusterAttributeRequest {
	s.ZoneId = &v
	return s
}

type DescribeClusterAttributeResponseBody struct {
	AutoRenew            *string                                             `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	ClusterId            *string                                             `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName          *string                                             `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	ClusterType          *string                                             `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	ColdStorageStatus    *string                                             `json:"ColdStorageStatus,omitempty" xml:"ColdStorageStatus,omitempty"`
	ConnectionInfo       *DescribeClusterAttributeResponseBodyConnectionInfo `json:"ConnectionInfo,omitempty" xml:"ConnectionInfo,omitempty" type:"Struct"`
	CoreDiskQuantity     *int32                                              `json:"CoreDiskQuantity,omitempty" xml:"CoreDiskQuantity,omitempty"`
	CoreDiskSize         *string                                             `json:"CoreDiskSize,omitempty" xml:"CoreDiskSize,omitempty"`
	CoreDiskType         *string                                             `json:"CoreDiskType,omitempty" xml:"CoreDiskType,omitempty"`
	CoreInstanceQuantity *int32                                              `json:"CoreInstanceQuantity,omitempty" xml:"CoreInstanceQuantity,omitempty"`
	CoreInstanceType     *string                                             `json:"CoreInstanceType,omitempty" xml:"CoreInstanceType,omitempty"`
	CreateTime           *string                                             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExpireTime           *string                                             `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	HaType               *string                                             `json:"HaType,omitempty" xml:"HaType,omitempty"`
	HasUser              *string                                             `json:"HasUser,omitempty" xml:"HasUser,omitempty"`
	LockMode             *string                                             `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	MainVersion          *string                                             `json:"MainVersion,omitempty" xml:"MainVersion,omitempty"`
	MasterDiskSize       *int32                                              `json:"MasterDiskSize,omitempty" xml:"MasterDiskSize,omitempty"`
	MasterDiskType       *string                                             `json:"MasterDiskType,omitempty" xml:"MasterDiskType,omitempty"`
	MasterInstanceType   *string                                             `json:"MasterInstanceType,omitempty" xml:"MasterInstanceType,omitempty"`
	MinorVersion         *string                                             `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	NetInfo              *DescribeClusterAttributeResponseBodyNetInfo        `json:"NetInfo,omitempty" xml:"NetInfo,omitempty" type:"Struct"`
	NodeList             *DescribeClusterAttributeResponseBodyNodeList       `json:"NodeList,omitempty" xml:"NodeList,omitempty" type:"Struct"`
	PayType              *string                                             `json:"PayType,omitempty" xml:"PayType,omitempty"`
	RegionId             *string                                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId            *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status               *string                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateStatus         *string                                             `json:"UpdateStatus,omitempty" xml:"UpdateStatus,omitempty"`
	ZoneId               *string                                             `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeClusterAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterAttributeResponseBody) SetAutoRenew(v string) *DescribeClusterAttributeResponseBody {
	s.AutoRenew = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetClusterId(v string) *DescribeClusterAttributeResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetClusterName(v string) *DescribeClusterAttributeResponseBody {
	s.ClusterName = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetClusterType(v string) *DescribeClusterAttributeResponseBody {
	s.ClusterType = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetColdStorageStatus(v string) *DescribeClusterAttributeResponseBody {
	s.ColdStorageStatus = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetConnectionInfo(v *DescribeClusterAttributeResponseBodyConnectionInfo) *DescribeClusterAttributeResponseBody {
	s.ConnectionInfo = v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetCoreDiskQuantity(v int32) *DescribeClusterAttributeResponseBody {
	s.CoreDiskQuantity = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetCoreDiskSize(v string) *DescribeClusterAttributeResponseBody {
	s.CoreDiskSize = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetCoreDiskType(v string) *DescribeClusterAttributeResponseBody {
	s.CoreDiskType = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetCoreInstanceQuantity(v int32) *DescribeClusterAttributeResponseBody {
	s.CoreInstanceQuantity = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetCoreInstanceType(v string) *DescribeClusterAttributeResponseBody {
	s.CoreInstanceType = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetCreateTime(v string) *DescribeClusterAttributeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetExpireTime(v string) *DescribeClusterAttributeResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetHaType(v string) *DescribeClusterAttributeResponseBody {
	s.HaType = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetHasUser(v string) *DescribeClusterAttributeResponseBody {
	s.HasUser = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetLockMode(v string) *DescribeClusterAttributeResponseBody {
	s.LockMode = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetMainVersion(v string) *DescribeClusterAttributeResponseBody {
	s.MainVersion = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetMasterDiskSize(v int32) *DescribeClusterAttributeResponseBody {
	s.MasterDiskSize = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetMasterDiskType(v string) *DescribeClusterAttributeResponseBody {
	s.MasterDiskType = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetMasterInstanceType(v string) *DescribeClusterAttributeResponseBody {
	s.MasterInstanceType = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetMinorVersion(v string) *DescribeClusterAttributeResponseBody {
	s.MinorVersion = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetNetInfo(v *DescribeClusterAttributeResponseBodyNetInfo) *DescribeClusterAttributeResponseBody {
	s.NetInfo = v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetNodeList(v *DescribeClusterAttributeResponseBodyNodeList) *DescribeClusterAttributeResponseBody {
	s.NodeList = v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetPayType(v string) *DescribeClusterAttributeResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetRegionId(v string) *DescribeClusterAttributeResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetRequestId(v string) *DescribeClusterAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetStatus(v string) *DescribeClusterAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetUpdateStatus(v string) *DescribeClusterAttributeResponseBody {
	s.UpdateStatus = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetZoneId(v string) *DescribeClusterAttributeResponseBody {
	s.ZoneId = &v
	return s
}

type DescribeClusterAttributeResponseBodyConnectionInfo struct {
	HaRestConnectionString        *string                                                                          `json:"HaRestConnectionString,omitempty" xml:"HaRestConnectionString,omitempty"`
	HaRestPort                    *string                                                                          `json:"HaRestPort,omitempty" xml:"HaRestPort,omitempty"`
	HaThriftConnectionString      *string                                                                          `json:"HaThriftConnectionString,omitempty" xml:"HaThriftConnectionString,omitempty"`
	HaThriftPort                  *string                                                                          `json:"HaThriftPort,omitempty" xml:"HaThriftPort,omitempty"`
	ThriftConnectionString        *string                                                                          `json:"ThriftConnectionString,omitempty" xml:"ThriftConnectionString,omitempty"`
	ThriftPort                    *string                                                                          `json:"ThriftPort,omitempty" xml:"ThriftPort,omitempty"`
	UIProxyConnectionString       *string                                                                          `json:"UIProxyConnectionString,omitempty" xml:"UIProxyConnectionString,omitempty"`
	ZKClassicConnectionStringList *DescribeClusterAttributeResponseBodyConnectionInfoZKClassicConnectionStringList `json:"ZKClassicConnectionStringList,omitempty" xml:"ZKClassicConnectionStringList,omitempty" type:"Struct"`
	ZKConnectionStringList        *DescribeClusterAttributeResponseBodyConnectionInfoZKConnectionStringList        `json:"ZKConnectionStringList,omitempty" xml:"ZKConnectionStringList,omitempty" type:"Struct"`
	ZKInnerConnectionStringList   *DescribeClusterAttributeResponseBodyConnectionInfoZKInnerConnectionStringList   `json:"ZKInnerConnectionStringList,omitempty" xml:"ZKInnerConnectionStringList,omitempty" type:"Struct"`
	ZKPort                        *int32                                                                           `json:"ZKPort,omitempty" xml:"ZKPort,omitempty"`
	ZKPublicConnectionStringList  *DescribeClusterAttributeResponseBodyConnectionInfoZKPublicConnectionStringList  `json:"ZKPublicConnectionStringList,omitempty" xml:"ZKPublicConnectionStringList,omitempty" type:"Struct"`
}

func (s DescribeClusterAttributeResponseBodyConnectionInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAttributeResponseBodyConnectionInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfo) SetHaRestConnectionString(v string) *DescribeClusterAttributeResponseBodyConnectionInfo {
	s.HaRestConnectionString = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfo) SetHaRestPort(v string) *DescribeClusterAttributeResponseBodyConnectionInfo {
	s.HaRestPort = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfo) SetHaThriftConnectionString(v string) *DescribeClusterAttributeResponseBodyConnectionInfo {
	s.HaThriftConnectionString = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfo) SetHaThriftPort(v string) *DescribeClusterAttributeResponseBodyConnectionInfo {
	s.HaThriftPort = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfo) SetThriftConnectionString(v string) *DescribeClusterAttributeResponseBodyConnectionInfo {
	s.ThriftConnectionString = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfo) SetThriftPort(v string) *DescribeClusterAttributeResponseBodyConnectionInfo {
	s.ThriftPort = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfo) SetUIProxyConnectionString(v string) *DescribeClusterAttributeResponseBodyConnectionInfo {
	s.UIProxyConnectionString = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfo) SetZKClassicConnectionStringList(v *DescribeClusterAttributeResponseBodyConnectionInfoZKClassicConnectionStringList) *DescribeClusterAttributeResponseBodyConnectionInfo {
	s.ZKClassicConnectionStringList = v
	return s
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfo) SetZKConnectionStringList(v *DescribeClusterAttributeResponseBodyConnectionInfoZKConnectionStringList) *DescribeClusterAttributeResponseBodyConnectionInfo {
	s.ZKConnectionStringList = v
	return s
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfo) SetZKInnerConnectionStringList(v *DescribeClusterAttributeResponseBodyConnectionInfoZKInnerConnectionStringList) *DescribeClusterAttributeResponseBodyConnectionInfo {
	s.ZKInnerConnectionStringList = v
	return s
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfo) SetZKPort(v int32) *DescribeClusterAttributeResponseBodyConnectionInfo {
	s.ZKPort = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfo) SetZKPublicConnectionStringList(v *DescribeClusterAttributeResponseBodyConnectionInfoZKPublicConnectionStringList) *DescribeClusterAttributeResponseBodyConnectionInfo {
	s.ZKPublicConnectionStringList = v
	return s
}

type DescribeClusterAttributeResponseBodyConnectionInfoZKClassicConnectionStringList struct {
	String_ []*string `json:"String,omitempty" xml:"String,omitempty" type:"Repeated"`
}

func (s DescribeClusterAttributeResponseBodyConnectionInfoZKClassicConnectionStringList) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAttributeResponseBodyConnectionInfoZKClassicConnectionStringList) GoString() string {
	return s.String()
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfoZKClassicConnectionStringList) SetString_(v []*string) *DescribeClusterAttributeResponseBodyConnectionInfoZKClassicConnectionStringList {
	s.String_ = v
	return s
}

type DescribeClusterAttributeResponseBodyConnectionInfoZKConnectionStringList struct {
	String_ []*string `json:"String,omitempty" xml:"String,omitempty" type:"Repeated"`
}

func (s DescribeClusterAttributeResponseBodyConnectionInfoZKConnectionStringList) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAttributeResponseBodyConnectionInfoZKConnectionStringList) GoString() string {
	return s.String()
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfoZKConnectionStringList) SetString_(v []*string) *DescribeClusterAttributeResponseBodyConnectionInfoZKConnectionStringList {
	s.String_ = v
	return s
}

type DescribeClusterAttributeResponseBodyConnectionInfoZKInnerConnectionStringList struct {
	String_ []*string `json:"String,omitempty" xml:"String,omitempty" type:"Repeated"`
}

func (s DescribeClusterAttributeResponseBodyConnectionInfoZKInnerConnectionStringList) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAttributeResponseBodyConnectionInfoZKInnerConnectionStringList) GoString() string {
	return s.String()
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfoZKInnerConnectionStringList) SetString_(v []*string) *DescribeClusterAttributeResponseBodyConnectionInfoZKInnerConnectionStringList {
	s.String_ = v
	return s
}

type DescribeClusterAttributeResponseBodyConnectionInfoZKPublicConnectionStringList struct {
	String_ []*string `json:"String,omitempty" xml:"String,omitempty" type:"Repeated"`
}

func (s DescribeClusterAttributeResponseBodyConnectionInfoZKPublicConnectionStringList) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAttributeResponseBodyConnectionInfoZKPublicConnectionStringList) GoString() string {
	return s.String()
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfoZKPublicConnectionStringList) SetString_(v []*string) *DescribeClusterAttributeResponseBodyConnectionInfoZKPublicConnectionStringList {
	s.String_ = v
	return s
}

type DescribeClusterAttributeResponseBodyNetInfo struct {
	InnerIpAddress  *string `json:"InnerIpAddress,omitempty" xml:"InnerIpAddress,omitempty"`
	NetType         *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	PublicIpAddress *string `json:"PublicIpAddress,omitempty" xml:"PublicIpAddress,omitempty"`
	SecurityIpList  *string `json:"SecurityIpList,omitempty" xml:"SecurityIpList,omitempty"`
	VSwitchId       *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId           *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeClusterAttributeResponseBodyNetInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAttributeResponseBodyNetInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterAttributeResponseBodyNetInfo) SetInnerIpAddress(v string) *DescribeClusterAttributeResponseBodyNetInfo {
	s.InnerIpAddress = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyNetInfo) SetNetType(v string) *DescribeClusterAttributeResponseBodyNetInfo {
	s.NetType = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyNetInfo) SetPublicIpAddress(v string) *DescribeClusterAttributeResponseBodyNetInfo {
	s.PublicIpAddress = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyNetInfo) SetSecurityIpList(v string) *DescribeClusterAttributeResponseBodyNetInfo {
	s.SecurityIpList = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyNetInfo) SetVSwitchId(v string) *DescribeClusterAttributeResponseBodyNetInfo {
	s.VSwitchId = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyNetInfo) SetVpcId(v string) *DescribeClusterAttributeResponseBodyNetInfo {
	s.VpcId = &v
	return s
}

type DescribeClusterAttributeResponseBodyNodeList struct {
	Node []*DescribeClusterAttributeResponseBodyNodeListNode `json:"Node,omitempty" xml:"Node,omitempty" type:"Repeated"`
}

func (s DescribeClusterAttributeResponseBodyNodeList) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAttributeResponseBodyNodeList) GoString() string {
	return s.String()
}

func (s *DescribeClusterAttributeResponseBodyNodeList) SetNode(v []*DescribeClusterAttributeResponseBodyNodeListNode) *DescribeClusterAttributeResponseBodyNodeList {
	s.Node = v
	return s
}

type DescribeClusterAttributeResponseBodyNodeListNode struct {
	DaemonList       *DescribeClusterAttributeResponseBodyNodeListNodeDaemonList `json:"DaemonList,omitempty" xml:"DaemonList,omitempty" type:"Struct"`
	MemStore         *string                                                     `json:"MemStore,omitempty" xml:"MemStore,omitempty"`
	NodeDiskQuantity *string                                                     `json:"NodeDiskQuantity,omitempty" xml:"NodeDiskQuantity,omitempty"`
	NodeDiskSize     *string                                                     `json:"NodeDiskSize,omitempty" xml:"NodeDiskSize,omitempty"`
	NodeDiskType     *string                                                     `json:"NodeDiskType,omitempty" xml:"NodeDiskType,omitempty"`
	NodeId           *string                                                     `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	NodeInstanceType *string                                                     `json:"NodeInstanceType,omitempty" xml:"NodeInstanceType,omitempty"`
	NodeStatus       *string                                                     `json:"NodeStatus,omitempty" xml:"NodeStatus,omitempty"`
	NodeType         *string                                                     `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	RegionQuantity   *string                                                     `json:"RegionQuantity,omitempty" xml:"RegionQuantity,omitempty"`
	ServiceType      *string                                                     `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	StoreFile        *string                                                     `json:"StoreFile,omitempty" xml:"StoreFile,omitempty"`
}

func (s DescribeClusterAttributeResponseBodyNodeListNode) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAttributeResponseBodyNodeListNode) GoString() string {
	return s.String()
}

func (s *DescribeClusterAttributeResponseBodyNodeListNode) SetDaemonList(v *DescribeClusterAttributeResponseBodyNodeListNodeDaemonList) *DescribeClusterAttributeResponseBodyNodeListNode {
	s.DaemonList = v
	return s
}

func (s *DescribeClusterAttributeResponseBodyNodeListNode) SetMemStore(v string) *DescribeClusterAttributeResponseBodyNodeListNode {
	s.MemStore = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyNodeListNode) SetNodeDiskQuantity(v string) *DescribeClusterAttributeResponseBodyNodeListNode {
	s.NodeDiskQuantity = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyNodeListNode) SetNodeDiskSize(v string) *DescribeClusterAttributeResponseBodyNodeListNode {
	s.NodeDiskSize = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyNodeListNode) SetNodeDiskType(v string) *DescribeClusterAttributeResponseBodyNodeListNode {
	s.NodeDiskType = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyNodeListNode) SetNodeId(v string) *DescribeClusterAttributeResponseBodyNodeListNode {
	s.NodeId = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyNodeListNode) SetNodeInstanceType(v string) *DescribeClusterAttributeResponseBodyNodeListNode {
	s.NodeInstanceType = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyNodeListNode) SetNodeStatus(v string) *DescribeClusterAttributeResponseBodyNodeListNode {
	s.NodeStatus = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyNodeListNode) SetNodeType(v string) *DescribeClusterAttributeResponseBodyNodeListNode {
	s.NodeType = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyNodeListNode) SetRegionQuantity(v string) *DescribeClusterAttributeResponseBodyNodeListNode {
	s.RegionQuantity = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyNodeListNode) SetServiceType(v string) *DescribeClusterAttributeResponseBodyNodeListNode {
	s.ServiceType = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyNodeListNode) SetStoreFile(v string) *DescribeClusterAttributeResponseBodyNodeListNode {
	s.StoreFile = &v
	return s
}

type DescribeClusterAttributeResponseBodyNodeListNodeDaemonList struct {
	Daemon []*DescribeClusterAttributeResponseBodyNodeListNodeDaemonListDaemon `json:"Daemon,omitempty" xml:"Daemon,omitempty" type:"Repeated"`
}

func (s DescribeClusterAttributeResponseBodyNodeListNodeDaemonList) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAttributeResponseBodyNodeListNodeDaemonList) GoString() string {
	return s.String()
}

func (s *DescribeClusterAttributeResponseBodyNodeListNodeDaemonList) SetDaemon(v []*DescribeClusterAttributeResponseBodyNodeListNodeDaemonListDaemon) *DescribeClusterAttributeResponseBodyNodeListNodeDaemonList {
	s.Daemon = v
	return s
}

type DescribeClusterAttributeResponseBodyNodeListNodeDaemonListDaemon struct {
	DaemonName   *string `json:"DaemonName,omitempty" xml:"DaemonName,omitempty"`
	DaemonStatus *string `json:"DaemonStatus,omitempty" xml:"DaemonStatus,omitempty"`
}

func (s DescribeClusterAttributeResponseBodyNodeListNodeDaemonListDaemon) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAttributeResponseBodyNodeListNodeDaemonListDaemon) GoString() string {
	return s.String()
}

func (s *DescribeClusterAttributeResponseBodyNodeListNodeDaemonListDaemon) SetDaemonName(v string) *DescribeClusterAttributeResponseBodyNodeListNodeDaemonListDaemon {
	s.DaemonName = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyNodeListNodeDaemonListDaemon) SetDaemonStatus(v string) *DescribeClusterAttributeResponseBodyNodeListNodeDaemonListDaemon {
	s.DaemonStatus = &v
	return s
}

type DescribeClusterAttributeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterAttributeResponse) SetHeaders(v map[string]*string) *DescribeClusterAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterAttributeResponse) SetStatusCode(v int32) *DescribeClusterAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterAttributeResponse) SetBody(v *DescribeClusterAttributeResponseBody) *DescribeClusterAttributeResponse {
	s.Body = v
	return s
}

type DescribeClusterConnectAddrsRequest struct {
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeClusterConnectAddrsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterConnectAddrsRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectAddrsRequest) SetClusterId(v string) *DescribeClusterConnectAddrsRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterConnectAddrsRequest) SetOwnerId(v int64) *DescribeClusterConnectAddrsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeClusterConnectAddrsRequest) SetRegionId(v string) *DescribeClusterConnectAddrsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeClusterConnectAddrsRequest) SetResourceOwnerAccount(v string) *DescribeClusterConnectAddrsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeClusterConnectAddrsRequest) SetResourceOwnerId(v int64) *DescribeClusterConnectAddrsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeClusterConnectAddrsRequest) SetZoneId(v string) *DescribeClusterConnectAddrsRequest {
	s.ZoneId = &v
	return s
}

type DescribeClusterConnectAddrsResponseBody struct {
	DbType              *string                                                     `json:"DbType,omitempty" xml:"DbType,omitempty"`
	IsMultimod          *string                                                     `json:"IsMultimod,omitempty" xml:"IsMultimod,omitempty"`
	NetType             *string                                                     `json:"NetType,omitempty" xml:"NetType,omitempty"`
	RequestId           *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceConnAddrs    *DescribeClusterConnectAddrsResponseBodyServiceConnAddrs    `json:"ServiceConnAddrs,omitempty" xml:"ServiceConnAddrs,omitempty" type:"Struct"`
	SlbConnAddrs        *DescribeClusterConnectAddrsResponseBodySlbConnAddrs        `json:"SlbConnAddrs,omitempty" xml:"SlbConnAddrs,omitempty" type:"Struct"`
	ThriftConn          *DescribeClusterConnectAddrsResponseBodyThriftConn          `json:"ThriftConn,omitempty" xml:"ThriftConn,omitempty" type:"Struct"`
	UiProxyConnAddrInfo *DescribeClusterConnectAddrsResponseBodyUiProxyConnAddrInfo `json:"UiProxyConnAddrInfo,omitempty" xml:"UiProxyConnAddrInfo,omitempty" type:"Struct"`
	VSwitchId           *string                                                     `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId               *string                                                     `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZkConnAddrs         *DescribeClusterConnectAddrsResponseBodyZkConnAddrs         `json:"ZkConnAddrs,omitempty" xml:"ZkConnAddrs,omitempty" type:"Struct"`
}

func (s DescribeClusterConnectAddrsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterConnectAddrsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectAddrsResponseBody) SetDbType(v string) *DescribeClusterConnectAddrsResponseBody {
	s.DbType = &v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBody) SetIsMultimod(v string) *DescribeClusterConnectAddrsResponseBody {
	s.IsMultimod = &v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBody) SetNetType(v string) *DescribeClusterConnectAddrsResponseBody {
	s.NetType = &v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBody) SetRequestId(v string) *DescribeClusterConnectAddrsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBody) SetServiceConnAddrs(v *DescribeClusterConnectAddrsResponseBodyServiceConnAddrs) *DescribeClusterConnectAddrsResponseBody {
	s.ServiceConnAddrs = v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBody) SetSlbConnAddrs(v *DescribeClusterConnectAddrsResponseBodySlbConnAddrs) *DescribeClusterConnectAddrsResponseBody {
	s.SlbConnAddrs = v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBody) SetThriftConn(v *DescribeClusterConnectAddrsResponseBodyThriftConn) *DescribeClusterConnectAddrsResponseBody {
	s.ThriftConn = v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBody) SetUiProxyConnAddrInfo(v *DescribeClusterConnectAddrsResponseBodyUiProxyConnAddrInfo) *DescribeClusterConnectAddrsResponseBody {
	s.UiProxyConnAddrInfo = v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBody) SetVSwitchId(v string) *DescribeClusterConnectAddrsResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBody) SetVpcId(v string) *DescribeClusterConnectAddrsResponseBody {
	s.VpcId = &v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBody) SetZkConnAddrs(v *DescribeClusterConnectAddrsResponseBodyZkConnAddrs) *DescribeClusterConnectAddrsResponseBody {
	s.ZkConnAddrs = v
	return s
}

type DescribeClusterConnectAddrsResponseBodyServiceConnAddrs struct {
	ServiceConnAddr []*DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddr `json:"ServiceConnAddr,omitempty" xml:"ServiceConnAddr,omitempty" type:"Repeated"`
}

func (s DescribeClusterConnectAddrsResponseBodyServiceConnAddrs) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterConnectAddrsResponseBodyServiceConnAddrs) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectAddrsResponseBodyServiceConnAddrs) SetServiceConnAddr(v []*DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddr) *DescribeClusterConnectAddrsResponseBodyServiceConnAddrs {
	s.ServiceConnAddr = v
	return s
}

type DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddr struct {
	ConnAddrInfo *DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo `json:"ConnAddrInfo,omitempty" xml:"ConnAddrInfo,omitempty" type:"Struct"`
	ConnType     *string                                                                             `json:"ConnType,omitempty" xml:"ConnType,omitempty"`
}

func (s DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddr) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddr) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddr) SetConnAddrInfo(v *DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo) *DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddr {
	s.ConnAddrInfo = v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddr) SetConnType(v string) *DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddr {
	s.ConnType = &v
	return s
}

type DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo struct {
	ConnAddr     *string `json:"ConnAddr,omitempty" xml:"ConnAddr,omitempty"`
	ConnAddrPort *string `json:"ConnAddrPort,omitempty" xml:"ConnAddrPort,omitempty"`
	NetType      *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
}

func (s DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo) SetConnAddr(v string) *DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo {
	s.ConnAddr = &v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo) SetConnAddrPort(v string) *DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo {
	s.ConnAddrPort = &v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo) SetNetType(v string) *DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo {
	s.NetType = &v
	return s
}

type DescribeClusterConnectAddrsResponseBodySlbConnAddrs struct {
	SlbConnAddr []*DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddr `json:"SlbConnAddr,omitempty" xml:"SlbConnAddr,omitempty" type:"Repeated"`
}

func (s DescribeClusterConnectAddrsResponseBodySlbConnAddrs) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterConnectAddrsResponseBodySlbConnAddrs) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectAddrsResponseBodySlbConnAddrs) SetSlbConnAddr(v []*DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddr) *DescribeClusterConnectAddrsResponseBodySlbConnAddrs {
	s.SlbConnAddr = v
	return s
}

type DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddr struct {
	ConnAddrInfo *DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo `json:"ConnAddrInfo,omitempty" xml:"ConnAddrInfo,omitempty" type:"Struct"`
	SlbType      *string                                                                     `json:"SlbType,omitempty" xml:"SlbType,omitempty"`
}

func (s DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddr) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddr) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddr) SetConnAddrInfo(v *DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo) *DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddr {
	s.ConnAddrInfo = v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddr) SetSlbType(v string) *DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddr {
	s.SlbType = &v
	return s
}

type DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo struct {
	ConnAddr     *string `json:"ConnAddr,omitempty" xml:"ConnAddr,omitempty"`
	ConnAddrPort *string `json:"ConnAddrPort,omitempty" xml:"ConnAddrPort,omitempty"`
	NetType      *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
}

func (s DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo) SetConnAddr(v string) *DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo {
	s.ConnAddr = &v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo) SetConnAddrPort(v string) *DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo {
	s.ConnAddrPort = &v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo) SetNetType(v string) *DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo {
	s.NetType = &v
	return s
}

type DescribeClusterConnectAddrsResponseBodyThriftConn struct {
	ConnAddr     *string `json:"ConnAddr,omitempty" xml:"ConnAddr,omitempty"`
	ConnAddrPort *string `json:"ConnAddrPort,omitempty" xml:"ConnAddrPort,omitempty"`
	NetType      *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
}

func (s DescribeClusterConnectAddrsResponseBodyThriftConn) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterConnectAddrsResponseBodyThriftConn) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectAddrsResponseBodyThriftConn) SetConnAddr(v string) *DescribeClusterConnectAddrsResponseBodyThriftConn {
	s.ConnAddr = &v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBodyThriftConn) SetConnAddrPort(v string) *DescribeClusterConnectAddrsResponseBodyThriftConn {
	s.ConnAddrPort = &v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBodyThriftConn) SetNetType(v string) *DescribeClusterConnectAddrsResponseBodyThriftConn {
	s.NetType = &v
	return s
}

type DescribeClusterConnectAddrsResponseBodyUiProxyConnAddrInfo struct {
	ConnAddr     *string `json:"ConnAddr,omitempty" xml:"ConnAddr,omitempty"`
	ConnAddrPort *string `json:"ConnAddrPort,omitempty" xml:"ConnAddrPort,omitempty"`
	NetType      *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
}

func (s DescribeClusterConnectAddrsResponseBodyUiProxyConnAddrInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterConnectAddrsResponseBodyUiProxyConnAddrInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectAddrsResponseBodyUiProxyConnAddrInfo) SetConnAddr(v string) *DescribeClusterConnectAddrsResponseBodyUiProxyConnAddrInfo {
	s.ConnAddr = &v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBodyUiProxyConnAddrInfo) SetConnAddrPort(v string) *DescribeClusterConnectAddrsResponseBodyUiProxyConnAddrInfo {
	s.ConnAddrPort = &v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBodyUiProxyConnAddrInfo) SetNetType(v string) *DescribeClusterConnectAddrsResponseBodyUiProxyConnAddrInfo {
	s.NetType = &v
	return s
}

type DescribeClusterConnectAddrsResponseBodyZkConnAddrs struct {
	ZkConnAddr []*DescribeClusterConnectAddrsResponseBodyZkConnAddrsZkConnAddr `json:"ZkConnAddr,omitempty" xml:"ZkConnAddr,omitempty" type:"Repeated"`
}

func (s DescribeClusterConnectAddrsResponseBodyZkConnAddrs) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterConnectAddrsResponseBodyZkConnAddrs) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectAddrsResponseBodyZkConnAddrs) SetZkConnAddr(v []*DescribeClusterConnectAddrsResponseBodyZkConnAddrsZkConnAddr) *DescribeClusterConnectAddrsResponseBodyZkConnAddrs {
	s.ZkConnAddr = v
	return s
}

type DescribeClusterConnectAddrsResponseBodyZkConnAddrsZkConnAddr struct {
	ConnAddr     *string `json:"ConnAddr,omitempty" xml:"ConnAddr,omitempty"`
	ConnAddrPort *string `json:"ConnAddrPort,omitempty" xml:"ConnAddrPort,omitempty"`
	NetType      *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
}

func (s DescribeClusterConnectAddrsResponseBodyZkConnAddrsZkConnAddr) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterConnectAddrsResponseBodyZkConnAddrsZkConnAddr) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectAddrsResponseBodyZkConnAddrsZkConnAddr) SetConnAddr(v string) *DescribeClusterConnectAddrsResponseBodyZkConnAddrsZkConnAddr {
	s.ConnAddr = &v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBodyZkConnAddrsZkConnAddr) SetConnAddrPort(v string) *DescribeClusterConnectAddrsResponseBodyZkConnAddrsZkConnAddr {
	s.ConnAddrPort = &v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBodyZkConnAddrsZkConnAddr) SetNetType(v string) *DescribeClusterConnectAddrsResponseBodyZkConnAddrsZkConnAddr {
	s.NetType = &v
	return s
}

type DescribeClusterConnectAddrsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterConnectAddrsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterConnectAddrsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterConnectAddrsResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectAddrsResponse) SetHeaders(v map[string]*string) *DescribeClusterConnectAddrsResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterConnectAddrsResponse) SetStatusCode(v int32) *DescribeClusterConnectAddrsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterConnectAddrsResponse) SetBody(v *DescribeClusterConnectAddrsResponseBody) *DescribeClusterConnectAddrsResponse {
	s.Body = v
	return s
}

type DescribeClusterListRequest struct {
	ClusterId            *string                          `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName          *string                          `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	DbType               *string                          `json:"DbType,omitempty" xml:"DbType,omitempty"`
	OwnerId              *int64                           `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId             *string                          `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string                          `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                           `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StatusList           []*string                        `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
	Tag                  []*DescribeClusterListRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	ZoneId               *string                          `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeClusterListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterListRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterListRequest) SetClusterId(v string) *DescribeClusterListRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterListRequest) SetClusterName(v string) *DescribeClusterListRequest {
	s.ClusterName = &v
	return s
}

func (s *DescribeClusterListRequest) SetDbType(v string) *DescribeClusterListRequest {
	s.DbType = &v
	return s
}

func (s *DescribeClusterListRequest) SetOwnerId(v int64) *DescribeClusterListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeClusterListRequest) SetPageNumber(v int32) *DescribeClusterListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeClusterListRequest) SetPageSize(v int32) *DescribeClusterListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeClusterListRequest) SetRegionId(v string) *DescribeClusterListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeClusterListRequest) SetResourceOwnerAccount(v string) *DescribeClusterListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeClusterListRequest) SetResourceOwnerId(v int64) *DescribeClusterListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeClusterListRequest) SetStatusList(v []*string) *DescribeClusterListRequest {
	s.StatusList = v
	return s
}

func (s *DescribeClusterListRequest) SetTag(v []*DescribeClusterListRequestTag) *DescribeClusterListRequest {
	s.Tag = v
	return s
}

func (s *DescribeClusterListRequest) SetZoneId(v string) *DescribeClusterListRequest {
	s.ZoneId = &v
	return s
}

type DescribeClusterListRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeClusterListRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterListRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeClusterListRequestTag) SetKey(v string) *DescribeClusterListRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeClusterListRequestTag) SetValue(v string) *DescribeClusterListRequestTag {
	s.Value = &v
	return s
}

type DescribeClusterListResponseBody struct {
	ClusterList      *DescribeClusterListResponseBodyClusterList `json:"ClusterList,omitempty" xml:"ClusterList,omitempty" type:"Struct"`
	PageNumber       *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageRecordCount  *int32                                      `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId        *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalRecordCount *int32                                      `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeClusterListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterListResponseBody) SetClusterList(v *DescribeClusterListResponseBodyClusterList) *DescribeClusterListResponseBody {
	s.ClusterList = v
	return s
}

func (s *DescribeClusterListResponseBody) SetPageNumber(v int32) *DescribeClusterListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeClusterListResponseBody) SetPageRecordCount(v int32) *DescribeClusterListResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeClusterListResponseBody) SetRequestId(v string) *DescribeClusterListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterListResponseBody) SetTotalRecordCount(v int32) *DescribeClusterListResponseBody {
	s.TotalRecordCount = &v
	return s
}

type DescribeClusterListResponseBodyClusterList struct {
	Cluster []*DescribeClusterListResponseBodyClusterListCluster `json:"Cluster,omitempty" xml:"Cluster,omitempty" type:"Repeated"`
}

func (s DescribeClusterListResponseBodyClusterList) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterListResponseBodyClusterList) GoString() string {
	return s.String()
}

func (s *DescribeClusterListResponseBodyClusterList) SetCluster(v []*DescribeClusterListResponseBodyClusterListCluster) *DescribeClusterListResponseBodyClusterList {
	s.Cluster = v
	return s
}

type DescribeClusterListResponseBodyClusterListCluster struct {
	ClusterId            *string                                                `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName          *string                                                `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	ClusterType          *string                                                `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	CoreDiskSize         *string                                                `json:"CoreDiskSize,omitempty" xml:"CoreDiskSize,omitempty"`
	CoreDiskType         *string                                                `json:"CoreDiskType,omitempty" xml:"CoreDiskType,omitempty"`
	CoreInstanceQuantity *int32                                                 `json:"CoreInstanceQuantity,omitempty" xml:"CoreInstanceQuantity,omitempty"`
	CreateTime           *string                                                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DbType               *string                                                `json:"DbType,omitempty" xml:"DbType,omitempty"`
	ExpireTime           *string                                                `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	LockMode             *string                                                `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	MainVersion          *string                                                `json:"MainVersion,omitempty" xml:"MainVersion,omitempty"`
	NetType              *string                                                `json:"NetType,omitempty" xml:"NetType,omitempty"`
	PayType              *string                                                `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Reason               *string                                                `json:"Reason,omitempty" xml:"Reason,omitempty"`
	RegionId             *string                                                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status               *string                                                `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags                 *DescribeClusterListResponseBodyClusterListClusterTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	UserId               *string                                                `json:"UserId,omitempty" xml:"UserId,omitempty"`
	VpcId                *string                                                `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId               *string                                                `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeClusterListResponseBodyClusterListCluster) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterListResponseBodyClusterListCluster) GoString() string {
	return s.String()
}

func (s *DescribeClusterListResponseBodyClusterListCluster) SetClusterId(v string) *DescribeClusterListResponseBodyClusterListCluster {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterListResponseBodyClusterListCluster) SetClusterName(v string) *DescribeClusterListResponseBodyClusterListCluster {
	s.ClusterName = &v
	return s
}

func (s *DescribeClusterListResponseBodyClusterListCluster) SetClusterType(v string) *DescribeClusterListResponseBodyClusterListCluster {
	s.ClusterType = &v
	return s
}

func (s *DescribeClusterListResponseBodyClusterListCluster) SetCoreDiskSize(v string) *DescribeClusterListResponseBodyClusterListCluster {
	s.CoreDiskSize = &v
	return s
}

func (s *DescribeClusterListResponseBodyClusterListCluster) SetCoreDiskType(v string) *DescribeClusterListResponseBodyClusterListCluster {
	s.CoreDiskType = &v
	return s
}

func (s *DescribeClusterListResponseBodyClusterListCluster) SetCoreInstanceQuantity(v int32) *DescribeClusterListResponseBodyClusterListCluster {
	s.CoreInstanceQuantity = &v
	return s
}

func (s *DescribeClusterListResponseBodyClusterListCluster) SetCreateTime(v string) *DescribeClusterListResponseBodyClusterListCluster {
	s.CreateTime = &v
	return s
}

func (s *DescribeClusterListResponseBodyClusterListCluster) SetDbType(v string) *DescribeClusterListResponseBodyClusterListCluster {
	s.DbType = &v
	return s
}

func (s *DescribeClusterListResponseBodyClusterListCluster) SetExpireTime(v string) *DescribeClusterListResponseBodyClusterListCluster {
	s.ExpireTime = &v
	return s
}

func (s *DescribeClusterListResponseBodyClusterListCluster) SetLockMode(v string) *DescribeClusterListResponseBodyClusterListCluster {
	s.LockMode = &v
	return s
}

func (s *DescribeClusterListResponseBodyClusterListCluster) SetMainVersion(v string) *DescribeClusterListResponseBodyClusterListCluster {
	s.MainVersion = &v
	return s
}

func (s *DescribeClusterListResponseBodyClusterListCluster) SetNetType(v string) *DescribeClusterListResponseBodyClusterListCluster {
	s.NetType = &v
	return s
}

func (s *DescribeClusterListResponseBodyClusterListCluster) SetPayType(v string) *DescribeClusterListResponseBodyClusterListCluster {
	s.PayType = &v
	return s
}

func (s *DescribeClusterListResponseBodyClusterListCluster) SetReason(v string) *DescribeClusterListResponseBodyClusterListCluster {
	s.Reason = &v
	return s
}

func (s *DescribeClusterListResponseBodyClusterListCluster) SetRegionId(v string) *DescribeClusterListResponseBodyClusterListCluster {
	s.RegionId = &v
	return s
}

func (s *DescribeClusterListResponseBodyClusterListCluster) SetStatus(v string) *DescribeClusterListResponseBodyClusterListCluster {
	s.Status = &v
	return s
}

func (s *DescribeClusterListResponseBodyClusterListCluster) SetTags(v *DescribeClusterListResponseBodyClusterListClusterTags) *DescribeClusterListResponseBodyClusterListCluster {
	s.Tags = v
	return s
}

func (s *DescribeClusterListResponseBodyClusterListCluster) SetUserId(v string) *DescribeClusterListResponseBodyClusterListCluster {
	s.UserId = &v
	return s
}

func (s *DescribeClusterListResponseBodyClusterListCluster) SetVpcId(v string) *DescribeClusterListResponseBodyClusterListCluster {
	s.VpcId = &v
	return s
}

func (s *DescribeClusterListResponseBodyClusterListCluster) SetZoneId(v string) *DescribeClusterListResponseBodyClusterListCluster {
	s.ZoneId = &v
	return s
}

type DescribeClusterListResponseBodyClusterListClusterTags struct {
	Tag []*DescribeClusterListResponseBodyClusterListClusterTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeClusterListResponseBodyClusterListClusterTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterListResponseBodyClusterListClusterTags) GoString() string {
	return s.String()
}

func (s *DescribeClusterListResponseBodyClusterListClusterTags) SetTag(v []*DescribeClusterListResponseBodyClusterListClusterTagsTag) *DescribeClusterListResponseBodyClusterListClusterTags {
	s.Tag = v
	return s
}

type DescribeClusterListResponseBodyClusterListClusterTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeClusterListResponseBodyClusterListClusterTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterListResponseBodyClusterListClusterTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeClusterListResponseBodyClusterListClusterTagsTag) SetKey(v string) *DescribeClusterListResponseBodyClusterListClusterTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeClusterListResponseBodyClusterListClusterTagsTag) SetValue(v string) *DescribeClusterListResponseBodyClusterListClusterTagsTag {
	s.Value = &v
	return s
}

type DescribeClusterListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterListResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterListResponse) SetHeaders(v map[string]*string) *DescribeClusterListResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterListResponse) SetStatusCode(v int32) *DescribeClusterListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterListResponse) SetBody(v *DescribeClusterListResponseBody) *DescribeClusterListResponse {
	s.Body = v
	return s
}

type DescribeClusterModelRequest struct {
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeClusterModelRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterModelRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterModelRequest) SetClusterId(v string) *DescribeClusterModelRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterModelRequest) SetOwnerId(v int64) *DescribeClusterModelRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeClusterModelRequest) SetRegionId(v string) *DescribeClusterModelRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeClusterModelRequest) SetResourceOwnerAccount(v string) *DescribeClusterModelRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeClusterModelRequest) SetResourceOwnerId(v int64) *DescribeClusterModelRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeClusterModelRequest) SetZoneId(v string) *DescribeClusterModelRequest {
	s.ZoneId = &v
	return s
}

type DescribeClusterModelResponseBody struct {
	AutoRenew            *string                               `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	BackupStatus         *string                               `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	ClusterId            *string                               `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName          *string                               `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	ClusterType          *string                               `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	ColdStorageStatus    *string                               `json:"ColdStorageStatus,omitempty" xml:"ColdStorageStatus,omitempty"`
	CoreDiskQuantity     *int32                                `json:"CoreDiskQuantity,omitempty" xml:"CoreDiskQuantity,omitempty"`
	CoreDiskSize         *string                               `json:"CoreDiskSize,omitempty" xml:"CoreDiskSize,omitempty"`
	CoreDiskType         *string                               `json:"CoreDiskType,omitempty" xml:"CoreDiskType,omitempty"`
	CoreInstanceQuantity *int32                                `json:"CoreInstanceQuantity,omitempty" xml:"CoreInstanceQuantity,omitempty"`
	CoreInstanceType     *string                               `json:"CoreInstanceType,omitempty" xml:"CoreInstanceType,omitempty"`
	CreateTime           *string                               `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DbType               *string                               `json:"DbType,omitempty" xml:"DbType,omitempty"`
	ExpireTime           *string                               `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	HaType               *string                               `json:"HaType,omitempty" xml:"HaType,omitempty"`
	HasUser              *string                               `json:"HasUser,omitempty" xml:"HasUser,omitempty"`
	IsMultimod           *string                               `json:"IsMultimod,omitempty" xml:"IsMultimod,omitempty"`
	LockMode             *string                               `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	MainVersion          *string                               `json:"MainVersion,omitempty" xml:"MainVersion,omitempty"`
	MasterDiskSize       *int32                                `json:"MasterDiskSize,omitempty" xml:"MasterDiskSize,omitempty"`
	MasterDiskType       *string                               `json:"MasterDiskType,omitempty" xml:"MasterDiskType,omitempty"`
	MasterInstanceType   *string                               `json:"MasterInstanceType,omitempty" xml:"MasterInstanceType,omitempty"`
	MinorVersion         *string                               `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	PayType              *string                               `json:"PayType,omitempty" xml:"PayType,omitempty"`
	RegionId             *string                               `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId            *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status               *string                               `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags                 *DescribeClusterModelResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	UpdateStatus         *string                               `json:"UpdateStatus,omitempty" xml:"UpdateStatus,omitempty"`
	ZoneId               *string                               `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeClusterModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterModelResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterModelResponseBody) SetAutoRenew(v string) *DescribeClusterModelResponseBody {
	s.AutoRenew = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetBackupStatus(v string) *DescribeClusterModelResponseBody {
	s.BackupStatus = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetClusterId(v string) *DescribeClusterModelResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetClusterName(v string) *DescribeClusterModelResponseBody {
	s.ClusterName = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetClusterType(v string) *DescribeClusterModelResponseBody {
	s.ClusterType = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetColdStorageStatus(v string) *DescribeClusterModelResponseBody {
	s.ColdStorageStatus = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetCoreDiskQuantity(v int32) *DescribeClusterModelResponseBody {
	s.CoreDiskQuantity = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetCoreDiskSize(v string) *DescribeClusterModelResponseBody {
	s.CoreDiskSize = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetCoreDiskType(v string) *DescribeClusterModelResponseBody {
	s.CoreDiskType = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetCoreInstanceQuantity(v int32) *DescribeClusterModelResponseBody {
	s.CoreInstanceQuantity = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetCoreInstanceType(v string) *DescribeClusterModelResponseBody {
	s.CoreInstanceType = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetCreateTime(v string) *DescribeClusterModelResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetDbType(v string) *DescribeClusterModelResponseBody {
	s.DbType = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetExpireTime(v string) *DescribeClusterModelResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetHaType(v string) *DescribeClusterModelResponseBody {
	s.HaType = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetHasUser(v string) *DescribeClusterModelResponseBody {
	s.HasUser = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetIsMultimod(v string) *DescribeClusterModelResponseBody {
	s.IsMultimod = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetLockMode(v string) *DescribeClusterModelResponseBody {
	s.LockMode = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetMainVersion(v string) *DescribeClusterModelResponseBody {
	s.MainVersion = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetMasterDiskSize(v int32) *DescribeClusterModelResponseBody {
	s.MasterDiskSize = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetMasterDiskType(v string) *DescribeClusterModelResponseBody {
	s.MasterDiskType = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetMasterInstanceType(v string) *DescribeClusterModelResponseBody {
	s.MasterInstanceType = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetMinorVersion(v string) *DescribeClusterModelResponseBody {
	s.MinorVersion = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetPayType(v string) *DescribeClusterModelResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetRegionId(v string) *DescribeClusterModelResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetRequestId(v string) *DescribeClusterModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetStatus(v string) *DescribeClusterModelResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetTags(v *DescribeClusterModelResponseBodyTags) *DescribeClusterModelResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeClusterModelResponseBody) SetUpdateStatus(v string) *DescribeClusterModelResponseBody {
	s.UpdateStatus = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetZoneId(v string) *DescribeClusterModelResponseBody {
	s.ZoneId = &v
	return s
}

type DescribeClusterModelResponseBodyTags struct {
	Tag []*DescribeClusterModelResponseBodyTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeClusterModelResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterModelResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeClusterModelResponseBodyTags) SetTag(v []*DescribeClusterModelResponseBodyTagsTag) *DescribeClusterModelResponseBodyTags {
	s.Tag = v
	return s
}

type DescribeClusterModelResponseBodyTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeClusterModelResponseBodyTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterModelResponseBodyTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeClusterModelResponseBodyTagsTag) SetKey(v string) *DescribeClusterModelResponseBodyTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeClusterModelResponseBodyTagsTag) SetValue(v string) *DescribeClusterModelResponseBodyTagsTag {
	s.Value = &v
	return s
}

type DescribeClusterModelResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterModelResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterModelResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterModelResponse) SetHeaders(v map[string]*string) *DescribeClusterModelResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterModelResponse) SetStatusCode(v int32) *DescribeClusterModelResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterModelResponse) SetBody(v *DescribeClusterModelResponseBody) *DescribeClusterModelResponse {
	s.Body = v
	return s
}

type DescribeClusterWhiteListRequest struct {
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeClusterWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterWhiteListRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterWhiteListRequest) SetClusterId(v string) *DescribeClusterWhiteListRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterWhiteListRequest) SetOwnerId(v int64) *DescribeClusterWhiteListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeClusterWhiteListRequest) SetRegionId(v string) *DescribeClusterWhiteListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeClusterWhiteListRequest) SetResourceOwnerAccount(v string) *DescribeClusterWhiteListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeClusterWhiteListRequest) SetResourceOwnerId(v int64) *DescribeClusterWhiteListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeClusterWhiteListRequest) SetZoneId(v string) *DescribeClusterWhiteListRequest {
	s.ZoneId = &v
	return s
}

type DescribeClusterWhiteListResponseBody struct {
	GroupItems *DescribeClusterWhiteListResponseBodyGroupItems `json:"GroupItems,omitempty" xml:"GroupItems,omitempty" type:"Struct"`
	RequestId  *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClusterWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterWhiteListResponseBody) SetGroupItems(v *DescribeClusterWhiteListResponseBodyGroupItems) *DescribeClusterWhiteListResponseBody {
	s.GroupItems = v
	return s
}

func (s *DescribeClusterWhiteListResponseBody) SetRequestId(v string) *DescribeClusterWhiteListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeClusterWhiteListResponseBodyGroupItems struct {
	WhiteIp []*string `json:"WhiteIp,omitempty" xml:"WhiteIp,omitempty" type:"Repeated"`
}

func (s DescribeClusterWhiteListResponseBodyGroupItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterWhiteListResponseBodyGroupItems) GoString() string {
	return s.String()
}

func (s *DescribeClusterWhiteListResponseBodyGroupItems) SetWhiteIp(v []*string) *DescribeClusterWhiteListResponseBodyGroupItems {
	s.WhiteIp = v
	return s
}

type DescribeClusterWhiteListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterWhiteListResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterWhiteListResponse) SetHeaders(v map[string]*string) *DescribeClusterWhiteListResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterWhiteListResponse) SetStatusCode(v int32) *DescribeClusterWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterWhiteListResponse) SetBody(v *DescribeClusterWhiteListResponseBody) *DescribeClusterWhiteListResponse {
	s.Body = v
	return s
}

type DescribeColdStorageRequest struct {
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeColdStorageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeColdStorageRequest) GoString() string {
	return s.String()
}

func (s *DescribeColdStorageRequest) SetClusterId(v string) *DescribeColdStorageRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeColdStorageRequest) SetOwnerId(v int64) *DescribeColdStorageRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeColdStorageRequest) SetRegionId(v string) *DescribeColdStorageRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeColdStorageRequest) SetResourceOwnerAccount(v string) *DescribeColdStorageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeColdStorageRequest) SetResourceOwnerId(v int64) *DescribeColdStorageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeColdStorageRequest) SetZoneId(v string) *DescribeColdStorageRequest {
	s.ZoneId = &v
	return s
}

type DescribeColdStorageResponseBody struct {
	ClusterId             *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ColdStorageSize       *string `json:"ColdStorageSize,omitempty" xml:"ColdStorageSize,omitempty"`
	ColdStorageUsePercent *string `json:"ColdStorageUsePercent,omitempty" xml:"ColdStorageUsePercent,omitempty"`
	OpenStatus            *string `json:"OpenStatus,omitempty" xml:"OpenStatus,omitempty"`
	PayType               *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	RequestId             *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeColdStorageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeColdStorageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeColdStorageResponseBody) SetClusterId(v string) *DescribeColdStorageResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeColdStorageResponseBody) SetColdStorageSize(v string) *DescribeColdStorageResponseBody {
	s.ColdStorageSize = &v
	return s
}

func (s *DescribeColdStorageResponseBody) SetColdStorageUsePercent(v string) *DescribeColdStorageResponseBody {
	s.ColdStorageUsePercent = &v
	return s
}

func (s *DescribeColdStorageResponseBody) SetOpenStatus(v string) *DescribeColdStorageResponseBody {
	s.OpenStatus = &v
	return s
}

func (s *DescribeColdStorageResponseBody) SetPayType(v string) *DescribeColdStorageResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeColdStorageResponseBody) SetRequestId(v string) *DescribeColdStorageResponseBody {
	s.RequestId = &v
	return s
}

type DescribeColdStorageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeColdStorageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeColdStorageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeColdStorageResponse) GoString() string {
	return s.String()
}

func (s *DescribeColdStorageResponse) SetHeaders(v map[string]*string) *DescribeColdStorageResponse {
	s.Headers = v
	return s
}

func (s *DescribeColdStorageResponse) SetStatusCode(v int32) *DescribeColdStorageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeColdStorageResponse) SetBody(v *DescribeColdStorageResponseBody) *DescribeColdStorageResponse {
	s.Body = v
	return s
}

type DescribeMultiModDbAttributeRequest struct {
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeMultiModDbAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiModDbAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeMultiModDbAttributeRequest) SetClusterId(v string) *DescribeMultiModDbAttributeRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeMultiModDbAttributeRequest) SetOwnerId(v int64) *DescribeMultiModDbAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeMultiModDbAttributeRequest) SetRegionId(v string) *DescribeMultiModDbAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMultiModDbAttributeRequest) SetResourceOwnerAccount(v string) *DescribeMultiModDbAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeMultiModDbAttributeRequest) SetResourceOwnerId(v int64) *DescribeMultiModDbAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeMultiModDbAttributeRequest) SetZoneId(v string) *DescribeMultiModDbAttributeRequest {
	s.ZoneId = &v
	return s
}

type DescribeMultiModDbAttributeResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMultiModDbAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiModDbAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMultiModDbAttributeResponseBody) SetData(v string) *DescribeMultiModDbAttributeResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeMultiModDbAttributeResponseBody) SetRequestId(v string) *DescribeMultiModDbAttributeResponseBody {
	s.RequestId = &v
	return s
}

type DescribeMultiModDbAttributeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMultiModDbAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMultiModDbAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiModDbAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeMultiModDbAttributeResponse) SetHeaders(v map[string]*string) *DescribeMultiModDbAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeMultiModDbAttributeResponse) SetStatusCode(v int32) *DescribeMultiModDbAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMultiModDbAttributeResponse) SetBody(v *DescribeMultiModDbAttributeResponseBody) *DescribeMultiModDbAttributeResponse {
	s.Body = v
	return s
}

type DescribeRdsVSwitchsRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	VpcId                *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRdsVSwitchsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsVSwitchsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRdsVSwitchsRequest) SetOwnerAccount(v string) *DescribeRdsVSwitchsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRdsVSwitchsRequest) SetOwnerId(v int64) *DescribeRdsVSwitchsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRdsVSwitchsRequest) SetRegionId(v string) *DescribeRdsVSwitchsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRdsVSwitchsRequest) SetResourceOwnerAccount(v string) *DescribeRdsVSwitchsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRdsVSwitchsRequest) SetResourceOwnerId(v int64) *DescribeRdsVSwitchsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRdsVSwitchsRequest) SetSecurityToken(v string) *DescribeRdsVSwitchsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeRdsVSwitchsRequest) SetVpcId(v string) *DescribeRdsVSwitchsRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeRdsVSwitchsRequest) SetZoneId(v string) *DescribeRdsVSwitchsRequest {
	s.ZoneId = &v
	return s
}

type DescribeRdsVSwitchsResponseBody struct {
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VSwitches *DescribeRdsVSwitchsResponseBodyVSwitches `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Struct"`
}

func (s DescribeRdsVSwitchsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsVSwitchsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRdsVSwitchsResponseBody) SetRequestId(v string) *DescribeRdsVSwitchsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBody) SetVSwitches(v *DescribeRdsVSwitchsResponseBodyVSwitches) *DescribeRdsVSwitchsResponseBody {
	s.VSwitches = v
	return s
}

type DescribeRdsVSwitchsResponseBodyVSwitches struct {
	VSwitch []*DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch `json:"VSwitch,omitempty" xml:"VSwitch,omitempty" type:"Repeated"`
}

func (s DescribeRdsVSwitchsResponseBodyVSwitches) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsVSwitchsResponseBodyVSwitches) GoString() string {
	return s.String()
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitches) SetVSwitch(v []*DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) *DescribeRdsVSwitchsResponseBodyVSwitches {
	s.VSwitch = v
	return s
}

type DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch struct {
	AliUid      *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	Bid         *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	CidrBlock   *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	GmtCreate   *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	IsDefault   *bool   `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	IzNo        *string `json:"IzNo,omitempty" xml:"IzNo,omitempty"`
	RegionNo    *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VSwitchId   *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
}

func (s DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) GoString() string {
	return s.String()
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetAliUid(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.AliUid = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetBid(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.Bid = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetCidrBlock(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.CidrBlock = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetGmtCreate(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.GmtCreate = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetGmtModified(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.GmtModified = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetIsDefault(v bool) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.IsDefault = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetIzNo(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.IzNo = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetRegionNo(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.RegionNo = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetStatus(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.Status = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetVSwitchId(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.VSwitchId = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetVSwitchName(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.VSwitchName = &v
	return s
}

type DescribeRdsVSwitchsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRdsVSwitchsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRdsVSwitchsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsVSwitchsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRdsVSwitchsResponse) SetHeaders(v map[string]*string) *DescribeRdsVSwitchsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRdsVSwitchsResponse) SetStatusCode(v int32) *DescribeRdsVSwitchsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRdsVSwitchsResponse) SetBody(v *DescribeRdsVSwitchsResponseBody) *DescribeRdsVSwitchsResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetOwnerId(v int64) *DescribeRegionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRegionsRequest) SetRegionId(v string) *DescribeRegionsRequest {
	s.RegionId = &v
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

func (s *DescribeRegionsRequest) SetZoneId(v string) *DescribeRegionsRequest {
	s.ZoneId = &v
	return s
}

type DescribeRegionsResponseBody struct {
	Regions   *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
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
	RegionId *string                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Zones    *DescribeRegionsResponseBodyRegionsRegionZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
}

func (s DescribeRegionsResponseBodyRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionId(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetZones(v *DescribeRegionsResponseBodyRegionsRegionZones) *DescribeRegionsResponseBodyRegionsRegion {
	s.Zones = v
	return s
}

type DescribeRegionsResponseBodyRegionsRegionZones struct {
	Zone []*DescribeRegionsResponseBodyRegionsRegionZonesZone `json:"Zone,omitempty" xml:"Zone,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegionsRegionZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegionZones) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegionZones) SetZone(v []*DescribeRegionsResponseBodyRegionsRegionZonesZone) *DescribeRegionsResponseBodyRegionsRegionZones {
	s.Zone = v
	return s
}

type DescribeRegionsResponseBodyRegionsRegionZonesZone struct {
	VpcEnabled *bool   `json:"VpcEnabled,omitempty" xml:"VpcEnabled,omitempty"`
	ZoneId     *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegionZonesZone) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegionZonesZone) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegionZonesZone) SetVpcEnabled(v bool) *DescribeRegionsResponseBodyRegionsRegionZonesZone {
	s.VpcEnabled = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegionZonesZone) SetZoneId(v string) *DescribeRegionsResponseBodyRegionsRegionZonesZone {
	s.ZoneId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeServerlessInstanceRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId      *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeServerlessInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeServerlessInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeServerlessInstanceRequest) SetClientToken(v string) *DescribeServerlessInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeServerlessInstanceRequest) SetInstanceId(v string) *DescribeServerlessInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeServerlessInstanceRequest) SetRegionId(v string) *DescribeServerlessInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeServerlessInstanceRequest) SetZoneId(v string) *DescribeServerlessInstanceRequest {
	s.ZoneId = &v
	return s
}

type DescribeServerlessInstanceResponseBody struct {
	AutoRenew            *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	ClusterType          *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	CreateTime           *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CuSize               *string `json:"CuSize,omitempty" xml:"CuSize,omitempty"`
	DiskSize             *string `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	ExpireTime           *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	HaType               *string `json:"HaType,omitempty" xml:"HaType,omitempty"`
	HasUser              *string `json:"HasUser,omitempty" xml:"HasUser,omitempty"`
	InnerEndpoint        *string `json:"InnerEndpoint,omitempty" xml:"InnerEndpoint,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName         *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	IsDeletionProtection *string `json:"IsDeletionProtection,omitempty" xml:"IsDeletionProtection,omitempty"`
	LockMode             *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	MainVersion          *string `json:"MainVersion,omitempty" xml:"MainVersion,omitempty"`
	OuterEndpoint        *string `json:"OuterEndpoint,omitempty" xml:"OuterEndpoint,omitempty"`
	PayType              *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ReserverMaxQpsNum    *string `json:"ReserverMaxQpsNum,omitempty" xml:"ReserverMaxQpsNum,omitempty"`
	ReserverMinQpsNum    *string `json:"ReserverMinQpsNum,omitempty" xml:"ReserverMinQpsNum,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateStatus         *string `json:"UpdateStatus,omitempty" xml:"UpdateStatus,omitempty"`
	VSwitchId            *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId                *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeServerlessInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeServerlessInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServerlessInstanceResponseBody) SetAutoRenew(v string) *DescribeServerlessInstanceResponseBody {
	s.AutoRenew = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetClusterType(v string) *DescribeServerlessInstanceResponseBody {
	s.ClusterType = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetCreateTime(v string) *DescribeServerlessInstanceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetCuSize(v string) *DescribeServerlessInstanceResponseBody {
	s.CuSize = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetDiskSize(v string) *DescribeServerlessInstanceResponseBody {
	s.DiskSize = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetExpireTime(v string) *DescribeServerlessInstanceResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetHaType(v string) *DescribeServerlessInstanceResponseBody {
	s.HaType = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetHasUser(v string) *DescribeServerlessInstanceResponseBody {
	s.HasUser = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetInnerEndpoint(v string) *DescribeServerlessInstanceResponseBody {
	s.InnerEndpoint = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetInstanceId(v string) *DescribeServerlessInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetInstanceName(v string) *DescribeServerlessInstanceResponseBody {
	s.InstanceName = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetIsDeletionProtection(v string) *DescribeServerlessInstanceResponseBody {
	s.IsDeletionProtection = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetLockMode(v string) *DescribeServerlessInstanceResponseBody {
	s.LockMode = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetMainVersion(v string) *DescribeServerlessInstanceResponseBody {
	s.MainVersion = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetOuterEndpoint(v string) *DescribeServerlessInstanceResponseBody {
	s.OuterEndpoint = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetPayType(v string) *DescribeServerlessInstanceResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetRegionId(v string) *DescribeServerlessInstanceResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetRequestId(v string) *DescribeServerlessInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetReserverMaxQpsNum(v string) *DescribeServerlessInstanceResponseBody {
	s.ReserverMaxQpsNum = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetReserverMinQpsNum(v string) *DescribeServerlessInstanceResponseBody {
	s.ReserverMinQpsNum = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetStatus(v string) *DescribeServerlessInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetUpdateStatus(v string) *DescribeServerlessInstanceResponseBody {
	s.UpdateStatus = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetVSwitchId(v string) *DescribeServerlessInstanceResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetVpcId(v string) *DescribeServerlessInstanceResponseBody {
	s.VpcId = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetZoneId(v string) *DescribeServerlessInstanceResponseBody {
	s.ZoneId = &v
	return s
}

type DescribeServerlessInstanceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeServerlessInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServerlessInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeServerlessInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeServerlessInstanceResponse) SetHeaders(v map[string]*string) *DescribeServerlessInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeServerlessInstanceResponse) SetStatusCode(v int32) *DescribeServerlessInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServerlessInstanceResponse) SetBody(v *DescribeServerlessInstanceResponseBody) *DescribeServerlessInstanceResponse {
	s.Body = v
	return s
}

type DescribeSubscriptionInitializeProgressRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SubscriptionId       *string `json:"SubscriptionId,omitempty" xml:"SubscriptionId,omitempty"`
}

func (s DescribeSubscriptionInitializeProgressRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionInitializeProgressRequest) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInitializeProgressRequest) SetOwnerId(v int64) *DescribeSubscriptionInitializeProgressRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSubscriptionInitializeProgressRequest) SetPageNumber(v int32) *DescribeSubscriptionInitializeProgressRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSubscriptionInitializeProgressRequest) SetPageSize(v int32) *DescribeSubscriptionInitializeProgressRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSubscriptionInitializeProgressRequest) SetResourceOwnerAccount(v string) *DescribeSubscriptionInitializeProgressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSubscriptionInitializeProgressRequest) SetResourceOwnerId(v int64) *DescribeSubscriptionInitializeProgressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSubscriptionInitializeProgressRequest) SetSubscriptionId(v string) *DescribeSubscriptionInitializeProgressRequest {
	s.SubscriptionId = &v
	return s
}

type DescribeSubscriptionInitializeProgressResponseBody struct {
	Items            []*DescribeSubscriptionInitializeProgressResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PageNumber       *int32                                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageRecordCount  *int32                                                     `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId        *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalRecordCount *int32                                                     `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeSubscriptionInitializeProgressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionInitializeProgressResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInitializeProgressResponseBody) SetItems(v []*DescribeSubscriptionInitializeProgressResponseBodyItems) *DescribeSubscriptionInitializeProgressResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSubscriptionInitializeProgressResponseBody) SetPageNumber(v int32) *DescribeSubscriptionInitializeProgressResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSubscriptionInitializeProgressResponseBody) SetPageRecordCount(v int32) *DescribeSubscriptionInitializeProgressResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSubscriptionInitializeProgressResponseBody) SetRequestId(v string) *DescribeSubscriptionInitializeProgressResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSubscriptionInitializeProgressResponseBody) SetTotalRecordCount(v int32) *DescribeSubscriptionInitializeProgressResponseBody {
	s.TotalRecordCount = &v
	return s
}

type DescribeSubscriptionInitializeProgressResponseBodyItems struct {
	FinishTime     *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	Progress       *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SubscriptionId *string `json:"SubscriptionId,omitempty" xml:"SubscriptionId,omitempty"`
}

func (s DescribeSubscriptionInitializeProgressResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionInitializeProgressResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInitializeProgressResponseBodyItems) SetFinishTime(v string) *DescribeSubscriptionInitializeProgressResponseBodyItems {
	s.FinishTime = &v
	return s
}

func (s *DescribeSubscriptionInitializeProgressResponseBodyItems) SetProgress(v string) *DescribeSubscriptionInitializeProgressResponseBodyItems {
	s.Progress = &v
	return s
}

func (s *DescribeSubscriptionInitializeProgressResponseBodyItems) SetStatus(v string) *DescribeSubscriptionInitializeProgressResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeSubscriptionInitializeProgressResponseBodyItems) SetSubscriptionId(v string) *DescribeSubscriptionInitializeProgressResponseBodyItems {
	s.SubscriptionId = &v
	return s
}

type DescribeSubscriptionInitializeProgressResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSubscriptionInitializeProgressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSubscriptionInitializeProgressResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionInitializeProgressResponse) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInitializeProgressResponse) SetHeaders(v map[string]*string) *DescribeSubscriptionInitializeProgressResponse {
	s.Headers = v
	return s
}

func (s *DescribeSubscriptionInitializeProgressResponse) SetStatusCode(v int32) *DescribeSubscriptionInitializeProgressResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSubscriptionInitializeProgressResponse) SetBody(v *DescribeSubscriptionInitializeProgressResponseBody) *DescribeSubscriptionInitializeProgressResponse {
	s.Body = v
	return s
}

type DescribeSubscriptionPerformanceRequest struct {
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Key                  *string `json:"Key,omitempty" xml:"Key,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SourceInstanceId     *string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	SubscriptionId       *string `json:"SubscriptionId,omitempty" xml:"SubscriptionId,omitempty"`
}

func (s DescribeSubscriptionPerformanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionPerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionPerformanceRequest) SetEndTime(v string) *DescribeSubscriptionPerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSubscriptionPerformanceRequest) SetKey(v string) *DescribeSubscriptionPerformanceRequest {
	s.Key = &v
	return s
}

func (s *DescribeSubscriptionPerformanceRequest) SetOwnerId(v int64) *DescribeSubscriptionPerformanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSubscriptionPerformanceRequest) SetResourceOwnerAccount(v string) *DescribeSubscriptionPerformanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSubscriptionPerformanceRequest) SetResourceOwnerId(v int64) *DescribeSubscriptionPerformanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSubscriptionPerformanceRequest) SetSourceInstanceId(v string) *DescribeSubscriptionPerformanceRequest {
	s.SourceInstanceId = &v
	return s
}

func (s *DescribeSubscriptionPerformanceRequest) SetStartTime(v string) *DescribeSubscriptionPerformanceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSubscriptionPerformanceRequest) SetSubscriptionId(v string) *DescribeSubscriptionPerformanceRequest {
	s.SubscriptionId = &v
	return s
}

type DescribeSubscriptionPerformanceResponseBody struct {
	EndTime         *string                                                     `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PerformanceKeys *DescribeSubscriptionPerformanceResponseBodyPerformanceKeys `json:"PerformanceKeys,omitempty" xml:"PerformanceKeys,omitempty" type:"Struct"`
	ReplicaId       *string                                                     `json:"ReplicaId,omitempty" xml:"ReplicaId,omitempty"`
	RequestId       *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime       *string                                                     `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSubscriptionPerformanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionPerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionPerformanceResponseBody) SetEndTime(v string) *DescribeSubscriptionPerformanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeSubscriptionPerformanceResponseBody) SetPerformanceKeys(v *DescribeSubscriptionPerformanceResponseBodyPerformanceKeys) *DescribeSubscriptionPerformanceResponseBody {
	s.PerformanceKeys = v
	return s
}

func (s *DescribeSubscriptionPerformanceResponseBody) SetReplicaId(v string) *DescribeSubscriptionPerformanceResponseBody {
	s.ReplicaId = &v
	return s
}

func (s *DescribeSubscriptionPerformanceResponseBody) SetRequestId(v string) *DescribeSubscriptionPerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSubscriptionPerformanceResponseBody) SetStartTime(v string) *DescribeSubscriptionPerformanceResponseBody {
	s.StartTime = &v
	return s
}

type DescribeSubscriptionPerformanceResponseBodyPerformanceKeys struct {
	PerformanceKey []*DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKey `json:"PerformanceKey,omitempty" xml:"PerformanceKey,omitempty" type:"Repeated"`
}

func (s DescribeSubscriptionPerformanceResponseBodyPerformanceKeys) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionPerformanceResponseBodyPerformanceKeys) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionPerformanceResponseBodyPerformanceKeys) SetPerformanceKey(v []*DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKey) *DescribeSubscriptionPerformanceResponseBodyPerformanceKeys {
	s.PerformanceKey = v
	return s
}

type DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKey struct {
	Key               *string                                                                                    `json:"Key,omitempty" xml:"Key,omitempty"`
	PerformanceValues *DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValues `json:"PerformanceValues,omitempty" xml:"PerformanceValues,omitempty" type:"Struct"`
	Unit              *string                                                                                    `json:"Unit,omitempty" xml:"Unit,omitempty"`
	ValueFormat       *string                                                                                    `json:"ValueFormat,omitempty" xml:"ValueFormat,omitempty"`
}

func (s DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKey) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKey) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKey) SetKey(v string) *DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKey {
	s.Key = &v
	return s
}

func (s *DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKey) SetPerformanceValues(v *DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValues) *DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKey {
	s.PerformanceValues = v
	return s
}

func (s *DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKey) SetUnit(v string) *DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKey {
	s.Unit = &v
	return s
}

func (s *DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKey) SetValueFormat(v string) *DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKey {
	s.ValueFormat = &v
	return s
}

type DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValues struct {
	PerformanceValue []*DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValuesPerformanceValue `json:"PerformanceValue,omitempty" xml:"PerformanceValue,omitempty" type:"Repeated"`
}

func (s DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValues) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValues) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValues) SetPerformanceValue(v []*DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValuesPerformanceValue) *DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValues {
	s.PerformanceValue = v
	return s
}

type DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValuesPerformanceValue struct {
	Date  *string `json:"Date,omitempty" xml:"Date,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValuesPerformanceValue) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValuesPerformanceValue) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValuesPerformanceValue) SetDate(v string) *DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValuesPerformanceValue {
	s.Date = &v
	return s
}

func (s *DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValuesPerformanceValue) SetValue(v string) *DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValuesPerformanceValue {
	s.Value = &v
	return s
}

type DescribeSubscriptionPerformanceResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSubscriptionPerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSubscriptionPerformanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionPerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionPerformanceResponse) SetHeaders(v map[string]*string) *DescribeSubscriptionPerformanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeSubscriptionPerformanceResponse) SetStatusCode(v int32) *DescribeSubscriptionPerformanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSubscriptionPerformanceResponse) SetBody(v *DescribeSubscriptionPerformanceResponseBody) *DescribeSubscriptionPerformanceResponse {
	s.Body = v
	return s
}

type DescribeSubscriptionPermissionRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeSubscriptionPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionPermissionRequest) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionPermissionRequest) SetOwnerId(v int64) *DescribeSubscriptionPermissionRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSubscriptionPermissionRequest) SetResourceOwnerAccount(v string) *DescribeSubscriptionPermissionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSubscriptionPermissionRequest) SetResourceOwnerId(v int64) *DescribeSubscriptionPermissionRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeSubscriptionPermissionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSubscriptionPermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionPermissionResponseBody) SetRequestId(v string) *DescribeSubscriptionPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSubscriptionPermissionResponseBody) SetStatus(v string) *DescribeSubscriptionPermissionResponseBody {
	s.Status = &v
	return s
}

type DescribeSubscriptionPermissionResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSubscriptionPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSubscriptionPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionPermissionResponse) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionPermissionResponse) SetHeaders(v map[string]*string) *DescribeSubscriptionPermissionResponse {
	s.Headers = v
	return s
}

func (s *DescribeSubscriptionPermissionResponse) SetStatusCode(v int32) *DescribeSubscriptionPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSubscriptionPermissionResponse) SetBody(v *DescribeSubscriptionPermissionResponseBody) *DescribeSubscriptionPermissionResponse {
	s.Body = v
	return s
}

type DescribeSubscriptionsRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SubscriptionId       *string `json:"SubscriptionId,omitempty" xml:"SubscriptionId,omitempty"`
}

func (s DescribeSubscriptionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionsRequest) SetOwnerId(v int64) *DescribeSubscriptionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSubscriptionsRequest) SetPageNumber(v int32) *DescribeSubscriptionsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSubscriptionsRequest) SetPageSize(v int32) *DescribeSubscriptionsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSubscriptionsRequest) SetResourceOwnerAccount(v string) *DescribeSubscriptionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSubscriptionsRequest) SetResourceOwnerId(v int64) *DescribeSubscriptionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSubscriptionsRequest) SetSubscriptionId(v string) *DescribeSubscriptionsRequest {
	s.SubscriptionId = &v
	return s
}

type DescribeSubscriptionsResponseBody struct {
	RequestId     *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Subscriptions []*DescribeSubscriptionsResponseBodySubscriptions `json:"Subscriptions,omitempty" xml:"Subscriptions,omitempty" type:"Repeated"`
}

func (s DescribeSubscriptionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionsResponseBody) SetRequestId(v string) *DescribeSubscriptionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSubscriptionsResponseBody) SetSubscriptions(v []*DescribeSubscriptionsResponseBodySubscriptions) *DescribeSubscriptionsResponseBody {
	s.Subscriptions = v
	return s
}

type DescribeSubscriptionsResponseBodySubscriptions struct {
	DBInstances             []*DescribeSubscriptionsResponseBodySubscriptionsDBInstances `json:"DBInstances,omitempty" xml:"DBInstances,omitempty" type:"Repeated"`
	Mapping                 *string                                                      `json:"Mapping,omitempty" xml:"Mapping,omitempty"`
	SubscriptionDescription *string                                                      `json:"SubscriptionDescription,omitempty" xml:"SubscriptionDescription,omitempty"`
	SubscriptionId          *string                                                      `json:"SubscriptionId,omitempty" xml:"SubscriptionId,omitempty"`
	SubscriptionStatus      *string                                                      `json:"SubscriptionStatus,omitempty" xml:"SubscriptionStatus,omitempty"`
	SubscriptionType        *string                                                      `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
}

func (s DescribeSubscriptionsResponseBodySubscriptions) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionsResponseBodySubscriptions) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionsResponseBodySubscriptions) SetDBInstances(v []*DescribeSubscriptionsResponseBodySubscriptionsDBInstances) *DescribeSubscriptionsResponseBodySubscriptions {
	s.DBInstances = v
	return s
}

func (s *DescribeSubscriptionsResponseBodySubscriptions) SetMapping(v string) *DescribeSubscriptionsResponseBodySubscriptions {
	s.Mapping = &v
	return s
}

func (s *DescribeSubscriptionsResponseBodySubscriptions) SetSubscriptionDescription(v string) *DescribeSubscriptionsResponseBodySubscriptions {
	s.SubscriptionDescription = &v
	return s
}

func (s *DescribeSubscriptionsResponseBodySubscriptions) SetSubscriptionId(v string) *DescribeSubscriptionsResponseBodySubscriptions {
	s.SubscriptionId = &v
	return s
}

func (s *DescribeSubscriptionsResponseBodySubscriptions) SetSubscriptionStatus(v string) *DescribeSubscriptionsResponseBodySubscriptions {
	s.SubscriptionStatus = &v
	return s
}

func (s *DescribeSubscriptionsResponseBodySubscriptions) SetSubscriptionType(v string) *DescribeSubscriptionsResponseBodySubscriptions {
	s.SubscriptionType = &v
	return s
}

type DescribeSubscriptionsResponseBodySubscriptionsDBInstances struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Role         *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s DescribeSubscriptionsResponseBodySubscriptionsDBInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionsResponseBodySubscriptionsDBInstances) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionsResponseBodySubscriptionsDBInstances) SetDBInstanceId(v string) *DescribeSubscriptionsResponseBodySubscriptionsDBInstances {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSubscriptionsResponseBodySubscriptionsDBInstances) SetRegionId(v string) *DescribeSubscriptionsResponseBodySubscriptionsDBInstances {
	s.RegionId = &v
	return s
}

func (s *DescribeSubscriptionsResponseBodySubscriptionsDBInstances) SetRole(v string) *DescribeSubscriptionsResponseBodySubscriptionsDBInstances {
	s.Role = &v
	return s
}

type DescribeSubscriptionsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSubscriptionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSubscriptionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionsResponse) SetHeaders(v map[string]*string) *DescribeSubscriptionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSubscriptionsResponse) SetStatusCode(v int32) *DescribeSubscriptionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSubscriptionsResponse) SetBody(v *DescribeSubscriptionsResponseBody) *DescribeSubscriptionsResponse {
	s.Body = v
	return s
}

type EnableServerlessPublicConnectionRequest struct {
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s EnableServerlessPublicConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableServerlessPublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *EnableServerlessPublicConnectionRequest) SetClientToken(v string) *EnableServerlessPublicConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *EnableServerlessPublicConnectionRequest) SetInstanceId(v string) *EnableServerlessPublicConnectionRequest {
	s.InstanceId = &v
	return s
}

func (s *EnableServerlessPublicConnectionRequest) SetOwnerId(v int64) *EnableServerlessPublicConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *EnableServerlessPublicConnectionRequest) SetRegionId(v string) *EnableServerlessPublicConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *EnableServerlessPublicConnectionRequest) SetResourceOwnerAccount(v string) *EnableServerlessPublicConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *EnableServerlessPublicConnectionRequest) SetResourceOwnerId(v int64) *EnableServerlessPublicConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *EnableServerlessPublicConnectionRequest) SetZoneId(v string) *EnableServerlessPublicConnectionRequest {
	s.ZoneId = &v
	return s
}

type EnableServerlessPublicConnectionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableServerlessPublicConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableServerlessPublicConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *EnableServerlessPublicConnectionResponseBody) SetRequestId(v string) *EnableServerlessPublicConnectionResponseBody {
	s.RequestId = &v
	return s
}

type EnableServerlessPublicConnectionResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableServerlessPublicConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableServerlessPublicConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableServerlessPublicConnectionResponse) GoString() string {
	return s.String()
}

func (s *EnableServerlessPublicConnectionResponse) SetHeaders(v map[string]*string) *EnableServerlessPublicConnectionResponse {
	s.Headers = v
	return s
}

func (s *EnableServerlessPublicConnectionResponse) SetStatusCode(v int32) *EnableServerlessPublicConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableServerlessPublicConnectionResponse) SetBody(v *EnableServerlessPublicConnectionResponseBody) *EnableServerlessPublicConnectionResponse {
	s.Body = v
	return s
}

type GetMultimodeCmsUrlRequest struct {
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetMultimodeCmsUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMultimodeCmsUrlRequest) GoString() string {
	return s.String()
}

func (s *GetMultimodeCmsUrlRequest) SetClusterId(v string) *GetMultimodeCmsUrlRequest {
	s.ClusterId = &v
	return s
}

func (s *GetMultimodeCmsUrlRequest) SetOwnerId(v int64) *GetMultimodeCmsUrlRequest {
	s.OwnerId = &v
	return s
}

func (s *GetMultimodeCmsUrlRequest) SetRegionId(v string) *GetMultimodeCmsUrlRequest {
	s.RegionId = &v
	return s
}

func (s *GetMultimodeCmsUrlRequest) SetResourceOwnerAccount(v string) *GetMultimodeCmsUrlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetMultimodeCmsUrlRequest) SetResourceOwnerId(v int64) *GetMultimodeCmsUrlRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetMultimodeCmsUrlRequest) SetZoneId(v string) *GetMultimodeCmsUrlRequest {
	s.ZoneId = &v
	return s
}

type GetMultimodeCmsUrlResponseBody struct {
	ClusterId      *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	MultimodCmsUrl *string `json:"MultimodCmsUrl,omitempty" xml:"MultimodCmsUrl,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMultimodeCmsUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMultimodeCmsUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetMultimodeCmsUrlResponseBody) SetClusterId(v string) *GetMultimodeCmsUrlResponseBody {
	s.ClusterId = &v
	return s
}

func (s *GetMultimodeCmsUrlResponseBody) SetMultimodCmsUrl(v string) *GetMultimodeCmsUrlResponseBody {
	s.MultimodCmsUrl = &v
	return s
}

func (s *GetMultimodeCmsUrlResponseBody) SetRequestId(v string) *GetMultimodeCmsUrlResponseBody {
	s.RequestId = &v
	return s
}

type GetMultimodeCmsUrlResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMultimodeCmsUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMultimodeCmsUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMultimodeCmsUrlResponse) GoString() string {
	return s.String()
}

func (s *GetMultimodeCmsUrlResponse) SetHeaders(v map[string]*string) *GetMultimodeCmsUrlResponse {
	s.Headers = v
	return s
}

func (s *GetMultimodeCmsUrlResponse) SetStatusCode(v int32) *GetMultimodeCmsUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMultimodeCmsUrlResponse) SetBody(v *GetMultimodeCmsUrlResponseBody) *GetMultimodeCmsUrlResponse {
	s.Body = v
	return s
}

type ListClusterServiceConfigRequest struct {
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListClusterServiceConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClusterServiceConfigRequest) GoString() string {
	return s.String()
}

func (s *ListClusterServiceConfigRequest) SetClusterId(v string) *ListClusterServiceConfigRequest {
	s.ClusterId = &v
	return s
}

func (s *ListClusterServiceConfigRequest) SetOwnerId(v int64) *ListClusterServiceConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *ListClusterServiceConfigRequest) SetPageNumber(v int32) *ListClusterServiceConfigRequest {
	s.PageNumber = &v
	return s
}

func (s *ListClusterServiceConfigRequest) SetPageSize(v int32) *ListClusterServiceConfigRequest {
	s.PageSize = &v
	return s
}

func (s *ListClusterServiceConfigRequest) SetResourceOwnerAccount(v string) *ListClusterServiceConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListClusterServiceConfigRequest) SetResourceOwnerId(v int64) *ListClusterServiceConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

type ListClusterServiceConfigResponseBody struct {
	ClusterId        *string                                         `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName      *string                                         `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	ConfigList       *ListClusterServiceConfigResponseBodyConfigList `json:"ConfigList,omitempty" xml:"ConfigList,omitempty" type:"Struct"`
	PageNumber       *int32                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageRecordCount  *int32                                          `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId        *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalRecordCount *int32                                          `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s ListClusterServiceConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClusterServiceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterServiceConfigResponseBody) SetClusterId(v string) *ListClusterServiceConfigResponseBody {
	s.ClusterId = &v
	return s
}

func (s *ListClusterServiceConfigResponseBody) SetClusterName(v string) *ListClusterServiceConfigResponseBody {
	s.ClusterName = &v
	return s
}

func (s *ListClusterServiceConfigResponseBody) SetConfigList(v *ListClusterServiceConfigResponseBodyConfigList) *ListClusterServiceConfigResponseBody {
	s.ConfigList = v
	return s
}

func (s *ListClusterServiceConfigResponseBody) SetPageNumber(v int32) *ListClusterServiceConfigResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListClusterServiceConfigResponseBody) SetPageRecordCount(v int32) *ListClusterServiceConfigResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *ListClusterServiceConfigResponseBody) SetRequestId(v string) *ListClusterServiceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterServiceConfigResponseBody) SetTotalRecordCount(v int32) *ListClusterServiceConfigResponseBody {
	s.TotalRecordCount = &v
	return s
}

type ListClusterServiceConfigResponseBodyConfigList struct {
	Config []*ListClusterServiceConfigResponseBodyConfigListConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Repeated"`
}

func (s ListClusterServiceConfigResponseBodyConfigList) String() string {
	return tea.Prettify(s)
}

func (s ListClusterServiceConfigResponseBodyConfigList) GoString() string {
	return s.String()
}

func (s *ListClusterServiceConfigResponseBodyConfigList) SetConfig(v []*ListClusterServiceConfigResponseBodyConfigListConfig) *ListClusterServiceConfigResponseBodyConfigList {
	s.Config = v
	return s
}

type ListClusterServiceConfigResponseBodyConfigListConfig struct {
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NeedRestart  *string `json:"NeedRestart,omitempty" xml:"NeedRestart,omitempty"`
	RunningValue *string `json:"RunningValue,omitempty" xml:"RunningValue,omitempty"`
	Unit         *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	ValueRange   *string `json:"ValueRange,omitempty" xml:"ValueRange,omitempty"`
}

func (s ListClusterServiceConfigResponseBodyConfigListConfig) String() string {
	return tea.Prettify(s)
}

func (s ListClusterServiceConfigResponseBodyConfigListConfig) GoString() string {
	return s.String()
}

func (s *ListClusterServiceConfigResponseBodyConfigListConfig) SetDefaultValue(v string) *ListClusterServiceConfigResponseBodyConfigListConfig {
	s.DefaultValue = &v
	return s
}

func (s *ListClusterServiceConfigResponseBodyConfigListConfig) SetDescription(v string) *ListClusterServiceConfigResponseBodyConfigListConfig {
	s.Description = &v
	return s
}

func (s *ListClusterServiceConfigResponseBodyConfigListConfig) SetName(v string) *ListClusterServiceConfigResponseBodyConfigListConfig {
	s.Name = &v
	return s
}

func (s *ListClusterServiceConfigResponseBodyConfigListConfig) SetNeedRestart(v string) *ListClusterServiceConfigResponseBodyConfigListConfig {
	s.NeedRestart = &v
	return s
}

func (s *ListClusterServiceConfigResponseBodyConfigListConfig) SetRunningValue(v string) *ListClusterServiceConfigResponseBodyConfigListConfig {
	s.RunningValue = &v
	return s
}

func (s *ListClusterServiceConfigResponseBodyConfigListConfig) SetUnit(v string) *ListClusterServiceConfigResponseBodyConfigListConfig {
	s.Unit = &v
	return s
}

func (s *ListClusterServiceConfigResponseBodyConfigListConfig) SetValueRange(v string) *ListClusterServiceConfigResponseBodyConfigListConfig {
	s.ValueRange = &v
	return s
}

type ListClusterServiceConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClusterServiceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClusterServiceConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClusterServiceConfigResponse) GoString() string {
	return s.String()
}

func (s *ListClusterServiceConfigResponse) SetHeaders(v map[string]*string) *ListClusterServiceConfigResponse {
	s.Headers = v
	return s
}

func (s *ListClusterServiceConfigResponse) SetStatusCode(v int32) *ListClusterServiceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClusterServiceConfigResponse) SetBody(v *ListClusterServiceConfigResponseBody) *ListClusterServiceConfigResponse {
	s.Body = v
	return s
}

type ListClusterServiceConfigHistoryRequest struct {
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListClusterServiceConfigHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClusterServiceConfigHistoryRequest) GoString() string {
	return s.String()
}

func (s *ListClusterServiceConfigHistoryRequest) SetClusterId(v string) *ListClusterServiceConfigHistoryRequest {
	s.ClusterId = &v
	return s
}

func (s *ListClusterServiceConfigHistoryRequest) SetOwnerId(v int64) *ListClusterServiceConfigHistoryRequest {
	s.OwnerId = &v
	return s
}

func (s *ListClusterServiceConfigHistoryRequest) SetPageNumber(v int32) *ListClusterServiceConfigHistoryRequest {
	s.PageNumber = &v
	return s
}

func (s *ListClusterServiceConfigHistoryRequest) SetPageSize(v int32) *ListClusterServiceConfigHistoryRequest {
	s.PageSize = &v
	return s
}

func (s *ListClusterServiceConfigHistoryRequest) SetRegionId(v string) *ListClusterServiceConfigHistoryRequest {
	s.RegionId = &v
	return s
}

func (s *ListClusterServiceConfigHistoryRequest) SetResourceOwnerAccount(v string) *ListClusterServiceConfigHistoryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListClusterServiceConfigHistoryRequest) SetResourceOwnerId(v int64) *ListClusterServiceConfigHistoryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListClusterServiceConfigHistoryRequest) SetZoneId(v string) *ListClusterServiceConfigHistoryRequest {
	s.ZoneId = &v
	return s
}

type ListClusterServiceConfigHistoryResponseBody struct {
	ConfigHistoryList *ListClusterServiceConfigHistoryResponseBodyConfigHistoryList `json:"ConfigHistoryList,omitempty" xml:"ConfigHistoryList,omitempty" type:"Struct"`
	PageNumber        *int32                                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageRecordCount   *int32                                                        `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId         *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalRecordCount  *int32                                                        `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s ListClusterServiceConfigHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClusterServiceConfigHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterServiceConfigHistoryResponseBody) SetConfigHistoryList(v *ListClusterServiceConfigHistoryResponseBodyConfigHistoryList) *ListClusterServiceConfigHistoryResponseBody {
	s.ConfigHistoryList = v
	return s
}

func (s *ListClusterServiceConfigHistoryResponseBody) SetPageNumber(v int32) *ListClusterServiceConfigHistoryResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListClusterServiceConfigHistoryResponseBody) SetPageRecordCount(v int32) *ListClusterServiceConfigHistoryResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *ListClusterServiceConfigHistoryResponseBody) SetRequestId(v string) *ListClusterServiceConfigHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterServiceConfigHistoryResponseBody) SetTotalRecordCount(v int32) *ListClusterServiceConfigHistoryResponseBody {
	s.TotalRecordCount = &v
	return s
}

type ListClusterServiceConfigHistoryResponseBodyConfigHistoryList struct {
	ConfigHistory []*ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory `json:"ConfigHistory,omitempty" xml:"ConfigHistory,omitempty" type:"Repeated"`
}

func (s ListClusterServiceConfigHistoryResponseBodyConfigHistoryList) String() string {
	return tea.Prettify(s)
}

func (s ListClusterServiceConfigHistoryResponseBodyConfigHistoryList) GoString() string {
	return s.String()
}

func (s *ListClusterServiceConfigHistoryResponseBodyConfigHistoryList) SetConfigHistory(v []*ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory) *ListClusterServiceConfigHistoryResponseBodyConfigHistoryList {
	s.ConfigHistory = v
	return s
}

type ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory struct {
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Effective  *string `json:"Effective,omitempty" xml:"Effective,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NewValue   *string `json:"NewValue,omitempty" xml:"NewValue,omitempty"`
	OldValue   *string `json:"OldValue,omitempty" xml:"OldValue,omitempty"`
}

func (s ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory) String() string {
	return tea.Prettify(s)
}

func (s ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory) GoString() string {
	return s.String()
}

func (s *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory) SetCreateTime(v string) *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory {
	s.CreateTime = &v
	return s
}

func (s *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory) SetEffective(v string) *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory {
	s.Effective = &v
	return s
}

func (s *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory) SetName(v string) *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory {
	s.Name = &v
	return s
}

func (s *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory) SetNewValue(v string) *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory {
	s.NewValue = &v
	return s
}

func (s *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory) SetOldValue(v string) *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory {
	s.OldValue = &v
	return s
}

type ListClusterServiceConfigHistoryResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClusterServiceConfigHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClusterServiceConfigHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClusterServiceConfigHistoryResponse) GoString() string {
	return s.String()
}

func (s *ListClusterServiceConfigHistoryResponse) SetHeaders(v map[string]*string) *ListClusterServiceConfigHistoryResponse {
	s.Headers = v
	return s
}

func (s *ListClusterServiceConfigHistoryResponse) SetStatusCode(v int32) *ListClusterServiceConfigHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClusterServiceConfigHistoryResponse) SetBody(v *ListClusterServiceConfigHistoryResponseBody) *ListClusterServiceConfigHistoryResponse {
	s.Body = v
	return s
}

type ListHbaseInstancesRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	VpcId                *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListHbaseInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHbaseInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListHbaseInstancesRequest) SetOwnerAccount(v string) *ListHbaseInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListHbaseInstancesRequest) SetOwnerId(v int64) *ListHbaseInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListHbaseInstancesRequest) SetResourceOwnerAccount(v string) *ListHbaseInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListHbaseInstancesRequest) SetResourceOwnerId(v int64) *ListHbaseInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListHbaseInstancesRequest) SetSecurityToken(v string) *ListHbaseInstancesRequest {
	s.SecurityToken = &v
	return s
}

func (s *ListHbaseInstancesRequest) SetVpcId(v string) *ListHbaseInstancesRequest {
	s.VpcId = &v
	return s
}

func (s *ListHbaseInstancesRequest) SetZoneId(v string) *ListHbaseInstancesRequest {
	s.ZoneId = &v
	return s
}

type ListHbaseInstancesResponseBody struct {
	Instances *ListHbaseInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListHbaseInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHbaseInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListHbaseInstancesResponseBody) SetInstances(v *ListHbaseInstancesResponseBodyInstances) *ListHbaseInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *ListHbaseInstancesResponseBody) SetRequestId(v string) *ListHbaseInstancesResponseBody {
	s.RequestId = &v
	return s
}

type ListHbaseInstancesResponseBodyInstances struct {
	Instance []*ListHbaseInstancesResponseBodyInstancesInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s ListHbaseInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s ListHbaseInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListHbaseInstancesResponseBodyInstances) SetInstance(v []*ListHbaseInstancesResponseBodyInstancesInstance) *ListHbaseInstancesResponseBodyInstances {
	s.Instance = v
	return s
}

type ListHbaseInstancesResponseBodyInstancesInstance struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	IsDefault    *bool   `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
}

func (s ListHbaseInstancesResponseBodyInstancesInstance) String() string {
	return tea.Prettify(s)
}

func (s ListHbaseInstancesResponseBodyInstancesInstance) GoString() string {
	return s.String()
}

func (s *ListHbaseInstancesResponseBodyInstancesInstance) SetInstanceId(v string) *ListHbaseInstancesResponseBodyInstancesInstance {
	s.InstanceId = &v
	return s
}

func (s *ListHbaseInstancesResponseBodyInstancesInstance) SetInstanceName(v string) *ListHbaseInstancesResponseBodyInstancesInstance {
	s.InstanceName = &v
	return s
}

func (s *ListHbaseInstancesResponseBodyInstancesInstance) SetIsDefault(v bool) *ListHbaseInstancesResponseBodyInstancesInstance {
	s.IsDefault = &v
	return s
}

type ListHbaseInstancesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHbaseInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHbaseInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHbaseInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListHbaseInstancesResponse) SetHeaders(v map[string]*string) *ListHbaseInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListHbaseInstancesResponse) SetStatusCode(v int32) *ListHbaseInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHbaseInstancesResponse) SetBody(v *ListHbaseInstancesResponseBody) *ListHbaseInstancesResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	NextToken            *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount         *string                       `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64                        `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId           []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string                       `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                        `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceType         *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag                  []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetOwnerAccount(v string) *ListTagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTagResourcesRequest) SetOwnerId(v int64) *ListTagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceOwnerAccount(v string) *ListTagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceOwnerId(v int64) *ListTagResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponseBody struct {
	NextToken    *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagResources *ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Struct"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v *ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	TagResource []*ListTagResourcesResponseBodyTagResourcesTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagResource(v []*ListTagResourcesResponseBodyTagResourcesTagResource) *ListTagResourcesResponseBodyTagResources {
	s.TagResource = v
	return s
}

type ListTagResourcesResponseBodyTagResourcesTagResource struct {
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ModifyBackupPolicyRequest struct {
	ClientToken                 *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ClusterId                   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId                     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PreferredBackupEndTimeUTC   *string `json:"PreferredBackupEndTimeUTC,omitempty" xml:"PreferredBackupEndTimeUTC,omitempty"`
	PreferredBackupPeriod       *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	PreferredBackupStartTimeUTC *string `json:"PreferredBackupStartTimeUTC,omitempty" xml:"PreferredBackupStartTimeUTC,omitempty"`
	PreferredBackupTime         *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	RegionId                    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount        *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId             *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId                      *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ModifyBackupPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyRequest) SetClientToken(v string) *ModifyBackupPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetClusterId(v string) *ModifyBackupPolicyRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetOwnerId(v int64) *ModifyBackupPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupEndTimeUTC(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupEndTimeUTC = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupPeriod(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupStartTimeUTC(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupStartTimeUTC = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupTime(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupTime = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetRegionId(v string) *ModifyBackupPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetResourceOwnerAccount(v string) *ModifyBackupPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetResourceOwnerId(v int64) *ModifyBackupPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetZoneId(v string) *ModifyBackupPolicyRequest {
	s.ZoneId = &v
	return s
}

type ModifyBackupPolicyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyBackupPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyResponseBody) SetRequestId(v string) *ModifyBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

type ModifyBackupPolicyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyBackupPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyResponse) SetHeaders(v map[string]*string) *ModifyBackupPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyBackupPolicyResponse) SetStatusCode(v int32) *ModifyBackupPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyBackupPolicyResponse) SetBody(v *ModifyBackupPolicyResponseBody) *ModifyBackupPolicyResponse {
	s.Body = v
	return s
}

type ModifyClusterNameRequest struct {
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName          *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ModifyClusterNameRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterNameRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterNameRequest) SetClusterId(v string) *ModifyClusterNameRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyClusterNameRequest) SetClusterName(v string) *ModifyClusterNameRequest {
	s.ClusterName = &v
	return s
}

func (s *ModifyClusterNameRequest) SetOwnerId(v int64) *ModifyClusterNameRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyClusterNameRequest) SetRegionId(v string) *ModifyClusterNameRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyClusterNameRequest) SetResourceOwnerAccount(v string) *ModifyClusterNameRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyClusterNameRequest) SetResourceOwnerId(v int64) *ModifyClusterNameRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyClusterNameRequest) SetZoneId(v string) *ModifyClusterNameRequest {
	s.ZoneId = &v
	return s
}

type ModifyClusterNameResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyClusterNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterNameResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterNameResponseBody) SetRequestId(v string) *ModifyClusterNameResponseBody {
	s.RequestId = &v
	return s
}

type ModifyClusterNameResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyClusterNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyClusterNameResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterNameResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterNameResponse) SetHeaders(v map[string]*string) *ModifyClusterNameResponse {
	s.Headers = v
	return s
}

func (s *ModifyClusterNameResponse) SetStatusCode(v int32) *ModifyClusterNameResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyClusterNameResponse) SetBody(v *ModifyClusterNameResponseBody) *ModifyClusterNameResponse {
	s.Body = v
	return s
}

type ModifyClusterNetTypeRequest struct {
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	NetType              *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	VSwitchId            *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId                *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ModifyClusterNetTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterNetTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterNetTypeRequest) SetClusterId(v string) *ModifyClusterNetTypeRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyClusterNetTypeRequest) SetNetType(v string) *ModifyClusterNetTypeRequest {
	s.NetType = &v
	return s
}

func (s *ModifyClusterNetTypeRequest) SetOwnerId(v int64) *ModifyClusterNetTypeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyClusterNetTypeRequest) SetRegionId(v string) *ModifyClusterNetTypeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyClusterNetTypeRequest) SetResourceOwnerAccount(v string) *ModifyClusterNetTypeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyClusterNetTypeRequest) SetResourceOwnerId(v int64) *ModifyClusterNetTypeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyClusterNetTypeRequest) SetVSwitchId(v string) *ModifyClusterNetTypeRequest {
	s.VSwitchId = &v
	return s
}

func (s *ModifyClusterNetTypeRequest) SetVpcId(v string) *ModifyClusterNetTypeRequest {
	s.VpcId = &v
	return s
}

func (s *ModifyClusterNetTypeRequest) SetZoneId(v string) *ModifyClusterNetTypeRequest {
	s.ZoneId = &v
	return s
}

type ModifyClusterNetTypeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyClusterNetTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterNetTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterNetTypeResponseBody) SetRequestId(v string) *ModifyClusterNetTypeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyClusterNetTypeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyClusterNetTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyClusterNetTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterNetTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterNetTypeResponse) SetHeaders(v map[string]*string) *ModifyClusterNetTypeResponse {
	s.Headers = v
	return s
}

func (s *ModifyClusterNetTypeResponse) SetStatusCode(v int32) *ModifyClusterNetTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyClusterNetTypeResponse) SetBody(v *ModifyClusterNetTypeResponseBody) *ModifyClusterNetTypeResponse {
	s.Body = v
	return s
}

type ModifyClusterSecurityIpListRequest struct {
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityIpList       *string `json:"SecurityIpList,omitempty" xml:"SecurityIpList,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ModifyClusterSecurityIpListRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterSecurityIpListRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterSecurityIpListRequest) SetClusterId(v string) *ModifyClusterSecurityIpListRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyClusterSecurityIpListRequest) SetOwnerId(v int64) *ModifyClusterSecurityIpListRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyClusterSecurityIpListRequest) SetRegionId(v string) *ModifyClusterSecurityIpListRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyClusterSecurityIpListRequest) SetResourceOwnerAccount(v string) *ModifyClusterSecurityIpListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyClusterSecurityIpListRequest) SetResourceOwnerId(v int64) *ModifyClusterSecurityIpListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyClusterSecurityIpListRequest) SetSecurityIpList(v string) *ModifyClusterSecurityIpListRequest {
	s.SecurityIpList = &v
	return s
}

func (s *ModifyClusterSecurityIpListRequest) SetZoneId(v string) *ModifyClusterSecurityIpListRequest {
	s.ZoneId = &v
	return s
}

type ModifyClusterSecurityIpListResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyClusterSecurityIpListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterSecurityIpListResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterSecurityIpListResponseBody) SetRequestId(v string) *ModifyClusterSecurityIpListResponseBody {
	s.RequestId = &v
	return s
}

type ModifyClusterSecurityIpListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyClusterSecurityIpListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyClusterSecurityIpListResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterSecurityIpListResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterSecurityIpListResponse) SetHeaders(v map[string]*string) *ModifyClusterSecurityIpListResponse {
	s.Headers = v
	return s
}

func (s *ModifyClusterSecurityIpListResponse) SetStatusCode(v int32) *ModifyClusterSecurityIpListResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyClusterSecurityIpListResponse) SetBody(v *ModifyClusterSecurityIpListResponseBody) *ModifyClusterSecurityIpListResponse {
	s.Body = v
	return s
}

type ModifyClusterServiceConfigRequest struct {
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Parameters           *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Restart              *bool   `json:"Restart,omitempty" xml:"Restart,omitempty"`
	Value                *string `json:"Value,omitempty" xml:"Value,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ModifyClusterServiceConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterServiceConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterServiceConfigRequest) SetClusterId(v string) *ModifyClusterServiceConfigRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyClusterServiceConfigRequest) SetName(v string) *ModifyClusterServiceConfigRequest {
	s.Name = &v
	return s
}

func (s *ModifyClusterServiceConfigRequest) SetOwnerId(v int64) *ModifyClusterServiceConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyClusterServiceConfigRequest) SetParameters(v string) *ModifyClusterServiceConfigRequest {
	s.Parameters = &v
	return s
}

func (s *ModifyClusterServiceConfigRequest) SetRegionId(v string) *ModifyClusterServiceConfigRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyClusterServiceConfigRequest) SetResourceOwnerAccount(v string) *ModifyClusterServiceConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyClusterServiceConfigRequest) SetResourceOwnerId(v int64) *ModifyClusterServiceConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyClusterServiceConfigRequest) SetRestart(v bool) *ModifyClusterServiceConfigRequest {
	s.Restart = &v
	return s
}

func (s *ModifyClusterServiceConfigRequest) SetValue(v string) *ModifyClusterServiceConfigRequest {
	s.Value = &v
	return s
}

func (s *ModifyClusterServiceConfigRequest) SetZoneId(v string) *ModifyClusterServiceConfigRequest {
	s.ZoneId = &v
	return s
}

type ModifyClusterServiceConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyClusterServiceConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterServiceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterServiceConfigResponseBody) SetRequestId(v string) *ModifyClusterServiceConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyClusterServiceConfigResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyClusterServiceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyClusterServiceConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterServiceConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterServiceConfigResponse) SetHeaders(v map[string]*string) *ModifyClusterServiceConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyClusterServiceConfigResponse) SetStatusCode(v int32) *ModifyClusterServiceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyClusterServiceConfigResponse) SetBody(v *ModifyClusterServiceConfigResponseBody) *ModifyClusterServiceConfigResponse {
	s.Body = v
	return s
}

type ModifyHasRootPasswordRequest struct {
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	HasPassword          *string `json:"HasPassword,omitempty" xml:"HasPassword,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ModifyHasRootPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyHasRootPasswordRequest) GoString() string {
	return s.String()
}

func (s *ModifyHasRootPasswordRequest) SetClusterId(v string) *ModifyHasRootPasswordRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyHasRootPasswordRequest) SetHasPassword(v string) *ModifyHasRootPasswordRequest {
	s.HasPassword = &v
	return s
}

func (s *ModifyHasRootPasswordRequest) SetOwnerId(v int64) *ModifyHasRootPasswordRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyHasRootPasswordRequest) SetRegionId(v string) *ModifyHasRootPasswordRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyHasRootPasswordRequest) SetResourceOwnerAccount(v string) *ModifyHasRootPasswordRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyHasRootPasswordRequest) SetResourceOwnerId(v int64) *ModifyHasRootPasswordRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyHasRootPasswordRequest) SetZoneId(v string) *ModifyHasRootPasswordRequest {
	s.ZoneId = &v
	return s
}

type ModifyHasRootPasswordResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyHasRootPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyHasRootPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHasRootPasswordResponseBody) SetRequestId(v string) *ModifyHasRootPasswordResponseBody {
	s.RequestId = &v
	return s
}

type ModifyHasRootPasswordResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyHasRootPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyHasRootPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyHasRootPasswordResponse) GoString() string {
	return s.String()
}

func (s *ModifyHasRootPasswordResponse) SetHeaders(v map[string]*string) *ModifyHasRootPasswordResponse {
	s.Headers = v
	return s
}

func (s *ModifyHasRootPasswordResponse) SetStatusCode(v int32) *ModifyHasRootPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyHasRootPasswordResponse) SetBody(v *ModifyHasRootPasswordResponseBody) *ModifyHasRootPasswordResponse {
	s.Body = v
	return s
}

type ModifyRestartClusterRequest struct {
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Components           *string `json:"Components,omitempty" xml:"Components,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ModifyRestartClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyRestartClusterRequest) GoString() string {
	return s.String()
}

func (s *ModifyRestartClusterRequest) SetClusterId(v string) *ModifyRestartClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyRestartClusterRequest) SetComponents(v string) *ModifyRestartClusterRequest {
	s.Components = &v
	return s
}

func (s *ModifyRestartClusterRequest) SetOwnerId(v int64) *ModifyRestartClusterRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyRestartClusterRequest) SetRegionId(v string) *ModifyRestartClusterRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyRestartClusterRequest) SetResourceOwnerAccount(v string) *ModifyRestartClusterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyRestartClusterRequest) SetResourceOwnerId(v int64) *ModifyRestartClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyRestartClusterRequest) SetZoneId(v string) *ModifyRestartClusterRequest {
	s.ZoneId = &v
	return s
}

type ModifyRestartClusterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyRestartClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyRestartClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRestartClusterResponseBody) SetRequestId(v string) *ModifyRestartClusterResponseBody {
	s.RequestId = &v
	return s
}

type ModifyRestartClusterResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyRestartClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyRestartClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyRestartClusterResponse) GoString() string {
	return s.String()
}

func (s *ModifyRestartClusterResponse) SetHeaders(v map[string]*string) *ModifyRestartClusterResponse {
	s.Headers = v
	return s
}

func (s *ModifyRestartClusterResponse) SetStatusCode(v int32) *ModifyRestartClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRestartClusterResponse) SetBody(v *ModifyRestartClusterResponseBody) *ModifyRestartClusterResponse {
	s.Body = v
	return s
}

type ModifyRollbackHasForHbaseRequest struct {
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ModifyRollbackHasForHbaseRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyRollbackHasForHbaseRequest) GoString() string {
	return s.String()
}

func (s *ModifyRollbackHasForHbaseRequest) SetClientToken(v string) *ModifyRollbackHasForHbaseRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyRollbackHasForHbaseRequest) SetClusterId(v string) *ModifyRollbackHasForHbaseRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyRollbackHasForHbaseRequest) SetOwnerId(v int64) *ModifyRollbackHasForHbaseRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyRollbackHasForHbaseRequest) SetRegionId(v string) *ModifyRollbackHasForHbaseRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyRollbackHasForHbaseRequest) SetResourceOwnerAccount(v string) *ModifyRollbackHasForHbaseRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyRollbackHasForHbaseRequest) SetResourceOwnerId(v int64) *ModifyRollbackHasForHbaseRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyRollbackHasForHbaseRequest) SetZoneId(v string) *ModifyRollbackHasForHbaseRequest {
	s.ZoneId = &v
	return s
}

type ModifyRollbackHasForHbaseResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyRollbackHasForHbaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyRollbackHasForHbaseResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRollbackHasForHbaseResponseBody) SetRequestId(v string) *ModifyRollbackHasForHbaseResponseBody {
	s.RequestId = &v
	return s
}

type ModifyRollbackHasForHbaseResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyRollbackHasForHbaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyRollbackHasForHbaseResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyRollbackHasForHbaseResponse) GoString() string {
	return s.String()
}

func (s *ModifyRollbackHasForHbaseResponse) SetHeaders(v map[string]*string) *ModifyRollbackHasForHbaseResponse {
	s.Headers = v
	return s
}

func (s *ModifyRollbackHasForHbaseResponse) SetStatusCode(v int32) *ModifyRollbackHasForHbaseResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRollbackHasForHbaseResponse) SetBody(v *ModifyRollbackHasForHbaseResponseBody) *ModifyRollbackHasForHbaseResponse {
	s.Body = v
	return s
}

type ModifySubscriptionDescriptionRequest struct {
	OwnerId                 *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount    *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId         *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SubscriptionDescription *string `json:"SubscriptionDescription,omitempty" xml:"SubscriptionDescription,omitempty"`
	SubscriptionId          *string `json:"SubscriptionId,omitempty" xml:"SubscriptionId,omitempty"`
}

func (s ModifySubscriptionDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySubscriptionDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifySubscriptionDescriptionRequest) SetOwnerId(v int64) *ModifySubscriptionDescriptionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySubscriptionDescriptionRequest) SetResourceOwnerAccount(v string) *ModifySubscriptionDescriptionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySubscriptionDescriptionRequest) SetResourceOwnerId(v int64) *ModifySubscriptionDescriptionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySubscriptionDescriptionRequest) SetSubscriptionDescription(v string) *ModifySubscriptionDescriptionRequest {
	s.SubscriptionDescription = &v
	return s
}

func (s *ModifySubscriptionDescriptionRequest) SetSubscriptionId(v string) *ModifySubscriptionDescriptionRequest {
	s.SubscriptionId = &v
	return s
}

type ModifySubscriptionDescriptionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySubscriptionDescriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySubscriptionDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySubscriptionDescriptionResponseBody) SetRequestId(v string) *ModifySubscriptionDescriptionResponseBody {
	s.RequestId = &v
	return s
}

type ModifySubscriptionDescriptionResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySubscriptionDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySubscriptionDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifySubscriptionDescriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifySubscriptionDescriptionResponse) SetHeaders(v map[string]*string) *ModifySubscriptionDescriptionResponse {
	s.Headers = v
	return s
}

func (s *ModifySubscriptionDescriptionResponse) SetStatusCode(v int32) *ModifySubscriptionDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySubscriptionDescriptionResponse) SetBody(v *ModifySubscriptionDescriptionResponseBody) *ModifySubscriptionDescriptionResponse {
	s.Body = v
	return s
}

type ModifySubscriptionMappingRequest struct {
	Mapping              *string `json:"Mapping,omitempty" xml:"Mapping,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SubscriptionId       *string `json:"SubscriptionId,omitempty" xml:"SubscriptionId,omitempty"`
}

func (s ModifySubscriptionMappingRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySubscriptionMappingRequest) GoString() string {
	return s.String()
}

func (s *ModifySubscriptionMappingRequest) SetMapping(v string) *ModifySubscriptionMappingRequest {
	s.Mapping = &v
	return s
}

func (s *ModifySubscriptionMappingRequest) SetOwnerId(v int64) *ModifySubscriptionMappingRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySubscriptionMappingRequest) SetResourceOwnerAccount(v string) *ModifySubscriptionMappingRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySubscriptionMappingRequest) SetResourceOwnerId(v int64) *ModifySubscriptionMappingRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySubscriptionMappingRequest) SetSubscriptionId(v string) *ModifySubscriptionMappingRequest {
	s.SubscriptionId = &v
	return s
}

type ModifySubscriptionMappingResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySubscriptionMappingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySubscriptionMappingResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySubscriptionMappingResponseBody) SetRequestId(v string) *ModifySubscriptionMappingResponseBody {
	s.RequestId = &v
	return s
}

type ModifySubscriptionMappingResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySubscriptionMappingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySubscriptionMappingResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifySubscriptionMappingResponse) GoString() string {
	return s.String()
}

func (s *ModifySubscriptionMappingResponse) SetHeaders(v map[string]*string) *ModifySubscriptionMappingResponse {
	s.Headers = v
	return s
}

func (s *ModifySubscriptionMappingResponse) SetStatusCode(v int32) *ModifySubscriptionMappingResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySubscriptionMappingResponse) SetBody(v *ModifySubscriptionMappingResponseBody) *ModifySubscriptionMappingResponse {
	s.Body = v
	return s
}

type ModifySubscriptionPermissionRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Status               *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifySubscriptionPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySubscriptionPermissionRequest) GoString() string {
	return s.String()
}

func (s *ModifySubscriptionPermissionRequest) SetOwnerId(v int64) *ModifySubscriptionPermissionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySubscriptionPermissionRequest) SetResourceOwnerAccount(v string) *ModifySubscriptionPermissionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySubscriptionPermissionRequest) SetResourceOwnerId(v int64) *ModifySubscriptionPermissionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySubscriptionPermissionRequest) SetStatus(v int32) *ModifySubscriptionPermissionRequest {
	s.Status = &v
	return s
}

type ModifySubscriptionPermissionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifySubscriptionPermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySubscriptionPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySubscriptionPermissionResponseBody) SetRequestId(v string) *ModifySubscriptionPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySubscriptionPermissionResponseBody) SetStatus(v string) *ModifySubscriptionPermissionResponseBody {
	s.Status = &v
	return s
}

type ModifySubscriptionPermissionResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySubscriptionPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySubscriptionPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifySubscriptionPermissionResponse) GoString() string {
	return s.String()
}

func (s *ModifySubscriptionPermissionResponse) SetHeaders(v map[string]*string) *ModifySubscriptionPermissionResponse {
	s.Headers = v
	return s
}

func (s *ModifySubscriptionPermissionResponse) SetStatusCode(v int32) *ModifySubscriptionPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySubscriptionPermissionResponse) SetBody(v *ModifySubscriptionPermissionResponseBody) *ModifySubscriptionPermissionResponse {
	s.Body = v
	return s
}

type ModifyUIProxyAccountPasswordRequest struct {
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountPassword      *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ModifyUIProxyAccountPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyUIProxyAccountPasswordRequest) GoString() string {
	return s.String()
}

func (s *ModifyUIProxyAccountPasswordRequest) SetAccountName(v string) *ModifyUIProxyAccountPasswordRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyUIProxyAccountPasswordRequest) SetAccountPassword(v string) *ModifyUIProxyAccountPasswordRequest {
	s.AccountPassword = &v
	return s
}

func (s *ModifyUIProxyAccountPasswordRequest) SetClusterId(v string) *ModifyUIProxyAccountPasswordRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyUIProxyAccountPasswordRequest) SetOwnerId(v int64) *ModifyUIProxyAccountPasswordRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyUIProxyAccountPasswordRequest) SetRegionId(v string) *ModifyUIProxyAccountPasswordRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyUIProxyAccountPasswordRequest) SetResourceOwnerAccount(v string) *ModifyUIProxyAccountPasswordRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyUIProxyAccountPasswordRequest) SetResourceOwnerId(v int64) *ModifyUIProxyAccountPasswordRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyUIProxyAccountPasswordRequest) SetZoneId(v string) *ModifyUIProxyAccountPasswordRequest {
	s.ZoneId = &v
	return s
}

type ModifyUIProxyAccountPasswordResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyUIProxyAccountPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyUIProxyAccountPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUIProxyAccountPasswordResponseBody) SetRequestId(v string) *ModifyUIProxyAccountPasswordResponseBody {
	s.RequestId = &v
	return s
}

type ModifyUIProxyAccountPasswordResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyUIProxyAccountPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyUIProxyAccountPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyUIProxyAccountPasswordResponse) GoString() string {
	return s.String()
}

func (s *ModifyUIProxyAccountPasswordResponse) SetHeaders(v map[string]*string) *ModifyUIProxyAccountPasswordResponse {
	s.Headers = v
	return s
}

func (s *ModifyUIProxyAccountPasswordResponse) SetStatusCode(v int32) *ModifyUIProxyAccountPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyUIProxyAccountPasswordResponse) SetBody(v *ModifyUIProxyAccountPasswordResponseBody) *ModifyUIProxyAccountPasswordResponse {
	s.Body = v
	return s
}

type ModifyUpgradeToHasForHbaseRequest struct {
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	HasPassword          *string `json:"HasPassword,omitempty" xml:"HasPassword,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ModifyUpgradeToHasForHbaseRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyUpgradeToHasForHbaseRequest) GoString() string {
	return s.String()
}

func (s *ModifyUpgradeToHasForHbaseRequest) SetClientToken(v string) *ModifyUpgradeToHasForHbaseRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyUpgradeToHasForHbaseRequest) SetClusterId(v string) *ModifyUpgradeToHasForHbaseRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyUpgradeToHasForHbaseRequest) SetHasPassword(v string) *ModifyUpgradeToHasForHbaseRequest {
	s.HasPassword = &v
	return s
}

func (s *ModifyUpgradeToHasForHbaseRequest) SetOwnerId(v int64) *ModifyUpgradeToHasForHbaseRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyUpgradeToHasForHbaseRequest) SetRegionId(v string) *ModifyUpgradeToHasForHbaseRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyUpgradeToHasForHbaseRequest) SetResourceOwnerAccount(v string) *ModifyUpgradeToHasForHbaseRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyUpgradeToHasForHbaseRequest) SetResourceOwnerId(v int64) *ModifyUpgradeToHasForHbaseRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyUpgradeToHasForHbaseRequest) SetZoneId(v string) *ModifyUpgradeToHasForHbaseRequest {
	s.ZoneId = &v
	return s
}

type ModifyUpgradeToHasForHbaseResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyUpgradeToHasForHbaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyUpgradeToHasForHbaseResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUpgradeToHasForHbaseResponseBody) SetRequestId(v string) *ModifyUpgradeToHasForHbaseResponseBody {
	s.RequestId = &v
	return s
}

type ModifyUpgradeToHasForHbaseResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyUpgradeToHasForHbaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyUpgradeToHasForHbaseResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyUpgradeToHasForHbaseResponse) GoString() string {
	return s.String()
}

func (s *ModifyUpgradeToHasForHbaseResponse) SetHeaders(v map[string]*string) *ModifyUpgradeToHasForHbaseResponse {
	s.Headers = v
	return s
}

func (s *ModifyUpgradeToHasForHbaseResponse) SetStatusCode(v int32) *ModifyUpgradeToHasForHbaseResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyUpgradeToHasForHbaseResponse) SetBody(v *ModifyUpgradeToHasForHbaseResponseBody) *ModifyUpgradeToHasForHbaseResponse {
	s.Body = v
	return s
}

type MultimodAddComponentsRequest struct {
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Components           *string `json:"Components,omitempty" xml:"Components,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s MultimodAddComponentsRequest) String() string {
	return tea.Prettify(s)
}

func (s MultimodAddComponentsRequest) GoString() string {
	return s.String()
}

func (s *MultimodAddComponentsRequest) SetClusterId(v string) *MultimodAddComponentsRequest {
	s.ClusterId = &v
	return s
}

func (s *MultimodAddComponentsRequest) SetComponents(v string) *MultimodAddComponentsRequest {
	s.Components = &v
	return s
}

func (s *MultimodAddComponentsRequest) SetOwnerId(v int64) *MultimodAddComponentsRequest {
	s.OwnerId = &v
	return s
}

func (s *MultimodAddComponentsRequest) SetRegionId(v string) *MultimodAddComponentsRequest {
	s.RegionId = &v
	return s
}

func (s *MultimodAddComponentsRequest) SetResourceOwnerAccount(v string) *MultimodAddComponentsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *MultimodAddComponentsRequest) SetResourceOwnerId(v int64) *MultimodAddComponentsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *MultimodAddComponentsRequest) SetZoneId(v string) *MultimodAddComponentsRequest {
	s.ZoneId = &v
	return s
}

type MultimodAddComponentsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MultimodAddComponentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MultimodAddComponentsResponseBody) GoString() string {
	return s.String()
}

func (s *MultimodAddComponentsResponseBody) SetRequestId(v string) *MultimodAddComponentsResponseBody {
	s.RequestId = &v
	return s
}

type MultimodAddComponentsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MultimodAddComponentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MultimodAddComponentsResponse) String() string {
	return tea.Prettify(s)
}

func (s MultimodAddComponentsResponse) GoString() string {
	return s.String()
}

func (s *MultimodAddComponentsResponse) SetHeaders(v map[string]*string) *MultimodAddComponentsResponse {
	s.Headers = v
	return s
}

func (s *MultimodAddComponentsResponse) SetStatusCode(v int32) *MultimodAddComponentsResponse {
	s.StatusCode = &v
	return s
}

func (s *MultimodAddComponentsResponse) SetBody(v *MultimodAddComponentsResponseBody) *MultimodAddComponentsResponse {
	s.Body = v
	return s
}

type OpenBackupRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId    *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s OpenBackupRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenBackupRequest) GoString() string {
	return s.String()
}

func (s *OpenBackupRequest) SetClusterId(v string) *OpenBackupRequest {
	s.ClusterId = &v
	return s
}

func (s *OpenBackupRequest) SetRegionId(v string) *OpenBackupRequest {
	s.RegionId = &v
	return s
}

func (s *OpenBackupRequest) SetZoneId(v string) *OpenBackupRequest {
	s.ZoneId = &v
	return s
}

type OpenBackupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenBackupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenBackupResponseBody) GoString() string {
	return s.String()
}

func (s *OpenBackupResponseBody) SetRequestId(v string) *OpenBackupResponseBody {
	s.RequestId = &v
	return s
}

type OpenBackupResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenBackupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenBackupResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenBackupResponse) GoString() string {
	return s.String()
}

func (s *OpenBackupResponse) SetHeaders(v map[string]*string) *OpenBackupResponse {
	s.Headers = v
	return s
}

func (s *OpenBackupResponse) SetStatusCode(v int32) *OpenBackupResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenBackupResponse) SetBody(v *OpenBackupResponseBody) *OpenBackupResponse {
	s.Body = v
	return s
}

type QueryHBaseHaDBRequest struct {
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s QueryHBaseHaDBRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryHBaseHaDBRequest) GoString() string {
	return s.String()
}

func (s *QueryHBaseHaDBRequest) SetClusterId(v string) *QueryHBaseHaDBRequest {
	s.ClusterId = &v
	return s
}

func (s *QueryHBaseHaDBRequest) SetOwnerId(v int64) *QueryHBaseHaDBRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryHBaseHaDBRequest) SetRegionId(v string) *QueryHBaseHaDBRequest {
	s.RegionId = &v
	return s
}

func (s *QueryHBaseHaDBRequest) SetResourceOwnerAccount(v string) *QueryHBaseHaDBRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryHBaseHaDBRequest) SetResourceOwnerId(v int64) *QueryHBaseHaDBRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryHBaseHaDBRequest) SetZoneId(v string) *QueryHBaseHaDBRequest {
	s.ZoneId = &v
	return s
}

type QueryHBaseHaDBResponseBody struct {
	ClusterList *QueryHBaseHaDBResponseBodyClusterList `json:"ClusterList,omitempty" xml:"ClusterList,omitempty" type:"Struct"`
	PageNumber  *int32                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int32                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryHBaseHaDBResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryHBaseHaDBResponseBody) GoString() string {
	return s.String()
}

func (s *QueryHBaseHaDBResponseBody) SetClusterList(v *QueryHBaseHaDBResponseBodyClusterList) *QueryHBaseHaDBResponseBody {
	s.ClusterList = v
	return s
}

func (s *QueryHBaseHaDBResponseBody) SetPageNumber(v int32) *QueryHBaseHaDBResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryHBaseHaDBResponseBody) SetPageSize(v int32) *QueryHBaseHaDBResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryHBaseHaDBResponseBody) SetRequestId(v string) *QueryHBaseHaDBResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryHBaseHaDBResponseBody) SetTotalCount(v int32) *QueryHBaseHaDBResponseBody {
	s.TotalCount = &v
	return s
}

type QueryHBaseHaDBResponseBodyClusterList struct {
	Cluster []*QueryHBaseHaDBResponseBodyClusterListCluster `json:"Cluster,omitempty" xml:"Cluster,omitempty" type:"Repeated"`
}

func (s QueryHBaseHaDBResponseBodyClusterList) String() string {
	return tea.Prettify(s)
}

func (s QueryHBaseHaDBResponseBodyClusterList) GoString() string {
	return s.String()
}

func (s *QueryHBaseHaDBResponseBodyClusterList) SetCluster(v []*QueryHBaseHaDBResponseBodyClusterListCluster) *QueryHBaseHaDBResponseBodyClusterList {
	s.Cluster = v
	return s
}

type QueryHBaseHaDBResponseBodyClusterListCluster struct {
	ActiveName  *string `json:"ActiveName,omitempty" xml:"ActiveName,omitempty"`
	BdsName     *string `json:"BdsName,omitempty" xml:"BdsName,omitempty"`
	HaName      *string `json:"HaName,omitempty" xml:"HaName,omitempty"`
	StandbyName *string `json:"StandbyName,omitempty" xml:"StandbyName,omitempty"`
}

func (s QueryHBaseHaDBResponseBodyClusterListCluster) String() string {
	return tea.Prettify(s)
}

func (s QueryHBaseHaDBResponseBodyClusterListCluster) GoString() string {
	return s.String()
}

func (s *QueryHBaseHaDBResponseBodyClusterListCluster) SetActiveName(v string) *QueryHBaseHaDBResponseBodyClusterListCluster {
	s.ActiveName = &v
	return s
}

func (s *QueryHBaseHaDBResponseBodyClusterListCluster) SetBdsName(v string) *QueryHBaseHaDBResponseBodyClusterListCluster {
	s.BdsName = &v
	return s
}

func (s *QueryHBaseHaDBResponseBodyClusterListCluster) SetHaName(v string) *QueryHBaseHaDBResponseBodyClusterListCluster {
	s.HaName = &v
	return s
}

func (s *QueryHBaseHaDBResponseBodyClusterListCluster) SetStandbyName(v string) *QueryHBaseHaDBResponseBodyClusterListCluster {
	s.StandbyName = &v
	return s
}

type QueryHBaseHaDBResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryHBaseHaDBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryHBaseHaDBResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryHBaseHaDBResponse) GoString() string {
	return s.String()
}

func (s *QueryHBaseHaDBResponse) SetHeaders(v map[string]*string) *QueryHBaseHaDBResponse {
	s.Headers = v
	return s
}

func (s *QueryHBaseHaDBResponse) SetStatusCode(v int32) *QueryHBaseHaDBResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryHBaseHaDBResponse) SetBody(v *QueryHBaseHaDBResponseBody) *QueryHBaseHaDBResponse {
	s.Body = v
	return s
}

type QuerySparkRelateHBaseRequest struct {
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s QuerySparkRelateHBaseRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySparkRelateHBaseRequest) GoString() string {
	return s.String()
}

func (s *QuerySparkRelateHBaseRequest) SetClusterId(v string) *QuerySparkRelateHBaseRequest {
	s.ClusterId = &v
	return s
}

func (s *QuerySparkRelateHBaseRequest) SetOwnerId(v int64) *QuerySparkRelateHBaseRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySparkRelateHBaseRequest) SetRegionId(v string) *QuerySparkRelateHBaseRequest {
	s.RegionId = &v
	return s
}

func (s *QuerySparkRelateHBaseRequest) SetResourceOwnerAccount(v string) *QuerySparkRelateHBaseRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySparkRelateHBaseRequest) SetResourceOwnerId(v int64) *QuerySparkRelateHBaseRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuerySparkRelateHBaseRequest) SetZoneId(v string) *QuerySparkRelateHBaseRequest {
	s.ZoneId = &v
	return s
}

type QuerySparkRelateHBaseResponseBody struct {
	ClusterList *QuerySparkRelateHBaseResponseBodyClusterList `json:"ClusterList,omitempty" xml:"ClusterList,omitempty" type:"Struct"`
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QuerySparkRelateHBaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySparkRelateHBaseResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySparkRelateHBaseResponseBody) SetClusterList(v *QuerySparkRelateHBaseResponseBodyClusterList) *QuerySparkRelateHBaseResponseBody {
	s.ClusterList = v
	return s
}

func (s *QuerySparkRelateHBaseResponseBody) SetRequestId(v string) *QuerySparkRelateHBaseResponseBody {
	s.RequestId = &v
	return s
}

type QuerySparkRelateHBaseResponseBodyClusterList struct {
	Cluster []*QuerySparkRelateHBaseResponseBodyClusterListCluster `json:"Cluster,omitempty" xml:"Cluster,omitempty" type:"Repeated"`
}

func (s QuerySparkRelateHBaseResponseBodyClusterList) String() string {
	return tea.Prettify(s)
}

func (s QuerySparkRelateHBaseResponseBodyClusterList) GoString() string {
	return s.String()
}

func (s *QuerySparkRelateHBaseResponseBodyClusterList) SetCluster(v []*QuerySparkRelateHBaseResponseBodyClusterListCluster) *QuerySparkRelateHBaseResponseBodyClusterList {
	s.Cluster = v
	return s
}

type QuerySparkRelateHBaseResponseBodyClusterListCluster struct {
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName          *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	ClusterType          *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	CoreDiskType         *string `json:"CoreDiskType,omitempty" xml:"CoreDiskType,omitempty"`
	CoreInstanceQuantity *int32  `json:"CoreInstanceQuantity,omitempty" xml:"CoreInstanceQuantity,omitempty"`
	CreateTime           *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DbType               *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	ExpireTime           *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	IsRelated            *bool   `json:"IsRelated,omitempty" xml:"IsRelated,omitempty"`
	LockMode             *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	MainVersion          *string `json:"MainVersion,omitempty" xml:"MainVersion,omitempty"`
	NetType              *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	PayType              *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Reason               *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId               *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s QuerySparkRelateHBaseResponseBodyClusterListCluster) String() string {
	return tea.Prettify(s)
}

func (s QuerySparkRelateHBaseResponseBodyClusterListCluster) GoString() string {
	return s.String()
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) SetClusterId(v string) *QuerySparkRelateHBaseResponseBodyClusterListCluster {
	s.ClusterId = &v
	return s
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) SetClusterName(v string) *QuerySparkRelateHBaseResponseBodyClusterListCluster {
	s.ClusterName = &v
	return s
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) SetClusterType(v string) *QuerySparkRelateHBaseResponseBodyClusterListCluster {
	s.ClusterType = &v
	return s
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) SetCoreDiskType(v string) *QuerySparkRelateHBaseResponseBodyClusterListCluster {
	s.CoreDiskType = &v
	return s
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) SetCoreInstanceQuantity(v int32) *QuerySparkRelateHBaseResponseBodyClusterListCluster {
	s.CoreInstanceQuantity = &v
	return s
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) SetCreateTime(v string) *QuerySparkRelateHBaseResponseBodyClusterListCluster {
	s.CreateTime = &v
	return s
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) SetDbType(v string) *QuerySparkRelateHBaseResponseBodyClusterListCluster {
	s.DbType = &v
	return s
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) SetExpireTime(v string) *QuerySparkRelateHBaseResponseBodyClusterListCluster {
	s.ExpireTime = &v
	return s
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) SetIsRelated(v bool) *QuerySparkRelateHBaseResponseBodyClusterListCluster {
	s.IsRelated = &v
	return s
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) SetLockMode(v string) *QuerySparkRelateHBaseResponseBodyClusterListCluster {
	s.LockMode = &v
	return s
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) SetMainVersion(v string) *QuerySparkRelateHBaseResponseBodyClusterListCluster {
	s.MainVersion = &v
	return s
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) SetNetType(v string) *QuerySparkRelateHBaseResponseBodyClusterListCluster {
	s.NetType = &v
	return s
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) SetPayType(v string) *QuerySparkRelateHBaseResponseBodyClusterListCluster {
	s.PayType = &v
	return s
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) SetReason(v string) *QuerySparkRelateHBaseResponseBodyClusterListCluster {
	s.Reason = &v
	return s
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) SetRegionId(v string) *QuerySparkRelateHBaseResponseBodyClusterListCluster {
	s.RegionId = &v
	return s
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) SetStatus(v string) *QuerySparkRelateHBaseResponseBodyClusterListCluster {
	s.Status = &v
	return s
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) SetUserId(v string) *QuerySparkRelateHBaseResponseBodyClusterListCluster {
	s.UserId = &v
	return s
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) SetZoneId(v string) *QuerySparkRelateHBaseResponseBodyClusterListCluster {
	s.ZoneId = &v
	return s
}

type QuerySparkRelateHBaseResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySparkRelateHBaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySparkRelateHBaseResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySparkRelateHBaseResponse) GoString() string {
	return s.String()
}

func (s *QuerySparkRelateHBaseResponse) SetHeaders(v map[string]*string) *QuerySparkRelateHBaseResponse {
	s.Headers = v
	return s
}

func (s *QuerySparkRelateHBaseResponse) SetStatusCode(v int32) *QuerySparkRelateHBaseResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySparkRelateHBaseResponse) SetBody(v *QuerySparkRelateHBaseResponseBody) *QuerySparkRelateHBaseResponse {
	s.Body = v
	return s
}

type QueryXpackRelatedDBRequest struct {
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RelateDbType         *string `json:"RelateDbType,omitempty" xml:"RelateDbType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s QueryXpackRelatedDBRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryXpackRelatedDBRequest) GoString() string {
	return s.String()
}

func (s *QueryXpackRelatedDBRequest) SetClusterId(v string) *QueryXpackRelatedDBRequest {
	s.ClusterId = &v
	return s
}

func (s *QueryXpackRelatedDBRequest) SetOwnerId(v int64) *QueryXpackRelatedDBRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryXpackRelatedDBRequest) SetRegionId(v string) *QueryXpackRelatedDBRequest {
	s.RegionId = &v
	return s
}

func (s *QueryXpackRelatedDBRequest) SetRelateDbType(v string) *QueryXpackRelatedDBRequest {
	s.RelateDbType = &v
	return s
}

func (s *QueryXpackRelatedDBRequest) SetResourceOwnerAccount(v string) *QueryXpackRelatedDBRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryXpackRelatedDBRequest) SetResourceOwnerId(v int64) *QueryXpackRelatedDBRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryXpackRelatedDBRequest) SetZoneId(v string) *QueryXpackRelatedDBRequest {
	s.ZoneId = &v
	return s
}

type QueryXpackRelatedDBResponseBody struct {
	ClusterList *QueryXpackRelatedDBResponseBodyClusterList `json:"ClusterList,omitempty" xml:"ClusterList,omitempty" type:"Struct"`
	PageNumber  *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryXpackRelatedDBResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryXpackRelatedDBResponseBody) GoString() string {
	return s.String()
}

func (s *QueryXpackRelatedDBResponseBody) SetClusterList(v *QueryXpackRelatedDBResponseBodyClusterList) *QueryXpackRelatedDBResponseBody {
	s.ClusterList = v
	return s
}

func (s *QueryXpackRelatedDBResponseBody) SetPageNumber(v int32) *QueryXpackRelatedDBResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryXpackRelatedDBResponseBody) SetPageSize(v int32) *QueryXpackRelatedDBResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryXpackRelatedDBResponseBody) SetRequestId(v string) *QueryXpackRelatedDBResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryXpackRelatedDBResponseBody) SetTotalCount(v int32) *QueryXpackRelatedDBResponseBody {
	s.TotalCount = &v
	return s
}

type QueryXpackRelatedDBResponseBodyClusterList struct {
	Cluster []*QueryXpackRelatedDBResponseBodyClusterListCluster `json:"Cluster,omitempty" xml:"Cluster,omitempty" type:"Repeated"`
}

func (s QueryXpackRelatedDBResponseBodyClusterList) String() string {
	return tea.Prettify(s)
}

func (s QueryXpackRelatedDBResponseBodyClusterList) GoString() string {
	return s.String()
}

func (s *QueryXpackRelatedDBResponseBodyClusterList) SetCluster(v []*QueryXpackRelatedDBResponseBodyClusterListCluster) *QueryXpackRelatedDBResponseBodyClusterList {
	s.Cluster = v
	return s
}

type QueryXpackRelatedDBResponseBodyClusterListCluster struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	DBType      *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	DBVersion   *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	IsRelated   *bool   `json:"IsRelated,omitempty" xml:"IsRelated,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryXpackRelatedDBResponseBodyClusterListCluster) String() string {
	return tea.Prettify(s)
}

func (s QueryXpackRelatedDBResponseBodyClusterListCluster) GoString() string {
	return s.String()
}

func (s *QueryXpackRelatedDBResponseBodyClusterListCluster) SetClusterId(v string) *QueryXpackRelatedDBResponseBodyClusterListCluster {
	s.ClusterId = &v
	return s
}

func (s *QueryXpackRelatedDBResponseBodyClusterListCluster) SetClusterName(v string) *QueryXpackRelatedDBResponseBodyClusterListCluster {
	s.ClusterName = &v
	return s
}

func (s *QueryXpackRelatedDBResponseBodyClusterListCluster) SetDBType(v string) *QueryXpackRelatedDBResponseBodyClusterListCluster {
	s.DBType = &v
	return s
}

func (s *QueryXpackRelatedDBResponseBodyClusterListCluster) SetDBVersion(v string) *QueryXpackRelatedDBResponseBodyClusterListCluster {
	s.DBVersion = &v
	return s
}

func (s *QueryXpackRelatedDBResponseBodyClusterListCluster) SetIsRelated(v bool) *QueryXpackRelatedDBResponseBodyClusterListCluster {
	s.IsRelated = &v
	return s
}

func (s *QueryXpackRelatedDBResponseBodyClusterListCluster) SetStatus(v string) *QueryXpackRelatedDBResponseBodyClusterListCluster {
	s.Status = &v
	return s
}

type QueryXpackRelatedDBResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryXpackRelatedDBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryXpackRelatedDBResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryXpackRelatedDBResponse) GoString() string {
	return s.String()
}

func (s *QueryXpackRelatedDBResponse) SetHeaders(v map[string]*string) *QueryXpackRelatedDBResponse {
	s.Headers = v
	return s
}

func (s *QueryXpackRelatedDBResponse) SetStatusCode(v int32) *QueryXpackRelatedDBResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryXpackRelatedDBResponse) SetBody(v *QueryXpackRelatedDBResponseBody) *QueryXpackRelatedDBResponse {
	s.Body = v
	return s
}

type RelateDbForHBaseHaRequest struct {
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	HaActive             *string `json:"HaActive,omitempty" xml:"HaActive,omitempty"`
	HaActiveClusterKey   *string `json:"HaActiveClusterKey,omitempty" xml:"HaActiveClusterKey,omitempty"`
	HaActiveDBType       *string `json:"HaActiveDBType,omitempty" xml:"HaActiveDBType,omitempty"`
	HaActiveHbaseFsDir   *string `json:"HaActiveHbaseFsDir,omitempty" xml:"HaActiveHbaseFsDir,omitempty"`
	HaActiveHdfsUri      *string `json:"HaActiveHdfsUri,omitempty" xml:"HaActiveHdfsUri,omitempty"`
	HaActivePassword     *string `json:"HaActivePassword,omitempty" xml:"HaActivePassword,omitempty"`
	HaActiveUser         *string `json:"HaActiveUser,omitempty" xml:"HaActiveUser,omitempty"`
	HaActiveVersion      *string `json:"HaActiveVersion,omitempty" xml:"HaActiveVersion,omitempty"`
	HaMigrateType        *string `json:"HaMigrateType,omitempty" xml:"HaMigrateType,omitempty"`
	HaStandby            *string `json:"HaStandby,omitempty" xml:"HaStandby,omitempty"`
	HaStandbyClusterKey  *string `json:"HaStandbyClusterKey,omitempty" xml:"HaStandbyClusterKey,omitempty"`
	HaStandbyDBType      *string `json:"HaStandbyDBType,omitempty" xml:"HaStandbyDBType,omitempty"`
	HaStandbyHbaseFsDir  *string `json:"HaStandbyHbaseFsDir,omitempty" xml:"HaStandbyHbaseFsDir,omitempty"`
	HaStandbyHdfsUri     *string `json:"HaStandbyHdfsUri,omitempty" xml:"HaStandbyHdfsUri,omitempty"`
	HaStandbyPassword    *string `json:"HaStandbyPassword,omitempty" xml:"HaStandbyPassword,omitempty"`
	HaStandbyUser        *string `json:"HaStandbyUser,omitempty" xml:"HaStandbyUser,omitempty"`
	HaStandbyVersion     *string `json:"HaStandbyVersion,omitempty" xml:"HaStandbyVersion,omitempty"`
	HaTables             *string `json:"HaTables,omitempty" xml:"HaTables,omitempty"`
	IsActiveStandard     *bool   `json:"IsActiveStandard,omitempty" xml:"IsActiveStandard,omitempty"`
	IsStandbyStandard    *bool   `json:"IsStandbyStandard,omitempty" xml:"IsStandbyStandard,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s RelateDbForHBaseHaRequest) String() string {
	return tea.Prettify(s)
}

func (s RelateDbForHBaseHaRequest) GoString() string {
	return s.String()
}

func (s *RelateDbForHBaseHaRequest) SetClusterId(v string) *RelateDbForHBaseHaRequest {
	s.ClusterId = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaActive(v string) *RelateDbForHBaseHaRequest {
	s.HaActive = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaActiveClusterKey(v string) *RelateDbForHBaseHaRequest {
	s.HaActiveClusterKey = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaActiveDBType(v string) *RelateDbForHBaseHaRequest {
	s.HaActiveDBType = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaActiveHbaseFsDir(v string) *RelateDbForHBaseHaRequest {
	s.HaActiveHbaseFsDir = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaActiveHdfsUri(v string) *RelateDbForHBaseHaRequest {
	s.HaActiveHdfsUri = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaActivePassword(v string) *RelateDbForHBaseHaRequest {
	s.HaActivePassword = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaActiveUser(v string) *RelateDbForHBaseHaRequest {
	s.HaActiveUser = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaActiveVersion(v string) *RelateDbForHBaseHaRequest {
	s.HaActiveVersion = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaMigrateType(v string) *RelateDbForHBaseHaRequest {
	s.HaMigrateType = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaStandby(v string) *RelateDbForHBaseHaRequest {
	s.HaStandby = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaStandbyClusterKey(v string) *RelateDbForHBaseHaRequest {
	s.HaStandbyClusterKey = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaStandbyDBType(v string) *RelateDbForHBaseHaRequest {
	s.HaStandbyDBType = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaStandbyHbaseFsDir(v string) *RelateDbForHBaseHaRequest {
	s.HaStandbyHbaseFsDir = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaStandbyHdfsUri(v string) *RelateDbForHBaseHaRequest {
	s.HaStandbyHdfsUri = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaStandbyPassword(v string) *RelateDbForHBaseHaRequest {
	s.HaStandbyPassword = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaStandbyUser(v string) *RelateDbForHBaseHaRequest {
	s.HaStandbyUser = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaStandbyVersion(v string) *RelateDbForHBaseHaRequest {
	s.HaStandbyVersion = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaTables(v string) *RelateDbForHBaseHaRequest {
	s.HaTables = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetIsActiveStandard(v bool) *RelateDbForHBaseHaRequest {
	s.IsActiveStandard = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetIsStandbyStandard(v bool) *RelateDbForHBaseHaRequest {
	s.IsStandbyStandard = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetOwnerId(v int64) *RelateDbForHBaseHaRequest {
	s.OwnerId = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetRegionId(v string) *RelateDbForHBaseHaRequest {
	s.RegionId = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetResourceOwnerAccount(v string) *RelateDbForHBaseHaRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetResourceOwnerId(v int64) *RelateDbForHBaseHaRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetZoneId(v string) *RelateDbForHBaseHaRequest {
	s.ZoneId = &v
	return s
}

type RelateDbForHBaseHaResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RelateDbForHBaseHaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RelateDbForHBaseHaResponseBody) GoString() string {
	return s.String()
}

func (s *RelateDbForHBaseHaResponseBody) SetRequestId(v string) *RelateDbForHBaseHaResponseBody {
	s.RequestId = &v
	return s
}

type RelateDbForHBaseHaResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RelateDbForHBaseHaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RelateDbForHBaseHaResponse) String() string {
	return tea.Prettify(s)
}

func (s RelateDbForHBaseHaResponse) GoString() string {
	return s.String()
}

func (s *RelateDbForHBaseHaResponse) SetHeaders(v map[string]*string) *RelateDbForHBaseHaResponse {
	s.Headers = v
	return s
}

func (s *RelateDbForHBaseHaResponse) SetStatusCode(v int32) *RelateDbForHBaseHaResponse {
	s.StatusCode = &v
	return s
}

func (s *RelateDbForHBaseHaResponse) SetBody(v *RelateDbForHBaseHaResponseBody) *RelateDbForHBaseHaResponse {
	s.Body = v
	return s
}

type ReleasePublicNetworkAddressRequest struct {
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ReleasePublicNetworkAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleasePublicNetworkAddressRequest) GoString() string {
	return s.String()
}

func (s *ReleasePublicNetworkAddressRequest) SetClusterId(v string) *ReleasePublicNetworkAddressRequest {
	s.ClusterId = &v
	return s
}

func (s *ReleasePublicNetworkAddressRequest) SetOwnerId(v int64) *ReleasePublicNetworkAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *ReleasePublicNetworkAddressRequest) SetResourceOwnerAccount(v string) *ReleasePublicNetworkAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReleasePublicNetworkAddressRequest) SetResourceOwnerId(v int64) *ReleasePublicNetworkAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

type ReleasePublicNetworkAddressResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleasePublicNetworkAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleasePublicNetworkAddressResponseBody) GoString() string {
	return s.String()
}

func (s *ReleasePublicNetworkAddressResponseBody) SetRequestId(v string) *ReleasePublicNetworkAddressResponseBody {
	s.RequestId = &v
	return s
}

type ReleasePublicNetworkAddressResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleasePublicNetworkAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleasePublicNetworkAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleasePublicNetworkAddressResponse) GoString() string {
	return s.String()
}

func (s *ReleasePublicNetworkAddressResponse) SetHeaders(v map[string]*string) *ReleasePublicNetworkAddressResponse {
	s.Headers = v
	return s
}

func (s *ReleasePublicNetworkAddressResponse) SetStatusCode(v int32) *ReleasePublicNetworkAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleasePublicNetworkAddressResponse) SetBody(v *ReleasePublicNetworkAddressResponseBody) *ReleasePublicNetworkAddressResponse {
	s.Body = v
	return s
}

type ReleaseSubscriptionRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SubscriptionId       *string `json:"SubscriptionId,omitempty" xml:"SubscriptionId,omitempty"`
}

func (s ReleaseSubscriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseSubscriptionRequest) GoString() string {
	return s.String()
}

func (s *ReleaseSubscriptionRequest) SetOwnerId(v int64) *ReleaseSubscriptionRequest {
	s.OwnerId = &v
	return s
}

func (s *ReleaseSubscriptionRequest) SetResourceOwnerAccount(v string) *ReleaseSubscriptionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReleaseSubscriptionRequest) SetResourceOwnerId(v int64) *ReleaseSubscriptionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReleaseSubscriptionRequest) SetSubscriptionId(v string) *ReleaseSubscriptionRequest {
	s.SubscriptionId = &v
	return s
}

type ReleaseSubscriptionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseSubscriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseSubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseSubscriptionResponseBody) SetRequestId(v string) *ReleaseSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

type ReleaseSubscriptionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseSubscriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseSubscriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseSubscriptionResponse) GoString() string {
	return s.String()
}

func (s *ReleaseSubscriptionResponse) SetHeaders(v map[string]*string) *ReleaseSubscriptionResponse {
	s.Headers = v
	return s
}

func (s *ReleaseSubscriptionResponse) SetStatusCode(v int32) *ReleaseSubscriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseSubscriptionResponse) SetBody(v *ReleaseSubscriptionResponseBody) *ReleaseSubscriptionResponse {
	s.Body = v
	return s
}

type RenewClusterRequest struct {
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Duration             *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PricingCycle         *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RenewClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s RenewClusterRequest) GoString() string {
	return s.String()
}

func (s *RenewClusterRequest) SetClusterId(v string) *RenewClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *RenewClusterRequest) SetDuration(v int32) *RenewClusterRequest {
	s.Duration = &v
	return s
}

func (s *RenewClusterRequest) SetOwnerId(v int64) *RenewClusterRequest {
	s.OwnerId = &v
	return s
}

func (s *RenewClusterRequest) SetPricingCycle(v string) *RenewClusterRequest {
	s.PricingCycle = &v
	return s
}

func (s *RenewClusterRequest) SetResourceOwnerAccount(v string) *RenewClusterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RenewClusterRequest) SetResourceOwnerId(v int64) *RenewClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

type RenewClusterResponseBody struct {
	OrderId   *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenewClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenewClusterResponseBody) GoString() string {
	return s.String()
}

func (s *RenewClusterResponseBody) SetOrderId(v int64) *RenewClusterResponseBody {
	s.OrderId = &v
	return s
}

func (s *RenewClusterResponseBody) SetRequestId(v string) *RenewClusterResponseBody {
	s.RequestId = &v
	return s
}

type RenewClusterResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s RenewClusterResponse) GoString() string {
	return s.String()
}

func (s *RenewClusterResponse) SetHeaders(v map[string]*string) *RenewClusterResponse {
	s.Headers = v
	return s
}

func (s *RenewClusterResponse) SetStatusCode(v int32) *RenewClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewClusterResponse) SetBody(v *RenewClusterResponseBody) *RenewClusterResponse {
	s.Body = v
	return s
}

type ResizeClusterRequest struct {
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CloudType            *string `json:"CloudType,omitempty" xml:"CloudType,omitempty"`
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ColdStorageSize      *string `json:"ColdStorageSize,omitempty" xml:"ColdStorageSize,omitempty"`
	CoreDiskQuantity     *string `json:"CoreDiskQuantity,omitempty" xml:"CoreDiskQuantity,omitempty"`
	CoreDiskSize         *string `json:"CoreDiskSize,omitempty" xml:"CoreDiskSize,omitempty"`
	CoreDiskType         *string `json:"CoreDiskType,omitempty" xml:"CoreDiskType,omitempty"`
	CoreInstanceQuantity *string `json:"CoreInstanceQuantity,omitempty" xml:"CoreInstanceQuantity,omitempty"`
	CoreInstanceType     *string `json:"CoreInstanceType,omitempty" xml:"CoreInstanceType,omitempty"`
	Engine               *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	IsColdStorage        *string `json:"IsColdStorage,omitempty" xml:"IsColdStorage,omitempty"`
	PayType              *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UpgradeType          *string `json:"UpgradeType,omitempty" xml:"UpgradeType,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ResizeClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s ResizeClusterRequest) GoString() string {
	return s.String()
}

func (s *ResizeClusterRequest) SetClientToken(v string) *ResizeClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *ResizeClusterRequest) SetCloudType(v string) *ResizeClusterRequest {
	s.CloudType = &v
	return s
}

func (s *ResizeClusterRequest) SetClusterId(v string) *ResizeClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *ResizeClusterRequest) SetColdStorageSize(v string) *ResizeClusterRequest {
	s.ColdStorageSize = &v
	return s
}

func (s *ResizeClusterRequest) SetCoreDiskQuantity(v string) *ResizeClusterRequest {
	s.CoreDiskQuantity = &v
	return s
}

func (s *ResizeClusterRequest) SetCoreDiskSize(v string) *ResizeClusterRequest {
	s.CoreDiskSize = &v
	return s
}

func (s *ResizeClusterRequest) SetCoreDiskType(v string) *ResizeClusterRequest {
	s.CoreDiskType = &v
	return s
}

func (s *ResizeClusterRequest) SetCoreInstanceQuantity(v string) *ResizeClusterRequest {
	s.CoreInstanceQuantity = &v
	return s
}

func (s *ResizeClusterRequest) SetCoreInstanceType(v string) *ResizeClusterRequest {
	s.CoreInstanceType = &v
	return s
}

func (s *ResizeClusterRequest) SetEngine(v string) *ResizeClusterRequest {
	s.Engine = &v
	return s
}

func (s *ResizeClusterRequest) SetIsColdStorage(v string) *ResizeClusterRequest {
	s.IsColdStorage = &v
	return s
}

func (s *ResizeClusterRequest) SetPayType(v string) *ResizeClusterRequest {
	s.PayType = &v
	return s
}

func (s *ResizeClusterRequest) SetRegionId(v string) *ResizeClusterRequest {
	s.RegionId = &v
	return s
}

func (s *ResizeClusterRequest) SetUpgradeType(v string) *ResizeClusterRequest {
	s.UpgradeType = &v
	return s
}

func (s *ResizeClusterRequest) SetZoneId(v string) *ResizeClusterRequest {
	s.ZoneId = &v
	return s
}

type ResizeClusterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResizeClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResizeClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ResizeClusterResponseBody) SetRequestId(v string) *ResizeClusterResponseBody {
	s.RequestId = &v
	return s
}

type ResizeClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResizeClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResizeClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s ResizeClusterResponse) GoString() string {
	return s.String()
}

func (s *ResizeClusterResponse) SetHeaders(v map[string]*string) *ResizeClusterResponse {
	s.Headers = v
	return s
}

func (s *ResizeClusterResponse) SetStatusCode(v int32) *ResizeClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *ResizeClusterResponse) SetBody(v *ResizeClusterResponseBody) *ResizeClusterResponse {
	s.Body = v
	return s
}

type SparkRelateHBaseRequest struct {
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	HBaseClusterIds      *string `json:"HBaseClusterIds,omitempty" xml:"HBaseClusterIds,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s SparkRelateHBaseRequest) String() string {
	return tea.Prettify(s)
}

func (s SparkRelateHBaseRequest) GoString() string {
	return s.String()
}

func (s *SparkRelateHBaseRequest) SetClusterId(v string) *SparkRelateHBaseRequest {
	s.ClusterId = &v
	return s
}

func (s *SparkRelateHBaseRequest) SetHBaseClusterIds(v string) *SparkRelateHBaseRequest {
	s.HBaseClusterIds = &v
	return s
}

func (s *SparkRelateHBaseRequest) SetOwnerId(v int64) *SparkRelateHBaseRequest {
	s.OwnerId = &v
	return s
}

func (s *SparkRelateHBaseRequest) SetRegionId(v string) *SparkRelateHBaseRequest {
	s.RegionId = &v
	return s
}

func (s *SparkRelateHBaseRequest) SetResourceOwnerAccount(v string) *SparkRelateHBaseRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SparkRelateHBaseRequest) SetResourceOwnerId(v int64) *SparkRelateHBaseRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SparkRelateHBaseRequest) SetZoneId(v string) *SparkRelateHBaseRequest {
	s.ZoneId = &v
	return s
}

type SparkRelateHBaseResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SparkRelateHBaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SparkRelateHBaseResponseBody) GoString() string {
	return s.String()
}

func (s *SparkRelateHBaseResponseBody) SetRequestId(v string) *SparkRelateHBaseResponseBody {
	s.RequestId = &v
	return s
}

type SparkRelateHBaseResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SparkRelateHBaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SparkRelateHBaseResponse) String() string {
	return tea.Prettify(s)
}

func (s SparkRelateHBaseResponse) GoString() string {
	return s.String()
}

func (s *SparkRelateHBaseResponse) SetHeaders(v map[string]*string) *SparkRelateHBaseResponse {
	s.Headers = v
	return s
}

func (s *SparkRelateHBaseResponse) SetStatusCode(v int32) *SparkRelateHBaseResponse {
	s.StatusCode = &v
	return s
}

func (s *SparkRelateHBaseResponse) SetBody(v *SparkRelateHBaseResponseBody) *SparkRelateHBaseResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	OwnerAccount         *string                   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64                    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId           []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string                   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceType         *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag                  []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetOwnerAccount(v string) *TagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *TagResourcesRequest) SetOwnerId(v int64) *TagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceOwnerAccount(v string) *TagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TagResourcesRequest) SetResourceOwnerId(v int64) *TagResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	All                  *bool     `json:"All,omitempty" xml:"All,omitempty"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceType         *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey               []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetOwnerAccount(v string) *UntagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UntagResourcesRequest) SetOwnerId(v int64) *UntagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceOwnerAccount(v string) *UntagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceOwnerId(v int64) *UntagResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UntagResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type UpgradeMinorVersionRequest struct {
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Components           *string `json:"Components,omitempty" xml:"Components,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	UpgradeVersion       *string `json:"UpgradeVersion,omitempty" xml:"UpgradeVersion,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s UpgradeMinorVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeMinorVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeMinorVersionRequest) SetClientToken(v string) *UpgradeMinorVersionRequest {
	s.ClientToken = &v
	return s
}

func (s *UpgradeMinorVersionRequest) SetClusterId(v string) *UpgradeMinorVersionRequest {
	s.ClusterId = &v
	return s
}

func (s *UpgradeMinorVersionRequest) SetComponents(v string) *UpgradeMinorVersionRequest {
	s.Components = &v
	return s
}

func (s *UpgradeMinorVersionRequest) SetOwnerId(v int64) *UpgradeMinorVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *UpgradeMinorVersionRequest) SetRegionId(v string) *UpgradeMinorVersionRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeMinorVersionRequest) SetResourceOwnerAccount(v string) *UpgradeMinorVersionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpgradeMinorVersionRequest) SetResourceOwnerId(v int64) *UpgradeMinorVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpgradeMinorVersionRequest) SetUpgradeVersion(v string) *UpgradeMinorVersionRequest {
	s.UpgradeVersion = &v
	return s
}

func (s *UpgradeMinorVersionRequest) SetZoneId(v string) *UpgradeMinorVersionRequest {
	s.ZoneId = &v
	return s
}

type UpgradeMinorVersionResponseBody struct {
	NewVersion *string `json:"NewVersion,omitempty" xml:"NewVersion,omitempty"`
	OldVersion *string `json:"OldVersion,omitempty" xml:"OldVersion,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeMinorVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeMinorVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeMinorVersionResponseBody) SetNewVersion(v string) *UpgradeMinorVersionResponseBody {
	s.NewVersion = &v
	return s
}

func (s *UpgradeMinorVersionResponseBody) SetOldVersion(v string) *UpgradeMinorVersionResponseBody {
	s.OldVersion = &v
	return s
}

func (s *UpgradeMinorVersionResponseBody) SetRequestId(v string) *UpgradeMinorVersionResponseBody {
	s.RequestId = &v
	return s
}

type UpgradeMinorVersionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeMinorVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeMinorVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeMinorVersionResponse) GoString() string {
	return s.String()
}

func (s *UpgradeMinorVersionResponse) SetHeaders(v map[string]*string) *UpgradeMinorVersionResponse {
	s.Headers = v
	return s
}

func (s *UpgradeMinorVersionResponse) SetStatusCode(v int32) *UpgradeMinorVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeMinorVersionResponse) SetBody(v *UpgradeMinorVersionResponseBody) *UpgradeMinorVersionResponse {
	s.Body = v
	return s
}

type XpackRelateDBRequest struct {
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DbClusterIds         *string `json:"DbClusterIds,omitempty" xml:"DbClusterIds,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RelateDbType         *string `json:"RelateDbType,omitempty" xml:"RelateDbType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s XpackRelateDBRequest) String() string {
	return tea.Prettify(s)
}

func (s XpackRelateDBRequest) GoString() string {
	return s.String()
}

func (s *XpackRelateDBRequest) SetClusterId(v string) *XpackRelateDBRequest {
	s.ClusterId = &v
	return s
}

func (s *XpackRelateDBRequest) SetDbClusterIds(v string) *XpackRelateDBRequest {
	s.DbClusterIds = &v
	return s
}

func (s *XpackRelateDBRequest) SetOwnerId(v int64) *XpackRelateDBRequest {
	s.OwnerId = &v
	return s
}

func (s *XpackRelateDBRequest) SetRegionId(v string) *XpackRelateDBRequest {
	s.RegionId = &v
	return s
}

func (s *XpackRelateDBRequest) SetRelateDbType(v string) *XpackRelateDBRequest {
	s.RelateDbType = &v
	return s
}

func (s *XpackRelateDBRequest) SetResourceOwnerAccount(v string) *XpackRelateDBRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *XpackRelateDBRequest) SetResourceOwnerId(v int64) *XpackRelateDBRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *XpackRelateDBRequest) SetZoneId(v string) *XpackRelateDBRequest {
	s.ZoneId = &v
	return s
}

type XpackRelateDBResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s XpackRelateDBResponseBody) String() string {
	return tea.Prettify(s)
}

func (s XpackRelateDBResponseBody) GoString() string {
	return s.String()
}

func (s *XpackRelateDBResponseBody) SetRequestId(v string) *XpackRelateDBResponseBody {
	s.RequestId = &v
	return s
}

type XpackRelateDBResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *XpackRelateDBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s XpackRelateDBResponse) String() string {
	return tea.Prettify(s)
}

func (s XpackRelateDBResponse) GoString() string {
	return s.String()
}

func (s *XpackRelateDBResponse) SetHeaders(v map[string]*string) *XpackRelateDBResponse {
	s.Headers = v
	return s
}

func (s *XpackRelateDBResponse) SetStatusCode(v int32) *XpackRelateDBResponse {
	s.StatusCode = &v
	return s
}

func (s *XpackRelateDBResponse) SetBody(v *XpackRelateDBResponseBody) *XpackRelateDBResponse {
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
		"ap-southeast-1":        tea.String("hbase.aliyuncs.com"),
		"cn-beijing":            tea.String("hbase.aliyuncs.com"),
		"cn-hangzhou":           tea.String("hbase.aliyuncs.com"),
		"cn-hangzhou-finance":   tea.String("hbase.aliyuncs.com"),
		"cn-hongkong":           tea.String("hbase.aliyuncs.com"),
		"cn-north-2-gov-1":      tea.String("hbase.aliyuncs.com"),
		"cn-qingdao":            tea.String("hbase.aliyuncs.com"),
		"cn-shanghai":           tea.String("hbase.aliyuncs.com"),
		"cn-shanghai-finance-1": tea.String("hbase.aliyuncs.com"),
		"cn-shenzhen":           tea.String("hbase.aliyuncs.com"),
		"cn-shenzhen-finance-1": tea.String("hbase.aliyuncs.com"),
		"cn-guangzhou":          tea.String("hbase.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("hbase"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddUserHdfsInfoWithOptions(request *AddUserHdfsInfoRequest, runtime *util.RuntimeOptions) (_result *AddUserHdfsInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ExtInfo)) {
		query["ExtInfo"] = request.ExtInfo
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddUserHdfsInfo"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddUserHdfsInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddUserHdfsInfo(request *AddUserHdfsInfoRequest) (_result *AddUserHdfsInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddUserHdfsInfoResponse{}
	_body, _err := client.AddUserHdfsInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AllocatePublicNetworkAddressWithOptions(request *AllocatePublicNetworkAddressRequest, runtime *util.RuntimeOptions) (_result *AllocatePublicNetworkAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AllocatePublicNetworkAddress"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AllocatePublicNetworkAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AllocatePublicNetworkAddress(request *AllocatePublicNetworkAddressRequest) (_result *AllocatePublicNetworkAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AllocatePublicNetworkAddressResponse{}
	_body, _err := client.AllocatePublicNetworkAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckVersionsOfComponentsWithOptions(request *CheckVersionsOfComponentsRequest, runtime *util.RuntimeOptions) (_result *CheckVersionsOfComponentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Components)) {
		query["Components"] = request.Components
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckVersionsOfComponents"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckVersionsOfComponentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckVersionsOfComponents(request *CheckVersionsOfComponentsRequest) (_result *CheckVersionsOfComponentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckVersionsOfComponentsResponse{}
	_body, _err := client.CheckVersionsOfComponentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CloseBackupWithOptions(request *CloseBackupRequest, runtime *util.RuntimeOptions) (_result *CloseBackupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CloseBackup"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CloseBackupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CloseBackup(request *CloseBackupRequest) (_result *CloseBackupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CloseBackupResponse{}
	_body, _err := client.CloseBackupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConvertClusterWithOptions(request *ConvertClusterRequest, runtime *util.RuntimeOptions) (_result *ConvertClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PricingCycle)) {
		query["PricingCycle"] = request.PricingCycle
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ConvertCluster"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConvertClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConvertCluster(request *ConvertClusterRequest) (_result *ConvertClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConvertClusterResponse{}
	_body, _err := client.ConvertClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateClusterWithOptions(request *CreateClusterRequest, runtime *util.RuntimeOptions) (_result *CreateClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoRenew)) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !tea.BoolValue(util.IsUnset(request.BackupId)) {
		query["BackupId"] = request.BackupId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CloudType)) {
		query["CloudType"] = request.CloudType
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterName)) {
		query["ClusterName"] = request.ClusterName
	}

	if !tea.BoolValue(util.IsUnset(request.ColdStorageSize)) {
		query["ColdStorageSize"] = request.ColdStorageSize
	}

	if !tea.BoolValue(util.IsUnset(request.CoreDiskQuantity)) {
		query["CoreDiskQuantity"] = request.CoreDiskQuantity
	}

	if !tea.BoolValue(util.IsUnset(request.CoreDiskSize)) {
		query["CoreDiskSize"] = request.CoreDiskSize
	}

	if !tea.BoolValue(util.IsUnset(request.CoreDiskType)) {
		query["CoreDiskType"] = request.CoreDiskType
	}

	if !tea.BoolValue(util.IsUnset(request.CoreInstanceQuantity)) {
		query["CoreInstanceQuantity"] = request.CoreInstanceQuantity
	}

	if !tea.BoolValue(util.IsUnset(request.CoreInstanceType)) {
		query["CoreInstanceType"] = request.CoreInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.DbInstanceConnType)) {
		query["DbInstanceConnType"] = request.DbInstanceConnType
	}

	if !tea.BoolValue(util.IsUnset(request.DbInstanceType)) {
		query["DbInstanceType"] = request.DbInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.DbType)) {
		query["DbType"] = request.DbType
	}

	if !tea.BoolValue(util.IsUnset(request.DepMode)) {
		query["DepMode"] = request.DepMode
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.Engine)) {
		query["Engine"] = request.Engine
	}

	if !tea.BoolValue(util.IsUnset(request.EngineVersion)) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !tea.BoolValue(util.IsUnset(request.IsColdStorage)) {
		query["IsColdStorage"] = request.IsColdStorage
	}

	if !tea.BoolValue(util.IsUnset(request.MasterInstanceType)) {
		query["MasterInstanceType"] = request.MasterInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.NetType)) {
		query["NetType"] = request.NetType
	}

	if !tea.BoolValue(util.IsUnset(request.PayType)) {
		query["PayType"] = request.PayType
	}

	if !tea.BoolValue(util.IsUnset(request.PricingCycle)) {
		query["PricingCycle"] = request.PricingCycle
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RestoreTime)) {
		query["RestoreTime"] = request.RestoreTime
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityIPList)) {
		query["SecurityIPList"] = request.SecurityIPList
	}

	if !tea.BoolValue(util.IsUnset(request.SrcDBInstanceId)) {
		query["SrcDBInstanceId"] = request.SrcDBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCluster"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCluster(request *CreateClusterRequest) (_result *CreateClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateClusterResponse{}
	_body, _err := client.CreateClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGlobalResourceWithOptions(request *CreateGlobalResourceRequest, runtime *util.RuntimeOptions) (_result *CreateGlobalResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceName)) {
		query["ResourceName"] = request.ResourceName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGlobalResource"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGlobalResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGlobalResource(request *CreateGlobalResourceRequest) (_result *CreateGlobalResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGlobalResourceResponse{}
	_body, _err := client.CreateGlobalResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateHbaseSlbServerWithOptions(request *CreateHbaseSlbServerRequest, runtime *util.RuntimeOptions) (_result *CreateHbaseSlbServerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SlbServer)) {
		query["SlbServer"] = request.SlbServer
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateHbaseSlbServer"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateHbaseSlbServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateHbaseSlbServer(request *CreateHbaseSlbServerRequest) (_result *CreateHbaseSlbServerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateHbaseSlbServerResponse{}
	_body, _err := client.CreateHbaseSlbServerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSubscriptionWithOptions(request *CreateSubscriptionRequest, runtime *util.RuntimeOptions) (_result *CreateSubscriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DestinationInstanceId)) {
		query["DestinationInstanceId"] = request.DestinationInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationInstanceRegionId)) {
		query["DestinationInstanceRegionId"] = request.DestinationInstanceRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraContext)) {
		query["ExtraContext"] = request.ExtraContext
	}

	if !tea.BoolValue(util.IsUnset(request.Mapping)) {
		query["Mapping"] = request.Mapping
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SlbServer)) {
		query["SlbServer"] = request.SlbServer
	}

	if !tea.BoolValue(util.IsUnset(request.SourceInstanceId)) {
		query["SourceInstanceId"] = request.SourceInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceInstanceRegionId)) {
		query["SourceInstanceRegionId"] = request.SourceInstanceRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SubscriptionDescription)) {
		query["SubscriptionDescription"] = request.SubscriptionDescription
	}

	if !tea.BoolValue(util.IsUnset(request.SubscriptionType)) {
		query["SubscriptionType"] = request.SubscriptionType
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSubscription"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSubscriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSubscription(request *CreateSubscriptionRequest) (_result *CreateSubscriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSubscriptionResponse{}
	_body, _err := client.CreateSubscriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteClusterWithOptions(request *DeleteClusterRequest, runtime *util.RuntimeOptions) (_result *DeleteClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCluster"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCluster(request *DeleteClusterRequest) (_result *DeleteClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteClusterResponse{}
	_body, _err := client.DeleteClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteGlobalResourceWithOptions(request *DeleteGlobalResourceRequest, runtime *util.RuntimeOptions) (_result *DeleteGlobalResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceName)) {
		query["ResourceName"] = request.ResourceName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGlobalResource"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGlobalResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGlobalResource(request *DeleteGlobalResourceRequest) (_result *DeleteGlobalResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGlobalResourceResponse{}
	_body, _err := client.DeleteGlobalResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteHbaseSlbServerWithOptions(request *DeleteHbaseSlbServerRequest, runtime *util.RuntimeOptions) (_result *DeleteHbaseSlbServerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SlbServer)) {
		query["SlbServer"] = request.SlbServer
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteHbaseSlbServer"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteHbaseSlbServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteHbaseSlbServer(request *DeleteHbaseSlbServerRequest) (_result *DeleteHbaseSlbServerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteHbaseSlbServerResponse{}
	_body, _err := client.DeleteHbaseSlbServerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteServerlessInstanceWithOptions(request *DeleteServerlessInstanceRequest, runtime *util.RuntimeOptions) (_result *DeleteServerlessInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteServerlessInstance"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteServerlessInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteServerlessInstance(request *DeleteServerlessInstanceRequest) (_result *DeleteServerlessInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteServerlessInstanceResponse{}
	_body, _err := client.DeleteServerlessInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUserHdfsInfoWithOptions(request *DeleteUserHdfsInfoRequest, runtime *util.RuntimeOptions) (_result *DeleteUserHdfsInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.NameService)) {
		query["NameService"] = request.NameService
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUserHdfsInfo"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUserHdfsInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUserHdfsInfo(request *DeleteUserHdfsInfoRequest) (_result *DeleteUserHdfsInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserHdfsInfoResponse{}
	_body, _err := client.DeleteUserHdfsInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBackupPolicyWithOptions(request *DescribeBackupPolicyRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBackupPolicy"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBackupPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackupPolicy(request *DescribeBackupPolicyRequest) (_result *DescribeBackupPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupPolicyResponse{}
	_body, _err := client.DescribeBackupPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBackupsWithOptions(request *DescribeBackupsRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupId)) {
		query["BackupId"] = request.BackupId
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTimeUTC)) {
		query["EndTimeUTC"] = request.EndTimeUTC
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimeUTC)) {
		query["StartTimeUTC"] = request.StartTimeUTC
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBackups"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBackupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackups(request *DescribeBackupsRequest) (_result *DescribeBackupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupsResponse{}
	_body, _err := client.DescribeBackupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterAttributeWithOptions(request *DescribeClusterAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeClusterAttribute"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeClusterAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterAttribute(request *DescribeClusterAttributeRequest) (_result *DescribeClusterAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterAttributeResponse{}
	_body, _err := client.DescribeClusterAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterConnectAddrsWithOptions(request *DescribeClusterConnectAddrsRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterConnectAddrsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeClusterConnectAddrs"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeClusterConnectAddrsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterConnectAddrs(request *DescribeClusterConnectAddrsRequest) (_result *DescribeClusterConnectAddrsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterConnectAddrsResponse{}
	_body, _err := client.DescribeClusterConnectAddrsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterListWithOptions(request *DescribeClusterListRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterName)) {
		query["ClusterName"] = request.ClusterName
	}

	if !tea.BoolValue(util.IsUnset(request.DbType)) {
		query["DbType"] = request.DbType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StatusList)) {
		query["StatusList"] = request.StatusList
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeClusterList"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeClusterListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterList(request *DescribeClusterListRequest) (_result *DescribeClusterListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterListResponse{}
	_body, _err := client.DescribeClusterListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterModelWithOptions(request *DescribeClusterModelRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeClusterModel"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeClusterModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterModel(request *DescribeClusterModelRequest) (_result *DescribeClusterModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterModelResponse{}
	_body, _err := client.DescribeClusterModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterWhiteListWithOptions(request *DescribeClusterWhiteListRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeClusterWhiteList"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeClusterWhiteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterWhiteList(request *DescribeClusterWhiteListRequest) (_result *DescribeClusterWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterWhiteListResponse{}
	_body, _err := client.DescribeClusterWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeColdStorageWithOptions(request *DescribeColdStorageRequest, runtime *util.RuntimeOptions) (_result *DescribeColdStorageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeColdStorage"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeColdStorageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeColdStorage(request *DescribeColdStorageRequest) (_result *DescribeColdStorageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeColdStorageResponse{}
	_body, _err := client.DescribeColdStorageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMultiModDbAttributeWithOptions(request *DescribeMultiModDbAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeMultiModDbAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMultiModDbAttribute"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMultiModDbAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMultiModDbAttribute(request *DescribeMultiModDbAttributeRequest) (_result *DescribeMultiModDbAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMultiModDbAttributeResponse{}
	_body, _err := client.DescribeMultiModDbAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRdsVSwitchsWithOptions(request *DescribeRdsVSwitchsRequest, runtime *util.RuntimeOptions) (_result *DescribeRdsVSwitchsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRdsVSwitchs"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRdsVSwitchsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRdsVSwitchs(request *DescribeRdsVSwitchsRequest) (_result *DescribeRdsVSwitchsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRdsVSwitchsResponse{}
	_body, _err := client.DescribeRdsVSwitchsWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeServerlessInstanceWithOptions(request *DescribeServerlessInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeServerlessInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeServerlessInstance"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeServerlessInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeServerlessInstance(request *DescribeServerlessInstanceRequest) (_result *DescribeServerlessInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeServerlessInstanceResponse{}
	_body, _err := client.DescribeServerlessInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSubscriptionInitializeProgressWithOptions(request *DescribeSubscriptionInitializeProgressRequest, runtime *util.RuntimeOptions) (_result *DescribeSubscriptionInitializeProgressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SubscriptionId)) {
		query["SubscriptionId"] = request.SubscriptionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSubscriptionInitializeProgress"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSubscriptionInitializeProgressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSubscriptionInitializeProgress(request *DescribeSubscriptionInitializeProgressRequest) (_result *DescribeSubscriptionInitializeProgressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSubscriptionInitializeProgressResponse{}
	_body, _err := client.DescribeSubscriptionInitializeProgressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSubscriptionPerformanceWithOptions(request *DescribeSubscriptionPerformanceRequest, runtime *util.RuntimeOptions) (_result *DescribeSubscriptionPerformanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceInstanceId)) {
		query["SourceInstanceId"] = request.SourceInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.SubscriptionId)) {
		query["SubscriptionId"] = request.SubscriptionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSubscriptionPerformance"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSubscriptionPerformanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSubscriptionPerformance(request *DescribeSubscriptionPerformanceRequest) (_result *DescribeSubscriptionPerformanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSubscriptionPerformanceResponse{}
	_body, _err := client.DescribeSubscriptionPerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSubscriptionPermissionWithOptions(request *DescribeSubscriptionPermissionRequest, runtime *util.RuntimeOptions) (_result *DescribeSubscriptionPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSubscriptionPermission"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSubscriptionPermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSubscriptionPermission(request *DescribeSubscriptionPermissionRequest) (_result *DescribeSubscriptionPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSubscriptionPermissionResponse{}
	_body, _err := client.DescribeSubscriptionPermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSubscriptionsWithOptions(request *DescribeSubscriptionsRequest, runtime *util.RuntimeOptions) (_result *DescribeSubscriptionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SubscriptionId)) {
		query["SubscriptionId"] = request.SubscriptionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSubscriptions"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSubscriptionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSubscriptions(request *DescribeSubscriptionsRequest) (_result *DescribeSubscriptionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSubscriptionsResponse{}
	_body, _err := client.DescribeSubscriptionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableServerlessPublicConnectionWithOptions(request *EnableServerlessPublicConnectionRequest, runtime *util.RuntimeOptions) (_result *EnableServerlessPublicConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableServerlessPublicConnection"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableServerlessPublicConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableServerlessPublicConnection(request *EnableServerlessPublicConnectionRequest) (_result *EnableServerlessPublicConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableServerlessPublicConnectionResponse{}
	_body, _err := client.EnableServerlessPublicConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMultimodeCmsUrlWithOptions(request *GetMultimodeCmsUrlRequest, runtime *util.RuntimeOptions) (_result *GetMultimodeCmsUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMultimodeCmsUrl"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMultimodeCmsUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMultimodeCmsUrl(request *GetMultimodeCmsUrlRequest) (_result *GetMultimodeCmsUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMultimodeCmsUrlResponse{}
	_body, _err := client.GetMultimodeCmsUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClusterServiceConfigWithOptions(request *ListClusterServiceConfigRequest, runtime *util.RuntimeOptions) (_result *ListClusterServiceConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListClusterServiceConfig"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListClusterServiceConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClusterServiceConfig(request *ListClusterServiceConfigRequest) (_result *ListClusterServiceConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClusterServiceConfigResponse{}
	_body, _err := client.ListClusterServiceConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClusterServiceConfigHistoryWithOptions(request *ListClusterServiceConfigHistoryRequest, runtime *util.RuntimeOptions) (_result *ListClusterServiceConfigHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListClusterServiceConfigHistory"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListClusterServiceConfigHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClusterServiceConfigHistory(request *ListClusterServiceConfigHistoryRequest) (_result *ListClusterServiceConfigHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClusterServiceConfigHistoryResponse{}
	_body, _err := client.ListClusterServiceConfigHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListHbaseInstancesWithOptions(request *ListHbaseInstancesRequest, runtime *util.RuntimeOptions) (_result *ListHbaseInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHbaseInstances"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHbaseInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListHbaseInstances(request *ListHbaseInstancesRequest) (_result *ListHbaseInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListHbaseInstancesResponse{}
	_body, _err := client.ListHbaseInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyBackupPolicyWithOptions(request *ModifyBackupPolicyRequest, runtime *util.RuntimeOptions) (_result *ModifyBackupPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PreferredBackupEndTimeUTC)) {
		query["PreferredBackupEndTimeUTC"] = request.PreferredBackupEndTimeUTC
	}

	if !tea.BoolValue(util.IsUnset(request.PreferredBackupPeriod)) {
		query["PreferredBackupPeriod"] = request.PreferredBackupPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.PreferredBackupStartTimeUTC)) {
		query["PreferredBackupStartTimeUTC"] = request.PreferredBackupStartTimeUTC
	}

	if !tea.BoolValue(util.IsUnset(request.PreferredBackupTime)) {
		query["PreferredBackupTime"] = request.PreferredBackupTime
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyBackupPolicy"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyBackupPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyBackupPolicy(request *ModifyBackupPolicyRequest) (_result *ModifyBackupPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyBackupPolicyResponse{}
	_body, _err := client.ModifyBackupPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyClusterNameWithOptions(request *ModifyClusterNameRequest, runtime *util.RuntimeOptions) (_result *ModifyClusterNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterName)) {
		query["ClusterName"] = request.ClusterName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyClusterName"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyClusterNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyClusterName(request *ModifyClusterNameRequest) (_result *ModifyClusterNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyClusterNameResponse{}
	_body, _err := client.ModifyClusterNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyClusterNetTypeWithOptions(request *ModifyClusterNetTypeRequest, runtime *util.RuntimeOptions) (_result *ModifyClusterNetTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.NetType)) {
		query["NetType"] = request.NetType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyClusterNetType"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyClusterNetTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyClusterNetType(request *ModifyClusterNetTypeRequest) (_result *ModifyClusterNetTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyClusterNetTypeResponse{}
	_body, _err := client.ModifyClusterNetTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyClusterSecurityIpListWithOptions(request *ModifyClusterSecurityIpListRequest, runtime *util.RuntimeOptions) (_result *ModifyClusterSecurityIpListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityIpList)) {
		query["SecurityIpList"] = request.SecurityIpList
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyClusterSecurityIpList"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyClusterSecurityIpListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyClusterSecurityIpList(request *ModifyClusterSecurityIpListRequest) (_result *ModifyClusterSecurityIpListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyClusterSecurityIpListResponse{}
	_body, _err := client.ModifyClusterSecurityIpListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyClusterServiceConfigWithOptions(request *ModifyClusterServiceConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyClusterServiceConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		query["Parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Restart)) {
		query["Restart"] = request.Restart
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		query["Value"] = request.Value
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyClusterServiceConfig"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyClusterServiceConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyClusterServiceConfig(request *ModifyClusterServiceConfigRequest) (_result *ModifyClusterServiceConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyClusterServiceConfigResponse{}
	_body, _err := client.ModifyClusterServiceConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyHasRootPasswordWithOptions(request *ModifyHasRootPasswordRequest, runtime *util.RuntimeOptions) (_result *ModifyHasRootPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.HasPassword)) {
		query["HasPassword"] = request.HasPassword
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyHasRootPassword"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyHasRootPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyHasRootPassword(request *ModifyHasRootPasswordRequest) (_result *ModifyHasRootPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyHasRootPasswordResponse{}
	_body, _err := client.ModifyHasRootPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyRestartClusterWithOptions(request *ModifyRestartClusterRequest, runtime *util.RuntimeOptions) (_result *ModifyRestartClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Components)) {
		query["Components"] = request.Components
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyRestartCluster"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyRestartClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyRestartCluster(request *ModifyRestartClusterRequest) (_result *ModifyRestartClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyRestartClusterResponse{}
	_body, _err := client.ModifyRestartClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyRollbackHasForHbaseWithOptions(request *ModifyRollbackHasForHbaseRequest, runtime *util.RuntimeOptions) (_result *ModifyRollbackHasForHbaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyRollbackHasForHbase"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyRollbackHasForHbaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyRollbackHasForHbase(request *ModifyRollbackHasForHbaseRequest) (_result *ModifyRollbackHasForHbaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyRollbackHasForHbaseResponse{}
	_body, _err := client.ModifyRollbackHasForHbaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifySubscriptionDescriptionWithOptions(request *ModifySubscriptionDescriptionRequest, runtime *util.RuntimeOptions) (_result *ModifySubscriptionDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SubscriptionDescription)) {
		query["SubscriptionDescription"] = request.SubscriptionDescription
	}

	if !tea.BoolValue(util.IsUnset(request.SubscriptionId)) {
		query["SubscriptionId"] = request.SubscriptionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifySubscriptionDescription"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifySubscriptionDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifySubscriptionDescription(request *ModifySubscriptionDescriptionRequest) (_result *ModifySubscriptionDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifySubscriptionDescriptionResponse{}
	_body, _err := client.ModifySubscriptionDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifySubscriptionMappingWithOptions(request *ModifySubscriptionMappingRequest, runtime *util.RuntimeOptions) (_result *ModifySubscriptionMappingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Mapping)) {
		query["Mapping"] = request.Mapping
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SubscriptionId)) {
		query["SubscriptionId"] = request.SubscriptionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifySubscriptionMapping"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifySubscriptionMappingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifySubscriptionMapping(request *ModifySubscriptionMappingRequest) (_result *ModifySubscriptionMappingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifySubscriptionMappingResponse{}
	_body, _err := client.ModifySubscriptionMappingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifySubscriptionPermissionWithOptions(request *ModifySubscriptionPermissionRequest, runtime *util.RuntimeOptions) (_result *ModifySubscriptionPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifySubscriptionPermission"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifySubscriptionPermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifySubscriptionPermission(request *ModifySubscriptionPermissionRequest) (_result *ModifySubscriptionPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifySubscriptionPermissionResponse{}
	_body, _err := client.ModifySubscriptionPermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyUIProxyAccountPasswordWithOptions(request *ModifyUIProxyAccountPasswordRequest, runtime *util.RuntimeOptions) (_result *ModifyUIProxyAccountPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.AccountPassword)) {
		query["AccountPassword"] = request.AccountPassword
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyUIProxyAccountPassword"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyUIProxyAccountPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyUIProxyAccountPassword(request *ModifyUIProxyAccountPasswordRequest) (_result *ModifyUIProxyAccountPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyUIProxyAccountPasswordResponse{}
	_body, _err := client.ModifyUIProxyAccountPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyUpgradeToHasForHbaseWithOptions(request *ModifyUpgradeToHasForHbaseRequest, runtime *util.RuntimeOptions) (_result *ModifyUpgradeToHasForHbaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.HasPassword)) {
		query["HasPassword"] = request.HasPassword
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyUpgradeToHasForHbase"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyUpgradeToHasForHbaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyUpgradeToHasForHbase(request *ModifyUpgradeToHasForHbaseRequest) (_result *ModifyUpgradeToHasForHbaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyUpgradeToHasForHbaseResponse{}
	_body, _err := client.ModifyUpgradeToHasForHbaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MultimodAddComponentsWithOptions(request *MultimodAddComponentsRequest, runtime *util.RuntimeOptions) (_result *MultimodAddComponentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Components)) {
		query["Components"] = request.Components
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("MultimodAddComponents"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MultimodAddComponentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MultimodAddComponents(request *MultimodAddComponentsRequest) (_result *MultimodAddComponentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MultimodAddComponentsResponse{}
	_body, _err := client.MultimodAddComponentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenBackupWithOptions(request *OpenBackupRequest, runtime *util.RuntimeOptions) (_result *OpenBackupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OpenBackup"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenBackupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenBackup(request *OpenBackupRequest) (_result *OpenBackupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenBackupResponse{}
	_body, _err := client.OpenBackupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryHBaseHaDBWithOptions(request *QueryHBaseHaDBRequest, runtime *util.RuntimeOptions) (_result *QueryHBaseHaDBResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryHBaseHaDB"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryHBaseHaDBResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryHBaseHaDB(request *QueryHBaseHaDBRequest) (_result *QueryHBaseHaDBResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryHBaseHaDBResponse{}
	_body, _err := client.QueryHBaseHaDBWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QuerySparkRelateHBaseWithOptions(request *QuerySparkRelateHBaseRequest, runtime *util.RuntimeOptions) (_result *QuerySparkRelateHBaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QuerySparkRelateHBase"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySparkRelateHBaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySparkRelateHBase(request *QuerySparkRelateHBaseRequest) (_result *QuerySparkRelateHBaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySparkRelateHBaseResponse{}
	_body, _err := client.QuerySparkRelateHBaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryXpackRelatedDBWithOptions(request *QueryXpackRelatedDBRequest, runtime *util.RuntimeOptions) (_result *QueryXpackRelatedDBResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RelateDbType)) {
		query["RelateDbType"] = request.RelateDbType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryXpackRelatedDB"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryXpackRelatedDBResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryXpackRelatedDB(request *QueryXpackRelatedDBRequest) (_result *QueryXpackRelatedDBResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryXpackRelatedDBResponse{}
	_body, _err := client.QueryXpackRelatedDBWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RelateDbForHBaseHaWithOptions(request *RelateDbForHBaseHaRequest, runtime *util.RuntimeOptions) (_result *RelateDbForHBaseHaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.HaActive)) {
		query["HaActive"] = request.HaActive
	}

	if !tea.BoolValue(util.IsUnset(request.HaActiveClusterKey)) {
		query["HaActiveClusterKey"] = request.HaActiveClusterKey
	}

	if !tea.BoolValue(util.IsUnset(request.HaActiveDBType)) {
		query["HaActiveDBType"] = request.HaActiveDBType
	}

	if !tea.BoolValue(util.IsUnset(request.HaActiveHbaseFsDir)) {
		query["HaActiveHbaseFsDir"] = request.HaActiveHbaseFsDir
	}

	if !tea.BoolValue(util.IsUnset(request.HaActiveHdfsUri)) {
		query["HaActiveHdfsUri"] = request.HaActiveHdfsUri
	}

	if !tea.BoolValue(util.IsUnset(request.HaActivePassword)) {
		query["HaActivePassword"] = request.HaActivePassword
	}

	if !tea.BoolValue(util.IsUnset(request.HaActiveUser)) {
		query["HaActiveUser"] = request.HaActiveUser
	}

	if !tea.BoolValue(util.IsUnset(request.HaActiveVersion)) {
		query["HaActiveVersion"] = request.HaActiveVersion
	}

	if !tea.BoolValue(util.IsUnset(request.HaMigrateType)) {
		query["HaMigrateType"] = request.HaMigrateType
	}

	if !tea.BoolValue(util.IsUnset(request.HaStandby)) {
		query["HaStandby"] = request.HaStandby
	}

	if !tea.BoolValue(util.IsUnset(request.HaStandbyClusterKey)) {
		query["HaStandbyClusterKey"] = request.HaStandbyClusterKey
	}

	if !tea.BoolValue(util.IsUnset(request.HaStandbyDBType)) {
		query["HaStandbyDBType"] = request.HaStandbyDBType
	}

	if !tea.BoolValue(util.IsUnset(request.HaStandbyHbaseFsDir)) {
		query["HaStandbyHbaseFsDir"] = request.HaStandbyHbaseFsDir
	}

	if !tea.BoolValue(util.IsUnset(request.HaStandbyHdfsUri)) {
		query["HaStandbyHdfsUri"] = request.HaStandbyHdfsUri
	}

	if !tea.BoolValue(util.IsUnset(request.HaStandbyPassword)) {
		query["HaStandbyPassword"] = request.HaStandbyPassword
	}

	if !tea.BoolValue(util.IsUnset(request.HaStandbyUser)) {
		query["HaStandbyUser"] = request.HaStandbyUser
	}

	if !tea.BoolValue(util.IsUnset(request.HaStandbyVersion)) {
		query["HaStandbyVersion"] = request.HaStandbyVersion
	}

	if !tea.BoolValue(util.IsUnset(request.HaTables)) {
		query["HaTables"] = request.HaTables
	}

	if !tea.BoolValue(util.IsUnset(request.IsActiveStandard)) {
		query["IsActiveStandard"] = request.IsActiveStandard
	}

	if !tea.BoolValue(util.IsUnset(request.IsStandbyStandard)) {
		query["IsStandbyStandard"] = request.IsStandbyStandard
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RelateDbForHBaseHa"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RelateDbForHBaseHaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RelateDbForHBaseHa(request *RelateDbForHBaseHaRequest) (_result *RelateDbForHBaseHaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RelateDbForHBaseHaResponse{}
	_body, _err := client.RelateDbForHBaseHaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleasePublicNetworkAddressWithOptions(request *ReleasePublicNetworkAddressRequest, runtime *util.RuntimeOptions) (_result *ReleasePublicNetworkAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleasePublicNetworkAddress"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleasePublicNetworkAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleasePublicNetworkAddress(request *ReleasePublicNetworkAddressRequest) (_result *ReleasePublicNetworkAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleasePublicNetworkAddressResponse{}
	_body, _err := client.ReleasePublicNetworkAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleaseSubscriptionWithOptions(request *ReleaseSubscriptionRequest, runtime *util.RuntimeOptions) (_result *ReleaseSubscriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SubscriptionId)) {
		query["SubscriptionId"] = request.SubscriptionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleaseSubscription"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleaseSubscriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseSubscription(request *ReleaseSubscriptionRequest) (_result *ReleaseSubscriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseSubscriptionResponse{}
	_body, _err := client.ReleaseSubscriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenewClusterWithOptions(request *RenewClusterRequest, runtime *util.RuntimeOptions) (_result *RenewClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PricingCycle)) {
		query["PricingCycle"] = request.PricingCycle
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RenewCluster"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RenewClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenewCluster(request *RenewClusterRequest) (_result *RenewClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenewClusterResponse{}
	_body, _err := client.RenewClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResizeClusterWithOptions(request *ResizeClusterRequest, runtime *util.RuntimeOptions) (_result *ResizeClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CloudType)) {
		query["CloudType"] = request.CloudType
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ColdStorageSize)) {
		query["ColdStorageSize"] = request.ColdStorageSize
	}

	if !tea.BoolValue(util.IsUnset(request.CoreDiskQuantity)) {
		query["CoreDiskQuantity"] = request.CoreDiskQuantity
	}

	if !tea.BoolValue(util.IsUnset(request.CoreDiskSize)) {
		query["CoreDiskSize"] = request.CoreDiskSize
	}

	if !tea.BoolValue(util.IsUnset(request.CoreDiskType)) {
		query["CoreDiskType"] = request.CoreDiskType
	}

	if !tea.BoolValue(util.IsUnset(request.CoreInstanceQuantity)) {
		query["CoreInstanceQuantity"] = request.CoreInstanceQuantity
	}

	if !tea.BoolValue(util.IsUnset(request.CoreInstanceType)) {
		query["CoreInstanceType"] = request.CoreInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.Engine)) {
		query["Engine"] = request.Engine
	}

	if !tea.BoolValue(util.IsUnset(request.IsColdStorage)) {
		query["IsColdStorage"] = request.IsColdStorage
	}

	if !tea.BoolValue(util.IsUnset(request.PayType)) {
		query["PayType"] = request.PayType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UpgradeType)) {
		query["UpgradeType"] = request.UpgradeType
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResizeCluster"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResizeClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResizeCluster(request *ResizeClusterRequest) (_result *ResizeClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResizeClusterResponse{}
	_body, _err := client.ResizeClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SparkRelateHBaseWithOptions(request *SparkRelateHBaseRequest, runtime *util.RuntimeOptions) (_result *SparkRelateHBaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.HBaseClusterIds)) {
		query["HBaseClusterIds"] = request.HBaseClusterIds
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SparkRelateHBase"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SparkRelateHBaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SparkRelateHBase(request *SparkRelateHBaseRequest) (_result *SparkRelateHBaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SparkRelateHBaseResponse{}
	_body, _err := client.SparkRelateHBaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeMinorVersionWithOptions(request *UpgradeMinorVersionRequest, runtime *util.RuntimeOptions) (_result *UpgradeMinorVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Components)) {
		query["Components"] = request.Components
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.UpgradeVersion)) {
		query["UpgradeVersion"] = request.UpgradeVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeMinorVersion"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradeMinorVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeMinorVersion(request *UpgradeMinorVersionRequest) (_result *UpgradeMinorVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeMinorVersionResponse{}
	_body, _err := client.UpgradeMinorVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) XpackRelateDBWithOptions(request *XpackRelateDBRequest, runtime *util.RuntimeOptions) (_result *XpackRelateDBResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DbClusterIds)) {
		query["DbClusterIds"] = request.DbClusterIds
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RelateDbType)) {
		query["RelateDbType"] = request.RelateDbType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("XpackRelateDB"),
		Version:     tea.String("2017-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &XpackRelateDBResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) XpackRelateDB(request *XpackRelateDBRequest) (_result *XpackRelateDBResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &XpackRelateDBResponse{}
	_body, _err := client.XpackRelateDBWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
