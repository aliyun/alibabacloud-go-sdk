// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateLindormInstanceRequest struct {
	ColdStorage          *int32  `json:"ColdStorage,omitempty" xml:"ColdStorage,omitempty"`
	CoreSpec             *string `json:"CoreSpec,omitempty" xml:"CoreSpec,omitempty"`
	DiskCategory         *string `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	Duration             *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	FilestoreNum         *int32  `json:"FilestoreNum,omitempty" xml:"FilestoreNum,omitempty"`
	FilestoreSpec        *string `json:"FilestoreSpec,omitempty" xml:"FilestoreSpec,omitempty"`
	InstanceAlias        *string `json:"InstanceAlias,omitempty" xml:"InstanceAlias,omitempty"`
	InstanceStorage      *string `json:"InstanceStorage,omitempty" xml:"InstanceStorage,omitempty"`
	LindormNum           *int32  `json:"LindormNum,omitempty" xml:"LindormNum,omitempty"`
	LindormSpec          *string `json:"LindormSpec,omitempty" xml:"LindormSpec,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PayType              *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	PricingCycle         *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	SolrNum              *int32  `json:"SolrNum,omitempty" xml:"SolrNum,omitempty"`
	SolrSpec             *string `json:"SolrSpec,omitempty" xml:"SolrSpec,omitempty"`
	TsdbNum              *int32  `json:"TsdbNum,omitempty" xml:"TsdbNum,omitempty"`
	TsdbSpec             *string `json:"TsdbSpec,omitempty" xml:"TsdbSpec,omitempty"`
	VPCId                *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId            *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateLindormInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLindormInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateLindormInstanceRequest) SetColdStorage(v int32) *CreateLindormInstanceRequest {
	s.ColdStorage = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetCoreSpec(v string) *CreateLindormInstanceRequest {
	s.CoreSpec = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetDiskCategory(v string) *CreateLindormInstanceRequest {
	s.DiskCategory = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetDuration(v string) *CreateLindormInstanceRequest {
	s.Duration = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetFilestoreNum(v int32) *CreateLindormInstanceRequest {
	s.FilestoreNum = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetFilestoreSpec(v string) *CreateLindormInstanceRequest {
	s.FilestoreSpec = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetInstanceAlias(v string) *CreateLindormInstanceRequest {
	s.InstanceAlias = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetInstanceStorage(v string) *CreateLindormInstanceRequest {
	s.InstanceStorage = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetLindormNum(v int32) *CreateLindormInstanceRequest {
	s.LindormNum = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetLindormSpec(v string) *CreateLindormInstanceRequest {
	s.LindormSpec = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetOwnerAccount(v string) *CreateLindormInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetOwnerId(v int64) *CreateLindormInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetPayType(v string) *CreateLindormInstanceRequest {
	s.PayType = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetPricingCycle(v string) *CreateLindormInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetRegionId(v string) *CreateLindormInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetResourceOwnerAccount(v string) *CreateLindormInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetResourceOwnerId(v int64) *CreateLindormInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetSecurityToken(v string) *CreateLindormInstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetSolrNum(v int32) *CreateLindormInstanceRequest {
	s.SolrNum = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetSolrSpec(v string) *CreateLindormInstanceRequest {
	s.SolrSpec = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetTsdbNum(v int32) *CreateLindormInstanceRequest {
	s.TsdbNum = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetTsdbSpec(v string) *CreateLindormInstanceRequest {
	s.TsdbSpec = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetVPCId(v string) *CreateLindormInstanceRequest {
	s.VPCId = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetVSwitchId(v string) *CreateLindormInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetZoneId(v string) *CreateLindormInstanceRequest {
	s.ZoneId = &v
	return s
}

type CreateLindormInstanceResponseBody struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OrderId    *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLindormInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLindormInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLindormInstanceResponseBody) SetInstanceId(v string) *CreateLindormInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateLindormInstanceResponseBody) SetOrderId(v int64) *CreateLindormInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateLindormInstanceResponseBody) SetRequestId(v string) *CreateLindormInstanceResponseBody {
	s.RequestId = &v
	return s
}

type CreateLindormInstanceResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateLindormInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateLindormInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLindormInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateLindormInstanceResponse) SetHeaders(v map[string]*string) *CreateLindormInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateLindormInstanceResponse) SetBody(v *CreateLindormInstanceResponseBody) *CreateLindormInstanceResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	AcceptLanguage       *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
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

func (s *DescribeRegionsRequest) SetSecurityToken(v string) *DescribeRegionsRequest {
	s.SecurityToken = &v
	return s
}

type DescribeRegionsResponseBody struct {
	Regions   []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	LocalName      *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetLocalName(v string) *DescribeRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionId(v string) *DescribeRegionsResponseBodyRegions {
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

type GetInstanceIpWhiteListRequest struct {
	GroupName            *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetInstanceIpWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceIpWhiteListRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceIpWhiteListRequest) SetGroupName(v string) *GetInstanceIpWhiteListRequest {
	s.GroupName = &v
	return s
}

func (s *GetInstanceIpWhiteListRequest) SetInstanceId(v string) *GetInstanceIpWhiteListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceIpWhiteListRequest) SetOwnerAccount(v string) *GetInstanceIpWhiteListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetInstanceIpWhiteListRequest) SetOwnerId(v int64) *GetInstanceIpWhiteListRequest {
	s.OwnerId = &v
	return s
}

func (s *GetInstanceIpWhiteListRequest) SetResourceOwnerAccount(v string) *GetInstanceIpWhiteListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetInstanceIpWhiteListRequest) SetResourceOwnerId(v int64) *GetInstanceIpWhiteListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetInstanceIpWhiteListRequest) SetSecurityToken(v string) *GetInstanceIpWhiteListRequest {
	s.SecurityToken = &v
	return s
}

type GetInstanceIpWhiteListResponseBody struct {
	InstanceId *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IpList     []*string `json:"IpList,omitempty" xml:"IpList,omitempty" type:"Repeated"`
	RequestId  *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInstanceIpWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceIpWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceIpWhiteListResponseBody) SetInstanceId(v string) *GetInstanceIpWhiteListResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceIpWhiteListResponseBody) SetIpList(v []*string) *GetInstanceIpWhiteListResponseBody {
	s.IpList = v
	return s
}

func (s *GetInstanceIpWhiteListResponseBody) SetRequestId(v string) *GetInstanceIpWhiteListResponseBody {
	s.RequestId = &v
	return s
}

type GetInstanceIpWhiteListResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetInstanceIpWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInstanceIpWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceIpWhiteListResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceIpWhiteListResponse) SetHeaders(v map[string]*string) *GetInstanceIpWhiteListResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceIpWhiteListResponse) SetBody(v *GetInstanceIpWhiteListResponseBody) *GetInstanceIpWhiteListResponse {
	s.Body = v
	return s
}

type GetLindormInstanceRequest struct {
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetLindormInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLindormInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceRequest) SetInstanceId(v string) *GetLindormInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *GetLindormInstanceRequest) SetOwnerAccount(v string) *GetLindormInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetLindormInstanceRequest) SetOwnerId(v int64) *GetLindormInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *GetLindormInstanceRequest) SetRegionId(v string) *GetLindormInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *GetLindormInstanceRequest) SetResourceOwnerAccount(v string) *GetLindormInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetLindormInstanceRequest) SetResourceOwnerId(v int64) *GetLindormInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetLindormInstanceRequest) SetSecurityToken(v string) *GetLindormInstanceRequest {
	s.SecurityToken = &v
	return s
}

