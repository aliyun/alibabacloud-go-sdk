// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DescribeUisRequest struct {
	UisId *string `json:"UisId,omitempty" xml:"UisId,omitempty" require:"true"`
}

func (s DescribeUisRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUisRequest) GoString() string {
	return s.String()
}

func (s *DescribeUisRequest) SetUisId(v string) *DescribeUisRequest {
	s.UisId = &v
	return s
}

type DescribeUisResponse struct {
	RequestId   *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	UisId       *string                           `json:"UisId,omitempty" xml:"UisId,omitempty" require:"true"`
	UisName     *string                           `json:"UisName,omitempty" xml:"UisName,omitempty" require:"true"`
	State       *string                           `json:"State,omitempty" xml:"State,omitempty" require:"true"`
	RegionId    *string                           `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	InternetIps []*DescribeUisResponseInternetIps `json:"InternetIps,omitempty" xml:"InternetIps,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeUisResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUisResponse) GoString() string {
	return s.String()
}

func (s *DescribeUisResponse) SetRequestId(v string) *DescribeUisResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeUisResponse) SetUisId(v string) *DescribeUisResponse {
	s.UisId = &v
	return s
}

func (s *DescribeUisResponse) SetUisName(v string) *DescribeUisResponse {
	s.UisName = &v
	return s
}

func (s *DescribeUisResponse) SetState(v string) *DescribeUisResponse {
	s.State = &v
	return s
}

func (s *DescribeUisResponse) SetRegionId(v string) *DescribeUisResponse {
	s.RegionId = &v
	return s
}

func (s *DescribeUisResponse) SetInternetIps(v []*DescribeUisResponseInternetIps) *DescribeUisResponse {
	s.InternetIps = v
	return s
}

type DescribeUisResponseInternetIps struct {
	Ip    *string `json:"Ip,omitempty" xml:"Ip,omitempty" require:"true"`
	Role  *string `json:"Role,omitempty" xml:"Role,omitempty" require:"true"`
	EipId *string `json:"EipId,omitempty" xml:"EipId,omitempty" require:"true"`
}

func (s DescribeUisResponseInternetIps) String() string {
	return tea.Prettify(s)
}

func (s DescribeUisResponseInternetIps) GoString() string {
	return s.String()
}

func (s *DescribeUisResponseInternetIps) SetIp(v string) *DescribeUisResponseInternetIps {
	s.Ip = &v
	return s
}

func (s *DescribeUisResponseInternetIps) SetRole(v string) *DescribeUisResponseInternetIps {
	s.Role = &v
	return s
}

func (s *DescribeUisResponseInternetIps) SetEipId(v string) *DescribeUisResponseInternetIps {
	s.EipId = &v
	return s
}

type ModifyVnetRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	UisId       *string `json:"UisId,omitempty" xml:"UisId,omitempty" require:"true"`
	VnetId      *string `json:"VnetId,omitempty" xml:"VnetId,omitempty" require:"true"`
	Cidrs       *string `json:"Cidrs,omitempty" xml:"Cidrs,omitempty"`
	MbpsQuota   *int    `json:"MbpsQuota,omitempty" xml:"MbpsQuota,omitempty"`
	KppsQuota   *int    `json:"KppsQuota,omitempty" xml:"KppsQuota,omitempty"`
	FlowQuota   *int    `json:"FlowQuota,omitempty" xml:"FlowQuota,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ModifyVnetRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyVnetRequest) GoString() string {
	return s.String()
}

func (s *ModifyVnetRequest) SetClientToken(v string) *ModifyVnetRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyVnetRequest) SetRegionId(v string) *ModifyVnetRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyVnetRequest) SetUisId(v string) *ModifyVnetRequest {
	s.UisId = &v
	return s
}

func (s *ModifyVnetRequest) SetVnetId(v string) *ModifyVnetRequest {
	s.VnetId = &v
	return s
}

func (s *ModifyVnetRequest) SetCidrs(v string) *ModifyVnetRequest {
	s.Cidrs = &v
	return s
}

func (s *ModifyVnetRequest) SetMbpsQuota(v int) *ModifyVnetRequest {
	s.MbpsQuota = &v
	return s
}

func (s *ModifyVnetRequest) SetKppsQuota(v int) *ModifyVnetRequest {
	s.KppsQuota = &v
	return s
}

func (s *ModifyVnetRequest) SetFlowQuota(v int) *ModifyVnetRequest {
	s.FlowQuota = &v
	return s
}

func (s *ModifyVnetRequest) SetName(v string) *ModifyVnetRequest {
	s.Name = &v
	return s
}

type ModifyVnetResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ModifyVnetResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyVnetResponse) GoString() string {
	return s.String()
}

func (s *ModifyVnetResponse) SetRequestId(v string) *ModifyVnetResponse {
	s.RequestId = &v
	return s
}

type DescribeGreConnectionsRequest struct {
	UisId      *string `json:"UisId,omitempty" xml:"UisId,omitempty" require:"true"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	PageNumber *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeGreConnectionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGreConnectionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeGreConnectionsRequest) SetUisId(v string) *DescribeGreConnectionsRequest {
	s.UisId = &v
	return s
}

func (s *DescribeGreConnectionsRequest) SetRegionId(v string) *DescribeGreConnectionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeGreConnectionsRequest) SetPageNumber(v int) *DescribeGreConnectionsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeGreConnectionsRequest) SetPageSize(v int) *DescribeGreConnectionsRequest {
	s.PageSize = &v
	return s
}

