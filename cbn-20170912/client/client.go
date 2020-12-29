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

type ActiveFlowLogRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	FlowLogId            *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
}

func (s ActiveFlowLogRequest) String() string {
	return tea.Prettify(s)
}

func (s ActiveFlowLogRequest) GoString() string {
	return s.String()
}

func (s *ActiveFlowLogRequest) SetOwnerAccount(v string) *ActiveFlowLogRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ActiveFlowLogRequest) SetOwnerId(v int64) *ActiveFlowLogRequest {
	s.OwnerId = &v
	return s
}

func (s *ActiveFlowLogRequest) SetResourceOwnerAccount(v string) *ActiveFlowLogRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ActiveFlowLogRequest) SetResourceOwnerId(v int64) *ActiveFlowLogRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ActiveFlowLogRequest) SetClientToken(v string) *ActiveFlowLogRequest {
	s.ClientToken = &v
	return s
}

func (s *ActiveFlowLogRequest) SetRegionId(v string) *ActiveFlowLogRequest {
	s.RegionId = &v
	return s
}

func (s *ActiveFlowLogRequest) SetFlowLogId(v string) *ActiveFlowLogRequest {
	s.FlowLogId = &v
	return s
}

func (s *ActiveFlowLogRequest) SetCenId(v string) *ActiveFlowLogRequest {
	s.CenId = &v
	return s
}

type ActiveFlowLogResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ActiveFlowLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ActiveFlowLogResponseBody) GoString() string {
	return s.String()
}

func (s *ActiveFlowLogResponseBody) SetRequestId(v string) *ActiveFlowLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *ActiveFlowLogResponseBody) SetSuccess(v string) *ActiveFlowLogResponseBody {
	s.Success = &v
	return s
}

type ActiveFlowLogResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ActiveFlowLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ActiveFlowLogResponse) String() string {
	return tea.Prettify(s)
}

func (s ActiveFlowLogResponse) GoString() string {
	return s.String()
}

func (s *ActiveFlowLogResponse) SetHeaders(v map[string]*string) *ActiveFlowLogResponse {
	s.Headers = v
	return s
}

func (s *ActiveFlowLogResponse) SetBody(v *ActiveFlowLogResponseBody) *ActiveFlowLogResponse {
	s.Body = v
	return s
}

type AssociateCenBandwidthPackageRequest struct {
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	CenId                 *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	CenBandwidthPackageId *string `json:"CenBandwidthPackageId,omitempty" xml:"CenBandwidthPackageId,omitempty"`
}

func (s AssociateCenBandwidthPackageRequest) String() string {
	return tea.Prettify(s)
}

func (s AssociateCenBandwidthPackageRequest) GoString() string {
	return s.String()
}

func (s *AssociateCenBandwidthPackageRequest) SetOwnerAccount(v string) *AssociateCenBandwidthPackageRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AssociateCenBandwidthPackageRequest) SetOwnerId(v int64) *AssociateCenBandwidthPackageRequest {
	s.OwnerId = &v
	return s
}

func (s *AssociateCenBandwidthPackageRequest) SetResourceOwnerAccount(v string) *AssociateCenBandwidthPackageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AssociateCenBandwidthPackageRequest) SetResourceOwnerId(v int64) *AssociateCenBandwidthPackageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AssociateCenBandwidthPackageRequest) SetCenId(v string) *AssociateCenBandwidthPackageRequest {
	s.CenId = &v
	return s
}

func (s *AssociateCenBandwidthPackageRequest) SetCenBandwidthPackageId(v string) *AssociateCenBandwidthPackageRequest {
	s.CenBandwidthPackageId = &v
	return s
}

type AssociateCenBandwidthPackageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateCenBandwidthPackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssociateCenBandwidthPackageResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateCenBandwidthPackageResponseBody) SetRequestId(v string) *AssociateCenBandwidthPackageResponseBody {
	s.RequestId = &v
	return s
}

type AssociateCenBandwidthPackageResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AssociateCenBandwidthPackageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AssociateCenBandwidthPackageResponse) String() string {
	return tea.Prettify(s)
}

func (s AssociateCenBandwidthPackageResponse) GoString() string {
	return s.String()
}

func (s *AssociateCenBandwidthPackageResponse) SetHeaders(v map[string]*string) *AssociateCenBandwidthPackageResponse {
	s.Headers = v
	return s
}

func (s *AssociateCenBandwidthPackageResponse) SetBody(v *AssociateCenBandwidthPackageResponseBody) *AssociateCenBandwidthPackageResponse {
	s.Body = v
	return s
}

type AttachCenChildInstanceRequest struct {
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	CenId                 *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ChildInstanceId       *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	ChildInstanceType     *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	ChildInstanceRegionId *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	ChildInstanceOwnerId  *int64  `json:"ChildInstanceOwnerId,omitempty" xml:"ChildInstanceOwnerId,omitempty"`
}

func (s AttachCenChildInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachCenChildInstanceRequest) GoString() string {
	return s.String()
}

func (s *AttachCenChildInstanceRequest) SetOwnerAccount(v string) *AttachCenChildInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AttachCenChildInstanceRequest) SetOwnerId(v int64) *AttachCenChildInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *AttachCenChildInstanceRequest) SetResourceOwnerAccount(v string) *AttachCenChildInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AttachCenChildInstanceRequest) SetResourceOwnerId(v int64) *AttachCenChildInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AttachCenChildInstanceRequest) SetCenId(v string) *AttachCenChildInstanceRequest {
	s.CenId = &v
	return s
}

func (s *AttachCenChildInstanceRequest) SetChildInstanceId(v string) *AttachCenChildInstanceRequest {
	s.ChildInstanceId = &v
	return s
}

func (s *AttachCenChildInstanceRequest) SetChildInstanceType(v string) *AttachCenChildInstanceRequest {
	s.ChildInstanceType = &v
	return s
}

func (s *AttachCenChildInstanceRequest) SetChildInstanceRegionId(v string) *AttachCenChildInstanceRequest {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *AttachCenChildInstanceRequest) SetChildInstanceOwnerId(v int64) *AttachCenChildInstanceRequest {
	s.ChildInstanceOwnerId = &v
	return s
}

type AttachCenChildInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachCenChildInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachCenChildInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *AttachCenChildInstanceResponseBody) SetRequestId(v string) *AttachCenChildInstanceResponseBody {
	s.RequestId = &v
	return s
}

type AttachCenChildInstanceResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AttachCenChildInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttachCenChildInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachCenChildInstanceResponse) GoString() string {
	return s.String()
}

func (s *AttachCenChildInstanceResponse) SetHeaders(v map[string]*string) *AttachCenChildInstanceResponse {
	s.Headers = v
	return s
}

func (s *AttachCenChildInstanceResponse) SetBody(v *AttachCenChildInstanceResponseBody) *AttachCenChildInstanceResponse {
	s.Body = v
	return s
}

type CreateCenRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ProtectionLevel      *string `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty"`
	Ipv6Level            *string `json:"Ipv6Level,omitempty" xml:"Ipv6Level,omitempty"`
}

func (s CreateCenRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCenRequest) GoString() string {
	return s.String()
}

func (s *CreateCenRequest) SetOwnerAccount(v string) *CreateCenRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateCenRequest) SetOwnerId(v int64) *CreateCenRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateCenRequest) SetResourceOwnerAccount(v string) *CreateCenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateCenRequest) SetResourceOwnerId(v int64) *CreateCenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateCenRequest) SetClientToken(v string) *CreateCenRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCenRequest) SetName(v string) *CreateCenRequest {
	s.Name = &v
	return s
}

func (s *CreateCenRequest) SetDescription(v string) *CreateCenRequest {
	s.Description = &v
	return s
}

func (s *CreateCenRequest) SetProtectionLevel(v string) *CreateCenRequest {
	s.ProtectionLevel = &v
	return s
}

func (s *CreateCenRequest) SetIpv6Level(v string) *CreateCenRequest {
	s.Ipv6Level = &v
	return s
}

type CreateCenResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CenId     *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
}

func (s CreateCenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCenResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCenResponseBody) SetRequestId(v string) *CreateCenResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCenResponseBody) SetCenId(v string) *CreateCenResponseBody {
	s.CenId = &v
	return s
}

type CreateCenResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCenResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCenResponse) GoString() string {
	return s.String()
}

func (s *CreateCenResponse) SetHeaders(v map[string]*string) *CreateCenResponse {
	s.Headers = v
	return s
}

func (s *CreateCenResponse) SetBody(v *CreateCenResponseBody) *CreateCenResponse {
	s.Body = v
	return s
}

type CreateCenBandwidthPackageRequest struct {
	OwnerAccount               *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount       *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId            *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ClientToken                *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Name                       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description                *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Bandwidth                  *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	GeographicRegionAId        *string `json:"GeographicRegionAId,omitempty" xml:"GeographicRegionAId,omitempty"`
	GeographicRegionBId        *string `json:"GeographicRegionBId,omitempty" xml:"GeographicRegionBId,omitempty"`
	BandwidthPackageChargeType *string `json:"BandwidthPackageChargeType,omitempty" xml:"BandwidthPackageChargeType,omitempty"`
	Period                     *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	PricingCycle               *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	AutoPay                    *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	AutoRenew                  *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AutoRenewDuration          *int32  `json:"AutoRenewDuration,omitempty" xml:"AutoRenewDuration,omitempty"`
}

func (s CreateCenBandwidthPackageRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCenBandwidthPackageRequest) GoString() string {
	return s.String()
}

func (s *CreateCenBandwidthPackageRequest) SetOwnerAccount(v string) *CreateCenBandwidthPackageRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateCenBandwidthPackageRequest) SetOwnerId(v int64) *CreateCenBandwidthPackageRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateCenBandwidthPackageRequest) SetResourceOwnerAccount(v string) *CreateCenBandwidthPackageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateCenBandwidthPackageRequest) SetResourceOwnerId(v int64) *CreateCenBandwidthPackageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateCenBandwidthPackageRequest) SetClientToken(v string) *CreateCenBandwidthPackageRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCenBandwidthPackageRequest) SetName(v string) *CreateCenBandwidthPackageRequest {
	s.Name = &v
	return s
}

func (s *CreateCenBandwidthPackageRequest) SetDescription(v string) *CreateCenBandwidthPackageRequest {
	s.Description = &v
	return s
}

func (s *CreateCenBandwidthPackageRequest) SetBandwidth(v int32) *CreateCenBandwidthPackageRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateCenBandwidthPackageRequest) SetGeographicRegionAId(v string) *CreateCenBandwidthPackageRequest {
	s.GeographicRegionAId = &v
	return s
}

func (s *CreateCenBandwidthPackageRequest) SetGeographicRegionBId(v string) *CreateCenBandwidthPackageRequest {
	s.GeographicRegionBId = &v
	return s
}

func (s *CreateCenBandwidthPackageRequest) SetBandwidthPackageChargeType(v string) *CreateCenBandwidthPackageRequest {
	s.BandwidthPackageChargeType = &v
	return s
}

func (s *CreateCenBandwidthPackageRequest) SetPeriod(v int32) *CreateCenBandwidthPackageRequest {
	s.Period = &v
	return s
}

func (s *CreateCenBandwidthPackageRequest) SetPricingCycle(v string) *CreateCenBandwidthPackageRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateCenBandwidthPackageRequest) SetAutoPay(v bool) *CreateCenBandwidthPackageRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateCenBandwidthPackageRequest) SetAutoRenew(v bool) *CreateCenBandwidthPackageRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateCenBandwidthPackageRequest) SetAutoRenewDuration(v int32) *CreateCenBandwidthPackageRequest {
	s.AutoRenewDuration = &v
	return s
}

type CreateCenBandwidthPackageResponseBody struct {
	RequestId                  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CenBandwidthPackageId      *string `json:"CenBandwidthPackageId,omitempty" xml:"CenBandwidthPackageId,omitempty"`
	CenBandwidthPackageOrderId *string `json:"CenBandwidthPackageOrderId,omitempty" xml:"CenBandwidthPackageOrderId,omitempty"`
}

func (s CreateCenBandwidthPackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCenBandwidthPackageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCenBandwidthPackageResponseBody) SetRequestId(v string) *CreateCenBandwidthPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCenBandwidthPackageResponseBody) SetCenBandwidthPackageId(v string) *CreateCenBandwidthPackageResponseBody {
	s.CenBandwidthPackageId = &v
	return s
}

func (s *CreateCenBandwidthPackageResponseBody) SetCenBandwidthPackageOrderId(v string) *CreateCenBandwidthPackageResponseBody {
	s.CenBandwidthPackageOrderId = &v
	return s
}

type CreateCenBandwidthPackageResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCenBandwidthPackageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCenBandwidthPackageResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCenBandwidthPackageResponse) GoString() string {
	return s.String()
}

func (s *CreateCenBandwidthPackageResponse) SetHeaders(v map[string]*string) *CreateCenBandwidthPackageResponse {
	s.Headers = v
	return s
}

func (s *CreateCenBandwidthPackageResponse) SetBody(v *CreateCenBandwidthPackageResponseBody) *CreateCenBandwidthPackageResponse {
	s.Body = v
	return s
}

type CreateCenChildInstanceRouteEntryToCenRequest struct {
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	CenId                 *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ChildInstanceId       *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	ChildInstanceType     *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	ChildInstanceRegionId *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	ChildInstanceAliUid   *int64  `json:"ChildInstanceAliUid,omitempty" xml:"ChildInstanceAliUid,omitempty"`
	RouteTableId          *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
	DestinationCidrBlock  *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
}

func (s CreateCenChildInstanceRouteEntryToCenRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCenChildInstanceRouteEntryToCenRequest) GoString() string {
	return s.String()
}

func (s *CreateCenChildInstanceRouteEntryToCenRequest) SetOwnerAccount(v string) *CreateCenChildInstanceRouteEntryToCenRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToCenRequest) SetOwnerId(v int64) *CreateCenChildInstanceRouteEntryToCenRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToCenRequest) SetResourceOwnerAccount(v string) *CreateCenChildInstanceRouteEntryToCenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToCenRequest) SetResourceOwnerId(v int64) *CreateCenChildInstanceRouteEntryToCenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToCenRequest) SetCenId(v string) *CreateCenChildInstanceRouteEntryToCenRequest {
	s.CenId = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToCenRequest) SetChildInstanceId(v string) *CreateCenChildInstanceRouteEntryToCenRequest {
	s.ChildInstanceId = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToCenRequest) SetChildInstanceType(v string) *CreateCenChildInstanceRouteEntryToCenRequest {
	s.ChildInstanceType = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToCenRequest) SetChildInstanceRegionId(v string) *CreateCenChildInstanceRouteEntryToCenRequest {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToCenRequest) SetChildInstanceAliUid(v int64) *CreateCenChildInstanceRouteEntryToCenRequest {
	s.ChildInstanceAliUid = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToCenRequest) SetRouteTableId(v string) *CreateCenChildInstanceRouteEntryToCenRequest {
	s.RouteTableId = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToCenRequest) SetDestinationCidrBlock(v string) *CreateCenChildInstanceRouteEntryToCenRequest {
	s.DestinationCidrBlock = &v
	return s
}

type CreateCenChildInstanceRouteEntryToCenResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCenChildInstanceRouteEntryToCenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCenChildInstanceRouteEntryToCenResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCenChildInstanceRouteEntryToCenResponseBody) SetRequestId(v string) *CreateCenChildInstanceRouteEntryToCenResponseBody {
	s.RequestId = &v
	return s
}

type CreateCenChildInstanceRouteEntryToCenResponse struct {
	Headers map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCenChildInstanceRouteEntryToCenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCenChildInstanceRouteEntryToCenResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCenChildInstanceRouteEntryToCenResponse) GoString() string {
	return s.String()
}

func (s *CreateCenChildInstanceRouteEntryToCenResponse) SetHeaders(v map[string]*string) *CreateCenChildInstanceRouteEntryToCenResponse {
	s.Headers = v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToCenResponse) SetBody(v *CreateCenChildInstanceRouteEntryToCenResponseBody) *CreateCenChildInstanceRouteEntryToCenResponse {
	s.Body = v
	return s
}

type CreateCenRouteMapRequest struct {
	OwnerAccount                       *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                            *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount               *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId                    *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	CenId                              *string   `json:"CenId,omitempty" xml:"CenId,omitempty"`
	CenRegionId                        *string   `json:"CenRegionId,omitempty" xml:"CenRegionId,omitempty"`
	TransmitDirection                  *string   `json:"TransmitDirection,omitempty" xml:"TransmitDirection,omitempty"`
	Description                        *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	Priority                           *int32    `json:"Priority,omitempty" xml:"Priority,omitempty"`
	MapResult                          *string   `json:"MapResult,omitempty" xml:"MapResult,omitempty"`
	NextPriority                       *int32    `json:"NextPriority,omitempty" xml:"NextPriority,omitempty"`
	CidrMatchMode                      *string   `json:"CidrMatchMode,omitempty" xml:"CidrMatchMode,omitempty"`
	AsPathMatchMode                    *string   `json:"AsPathMatchMode,omitempty" xml:"AsPathMatchMode,omitempty"`
	CommunityMatchMode                 *string   `json:"CommunityMatchMode,omitempty" xml:"CommunityMatchMode,omitempty"`
	CommunityOperateMode               *string   `json:"CommunityOperateMode,omitempty" xml:"CommunityOperateMode,omitempty"`
	Preference                         *int32    `json:"Preference,omitempty" xml:"Preference,omitempty"`
	SourceInstanceIdsReverseMatch      *bool     `json:"SourceInstanceIdsReverseMatch,omitempty" xml:"SourceInstanceIdsReverseMatch,omitempty"`
	DestinationInstanceIdsReverseMatch *bool     `json:"DestinationInstanceIdsReverseMatch,omitempty" xml:"DestinationInstanceIdsReverseMatch,omitempty"`
	GatewayZoneId                      *string   `json:"GatewayZoneId,omitempty" xml:"GatewayZoneId,omitempty"`
	SystemPolicy                       *bool     `json:"SystemPolicy,omitempty" xml:"SystemPolicy,omitempty"`
	MatchAddressType                   *string   `json:"MatchAddressType,omitempty" xml:"MatchAddressType,omitempty"`
	SourceInstanceIds                  []*string `json:"SourceInstanceIds,omitempty" xml:"SourceInstanceIds,omitempty" type:"Repeated"`
	DestinationInstanceIds             []*string `json:"DestinationInstanceIds,omitempty" xml:"DestinationInstanceIds,omitempty" type:"Repeated"`
	SourceRouteTableIds                []*string `json:"SourceRouteTableIds,omitempty" xml:"SourceRouteTableIds,omitempty" type:"Repeated"`
	DestinationRouteTableIds           []*string `json:"DestinationRouteTableIds,omitempty" xml:"DestinationRouteTableIds,omitempty" type:"Repeated"`
	SourceRegionIds                    []*string `json:"SourceRegionIds,omitempty" xml:"SourceRegionIds,omitempty" type:"Repeated"`
	SourceChildInstanceTypes           []*string `json:"SourceChildInstanceTypes,omitempty" xml:"SourceChildInstanceTypes,omitempty" type:"Repeated"`
	DestinationChildInstanceTypes      []*string `json:"DestinationChildInstanceTypes,omitempty" xml:"DestinationChildInstanceTypes,omitempty" type:"Repeated"`
	DestinationCidrBlocks              []*string `json:"DestinationCidrBlocks,omitempty" xml:"DestinationCidrBlocks,omitempty" type:"Repeated"`
	RouteTypes                         []*string `json:"RouteTypes,omitempty" xml:"RouteTypes,omitempty" type:"Repeated"`
	MatchAsns                          []*int    `json:"MatchAsns,omitempty" xml:"MatchAsns,omitempty" type:"Repeated"`
	MatchCommunitySet                  []*string `json:"MatchCommunitySet,omitempty" xml:"MatchCommunitySet,omitempty" type:"Repeated"`
	OperateCommunitySet                []*string `json:"OperateCommunitySet,omitempty" xml:"OperateCommunitySet,omitempty" type:"Repeated"`
	PrependAsPath                      []*int    `json:"PrependAsPath,omitempty" xml:"PrependAsPath,omitempty" type:"Repeated"`
	DestinationRegionIds               []*string `json:"DestinationRegionIds,omitempty" xml:"DestinationRegionIds,omitempty" type:"Repeated"`
	SourceZoneIds                      []*string `json:"SourceZoneIds,omitempty" xml:"SourceZoneIds,omitempty" type:"Repeated"`
	OriginalRouteTableIds              []*string `json:"OriginalRouteTableIds,omitempty" xml:"OriginalRouteTableIds,omitempty" type:"Repeated"`
}

func (s CreateCenRouteMapRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCenRouteMapRequest) GoString() string {
	return s.String()
}