type GetLindormInstanceResponseBody struct {
	AliUid              *int64                                      `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	AutoRenew           *bool                                       `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	ColdStorage         *int32                                      `json:"ColdStorage,omitempty" xml:"ColdStorage,omitempty"`
	CreateMilliseconds  *int64                                      `json:"CreateMilliseconds,omitempty" xml:"CreateMilliseconds,omitempty"`
	CreateTime          *string                                     `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DeletionProtection  *string                                     `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	DiskCategory        *string                                     `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	DiskThreshold       *string                                     `json:"DiskThreshold,omitempty" xml:"DiskThreshold,omitempty"`
	DiskUsage           *string                                     `json:"DiskUsage,omitempty" xml:"DiskUsage,omitempty"`
	EnableCdc           *bool                                       `json:"EnableCdc,omitempty" xml:"EnableCdc,omitempty"`
	EnableCompute       *bool                                       `json:"EnableCompute,omitempty" xml:"EnableCompute,omitempty"`
	EnableKms           *bool                                       `json:"EnableKms,omitempty" xml:"EnableKms,omitempty"`
	EnableSSL           *bool                                       `json:"EnableSSL,omitempty" xml:"EnableSSL,omitempty"`
	EnableShs           *bool                                       `json:"EnableShs,omitempty" xml:"EnableShs,omitempty"`
	EnableStream        *bool                                       `json:"EnableStream,omitempty" xml:"EnableStream,omitempty"`
	EngineList          []*GetLindormInstanceResponseBodyEngineList `json:"EngineList,omitempty" xml:"EngineList,omitempty" type:"Repeated"`
	EngineType          *int32                                      `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	ExpireTime          *string                                     `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	ExpiredMilliseconds *int64                                      `json:"ExpiredMilliseconds,omitempty" xml:"ExpiredMilliseconds,omitempty"`
	InstanceAlias       *string                                     `json:"InstanceAlias,omitempty" xml:"InstanceAlias,omitempty"`
	InstanceId          *string                                     `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceStatus      *string                                     `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	InstanceStorage     *string                                     `json:"InstanceStorage,omitempty" xml:"InstanceStorage,omitempty"`
	MaintainEndTime     *string                                     `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	MaintainStartTime   *string                                     `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	NetworkType         *string                                     `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	PayType             *string                                     `json:"PayType,omitempty" xml:"PayType,omitempty"`
	RegionId            *string                                     `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId           *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceType         *string                                     `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	VpcId               *string                                     `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitchId           *string                                     `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	ZoneId              *string                                     `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetLindormInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLindormInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceResponseBody) SetAliUid(v int64) *GetLindormInstanceResponseBody {
	s.AliUid = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetAutoRenew(v bool) *GetLindormInstanceResponseBody {
	s.AutoRenew = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetColdStorage(v int32) *GetLindormInstanceResponseBody {
	s.ColdStorage = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetCreateMilliseconds(v int64) *GetLindormInstanceResponseBody {
	s.CreateMilliseconds = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetCreateTime(v string) *GetLindormInstanceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetDeletionProtection(v string) *GetLindormInstanceResponseBody {
	s.DeletionProtection = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetDiskCategory(v string) *GetLindormInstanceResponseBody {
	s.DiskCategory = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetDiskThreshold(v string) *GetLindormInstanceResponseBody {
	s.DiskThreshold = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetDiskUsage(v string) *GetLindormInstanceResponseBody {
	s.DiskUsage = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetEnableCdc(v bool) *GetLindormInstanceResponseBody {
	s.EnableCdc = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetEnableCompute(v bool) *GetLindormInstanceResponseBody {
	s.EnableCompute = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetEnableKms(v bool) *GetLindormInstanceResponseBody {
	s.EnableKms = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetEnableSSL(v bool) *GetLindormInstanceResponseBody {
	s.EnableSSL = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetEnableShs(v bool) *GetLindormInstanceResponseBody {
	s.EnableShs = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetEnableStream(v bool) *GetLindormInstanceResponseBody {
	s.EnableStream = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetEngineList(v []*GetLindormInstanceResponseBodyEngineList) *GetLindormInstanceResponseBody {
	s.EngineList = v
	return s
}

func (s *GetLindormInstanceResponseBody) SetEngineType(v int32) *GetLindormInstanceResponseBody {
	s.EngineType = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetExpireTime(v string) *GetLindormInstanceResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetExpiredMilliseconds(v int64) *GetLindormInstanceResponseBody {
	s.ExpiredMilliseconds = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetInstanceAlias(v string) *GetLindormInstanceResponseBody {
	s.InstanceAlias = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetInstanceId(v string) *GetLindormInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetInstanceStatus(v string) *GetLindormInstanceResponseBody {
	s.InstanceStatus = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetInstanceStorage(v string) *GetLindormInstanceResponseBody {
	s.InstanceStorage = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetMaintainEndTime(v string) *GetLindormInstanceResponseBody {
	s.MaintainEndTime = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetMaintainStartTime(v string) *GetLindormInstanceResponseBody {
	s.MaintainStartTime = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetNetworkType(v string) *GetLindormInstanceResponseBody {
	s.NetworkType = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetPayType(v string) *GetLindormInstanceResponseBody {
	s.PayType = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetRegionId(v string) *GetLindormInstanceResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetRequestId(v string) *GetLindormInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetServiceType(v string) *GetLindormInstanceResponseBody {
	s.ServiceType = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetVpcId(v string) *GetLindormInstanceResponseBody {
	s.VpcId = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetVswitchId(v string) *GetLindormInstanceResponseBody {
	s.VswitchId = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetZoneId(v string) *GetLindormInstanceResponseBody {
	s.ZoneId = &v
	return s
}

type GetLindormInstanceResponseBodyEngineList struct {
	CoreCount     *string `json:"CoreCount,omitempty" xml:"CoreCount,omitempty"`
	CpuCount      *string `json:"CpuCount,omitempty" xml:"CpuCount,omitempty"`
	Engine        *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	IsLastVersion *bool   `json:"IsLastVersion,omitempty" xml:"IsLastVersion,omitempty"`
	LatestVersion *string `json:"LatestVersion,omitempty" xml:"LatestVersion,omitempty"`
	MemorySize    *string `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetLindormInstanceResponseBodyEngineList) String() string {
	return tea.Prettify(s)
}