type DescribeGreConnectionsResponse struct {
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	UisId          *string                                         `json:"UisId,omitempty" xml:"UisId,omitempty" require:"true"`
	UisName        *string                                         `json:"UisName,omitempty" xml:"UisName,omitempty" require:"true"`
	State          *string                                         `json:"State,omitempty" xml:"State,omitempty" require:"true"`
	Page           *int64                                          `json:"Page,omitempty" xml:"Page,omitempty" require:"true"`
	PageSize       *int64                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	TotalCount     *int                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	GreConnections []*DescribeGreConnectionsResponseGreConnections `json:"GreConnections,omitempty" xml:"GreConnections,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeGreConnectionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGreConnectionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeGreConnectionsResponse) SetRequestId(v string) *DescribeGreConnectionsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeGreConnectionsResponse) SetUisId(v string) *DescribeGreConnectionsResponse {
	s.UisId = &v
	return s
}

func (s *DescribeGreConnectionsResponse) SetUisName(v string) *DescribeGreConnectionsResponse {
	s.UisName = &v
	return s
}

func (s *DescribeGreConnectionsResponse) SetState(v string) *DescribeGreConnectionsResponse {
	s.State = &v
	return s
}

func (s *DescribeGreConnectionsResponse) SetPage(v int64) *DescribeGreConnectionsResponse {
	s.Page = &v
	return s
}

func (s *DescribeGreConnectionsResponse) SetPageSize(v int64) *DescribeGreConnectionsResponse {
	s.PageSize = &v
	return s
}

func (s *DescribeGreConnectionsResponse) SetTotalCount(v int) *DescribeGreConnectionsResponse {
	s.TotalCount = &v
	return s
}

func (s *DescribeGreConnectionsResponse) SetGreConnections(v []*DescribeGreConnectionsResponseGreConnections) *DescribeGreConnectionsResponse {
	s.GreConnections = v
	return s
}

type DescribeGreConnectionsResponseGreConnections struct {
	GreConnectionId  *string `json:"GreConnectionId,omitempty" xml:"GreConnectionId,omitempty" require:"true"`
	State            *string `json:"State,omitempty" xml:"State,omitempty" require:"true"`
	Cidrs            *string `json:"Cidrs,omitempty" xml:"Cidrs,omitempty" require:"true"`
	LocalIp          *string `json:"LocalIp,omitempty" xml:"LocalIp,omitempty" require:"true"`
	LocalTunnelIp    *string `json:"LocalTunnelIp,omitempty" xml:"LocalTunnelIp,omitempty" require:"true"`
	CustomerIp       *string `json:"CustomerIp,omitempty" xml:"CustomerIp,omitempty" require:"true"`
	CustomerTunnelIp *string `json:"CustomerTunnelIp,omitempty" xml:"CustomerTunnelIp,omitempty" require:"true"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	CreateTime       *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
}

func (s DescribeGreConnectionsResponseGreConnections) String() string {
	return tea.Prettify(s)
}

func (s DescribeGreConnectionsResponseGreConnections) GoString() string {
	return s.String()
}

func (s *DescribeGreConnectionsResponseGreConnections) SetGreConnectionId(v string) *DescribeGreConnectionsResponseGreConnections {
	s.GreConnectionId = &v
	return s
}

func (s *DescribeGreConnectionsResponseGreConnections) SetState(v string) *DescribeGreConnectionsResponseGreConnections {
	s.State = &v
	return s
}

func (s *DescribeGreConnectionsResponseGreConnections) SetCidrs(v string) *DescribeGreConnectionsResponseGreConnections {
	s.Cidrs = &v
	return s
}

func (s *DescribeGreConnectionsResponseGreConnections) SetLocalIp(v string) *DescribeGreConnectionsResponseGreConnections {
	s.LocalIp = &v
	return s
}

func (s *DescribeGreConnectionsResponseGreConnections) SetLocalTunnelIp(v string) *DescribeGreConnectionsResponseGreConnections {
	s.LocalTunnelIp = &v
	return s
}

func (s *DescribeGreConnectionsResponseGreConnections) SetCustomerIp(v string) *DescribeGreConnectionsResponseGreConnections {
	s.CustomerIp = &v
	return s
}

func (s *DescribeGreConnectionsResponseGreConnections) SetCustomerTunnelIp(v string) *DescribeGreConnectionsResponseGreConnections {
	s.CustomerTunnelIp = &v
	return s
}

func (s *DescribeGreConnectionsResponseGreConnections) SetName(v string) *DescribeGreConnectionsResponseGreConnections {
	s.Name = &v
	return s
}

func (s *DescribeGreConnectionsResponseGreConnections) SetDescription(v string) *DescribeGreConnectionsResponseGreConnections {
	s.Description = &v
	return s
}

func (s *DescribeGreConnectionsResponseGreConnections) SetCreateTime(v int64) *DescribeGreConnectionsResponseGreConnections {
	s.CreateTime = &v
	return s
}

type DescribeUissRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	PageNumber *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeUissRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUissRequest) GoString() string {
	return s.String()
}

func (s *DescribeUissRequest) SetRegionId(v string) *DescribeUissRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUissRequest) SetPageNumber(v int) *DescribeUissRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeUissRequest) SetPageSize(v int) *DescribeUissRequest {
	s.PageSize = &v
	return s
}

type DescribeUissResponse struct {
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Page      *int64                      `json:"Page,omitempty" xml:"Page,omitempty" require:"true"`
	PageSize  *int64                      `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	Uiss      []*DescribeUissResponseUiss `json:"Uiss,omitempty" xml:"Uiss,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeUissResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUissResponse) GoString() string {
	return s.String()
}

func (s *DescribeUissResponse) SetRequestId(v string) *DescribeUissResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeUissResponse) SetPage(v int64) *DescribeUissResponse {
	s.Page = &v
	return s
}

func (s *DescribeUissResponse) SetPageSize(v int64) *DescribeUissResponse {
	s.PageSize = &v
	return s
}

func (s *DescribeUissResponse) SetUiss(v []*DescribeUissResponseUiss) *DescribeUissResponse {
	s.Uiss = v
	return s
}

type DescribeUissResponseUiss struct {
	UisId           *string `json:"UisId,omitempty" xml:"UisId,omitempty" require:"true"`
	UisName         *string `json:"UisName,omitempty" xml:"UisName,omitempty" require:"true"`
	Spec            *string `json:"Spec,omitempty" xml:"Spec,omitempty" require:"true"`
	State           *string `json:"State,omitempty" xml:"State,omitempty" require:"true"`
	CreateTime      *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty" require:"true"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
}

func (s DescribeUissResponseUiss) String() string {
	return tea.Prettify(s)
}

func (s DescribeUissResponseUiss) GoString() string {
	return s.String()
}

func (s *DescribeUissResponseUiss) SetUisId(v string) *DescribeUissResponseUiss {
	s.UisId = &v
	return s
}

func (s *DescribeUissResponseUiss) SetUisName(v string) *DescribeUissResponseUiss {
	s.UisName = &v
	return s
}

func (s *DescribeUissResponseUiss) SetSpec(v string) *DescribeUissResponseUiss {
	s.Spec = &v
	return s
}

func (s *DescribeUissResponseUiss) SetState(v string) *DescribeUissResponseUiss {
	s.State = &v
	return s
}

func (s *DescribeUissResponseUiss) SetCreateTime(v int64) *DescribeUissResponseUiss {
	s.CreateTime = &v
	return s
}

func (s *DescribeUissResponseUiss) SetDescription(v string) *DescribeUissResponseUiss {
	s.Description = &v
	return s
}

func (s *DescribeUissResponseUiss) SetResourceGroupId(v string) *DescribeUissResponseUiss {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeUissResponseUiss) SetRegionId(v string) *DescribeUissResponseUiss {
	s.RegionId = &v
	return s
}

