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

type AddUserHdfsInfoRequest struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ExtInfo     *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
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

func (s *AddUserHdfsInfoRequest) SetClientToken(v string) *AddUserHdfsInfoRequest {
	s.ClientToken = &v
	return s
}

type AddUserHdfsInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddUserHdfsInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddUserHdfsInfoResponseBody) GoString() string {
	return s.String()
}

func (s *AddUserHdfsInfoResponseBody) SetRequestId(v string) *AddUserHdfsInfoResponseBody {
	s.RequestId = &v
	return s
}

type AddUserHdfsInfoResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddUserHdfsInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddUserHdfsInfoResponse) SetBody(v *AddUserHdfsInfoResponseBody) *AddUserHdfsInfoResponse {
	s.Body = v
	return s
}

type AllocatePublicNetworkAddressRequest struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
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

func (s *AllocatePublicNetworkAddressRequest) SetClientToken(v string) *AllocatePublicNetworkAddressRequest {
	s.ClientToken = &v
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
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AllocatePublicNetworkAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AllocatePublicNetworkAddressResponse) SetBody(v *AllocatePublicNetworkAddressResponseBody) *AllocatePublicNetworkAddressResponse {
	s.Body = v
	return s
}

type CheckComponentsVersionRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Components *string `json:"Components,omitempty" xml:"Components,omitempty"`
}

func (s CheckComponentsVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckComponentsVersionRequest) GoString() string {
	return s.String()
}

func (s *CheckComponentsVersionRequest) SetClusterId(v string) *CheckComponentsVersionRequest {
	s.ClusterId = &v
	return s
}

func (s *CheckComponentsVersionRequest) SetComponents(v string) *CheckComponentsVersionRequest {
	s.Components = &v
	return s
}

type CheckComponentsVersionResponseBody struct {
	Components *CheckComponentsVersionResponseBodyComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Struct"`
	RequestId  *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckComponentsVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckComponentsVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CheckComponentsVersionResponseBody) SetComponents(v *CheckComponentsVersionResponseBodyComponents) *CheckComponentsVersionResponseBody {
	s.Components = v
	return s
}

func (s *CheckComponentsVersionResponseBody) SetRequestId(v string) *CheckComponentsVersionResponseBody {
	s.RequestId = &v
	return s
}

type CheckComponentsVersionResponseBodyComponents struct {
	Component []*CheckComponentsVersionResponseBodyComponentsComponent `json:"Component,omitempty" xml:"Component,omitempty" type:"Repeated"`
}

func (s CheckComponentsVersionResponseBodyComponents) String() string {
	return tea.Prettify(s)
}

func (s CheckComponentsVersionResponseBodyComponents) GoString() string {
	return s.String()
}

func (s *CheckComponentsVersionResponseBodyComponents) SetComponent(v []*CheckComponentsVersionResponseBodyComponentsComponent) *CheckComponentsVersionResponseBodyComponents {
	s.Component = v
	return s
}

type CheckComponentsVersionResponseBodyComponentsComponent struct {
	IsLatestVersion *string `json:"IsLatestVersion,omitempty" xml:"IsLatestVersion,omitempty"`
	Component       *string `json:"Component,omitempty" xml:"Component,omitempty"`
}

func (s CheckComponentsVersionResponseBodyComponentsComponent) String() string {
	return tea.Prettify(s)
}

func (s CheckComponentsVersionResponseBodyComponentsComponent) GoString() string {
	return s.String()
}

func (s *CheckComponentsVersionResponseBodyComponentsComponent) SetIsLatestVersion(v string) *CheckComponentsVersionResponseBodyComponentsComponent {
	s.IsLatestVersion = &v
	return s
}

func (s *CheckComponentsVersionResponseBodyComponentsComponent) SetComponent(v string) *CheckComponentsVersionResponseBodyComponentsComponent {
	s.Component = &v
	return s
}

type CheckComponentsVersionResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckComponentsVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckComponentsVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckComponentsVersionResponse) GoString() string {
	return s.String()
}

func (s *CheckComponentsVersionResponse) SetHeaders(v map[string]*string) *CheckComponentsVersionResponse {
	s.Headers = v
	return s
}

func (s *CheckComponentsVersionResponse) SetBody(v *CheckComponentsVersionResponseBody) *CheckComponentsVersionResponse {
	s.Body = v
	return s
}

type CloseBackupRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
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
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CloseBackupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CloseBackupResponse) SetBody(v *CloseBackupResponseBody) *CloseBackupResponse {
	s.Body = v
	return s
}

type ConvertInstanceRequest struct {
	ClusterId    *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	Duration     *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
}

func (s ConvertInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ConvertInstanceRequest) GoString() string {
	return s.String()
}

func (s *ConvertInstanceRequest) SetClusterId(v string) *ConvertInstanceRequest {
	s.ClusterId = &v
	return s
}

func (s *ConvertInstanceRequest) SetPricingCycle(v string) *ConvertInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *ConvertInstanceRequest) SetDuration(v int32) *ConvertInstanceRequest {
	s.Duration = &v
	return s
}

type ConvertInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId   *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s ConvertInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConvertInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ConvertInstanceResponseBody) SetRequestId(v string) *ConvertInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConvertInstanceResponseBody) SetOrderId(v int64) *ConvertInstanceResponseBody {
	s.OrderId = &v
	return s
}

type ConvertInstanceResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConvertInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConvertInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ConvertInstanceResponse) GoString() string {
	return s.String()
}

func (s *ConvertInstanceResponse) SetHeaders(v map[string]*string) *ConvertInstanceResponse {
	s.Headers = v
	return s
}

func (s *ConvertInstanceResponse) SetBody(v *ConvertInstanceResponseBody) *ConvertInstanceResponse {
	s.Body = v
	return s
}

type CreateBackupPlanRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s CreateBackupPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBackupPlanRequest) GoString() string {
	return s.String()
}

func (s *CreateBackupPlanRequest) SetClusterId(v string) *CreateBackupPlanRequest {
	s.ClusterId = &v
	return s
}

type CreateBackupPlanResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBackupPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBackupPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBackupPlanResponseBody) SetRequestId(v string) *CreateBackupPlanResponseBody {
	s.RequestId = &v
	return s
}

type CreateBackupPlanResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateBackupPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateBackupPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBackupPlanResponse) GoString() string {
	return s.String()
}

func (s *CreateBackupPlanResponse) SetHeaders(v map[string]*string) *CreateBackupPlanResponse {
	s.Headers = v
	return s
}

func (s *CreateBackupPlanResponse) SetBody(v *CreateBackupPlanResponseBody) *CreateBackupPlanResponse {
	s.Body = v
	return s
}

type CreateClusterRequest struct {
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId             *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ClusterName        *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	VpcId              *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VSwitchId          *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	MasterInstanceType *string `json:"MasterInstanceType,omitempty" xml:"MasterInstanceType,omitempty"`
	CoreInstanceType   *string `json:"CoreInstanceType,omitempty" xml:"CoreInstanceType,omitempty"`
	DiskType           *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	DiskSize           *int32  `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	NodeCount          *int32  `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	PayType            *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Engine             *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	SecurityIPList     *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	EngineVersion      *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	ColdStorageSize    *int32  `json:"ColdStorageSize,omitempty" xml:"ColdStorageSize,omitempty"`
	PeriodUnit         *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	Period             *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	AutoRenewPeriod    *int32  `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	ClientToken        *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ResourceGroupId    *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterRequest) SetRegionId(v string) *CreateClusterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateClusterRequest) SetZoneId(v string) *CreateClusterRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateClusterRequest) SetClusterName(v string) *CreateClusterRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateClusterRequest) SetVpcId(v string) *CreateClusterRequest {
	s.VpcId = &v
	return s
}

func (s *CreateClusterRequest) SetVSwitchId(v string) *CreateClusterRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateClusterRequest) SetMasterInstanceType(v string) *CreateClusterRequest {
	s.MasterInstanceType = &v
	return s
}

func (s *CreateClusterRequest) SetCoreInstanceType(v string) *CreateClusterRequest {
	s.CoreInstanceType = &v
	return s
}

func (s *CreateClusterRequest) SetDiskType(v string) *CreateClusterRequest {
	s.DiskType = &v
	return s
}

func (s *CreateClusterRequest) SetDiskSize(v int32) *CreateClusterRequest {
	s.DiskSize = &v
	return s
}

func (s *CreateClusterRequest) SetNodeCount(v int32) *CreateClusterRequest {
	s.NodeCount = &v
	return s
}

func (s *CreateClusterRequest) SetPayType(v string) *CreateClusterRequest {
	s.PayType = &v
	return s
}

func (s *CreateClusterRequest) SetEngine(v string) *CreateClusterRequest {
	s.Engine = &v
	return s
}

func (s *CreateClusterRequest) SetSecurityIPList(v string) *CreateClusterRequest {
	s.SecurityIPList = &v
	return s
}

func (s *CreateClusterRequest) SetEngineVersion(v string) *CreateClusterRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateClusterRequest) SetColdStorageSize(v int32) *CreateClusterRequest {
	s.ColdStorageSize = &v
	return s
}

func (s *CreateClusterRequest) SetPeriodUnit(v string) *CreateClusterRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateClusterRequest) SetPeriod(v int32) *CreateClusterRequest {
	s.Period = &v
	return s
}

func (s *CreateClusterRequest) SetAutoRenewPeriod(v int32) *CreateClusterRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateClusterRequest) SetClientToken(v string) *CreateClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateClusterRequest) SetResourceGroupId(v string) *CreateClusterRequest {
	s.ResourceGroupId = &v
	return s
}

type CreateClusterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClusterResponseBody) SetRequestId(v string) *CreateClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateClusterResponseBody) SetClusterId(v string) *CreateClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateClusterResponseBody) SetOrderId(v string) *CreateClusterResponseBody {
	s.OrderId = &v
	return s
}

type CreateClusterResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateClusterResponse) SetBody(v *CreateClusterResponseBody) *CreateClusterResponse {
	s.Body = v
	return s
}

type CreateGlobalResourceRequest struct {
	ClusterId    *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
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

func (s *CreateGlobalResourceRequest) SetResourceName(v string) *CreateGlobalResourceRequest {
	s.ResourceName = &v
	return s
}

func (s *CreateGlobalResourceRequest) SetResourceType(v string) *CreateGlobalResourceRequest {
	s.ResourceType = &v
	return s
}

func (s *CreateGlobalResourceRequest) SetClientToken(v string) *CreateGlobalResourceRequest {
	s.ClientToken = &v
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
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateGlobalResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateGlobalResourceResponse) SetBody(v *CreateGlobalResourceResponseBody) *CreateGlobalResourceResponse {
	s.Body = v
	return s
}

type CreateHbaseHaSlbRequest struct {
	BdsId       *string `json:"BdsId,omitempty" xml:"BdsId,omitempty"`
	HaId        *string `json:"HaId,omitempty" xml:"HaId,omitempty"`
	HaTypes     *string `json:"HaTypes,omitempty" xml:"HaTypes,omitempty"`
	HbaseType   *string `json:"HbaseType,omitempty" xml:"HbaseType,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateHbaseHaSlbRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateHbaseHaSlbRequest) GoString() string {
	return s.String()
}

func (s *CreateHbaseHaSlbRequest) SetBdsId(v string) *CreateHbaseHaSlbRequest {
	s.BdsId = &v
	return s
}

func (s *CreateHbaseHaSlbRequest) SetHaId(v string) *CreateHbaseHaSlbRequest {
	s.HaId = &v
	return s
}

func (s *CreateHbaseHaSlbRequest) SetHaTypes(v string) *CreateHbaseHaSlbRequest {
	s.HaTypes = &v
	return s
}

func (s *CreateHbaseHaSlbRequest) SetHbaseType(v string) *CreateHbaseHaSlbRequest {
	s.HbaseType = &v
	return s
}

func (s *CreateHbaseHaSlbRequest) SetClientToken(v string) *CreateHbaseHaSlbRequest {
	s.ClientToken = &v
	return s
}

type CreateHbaseHaSlbResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateHbaseHaSlbResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateHbaseHaSlbResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHbaseHaSlbResponseBody) SetRequestId(v string) *CreateHbaseHaSlbResponseBody {
	s.RequestId = &v
	return s
}

type CreateHbaseHaSlbResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateHbaseHaSlbResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateHbaseHaSlbResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateHbaseHaSlbResponse) GoString() string {
	return s.String()
}

func (s *CreateHbaseHaSlbResponse) SetHeaders(v map[string]*string) *CreateHbaseHaSlbResponse {
	s.Headers = v
	return s
}

func (s *CreateHbaseHaSlbResponse) SetBody(v *CreateHbaseHaSlbResponseBody) *CreateHbaseHaSlbResponse {
	s.Body = v
	return s
}