func (s GetLindormInstanceResponseBodyEngineList) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceResponseBodyEngineList) SetCoreCount(v string) *GetLindormInstanceResponseBodyEngineList {
	s.CoreCount = &v
	return s
}

func (s *GetLindormInstanceResponseBodyEngineList) SetCpuCount(v string) *GetLindormInstanceResponseBodyEngineList {
	s.CpuCount = &v
	return s
}

func (s *GetLindormInstanceResponseBodyEngineList) SetEngine(v string) *GetLindormInstanceResponseBodyEngineList {
	s.Engine = &v
	return s
}

func (s *GetLindormInstanceResponseBodyEngineList) SetIsLastVersion(v bool) *GetLindormInstanceResponseBodyEngineList {
	s.IsLastVersion = &v
	return s
}

func (s *GetLindormInstanceResponseBodyEngineList) SetLatestVersion(v string) *GetLindormInstanceResponseBodyEngineList {
	s.LatestVersion = &v
	return s
}

func (s *GetLindormInstanceResponseBodyEngineList) SetMemorySize(v string) *GetLindormInstanceResponseBodyEngineList {
	s.MemorySize = &v
	return s
}

func (s *GetLindormInstanceResponseBodyEngineList) SetVersion(v string) *GetLindormInstanceResponseBodyEngineList {
	s.Version = &v
	return s
}

type GetLindormInstanceResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetLindormInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLindormInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLindormInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceResponse) SetHeaders(v map[string]*string) *GetLindormInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetLindormInstanceResponse) SetBody(v *GetLindormInstanceResponseBody) *GetLindormInstanceResponse {
	s.Body = v
	return s
}

type GetLindormInstanceEngineListRequest struct {
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetLindormInstanceEngineListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLindormInstanceEngineListRequest) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceEngineListRequest) SetInstanceId(v string) *GetLindormInstanceEngineListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetLindormInstanceEngineListRequest) SetOwnerAccount(v string) *GetLindormInstanceEngineListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetLindormInstanceEngineListRequest) SetOwnerId(v int64) *GetLindormInstanceEngineListRequest {
	s.OwnerId = &v
	return s
}

func (s *GetLindormInstanceEngineListRequest) SetRegionId(v string) *GetLindormInstanceEngineListRequest {
	s.RegionId = &v
	return s
}

func (s *GetLindormInstanceEngineListRequest) SetResourceOwnerAccount(v string) *GetLindormInstanceEngineListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetLindormInstanceEngineListRequest) SetResourceOwnerId(v int64) *GetLindormInstanceEngineListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetLindormInstanceEngineListRequest) SetSecurityToken(v string) *GetLindormInstanceEngineListRequest {
	s.SecurityToken = &v
	return s
}