type DescribeVnetRouteEntryListRequest struct {
	UisId      *string `json:"UisId,omitempty" xml:"UisId,omitempty" require:"true"`
	VnetId     *string `json:"VnetId,omitempty" xml:"VnetId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	PageNumber *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeVnetRouteEntryListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVnetRouteEntryListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVnetRouteEntryListRequest) SetUisId(v string) *DescribeVnetRouteEntryListRequest {
	s.UisId = &v
	return s
}

func (s *DescribeVnetRouteEntryListRequest) SetVnetId(v string) *DescribeVnetRouteEntryListRequest {
	s.VnetId = &v
	return s
}

func (s *DescribeVnetRouteEntryListRequest) SetRegionId(v string) *DescribeVnetRouteEntryListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVnetRouteEntryListRequest) SetPageNumber(v int) *DescribeVnetRouteEntryListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVnetRouteEntryListRequest) SetPageSize(v int) *DescribeVnetRouteEntryListRequest {
	s.PageSize = &v
	return s
}

type DescribeVnetRouteEntryListResponse struct {
	RequestId  *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PageSize   *int                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	PageNumber *int                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	TotalCount *int                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	Vnets      []*DescribeVnetRouteEntryListResponseVnets `json:"Vnets,omitempty" xml:"Vnets,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeVnetRouteEntryListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVnetRouteEntryListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVnetRouteEntryListResponse) SetRequestId(v string) *DescribeVnetRouteEntryListResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeVnetRouteEntryListResponse) SetPageSize(v int) *DescribeVnetRouteEntryListResponse {
	s.PageSize = &v
	return s
}

func (s *DescribeVnetRouteEntryListResponse) SetPageNumber(v int) *DescribeVnetRouteEntryListResponse {
	s.PageNumber = &v
	return s
}

func (s *DescribeVnetRouteEntryListResponse) SetTotalCount(v int) *DescribeVnetRouteEntryListResponse {
	s.TotalCount = &v
	return s
}

func (s *DescribeVnetRouteEntryListResponse) SetVnets(v []*DescribeVnetRouteEntryListResponseVnets) *DescribeVnetRouteEntryListResponse {
	s.Vnets = v
	return s
}

type DescribeVnetRouteEntryListResponseVnets struct {
	UisId       *string `json:"UisId,omitempty" xml:"UisId,omitempty" require:"true"`
	VnetId      *string `json:"VnetId,omitempty" xml:"VnetId,omitempty" require:"true"`
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty" require:"true"`
	NexthopType *string `json:"NexthopType,omitempty" xml:"NexthopType,omitempty" require:"true"`
	Nexthop     *string `json:"Nexthop,omitempty" xml:"Nexthop,omitempty" require:"true"`
}

func (s DescribeVnetRouteEntryListResponseVnets) String() string {
	return tea.Prettify(s)
}

func (s DescribeVnetRouteEntryListResponseVnets) GoString() string {
	return s.String()
}

func (s *DescribeVnetRouteEntryListResponseVnets) SetUisId(v string) *DescribeVnetRouteEntryListResponseVnets {
	s.UisId = &v
	return s
}

func (s *DescribeVnetRouteEntryListResponseVnets) SetVnetId(v string) *DescribeVnetRouteEntryListResponseVnets {
	s.VnetId = &v
	return s
}

func (s *DescribeVnetRouteEntryListResponseVnets) SetDestination(v string) *DescribeVnetRouteEntryListResponseVnets {
	s.Destination = &v
	return s
}

func (s *DescribeVnetRouteEntryListResponseVnets) SetNexthopType(v string) *DescribeVnetRouteEntryListResponseVnets {
	s.NexthopType = &v
	return s
}

func (s *DescribeVnetRouteEntryListResponseVnets) SetNexthop(v string) *DescribeVnetRouteEntryListResponseVnets {
	s.Nexthop = &v
	return s
}

type CreateVnetRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	UisId       *string `json:"UisId,omitempty" xml:"UisId,omitempty" require:"true"`
	Cidrs       *string `json:"Cidrs,omitempty" xml:"Cidrs,omitempty" require:"true"`
	MbpsQuota   *int    `json:"MbpsQuota,omitempty" xml:"MbpsQuota,omitempty"`
	KppsQuota   *int    `json:"KppsQuota,omitempty" xml:"KppsQuota,omitempty"`
	FlowQuota   *int    `json:"FlowQuota,omitempty" xml:"FlowQuota,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateVnetRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVnetRequest) GoString() string {
	return s.String()
}

func (s *CreateVnetRequest) SetClientToken(v string) *CreateVnetRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateVnetRequest) SetRegionId(v string) *CreateVnetRequest {
	s.RegionId = &v
	return s
}

func (s *CreateVnetRequest) SetUisId(v string) *CreateVnetRequest {
	s.UisId = &v
	return s
}

func (s *CreateVnetRequest) SetCidrs(v string) *CreateVnetRequest {
	s.Cidrs = &v
	return s
}

func (s *CreateVnetRequest) SetMbpsQuota(v int) *CreateVnetRequest {
	s.MbpsQuota = &v
	return s
}

func (s *CreateVnetRequest) SetKppsQuota(v int) *CreateVnetRequest {
	s.KppsQuota = &v
	return s
}

func (s *CreateVnetRequest) SetFlowQuota(v int) *CreateVnetRequest {
	s.FlowQuota = &v
	return s
}

func (s *CreateVnetRequest) SetName(v string) *CreateVnetRequest {
	s.Name = &v
	return s
}

type CreateVnetResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	VnetId    *string `json:"VnetId,omitempty" xml:"VnetId,omitempty" require:"true"`
}

func (s CreateVnetResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVnetResponse) GoString() string {
	return s.String()
}

func (s *CreateVnetResponse) SetRequestId(v string) *CreateVnetResponse {
	s.RequestId = &v
	return s
}

func (s *CreateVnetResponse) SetVnetId(v string) *CreateVnetResponse {
	s.VnetId = &v
	return s
}

type DeleteVnetRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	UisId       *string `json:"UisId,omitempty" xml:"UisId,omitempty" require:"true"`
	VnetId      *string `json:"VnetId,omitempty" xml:"VnetId,omitempty" require:"true"`
}

func (s DeleteVnetRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVnetRequest) GoString() string {
	return s.String()
}

func (s *DeleteVnetRequest) SetClientToken(v string) *DeleteVnetRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteVnetRequest) SetRegionId(v string) *DeleteVnetRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteVnetRequest) SetUisId(v string) *DeleteVnetRequest {
	s.UisId = &v
	return s
}

func (s *DeleteVnetRequest) SetVnetId(v string) *DeleteVnetRequest {
	s.VnetId = &v
	return s
}

type DeleteVnetResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteVnetResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVnetResponse) GoString() string {
	return s.String()
}

func (s *DeleteVnetResponse) SetRequestId(v string) *DeleteVnetResponse {
	s.RequestId = &v
	return s
}

type DescribeVnetsRequest struct {
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	PageNumber *int      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	VnetId     []*string `json:"VnetId,omitempty" xml:"VnetId,omitempty" type:"Repeated"`
	UisId      *string   `json:"UisId,omitempty" xml:"UisId,omitempty" require:"true"`
}

func (s DescribeVnetsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVnetsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVnetsRequest) SetRegionId(v string) *DescribeVnetsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVnetsRequest) SetPageNumber(v int) *DescribeVnetsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVnetsRequest) SetPageSize(v int) *DescribeVnetsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVnetsRequest) SetVnetId(v []*string) *DescribeVnetsRequest {
	s.VnetId = v
	return s
}

func (s *DescribeVnetsRequest) SetUisId(v string) *DescribeVnetsRequest {
	s.UisId = &v
	return s
}

type DescribeVnetsResponse struct {
	RequestId  *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PageSize   *int                          `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	PageNumber *int                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	TotalCount *int                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	Vnets      []*DescribeVnetsResponseVnets `json:"Vnets,omitempty" xml:"Vnets,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeVnetsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVnetsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVnetsResponse) SetRequestId(v string) *DescribeVnetsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeVnetsResponse) SetPageSize(v int) *DescribeVnetsResponse {
	s.PageSize = &v
	return s
}

func (s *DescribeVnetsResponse) SetPageNumber(v int) *DescribeVnetsResponse {
	s.PageNumber = &v
	return s
}

func (s *DescribeVnetsResponse) SetTotalCount(v int) *DescribeVnetsResponse {
	s.TotalCount = &v
	return s
}

func (s *DescribeVnetsResponse) SetVnets(v []*DescribeVnetsResponseVnets) *DescribeVnetsResponse {
	s.Vnets = v
	return s
}

type DescribeVnetsResponseVnets struct {
	UisId     *string `json:"UisId,omitempty" xml:"UisId,omitempty" require:"true"`
	VnetId    *string `json:"VnetId,omitempty" xml:"VnetId,omitempty" require:"true"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Cidrs     *string `json:"Cidrs,omitempty" xml:"Cidrs,omitempty" require:"true"`
	MbpsQuota *int    `json:"MbpsQuota,omitempty" xml:"MbpsQuota,omitempty" require:"true"`
	KppsQuota *int    `json:"KppsQuota,omitempty" xml:"KppsQuota,omitempty" require:"true"`
	FlowQuota *int    `json:"FlowQuota,omitempty" xml:"FlowQuota,omitempty" require:"true"`
	State     *string `json:"State,omitempty" xml:"State,omitempty" require:"true"`
}

