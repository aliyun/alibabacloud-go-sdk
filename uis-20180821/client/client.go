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

type AddHighPriorityIpRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UisId                *string `json:"UisId,omitempty" xml:"UisId,omitempty"`
	HighPriorityIp       *string `json:"HighPriorityIp,omitempty" xml:"HighPriorityIp,omitempty"`
}

func (s AddHighPriorityIpRequest) String() string {
	return tea.Prettify(s)
}

func (s AddHighPriorityIpRequest) GoString() string {
	return s.String()
}

func (s *AddHighPriorityIpRequest) SetOwnerAccount(v string) *AddHighPriorityIpRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddHighPriorityIpRequest) SetOwnerId(v int64) *AddHighPriorityIpRequest {
	s.OwnerId = &v
	return s
}

func (s *AddHighPriorityIpRequest) SetResourceOwnerAccount(v string) *AddHighPriorityIpRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddHighPriorityIpRequest) SetResourceOwnerId(v int64) *AddHighPriorityIpRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddHighPriorityIpRequest) SetClientToken(v string) *AddHighPriorityIpRequest {
	s.ClientToken = &v
	return s
}

func (s *AddHighPriorityIpRequest) SetRegionId(v string) *AddHighPriorityIpRequest {
	s.RegionId = &v
	return s
}

func (s *AddHighPriorityIpRequest) SetUisId(v string) *AddHighPriorityIpRequest {
	s.UisId = &v
	return s
}

func (s *AddHighPriorityIpRequest) SetHighPriorityIp(v string) *AddHighPriorityIpRequest {
	s.HighPriorityIp = &v
	return s
}

type AddHighPriorityIpResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddHighPriorityIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddHighPriorityIpResponseBody) GoString() string {
	return s.String()
}

func (s *AddHighPriorityIpResponseBody) SetRequestId(v string) *AddHighPriorityIpResponseBody {
	s.RequestId = &v
	return s
}

type AddHighPriorityIpResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddHighPriorityIpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddHighPriorityIpResponse) String() string {
	return tea.Prettify(s)
}

func (s AddHighPriorityIpResponse) GoString() string {
	return s.String()
}

func (s *AddHighPriorityIpResponse) SetHeaders(v map[string]*string) *AddHighPriorityIpResponse {
	s.Headers = v
	return s
}

func (s *AddHighPriorityIpResponse) SetBody(v *AddHighPriorityIpResponseBody) *AddHighPriorityIpResponse {
	s.Body = v
	return s
}

type AddUisNodeIpRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UisNodeId            *string `json:"UisNodeId,omitempty" xml:"UisNodeId,omitempty"`
	IpAddrsNum           *int32  `json:"IpAddrsNum,omitempty" xml:"IpAddrsNum,omitempty"`
}

func (s AddUisNodeIpRequest) String() string {
	return tea.Prettify(s)
}

func (s AddUisNodeIpRequest) GoString() string {
	return s.String()
}

func (s *AddUisNodeIpRequest) SetOwnerAccount(v string) *AddUisNodeIpRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddUisNodeIpRequest) SetOwnerId(v int64) *AddUisNodeIpRequest {
	s.OwnerId = &v
	return s
}

func (s *AddUisNodeIpRequest) SetResourceOwnerAccount(v string) *AddUisNodeIpRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddUisNodeIpRequest) SetResourceOwnerId(v int64) *AddUisNodeIpRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddUisNodeIpRequest) SetClientToken(v string) *AddUisNodeIpRequest {
	s.ClientToken = &v
	return s
}

func (s *AddUisNodeIpRequest) SetRegionId(v string) *AddUisNodeIpRequest {
	s.RegionId = &v
	return s
}

func (s *AddUisNodeIpRequest) SetUisNodeId(v string) *AddUisNodeIpRequest {
	s.UisNodeId = &v
	return s
}

func (s *AddUisNodeIpRequest) SetIpAddrsNum(v int32) *AddUisNodeIpRequest {
	s.IpAddrsNum = &v
	return s
}

type AddUisNodeIpResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddUisNodeIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddUisNodeIpResponseBody) GoString() string {
	return s.String()
}

func (s *AddUisNodeIpResponseBody) SetRequestId(v string) *AddUisNodeIpResponseBody {
	s.RequestId = &v
	return s
}

type AddUisNodeIpResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddUisNodeIpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddUisNodeIpResponse) String() string {
	return tea.Prettify(s)
}

func (s AddUisNodeIpResponse) GoString() string {
	return s.String()
}

func (s *AddUisNodeIpResponse) SetHeaders(v map[string]*string) *AddUisNodeIpResponse {
	s.Headers = v
	return s
}

func (s *AddUisNodeIpResponse) SetBody(v *AddUisNodeIpResponseBody) *AddUisNodeIpResponse {
	s.Body = v
	return s
}

type CreateDnatEntryRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	UisNodeId            *string `json:"UisNodeId,omitempty" xml:"UisNodeId,omitempty"`
	DestinationIp        *string `json:"DestinationIp,omitempty" xml:"DestinationIp,omitempty"`
	DestinationPort      *int32  `json:"DestinationPort,omitempty" xml:"DestinationPort,omitempty"`
	OriginalIp           *string `json:"OriginalIp,omitempty" xml:"OriginalIp,omitempty"`
	OriginalPort         *int32  `json:"OriginalPort,omitempty" xml:"OriginalPort,omitempty"`
	IpProtocol           *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateDnatEntryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDnatEntryRequest) GoString() string {
	return s.String()
}

func (s *CreateDnatEntryRequest) SetOwnerAccount(v string) *CreateDnatEntryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDnatEntryRequest) SetOwnerId(v int64) *CreateDnatEntryRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDnatEntryRequest) SetResourceOwnerAccount(v string) *CreateDnatEntryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDnatEntryRequest) SetResourceOwnerId(v int64) *CreateDnatEntryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDnatEntryRequest) SetUisNodeId(v string) *CreateDnatEntryRequest {
	s.UisNodeId = &v
	return s
}

func (s *CreateDnatEntryRequest) SetDestinationIp(v string) *CreateDnatEntryRequest {
	s.DestinationIp = &v
	return s
}

func (s *CreateDnatEntryRequest) SetDestinationPort(v int32) *CreateDnatEntryRequest {
	s.DestinationPort = &v
	return s
}

func (s *CreateDnatEntryRequest) SetOriginalIp(v string) *CreateDnatEntryRequest {
	s.OriginalIp = &v
	return s
}

func (s *CreateDnatEntryRequest) SetOriginalPort(v int32) *CreateDnatEntryRequest {
	s.OriginalPort = &v
	return s
}

func (s *CreateDnatEntryRequest) SetIpProtocol(v string) *CreateDnatEntryRequest {
	s.IpProtocol = &v
	return s
}

func (s *CreateDnatEntryRequest) SetName(v string) *CreateDnatEntryRequest {
	s.Name = &v
	return s
}

type CreateDnatEntryResponseBody struct {
	UisDnatId *string `json:"UisDnatId,omitempty" xml:"UisDnatId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDnatEntryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDnatEntryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDnatEntryResponseBody) SetUisDnatId(v string) *CreateDnatEntryResponseBody {
	s.UisDnatId = &v
	return s
}

func (s *CreateDnatEntryResponseBody) SetRequestId(v string) *CreateDnatEntryResponseBody {
	s.RequestId = &v
	return s
}

type CreateDnatEntryResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDnatEntryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDnatEntryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDnatEntryResponse) GoString() string {
	return s.String()
}

func (s *CreateDnatEntryResponse) SetHeaders(v map[string]*string) *CreateDnatEntryResponse {
	s.Headers = v
	return s
}

func (s *CreateDnatEntryResponse) SetBody(v *CreateDnatEntryResponseBody) *CreateDnatEntryResponse {
	s.Body = v
	return s
}

type CreateSubConnectionRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	UisNodeId            *string `json:"UisNodeId,omitempty" xml:"UisNodeId,omitempty"`
	Ip                   *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	GreConfig            *string `json:"GreConfig,omitempty" xml:"GreConfig,omitempty"`
}

func (s CreateSubConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSubConnectionRequest) GoString() string {
	return s.String()
}

func (s *CreateSubConnectionRequest) SetOwnerAccount(v string) *CreateSubConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateSubConnectionRequest) SetOwnerId(v int64) *CreateSubConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSubConnectionRequest) SetResourceOwnerAccount(v string) *CreateSubConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateSubConnectionRequest) SetResourceOwnerId(v int64) *CreateSubConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateSubConnectionRequest) SetUisNodeId(v string) *CreateSubConnectionRequest {
	s.UisNodeId = &v
	return s
}

func (s *CreateSubConnectionRequest) SetIp(v string) *CreateSubConnectionRequest {
	s.Ip = &v
	return s
}

func (s *CreateSubConnectionRequest) SetDescription(v string) *CreateSubConnectionRequest {
	s.Description = &v
	return s
}

func (s *CreateSubConnectionRequest) SetName(v string) *CreateSubConnectionRequest {
	s.Name = &v
	return s
}

func (s *CreateSubConnectionRequest) SetGreConfig(v string) *CreateSubConnectionRequest {
	s.GreConfig = &v
	return s
}

type CreateSubConnectionResponseBody struct {
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UisSubConnectionId *string `json:"UisSubConnectionId,omitempty" xml:"UisSubConnectionId,omitempty"`
}

func (s CreateSubConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSubConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSubConnectionResponseBody) SetRequestId(v string) *CreateSubConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSubConnectionResponseBody) SetUisSubConnectionId(v string) *CreateSubConnectionResponseBody {
	s.UisSubConnectionId = &v
	return s
}

type CreateSubConnectionResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSubConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSubConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSubConnectionResponse) GoString() string {
	return s.String()
}

func (s *CreateSubConnectionResponse) SetHeaders(v map[string]*string) *CreateSubConnectionResponse {
	s.Headers = v
	return s
}

func (s *CreateSubConnectionResponse) SetBody(v *CreateSubConnectionResponseBody) *CreateSubConnectionResponse {
	s.Body = v
	return s
}

type CreateUisRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceChargeType   *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	InternetChargeType   *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	AutoPay              *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	Duration             *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	PricingCycle         *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	Bandwidth            *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	BandwidthType        *string `json:"BandwidthType,omitempty" xml:"BandwidthType,omitempty"`
	UisProtocol          *string `json:"UisProtocol,omitempty" xml:"UisProtocol,omitempty"`
	ConnectionBandwidth  *int32  `json:"ConnectionBandwidth,omitempty" xml:"ConnectionBandwidth,omitempty"`
	ConnectionCount      *int32  `json:"ConnectionCount,omitempty" xml:"ConnectionCount,omitempty"`
	ServiceRegion        *string `json:"ServiceRegion,omitempty" xml:"ServiceRegion,omitempty"`
	AccessType           *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
}

func (s CreateUisRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUisRequest) GoString() string {
	return s.String()
}

func (s *CreateUisRequest) SetOwnerAccount(v string) *CreateUisRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateUisRequest) SetOwnerId(v int64) *CreateUisRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateUisRequest) SetResourceOwnerAccount(v string) *CreateUisRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateUisRequest) SetResourceOwnerId(v int64) *CreateUisRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateUisRequest) SetRegionId(v string) *CreateUisRequest {
	s.RegionId = &v
	return s
}

func (s *CreateUisRequest) SetClientToken(v string) *CreateUisRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateUisRequest) SetName(v string) *CreateUisRequest {
	s.Name = &v
	return s
}

func (s *CreateUisRequest) SetDescription(v string) *CreateUisRequest {
	s.Description = &v
	return s
}

func (s *CreateUisRequest) SetInstanceChargeType(v string) *CreateUisRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *CreateUisRequest) SetInternetChargeType(v string) *CreateUisRequest {
	s.InternetChargeType = &v
	return s
}

func (s *CreateUisRequest) SetAutoPay(v bool) *CreateUisRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateUisRequest) SetDuration(v int32) *CreateUisRequest {
	s.Duration = &v
	return s
}

func (s *CreateUisRequest) SetPricingCycle(v string) *CreateUisRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateUisRequest) SetBandwidth(v int32) *CreateUisRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateUisRequest) SetBandwidthType(v string) *CreateUisRequest {
	s.BandwidthType = &v
	return s
}

func (s *CreateUisRequest) SetUisProtocol(v string) *CreateUisRequest {
	s.UisProtocol = &v
	return s
}

func (s *CreateUisRequest) SetConnectionBandwidth(v int32) *CreateUisRequest {
	s.ConnectionBandwidth = &v
	return s
}

func (s *CreateUisRequest) SetConnectionCount(v int32) *CreateUisRequest {
	s.ConnectionCount = &v
	return s
}

func (s *CreateUisRequest) SetServiceRegion(v string) *CreateUisRequest {
	s.ServiceRegion = &v
	return s
}

func (s *CreateUisRequest) SetAccessType(v string) *CreateUisRequest {
	s.AccessType = &v
	return s
}

type CreateUisResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UisId     *string `json:"UisId,omitempty" xml:"UisId,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateUisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUisResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUisResponseBody) SetRequestId(v string) *CreateUisResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUisResponseBody) SetUisId(v string) *CreateUisResponseBody {
	s.UisId = &v
	return s
}

func (s *CreateUisResponseBody) SetOrderId(v string) *CreateUisResponseBody {
	s.OrderId = &v
	return s
}

type CreateUisResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateUisResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateUisResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUisResponse) GoString() string {
	return s.String()
}

func (s *CreateUisResponse) SetHeaders(v map[string]*string) *CreateUisResponse {
	s.Headers = v
	return s
}

func (s *CreateUisResponse) SetBody(v *CreateUisResponseBody) *CreateUisResponse {
	s.Body = v
	return s
}

type CreateUisConnectionRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UisNodeId            *string `json:"UisNodeId,omitempty" xml:"UisNodeId,omitempty"`
	UisProtocol          *string `json:"UisProtocol,omitempty" xml:"UisProtocol,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	GreConfig            *string `json:"GreConfig,omitempty" xml:"GreConfig,omitempty"`
	SslConfig            *string `json:"SslConfig,omitempty" xml:"SslConfig,omitempty"`
}

func (s CreateUisConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUisConnectionRequest) GoString() string {
	return s.String()
}

func (s *CreateUisConnectionRequest) SetOwnerAccount(v string) *CreateUisConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateUisConnectionRequest) SetOwnerId(v int64) *CreateUisConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateUisConnectionRequest) SetResourceOwnerAccount(v string) *CreateUisConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateUisConnectionRequest) SetResourceOwnerId(v int64) *CreateUisConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateUisConnectionRequest) SetRegionId(v string) *CreateUisConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *CreateUisConnectionRequest) SetUisNodeId(v string) *CreateUisConnectionRequest {
	s.UisNodeId = &v
	return s
}

func (s *CreateUisConnectionRequest) SetUisProtocol(v string) *CreateUisConnectionRequest {
	s.UisProtocol = &v
	return s
}

func (s *CreateUisConnectionRequest) SetDescription(v string) *CreateUisConnectionRequest {
	s.Description = &v
	return s
}

func (s *CreateUisConnectionRequest) SetName(v string) *CreateUisConnectionRequest {
	s.Name = &v
	return s
}

func (s *CreateUisConnectionRequest) SetGreConfig(v string) *CreateUisConnectionRequest {
	s.GreConfig = &v
	return s
}

func (s *CreateUisConnectionRequest) SetSslConfig(v string) *CreateUisConnectionRequest {
	s.SslConfig = &v
	return s
}

type CreateUisConnectionResponseBody struct {
	UisConnectionId *string `json:"UisConnectionId,omitempty" xml:"UisConnectionId,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateUisConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUisConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUisConnectionResponseBody) SetUisConnectionId(v string) *CreateUisConnectionResponseBody {
	s.UisConnectionId = &v
	return s
}

func (s *CreateUisConnectionResponseBody) SetRequestId(v string) *CreateUisConnectionResponseBody {
	s.RequestId = &v
	return s
}

type CreateUisConnectionResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateUisConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateUisConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUisConnectionResponse) GoString() string {
	return s.String()
}

func (s *CreateUisConnectionResponse) SetHeaders(v map[string]*string) *CreateUisConnectionResponse {
	s.Headers = v
	return s
}

func (s *CreateUisConnectionResponse) SetBody(v *CreateUisConnectionResponseBody) *CreateUisConnectionResponse {
	s.Body = v
	return s
}

type CreateUisNetworkInterfaceRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	UisNodeId            *string `json:"UisNodeId,omitempty" xml:"UisNodeId,omitempty"`
	VswitchId            *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	SecurityGroupId      *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	IpAddress            *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s CreateUisNetworkInterfaceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUisNetworkInterfaceRequest) GoString() string {
	return s.String()
}

func (s *CreateUisNetworkInterfaceRequest) SetOwnerAccount(v string) *CreateUisNetworkInterfaceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateUisNetworkInterfaceRequest) SetOwnerId(v int64) *CreateUisNetworkInterfaceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateUisNetworkInterfaceRequest) SetResourceOwnerAccount(v string) *CreateUisNetworkInterfaceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateUisNetworkInterfaceRequest) SetResourceOwnerId(v int64) *CreateUisNetworkInterfaceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateUisNetworkInterfaceRequest) SetUisNodeId(v string) *CreateUisNetworkInterfaceRequest {
	s.UisNodeId = &v
	return s
}

func (s *CreateUisNetworkInterfaceRequest) SetVswitchId(v string) *CreateUisNetworkInterfaceRequest {
	s.VswitchId = &v
	return s
}

func (s *CreateUisNetworkInterfaceRequest) SetSecurityGroupId(v string) *CreateUisNetworkInterfaceRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateUisNetworkInterfaceRequest) SetIpAddress(v string) *CreateUisNetworkInterfaceRequest {
	s.IpAddress = &v
	return s
}

func (s *CreateUisNetworkInterfaceRequest) SetName(v string) *CreateUisNetworkInterfaceRequest {
	s.Name = &v
	return s
}

func (s *CreateUisNetworkInterfaceRequest) SetDescription(v string) *CreateUisNetworkInterfaceRequest {
	s.Description = &v
	return s
}

type CreateUisNetworkInterfaceResponseBody struct {
	UisEniId  *string `json:"UisEniId,omitempty" xml:"UisEniId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateUisNetworkInterfaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUisNetworkInterfaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUisNetworkInterfaceResponseBody) SetUisEniId(v string) *CreateUisNetworkInterfaceResponseBody {
	s.UisEniId = &v
	return s
}

func (s *CreateUisNetworkInterfaceResponseBody) SetRequestId(v string) *CreateUisNetworkInterfaceResponseBody {
	s.RequestId = &v
	return s
}

type CreateUisNetworkInterfaceResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateUisNetworkInterfaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateUisNetworkInterfaceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUisNetworkInterfaceResponse) GoString() string {
	return s.String()
}

func (s *CreateUisNetworkInterfaceResponse) SetHeaders(v map[string]*string) *CreateUisNetworkInterfaceResponse {
	s.Headers = v
	return s
}

func (s *CreateUisNetworkInterfaceResponse) SetBody(v *CreateUisNetworkInterfaceResponseBody) *CreateUisNetworkInterfaceResponse {
	s.Body = v
	return s
}

type CreateUisNodeRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UisId                *string `json:"UisId,omitempty" xml:"UisId,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	UisNodeBandwidth     *int32  `json:"UisNodeBandwidth,omitempty" xml:"UisNodeBandwidth,omitempty"`
	IpAddrsNum           *int32  `json:"IpAddrsNum,omitempty" xml:"IpAddrsNum,omitempty"`
	UisNodeAreaId        *string `json:"UisNodeAreaId,omitempty" xml:"UisNodeAreaId,omitempty"`
}

func (s CreateUisNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUisNodeRequest) GoString() string {
	return s.String()
}

func (s *CreateUisNodeRequest) SetOwnerAccount(v string) *CreateUisNodeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateUisNodeRequest) SetOwnerId(v int64) *CreateUisNodeRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateUisNodeRequest) SetResourceOwnerAccount(v string) *CreateUisNodeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateUisNodeRequest) SetResourceOwnerId(v int64) *CreateUisNodeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateUisNodeRequest) SetRegionId(v string) *CreateUisNodeRequest {
	s.RegionId = &v
	return s
}

func (s *CreateUisNodeRequest) SetUisId(v string) *CreateUisNodeRequest {
	s.UisId = &v
	return s
}

func (s *CreateUisNodeRequest) SetName(v string) *CreateUisNodeRequest {
	s.Name = &v
	return s
}

func (s *CreateUisNodeRequest) SetDescription(v string) *CreateUisNodeRequest {
	s.Description = &v
	return s
}

func (s *CreateUisNodeRequest) SetUisNodeBandwidth(v int32) *CreateUisNodeRequest {
	s.UisNodeBandwidth = &v
	return s
}

func (s *CreateUisNodeRequest) SetIpAddrsNum(v int32) *CreateUisNodeRequest {
	s.IpAddrsNum = &v
	return s
}

func (s *CreateUisNodeRequest) SetUisNodeAreaId(v string) *CreateUisNodeRequest {
	s.UisNodeAreaId = &v
	return s
}

type CreateUisNodeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UisNodeId *string `json:"UisNodeId,omitempty" xml:"UisNodeId,omitempty"`
}

func (s CreateUisNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUisNodeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUisNodeResponseBody) SetRequestId(v string) *CreateUisNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUisNodeResponseBody) SetUisNodeId(v string) *CreateUisNodeResponseBody {
	s.UisNodeId = &v
	return s
}

type CreateUisNodeResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateUisNodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateUisNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUisNodeResponse) GoString() string {
	return s.String()
}

func (s *CreateUisNodeResponse) SetHeaders(v map[string]*string) *CreateUisNodeResponse {
	s.Headers = v
	return s
}

func (s *CreateUisNodeResponse) SetBody(v *CreateUisNodeResponseBody) *CreateUisNodeResponse {
	s.Body = v
	return s
}

type DeleteDnatEntryRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	UisNodeId            *string `json:"UisNodeId,omitempty" xml:"UisNodeId,omitempty"`
	UisDnatId            *string `json:"UisDnatId,omitempty" xml:"UisDnatId,omitempty"`
}

func (s DeleteDnatEntryRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDnatEntryRequest) GoString() string {
	return s.String()
}

func (s *DeleteDnatEntryRequest) SetOwnerAccount(v string) *DeleteDnatEntryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDnatEntryRequest) SetOwnerId(v int64) *DeleteDnatEntryRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDnatEntryRequest) SetResourceOwnerAccount(v string) *DeleteDnatEntryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDnatEntryRequest) SetResourceOwnerId(v int64) *DeleteDnatEntryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteDnatEntryRequest) SetUisNodeId(v string) *DeleteDnatEntryRequest {
	s.UisNodeId = &v
	return s
}

func (s *DeleteDnatEntryRequest) SetUisDnatId(v string) *DeleteDnatEntryRequest {
	s.UisDnatId = &v
	return s
}

type DeleteDnatEntryResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDnatEntryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDnatEntryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDnatEntryResponseBody) SetRequestId(v string) *DeleteDnatEntryResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDnatEntryResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDnatEntryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDnatEntryResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDnatEntryResponse) GoString() string {
	return s.String()
}

func (s *DeleteDnatEntryResponse) SetHeaders(v map[string]*string) *DeleteDnatEntryResponse {
	s.Headers = v
	return s
}

func (s *DeleteDnatEntryResponse) SetBody(v *DeleteDnatEntryResponseBody) *DeleteDnatEntryResponse {
	s.Body = v
	return s
}

type DeleteHighPriorityIpRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UisId                *string `json:"UisId,omitempty" xml:"UisId,omitempty"`
	HighPriorityIp       *string `json:"HighPriorityIp,omitempty" xml:"HighPriorityIp,omitempty"`
}

func (s DeleteHighPriorityIpRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteHighPriorityIpRequest) GoString() string {
	return s.String()
}

func (s *DeleteHighPriorityIpRequest) SetOwnerAccount(v string) *DeleteHighPriorityIpRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteHighPriorityIpRequest) SetOwnerId(v int64) *DeleteHighPriorityIpRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteHighPriorityIpRequest) SetResourceOwnerAccount(v string) *DeleteHighPriorityIpRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteHighPriorityIpRequest) SetResourceOwnerId(v int64) *DeleteHighPriorityIpRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteHighPriorityIpRequest) SetClientToken(v string) *DeleteHighPriorityIpRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteHighPriorityIpRequest) SetRegionId(v string) *DeleteHighPriorityIpRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteHighPriorityIpRequest) SetUisId(v string) *DeleteHighPriorityIpRequest {
	s.UisId = &v
	return s
}

func (s *DeleteHighPriorityIpRequest) SetHighPriorityIp(v string) *DeleteHighPriorityIpRequest {
	s.HighPriorityIp = &v
	return s
}

type DeleteHighPriorityIpResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteHighPriorityIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteHighPriorityIpResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHighPriorityIpResponseBody) SetRequestId(v string) *DeleteHighPriorityIpResponseBody {
	s.RequestId = &v
	return s
}

type DeleteHighPriorityIpResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteHighPriorityIpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteHighPriorityIpResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteHighPriorityIpResponse) GoString() string {
	return s.String()
}

func (s *DeleteHighPriorityIpResponse) SetHeaders(v map[string]*string) *DeleteHighPriorityIpResponse {
	s.Headers = v
	return s
}

func (s *DeleteHighPriorityIpResponse) SetBody(v *DeleteHighPriorityIpResponseBody) *DeleteHighPriorityIpResponse {
	s.Body = v
	return s
}

type DeleteSubConnectionRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	UisSubConnectionId   *string `json:"UisSubConnectionId,omitempty" xml:"UisSubConnectionId,omitempty"`
}

func (s DeleteSubConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSubConnectionRequest) GoString() string {
	return s.String()
}

func (s *DeleteSubConnectionRequest) SetOwnerAccount(v string) *DeleteSubConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteSubConnectionRequest) SetOwnerId(v int64) *DeleteSubConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSubConnectionRequest) SetResourceOwnerAccount(v string) *DeleteSubConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteSubConnectionRequest) SetResourceOwnerId(v int64) *DeleteSubConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteSubConnectionRequest) SetUisSubConnectionId(v string) *DeleteSubConnectionRequest {
	s.UisSubConnectionId = &v
	return s
}

type DeleteSubConnectionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSubConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSubConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSubConnectionResponseBody) SetRequestId(v string) *DeleteSubConnectionResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSubConnectionResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSubConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSubConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSubConnectionResponse) GoString() string {
	return s.String()
}

func (s *DeleteSubConnectionResponse) SetHeaders(v map[string]*string) *DeleteSubConnectionResponse {
	s.Headers = v
	return s
}

func (s *DeleteSubConnectionResponse) SetBody(v *DeleteSubConnectionResponseBody) *DeleteSubConnectionResponse {
	s.Body = v
	return s
}

type DeleteUisRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UisId                *string `json:"UisId,omitempty" xml:"UisId,omitempty"`
}

func (s DeleteUisRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUisRequest) GoString() string {
	return s.String()
}

func (s *DeleteUisRequest) SetOwnerAccount(v string) *DeleteUisRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteUisRequest) SetOwnerId(v int64) *DeleteUisRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteUisRequest) SetResourceOwnerAccount(v string) *DeleteUisRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteUisRequest) SetResourceOwnerId(v int64) *DeleteUisRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteUisRequest) SetClientToken(v string) *DeleteUisRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteUisRequest) SetRegionId(v string) *DeleteUisRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteUisRequest) SetUisId(v string) *DeleteUisRequest {
	s.UisId = &v
	return s
}

type DeleteUisResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUisResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUisResponseBody) SetRequestId(v string) *DeleteUisResponseBody {
	s.RequestId = &v
	return s
}

type DeleteUisResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteUisResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteUisResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUisResponse) GoString() string {
	return s.String()
}

func (s *DeleteUisResponse) SetHeaders(v map[string]*string) *DeleteUisResponse {
	s.Headers = v
	return s
}

func (s *DeleteUisResponse) SetBody(v *DeleteUisResponseBody) *DeleteUisResponse {
	s.Body = v
	return s
}

type DeleteUisConnectionRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	UisConnectionId      *string `json:"UisConnectionId,omitempty" xml:"UisConnectionId,omitempty"`
	UisNodeId            *string `json:"UisNodeId,omitempty" xml:"UisNodeId,omitempty"`
}

func (s DeleteUisConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUisConnectionRequest) GoString() string {
	return s.String()
}

func (s *DeleteUisConnectionRequest) SetOwnerAccount(v string) *DeleteUisConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteUisConnectionRequest) SetOwnerId(v int64) *DeleteUisConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteUisConnectionRequest) SetResourceOwnerAccount(v string) *DeleteUisConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteUisConnectionRequest) SetResourceOwnerId(v int64) *DeleteUisConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteUisConnectionRequest) SetClientToken(v string) *DeleteUisConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteUisConnectionRequest) SetUisConnectionId(v string) *DeleteUisConnectionRequest {
	s.UisConnectionId = &v
	return s
}

func (s *DeleteUisConnectionRequest) SetUisNodeId(v string) *DeleteUisConnectionRequest {
	s.UisNodeId = &v
	return s
}

type DeleteUisConnectionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUisConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUisConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUisConnectionResponseBody) SetRequestId(v string) *DeleteUisConnectionResponseBody {
	s.RequestId = &v
	return s
}

type DeleteUisConnectionResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteUisConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteUisConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUisConnectionResponse) GoString() string {
	return s.String()
}

func (s *DeleteUisConnectionResponse) SetHeaders(v map[string]*string) *DeleteUisConnectionResponse {
	s.Headers = v
	return s
}

func (s *DeleteUisConnectionResponse) SetBody(v *DeleteUisConnectionResponseBody) *DeleteUisConnectionResponse {
	s.Body = v
	return s
}

type DeleteUisNetworkInterfaceRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	UisNodeId            *string `json:"UisNodeId,omitempty" xml:"UisNodeId,omitempty"`
	UisEniId             *string `json:"UisEniId,omitempty" xml:"UisEniId,omitempty"`
}

func (s DeleteUisNetworkInterfaceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUisNetworkInterfaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteUisNetworkInterfaceRequest) SetOwnerAccount(v string) *DeleteUisNetworkInterfaceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteUisNetworkInterfaceRequest) SetOwnerId(v int64) *DeleteUisNetworkInterfaceRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteUisNetworkInterfaceRequest) SetResourceOwnerAccount(v string) *DeleteUisNetworkInterfaceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteUisNetworkInterfaceRequest) SetResourceOwnerId(v int64) *DeleteUisNetworkInterfaceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteUisNetworkInterfaceRequest) SetUisNodeId(v string) *DeleteUisNetworkInterfaceRequest {
	s.UisNodeId = &v
	return s
}

func (s *DeleteUisNetworkInterfaceRequest) SetUisEniId(v string) *DeleteUisNetworkInterfaceRequest {
	s.UisEniId = &v
	return s
}

type DeleteUisNetworkInterfaceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUisNetworkInterfaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUisNetworkInterfaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUisNetworkInterfaceResponseBody) SetRequestId(v string) *DeleteUisNetworkInterfaceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteUisNetworkInterfaceResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteUisNetworkInterfaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteUisNetworkInterfaceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUisNetworkInterfaceResponse) GoString() string {
	return s.String()
}

func (s *DeleteUisNetworkInterfaceResponse) SetHeaders(v map[string]*string) *DeleteUisNetworkInterfaceResponse {
	s.Headers = v
	return s
}

func (s *DeleteUisNetworkInterfaceResponse) SetBody(v *DeleteUisNetworkInterfaceResponseBody) *DeleteUisNetworkInterfaceResponse {
	s.Body = v
	return s
}

type DeleteUisNodeRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	UisId                *string `json:"UisId,omitempty" xml:"UisId,omitempty"`
	UisNodeId            *string `json:"UisNodeId,omitempty" xml:"UisNodeId,omitempty"`
}

func (s DeleteUisNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUisNodeRequest) GoString() string {
	return s.String()
}

func (s *DeleteUisNodeRequest) SetOwnerAccount(v string) *DeleteUisNodeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteUisNodeRequest) SetOwnerId(v int64) *DeleteUisNodeRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteUisNodeRequest) SetResourceOwnerAccount(v string) *DeleteUisNodeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteUisNodeRequest) SetResourceOwnerId(v int64) *DeleteUisNodeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteUisNodeRequest) SetClientToken(v string) *DeleteUisNodeRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteUisNodeRequest) SetUisId(v string) *DeleteUisNodeRequest {
	s.UisId = &v
	return s
}

func (s *DeleteUisNodeRequest) SetUisNodeId(v string) *DeleteUisNodeRequest {
	s.UisNodeId = &v
	return s
}

type DeleteUisNodeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUisNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUisNodeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUisNodeResponseBody) SetRequestId(v string) *DeleteUisNodeResponseBody {
	s.RequestId = &v
	return s
}

type DeleteUisNodeResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteUisNodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteUisNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUisNodeResponse) GoString() string {
	return s.String()
}

func (s *DeleteUisNodeResponse) SetHeaders(v map[string]*string) *DeleteUisNodeResponse {
	s.Headers = v
	return s
}

func (s *DeleteUisNodeResponse) SetBody(v *DeleteUisNodeResponseBody) *DeleteUisNodeResponse {
	s.Body = v
	return s
}

type DeleteUisNodeIpRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UisNodeId            *string `json:"UisNodeId,omitempty" xml:"UisNodeId,omitempty"`
	UisNodeIpAddress     *string `json:"UisNodeIpAddress,omitempty" xml:"UisNodeIpAddress,omitempty"`
}

func (s DeleteUisNodeIpRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUisNodeIpRequest) GoString() string {
	return s.String()
}

func (s *DeleteUisNodeIpRequest) SetOwnerAccount(v string) *DeleteUisNodeIpRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteUisNodeIpRequest) SetOwnerId(v int64) *DeleteUisNodeIpRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteUisNodeIpRequest) SetResourceOwnerAccount(v string) *DeleteUisNodeIpRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteUisNodeIpRequest) SetResourceOwnerId(v int64) *DeleteUisNodeIpRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteUisNodeIpRequest) SetClientToken(v string) *DeleteUisNodeIpRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteUisNodeIpRequest) SetRegionId(v string) *DeleteUisNodeIpRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteUisNodeIpRequest) SetUisNodeId(v string) *DeleteUisNodeIpRequest {
	s.UisNodeId = &v
	return s
}

func (s *DeleteUisNodeIpRequest) SetUisNodeIpAddress(v string) *DeleteUisNodeIpRequest {
	s.UisNodeIpAddress = &v
	return s
}

type DeleteUisNodeIpResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUisNodeIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUisNodeIpResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUisNodeIpResponseBody) SetRequestId(v string) *DeleteUisNodeIpResponseBody {
	s.RequestId = &v
	return s
}

type DeleteUisNodeIpResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteUisNodeIpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteUisNodeIpResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUisNodeIpResponse) GoString() string {
	return s.String()
}

func (s *DeleteUisNodeIpResponse) SetHeaders(v map[string]*string) *DeleteUisNodeIpResponse {
	s.Headers = v
	return s
}

func (s *DeleteUisNodeIpResponse) SetBody(v *DeleteUisNodeIpResponseBody) *DeleteUisNodeIpResponse {
	s.Body = v
	return s
}

type DescribeAreasRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAreasRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAreasRequest) GoString() string {
	return s.String()
}

func (s *DescribeAreasRequest) SetOwnerAccount(v string) *DescribeAreasRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAreasRequest) SetOwnerId(v int64) *DescribeAreasRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAreasRequest) SetResourceOwnerAccount(v string) *DescribeAreasRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAreasRequest) SetResourceOwnerId(v int64) *DescribeAreasRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAreasRequest) SetRegionId(v string) *DescribeAreasRequest {
	s.RegionId = &v
	return s
}

type DescribeAreasResponseBody struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Areas     *DescribeAreasResponseBodyAreas `json:"Areas,omitempty" xml:"Areas,omitempty" type:"Struct"`
}

func (s DescribeAreasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAreasResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAreasResponseBody) SetRequestId(v string) *DescribeAreasResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAreasResponseBody) SetAreas(v *DescribeAreasResponseBodyAreas) *DescribeAreasResponseBody {
	s.Areas = v
	return s
}

type DescribeAreasResponseBodyAreas struct {
	Area []*DescribeAreasResponseBodyAreasArea `json:"Area,omitempty" xml:"Area,omitempty" type:"Repeated"`
}

func (s DescribeAreasResponseBodyAreas) String() string {
	return tea.Prettify(s)
}

func (s DescribeAreasResponseBodyAreas) GoString() string {
	return s.String()
}

func (s *DescribeAreasResponseBodyAreas) SetArea(v []*DescribeAreasResponseBodyAreasArea) *DescribeAreasResponseBodyAreas {
	s.Area = v
	return s
}

type DescribeAreasResponseBodyAreasArea struct {
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	AreaId    *string `json:"AreaId,omitempty" xml:"AreaId,omitempty"`
}

func (s DescribeAreasResponseBodyAreasArea) String() string {
	return tea.Prettify(s)
}

func (s DescribeAreasResponseBodyAreasArea) GoString() string {
	return s.String()
}

func (s *DescribeAreasResponseBodyAreasArea) SetLocalName(v string) *DescribeAreasResponseBodyAreasArea {
	s.LocalName = &v
	return s
}

func (s *DescribeAreasResponseBodyAreasArea) SetAreaId(v string) *DescribeAreasResponseBodyAreasArea {
	s.AreaId = &v
	return s
}

type DescribeAreasResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAreasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAreasResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAreasResponse) GoString() string {
	return s.String()
}

func (s *DescribeAreasResponse) SetHeaders(v map[string]*string) *DescribeAreasResponse {
	s.Headers = v
	return s
}

func (s *DescribeAreasResponse) SetBody(v *DescribeAreasResponseBody) *DescribeAreasResponse {
	s.Body = v
	return s
}

type DescribeDnatEntriesRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UisNodeId            *string `json:"UisNodeId,omitempty" xml:"UisNodeId,omitempty"`
	UisDnatId            *string `json:"UisDnatId,omitempty" xml:"UisDnatId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDnatEntriesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnatEntriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDnatEntriesRequest) SetOwnerAccount(v string) *DescribeDnatEntriesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDnatEntriesRequest) SetOwnerId(v int64) *DescribeDnatEntriesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDnatEntriesRequest) SetResourceOwnerAccount(v string) *DescribeDnatEntriesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDnatEntriesRequest) SetResourceOwnerId(v int64) *DescribeDnatEntriesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDnatEntriesRequest) SetRegionId(v string) *DescribeDnatEntriesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDnatEntriesRequest) SetUisNodeId(v string) *DescribeDnatEntriesRequest {
	s.UisNodeId = &v
	return s
}

func (s *DescribeDnatEntriesRequest) SetUisDnatId(v string) *DescribeDnatEntriesRequest {
	s.UisDnatId = &v
	return s
}

func (s *DescribeDnatEntriesRequest) SetPageNumber(v int32) *DescribeDnatEntriesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDnatEntriesRequest) SetPageSize(v int32) *DescribeDnatEntriesRequest {
	s.PageSize = &v
	return s
}

type DescribeDnatEntriesResponseBody struct {
	TotalCount     *int32                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize       *int32                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber     *int32                                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	UisDnatEntries *DescribeDnatEntriesResponseBodyUisDnatEntries `json:"UisDnatEntries,omitempty" xml:"UisDnatEntries,omitempty" type:"Struct"`
}

func (s DescribeDnatEntriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnatEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDnatEntriesResponseBody) SetTotalCount(v int32) *DescribeDnatEntriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDnatEntriesResponseBody) SetPageSize(v int32) *DescribeDnatEntriesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDnatEntriesResponseBody) SetRequestId(v string) *DescribeDnatEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDnatEntriesResponseBody) SetPageNumber(v int32) *DescribeDnatEntriesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDnatEntriesResponseBody) SetUisDnatEntries(v *DescribeDnatEntriesResponseBodyUisDnatEntries) *DescribeDnatEntriesResponseBody {
	s.UisDnatEntries = v
	return s
}

type DescribeDnatEntriesResponseBodyUisDnatEntries struct {
	UisDnatEntry []*DescribeDnatEntriesResponseBodyUisDnatEntriesUisDnatEntry `json:"UisDnatEntry,omitempty" xml:"UisDnatEntry,omitempty" type:"Repeated"`
}

func (s DescribeDnatEntriesResponseBodyUisDnatEntries) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnatEntriesResponseBodyUisDnatEntries) GoString() string {
	return s.String()
}

func (s *DescribeDnatEntriesResponseBodyUisDnatEntries) SetUisDnatEntry(v []*DescribeDnatEntriesResponseBodyUisDnatEntriesUisDnatEntry) *DescribeDnatEntriesResponseBodyUisDnatEntries {
	s.UisDnatEntry = v
	return s
}

type DescribeDnatEntriesResponseBodyUisDnatEntriesUisDnatEntry struct {
	OriginalPort    *int32  `json:"OriginalPort,omitempty" xml:"OriginalPort,omitempty"`
	DestinationPort *int32  `json:"DestinationPort,omitempty" xml:"DestinationPort,omitempty"`
	OriginalIp      *string `json:"OriginalIp,omitempty" xml:"OriginalIp,omitempty"`
	IpProtocol      *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	UisDnatId       *string `json:"UisDnatId,omitempty" xml:"UisDnatId,omitempty"`
	DestinationIp   *string `json:"DestinationIp,omitempty" xml:"DestinationIp,omitempty"`
}

func (s DescribeDnatEntriesResponseBodyUisDnatEntriesUisDnatEntry) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnatEntriesResponseBodyUisDnatEntriesUisDnatEntry) GoString() string {
	return s.String()
}

func (s *DescribeDnatEntriesResponseBodyUisDnatEntriesUisDnatEntry) SetOriginalPort(v int32) *DescribeDnatEntriesResponseBodyUisDnatEntriesUisDnatEntry {
	s.OriginalPort = &v
	return s
}

func (s *DescribeDnatEntriesResponseBodyUisDnatEntriesUisDnatEntry) SetDestinationPort(v int32) *DescribeDnatEntriesResponseBodyUisDnatEntriesUisDnatEntry {
	s.DestinationPort = &v
	return s
}

func (s *DescribeDnatEntriesResponseBodyUisDnatEntriesUisDnatEntry) SetOriginalIp(v string) *DescribeDnatEntriesResponseBodyUisDnatEntriesUisDnatEntry {
	s.OriginalIp = &v
	return s
}

func (s *DescribeDnatEntriesResponseBodyUisDnatEntriesUisDnatEntry) SetIpProtocol(v string) *DescribeDnatEntriesResponseBodyUisDnatEntriesUisDnatEntry {
	s.IpProtocol = &v
	return s
}

func (s *DescribeDnatEntriesResponseBodyUisDnatEntriesUisDnatEntry) SetUisDnatId(v string) *DescribeDnatEntriesResponseBodyUisDnatEntriesUisDnatEntry {
	s.UisDnatId = &v
	return s
}

func (s *DescribeDnatEntriesResponseBodyUisDnatEntriesUisDnatEntry) SetDestinationIp(v string) *DescribeDnatEntriesResponseBodyUisDnatEntriesUisDnatEntry {
	s.DestinationIp = &v
	return s
}

type DescribeDnatEntriesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDnatEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDnatEntriesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDnatEntriesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDnatEntriesResponse) SetHeaders(v map[string]*string) *DescribeDnatEntriesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDnatEntriesResponse) SetBody(v *DescribeDnatEntriesResponseBody) *DescribeDnatEntriesResponse {
	s.Body = v
	return s
}

type DescribeHighPriorityIpRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UisId                *string `json:"UisId,omitempty" xml:"UisId,omitempty"`
	HighPriorityIp       *string `json:"HighPriorityIp,omitempty" xml:"HighPriorityIp,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s DescribeHighPriorityIpRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHighPriorityIpRequest) GoString() string {
	return s.String()
}

func (s *DescribeHighPriorityIpRequest) SetOwnerAccount(v string) *DescribeHighPriorityIpRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeHighPriorityIpRequest) SetOwnerId(v int64) *DescribeHighPriorityIpRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeHighPriorityIpRequest) SetResourceOwnerAccount(v string) *DescribeHighPriorityIpRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeHighPriorityIpRequest) SetResourceOwnerId(v int64) *DescribeHighPriorityIpRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeHighPriorityIpRequest) SetRegionId(v string) *DescribeHighPriorityIpRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHighPriorityIpRequest) SetUisId(v string) *DescribeHighPriorityIpRequest {
	s.UisId = &v
	return s
}

func (s *DescribeHighPriorityIpRequest) SetHighPriorityIp(v string) *DescribeHighPriorityIpRequest {
	s.HighPriorityIp = &v
	return s
}

func (s *DescribeHighPriorityIpRequest) SetPageNumber(v int32) *DescribeHighPriorityIpRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeHighPriorityIpRequest) SetPageSize(v int32) *DescribeHighPriorityIpRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeHighPriorityIpRequest) SetClientToken(v string) *DescribeHighPriorityIpRequest {
	s.ClientToken = &v
	return s
}

type DescribeHighPriorityIpResponseBody struct {
	TotalCount      *int32                                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize        *int32                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId       *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber      *int32                                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	HighPriorityIps *DescribeHighPriorityIpResponseBodyHighPriorityIps `json:"HighPriorityIps,omitempty" xml:"HighPriorityIps,omitempty" type:"Struct"`
}

func (s DescribeHighPriorityIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHighPriorityIpResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHighPriorityIpResponseBody) SetTotalCount(v int32) *DescribeHighPriorityIpResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeHighPriorityIpResponseBody) SetPageSize(v int32) *DescribeHighPriorityIpResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeHighPriorityIpResponseBody) SetRequestId(v string) *DescribeHighPriorityIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHighPriorityIpResponseBody) SetPageNumber(v int32) *DescribeHighPriorityIpResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeHighPriorityIpResponseBody) SetHighPriorityIps(v *DescribeHighPriorityIpResponseBodyHighPriorityIps) *DescribeHighPriorityIpResponseBody {
	s.HighPriorityIps = v
	return s
}

type DescribeHighPriorityIpResponseBodyHighPriorityIps struct {
	HighPriorityIp []*DescribeHighPriorityIpResponseBodyHighPriorityIpsHighPriorityIp `json:"HighPriorityIp,omitempty" xml:"HighPriorityIp,omitempty" type:"Repeated"`
}

func (s DescribeHighPriorityIpResponseBodyHighPriorityIps) String() string {
	return tea.Prettify(s)
}

func (s DescribeHighPriorityIpResponseBodyHighPriorityIps) GoString() string {
	return s.String()
}

func (s *DescribeHighPriorityIpResponseBodyHighPriorityIps) SetHighPriorityIp(v []*DescribeHighPriorityIpResponseBodyHighPriorityIpsHighPriorityIp) *DescribeHighPriorityIpResponseBodyHighPriorityIps {
	s.HighPriorityIp = v
	return s
}

type DescribeHighPriorityIpResponseBodyHighPriorityIpsHighPriorityIp struct {
	StaticOffAreaId  *string `json:"StaticOffAreaId,omitempty" xml:"StaticOffAreaId,omitempty"`
	Domain           *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	DynamicOffAreaId *string `json:"DynamicOffAreaId,omitempty" xml:"DynamicOffAreaId,omitempty"`
	State            *string `json:"State,omitempty" xml:"State,omitempty"`
	Ip               *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	BoardAreaId      *string `json:"BoardAreaId,omitempty" xml:"BoardAreaId,omitempty"`
}

func (s DescribeHighPriorityIpResponseBodyHighPriorityIpsHighPriorityIp) String() string {
	return tea.Prettify(s)
}

func (s DescribeHighPriorityIpResponseBodyHighPriorityIpsHighPriorityIp) GoString() string {
	return s.String()
}

func (s *DescribeHighPriorityIpResponseBodyHighPriorityIpsHighPriorityIp) SetStaticOffAreaId(v string) *DescribeHighPriorityIpResponseBodyHighPriorityIpsHighPriorityIp {
	s.StaticOffAreaId = &v
	return s
}

func (s *DescribeHighPriorityIpResponseBodyHighPriorityIpsHighPriorityIp) SetDomain(v string) *DescribeHighPriorityIpResponseBodyHighPriorityIpsHighPriorityIp {
	s.Domain = &v
	return s
}

func (s *DescribeHighPriorityIpResponseBodyHighPriorityIpsHighPriorityIp) SetDynamicOffAreaId(v string) *DescribeHighPriorityIpResponseBodyHighPriorityIpsHighPriorityIp {
	s.DynamicOffAreaId = &v
	return s
}

func (s *DescribeHighPriorityIpResponseBodyHighPriorityIpsHighPriorityIp) SetState(v string) *DescribeHighPriorityIpResponseBodyHighPriorityIpsHighPriorityIp {
	s.State = &v
	return s
}

func (s *DescribeHighPriorityIpResponseBodyHighPriorityIpsHighPriorityIp) SetIp(v string) *DescribeHighPriorityIpResponseBodyHighPriorityIpsHighPriorityIp {
	s.Ip = &v
	return s
}