type GetLindormInstanceEngineListResponseBody struct {
	EngineList []*GetLindormInstanceEngineListResponseBodyEngineList `json:"EngineList,omitempty" xml:"EngineList,omitempty" type:"Repeated"`
	InstanceId *string                                               `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RequestId  *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetLindormInstanceEngineListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLindormInstanceEngineListResponseBody) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceEngineListResponseBody) SetEngineList(v []*GetLindormInstanceEngineListResponseBodyEngineList) *GetLindormInstanceEngineListResponseBody {
	s.EngineList = v
	return s
}

func (s *GetLindormInstanceEngineListResponseBody) SetInstanceId(v string) *GetLindormInstanceEngineListResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetLindormInstanceEngineListResponseBody) SetRequestId(v string) *GetLindormInstanceEngineListResponseBody {
	s.RequestId = &v
	return s
}

type GetLindormInstanceEngineListResponseBodyEngineList struct {
	EngineType  *string                                                          `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	NetInfoList []*GetLindormInstanceEngineListResponseBodyEngineListNetInfoList `json:"NetInfoList,omitempty" xml:"NetInfoList,omitempty" type:"Repeated"`
}

func (s GetLindormInstanceEngineListResponseBodyEngineList) String() string {
	return tea.Prettify(s)
}

func (s GetLindormInstanceEngineListResponseBodyEngineList) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceEngineListResponseBodyEngineList) SetEngineType(v string) *GetLindormInstanceEngineListResponseBodyEngineList {
	s.EngineType = &v
	return s
}

func (s *GetLindormInstanceEngineListResponseBodyEngineList) SetNetInfoList(v []*GetLindormInstanceEngineListResponseBodyEngineListNetInfoList) *GetLindormInstanceEngineListResponseBodyEngineList {
	s.NetInfoList = v
	return s
}

type GetLindormInstanceEngineListResponseBodyEngineListNetInfoList struct {
	AccessType       *int32  `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	NetType          *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	Port             *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s GetLindormInstanceEngineListResponseBodyEngineListNetInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetLindormInstanceEngineListResponseBodyEngineListNetInfoList) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceEngineListResponseBodyEngineListNetInfoList) SetAccessType(v int32) *GetLindormInstanceEngineListResponseBodyEngineListNetInfoList {
	s.AccessType = &v
	return s
}

func (s *GetLindormInstanceEngineListResponseBodyEngineListNetInfoList) SetConnectionString(v string) *GetLindormInstanceEngineListResponseBodyEngineListNetInfoList {
	s.ConnectionString = &v
	return s
}

func (s *GetLindormInstanceEngineListResponseBodyEngineListNetInfoList) SetNetType(v string) *GetLindormInstanceEngineListResponseBodyEngineListNetInfoList {
	s.NetType = &v
	return s
}

func (s *GetLindormInstanceEngineListResponseBodyEngineListNetInfoList) SetPort(v int32) *GetLindormInstanceEngineListResponseBodyEngineListNetInfoList {
	s.Port = &v
	return s
}

type GetLindormInstanceEngineListResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetLindormInstanceEngineListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLindormInstanceEngineListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLindormInstanceEngineListResponse) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceEngineListResponse) SetHeaders(v map[string]*string) *GetLindormInstanceEngineListResponse {
	s.Headers = v
	return s
}

func (s *GetLindormInstanceEngineListResponse) SetBody(v *GetLindormInstanceEngineListResponseBody) *GetLindormInstanceEngineListResponse {
	s.Body = v
	return s
}

type GetLindormInstanceListRequest struct {
	OwnerAccount         *string                             `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64                              `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueryStr             *string                             `json:"QueryStr,omitempty" xml:"QueryStr,omitempty"`
	RegionId             *string                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string                             `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                              `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string                             `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	ServiceType          *string                             `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	SupportEngine        *int32                              `json:"SupportEngine,omitempty" xml:"SupportEngine,omitempty"`
	Tag                  []*GetLindormInstanceListRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s GetLindormInstanceListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLindormInstanceListRequest) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceListRequest) SetOwnerAccount(v string) *GetLindormInstanceListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetLindormInstanceListRequest) SetOwnerId(v int64) *GetLindormInstanceListRequest {
	s.OwnerId = &v
	return s
}

func (s *GetLindormInstanceListRequest) SetPageNumber(v int32) *GetLindormInstanceListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetLindormInstanceListRequest) SetPageSize(v int32) *GetLindormInstanceListRequest {
	s.PageSize = &v
	return s
}

func (s *GetLindormInstanceListRequest) SetQueryStr(v string) *GetLindormInstanceListRequest {
	s.QueryStr = &v
	return s
}

func (s *GetLindormInstanceListRequest) SetRegionId(v string) *GetLindormInstanceListRequest {
	s.RegionId = &v
	return s
}

func (s *GetLindormInstanceListRequest) SetResourceOwnerAccount(v string) *GetLindormInstanceListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetLindormInstanceListRequest) SetResourceOwnerId(v int64) *GetLindormInstanceListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetLindormInstanceListRequest) SetSecurityToken(v string) *GetLindormInstanceListRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetLindormInstanceListRequest) SetServiceType(v string) *GetLindormInstanceListRequest {
	s.ServiceType = &v
	return s
}

func (s *GetLindormInstanceListRequest) SetSupportEngine(v int32) *GetLindormInstanceListRequest {
	s.SupportEngine = &v
	return s
}

func (s *GetLindormInstanceListRequest) SetTag(v []*GetLindormInstanceListRequestTag) *GetLindormInstanceListRequest {
	s.Tag = v
	return s
}

type GetLindormInstanceListRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetLindormInstanceListRequestTag) String() string {
	return tea.Prettify(s)
}

func (s GetLindormInstanceListRequestTag) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceListRequestTag) SetKey(v string) *GetLindormInstanceListRequestTag {
	s.Key = &v
	return s
}

func (s *GetLindormInstanceListRequestTag) SetValue(v string) *GetLindormInstanceListRequestTag {
	s.Value = &v
	return s
}

type GetLindormInstanceListResponseBody struct {
	InstanceList []*GetLindormInstanceListResponseBodyInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
	PageNumber   *int32                                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId    *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total        *int32                                            `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetLindormInstanceListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLindormInstanceListResponseBody) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceListResponseBody) SetInstanceList(v []*GetLindormInstanceListResponseBodyInstanceList) *GetLindormInstanceListResponseBody {
	s.InstanceList = v
	return s
}