func (s DescribeVnetsResponseVnets) String() string {
	return tea.Prettify(s)
}

func (s DescribeVnetsResponseVnets) GoString() string {
	return s.String()
}

func (s *DescribeVnetsResponseVnets) SetUisId(v string) *DescribeVnetsResponseVnets {
	s.UisId = &v
	return s
}

func (s *DescribeVnetsResponseVnets) SetVnetId(v string) *DescribeVnetsResponseVnets {
	s.VnetId = &v
	return s
}

func (s *DescribeVnetsResponseVnets) SetName(v string) *DescribeVnetsResponseVnets {
	s.Name = &v
	return s
}

func (s *DescribeVnetsResponseVnets) SetCidrs(v string) *DescribeVnetsResponseVnets {
	s.Cidrs = &v
	return s
}

func (s *DescribeVnetsResponseVnets) SetMbpsQuota(v int) *DescribeVnetsResponseVnets {
	s.MbpsQuota = &v
	return s
}

func (s *DescribeVnetsResponseVnets) SetKppsQuota(v int) *DescribeVnetsResponseVnets {
	s.KppsQuota = &v
	return s
}

func (s *DescribeVnetsResponseVnets) SetFlowQuota(v int) *DescribeVnetsResponseVnets {
	s.FlowQuota = &v
	return s
}

func (s *DescribeVnetsResponseVnets) SetState(v string) *DescribeVnetsResponseVnets {
	s.State = &v
	return s
}

type DeleteVnetRouteEntryRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	UisId       *string `json:"UisId,omitempty" xml:"UisId,omitempty" require:"true"`
	VnetId      *string `json:"VnetId,omitempty" xml:"VnetId,omitempty" require:"true"`
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty" require:"true"`
}

func (s DeleteVnetRouteEntryRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVnetRouteEntryRequest) GoString() string {
	return s.String()
}

func (s *DeleteVnetRouteEntryRequest) SetClientToken(v string) *DeleteVnetRouteEntryRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteVnetRouteEntryRequest) SetRegionId(v string) *DeleteVnetRouteEntryRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteVnetRouteEntryRequest) SetUisId(v string) *DeleteVnetRouteEntryRequest {
	s.UisId = &v
	return s
}

func (s *DeleteVnetRouteEntryRequest) SetVnetId(v string) *DeleteVnetRouteEntryRequest {
	s.VnetId = &v
	return s
}

func (s *DeleteVnetRouteEntryRequest) SetDestination(v string) *DeleteVnetRouteEntryRequest {
	s.Destination = &v
	return s
}

type DeleteVnetRouteEntryResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteVnetRouteEntryResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVnetRouteEntryResponse) GoString() string {
	return s.String()
}

func (s *DeleteVnetRouteEntryResponse) SetRequestId(v string) *DeleteVnetRouteEntryResponse {
	s.RequestId = &v
	return s
}

type AssociateEipRequest struct {
	UisId       *string `json:"UisId,omitempty" xml:"UisId,omitempty" require:"true"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	Role        *string `json:"Role,omitempty" xml:"Role,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AssociateEipRequest) String() string {
	return tea.Prettify(s)
}

func (s AssociateEipRequest) GoString() string {
	return s.String()
}

func (s *AssociateEipRequest) SetUisId(v string) *AssociateEipRequest {
	s.UisId = &v
	return s
}

func (s *AssociateEipRequest) SetType(v string) *AssociateEipRequest {
	s.Type = &v
	return s
}

func (s *AssociateEipRequest) SetInstanceId(v string) *AssociateEipRequest {
	s.InstanceId = &v
	return s
}

func (s *AssociateEipRequest) SetRole(v string) *AssociateEipRequest {
	s.Role = &v
	return s
}

func (s *AssociateEipRequest) SetClientToken(v string) *AssociateEipRequest {
	s.ClientToken = &v
	return s
}

func (s *AssociateEipRequest) SetRegionId(v string) *AssociateEipRequest {
	s.RegionId = &v
	return s
}

type AssociateEipResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s AssociateEipResponse) String() string {
	return tea.Prettify(s)
}

func (s AssociateEipResponse) GoString() string {
	return s.String()
}

func (s *AssociateEipResponse) SetRequestId(v string) *AssociateEipResponse {
	s.RequestId = &v
	return s
}