func (s *CreateCenRouteMapRequest) SetOwnerAccount(v string) *CreateCenRouteMapRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetOwnerId(v int64) *CreateCenRouteMapRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetResourceOwnerAccount(v string) *CreateCenRouteMapRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetResourceOwnerId(v int64) *CreateCenRouteMapRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetCenId(v string) *CreateCenRouteMapRequest {
	s.CenId = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetCenRegionId(v string) *CreateCenRouteMapRequest {
	s.CenRegionId = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetTransmitDirection(v string) *CreateCenRouteMapRequest {
	s.TransmitDirection = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetDescription(v string) *CreateCenRouteMapRequest {
	s.Description = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetPriority(v int32) *CreateCenRouteMapRequest {
	s.Priority = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetMapResult(v string) *CreateCenRouteMapRequest {
	s.MapResult = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetNextPriority(v int32) *CreateCenRouteMapRequest {
	s.NextPriority = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetCidrMatchMode(v string) *CreateCenRouteMapRequest {
	s.CidrMatchMode = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetAsPathMatchMode(v string) *CreateCenRouteMapRequest {
	s.AsPathMatchMode = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetCommunityMatchMode(v string) *CreateCenRouteMapRequest {
	s.CommunityMatchMode = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetCommunityOperateMode(v string) *CreateCenRouteMapRequest {
	s.CommunityOperateMode = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetPreference(v int32) *CreateCenRouteMapRequest {
	s.Preference = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetSourceInstanceIdsReverseMatch(v bool) *CreateCenRouteMapRequest {
	s.SourceInstanceIdsReverseMatch = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetDestinationInstanceIdsReverseMatch(v bool) *CreateCenRouteMapRequest {
	s.DestinationInstanceIdsReverseMatch = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetGatewayZoneId(v string) *CreateCenRouteMapRequest {
	s.GatewayZoneId = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetSystemPolicy(v bool) *CreateCenRouteMapRequest {
	s.SystemPolicy = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetMatchAddressType(v string) *CreateCenRouteMapRequest {
	s.MatchAddressType = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetSourceInstanceIds(v []*string) *CreateCenRouteMapRequest {
	s.SourceInstanceIds = v
	return s
}

func (s *CreateCenRouteMapRequest) SetDestinationInstanceIds(v []*string) *CreateCenRouteMapRequest {
	s.DestinationInstanceIds = v
	return s
}

func (s *CreateCenRouteMapRequest) SetSourceRouteTableIds(v []*string) *CreateCenRouteMapRequest {
	s.SourceRouteTableIds = v
	return s
}

func (s *CreateCenRouteMapRequest) SetDestinationRouteTableIds(v []*string) *CreateCenRouteMapRequest {
	s.DestinationRouteTableIds = v
	return s
}

func (s *CreateCenRouteMapRequest) SetSourceRegionIds(v []*string) *CreateCenRouteMapRequest {
	s.SourceRegionIds = v
	return s
}

func (s *CreateCenRouteMapRequest) SetSourceChildInstanceTypes(v []*string) *CreateCenRouteMapRequest {
	s.SourceChildInstanceTypes = v
	return s
}

func (s *CreateCenRouteMapRequest) SetDestinationChildInstanceTypes(v []*string) *CreateCenRouteMapRequest {
	s.DestinationChildInstanceTypes = v
	return s
}

func (s *CreateCenRouteMapRequest) SetDestinationCidrBlocks(v []*string) *CreateCenRouteMapRequest {
	s.DestinationCidrBlocks = v
	return s
}

func (s *CreateCenRouteMapRequest) SetRouteTypes(v []*string) *CreateCenRouteMapRequest {
	s.RouteTypes = v
	return s
}

func (s *CreateCenRouteMapRequest) SetMatchAsns(v []*int) *CreateCenRouteMapRequest {
	s.MatchAsns = v
	return s
}

func (s *CreateCenRouteMapRequest) SetMatchCommunitySet(v []*string) *CreateCenRouteMapRequest {
	s.MatchCommunitySet = v
	return s
}

func (s *CreateCenRouteMapRequest) SetOperateCommunitySet(v []*string) *CreateCenRouteMapRequest {
	s.OperateCommunitySet = v
	return s
}

func (s *CreateCenRouteMapRequest) SetPrependAsPath(v []*int) *CreateCenRouteMapRequest {
	s.PrependAsPath = v
	return s
}

func (s *CreateCenRouteMapRequest) SetDestinationRegionIds(v []*string) *CreateCenRouteMapRequest {
	s.DestinationRegionIds = v
	return s
}

func (s *CreateCenRouteMapRequest) SetSourceZoneIds(v []*string) *CreateCenRouteMapRequest {
	s.SourceZoneIds = v
	return s
}

func (s *CreateCenRouteMapRequest) SetOriginalRouteTableIds(v []*string) *CreateCenRouteMapRequest {
	s.OriginalRouteTableIds = v
	return s
}

type CreateCenRouteMapResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RouteMapId *string `json:"RouteMapId,omitempty" xml:"RouteMapId,omitempty"`
}

func (s CreateCenRouteMapResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCenRouteMapResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCenRouteMapResponseBody) SetRequestId(v string) *CreateCenRouteMapResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCenRouteMapResponseBody) SetRouteMapId(v string) *CreateCenRouteMapResponseBody {
	s.RouteMapId = &v
	return s
}

type CreateCenRouteMapResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCenRouteMapResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCenRouteMapResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCenRouteMapResponse) GoString() string {
	return s.String()
}

func (s *CreateCenRouteMapResponse) SetHeaders(v map[string]*string) *CreateCenRouteMapResponse {
	s.Headers = v
	return s
}

func (s *CreateCenRouteMapResponse) SetBody(v *CreateCenRouteMapResponseBody) *CreateCenRouteMapResponse {
	s.Body = v
	return s
}

type CreateFlowlogRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	FlowLogName          *string `json:"FlowLogName,omitempty" xml:"FlowLogName,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ProjectName          *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	LogStoreName         *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
}

func (s CreateFlowlogRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowlogRequest) GoString() string {
	return s.String()
}

func (s *CreateFlowlogRequest) SetOwnerAccount(v string) *CreateFlowlogRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateFlowlogRequest) SetOwnerId(v int64) *CreateFlowlogRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateFlowlogRequest) SetResourceOwnerAccount(v string) *CreateFlowlogRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateFlowlogRequest) SetResourceOwnerId(v int64) *CreateFlowlogRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateFlowlogRequest) SetClientToken(v string) *CreateFlowlogRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateFlowlogRequest) SetRegionId(v string) *CreateFlowlogRequest {
	s.RegionId = &v
	return s
}

func (s *CreateFlowlogRequest) SetFlowLogName(v string) *CreateFlowlogRequest {
	s.FlowLogName = &v
	return s
}

func (s *CreateFlowlogRequest) SetDescription(v string) *CreateFlowlogRequest {
	s.Description = &v
	return s
}

func (s *CreateFlowlogRequest) SetCenId(v string) *CreateFlowlogRequest {
	s.CenId = &v
	return s
}

func (s *CreateFlowlogRequest) SetProjectName(v string) *CreateFlowlogRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateFlowlogRequest) SetLogStoreName(v string) *CreateFlowlogRequest {
	s.LogStoreName = &v
	return s
}

type CreateFlowlogResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FlowLogId *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateFlowlogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowlogResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFlowlogResponseBody) SetRequestId(v string) *CreateFlowlogResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFlowlogResponseBody) SetFlowLogId(v string) *CreateFlowlogResponseBody {
	s.FlowLogId = &v
	return s
}

func (s *CreateFlowlogResponseBody) SetSuccess(v string) *CreateFlowlogResponseBody {
	s.Success = &v
	return s
}

type CreateFlowlogResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFlowlogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFlowlogResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowlogResponse) GoString() string {
	return s.String()
}

func (s *CreateFlowlogResponse) SetHeaders(v map[string]*string) *CreateFlowlogResponse {
	s.Headers = v
	return s
}

func (s *CreateFlowlogResponse) SetBody(v *CreateFlowlogResponseBody) *CreateFlowlogResponse {
	s.Body = v
	return s
}

type DeactiveFlowLogRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	FlowLogId            *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
}

func (s DeactiveFlowLogRequest) String() string {
	return tea.Prettify(s)
}

func (s DeactiveFlowLogRequest) GoString() string {
	return s.String()
}

func (s *DeactiveFlowLogRequest) SetOwnerAccount(v string) *DeactiveFlowLogRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeactiveFlowLogRequest) SetOwnerId(v int64) *DeactiveFlowLogRequest {
	s.OwnerId = &v
	return s
}

func (s *DeactiveFlowLogRequest) SetResourceOwnerAccount(v string) *DeactiveFlowLogRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeactiveFlowLogRequest) SetResourceOwnerId(v int64) *DeactiveFlowLogRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeactiveFlowLogRequest) SetClientToken(v string) *DeactiveFlowLogRequest {
	s.ClientToken = &v
	return s
}

func (s *DeactiveFlowLogRequest) SetRegionId(v string) *DeactiveFlowLogRequest {
	s.RegionId = &v
	return s
}

func (s *DeactiveFlowLogRequest) SetFlowLogId(v string) *DeactiveFlowLogRequest {
	s.FlowLogId = &v
	return s
}

func (s *DeactiveFlowLogRequest) SetCenId(v string) *DeactiveFlowLogRequest {
	s.CenId = &v
	return s
}

type DeactiveFlowLogResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeactiveFlowLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeactiveFlowLogResponseBody) GoString() string {
	return s.String()
}

func (s *DeactiveFlowLogResponseBody) SetRequestId(v string) *DeactiveFlowLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeactiveFlowLogResponseBody) SetSuccess(v string) *DeactiveFlowLogResponseBody {
	s.Success = &v
	return s
}

type DeactiveFlowLogResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeactiveFlowLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeactiveFlowLogResponse) String() string {
	return tea.Prettify(s)
}

func (s DeactiveFlowLogResponse) GoString() string {
	return s.String()
}

func (s *DeactiveFlowLogResponse) SetHeaders(v map[string]*string) *DeactiveFlowLogResponse {
	s.Headers = v
	return s
}

func (s *DeactiveFlowLogResponse) SetBody(v *DeactiveFlowLogResponseBody) *DeactiveFlowLogResponse {
	s.Body = v
	return s
}

type DeleteCenRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
}

func (s DeleteCenRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCenRequest) GoString() string {
	return s.String()
}

func (s *DeleteCenRequest) SetOwnerAccount(v string) *DeleteCenRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteCenRequest) SetOwnerId(v int64) *DeleteCenRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteCenRequest) SetResourceOwnerAccount(v string) *DeleteCenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteCenRequest) SetResourceOwnerId(v int64) *DeleteCenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteCenRequest) SetCenId(v string) *DeleteCenRequest {
	s.CenId = &v
	return s
}

type DeleteCenResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCenResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCenResponseBody) SetRequestId(v string) *DeleteCenResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCenResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteCenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCenResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCenResponse) GoString() string {
	return s.String()
}

func (s *DeleteCenResponse) SetHeaders(v map[string]*string) *DeleteCenResponse {
	s.Headers = v
	return s
}

func (s *DeleteCenResponse) SetBody(v *DeleteCenResponseBody) *DeleteCenResponse {
	s.Body = v
	return s
}

type DeleteCenBandwidthPackageRequest struct {
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	CenBandwidthPackageId *string `json:"CenBandwidthPackageId,omitempty" xml:"CenBandwidthPackageId,omitempty"`
}

func (s DeleteCenBandwidthPackageRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCenBandwidthPackageRequest) GoString() string {
	return s.String()
}

func (s *DeleteCenBandwidthPackageRequest) SetOwnerAccount(v string) *DeleteCenBandwidthPackageRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteCenBandwidthPackageRequest) SetOwnerId(v int64) *DeleteCenBandwidthPackageRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteCenBandwidthPackageRequest) SetResourceOwnerAccount(v string) *DeleteCenBandwidthPackageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteCenBandwidthPackageRequest) SetResourceOwnerId(v int64) *DeleteCenBandwidthPackageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteCenBandwidthPackageRequest) SetCenBandwidthPackageId(v string) *DeleteCenBandwidthPackageRequest {
	s.CenBandwidthPackageId = &v
	return s
}

type DeleteCenBandwidthPackageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCenBandwidthPackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCenBandwidthPackageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCenBandwidthPackageResponseBody) SetRequestId(v string) *DeleteCenBandwidthPackageResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCenBandwidthPackageResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteCenBandwidthPackageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCenBandwidthPackageResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCenBandwidthPackageResponse) GoString() string {
	return s.String()
}

func (s *DeleteCenBandwidthPackageResponse) SetHeaders(v map[string]*string) *DeleteCenBandwidthPackageResponse {
	s.Headers = v
	return s
}

func (s *DeleteCenBandwidthPackageResponse) SetBody(v *DeleteCenBandwidthPackageResponseBody) *DeleteCenBandwidthPackageResponse {
	s.Body = v
	return s
}

type DeleteCenChildInstanceRouteEntryToCenRequest struct {
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	CenId                 *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ChildInstanceId       *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	ChildInstanceType     *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	ChildInstanceRegionId *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	ChildInstanceAliUid   *int64  `json:"ChildInstanceAliUid,omitempty" xml:"ChildInstanceAliUid,omitempty"`
	RouteTableId          *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
	DestinationCidrBlock  *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
}

func (s DeleteCenChildInstanceRouteEntryToCenRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCenChildInstanceRouteEntryToCenRequest) GoString() string {
	return s.String()
}

func (s *DeleteCenChildInstanceRouteEntryToCenRequest) SetOwnerAccount(v string) *DeleteCenChildInstanceRouteEntryToCenRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToCenRequest) SetOwnerId(v int64) *DeleteCenChildInstanceRouteEntryToCenRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToCenRequest) SetResourceOwnerAccount(v string) *DeleteCenChildInstanceRouteEntryToCenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToCenRequest) SetResourceOwnerId(v int64) *DeleteCenChildInstanceRouteEntryToCenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToCenRequest) SetCenId(v string) *DeleteCenChildInstanceRouteEntryToCenRequest {
	s.CenId = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToCenRequest) SetChildInstanceId(v string) *DeleteCenChildInstanceRouteEntryToCenRequest {
	s.ChildInstanceId = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToCenRequest) SetChildInstanceType(v string) *DeleteCenChildInstanceRouteEntryToCenRequest {
	s.ChildInstanceType = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToCenRequest) SetChildInstanceRegionId(v string) *DeleteCenChildInstanceRouteEntryToCenRequest {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToCenRequest) SetChildInstanceAliUid(v int64) *DeleteCenChildInstanceRouteEntryToCenRequest {
	s.ChildInstanceAliUid = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToCenRequest) SetRouteTableId(v string) *DeleteCenChildInstanceRouteEntryToCenRequest {
	s.RouteTableId = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToCenRequest) SetDestinationCidrBlock(v string) *DeleteCenChildInstanceRouteEntryToCenRequest {
	s.DestinationCidrBlock = &v
	return s
}

type DeleteCenChildInstanceRouteEntryToCenResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCenChildInstanceRouteEntryToCenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCenChildInstanceRouteEntryToCenResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCenChildInstanceRouteEntryToCenResponseBody) SetRequestId(v string) *DeleteCenChildInstanceRouteEntryToCenResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCenChildInstanceRouteEntryToCenResponse struct {
	Headers map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteCenChildInstanceRouteEntryToCenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCenChildInstanceRouteEntryToCenResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCenChildInstanceRouteEntryToCenResponse) GoString() string {
	return s.String()
}

func (s *DeleteCenChildInstanceRouteEntryToCenResponse) SetHeaders(v map[string]*string) *DeleteCenChildInstanceRouteEntryToCenResponse {
	s.Headers = v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToCenResponse) SetBody(v *DeleteCenChildInstanceRouteEntryToCenResponseBody) *DeleteCenChildInstanceRouteEntryToCenResponse {
	s.Body = v
	return s
}

type DeleteCenRouteMapRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	CenRegionId          *string `json:"CenRegionId,omitempty" xml:"CenRegionId,omitempty"`
	RouteMapId           *string `json:"RouteMapId,omitempty" xml:"RouteMapId,omitempty"`
}

func (s DeleteCenRouteMapRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCenRouteMapRequest) GoString() string {
	return s.String()
}

func (s *DeleteCenRouteMapRequest) SetOwnerAccount(v string) *DeleteCenRouteMapRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteCenRouteMapRequest) SetOwnerId(v int64) *DeleteCenRouteMapRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteCenRouteMapRequest) SetResourceOwnerAccount(v string) *DeleteCenRouteMapRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteCenRouteMapRequest) SetResourceOwnerId(v int64) *DeleteCenRouteMapRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteCenRouteMapRequest) SetCenId(v string) *DeleteCenRouteMapRequest {
	s.CenId = &v
	return s
}

func (s *DeleteCenRouteMapRequest) SetCenRegionId(v string) *DeleteCenRouteMapRequest {
	s.CenRegionId = &v
	return s
}

func (s *DeleteCenRouteMapRequest) SetRouteMapId(v string) *DeleteCenRouteMapRequest {
	s.RouteMapId = &v
	return s
}

type DeleteCenRouteMapResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCenRouteMapResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCenRouteMapResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCenRouteMapResponseBody) SetRequestId(v string) *DeleteCenRouteMapResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCenRouteMapResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteCenRouteMapResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCenRouteMapResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCenRouteMapResponse) GoString() string {
	return s.String()
}

func (s *DeleteCenRouteMapResponse) SetHeaders(v map[string]*string) *DeleteCenRouteMapResponse {
	s.Headers = v
	return s
}

func (s *DeleteCenRouteMapResponse) SetBody(v *DeleteCenRouteMapResponseBody) *DeleteCenRouteMapResponse {
	s.Body = v
	return s
}

type DeleteFlowlogRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	FlowLogId            *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
}

func (s DeleteFlowlogRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowlogRequest) GoString() string {
	return s.String()
}

func (s *DeleteFlowlogRequest) SetOwnerAccount(v string) *DeleteFlowlogRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteFlowlogRequest) SetOwnerId(v int64) *DeleteFlowlogRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteFlowlogRequest) SetResourceOwnerAccount(v string) *DeleteFlowlogRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteFlowlogRequest) SetResourceOwnerId(v int64) *DeleteFlowlogRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteFlowlogRequest) SetClientToken(v string) *DeleteFlowlogRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteFlowlogRequest) SetRegionId(v string) *DeleteFlowlogRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteFlowlogRequest) SetFlowLogId(v string) *DeleteFlowlogRequest {
	s.FlowLogId = &v
	return s
}

func (s *DeleteFlowlogRequest) SetCenId(v string) *DeleteFlowlogRequest {
	s.CenId = &v
	return s
}

type DeleteFlowlogResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteFlowlogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowlogResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFlowlogResponseBody) SetRequestId(v string) *DeleteFlowlogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFlowlogResponseBody) SetSuccess(v string) *DeleteFlowlogResponseBody {
	s.Success = &v
	return s
}

type DeleteFlowlogResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFlowlogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFlowlogResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowlogResponse) GoString() string {
	return s.String()
}

func (s *DeleteFlowlogResponse) SetHeaders(v map[string]*string) *DeleteFlowlogResponse {
	s.Headers = v
	return s
}

func (s *DeleteFlowlogResponse) SetBody(v *DeleteFlowlogResponseBody) *DeleteFlowlogResponse {
	s.Body = v
	return s
}

type DeleteRouteServiceInCenRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	Host                 *string `json:"Host,omitempty" xml:"Host,omitempty"`
	HostRegionId         *string `json:"HostRegionId,omitempty" xml:"HostRegionId,omitempty"`
	AccessRegionId       *string `json:"AccessRegionId,omitempty" xml:"AccessRegionId,omitempty"`
	HostVpcId            *string `json:"HostVpcId,omitempty" xml:"HostVpcId,omitempty"`
}

func (s DeleteRouteServiceInCenRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRouteServiceInCenRequest) GoString() string {
	return s.String()
}

func (s *DeleteRouteServiceInCenRequest) SetOwnerAccount(v string) *DeleteRouteServiceInCenRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteRouteServiceInCenRequest) SetOwnerId(v int64) *DeleteRouteServiceInCenRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteRouteServiceInCenRequest) SetResourceOwnerAccount(v string) *DeleteRouteServiceInCenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteRouteServiceInCenRequest) SetResourceOwnerId(v int64) *DeleteRouteServiceInCenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteRouteServiceInCenRequest) SetCenId(v string) *DeleteRouteServiceInCenRequest {
	s.CenId = &v
	return s
}

func (s *DeleteRouteServiceInCenRequest) SetHost(v string) *DeleteRouteServiceInCenRequest {
	s.Host = &v
	return s
}

func (s *DeleteRouteServiceInCenRequest) SetHostRegionId(v string) *DeleteRouteServiceInCenRequest {
	s.HostRegionId = &v
	return s
}

func (s *DeleteRouteServiceInCenRequest) SetAccessRegionId(v string) *DeleteRouteServiceInCenRequest {
	s.AccessRegionId = &v
	return s
}

func (s *DeleteRouteServiceInCenRequest) SetHostVpcId(v string) *DeleteRouteServiceInCenRequest {
	s.HostVpcId = &v
	return s
}

type DeleteRouteServiceInCenResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRouteServiceInCenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRouteServiceInCenResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRouteServiceInCenResponseBody) SetRequestId(v string) *DeleteRouteServiceInCenResponseBody {
	s.RequestId = &v
	return s
}

type DeleteRouteServiceInCenResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteRouteServiceInCenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRouteServiceInCenResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRouteServiceInCenResponse) GoString() string {
	return s.String()
}

func (s *DeleteRouteServiceInCenResponse) SetHeaders(v map[string]*string) *DeleteRouteServiceInCenResponse {
	s.Headers = v
	return s
}

func (s *DeleteRouteServiceInCenResponse) SetBody(v *DeleteRouteServiceInCenResponseBody) *DeleteRouteServiceInCenResponse {
	s.Body = v
	return s
}

type DescribeCenAttachedChildInstanceAttributeRequest struct {
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	CenId                 *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ChildInstanceId       *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	ChildInstanceType     *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	ChildInstanceRegionId *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	IncludeRouteTable     *bool   `json:"IncludeRouteTable,omitempty" xml:"IncludeRouteTable,omitempty"`
}

func (s DescribeCenAttachedChildInstanceAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenAttachedChildInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeCenAttachedChildInstanceAttributeRequest) SetOwnerAccount(v string) *DescribeCenAttachedChildInstanceAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeRequest) SetOwnerId(v int64) *DescribeCenAttachedChildInstanceAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeRequest) SetResourceOwnerAccount(v string) *DescribeCenAttachedChildInstanceAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeRequest) SetResourceOwnerId(v int64) *DescribeCenAttachedChildInstanceAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeRequest) SetCenId(v string) *DescribeCenAttachedChildInstanceAttributeRequest {
	s.CenId = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeRequest) SetChildInstanceId(v string) *DescribeCenAttachedChildInstanceAttributeRequest {
	s.ChildInstanceId = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeRequest) SetChildInstanceType(v string) *DescribeCenAttachedChildInstanceAttributeRequest {
	s.ChildInstanceType = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeRequest) SetChildInstanceRegionId(v string) *DescribeCenAttachedChildInstanceAttributeRequest {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeRequest) SetIncludeRouteTable(v bool) *DescribeCenAttachedChildInstanceAttributeRequest {
	s.IncludeRouteTable = &v
	return s
}

type DescribeCenAttachedChildInstanceAttributeResponseBody struct {
	Status                  *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ChildInstanceType       *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	RequestId               *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CenId                   *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ChildInstanceAttachTime *string `json:"ChildInstanceAttachTime,omitempty" xml:"ChildInstanceAttachTime,omitempty"`
	ChildInstanceOwnerId    *int64  `json:"ChildInstanceOwnerId,omitempty" xml:"ChildInstanceOwnerId,omitempty"`
	ChildInstanceName       *string `json:"ChildInstanceName,omitempty" xml:"ChildInstanceName,omitempty"`
	ChildInstanceId         *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	ChildInstanceRegionId   *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
}

func (s DescribeCenAttachedChildInstanceAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenAttachedChildInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) SetStatus(v string) *DescribeCenAttachedChildInstanceAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) SetChildInstanceType(v string) *DescribeCenAttachedChildInstanceAttributeResponseBody {
	s.ChildInstanceType = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) SetRequestId(v string) *DescribeCenAttachedChildInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) SetCenId(v string) *DescribeCenAttachedChildInstanceAttributeResponseBody {
	s.CenId = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) SetChildInstanceAttachTime(v string) *DescribeCenAttachedChildInstanceAttributeResponseBody {
	s.ChildInstanceAttachTime = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) SetChildInstanceOwnerId(v int64) *DescribeCenAttachedChildInstanceAttributeResponseBody {
	s.ChildInstanceOwnerId = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) SetChildInstanceName(v string) *DescribeCenAttachedChildInstanceAttributeResponseBody {
	s.ChildInstanceName = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) SetChildInstanceId(v string) *DescribeCenAttachedChildInstanceAttributeResponseBody {
	s.ChildInstanceId = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) SetChildInstanceRegionId(v string) *DescribeCenAttachedChildInstanceAttributeResponseBody {
	s.ChildInstanceRegionId = &v
	return s
}

type DescribeCenAttachedChildInstanceAttributeResponse struct {
	Headers map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCenAttachedChildInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCenAttachedChildInstanceAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenAttachedChildInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeCenAttachedChildInstanceAttributeResponse) SetHeaders(v map[string]*string) *DescribeCenAttachedChildInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeResponse) SetBody(v *DescribeCenAttachedChildInstanceAttributeResponseBody) *DescribeCenAttachedChildInstanceAttributeResponse {
	s.Body = v
	return s
}

type DescribeCenAttachedChildInstancesRequest struct {
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PageNumber            *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize              *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CenId                 *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ChildInstanceType     *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	ChildInstanceRegionId *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
}

func (s DescribeCenAttachedChildInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenAttachedChildInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCenAttachedChildInstancesRequest) SetOwnerAccount(v string) *DescribeCenAttachedChildInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesRequest) SetOwnerId(v int64) *DescribeCenAttachedChildInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesRequest) SetResourceOwnerAccount(v string) *DescribeCenAttachedChildInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesRequest) SetResourceOwnerId(v int64) *DescribeCenAttachedChildInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesRequest) SetPageNumber(v int32) *DescribeCenAttachedChildInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesRequest) SetPageSize(v int32) *DescribeCenAttachedChildInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesRequest) SetCenId(v string) *DescribeCenAttachedChildInstancesRequest {
	s.CenId = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesRequest) SetChildInstanceType(v string) *DescribeCenAttachedChildInstancesRequest {
	s.ChildInstanceType = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesRequest) SetChildInstanceRegionId(v string) *DescribeCenAttachedChildInstancesRequest {
	s.ChildInstanceRegionId = &v
	return s
}

type DescribeCenAttachedChildInstancesResponseBody struct {
	TotalCount     *int32                                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize       *int32                                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber     *int32                                                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	ChildInstances []*DescribeCenAttachedChildInstancesResponseBodyChildInstances `json:"ChildInstances,omitempty" xml:"ChildInstances,omitempty" type:"Repeated"`
}

func (s DescribeCenAttachedChildInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenAttachedChildInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCenAttachedChildInstancesResponseBody) SetTotalCount(v int32) *DescribeCenAttachedChildInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesResponseBody) SetPageSize(v int32) *DescribeCenAttachedChildInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesResponseBody) SetRequestId(v string) *DescribeCenAttachedChildInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesResponseBody) SetPageNumber(v int32) *DescribeCenAttachedChildInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesResponseBody) SetChildInstances(v []*DescribeCenAttachedChildInstancesResponseBodyChildInstances) *DescribeCenAttachedChildInstancesResponseBody {
	s.ChildInstances = v
	return s
}

type DescribeCenAttachedChildInstancesResponseBodyChildInstances struct {
	ChildInstanceType       *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	Status                  *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ChildInstanceRegionId   *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	ChildInstanceOwnerId    *int64  `json:"ChildInstanceOwnerId,omitempty" xml:"ChildInstanceOwnerId,omitempty"`
	ChildInstanceId         *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	CenId                   *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ChildInstanceAttachTime *string `json:"ChildInstanceAttachTime,omitempty" xml:"ChildInstanceAttachTime,omitempty"`
}

func (s DescribeCenAttachedChildInstancesResponseBodyChildInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenAttachedChildInstancesResponseBodyChildInstances) GoString() string {
	return s.String()
}

func (s *DescribeCenAttachedChildInstancesResponseBodyChildInstances) SetChildInstanceType(v string) *DescribeCenAttachedChildInstancesResponseBodyChildInstances {
	s.ChildInstanceType = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesResponseBodyChildInstances) SetStatus(v string) *DescribeCenAttachedChildInstancesResponseBodyChildInstances {
	s.Status = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesResponseBodyChildInstances) SetChildInstanceRegionId(v string) *DescribeCenAttachedChildInstancesResponseBodyChildInstances {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesResponseBodyChildInstances) SetChildInstanceOwnerId(v int64) *DescribeCenAttachedChildInstancesResponseBodyChildInstances {
	s.ChildInstanceOwnerId = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesResponseBodyChildInstances) SetChildInstanceId(v string) *DescribeCenAttachedChildInstancesResponseBodyChildInstances {
	s.ChildInstanceId = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesResponseBodyChildInstances) SetCenId(v string) *DescribeCenAttachedChildInstancesResponseBodyChildInstances {
	s.CenId = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesResponseBodyChildInstances) SetChildInstanceAttachTime(v string) *DescribeCenAttachedChildInstancesResponseBodyChildInstances {
	s.ChildInstanceAttachTime = &v
	return s
}

type DescribeCenAttachedChildInstancesResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCenAttachedChildInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCenAttachedChildInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenAttachedChildInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCenAttachedChildInstancesResponse) SetHeaders(v map[string]*string) *DescribeCenAttachedChildInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCenAttachedChildInstancesResponse) SetBody(v *DescribeCenAttachedChildInstancesResponseBody) *DescribeCenAttachedChildInstancesResponse {
	s.Body = v
	return s
}

type DescribeCenBandwidthPackagesRequest struct {
	IncludeReservationData *bool                                        `json:"IncludeReservationData,omitempty" xml:"IncludeReservationData,omitempty"`
	OwnerAccount           *string                                      `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                *int64                                       `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount   *string                                      `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId        *int64                                       `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PageNumber             *int32                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize               *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	IsOrKey                *bool                                        `json:"IsOrKey,omitempty" xml:"IsOrKey,omitempty"`
	Filter                 []*DescribeCenBandwidthPackagesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
}

func (s DescribeCenBandwidthPackagesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenBandwidthPackagesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCenBandwidthPackagesRequest) SetIncludeReservationData(v bool) *DescribeCenBandwidthPackagesRequest {
	s.IncludeReservationData = &v
	return s
}

func (s *DescribeCenBandwidthPackagesRequest) SetOwnerAccount(v string) *DescribeCenBandwidthPackagesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCenBandwidthPackagesRequest) SetOwnerId(v int64) *DescribeCenBandwidthPackagesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCenBandwidthPackagesRequest) SetResourceOwnerAccount(v string) *DescribeCenBandwidthPackagesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCenBandwidthPackagesRequest) SetResourceOwnerId(v int64) *DescribeCenBandwidthPackagesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCenBandwidthPackagesRequest) SetPageNumber(v int32) *DescribeCenBandwidthPackagesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCenBandwidthPackagesRequest) SetPageSize(v int32) *DescribeCenBandwidthPackagesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCenBandwidthPackagesRequest) SetIsOrKey(v bool) *DescribeCenBandwidthPackagesRequest {
	s.IsOrKey = &v
	return s
}

func (s *DescribeCenBandwidthPackagesRequest) SetFilter(v []*DescribeCenBandwidthPackagesRequestFilter) *DescribeCenBandwidthPackagesRequest {
	s.Filter = v
	return s
}

type DescribeCenBandwidthPackagesRequestFilter struct {
	Key   *string   `json:"Key,omitempty" xml:"Key,omitempty"`
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s DescribeCenBandwidthPackagesRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenBandwidthPackagesRequestFilter) GoString() string {
	return s.String()
}

func (s *DescribeCenBandwidthPackagesRequestFilter) SetKey(v string) *DescribeCenBandwidthPackagesRequestFilter {
	s.Key = &v
	return s
}

func (s *DescribeCenBandwidthPackagesRequestFilter) SetValue(v []*string) *DescribeCenBandwidthPackagesRequestFilter {
	s.Value = v
	return s
}

type DescribeCenBandwidthPackagesResponseBody struct {
	TotalCount           *int32                                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize             *int32                                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId            *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber           *int32                                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	CenBandwidthPackages []*DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages `json:"CenBandwidthPackages,omitempty" xml:"CenBandwidthPackages,omitempty" type:"Repeated"`
}

func (s DescribeCenBandwidthPackagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenBandwidthPackagesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCenBandwidthPackagesResponseBody) SetTotalCount(v int32) *DescribeCenBandwidthPackagesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBody) SetPageSize(v int32) *DescribeCenBandwidthPackagesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBody) SetRequestId(v string) *DescribeCenBandwidthPackagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBody) SetPageNumber(v int32) *DescribeCenBandwidthPackagesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBody) SetCenBandwidthPackages(v []*DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages) *DescribeCenBandwidthPackagesResponseBody {
	s.CenBandwidthPackages = v
	return s
}

type DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages struct {
	CreationTime                    *string                                                                                        `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Status                          *string                                                                                        `json:"Status,omitempty" xml:"Status,omitempty"`
	ReservationActiveTime           *string                                                                                        `json:"ReservationActiveTime,omitempty" xml:"ReservationActiveTime,omitempty"`
	ReservationOrderType            *string                                                                                        `json:"ReservationOrderType,omitempty" xml:"ReservationOrderType,omitempty"`
	BandwidthPackageChargeType      *string                                                                                        `json:"BandwidthPackageChargeType,omitempty" xml:"BandwidthPackageChargeType,omitempty"`
	CenBandwidthPackageId           *string                                                                                        `json:"CenBandwidthPackageId,omitempty" xml:"CenBandwidthPackageId,omitempty"`
	OrginInterRegionBandwidthLimits []*DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesOrginInterRegionBandwidthLimits `json:"OrginInterRegionBandwidthLimits,omitempty" xml:"OrginInterRegionBandwidthLimits,omitempty" type:"Repeated"`
	CenIds                          []*string                                                                                      `json:"CenIds,omitempty" xml:"CenIds,omitempty" type:"Repeated"`
	GeographicRegionAId             *string                                                                                        `json:"GeographicRegionAId,omitempty" xml:"GeographicRegionAId,omitempty"`
	Ratio                           *string                                                                                        `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	ReservationInternetChargeType   *string                                                                                        `json:"ReservationInternetChargeType,omitempty" xml:"ReservationInternetChargeType,omitempty"`
	TypeFor95                       *string                                                                                        `json:"TypeFor95,omitempty" xml:"TypeFor95,omitempty"`
	Bandwidth                       *int64                                                                                         `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	Description                     *string                                                                                        `json:"Description,omitempty" xml:"Description,omitempty"`
	ExpiredTime                     *string                                                                                        `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	ReservationBandwidth            *string                                                                                        `json:"ReservationBandwidth,omitempty" xml:"ReservationBandwidth,omitempty"`
	GeographicSpanId                *string                                                                                        `json:"GeographicSpanId,omitempty" xml:"GeographicSpanId,omitempty"`
	GeographicRegionBId             *string                                                                                        `json:"GeographicRegionBId,omitempty" xml:"GeographicRegionBId,omitempty"`
	IsCrossBorder                   *bool                                                                                          `json:"IsCrossBorder,omitempty" xml:"IsCrossBorder,omitempty"`
	BusinessStatus                  *string                                                                                        `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	Name                            *string                                                                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	HasReservationData              *string                                                                                        `json:"HasReservationData,omitempty" xml:"HasReservationData,omitempty"`
}

func (s DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages) GoString() string {
	return s.String()
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages) SetCreationTime(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages {
	s.CreationTime = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages) SetStatus(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages {
	s.Status = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages) SetReservationActiveTime(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages {
	s.ReservationActiveTime = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages) SetReservationOrderType(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages {
	s.ReservationOrderType = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages) SetBandwidthPackageChargeType(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages {
	s.BandwidthPackageChargeType = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages) SetCenBandwidthPackageId(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages {
	s.CenBandwidthPackageId = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages) SetOrginInterRegionBandwidthLimits(v []*DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesOrginInterRegionBandwidthLimits) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages {
	s.OrginInterRegionBandwidthLimits = v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages) SetCenIds(v []*string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages {
	s.CenIds = v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages) SetGeographicRegionAId(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages {
	s.GeographicRegionAId = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages) SetRatio(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages {
	s.Ratio = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages) SetReservationInternetChargeType(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages {
	s.ReservationInternetChargeType = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages) SetTypeFor95(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages {
	s.TypeFor95 = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages) SetBandwidth(v int64) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages {
	s.Bandwidth = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages) SetDescription(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages {
	s.Description = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages) SetExpiredTime(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages) SetReservationBandwidth(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages {
	s.ReservationBandwidth = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages) SetGeographicSpanId(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages {
	s.GeographicSpanId = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages) SetGeographicRegionBId(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages {
	s.GeographicRegionBId = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages) SetIsCrossBorder(v bool) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages {
	s.IsCrossBorder = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages) SetBusinessStatus(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages {
	s.BusinessStatus = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages) SetName(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages {
	s.Name = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages) SetHasReservationData(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages {
	s.HasReservationData = &v
	return s
}

type DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesOrginInterRegionBandwidthLimits struct {
	OppositeRegionId *string `json:"OppositeRegionId,omitempty" xml:"OppositeRegionId,omitempty"`
	GeographicSpanId *string `json:"GeographicSpanId,omitempty" xml:"GeographicSpanId,omitempty"`
	LocalRegionId    *string `json:"LocalRegionId,omitempty" xml:"LocalRegionId,omitempty"`
	BandwidthLimit   *string `json:"BandwidthLimit,omitempty" xml:"BandwidthLimit,omitempty"`
}

func (s DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesOrginInterRegionBandwidthLimits) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesOrginInterRegionBandwidthLimits) GoString() string {
	return s.String()
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesOrginInterRegionBandwidthLimits) SetOppositeRegionId(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesOrginInterRegionBandwidthLimits {
	s.OppositeRegionId = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesOrginInterRegionBandwidthLimits) SetGeographicSpanId(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesOrginInterRegionBandwidthLimits {
	s.GeographicSpanId = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesOrginInterRegionBandwidthLimits) SetLocalRegionId(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesOrginInterRegionBandwidthLimits {
	s.LocalRegionId = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesOrginInterRegionBandwidthLimits) SetBandwidthLimit(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesOrginInterRegionBandwidthLimits {
	s.BandwidthLimit = &v
	return s
}

type DescribeCenBandwidthPackagesResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCenBandwidthPackagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCenBandwidthPackagesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenBandwidthPackagesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCenBandwidthPackagesResponse) SetHeaders(v map[string]*string) *DescribeCenBandwidthPackagesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCenBandwidthPackagesResponse) SetBody(v *DescribeCenBandwidthPackagesResponseBody) *DescribeCenBandwidthPackagesResponse {
	s.Body = v
	return s
}

type DescribeCenChildInstanceRouteEntriesRequest struct {
	OwnerAccount              *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount      *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId           *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	CenId                     *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ChildInstanceId           *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	ChildInstanceType         *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	Status                    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	PageNumber                *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize                  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ChildInstanceRegionId     *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	ChildInstanceRouteTableId *string `json:"ChildInstanceRouteTableId,omitempty" xml:"ChildInstanceRouteTableId,omitempty"`
	DestinationCidrBlock      *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
}

func (s DescribeCenChildInstanceRouteEntriesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenChildInstanceRouteEntriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) SetOwnerAccount(v string) *DescribeCenChildInstanceRouteEntriesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) SetOwnerId(v int64) *DescribeCenChildInstanceRouteEntriesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) SetResourceOwnerAccount(v string) *DescribeCenChildInstanceRouteEntriesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) SetResourceOwnerId(v int64) *DescribeCenChildInstanceRouteEntriesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) SetCenId(v string) *DescribeCenChildInstanceRouteEntriesRequest {
	s.CenId = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) SetChildInstanceId(v string) *DescribeCenChildInstanceRouteEntriesRequest {
	s.ChildInstanceId = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) SetChildInstanceType(v string) *DescribeCenChildInstanceRouteEntriesRequest {
	s.ChildInstanceType = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) SetStatus(v string) *DescribeCenChildInstanceRouteEntriesRequest {
	s.Status = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) SetPageNumber(v int32) *DescribeCenChildInstanceRouteEntriesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) SetPageSize(v int32) *DescribeCenChildInstanceRouteEntriesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) SetChildInstanceRegionId(v string) *DescribeCenChildInstanceRouteEntriesRequest {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) SetChildInstanceRouteTableId(v string) *DescribeCenChildInstanceRouteEntriesRequest {
	s.ChildInstanceRouteTableId = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) SetDestinationCidrBlock(v string) *DescribeCenChildInstanceRouteEntriesRequest {
	s.DestinationCidrBlock = &v
	return s
}

type DescribeCenChildInstanceRouteEntriesResponseBody struct {
	TotalCount      *int32                                                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize        *int32                                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId       *string                                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber      *int32                                                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	CenRouteEntries []*DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries `json:"CenRouteEntries,omitempty" xml:"CenRouteEntries,omitempty" type:"Repeated"`
}

func (s DescribeCenChildInstanceRouteEntriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenChildInstanceRouteEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBody) SetTotalCount(v int32) *DescribeCenChildInstanceRouteEntriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBody) SetPageSize(v int32) *DescribeCenChildInstanceRouteEntriesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBody) SetRequestId(v string) *DescribeCenChildInstanceRouteEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBody) SetPageNumber(v int32) *DescribeCenChildInstanceRouteEntriesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBody) SetCenRouteEntries(v []*DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries) *DescribeCenChildInstanceRouteEntriesResponseBody {
	s.CenRouteEntries = v
	return s
}

type DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries struct {
	Status               *string                                                                              `json:"Status,omitempty" xml:"Status,omitempty"`
	Type                 *string                                                                              `json:"Type,omitempty" xml:"Type,omitempty"`
	PublishStatus        *string                                                                              `json:"PublishStatus,omitempty" xml:"PublishStatus,omitempty"`
	NextHopType          *string                                                                              `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	OperationalMode      *bool                                                                                `json:"OperationalMode,omitempty" xml:"OperationalMode,omitempty"`
	NextHopRegionId      *string                                                                              `json:"NextHopRegionId,omitempty" xml:"NextHopRegionId,omitempty"`
	NextHopInstanceId    *string                                                                              `json:"NextHopInstanceId,omitempty" xml:"NextHopInstanceId,omitempty"`
	RouteTableId         *string                                                                              `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
	AsPaths              []*string                                                                            `json:"AsPaths,omitempty" xml:"AsPaths,omitempty" type:"Repeated"`
	Conflicts            []*DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesConflicts          `json:"Conflicts,omitempty" xml:"Conflicts,omitempty" type:"Repeated"`
	Communities          []*string                                                                            `json:"Communities,omitempty" xml:"Communities,omitempty" type:"Repeated"`
	DestinationCidrBlock *string                                                                              `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	CenRouteMapRecords   []*DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteMapRecords `json:"CenRouteMapRecords,omitempty" xml:"CenRouteMapRecords,omitempty" type:"Repeated"`
}

func (s DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries) GoString() string {
	return s.String()
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries) SetStatus(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries {
	s.Status = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries) SetType(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries {
	s.Type = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries) SetPublishStatus(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries {
	s.PublishStatus = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries) SetNextHopType(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries {
	s.NextHopType = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries) SetOperationalMode(v bool) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries {
	s.OperationalMode = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries) SetNextHopRegionId(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries {
	s.NextHopRegionId = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries) SetNextHopInstanceId(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries {
	s.NextHopInstanceId = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries) SetRouteTableId(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries {
	s.RouteTableId = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries) SetAsPaths(v []*string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries {
	s.AsPaths = v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries) SetConflicts(v []*DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesConflicts) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries {
	s.Conflicts = v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries) SetCommunities(v []*string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries {
	s.Communities = v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries) SetDestinationCidrBlock(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries) SetCenRouteMapRecords(v []*DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteMapRecords) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries {
	s.CenRouteMapRecords = v
	return s
}

type DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesConflicts struct {
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType         *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesConflicts) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesConflicts) GoString() string {
	return s.String()
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesConflicts) SetStatus(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesConflicts {
	s.Status = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesConflicts) SetDestinationCidrBlock(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesConflicts {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesConflicts) SetInstanceId(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesConflicts {
	s.InstanceId = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesConflicts) SetInstanceType(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesConflicts {
	s.InstanceType = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesConflicts) SetRegionId(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesConflicts {
	s.RegionId = &v
	return s
}

type DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteMapRecords struct {
	RouteMapId *string `json:"RouteMapId,omitempty" xml:"RouteMapId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteMapRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteMapRecords) GoString() string {
	return s.String()
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteMapRecords) SetRouteMapId(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteMapRecords {
	s.RouteMapId = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteMapRecords) SetRegionId(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteMapRecords {
	s.RegionId = &v
	return s
}

type DescribeCenChildInstanceRouteEntriesResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCenChildInstanceRouteEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCenChildInstanceRouteEntriesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenChildInstanceRouteEntriesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCenChildInstanceRouteEntriesResponse) SetHeaders(v map[string]*string) *DescribeCenChildInstanceRouteEntriesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponse) SetBody(v *DescribeCenChildInstanceRouteEntriesResponseBody) *DescribeCenChildInstanceRouteEntriesResponse {
	s.Body = v
	return s
}

type DescribeCenGeographicSpanRemainingBandwidthRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	GeographicRegionAId  *string `json:"GeographicRegionAId,omitempty" xml:"GeographicRegionAId,omitempty"`
	GeographicRegionBId  *string `json:"GeographicRegionBId,omitempty" xml:"GeographicRegionBId,omitempty"`
}

func (s DescribeCenGeographicSpanRemainingBandwidthRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenGeographicSpanRemainingBandwidthRequest) GoString() string {
	return s.String()
}

func (s *DescribeCenGeographicSpanRemainingBandwidthRequest) SetOwnerAccount(v string) *DescribeCenGeographicSpanRemainingBandwidthRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCenGeographicSpanRemainingBandwidthRequest) SetOwnerId(v int64) *DescribeCenGeographicSpanRemainingBandwidthRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCenGeographicSpanRemainingBandwidthRequest) SetResourceOwnerAccount(v string) *DescribeCenGeographicSpanRemainingBandwidthRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCenGeographicSpanRemainingBandwidthRequest) SetResourceOwnerId(v int64) *DescribeCenGeographicSpanRemainingBandwidthRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCenGeographicSpanRemainingBandwidthRequest) SetPageNumber(v int32) *DescribeCenGeographicSpanRemainingBandwidthRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCenGeographicSpanRemainingBandwidthRequest) SetPageSize(v int32) *DescribeCenGeographicSpanRemainingBandwidthRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCenGeographicSpanRemainingBandwidthRequest) SetCenId(v string) *DescribeCenGeographicSpanRemainingBandwidthRequest {
	s.CenId = &v
	return s
}

func (s *DescribeCenGeographicSpanRemainingBandwidthRequest) SetGeographicRegionAId(v string) *DescribeCenGeographicSpanRemainingBandwidthRequest {
	s.GeographicRegionAId = &v
	return s
}

func (s *DescribeCenGeographicSpanRemainingBandwidthRequest) SetGeographicRegionBId(v string) *DescribeCenGeographicSpanRemainingBandwidthRequest {
	s.GeographicRegionBId = &v
	return s
}

type DescribeCenGeographicSpanRemainingBandwidthResponseBody struct {
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RemainingBandwidth *int64  `json:"RemainingBandwidth,omitempty" xml:"RemainingBandwidth,omitempty"`
}

func (s DescribeCenGeographicSpanRemainingBandwidthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenGeographicSpanRemainingBandwidthResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCenGeographicSpanRemainingBandwidthResponseBody) SetRequestId(v string) *DescribeCenGeographicSpanRemainingBandwidthResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCenGeographicSpanRemainingBandwidthResponseBody) SetRemainingBandwidth(v int64) *DescribeCenGeographicSpanRemainingBandwidthResponseBody {
	s.RemainingBandwidth = &v
	return s
}

type DescribeCenGeographicSpanRemainingBandwidthResponse struct {
	Headers map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCenGeographicSpanRemainingBandwidthResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCenGeographicSpanRemainingBandwidthResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenGeographicSpanRemainingBandwidthResponse) GoString() string {
	return s.String()
}

func (s *DescribeCenGeographicSpanRemainingBandwidthResponse) SetHeaders(v map[string]*string) *DescribeCenGeographicSpanRemainingBandwidthResponse {
	s.Headers = v
	return s
}

func (s *DescribeCenGeographicSpanRemainingBandwidthResponse) SetBody(v *DescribeCenGeographicSpanRemainingBandwidthResponseBody) *DescribeCenGeographicSpanRemainingBandwidthResponse {
	s.Body = v
	return s
}

type DescribeCenGeographicSpansRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	GeographicSpanId     *string `json:"GeographicSpanId,omitempty" xml:"GeographicSpanId,omitempty"`
}

func (s DescribeCenGeographicSpansRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenGeographicSpansRequest) GoString() string {
	return s.String()
}

func (s *DescribeCenGeographicSpansRequest) SetOwnerAccount(v string) *DescribeCenGeographicSpansRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCenGeographicSpansRequest) SetOwnerId(v int64) *DescribeCenGeographicSpansRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCenGeographicSpansRequest) SetResourceOwnerAccount(v string) *DescribeCenGeographicSpansRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCenGeographicSpansRequest) SetResourceOwnerId(v int64) *DescribeCenGeographicSpansRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCenGeographicSpansRequest) SetPageNumber(v int32) *DescribeCenGeographicSpansRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCenGeographicSpansRequest) SetPageSize(v int32) *DescribeCenGeographicSpansRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCenGeographicSpansRequest) SetGeographicSpanId(v string) *DescribeCenGeographicSpansRequest {
	s.GeographicSpanId = &v
	return s
}

type DescribeCenGeographicSpansResponseBody struct {
	TotalCount           *int32                                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize             *int32                                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId            *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber           *int32                                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	GeographicSpanModels []*DescribeCenGeographicSpansResponseBodyGeographicSpanModels `json:"GeographicSpanModels,omitempty" xml:"GeographicSpanModels,omitempty" type:"Repeated"`
}

func (s DescribeCenGeographicSpansResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenGeographicSpansResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCenGeographicSpansResponseBody) SetTotalCount(v int32) *DescribeCenGeographicSpansResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCenGeographicSpansResponseBody) SetPageSize(v int32) *DescribeCenGeographicSpansResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCenGeographicSpansResponseBody) SetRequestId(v string) *DescribeCenGeographicSpansResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCenGeographicSpansResponseBody) SetPageNumber(v int32) *DescribeCenGeographicSpansResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCenGeographicSpansResponseBody) SetGeographicSpanModels(v []*DescribeCenGeographicSpansResponseBodyGeographicSpanModels) *DescribeCenGeographicSpansResponseBody {
	s.GeographicSpanModels = v
	return s
}

type DescribeCenGeographicSpansResponseBodyGeographicSpanModels struct {
	GeographicSpanId    *string `json:"GeographicSpanId,omitempty" xml:"GeographicSpanId,omitempty"`
	OppositeGeoRegionId *string `json:"OppositeGeoRegionId,omitempty" xml:"OppositeGeoRegionId,omitempty"`
	LocalGeoRegionId    *string `json:"LocalGeoRegionId,omitempty" xml:"LocalGeoRegionId,omitempty"`
}

func (s DescribeCenGeographicSpansResponseBodyGeographicSpanModels) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenGeographicSpansResponseBodyGeographicSpanModels) GoString() string {
	return s.String()
}

func (s *DescribeCenGeographicSpansResponseBodyGeographicSpanModels) SetGeographicSpanId(v string) *DescribeCenGeographicSpansResponseBodyGeographicSpanModels {
	s.GeographicSpanId = &v
	return s
}

func (s *DescribeCenGeographicSpansResponseBodyGeographicSpanModels) SetOppositeGeoRegionId(v string) *DescribeCenGeographicSpansResponseBodyGeographicSpanModels {
	s.OppositeGeoRegionId = &v
	return s
}

func (s *DescribeCenGeographicSpansResponseBodyGeographicSpanModels) SetLocalGeoRegionId(v string) *DescribeCenGeographicSpansResponseBodyGeographicSpanModels {
	s.LocalGeoRegionId = &v
	return s
}

type DescribeCenGeographicSpansResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCenGeographicSpansResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCenGeographicSpansResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenGeographicSpansResponse) GoString() string {
	return s.String()
}

func (s *DescribeCenGeographicSpansResponse) SetHeaders(v map[string]*string) *DescribeCenGeographicSpansResponse {
	s.Headers = v
	return s
}

func (s *DescribeCenGeographicSpansResponse) SetBody(v *DescribeCenGeographicSpansResponseBody) *DescribeCenGeographicSpansResponse {
	s.Body = v
	return s
}

type DescribeCenInterRegionBandwidthLimitsRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
}