type CreateHBaseSlbServerRequest struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	SlbServer   *string `json:"SlbServer,omitempty" xml:"SlbServer,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateHBaseSlbServerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateHBaseSlbServerRequest) GoString() string {
	return s.String()
}

func (s *CreateHBaseSlbServerRequest) SetClusterId(v string) *CreateHBaseSlbServerRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateHBaseSlbServerRequest) SetSlbServer(v string) *CreateHBaseSlbServerRequest {
	s.SlbServer = &v
	return s
}

func (s *CreateHBaseSlbServerRequest) SetClientToken(v string) *CreateHBaseSlbServerRequest {
	s.ClientToken = &v
	return s
}

type CreateHBaseSlbServerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateHBaseSlbServerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateHBaseSlbServerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHBaseSlbServerResponseBody) SetRequestId(v string) *CreateHBaseSlbServerResponseBody {
	s.RequestId = &v
	return s
}

type CreateHBaseSlbServerResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateHBaseSlbServerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateHBaseSlbServerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateHBaseSlbServerResponse) GoString() string {
	return s.String()
}

func (s *CreateHBaseSlbServerResponse) SetHeaders(v map[string]*string) *CreateHBaseSlbServerResponse {
	s.Headers = v
	return s
}

func (s *CreateHBaseSlbServerResponse) SetBody(v *CreateHBaseSlbServerResponseBody) *CreateHBaseSlbServerResponse {
	s.Body = v
	return s
}

type CreateMultiZoneClusterRequest struct {
	Engine               *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion        *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	ArchVersion          *string `json:"ArchVersion,omitempty" xml:"ArchVersion,omitempty"`
	ClusterName          *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VpcId                *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	MultiZoneCombination *string `json:"MultiZoneCombination,omitempty" xml:"MultiZoneCombination,omitempty"`
	PrimaryZoneId        *string `json:"PrimaryZoneId,omitempty" xml:"PrimaryZoneId,omitempty"`
	PrimaryVSwitchId     *string `json:"PrimaryVSwitchId,omitempty" xml:"PrimaryVSwitchId,omitempty"`
	StandbyZoneId        *string `json:"StandbyZoneId,omitempty" xml:"StandbyZoneId,omitempty"`
	StandbyVSwitchId     *string `json:"StandbyVSwitchId,omitempty" xml:"StandbyVSwitchId,omitempty"`
	ArbiterZoneId        *string `json:"ArbiterZoneId,omitempty" xml:"ArbiterZoneId,omitempty"`
	ArbiterVSwitchId     *string `json:"ArbiterVSwitchId,omitempty" xml:"ArbiterVSwitchId,omitempty"`
	MasterInstanceType   *string `json:"MasterInstanceType,omitempty" xml:"MasterInstanceType,omitempty"`
	CoreInstanceType     *string `json:"CoreInstanceType,omitempty" xml:"CoreInstanceType,omitempty"`
	CoreDiskType         *string `json:"CoreDiskType,omitempty" xml:"CoreDiskType,omitempty"`
	CoreDiskSize         *int32  `json:"CoreDiskSize,omitempty" xml:"CoreDiskSize,omitempty"`
	CoreNodeCount        *int32  `json:"CoreNodeCount,omitempty" xml:"CoreNodeCount,omitempty"`
	LogInstanceType      *string `json:"LogInstanceType,omitempty" xml:"LogInstanceType,omitempty"`
	LogDiskType          *string `json:"LogDiskType,omitempty" xml:"LogDiskType,omitempty"`
	LogDiskSize          *int32  `json:"LogDiskSize,omitempty" xml:"LogDiskSize,omitempty"`
	LogNodeCount         *int32  `json:"LogNodeCount,omitempty" xml:"LogNodeCount,omitempty"`
	SecurityIPList       *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	PayType              *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	PeriodUnit           *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	Period               *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	AutoRenewPeriod      *int32  `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateMultiZoneClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMultiZoneClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateMultiZoneClusterRequest) SetEngine(v string) *CreateMultiZoneClusterRequest {
	s.Engine = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetEngineVersion(v string) *CreateMultiZoneClusterRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetArchVersion(v string) *CreateMultiZoneClusterRequest {
	s.ArchVersion = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetClusterName(v string) *CreateMultiZoneClusterRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetRegionId(v string) *CreateMultiZoneClusterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetVpcId(v string) *CreateMultiZoneClusterRequest {
	s.VpcId = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetMultiZoneCombination(v string) *CreateMultiZoneClusterRequest {
	s.MultiZoneCombination = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetPrimaryZoneId(v string) *CreateMultiZoneClusterRequest {
	s.PrimaryZoneId = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetPrimaryVSwitchId(v string) *CreateMultiZoneClusterRequest {
	s.PrimaryVSwitchId = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetStandbyZoneId(v string) *CreateMultiZoneClusterRequest {
	s.StandbyZoneId = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetStandbyVSwitchId(v string) *CreateMultiZoneClusterRequest {
	s.StandbyVSwitchId = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetArbiterZoneId(v string) *CreateMultiZoneClusterRequest {
	s.ArbiterZoneId = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetArbiterVSwitchId(v string) *CreateMultiZoneClusterRequest {
	s.ArbiterVSwitchId = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetMasterInstanceType(v string) *CreateMultiZoneClusterRequest {
	s.MasterInstanceType = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetCoreInstanceType(v string) *CreateMultiZoneClusterRequest {
	s.CoreInstanceType = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetCoreDiskType(v string) *CreateMultiZoneClusterRequest {
	s.CoreDiskType = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetCoreDiskSize(v int32) *CreateMultiZoneClusterRequest {
	s.CoreDiskSize = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetCoreNodeCount(v int32) *CreateMultiZoneClusterRequest {
	s.CoreNodeCount = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetLogInstanceType(v string) *CreateMultiZoneClusterRequest {
	s.LogInstanceType = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetLogDiskType(v string) *CreateMultiZoneClusterRequest {
	s.LogDiskType = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetLogDiskSize(v int32) *CreateMultiZoneClusterRequest {
	s.LogDiskSize = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetLogNodeCount(v int32) *CreateMultiZoneClusterRequest {
	s.LogNodeCount = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetSecurityIPList(v string) *CreateMultiZoneClusterRequest {
	s.SecurityIPList = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetPayType(v string) *CreateMultiZoneClusterRequest {
	s.PayType = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetPeriodUnit(v string) *CreateMultiZoneClusterRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetPeriod(v int32) *CreateMultiZoneClusterRequest {
	s.Period = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetAutoRenewPeriod(v int32) *CreateMultiZoneClusterRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetClientToken(v string) *CreateMultiZoneClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetResourceGroupId(v string) *CreateMultiZoneClusterRequest {
	s.ResourceGroupId = &v
	return s
}

type CreateMultiZoneClusterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateMultiZoneClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMultiZoneClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMultiZoneClusterResponseBody) SetRequestId(v string) *CreateMultiZoneClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMultiZoneClusterResponseBody) SetClusterId(v string) *CreateMultiZoneClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateMultiZoneClusterResponseBody) SetOrderId(v string) *CreateMultiZoneClusterResponseBody {
	s.OrderId = &v
	return s
}

type CreateMultiZoneClusterResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateMultiZoneClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMultiZoneClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMultiZoneClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateMultiZoneClusterResponse) SetHeaders(v map[string]*string) *CreateMultiZoneClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateMultiZoneClusterResponse) SetBody(v *CreateMultiZoneClusterResponseBody) *CreateMultiZoneClusterResponse {
	s.Body = v
	return s
}

type CreateRestorePlanRequest struct {
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	TargetClusterId *string `json:"TargetClusterId,omitempty" xml:"TargetClusterId,omitempty"`
	RestoreAllTable *bool   `json:"RestoreAllTable,omitempty" xml:"RestoreAllTable,omitempty"`
	RestoreByCopy   *bool   `json:"RestoreByCopy,omitempty" xml:"RestoreByCopy,omitempty"`
	RestoreToDate   *string `json:"RestoreToDate,omitempty" xml:"RestoreToDate,omitempty"`
	Tables          *string `json:"Tables,omitempty" xml:"Tables,omitempty"`
}

func (s CreateRestorePlanRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRestorePlanRequest) GoString() string {
	return s.String()
}

func (s *CreateRestorePlanRequest) SetClusterId(v string) *CreateRestorePlanRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateRestorePlanRequest) SetTargetClusterId(v string) *CreateRestorePlanRequest {
	s.TargetClusterId = &v
	return s
}

func (s *CreateRestorePlanRequest) SetRestoreAllTable(v bool) *CreateRestorePlanRequest {
	s.RestoreAllTable = &v
	return s
}

func (s *CreateRestorePlanRequest) SetRestoreByCopy(v bool) *CreateRestorePlanRequest {
	s.RestoreByCopy = &v
	return s
}

func (s *CreateRestorePlanRequest) SetRestoreToDate(v string) *CreateRestorePlanRequest {
	s.RestoreToDate = &v
	return s
}

func (s *CreateRestorePlanRequest) SetTables(v string) *CreateRestorePlanRequest {
	s.Tables = &v
	return s
}

type CreateRestorePlanResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRestorePlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRestorePlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRestorePlanResponseBody) SetRequestId(v string) *CreateRestorePlanResponseBody {
	s.RequestId = &v
	return s
}

type CreateRestorePlanResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateRestorePlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRestorePlanResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRestorePlanResponse) GoString() string {
	return s.String()
}

func (s *CreateRestorePlanResponse) SetHeaders(v map[string]*string) *CreateRestorePlanResponse {
	s.Headers = v
	return s
}

func (s *CreateRestorePlanResponse) SetBody(v *CreateRestorePlanResponseBody) *CreateRestorePlanResponse {
	s.Body = v
	return s
}

type CreateServerlessClusterRequest struct {
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ClusterName          *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	VpcId                *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VSwitchId            *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	PayType              *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	PeriodUnit           *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	Period               *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	AutoRenewPeriod      *int32  `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	ServerlessSpec       *string `json:"ServerlessSpec,omitempty" xml:"ServerlessSpec,omitempty"`
	ServerlessCapability *int32  `json:"ServerlessCapability,omitempty" xml:"ServerlessCapability,omitempty"`
	ServerlessStorage    *int32  `json:"ServerlessStorage,omitempty" xml:"ServerlessStorage,omitempty"`
	Engine               *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion        *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ClientType           *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateServerlessClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServerlessClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateServerlessClusterRequest) SetRegionId(v string) *CreateServerlessClusterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetZoneId(v string) *CreateServerlessClusterRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetClusterName(v string) *CreateServerlessClusterRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetVpcId(v string) *CreateServerlessClusterRequest {
	s.VpcId = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetVSwitchId(v string) *CreateServerlessClusterRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetPayType(v string) *CreateServerlessClusterRequest {
	s.PayType = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetPeriodUnit(v string) *CreateServerlessClusterRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetPeriod(v int32) *CreateServerlessClusterRequest {
	s.Period = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetAutoRenewPeriod(v int32) *CreateServerlessClusterRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetServerlessSpec(v string) *CreateServerlessClusterRequest {
	s.ServerlessSpec = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetServerlessCapability(v int32) *CreateServerlessClusterRequest {
	s.ServerlessCapability = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetServerlessStorage(v int32) *CreateServerlessClusterRequest {
	s.ServerlessStorage = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetEngine(v string) *CreateServerlessClusterRequest {
	s.Engine = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetEngineVersion(v string) *CreateServerlessClusterRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetClientToken(v string) *CreateServerlessClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetClientType(v string) *CreateServerlessClusterRequest {
	s.ClientType = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetResourceGroupId(v string) *CreateServerlessClusterRequest {
	s.ResourceGroupId = &v
	return s
}

type CreateServerlessClusterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	PassWord  *string `json:"PassWord,omitempty" xml:"PassWord,omitempty"`
}

func (s CreateServerlessClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServerlessClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServerlessClusterResponseBody) SetRequestId(v string) *CreateServerlessClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServerlessClusterResponseBody) SetClusterId(v string) *CreateServerlessClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateServerlessClusterResponseBody) SetOrderId(v string) *CreateServerlessClusterResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateServerlessClusterResponseBody) SetPassWord(v string) *CreateServerlessClusterResponseBody {
	s.PassWord = &v
	return s
}

type CreateServerlessClusterResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateServerlessClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateServerlessClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServerlessClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateServerlessClusterResponse) SetHeaders(v map[string]*string) *CreateServerlessClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateServerlessClusterResponse) SetBody(v *CreateServerlessClusterResponseBody) *CreateServerlessClusterResponse {
	s.Body = v
	return s
}

type DeleteGlobalResourceRequest struct {
	ClusterId    *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
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

func (s *DeleteGlobalResourceRequest) SetResourceName(v string) *DeleteGlobalResourceRequest {
	s.ResourceName = &v
	return s
}

func (s *DeleteGlobalResourceRequest) SetResourceType(v string) *DeleteGlobalResourceRequest {
	s.ResourceType = &v
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
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteGlobalResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteGlobalResourceResponse) SetBody(v *DeleteGlobalResourceResponseBody) *DeleteGlobalResourceResponse {
	s.Body = v
	return s
}

type DeleteHBaseHaDBRequest struct {
	BdsId *string `json:"BdsId,omitempty" xml:"BdsId,omitempty"`
	HaId  *string `json:"HaId,omitempty" xml:"HaId,omitempty"`
}

func (s DeleteHBaseHaDBRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteHBaseHaDBRequest) GoString() string {
	return s.String()
}

func (s *DeleteHBaseHaDBRequest) SetBdsId(v string) *DeleteHBaseHaDBRequest {
	s.BdsId = &v
	return s
}

func (s *DeleteHBaseHaDBRequest) SetHaId(v string) *DeleteHBaseHaDBRequest {
	s.HaId = &v
	return s
}

type DeleteHBaseHaDBResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteHBaseHaDBResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteHBaseHaDBResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHBaseHaDBResponseBody) SetRequestId(v string) *DeleteHBaseHaDBResponseBody {
	s.RequestId = &v
	return s
}

type DeleteHBaseHaDBResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteHBaseHaDBResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteHBaseHaDBResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteHBaseHaDBResponse) GoString() string {
	return s.String()
}

func (s *DeleteHBaseHaDBResponse) SetHeaders(v map[string]*string) *DeleteHBaseHaDBResponse {
	s.Headers = v
	return s
}

func (s *DeleteHBaseHaDBResponse) SetBody(v *DeleteHBaseHaDBResponseBody) *DeleteHBaseHaDBResponse {
	s.Body = v
	return s
}

type DeleteHbaseHaSlbRequest struct {
	BdsId   *string `json:"BdsId,omitempty" xml:"BdsId,omitempty"`
	HaId    *string `json:"HaId,omitempty" xml:"HaId,omitempty"`
	HaTypes *string `json:"HaTypes,omitempty" xml:"HaTypes,omitempty"`
}

func (s DeleteHbaseHaSlbRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteHbaseHaSlbRequest) GoString() string {
	return s.String()
}

func (s *DeleteHbaseHaSlbRequest) SetBdsId(v string) *DeleteHbaseHaSlbRequest {
	s.BdsId = &v
	return s
}

func (s *DeleteHbaseHaSlbRequest) SetHaId(v string) *DeleteHbaseHaSlbRequest {
	s.HaId = &v
	return s
}

func (s *DeleteHbaseHaSlbRequest) SetHaTypes(v string) *DeleteHbaseHaSlbRequest {
	s.HaTypes = &v
	return s
}

type DeleteHbaseHaSlbResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteHbaseHaSlbResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteHbaseHaSlbResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHbaseHaSlbResponseBody) SetRequestId(v string) *DeleteHbaseHaSlbResponseBody {
	s.RequestId = &v
	return s
}

type DeleteHbaseHaSlbResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteHbaseHaSlbResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteHbaseHaSlbResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteHbaseHaSlbResponse) GoString() string {
	return s.String()
}

func (s *DeleteHbaseHaSlbResponse) SetHeaders(v map[string]*string) *DeleteHbaseHaSlbResponse {
	s.Headers = v
	return s
}

func (s *DeleteHbaseHaSlbResponse) SetBody(v *DeleteHbaseHaSlbResponseBody) *DeleteHbaseHaSlbResponse {
	s.Body = v
	return s
}

type DeleteHBaseSlbServerRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	SlbServer *string `json:"SlbServer,omitempty" xml:"SlbServer,omitempty"`
}

func (s DeleteHBaseSlbServerRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteHBaseSlbServerRequest) GoString() string {
	return s.String()
}

func (s *DeleteHBaseSlbServerRequest) SetClusterId(v string) *DeleteHBaseSlbServerRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteHBaseSlbServerRequest) SetSlbServer(v string) *DeleteHBaseSlbServerRequest {
	s.SlbServer = &v
	return s
}

type DeleteHBaseSlbServerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteHBaseSlbServerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteHBaseSlbServerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHBaseSlbServerResponseBody) SetRequestId(v string) *DeleteHBaseSlbServerResponseBody {
	s.RequestId = &v
	return s
}

type DeleteHBaseSlbServerResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteHBaseSlbServerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteHBaseSlbServerResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteHBaseSlbServerResponse) GoString() string {
	return s.String()
}

func (s *DeleteHBaseSlbServerResponse) SetHeaders(v map[string]*string) *DeleteHBaseSlbServerResponse {
	s.Headers = v
	return s
}

func (s *DeleteHBaseSlbServerResponse) SetBody(v *DeleteHBaseSlbServerResponseBody) *DeleteHBaseSlbServerResponse {
	s.Body = v
	return s
}

type DeleteInstanceRequest struct {
	ClusterId           *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ImmediateDeleteFlag *bool   `json:"ImmediateDeleteFlag,omitempty" xml:"ImmediateDeleteFlag,omitempty"`
}

func (s DeleteInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceRequest) SetClusterId(v string) *DeleteInstanceRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteInstanceRequest) SetImmediateDeleteFlag(v bool) *DeleteInstanceRequest {
	s.ImmediateDeleteFlag = &v
	return s
}

type DeleteInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponseBody) SetRequestId(v string) *DeleteInstanceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteInstanceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponse) SetHeaders(v map[string]*string) *DeleteInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceResponse) SetBody(v *DeleteInstanceResponseBody) *DeleteInstanceResponse {
	s.Body = v
	return s
}

type DeleteMultiZoneClusterRequest struct {
	ClusterId           *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ImmediateDeleteFlag *bool   `json:"ImmediateDeleteFlag,omitempty" xml:"ImmediateDeleteFlag,omitempty"`
}

func (s DeleteMultiZoneClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMultiZoneClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteMultiZoneClusterRequest) SetClusterId(v string) *DeleteMultiZoneClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteMultiZoneClusterRequest) SetImmediateDeleteFlag(v bool) *DeleteMultiZoneClusterRequest {
	s.ImmediateDeleteFlag = &v
	return s
}

type DeleteMultiZoneClusterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMultiZoneClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMultiZoneClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMultiZoneClusterResponseBody) SetRequestId(v string) *DeleteMultiZoneClusterResponseBody {
	s.RequestId = &v
	return s
}

type DeleteMultiZoneClusterResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteMultiZoneClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteMultiZoneClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMultiZoneClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteMultiZoneClusterResponse) SetHeaders(v map[string]*string) *DeleteMultiZoneClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteMultiZoneClusterResponse) SetBody(v *DeleteMultiZoneClusterResponseBody) *DeleteMultiZoneClusterResponse {
	s.Body = v
	return s
}

type DeleteServerlessClusterRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId    *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DeleteServerlessClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteServerlessClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteServerlessClusterRequest) SetClusterId(v string) *DeleteServerlessClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteServerlessClusterRequest) SetRegionId(v string) *DeleteServerlessClusterRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteServerlessClusterRequest) SetZoneId(v string) *DeleteServerlessClusterRequest {
	s.ZoneId = &v
	return s
}

type DeleteServerlessClusterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServerlessClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteServerlessClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServerlessClusterResponseBody) SetRequestId(v string) *DeleteServerlessClusterResponseBody {
	s.RequestId = &v
	return s
}

type DeleteServerlessClusterResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteServerlessClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteServerlessClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServerlessClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteServerlessClusterResponse) SetHeaders(v map[string]*string) *DeleteServerlessClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteServerlessClusterResponse) SetBody(v *DeleteServerlessClusterResponseBody) *DeleteServerlessClusterResponse {
	s.Body = v
	return s
}

type DeleteUserHdfsInfoRequest struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	NameService *string `json:"NameService,omitempty" xml:"NameService,omitempty"`
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

type DeleteUserHdfsInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserHdfsInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserHdfsInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserHdfsInfoResponseBody) SetRequestId(v string) *DeleteUserHdfsInfoResponseBody {
	s.RequestId = &v
	return s
}

type DeleteUserHdfsInfoResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteUserHdfsInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteUserHdfsInfoResponse) SetBody(v *DeleteUserHdfsInfoResponseBody) *DeleteUserHdfsInfoResponse {
	s.Body = v
	return s
}

type DescribeAvailableResourceRequest struct {
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId     *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAvailableResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceRequest) SetChargeType(v string) *DescribeAvailableResourceRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetRegionId(v string) *DescribeAvailableResourceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetZoneId(v string) *DescribeAvailableResourceRequest {
	s.ZoneId = &v
	return s
}

type DescribeAvailableResourceResponseBody struct {
	RequestId      *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AvailableZones *DescribeAvailableResourceResponseBodyAvailableZones `json:"AvailableZones,omitempty" xml:"AvailableZones,omitempty" type:"Struct"`
}

func (s DescribeAvailableResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBody) SetRequestId(v string) *DescribeAvailableResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAvailableResourceResponseBody) SetAvailableZones(v *DescribeAvailableResourceResponseBodyAvailableZones) *DescribeAvailableResourceResponseBody {
	s.AvailableZones = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZones struct {
	AvailableZone []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone `json:"AvailableZone,omitempty" xml:"AvailableZone,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZones) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZones) SetAvailableZone(v []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) *DescribeAvailableResourceResponseBodyAvailableZones {
	s.AvailableZone = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone struct {
	SupportedEngines *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEngines `json:"SupportedEngines,omitempty" xml:"SupportedEngines,omitempty" type:"Struct"`
	ZoneId           *string                                                                           `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	MasterResources  *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResources  `json:"MasterResources,omitempty" xml:"MasterResources,omitempty" type:"Struct"`
	RegionId         *string                                                                           `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) SetSupportedEngines(v *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEngines) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone {
	s.SupportedEngines = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) SetZoneId(v string) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone {
	s.ZoneId = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) SetMasterResources(v *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResources) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone {
	s.MasterResources = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) SetRegionId(v string) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone {
	s.RegionId = &v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEngines struct {
	SupportedEngine []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine `json:"SupportedEngine,omitempty" xml:"SupportedEngine,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEngines) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEngines) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEngines) SetSupportedEngine(v []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEngines {
	s.SupportedEngine = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine struct {
	SupportedEngineVersions *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersions `json:"SupportedEngineVersions,omitempty" xml:"SupportedEngineVersions,omitempty" type:"Struct"`
	Engine                  *string                                                                                                                 `json:"Engine,omitempty" xml:"Engine,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine) SetSupportedEngineVersions(v *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersions) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine {
	s.SupportedEngineVersions = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine) SetEngine(v string) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine {
	s.Engine = &v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersions struct {
	SupportedEngineVersion []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion `json:"SupportedEngineVersion,omitempty" xml:"SupportedEngineVersion,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersions) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersions) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersions) SetSupportedEngineVersion(v []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersions {
	s.SupportedEngineVersion = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion struct {
	Version             *string                                                                                                                                                          `json:"Version,omitempty" xml:"Version,omitempty"`
	SupportedCategories *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategories `json:"SupportedCategories,omitempty" xml:"SupportedCategories,omitempty" type:"Struct"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion) SetVersion(v string) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion {
	s.Version = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion) SetSupportedCategories(v *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategories) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion {
	s.SupportedCategories = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategories struct {
	SupportedCategories []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories `json:"SupportedCategories,omitempty" xml:"SupportedCategories,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategories) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategories) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategories) SetSupportedCategories(v []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategories {
	s.SupportedCategories = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories struct {
	SupportedStorageTypes *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypes `json:"SupportedStorageTypes,omitempty" xml:"SupportedStorageTypes,omitempty" type:"Struct"`
	Category              *string                                                                                                                                                                                                  `json:"Category,omitempty" xml:"Category,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories) SetSupportedStorageTypes(v *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypes) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories {
	s.SupportedStorageTypes = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories) SetCategory(v string) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories {
	s.Category = &v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypes struct {
	SupportedStorageType []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType `json:"SupportedStorageType,omitempty" xml:"SupportedStorageType,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypes) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypes) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypes) SetSupportedStorageType(v []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypes {
	s.SupportedStorageType = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType struct {
	CoreResources *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResources `json:"CoreResources,omitempty" xml:"CoreResources,omitempty" type:"Struct"`
	StorageType   *string                                                                                                                                                                                                                                   `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType) SetCoreResources(v *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResources) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType {
	s.CoreResources = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType) SetStorageType(v string) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType {
	s.StorageType = &v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResources struct {
	CoreResource []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource `json:"CoreResource,omitempty" xml:"CoreResource,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResources) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResources) SetCoreResource(v []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResources {
	s.CoreResource = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource struct {
	InstanceTypeDetail     *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail     `json:"InstanceTypeDetail,omitempty" xml:"InstanceTypeDetail,omitempty" type:"Struct"`
	DBInstanceStorageRange *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange `json:"DBInstanceStorageRange,omitempty" xml:"DBInstanceStorageRange,omitempty" type:"Struct"`
	MaxCoreCount           *int32                                                                                                                                                                                                                                                                      `json:"MaxCoreCount,omitempty" xml:"MaxCoreCount,omitempty"`
	InstanceType           *string                                                                                                                                                                                                                                                                     `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource) SetInstanceTypeDetail(v *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource {
	s.InstanceTypeDetail = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource) SetDBInstanceStorageRange(v *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource {
	s.DBInstanceStorageRange = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource) SetMaxCoreCount(v int32) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource {
	s.MaxCoreCount = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource) SetInstanceType(v string) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource {
	s.InstanceType = &v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail struct {
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Mem *int32 `json:"Mem,omitempty" xml:"Mem,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail) SetCpu(v int32) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail {
	s.Cpu = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail) SetMem(v int32) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail {
	s.Mem = &v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange struct {
	MaxSize  *int32 `json:"MaxSize,omitempty" xml:"MaxSize,omitempty"`
	StepSize *int32 `json:"StepSize,omitempty" xml:"StepSize,omitempty"`
	MinSize  *int32 `json:"MinSize,omitempty" xml:"MinSize,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange) SetMaxSize(v int32) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange {
	s.MaxSize = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange) SetStepSize(v int32) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange {
	s.StepSize = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange) SetMinSize(v int32) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange {
	s.MinSize = &v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResources struct {
	MasterResource []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource `json:"MasterResource,omitempty" xml:"MasterResource,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResources) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResources) SetMasterResource(v []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResources {
	s.MasterResource = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource struct {
	InstanceTypeDetail *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail `json:"InstanceTypeDetail,omitempty" xml:"InstanceTypeDetail,omitempty" type:"Struct"`
	InstanceType       *string                                                                                                          `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource) SetInstanceTypeDetail(v *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource {
	s.InstanceTypeDetail = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource) SetInstanceType(v string) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource {
	s.InstanceType = &v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail struct {
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Mem *int32 `json:"Mem,omitempty" xml:"Mem,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail) SetCpu(v int32) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail {
	s.Cpu = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail) SetMem(v int32) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail {
	s.Mem = &v
	return s
}

type DescribeAvailableResourceResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAvailableResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAvailableResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponse) SetHeaders(v map[string]*string) *DescribeAvailableResourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeAvailableResourceResponse) SetBody(v *DescribeAvailableResourceResponseBody) *DescribeAvailableResourceResponse {
	s.Body = v
	return s
}

type DescribeBackupPlanConfigRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeBackupPlanConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPlanConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlanConfigRequest) SetClusterId(v string) *DescribeBackupPlanConfigRequest {
	s.ClusterId = &v
	return s
}

type DescribeBackupPlanConfigResponseBody struct {
	RequestId           *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FullBackupCycle     *int32                                      `json:"FullBackupCycle,omitempty" xml:"FullBackupCycle,omitempty"`
	NextFullBackupDate  *string                                     `json:"NextFullBackupDate,omitempty" xml:"NextFullBackupDate,omitempty"`
	Tables              *DescribeBackupPlanConfigResponseBodyTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Struct"`
	MinHFileBackupCount *int32                                      `json:"MinHFileBackupCount,omitempty" xml:"MinHFileBackupCount,omitempty"`
}

func (s DescribeBackupPlanConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPlanConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlanConfigResponseBody) SetRequestId(v string) *DescribeBackupPlanConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupPlanConfigResponseBody) SetFullBackupCycle(v int32) *DescribeBackupPlanConfigResponseBody {
	s.FullBackupCycle = &v
	return s
}

func (s *DescribeBackupPlanConfigResponseBody) SetNextFullBackupDate(v string) *DescribeBackupPlanConfigResponseBody {
	s.NextFullBackupDate = &v
	return s
}

func (s *DescribeBackupPlanConfigResponseBody) SetTables(v *DescribeBackupPlanConfigResponseBodyTables) *DescribeBackupPlanConfigResponseBody {
	s.Tables = v
	return s
}

func (s *DescribeBackupPlanConfigResponseBody) SetMinHFileBackupCount(v int32) *DescribeBackupPlanConfigResponseBody {
	s.MinHFileBackupCount = &v
	return s
}

type DescribeBackupPlanConfigResponseBodyTables struct {
	Table []*string `json:"Table,omitempty" xml:"Table,omitempty" type:"Repeated"`
}

func (s DescribeBackupPlanConfigResponseBodyTables) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPlanConfigResponseBodyTables) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlanConfigResponseBodyTables) SetTable(v []*string) *DescribeBackupPlanConfigResponseBodyTables {
	s.Table = v
	return s
}

type DescribeBackupPlanConfigResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBackupPlanConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBackupPlanConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPlanConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlanConfigResponse) SetHeaders(v map[string]*string) *DescribeBackupPlanConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupPlanConfigResponse) SetBody(v *DescribeBackupPlanConfigResponseBody) *DescribeBackupPlanConfigResponse {
	s.Body = v
	return s
}

type DescribeBackupPolicyRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
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

type DescribeBackupPolicyResponseBody struct {
	PreferredBackupPeriod       *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	PreferredBackupEndTimeUTC   *string `json:"PreferredBackupEndTimeUTC,omitempty" xml:"PreferredBackupEndTimeUTC,omitempty"`
	PreferredBackupStartTimeUTC *string `json:"PreferredBackupStartTimeUTC,omitempty" xml:"PreferredBackupStartTimeUTC,omitempty"`
	RequestId                   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PreferredBackupTime         *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	BackupRetentionPeriod       *string `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
}

func (s DescribeBackupPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBody) SetPreferredBackupPeriod(v string) *DescribeBackupPolicyResponseBody {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetPreferredBackupEndTimeUTC(v string) *DescribeBackupPolicyResponseBody {
	s.PreferredBackupEndTimeUTC = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetPreferredBackupStartTimeUTC(v string) *DescribeBackupPolicyResponseBody {
	s.PreferredBackupStartTimeUTC = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetRequestId(v string) *DescribeBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetPreferredBackupTime(v string) *DescribeBackupPolicyResponseBody {
	s.PreferredBackupTime = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetBackupRetentionPeriod(v string) *DescribeBackupPolicyResponseBody {
	s.BackupRetentionPeriod = &v
	return s
}

type DescribeBackupPolicyResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeBackupPolicyResponse) SetBody(v *DescribeBackupPolicyResponseBody) *DescribeBackupPolicyResponse {
	s.Body = v
	return s
}

type DescribeBackupsRequest struct {
	ClusterId    *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	BackupId     *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	PageNumber   *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTimeUTC *string `json:"StartTimeUTC,omitempty" xml:"StartTimeUTC,omitempty"`
	EndTimeUTC   *string `json:"EndTimeUTC,omitempty" xml:"EndTimeUTC,omitempty"`
}

func (s DescribeBackupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupsRequest) SetClusterId(v string) *DescribeBackupsRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeBackupsRequest) SetBackupId(v string) *DescribeBackupsRequest {
	s.BackupId = &v
	return s
}

func (s *DescribeBackupsRequest) SetPageNumber(v string) *DescribeBackupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupsRequest) SetPageSize(v string) *DescribeBackupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupsRequest) SetStartTime(v string) *DescribeBackupsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeBackupsRequest) SetEndTime(v string) *DescribeBackupsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeBackupsRequest) SetStartTimeUTC(v string) *DescribeBackupsRequest {
	s.StartTimeUTC = &v
	return s
}

func (s *DescribeBackupsRequest) SetEndTimeUTC(v string) *DescribeBackupsRequest {
	s.EndTimeUTC = &v
	return s
}

type DescribeBackupsResponseBody struct {
	TotalCount   *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize     *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId    *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber   *int32                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	EnableStatus *string                             `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	Backups      *DescribeBackupsResponseBodyBackups `json:"Backups,omitempty" xml:"Backups,omitempty" type:"Struct"`
}

func (s DescribeBackupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBody) SetTotalCount(v int32) *DescribeBackupsResponseBody {
	s.TotalCount = &v
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

func (s *DescribeBackupsResponseBody) SetPageNumber(v int32) *DescribeBackupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetEnableStatus(v string) *DescribeBackupsResponseBody {
	s.EnableStatus = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetBackups(v *DescribeBackupsResponseBodyBackups) *DescribeBackupsResponseBody {
	s.Backups = v
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
	BackupStatus       *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	BackupType         *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	BackupStartTime    *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	BackupDownloadURL  *string `json:"BackupDownloadURL,omitempty" xml:"BackupDownloadURL,omitempty"`
	BackupStartTimeUTC *string `json:"BackupStartTimeUTC,omitempty" xml:"BackupStartTimeUTC,omitempty"`
	BackupEndTime      *string `json:"BackupEndTime,omitempty" xml:"BackupEndTime,omitempty"`
	BackupId           *int32  `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	BackupDBNames      *string `json:"BackupDBNames,omitempty" xml:"BackupDBNames,omitempty"`
	BackupEndTimeUTC   *string `json:"BackupEndTimeUTC,omitempty" xml:"BackupEndTimeUTC,omitempty"`
	BackupSize         *string `json:"BackupSize,omitempty" xml:"BackupSize,omitempty"`
	BackupMode         *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	BackupMethod       *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
}

func (s DescribeBackupsResponseBodyBackupsBackup) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsResponseBodyBackupsBackup) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupStatus(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupStatus = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupType(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupStartTime(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupStartTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupDownloadURL(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupDownloadURL = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupStartTimeUTC(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupStartTimeUTC = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupEndTime(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupEndTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupId(v int32) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupId = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupDBNames(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupDBNames = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupEndTimeUTC(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupEndTimeUTC = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupSize(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupSize = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupMode(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupMode = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupMethod(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupMethod = &v
	return s
}

type DescribeBackupsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBackupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeBackupsResponse) SetBody(v *DescribeBackupsResponseBody) *DescribeBackupsResponse {
	s.Body = v
	return s
}

type DescribeBackupStatusRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeBackupStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupStatusRequest) SetClusterId(v string) *DescribeBackupStatusRequest {
	s.ClusterId = &v
	return s
}

type DescribeBackupStatusResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	BdsClusterId *string `json:"BdsClusterId,omitempty" xml:"BdsClusterId,omitempty"`
	ClusterId    *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	BackupStatus *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
}

func (s DescribeBackupStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupStatusResponseBody) SetRequestId(v string) *DescribeBackupStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupStatusResponseBody) SetBdsClusterId(v string) *DescribeBackupStatusResponseBody {
	s.BdsClusterId = &v
	return s
}

func (s *DescribeBackupStatusResponseBody) SetClusterId(v string) *DescribeBackupStatusResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeBackupStatusResponseBody) SetBackupStatus(v string) *DescribeBackupStatusResponseBody {
	s.BackupStatus = &v
	return s
}

type DescribeBackupStatusResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBackupStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBackupStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupStatusResponse) SetHeaders(v map[string]*string) *DescribeBackupStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupStatusResponse) SetBody(v *DescribeBackupStatusResponseBody) *DescribeBackupStatusResponse {
	s.Body = v
	return s
}

type DescribeBackupSummaryRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeBackupSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupSummaryRequest) SetClusterId(v string) *DescribeBackupSummaryRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeBackupSummaryRequest) SetPageNumber(v int32) *DescribeBackupSummaryRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupSummaryRequest) SetPageSize(v int32) *DescribeBackupSummaryRequest {
	s.PageSize = &v
	return s
}

type DescribeBackupSummaryResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Incr      *DescribeBackupSummaryResponseBodyIncr `json:"Incr,omitempty" xml:"Incr,omitempty" type:"Struct"`
	Full      *DescribeBackupSummaryResponseBodyFull `json:"Full,omitempty" xml:"Full,omitempty" type:"Struct"`
}

func (s DescribeBackupSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupSummaryResponseBody) SetRequestId(v string) *DescribeBackupSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupSummaryResponseBody) SetIncr(v *DescribeBackupSummaryResponseBodyIncr) *DescribeBackupSummaryResponseBody {
	s.Incr = v
	return s
}

func (s *DescribeBackupSummaryResponseBody) SetFull(v *DescribeBackupSummaryResponseBodyFull) *DescribeBackupSummaryResponseBody {
	s.Full = v
	return s
}

type DescribeBackupSummaryResponseBodyIncr struct {
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Speed         *string `json:"Speed,omitempty" xml:"Speed,omitempty"`
	Pos           *string `json:"Pos,omitempty" xml:"Pos,omitempty"`
	QueueLogNum   *string `json:"QueueLogNum,omitempty" xml:"QueueLogNum,omitempty"`
	BackupLogSize *string `json:"BackupLogSize,omitempty" xml:"BackupLogSize,omitempty"`
	RunningLogNum *string `json:"RunningLogNum,omitempty" xml:"RunningLogNum,omitempty"`
}

func (s DescribeBackupSummaryResponseBodyIncr) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupSummaryResponseBodyIncr) GoString() string {
	return s.String()
}

func (s *DescribeBackupSummaryResponseBodyIncr) SetStatus(v string) *DescribeBackupSummaryResponseBodyIncr {
	s.Status = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyIncr) SetSpeed(v string) *DescribeBackupSummaryResponseBodyIncr {
	s.Speed = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyIncr) SetPos(v string) *DescribeBackupSummaryResponseBodyIncr {
	s.Pos = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyIncr) SetQueueLogNum(v string) *DescribeBackupSummaryResponseBodyIncr {
	s.QueueLogNum = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyIncr) SetBackupLogSize(v string) *DescribeBackupSummaryResponseBodyIncr {
	s.BackupLogSize = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyIncr) SetRunningLogNum(v string) *DescribeBackupSummaryResponseBodyIncr {
	s.RunningLogNum = &v
	return s
}

type DescribeBackupSummaryResponseBodyFull struct {
	NextFullBackupDate *string                                       `json:"NextFullBackupDate,omitempty" xml:"NextFullBackupDate,omitempty"`
	Records            *DescribeBackupSummaryResponseBodyFullRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Struct"`
	HasMore            *string                                       `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	PageSize           *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber         *int32                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Total              *int32                                        `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeBackupSummaryResponseBodyFull) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupSummaryResponseBodyFull) GoString() string {
	return s.String()
}

func (s *DescribeBackupSummaryResponseBodyFull) SetNextFullBackupDate(v string) *DescribeBackupSummaryResponseBodyFull {
	s.NextFullBackupDate = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyFull) SetRecords(v *DescribeBackupSummaryResponseBodyFullRecords) *DescribeBackupSummaryResponseBodyFull {
	s.Records = v
	return s
}

func (s *DescribeBackupSummaryResponseBodyFull) SetHasMore(v string) *DescribeBackupSummaryResponseBodyFull {
	s.HasMore = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyFull) SetPageSize(v int32) *DescribeBackupSummaryResponseBodyFull {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyFull) SetPageNumber(v int32) *DescribeBackupSummaryResponseBodyFull {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyFull) SetTotal(v int32) *DescribeBackupSummaryResponseBodyFull {
	s.Total = &v
	return s
}

type DescribeBackupSummaryResponseBodyFullRecords struct {
	Record []*DescribeBackupSummaryResponseBodyFullRecordsRecord `json:"Record,omitempty" xml:"Record,omitempty" type:"Repeated"`
}

func (s DescribeBackupSummaryResponseBodyFullRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupSummaryResponseBodyFullRecords) GoString() string {
	return s.String()
}

func (s *DescribeBackupSummaryResponseBodyFullRecords) SetRecord(v []*DescribeBackupSummaryResponseBodyFullRecordsRecord) *DescribeBackupSummaryResponseBodyFullRecords {
	s.Record = v
	return s
}

type DescribeBackupSummaryResponseBodyFullRecordsRecord struct {
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	Process    *string `json:"Process,omitempty" xml:"Process,omitempty"`
	DataSize   *string `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	Speed      *string `json:"Speed,omitempty" xml:"Speed,omitempty"`
	RecordId   *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
}

func (s DescribeBackupSummaryResponseBodyFullRecordsRecord) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupSummaryResponseBodyFullRecordsRecord) GoString() string {
	return s.String()
}

func (s *DescribeBackupSummaryResponseBodyFullRecordsRecord) SetStatus(v string) *DescribeBackupSummaryResponseBodyFullRecordsRecord {
	s.Status = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyFullRecordsRecord) SetFinishTime(v string) *DescribeBackupSummaryResponseBodyFullRecordsRecord {
	s.FinishTime = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyFullRecordsRecord) SetProcess(v string) *DescribeBackupSummaryResponseBodyFullRecordsRecord {
	s.Process = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyFullRecordsRecord) SetDataSize(v string) *DescribeBackupSummaryResponseBodyFullRecordsRecord {
	s.DataSize = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyFullRecordsRecord) SetSpeed(v string) *DescribeBackupSummaryResponseBodyFullRecordsRecord {
	s.Speed = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyFullRecordsRecord) SetRecordId(v string) *DescribeBackupSummaryResponseBodyFullRecordsRecord {
	s.RecordId = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyFullRecordsRecord) SetCreateTime(v string) *DescribeBackupSummaryResponseBodyFullRecordsRecord {
	s.CreateTime = &v
	return s
}

type DescribeBackupSummaryResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBackupSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBackupSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupSummaryResponse) SetHeaders(v map[string]*string) *DescribeBackupSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupSummaryResponse) SetBody(v *DescribeBackupSummaryResponseBody) *DescribeBackupSummaryResponse {
	s.Body = v
	return s
}

type DescribeBackupTablesRequest struct {
	ClusterId      *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	BackupRecordId *string `json:"BackupRecordId,omitempty" xml:"BackupRecordId,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeBackupTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupTablesRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupTablesRequest) SetClusterId(v string) *DescribeBackupTablesRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeBackupTablesRequest) SetBackupRecordId(v string) *DescribeBackupTablesRequest {
	s.BackupRecordId = &v
	return s
}

func (s *DescribeBackupTablesRequest) SetPageNumber(v int32) *DescribeBackupTablesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupTablesRequest) SetPageSize(v int32) *DescribeBackupTablesRequest {
	s.PageSize = &v
	return s
}

type DescribeBackupTablesResponseBody struct {
	RequestId     *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize      *int32                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	BackupRecords *DescribeBackupTablesResponseBodyBackupRecords `json:"BackupRecords,omitempty" xml:"BackupRecords,omitempty" type:"Struct"`
	PageNumber    *int32                                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Total         *int64                                         `json:"Total,omitempty" xml:"Total,omitempty"`
	Tables        *DescribeBackupTablesResponseBodyTables        `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Struct"`
}

func (s DescribeBackupTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupTablesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupTablesResponseBody) SetRequestId(v string) *DescribeBackupTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupTablesResponseBody) SetPageSize(v int32) *DescribeBackupTablesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupTablesResponseBody) SetBackupRecords(v *DescribeBackupTablesResponseBodyBackupRecords) *DescribeBackupTablesResponseBody {
	s.BackupRecords = v
	return s
}

func (s *DescribeBackupTablesResponseBody) SetPageNumber(v int32) *DescribeBackupTablesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupTablesResponseBody) SetTotal(v int64) *DescribeBackupTablesResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeBackupTablesResponseBody) SetTables(v *DescribeBackupTablesResponseBodyTables) *DescribeBackupTablesResponseBody {
	s.Tables = v
	return s
}

type DescribeBackupTablesResponseBodyBackupRecords struct {
	BackupRecord []*DescribeBackupTablesResponseBodyBackupRecordsBackupRecord `json:"BackupRecord,omitempty" xml:"BackupRecord,omitempty" type:"Repeated"`
}

func (s DescribeBackupTablesResponseBodyBackupRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupTablesResponseBodyBackupRecords) GoString() string {
	return s.String()
}

func (s *DescribeBackupTablesResponseBodyBackupRecords) SetBackupRecord(v []*DescribeBackupTablesResponseBodyBackupRecordsBackupRecord) *DescribeBackupTablesResponseBodyBackupRecords {
	s.BackupRecord = v
	return s
}

type DescribeBackupTablesResponseBodyBackupRecordsBackupRecord struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Process   *string `json:"Process,omitempty" xml:"Process,omitempty"`
	DataSize  *string `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	Speed     *string `json:"Speed,omitempty" xml:"Speed,omitempty"`
	State     *string `json:"State,omitempty" xml:"State,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Table     *string `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s DescribeBackupTablesResponseBodyBackupRecordsBackupRecord) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupTablesResponseBodyBackupRecordsBackupRecord) GoString() string {
	return s.String()
}

func (s *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord) SetEndTime(v string) *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord {
	s.EndTime = &v
	return s
}

func (s *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord) SetStartTime(v string) *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord {
	s.StartTime = &v
	return s
}

func (s *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord) SetProcess(v string) *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord {
	s.Process = &v
	return s
}

func (s *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord) SetDataSize(v string) *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord {
	s.DataSize = &v
	return s
}

func (s *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord) SetSpeed(v string) *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord {
	s.Speed = &v
	return s
}

func (s *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord) SetState(v string) *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord {
	s.State = &v
	return s
}

func (s *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord) SetMessage(v string) *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord {
	s.Message = &v
	return s
}

func (s *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord) SetTable(v string) *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord {
	s.Table = &v
	return s
}

type DescribeBackupTablesResponseBodyTables struct {
	Table []*string `json:"Table,omitempty" xml:"Table,omitempty" type:"Repeated"`
}

func (s DescribeBackupTablesResponseBodyTables) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupTablesResponseBodyTables) GoString() string {
	return s.String()
}

func (s *DescribeBackupTablesResponseBodyTables) SetTable(v []*string) *DescribeBackupTablesResponseBodyTables {
	s.Table = v
	return s
}

type DescribeBackupTablesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBackupTablesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBackupTablesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupTablesResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupTablesResponse) SetHeaders(v map[string]*string) *DescribeBackupTablesResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupTablesResponse) SetBody(v *DescribeBackupTablesResponseBody) *DescribeBackupTablesResponse {
	s.Body = v
	return s
}

type DescribeClusterConnectionRequest struct {
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeClusterConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterConnectionRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectionRequest) SetRegionId(v string) *DescribeClusterConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeClusterConnectionRequest) SetClusterId(v string) *DescribeClusterConnectionRequest {
	s.ClusterId = &v
	return s
}

type DescribeClusterConnectionResponseBody struct {
	IsMultimod          *string                                                   `json:"IsMultimod,omitempty" xml:"IsMultimod,omitempty"`
	VpcId               *string                                                   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	RequestId           *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UiProxyConnAddrInfo *DescribeClusterConnectionResponseBodyUiProxyConnAddrInfo `json:"UiProxyConnAddrInfo,omitempty" xml:"UiProxyConnAddrInfo,omitempty" type:"Struct"`
	VSwitchId           *string                                                   `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ServiceConnAddrs    *DescribeClusterConnectionResponseBodyServiceConnAddrs    `json:"ServiceConnAddrs,omitempty" xml:"ServiceConnAddrs,omitempty" type:"Struct"`
	NetType             *string                                                   `json:"NetType,omitempty" xml:"NetType,omitempty"`
	DbType              *string                                                   `json:"DbType,omitempty" xml:"DbType,omitempty"`
	ThriftConn          *DescribeClusterConnectionResponseBodyThriftConn          `json:"ThriftConn,omitempty" xml:"ThriftConn,omitempty" type:"Struct"`
	ZkConnAddrs         *DescribeClusterConnectionResponseBodyZkConnAddrs         `json:"ZkConnAddrs,omitempty" xml:"ZkConnAddrs,omitempty" type:"Struct"`
	SlbConnAddrs        *DescribeClusterConnectionResponseBodySlbConnAddrs        `json:"SlbConnAddrs,omitempty" xml:"SlbConnAddrs,omitempty" type:"Struct"`
}

func (s DescribeClusterConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectionResponseBody) SetIsMultimod(v string) *DescribeClusterConnectionResponseBody {
	s.IsMultimod = &v
	return s
}

func (s *DescribeClusterConnectionResponseBody) SetVpcId(v string) *DescribeClusterConnectionResponseBody {
	s.VpcId = &v
	return s
}

func (s *DescribeClusterConnectionResponseBody) SetRequestId(v string) *DescribeClusterConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterConnectionResponseBody) SetUiProxyConnAddrInfo(v *DescribeClusterConnectionResponseBodyUiProxyConnAddrInfo) *DescribeClusterConnectionResponseBody {
	s.UiProxyConnAddrInfo = v
	return s
}

func (s *DescribeClusterConnectionResponseBody) SetVSwitchId(v string) *DescribeClusterConnectionResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *DescribeClusterConnectionResponseBody) SetServiceConnAddrs(v *DescribeClusterConnectionResponseBodyServiceConnAddrs) *DescribeClusterConnectionResponseBody {
	s.ServiceConnAddrs = v
	return s
}

func (s *DescribeClusterConnectionResponseBody) SetNetType(v string) *DescribeClusterConnectionResponseBody {
	s.NetType = &v
	return s
}

func (s *DescribeClusterConnectionResponseBody) SetDbType(v string) *DescribeClusterConnectionResponseBody {
	s.DbType = &v
	return s
}

func (s *DescribeClusterConnectionResponseBody) SetThriftConn(v *DescribeClusterConnectionResponseBodyThriftConn) *DescribeClusterConnectionResponseBody {
	s.ThriftConn = v
	return s
}

func (s *DescribeClusterConnectionResponseBody) SetZkConnAddrs(v *DescribeClusterConnectionResponseBodyZkConnAddrs) *DescribeClusterConnectionResponseBody {
	s.ZkConnAddrs = v
	return s
}

func (s *DescribeClusterConnectionResponseBody) SetSlbConnAddrs(v *DescribeClusterConnectionResponseBodySlbConnAddrs) *DescribeClusterConnectionResponseBody {
	s.SlbConnAddrs = v
	return s
}

type DescribeClusterConnectionResponseBodyUiProxyConnAddrInfo struct {
	ConnAddrPort *string `json:"ConnAddrPort,omitempty" xml:"ConnAddrPort,omitempty"`
	NetType      *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	ConnAddr     *string `json:"ConnAddr,omitempty" xml:"ConnAddr,omitempty"`
}

func (s DescribeClusterConnectionResponseBodyUiProxyConnAddrInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterConnectionResponseBodyUiProxyConnAddrInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectionResponseBodyUiProxyConnAddrInfo) SetConnAddrPort(v string) *DescribeClusterConnectionResponseBodyUiProxyConnAddrInfo {
	s.ConnAddrPort = &v
	return s
}

func (s *DescribeClusterConnectionResponseBodyUiProxyConnAddrInfo) SetNetType(v string) *DescribeClusterConnectionResponseBodyUiProxyConnAddrInfo {
	s.NetType = &v
	return s
}

func (s *DescribeClusterConnectionResponseBodyUiProxyConnAddrInfo) SetConnAddr(v string) *DescribeClusterConnectionResponseBodyUiProxyConnAddrInfo {
	s.ConnAddr = &v
	return s
}

type DescribeClusterConnectionResponseBodyServiceConnAddrs struct {
	ServiceConnAddr []*DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddr `json:"ServiceConnAddr,omitempty" xml:"ServiceConnAddr,omitempty" type:"Repeated"`
}

func (s DescribeClusterConnectionResponseBodyServiceConnAddrs) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterConnectionResponseBodyServiceConnAddrs) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectionResponseBodyServiceConnAddrs) SetServiceConnAddr(v []*DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddr) *DescribeClusterConnectionResponseBodyServiceConnAddrs {
	s.ServiceConnAddr = v
	return s
}

type DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddr struct {
	ConnAddrInfo *DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo `json:"ConnAddrInfo,omitempty" xml:"ConnAddrInfo,omitempty" type:"Struct"`
	ConnType     *string                                                                           `json:"ConnType,omitempty" xml:"ConnType,omitempty"`
}

func (s DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddr) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddr) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddr) SetConnAddrInfo(v *DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo) *DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddr {
	s.ConnAddrInfo = v
	return s
}

func (s *DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddr) SetConnType(v string) *DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddr {
	s.ConnType = &v
	return s
}

type DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo struct {
	ConnAddrPort *string `json:"ConnAddrPort,omitempty" xml:"ConnAddrPort,omitempty"`
	NetType      *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	ConnAddr     *string `json:"ConnAddr,omitempty" xml:"ConnAddr,omitempty"`
}

func (s DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo) SetConnAddrPort(v string) *DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo {
	s.ConnAddrPort = &v
	return s
}

func (s *DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo) SetNetType(v string) *DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo {
	s.NetType = &v
	return s
}

func (s *DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo) SetConnAddr(v string) *DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo {
	s.ConnAddr = &v
	return s
}

type DescribeClusterConnectionResponseBodyThriftConn struct {
	ConnAddrPort *string `json:"ConnAddrPort,omitempty" xml:"ConnAddrPort,omitempty"`
	NetType      *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	ConnAddr     *string `json:"ConnAddr,omitempty" xml:"ConnAddr,omitempty"`
}

func (s DescribeClusterConnectionResponseBodyThriftConn) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterConnectionResponseBodyThriftConn) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectionResponseBodyThriftConn) SetConnAddrPort(v string) *DescribeClusterConnectionResponseBodyThriftConn {
	s.ConnAddrPort = &v
	return s
}

func (s *DescribeClusterConnectionResponseBodyThriftConn) SetNetType(v string) *DescribeClusterConnectionResponseBodyThriftConn {
	s.NetType = &v
	return s
}

func (s *DescribeClusterConnectionResponseBodyThriftConn) SetConnAddr(v string) *DescribeClusterConnectionResponseBodyThriftConn {
	s.ConnAddr = &v
	return s
}

type DescribeClusterConnectionResponseBodyZkConnAddrs struct {
	ZkConnAddr []*DescribeClusterConnectionResponseBodyZkConnAddrsZkConnAddr `json:"ZkConnAddr,omitempty" xml:"ZkConnAddr,omitempty" type:"Repeated"`
}

func (s DescribeClusterConnectionResponseBodyZkConnAddrs) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterConnectionResponseBodyZkConnAddrs) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectionResponseBodyZkConnAddrs) SetZkConnAddr(v []*DescribeClusterConnectionResponseBodyZkConnAddrsZkConnAddr) *DescribeClusterConnectionResponseBodyZkConnAddrs {
	s.ZkConnAddr = v
	return s
}

type DescribeClusterConnectionResponseBodyZkConnAddrsZkConnAddr struct {
	ConnAddrPort *string `json:"ConnAddrPort,omitempty" xml:"ConnAddrPort,omitempty"`
	NetType      *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	ConnAddr     *string `json:"ConnAddr,omitempty" xml:"ConnAddr,omitempty"`
}

func (s DescribeClusterConnectionResponseBodyZkConnAddrsZkConnAddr) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterConnectionResponseBodyZkConnAddrsZkConnAddr) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectionResponseBodyZkConnAddrsZkConnAddr) SetConnAddrPort(v string) *DescribeClusterConnectionResponseBodyZkConnAddrsZkConnAddr {
	s.ConnAddrPort = &v
	return s
}

func (s *DescribeClusterConnectionResponseBodyZkConnAddrsZkConnAddr) SetNetType(v string) *DescribeClusterConnectionResponseBodyZkConnAddrsZkConnAddr {
	s.NetType = &v
	return s
}

func (s *DescribeClusterConnectionResponseBodyZkConnAddrsZkConnAddr) SetConnAddr(v string) *DescribeClusterConnectionResponseBodyZkConnAddrsZkConnAddr {
	s.ConnAddr = &v
	return s
}

type DescribeClusterConnectionResponseBodySlbConnAddrs struct {
	SlbConnAddr []*DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddr `json:"SlbConnAddr,omitempty" xml:"SlbConnAddr,omitempty" type:"Repeated"`
}

func (s DescribeClusterConnectionResponseBodySlbConnAddrs) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterConnectionResponseBodySlbConnAddrs) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectionResponseBodySlbConnAddrs) SetSlbConnAddr(v []*DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddr) *DescribeClusterConnectionResponseBodySlbConnAddrs {
	s.SlbConnAddr = v
	return s
}

type DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddr struct {
	ConnAddrInfo *DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo `json:"ConnAddrInfo,omitempty" xml:"ConnAddrInfo,omitempty" type:"Struct"`
	SlbType      *string                                                                   `json:"SlbType,omitempty" xml:"SlbType,omitempty"`
}

func (s DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddr) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddr) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddr) SetConnAddrInfo(v *DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo) *DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddr {
	s.ConnAddrInfo = v
	return s
}

func (s *DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddr) SetSlbType(v string) *DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddr {
	s.SlbType = &v
	return s
}

type DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo struct {
	ConnAddrPort *string `json:"ConnAddrPort,omitempty" xml:"ConnAddrPort,omitempty"`
	NetType      *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	ConnAddr     *string `json:"ConnAddr,omitempty" xml:"ConnAddr,omitempty"`
}

func (s DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo) SetConnAddrPort(v string) *DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo {
	s.ConnAddrPort = &v
	return s
}

func (s *DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo) SetNetType(v string) *DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo {
	s.NetType = &v
	return s
}

func (s *DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo) SetConnAddr(v string) *DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo {
	s.ConnAddr = &v
	return s
}

type DescribeClusterConnectionResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeClusterConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeClusterConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterConnectionResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectionResponse) SetHeaders(v map[string]*string) *DescribeClusterConnectionResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterConnectionResponse) SetBody(v *DescribeClusterConnectionResponseBody) *DescribeClusterConnectionResponse {
	s.Body = v
	return s
}

type DescribeColdStorageRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
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

type DescribeColdStorageResponseBody struct {
	ColdStorageUsePercent *string `json:"ColdStorageUsePercent,omitempty" xml:"ColdStorageUsePercent,omitempty"`
	ColdStorageSize       *string `json:"ColdStorageSize,omitempty" xml:"ColdStorageSize,omitempty"`
	RequestId             *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ClusterId             *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PayType               *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	OpenStatus            *string `json:"OpenStatus,omitempty" xml:"OpenStatus,omitempty"`
}

func (s DescribeColdStorageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeColdStorageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeColdStorageResponseBody) SetColdStorageUsePercent(v string) *DescribeColdStorageResponseBody {
	s.ColdStorageUsePercent = &v
	return s
}

func (s *DescribeColdStorageResponseBody) SetColdStorageSize(v string) *DescribeColdStorageResponseBody {
	s.ColdStorageSize = &v
	return s
}

func (s *DescribeColdStorageResponseBody) SetRequestId(v string) *DescribeColdStorageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeColdStorageResponseBody) SetClusterId(v string) *DescribeColdStorageResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeColdStorageResponseBody) SetPayType(v string) *DescribeColdStorageResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeColdStorageResponseBody) SetOpenStatus(v string) *DescribeColdStorageResponseBody {
	s.OpenStatus = &v
	return s
}

type DescribeColdStorageResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeColdStorageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeColdStorageResponse) SetBody(v *DescribeColdStorageResponseBody) *DescribeColdStorageResponse {
	s.Body = v
	return s
}

type DescribeDBInstanceUsageRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeDBInstanceUsageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceUsageRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceUsageRequest) SetClusterId(v string) *DescribeDBInstanceUsageRequest {
	s.ClusterId = &v
	return s
}

type DescribeDBInstanceUsageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DescribeDBInstanceUsageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceUsageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceUsageResponseBody) SetRequestId(v string) *DescribeDBInstanceUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceUsageResponseBody) SetResult(v string) *DescribeDBInstanceUsageResponseBody {
	s.Result = &v
	return s
}

type DescribeDBInstanceUsageResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBInstanceUsageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBInstanceUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceUsageResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceUsageResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceUsageResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceUsageResponse) SetBody(v *DescribeDBInstanceUsageResponseBody) *DescribeDBInstanceUsageResponse {
	s.Body = v
	return s
}

type DescribeDeletedInstancesRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDeletedInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeletedInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeletedInstancesRequest) SetRegionId(v string) *DescribeDeletedInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDeletedInstancesRequest) SetPageNumber(v int32) *DescribeDeletedInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDeletedInstancesRequest) SetPageSize(v int32) *DescribeDeletedInstancesRequest {
	s.PageSize = &v
	return s
}

type DescribeDeletedInstancesResponseBody struct {
	Instances  *DescribeDeletedInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	TotalCount *int64                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int32                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeDeletedInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeletedInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeletedInstancesResponseBody) SetInstances(v *DescribeDeletedInstancesResponseBodyInstances) *DescribeDeletedInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeDeletedInstancesResponseBody) SetTotalCount(v int64) *DescribeDeletedInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDeletedInstancesResponseBody) SetPageSize(v int32) *DescribeDeletedInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDeletedInstancesResponseBody) SetRequestId(v string) *DescribeDeletedInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeletedInstancesResponseBody) SetPageNumber(v int32) *DescribeDeletedInstancesResponseBody {
	s.PageNumber = &v
	return s
}

type DescribeDeletedInstancesResponseBodyInstances struct {
	Instance []*DescribeDeletedInstancesResponseBodyInstancesInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s DescribeDeletedInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeletedInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeDeletedInstancesResponseBodyInstances) SetInstance(v []*DescribeDeletedInstancesResponseBodyInstancesInstance) *DescribeDeletedInstancesResponseBodyInstances {
	s.Instance = v
	return s
}

type DescribeDeletedInstancesResponseBodyInstancesInstance struct {
	Status             *string `json:"Status,omitempty" xml:"Status,omitempty"`
	InstanceId         *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ParentId           *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	ClusterType        *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	InstanceName       *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	DeleteTime         *string `json:"DeleteTime,omitempty" xml:"DeleteTime,omitempty"`
	ZoneId             *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ModuleStackVersion *string `json:"ModuleStackVersion,omitempty" xml:"ModuleStackVersion,omitempty"`
	Engine             *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	MajorVersion       *string `json:"MajorVersion,omitempty" xml:"MajorVersion,omitempty"`
	CreatedTime        *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
}

func (s DescribeDeletedInstancesResponseBodyInstancesInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeletedInstancesResponseBodyInstancesInstance) GoString() string {
	return s.String()
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) SetStatus(v string) *DescribeDeletedInstancesResponseBodyInstancesInstance {
	s.Status = &v
	return s
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) SetInstanceId(v string) *DescribeDeletedInstancesResponseBodyInstancesInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) SetRegionId(v string) *DescribeDeletedInstancesResponseBodyInstancesInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) SetParentId(v string) *DescribeDeletedInstancesResponseBodyInstancesInstance {
	s.ParentId = &v
	return s
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) SetClusterType(v string) *DescribeDeletedInstancesResponseBodyInstancesInstance {
	s.ClusterType = &v
	return s
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) SetInstanceName(v string) *DescribeDeletedInstancesResponseBodyInstancesInstance {
	s.InstanceName = &v
	return s
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) SetDeleteTime(v string) *DescribeDeletedInstancesResponseBodyInstancesInstance {
	s.DeleteTime = &v
	return s
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) SetZoneId(v string) *DescribeDeletedInstancesResponseBodyInstancesInstance {
	s.ZoneId = &v
	return s
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) SetModuleStackVersion(v string) *DescribeDeletedInstancesResponseBodyInstancesInstance {
	s.ModuleStackVersion = &v
	return s
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) SetEngine(v string) *DescribeDeletedInstancesResponseBodyInstancesInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) SetMajorVersion(v string) *DescribeDeletedInstancesResponseBodyInstancesInstance {
	s.MajorVersion = &v
	return s
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) SetCreatedTime(v string) *DescribeDeletedInstancesResponseBodyInstancesInstance {
	s.CreatedTime = &v
	return s
}

type DescribeDeletedInstancesResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDeletedInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDeletedInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeletedInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeletedInstancesResponse) SetHeaders(v map[string]*string) *DescribeDeletedInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeletedInstancesResponse) SetBody(v *DescribeDeletedInstancesResponseBody) *DescribeDeletedInstancesResponse {
	s.Body = v
	return s
}

type DescribeDiskWarningLineRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeDiskWarningLineRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiskWarningLineRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiskWarningLineRequest) SetClusterId(v string) *DescribeDiskWarningLineRequest {
	s.ClusterId = &v
	return s
}

type DescribeDiskWarningLineResponseBody struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WarningLine *string `json:"WarningLine,omitempty" xml:"WarningLine,omitempty"`
}

func (s DescribeDiskWarningLineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiskWarningLineResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiskWarningLineResponseBody) SetRequestId(v string) *DescribeDiskWarningLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiskWarningLineResponseBody) SetWarningLine(v string) *DescribeDiskWarningLineResponseBody {
	s.WarningLine = &v
	return s
}

type DescribeDiskWarningLineResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDiskWarningLineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDiskWarningLineResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiskWarningLineResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiskWarningLineResponse) SetHeaders(v map[string]*string) *DescribeDiskWarningLineResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiskWarningLineResponse) SetBody(v *DescribeDiskWarningLineResponseBody) *DescribeDiskWarningLineResponse {
	s.Body = v
	return s
}

type DescribeEndpointsRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeEndpointsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndpointsRequest) GoString() string {
	return s.String()
}

func (s *DescribeEndpointsRequest) SetClusterId(v string) *DescribeEndpointsRequest {
	s.ClusterId = &v
	return s
}

type DescribeEndpointsResponseBody struct {
	VpcId     *string                                 `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ConnAddrs *DescribeEndpointsResponseBodyConnAddrs `json:"ConnAddrs,omitempty" xml:"ConnAddrs,omitempty" type:"Struct"`
	VSwitchId *string                                 `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	NetType   *string                                 `json:"NetType,omitempty" xml:"NetType,omitempty"`
	Engine    *string                                 `json:"Engine,omitempty" xml:"Engine,omitempty"`
}

func (s DescribeEndpointsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEndpointsResponseBody) SetVpcId(v string) *DescribeEndpointsResponseBody {
	s.VpcId = &v
	return s
}

func (s *DescribeEndpointsResponseBody) SetRequestId(v string) *DescribeEndpointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEndpointsResponseBody) SetConnAddrs(v *DescribeEndpointsResponseBodyConnAddrs) *DescribeEndpointsResponseBody {
	s.ConnAddrs = v
	return s
}

func (s *DescribeEndpointsResponseBody) SetVSwitchId(v string) *DescribeEndpointsResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *DescribeEndpointsResponseBody) SetNetType(v string) *DescribeEndpointsResponseBody {
	s.NetType = &v
	return s
}

func (s *DescribeEndpointsResponseBody) SetEngine(v string) *DescribeEndpointsResponseBody {
	s.Engine = &v
	return s
}

type DescribeEndpointsResponseBodyConnAddrs struct {
	ConnAddrInfo []*DescribeEndpointsResponseBodyConnAddrsConnAddrInfo `json:"ConnAddrInfo,omitempty" xml:"ConnAddrInfo,omitempty" type:"Repeated"`
}

func (s DescribeEndpointsResponseBodyConnAddrs) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndpointsResponseBodyConnAddrs) GoString() string {
	return s.String()
}

func (s *DescribeEndpointsResponseBodyConnAddrs) SetConnAddrInfo(v []*DescribeEndpointsResponseBodyConnAddrsConnAddrInfo) *DescribeEndpointsResponseBodyConnAddrs {
	s.ConnAddrInfo = v
	return s
}

type DescribeEndpointsResponseBodyConnAddrsConnAddrInfo struct {
	ConnAddrPort *string `json:"ConnAddrPort,omitempty" xml:"ConnAddrPort,omitempty"`
	NetType      *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	ConnAddr     *string `json:"ConnAddr,omitempty" xml:"ConnAddr,omitempty"`
	ConnType     *string `json:"ConnType,omitempty" xml:"ConnType,omitempty"`
}

func (s DescribeEndpointsResponseBodyConnAddrsConnAddrInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndpointsResponseBodyConnAddrsConnAddrInfo) GoString() string {
	return s.String()
}

func (s *DescribeEndpointsResponseBodyConnAddrsConnAddrInfo) SetConnAddrPort(v string) *DescribeEndpointsResponseBodyConnAddrsConnAddrInfo {
	s.ConnAddrPort = &v
	return s
}

func (s *DescribeEndpointsResponseBodyConnAddrsConnAddrInfo) SetNetType(v string) *DescribeEndpointsResponseBodyConnAddrsConnAddrInfo {
	s.NetType = &v
	return s
}

func (s *DescribeEndpointsResponseBodyConnAddrsConnAddrInfo) SetConnAddr(v string) *DescribeEndpointsResponseBodyConnAddrsConnAddrInfo {
	s.ConnAddr = &v
	return s
}

func (s *DescribeEndpointsResponseBodyConnAddrsConnAddrInfo) SetConnType(v string) *DescribeEndpointsResponseBodyConnAddrsConnAddrInfo {
	s.ConnType = &v
	return s
}

type DescribeEndpointsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEndpointsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndpointsResponse) GoString() string {
	return s.String()
}

func (s *DescribeEndpointsResponse) SetHeaders(v map[string]*string) *DescribeEndpointsResponse {
	s.Headers = v
	return s
}

func (s *DescribeEndpointsResponse) SetBody(v *DescribeEndpointsResponseBody) *DescribeEndpointsResponse {
	s.Body = v
	return s
}

type DescribeInstanceRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRequest) SetClusterId(v string) *DescribeInstanceRequest {
	s.ClusterId = &v
	return s
}

type DescribeInstanceResponseBody struct {
	ModuleStackVersion   *string                           `json:"ModuleStackVersion,omitempty" xml:"ModuleStackVersion,omitempty"`
	IsHa                 *bool                             `json:"IsHa,omitempty" xml:"IsHa,omitempty"`
	CreatedTime          *string                           `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	ResourceGroupId      *string                           `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	MasterInstanceType   *string                           `json:"MasterInstanceType,omitempty" xml:"MasterInstanceType,omitempty"`
	IsDeletionProtection *bool                             `json:"IsDeletionProtection,omitempty" xml:"IsDeletionProtection,omitempty"`
	ModuleId             *int32                            `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	IsLatestVersion      *bool                             `json:"IsLatestVersion,omitempty" xml:"IsLatestVersion,omitempty"`
	MaintainEndTime      *string                           `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	NetworkType          *string                           `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	CoreInstanceType     *string                           `json:"CoreInstanceType,omitempty" xml:"CoreInstanceType,omitempty"`
	ClusterName          *string                           `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	MasterDiskType       *string                           `json:"MasterDiskType,omitempty" xml:"MasterDiskType,omitempty"`
	MaintainStartTime    *string                           `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	Engine               *string                           `json:"Engine,omitempty" xml:"Engine,omitempty"`
	Tags                 *DescribeInstanceResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	Status               *string                           `json:"Status,omitempty" xml:"Status,omitempty"`
	ColdStorageSize      *int32                            `json:"ColdStorageSize,omitempty" xml:"ColdStorageSize,omitempty"`
	ParentId             *string                           `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	MajorVersion         *string                           `json:"MajorVersion,omitempty" xml:"MajorVersion,omitempty"`
	CoreDiskCount        *string                           `json:"CoreDiskCount,omitempty" xml:"CoreDiskCount,omitempty"`
	ExpireTimeUTC        *string                           `json:"ExpireTimeUTC,omitempty" xml:"ExpireTimeUTC,omitempty"`
	MasterDiskSize       *int32                            `json:"MasterDiskSize,omitempty" xml:"MasterDiskSize,omitempty"`
	RequestId            *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ZoneId               *string                           `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ClusterId            *string                           `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	InstanceId           *string                           `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	CoreNodeCount        *int32                            `json:"CoreNodeCount,omitempty" xml:"CoreNodeCount,omitempty"`
	ColdStorageStatus    *string                           `json:"ColdStorageStatus,omitempty" xml:"ColdStorageStatus,omitempty"`
	CreatedTimeUTC       *string                           `json:"CreatedTimeUTC,omitempty" xml:"CreatedTimeUTC,omitempty"`
	MinorVersion         *string                           `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	Duration             *int32                            `json:"Duration,omitempty" xml:"Duration,omitempty"`
	PayType              *string                           `json:"PayType,omitempty" xml:"PayType,omitempty"`
	IsMultiModel         *bool                             `json:"IsMultiModel,omitempty" xml:"IsMultiModel,omitempty"`
	ClusterType          *string                           `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	VswitchId            *string                           `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	MasterNodeCount      *int32                            `json:"MasterNodeCount,omitempty" xml:"MasterNodeCount,omitempty"`
	InstanceName         *string                           `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	VpcId                *string                           `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	AutoRenewal          *bool                             `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	CoreDiskType         *string                           `json:"CoreDiskType,omitempty" xml:"CoreDiskType,omitempty"`
	RegionId             *string                           `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ExpireTime           *string                           `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	BackupStatus         *string                           `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	CoreDiskSize         *int32                            `json:"CoreDiskSize,omitempty" xml:"CoreDiskSize,omitempty"`
}

func (s DescribeInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBody) SetModuleStackVersion(v string) *DescribeInstanceResponseBody {
	s.ModuleStackVersion = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetIsHa(v bool) *DescribeInstanceResponseBody {
	s.IsHa = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetCreatedTime(v string) *DescribeInstanceResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetResourceGroupId(v string) *DescribeInstanceResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetMasterInstanceType(v string) *DescribeInstanceResponseBody {
	s.MasterInstanceType = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetIsDeletionProtection(v bool) *DescribeInstanceResponseBody {
	s.IsDeletionProtection = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetModuleId(v int32) *DescribeInstanceResponseBody {
	s.ModuleId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetIsLatestVersion(v bool) *DescribeInstanceResponseBody {
	s.IsLatestVersion = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetMaintainEndTime(v string) *DescribeInstanceResponseBody {
	s.MaintainEndTime = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetNetworkType(v string) *DescribeInstanceResponseBody {
	s.NetworkType = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetCoreInstanceType(v string) *DescribeInstanceResponseBody {
	s.CoreInstanceType = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetClusterName(v string) *DescribeInstanceResponseBody {
	s.ClusterName = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetMasterDiskType(v string) *DescribeInstanceResponseBody {
	s.MasterDiskType = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetMaintainStartTime(v string) *DescribeInstanceResponseBody {
	s.MaintainStartTime = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetEngine(v string) *DescribeInstanceResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetTags(v *DescribeInstanceResponseBodyTags) *DescribeInstanceResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeInstanceResponseBody) SetStatus(v string) *DescribeInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetColdStorageSize(v int32) *DescribeInstanceResponseBody {
	s.ColdStorageSize = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetParentId(v string) *DescribeInstanceResponseBody {
	s.ParentId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetMajorVersion(v string) *DescribeInstanceResponseBody {
	s.MajorVersion = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetCoreDiskCount(v string) *DescribeInstanceResponseBody {
	s.CoreDiskCount = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetExpireTimeUTC(v string) *DescribeInstanceResponseBody {
	s.ExpireTimeUTC = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetMasterDiskSize(v int32) *DescribeInstanceResponseBody {
	s.MasterDiskSize = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetRequestId(v string) *DescribeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetZoneId(v string) *DescribeInstanceResponseBody {
	s.ZoneId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetClusterId(v string) *DescribeInstanceResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetInstanceId(v string) *DescribeInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetCoreNodeCount(v int32) *DescribeInstanceResponseBody {
	s.CoreNodeCount = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetColdStorageStatus(v string) *DescribeInstanceResponseBody {
	s.ColdStorageStatus = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetCreatedTimeUTC(v string) *DescribeInstanceResponseBody {
	s.CreatedTimeUTC = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetMinorVersion(v string) *DescribeInstanceResponseBody {
	s.MinorVersion = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetDuration(v int32) *DescribeInstanceResponseBody {
	s.Duration = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetPayType(v string) *DescribeInstanceResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetIsMultiModel(v bool) *DescribeInstanceResponseBody {
	s.IsMultiModel = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetClusterType(v string) *DescribeInstanceResponseBody {
	s.ClusterType = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetVswitchId(v string) *DescribeInstanceResponseBody {
	s.VswitchId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetMasterNodeCount(v int32) *DescribeInstanceResponseBody {
	s.MasterNodeCount = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetInstanceName(v string) *DescribeInstanceResponseBody {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetVpcId(v string) *DescribeInstanceResponseBody {
	s.VpcId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetAutoRenewal(v bool) *DescribeInstanceResponseBody {
	s.AutoRenewal = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetCoreDiskType(v string) *DescribeInstanceResponseBody {
	s.CoreDiskType = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetRegionId(v string) *DescribeInstanceResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetExpireTime(v string) *DescribeInstanceResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetBackupStatus(v string) *DescribeInstanceResponseBody {
	s.BackupStatus = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetCoreDiskSize(v int32) *DescribeInstanceResponseBody {
	s.CoreDiskSize = &v
	return s
}

type DescribeInstanceResponseBodyTags struct {
	Tag []*DescribeInstanceResponseBodyTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeInstanceResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyTags) SetTag(v []*DescribeInstanceResponseBodyTagsTag) *DescribeInstanceResponseBodyTags {
	s.Tag = v
	return s
}

type DescribeInstanceResponseBodyTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstanceResponseBodyTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyTagsTag) SetKey(v string) *DescribeInstanceResponseBodyTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeInstanceResponseBodyTagsTag) SetValue(v string) *DescribeInstanceResponseBodyTagsTag {
	s.Value = &v
	return s
}

type DescribeInstanceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponse) SetHeaders(v map[string]*string) *DescribeInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceResponse) SetBody(v *DescribeInstanceResponseBody) *DescribeInstanceResponse {
	s.Body = v
	return s
}

type DescribeInstancesRequest struct {
	RegionId        *string                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	PageNumber      *int32                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	DbType          *string                        `json:"DbType,omitempty" xml:"DbType,omitempty"`
	ClusterName     *string                        `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	ResourceGroupId *string                        `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tag             []*DescribeInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequest) SetRegionId(v string) *DescribeInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageNumber(v int32) *DescribeInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageSize(v int32) *DescribeInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesRequest) SetDbType(v string) *DescribeInstancesRequest {
	s.DbType = &v
	return s
}

func (s *DescribeInstancesRequest) SetClusterName(v string) *DescribeInstancesRequest {
	s.ClusterName = &v
	return s
}

func (s *DescribeInstancesRequest) SetResourceGroupId(v string) *DescribeInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstancesRequest) SetTag(v []*DescribeInstancesRequestTag) *DescribeInstancesRequest {
	s.Tag = v
	return s
}

type DescribeInstancesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstancesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequestTag) SetKey(v string) *DescribeInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeInstancesRequestTag) SetValue(v string) *DescribeInstancesRequestTag {
	s.Value = &v
	return s
}

type DescribeInstancesResponseBody struct {
	Instances  *DescribeInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	TotalCount *int64                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBody) SetInstances(v *DescribeInstancesResponseBodyInstances) *DescribeInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeInstancesResponseBody) SetTotalCount(v int64) *DescribeInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetPageSize(v int32) *DescribeInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetRequestId(v string) *DescribeInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetPageNumber(v int32) *DescribeInstancesResponseBody {
	s.PageNumber = &v
	return s
}

type DescribeInstancesResponseBodyInstances struct {
	Instance []*DescribeInstancesResponseBodyInstancesInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstances) SetInstance(v []*DescribeInstancesResponseBodyInstancesInstance) *DescribeInstancesResponseBodyInstances {
	s.Instance = v
	return s
}

type DescribeInstancesResponseBodyInstancesInstance struct {
	VpcId                *string                                             `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	Status               *string                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	ModuleId             *int32                                              `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	VswitchId            *string                                             `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	BackupStatus         *string                                             `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	PayType              *string                                             `json:"PayType,omitempty" xml:"PayType,omitempty"`
	CoreDiskType         *string                                             `json:"CoreDiskType,omitempty" xml:"CoreDiskType,omitempty"`
	Tags                 *DescribeInstancesResponseBodyInstancesInstanceTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	MasterNodeCount      *int32                                              `json:"MasterNodeCount,omitempty" xml:"MasterNodeCount,omitempty"`
	NetworkType          *string                                             `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	CreatedTimeUTC       *string                                             `json:"CreatedTimeUTC,omitempty" xml:"CreatedTimeUTC,omitempty"`
	ParentId             *string                                             `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	ExpireTimeUTC        *string                                             `json:"ExpireTimeUTC,omitempty" xml:"ExpireTimeUTC,omitempty"`
	InstanceName         *string                                             `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	MasterInstanceType   *string                                             `json:"MasterInstanceType,omitempty" xml:"MasterInstanceType,omitempty"`
	CoreInstanceType     *string                                             `json:"CoreInstanceType,omitempty" xml:"CoreInstanceType,omitempty"`
	CreatedTime          *string                                             `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	CoreDiskSize         *int32                                              `json:"CoreDiskSize,omitempty" xml:"CoreDiskSize,omitempty"`
	ClusterId            *string                                             `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ExpireTime           *string                                             `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	IsHa                 *bool                                               `json:"IsHa,omitempty" xml:"IsHa,omitempty"`
	InstanceId           *string                                             `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ColdStorageStatus    *string                                             `json:"ColdStorageStatus,omitempty" xml:"ColdStorageStatus,omitempty"`
	IsDeletionProtection *bool                                               `json:"IsDeletionProtection,omitempty" xml:"IsDeletionProtection,omitempty"`
	RegionId             *string                                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	MasterDiskSize       *int32                                              `json:"MasterDiskSize,omitempty" xml:"MasterDiskSize,omitempty"`
	MasterDiskType       *string                                             `json:"MasterDiskType,omitempty" xml:"MasterDiskType,omitempty"`
	AutoRenewal          *bool                                               `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	ClusterType          *string                                             `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	ResourceGroupId      *string                                             `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ClusterName          *string                                             `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	ZoneId               *string                                             `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	Duration             *int32                                              `json:"Duration,omitempty" xml:"Duration,omitempty"`
	ModuleStackVersion   *string                                             `json:"ModuleStackVersion,omitempty" xml:"ModuleStackVersion,omitempty"`
	Engine               *string                                             `json:"Engine,omitempty" xml:"Engine,omitempty"`
	MajorVersion         *string                                             `json:"MajorVersion,omitempty" xml:"MajorVersion,omitempty"`
	CoreDiskCount        *string                                             `json:"CoreDiskCount,omitempty" xml:"CoreDiskCount,omitempty"`
	CoreNodeCount        *int32                                              `json:"CoreNodeCount,omitempty" xml:"CoreNodeCount,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstance) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetVpcId(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.VpcId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetStatus(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.Status = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetModuleId(v int32) *DescribeInstancesResponseBodyInstancesInstance {
	s.ModuleId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetVswitchId(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.VswitchId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetBackupStatus(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.BackupStatus = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetPayType(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.PayType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetCoreDiskType(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.CoreDiskType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetTags(v *DescribeInstancesResponseBodyInstancesInstanceTags) *DescribeInstancesResponseBodyInstancesInstance {
	s.Tags = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetMasterNodeCount(v int32) *DescribeInstancesResponseBodyInstancesInstance {
	s.MasterNodeCount = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetNetworkType(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.NetworkType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetCreatedTimeUTC(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.CreatedTimeUTC = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetParentId(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.ParentId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetExpireTimeUTC(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.ExpireTimeUTC = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetInstanceName(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetMasterInstanceType(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.MasterInstanceType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetCoreInstanceType(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.CoreInstanceType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetCreatedTime(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.CreatedTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetCoreDiskSize(v int32) *DescribeInstancesResponseBodyInstancesInstance {
	s.CoreDiskSize = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetClusterId(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.ClusterId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetExpireTime(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetIsHa(v bool) *DescribeInstancesResponseBodyInstancesInstance {
	s.IsHa = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetInstanceId(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetColdStorageStatus(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.ColdStorageStatus = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetIsDeletionProtection(v bool) *DescribeInstancesResponseBodyInstancesInstance {
	s.IsDeletionProtection = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetRegionId(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetMasterDiskSize(v int32) *DescribeInstancesResponseBodyInstancesInstance {
	s.MasterDiskSize = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetMasterDiskType(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.MasterDiskType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetAutoRenewal(v bool) *DescribeInstancesResponseBodyInstancesInstance {
	s.AutoRenewal = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetClusterType(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.ClusterType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetResourceGroupId(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetClusterName(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.ClusterName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetZoneId(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.ZoneId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetDuration(v int32) *DescribeInstancesResponseBodyInstancesInstance {
	s.Duration = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetModuleStackVersion(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.ModuleStackVersion = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetEngine(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.Engine = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetMajorVersion(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.MajorVersion = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetCoreDiskCount(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.CoreDiskCount = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetCoreNodeCount(v int32) *DescribeInstancesResponseBodyInstancesInstance {
	s.CoreNodeCount = &v
	return s
}

type DescribeInstancesResponseBodyInstancesInstanceTags struct {
	Tag []*DescribeInstancesResponseBodyInstancesInstanceTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceTags) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceTags) SetTag(v []*DescribeInstancesResponseBodyInstancesInstanceTagsTag) *DescribeInstancesResponseBodyInstancesInstanceTags {
	s.Tag = v
	return s
}

type DescribeInstancesResponseBodyInstancesInstanceTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceTagsTag) SetKey(v string) *DescribeInstancesResponseBodyInstancesInstanceTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceTagsTag) SetValue(v string) *DescribeInstancesResponseBodyInstancesInstanceTagsTag {
	s.Value = &v
	return s
}

type DescribeInstancesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponse) SetHeaders(v map[string]*string) *DescribeInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstancesResponse) SetBody(v *DescribeInstancesResponseBody) *DescribeInstancesResponse {
	s.Body = v
	return s
}

type DescribeInstanceTypeRequest struct {
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s DescribeInstanceTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypeRequest) SetInstanceType(v string) *DescribeInstanceTypeRequest {
	s.InstanceType = &v
	return s
}

type DescribeInstanceTypeResponseBody struct {
	RequestId            *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceTypeSpecList *DescribeInstanceTypeResponseBodyInstanceTypeSpecList `json:"InstanceTypeSpecList,omitempty" xml:"InstanceTypeSpecList,omitempty" type:"Struct"`
}

func (s DescribeInstanceTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypeResponseBody) SetRequestId(v string) *DescribeInstanceTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceTypeResponseBody) SetInstanceTypeSpecList(v *DescribeInstanceTypeResponseBodyInstanceTypeSpecList) *DescribeInstanceTypeResponseBody {
	s.InstanceTypeSpecList = v
	return s
}

type DescribeInstanceTypeResponseBodyInstanceTypeSpecList struct {
	InstanceTypeSpec []*DescribeInstanceTypeResponseBodyInstanceTypeSpecListInstanceTypeSpec `json:"InstanceTypeSpec,omitempty" xml:"InstanceTypeSpec,omitempty" type:"Repeated"`
}

func (s DescribeInstanceTypeResponseBodyInstanceTypeSpecList) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTypeResponseBodyInstanceTypeSpecList) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypeResponseBodyInstanceTypeSpecList) SetInstanceTypeSpec(v []*DescribeInstanceTypeResponseBodyInstanceTypeSpecListInstanceTypeSpec) *DescribeInstanceTypeResponseBodyInstanceTypeSpecList {
	s.InstanceTypeSpec = v
	return s
}

type DescribeInstanceTypeResponseBodyInstanceTypeSpecListInstanceTypeSpec struct {
	CpuSize      *int64  `json:"CpuSize,omitempty" xml:"CpuSize,omitempty"`
	MemSize      *int64  `json:"MemSize,omitempty" xml:"MemSize,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s DescribeInstanceTypeResponseBodyInstanceTypeSpecListInstanceTypeSpec) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTypeResponseBodyInstanceTypeSpecListInstanceTypeSpec) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypeResponseBodyInstanceTypeSpecListInstanceTypeSpec) SetCpuSize(v int64) *DescribeInstanceTypeResponseBodyInstanceTypeSpecListInstanceTypeSpec {
	s.CpuSize = &v
	return s
}

func (s *DescribeInstanceTypeResponseBodyInstanceTypeSpecListInstanceTypeSpec) SetMemSize(v int64) *DescribeInstanceTypeResponseBodyInstanceTypeSpecListInstanceTypeSpec {
	s.MemSize = &v
	return s
}

func (s *DescribeInstanceTypeResponseBodyInstanceTypeSpecListInstanceTypeSpec) SetInstanceType(v string) *DescribeInstanceTypeResponseBodyInstanceTypeSpecListInstanceTypeSpec {
	s.InstanceType = &v
	return s
}

type DescribeInstanceTypeResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypeResponse) SetHeaders(v map[string]*string) *DescribeInstanceTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceTypeResponse) SetBody(v *DescribeInstanceTypeResponseBody) *DescribeInstanceTypeResponse {
	s.Body = v
	return s
}

type DescribeIpWhitelistRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeIpWhitelistRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpWhitelistRequest) GoString() string {
	return s.String()
}

func (s *DescribeIpWhitelistRequest) SetClusterId(v string) *DescribeIpWhitelistRequest {
	s.ClusterId = &v
	return s
}

type DescribeIpWhitelistResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Groups    *DescribeIpWhitelistResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Struct"`
}

func (s DescribeIpWhitelistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIpWhitelistResponseBody) SetRequestId(v string) *DescribeIpWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIpWhitelistResponseBody) SetGroups(v *DescribeIpWhitelistResponseBodyGroups) *DescribeIpWhitelistResponseBody {
	s.Groups = v
	return s
}

type DescribeIpWhitelistResponseBodyGroups struct {
	Group []*DescribeIpWhitelistResponseBodyGroupsGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Repeated"`
}

func (s DescribeIpWhitelistResponseBodyGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpWhitelistResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *DescribeIpWhitelistResponseBodyGroups) SetGroup(v []*DescribeIpWhitelistResponseBodyGroupsGroup) *DescribeIpWhitelistResponseBodyGroups {
	s.Group = v
	return s
}

type DescribeIpWhitelistResponseBodyGroupsGroup struct {
	IpVersion *int32                                            `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	GroupName *string                                           `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	IpList    *DescribeIpWhitelistResponseBodyGroupsGroupIpList `json:"IpList,omitempty" xml:"IpList,omitempty" type:"Struct"`
}

func (s DescribeIpWhitelistResponseBodyGroupsGroup) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpWhitelistResponseBodyGroupsGroup) GoString() string {
	return s.String()
}

func (s *DescribeIpWhitelistResponseBodyGroupsGroup) SetIpVersion(v int32) *DescribeIpWhitelistResponseBodyGroupsGroup {
	s.IpVersion = &v
	return s
}

func (s *DescribeIpWhitelistResponseBodyGroupsGroup) SetGroupName(v string) *DescribeIpWhitelistResponseBodyGroupsGroup {
	s.GroupName = &v
	return s
}

func (s *DescribeIpWhitelistResponseBodyGroupsGroup) SetIpList(v *DescribeIpWhitelistResponseBodyGroupsGroupIpList) *DescribeIpWhitelistResponseBodyGroupsGroup {
	s.IpList = v
	return s
}

type DescribeIpWhitelistResponseBodyGroupsGroupIpList struct {
	Ip []*string `json:"Ip,omitempty" xml:"Ip,omitempty" type:"Repeated"`
}

func (s DescribeIpWhitelistResponseBodyGroupsGroupIpList) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpWhitelistResponseBodyGroupsGroupIpList) GoString() string {
	return s.String()
}

func (s *DescribeIpWhitelistResponseBodyGroupsGroupIpList) SetIp(v []*string) *DescribeIpWhitelistResponseBodyGroupsGroupIpList {
	s.Ip = v
	return s
}

type DescribeIpWhitelistResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeIpWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeIpWhitelistResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpWhitelistResponse) GoString() string {
	return s.String()
}

func (s *DescribeIpWhitelistResponse) SetHeaders(v map[string]*string) *DescribeIpWhitelistResponse {
	s.Headers = v
	return s
}

func (s *DescribeIpWhitelistResponse) SetBody(v *DescribeIpWhitelistResponseBody) *DescribeIpWhitelistResponse {
	s.Body = v
	return s
}

type DescribeMultiZoneAvailableRegionsRequest struct {
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
}

func (s DescribeMultiZoneAvailableRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableRegionsRequest) SetAcceptLanguage(v string) *DescribeMultiZoneAvailableRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

type DescribeMultiZoneAvailableRegionsResponseBody struct {
	RequestId *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Regions   *DescribeMultiZoneAvailableRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
}

func (s DescribeMultiZoneAvailableRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableRegionsResponseBody) SetRequestId(v string) *DescribeMultiZoneAvailableRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMultiZoneAvailableRegionsResponseBody) SetRegions(v *DescribeMultiZoneAvailableRegionsResponseBodyRegions) *DescribeMultiZoneAvailableRegionsResponseBody {
	s.Regions = v
	return s
}

type DescribeMultiZoneAvailableRegionsResponseBodyRegions struct {
	Region []*DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s DescribeMultiZoneAvailableRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableRegionsResponseBodyRegions) SetRegion(v []*DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegion) *DescribeMultiZoneAvailableRegionsResponseBodyRegions {
	s.Region = v
	return s
}

type DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegion struct {
	LocalName         *string                                                                      `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionEndpoint    *string                                                                      `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	AvailableCombines *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombines `json:"AvailableCombines,omitempty" xml:"AvailableCombines,omitempty" type:"Struct"`
	RegionId          *string                                                                      `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegion) SetLocalName(v string) *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegion {
	s.LocalName = &v
	return s
}

func (s *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegion) SetRegionEndpoint(v string) *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegion {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegion) SetAvailableCombines(v *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombines) *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegion {
	s.AvailableCombines = v
	return s
}

func (s *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegion) SetRegionId(v string) *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegion {
	s.RegionId = &v
	return s
}

type DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombines struct {
	AvailableCombine []*DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombine `json:"AvailableCombine,omitempty" xml:"AvailableCombine,omitempty" type:"Repeated"`
}

func (s DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombines) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombines) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombines) SetAvailableCombine(v []*DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombine) *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombines {
	s.AvailableCombine = v
	return s
}

type DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombine struct {
	Zones *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombineZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
	Id    *string                                                                                           `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombine) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombine) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombine) SetZones(v *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombineZones) *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombine {
	s.Zones = v
	return s
}

func (s *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombine) SetId(v string) *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombine {
	s.Id = &v
	return s
}

type DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombineZones struct {
	Zone []*string `json:"Zone,omitempty" xml:"Zone,omitempty" type:"Repeated"`
}

func (s DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombineZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombineZones) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombineZones) SetZone(v []*string) *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombineZones {
	s.Zone = v
	return s
}

type DescribeMultiZoneAvailableRegionsResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMultiZoneAvailableRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMultiZoneAvailableRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableRegionsResponse) SetHeaders(v map[string]*string) *DescribeMultiZoneAvailableRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeMultiZoneAvailableRegionsResponse) SetBody(v *DescribeMultiZoneAvailableRegionsResponseBody) *DescribeMultiZoneAvailableRegionsResponse {
	s.Body = v
	return s
}

type DescribeMultiZoneAvailableResourceRequest struct {
	ChargeType      *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneCombination *string `json:"ZoneCombination,omitempty" xml:"ZoneCombination,omitempty"`
}

func (s DescribeMultiZoneAvailableResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceRequest) SetChargeType(v string) *DescribeMultiZoneAvailableResourceRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribeMultiZoneAvailableResourceRequest) SetRegionId(v string) *DescribeMultiZoneAvailableResourceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMultiZoneAvailableResourceRequest) SetZoneCombination(v string) *DescribeMultiZoneAvailableResourceRequest {
	s.ZoneCombination = &v
	return s
}

type DescribeMultiZoneAvailableResourceResponseBody struct {
	RequestId      *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AvailableZones *DescribeMultiZoneAvailableResourceResponseBodyAvailableZones `json:"AvailableZones,omitempty" xml:"AvailableZones,omitempty" type:"Struct"`
}

func (s DescribeMultiZoneAvailableResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBody) SetRequestId(v string) *DescribeMultiZoneAvailableResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBody) SetAvailableZones(v *DescribeMultiZoneAvailableResourceResponseBodyAvailableZones) *DescribeMultiZoneAvailableResourceResponseBody {
	s.AvailableZones = v
	return s
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZones struct {
	AvailableZone []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZone `json:"AvailableZone,omitempty" xml:"AvailableZone,omitempty" type:"Repeated"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZones) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZones) SetAvailableZone(v []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZone) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZones {
	s.AvailableZone = v
	return s
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZone struct {
	SupportedEngines *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEngines `json:"SupportedEngines,omitempty" xml:"SupportedEngines,omitempty" type:"Struct"`
	ZoneCombination  *string                                                                                    `json:"ZoneCombination,omitempty" xml:"ZoneCombination,omitempty"`
	MasterResources  *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResources  `json:"MasterResources,omitempty" xml:"MasterResources,omitempty" type:"Struct"`
	RegionId         *string                                                                                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZone) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZone) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZone) SetSupportedEngines(v *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEngines) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZone {
	s.SupportedEngines = v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZone) SetZoneCombination(v string) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZone {
	s.ZoneCombination = &v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZone) SetMasterResources(v *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResources) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZone {
	s.MasterResources = v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZone) SetRegionId(v string) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZone {
	s.RegionId = &v
	return s
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEngines struct {
	SupportedEngine []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine `json:"SupportedEngine,omitempty" xml:"SupportedEngine,omitempty" type:"Repeated"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEngines) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEngines) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEngines) SetSupportedEngine(v []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEngines {
	s.SupportedEngine = v
	return s
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine struct {
	SupportedEngineVersions *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersions `json:"SupportedEngineVersions,omitempty" xml:"SupportedEngineVersions,omitempty" type:"Struct"`
	Engine                  *string                                                                                                                          `json:"Engine,omitempty" xml:"Engine,omitempty"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine) SetSupportedEngineVersions(v *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersions) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine {
	s.SupportedEngineVersions = v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine) SetEngine(v string) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine {
	s.Engine = &v
	return s
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersions struct {
	SupportedEngineVersion []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion `json:"SupportedEngineVersion,omitempty" xml:"SupportedEngineVersion,omitempty" type:"Repeated"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersions) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersions) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersions) SetSupportedEngineVersion(v []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersions {
	s.SupportedEngineVersion = v
	return s
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion struct {
	Version             *string                                                                                                                                                                   `json:"Version,omitempty" xml:"Version,omitempty"`
	SupportedCategories *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategories `json:"SupportedCategories,omitempty" xml:"SupportedCategories,omitempty" type:"Struct"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion) SetVersion(v string) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion {
	s.Version = &v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion) SetSupportedCategories(v *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategories) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion {
	s.SupportedCategories = v
	return s
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategories struct {
	SupportedCategories []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories `json:"SupportedCategories,omitempty" xml:"SupportedCategories,omitempty" type:"Repeated"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategories) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategories) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategories) SetSupportedCategories(v []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategories {
	s.SupportedCategories = v
	return s
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories struct {
	SupportedStorageTypes *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypes `json:"SupportedStorageTypes,omitempty" xml:"SupportedStorageTypes,omitempty" type:"Struct"`
	Category              *string                                                                                                                                                                                                           `json:"Category,omitempty" xml:"Category,omitempty"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories) SetSupportedStorageTypes(v *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypes) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories {
	s.SupportedStorageTypes = v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories) SetCategory(v string) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories {
	s.Category = &v
	return s
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypes struct {
	SupportedStorageType []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType `json:"SupportedStorageType,omitempty" xml:"SupportedStorageType,omitempty" type:"Repeated"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypes) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypes) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypes) SetSupportedStorageType(v []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypes {
	s.SupportedStorageType = v
	return s
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType struct {
	CoreResources *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResources `json:"CoreResources,omitempty" xml:"CoreResources,omitempty" type:"Struct"`
	StorageType   *string                                                                                                                                                                                                                                            `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType) SetCoreResources(v *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResources) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType {
	s.CoreResources = v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType) SetStorageType(v string) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType {
	s.StorageType = &v
	return s
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResources struct {
	CoreResource []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource `json:"CoreResource,omitempty" xml:"CoreResource,omitempty" type:"Repeated"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResources) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResources) SetCoreResource(v []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResources {
	s.CoreResource = v
	return s
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource struct {
	InstanceTypeDetail     *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail     `json:"InstanceTypeDetail,omitempty" xml:"InstanceTypeDetail,omitempty" type:"Struct"`
	DBInstanceStorageRange *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange `json:"DBInstanceStorageRange,omitempty" xml:"DBInstanceStorageRange,omitempty" type:"Struct"`
	MaxCoreCount           *int32                                                                                                                                                                                                                                                                               `json:"MaxCoreCount,omitempty" xml:"MaxCoreCount,omitempty"`
	InstanceType           *string                                                                                                                                                                                                                                                                              `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource) SetInstanceTypeDetail(v *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource {
	s.InstanceTypeDetail = v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource) SetDBInstanceStorageRange(v *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource {
	s.DBInstanceStorageRange = v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource) SetMaxCoreCount(v int32) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource {
	s.MaxCoreCount = &v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource) SetInstanceType(v string) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource {
	s.InstanceType = &v
	return s
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail struct {
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Mem *int32 `json:"Mem,omitempty" xml:"Mem,omitempty"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail) SetCpu(v int32) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail {
	s.Cpu = &v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail) SetMem(v int32) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail {
	s.Mem = &v
	return s
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange struct {
	MaxSize  *int32 `json:"MaxSize,omitempty" xml:"MaxSize,omitempty"`
	StepSize *int32 `json:"StepSize,omitempty" xml:"StepSize,omitempty"`
	MinSize  *int32 `json:"MinSize,omitempty" xml:"MinSize,omitempty"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange) SetMaxSize(v int32) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange {
	s.MaxSize = &v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange) SetStepSize(v int32) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange {
	s.StepSize = &v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange) SetMinSize(v int32) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange {
	s.MinSize = &v
	return s
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResources struct {
	MasterResource []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource `json:"MasterResource,omitempty" xml:"MasterResource,omitempty" type:"Repeated"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResources) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResources) SetMasterResource(v []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResources {
	s.MasterResource = v
	return s
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource struct {
	InstanceTypeDetail *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail `json:"InstanceTypeDetail,omitempty" xml:"InstanceTypeDetail,omitempty" type:"Struct"`
	InstanceType       *string                                                                                                                   `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource) SetInstanceTypeDetail(v *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource {
	s.InstanceTypeDetail = v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource) SetInstanceType(v string) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource {
	s.InstanceType = &v
	return s
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail struct {
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Mem *int32 `json:"Mem,omitempty" xml:"Mem,omitempty"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail) SetCpu(v int32) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail {
	s.Cpu = &v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail) SetMem(v int32) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail {
	s.Mem = &v
	return s
}

type DescribeMultiZoneAvailableResourceResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMultiZoneAvailableResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMultiZoneAvailableResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponse) SetHeaders(v map[string]*string) *DescribeMultiZoneAvailableResourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponse) SetBody(v *DescribeMultiZoneAvailableResourceResponseBody) *DescribeMultiZoneAvailableResourceResponse {
	s.Body = v
	return s
}

type DescribeMultiZoneClusterRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeMultiZoneClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneClusterRequest) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneClusterRequest) SetClusterId(v string) *DescribeMultiZoneClusterRequest {
	s.ClusterId = &v
	return s
}

type DescribeMultiZoneClusterResponseBody struct {
	StandbyZoneId           *string                                                      `json:"StandbyZoneId,omitempty" xml:"StandbyZoneId,omitempty"`
	ModuleStackVersion      *string                                                      `json:"ModuleStackVersion,omitempty" xml:"ModuleStackVersion,omitempty"`
	CreatedTime             *string                                                      `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	ResourceGroupId         *string                                                      `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	PrimaryVSwitchIds       *string                                                      `json:"PrimaryVSwitchIds,omitempty" xml:"PrimaryVSwitchIds,omitempty"`
	MasterInstanceType      *string                                                      `json:"MasterInstanceType,omitempty" xml:"MasterInstanceType,omitempty"`
	LogDiskCount            *string                                                      `json:"LogDiskCount,omitempty" xml:"LogDiskCount,omitempty"`
	IsDeletionProtection    *bool                                                        `json:"IsDeletionProtection,omitempty" xml:"IsDeletionProtection,omitempty"`
	LogDiskSize             *int32                                                       `json:"LogDiskSize,omitempty" xml:"LogDiskSize,omitempty"`
	ModuleId                *int32                                                       `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ArbiterVSwitchIds       *string                                                      `json:"ArbiterVSwitchIds,omitempty" xml:"ArbiterVSwitchIds,omitempty"`
	MaintainEndTime         *string                                                      `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	StandbyVSwitchIds       *string                                                      `json:"StandbyVSwitchIds,omitempty" xml:"StandbyVSwitchIds,omitempty"`
	NetworkType             *string                                                      `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	CoreInstanceType        *string                                                      `json:"CoreInstanceType,omitempty" xml:"CoreInstanceType,omitempty"`
	ClusterName             *string                                                      `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	MasterDiskType          *string                                                      `json:"MasterDiskType,omitempty" xml:"MasterDiskType,omitempty"`
	MaintainStartTime       *string                                                      `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	Engine                  *string                                                      `json:"Engine,omitempty" xml:"Engine,omitempty"`
	Tags                    *DescribeMultiZoneClusterResponseBodyTags                    `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	ArbiterZoneId           *string                                                      `json:"ArbiterZoneId,omitempty" xml:"ArbiterZoneId,omitempty"`
	Status                  *string                                                      `json:"Status,omitempty" xml:"Status,omitempty"`
	ParentId                *string                                                      `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	MajorVersion            *string                                                      `json:"MajorVersion,omitempty" xml:"MajorVersion,omitempty"`
	CoreDiskCount           *string                                                      `json:"CoreDiskCount,omitempty" xml:"CoreDiskCount,omitempty"`
	MultiZoneInstanceModels *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModels `json:"MultiZoneInstanceModels,omitempty" xml:"MultiZoneInstanceModels,omitempty" type:"Struct"`
	ExpireTimeUTC           *string                                                      `json:"ExpireTimeUTC,omitempty" xml:"ExpireTimeUTC,omitempty"`
	PrimaryZoneId           *string                                                      `json:"PrimaryZoneId,omitempty" xml:"PrimaryZoneId,omitempty"`
	MasterDiskSize          *int32                                                       `json:"MasterDiskSize,omitempty" xml:"MasterDiskSize,omitempty"`
	RequestId               *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MultiZoneCombination    *string                                                      `json:"MultiZoneCombination,omitempty" xml:"MultiZoneCombination,omitempty"`
	ClusterId               *string                                                      `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	InstanceId              *string                                                      `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	CoreNodeCount           *int32                                                       `json:"CoreNodeCount,omitempty" xml:"CoreNodeCount,omitempty"`
	CreatedTimeUTC          *string                                                      `json:"CreatedTimeUTC,omitempty" xml:"CreatedTimeUTC,omitempty"`
	LogInstanceType         *string                                                      `json:"LogInstanceType,omitempty" xml:"LogInstanceType,omitempty"`
	Duration                *int32                                                       `json:"Duration,omitempty" xml:"Duration,omitempty"`
	PayType                 *string                                                      `json:"PayType,omitempty" xml:"PayType,omitempty"`
	MasterNodeCount         *int32                                                       `json:"MasterNodeCount,omitempty" xml:"MasterNodeCount,omitempty"`
	InstanceName            *string                                                      `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	VpcId                   *string                                                      `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	AutoRenewal             *bool                                                        `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	CoreDiskType            *string                                                      `json:"CoreDiskType,omitempty" xml:"CoreDiskType,omitempty"`
	LogNodeCount            *int32                                                       `json:"LogNodeCount,omitempty" xml:"LogNodeCount,omitempty"`
	LogDiskType             *string                                                      `json:"LogDiskType,omitempty" xml:"LogDiskType,omitempty"`
	RegionId                *string                                                      `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ExpireTime              *string                                                      `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	CoreDiskSize            *int32                                                       `json:"CoreDiskSize,omitempty" xml:"CoreDiskSize,omitempty"`
}

func (s DescribeMultiZoneClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneClusterResponseBody) SetStandbyZoneId(v string) *DescribeMultiZoneClusterResponseBody {
	s.StandbyZoneId = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetModuleStackVersion(v string) *DescribeMultiZoneClusterResponseBody {
	s.ModuleStackVersion = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetCreatedTime(v string) *DescribeMultiZoneClusterResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetResourceGroupId(v string) *DescribeMultiZoneClusterResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetPrimaryVSwitchIds(v string) *DescribeMultiZoneClusterResponseBody {
	s.PrimaryVSwitchIds = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetMasterInstanceType(v string) *DescribeMultiZoneClusterResponseBody {
	s.MasterInstanceType = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetLogDiskCount(v string) *DescribeMultiZoneClusterResponseBody {
	s.LogDiskCount = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetIsDeletionProtection(v bool) *DescribeMultiZoneClusterResponseBody {
	s.IsDeletionProtection = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetLogDiskSize(v int32) *DescribeMultiZoneClusterResponseBody {
	s.LogDiskSize = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetModuleId(v int32) *DescribeMultiZoneClusterResponseBody {
	s.ModuleId = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetArbiterVSwitchIds(v string) *DescribeMultiZoneClusterResponseBody {
	s.ArbiterVSwitchIds = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetMaintainEndTime(v string) *DescribeMultiZoneClusterResponseBody {
	s.MaintainEndTime = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetStandbyVSwitchIds(v string) *DescribeMultiZoneClusterResponseBody {
	s.StandbyVSwitchIds = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetNetworkType(v string) *DescribeMultiZoneClusterResponseBody {
	s.NetworkType = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetCoreInstanceType(v string) *DescribeMultiZoneClusterResponseBody {
	s.CoreInstanceType = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetClusterName(v string) *DescribeMultiZoneClusterResponseBody {
	s.ClusterName = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetMasterDiskType(v string) *DescribeMultiZoneClusterResponseBody {
	s.MasterDiskType = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetMaintainStartTime(v string) *DescribeMultiZoneClusterResponseBody {
	s.MaintainStartTime = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetEngine(v string) *DescribeMultiZoneClusterResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetTags(v *DescribeMultiZoneClusterResponseBodyTags) *DescribeMultiZoneClusterResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetArbiterZoneId(v string) *DescribeMultiZoneClusterResponseBody {
	s.ArbiterZoneId = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetStatus(v string) *DescribeMultiZoneClusterResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetParentId(v string) *DescribeMultiZoneClusterResponseBody {
	s.ParentId = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetMajorVersion(v string) *DescribeMultiZoneClusterResponseBody {
	s.MajorVersion = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetCoreDiskCount(v string) *DescribeMultiZoneClusterResponseBody {
	s.CoreDiskCount = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetMultiZoneInstanceModels(v *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModels) *DescribeMultiZoneClusterResponseBody {
	s.MultiZoneInstanceModels = v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetExpireTimeUTC(v string) *DescribeMultiZoneClusterResponseBody {
	s.ExpireTimeUTC = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetPrimaryZoneId(v string) *DescribeMultiZoneClusterResponseBody {
	s.PrimaryZoneId = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetMasterDiskSize(v int32) *DescribeMultiZoneClusterResponseBody {
	s.MasterDiskSize = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetRequestId(v string) *DescribeMultiZoneClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetMultiZoneCombination(v string) *DescribeMultiZoneClusterResponseBody {
	s.MultiZoneCombination = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetClusterId(v string) *DescribeMultiZoneClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetInstanceId(v string) *DescribeMultiZoneClusterResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetCoreNodeCount(v int32) *DescribeMultiZoneClusterResponseBody {
	s.CoreNodeCount = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetCreatedTimeUTC(v string) *DescribeMultiZoneClusterResponseBody {
	s.CreatedTimeUTC = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetLogInstanceType(v string) *DescribeMultiZoneClusterResponseBody {
	s.LogInstanceType = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetDuration(v int32) *DescribeMultiZoneClusterResponseBody {
	s.Duration = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetPayType(v string) *DescribeMultiZoneClusterResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetMasterNodeCount(v int32) *DescribeMultiZoneClusterResponseBody {
	s.MasterNodeCount = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetInstanceName(v string) *DescribeMultiZoneClusterResponseBody {
	s.InstanceName = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetVpcId(v string) *DescribeMultiZoneClusterResponseBody {
	s.VpcId = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetAutoRenewal(v bool) *DescribeMultiZoneClusterResponseBody {
	s.AutoRenewal = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetCoreDiskType(v string) *DescribeMultiZoneClusterResponseBody {
	s.CoreDiskType = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetLogNodeCount(v int32) *DescribeMultiZoneClusterResponseBody {
	s.LogNodeCount = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetLogDiskType(v string) *DescribeMultiZoneClusterResponseBody {
	s.LogDiskType = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetRegionId(v string) *DescribeMultiZoneClusterResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetExpireTime(v string) *DescribeMultiZoneClusterResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetCoreDiskSize(v int32) *DescribeMultiZoneClusterResponseBody {
	s.CoreDiskSize = &v
	return s
}

type DescribeMultiZoneClusterResponseBodyTags struct {
	Tag []*DescribeMultiZoneClusterResponseBodyTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeMultiZoneClusterResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneClusterResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneClusterResponseBodyTags) SetTag(v []*DescribeMultiZoneClusterResponseBodyTagsTag) *DescribeMultiZoneClusterResponseBodyTags {
	s.Tag = v
	return s
}

type DescribeMultiZoneClusterResponseBodyTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeMultiZoneClusterResponseBodyTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneClusterResponseBodyTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneClusterResponseBodyTagsTag) SetKey(v string) *DescribeMultiZoneClusterResponseBodyTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBodyTagsTag) SetValue(v string) *DescribeMultiZoneClusterResponseBodyTagsTag {
	s.Value = &v
	return s
}

type DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModels struct {
	MultiZoneInstanceModel []*DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel `json:"MultiZoneInstanceModel,omitempty" xml:"MultiZoneInstanceModel,omitempty" type:"Repeated"`
}

func (s DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModels) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModels) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModels) SetMultiZoneInstanceModel(v []*DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel) *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModels {
	s.MultiZoneInstanceModel = v
	return s
}

type DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel struct {
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	IsLatestVersion *bool   `json:"IsLatestVersion,omitempty" xml:"IsLatestVersion,omitempty"`
	InsName         *string `json:"InsName,omitempty" xml:"InsName,omitempty"`
	Role            *string `json:"Role,omitempty" xml:"Role,omitempty"`
	MinorVersion    *string `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
}

func (s DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel) SetStatus(v string) *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel {
	s.Status = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel) SetIsLatestVersion(v bool) *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel {
	s.IsLatestVersion = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel) SetInsName(v string) *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel {
	s.InsName = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel) SetRole(v string) *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel {
	s.Role = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel) SetMinorVersion(v string) *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel {
	s.MinorVersion = &v
	return s
}

type DescribeMultiZoneClusterResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMultiZoneClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMultiZoneClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneClusterResponse) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneClusterResponse) SetHeaders(v map[string]*string) *DescribeMultiZoneClusterResponse {
	s.Headers = v
	return s
}

func (s *DescribeMultiZoneClusterResponse) SetBody(v *DescribeMultiZoneClusterResponseBody) *DescribeMultiZoneClusterResponse {
	s.Body = v
	return s
}

type DescribeRecoverableTimeRangeRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeRecoverableTimeRangeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecoverableTimeRangeRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecoverableTimeRangeRequest) SetClusterId(v string) *DescribeRecoverableTimeRangeRequest {
	s.ClusterId = &v
	return s
}

type DescribeRecoverableTimeRangeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TimeEnd   *string `json:"TimeEnd,omitempty" xml:"TimeEnd,omitempty"`
	TimeBegin *string `json:"TimeBegin,omitempty" xml:"TimeBegin,omitempty"`
}

func (s DescribeRecoverableTimeRangeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecoverableTimeRangeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRecoverableTimeRangeResponseBody) SetRequestId(v string) *DescribeRecoverableTimeRangeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRecoverableTimeRangeResponseBody) SetTimeEnd(v string) *DescribeRecoverableTimeRangeResponseBody {
	s.TimeEnd = &v
	return s
}

func (s *DescribeRecoverableTimeRangeResponseBody) SetTimeBegin(v string) *DescribeRecoverableTimeRangeResponseBody {
	s.TimeBegin = &v
	return s
}

type DescribeRecoverableTimeRangeResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRecoverableTimeRangeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRecoverableTimeRangeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecoverableTimeRangeResponse) GoString() string {
	return s.String()
}

func (s *DescribeRecoverableTimeRangeResponse) SetHeaders(v map[string]*string) *DescribeRecoverableTimeRangeResponse {
	s.Headers = v
	return s
}

func (s *DescribeRecoverableTimeRangeResponse) SetBody(v *DescribeRecoverableTimeRangeResponseBody) *DescribeRecoverableTimeRangeResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
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
	Zones          *DescribeRegionsResponseBodyRegionsRegionZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
	LocalName      *string                                        `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionEndpoint *string                                        `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	RegionId       *string                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetZones(v *DescribeRegionsResponseBodyRegionsRegionZones) *DescribeRegionsResponseBodyRegionsRegion {
	s.Zones = v
	return s
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
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegionZonesZone) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegionZonesZone) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegionZonesZone) SetId(v string) *DescribeRegionsResponseBodyRegionsRegionZonesZone {
	s.Id = &v
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

type DescribeRestoreFullDetailsRequest struct {
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RestoreRecordId *string `json:"RestoreRecordId,omitempty" xml:"RestoreRecordId,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeRestoreFullDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreFullDetailsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRestoreFullDetailsRequest) SetClusterId(v string) *DescribeRestoreFullDetailsRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeRestoreFullDetailsRequest) SetRestoreRecordId(v string) *DescribeRestoreFullDetailsRequest {
	s.RestoreRecordId = &v
	return s
}

func (s *DescribeRestoreFullDetailsRequest) SetPageNumber(v int32) *DescribeRestoreFullDetailsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRestoreFullDetailsRequest) SetPageSize(v int32) *DescribeRestoreFullDetailsRequest {
	s.PageSize = &v
	return s
}

type DescribeRestoreFullDetailsResponseBody struct {
	RestoreFull *DescribeRestoreFullDetailsResponseBodyRestoreFull `json:"RestoreFull,omitempty" xml:"RestoreFull,omitempty" type:"Struct"`
	RequestId   *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRestoreFullDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreFullDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRestoreFullDetailsResponseBody) SetRestoreFull(v *DescribeRestoreFullDetailsResponseBodyRestoreFull) *DescribeRestoreFullDetailsResponseBody {
	s.RestoreFull = v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBody) SetRequestId(v string) *DescribeRestoreFullDetailsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRestoreFullDetailsResponseBodyRestoreFull struct {
	Succeed            *int32                                                               `json:"Succeed,omitempty" xml:"Succeed,omitempty"`
	DataSize           *string                                                              `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	Speed              *string                                                              `json:"Speed,omitempty" xml:"Speed,omitempty"`
	RestoreFullDetails *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetails `json:"RestoreFullDetails,omitempty" xml:"RestoreFullDetails,omitempty" type:"Struct"`
	PageSize           *int32                                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Fail               *int32                                                               `json:"Fail,omitempty" xml:"Fail,omitempty"`
	PageNumber         *int32                                                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Total              *int64                                                               `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeRestoreFullDetailsResponseBodyRestoreFull) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreFullDetailsResponseBodyRestoreFull) GoString() string {
	return s.String()
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFull) SetSucceed(v int32) *DescribeRestoreFullDetailsResponseBodyRestoreFull {
	s.Succeed = &v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFull) SetDataSize(v string) *DescribeRestoreFullDetailsResponseBodyRestoreFull {
	s.DataSize = &v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFull) SetSpeed(v string) *DescribeRestoreFullDetailsResponseBodyRestoreFull {
	s.Speed = &v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFull) SetRestoreFullDetails(v *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetails) *DescribeRestoreFullDetailsResponseBodyRestoreFull {
	s.RestoreFullDetails = v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFull) SetPageSize(v int32) *DescribeRestoreFullDetailsResponseBodyRestoreFull {
	s.PageSize = &v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFull) SetFail(v int32) *DescribeRestoreFullDetailsResponseBodyRestoreFull {
	s.Fail = &v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFull) SetPageNumber(v int32) *DescribeRestoreFullDetailsResponseBodyRestoreFull {
	s.PageNumber = &v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFull) SetTotal(v int64) *DescribeRestoreFullDetailsResponseBodyRestoreFull {
	s.Total = &v
	return s
}

type DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetails struct {
	RestoreFullDetail []*DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail `json:"RestoreFullDetail,omitempty" xml:"RestoreFullDetail,omitempty" type:"Repeated"`
}

func (s DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetails) GoString() string {
	return s.String()
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetails) SetRestoreFullDetail(v []*DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetails {
	s.RestoreFullDetail = v
	return s
}

type DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Process   *string `json:"Process,omitempty" xml:"Process,omitempty"`
	DataSize  *string `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	Speed     *string `json:"Speed,omitempty" xml:"Speed,omitempty"`
	State     *string `json:"State,omitempty" xml:"State,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Table     *string `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) GoString() string {
	return s.String()
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetEndTime(v string) *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.EndTime = &v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetStartTime(v string) *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.StartTime = &v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetProcess(v string) *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.Process = &v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetDataSize(v string) *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.DataSize = &v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetSpeed(v string) *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.Speed = &v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetState(v string) *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.State = &v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetMessage(v string) *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.Message = &v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetTable(v string) *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.Table = &v
	return s
}

type DescribeRestoreFullDetailsResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRestoreFullDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRestoreFullDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreFullDetailsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRestoreFullDetailsResponse) SetHeaders(v map[string]*string) *DescribeRestoreFullDetailsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRestoreFullDetailsResponse) SetBody(v *DescribeRestoreFullDetailsResponseBody) *DescribeRestoreFullDetailsResponse {
	s.Body = v
	return s
}

type DescribeRestoreIncrDetailRequest struct {
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RestoreRecordId *string `json:"RestoreRecordId,omitempty" xml:"RestoreRecordId,omitempty"`
}

func (s DescribeRestoreIncrDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreIncrDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeRestoreIncrDetailRequest) SetClusterId(v string) *DescribeRestoreIncrDetailRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeRestoreIncrDetailRequest) SetRestoreRecordId(v string) *DescribeRestoreIncrDetailRequest {
	s.RestoreRecordId = &v
	return s
}

type DescribeRestoreIncrDetailResponseBody struct {
	RequestId         *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RestoreIncrDetail *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail `json:"RestoreIncrDetail,omitempty" xml:"RestoreIncrDetail,omitempty" type:"Struct"`
}

func (s DescribeRestoreIncrDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreIncrDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRestoreIncrDetailResponseBody) SetRequestId(v string) *DescribeRestoreIncrDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRestoreIncrDetailResponseBody) SetRestoreIncrDetail(v *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail) *DescribeRestoreIncrDetailResponseBody {
	s.RestoreIncrDetail = v
	return s
}

type DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail struct {
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Process        *string `json:"Process,omitempty" xml:"Process,omitempty"`
	RestoreStartTs *string `json:"RestoreStartTs,omitempty" xml:"RestoreStartTs,omitempty"`
	State          *string `json:"State,omitempty" xml:"State,omitempty"`
	RestoredTs     *string `json:"RestoredTs,omitempty" xml:"RestoredTs,omitempty"`
	RestoreDelay   *string `json:"RestoreDelay,omitempty" xml:"RestoreDelay,omitempty"`
}

func (s DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail) GoString() string {
	return s.String()
}

func (s *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail) SetEndTime(v string) *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail {
	s.EndTime = &v
	return s
}

func (s *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail) SetStartTime(v string) *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail {
	s.StartTime = &v
	return s
}

func (s *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail) SetProcess(v string) *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail {
	s.Process = &v
	return s
}

func (s *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail) SetRestoreStartTs(v string) *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail {
	s.RestoreStartTs = &v
	return s
}

func (s *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail) SetState(v string) *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail {
	s.State = &v
	return s
}

func (s *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail) SetRestoredTs(v string) *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail {
	s.RestoredTs = &v
	return s
}

func (s *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail) SetRestoreDelay(v string) *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail {
	s.RestoreDelay = &v
	return s
}

type DescribeRestoreIncrDetailResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRestoreIncrDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRestoreIncrDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreIncrDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeRestoreIncrDetailResponse) SetHeaders(v map[string]*string) *DescribeRestoreIncrDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeRestoreIncrDetailResponse) SetBody(v *DescribeRestoreIncrDetailResponseBody) *DescribeRestoreIncrDetailResponse {
	s.Body = v
	return s
}

type DescribeRestoreSchemaDetailsRequest struct {
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RestoreRecordId *string `json:"RestoreRecordId,omitempty" xml:"RestoreRecordId,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeRestoreSchemaDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreSchemaDetailsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRestoreSchemaDetailsRequest) SetClusterId(v string) *DescribeRestoreSchemaDetailsRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeRestoreSchemaDetailsRequest) SetRestoreRecordId(v string) *DescribeRestoreSchemaDetailsRequest {
	s.RestoreRecordId = &v
	return s
}

func (s *DescribeRestoreSchemaDetailsRequest) SetPageNumber(v int32) *DescribeRestoreSchemaDetailsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRestoreSchemaDetailsRequest) SetPageSize(v int32) *DescribeRestoreSchemaDetailsRequest {
	s.PageSize = &v
	return s
}

type DescribeRestoreSchemaDetailsResponseBody struct {
	RestoreSchema *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema `json:"RestoreSchema,omitempty" xml:"RestoreSchema,omitempty" type:"Struct"`
	RequestId     *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRestoreSchemaDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreSchemaDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRestoreSchemaDetailsResponseBody) SetRestoreSchema(v *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema) *DescribeRestoreSchemaDetailsResponseBody {
	s.RestoreSchema = v
	return s
}

func (s *DescribeRestoreSchemaDetailsResponseBody) SetRequestId(v string) *DescribeRestoreSchemaDetailsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRestoreSchemaDetailsResponseBodyRestoreSchema struct {
	Succeed              *int32                                                                     `json:"Succeed,omitempty" xml:"Succeed,omitempty"`
	RestoreSchemaDetails *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetails `json:"RestoreSchemaDetails,omitempty" xml:"RestoreSchemaDetails,omitempty" type:"Struct"`
	PageSize             *int32                                                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber           *int32                                                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Fail                 *int32                                                                     `json:"Fail,omitempty" xml:"Fail,omitempty"`
	Total                *int64                                                                     `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeRestoreSchemaDetailsResponseBodyRestoreSchema) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreSchemaDetailsResponseBodyRestoreSchema) GoString() string {
	return s.String()
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema) SetSucceed(v int32) *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema {
	s.Succeed = &v
	return s
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema) SetRestoreSchemaDetails(v *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetails) *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema {
	s.RestoreSchemaDetails = v
	return s
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema) SetPageSize(v int32) *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema {
	s.PageSize = &v
	return s
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema) SetPageNumber(v int32) *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema {
	s.PageNumber = &v
	return s
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema) SetFail(v int32) *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema {
	s.Fail = &v
	return s
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema) SetTotal(v int64) *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema {
	s.Total = &v
	return s
}

type DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetails struct {
	RestoreSchemaDetail []*DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail `json:"RestoreSchemaDetail,omitempty" xml:"RestoreSchemaDetail,omitempty" type:"Repeated"`
}

func (s DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetails) GoString() string {
	return s.String()
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetails) SetRestoreSchemaDetail(v []*DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetails {
	s.RestoreSchemaDetail = v
	return s
}

type DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	State     *string `json:"State,omitempty" xml:"State,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Table     *string `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) GoString() string {
	return s.String()
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) SetEndTime(v string) *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail {
	s.EndTime = &v
	return s
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) SetStartTime(v string) *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail {
	s.StartTime = &v
	return s
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) SetState(v string) *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail {
	s.State = &v
	return s
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) SetMessage(v string) *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail {
	s.Message = &v
	return s
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) SetTable(v string) *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail {
	s.Table = &v
	return s
}

type DescribeRestoreSchemaDetailsResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRestoreSchemaDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRestoreSchemaDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreSchemaDetailsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRestoreSchemaDetailsResponse) SetHeaders(v map[string]*string) *DescribeRestoreSchemaDetailsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRestoreSchemaDetailsResponse) SetBody(v *DescribeRestoreSchemaDetailsResponseBody) *DescribeRestoreSchemaDetailsResponse {
	s.Body = v
	return s
}

type DescribeRestoreSummaryRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeRestoreSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeRestoreSummaryRequest) SetClusterId(v string) *DescribeRestoreSummaryRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeRestoreSummaryRequest) SetPageNumber(v int32) *DescribeRestoreSummaryRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRestoreSummaryRequest) SetPageSize(v int32) *DescribeRestoreSummaryRequest {
	s.PageSize = &v
	return s
}

type DescribeRestoreSummaryResponseBody struct {
	PageSize             *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId            *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber           *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Total                *int32                                      `json:"Total,omitempty" xml:"Total,omitempty"`
	HasMoreRestoreRecord *int32                                      `json:"HasMoreRestoreRecord,omitempty" xml:"HasMoreRestoreRecord,omitempty"`
	Rescords             *DescribeRestoreSummaryResponseBodyRescords `json:"Rescords,omitempty" xml:"Rescords,omitempty" type:"Struct"`
}

func (s DescribeRestoreSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRestoreSummaryResponseBody) SetPageSize(v int32) *DescribeRestoreSummaryResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRestoreSummaryResponseBody) SetRequestId(v string) *DescribeRestoreSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRestoreSummaryResponseBody) SetPageNumber(v int32) *DescribeRestoreSummaryResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeRestoreSummaryResponseBody) SetTotal(v int32) *DescribeRestoreSummaryResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeRestoreSummaryResponseBody) SetHasMoreRestoreRecord(v int32) *DescribeRestoreSummaryResponseBody {
	s.HasMoreRestoreRecord = &v
	return s
}

func (s *DescribeRestoreSummaryResponseBody) SetRescords(v *DescribeRestoreSummaryResponseBodyRescords) *DescribeRestoreSummaryResponseBody {
	s.Rescords = v
	return s
}

type DescribeRestoreSummaryResponseBodyRescords struct {
	Rescord []*DescribeRestoreSummaryResponseBodyRescordsRescord `json:"Rescord,omitempty" xml:"Rescord,omitempty" type:"Repeated"`
}

func (s DescribeRestoreSummaryResponseBodyRescords) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreSummaryResponseBodyRescords) GoString() string {
	return s.String()
}

func (s *DescribeRestoreSummaryResponseBodyRescords) SetRescord(v []*DescribeRestoreSummaryResponseBodyRescordsRescord) *DescribeRestoreSummaryResponseBodyRescords {
	s.Rescord = v
	return s
}

type DescribeRestoreSummaryResponseBodyRescordsRescord struct {
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	FinishTime          *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	SchemaProcess       *string `json:"SchemaProcess,omitempty" xml:"SchemaProcess,omitempty"`
	BulkLoadProcess     *string `json:"BulkLoadProcess,omitempty" xml:"BulkLoadProcess,omitempty"`
	RecordId            *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	LogProcess          *string `json:"LogProcess,omitempty" xml:"LogProcess,omitempty"`
	HfileRestoreProcess *string `json:"HfileRestoreProcess,omitempty" xml:"HfileRestoreProcess,omitempty"`
}

func (s DescribeRestoreSummaryResponseBodyRescordsRescord) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreSummaryResponseBodyRescordsRescord) GoString() string {
	return s.String()
}

func (s *DescribeRestoreSummaryResponseBodyRescordsRescord) SetStatus(v string) *DescribeRestoreSummaryResponseBodyRescordsRescord {
	s.Status = &v
	return s
}

func (s *DescribeRestoreSummaryResponseBodyRescordsRescord) SetFinishTime(v string) *DescribeRestoreSummaryResponseBodyRescordsRescord {
	s.FinishTime = &v
	return s
}

func (s *DescribeRestoreSummaryResponseBodyRescordsRescord) SetSchemaProcess(v string) *DescribeRestoreSummaryResponseBodyRescordsRescord {
	s.SchemaProcess = &v
	return s
}

func (s *DescribeRestoreSummaryResponseBodyRescordsRescord) SetBulkLoadProcess(v string) *DescribeRestoreSummaryResponseBodyRescordsRescord {
	s.BulkLoadProcess = &v
	return s
}

func (s *DescribeRestoreSummaryResponseBodyRescordsRescord) SetRecordId(v string) *DescribeRestoreSummaryResponseBodyRescordsRescord {
	s.RecordId = &v
	return s
}

func (s *DescribeRestoreSummaryResponseBodyRescordsRescord) SetCreateTime(v string) *DescribeRestoreSummaryResponseBodyRescordsRescord {
	s.CreateTime = &v
	return s
}

func (s *DescribeRestoreSummaryResponseBodyRescordsRescord) SetLogProcess(v string) *DescribeRestoreSummaryResponseBodyRescordsRescord {
	s.LogProcess = &v
	return s
}

func (s *DescribeRestoreSummaryResponseBodyRescordsRescord) SetHfileRestoreProcess(v string) *DescribeRestoreSummaryResponseBodyRescordsRescord {
	s.HfileRestoreProcess = &v
	return s
}

type DescribeRestoreSummaryResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRestoreSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRestoreSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeRestoreSummaryResponse) SetHeaders(v map[string]*string) *DescribeRestoreSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeRestoreSummaryResponse) SetBody(v *DescribeRestoreSummaryResponseBody) *DescribeRestoreSummaryResponse {
	s.Body = v
	return s
}

type DescribeRestoreTablesRequest struct {
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RestoreRecordId *string `json:"RestoreRecordId,omitempty" xml:"RestoreRecordId,omitempty"`
}

func (s DescribeRestoreTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreTablesRequest) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTablesRequest) SetClusterId(v string) *DescribeRestoreTablesRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeRestoreTablesRequest) SetRestoreRecordId(v string) *DescribeRestoreTablesRequest {
	s.RestoreRecordId = &v
	return s
}

type DescribeRestoreTablesResponseBody struct {
	RestoreSchema     *DescribeRestoreTablesResponseBodyRestoreSchema     `json:"RestoreSchema,omitempty" xml:"RestoreSchema,omitempty" type:"Struct"`
	RestoreFull       *DescribeRestoreTablesResponseBodyRestoreFull       `json:"RestoreFull,omitempty" xml:"RestoreFull,omitempty" type:"Struct"`
	RequestId         *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tables            *DescribeRestoreTablesResponseBodyTables            `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Struct"`
	RestoreSummary    *DescribeRestoreTablesResponseBodyRestoreSummary    `json:"RestoreSummary,omitempty" xml:"RestoreSummary,omitempty" type:"Struct"`
	RestoreIncrDetail *DescribeRestoreTablesResponseBodyRestoreIncrDetail `json:"RestoreIncrDetail,omitempty" xml:"RestoreIncrDetail,omitempty" type:"Struct"`
}

func (s DescribeRestoreTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreTablesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTablesResponseBody) SetRestoreSchema(v *DescribeRestoreTablesResponseBodyRestoreSchema) *DescribeRestoreTablesResponseBody {
	s.RestoreSchema = v
	return s
}

func (s *DescribeRestoreTablesResponseBody) SetRestoreFull(v *DescribeRestoreTablesResponseBodyRestoreFull) *DescribeRestoreTablesResponseBody {
	s.RestoreFull = v
	return s
}

func (s *DescribeRestoreTablesResponseBody) SetRequestId(v string) *DescribeRestoreTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRestoreTablesResponseBody) SetTables(v *DescribeRestoreTablesResponseBodyTables) *DescribeRestoreTablesResponseBody {
	s.Tables = v
	return s
}

func (s *DescribeRestoreTablesResponseBody) SetRestoreSummary(v *DescribeRestoreTablesResponseBodyRestoreSummary) *DescribeRestoreTablesResponseBody {
	s.RestoreSummary = v
	return s
}

func (s *DescribeRestoreTablesResponseBody) SetRestoreIncrDetail(v *DescribeRestoreTablesResponseBodyRestoreIncrDetail) *DescribeRestoreTablesResponseBody {
	s.RestoreIncrDetail = v
	return s
}

type DescribeRestoreTablesResponseBodyRestoreSchema struct {
	Succeed              *int32                                                              `json:"Succeed,omitempty" xml:"Succeed,omitempty"`
	RestoreSchemaDetails *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetails `json:"RestoreSchemaDetails,omitempty" xml:"RestoreSchemaDetails,omitempty" type:"Struct"`
	PageSize             *int32                                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber           *int32                                                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Fail                 *int32                                                              `json:"Fail,omitempty" xml:"Fail,omitempty"`
	Total                *int64                                                              `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeRestoreTablesResponseBodyRestoreSchema) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreTablesResponseBodyRestoreSchema) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchema) SetSucceed(v int32) *DescribeRestoreTablesResponseBodyRestoreSchema {
	s.Succeed = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchema) SetRestoreSchemaDetails(v *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetails) *DescribeRestoreTablesResponseBodyRestoreSchema {
	s.RestoreSchemaDetails = v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchema) SetPageSize(v int32) *DescribeRestoreTablesResponseBodyRestoreSchema {
	s.PageSize = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchema) SetPageNumber(v int32) *DescribeRestoreTablesResponseBodyRestoreSchema {
	s.PageNumber = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchema) SetFail(v int32) *DescribeRestoreTablesResponseBodyRestoreSchema {
	s.Fail = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchema) SetTotal(v int64) *DescribeRestoreTablesResponseBodyRestoreSchema {
	s.Total = &v
	return s
}

type DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetails struct {
	RestoreSchemaDetail []*DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail `json:"RestoreSchemaDetail,omitempty" xml:"RestoreSchemaDetail,omitempty" type:"Repeated"`
}

func (s DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetails) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetails) SetRestoreSchemaDetail(v []*DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetails {
	s.RestoreSchemaDetail = v
	return s
}

type DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	State     *string `json:"State,omitempty" xml:"State,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Table     *string `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) SetEndTime(v string) *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail {
	s.EndTime = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) SetStartTime(v string) *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail {
	s.StartTime = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) SetState(v string) *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail {
	s.State = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) SetMessage(v string) *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail {
	s.Message = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) SetTable(v string) *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail {
	s.Table = &v
	return s
}

type DescribeRestoreTablesResponseBodyRestoreFull struct {
	Succeed            *int32                                                          `json:"Succeed,omitempty" xml:"Succeed,omitempty"`
	DataSize           *string                                                         `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	Speed              *string                                                         `json:"Speed,omitempty" xml:"Speed,omitempty"`
	RestoreFullDetails *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetails `json:"RestoreFullDetails,omitempty" xml:"RestoreFullDetails,omitempty" type:"Struct"`
	PageSize           *int32                                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Fail               *int32                                                          `json:"Fail,omitempty" xml:"Fail,omitempty"`
	PageNumber         *int32                                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Total              *int64                                                          `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeRestoreTablesResponseBodyRestoreFull) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreTablesResponseBodyRestoreFull) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTablesResponseBodyRestoreFull) SetSucceed(v int32) *DescribeRestoreTablesResponseBodyRestoreFull {
	s.Succeed = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreFull) SetDataSize(v string) *DescribeRestoreTablesResponseBodyRestoreFull {
	s.DataSize = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreFull) SetSpeed(v string) *DescribeRestoreTablesResponseBodyRestoreFull {
	s.Speed = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreFull) SetRestoreFullDetails(v *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetails) *DescribeRestoreTablesResponseBodyRestoreFull {
	s.RestoreFullDetails = v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreFull) SetPageSize(v int32) *DescribeRestoreTablesResponseBodyRestoreFull {
	s.PageSize = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreFull) SetFail(v int32) *DescribeRestoreTablesResponseBodyRestoreFull {
	s.Fail = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreFull) SetPageNumber(v int32) *DescribeRestoreTablesResponseBodyRestoreFull {
	s.PageNumber = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreFull) SetTotal(v int64) *DescribeRestoreTablesResponseBodyRestoreFull {
	s.Total = &v
	return s
}

type DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetails struct {
	RestoreFullDetail []*DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail `json:"RestoreFullDetail,omitempty" xml:"RestoreFullDetail,omitempty" type:"Repeated"`
}

func (s DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetails) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetails) SetRestoreFullDetail(v []*DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetails {
	s.RestoreFullDetail = v
	return s
}

type DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Process   *string `json:"Process,omitempty" xml:"Process,omitempty"`
	DataSize  *string `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	Speed     *string `json:"Speed,omitempty" xml:"Speed,omitempty"`
	State     *string `json:"State,omitempty" xml:"State,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Table     *string `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetEndTime(v string) *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.EndTime = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetStartTime(v string) *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.StartTime = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetProcess(v string) *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.Process = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetDataSize(v string) *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.DataSize = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetSpeed(v string) *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.Speed = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetState(v string) *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.State = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetMessage(v string) *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.Message = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetTable(v string) *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.Table = &v
	return s
}

type DescribeRestoreTablesResponseBodyTables struct {
	Table []*string `json:"Table,omitempty" xml:"Table,omitempty" type:"Repeated"`
}

func (s DescribeRestoreTablesResponseBodyTables) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreTablesResponseBodyTables) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTablesResponseBodyTables) SetTable(v []*string) *DescribeRestoreTablesResponseBodyTables {
	s.Table = v
	return s
}

type DescribeRestoreTablesResponseBodyRestoreSummary struct {
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	State         *string `json:"State,omitempty" xml:"State,omitempty"`
	RecordId      *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	RestoreToDate *string `json:"RestoreToDate,omitempty" xml:"RestoreToDate,omitempty"`
	TargetCluster *string `json:"TargetCluster,omitempty" xml:"TargetCluster,omitempty"`
}

func (s DescribeRestoreTablesResponseBodyRestoreSummary) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreTablesResponseBodyRestoreSummary) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTablesResponseBodyRestoreSummary) SetEndTime(v string) *DescribeRestoreTablesResponseBodyRestoreSummary {
	s.EndTime = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSummary) SetStartTime(v string) *DescribeRestoreTablesResponseBodyRestoreSummary {
	s.StartTime = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSummary) SetState(v string) *DescribeRestoreTablesResponseBodyRestoreSummary {
	s.State = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSummary) SetRecordId(v string) *DescribeRestoreTablesResponseBodyRestoreSummary {
	s.RecordId = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSummary) SetRestoreToDate(v string) *DescribeRestoreTablesResponseBodyRestoreSummary {
	s.RestoreToDate = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSummary) SetTargetCluster(v string) *DescribeRestoreTablesResponseBodyRestoreSummary {
	s.TargetCluster = &v
	return s
}

type DescribeRestoreTablesResponseBodyRestoreIncrDetail struct {
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Process        *string `json:"Process,omitempty" xml:"Process,omitempty"`
	RestoreStartTs *string `json:"RestoreStartTs,omitempty" xml:"RestoreStartTs,omitempty"`
	State          *string `json:"State,omitempty" xml:"State,omitempty"`
	RestoredTs     *string `json:"RestoredTs,omitempty" xml:"RestoredTs,omitempty"`
	RestoreDelay   *string `json:"RestoreDelay,omitempty" xml:"RestoreDelay,omitempty"`
}

func (s DescribeRestoreTablesResponseBodyRestoreIncrDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreTablesResponseBodyRestoreIncrDetail) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTablesResponseBodyRestoreIncrDetail) SetEndTime(v string) *DescribeRestoreTablesResponseBodyRestoreIncrDetail {
	s.EndTime = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreIncrDetail) SetStartTime(v string) *DescribeRestoreTablesResponseBodyRestoreIncrDetail {
	s.StartTime = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreIncrDetail) SetProcess(v string) *DescribeRestoreTablesResponseBodyRestoreIncrDetail {
	s.Process = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreIncrDetail) SetRestoreStartTs(v string) *DescribeRestoreTablesResponseBodyRestoreIncrDetail {
	s.RestoreStartTs = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreIncrDetail) SetState(v string) *DescribeRestoreTablesResponseBodyRestoreIncrDetail {
	s.State = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreIncrDetail) SetRestoredTs(v string) *DescribeRestoreTablesResponseBodyRestoreIncrDetail {
	s.RestoredTs = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreIncrDetail) SetRestoreDelay(v string) *DescribeRestoreTablesResponseBodyRestoreIncrDetail {
	s.RestoreDelay = &v
	return s
}

type DescribeRestoreTablesResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRestoreTablesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRestoreTablesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreTablesResponse) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTablesResponse) SetHeaders(v map[string]*string) *DescribeRestoreTablesResponse {
	s.Headers = v
	return s
}

func (s *DescribeRestoreTablesResponse) SetBody(v *DescribeRestoreTablesResponseBody) *DescribeRestoreTablesResponse {
	s.Body = v
	return s
}

type DescribeSecurityGroupsRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeSecurityGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupsRequest) SetClusterId(v string) *DescribeSecurityGroupsRequest {
	s.ClusterId = &v
	return s
}

type DescribeSecurityGroupsResponseBody struct {
	RequestId        *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecurityGroupIds *DescribeSecurityGroupsResponseBodySecurityGroupIds `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Struct"`
}

func (s DescribeSecurityGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupsResponseBody) SetRequestId(v string) *DescribeSecurityGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityGroupsResponseBody) SetSecurityGroupIds(v *DescribeSecurityGroupsResponseBodySecurityGroupIds) *DescribeSecurityGroupsResponseBody {
	s.SecurityGroupIds = v
	return s
}

type DescribeSecurityGroupsResponseBodySecurityGroupIds struct {
	SecurityGroupId []*string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" type:"Repeated"`
}

func (s DescribeSecurityGroupsResponseBodySecurityGroupIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityGroupsResponseBodySecurityGroupIds) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupIds) SetSecurityGroupId(v []*string) *DescribeSecurityGroupsResponseBodySecurityGroupIds {
	s.SecurityGroupId = v
	return s
}

type DescribeSecurityGroupsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSecurityGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSecurityGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupsResponse) SetHeaders(v map[string]*string) *DescribeSecurityGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecurityGroupsResponse) SetBody(v *DescribeSecurityGroupsResponseBody) *DescribeSecurityGroupsResponse {
	s.Body = v
	return s
}

type DescribeSubDomainRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId   *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeSubDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubDomainRequest) GoString() string {
	return s.String()
}

func (s *DescribeSubDomainRequest) SetRegionId(v string) *DescribeSubDomainRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSubDomainRequest) SetZoneId(v string) *DescribeSubDomainRequest {
	s.ZoneId = &v
	return s
}

type DescribeSubDomainResponseBody struct {
	SubDomain *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSubDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSubDomainResponseBody) SetSubDomain(v string) *DescribeSubDomainResponseBody {
	s.SubDomain = &v
	return s
}

func (s *DescribeSubDomainResponseBody) SetRequestId(v string) *DescribeSubDomainResponseBody {
	s.RequestId = &v
	return s
}

type DescribeSubDomainResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSubDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSubDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubDomainResponse) GoString() string {
	return s.String()
}

func (s *DescribeSubDomainResponse) SetHeaders(v map[string]*string) *DescribeSubDomainResponse {
	s.Headers = v
	return s
}

func (s *DescribeSubDomainResponse) SetBody(v *DescribeSubDomainResponseBody) *DescribeSubDomainResponse {
	s.Body = v
	return s
}

type EnableHBaseueBackupRequest struct {
	HbaseueClusterId *string `json:"HbaseueClusterId,omitempty" xml:"HbaseueClusterId,omitempty"`
	NodeCount        *int32  `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	ColdStorageSize  *int32  `json:"ColdStorageSize,omitempty" xml:"ColdStorageSize,omitempty"`
	ClientToken      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s EnableHBaseueBackupRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableHBaseueBackupRequest) GoString() string {
	return s.String()
}

func (s *EnableHBaseueBackupRequest) SetHbaseueClusterId(v string) *EnableHBaseueBackupRequest {
	s.HbaseueClusterId = &v
	return s
}

func (s *EnableHBaseueBackupRequest) SetNodeCount(v int32) *EnableHBaseueBackupRequest {
	s.NodeCount = &v
	return s
}

func (s *EnableHBaseueBackupRequest) SetColdStorageSize(v int32) *EnableHBaseueBackupRequest {
	s.ColdStorageSize = &v
	return s
}

func (s *EnableHBaseueBackupRequest) SetClientToken(v string) *EnableHBaseueBackupRequest {
	s.ClientToken = &v
	return s
}

type EnableHBaseueBackupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s EnableHBaseueBackupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableHBaseueBackupResponseBody) GoString() string {
	return s.String()
}

func (s *EnableHBaseueBackupResponseBody) SetRequestId(v string) *EnableHBaseueBackupResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableHBaseueBackupResponseBody) SetClusterId(v string) *EnableHBaseueBackupResponseBody {
	s.ClusterId = &v
	return s
}

func (s *EnableHBaseueBackupResponseBody) SetOrderId(v string) *EnableHBaseueBackupResponseBody {
	s.OrderId = &v
	return s
}

type EnableHBaseueBackupResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableHBaseueBackupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableHBaseueBackupResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableHBaseueBackupResponse) GoString() string {
	return s.String()
}

func (s *EnableHBaseueBackupResponse) SetHeaders(v map[string]*string) *EnableHBaseueBackupResponse {
	s.Headers = v
	return s
}

func (s *EnableHBaseueBackupResponse) SetBody(v *EnableHBaseueBackupResponseBody) *EnableHBaseueBackupResponse {
	s.Body = v
	return s
}

type EnableHBaseueModuleRequest struct {
	ModuleClusterName  *string `json:"ModuleClusterName,omitempty" xml:"ModuleClusterName,omitempty"`
	VpcId              *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitchId          *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	MasterInstanceType *string `json:"MasterInstanceType,omitempty" xml:"MasterInstanceType,omitempty"`
	CoreInstanceType   *string `json:"CoreInstanceType,omitempty" xml:"CoreInstanceType,omitempty"`
	DiskType           *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	DiskSize           *int32  `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	NodeCount          *int32  `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	PayType            *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	PeriodUnit         *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	Period             *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	AutoRenewPeriod    *int32  `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	ClientToken        *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	HbaseueClusterId   *string `json:"HbaseueClusterId,omitempty" xml:"HbaseueClusterId,omitempty"`
	BdsId              *string `json:"BdsId,omitempty" xml:"BdsId,omitempty"`
	ModuleTypeName     *string `json:"ModuleTypeName,omitempty" xml:"ModuleTypeName,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId             *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s EnableHBaseueModuleRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableHBaseueModuleRequest) GoString() string {
	return s.String()
}

func (s *EnableHBaseueModuleRequest) SetModuleClusterName(v string) *EnableHBaseueModuleRequest {
	s.ModuleClusterName = &v
	return s
}

func (s *EnableHBaseueModuleRequest) SetVpcId(v string) *EnableHBaseueModuleRequest {
	s.VpcId = &v
	return s
}

func (s *EnableHBaseueModuleRequest) SetVswitchId(v string) *EnableHBaseueModuleRequest {
	s.VswitchId = &v
	return s
}

func (s *EnableHBaseueModuleRequest) SetMasterInstanceType(v string) *EnableHBaseueModuleRequest {
	s.MasterInstanceType = &v
	return s
}

func (s *EnableHBaseueModuleRequest) SetCoreInstanceType(v string) *EnableHBaseueModuleRequest {
	s.CoreInstanceType = &v
	return s
}

func (s *EnableHBaseueModuleRequest) SetDiskType(v string) *EnableHBaseueModuleRequest {
	s.DiskType = &v
	return s
}

func (s *EnableHBaseueModuleRequest) SetDiskSize(v int32) *EnableHBaseueModuleRequest {
	s.DiskSize = &v
	return s
}

func (s *EnableHBaseueModuleRequest) SetNodeCount(v int32) *EnableHBaseueModuleRequest {
	s.NodeCount = &v
	return s
}

func (s *EnableHBaseueModuleRequest) SetPayType(v string) *EnableHBaseueModuleRequest {
	s.PayType = &v
	return s
}

func (s *EnableHBaseueModuleRequest) SetPeriodUnit(v string) *EnableHBaseueModuleRequest {
	s.PeriodUnit = &v
	return s
}

func (s *EnableHBaseueModuleRequest) SetPeriod(v int32) *EnableHBaseueModuleRequest {
	s.Period = &v
	return s
}

func (s *EnableHBaseueModuleRequest) SetAutoRenewPeriod(v int32) *EnableHBaseueModuleRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *EnableHBaseueModuleRequest) SetClientToken(v string) *EnableHBaseueModuleRequest {
	s.ClientToken = &v
	return s
}

func (s *EnableHBaseueModuleRequest) SetHbaseueClusterId(v string) *EnableHBaseueModuleRequest {
	s.HbaseueClusterId = &v
	return s
}

func (s *EnableHBaseueModuleRequest) SetBdsId(v string) *EnableHBaseueModuleRequest {
	s.BdsId = &v
	return s
}

func (s *EnableHBaseueModuleRequest) SetModuleTypeName(v string) *EnableHBaseueModuleRequest {
	s.ModuleTypeName = &v
	return s
}

func (s *EnableHBaseueModuleRequest) SetRegionId(v string) *EnableHBaseueModuleRequest {
	s.RegionId = &v
	return s
}

func (s *EnableHBaseueModuleRequest) SetZoneId(v string) *EnableHBaseueModuleRequest {
	s.ZoneId = &v
	return s
}

type EnableHBaseueModuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s EnableHBaseueModuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableHBaseueModuleResponseBody) GoString() string {
	return s.String()
}

func (s *EnableHBaseueModuleResponseBody) SetRequestId(v string) *EnableHBaseueModuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableHBaseueModuleResponseBody) SetClusterId(v string) *EnableHBaseueModuleResponseBody {
	s.ClusterId = &v
	return s
}

func (s *EnableHBaseueModuleResponseBody) SetOrderId(v string) *EnableHBaseueModuleResponseBody {
	s.OrderId = &v
	return s
}

type EnableHBaseueModuleResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableHBaseueModuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableHBaseueModuleResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableHBaseueModuleResponse) GoString() string {
	return s.String()
}

func (s *EnableHBaseueModuleResponse) SetHeaders(v map[string]*string) *EnableHBaseueModuleResponse {
	s.Headers = v
	return s
}

func (s *EnableHBaseueModuleResponse) SetBody(v *EnableHBaseueModuleResponseBody) *EnableHBaseueModuleResponse {
	s.Body = v
	return s
}

type EvaluateMultiZoneResourceRequest struct {
	Engine               *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion        *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	ArchVersion          *string `json:"ArchVersion,omitempty" xml:"ArchVersion,omitempty"`
	ClusterName          *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VpcId                *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	MultiZoneCombination *string `json:"MultiZoneCombination,omitempty" xml:"MultiZoneCombination,omitempty"`
	PrimaryZoneId        *string `json:"PrimaryZoneId,omitempty" xml:"PrimaryZoneId,omitempty"`
	PrimaryVSwitchId     *string `json:"PrimaryVSwitchId,omitempty" xml:"PrimaryVSwitchId,omitempty"`
	StandbyZoneId        *string `json:"StandbyZoneId,omitempty" xml:"StandbyZoneId,omitempty"`
	StandbyVSwitchId     *string `json:"StandbyVSwitchId,omitempty" xml:"StandbyVSwitchId,omitempty"`
	ArbiterZoneId        *string `json:"ArbiterZoneId,omitempty" xml:"ArbiterZoneId,omitempty"`
	ArbiterVSwitchId     *string `json:"ArbiterVSwitchId,omitempty" xml:"ArbiterVSwitchId,omitempty"`
	MasterInstanceType   *string `json:"MasterInstanceType,omitempty" xml:"MasterInstanceType,omitempty"`
	CoreInstanceType     *string `json:"CoreInstanceType,omitempty" xml:"CoreInstanceType,omitempty"`
	CoreDiskType         *string `json:"CoreDiskType,omitempty" xml:"CoreDiskType,omitempty"`
	CoreDiskSize         *int32  `json:"CoreDiskSize,omitempty" xml:"CoreDiskSize,omitempty"`
	CoreNodeCount        *int32  `json:"CoreNodeCount,omitempty" xml:"CoreNodeCount,omitempty"`
	LogInstanceType      *string `json:"LogInstanceType,omitempty" xml:"LogInstanceType,omitempty"`
	LogDiskType          *string `json:"LogDiskType,omitempty" xml:"LogDiskType,omitempty"`
	LogDiskSize          *int32  `json:"LogDiskSize,omitempty" xml:"LogDiskSize,omitempty"`
	LogNodeCount         *int32  `json:"LogNodeCount,omitempty" xml:"LogNodeCount,omitempty"`
	SecurityIPList       *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	PayType              *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	PeriodUnit           *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	Period               *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	AutoRenewPeriod      *int32  `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s EvaluateMultiZoneResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s EvaluateMultiZoneResourceRequest) GoString() string {
	return s.String()
}

func (s *EvaluateMultiZoneResourceRequest) SetEngine(v string) *EvaluateMultiZoneResourceRequest {
	s.Engine = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetEngineVersion(v string) *EvaluateMultiZoneResourceRequest {
	s.EngineVersion = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetArchVersion(v string) *EvaluateMultiZoneResourceRequest {
	s.ArchVersion = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetClusterName(v string) *EvaluateMultiZoneResourceRequest {
	s.ClusterName = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetRegionId(v string) *EvaluateMultiZoneResourceRequest {
	s.RegionId = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetVpcId(v string) *EvaluateMultiZoneResourceRequest {
	s.VpcId = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetMultiZoneCombination(v string) *EvaluateMultiZoneResourceRequest {
	s.MultiZoneCombination = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetPrimaryZoneId(v string) *EvaluateMultiZoneResourceRequest {
	s.PrimaryZoneId = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetPrimaryVSwitchId(v string) *EvaluateMultiZoneResourceRequest {
	s.PrimaryVSwitchId = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetStandbyZoneId(v string) *EvaluateMultiZoneResourceRequest {
	s.StandbyZoneId = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetStandbyVSwitchId(v string) *EvaluateMultiZoneResourceRequest {
	s.StandbyVSwitchId = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetArbiterZoneId(v string) *EvaluateMultiZoneResourceRequest {
	s.ArbiterZoneId = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetArbiterVSwitchId(v string) *EvaluateMultiZoneResourceRequest {
	s.ArbiterVSwitchId = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetMasterInstanceType(v string) *EvaluateMultiZoneResourceRequest {
	s.MasterInstanceType = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetCoreInstanceType(v string) *EvaluateMultiZoneResourceRequest {
	s.CoreInstanceType = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetCoreDiskType(v string) *EvaluateMultiZoneResourceRequest {
	s.CoreDiskType = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetCoreDiskSize(v int32) *EvaluateMultiZoneResourceRequest {
	s.CoreDiskSize = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetCoreNodeCount(v int32) *EvaluateMultiZoneResourceRequest {
	s.CoreNodeCount = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetLogInstanceType(v string) *EvaluateMultiZoneResourceRequest {
	s.LogInstanceType = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetLogDiskType(v string) *EvaluateMultiZoneResourceRequest {
	s.LogDiskType = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetLogDiskSize(v int32) *EvaluateMultiZoneResourceRequest {
	s.LogDiskSize = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetLogNodeCount(v int32) *EvaluateMultiZoneResourceRequest {
	s.LogNodeCount = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetSecurityIPList(v string) *EvaluateMultiZoneResourceRequest {
	s.SecurityIPList = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetPayType(v string) *EvaluateMultiZoneResourceRequest {
	s.PayType = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetPeriodUnit(v string) *EvaluateMultiZoneResourceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetPeriod(v int32) *EvaluateMultiZoneResourceRequest {
	s.Period = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetAutoRenewPeriod(v int32) *EvaluateMultiZoneResourceRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetClientToken(v string) *EvaluateMultiZoneResourceRequest {
	s.ClientToken = &v
	return s
}

type EvaluateMultiZoneResourceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EvaluateMultiZoneResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EvaluateMultiZoneResourceResponseBody) GoString() string {
	return s.String()
}

func (s *EvaluateMultiZoneResourceResponseBody) SetRequestId(v string) *EvaluateMultiZoneResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *EvaluateMultiZoneResourceResponseBody) SetSuccess(v bool) *EvaluateMultiZoneResourceResponseBody {
	s.Success = &v
	return s
}

type EvaluateMultiZoneResourceResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EvaluateMultiZoneResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EvaluateMultiZoneResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s EvaluateMultiZoneResourceResponse) GoString() string {
	return s.String()
}

func (s *EvaluateMultiZoneResourceResponse) SetHeaders(v map[string]*string) *EvaluateMultiZoneResourceResponse {
	s.Headers = v
	return s
}

func (s *EvaluateMultiZoneResourceResponse) SetBody(v *EvaluateMultiZoneResourceResponseBody) *EvaluateMultiZoneResourceResponse {
	s.Body = v
	return s
}

type GetMultimodeCmsUrlRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *GetMultimodeCmsUrlRequest) SetRegionId(v string) *GetMultimodeCmsUrlRequest {
	s.RegionId = &v
	return s
}

type GetMultimodeCmsUrlResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ClusterId      *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	MultimodCmsUrl *string `json:"MultimodCmsUrl,omitempty" xml:"MultimodCmsUrl,omitempty"`
}

func (s GetMultimodeCmsUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMultimodeCmsUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetMultimodeCmsUrlResponseBody) SetRequestId(v string) *GetMultimodeCmsUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMultimodeCmsUrlResponseBody) SetClusterId(v string) *GetMultimodeCmsUrlResponseBody {
	s.ClusterId = &v
	return s
}

func (s *GetMultimodeCmsUrlResponseBody) SetMultimodCmsUrl(v string) *GetMultimodeCmsUrlResponseBody {
	s.MultimodCmsUrl = &v
	return s
}

type GetMultimodeCmsUrlResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetMultimodeCmsUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetMultimodeCmsUrlResponse) SetBody(v *GetMultimodeCmsUrlResponseBody) *GetMultimodeCmsUrlResponse {
	s.Body = v
	return s
}

type ListHBaseInstancesRequest struct {
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListHBaseInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHBaseInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListHBaseInstancesRequest) SetVpcId(v string) *ListHBaseInstancesRequest {
	s.VpcId = &v
	return s
}

type ListHBaseInstancesResponseBody struct {
	Instances *ListHBaseInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListHBaseInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHBaseInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListHBaseInstancesResponseBody) SetInstances(v *ListHBaseInstancesResponseBodyInstances) *ListHBaseInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *ListHBaseInstancesResponseBody) SetRequestId(v string) *ListHBaseInstancesResponseBody {
	s.RequestId = &v
	return s
}

type ListHBaseInstancesResponseBodyInstances struct {
	Instance []*ListHBaseInstancesResponseBodyInstancesInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s ListHBaseInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s ListHBaseInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListHBaseInstancesResponseBodyInstances) SetInstance(v []*ListHBaseInstancesResponseBodyInstancesInstance) *ListHBaseInstancesResponseBodyInstances {
	s.Instance = v
	return s
}

type ListHBaseInstancesResponseBodyInstancesInstance struct {
	IsDefault    *bool   `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListHBaseInstancesResponseBodyInstancesInstance) String() string {
	return tea.Prettify(s)
}

func (s ListHBaseInstancesResponseBodyInstancesInstance) GoString() string {
	return s.String()
}

func (s *ListHBaseInstancesResponseBodyInstancesInstance) SetIsDefault(v bool) *ListHBaseInstancesResponseBodyInstancesInstance {
	s.IsDefault = &v
	return s
}

func (s *ListHBaseInstancesResponseBodyInstancesInstance) SetInstanceName(v string) *ListHBaseInstancesResponseBodyInstancesInstance {
	s.InstanceName = &v
	return s
}

func (s *ListHBaseInstancesResponseBodyInstancesInstance) SetInstanceId(v string) *ListHBaseInstancesResponseBodyInstancesInstance {
	s.InstanceId = &v
	return s
}

type ListHBaseInstancesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListHBaseInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListHBaseInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHBaseInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListHBaseInstancesResponse) SetHeaders(v map[string]*string) *ListHBaseInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListHBaseInstancesResponse) SetBody(v *ListHBaseInstancesResponseBody) *ListHBaseInstancesResponse {
	s.Body = v
	return s
}

type ListInstanceServiceConfigHistoriesRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s ListInstanceServiceConfigHistoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceServiceConfigHistoriesRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceServiceConfigHistoriesRequest) SetClusterId(v string) *ListInstanceServiceConfigHistoriesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListInstanceServiceConfigHistoriesRequest) SetPageSize(v int32) *ListInstanceServiceConfigHistoriesRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstanceServiceConfigHistoriesRequest) SetPageNumber(v int32) *ListInstanceServiceConfigHistoriesRequest {
	s.PageNumber = &v
	return s
}

type ListInstanceServiceConfigHistoriesResponseBody struct {
	TotalRecordCount     *int64                                                              `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	PageRecordCount      *int32                                                              `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId            *string                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber           *int32                                                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	ConfigureHistoryList *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryList `json:"ConfigureHistoryList,omitempty" xml:"ConfigureHistoryList,omitempty" type:"Struct"`
}

func (s ListInstanceServiceConfigHistoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceServiceConfigHistoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceServiceConfigHistoriesResponseBody) SetTotalRecordCount(v int64) *ListInstanceServiceConfigHistoriesResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *ListInstanceServiceConfigHistoriesResponseBody) SetPageRecordCount(v int32) *ListInstanceServiceConfigHistoriesResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *ListInstanceServiceConfigHistoriesResponseBody) SetRequestId(v string) *ListInstanceServiceConfigHistoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceServiceConfigHistoriesResponseBody) SetPageNumber(v int32) *ListInstanceServiceConfigHistoriesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListInstanceServiceConfigHistoriesResponseBody) SetConfigureHistoryList(v *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryList) *ListInstanceServiceConfigHistoriesResponseBody {
	s.ConfigureHistoryList = v
	return s
}

type ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryList struct {
	Config []*ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Repeated"`
}

func (s ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryList) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryList) GoString() string {
	return s.String()
}

func (s *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryList) SetConfig(v []*ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig) *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryList {
	s.Config = v
	return s
}

type ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig struct {
	Effective     *string `json:"Effective,omitempty" xml:"Effective,omitempty"`
	OldValue      *string `json:"OldValue,omitempty" xml:"OldValue,omitempty"`
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ConfigureName *string `json:"ConfigureName,omitempty" xml:"ConfigureName,omitempty"`
	NewValue      *string `json:"NewValue,omitempty" xml:"NewValue,omitempty"`
}

func (s ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig) GoString() string {
	return s.String()
}

func (s *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig) SetEffective(v string) *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig {
	s.Effective = &v
	return s
}

func (s *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig) SetOldValue(v string) *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig {
	s.OldValue = &v
	return s
}

func (s *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig) SetCreateTime(v string) *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig {
	s.CreateTime = &v
	return s
}

func (s *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig) SetConfigureName(v string) *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig {
	s.ConfigureName = &v
	return s
}

func (s *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig) SetNewValue(v string) *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig {
	s.NewValue = &v
	return s
}

type ListInstanceServiceConfigHistoriesResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInstanceServiceConfigHistoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInstanceServiceConfigHistoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceServiceConfigHistoriesResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceServiceConfigHistoriesResponse) SetHeaders(v map[string]*string) *ListInstanceServiceConfigHistoriesResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceServiceConfigHistoriesResponse) SetBody(v *ListInstanceServiceConfigHistoriesResponseBody) *ListInstanceServiceConfigHistoriesResponse {
	s.Body = v
	return s
}

type ListInstanceServiceConfigurationsRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s ListInstanceServiceConfigurationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceServiceConfigurationsRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceServiceConfigurationsRequest) SetClusterId(v string) *ListInstanceServiceConfigurationsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListInstanceServiceConfigurationsRequest) SetPageSize(v int32) *ListInstanceServiceConfigurationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstanceServiceConfigurationsRequest) SetPageNumber(v int32) *ListInstanceServiceConfigurationsRequest {
	s.PageNumber = &v
	return s
}

type ListInstanceServiceConfigurationsResponseBody struct {
	TotalRecordCount *int64                                                      `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	PageRecordCount  *int32                                                      `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId        *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber       *int32                                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	ConfigureList    *ListInstanceServiceConfigurationsResponseBodyConfigureList `json:"ConfigureList,omitempty" xml:"ConfigureList,omitempty" type:"Struct"`
}

func (s ListInstanceServiceConfigurationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceServiceConfigurationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceServiceConfigurationsResponseBody) SetTotalRecordCount(v int64) *ListInstanceServiceConfigurationsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *ListInstanceServiceConfigurationsResponseBody) SetPageRecordCount(v int32) *ListInstanceServiceConfigurationsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *ListInstanceServiceConfigurationsResponseBody) SetRequestId(v string) *ListInstanceServiceConfigurationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceServiceConfigurationsResponseBody) SetPageNumber(v int32) *ListInstanceServiceConfigurationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListInstanceServiceConfigurationsResponseBody) SetConfigureList(v *ListInstanceServiceConfigurationsResponseBodyConfigureList) *ListInstanceServiceConfigurationsResponseBody {
	s.ConfigureList = v
	return s
}

type ListInstanceServiceConfigurationsResponseBodyConfigureList struct {
	Config []*ListInstanceServiceConfigurationsResponseBodyConfigureListConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Repeated"`
}

func (s ListInstanceServiceConfigurationsResponseBodyConfigureList) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceServiceConfigurationsResponseBodyConfigureList) GoString() string {
	return s.String()
}

func (s *ListInstanceServiceConfigurationsResponseBodyConfigureList) SetConfig(v []*ListInstanceServiceConfigurationsResponseBodyConfigureListConfig) *ListInstanceServiceConfigurationsResponseBodyConfigureList {
	s.Config = v
	return s
}

type ListInstanceServiceConfigurationsResponseBodyConfigureListConfig struct {
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RunningValue  *string `json:"RunningValue,omitempty" xml:"RunningValue,omitempty"`
	ConfigureUnit *string `json:"ConfigureUnit,omitempty" xml:"ConfigureUnit,omitempty"`
	ConfigureName *string `json:"ConfigureName,omitempty" xml:"ConfigureName,omitempty"`
	ValueRange    *string `json:"ValueRange,omitempty" xml:"ValueRange,omitempty"`
	DefaultValue  *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	NeedRestart   *string `json:"NeedRestart,omitempty" xml:"NeedRestart,omitempty"`
}

func (s ListInstanceServiceConfigurationsResponseBodyConfigureListConfig) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceServiceConfigurationsResponseBodyConfigureListConfig) GoString() string {
	return s.String()
}

func (s *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig) SetDescription(v string) *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig {
	s.Description = &v
	return s
}

func (s *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig) SetRunningValue(v string) *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig {
	s.RunningValue = &v
	return s
}

func (s *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig) SetConfigureUnit(v string) *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig {
	s.ConfigureUnit = &v
	return s
}

func (s *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig) SetConfigureName(v string) *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig {
	s.ConfigureName = &v
	return s
}

func (s *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig) SetValueRange(v string) *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig {
	s.ValueRange = &v
	return s
}

func (s *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig) SetDefaultValue(v string) *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig {
	s.DefaultValue = &v
	return s
}

func (s *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig) SetNeedRestart(v string) *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig {
	s.NeedRestart = &v
	return s
}

type ListInstanceServiceConfigurationsResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInstanceServiceConfigurationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInstanceServiceConfigurationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceServiceConfigurationsResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceServiceConfigurationsResponse) SetHeaders(v map[string]*string) *ListInstanceServiceConfigurationsResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceServiceConfigurationsResponse) SetBody(v *ListInstanceServiceConfigurationsResponseBody) *ListInstanceServiceConfigurationsResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	RegionId   *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	NextToken  *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceId []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	Tag        []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
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
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ListTagsRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagsRequest) GoString() string {
	return s.String()
}

func (s *ListTagsRequest) SetRegionId(v string) *ListTagsRequest {
	s.RegionId = &v
	return s
}

type ListTagsResponseBody struct {
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tags      *ListTagsResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s ListTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagsResponseBody) SetRequestId(v string) *ListTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagsResponseBody) SetTags(v *ListTagsResponseBodyTags) *ListTagsResponseBody {
	s.Tags = v
	return s
}

type ListTagsResponseBodyTags struct {
	Tag []*ListTagsResponseBodyTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagsResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s ListTagsResponseBodyTags) GoString() string {
	return s.String()
}

func (s *ListTagsResponseBodyTags) SetTag(v []*ListTagsResponseBodyTagsTag) *ListTagsResponseBodyTags {
	s.Tag = v
	return s
}

type ListTagsResponseBodyTagsTag struct {
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagsResponseBodyTagsTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagsResponseBodyTagsTag) GoString() string {
	return s.String()
}

func (s *ListTagsResponseBodyTagsTag) SetTagValue(v string) *ListTagsResponseBodyTagsTag {
	s.TagValue = &v
	return s
}

func (s *ListTagsResponseBodyTagsTag) SetTagKey(v string) *ListTagsResponseBodyTagsTag {
	s.TagKey = &v
	return s
}

type ListTagsResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagsResponse) GoString() string {
	return s.String()
}

func (s *ListTagsResponse) SetHeaders(v map[string]*string) *ListTagsResponse {
	s.Headers = v
	return s
}

func (s *ListTagsResponse) SetBody(v *ListTagsResponseBody) *ListTagsResponse {
	s.Body = v
	return s
}

type ModifyBackupPlanConfigRequest struct {
	ClusterId           *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Tables              *string `json:"Tables,omitempty" xml:"Tables,omitempty"`
	FullBackupCycle     *string `json:"FullBackupCycle,omitempty" xml:"FullBackupCycle,omitempty"`
	MinHFileBackupCount *string `json:"MinHFileBackupCount,omitempty" xml:"MinHFileBackupCount,omitempty"`
	NextFullBackupDate  *string `json:"NextFullBackupDate,omitempty" xml:"NextFullBackupDate,omitempty"`
}

func (s ModifyBackupPlanConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupPlanConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyBackupPlanConfigRequest) SetClusterId(v string) *ModifyBackupPlanConfigRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyBackupPlanConfigRequest) SetTables(v string) *ModifyBackupPlanConfigRequest {
	s.Tables = &v
	return s
}

func (s *ModifyBackupPlanConfigRequest) SetFullBackupCycle(v string) *ModifyBackupPlanConfigRequest {
	s.FullBackupCycle = &v
	return s
}

func (s *ModifyBackupPlanConfigRequest) SetMinHFileBackupCount(v string) *ModifyBackupPlanConfigRequest {
	s.MinHFileBackupCount = &v
	return s
}

func (s *ModifyBackupPlanConfigRequest) SetNextFullBackupDate(v string) *ModifyBackupPlanConfigRequest {
	s.NextFullBackupDate = &v
	return s
}

type ModifyBackupPlanConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyBackupPlanConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupPlanConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBackupPlanConfigResponseBody) SetRequestId(v string) *ModifyBackupPlanConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyBackupPlanConfigResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyBackupPlanConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyBackupPlanConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupPlanConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyBackupPlanConfigResponse) SetHeaders(v map[string]*string) *ModifyBackupPlanConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyBackupPlanConfigResponse) SetBody(v *ModifyBackupPlanConfigResponseBody) *ModifyBackupPlanConfigResponse {
	s.Body = v
	return s
}

type ModifyBackupPolicyRequest struct {
	ClusterId                   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PreferredBackupTime         *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	PreferredBackupPeriod       *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	PreferredBackupStartTimeUTC *string `json:"PreferredBackupStartTimeUTC,omitempty" xml:"PreferredBackupStartTimeUTC,omitempty"`
	PreferredBackupEndTimeUTC   *string `json:"PreferredBackupEndTimeUTC,omitempty" xml:"PreferredBackupEndTimeUTC,omitempty"`
}

func (s ModifyBackupPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyRequest) SetClusterId(v string) *ModifyBackupPolicyRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupTime(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupTime = &v
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

func (s *ModifyBackupPolicyRequest) SetPreferredBackupEndTimeUTC(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupEndTimeUTC = &v
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyBackupPolicyResponse) SetBody(v *ModifyBackupPolicyResponseBody) *ModifyBackupPolicyResponse {
	s.Body = v
	return s
}

type ModifyClusterDeletionProtectionRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Protection *bool   `json:"Protection,omitempty" xml:"Protection,omitempty"`
}

func (s ModifyClusterDeletionProtectionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterDeletionProtectionRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterDeletionProtectionRequest) SetClusterId(v string) *ModifyClusterDeletionProtectionRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyClusterDeletionProtectionRequest) SetProtection(v bool) *ModifyClusterDeletionProtectionRequest {
	s.Protection = &v
	return s
}

type ModifyClusterDeletionProtectionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyClusterDeletionProtectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterDeletionProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterDeletionProtectionResponseBody) SetRequestId(v string) *ModifyClusterDeletionProtectionResponseBody {
	s.RequestId = &v
	return s
}

type ModifyClusterDeletionProtectionResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyClusterDeletionProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyClusterDeletionProtectionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterDeletionProtectionResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterDeletionProtectionResponse) SetHeaders(v map[string]*string) *ModifyClusterDeletionProtectionResponse {
	s.Headers = v
	return s
}

func (s *ModifyClusterDeletionProtectionResponse) SetBody(v *ModifyClusterDeletionProtectionResponseBody) *ModifyClusterDeletionProtectionResponse {
	s.Body = v
	return s
}

type ModifyDiskWarningLineRequest struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	WarningLine *int32  `json:"WarningLine,omitempty" xml:"WarningLine,omitempty"`
}

func (s ModifyDiskWarningLineRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDiskWarningLineRequest) GoString() string {
	return s.String()
}

func (s *ModifyDiskWarningLineRequest) SetClusterId(v string) *ModifyDiskWarningLineRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyDiskWarningLineRequest) SetWarningLine(v int32) *ModifyDiskWarningLineRequest {
	s.WarningLine = &v
	return s
}

type ModifyDiskWarningLineResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDiskWarningLineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDiskWarningLineResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDiskWarningLineResponseBody) SetRequestId(v string) *ModifyDiskWarningLineResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDiskWarningLineResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDiskWarningLineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDiskWarningLineResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDiskWarningLineResponse) GoString() string {
	return s.String()
}

func (s *ModifyDiskWarningLineResponse) SetHeaders(v map[string]*string) *ModifyDiskWarningLineResponse {
	s.Headers = v
	return s
}

func (s *ModifyDiskWarningLineResponse) SetBody(v *ModifyDiskWarningLineResponseBody) *ModifyDiskWarningLineResponse {
	s.Body = v
	return s
}

type ModifyInstanceMaintainTimeRequest struct {
	ClusterId         *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	MaintainEndTime   *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
}

func (s ModifyInstanceMaintainTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceMaintainTimeRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMaintainTimeRequest) SetClusterId(v string) *ModifyInstanceMaintainTimeRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyInstanceMaintainTimeRequest) SetMaintainStartTime(v string) *ModifyInstanceMaintainTimeRequest {
	s.MaintainStartTime = &v
	return s
}

func (s *ModifyInstanceMaintainTimeRequest) SetMaintainEndTime(v string) *ModifyInstanceMaintainTimeRequest {
	s.MaintainEndTime = &v
	return s
}

type ModifyInstanceMaintainTimeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceMaintainTimeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceMaintainTimeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMaintainTimeResponseBody) SetRequestId(v string) *ModifyInstanceMaintainTimeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyInstanceMaintainTimeResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyInstanceMaintainTimeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyInstanceMaintainTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceMaintainTimeResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMaintainTimeResponse) SetHeaders(v map[string]*string) *ModifyInstanceMaintainTimeResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceMaintainTimeResponse) SetBody(v *ModifyInstanceMaintainTimeResponseBody) *ModifyInstanceMaintainTimeResponse {
	s.Body = v
	return s
}

type ModifyInstanceNameRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId      *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
}

func (s ModifyInstanceNameRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceNameRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceNameRequest) SetClientToken(v string) *ModifyInstanceNameRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyInstanceNameRequest) SetRegionId(v string) *ModifyInstanceNameRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceNameRequest) SetZoneId(v string) *ModifyInstanceNameRequest {
	s.ZoneId = &v
	return s
}

func (s *ModifyInstanceNameRequest) SetClusterId(v string) *ModifyInstanceNameRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyInstanceNameRequest) SetClusterName(v string) *ModifyInstanceNameRequest {
	s.ClusterName = &v
	return s
}

type ModifyInstanceNameResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceNameResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceNameResponseBody) SetRequestId(v string) *ModifyInstanceNameResponseBody {
	s.RequestId = &v
	return s
}

type ModifyInstanceNameResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyInstanceNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyInstanceNameResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceNameResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceNameResponse) SetHeaders(v map[string]*string) *ModifyInstanceNameResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceNameResponse) SetBody(v *ModifyInstanceNameResponseBody) *ModifyInstanceNameResponse {
	s.Body = v
	return s
}

type ModifyInstanceServiceConfigRequest struct {
	ClusterId      *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Restart        *bool   `json:"Restart,omitempty" xml:"Restart,omitempty"`
	ConfigureName  *string `json:"ConfigureName,omitempty" xml:"ConfigureName,omitempty"`
	ConfigureValue *string `json:"ConfigureValue,omitempty" xml:"ConfigureValue,omitempty"`
	Parameters     *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
}