func (s *DescribeHighPriorityIpResponseBodyHighPriorityIpsHighPriorityIp) SetBoardAreaId(v string) *DescribeHighPriorityIpResponseBodyHighPriorityIpsHighPriorityIp {
	s.BoardAreaId = &v
	return s
}

type DescribeHighPriorityIpResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeHighPriorityIpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeHighPriorityIpResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHighPriorityIpResponse) GoString() string {
	return s.String()
}

func (s *DescribeHighPriorityIpResponse) SetHeaders(v map[string]*string) *DescribeHighPriorityIpResponse {
	s.Headers = v
	return s
}

func (s *DescribeHighPriorityIpResponse) SetBody(v *DescribeHighPriorityIpResponseBody) *DescribeHighPriorityIpResponse {
	s.Body = v
	return s
}

type DescribeHighPriorityIpsRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UisId                *string `json:"UisId,omitempty" xml:"UisId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeHighPriorityIpsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHighPriorityIpsRequest) GoString() string {
	return s.String()
}

func (s *DescribeHighPriorityIpsRequest) SetOwnerAccount(v string) *DescribeHighPriorityIpsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeHighPriorityIpsRequest) SetOwnerId(v int64) *DescribeHighPriorityIpsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeHighPriorityIpsRequest) SetResourceOwnerAccount(v string) *DescribeHighPriorityIpsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeHighPriorityIpsRequest) SetResourceOwnerId(v int64) *DescribeHighPriorityIpsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeHighPriorityIpsRequest) SetRegionId(v string) *DescribeHighPriorityIpsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHighPriorityIpsRequest) SetUisId(v string) *DescribeHighPriorityIpsRequest {
	s.UisId = &v
	return s
}

func (s *DescribeHighPriorityIpsRequest) SetPageNumber(v int32) *DescribeHighPriorityIpsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeHighPriorityIpsRequest) SetPageSize(v int32) *DescribeHighPriorityIpsRequest {
	s.PageSize = &v
	return s
}

type DescribeHighPriorityIpsResponseBody struct {
	TotalCount      *int32                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize        *int32                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId       *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber      *int32                                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	HighPriorityIps *DescribeHighPriorityIpsResponseBodyHighPriorityIps `json:"HighPriorityIps,omitempty" xml:"HighPriorityIps,omitempty" type:"Struct"`
}

func (s DescribeHighPriorityIpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHighPriorityIpsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHighPriorityIpsResponseBody) SetTotalCount(v int32) *DescribeHighPriorityIpsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeHighPriorityIpsResponseBody) SetPageSize(v int32) *DescribeHighPriorityIpsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeHighPriorityIpsResponseBody) SetRequestId(v string) *DescribeHighPriorityIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHighPriorityIpsResponseBody) SetPageNumber(v int32) *DescribeHighPriorityIpsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeHighPriorityIpsResponseBody) SetHighPriorityIps(v *DescribeHighPriorityIpsResponseBodyHighPriorityIps) *DescribeHighPriorityIpsResponseBody {
	s.HighPriorityIps = v
	return s
}

type DescribeHighPriorityIpsResponseBodyHighPriorityIps struct {
	HighPriorityIp []*DescribeHighPriorityIpsResponseBodyHighPriorityIpsHighPriorityIp `json:"HighPriorityIp,omitempty" xml:"HighPriorityIp,omitempty" type:"Repeated"`
}

func (s DescribeHighPriorityIpsResponseBodyHighPriorityIps) String() string {
	return tea.Prettify(s)
}

func (s DescribeHighPriorityIpsResponseBodyHighPriorityIps) GoString() string {
	return s.String()
}

func (s *DescribeHighPriorityIpsResponseBodyHighPriorityIps) SetHighPriorityIp(v []*DescribeHighPriorityIpsResponseBodyHighPriorityIpsHighPriorityIp) *DescribeHighPriorityIpsResponseBodyHighPriorityIps {
	s.HighPriorityIp = v
	return s
}

type DescribeHighPriorityIpsResponseBodyHighPriorityIpsHighPriorityIp struct {
	AreaId      *string `json:"AreaId,omitempty" xml:"AreaId,omitempty"`
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
}

func (s DescribeHighPriorityIpsResponseBodyHighPriorityIpsHighPriorityIp) String() string {
	return tea.Prettify(s)
}

func (s DescribeHighPriorityIpsResponseBodyHighPriorityIpsHighPriorityIp) GoString() string {
	return s.String()
}

func (s *DescribeHighPriorityIpsResponseBodyHighPriorityIpsHighPriorityIp) SetAreaId(v string) *DescribeHighPriorityIpsResponseBodyHighPriorityIpsHighPriorityIp {
	s.AreaId = &v
	return s
}

func (s *DescribeHighPriorityIpsResponseBodyHighPriorityIpsHighPriorityIp) SetDestination(v string) *DescribeHighPriorityIpsResponseBodyHighPriorityIpsHighPriorityIp {
	s.Destination = &v
	return s
}

type DescribeHighPriorityIpsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeHighPriorityIpsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeHighPriorityIpsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHighPriorityIpsResponse) GoString() string {
	return s.String()
}

func (s *DescribeHighPriorityIpsResponse) SetHeaders(v map[string]*string) *DescribeHighPriorityIpsResponse {
	s.Headers = v
	return s
}

func (s *DescribeHighPriorityIpsResponse) SetBody(v *DescribeHighPriorityIpsResponseBody) *DescribeHighPriorityIpsResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetOwnerAccount(v string) *DescribeRegionsRequest {
	s.OwnerAccount = &v
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
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

type DescribeSubConnectionRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	UisSubConnectionId   *string `json:"UisSubConnectionId,omitempty" xml:"UisSubConnectionId,omitempty"`
}

func (s DescribeSubConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubConnectionRequest) GoString() string {
	return s.String()
}

func (s *DescribeSubConnectionRequest) SetOwnerAccount(v string) *DescribeSubConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSubConnectionRequest) SetOwnerId(v int64) *DescribeSubConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSubConnectionRequest) SetResourceOwnerAccount(v string) *DescribeSubConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSubConnectionRequest) SetResourceOwnerId(v int64) *DescribeSubConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSubConnectionRequest) SetUisSubConnectionId(v string) *DescribeSubConnectionRequest {
	s.UisSubConnectionId = &v
	return s
}

type DescribeSubConnectionResponseBody struct {
	CustomerTunnelIp   *string `json:"CustomerTunnelIp,omitempty" xml:"CustomerTunnelIp,omitempty"`
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	LocalTunnelIp      *string `json:"LocalTunnelIp,omitempty" xml:"LocalTunnelIp,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UisId              *string `json:"UisId,omitempty" xml:"UisId,omitempty"`
	CustomerIp         *string `json:"CustomerIp,omitempty" xml:"CustomerIp,omitempty"`
	Ip                 *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	UisNodeId          *string `json:"UisNodeId,omitempty" xml:"UisNodeId,omitempty"`
	UisSubConnectionId *string `json:"UisSubConnectionId,omitempty" xml:"UisSubConnectionId,omitempty"`
	CustomerSubnet     *string `json:"CustomerSubnet,omitempty" xml:"CustomerSubnet,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeSubConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSubConnectionResponseBody) SetCustomerTunnelIp(v string) *DescribeSubConnectionResponseBody {
	s.CustomerTunnelIp = &v
	return s
}

func (s *DescribeSubConnectionResponseBody) SetDescription(v string) *DescribeSubConnectionResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeSubConnectionResponseBody) SetLocalTunnelIp(v string) *DescribeSubConnectionResponseBody {
	s.LocalTunnelIp = &v
	return s
}

func (s *DescribeSubConnectionResponseBody) SetRequestId(v string) *DescribeSubConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSubConnectionResponseBody) SetUisId(v string) *DescribeSubConnectionResponseBody {
	s.UisId = &v
	return s
}

func (s *DescribeSubConnectionResponseBody) SetCustomerIp(v string) *DescribeSubConnectionResponseBody {
	s.CustomerIp = &v
	return s
}

func (s *DescribeSubConnectionResponseBody) SetIp(v string) *DescribeSubConnectionResponseBody {
	s.Ip = &v
	return s
}

func (s *DescribeSubConnectionResponseBody) SetUisNodeId(v string) *DescribeSubConnectionResponseBody {
	s.UisNodeId = &v
	return s
}

func (s *DescribeSubConnectionResponseBody) SetUisSubConnectionId(v string) *DescribeSubConnectionResponseBody {
	s.UisSubConnectionId = &v
	return s
}

func (s *DescribeSubConnectionResponseBody) SetCustomerSubnet(v string) *DescribeSubConnectionResponseBody {
	s.CustomerSubnet = &v
	return s
}

func (s *DescribeSubConnectionResponseBody) SetName(v string) *DescribeSubConnectionResponseBody {
	s.Name = &v
	return s
}

type DescribeSubConnectionResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSubConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSubConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubConnectionResponse) GoString() string {
	return s.String()
}

func (s *DescribeSubConnectionResponse) SetHeaders(v map[string]*string) *DescribeSubConnectionResponse {
	s.Headers = v
	return s
}

func (s *DescribeSubConnectionResponse) SetBody(v *DescribeSubConnectionResponseBody) *DescribeSubConnectionResponse {
	s.Body = v
	return s
}

type DescribeSubConnectionsRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	UisNodeId            *string `json:"UisNodeId,omitempty" xml:"UisNodeId,omitempty"`
	IP                   *string `json:"IP,omitempty" xml:"IP,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeSubConnectionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubConnectionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSubConnectionsRequest) SetOwnerAccount(v string) *DescribeSubConnectionsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSubConnectionsRequest) SetOwnerId(v int64) *DescribeSubConnectionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSubConnectionsRequest) SetResourceOwnerAccount(v string) *DescribeSubConnectionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSubConnectionsRequest) SetResourceOwnerId(v int64) *DescribeSubConnectionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSubConnectionsRequest) SetUisNodeId(v string) *DescribeSubConnectionsRequest {
	s.UisNodeId = &v
	return s
}

func (s *DescribeSubConnectionsRequest) SetIP(v string) *DescribeSubConnectionsRequest {
	s.IP = &v
	return s
}

func (s *DescribeSubConnectionsRequest) SetPageNumber(v int32) *DescribeSubConnectionsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSubConnectionsRequest) SetPageSize(v int32) *DescribeSubConnectionsRequest {
	s.PageSize = &v
	return s
}

type DescribeSubConnectionsResponseBody struct {
	TotalCount        *int32                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize          *int32                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId         *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber        *int32                                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	UisSubConnections *DescribeSubConnectionsResponseBodyUisSubConnections `json:"UisSubConnections,omitempty" xml:"UisSubConnections,omitempty" type:"Struct"`
}

func (s DescribeSubConnectionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubConnectionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSubConnectionsResponseBody) SetTotalCount(v int32) *DescribeSubConnectionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSubConnectionsResponseBody) SetPageSize(v int32) *DescribeSubConnectionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSubConnectionsResponseBody) SetRequestId(v string) *DescribeSubConnectionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSubConnectionsResponseBody) SetPageNumber(v int32) *DescribeSubConnectionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSubConnectionsResponseBody) SetUisSubConnections(v *DescribeSubConnectionsResponseBodyUisSubConnections) *DescribeSubConnectionsResponseBody {
	s.UisSubConnections = v
	return s
}

type DescribeSubConnectionsResponseBodyUisSubConnections struct {
	UisSubConnection []*DescribeSubConnectionsResponseBodyUisSubConnectionsUisSubConnection `json:"UisSubConnection,omitempty" xml:"UisSubConnection,omitempty" type:"Repeated"`
}

func (s DescribeSubConnectionsResponseBodyUisSubConnections) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubConnectionsResponseBodyUisSubConnections) GoString() string {
	return s.String()
}

func (s *DescribeSubConnectionsResponseBodyUisSubConnections) SetUisSubConnection(v []*DescribeSubConnectionsResponseBodyUisSubConnectionsUisSubConnection) *DescribeSubConnectionsResponseBodyUisSubConnections {
	s.UisSubConnection = v
	return s
}

type DescribeSubConnectionsResponseBodyUisSubConnectionsUisSubConnection struct {
	UisSubConnectionId *string `json:"UisSubConnectionId,omitempty" xml:"UisSubConnectionId,omitempty"`
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	CreateTime         *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Ip                 *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeSubConnectionsResponseBodyUisSubConnectionsUisSubConnection) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubConnectionsResponseBodyUisSubConnectionsUisSubConnection) GoString() string {
	return s.String()
}

func (s *DescribeSubConnectionsResponseBodyUisSubConnectionsUisSubConnection) SetUisSubConnectionId(v string) *DescribeSubConnectionsResponseBodyUisSubConnectionsUisSubConnection {
	s.UisSubConnectionId = &v
	return s
}

func (s *DescribeSubConnectionsResponseBodyUisSubConnectionsUisSubConnection) SetDescription(v string) *DescribeSubConnectionsResponseBodyUisSubConnectionsUisSubConnection {
	s.Description = &v
	return s
}

func (s *DescribeSubConnectionsResponseBodyUisSubConnectionsUisSubConnection) SetCreateTime(v int64) *DescribeSubConnectionsResponseBodyUisSubConnectionsUisSubConnection {
	s.CreateTime = &v
	return s
}

func (s *DescribeSubConnectionsResponseBodyUisSubConnectionsUisSubConnection) SetIp(v string) *DescribeSubConnectionsResponseBodyUisSubConnectionsUisSubConnection {
	s.Ip = &v
	return s
}

func (s *DescribeSubConnectionsResponseBodyUisSubConnectionsUisSubConnection) SetName(v string) *DescribeSubConnectionsResponseBodyUisSubConnectionsUisSubConnection {
	s.Name = &v
	return s
}

type DescribeSubConnectionsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSubConnectionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSubConnectionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubConnectionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSubConnectionsResponse) SetHeaders(v map[string]*string) *DescribeSubConnectionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSubConnectionsResponse) SetBody(v *DescribeSubConnectionsResponseBody) *DescribeSubConnectionsResponse {
	s.Body = v
	return s
}

type DescribeUisConnectionsRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UisNodeId            *string `json:"UisNodeId,omitempty" xml:"UisNodeId,omitempty"`
	UisConnectionId      *string `json:"UisConnectionId,omitempty" xml:"UisConnectionId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s DescribeUisConnectionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUisConnectionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeUisConnectionsRequest) SetOwnerAccount(v string) *DescribeUisConnectionsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeUisConnectionsRequest) SetOwnerId(v int64) *DescribeUisConnectionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeUisConnectionsRequest) SetResourceOwnerAccount(v string) *DescribeUisConnectionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeUisConnectionsRequest) SetResourceOwnerId(v int64) *DescribeUisConnectionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeUisConnectionsRequest) SetRegionId(v string) *DescribeUisConnectionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUisConnectionsRequest) SetUisNodeId(v string) *DescribeUisConnectionsRequest {
	s.UisNodeId = &v
	return s
}

func (s *DescribeUisConnectionsRequest) SetUisConnectionId(v string) *DescribeUisConnectionsRequest {
	s.UisConnectionId = &v
	return s
}

func (s *DescribeUisConnectionsRequest) SetPageNumber(v int32) *DescribeUisConnectionsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeUisConnectionsRequest) SetPageSize(v int32) *DescribeUisConnectionsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeUisConnectionsRequest) SetClientToken(v string) *DescribeUisConnectionsRequest {
	s.ClientToken = &v
	return s
}

type DescribeUisConnectionsResponseBody struct {
	TotalCount     *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize       *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber     *int32                                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	UisConnections *DescribeUisConnectionsResponseBodyUisConnections `json:"UisConnections,omitempty" xml:"UisConnections,omitempty" type:"Struct"`
}

func (s DescribeUisConnectionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUisConnectionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUisConnectionsResponseBody) SetTotalCount(v int32) *DescribeUisConnectionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeUisConnectionsResponseBody) SetPageSize(v int32) *DescribeUisConnectionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeUisConnectionsResponseBody) SetRequestId(v string) *DescribeUisConnectionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUisConnectionsResponseBody) SetPageNumber(v int32) *DescribeUisConnectionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeUisConnectionsResponseBody) SetUisConnections(v *DescribeUisConnectionsResponseBodyUisConnections) *DescribeUisConnectionsResponseBody {
	s.UisConnections = v
	return s
}

type DescribeUisConnectionsResponseBodyUisConnections struct {
	UisConnection []*DescribeUisConnectionsResponseBodyUisConnectionsUisConnection `json:"UisConnection,omitempty" xml:"UisConnection,omitempty" type:"Repeated"`
}

func (s DescribeUisConnectionsResponseBodyUisConnections) String() string {
	return tea.Prettify(s)
}

func (s DescribeUisConnectionsResponseBodyUisConnections) GoString() string {
	return s.String()
}

func (s *DescribeUisConnectionsResponseBodyUisConnections) SetUisConnection(v []*DescribeUisConnectionsResponseBodyUisConnectionsUisConnection) *DescribeUisConnectionsResponseBodyUisConnections {
	s.UisConnection = v
	return s
}

type DescribeUisConnectionsResponseBodyUisConnectionsUisConnection struct {
	GreConfig       *string `json:"GreConfig,omitempty" xml:"GreConfig,omitempty"`
	UisId           *string `json:"UisId,omitempty" xml:"UisId,omitempty"`
	UisNodeId       *string `json:"UisNodeId,omitempty" xml:"UisNodeId,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	State           *string `json:"State,omitempty" xml:"State,omitempty"`
	UisProtocol     *string `json:"UisProtocol,omitempty" xml:"UisProtocol,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SslConfig       *string `json:"SslConfig,omitempty" xml:"SslConfig,omitempty"`
	UisConnectionId *string `json:"UisConnectionId,omitempty" xml:"UisConnectionId,omitempty"`
}

func (s DescribeUisConnectionsResponseBodyUisConnectionsUisConnection) String() string {
	return tea.Prettify(s)
}

func (s DescribeUisConnectionsResponseBodyUisConnectionsUisConnection) GoString() string {
	return s.String()
}

func (s *DescribeUisConnectionsResponseBodyUisConnectionsUisConnection) SetGreConfig(v string) *DescribeUisConnectionsResponseBodyUisConnectionsUisConnection {
	s.GreConfig = &v
	return s
}

func (s *DescribeUisConnectionsResponseBodyUisConnectionsUisConnection) SetUisId(v string) *DescribeUisConnectionsResponseBodyUisConnectionsUisConnection {
	s.UisId = &v
	return s
}

func (s *DescribeUisConnectionsResponseBodyUisConnectionsUisConnection) SetUisNodeId(v string) *DescribeUisConnectionsResponseBodyUisConnectionsUisConnection {
	s.UisNodeId = &v
	return s
}

func (s *DescribeUisConnectionsResponseBodyUisConnectionsUisConnection) SetDescription(v string) *DescribeUisConnectionsResponseBodyUisConnectionsUisConnection {
	s.Description = &v
	return s
}

func (s *DescribeUisConnectionsResponseBodyUisConnectionsUisConnection) SetState(v string) *DescribeUisConnectionsResponseBodyUisConnectionsUisConnection {
	s.State = &v
	return s
}

func (s *DescribeUisConnectionsResponseBodyUisConnectionsUisConnection) SetUisProtocol(v string) *DescribeUisConnectionsResponseBodyUisConnectionsUisConnection {
	s.UisProtocol = &v
	return s
}

func (s *DescribeUisConnectionsResponseBodyUisConnectionsUisConnection) SetName(v string) *DescribeUisConnectionsResponseBodyUisConnectionsUisConnection {
	s.Name = &v
	return s
}

func (s *DescribeUisConnectionsResponseBodyUisConnectionsUisConnection) SetSslConfig(v string) *DescribeUisConnectionsResponseBodyUisConnectionsUisConnection {
	s.SslConfig = &v
	return s
}

func (s *DescribeUisConnectionsResponseBodyUisConnectionsUisConnection) SetUisConnectionId(v string) *DescribeUisConnectionsResponseBodyUisConnectionsUisConnection {
	s.UisConnectionId = &v
	return s
}

type DescribeUisConnectionsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUisConnectionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUisConnectionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUisConnectionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeUisConnectionsResponse) SetHeaders(v map[string]*string) *DescribeUisConnectionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeUisConnectionsResponse) SetBody(v *DescribeUisConnectionsResponseBody) *DescribeUisConnectionsResponse {
	s.Body = v
	return s
}

type DescribeUiseNodeStatusRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	UisNodeId            *string `json:"UisNodeId,omitempty" xml:"UisNodeId,omitempty"`
	Ip                   *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s DescribeUiseNodeStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUiseNodeStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeUiseNodeStatusRequest) SetOwnerAccount(v string) *DescribeUiseNodeStatusRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeUiseNodeStatusRequest) SetOwnerId(v int64) *DescribeUiseNodeStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeUiseNodeStatusRequest) SetResourceOwnerAccount(v string) *DescribeUiseNodeStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeUiseNodeStatusRequest) SetResourceOwnerId(v int64) *DescribeUiseNodeStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeUiseNodeStatusRequest) SetUisNodeId(v string) *DescribeUiseNodeStatusRequest {
	s.UisNodeId = &v
	return s
}

func (s *DescribeUiseNodeStatusRequest) SetIp(v string) *DescribeUiseNodeStatusRequest {
	s.Ip = &v
	return s
}

type DescribeUiseNodeStatusResponseBody struct {
	RequestId    *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	IpStatusList *DescribeUiseNodeStatusResponseBodyIpStatusList `json:"IpStatusList,omitempty" xml:"IpStatusList,omitempty" type:"Struct"`
}

func (s DescribeUiseNodeStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUiseNodeStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUiseNodeStatusResponseBody) SetRequestId(v string) *DescribeUiseNodeStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUiseNodeStatusResponseBody) SetIpStatusList(v *DescribeUiseNodeStatusResponseBodyIpStatusList) *DescribeUiseNodeStatusResponseBody {
	s.IpStatusList = v
	return s
}

type DescribeUiseNodeStatusResponseBodyIpStatusList struct {
	IpStatus []*DescribeUiseNodeStatusResponseBodyIpStatusListIpStatus `json:"IpStatus,omitempty" xml:"IpStatus,omitempty" type:"Repeated"`
}

func (s DescribeUiseNodeStatusResponseBodyIpStatusList) String() string {
	return tea.Prettify(s)
}

func (s DescribeUiseNodeStatusResponseBodyIpStatusList) GoString() string {
	return s.String()
}

func (s *DescribeUiseNodeStatusResponseBodyIpStatusList) SetIpStatus(v []*DescribeUiseNodeStatusResponseBodyIpStatusListIpStatus) *DescribeUiseNodeStatusResponseBodyIpStatusList {
	s.IpStatus = v
	return s
}

type DescribeUiseNodeStatusResponseBodyIpStatusListIpStatus struct {
	CurrentConnections *int32  `json:"CurrentConnections,omitempty" xml:"CurrentConnections,omitempty"`
	Ip                 *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	MaxConnections     *int32  `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty"`
}

func (s DescribeUiseNodeStatusResponseBodyIpStatusListIpStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeUiseNodeStatusResponseBodyIpStatusListIpStatus) GoString() string {
	return s.String()
}

func (s *DescribeUiseNodeStatusResponseBodyIpStatusListIpStatus) SetCurrentConnections(v int32) *DescribeUiseNodeStatusResponseBodyIpStatusListIpStatus {
	s.CurrentConnections = &v
	return s
}

func (s *DescribeUiseNodeStatusResponseBodyIpStatusListIpStatus) SetIp(v string) *DescribeUiseNodeStatusResponseBodyIpStatusListIpStatus {
	s.Ip = &v
	return s
}

func (s *DescribeUiseNodeStatusResponseBodyIpStatusListIpStatus) SetMaxConnections(v int32) *DescribeUiseNodeStatusResponseBodyIpStatusListIpStatus {
	s.MaxConnections = &v
	return s
}

type DescribeUiseNodeStatusResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUiseNodeStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUiseNodeStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUiseNodeStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeUiseNodeStatusResponse) SetHeaders(v map[string]*string) *DescribeUiseNodeStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeUiseNodeStatusResponse) SetBody(v *DescribeUiseNodeStatusResponseBody) *DescribeUiseNodeStatusResponse {
	s.Body = v
	return s
}

type DescribeUisesRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	UisId                *string `json:"UisId,omitempty" xml:"UisId,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeUisesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUisesRequest) GoString() string {
	return s.String()
}

func (s *DescribeUisesRequest) SetOwnerAccount(v string) *DescribeUisesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeUisesRequest) SetOwnerId(v int64) *DescribeUisesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeUisesRequest) SetResourceOwnerAccount(v string) *DescribeUisesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeUisesRequest) SetResourceOwnerId(v int64) *DescribeUisesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeUisesRequest) SetUisId(v string) *DescribeUisesRequest {
	s.UisId = &v
	return s
}

func (s *DescribeUisesRequest) SetName(v string) *DescribeUisesRequest {
	s.Name = &v
	return s
}

func (s *DescribeUisesRequest) SetPageNumber(v int32) *DescribeUisesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeUisesRequest) SetPageSize(v int32) *DescribeUisesRequest {
	s.PageSize = &v
	return s
}

type DescribeUisesResponseBody struct {
	TotalCount *int32                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Uises      *DescribeUisesResponseBodyUises `json:"Uises,omitempty" xml:"Uises,omitempty" type:"Struct"`
	PageSize   *int32                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeUisesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUisesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUisesResponseBody) SetTotalCount(v int32) *DescribeUisesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeUisesResponseBody) SetUises(v *DescribeUisesResponseBodyUises) *DescribeUisesResponseBody {
	s.Uises = v
	return s
}

func (s *DescribeUisesResponseBody) SetPageSize(v int32) *DescribeUisesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeUisesResponseBody) SetRequestId(v string) *DescribeUisesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUisesResponseBody) SetPageNumber(v int32) *DescribeUisesResponseBody {
	s.PageNumber = &v
	return s
}

type DescribeUisesResponseBodyUises struct {
	Uis []*DescribeUisesResponseBodyUisesUis `json:"Uis,omitempty" xml:"Uis,omitempty" type:"Repeated"`
}

func (s DescribeUisesResponseBodyUises) String() string {
	return tea.Prettify(s)
}

func (s DescribeUisesResponseBodyUises) GoString() string {
	return s.String()
}

func (s *DescribeUisesResponseBodyUises) SetUis(v []*DescribeUisesResponseBodyUisesUis) *DescribeUisesResponseBodyUises {
	s.Uis = v
	return s
}

type DescribeUisesResponseBodyUisesUis struct {
	Status              *string                                      `json:"Status,omitempty" xml:"Status,omitempty"`
	UisId               *string                                      `json:"UisId,omitempty" xml:"UisId,omitempty"`
	ConnectionType      *string                                      `json:"ConnectionType,omitempty" xml:"ConnectionType,omitempty"`
	BandwidthType       *string                                      `json:"BandwidthType,omitempty" xml:"BandwidthType,omitempty"`
	CreateTime          *int64                                       `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ChargeType          *string                                      `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	PayType             *string                                      `json:"PayType,omitempty" xml:"PayType,omitempty"`
	UserInfo            *DescribeUisesResponseBodyUisesUisUserInfo   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
	SslClientCertUrl    *string                                      `json:"SslClientCertUrl,omitempty" xml:"SslClientCertUrl,omitempty"`
	ConnectionCount     *int32                                       `json:"ConnectionCount,omitempty" xml:"ConnectionCount,omitempty"`
	UisNodeIds          *DescribeUisesResponseBodyUisesUisUisNodeIds `json:"UisNodeIds,omitempty" xml:"UisNodeIds,omitempty" type:"Struct"`
	EndTime             *int64                                       `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Bandwidth           *int32                                       `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	Description         *string                                      `json:"Description,omitempty" xml:"Description,omitempty"`
	ServiceRegion       *string                                      `json:"ServiceRegion,omitempty" xml:"ServiceRegion,omitempty"`
	ConnectionBandwidth *int32                                       `json:"ConnectionBandwidth,omitempty" xml:"ConnectionBandwidth,omitempty"`
	BusinessStatus      *string                                      `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	Name                *string                                      `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeUisesResponseBodyUisesUis) String() string {
	return tea.Prettify(s)
}

func (s DescribeUisesResponseBodyUisesUis) GoString() string {
	return s.String()
}

func (s *DescribeUisesResponseBodyUisesUis) SetStatus(v string) *DescribeUisesResponseBodyUisesUis {
	s.Status = &v
	return s
}

func (s *DescribeUisesResponseBodyUisesUis) SetUisId(v string) *DescribeUisesResponseBodyUisesUis {
	s.UisId = &v
	return s
}

func (s *DescribeUisesResponseBodyUisesUis) SetConnectionType(v string) *DescribeUisesResponseBodyUisesUis {
	s.ConnectionType = &v
	return s
}

func (s *DescribeUisesResponseBodyUisesUis) SetBandwidthType(v string) *DescribeUisesResponseBodyUisesUis {
	s.BandwidthType = &v
	return s
}

func (s *DescribeUisesResponseBodyUisesUis) SetCreateTime(v int64) *DescribeUisesResponseBodyUisesUis {
	s.CreateTime = &v
	return s
}

func (s *DescribeUisesResponseBodyUisesUis) SetChargeType(v string) *DescribeUisesResponseBodyUisesUis {
	s.ChargeType = &v
	return s
}

func (s *DescribeUisesResponseBodyUisesUis) SetPayType(v string) *DescribeUisesResponseBodyUisesUis {
	s.PayType = &v
	return s
}

func (s *DescribeUisesResponseBodyUisesUis) SetUserInfo(v *DescribeUisesResponseBodyUisesUisUserInfo) *DescribeUisesResponseBodyUisesUis {
	s.UserInfo = v
	return s
}

func (s *DescribeUisesResponseBodyUisesUis) SetSslClientCertUrl(v string) *DescribeUisesResponseBodyUisesUis {
	s.SslClientCertUrl = &v
	return s
}

func (s *DescribeUisesResponseBodyUisesUis) SetConnectionCount(v int32) *DescribeUisesResponseBodyUisesUis {
	s.ConnectionCount = &v
	return s
}

func (s *DescribeUisesResponseBodyUisesUis) SetUisNodeIds(v *DescribeUisesResponseBodyUisesUisUisNodeIds) *DescribeUisesResponseBodyUisesUis {
	s.UisNodeIds = v
	return s
}

func (s *DescribeUisesResponseBodyUisesUis) SetEndTime(v int64) *DescribeUisesResponseBodyUisesUis {
	s.EndTime = &v
	return s
}

func (s *DescribeUisesResponseBodyUisesUis) SetBandwidth(v int32) *DescribeUisesResponseBodyUisesUis {
	s.Bandwidth = &v
	return s
}

func (s *DescribeUisesResponseBodyUisesUis) SetDescription(v string) *DescribeUisesResponseBodyUisesUis {
	s.Description = &v
	return s
}

func (s *DescribeUisesResponseBodyUisesUis) SetServiceRegion(v string) *DescribeUisesResponseBodyUisesUis {
	s.ServiceRegion = &v
	return s
}

func (s *DescribeUisesResponseBodyUisesUis) SetConnectionBandwidth(v int32) *DescribeUisesResponseBodyUisesUis {
	s.ConnectionBandwidth = &v
	return s
}

func (s *DescribeUisesResponseBodyUisesUis) SetBusinessStatus(v string) *DescribeUisesResponseBodyUisesUis {
	s.BusinessStatus = &v
	return s
}

func (s *DescribeUisesResponseBodyUisesUis) SetName(v string) *DescribeUisesResponseBodyUisesUis {
	s.Name = &v
	return s
}

type DescribeUisesResponseBodyUisesUisUserInfo struct {
	ClientInfoDBAccount  *string `json:"ClientInfoDBAccount,omitempty" xml:"ClientInfoDBAccount,omitempty"`
	ClientInfoDB         *string `json:"ClientInfoDB,omitempty" xml:"ClientInfoDB,omitempty"`
	ClientInfoDBPassword *string `json:"ClientInfoDBPassword,omitempty" xml:"ClientInfoDBPassword,omitempty"`
}

func (s DescribeUisesResponseBodyUisesUisUserInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeUisesResponseBodyUisesUisUserInfo) GoString() string {
	return s.String()
}

func (s *DescribeUisesResponseBodyUisesUisUserInfo) SetClientInfoDBAccount(v string) *DescribeUisesResponseBodyUisesUisUserInfo {
	s.ClientInfoDBAccount = &v
	return s
}

func (s *DescribeUisesResponseBodyUisesUisUserInfo) SetClientInfoDB(v string) *DescribeUisesResponseBodyUisesUisUserInfo {
	s.ClientInfoDB = &v
	return s
}

func (s *DescribeUisesResponseBodyUisesUisUserInfo) SetClientInfoDBPassword(v string) *DescribeUisesResponseBodyUisesUisUserInfo {
	s.ClientInfoDBPassword = &v
	return s
}

type DescribeUisesResponseBodyUisesUisUisNodeIds struct {
	UisNodeIds []*string `json:"UisNodeIds,omitempty" xml:"UisNodeIds,omitempty" type:"Repeated"`
}

func (s DescribeUisesResponseBodyUisesUisUisNodeIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeUisesResponseBodyUisesUisUisNodeIds) GoString() string {
	return s.String()
}

func (s *DescribeUisesResponseBodyUisesUisUisNodeIds) SetUisNodeIds(v []*string) *DescribeUisesResponseBodyUisesUisUisNodeIds {
	s.UisNodeIds = v
	return s
}

type DescribeUisesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUisesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUisesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUisesResponse) GoString() string {
	return s.String()
}

func (s *DescribeUisesResponse) SetHeaders(v map[string]*string) *DescribeUisesResponse {
	s.Headers = v
	return s
}

func (s *DescribeUisesResponse) SetBody(v *DescribeUisesResponseBody) *DescribeUisesResponse {
	s.Body = v
	return s
}

type DescribeUisNetworkInterfacesRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UisNodeId            *string `json:"UisNodeId,omitempty" xml:"UisNodeId,omitempty"`
	UisEniId             *string `json:"UisEniId,omitempty" xml:"UisEniId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeUisNetworkInterfacesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUisNetworkInterfacesRequest) GoString() string {
	return s.String()
}