type ModifyGreConnectionRequest struct {
	ClientToken      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	UisId            *string `json:"UisId,omitempty" xml:"UisId,omitempty" require:"true"`
	GreConnectionId  *string `json:"GreConnectionId,omitempty" xml:"GreConnectionId,omitempty" require:"true"`
	Cidrs            *string `json:"Cidrs,omitempty" xml:"Cidrs,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	LocalIp          *string `json:"LocalIp,omitempty" xml:"LocalIp,omitempty"`
	CustomerIp       *string `json:"CustomerIp,omitempty" xml:"CustomerIp,omitempty"`
	LocalTunnelIp    *string `json:"LocalTunnelIp,omitempty" xml:"LocalTunnelIp,omitempty"`
	CustomerTunnelIp *string `json:"CustomerTunnelIp,omitempty" xml:"CustomerTunnelIp,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s ModifyGreConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyGreConnectionRequest) GoString() string {
	return s.String()
}

func (s *ModifyGreConnectionRequest) SetClientToken(v string) *ModifyGreConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyGreConnectionRequest) SetRegionId(v string) *ModifyGreConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyGreConnectionRequest) SetUisId(v string) *ModifyGreConnectionRequest {
	s.UisId = &v
	return s
}

func (s *ModifyGreConnectionRequest) SetGreConnectionId(v string) *ModifyGreConnectionRequest {
	s.GreConnectionId = &v
	return s
}

func (s *ModifyGreConnectionRequest) SetCidrs(v string) *ModifyGreConnectionRequest {
	s.Cidrs = &v
	return s
}

func (s *ModifyGreConnectionRequest) SetName(v string) *ModifyGreConnectionRequest {
	s.Name = &v
	return s
}

func (s *ModifyGreConnectionRequest) SetLocalIp(v string) *ModifyGreConnectionRequest {
	s.LocalIp = &v
	return s
}

func (s *ModifyGreConnectionRequest) SetCustomerIp(v string) *ModifyGreConnectionRequest {
	s.CustomerIp = &v
	return s
}

func (s *ModifyGreConnectionRequest) SetLocalTunnelIp(v string) *ModifyGreConnectionRequest {
	s.LocalTunnelIp = &v
	return s
}

func (s *ModifyGreConnectionRequest) SetCustomerTunnelIp(v string) *ModifyGreConnectionRequest {
	s.CustomerTunnelIp = &v
	return s
}

func (s *ModifyGreConnectionRequest) SetDescription(v string) *ModifyGreConnectionRequest {
	s.Description = &v
	return s
}

type ModifyGreConnectionResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ModifyGreConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyGreConnectionResponse) GoString() string {
	return s.String()
}

func (s *ModifyGreConnectionResponse) SetRequestId(v string) *ModifyGreConnectionResponse {
	s.RequestId = &v
	return s
}

type DescribeVnetRequest struct {
	UisId    *string `json:"UisId,omitempty" xml:"UisId,omitempty" require:"true"`
	VnetId   *string `json:"VnetId,omitempty" xml:"VnetId,omitempty" require:"true"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
}

func (s DescribeVnetRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVnetRequest) GoString() string {
	return s.String()
}

func (s *DescribeVnetRequest) SetUisId(v string) *DescribeVnetRequest {
	s.UisId = &v
	return s
}

func (s *DescribeVnetRequest) SetVnetId(v string) *DescribeVnetRequest {
	s.VnetId = &v
	return s
}

func (s *DescribeVnetRequest) SetRegionId(v string) *DescribeVnetRequest {
	s.RegionId = &v
	return s
}

type DescribeVnetResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	UisId     *string `json:"UisId,omitempty" xml:"UisId,omitempty" require:"true"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	State     *string `json:"State,omitempty" xml:"State,omitempty" require:"true"`
	VnetId    *string `json:"VnetId,omitempty" xml:"VnetId,omitempty" require:"true"`
	Cidrs     *string `json:"Cidrs,omitempty" xml:"Cidrs,omitempty" require:"true"`
	MbpsQuota *string `json:"MbpsQuota,omitempty" xml:"MbpsQuota,omitempty" require:"true"`
	KppsQuota *string `json:"KppsQuota,omitempty" xml:"KppsQuota,omitempty" require:"true"`
	FlowQuota *string `json:"FlowQuota,omitempty" xml:"FlowQuota,omitempty" require:"true"`
}

func (s DescribeVnetResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVnetResponse) GoString() string {
	return s.String()
}

func (s *DescribeVnetResponse) SetRequestId(v string) *DescribeVnetResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeVnetResponse) SetUisId(v string) *DescribeVnetResponse {
	s.UisId = &v
	return s
}

func (s *DescribeVnetResponse) SetName(v string) *DescribeVnetResponse {
	s.Name = &v
	return s
}

func (s *DescribeVnetResponse) SetState(v string) *DescribeVnetResponse {
	s.State = &v
	return s
}

func (s *DescribeVnetResponse) SetVnetId(v string) *DescribeVnetResponse {
	s.VnetId = &v
	return s
}

func (s *DescribeVnetResponse) SetCidrs(v string) *DescribeVnetResponse {
	s.Cidrs = &v
	return s
}

func (s *DescribeVnetResponse) SetMbpsQuota(v string) *DescribeVnetResponse {
	s.MbpsQuota = &v
	return s
}

func (s *DescribeVnetResponse) SetKppsQuota(v string) *DescribeVnetResponse {
	s.KppsQuota = &v
	return s
}

func (s *DescribeVnetResponse) SetFlowQuota(v string) *DescribeVnetResponse {
	s.FlowQuota = &v
	return s
}