func (s ModifyInstanceServiceConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceServiceConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceServiceConfigRequest) SetClusterId(v string) *ModifyInstanceServiceConfigRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyInstanceServiceConfigRequest) SetRestart(v bool) *ModifyInstanceServiceConfigRequest {
	s.Restart = &v
	return s
}

func (s *ModifyInstanceServiceConfigRequest) SetConfigureName(v string) *ModifyInstanceServiceConfigRequest {
	s.ConfigureName = &v
	return s
}

func (s *ModifyInstanceServiceConfigRequest) SetConfigureValue(v string) *ModifyInstanceServiceConfigRequest {
	s.ConfigureValue = &v
	return s
}

func (s *ModifyInstanceServiceConfigRequest) SetParameters(v string) *ModifyInstanceServiceConfigRequest {
	s.Parameters = &v
	return s
}

type ModifyInstanceServiceConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceServiceConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceServiceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceServiceConfigResponseBody) SetRequestId(v string) *ModifyInstanceServiceConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyInstanceServiceConfigResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyInstanceServiceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyInstanceServiceConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceServiceConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceServiceConfigResponse) SetHeaders(v map[string]*string) *ModifyInstanceServiceConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceServiceConfigResponse) SetBody(v *ModifyInstanceServiceConfigResponseBody) *ModifyInstanceServiceConfigResponse {
	s.Body = v
	return s
}