func (s *DescribeUisNetworkInterfacesRequest) SetOwnerAccount(v string) *DescribeUisNetworkInterfacesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeUisNetworkInterfacesRequest) SetOwnerId(v int64) *DescribeUisNetworkInterfacesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeUisNetworkInterfacesRequest) SetResourceOwnerAccount(v string) *DescribeUisNetworkInterfacesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeUisNetworkInterfacesRequest) SetResourceOwnerId(v int64) *DescribeUisNetworkInterfacesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeUisNetworkInterfacesRequest) SetRegionId(v string) *DescribeUisNetworkInterfacesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUisNetworkInterfacesRequest) SetUisNodeId(v string) *DescribeUisNetworkInterfacesRequest {
	s.UisNodeId = &v
	return s
}

func (s *DescribeUisNetworkInterfacesRequest) SetUisEniId(v string) *DescribeUisNetworkInterfacesRequest {
	s.UisEniId = &v
	return s
}

func (s *DescribeUisNetworkInterfacesRequest) SetPageNumber(v int32) *DescribeUisNetworkInterfacesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeUisNetworkInterfacesRequest) SetPageSize(v int32) *DescribeUisNetworkInterfacesRequest {
	s.PageSize = &v
	return s
}

type DescribeUisNetworkInterfacesResponseBody struct {
	TotalCount        *int32                                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize          *int32                                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId         *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber        *int32                                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	NetworkInterfaces *DescribeUisNetworkInterfacesResponseBodyNetworkInterfaces `json:"NetworkInterfaces,omitempty" xml:"NetworkInterfaces,omitempty" type:"Struct"`
}

func (s DescribeUisNetworkInterfacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUisNetworkInterfacesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUisNetworkInterfacesResponseBody) SetTotalCount(v int32) *DescribeUisNetworkInterfacesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeUisNetworkInterfacesResponseBody) SetPageSize(v int32) *DescribeUisNetworkInterfacesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeUisNetworkInterfacesResponseBody) SetRequestId(v string) *DescribeUisNetworkInterfacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUisNetworkInterfacesResponseBody) SetPageNumber(v int32) *DescribeUisNetworkInterfacesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeUisNetworkInterfacesResponseBody) SetNetworkInterfaces(v *DescribeUisNetworkInterfacesResponseBodyNetworkInterfaces) *DescribeUisNetworkInterfacesResponseBody {
	s.NetworkInterfaces = v
	return s
}

type DescribeUisNetworkInterfacesResponseBodyNetworkInterfaces struct {
	NetworkInterface []*DescribeUisNetworkInterfacesResponseBodyNetworkInterfacesNetworkInterface `json:"NetworkInterface,omitempty" xml:"NetworkInterface,omitempty" type:"Repeated"`
}

func (s DescribeUisNetworkInterfacesResponseBodyNetworkInterfaces) String() string {
	return tea.Prettify(s)
}

func (s DescribeUisNetworkInterfacesResponseBodyNetworkInterfaces) GoString() string {
	return s.String()
}

func (s *DescribeUisNetworkInterfacesResponseBodyNetworkInterfaces) SetNetworkInterface(v []*DescribeUisNetworkInterfacesResponseBodyNetworkInterfacesNetworkInterface) *DescribeUisNetworkInterfacesResponseBodyNetworkInterfaces {
	s.NetworkInterface = v
	return s
}

type DescribeUisNetworkInterfacesResponseBodyNetworkInterfacesNetworkInterface struct {
	UisEniId           *string `json:"UisEniId,omitempty" xml:"UisEniId,omitempty"`
	UisNodeId          *string `json:"UisNodeId,omitempty" xml:"UisNodeId,omitempty"`
	IpAddress          *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	VswitchId          *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	State              *string `json:"State,omitempty" xml:"State,omitempty"`
	SecurityGroupID    *string `json:"SecurityGroupID,omitempty" xml:"SecurityGroupID,omitempty"`
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Log                *string `json:"Log,omitempty" xml:"Log,omitempty"`
}

func (s DescribeUisNetworkInterfacesResponseBodyNetworkInterfacesNetworkInterface) String() string {
	return tea.Prettify(s)
}

func (s DescribeUisNetworkInterfacesResponseBodyNetworkInterfacesNetworkInterface) GoString() string {
	return s.String()
}

func (s *DescribeUisNetworkInterfacesResponseBodyNetworkInterfacesNetworkInterface) SetUisEniId(v string) *DescribeUisNetworkInterfacesResponseBodyNetworkInterfacesNetworkInterface {
	s.UisEniId = &v
	return s
}

func (s *DescribeUisNetworkInterfacesResponseBodyNetworkInterfacesNetworkInterface) SetUisNodeId(v string) *DescribeUisNetworkInterfacesResponseBodyNetworkInterfacesNetworkInterface {
	s.UisNodeId = &v
	return s
}

func (s *DescribeUisNetworkInterfacesResponseBodyNetworkInterfacesNetworkInterface) SetIpAddress(v string) *DescribeUisNetworkInterfacesResponseBodyNetworkInterfacesNetworkInterface {
	s.IpAddress = &v
	return s
}

func (s *DescribeUisNetworkInterfacesResponseBodyNetworkInterfacesNetworkInterface) SetDescription(v string) *DescribeUisNetworkInterfacesResponseBodyNetworkInterfacesNetworkInterface {
	s.Description = &v
	return s
}

func (s *DescribeUisNetworkInterfacesResponseBodyNetworkInterfacesNetworkInterface) SetVswitchId(v string) *DescribeUisNetworkInterfacesResponseBodyNetworkInterfacesNetworkInterface {
	s.VswitchId = &v
	return s
}

func (s *DescribeUisNetworkInterfacesResponseBodyNetworkInterfacesNetworkInterface) SetState(v string) *DescribeUisNetworkInterfacesResponseBodyNetworkInterfacesNetworkInterface {
	s.State = &v
	return s
}

func (s *DescribeUisNetworkInterfacesResponseBodyNetworkInterfacesNetworkInterface) SetSecurityGroupID(v string) *DescribeUisNetworkInterfacesResponseBodyNetworkInterfacesNetworkInterface {
	s.SecurityGroupID = &v
	return s
}

func (s *DescribeUisNetworkInterfacesResponseBodyNetworkInterfacesNetworkInterface) SetNetworkInterfaceId(v string) *DescribeUisNetworkInterfacesResponseBodyNetworkInterfacesNetworkInterface {
	s.NetworkInterfaceId = &v
	return s
}

func (s *DescribeUisNetworkInterfacesResponseBodyNetworkInterfacesNetworkInterface) SetName(v string) *DescribeUisNetworkInterfacesResponseBodyNetworkInterfacesNetworkInterface {
	s.Name = &v
	return s
}

func (s *DescribeUisNetworkInterfacesResponseBodyNetworkInterfacesNetworkInterface) SetLog(v string) *DescribeUisNetworkInterfacesResponseBodyNetworkInterfacesNetworkInterface {
	s.Log = &v
	return s
}

type DescribeUisNetworkInterfacesResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUisNetworkInterfacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUisNetworkInterfacesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUisNetworkInterfacesResponse) GoString() string {
	return s.String()
}

func (s *DescribeUisNetworkInterfacesResponse) SetHeaders(v map[string]*string) *DescribeUisNetworkInterfacesResponse {
	s.Headers = v
	return s
}

func (s *DescribeUisNetworkInterfacesResponse) SetBody(v *DescribeUisNetworkInterfacesResponseBody) *DescribeUisNetworkInterfacesResponse {
	s.Body = v
	return s
}

type DescribeUisNodesRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UisId                *string `json:"UisId,omitempty" xml:"UisId,omitempty"`
	UisNodeId            *string `json:"UisNodeId,omitempty" xml:"UisNodeId,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s DescribeUisNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUisNodesRequest) GoString() string {
	return s.String()
}

func (s *DescribeUisNodesRequest) SetOwnerAccount(v string) *DescribeUisNodesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeUisNodesRequest) SetOwnerId(v int64) *DescribeUisNodesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeUisNodesRequest) SetResourceOwnerAccount(v string) *DescribeUisNodesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeUisNodesRequest) SetResourceOwnerId(v int64) *DescribeUisNodesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeUisNodesRequest) SetRegionId(v string) *DescribeUisNodesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUisNodesRequest) SetUisId(v string) *DescribeUisNodesRequest {
	s.UisId = &v
	return s
}

func (s *DescribeUisNodesRequest) SetUisNodeId(v string) *DescribeUisNodesRequest {
	s.UisNodeId = &v
	return s
}

func (s *DescribeUisNodesRequest) SetName(v string) *DescribeUisNodesRequest {
	s.Name = &v
	return s
}

func (s *DescribeUisNodesRequest) SetStatus(v string) *DescribeUisNodesRequest {
	s.Status = &v
	return s
}

func (s *DescribeUisNodesRequest) SetPageNumber(v int32) *DescribeUisNodesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeUisNodesRequest) SetPageSize(v int32) *DescribeUisNodesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeUisNodesRequest) SetClientToken(v string) *DescribeUisNodesRequest {
	s.ClientToken = &v
	return s
}

type DescribeUisNodesResponseBody struct {
	TotalCount  *int32                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize    *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber  *int32                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	UisNodeList *DescribeUisNodesResponseBodyUisNodeList `json:"UisNodeList,omitempty" xml:"UisNodeList,omitempty" type:"Struct"`
}

func (s DescribeUisNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUisNodesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUisNodesResponseBody) SetTotalCount(v int32) *DescribeUisNodesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeUisNodesResponseBody) SetPageSize(v int32) *DescribeUisNodesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeUisNodesResponseBody) SetRequestId(v string) *DescribeUisNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUisNodesResponseBody) SetPageNumber(v int32) *DescribeUisNodesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeUisNodesResponseBody) SetUisNodeList(v *DescribeUisNodesResponseBodyUisNodeList) *DescribeUisNodesResponseBody {
	s.UisNodeList = v
	return s
}

type DescribeUisNodesResponseBodyUisNodeList struct {
	UisNode []*DescribeUisNodesResponseBodyUisNodeListUisNode `json:"UisNode,omitempty" xml:"UisNode,omitempty" type:"Repeated"`
}

func (s DescribeUisNodesResponseBodyUisNodeList) String() string {
	return tea.Prettify(s)
}

func (s DescribeUisNodesResponseBodyUisNodeList) GoString() string {
	return s.String()
}

func (s *DescribeUisNodesResponseBodyUisNodeList) SetUisNode(v []*DescribeUisNodesResponseBodyUisNodeListUisNode) *DescribeUisNodesResponseBodyUisNodeList {
	s.UisNode = v
	return s
}

type DescribeUisNodesResponseBodyUisNodeListUisNode struct {
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UisId            *string `json:"UisId,omitempty" xml:"UisId,omitempty"`
	UisNodeActiveIp  *string `json:"UisNodeActiveIp,omitempty" xml:"UisNodeActiveIp,omitempty"`
	CreateTime       *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	UisEniIps        *string `json:"UisEniIps,omitempty" xml:"UisEniIps,omitempty"`
	UisNodeAreaId    *string `json:"UisNodeAreaId,omitempty" xml:"UisNodeAreaId,omitempty"`
	UisNodeId        *string `json:"UisNodeId,omitempty" xml:"UisNodeId,omitempty"`
	UisNodeIps       *string `json:"UisNodeIps,omitempty" xml:"UisNodeIps,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	UisNodeBandwidth *int32  `json:"UisNodeBandwidth,omitempty" xml:"UisNodeBandwidth,omitempty"`
	IpAddrsNum       *int32  `json:"IpAddrsNum,omitempty" xml:"IpAddrsNum,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeUisNodesResponseBodyUisNodeListUisNode) String() string {
	return tea.Prettify(s)
}

func (s DescribeUisNodesResponseBodyUisNodeListUisNode) GoString() string {
	return s.String()
}

func (s *DescribeUisNodesResponseBodyUisNodeListUisNode) SetStatus(v string) *DescribeUisNodesResponseBodyUisNodeListUisNode {
	s.Status = &v
	return s
}

func (s *DescribeUisNodesResponseBodyUisNodeListUisNode) SetUisId(v string) *DescribeUisNodesResponseBodyUisNodeListUisNode {
	s.UisId = &v
	return s
}

func (s *DescribeUisNodesResponseBodyUisNodeListUisNode) SetUisNodeActiveIp(v string) *DescribeUisNodesResponseBodyUisNodeListUisNode {
	s.UisNodeActiveIp = &v
	return s
}

func (s *DescribeUisNodesResponseBodyUisNodeListUisNode) SetCreateTime(v int64) *DescribeUisNodesResponseBodyUisNodeListUisNode {
	s.CreateTime = &v
	return s
}

func (s *DescribeUisNodesResponseBodyUisNodeListUisNode) SetUisEniIps(v string) *DescribeUisNodesResponseBodyUisNodeListUisNode {
	s.UisEniIps = &v
	return s
}

func (s *DescribeUisNodesResponseBodyUisNodeListUisNode) SetUisNodeAreaId(v string) *DescribeUisNodesResponseBodyUisNodeListUisNode {
	s.UisNodeAreaId = &v
	return s
}

func (s *DescribeUisNodesResponseBodyUisNodeListUisNode) SetUisNodeId(v string) *DescribeUisNodesResponseBodyUisNodeListUisNode {
	s.UisNodeId = &v
	return s
}

func (s *DescribeUisNodesResponseBodyUisNodeListUisNode) SetUisNodeIps(v string) *DescribeUisNodesResponseBodyUisNodeListUisNode {
	s.UisNodeIps = &v
	return s
}

func (s *DescribeUisNodesResponseBodyUisNodeListUisNode) SetDescription(v string) *DescribeUisNodesResponseBodyUisNodeListUisNode {
	s.Description = &v
	return s
}

func (s *DescribeUisNodesResponseBodyUisNodeListUisNode) SetUisNodeBandwidth(v int32) *DescribeUisNodesResponseBodyUisNodeListUisNode {
	s.UisNodeBandwidth = &v
	return s
}

func (s *DescribeUisNodesResponseBodyUisNodeListUisNode) SetIpAddrsNum(v int32) *DescribeUisNodesResponseBodyUisNodeListUisNode {
	s.IpAddrsNum = &v
	return s
}

func (s *DescribeUisNodesResponseBodyUisNodeListUisNode) SetName(v string) *DescribeUisNodesResponseBodyUisNodeListUisNode {
	s.Name = &v
	return s
}

type DescribeUisNodesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUisNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUisNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUisNodesResponse) GoString() string {
	return s.String()
}

func (s *DescribeUisNodesResponse) SetHeaders(v map[string]*string) *DescribeUisNodesResponse {
	s.Headers = v
	return s
}

func (s *DescribeUisNodesResponse) SetBody(v *DescribeUisNodesResponseBody) *DescribeUisNodesResponse {
	s.Body = v
	return s
}

type DescribeWhiteListRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UisId                *string `json:"UisId,omitempty" xml:"UisId,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s DescribeWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWhiteListRequest) GoString() string {
	return s.String()
}

