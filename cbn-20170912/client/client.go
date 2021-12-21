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

type ActiveFlowLogRequest struct {
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	FlowLogId            *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ActiveFlowLogRequest) String() string {
	return tea.Prettify(s)
}

func (s ActiveFlowLogRequest) GoString() string {
	return s.String()
}

func (s *ActiveFlowLogRequest) SetCenId(v string) *ActiveFlowLogRequest {
	s.CenId = &v
	return s
}

func (s *ActiveFlowLogRequest) SetClientToken(v string) *ActiveFlowLogRequest {
	s.ClientToken = &v
	return s
}

func (s *ActiveFlowLogRequest) SetFlowLogId(v string) *ActiveFlowLogRequest {
	s.FlowLogId = &v
	return s
}

func (s *ActiveFlowLogRequest) SetOwnerAccount(v string) *ActiveFlowLogRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ActiveFlowLogRequest) SetOwnerId(v int64) *ActiveFlowLogRequest {
	s.OwnerId = &v
	return s
}

func (s *ActiveFlowLogRequest) SetRegionId(v string) *ActiveFlowLogRequest {
	s.RegionId = &v
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

type AddTraficMatchRuleToTrafficMarkingPolicyRequest struct {
	ClientToken            *string                                                             `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                 *bool                                                               `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount           *string                                                             `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                *int64                                                              `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount   *string                                                             `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId        *int64                                                              `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TrafficMarkingPolicyId *string                                                             `json:"TrafficMarkingPolicyId,omitempty" xml:"TrafficMarkingPolicyId,omitempty"`
	TrafficMatchRules      []*AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules `json:"TrafficMatchRules,omitempty" xml:"TrafficMatchRules,omitempty" type:"Repeated"`
}

func (s AddTraficMatchRuleToTrafficMarkingPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s AddTraficMatchRuleToTrafficMarkingPolicyRequest) GoString() string {
	return s.String()
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequest) SetClientToken(v string) *AddTraficMatchRuleToTrafficMarkingPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequest) SetDryRun(v bool) *AddTraficMatchRuleToTrafficMarkingPolicyRequest {
	s.DryRun = &v
	return s
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequest) SetOwnerAccount(v string) *AddTraficMatchRuleToTrafficMarkingPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequest) SetOwnerId(v int64) *AddTraficMatchRuleToTrafficMarkingPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequest) SetResourceOwnerAccount(v string) *AddTraficMatchRuleToTrafficMarkingPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequest) SetResourceOwnerId(v int64) *AddTraficMatchRuleToTrafficMarkingPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequest) SetTrafficMarkingPolicyId(v string) *AddTraficMatchRuleToTrafficMarkingPolicyRequest {
	s.TrafficMarkingPolicyId = &v
	return s
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequest) SetTrafficMatchRules(v []*AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) *AddTraficMatchRuleToTrafficMarkingPolicyRequest {
	s.TrafficMatchRules = v
	return s
}

type AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules struct {
	DstCidr                     *string  `json:"DstCidr,omitempty" xml:"DstCidr,omitempty"`
	DstPortRange                []*int32 `json:"DstPortRange,omitempty" xml:"DstPortRange,omitempty" type:"Repeated"`
	MatchDscp                   *int32   `json:"MatchDscp,omitempty" xml:"MatchDscp,omitempty"`
	Protocol                    *string  `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	SrcCidr                     *string  `json:"SrcCidr,omitempty" xml:"SrcCidr,omitempty"`
	SrcPortRange                []*int32 `json:"SrcPortRange,omitempty" xml:"SrcPortRange,omitempty" type:"Repeated"`
	TrafficMatchRuleDescription *string  `json:"TrafficMatchRuleDescription,omitempty" xml:"TrafficMatchRuleDescription,omitempty"`
	TrafficMatchRuleName        *string  `json:"TrafficMatchRuleName,omitempty" xml:"TrafficMatchRuleName,omitempty"`
}

func (s AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) String() string {
	return tea.Prettify(s)
}

func (s AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) GoString() string {
	return s.String()
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) SetDstCidr(v string) *AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules {
	s.DstCidr = &v
	return s
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) SetDstPortRange(v []*int32) *AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules {
	s.DstPortRange = v
	return s
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) SetMatchDscp(v int32) *AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules {
	s.MatchDscp = &v
	return s
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) SetProtocol(v string) *AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules {
	s.Protocol = &v
	return s
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) SetSrcCidr(v string) *AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules {
	s.SrcCidr = &v
	return s
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) SetSrcPortRange(v []*int32) *AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules {
	s.SrcPortRange = v
	return s
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) SetTrafficMatchRuleDescription(v string) *AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules {
	s.TrafficMatchRuleDescription = &v
	return s
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules) SetTrafficMatchRuleName(v string) *AddTraficMatchRuleToTrafficMarkingPolicyRequestTrafficMatchRules {
	s.TrafficMatchRuleName = &v
	return s
}

type AddTraficMatchRuleToTrafficMarkingPolicyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddTraficMatchRuleToTrafficMarkingPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddTraficMatchRuleToTrafficMarkingPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyResponseBody) SetRequestId(v string) *AddTraficMatchRuleToTrafficMarkingPolicyResponseBody {
	s.RequestId = &v
	return s
}

type AddTraficMatchRuleToTrafficMarkingPolicyResponse struct {
	Headers map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddTraficMatchRuleToTrafficMarkingPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddTraficMatchRuleToTrafficMarkingPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s AddTraficMatchRuleToTrafficMarkingPolicyResponse) GoString() string {
	return s.String()
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyResponse) SetHeaders(v map[string]*string) *AddTraficMatchRuleToTrafficMarkingPolicyResponse {
	s.Headers = v
	return s
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyResponse) SetBody(v *AddTraficMatchRuleToTrafficMarkingPolicyResponseBody) *AddTraficMatchRuleToTrafficMarkingPolicyResponse {
	s.Body = v
	return s
}

type AssociateCenBandwidthPackageRequest struct {
	CenBandwidthPackageId *string `json:"CenBandwidthPackageId,omitempty" xml:"CenBandwidthPackageId,omitempty"`
	CenId                 *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AssociateCenBandwidthPackageRequest) String() string {
	return tea.Prettify(s)
}

func (s AssociateCenBandwidthPackageRequest) GoString() string {
	return s.String()
}

func (s *AssociateCenBandwidthPackageRequest) SetCenBandwidthPackageId(v string) *AssociateCenBandwidthPackageRequest {
	s.CenBandwidthPackageId = &v
	return s
}

func (s *AssociateCenBandwidthPackageRequest) SetCenId(v string) *AssociateCenBandwidthPackageRequest {
	s.CenId = &v
	return s
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

type AssociateTransitRouterAttachmentWithRouteTableRequest struct {
	ClientToken               *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                    *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount              *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount      *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId           *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	TransitRouterRouteTableId *string `json:"TransitRouterRouteTableId,omitempty" xml:"TransitRouterRouteTableId,omitempty"`
}

func (s AssociateTransitRouterAttachmentWithRouteTableRequest) String() string {
	return tea.Prettify(s)
}

func (s AssociateTransitRouterAttachmentWithRouteTableRequest) GoString() string {
	return s.String()
}

func (s *AssociateTransitRouterAttachmentWithRouteTableRequest) SetClientToken(v string) *AssociateTransitRouterAttachmentWithRouteTableRequest {
	s.ClientToken = &v
	return s
}

func (s *AssociateTransitRouterAttachmentWithRouteTableRequest) SetDryRun(v bool) *AssociateTransitRouterAttachmentWithRouteTableRequest {
	s.DryRun = &v
	return s
}

func (s *AssociateTransitRouterAttachmentWithRouteTableRequest) SetOwnerAccount(v string) *AssociateTransitRouterAttachmentWithRouteTableRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AssociateTransitRouterAttachmentWithRouteTableRequest) SetOwnerId(v int64) *AssociateTransitRouterAttachmentWithRouteTableRequest {
	s.OwnerId = &v
	return s
}

func (s *AssociateTransitRouterAttachmentWithRouteTableRequest) SetResourceOwnerAccount(v string) *AssociateTransitRouterAttachmentWithRouteTableRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AssociateTransitRouterAttachmentWithRouteTableRequest) SetResourceOwnerId(v int64) *AssociateTransitRouterAttachmentWithRouteTableRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AssociateTransitRouterAttachmentWithRouteTableRequest) SetTransitRouterAttachmentId(v string) *AssociateTransitRouterAttachmentWithRouteTableRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *AssociateTransitRouterAttachmentWithRouteTableRequest) SetTransitRouterRouteTableId(v string) *AssociateTransitRouterAttachmentWithRouteTableRequest {
	s.TransitRouterRouteTableId = &v
	return s
}

type AssociateTransitRouterAttachmentWithRouteTableResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateTransitRouterAttachmentWithRouteTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssociateTransitRouterAttachmentWithRouteTableResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateTransitRouterAttachmentWithRouteTableResponseBody) SetRequestId(v string) *AssociateTransitRouterAttachmentWithRouteTableResponseBody {
	s.RequestId = &v
	return s
}

type AssociateTransitRouterAttachmentWithRouteTableResponse struct {
	Headers map[string]*string                                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AssociateTransitRouterAttachmentWithRouteTableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AssociateTransitRouterAttachmentWithRouteTableResponse) String() string {
	return tea.Prettify(s)
}

func (s AssociateTransitRouterAttachmentWithRouteTableResponse) GoString() string {
	return s.String()
}

func (s *AssociateTransitRouterAttachmentWithRouteTableResponse) SetHeaders(v map[string]*string) *AssociateTransitRouterAttachmentWithRouteTableResponse {
	s.Headers = v
	return s
}

func (s *AssociateTransitRouterAttachmentWithRouteTableResponse) SetBody(v *AssociateTransitRouterAttachmentWithRouteTableResponseBody) *AssociateTransitRouterAttachmentWithRouteTableResponse {
	s.Body = v
	return s
}

type AttachCenChildInstanceRequest struct {
	CenId                 *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ChildInstanceId       *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	ChildInstanceOwnerId  *int64  `json:"ChildInstanceOwnerId,omitempty" xml:"ChildInstanceOwnerId,omitempty"`
	ChildInstanceRegionId *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	ChildInstanceType     *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AttachCenChildInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachCenChildInstanceRequest) GoString() string {
	return s.String()
}

func (s *AttachCenChildInstanceRequest) SetCenId(v string) *AttachCenChildInstanceRequest {
	s.CenId = &v
	return s
}

func (s *AttachCenChildInstanceRequest) SetChildInstanceId(v string) *AttachCenChildInstanceRequest {
	s.ChildInstanceId = &v
	return s
}

func (s *AttachCenChildInstanceRequest) SetChildInstanceOwnerId(v int64) *AttachCenChildInstanceRequest {
	s.ChildInstanceOwnerId = &v
	return s
}

func (s *AttachCenChildInstanceRequest) SetChildInstanceRegionId(v string) *AttachCenChildInstanceRequest {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *AttachCenChildInstanceRequest) SetChildInstanceType(v string) *AttachCenChildInstanceRequest {
	s.ChildInstanceType = &v
	return s
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

type CheckTransitRouterServiceRequest struct {
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CheckTransitRouterServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckTransitRouterServiceRequest) GoString() string {
	return s.String()
}

func (s *CheckTransitRouterServiceRequest) SetClientToken(v string) *CheckTransitRouterServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *CheckTransitRouterServiceRequest) SetOwnerAccount(v string) *CheckTransitRouterServiceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CheckTransitRouterServiceRequest) SetOwnerId(v int64) *CheckTransitRouterServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckTransitRouterServiceRequest) SetResourceOwnerAccount(v string) *CheckTransitRouterServiceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckTransitRouterServiceRequest) SetResourceOwnerId(v int64) *CheckTransitRouterServiceRequest {
	s.ResourceOwnerId = &v
	return s
}

type CheckTransitRouterServiceResponseBody struct {
	Enabled   *string `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckTransitRouterServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckTransitRouterServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CheckTransitRouterServiceResponseBody) SetEnabled(v string) *CheckTransitRouterServiceResponseBody {
	s.Enabled = &v
	return s
}

func (s *CheckTransitRouterServiceResponseBody) SetRequestId(v string) *CheckTransitRouterServiceResponseBody {
	s.RequestId = &v
	return s
}

type CheckTransitRouterServiceResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckTransitRouterServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckTransitRouterServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckTransitRouterServiceResponse) GoString() string {
	return s.String()
}

func (s *CheckTransitRouterServiceResponse) SetHeaders(v map[string]*string) *CheckTransitRouterServiceResponse {
	s.Headers = v
	return s
}

func (s *CheckTransitRouterServiceResponse) SetBody(v *CheckTransitRouterServiceResponseBody) *CheckTransitRouterServiceResponse {
	s.Body = v
	return s
}

type CreateCenRequest struct {
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProtectionLevel      *string `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateCenRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCenRequest) GoString() string {
	return s.String()
}

func (s *CreateCenRequest) SetClientToken(v string) *CreateCenRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCenRequest) SetDescription(v string) *CreateCenRequest {
	s.Description = &v
	return s
}

func (s *CreateCenRequest) SetName(v string) *CreateCenRequest {
	s.Name = &v
	return s
}

func (s *CreateCenRequest) SetOwnerAccount(v string) *CreateCenRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateCenRequest) SetOwnerId(v int64) *CreateCenRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateCenRequest) SetProtectionLevel(v string) *CreateCenRequest {
	s.ProtectionLevel = &v
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
	AutoPay                    *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	AutoRenew                  *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AutoRenewDuration          *int32  `json:"AutoRenewDuration,omitempty" xml:"AutoRenewDuration,omitempty"`
	Bandwidth                  *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	BandwidthPackageChargeType *string `json:"BandwidthPackageChargeType,omitempty" xml:"BandwidthPackageChargeType,omitempty"`
	ClientToken                *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description                *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GeographicRegionAId        *string `json:"GeographicRegionAId,omitempty" xml:"GeographicRegionAId,omitempty"`
	GeographicRegionBId        *string `json:"GeographicRegionBId,omitempty" xml:"GeographicRegionBId,omitempty"`
	Name                       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount               *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Period                     *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	PricingCycle               *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	ResourceOwnerAccount       *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId            *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateCenBandwidthPackageRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCenBandwidthPackageRequest) GoString() string {
	return s.String()
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

func (s *CreateCenBandwidthPackageRequest) SetBandwidth(v int32) *CreateCenBandwidthPackageRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateCenBandwidthPackageRequest) SetBandwidthPackageChargeType(v string) *CreateCenBandwidthPackageRequest {
	s.BandwidthPackageChargeType = &v
	return s
}

func (s *CreateCenBandwidthPackageRequest) SetClientToken(v string) *CreateCenBandwidthPackageRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCenBandwidthPackageRequest) SetDescription(v string) *CreateCenBandwidthPackageRequest {
	s.Description = &v
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

func (s *CreateCenBandwidthPackageRequest) SetName(v string) *CreateCenBandwidthPackageRequest {
	s.Name = &v
	return s
}

func (s *CreateCenBandwidthPackageRequest) SetOwnerAccount(v string) *CreateCenBandwidthPackageRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateCenBandwidthPackageRequest) SetOwnerId(v int64) *CreateCenBandwidthPackageRequest {
	s.OwnerId = &v
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

func (s *CreateCenBandwidthPackageRequest) SetResourceOwnerAccount(v string) *CreateCenBandwidthPackageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateCenBandwidthPackageRequest) SetResourceOwnerId(v int64) *CreateCenBandwidthPackageRequest {
	s.ResourceOwnerId = &v
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

type CreateCenChildInstanceRouteEntryToAttachmentRequest struct {
	CenId                     *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ClientToken               *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DestinationCidrBlock      *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	DryRun                    *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount              *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount      *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId           *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RouteTableId              *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
}

func (s CreateCenChildInstanceRouteEntryToAttachmentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCenChildInstanceRouteEntryToAttachmentRequest) GoString() string {
	return s.String()
}

func (s *CreateCenChildInstanceRouteEntryToAttachmentRequest) SetCenId(v string) *CreateCenChildInstanceRouteEntryToAttachmentRequest {
	s.CenId = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToAttachmentRequest) SetClientToken(v string) *CreateCenChildInstanceRouteEntryToAttachmentRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToAttachmentRequest) SetDestinationCidrBlock(v string) *CreateCenChildInstanceRouteEntryToAttachmentRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToAttachmentRequest) SetDryRun(v bool) *CreateCenChildInstanceRouteEntryToAttachmentRequest {
	s.DryRun = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToAttachmentRequest) SetOwnerAccount(v string) *CreateCenChildInstanceRouteEntryToAttachmentRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToAttachmentRequest) SetOwnerId(v int64) *CreateCenChildInstanceRouteEntryToAttachmentRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToAttachmentRequest) SetResourceOwnerAccount(v string) *CreateCenChildInstanceRouteEntryToAttachmentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToAttachmentRequest) SetResourceOwnerId(v int64) *CreateCenChildInstanceRouteEntryToAttachmentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToAttachmentRequest) SetRouteTableId(v string) *CreateCenChildInstanceRouteEntryToAttachmentRequest {
	s.RouteTableId = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToAttachmentRequest) SetTransitRouterAttachmentId(v string) *CreateCenChildInstanceRouteEntryToAttachmentRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

type CreateCenChildInstanceRouteEntryToAttachmentResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCenChildInstanceRouteEntryToAttachmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCenChildInstanceRouteEntryToAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCenChildInstanceRouteEntryToAttachmentResponseBody) SetRequestId(v string) *CreateCenChildInstanceRouteEntryToAttachmentResponseBody {
	s.RequestId = &v
	return s
}

type CreateCenChildInstanceRouteEntryToAttachmentResponse struct {
	Headers map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCenChildInstanceRouteEntryToAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCenChildInstanceRouteEntryToAttachmentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCenChildInstanceRouteEntryToAttachmentResponse) GoString() string {
	return s.String()
}

func (s *CreateCenChildInstanceRouteEntryToAttachmentResponse) SetHeaders(v map[string]*string) *CreateCenChildInstanceRouteEntryToAttachmentResponse {
	s.Headers = v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToAttachmentResponse) SetBody(v *CreateCenChildInstanceRouteEntryToAttachmentResponseBody) *CreateCenChildInstanceRouteEntryToAttachmentResponse {
	s.Body = v
	return s
}

type CreateCenChildInstanceRouteEntryToCenRequest struct {
	CenId                 *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ChildInstanceAliUid   *int64  `json:"ChildInstanceAliUid,omitempty" xml:"ChildInstanceAliUid,omitempty"`
	ChildInstanceId       *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	ChildInstanceRegionId *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	ChildInstanceType     *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	DestinationCidrBlock  *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RouteTableId          *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s CreateCenChildInstanceRouteEntryToCenRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCenChildInstanceRouteEntryToCenRequest) GoString() string {
	return s.String()
}

func (s *CreateCenChildInstanceRouteEntryToCenRequest) SetCenId(v string) *CreateCenChildInstanceRouteEntryToCenRequest {
	s.CenId = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToCenRequest) SetChildInstanceAliUid(v int64) *CreateCenChildInstanceRouteEntryToCenRequest {
	s.ChildInstanceAliUid = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToCenRequest) SetChildInstanceId(v string) *CreateCenChildInstanceRouteEntryToCenRequest {
	s.ChildInstanceId = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToCenRequest) SetChildInstanceRegionId(v string) *CreateCenChildInstanceRouteEntryToCenRequest {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToCenRequest) SetChildInstanceType(v string) *CreateCenChildInstanceRouteEntryToCenRequest {
	s.ChildInstanceType = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToCenRequest) SetDestinationCidrBlock(v string) *CreateCenChildInstanceRouteEntryToCenRequest {
	s.DestinationCidrBlock = &v
	return s
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

func (s *CreateCenChildInstanceRouteEntryToCenRequest) SetRouteTableId(v string) *CreateCenChildInstanceRouteEntryToCenRequest {
	s.RouteTableId = &v
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

type CreateCenInterRegionTrafficQosPolicyRequest struct {
	ClientToken                 *string                                                        `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                      *bool                                                          `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount                *string                                                        `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                     *int64                                                         `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount        *string                                                        `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId             *int64                                                         `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TrafficQosPolicyDescription *string                                                        `json:"TrafficQosPolicyDescription,omitempty" xml:"TrafficQosPolicyDescription,omitempty"`
	TrafficQosPolicyName        *string                                                        `json:"TrafficQosPolicyName,omitempty" xml:"TrafficQosPolicyName,omitempty"`
	TrafficQosQueues            []*CreateCenInterRegionTrafficQosPolicyRequestTrafficQosQueues `json:"TrafficQosQueues,omitempty" xml:"TrafficQosQueues,omitempty" type:"Repeated"`
	TransitRouterAttachmentId   *string                                                        `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	TransitRouterId             *string                                                        `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s CreateCenInterRegionTrafficQosPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCenInterRegionTrafficQosPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateCenInterRegionTrafficQosPolicyRequest) SetClientToken(v string) *CreateCenInterRegionTrafficQosPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCenInterRegionTrafficQosPolicyRequest) SetDryRun(v bool) *CreateCenInterRegionTrafficQosPolicyRequest {
	s.DryRun = &v
	return s
}

func (s *CreateCenInterRegionTrafficQosPolicyRequest) SetOwnerAccount(v string) *CreateCenInterRegionTrafficQosPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateCenInterRegionTrafficQosPolicyRequest) SetOwnerId(v int64) *CreateCenInterRegionTrafficQosPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateCenInterRegionTrafficQosPolicyRequest) SetResourceOwnerAccount(v string) *CreateCenInterRegionTrafficQosPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateCenInterRegionTrafficQosPolicyRequest) SetResourceOwnerId(v int64) *CreateCenInterRegionTrafficQosPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateCenInterRegionTrafficQosPolicyRequest) SetTrafficQosPolicyDescription(v string) *CreateCenInterRegionTrafficQosPolicyRequest {
	s.TrafficQosPolicyDescription = &v
	return s
}

func (s *CreateCenInterRegionTrafficQosPolicyRequest) SetTrafficQosPolicyName(v string) *CreateCenInterRegionTrafficQosPolicyRequest {
	s.TrafficQosPolicyName = &v
	return s
}

func (s *CreateCenInterRegionTrafficQosPolicyRequest) SetTrafficQosQueues(v []*CreateCenInterRegionTrafficQosPolicyRequestTrafficQosQueues) *CreateCenInterRegionTrafficQosPolicyRequest {
	s.TrafficQosQueues = v
	return s
}

func (s *CreateCenInterRegionTrafficQosPolicyRequest) SetTransitRouterAttachmentId(v string) *CreateCenInterRegionTrafficQosPolicyRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *CreateCenInterRegionTrafficQosPolicyRequest) SetTransitRouterId(v string) *CreateCenInterRegionTrafficQosPolicyRequest {
	s.TransitRouterId = &v
	return s
}

type CreateCenInterRegionTrafficQosPolicyRequestTrafficQosQueues struct {
	Dscps                  []*int32 `json:"Dscps,omitempty" xml:"Dscps,omitempty" type:"Repeated"`
	QosQueueDescription    *string  `json:"QosQueueDescription,omitempty" xml:"QosQueueDescription,omitempty"`
	QosQueueName           *string  `json:"QosQueueName,omitempty" xml:"QosQueueName,omitempty"`
	RemainBandwidthPercent *string  `json:"RemainBandwidthPercent,omitempty" xml:"RemainBandwidthPercent,omitempty"`
}

func (s CreateCenInterRegionTrafficQosPolicyRequestTrafficQosQueues) String() string {
	return tea.Prettify(s)
}

func (s CreateCenInterRegionTrafficQosPolicyRequestTrafficQosQueues) GoString() string {
	return s.String()
}

func (s *CreateCenInterRegionTrafficQosPolicyRequestTrafficQosQueues) SetDscps(v []*int32) *CreateCenInterRegionTrafficQosPolicyRequestTrafficQosQueues {
	s.Dscps = v
	return s
}

func (s *CreateCenInterRegionTrafficQosPolicyRequestTrafficQosQueues) SetQosQueueDescription(v string) *CreateCenInterRegionTrafficQosPolicyRequestTrafficQosQueues {
	s.QosQueueDescription = &v
	return s
}

func (s *CreateCenInterRegionTrafficQosPolicyRequestTrafficQosQueues) SetQosQueueName(v string) *CreateCenInterRegionTrafficQosPolicyRequestTrafficQosQueues {
	s.QosQueueName = &v
	return s
}

func (s *CreateCenInterRegionTrafficQosPolicyRequestTrafficQosQueues) SetRemainBandwidthPercent(v string) *CreateCenInterRegionTrafficQosPolicyRequestTrafficQosQueues {
	s.RemainBandwidthPercent = &v
	return s
}

type CreateCenInterRegionTrafficQosPolicyResponseBody struct {
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TrafficQosPolicyId *string `json:"TrafficQosPolicyId,omitempty" xml:"TrafficQosPolicyId,omitempty"`
}

func (s CreateCenInterRegionTrafficQosPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCenInterRegionTrafficQosPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCenInterRegionTrafficQosPolicyResponseBody) SetRequestId(v string) *CreateCenInterRegionTrafficQosPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCenInterRegionTrafficQosPolicyResponseBody) SetTrafficQosPolicyId(v string) *CreateCenInterRegionTrafficQosPolicyResponseBody {
	s.TrafficQosPolicyId = &v
	return s
}

type CreateCenInterRegionTrafficQosPolicyResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCenInterRegionTrafficQosPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCenInterRegionTrafficQosPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCenInterRegionTrafficQosPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateCenInterRegionTrafficQosPolicyResponse) SetHeaders(v map[string]*string) *CreateCenInterRegionTrafficQosPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateCenInterRegionTrafficQosPolicyResponse) SetBody(v *CreateCenInterRegionTrafficQosPolicyResponseBody) *CreateCenInterRegionTrafficQosPolicyResponse {
	s.Body = v
	return s
}

type CreateCenRouteMapRequest struct {
	AsPathMatchMode                    *string   `json:"AsPathMatchMode,omitempty" xml:"AsPathMatchMode,omitempty"`
	CenId                              *string   `json:"CenId,omitempty" xml:"CenId,omitempty"`
	CenRegionId                        *string   `json:"CenRegionId,omitempty" xml:"CenRegionId,omitempty"`
	CidrMatchMode                      *string   `json:"CidrMatchMode,omitempty" xml:"CidrMatchMode,omitempty"`
	CommunityMatchMode                 *string   `json:"CommunityMatchMode,omitempty" xml:"CommunityMatchMode,omitempty"`
	CommunityOperateMode               *string   `json:"CommunityOperateMode,omitempty" xml:"CommunityOperateMode,omitempty"`
	Description                        *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	DestinationChildInstanceTypes      []*string `json:"DestinationChildInstanceTypes,omitempty" xml:"DestinationChildInstanceTypes,omitempty" type:"Repeated"`
	DestinationCidrBlocks              []*string `json:"DestinationCidrBlocks,omitempty" xml:"DestinationCidrBlocks,omitempty" type:"Repeated"`
	DestinationInstanceIds             []*string `json:"DestinationInstanceIds,omitempty" xml:"DestinationInstanceIds,omitempty" type:"Repeated"`
	DestinationInstanceIdsReverseMatch *bool     `json:"DestinationInstanceIdsReverseMatch,omitempty" xml:"DestinationInstanceIdsReverseMatch,omitempty"`
	DestinationRouteTableIds           []*string `json:"DestinationRouteTableIds,omitempty" xml:"DestinationRouteTableIds,omitempty" type:"Repeated"`
	MapResult                          *string   `json:"MapResult,omitempty" xml:"MapResult,omitempty"`
	MatchAsns                          []*int64  `json:"MatchAsns,omitempty" xml:"MatchAsns,omitempty" type:"Repeated"`
	MatchCommunitySet                  []*string `json:"MatchCommunitySet,omitempty" xml:"MatchCommunitySet,omitempty" type:"Repeated"`
	NextPriority                       *int32    `json:"NextPriority,omitempty" xml:"NextPriority,omitempty"`
	OperateCommunitySet                []*string `json:"OperateCommunitySet,omitempty" xml:"OperateCommunitySet,omitempty" type:"Repeated"`
	OwnerAccount                       *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                            *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Preference                         *int32    `json:"Preference,omitempty" xml:"Preference,omitempty"`
	PrependAsPath                      []*int64  `json:"PrependAsPath,omitempty" xml:"PrependAsPath,omitempty" type:"Repeated"`
	Priority                           *int32    `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ResourceOwnerAccount               *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId                    *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RouteTypes                         []*string `json:"RouteTypes,omitempty" xml:"RouteTypes,omitempty" type:"Repeated"`
	SourceChildInstanceTypes           []*string `json:"SourceChildInstanceTypes,omitempty" xml:"SourceChildInstanceTypes,omitempty" type:"Repeated"`
	SourceInstanceIds                  []*string `json:"SourceInstanceIds,omitempty" xml:"SourceInstanceIds,omitempty" type:"Repeated"`
	SourceInstanceIdsReverseMatch      *bool     `json:"SourceInstanceIdsReverseMatch,omitempty" xml:"SourceInstanceIdsReverseMatch,omitempty"`
	SourceRegionIds                    []*string `json:"SourceRegionIds,omitempty" xml:"SourceRegionIds,omitempty" type:"Repeated"`
	SourceRouteTableIds                []*string `json:"SourceRouteTableIds,omitempty" xml:"SourceRouteTableIds,omitempty" type:"Repeated"`
	TransmitDirection                  *string   `json:"TransmitDirection,omitempty" xml:"TransmitDirection,omitempty"`
}

func (s CreateCenRouteMapRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCenRouteMapRequest) GoString() string {
	return s.String()
}

func (s *CreateCenRouteMapRequest) SetAsPathMatchMode(v string) *CreateCenRouteMapRequest {
	s.AsPathMatchMode = &v
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

func (s *CreateCenRouteMapRequest) SetCidrMatchMode(v string) *CreateCenRouteMapRequest {
	s.CidrMatchMode = &v
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

func (s *CreateCenRouteMapRequest) SetDescription(v string) *CreateCenRouteMapRequest {
	s.Description = &v
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

func (s *CreateCenRouteMapRequest) SetDestinationInstanceIds(v []*string) *CreateCenRouteMapRequest {
	s.DestinationInstanceIds = v
	return s
}

func (s *CreateCenRouteMapRequest) SetDestinationInstanceIdsReverseMatch(v bool) *CreateCenRouteMapRequest {
	s.DestinationInstanceIdsReverseMatch = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetDestinationRouteTableIds(v []*string) *CreateCenRouteMapRequest {
	s.DestinationRouteTableIds = v
	return s
}

func (s *CreateCenRouteMapRequest) SetMapResult(v string) *CreateCenRouteMapRequest {
	s.MapResult = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetMatchAsns(v []*int64) *CreateCenRouteMapRequest {
	s.MatchAsns = v
	return s
}

func (s *CreateCenRouteMapRequest) SetMatchCommunitySet(v []*string) *CreateCenRouteMapRequest {
	s.MatchCommunitySet = v
	return s
}

func (s *CreateCenRouteMapRequest) SetNextPriority(v int32) *CreateCenRouteMapRequest {
	s.NextPriority = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetOperateCommunitySet(v []*string) *CreateCenRouteMapRequest {
	s.OperateCommunitySet = v
	return s
}

func (s *CreateCenRouteMapRequest) SetOwnerAccount(v string) *CreateCenRouteMapRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetOwnerId(v int64) *CreateCenRouteMapRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetPreference(v int32) *CreateCenRouteMapRequest {
	s.Preference = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetPrependAsPath(v []*int64) *CreateCenRouteMapRequest {
	s.PrependAsPath = v
	return s
}

func (s *CreateCenRouteMapRequest) SetPriority(v int32) *CreateCenRouteMapRequest {
	s.Priority = &v
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

func (s *CreateCenRouteMapRequest) SetRouteTypes(v []*string) *CreateCenRouteMapRequest {
	s.RouteTypes = v
	return s
}

func (s *CreateCenRouteMapRequest) SetSourceChildInstanceTypes(v []*string) *CreateCenRouteMapRequest {
	s.SourceChildInstanceTypes = v
	return s
}

func (s *CreateCenRouteMapRequest) SetSourceInstanceIds(v []*string) *CreateCenRouteMapRequest {
	s.SourceInstanceIds = v
	return s
}

func (s *CreateCenRouteMapRequest) SetSourceInstanceIdsReverseMatch(v bool) *CreateCenRouteMapRequest {
	s.SourceInstanceIdsReverseMatch = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetSourceRegionIds(v []*string) *CreateCenRouteMapRequest {
	s.SourceRegionIds = v
	return s
}

func (s *CreateCenRouteMapRequest) SetSourceRouteTableIds(v []*string) *CreateCenRouteMapRequest {
	s.SourceRouteTableIds = v
	return s
}

func (s *CreateCenRouteMapRequest) SetTransmitDirection(v string) *CreateCenRouteMapRequest {
	s.TransmitDirection = &v
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
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FlowLogName          *string `json:"FlowLogName,omitempty" xml:"FlowLogName,omitempty"`
	LogStoreName         *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProjectName          *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateFlowlogRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowlogRequest) GoString() string {
	return s.String()
}

func (s *CreateFlowlogRequest) SetCenId(v string) *CreateFlowlogRequest {
	s.CenId = &v
	return s
}

func (s *CreateFlowlogRequest) SetClientToken(v string) *CreateFlowlogRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateFlowlogRequest) SetDescription(v string) *CreateFlowlogRequest {
	s.Description = &v
	return s
}

func (s *CreateFlowlogRequest) SetFlowLogName(v string) *CreateFlowlogRequest {
	s.FlowLogName = &v
	return s
}

func (s *CreateFlowlogRequest) SetLogStoreName(v string) *CreateFlowlogRequest {
	s.LogStoreName = &v
	return s
}

func (s *CreateFlowlogRequest) SetOwnerAccount(v string) *CreateFlowlogRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateFlowlogRequest) SetOwnerId(v int64) *CreateFlowlogRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateFlowlogRequest) SetProjectName(v string) *CreateFlowlogRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateFlowlogRequest) SetRegionId(v string) *CreateFlowlogRequest {
	s.RegionId = &v
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

type CreateFlowlogResponseBody struct {
	FlowLogId *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateFlowlogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowlogResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFlowlogResponseBody) SetFlowLogId(v string) *CreateFlowlogResponseBody {
	s.FlowLogId = &v
	return s
}

func (s *CreateFlowlogResponseBody) SetRequestId(v string) *CreateFlowlogResponseBody {
	s.RequestId = &v
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

type CreateTrafficMarkingPolicyRequest struct {
	ClientToken                     *string                                               `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                          *bool                                                 `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	MarkingDscp                     *int32                                                `json:"MarkingDscp,omitempty" xml:"MarkingDscp,omitempty"`
	OwnerAccount                    *string                                               `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                         *int64                                                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Priority                        *int32                                                `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ResourceOwnerAccount            *string                                               `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId                 *int64                                                `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TrafficMarkingPolicyDescription *string                                               `json:"TrafficMarkingPolicyDescription,omitempty" xml:"TrafficMarkingPolicyDescription,omitempty"`
	TrafficMarkingPolicyName        *string                                               `json:"TrafficMarkingPolicyName,omitempty" xml:"TrafficMarkingPolicyName,omitempty"`
	TrafficMatchRules               []*CreateTrafficMarkingPolicyRequestTrafficMatchRules `json:"TrafficMatchRules,omitempty" xml:"TrafficMatchRules,omitempty" type:"Repeated"`
	TransitRouterId                 *string                                               `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s CreateTrafficMarkingPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTrafficMarkingPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateTrafficMarkingPolicyRequest) SetClientToken(v string) *CreateTrafficMarkingPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTrafficMarkingPolicyRequest) SetDryRun(v bool) *CreateTrafficMarkingPolicyRequest {
	s.DryRun = &v
	return s
}

func (s *CreateTrafficMarkingPolicyRequest) SetMarkingDscp(v int32) *CreateTrafficMarkingPolicyRequest {
	s.MarkingDscp = &v
	return s
}

func (s *CreateTrafficMarkingPolicyRequest) SetOwnerAccount(v string) *CreateTrafficMarkingPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateTrafficMarkingPolicyRequest) SetOwnerId(v int64) *CreateTrafficMarkingPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTrafficMarkingPolicyRequest) SetPriority(v int32) *CreateTrafficMarkingPolicyRequest {
	s.Priority = &v
	return s
}

func (s *CreateTrafficMarkingPolicyRequest) SetResourceOwnerAccount(v string) *CreateTrafficMarkingPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateTrafficMarkingPolicyRequest) SetResourceOwnerId(v int64) *CreateTrafficMarkingPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateTrafficMarkingPolicyRequest) SetTrafficMarkingPolicyDescription(v string) *CreateTrafficMarkingPolicyRequest {
	s.TrafficMarkingPolicyDescription = &v
	return s
}

func (s *CreateTrafficMarkingPolicyRequest) SetTrafficMarkingPolicyName(v string) *CreateTrafficMarkingPolicyRequest {
	s.TrafficMarkingPolicyName = &v
	return s
}

func (s *CreateTrafficMarkingPolicyRequest) SetTrafficMatchRules(v []*CreateTrafficMarkingPolicyRequestTrafficMatchRules) *CreateTrafficMarkingPolicyRequest {
	s.TrafficMatchRules = v
	return s
}

func (s *CreateTrafficMarkingPolicyRequest) SetTransitRouterId(v string) *CreateTrafficMarkingPolicyRequest {
	s.TransitRouterId = &v
	return s
}

type CreateTrafficMarkingPolicyRequestTrafficMatchRules struct {
	DstCidr                     *string  `json:"DstCidr,omitempty" xml:"DstCidr,omitempty"`
	DstPortRange                []*int32 `json:"DstPortRange,omitempty" xml:"DstPortRange,omitempty" type:"Repeated"`
	MatchDscp                   *int32   `json:"MatchDscp,omitempty" xml:"MatchDscp,omitempty"`
	Protocol                    *string  `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	SrcCidr                     *string  `json:"SrcCidr,omitempty" xml:"SrcCidr,omitempty"`
	SrcPortRange                []*int32 `json:"SrcPortRange,omitempty" xml:"SrcPortRange,omitempty" type:"Repeated"`
	TrafficMatchRuleDescription *string  `json:"TrafficMatchRuleDescription,omitempty" xml:"TrafficMatchRuleDescription,omitempty"`
	TrafficMatchRuleName        *string  `json:"TrafficMatchRuleName,omitempty" xml:"TrafficMatchRuleName,omitempty"`
}

func (s CreateTrafficMarkingPolicyRequestTrafficMatchRules) String() string {
	return tea.Prettify(s)
}

func (s CreateTrafficMarkingPolicyRequestTrafficMatchRules) GoString() string {
	return s.String()
}

func (s *CreateTrafficMarkingPolicyRequestTrafficMatchRules) SetDstCidr(v string) *CreateTrafficMarkingPolicyRequestTrafficMatchRules {
	s.DstCidr = &v
	return s
}

func (s *CreateTrafficMarkingPolicyRequestTrafficMatchRules) SetDstPortRange(v []*int32) *CreateTrafficMarkingPolicyRequestTrafficMatchRules {
	s.DstPortRange = v
	return s
}

func (s *CreateTrafficMarkingPolicyRequestTrafficMatchRules) SetMatchDscp(v int32) *CreateTrafficMarkingPolicyRequestTrafficMatchRules {
	s.MatchDscp = &v
	return s
}

func (s *CreateTrafficMarkingPolicyRequestTrafficMatchRules) SetProtocol(v string) *CreateTrafficMarkingPolicyRequestTrafficMatchRules {
	s.Protocol = &v
	return s
}

func (s *CreateTrafficMarkingPolicyRequestTrafficMatchRules) SetSrcCidr(v string) *CreateTrafficMarkingPolicyRequestTrafficMatchRules {
	s.SrcCidr = &v
	return s
}

func (s *CreateTrafficMarkingPolicyRequestTrafficMatchRules) SetSrcPortRange(v []*int32) *CreateTrafficMarkingPolicyRequestTrafficMatchRules {
	s.SrcPortRange = v
	return s
}

func (s *CreateTrafficMarkingPolicyRequestTrafficMatchRules) SetTrafficMatchRuleDescription(v string) *CreateTrafficMarkingPolicyRequestTrafficMatchRules {
	s.TrafficMatchRuleDescription = &v
	return s
}

func (s *CreateTrafficMarkingPolicyRequestTrafficMatchRules) SetTrafficMatchRuleName(v string) *CreateTrafficMarkingPolicyRequestTrafficMatchRules {
	s.TrafficMatchRuleName = &v
	return s
}

type CreateTrafficMarkingPolicyResponseBody struct {
	RequestId              *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TrafficMarkingPolicyId *string `json:"TrafficMarkingPolicyId,omitempty" xml:"TrafficMarkingPolicyId,omitempty"`
}

func (s CreateTrafficMarkingPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTrafficMarkingPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTrafficMarkingPolicyResponseBody) SetRequestId(v string) *CreateTrafficMarkingPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTrafficMarkingPolicyResponseBody) SetTrafficMarkingPolicyId(v string) *CreateTrafficMarkingPolicyResponseBody {
	s.TrafficMarkingPolicyId = &v
	return s
}

type CreateTrafficMarkingPolicyResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTrafficMarkingPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTrafficMarkingPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTrafficMarkingPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateTrafficMarkingPolicyResponse) SetHeaders(v map[string]*string) *CreateTrafficMarkingPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateTrafficMarkingPolicyResponse) SetBody(v *CreateTrafficMarkingPolicyResponseBody) *CreateTrafficMarkingPolicyResponse {
	s.Body = v
	return s
}

type CreateTransitRouterRequest struct {
	CenId                    *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ClientToken              *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                   *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount             *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount     *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId          *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TransitRouterDescription *string `json:"TransitRouterDescription,omitempty" xml:"TransitRouterDescription,omitempty"`
	TransitRouterName        *string `json:"TransitRouterName,omitempty" xml:"TransitRouterName,omitempty"`
}

func (s CreateTransitRouterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTransitRouterRequest) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterRequest) SetCenId(v string) *CreateTransitRouterRequest {
	s.CenId = &v
	return s
}

func (s *CreateTransitRouterRequest) SetClientToken(v string) *CreateTransitRouterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTransitRouterRequest) SetDryRun(v bool) *CreateTransitRouterRequest {
	s.DryRun = &v
	return s
}

func (s *CreateTransitRouterRequest) SetOwnerAccount(v string) *CreateTransitRouterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateTransitRouterRequest) SetOwnerId(v int64) *CreateTransitRouterRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTransitRouterRequest) SetRegionId(v string) *CreateTransitRouterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTransitRouterRequest) SetResourceOwnerAccount(v string) *CreateTransitRouterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateTransitRouterRequest) SetResourceOwnerId(v int64) *CreateTransitRouterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateTransitRouterRequest) SetTransitRouterDescription(v string) *CreateTransitRouterRequest {
	s.TransitRouterDescription = &v
	return s
}

func (s *CreateTransitRouterRequest) SetTransitRouterName(v string) *CreateTransitRouterRequest {
	s.TransitRouterName = &v
	return s
}

type CreateTransitRouterResponseBody struct {
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s CreateTransitRouterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTransitRouterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterResponseBody) SetRequestId(v string) *CreateTransitRouterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTransitRouterResponseBody) SetTransitRouterId(v string) *CreateTransitRouterResponseBody {
	s.TransitRouterId = &v
	return s
}

type CreateTransitRouterResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTransitRouterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTransitRouterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTransitRouterResponse) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterResponse) SetHeaders(v map[string]*string) *CreateTransitRouterResponse {
	s.Headers = v
	return s
}

func (s *CreateTransitRouterResponse) SetBody(v *CreateTransitRouterResponseBody) *CreateTransitRouterResponse {
	s.Body = v
	return s
}

type CreateTransitRouterPeerAttachmentRequest struct {
	AutoPublishRouteEnabled            *bool   `json:"AutoPublishRouteEnabled,omitempty" xml:"AutoPublishRouteEnabled,omitempty"`
	Bandwidth                          *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	BandwidthType                      *string `json:"BandwidthType,omitempty" xml:"BandwidthType,omitempty"`
	CenBandwidthPackageId              *string `json:"CenBandwidthPackageId,omitempty" xml:"CenBandwidthPackageId,omitempty"`
	CenId                              *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ClientToken                        *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                             *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount                       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PeerTransitRouterId                *string `json:"PeerTransitRouterId,omitempty" xml:"PeerTransitRouterId,omitempty"`
	PeerTransitRouterRegionId          *string `json:"PeerTransitRouterRegionId,omitempty" xml:"PeerTransitRouterRegionId,omitempty"`
	RegionId                           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount               *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId                    *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TransitRouterAttachmentDescription *string `json:"TransitRouterAttachmentDescription,omitempty" xml:"TransitRouterAttachmentDescription,omitempty"`
	TransitRouterAttachmentName        *string `json:"TransitRouterAttachmentName,omitempty" xml:"TransitRouterAttachmentName,omitempty"`
	TransitRouterId                    *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s CreateTransitRouterPeerAttachmentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTransitRouterPeerAttachmentRequest) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterPeerAttachmentRequest) SetAutoPublishRouteEnabled(v bool) *CreateTransitRouterPeerAttachmentRequest {
	s.AutoPublishRouteEnabled = &v
	return s
}

func (s *CreateTransitRouterPeerAttachmentRequest) SetBandwidth(v int32) *CreateTransitRouterPeerAttachmentRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateTransitRouterPeerAttachmentRequest) SetBandwidthType(v string) *CreateTransitRouterPeerAttachmentRequest {
	s.BandwidthType = &v
	return s
}

func (s *CreateTransitRouterPeerAttachmentRequest) SetCenBandwidthPackageId(v string) *CreateTransitRouterPeerAttachmentRequest {
	s.CenBandwidthPackageId = &v
	return s
}

func (s *CreateTransitRouterPeerAttachmentRequest) SetCenId(v string) *CreateTransitRouterPeerAttachmentRequest {
	s.CenId = &v
	return s
}

func (s *CreateTransitRouterPeerAttachmentRequest) SetClientToken(v string) *CreateTransitRouterPeerAttachmentRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTransitRouterPeerAttachmentRequest) SetDryRun(v bool) *CreateTransitRouterPeerAttachmentRequest {
	s.DryRun = &v
	return s
}

func (s *CreateTransitRouterPeerAttachmentRequest) SetOwnerAccount(v string) *CreateTransitRouterPeerAttachmentRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateTransitRouterPeerAttachmentRequest) SetOwnerId(v int64) *CreateTransitRouterPeerAttachmentRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTransitRouterPeerAttachmentRequest) SetPeerTransitRouterId(v string) *CreateTransitRouterPeerAttachmentRequest {
	s.PeerTransitRouterId = &v
	return s
}

func (s *CreateTransitRouterPeerAttachmentRequest) SetPeerTransitRouterRegionId(v string) *CreateTransitRouterPeerAttachmentRequest {
	s.PeerTransitRouterRegionId = &v
	return s
}

func (s *CreateTransitRouterPeerAttachmentRequest) SetRegionId(v string) *CreateTransitRouterPeerAttachmentRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTransitRouterPeerAttachmentRequest) SetResourceOwnerAccount(v string) *CreateTransitRouterPeerAttachmentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateTransitRouterPeerAttachmentRequest) SetResourceOwnerId(v int64) *CreateTransitRouterPeerAttachmentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateTransitRouterPeerAttachmentRequest) SetTransitRouterAttachmentDescription(v string) *CreateTransitRouterPeerAttachmentRequest {
	s.TransitRouterAttachmentDescription = &v
	return s
}

func (s *CreateTransitRouterPeerAttachmentRequest) SetTransitRouterAttachmentName(v string) *CreateTransitRouterPeerAttachmentRequest {
	s.TransitRouterAttachmentName = &v
	return s
}

func (s *CreateTransitRouterPeerAttachmentRequest) SetTransitRouterId(v string) *CreateTransitRouterPeerAttachmentRequest {
	s.TransitRouterId = &v
	return s
}

type CreateTransitRouterPeerAttachmentResponseBody struct {
	RequestId                 *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
}

func (s CreateTransitRouterPeerAttachmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTransitRouterPeerAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterPeerAttachmentResponseBody) SetRequestId(v string) *CreateTransitRouterPeerAttachmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTransitRouterPeerAttachmentResponseBody) SetTransitRouterAttachmentId(v string) *CreateTransitRouterPeerAttachmentResponseBody {
	s.TransitRouterAttachmentId = &v
	return s
}

type CreateTransitRouterPeerAttachmentResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTransitRouterPeerAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTransitRouterPeerAttachmentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTransitRouterPeerAttachmentResponse) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterPeerAttachmentResponse) SetHeaders(v map[string]*string) *CreateTransitRouterPeerAttachmentResponse {
	s.Headers = v
	return s
}

func (s *CreateTransitRouterPeerAttachmentResponse) SetBody(v *CreateTransitRouterPeerAttachmentResponseBody) *CreateTransitRouterPeerAttachmentResponse {
	s.Body = v
	return s
}

type CreateTransitRouterRouteEntryRequest struct {
	ClientToken                                 *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                                      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount                                *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                                     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount                        *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId                             *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TransitRouterRouteEntryDescription          *string `json:"TransitRouterRouteEntryDescription,omitempty" xml:"TransitRouterRouteEntryDescription,omitempty"`
	TransitRouterRouteEntryDestinationCidrBlock *string `json:"TransitRouterRouteEntryDestinationCidrBlock,omitempty" xml:"TransitRouterRouteEntryDestinationCidrBlock,omitempty"`
	TransitRouterRouteEntryName                 *string `json:"TransitRouterRouteEntryName,omitempty" xml:"TransitRouterRouteEntryName,omitempty"`
	TransitRouterRouteEntryNextHopId            *string `json:"TransitRouterRouteEntryNextHopId,omitempty" xml:"TransitRouterRouteEntryNextHopId,omitempty"`
	TransitRouterRouteEntryNextHopType          *string `json:"TransitRouterRouteEntryNextHopType,omitempty" xml:"TransitRouterRouteEntryNextHopType,omitempty"`
	TransitRouterRouteTableId                   *string `json:"TransitRouterRouteTableId,omitempty" xml:"TransitRouterRouteTableId,omitempty"`
}

func (s CreateTransitRouterRouteEntryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTransitRouterRouteEntryRequest) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterRouteEntryRequest) SetClientToken(v string) *CreateTransitRouterRouteEntryRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTransitRouterRouteEntryRequest) SetDryRun(v bool) *CreateTransitRouterRouteEntryRequest {
	s.DryRun = &v
	return s
}

func (s *CreateTransitRouterRouteEntryRequest) SetOwnerAccount(v string) *CreateTransitRouterRouteEntryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateTransitRouterRouteEntryRequest) SetOwnerId(v int64) *CreateTransitRouterRouteEntryRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTransitRouterRouteEntryRequest) SetResourceOwnerAccount(v string) *CreateTransitRouterRouteEntryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateTransitRouterRouteEntryRequest) SetResourceOwnerId(v int64) *CreateTransitRouterRouteEntryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateTransitRouterRouteEntryRequest) SetTransitRouterRouteEntryDescription(v string) *CreateTransitRouterRouteEntryRequest {
	s.TransitRouterRouteEntryDescription = &v
	return s
}

func (s *CreateTransitRouterRouteEntryRequest) SetTransitRouterRouteEntryDestinationCidrBlock(v string) *CreateTransitRouterRouteEntryRequest {
	s.TransitRouterRouteEntryDestinationCidrBlock = &v
	return s
}

func (s *CreateTransitRouterRouteEntryRequest) SetTransitRouterRouteEntryName(v string) *CreateTransitRouterRouteEntryRequest {
	s.TransitRouterRouteEntryName = &v
	return s
}

func (s *CreateTransitRouterRouteEntryRequest) SetTransitRouterRouteEntryNextHopId(v string) *CreateTransitRouterRouteEntryRequest {
	s.TransitRouterRouteEntryNextHopId = &v
	return s
}

func (s *CreateTransitRouterRouteEntryRequest) SetTransitRouterRouteEntryNextHopType(v string) *CreateTransitRouterRouteEntryRequest {
	s.TransitRouterRouteEntryNextHopType = &v
	return s
}

func (s *CreateTransitRouterRouteEntryRequest) SetTransitRouterRouteTableId(v string) *CreateTransitRouterRouteEntryRequest {
	s.TransitRouterRouteTableId = &v
	return s
}

type CreateTransitRouterRouteEntryResponseBody struct {
	RequestId                 *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TransitRouterRouteEntryId *string `json:"TransitRouterRouteEntryId,omitempty" xml:"TransitRouterRouteEntryId,omitempty"`
}

func (s CreateTransitRouterRouteEntryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTransitRouterRouteEntryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterRouteEntryResponseBody) SetRequestId(v string) *CreateTransitRouterRouteEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTransitRouterRouteEntryResponseBody) SetTransitRouterRouteEntryId(v string) *CreateTransitRouterRouteEntryResponseBody {
	s.TransitRouterRouteEntryId = &v
	return s
}

type CreateTransitRouterRouteEntryResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTransitRouterRouteEntryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTransitRouterRouteEntryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTransitRouterRouteEntryResponse) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterRouteEntryResponse) SetHeaders(v map[string]*string) *CreateTransitRouterRouteEntryResponse {
	s.Headers = v
	return s
}

func (s *CreateTransitRouterRouteEntryResponse) SetBody(v *CreateTransitRouterRouteEntryResponseBody) *CreateTransitRouterRouteEntryResponse {
	s.Body = v
	return s
}

type CreateTransitRouterRouteTableRequest struct {
	ClientToken                        *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                             *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount                       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount               *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId                    *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TransitRouterId                    *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	TransitRouterRouteTableDescription *string `json:"TransitRouterRouteTableDescription,omitempty" xml:"TransitRouterRouteTableDescription,omitempty"`
	TransitRouterRouteTableName        *string `json:"TransitRouterRouteTableName,omitempty" xml:"TransitRouterRouteTableName,omitempty"`
}

func (s CreateTransitRouterRouteTableRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTransitRouterRouteTableRequest) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterRouteTableRequest) SetClientToken(v string) *CreateTransitRouterRouteTableRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTransitRouterRouteTableRequest) SetDryRun(v bool) *CreateTransitRouterRouteTableRequest {
	s.DryRun = &v
	return s
}

func (s *CreateTransitRouterRouteTableRequest) SetOwnerAccount(v string) *CreateTransitRouterRouteTableRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateTransitRouterRouteTableRequest) SetOwnerId(v int64) *CreateTransitRouterRouteTableRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTransitRouterRouteTableRequest) SetResourceOwnerAccount(v string) *CreateTransitRouterRouteTableRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateTransitRouterRouteTableRequest) SetResourceOwnerId(v int64) *CreateTransitRouterRouteTableRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateTransitRouterRouteTableRequest) SetTransitRouterId(v string) *CreateTransitRouterRouteTableRequest {
	s.TransitRouterId = &v
	return s
}

func (s *CreateTransitRouterRouteTableRequest) SetTransitRouterRouteTableDescription(v string) *CreateTransitRouterRouteTableRequest {
	s.TransitRouterRouteTableDescription = &v
	return s
}

func (s *CreateTransitRouterRouteTableRequest) SetTransitRouterRouteTableName(v string) *CreateTransitRouterRouteTableRequest {
	s.TransitRouterRouteTableName = &v
	return s
}

type CreateTransitRouterRouteTableResponseBody struct {
	RequestId                 *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TransitRouterRouteTableId *string `json:"TransitRouterRouteTableId,omitempty" xml:"TransitRouterRouteTableId,omitempty"`
}

func (s CreateTransitRouterRouteTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTransitRouterRouteTableResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterRouteTableResponseBody) SetRequestId(v string) *CreateTransitRouterRouteTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTransitRouterRouteTableResponseBody) SetTransitRouterRouteTableId(v string) *CreateTransitRouterRouteTableResponseBody {
	s.TransitRouterRouteTableId = &v
	return s
}

type CreateTransitRouterRouteTableResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTransitRouterRouteTableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTransitRouterRouteTableResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTransitRouterRouteTableResponse) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterRouteTableResponse) SetHeaders(v map[string]*string) *CreateTransitRouterRouteTableResponse {
	s.Headers = v
	return s
}

func (s *CreateTransitRouterRouteTableResponse) SetBody(v *CreateTransitRouterRouteTableResponseBody) *CreateTransitRouterRouteTableResponse {
	s.Body = v
	return s
}

type CreateTransitRouterVbrAttachmentRequest struct {
	AutoPublishRouteEnabled            *bool   `json:"AutoPublishRouteEnabled,omitempty" xml:"AutoPublishRouteEnabled,omitempty"`
	CenId                              *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ClientToken                        *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                             *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount                       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId                           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount               *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId                    *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TransitRouterAttachmentDescription *string `json:"TransitRouterAttachmentDescription,omitempty" xml:"TransitRouterAttachmentDescription,omitempty"`
	TransitRouterAttachmentName        *string `json:"TransitRouterAttachmentName,omitempty" xml:"TransitRouterAttachmentName,omitempty"`
	TransitRouterId                    *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	VbrId                              *string `json:"VbrId,omitempty" xml:"VbrId,omitempty"`
	VbrOwnerId                         *int64  `json:"VbrOwnerId,omitempty" xml:"VbrOwnerId,omitempty"`
}

func (s CreateTransitRouterVbrAttachmentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTransitRouterVbrAttachmentRequest) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterVbrAttachmentRequest) SetAutoPublishRouteEnabled(v bool) *CreateTransitRouterVbrAttachmentRequest {
	s.AutoPublishRouteEnabled = &v
	return s
}

func (s *CreateTransitRouterVbrAttachmentRequest) SetCenId(v string) *CreateTransitRouterVbrAttachmentRequest {
	s.CenId = &v
	return s
}

func (s *CreateTransitRouterVbrAttachmentRequest) SetClientToken(v string) *CreateTransitRouterVbrAttachmentRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTransitRouterVbrAttachmentRequest) SetDryRun(v bool) *CreateTransitRouterVbrAttachmentRequest {
	s.DryRun = &v
	return s
}

func (s *CreateTransitRouterVbrAttachmentRequest) SetOwnerAccount(v string) *CreateTransitRouterVbrAttachmentRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateTransitRouterVbrAttachmentRequest) SetOwnerId(v int64) *CreateTransitRouterVbrAttachmentRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTransitRouterVbrAttachmentRequest) SetRegionId(v string) *CreateTransitRouterVbrAttachmentRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTransitRouterVbrAttachmentRequest) SetResourceOwnerAccount(v string) *CreateTransitRouterVbrAttachmentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateTransitRouterVbrAttachmentRequest) SetResourceOwnerId(v int64) *CreateTransitRouterVbrAttachmentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateTransitRouterVbrAttachmentRequest) SetTransitRouterAttachmentDescription(v string) *CreateTransitRouterVbrAttachmentRequest {
	s.TransitRouterAttachmentDescription = &v
	return s
}

func (s *CreateTransitRouterVbrAttachmentRequest) SetTransitRouterAttachmentName(v string) *CreateTransitRouterVbrAttachmentRequest {
	s.TransitRouterAttachmentName = &v
	return s
}

func (s *CreateTransitRouterVbrAttachmentRequest) SetTransitRouterId(v string) *CreateTransitRouterVbrAttachmentRequest {
	s.TransitRouterId = &v
	return s
}

func (s *CreateTransitRouterVbrAttachmentRequest) SetVbrId(v string) *CreateTransitRouterVbrAttachmentRequest {
	s.VbrId = &v
	return s
}

func (s *CreateTransitRouterVbrAttachmentRequest) SetVbrOwnerId(v int64) *CreateTransitRouterVbrAttachmentRequest {
	s.VbrOwnerId = &v
	return s
}

type CreateTransitRouterVbrAttachmentResponseBody struct {
	RequestId                 *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
}

func (s CreateTransitRouterVbrAttachmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTransitRouterVbrAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterVbrAttachmentResponseBody) SetRequestId(v string) *CreateTransitRouterVbrAttachmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTransitRouterVbrAttachmentResponseBody) SetTransitRouterAttachmentId(v string) *CreateTransitRouterVbrAttachmentResponseBody {
	s.TransitRouterAttachmentId = &v
	return s
}

type CreateTransitRouterVbrAttachmentResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTransitRouterVbrAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTransitRouterVbrAttachmentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTransitRouterVbrAttachmentResponse) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterVbrAttachmentResponse) SetHeaders(v map[string]*string) *CreateTransitRouterVbrAttachmentResponse {
	s.Headers = v
	return s
}

func (s *CreateTransitRouterVbrAttachmentResponse) SetBody(v *CreateTransitRouterVbrAttachmentResponseBody) *CreateTransitRouterVbrAttachmentResponse {
	s.Body = v
	return s
}

type CreateTransitRouterVpcAttachmentRequest struct {
	CenId                              *string                                                `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ChargeType                         *string                                                `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ClientToken                        *string                                                `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                             *bool                                                  `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount                       *string                                                `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                            *int64                                                 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId                           *string                                                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount               *string                                                `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId                    *int64                                                 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TransitRouterAttachmentDescription *string                                                `json:"TransitRouterAttachmentDescription,omitempty" xml:"TransitRouterAttachmentDescription,omitempty"`
	TransitRouterAttachmentName        *string                                                `json:"TransitRouterAttachmentName,omitempty" xml:"TransitRouterAttachmentName,omitempty"`
	TransitRouterId                    *string                                                `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	VpcId                              *string                                                `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcOwnerId                         *int64                                                 `json:"VpcOwnerId,omitempty" xml:"VpcOwnerId,omitempty"`
	ZoneMappings                       []*CreateTransitRouterVpcAttachmentRequestZoneMappings `json:"ZoneMappings,omitempty" xml:"ZoneMappings,omitempty" type:"Repeated"`
}

func (s CreateTransitRouterVpcAttachmentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTransitRouterVpcAttachmentRequest) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterVpcAttachmentRequest) SetCenId(v string) *CreateTransitRouterVpcAttachmentRequest {
	s.CenId = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentRequest) SetChargeType(v string) *CreateTransitRouterVpcAttachmentRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentRequest) SetClientToken(v string) *CreateTransitRouterVpcAttachmentRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentRequest) SetDryRun(v bool) *CreateTransitRouterVpcAttachmentRequest {
	s.DryRun = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentRequest) SetOwnerAccount(v string) *CreateTransitRouterVpcAttachmentRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentRequest) SetOwnerId(v int64) *CreateTransitRouterVpcAttachmentRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentRequest) SetRegionId(v string) *CreateTransitRouterVpcAttachmentRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentRequest) SetResourceOwnerAccount(v string) *CreateTransitRouterVpcAttachmentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentRequest) SetResourceOwnerId(v int64) *CreateTransitRouterVpcAttachmentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentRequest) SetTransitRouterAttachmentDescription(v string) *CreateTransitRouterVpcAttachmentRequest {
	s.TransitRouterAttachmentDescription = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentRequest) SetTransitRouterAttachmentName(v string) *CreateTransitRouterVpcAttachmentRequest {
	s.TransitRouterAttachmentName = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentRequest) SetTransitRouterId(v string) *CreateTransitRouterVpcAttachmentRequest {
	s.TransitRouterId = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentRequest) SetVpcId(v string) *CreateTransitRouterVpcAttachmentRequest {
	s.VpcId = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentRequest) SetVpcOwnerId(v int64) *CreateTransitRouterVpcAttachmentRequest {
	s.VpcOwnerId = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentRequest) SetZoneMappings(v []*CreateTransitRouterVpcAttachmentRequestZoneMappings) *CreateTransitRouterVpcAttachmentRequest {
	s.ZoneMappings = v
	return s
}

type CreateTransitRouterVpcAttachmentRequestZoneMappings struct {
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId    *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateTransitRouterVpcAttachmentRequestZoneMappings) String() string {
	return tea.Prettify(s)
}

func (s CreateTransitRouterVpcAttachmentRequestZoneMappings) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterVpcAttachmentRequestZoneMappings) SetVSwitchId(v string) *CreateTransitRouterVpcAttachmentRequestZoneMappings {
	s.VSwitchId = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentRequestZoneMappings) SetZoneId(v string) *CreateTransitRouterVpcAttachmentRequestZoneMappings {
	s.ZoneId = &v
	return s
}

type CreateTransitRouterVpcAttachmentResponseBody struct {
	RequestId                 *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
}

func (s CreateTransitRouterVpcAttachmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTransitRouterVpcAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterVpcAttachmentResponseBody) SetRequestId(v string) *CreateTransitRouterVpcAttachmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentResponseBody) SetTransitRouterAttachmentId(v string) *CreateTransitRouterVpcAttachmentResponseBody {
	s.TransitRouterAttachmentId = &v
	return s
}

type CreateTransitRouterVpcAttachmentResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTransitRouterVpcAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTransitRouterVpcAttachmentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTransitRouterVpcAttachmentResponse) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterVpcAttachmentResponse) SetHeaders(v map[string]*string) *CreateTransitRouterVpcAttachmentResponse {
	s.Headers = v
	return s
}

func (s *CreateTransitRouterVpcAttachmentResponse) SetBody(v *CreateTransitRouterVpcAttachmentResponseBody) *CreateTransitRouterVpcAttachmentResponse {
	s.Body = v
	return s
}

type DeactiveFlowLogRequest struct {
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	FlowLogId            *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeactiveFlowLogRequest) String() string {
	return tea.Prettify(s)
}

func (s DeactiveFlowLogRequest) GoString() string {
	return s.String()
}

func (s *DeactiveFlowLogRequest) SetCenId(v string) *DeactiveFlowLogRequest {
	s.CenId = &v
	return s
}

func (s *DeactiveFlowLogRequest) SetClientToken(v string) *DeactiveFlowLogRequest {
	s.ClientToken = &v
	return s
}

func (s *DeactiveFlowLogRequest) SetFlowLogId(v string) *DeactiveFlowLogRequest {
	s.FlowLogId = &v
	return s
}

func (s *DeactiveFlowLogRequest) SetOwnerAccount(v string) *DeactiveFlowLogRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeactiveFlowLogRequest) SetOwnerId(v int64) *DeactiveFlowLogRequest {
	s.OwnerId = &v
	return s
}

func (s *DeactiveFlowLogRequest) SetRegionId(v string) *DeactiveFlowLogRequest {
	s.RegionId = &v
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
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteCenRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCenRequest) GoString() string {
	return s.String()
}

func (s *DeleteCenRequest) SetCenId(v string) *DeleteCenRequest {
	s.CenId = &v
	return s
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
	CenBandwidthPackageId *string `json:"CenBandwidthPackageId,omitempty" xml:"CenBandwidthPackageId,omitempty"`
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteCenBandwidthPackageRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCenBandwidthPackageRequest) GoString() string {
	return s.String()
}

func (s *DeleteCenBandwidthPackageRequest) SetCenBandwidthPackageId(v string) *DeleteCenBandwidthPackageRequest {
	s.CenBandwidthPackageId = &v
	return s
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

type DeleteCenChildInstanceRouteEntryToAttachmentRequest struct {
	CenId                     *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ClientToken               *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DestinationCidrBlock      *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	DryRun                    *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount              *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount      *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId           *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RouteTableId              *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
}

func (s DeleteCenChildInstanceRouteEntryToAttachmentRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCenChildInstanceRouteEntryToAttachmentRequest) GoString() string {
	return s.String()
}

func (s *DeleteCenChildInstanceRouteEntryToAttachmentRequest) SetCenId(v string) *DeleteCenChildInstanceRouteEntryToAttachmentRequest {
	s.CenId = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToAttachmentRequest) SetClientToken(v string) *DeleteCenChildInstanceRouteEntryToAttachmentRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToAttachmentRequest) SetDestinationCidrBlock(v string) *DeleteCenChildInstanceRouteEntryToAttachmentRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToAttachmentRequest) SetDryRun(v bool) *DeleteCenChildInstanceRouteEntryToAttachmentRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToAttachmentRequest) SetOwnerAccount(v string) *DeleteCenChildInstanceRouteEntryToAttachmentRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToAttachmentRequest) SetOwnerId(v int64) *DeleteCenChildInstanceRouteEntryToAttachmentRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToAttachmentRequest) SetResourceOwnerAccount(v string) *DeleteCenChildInstanceRouteEntryToAttachmentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToAttachmentRequest) SetResourceOwnerId(v int64) *DeleteCenChildInstanceRouteEntryToAttachmentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToAttachmentRequest) SetRouteTableId(v string) *DeleteCenChildInstanceRouteEntryToAttachmentRequest {
	s.RouteTableId = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToAttachmentRequest) SetTransitRouterAttachmentId(v string) *DeleteCenChildInstanceRouteEntryToAttachmentRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

type DeleteCenChildInstanceRouteEntryToAttachmentResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCenChildInstanceRouteEntryToAttachmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCenChildInstanceRouteEntryToAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCenChildInstanceRouteEntryToAttachmentResponseBody) SetRequestId(v string) *DeleteCenChildInstanceRouteEntryToAttachmentResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCenChildInstanceRouteEntryToAttachmentResponse struct {
	Headers map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteCenChildInstanceRouteEntryToAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCenChildInstanceRouteEntryToAttachmentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCenChildInstanceRouteEntryToAttachmentResponse) GoString() string {
	return s.String()
}

func (s *DeleteCenChildInstanceRouteEntryToAttachmentResponse) SetHeaders(v map[string]*string) *DeleteCenChildInstanceRouteEntryToAttachmentResponse {
	s.Headers = v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToAttachmentResponse) SetBody(v *DeleteCenChildInstanceRouteEntryToAttachmentResponseBody) *DeleteCenChildInstanceRouteEntryToAttachmentResponse {
	s.Body = v
	return s
}

type DeleteCenChildInstanceRouteEntryToCenRequest struct {
	CenId                 *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ChildInstanceAliUid   *int64  `json:"ChildInstanceAliUid,omitempty" xml:"ChildInstanceAliUid,omitempty"`
	ChildInstanceId       *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	ChildInstanceRegionId *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	ChildInstanceType     *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	DestinationCidrBlock  *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RouteTableId          *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s DeleteCenChildInstanceRouteEntryToCenRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCenChildInstanceRouteEntryToCenRequest) GoString() string {
	return s.String()
}

func (s *DeleteCenChildInstanceRouteEntryToCenRequest) SetCenId(v string) *DeleteCenChildInstanceRouteEntryToCenRequest {
	s.CenId = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToCenRequest) SetChildInstanceAliUid(v int64) *DeleteCenChildInstanceRouteEntryToCenRequest {
	s.ChildInstanceAliUid = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToCenRequest) SetChildInstanceId(v string) *DeleteCenChildInstanceRouteEntryToCenRequest {
	s.ChildInstanceId = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToCenRequest) SetChildInstanceRegionId(v string) *DeleteCenChildInstanceRouteEntryToCenRequest {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToCenRequest) SetChildInstanceType(v string) *DeleteCenChildInstanceRouteEntryToCenRequest {
	s.ChildInstanceType = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToCenRequest) SetDestinationCidrBlock(v string) *DeleteCenChildInstanceRouteEntryToCenRequest {
	s.DestinationCidrBlock = &v
	return s
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

func (s *DeleteCenChildInstanceRouteEntryToCenRequest) SetRouteTableId(v string) *DeleteCenChildInstanceRouteEntryToCenRequest {
	s.RouteTableId = &v
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

type DeleteCenInterRegionTrafficQosPolicyRequest struct {
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun               *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TrafficQosPolicyId   *string `json:"TrafficQosPolicyId,omitempty" xml:"TrafficQosPolicyId,omitempty"`
}

func (s DeleteCenInterRegionTrafficQosPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCenInterRegionTrafficQosPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteCenInterRegionTrafficQosPolicyRequest) SetClientToken(v string) *DeleteCenInterRegionTrafficQosPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteCenInterRegionTrafficQosPolicyRequest) SetDryRun(v bool) *DeleteCenInterRegionTrafficQosPolicyRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteCenInterRegionTrafficQosPolicyRequest) SetOwnerAccount(v string) *DeleteCenInterRegionTrafficQosPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteCenInterRegionTrafficQosPolicyRequest) SetOwnerId(v int64) *DeleteCenInterRegionTrafficQosPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteCenInterRegionTrafficQosPolicyRequest) SetResourceOwnerAccount(v string) *DeleteCenInterRegionTrafficQosPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteCenInterRegionTrafficQosPolicyRequest) SetResourceOwnerId(v int64) *DeleteCenInterRegionTrafficQosPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteCenInterRegionTrafficQosPolicyRequest) SetTrafficQosPolicyId(v string) *DeleteCenInterRegionTrafficQosPolicyRequest {
	s.TrafficQosPolicyId = &v
	return s
}

type DeleteCenInterRegionTrafficQosPolicyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCenInterRegionTrafficQosPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCenInterRegionTrafficQosPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCenInterRegionTrafficQosPolicyResponseBody) SetRequestId(v string) *DeleteCenInterRegionTrafficQosPolicyResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCenInterRegionTrafficQosPolicyResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteCenInterRegionTrafficQosPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCenInterRegionTrafficQosPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCenInterRegionTrafficQosPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteCenInterRegionTrafficQosPolicyResponse) SetHeaders(v map[string]*string) *DeleteCenInterRegionTrafficQosPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteCenInterRegionTrafficQosPolicyResponse) SetBody(v *DeleteCenInterRegionTrafficQosPolicyResponseBody) *DeleteCenInterRegionTrafficQosPolicyResponse {
	s.Body = v
	return s
}

type DeleteCenInterRegionTrafficQosQueueRequest struct {
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun               *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	QosQueueId           *string `json:"QosQueueId,omitempty" xml:"QosQueueId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteCenInterRegionTrafficQosQueueRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCenInterRegionTrafficQosQueueRequest) GoString() string {
	return s.String()
}

func (s *DeleteCenInterRegionTrafficQosQueueRequest) SetClientToken(v string) *DeleteCenInterRegionTrafficQosQueueRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteCenInterRegionTrafficQosQueueRequest) SetDryRun(v bool) *DeleteCenInterRegionTrafficQosQueueRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteCenInterRegionTrafficQosQueueRequest) SetOwnerAccount(v string) *DeleteCenInterRegionTrafficQosQueueRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteCenInterRegionTrafficQosQueueRequest) SetOwnerId(v int64) *DeleteCenInterRegionTrafficQosQueueRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteCenInterRegionTrafficQosQueueRequest) SetQosQueueId(v string) *DeleteCenInterRegionTrafficQosQueueRequest {
	s.QosQueueId = &v
	return s
}

func (s *DeleteCenInterRegionTrafficQosQueueRequest) SetResourceOwnerAccount(v string) *DeleteCenInterRegionTrafficQosQueueRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteCenInterRegionTrafficQosQueueRequest) SetResourceOwnerId(v int64) *DeleteCenInterRegionTrafficQosQueueRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteCenInterRegionTrafficQosQueueResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCenInterRegionTrafficQosQueueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCenInterRegionTrafficQosQueueResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCenInterRegionTrafficQosQueueResponseBody) SetRequestId(v string) *DeleteCenInterRegionTrafficQosQueueResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCenInterRegionTrafficQosQueueResponse struct {
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteCenInterRegionTrafficQosQueueResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCenInterRegionTrafficQosQueueResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCenInterRegionTrafficQosQueueResponse) GoString() string {
	return s.String()
}

func (s *DeleteCenInterRegionTrafficQosQueueResponse) SetHeaders(v map[string]*string) *DeleteCenInterRegionTrafficQosQueueResponse {
	s.Headers = v
	return s
}

func (s *DeleteCenInterRegionTrafficQosQueueResponse) SetBody(v *DeleteCenInterRegionTrafficQosQueueResponseBody) *DeleteCenInterRegionTrafficQosQueueResponse {
	s.Body = v
	return s
}

type DeleteCenRouteMapRequest struct {
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	CenRegionId          *string `json:"CenRegionId,omitempty" xml:"CenRegionId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RouteMapId           *string `json:"RouteMapId,omitempty" xml:"RouteMapId,omitempty"`
}

func (s DeleteCenRouteMapRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCenRouteMapRequest) GoString() string {
	return s.String()
}

func (s *DeleteCenRouteMapRequest) SetCenId(v string) *DeleteCenRouteMapRequest {
	s.CenId = &v
	return s
}

func (s *DeleteCenRouteMapRequest) SetCenRegionId(v string) *DeleteCenRouteMapRequest {
	s.CenRegionId = &v
	return s
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
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	FlowLogId            *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteFlowlogRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowlogRequest) GoString() string {
	return s.String()
}

func (s *DeleteFlowlogRequest) SetCenId(v string) *DeleteFlowlogRequest {
	s.CenId = &v
	return s
}

func (s *DeleteFlowlogRequest) SetClientToken(v string) *DeleteFlowlogRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteFlowlogRequest) SetFlowLogId(v string) *DeleteFlowlogRequest {
	s.FlowLogId = &v
	return s
}

func (s *DeleteFlowlogRequest) SetOwnerAccount(v string) *DeleteFlowlogRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteFlowlogRequest) SetOwnerId(v int64) *DeleteFlowlogRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteFlowlogRequest) SetRegionId(v string) *DeleteFlowlogRequest {
	s.RegionId = &v
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
	AccessRegionId       *string `json:"AccessRegionId,omitempty" xml:"AccessRegionId,omitempty"`
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	Host                 *string `json:"Host,omitempty" xml:"Host,omitempty"`
	HostRegionId         *string `json:"HostRegionId,omitempty" xml:"HostRegionId,omitempty"`
	HostVpcId            *string `json:"HostVpcId,omitempty" xml:"HostVpcId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteRouteServiceInCenRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRouteServiceInCenRequest) GoString() string {
	return s.String()
}

func (s *DeleteRouteServiceInCenRequest) SetAccessRegionId(v string) *DeleteRouteServiceInCenRequest {
	s.AccessRegionId = &v
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

func (s *DeleteRouteServiceInCenRequest) SetHostVpcId(v string) *DeleteRouteServiceInCenRequest {
	s.HostVpcId = &v
	return s
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

type DeleteTrafficMarkingPolicyRequest struct {
	ClientToken            *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                 *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount           *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount   *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId        *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TrafficMarkingPolicyId *string `json:"TrafficMarkingPolicyId,omitempty" xml:"TrafficMarkingPolicyId,omitempty"`
}

func (s DeleteTrafficMarkingPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTrafficMarkingPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteTrafficMarkingPolicyRequest) SetClientToken(v string) *DeleteTrafficMarkingPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteTrafficMarkingPolicyRequest) SetDryRun(v bool) *DeleteTrafficMarkingPolicyRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteTrafficMarkingPolicyRequest) SetOwnerAccount(v string) *DeleteTrafficMarkingPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteTrafficMarkingPolicyRequest) SetOwnerId(v int64) *DeleteTrafficMarkingPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteTrafficMarkingPolicyRequest) SetResourceOwnerAccount(v string) *DeleteTrafficMarkingPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteTrafficMarkingPolicyRequest) SetResourceOwnerId(v int64) *DeleteTrafficMarkingPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteTrafficMarkingPolicyRequest) SetTrafficMarkingPolicyId(v string) *DeleteTrafficMarkingPolicyRequest {
	s.TrafficMarkingPolicyId = &v
	return s
}

type DeleteTrafficMarkingPolicyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTrafficMarkingPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTrafficMarkingPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTrafficMarkingPolicyResponseBody) SetRequestId(v string) *DeleteTrafficMarkingPolicyResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTrafficMarkingPolicyResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteTrafficMarkingPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTrafficMarkingPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTrafficMarkingPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteTrafficMarkingPolicyResponse) SetHeaders(v map[string]*string) *DeleteTrafficMarkingPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteTrafficMarkingPolicyResponse) SetBody(v *DeleteTrafficMarkingPolicyResponseBody) *DeleteTrafficMarkingPolicyResponse {
	s.Body = v
	return s
}

type DeleteTransitRouterPeerAttachmentRequest struct {
	ClientToken               *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                    *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount              *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount      *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId           *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
}

func (s DeleteTransitRouterPeerAttachmentRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTransitRouterPeerAttachmentRequest) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterPeerAttachmentRequest) SetClientToken(v string) *DeleteTransitRouterPeerAttachmentRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteTransitRouterPeerAttachmentRequest) SetDryRun(v bool) *DeleteTransitRouterPeerAttachmentRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteTransitRouterPeerAttachmentRequest) SetOwnerAccount(v string) *DeleteTransitRouterPeerAttachmentRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteTransitRouterPeerAttachmentRequest) SetOwnerId(v int64) *DeleteTransitRouterPeerAttachmentRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteTransitRouterPeerAttachmentRequest) SetResourceOwnerAccount(v string) *DeleteTransitRouterPeerAttachmentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteTransitRouterPeerAttachmentRequest) SetResourceOwnerId(v int64) *DeleteTransitRouterPeerAttachmentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteTransitRouterPeerAttachmentRequest) SetTransitRouterAttachmentId(v string) *DeleteTransitRouterPeerAttachmentRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

type DeleteTransitRouterPeerAttachmentResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTransitRouterPeerAttachmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTransitRouterPeerAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterPeerAttachmentResponseBody) SetRequestId(v string) *DeleteTransitRouterPeerAttachmentResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTransitRouterPeerAttachmentResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteTransitRouterPeerAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTransitRouterPeerAttachmentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTransitRouterPeerAttachmentResponse) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterPeerAttachmentResponse) SetHeaders(v map[string]*string) *DeleteTransitRouterPeerAttachmentResponse {
	s.Headers = v
	return s
}

func (s *DeleteTransitRouterPeerAttachmentResponse) SetBody(v *DeleteTransitRouterPeerAttachmentResponseBody) *DeleteTransitRouterPeerAttachmentResponse {
	s.Body = v
	return s
}

type DeleteTransitRouterRouteEntryRequest struct {
	ClientToken                                 *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                                      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount                                *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                                     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount                        *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId                             *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TransitRouterRouteEntryDestinationCidrBlock *string `json:"TransitRouterRouteEntryDestinationCidrBlock,omitempty" xml:"TransitRouterRouteEntryDestinationCidrBlock,omitempty"`
	TransitRouterRouteEntryId                   *string `json:"TransitRouterRouteEntryId,omitempty" xml:"TransitRouterRouteEntryId,omitempty"`
	TransitRouterRouteEntryNextHopId            *string `json:"TransitRouterRouteEntryNextHopId,omitempty" xml:"TransitRouterRouteEntryNextHopId,omitempty"`
	TransitRouterRouteEntryNextHopType          *string `json:"TransitRouterRouteEntryNextHopType,omitempty" xml:"TransitRouterRouteEntryNextHopType,omitempty"`
	TransitRouterRouteTableId                   *string `json:"TransitRouterRouteTableId,omitempty" xml:"TransitRouterRouteTableId,omitempty"`
}

func (s DeleteTransitRouterRouteEntryRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTransitRouterRouteEntryRequest) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterRouteEntryRequest) SetClientToken(v string) *DeleteTransitRouterRouteEntryRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteTransitRouterRouteEntryRequest) SetDryRun(v bool) *DeleteTransitRouterRouteEntryRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteTransitRouterRouteEntryRequest) SetOwnerAccount(v string) *DeleteTransitRouterRouteEntryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteTransitRouterRouteEntryRequest) SetOwnerId(v int64) *DeleteTransitRouterRouteEntryRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteTransitRouterRouteEntryRequest) SetResourceOwnerAccount(v string) *DeleteTransitRouterRouteEntryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteTransitRouterRouteEntryRequest) SetResourceOwnerId(v int64) *DeleteTransitRouterRouteEntryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteTransitRouterRouteEntryRequest) SetTransitRouterRouteEntryDestinationCidrBlock(v string) *DeleteTransitRouterRouteEntryRequest {
	s.TransitRouterRouteEntryDestinationCidrBlock = &v
	return s
}

func (s *DeleteTransitRouterRouteEntryRequest) SetTransitRouterRouteEntryId(v string) *DeleteTransitRouterRouteEntryRequest {
	s.TransitRouterRouteEntryId = &v
	return s
}

func (s *DeleteTransitRouterRouteEntryRequest) SetTransitRouterRouteEntryNextHopId(v string) *DeleteTransitRouterRouteEntryRequest {
	s.TransitRouterRouteEntryNextHopId = &v
	return s
}

func (s *DeleteTransitRouterRouteEntryRequest) SetTransitRouterRouteEntryNextHopType(v string) *DeleteTransitRouterRouteEntryRequest {
	s.TransitRouterRouteEntryNextHopType = &v
	return s
}

func (s *DeleteTransitRouterRouteEntryRequest) SetTransitRouterRouteTableId(v string) *DeleteTransitRouterRouteEntryRequest {
	s.TransitRouterRouteTableId = &v
	return s
}

type DeleteTransitRouterRouteEntryResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTransitRouterRouteEntryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTransitRouterRouteEntryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterRouteEntryResponseBody) SetRequestId(v string) *DeleteTransitRouterRouteEntryResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTransitRouterRouteEntryResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteTransitRouterRouteEntryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTransitRouterRouteEntryResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTransitRouterRouteEntryResponse) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterRouteEntryResponse) SetHeaders(v map[string]*string) *DeleteTransitRouterRouteEntryResponse {
	s.Headers = v
	return s
}

func (s *DeleteTransitRouterRouteEntryResponse) SetBody(v *DeleteTransitRouterRouteEntryResponseBody) *DeleteTransitRouterRouteEntryResponse {
	s.Body = v
	return s
}

type DeleteTransitRouterRouteTableRequest struct {
	ClientToken               *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                    *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount              *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount      *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId           *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TransitRouterRouteTableId *string `json:"TransitRouterRouteTableId,omitempty" xml:"TransitRouterRouteTableId,omitempty"`
}

func (s DeleteTransitRouterRouteTableRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTransitRouterRouteTableRequest) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterRouteTableRequest) SetClientToken(v string) *DeleteTransitRouterRouteTableRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteTransitRouterRouteTableRequest) SetDryRun(v bool) *DeleteTransitRouterRouteTableRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteTransitRouterRouteTableRequest) SetOwnerAccount(v string) *DeleteTransitRouterRouteTableRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteTransitRouterRouteTableRequest) SetOwnerId(v int64) *DeleteTransitRouterRouteTableRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteTransitRouterRouteTableRequest) SetResourceOwnerAccount(v string) *DeleteTransitRouterRouteTableRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteTransitRouterRouteTableRequest) SetResourceOwnerId(v int64) *DeleteTransitRouterRouteTableRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteTransitRouterRouteTableRequest) SetTransitRouterRouteTableId(v string) *DeleteTransitRouterRouteTableRequest {
	s.TransitRouterRouteTableId = &v
	return s
}

type DeleteTransitRouterRouteTableResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTransitRouterRouteTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTransitRouterRouteTableResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterRouteTableResponseBody) SetRequestId(v string) *DeleteTransitRouterRouteTableResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTransitRouterRouteTableResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteTransitRouterRouteTableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTransitRouterRouteTableResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTransitRouterRouteTableResponse) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterRouteTableResponse) SetHeaders(v map[string]*string) *DeleteTransitRouterRouteTableResponse {
	s.Headers = v
	return s
}

func (s *DeleteTransitRouterRouteTableResponse) SetBody(v *DeleteTransitRouterRouteTableResponseBody) *DeleteTransitRouterRouteTableResponse {
	s.Body = v
	return s
}

type DeleteTransitRouterVbrAttachmentRequest struct {
	ClientToken               *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                    *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount              *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount      *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId           *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
}

func (s DeleteTransitRouterVbrAttachmentRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTransitRouterVbrAttachmentRequest) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterVbrAttachmentRequest) SetClientToken(v string) *DeleteTransitRouterVbrAttachmentRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteTransitRouterVbrAttachmentRequest) SetDryRun(v bool) *DeleteTransitRouterVbrAttachmentRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteTransitRouterVbrAttachmentRequest) SetOwnerAccount(v string) *DeleteTransitRouterVbrAttachmentRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteTransitRouterVbrAttachmentRequest) SetOwnerId(v int64) *DeleteTransitRouterVbrAttachmentRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteTransitRouterVbrAttachmentRequest) SetResourceOwnerAccount(v string) *DeleteTransitRouterVbrAttachmentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteTransitRouterVbrAttachmentRequest) SetResourceOwnerId(v int64) *DeleteTransitRouterVbrAttachmentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteTransitRouterVbrAttachmentRequest) SetTransitRouterAttachmentId(v string) *DeleteTransitRouterVbrAttachmentRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

type DeleteTransitRouterVbrAttachmentResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTransitRouterVbrAttachmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTransitRouterVbrAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterVbrAttachmentResponseBody) SetRequestId(v string) *DeleteTransitRouterVbrAttachmentResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTransitRouterVbrAttachmentResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteTransitRouterVbrAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTransitRouterVbrAttachmentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTransitRouterVbrAttachmentResponse) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterVbrAttachmentResponse) SetHeaders(v map[string]*string) *DeleteTransitRouterVbrAttachmentResponse {
	s.Headers = v
	return s
}

func (s *DeleteTransitRouterVbrAttachmentResponse) SetBody(v *DeleteTransitRouterVbrAttachmentResponseBody) *DeleteTransitRouterVbrAttachmentResponse {
	s.Body = v
	return s
}

type DeleteTransitRouterVpcAttachmentRequest struct {
	ClientToken               *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                    *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount              *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount      *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId           *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
}

func (s DeleteTransitRouterVpcAttachmentRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTransitRouterVpcAttachmentRequest) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterVpcAttachmentRequest) SetClientToken(v string) *DeleteTransitRouterVpcAttachmentRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteTransitRouterVpcAttachmentRequest) SetDryRun(v bool) *DeleteTransitRouterVpcAttachmentRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteTransitRouterVpcAttachmentRequest) SetOwnerAccount(v string) *DeleteTransitRouterVpcAttachmentRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteTransitRouterVpcAttachmentRequest) SetOwnerId(v int64) *DeleteTransitRouterVpcAttachmentRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteTransitRouterVpcAttachmentRequest) SetResourceOwnerAccount(v string) *DeleteTransitRouterVpcAttachmentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteTransitRouterVpcAttachmentRequest) SetResourceOwnerId(v int64) *DeleteTransitRouterVpcAttachmentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteTransitRouterVpcAttachmentRequest) SetTransitRouterAttachmentId(v string) *DeleteTransitRouterVpcAttachmentRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

type DeleteTransitRouterVpcAttachmentResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTransitRouterVpcAttachmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTransitRouterVpcAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterVpcAttachmentResponseBody) SetRequestId(v string) *DeleteTransitRouterVpcAttachmentResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTransitRouterVpcAttachmentResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteTransitRouterVpcAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTransitRouterVpcAttachmentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTransitRouterVpcAttachmentResponse) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterVpcAttachmentResponse) SetHeaders(v map[string]*string) *DeleteTransitRouterVpcAttachmentResponse {
	s.Headers = v
	return s
}

func (s *DeleteTransitRouterVpcAttachmentResponse) SetBody(v *DeleteTransitRouterVpcAttachmentResponseBody) *DeleteTransitRouterVpcAttachmentResponse {
	s.Body = v
	return s
}

type DescribeCenAttachedChildInstanceAttributeRequest struct {
	CenId                 *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ChildInstanceId       *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	ChildInstanceRegionId *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	ChildInstanceType     *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeCenAttachedChildInstanceAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenAttachedChildInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeCenAttachedChildInstanceAttributeRequest) SetCenId(v string) *DescribeCenAttachedChildInstanceAttributeRequest {
	s.CenId = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeRequest) SetChildInstanceId(v string) *DescribeCenAttachedChildInstanceAttributeRequest {
	s.ChildInstanceId = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeRequest) SetChildInstanceRegionId(v string) *DescribeCenAttachedChildInstanceAttributeRequest {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeRequest) SetChildInstanceType(v string) *DescribeCenAttachedChildInstanceAttributeRequest {
	s.ChildInstanceType = &v
	return s
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

type DescribeCenAttachedChildInstanceAttributeResponseBody struct {
	CenId                   *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ChildInstanceAttachTime *string `json:"ChildInstanceAttachTime,omitempty" xml:"ChildInstanceAttachTime,omitempty"`
	ChildInstanceId         *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	ChildInstanceName       *string `json:"ChildInstanceName,omitempty" xml:"ChildInstanceName,omitempty"`
	ChildInstanceOwnerId    *int64  `json:"ChildInstanceOwnerId,omitempty" xml:"ChildInstanceOwnerId,omitempty"`
	ChildInstanceRegionId   *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	ChildInstanceType       *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	RequestId               *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status                  *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCenAttachedChildInstanceAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenAttachedChildInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) SetCenId(v string) *DescribeCenAttachedChildInstanceAttributeResponseBody {
	s.CenId = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) SetChildInstanceAttachTime(v string) *DescribeCenAttachedChildInstanceAttributeResponseBody {
	s.ChildInstanceAttachTime = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) SetChildInstanceId(v string) *DescribeCenAttachedChildInstanceAttributeResponseBody {
	s.ChildInstanceId = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) SetChildInstanceName(v string) *DescribeCenAttachedChildInstanceAttributeResponseBody {
	s.ChildInstanceName = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) SetChildInstanceOwnerId(v int64) *DescribeCenAttachedChildInstanceAttributeResponseBody {
	s.ChildInstanceOwnerId = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) SetChildInstanceRegionId(v string) *DescribeCenAttachedChildInstanceAttributeResponseBody {
	s.ChildInstanceRegionId = &v
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

func (s *DescribeCenAttachedChildInstanceAttributeResponseBody) SetStatus(v string) *DescribeCenAttachedChildInstanceAttributeResponseBody {
	s.Status = &v
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
	CenId                 *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ChildInstanceRegionId *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	ChildInstanceType     *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber            *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize              *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeCenAttachedChildInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenAttachedChildInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCenAttachedChildInstancesRequest) SetCenId(v string) *DescribeCenAttachedChildInstancesRequest {
	s.CenId = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesRequest) SetChildInstanceRegionId(v string) *DescribeCenAttachedChildInstancesRequest {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesRequest) SetChildInstanceType(v string) *DescribeCenAttachedChildInstancesRequest {
	s.ChildInstanceType = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesRequest) SetOwnerAccount(v string) *DescribeCenAttachedChildInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesRequest) SetOwnerId(v int64) *DescribeCenAttachedChildInstancesRequest {
	s.OwnerId = &v
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

func (s *DescribeCenAttachedChildInstancesRequest) SetResourceOwnerAccount(v string) *DescribeCenAttachedChildInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesRequest) SetResourceOwnerId(v int64) *DescribeCenAttachedChildInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeCenAttachedChildInstancesResponseBody struct {
	ChildInstances *DescribeCenAttachedChildInstancesResponseBodyChildInstances `json:"ChildInstances,omitempty" xml:"ChildInstances,omitempty" type:"Struct"`
	PageNumber     *int32                                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount     *int32                                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCenAttachedChildInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenAttachedChildInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCenAttachedChildInstancesResponseBody) SetChildInstances(v *DescribeCenAttachedChildInstancesResponseBodyChildInstances) *DescribeCenAttachedChildInstancesResponseBody {
	s.ChildInstances = v
	return s
}

func (s *DescribeCenAttachedChildInstancesResponseBody) SetPageNumber(v int32) *DescribeCenAttachedChildInstancesResponseBody {
	s.PageNumber = &v
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

func (s *DescribeCenAttachedChildInstancesResponseBody) SetTotalCount(v int32) *DescribeCenAttachedChildInstancesResponseBody {
	s.TotalCount = &v
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
	CenId                   *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ChildInstanceAttachTime *string `json:"ChildInstanceAttachTime,omitempty" xml:"ChildInstanceAttachTime,omitempty"`
	ChildInstanceId         *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	ChildInstanceOwnerId    *int64  `json:"ChildInstanceOwnerId,omitempty" xml:"ChildInstanceOwnerId,omitempty"`
	ChildInstanceRegionId   *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	ChildInstanceType       *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	Status                  *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance) GoString() string {
	return s.String()
}

func (s *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance) SetCenId(v string) *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance {
	s.CenId = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance) SetChildInstanceAttachTime(v string) *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance {
	s.ChildInstanceAttachTime = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance) SetChildInstanceId(v string) *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance {
	s.ChildInstanceId = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance) SetChildInstanceOwnerId(v int64) *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance {
	s.ChildInstanceOwnerId = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance) SetChildInstanceRegionId(v string) *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance) SetChildInstanceType(v string) *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance {
	s.ChildInstanceType = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance) SetStatus(v string) *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance {
	s.Status = &v
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
	Filter                 []*DescribeCenBandwidthPackagesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	IncludeReservationData *bool                                        `json:"IncludeReservationData,omitempty" xml:"IncludeReservationData,omitempty"`
	IsOrKey                *bool                                        `json:"IsOrKey,omitempty" xml:"IsOrKey,omitempty"`
	OwnerAccount           *string                                      `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                *int64                                       `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber             *int32                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize               *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount   *string                                      `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId        *int64                                       `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeCenBandwidthPackagesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenBandwidthPackagesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCenBandwidthPackagesRequest) SetFilter(v []*DescribeCenBandwidthPackagesRequestFilter) *DescribeCenBandwidthPackagesRequest {
	s.Filter = v
	return s
}

func (s *DescribeCenBandwidthPackagesRequest) SetIncludeReservationData(v bool) *DescribeCenBandwidthPackagesRequest {
	s.IncludeReservationData = &v
	return s
}

func (s *DescribeCenBandwidthPackagesRequest) SetIsOrKey(v bool) *DescribeCenBandwidthPackagesRequest {
	s.IsOrKey = &v
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

func (s *DescribeCenBandwidthPackagesRequest) SetPageNumber(v int32) *DescribeCenBandwidthPackagesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCenBandwidthPackagesRequest) SetPageSize(v int32) *DescribeCenBandwidthPackagesRequest {
	s.PageSize = &v
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
	CenBandwidthPackages *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages `json:"CenBandwidthPackages,omitempty" xml:"CenBandwidthPackages,omitempty" type:"Struct"`
	PageNumber           *int32                                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32                                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId            *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount           *int32                                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCenBandwidthPackagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenBandwidthPackagesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCenBandwidthPackagesResponseBody) SetCenBandwidthPackages(v *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages) *DescribeCenBandwidthPackagesResponseBody {
	s.CenBandwidthPackages = v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBody) SetPageNumber(v int32) *DescribeCenBandwidthPackagesResponseBody {
	s.PageNumber = &v
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

func (s *DescribeCenBandwidthPackagesResponseBody) SetTotalCount(v int32) *DescribeCenBandwidthPackagesResponseBody {
	s.TotalCount = &v
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
	Bandwidth                       *int64                                                                                                          `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	BandwidthPackageChargeType      *string                                                                                                         `json:"BandwidthPackageChargeType,omitempty" xml:"BandwidthPackageChargeType,omitempty"`
	BusinessStatus                  *string                                                                                                         `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	CenBandwidthPackageId           *string                                                                                                         `json:"CenBandwidthPackageId,omitempty" xml:"CenBandwidthPackageId,omitempty"`
	CenIds                          *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageCenIds                          `json:"CenIds,omitempty" xml:"CenIds,omitempty" type:"Struct"`
	CreationTime                    *string                                                                                                         `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description                     *string                                                                                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	ExpiredTime                     *string                                                                                                         `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	GeographicRegionAId             *string                                                                                                         `json:"GeographicRegionAId,omitempty" xml:"GeographicRegionAId,omitempty"`
	GeographicRegionBId             *string                                                                                                         `json:"GeographicRegionBId,omitempty" xml:"GeographicRegionBId,omitempty"`
	GeographicSpanId                *string                                                                                                         `json:"GeographicSpanId,omitempty" xml:"GeographicSpanId,omitempty"`
	HasReservationData              *string                                                                                                         `json:"HasReservationData,omitempty" xml:"HasReservationData,omitempty"`
	IsCrossBorder                   *bool                                                                                                           `json:"IsCrossBorder,omitempty" xml:"IsCrossBorder,omitempty"`
	Name                            *string                                                                                                         `json:"Name,omitempty" xml:"Name,omitempty"`
	OrginInterRegionBandwidthLimits *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimits `json:"OrginInterRegionBandwidthLimits,omitempty" xml:"OrginInterRegionBandwidthLimits,omitempty" type:"Struct"`
	ReservationActiveTime           *string                                                                                                         `json:"ReservationActiveTime,omitempty" xml:"ReservationActiveTime,omitempty"`
	ReservationBandwidth            *string                                                                                                         `json:"ReservationBandwidth,omitempty" xml:"ReservationBandwidth,omitempty"`
	ReservationInternetChargeType   *string                                                                                                         `json:"ReservationInternetChargeType,omitempty" xml:"ReservationInternetChargeType,omitempty"`
	ReservationOrderType            *string                                                                                                         `json:"ReservationOrderType,omitempty" xml:"ReservationOrderType,omitempty"`
	Status                          *string                                                                                                         `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) GoString() string {
	return s.String()
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetBandwidth(v int64) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.Bandwidth = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetBandwidthPackageChargeType(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.BandwidthPackageChargeType = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetBusinessStatus(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.BusinessStatus = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetCenBandwidthPackageId(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.CenBandwidthPackageId = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetCenIds(v *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageCenIds) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.CenIds = v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetCreationTime(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.CreationTime = &v
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

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetGeographicRegionAId(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.GeographicRegionAId = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetGeographicRegionBId(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.GeographicRegionBId = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetGeographicSpanId(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.GeographicSpanId = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetHasReservationData(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.HasReservationData = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetIsCrossBorder(v bool) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.IsCrossBorder = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetName(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.Name = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetOrginInterRegionBandwidthLimits(v *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimits) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.OrginInterRegionBandwidthLimits = v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetReservationActiveTime(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.ReservationActiveTime = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetReservationBandwidth(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.ReservationBandwidth = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetReservationInternetChargeType(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.ReservationInternetChargeType = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetReservationOrderType(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.ReservationOrderType = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetStatus(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.Status = &v
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
	GeographicSpanId *string `json:"GeographicSpanId,omitempty" xml:"GeographicSpanId,omitempty"`
	LocalRegionId    *string `json:"LocalRegionId,omitempty" xml:"LocalRegionId,omitempty"`
	OppositeRegionId *string `json:"OppositeRegionId,omitempty" xml:"OppositeRegionId,omitempty"`
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

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimitsOrginInterRegionBandwidthLimit) SetGeographicSpanId(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimitsOrginInterRegionBandwidthLimit {
	s.GeographicSpanId = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimitsOrginInterRegionBandwidthLimit) SetLocalRegionId(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimitsOrginInterRegionBandwidthLimit {
	s.LocalRegionId = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimitsOrginInterRegionBandwidthLimit) SetOppositeRegionId(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimitsOrginInterRegionBandwidthLimit {
	s.OppositeRegionId = &v
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
	CenId                 *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ChildInstanceId       *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	ChildInstanceRegionId *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	ChildInstanceType     *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber            *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize              *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCenChildInstanceRouteEntriesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenChildInstanceRouteEntriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) SetCenId(v string) *DescribeCenChildInstanceRouteEntriesRequest {
	s.CenId = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) SetChildInstanceId(v string) *DescribeCenChildInstanceRouteEntriesRequest {
	s.ChildInstanceId = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) SetChildInstanceRegionId(v string) *DescribeCenChildInstanceRouteEntriesRequest {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) SetChildInstanceType(v string) *DescribeCenChildInstanceRouteEntriesRequest {
	s.ChildInstanceType = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) SetOwnerAccount(v string) *DescribeCenChildInstanceRouteEntriesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) SetOwnerId(v int64) *DescribeCenChildInstanceRouteEntriesRequest {
	s.OwnerId = &v
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

func (s *DescribeCenChildInstanceRouteEntriesRequest) SetResourceOwnerAccount(v string) *DescribeCenChildInstanceRouteEntriesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) SetResourceOwnerId(v int64) *DescribeCenChildInstanceRouteEntriesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesRequest) SetStatus(v string) *DescribeCenChildInstanceRouteEntriesRequest {
	s.Status = &v
	return s
}

type DescribeCenChildInstanceRouteEntriesResponseBody struct {
	CenRouteEntries *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries `json:"CenRouteEntries,omitempty" xml:"CenRouteEntries,omitempty" type:"Struct"`
	PageNumber      *int32                                                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32                                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId       *string                                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount      *int32                                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCenChildInstanceRouteEntriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenChildInstanceRouteEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBody) SetCenRouteEntries(v *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries) *DescribeCenChildInstanceRouteEntriesResponseBody {
	s.CenRouteEntries = v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBody) SetPageNumber(v int32) *DescribeCenChildInstanceRouteEntriesResponseBody {
	s.PageNumber = &v
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

func (s *DescribeCenChildInstanceRouteEntriesResponseBody) SetTotalCount(v int32) *DescribeCenChildInstanceRouteEntriesResponseBody {
	s.TotalCount = &v
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
	AsPaths              *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryAsPaths            `json:"AsPaths,omitempty" xml:"AsPaths,omitempty" type:"Struct"`
	CenRouteMapRecords   *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecords `json:"CenRouteMapRecords,omitempty" xml:"CenRouteMapRecords,omitempty" type:"Struct"`
	Communities          *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCommunities        `json:"Communities,omitempty" xml:"Communities,omitempty" type:"Struct"`
	Conflicts            *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflicts          `json:"Conflicts,omitempty" xml:"Conflicts,omitempty" type:"Struct"`
	DestinationCidrBlock *string                                                                                         `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	NextHopInstanceId    *string                                                                                         `json:"NextHopInstanceId,omitempty" xml:"NextHopInstanceId,omitempty"`
	NextHopRegionId      *string                                                                                         `json:"NextHopRegionId,omitempty" xml:"NextHopRegionId,omitempty"`
	NextHopType          *string                                                                                         `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	OperationalMode      *bool                                                                                           `json:"OperationalMode,omitempty" xml:"OperationalMode,omitempty"`
	PublishStatus        *string                                                                                         `json:"PublishStatus,omitempty" xml:"PublishStatus,omitempty"`
	RouteTableId         *string                                                                                         `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
	Status               *string                                                                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	Type                 *string                                                                                         `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) GoString() string {
	return s.String()
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetAsPaths(v *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryAsPaths) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.AsPaths = v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetCenRouteMapRecords(v *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecords) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.CenRouteMapRecords = v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetCommunities(v *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCommunities) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.Communities = v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetConflicts(v *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflicts) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.Conflicts = v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetDestinationCidrBlock(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetNextHopInstanceId(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.NextHopInstanceId = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetNextHopRegionId(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.NextHopRegionId = &v
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

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetPublishStatus(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.PublishStatus = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetRouteTableId(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.RouteTableId = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetStatus(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.Status = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetType(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.Type = &v
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
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RouteMapId *string `json:"RouteMapId,omitempty" xml:"RouteMapId,omitempty"`
}

func (s DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord) GoString() string {
	return s.String()
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord) SetRegionId(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord {
	s.RegionId = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord) SetRouteMapId(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord {
	s.RouteMapId = &v
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
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType         *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflictsConflict) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflictsConflict) GoString() string {
	return s.String()
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

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflictsConflict) SetStatus(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflictsConflict {
	s.Status = &v
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
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	GeographicRegionAId  *string `json:"GeographicRegionAId,omitempty" xml:"GeographicRegionAId,omitempty"`
	GeographicRegionBId  *string `json:"GeographicRegionBId,omitempty" xml:"GeographicRegionBId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeCenGeographicSpanRemainingBandwidthRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenGeographicSpanRemainingBandwidthRequest) GoString() string {
	return s.String()
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

func (s *DescribeCenGeographicSpanRemainingBandwidthRequest) SetOwnerAccount(v string) *DescribeCenGeographicSpanRemainingBandwidthRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCenGeographicSpanRemainingBandwidthRequest) SetOwnerId(v int64) *DescribeCenGeographicSpanRemainingBandwidthRequest {
	s.OwnerId = &v
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

func (s *DescribeCenGeographicSpanRemainingBandwidthRequest) SetResourceOwnerAccount(v string) *DescribeCenGeographicSpanRemainingBandwidthRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCenGeographicSpanRemainingBandwidthRequest) SetResourceOwnerId(v int64) *DescribeCenGeographicSpanRemainingBandwidthRequest {
	s.ResourceOwnerId = &v
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
	GeographicSpanId     *string `json:"GeographicSpanId,omitempty" xml:"GeographicSpanId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeCenGeographicSpansRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenGeographicSpansRequest) GoString() string {
	return s.String()
}

func (s *DescribeCenGeographicSpansRequest) SetGeographicSpanId(v string) *DescribeCenGeographicSpansRequest {
	s.GeographicSpanId = &v
	return s
}

func (s *DescribeCenGeographicSpansRequest) SetOwnerAccount(v string) *DescribeCenGeographicSpansRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCenGeographicSpansRequest) SetOwnerId(v int64) *DescribeCenGeographicSpansRequest {
	s.OwnerId = &v
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

func (s *DescribeCenGeographicSpansRequest) SetResourceOwnerAccount(v string) *DescribeCenGeographicSpansRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCenGeographicSpansRequest) SetResourceOwnerId(v int64) *DescribeCenGeographicSpansRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeCenGeographicSpansResponseBody struct {
	GeographicSpanModels *DescribeCenGeographicSpansResponseBodyGeographicSpanModels `json:"GeographicSpanModels,omitempty" xml:"GeographicSpanModels,omitempty" type:"Struct"`
	PageNumber           *int32                                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32                                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId            *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount           *int32                                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCenGeographicSpansResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenGeographicSpansResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCenGeographicSpansResponseBody) SetGeographicSpanModels(v *DescribeCenGeographicSpansResponseBodyGeographicSpanModels) *DescribeCenGeographicSpansResponseBody {
	s.GeographicSpanModels = v
	return s
}

func (s *DescribeCenGeographicSpansResponseBody) SetPageNumber(v int32) *DescribeCenGeographicSpansResponseBody {
	s.PageNumber = &v
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

func (s *DescribeCenGeographicSpansResponseBody) SetTotalCount(v int32) *DescribeCenGeographicSpansResponseBody {
	s.TotalCount = &v
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
	GeographicSpanId    *string `json:"GeographicSpanId,omitempty" xml:"GeographicSpanId,omitempty"`
	LocalGeoRegionId    *string `json:"LocalGeoRegionId,omitempty" xml:"LocalGeoRegionId,omitempty"`
	OppositeGeoRegionId *string `json:"OppositeGeoRegionId,omitempty" xml:"OppositeGeoRegionId,omitempty"`
}

func (s DescribeCenGeographicSpansResponseBodyGeographicSpanModelsGeographicSpanModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenGeographicSpansResponseBodyGeographicSpanModelsGeographicSpanModel) GoString() string {
	return s.String()
}

func (s *DescribeCenGeographicSpansResponseBodyGeographicSpanModelsGeographicSpanModel) SetGeographicSpanId(v string) *DescribeCenGeographicSpansResponseBodyGeographicSpanModelsGeographicSpanModel {
	s.GeographicSpanId = &v
	return s
}

func (s *DescribeCenGeographicSpansResponseBodyGeographicSpanModelsGeographicSpanModel) SetLocalGeoRegionId(v string) *DescribeCenGeographicSpansResponseBodyGeographicSpanModelsGeographicSpanModel {
	s.LocalGeoRegionId = &v
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
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeCenInterRegionBandwidthLimitsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenInterRegionBandwidthLimitsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCenInterRegionBandwidthLimitsRequest) SetCenId(v string) *DescribeCenInterRegionBandwidthLimitsRequest {
	s.CenId = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsRequest) SetOwnerAccount(v string) *DescribeCenInterRegionBandwidthLimitsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsRequest) SetOwnerId(v int64) *DescribeCenInterRegionBandwidthLimitsRequest {
	s.OwnerId = &v
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

func (s *DescribeCenInterRegionBandwidthLimitsRequest) SetResourceOwnerAccount(v string) *DescribeCenInterRegionBandwidthLimitsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsRequest) SetResourceOwnerId(v int64) *DescribeCenInterRegionBandwidthLimitsRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeCenInterRegionBandwidthLimitsResponseBody struct {
	CenInterRegionBandwidthLimits *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimits `json:"CenInterRegionBandwidthLimits,omitempty" xml:"CenInterRegionBandwidthLimits,omitempty" type:"Struct"`
	PageNumber                    *int32                                                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize                      *int32                                                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId                     *string                                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount                    *int32                                                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCenInterRegionBandwidthLimitsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenInterRegionBandwidthLimitsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBody) SetCenInterRegionBandwidthLimits(v *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimits) *DescribeCenInterRegionBandwidthLimitsResponseBody {
	s.CenInterRegionBandwidthLimits = v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBody) SetPageNumber(v int32) *DescribeCenInterRegionBandwidthLimitsResponseBody {
	s.PageNumber = &v
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

func (s *DescribeCenInterRegionBandwidthLimitsResponseBody) SetTotalCount(v int32) *DescribeCenInterRegionBandwidthLimitsResponseBody {
	s.TotalCount = &v
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
	BandwidthLimit     *int64  `json:"BandwidthLimit,omitempty" xml:"BandwidthLimit,omitempty"`
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	CenId              *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	GeographicSpanId   *string `json:"GeographicSpanId,omitempty" xml:"GeographicSpanId,omitempty"`
	LocalRegionId      *string `json:"LocalRegionId,omitempty" xml:"LocalRegionId,omitempty"`
	OppositeRegionId   *string `json:"OppositeRegionId,omitempty" xml:"OppositeRegionId,omitempty"`
	Status             *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit) GoString() string {
	return s.String()
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit) SetBandwidthLimit(v int64) *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit {
	s.BandwidthLimit = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit) SetBandwidthPackageId(v string) *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit {
	s.BandwidthPackageId = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit) SetCenId(v string) *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit {
	s.CenId = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit) SetGeographicSpanId(v string) *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit {
	s.GeographicSpanId = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit) SetLocalRegionId(v string) *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit {
	s.LocalRegionId = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit) SetOppositeRegionId(v string) *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit {
	s.OppositeRegionId = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit) SetStatus(v string) *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit {
	s.Status = &v
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
	AccessRegionId       *string `json:"AccessRegionId,omitempty" xml:"AccessRegionId,omitempty"`
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	HostRegionId         *string `json:"HostRegionId,omitempty" xml:"HostRegionId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeCenPrivateZoneRoutesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenPrivateZoneRoutesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCenPrivateZoneRoutesRequest) SetAccessRegionId(v string) *DescribeCenPrivateZoneRoutesRequest {
	s.AccessRegionId = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesRequest) SetCenId(v string) *DescribeCenPrivateZoneRoutesRequest {
	s.CenId = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesRequest) SetHostRegionId(v string) *DescribeCenPrivateZoneRoutesRequest {
	s.HostRegionId = &v
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

func (s *DescribeCenPrivateZoneRoutesRequest) SetResourceOwnerAccount(v string) *DescribeCenPrivateZoneRoutesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesRequest) SetResourceOwnerId(v int64) *DescribeCenPrivateZoneRoutesRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeCenPrivateZoneRoutesResponseBody struct {
	CenId                 *string                                                   `json:"CenId,omitempty" xml:"CenId,omitempty"`
	PageNumber            *int32                                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize              *int32                                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PrivateZoneDnsServers *string                                                   `json:"PrivateZoneDnsServers,omitempty" xml:"PrivateZoneDnsServers,omitempty"`
	PrivateZoneInfos      *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfos `json:"PrivateZoneInfos,omitempty" xml:"PrivateZoneInfos,omitempty" type:"Struct"`
	RequestId             *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount            *int32                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCenPrivateZoneRoutesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenPrivateZoneRoutesResponseBody) GoString() string {
	return s.String()
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

func (s *DescribeCenPrivateZoneRoutesResponseBody) SetPrivateZoneDnsServers(v string) *DescribeCenPrivateZoneRoutesResponseBody {
	s.PrivateZoneDnsServers = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesResponseBody) SetPrivateZoneInfos(v *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfos) *DescribeCenPrivateZoneRoutesResponseBody {
	s.PrivateZoneInfos = v
	return s
}

func (s *DescribeCenPrivateZoneRoutesResponseBody) SetRequestId(v string) *DescribeCenPrivateZoneRoutesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesResponseBody) SetTotalCount(v int32) *DescribeCenPrivateZoneRoutesResponseBody {
	s.TotalCount = &v
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
	AccessRegionId *string `json:"AccessRegionId,omitempty" xml:"AccessRegionId,omitempty"`
	HostRegionId   *string `json:"HostRegionId,omitempty" xml:"HostRegionId,omitempty"`
	HostVpcId      *string `json:"HostVpcId,omitempty" xml:"HostVpcId,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfosPrivateZoneInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfosPrivateZoneInfo) GoString() string {
	return s.String()
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

func (s *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfosPrivateZoneInfo) SetStatus(v string) *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfosPrivateZoneInfo {
	s.Status = &v
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
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	CenRegionId          *string `json:"CenRegionId,omitempty" xml:"CenRegionId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCenRegionDomainRouteEntriesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenRegionDomainRouteEntriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCenRegionDomainRouteEntriesRequest) SetCenId(v string) *DescribeCenRegionDomainRouteEntriesRequest {
	s.CenId = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesRequest) SetCenRegionId(v string) *DescribeCenRegionDomainRouteEntriesRequest {
	s.CenRegionId = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesRequest) SetOwnerAccount(v string) *DescribeCenRegionDomainRouteEntriesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesRequest) SetOwnerId(v int64) *DescribeCenRegionDomainRouteEntriesRequest {
	s.OwnerId = &v
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

func (s *DescribeCenRegionDomainRouteEntriesRequest) SetResourceOwnerAccount(v string) *DescribeCenRegionDomainRouteEntriesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesRequest) SetResourceOwnerId(v int64) *DescribeCenRegionDomainRouteEntriesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesRequest) SetStatus(v string) *DescribeCenRegionDomainRouteEntriesRequest {
	s.Status = &v
	return s
}

type DescribeCenRegionDomainRouteEntriesResponseBody struct {
	CenRouteEntries *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries `json:"CenRouteEntries,omitempty" xml:"CenRouteEntries,omitempty" type:"Struct"`
	PageNumber      *int32                                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32                                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId       *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount      *int32                                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCenRegionDomainRouteEntriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenRegionDomainRouteEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBody) SetCenRouteEntries(v *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries) *DescribeCenRegionDomainRouteEntriesResponseBody {
	s.CenRouteEntries = v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBody) SetPageNumber(v int32) *DescribeCenRegionDomainRouteEntriesResponseBody {
	s.PageNumber = &v
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

func (s *DescribeCenRegionDomainRouteEntriesResponseBody) SetTotalCount(v int32) *DescribeCenRegionDomainRouteEntriesResponseBody {
	s.TotalCount = &v
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
	AsPaths               *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryAsPaths               `json:"AsPaths,omitempty" xml:"AsPaths,omitempty" type:"Struct"`
	CenOutRouteMapRecords *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecords `json:"CenOutRouteMapRecords,omitempty" xml:"CenOutRouteMapRecords,omitempty" type:"Struct"`
	CenRouteMapRecords    *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecords    `json:"CenRouteMapRecords,omitempty" xml:"CenRouteMapRecords,omitempty" type:"Struct"`
	Communities           *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCommunities           `json:"Communities,omitempty" xml:"Communities,omitempty" type:"Struct"`
	DestinationCidrBlock  *string                                                                                           `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	NextHopInstanceId     *string                                                                                           `json:"NextHopInstanceId,omitempty" xml:"NextHopInstanceId,omitempty"`
	NextHopRegionId       *string                                                                                           `json:"NextHopRegionId,omitempty" xml:"NextHopRegionId,omitempty"`
	NextHopType           *string                                                                                           `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	Preference            *int32                                                                                            `json:"Preference,omitempty" xml:"Preference,omitempty"`
	Status                *string                                                                                           `json:"Status,omitempty" xml:"Status,omitempty"`
	ToOtherRegionStatus   *string                                                                                           `json:"ToOtherRegionStatus,omitempty" xml:"ToOtherRegionStatus,omitempty"`
	Type                  *string                                                                                           `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) GoString() string {
	return s.String()
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetAsPaths(v *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryAsPaths) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.AsPaths = v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetCenOutRouteMapRecords(v *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecords) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.CenOutRouteMapRecords = v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetCenRouteMapRecords(v *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecords) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.CenRouteMapRecords = v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetCommunities(v *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCommunities) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.Communities = v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetDestinationCidrBlock(v string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.DestinationCidrBlock = &v
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

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetNextHopType(v string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.NextHopType = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetPreference(v int32) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.Preference = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetStatus(v string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.Status = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetToOtherRegionStatus(v string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.ToOtherRegionStatus = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetType(v string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.Type = &v
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
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RouteMapId *string `json:"RouteMapId,omitempty" xml:"RouteMapId,omitempty"`
}

func (s DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecordsCenOutRouteMapRecord) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecordsCenOutRouteMapRecord) GoString() string {
	return s.String()
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecordsCenOutRouteMapRecord) SetRegionId(v string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecordsCenOutRouteMapRecord {
	s.RegionId = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecordsCenOutRouteMapRecord) SetRouteMapId(v string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecordsCenOutRouteMapRecord {
	s.RouteMapId = &v
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
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RouteMapId *string `json:"RouteMapId,omitempty" xml:"RouteMapId,omitempty"`
}

func (s DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord) GoString() string {
	return s.String()
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord) SetRegionId(v string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord {
	s.RegionId = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord) SetRouteMapId(v string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord {
	s.RouteMapId = &v
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
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	CenRegionId          *string `json:"CenRegionId,omitempty" xml:"CenRegionId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RouteMapId           *string `json:"RouteMapId,omitempty" xml:"RouteMapId,omitempty"`
	TransmitDirection    *string `json:"TransmitDirection,omitempty" xml:"TransmitDirection,omitempty"`
}

func (s DescribeCenRouteMapsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenRouteMapsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCenRouteMapsRequest) SetCenId(v string) *DescribeCenRouteMapsRequest {
	s.CenId = &v
	return s
}

func (s *DescribeCenRouteMapsRequest) SetCenRegionId(v string) *DescribeCenRouteMapsRequest {
	s.CenRegionId = &v
	return s
}

func (s *DescribeCenRouteMapsRequest) SetOwnerAccount(v string) *DescribeCenRouteMapsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCenRouteMapsRequest) SetOwnerId(v int64) *DescribeCenRouteMapsRequest {
	s.OwnerId = &v
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

func (s *DescribeCenRouteMapsRequest) SetResourceOwnerAccount(v string) *DescribeCenRouteMapsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCenRouteMapsRequest) SetResourceOwnerId(v int64) *DescribeCenRouteMapsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCenRouteMapsRequest) SetRouteMapId(v string) *DescribeCenRouteMapsRequest {
	s.RouteMapId = &v
	return s
}

func (s *DescribeCenRouteMapsRequest) SetTransmitDirection(v string) *DescribeCenRouteMapsRequest {
	s.TransmitDirection = &v
	return s
}

type DescribeCenRouteMapsResponseBody struct {
	PageNumber *int32                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RouteMaps  *DescribeCenRouteMapsResponseBodyRouteMaps `json:"RouteMaps,omitempty" xml:"RouteMaps,omitempty" type:"Struct"`
	TotalCount *int32                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCenRouteMapsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenRouteMapsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCenRouteMapsResponseBody) SetPageNumber(v int32) *DescribeCenRouteMapsResponseBody {
	s.PageNumber = &v
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

func (s *DescribeCenRouteMapsResponseBody) SetRouteMaps(v *DescribeCenRouteMapsResponseBodyRouteMaps) *DescribeCenRouteMapsResponseBody {
	s.RouteMaps = v
	return s
}

func (s *DescribeCenRouteMapsResponseBody) SetTotalCount(v int32) *DescribeCenRouteMapsResponseBody {
	s.TotalCount = &v
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
	AsPathMatchMode                    *string                                                                         `json:"AsPathMatchMode,omitempty" xml:"AsPathMatchMode,omitempty"`
	CenId                              *string                                                                         `json:"CenId,omitempty" xml:"CenId,omitempty"`
	CenRegionId                        *string                                                                         `json:"CenRegionId,omitempty" xml:"CenRegionId,omitempty"`
	CidrMatchMode                      *string                                                                         `json:"CidrMatchMode,omitempty" xml:"CidrMatchMode,omitempty"`
	CommunityMatchMode                 *string                                                                         `json:"CommunityMatchMode,omitempty" xml:"CommunityMatchMode,omitempty"`
	CommunityOperateMode               *string                                                                         `json:"CommunityOperateMode,omitempty" xml:"CommunityOperateMode,omitempty"`
	Description                        *string                                                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	DestinationChildInstanceTypes      *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationChildInstanceTypes `json:"DestinationChildInstanceTypes,omitempty" xml:"DestinationChildInstanceTypes,omitempty" type:"Struct"`
	DestinationCidrBlocks              *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationCidrBlocks         `json:"DestinationCidrBlocks,omitempty" xml:"DestinationCidrBlocks,omitempty" type:"Struct"`
	DestinationInstanceIds             *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationInstanceIds        `json:"DestinationInstanceIds,omitempty" xml:"DestinationInstanceIds,omitempty" type:"Struct"`
	DestinationInstanceIdsReverseMatch *bool                                                                           `json:"DestinationInstanceIdsReverseMatch,omitempty" xml:"DestinationInstanceIdsReverseMatch,omitempty"`
	DestinationRouteTableIds           *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationRouteTableIds      `json:"DestinationRouteTableIds,omitempty" xml:"DestinationRouteTableIds,omitempty" type:"Struct"`
	MapResult                          *string                                                                         `json:"MapResult,omitempty" xml:"MapResult,omitempty"`
	MatchAsns                          *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapMatchAsns                     `json:"MatchAsns,omitempty" xml:"MatchAsns,omitempty" type:"Struct"`
	MatchCommunitySet                  *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapMatchCommunitySet             `json:"MatchCommunitySet,omitempty" xml:"MatchCommunitySet,omitempty" type:"Struct"`
	NextPriority                       *int32                                                                          `json:"NextPriority,omitempty" xml:"NextPriority,omitempty"`
	OperateCommunitySet                *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapOperateCommunitySet           `json:"OperateCommunitySet,omitempty" xml:"OperateCommunitySet,omitempty" type:"Struct"`
	Preference                         *int32                                                                          `json:"Preference,omitempty" xml:"Preference,omitempty"`
	PrependAsPath                      *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapPrependAsPath                 `json:"PrependAsPath,omitempty" xml:"PrependAsPath,omitempty" type:"Struct"`
	Priority                           *int32                                                                          `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RouteMapId                         *string                                                                         `json:"RouteMapId,omitempty" xml:"RouteMapId,omitempty"`
	RouteTypes                         *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapRouteTypes                    `json:"RouteTypes,omitempty" xml:"RouteTypes,omitempty" type:"Struct"`
	SourceChildInstanceTypes           *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceChildInstanceTypes      `json:"SourceChildInstanceTypes,omitempty" xml:"SourceChildInstanceTypes,omitempty" type:"Struct"`
	SourceInstanceIds                  *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceInstanceIds             `json:"SourceInstanceIds,omitempty" xml:"SourceInstanceIds,omitempty" type:"Struct"`
	SourceInstanceIdsReverseMatch      *bool                                                                           `json:"SourceInstanceIdsReverseMatch,omitempty" xml:"SourceInstanceIdsReverseMatch,omitempty"`
	SourceRegionIds                    *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceRegionIds               `json:"SourceRegionIds,omitempty" xml:"SourceRegionIds,omitempty" type:"Struct"`
	SourceRouteTableIds                *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceRouteTableIds           `json:"SourceRouteTableIds,omitempty" xml:"SourceRouteTableIds,omitempty" type:"Struct"`
	Status                             *string                                                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	TransmitDirection                  *string                                                                         `json:"TransmitDirection,omitempty" xml:"TransmitDirection,omitempty"`
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) GoString() string {
	return s.String()
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetAsPathMatchMode(v string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.AsPathMatchMode = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetCenId(v string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.CenId = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetCenRegionId(v string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.CenRegionId = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetCidrMatchMode(v string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.CidrMatchMode = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetCommunityMatchMode(v string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.CommunityMatchMode = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetCommunityOperateMode(v string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.CommunityOperateMode = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetDescription(v string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.Description = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetDestinationChildInstanceTypes(v *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationChildInstanceTypes) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.DestinationChildInstanceTypes = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetDestinationCidrBlocks(v *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationCidrBlocks) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.DestinationCidrBlocks = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetDestinationInstanceIds(v *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationInstanceIds) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.DestinationInstanceIds = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetDestinationInstanceIdsReverseMatch(v bool) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.DestinationInstanceIdsReverseMatch = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetDestinationRouteTableIds(v *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationRouteTableIds) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.DestinationRouteTableIds = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetMapResult(v string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.MapResult = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetMatchAsns(v *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapMatchAsns) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.MatchAsns = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetMatchCommunitySet(v *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapMatchCommunitySet) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.MatchCommunitySet = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetNextPriority(v int32) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.NextPriority = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetOperateCommunitySet(v *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapOperateCommunitySet) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.OperateCommunitySet = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetPreference(v int32) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.Preference = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetPrependAsPath(v *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapPrependAsPath) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.PrependAsPath = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetPriority(v int32) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.Priority = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetRouteMapId(v string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.RouteMapId = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetRouteTypes(v *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapRouteTypes) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.RouteTypes = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetSourceChildInstanceTypes(v *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceChildInstanceTypes) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.SourceChildInstanceTypes = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetSourceInstanceIds(v *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceInstanceIds) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.SourceInstanceIds = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetSourceInstanceIdsReverseMatch(v bool) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.SourceInstanceIdsReverseMatch = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetSourceRegionIds(v *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceRegionIds) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.SourceRegionIds = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetSourceRouteTableIds(v *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceRouteTableIds) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.SourceRouteTableIds = v
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

type DescribeCenVbrHealthCheckRequest struct {
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	VbrInstanceId        *string `json:"VbrInstanceId,omitempty" xml:"VbrInstanceId,omitempty"`
	VbrInstanceOwnerId   *int64  `json:"VbrInstanceOwnerId,omitempty" xml:"VbrInstanceOwnerId,omitempty"`
	VbrInstanceRegionId  *string `json:"VbrInstanceRegionId,omitempty" xml:"VbrInstanceRegionId,omitempty"`
}

func (s DescribeCenVbrHealthCheckRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenVbrHealthCheckRequest) GoString() string {
	return s.String()
}

func (s *DescribeCenVbrHealthCheckRequest) SetCenId(v string) *DescribeCenVbrHealthCheckRequest {
	s.CenId = &v
	return s
}

func (s *DescribeCenVbrHealthCheckRequest) SetOwnerAccount(v string) *DescribeCenVbrHealthCheckRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCenVbrHealthCheckRequest) SetOwnerId(v int64) *DescribeCenVbrHealthCheckRequest {
	s.OwnerId = &v
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

func (s *DescribeCenVbrHealthCheckRequest) SetResourceOwnerAccount(v string) *DescribeCenVbrHealthCheckRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCenVbrHealthCheckRequest) SetResourceOwnerId(v int64) *DescribeCenVbrHealthCheckRequest {
	s.ResourceOwnerId = &v
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

func (s *DescribeCenVbrHealthCheckRequest) SetVbrInstanceRegionId(v string) *DescribeCenVbrHealthCheckRequest {
	s.VbrInstanceRegionId = &v
	return s
}

type DescribeCenVbrHealthCheckResponseBody struct {
	PageNumber      *int32                                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId       *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount      *int32                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VbrHealthChecks *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecks `json:"VbrHealthChecks,omitempty" xml:"VbrHealthChecks,omitempty" type:"Struct"`
}

func (s DescribeCenVbrHealthCheckResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenVbrHealthCheckResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCenVbrHealthCheckResponseBody) SetPageNumber(v int32) *DescribeCenVbrHealthCheckResponseBody {
	s.PageNumber = &v
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
	CenId               *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	HealthCheckInterval *int32  `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	HealthCheckOnly     *bool   `json:"HealthCheckOnly,omitempty" xml:"HealthCheckOnly,omitempty"`
	HealthCheckSourceIp *string `json:"HealthCheckSourceIp,omitempty" xml:"HealthCheckSourceIp,omitempty"`
	HealthCheckTargetIp *string `json:"HealthCheckTargetIp,omitempty" xml:"HealthCheckTargetIp,omitempty"`
	HealthyThreshold    *int32  `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	VbrInstanceId       *string `json:"VbrInstanceId,omitempty" xml:"VbrInstanceId,omitempty"`
	VbrInstanceRegionId *string `json:"VbrInstanceRegionId,omitempty" xml:"VbrInstanceRegionId,omitempty"`
}

func (s DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck) String() string {
	return tea.Prettify(s)
}

func (s DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck) GoString() string {
	return s.String()
}

func (s *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck) SetCenId(v string) *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck {
	s.CenId = &v
	return s
}

func (s *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck) SetHealthCheckInterval(v int32) *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck {
	s.HealthCheckInterval = &v
	return s
}

func (s *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck) SetHealthCheckOnly(v bool) *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck {
	s.HealthCheckOnly = &v
	return s
}

func (s *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck) SetHealthCheckSourceIp(v string) *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck {
	s.HealthCheckSourceIp = &v
	return s
}

func (s *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck) SetHealthCheckTargetIp(v string) *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck {
	s.HealthCheckTargetIp = &v
	return s
}

func (s *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck) SetHealthyThreshold(v int32) *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck {
	s.HealthyThreshold = &v
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

type DescribeCensRequest struct {
	Filter               []*DescribeCensRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	OwnerAccount         *string                      `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64                       `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string                      `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                       `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Tag                  []*DescribeCensRequestTag    `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeCensRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCensRequest) GoString() string {
	return s.String()
}

func (s *DescribeCensRequest) SetFilter(v []*DescribeCensRequestFilter) *DescribeCensRequest {
	s.Filter = v
	return s
}

func (s *DescribeCensRequest) SetOwnerAccount(v string) *DescribeCensRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCensRequest) SetOwnerId(v int64) *DescribeCensRequest {
	s.OwnerId = &v
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

func (s *DescribeCensRequest) SetResourceOwnerAccount(v string) *DescribeCensRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCensRequest) SetResourceOwnerId(v int64) *DescribeCensRequest {
	s.ResourceOwnerId = &v
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
	Cens       *DescribeCensResponseBodyCens `json:"Cens,omitempty" xml:"Cens,omitempty" type:"Struct"`
	PageNumber *int32                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCensResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCensResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCensResponseBody) SetCens(v *DescribeCensResponseBodyCens) *DescribeCensResponseBody {
	s.Cens = v
	return s
}

func (s *DescribeCensResponseBody) SetPageNumber(v int32) *DescribeCensResponseBody {
	s.PageNumber = &v
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

func (s *DescribeCensResponseBody) SetTotalCount(v int32) *DescribeCensResponseBody {
	s.TotalCount = &v
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
	CenBandwidthPackageIds *DescribeCensResponseBodyCensCenCenBandwidthPackageIds `json:"CenBandwidthPackageIds,omitempty" xml:"CenBandwidthPackageIds,omitempty" type:"Struct"`
	CenId                  *string                                                `json:"CenId,omitempty" xml:"CenId,omitempty"`
	CreationTime           *string                                                `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description            *string                                                `json:"Description,omitempty" xml:"Description,omitempty"`
	Name                   *string                                                `json:"Name,omitempty" xml:"Name,omitempty"`
	ProtectionLevel        *string                                                `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty"`
	Status                 *string                                                `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags                   *DescribeCensResponseBodyCensCenTags                   `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s DescribeCensResponseBodyCensCen) String() string {
	return tea.Prettify(s)
}

func (s DescribeCensResponseBodyCensCen) GoString() string {
	return s.String()
}

func (s *DescribeCensResponseBodyCensCen) SetCenBandwidthPackageIds(v *DescribeCensResponseBodyCensCenCenBandwidthPackageIds) *DescribeCensResponseBodyCensCen {
	s.CenBandwidthPackageIds = v
	return s
}

func (s *DescribeCensResponseBodyCensCen) SetCenId(v string) *DescribeCensResponseBodyCensCen {
	s.CenId = &v
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

func (s *DescribeCensResponseBodyCensCen) SetName(v string) *DescribeCensResponseBodyCensCen {
	s.Name = &v
	return s
}

func (s *DescribeCensResponseBodyCensCen) SetProtectionLevel(v string) *DescribeCensResponseBodyCensCen {
	s.ProtectionLevel = &v
	return s
}

func (s *DescribeCensResponseBodyCensCen) SetStatus(v string) *DescribeCensResponseBodyCensCen {
	s.Status = &v
	return s
}

func (s *DescribeCensResponseBodyCensCen) SetTags(v *DescribeCensResponseBodyCensCenTags) *DescribeCensResponseBodyCensCen {
	s.Tags = v
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

type DescribeChildInstanceRegionsRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProductType          *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
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

func (s *DescribeChildInstanceRegionsRequest) SetProductType(v string) *DescribeChildInstanceRegionsRequest {
	s.ProductType = &v
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

type DescribeChildInstanceRegionsResponseBody struct {
	Regions   *DescribeChildInstanceRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeChildInstanceRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeChildInstanceRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChildInstanceRegionsResponseBody) SetRegions(v *DescribeChildInstanceRegionsResponseBodyRegions) *DescribeChildInstanceRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeChildInstanceRegionsResponseBody) SetRequestId(v string) *DescribeChildInstanceRegionsResponseBody {
	s.RequestId = &v
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
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FlowLogId            *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
	FlowLogName          *string `json:"FlowLogName,omitempty" xml:"FlowLogName,omitempty"`
	LogStoreName         *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectName          *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeFlowlogsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowlogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowlogsRequest) SetCenId(v string) *DescribeFlowlogsRequest {
	s.CenId = &v
	return s
}

func (s *DescribeFlowlogsRequest) SetClientToken(v string) *DescribeFlowlogsRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeFlowlogsRequest) SetDescription(v string) *DescribeFlowlogsRequest {
	s.Description = &v
	return s
}

func (s *DescribeFlowlogsRequest) SetFlowLogId(v string) *DescribeFlowlogsRequest {
	s.FlowLogId = &v
	return s
}

func (s *DescribeFlowlogsRequest) SetFlowLogName(v string) *DescribeFlowlogsRequest {
	s.FlowLogName = &v
	return s
}

func (s *DescribeFlowlogsRequest) SetLogStoreName(v string) *DescribeFlowlogsRequest {
	s.LogStoreName = &v
	return s
}

func (s *DescribeFlowlogsRequest) SetOwnerAccount(v string) *DescribeFlowlogsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeFlowlogsRequest) SetOwnerId(v int64) *DescribeFlowlogsRequest {
	s.OwnerId = &v
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

func (s *DescribeFlowlogsRequest) SetProjectName(v string) *DescribeFlowlogsRequest {
	s.ProjectName = &v
	return s
}

func (s *DescribeFlowlogsRequest) SetRegionId(v string) *DescribeFlowlogsRequest {
	s.RegionId = &v
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

func (s *DescribeFlowlogsRequest) SetStatus(v string) *DescribeFlowlogsRequest {
	s.Status = &v
	return s
}

type DescribeFlowlogsResponseBody struct {
	FlowLogs   *DescribeFlowlogsResponseBodyFlowLogs `json:"FlowLogs,omitempty" xml:"FlowLogs,omitempty" type:"Struct"`
	PageNumber *string                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *string                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string                               `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *string                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeFlowlogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowlogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowlogsResponseBody) SetFlowLogs(v *DescribeFlowlogsResponseBodyFlowLogs) *DescribeFlowlogsResponseBody {
	s.FlowLogs = v
	return s
}

func (s *DescribeFlowlogsResponseBody) SetPageNumber(v string) *DescribeFlowlogsResponseBody {
	s.PageNumber = &v
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

func (s *DescribeFlowlogsResponseBody) SetSuccess(v string) *DescribeFlowlogsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeFlowlogsResponseBody) SetTotalCount(v string) *DescribeFlowlogsResponseBody {
	s.TotalCount = &v
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
	CenId        *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FlowLogId    *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
	FlowLogName  *string `json:"FlowLogName,omitempty" xml:"FlowLogName,omitempty"`
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	ProjectName  *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeFlowlogsResponseBodyFlowLogsFlowLog) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowlogsResponseBodyFlowLogsFlowLog) GoString() string {
	return s.String()
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) SetCenId(v string) *DescribeFlowlogsResponseBodyFlowLogsFlowLog {
	s.CenId = &v
	return s
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) SetCreationTime(v string) *DescribeFlowlogsResponseBodyFlowLogsFlowLog {
	s.CreationTime = &v
	return s
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) SetDescription(v string) *DescribeFlowlogsResponseBodyFlowLogsFlowLog {
	s.Description = &v
	return s
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) SetFlowLogId(v string) *DescribeFlowlogsResponseBodyFlowLogsFlowLog {
	s.FlowLogId = &v
	return s
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) SetFlowLogName(v string) *DescribeFlowlogsResponseBodyFlowLogsFlowLog {
	s.FlowLogName = &v
	return s
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) SetLogStoreName(v string) *DescribeFlowlogsResponseBodyFlowLogsFlowLog {
	s.LogStoreName = &v
	return s
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) SetProjectName(v string) *DescribeFlowlogsResponseBodyFlowLogsFlowLog {
	s.ProjectName = &v
	return s
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) SetRegionId(v string) *DescribeFlowlogsResponseBodyFlowLogsFlowLog {
	s.RegionId = &v
	return s
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) SetStatus(v string) *DescribeFlowlogsResponseBodyFlowLogsFlowLog {
	s.Status = &v
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
	GeographicRegionId   *string `json:"GeographicRegionId,omitempty" xml:"GeographicRegionId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeGeographicRegionMembershipRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGeographicRegionMembershipRequest) GoString() string {
	return s.String()
}

func (s *DescribeGeographicRegionMembershipRequest) SetGeographicRegionId(v string) *DescribeGeographicRegionMembershipRequest {
	s.GeographicRegionId = &v
	return s
}

func (s *DescribeGeographicRegionMembershipRequest) SetOwnerAccount(v string) *DescribeGeographicRegionMembershipRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeGeographicRegionMembershipRequest) SetOwnerId(v int64) *DescribeGeographicRegionMembershipRequest {
	s.OwnerId = &v
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

func (s *DescribeGeographicRegionMembershipRequest) SetResourceOwnerAccount(v string) *DescribeGeographicRegionMembershipRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeGeographicRegionMembershipRequest) SetResourceOwnerId(v int64) *DescribeGeographicRegionMembershipRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeGeographicRegionMembershipResponseBody struct {
	PageNumber *int32                                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionIds  *DescribeGeographicRegionMembershipResponseBodyRegionIds `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Struct"`
	RequestId  *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeGeographicRegionMembershipResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGeographicRegionMembershipResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGeographicRegionMembershipResponseBody) SetPageNumber(v int32) *DescribeGeographicRegionMembershipResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeGeographicRegionMembershipResponseBody) SetPageSize(v int32) *DescribeGeographicRegionMembershipResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGeographicRegionMembershipResponseBody) SetRegionIds(v *DescribeGeographicRegionMembershipResponseBodyRegionIds) *DescribeGeographicRegionMembershipResponseBody {
	s.RegionIds = v
	return s
}

func (s *DescribeGeographicRegionMembershipResponseBody) SetRequestId(v string) *DescribeGeographicRegionMembershipResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGeographicRegionMembershipResponseBody) SetTotalCount(v int32) *DescribeGeographicRegionMembershipResponseBody {
	s.TotalCount = &v
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
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProductType          *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeGrantRulesToCenRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGrantRulesToCenRequest) GoString() string {
	return s.String()
}

func (s *DescribeGrantRulesToCenRequest) SetCenId(v string) *DescribeGrantRulesToCenRequest {
	s.CenId = &v
	return s
}

func (s *DescribeGrantRulesToCenRequest) SetOwnerAccount(v string) *DescribeGrantRulesToCenRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeGrantRulesToCenRequest) SetOwnerId(v int64) *DescribeGrantRulesToCenRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeGrantRulesToCenRequest) SetProductType(v string) *DescribeGrantRulesToCenRequest {
	s.ProductType = &v
	return s
}

func (s *DescribeGrantRulesToCenRequest) SetRegionId(v string) *DescribeGrantRulesToCenRequest {
	s.RegionId = &v
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

type DescribeGrantRulesToCenResponseBody struct {
	GrantRules *DescribeGrantRulesToCenResponseBodyGrantRules `json:"GrantRules,omitempty" xml:"GrantRules,omitempty" type:"Struct"`
	RequestId  *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeGrantRulesToCenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGrantRulesToCenResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGrantRulesToCenResponseBody) SetGrantRules(v *DescribeGrantRulesToCenResponseBodyGrantRules) *DescribeGrantRulesToCenResponseBody {
	s.GrantRules = v
	return s
}

func (s *DescribeGrantRulesToCenResponseBody) SetRequestId(v string) *DescribeGrantRulesToCenResponseBody {
	s.RequestId = &v
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
	CenId                 *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ChildInstanceId       *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	ChildInstanceOwnerId  *int64  `json:"ChildInstanceOwnerId,omitempty" xml:"ChildInstanceOwnerId,omitempty"`
	ChildInstanceRegionId *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	ChildInstanceType     *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	OrderType             *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
}

func (s DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule) String() string {
	return tea.Prettify(s)
}

func (s DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule) GoString() string {
	return s.String()
}

func (s *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule) SetCenId(v string) *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule {
	s.CenId = &v
	return s
}

func (s *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule) SetChildInstanceId(v string) *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule {
	s.ChildInstanceId = &v
	return s
}

func (s *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule) SetChildInstanceOwnerId(v int64) *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule {
	s.ChildInstanceOwnerId = &v
	return s
}

func (s *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule) SetChildInstanceRegionId(v string) *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule) SetChildInstanceType(v string) *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule {
	s.ChildInstanceType = &v
	return s
}

func (s *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule) SetOrderType(v string) *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule {
	s.OrderType = &v
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
	CenId                     *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ChildInstanceId           *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	ChildInstanceRegionId     *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	ChildInstanceRouteTableId *string `json:"ChildInstanceRouteTableId,omitempty" xml:"ChildInstanceRouteTableId,omitempty"`
	ChildInstanceType         *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	DestinationCidrBlock      *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	PageNumber                *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize                  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount      *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId           *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribePublishedRouteEntriesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePublishedRouteEntriesRequest) GoString() string {
	return s.String()
}

func (s *DescribePublishedRouteEntriesRequest) SetCenId(v string) *DescribePublishedRouteEntriesRequest {
	s.CenId = &v
	return s
}

func (s *DescribePublishedRouteEntriesRequest) SetChildInstanceId(v string) *DescribePublishedRouteEntriesRequest {
	s.ChildInstanceId = &v
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

func (s *DescribePublishedRouteEntriesRequest) SetChildInstanceType(v string) *DescribePublishedRouteEntriesRequest {
	s.ChildInstanceType = &v
	return s
}

func (s *DescribePublishedRouteEntriesRequest) SetDestinationCidrBlock(v string) *DescribePublishedRouteEntriesRequest {
	s.DestinationCidrBlock = &v
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

func (s *DescribePublishedRouteEntriesRequest) SetResourceOwnerAccount(v string) *DescribePublishedRouteEntriesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePublishedRouteEntriesRequest) SetResourceOwnerId(v int64) *DescribePublishedRouteEntriesRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribePublishedRouteEntriesResponseBody struct {
	PageNumber            *int32                                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize              *int32                                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PublishedRouteEntries *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntries `json:"PublishedRouteEntries,omitempty" xml:"PublishedRouteEntries,omitempty" type:"Struct"`
	RequestId             *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount            *int32                                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePublishedRouteEntriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePublishedRouteEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePublishedRouteEntriesResponseBody) SetPageNumber(v int32) *DescribePublishedRouteEntriesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBody) SetPageSize(v int32) *DescribePublishedRouteEntriesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBody) SetPublishedRouteEntries(v *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntries) *DescribePublishedRouteEntriesResponseBody {
	s.PublishedRouteEntries = v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBody) SetRequestId(v string) *DescribePublishedRouteEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBody) SetTotalCount(v int32) *DescribePublishedRouteEntriesResponseBody {
	s.TotalCount = &v
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
	ChildInstanceRouteTableId *string                                                                                     `json:"ChildInstanceRouteTableId,omitempty" xml:"ChildInstanceRouteTableId,omitempty"`
	Conflicts                 *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflicts `json:"Conflicts,omitempty" xml:"Conflicts,omitempty" type:"Struct"`
	DestinationCidrBlock      *string                                                                                     `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	NextHopId                 *string                                                                                     `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	NextHopType               *string                                                                                     `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	OperationalMode           *bool                                                                                       `json:"OperationalMode,omitempty" xml:"OperationalMode,omitempty"`
	PublishStatus             *string                                                                                     `json:"PublishStatus,omitempty" xml:"PublishStatus,omitempty"`
	RouteType                 *string                                                                                     `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
}

func (s DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry) String() string {
	return tea.Prettify(s)
}

func (s DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry) GoString() string {
	return s.String()
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry) SetChildInstanceRouteTableId(v string) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry {
	s.ChildInstanceRouteTableId = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry) SetConflicts(v *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflicts) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry {
	s.Conflicts = v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry) SetDestinationCidrBlock(v string) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry) SetNextHopId(v string) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry {
	s.NextHopId = &v
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

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry) SetPublishStatus(v string) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry {
	s.PublishStatus = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry) SetRouteType(v string) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry {
	s.RouteType = &v
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
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType         *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflictsConflict) String() string {
	return tea.Prettify(s)
}

func (s DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflictsConflict) GoString() string {
	return s.String()
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

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflictsConflict) SetStatus(v string) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflictsConflict {
	s.Status = &v
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
	ChildInstanceId           *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	ChildInstanceRegionId     *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	ChildInstanceRouteTableId *string `json:"ChildInstanceRouteTableId,omitempty" xml:"ChildInstanceRouteTableId,omitempty"`
	ChildInstanceType         *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	DestinationCidrBlock      *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	OwnerAccount              *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber                *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize                  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount      *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId           *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeRouteConflictRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRouteConflictRequest) GoString() string {
	return s.String()
}

func (s *DescribeRouteConflictRequest) SetChildInstanceId(v string) *DescribeRouteConflictRequest {
	s.ChildInstanceId = &v
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

func (s *DescribeRouteConflictRequest) SetChildInstanceType(v string) *DescribeRouteConflictRequest {
	s.ChildInstanceType = &v
	return s
}

func (s *DescribeRouteConflictRequest) SetDestinationCidrBlock(v string) *DescribeRouteConflictRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DescribeRouteConflictRequest) SetOwnerAccount(v string) *DescribeRouteConflictRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRouteConflictRequest) SetOwnerId(v int64) *DescribeRouteConflictRequest {
	s.OwnerId = &v
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

func (s *DescribeRouteConflictRequest) SetResourceOwnerAccount(v string) *DescribeRouteConflictRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRouteConflictRequest) SetResourceOwnerId(v int64) *DescribeRouteConflictRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeRouteConflictResponseBody struct {
	PageNumber     *int32                                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RouteConflicts *DescribeRouteConflictResponseBodyRouteConflicts `json:"RouteConflicts,omitempty" xml:"RouteConflicts,omitempty" type:"Struct"`
	TotalCount     *int32                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRouteConflictResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRouteConflictResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRouteConflictResponseBody) SetPageNumber(v int32) *DescribeRouteConflictResponseBody {
	s.PageNumber = &v
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

func (s *DescribeRouteConflictResponseBody) SetRouteConflicts(v *DescribeRouteConflictResponseBodyRouteConflicts) *DescribeRouteConflictResponseBody {
	s.RouteConflicts = v
	return s
}

func (s *DescribeRouteConflictResponseBody) SetTotalCount(v int32) *DescribeRouteConflictResponseBody {
	s.TotalCount = &v
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
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType         *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeRouteConflictResponseBodyRouteConflictsRouteConflict) String() string {
	return tea.Prettify(s)
}

func (s DescribeRouteConflictResponseBodyRouteConflictsRouteConflict) GoString() string {
	return s.String()
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

func (s *DescribeRouteConflictResponseBodyRouteConflictsRouteConflict) SetStatus(v string) *DescribeRouteConflictResponseBodyRouteConflictsRouteConflict {
	s.Status = &v
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
	AccessRegionId       *string `json:"AccessRegionId,omitempty" xml:"AccessRegionId,omitempty"`
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	Host                 *string `json:"Host,omitempty" xml:"Host,omitempty"`
	HostRegionId         *string `json:"HostRegionId,omitempty" xml:"HostRegionId,omitempty"`
	HostVpcId            *string `json:"HostVpcId,omitempty" xml:"HostVpcId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeRouteServicesInCenRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRouteServicesInCenRequest) GoString() string {
	return s.String()
}

func (s *DescribeRouteServicesInCenRequest) SetAccessRegionId(v string) *DescribeRouteServicesInCenRequest {
	s.AccessRegionId = &v
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

func (s *DescribeRouteServicesInCenRequest) SetHostVpcId(v string) *DescribeRouteServicesInCenRequest {
	s.HostVpcId = &v
	return s
}

func (s *DescribeRouteServicesInCenRequest) SetOwnerAccount(v string) *DescribeRouteServicesInCenRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRouteServicesInCenRequest) SetOwnerId(v int64) *DescribeRouteServicesInCenRequest {
	s.OwnerId = &v
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

func (s *DescribeRouteServicesInCenRequest) SetResourceOwnerAccount(v string) *DescribeRouteServicesInCenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRouteServicesInCenRequest) SetResourceOwnerId(v int64) *DescribeRouteServicesInCenRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeRouteServicesInCenResponseBody struct {
	PageNumber          *int32                                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize            *int32                                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId           *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RouteServiceEntries *DescribeRouteServicesInCenResponseBodyRouteServiceEntries `json:"RouteServiceEntries,omitempty" xml:"RouteServiceEntries,omitempty" type:"Struct"`
	TotalCount          *int32                                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRouteServicesInCenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRouteServicesInCenResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRouteServicesInCenResponseBody) SetPageNumber(v int32) *DescribeRouteServicesInCenResponseBody {
	s.PageNumber = &v
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

func (s *DescribeRouteServicesInCenResponseBody) SetRouteServiceEntries(v *DescribeRouteServicesInCenResponseBodyRouteServiceEntries) *DescribeRouteServicesInCenResponseBody {
	s.RouteServiceEntries = v
	return s
}

func (s *DescribeRouteServicesInCenResponseBody) SetTotalCount(v int32) *DescribeRouteServicesInCenResponseBody {
	s.TotalCount = &v
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
	AccessRegionId *string                                                                          `json:"AccessRegionId,omitempty" xml:"AccessRegionId,omitempty"`
	CenId          *string                                                                          `json:"CenId,omitempty" xml:"CenId,omitempty"`
	Cidrs          *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntryCidrs `json:"Cidrs,omitempty" xml:"Cidrs,omitempty" type:"Struct"`
	Description    *string                                                                          `json:"Description,omitempty" xml:"Description,omitempty"`
	Host           *string                                                                          `json:"Host,omitempty" xml:"Host,omitempty"`
	HostRegionId   *string                                                                          `json:"HostRegionId,omitempty" xml:"HostRegionId,omitempty"`
	HostVpcId      *string                                                                          `json:"HostVpcId,omitempty" xml:"HostVpcId,omitempty"`
	Status         *string                                                                          `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry) String() string {
	return tea.Prettify(s)
}

func (s DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry) GoString() string {
	return s.String()
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry) SetAccessRegionId(v string) *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry {
	s.AccessRegionId = &v
	return s
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry) SetCenId(v string) *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry {
	s.CenId = &v
	return s
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry) SetCidrs(v *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntryCidrs) *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry {
	s.Cidrs = v
	return s
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry) SetDescription(v string) *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry {
	s.Description = &v
	return s
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry) SetHost(v string) *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry {
	s.Host = &v
	return s
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry) SetHostRegionId(v string) *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry {
	s.HostRegionId = &v
	return s
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry) SetHostVpcId(v string) *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry {
	s.HostVpcId = &v
	return s
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry) SetStatus(v string) *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry {
	s.Status = &v
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
	CenId                 *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	CenOwnerId            *int64  `json:"CenOwnerId,omitempty" xml:"CenOwnerId,omitempty"`
	ChildInstanceId       *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	ChildInstanceOwnerId  *int64  `json:"ChildInstanceOwnerId,omitempty" xml:"ChildInstanceOwnerId,omitempty"`
	ChildInstanceRegionId *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	ChildInstanceType     *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DetachCenChildInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachCenChildInstanceRequest) GoString() string {
	return s.String()
}

func (s *DetachCenChildInstanceRequest) SetCenId(v string) *DetachCenChildInstanceRequest {
	s.CenId = &v
	return s
}

func (s *DetachCenChildInstanceRequest) SetCenOwnerId(v int64) *DetachCenChildInstanceRequest {
	s.CenOwnerId = &v
	return s
}

func (s *DetachCenChildInstanceRequest) SetChildInstanceId(v string) *DetachCenChildInstanceRequest {
	s.ChildInstanceId = &v
	return s
}

func (s *DetachCenChildInstanceRequest) SetChildInstanceOwnerId(v int64) *DetachCenChildInstanceRequest {
	s.ChildInstanceOwnerId = &v
	return s
}

func (s *DetachCenChildInstanceRequest) SetChildInstanceRegionId(v string) *DetachCenChildInstanceRequest {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *DetachCenChildInstanceRequest) SetChildInstanceType(v string) *DetachCenChildInstanceRequest {
	s.ChildInstanceType = &v
	return s
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
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	VbrInstanceId        *string `json:"VbrInstanceId,omitempty" xml:"VbrInstanceId,omitempty"`
	VbrInstanceOwnerId   *int64  `json:"VbrInstanceOwnerId,omitempty" xml:"VbrInstanceOwnerId,omitempty"`
	VbrInstanceRegionId  *string `json:"VbrInstanceRegionId,omitempty" xml:"VbrInstanceRegionId,omitempty"`
}

func (s DisableCenVbrHealthCheckRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableCenVbrHealthCheckRequest) GoString() string {
	return s.String()
}

func (s *DisableCenVbrHealthCheckRequest) SetCenId(v string) *DisableCenVbrHealthCheckRequest {
	s.CenId = &v
	return s
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

func (s *DisableCenVbrHealthCheckRequest) SetVbrInstanceId(v string) *DisableCenVbrHealthCheckRequest {
	s.VbrInstanceId = &v
	return s
}

func (s *DisableCenVbrHealthCheckRequest) SetVbrInstanceOwnerId(v int64) *DisableCenVbrHealthCheckRequest {
	s.VbrInstanceOwnerId = &v
	return s
}

func (s *DisableCenVbrHealthCheckRequest) SetVbrInstanceRegionId(v string) *DisableCenVbrHealthCheckRequest {
	s.VbrInstanceRegionId = &v
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

type DisableTransitRouterRouteTablePropagationRequest struct {
	ClientToken               *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                    *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount              *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount      *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId           *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	TransitRouterRouteTableId *string `json:"TransitRouterRouteTableId,omitempty" xml:"TransitRouterRouteTableId,omitempty"`
}

func (s DisableTransitRouterRouteTablePropagationRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableTransitRouterRouteTablePropagationRequest) GoString() string {
	return s.String()
}

func (s *DisableTransitRouterRouteTablePropagationRequest) SetClientToken(v string) *DisableTransitRouterRouteTablePropagationRequest {
	s.ClientToken = &v
	return s
}

func (s *DisableTransitRouterRouteTablePropagationRequest) SetDryRun(v bool) *DisableTransitRouterRouteTablePropagationRequest {
	s.DryRun = &v
	return s
}

func (s *DisableTransitRouterRouteTablePropagationRequest) SetOwnerAccount(v string) *DisableTransitRouterRouteTablePropagationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DisableTransitRouterRouteTablePropagationRequest) SetOwnerId(v int64) *DisableTransitRouterRouteTablePropagationRequest {
	s.OwnerId = &v
	return s
}

func (s *DisableTransitRouterRouteTablePropagationRequest) SetResourceOwnerAccount(v string) *DisableTransitRouterRouteTablePropagationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DisableTransitRouterRouteTablePropagationRequest) SetResourceOwnerId(v int64) *DisableTransitRouterRouteTablePropagationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DisableTransitRouterRouteTablePropagationRequest) SetTransitRouterAttachmentId(v string) *DisableTransitRouterRouteTablePropagationRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *DisableTransitRouterRouteTablePropagationRequest) SetTransitRouterRouteTableId(v string) *DisableTransitRouterRouteTablePropagationRequest {
	s.TransitRouterRouteTableId = &v
	return s
}

type DisableTransitRouterRouteTablePropagationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableTransitRouterRouteTablePropagationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableTransitRouterRouteTablePropagationResponseBody) GoString() string {
	return s.String()
}

func (s *DisableTransitRouterRouteTablePropagationResponseBody) SetRequestId(v string) *DisableTransitRouterRouteTablePropagationResponseBody {
	s.RequestId = &v
	return s
}

type DisableTransitRouterRouteTablePropagationResponse struct {
	Headers map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableTransitRouterRouteTablePropagationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableTransitRouterRouteTablePropagationResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableTransitRouterRouteTablePropagationResponse) GoString() string {
	return s.String()
}

func (s *DisableTransitRouterRouteTablePropagationResponse) SetHeaders(v map[string]*string) *DisableTransitRouterRouteTablePropagationResponse {
	s.Headers = v
	return s
}

func (s *DisableTransitRouterRouteTablePropagationResponse) SetBody(v *DisableTransitRouterRouteTablePropagationResponseBody) *DisableTransitRouterRouteTablePropagationResponse {
	s.Body = v
	return s
}

type DissociateTransitRouterAttachmentFromRouteTableRequest struct {
	ClientToken               *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                    *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount              *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount      *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId           *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	TransitRouterRouteTableId *string `json:"TransitRouterRouteTableId,omitempty" xml:"TransitRouterRouteTableId,omitempty"`
}

func (s DissociateTransitRouterAttachmentFromRouteTableRequest) String() string {
	return tea.Prettify(s)
}

func (s DissociateTransitRouterAttachmentFromRouteTableRequest) GoString() string {
	return s.String()
}

func (s *DissociateTransitRouterAttachmentFromRouteTableRequest) SetClientToken(v string) *DissociateTransitRouterAttachmentFromRouteTableRequest {
	s.ClientToken = &v
	return s
}

func (s *DissociateTransitRouterAttachmentFromRouteTableRequest) SetDryRun(v bool) *DissociateTransitRouterAttachmentFromRouteTableRequest {
	s.DryRun = &v
	return s
}

func (s *DissociateTransitRouterAttachmentFromRouteTableRequest) SetOwnerAccount(v string) *DissociateTransitRouterAttachmentFromRouteTableRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DissociateTransitRouterAttachmentFromRouteTableRequest) SetOwnerId(v int64) *DissociateTransitRouterAttachmentFromRouteTableRequest {
	s.OwnerId = &v
	return s
}

func (s *DissociateTransitRouterAttachmentFromRouteTableRequest) SetResourceOwnerAccount(v string) *DissociateTransitRouterAttachmentFromRouteTableRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DissociateTransitRouterAttachmentFromRouteTableRequest) SetResourceOwnerId(v int64) *DissociateTransitRouterAttachmentFromRouteTableRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DissociateTransitRouterAttachmentFromRouteTableRequest) SetTransitRouterAttachmentId(v string) *DissociateTransitRouterAttachmentFromRouteTableRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *DissociateTransitRouterAttachmentFromRouteTableRequest) SetTransitRouterRouteTableId(v string) *DissociateTransitRouterAttachmentFromRouteTableRequest {
	s.TransitRouterRouteTableId = &v
	return s
}

type DissociateTransitRouterAttachmentFromRouteTableResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DissociateTransitRouterAttachmentFromRouteTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DissociateTransitRouterAttachmentFromRouteTableResponseBody) GoString() string {
	return s.String()
}

func (s *DissociateTransitRouterAttachmentFromRouteTableResponseBody) SetRequestId(v string) *DissociateTransitRouterAttachmentFromRouteTableResponseBody {
	s.RequestId = &v
	return s
}

type DissociateTransitRouterAttachmentFromRouteTableResponse struct {
	Headers map[string]*string                                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DissociateTransitRouterAttachmentFromRouteTableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DissociateTransitRouterAttachmentFromRouteTableResponse) String() string {
	return tea.Prettify(s)
}

func (s DissociateTransitRouterAttachmentFromRouteTableResponse) GoString() string {
	return s.String()
}

func (s *DissociateTransitRouterAttachmentFromRouteTableResponse) SetHeaders(v map[string]*string) *DissociateTransitRouterAttachmentFromRouteTableResponse {
	s.Headers = v
	return s
}

func (s *DissociateTransitRouterAttachmentFromRouteTableResponse) SetBody(v *DissociateTransitRouterAttachmentFromRouteTableResponseBody) *DissociateTransitRouterAttachmentFromRouteTableResponse {
	s.Body = v
	return s
}

type EnableCenVbrHealthCheckRequest struct {
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	HealthCheckInterval  *int32  `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	HealthCheckOnly      *bool   `json:"HealthCheckOnly,omitempty" xml:"HealthCheckOnly,omitempty"`
	HealthCheckSourceIp  *string `json:"HealthCheckSourceIp,omitempty" xml:"HealthCheckSourceIp,omitempty"`
	HealthCheckTargetIp  *string `json:"HealthCheckTargetIp,omitempty" xml:"HealthCheckTargetIp,omitempty"`
	HealthyThreshold     *int32  `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	VbrInstanceId        *string `json:"VbrInstanceId,omitempty" xml:"VbrInstanceId,omitempty"`
	VbrInstanceOwnerId   *int64  `json:"VbrInstanceOwnerId,omitempty" xml:"VbrInstanceOwnerId,omitempty"`
	VbrInstanceRegionId  *string `json:"VbrInstanceRegionId,omitempty" xml:"VbrInstanceRegionId,omitempty"`
}

func (s EnableCenVbrHealthCheckRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableCenVbrHealthCheckRequest) GoString() string {
	return s.String()
}

func (s *EnableCenVbrHealthCheckRequest) SetCenId(v string) *EnableCenVbrHealthCheckRequest {
	s.CenId = &v
	return s
}

func (s *EnableCenVbrHealthCheckRequest) SetHealthCheckInterval(v int32) *EnableCenVbrHealthCheckRequest {
	s.HealthCheckInterval = &v
	return s
}

func (s *EnableCenVbrHealthCheckRequest) SetHealthCheckOnly(v bool) *EnableCenVbrHealthCheckRequest {
	s.HealthCheckOnly = &v
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

func (s *EnableCenVbrHealthCheckRequest) SetHealthyThreshold(v int32) *EnableCenVbrHealthCheckRequest {
	s.HealthyThreshold = &v
	return s
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

func (s *EnableCenVbrHealthCheckRequest) SetVbrInstanceId(v string) *EnableCenVbrHealthCheckRequest {
	s.VbrInstanceId = &v
	return s
}

func (s *EnableCenVbrHealthCheckRequest) SetVbrInstanceOwnerId(v int64) *EnableCenVbrHealthCheckRequest {
	s.VbrInstanceOwnerId = &v
	return s
}

func (s *EnableCenVbrHealthCheckRequest) SetVbrInstanceRegionId(v string) *EnableCenVbrHealthCheckRequest {
	s.VbrInstanceRegionId = &v
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

type EnableTransitRouterRouteTablePropagationRequest struct {
	ClientToken               *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                    *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount              *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount      *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId           *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	TransitRouterRouteTableId *string `json:"TransitRouterRouteTableId,omitempty" xml:"TransitRouterRouteTableId,omitempty"`
}

func (s EnableTransitRouterRouteTablePropagationRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableTransitRouterRouteTablePropagationRequest) GoString() string {
	return s.String()
}

func (s *EnableTransitRouterRouteTablePropagationRequest) SetClientToken(v string) *EnableTransitRouterRouteTablePropagationRequest {
	s.ClientToken = &v
	return s
}

func (s *EnableTransitRouterRouteTablePropagationRequest) SetDryRun(v bool) *EnableTransitRouterRouteTablePropagationRequest {
	s.DryRun = &v
	return s
}

func (s *EnableTransitRouterRouteTablePropagationRequest) SetOwnerAccount(v string) *EnableTransitRouterRouteTablePropagationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *EnableTransitRouterRouteTablePropagationRequest) SetOwnerId(v int64) *EnableTransitRouterRouteTablePropagationRequest {
	s.OwnerId = &v
	return s
}

func (s *EnableTransitRouterRouteTablePropagationRequest) SetResourceOwnerAccount(v string) *EnableTransitRouterRouteTablePropagationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *EnableTransitRouterRouteTablePropagationRequest) SetResourceOwnerId(v int64) *EnableTransitRouterRouteTablePropagationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *EnableTransitRouterRouteTablePropagationRequest) SetTransitRouterAttachmentId(v string) *EnableTransitRouterRouteTablePropagationRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *EnableTransitRouterRouteTablePropagationRequest) SetTransitRouterRouteTableId(v string) *EnableTransitRouterRouteTablePropagationRequest {
	s.TransitRouterRouteTableId = &v
	return s
}

type EnableTransitRouterRouteTablePropagationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableTransitRouterRouteTablePropagationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableTransitRouterRouteTablePropagationResponseBody) GoString() string {
	return s.String()
}

func (s *EnableTransitRouterRouteTablePropagationResponseBody) SetRequestId(v string) *EnableTransitRouterRouteTablePropagationResponseBody {
	s.RequestId = &v
	return s
}

type EnableTransitRouterRouteTablePropagationResponse struct {
	Headers map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableTransitRouterRouteTablePropagationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableTransitRouterRouteTablePropagationResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableTransitRouterRouteTablePropagationResponse) GoString() string {
	return s.String()
}

func (s *EnableTransitRouterRouteTablePropagationResponse) SetHeaders(v map[string]*string) *EnableTransitRouterRouteTablePropagationResponse {
	s.Headers = v
	return s
}

func (s *EnableTransitRouterRouteTablePropagationResponse) SetBody(v *EnableTransitRouterRouteTablePropagationResponseBody) *EnableTransitRouterRouteTablePropagationResponse {
	s.Body = v
	return s
}

type GrantInstanceToTransitRouterRequest struct {
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	CenOwnerId           *int64  `json:"CenOwnerId,omitempty" xml:"CenOwnerId,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType         *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	OrderType            *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GrantInstanceToTransitRouterRequest) String() string {
	return tea.Prettify(s)
}

func (s GrantInstanceToTransitRouterRequest) GoString() string {
	return s.String()
}

func (s *GrantInstanceToTransitRouterRequest) SetCenId(v string) *GrantInstanceToTransitRouterRequest {
	s.CenId = &v
	return s
}

func (s *GrantInstanceToTransitRouterRequest) SetCenOwnerId(v int64) *GrantInstanceToTransitRouterRequest {
	s.CenOwnerId = &v
	return s
}

func (s *GrantInstanceToTransitRouterRequest) SetInstanceId(v string) *GrantInstanceToTransitRouterRequest {
	s.InstanceId = &v
	return s
}

func (s *GrantInstanceToTransitRouterRequest) SetInstanceType(v string) *GrantInstanceToTransitRouterRequest {
	s.InstanceType = &v
	return s
}

func (s *GrantInstanceToTransitRouterRequest) SetOrderType(v string) *GrantInstanceToTransitRouterRequest {
	s.OrderType = &v
	return s
}

func (s *GrantInstanceToTransitRouterRequest) SetOwnerAccount(v string) *GrantInstanceToTransitRouterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GrantInstanceToTransitRouterRequest) SetOwnerId(v int64) *GrantInstanceToTransitRouterRequest {
	s.OwnerId = &v
	return s
}

func (s *GrantInstanceToTransitRouterRequest) SetRegionId(v string) *GrantInstanceToTransitRouterRequest {
	s.RegionId = &v
	return s
}

func (s *GrantInstanceToTransitRouterRequest) SetResourceOwnerAccount(v string) *GrantInstanceToTransitRouterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GrantInstanceToTransitRouterRequest) SetResourceOwnerId(v int64) *GrantInstanceToTransitRouterRequest {
	s.ResourceOwnerId = &v
	return s
}

type GrantInstanceToTransitRouterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GrantInstanceToTransitRouterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GrantInstanceToTransitRouterResponseBody) GoString() string {
	return s.String()
}

func (s *GrantInstanceToTransitRouterResponseBody) SetRequestId(v string) *GrantInstanceToTransitRouterResponseBody {
	s.RequestId = &v
	return s
}

type GrantInstanceToTransitRouterResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GrantInstanceToTransitRouterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GrantInstanceToTransitRouterResponse) String() string {
	return tea.Prettify(s)
}

func (s GrantInstanceToTransitRouterResponse) GoString() string {
	return s.String()
}

func (s *GrantInstanceToTransitRouterResponse) SetHeaders(v map[string]*string) *GrantInstanceToTransitRouterResponse {
	s.Headers = v
	return s
}

func (s *GrantInstanceToTransitRouterResponse) SetBody(v *GrantInstanceToTransitRouterResponseBody) *GrantInstanceToTransitRouterResponse {
	s.Body = v
	return s
}

type ListCenInterRegionTrafficQosPoliciesRequest struct {
	MaxResults                  *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken                   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount                *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount        *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId             *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TrafficQosPolicyDescription *string `json:"TrafficQosPolicyDescription,omitempty" xml:"TrafficQosPolicyDescription,omitempty"`
	TrafficQosPolicyId          *string `json:"TrafficQosPolicyId,omitempty" xml:"TrafficQosPolicyId,omitempty"`
	TrafficQosPolicyName        *string `json:"TrafficQosPolicyName,omitempty" xml:"TrafficQosPolicyName,omitempty"`
	TransitRouterAttachmentId   *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	TransitRouterId             *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s ListCenInterRegionTrafficQosPoliciesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCenInterRegionTrafficQosPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListCenInterRegionTrafficQosPoliciesRequest) SetMaxResults(v int32) *ListCenInterRegionTrafficQosPoliciesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesRequest) SetNextToken(v string) *ListCenInterRegionTrafficQosPoliciesRequest {
	s.NextToken = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesRequest) SetOwnerAccount(v string) *ListCenInterRegionTrafficQosPoliciesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesRequest) SetOwnerId(v int64) *ListCenInterRegionTrafficQosPoliciesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesRequest) SetResourceOwnerAccount(v string) *ListCenInterRegionTrafficQosPoliciesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesRequest) SetResourceOwnerId(v int64) *ListCenInterRegionTrafficQosPoliciesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesRequest) SetTrafficQosPolicyDescription(v string) *ListCenInterRegionTrafficQosPoliciesRequest {
	s.TrafficQosPolicyDescription = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesRequest) SetTrafficQosPolicyId(v string) *ListCenInterRegionTrafficQosPoliciesRequest {
	s.TrafficQosPolicyId = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesRequest) SetTrafficQosPolicyName(v string) *ListCenInterRegionTrafficQosPoliciesRequest {
	s.TrafficQosPolicyName = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesRequest) SetTransitRouterAttachmentId(v string) *ListCenInterRegionTrafficQosPoliciesRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesRequest) SetTransitRouterId(v string) *ListCenInterRegionTrafficQosPoliciesRequest {
	s.TransitRouterId = &v
	return s
}

type ListCenInterRegionTrafficQosPoliciesResponseBody struct {
	MaxResults         *int32                                                                `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken          *string                                                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId          *string                                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount         *int32                                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TrafficQosPolicies []*ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies `json:"TrafficQosPolicies,omitempty" xml:"TrafficQosPolicies,omitempty" type:"Repeated"`
}

func (s ListCenInterRegionTrafficQosPoliciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCenInterRegionTrafficQosPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBody) SetMaxResults(v int32) *ListCenInterRegionTrafficQosPoliciesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBody) SetNextToken(v string) *ListCenInterRegionTrafficQosPoliciesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBody) SetRequestId(v string) *ListCenInterRegionTrafficQosPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBody) SetTotalCount(v int32) *ListCenInterRegionTrafficQosPoliciesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBody) SetTrafficQosPolicies(v []*ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies) *ListCenInterRegionTrafficQosPoliciesResponseBody {
	s.TrafficQosPolicies = v
	return s
}

type ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies struct {
	TrafficQosPolicyDescription *string                                                                               `json:"TrafficQosPolicyDescription,omitempty" xml:"TrafficQosPolicyDescription,omitempty"`
	TrafficQosPolicyId          *string                                                                               `json:"TrafficQosPolicyId,omitempty" xml:"TrafficQosPolicyId,omitempty"`
	TrafficQosPolicyName        *string                                                                               `json:"TrafficQosPolicyName,omitempty" xml:"TrafficQosPolicyName,omitempty"`
	TrafficQosPolicyStatus      *string                                                                               `json:"TrafficQosPolicyStatus,omitempty" xml:"TrafficQosPolicyStatus,omitempty"`
	TrafficQosQueues            []*ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPoliciesTrafficQosQueues `json:"TrafficQosQueues,omitempty" xml:"TrafficQosQueues,omitempty" type:"Repeated"`
}

func (s ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies) String() string {
	return tea.Prettify(s)
}

func (s ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies) GoString() string {
	return s.String()
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies) SetTrafficQosPolicyDescription(v string) *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies {
	s.TrafficQosPolicyDescription = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies) SetTrafficQosPolicyId(v string) *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies {
	s.TrafficQosPolicyId = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies) SetTrafficQosPolicyName(v string) *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies {
	s.TrafficQosPolicyName = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies) SetTrafficQosPolicyStatus(v string) *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies {
	s.TrafficQosPolicyStatus = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies) SetTrafficQosQueues(v []*ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPoliciesTrafficQosQueues) *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies {
	s.TrafficQosQueues = v
	return s
}

type ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPoliciesTrafficQosQueues struct {
	Dscps                  []*int32 `json:"Dscps,omitempty" xml:"Dscps,omitempty" type:"Repeated"`
	QosQueueDescription    *string  `json:"QosQueueDescription,omitempty" xml:"QosQueueDescription,omitempty"`
	QosQueueId             *string  `json:"QosQueueId,omitempty" xml:"QosQueueId,omitempty"`
	QosQueueName           *string  `json:"QosQueueName,omitempty" xml:"QosQueueName,omitempty"`
	RemainBandwidthPercent *int32   `json:"RemainBandwidthPercent,omitempty" xml:"RemainBandwidthPercent,omitempty"`
}

func (s ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPoliciesTrafficQosQueues) String() string {
	return tea.Prettify(s)
}

func (s ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPoliciesTrafficQosQueues) GoString() string {
	return s.String()
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPoliciesTrafficQosQueues) SetDscps(v []*int32) *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPoliciesTrafficQosQueues {
	s.Dscps = v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPoliciesTrafficQosQueues) SetQosQueueDescription(v string) *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPoliciesTrafficQosQueues {
	s.QosQueueDescription = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPoliciesTrafficQosQueues) SetQosQueueId(v string) *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPoliciesTrafficQosQueues {
	s.QosQueueId = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPoliciesTrafficQosQueues) SetQosQueueName(v string) *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPoliciesTrafficQosQueues {
	s.QosQueueName = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPoliciesTrafficQosQueues) SetRemainBandwidthPercent(v int32) *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPoliciesTrafficQosQueues {
	s.RemainBandwidthPercent = &v
	return s
}

type ListCenInterRegionTrafficQosPoliciesResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListCenInterRegionTrafficQosPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCenInterRegionTrafficQosPoliciesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCenInterRegionTrafficQosPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListCenInterRegionTrafficQosPoliciesResponse) SetHeaders(v map[string]*string) *ListCenInterRegionTrafficQosPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesResponse) SetBody(v *ListCenInterRegionTrafficQosPoliciesResponseBody) *ListCenInterRegionTrafficQosPoliciesResponse {
	s.Body = v
	return s
}

type ListGrantVSwitchesToCenRequest struct {
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	VpcId                *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListGrantVSwitchesToCenRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGrantVSwitchesToCenRequest) GoString() string {
	return s.String()
}

func (s *ListGrantVSwitchesToCenRequest) SetCenId(v string) *ListGrantVSwitchesToCenRequest {
	s.CenId = &v
	return s
}

func (s *ListGrantVSwitchesToCenRequest) SetOwnerAccount(v string) *ListGrantVSwitchesToCenRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListGrantVSwitchesToCenRequest) SetOwnerId(v int64) *ListGrantVSwitchesToCenRequest {
	s.OwnerId = &v
	return s
}

func (s *ListGrantVSwitchesToCenRequest) SetPageNumber(v int32) *ListGrantVSwitchesToCenRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGrantVSwitchesToCenRequest) SetPageSize(v int32) *ListGrantVSwitchesToCenRequest {
	s.PageSize = &v
	return s
}

func (s *ListGrantVSwitchesToCenRequest) SetRegionId(v string) *ListGrantVSwitchesToCenRequest {
	s.RegionId = &v
	return s
}

func (s *ListGrantVSwitchesToCenRequest) SetResourceOwnerAccount(v string) *ListGrantVSwitchesToCenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListGrantVSwitchesToCenRequest) SetResourceOwnerId(v int64) *ListGrantVSwitchesToCenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListGrantVSwitchesToCenRequest) SetVpcId(v string) *ListGrantVSwitchesToCenRequest {
	s.VpcId = &v
	return s
}

func (s *ListGrantVSwitchesToCenRequest) SetZoneId(v string) *ListGrantVSwitchesToCenRequest {
	s.ZoneId = &v
	return s
}

type ListGrantVSwitchesToCenResponseBody struct {
	PageNumber *int32                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VSwitches  []*ListGrantVSwitchesToCenResponseBodyVSwitches `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
}

func (s ListGrantVSwitchesToCenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGrantVSwitchesToCenResponseBody) GoString() string {
	return s.String()
}

func (s *ListGrantVSwitchesToCenResponseBody) SetPageNumber(v int32) *ListGrantVSwitchesToCenResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListGrantVSwitchesToCenResponseBody) SetPageSize(v int32) *ListGrantVSwitchesToCenResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListGrantVSwitchesToCenResponseBody) SetRequestId(v string) *ListGrantVSwitchesToCenResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGrantVSwitchesToCenResponseBody) SetTotalCount(v int32) *ListGrantVSwitchesToCenResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListGrantVSwitchesToCenResponseBody) SetVSwitches(v []*ListGrantVSwitchesToCenResponseBodyVSwitches) *ListGrantVSwitchesToCenResponseBody {
	s.VSwitches = v
	return s
}

type ListGrantVSwitchesToCenResponseBodyVSwitches struct {
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId     *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId    *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListGrantVSwitchesToCenResponseBodyVSwitches) String() string {
	return tea.Prettify(s)
}

func (s ListGrantVSwitchesToCenResponseBodyVSwitches) GoString() string {
	return s.String()
}

func (s *ListGrantVSwitchesToCenResponseBodyVSwitches) SetVSwitchId(v string) *ListGrantVSwitchesToCenResponseBodyVSwitches {
	s.VSwitchId = &v
	return s
}

func (s *ListGrantVSwitchesToCenResponseBodyVSwitches) SetVpcId(v string) *ListGrantVSwitchesToCenResponseBodyVSwitches {
	s.VpcId = &v
	return s
}

func (s *ListGrantVSwitchesToCenResponseBodyVSwitches) SetZoneId(v string) *ListGrantVSwitchesToCenResponseBodyVSwitches {
	s.ZoneId = &v
	return s
}

type ListGrantVSwitchesToCenResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListGrantVSwitchesToCenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListGrantVSwitchesToCenResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGrantVSwitchesToCenResponse) GoString() string {
	return s.String()
}

func (s *ListGrantVSwitchesToCenResponse) SetHeaders(v map[string]*string) *ListGrantVSwitchesToCenResponse {
	s.Headers = v
	return s
}

func (s *ListGrantVSwitchesToCenResponse) SetBody(v *ListGrantVSwitchesToCenResponseBody) *ListGrantVSwitchesToCenResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	NextToken            *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount         *string                       `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64                        `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageSize             *int32                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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

func (s *ListTagResourcesRequest) SetPageSize(v int32) *ListTagResourcesRequest {
	s.PageSize = &v
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

type ListTrafficMarkingPoliciesRequest struct {
	MaxResults                      *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken                       *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount                    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount            *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId                 *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TrafficMarkingPolicyDescription *string `json:"TrafficMarkingPolicyDescription,omitempty" xml:"TrafficMarkingPolicyDescription,omitempty"`
	TrafficMarkingPolicyId          *string `json:"TrafficMarkingPolicyId,omitempty" xml:"TrafficMarkingPolicyId,omitempty"`
	TrafficMarkingPolicyName        *string `json:"TrafficMarkingPolicyName,omitempty" xml:"TrafficMarkingPolicyName,omitempty"`
	TransitRouterId                 *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s ListTrafficMarkingPoliciesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTrafficMarkingPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListTrafficMarkingPoliciesRequest) SetMaxResults(v int32) *ListTrafficMarkingPoliciesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTrafficMarkingPoliciesRequest) SetNextToken(v string) *ListTrafficMarkingPoliciesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTrafficMarkingPoliciesRequest) SetOwnerAccount(v string) *ListTrafficMarkingPoliciesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTrafficMarkingPoliciesRequest) SetOwnerId(v int64) *ListTrafficMarkingPoliciesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTrafficMarkingPoliciesRequest) SetResourceOwnerAccount(v string) *ListTrafficMarkingPoliciesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTrafficMarkingPoliciesRequest) SetResourceOwnerId(v int64) *ListTrafficMarkingPoliciesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTrafficMarkingPoliciesRequest) SetTrafficMarkingPolicyDescription(v string) *ListTrafficMarkingPoliciesRequest {
	s.TrafficMarkingPolicyDescription = &v
	return s
}

func (s *ListTrafficMarkingPoliciesRequest) SetTrafficMarkingPolicyId(v string) *ListTrafficMarkingPoliciesRequest {
	s.TrafficMarkingPolicyId = &v
	return s
}

func (s *ListTrafficMarkingPoliciesRequest) SetTrafficMarkingPolicyName(v string) *ListTrafficMarkingPoliciesRequest {
	s.TrafficMarkingPolicyName = &v
	return s
}

func (s *ListTrafficMarkingPoliciesRequest) SetTransitRouterId(v string) *ListTrafficMarkingPoliciesRequest {
	s.TransitRouterId = &v
	return s
}

type ListTrafficMarkingPoliciesResponseBody struct {
	MaxResults             *int32                                                          `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken              *string                                                         `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId              *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount             *int32                                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TrafficMarkingPolicies []*ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies `json:"TrafficMarkingPolicies,omitempty" xml:"TrafficMarkingPolicies,omitempty" type:"Repeated"`
}

func (s ListTrafficMarkingPoliciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTrafficMarkingPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTrafficMarkingPoliciesResponseBody) SetMaxResults(v int32) *ListTrafficMarkingPoliciesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTrafficMarkingPoliciesResponseBody) SetNextToken(v string) *ListTrafficMarkingPoliciesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTrafficMarkingPoliciesResponseBody) SetRequestId(v string) *ListTrafficMarkingPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTrafficMarkingPoliciesResponseBody) SetTotalCount(v int32) *ListTrafficMarkingPoliciesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTrafficMarkingPoliciesResponseBody) SetTrafficMarkingPolicies(v []*ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies) *ListTrafficMarkingPoliciesResponseBody {
	s.TrafficMarkingPolicies = v
	return s
}

type ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies struct {
	MarkingDscp                     *int32                                                                           `json:"MarkingDscp,omitempty" xml:"MarkingDscp,omitempty"`
	Priority                        *int32                                                                           `json:"Priority,omitempty" xml:"Priority,omitempty"`
	TrafficMarkingPolicyDescription *string                                                                          `json:"TrafficMarkingPolicyDescription,omitempty" xml:"TrafficMarkingPolicyDescription,omitempty"`
	TrafficMarkingPolicyId          *string                                                                          `json:"TrafficMarkingPolicyId,omitempty" xml:"TrafficMarkingPolicyId,omitempty"`
	TrafficMarkingPolicyName        *string                                                                          `json:"TrafficMarkingPolicyName,omitempty" xml:"TrafficMarkingPolicyName,omitempty"`
	TrafficMarkingPolicyStatus      *string                                                                          `json:"TrafficMarkingPolicyStatus,omitempty" xml:"TrafficMarkingPolicyStatus,omitempty"`
	TrafficMatchRules               []*ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules `json:"TrafficMatchRules,omitempty" xml:"TrafficMatchRules,omitempty" type:"Repeated"`
}

func (s ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies) String() string {
	return tea.Prettify(s)
}

func (s ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies) GoString() string {
	return s.String()
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies) SetMarkingDscp(v int32) *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies {
	s.MarkingDscp = &v
	return s
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies) SetPriority(v int32) *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies {
	s.Priority = &v
	return s
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies) SetTrafficMarkingPolicyDescription(v string) *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies {
	s.TrafficMarkingPolicyDescription = &v
	return s
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies) SetTrafficMarkingPolicyId(v string) *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies {
	s.TrafficMarkingPolicyId = &v
	return s
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies) SetTrafficMarkingPolicyName(v string) *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies {
	s.TrafficMarkingPolicyName = &v
	return s
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies) SetTrafficMarkingPolicyStatus(v string) *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies {
	s.TrafficMarkingPolicyStatus = &v
	return s
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies) SetTrafficMatchRules(v []*ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules) *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPolicies {
	s.TrafficMatchRules = v
	return s
}

type ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules struct {
	DstCidr                     *string  `json:"DstCidr,omitempty" xml:"DstCidr,omitempty"`
	DstPortRange                []*int32 `json:"DstPortRange,omitempty" xml:"DstPortRange,omitempty" type:"Repeated"`
	MatchDscp                   *int32   `json:"MatchDscp,omitempty" xml:"MatchDscp,omitempty"`
	Protocol                    *string  `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	SrcCidr                     *string  `json:"SrcCidr,omitempty" xml:"SrcCidr,omitempty"`
	SrcPortRange                []*int32 `json:"SrcPortRange,omitempty" xml:"SrcPortRange,omitempty" type:"Repeated"`
	TrafficMatchRuleDescription *string  `json:"TrafficMatchRuleDescription,omitempty" xml:"TrafficMatchRuleDescription,omitempty"`
	TrafficMatchRuleId          *string  `json:"TrafficMatchRuleId,omitempty" xml:"TrafficMatchRuleId,omitempty"`
	TrafficMatchRuleName        *string  `json:"TrafficMatchRuleName,omitempty" xml:"TrafficMatchRuleName,omitempty"`
	TrafficMatchRuleStatus      *string  `json:"TrafficMatchRuleStatus,omitempty" xml:"TrafficMatchRuleStatus,omitempty"`
}

func (s ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules) String() string {
	return tea.Prettify(s)
}

func (s ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules) GoString() string {
	return s.String()
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules) SetDstCidr(v string) *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules {
	s.DstCidr = &v
	return s
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules) SetDstPortRange(v []*int32) *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules {
	s.DstPortRange = v
	return s
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules) SetMatchDscp(v int32) *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules {
	s.MatchDscp = &v
	return s
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules) SetProtocol(v string) *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules {
	s.Protocol = &v
	return s
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules) SetSrcCidr(v string) *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules {
	s.SrcCidr = &v
	return s
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules) SetSrcPortRange(v []*int32) *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules {
	s.SrcPortRange = v
	return s
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules) SetTrafficMatchRuleDescription(v string) *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules {
	s.TrafficMatchRuleDescription = &v
	return s
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules) SetTrafficMatchRuleId(v string) *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules {
	s.TrafficMatchRuleId = &v
	return s
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules) SetTrafficMatchRuleName(v string) *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules {
	s.TrafficMatchRuleName = &v
	return s
}

func (s *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules) SetTrafficMatchRuleStatus(v string) *ListTrafficMarkingPoliciesResponseBodyTrafficMarkingPoliciesTrafficMatchRules {
	s.TrafficMatchRuleStatus = &v
	return s
}

type ListTrafficMarkingPoliciesResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTrafficMarkingPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTrafficMarkingPoliciesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTrafficMarkingPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListTrafficMarkingPoliciesResponse) SetHeaders(v map[string]*string) *ListTrafficMarkingPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListTrafficMarkingPoliciesResponse) SetBody(v *ListTrafficMarkingPoliciesResponseBody) *ListTrafficMarkingPoliciesResponse {
	s.Body = v
	return s
}

type ListTransitRouterAvailableResourceRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListTransitRouterAvailableResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTransitRouterAvailableResourceRequest) GoString() string {
	return s.String()
}

func (s *ListTransitRouterAvailableResourceRequest) SetOwnerAccount(v string) *ListTransitRouterAvailableResourceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTransitRouterAvailableResourceRequest) SetOwnerId(v int64) *ListTransitRouterAvailableResourceRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTransitRouterAvailableResourceRequest) SetRegionId(v string) *ListTransitRouterAvailableResourceRequest {
	s.RegionId = &v
	return s
}

func (s *ListTransitRouterAvailableResourceRequest) SetResourceOwnerAccount(v string) *ListTransitRouterAvailableResourceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTransitRouterAvailableResourceRequest) SetResourceOwnerId(v int64) *ListTransitRouterAvailableResourceRequest {
	s.ResourceOwnerId = &v
	return s
}

type ListTransitRouterAvailableResourceResponseBody struct {
	MasterZones []*string `json:"MasterZones,omitempty" xml:"MasterZones,omitempty" type:"Repeated"`
	RequestId   *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SlaveZones  []*string `json:"SlaveZones,omitempty" xml:"SlaveZones,omitempty" type:"Repeated"`
}

func (s ListTransitRouterAvailableResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTransitRouterAvailableResourceResponseBody) GoString() string {
	return s.String()
}

func (s *ListTransitRouterAvailableResourceResponseBody) SetMasterZones(v []*string) *ListTransitRouterAvailableResourceResponseBody {
	s.MasterZones = v
	return s
}

func (s *ListTransitRouterAvailableResourceResponseBody) SetRequestId(v string) *ListTransitRouterAvailableResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTransitRouterAvailableResourceResponseBody) SetSlaveZones(v []*string) *ListTransitRouterAvailableResourceResponseBody {
	s.SlaveZones = v
	return s
}

type ListTransitRouterAvailableResourceResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTransitRouterAvailableResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTransitRouterAvailableResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTransitRouterAvailableResourceResponse) GoString() string {
	return s.String()
}

func (s *ListTransitRouterAvailableResourceResponse) SetHeaders(v map[string]*string) *ListTransitRouterAvailableResourceResponse {
	s.Headers = v
	return s
}

func (s *ListTransitRouterAvailableResourceResponse) SetBody(v *ListTransitRouterAvailableResourceResponseBody) *ListTransitRouterAvailableResourceResponse {
	s.Body = v
	return s
}

type ListTransitRouterPeerAttachmentsRequest struct {
	CenId                     *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	MaxResults                *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken                 *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount              *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId                  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount      *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId           *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	TransitRouterId           *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s ListTransitRouterPeerAttachmentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTransitRouterPeerAttachmentsRequest) GoString() string {
	return s.String()
}

func (s *ListTransitRouterPeerAttachmentsRequest) SetCenId(v string) *ListTransitRouterPeerAttachmentsRequest {
	s.CenId = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsRequest) SetMaxResults(v int32) *ListTransitRouterPeerAttachmentsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsRequest) SetNextToken(v string) *ListTransitRouterPeerAttachmentsRequest {
	s.NextToken = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsRequest) SetOwnerAccount(v string) *ListTransitRouterPeerAttachmentsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsRequest) SetOwnerId(v int64) *ListTransitRouterPeerAttachmentsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsRequest) SetRegionId(v string) *ListTransitRouterPeerAttachmentsRequest {
	s.RegionId = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsRequest) SetResourceOwnerAccount(v string) *ListTransitRouterPeerAttachmentsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsRequest) SetResourceOwnerId(v int64) *ListTransitRouterPeerAttachmentsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsRequest) SetTransitRouterAttachmentId(v string) *ListTransitRouterPeerAttachmentsRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsRequest) SetTransitRouterId(v string) *ListTransitRouterPeerAttachmentsRequest {
	s.TransitRouterId = &v
	return s
}

type ListTransitRouterPeerAttachmentsResponseBody struct {
	MaxResults               *int32                                                                  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken                *string                                                                 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId                *string                                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount               *int32                                                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TransitRouterAttachments []*ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments `json:"TransitRouterAttachments,omitempty" xml:"TransitRouterAttachments,omitempty" type:"Repeated"`
}

func (s ListTransitRouterPeerAttachmentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTransitRouterPeerAttachmentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTransitRouterPeerAttachmentsResponseBody) SetMaxResults(v int32) *ListTransitRouterPeerAttachmentsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponseBody) SetNextToken(v string) *ListTransitRouterPeerAttachmentsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponseBody) SetRequestId(v string) *ListTransitRouterPeerAttachmentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponseBody) SetTotalCount(v int32) *ListTransitRouterPeerAttachmentsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponseBody) SetTransitRouterAttachments(v []*ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) *ListTransitRouterPeerAttachmentsResponseBody {
	s.TransitRouterAttachments = v
	return s
}

type ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments struct {
	AutoPublishRouteEnabled            *bool   `json:"AutoPublishRouteEnabled,omitempty" xml:"AutoPublishRouteEnabled,omitempty"`
	Bandwidth                          *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	BandwidthType                      *string `json:"BandwidthType,omitempty" xml:"BandwidthType,omitempty"`
	CenBandwidthPackageId              *string `json:"CenBandwidthPackageId,omitempty" xml:"CenBandwidthPackageId,omitempty"`
	CreationTime                       *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	GeographicSpanId                   *string `json:"GeographicSpanId,omitempty" xml:"GeographicSpanId,omitempty"`
	PeerTransitRouterId                *string `json:"PeerTransitRouterId,omitempty" xml:"PeerTransitRouterId,omitempty"`
	PeerTransitRouterOwnerId           *int64  `json:"PeerTransitRouterOwnerId,omitempty" xml:"PeerTransitRouterOwnerId,omitempty"`
	PeerTransitRouterRegionId          *string `json:"PeerTransitRouterRegionId,omitempty" xml:"PeerTransitRouterRegionId,omitempty"`
	RegionId                           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType                       *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Status                             *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TransitRouterAttachmentDescription *string `json:"TransitRouterAttachmentDescription,omitempty" xml:"TransitRouterAttachmentDescription,omitempty"`
	TransitRouterAttachmentId          *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	TransitRouterAttachmentName        *string `json:"TransitRouterAttachmentName,omitempty" xml:"TransitRouterAttachmentName,omitempty"`
	TransitRouterId                    *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) String() string {
	return tea.Prettify(s)
}

func (s ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) GoString() string {
	return s.String()
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) SetAutoPublishRouteEnabled(v bool) *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments {
	s.AutoPublishRouteEnabled = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) SetBandwidth(v int32) *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments {
	s.Bandwidth = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) SetBandwidthType(v string) *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments {
	s.BandwidthType = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) SetCenBandwidthPackageId(v string) *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments {
	s.CenBandwidthPackageId = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) SetCreationTime(v string) *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments {
	s.CreationTime = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) SetGeographicSpanId(v string) *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments {
	s.GeographicSpanId = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) SetPeerTransitRouterId(v string) *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments {
	s.PeerTransitRouterId = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) SetPeerTransitRouterOwnerId(v int64) *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments {
	s.PeerTransitRouterOwnerId = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) SetPeerTransitRouterRegionId(v string) *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments {
	s.PeerTransitRouterRegionId = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) SetRegionId(v string) *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments {
	s.RegionId = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) SetResourceType(v string) *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments {
	s.ResourceType = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) SetStatus(v string) *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments {
	s.Status = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) SetTransitRouterAttachmentDescription(v string) *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments {
	s.TransitRouterAttachmentDescription = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) SetTransitRouterAttachmentId(v string) *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) SetTransitRouterAttachmentName(v string) *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments {
	s.TransitRouterAttachmentName = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) SetTransitRouterId(v string) *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments {
	s.TransitRouterId = &v
	return s
}

type ListTransitRouterPeerAttachmentsResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTransitRouterPeerAttachmentsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTransitRouterPeerAttachmentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTransitRouterPeerAttachmentsResponse) GoString() string {
	return s.String()
}

func (s *ListTransitRouterPeerAttachmentsResponse) SetHeaders(v map[string]*string) *ListTransitRouterPeerAttachmentsResponse {
	s.Headers = v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponse) SetBody(v *ListTransitRouterPeerAttachmentsResponseBody) *ListTransitRouterPeerAttachmentsResponse {
	s.Body = v
	return s
}

type ListTransitRouterRouteEntriesRequest struct {
	MaxResults                                  *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken                                   *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount                                *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                                     *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount                        *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId                             *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TransitRouterRouteEntryDestinationCidrBlock *string   `json:"TransitRouterRouteEntryDestinationCidrBlock,omitempty" xml:"TransitRouterRouteEntryDestinationCidrBlock,omitempty"`
	TransitRouterRouteEntryIds                  []*string `json:"TransitRouterRouteEntryIds,omitempty" xml:"TransitRouterRouteEntryIds,omitempty" type:"Repeated"`
	TransitRouterRouteEntryNames                []*string `json:"TransitRouterRouteEntryNames,omitempty" xml:"TransitRouterRouteEntryNames,omitempty" type:"Repeated"`
	TransitRouterRouteEntryStatus               *string   `json:"TransitRouterRouteEntryStatus,omitempty" xml:"TransitRouterRouteEntryStatus,omitempty"`
	TransitRouterRouteTableId                   *string   `json:"TransitRouterRouteTableId,omitempty" xml:"TransitRouterRouteTableId,omitempty"`
}

func (s ListTransitRouterRouteEntriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTransitRouterRouteEntriesRequest) GoString() string {
	return s.String()
}

func (s *ListTransitRouterRouteEntriesRequest) SetMaxResults(v int32) *ListTransitRouterRouteEntriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTransitRouterRouteEntriesRequest) SetNextToken(v string) *ListTransitRouterRouteEntriesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTransitRouterRouteEntriesRequest) SetOwnerAccount(v string) *ListTransitRouterRouteEntriesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTransitRouterRouteEntriesRequest) SetOwnerId(v int64) *ListTransitRouterRouteEntriesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTransitRouterRouteEntriesRequest) SetResourceOwnerAccount(v string) *ListTransitRouterRouteEntriesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTransitRouterRouteEntriesRequest) SetResourceOwnerId(v int64) *ListTransitRouterRouteEntriesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTransitRouterRouteEntriesRequest) SetTransitRouterRouteEntryDestinationCidrBlock(v string) *ListTransitRouterRouteEntriesRequest {
	s.TransitRouterRouteEntryDestinationCidrBlock = &v
	return s
}

func (s *ListTransitRouterRouteEntriesRequest) SetTransitRouterRouteEntryIds(v []*string) *ListTransitRouterRouteEntriesRequest {
	s.TransitRouterRouteEntryIds = v
	return s
}

func (s *ListTransitRouterRouteEntriesRequest) SetTransitRouterRouteEntryNames(v []*string) *ListTransitRouterRouteEntriesRequest {
	s.TransitRouterRouteEntryNames = v
	return s
}

func (s *ListTransitRouterRouteEntriesRequest) SetTransitRouterRouteEntryStatus(v string) *ListTransitRouterRouteEntriesRequest {
	s.TransitRouterRouteEntryStatus = &v
	return s
}

func (s *ListTransitRouterRouteEntriesRequest) SetTransitRouterRouteTableId(v string) *ListTransitRouterRouteEntriesRequest {
	s.TransitRouterRouteTableId = &v
	return s
}

type ListTransitRouterRouteEntriesResponseBody struct {
	MaxResults                *int32                                                                `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken                 *string                                                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId                 *string                                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount                *int32                                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TransitRouterRouteEntries []*ListTransitRouterRouteEntriesResponseBodyTransitRouterRouteEntries `json:"TransitRouterRouteEntries,omitempty" xml:"TransitRouterRouteEntries,omitempty" type:"Repeated"`
}

func (s ListTransitRouterRouteEntriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTransitRouterRouteEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTransitRouterRouteEntriesResponseBody) SetMaxResults(v int32) *ListTransitRouterRouteEntriesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTransitRouterRouteEntriesResponseBody) SetNextToken(v string) *ListTransitRouterRouteEntriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTransitRouterRouteEntriesResponseBody) SetRequestId(v string) *ListTransitRouterRouteEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTransitRouterRouteEntriesResponseBody) SetTotalCount(v int32) *ListTransitRouterRouteEntriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTransitRouterRouteEntriesResponseBody) SetTransitRouterRouteEntries(v []*ListTransitRouterRouteEntriesResponseBodyTransitRouterRouteEntries) *ListTransitRouterRouteEntriesResponseBody {
	s.TransitRouterRouteEntries = v
	return s
}

type ListTransitRouterRouteEntriesResponseBodyTransitRouterRouteEntries struct {
	CreateTime                                  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	TransitRouterRouteEntryDescription          *string `json:"TransitRouterRouteEntryDescription,omitempty" xml:"TransitRouterRouteEntryDescription,omitempty"`
	TransitRouterRouteEntryDestinationCidrBlock *string `json:"TransitRouterRouteEntryDestinationCidrBlock,omitempty" xml:"TransitRouterRouteEntryDestinationCidrBlock,omitempty"`
	TransitRouterRouteEntryId                   *string `json:"TransitRouterRouteEntryId,omitempty" xml:"TransitRouterRouteEntryId,omitempty"`
	TransitRouterRouteEntryName                 *string `json:"TransitRouterRouteEntryName,omitempty" xml:"TransitRouterRouteEntryName,omitempty"`
	TransitRouterRouteEntryNextHopId            *string `json:"TransitRouterRouteEntryNextHopId,omitempty" xml:"TransitRouterRouteEntryNextHopId,omitempty"`
	TransitRouterRouteEntryNextHopType          *string `json:"TransitRouterRouteEntryNextHopType,omitempty" xml:"TransitRouterRouteEntryNextHopType,omitempty"`
	TransitRouterRouteEntryStatus               *string `json:"TransitRouterRouteEntryStatus,omitempty" xml:"TransitRouterRouteEntryStatus,omitempty"`
	TransitRouterRouteEntryType                 *string `json:"TransitRouterRouteEntryType,omitempty" xml:"TransitRouterRouteEntryType,omitempty"`
}

func (s ListTransitRouterRouteEntriesResponseBodyTransitRouterRouteEntries) String() string {
	return tea.Prettify(s)
}

func (s ListTransitRouterRouteEntriesResponseBodyTransitRouterRouteEntries) GoString() string {
	return s.String()
}

func (s *ListTransitRouterRouteEntriesResponseBodyTransitRouterRouteEntries) SetCreateTime(v string) *ListTransitRouterRouteEntriesResponseBodyTransitRouterRouteEntries {
	s.CreateTime = &v
	return s
}

func (s *ListTransitRouterRouteEntriesResponseBodyTransitRouterRouteEntries) SetTransitRouterRouteEntryDescription(v string) *ListTransitRouterRouteEntriesResponseBodyTransitRouterRouteEntries {
	s.TransitRouterRouteEntryDescription = &v
	return s
}

func (s *ListTransitRouterRouteEntriesResponseBodyTransitRouterRouteEntries) SetTransitRouterRouteEntryDestinationCidrBlock(v string) *ListTransitRouterRouteEntriesResponseBodyTransitRouterRouteEntries {
	s.TransitRouterRouteEntryDestinationCidrBlock = &v
	return s
}

func (s *ListTransitRouterRouteEntriesResponseBodyTransitRouterRouteEntries) SetTransitRouterRouteEntryId(v string) *ListTransitRouterRouteEntriesResponseBodyTransitRouterRouteEntries {
	s.TransitRouterRouteEntryId = &v
	return s
}

func (s *ListTransitRouterRouteEntriesResponseBodyTransitRouterRouteEntries) SetTransitRouterRouteEntryName(v string) *ListTransitRouterRouteEntriesResponseBodyTransitRouterRouteEntries {
	s.TransitRouterRouteEntryName = &v
	return s
}

func (s *ListTransitRouterRouteEntriesResponseBodyTransitRouterRouteEntries) SetTransitRouterRouteEntryNextHopId(v string) *ListTransitRouterRouteEntriesResponseBodyTransitRouterRouteEntries {
	s.TransitRouterRouteEntryNextHopId = &v
	return s
}

func (s *ListTransitRouterRouteEntriesResponseBodyTransitRouterRouteEntries) SetTransitRouterRouteEntryNextHopType(v string) *ListTransitRouterRouteEntriesResponseBodyTransitRouterRouteEntries {
	s.TransitRouterRouteEntryNextHopType = &v
	return s
}

func (s *ListTransitRouterRouteEntriesResponseBodyTransitRouterRouteEntries) SetTransitRouterRouteEntryStatus(v string) *ListTransitRouterRouteEntriesResponseBodyTransitRouterRouteEntries {
	s.TransitRouterRouteEntryStatus = &v
	return s
}

func (s *ListTransitRouterRouteEntriesResponseBodyTransitRouterRouteEntries) SetTransitRouterRouteEntryType(v string) *ListTransitRouterRouteEntriesResponseBodyTransitRouterRouteEntries {
	s.TransitRouterRouteEntryType = &v
	return s
}

type ListTransitRouterRouteEntriesResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTransitRouterRouteEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTransitRouterRouteEntriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTransitRouterRouteEntriesResponse) GoString() string {
	return s.String()
}

func (s *ListTransitRouterRouteEntriesResponse) SetHeaders(v map[string]*string) *ListTransitRouterRouteEntriesResponse {
	s.Headers = v
	return s
}

func (s *ListTransitRouterRouteEntriesResponse) SetBody(v *ListTransitRouterRouteEntriesResponseBody) *ListTransitRouterRouteEntriesResponse {
	s.Body = v
	return s
}

type ListTransitRouterRouteTableAssociationsRequest struct {
	MaxResults                *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken                 *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount              *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount      *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId           *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	TransitRouterRouteTableId *string `json:"TransitRouterRouteTableId,omitempty" xml:"TransitRouterRouteTableId,omitempty"`
}

func (s ListTransitRouterRouteTableAssociationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTransitRouterRouteTableAssociationsRequest) GoString() string {
	return s.String()
}

func (s *ListTransitRouterRouteTableAssociationsRequest) SetMaxResults(v int32) *ListTransitRouterRouteTableAssociationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTransitRouterRouteTableAssociationsRequest) SetNextToken(v string) *ListTransitRouterRouteTableAssociationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListTransitRouterRouteTableAssociationsRequest) SetOwnerAccount(v string) *ListTransitRouterRouteTableAssociationsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTransitRouterRouteTableAssociationsRequest) SetOwnerId(v int64) *ListTransitRouterRouteTableAssociationsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTransitRouterRouteTableAssociationsRequest) SetResourceOwnerAccount(v string) *ListTransitRouterRouteTableAssociationsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTransitRouterRouteTableAssociationsRequest) SetResourceOwnerId(v int64) *ListTransitRouterRouteTableAssociationsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTransitRouterRouteTableAssociationsRequest) SetTransitRouterAttachmentId(v string) *ListTransitRouterRouteTableAssociationsRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *ListTransitRouterRouteTableAssociationsRequest) SetTransitRouterRouteTableId(v string) *ListTransitRouterRouteTableAssociationsRequest {
	s.TransitRouterRouteTableId = &v
	return s
}

type ListTransitRouterRouteTableAssociationsResponseBody struct {
	MaxResults                *int32                                                                          `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken                 *string                                                                         `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId                 *string                                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount                *int32                                                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TransitRouterAssociations []*ListTransitRouterRouteTableAssociationsResponseBodyTransitRouterAssociations `json:"TransitRouterAssociations,omitempty" xml:"TransitRouterAssociations,omitempty" type:"Repeated"`
}

func (s ListTransitRouterRouteTableAssociationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTransitRouterRouteTableAssociationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTransitRouterRouteTableAssociationsResponseBody) SetMaxResults(v int32) *ListTransitRouterRouteTableAssociationsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTransitRouterRouteTableAssociationsResponseBody) SetNextToken(v string) *ListTransitRouterRouteTableAssociationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTransitRouterRouteTableAssociationsResponseBody) SetRequestId(v string) *ListTransitRouterRouteTableAssociationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTransitRouterRouteTableAssociationsResponseBody) SetTotalCount(v int32) *ListTransitRouterRouteTableAssociationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTransitRouterRouteTableAssociationsResponseBody) SetTransitRouterAssociations(v []*ListTransitRouterRouteTableAssociationsResponseBodyTransitRouterAssociations) *ListTransitRouterRouteTableAssociationsResponseBody {
	s.TransitRouterAssociations = v
	return s
}

type ListTransitRouterRouteTableAssociationsResponseBodyTransitRouterAssociations struct {
	ResourceId                *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType              *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Status                    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	TransitRouterRouteTableId *string `json:"TransitRouterRouteTableId,omitempty" xml:"TransitRouterRouteTableId,omitempty"`
}

func (s ListTransitRouterRouteTableAssociationsResponseBodyTransitRouterAssociations) String() string {
	return tea.Prettify(s)
}

func (s ListTransitRouterRouteTableAssociationsResponseBodyTransitRouterAssociations) GoString() string {
	return s.String()
}

func (s *ListTransitRouterRouteTableAssociationsResponseBodyTransitRouterAssociations) SetResourceId(v string) *ListTransitRouterRouteTableAssociationsResponseBodyTransitRouterAssociations {
	s.ResourceId = &v
	return s
}

func (s *ListTransitRouterRouteTableAssociationsResponseBodyTransitRouterAssociations) SetResourceType(v string) *ListTransitRouterRouteTableAssociationsResponseBodyTransitRouterAssociations {
	s.ResourceType = &v
	return s
}

func (s *ListTransitRouterRouteTableAssociationsResponseBodyTransitRouterAssociations) SetStatus(v string) *ListTransitRouterRouteTableAssociationsResponseBodyTransitRouterAssociations {
	s.Status = &v
	return s
}

func (s *ListTransitRouterRouteTableAssociationsResponseBodyTransitRouterAssociations) SetTransitRouterAttachmentId(v string) *ListTransitRouterRouteTableAssociationsResponseBodyTransitRouterAssociations {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *ListTransitRouterRouteTableAssociationsResponseBodyTransitRouterAssociations) SetTransitRouterRouteTableId(v string) *ListTransitRouterRouteTableAssociationsResponseBodyTransitRouterAssociations {
	s.TransitRouterRouteTableId = &v
	return s
}

type ListTransitRouterRouteTableAssociationsResponse struct {
	Headers map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTransitRouterRouteTableAssociationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTransitRouterRouteTableAssociationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTransitRouterRouteTableAssociationsResponse) GoString() string {
	return s.String()
}

func (s *ListTransitRouterRouteTableAssociationsResponse) SetHeaders(v map[string]*string) *ListTransitRouterRouteTableAssociationsResponse {
	s.Headers = v
	return s
}

func (s *ListTransitRouterRouteTableAssociationsResponse) SetBody(v *ListTransitRouterRouteTableAssociationsResponseBody) *ListTransitRouterRouteTableAssociationsResponse {
	s.Body = v
	return s
}

type ListTransitRouterRouteTablePropagationsRequest struct {
	MaxResults                *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken                 *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount              *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount      *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId           *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	TransitRouterRouteTableId *string `json:"TransitRouterRouteTableId,omitempty" xml:"TransitRouterRouteTableId,omitempty"`
}

func (s ListTransitRouterRouteTablePropagationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTransitRouterRouteTablePropagationsRequest) GoString() string {
	return s.String()
}

func (s *ListTransitRouterRouteTablePropagationsRequest) SetMaxResults(v int32) *ListTransitRouterRouteTablePropagationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTransitRouterRouteTablePropagationsRequest) SetNextToken(v string) *ListTransitRouterRouteTablePropagationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListTransitRouterRouteTablePropagationsRequest) SetOwnerAccount(v string) *ListTransitRouterRouteTablePropagationsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTransitRouterRouteTablePropagationsRequest) SetOwnerId(v int64) *ListTransitRouterRouteTablePropagationsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTransitRouterRouteTablePropagationsRequest) SetResourceOwnerAccount(v string) *ListTransitRouterRouteTablePropagationsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTransitRouterRouteTablePropagationsRequest) SetResourceOwnerId(v int64) *ListTransitRouterRouteTablePropagationsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTransitRouterRouteTablePropagationsRequest) SetTransitRouterAttachmentId(v string) *ListTransitRouterRouteTablePropagationsRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *ListTransitRouterRouteTablePropagationsRequest) SetTransitRouterRouteTableId(v string) *ListTransitRouterRouteTablePropagationsRequest {
	s.TransitRouterRouteTableId = &v
	return s
}

type ListTransitRouterRouteTablePropagationsResponseBody struct {
	MaxResults                *int32                                                                          `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken                 *string                                                                         `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId                 *string                                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount                *int32                                                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TransitRouterPropagations []*ListTransitRouterRouteTablePropagationsResponseBodyTransitRouterPropagations `json:"TransitRouterPropagations,omitempty" xml:"TransitRouterPropagations,omitempty" type:"Repeated"`
}

func (s ListTransitRouterRouteTablePropagationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTransitRouterRouteTablePropagationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTransitRouterRouteTablePropagationsResponseBody) SetMaxResults(v int32) *ListTransitRouterRouteTablePropagationsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTransitRouterRouteTablePropagationsResponseBody) SetNextToken(v string) *ListTransitRouterRouteTablePropagationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTransitRouterRouteTablePropagationsResponseBody) SetRequestId(v string) *ListTransitRouterRouteTablePropagationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTransitRouterRouteTablePropagationsResponseBody) SetTotalCount(v int32) *ListTransitRouterRouteTablePropagationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTransitRouterRouteTablePropagationsResponseBody) SetTransitRouterPropagations(v []*ListTransitRouterRouteTablePropagationsResponseBodyTransitRouterPropagations) *ListTransitRouterRouteTablePropagationsResponseBody {
	s.TransitRouterPropagations = v
	return s
}

type ListTransitRouterRouteTablePropagationsResponseBodyTransitRouterPropagations struct {
	ResourceId                *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType              *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Status                    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	TransitRouterRouteTableId *string `json:"TransitRouterRouteTableId,omitempty" xml:"TransitRouterRouteTableId,omitempty"`
}

func (s ListTransitRouterRouteTablePropagationsResponseBodyTransitRouterPropagations) String() string {
	return tea.Prettify(s)
}

func (s ListTransitRouterRouteTablePropagationsResponseBodyTransitRouterPropagations) GoString() string {
	return s.String()
}

func (s *ListTransitRouterRouteTablePropagationsResponseBodyTransitRouterPropagations) SetResourceId(v string) *ListTransitRouterRouteTablePropagationsResponseBodyTransitRouterPropagations {
	s.ResourceId = &v
	return s
}

func (s *ListTransitRouterRouteTablePropagationsResponseBodyTransitRouterPropagations) SetResourceType(v string) *ListTransitRouterRouteTablePropagationsResponseBodyTransitRouterPropagations {
	s.ResourceType = &v
	return s
}

func (s *ListTransitRouterRouteTablePropagationsResponseBodyTransitRouterPropagations) SetStatus(v string) *ListTransitRouterRouteTablePropagationsResponseBodyTransitRouterPropagations {
	s.Status = &v
	return s
}

func (s *ListTransitRouterRouteTablePropagationsResponseBodyTransitRouterPropagations) SetTransitRouterAttachmentId(v string) *ListTransitRouterRouteTablePropagationsResponseBodyTransitRouterPropagations {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *ListTransitRouterRouteTablePropagationsResponseBodyTransitRouterPropagations) SetTransitRouterRouteTableId(v string) *ListTransitRouterRouteTablePropagationsResponseBodyTransitRouterPropagations {
	s.TransitRouterRouteTableId = &v
	return s
}

type ListTransitRouterRouteTablePropagationsResponse struct {
	Headers map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTransitRouterRouteTablePropagationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTransitRouterRouteTablePropagationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTransitRouterRouteTablePropagationsResponse) GoString() string {
	return s.String()
}

func (s *ListTransitRouterRouteTablePropagationsResponse) SetHeaders(v map[string]*string) *ListTransitRouterRouteTablePropagationsResponse {
	s.Headers = v
	return s
}

func (s *ListTransitRouterRouteTablePropagationsResponse) SetBody(v *ListTransitRouterRouteTablePropagationsResponseBody) *ListTransitRouterRouteTablePropagationsResponse {
	s.Body = v
	return s
}

type ListTransitRouterRouteTablesRequest struct {
	MaxResults                    *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken                     *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount                  *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                       *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount          *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId               *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TransitRouterId               *string   `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	TransitRouterRouteTableIds    []*string `json:"TransitRouterRouteTableIds,omitempty" xml:"TransitRouterRouteTableIds,omitempty" type:"Repeated"`
	TransitRouterRouteTableNames  []*string `json:"TransitRouterRouteTableNames,omitempty" xml:"TransitRouterRouteTableNames,omitempty" type:"Repeated"`
	TransitRouterRouteTableStatus *string   `json:"TransitRouterRouteTableStatus,omitempty" xml:"TransitRouterRouteTableStatus,omitempty"`
	TransitRouterRouteTableType   *string   `json:"TransitRouterRouteTableType,omitempty" xml:"TransitRouterRouteTableType,omitempty"`
}

func (s ListTransitRouterRouteTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTransitRouterRouteTablesRequest) GoString() string {
	return s.String()
}

func (s *ListTransitRouterRouteTablesRequest) SetMaxResults(v int32) *ListTransitRouterRouteTablesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTransitRouterRouteTablesRequest) SetNextToken(v string) *ListTransitRouterRouteTablesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTransitRouterRouteTablesRequest) SetOwnerAccount(v string) *ListTransitRouterRouteTablesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTransitRouterRouteTablesRequest) SetOwnerId(v int64) *ListTransitRouterRouteTablesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTransitRouterRouteTablesRequest) SetResourceOwnerAccount(v string) *ListTransitRouterRouteTablesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTransitRouterRouteTablesRequest) SetResourceOwnerId(v int64) *ListTransitRouterRouteTablesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTransitRouterRouteTablesRequest) SetTransitRouterId(v string) *ListTransitRouterRouteTablesRequest {
	s.TransitRouterId = &v
	return s
}

func (s *ListTransitRouterRouteTablesRequest) SetTransitRouterRouteTableIds(v []*string) *ListTransitRouterRouteTablesRequest {
	s.TransitRouterRouteTableIds = v
	return s
}

func (s *ListTransitRouterRouteTablesRequest) SetTransitRouterRouteTableNames(v []*string) *ListTransitRouterRouteTablesRequest {
	s.TransitRouterRouteTableNames = v
	return s
}

func (s *ListTransitRouterRouteTablesRequest) SetTransitRouterRouteTableStatus(v string) *ListTransitRouterRouteTablesRequest {
	s.TransitRouterRouteTableStatus = &v
	return s
}

func (s *ListTransitRouterRouteTablesRequest) SetTransitRouterRouteTableType(v string) *ListTransitRouterRouteTablesRequest {
	s.TransitRouterRouteTableType = &v
	return s
}

type ListTransitRouterRouteTablesResponseBody struct {
	MaxResults               *int32                                                              `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken                *string                                                             `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId                *string                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount               *int32                                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TransitRouterRouteTables []*ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables `json:"TransitRouterRouteTables,omitempty" xml:"TransitRouterRouteTables,omitempty" type:"Repeated"`
}

func (s ListTransitRouterRouteTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTransitRouterRouteTablesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTransitRouterRouteTablesResponseBody) SetMaxResults(v int32) *ListTransitRouterRouteTablesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTransitRouterRouteTablesResponseBody) SetNextToken(v string) *ListTransitRouterRouteTablesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTransitRouterRouteTablesResponseBody) SetRequestId(v string) *ListTransitRouterRouteTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTransitRouterRouteTablesResponseBody) SetTotalCount(v int32) *ListTransitRouterRouteTablesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTransitRouterRouteTablesResponseBody) SetTransitRouterRouteTables(v []*ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables) *ListTransitRouterRouteTablesResponseBody {
	s.TransitRouterRouteTables = v
	return s
}

type ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables struct {
	CreateTime                         *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	TransitRouterRouteTableDescription *string `json:"TransitRouterRouteTableDescription,omitempty" xml:"TransitRouterRouteTableDescription,omitempty"`
	TransitRouterRouteTableId          *string `json:"TransitRouterRouteTableId,omitempty" xml:"TransitRouterRouteTableId,omitempty"`
	TransitRouterRouteTableName        *string `json:"TransitRouterRouteTableName,omitempty" xml:"TransitRouterRouteTableName,omitempty"`
	TransitRouterRouteTableStatus      *string `json:"TransitRouterRouteTableStatus,omitempty" xml:"TransitRouterRouteTableStatus,omitempty"`
	TransitRouterRouteTableType        *string `json:"TransitRouterRouteTableType,omitempty" xml:"TransitRouterRouteTableType,omitempty"`
}

func (s ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables) String() string {
	return tea.Prettify(s)
}

func (s ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables) GoString() string {
	return s.String()
}

func (s *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables) SetCreateTime(v string) *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables {
	s.CreateTime = &v
	return s
}

func (s *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables) SetTransitRouterRouteTableDescription(v string) *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables {
	s.TransitRouterRouteTableDescription = &v
	return s
}

func (s *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables) SetTransitRouterRouteTableId(v string) *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables {
	s.TransitRouterRouteTableId = &v
	return s
}

func (s *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables) SetTransitRouterRouteTableName(v string) *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables {
	s.TransitRouterRouteTableName = &v
	return s
}

func (s *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables) SetTransitRouterRouteTableStatus(v string) *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables {
	s.TransitRouterRouteTableStatus = &v
	return s
}

func (s *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables) SetTransitRouterRouteTableType(v string) *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables {
	s.TransitRouterRouteTableType = &v
	return s
}

type ListTransitRouterRouteTablesResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTransitRouterRouteTablesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTransitRouterRouteTablesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTransitRouterRouteTablesResponse) GoString() string {
	return s.String()
}

func (s *ListTransitRouterRouteTablesResponse) SetHeaders(v map[string]*string) *ListTransitRouterRouteTablesResponse {
	s.Headers = v
	return s
}

func (s *ListTransitRouterRouteTablesResponse) SetBody(v *ListTransitRouterRouteTablesResponseBody) *ListTransitRouterRouteTablesResponse {
	s.Body = v
	return s
}

type ListTransitRouterVbrAttachmentsRequest struct {
	CenId                     *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	MaxResults                *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken                 *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount              *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId                  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount      *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId           *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	TransitRouterId           *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s ListTransitRouterVbrAttachmentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTransitRouterVbrAttachmentsRequest) GoString() string {
	return s.String()
}

func (s *ListTransitRouterVbrAttachmentsRequest) SetCenId(v string) *ListTransitRouterVbrAttachmentsRequest {
	s.CenId = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsRequest) SetMaxResults(v int32) *ListTransitRouterVbrAttachmentsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsRequest) SetNextToken(v string) *ListTransitRouterVbrAttachmentsRequest {
	s.NextToken = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsRequest) SetOwnerAccount(v string) *ListTransitRouterVbrAttachmentsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsRequest) SetOwnerId(v int64) *ListTransitRouterVbrAttachmentsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsRequest) SetRegionId(v string) *ListTransitRouterVbrAttachmentsRequest {
	s.RegionId = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsRequest) SetResourceOwnerAccount(v string) *ListTransitRouterVbrAttachmentsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsRequest) SetResourceOwnerId(v int64) *ListTransitRouterVbrAttachmentsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsRequest) SetTransitRouterAttachmentId(v string) *ListTransitRouterVbrAttachmentsRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsRequest) SetTransitRouterId(v string) *ListTransitRouterVbrAttachmentsRequest {
	s.TransitRouterId = &v
	return s
}

type ListTransitRouterVbrAttachmentsResponseBody struct {
	MaxResults               *int32                                                                 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken                *string                                                                `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId                *string                                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount               *int32                                                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TransitRouterAttachments []*ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments `json:"TransitRouterAttachments,omitempty" xml:"TransitRouterAttachments,omitempty" type:"Repeated"`
}

func (s ListTransitRouterVbrAttachmentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTransitRouterVbrAttachmentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTransitRouterVbrAttachmentsResponseBody) SetMaxResults(v int32) *ListTransitRouterVbrAttachmentsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsResponseBody) SetNextToken(v string) *ListTransitRouterVbrAttachmentsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsResponseBody) SetRequestId(v string) *ListTransitRouterVbrAttachmentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsResponseBody) SetTotalCount(v int32) *ListTransitRouterVbrAttachmentsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsResponseBody) SetTransitRouterAttachments(v []*ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) *ListTransitRouterVbrAttachmentsResponseBody {
	s.TransitRouterAttachments = v
	return s
}

type ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments struct {
	AutoPublishRouteEnabled            *bool   `json:"AutoPublishRouteEnabled,omitempty" xml:"AutoPublishRouteEnabled,omitempty"`
	CreationTime                       *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	ResourceType                       *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Status                             *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TransitRouterAttachmentDescription *string `json:"TransitRouterAttachmentDescription,omitempty" xml:"TransitRouterAttachmentDescription,omitempty"`
	TransitRouterAttachmentId          *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	TransitRouterAttachmentName        *string `json:"TransitRouterAttachmentName,omitempty" xml:"TransitRouterAttachmentName,omitempty"`
	TransitRouterId                    *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	VbrId                              *string `json:"VbrId,omitempty" xml:"VbrId,omitempty"`
	VbrOwnerId                         *int64  `json:"VbrOwnerId,omitempty" xml:"VbrOwnerId,omitempty"`
	VbrRegionId                        *string `json:"VbrRegionId,omitempty" xml:"VbrRegionId,omitempty"`
}

func (s ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) String() string {
	return tea.Prettify(s)
}

func (s ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) GoString() string {
	return s.String()
}

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) SetAutoPublishRouteEnabled(v bool) *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments {
	s.AutoPublishRouteEnabled = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) SetCreationTime(v string) *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments {
	s.CreationTime = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) SetResourceType(v string) *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments {
	s.ResourceType = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) SetStatus(v string) *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments {
	s.Status = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) SetTransitRouterAttachmentDescription(v string) *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments {
	s.TransitRouterAttachmentDescription = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) SetTransitRouterAttachmentId(v string) *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) SetTransitRouterAttachmentName(v string) *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments {
	s.TransitRouterAttachmentName = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) SetTransitRouterId(v string) *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments {
	s.TransitRouterId = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) SetVbrId(v string) *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments {
	s.VbrId = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) SetVbrOwnerId(v int64) *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments {
	s.VbrOwnerId = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) SetVbrRegionId(v string) *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments {
	s.VbrRegionId = &v
	return s
}

type ListTransitRouterVbrAttachmentsResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTransitRouterVbrAttachmentsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTransitRouterVbrAttachmentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTransitRouterVbrAttachmentsResponse) GoString() string {
	return s.String()
}

func (s *ListTransitRouterVbrAttachmentsResponse) SetHeaders(v map[string]*string) *ListTransitRouterVbrAttachmentsResponse {
	s.Headers = v
	return s
}

func (s *ListTransitRouterVbrAttachmentsResponse) SetBody(v *ListTransitRouterVbrAttachmentsResponseBody) *ListTransitRouterVbrAttachmentsResponse {
	s.Body = v
	return s
}

type ListTransitRouterVpcAttachmentsRequest struct {
	CenId                     *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	MaxResults                *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken                 *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount              *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId                  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount      *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId           *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	TransitRouterId           *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s ListTransitRouterVpcAttachmentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTransitRouterVpcAttachmentsRequest) GoString() string {
	return s.String()
}

func (s *ListTransitRouterVpcAttachmentsRequest) SetCenId(v string) *ListTransitRouterVpcAttachmentsRequest {
	s.CenId = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsRequest) SetMaxResults(v int32) *ListTransitRouterVpcAttachmentsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsRequest) SetNextToken(v string) *ListTransitRouterVpcAttachmentsRequest {
	s.NextToken = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsRequest) SetOwnerAccount(v string) *ListTransitRouterVpcAttachmentsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsRequest) SetOwnerId(v int64) *ListTransitRouterVpcAttachmentsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsRequest) SetRegionId(v string) *ListTransitRouterVpcAttachmentsRequest {
	s.RegionId = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsRequest) SetResourceOwnerAccount(v string) *ListTransitRouterVpcAttachmentsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsRequest) SetResourceOwnerId(v int64) *ListTransitRouterVpcAttachmentsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsRequest) SetTransitRouterAttachmentId(v string) *ListTransitRouterVpcAttachmentsRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsRequest) SetTransitRouterId(v string) *ListTransitRouterVpcAttachmentsRequest {
	s.TransitRouterId = &v
	return s
}

type ListTransitRouterVpcAttachmentsResponseBody struct {
	MaxResults               *int32                                                                 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken                *string                                                                `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId                *string                                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount               *int32                                                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TransitRouterAttachments []*ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments `json:"TransitRouterAttachments,omitempty" xml:"TransitRouterAttachments,omitempty" type:"Repeated"`
}

func (s ListTransitRouterVpcAttachmentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTransitRouterVpcAttachmentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTransitRouterVpcAttachmentsResponseBody) SetMaxResults(v int32) *ListTransitRouterVpcAttachmentsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponseBody) SetNextToken(v string) *ListTransitRouterVpcAttachmentsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponseBody) SetRequestId(v string) *ListTransitRouterVpcAttachmentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponseBody) SetTotalCount(v int32) *ListTransitRouterVpcAttachmentsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponseBody) SetTransitRouterAttachments(v []*ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) *ListTransitRouterVpcAttachmentsResponseBody {
	s.TransitRouterAttachments = v
	return s
}

type ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments struct {
	CreationTime                       *string                                                                            `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	ResourceType                       *string                                                                            `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Status                             *string                                                                            `json:"Status,omitempty" xml:"Status,omitempty"`
	TransitRouterAttachmentDescription *string                                                                            `json:"TransitRouterAttachmentDescription,omitempty" xml:"TransitRouterAttachmentDescription,omitempty"`
	TransitRouterAttachmentId          *string                                                                            `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	TransitRouterAttachmentName        *string                                                                            `json:"TransitRouterAttachmentName,omitempty" xml:"TransitRouterAttachmentName,omitempty"`
	TransitRouterId                    *string                                                                            `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	VpcId                              *string                                                                            `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcOwnerId                         *int64                                                                             `json:"VpcOwnerId,omitempty" xml:"VpcOwnerId,omitempty"`
	VpcRegionId                        *string                                                                            `json:"VpcRegionId,omitempty" xml:"VpcRegionId,omitempty"`
	ZoneMappings                       []*ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsZoneMappings `json:"ZoneMappings,omitempty" xml:"ZoneMappings,omitempty" type:"Repeated"`
}

func (s ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) String() string {
	return tea.Prettify(s)
}

func (s ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) GoString() string {
	return s.String()
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) SetCreationTime(v string) *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments {
	s.CreationTime = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) SetResourceType(v string) *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments {
	s.ResourceType = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) SetStatus(v string) *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments {
	s.Status = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) SetTransitRouterAttachmentDescription(v string) *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments {
	s.TransitRouterAttachmentDescription = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) SetTransitRouterAttachmentId(v string) *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) SetTransitRouterAttachmentName(v string) *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments {
	s.TransitRouterAttachmentName = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) SetTransitRouterId(v string) *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments {
	s.TransitRouterId = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) SetVpcId(v string) *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments {
	s.VpcId = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) SetVpcOwnerId(v int64) *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments {
	s.VpcOwnerId = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) SetVpcRegionId(v string) *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments {
	s.VpcRegionId = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) SetZoneMappings(v []*ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsZoneMappings) *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments {
	s.ZoneMappings = v
	return s
}

type ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsZoneMappings struct {
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId    *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsZoneMappings) String() string {
	return tea.Prettify(s)
}

func (s ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsZoneMappings) GoString() string {
	return s.String()
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsZoneMappings) SetVSwitchId(v string) *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsZoneMappings {
	s.VSwitchId = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsZoneMappings) SetZoneId(v string) *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsZoneMappings {
	s.ZoneId = &v
	return s
}

type ListTransitRouterVpcAttachmentsResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTransitRouterVpcAttachmentsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTransitRouterVpcAttachmentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTransitRouterVpcAttachmentsResponse) GoString() string {
	return s.String()
}

func (s *ListTransitRouterVpcAttachmentsResponse) SetHeaders(v map[string]*string) *ListTransitRouterVpcAttachmentsResponse {
	s.Headers = v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponse) SetBody(v *ListTransitRouterVpcAttachmentsResponseBody) *ListTransitRouterVpcAttachmentsResponse {
	s.Body = v
	return s
}

type ListTransitRoutersRequest struct {
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TransitRouterId      *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s ListTransitRoutersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTransitRoutersRequest) GoString() string {
	return s.String()
}

func (s *ListTransitRoutersRequest) SetCenId(v string) *ListTransitRoutersRequest {
	s.CenId = &v
	return s
}

func (s *ListTransitRoutersRequest) SetOwnerAccount(v string) *ListTransitRoutersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTransitRoutersRequest) SetOwnerId(v int64) *ListTransitRoutersRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTransitRoutersRequest) SetPageNumber(v int32) *ListTransitRoutersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTransitRoutersRequest) SetPageSize(v int32) *ListTransitRoutersRequest {
	s.PageSize = &v
	return s
}

func (s *ListTransitRoutersRequest) SetRegionId(v string) *ListTransitRoutersRequest {
	s.RegionId = &v
	return s
}

func (s *ListTransitRoutersRequest) SetResourceOwnerAccount(v string) *ListTransitRoutersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTransitRoutersRequest) SetResourceOwnerId(v int64) *ListTransitRoutersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTransitRoutersRequest) SetTransitRouterId(v string) *ListTransitRoutersRequest {
	s.TransitRouterId = &v
	return s
}

type ListTransitRoutersResponseBody struct {
	PageNumber     *int32                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount     *int32                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TransitRouters []*ListTransitRoutersResponseBodyTransitRouters `json:"TransitRouters,omitempty" xml:"TransitRouters,omitempty" type:"Repeated"`
}

func (s ListTransitRoutersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTransitRoutersResponseBody) GoString() string {
	return s.String()
}

func (s *ListTransitRoutersResponseBody) SetPageNumber(v int32) *ListTransitRoutersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTransitRoutersResponseBody) SetPageSize(v int32) *ListTransitRoutersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTransitRoutersResponseBody) SetRequestId(v string) *ListTransitRoutersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTransitRoutersResponseBody) SetTotalCount(v int32) *ListTransitRoutersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTransitRoutersResponseBody) SetTransitRouters(v []*ListTransitRoutersResponseBodyTransitRouters) *ListTransitRoutersResponseBody {
	s.TransitRouters = v
	return s
}

type ListTransitRoutersResponseBodyTransitRouters struct {
	AliUid                   *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	CenId                    *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	CreationTime             *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status                   *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TransitRouterDescription *string `json:"TransitRouterDescription,omitempty" xml:"TransitRouterDescription,omitempty"`
	TransitRouterId          *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	TransitRouterName        *string `json:"TransitRouterName,omitempty" xml:"TransitRouterName,omitempty"`
	Type                     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListTransitRoutersResponseBodyTransitRouters) String() string {
	return tea.Prettify(s)
}

func (s ListTransitRoutersResponseBodyTransitRouters) GoString() string {
	return s.String()
}

func (s *ListTransitRoutersResponseBodyTransitRouters) SetAliUid(v int64) *ListTransitRoutersResponseBodyTransitRouters {
	s.AliUid = &v
	return s
}

func (s *ListTransitRoutersResponseBodyTransitRouters) SetCenId(v string) *ListTransitRoutersResponseBodyTransitRouters {
	s.CenId = &v
	return s
}

func (s *ListTransitRoutersResponseBodyTransitRouters) SetCreationTime(v string) *ListTransitRoutersResponseBodyTransitRouters {
	s.CreationTime = &v
	return s
}

func (s *ListTransitRoutersResponseBodyTransitRouters) SetRegionId(v string) *ListTransitRoutersResponseBodyTransitRouters {
	s.RegionId = &v
	return s
}

func (s *ListTransitRoutersResponseBodyTransitRouters) SetStatus(v string) *ListTransitRoutersResponseBodyTransitRouters {
	s.Status = &v
	return s
}

func (s *ListTransitRoutersResponseBodyTransitRouters) SetTransitRouterDescription(v string) *ListTransitRoutersResponseBodyTransitRouters {
	s.TransitRouterDescription = &v
	return s
}

func (s *ListTransitRoutersResponseBodyTransitRouters) SetTransitRouterId(v string) *ListTransitRoutersResponseBodyTransitRouters {
	s.TransitRouterId = &v
	return s
}

func (s *ListTransitRoutersResponseBodyTransitRouters) SetTransitRouterName(v string) *ListTransitRoutersResponseBodyTransitRouters {
	s.TransitRouterName = &v
	return s
}

func (s *ListTransitRoutersResponseBodyTransitRouters) SetType(v string) *ListTransitRoutersResponseBodyTransitRouters {
	s.Type = &v
	return s
}

type ListTransitRoutersResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTransitRoutersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTransitRoutersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTransitRoutersResponse) GoString() string {
	return s.String()
}

func (s *ListTransitRoutersResponse) SetHeaders(v map[string]*string) *ListTransitRoutersResponse {
	s.Headers = v
	return s
}

func (s *ListTransitRoutersResponse) SetBody(v *ListTransitRoutersResponseBody) *ListTransitRoutersResponse {
	s.Body = v
	return s
}

type ModifyCenAttributeRequest struct {
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProtectionLevel      *string `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyCenAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyCenAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyCenAttributeRequest) SetCenId(v string) *ModifyCenAttributeRequest {
	s.CenId = &v
	return s
}

func (s *ModifyCenAttributeRequest) SetDescription(v string) *ModifyCenAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyCenAttributeRequest) SetName(v string) *ModifyCenAttributeRequest {
	s.Name = &v
	return s
}

func (s *ModifyCenAttributeRequest) SetOwnerAccount(v string) *ModifyCenAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyCenAttributeRequest) SetOwnerId(v int64) *ModifyCenAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyCenAttributeRequest) SetProtectionLevel(v string) *ModifyCenAttributeRequest {
	s.ProtectionLevel = &v
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
	CenBandwidthPackageId *string `json:"CenBandwidthPackageId,omitempty" xml:"CenBandwidthPackageId,omitempty"`
	Description           *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name                  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyCenBandwidthPackageAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyCenBandwidthPackageAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyCenBandwidthPackageAttributeRequest) SetCenBandwidthPackageId(v string) *ModifyCenBandwidthPackageAttributeRequest {
	s.CenBandwidthPackageId = &v
	return s
}

func (s *ModifyCenBandwidthPackageAttributeRequest) SetDescription(v string) *ModifyCenBandwidthPackageAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyCenBandwidthPackageAttributeRequest) SetName(v string) *ModifyCenBandwidthPackageAttributeRequest {
	s.Name = &v
	return s
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
	Bandwidth             *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	CenBandwidthPackageId *string `json:"CenBandwidthPackageId,omitempty" xml:"CenBandwidthPackageId,omitempty"`
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyCenBandwidthPackageSpecRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyCenBandwidthPackageSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyCenBandwidthPackageSpecRequest) SetBandwidth(v int32) *ModifyCenBandwidthPackageSpecRequest {
	s.Bandwidth = &v
	return s
}

func (s *ModifyCenBandwidthPackageSpecRequest) SetCenBandwidthPackageId(v string) *ModifyCenBandwidthPackageSpecRequest {
	s.CenBandwidthPackageId = &v
	return s
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
	AsPathMatchMode                    *string   `json:"AsPathMatchMode,omitempty" xml:"AsPathMatchMode,omitempty"`
	CenId                              *string   `json:"CenId,omitempty" xml:"CenId,omitempty"`
	CenRegionId                        *string   `json:"CenRegionId,omitempty" xml:"CenRegionId,omitempty"`
	CidrMatchMode                      *string   `json:"CidrMatchMode,omitempty" xml:"CidrMatchMode,omitempty"`
	CommunityMatchMode                 *string   `json:"CommunityMatchMode,omitempty" xml:"CommunityMatchMode,omitempty"`
	CommunityOperateMode               *string   `json:"CommunityOperateMode,omitempty" xml:"CommunityOperateMode,omitempty"`
	Description                        *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	DestinationChildInstanceTypes      []*string `json:"DestinationChildInstanceTypes,omitempty" xml:"DestinationChildInstanceTypes,omitempty" type:"Repeated"`
	DestinationCidrBlocks              []*string `json:"DestinationCidrBlocks,omitempty" xml:"DestinationCidrBlocks,omitempty" type:"Repeated"`
	DestinationInstanceIds             []*string `json:"DestinationInstanceIds,omitempty" xml:"DestinationInstanceIds,omitempty" type:"Repeated"`
	DestinationInstanceIdsReverseMatch *bool     `json:"DestinationInstanceIdsReverseMatch,omitempty" xml:"DestinationInstanceIdsReverseMatch,omitempty"`
	DestinationRouteTableIds           []*string `json:"DestinationRouteTableIds,omitempty" xml:"DestinationRouteTableIds,omitempty" type:"Repeated"`
	MapResult                          *string   `json:"MapResult,omitempty" xml:"MapResult,omitempty"`
	MatchAsns                          []*int32  `json:"MatchAsns,omitempty" xml:"MatchAsns,omitempty" type:"Repeated"`
	MatchCommunitySet                  []*string `json:"MatchCommunitySet,omitempty" xml:"MatchCommunitySet,omitempty" type:"Repeated"`
	NextPriority                       *int32    `json:"NextPriority,omitempty" xml:"NextPriority,omitempty"`
	OperateCommunitySet                []*string `json:"OperateCommunitySet,omitempty" xml:"OperateCommunitySet,omitempty" type:"Repeated"`
	OwnerAccount                       *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                            *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Preference                         *int32    `json:"Preference,omitempty" xml:"Preference,omitempty"`
	PrependAsPath                      []*int64  `json:"PrependAsPath,omitempty" xml:"PrependAsPath,omitempty" type:"Repeated"`
	Priority                           *int32    `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ResourceOwnerAccount               *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId                    *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RouteMapId                         *string   `json:"RouteMapId,omitempty" xml:"RouteMapId,omitempty"`
	RouteTypes                         []*string `json:"RouteTypes,omitempty" xml:"RouteTypes,omitempty" type:"Repeated"`
	SourceChildInstanceTypes           []*string `json:"SourceChildInstanceTypes,omitempty" xml:"SourceChildInstanceTypes,omitempty" type:"Repeated"`
	SourceInstanceIds                  []*string `json:"SourceInstanceIds,omitempty" xml:"SourceInstanceIds,omitempty" type:"Repeated"`
	SourceInstanceIdsReverseMatch      *bool     `json:"SourceInstanceIdsReverseMatch,omitempty" xml:"SourceInstanceIdsReverseMatch,omitempty"`
	SourceRegionIds                    []*string `json:"SourceRegionIds,omitempty" xml:"SourceRegionIds,omitempty" type:"Repeated"`
	SourceRouteTableIds                []*string `json:"SourceRouteTableIds,omitempty" xml:"SourceRouteTableIds,omitempty" type:"Repeated"`
}

func (s ModifyCenRouteMapRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyCenRouteMapRequest) GoString() string {
	return s.String()
}

func (s *ModifyCenRouteMapRequest) SetAsPathMatchMode(v string) *ModifyCenRouteMapRequest {
	s.AsPathMatchMode = &v
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

func (s *ModifyCenRouteMapRequest) SetCidrMatchMode(v string) *ModifyCenRouteMapRequest {
	s.CidrMatchMode = &v
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

func (s *ModifyCenRouteMapRequest) SetDescription(v string) *ModifyCenRouteMapRequest {
	s.Description = &v
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

func (s *ModifyCenRouteMapRequest) SetDestinationInstanceIds(v []*string) *ModifyCenRouteMapRequest {
	s.DestinationInstanceIds = v
	return s
}

func (s *ModifyCenRouteMapRequest) SetDestinationInstanceIdsReverseMatch(v bool) *ModifyCenRouteMapRequest {
	s.DestinationInstanceIdsReverseMatch = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetDestinationRouteTableIds(v []*string) *ModifyCenRouteMapRequest {
	s.DestinationRouteTableIds = v
	return s
}

func (s *ModifyCenRouteMapRequest) SetMapResult(v string) *ModifyCenRouteMapRequest {
	s.MapResult = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetMatchAsns(v []*int32) *ModifyCenRouteMapRequest {
	s.MatchAsns = v
	return s
}

func (s *ModifyCenRouteMapRequest) SetMatchCommunitySet(v []*string) *ModifyCenRouteMapRequest {
	s.MatchCommunitySet = v
	return s
}

func (s *ModifyCenRouteMapRequest) SetNextPriority(v int32) *ModifyCenRouteMapRequest {
	s.NextPriority = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetOperateCommunitySet(v []*string) *ModifyCenRouteMapRequest {
	s.OperateCommunitySet = v
	return s
}

func (s *ModifyCenRouteMapRequest) SetOwnerAccount(v string) *ModifyCenRouteMapRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetOwnerId(v int64) *ModifyCenRouteMapRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetPreference(v int32) *ModifyCenRouteMapRequest {
	s.Preference = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetPrependAsPath(v []*int64) *ModifyCenRouteMapRequest {
	s.PrependAsPath = v
	return s
}

func (s *ModifyCenRouteMapRequest) SetPriority(v int32) *ModifyCenRouteMapRequest {
	s.Priority = &v
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

func (s *ModifyCenRouteMapRequest) SetRouteMapId(v string) *ModifyCenRouteMapRequest {
	s.RouteMapId = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetRouteTypes(v []*string) *ModifyCenRouteMapRequest {
	s.RouteTypes = v
	return s
}

func (s *ModifyCenRouteMapRequest) SetSourceChildInstanceTypes(v []*string) *ModifyCenRouteMapRequest {
	s.SourceChildInstanceTypes = v
	return s
}

func (s *ModifyCenRouteMapRequest) SetSourceInstanceIds(v []*string) *ModifyCenRouteMapRequest {
	s.SourceInstanceIds = v
	return s
}

func (s *ModifyCenRouteMapRequest) SetSourceInstanceIdsReverseMatch(v bool) *ModifyCenRouteMapRequest {
	s.SourceInstanceIdsReverseMatch = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetSourceRegionIds(v []*string) *ModifyCenRouteMapRequest {
	s.SourceRegionIds = v
	return s
}

func (s *ModifyCenRouteMapRequest) SetSourceRouteTableIds(v []*string) *ModifyCenRouteMapRequest {
	s.SourceRouteTableIds = v
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
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FlowLogId            *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
	FlowLogName          *string `json:"FlowLogName,omitempty" xml:"FlowLogName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyFlowLogAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowLogAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyFlowLogAttributeRequest) SetCenId(v string) *ModifyFlowLogAttributeRequest {
	s.CenId = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetClientToken(v string) *ModifyFlowLogAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetDescription(v string) *ModifyFlowLogAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetFlowLogId(v string) *ModifyFlowLogAttributeRequest {
	s.FlowLogId = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetFlowLogName(v string) *ModifyFlowLogAttributeRequest {
	s.FlowLogName = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetOwnerAccount(v string) *ModifyFlowLogAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetOwnerId(v int64) *ModifyFlowLogAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetRegionId(v string) *ModifyFlowLogAttributeRequest {
	s.RegionId = &v
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

type MoveResourceGroupRequest struct {
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun               *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	NewResourceGroupId   *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceId           *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceType         *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s MoveResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupRequest) SetClientToken(v string) *MoveResourceGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *MoveResourceGroupRequest) SetDryRun(v bool) *MoveResourceGroupRequest {
	s.DryRun = &v
	return s
}

func (s *MoveResourceGroupRequest) SetNewResourceGroupId(v string) *MoveResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetOwnerAccount(v string) *MoveResourceGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *MoveResourceGroupRequest) SetOwnerId(v int64) *MoveResourceGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceId(v string) *MoveResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceOwnerAccount(v string) *MoveResourceGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceOwnerId(v int64) *MoveResourceGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceType(v string) *MoveResourceGroupRequest {
	s.ResourceType = &v
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

type OpenTransitRouterServiceRequest struct {
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s OpenTransitRouterServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenTransitRouterServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenTransitRouterServiceRequest) SetClientToken(v string) *OpenTransitRouterServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *OpenTransitRouterServiceRequest) SetOwnerAccount(v string) *OpenTransitRouterServiceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *OpenTransitRouterServiceRequest) SetOwnerId(v int64) *OpenTransitRouterServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *OpenTransitRouterServiceRequest) SetResourceOwnerAccount(v string) *OpenTransitRouterServiceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *OpenTransitRouterServiceRequest) SetResourceOwnerId(v int64) *OpenTransitRouterServiceRequest {
	s.ResourceOwnerId = &v
	return s
}

type OpenTransitRouterServiceResponseBody struct {
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenTransitRouterServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenTransitRouterServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenTransitRouterServiceResponseBody) SetOrderId(v string) *OpenTransitRouterServiceResponseBody {
	s.OrderId = &v
	return s
}

func (s *OpenTransitRouterServiceResponseBody) SetRequestId(v string) *OpenTransitRouterServiceResponseBody {
	s.RequestId = &v
	return s
}

type OpenTransitRouterServiceResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OpenTransitRouterServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpenTransitRouterServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenTransitRouterServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenTransitRouterServiceResponse) SetHeaders(v map[string]*string) *OpenTransitRouterServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenTransitRouterServiceResponse) SetBody(v *OpenTransitRouterServiceResponseBody) *OpenTransitRouterServiceResponse {
	s.Body = v
	return s
}

type PublishRouteEntriesRequest struct {
	CenId                     *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ChildInstanceId           *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	ChildInstanceRegionId     *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	ChildInstanceRouteTableId *string `json:"ChildInstanceRouteTableId,omitempty" xml:"ChildInstanceRouteTableId,omitempty"`
	ChildInstanceType         *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	DestinationCidrBlock      *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	ResourceOwnerAccount      *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId           *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s PublishRouteEntriesRequest) String() string {
	return tea.Prettify(s)
}

func (s PublishRouteEntriesRequest) GoString() string {
	return s.String()
}

func (s *PublishRouteEntriesRequest) SetCenId(v string) *PublishRouteEntriesRequest {
	s.CenId = &v
	return s
}

func (s *PublishRouteEntriesRequest) SetChildInstanceId(v string) *PublishRouteEntriesRequest {
	s.ChildInstanceId = &v
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

func (s *PublishRouteEntriesRequest) SetChildInstanceType(v string) *PublishRouteEntriesRequest {
	s.ChildInstanceType = &v
	return s
}

func (s *PublishRouteEntriesRequest) SetDestinationCidrBlock(v string) *PublishRouteEntriesRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *PublishRouteEntriesRequest) SetResourceOwnerAccount(v string) *PublishRouteEntriesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PublishRouteEntriesRequest) SetResourceOwnerId(v int64) *PublishRouteEntriesRequest {
	s.ResourceOwnerId = &v
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

type RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest struct {
	ClientToken            *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                 *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount           *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount   *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId        *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TrafficMarkRuleIds     []*string `json:"TrafficMarkRuleIds,omitempty" xml:"TrafficMarkRuleIds,omitempty" type:"Repeated"`
	TrafficMarkingPolicyId *string   `json:"TrafficMarkingPolicyId,omitempty" xml:"TrafficMarkingPolicyId,omitempty"`
}

func (s RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest) GoString() string {
	return s.String()
}

func (s *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest) SetClientToken(v string) *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest) SetDryRun(v bool) *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest {
	s.DryRun = &v
	return s
}

func (s *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest) SetOwnerAccount(v string) *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest) SetOwnerId(v int64) *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest) SetResourceOwnerAccount(v string) *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest) SetResourceOwnerId(v int64) *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest) SetTrafficMarkRuleIds(v []*string) *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest {
	s.TrafficMarkRuleIds = v
	return s
}

func (s *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest) SetTrafficMarkingPolicyId(v string) *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest {
	s.TrafficMarkingPolicyId = &v
	return s
}

type RemoveTraficMatchRuleFromTrafficMarkingPolicyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveTraficMatchRuleFromTrafficMarkingPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveTraficMatchRuleFromTrafficMarkingPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveTraficMatchRuleFromTrafficMarkingPolicyResponseBody) SetRequestId(v string) *RemoveTraficMatchRuleFromTrafficMarkingPolicyResponseBody {
	s.RequestId = &v
	return s
}

type RemoveTraficMatchRuleFromTrafficMarkingPolicyResponse struct {
	Headers map[string]*string                                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveTraficMatchRuleFromTrafficMarkingPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveTraficMatchRuleFromTrafficMarkingPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveTraficMatchRuleFromTrafficMarkingPolicyResponse) GoString() string {
	return s.String()
}

func (s *RemoveTraficMatchRuleFromTrafficMarkingPolicyResponse) SetHeaders(v map[string]*string) *RemoveTraficMatchRuleFromTrafficMarkingPolicyResponse {
	s.Headers = v
	return s
}

func (s *RemoveTraficMatchRuleFromTrafficMarkingPolicyResponse) SetBody(v *RemoveTraficMatchRuleFromTrafficMarkingPolicyResponseBody) *RemoveTraficMatchRuleFromTrafficMarkingPolicyResponse {
	s.Body = v
	return s
}

type ReplaceTransitRouterRouteTableAssociationRequest struct {
	ClientToken               *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                    *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount              *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount      *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId           *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	TransitRouterRouteTableId *string `json:"TransitRouterRouteTableId,omitempty" xml:"TransitRouterRouteTableId,omitempty"`
}

func (s ReplaceTransitRouterRouteTableAssociationRequest) String() string {
	return tea.Prettify(s)
}

func (s ReplaceTransitRouterRouteTableAssociationRequest) GoString() string {
	return s.String()
}

func (s *ReplaceTransitRouterRouteTableAssociationRequest) SetClientToken(v string) *ReplaceTransitRouterRouteTableAssociationRequest {
	s.ClientToken = &v
	return s
}

func (s *ReplaceTransitRouterRouteTableAssociationRequest) SetDryRun(v bool) *ReplaceTransitRouterRouteTableAssociationRequest {
	s.DryRun = &v
	return s
}

func (s *ReplaceTransitRouterRouteTableAssociationRequest) SetOwnerAccount(v string) *ReplaceTransitRouterRouteTableAssociationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ReplaceTransitRouterRouteTableAssociationRequest) SetOwnerId(v int64) *ReplaceTransitRouterRouteTableAssociationRequest {
	s.OwnerId = &v
	return s
}

func (s *ReplaceTransitRouterRouteTableAssociationRequest) SetResourceOwnerAccount(v string) *ReplaceTransitRouterRouteTableAssociationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReplaceTransitRouterRouteTableAssociationRequest) SetResourceOwnerId(v int64) *ReplaceTransitRouterRouteTableAssociationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReplaceTransitRouterRouteTableAssociationRequest) SetTransitRouterAttachmentId(v string) *ReplaceTransitRouterRouteTableAssociationRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *ReplaceTransitRouterRouteTableAssociationRequest) SetTransitRouterRouteTableId(v string) *ReplaceTransitRouterRouteTableAssociationRequest {
	s.TransitRouterRouteTableId = &v
	return s
}

type ReplaceTransitRouterRouteTableAssociationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReplaceTransitRouterRouteTableAssociationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReplaceTransitRouterRouteTableAssociationResponseBody) GoString() string {
	return s.String()
}

func (s *ReplaceTransitRouterRouteTableAssociationResponseBody) SetRequestId(v string) *ReplaceTransitRouterRouteTableAssociationResponseBody {
	s.RequestId = &v
	return s
}

type ReplaceTransitRouterRouteTableAssociationResponse struct {
	Headers map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReplaceTransitRouterRouteTableAssociationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReplaceTransitRouterRouteTableAssociationResponse) String() string {
	return tea.Prettify(s)
}

func (s ReplaceTransitRouterRouteTableAssociationResponse) GoString() string {
	return s.String()
}

func (s *ReplaceTransitRouterRouteTableAssociationResponse) SetHeaders(v map[string]*string) *ReplaceTransitRouterRouteTableAssociationResponse {
	s.Headers = v
	return s
}

func (s *ReplaceTransitRouterRouteTableAssociationResponse) SetBody(v *ReplaceTransitRouterRouteTableAssociationResponseBody) *ReplaceTransitRouterRouteTableAssociationResponse {
	s.Body = v
	return s
}

type ResolveAndRouteServiceInCenRequest struct {
	AccessRegionIds      []*string `json:"AccessRegionIds,omitempty" xml:"AccessRegionIds,omitempty" type:"Repeated"`
	CenId                *string   `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ClientToken          *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description          *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	Host                 *string   `json:"Host,omitempty" xml:"Host,omitempty"`
	HostRegionId         *string   `json:"HostRegionId,omitempty" xml:"HostRegionId,omitempty"`
	HostVpcId            *string   `json:"HostVpcId,omitempty" xml:"HostVpcId,omitempty"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ResolveAndRouteServiceInCenRequest) String() string {
	return tea.Prettify(s)
}

func (s ResolveAndRouteServiceInCenRequest) GoString() string {
	return s.String()
}

func (s *ResolveAndRouteServiceInCenRequest) SetAccessRegionIds(v []*string) *ResolveAndRouteServiceInCenRequest {
	s.AccessRegionIds = v
	return s
}

func (s *ResolveAndRouteServiceInCenRequest) SetCenId(v string) *ResolveAndRouteServiceInCenRequest {
	s.CenId = &v
	return s
}

func (s *ResolveAndRouteServiceInCenRequest) SetClientToken(v string) *ResolveAndRouteServiceInCenRequest {
	s.ClientToken = &v
	return s
}

func (s *ResolveAndRouteServiceInCenRequest) SetDescription(v string) *ResolveAndRouteServiceInCenRequest {
	s.Description = &v
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

type RevokeInstanceFromTransitRouterRequest struct {
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	CenOwnerId           *int64  `json:"CenOwnerId,omitempty" xml:"CenOwnerId,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType         *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RevokeInstanceFromTransitRouterRequest) String() string {
	return tea.Prettify(s)
}

func (s RevokeInstanceFromTransitRouterRequest) GoString() string {
	return s.String()
}

func (s *RevokeInstanceFromTransitRouterRequest) SetCenId(v string) *RevokeInstanceFromTransitRouterRequest {
	s.CenId = &v
	return s
}

func (s *RevokeInstanceFromTransitRouterRequest) SetCenOwnerId(v int64) *RevokeInstanceFromTransitRouterRequest {
	s.CenOwnerId = &v
	return s
}

func (s *RevokeInstanceFromTransitRouterRequest) SetInstanceId(v string) *RevokeInstanceFromTransitRouterRequest {
	s.InstanceId = &v
	return s
}

func (s *RevokeInstanceFromTransitRouterRequest) SetInstanceType(v string) *RevokeInstanceFromTransitRouterRequest {
	s.InstanceType = &v
	return s
}

func (s *RevokeInstanceFromTransitRouterRequest) SetOwnerAccount(v string) *RevokeInstanceFromTransitRouterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RevokeInstanceFromTransitRouterRequest) SetOwnerId(v int64) *RevokeInstanceFromTransitRouterRequest {
	s.OwnerId = &v
	return s
}

func (s *RevokeInstanceFromTransitRouterRequest) SetRegionId(v string) *RevokeInstanceFromTransitRouterRequest {
	s.RegionId = &v
	return s
}

func (s *RevokeInstanceFromTransitRouterRequest) SetResourceOwnerAccount(v string) *RevokeInstanceFromTransitRouterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RevokeInstanceFromTransitRouterRequest) SetResourceOwnerId(v int64) *RevokeInstanceFromTransitRouterRequest {
	s.ResourceOwnerId = &v
	return s
}

type RevokeInstanceFromTransitRouterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeInstanceFromTransitRouterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RevokeInstanceFromTransitRouterResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeInstanceFromTransitRouterResponseBody) SetRequestId(v string) *RevokeInstanceFromTransitRouterResponseBody {
	s.RequestId = &v
	return s
}

type RevokeInstanceFromTransitRouterResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RevokeInstanceFromTransitRouterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RevokeInstanceFromTransitRouterResponse) String() string {
	return tea.Prettify(s)
}

func (s RevokeInstanceFromTransitRouterResponse) GoString() string {
	return s.String()
}

func (s *RevokeInstanceFromTransitRouterResponse) SetHeaders(v map[string]*string) *RevokeInstanceFromTransitRouterResponse {
	s.Headers = v
	return s
}

func (s *RevokeInstanceFromTransitRouterResponse) SetBody(v *RevokeInstanceFromTransitRouterResponseBody) *RevokeInstanceFromTransitRouterResponse {
	s.Body = v
	return s
}

type RoutePrivateZoneInCenToVpcRequest struct {
	AccessRegionId       *string `json:"AccessRegionId,omitempty" xml:"AccessRegionId,omitempty"`
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	HostRegionId         *string `json:"HostRegionId,omitempty" xml:"HostRegionId,omitempty"`
	HostVpcId            *string `json:"HostVpcId,omitempty" xml:"HostVpcId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RoutePrivateZoneInCenToVpcRequest) String() string {
	return tea.Prettify(s)
}

func (s RoutePrivateZoneInCenToVpcRequest) GoString() string {
	return s.String()
}

func (s *RoutePrivateZoneInCenToVpcRequest) SetAccessRegionId(v string) *RoutePrivateZoneInCenToVpcRequest {
	s.AccessRegionId = &v
	return s
}

func (s *RoutePrivateZoneInCenToVpcRequest) SetCenId(v string) *RoutePrivateZoneInCenToVpcRequest {
	s.CenId = &v
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
	BandwidthLimit       *int64  `json:"BandwidthLimit,omitempty" xml:"BandwidthLimit,omitempty"`
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	LocalRegionId        *string `json:"LocalRegionId,omitempty" xml:"LocalRegionId,omitempty"`
	OppositeRegionId     *string `json:"OppositeRegionId,omitempty" xml:"OppositeRegionId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s SetCenInterRegionBandwidthLimitRequest) String() string {
	return tea.Prettify(s)
}

func (s SetCenInterRegionBandwidthLimitRequest) GoString() string {
	return s.String()
}

func (s *SetCenInterRegionBandwidthLimitRequest) SetBandwidthLimit(v int64) *SetCenInterRegionBandwidthLimitRequest {
	s.BandwidthLimit = &v
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
	Bandwidth             *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	CenBandwidthPackageId *string `json:"CenBandwidthPackageId,omitempty" xml:"CenBandwidthPackageId,omitempty"`
	EndTime               *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s TempUpgradeCenBandwidthPackageSpecRequest) String() string {
	return tea.Prettify(s)
}

func (s TempUpgradeCenBandwidthPackageSpecRequest) GoString() string {
	return s.String()
}

func (s *TempUpgradeCenBandwidthPackageSpecRequest) SetBandwidth(v int32) *TempUpgradeCenBandwidthPackageSpecRequest {
	s.Bandwidth = &v
	return s
}

func (s *TempUpgradeCenBandwidthPackageSpecRequest) SetCenBandwidthPackageId(v string) *TempUpgradeCenBandwidthPackageSpecRequest {
	s.CenBandwidthPackageId = &v
	return s
}

func (s *TempUpgradeCenBandwidthPackageSpecRequest) SetEndTime(v string) *TempUpgradeCenBandwidthPackageSpecRequest {
	s.EndTime = &v
	return s
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
	CenBandwidthPackageId *string `json:"CenBandwidthPackageId,omitempty" xml:"CenBandwidthPackageId,omitempty"`
	CenId                 *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UnassociateCenBandwidthPackageRequest) String() string {
	return tea.Prettify(s)
}

func (s UnassociateCenBandwidthPackageRequest) GoString() string {
	return s.String()
}

func (s *UnassociateCenBandwidthPackageRequest) SetCenBandwidthPackageId(v string) *UnassociateCenBandwidthPackageRequest {
	s.CenBandwidthPackageId = &v
	return s
}

func (s *UnassociateCenBandwidthPackageRequest) SetCenId(v string) *UnassociateCenBandwidthPackageRequest {
	s.CenId = &v
	return s
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
	AccessRegionId       *string `json:"AccessRegionId,omitempty" xml:"AccessRegionId,omitempty"`
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UnroutePrivateZoneInCenToVpcRequest) String() string {
	return tea.Prettify(s)
}

func (s UnroutePrivateZoneInCenToVpcRequest) GoString() string {
	return s.String()
}

func (s *UnroutePrivateZoneInCenToVpcRequest) SetAccessRegionId(v string) *UnroutePrivateZoneInCenToVpcRequest {
	s.AccessRegionId = &v
	return s
}

func (s *UnroutePrivateZoneInCenToVpcRequest) SetCenId(v string) *UnroutePrivateZoneInCenToVpcRequest {
	s.CenId = &v
	return s
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
	All                  *bool     `json:"All,omitempty" xml:"All,omitempty"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

type UpdateCenInterRegionTrafficQosPolicyAttributeRequest struct {
	ClientToken                 *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount                *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount        *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId             *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TrafficQosPolicyDescription *string `json:"TrafficQosPolicyDescription,omitempty" xml:"TrafficQosPolicyDescription,omitempty"`
	TrafficQosPolicyId          *string `json:"TrafficQosPolicyId,omitempty" xml:"TrafficQosPolicyId,omitempty"`
	TrafficQosPolicyName        *string `json:"TrafficQosPolicyName,omitempty" xml:"TrafficQosPolicyName,omitempty"`
}

func (s UpdateCenInterRegionTrafficQosPolicyAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCenInterRegionTrafficQosPolicyAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateCenInterRegionTrafficQosPolicyAttributeRequest) SetClientToken(v string) *UpdateCenInterRegionTrafficQosPolicyAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCenInterRegionTrafficQosPolicyAttributeRequest) SetDryRun(v bool) *UpdateCenInterRegionTrafficQosPolicyAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateCenInterRegionTrafficQosPolicyAttributeRequest) SetOwnerAccount(v string) *UpdateCenInterRegionTrafficQosPolicyAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateCenInterRegionTrafficQosPolicyAttributeRequest) SetOwnerId(v int64) *UpdateCenInterRegionTrafficQosPolicyAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateCenInterRegionTrafficQosPolicyAttributeRequest) SetResourceOwnerAccount(v string) *UpdateCenInterRegionTrafficQosPolicyAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateCenInterRegionTrafficQosPolicyAttributeRequest) SetResourceOwnerId(v int64) *UpdateCenInterRegionTrafficQosPolicyAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateCenInterRegionTrafficQosPolicyAttributeRequest) SetTrafficQosPolicyDescription(v string) *UpdateCenInterRegionTrafficQosPolicyAttributeRequest {
	s.TrafficQosPolicyDescription = &v
	return s
}

func (s *UpdateCenInterRegionTrafficQosPolicyAttributeRequest) SetTrafficQosPolicyId(v string) *UpdateCenInterRegionTrafficQosPolicyAttributeRequest {
	s.TrafficQosPolicyId = &v
	return s
}

func (s *UpdateCenInterRegionTrafficQosPolicyAttributeRequest) SetTrafficQosPolicyName(v string) *UpdateCenInterRegionTrafficQosPolicyAttributeRequest {
	s.TrafficQosPolicyName = &v
	return s
}

type UpdateCenInterRegionTrafficQosPolicyAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCenInterRegionTrafficQosPolicyAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCenInterRegionTrafficQosPolicyAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCenInterRegionTrafficQosPolicyAttributeResponseBody) SetRequestId(v string) *UpdateCenInterRegionTrafficQosPolicyAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateCenInterRegionTrafficQosPolicyAttributeResponse struct {
	Headers map[string]*string                                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateCenInterRegionTrafficQosPolicyAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCenInterRegionTrafficQosPolicyAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCenInterRegionTrafficQosPolicyAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateCenInterRegionTrafficQosPolicyAttributeResponse) SetHeaders(v map[string]*string) *UpdateCenInterRegionTrafficQosPolicyAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateCenInterRegionTrafficQosPolicyAttributeResponse) SetBody(v *UpdateCenInterRegionTrafficQosPolicyAttributeResponseBody) *UpdateCenInterRegionTrafficQosPolicyAttributeResponse {
	s.Body = v
	return s
}

type UpdateCenInterRegionTrafficQosQueueAttributeRequest struct {
	ClientToken            *string  `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                 *bool    `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	Dscps                  []*int32 `json:"Dscps,omitempty" xml:"Dscps,omitempty" type:"Repeated"`
	OwnerAccount           *string  `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                *int64   `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	QosQueueDescription    *string  `json:"QosQueueDescription,omitempty" xml:"QosQueueDescription,omitempty"`
	QosQueueId             *string  `json:"QosQueueId,omitempty" xml:"QosQueueId,omitempty"`
	QosQueueName           *string  `json:"QosQueueName,omitempty" xml:"QosQueueName,omitempty"`
	RemainBandwidthPercent *string  `json:"RemainBandwidthPercent,omitempty" xml:"RemainBandwidthPercent,omitempty"`
	ResourceOwnerAccount   *string  `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId        *int64   `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpdateCenInterRegionTrafficQosQueueAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCenInterRegionTrafficQosQueueAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeRequest) SetClientToken(v string) *UpdateCenInterRegionTrafficQosQueueAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeRequest) SetDryRun(v bool) *UpdateCenInterRegionTrafficQosQueueAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeRequest) SetDscps(v []*int32) *UpdateCenInterRegionTrafficQosQueueAttributeRequest {
	s.Dscps = v
	return s
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeRequest) SetOwnerAccount(v string) *UpdateCenInterRegionTrafficQosQueueAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeRequest) SetOwnerId(v int64) *UpdateCenInterRegionTrafficQosQueueAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeRequest) SetQosQueueDescription(v string) *UpdateCenInterRegionTrafficQosQueueAttributeRequest {
	s.QosQueueDescription = &v
	return s
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeRequest) SetQosQueueId(v string) *UpdateCenInterRegionTrafficQosQueueAttributeRequest {
	s.QosQueueId = &v
	return s
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeRequest) SetQosQueueName(v string) *UpdateCenInterRegionTrafficQosQueueAttributeRequest {
	s.QosQueueName = &v
	return s
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeRequest) SetRemainBandwidthPercent(v string) *UpdateCenInterRegionTrafficQosQueueAttributeRequest {
	s.RemainBandwidthPercent = &v
	return s
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeRequest) SetResourceOwnerAccount(v string) *UpdateCenInterRegionTrafficQosQueueAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeRequest) SetResourceOwnerId(v int64) *UpdateCenInterRegionTrafficQosQueueAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

type UpdateCenInterRegionTrafficQosQueueAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCenInterRegionTrafficQosQueueAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCenInterRegionTrafficQosQueueAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeResponseBody) SetRequestId(v string) *UpdateCenInterRegionTrafficQosQueueAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateCenInterRegionTrafficQosQueueAttributeResponse struct {
	Headers map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateCenInterRegionTrafficQosQueueAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCenInterRegionTrafficQosQueueAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCenInterRegionTrafficQosQueueAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeResponse) SetHeaders(v map[string]*string) *UpdateCenInterRegionTrafficQosQueueAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeResponse) SetBody(v *UpdateCenInterRegionTrafficQosQueueAttributeResponseBody) *UpdateCenInterRegionTrafficQosQueueAttributeResponse {
	s.Body = v
	return s
}

type UpdateTrafficMarkingPolicyAttributeRequest struct {
	ClientToken                     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                          *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount                    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount            *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId                 *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TrafficMarkingPolicyDescription *string `json:"TrafficMarkingPolicyDescription,omitempty" xml:"TrafficMarkingPolicyDescription,omitempty"`
	TrafficMarkingPolicyId          *string `json:"TrafficMarkingPolicyId,omitempty" xml:"TrafficMarkingPolicyId,omitempty"`
	TrafficMarkingPolicyName        *string `json:"TrafficMarkingPolicyName,omitempty" xml:"TrafficMarkingPolicyName,omitempty"`
}

func (s UpdateTrafficMarkingPolicyAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTrafficMarkingPolicyAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateTrafficMarkingPolicyAttributeRequest) SetClientToken(v string) *UpdateTrafficMarkingPolicyAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateTrafficMarkingPolicyAttributeRequest) SetDryRun(v bool) *UpdateTrafficMarkingPolicyAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateTrafficMarkingPolicyAttributeRequest) SetOwnerAccount(v string) *UpdateTrafficMarkingPolicyAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateTrafficMarkingPolicyAttributeRequest) SetOwnerId(v int64) *UpdateTrafficMarkingPolicyAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateTrafficMarkingPolicyAttributeRequest) SetResourceOwnerAccount(v string) *UpdateTrafficMarkingPolicyAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateTrafficMarkingPolicyAttributeRequest) SetResourceOwnerId(v int64) *UpdateTrafficMarkingPolicyAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateTrafficMarkingPolicyAttributeRequest) SetTrafficMarkingPolicyDescription(v string) *UpdateTrafficMarkingPolicyAttributeRequest {
	s.TrafficMarkingPolicyDescription = &v
	return s
}

func (s *UpdateTrafficMarkingPolicyAttributeRequest) SetTrafficMarkingPolicyId(v string) *UpdateTrafficMarkingPolicyAttributeRequest {
	s.TrafficMarkingPolicyId = &v
	return s
}

func (s *UpdateTrafficMarkingPolicyAttributeRequest) SetTrafficMarkingPolicyName(v string) *UpdateTrafficMarkingPolicyAttributeRequest {
	s.TrafficMarkingPolicyName = &v
	return s
}

type UpdateTrafficMarkingPolicyAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTrafficMarkingPolicyAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTrafficMarkingPolicyAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTrafficMarkingPolicyAttributeResponseBody) SetRequestId(v string) *UpdateTrafficMarkingPolicyAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateTrafficMarkingPolicyAttributeResponse struct {
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateTrafficMarkingPolicyAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTrafficMarkingPolicyAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTrafficMarkingPolicyAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateTrafficMarkingPolicyAttributeResponse) SetHeaders(v map[string]*string) *UpdateTrafficMarkingPolicyAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateTrafficMarkingPolicyAttributeResponse) SetBody(v *UpdateTrafficMarkingPolicyAttributeResponseBody) *UpdateTrafficMarkingPolicyAttributeResponse {
	s.Body = v
	return s
}

type UpdateTransitRouterRequest struct {
	ClientToken              *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                   *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount             *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount     *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId          *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TransitRouterDescription *string `json:"TransitRouterDescription,omitempty" xml:"TransitRouterDescription,omitempty"`
	TransitRouterId          *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	TransitRouterName        *string `json:"TransitRouterName,omitempty" xml:"TransitRouterName,omitempty"`
}

func (s UpdateTransitRouterRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTransitRouterRequest) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterRequest) SetClientToken(v string) *UpdateTransitRouterRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateTransitRouterRequest) SetDryRun(v bool) *UpdateTransitRouterRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateTransitRouterRequest) SetOwnerAccount(v string) *UpdateTransitRouterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateTransitRouterRequest) SetOwnerId(v int64) *UpdateTransitRouterRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateTransitRouterRequest) SetRegionId(v string) *UpdateTransitRouterRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateTransitRouterRequest) SetResourceOwnerAccount(v string) *UpdateTransitRouterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateTransitRouterRequest) SetResourceOwnerId(v int64) *UpdateTransitRouterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateTransitRouterRequest) SetTransitRouterDescription(v string) *UpdateTransitRouterRequest {
	s.TransitRouterDescription = &v
	return s
}

func (s *UpdateTransitRouterRequest) SetTransitRouterId(v string) *UpdateTransitRouterRequest {
	s.TransitRouterId = &v
	return s
}

func (s *UpdateTransitRouterRequest) SetTransitRouterName(v string) *UpdateTransitRouterRequest {
	s.TransitRouterName = &v
	return s
}

type UpdateTransitRouterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTransitRouterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTransitRouterResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterResponseBody) SetRequestId(v string) *UpdateTransitRouterResponseBody {
	s.RequestId = &v
	return s
}

type UpdateTransitRouterResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateTransitRouterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTransitRouterResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTransitRouterResponse) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterResponse) SetHeaders(v map[string]*string) *UpdateTransitRouterResponse {
	s.Headers = v
	return s
}

func (s *UpdateTransitRouterResponse) SetBody(v *UpdateTransitRouterResponseBody) *UpdateTransitRouterResponse {
	s.Body = v
	return s
}

type UpdateTransitRouterPeerAttachmentAttributeRequest struct {
	AutoPublishRouteEnabled            *bool   `json:"AutoPublishRouteEnabled,omitempty" xml:"AutoPublishRouteEnabled,omitempty"`
	Bandwidth                          *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	BandwidthType                      *string `json:"BandwidthType,omitempty" xml:"BandwidthType,omitempty"`
	CenBandwidthPackageId              *string `json:"CenBandwidthPackageId,omitempty" xml:"CenBandwidthPackageId,omitempty"`
	ClientToken                        *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                             *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount                       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount               *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId                    *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TransitRouterAttachmentDescription *string `json:"TransitRouterAttachmentDescription,omitempty" xml:"TransitRouterAttachmentDescription,omitempty"`
	TransitRouterAttachmentId          *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	TransitRouterAttachmentName        *string `json:"TransitRouterAttachmentName,omitempty" xml:"TransitRouterAttachmentName,omitempty"`
}

func (s UpdateTransitRouterPeerAttachmentAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTransitRouterPeerAttachmentAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterPeerAttachmentAttributeRequest) SetAutoPublishRouteEnabled(v bool) *UpdateTransitRouterPeerAttachmentAttributeRequest {
	s.AutoPublishRouteEnabled = &v
	return s
}

func (s *UpdateTransitRouterPeerAttachmentAttributeRequest) SetBandwidth(v int32) *UpdateTransitRouterPeerAttachmentAttributeRequest {
	s.Bandwidth = &v
	return s
}

func (s *UpdateTransitRouterPeerAttachmentAttributeRequest) SetBandwidthType(v string) *UpdateTransitRouterPeerAttachmentAttributeRequest {
	s.BandwidthType = &v
	return s
}

func (s *UpdateTransitRouterPeerAttachmentAttributeRequest) SetCenBandwidthPackageId(v string) *UpdateTransitRouterPeerAttachmentAttributeRequest {
	s.CenBandwidthPackageId = &v
	return s
}

func (s *UpdateTransitRouterPeerAttachmentAttributeRequest) SetClientToken(v string) *UpdateTransitRouterPeerAttachmentAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateTransitRouterPeerAttachmentAttributeRequest) SetDryRun(v bool) *UpdateTransitRouterPeerAttachmentAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateTransitRouterPeerAttachmentAttributeRequest) SetOwnerAccount(v string) *UpdateTransitRouterPeerAttachmentAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateTransitRouterPeerAttachmentAttributeRequest) SetOwnerId(v int64) *UpdateTransitRouterPeerAttachmentAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateTransitRouterPeerAttachmentAttributeRequest) SetResourceOwnerAccount(v string) *UpdateTransitRouterPeerAttachmentAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateTransitRouterPeerAttachmentAttributeRequest) SetResourceOwnerId(v int64) *UpdateTransitRouterPeerAttachmentAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateTransitRouterPeerAttachmentAttributeRequest) SetTransitRouterAttachmentDescription(v string) *UpdateTransitRouterPeerAttachmentAttributeRequest {
	s.TransitRouterAttachmentDescription = &v
	return s
}

func (s *UpdateTransitRouterPeerAttachmentAttributeRequest) SetTransitRouterAttachmentId(v string) *UpdateTransitRouterPeerAttachmentAttributeRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *UpdateTransitRouterPeerAttachmentAttributeRequest) SetTransitRouterAttachmentName(v string) *UpdateTransitRouterPeerAttachmentAttributeRequest {
	s.TransitRouterAttachmentName = &v
	return s
}

type UpdateTransitRouterPeerAttachmentAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTransitRouterPeerAttachmentAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTransitRouterPeerAttachmentAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterPeerAttachmentAttributeResponseBody) SetRequestId(v string) *UpdateTransitRouterPeerAttachmentAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateTransitRouterPeerAttachmentAttributeResponse struct {
	Headers map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateTransitRouterPeerAttachmentAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTransitRouterPeerAttachmentAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTransitRouterPeerAttachmentAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterPeerAttachmentAttributeResponse) SetHeaders(v map[string]*string) *UpdateTransitRouterPeerAttachmentAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateTransitRouterPeerAttachmentAttributeResponse) SetBody(v *UpdateTransitRouterPeerAttachmentAttributeResponseBody) *UpdateTransitRouterPeerAttachmentAttributeResponse {
	s.Body = v
	return s
}

type UpdateTransitRouterRouteEntryRequest struct {
	ClientToken                        *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                             *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount                       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount               *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId                    *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TransitRouterRouteEntryDescription *string `json:"TransitRouterRouteEntryDescription,omitempty" xml:"TransitRouterRouteEntryDescription,omitempty"`
	TransitRouterRouteEntryId          *string `json:"TransitRouterRouteEntryId,omitempty" xml:"TransitRouterRouteEntryId,omitempty"`
	TransitRouterRouteEntryName        *string `json:"TransitRouterRouteEntryName,omitempty" xml:"TransitRouterRouteEntryName,omitempty"`
}

func (s UpdateTransitRouterRouteEntryRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTransitRouterRouteEntryRequest) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterRouteEntryRequest) SetClientToken(v string) *UpdateTransitRouterRouteEntryRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateTransitRouterRouteEntryRequest) SetDryRun(v bool) *UpdateTransitRouterRouteEntryRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateTransitRouterRouteEntryRequest) SetOwnerAccount(v string) *UpdateTransitRouterRouteEntryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateTransitRouterRouteEntryRequest) SetOwnerId(v int64) *UpdateTransitRouterRouteEntryRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateTransitRouterRouteEntryRequest) SetResourceOwnerAccount(v string) *UpdateTransitRouterRouteEntryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateTransitRouterRouteEntryRequest) SetResourceOwnerId(v int64) *UpdateTransitRouterRouteEntryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateTransitRouterRouteEntryRequest) SetTransitRouterRouteEntryDescription(v string) *UpdateTransitRouterRouteEntryRequest {
	s.TransitRouterRouteEntryDescription = &v
	return s
}

func (s *UpdateTransitRouterRouteEntryRequest) SetTransitRouterRouteEntryId(v string) *UpdateTransitRouterRouteEntryRequest {
	s.TransitRouterRouteEntryId = &v
	return s
}

func (s *UpdateTransitRouterRouteEntryRequest) SetTransitRouterRouteEntryName(v string) *UpdateTransitRouterRouteEntryRequest {
	s.TransitRouterRouteEntryName = &v
	return s
}

type UpdateTransitRouterRouteEntryResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTransitRouterRouteEntryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTransitRouterRouteEntryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterRouteEntryResponseBody) SetRequestId(v string) *UpdateTransitRouterRouteEntryResponseBody {
	s.RequestId = &v
	return s
}

type UpdateTransitRouterRouteEntryResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateTransitRouterRouteEntryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTransitRouterRouteEntryResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTransitRouterRouteEntryResponse) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterRouteEntryResponse) SetHeaders(v map[string]*string) *UpdateTransitRouterRouteEntryResponse {
	s.Headers = v
	return s
}

func (s *UpdateTransitRouterRouteEntryResponse) SetBody(v *UpdateTransitRouterRouteEntryResponseBody) *UpdateTransitRouterRouteEntryResponse {
	s.Body = v
	return s
}

type UpdateTransitRouterRouteTableRequest struct {
	ClientToken                        *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                             *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount                       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount               *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId                    *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TransitRouterRouteTableDescription *string `json:"TransitRouterRouteTableDescription,omitempty" xml:"TransitRouterRouteTableDescription,omitempty"`
	TransitRouterRouteTableId          *string `json:"TransitRouterRouteTableId,omitempty" xml:"TransitRouterRouteTableId,omitempty"`
	TransitRouterRouteTableName        *string `json:"TransitRouterRouteTableName,omitempty" xml:"TransitRouterRouteTableName,omitempty"`
}

func (s UpdateTransitRouterRouteTableRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTransitRouterRouteTableRequest) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterRouteTableRequest) SetClientToken(v string) *UpdateTransitRouterRouteTableRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateTransitRouterRouteTableRequest) SetDryRun(v bool) *UpdateTransitRouterRouteTableRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateTransitRouterRouteTableRequest) SetOwnerAccount(v string) *UpdateTransitRouterRouteTableRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateTransitRouterRouteTableRequest) SetOwnerId(v int64) *UpdateTransitRouterRouteTableRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateTransitRouterRouteTableRequest) SetResourceOwnerAccount(v string) *UpdateTransitRouterRouteTableRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateTransitRouterRouteTableRequest) SetResourceOwnerId(v int64) *UpdateTransitRouterRouteTableRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateTransitRouterRouteTableRequest) SetTransitRouterRouteTableDescription(v string) *UpdateTransitRouterRouteTableRequest {
	s.TransitRouterRouteTableDescription = &v
	return s
}

func (s *UpdateTransitRouterRouteTableRequest) SetTransitRouterRouteTableId(v string) *UpdateTransitRouterRouteTableRequest {
	s.TransitRouterRouteTableId = &v
	return s
}

func (s *UpdateTransitRouterRouteTableRequest) SetTransitRouterRouteTableName(v string) *UpdateTransitRouterRouteTableRequest {
	s.TransitRouterRouteTableName = &v
	return s
}

type UpdateTransitRouterRouteTableResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTransitRouterRouteTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTransitRouterRouteTableResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterRouteTableResponseBody) SetRequestId(v string) *UpdateTransitRouterRouteTableResponseBody {
	s.RequestId = &v
	return s
}

type UpdateTransitRouterRouteTableResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateTransitRouterRouteTableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTransitRouterRouteTableResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTransitRouterRouteTableResponse) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterRouteTableResponse) SetHeaders(v map[string]*string) *UpdateTransitRouterRouteTableResponse {
	s.Headers = v
	return s
}

func (s *UpdateTransitRouterRouteTableResponse) SetBody(v *UpdateTransitRouterRouteTableResponseBody) *UpdateTransitRouterRouteTableResponse {
	s.Body = v
	return s
}

type UpdateTransitRouterVbrAttachmentAttributeRequest struct {
	ClientToken                        *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                             *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount                       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount               *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId                    *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TransitRouterAttachmentDescription *string `json:"TransitRouterAttachmentDescription,omitempty" xml:"TransitRouterAttachmentDescription,omitempty"`
	TransitRouterAttachmentId          *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	TransitRouterAttachmentName        *string `json:"TransitRouterAttachmentName,omitempty" xml:"TransitRouterAttachmentName,omitempty"`
}

func (s UpdateTransitRouterVbrAttachmentAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTransitRouterVbrAttachmentAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterVbrAttachmentAttributeRequest) SetClientToken(v string) *UpdateTransitRouterVbrAttachmentAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateTransitRouterVbrAttachmentAttributeRequest) SetDryRun(v bool) *UpdateTransitRouterVbrAttachmentAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateTransitRouterVbrAttachmentAttributeRequest) SetOwnerAccount(v string) *UpdateTransitRouterVbrAttachmentAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateTransitRouterVbrAttachmentAttributeRequest) SetOwnerId(v int64) *UpdateTransitRouterVbrAttachmentAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateTransitRouterVbrAttachmentAttributeRequest) SetResourceOwnerAccount(v string) *UpdateTransitRouterVbrAttachmentAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateTransitRouterVbrAttachmentAttributeRequest) SetResourceOwnerId(v int64) *UpdateTransitRouterVbrAttachmentAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateTransitRouterVbrAttachmentAttributeRequest) SetTransitRouterAttachmentDescription(v string) *UpdateTransitRouterVbrAttachmentAttributeRequest {
	s.TransitRouterAttachmentDescription = &v
	return s
}

func (s *UpdateTransitRouterVbrAttachmentAttributeRequest) SetTransitRouterAttachmentId(v string) *UpdateTransitRouterVbrAttachmentAttributeRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *UpdateTransitRouterVbrAttachmentAttributeRequest) SetTransitRouterAttachmentName(v string) *UpdateTransitRouterVbrAttachmentAttributeRequest {
	s.TransitRouterAttachmentName = &v
	return s
}

type UpdateTransitRouterVbrAttachmentAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTransitRouterVbrAttachmentAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTransitRouterVbrAttachmentAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterVbrAttachmentAttributeResponseBody) SetRequestId(v string) *UpdateTransitRouterVbrAttachmentAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateTransitRouterVbrAttachmentAttributeResponse struct {
	Headers map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateTransitRouterVbrAttachmentAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTransitRouterVbrAttachmentAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTransitRouterVbrAttachmentAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterVbrAttachmentAttributeResponse) SetHeaders(v map[string]*string) *UpdateTransitRouterVbrAttachmentAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateTransitRouterVbrAttachmentAttributeResponse) SetBody(v *UpdateTransitRouterVbrAttachmentAttributeResponseBody) *UpdateTransitRouterVbrAttachmentAttributeResponse {
	s.Body = v
	return s
}

type UpdateTransitRouterVpcAttachmentAttributeRequest struct {
	ClientToken                        *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                             *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount                       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount               *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId                    *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TransitRouterAttachmentDescription *string `json:"TransitRouterAttachmentDescription,omitempty" xml:"TransitRouterAttachmentDescription,omitempty"`
	TransitRouterAttachmentId          *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	TransitRouterAttachmentName        *string `json:"TransitRouterAttachmentName,omitempty" xml:"TransitRouterAttachmentName,omitempty"`
}

func (s UpdateTransitRouterVpcAttachmentAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTransitRouterVpcAttachmentAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterVpcAttachmentAttributeRequest) SetClientToken(v string) *UpdateTransitRouterVpcAttachmentAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentAttributeRequest) SetDryRun(v bool) *UpdateTransitRouterVpcAttachmentAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentAttributeRequest) SetOwnerAccount(v string) *UpdateTransitRouterVpcAttachmentAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentAttributeRequest) SetOwnerId(v int64) *UpdateTransitRouterVpcAttachmentAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentAttributeRequest) SetResourceOwnerAccount(v string) *UpdateTransitRouterVpcAttachmentAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentAttributeRequest) SetResourceOwnerId(v int64) *UpdateTransitRouterVpcAttachmentAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentAttributeRequest) SetTransitRouterAttachmentDescription(v string) *UpdateTransitRouterVpcAttachmentAttributeRequest {
	s.TransitRouterAttachmentDescription = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentAttributeRequest) SetTransitRouterAttachmentId(v string) *UpdateTransitRouterVpcAttachmentAttributeRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentAttributeRequest) SetTransitRouterAttachmentName(v string) *UpdateTransitRouterVpcAttachmentAttributeRequest {
	s.TransitRouterAttachmentName = &v
	return s
}

type UpdateTransitRouterVpcAttachmentAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTransitRouterVpcAttachmentAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTransitRouterVpcAttachmentAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterVpcAttachmentAttributeResponseBody) SetRequestId(v string) *UpdateTransitRouterVpcAttachmentAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateTransitRouterVpcAttachmentAttributeResponse struct {
	Headers map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateTransitRouterVpcAttachmentAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTransitRouterVpcAttachmentAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTransitRouterVpcAttachmentAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterVpcAttachmentAttributeResponse) SetHeaders(v map[string]*string) *UpdateTransitRouterVpcAttachmentAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentAttributeResponse) SetBody(v *UpdateTransitRouterVpcAttachmentAttributeResponseBody) *UpdateTransitRouterVpcAttachmentAttributeResponse {
	s.Body = v
	return s
}

type WithdrawPublishedRouteEntriesRequest struct {
	CenId                     *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ChildInstanceId           *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	ChildInstanceRegionId     *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	ChildInstanceRouteTableId *string `json:"ChildInstanceRouteTableId,omitempty" xml:"ChildInstanceRouteTableId,omitempty"`
	ChildInstanceType         *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	DestinationCidrBlock      *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	ResourceOwnerAccount      *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId           *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s WithdrawPublishedRouteEntriesRequest) String() string {
	return tea.Prettify(s)
}

func (s WithdrawPublishedRouteEntriesRequest) GoString() string {
	return s.String()
}

func (s *WithdrawPublishedRouteEntriesRequest) SetCenId(v string) *WithdrawPublishedRouteEntriesRequest {
	s.CenId = &v
	return s
}

func (s *WithdrawPublishedRouteEntriesRequest) SetChildInstanceId(v string) *WithdrawPublishedRouteEntriesRequest {
	s.ChildInstanceId = &v
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

func (s *WithdrawPublishedRouteEntriesRequest) SetChildInstanceType(v string) *WithdrawPublishedRouteEntriesRequest {
	s.ChildInstanceType = &v
	return s
}

func (s *WithdrawPublishedRouteEntriesRequest) SetDestinationCidrBlock(v string) *WithdrawPublishedRouteEntriesRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *WithdrawPublishedRouteEntriesRequest) SetResourceOwnerAccount(v string) *WithdrawPublishedRouteEntriesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *WithdrawPublishedRouteEntriesRequest) SetResourceOwnerId(v int64) *WithdrawPublishedRouteEntriesRequest {
	s.ResourceOwnerId = &v
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
	query := map[string]interface{}{}
	query["CenId"] = request.CenId
	query["ClientToken"] = request.ClientToken
	query["FlowLogId"] = request.FlowLogId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ActiveFlowLog"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ActiveFlowLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) AddTraficMatchRuleToTrafficMarkingPolicyWithOptions(request *AddTraficMatchRuleToTrafficMarkingPolicyRequest, runtime *util.RuntimeOptions) (_result *AddTraficMatchRuleToTrafficMarkingPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TrafficMarkingPolicyId"] = request.TrafficMarkingPolicyId
	query["TrafficMatchRules"] = request.TrafficMatchRules
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddTraficMatchRuleToTrafficMarkingPolicy"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddTraficMatchRuleToTrafficMarkingPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddTraficMatchRuleToTrafficMarkingPolicy(request *AddTraficMatchRuleToTrafficMarkingPolicyRequest) (_result *AddTraficMatchRuleToTrafficMarkingPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddTraficMatchRuleToTrafficMarkingPolicyResponse{}
	_body, _err := client.AddTraficMatchRuleToTrafficMarkingPolicyWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	query["CenBandwidthPackageId"] = request.CenBandwidthPackageId
	query["CenId"] = request.CenId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AssociateCenBandwidthPackage"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AssociateCenBandwidthPackageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) AssociateTransitRouterAttachmentWithRouteTableWithOptions(request *AssociateTransitRouterAttachmentWithRouteTableRequest, runtime *util.RuntimeOptions) (_result *AssociateTransitRouterAttachmentWithRouteTableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	query["TransitRouterRouteTableId"] = request.TransitRouterRouteTableId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AssociateTransitRouterAttachmentWithRouteTable"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AssociateTransitRouterAttachmentWithRouteTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AssociateTransitRouterAttachmentWithRouteTable(request *AssociateTransitRouterAttachmentWithRouteTableRequest) (_result *AssociateTransitRouterAttachmentWithRouteTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssociateTransitRouterAttachmentWithRouteTableResponse{}
	_body, _err := client.AssociateTransitRouterAttachmentWithRouteTableWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	query["CenId"] = request.CenId
	query["ChildInstanceId"] = request.ChildInstanceId
	query["ChildInstanceOwnerId"] = request.ChildInstanceOwnerId
	query["ChildInstanceRegionId"] = request.ChildInstanceRegionId
	query["ChildInstanceType"] = request.ChildInstanceType
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachCenChildInstance"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachCenChildInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) CheckTransitRouterServiceWithOptions(request *CheckTransitRouterServiceRequest, runtime *util.RuntimeOptions) (_result *CheckTransitRouterServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckTransitRouterService"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckTransitRouterServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckTransitRouterService(request *CheckTransitRouterServiceRequest) (_result *CheckTransitRouterServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckTransitRouterServiceResponse{}
	_body, _err := client.CheckTransitRouterServiceWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["Description"] = request.Description
	query["Name"] = request.Name
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ProtectionLevel"] = request.ProtectionLevel
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCen"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["AutoPay"] = request.AutoPay
	query["AutoRenew"] = request.AutoRenew
	query["AutoRenewDuration"] = request.AutoRenewDuration
	query["Bandwidth"] = request.Bandwidth
	query["BandwidthPackageChargeType"] = request.BandwidthPackageChargeType
	query["ClientToken"] = request.ClientToken
	query["Description"] = request.Description
	query["GeographicRegionAId"] = request.GeographicRegionAId
	query["GeographicRegionBId"] = request.GeographicRegionBId
	query["Name"] = request.Name
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["Period"] = request.Period
	query["PricingCycle"] = request.PricingCycle
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCenBandwidthPackage"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCenBandwidthPackageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) CreateCenChildInstanceRouteEntryToAttachmentWithOptions(request *CreateCenChildInstanceRouteEntryToAttachmentRequest, runtime *util.RuntimeOptions) (_result *CreateCenChildInstanceRouteEntryToAttachmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["CenId"] = request.CenId
	query["ClientToken"] = request.ClientToken
	query["DestinationCidrBlock"] = request.DestinationCidrBlock
	query["DryRun"] = request.DryRun
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["RouteTableId"] = request.RouteTableId
	query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCenChildInstanceRouteEntryToAttachment"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCenChildInstanceRouteEntryToAttachmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCenChildInstanceRouteEntryToAttachment(request *CreateCenChildInstanceRouteEntryToAttachmentRequest) (_result *CreateCenChildInstanceRouteEntryToAttachmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCenChildInstanceRouteEntryToAttachmentResponse{}
	_body, _err := client.CreateCenChildInstanceRouteEntryToAttachmentWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	query["CenId"] = request.CenId
	query["ChildInstanceAliUid"] = request.ChildInstanceAliUid
	query["ChildInstanceId"] = request.ChildInstanceId
	query["ChildInstanceRegionId"] = request.ChildInstanceRegionId
	query["ChildInstanceType"] = request.ChildInstanceType
	query["DestinationCidrBlock"] = request.DestinationCidrBlock
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["RouteTableId"] = request.RouteTableId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCenChildInstanceRouteEntryToCen"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCenChildInstanceRouteEntryToCenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) CreateCenInterRegionTrafficQosPolicyWithOptions(request *CreateCenInterRegionTrafficQosPolicyRequest, runtime *util.RuntimeOptions) (_result *CreateCenInterRegionTrafficQosPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TrafficQosPolicyDescription"] = request.TrafficQosPolicyDescription
	query["TrafficQosPolicyName"] = request.TrafficQosPolicyName
	query["TrafficQosQueues"] = request.TrafficQosQueues
	query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	query["TransitRouterId"] = request.TransitRouterId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCenInterRegionTrafficQosPolicy"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCenInterRegionTrafficQosPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCenInterRegionTrafficQosPolicy(request *CreateCenInterRegionTrafficQosPolicyRequest) (_result *CreateCenInterRegionTrafficQosPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCenInterRegionTrafficQosPolicyResponse{}
	_body, _err := client.CreateCenInterRegionTrafficQosPolicyWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	query["AsPathMatchMode"] = request.AsPathMatchMode
	query["CenId"] = request.CenId
	query["CenRegionId"] = request.CenRegionId
	query["CidrMatchMode"] = request.CidrMatchMode
	query["CommunityMatchMode"] = request.CommunityMatchMode
	query["CommunityOperateMode"] = request.CommunityOperateMode
	query["Description"] = request.Description
	query["DestinationChildInstanceTypes"] = request.DestinationChildInstanceTypes
	query["DestinationCidrBlocks"] = request.DestinationCidrBlocks
	query["DestinationInstanceIds"] = request.DestinationInstanceIds
	query["DestinationInstanceIdsReverseMatch"] = request.DestinationInstanceIdsReverseMatch
	query["DestinationRouteTableIds"] = request.DestinationRouteTableIds
	query["MapResult"] = request.MapResult
	query["MatchAsns"] = request.MatchAsns
	query["MatchCommunitySet"] = request.MatchCommunitySet
	query["NextPriority"] = request.NextPriority
	query["OperateCommunitySet"] = request.OperateCommunitySet
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["Preference"] = request.Preference
	query["PrependAsPath"] = request.PrependAsPath
	query["Priority"] = request.Priority
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["RouteTypes"] = request.RouteTypes
	query["SourceChildInstanceTypes"] = request.SourceChildInstanceTypes
	query["SourceInstanceIds"] = request.SourceInstanceIds
	query["SourceInstanceIdsReverseMatch"] = request.SourceInstanceIdsReverseMatch
	query["SourceRegionIds"] = request.SourceRegionIds
	query["SourceRouteTableIds"] = request.SourceRouteTableIds
	query["TransmitDirection"] = request.TransmitDirection
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCenRouteMap"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCenRouteMapResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["CenId"] = request.CenId
	query["ClientToken"] = request.ClientToken
	query["Description"] = request.Description
	query["FlowLogName"] = request.FlowLogName
	query["LogStoreName"] = request.LogStoreName
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ProjectName"] = request.ProjectName
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFlowlog"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFlowlogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) CreateTrafficMarkingPolicyWithOptions(request *CreateTrafficMarkingPolicyRequest, runtime *util.RuntimeOptions) (_result *CreateTrafficMarkingPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["MarkingDscp"] = request.MarkingDscp
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["Priority"] = request.Priority
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TrafficMarkingPolicyDescription"] = request.TrafficMarkingPolicyDescription
	query["TrafficMarkingPolicyName"] = request.TrafficMarkingPolicyName
	query["TrafficMatchRules"] = request.TrafficMatchRules
	query["TransitRouterId"] = request.TransitRouterId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTrafficMarkingPolicy"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTrafficMarkingPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTrafficMarkingPolicy(request *CreateTrafficMarkingPolicyRequest) (_result *CreateTrafficMarkingPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTrafficMarkingPolicyResponse{}
	_body, _err := client.CreateTrafficMarkingPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTransitRouterWithOptions(request *CreateTransitRouterRequest, runtime *util.RuntimeOptions) (_result *CreateTransitRouterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["CenId"] = request.CenId
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TransitRouterDescription"] = request.TransitRouterDescription
	query["TransitRouterName"] = request.TransitRouterName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTransitRouter"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTransitRouterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTransitRouter(request *CreateTransitRouterRequest) (_result *CreateTransitRouterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTransitRouterResponse{}
	_body, _err := client.CreateTransitRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTransitRouterPeerAttachmentWithOptions(request *CreateTransitRouterPeerAttachmentRequest, runtime *util.RuntimeOptions) (_result *CreateTransitRouterPeerAttachmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AutoPublishRouteEnabled"] = request.AutoPublishRouteEnabled
	query["Bandwidth"] = request.Bandwidth
	query["BandwidthType"] = request.BandwidthType
	query["CenBandwidthPackageId"] = request.CenBandwidthPackageId
	query["CenId"] = request.CenId
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["PeerTransitRouterId"] = request.PeerTransitRouterId
	query["PeerTransitRouterRegionId"] = request.PeerTransitRouterRegionId
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TransitRouterAttachmentDescription"] = request.TransitRouterAttachmentDescription
	query["TransitRouterAttachmentName"] = request.TransitRouterAttachmentName
	query["TransitRouterId"] = request.TransitRouterId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTransitRouterPeerAttachment"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTransitRouterPeerAttachmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTransitRouterPeerAttachment(request *CreateTransitRouterPeerAttachmentRequest) (_result *CreateTransitRouterPeerAttachmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTransitRouterPeerAttachmentResponse{}
	_body, _err := client.CreateTransitRouterPeerAttachmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTransitRouterRouteEntryWithOptions(request *CreateTransitRouterRouteEntryRequest, runtime *util.RuntimeOptions) (_result *CreateTransitRouterRouteEntryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TransitRouterRouteEntryDescription"] = request.TransitRouterRouteEntryDescription
	query["TransitRouterRouteEntryDestinationCidrBlock"] = request.TransitRouterRouteEntryDestinationCidrBlock
	query["TransitRouterRouteEntryName"] = request.TransitRouterRouteEntryName
	query["TransitRouterRouteEntryNextHopId"] = request.TransitRouterRouteEntryNextHopId
	query["TransitRouterRouteEntryNextHopType"] = request.TransitRouterRouteEntryNextHopType
	query["TransitRouterRouteTableId"] = request.TransitRouterRouteTableId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTransitRouterRouteEntry"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTransitRouterRouteEntryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTransitRouterRouteEntry(request *CreateTransitRouterRouteEntryRequest) (_result *CreateTransitRouterRouteEntryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTransitRouterRouteEntryResponse{}
	_body, _err := client.CreateTransitRouterRouteEntryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTransitRouterRouteTableWithOptions(request *CreateTransitRouterRouteTableRequest, runtime *util.RuntimeOptions) (_result *CreateTransitRouterRouteTableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TransitRouterId"] = request.TransitRouterId
	query["TransitRouterRouteTableDescription"] = request.TransitRouterRouteTableDescription
	query["TransitRouterRouteTableName"] = request.TransitRouterRouteTableName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTransitRouterRouteTable"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTransitRouterRouteTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTransitRouterRouteTable(request *CreateTransitRouterRouteTableRequest) (_result *CreateTransitRouterRouteTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTransitRouterRouteTableResponse{}
	_body, _err := client.CreateTransitRouterRouteTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTransitRouterVbrAttachmentWithOptions(request *CreateTransitRouterVbrAttachmentRequest, runtime *util.RuntimeOptions) (_result *CreateTransitRouterVbrAttachmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AutoPublishRouteEnabled"] = request.AutoPublishRouteEnabled
	query["CenId"] = request.CenId
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TransitRouterAttachmentDescription"] = request.TransitRouterAttachmentDescription
	query["TransitRouterAttachmentName"] = request.TransitRouterAttachmentName
	query["TransitRouterId"] = request.TransitRouterId
	query["VbrId"] = request.VbrId
	query["VbrOwnerId"] = request.VbrOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTransitRouterVbrAttachment"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTransitRouterVbrAttachmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTransitRouterVbrAttachment(request *CreateTransitRouterVbrAttachmentRequest) (_result *CreateTransitRouterVbrAttachmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTransitRouterVbrAttachmentResponse{}
	_body, _err := client.CreateTransitRouterVbrAttachmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTransitRouterVpcAttachmentWithOptions(request *CreateTransitRouterVpcAttachmentRequest, runtime *util.RuntimeOptions) (_result *CreateTransitRouterVpcAttachmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["CenId"] = request.CenId
	query["ChargeType"] = request.ChargeType
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TransitRouterAttachmentDescription"] = request.TransitRouterAttachmentDescription
	query["TransitRouterAttachmentName"] = request.TransitRouterAttachmentName
	query["TransitRouterId"] = request.TransitRouterId
	query["VpcId"] = request.VpcId
	query["VpcOwnerId"] = request.VpcOwnerId
	query["ZoneMappings"] = request.ZoneMappings
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTransitRouterVpcAttachment"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTransitRouterVpcAttachmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTransitRouterVpcAttachment(request *CreateTransitRouterVpcAttachmentRequest) (_result *CreateTransitRouterVpcAttachmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTransitRouterVpcAttachmentResponse{}
	_body, _err := client.CreateTransitRouterVpcAttachmentWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	query["CenId"] = request.CenId
	query["ClientToken"] = request.ClientToken
	query["FlowLogId"] = request.FlowLogId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeactiveFlowLog"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeactiveFlowLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["CenId"] = request.CenId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCen"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["CenBandwidthPackageId"] = request.CenBandwidthPackageId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCenBandwidthPackage"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCenBandwidthPackageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DeleteCenChildInstanceRouteEntryToAttachmentWithOptions(request *DeleteCenChildInstanceRouteEntryToAttachmentRequest, runtime *util.RuntimeOptions) (_result *DeleteCenChildInstanceRouteEntryToAttachmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["CenId"] = request.CenId
	query["ClientToken"] = request.ClientToken
	query["DestinationCidrBlock"] = request.DestinationCidrBlock
	query["DryRun"] = request.DryRun
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["RouteTableId"] = request.RouteTableId
	query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCenChildInstanceRouteEntryToAttachment"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCenChildInstanceRouteEntryToAttachmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCenChildInstanceRouteEntryToAttachment(request *DeleteCenChildInstanceRouteEntryToAttachmentRequest) (_result *DeleteCenChildInstanceRouteEntryToAttachmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCenChildInstanceRouteEntryToAttachmentResponse{}
	_body, _err := client.DeleteCenChildInstanceRouteEntryToAttachmentWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	query["CenId"] = request.CenId
	query["ChildInstanceAliUid"] = request.ChildInstanceAliUid
	query["ChildInstanceId"] = request.ChildInstanceId
	query["ChildInstanceRegionId"] = request.ChildInstanceRegionId
	query["ChildInstanceType"] = request.ChildInstanceType
	query["DestinationCidrBlock"] = request.DestinationCidrBlock
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["RouteTableId"] = request.RouteTableId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCenChildInstanceRouteEntryToCen"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCenChildInstanceRouteEntryToCenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DeleteCenInterRegionTrafficQosPolicyWithOptions(request *DeleteCenInterRegionTrafficQosPolicyRequest, runtime *util.RuntimeOptions) (_result *DeleteCenInterRegionTrafficQosPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TrafficQosPolicyId"] = request.TrafficQosPolicyId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCenInterRegionTrafficQosPolicy"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCenInterRegionTrafficQosPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCenInterRegionTrafficQosPolicy(request *DeleteCenInterRegionTrafficQosPolicyRequest) (_result *DeleteCenInterRegionTrafficQosPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCenInterRegionTrafficQosPolicyResponse{}
	_body, _err := client.DeleteCenInterRegionTrafficQosPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCenInterRegionTrafficQosQueueWithOptions(request *DeleteCenInterRegionTrafficQosQueueRequest, runtime *util.RuntimeOptions) (_result *DeleteCenInterRegionTrafficQosQueueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["QosQueueId"] = request.QosQueueId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCenInterRegionTrafficQosQueue"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCenInterRegionTrafficQosQueueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCenInterRegionTrafficQosQueue(request *DeleteCenInterRegionTrafficQosQueueRequest) (_result *DeleteCenInterRegionTrafficQosQueueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCenInterRegionTrafficQosQueueResponse{}
	_body, _err := client.DeleteCenInterRegionTrafficQosQueueWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	query["CenId"] = request.CenId
	query["CenRegionId"] = request.CenRegionId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["RouteMapId"] = request.RouteMapId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCenRouteMap"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCenRouteMapResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["CenId"] = request.CenId
	query["ClientToken"] = request.ClientToken
	query["FlowLogId"] = request.FlowLogId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFlowlog"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFlowlogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["AccessRegionId"] = request.AccessRegionId
	query["CenId"] = request.CenId
	query["Host"] = request.Host
	query["HostRegionId"] = request.HostRegionId
	query["HostVpcId"] = request.HostVpcId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRouteServiceInCen"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRouteServiceInCenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DeleteTrafficMarkingPolicyWithOptions(request *DeleteTrafficMarkingPolicyRequest, runtime *util.RuntimeOptions) (_result *DeleteTrafficMarkingPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TrafficMarkingPolicyId"] = request.TrafficMarkingPolicyId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTrafficMarkingPolicy"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTrafficMarkingPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTrafficMarkingPolicy(request *DeleteTrafficMarkingPolicyRequest) (_result *DeleteTrafficMarkingPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTrafficMarkingPolicyResponse{}
	_body, _err := client.DeleteTrafficMarkingPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTransitRouterPeerAttachmentWithOptions(request *DeleteTransitRouterPeerAttachmentRequest, runtime *util.RuntimeOptions) (_result *DeleteTransitRouterPeerAttachmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTransitRouterPeerAttachment"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTransitRouterPeerAttachmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTransitRouterPeerAttachment(request *DeleteTransitRouterPeerAttachmentRequest) (_result *DeleteTransitRouterPeerAttachmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTransitRouterPeerAttachmentResponse{}
	_body, _err := client.DeleteTransitRouterPeerAttachmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTransitRouterRouteEntryWithOptions(request *DeleteTransitRouterRouteEntryRequest, runtime *util.RuntimeOptions) (_result *DeleteTransitRouterRouteEntryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TransitRouterRouteEntryDestinationCidrBlock"] = request.TransitRouterRouteEntryDestinationCidrBlock
	query["TransitRouterRouteEntryId"] = request.TransitRouterRouteEntryId
	query["TransitRouterRouteEntryNextHopId"] = request.TransitRouterRouteEntryNextHopId
	query["TransitRouterRouteEntryNextHopType"] = request.TransitRouterRouteEntryNextHopType
	query["TransitRouterRouteTableId"] = request.TransitRouterRouteTableId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTransitRouterRouteEntry"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTransitRouterRouteEntryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTransitRouterRouteEntry(request *DeleteTransitRouterRouteEntryRequest) (_result *DeleteTransitRouterRouteEntryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTransitRouterRouteEntryResponse{}
	_body, _err := client.DeleteTransitRouterRouteEntryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTransitRouterRouteTableWithOptions(request *DeleteTransitRouterRouteTableRequest, runtime *util.RuntimeOptions) (_result *DeleteTransitRouterRouteTableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TransitRouterRouteTableId"] = request.TransitRouterRouteTableId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTransitRouterRouteTable"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTransitRouterRouteTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTransitRouterRouteTable(request *DeleteTransitRouterRouteTableRequest) (_result *DeleteTransitRouterRouteTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTransitRouterRouteTableResponse{}
	_body, _err := client.DeleteTransitRouterRouteTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTransitRouterVbrAttachmentWithOptions(request *DeleteTransitRouterVbrAttachmentRequest, runtime *util.RuntimeOptions) (_result *DeleteTransitRouterVbrAttachmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTransitRouterVbrAttachment"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTransitRouterVbrAttachmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTransitRouterVbrAttachment(request *DeleteTransitRouterVbrAttachmentRequest) (_result *DeleteTransitRouterVbrAttachmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTransitRouterVbrAttachmentResponse{}
	_body, _err := client.DeleteTransitRouterVbrAttachmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTransitRouterVpcAttachmentWithOptions(request *DeleteTransitRouterVpcAttachmentRequest, runtime *util.RuntimeOptions) (_result *DeleteTransitRouterVpcAttachmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTransitRouterVpcAttachment"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTransitRouterVpcAttachmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTransitRouterVpcAttachment(request *DeleteTransitRouterVpcAttachmentRequest) (_result *DeleteTransitRouterVpcAttachmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTransitRouterVpcAttachmentResponse{}
	_body, _err := client.DeleteTransitRouterVpcAttachmentWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	query["CenId"] = request.CenId
	query["ChildInstanceId"] = request.ChildInstanceId
	query["ChildInstanceRegionId"] = request.ChildInstanceRegionId
	query["ChildInstanceType"] = request.ChildInstanceType
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCenAttachedChildInstanceAttribute"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCenAttachedChildInstanceAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["CenId"] = request.CenId
	query["ChildInstanceRegionId"] = request.ChildInstanceRegionId
	query["ChildInstanceType"] = request.ChildInstanceType
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCenAttachedChildInstances"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCenAttachedChildInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["Filter"] = request.Filter
	query["IncludeReservationData"] = request.IncludeReservationData
	query["IsOrKey"] = request.IsOrKey
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCenBandwidthPackages"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCenBandwidthPackagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["CenId"] = request.CenId
	query["ChildInstanceId"] = request.ChildInstanceId
	query["ChildInstanceRegionId"] = request.ChildInstanceRegionId
	query["ChildInstanceType"] = request.ChildInstanceType
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["Status"] = request.Status
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCenChildInstanceRouteEntries"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCenChildInstanceRouteEntriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["CenId"] = request.CenId
	query["GeographicRegionAId"] = request.GeographicRegionAId
	query["GeographicRegionBId"] = request.GeographicRegionBId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCenGeographicSpanRemainingBandwidth"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCenGeographicSpanRemainingBandwidthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["GeographicSpanId"] = request.GeographicSpanId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCenGeographicSpans"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCenGeographicSpansResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["CenId"] = request.CenId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCenInterRegionBandwidthLimits"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCenInterRegionBandwidthLimitsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["AccessRegionId"] = request.AccessRegionId
	query["CenId"] = request.CenId
	query["HostRegionId"] = request.HostRegionId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCenPrivateZoneRoutes"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCenPrivateZoneRoutesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["CenId"] = request.CenId
	query["CenRegionId"] = request.CenRegionId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["Status"] = request.Status
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCenRegionDomainRouteEntries"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCenRegionDomainRouteEntriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["CenId"] = request.CenId
	query["CenRegionId"] = request.CenRegionId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["RouteMapId"] = request.RouteMapId
	query["TransmitDirection"] = request.TransmitDirection
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCenRouteMaps"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCenRouteMapsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeCenVbrHealthCheckWithOptions(request *DescribeCenVbrHealthCheckRequest, runtime *util.RuntimeOptions) (_result *DescribeCenVbrHealthCheckResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["CenId"] = request.CenId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["VbrInstanceId"] = request.VbrInstanceId
	query["VbrInstanceOwnerId"] = request.VbrInstanceOwnerId
	query["VbrInstanceRegionId"] = request.VbrInstanceRegionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCenVbrHealthCheck"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCenVbrHealthCheckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeCensWithOptions(request *DescribeCensRequest, runtime *util.RuntimeOptions) (_result *DescribeCensResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Filter"] = request.Filter
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["Tag"] = request.Tag
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCens"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCensResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeChildInstanceRegionsWithOptions(request *DescribeChildInstanceRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeChildInstanceRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ProductType"] = request.ProductType
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeChildInstanceRegions"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeChildInstanceRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["CenId"] = request.CenId
	query["ClientToken"] = request.ClientToken
	query["Description"] = request.Description
	query["FlowLogId"] = request.FlowLogId
	query["FlowLogName"] = request.FlowLogName
	query["LogStoreName"] = request.LogStoreName
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["ProjectName"] = request.ProjectName
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["Status"] = request.Status
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFlowlogs"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFlowlogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["GeographicRegionId"] = request.GeographicRegionId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGeographicRegionMembership"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGeographicRegionMembershipResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["CenId"] = request.CenId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ProductType"] = request.ProductType
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGrantRulesToCen"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGrantRulesToCenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["CenId"] = request.CenId
	query["ChildInstanceId"] = request.ChildInstanceId
	query["ChildInstanceRegionId"] = request.ChildInstanceRegionId
	query["ChildInstanceRouteTableId"] = request.ChildInstanceRouteTableId
	query["ChildInstanceType"] = request.ChildInstanceType
	query["DestinationCidrBlock"] = request.DestinationCidrBlock
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePublishedRouteEntries"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePublishedRouteEntriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["ChildInstanceId"] = request.ChildInstanceId
	query["ChildInstanceRegionId"] = request.ChildInstanceRegionId
	query["ChildInstanceRouteTableId"] = request.ChildInstanceRouteTableId
	query["ChildInstanceType"] = request.ChildInstanceType
	query["DestinationCidrBlock"] = request.DestinationCidrBlock
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRouteConflict"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRouteConflictResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["AccessRegionId"] = request.AccessRegionId
	query["CenId"] = request.CenId
	query["Host"] = request.Host
	query["HostRegionId"] = request.HostRegionId
	query["HostVpcId"] = request.HostVpcId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRouteServicesInCen"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRouteServicesInCenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["CenId"] = request.CenId
	query["CenOwnerId"] = request.CenOwnerId
	query["ChildInstanceId"] = request.ChildInstanceId
	query["ChildInstanceOwnerId"] = request.ChildInstanceOwnerId
	query["ChildInstanceRegionId"] = request.ChildInstanceRegionId
	query["ChildInstanceType"] = request.ChildInstanceType
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachCenChildInstance"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachCenChildInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["CenId"] = request.CenId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["VbrInstanceId"] = request.VbrInstanceId
	query["VbrInstanceOwnerId"] = request.VbrInstanceOwnerId
	query["VbrInstanceRegionId"] = request.VbrInstanceRegionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableCenVbrHealthCheck"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableCenVbrHealthCheckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DisableTransitRouterRouteTablePropagationWithOptions(request *DisableTransitRouterRouteTablePropagationRequest, runtime *util.RuntimeOptions) (_result *DisableTransitRouterRouteTablePropagationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	query["TransitRouterRouteTableId"] = request.TransitRouterRouteTableId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableTransitRouterRouteTablePropagation"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableTransitRouterRouteTablePropagationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableTransitRouterRouteTablePropagation(request *DisableTransitRouterRouteTablePropagationRequest) (_result *DisableTransitRouterRouteTablePropagationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableTransitRouterRouteTablePropagationResponse{}
	_body, _err := client.DisableTransitRouterRouteTablePropagationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DissociateTransitRouterAttachmentFromRouteTableWithOptions(request *DissociateTransitRouterAttachmentFromRouteTableRequest, runtime *util.RuntimeOptions) (_result *DissociateTransitRouterAttachmentFromRouteTableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	query["TransitRouterRouteTableId"] = request.TransitRouterRouteTableId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DissociateTransitRouterAttachmentFromRouteTable"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DissociateTransitRouterAttachmentFromRouteTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DissociateTransitRouterAttachmentFromRouteTable(request *DissociateTransitRouterAttachmentFromRouteTableRequest) (_result *DissociateTransitRouterAttachmentFromRouteTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DissociateTransitRouterAttachmentFromRouteTableResponse{}
	_body, _err := client.DissociateTransitRouterAttachmentFromRouteTableWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	query["CenId"] = request.CenId
	query["HealthCheckInterval"] = request.HealthCheckInterval
	query["HealthCheckOnly"] = request.HealthCheckOnly
	query["HealthCheckSourceIp"] = request.HealthCheckSourceIp
	query["HealthCheckTargetIp"] = request.HealthCheckTargetIp
	query["HealthyThreshold"] = request.HealthyThreshold
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["VbrInstanceId"] = request.VbrInstanceId
	query["VbrInstanceOwnerId"] = request.VbrInstanceOwnerId
	query["VbrInstanceRegionId"] = request.VbrInstanceRegionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableCenVbrHealthCheck"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableCenVbrHealthCheckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) EnableTransitRouterRouteTablePropagationWithOptions(request *EnableTransitRouterRouteTablePropagationRequest, runtime *util.RuntimeOptions) (_result *EnableTransitRouterRouteTablePropagationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	query["TransitRouterRouteTableId"] = request.TransitRouterRouteTableId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableTransitRouterRouteTablePropagation"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableTransitRouterRouteTablePropagationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableTransitRouterRouteTablePropagation(request *EnableTransitRouterRouteTablePropagationRequest) (_result *EnableTransitRouterRouteTablePropagationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableTransitRouterRouteTablePropagationResponse{}
	_body, _err := client.EnableTransitRouterRouteTablePropagationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GrantInstanceToTransitRouterWithOptions(request *GrantInstanceToTransitRouterRequest, runtime *util.RuntimeOptions) (_result *GrantInstanceToTransitRouterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["CenId"] = request.CenId
	query["CenOwnerId"] = request.CenOwnerId
	query["InstanceId"] = request.InstanceId
	query["InstanceType"] = request.InstanceType
	query["OrderType"] = request.OrderType
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GrantInstanceToTransitRouter"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GrantInstanceToTransitRouterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GrantInstanceToTransitRouter(request *GrantInstanceToTransitRouterRequest) (_result *GrantInstanceToTransitRouterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GrantInstanceToTransitRouterResponse{}
	_body, _err := client.GrantInstanceToTransitRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCenInterRegionTrafficQosPoliciesWithOptions(request *ListCenInterRegionTrafficQosPoliciesRequest, runtime *util.RuntimeOptions) (_result *ListCenInterRegionTrafficQosPoliciesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["MaxResults"] = request.MaxResults
	query["NextToken"] = request.NextToken
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TrafficQosPolicyDescription"] = request.TrafficQosPolicyDescription
	query["TrafficQosPolicyId"] = request.TrafficQosPolicyId
	query["TrafficQosPolicyName"] = request.TrafficQosPolicyName
	query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	query["TransitRouterId"] = request.TransitRouterId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCenInterRegionTrafficQosPolicies"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCenInterRegionTrafficQosPoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCenInterRegionTrafficQosPolicies(request *ListCenInterRegionTrafficQosPoliciesRequest) (_result *ListCenInterRegionTrafficQosPoliciesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCenInterRegionTrafficQosPoliciesResponse{}
	_body, _err := client.ListCenInterRegionTrafficQosPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListGrantVSwitchesToCenWithOptions(request *ListGrantVSwitchesToCenRequest, runtime *util.RuntimeOptions) (_result *ListGrantVSwitchesToCenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["CenId"] = request.CenId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["VpcId"] = request.VpcId
	query["ZoneId"] = request.ZoneId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGrantVSwitchesToCen"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGrantVSwitchesToCenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListGrantVSwitchesToCen(request *ListGrantVSwitchesToCenRequest) (_result *ListGrantVSwitchesToCenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListGrantVSwitchesToCenResponse{}
	_body, _err := client.ListGrantVSwitchesToCenWithOptions(request, runtime)
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
	query["NextToken"] = request.NextToken
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["PageSize"] = request.PageSize
	query["ResourceId"] = request.ResourceId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["ResourceType"] = request.ResourceType
	query["Tag"] = request.Tag
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2017-09-12"),
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

func (client *Client) ListTrafficMarkingPoliciesWithOptions(request *ListTrafficMarkingPoliciesRequest, runtime *util.RuntimeOptions) (_result *ListTrafficMarkingPoliciesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["MaxResults"] = request.MaxResults
	query["NextToken"] = request.NextToken
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TrafficMarkingPolicyDescription"] = request.TrafficMarkingPolicyDescription
	query["TrafficMarkingPolicyId"] = request.TrafficMarkingPolicyId
	query["TrafficMarkingPolicyName"] = request.TrafficMarkingPolicyName
	query["TransitRouterId"] = request.TransitRouterId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTrafficMarkingPolicies"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTrafficMarkingPoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTrafficMarkingPolicies(request *ListTrafficMarkingPoliciesRequest) (_result *ListTrafficMarkingPoliciesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTrafficMarkingPoliciesResponse{}
	_body, _err := client.ListTrafficMarkingPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTransitRouterAvailableResourceWithOptions(request *ListTransitRouterAvailableResourceRequest, runtime *util.RuntimeOptions) (_result *ListTransitRouterAvailableResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTransitRouterAvailableResource"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTransitRouterAvailableResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTransitRouterAvailableResource(request *ListTransitRouterAvailableResourceRequest) (_result *ListTransitRouterAvailableResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTransitRouterAvailableResourceResponse{}
	_body, _err := client.ListTransitRouterAvailableResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTransitRouterPeerAttachmentsWithOptions(request *ListTransitRouterPeerAttachmentsRequest, runtime *util.RuntimeOptions) (_result *ListTransitRouterPeerAttachmentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["CenId"] = request.CenId
	query["MaxResults"] = request.MaxResults
	query["NextToken"] = request.NextToken
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	query["TransitRouterId"] = request.TransitRouterId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTransitRouterPeerAttachments"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTransitRouterPeerAttachmentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTransitRouterPeerAttachments(request *ListTransitRouterPeerAttachmentsRequest) (_result *ListTransitRouterPeerAttachmentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTransitRouterPeerAttachmentsResponse{}
	_body, _err := client.ListTransitRouterPeerAttachmentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTransitRouterRouteEntriesWithOptions(request *ListTransitRouterRouteEntriesRequest, runtime *util.RuntimeOptions) (_result *ListTransitRouterRouteEntriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["MaxResults"] = request.MaxResults
	query["NextToken"] = request.NextToken
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TransitRouterRouteEntryDestinationCidrBlock"] = request.TransitRouterRouteEntryDestinationCidrBlock
	query["TransitRouterRouteEntryIds"] = request.TransitRouterRouteEntryIds
	query["TransitRouterRouteEntryNames"] = request.TransitRouterRouteEntryNames
	query["TransitRouterRouteEntryStatus"] = request.TransitRouterRouteEntryStatus
	query["TransitRouterRouteTableId"] = request.TransitRouterRouteTableId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTransitRouterRouteEntries"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTransitRouterRouteEntriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTransitRouterRouteEntries(request *ListTransitRouterRouteEntriesRequest) (_result *ListTransitRouterRouteEntriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTransitRouterRouteEntriesResponse{}
	_body, _err := client.ListTransitRouterRouteEntriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTransitRouterRouteTableAssociationsWithOptions(request *ListTransitRouterRouteTableAssociationsRequest, runtime *util.RuntimeOptions) (_result *ListTransitRouterRouteTableAssociationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["MaxResults"] = request.MaxResults
	query["NextToken"] = request.NextToken
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	query["TransitRouterRouteTableId"] = request.TransitRouterRouteTableId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTransitRouterRouteTableAssociations"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTransitRouterRouteTableAssociationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTransitRouterRouteTableAssociations(request *ListTransitRouterRouteTableAssociationsRequest) (_result *ListTransitRouterRouteTableAssociationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTransitRouterRouteTableAssociationsResponse{}
	_body, _err := client.ListTransitRouterRouteTableAssociationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTransitRouterRouteTablePropagationsWithOptions(request *ListTransitRouterRouteTablePropagationsRequest, runtime *util.RuntimeOptions) (_result *ListTransitRouterRouteTablePropagationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["MaxResults"] = request.MaxResults
	query["NextToken"] = request.NextToken
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	query["TransitRouterRouteTableId"] = request.TransitRouterRouteTableId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTransitRouterRouteTablePropagations"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTransitRouterRouteTablePropagationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTransitRouterRouteTablePropagations(request *ListTransitRouterRouteTablePropagationsRequest) (_result *ListTransitRouterRouteTablePropagationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTransitRouterRouteTablePropagationsResponse{}
	_body, _err := client.ListTransitRouterRouteTablePropagationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTransitRouterRouteTablesWithOptions(request *ListTransitRouterRouteTablesRequest, runtime *util.RuntimeOptions) (_result *ListTransitRouterRouteTablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["MaxResults"] = request.MaxResults
	query["NextToken"] = request.NextToken
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TransitRouterId"] = request.TransitRouterId
	query["TransitRouterRouteTableIds"] = request.TransitRouterRouteTableIds
	query["TransitRouterRouteTableNames"] = request.TransitRouterRouteTableNames
	query["TransitRouterRouteTableStatus"] = request.TransitRouterRouteTableStatus
	query["TransitRouterRouteTableType"] = request.TransitRouterRouteTableType
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTransitRouterRouteTables"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTransitRouterRouteTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTransitRouterRouteTables(request *ListTransitRouterRouteTablesRequest) (_result *ListTransitRouterRouteTablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTransitRouterRouteTablesResponse{}
	_body, _err := client.ListTransitRouterRouteTablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTransitRouterVbrAttachmentsWithOptions(request *ListTransitRouterVbrAttachmentsRequest, runtime *util.RuntimeOptions) (_result *ListTransitRouterVbrAttachmentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["CenId"] = request.CenId
	query["MaxResults"] = request.MaxResults
	query["NextToken"] = request.NextToken
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	query["TransitRouterId"] = request.TransitRouterId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTransitRouterVbrAttachments"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTransitRouterVbrAttachmentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTransitRouterVbrAttachments(request *ListTransitRouterVbrAttachmentsRequest) (_result *ListTransitRouterVbrAttachmentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTransitRouterVbrAttachmentsResponse{}
	_body, _err := client.ListTransitRouterVbrAttachmentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTransitRouterVpcAttachmentsWithOptions(request *ListTransitRouterVpcAttachmentsRequest, runtime *util.RuntimeOptions) (_result *ListTransitRouterVpcAttachmentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["CenId"] = request.CenId
	query["MaxResults"] = request.MaxResults
	query["NextToken"] = request.NextToken
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	query["TransitRouterId"] = request.TransitRouterId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTransitRouterVpcAttachments"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTransitRouterVpcAttachmentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTransitRouterVpcAttachments(request *ListTransitRouterVpcAttachmentsRequest) (_result *ListTransitRouterVpcAttachmentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTransitRouterVpcAttachmentsResponse{}
	_body, _err := client.ListTransitRouterVpcAttachmentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTransitRoutersWithOptions(request *ListTransitRoutersRequest, runtime *util.RuntimeOptions) (_result *ListTransitRoutersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["CenId"] = request.CenId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TransitRouterId"] = request.TransitRouterId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTransitRouters"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTransitRoutersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTransitRouters(request *ListTransitRoutersRequest) (_result *ListTransitRoutersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTransitRoutersResponse{}
	_body, _err := client.ListTransitRoutersWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	query["CenId"] = request.CenId
	query["Description"] = request.Description
	query["Name"] = request.Name
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ProtectionLevel"] = request.ProtectionLevel
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyCenAttribute"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyCenAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["CenBandwidthPackageId"] = request.CenBandwidthPackageId
	query["Description"] = request.Description
	query["Name"] = request.Name
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyCenBandwidthPackageAttribute"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyCenBandwidthPackageAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["Bandwidth"] = request.Bandwidth
	query["CenBandwidthPackageId"] = request.CenBandwidthPackageId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyCenBandwidthPackageSpec"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyCenBandwidthPackageSpecResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["AsPathMatchMode"] = request.AsPathMatchMode
	query["CenId"] = request.CenId
	query["CenRegionId"] = request.CenRegionId
	query["CidrMatchMode"] = request.CidrMatchMode
	query["CommunityMatchMode"] = request.CommunityMatchMode
	query["CommunityOperateMode"] = request.CommunityOperateMode
	query["Description"] = request.Description
	query["DestinationChildInstanceTypes"] = request.DestinationChildInstanceTypes
	query["DestinationCidrBlocks"] = request.DestinationCidrBlocks
	query["DestinationInstanceIds"] = request.DestinationInstanceIds
	query["DestinationInstanceIdsReverseMatch"] = request.DestinationInstanceIdsReverseMatch
	query["DestinationRouteTableIds"] = request.DestinationRouteTableIds
	query["MapResult"] = request.MapResult
	query["MatchAsns"] = request.MatchAsns
	query["MatchCommunitySet"] = request.MatchCommunitySet
	query["NextPriority"] = request.NextPriority
	query["OperateCommunitySet"] = request.OperateCommunitySet
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["Preference"] = request.Preference
	query["PrependAsPath"] = request.PrependAsPath
	query["Priority"] = request.Priority
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["RouteMapId"] = request.RouteMapId
	query["RouteTypes"] = request.RouteTypes
	query["SourceChildInstanceTypes"] = request.SourceChildInstanceTypes
	query["SourceInstanceIds"] = request.SourceInstanceIds
	query["SourceInstanceIdsReverseMatch"] = request.SourceInstanceIdsReverseMatch
	query["SourceRegionIds"] = request.SourceRegionIds
	query["SourceRouteTableIds"] = request.SourceRouteTableIds
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyCenRouteMap"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyCenRouteMapResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["CenId"] = request.CenId
	query["ClientToken"] = request.ClientToken
	query["Description"] = request.Description
	query["FlowLogId"] = request.FlowLogId
	query["FlowLogName"] = request.FlowLogName
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyFlowLogAttribute"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyFlowLogAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) MoveResourceGroupWithOptions(request *MoveResourceGroupRequest, runtime *util.RuntimeOptions) (_result *MoveResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["NewResourceGroupId"] = request.NewResourceGroupId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceId"] = request.ResourceId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["ResourceType"] = request.ResourceType
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("MoveResourceGroup"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) OpenTransitRouterServiceWithOptions(request *OpenTransitRouterServiceRequest, runtime *util.RuntimeOptions) (_result *OpenTransitRouterServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OpenTransitRouterService"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenTransitRouterServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenTransitRouterService(request *OpenTransitRouterServiceRequest) (_result *OpenTransitRouterServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenTransitRouterServiceResponse{}
	_body, _err := client.OpenTransitRouterServiceWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	query["CenId"] = request.CenId
	query["ChildInstanceId"] = request.ChildInstanceId
	query["ChildInstanceRegionId"] = request.ChildInstanceRegionId
	query["ChildInstanceRouteTableId"] = request.ChildInstanceRouteTableId
	query["ChildInstanceType"] = request.ChildInstanceType
	query["DestinationCidrBlock"] = request.DestinationCidrBlock
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PublishRouteEntries"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PublishRouteEntriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) RemoveTraficMatchRuleFromTrafficMarkingPolicyWithOptions(request *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest, runtime *util.RuntimeOptions) (_result *RemoveTraficMatchRuleFromTrafficMarkingPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TrafficMarkRuleIds"] = request.TrafficMarkRuleIds
	query["TrafficMarkingPolicyId"] = request.TrafficMarkingPolicyId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveTraficMatchRuleFromTrafficMarkingPolicy"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveTraficMatchRuleFromTrafficMarkingPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveTraficMatchRuleFromTrafficMarkingPolicy(request *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest) (_result *RemoveTraficMatchRuleFromTrafficMarkingPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveTraficMatchRuleFromTrafficMarkingPolicyResponse{}
	_body, _err := client.RemoveTraficMatchRuleFromTrafficMarkingPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReplaceTransitRouterRouteTableAssociationWithOptions(request *ReplaceTransitRouterRouteTableAssociationRequest, runtime *util.RuntimeOptions) (_result *ReplaceTransitRouterRouteTableAssociationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	query["TransitRouterRouteTableId"] = request.TransitRouterRouteTableId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReplaceTransitRouterRouteTableAssociation"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReplaceTransitRouterRouteTableAssociationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReplaceTransitRouterRouteTableAssociation(request *ReplaceTransitRouterRouteTableAssociationRequest) (_result *ReplaceTransitRouterRouteTableAssociationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReplaceTransitRouterRouteTableAssociationResponse{}
	_body, _err := client.ReplaceTransitRouterRouteTableAssociationWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	query["AccessRegionIds"] = request.AccessRegionIds
	query["CenId"] = request.CenId
	query["ClientToken"] = request.ClientToken
	query["Description"] = request.Description
	query["Host"] = request.Host
	query["HostRegionId"] = request.HostRegionId
	query["HostVpcId"] = request.HostVpcId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResolveAndRouteServiceInCen"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResolveAndRouteServiceInCenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) RevokeInstanceFromTransitRouterWithOptions(request *RevokeInstanceFromTransitRouterRequest, runtime *util.RuntimeOptions) (_result *RevokeInstanceFromTransitRouterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["CenId"] = request.CenId
	query["CenOwnerId"] = request.CenOwnerId
	query["InstanceId"] = request.InstanceId
	query["InstanceType"] = request.InstanceType
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RevokeInstanceFromTransitRouter"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RevokeInstanceFromTransitRouterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RevokeInstanceFromTransitRouter(request *RevokeInstanceFromTransitRouterRequest) (_result *RevokeInstanceFromTransitRouterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RevokeInstanceFromTransitRouterResponse{}
	_body, _err := client.RevokeInstanceFromTransitRouterWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	query["AccessRegionId"] = request.AccessRegionId
	query["CenId"] = request.CenId
	query["HostRegionId"] = request.HostRegionId
	query["HostVpcId"] = request.HostVpcId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RoutePrivateZoneInCenToVpc"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RoutePrivateZoneInCenToVpcResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["BandwidthLimit"] = request.BandwidthLimit
	query["CenId"] = request.CenId
	query["LocalRegionId"] = request.LocalRegionId
	query["OppositeRegionId"] = request.OppositeRegionId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetCenInterRegionBandwidthLimit"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetCenInterRegionBandwidthLimitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceId"] = request.ResourceId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["ResourceType"] = request.ResourceType
	query["Tag"] = request.Tag
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2017-09-12"),
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

func (client *Client) TempUpgradeCenBandwidthPackageSpecWithOptions(request *TempUpgradeCenBandwidthPackageSpecRequest, runtime *util.RuntimeOptions) (_result *TempUpgradeCenBandwidthPackageSpecResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Bandwidth"] = request.Bandwidth
	query["CenBandwidthPackageId"] = request.CenBandwidthPackageId
	query["EndTime"] = request.EndTime
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TempUpgradeCenBandwidthPackageSpec"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TempUpgradeCenBandwidthPackageSpecResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["CenBandwidthPackageId"] = request.CenBandwidthPackageId
	query["CenId"] = request.CenId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnassociateCenBandwidthPackage"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnassociateCenBandwidthPackageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["AccessRegionId"] = request.AccessRegionId
	query["CenId"] = request.CenId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnroutePrivateZoneInCenToVpc"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnroutePrivateZoneInCenToVpcResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["All"] = request.All
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceId"] = request.ResourceId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["ResourceType"] = request.ResourceType
	query["TagKey"] = request.TagKey
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2017-09-12"),
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

func (client *Client) UpdateCenInterRegionTrafficQosPolicyAttributeWithOptions(request *UpdateCenInterRegionTrafficQosPolicyAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateCenInterRegionTrafficQosPolicyAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TrafficQosPolicyDescription"] = request.TrafficQosPolicyDescription
	query["TrafficQosPolicyId"] = request.TrafficQosPolicyId
	query["TrafficQosPolicyName"] = request.TrafficQosPolicyName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateCenInterRegionTrafficQosPolicyAttribute"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateCenInterRegionTrafficQosPolicyAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCenInterRegionTrafficQosPolicyAttribute(request *UpdateCenInterRegionTrafficQosPolicyAttributeRequest) (_result *UpdateCenInterRegionTrafficQosPolicyAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCenInterRegionTrafficQosPolicyAttributeResponse{}
	_body, _err := client.UpdateCenInterRegionTrafficQosPolicyAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCenInterRegionTrafficQosQueueAttributeWithOptions(request *UpdateCenInterRegionTrafficQosQueueAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateCenInterRegionTrafficQosQueueAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["Dscps"] = request.Dscps
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["QosQueueDescription"] = request.QosQueueDescription
	query["QosQueueId"] = request.QosQueueId
	query["QosQueueName"] = request.QosQueueName
	query["RemainBandwidthPercent"] = request.RemainBandwidthPercent
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateCenInterRegionTrafficQosQueueAttribute"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateCenInterRegionTrafficQosQueueAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCenInterRegionTrafficQosQueueAttribute(request *UpdateCenInterRegionTrafficQosQueueAttributeRequest) (_result *UpdateCenInterRegionTrafficQosQueueAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCenInterRegionTrafficQosQueueAttributeResponse{}
	_body, _err := client.UpdateCenInterRegionTrafficQosQueueAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTrafficMarkingPolicyAttributeWithOptions(request *UpdateTrafficMarkingPolicyAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateTrafficMarkingPolicyAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TrafficMarkingPolicyDescription"] = request.TrafficMarkingPolicyDescription
	query["TrafficMarkingPolicyId"] = request.TrafficMarkingPolicyId
	query["TrafficMarkingPolicyName"] = request.TrafficMarkingPolicyName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTrafficMarkingPolicyAttribute"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTrafficMarkingPolicyAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTrafficMarkingPolicyAttribute(request *UpdateTrafficMarkingPolicyAttributeRequest) (_result *UpdateTrafficMarkingPolicyAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateTrafficMarkingPolicyAttributeResponse{}
	_body, _err := client.UpdateTrafficMarkingPolicyAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTransitRouterWithOptions(request *UpdateTransitRouterRequest, runtime *util.RuntimeOptions) (_result *UpdateTransitRouterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TransitRouterDescription"] = request.TransitRouterDescription
	query["TransitRouterId"] = request.TransitRouterId
	query["TransitRouterName"] = request.TransitRouterName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTransitRouter"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTransitRouterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTransitRouter(request *UpdateTransitRouterRequest) (_result *UpdateTransitRouterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateTransitRouterResponse{}
	_body, _err := client.UpdateTransitRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTransitRouterPeerAttachmentAttributeWithOptions(request *UpdateTransitRouterPeerAttachmentAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateTransitRouterPeerAttachmentAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AutoPublishRouteEnabled"] = request.AutoPublishRouteEnabled
	query["Bandwidth"] = request.Bandwidth
	query["BandwidthType"] = request.BandwidthType
	query["CenBandwidthPackageId"] = request.CenBandwidthPackageId
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TransitRouterAttachmentDescription"] = request.TransitRouterAttachmentDescription
	query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	query["TransitRouterAttachmentName"] = request.TransitRouterAttachmentName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTransitRouterPeerAttachmentAttribute"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTransitRouterPeerAttachmentAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTransitRouterPeerAttachmentAttribute(request *UpdateTransitRouterPeerAttachmentAttributeRequest) (_result *UpdateTransitRouterPeerAttachmentAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateTransitRouterPeerAttachmentAttributeResponse{}
	_body, _err := client.UpdateTransitRouterPeerAttachmentAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTransitRouterRouteEntryWithOptions(request *UpdateTransitRouterRouteEntryRequest, runtime *util.RuntimeOptions) (_result *UpdateTransitRouterRouteEntryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TransitRouterRouteEntryDescription"] = request.TransitRouterRouteEntryDescription
	query["TransitRouterRouteEntryId"] = request.TransitRouterRouteEntryId
	query["TransitRouterRouteEntryName"] = request.TransitRouterRouteEntryName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTransitRouterRouteEntry"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTransitRouterRouteEntryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTransitRouterRouteEntry(request *UpdateTransitRouterRouteEntryRequest) (_result *UpdateTransitRouterRouteEntryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateTransitRouterRouteEntryResponse{}
	_body, _err := client.UpdateTransitRouterRouteEntryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTransitRouterRouteTableWithOptions(request *UpdateTransitRouterRouteTableRequest, runtime *util.RuntimeOptions) (_result *UpdateTransitRouterRouteTableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TransitRouterRouteTableDescription"] = request.TransitRouterRouteTableDescription
	query["TransitRouterRouteTableId"] = request.TransitRouterRouteTableId
	query["TransitRouterRouteTableName"] = request.TransitRouterRouteTableName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTransitRouterRouteTable"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTransitRouterRouteTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTransitRouterRouteTable(request *UpdateTransitRouterRouteTableRequest) (_result *UpdateTransitRouterRouteTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateTransitRouterRouteTableResponse{}
	_body, _err := client.UpdateTransitRouterRouteTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTransitRouterVbrAttachmentAttributeWithOptions(request *UpdateTransitRouterVbrAttachmentAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateTransitRouterVbrAttachmentAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TransitRouterAttachmentDescription"] = request.TransitRouterAttachmentDescription
	query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	query["TransitRouterAttachmentName"] = request.TransitRouterAttachmentName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTransitRouterVbrAttachmentAttribute"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTransitRouterVbrAttachmentAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTransitRouterVbrAttachmentAttribute(request *UpdateTransitRouterVbrAttachmentAttributeRequest) (_result *UpdateTransitRouterVbrAttachmentAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateTransitRouterVbrAttachmentAttributeResponse{}
	_body, _err := client.UpdateTransitRouterVbrAttachmentAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTransitRouterVpcAttachmentAttributeWithOptions(request *UpdateTransitRouterVpcAttachmentAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateTransitRouterVpcAttachmentAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TransitRouterAttachmentDescription"] = request.TransitRouterAttachmentDescription
	query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	query["TransitRouterAttachmentName"] = request.TransitRouterAttachmentName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTransitRouterVpcAttachmentAttribute"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTransitRouterVpcAttachmentAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTransitRouterVpcAttachmentAttribute(request *UpdateTransitRouterVpcAttachmentAttributeRequest) (_result *UpdateTransitRouterVpcAttachmentAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateTransitRouterVpcAttachmentAttributeResponse{}
	_body, _err := client.UpdateTransitRouterVpcAttachmentAttributeWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	query["CenId"] = request.CenId
	query["ChildInstanceId"] = request.ChildInstanceId
	query["ChildInstanceRegionId"] = request.ChildInstanceRegionId
	query["ChildInstanceRouteTableId"] = request.ChildInstanceRouteTableId
	query["ChildInstanceType"] = request.ChildInstanceType
	query["DestinationCidrBlock"] = request.DestinationCidrBlock
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("WithdrawPublishedRouteEntries"),
		Version:     tea.String("2017-09-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &WithdrawPublishedRouteEntriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