type ModifyInstanceTypeRequest struct {
	ClusterId          *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	MasterInstanceType *string `json:"MasterInstanceType,omitempty" xml:"MasterInstanceType,omitempty"`
	CoreInstanceType   *string `json:"CoreInstanceType,omitempty" xml:"CoreInstanceType,omitempty"`
}

func (s ModifyInstanceTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceTypeRequest) SetClusterId(v string) *ModifyInstanceTypeRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyInstanceTypeRequest) SetMasterInstanceType(v string) *ModifyInstanceTypeRequest {
	s.MasterInstanceType = &v
	return s
}

func (s *ModifyInstanceTypeRequest) SetCoreInstanceType(v string) *ModifyInstanceTypeRequest {
	s.CoreInstanceType = &v
	return s
}

type ModifyInstanceTypeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s ModifyInstanceTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceTypeResponseBody) SetRequestId(v string) *ModifyInstanceTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceTypeResponseBody) SetOrderId(v string) *ModifyInstanceTypeResponseBody {
	s.OrderId = &v
	return s
}

type ModifyInstanceTypeResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyInstanceTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyInstanceTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceTypeResponse) SetHeaders(v map[string]*string) *ModifyInstanceTypeResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceTypeResponse) SetBody(v *ModifyInstanceTypeResponseBody) *ModifyInstanceTypeResponse {
	s.Body = v
	return s
}

type ModifyIpWhitelistRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	IpList    *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
}

func (s ModifyIpWhitelistRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyIpWhitelistRequest) GoString() string {
	return s.String()
}

func (s *ModifyIpWhitelistRequest) SetClusterId(v string) *ModifyIpWhitelistRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyIpWhitelistRequest) SetIpList(v string) *ModifyIpWhitelistRequest {
	s.IpList = &v
	return s
}

func (s *ModifyIpWhitelistRequest) SetGroupName(v string) *ModifyIpWhitelistRequest {
	s.GroupName = &v
	return s
}

func (s *ModifyIpWhitelistRequest) SetIpVersion(v string) *ModifyIpWhitelistRequest {
	s.IpVersion = &v
	return s
}

type ModifyIpWhitelistResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyIpWhitelistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyIpWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyIpWhitelistResponseBody) SetRequestId(v string) *ModifyIpWhitelistResponseBody {
	s.RequestId = &v
	return s
}

type ModifyIpWhitelistResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyIpWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyIpWhitelistResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyIpWhitelistResponse) GoString() string {
	return s.String()
}

func (s *ModifyIpWhitelistResponse) SetHeaders(v map[string]*string) *ModifyIpWhitelistResponse {
	s.Headers = v
	return s
}

func (s *ModifyIpWhitelistResponse) SetBody(v *ModifyIpWhitelistResponseBody) *ModifyIpWhitelistResponse {
	s.Body = v
	return s
}

type ModifyMultiZoneClusterNodeTypeRequest struct {
	ClusterId          *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	MasterInstanceType *string `json:"MasterInstanceType,omitempty" xml:"MasterInstanceType,omitempty"`
	CoreInstanceType   *string `json:"CoreInstanceType,omitempty" xml:"CoreInstanceType,omitempty"`
	LogInstanceType    *string `json:"LogInstanceType,omitempty" xml:"LogInstanceType,omitempty"`
}

func (s ModifyMultiZoneClusterNodeTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyMultiZoneClusterNodeTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyMultiZoneClusterNodeTypeRequest) SetClusterId(v string) *ModifyMultiZoneClusterNodeTypeRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyMultiZoneClusterNodeTypeRequest) SetMasterInstanceType(v string) *ModifyMultiZoneClusterNodeTypeRequest {
	s.MasterInstanceType = &v
	return s
}

func (s *ModifyMultiZoneClusterNodeTypeRequest) SetCoreInstanceType(v string) *ModifyMultiZoneClusterNodeTypeRequest {
	s.CoreInstanceType = &v
	return s
}

func (s *ModifyMultiZoneClusterNodeTypeRequest) SetLogInstanceType(v string) *ModifyMultiZoneClusterNodeTypeRequest {
	s.LogInstanceType = &v
	return s
}

type ModifyMultiZoneClusterNodeTypeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s ModifyMultiZoneClusterNodeTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyMultiZoneClusterNodeTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMultiZoneClusterNodeTypeResponseBody) SetRequestId(v string) *ModifyMultiZoneClusterNodeTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyMultiZoneClusterNodeTypeResponseBody) SetOrderId(v string) *ModifyMultiZoneClusterNodeTypeResponseBody {
	s.OrderId = &v
	return s
}

type ModifyMultiZoneClusterNodeTypeResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyMultiZoneClusterNodeTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyMultiZoneClusterNodeTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyMultiZoneClusterNodeTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifyMultiZoneClusterNodeTypeResponse) SetHeaders(v map[string]*string) *ModifyMultiZoneClusterNodeTypeResponse {
	s.Headers = v
	return s
}

func (s *ModifyMultiZoneClusterNodeTypeResponse) SetBody(v *ModifyMultiZoneClusterNodeTypeResponseBody) *ModifyMultiZoneClusterNodeTypeResponse {
	s.Body = v
	return s
}

type ModifySecurityGroupsRequest struct {
	ClusterId        *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	SecurityGroupIds *string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty"`
}

func (s ModifySecurityGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySecurityGroupsRequest) GoString() string {
	return s.String()
}

func (s *ModifySecurityGroupsRequest) SetClusterId(v string) *ModifySecurityGroupsRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifySecurityGroupsRequest) SetSecurityGroupIds(v string) *ModifySecurityGroupsRequest {
	s.SecurityGroupIds = &v
	return s
}

type ModifySecurityGroupsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySecurityGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySecurityGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySecurityGroupsResponseBody) SetRequestId(v string) *ModifySecurityGroupsResponseBody {
	s.RequestId = &v
	return s
}

type ModifySecurityGroupsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifySecurityGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifySecurityGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifySecurityGroupsResponse) GoString() string {
	return s.String()
}

func (s *ModifySecurityGroupsResponse) SetHeaders(v map[string]*string) *ModifySecurityGroupsResponse {
	s.Headers = v
	return s
}

func (s *ModifySecurityGroupsResponse) SetBody(v *ModifySecurityGroupsResponseBody) *ModifySecurityGroupsResponse {
	s.Body = v
	return s
}

type ModifyUIAccountPasswordRequest struct {
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	AccountName     *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
}

func (s ModifyUIAccountPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyUIAccountPasswordRequest) GoString() string {
	return s.String()
}

func (s *ModifyUIAccountPasswordRequest) SetClusterId(v string) *ModifyUIAccountPasswordRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyUIAccountPasswordRequest) SetAccountName(v string) *ModifyUIAccountPasswordRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyUIAccountPasswordRequest) SetAccountPassword(v string) *ModifyUIAccountPasswordRequest {
	s.AccountPassword = &v
	return s
}

type ModifyUIAccountPasswordResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyUIAccountPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyUIAccountPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUIAccountPasswordResponseBody) SetRequestId(v string) *ModifyUIAccountPasswordResponseBody {
	s.RequestId = &v
	return s
}

type ModifyUIAccountPasswordResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyUIAccountPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyUIAccountPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyUIAccountPasswordResponse) GoString() string {
	return s.String()
}

func (s *ModifyUIAccountPasswordResponse) SetHeaders(v map[string]*string) *ModifyUIAccountPasswordResponse {
	s.Headers = v
	return s
}

func (s *ModifyUIAccountPasswordResponse) SetBody(v *ModifyUIAccountPasswordResponseBody) *ModifyUIAccountPasswordResponse {
	s.Body = v
	return s
}

type MoveResourceGroupRequest struct {
	ClusterId          *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
}

func (s MoveResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupRequest) SetClusterId(v string) *MoveResourceGroupRequest {
	s.ClusterId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetNewResourceGroupId(v string) *MoveResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

type MoveResourceGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MoveResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponseBody) SetRequestId(v string) *MoveResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type MoveResourceGroupResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MoveResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MoveResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponse) SetHeaders(v map[string]*string) *MoveResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *MoveResourceGroupResponse) SetBody(v *MoveResourceGroupResponseBody) *MoveResourceGroupResponse {
	s.Body = v
	return s
}

type OpenBackupRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
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
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OpenBackupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *OpenBackupResponse) SetBody(v *OpenBackupResponseBody) *OpenBackupResponse {
	s.Body = v
	return s
}

type PurgeInstanceRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s PurgeInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s PurgeInstanceRequest) GoString() string {
	return s.String()
}

func (s *PurgeInstanceRequest) SetClusterId(v string) *PurgeInstanceRequest {
	s.ClusterId = &v
	return s
}

type PurgeInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PurgeInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PurgeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *PurgeInstanceResponseBody) SetRequestId(v string) *PurgeInstanceResponseBody {
	s.RequestId = &v
	return s
}

type PurgeInstanceResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PurgeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PurgeInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s PurgeInstanceResponse) GoString() string {
	return s.String()
}

func (s *PurgeInstanceResponse) SetHeaders(v map[string]*string) *PurgeInstanceResponse {
	s.Headers = v
	return s
}

func (s *PurgeInstanceResponse) SetBody(v *PurgeInstanceResponseBody) *PurgeInstanceResponse {
	s.Body = v
	return s
}

type QueryHBaseHaDBRequest struct {
	BdsId *string `json:"BdsId,omitempty" xml:"BdsId,omitempty"`
}

func (s QueryHBaseHaDBRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryHBaseHaDBRequest) GoString() string {
	return s.String()
}

func (s *QueryHBaseHaDBRequest) SetBdsId(v string) *QueryHBaseHaDBRequest {
	s.BdsId = &v
	return s
}

type QueryHBaseHaDBResponseBody struct {
	TotalCount  *int64                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	ClusterList *QueryHBaseHaDBResponseBodyClusterList `json:"ClusterList,omitempty" xml:"ClusterList,omitempty" type:"Struct"`
	PageSize    *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber  *int32                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s QueryHBaseHaDBResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryHBaseHaDBResponseBody) GoString() string {
	return s.String()
}

func (s *QueryHBaseHaDBResponseBody) SetTotalCount(v int64) *QueryHBaseHaDBResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryHBaseHaDBResponseBody) SetClusterList(v *QueryHBaseHaDBResponseBodyClusterList) *QueryHBaseHaDBResponseBody {
	s.ClusterList = v
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

func (s *QueryHBaseHaDBResponseBody) SetPageNumber(v int32) *QueryHBaseHaDBResponseBody {
	s.PageNumber = &v
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
	HaName        *string                                                    `json:"HaName,omitempty" xml:"HaName,omitempty"`
	BdsName       *string                                                    `json:"BdsName,omitempty" xml:"BdsName,omitempty"`
	HaSlbConnList *QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnList `json:"HaSlbConnList,omitempty" xml:"HaSlbConnList,omitempty" type:"Struct"`
	ActiveName    *string                                                    `json:"ActiveName,omitempty" xml:"ActiveName,omitempty"`
	StandbyName   *string                                                    `json:"StandbyName,omitempty" xml:"StandbyName,omitempty"`
}

func (s QueryHBaseHaDBResponseBodyClusterListCluster) String() string {
	return tea.Prettify(s)
}

func (s QueryHBaseHaDBResponseBodyClusterListCluster) GoString() string {
	return s.String()
}

func (s *QueryHBaseHaDBResponseBodyClusterListCluster) SetHaName(v string) *QueryHBaseHaDBResponseBodyClusterListCluster {
	s.HaName = &v
	return s
}

func (s *QueryHBaseHaDBResponseBodyClusterListCluster) SetBdsName(v string) *QueryHBaseHaDBResponseBodyClusterListCluster {
	s.BdsName = &v
	return s
}

func (s *QueryHBaseHaDBResponseBodyClusterListCluster) SetHaSlbConnList(v *QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnList) *QueryHBaseHaDBResponseBodyClusterListCluster {
	s.HaSlbConnList = v
	return s
}

func (s *QueryHBaseHaDBResponseBodyClusterListCluster) SetActiveName(v string) *QueryHBaseHaDBResponseBodyClusterListCluster {
	s.ActiveName = &v
	return s
}

func (s *QueryHBaseHaDBResponseBodyClusterListCluster) SetStandbyName(v string) *QueryHBaseHaDBResponseBodyClusterListCluster {
	s.StandbyName = &v
	return s
}

type QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnList struct {
	HaSlbConn []*QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnListHaSlbConn `json:"HaSlbConn,omitempty" xml:"HaSlbConn,omitempty" type:"Repeated"`
}

func (s QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnList) String() string {
	return tea.Prettify(s)
}

func (s QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnList) GoString() string {
	return s.String()
}

func (s *QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnList) SetHaSlbConn(v []*QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnListHaSlbConn) *QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnList {
	s.HaSlbConn = v
	return s
}

type QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnListHaSlbConn struct {
	SlbConnAddr *string `json:"SlbConnAddr,omitempty" xml:"SlbConnAddr,omitempty"`
	HbaseType   *string `json:"HbaseType,omitempty" xml:"HbaseType,omitempty"`
	SlbType     *string `json:"SlbType,omitempty" xml:"SlbType,omitempty"`
}

func (s QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnListHaSlbConn) String() string {
	return tea.Prettify(s)
}

func (s QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnListHaSlbConn) GoString() string {
	return s.String()
}

func (s *QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnListHaSlbConn) SetSlbConnAddr(v string) *QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnListHaSlbConn {
	s.SlbConnAddr = &v
	return s
}

func (s *QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnListHaSlbConn) SetHbaseType(v string) *QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnListHaSlbConn {
	s.HbaseType = &v
	return s
}

func (s *QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnListHaSlbConn) SetSlbType(v string) *QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnListHaSlbConn {
	s.SlbType = &v
	return s
}

type QueryHBaseHaDBResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryHBaseHaDBResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryHBaseHaDBResponse) SetBody(v *QueryHBaseHaDBResponseBody) *QueryHBaseHaDBResponse {
	s.Body = v
	return s
}

type QueryXpackRelateDBRequest struct {
	ClusterId     *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RelateDbType  *string `json:"RelateDbType,omitempty" xml:"RelateDbType,omitempty"`
	HasSingleNode *bool   `json:"HasSingleNode,omitempty" xml:"HasSingleNode,omitempty"`
}

func (s QueryXpackRelateDBRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryXpackRelateDBRequest) GoString() string {
	return s.String()
}

func (s *QueryXpackRelateDBRequest) SetClusterId(v string) *QueryXpackRelateDBRequest {
	s.ClusterId = &v
	return s
}

func (s *QueryXpackRelateDBRequest) SetRelateDbType(v string) *QueryXpackRelateDBRequest {
	s.RelateDbType = &v
	return s
}

func (s *QueryXpackRelateDBRequest) SetHasSingleNode(v bool) *QueryXpackRelateDBRequest {
	s.HasSingleNode = &v
	return s
}

type QueryXpackRelateDBResponseBody struct {
	ClusterList *QueryXpackRelateDBResponseBodyClusterList `json:"ClusterList,omitempty" xml:"ClusterList,omitempty" type:"Struct"`
	RequestId   *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryXpackRelateDBResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryXpackRelateDBResponseBody) GoString() string {
	return s.String()
}

func (s *QueryXpackRelateDBResponseBody) SetClusterList(v *QueryXpackRelateDBResponseBodyClusterList) *QueryXpackRelateDBResponseBody {
	s.ClusterList = v
	return s
}

func (s *QueryXpackRelateDBResponseBody) SetRequestId(v string) *QueryXpackRelateDBResponseBody {
	s.RequestId = &v
	return s
}

type QueryXpackRelateDBResponseBodyClusterList struct {
	Cluster []*QueryXpackRelateDBResponseBodyClusterListCluster `json:"Cluster,omitempty" xml:"Cluster,omitempty" type:"Repeated"`
}

func (s QueryXpackRelateDBResponseBodyClusterList) String() string {
	return tea.Prettify(s)
}

func (s QueryXpackRelateDBResponseBodyClusterList) GoString() string {
	return s.String()
}

func (s *QueryXpackRelateDBResponseBodyClusterList) SetCluster(v []*QueryXpackRelateDBResponseBodyClusterListCluster) *QueryXpackRelateDBResponseBodyClusterList {
	s.Cluster = v
	return s
}

type QueryXpackRelateDBResponseBodyClusterListCluster struct {
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	DBVersion   *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	IsRelated   *bool   `json:"IsRelated,omitempty" xml:"IsRelated,omitempty"`
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	DBType      *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	LockMode    *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s QueryXpackRelateDBResponseBodyClusterListCluster) String() string {
	return tea.Prettify(s)
}

func (s QueryXpackRelateDBResponseBodyClusterListCluster) GoString() string {
	return s.String()
}

func (s *QueryXpackRelateDBResponseBodyClusterListCluster) SetStatus(v string) *QueryXpackRelateDBResponseBodyClusterListCluster {
	s.Status = &v
	return s
}

func (s *QueryXpackRelateDBResponseBodyClusterListCluster) SetDBVersion(v string) *QueryXpackRelateDBResponseBodyClusterListCluster {
	s.DBVersion = &v
	return s
}

func (s *QueryXpackRelateDBResponseBodyClusterListCluster) SetIsRelated(v bool) *QueryXpackRelateDBResponseBodyClusterListCluster {
	s.IsRelated = &v
	return s
}

func (s *QueryXpackRelateDBResponseBodyClusterListCluster) SetClusterName(v string) *QueryXpackRelateDBResponseBodyClusterListCluster {
	s.ClusterName = &v
	return s
}

func (s *QueryXpackRelateDBResponseBodyClusterListCluster) SetDBType(v string) *QueryXpackRelateDBResponseBodyClusterListCluster {
	s.DBType = &v
	return s
}

func (s *QueryXpackRelateDBResponseBodyClusterListCluster) SetLockMode(v string) *QueryXpackRelateDBResponseBodyClusterListCluster {
	s.LockMode = &v
	return s
}

func (s *QueryXpackRelateDBResponseBodyClusterListCluster) SetClusterId(v string) *QueryXpackRelateDBResponseBodyClusterListCluster {
	s.ClusterId = &v
	return s
}

type QueryXpackRelateDBResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryXpackRelateDBResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryXpackRelateDBResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryXpackRelateDBResponse) GoString() string {
	return s.String()
}

func (s *QueryXpackRelateDBResponse) SetHeaders(v map[string]*string) *QueryXpackRelateDBResponse {
	s.Headers = v
	return s
}

func (s *QueryXpackRelateDBResponse) SetBody(v *QueryXpackRelateDBResponseBody) *QueryXpackRelateDBResponse {
	s.Body = v
	return s
}