func (s *GetLindormInstanceListResponseBody) SetPageNumber(v int32) *GetLindormInstanceListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetLindormInstanceListResponseBody) SetPageSize(v int32) *GetLindormInstanceListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetLindormInstanceListResponseBody) SetRequestId(v string) *GetLindormInstanceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLindormInstanceListResponseBody) SetTotal(v int32) *GetLindormInstanceListResponseBody {
	s.Total = &v
	return s
}

type GetLindormInstanceListResponseBodyInstanceList struct {
	AliUid              *int64                                                `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	CreateMilliseconds  *int64                                                `json:"CreateMilliseconds,omitempty" xml:"CreateMilliseconds,omitempty"`
	CreateTime          *string                                               `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EnableStream        *bool                                                 `json:"EnableStream,omitempty" xml:"EnableStream,omitempty"`
	EngineType          *string                                               `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	ExpireTime          *string                                               `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	ExpiredMilliseconds *int64                                                `json:"ExpiredMilliseconds,omitempty" xml:"ExpiredMilliseconds,omitempty"`
	InstanceAlias       *string                                               `json:"InstanceAlias,omitempty" xml:"InstanceAlias,omitempty"`
	InstanceId          *string                                               `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceStatus      *string                                               `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	InstanceStorage     *string                                               `json:"InstanceStorage,omitempty" xml:"InstanceStorage,omitempty"`
	NetworkType         *string                                               `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	PayType             *string                                               `json:"PayType,omitempty" xml:"PayType,omitempty"`
	RegionId            *string                                               `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceType         *string                                               `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	Tags                []*GetLindormInstanceListResponseBodyInstanceListTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	VpcId               *string                                               `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId              *string                                               `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetLindormInstanceListResponseBodyInstanceList) String() string {
	return tea.Prettify(s)
}

func (s GetLindormInstanceListResponseBodyInstanceList) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetAliUid(v int64) *GetLindormInstanceListResponseBodyInstanceList {
	s.AliUid = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetCreateMilliseconds(v int64) *GetLindormInstanceListResponseBodyInstanceList {
	s.CreateMilliseconds = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetCreateTime(v string) *GetLindormInstanceListResponseBodyInstanceList {
	s.CreateTime = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetEnableStream(v bool) *GetLindormInstanceListResponseBodyInstanceList {
	s.EnableStream = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetEngineType(v string) *GetLindormInstanceListResponseBodyInstanceList {
	s.EngineType = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetExpireTime(v string) *GetLindormInstanceListResponseBodyInstanceList {
	s.ExpireTime = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetExpiredMilliseconds(v int64) *GetLindormInstanceListResponseBodyInstanceList {
	s.ExpiredMilliseconds = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetInstanceAlias(v string) *GetLindormInstanceListResponseBodyInstanceList {
	s.InstanceAlias = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetInstanceId(v string) *GetLindormInstanceListResponseBodyInstanceList {
	s.InstanceId = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetInstanceStatus(v string) *GetLindormInstanceListResponseBodyInstanceList {
	s.InstanceStatus = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetInstanceStorage(v string) *GetLindormInstanceListResponseBodyInstanceList {
	s.InstanceStorage = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetNetworkType(v string) *GetLindormInstanceListResponseBodyInstanceList {
	s.NetworkType = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetPayType(v string) *GetLindormInstanceListResponseBodyInstanceList {
	s.PayType = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetRegionId(v string) *GetLindormInstanceListResponseBodyInstanceList {
	s.RegionId = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetServiceType(v string) *GetLindormInstanceListResponseBodyInstanceList {
	s.ServiceType = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetTags(v []*GetLindormInstanceListResponseBodyInstanceListTags) *GetLindormInstanceListResponseBodyInstanceList {
	s.Tags = v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetVpcId(v string) *GetLindormInstanceListResponseBodyInstanceList {
	s.VpcId = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetZoneId(v string) *GetLindormInstanceListResponseBodyInstanceList {
	s.ZoneId = &v
	return s
}

type GetLindormInstanceListResponseBodyInstanceListTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetLindormInstanceListResponseBodyInstanceListTags) String() string {
	return tea.Prettify(s)
}

func (s GetLindormInstanceListResponseBodyInstanceListTags) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceListResponseBodyInstanceListTags) SetKey(v string) *GetLindormInstanceListResponseBodyInstanceListTags {
	s.Key = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceListTags) SetValue(v string) *GetLindormInstanceListResponseBodyInstanceListTags {
	s.Value = &v
	return s
}

type GetLindormInstanceListResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetLindormInstanceListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLindormInstanceListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLindormInstanceListResponse) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceListResponse) SetHeaders(v map[string]*string) *GetLindormInstanceListResponse {
	s.Headers = v
	return s
}

func (s *GetLindormInstanceListResponse) SetBody(v *GetLindormInstanceListResponseBody) *GetLindormInstanceListResponse {
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
	SecurityToken        *string                       `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
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

func (s *ListTagResourcesRequest) SetSecurityToken(v string) *ListTagResourcesRequest {
	s.SecurityToken = &v
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
	NextToken    *string                                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagResources []*ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
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

func (s *ListTagResourcesResponseBody) SetTagResources(v []*ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceId(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagKey(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagValue(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagValue = &v
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

type ReleaseLindormInstanceRequest struct {
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ReleaseLindormInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseLindormInstanceRequest) GoString() string {
	return s.String()
}

func (s *ReleaseLindormInstanceRequest) SetInstanceId(v string) *ReleaseLindormInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ReleaseLindormInstanceRequest) SetOwnerAccount(v string) *ReleaseLindormInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ReleaseLindormInstanceRequest) SetOwnerId(v int64) *ReleaseLindormInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *ReleaseLindormInstanceRequest) SetResourceOwnerAccount(v string) *ReleaseLindormInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReleaseLindormInstanceRequest) SetResourceOwnerId(v int64) *ReleaseLindormInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReleaseLindormInstanceRequest) SetSecurityToken(v string) *ReleaseLindormInstanceRequest {
	s.SecurityToken = &v
	return s
}

type ReleaseLindormInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseLindormInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseLindormInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseLindormInstanceResponseBody) SetRequestId(v string) *ReleaseLindormInstanceResponseBody {
	s.RequestId = &v
	return s
}

type ReleaseLindormInstanceResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReleaseLindormInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReleaseLindormInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseLindormInstanceResponse) GoString() string {
	return s.String()
}

func (s *ReleaseLindormInstanceResponse) SetHeaders(v map[string]*string) *ReleaseLindormInstanceResponse {
	s.Headers = v
	return s
}

func (s *ReleaseLindormInstanceResponse) SetBody(v *ReleaseLindormInstanceResponseBody) *ReleaseLindormInstanceResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	OwnerAccount         *string                   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64                    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceId           []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string                   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceType         *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	SecurityToken        *string                   `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
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

func (s *TagResourcesRequest) SetSecurityToken(v string) *TagResourcesRequest {
	s.SecurityToken = &v
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

type UntagResourcesRequest struct {
	All                  *bool     `json:"All,omitempty" xml:"All,omitempty"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceType         *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	SecurityToken        *string   `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
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

func (s *UntagResourcesRequest) SetSecurityToken(v string) *UntagResourcesRequest {
	s.SecurityToken = &v
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type UpdateInstanceIpWhiteListRequest struct {
	GroupName            *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityIpList       *string `json:"SecurityIpList,omitempty" xml:"SecurityIpList,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s UpdateInstanceIpWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceIpWhiteListRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceIpWhiteListRequest) SetGroupName(v string) *UpdateInstanceIpWhiteListRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateInstanceIpWhiteListRequest) SetInstanceId(v string) *UpdateInstanceIpWhiteListRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceIpWhiteListRequest) SetOwnerAccount(v string) *UpdateInstanceIpWhiteListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateInstanceIpWhiteListRequest) SetOwnerId(v int64) *UpdateInstanceIpWhiteListRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateInstanceIpWhiteListRequest) SetResourceOwnerAccount(v string) *UpdateInstanceIpWhiteListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateInstanceIpWhiteListRequest) SetResourceOwnerId(v int64) *UpdateInstanceIpWhiteListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateInstanceIpWhiteListRequest) SetSecurityIpList(v string) *UpdateInstanceIpWhiteListRequest {
	s.SecurityIpList = &v
	return s
}

func (s *UpdateInstanceIpWhiteListRequest) SetSecurityToken(v string) *UpdateInstanceIpWhiteListRequest {
	s.SecurityToken = &v
	return s
}

type UpdateInstanceIpWhiteListResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateInstanceIpWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceIpWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceIpWhiteListResponseBody) SetRequestId(v string) *UpdateInstanceIpWhiteListResponseBody {
	s.RequestId = &v
	return s
}

type UpdateInstanceIpWhiteListResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateInstanceIpWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateInstanceIpWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceIpWhiteListResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceIpWhiteListResponse) SetHeaders(v map[string]*string) *UpdateInstanceIpWhiteListResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceIpWhiteListResponse) SetBody(v *UpdateInstanceIpWhiteListResponseBody) *UpdateInstanceIpWhiteListResponse {
	s.Body = v
	return s
}