func (s DescribeCenInterRegionBandwidthLimitsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenInterRegionBandwidthLimitsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCenInterRegionBandwidthLimitsRequest) SetOwnerAccount(v string) *DescribeCenInterRegionBandwidthLimitsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsRequest) SetOwnerId(v int64) *DescribeCenInterRegionBandwidthLimitsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsRequest) SetResourceOwnerAccount(v string) *DescribeCenInterRegionBandwidthLimitsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsRequest) SetResourceOwnerId(v int64) *DescribeCenInterRegionBandwidthLimitsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsRequest) SetPageNumber(v int32) *DescribeCenInterRegionBandwidthLimitsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsRequest) SetPageSize(v int32) *DescribeCenInterRegionBandwidthLimitsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsRequest) SetCenId(v string) *DescribeCenInterRegionBandwidthLimitsRequest {
	s.CenId = &v
	return s
}

type DescribeCenInterRegionBandwidthLimitsResponseBody struct {
	CenInterRegionBandwidthLimits []*DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimits `json:"CenInterRegionBandwidthLimits,omitempty" xml:"CenInterRegionBandwidthLimits,omitempty" type:"Repeated"`
	TotalCount                    *int32                                                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize                      *int32                                                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId                     *string                                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber                    *int32                                                                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeCenInterRegionBandwidthLimitsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenInterRegionBandwidthLimitsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBody) SetCenInterRegionBandwidthLimits(v []*DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimits) *DescribeCenInterRegionBandwidthLimitsResponseBody {
	s.CenInterRegionBandwidthLimits = v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBody) SetTotalCount(v int32) *DescribeCenInterRegionBandwidthLimitsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBody) SetPageSize(v int32) *DescribeCenInterRegionBandwidthLimitsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBody) SetRequestId(v string) *DescribeCenInterRegionBandwidthLimitsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBody) SetPageNumber(v int32) *DescribeCenInterRegionBandwidthLimitsResponseBody {
	s.PageNumber = &v
	return s
}

type DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimits struct {
	Status             *string `json:"Status,omitempty" xml:"Status,omitempty"`
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	OppositeRegionId   *string `json:"OppositeRegionId,omitempty" xml:"OppositeRegionId,omitempty"`
	GeographicSpanId   *string `json:"GeographicSpanId,omitempty" xml:"GeographicSpanId,omitempty"`
	CenId              *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	LocalRegionId      *string `json:"LocalRegionId,omitempty" xml:"LocalRegionId,omitempty"`
	BandwidthLimit     *int64  `json:"BandwidthLimit,omitempty" xml:"BandwidthLimit,omitempty"`
}

func (s DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimits) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimits) GoString() string {
	return s.String()
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimits) SetStatus(v string) *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimits {
	s.Status = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimits) SetBandwidthPackageId(v string) *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimits {
	s.BandwidthPackageId = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimits) SetOppositeRegionId(v string) *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimits {
	s.OppositeRegionId = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimits) SetGeographicSpanId(v string) *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimits {
	s.GeographicSpanId = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimits) SetCenId(v string) *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimits {
	s.CenId = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimits) SetLocalRegionId(v string) *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimits {
	s.LocalRegionId = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimits) SetBandwidthLimit(v int64) *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimits {
	s.BandwidthLimit = &v
	return s
}

type DescribeCenInterRegionBandwidthLimitsResponse struct {
	Headers map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCenInterRegionBandwidthLimitsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCenInterRegionBandwidthLimitsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenInterRegionBandwidthLimitsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCenInterRegionBandwidthLimitsResponse) SetHeaders(v map[string]*string) *DescribeCenInterRegionBandwidthLimitsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsResponse) SetBody(v *DescribeCenInterRegionBandwidthLimitsResponseBody) *DescribeCenInterRegionBandwidthLimitsResponse {
	s.Body = v
	return s
}

type DescribeCenPrivateZoneRoutesRequest struct {
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	AccessRegionId       *string `json:"AccessRegionId,omitempty" xml:"AccessRegionId,omitempty"`
	HostRegionId         *string `json:"HostRegionId,omitempty" xml:"HostRegionId,omitempty"`
}

func (s DescribeCenPrivateZoneRoutesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenPrivateZoneRoutesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCenPrivateZoneRoutesRequest) SetResourceOwnerAccount(v string) *DescribeCenPrivateZoneRoutesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesRequest) SetResourceOwnerId(v int64) *DescribeCenPrivateZoneRoutesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesRequest) SetPageNumber(v int32) *DescribeCenPrivateZoneRoutesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesRequest) SetPageSize(v int32) *DescribeCenPrivateZoneRoutesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesRequest) SetCenId(v string) *DescribeCenPrivateZoneRoutesRequest {
	s.CenId = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesRequest) SetAccessRegionId(v string) *DescribeCenPrivateZoneRoutesRequest {
	s.AccessRegionId = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesRequest) SetHostRegionId(v string) *DescribeCenPrivateZoneRoutesRequest {
	s.HostRegionId = &v
	return s
}

type DescribeCenPrivateZoneRoutesResponseBody struct {
	TotalCount            *int32                                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId             *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize              *int32                                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PrivateZoneInfos      []*DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfos `json:"PrivateZoneInfos,omitempty" xml:"PrivateZoneInfos,omitempty" type:"Repeated"`
	PageNumber            *int32                                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	CenId                 *string                                                     `json:"CenId,omitempty" xml:"CenId,omitempty"`
	PrivateZoneDnsServers *string                                                     `json:"PrivateZoneDnsServers,omitempty" xml:"PrivateZoneDnsServers,omitempty"`
}

func (s DescribeCenPrivateZoneRoutesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenPrivateZoneRoutesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCenPrivateZoneRoutesResponseBody) SetTotalCount(v int32) *DescribeCenPrivateZoneRoutesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesResponseBody) SetRequestId(v string) *DescribeCenPrivateZoneRoutesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesResponseBody) SetPageSize(v int32) *DescribeCenPrivateZoneRoutesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesResponseBody) SetPrivateZoneInfos(v []*DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfos) *DescribeCenPrivateZoneRoutesResponseBody {
	s.PrivateZoneInfos = v
	return s
}

func (s *DescribeCenPrivateZoneRoutesResponseBody) SetPageNumber(v int32) *DescribeCenPrivateZoneRoutesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesResponseBody) SetCenId(v string) *DescribeCenPrivateZoneRoutesResponseBody {
	s.CenId = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesResponseBody) SetPrivateZoneDnsServers(v string) *DescribeCenPrivateZoneRoutesResponseBody {
	s.PrivateZoneDnsServers = &v
	return s
}

type DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfos struct {
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	HostVpcId      *string `json:"HostVpcId,omitempty" xml:"HostVpcId,omitempty"`
	AccessRegionId *string `json:"AccessRegionId,omitempty" xml:"AccessRegionId,omitempty"`
	HostRegionId   *string `json:"HostRegionId,omitempty" xml:"HostRegionId,omitempty"`
}

func (s DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfos) GoString() string {
	return s.String()
}

func (s *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfos) SetStatus(v string) *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfos {
	s.Status = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfos) SetHostVpcId(v string) *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfos {
	s.HostVpcId = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfos) SetAccessRegionId(v string) *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfos {
	s.AccessRegionId = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfos) SetHostRegionId(v string) *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfos {
	s.HostRegionId = &v
	return s
}

type DescribeCenPrivateZoneRoutesResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCenPrivateZoneRoutesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCenPrivateZoneRoutesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenPrivateZoneRoutesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCenPrivateZoneRoutesResponse) SetHeaders(v map[string]*string) *DescribeCenPrivateZoneRoutesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCenPrivateZoneRoutesResponse) SetBody(v *DescribeCenPrivateZoneRoutesResponseBody) *DescribeCenPrivateZoneRoutesResponse {
	s.Body = v
	return s
}

type DescribeCenRegionDomainRouteEntriesRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	CenRegionId          *string `json:"CenRegionId,omitempty" xml:"CenRegionId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
}

func (s DescribeCenRegionDomainRouteEntriesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenRegionDomainRouteEntriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCenRegionDomainRouteEntriesRequest) SetOwnerAccount(v string) *DescribeCenRegionDomainRouteEntriesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesRequest) SetOwnerId(v int64) *DescribeCenRegionDomainRouteEntriesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesRequest) SetResourceOwnerAccount(v string) *DescribeCenRegionDomainRouteEntriesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesRequest) SetResourceOwnerId(v int64) *DescribeCenRegionDomainRouteEntriesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesRequest) SetCenId(v string) *DescribeCenRegionDomainRouteEntriesRequest {
	s.CenId = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesRequest) SetCenRegionId(v string) *DescribeCenRegionDomainRouteEntriesRequest {
	s.CenRegionId = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesRequest) SetPageNumber(v int32) *DescribeCenRegionDomainRouteEntriesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesRequest) SetPageSize(v int32) *DescribeCenRegionDomainRouteEntriesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesRequest) SetStatus(v string) *DescribeCenRegionDomainRouteEntriesRequest {
	s.Status = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesRequest) SetDestinationCidrBlock(v string) *DescribeCenRegionDomainRouteEntriesRequest {
	s.DestinationCidrBlock = &v
	return s
}

type DescribeCenRegionDomainRouteEntriesResponseBody struct {
	TotalCount      *int32                                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize        *int32                                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId       *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber      *int32                                                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	CenRouteEntries []*DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries `json:"CenRouteEntries,omitempty" xml:"CenRouteEntries,omitempty" type:"Repeated"`
}

func (s DescribeCenRegionDomainRouteEntriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenRegionDomainRouteEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBody) SetTotalCount(v int32) *DescribeCenRegionDomainRouteEntriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBody) SetPageSize(v int32) *DescribeCenRegionDomainRouteEntriesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBody) SetRequestId(v string) *DescribeCenRegionDomainRouteEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBody) SetPageNumber(v int32) *DescribeCenRegionDomainRouteEntriesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBody) SetCenRouteEntries(v []*DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries) *DescribeCenRegionDomainRouteEntriesResponseBody {
	s.CenRouteEntries = v
	return s
}

type DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries struct {
	Type                  *string                                                                                `json:"Type,omitempty" xml:"Type,omitempty"`
	Status                *string                                                                                `json:"Status,omitempty" xml:"Status,omitempty"`
	CenOutRouteMapRecords []*DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenOutRouteMapRecords `json:"CenOutRouteMapRecords,omitempty" xml:"CenOutRouteMapRecords,omitempty" type:"Repeated"`
	NextHopType           *string                                                                                `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	NextHopInstanceId     *string                                                                                `json:"NextHopInstanceId,omitempty" xml:"NextHopInstanceId,omitempty"`
	NextHopRegionId       *string                                                                                `json:"NextHopRegionId,omitempty" xml:"NextHopRegionId,omitempty"`
	AsPaths               []*string                                                                              `json:"AsPaths,omitempty" xml:"AsPaths,omitempty" type:"Repeated"`
	ToOtherRegionStatus   *string                                                                                `json:"ToOtherRegionStatus,omitempty" xml:"ToOtherRegionStatus,omitempty"`
	Communities           []*string                                                                              `json:"Communities,omitempty" xml:"Communities,omitempty" type:"Repeated"`
	DestinationCidrBlock  *string                                                                                `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	Preference            *int32                                                                                 `json:"Preference,omitempty" xml:"Preference,omitempty"`
	CenRouteMapRecords    []*DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteMapRecords    `json:"CenRouteMapRecords,omitempty" xml:"CenRouteMapRecords,omitempty" type:"Repeated"`
}

func (s DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries) GoString() string {
	return s.String()
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries) SetType(v string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries {
	s.Type = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries) SetStatus(v string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries {
	s.Status = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries) SetCenOutRouteMapRecords(v []*DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenOutRouteMapRecords) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries {
	s.CenOutRouteMapRecords = v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries) SetNextHopType(v string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries {
	s.NextHopType = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries) SetNextHopInstanceId(v string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries {
	s.NextHopInstanceId = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries) SetNextHopRegionId(v string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries {
	s.NextHopRegionId = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries) SetAsPaths(v []*string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries {
	s.AsPaths = v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries) SetToOtherRegionStatus(v string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries {
	s.ToOtherRegionStatus = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries) SetCommunities(v []*string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries {
	s.Communities = v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries) SetDestinationCidrBlock(v string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries) SetPreference(v int32) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries {
	s.Preference = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries) SetCenRouteMapRecords(v []*DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteMapRecords) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries {
	s.CenRouteMapRecords = v
	return s
}

type DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenOutRouteMapRecords struct {
	RouteMapId *string `json:"RouteMapId,omitempty" xml:"RouteMapId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenOutRouteMapRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenOutRouteMapRecords) GoString() string {
	return s.String()
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenOutRouteMapRecords) SetRouteMapId(v string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenOutRouteMapRecords {
	s.RouteMapId = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenOutRouteMapRecords) SetRegionId(v string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenOutRouteMapRecords {
	s.RegionId = &v
	return s
}

type DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteMapRecords struct {
	RouteMapId *string `json:"RouteMapId,omitempty" xml:"RouteMapId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteMapRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteMapRecords) GoString() string {
	return s.String()
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteMapRecords) SetRouteMapId(v string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteMapRecords {
	s.RouteMapId = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteMapRecords) SetRegionId(v string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteMapRecords {
	s.RegionId = &v
	return s
}

type DescribeCenRegionDomainRouteEntriesResponse struct {
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCenRegionDomainRouteEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCenRegionDomainRouteEntriesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenRegionDomainRouteEntriesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCenRegionDomainRouteEntriesResponse) SetHeaders(v map[string]*string) *DescribeCenRegionDomainRouteEntriesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponse) SetBody(v *DescribeCenRegionDomainRouteEntriesResponseBody) *DescribeCenRegionDomainRouteEntriesResponse {
	s.Body = v
	return s
}

type DescribeCenRouteMapsRequest struct {
	OwnerAccount         *string                              `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64                               `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string                              `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                               `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PageNumber           *int32                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CenId                *string                              `json:"CenId,omitempty" xml:"CenId,omitempty"`
	RouteMapId           *string                              `json:"RouteMapId,omitempty" xml:"RouteMapId,omitempty"`
	CenRegionId          *string                              `json:"CenRegionId,omitempty" xml:"CenRegionId,omitempty"`
	TransmitDirection    *string                              `json:"TransmitDirection,omitempty" xml:"TransmitDirection,omitempty"`
	Filter               []*DescribeCenRouteMapsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
}

func (s DescribeCenRouteMapsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenRouteMapsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCenRouteMapsRequest) SetOwnerAccount(v string) *DescribeCenRouteMapsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCenRouteMapsRequest) SetOwnerId(v int64) *DescribeCenRouteMapsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCenRouteMapsRequest) SetResourceOwnerAccount(v string) *DescribeCenRouteMapsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCenRouteMapsRequest) SetResourceOwnerId(v int64) *DescribeCenRouteMapsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCenRouteMapsRequest) SetPageNumber(v int32) *DescribeCenRouteMapsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCenRouteMapsRequest) SetPageSize(v int32) *DescribeCenRouteMapsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCenRouteMapsRequest) SetCenId(v string) *DescribeCenRouteMapsRequest {
	s.CenId = &v
	return s
}

func (s *DescribeCenRouteMapsRequest) SetRouteMapId(v string) *DescribeCenRouteMapsRequest {
	s.RouteMapId = &v
	return s
}

func (s *DescribeCenRouteMapsRequest) SetCenRegionId(v string) *DescribeCenRouteMapsRequest {
	s.CenRegionId = &v
	return s
}

func (s *DescribeCenRouteMapsRequest) SetTransmitDirection(v string) *DescribeCenRouteMapsRequest {
	s.TransmitDirection = &v
	return s
}

func (s *DescribeCenRouteMapsRequest) SetFilter(v []*DescribeCenRouteMapsRequestFilter) *DescribeCenRouteMapsRequest {
	s.Filter = v
	return s
}

type DescribeCenRouteMapsRequestFilter struct {
}

func (s DescribeCenRouteMapsRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenRouteMapsRequestFilter) GoString() string {
	return s.String()
}

type DescribeCenRouteMapsResponseBody struct {
	TotalCount *int32                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	RouteMaps  []*DescribeCenRouteMapsResponseBodyRouteMaps `json:"RouteMaps,omitempty" xml:"RouteMaps,omitempty" type:"Repeated"`
}

func (s DescribeCenRouteMapsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenRouteMapsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCenRouteMapsResponseBody) SetTotalCount(v int32) *DescribeCenRouteMapsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBody) SetPageSize(v int32) *DescribeCenRouteMapsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBody) SetRequestId(v string) *DescribeCenRouteMapsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBody) SetPageNumber(v int32) *DescribeCenRouteMapsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBody) SetRouteMaps(v []*DescribeCenRouteMapsResponseBodyRouteMaps) *DescribeCenRouteMapsResponseBody {
	s.RouteMaps = v
	return s
}

type DescribeCenRouteMapsResponseBodyRouteMaps struct {
	Status                             *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	SourceInstanceIdsReverseMatch      *bool     `json:"SourceInstanceIdsReverseMatch,omitempty" xml:"SourceInstanceIdsReverseMatch,omitempty"`
	SourceRegionIds                    []*string `json:"SourceRegionIds,omitempty" xml:"SourceRegionIds,omitempty" type:"Repeated"`
	MatchCommunitySet                  []*string `json:"MatchCommunitySet,omitempty" xml:"MatchCommunitySet,omitempty" type:"Repeated"`
	Priority                           *int32    `json:"Priority,omitempty" xml:"Priority,omitempty"`
	CommunityOperateMode               *string   `json:"CommunityOperateMode,omitempty" xml:"CommunityOperateMode,omitempty"`
	PrependAsPath                      []*string `json:"PrependAsPath,omitempty" xml:"PrependAsPath,omitempty" type:"Repeated"`
	RouteTypes                         []*string `json:"RouteTypes,omitempty" xml:"RouteTypes,omitempty" type:"Repeated"`
	Description                        *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	DestinationInstanceIds             []*string `json:"DestinationInstanceIds,omitempty" xml:"DestinationInstanceIds,omitempty" type:"Repeated"`
	MatchAsns                          []*string `json:"MatchAsns,omitempty" xml:"MatchAsns,omitempty" type:"Repeated"`
	DestinationInstanceIdsReverseMatch *bool     `json:"DestinationInstanceIdsReverseMatch,omitempty" xml:"DestinationInstanceIdsReverseMatch,omitempty"`
	OperateCommunitySet                []*string `json:"OperateCommunitySet,omitempty" xml:"OperateCommunitySet,omitempty" type:"Repeated"`
	NextPriority                       *int32    `json:"NextPriority,omitempty" xml:"NextPriority,omitempty"`
	RouteMapId                         *string   `json:"RouteMapId,omitempty" xml:"RouteMapId,omitempty"`
	TransmitDirection                  *string   `json:"TransmitDirection,omitempty" xml:"TransmitDirection,omitempty"`
	SourceChildInstanceTypes           []*string `json:"SourceChildInstanceTypes,omitempty" xml:"SourceChildInstanceTypes,omitempty" type:"Repeated"`
	DestinationRouteTableIds           []*string `json:"DestinationRouteTableIds,omitempty" xml:"DestinationRouteTableIds,omitempty" type:"Repeated"`
	SourceInstanceIds                  []*string `json:"SourceInstanceIds,omitempty" xml:"SourceInstanceIds,omitempty" type:"Repeated"`
	CenRegionId                        *string   `json:"CenRegionId,omitempty" xml:"CenRegionId,omitempty"`
	DestinationCidrBlocks              []*string `json:"DestinationCidrBlocks,omitempty" xml:"DestinationCidrBlocks,omitempty" type:"Repeated"`
	CenId                              *string   `json:"CenId,omitempty" xml:"CenId,omitempty"`
	SourceRouteTableIds                []*string `json:"SourceRouteTableIds,omitempty" xml:"SourceRouteTableIds,omitempty" type:"Repeated"`
	MapResult                          *string   `json:"MapResult,omitempty" xml:"MapResult,omitempty"`
	CommunityMatchMode                 *string   `json:"CommunityMatchMode,omitempty" xml:"CommunityMatchMode,omitempty"`
	DestinationChildInstanceTypes      []*string `json:"DestinationChildInstanceTypes,omitempty" xml:"DestinationChildInstanceTypes,omitempty" type:"Repeated"`
	AsPathMatchMode                    *string   `json:"AsPathMatchMode,omitempty" xml:"AsPathMatchMode,omitempty"`
	Preference                         *int32    `json:"Preference,omitempty" xml:"Preference,omitempty"`
	CidrMatchMode                      *string   `json:"CidrMatchMode,omitempty" xml:"CidrMatchMode,omitempty"`
}

func (s DescribeCenRouteMapsResponseBodyRouteMaps) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenRouteMapsResponseBodyRouteMaps) GoString() string {
	return s.String()
}

func (s *DescribeCenRouteMapsResponseBodyRouteMaps) SetStatus(v string) *DescribeCenRouteMapsResponseBodyRouteMaps {
	s.Status = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMaps) SetSourceInstanceIdsReverseMatch(v bool) *DescribeCenRouteMapsResponseBodyRouteMaps {
	s.SourceInstanceIdsReverseMatch = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMaps) SetSourceRegionIds(v []*string) *DescribeCenRouteMapsResponseBodyRouteMaps {
	s.SourceRegionIds = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMaps) SetMatchCommunitySet(v []*string) *DescribeCenRouteMapsResponseBodyRouteMaps {
	s.MatchCommunitySet = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMaps) SetPriority(v int32) *DescribeCenRouteMapsResponseBodyRouteMaps {
	s.Priority = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMaps) SetCommunityOperateMode(v string) *DescribeCenRouteMapsResponseBodyRouteMaps {
	s.CommunityOperateMode = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMaps) SetPrependAsPath(v []*string) *DescribeCenRouteMapsResponseBodyRouteMaps {
	s.PrependAsPath = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMaps) SetRouteTypes(v []*string) *DescribeCenRouteMapsResponseBodyRouteMaps {
	s.RouteTypes = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMaps) SetDescription(v string) *DescribeCenRouteMapsResponseBodyRouteMaps {
	s.Description = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMaps) SetDestinationInstanceIds(v []*string) *DescribeCenRouteMapsResponseBodyRouteMaps {
	s.DestinationInstanceIds = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMaps) SetMatchAsns(v []*string) *DescribeCenRouteMapsResponseBodyRouteMaps {
	s.MatchAsns = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMaps) SetDestinationInstanceIdsReverseMatch(v bool) *DescribeCenRouteMapsResponseBodyRouteMaps {
	s.DestinationInstanceIdsReverseMatch = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMaps) SetOperateCommunitySet(v []*string) *DescribeCenRouteMapsResponseBodyRouteMaps {
	s.OperateCommunitySet = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMaps) SetNextPriority(v int32) *DescribeCenRouteMapsResponseBodyRouteMaps {
	s.NextPriority = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMaps) SetRouteMapId(v string) *DescribeCenRouteMapsResponseBodyRouteMaps {
	s.RouteMapId = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMaps) SetTransmitDirection(v string) *DescribeCenRouteMapsResponseBodyRouteMaps {
	s.TransmitDirection = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMaps) SetSourceChildInstanceTypes(v []*string) *DescribeCenRouteMapsResponseBodyRouteMaps {
	s.SourceChildInstanceTypes = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMaps) SetDestinationRouteTableIds(v []*string) *DescribeCenRouteMapsResponseBodyRouteMaps {
	s.DestinationRouteTableIds = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMaps) SetSourceInstanceIds(v []*string) *DescribeCenRouteMapsResponseBodyRouteMaps {
	s.SourceInstanceIds = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMaps) SetCenRegionId(v string) *DescribeCenRouteMapsResponseBodyRouteMaps {
	s.CenRegionId = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMaps) SetDestinationCidrBlocks(v []*string) *DescribeCenRouteMapsResponseBodyRouteMaps {
	s.DestinationCidrBlocks = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMaps) SetCenId(v string) *DescribeCenRouteMapsResponseBodyRouteMaps {
	s.CenId = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMaps) SetSourceRouteTableIds(v []*string) *DescribeCenRouteMapsResponseBodyRouteMaps {
	s.SourceRouteTableIds = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMaps) SetMapResult(v string) *DescribeCenRouteMapsResponseBodyRouteMaps {
	s.MapResult = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMaps) SetCommunityMatchMode(v string) *DescribeCenRouteMapsResponseBodyRouteMaps {
	s.CommunityMatchMode = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMaps) SetDestinationChildInstanceTypes(v []*string) *DescribeCenRouteMapsResponseBodyRouteMaps {
	s.DestinationChildInstanceTypes = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMaps) SetAsPathMatchMode(v string) *DescribeCenRouteMapsResponseBodyRouteMaps {
	s.AsPathMatchMode = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMaps) SetPreference(v int32) *DescribeCenRouteMapsResponseBodyRouteMaps {
	s.Preference = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMaps) SetCidrMatchMode(v string) *DescribeCenRouteMapsResponseBodyRouteMaps {
	s.CidrMatchMode = &v
	return s
}

type DescribeCenRouteMapsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCenRouteMapsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCenRouteMapsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenRouteMapsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCenRouteMapsResponse) SetHeaders(v map[string]*string) *DescribeCenRouteMapsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCenRouteMapsResponse) SetBody(v *DescribeCenRouteMapsResponseBody) *DescribeCenRouteMapsResponse {
	s.Body = v
	return s
}

type DescribeCensRequest struct {
	OwnerAccount         *string                      `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64                       `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string                      `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                       `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PageNumber           *int32                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Filter               []*DescribeCensRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	Tag                  []*DescribeCensRequestTag    `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeCensRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCensRequest) GoString() string {
	return s.String()
}

func (s *DescribeCensRequest) SetOwnerAccount(v string) *DescribeCensRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCensRequest) SetOwnerId(v int64) *DescribeCensRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCensRequest) SetResourceOwnerAccount(v string) *DescribeCensRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCensRequest) SetResourceOwnerId(v int64) *DescribeCensRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCensRequest) SetPageNumber(v int32) *DescribeCensRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCensRequest) SetPageSize(v int32) *DescribeCensRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCensRequest) SetFilter(v []*DescribeCensRequestFilter) *DescribeCensRequest {
	s.Filter = v
	return s
}

func (s *DescribeCensRequest) SetTag(v []*DescribeCensRequestTag) *DescribeCensRequest {
	s.Tag = v
	return s
}

type DescribeCensRequestFilter struct {
	Key   *string   `json:"Key,omitempty" xml:"Key,omitempty"`
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s DescribeCensRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s DescribeCensRequestFilter) GoString() string {
	return s.String()
}

func (s *DescribeCensRequestFilter) SetKey(v string) *DescribeCensRequestFilter {
	s.Key = &v
	return s
}

func (s *DescribeCensRequestFilter) SetValue(v []*string) *DescribeCensRequestFilter {
	s.Value = v
	return s
}

type DescribeCensRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeCensRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeCensRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeCensRequestTag) SetKey(v string) *DescribeCensRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeCensRequestTag) SetValue(v string) *DescribeCensRequestTag {
	s.Value = &v
	return s
}

type DescribeCensResponseBody struct {
	TotalCount *int32                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int32                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Cens       []*DescribeCensResponseBodyCens `json:"Cens,omitempty" xml:"Cens,omitempty" type:"Repeated"`
}

func (s DescribeCensResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCensResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCensResponseBody) SetTotalCount(v int32) *DescribeCensResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCensResponseBody) SetPageSize(v int32) *DescribeCensResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCensResponseBody) SetRequestId(v string) *DescribeCensResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCensResponseBody) SetPageNumber(v int32) *DescribeCensResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCensResponseBody) SetCens(v []*DescribeCensResponseBodyCens) *DescribeCensResponseBody {
	s.Cens = v
	return s
}

type DescribeCensResponseBodyCens struct {
	Status                 *string                             `json:"Status,omitempty" xml:"Status,omitempty"`
	CreationTime           *string                             `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	CenBandwidthPackageIds []*string                           `json:"CenBandwidthPackageIds,omitempty" xml:"CenBandwidthPackageIds,omitempty" type:"Repeated"`
	Description            *string                             `json:"Description,omitempty" xml:"Description,omitempty"`
	CenId                  *string                             `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ProtectionLevel        *string                             `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty"`
	Tags                   []*DescribeCensResponseBodyCensTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	Name                   *string                             `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeCensResponseBodyCens) String() string {
	return tea.Prettify(s)
}

func (s DescribeCensResponseBodyCens) GoString() string {
	return s.String()
}

func (s *DescribeCensResponseBodyCens) SetStatus(v string) *DescribeCensResponseBodyCens {
	s.Status = &v
	return s
}

func (s *DescribeCensResponseBodyCens) SetCreationTime(v string) *DescribeCensResponseBodyCens {
	s.CreationTime = &v
	return s
}

func (s *DescribeCensResponseBodyCens) SetCenBandwidthPackageIds(v []*string) *DescribeCensResponseBodyCens {
	s.CenBandwidthPackageIds = v
	return s
}

func (s *DescribeCensResponseBodyCens) SetDescription(v string) *DescribeCensResponseBodyCens {
	s.Description = &v
	return s
}

func (s *DescribeCensResponseBodyCens) SetCenId(v string) *DescribeCensResponseBodyCens {
	s.CenId = &v
	return s
}

func (s *DescribeCensResponseBodyCens) SetProtectionLevel(v string) *DescribeCensResponseBodyCens {
	s.ProtectionLevel = &v
	return s
}

func (s *DescribeCensResponseBodyCens) SetTags(v []*DescribeCensResponseBodyCensTags) *DescribeCensResponseBodyCens {
	s.Tags = v
	return s
}

func (s *DescribeCensResponseBodyCens) SetName(v string) *DescribeCensResponseBodyCens {
	s.Name = &v
	return s
}

type DescribeCensResponseBodyCensTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeCensResponseBodyCensTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeCensResponseBodyCensTags) GoString() string {
	return s.String()
}

func (s *DescribeCensResponseBodyCensTags) SetKey(v string) *DescribeCensResponseBodyCensTags {
	s.Key = &v
	return s
}

func (s *DescribeCensResponseBodyCensTags) SetValue(v string) *DescribeCensResponseBodyCensTags {
	s.Value = &v
	return s
}

type DescribeCensResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCensResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCensResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCensResponse) GoString() string {
	return s.String()
}

func (s *DescribeCensResponse) SetHeaders(v map[string]*string) *DescribeCensResponse {
	s.Headers = v
	return s
}

func (s *DescribeCensResponse) SetBody(v *DescribeCensResponseBody) *DescribeCensResponse {
	s.Body = v
	return s
}

type DescribeCenVbrHealthCheckRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	VbrInstanceRegionId  *string `json:"VbrInstanceRegionId,omitempty" xml:"VbrInstanceRegionId,omitempty"`
	VbrInstanceId        *string `json:"VbrInstanceId,omitempty" xml:"VbrInstanceId,omitempty"`
	VbrInstanceOwnerId   *int64  `json:"VbrInstanceOwnerId,omitempty" xml:"VbrInstanceOwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeCenVbrHealthCheckRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenVbrHealthCheckRequest) GoString() string {
	return s.String()
}

func (s *DescribeCenVbrHealthCheckRequest) SetOwnerAccount(v string) *DescribeCenVbrHealthCheckRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCenVbrHealthCheckRequest) SetOwnerId(v int64) *DescribeCenVbrHealthCheckRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCenVbrHealthCheckRequest) SetResourceOwnerAccount(v string) *DescribeCenVbrHealthCheckRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCenVbrHealthCheckRequest) SetResourceOwnerId(v int64) *DescribeCenVbrHealthCheckRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCenVbrHealthCheckRequest) SetCenId(v string) *DescribeCenVbrHealthCheckRequest {
	s.CenId = &v
	return s
}

func (s *DescribeCenVbrHealthCheckRequest) SetVbrInstanceRegionId(v string) *DescribeCenVbrHealthCheckRequest {
	s.VbrInstanceRegionId = &v
	return s
}

func (s *DescribeCenVbrHealthCheckRequest) SetVbrInstanceId(v string) *DescribeCenVbrHealthCheckRequest {
	s.VbrInstanceId = &v
	return s
}

func (s *DescribeCenVbrHealthCheckRequest) SetVbrInstanceOwnerId(v int64) *DescribeCenVbrHealthCheckRequest {
	s.VbrInstanceOwnerId = &v
	return s
}

func (s *DescribeCenVbrHealthCheckRequest) SetPageNumber(v int32) *DescribeCenVbrHealthCheckRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCenVbrHealthCheckRequest) SetPageSize(v int32) *DescribeCenVbrHealthCheckRequest {
	s.PageSize = &v
	return s
}

type DescribeCenVbrHealthCheckResponseBody struct {
	TotalCount      *int32                                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize        *int32                                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId       *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber      *int32                                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	VbrHealthChecks []*DescribeCenVbrHealthCheckResponseBodyVbrHealthChecks `json:"VbrHealthChecks,omitempty" xml:"VbrHealthChecks,omitempty" type:"Repeated"`
}

func (s DescribeCenVbrHealthCheckResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenVbrHealthCheckResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCenVbrHealthCheckResponseBody) SetTotalCount(v int32) *DescribeCenVbrHealthCheckResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCenVbrHealthCheckResponseBody) SetPageSize(v int32) *DescribeCenVbrHealthCheckResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCenVbrHealthCheckResponseBody) SetRequestId(v string) *DescribeCenVbrHealthCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCenVbrHealthCheckResponseBody) SetPageNumber(v int32) *DescribeCenVbrHealthCheckResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCenVbrHealthCheckResponseBody) SetVbrHealthChecks(v []*DescribeCenVbrHealthCheckResponseBodyVbrHealthChecks) *DescribeCenVbrHealthCheckResponseBody {
	s.VbrHealthChecks = v
	return s
}

type DescribeCenVbrHealthCheckResponseBodyVbrHealthChecks struct {
	HealthCheckTargetIp *string `json:"HealthCheckTargetIp,omitempty" xml:"HealthCheckTargetIp,omitempty"`
	VbrInstanceId       *string `json:"VbrInstanceId,omitempty" xml:"VbrInstanceId,omitempty"`
	VbrInstanceRegionId *string `json:"VbrInstanceRegionId,omitempty" xml:"VbrInstanceRegionId,omitempty"`
	CenId               *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	HealthyThreshold    *int32  `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	HealthCheckInterval *int32  `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	HealthCheckSourceIp *string `json:"HealthCheckSourceIp,omitempty" xml:"HealthCheckSourceIp,omitempty"`
}

