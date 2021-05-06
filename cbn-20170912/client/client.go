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
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ActiveFlowLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ActiveFlowLogResponseBody) GoString() string {
	return s.String()
}

func (s *ActiveFlowLogResponseBody) SetSuccess(v string) *ActiveFlowLogResponseBody {
	s.Success = &v
	return s
}

func (s *ActiveFlowLogResponseBody) SetRequestId(v string) *ActiveFlowLogResponseBody {
	s.RequestId = &v
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

type CreateCenResponseBody struct {
	CenId     *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCenResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCenResponseBody) SetCenId(v string) *CreateCenResponseBody {
	s.CenId = &v
	return s
}

func (s *CreateCenResponseBody) SetRequestId(v string) *CreateCenResponseBody {
	s.RequestId = &v
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
	CenBandwidthPackageId      *string `json:"CenBandwidthPackageId,omitempty" xml:"CenBandwidthPackageId,omitempty"`
	CenBandwidthPackageOrderId *string `json:"CenBandwidthPackageOrderId,omitempty" xml:"CenBandwidthPackageOrderId,omitempty"`
	RequestId                  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCenBandwidthPackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCenBandwidthPackageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCenBandwidthPackageResponseBody) SetCenBandwidthPackageId(v string) *CreateCenBandwidthPackageResponseBody {
	s.CenBandwidthPackageId = &v
	return s
}

func (s *CreateCenBandwidthPackageResponseBody) SetCenBandwidthPackageOrderId(v string) *CreateCenBandwidthPackageResponseBody {
	s.CenBandwidthPackageOrderId = &v
	return s
}

func (s *CreateCenBandwidthPackageResponseBody) SetRequestId(v string) *CreateCenBandwidthPackageResponseBody {
	s.RequestId = &v
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

type CreateCenRouteMapResponseBody struct {
	RouteMapId *string `json:"RouteMapId,omitempty" xml:"RouteMapId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCenRouteMapResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCenRouteMapResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCenRouteMapResponseBody) SetRouteMapId(v string) *CreateCenRouteMapResponseBody {
	s.RouteMapId = &v
	return s
}

func (s *CreateCenRouteMapResponseBody) SetRequestId(v string) *CreateCenRouteMapResponseBody {
	s.RequestId = &v
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
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
	FlowLogId *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
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

func (s *CreateFlowlogResponseBody) SetSuccess(v string) *CreateFlowlogResponseBody {
	s.Success = &v
	return s
}

func (s *CreateFlowlogResponseBody) SetFlowLogId(v string) *CreateFlowlogResponseBody {
	s.FlowLogId = &v
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
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeactiveFlowLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeactiveFlowLogResponseBody) GoString() string {
	return s.String()
}

func (s *DeactiveFlowLogResponseBody) SetSuccess(v string) *DeactiveFlowLogResponseBody {
	s.Success = &v
	return s
}

func (s *DeactiveFlowLogResponseBody) SetRequestId(v string) *DeactiveFlowLogResponseBody {
	s.RequestId = &v
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
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFlowlogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowlogResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFlowlogResponseBody) SetSuccess(v string) *DeleteFlowlogResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteFlowlogResponseBody) SetRequestId(v string) *DeleteFlowlogResponseBody {
	s.RequestId = &v
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

type DescribeCenAttachedChildInstanceAttributeResponseBody struct {
	Status                  *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ChildInstanceType       *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	ChildInstanceRegionId   *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	RequestId               *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ChildInstanceOwnerId    *int64  `json:"ChildInstanceOwnerId,omitempty" xml:"ChildInstanceOwnerId,omitempty"`
	ChildInstanceId         *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	CenId                   *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ChildInstanceName       *string `json:"ChildInstanceName,omitempty" xml:"ChildInstanceName,omitempty"`
	ChildInstanceAttachTime *string `json:"ChildInstanceAttachTime,omitempty" xml:"ChildInstanceAttachTime,omitempty"`
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

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) SetChildInstanceRegionId(v string) *DescribeCenAttachedChildInstanceAttributeResponseBody {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) SetRequestId(v string) *DescribeCenAttachedChildInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) SetChildInstanceOwnerId(v int64) *DescribeCenAttachedChildInstanceAttributeResponseBody {
	s.ChildInstanceOwnerId = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) SetChildInstanceId(v string) *DescribeCenAttachedChildInstanceAttributeResponseBody {
	s.ChildInstanceId = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) SetCenId(v string) *DescribeCenAttachedChildInstanceAttributeResponseBody {
	s.CenId = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) SetChildInstanceName(v string) *DescribeCenAttachedChildInstanceAttributeResponseBody {
	s.ChildInstanceName = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) SetChildInstanceAttachTime(v string) *DescribeCenAttachedChildInstanceAttributeResponseBody {
	s.ChildInstanceAttachTime = &v
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
	PageSize       *int32                                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber     *int32                                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount     *int32                                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	ChildInstances *DescribeCenAttachedChildInstancesResponseBodyChildInstances `json:"ChildInstances,omitempty" xml:"ChildInstances,omitempty" type:"Struct"`
}

func (s DescribeCenAttachedChildInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenAttachedChildInstancesResponseBody) GoString() string {
	return s.String()
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

func (s *DescribeCenAttachedChildInstancesResponseBody) SetTotalCount(v int32) *DescribeCenAttachedChildInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesResponseBody) SetChildInstances(v *DescribeCenAttachedChildInstancesResponseBodyChildInstances) *DescribeCenAttachedChildInstancesResponseBody {
	s.ChildInstances = v
	return s
}

type DescribeCenAttachedChildInstancesResponseBodyChildInstances struct {
	ChildInstance []*DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance `json:"ChildInstance,omitempty" xml:"ChildInstance,omitempty" type:"Repeated"`
}

func (s DescribeCenAttachedChildInstancesResponseBodyChildInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenAttachedChildInstancesResponseBodyChildInstances) GoString() string {
	return s.String()
}

func (s *DescribeCenAttachedChildInstancesResponseBodyChildInstances) SetChildInstance(v []*DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance) *DescribeCenAttachedChildInstancesResponseBodyChildInstances {
	s.ChildInstance = v
	return s
}

type DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance struct {
	ChildInstanceType       *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	Status                  *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ChildInstanceRegionId   *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	ChildInstanceOwnerId    *int64  `json:"ChildInstanceOwnerId,omitempty" xml:"ChildInstanceOwnerId,omitempty"`
	ChildInstanceId         *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	CenId                   *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ChildInstanceAttachTime *string `json:"ChildInstanceAttachTime,omitempty" xml:"ChildInstanceAttachTime,omitempty"`
}

func (s DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance) GoString() string {
	return s.String()
}

func (s *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance) SetChildInstanceType(v string) *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance {
	s.ChildInstanceType = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance) SetStatus(v string) *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance {
	s.Status = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance) SetChildInstanceRegionId(v string) *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance) SetChildInstanceOwnerId(v int64) *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance {
	s.ChildInstanceOwnerId = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance) SetChildInstanceId(v string) *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance {
	s.ChildInstanceId = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance) SetCenId(v string) *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance {
	s.CenId = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance) SetChildInstanceAttachTime(v string) *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance {
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
	PageSize             *int32                                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId            *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber           *int32                                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount           *int32                                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	CenBandwidthPackages *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages `json:"CenBandwidthPackages,omitempty" xml:"CenBandwidthPackages,omitempty" type:"Struct"`
}

func (s DescribeCenBandwidthPackagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenBandwidthPackagesResponseBody) GoString() string {
	return s.String()
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

func (s *DescribeCenBandwidthPackagesResponseBody) SetTotalCount(v int32) *DescribeCenBandwidthPackagesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBody) SetCenBandwidthPackages(v *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages) *DescribeCenBandwidthPackagesResponseBody {
	s.CenBandwidthPackages = v
	return s
}

type DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages struct {
	CenBandwidthPackage []*DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage `json:"CenBandwidthPackage,omitempty" xml:"CenBandwidthPackage,omitempty" type:"Repeated"`
}

func (s DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages) GoString() string {
	return s.String()
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages) SetCenBandwidthPackage(v []*DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages {
	s.CenBandwidthPackage = v
	return s
}

type DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage struct {
	ReservationActiveTime           *string                                                                                                         `json:"ReservationActiveTime,omitempty" xml:"ReservationActiveTime,omitempty"`
	Status                          *string                                                                                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	CreationTime                    *string                                                                                                         `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	ReservationOrderType            *string                                                                                                         `json:"ReservationOrderType,omitempty" xml:"ReservationOrderType,omitempty"`
	BandwidthPackageChargeType      *string                                                                                                         `json:"BandwidthPackageChargeType,omitempty" xml:"BandwidthPackageChargeType,omitempty"`
	CenBandwidthPackageId           *string                                                                                                         `json:"CenBandwidthPackageId,omitempty" xml:"CenBandwidthPackageId,omitempty"`
	ReservationInternetChargeType   *string                                                                                                         `json:"ReservationInternetChargeType,omitempty" xml:"ReservationInternetChargeType,omitempty"`
	Ratio                           *string                                                                                                         `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	GeographicRegionAId             *string                                                                                                         `json:"GeographicRegionAId,omitempty" xml:"GeographicRegionAId,omitempty"`
	TypeFor95                       *string                                                                                                         `json:"TypeFor95,omitempty" xml:"TypeFor95,omitempty"`
	Bandwidth                       *int64                                                                                                          `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	Description                     *string                                                                                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	ExpiredTime                     *string                                                                                                         `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	ReservationBandwidth            *string                                                                                                         `json:"ReservationBandwidth,omitempty" xml:"ReservationBandwidth,omitempty"`
	GeographicSpanId                *string                                                                                                         `json:"GeographicSpanId,omitempty" xml:"GeographicSpanId,omitempty"`
	GeographicRegionBId             *string                                                                                                         `json:"GeographicRegionBId,omitempty" xml:"GeographicRegionBId,omitempty"`
	IsCrossBorder                   *bool                                                                                                           `json:"IsCrossBorder,omitempty" xml:"IsCrossBorder,omitempty"`
	BusinessStatus                  *string                                                                                                         `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	Name                            *string                                                                                                         `json:"Name,omitempty" xml:"Name,omitempty"`
	HasReservationData              *string                                                                                                         `json:"HasReservationData,omitempty" xml:"HasReservationData,omitempty"`
	OrginInterRegionBandwidthLimits *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimits `json:"OrginInterRegionBandwidthLimits,omitempty" xml:"OrginInterRegionBandwidthLimits,omitempty" type:"Struct"`
	CenIds                          *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageCenIds                          `json:"CenIds,omitempty" xml:"CenIds,omitempty" type:"Struct"`
}

func (s DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) GoString() string {
	return s.String()
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetReservationActiveTime(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.ReservationActiveTime = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetStatus(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.Status = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetCreationTime(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.CreationTime = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetReservationOrderType(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.ReservationOrderType = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetBandwidthPackageChargeType(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.BandwidthPackageChargeType = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetCenBandwidthPackageId(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.CenBandwidthPackageId = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetReservationInternetChargeType(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.ReservationInternetChargeType = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetRatio(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.Ratio = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetGeographicRegionAId(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.GeographicRegionAId = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetTypeFor95(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.TypeFor95 = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetBandwidth(v int64) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.Bandwidth = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetDescription(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.Description = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetExpiredTime(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetReservationBandwidth(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.ReservationBandwidth = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetGeographicSpanId(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.GeographicSpanId = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetGeographicRegionBId(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.GeographicRegionBId = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetIsCrossBorder(v bool) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.IsCrossBorder = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetBusinessStatus(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.BusinessStatus = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetName(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.Name = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetHasReservationData(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.HasReservationData = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetOrginInterRegionBandwidthLimits(v *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimits) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.OrginInterRegionBandwidthLimits = v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetCenIds(v *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageCenIds) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.CenIds = v
	return s
}

type DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimits struct {
	OrginInterRegionBandwidthLimit []*DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimitsOrginInterRegionBandwidthLimit `json:"OrginInterRegionBandwidthLimit,omitempty" xml:"OrginInterRegionBandwidthLimit,omitempty" type:"Repeated"`
}

func (s DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimits) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimits) GoString() string {
	return s.String()
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimits) SetOrginInterRegionBandwidthLimit(v []*DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimitsOrginInterRegionBandwidthLimit) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimits {
	s.OrginInterRegionBandwidthLimit = v
	return s
}

type DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimitsOrginInterRegionBandwidthLimit struct {
	BandwidthLimit   *string `json:"BandwidthLimit,omitempty" xml:"BandwidthLimit,omitempty"`
	OppositeRegionId *string `json:"OppositeRegionId,omitempty" xml:"OppositeRegionId,omitempty"`
	GeographicSpanId *string `json:"GeographicSpanId,omitempty" xml:"GeographicSpanId,omitempty"`
	LocalRegionId    *string `json:"LocalRegionId,omitempty" xml:"LocalRegionId,omitempty"`
}

func (s DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimitsOrginInterRegionBandwidthLimit) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimitsOrginInterRegionBandwidthLimit) GoString() string {
	return s.String()
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimitsOrginInterRegionBandwidthLimit) SetBandwidthLimit(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimitsOrginInterRegionBandwidthLimit {
	s.BandwidthLimit = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimitsOrginInterRegionBandwidthLimit) SetOppositeRegionId(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimitsOrginInterRegionBandwidthLimit {
	s.OppositeRegionId = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimitsOrginInterRegionBandwidthLimit) SetGeographicSpanId(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimitsOrginInterRegionBandwidthLimit {
	s.GeographicSpanId = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimitsOrginInterRegionBandwidthLimit) SetLocalRegionId(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimitsOrginInterRegionBandwidthLimit {
	s.LocalRegionId = &v
	return s
}

type DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageCenIds struct {
	CenId []*string `json:"CenId,omitempty" xml:"CenId,omitempty" type:"Repeated"`
}

func (s DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageCenIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageCenIds) GoString() string {
	return s.String()
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageCenIds) SetCenId(v []*string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageCenIds {
	s.CenId = v
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
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	CenId                 *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ChildInstanceId       *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	ChildInstanceType     *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty"`
	PageNumber            *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize              *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ChildInstanceRegionId *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
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

type DescribeCenChildInstanceRouteEntriesResponseBody struct {
	PageSize        *int32                                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId       *string                                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber      *int32                                                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount      *int32                                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	CenRouteEntries *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries `json:"CenRouteEntries,omitempty" xml:"CenRouteEntries,omitempty" type:"Struct"`
}

func (s DescribeCenChildInstanceRouteEntriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenChildInstanceRouteEntriesResponseBody) GoString() string {
	return s.String()
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

func (s *DescribeCenChildInstanceRouteEntriesResponseBody) SetTotalCount(v int32) *DescribeCenChildInstanceRouteEntriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBody) SetCenRouteEntries(v *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries) *DescribeCenChildInstanceRouteEntriesResponseBody {
	s.CenRouteEntries = v
	return s
}

type DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries struct {
	CenRouteEntry []*DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry `json:"CenRouteEntry,omitempty" xml:"CenRouteEntry,omitempty" type:"Repeated"`
}

func (s DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries) GoString() string {
	return s.String()
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries) SetCenRouteEntry(v []*DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries {
	s.CenRouteEntry = v
	return s
}

type DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry struct {
	Status               *string                                                                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	Type                 *string                                                                                         `json:"Type,omitempty" xml:"Type,omitempty"`
	PublishStatus        *string                                                                                         `json:"PublishStatus,omitempty" xml:"PublishStatus,omitempty"`
	NextHopType          *string                                                                                         `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	OperationalMode      *bool                                                                                           `json:"OperationalMode,omitempty" xml:"OperationalMode,omitempty"`
	NextHopRegionId      *string                                                                                         `json:"NextHopRegionId,omitempty" xml:"NextHopRegionId,omitempty"`
	NextHopInstanceId    *string                                                                                         `json:"NextHopInstanceId,omitempty" xml:"NextHopInstanceId,omitempty"`
	DestinationCidrBlock *string                                                                                         `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	RouteTableId         *string                                                                                         `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
	CenRouteMapRecords   *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecords `json:"CenRouteMapRecords,omitempty" xml:"CenRouteMapRecords,omitempty" type:"Struct"`
	Conflicts            *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflicts          `json:"Conflicts,omitempty" xml:"Conflicts,omitempty" type:"Struct"`
	Communities          *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCommunities        `json:"Communities,omitempty" xml:"Communities,omitempty" type:"Struct"`
	AsPaths              *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryAsPaths            `json:"AsPaths,omitempty" xml:"AsPaths,omitempty" type:"Struct"`
}

func (s DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) GoString() string {
	return s.String()
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetStatus(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.Status = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetType(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.Type = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetPublishStatus(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.PublishStatus = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetNextHopType(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.NextHopType = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetOperationalMode(v bool) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.OperationalMode = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetNextHopRegionId(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.NextHopRegionId = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetNextHopInstanceId(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.NextHopInstanceId = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetDestinationCidrBlock(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetRouteTableId(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.RouteTableId = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetCenRouteMapRecords(v *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecords) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.CenRouteMapRecords = v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetConflicts(v *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflicts) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.Conflicts = v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetCommunities(v *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCommunities) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.Communities = v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetAsPaths(v *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryAsPaths) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.AsPaths = v
	return s
}

type DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecords struct {
	CenRouteMapRecord []*DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord `json:"CenRouteMapRecord,omitempty" xml:"CenRouteMapRecord,omitempty" type:"Repeated"`
}

func (s DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecords) GoString() string {
	return s.String()
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecords) SetCenRouteMapRecord(v []*DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecords {
	s.CenRouteMapRecord = v
	return s
}

type DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord struct {
	RouteMapId *string `json:"RouteMapId,omitempty" xml:"RouteMapId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord) GoString() string {
	return s.String()
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord) SetRouteMapId(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord {
	s.RouteMapId = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord) SetRegionId(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord {
	s.RegionId = &v
	return s
}

type DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflicts struct {
	Conflict []*DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflictsConflict `json:"Conflict,omitempty" xml:"Conflict,omitempty" type:"Repeated"`
}

func (s DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflicts) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflicts) GoString() string {
	return s.String()
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflicts) SetConflict(v []*DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflictsConflict) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflicts {
	s.Conflict = v
	return s
}

type DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflictsConflict struct {
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType         *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflictsConflict) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflictsConflict) GoString() string {
	return s.String()
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflictsConflict) SetStatus(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflictsConflict {
	s.Status = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflictsConflict) SetDestinationCidrBlock(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflictsConflict {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflictsConflict) SetInstanceId(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflictsConflict {
	s.InstanceId = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflictsConflict) SetInstanceType(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflictsConflict {
	s.InstanceType = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflictsConflict) SetRegionId(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflictsConflict {
	s.RegionId = &v
	return s
}

type DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCommunities struct {
	Community []*string `json:"Community,omitempty" xml:"Community,omitempty" type:"Repeated"`
}

func (s DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCommunities) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCommunities) GoString() string {
	return s.String()
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCommunities) SetCommunity(v []*string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCommunities {
	s.Community = v
	return s
}

type DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryAsPaths struct {
	AsPath []*string `json:"AsPath,omitempty" xml:"AsPath,omitempty" type:"Repeated"`
}

func (s DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryAsPaths) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryAsPaths) GoString() string {
	return s.String()
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryAsPaths) SetAsPath(v []*string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryAsPaths {
	s.AsPath = v
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
	RemainingBandwidth *int64  `json:"RemainingBandwidth,omitempty" xml:"RemainingBandwidth,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCenGeographicSpanRemainingBandwidthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenGeographicSpanRemainingBandwidthResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCenGeographicSpanRemainingBandwidthResponseBody) SetRemainingBandwidth(v int64) *DescribeCenGeographicSpanRemainingBandwidthResponseBody {
	s.RemainingBandwidth = &v
	return s
}

func (s *DescribeCenGeographicSpanRemainingBandwidthResponseBody) SetRequestId(v string) *DescribeCenGeographicSpanRemainingBandwidthResponseBody {
	s.RequestId = &v
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
	PageSize             *int32                                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId            *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber           *int32                                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount           *int32                                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	GeographicSpanModels *DescribeCenGeographicSpansResponseBodyGeographicSpanModels `json:"GeographicSpanModels,omitempty" xml:"GeographicSpanModels,omitempty" type:"Struct"`
}

func (s DescribeCenGeographicSpansResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenGeographicSpansResponseBody) GoString() string {
	return s.String()
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

func (s *DescribeCenGeographicSpansResponseBody) SetTotalCount(v int32) *DescribeCenGeographicSpansResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCenGeographicSpansResponseBody) SetGeographicSpanModels(v *DescribeCenGeographicSpansResponseBodyGeographicSpanModels) *DescribeCenGeographicSpansResponseBody {
	s.GeographicSpanModels = v
	return s
}

type DescribeCenGeographicSpansResponseBodyGeographicSpanModels struct {
	GeographicSpanModel []*DescribeCenGeographicSpansResponseBodyGeographicSpanModelsGeographicSpanModel `json:"GeographicSpanModel,omitempty" xml:"GeographicSpanModel,omitempty" type:"Repeated"`
}

func (s DescribeCenGeographicSpansResponseBodyGeographicSpanModels) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenGeographicSpansResponseBodyGeographicSpanModels) GoString() string {
	return s.String()
}

func (s *DescribeCenGeographicSpansResponseBodyGeographicSpanModels) SetGeographicSpanModel(v []*DescribeCenGeographicSpansResponseBodyGeographicSpanModelsGeographicSpanModel) *DescribeCenGeographicSpansResponseBodyGeographicSpanModels {
	s.GeographicSpanModel = v
	return s
}

type DescribeCenGeographicSpansResponseBodyGeographicSpanModelsGeographicSpanModel struct {
	LocalGeoRegionId    *string `json:"LocalGeoRegionId,omitempty" xml:"LocalGeoRegionId,omitempty"`
	GeographicSpanId    *string `json:"GeographicSpanId,omitempty" xml:"GeographicSpanId,omitempty"`
	OppositeGeoRegionId *string `json:"OppositeGeoRegionId,omitempty" xml:"OppositeGeoRegionId,omitempty"`
}

func (s DescribeCenGeographicSpansResponseBodyGeographicSpanModelsGeographicSpanModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenGeographicSpansResponseBodyGeographicSpanModelsGeographicSpanModel) GoString() string {
	return s.String()
}

func (s *DescribeCenGeographicSpansResponseBodyGeographicSpanModelsGeographicSpanModel) SetLocalGeoRegionId(v string) *DescribeCenGeographicSpansResponseBodyGeographicSpanModelsGeographicSpanModel {
	s.LocalGeoRegionId = &v
	return s
}

func (s *DescribeCenGeographicSpansResponseBodyGeographicSpanModelsGeographicSpanModel) SetGeographicSpanId(v string) *DescribeCenGeographicSpansResponseBodyGeographicSpanModelsGeographicSpanModel {
	s.GeographicSpanId = &v
	return s
}

func (s *DescribeCenGeographicSpansResponseBodyGeographicSpanModelsGeographicSpanModel) SetOppositeGeoRegionId(v string) *DescribeCenGeographicSpansResponseBodyGeographicSpanModelsGeographicSpanModel {
	s.OppositeGeoRegionId = &v
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
	PageSize                      *int32                                                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId                     *string                                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber                    *int32                                                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount                    *int32                                                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	CenInterRegionBandwidthLimits *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimits `json:"CenInterRegionBandwidthLimits,omitempty" xml:"CenInterRegionBandwidthLimits,omitempty" type:"Struct"`
}

func (s DescribeCenInterRegionBandwidthLimitsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenInterRegionBandwidthLimitsResponseBody) GoString() string {
	return s.String()
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

func (s *DescribeCenInterRegionBandwidthLimitsResponseBody) SetTotalCount(v int32) *DescribeCenInterRegionBandwidthLimitsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBody) SetCenInterRegionBandwidthLimits(v *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimits) *DescribeCenInterRegionBandwidthLimitsResponseBody {
	s.CenInterRegionBandwidthLimits = v
	return s
}

type DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimits struct {
	CenInterRegionBandwidthLimit []*DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit `json:"CenInterRegionBandwidthLimit,omitempty" xml:"CenInterRegionBandwidthLimit,omitempty" type:"Repeated"`
}

func (s DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimits) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimits) GoString() string {
	return s.String()
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimits) SetCenInterRegionBandwidthLimit(v []*DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit) *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimits {
	s.CenInterRegionBandwidthLimit = v
	return s
}

type DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit struct {
	Status             *string `json:"Status,omitempty" xml:"Status,omitempty"`
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	OppositeRegionId   *string `json:"OppositeRegionId,omitempty" xml:"OppositeRegionId,omitempty"`
	GeographicSpanId   *string `json:"GeographicSpanId,omitempty" xml:"GeographicSpanId,omitempty"`
	CenId              *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	LocalRegionId      *string `json:"LocalRegionId,omitempty" xml:"LocalRegionId,omitempty"`
	BandwidthLimit     *int64  `json:"BandwidthLimit,omitempty" xml:"BandwidthLimit,omitempty"`
}

func (s DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit) GoString() string {
	return s.String()
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit) SetStatus(v string) *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit {
	s.Status = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit) SetBandwidthPackageId(v string) *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit {
	s.BandwidthPackageId = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit) SetOppositeRegionId(v string) *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit {
	s.OppositeRegionId = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit) SetGeographicSpanId(v string) *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit {
	s.GeographicSpanId = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit) SetCenId(v string) *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit {
	s.CenId = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit) SetLocalRegionId(v string) *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit {
	s.LocalRegionId = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit) SetBandwidthLimit(v int64) *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit {
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
	RequestId             *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PrivateZoneDnsServers *string                                                   `json:"PrivateZoneDnsServers,omitempty" xml:"PrivateZoneDnsServers,omitempty"`
	CenId                 *string                                                   `json:"CenId,omitempty" xml:"CenId,omitempty"`
	PageNumber            *int32                                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize              *int32                                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount            *int32                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PrivateZoneInfos      *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfos `json:"PrivateZoneInfos,omitempty" xml:"PrivateZoneInfos,omitempty" type:"Struct"`
}

func (s DescribeCenPrivateZoneRoutesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenPrivateZoneRoutesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCenPrivateZoneRoutesResponseBody) SetRequestId(v string) *DescribeCenPrivateZoneRoutesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesResponseBody) SetPrivateZoneDnsServers(v string) *DescribeCenPrivateZoneRoutesResponseBody {
	s.PrivateZoneDnsServers = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesResponseBody) SetCenId(v string) *DescribeCenPrivateZoneRoutesResponseBody {
	s.CenId = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesResponseBody) SetPageNumber(v int32) *DescribeCenPrivateZoneRoutesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesResponseBody) SetPageSize(v int32) *DescribeCenPrivateZoneRoutesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesResponseBody) SetTotalCount(v int32) *DescribeCenPrivateZoneRoutesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesResponseBody) SetPrivateZoneInfos(v *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfos) *DescribeCenPrivateZoneRoutesResponseBody {
	s.PrivateZoneInfos = v
	return s
}

type DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfos struct {
	PrivateZoneInfo []*DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfosPrivateZoneInfo `json:"PrivateZoneInfo,omitempty" xml:"PrivateZoneInfo,omitempty" type:"Repeated"`
}

func (s DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfos) GoString() string {
	return s.String()
}

func (s *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfos) SetPrivateZoneInfo(v []*DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfosPrivateZoneInfo) *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfos {
	s.PrivateZoneInfo = v
	return s
}

type DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfosPrivateZoneInfo struct {
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	AccessRegionId *string `json:"AccessRegionId,omitempty" xml:"AccessRegionId,omitempty"`
	HostRegionId   *string `json:"HostRegionId,omitempty" xml:"HostRegionId,omitempty"`
	HostVpcId      *string `json:"HostVpcId,omitempty" xml:"HostVpcId,omitempty"`
}

func (s DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfosPrivateZoneInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfosPrivateZoneInfo) GoString() string {
	return s.String()
}

func (s *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfosPrivateZoneInfo) SetStatus(v string) *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfosPrivateZoneInfo {
	s.Status = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfosPrivateZoneInfo) SetAccessRegionId(v string) *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfosPrivateZoneInfo {
	s.AccessRegionId = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfosPrivateZoneInfo) SetHostRegionId(v string) *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfosPrivateZoneInfo {
	s.HostRegionId = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfosPrivateZoneInfo) SetHostVpcId(v string) *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfosPrivateZoneInfo {
	s.HostVpcId = &v
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

type DescribeCenRegionDomainRouteEntriesResponseBody struct {
	PageSize        *int32                                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId       *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber      *int32                                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount      *int32                                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	CenRouteEntries *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries `json:"CenRouteEntries,omitempty" xml:"CenRouteEntries,omitempty" type:"Struct"`
}

func (s DescribeCenRegionDomainRouteEntriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenRegionDomainRouteEntriesResponseBody) GoString() string {
	return s.String()
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

func (s *DescribeCenRegionDomainRouteEntriesResponseBody) SetTotalCount(v int32) *DescribeCenRegionDomainRouteEntriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBody) SetCenRouteEntries(v *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries) *DescribeCenRegionDomainRouteEntriesResponseBody {
	s.CenRouteEntries = v
	return s
}

type DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries struct {
	CenRouteEntry []*DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry `json:"CenRouteEntry,omitempty" xml:"CenRouteEntry,omitempty" type:"Repeated"`
}

func (s DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries) GoString() string {
	return s.String()
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries) SetCenRouteEntry(v []*DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries {
	s.CenRouteEntry = v
	return s
}

type DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry struct {
	ToOtherRegionStatus   *string                                                                                           `json:"ToOtherRegionStatus,omitempty" xml:"ToOtherRegionStatus,omitempty"`
	Type                  *string                                                                                           `json:"Type,omitempty" xml:"Type,omitempty"`
	Status                *string                                                                                           `json:"Status,omitempty" xml:"Status,omitempty"`
	NextHopType           *string                                                                                           `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	NextHopInstanceId     *string                                                                                           `json:"NextHopInstanceId,omitempty" xml:"NextHopInstanceId,omitempty"`
	NextHopRegionId       *string                                                                                           `json:"NextHopRegionId,omitempty" xml:"NextHopRegionId,omitempty"`
	DestinationCidrBlock  *string                                                                                           `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	Preference            *int32                                                                                            `json:"Preference,omitempty" xml:"Preference,omitempty"`
	CenRouteMapRecords    *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecords    `json:"CenRouteMapRecords,omitempty" xml:"CenRouteMapRecords,omitempty" type:"Struct"`
	CenOutRouteMapRecords *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecords `json:"CenOutRouteMapRecords,omitempty" xml:"CenOutRouteMapRecords,omitempty" type:"Struct"`
	Communities           *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCommunities           `json:"Communities,omitempty" xml:"Communities,omitempty" type:"Struct"`
	AsPaths               *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryAsPaths               `json:"AsPaths,omitempty" xml:"AsPaths,omitempty" type:"Struct"`
}

func (s DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) GoString() string {
	return s.String()
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetToOtherRegionStatus(v string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.ToOtherRegionStatus = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetType(v string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.Type = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetStatus(v string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.Status = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetNextHopType(v string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.NextHopType = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetNextHopInstanceId(v string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.NextHopInstanceId = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetNextHopRegionId(v string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.NextHopRegionId = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetDestinationCidrBlock(v string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetPreference(v int32) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.Preference = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetCenRouteMapRecords(v *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecords) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.CenRouteMapRecords = v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetCenOutRouteMapRecords(v *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecords) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.CenOutRouteMapRecords = v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetCommunities(v *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCommunities) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.Communities = v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetAsPaths(v *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryAsPaths) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.AsPaths = v
	return s
}

type DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecords struct {
	CenRouteMapRecord []*DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord `json:"CenRouteMapRecord,omitempty" xml:"CenRouteMapRecord,omitempty" type:"Repeated"`
}

func (s DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecords) GoString() string {
	return s.String()
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecords) SetCenRouteMapRecord(v []*DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecords {
	s.CenRouteMapRecord = v
	return s
}

type DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord struct {
	RouteMapId *string `json:"RouteMapId,omitempty" xml:"RouteMapId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord) GoString() string {
	return s.String()
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord) SetRouteMapId(v string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord {
	s.RouteMapId = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord) SetRegionId(v string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord {
	s.RegionId = &v
	return s
}

type DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecords struct {
	CenOutRouteMapRecord []*DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecordsCenOutRouteMapRecord `json:"CenOutRouteMapRecord,omitempty" xml:"CenOutRouteMapRecord,omitempty" type:"Repeated"`
}

func (s DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecords) GoString() string {
	return s.String()
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecords) SetCenOutRouteMapRecord(v []*DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecordsCenOutRouteMapRecord) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecords {
	s.CenOutRouteMapRecord = v
	return s
}

type DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecordsCenOutRouteMapRecord struct {
	RouteMapId *string `json:"RouteMapId,omitempty" xml:"RouteMapId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecordsCenOutRouteMapRecord) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecordsCenOutRouteMapRecord) GoString() string {
	return s.String()
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecordsCenOutRouteMapRecord) SetRouteMapId(v string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecordsCenOutRouteMapRecord {
	s.RouteMapId = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecordsCenOutRouteMapRecord) SetRegionId(v string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecordsCenOutRouteMapRecord {
	s.RegionId = &v
	return s
}

type DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCommunities struct {
	Community []*string `json:"Community,omitempty" xml:"Community,omitempty" type:"Repeated"`
}

func (s DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCommunities) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCommunities) GoString() string {
	return s.String()
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCommunities) SetCommunity(v []*string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCommunities {
	s.Community = v
	return s
}

type DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryAsPaths struct {
	AsPath []*string `json:"AsPath,omitempty" xml:"AsPath,omitempty" type:"Repeated"`
}

func (s DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryAsPaths) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryAsPaths) GoString() string {
	return s.String()
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryAsPaths) SetAsPath(v []*string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryAsPaths {
	s.AsPath = v
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
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	RouteMapId           *string `json:"RouteMapId,omitempty" xml:"RouteMapId,omitempty"`
	CenRegionId          *string `json:"CenRegionId,omitempty" xml:"CenRegionId,omitempty"`
	TransmitDirection    *string `json:"TransmitDirection,omitempty" xml:"TransmitDirection,omitempty"`
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

type DescribeCenRouteMapsResponseBody struct {
	PageSize   *int32                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount *int32                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RouteMaps  *DescribeCenRouteMapsResponseBodyRouteMaps `json:"RouteMaps,omitempty" xml:"RouteMaps,omitempty" type:"Struct"`
}

func (s DescribeCenRouteMapsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenRouteMapsResponseBody) GoString() string {
	return s.String()
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

func (s *DescribeCenRouteMapsResponseBody) SetTotalCount(v int32) *DescribeCenRouteMapsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBody) SetRouteMaps(v *DescribeCenRouteMapsResponseBodyRouteMaps) *DescribeCenRouteMapsResponseBody {
	s.RouteMaps = v
	return s
}

type DescribeCenRouteMapsResponseBodyRouteMaps struct {
	RouteMap []*DescribeCenRouteMapsResponseBodyRouteMapsRouteMap `json:"RouteMap,omitempty" xml:"RouteMap,omitempty" type:"Repeated"`
}

func (s DescribeCenRouteMapsResponseBodyRouteMaps) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenRouteMapsResponseBodyRouteMaps) GoString() string {
	return s.String()
}

func (s *DescribeCenRouteMapsResponseBodyRouteMaps) SetRouteMap(v []*DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) *DescribeCenRouteMapsResponseBodyRouteMaps {
	s.RouteMap = v
	return s
}

type DescribeCenRouteMapsResponseBodyRouteMapsRouteMap struct {
	RouteMapId                         *string                                                                         `json:"RouteMapId,omitempty" xml:"RouteMapId,omitempty"`
	Status                             *string                                                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	TransmitDirection                  *string                                                                         `json:"TransmitDirection,omitempty" xml:"TransmitDirection,omitempty"`
	SourceInstanceIdsReverseMatch      *bool                                                                           `json:"SourceInstanceIdsReverseMatch,omitempty" xml:"SourceInstanceIdsReverseMatch,omitempty"`
	CenRegionId                        *string                                                                         `json:"CenRegionId,omitempty" xml:"CenRegionId,omitempty"`
	CenId                              *string                                                                         `json:"CenId,omitempty" xml:"CenId,omitempty"`
	Priority                           *int32                                                                          `json:"Priority,omitempty" xml:"Priority,omitempty"`
	CommunityOperateMode               *string                                                                         `json:"CommunityOperateMode,omitempty" xml:"CommunityOperateMode,omitempty"`
	MapResult                          *string                                                                         `json:"MapResult,omitempty" xml:"MapResult,omitempty"`
	CommunityMatchMode                 *string                                                                         `json:"CommunityMatchMode,omitempty" xml:"CommunityMatchMode,omitempty"`
	Description                        *string                                                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	AsPathMatchMode                    *string                                                                         `json:"AsPathMatchMode,omitempty" xml:"AsPathMatchMode,omitempty"`
	Preference                         *int32                                                                          `json:"Preference,omitempty" xml:"Preference,omitempty"`
	DestinationInstanceIdsReverseMatch *bool                                                                           `json:"DestinationInstanceIdsReverseMatch,omitempty" xml:"DestinationInstanceIdsReverseMatch,omitempty"`
	CidrMatchMode                      *string                                                                         `json:"CidrMatchMode,omitempty" xml:"CidrMatchMode,omitempty"`
	NextPriority                       *int32                                                                          `json:"NextPriority,omitempty" xml:"NextPriority,omitempty"`
	SourceRegionIds                    *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceRegionIds               `json:"SourceRegionIds,omitempty" xml:"SourceRegionIds,omitempty" type:"Struct"`
	SourceChildInstanceTypes           *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceChildInstanceTypes      `json:"SourceChildInstanceTypes,omitempty" xml:"SourceChildInstanceTypes,omitempty" type:"Struct"`
	DestinationRouteTableIds           *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationRouteTableIds      `json:"DestinationRouteTableIds,omitempty" xml:"DestinationRouteTableIds,omitempty" type:"Struct"`
	SourceInstanceIds                  *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceInstanceIds             `json:"SourceInstanceIds,omitempty" xml:"SourceInstanceIds,omitempty" type:"Struct"`
	DestinationCidrBlocks              *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationCidrBlocks         `json:"DestinationCidrBlocks,omitempty" xml:"DestinationCidrBlocks,omitempty" type:"Struct"`
	SourceRouteTableIds                *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceRouteTableIds           `json:"SourceRouteTableIds,omitempty" xml:"SourceRouteTableIds,omitempty" type:"Struct"`
	MatchCommunitySet                  *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapMatchCommunitySet             `json:"MatchCommunitySet,omitempty" xml:"MatchCommunitySet,omitempty" type:"Struct"`
	PrependAsPath                      *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapPrependAsPath                 `json:"PrependAsPath,omitempty" xml:"PrependAsPath,omitempty" type:"Struct"`
	RouteTypes                         *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapRouteTypes                    `json:"RouteTypes,omitempty" xml:"RouteTypes,omitempty" type:"Struct"`
	DestinationChildInstanceTypes      *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationChildInstanceTypes `json:"DestinationChildInstanceTypes,omitempty" xml:"DestinationChildInstanceTypes,omitempty" type:"Struct"`
	DestinationInstanceIds             *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationInstanceIds        `json:"DestinationInstanceIds,omitempty" xml:"DestinationInstanceIds,omitempty" type:"Struct"`
	MatchAsns                          *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapMatchAsns                     `json:"MatchAsns,omitempty" xml:"MatchAsns,omitempty" type:"Struct"`
	OperateCommunitySet                *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapOperateCommunitySet           `json:"OperateCommunitySet,omitempty" xml:"OperateCommunitySet,omitempty" type:"Struct"`
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) GoString() string {
	return s.String()
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetRouteMapId(v string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.RouteMapId = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetStatus(v string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.Status = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetTransmitDirection(v string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.TransmitDirection = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetSourceInstanceIdsReverseMatch(v bool) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.SourceInstanceIdsReverseMatch = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetCenRegionId(v string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.CenRegionId = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetCenId(v string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.CenId = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetPriority(v int32) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.Priority = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetCommunityOperateMode(v string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.CommunityOperateMode = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetMapResult(v string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.MapResult = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetCommunityMatchMode(v string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.CommunityMatchMode = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetDescription(v string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.Description = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetAsPathMatchMode(v string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.AsPathMatchMode = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetPreference(v int32) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.Preference = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetDestinationInstanceIdsReverseMatch(v bool) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.DestinationInstanceIdsReverseMatch = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetCidrMatchMode(v string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.CidrMatchMode = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetNextPriority(v int32) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.NextPriority = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetSourceRegionIds(v *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceRegionIds) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.SourceRegionIds = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetSourceChildInstanceTypes(v *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceChildInstanceTypes) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.SourceChildInstanceTypes = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetDestinationRouteTableIds(v *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationRouteTableIds) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.DestinationRouteTableIds = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetSourceInstanceIds(v *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceInstanceIds) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.SourceInstanceIds = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetDestinationCidrBlocks(v *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationCidrBlocks) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.DestinationCidrBlocks = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetSourceRouteTableIds(v *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceRouteTableIds) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.SourceRouteTableIds = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetMatchCommunitySet(v *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapMatchCommunitySet) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.MatchCommunitySet = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetPrependAsPath(v *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapPrependAsPath) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.PrependAsPath = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetRouteTypes(v *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapRouteTypes) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.RouteTypes = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetDestinationChildInstanceTypes(v *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationChildInstanceTypes) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.DestinationChildInstanceTypes = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetDestinationInstanceIds(v *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationInstanceIds) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.DestinationInstanceIds = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetMatchAsns(v *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapMatchAsns) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.MatchAsns = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetOperateCommunitySet(v *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapOperateCommunitySet) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.OperateCommunitySet = v
	return s
}

type DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceRegionIds struct {
	SourceRegionId []*string `json:"SourceRegionId,omitempty" xml:"SourceRegionId,omitempty" type:"Repeated"`
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceRegionIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceRegionIds) GoString() string {
	return s.String()
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceRegionIds) SetSourceRegionId(v []*string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceRegionIds {
	s.SourceRegionId = v
	return s
}

type DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceChildInstanceTypes struct {
	SourceChildInstanceType []*string `json:"SourceChildInstanceType,omitempty" xml:"SourceChildInstanceType,omitempty" type:"Repeated"`
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceChildInstanceTypes) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceChildInstanceTypes) GoString() string {
	return s.String()
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceChildInstanceTypes) SetSourceChildInstanceType(v []*string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceChildInstanceTypes {
	s.SourceChildInstanceType = v
	return s
}

type DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationRouteTableIds struct {
	DestinationRouteTableId []*string `json:"DestinationRouteTableId,omitempty" xml:"DestinationRouteTableId,omitempty" type:"Repeated"`
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationRouteTableIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationRouteTableIds) GoString() string {
	return s.String()
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationRouteTableIds) SetDestinationRouteTableId(v []*string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationRouteTableIds {
	s.DestinationRouteTableId = v
	return s
}

type DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceInstanceIds struct {
	SourceInstanceId []*string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty" type:"Repeated"`
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceInstanceIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceInstanceIds) GoString() string {
	return s.String()
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceInstanceIds) SetSourceInstanceId(v []*string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceInstanceIds {
	s.SourceInstanceId = v
	return s
}

type DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationCidrBlocks struct {
	DestinationCidrBlock []*string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty" type:"Repeated"`
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationCidrBlocks) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationCidrBlocks) GoString() string {
	return s.String()
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationCidrBlocks) SetDestinationCidrBlock(v []*string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationCidrBlocks {
	s.DestinationCidrBlock = v
	return s
}

type DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceRouteTableIds struct {
	SourceRouteTableId []*string `json:"SourceRouteTableId,omitempty" xml:"SourceRouteTableId,omitempty" type:"Repeated"`
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceRouteTableIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceRouteTableIds) GoString() string {
	return s.String()
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceRouteTableIds) SetSourceRouteTableId(v []*string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceRouteTableIds {
	s.SourceRouteTableId = v
	return s
}

type DescribeCenRouteMapsResponseBodyRouteMapsRouteMapMatchCommunitySet struct {
	MatchCommunity []*string `json:"MatchCommunity,omitempty" xml:"MatchCommunity,omitempty" type:"Repeated"`
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapMatchCommunitySet) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapMatchCommunitySet) GoString() string {
	return s.String()
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapMatchCommunitySet) SetMatchCommunity(v []*string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapMatchCommunitySet {
	s.MatchCommunity = v
	return s
}

type DescribeCenRouteMapsResponseBodyRouteMapsRouteMapPrependAsPath struct {
	AsPath []*string `json:"AsPath,omitempty" xml:"AsPath,omitempty" type:"Repeated"`
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapPrependAsPath) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapPrependAsPath) GoString() string {
	return s.String()
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapPrependAsPath) SetAsPath(v []*string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapPrependAsPath {
	s.AsPath = v
	return s
}

type DescribeCenRouteMapsResponseBodyRouteMapsRouteMapRouteTypes struct {
	RouteType []*string `json:"RouteType,omitempty" xml:"RouteType,omitempty" type:"Repeated"`
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapRouteTypes) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapRouteTypes) GoString() string {
	return s.String()
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapRouteTypes) SetRouteType(v []*string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapRouteTypes {
	s.RouteType = v
	return s
}

type DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationChildInstanceTypes struct {
	DestinationChildInstanceType []*string `json:"DestinationChildInstanceType,omitempty" xml:"DestinationChildInstanceType,omitempty" type:"Repeated"`
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationChildInstanceTypes) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationChildInstanceTypes) GoString() string {
	return s.String()
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationChildInstanceTypes) SetDestinationChildInstanceType(v []*string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationChildInstanceTypes {
	s.DestinationChildInstanceType = v
	return s
}

type DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationInstanceIds struct {
	DestinationInstanceId []*string `json:"DestinationInstanceId,omitempty" xml:"DestinationInstanceId,omitempty" type:"Repeated"`
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationInstanceIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationInstanceIds) GoString() string {
	return s.String()
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationInstanceIds) SetDestinationInstanceId(v []*string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationInstanceIds {
	s.DestinationInstanceId = v
	return s
}

type DescribeCenRouteMapsResponseBodyRouteMapsRouteMapMatchAsns struct {
	MatchAsn []*string `json:"MatchAsn,omitempty" xml:"MatchAsn,omitempty" type:"Repeated"`
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapMatchAsns) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapMatchAsns) GoString() string {
	return s.String()
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapMatchAsns) SetMatchAsn(v []*string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapMatchAsns {
	s.MatchAsn = v
	return s
}

type DescribeCenRouteMapsResponseBodyRouteMapsRouteMapOperateCommunitySet struct {
	OperateCommunity []*string `json:"OperateCommunity,omitempty" xml:"OperateCommunity,omitempty" type:"Repeated"`
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapOperateCommunitySet) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapOperateCommunitySet) GoString() string {
	return s.String()
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapOperateCommunitySet) SetOperateCommunity(v []*string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapOperateCommunitySet {
	s.OperateCommunity = v
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
	PageSize   *int32                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount *int32                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Cens       *DescribeCensResponseBodyCens `json:"Cens,omitempty" xml:"Cens,omitempty" type:"Struct"`
}

func (s DescribeCensResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCensResponseBody) GoString() string {
	return s.String()
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

func (s *DescribeCensResponseBody) SetTotalCount(v int32) *DescribeCensResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCensResponseBody) SetCens(v *DescribeCensResponseBodyCens) *DescribeCensResponseBody {
	s.Cens = v
	return s
}

type DescribeCensResponseBodyCens struct {
	Cen []*DescribeCensResponseBodyCensCen `json:"Cen,omitempty" xml:"Cen,omitempty" type:"Repeated"`
}

func (s DescribeCensResponseBodyCens) String() string {
	return tea.Prettify(s)
}

func (s DescribeCensResponseBodyCens) GoString() string {
	return s.String()
}

func (s *DescribeCensResponseBodyCens) SetCen(v []*DescribeCensResponseBodyCensCen) *DescribeCensResponseBodyCens {
	s.Cen = v
	return s
}

type DescribeCensResponseBodyCensCen struct {
	Status                 *string                                                `json:"Status,omitempty" xml:"Status,omitempty"`
	CreationTime           *string                                                `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description            *string                                                `json:"Description,omitempty" xml:"Description,omitempty"`
	CenId                  *string                                                `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ProtectionLevel        *string                                                `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty"`
	Name                   *string                                                `json:"Name,omitempty" xml:"Name,omitempty"`
	Tags                   *DescribeCensResponseBodyCensCenTags                   `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	CenBandwidthPackageIds *DescribeCensResponseBodyCensCenCenBandwidthPackageIds `json:"CenBandwidthPackageIds,omitempty" xml:"CenBandwidthPackageIds,omitempty" type:"Struct"`
}

func (s DescribeCensResponseBodyCensCen) String() string {
	return tea.Prettify(s)
}

func (s DescribeCensResponseBodyCensCen) GoString() string {
	return s.String()
}

func (s *DescribeCensResponseBodyCensCen) SetStatus(v string) *DescribeCensResponseBodyCensCen {
	s.Status = &v
	return s
}

func (s *DescribeCensResponseBodyCensCen) SetCreationTime(v string) *DescribeCensResponseBodyCensCen {
	s.CreationTime = &v
	return s
}

func (s *DescribeCensResponseBodyCensCen) SetDescription(v string) *DescribeCensResponseBodyCensCen {
	s.Description = &v
	return s
}

func (s *DescribeCensResponseBodyCensCen) SetCenId(v string) *DescribeCensResponseBodyCensCen {
	s.CenId = &v
	return s
}

func (s *DescribeCensResponseBodyCensCen) SetProtectionLevel(v string) *DescribeCensResponseBodyCensCen {
	s.ProtectionLevel = &v
	return s
}

func (s *DescribeCensResponseBodyCensCen) SetName(v string) *DescribeCensResponseBodyCensCen {
	s.Name = &v
	return s
}

func (s *DescribeCensResponseBodyCensCen) SetTags(v *DescribeCensResponseBodyCensCenTags) *DescribeCensResponseBodyCensCen {
	s.Tags = v
	return s
}

func (s *DescribeCensResponseBodyCensCen) SetCenBandwidthPackageIds(v *DescribeCensResponseBodyCensCenCenBandwidthPackageIds) *DescribeCensResponseBodyCensCen {
	s.CenBandwidthPackageIds = v
	return s
}

type DescribeCensResponseBodyCensCenTags struct {
	Tag []*DescribeCensResponseBodyCensCenTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeCensResponseBodyCensCenTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeCensResponseBodyCensCenTags) GoString() string {
	return s.String()
}

func (s *DescribeCensResponseBodyCensCenTags) SetTag(v []*DescribeCensResponseBodyCensCenTagsTag) *DescribeCensResponseBodyCensCenTags {
	s.Tag = v
	return s
}

type DescribeCensResponseBodyCensCenTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeCensResponseBodyCensCenTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeCensResponseBodyCensCenTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeCensResponseBodyCensCenTagsTag) SetKey(v string) *DescribeCensResponseBodyCensCenTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeCensResponseBodyCensCenTagsTag) SetValue(v string) *DescribeCensResponseBodyCensCenTagsTag {
	s.Value = &v
	return s
}

type DescribeCensResponseBodyCensCenCenBandwidthPackageIds struct {
	CenBandwidthPackageId []*string `json:"CenBandwidthPackageId,omitempty" xml:"CenBandwidthPackageId,omitempty" type:"Repeated"`
}

func (s DescribeCensResponseBodyCensCenCenBandwidthPackageIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeCensResponseBodyCensCenCenBandwidthPackageIds) GoString() string {
	return s.String()
}

func (s *DescribeCensResponseBodyCensCenCenBandwidthPackageIds) SetCenBandwidthPackageId(v []*string) *DescribeCensResponseBodyCensCenCenBandwidthPackageIds {
	s.CenBandwidthPackageId = v
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
	PageSize        *int32                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId       *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber      *int32                                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount      *int32                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VbrHealthChecks *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecks `json:"VbrHealthChecks,omitempty" xml:"VbrHealthChecks,omitempty" type:"Struct"`
}

func (s DescribeCenVbrHealthCheckResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenVbrHealthCheckResponseBody) GoString() string {
	return s.String()
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

func (s *DescribeCenVbrHealthCheckResponseBody) SetTotalCount(v int32) *DescribeCenVbrHealthCheckResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCenVbrHealthCheckResponseBody) SetVbrHealthChecks(v *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecks) *DescribeCenVbrHealthCheckResponseBody {
	s.VbrHealthChecks = v
	return s
}

type DescribeCenVbrHealthCheckResponseBodyVbrHealthChecks struct {
	VbrHealthCheck []*DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck `json:"VbrHealthCheck,omitempty" xml:"VbrHealthCheck,omitempty" type:"Repeated"`
}

func (s DescribeCenVbrHealthCheckResponseBodyVbrHealthChecks) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenVbrHealthCheckResponseBodyVbrHealthChecks) GoString() string {
	return s.String()
}

func (s *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecks) SetVbrHealthCheck(v []*DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck) *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecks {
	s.VbrHealthCheck = v
	return s
}

type DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck struct {
	HealthCheckTargetIp *string `json:"HealthCheckTargetIp,omitempty" xml:"HealthCheckTargetIp,omitempty"`
	VbrInstanceId       *string `json:"VbrInstanceId,omitempty" xml:"VbrInstanceId,omitempty"`
	VbrInstanceRegionId *string `json:"VbrInstanceRegionId,omitempty" xml:"VbrInstanceRegionId,omitempty"`
	CenId               *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	HealthyThreshold    *int32  `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	HealthCheckInterval *int32  `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	HealthCheckSourceIp *string `json:"HealthCheckSourceIp,omitempty" xml:"HealthCheckSourceIp,omitempty"`
}

func (s DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck) GoString() string {
	return s.String()
}

func (s *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck) SetHealthCheckTargetIp(v string) *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck {
	s.HealthCheckTargetIp = &v
	return s
}

func (s *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck) SetVbrInstanceId(v string) *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck {
	s.VbrInstanceId = &v
	return s
}

func (s *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck) SetVbrInstanceRegionId(v string) *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck {
	s.VbrInstanceRegionId = &v
	return s
}

func (s *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck) SetCenId(v string) *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck {
	s.CenId = &v
	return s
}

func (s *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck) SetHealthyThreshold(v int32) *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck {
	s.HealthyThreshold = &v
	return s
}

func (s *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck) SetHealthCheckInterval(v int32) *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck {
	s.HealthCheckInterval = &v
	return s
}

func (s *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck) SetHealthCheckSourceIp(v string) *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck {
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

type DescribeChildInstanceRegionsResponseBody struct {
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Regions   *DescribeChildInstanceRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
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

func (s *DescribeChildInstanceRegionsResponseBody) SetRegions(v *DescribeChildInstanceRegionsResponseBodyRegions) *DescribeChildInstanceRegionsResponseBody {
	s.Regions = v
	return s
}

type DescribeChildInstanceRegionsResponseBodyRegions struct {
	Region []*DescribeChildInstanceRegionsResponseBodyRegionsRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s DescribeChildInstanceRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeChildInstanceRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeChildInstanceRegionsResponseBodyRegions) SetRegion(v []*DescribeChildInstanceRegionsResponseBodyRegionsRegion) *DescribeChildInstanceRegionsResponseBodyRegions {
	s.Region = v
	return s
}

type DescribeChildInstanceRegionsResponseBodyRegionsRegion struct {
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeChildInstanceRegionsResponseBodyRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeChildInstanceRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeChildInstanceRegionsResponseBodyRegionsRegion) SetLocalName(v string) *DescribeChildInstanceRegionsResponseBodyRegionsRegion {
	s.LocalName = &v
	return s
}

func (s *DescribeChildInstanceRegionsResponseBodyRegionsRegion) SetRegionId(v string) *DescribeChildInstanceRegionsResponseBodyRegionsRegion {
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
	PageSize   *string                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *string                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *string                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Success    *string                               `json:"Success,omitempty" xml:"Success,omitempty"`
	FlowLogs   *DescribeFlowlogsResponseBodyFlowLogs `json:"FlowLogs,omitempty" xml:"FlowLogs,omitempty" type:"Struct"`
}

func (s DescribeFlowlogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowlogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowlogsResponseBody) SetPageSize(v string) *DescribeFlowlogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeFlowlogsResponseBody) SetPageNumber(v string) *DescribeFlowlogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeFlowlogsResponseBody) SetRequestId(v string) *DescribeFlowlogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlowlogsResponseBody) SetTotalCount(v string) *DescribeFlowlogsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeFlowlogsResponseBody) SetSuccess(v string) *DescribeFlowlogsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeFlowlogsResponseBody) SetFlowLogs(v *DescribeFlowlogsResponseBodyFlowLogs) *DescribeFlowlogsResponseBody {
	s.FlowLogs = v
	return s
}

type DescribeFlowlogsResponseBodyFlowLogs struct {
	FlowLog []*DescribeFlowlogsResponseBodyFlowLogsFlowLog `json:"FlowLog,omitempty" xml:"FlowLog,omitempty" type:"Repeated"`
}

func (s DescribeFlowlogsResponseBodyFlowLogs) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowlogsResponseBodyFlowLogs) GoString() string {
	return s.String()
}

func (s *DescribeFlowlogsResponseBodyFlowLogs) SetFlowLog(v []*DescribeFlowlogsResponseBodyFlowLogsFlowLog) *DescribeFlowlogsResponseBodyFlowLogs {
	s.FlowLog = v
	return s
}

type DescribeFlowlogsResponseBodyFlowLogsFlowLog struct {
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

func (s DescribeFlowlogsResponseBodyFlowLogsFlowLog) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowlogsResponseBodyFlowLogsFlowLog) GoString() string {
	return s.String()
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) SetStatus(v string) *DescribeFlowlogsResponseBodyFlowLogsFlowLog {
	s.Status = &v
	return s
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) SetCreationTime(v string) *DescribeFlowlogsResponseBodyFlowLogsFlowLog {
	s.CreationTime = &v
	return s
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) SetFlowLogName(v string) *DescribeFlowlogsResponseBodyFlowLogsFlowLog {
	s.FlowLogName = &v
	return s
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) SetDescription(v string) *DescribeFlowlogsResponseBodyFlowLogsFlowLog {
	s.Description = &v
	return s
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) SetProjectName(v string) *DescribeFlowlogsResponseBodyFlowLogsFlowLog {
	s.ProjectName = &v
	return s
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) SetCenId(v string) *DescribeFlowlogsResponseBodyFlowLogsFlowLog {
	s.CenId = &v
	return s
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) SetLogStoreName(v string) *DescribeFlowlogsResponseBodyFlowLogsFlowLog {
	s.LogStoreName = &v
	return s
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) SetRegionId(v string) *DescribeFlowlogsResponseBodyFlowLogsFlowLog {
	s.RegionId = &v
	return s
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) SetFlowLogId(v string) *DescribeFlowlogsResponseBodyFlowLogsFlowLog {
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
	PageSize   *int32                                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount *int32                                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RegionIds  *DescribeGeographicRegionMembershipResponseBodyRegionIds `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Struct"`
}

func (s DescribeGeographicRegionMembershipResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGeographicRegionMembershipResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGeographicRegionMembershipResponseBody) SetPageSize(v int32) *DescribeGeographicRegionMembershipResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGeographicRegionMembershipResponseBody) SetRequestId(v string) *DescribeGeographicRegionMembershipResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGeographicRegionMembershipResponseBody) SetPageNumber(v int32) *DescribeGeographicRegionMembershipResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeGeographicRegionMembershipResponseBody) SetTotalCount(v int32) *DescribeGeographicRegionMembershipResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeGeographicRegionMembershipResponseBody) SetRegionIds(v *DescribeGeographicRegionMembershipResponseBodyRegionIds) *DescribeGeographicRegionMembershipResponseBody {
	s.RegionIds = v
	return s
}

type DescribeGeographicRegionMembershipResponseBodyRegionIds struct {
	RegionId []*DescribeGeographicRegionMembershipResponseBodyRegionIdsRegionId `json:"RegionId,omitempty" xml:"RegionId,omitempty" type:"Repeated"`
}

func (s DescribeGeographicRegionMembershipResponseBodyRegionIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeGeographicRegionMembershipResponseBodyRegionIds) GoString() string {
	return s.String()
}

func (s *DescribeGeographicRegionMembershipResponseBodyRegionIds) SetRegionId(v []*DescribeGeographicRegionMembershipResponseBodyRegionIdsRegionId) *DescribeGeographicRegionMembershipResponseBodyRegionIds {
	s.RegionId = v
	return s
}

type DescribeGeographicRegionMembershipResponseBodyRegionIdsRegionId struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeGeographicRegionMembershipResponseBodyRegionIdsRegionId) String() string {
	return tea.Prettify(s)
}

func (s DescribeGeographicRegionMembershipResponseBodyRegionIdsRegionId) GoString() string {
	return s.String()
}

func (s *DescribeGeographicRegionMembershipResponseBodyRegionIdsRegionId) SetRegionId(v string) *DescribeGeographicRegionMembershipResponseBodyRegionIdsRegionId {
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
	RequestId  *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	GrantRules *DescribeGrantRulesToCenResponseBodyGrantRules `json:"GrantRules,omitempty" xml:"GrantRules,omitempty" type:"Struct"`
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

func (s *DescribeGrantRulesToCenResponseBody) SetGrantRules(v *DescribeGrantRulesToCenResponseBodyGrantRules) *DescribeGrantRulesToCenResponseBody {
	s.GrantRules = v
	return s
}

type DescribeGrantRulesToCenResponseBodyGrantRules struct {
	GrantRule []*DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule `json:"GrantRule,omitempty" xml:"GrantRule,omitempty" type:"Repeated"`
}

func (s DescribeGrantRulesToCenResponseBodyGrantRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeGrantRulesToCenResponseBodyGrantRules) GoString() string {
	return s.String()
}

func (s *DescribeGrantRulesToCenResponseBodyGrantRules) SetGrantRule(v []*DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule) *DescribeGrantRulesToCenResponseBodyGrantRules {
	s.GrantRule = v
	return s
}

type DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule struct {
	ChildInstanceType     *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	ChildInstanceRegionId *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	ChildInstanceOwnerId  *int64  `json:"ChildInstanceOwnerId,omitempty" xml:"ChildInstanceOwnerId,omitempty"`
	ChildInstanceId       *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	CenId                 *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
}

func (s DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule) String() string {
	return tea.Prettify(s)
}

func (s DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule) GoString() string {
	return s.String()
}

func (s *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule) SetChildInstanceType(v string) *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule {
	s.ChildInstanceType = &v
	return s
}

func (s *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule) SetChildInstanceRegionId(v string) *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule) SetChildInstanceOwnerId(v int64) *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule {
	s.ChildInstanceOwnerId = &v
	return s
}

func (s *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule) SetChildInstanceId(v string) *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule {
	s.ChildInstanceId = &v
	return s
}

func (s *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule) SetCenId(v string) *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule {
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
	PageSize              *int32                                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId             *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber            *int32                                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount            *int32                                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PublishedRouteEntries *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntries `json:"PublishedRouteEntries,omitempty" xml:"PublishedRouteEntries,omitempty" type:"Struct"`
}

func (s DescribePublishedRouteEntriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePublishedRouteEntriesResponseBody) GoString() string {
	return s.String()
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

func (s *DescribePublishedRouteEntriesResponseBody) SetTotalCount(v int32) *DescribePublishedRouteEntriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBody) SetPublishedRouteEntries(v *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntries) *DescribePublishedRouteEntriesResponseBody {
	s.PublishedRouteEntries = v
	return s
}

type DescribePublishedRouteEntriesResponseBodyPublishedRouteEntries struct {
	PublishedRouteEntry []*DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry `json:"PublishedRouteEntry,omitempty" xml:"PublishedRouteEntry,omitempty" type:"Repeated"`
}

func (s DescribePublishedRouteEntriesResponseBodyPublishedRouteEntries) String() string {
	return tea.Prettify(s)
}

func (s DescribePublishedRouteEntriesResponseBodyPublishedRouteEntries) GoString() string {
	return s.String()
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntries) SetPublishedRouteEntry(v []*DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntries {
	s.PublishedRouteEntry = v
	return s
}

type DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry struct {
	NextHopId                 *string                                                                                     `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	PublishStatus             *string                                                                                     `json:"PublishStatus,omitempty" xml:"PublishStatus,omitempty"`
	ChildInstanceRouteTableId *string                                                                                     `json:"ChildInstanceRouteTableId,omitempty" xml:"ChildInstanceRouteTableId,omitempty"`
	NextHopType               *string                                                                                     `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	OperationalMode           *bool                                                                                       `json:"OperationalMode,omitempty" xml:"OperationalMode,omitempty"`
	DestinationCidrBlock      *string                                                                                     `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	RouteType                 *string                                                                                     `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
	Conflicts                 *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflicts `json:"Conflicts,omitempty" xml:"Conflicts,omitempty" type:"Struct"`
}

func (s DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry) String() string {
	return tea.Prettify(s)
}

func (s DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry) GoString() string {
	return s.String()
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry) SetNextHopId(v string) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry {
	s.NextHopId = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry) SetPublishStatus(v string) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry {
	s.PublishStatus = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry) SetChildInstanceRouteTableId(v string) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry {
	s.ChildInstanceRouteTableId = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry) SetNextHopType(v string) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry {
	s.NextHopType = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry) SetOperationalMode(v bool) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry {
	s.OperationalMode = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry) SetDestinationCidrBlock(v string) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry) SetRouteType(v string) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry {
	s.RouteType = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry) SetConflicts(v *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflicts) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry {
	s.Conflicts = v
	return s
}

type DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflicts struct {
	Conflict []*DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflictsConflict `json:"Conflict,omitempty" xml:"Conflict,omitempty" type:"Repeated"`
}

func (s DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflicts) String() string {
	return tea.Prettify(s)
}

func (s DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflicts) GoString() string {
	return s.String()
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflicts) SetConflict(v []*DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflictsConflict) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflicts {
	s.Conflict = v
	return s
}

type DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflictsConflict struct {
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType         *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflictsConflict) String() string {
	return tea.Prettify(s)
}

func (s DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflictsConflict) GoString() string {
	return s.String()
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflictsConflict) SetStatus(v string) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflictsConflict {
	s.Status = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflictsConflict) SetDestinationCidrBlock(v string) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflictsConflict {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflictsConflict) SetInstanceId(v string) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflictsConflict {
	s.InstanceId = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflictsConflict) SetInstanceType(v string) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflictsConflict {
	s.InstanceType = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflictsConflict) SetRegionId(v string) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflictsConflict {
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
	PageSize       *int32                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber     *int32                                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount     *int32                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RouteConflicts *DescribeRouteConflictResponseBodyRouteConflicts `json:"RouteConflicts,omitempty" xml:"RouteConflicts,omitempty" type:"Struct"`
}

func (s DescribeRouteConflictResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRouteConflictResponseBody) GoString() string {
	return s.String()
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

func (s *DescribeRouteConflictResponseBody) SetTotalCount(v int32) *DescribeRouteConflictResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRouteConflictResponseBody) SetRouteConflicts(v *DescribeRouteConflictResponseBodyRouteConflicts) *DescribeRouteConflictResponseBody {
	s.RouteConflicts = v
	return s
}

type DescribeRouteConflictResponseBodyRouteConflicts struct {
	RouteConflict []*DescribeRouteConflictResponseBodyRouteConflictsRouteConflict `json:"RouteConflict,omitempty" xml:"RouteConflict,omitempty" type:"Repeated"`
}

func (s DescribeRouteConflictResponseBodyRouteConflicts) String() string {
	return tea.Prettify(s)
}

func (s DescribeRouteConflictResponseBodyRouteConflicts) GoString() string {
	return s.String()
}

func (s *DescribeRouteConflictResponseBodyRouteConflicts) SetRouteConflict(v []*DescribeRouteConflictResponseBodyRouteConflictsRouteConflict) *DescribeRouteConflictResponseBodyRouteConflicts {
	s.RouteConflict = v
	return s
}

type DescribeRouteConflictResponseBodyRouteConflictsRouteConflict struct {
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType         *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRouteConflictResponseBodyRouteConflictsRouteConflict) String() string {
	return tea.Prettify(s)
}

func (s DescribeRouteConflictResponseBodyRouteConflictsRouteConflict) GoString() string {
	return s.String()
}

func (s *DescribeRouteConflictResponseBodyRouteConflictsRouteConflict) SetStatus(v string) *DescribeRouteConflictResponseBodyRouteConflictsRouteConflict {
	s.Status = &v
	return s
}

func (s *DescribeRouteConflictResponseBodyRouteConflictsRouteConflict) SetDestinationCidrBlock(v string) *DescribeRouteConflictResponseBodyRouteConflictsRouteConflict {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DescribeRouteConflictResponseBodyRouteConflictsRouteConflict) SetInstanceId(v string) *DescribeRouteConflictResponseBodyRouteConflictsRouteConflict {
	s.InstanceId = &v
	return s
}

func (s *DescribeRouteConflictResponseBodyRouteConflictsRouteConflict) SetInstanceType(v string) *DescribeRouteConflictResponseBodyRouteConflictsRouteConflict {
	s.InstanceType = &v
	return s
}

func (s *DescribeRouteConflictResponseBodyRouteConflictsRouteConflict) SetRegionId(v string) *DescribeRouteConflictResponseBodyRouteConflictsRouteConflict {
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
	PageSize            *int32                                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId           *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber          *int32                                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount          *int32                                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RouteServiceEntries *DescribeRouteServicesInCenResponseBodyRouteServiceEntries `json:"RouteServiceEntries,omitempty" xml:"RouteServiceEntries,omitempty" type:"Struct"`
}

func (s DescribeRouteServicesInCenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRouteServicesInCenResponseBody) GoString() string {
	return s.String()
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

func (s *DescribeRouteServicesInCenResponseBody) SetTotalCount(v int32) *DescribeRouteServicesInCenResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRouteServicesInCenResponseBody) SetRouteServiceEntries(v *DescribeRouteServicesInCenResponseBodyRouteServiceEntries) *DescribeRouteServicesInCenResponseBody {
	s.RouteServiceEntries = v
	return s
}

type DescribeRouteServicesInCenResponseBodyRouteServiceEntries struct {
	RouteServiceEntry []*DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry `json:"RouteServiceEntry,omitempty" xml:"RouteServiceEntry,omitempty" type:"Repeated"`
}

func (s DescribeRouteServicesInCenResponseBodyRouteServiceEntries) String() string {
	return tea.Prettify(s)
}

func (s DescribeRouteServicesInCenResponseBodyRouteServiceEntries) GoString() string {
	return s.String()
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntries) SetRouteServiceEntry(v []*DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry) *DescribeRouteServicesInCenResponseBodyRouteServiceEntries {
	s.RouteServiceEntry = v
	return s
}

type DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry struct {
	Status         *string                                                                          `json:"Status,omitempty" xml:"Status,omitempty"`
	Host           *string                                                                          `json:"Host,omitempty" xml:"Host,omitempty"`
	Description    *string                                                                          `json:"Description,omitempty" xml:"Description,omitempty"`
	HostVpcId      *string                                                                          `json:"HostVpcId,omitempty" xml:"HostVpcId,omitempty"`
	CenId          *string                                                                          `json:"CenId,omitempty" xml:"CenId,omitempty"`
	AccessRegionId *string                                                                          `json:"AccessRegionId,omitempty" xml:"AccessRegionId,omitempty"`
	HostRegionId   *string                                                                          `json:"HostRegionId,omitempty" xml:"HostRegionId,omitempty"`
	UpdateInterval *string                                                                          `json:"UpdateInterval,omitempty" xml:"UpdateInterval,omitempty"`
	Cidrs          *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntryCidrs `json:"Cidrs,omitempty" xml:"Cidrs,omitempty" type:"Struct"`
}

func (s DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry) String() string {
	return tea.Prettify(s)
}

func (s DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry) GoString() string {
	return s.String()
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry) SetStatus(v string) *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry {
	s.Status = &v
	return s
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry) SetHost(v string) *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry {
	s.Host = &v
	return s
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry) SetDescription(v string) *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry {
	s.Description = &v
	return s
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry) SetHostVpcId(v string) *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry {
	s.HostVpcId = &v
	return s
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry) SetCenId(v string) *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry {
	s.CenId = &v
	return s
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry) SetAccessRegionId(v string) *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry {
	s.AccessRegionId = &v
	return s
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry) SetHostRegionId(v string) *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry {
	s.HostRegionId = &v
	return s
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry) SetUpdateInterval(v string) *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry {
	s.UpdateInterval = &v
	return s
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry) SetCidrs(v *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntryCidrs) *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry {
	s.Cidrs = v
	return s
}

type DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntryCidrs struct {
	Cidr []*string `json:"Cidr,omitempty" xml:"Cidr,omitempty" type:"Repeated"`
}

func (s DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntryCidrs) String() string {
	return tea.Prettify(s)
}

func (s DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntryCidrs) GoString() string {
	return s.String()
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntryCidrs) SetCidr(v []*string) *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntryCidrs {
	s.Cidr = v
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

type ModifyCenAttributeRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ProtectionLevel      *string `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty"`
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
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyFlowLogAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowLogAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFlowLogAttributeResponseBody) SetSuccess(v string) *ModifyFlowLogAttributeResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyFlowLogAttributeResponseBody) SetRequestId(v string) *ModifyFlowLogAttributeResponseBody {
	s.RequestId = &v
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
	client.EndpointRule = tea.String("central")
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