type RelateDbForHBaseHaRequest struct {
	ClusterId           *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	HaTables            *string `json:"HaTables,omitempty" xml:"HaTables,omitempty"`
	HaActiveHdfsUri     *string `json:"HaActiveHdfsUri,omitempty" xml:"HaActiveHdfsUri,omitempty"`
	HaActiveHbaseFsDir  *string `json:"HaActiveHbaseFsDir,omitempty" xml:"HaActiveHbaseFsDir,omitempty"`
	HaActiveClusterKey  *string `json:"HaActiveClusterKey,omitempty" xml:"HaActiveClusterKey,omitempty"`
	HaActiveVersion     *string `json:"HaActiveVersion,omitempty" xml:"HaActiveVersion,omitempty"`
	HaActiveUser        *string `json:"HaActiveUser,omitempty" xml:"HaActiveUser,omitempty"`
	HaActivePassword    *string `json:"HaActivePassword,omitempty" xml:"HaActivePassword,omitempty"`
	HaStandbyHdfsUri    *string `json:"HaStandbyHdfsUri,omitempty" xml:"HaStandbyHdfsUri,omitempty"`
	HaStandbyHbaseFsDir *string `json:"HaStandbyHbaseFsDir,omitempty" xml:"HaStandbyHbaseFsDir,omitempty"`
	HaStandbyClusterKey *string `json:"HaStandbyClusterKey,omitempty" xml:"HaStandbyClusterKey,omitempty"`
	HaStandbyVersion    *string `json:"HaStandbyVersion,omitempty" xml:"HaStandbyVersion,omitempty"`
	HaStandbyUser       *string `json:"HaStandbyUser,omitempty" xml:"HaStandbyUser,omitempty"`
	HaStandbyPassword   *string `json:"HaStandbyPassword,omitempty" xml:"HaStandbyPassword,omitempty"`
	HaActive            *string `json:"HaActive,omitempty" xml:"HaActive,omitempty"`
	HaStandby           *string `json:"HaStandby,omitempty" xml:"HaStandby,omitempty"`
	HaActiveDBType      *string `json:"HaActiveDBType,omitempty" xml:"HaActiveDBType,omitempty"`
	HaStandbyDBType     *string `json:"HaStandbyDBType,omitempty" xml:"HaStandbyDBType,omitempty"`
	IsActiveStandard    *bool   `json:"IsActiveStandard,omitempty" xml:"IsActiveStandard,omitempty"`
	IsStandbyStandard   *bool   `json:"IsStandbyStandard,omitempty" xml:"IsStandbyStandard,omitempty"`
	HaMigrateType       *string `json:"HaMigrateType,omitempty" xml:"HaMigrateType,omitempty"`
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

func (s *RelateDbForHBaseHaRequest) SetHaTables(v string) *RelateDbForHBaseHaRequest {
	s.HaTables = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaActiveHdfsUri(v string) *RelateDbForHBaseHaRequest {
	s.HaActiveHdfsUri = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaActiveHbaseFsDir(v string) *RelateDbForHBaseHaRequest {
	s.HaActiveHbaseFsDir = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaActiveClusterKey(v string) *RelateDbForHBaseHaRequest {
	s.HaActiveClusterKey = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaActiveVersion(v string) *RelateDbForHBaseHaRequest {
	s.HaActiveVersion = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaActiveUser(v string) *RelateDbForHBaseHaRequest {
	s.HaActiveUser = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaActivePassword(v string) *RelateDbForHBaseHaRequest {
	s.HaActivePassword = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaStandbyHdfsUri(v string) *RelateDbForHBaseHaRequest {
	s.HaStandbyHdfsUri = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaStandbyHbaseFsDir(v string) *RelateDbForHBaseHaRequest {
	s.HaStandbyHbaseFsDir = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaStandbyClusterKey(v string) *RelateDbForHBaseHaRequest {
	s.HaStandbyClusterKey = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaStandbyVersion(v string) *RelateDbForHBaseHaRequest {
	s.HaStandbyVersion = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaStandbyUser(v string) *RelateDbForHBaseHaRequest {
	s.HaStandbyUser = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaStandbyPassword(v string) *RelateDbForHBaseHaRequest {
	s.HaStandbyPassword = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaActive(v string) *RelateDbForHBaseHaRequest {
	s.HaActive = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaStandby(v string) *RelateDbForHBaseHaRequest {
	s.HaStandby = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaActiveDBType(v string) *RelateDbForHBaseHaRequest {
	s.HaActiveDBType = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaStandbyDBType(v string) *RelateDbForHBaseHaRequest {
	s.HaStandbyDBType = &v
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

func (s *RelateDbForHBaseHaRequest) SetHaMigrateType(v string) *RelateDbForHBaseHaRequest {
	s.HaMigrateType = &v
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RelateDbForHBaseHaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RelateDbForHBaseHaResponse) SetBody(v *RelateDbForHBaseHaResponseBody) *RelateDbForHBaseHaResponse {
	s.Body = v
	return s
}

type ReleasePublicNetworkAddressRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
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
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReleasePublicNetworkAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ReleasePublicNetworkAddressResponse) SetBody(v *ReleasePublicNetworkAddressResponseBody) *ReleasePublicNetworkAddressResponse {
	s.Body = v
	return s
}

type RenewInstanceRequest struct {
	ClusterId    *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	Duration     *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
}

func (s RenewInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RenewInstanceRequest) GoString() string {
	return s.String()
}

func (s *RenewInstanceRequest) SetClusterId(v string) *RenewInstanceRequest {
	s.ClusterId = &v
	return s
}

func (s *RenewInstanceRequest) SetPricingCycle(v string) *RenewInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *RenewInstanceRequest) SetDuration(v int32) *RenewInstanceRequest {
	s.Duration = &v
	return s
}

type RenewInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId   *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s RenewInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenewInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RenewInstanceResponseBody) SetRequestId(v string) *RenewInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewInstanceResponseBody) SetOrderId(v int64) *RenewInstanceResponseBody {
	s.OrderId = &v
	return s
}

type RenewInstanceResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RenewInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RenewInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RenewInstanceResponse) GoString() string {
	return s.String()
}

func (s *RenewInstanceResponse) SetHeaders(v map[string]*string) *RenewInstanceResponse {
	s.Headers = v
	return s
}

func (s *RenewInstanceResponse) SetBody(v *RenewInstanceResponseBody) *RenewInstanceResponse {
	s.Body = v
	return s
}

type ResizeColdStorageSizeRequest struct {
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ColdStorageSize *int32  `json:"ColdStorageSize,omitempty" xml:"ColdStorageSize,omitempty"`
}

func (s ResizeColdStorageSizeRequest) String() string {
	return tea.Prettify(s)
}

func (s ResizeColdStorageSizeRequest) GoString() string {
	return s.String()
}

func (s *ResizeColdStorageSizeRequest) SetClusterId(v string) *ResizeColdStorageSizeRequest {
	s.ClusterId = &v
	return s
}

func (s *ResizeColdStorageSizeRequest) SetColdStorageSize(v int32) *ResizeColdStorageSizeRequest {
	s.ColdStorageSize = &v
	return s
}

type ResizeColdStorageSizeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s ResizeColdStorageSizeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResizeColdStorageSizeResponseBody) GoString() string {
	return s.String()
}

func (s *ResizeColdStorageSizeResponseBody) SetRequestId(v string) *ResizeColdStorageSizeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResizeColdStorageSizeResponseBody) SetOrderId(v string) *ResizeColdStorageSizeResponseBody {
	s.OrderId = &v
	return s
}

type ResizeColdStorageSizeResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResizeColdStorageSizeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResizeColdStorageSizeResponse) String() string {
	return tea.Prettify(s)
}

func (s ResizeColdStorageSizeResponse) GoString() string {
	return s.String()
}

func (s *ResizeColdStorageSizeResponse) SetHeaders(v map[string]*string) *ResizeColdStorageSizeResponse {
	s.Headers = v
	return s
}

func (s *ResizeColdStorageSizeResponse) SetBody(v *ResizeColdStorageSizeResponseBody) *ResizeColdStorageSizeResponse {
	s.Body = v
	return s
}

type ResizeDiskSizeRequest struct {
	ClusterId    *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	NodeDiskSize *int32  `json:"NodeDiskSize,omitempty" xml:"NodeDiskSize,omitempty"`
}

func (s ResizeDiskSizeRequest) String() string {
	return tea.Prettify(s)
}

func (s ResizeDiskSizeRequest) GoString() string {
	return s.String()
}

func (s *ResizeDiskSizeRequest) SetClusterId(v string) *ResizeDiskSizeRequest {
	s.ClusterId = &v
	return s
}

func (s *ResizeDiskSizeRequest) SetNodeDiskSize(v int32) *ResizeDiskSizeRequest {
	s.NodeDiskSize = &v
	return s
}

type ResizeDiskSizeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s ResizeDiskSizeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResizeDiskSizeResponseBody) GoString() string {
	return s.String()
}

func (s *ResizeDiskSizeResponseBody) SetRequestId(v string) *ResizeDiskSizeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResizeDiskSizeResponseBody) SetOrderId(v string) *ResizeDiskSizeResponseBody {
	s.OrderId = &v
	return s
}

type ResizeDiskSizeResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResizeDiskSizeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResizeDiskSizeResponse) String() string {
	return tea.Prettify(s)
}

func (s ResizeDiskSizeResponse) GoString() string {
	return s.String()
}

func (s *ResizeDiskSizeResponse) SetHeaders(v map[string]*string) *ResizeDiskSizeResponse {
	s.Headers = v
	return s
}

func (s *ResizeDiskSizeResponse) SetBody(v *ResizeDiskSizeResponseBody) *ResizeDiskSizeResponse {
	s.Body = v
	return s
}

type ResizeMultiZoneClusterDiskSizeRequest struct {
	ClusterId    *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CoreDiskSize *int32  `json:"CoreDiskSize,omitempty" xml:"CoreDiskSize,omitempty"`
	LogDiskSize  *int32  `json:"LogDiskSize,omitempty" xml:"LogDiskSize,omitempty"`
}

func (s ResizeMultiZoneClusterDiskSizeRequest) String() string {
	return tea.Prettify(s)
}

func (s ResizeMultiZoneClusterDiskSizeRequest) GoString() string {
	return s.String()
}

func (s *ResizeMultiZoneClusterDiskSizeRequest) SetClusterId(v string) *ResizeMultiZoneClusterDiskSizeRequest {
	s.ClusterId = &v
	return s
}

func (s *ResizeMultiZoneClusterDiskSizeRequest) SetCoreDiskSize(v int32) *ResizeMultiZoneClusterDiskSizeRequest {
	s.CoreDiskSize = &v
	return s
}

func (s *ResizeMultiZoneClusterDiskSizeRequest) SetLogDiskSize(v int32) *ResizeMultiZoneClusterDiskSizeRequest {
	s.LogDiskSize = &v
	return s
}

type ResizeMultiZoneClusterDiskSizeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s ResizeMultiZoneClusterDiskSizeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResizeMultiZoneClusterDiskSizeResponseBody) GoString() string {
	return s.String()
}

func (s *ResizeMultiZoneClusterDiskSizeResponseBody) SetRequestId(v string) *ResizeMultiZoneClusterDiskSizeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResizeMultiZoneClusterDiskSizeResponseBody) SetOrderId(v string) *ResizeMultiZoneClusterDiskSizeResponseBody {
	s.OrderId = &v
	return s
}

type ResizeMultiZoneClusterDiskSizeResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResizeMultiZoneClusterDiskSizeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResizeMultiZoneClusterDiskSizeResponse) String() string {
	return tea.Prettify(s)
}

func (s ResizeMultiZoneClusterDiskSizeResponse) GoString() string {
	return s.String()
}

func (s *ResizeMultiZoneClusterDiskSizeResponse) SetHeaders(v map[string]*string) *ResizeMultiZoneClusterDiskSizeResponse {
	s.Headers = v
	return s
}

func (s *ResizeMultiZoneClusterDiskSizeResponse) SetBody(v *ResizeMultiZoneClusterDiskSizeResponseBody) *ResizeMultiZoneClusterDiskSizeResponse {
	s.Body = v
	return s
}

type ResizeMultiZoneClusterNodeCountRequest struct {
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PrimaryVSwitchId     *string `json:"PrimaryVSwitchId,omitempty" xml:"PrimaryVSwitchId,omitempty"`
	StandbyVSwitchId     *string `json:"StandbyVSwitchId,omitempty" xml:"StandbyVSwitchId,omitempty"`
	ArbiterVSwitchId     *string `json:"ArbiterVSwitchId,omitempty" xml:"ArbiterVSwitchId,omitempty"`
	CoreNodeCount        *int32  `json:"CoreNodeCount,omitempty" xml:"CoreNodeCount,omitempty"`
	PrimaryCoreNodeCount *int32  `json:"PrimaryCoreNodeCount,omitempty" xml:"PrimaryCoreNodeCount,omitempty"`
	StandbyCoreNodeCount *int32  `json:"StandbyCoreNodeCount,omitempty" xml:"StandbyCoreNodeCount,omitempty"`
	LogNodeCount         *int32  `json:"LogNodeCount,omitempty" xml:"LogNodeCount,omitempty"`
}

func (s ResizeMultiZoneClusterNodeCountRequest) String() string {
	return tea.Prettify(s)
}

func (s ResizeMultiZoneClusterNodeCountRequest) GoString() string {
	return s.String()
}

func (s *ResizeMultiZoneClusterNodeCountRequest) SetClusterId(v string) *ResizeMultiZoneClusterNodeCountRequest {
	s.ClusterId = &v
	return s
}

func (s *ResizeMultiZoneClusterNodeCountRequest) SetPrimaryVSwitchId(v string) *ResizeMultiZoneClusterNodeCountRequest {
	s.PrimaryVSwitchId = &v
	return s
}

func (s *ResizeMultiZoneClusterNodeCountRequest) SetStandbyVSwitchId(v string) *ResizeMultiZoneClusterNodeCountRequest {
	s.StandbyVSwitchId = &v
	return s
}

func (s *ResizeMultiZoneClusterNodeCountRequest) SetArbiterVSwitchId(v string) *ResizeMultiZoneClusterNodeCountRequest {
	s.ArbiterVSwitchId = &v
	return s
}

func (s *ResizeMultiZoneClusterNodeCountRequest) SetCoreNodeCount(v int32) *ResizeMultiZoneClusterNodeCountRequest {
	s.CoreNodeCount = &v
	return s
}

func (s *ResizeMultiZoneClusterNodeCountRequest) SetPrimaryCoreNodeCount(v int32) *ResizeMultiZoneClusterNodeCountRequest {
	s.PrimaryCoreNodeCount = &v
	return s
}

func (s *ResizeMultiZoneClusterNodeCountRequest) SetStandbyCoreNodeCount(v int32) *ResizeMultiZoneClusterNodeCountRequest {
	s.StandbyCoreNodeCount = &v
	return s
}

func (s *ResizeMultiZoneClusterNodeCountRequest) SetLogNodeCount(v int32) *ResizeMultiZoneClusterNodeCountRequest {
	s.LogNodeCount = &v
	return s
}

type ResizeMultiZoneClusterNodeCountResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s ResizeMultiZoneClusterNodeCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResizeMultiZoneClusterNodeCountResponseBody) GoString() string {
	return s.String()
}

func (s *ResizeMultiZoneClusterNodeCountResponseBody) SetRequestId(v string) *ResizeMultiZoneClusterNodeCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResizeMultiZoneClusterNodeCountResponseBody) SetOrderId(v string) *ResizeMultiZoneClusterNodeCountResponseBody {
	s.OrderId = &v
	return s
}

type ResizeMultiZoneClusterNodeCountResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResizeMultiZoneClusterNodeCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResizeMultiZoneClusterNodeCountResponse) String() string {
	return tea.Prettify(s)
}

func (s ResizeMultiZoneClusterNodeCountResponse) GoString() string {
	return s.String()
}

func (s *ResizeMultiZoneClusterNodeCountResponse) SetHeaders(v map[string]*string) *ResizeMultiZoneClusterNodeCountResponse {
	s.Headers = v
	return s
}

func (s *ResizeMultiZoneClusterNodeCountResponse) SetBody(v *ResizeMultiZoneClusterNodeCountResponseBody) *ResizeMultiZoneClusterNodeCountResponse {
	s.Body = v
	return s
}

type ResizeNodeCountRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	NodeCount *int32  `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	ZoneId    *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s ResizeNodeCountRequest) String() string {
	return tea.Prettify(s)
}

func (s ResizeNodeCountRequest) GoString() string {
	return s.String()
}

func (s *ResizeNodeCountRequest) SetClusterId(v string) *ResizeNodeCountRequest {
	s.ClusterId = &v
	return s
}

func (s *ResizeNodeCountRequest) SetNodeCount(v int32) *ResizeNodeCountRequest {
	s.NodeCount = &v
	return s
}

func (s *ResizeNodeCountRequest) SetZoneId(v string) *ResizeNodeCountRequest {
	s.ZoneId = &v
	return s
}

func (s *ResizeNodeCountRequest) SetVSwitchId(v string) *ResizeNodeCountRequest {
	s.VSwitchId = &v
	return s
}

type ResizeNodeCountResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s ResizeNodeCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResizeNodeCountResponseBody) GoString() string {
	return s.String()
}

func (s *ResizeNodeCountResponseBody) SetRequestId(v string) *ResizeNodeCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResizeNodeCountResponseBody) SetOrderId(v string) *ResizeNodeCountResponseBody {
	s.OrderId = &v
	return s
}

type ResizeNodeCountResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResizeNodeCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResizeNodeCountResponse) String() string {
	return tea.Prettify(s)
}

func (s ResizeNodeCountResponse) GoString() string {
	return s.String()
}

func (s *ResizeNodeCountResponse) SetHeaders(v map[string]*string) *ResizeNodeCountResponse {
	s.Headers = v
	return s
}

func (s *ResizeNodeCountResponse) SetBody(v *ResizeNodeCountResponseBody) *ResizeNodeCountResponse {
	s.Body = v
	return s
}

type RestartInstanceRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Components *string `json:"Components,omitempty" xml:"Components,omitempty"`
}

func (s RestartInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestartInstanceRequest) SetClusterId(v string) *RestartInstanceRequest {
	s.ClusterId = &v
	return s
}

func (s *RestartInstanceRequest) SetComponents(v string) *RestartInstanceRequest {
	s.Components = &v
	return s
}

type RestartInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestartInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RestartInstanceResponseBody) SetRequestId(v string) *RestartInstanceResponseBody {
	s.RequestId = &v
	return s
}

type RestartInstanceResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RestartInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RestartInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartInstanceResponse) GoString() string {
	return s.String()
}

func (s *RestartInstanceResponse) SetHeaders(v map[string]*string) *RestartInstanceResponse {
	s.Headers = v
	return s
}

func (s *RestartInstanceResponse) SetBody(v *RestartInstanceResponseBody) *RestartInstanceResponse {
	s.Body = v
	return s
}

type SwitchHbaseHaSlbRequest struct {
	BdsId     *string `json:"BdsId,omitempty" xml:"BdsId,omitempty"`
	HaId      *string `json:"HaId,omitempty" xml:"HaId,omitempty"`
	HaTypes   *string `json:"HaTypes,omitempty" xml:"HaTypes,omitempty"`
	HbaseType *string `json:"HbaseType,omitempty" xml:"HbaseType,omitempty"`
}

func (s SwitchHbaseHaSlbRequest) String() string {
	return tea.Prettify(s)
}

func (s SwitchHbaseHaSlbRequest) GoString() string {
	return s.String()
}

func (s *SwitchHbaseHaSlbRequest) SetBdsId(v string) *SwitchHbaseHaSlbRequest {
	s.BdsId = &v
	return s
}

func (s *SwitchHbaseHaSlbRequest) SetHaId(v string) *SwitchHbaseHaSlbRequest {
	s.HaId = &v
	return s
}

func (s *SwitchHbaseHaSlbRequest) SetHaTypes(v string) *SwitchHbaseHaSlbRequest {
	s.HaTypes = &v
	return s
}

func (s *SwitchHbaseHaSlbRequest) SetHbaseType(v string) *SwitchHbaseHaSlbRequest {
	s.HbaseType = &v
	return s
}

type SwitchHbaseHaSlbResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SwitchHbaseHaSlbResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SwitchHbaseHaSlbResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchHbaseHaSlbResponseBody) SetRequestId(v string) *SwitchHbaseHaSlbResponseBody {
	s.RequestId = &v
	return s
}

type SwitchHbaseHaSlbResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SwitchHbaseHaSlbResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SwitchHbaseHaSlbResponse) String() string {
	return tea.Prettify(s)
}

func (s SwitchHbaseHaSlbResponse) GoString() string {
	return s.String()
}

func (s *SwitchHbaseHaSlbResponse) SetHeaders(v map[string]*string) *SwitchHbaseHaSlbResponse {
	s.Headers = v
	return s
}

func (s *SwitchHbaseHaSlbResponse) SetBody(v *SwitchHbaseHaSlbResponseBody) *SwitchHbaseHaSlbResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	RegionId   *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	Tag        []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
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
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UnTagResourcesRequest struct {
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	All        *bool     `json:"All,omitempty" xml:"All,omitempty"`
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	TagKey     []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UnTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UnTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UnTagResourcesRequest) SetRegionId(v string) *UnTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UnTagResourcesRequest) SetAll(v bool) *UnTagResourcesRequest {
	s.All = &v
	return s
}

func (s *UnTagResourcesRequest) SetResourceId(v []*string) *UnTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UnTagResourcesRequest) SetTagKey(v []*string) *UnTagResourcesRequest {
	s.TagKey = v
	return s
}

type UnTagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UnTagResourcesResponseBody) SetRequestId(v string) *UnTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UnTagResourcesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UnTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UnTagResourcesResponse) SetHeaders(v map[string]*string) *UnTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UnTagResourcesResponse) SetBody(v *UnTagResourcesResponseBody) *UnTagResourcesResponse {
	s.Body = v
	return s
}

type UpgradeMinorVersionRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Components *string `json:"Components,omitempty" xml:"Components,omitempty"`
}

func (s UpgradeMinorVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeMinorVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeMinorVersionRequest) SetClusterId(v string) *UpgradeMinorVersionRequest {
	s.ClusterId = &v
	return s
}

func (s *UpgradeMinorVersionRequest) SetComponents(v string) *UpgradeMinorVersionRequest {
	s.Components = &v
	return s
}

type UpgradeMinorVersionResponseBody struct {
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UpgradingComponents *string `json:"UpgradingComponents,omitempty" xml:"UpgradingComponents,omitempty"`
}

func (s UpgradeMinorVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeMinorVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeMinorVersionResponseBody) SetRequestId(v string) *UpgradeMinorVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeMinorVersionResponseBody) SetUpgradingComponents(v string) *UpgradeMinorVersionResponseBody {
	s.UpgradingComponents = &v
	return s
}

type UpgradeMinorVersionResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpgradeMinorVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpgradeMinorVersionResponse) SetBody(v *UpgradeMinorVersionResponseBody) *UpgradeMinorVersionResponse {
	s.Body = v
	return s
}

type UpgradeMultiZoneClusterRequest struct {
	ClusterId         *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RunMode           *string `json:"RunMode,omitempty" xml:"RunMode,omitempty"`
	UpgradeInsName    *string `json:"UpgradeInsName,omitempty" xml:"UpgradeInsName,omitempty"`
	Components        *string `json:"Components,omitempty" xml:"Components,omitempty"`
	Versions          *string `json:"Versions,omitempty" xml:"Versions,omitempty"`
	RestartComponents *string `json:"RestartComponents,omitempty" xml:"RestartComponents,omitempty"`
}

func (s UpgradeMultiZoneClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeMultiZoneClusterRequest) GoString() string {
	return s.String()
}

func (s *UpgradeMultiZoneClusterRequest) SetClusterId(v string) *UpgradeMultiZoneClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *UpgradeMultiZoneClusterRequest) SetRunMode(v string) *UpgradeMultiZoneClusterRequest {
	s.RunMode = &v
	return s
}

func (s *UpgradeMultiZoneClusterRequest) SetUpgradeInsName(v string) *UpgradeMultiZoneClusterRequest {
	s.UpgradeInsName = &v
	return s
}

func (s *UpgradeMultiZoneClusterRequest) SetComponents(v string) *UpgradeMultiZoneClusterRequest {
	s.Components = &v
	return s
}

func (s *UpgradeMultiZoneClusterRequest) SetVersions(v string) *UpgradeMultiZoneClusterRequest {
	s.Versions = &v
	return s
}

func (s *UpgradeMultiZoneClusterRequest) SetRestartComponents(v string) *UpgradeMultiZoneClusterRequest {
	s.RestartComponents = &v
	return s
}

type UpgradeMultiZoneClusterResponseBody struct {
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UpgradingComponents *string `json:"UpgradingComponents,omitempty" xml:"UpgradingComponents,omitempty"`
}

func (s UpgradeMultiZoneClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeMultiZoneClusterResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeMultiZoneClusterResponseBody) SetRequestId(v string) *UpgradeMultiZoneClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeMultiZoneClusterResponseBody) SetUpgradingComponents(v string) *UpgradeMultiZoneClusterResponseBody {
	s.UpgradingComponents = &v
	return s
}

type UpgradeMultiZoneClusterResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpgradeMultiZoneClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpgradeMultiZoneClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeMultiZoneClusterResponse) GoString() string {
	return s.String()
}

func (s *UpgradeMultiZoneClusterResponse) SetHeaders(v map[string]*string) *UpgradeMultiZoneClusterResponse {
	s.Headers = v
	return s
}

func (s *UpgradeMultiZoneClusterResponse) SetBody(v *UpgradeMultiZoneClusterResponseBody) *UpgradeMultiZoneClusterResponse {
	s.Body = v
	return s
}

type XpackRelateDBRequest struct {
	ClusterId    *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DbClusterIds *string `json:"DbClusterIds,omitempty" xml:"DbClusterIds,omitempty"`
	RelateDbType *string `json:"RelateDbType,omitempty" xml:"RelateDbType,omitempty"`
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

func (s *XpackRelateDBRequest) SetRelateDbType(v string) *XpackRelateDBRequest {
	s.RelateDbType = &v
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *XpackRelateDBResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
		"ap-northeast-2-pop":          tea.String("hbase.aliyuncs.com"),
		"ap-southeast-1":              tea.String("hbase.aliyuncs.com"),
		"cn-beijing":                  tea.String("hbase.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("hbase.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("hbase.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("hbase.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("hbase.aliyuncs.com"),
		"cn-edge-1":                   tea.String("hbase.aliyuncs.com"),
		"cn-fujian":                   tea.String("hbase.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("hbase.aliyuncs.com"),
		"cn-hangzhou":                 tea.String("hbase.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("hbase.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("hbase.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("hbase.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("hbase.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("hbase.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("hbase.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("hbase.aliyuncs.com"),
		"cn-hongkong":                 tea.String("hbase.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("hbase.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("hbase.aliyuncs.com"),
		"cn-qingdao":                  tea.String("hbase.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("hbase.aliyuncs.com"),
		"cn-shanghai":                 tea.String("hbase.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("hbase.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("hbase.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("hbase.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("hbase.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("hbase.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("hbase.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("hbase.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("hbase.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("hbase.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("hbase.aliyuncs.com"),
		"cn-wuhan":                    tea.String("hbase.aliyuncs.com"),
		"cn-yushanfang":               tea.String("hbase.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("hbase.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("hbase.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("hbase.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("hbase.ap-northeast-1.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("hbase.ap-northeast-1.aliyuncs.com"),
		"us-east-1":                   tea.String("hbase.aliyuncs.com"),
		"us-west-1":                   tea.String("hbase.aliyuncs.com"),
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddUserHdfsInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddUserHdfsInfo"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AllocatePublicNetworkAddressResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AllocatePublicNetworkAddress"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) CheckComponentsVersionWithOptions(request *CheckComponentsVersionRequest, runtime *util.RuntimeOptions) (_result *CheckComponentsVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CheckComponentsVersionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CheckComponentsVersion"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckComponentsVersion(request *CheckComponentsVersionRequest) (_result *CheckComponentsVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckComponentsVersionResponse{}
	_body, _err := client.CheckComponentsVersionWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CloseBackupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CloseBackup"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ConvertInstanceWithOptions(request *ConvertInstanceRequest, runtime *util.RuntimeOptions) (_result *ConvertInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConvertInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConvertInstance"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConvertInstance(request *ConvertInstanceRequest) (_result *ConvertInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConvertInstanceResponse{}
	_body, _err := client.ConvertInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateBackupPlanWithOptions(request *CreateBackupPlanRequest, runtime *util.RuntimeOptions) (_result *CreateBackupPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateBackupPlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateBackupPlan"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateBackupPlan(request *CreateBackupPlanRequest) (_result *CreateBackupPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateBackupPlanResponse{}
	_body, _err := client.CreateBackupPlanWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateClusterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateCluster"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateGlobalResourceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateGlobalResource"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) CreateHbaseHaSlbWithOptions(request *CreateHbaseHaSlbRequest, runtime *util.RuntimeOptions) (_result *CreateHbaseHaSlbResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateHbaseHaSlbResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateHbaseHaSlb"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateHbaseHaSlb(request *CreateHbaseHaSlbRequest) (_result *CreateHbaseHaSlbResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateHbaseHaSlbResponse{}
	_body, _err := client.CreateHbaseHaSlbWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateHBaseSlbServerWithOptions(request *CreateHBaseSlbServerRequest, runtime *util.RuntimeOptions) (_result *CreateHBaseSlbServerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateHBaseSlbServerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateHBaseSlbServer"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateHBaseSlbServer(request *CreateHBaseSlbServerRequest) (_result *CreateHBaseSlbServerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateHBaseSlbServerResponse{}
	_body, _err := client.CreateHBaseSlbServerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMultiZoneClusterWithOptions(request *CreateMultiZoneClusterRequest, runtime *util.RuntimeOptions) (_result *CreateMultiZoneClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateMultiZoneClusterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateMultiZoneCluster"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMultiZoneCluster(request *CreateMultiZoneClusterRequest) (_result *CreateMultiZoneClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMultiZoneClusterResponse{}
	_body, _err := client.CreateMultiZoneClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRestorePlanWithOptions(request *CreateRestorePlanRequest, runtime *util.RuntimeOptions) (_result *CreateRestorePlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateRestorePlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateRestorePlan"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRestorePlan(request *CreateRestorePlanRequest) (_result *CreateRestorePlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRestorePlanResponse{}
	_body, _err := client.CreateRestorePlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateServerlessClusterWithOptions(request *CreateServerlessClusterRequest, runtime *util.RuntimeOptions) (_result *CreateServerlessClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateServerlessClusterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateServerlessCluster"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateServerlessCluster(request *CreateServerlessClusterRequest) (_result *CreateServerlessClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateServerlessClusterResponse{}
	_body, _err := client.CreateServerlessClusterWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteGlobalResourceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteGlobalResource"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DeleteHBaseHaDBWithOptions(request *DeleteHBaseHaDBRequest, runtime *util.RuntimeOptions) (_result *DeleteHBaseHaDBResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteHBaseHaDBResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteHBaseHaDB"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteHBaseHaDB(request *DeleteHBaseHaDBRequest) (_result *DeleteHBaseHaDBResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteHBaseHaDBResponse{}
	_body, _err := client.DeleteHBaseHaDBWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteHbaseHaSlbWithOptions(request *DeleteHbaseHaSlbRequest, runtime *util.RuntimeOptions) (_result *DeleteHbaseHaSlbResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteHbaseHaSlbResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteHbaseHaSlb"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteHbaseHaSlb(request *DeleteHbaseHaSlbRequest) (_result *DeleteHbaseHaSlbResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteHbaseHaSlbResponse{}
	_body, _err := client.DeleteHbaseHaSlbWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteHBaseSlbServerWithOptions(request *DeleteHBaseSlbServerRequest, runtime *util.RuntimeOptions) (_result *DeleteHBaseSlbServerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteHBaseSlbServerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteHBaseSlbServer"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteHBaseSlbServer(request *DeleteHBaseSlbServerRequest) (_result *DeleteHBaseSlbServerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteHBaseSlbServerResponse{}
	_body, _err := client.DeleteHBaseSlbServerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteInstanceWithOptions(request *DeleteInstanceRequest, runtime *util.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteInstance"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteInstance(request *DeleteInstanceRequest) (_result *DeleteInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DeleteInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteMultiZoneClusterWithOptions(request *DeleteMultiZoneClusterRequest, runtime *util.RuntimeOptions) (_result *DeleteMultiZoneClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteMultiZoneClusterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteMultiZoneCluster"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMultiZoneCluster(request *DeleteMultiZoneClusterRequest) (_result *DeleteMultiZoneClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMultiZoneClusterResponse{}
	_body, _err := client.DeleteMultiZoneClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteServerlessClusterWithOptions(request *DeleteServerlessClusterRequest, runtime *util.RuntimeOptions) (_result *DeleteServerlessClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteServerlessClusterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteServerlessCluster"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteServerlessCluster(request *DeleteServerlessClusterRequest) (_result *DeleteServerlessClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteServerlessClusterResponse{}
	_body, _err := client.DeleteServerlessClusterWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteUserHdfsInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteUserHdfsInfo"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeAvailableResourceWithOptions(request *DescribeAvailableResourceRequest, runtime *util.RuntimeOptions) (_result *DescribeAvailableResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAvailableResourceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAvailableResource"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAvailableResource(request *DescribeAvailableResourceRequest) (_result *DescribeAvailableResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAvailableResourceResponse{}
	_body, _err := client.DescribeAvailableResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBackupPlanConfigWithOptions(request *DescribeBackupPlanConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupPlanConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBackupPlanConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBackupPlanConfig"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackupPlanConfig(request *DescribeBackupPlanConfigRequest) (_result *DescribeBackupPlanConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupPlanConfigResponse{}
	_body, _err := client.DescribeBackupPlanConfigWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBackupPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBackupPolicy"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBackupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBackups"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeBackupStatusWithOptions(request *DescribeBackupStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBackupStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBackupStatus"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackupStatus(request *DescribeBackupStatusRequest) (_result *DescribeBackupStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupStatusResponse{}
	_body, _err := client.DescribeBackupStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBackupSummaryWithOptions(request *DescribeBackupSummaryRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBackupSummaryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBackupSummary"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackupSummary(request *DescribeBackupSummaryRequest) (_result *DescribeBackupSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupSummaryResponse{}
	_body, _err := client.DescribeBackupSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBackupTablesWithOptions(request *DescribeBackupTablesRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupTablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBackupTablesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBackupTables"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackupTables(request *DescribeBackupTablesRequest) (_result *DescribeBackupTablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupTablesResponse{}
	_body, _err := client.DescribeBackupTablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterConnectionWithOptions(request *DescribeClusterConnectionRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeClusterConnectionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeClusterConnection"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterConnection(request *DescribeClusterConnectionRequest) (_result *DescribeClusterConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterConnectionResponse{}
	_body, _err := client.DescribeClusterConnectionWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeColdStorageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeColdStorage"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeDBInstanceUsageWithOptions(request *DescribeDBInstanceUsageRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceUsageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBInstanceUsageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBInstanceUsage"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBInstanceUsage(request *DescribeDBInstanceUsageRequest) (_result *DescribeDBInstanceUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstanceUsageResponse{}
	_body, _err := client.DescribeDBInstanceUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDeletedInstancesWithOptions(request *DescribeDeletedInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeDeletedInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDeletedInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDeletedInstances"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDeletedInstances(request *DescribeDeletedInstancesRequest) (_result *DescribeDeletedInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDeletedInstancesResponse{}
	_body, _err := client.DescribeDeletedInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDiskWarningLineWithOptions(request *DescribeDiskWarningLineRequest, runtime *util.RuntimeOptions) (_result *DescribeDiskWarningLineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDiskWarningLineResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDiskWarningLine"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDiskWarningLine(request *DescribeDiskWarningLineRequest) (_result *DescribeDiskWarningLineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDiskWarningLineResponse{}
	_body, _err := client.DescribeDiskWarningLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEndpointsWithOptions(request *DescribeEndpointsRequest, runtime *util.RuntimeOptions) (_result *DescribeEndpointsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeEndpointsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeEndpoints"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEndpoints(request *DescribeEndpointsRequest) (_result *DescribeEndpointsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEndpointsResponse{}
	_body, _err := client.DescribeEndpointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceWithOptions(request *DescribeInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstance"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstance(request *DescribeInstanceRequest) (_result *DescribeInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceResponse{}
	_body, _err := client.DescribeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstancesWithOptions(request *DescribeInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstances"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstances(request *DescribeInstancesRequest) (_result *DescribeInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.DescribeInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceTypeWithOptions(request *DescribeInstanceTypeRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstanceTypeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstanceType"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceType(request *DescribeInstanceTypeRequest) (_result *DescribeInstanceTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceTypeResponse{}
	_body, _err := client.DescribeInstanceTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeIpWhitelistWithOptions(request *DescribeIpWhitelistRequest, runtime *util.RuntimeOptions) (_result *DescribeIpWhitelistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeIpWhitelistResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeIpWhitelist"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeIpWhitelist(request *DescribeIpWhitelistRequest) (_result *DescribeIpWhitelistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeIpWhitelistResponse{}
	_body, _err := client.DescribeIpWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMultiZoneAvailableRegionsWithOptions(request *DescribeMultiZoneAvailableRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeMultiZoneAvailableRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMultiZoneAvailableRegionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMultiZoneAvailableRegions"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMultiZoneAvailableRegions(request *DescribeMultiZoneAvailableRegionsRequest) (_result *DescribeMultiZoneAvailableRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMultiZoneAvailableRegionsResponse{}
	_body, _err := client.DescribeMultiZoneAvailableRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMultiZoneAvailableResourceWithOptions(request *DescribeMultiZoneAvailableResourceRequest, runtime *util.RuntimeOptions) (_result *DescribeMultiZoneAvailableResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMultiZoneAvailableResourceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMultiZoneAvailableResource"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMultiZoneAvailableResource(request *DescribeMultiZoneAvailableResourceRequest) (_result *DescribeMultiZoneAvailableResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMultiZoneAvailableResourceResponse{}
	_body, _err := client.DescribeMultiZoneAvailableResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMultiZoneClusterWithOptions(request *DescribeMultiZoneClusterRequest, runtime *util.RuntimeOptions) (_result *DescribeMultiZoneClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMultiZoneClusterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMultiZoneCluster"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMultiZoneCluster(request *DescribeMultiZoneClusterRequest) (_result *DescribeMultiZoneClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMultiZoneClusterResponse{}
	_body, _err := client.DescribeMultiZoneClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRecoverableTimeRangeWithOptions(request *DescribeRecoverableTimeRangeRequest, runtime *util.RuntimeOptions) (_result *DescribeRecoverableTimeRangeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRecoverableTimeRangeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRecoverableTimeRange"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRecoverableTimeRange(request *DescribeRecoverableTimeRangeRequest) (_result *DescribeRecoverableTimeRangeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRecoverableTimeRangeResponse{}
	_body, _err := client.DescribeRecoverableTimeRangeWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DescribeRegions"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeRestoreFullDetailsWithOptions(request *DescribeRestoreFullDetailsRequest, runtime *util.RuntimeOptions) (_result *DescribeRestoreFullDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRestoreFullDetailsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRestoreFullDetails"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRestoreFullDetails(request *DescribeRestoreFullDetailsRequest) (_result *DescribeRestoreFullDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRestoreFullDetailsResponse{}
	_body, _err := client.DescribeRestoreFullDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRestoreIncrDetailWithOptions(request *DescribeRestoreIncrDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeRestoreIncrDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRestoreIncrDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRestoreIncrDetail"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRestoreIncrDetail(request *DescribeRestoreIncrDetailRequest) (_result *DescribeRestoreIncrDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRestoreIncrDetailResponse{}
	_body, _err := client.DescribeRestoreIncrDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRestoreSchemaDetailsWithOptions(request *DescribeRestoreSchemaDetailsRequest, runtime *util.RuntimeOptions) (_result *DescribeRestoreSchemaDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRestoreSchemaDetailsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRestoreSchemaDetails"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRestoreSchemaDetails(request *DescribeRestoreSchemaDetailsRequest) (_result *DescribeRestoreSchemaDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRestoreSchemaDetailsResponse{}
	_body, _err := client.DescribeRestoreSchemaDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRestoreSummaryWithOptions(request *DescribeRestoreSummaryRequest, runtime *util.RuntimeOptions) (_result *DescribeRestoreSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRestoreSummaryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRestoreSummary"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRestoreSummary(request *DescribeRestoreSummaryRequest) (_result *DescribeRestoreSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRestoreSummaryResponse{}
	_body, _err := client.DescribeRestoreSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRestoreTablesWithOptions(request *DescribeRestoreTablesRequest, runtime *util.RuntimeOptions) (_result *DescribeRestoreTablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRestoreTablesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRestoreTables"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRestoreTables(request *DescribeRestoreTablesRequest) (_result *DescribeRestoreTablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRestoreTablesResponse{}
	_body, _err := client.DescribeRestoreTablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSecurityGroupsWithOptions(request *DescribeSecurityGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeSecurityGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSecurityGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSecurityGroups"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSecurityGroups(request *DescribeSecurityGroupsRequest) (_result *DescribeSecurityGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSecurityGroupsResponse{}
	_body, _err := client.DescribeSecurityGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSubDomainWithOptions(request *DescribeSubDomainRequest, runtime *util.RuntimeOptions) (_result *DescribeSubDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSubDomainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSubDomain"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSubDomain(request *DescribeSubDomainRequest) (_result *DescribeSubDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSubDomainResponse{}
	_body, _err := client.DescribeSubDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableHBaseueBackupWithOptions(request *EnableHBaseueBackupRequest, runtime *util.RuntimeOptions) (_result *EnableHBaseueBackupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnableHBaseueBackupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableHBaseueBackup"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableHBaseueBackup(request *EnableHBaseueBackupRequest) (_result *EnableHBaseueBackupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableHBaseueBackupResponse{}
	_body, _err := client.EnableHBaseueBackupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableHBaseueModuleWithOptions(request *EnableHBaseueModuleRequest, runtime *util.RuntimeOptions) (_result *EnableHBaseueModuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnableHBaseueModuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableHBaseueModule"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableHBaseueModule(request *EnableHBaseueModuleRequest) (_result *EnableHBaseueModuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableHBaseueModuleResponse{}
	_body, _err := client.EnableHBaseueModuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EvaluateMultiZoneResourceWithOptions(request *EvaluateMultiZoneResourceRequest, runtime *util.RuntimeOptions) (_result *EvaluateMultiZoneResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EvaluateMultiZoneResourceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EvaluateMultiZoneResource"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EvaluateMultiZoneResource(request *EvaluateMultiZoneResourceRequest) (_result *EvaluateMultiZoneResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EvaluateMultiZoneResourceResponse{}
	_body, _err := client.EvaluateMultiZoneResourceWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetMultimodeCmsUrlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetMultimodeCmsUrl"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListHBaseInstancesWithOptions(request *ListHBaseInstancesRequest, runtime *util.RuntimeOptions) (_result *ListHBaseInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListHBaseInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListHBaseInstances"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListHBaseInstances(request *ListHBaseInstancesRequest) (_result *ListHBaseInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListHBaseInstancesResponse{}
	_body, _err := client.ListHBaseInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInstanceServiceConfigHistoriesWithOptions(request *ListInstanceServiceConfigHistoriesRequest, runtime *util.RuntimeOptions) (_result *ListInstanceServiceConfigHistoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListInstanceServiceConfigHistoriesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListInstanceServiceConfigHistories"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInstanceServiceConfigHistories(request *ListInstanceServiceConfigHistoriesRequest) (_result *ListInstanceServiceConfigHistoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInstanceServiceConfigHistoriesResponse{}
	_body, _err := client.ListInstanceServiceConfigHistoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInstanceServiceConfigurationsWithOptions(request *ListInstanceServiceConfigurationsRequest, runtime *util.RuntimeOptions) (_result *ListInstanceServiceConfigurationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListInstanceServiceConfigurationsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListInstanceServiceConfigurations"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInstanceServiceConfigurations(request *ListInstanceServiceConfigurationsRequest) (_result *ListInstanceServiceConfigurationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInstanceServiceConfigurationsResponse{}
	_body, _err := client.ListInstanceServiceConfigurationsWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTagResources"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListTagsWithOptions(request *ListTagsRequest, runtime *util.RuntimeOptions) (_result *ListTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTags"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTags(request *ListTagsRequest) (_result *ListTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagsResponse{}
	_body, _err := client.ListTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyBackupPlanConfigWithOptions(request *ModifyBackupPlanConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyBackupPlanConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyBackupPlanConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyBackupPlanConfig"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyBackupPlanConfig(request *ModifyBackupPlanConfigRequest) (_result *ModifyBackupPlanConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyBackupPlanConfigResponse{}
	_body, _err := client.ModifyBackupPlanConfigWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyBackupPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyBackupPolicy"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ModifyClusterDeletionProtectionWithOptions(request *ModifyClusterDeletionProtectionRequest, runtime *util.RuntimeOptions) (_result *ModifyClusterDeletionProtectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyClusterDeletionProtectionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyClusterDeletionProtection"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyClusterDeletionProtection(request *ModifyClusterDeletionProtectionRequest) (_result *ModifyClusterDeletionProtectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyClusterDeletionProtectionResponse{}
	_body, _err := client.ModifyClusterDeletionProtectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDiskWarningLineWithOptions(request *ModifyDiskWarningLineRequest, runtime *util.RuntimeOptions) (_result *ModifyDiskWarningLineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDiskWarningLineResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDiskWarningLine"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDiskWarningLine(request *ModifyDiskWarningLineRequest) (_result *ModifyDiskWarningLineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDiskWarningLineResponse{}
	_body, _err := client.ModifyDiskWarningLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyInstanceMaintainTimeWithOptions(request *ModifyInstanceMaintainTimeRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceMaintainTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyInstanceMaintainTimeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyInstanceMaintainTime"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyInstanceMaintainTime(request *ModifyInstanceMaintainTimeRequest) (_result *ModifyInstanceMaintainTimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstanceMaintainTimeResponse{}
	_body, _err := client.ModifyInstanceMaintainTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyInstanceNameWithOptions(request *ModifyInstanceNameRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyInstanceNameResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyInstanceName"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyInstanceName(request *ModifyInstanceNameRequest) (_result *ModifyInstanceNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstanceNameResponse{}
	_body, _err := client.ModifyInstanceNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyInstanceServiceConfigWithOptions(request *ModifyInstanceServiceConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceServiceConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyInstanceServiceConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyInstanceServiceConfig"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyInstanceServiceConfig(request *ModifyInstanceServiceConfigRequest) (_result *ModifyInstanceServiceConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstanceServiceConfigResponse{}
	_body, _err := client.ModifyInstanceServiceConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyInstanceTypeWithOptions(request *ModifyInstanceTypeRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyInstanceTypeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyInstanceType"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyInstanceType(request *ModifyInstanceTypeRequest) (_result *ModifyInstanceTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstanceTypeResponse{}
	_body, _err := client.ModifyInstanceTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyIpWhitelistWithOptions(request *ModifyIpWhitelistRequest, runtime *util.RuntimeOptions) (_result *ModifyIpWhitelistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyIpWhitelistResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyIpWhitelist"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyIpWhitelist(request *ModifyIpWhitelistRequest) (_result *ModifyIpWhitelistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyIpWhitelistResponse{}
	_body, _err := client.ModifyIpWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyMultiZoneClusterNodeTypeWithOptions(request *ModifyMultiZoneClusterNodeTypeRequest, runtime *util.RuntimeOptions) (_result *ModifyMultiZoneClusterNodeTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyMultiZoneClusterNodeTypeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyMultiZoneClusterNodeType"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyMultiZoneClusterNodeType(request *ModifyMultiZoneClusterNodeTypeRequest) (_result *ModifyMultiZoneClusterNodeTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyMultiZoneClusterNodeTypeResponse{}
	_body, _err := client.ModifyMultiZoneClusterNodeTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifySecurityGroupsWithOptions(request *ModifySecurityGroupsRequest, runtime *util.RuntimeOptions) (_result *ModifySecurityGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifySecurityGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifySecurityGroups"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifySecurityGroups(request *ModifySecurityGroupsRequest) (_result *ModifySecurityGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifySecurityGroupsResponse{}
	_body, _err := client.ModifySecurityGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyUIAccountPasswordWithOptions(request *ModifyUIAccountPasswordRequest, runtime *util.RuntimeOptions) (_result *ModifyUIAccountPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyUIAccountPasswordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyUIAccountPassword"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyUIAccountPassword(request *ModifyUIAccountPasswordRequest) (_result *ModifyUIAccountPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyUIAccountPasswordResponse{}
	_body, _err := client.ModifyUIAccountPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MoveResourceGroupWithOptions(request *MoveResourceGroupRequest, runtime *util.RuntimeOptions) (_result *MoveResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("MoveResourceGroup"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MoveResourceGroup(request *MoveResourceGroupRequest) (_result *MoveResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.MoveResourceGroupWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OpenBackupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OpenBackup"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) PurgeInstanceWithOptions(request *PurgeInstanceRequest, runtime *util.RuntimeOptions) (_result *PurgeInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PurgeInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PurgeInstance"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PurgeInstance(request *PurgeInstanceRequest) (_result *PurgeInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PurgeInstanceResponse{}
	_body, _err := client.PurgeInstanceWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryHBaseHaDBResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryHBaseHaDB"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) QueryXpackRelateDBWithOptions(request *QueryXpackRelateDBRequest, runtime *util.RuntimeOptions) (_result *QueryXpackRelateDBResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryXpackRelateDBResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryXpackRelateDB"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryXpackRelateDB(request *QueryXpackRelateDBRequest) (_result *QueryXpackRelateDBResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryXpackRelateDBResponse{}
	_body, _err := client.QueryXpackRelateDBWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RelateDbForHBaseHaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RelateDbForHBaseHa"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ReleasePublicNetworkAddressResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ReleasePublicNetworkAddress"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RenewInstanceWithOptions(request *RenewInstanceRequest, runtime *util.RuntimeOptions) (_result *RenewInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RenewInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RenewInstance"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenewInstance(request *RenewInstanceRequest) (_result *RenewInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenewInstanceResponse{}
	_body, _err := client.RenewInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResizeColdStorageSizeWithOptions(request *ResizeColdStorageSizeRequest, runtime *util.RuntimeOptions) (_result *ResizeColdStorageSizeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ResizeColdStorageSizeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ResizeColdStorageSize"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResizeColdStorageSize(request *ResizeColdStorageSizeRequest) (_result *ResizeColdStorageSizeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResizeColdStorageSizeResponse{}
	_body, _err := client.ResizeColdStorageSizeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResizeDiskSizeWithOptions(request *ResizeDiskSizeRequest, runtime *util.RuntimeOptions) (_result *ResizeDiskSizeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ResizeDiskSizeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ResizeDiskSize"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResizeDiskSize(request *ResizeDiskSizeRequest) (_result *ResizeDiskSizeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResizeDiskSizeResponse{}
	_body, _err := client.ResizeDiskSizeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResizeMultiZoneClusterDiskSizeWithOptions(request *ResizeMultiZoneClusterDiskSizeRequest, runtime *util.RuntimeOptions) (_result *ResizeMultiZoneClusterDiskSizeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ResizeMultiZoneClusterDiskSizeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ResizeMultiZoneClusterDiskSize"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResizeMultiZoneClusterDiskSize(request *ResizeMultiZoneClusterDiskSizeRequest) (_result *ResizeMultiZoneClusterDiskSizeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResizeMultiZoneClusterDiskSizeResponse{}
	_body, _err := client.ResizeMultiZoneClusterDiskSizeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResizeMultiZoneClusterNodeCountWithOptions(request *ResizeMultiZoneClusterNodeCountRequest, runtime *util.RuntimeOptions) (_result *ResizeMultiZoneClusterNodeCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ResizeMultiZoneClusterNodeCountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ResizeMultiZoneClusterNodeCount"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResizeMultiZoneClusterNodeCount(request *ResizeMultiZoneClusterNodeCountRequest) (_result *ResizeMultiZoneClusterNodeCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResizeMultiZoneClusterNodeCountResponse{}
	_body, _err := client.ResizeMultiZoneClusterNodeCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResizeNodeCountWithOptions(request *ResizeNodeCountRequest, runtime *util.RuntimeOptions) (_result *ResizeNodeCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ResizeNodeCountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ResizeNodeCount"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResizeNodeCount(request *ResizeNodeCountRequest) (_result *ResizeNodeCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResizeNodeCountResponse{}
	_body, _err := client.ResizeNodeCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RestartInstanceWithOptions(request *RestartInstanceRequest, runtime *util.RuntimeOptions) (_result *RestartInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RestartInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RestartInstance"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RestartInstance(request *RestartInstanceRequest) (_result *RestartInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestartInstanceResponse{}
	_body, _err := client.RestartInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SwitchHbaseHaSlbWithOptions(request *SwitchHbaseHaSlbRequest, runtime *util.RuntimeOptions) (_result *SwitchHbaseHaSlbResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SwitchHbaseHaSlbResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SwitchHbaseHaSlb"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SwitchHbaseHaSlb(request *SwitchHbaseHaSlbRequest) (_result *SwitchHbaseHaSlbResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SwitchHbaseHaSlbResponse{}
	_body, _err := client.SwitchHbaseHaSlbWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TagResources"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) UnTagResourcesWithOptions(request *UnTagResourcesRequest, runtime *util.RuntimeOptions) (_result *UnTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UnTagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UnTagResources"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnTagResources(request *UnTagResourcesRequest) (_result *UnTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnTagResourcesResponse{}
	_body, _err := client.UnTagResourcesWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpgradeMinorVersionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpgradeMinorVersion"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) UpgradeMultiZoneClusterWithOptions(request *UpgradeMultiZoneClusterRequest, runtime *util.RuntimeOptions) (_result *UpgradeMultiZoneClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpgradeMultiZoneClusterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpgradeMultiZoneCluster"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeMultiZoneCluster(request *UpgradeMultiZoneClusterRequest) (_result *UpgradeMultiZoneClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeMultiZoneClusterResponse{}
	_body, _err := client.UpgradeMultiZoneClusterWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &XpackRelateDBResponse{}
	_body, _err := client.DoRPCRequest(tea.String("XpackRelateDB"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