func (s DescribeCenVbrHealthCheckResponseBodyVbrHealthChecks) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenVbrHealthCheckResponseBodyVbrHealthChecks) GoString() string {
	return s.String()
}

func (s *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecks) SetHealthCheckTargetIp(v string) *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecks {
	s.HealthCheckTargetIp = &v
	return s
}

func (s *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecks) SetVbrInstanceId(v string) *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecks {
	s.VbrInstanceId = &v
	return s
}

func (s *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecks) SetVbrInstanceRegionId(v string) *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecks {
	s.VbrInstanceRegionId = &v
	return s
}

func (s *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecks) SetCenId(v string) *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecks {
	s.CenId = &v
	return s
}

func (s *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecks) SetHealthyThreshold(v int32) *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecks {
	s.HealthyThreshold = &v
	return s
}

func (s *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecks) SetHealthCheckInterval(v int32) *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecks {
	s.HealthCheckInterval = &v
	return s
}

func (s *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecks) SetHealthCheckSourceIp(v string) *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecks {
	s.HealthCheckSourceIp = &v
	return s
}

type DescribeCenVbrHealthCheckResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCenVbrHealthCheckResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCenVbrHealthCheckResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenVbrHealthCheckResponse) GoString() string {
	return s.String()
}

func (s *DescribeCenVbrHealthCheckResponse) SetHeaders(v map[string]*string) *DescribeCenVbrHealthCheckResponse {
	s.Headers = v
	return s
}

func (s *DescribeCenVbrHealthCheckResponse) SetBody(v *DescribeCenVbrHealthCheckResponseBody) *DescribeCenVbrHealthCheckResponse {
	s.Body = v
	return s
}

type DescribeChildInstanceRegionsRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ProductType          *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	ChildInstanceOwnerId *int64  `json:"ChildInstanceOwnerId,omitempty" xml:"ChildInstanceOwnerId,omitempty"`
}

func (s DescribeChildInstanceRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeChildInstanceRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeChildInstanceRegionsRequest) SetOwnerAccount(v string) *DescribeChildInstanceRegionsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeChildInstanceRegionsRequest) SetOwnerId(v int64) *DescribeChildInstanceRegionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeChildInstanceRegionsRequest) SetResourceOwnerAccount(v string) *DescribeChildInstanceRegionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeChildInstanceRegionsRequest) SetResourceOwnerId(v int64) *DescribeChildInstanceRegionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeChildInstanceRegionsRequest) SetProductType(v string) *DescribeChildInstanceRegionsRequest {
	s.ProductType = &v
	return s
}

func (s *DescribeChildInstanceRegionsRequest) SetChildInstanceOwnerId(v int64) *DescribeChildInstanceRegionsRequest {
	s.ChildInstanceOwnerId = &v
	return s
}

type DescribeChildInstanceRegionsResponseBody struct {
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Regions   []*DescribeChildInstanceRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
}

func (s DescribeChildInstanceRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeChildInstanceRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChildInstanceRegionsResponseBody) SetRequestId(v string) *DescribeChildInstanceRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeChildInstanceRegionsResponseBody) SetRegions(v []*DescribeChildInstanceRegionsResponseBodyRegions) *DescribeChildInstanceRegionsResponseBody {
	s.Regions = v
	return s
}

type DescribeChildInstanceRegionsResponseBodyRegions struct {
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeChildInstanceRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeChildInstanceRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeChildInstanceRegionsResponseBodyRegions) SetLocalName(v string) *DescribeChildInstanceRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *DescribeChildInstanceRegionsResponseBodyRegions) SetRegionId(v string) *DescribeChildInstanceRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

type DescribeChildInstanceRegionsResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeChildInstanceRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeChildInstanceRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeChildInstanceRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeChildInstanceRegionsResponse) SetHeaders(v map[string]*string) *DescribeChildInstanceRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeChildInstanceRegionsResponse) SetBody(v *DescribeChildInstanceRegionsResponseBody) *DescribeChildInstanceRegionsResponse {
	s.Body = v
	return s
}

type DescribeFlowlogsRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	FlowLogName          *string `json:"FlowLogName,omitempty" xml:"FlowLogName,omitempty"`
	FlowLogId            *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ProjectName          *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	LogStoreName         *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeFlowlogsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowlogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowlogsRequest) SetOwnerAccount(v string) *DescribeFlowlogsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeFlowlogsRequest) SetOwnerId(v int64) *DescribeFlowlogsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeFlowlogsRequest) SetResourceOwnerAccount(v string) *DescribeFlowlogsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeFlowlogsRequest) SetResourceOwnerId(v int64) *DescribeFlowlogsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeFlowlogsRequest) SetClientToken(v string) *DescribeFlowlogsRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeFlowlogsRequest) SetRegionId(v string) *DescribeFlowlogsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFlowlogsRequest) SetFlowLogName(v string) *DescribeFlowlogsRequest {
	s.FlowLogName = &v
	return s
}

func (s *DescribeFlowlogsRequest) SetFlowLogId(v string) *DescribeFlowlogsRequest {
	s.FlowLogId = &v
	return s
}

func (s *DescribeFlowlogsRequest) SetDescription(v string) *DescribeFlowlogsRequest {
	s.Description = &v
	return s
}

func (s *DescribeFlowlogsRequest) SetCenId(v string) *DescribeFlowlogsRequest {
	s.CenId = &v
	return s
}

func (s *DescribeFlowlogsRequest) SetProjectName(v string) *DescribeFlowlogsRequest {
	s.ProjectName = &v
	return s
}

func (s *DescribeFlowlogsRequest) SetLogStoreName(v string) *DescribeFlowlogsRequest {
	s.LogStoreName = &v
	return s
}

func (s *DescribeFlowlogsRequest) SetStatus(v string) *DescribeFlowlogsRequest {
	s.Status = &v
	return s
}

func (s *DescribeFlowlogsRequest) SetPageNumber(v int32) *DescribeFlowlogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeFlowlogsRequest) SetPageSize(v int32) *DescribeFlowlogsRequest {
	s.PageSize = &v
	return s
}

type DescribeFlowlogsResponseBody struct {
	TotalCount *string                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	FlowLogs   []*DescribeFlowlogsResponseBodyFlowLogs `json:"FlowLogs,omitempty" xml:"FlowLogs,omitempty" type:"Repeated"`
	PageSize   *string                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *string                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Success    *string                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeFlowlogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowlogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowlogsResponseBody) SetTotalCount(v string) *DescribeFlowlogsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeFlowlogsResponseBody) SetFlowLogs(v []*DescribeFlowlogsResponseBodyFlowLogs) *DescribeFlowlogsResponseBody {
	s.FlowLogs = v
	return s
}

func (s *DescribeFlowlogsResponseBody) SetPageSize(v string) *DescribeFlowlogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeFlowlogsResponseBody) SetRequestId(v string) *DescribeFlowlogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlowlogsResponseBody) SetPageNumber(v string) *DescribeFlowlogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeFlowlogsResponseBody) SetSuccess(v string) *DescribeFlowlogsResponseBody {
	s.Success = &v
	return s
}

type DescribeFlowlogsResponseBodyFlowLogs struct {
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	FlowLogName  *string `json:"FlowLogName,omitempty" xml:"FlowLogName,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ProjectName  *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	CenId        *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	FlowLogId    *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
}

func (s DescribeFlowlogsResponseBodyFlowLogs) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowlogsResponseBodyFlowLogs) GoString() string {
	return s.String()
}

func (s *DescribeFlowlogsResponseBodyFlowLogs) SetStatus(v string) *DescribeFlowlogsResponseBodyFlowLogs {
	s.Status = &v
	return s
}

func (s *DescribeFlowlogsResponseBodyFlowLogs) SetCreationTime(v string) *DescribeFlowlogsResponseBodyFlowLogs {
	s.CreationTime = &v
	return s
}

func (s *DescribeFlowlogsResponseBodyFlowLogs) SetFlowLogName(v string) *DescribeFlowlogsResponseBodyFlowLogs {
	s.FlowLogName = &v
	return s
}

func (s *DescribeFlowlogsResponseBodyFlowLogs) SetDescription(v string) *DescribeFlowlogsResponseBodyFlowLogs {
	s.Description = &v
	return s
}

func (s *DescribeFlowlogsResponseBodyFlowLogs) SetProjectName(v string) *DescribeFlowlogsResponseBodyFlowLogs {
	s.ProjectName = &v
	return s
}

func (s *DescribeFlowlogsResponseBodyFlowLogs) SetCenId(v string) *DescribeFlowlogsResponseBodyFlowLogs {
	s.CenId = &v
	return s
}

func (s *DescribeFlowlogsResponseBodyFlowLogs) SetLogStoreName(v string) *DescribeFlowlogsResponseBodyFlowLogs {
	s.LogStoreName = &v
	return s
}

func (s *DescribeFlowlogsResponseBodyFlowLogs) SetRegionId(v string) *DescribeFlowlogsResponseBodyFlowLogs {
	s.RegionId = &v
	return s
}

func (s *DescribeFlowlogsResponseBodyFlowLogs) SetFlowLogId(v string) *DescribeFlowlogsResponseBodyFlowLogs {
	s.FlowLogId = &v
	return s
}

type DescribeFlowlogsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFlowlogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFlowlogsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowlogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlowlogsResponse) SetHeaders(v map[string]*string) *DescribeFlowlogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlowlogsResponse) SetBody(v *DescribeFlowlogsResponseBody) *DescribeFlowlogsResponse {
	s.Body = v
	return s
}

type DescribeGeographicRegionMembershipRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	GeographicRegionId   *string `json:"GeographicRegionId,omitempty" xml:"GeographicRegionId,omitempty"`
}

func (s DescribeGeographicRegionMembershipRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGeographicRegionMembershipRequest) GoString() string {
	return s.String()
}

func (s *DescribeGeographicRegionMembershipRequest) SetOwnerAccount(v string) *DescribeGeographicRegionMembershipRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeGeographicRegionMembershipRequest) SetOwnerId(v int64) *DescribeGeographicRegionMembershipRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeGeographicRegionMembershipRequest) SetResourceOwnerAccount(v string) *DescribeGeographicRegionMembershipRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeGeographicRegionMembershipRequest) SetResourceOwnerId(v int64) *DescribeGeographicRegionMembershipRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeGeographicRegionMembershipRequest) SetPageNumber(v int32) *DescribeGeographicRegionMembershipRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeGeographicRegionMembershipRequest) SetPageSize(v int32) *DescribeGeographicRegionMembershipRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGeographicRegionMembershipRequest) SetGeographicRegionId(v string) *DescribeGeographicRegionMembershipRequest {
	s.GeographicRegionId = &v
	return s
}

type DescribeGeographicRegionMembershipResponseBody struct {
	TotalCount *int32                                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int32                                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RegionIds  []*DescribeGeographicRegionMembershipResponseBodyRegionIds `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Repeated"`
	PageNumber *int32                                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeGeographicRegionMembershipResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGeographicRegionMembershipResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGeographicRegionMembershipResponseBody) SetTotalCount(v int32) *DescribeGeographicRegionMembershipResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeGeographicRegionMembershipResponseBody) SetPageSize(v int32) *DescribeGeographicRegionMembershipResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGeographicRegionMembershipResponseBody) SetRequestId(v string) *DescribeGeographicRegionMembershipResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGeographicRegionMembershipResponseBody) SetRegionIds(v []*DescribeGeographicRegionMembershipResponseBodyRegionIds) *DescribeGeographicRegionMembershipResponseBody {
	s.RegionIds = v
	return s
}

func (s *DescribeGeographicRegionMembershipResponseBody) SetPageNumber(v int32) *DescribeGeographicRegionMembershipResponseBody {
	s.PageNumber = &v
	return s
}

type DescribeGeographicRegionMembershipResponseBodyRegionIds struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeGeographicRegionMembershipResponseBodyRegionIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeGeographicRegionMembershipResponseBodyRegionIds) GoString() string {
	return s.String()
}

func (s *DescribeGeographicRegionMembershipResponseBodyRegionIds) SetRegionId(v string) *DescribeGeographicRegionMembershipResponseBodyRegionIds {
	s.RegionId = &v
	return s
}

type DescribeGeographicRegionMembershipResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGeographicRegionMembershipResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGeographicRegionMembershipResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGeographicRegionMembershipResponse) GoString() string {
	return s.String()
}

func (s *DescribeGeographicRegionMembershipResponse) SetHeaders(v map[string]*string) *DescribeGeographicRegionMembershipResponse {
	s.Headers = v
	return s
}

func (s *DescribeGeographicRegionMembershipResponse) SetBody(v *DescribeGeographicRegionMembershipResponseBody) *DescribeGeographicRegionMembershipResponse {
	s.Body = v
	return s
}

type DescribeGrantRulesToCenRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ProductType          *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s DescribeGrantRulesToCenRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGrantRulesToCenRequest) GoString() string {
	return s.String()
}

func (s *DescribeGrantRulesToCenRequest) SetOwnerAccount(v string) *DescribeGrantRulesToCenRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeGrantRulesToCenRequest) SetOwnerId(v int64) *DescribeGrantRulesToCenRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeGrantRulesToCenRequest) SetResourceOwnerAccount(v string) *DescribeGrantRulesToCenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeGrantRulesToCenRequest) SetResourceOwnerId(v int64) *DescribeGrantRulesToCenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeGrantRulesToCenRequest) SetRegionId(v string) *DescribeGrantRulesToCenRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeGrantRulesToCenRequest) SetCenId(v string) *DescribeGrantRulesToCenRequest {
	s.CenId = &v
	return s
}

func (s *DescribeGrantRulesToCenRequest) SetProductType(v string) *DescribeGrantRulesToCenRequest {
	s.ProductType = &v
	return s
}

type DescribeGrantRulesToCenResponseBody struct {
	RequestId  *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	GrantRules []*DescribeGrantRulesToCenResponseBodyGrantRules `json:"GrantRules,omitempty" xml:"GrantRules,omitempty" type:"Repeated"`
}

func (s DescribeGrantRulesToCenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGrantRulesToCenResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGrantRulesToCenResponseBody) SetRequestId(v string) *DescribeGrantRulesToCenResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGrantRulesToCenResponseBody) SetGrantRules(v []*DescribeGrantRulesToCenResponseBodyGrantRules) *DescribeGrantRulesToCenResponseBody {
	s.GrantRules = v
	return s
}

type DescribeGrantRulesToCenResponseBodyGrantRules struct {
	ChildInstanceType     *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	ChildInstanceRegionId *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	ChildInstanceOwnerId  *int64  `json:"ChildInstanceOwnerId,omitempty" xml:"ChildInstanceOwnerId,omitempty"`
	ChildInstanceId       *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	CenId                 *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
}

func (s DescribeGrantRulesToCenResponseBodyGrantRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeGrantRulesToCenResponseBodyGrantRules) GoString() string {
	return s.String()
}

func (s *DescribeGrantRulesToCenResponseBodyGrantRules) SetChildInstanceType(v string) *DescribeGrantRulesToCenResponseBodyGrantRules {
	s.ChildInstanceType = &v
	return s
}

func (s *DescribeGrantRulesToCenResponseBodyGrantRules) SetChildInstanceRegionId(v string) *DescribeGrantRulesToCenResponseBodyGrantRules {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *DescribeGrantRulesToCenResponseBodyGrantRules) SetChildInstanceOwnerId(v int64) *DescribeGrantRulesToCenResponseBodyGrantRules {
	s.ChildInstanceOwnerId = &v
	return s
}

func (s *DescribeGrantRulesToCenResponseBodyGrantRules) SetChildInstanceId(v string) *DescribeGrantRulesToCenResponseBodyGrantRules {
	s.ChildInstanceId = &v
	return s
}

func (s *DescribeGrantRulesToCenResponseBodyGrantRules) SetCenId(v string) *DescribeGrantRulesToCenResponseBodyGrantRules {
	s.CenId = &v
	return s
}

type DescribeGrantRulesToCenResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGrantRulesToCenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGrantRulesToCenResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGrantRulesToCenResponse) GoString() string {
	return s.String()
}

func (s *DescribeGrantRulesToCenResponse) SetHeaders(v map[string]*string) *DescribeGrantRulesToCenResponse {
	s.Headers = v
	return s
}

func (s *DescribeGrantRulesToCenResponse) SetBody(v *DescribeGrantRulesToCenResponseBody) *DescribeGrantRulesToCenResponse {
	s.Body = v
	return s
}

type DescribePublishedRouteEntriesRequest struct {
	ResourceOwnerAccount      *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId           *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PageNumber                *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize                  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CenId                     *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ChildInstanceId           *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	ChildInstanceType         *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	ChildInstanceRegionId     *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	ChildInstanceRouteTableId *string `json:"ChildInstanceRouteTableId,omitempty" xml:"ChildInstanceRouteTableId,omitempty"`
	DestinationCidrBlock      *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
}

func (s DescribePublishedRouteEntriesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePublishedRouteEntriesRequest) GoString() string {
	return s.String()
}

func (s *DescribePublishedRouteEntriesRequest) SetResourceOwnerAccount(v string) *DescribePublishedRouteEntriesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePublishedRouteEntriesRequest) SetResourceOwnerId(v int64) *DescribePublishedRouteEntriesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribePublishedRouteEntriesRequest) SetPageNumber(v int32) *DescribePublishedRouteEntriesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePublishedRouteEntriesRequest) SetPageSize(v int32) *DescribePublishedRouteEntriesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePublishedRouteEntriesRequest) SetCenId(v string) *DescribePublishedRouteEntriesRequest {
	s.CenId = &v
	return s
}

func (s *DescribePublishedRouteEntriesRequest) SetChildInstanceId(v string) *DescribePublishedRouteEntriesRequest {
	s.ChildInstanceId = &v
	return s
}

func (s *DescribePublishedRouteEntriesRequest) SetChildInstanceType(v string) *DescribePublishedRouteEntriesRequest {
	s.ChildInstanceType = &v
	return s
}

func (s *DescribePublishedRouteEntriesRequest) SetChildInstanceRegionId(v string) *DescribePublishedRouteEntriesRequest {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *DescribePublishedRouteEntriesRequest) SetChildInstanceRouteTableId(v string) *DescribePublishedRouteEntriesRequest {
	s.ChildInstanceRouteTableId = &v
	return s
}

func (s *DescribePublishedRouteEntriesRequest) SetDestinationCidrBlock(v string) *DescribePublishedRouteEntriesRequest {
	s.DestinationCidrBlock = &v
	return s
}

type DescribePublishedRouteEntriesResponseBody struct {
	TotalCount            *int32                                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize              *int32                                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId             *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber            *int32                                                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PublishedRouteEntries []*DescribePublishedRouteEntriesResponseBodyPublishedRouteEntries `json:"PublishedRouteEntries,omitempty" xml:"PublishedRouteEntries,omitempty" type:"Repeated"`
}

func (s DescribePublishedRouteEntriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePublishedRouteEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePublishedRouteEntriesResponseBody) SetTotalCount(v int32) *DescribePublishedRouteEntriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBody) SetPageSize(v int32) *DescribePublishedRouteEntriesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBody) SetRequestId(v string) *DescribePublishedRouteEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBody) SetPageNumber(v int32) *DescribePublishedRouteEntriesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBody) SetPublishedRouteEntries(v []*DescribePublishedRouteEntriesResponseBodyPublishedRouteEntries) *DescribePublishedRouteEntriesResponseBody {
	s.PublishedRouteEntries = v
	return s
}

type DescribePublishedRouteEntriesResponseBodyPublishedRouteEntries struct {
	NextHopId                 *string                                                                    `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	Conflicts                 []*DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesConflicts `json:"Conflicts,omitempty" xml:"Conflicts,omitempty" type:"Repeated"`
	PublishStatus             *string                                                                    `json:"PublishStatus,omitempty" xml:"PublishStatus,omitempty"`
	ChildInstanceRouteTableId *string                                                                    `json:"ChildInstanceRouteTableId,omitempty" xml:"ChildInstanceRouteTableId,omitempty"`
	NextHopType               *string                                                                    `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	OperationalMode           *bool                                                                      `json:"OperationalMode,omitempty" xml:"OperationalMode,omitempty"`
	DestinationCidrBlock      *string                                                                    `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	RouteType                 *string                                                                    `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
}

func (s DescribePublishedRouteEntriesResponseBodyPublishedRouteEntries) String() string {
	return tea.Prettify(s)
}

func (s DescribePublishedRouteEntriesResponseBodyPublishedRouteEntries) GoString() string {
	return s.String()
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntries) SetNextHopId(v string) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntries {
	s.NextHopId = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntries) SetConflicts(v []*DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesConflicts) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntries {
	s.Conflicts = v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntries) SetPublishStatus(v string) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntries {
	s.PublishStatus = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntries) SetChildInstanceRouteTableId(v string) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntries {
	s.ChildInstanceRouteTableId = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntries) SetNextHopType(v string) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntries {
	s.NextHopType = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntries) SetOperationalMode(v bool) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntries {
	s.OperationalMode = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntries) SetDestinationCidrBlock(v string) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntries {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntries) SetRouteType(v string) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntries {
	s.RouteType = &v
	return s
}

type DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesConflicts struct {
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType         *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesConflicts) String() string {
	return tea.Prettify(s)
}

func (s DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesConflicts) GoString() string {
	return s.String()
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesConflicts) SetStatus(v string) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesConflicts {
	s.Status = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesConflicts) SetDestinationCidrBlock(v string) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesConflicts {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesConflicts) SetInstanceId(v string) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesConflicts {
	s.InstanceId = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesConflicts) SetInstanceType(v string) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesConflicts {
	s.InstanceType = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesConflicts) SetRegionId(v string) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesConflicts {
	s.RegionId = &v
	return s
}

type DescribePublishedRouteEntriesResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePublishedRouteEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePublishedRouteEntriesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePublishedRouteEntriesResponse) GoString() string {
	return s.String()
}

func (s *DescribePublishedRouteEntriesResponse) SetHeaders(v map[string]*string) *DescribePublishedRouteEntriesResponse {
	s.Headers = v
	return s
}

func (s *DescribePublishedRouteEntriesResponse) SetBody(v *DescribePublishedRouteEntriesResponseBody) *DescribePublishedRouteEntriesResponse {
	s.Body = v
	return s
}

type DescribeRouteConflictRequest struct {
	OwnerAccount              *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount      *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId           *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PageNumber                *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize                  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ChildInstanceId           *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	ChildInstanceType         *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	ChildInstanceRegionId     *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	ChildInstanceRouteTableId *string `json:"ChildInstanceRouteTableId,omitempty" xml:"ChildInstanceRouteTableId,omitempty"`
	DestinationCidrBlock      *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
}

func (s DescribeRouteConflictRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRouteConflictRequest) GoString() string {
	return s.String()
}

func (s *DescribeRouteConflictRequest) SetOwnerAccount(v string) *DescribeRouteConflictRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRouteConflictRequest) SetOwnerId(v int64) *DescribeRouteConflictRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRouteConflictRequest) SetResourceOwnerAccount(v string) *DescribeRouteConflictRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRouteConflictRequest) SetResourceOwnerId(v int64) *DescribeRouteConflictRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRouteConflictRequest) SetPageNumber(v int32) *DescribeRouteConflictRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRouteConflictRequest) SetPageSize(v int32) *DescribeRouteConflictRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRouteConflictRequest) SetChildInstanceId(v string) *DescribeRouteConflictRequest {
	s.ChildInstanceId = &v
	return s
}

func (s *DescribeRouteConflictRequest) SetChildInstanceType(v string) *DescribeRouteConflictRequest {
	s.ChildInstanceType = &v
	return s
}

func (s *DescribeRouteConflictRequest) SetChildInstanceRegionId(v string) *DescribeRouteConflictRequest {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *DescribeRouteConflictRequest) SetChildInstanceRouteTableId(v string) *DescribeRouteConflictRequest {
	s.ChildInstanceRouteTableId = &v
	return s
}

func (s *DescribeRouteConflictRequest) SetDestinationCidrBlock(v string) *DescribeRouteConflictRequest {
	s.DestinationCidrBlock = &v
	return s
}

type DescribeRouteConflictResponseBody struct {
	TotalCount     *int32                                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize       *int32                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber     *int32                                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	RouteConflicts []*DescribeRouteConflictResponseBodyRouteConflicts `json:"RouteConflicts,omitempty" xml:"RouteConflicts,omitempty" type:"Repeated"`
}

func (s DescribeRouteConflictResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRouteConflictResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRouteConflictResponseBody) SetTotalCount(v int32) *DescribeRouteConflictResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRouteConflictResponseBody) SetPageSize(v int32) *DescribeRouteConflictResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRouteConflictResponseBody) SetRequestId(v string) *DescribeRouteConflictResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRouteConflictResponseBody) SetPageNumber(v int32) *DescribeRouteConflictResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeRouteConflictResponseBody) SetRouteConflicts(v []*DescribeRouteConflictResponseBodyRouteConflicts) *DescribeRouteConflictResponseBody {
	s.RouteConflicts = v
	return s
}

type DescribeRouteConflictResponseBodyRouteConflicts struct {
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType         *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRouteConflictResponseBodyRouteConflicts) String() string {
	return tea.Prettify(s)
}

func (s DescribeRouteConflictResponseBodyRouteConflicts) GoString() string {
	return s.String()
}

func (s *DescribeRouteConflictResponseBodyRouteConflicts) SetStatus(v string) *DescribeRouteConflictResponseBodyRouteConflicts {
	s.Status = &v
	return s
}

func (s *DescribeRouteConflictResponseBodyRouteConflicts) SetDestinationCidrBlock(v string) *DescribeRouteConflictResponseBodyRouteConflicts {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DescribeRouteConflictResponseBodyRouteConflicts) SetInstanceId(v string) *DescribeRouteConflictResponseBodyRouteConflicts {
	s.InstanceId = &v
	return s
}

func (s *DescribeRouteConflictResponseBodyRouteConflicts) SetInstanceType(v string) *DescribeRouteConflictResponseBodyRouteConflicts {
	s.InstanceType = &v
	return s
}

func (s *DescribeRouteConflictResponseBodyRouteConflicts) SetRegionId(v string) *DescribeRouteConflictResponseBodyRouteConflicts {
	s.RegionId = &v
	return s
}

type DescribeRouteConflictResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRouteConflictResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRouteConflictResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRouteConflictResponse) GoString() string {
	return s.String()
}

func (s *DescribeRouteConflictResponse) SetHeaders(v map[string]*string) *DescribeRouteConflictResponse {
	s.Headers = v
	return s
}

func (s *DescribeRouteConflictResponse) SetBody(v *DescribeRouteConflictResponseBody) *DescribeRouteConflictResponse {
	s.Body = v
	return s
}

type DescribeRouteServicesInCenRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	Host                 *string `json:"Host,omitempty" xml:"Host,omitempty"`
	HostRegionId         *string `json:"HostRegionId,omitempty" xml:"HostRegionId,omitempty"`
	AccessRegionId       *string `json:"AccessRegionId,omitempty" xml:"AccessRegionId,omitempty"`
	HostVpcId            *string `json:"HostVpcId,omitempty" xml:"HostVpcId,omitempty"`
}

func (s DescribeRouteServicesInCenRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRouteServicesInCenRequest) GoString() string {
	return s.String()
}

func (s *DescribeRouteServicesInCenRequest) SetOwnerAccount(v string) *DescribeRouteServicesInCenRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRouteServicesInCenRequest) SetOwnerId(v int64) *DescribeRouteServicesInCenRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRouteServicesInCenRequest) SetResourceOwnerAccount(v string) *DescribeRouteServicesInCenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRouteServicesInCenRequest) SetResourceOwnerId(v int64) *DescribeRouteServicesInCenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRouteServicesInCenRequest) SetPageNumber(v int32) *DescribeRouteServicesInCenRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRouteServicesInCenRequest) SetPageSize(v int32) *DescribeRouteServicesInCenRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRouteServicesInCenRequest) SetCenId(v string) *DescribeRouteServicesInCenRequest {
	s.CenId = &v
	return s
}

func (s *DescribeRouteServicesInCenRequest) SetHost(v string) *DescribeRouteServicesInCenRequest {
	s.Host = &v
	return s
}

func (s *DescribeRouteServicesInCenRequest) SetHostRegionId(v string) *DescribeRouteServicesInCenRequest {
	s.HostRegionId = &v
	return s
}

func (s *DescribeRouteServicesInCenRequest) SetAccessRegionId(v string) *DescribeRouteServicesInCenRequest {
	s.AccessRegionId = &v
	return s
}

func (s *DescribeRouteServicesInCenRequest) SetHostVpcId(v string) *DescribeRouteServicesInCenRequest {
	s.HostVpcId = &v
	return s
}

type DescribeRouteServicesInCenResponseBody struct {
	TotalCount          *int32                                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize            *int32                                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId           *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber          *int32                                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	RouteServiceEntries []*DescribeRouteServicesInCenResponseBodyRouteServiceEntries `json:"RouteServiceEntries,omitempty" xml:"RouteServiceEntries,omitempty" type:"Repeated"`
}

func (s DescribeRouteServicesInCenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRouteServicesInCenResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRouteServicesInCenResponseBody) SetTotalCount(v int32) *DescribeRouteServicesInCenResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRouteServicesInCenResponseBody) SetPageSize(v int32) *DescribeRouteServicesInCenResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRouteServicesInCenResponseBody) SetRequestId(v string) *DescribeRouteServicesInCenResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRouteServicesInCenResponseBody) SetPageNumber(v int32) *DescribeRouteServicesInCenResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeRouteServicesInCenResponseBody) SetRouteServiceEntries(v []*DescribeRouteServicesInCenResponseBodyRouteServiceEntries) *DescribeRouteServicesInCenResponseBody {
	s.RouteServiceEntries = v
	return s
}

type DescribeRouteServicesInCenResponseBodyRouteServiceEntries struct {
	Status         *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	Host           *string   `json:"Host,omitempty" xml:"Host,omitempty"`
	Description    *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	HostVpcId      *string   `json:"HostVpcId,omitempty" xml:"HostVpcId,omitempty"`
	Cidrs          []*string `json:"Cidrs,omitempty" xml:"Cidrs,omitempty" type:"Repeated"`
	CenId          *string   `json:"CenId,omitempty" xml:"CenId,omitempty"`
	AccessRegionId *string   `json:"AccessRegionId,omitempty" xml:"AccessRegionId,omitempty"`
	HostRegionId   *string   `json:"HostRegionId,omitempty" xml:"HostRegionId,omitempty"`
	UpdateInterval *string   `json:"UpdateInterval,omitempty" xml:"UpdateInterval,omitempty"`
}

func (s DescribeRouteServicesInCenResponseBodyRouteServiceEntries) String() string {
	return tea.Prettify(s)
}

func (s DescribeRouteServicesInCenResponseBodyRouteServiceEntries) GoString() string {
	return s.String()
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntries) SetStatus(v string) *DescribeRouteServicesInCenResponseBodyRouteServiceEntries {
	s.Status = &v
	return s
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntries) SetHost(v string) *DescribeRouteServicesInCenResponseBodyRouteServiceEntries {
	s.Host = &v
	return s
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntries) SetDescription(v string) *DescribeRouteServicesInCenResponseBodyRouteServiceEntries {
	s.Description = &v
	return s
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntries) SetHostVpcId(v string) *DescribeRouteServicesInCenResponseBodyRouteServiceEntries {
	s.HostVpcId = &v
	return s
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntries) SetCidrs(v []*string) *DescribeRouteServicesInCenResponseBodyRouteServiceEntries {
	s.Cidrs = v
	return s
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntries) SetCenId(v string) *DescribeRouteServicesInCenResponseBodyRouteServiceEntries {
	s.CenId = &v
	return s
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntries) SetAccessRegionId(v string) *DescribeRouteServicesInCenResponseBodyRouteServiceEntries {
	s.AccessRegionId = &v
	return s
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntries) SetHostRegionId(v string) *DescribeRouteServicesInCenResponseBodyRouteServiceEntries {
	s.HostRegionId = &v
	return s
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntries) SetUpdateInterval(v string) *DescribeRouteServicesInCenResponseBodyRouteServiceEntries {
	s.UpdateInterval = &v
	return s
}

type DescribeRouteServicesInCenResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRouteServicesInCenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRouteServicesInCenResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRouteServicesInCenResponse) GoString() string {
	return s.String()
}

func (s *DescribeRouteServicesInCenResponse) SetHeaders(v map[string]*string) *DescribeRouteServicesInCenResponse {
	s.Headers = v
	return s
}

func (s *DescribeRouteServicesInCenResponse) SetBody(v *DescribeRouteServicesInCenResponseBody) *DescribeRouteServicesInCenResponse {
	s.Body = v
	return s
}

type DetachCenChildInstanceRequest struct {
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	CenId                 *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ChildInstanceId       *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	ChildInstanceType     *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	ChildInstanceRegionId *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	ChildInstanceOwnerId  *int64  `json:"ChildInstanceOwnerId,omitempty" xml:"ChildInstanceOwnerId,omitempty"`
	CenOwnerId            *int64  `json:"CenOwnerId,omitempty" xml:"CenOwnerId,omitempty"`
}

func (s DetachCenChildInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachCenChildInstanceRequest) GoString() string {
	return s.String()
}

func (s *DetachCenChildInstanceRequest) SetOwnerAccount(v string) *DetachCenChildInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DetachCenChildInstanceRequest) SetOwnerId(v int64) *DetachCenChildInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DetachCenChildInstanceRequest) SetResourceOwnerAccount(v string) *DetachCenChildInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DetachCenChildInstanceRequest) SetResourceOwnerId(v int64) *DetachCenChildInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DetachCenChildInstanceRequest) SetCenId(v string) *DetachCenChildInstanceRequest {
	s.CenId = &v
	return s
}

func (s *DetachCenChildInstanceRequest) SetChildInstanceId(v string) *DetachCenChildInstanceRequest {
	s.ChildInstanceId = &v
	return s
}

func (s *DetachCenChildInstanceRequest) SetChildInstanceType(v string) *DetachCenChildInstanceRequest {
	s.ChildInstanceType = &v
	return s
}

func (s *DetachCenChildInstanceRequest) SetChildInstanceRegionId(v string) *DetachCenChildInstanceRequest {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *DetachCenChildInstanceRequest) SetChildInstanceOwnerId(v int64) *DetachCenChildInstanceRequest {
	s.ChildInstanceOwnerId = &v
	return s
}

func (s *DetachCenChildInstanceRequest) SetCenOwnerId(v int64) *DetachCenChildInstanceRequest {
	s.CenOwnerId = &v
	return s
}

type DetachCenChildInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachCenChildInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachCenChildInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DetachCenChildInstanceResponseBody) SetRequestId(v string) *DetachCenChildInstanceResponseBody {
	s.RequestId = &v
	return s
}

type DetachCenChildInstanceResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetachCenChildInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetachCenChildInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachCenChildInstanceResponse) GoString() string {
	return s.String()
}

func (s *DetachCenChildInstanceResponse) SetHeaders(v map[string]*string) *DetachCenChildInstanceResponse {
	s.Headers = v
	return s
}

func (s *DetachCenChildInstanceResponse) SetBody(v *DetachCenChildInstanceResponseBody) *DetachCenChildInstanceResponse {
	s.Body = v
	return s
}

type DisableCenVbrHealthCheckRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	VbrInstanceRegionId  *string `json:"VbrInstanceRegionId,omitempty" xml:"VbrInstanceRegionId,omitempty"`
	VbrInstanceId        *string `json:"VbrInstanceId,omitempty" xml:"VbrInstanceId,omitempty"`
	VbrInstanceOwnerId   *int64  `json:"VbrInstanceOwnerId,omitempty" xml:"VbrInstanceOwnerId,omitempty"`
}

func (s DisableCenVbrHealthCheckRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableCenVbrHealthCheckRequest) GoString() string {
	return s.String()
}

func (s *DisableCenVbrHealthCheckRequest) SetOwnerAccount(v string) *DisableCenVbrHealthCheckRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DisableCenVbrHealthCheckRequest) SetOwnerId(v int64) *DisableCenVbrHealthCheckRequest {
	s.OwnerId = &v
	return s
}

func (s *DisableCenVbrHealthCheckRequest) SetResourceOwnerAccount(v string) *DisableCenVbrHealthCheckRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DisableCenVbrHealthCheckRequest) SetResourceOwnerId(v int64) *DisableCenVbrHealthCheckRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DisableCenVbrHealthCheckRequest) SetCenId(v string) *DisableCenVbrHealthCheckRequest {
	s.CenId = &v
	return s
}

func (s *DisableCenVbrHealthCheckRequest) SetVbrInstanceRegionId(v string) *DisableCenVbrHealthCheckRequest {
	s.VbrInstanceRegionId = &v
	return s
}

func (s *DisableCenVbrHealthCheckRequest) SetVbrInstanceId(v string) *DisableCenVbrHealthCheckRequest {
	s.VbrInstanceId = &v
	return s
}

func (s *DisableCenVbrHealthCheckRequest) SetVbrInstanceOwnerId(v int64) *DisableCenVbrHealthCheckRequest {
	s.VbrInstanceOwnerId = &v
	return s
}

type DisableCenVbrHealthCheckResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableCenVbrHealthCheckResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableCenVbrHealthCheckResponseBody) GoString() string {
	return s.String()
}

func (s *DisableCenVbrHealthCheckResponseBody) SetRequestId(v string) *DisableCenVbrHealthCheckResponseBody {
	s.RequestId = &v
	return s
}

type DisableCenVbrHealthCheckResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableCenVbrHealthCheckResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableCenVbrHealthCheckResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableCenVbrHealthCheckResponse) GoString() string {
	return s.String()
}

func (s *DisableCenVbrHealthCheckResponse) SetHeaders(v map[string]*string) *DisableCenVbrHealthCheckResponse {
	s.Headers = v
	return s
}

func (s *DisableCenVbrHealthCheckResponse) SetBody(v *DisableCenVbrHealthCheckResponseBody) *DisableCenVbrHealthCheckResponse {
	s.Body = v
	return s
}

type EnableCenVbrHealthCheckRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	VbrInstanceRegionId  *string `json:"VbrInstanceRegionId,omitempty" xml:"VbrInstanceRegionId,omitempty"`
	VbrInstanceId        *string `json:"VbrInstanceId,omitempty" xml:"VbrInstanceId,omitempty"`
	HealthCheckSourceIp  *string `json:"HealthCheckSourceIp,omitempty" xml:"HealthCheckSourceIp,omitempty"`
	HealthCheckTargetIp  *string `json:"HealthCheckTargetIp,omitempty" xml:"HealthCheckTargetIp,omitempty"`
	VbrInstanceOwnerId   *int64  `json:"VbrInstanceOwnerId,omitempty" xml:"VbrInstanceOwnerId,omitempty"`
	HealthCheckInterval  *int32  `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	HealthyThreshold     *int32  `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
}

func (s EnableCenVbrHealthCheckRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableCenVbrHealthCheckRequest) GoString() string {
	return s.String()
}

func (s *EnableCenVbrHealthCheckRequest) SetOwnerAccount(v string) *EnableCenVbrHealthCheckRequest {
	s.OwnerAccount = &v
	return s
}

func (s *EnableCenVbrHealthCheckRequest) SetOwnerId(v int64) *EnableCenVbrHealthCheckRequest {
	s.OwnerId = &v
	return s
}

func (s *EnableCenVbrHealthCheckRequest) SetResourceOwnerAccount(v string) *EnableCenVbrHealthCheckRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *EnableCenVbrHealthCheckRequest) SetResourceOwnerId(v int64) *EnableCenVbrHealthCheckRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *EnableCenVbrHealthCheckRequest) SetCenId(v string) *EnableCenVbrHealthCheckRequest {
	s.CenId = &v
	return s
}

func (s *EnableCenVbrHealthCheckRequest) SetVbrInstanceRegionId(v string) *EnableCenVbrHealthCheckRequest {
	s.VbrInstanceRegionId = &v
	return s
}

func (s *EnableCenVbrHealthCheckRequest) SetVbrInstanceId(v string) *EnableCenVbrHealthCheckRequest {
	s.VbrInstanceId = &v
	return s
}

func (s *EnableCenVbrHealthCheckRequest) SetHealthCheckSourceIp(v string) *EnableCenVbrHealthCheckRequest {
	s.HealthCheckSourceIp = &v
	return s
}

func (s *EnableCenVbrHealthCheckRequest) SetHealthCheckTargetIp(v string) *EnableCenVbrHealthCheckRequest {
	s.HealthCheckTargetIp = &v
	return s
}

func (s *EnableCenVbrHealthCheckRequest) SetVbrInstanceOwnerId(v int64) *EnableCenVbrHealthCheckRequest {
	s.VbrInstanceOwnerId = &v
	return s
}

func (s *EnableCenVbrHealthCheckRequest) SetHealthCheckInterval(v int32) *EnableCenVbrHealthCheckRequest {
	s.HealthCheckInterval = &v
	return s
}

func (s *EnableCenVbrHealthCheckRequest) SetHealthyThreshold(v int32) *EnableCenVbrHealthCheckRequest {
	s.HealthyThreshold = &v
	return s
}

type EnableCenVbrHealthCheckResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableCenVbrHealthCheckResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableCenVbrHealthCheckResponseBody) GoString() string {
	return s.String()
}

func (s *EnableCenVbrHealthCheckResponseBody) SetRequestId(v string) *EnableCenVbrHealthCheckResponseBody {
	s.RequestId = &v
	return s
}

type EnableCenVbrHealthCheckResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableCenVbrHealthCheckResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableCenVbrHealthCheckResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableCenVbrHealthCheckResponse) GoString() string {
	return s.String()
}

func (s *EnableCenVbrHealthCheckResponse) SetHeaders(v map[string]*string) *EnableCenVbrHealthCheckResponse {
	s.Headers = v
	return s
}

func (s *EnableCenVbrHealthCheckResponse) SetBody(v *EnableCenVbrHealthCheckResponseBody) *EnableCenVbrHealthCheckResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	OwnerAccount         *string                       `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64                        `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string                       `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                        `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceType         *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	NextToken            *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PageSize             *int32                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceId           []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	Tag                  []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetOwnerAccount(v string) *ListTagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTagResourcesRequest) SetOwnerId(v int64) *ListTagResourcesRequest {
	s.OwnerId = &v
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

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetPageSize(v int32) *ListTagResourcesRequest {
	s.PageSize = &v
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
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagValue(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagValue = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceId(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagKey(v string) *ListTagResourcesResponseBodyTagResources {
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

type ModifyCenAttributeRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ProtectionLevel      *string `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty"`
	Ipv6Level            *string `json:"Ipv6Level,omitempty" xml:"Ipv6Level,omitempty"`
}

func (s ModifyCenAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyCenAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyCenAttributeRequest) SetOwnerAccount(v string) *ModifyCenAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyCenAttributeRequest) SetOwnerId(v int64) *ModifyCenAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyCenAttributeRequest) SetResourceOwnerAccount(v string) *ModifyCenAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyCenAttributeRequest) SetResourceOwnerId(v int64) *ModifyCenAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyCenAttributeRequest) SetCenId(v string) *ModifyCenAttributeRequest {
	s.CenId = &v
	return s
}

func (s *ModifyCenAttributeRequest) SetName(v string) *ModifyCenAttributeRequest {
	s.Name = &v
	return s
}

func (s *ModifyCenAttributeRequest) SetDescription(v string) *ModifyCenAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyCenAttributeRequest) SetProtectionLevel(v string) *ModifyCenAttributeRequest {
	s.ProtectionLevel = &v
	return s
}

func (s *ModifyCenAttributeRequest) SetIpv6Level(v string) *ModifyCenAttributeRequest {
	s.Ipv6Level = &v
	return s
}

type ModifyCenAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCenAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyCenAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCenAttributeResponseBody) SetRequestId(v string) *ModifyCenAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyCenAttributeResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyCenAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyCenAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyCenAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyCenAttributeResponse) SetHeaders(v map[string]*string) *ModifyCenAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyCenAttributeResponse) SetBody(v *ModifyCenAttributeResponseBody) *ModifyCenAttributeResponse {
	s.Body = v
	return s
}

type ModifyCenBandwidthPackageAttributeRequest struct {
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Name                  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description           *string `json:"Description,omitempty" xml:"Description,omitempty"`
	CenBandwidthPackageId *string `json:"CenBandwidthPackageId,omitempty" xml:"CenBandwidthPackageId,omitempty"`
}

func (s ModifyCenBandwidthPackageAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyCenBandwidthPackageAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyCenBandwidthPackageAttributeRequest) SetOwnerAccount(v string) *ModifyCenBandwidthPackageAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyCenBandwidthPackageAttributeRequest) SetOwnerId(v int64) *ModifyCenBandwidthPackageAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyCenBandwidthPackageAttributeRequest) SetResourceOwnerAccount(v string) *ModifyCenBandwidthPackageAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyCenBandwidthPackageAttributeRequest) SetResourceOwnerId(v int64) *ModifyCenBandwidthPackageAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyCenBandwidthPackageAttributeRequest) SetName(v string) *ModifyCenBandwidthPackageAttributeRequest {
	s.Name = &v
	return s
}

func (s *ModifyCenBandwidthPackageAttributeRequest) SetDescription(v string) *ModifyCenBandwidthPackageAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyCenBandwidthPackageAttributeRequest) SetCenBandwidthPackageId(v string) *ModifyCenBandwidthPackageAttributeRequest {
	s.CenBandwidthPackageId = &v
	return s
}

type ModifyCenBandwidthPackageAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCenBandwidthPackageAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyCenBandwidthPackageAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCenBandwidthPackageAttributeResponseBody) SetRequestId(v string) *ModifyCenBandwidthPackageAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyCenBandwidthPackageAttributeResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyCenBandwidthPackageAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyCenBandwidthPackageAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyCenBandwidthPackageAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyCenBandwidthPackageAttributeResponse) SetHeaders(v map[string]*string) *ModifyCenBandwidthPackageAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyCenBandwidthPackageAttributeResponse) SetBody(v *ModifyCenBandwidthPackageAttributeResponseBody) *ModifyCenBandwidthPackageAttributeResponse {
	s.Body = v
	return s
}

type ModifyCenBandwidthPackageSpecRequest struct {
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	CenBandwidthPackageId *string `json:"CenBandwidthPackageId,omitempty" xml:"CenBandwidthPackageId,omitempty"`
	Bandwidth             *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
}

func (s ModifyCenBandwidthPackageSpecRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyCenBandwidthPackageSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyCenBandwidthPackageSpecRequest) SetOwnerAccount(v string) *ModifyCenBandwidthPackageSpecRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyCenBandwidthPackageSpecRequest) SetOwnerId(v int64) *ModifyCenBandwidthPackageSpecRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyCenBandwidthPackageSpecRequest) SetResourceOwnerAccount(v string) *ModifyCenBandwidthPackageSpecRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyCenBandwidthPackageSpecRequest) SetResourceOwnerId(v int64) *ModifyCenBandwidthPackageSpecRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyCenBandwidthPackageSpecRequest) SetCenBandwidthPackageId(v string) *ModifyCenBandwidthPackageSpecRequest {
	s.CenBandwidthPackageId = &v
	return s
}

func (s *ModifyCenBandwidthPackageSpecRequest) SetBandwidth(v int32) *ModifyCenBandwidthPackageSpecRequest {
	s.Bandwidth = &v
	return s
}

type ModifyCenBandwidthPackageSpecResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCenBandwidthPackageSpecResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyCenBandwidthPackageSpecResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCenBandwidthPackageSpecResponseBody) SetRequestId(v string) *ModifyCenBandwidthPackageSpecResponseBody {
	s.RequestId = &v
	return s
}

type ModifyCenBandwidthPackageSpecResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyCenBandwidthPackageSpecResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyCenBandwidthPackageSpecResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyCenBandwidthPackageSpecResponse) GoString() string {
	return s.String()
}

func (s *ModifyCenBandwidthPackageSpecResponse) SetHeaders(v map[string]*string) *ModifyCenBandwidthPackageSpecResponse {
	s.Headers = v
	return s
}

func (s *ModifyCenBandwidthPackageSpecResponse) SetBody(v *ModifyCenBandwidthPackageSpecResponseBody) *ModifyCenBandwidthPackageSpecResponse {
	s.Body = v
	return s
}

type ModifyCenRouteMapRequest struct {
	OwnerAccount                       *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                            *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount               *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId                    *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	CenId                              *string   `json:"CenId,omitempty" xml:"CenId,omitempty"`
	CenRegionId                        *string   `json:"CenRegionId,omitempty" xml:"CenRegionId,omitempty"`
	RouteMapId                         *string   `json:"RouteMapId,omitempty" xml:"RouteMapId,omitempty"`
	Description                        *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	MapResult                          *string   `json:"MapResult,omitempty" xml:"MapResult,omitempty"`
	NextPriority                       *int32    `json:"NextPriority,omitempty" xml:"NextPriority,omitempty"`
	CidrMatchMode                      *string   `json:"CidrMatchMode,omitempty" xml:"CidrMatchMode,omitempty"`
	AsPathMatchMode                    *string   `json:"AsPathMatchMode,omitempty" xml:"AsPathMatchMode,omitempty"`
	CommunityMatchMode                 *string   `json:"CommunityMatchMode,omitempty" xml:"CommunityMatchMode,omitempty"`
	CommunityOperateMode               *string   `json:"CommunityOperateMode,omitempty" xml:"CommunityOperateMode,omitempty"`
	Preference                         *int32    `json:"Preference,omitempty" xml:"Preference,omitempty"`
	Priority                           *int32    `json:"Priority,omitempty" xml:"Priority,omitempty"`
	SourceInstanceIdsReverseMatch      *bool     `json:"SourceInstanceIdsReverseMatch,omitempty" xml:"SourceInstanceIdsReverseMatch,omitempty"`
	DestinationInstanceIdsReverseMatch *bool     `json:"DestinationInstanceIdsReverseMatch,omitempty" xml:"DestinationInstanceIdsReverseMatch,omitempty"`
	GatewayZoneId                      *string   `json:"GatewayZoneId,omitempty" xml:"GatewayZoneId,omitempty"`
	MatchAddressType                   *string   `json:"MatchAddressType,omitempty" xml:"MatchAddressType,omitempty"`
	SourceInstanceIds                  []*string `json:"SourceInstanceIds,omitempty" xml:"SourceInstanceIds,omitempty" type:"Repeated"`
	DestinationInstanceIds             []*string `json:"DestinationInstanceIds,omitempty" xml:"DestinationInstanceIds,omitempty" type:"Repeated"`
	SourceRouteTableIds                []*string `json:"SourceRouteTableIds,omitempty" xml:"SourceRouteTableIds,omitempty" type:"Repeated"`
	DestinationRouteTableIds           []*string `json:"DestinationRouteTableIds,omitempty" xml:"DestinationRouteTableIds,omitempty" type:"Repeated"`
	SourceRegionIds                    []*string `json:"SourceRegionIds,omitempty" xml:"SourceRegionIds,omitempty" type:"Repeated"`
	SourceChildInstanceTypes           []*string `json:"SourceChildInstanceTypes,omitempty" xml:"SourceChildInstanceTypes,omitempty" type:"Repeated"`
	DestinationChildInstanceTypes      []*string `json:"DestinationChildInstanceTypes,omitempty" xml:"DestinationChildInstanceTypes,omitempty" type:"Repeated"`
	DestinationCidrBlocks              []*string `json:"DestinationCidrBlocks,omitempty" xml:"DestinationCidrBlocks,omitempty" type:"Repeated"`
	RouteTypes                         []*string `json:"RouteTypes,omitempty" xml:"RouteTypes,omitempty" type:"Repeated"`
	MatchAsns                          []*int    `json:"MatchAsns,omitempty" xml:"MatchAsns,omitempty" type:"Repeated"`
	MatchCommunitySet                  []*string `json:"MatchCommunitySet,omitempty" xml:"MatchCommunitySet,omitempty" type:"Repeated"`
	OperateCommunitySet                []*string `json:"OperateCommunitySet,omitempty" xml:"OperateCommunitySet,omitempty" type:"Repeated"`
	PrependAsPath                      []*int    `json:"PrependAsPath,omitempty" xml:"PrependAsPath,omitempty" type:"Repeated"`
	DestinationRegionIds               []*string `json:"DestinationRegionIds,omitempty" xml:"DestinationRegionIds,omitempty" type:"Repeated"`
	SourceZoneIds                      []*string `json:"SourceZoneIds,omitempty" xml:"SourceZoneIds,omitempty" type:"Repeated"`
	OriginalRouteTableIds              []*string `json:"OriginalRouteTableIds,omitempty" xml:"OriginalRouteTableIds,omitempty" type:"Repeated"`
}

func (s ModifyCenRouteMapRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyCenRouteMapRequest) GoString() string {
	return s.String()
}