func (s *DescribeWhiteListRequest) SetOwnerAccount(v string) *DescribeWhiteListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeWhiteListRequest) SetOwnerId(v int64) *DescribeWhiteListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeWhiteListRequest) SetResourceOwnerAccount(v string) *DescribeWhiteListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeWhiteListRequest) SetResourceOwnerId(v int64) *DescribeWhiteListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeWhiteListRequest) SetRegionId(v string) *DescribeWhiteListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeWhiteListRequest) SetUisId(v string) *DescribeWhiteListRequest {
	s.UisId = &v
	return s
}

func (s *DescribeWhiteListRequest) SetClientToken(v string) *DescribeWhiteListRequest {
	s.ClientToken = &v
	return s
}

type DescribeWhiteListResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Whitelist *string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty"`
}

func (s DescribeWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWhiteListResponseBody) SetRequestId(v string) *DescribeWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWhiteListResponseBody) SetWhitelist(v string) *DescribeWhiteListResponseBody {
	s.Whitelist = &v
	return s
}

type DescribeWhiteListResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWhiteListResponse) GoString() string {
	return s.String()
}

func (s *DescribeWhiteListResponse) SetHeaders(v map[string]*string) *DescribeWhiteListResponse {
	s.Headers = v
	return s
}

func (s *DescribeWhiteListResponse) SetBody(v *DescribeWhiteListResponseBody) *DescribeWhiteListResponse {
	s.Body = v
	return s
}

type GetDroppedIpListRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	UisId                *string `json:"UisId,omitempty" xml:"UisId,omitempty"`
	ChartDate            *string `json:"ChartDate,omitempty" xml:"ChartDate,omitempty"`
}

func (s GetDroppedIpListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDroppedIpListRequest) GoString() string {
	return s.String()
}

func (s *GetDroppedIpListRequest) SetOwnerAccount(v string) *GetDroppedIpListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetDroppedIpListRequest) SetOwnerId(v int64) *GetDroppedIpListRequest {
	s.OwnerId = &v
	return s
}

func (s *GetDroppedIpListRequest) SetResourceOwnerAccount(v string) *GetDroppedIpListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetDroppedIpListRequest) SetResourceOwnerId(v int64) *GetDroppedIpListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetDroppedIpListRequest) SetUisId(v string) *GetDroppedIpListRequest {
	s.UisId = &v
	return s
}

func (s *GetDroppedIpListRequest) SetChartDate(v string) *GetDroppedIpListRequest {
	s.ChartDate = &v
	return s
}

type GetDroppedIpListResponseBody struct {
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DroppedIpListUrl *string `json:"DroppedIpListUrl,omitempty" xml:"DroppedIpListUrl,omitempty"`
}

func (s GetDroppedIpListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDroppedIpListResponseBody) GoString() string {
	return s.String()
}

func (s *GetDroppedIpListResponseBody) SetRequestId(v string) *GetDroppedIpListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDroppedIpListResponseBody) SetDroppedIpListUrl(v string) *GetDroppedIpListResponseBody {
	s.DroppedIpListUrl = &v
	return s
}

type GetDroppedIpListResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDroppedIpListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDroppedIpListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDroppedIpListResponse) GoString() string {
	return s.String()
}

func (s *GetDroppedIpListResponse) SetHeaders(v map[string]*string) *GetDroppedIpListResponse {
	s.Headers = v
	return s
}

func (s *GetDroppedIpListResponse) SetBody(v *GetDroppedIpListResponseBody) *GetDroppedIpListResponse {
	s.Body = v
	return s
}

type ModifyDnatEntryRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UisNodeId            *string `json:"UisNodeId,omitempty" xml:"UisNodeId,omitempty"`
	UisDnatId            *string `json:"UisDnatId,omitempty" xml:"UisDnatId,omitempty"`
	DestinationIp        *string `json:"DestinationIp,omitempty" xml:"DestinationIp,omitempty"`
	DestinationPort      *int32  `json:"DestinationPort,omitempty" xml:"DestinationPort,omitempty"`
	OriginalIp           *string `json:"OriginalIp,omitempty" xml:"OriginalIp,omitempty"`
	OriginalPort         *int32  `json:"OriginalPort,omitempty" xml:"OriginalPort,omitempty"`
	IpProtocol           *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ModifyDnatEntryRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDnatEntryRequest) GoString() string {
	return s.String()
}

func (s *ModifyDnatEntryRequest) SetOwnerAccount(v string) *ModifyDnatEntryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDnatEntryRequest) SetOwnerId(v int64) *ModifyDnatEntryRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDnatEntryRequest) SetResourceOwnerAccount(v string) *ModifyDnatEntryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDnatEntryRequest) SetResourceOwnerId(v int64) *ModifyDnatEntryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDnatEntryRequest) SetRegionId(v string) *ModifyDnatEntryRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDnatEntryRequest) SetUisNodeId(v string) *ModifyDnatEntryRequest {
	s.UisNodeId = &v
	return s
}

func (s *ModifyDnatEntryRequest) SetUisDnatId(v string) *ModifyDnatEntryRequest {
	s.UisDnatId = &v
	return s
}

func (s *ModifyDnatEntryRequest) SetDestinationIp(v string) *ModifyDnatEntryRequest {
	s.DestinationIp = &v
	return s
}

func (s *ModifyDnatEntryRequest) SetDestinationPort(v int32) *ModifyDnatEntryRequest {
	s.DestinationPort = &v
	return s
}

func (s *ModifyDnatEntryRequest) SetOriginalIp(v string) *ModifyDnatEntryRequest {
	s.OriginalIp = &v
	return s
}

func (s *ModifyDnatEntryRequest) SetOriginalPort(v int32) *ModifyDnatEntryRequest {
	s.OriginalPort = &v
	return s
}

func (s *ModifyDnatEntryRequest) SetIpProtocol(v string) *ModifyDnatEntryRequest {
	s.IpProtocol = &v
	return s
}

func (s *ModifyDnatEntryRequest) SetName(v string) *ModifyDnatEntryRequest {
	s.Name = &v
	return s
}

type ModifyDnatEntryResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDnatEntryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDnatEntryResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDnatEntryResponseBody) SetRequestId(v string) *ModifyDnatEntryResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDnatEntryResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDnatEntryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDnatEntryResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDnatEntryResponse) GoString() string {
	return s.String()
}

func (s *ModifyDnatEntryResponse) SetHeaders(v map[string]*string) *ModifyDnatEntryResponse {
	s.Headers = v
	return s
}

func (s *ModifyDnatEntryResponse) SetBody(v *ModifyDnatEntryResponseBody) *ModifyDnatEntryResponse {
	s.Body = v
	return s
}

type ModifyHighPriorityIpRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	UisId                *string `json:"UisId,omitempty" xml:"UisId,omitempty"`
	HighPriorityIp       *string `json:"HighPriorityIp,omitempty" xml:"HighPriorityIp,omitempty"`
}

func (s ModifyHighPriorityIpRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyHighPriorityIpRequest) GoString() string {
	return s.String()
}

func (s *ModifyHighPriorityIpRequest) SetOwnerAccount(v string) *ModifyHighPriorityIpRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyHighPriorityIpRequest) SetOwnerId(v int64) *ModifyHighPriorityIpRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyHighPriorityIpRequest) SetResourceOwnerAccount(v string) *ModifyHighPriorityIpRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyHighPriorityIpRequest) SetResourceOwnerId(v int64) *ModifyHighPriorityIpRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyHighPriorityIpRequest) SetRegionId(v string) *ModifyHighPriorityIpRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyHighPriorityIpRequest) SetClientToken(v string) *ModifyHighPriorityIpRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyHighPriorityIpRequest) SetUisId(v string) *ModifyHighPriorityIpRequest {
	s.UisId = &v
	return s
}

func (s *ModifyHighPriorityIpRequest) SetHighPriorityIp(v string) *ModifyHighPriorityIpRequest {
	s.HighPriorityIp = &v
	return s
}

type ModifyHighPriorityIpResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyHighPriorityIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyHighPriorityIpResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHighPriorityIpResponseBody) SetRequestId(v string) *ModifyHighPriorityIpResponseBody {
	s.RequestId = &v
	return s
}

type ModifyHighPriorityIpResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyHighPriorityIpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyHighPriorityIpResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyHighPriorityIpResponse) GoString() string {
	return s.String()
}

func (s *ModifyHighPriorityIpResponse) SetHeaders(v map[string]*string) *ModifyHighPriorityIpResponse {
	s.Headers = v
	return s
}

func (s *ModifyHighPriorityIpResponse) SetBody(v *ModifyHighPriorityIpResponseBody) *ModifyHighPriorityIpResponse {
	s.Body = v
	return s
}

type ModifySubConnectionRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	UisSubConnectionId   *string `json:"UisSubConnectionId,omitempty" xml:"UisSubConnectionId,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	GreConfig            *string `json:"GreConfig,omitempty" xml:"GreConfig,omitempty"`
}

func (s ModifySubConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySubConnectionRequest) GoString() string {
	return s.String()
}

func (s *ModifySubConnectionRequest) SetOwnerAccount(v string) *ModifySubConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySubConnectionRequest) SetOwnerId(v int64) *ModifySubConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySubConnectionRequest) SetResourceOwnerAccount(v string) *ModifySubConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySubConnectionRequest) SetResourceOwnerId(v int64) *ModifySubConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySubConnectionRequest) SetUisSubConnectionId(v string) *ModifySubConnectionRequest {
	s.UisSubConnectionId = &v
	return s
}

func (s *ModifySubConnectionRequest) SetDescription(v string) *ModifySubConnectionRequest {
	s.Description = &v
	return s
}

func (s *ModifySubConnectionRequest) SetName(v string) *ModifySubConnectionRequest {
	s.Name = &v
	return s
}

func (s *ModifySubConnectionRequest) SetGreConfig(v string) *ModifySubConnectionRequest {
	s.GreConfig = &v
	return s
}

type ModifySubConnectionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySubConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySubConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySubConnectionResponseBody) SetRequestId(v string) *ModifySubConnectionResponseBody {
	s.RequestId = &v
	return s
}

type ModifySubConnectionResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifySubConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifySubConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifySubConnectionResponse) GoString() string {
	return s.String()
}

func (s *ModifySubConnectionResponse) SetHeaders(v map[string]*string) *ModifySubConnectionResponse {
	s.Headers = v
	return s
}

func (s *ModifySubConnectionResponse) SetBody(v *ModifySubConnectionResponseBody) *ModifySubConnectionResponse {
	s.Body = v
	return s
}

type ModifyUisAttributeRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	UisId                *string `json:"UisId,omitempty" xml:"UisId,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyUisAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyUisAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyUisAttributeRequest) SetOwnerAccount(v string) *ModifyUisAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyUisAttributeRequest) SetOwnerId(v int64) *ModifyUisAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyUisAttributeRequest) SetResourceOwnerAccount(v string) *ModifyUisAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyUisAttributeRequest) SetResourceOwnerId(v int64) *ModifyUisAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyUisAttributeRequest) SetClientToken(v string) *ModifyUisAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyUisAttributeRequest) SetUisId(v string) *ModifyUisAttributeRequest {
	s.UisId = &v
	return s
}

func (s *ModifyUisAttributeRequest) SetName(v string) *ModifyUisAttributeRequest {
	s.Name = &v
	return s
}

func (s *ModifyUisAttributeRequest) SetDescription(v string) *ModifyUisAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyUisAttributeRequest) SetRegionId(v string) *ModifyUisAttributeRequest {
	s.RegionId = &v
	return s
}

type ModifyUisAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyUisAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyUisAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUisAttributeResponseBody) SetRequestId(v string) *ModifyUisAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyUisAttributeResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyUisAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyUisAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyUisAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyUisAttributeResponse) SetHeaders(v map[string]*string) *ModifyUisAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyUisAttributeResponse) SetBody(v *ModifyUisAttributeResponseBody) *ModifyUisAttributeResponse {
	s.Body = v
	return s
}

type ModifyUisConnectionAttributeRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UisNodeId            *string `json:"UisNodeId,omitempty" xml:"UisNodeId,omitempty"`
	UisConnectionId      *string `json:"UisConnectionId,omitempty" xml:"UisConnectionId,omitempty"`
	UisProtocol          *string `json:"UisProtocol,omitempty" xml:"UisProtocol,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	GreConfig            *string `json:"GreConfig,omitempty" xml:"GreConfig,omitempty"`
	SslConfig            *string `json:"SslConfig,omitempty" xml:"SslConfig,omitempty"`
}

func (s ModifyUisConnectionAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyUisConnectionAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyUisConnectionAttributeRequest) SetOwnerAccount(v string) *ModifyUisConnectionAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyUisConnectionAttributeRequest) SetOwnerId(v int64) *ModifyUisConnectionAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyUisConnectionAttributeRequest) SetResourceOwnerAccount(v string) *ModifyUisConnectionAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyUisConnectionAttributeRequest) SetResourceOwnerId(v int64) *ModifyUisConnectionAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyUisConnectionAttributeRequest) SetRegionId(v string) *ModifyUisConnectionAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyUisConnectionAttributeRequest) SetUisNodeId(v string) *ModifyUisConnectionAttributeRequest {
	s.UisNodeId = &v
	return s
}

func (s *ModifyUisConnectionAttributeRequest) SetUisConnectionId(v string) *ModifyUisConnectionAttributeRequest {
	s.UisConnectionId = &v
	return s
}

func (s *ModifyUisConnectionAttributeRequest) SetUisProtocol(v string) *ModifyUisConnectionAttributeRequest {
	s.UisProtocol = &v
	return s
}

func (s *ModifyUisConnectionAttributeRequest) SetDescription(v string) *ModifyUisConnectionAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyUisConnectionAttributeRequest) SetName(v string) *ModifyUisConnectionAttributeRequest {
	s.Name = &v
	return s
}

func (s *ModifyUisConnectionAttributeRequest) SetGreConfig(v string) *ModifyUisConnectionAttributeRequest {
	s.GreConfig = &v
	return s
}

func (s *ModifyUisConnectionAttributeRequest) SetSslConfig(v string) *ModifyUisConnectionAttributeRequest {
	s.SslConfig = &v
	return s
}

type ModifyUisConnectionAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyUisConnectionAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyUisConnectionAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUisConnectionAttributeResponseBody) SetRequestId(v string) *ModifyUisConnectionAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyUisConnectionAttributeResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyUisConnectionAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyUisConnectionAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyUisConnectionAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyUisConnectionAttributeResponse) SetHeaders(v map[string]*string) *ModifyUisConnectionAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyUisConnectionAttributeResponse) SetBody(v *ModifyUisConnectionAttributeResponseBody) *ModifyUisConnectionAttributeResponse {
	s.Body = v
	return s
}

type ModifyUisNodeAttributeRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UisId                *string `json:"UisId,omitempty" xml:"UisId,omitempty"`
	UisNodeId            *string `json:"UisNodeId,omitempty" xml:"UisNodeId,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	UisNodeBandwidth     *int32  `json:"UisNodeBandwidth,omitempty" xml:"UisNodeBandwidth,omitempty"`
}

func (s ModifyUisNodeAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyUisNodeAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyUisNodeAttributeRequest) SetOwnerAccount(v string) *ModifyUisNodeAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyUisNodeAttributeRequest) SetOwnerId(v int64) *ModifyUisNodeAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyUisNodeAttributeRequest) SetResourceOwnerAccount(v string) *ModifyUisNodeAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyUisNodeAttributeRequest) SetResourceOwnerId(v int64) *ModifyUisNodeAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyUisNodeAttributeRequest) SetRegionId(v string) *ModifyUisNodeAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyUisNodeAttributeRequest) SetUisId(v string) *ModifyUisNodeAttributeRequest {
	s.UisId = &v
	return s
}

func (s *ModifyUisNodeAttributeRequest) SetUisNodeId(v string) *ModifyUisNodeAttributeRequest {
	s.UisNodeId = &v
	return s
}

func (s *ModifyUisNodeAttributeRequest) SetName(v string) *ModifyUisNodeAttributeRequest {
	s.Name = &v
	return s
}

func (s *ModifyUisNodeAttributeRequest) SetDescription(v string) *ModifyUisNodeAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyUisNodeAttributeRequest) SetUisNodeBandwidth(v int32) *ModifyUisNodeAttributeRequest {
	s.UisNodeBandwidth = &v
	return s
}

type ModifyUisNodeAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyUisNodeAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyUisNodeAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUisNodeAttributeResponseBody) SetRequestId(v string) *ModifyUisNodeAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyUisNodeAttributeResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyUisNodeAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyUisNodeAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyUisNodeAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyUisNodeAttributeResponse) SetHeaders(v map[string]*string) *ModifyUisNodeAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyUisNodeAttributeResponse) SetBody(v *ModifyUisNodeAttributeResponseBody) *ModifyUisNodeAttributeResponse {
	s.Body = v
	return s
}

type ModifyWhiteListRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	UisId                *string `json:"UisId,omitempty" xml:"UisId,omitempty"`
	Whitelist            *string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty"`
	ModifyMode           *string `json:"ModifyMode,omitempty" xml:"ModifyMode,omitempty"`
}

func (s ModifyWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyWhiteListRequest) GoString() string {
	return s.String()
}

func (s *ModifyWhiteListRequest) SetOwnerAccount(v string) *ModifyWhiteListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyWhiteListRequest) SetOwnerId(v int64) *ModifyWhiteListRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyWhiteListRequest) SetResourceOwnerAccount(v string) *ModifyWhiteListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyWhiteListRequest) SetResourceOwnerId(v int64) *ModifyWhiteListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyWhiteListRequest) SetRegionId(v string) *ModifyWhiteListRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyWhiteListRequest) SetClientToken(v string) *ModifyWhiteListRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyWhiteListRequest) SetUisId(v string) *ModifyWhiteListRequest {
	s.UisId = &v
	return s
}

func (s *ModifyWhiteListRequest) SetWhitelist(v string) *ModifyWhiteListRequest {
	s.Whitelist = &v
	return s
}

func (s *ModifyWhiteListRequest) SetModifyMode(v string) *ModifyWhiteListRequest {
	s.ModifyMode = &v
	return s
}

type ModifyWhiteListResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWhiteListResponseBody) SetRequestId(v string) *ModifyWhiteListResponseBody {
	s.RequestId = &v
	return s
}

type ModifyWhiteListResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyWhiteListResponse) GoString() string {
	return s.String()
}

func (s *ModifyWhiteListResponse) SetHeaders(v map[string]*string) *ModifyWhiteListResponse {
	s.Headers = v
	return s
}

func (s *ModifyWhiteListResponse) SetBody(v *ModifyWhiteListResponseBody) *ModifyWhiteListResponse {
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
		"cn-north-2-gov-1":            tea.String("uis.cn-hangzhou.aliyuncs.com"),
		"ap-northeast-1":              tea.String("uis.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("uis.aliyuncs.com"),
		"ap-south-1":                  tea.String("uis.aliyuncs.com"),
		"ap-southeast-1":              tea.String("uis.aliyuncs.com"),
		"ap-southeast-2":              tea.String("uis.aliyuncs.com"),
		"ap-southeast-3":              tea.String("uis.aliyuncs.com"),
		"ap-southeast-5":              tea.String("uis.aliyuncs.com"),
		"cn-beijing":                  tea.String("uis.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("uis.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("uis.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("uis.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("uis.aliyuncs.com"),
		"cn-chengdu":                  tea.String("uis.aliyuncs.com"),
		"cn-edge-1":                   tea.String("uis.aliyuncs.com"),
		"cn-fujian":                   tea.String("uis.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("uis.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("uis.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("uis.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("uis.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("uis.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("uis.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("uis.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("uis.aliyuncs.com"),
		"cn-hongkong":                 tea.String("uis.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("uis.aliyuncs.com"),
		"cn-huhehaote":                tea.String("uis.aliyuncs.com"),
		"cn-qingdao":                  tea.String("uis.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("uis.aliyuncs.com"),
		"cn-shanghai":                 tea.String("uis.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("uis.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("uis.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("uis.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("uis.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("uis.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("uis.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("uis.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("uis.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("uis.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("uis.aliyuncs.com"),
		"cn-wuhan":                    tea.String("uis.aliyuncs.com"),
		"cn-yushanfang":               tea.String("uis.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("uis.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("uis.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("uis.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("uis.aliyuncs.com"),
		"eu-central-1":                tea.String("uis.aliyuncs.com"),
		"eu-west-1":                   tea.String("uis.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("uis.aliyuncs.com"),
		"me-east-1":                   tea.String("uis.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("uis.aliyuncs.com"),
		"us-east-1":                   tea.String("uis.aliyuncs.com"),
		"us-west-1":                   tea.String("uis.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("uis"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddHighPriorityIpWithOptions(request *AddHighPriorityIpRequest, runtime *util.RuntimeOptions) (_result *AddHighPriorityIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddHighPriorityIpResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddHighPriorityIp"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddHighPriorityIp(request *AddHighPriorityIpRequest) (_result *AddHighPriorityIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddHighPriorityIpResponse{}
	_body, _err := client.AddHighPriorityIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddUisNodeIpWithOptions(request *AddUisNodeIpRequest, runtime *util.RuntimeOptions) (_result *AddUisNodeIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddUisNodeIpResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddUisNodeIp"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddUisNodeIp(request *AddUisNodeIpRequest) (_result *AddUisNodeIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddUisNodeIpResponse{}
	_body, _err := client.AddUisNodeIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDnatEntryWithOptions(request *CreateDnatEntryRequest, runtime *util.RuntimeOptions) (_result *CreateDnatEntryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDnatEntryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDnatEntry"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDnatEntry(request *CreateDnatEntryRequest) (_result *CreateDnatEntryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDnatEntryResponse{}
	_body, _err := client.CreateDnatEntryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSubConnectionWithOptions(request *CreateSubConnectionRequest, runtime *util.RuntimeOptions) (_result *CreateSubConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateSubConnectionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateSubConnection"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSubConnection(request *CreateSubConnectionRequest) (_result *CreateSubConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSubConnectionResponse{}
	_body, _err := client.CreateSubConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateUisWithOptions(request *CreateUisRequest, runtime *util.RuntimeOptions) (_result *CreateUisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateUisResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateUis"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateUis(request *CreateUisRequest) (_result *CreateUisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUisResponse{}
	_body, _err := client.CreateUisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateUisConnectionWithOptions(request *CreateUisConnectionRequest, runtime *util.RuntimeOptions) (_result *CreateUisConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateUisConnectionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateUisConnection"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateUisConnection(request *CreateUisConnectionRequest) (_result *CreateUisConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUisConnectionResponse{}
	_body, _err := client.CreateUisConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateUisNetworkInterfaceWithOptions(request *CreateUisNetworkInterfaceRequest, runtime *util.RuntimeOptions) (_result *CreateUisNetworkInterfaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateUisNetworkInterfaceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateUisNetworkInterface"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateUisNetworkInterface(request *CreateUisNetworkInterfaceRequest) (_result *CreateUisNetworkInterfaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUisNetworkInterfaceResponse{}
	_body, _err := client.CreateUisNetworkInterfaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateUisNodeWithOptions(request *CreateUisNodeRequest, runtime *util.RuntimeOptions) (_result *CreateUisNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateUisNodeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateUisNode"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateUisNode(request *CreateUisNodeRequest) (_result *CreateUisNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUisNodeResponse{}
	_body, _err := client.CreateUisNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDnatEntryWithOptions(request *DeleteDnatEntryRequest, runtime *util.RuntimeOptions) (_result *DeleteDnatEntryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDnatEntryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDnatEntry"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDnatEntry(request *DeleteDnatEntryRequest) (_result *DeleteDnatEntryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDnatEntryResponse{}
	_body, _err := client.DeleteDnatEntryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteHighPriorityIpWithOptions(request *DeleteHighPriorityIpRequest, runtime *util.RuntimeOptions) (_result *DeleteHighPriorityIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteHighPriorityIpResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteHighPriorityIp"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteHighPriorityIp(request *DeleteHighPriorityIpRequest) (_result *DeleteHighPriorityIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteHighPriorityIpResponse{}
	_body, _err := client.DeleteHighPriorityIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSubConnectionWithOptions(request *DeleteSubConnectionRequest, runtime *util.RuntimeOptions) (_result *DeleteSubConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteSubConnectionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteSubConnection"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSubConnection(request *DeleteSubConnectionRequest) (_result *DeleteSubConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSubConnectionResponse{}
	_body, _err := client.DeleteSubConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUisWithOptions(request *DeleteUisRequest, runtime *util.RuntimeOptions) (_result *DeleteUisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteUisResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteUis"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUis(request *DeleteUisRequest) (_result *DeleteUisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUisResponse{}
	_body, _err := client.DeleteUisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUisConnectionWithOptions(request *DeleteUisConnectionRequest, runtime *util.RuntimeOptions) (_result *DeleteUisConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteUisConnectionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteUisConnection"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUisConnection(request *DeleteUisConnectionRequest) (_result *DeleteUisConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUisConnectionResponse{}
	_body, _err := client.DeleteUisConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUisNetworkInterfaceWithOptions(request *DeleteUisNetworkInterfaceRequest, runtime *util.RuntimeOptions) (_result *DeleteUisNetworkInterfaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteUisNetworkInterfaceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteUisNetworkInterface"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUisNetworkInterface(request *DeleteUisNetworkInterfaceRequest) (_result *DeleteUisNetworkInterfaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUisNetworkInterfaceResponse{}
	_body, _err := client.DeleteUisNetworkInterfaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUisNodeWithOptions(request *DeleteUisNodeRequest, runtime *util.RuntimeOptions) (_result *DeleteUisNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteUisNodeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteUisNode"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUisNode(request *DeleteUisNodeRequest) (_result *DeleteUisNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUisNodeResponse{}
	_body, _err := client.DeleteUisNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUisNodeIpWithOptions(request *DeleteUisNodeIpRequest, runtime *util.RuntimeOptions) (_result *DeleteUisNodeIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteUisNodeIpResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteUisNodeIp"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUisNodeIp(request *DeleteUisNodeIpRequest) (_result *DeleteUisNodeIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUisNodeIpResponse{}
	_body, _err := client.DeleteUisNodeIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAreasWithOptions(request *DescribeAreasRequest, runtime *util.RuntimeOptions) (_result *DescribeAreasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAreasResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAreas"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAreas(request *DescribeAreasRequest) (_result *DescribeAreasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAreasResponse{}
	_body, _err := client.DescribeAreasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDnatEntriesWithOptions(request *DescribeDnatEntriesRequest, runtime *util.RuntimeOptions) (_result *DescribeDnatEntriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDnatEntriesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDnatEntries"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDnatEntries(request *DescribeDnatEntriesRequest) (_result *DescribeDnatEntriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDnatEntriesResponse{}
	_body, _err := client.DescribeDnatEntriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHighPriorityIpWithOptions(request *DescribeHighPriorityIpRequest, runtime *util.RuntimeOptions) (_result *DescribeHighPriorityIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeHighPriorityIpResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeHighPriorityIp"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHighPriorityIp(request *DescribeHighPriorityIpRequest) (_result *DescribeHighPriorityIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHighPriorityIpResponse{}
	_body, _err := client.DescribeHighPriorityIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHighPriorityIpsWithOptions(request *DescribeHighPriorityIpsRequest, runtime *util.RuntimeOptions) (_result *DescribeHighPriorityIpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeHighPriorityIpsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeHighPriorityIps"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHighPriorityIps(request *DescribeHighPriorityIpsRequest) (_result *DescribeHighPriorityIpsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHighPriorityIpsResponse{}
	_body, _err := client.DescribeHighPriorityIpsWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DescribeRegions"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeSubConnectionWithOptions(request *DescribeSubConnectionRequest, runtime *util.RuntimeOptions) (_result *DescribeSubConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSubConnectionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSubConnection"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSubConnection(request *DescribeSubConnectionRequest) (_result *DescribeSubConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSubConnectionResponse{}
	_body, _err := client.DescribeSubConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSubConnectionsWithOptions(request *DescribeSubConnectionsRequest, runtime *util.RuntimeOptions) (_result *DescribeSubConnectionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSubConnectionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSubConnections"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSubConnections(request *DescribeSubConnectionsRequest) (_result *DescribeSubConnectionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSubConnectionsResponse{}
	_body, _err := client.DescribeSubConnectionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUisConnectionsWithOptions(request *DescribeUisConnectionsRequest, runtime *util.RuntimeOptions) (_result *DescribeUisConnectionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeUisConnectionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeUisConnections"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUisConnections(request *DescribeUisConnectionsRequest) (_result *DescribeUisConnectionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUisConnectionsResponse{}
	_body, _err := client.DescribeUisConnectionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUiseNodeStatusWithOptions(request *DescribeUiseNodeStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeUiseNodeStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeUiseNodeStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeUiseNodeStatus"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUiseNodeStatus(request *DescribeUiseNodeStatusRequest) (_result *DescribeUiseNodeStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUiseNodeStatusResponse{}
	_body, _err := client.DescribeUiseNodeStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUisesWithOptions(request *DescribeUisesRequest, runtime *util.RuntimeOptions) (_result *DescribeUisesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeUisesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeUises"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUises(request *DescribeUisesRequest) (_result *DescribeUisesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUisesResponse{}
	_body, _err := client.DescribeUisesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUisNetworkInterfacesWithOptions(request *DescribeUisNetworkInterfacesRequest, runtime *util.RuntimeOptions) (_result *DescribeUisNetworkInterfacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeUisNetworkInterfacesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeUisNetworkInterfaces"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUisNetworkInterfaces(request *DescribeUisNetworkInterfacesRequest) (_result *DescribeUisNetworkInterfacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUisNetworkInterfacesResponse{}
	_body, _err := client.DescribeUisNetworkInterfacesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUisNodesWithOptions(request *DescribeUisNodesRequest, runtime *util.RuntimeOptions) (_result *DescribeUisNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeUisNodesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeUisNodes"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUisNodes(request *DescribeUisNodesRequest) (_result *DescribeUisNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUisNodesResponse{}
	_body, _err := client.DescribeUisNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeWhiteListWithOptions(request *DescribeWhiteListRequest, runtime *util.RuntimeOptions) (_result *DescribeWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeWhiteListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeWhiteList"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeWhiteList(request *DescribeWhiteListRequest) (_result *DescribeWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWhiteListResponse{}
	_body, _err := client.DescribeWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDroppedIpListWithOptions(request *GetDroppedIpListRequest, runtime *util.RuntimeOptions) (_result *GetDroppedIpListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDroppedIpListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDroppedIpList"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDroppedIpList(request *GetDroppedIpListRequest) (_result *GetDroppedIpListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDroppedIpListResponse{}
	_body, _err := client.GetDroppedIpListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDnatEntryWithOptions(request *ModifyDnatEntryRequest, runtime *util.RuntimeOptions) (_result *ModifyDnatEntryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDnatEntryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDnatEntry"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDnatEntry(request *ModifyDnatEntryRequest) (_result *ModifyDnatEntryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDnatEntryResponse{}
	_body, _err := client.ModifyDnatEntryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyHighPriorityIpWithOptions(request *ModifyHighPriorityIpRequest, runtime *util.RuntimeOptions) (_result *ModifyHighPriorityIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyHighPriorityIpResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyHighPriorityIp"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyHighPriorityIp(request *ModifyHighPriorityIpRequest) (_result *ModifyHighPriorityIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyHighPriorityIpResponse{}
	_body, _err := client.ModifyHighPriorityIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifySubConnectionWithOptions(request *ModifySubConnectionRequest, runtime *util.RuntimeOptions) (_result *ModifySubConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifySubConnectionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifySubConnection"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifySubConnection(request *ModifySubConnectionRequest) (_result *ModifySubConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifySubConnectionResponse{}
	_body, _err := client.ModifySubConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyUisAttributeWithOptions(request *ModifyUisAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyUisAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyUisAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyUisAttribute"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyUisAttribute(request *ModifyUisAttributeRequest) (_result *ModifyUisAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyUisAttributeResponse{}
	_body, _err := client.ModifyUisAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyUisConnectionAttributeWithOptions(request *ModifyUisConnectionAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyUisConnectionAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyUisConnectionAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyUisConnectionAttribute"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyUisConnectionAttribute(request *ModifyUisConnectionAttributeRequest) (_result *ModifyUisConnectionAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyUisConnectionAttributeResponse{}
	_body, _err := client.ModifyUisConnectionAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyUisNodeAttributeWithOptions(request *ModifyUisNodeAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyUisNodeAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyUisNodeAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyUisNodeAttribute"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyUisNodeAttribute(request *ModifyUisNodeAttributeRequest) (_result *ModifyUisNodeAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyUisNodeAttributeResponse{}
	_body, _err := client.ModifyUisNodeAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyWhiteListWithOptions(request *ModifyWhiteListRequest, runtime *util.RuntimeOptions) (_result *ModifyWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyWhiteListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyWhiteList"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyWhiteList(request *ModifyWhiteListRequest) (_result *ModifyWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyWhiteListResponse{}
	_body, _err := client.ModifyWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