type UnAssociateEipRequest struct {
	UisId       *string `json:"UisId,omitempty" xml:"UisId,omitempty" require:"true"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UnAssociateEipRequest) String() string {
	return tea.Prettify(s)
}

func (s UnAssociateEipRequest) GoString() string {
	return s.String()
}

func (s *UnAssociateEipRequest) SetUisId(v string) *UnAssociateEipRequest {
	s.UisId = &v
	return s
}

func (s *UnAssociateEipRequest) SetType(v string) *UnAssociateEipRequest {
	s.Type = &v
	return s
}

func (s *UnAssociateEipRequest) SetInstanceId(v string) *UnAssociateEipRequest {
	s.InstanceId = &v
	return s
}

func (s *UnAssociateEipRequest) SetClientToken(v string) *UnAssociateEipRequest {
	s.ClientToken = &v
	return s
}

func (s *UnAssociateEipRequest) SetRegionId(v string) *UnAssociateEipRequest {
	s.RegionId = &v
	return s
}

type UnAssociateEipResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UnAssociateEipResponse) String() string {
	return tea.Prettify(s)
}

func (s UnAssociateEipResponse) GoString() string {
	return s.String()
}

func (s *UnAssociateEipResponse) SetRequestId(v string) *UnAssociateEipResponse {
	s.RequestId = &v
	return s
}

type DeleteGreConnectionRequest struct {
	UisId           *string `json:"UisId,omitempty" xml:"UisId,omitempty" require:"true"`
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	GreConnectionId *string `json:"GreConnectionId,omitempty" xml:"GreConnectionId,omitempty" require:"true"`
}

func (s DeleteGreConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGreConnectionRequest) GoString() string {
	return s.String()
}

func (s *DeleteGreConnectionRequest) SetUisId(v string) *DeleteGreConnectionRequest {
	s.UisId = &v
	return s
}

func (s *DeleteGreConnectionRequest) SetClientToken(v string) *DeleteGreConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteGreConnectionRequest) SetRegionId(v string) *DeleteGreConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteGreConnectionRequest) SetGreConnectionId(v string) *DeleteGreConnectionRequest {
	s.GreConnectionId = &v
	return s
}

type DeleteGreConnectionResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteGreConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGreConnectionResponse) GoString() string {
	return s.String()
}

func (s *DeleteGreConnectionResponse) SetRequestId(v string) *DeleteGreConnectionResponse {
	s.RequestId = &v
	return s
}

type CreateGreConnectionRequest struct {
	UisId            *string `json:"UisId,omitempty" xml:"UisId,omitempty" require:"true"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Cidrs            *string `json:"Cidrs,omitempty" xml:"Cidrs,omitempty" require:"true"`
	LocalIp          *string `json:"LocalIp,omitempty" xml:"LocalIp,omitempty" require:"true"`
	LocalTunnelIp    *string `json:"LocalTunnelIp,omitempty" xml:"LocalTunnelIp,omitempty" require:"true"`
	CustomerIp       *string `json:"CustomerIp,omitempty" xml:"CustomerIp,omitempty" require:"true"`
	CustomerTunnelIp *string `json:"CustomerTunnelIp,omitempty" xml:"CustomerTunnelIp,omitempty" require:"true"`
	ClientToken      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateGreConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGreConnectionRequest) GoString() string {
	return s.String()
}

func (s *CreateGreConnectionRequest) SetUisId(v string) *CreateGreConnectionRequest {
	s.UisId = &v
	return s
}

func (s *CreateGreConnectionRequest) SetName(v string) *CreateGreConnectionRequest {
	s.Name = &v
	return s
}

func (s *CreateGreConnectionRequest) SetDescription(v string) *CreateGreConnectionRequest {
	s.Description = &v
	return s
}

func (s *CreateGreConnectionRequest) SetCidrs(v string) *CreateGreConnectionRequest {
	s.Cidrs = &v
	return s
}

func (s *CreateGreConnectionRequest) SetLocalIp(v string) *CreateGreConnectionRequest {
	s.LocalIp = &v
	return s
}

func (s *CreateGreConnectionRequest) SetLocalTunnelIp(v string) *CreateGreConnectionRequest {
	s.LocalTunnelIp = &v
	return s
}

func (s *CreateGreConnectionRequest) SetCustomerIp(v string) *CreateGreConnectionRequest {
	s.CustomerIp = &v
	return s
}

func (s *CreateGreConnectionRequest) SetCustomerTunnelIp(v string) *CreateGreConnectionRequest {
	s.CustomerTunnelIp = &v
	return s
}

func (s *CreateGreConnectionRequest) SetClientToken(v string) *CreateGreConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateGreConnectionRequest) SetRegionId(v string) *CreateGreConnectionRequest {
	s.RegionId = &v
	return s
}

type CreateGreConnectionResponse struct {
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	GreConnectionId *string `json:"GreConnectionId,omitempty" xml:"GreConnectionId,omitempty" require:"true"`
}

func (s CreateGreConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGreConnectionResponse) GoString() string {
	return s.String()
}

func (s *CreateGreConnectionResponse) SetRequestId(v string) *CreateGreConnectionResponse {
	s.RequestId = &v
	return s
}

func (s *CreateGreConnectionResponse) SetGreConnectionId(v string) *CreateGreConnectionResponse {
	s.GreConnectionId = &v
	return s
}

type CreateUisRequest struct {
	UisName         *string `json:"UisName,omitempty" xml:"UisName,omitempty" require:"true"`
	Spec            *string `json:"Spec,omitempty" xml:"Spec,omitempty" require:"true"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
}

func (s CreateUisRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUisRequest) GoString() string {
	return s.String()
}

func (s *CreateUisRequest) SetUisName(v string) *CreateUisRequest {
	s.UisName = &v
	return s
}

func (s *CreateUisRequest) SetSpec(v string) *CreateUisRequest {
	s.Spec = &v
	return s
}

func (s *CreateUisRequest) SetDescription(v string) *CreateUisRequest {
	s.Description = &v
	return s
}

func (s *CreateUisRequest) SetResourceGroupId(v string) *CreateUisRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateUisRequest) SetClientToken(v string) *CreateUisRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateUisRequest) SetRegionId(v string) *CreateUisRequest {
	s.RegionId = &v
	return s
}

type CreateUisResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	UisId     *string `json:"UisId,omitempty" xml:"UisId,omitempty" require:"true"`
}

func (s CreateUisResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUisResponse) GoString() string {
	return s.String()
}

func (s *CreateUisResponse) SetRequestId(v string) *CreateUisResponse {
	s.RequestId = &v
	return s
}

func (s *CreateUisResponse) SetUisId(v string) *CreateUisResponse {
	s.UisId = &v
	return s
}

type CreateVnetRouteEntryRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	UisId       *string `json:"UisId,omitempty" xml:"UisId,omitempty" require:"true"`
	VnetId      *string `json:"VnetId,omitempty" xml:"VnetId,omitempty" require:"true"`
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty" require:"true"`
	NexthopType *string `json:"NexthopType,omitempty" xml:"NexthopType,omitempty" require:"true"`
	Nexthop     *string `json:"Nexthop,omitempty" xml:"Nexthop,omitempty"`
}

func (s CreateVnetRouteEntryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVnetRouteEntryRequest) GoString() string {
	return s.String()
}

func (s *CreateVnetRouteEntryRequest) SetClientToken(v string) *CreateVnetRouteEntryRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateVnetRouteEntryRequest) SetRegionId(v string) *CreateVnetRouteEntryRequest {
	s.RegionId = &v
	return s
}

func (s *CreateVnetRouteEntryRequest) SetUisId(v string) *CreateVnetRouteEntryRequest {
	s.UisId = &v
	return s
}

func (s *CreateVnetRouteEntryRequest) SetVnetId(v string) *CreateVnetRouteEntryRequest {
	s.VnetId = &v
	return s
}

func (s *CreateVnetRouteEntryRequest) SetDestination(v string) *CreateVnetRouteEntryRequest {
	s.Destination = &v
	return s
}