func (s *ModifyCenRouteMapRequest) SetOwnerAccount(v string) *ModifyCenRouteMapRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetOwnerId(v int64) *ModifyCenRouteMapRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetResourceOwnerAccount(v string) *ModifyCenRouteMapRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetResourceOwnerId(v int64) *ModifyCenRouteMapRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetCenId(v string) *ModifyCenRouteMapRequest {
	s.CenId = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetCenRegionId(v string) *ModifyCenRouteMapRequest {
	s.CenRegionId = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetRouteMapId(v string) *ModifyCenRouteMapRequest {
	s.RouteMapId = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetDescription(v string) *ModifyCenRouteMapRequest {
	s.Description = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetMapResult(v string) *ModifyCenRouteMapRequest {
	s.MapResult = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetNextPriority(v int32) *ModifyCenRouteMapRequest {
	s.NextPriority = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetCidrMatchMode(v string) *ModifyCenRouteMapRequest {
	s.CidrMatchMode = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetAsPathMatchMode(v string) *ModifyCenRouteMapRequest {
	s.AsPathMatchMode = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetCommunityMatchMode(v string) *ModifyCenRouteMapRequest {
	s.CommunityMatchMode = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetCommunityOperateMode(v string) *ModifyCenRouteMapRequest {
	s.CommunityOperateMode = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetPreference(v int32) *ModifyCenRouteMapRequest {
	s.Preference = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetPriority(v int32) *ModifyCenRouteMapRequest {
	s.Priority = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetSourceInstanceIdsReverseMatch(v bool) *ModifyCenRouteMapRequest {
	s.SourceInstanceIdsReverseMatch = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetDestinationInstanceIdsReverseMatch(v bool) *ModifyCenRouteMapRequest {
	s.DestinationInstanceIdsReverseMatch = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetGatewayZoneId(v string) *ModifyCenRouteMapRequest {
	s.GatewayZoneId = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetMatchAddressType(v string) *ModifyCenRouteMapRequest {
	s.MatchAddressType = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetSourceInstanceIds(v []*string) *ModifyCenRouteMapRequest {
	s.SourceInstanceIds = v
	return s
}

func (s *ModifyCenRouteMapRequest) SetDestinationInstanceIds(v []*string) *ModifyCenRouteMapRequest {
	s.DestinationInstanceIds = v
	return s
}

func (s *ModifyCenRouteMapRequest) SetSourceRouteTableIds(v []*string) *ModifyCenRouteMapRequest {
	s.SourceRouteTableIds = v
	return s
}

func (s *ModifyCenRouteMapRequest) SetDestinationRouteTableIds(v []*string) *ModifyCenRouteMapRequest {
	s.DestinationRouteTableIds = v
	return s
}

func (s *ModifyCenRouteMapRequest) SetSourceRegionIds(v []*string) *ModifyCenRouteMapRequest {
	s.SourceRegionIds = v
	return s
}

func (s *ModifyCenRouteMapRequest) SetSourceChildInstanceTypes(v []*string) *ModifyCenRouteMapRequest {
	s.SourceChildInstanceTypes = v
	return s
}

func (s *ModifyCenRouteMapRequest) SetDestinationChildInstanceTypes(v []*string) *ModifyCenRouteMapRequest {
	s.DestinationChildInstanceTypes = v
	return s
}

func (s *ModifyCenRouteMapRequest) SetDestinationCidrBlocks(v []*string) *ModifyCenRouteMapRequest {
	s.DestinationCidrBlocks = v
	return s
}

func (s *ModifyCenRouteMapRequest) SetRouteTypes(v []*string) *ModifyCenRouteMapRequest {
	s.RouteTypes = v
	return s
}

func (s *ModifyCenRouteMapRequest) SetMatchAsns(v []*int) *ModifyCenRouteMapRequest {
	s.MatchAsns = v
	return s
}

func (s *ModifyCenRouteMapRequest) SetMatchCommunitySet(v []*string) *ModifyCenRouteMapRequest {
	s.MatchCommunitySet = v
	return s
}

func (s *ModifyCenRouteMapRequest) SetOperateCommunitySet(v []*string) *ModifyCenRouteMapRequest {
	s.OperateCommunitySet = v
	return s
}

func (s *ModifyCenRouteMapRequest) SetPrependAsPath(v []*int) *ModifyCenRouteMapRequest {
	s.PrependAsPath = v
	return s
}

func (s *ModifyCenRouteMapRequest) SetDestinationRegionIds(v []*string) *ModifyCenRouteMapRequest {
	s.DestinationRegionIds = v
	return s
}

func (s *ModifyCenRouteMapRequest) SetSourceZoneIds(v []*string) *ModifyCenRouteMapRequest {
	s.SourceZoneIds = v
	return s
}

func (s *ModifyCenRouteMapRequest) SetOriginalRouteTableIds(v []*string) *ModifyCenRouteMapRequest {
	s.OriginalRouteTableIds = v
	return s
}

type ModifyCenRouteMapResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCenRouteMapResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyCenRouteMapResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCenRouteMapResponseBody) SetRequestId(v string) *ModifyCenRouteMapResponseBody {
	s.RequestId = &v
	return s
}

type ModifyCenRouteMapResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyCenRouteMapResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyCenRouteMapResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyCenRouteMapResponse) GoString() string {
	return s.String()
}

func (s *ModifyCenRouteMapResponse) SetHeaders(v map[string]*string) *ModifyCenRouteMapResponse {
	s.Headers = v
	return s
}

func (s *ModifyCenRouteMapResponse) SetBody(v *ModifyCenRouteMapResponseBody) *ModifyCenRouteMapResponse {
	s.Body = v
	return s
}

type ModifyFlowLogAttributeRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	FlowLogName          *string `json:"FlowLogName,omitempty" xml:"FlowLogName,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	FlowLogId            *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
}

func (s ModifyFlowLogAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowLogAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyFlowLogAttributeRequest) SetOwnerAccount(v string) *ModifyFlowLogAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetOwnerId(v int64) *ModifyFlowLogAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetResourceOwnerAccount(v string) *ModifyFlowLogAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetResourceOwnerId(v int64) *ModifyFlowLogAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetClientToken(v string) *ModifyFlowLogAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetFlowLogName(v string) *ModifyFlowLogAttributeRequest {
	s.FlowLogName = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetDescription(v string) *ModifyFlowLogAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetRegionId(v string) *ModifyFlowLogAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetFlowLogId(v string) *ModifyFlowLogAttributeRequest {
	s.FlowLogId = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetCenId(v string) *ModifyFlowLogAttributeRequest {
	s.CenId = &v
	return s
}

type ModifyFlowLogAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyFlowLogAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowLogAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFlowLogAttributeResponseBody) SetRequestId(v string) *ModifyFlowLogAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyFlowLogAttributeResponseBody) SetSuccess(v string) *ModifyFlowLogAttributeResponseBody {
	s.Success = &v
	return s
}

type ModifyFlowLogAttributeResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyFlowLogAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyFlowLogAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowLogAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyFlowLogAttributeResponse) SetHeaders(v map[string]*string) *ModifyFlowLogAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyFlowLogAttributeResponse) SetBody(v *ModifyFlowLogAttributeResponseBody) *ModifyFlowLogAttributeResponse {
	s.Body = v
	return s
}

type PublishRouteEntriesRequest struct {
	ResourceOwnerAccount      *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId           *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	CenId                     *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ChildInstanceId           *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	ChildInstanceType         *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	ChildInstanceRegionId     *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	ChildInstanceRouteTableId *string `json:"ChildInstanceRouteTableId,omitempty" xml:"ChildInstanceRouteTableId,omitempty"`
	DestinationCidrBlock      *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
}

func (s PublishRouteEntriesRequest) String() string {
	return tea.Prettify(s)
}

func (s PublishRouteEntriesRequest) GoString() string {
	return s.String()
}

func (s *PublishRouteEntriesRequest) SetResourceOwnerAccount(v string) *PublishRouteEntriesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PublishRouteEntriesRequest) SetResourceOwnerId(v int64) *PublishRouteEntriesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *PublishRouteEntriesRequest) SetCenId(v string) *PublishRouteEntriesRequest {
	s.CenId = &v
	return s
}

func (s *PublishRouteEntriesRequest) SetChildInstanceId(v string) *PublishRouteEntriesRequest {
	s.ChildInstanceId = &v
	return s
}

func (s *PublishRouteEntriesRequest) SetChildInstanceType(v string) *PublishRouteEntriesRequest {
	s.ChildInstanceType = &v
	return s
}

func (s *PublishRouteEntriesRequest) SetChildInstanceRegionId(v string) *PublishRouteEntriesRequest {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *PublishRouteEntriesRequest) SetChildInstanceRouteTableId(v string) *PublishRouteEntriesRequest {
	s.ChildInstanceRouteTableId = &v
	return s
}

func (s *PublishRouteEntriesRequest) SetDestinationCidrBlock(v string) *PublishRouteEntriesRequest {
	s.DestinationCidrBlock = &v
	return s
}

type PublishRouteEntriesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PublishRouteEntriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PublishRouteEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *PublishRouteEntriesResponseBody) SetRequestId(v string) *PublishRouteEntriesResponseBody {
	s.RequestId = &v
	return s
}

type PublishRouteEntriesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PublishRouteEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PublishRouteEntriesResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishRouteEntriesResponse) GoString() string {
	return s.String()
}

func (s *PublishRouteEntriesResponse) SetHeaders(v map[string]*string) *PublishRouteEntriesResponse {
	s.Headers = v
	return s
}

func (s *PublishRouteEntriesResponse) SetBody(v *PublishRouteEntriesResponseBody) *PublishRouteEntriesResponse {
	s.Body = v
	return s
}

type ResolveAndRouteServiceInCenRequest struct {
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ClientToken          *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CenId                *string   `json:"CenId,omitempty" xml:"CenId,omitempty"`
	Host                 *string   `json:"Host,omitempty" xml:"Host,omitempty"`
	HostRegionId         *string   `json:"HostRegionId,omitempty" xml:"HostRegionId,omitempty"`
	UpdateInterval       *int64    `json:"UpdateInterval,omitempty" xml:"UpdateInterval,omitempty"`
	HostVpcId            *string   `json:"HostVpcId,omitempty" xml:"HostVpcId,omitempty"`
	Description          *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	AccessRegionIds      []*string `json:"AccessRegionIds,omitempty" xml:"AccessRegionIds,omitempty" type:"Repeated"`
}

func (s ResolveAndRouteServiceInCenRequest) String() string {
	return tea.Prettify(s)
}

func (s ResolveAndRouteServiceInCenRequest) GoString() string {
	return s.String()
}

func (s *ResolveAndRouteServiceInCenRequest) SetOwnerAccount(v string) *ResolveAndRouteServiceInCenRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ResolveAndRouteServiceInCenRequest) SetOwnerId(v int64) *ResolveAndRouteServiceInCenRequest {
	s.OwnerId = &v
	return s
}

func (s *ResolveAndRouteServiceInCenRequest) SetResourceOwnerAccount(v string) *ResolveAndRouteServiceInCenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ResolveAndRouteServiceInCenRequest) SetResourceOwnerId(v int64) *ResolveAndRouteServiceInCenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ResolveAndRouteServiceInCenRequest) SetClientToken(v string) *ResolveAndRouteServiceInCenRequest {
	s.ClientToken = &v
	return s
}

func (s *ResolveAndRouteServiceInCenRequest) SetCenId(v string) *ResolveAndRouteServiceInCenRequest {
	s.CenId = &v
	return s
}

func (s *ResolveAndRouteServiceInCenRequest) SetHost(v string) *ResolveAndRouteServiceInCenRequest {
	s.Host = &v
	return s
}

func (s *ResolveAndRouteServiceInCenRequest) SetHostRegionId(v string) *ResolveAndRouteServiceInCenRequest {
	s.HostRegionId = &v
	return s
}

func (s *ResolveAndRouteServiceInCenRequest) SetUpdateInterval(v int64) *ResolveAndRouteServiceInCenRequest {
	s.UpdateInterval = &v
	return s
}

func (s *ResolveAndRouteServiceInCenRequest) SetHostVpcId(v string) *ResolveAndRouteServiceInCenRequest {
	s.HostVpcId = &v
	return s
}

func (s *ResolveAndRouteServiceInCenRequest) SetDescription(v string) *ResolveAndRouteServiceInCenRequest {
	s.Description = &v
	return s
}

func (s *ResolveAndRouteServiceInCenRequest) SetAccessRegionIds(v []*string) *ResolveAndRouteServiceInCenRequest {
	s.AccessRegionIds = v
	return s
}

type ResolveAndRouteServiceInCenResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResolveAndRouteServiceInCenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResolveAndRouteServiceInCenResponseBody) GoString() string {
	return s.String()
}

func (s *ResolveAndRouteServiceInCenResponseBody) SetRequestId(v string) *ResolveAndRouteServiceInCenResponseBody {
	s.RequestId = &v
	return s
}

type ResolveAndRouteServiceInCenResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResolveAndRouteServiceInCenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResolveAndRouteServiceInCenResponse) String() string {
	return tea.Prettify(s)
}

func (s ResolveAndRouteServiceInCenResponse) GoString() string {
	return s.String()
}

func (s *ResolveAndRouteServiceInCenResponse) SetHeaders(v map[string]*string) *ResolveAndRouteServiceInCenResponse {
	s.Headers = v
	return s
}

func (s *ResolveAndRouteServiceInCenResponse) SetBody(v *ResolveAndRouteServiceInCenResponseBody) *ResolveAndRouteServiceInCenResponse {
	s.Body = v
	return s
}

type RoutePrivateZoneInCenToVpcRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	AccessRegionId       *string `json:"AccessRegionId,omitempty" xml:"AccessRegionId,omitempty"`
	HostRegionId         *string `json:"HostRegionId,omitempty" xml:"HostRegionId,omitempty"`
	HostVpcId            *string `json:"HostVpcId,omitempty" xml:"HostVpcId,omitempty"`
}

func (s RoutePrivateZoneInCenToVpcRequest) String() string {
	return tea.Prettify(s)
}

func (s RoutePrivateZoneInCenToVpcRequest) GoString() string {
	return s.String()
}

func (s *RoutePrivateZoneInCenToVpcRequest) SetOwnerAccount(v string) *RoutePrivateZoneInCenToVpcRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RoutePrivateZoneInCenToVpcRequest) SetOwnerId(v int64) *RoutePrivateZoneInCenToVpcRequest {
	s.OwnerId = &v
	return s
}

func (s *RoutePrivateZoneInCenToVpcRequest) SetResourceOwnerAccount(v string) *RoutePrivateZoneInCenToVpcRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RoutePrivateZoneInCenToVpcRequest) SetResourceOwnerId(v int64) *RoutePrivateZoneInCenToVpcRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RoutePrivateZoneInCenToVpcRequest) SetCenId(v string) *RoutePrivateZoneInCenToVpcRequest {
	s.CenId = &v
	return s
}

func (s *RoutePrivateZoneInCenToVpcRequest) SetAccessRegionId(v string) *RoutePrivateZoneInCenToVpcRequest {
	s.AccessRegionId = &v
	return s
}

func (s *RoutePrivateZoneInCenToVpcRequest) SetHostRegionId(v string) *RoutePrivateZoneInCenToVpcRequest {
	s.HostRegionId = &v
	return s
}

func (s *RoutePrivateZoneInCenToVpcRequest) SetHostVpcId(v string) *RoutePrivateZoneInCenToVpcRequest {
	s.HostVpcId = &v
	return s
}

type RoutePrivateZoneInCenToVpcResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RoutePrivateZoneInCenToVpcResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RoutePrivateZoneInCenToVpcResponseBody) GoString() string {
	return s.String()
}

func (s *RoutePrivateZoneInCenToVpcResponseBody) SetRequestId(v string) *RoutePrivateZoneInCenToVpcResponseBody {
	s.RequestId = &v
	return s
}

type RoutePrivateZoneInCenToVpcResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RoutePrivateZoneInCenToVpcResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RoutePrivateZoneInCenToVpcResponse) String() string {
	return tea.Prettify(s)
}

func (s RoutePrivateZoneInCenToVpcResponse) GoString() string {
	return s.String()
}

func (s *RoutePrivateZoneInCenToVpcResponse) SetHeaders(v map[string]*string) *RoutePrivateZoneInCenToVpcResponse {
	s.Headers = v
	return s
}

func (s *RoutePrivateZoneInCenToVpcResponse) SetBody(v *RoutePrivateZoneInCenToVpcResponseBody) *RoutePrivateZoneInCenToVpcResponse {
	s.Body = v
	return s
}

type SetCenInterRegionBandwidthLimitRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	LocalRegionId        *string `json:"LocalRegionId,omitempty" xml:"LocalRegionId,omitempty"`
	OppositeRegionId     *string `json:"OppositeRegionId,omitempty" xml:"OppositeRegionId,omitempty"`
	BandwidthLimit       *int64  `json:"BandwidthLimit,omitempty" xml:"BandwidthLimit,omitempty"`
	BandwidthPackageId   *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
}

func (s SetCenInterRegionBandwidthLimitRequest) String() string {
	return tea.Prettify(s)
}

func (s SetCenInterRegionBandwidthLimitRequest) GoString() string {
	return s.String()
}

func (s *SetCenInterRegionBandwidthLimitRequest) SetOwnerAccount(v string) *SetCenInterRegionBandwidthLimitRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetCenInterRegionBandwidthLimitRequest) SetOwnerId(v int64) *SetCenInterRegionBandwidthLimitRequest {
	s.OwnerId = &v
	return s
}

func (s *SetCenInterRegionBandwidthLimitRequest) SetResourceOwnerAccount(v string) *SetCenInterRegionBandwidthLimitRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetCenInterRegionBandwidthLimitRequest) SetResourceOwnerId(v int64) *SetCenInterRegionBandwidthLimitRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetCenInterRegionBandwidthLimitRequest) SetCenId(v string) *SetCenInterRegionBandwidthLimitRequest {
	s.CenId = &v
	return s
}

func (s *SetCenInterRegionBandwidthLimitRequest) SetLocalRegionId(v string) *SetCenInterRegionBandwidthLimitRequest {
	s.LocalRegionId = &v
	return s
}

func (s *SetCenInterRegionBandwidthLimitRequest) SetOppositeRegionId(v string) *SetCenInterRegionBandwidthLimitRequest {
	s.OppositeRegionId = &v
	return s
}

func (s *SetCenInterRegionBandwidthLimitRequest) SetBandwidthLimit(v int64) *SetCenInterRegionBandwidthLimitRequest {
	s.BandwidthLimit = &v
	return s
}

func (s *SetCenInterRegionBandwidthLimitRequest) SetBandwidthPackageId(v string) *SetCenInterRegionBandwidthLimitRequest {
	s.BandwidthPackageId = &v
	return s
}

type SetCenInterRegionBandwidthLimitResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetCenInterRegionBandwidthLimitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetCenInterRegionBandwidthLimitResponseBody) GoString() string {
	return s.String()
}

func (s *SetCenInterRegionBandwidthLimitResponseBody) SetRequestId(v string) *SetCenInterRegionBandwidthLimitResponseBody {
	s.RequestId = &v
	return s
}

type SetCenInterRegionBandwidthLimitResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetCenInterRegionBandwidthLimitResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetCenInterRegionBandwidthLimitResponse) String() string {
	return tea.Prettify(s)
}

func (s SetCenInterRegionBandwidthLimitResponse) GoString() string {
	return s.String()
}

func (s *SetCenInterRegionBandwidthLimitResponse) SetHeaders(v map[string]*string) *SetCenInterRegionBandwidthLimitResponse {
	s.Headers = v
	return s
}

func (s *SetCenInterRegionBandwidthLimitResponse) SetBody(v *SetCenInterRegionBandwidthLimitResponseBody) *SetCenInterRegionBandwidthLimitResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	OwnerAccount         *string                   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64                    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string                   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceType         *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ResourceId           []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
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

type TempUpgradeCenBandwidthPackageSpecRequest struct {
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	CenBandwidthPackageId *string `json:"CenBandwidthPackageId,omitempty" xml:"CenBandwidthPackageId,omitempty"`
	Bandwidth             *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	EndTime               *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s TempUpgradeCenBandwidthPackageSpecRequest) String() string {
	return tea.Prettify(s)
}

func (s TempUpgradeCenBandwidthPackageSpecRequest) GoString() string {
	return s.String()
}

func (s *TempUpgradeCenBandwidthPackageSpecRequest) SetOwnerAccount(v string) *TempUpgradeCenBandwidthPackageSpecRequest {
	s.OwnerAccount = &v
	return s
}

func (s *TempUpgradeCenBandwidthPackageSpecRequest) SetOwnerId(v int64) *TempUpgradeCenBandwidthPackageSpecRequest {
	s.OwnerId = &v
	return s
}

func (s *TempUpgradeCenBandwidthPackageSpecRequest) SetResourceOwnerAccount(v string) *TempUpgradeCenBandwidthPackageSpecRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TempUpgradeCenBandwidthPackageSpecRequest) SetResourceOwnerId(v int64) *TempUpgradeCenBandwidthPackageSpecRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TempUpgradeCenBandwidthPackageSpecRequest) SetCenBandwidthPackageId(v string) *TempUpgradeCenBandwidthPackageSpecRequest {
	s.CenBandwidthPackageId = &v
	return s
}

func (s *TempUpgradeCenBandwidthPackageSpecRequest) SetBandwidth(v int32) *TempUpgradeCenBandwidthPackageSpecRequest {
	s.Bandwidth = &v
	return s
}

func (s *TempUpgradeCenBandwidthPackageSpecRequest) SetEndTime(v string) *TempUpgradeCenBandwidthPackageSpecRequest {
	s.EndTime = &v
	return s
}

type TempUpgradeCenBandwidthPackageSpecResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TempUpgradeCenBandwidthPackageSpecResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TempUpgradeCenBandwidthPackageSpecResponseBody) GoString() string {
	return s.String()
}

func (s *TempUpgradeCenBandwidthPackageSpecResponseBody) SetRequestId(v string) *TempUpgradeCenBandwidthPackageSpecResponseBody {
	s.RequestId = &v
	return s
}

type TempUpgradeCenBandwidthPackageSpecResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TempUpgradeCenBandwidthPackageSpecResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TempUpgradeCenBandwidthPackageSpecResponse) String() string {
	return tea.Prettify(s)
}

func (s TempUpgradeCenBandwidthPackageSpecResponse) GoString() string {
	return s.String()
}

func (s *TempUpgradeCenBandwidthPackageSpecResponse) SetHeaders(v map[string]*string) *TempUpgradeCenBandwidthPackageSpecResponse {
	s.Headers = v
	return s
}

func (s *TempUpgradeCenBandwidthPackageSpecResponse) SetBody(v *TempUpgradeCenBandwidthPackageSpecResponseBody) *TempUpgradeCenBandwidthPackageSpecResponse {
	s.Body = v
	return s
}

type UnassociateCenBandwidthPackageRequest struct {
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	CenId                 *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	CenBandwidthPackageId *string `json:"CenBandwidthPackageId,omitempty" xml:"CenBandwidthPackageId,omitempty"`
}

func (s UnassociateCenBandwidthPackageRequest) String() string {
	return tea.Prettify(s)
}

func (s UnassociateCenBandwidthPackageRequest) GoString() string {
	return s.String()
}

func (s *UnassociateCenBandwidthPackageRequest) SetOwnerAccount(v string) *UnassociateCenBandwidthPackageRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UnassociateCenBandwidthPackageRequest) SetOwnerId(v int64) *UnassociateCenBandwidthPackageRequest {
	s.OwnerId = &v
	return s
}

func (s *UnassociateCenBandwidthPackageRequest) SetResourceOwnerAccount(v string) *UnassociateCenBandwidthPackageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnassociateCenBandwidthPackageRequest) SetResourceOwnerId(v int64) *UnassociateCenBandwidthPackageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UnassociateCenBandwidthPackageRequest) SetCenId(v string) *UnassociateCenBandwidthPackageRequest {
	s.CenId = &v
	return s
}

func (s *UnassociateCenBandwidthPackageRequest) SetCenBandwidthPackageId(v string) *UnassociateCenBandwidthPackageRequest {
	s.CenBandwidthPackageId = &v
	return s
}

type UnassociateCenBandwidthPackageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnassociateCenBandwidthPackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnassociateCenBandwidthPackageResponseBody) GoString() string {
	return s.String()
}

func (s *UnassociateCenBandwidthPackageResponseBody) SetRequestId(v string) *UnassociateCenBandwidthPackageResponseBody {
	s.RequestId = &v
	return s
}

type UnassociateCenBandwidthPackageResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnassociateCenBandwidthPackageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnassociateCenBandwidthPackageResponse) String() string {
	return tea.Prettify(s)
}

func (s UnassociateCenBandwidthPackageResponse) GoString() string {
	return s.String()
}

func (s *UnassociateCenBandwidthPackageResponse) SetHeaders(v map[string]*string) *UnassociateCenBandwidthPackageResponse {
	s.Headers = v
	return s
}

func (s *UnassociateCenBandwidthPackageResponse) SetBody(v *UnassociateCenBandwidthPackageResponseBody) *UnassociateCenBandwidthPackageResponse {
	s.Body = v
	return s
}

type UnroutePrivateZoneInCenToVpcRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	AccessRegionId       *string `json:"AccessRegionId,omitempty" xml:"AccessRegionId,omitempty"`
}

func (s UnroutePrivateZoneInCenToVpcRequest) String() string {
	return tea.Prettify(s)
}

func (s UnroutePrivateZoneInCenToVpcRequest) GoString() string {
	return s.String()
}

func (s *UnroutePrivateZoneInCenToVpcRequest) SetOwnerAccount(v string) *UnroutePrivateZoneInCenToVpcRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UnroutePrivateZoneInCenToVpcRequest) SetOwnerId(v int64) *UnroutePrivateZoneInCenToVpcRequest {
	s.OwnerId = &v
	return s
}

func (s *UnroutePrivateZoneInCenToVpcRequest) SetResourceOwnerAccount(v string) *UnroutePrivateZoneInCenToVpcRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnroutePrivateZoneInCenToVpcRequest) SetResourceOwnerId(v int64) *UnroutePrivateZoneInCenToVpcRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UnroutePrivateZoneInCenToVpcRequest) SetCenId(v string) *UnroutePrivateZoneInCenToVpcRequest {
	s.CenId = &v
	return s
}

func (s *UnroutePrivateZoneInCenToVpcRequest) SetAccessRegionId(v string) *UnroutePrivateZoneInCenToVpcRequest {
	s.AccessRegionId = &v
	return s
}

type UnroutePrivateZoneInCenToVpcResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnroutePrivateZoneInCenToVpcResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnroutePrivateZoneInCenToVpcResponseBody) GoString() string {
	return s.String()
}

func (s *UnroutePrivateZoneInCenToVpcResponseBody) SetRequestId(v string) *UnroutePrivateZoneInCenToVpcResponseBody {
	s.RequestId = &v
	return s
}

type UnroutePrivateZoneInCenToVpcResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnroutePrivateZoneInCenToVpcResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnroutePrivateZoneInCenToVpcResponse) String() string {
	return tea.Prettify(s)
}

func (s UnroutePrivateZoneInCenToVpcResponse) GoString() string {
	return s.String()
}

func (s *UnroutePrivateZoneInCenToVpcResponse) SetHeaders(v map[string]*string) *UnroutePrivateZoneInCenToVpcResponse {
	s.Headers = v
	return s
}

func (s *UnroutePrivateZoneInCenToVpcResponse) SetBody(v *UnroutePrivateZoneInCenToVpcResponseBody) *UnroutePrivateZoneInCenToVpcResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceType         *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	All                  *bool     `json:"All,omitempty" xml:"All,omitempty"`
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	TagKey               []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetOwnerAccount(v string) *UntagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UntagResourcesRequest) SetOwnerId(v int64) *UntagResourcesRequest {
	s.OwnerId = &v
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

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
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

type WithdrawPublishedRouteEntriesRequest struct {
	ResourceOwnerAccount      *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId           *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	CenId                     *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ChildInstanceId           *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	ChildInstanceType         *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	ChildInstanceRegionId     *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	ChildInstanceRouteTableId *string `json:"ChildInstanceRouteTableId,omitempty" xml:"ChildInstanceRouteTableId,omitempty"`
	DestinationCidrBlock      *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
}

func (s WithdrawPublishedRouteEntriesRequest) String() string {
	return tea.Prettify(s)
}

func (s WithdrawPublishedRouteEntriesRequest) GoString() string {
	return s.String()
}

func (s *WithdrawPublishedRouteEntriesRequest) SetResourceOwnerAccount(v string) *WithdrawPublishedRouteEntriesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *WithdrawPublishedRouteEntriesRequest) SetResourceOwnerId(v int64) *WithdrawPublishedRouteEntriesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *WithdrawPublishedRouteEntriesRequest) SetCenId(v string) *WithdrawPublishedRouteEntriesRequest {
	s.CenId = &v
	return s
}

func (s *WithdrawPublishedRouteEntriesRequest) SetChildInstanceId(v string) *WithdrawPublishedRouteEntriesRequest {
	s.ChildInstanceId = &v
	return s
}

func (s *WithdrawPublishedRouteEntriesRequest) SetChildInstanceType(v string) *WithdrawPublishedRouteEntriesRequest {
	s.ChildInstanceType = &v
	return s
}

func (s *WithdrawPublishedRouteEntriesRequest) SetChildInstanceRegionId(v string) *WithdrawPublishedRouteEntriesRequest {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *WithdrawPublishedRouteEntriesRequest) SetChildInstanceRouteTableId(v string) *WithdrawPublishedRouteEntriesRequest {
	s.ChildInstanceRouteTableId = &v
	return s
}

func (s *WithdrawPublishedRouteEntriesRequest) SetDestinationCidrBlock(v string) *WithdrawPublishedRouteEntriesRequest {
	s.DestinationCidrBlock = &v
	return s
}

type WithdrawPublishedRouteEntriesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s WithdrawPublishedRouteEntriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s WithdrawPublishedRouteEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *WithdrawPublishedRouteEntriesResponseBody) SetRequestId(v string) *WithdrawPublishedRouteEntriesResponseBody {
	s.RequestId = &v
	return s
}

type WithdrawPublishedRouteEntriesResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *WithdrawPublishedRouteEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s WithdrawPublishedRouteEntriesResponse) String() string {
	return tea.Prettify(s)
}

func (s WithdrawPublishedRouteEntriesResponse) GoString() string {
	return s.String()
}

func (s *WithdrawPublishedRouteEntriesResponse) SetHeaders(v map[string]*string) *WithdrawPublishedRouteEntriesResponse {
	s.Headers = v
	return s
}

func (s *WithdrawPublishedRouteEntriesResponse) SetBody(v *WithdrawPublishedRouteEntriesResponseBody) *WithdrawPublishedRouteEntriesResponse {
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
		"cn-qingdao":                  tea.String("cbn.aliyuncs.com"),
		"cn-beijing":                  tea.String("cbn.aliyuncs.com"),
		"cn-chengdu":                  tea.String("cbn.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("cbn.aliyuncs.com"),
		"cn-huhehaote":                tea.String("cbn.aliyuncs.com"),
		"cn-hangzhou":                 tea.String("cbn.aliyuncs.com"),
		"cn-shanghai":                 tea.String("cbn.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("cbn.aliyuncs.com"),
		"cn-heyuan":                   tea.String("cbn.aliyuncs.com"),
		"cn-wulanchabu":               tea.String("cbn.aliyuncs.com"),
		"cn-hongkong":                 tea.String("cbn.aliyuncs.com"),
		"ap-southeast-1":              tea.String("cbn.aliyuncs.com"),
		"ap-southeast-2":              tea.String("cbn.aliyuncs.com"),
		"ap-southeast-3":              tea.String("cbn.aliyuncs.com"),
		"ap-southeast-5":              tea.String("cbn.aliyuncs.com"),
		"ap-northeast-1":              tea.String("cbn.aliyuncs.com"),
		"eu-west-1":                   tea.String("cbn.aliyuncs.com"),
		"us-west-1":                   tea.String("cbn.aliyuncs.com"),
		"us-east-1":                   tea.String("cbn.aliyuncs.com"),
		"eu-central-1":                tea.String("cbn.aliyuncs.com"),
		"me-east-1":                   tea.String("cbn.aliyuncs.com"),
		"ap-south-1":                  tea.String("cbn.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("cbn.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("cbn.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("cbn.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("cbn.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("cbn.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("cbn.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("cbn.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("cbn.aliyuncs.com"),
		"cn-edge-1":                   tea.String("cbn.aliyuncs.com"),
		"cn-fujian":                   tea.String("cbn.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("cbn.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("cbn.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("cbn.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("cbn.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("cbn.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("cbn.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("cbn.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("cbn.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("cbn.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("cbn.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("cbn.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("cbn.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("cbn.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("cbn.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("cbn.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("cbn.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("cbn.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("cbn.aliyuncs.com"),
		"cn-wuhan":                    tea.String("cbn.aliyuncs.com"),
		"cn-yushanfang":               tea.String("cbn.aliyuncs.com"),
		"cn-zhangbei":                 tea.String("cbn.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("cbn.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("cbn.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("cbn.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("cbn.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("cbn-share.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("cbn"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) ActiveFlowLogWithOptions(request *ActiveFlowLogRequest, runtime *util.RuntimeOptions) (_result *ActiveFlowLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ActiveFlowLogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ActiveFlowLog"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ActiveFlowLog(request *ActiveFlowLogRequest) (_result *ActiveFlowLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ActiveFlowLogResponse{}
	_body, _err := client.ActiveFlowLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AssociateCenBandwidthPackageWithOptions(request *AssociateCenBandwidthPackageRequest, runtime *util.RuntimeOptions) (_result *AssociateCenBandwidthPackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AssociateCenBandwidthPackageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AssociateCenBandwidthPackage"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AssociateCenBandwidthPackage(request *AssociateCenBandwidthPackageRequest) (_result *AssociateCenBandwidthPackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssociateCenBandwidthPackageResponse{}
	_body, _err := client.AssociateCenBandwidthPackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AttachCenChildInstanceWithOptions(request *AttachCenChildInstanceRequest, runtime *util.RuntimeOptions) (_result *AttachCenChildInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AttachCenChildInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AttachCenChildInstance"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttachCenChildInstance(request *AttachCenChildInstanceRequest) (_result *AttachCenChildInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachCenChildInstanceResponse{}
	_body, _err := client.AttachCenChildInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCenWithOptions(request *CreateCenRequest, runtime *util.RuntimeOptions) (_result *CreateCenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateCenResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateCen"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCen(request *CreateCenRequest) (_result *CreateCenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCenResponse{}
	_body, _err := client.CreateCenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCenBandwidthPackageWithOptions(request *CreateCenBandwidthPackageRequest, runtime *util.RuntimeOptions) (_result *CreateCenBandwidthPackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateCenBandwidthPackageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateCenBandwidthPackage"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCenBandwidthPackage(request *CreateCenBandwidthPackageRequest) (_result *CreateCenBandwidthPackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCenBandwidthPackageResponse{}
	_body, _err := client.CreateCenBandwidthPackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCenChildInstanceRouteEntryToCenWithOptions(request *CreateCenChildInstanceRouteEntryToCenRequest, runtime *util.RuntimeOptions) (_result *CreateCenChildInstanceRouteEntryToCenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateCenChildInstanceRouteEntryToCenResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateCenChildInstanceRouteEntryToCen"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCenChildInstanceRouteEntryToCen(request *CreateCenChildInstanceRouteEntryToCenRequest) (_result *CreateCenChildInstanceRouteEntryToCenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCenChildInstanceRouteEntryToCenResponse{}
	_body, _err := client.CreateCenChildInstanceRouteEntryToCenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCenRouteMapWithOptions(request *CreateCenRouteMapRequest, runtime *util.RuntimeOptions) (_result *CreateCenRouteMapResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateCenRouteMapResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateCenRouteMap"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCenRouteMap(request *CreateCenRouteMapRequest) (_result *CreateCenRouteMapResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCenRouteMapResponse{}
	_body, _err := client.CreateCenRouteMapWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFlowlogWithOptions(request *CreateFlowlogRequest, runtime *util.RuntimeOptions) (_result *CreateFlowlogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateFlowlogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateFlowlog"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFlowlog(request *CreateFlowlogRequest) (_result *CreateFlowlogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFlowlogResponse{}
	_body, _err := client.CreateFlowlogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeactiveFlowLogWithOptions(request *DeactiveFlowLogRequest, runtime *util.RuntimeOptions) (_result *DeactiveFlowLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeactiveFlowLogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeactiveFlowLog"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeactiveFlowLog(request *DeactiveFlowLogRequest) (_result *DeactiveFlowLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeactiveFlowLogResponse{}
	_body, _err := client.DeactiveFlowLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCenWithOptions(request *DeleteCenRequest, runtime *util.RuntimeOptions) (_result *DeleteCenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteCenResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteCen"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCen(request *DeleteCenRequest) (_result *DeleteCenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCenResponse{}
	_body, _err := client.DeleteCenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCenBandwidthPackageWithOptions(request *DeleteCenBandwidthPackageRequest, runtime *util.RuntimeOptions) (_result *DeleteCenBandwidthPackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteCenBandwidthPackageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteCenBandwidthPackage"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCenBandwidthPackage(request *DeleteCenBandwidthPackageRequest) (_result *DeleteCenBandwidthPackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCenBandwidthPackageResponse{}
	_body, _err := client.DeleteCenBandwidthPackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCenChildInstanceRouteEntryToCenWithOptions(request *DeleteCenChildInstanceRouteEntryToCenRequest, runtime *util.RuntimeOptions) (_result *DeleteCenChildInstanceRouteEntryToCenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteCenChildInstanceRouteEntryToCenResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteCenChildInstanceRouteEntryToCen"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCenChildInstanceRouteEntryToCen(request *DeleteCenChildInstanceRouteEntryToCenRequest) (_result *DeleteCenChildInstanceRouteEntryToCenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCenChildInstanceRouteEntryToCenResponse{}
	_body, _err := client.DeleteCenChildInstanceRouteEntryToCenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCenRouteMapWithOptions(request *DeleteCenRouteMapRequest, runtime *util.RuntimeOptions) (_result *DeleteCenRouteMapResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteCenRouteMapResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteCenRouteMap"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCenRouteMap(request *DeleteCenRouteMapRequest) (_result *DeleteCenRouteMapResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCenRouteMapResponse{}
	_body, _err := client.DeleteCenRouteMapWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFlowlogWithOptions(request *DeleteFlowlogRequest, runtime *util.RuntimeOptions) (_result *DeleteFlowlogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteFlowlogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteFlowlog"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFlowlog(request *DeleteFlowlogRequest) (_result *DeleteFlowlogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFlowlogResponse{}
	_body, _err := client.DeleteFlowlogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRouteServiceInCenWithOptions(request *DeleteRouteServiceInCenRequest, runtime *util.RuntimeOptions) (_result *DeleteRouteServiceInCenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteRouteServiceInCenResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteRouteServiceInCen"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRouteServiceInCen(request *DeleteRouteServiceInCenRequest) (_result *DeleteRouteServiceInCenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRouteServiceInCenResponse{}
	_body, _err := client.DeleteRouteServiceInCenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCenAttachedChildInstanceAttributeWithOptions(request *DescribeCenAttachedChildInstanceAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeCenAttachedChildInstanceAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCenAttachedChildInstanceAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCenAttachedChildInstanceAttribute"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCenAttachedChildInstanceAttribute(request *DescribeCenAttachedChildInstanceAttributeRequest) (_result *DescribeCenAttachedChildInstanceAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCenAttachedChildInstanceAttributeResponse{}
	_body, _err := client.DescribeCenAttachedChildInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCenAttachedChildInstancesWithOptions(request *DescribeCenAttachedChildInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeCenAttachedChildInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCenAttachedChildInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCenAttachedChildInstances"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCenAttachedChildInstances(request *DescribeCenAttachedChildInstancesRequest) (_result *DescribeCenAttachedChildInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCenAttachedChildInstancesResponse{}
	_body, _err := client.DescribeCenAttachedChildInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCenBandwidthPackagesWithOptions(request *DescribeCenBandwidthPackagesRequest, runtime *util.RuntimeOptions) (_result *DescribeCenBandwidthPackagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCenBandwidthPackagesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCenBandwidthPackages"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCenBandwidthPackages(request *DescribeCenBandwidthPackagesRequest) (_result *DescribeCenBandwidthPackagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCenBandwidthPackagesResponse{}
	_body, _err := client.DescribeCenBandwidthPackagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCenChildInstanceRouteEntriesWithOptions(request *DescribeCenChildInstanceRouteEntriesRequest, runtime *util.RuntimeOptions) (_result *DescribeCenChildInstanceRouteEntriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCenChildInstanceRouteEntriesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCenChildInstanceRouteEntries"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCenChildInstanceRouteEntries(request *DescribeCenChildInstanceRouteEntriesRequest) (_result *DescribeCenChildInstanceRouteEntriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCenChildInstanceRouteEntriesResponse{}
	_body, _err := client.DescribeCenChildInstanceRouteEntriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCenGeographicSpanRemainingBandwidthWithOptions(request *DescribeCenGeographicSpanRemainingBandwidthRequest, runtime *util.RuntimeOptions) (_result *DescribeCenGeographicSpanRemainingBandwidthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCenGeographicSpanRemainingBandwidthResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCenGeographicSpanRemainingBandwidth"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCenGeographicSpanRemainingBandwidth(request *DescribeCenGeographicSpanRemainingBandwidthRequest) (_result *DescribeCenGeographicSpanRemainingBandwidthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCenGeographicSpanRemainingBandwidthResponse{}
	_body, _err := client.DescribeCenGeographicSpanRemainingBandwidthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCenGeographicSpansWithOptions(request *DescribeCenGeographicSpansRequest, runtime *util.RuntimeOptions) (_result *DescribeCenGeographicSpansResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCenGeographicSpansResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCenGeographicSpans"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCenGeographicSpans(request *DescribeCenGeographicSpansRequest) (_result *DescribeCenGeographicSpansResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCenGeographicSpansResponse{}
	_body, _err := client.DescribeCenGeographicSpansWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCenInterRegionBandwidthLimitsWithOptions(request *DescribeCenInterRegionBandwidthLimitsRequest, runtime *util.RuntimeOptions) (_result *DescribeCenInterRegionBandwidthLimitsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCenInterRegionBandwidthLimitsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCenInterRegionBandwidthLimits"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCenInterRegionBandwidthLimits(request *DescribeCenInterRegionBandwidthLimitsRequest) (_result *DescribeCenInterRegionBandwidthLimitsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCenInterRegionBandwidthLimitsResponse{}
	_body, _err := client.DescribeCenInterRegionBandwidthLimitsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCenPrivateZoneRoutesWithOptions(request *DescribeCenPrivateZoneRoutesRequest, runtime *util.RuntimeOptions) (_result *DescribeCenPrivateZoneRoutesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCenPrivateZoneRoutesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCenPrivateZoneRoutes"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCenPrivateZoneRoutes(request *DescribeCenPrivateZoneRoutesRequest) (_result *DescribeCenPrivateZoneRoutesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCenPrivateZoneRoutesResponse{}
	_body, _err := client.DescribeCenPrivateZoneRoutesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCenRegionDomainRouteEntriesWithOptions(request *DescribeCenRegionDomainRouteEntriesRequest, runtime *util.RuntimeOptions) (_result *DescribeCenRegionDomainRouteEntriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCenRegionDomainRouteEntriesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCenRegionDomainRouteEntries"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCenRegionDomainRouteEntries(request *DescribeCenRegionDomainRouteEntriesRequest) (_result *DescribeCenRegionDomainRouteEntriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCenRegionDomainRouteEntriesResponse{}
	_body, _err := client.DescribeCenRegionDomainRouteEntriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCenRouteMapsWithOptions(request *DescribeCenRouteMapsRequest, runtime *util.RuntimeOptions) (_result *DescribeCenRouteMapsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCenRouteMapsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCenRouteMaps"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCenRouteMaps(request *DescribeCenRouteMapsRequest) (_result *DescribeCenRouteMapsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCenRouteMapsResponse{}
	_body, _err := client.DescribeCenRouteMapsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCensWithOptions(request *DescribeCensRequest, runtime *util.RuntimeOptions) (_result *DescribeCensResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCensResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCens"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCens(request *DescribeCensRequest) (_result *DescribeCensResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCensResponse{}
	_body, _err := client.DescribeCensWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCenVbrHealthCheckWithOptions(request *DescribeCenVbrHealthCheckRequest, runtime *util.RuntimeOptions) (_result *DescribeCenVbrHealthCheckResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCenVbrHealthCheckResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCenVbrHealthCheck"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCenVbrHealthCheck(request *DescribeCenVbrHealthCheckRequest) (_result *DescribeCenVbrHealthCheckResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCenVbrHealthCheckResponse{}
	_body, _err := client.DescribeCenVbrHealthCheckWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeChildInstanceRegionsWithOptions(request *DescribeChildInstanceRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeChildInstanceRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeChildInstanceRegionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeChildInstanceRegions"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeChildInstanceRegions(request *DescribeChildInstanceRegionsRequest) (_result *DescribeChildInstanceRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeChildInstanceRegionsResponse{}
	_body, _err := client.DescribeChildInstanceRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFlowlogsWithOptions(request *DescribeFlowlogsRequest, runtime *util.RuntimeOptions) (_result *DescribeFlowlogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFlowlogsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFlowlogs"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFlowlogs(request *DescribeFlowlogsRequest) (_result *DescribeFlowlogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFlowlogsResponse{}
	_body, _err := client.DescribeFlowlogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGeographicRegionMembershipWithOptions(request *DescribeGeographicRegionMembershipRequest, runtime *util.RuntimeOptions) (_result *DescribeGeographicRegionMembershipResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeGeographicRegionMembershipResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeGeographicRegionMembership"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGeographicRegionMembership(request *DescribeGeographicRegionMembershipRequest) (_result *DescribeGeographicRegionMembershipResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGeographicRegionMembershipResponse{}
	_body, _err := client.DescribeGeographicRegionMembershipWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGrantRulesToCenWithOptions(request *DescribeGrantRulesToCenRequest, runtime *util.RuntimeOptions) (_result *DescribeGrantRulesToCenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeGrantRulesToCenResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeGrantRulesToCen"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGrantRulesToCen(request *DescribeGrantRulesToCenRequest) (_result *DescribeGrantRulesToCenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGrantRulesToCenResponse{}
	_body, _err := client.DescribeGrantRulesToCenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePublishedRouteEntriesWithOptions(request *DescribePublishedRouteEntriesRequest, runtime *util.RuntimeOptions) (_result *DescribePublishedRouteEntriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePublishedRouteEntriesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePublishedRouteEntries"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePublishedRouteEntries(request *DescribePublishedRouteEntriesRequest) (_result *DescribePublishedRouteEntriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePublishedRouteEntriesResponse{}
	_body, _err := client.DescribePublishedRouteEntriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRouteConflictWithOptions(request *DescribeRouteConflictRequest, runtime *util.RuntimeOptions) (_result *DescribeRouteConflictResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRouteConflictResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRouteConflict"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRouteConflict(request *DescribeRouteConflictRequest) (_result *DescribeRouteConflictResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRouteConflictResponse{}
	_body, _err := client.DescribeRouteConflictWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRouteServicesInCenWithOptions(request *DescribeRouteServicesInCenRequest, runtime *util.RuntimeOptions) (_result *DescribeRouteServicesInCenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRouteServicesInCenResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRouteServicesInCen"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRouteServicesInCen(request *DescribeRouteServicesInCenRequest) (_result *DescribeRouteServicesInCenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRouteServicesInCenResponse{}
	_body, _err := client.DescribeRouteServicesInCenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetachCenChildInstanceWithOptions(request *DetachCenChildInstanceRequest, runtime *util.RuntimeOptions) (_result *DetachCenChildInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetachCenChildInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetachCenChildInstance"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetachCenChildInstance(request *DetachCenChildInstanceRequest) (_result *DetachCenChildInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachCenChildInstanceResponse{}
	_body, _err := client.DetachCenChildInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableCenVbrHealthCheckWithOptions(request *DisableCenVbrHealthCheckRequest, runtime *util.RuntimeOptions) (_result *DisableCenVbrHealthCheckResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DisableCenVbrHealthCheckResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DisableCenVbrHealthCheck"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableCenVbrHealthCheck(request *DisableCenVbrHealthCheckRequest) (_result *DisableCenVbrHealthCheckResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableCenVbrHealthCheckResponse{}
	_body, _err := client.DisableCenVbrHealthCheckWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableCenVbrHealthCheckWithOptions(request *EnableCenVbrHealthCheckRequest, runtime *util.RuntimeOptions) (_result *EnableCenVbrHealthCheckResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnableCenVbrHealthCheckResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableCenVbrHealthCheck"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableCenVbrHealthCheck(request *EnableCenVbrHealthCheckRequest) (_result *EnableCenVbrHealthCheckResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableCenVbrHealthCheckResponse{}
	_body, _err := client.EnableCenVbrHealthCheckWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("ListTagResources"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ModifyCenAttributeWithOptions(request *ModifyCenAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyCenAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyCenAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyCenAttribute"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyCenAttribute(request *ModifyCenAttributeRequest) (_result *ModifyCenAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyCenAttributeResponse{}
	_body, _err := client.ModifyCenAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyCenBandwidthPackageAttributeWithOptions(request *ModifyCenBandwidthPackageAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyCenBandwidthPackageAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyCenBandwidthPackageAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyCenBandwidthPackageAttribute"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyCenBandwidthPackageAttribute(request *ModifyCenBandwidthPackageAttributeRequest) (_result *ModifyCenBandwidthPackageAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyCenBandwidthPackageAttributeResponse{}
	_body, _err := client.ModifyCenBandwidthPackageAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyCenBandwidthPackageSpecWithOptions(request *ModifyCenBandwidthPackageSpecRequest, runtime *util.RuntimeOptions) (_result *ModifyCenBandwidthPackageSpecResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyCenBandwidthPackageSpecResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyCenBandwidthPackageSpec"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyCenBandwidthPackageSpec(request *ModifyCenBandwidthPackageSpecRequest) (_result *ModifyCenBandwidthPackageSpecResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyCenBandwidthPackageSpecResponse{}
	_body, _err := client.ModifyCenBandwidthPackageSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyCenRouteMapWithOptions(request *ModifyCenRouteMapRequest, runtime *util.RuntimeOptions) (_result *ModifyCenRouteMapResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyCenRouteMapResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyCenRouteMap"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyCenRouteMap(request *ModifyCenRouteMapRequest) (_result *ModifyCenRouteMapResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyCenRouteMapResponse{}
	_body, _err := client.ModifyCenRouteMapWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyFlowLogAttributeWithOptions(request *ModifyFlowLogAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyFlowLogAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyFlowLogAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyFlowLogAttribute"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyFlowLogAttribute(request *ModifyFlowLogAttributeRequest) (_result *ModifyFlowLogAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyFlowLogAttributeResponse{}
	_body, _err := client.ModifyFlowLogAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PublishRouteEntriesWithOptions(request *PublishRouteEntriesRequest, runtime *util.RuntimeOptions) (_result *PublishRouteEntriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PublishRouteEntriesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PublishRouteEntries"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PublishRouteEntries(request *PublishRouteEntriesRequest) (_result *PublishRouteEntriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PublishRouteEntriesResponse{}
	_body, _err := client.PublishRouteEntriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResolveAndRouteServiceInCenWithOptions(request *ResolveAndRouteServiceInCenRequest, runtime *util.RuntimeOptions) (_result *ResolveAndRouteServiceInCenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ResolveAndRouteServiceInCenResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ResolveAndRouteServiceInCen"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResolveAndRouteServiceInCen(request *ResolveAndRouteServiceInCenRequest) (_result *ResolveAndRouteServiceInCenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResolveAndRouteServiceInCenResponse{}
	_body, _err := client.ResolveAndRouteServiceInCenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RoutePrivateZoneInCenToVpcWithOptions(request *RoutePrivateZoneInCenToVpcRequest, runtime *util.RuntimeOptions) (_result *RoutePrivateZoneInCenToVpcResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RoutePrivateZoneInCenToVpcResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RoutePrivateZoneInCenToVpc"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RoutePrivateZoneInCenToVpc(request *RoutePrivateZoneInCenToVpcRequest) (_result *RoutePrivateZoneInCenToVpcResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RoutePrivateZoneInCenToVpcResponse{}
	_body, _err := client.RoutePrivateZoneInCenToVpcWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetCenInterRegionBandwidthLimitWithOptions(request *SetCenInterRegionBandwidthLimitRequest, runtime *util.RuntimeOptions) (_result *SetCenInterRegionBandwidthLimitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetCenInterRegionBandwidthLimitResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetCenInterRegionBandwidthLimit"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetCenInterRegionBandwidthLimit(request *SetCenInterRegionBandwidthLimitRequest) (_result *SetCenInterRegionBandwidthLimitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetCenInterRegionBandwidthLimitResponse{}
	_body, _err := client.SetCenInterRegionBandwidthLimitWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("TagResources"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) TempUpgradeCenBandwidthPackageSpecWithOptions(request *TempUpgradeCenBandwidthPackageSpecRequest, runtime *util.RuntimeOptions) (_result *TempUpgradeCenBandwidthPackageSpecResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TempUpgradeCenBandwidthPackageSpecResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TempUpgradeCenBandwidthPackageSpec"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TempUpgradeCenBandwidthPackageSpec(request *TempUpgradeCenBandwidthPackageSpecRequest) (_result *TempUpgradeCenBandwidthPackageSpecResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TempUpgradeCenBandwidthPackageSpecResponse{}
	_body, _err := client.TempUpgradeCenBandwidthPackageSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnassociateCenBandwidthPackageWithOptions(request *UnassociateCenBandwidthPackageRequest, runtime *util.RuntimeOptions) (_result *UnassociateCenBandwidthPackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UnassociateCenBandwidthPackageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UnassociateCenBandwidthPackage"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnassociateCenBandwidthPackage(request *UnassociateCenBandwidthPackageRequest) (_result *UnassociateCenBandwidthPackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnassociateCenBandwidthPackageResponse{}
	_body, _err := client.UnassociateCenBandwidthPackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnroutePrivateZoneInCenToVpcWithOptions(request *UnroutePrivateZoneInCenToVpcRequest, runtime *util.RuntimeOptions) (_result *UnroutePrivateZoneInCenToVpcResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UnroutePrivateZoneInCenToVpcResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UnroutePrivateZoneInCenToVpc"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnroutePrivateZoneInCenToVpc(request *UnroutePrivateZoneInCenToVpcRequest) (_result *UnroutePrivateZoneInCenToVpcResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnroutePrivateZoneInCenToVpcResponse{}
	_body, _err := client.UnroutePrivateZoneInCenToVpcWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UntagResources"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) WithdrawPublishedRouteEntriesWithOptions(request *WithdrawPublishedRouteEntriesRequest, runtime *util.RuntimeOptions) (_result *WithdrawPublishedRouteEntriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &WithdrawPublishedRouteEntriesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("WithdrawPublishedRouteEntries"), tea.String("2017-09-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) WithdrawPublishedRouteEntries(request *WithdrawPublishedRouteEntriesRequest) (_result *WithdrawPublishedRouteEntriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &WithdrawPublishedRouteEntriesResponse{}
	_body, _err := client.WithdrawPublishedRouteEntriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