type UpgradeLindormInstanceRequest struct {
	ClusterStorage       *int32  `json:"ClusterStorage,omitempty" xml:"ClusterStorage,omitempty"`
	ColdStorage          *int32  `json:"ColdStorage,omitempty" xml:"ColdStorage,omitempty"`
	CoreNum              *int32  `json:"CoreNum,omitempty" xml:"CoreNum,omitempty"`
	CoreSpec             *string `json:"CoreSpec,omitempty" xml:"CoreSpec,omitempty"`
	FilestoreNum         *int32  `json:"FilestoreNum,omitempty" xml:"FilestoreNum,omitempty"`
	FilestoreSpec        *string `json:"FilestoreSpec,omitempty" xml:"FilestoreSpec,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LindormNum           *int32  `json:"LindormNum,omitempty" xml:"LindormNum,omitempty"`
	LindormSpec          *string `json:"LindormSpec,omitempty" xml:"LindormSpec,omitempty"`
	LtsCoreNum           *int32  `json:"LtsCoreNum,omitempty" xml:"LtsCoreNum,omitempty"`
	LtsCoreSpec          *string `json:"LtsCoreSpec,omitempty" xml:"LtsCoreSpec,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PhoenixCoreNum       *int32  `json:"PhoenixCoreNum,omitempty" xml:"PhoenixCoreNum,omitempty"`
	PhoenixCoreSpec      *string `json:"PhoenixCoreSpec,omitempty" xml:"PhoenixCoreSpec,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	SolrNum              *int32  `json:"SolrNum,omitempty" xml:"SolrNum,omitempty"`
	SolrSpec             *string `json:"SolrSpec,omitempty" xml:"SolrSpec,omitempty"`
	TsdbNum              *int32  `json:"TsdbNum,omitempty" xml:"TsdbNum,omitempty"`
	TsdbSpec             *string `json:"TsdbSpec,omitempty" xml:"TsdbSpec,omitempty"`
	UpgradeType          *string `json:"UpgradeType,omitempty" xml:"UpgradeType,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s UpgradeLindormInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeLindormInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpgradeLindormInstanceRequest) SetClusterStorage(v int32) *UpgradeLindormInstanceRequest {
	s.ClusterStorage = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetColdStorage(v int32) *UpgradeLindormInstanceRequest {
	s.ColdStorage = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetCoreNum(v int32) *UpgradeLindormInstanceRequest {
	s.CoreNum = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetCoreSpec(v string) *UpgradeLindormInstanceRequest {
	s.CoreSpec = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetFilestoreNum(v int32) *UpgradeLindormInstanceRequest {
	s.FilestoreNum = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetFilestoreSpec(v string) *UpgradeLindormInstanceRequest {
	s.FilestoreSpec = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetInstanceId(v string) *UpgradeLindormInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetLindormNum(v int32) *UpgradeLindormInstanceRequest {
	s.LindormNum = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetLindormSpec(v string) *UpgradeLindormInstanceRequest {
	s.LindormSpec = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetLtsCoreNum(v int32) *UpgradeLindormInstanceRequest {
	s.LtsCoreNum = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetLtsCoreSpec(v string) *UpgradeLindormInstanceRequest {
	s.LtsCoreSpec = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetOwnerAccount(v string) *UpgradeLindormInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetOwnerId(v int64) *UpgradeLindormInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetPhoenixCoreNum(v int32) *UpgradeLindormInstanceRequest {
	s.PhoenixCoreNum = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetPhoenixCoreSpec(v string) *UpgradeLindormInstanceRequest {
	s.PhoenixCoreSpec = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetRegionId(v string) *UpgradeLindormInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetResourceOwnerAccount(v string) *UpgradeLindormInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetResourceOwnerId(v int64) *UpgradeLindormInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetSecurityToken(v string) *UpgradeLindormInstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetSolrNum(v int32) *UpgradeLindormInstanceRequest {
	s.SolrNum = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetSolrSpec(v string) *UpgradeLindormInstanceRequest {
	s.SolrSpec = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetTsdbNum(v int32) *UpgradeLindormInstanceRequest {
	s.TsdbNum = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetTsdbSpec(v string) *UpgradeLindormInstanceRequest {
	s.TsdbSpec = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetUpgradeType(v string) *UpgradeLindormInstanceRequest {
	s.UpgradeType = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetZoneId(v string) *UpgradeLindormInstanceRequest {
	s.ZoneId = &v
	return s
}

type UpgradeLindormInstanceResponseBody struct {
	OrderId   *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeLindormInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeLindormInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeLindormInstanceResponseBody) SetOrderId(v int64) *UpgradeLindormInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *UpgradeLindormInstanceResponseBody) SetRequestId(v string) *UpgradeLindormInstanceResponseBody {
	s.RequestId = &v
	return s
}

type UpgradeLindormInstanceResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpgradeLindormInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpgradeLindormInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeLindormInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpgradeLindormInstanceResponse) SetHeaders(v map[string]*string) *UpgradeLindormInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpgradeLindormInstanceResponse) SetBody(v *UpgradeLindormInstanceResponseBody) *UpgradeLindormInstanceResponse {
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

func (client *Client) CreateLindormInstanceWithOptions(request *CreateLindormInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateLindormInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ColdStorage)) {
		query["ColdStorage"] = request.ColdStorage
	}

	if !tea.BoolValue(util.IsUnset(request.CoreSpec)) {
		query["CoreSpec"] = request.CoreSpec
	}

	if !tea.BoolValue(util.IsUnset(request.DiskCategory)) {
		query["DiskCategory"] = request.DiskCategory
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.FilestoreNum)) {
		query["FilestoreNum"] = request.FilestoreNum
	}

	if !tea.BoolValue(util.IsUnset(request.FilestoreSpec)) {
		query["FilestoreSpec"] = request.FilestoreSpec
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceAlias)) {
		query["InstanceAlias"] = request.InstanceAlias
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceStorage)) {
		query["InstanceStorage"] = request.InstanceStorage
	}

	if !tea.BoolValue(util.IsUnset(request.LindormNum)) {
		query["LindormNum"] = request.LindormNum
	}

	if !tea.BoolValue(util.IsUnset(request.LindormSpec)) {
		query["LindormSpec"] = request.LindormSpec
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
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

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.SolrNum)) {
		query["SolrNum"] = request.SolrNum
	}

	if !tea.BoolValue(util.IsUnset(request.SolrSpec)) {
		query["SolrSpec"] = request.SolrSpec
	}

	if !tea.BoolValue(util.IsUnset(request.TsdbNum)) {
		query["TsdbNum"] = request.TsdbNum
	}

	if !tea.BoolValue(util.IsUnset(request.TsdbSpec)) {
		query["TsdbSpec"] = request.TsdbSpec
	}

	if !tea.BoolValue(util.IsUnset(request.VPCId)) {
		query["VPCId"] = request.VPCId
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLindormInstance"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLindormInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLindormInstance(request *CreateLindormInstanceRequest) (_result *CreateLindormInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLindormInstanceResponse{}
	_body, _err := client.CreateLindormInstanceWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2020-06-15"),
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

func (client *Client) GetInstanceIpWhiteListWithOptions(request *GetInstanceIpWhiteListRequest, runtime *util.RuntimeOptions) (_result *GetInstanceIpWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstanceIpWhiteList"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstanceIpWhiteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInstanceIpWhiteList(request *GetInstanceIpWhiteListRequest) (_result *GetInstanceIpWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInstanceIpWhiteListResponse{}
	_body, _err := client.GetInstanceIpWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLindormInstanceWithOptions(request *GetLindormInstanceRequest, runtime *util.RuntimeOptions) (_result *GetLindormInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
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

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLindormInstance"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLindormInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLindormInstance(request *GetLindormInstanceRequest) (_result *GetLindormInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLindormInstanceResponse{}
	_body, _err := client.GetLindormInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLindormInstanceEngineListWithOptions(request *GetLindormInstanceEngineListRequest, runtime *util.RuntimeOptions) (_result *GetLindormInstanceEngineListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
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

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLindormInstanceEngineList"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLindormInstanceEngineListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLindormInstanceEngineList(request *GetLindormInstanceEngineListRequest) (_result *GetLindormInstanceEngineListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLindormInstanceEngineListResponse{}
	_body, _err := client.GetLindormInstanceEngineListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLindormInstanceListWithOptions(request *GetLindormInstanceListRequest, runtime *util.RuntimeOptions) (_result *GetLindormInstanceListResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryStr)) {
		query["QueryStr"] = request.QueryStr
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

	if !tea.BoolValue(util.IsUnset(request.ServiceType)) {
		query["ServiceType"] = request.ServiceType
	}

	if !tea.BoolValue(util.IsUnset(request.SupportEngine)) {
		query["SupportEngine"] = request.SupportEngine
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLindormInstanceList"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLindormInstanceListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLindormInstanceList(request *GetLindormInstanceListRequest) (_result *GetLindormInstanceListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLindormInstanceListResponse{}
	_body, _err := client.GetLindormInstanceListWithOptions(request, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2020-06-15"),
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

func (client *Client) ReleaseLindormInstanceWithOptions(request *ReleaseLindormInstanceRequest, runtime *util.RuntimeOptions) (_result *ReleaseLindormInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleaseLindormInstance"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleaseLindormInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseLindormInstance(request *ReleaseLindormInstanceRequest) (_result *ReleaseLindormInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseLindormInstanceResponse{}
	_body, _err := client.ReleaseLindormInstanceWithOptions(request, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2020-06-15"),
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

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2020-06-15"),
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

func (client *Client) UpdateInstanceIpWhiteListWithOptions(request *UpdateInstanceIpWhiteListRequest, runtime *util.RuntimeOptions) (_result *UpdateInstanceIpWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

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

	if !tea.BoolValue(util.IsUnset(request.SecurityIpList)) {
		query["SecurityIpList"] = request.SecurityIpList
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInstanceIpWhiteList"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateInstanceIpWhiteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateInstanceIpWhiteList(request *UpdateInstanceIpWhiteListRequest) (_result *UpdateInstanceIpWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateInstanceIpWhiteListResponse{}
	_body, _err := client.UpdateInstanceIpWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeLindormInstanceWithOptions(request *UpgradeLindormInstanceRequest, runtime *util.RuntimeOptions) (_result *UpgradeLindormInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterStorage)) {
		query["ClusterStorage"] = request.ClusterStorage
	}

	if !tea.BoolValue(util.IsUnset(request.ColdStorage)) {
		query["ColdStorage"] = request.ColdStorage
	}

	if !tea.BoolValue(util.IsUnset(request.CoreNum)) {
		query["CoreNum"] = request.CoreNum
	}

	if !tea.BoolValue(util.IsUnset(request.CoreSpec)) {
		query["CoreSpec"] = request.CoreSpec
	}

	if !tea.BoolValue(util.IsUnset(request.FilestoreNum)) {
		query["FilestoreNum"] = request.FilestoreNum
	}

	if !tea.BoolValue(util.IsUnset(request.FilestoreSpec)) {
		query["FilestoreSpec"] = request.FilestoreSpec
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.LindormNum)) {
		query["LindormNum"] = request.LindormNum
	}

	if !tea.BoolValue(util.IsUnset(request.LindormSpec)) {
		query["LindormSpec"] = request.LindormSpec
	}

	if !tea.BoolValue(util.IsUnset(request.LtsCoreNum)) {
		query["LtsCoreNum"] = request.LtsCoreNum
	}

	if !tea.BoolValue(util.IsUnset(request.LtsCoreSpec)) {
		query["LtsCoreSpec"] = request.LtsCoreSpec
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoenixCoreNum)) {
		query["PhoenixCoreNum"] = request.PhoenixCoreNum
	}

	if !tea.BoolValue(util.IsUnset(request.PhoenixCoreSpec)) {
		query["PhoenixCoreSpec"] = request.PhoenixCoreSpec
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

	if !tea.BoolValue(util.IsUnset(request.SolrNum)) {
		query["SolrNum"] = request.SolrNum
	}

	if !tea.BoolValue(util.IsUnset(request.SolrSpec)) {
		query["SolrSpec"] = request.SolrSpec
	}

	if !tea.BoolValue(util.IsUnset(request.TsdbNum)) {
		query["TsdbNum"] = request.TsdbNum
	}

	if !tea.BoolValue(util.IsUnset(request.TsdbSpec)) {
		query["TsdbSpec"] = request.TsdbSpec
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
		Action:      tea.String("UpgradeLindormInstance"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradeLindormInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeLindormInstance(request *UpgradeLindormInstanceRequest) (_result *UpgradeLindormInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeLindormInstanceResponse{}
	_body, _err := client.UpgradeLindormInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