func (s *CreateVnetRouteEntryRequest) SetNexthopType(v string) *CreateVnetRouteEntryRequest {
	s.NexthopType = &v
	return s
}

func (s *CreateVnetRouteEntryRequest) SetNexthop(v string) *CreateVnetRouteEntryRequest {
	s.Nexthop = &v
	return s
}

type CreateVnetRouteEntryResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s CreateVnetRouteEntryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVnetRouteEntryResponse) GoString() string {
	return s.String()
}

func (s *CreateVnetRouteEntryResponse) SetRequestId(v string) *CreateVnetRouteEntryResponse {
	s.RequestId = &v
	return s
}

type DeleteUisRequest struct {
	UisId       *string `json:"UisId,omitempty" xml:"UisId,omitempty" require:"true"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
}

func (s DeleteUisRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUisRequest) GoString() string {
	return s.String()
}

func (s *DeleteUisRequest) SetUisId(v string) *DeleteUisRequest {
	s.UisId = &v
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

type DeleteUisResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteUisResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUisResponse) GoString() string {
	return s.String()
}

func (s *DeleteUisResponse) SetRequestId(v string) *DeleteUisResponse {
	s.RequestId = &v
	return s
}

type DescribeGreConnectionRequest struct {
	UisId           *string `json:"UisId,omitempty" xml:"UisId,omitempty" require:"true"`
	GreConnectionId *string `json:"GreConnectionId,omitempty" xml:"GreConnectionId,omitempty" require:"true"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
}

func (s DescribeGreConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGreConnectionRequest) GoString() string {
	return s.String()
}

func (s *DescribeGreConnectionRequest) SetUisId(v string) *DescribeGreConnectionRequest {
	s.UisId = &v
	return s
}

func (s *DescribeGreConnectionRequest) SetGreConnectionId(v string) *DescribeGreConnectionRequest {
	s.GreConnectionId = &v
	return s
}

func (s *DescribeGreConnectionRequest) SetRegionId(v string) *DescribeGreConnectionRequest {
	s.RegionId = &v
	return s
}

type DescribeGreConnectionResponse struct {
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	UisId            *string `json:"UisId,omitempty" xml:"UisId,omitempty" require:"true"`
	UisName          *string `json:"UisName,omitempty" xml:"UisName,omitempty" require:"true"`
	GreConnectionId  *string `json:"GreConnectionId,omitempty" xml:"GreConnectionId,omitempty" require:"true"`
	State            *string `json:"State,omitempty" xml:"State,omitempty" require:"true"`
	Cidrs            *string `json:"Cidrs,omitempty" xml:"Cidrs,omitempty" require:"true"`
	LocalIp          *string `json:"LocalIp,omitempty" xml:"LocalIp,omitempty" require:"true"`
	LocalTunnelIp    *string `json:"LocalTunnelIp,omitempty" xml:"LocalTunnelIp,omitempty" require:"true"`
	CustomerIp       *string `json:"CustomerIp,omitempty" xml:"CustomerIp,omitempty" require:"true"`
	CustomerTunnelIp *string `json:"CustomerTunnelIp,omitempty" xml:"CustomerTunnelIp,omitempty" require:"true"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	CreateTime       *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
}

func (s DescribeGreConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGreConnectionResponse) GoString() string {
	return s.String()
}

func (s *DescribeGreConnectionResponse) SetRequestId(v string) *DescribeGreConnectionResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeGreConnectionResponse) SetUisId(v string) *DescribeGreConnectionResponse {
	s.UisId = &v
	return s
}

func (s *DescribeGreConnectionResponse) SetUisName(v string) *DescribeGreConnectionResponse {
	s.UisName = &v
	return s
}

func (s *DescribeGreConnectionResponse) SetGreConnectionId(v string) *DescribeGreConnectionResponse {
	s.GreConnectionId = &v
	return s
}

func (s *DescribeGreConnectionResponse) SetState(v string) *DescribeGreConnectionResponse {
	s.State = &v
	return s
}

func (s *DescribeGreConnectionResponse) SetCidrs(v string) *DescribeGreConnectionResponse {
	s.Cidrs = &v
	return s
}

func (s *DescribeGreConnectionResponse) SetLocalIp(v string) *DescribeGreConnectionResponse {
	s.LocalIp = &v
	return s
}

func (s *DescribeGreConnectionResponse) SetLocalTunnelIp(v string) *DescribeGreConnectionResponse {
	s.LocalTunnelIp = &v
	return s
}

func (s *DescribeGreConnectionResponse) SetCustomerIp(v string) *DescribeGreConnectionResponse {
	s.CustomerIp = &v
	return s
}

func (s *DescribeGreConnectionResponse) SetCustomerTunnelIp(v string) *DescribeGreConnectionResponse {
	s.CustomerTunnelIp = &v
	return s
}

func (s *DescribeGreConnectionResponse) SetName(v string) *DescribeGreConnectionResponse {
	s.Name = &v
	return s
}

func (s *DescribeGreConnectionResponse) SetDescription(v string) *DescribeGreConnectionResponse {
	s.Description = &v
	return s
}

func (s *DescribeGreConnectionResponse) SetCreateTime(v int64) *DescribeGreConnectionResponse {
	s.CreateTime = &v
	return s
}

type Client struct {
	rpc.Client
}

func NewClient(config *rpc.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *rpc.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("uisplus"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) DescribeUisWithOptions(request *DescribeUisRequest, runtime *util.RuntimeOptions) (_result *DescribeUisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeUisResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeUis"), tea.String("HTTPS"), tea.String("GET"), tea.String("2020-07-07"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUis(request *DescribeUisRequest) (_result *DescribeUisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUisResponse{}
	_body, _err := client.DescribeUisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyVnetWithOptions(request *ModifyVnetRequest, runtime *util.RuntimeOptions) (_result *ModifyVnetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyVnetResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyVnet"), tea.String("HTTPS"), tea.String("GET"), tea.String("2020-07-07"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyVnet(request *ModifyVnetRequest) (_result *ModifyVnetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyVnetResponse{}
	_body, _err := client.ModifyVnetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGreConnectionsWithOptions(request *DescribeGreConnectionsRequest, runtime *util.RuntimeOptions) (_result *DescribeGreConnectionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeGreConnectionsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeGreConnections"), tea.String("HTTPS"), tea.String("GET"), tea.String("2020-07-07"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGreConnections(request *DescribeGreConnectionsRequest) (_result *DescribeGreConnectionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGreConnectionsResponse{}
	_body, _err := client.DescribeGreConnectionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUissWithOptions(request *DescribeUissRequest, runtime *util.RuntimeOptions) (_result *DescribeUissResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeUissResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeUiss"), tea.String("HTTPS"), tea.String("GET"), tea.String("2020-07-07"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUiss(request *DescribeUissRequest) (_result *DescribeUissResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUissResponse{}
	_body, _err := client.DescribeUissWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVnetRouteEntryListWithOptions(request *DescribeVnetRouteEntryListRequest, runtime *util.RuntimeOptions) (_result *DescribeVnetRouteEntryListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeVnetRouteEntryListResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeVnetRouteEntryList"), tea.String("HTTPS"), tea.String("GET"), tea.String("2020-07-07"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVnetRouteEntryList(request *DescribeVnetRouteEntryListRequest) (_result *DescribeVnetRouteEntryListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVnetRouteEntryListResponse{}
	_body, _err := client.DescribeVnetRouteEntryListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVnetWithOptions(request *CreateVnetRequest, runtime *util.RuntimeOptions) (_result *CreateVnetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateVnetResponse{}
	_body, _err := client.DoRequest(tea.String("CreateVnet"), tea.String("HTTPS"), tea.String("GET"), tea.String("2020-07-07"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVnet(request *CreateVnetRequest) (_result *CreateVnetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVnetResponse{}
	_body, _err := client.CreateVnetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVnetWithOptions(request *DeleteVnetRequest, runtime *util.RuntimeOptions) (_result *DeleteVnetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteVnetResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteVnet"), tea.String("HTTPS"), tea.String("GET"), tea.String("2020-07-07"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVnet(request *DeleteVnetRequest) (_result *DeleteVnetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVnetResponse{}
	_body, _err := client.DeleteVnetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVnetsWithOptions(request *DescribeVnetsRequest, runtime *util.RuntimeOptions) (_result *DescribeVnetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeVnetsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeVnets"), tea.String("HTTPS"), tea.String("GET"), tea.String("2020-07-07"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVnets(request *DescribeVnetsRequest) (_result *DescribeVnetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVnetsResponse{}
	_body, _err := client.DescribeVnetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVnetRouteEntryWithOptions(request *DeleteVnetRouteEntryRequest, runtime *util.RuntimeOptions) (_result *DeleteVnetRouteEntryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteVnetRouteEntryResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteVnetRouteEntry"), tea.String("HTTPS"), tea.String("GET"), tea.String("2020-07-07"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVnetRouteEntry(request *DeleteVnetRouteEntryRequest) (_result *DeleteVnetRouteEntryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVnetRouteEntryResponse{}
	_body, _err := client.DeleteVnetRouteEntryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AssociateEipWithOptions(request *AssociateEipRequest, runtime *util.RuntimeOptions) (_result *AssociateEipResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AssociateEipResponse{}
	_body, _err := client.DoRequest(tea.String("AssociateEip"), tea.String("HTTPS"), tea.String("GET"), tea.String("2020-07-07"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AssociateEip(request *AssociateEipRequest) (_result *AssociateEipResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssociateEipResponse{}
	_body, _err := client.AssociateEipWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyGreConnectionWithOptions(request *ModifyGreConnectionRequest, runtime *util.RuntimeOptions) (_result *ModifyGreConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyGreConnectionResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyGreConnection"), tea.String("HTTPS"), tea.String("GET"), tea.String("2020-07-07"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyGreConnection(request *ModifyGreConnectionRequest) (_result *ModifyGreConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyGreConnectionResponse{}
	_body, _err := client.ModifyGreConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVnetWithOptions(request *DescribeVnetRequest, runtime *util.RuntimeOptions) (_result *DescribeVnetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeVnetResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeVnet"), tea.String("HTTPS"), tea.String("GET"), tea.String("2020-07-07"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVnet(request *DescribeVnetRequest) (_result *DescribeVnetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVnetResponse{}
	_body, _err := client.DescribeVnetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnAssociateEipWithOptions(request *UnAssociateEipRequest, runtime *util.RuntimeOptions) (_result *UnAssociateEipResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UnAssociateEipResponse{}
	_body, _err := client.DoRequest(tea.String("UnAssociateEip"), tea.String("HTTPS"), tea.String("GET"), tea.String("2020-07-07"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnAssociateEip(request *UnAssociateEipRequest) (_result *UnAssociateEipResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnAssociateEipResponse{}
	_body, _err := client.UnAssociateEipWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteGreConnectionWithOptions(request *DeleteGreConnectionRequest, runtime *util.RuntimeOptions) (_result *DeleteGreConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteGreConnectionResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteGreConnection"), tea.String("HTTPS"), tea.String("GET"), tea.String("2020-07-07"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGreConnection(request *DeleteGreConnectionRequest) (_result *DeleteGreConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGreConnectionResponse{}
	_body, _err := client.DeleteGreConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGreConnectionWithOptions(request *CreateGreConnectionRequest, runtime *util.RuntimeOptions) (_result *CreateGreConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateGreConnectionResponse{}
	_body, _err := client.DoRequest(tea.String("CreateGreConnection"), tea.String("HTTPS"), tea.String("GET"), tea.String("2020-07-07"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGreConnection(request *CreateGreConnectionRequest) (_result *CreateGreConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGreConnectionResponse{}
	_body, _err := client.CreateGreConnectionWithOptions(request, runtime)
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
	_result = &CreateUisResponse{}
	_body, _err := client.DoRequest(tea.String("CreateUis"), tea.String("HTTPS"), tea.String("GET"), tea.String("2020-07-07"), tea.String("AK"), tea.ToMap(request), nil, runtime)
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

func (client *Client) CreateVnetRouteEntryWithOptions(request *CreateVnetRouteEntryRequest, runtime *util.RuntimeOptions) (_result *CreateVnetRouteEntryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateVnetRouteEntryResponse{}
	_body, _err := client.DoRequest(tea.String("CreateVnetRouteEntry"), tea.String("HTTPS"), tea.String("GET"), tea.String("2020-07-07"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVnetRouteEntry(request *CreateVnetRouteEntryRequest) (_result *CreateVnetRouteEntryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVnetRouteEntryResponse{}
	_body, _err := client.CreateVnetRouteEntryWithOptions(request, runtime)
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
	_result = &DeleteUisResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteUis"), tea.String("HTTPS"), tea.String("GET"), tea.String("2020-07-07"), tea.String("AK"), tea.ToMap(request), nil, runtime)
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

func (client *Client) DescribeGreConnectionWithOptions(request *DescribeGreConnectionRequest, runtime *util.RuntimeOptions) (_result *DescribeGreConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeGreConnectionResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeGreConnection"), tea.String("HTTPS"), tea.String("GET"), tea.String("2020-07-07"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGreConnection(request *DescribeGreConnectionRequest) (_result *DescribeGreConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGreConnectionResponse{}
	_body, _err := client.DescribeGreConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
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
