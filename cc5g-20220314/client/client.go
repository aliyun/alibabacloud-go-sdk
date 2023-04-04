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

type AddDNSAuthorizationRuleRequest struct {
	ClientToken              *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description              *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DestinationIp            *string `json:"DestinationIp,omitempty" xml:"DestinationIp,omitempty"`
	DryRun                   *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	Name                     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SourceDNSIp              *string `json:"SourceDNSIp,omitempty" xml:"SourceDNSIp,omitempty"`
	WirelessCloudConnectorId *string `json:"WirelessCloudConnectorId,omitempty" xml:"WirelessCloudConnectorId,omitempty"`
}

func (s AddDNSAuthorizationRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s AddDNSAuthorizationRuleRequest) GoString() string {
	return s.String()
}

func (s *AddDNSAuthorizationRuleRequest) SetClientToken(v string) *AddDNSAuthorizationRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *AddDNSAuthorizationRuleRequest) SetDescription(v string) *AddDNSAuthorizationRuleRequest {
	s.Description = &v
	return s
}

func (s *AddDNSAuthorizationRuleRequest) SetDestinationIp(v string) *AddDNSAuthorizationRuleRequest {
	s.DestinationIp = &v
	return s
}

func (s *AddDNSAuthorizationRuleRequest) SetDryRun(v bool) *AddDNSAuthorizationRuleRequest {
	s.DryRun = &v
	return s
}

func (s *AddDNSAuthorizationRuleRequest) SetName(v string) *AddDNSAuthorizationRuleRequest {
	s.Name = &v
	return s
}

func (s *AddDNSAuthorizationRuleRequest) SetSourceDNSIp(v string) *AddDNSAuthorizationRuleRequest {
	s.SourceDNSIp = &v
	return s
}

func (s *AddDNSAuthorizationRuleRequest) SetWirelessCloudConnectorId(v string) *AddDNSAuthorizationRuleRequest {
	s.WirelessCloudConnectorId = &v
	return s
}

type AddDNSAuthorizationRuleResponseBody struct {
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddDNSAuthorizationRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddDNSAuthorizationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *AddDNSAuthorizationRuleResponseBody) SetAuthorizationRuleId(v string) *AddDNSAuthorizationRuleResponseBody {
	s.AuthorizationRuleId = &v
	return s
}

func (s *AddDNSAuthorizationRuleResponseBody) SetRequestId(v string) *AddDNSAuthorizationRuleResponseBody {
	s.RequestId = &v
	return s
}

type AddDNSAuthorizationRuleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddDNSAuthorizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddDNSAuthorizationRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s AddDNSAuthorizationRuleResponse) GoString() string {
	return s.String()
}

func (s *AddDNSAuthorizationRuleResponse) SetHeaders(v map[string]*string) *AddDNSAuthorizationRuleResponse {
	s.Headers = v
	return s
}

func (s *AddDNSAuthorizationRuleResponse) SetStatusCode(v int32) *AddDNSAuthorizationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDNSAuthorizationRuleResponse) SetBody(v *AddDNSAuthorizationRuleResponseBody) *AddDNSAuthorizationRuleResponse {
	s.Body = v
	return s
}

type AddGroupDnsAuthorizationRuleRequest struct {
	ClientToken                   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description                   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DestinationIp                 *string `json:"DestinationIp,omitempty" xml:"DestinationIp,omitempty"`
	DryRun                        *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	Name                          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SourceDNSIp                   *string `json:"SourceDNSIp,omitempty" xml:"SourceDNSIp,omitempty"`
	WirelessCloudConnectorGroupId *string `json:"WirelessCloudConnectorGroupId,omitempty" xml:"WirelessCloudConnectorGroupId,omitempty"`
}

func (s AddGroupDnsAuthorizationRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s AddGroupDnsAuthorizationRuleRequest) GoString() string {
	return s.String()
}

func (s *AddGroupDnsAuthorizationRuleRequest) SetClientToken(v string) *AddGroupDnsAuthorizationRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *AddGroupDnsAuthorizationRuleRequest) SetDescription(v string) *AddGroupDnsAuthorizationRuleRequest {
	s.Description = &v
	return s
}

func (s *AddGroupDnsAuthorizationRuleRequest) SetDestinationIp(v string) *AddGroupDnsAuthorizationRuleRequest {
	s.DestinationIp = &v
	return s
}

func (s *AddGroupDnsAuthorizationRuleRequest) SetDryRun(v bool) *AddGroupDnsAuthorizationRuleRequest {
	s.DryRun = &v
	return s
}

func (s *AddGroupDnsAuthorizationRuleRequest) SetName(v string) *AddGroupDnsAuthorizationRuleRequest {
	s.Name = &v
	return s
}

func (s *AddGroupDnsAuthorizationRuleRequest) SetSourceDNSIp(v string) *AddGroupDnsAuthorizationRuleRequest {
	s.SourceDNSIp = &v
	return s
}

func (s *AddGroupDnsAuthorizationRuleRequest) SetWirelessCloudConnectorGroupId(v string) *AddGroupDnsAuthorizationRuleRequest {
	s.WirelessCloudConnectorGroupId = &v
	return s
}

type AddGroupDnsAuthorizationRuleResponseBody struct {
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddGroupDnsAuthorizationRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddGroupDnsAuthorizationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *AddGroupDnsAuthorizationRuleResponseBody) SetAuthorizationRuleId(v string) *AddGroupDnsAuthorizationRuleResponseBody {
	s.AuthorizationRuleId = &v
	return s
}

func (s *AddGroupDnsAuthorizationRuleResponseBody) SetRequestId(v string) *AddGroupDnsAuthorizationRuleResponseBody {
	s.RequestId = &v
	return s
}

type AddGroupDnsAuthorizationRuleResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddGroupDnsAuthorizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddGroupDnsAuthorizationRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s AddGroupDnsAuthorizationRuleResponse) GoString() string {
	return s.String()
}

func (s *AddGroupDnsAuthorizationRuleResponse) SetHeaders(v map[string]*string) *AddGroupDnsAuthorizationRuleResponse {
	s.Headers = v
	return s
}

func (s *AddGroupDnsAuthorizationRuleResponse) SetStatusCode(v int32) *AddGroupDnsAuthorizationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *AddGroupDnsAuthorizationRuleResponse) SetBody(v *AddGroupDnsAuthorizationRuleResponseBody) *AddGroupDnsAuthorizationRuleResponse {
	s.Body = v
	return s
}

type AddWirelessCloudConnectorToGroupRequest struct {
	ClientToken                   *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                        *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	RegionId                      *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	WirelessCloudConnectorGroupId *string   `json:"WirelessCloudConnectorGroupId,omitempty" xml:"WirelessCloudConnectorGroupId,omitempty"`
	WirelessCloudConnectorIds     []*string `json:"WirelessCloudConnectorIds,omitempty" xml:"WirelessCloudConnectorIds,omitempty" type:"Repeated"`
}

func (s AddWirelessCloudConnectorToGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddWirelessCloudConnectorToGroupRequest) GoString() string {
	return s.String()
}

func (s *AddWirelessCloudConnectorToGroupRequest) SetClientToken(v string) *AddWirelessCloudConnectorToGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *AddWirelessCloudConnectorToGroupRequest) SetDryRun(v bool) *AddWirelessCloudConnectorToGroupRequest {
	s.DryRun = &v
	return s
}

func (s *AddWirelessCloudConnectorToGroupRequest) SetRegionId(v string) *AddWirelessCloudConnectorToGroupRequest {
	s.RegionId = &v
	return s
}

func (s *AddWirelessCloudConnectorToGroupRequest) SetWirelessCloudConnectorGroupId(v string) *AddWirelessCloudConnectorToGroupRequest {
	s.WirelessCloudConnectorGroupId = &v
	return s
}

func (s *AddWirelessCloudConnectorToGroupRequest) SetWirelessCloudConnectorIds(v []*string) *AddWirelessCloudConnectorToGroupRequest {
	s.WirelessCloudConnectorIds = v
	return s
}

type AddWirelessCloudConnectorToGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddWirelessCloudConnectorToGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddWirelessCloudConnectorToGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddWirelessCloudConnectorToGroupResponseBody) SetRequestId(v string) *AddWirelessCloudConnectorToGroupResponseBody {
	s.RequestId = &v
	return s
}

type AddWirelessCloudConnectorToGroupResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddWirelessCloudConnectorToGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddWirelessCloudConnectorToGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AddWirelessCloudConnectorToGroupResponse) GoString() string {
	return s.String()
}

func (s *AddWirelessCloudConnectorToGroupResponse) SetHeaders(v map[string]*string) *AddWirelessCloudConnectorToGroupResponse {
	s.Headers = v
	return s
}

func (s *AddWirelessCloudConnectorToGroupResponse) SetStatusCode(v int32) *AddWirelessCloudConnectorToGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddWirelessCloudConnectorToGroupResponse) SetBody(v *AddWirelessCloudConnectorToGroupResponseBody) *AddWirelessCloudConnectorToGroupResponse {
	s.Body = v
	return s
}

type AttachVpcToNetLinkRequest struct {
	ClientToken              *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                   *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	NetLinkId                *string   `json:"NetLinkId,omitempty" xml:"NetLinkId,omitempty"`
	RegionId                 *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VSwitches                []*string `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
	VpcId                    *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	WirelessCloudConnectorId *string   `json:"WirelessCloudConnectorId,omitempty" xml:"WirelessCloudConnectorId,omitempty"`
}

func (s AttachVpcToNetLinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachVpcToNetLinkRequest) GoString() string {
	return s.String()
}

func (s *AttachVpcToNetLinkRequest) SetClientToken(v string) *AttachVpcToNetLinkRequest {
	s.ClientToken = &v
	return s
}

func (s *AttachVpcToNetLinkRequest) SetDryRun(v bool) *AttachVpcToNetLinkRequest {
	s.DryRun = &v
	return s
}

func (s *AttachVpcToNetLinkRequest) SetNetLinkId(v string) *AttachVpcToNetLinkRequest {
	s.NetLinkId = &v
	return s
}

func (s *AttachVpcToNetLinkRequest) SetRegionId(v string) *AttachVpcToNetLinkRequest {
	s.RegionId = &v
	return s
}

func (s *AttachVpcToNetLinkRequest) SetVSwitches(v []*string) *AttachVpcToNetLinkRequest {
	s.VSwitches = v
	return s
}

func (s *AttachVpcToNetLinkRequest) SetVpcId(v string) *AttachVpcToNetLinkRequest {
	s.VpcId = &v
	return s
}

func (s *AttachVpcToNetLinkRequest) SetWirelessCloudConnectorId(v string) *AttachVpcToNetLinkRequest {
	s.WirelessCloudConnectorId = &v
	return s
}

type AttachVpcToNetLinkResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachVpcToNetLinkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachVpcToNetLinkResponseBody) GoString() string {
	return s.String()
}

func (s *AttachVpcToNetLinkResponseBody) SetRequestId(v string) *AttachVpcToNetLinkResponseBody {
	s.RequestId = &v
	return s
}

type AttachVpcToNetLinkResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AttachVpcToNetLinkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttachVpcToNetLinkResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachVpcToNetLinkResponse) GoString() string {
	return s.String()
}

func (s *AttachVpcToNetLinkResponse) SetHeaders(v map[string]*string) *AttachVpcToNetLinkResponse {
	s.Headers = v
	return s
}

func (s *AttachVpcToNetLinkResponse) SetStatusCode(v int32) *AttachVpcToNetLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachVpcToNetLinkResponse) SetBody(v *AttachVpcToNetLinkResponseBody) *AttachVpcToNetLinkResponse {
	s.Body = v
	return s
}

type CreateAuthorizationRuleRequest struct {
	ClientToken              *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description              *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Destination              *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	DestinationPort          *string `json:"DestinationPort,omitempty" xml:"DestinationPort,omitempty"`
	DestinationType          *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	DryRun                   *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	Name                     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Policy                   *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Protocol                 *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	SourceCidr               *string `json:"SourceCidr,omitempty" xml:"SourceCidr,omitempty"`
	WirelessCloudConnectorId *string `json:"WirelessCloudConnectorId,omitempty" xml:"WirelessCloudConnectorId,omitempty"`
}

func (s CreateAuthorizationRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAuthorizationRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateAuthorizationRuleRequest) SetClientToken(v string) *CreateAuthorizationRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAuthorizationRuleRequest) SetDescription(v string) *CreateAuthorizationRuleRequest {
	s.Description = &v
	return s
}

func (s *CreateAuthorizationRuleRequest) SetDestination(v string) *CreateAuthorizationRuleRequest {
	s.Destination = &v
	return s
}

func (s *CreateAuthorizationRuleRequest) SetDestinationPort(v string) *CreateAuthorizationRuleRequest {
	s.DestinationPort = &v
	return s
}

func (s *CreateAuthorizationRuleRequest) SetDestinationType(v string) *CreateAuthorizationRuleRequest {
	s.DestinationType = &v
	return s
}

func (s *CreateAuthorizationRuleRequest) SetDryRun(v bool) *CreateAuthorizationRuleRequest {
	s.DryRun = &v
	return s
}

func (s *CreateAuthorizationRuleRequest) SetName(v string) *CreateAuthorizationRuleRequest {
	s.Name = &v
	return s
}

func (s *CreateAuthorizationRuleRequest) SetPolicy(v string) *CreateAuthorizationRuleRequest {
	s.Policy = &v
	return s
}

func (s *CreateAuthorizationRuleRequest) SetProtocol(v string) *CreateAuthorizationRuleRequest {
	s.Protocol = &v
	return s
}

func (s *CreateAuthorizationRuleRequest) SetSourceCidr(v string) *CreateAuthorizationRuleRequest {
	s.SourceCidr = &v
	return s
}

func (s *CreateAuthorizationRuleRequest) SetWirelessCloudConnectorId(v string) *CreateAuthorizationRuleRequest {
	s.WirelessCloudConnectorId = &v
	return s
}

type CreateAuthorizationRuleResponseBody struct {
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAuthorizationRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAuthorizationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAuthorizationRuleResponseBody) SetAuthorizationRuleId(v string) *CreateAuthorizationRuleResponseBody {
	s.AuthorizationRuleId = &v
	return s
}

func (s *CreateAuthorizationRuleResponseBody) SetRequestId(v string) *CreateAuthorizationRuleResponseBody {
	s.RequestId = &v
	return s
}

type CreateAuthorizationRuleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAuthorizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAuthorizationRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAuthorizationRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateAuthorizationRuleResponse) SetHeaders(v map[string]*string) *CreateAuthorizationRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateAuthorizationRuleResponse) SetStatusCode(v int32) *CreateAuthorizationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAuthorizationRuleResponse) SetBody(v *CreateAuthorizationRuleResponseBody) *CreateAuthorizationRuleResponse {
	s.Body = v
	return s
}

type CreateBatchOperateCardsTaskRequest struct {
	ClientToken               *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description               *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	DryRun                    *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EffectType                *string   `json:"EffectType,omitempty" xml:"EffectType,omitempty"`
	Iccids                    []*string `json:"Iccids,omitempty" xml:"Iccids,omitempty" type:"Repeated"`
	IccidsOssFilePath         *string   `json:"IccidsOssFilePath,omitempty" xml:"IccidsOssFilePath,omitempty"`
	Name                      *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	OperateType               *string   `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	RegionId                  *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Threshold                 *int64    `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	WirelessCloudConnectorIds []*string `json:"WirelessCloudConnectorIds,omitempty" xml:"WirelessCloudConnectorIds,omitempty" type:"Repeated"`
}

func (s CreateBatchOperateCardsTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBatchOperateCardsTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateBatchOperateCardsTaskRequest) SetClientToken(v string) *CreateBatchOperateCardsTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateBatchOperateCardsTaskRequest) SetDescription(v string) *CreateBatchOperateCardsTaskRequest {
	s.Description = &v
	return s
}

func (s *CreateBatchOperateCardsTaskRequest) SetDryRun(v bool) *CreateBatchOperateCardsTaskRequest {
	s.DryRun = &v
	return s
}

func (s *CreateBatchOperateCardsTaskRequest) SetEffectType(v string) *CreateBatchOperateCardsTaskRequest {
	s.EffectType = &v
	return s
}

func (s *CreateBatchOperateCardsTaskRequest) SetIccids(v []*string) *CreateBatchOperateCardsTaskRequest {
	s.Iccids = v
	return s
}

func (s *CreateBatchOperateCardsTaskRequest) SetIccidsOssFilePath(v string) *CreateBatchOperateCardsTaskRequest {
	s.IccidsOssFilePath = &v
	return s
}

func (s *CreateBatchOperateCardsTaskRequest) SetName(v string) *CreateBatchOperateCardsTaskRequest {
	s.Name = &v
	return s
}

func (s *CreateBatchOperateCardsTaskRequest) SetOperateType(v string) *CreateBatchOperateCardsTaskRequest {
	s.OperateType = &v
	return s
}

func (s *CreateBatchOperateCardsTaskRequest) SetRegionId(v string) *CreateBatchOperateCardsTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CreateBatchOperateCardsTaskRequest) SetThreshold(v int64) *CreateBatchOperateCardsTaskRequest {
	s.Threshold = &v
	return s
}

func (s *CreateBatchOperateCardsTaskRequest) SetWirelessCloudConnectorIds(v []*string) *CreateBatchOperateCardsTaskRequest {
	s.WirelessCloudConnectorIds = v
	return s
}

type CreateBatchOperateCardsTaskResponseBody struct {
	BatchOperateCardsTaskId  *string `json:"BatchOperateCardsTaskId,omitempty" xml:"BatchOperateCardsTaskId,omitempty"`
	OperateResultOssFilePath *string `json:"OperateResultOssFilePath,omitempty" xml:"OperateResultOssFilePath,omitempty"`
	RequestId                *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBatchOperateCardsTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBatchOperateCardsTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBatchOperateCardsTaskResponseBody) SetBatchOperateCardsTaskId(v string) *CreateBatchOperateCardsTaskResponseBody {
	s.BatchOperateCardsTaskId = &v
	return s
}

func (s *CreateBatchOperateCardsTaskResponseBody) SetOperateResultOssFilePath(v string) *CreateBatchOperateCardsTaskResponseBody {
	s.OperateResultOssFilePath = &v
	return s
}

func (s *CreateBatchOperateCardsTaskResponseBody) SetRequestId(v string) *CreateBatchOperateCardsTaskResponseBody {
	s.RequestId = &v
	return s
}

type CreateBatchOperateCardsTaskResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateBatchOperateCardsTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateBatchOperateCardsTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBatchOperateCardsTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateBatchOperateCardsTaskResponse) SetHeaders(v map[string]*string) *CreateBatchOperateCardsTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateBatchOperateCardsTaskResponse) SetStatusCode(v int32) *CreateBatchOperateCardsTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBatchOperateCardsTaskResponse) SetBody(v *CreateBatchOperateCardsTaskResponseBody) *CreateBatchOperateCardsTaskResponse {
	s.Body = v
	return s
}

type CreateGroupAuthorizationRuleRequest struct {
	ClientToken                   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description                   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Destination                   *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	DestinationPort               *string `json:"DestinationPort,omitempty" xml:"DestinationPort,omitempty"`
	DestinationType               *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	DryRun                        *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	Name                          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Policy                        *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Protocol                      *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	SourceCidr                    *string `json:"SourceCidr,omitempty" xml:"SourceCidr,omitempty"`
	WirelessCloudConnectorGroupId *string `json:"WirelessCloudConnectorGroupId,omitempty" xml:"WirelessCloudConnectorGroupId,omitempty"`
}

func (s CreateGroupAuthorizationRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupAuthorizationRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupAuthorizationRuleRequest) SetClientToken(v string) *CreateGroupAuthorizationRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateGroupAuthorizationRuleRequest) SetDescription(v string) *CreateGroupAuthorizationRuleRequest {
	s.Description = &v
	return s
}

func (s *CreateGroupAuthorizationRuleRequest) SetDestination(v string) *CreateGroupAuthorizationRuleRequest {
	s.Destination = &v
	return s
}

func (s *CreateGroupAuthorizationRuleRequest) SetDestinationPort(v string) *CreateGroupAuthorizationRuleRequest {
	s.DestinationPort = &v
	return s
}

func (s *CreateGroupAuthorizationRuleRequest) SetDestinationType(v string) *CreateGroupAuthorizationRuleRequest {
	s.DestinationType = &v
	return s
}

func (s *CreateGroupAuthorizationRuleRequest) SetDryRun(v bool) *CreateGroupAuthorizationRuleRequest {
	s.DryRun = &v
	return s
}

func (s *CreateGroupAuthorizationRuleRequest) SetName(v string) *CreateGroupAuthorizationRuleRequest {
	s.Name = &v
	return s
}

func (s *CreateGroupAuthorizationRuleRequest) SetPolicy(v string) *CreateGroupAuthorizationRuleRequest {
	s.Policy = &v
	return s
}

func (s *CreateGroupAuthorizationRuleRequest) SetProtocol(v string) *CreateGroupAuthorizationRuleRequest {
	s.Protocol = &v
	return s
}

func (s *CreateGroupAuthorizationRuleRequest) SetSourceCidr(v string) *CreateGroupAuthorizationRuleRequest {
	s.SourceCidr = &v
	return s
}

func (s *CreateGroupAuthorizationRuleRequest) SetWirelessCloudConnectorGroupId(v string) *CreateGroupAuthorizationRuleRequest {
	s.WirelessCloudConnectorGroupId = &v
	return s
}

type CreateGroupAuthorizationRuleResponseBody struct {
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateGroupAuthorizationRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupAuthorizationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupAuthorizationRuleResponseBody) SetAuthorizationRuleId(v string) *CreateGroupAuthorizationRuleResponseBody {
	s.AuthorizationRuleId = &v
	return s
}

func (s *CreateGroupAuthorizationRuleResponseBody) SetRequestId(v string) *CreateGroupAuthorizationRuleResponseBody {
	s.RequestId = &v
	return s
}

type CreateGroupAuthorizationRuleResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateGroupAuthorizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGroupAuthorizationRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupAuthorizationRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupAuthorizationRuleResponse) SetHeaders(v map[string]*string) *CreateGroupAuthorizationRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateGroupAuthorizationRuleResponse) SetStatusCode(v int32) *CreateGroupAuthorizationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGroupAuthorizationRuleResponse) SetBody(v *CreateGroupAuthorizationRuleResponseBody) *CreateGroupAuthorizationRuleResponse {
	s.Body = v
	return s
}

type CreateIoTCloudConnectorBackhaulRouteRequest struct {
	ClientToken              *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                   *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	NetLinkId                *string `json:"NetLinkId,omitempty" xml:"NetLinkId,omitempty"`
	WirelessCloudConnectorId *string `json:"WirelessCloudConnectorId,omitempty" xml:"WirelessCloudConnectorId,omitempty"`
}

func (s CreateIoTCloudConnectorBackhaulRouteRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIoTCloudConnectorBackhaulRouteRequest) GoString() string {
	return s.String()
}

func (s *CreateIoTCloudConnectorBackhaulRouteRequest) SetClientToken(v string) *CreateIoTCloudConnectorBackhaulRouteRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateIoTCloudConnectorBackhaulRouteRequest) SetDryRun(v bool) *CreateIoTCloudConnectorBackhaulRouteRequest {
	s.DryRun = &v
	return s
}

func (s *CreateIoTCloudConnectorBackhaulRouteRequest) SetNetLinkId(v string) *CreateIoTCloudConnectorBackhaulRouteRequest {
	s.NetLinkId = &v
	return s
}

func (s *CreateIoTCloudConnectorBackhaulRouteRequest) SetWirelessCloudConnectorId(v string) *CreateIoTCloudConnectorBackhaulRouteRequest {
	s.WirelessCloudConnectorId = &v
	return s
}

type CreateIoTCloudConnectorBackhaulRouteResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateIoTCloudConnectorBackhaulRouteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIoTCloudConnectorBackhaulRouteResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIoTCloudConnectorBackhaulRouteResponseBody) SetRequestId(v string) *CreateIoTCloudConnectorBackhaulRouteResponseBody {
	s.RequestId = &v
	return s
}

type CreateIoTCloudConnectorBackhaulRouteResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateIoTCloudConnectorBackhaulRouteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateIoTCloudConnectorBackhaulRouteResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIoTCloudConnectorBackhaulRouteResponse) GoString() string {
	return s.String()
}

func (s *CreateIoTCloudConnectorBackhaulRouteResponse) SetHeaders(v map[string]*string) *CreateIoTCloudConnectorBackhaulRouteResponse {
	s.Headers = v
	return s
}

func (s *CreateIoTCloudConnectorBackhaulRouteResponse) SetStatusCode(v int32) *CreateIoTCloudConnectorBackhaulRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIoTCloudConnectorBackhaulRouteResponse) SetBody(v *CreateIoTCloudConnectorBackhaulRouteResponseBody) *CreateIoTCloudConnectorBackhaulRouteResponse {
	s.Body = v
	return s
}

type CreateWirelessCloudConnectorRequest struct {
	BusinessType *string                                        `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	ClientToken  *string                                        `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description  *string                                        `json:"Description,omitempty" xml:"Description,omitempty"`
	DryRun       *bool                                          `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	ISP          *string                                        `json:"ISP,omitempty" xml:"ISP,omitempty"`
	Name         *string                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	NetLinks     []*CreateWirelessCloudConnectorRequestNetLinks `json:"NetLinks,omitempty" xml:"NetLinks,omitempty" type:"Repeated"`
	RegionId     *string                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UseCase      *string                                        `json:"UseCase,omitempty" xml:"UseCase,omitempty"`
}

func (s CreateWirelessCloudConnectorRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWirelessCloudConnectorRequest) GoString() string {
	return s.String()
}

func (s *CreateWirelessCloudConnectorRequest) SetBusinessType(v string) *CreateWirelessCloudConnectorRequest {
	s.BusinessType = &v
	return s
}

func (s *CreateWirelessCloudConnectorRequest) SetClientToken(v string) *CreateWirelessCloudConnectorRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateWirelessCloudConnectorRequest) SetDescription(v string) *CreateWirelessCloudConnectorRequest {
	s.Description = &v
	return s
}

func (s *CreateWirelessCloudConnectorRequest) SetDryRun(v bool) *CreateWirelessCloudConnectorRequest {
	s.DryRun = &v
	return s
}

func (s *CreateWirelessCloudConnectorRequest) SetISP(v string) *CreateWirelessCloudConnectorRequest {
	s.ISP = &v
	return s
}

func (s *CreateWirelessCloudConnectorRequest) SetName(v string) *CreateWirelessCloudConnectorRequest {
	s.Name = &v
	return s
}

func (s *CreateWirelessCloudConnectorRequest) SetNetLinks(v []*CreateWirelessCloudConnectorRequestNetLinks) *CreateWirelessCloudConnectorRequest {
	s.NetLinks = v
	return s
}

func (s *CreateWirelessCloudConnectorRequest) SetRegionId(v string) *CreateWirelessCloudConnectorRequest {
	s.RegionId = &v
	return s
}

func (s *CreateWirelessCloudConnectorRequest) SetUseCase(v string) *CreateWirelessCloudConnectorRequest {
	s.UseCase = &v
	return s
}

type CreateWirelessCloudConnectorRequestNetLinks struct {
	APN      *string   `json:"APN,omitempty" xml:"APN,omitempty"`
	RegionId *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VSwitchs []*string `json:"VSwitchs,omitempty" xml:"VSwitchs,omitempty" type:"Repeated"`
	VpcId    *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateWirelessCloudConnectorRequestNetLinks) String() string {
	return tea.Prettify(s)
}

func (s CreateWirelessCloudConnectorRequestNetLinks) GoString() string {
	return s.String()
}

func (s *CreateWirelessCloudConnectorRequestNetLinks) SetAPN(v string) *CreateWirelessCloudConnectorRequestNetLinks {
	s.APN = &v
	return s
}

func (s *CreateWirelessCloudConnectorRequestNetLinks) SetRegionId(v string) *CreateWirelessCloudConnectorRequestNetLinks {
	s.RegionId = &v
	return s
}

func (s *CreateWirelessCloudConnectorRequestNetLinks) SetVSwitchs(v []*string) *CreateWirelessCloudConnectorRequestNetLinks {
	s.VSwitchs = v
	return s
}

func (s *CreateWirelessCloudConnectorRequestNetLinks) SetVpcId(v string) *CreateWirelessCloudConnectorRequestNetLinks {
	s.VpcId = &v
	return s
}

type CreateWirelessCloudConnectorResponseBody struct {
	RequestId                *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WirelessCloudConnectorId *string `json:"WirelessCloudConnectorId,omitempty" xml:"WirelessCloudConnectorId,omitempty"`
}

func (s CreateWirelessCloudConnectorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWirelessCloudConnectorResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWirelessCloudConnectorResponseBody) SetRequestId(v string) *CreateWirelessCloudConnectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWirelessCloudConnectorResponseBody) SetWirelessCloudConnectorId(v string) *CreateWirelessCloudConnectorResponseBody {
	s.WirelessCloudConnectorId = &v
	return s
}

type CreateWirelessCloudConnectorResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateWirelessCloudConnectorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateWirelessCloudConnectorResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWirelessCloudConnectorResponse) GoString() string {
	return s.String()
}

func (s *CreateWirelessCloudConnectorResponse) SetHeaders(v map[string]*string) *CreateWirelessCloudConnectorResponse {
	s.Headers = v
	return s
}

func (s *CreateWirelessCloudConnectorResponse) SetStatusCode(v int32) *CreateWirelessCloudConnectorResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWirelessCloudConnectorResponse) SetBody(v *CreateWirelessCloudConnectorResponseBody) *CreateWirelessCloudConnectorResponse {
	s.Body = v
	return s
}

type CreateWirelessCloudConnectorGroupRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DryRun      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateWirelessCloudConnectorGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWirelessCloudConnectorGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateWirelessCloudConnectorGroupRequest) SetClientToken(v string) *CreateWirelessCloudConnectorGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateWirelessCloudConnectorGroupRequest) SetDescription(v string) *CreateWirelessCloudConnectorGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateWirelessCloudConnectorGroupRequest) SetDryRun(v bool) *CreateWirelessCloudConnectorGroupRequest {
	s.DryRun = &v
	return s
}

func (s *CreateWirelessCloudConnectorGroupRequest) SetName(v string) *CreateWirelessCloudConnectorGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateWirelessCloudConnectorGroupRequest) SetRegionId(v string) *CreateWirelessCloudConnectorGroupRequest {
	s.RegionId = &v
	return s
}

type CreateWirelessCloudConnectorGroupResponseBody struct {
	RequestId                     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WirelessCloudConnectorGroupId *string `json:"WirelessCloudConnectorGroupId,omitempty" xml:"WirelessCloudConnectorGroupId,omitempty"`
}

func (s CreateWirelessCloudConnectorGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWirelessCloudConnectorGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWirelessCloudConnectorGroupResponseBody) SetRequestId(v string) *CreateWirelessCloudConnectorGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWirelessCloudConnectorGroupResponseBody) SetWirelessCloudConnectorGroupId(v string) *CreateWirelessCloudConnectorGroupResponseBody {
	s.WirelessCloudConnectorGroupId = &v
	return s
}

type CreateWirelessCloudConnectorGroupResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateWirelessCloudConnectorGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateWirelessCloudConnectorGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWirelessCloudConnectorGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateWirelessCloudConnectorGroupResponse) SetHeaders(v map[string]*string) *CreateWirelessCloudConnectorGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateWirelessCloudConnectorGroupResponse) SetStatusCode(v int32) *CreateWirelessCloudConnectorGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWirelessCloudConnectorGroupResponse) SetBody(v *CreateWirelessCloudConnectorGroupResponseBody) *CreateWirelessCloudConnectorGroupResponse {
	s.Body = v
	return s
}

type DeleteAuthorizationRuleRequest struct {
	AuthorizationRuleId      *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	ClientToken              *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                   *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	WirelessCloudConnectorId *string `json:"WirelessCloudConnectorId,omitempty" xml:"WirelessCloudConnectorId,omitempty"`
}

func (s DeleteAuthorizationRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAuthorizationRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteAuthorizationRuleRequest) SetAuthorizationRuleId(v string) *DeleteAuthorizationRuleRequest {
	s.AuthorizationRuleId = &v
	return s
}

func (s *DeleteAuthorizationRuleRequest) SetClientToken(v string) *DeleteAuthorizationRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteAuthorizationRuleRequest) SetDryRun(v bool) *DeleteAuthorizationRuleRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteAuthorizationRuleRequest) SetWirelessCloudConnectorId(v string) *DeleteAuthorizationRuleRequest {
	s.WirelessCloudConnectorId = &v
	return s
}

type DeleteAuthorizationRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAuthorizationRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAuthorizationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAuthorizationRuleResponseBody) SetRequestId(v string) *DeleteAuthorizationRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAuthorizationRuleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAuthorizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAuthorizationRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAuthorizationRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteAuthorizationRuleResponse) SetHeaders(v map[string]*string) *DeleteAuthorizationRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteAuthorizationRuleResponse) SetStatusCode(v int32) *DeleteAuthorizationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAuthorizationRuleResponse) SetBody(v *DeleteAuthorizationRuleResponseBody) *DeleteAuthorizationRuleResponse {
	s.Body = v
	return s
}

type DeleteBatchOperateCardsTaskRequest struct {
	BatchOperateCardsTaskId *string `json:"BatchOperateCardsTaskId,omitempty" xml:"BatchOperateCardsTaskId,omitempty"`
	ClientToken             *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                  *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	RegionId                *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteBatchOperateCardsTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBatchOperateCardsTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteBatchOperateCardsTaskRequest) SetBatchOperateCardsTaskId(v string) *DeleteBatchOperateCardsTaskRequest {
	s.BatchOperateCardsTaskId = &v
	return s
}

func (s *DeleteBatchOperateCardsTaskRequest) SetClientToken(v string) *DeleteBatchOperateCardsTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteBatchOperateCardsTaskRequest) SetDryRun(v bool) *DeleteBatchOperateCardsTaskRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteBatchOperateCardsTaskRequest) SetRegionId(v string) *DeleteBatchOperateCardsTaskRequest {
	s.RegionId = &v
	return s
}

type DeleteBatchOperateCardsTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBatchOperateCardsTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteBatchOperateCardsTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBatchOperateCardsTaskResponseBody) SetRequestId(v string) *DeleteBatchOperateCardsTaskResponseBody {
	s.RequestId = &v
	return s
}

type DeleteBatchOperateCardsTaskResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteBatchOperateCardsTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteBatchOperateCardsTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBatchOperateCardsTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteBatchOperateCardsTaskResponse) SetHeaders(v map[string]*string) *DeleteBatchOperateCardsTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteBatchOperateCardsTaskResponse) SetStatusCode(v int32) *DeleteBatchOperateCardsTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBatchOperateCardsTaskResponse) SetBody(v *DeleteBatchOperateCardsTaskResponseBody) *DeleteBatchOperateCardsTaskResponse {
	s.Body = v
	return s
}

type DeleteGroupAuthorizationRuleRequest struct {
	AuthorizationRuleId           *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	ClientToken                   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                        *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	WirelessCloudConnectorGroupId *string `json:"WirelessCloudConnectorGroupId,omitempty" xml:"WirelessCloudConnectorGroupId,omitempty"`
}

func (s DeleteGroupAuthorizationRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupAuthorizationRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteGroupAuthorizationRuleRequest) SetAuthorizationRuleId(v string) *DeleteGroupAuthorizationRuleRequest {
	s.AuthorizationRuleId = &v
	return s
}

func (s *DeleteGroupAuthorizationRuleRequest) SetClientToken(v string) *DeleteGroupAuthorizationRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteGroupAuthorizationRuleRequest) SetDryRun(v bool) *DeleteGroupAuthorizationRuleRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteGroupAuthorizationRuleRequest) SetWirelessCloudConnectorGroupId(v string) *DeleteGroupAuthorizationRuleRequest {
	s.WirelessCloudConnectorGroupId = &v
	return s
}

type DeleteGroupAuthorizationRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGroupAuthorizationRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupAuthorizationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGroupAuthorizationRuleResponseBody) SetRequestId(v string) *DeleteGroupAuthorizationRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteGroupAuthorizationRuleResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteGroupAuthorizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteGroupAuthorizationRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupAuthorizationRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteGroupAuthorizationRuleResponse) SetHeaders(v map[string]*string) *DeleteGroupAuthorizationRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteGroupAuthorizationRuleResponse) SetStatusCode(v int32) *DeleteGroupAuthorizationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGroupAuthorizationRuleResponse) SetBody(v *DeleteGroupAuthorizationRuleResponseBody) *DeleteGroupAuthorizationRuleResponse {
	s.Body = v
	return s
}

type DeleteIoTCloudConnectorBackhaulRouteRequest struct {
	ClientToken              *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                   *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	NetLinkId                *string `json:"NetLinkId,omitempty" xml:"NetLinkId,omitempty"`
	WirelessCloudConnectorId *string `json:"WirelessCloudConnectorId,omitempty" xml:"WirelessCloudConnectorId,omitempty"`
}

func (s DeleteIoTCloudConnectorBackhaulRouteRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIoTCloudConnectorBackhaulRouteRequest) GoString() string {
	return s.String()
}

func (s *DeleteIoTCloudConnectorBackhaulRouteRequest) SetClientToken(v string) *DeleteIoTCloudConnectorBackhaulRouteRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteIoTCloudConnectorBackhaulRouteRequest) SetDryRun(v bool) *DeleteIoTCloudConnectorBackhaulRouteRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteIoTCloudConnectorBackhaulRouteRequest) SetNetLinkId(v string) *DeleteIoTCloudConnectorBackhaulRouteRequest {
	s.NetLinkId = &v
	return s
}

func (s *DeleteIoTCloudConnectorBackhaulRouteRequest) SetWirelessCloudConnectorId(v string) *DeleteIoTCloudConnectorBackhaulRouteRequest {
	s.WirelessCloudConnectorId = &v
	return s
}

type DeleteIoTCloudConnectorBackhaulRouteResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIoTCloudConnectorBackhaulRouteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteIoTCloudConnectorBackhaulRouteResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIoTCloudConnectorBackhaulRouteResponseBody) SetRequestId(v string) *DeleteIoTCloudConnectorBackhaulRouteResponseBody {
	s.RequestId = &v
	return s
}

type DeleteIoTCloudConnectorBackhaulRouteResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteIoTCloudConnectorBackhaulRouteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteIoTCloudConnectorBackhaulRouteResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIoTCloudConnectorBackhaulRouteResponse) GoString() string {
	return s.String()
}

func (s *DeleteIoTCloudConnectorBackhaulRouteResponse) SetHeaders(v map[string]*string) *DeleteIoTCloudConnectorBackhaulRouteResponse {
	s.Headers = v
	return s
}

func (s *DeleteIoTCloudConnectorBackhaulRouteResponse) SetStatusCode(v int32) *DeleteIoTCloudConnectorBackhaulRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIoTCloudConnectorBackhaulRouteResponse) SetBody(v *DeleteIoTCloudConnectorBackhaulRouteResponseBody) *DeleteIoTCloudConnectorBackhaulRouteResponse {
	s.Body = v
	return s
}

type DeleteWirelessCloudConnectorRequest struct {
	ClientToken              *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                   *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	WirelessCloudConnectorId *string `json:"WirelessCloudConnectorId,omitempty" xml:"WirelessCloudConnectorId,omitempty"`
}

func (s DeleteWirelessCloudConnectorRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWirelessCloudConnectorRequest) GoString() string {
	return s.String()
}

func (s *DeleteWirelessCloudConnectorRequest) SetClientToken(v string) *DeleteWirelessCloudConnectorRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteWirelessCloudConnectorRequest) SetDryRun(v bool) *DeleteWirelessCloudConnectorRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteWirelessCloudConnectorRequest) SetWirelessCloudConnectorId(v string) *DeleteWirelessCloudConnectorRequest {
	s.WirelessCloudConnectorId = &v
	return s
}

type DeleteWirelessCloudConnectorResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWirelessCloudConnectorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWirelessCloudConnectorResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWirelessCloudConnectorResponseBody) SetRequestId(v string) *DeleteWirelessCloudConnectorResponseBody {
	s.RequestId = &v
	return s
}

type DeleteWirelessCloudConnectorResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteWirelessCloudConnectorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteWirelessCloudConnectorResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWirelessCloudConnectorResponse) GoString() string {
	return s.String()
}

func (s *DeleteWirelessCloudConnectorResponse) SetHeaders(v map[string]*string) *DeleteWirelessCloudConnectorResponse {
	s.Headers = v
	return s
}

func (s *DeleteWirelessCloudConnectorResponse) SetStatusCode(v int32) *DeleteWirelessCloudConnectorResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWirelessCloudConnectorResponse) SetBody(v *DeleteWirelessCloudConnectorResponseBody) *DeleteWirelessCloudConnectorResponse {
	s.Body = v
	return s
}

type DeleteWirelessCloudConnectorGroupRequest struct {
	ClientToken                   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                        *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	WirelessCloudConnectorGroupId *string `json:"WirelessCloudConnectorGroupId,omitempty" xml:"WirelessCloudConnectorGroupId,omitempty"`
}

func (s DeleteWirelessCloudConnectorGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWirelessCloudConnectorGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteWirelessCloudConnectorGroupRequest) SetClientToken(v string) *DeleteWirelessCloudConnectorGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteWirelessCloudConnectorGroupRequest) SetDryRun(v bool) *DeleteWirelessCloudConnectorGroupRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteWirelessCloudConnectorGroupRequest) SetWirelessCloudConnectorGroupId(v string) *DeleteWirelessCloudConnectorGroupRequest {
	s.WirelessCloudConnectorGroupId = &v
	return s
}

type DeleteWirelessCloudConnectorGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWirelessCloudConnectorGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWirelessCloudConnectorGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWirelessCloudConnectorGroupResponseBody) SetRequestId(v string) *DeleteWirelessCloudConnectorGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteWirelessCloudConnectorGroupResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteWirelessCloudConnectorGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteWirelessCloudConnectorGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWirelessCloudConnectorGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteWirelessCloudConnectorGroupResponse) SetHeaders(v map[string]*string) *DeleteWirelessCloudConnectorGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteWirelessCloudConnectorGroupResponse) SetStatusCode(v int32) *DeleteWirelessCloudConnectorGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWirelessCloudConnectorGroupResponse) SetBody(v *DeleteWirelessCloudConnectorGroupResponseBody) *DeleteWirelessCloudConnectorGroupResponse {
	s.Body = v
	return s
}

type DetachVpcFromNetLinkRequest struct {
	ClientToken              *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                   *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	NetLinkId                *string `json:"NetLinkId,omitempty" xml:"NetLinkId,omitempty"`
	WirelessCloudConnectorId *string `json:"WirelessCloudConnectorId,omitempty" xml:"WirelessCloudConnectorId,omitempty"`
}

func (s DetachVpcFromNetLinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachVpcFromNetLinkRequest) GoString() string {
	return s.String()
}

func (s *DetachVpcFromNetLinkRequest) SetClientToken(v string) *DetachVpcFromNetLinkRequest {
	s.ClientToken = &v
	return s
}

func (s *DetachVpcFromNetLinkRequest) SetDryRun(v bool) *DetachVpcFromNetLinkRequest {
	s.DryRun = &v
	return s
}

func (s *DetachVpcFromNetLinkRequest) SetNetLinkId(v string) *DetachVpcFromNetLinkRequest {
	s.NetLinkId = &v
	return s
}

func (s *DetachVpcFromNetLinkRequest) SetWirelessCloudConnectorId(v string) *DetachVpcFromNetLinkRequest {
	s.WirelessCloudConnectorId = &v
	return s
}

type DetachVpcFromNetLinkResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachVpcFromNetLinkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachVpcFromNetLinkResponseBody) GoString() string {
	return s.String()
}

func (s *DetachVpcFromNetLinkResponseBody) SetRequestId(v string) *DetachVpcFromNetLinkResponseBody {
	s.RequestId = &v
	return s
}

type DetachVpcFromNetLinkResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetachVpcFromNetLinkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetachVpcFromNetLinkResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachVpcFromNetLinkResponse) GoString() string {
	return s.String()
}

func (s *DetachVpcFromNetLinkResponse) SetHeaders(v map[string]*string) *DetachVpcFromNetLinkResponse {
	s.Headers = v
	return s
}

func (s *DetachVpcFromNetLinkResponse) SetStatusCode(v int32) *DetachVpcFromNetLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachVpcFromNetLinkResponse) SetBody(v *DetachVpcFromNetLinkResponseBody) *DetachVpcFromNetLinkResponse {
	s.Body = v
	return s
}

type FailCardsRequest struct {
	ClientToken *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	Iccids      []*string `json:"Iccids,omitempty" xml:"Iccids,omitempty" type:"Repeated"`
	RegionId    *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s FailCardsRequest) String() string {
	return tea.Prettify(s)
}

func (s FailCardsRequest) GoString() string {
	return s.String()
}

func (s *FailCardsRequest) SetClientToken(v string) *FailCardsRequest {
	s.ClientToken = &v
	return s
}

func (s *FailCardsRequest) SetDryRun(v bool) *FailCardsRequest {
	s.DryRun = &v
	return s
}

func (s *FailCardsRequest) SetIccids(v []*string) *FailCardsRequest {
	s.Iccids = v
	return s
}

func (s *FailCardsRequest) SetRegionId(v string) *FailCardsRequest {
	s.RegionId = &v
	return s
}

type FailCardsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FailCardsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FailCardsResponseBody) GoString() string {
	return s.String()
}

func (s *FailCardsResponseBody) SetRequestId(v string) *FailCardsResponseBody {
	s.RequestId = &v
	return s
}

type FailCardsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FailCardsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FailCardsResponse) String() string {
	return tea.Prettify(s)
}

func (s FailCardsResponse) GoString() string {
	return s.String()
}

func (s *FailCardsResponse) SetHeaders(v map[string]*string) *FailCardsResponse {
	s.Headers = v
	return s
}

func (s *FailCardsResponse) SetStatusCode(v int32) *FailCardsResponse {
	s.StatusCode = &v
	return s
}

func (s *FailCardsResponse) SetBody(v *FailCardsResponseBody) *FailCardsResponse {
	s.Body = v
	return s
}

type GetCardRequest struct {
	Iccid *string `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
}

func (s GetCardRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCardRequest) GoString() string {
	return s.String()
}

func (s *GetCardRequest) SetIccid(v string) *GetCardRequest {
	s.Iccid = &v
	return s
}

type GetCardResponseBody struct {
	APN                      *string `json:"APN,omitempty" xml:"APN,omitempty"`
	ActivatedTime            *string `json:"ActivatedTime,omitempty" xml:"ActivatedTime,omitempty"`
	AlarmThreshold           *int64  `json:"AlarmThreshold,omitempty" xml:"AlarmThreshold,omitempty"`
	CloudConnectorId         *string `json:"CloudConnectorId,omitempty" xml:"CloudConnectorId,omitempty"`
	Description              *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ISP                      *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	Iccid                    *string `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	Imei                     *string `json:"Imei,omitempty" xml:"Imei,omitempty"`
	Imsi                     *string `json:"Imsi,omitempty" xml:"Imsi,omitempty"`
	IpAddress                *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	LimitThreshold           *int64  `json:"LimitThreshold,omitempty" xml:"LimitThreshold,omitempty"`
	Lock                     *string `json:"Lock,omitempty" xml:"Lock,omitempty"`
	Msisdn                   *string `json:"Msisdn,omitempty" xml:"Msisdn,omitempty"`
	Name                     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NetType                  *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	OrderId                  *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId                *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SimStatus                *string `json:"SimStatus,omitempty" xml:"SimStatus,omitempty"`
	Spec                     *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	Status                   *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StopThreshold            *int64  `json:"StopThreshold,omitempty" xml:"StopThreshold,omitempty"`
	UsageDataMonth           *int32  `json:"UsageDataMonth,omitempty" xml:"UsageDataMonth,omitempty"`
	UsageDataTotal           *int64  `json:"UsageDataTotal,omitempty" xml:"UsageDataTotal,omitempty"`
	WirelessCloudConnectorId *string `json:"WirelessCloudConnectorId,omitempty" xml:"WirelessCloudConnectorId,omitempty"`
}

func (s GetCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCardResponseBody) GoString() string {
	return s.String()
}

func (s *GetCardResponseBody) SetAPN(v string) *GetCardResponseBody {
	s.APN = &v
	return s
}

func (s *GetCardResponseBody) SetActivatedTime(v string) *GetCardResponseBody {
	s.ActivatedTime = &v
	return s
}

func (s *GetCardResponseBody) SetAlarmThreshold(v int64) *GetCardResponseBody {
	s.AlarmThreshold = &v
	return s
}

func (s *GetCardResponseBody) SetCloudConnectorId(v string) *GetCardResponseBody {
	s.CloudConnectorId = &v
	return s
}

func (s *GetCardResponseBody) SetDescription(v string) *GetCardResponseBody {
	s.Description = &v
	return s
}

func (s *GetCardResponseBody) SetISP(v string) *GetCardResponseBody {
	s.ISP = &v
	return s
}

func (s *GetCardResponseBody) SetIccid(v string) *GetCardResponseBody {
	s.Iccid = &v
	return s
}

func (s *GetCardResponseBody) SetImei(v string) *GetCardResponseBody {
	s.Imei = &v
	return s
}

func (s *GetCardResponseBody) SetImsi(v string) *GetCardResponseBody {
	s.Imsi = &v
	return s
}

func (s *GetCardResponseBody) SetIpAddress(v string) *GetCardResponseBody {
	s.IpAddress = &v
	return s
}

func (s *GetCardResponseBody) SetLimitThreshold(v int64) *GetCardResponseBody {
	s.LimitThreshold = &v
	return s
}

func (s *GetCardResponseBody) SetLock(v string) *GetCardResponseBody {
	s.Lock = &v
	return s
}

func (s *GetCardResponseBody) SetMsisdn(v string) *GetCardResponseBody {
	s.Msisdn = &v
	return s
}

func (s *GetCardResponseBody) SetName(v string) *GetCardResponseBody {
	s.Name = &v
	return s
}

func (s *GetCardResponseBody) SetNetType(v string) *GetCardResponseBody {
	s.NetType = &v
	return s
}

func (s *GetCardResponseBody) SetOrderId(v string) *GetCardResponseBody {
	s.OrderId = &v
	return s
}

func (s *GetCardResponseBody) SetRequestId(v string) *GetCardResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCardResponseBody) SetSimStatus(v string) *GetCardResponseBody {
	s.SimStatus = &v
	return s
}

func (s *GetCardResponseBody) SetSpec(v string) *GetCardResponseBody {
	s.Spec = &v
	return s
}

func (s *GetCardResponseBody) SetStatus(v string) *GetCardResponseBody {
	s.Status = &v
	return s
}

func (s *GetCardResponseBody) SetStopThreshold(v int64) *GetCardResponseBody {
	s.StopThreshold = &v
	return s
}

func (s *GetCardResponseBody) SetUsageDataMonth(v int32) *GetCardResponseBody {
	s.UsageDataMonth = &v
	return s
}

func (s *GetCardResponseBody) SetUsageDataTotal(v int64) *GetCardResponseBody {
	s.UsageDataTotal = &v
	return s
}

func (s *GetCardResponseBody) SetWirelessCloudConnectorId(v string) *GetCardResponseBody {
	s.WirelessCloudConnectorId = &v
	return s
}

type GetCardResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCardResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCardResponse) GoString() string {
	return s.String()
}

func (s *GetCardResponse) SetHeaders(v map[string]*string) *GetCardResponse {
	s.Headers = v
	return s
}

func (s *GetCardResponse) SetStatusCode(v int32) *GetCardResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCardResponse) SetBody(v *GetCardResponseBody) *GetCardResponse {
	s.Body = v
	return s
}

type GetCardLockReasonRequest struct {
	Iccid *string `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
}

func (s GetCardLockReasonRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCardLockReasonRequest) GoString() string {
	return s.String()
}

func (s *GetCardLockReasonRequest) SetIccid(v string) *GetCardLockReasonRequest {
	s.Iccid = &v
	return s
}

type GetCardLockReasonResponseBody struct {
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCardLockReasonResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCardLockReasonResponseBody) GoString() string {
	return s.String()
}

func (s *GetCardLockReasonResponseBody) SetLockReason(v string) *GetCardLockReasonResponseBody {
	s.LockReason = &v
	return s
}

func (s *GetCardLockReasonResponseBody) SetRequestId(v string) *GetCardLockReasonResponseBody {
	s.RequestId = &v
	return s
}

type GetCardLockReasonResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCardLockReasonResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCardLockReasonResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCardLockReasonResponse) GoString() string {
	return s.String()
}

func (s *GetCardLockReasonResponse) SetHeaders(v map[string]*string) *GetCardLockReasonResponse {
	s.Headers = v
	return s
}

func (s *GetCardLockReasonResponse) SetStatusCode(v int32) *GetCardLockReasonResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCardLockReasonResponse) SetBody(v *GetCardLockReasonResponseBody) *GetCardLockReasonResponse {
	s.Body = v
	return s
}

type GetCreateCustomerInformationRequest struct {
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	WirelessCloudConnectorId *string `json:"WirelessCloudConnectorId,omitempty" xml:"WirelessCloudConnectorId,omitempty"`
}

func (s GetCreateCustomerInformationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCreateCustomerInformationRequest) GoString() string {
	return s.String()
}

func (s *GetCreateCustomerInformationRequest) SetRegionId(v string) *GetCreateCustomerInformationRequest {
	s.RegionId = &v
	return s
}

func (s *GetCreateCustomerInformationRequest) SetWirelessCloudConnectorId(v string) *GetCreateCustomerInformationRequest {
	s.WirelessCloudConnectorId = &v
	return s
}

type GetCreateCustomerInformationResponseBody struct {
	CanBuyCard *string `json:"CanBuyCard,omitempty" xml:"CanBuyCard,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	URL        *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s GetCreateCustomerInformationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCreateCustomerInformationResponseBody) GoString() string {
	return s.String()
}

func (s *GetCreateCustomerInformationResponseBody) SetCanBuyCard(v string) *GetCreateCustomerInformationResponseBody {
	s.CanBuyCard = &v
	return s
}

func (s *GetCreateCustomerInformationResponseBody) SetRequestId(v string) *GetCreateCustomerInformationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCreateCustomerInformationResponseBody) SetURL(v string) *GetCreateCustomerInformationResponseBody {
	s.URL = &v
	return s
}

type GetCreateCustomerInformationResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCreateCustomerInformationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCreateCustomerInformationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCreateCustomerInformationResponse) GoString() string {
	return s.String()
}

func (s *GetCreateCustomerInformationResponse) SetHeaders(v map[string]*string) *GetCreateCustomerInformationResponse {
	s.Headers = v
	return s
}

func (s *GetCreateCustomerInformationResponse) SetStatusCode(v int32) *GetCreateCustomerInformationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCreateCustomerInformationResponse) SetBody(v *GetCreateCustomerInformationResponseBody) *GetCreateCustomerInformationResponse {
	s.Body = v
	return s
}

type GetDiagnoseResultForSingleCardRequest struct {
	DiagnoseTaskId *string `json:"DiagnoseTaskId,omitempty" xml:"DiagnoseTaskId,omitempty"`
	RegionNo       *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
}

func (s GetDiagnoseResultForSingleCardRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDiagnoseResultForSingleCardRequest) GoString() string {
	return s.String()
}

func (s *GetDiagnoseResultForSingleCardRequest) SetDiagnoseTaskId(v string) *GetDiagnoseResultForSingleCardRequest {
	s.DiagnoseTaskId = &v
	return s
}

func (s *GetDiagnoseResultForSingleCardRequest) SetRegionNo(v string) *GetDiagnoseResultForSingleCardRequest {
	s.RegionNo = &v
	return s
}

type GetDiagnoseResultForSingleCardResponseBody struct {
	BeginTime                *int64                                                    `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	CardIp                   *string                                                   `json:"CardIp,omitempty" xml:"CardIp,omitempty"`
	Destination              *string                                                   `json:"Destination,omitempty" xml:"Destination,omitempty"`
	DestinationType          *string                                                   `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	DiagnoseItem             []*GetDiagnoseResultForSingleCardResponseBodyDiagnoseItem `json:"DiagnoseItem,omitempty" xml:"DiagnoseItem,omitempty" type:"Repeated"`
	EndTime                  *int64                                                    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ErrorResult              []*GetDiagnoseResultForSingleCardResponseBodyErrorResult  `json:"ErrorResult,omitempty" xml:"ErrorResult,omitempty" type:"Repeated"`
	IccId                    *string                                                   `json:"IccId,omitempty" xml:"IccId,omitempty"`
	RequestId                *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status                   *string                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	WirelessCloudConnectorId *string                                                   `json:"WirelessCloudConnectorId,omitempty" xml:"WirelessCloudConnectorId,omitempty"`
}

func (s GetDiagnoseResultForSingleCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDiagnoseResultForSingleCardResponseBody) GoString() string {
	return s.String()
}

func (s *GetDiagnoseResultForSingleCardResponseBody) SetBeginTime(v int64) *GetDiagnoseResultForSingleCardResponseBody {
	s.BeginTime = &v
	return s
}

func (s *GetDiagnoseResultForSingleCardResponseBody) SetCardIp(v string) *GetDiagnoseResultForSingleCardResponseBody {
	s.CardIp = &v
	return s
}

func (s *GetDiagnoseResultForSingleCardResponseBody) SetDestination(v string) *GetDiagnoseResultForSingleCardResponseBody {
	s.Destination = &v
	return s
}

func (s *GetDiagnoseResultForSingleCardResponseBody) SetDestinationType(v string) *GetDiagnoseResultForSingleCardResponseBody {
	s.DestinationType = &v
	return s
}

func (s *GetDiagnoseResultForSingleCardResponseBody) SetDiagnoseItem(v []*GetDiagnoseResultForSingleCardResponseBodyDiagnoseItem) *GetDiagnoseResultForSingleCardResponseBody {
	s.DiagnoseItem = v
	return s
}

func (s *GetDiagnoseResultForSingleCardResponseBody) SetEndTime(v int64) *GetDiagnoseResultForSingleCardResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetDiagnoseResultForSingleCardResponseBody) SetErrorResult(v []*GetDiagnoseResultForSingleCardResponseBodyErrorResult) *GetDiagnoseResultForSingleCardResponseBody {
	s.ErrorResult = v
	return s
}

func (s *GetDiagnoseResultForSingleCardResponseBody) SetIccId(v string) *GetDiagnoseResultForSingleCardResponseBody {
	s.IccId = &v
	return s
}

func (s *GetDiagnoseResultForSingleCardResponseBody) SetRequestId(v string) *GetDiagnoseResultForSingleCardResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDiagnoseResultForSingleCardResponseBody) SetStatus(v string) *GetDiagnoseResultForSingleCardResponseBody {
	s.Status = &v
	return s
}

func (s *GetDiagnoseResultForSingleCardResponseBody) SetWirelessCloudConnectorId(v string) *GetDiagnoseResultForSingleCardResponseBody {
	s.WirelessCloudConnectorId = &v
	return s
}

type GetDiagnoseResultForSingleCardResponseBodyDiagnoseItem struct {
	Part     *string                                                           `json:"Part,omitempty" xml:"Part,omitempty"`
	Status   *string                                                           `json:"Status,omitempty" xml:"Status,omitempty"`
	SubItems []*GetDiagnoseResultForSingleCardResponseBodyDiagnoseItemSubItems `json:"SubItems,omitempty" xml:"SubItems,omitempty" type:"Repeated"`
}

func (s GetDiagnoseResultForSingleCardResponseBodyDiagnoseItem) String() string {
	return tea.Prettify(s)
}

func (s GetDiagnoseResultForSingleCardResponseBodyDiagnoseItem) GoString() string {
	return s.String()
}

func (s *GetDiagnoseResultForSingleCardResponseBodyDiagnoseItem) SetPart(v string) *GetDiagnoseResultForSingleCardResponseBodyDiagnoseItem {
	s.Part = &v
	return s
}

func (s *GetDiagnoseResultForSingleCardResponseBodyDiagnoseItem) SetStatus(v string) *GetDiagnoseResultForSingleCardResponseBodyDiagnoseItem {
	s.Status = &v
	return s
}

func (s *GetDiagnoseResultForSingleCardResponseBodyDiagnoseItem) SetSubItems(v []*GetDiagnoseResultForSingleCardResponseBodyDiagnoseItemSubItems) *GetDiagnoseResultForSingleCardResponseBodyDiagnoseItem {
	s.SubItems = v
	return s
}

type GetDiagnoseResultForSingleCardResponseBodyDiagnoseItemSubItems struct {
	SubItem       *string `json:"SubItem,omitempty" xml:"SubItem,omitempty"`
	SubItemInfo   *string `json:"SubItemInfo,omitempty" xml:"SubItemInfo,omitempty"`
	SubItemStatus *string `json:"SubItemStatus,omitempty" xml:"SubItemStatus,omitempty"`
}

func (s GetDiagnoseResultForSingleCardResponseBodyDiagnoseItemSubItems) String() string {
	return tea.Prettify(s)
}

func (s GetDiagnoseResultForSingleCardResponseBodyDiagnoseItemSubItems) GoString() string {
	return s.String()
}

func (s *GetDiagnoseResultForSingleCardResponseBodyDiagnoseItemSubItems) SetSubItem(v string) *GetDiagnoseResultForSingleCardResponseBodyDiagnoseItemSubItems {
	s.SubItem = &v
	return s
}

func (s *GetDiagnoseResultForSingleCardResponseBodyDiagnoseItemSubItems) SetSubItemInfo(v string) *GetDiagnoseResultForSingleCardResponseBodyDiagnoseItemSubItems {
	s.SubItemInfo = &v
	return s
}

func (s *GetDiagnoseResultForSingleCardResponseBodyDiagnoseItemSubItems) SetSubItemStatus(v string) *GetDiagnoseResultForSingleCardResponseBodyDiagnoseItemSubItems {
	s.SubItemStatus = &v
	return s
}

type GetDiagnoseResultForSingleCardResponseBodyErrorResult struct {
	ErrorDesc       *string `json:"ErrorDesc,omitempty" xml:"ErrorDesc,omitempty"`
	ErrorItem       *string `json:"ErrorItem,omitempty" xml:"ErrorItem,omitempty"`
	ErrorLevel      *string `json:"ErrorLevel,omitempty" xml:"ErrorLevel,omitempty"`
	ErrorPart       *string `json:"ErrorPart,omitempty" xml:"ErrorPart,omitempty"`
	ErrorSuggestion *string `json:"ErrorSuggestion,omitempty" xml:"ErrorSuggestion,omitempty"`
}

func (s GetDiagnoseResultForSingleCardResponseBodyErrorResult) String() string {
	return tea.Prettify(s)
}

func (s GetDiagnoseResultForSingleCardResponseBodyErrorResult) GoString() string {
	return s.String()
}

func (s *GetDiagnoseResultForSingleCardResponseBodyErrorResult) SetErrorDesc(v string) *GetDiagnoseResultForSingleCardResponseBodyErrorResult {
	s.ErrorDesc = &v
	return s
}

func (s *GetDiagnoseResultForSingleCardResponseBodyErrorResult) SetErrorItem(v string) *GetDiagnoseResultForSingleCardResponseBodyErrorResult {
	s.ErrorItem = &v
	return s
}

func (s *GetDiagnoseResultForSingleCardResponseBodyErrorResult) SetErrorLevel(v string) *GetDiagnoseResultForSingleCardResponseBodyErrorResult {
	s.ErrorLevel = &v
	return s
}

func (s *GetDiagnoseResultForSingleCardResponseBodyErrorResult) SetErrorPart(v string) *GetDiagnoseResultForSingleCardResponseBodyErrorResult {
	s.ErrorPart = &v
	return s
}

func (s *GetDiagnoseResultForSingleCardResponseBodyErrorResult) SetErrorSuggestion(v string) *GetDiagnoseResultForSingleCardResponseBodyErrorResult {
	s.ErrorSuggestion = &v
	return s
}

type GetDiagnoseResultForSingleCardResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDiagnoseResultForSingleCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDiagnoseResultForSingleCardResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDiagnoseResultForSingleCardResponse) GoString() string {
	return s.String()
}

func (s *GetDiagnoseResultForSingleCardResponse) SetHeaders(v map[string]*string) *GetDiagnoseResultForSingleCardResponse {
	s.Headers = v
	return s
}

func (s *GetDiagnoseResultForSingleCardResponse) SetStatusCode(v int32) *GetDiagnoseResultForSingleCardResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDiagnoseResultForSingleCardResponse) SetBody(v *GetDiagnoseResultForSingleCardResponseBody) *GetDiagnoseResultForSingleCardResponse {
	s.Body = v
	return s
}

type GetWirelessCloudConnectorRequest struct {
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	WirelessCloudConnectorId *string `json:"WirelessCloudConnectorId,omitempty" xml:"WirelessCloudConnectorId,omitempty"`
}

func (s GetWirelessCloudConnectorRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWirelessCloudConnectorRequest) GoString() string {
	return s.String()
}

func (s *GetWirelessCloudConnectorRequest) SetRegionId(v string) *GetWirelessCloudConnectorRequest {
	s.RegionId = &v
	return s
}

func (s *GetWirelessCloudConnectorRequest) SetWirelessCloudConnectorId(v string) *GetWirelessCloudConnectorRequest {
	s.WirelessCloudConnectorId = &v
	return s
}

type GetWirelessCloudConnectorResponseBody struct {
	BusinessType                  *string                                          `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	CardCount                     *string                                          `json:"CardCount,omitempty" xml:"CardCount,omitempty"`
	CreateTime                    *string                                          `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DataPackageId                 *string                                          `json:"DataPackageId,omitempty" xml:"DataPackageId,omitempty"`
	DataPackageType               *string                                          `json:"DataPackageType,omitempty" xml:"DataPackageType,omitempty"`
	Description                   *string                                          `json:"Description,omitempty" xml:"Description,omitempty"`
	Features                      []*string                                        `json:"Features,omitempty" xml:"Features,omitempty" type:"Repeated"`
	Name                          *string                                          `json:"Name,omitempty" xml:"Name,omitempty"`
	NetLinks                      []*GetWirelessCloudConnectorResponseBodyNetLinks `json:"NetLinks,omitempty" xml:"NetLinks,omitempty" type:"Repeated"`
	RegionId                      *string                                          `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId                     *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceType                   *string                                          `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	Status                        *string                                          `json:"Status,omitempty" xml:"Status,omitempty"`
	UseCase                       *string                                          `json:"UseCase,omitempty" xml:"UseCase,omitempty"`
	WirelessCloudConnectorGroupId *string                                          `json:"WirelessCloudConnectorGroupId,omitempty" xml:"WirelessCloudConnectorGroupId,omitempty"`
	WirelessCloudConnectorId      *string                                          `json:"WirelessCloudConnectorId,omitempty" xml:"WirelessCloudConnectorId,omitempty"`
}

func (s GetWirelessCloudConnectorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWirelessCloudConnectorResponseBody) GoString() string {
	return s.String()
}

func (s *GetWirelessCloudConnectorResponseBody) SetBusinessType(v string) *GetWirelessCloudConnectorResponseBody {
	s.BusinessType = &v
	return s
}

func (s *GetWirelessCloudConnectorResponseBody) SetCardCount(v string) *GetWirelessCloudConnectorResponseBody {
	s.CardCount = &v
	return s
}

func (s *GetWirelessCloudConnectorResponseBody) SetCreateTime(v string) *GetWirelessCloudConnectorResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetWirelessCloudConnectorResponseBody) SetDataPackageId(v string) *GetWirelessCloudConnectorResponseBody {
	s.DataPackageId = &v
	return s
}

func (s *GetWirelessCloudConnectorResponseBody) SetDataPackageType(v string) *GetWirelessCloudConnectorResponseBody {
	s.DataPackageType = &v
	return s
}

func (s *GetWirelessCloudConnectorResponseBody) SetDescription(v string) *GetWirelessCloudConnectorResponseBody {
	s.Description = &v
	return s
}

func (s *GetWirelessCloudConnectorResponseBody) SetFeatures(v []*string) *GetWirelessCloudConnectorResponseBody {
	s.Features = v
	return s
}

func (s *GetWirelessCloudConnectorResponseBody) SetName(v string) *GetWirelessCloudConnectorResponseBody {
	s.Name = &v
	return s
}

func (s *GetWirelessCloudConnectorResponseBody) SetNetLinks(v []*GetWirelessCloudConnectorResponseBodyNetLinks) *GetWirelessCloudConnectorResponseBody {
	s.NetLinks = v
	return s
}

func (s *GetWirelessCloudConnectorResponseBody) SetRegionId(v string) *GetWirelessCloudConnectorResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetWirelessCloudConnectorResponseBody) SetRequestId(v string) *GetWirelessCloudConnectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWirelessCloudConnectorResponseBody) SetServiceType(v string) *GetWirelessCloudConnectorResponseBody {
	s.ServiceType = &v
	return s
}

func (s *GetWirelessCloudConnectorResponseBody) SetStatus(v string) *GetWirelessCloudConnectorResponseBody {
	s.Status = &v
	return s
}

func (s *GetWirelessCloudConnectorResponseBody) SetUseCase(v string) *GetWirelessCloudConnectorResponseBody {
	s.UseCase = &v
	return s
}

func (s *GetWirelessCloudConnectorResponseBody) SetWirelessCloudConnectorGroupId(v string) *GetWirelessCloudConnectorResponseBody {
	s.WirelessCloudConnectorGroupId = &v
	return s
}

func (s *GetWirelessCloudConnectorResponseBody) SetWirelessCloudConnectorId(v string) *GetWirelessCloudConnectorResponseBody {
	s.WirelessCloudConnectorId = &v
	return s
}

type GetWirelessCloudConnectorResponseBodyNetLinks struct {
	APN         *string `json:"APN,omitempty" xml:"APN,omitempty"`
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GrantAliUid *string `json:"GrantAliUid,omitempty" xml:"GrantAliUid,omitempty"`
	ISP         *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	// 代表创建时间的资源属性字段
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 代表资源名称的资源属性字段
	NetLinkId *string   `json:"NetLinkId,omitempty" xml:"NetLinkId,omitempty"`
	RegionId  *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status    *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	VSwitchs  []*string `json:"VSwitchs,omitempty" xml:"VSwitchs,omitempty" type:"Repeated"`
	VpcId     *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetWirelessCloudConnectorResponseBodyNetLinks) String() string {
	return tea.Prettify(s)
}

func (s GetWirelessCloudConnectorResponseBodyNetLinks) GoString() string {
	return s.String()
}

func (s *GetWirelessCloudConnectorResponseBodyNetLinks) SetAPN(v string) *GetWirelessCloudConnectorResponseBodyNetLinks {
	s.APN = &v
	return s
}

func (s *GetWirelessCloudConnectorResponseBodyNetLinks) SetCreateTime(v string) *GetWirelessCloudConnectorResponseBodyNetLinks {
	s.CreateTime = &v
	return s
}

func (s *GetWirelessCloudConnectorResponseBodyNetLinks) SetDescription(v string) *GetWirelessCloudConnectorResponseBodyNetLinks {
	s.Description = &v
	return s
}

func (s *GetWirelessCloudConnectorResponseBodyNetLinks) SetGrantAliUid(v string) *GetWirelessCloudConnectorResponseBodyNetLinks {
	s.GrantAliUid = &v
	return s
}

func (s *GetWirelessCloudConnectorResponseBodyNetLinks) SetISP(v string) *GetWirelessCloudConnectorResponseBodyNetLinks {
	s.ISP = &v
	return s
}

func (s *GetWirelessCloudConnectorResponseBodyNetLinks) SetName(v string) *GetWirelessCloudConnectorResponseBodyNetLinks {
	s.Name = &v
	return s
}

func (s *GetWirelessCloudConnectorResponseBodyNetLinks) SetNetLinkId(v string) *GetWirelessCloudConnectorResponseBodyNetLinks {
	s.NetLinkId = &v
	return s
}

func (s *GetWirelessCloudConnectorResponseBodyNetLinks) SetRegionId(v string) *GetWirelessCloudConnectorResponseBodyNetLinks {
	s.RegionId = &v
	return s
}

func (s *GetWirelessCloudConnectorResponseBodyNetLinks) SetStatus(v string) *GetWirelessCloudConnectorResponseBodyNetLinks {
	s.Status = &v
	return s
}

func (s *GetWirelessCloudConnectorResponseBodyNetLinks) SetVSwitchs(v []*string) *GetWirelessCloudConnectorResponseBodyNetLinks {
	s.VSwitchs = v
	return s
}

func (s *GetWirelessCloudConnectorResponseBodyNetLinks) SetVpcId(v string) *GetWirelessCloudConnectorResponseBodyNetLinks {
	s.VpcId = &v
	return s
}

type GetWirelessCloudConnectorResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetWirelessCloudConnectorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetWirelessCloudConnectorResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWirelessCloudConnectorResponse) GoString() string {
	return s.String()
}

func (s *GetWirelessCloudConnectorResponse) SetHeaders(v map[string]*string) *GetWirelessCloudConnectorResponse {
	s.Headers = v
	return s
}

func (s *GetWirelessCloudConnectorResponse) SetStatusCode(v int32) *GetWirelessCloudConnectorResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWirelessCloudConnectorResponse) SetBody(v *GetWirelessCloudConnectorResponseBody) *GetWirelessCloudConnectorResponse {
	s.Body = v
	return s
}

type GrantNetLinkRequest struct {
	ClientToken              *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                   *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	GrantAliUid              *int64  `json:"GrantAliUid,omitempty" xml:"GrantAliUid,omitempty"`
	NetLinkId                *string `json:"NetLinkId,omitempty" xml:"NetLinkId,omitempty"`
	WirelessCloudConnectorId *string `json:"WirelessCloudConnectorId,omitempty" xml:"WirelessCloudConnectorId,omitempty"`
}

func (s GrantNetLinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GrantNetLinkRequest) GoString() string {
	return s.String()
}

func (s *GrantNetLinkRequest) SetClientToken(v string) *GrantNetLinkRequest {
	s.ClientToken = &v
	return s
}

func (s *GrantNetLinkRequest) SetDryRun(v bool) *GrantNetLinkRequest {
	s.DryRun = &v
	return s
}

func (s *GrantNetLinkRequest) SetGrantAliUid(v int64) *GrantNetLinkRequest {
	s.GrantAliUid = &v
	return s
}

func (s *GrantNetLinkRequest) SetNetLinkId(v string) *GrantNetLinkRequest {
	s.NetLinkId = &v
	return s
}

func (s *GrantNetLinkRequest) SetWirelessCloudConnectorId(v string) *GrantNetLinkRequest {
	s.WirelessCloudConnectorId = &v
	return s
}

type GrantNetLinkResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GrantNetLinkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GrantNetLinkResponseBody) GoString() string {
	return s.String()
}

func (s *GrantNetLinkResponseBody) SetRequestId(v string) *GrantNetLinkResponseBody {
	s.RequestId = &v
	return s
}

type GrantNetLinkResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GrantNetLinkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GrantNetLinkResponse) String() string {
	return tea.Prettify(s)
}

func (s GrantNetLinkResponse) GoString() string {
	return s.String()
}

func (s *GrantNetLinkResponse) SetHeaders(v map[string]*string) *GrantNetLinkResponse {
	s.Headers = v
	return s
}

func (s *GrantNetLinkResponse) SetStatusCode(v int32) *GrantNetLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantNetLinkResponse) SetBody(v *GrantNetLinkResponseBody) *GrantNetLinkResponse {
	s.Body = v
	return s
}

type ListAPNsRequest struct {
	APN        *string `json:"APN,omitempty" xml:"APN,omitempty"`
	ISP        *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	MaxResults *int64  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListAPNsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAPNsRequest) GoString() string {
	return s.String()
}

func (s *ListAPNsRequest) SetAPN(v string) *ListAPNsRequest {
	s.APN = &v
	return s
}

func (s *ListAPNsRequest) SetISP(v string) *ListAPNsRequest {
	s.ISP = &v
	return s
}

func (s *ListAPNsRequest) SetMaxResults(v int64) *ListAPNsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAPNsRequest) SetNextToken(v string) *ListAPNsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAPNsRequest) SetRegionId(v string) *ListAPNsRequest {
	s.RegionId = &v
	return s
}

type ListAPNsResponseBody struct {
	APNs       []*ListAPNsResponseBodyAPNs `json:"APNs,omitempty" xml:"APNs,omitempty" type:"Repeated"`
	MaxResults *string                     `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAPNsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAPNsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAPNsResponseBody) SetAPNs(v []*ListAPNsResponseBodyAPNs) *ListAPNsResponseBody {
	s.APNs = v
	return s
}

func (s *ListAPNsResponseBody) SetMaxResults(v string) *ListAPNsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAPNsResponseBody) SetNextToken(v string) *ListAPNsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAPNsResponseBody) SetRequestId(v string) *ListAPNsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAPNsResponseBody) SetTotalCount(v string) *ListAPNsResponseBody {
	s.TotalCount = &v
	return s
}

type ListAPNsResponseBodyAPNs struct {
	// apn
	APN         *string `json:"APN,omitempty" xml:"APN,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 代表资源一级ID的资源属性字段
	ISP *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	// 代表创建时间的资源属性字段
	Name  *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Zones []*string `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s ListAPNsResponseBodyAPNs) String() string {
	return tea.Prettify(s)
}

func (s ListAPNsResponseBodyAPNs) GoString() string {
	return s.String()
}

func (s *ListAPNsResponseBodyAPNs) SetAPN(v string) *ListAPNsResponseBodyAPNs {
	s.APN = &v
	return s
}

func (s *ListAPNsResponseBodyAPNs) SetDescription(v string) *ListAPNsResponseBodyAPNs {
	s.Description = &v
	return s
}

func (s *ListAPNsResponseBodyAPNs) SetISP(v string) *ListAPNsResponseBodyAPNs {
	s.ISP = &v
	return s
}

func (s *ListAPNsResponseBodyAPNs) SetName(v string) *ListAPNsResponseBodyAPNs {
	s.Name = &v
	return s
}

func (s *ListAPNsResponseBodyAPNs) SetZones(v []*string) *ListAPNsResponseBodyAPNs {
	s.Zones = v
	return s
}

type ListAPNsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAPNsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAPNsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAPNsResponse) GoString() string {
	return s.String()
}

func (s *ListAPNsResponse) SetHeaders(v map[string]*string) *ListAPNsResponse {
	s.Headers = v
	return s
}

func (s *ListAPNsResponse) SetStatusCode(v int32) *ListAPNsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAPNsResponse) SetBody(v *ListAPNsResponseBody) *ListAPNsResponse {
	s.Body = v
	return s
}

type ListAuthorizationRulesRequest struct {
	AuthorizationRuleIds     []*string `json:"AuthorizationRuleIds,omitempty" xml:"AuthorizationRuleIds,omitempty" type:"Repeated"`
	Destination              *string   `json:"Destination,omitempty" xml:"Destination,omitempty"`
	DestinationPort          *string   `json:"DestinationPort,omitempty" xml:"DestinationPort,omitempty"`
	DestinationType          *string   `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	Dns                      *bool     `json:"Dns,omitempty" xml:"Dns,omitempty"`
	MaxResults               *int64    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Names                    []*string `json:"Names,omitempty" xml:"Names,omitempty" type:"Repeated"`
	NextToken                *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Policy                   *string   `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Protocol                 *string   `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	Statuses                 []*string `json:"Statuses,omitempty" xml:"Statuses,omitempty" type:"Repeated"`
	Type                     *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	WirelessCloudConnectorId *string   `json:"WirelessCloudConnectorId,omitempty" xml:"WirelessCloudConnectorId,omitempty"`
}

func (s ListAuthorizationRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAuthorizationRulesRequest) GoString() string {
	return s.String()
}

func (s *ListAuthorizationRulesRequest) SetAuthorizationRuleIds(v []*string) *ListAuthorizationRulesRequest {
	s.AuthorizationRuleIds = v
	return s
}

func (s *ListAuthorizationRulesRequest) SetDestination(v string) *ListAuthorizationRulesRequest {
	s.Destination = &v
	return s
}

func (s *ListAuthorizationRulesRequest) SetDestinationPort(v string) *ListAuthorizationRulesRequest {
	s.DestinationPort = &v
	return s
}

func (s *ListAuthorizationRulesRequest) SetDestinationType(v string) *ListAuthorizationRulesRequest {
	s.DestinationType = &v
	return s
}

func (s *ListAuthorizationRulesRequest) SetDns(v bool) *ListAuthorizationRulesRequest {
	s.Dns = &v
	return s
}

func (s *ListAuthorizationRulesRequest) SetMaxResults(v int64) *ListAuthorizationRulesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAuthorizationRulesRequest) SetNames(v []*string) *ListAuthorizationRulesRequest {
	s.Names = v
	return s
}

func (s *ListAuthorizationRulesRequest) SetNextToken(v string) *ListAuthorizationRulesRequest {
	s.NextToken = &v
	return s
}

func (s *ListAuthorizationRulesRequest) SetPolicy(v string) *ListAuthorizationRulesRequest {
	s.Policy = &v
	return s
}

func (s *ListAuthorizationRulesRequest) SetProtocol(v string) *ListAuthorizationRulesRequest {
	s.Protocol = &v
	return s
}

func (s *ListAuthorizationRulesRequest) SetStatuses(v []*string) *ListAuthorizationRulesRequest {
	s.Statuses = v
	return s
}

func (s *ListAuthorizationRulesRequest) SetType(v string) *ListAuthorizationRulesRequest {
	s.Type = &v
	return s
}

func (s *ListAuthorizationRulesRequest) SetWirelessCloudConnectorId(v string) *ListAuthorizationRulesRequest {
	s.WirelessCloudConnectorId = &v
	return s
}

type ListAuthorizationRulesResponseBody struct {
	AuthorizationRules []*ListAuthorizationRulesResponseBodyAuthorizationRules `json:"AuthorizationRules,omitempty" xml:"AuthorizationRules,omitempty" type:"Repeated"`
	MaxResults         *string                                                 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken          *string                                                 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId          *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount         *string                                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAuthorizationRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAuthorizationRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAuthorizationRulesResponseBody) SetAuthorizationRules(v []*ListAuthorizationRulesResponseBodyAuthorizationRules) *ListAuthorizationRulesResponseBody {
	s.AuthorizationRules = v
	return s
}

func (s *ListAuthorizationRulesResponseBody) SetMaxResults(v string) *ListAuthorizationRulesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAuthorizationRulesResponseBody) SetNextToken(v string) *ListAuthorizationRulesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAuthorizationRulesResponseBody) SetRequestId(v string) *ListAuthorizationRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAuthorizationRulesResponseBody) SetTotalCount(v string) *ListAuthorizationRulesResponseBody {
	s.TotalCount = &v
	return s
}

type ListAuthorizationRulesResponseBodyAuthorizationRules struct {
	// 代表资源一级ID的资源属性字段
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Destination         *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	DestinationPort     *string `json:"DestinationPort,omitempty" xml:"DestinationPort,omitempty"`
	DestinationType     *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	Dns                 *bool   `json:"Dns,omitempty" xml:"Dns,omitempty"`
	// 代表创建时间的资源属性字段
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Policy     *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Protocol   *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	SourceCidr *string `json:"SourceCidr,omitempty" xml:"SourceCidr,omitempty"`
	// 代表资源名称的资源属性字段
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListAuthorizationRulesResponseBodyAuthorizationRules) String() string {
	return tea.Prettify(s)
}

func (s ListAuthorizationRulesResponseBodyAuthorizationRules) GoString() string {
	return s.String()
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) SetAuthorizationRuleId(v string) *ListAuthorizationRulesResponseBodyAuthorizationRules {
	s.AuthorizationRuleId = &v
	return s
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) SetCreateTime(v string) *ListAuthorizationRulesResponseBodyAuthorizationRules {
	s.CreateTime = &v
	return s
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) SetDescription(v string) *ListAuthorizationRulesResponseBodyAuthorizationRules {
	s.Description = &v
	return s
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) SetDestination(v string) *ListAuthorizationRulesResponseBodyAuthorizationRules {
	s.Destination = &v
	return s
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) SetDestinationPort(v string) *ListAuthorizationRulesResponseBodyAuthorizationRules {
	s.DestinationPort = &v
	return s
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) SetDestinationType(v string) *ListAuthorizationRulesResponseBodyAuthorizationRules {
	s.DestinationType = &v
	return s
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) SetDns(v bool) *ListAuthorizationRulesResponseBodyAuthorizationRules {
	s.Dns = &v
	return s
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) SetName(v string) *ListAuthorizationRulesResponseBodyAuthorizationRules {
	s.Name = &v
	return s
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) SetPolicy(v string) *ListAuthorizationRulesResponseBodyAuthorizationRules {
	s.Policy = &v
	return s
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) SetProtocol(v string) *ListAuthorizationRulesResponseBodyAuthorizationRules {
	s.Protocol = &v
	return s
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) SetSourceCidr(v string) *ListAuthorizationRulesResponseBodyAuthorizationRules {
	s.SourceCidr = &v
	return s
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) SetStatus(v string) *ListAuthorizationRulesResponseBodyAuthorizationRules {
	s.Status = &v
	return s
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) SetType(v string) *ListAuthorizationRulesResponseBodyAuthorizationRules {
	s.Type = &v
	return s
}

type ListAuthorizationRulesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAuthorizationRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAuthorizationRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAuthorizationRulesResponse) GoString() string {
	return s.String()
}

func (s *ListAuthorizationRulesResponse) SetHeaders(v map[string]*string) *ListAuthorizationRulesResponse {
	s.Headers = v
	return s
}

func (s *ListAuthorizationRulesResponse) SetStatusCode(v int32) *ListAuthorizationRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAuthorizationRulesResponse) SetBody(v *ListAuthorizationRulesResponseBody) *ListAuthorizationRulesResponse {
	s.Body = v
	return s
}

type ListBatchOperateCardsTasksRequest struct {
	BatchOperateCardsTaskIds []*string `json:"BatchOperateCardsTaskIds,omitempty" xml:"BatchOperateCardsTaskIds,omitempty" type:"Repeated"`
	MaxResults               *int64    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Names                    []*string `json:"Names,omitempty" xml:"Names,omitempty" type:"Repeated"`
	NextToken                *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId                 *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Statuses                 []*string `json:"Statuses,omitempty" xml:"Statuses,omitempty" type:"Repeated"`
}

func (s ListBatchOperateCardsTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBatchOperateCardsTasksRequest) GoString() string {
	return s.String()
}

func (s *ListBatchOperateCardsTasksRequest) SetBatchOperateCardsTaskIds(v []*string) *ListBatchOperateCardsTasksRequest {
	s.BatchOperateCardsTaskIds = v
	return s
}

func (s *ListBatchOperateCardsTasksRequest) SetMaxResults(v int64) *ListBatchOperateCardsTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *ListBatchOperateCardsTasksRequest) SetNames(v []*string) *ListBatchOperateCardsTasksRequest {
	s.Names = v
	return s
}

func (s *ListBatchOperateCardsTasksRequest) SetNextToken(v string) *ListBatchOperateCardsTasksRequest {
	s.NextToken = &v
	return s
}

func (s *ListBatchOperateCardsTasksRequest) SetRegionId(v string) *ListBatchOperateCardsTasksRequest {
	s.RegionId = &v
	return s
}

func (s *ListBatchOperateCardsTasksRequest) SetStatuses(v []*string) *ListBatchOperateCardsTasksRequest {
	s.Statuses = v
	return s
}

type ListBatchOperateCardsTasksResponseBody struct {
	BatchOperateCardsTasks []*ListBatchOperateCardsTasksResponseBodyBatchOperateCardsTasks `json:"BatchOperateCardsTasks,omitempty" xml:"BatchOperateCardsTasks,omitempty" type:"Repeated"`
	MaxResults             *string                                                         `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken              *string                                                         `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId              *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount             *string                                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListBatchOperateCardsTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBatchOperateCardsTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListBatchOperateCardsTasksResponseBody) SetBatchOperateCardsTasks(v []*ListBatchOperateCardsTasksResponseBodyBatchOperateCardsTasks) *ListBatchOperateCardsTasksResponseBody {
	s.BatchOperateCardsTasks = v
	return s
}

func (s *ListBatchOperateCardsTasksResponseBody) SetMaxResults(v string) *ListBatchOperateCardsTasksResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListBatchOperateCardsTasksResponseBody) SetNextToken(v string) *ListBatchOperateCardsTasksResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListBatchOperateCardsTasksResponseBody) SetRequestId(v string) *ListBatchOperateCardsTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBatchOperateCardsTasksResponseBody) SetTotalCount(v string) *ListBatchOperateCardsTasksResponseBody {
	s.TotalCount = &v
	return s
}

type ListBatchOperateCardsTasksResponseBodyBatchOperateCardsTasks struct {
	// 代表资源一级ID的资源属性字段
	BatchOperateCardsTaskId  *string `json:"BatchOperateCardsTaskId,omitempty" xml:"BatchOperateCardsTaskId,omitempty"`
	CreateTime               *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description              *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EffectType               *string `json:"EffectType,omitempty" xml:"EffectType,omitempty"`
	IccidsOssFilePath        *string `json:"IccidsOssFilePath,omitempty" xml:"IccidsOssFilePath,omitempty"`
	Name                     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OperateResultOssFilePath *string `json:"OperateResultOssFilePath,omitempty" xml:"OperateResultOssFilePath,omitempty"`
	// 代表创建时间的资源属性字段
	OperateType             *string                                                                                `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	Status                  *string                                                                                `json:"Status,omitempty" xml:"Status,omitempty"`
	Threshold               *string                                                                                `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	WirelessCloudConnectors []*ListBatchOperateCardsTasksResponseBodyBatchOperateCardsTasksWirelessCloudConnectors `json:"WirelessCloudConnectors,omitempty" xml:"WirelessCloudConnectors,omitempty" type:"Repeated"`
}

func (s ListBatchOperateCardsTasksResponseBodyBatchOperateCardsTasks) String() string {
	return tea.Prettify(s)
}

func (s ListBatchOperateCardsTasksResponseBodyBatchOperateCardsTasks) GoString() string {
	return s.String()
}

func (s *ListBatchOperateCardsTasksResponseBodyBatchOperateCardsTasks) SetBatchOperateCardsTaskId(v string) *ListBatchOperateCardsTasksResponseBodyBatchOperateCardsTasks {
	s.BatchOperateCardsTaskId = &v
	return s
}

func (s *ListBatchOperateCardsTasksResponseBodyBatchOperateCardsTasks) SetCreateTime(v string) *ListBatchOperateCardsTasksResponseBodyBatchOperateCardsTasks {
	s.CreateTime = &v
	return s
}

func (s *ListBatchOperateCardsTasksResponseBodyBatchOperateCardsTasks) SetDescription(v string) *ListBatchOperateCardsTasksResponseBodyBatchOperateCardsTasks {
	s.Description = &v
	return s
}

func (s *ListBatchOperateCardsTasksResponseBodyBatchOperateCardsTasks) SetEffectType(v string) *ListBatchOperateCardsTasksResponseBodyBatchOperateCardsTasks {
	s.EffectType = &v
	return s
}

func (s *ListBatchOperateCardsTasksResponseBodyBatchOperateCardsTasks) SetIccidsOssFilePath(v string) *ListBatchOperateCardsTasksResponseBodyBatchOperateCardsTasks {
	s.IccidsOssFilePath = &v
	return s
}

func (s *ListBatchOperateCardsTasksResponseBodyBatchOperateCardsTasks) SetName(v string) *ListBatchOperateCardsTasksResponseBodyBatchOperateCardsTasks {
	s.Name = &v
	return s
}

func (s *ListBatchOperateCardsTasksResponseBodyBatchOperateCardsTasks) SetOperateResultOssFilePath(v string) *ListBatchOperateCardsTasksResponseBodyBatchOperateCardsTasks {
	s.OperateResultOssFilePath = &v
	return s
}

func (s *ListBatchOperateCardsTasksResponseBodyBatchOperateCardsTasks) SetOperateType(v string) *ListBatchOperateCardsTasksResponseBodyBatchOperateCardsTasks {
	s.OperateType = &v
	return s
}

func (s *ListBatchOperateCardsTasksResponseBodyBatchOperateCardsTasks) SetStatus(v string) *ListBatchOperateCardsTasksResponseBodyBatchOperateCardsTasks {
	s.Status = &v
	return s
}

func (s *ListBatchOperateCardsTasksResponseBodyBatchOperateCardsTasks) SetThreshold(v string) *ListBatchOperateCardsTasksResponseBodyBatchOperateCardsTasks {
	s.Threshold = &v
	return s
}

func (s *ListBatchOperateCardsTasksResponseBodyBatchOperateCardsTasks) SetWirelessCloudConnectors(v []*ListBatchOperateCardsTasksResponseBodyBatchOperateCardsTasksWirelessCloudConnectors) *ListBatchOperateCardsTasksResponseBodyBatchOperateCardsTasks {
	s.WirelessCloudConnectors = v
	return s
}

type ListBatchOperateCardsTasksResponseBodyBatchOperateCardsTasksWirelessCloudConnectors struct {
	Status                   *string `json:"Status,omitempty" xml:"Status,omitempty"`
	WirelessCloudConnectorId *string `json:"WirelessCloudConnectorId,omitempty" xml:"WirelessCloudConnectorId,omitempty"`
}

func (s ListBatchOperateCardsTasksResponseBodyBatchOperateCardsTasksWirelessCloudConnectors) String() string {
	return tea.Prettify(s)
}

func (s ListBatchOperateCardsTasksResponseBodyBatchOperateCardsTasksWirelessCloudConnectors) GoString() string {
	return s.String()
}

func (s *ListBatchOperateCardsTasksResponseBodyBatchOperateCardsTasksWirelessCloudConnectors) SetStatus(v string) *ListBatchOperateCardsTasksResponseBodyBatchOperateCardsTasksWirelessCloudConnectors {
	s.Status = &v
	return s
}

func (s *ListBatchOperateCardsTasksResponseBodyBatchOperateCardsTasksWirelessCloudConnectors) SetWirelessCloudConnectorId(v string) *ListBatchOperateCardsTasksResponseBodyBatchOperateCardsTasksWirelessCloudConnectors {
	s.WirelessCloudConnectorId = &v
	return s
}

type ListBatchOperateCardsTasksResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListBatchOperateCardsTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListBatchOperateCardsTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBatchOperateCardsTasksResponse) GoString() string {
	return s.String()
}

func (s *ListBatchOperateCardsTasksResponse) SetHeaders(v map[string]*string) *ListBatchOperateCardsTasksResponse {
	s.Headers = v
	return s
}

func (s *ListBatchOperateCardsTasksResponse) SetStatusCode(v int32) *ListBatchOperateCardsTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBatchOperateCardsTasksResponse) SetBody(v *ListBatchOperateCardsTasksResponseBody) *ListBatchOperateCardsTasksResponse {
	s.Body = v
	return s
}

type ListCardUsagesRequest struct {
	Iccids                   []*string `json:"Iccids,omitempty" xml:"Iccids,omitempty" type:"Repeated"`
	WirelessCloudConnectorId *string   `json:"WirelessCloudConnectorId,omitempty" xml:"WirelessCloudConnectorId,omitempty"`
}

func (s ListCardUsagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCardUsagesRequest) GoString() string {
	return s.String()
}

func (s *ListCardUsagesRequest) SetIccids(v []*string) *ListCardUsagesRequest {
	s.Iccids = v
	return s
}

func (s *ListCardUsagesRequest) SetWirelessCloudConnectorId(v string) *ListCardUsagesRequest {
	s.WirelessCloudConnectorId = &v
	return s
}

type ListCardUsagesResponseBody struct {
	Cards      []*ListCardUsagesResponseBodyCards `json:"Cards,omitempty" xml:"Cards,omitempty" type:"Repeated"`
	RequestId  *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *string                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCardUsagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCardUsagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCardUsagesResponseBody) SetCards(v []*ListCardUsagesResponseBodyCards) *ListCardUsagesResponseBody {
	s.Cards = v
	return s
}

func (s *ListCardUsagesResponseBody) SetRequestId(v string) *ListCardUsagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCardUsagesResponseBody) SetTotalCount(v string) *ListCardUsagesResponseBody {
	s.TotalCount = &v
	return s
}

type ListCardUsagesResponseBodyCards struct {
	// 代表资源一级ID的资源属性字段
	Iccid          *string `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	UsageDataMonth *int64  `json:"UsageDataMonth,omitempty" xml:"UsageDataMonth,omitempty"`
}

func (s ListCardUsagesResponseBodyCards) String() string {
	return tea.Prettify(s)
}

func (s ListCardUsagesResponseBodyCards) GoString() string {
	return s.String()
}

func (s *ListCardUsagesResponseBodyCards) SetIccid(v string) *ListCardUsagesResponseBodyCards {
	s.Iccid = &v
	return s
}

func (s *ListCardUsagesResponseBodyCards) SetUsageDataMonth(v int64) *ListCardUsagesResponseBodyCards {
	s.UsageDataMonth = &v
	return s
}

type ListCardUsagesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCardUsagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCardUsagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCardUsagesResponse) GoString() string {
	return s.String()
}

func (s *ListCardUsagesResponse) SetHeaders(v map[string]*string) *ListCardUsagesResponse {
	s.Headers = v
	return s
}

func (s *ListCardUsagesResponse) SetStatusCode(v int32) *ListCardUsagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCardUsagesResponse) SetBody(v *ListCardUsagesResponseBody) *ListCardUsagesResponse {
	s.Body = v
	return s
}

type ListCardsRequest struct {
	Apn                      *string   `json:"Apn,omitempty" xml:"Apn,omitempty"`
	Iccid                    *string   `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	Iccids                   []*string `json:"Iccids,omitempty" xml:"Iccids,omitempty" type:"Repeated"`
	IpAddress                *string   `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	Lock                     *bool     `json:"Lock,omitempty" xml:"Lock,omitempty"`
	MaxResults               *int64    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Msisdn                   *string   `json:"Msisdn,omitempty" xml:"Msisdn,omitempty"`
	NetLinkId                *string   `json:"NetLinkId,omitempty" xml:"NetLinkId,omitempty"`
	NextToken                *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Online                   *bool     `json:"Online,omitempty" xml:"Online,omitempty"`
	Statuses                 []*string `json:"Statuses,omitempty" xml:"Statuses,omitempty" type:"Repeated"`
	WirelessCloudConnectorId *string   `json:"WirelessCloudConnectorId,omitempty" xml:"WirelessCloudConnectorId,omitempty"`
}

func (s ListCardsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCardsRequest) GoString() string {
	return s.String()
}

func (s *ListCardsRequest) SetApn(v string) *ListCardsRequest {
	s.Apn = &v
	return s
}

func (s *ListCardsRequest) SetIccid(v string) *ListCardsRequest {
	s.Iccid = &v
	return s
}

func (s *ListCardsRequest) SetIccids(v []*string) *ListCardsRequest {
	s.Iccids = v
	return s
}

func (s *ListCardsRequest) SetIpAddress(v string) *ListCardsRequest {
	s.IpAddress = &v
	return s
}

func (s *ListCardsRequest) SetLock(v bool) *ListCardsRequest {
	s.Lock = &v
	return s
}

func (s *ListCardsRequest) SetMaxResults(v int64) *ListCardsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListCardsRequest) SetMsisdn(v string) *ListCardsRequest {
	s.Msisdn = &v
	return s
}

func (s *ListCardsRequest) SetNetLinkId(v string) *ListCardsRequest {
	s.NetLinkId = &v
	return s
}

func (s *ListCardsRequest) SetNextToken(v string) *ListCardsRequest {
	s.NextToken = &v
	return s
}

func (s *ListCardsRequest) SetOnline(v bool) *ListCardsRequest {
	s.Online = &v
	return s
}

func (s *ListCardsRequest) SetStatuses(v []*string) *ListCardsRequest {
	s.Statuses = v
	return s
}

func (s *ListCardsRequest) SetWirelessCloudConnectorId(v string) *ListCardsRequest {
	s.WirelessCloudConnectorId = &v
	return s
}

type ListCardsResponseBody struct {
	Cards      []*ListCardsResponseBodyCards `json:"Cards,omitempty" xml:"Cards,omitempty" type:"Repeated"`
	MaxResults *string                       `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *string                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCardsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCardsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCardsResponseBody) SetCards(v []*ListCardsResponseBodyCards) *ListCardsResponseBody {
	s.Cards = v
	return s
}

func (s *ListCardsResponseBody) SetMaxResults(v string) *ListCardsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListCardsResponseBody) SetNextToken(v string) *ListCardsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListCardsResponseBody) SetRequestId(v string) *ListCardsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCardsResponseBody) SetTotalCount(v string) *ListCardsResponseBody {
	s.TotalCount = &v
	return s
}

type ListCardsResponseBodyCards struct {
	// 代表创建时间的资源属性字段
	APN            *string `json:"APN,omitempty" xml:"APN,omitempty"`
	ActivatedTime  *string `json:"ActivatedTime,omitempty" xml:"ActivatedTime,omitempty"`
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ISP            *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	// 代表资源一级ID的资源属性字段
	Iccid     *string `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	Imei      *string `json:"Imei,omitempty" xml:"Imei,omitempty"`
	Imsi      *string `json:"Imsi,omitempty" xml:"Imsi,omitempty"`
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	Lock      *bool   `json:"Lock,omitempty" xml:"Lock,omitempty"`
	Msisdn    *string `json:"Msisdn,omitempty" xml:"Msisdn,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 代表资源名称的资源属性字段
	NetType        *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	OrderId        *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	Spec           *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UsageDataMonth *int64  `json:"UsageDataMonth,omitempty" xml:"UsageDataMonth,omitempty"`
	UsageDataTotal *string `json:"UsageDataTotal,omitempty" xml:"UsageDataTotal,omitempty"`
}

func (s ListCardsResponseBodyCards) String() string {
	return tea.Prettify(s)
}

func (s ListCardsResponseBodyCards) GoString() string {
	return s.String()
}

func (s *ListCardsResponseBodyCards) SetAPN(v string) *ListCardsResponseBodyCards {
	s.APN = &v
	return s
}

func (s *ListCardsResponseBodyCards) SetActivatedTime(v string) *ListCardsResponseBodyCards {
	s.ActivatedTime = &v
	return s
}

func (s *ListCardsResponseBodyCards) SetBusinessStatus(v string) *ListCardsResponseBodyCards {
	s.BusinessStatus = &v
	return s
}

func (s *ListCardsResponseBodyCards) SetDescription(v string) *ListCardsResponseBodyCards {
	s.Description = &v
	return s
}

func (s *ListCardsResponseBodyCards) SetISP(v string) *ListCardsResponseBodyCards {
	s.ISP = &v
	return s
}

func (s *ListCardsResponseBodyCards) SetIccid(v string) *ListCardsResponseBodyCards {
	s.Iccid = &v
	return s
}

func (s *ListCardsResponseBodyCards) SetImei(v string) *ListCardsResponseBodyCards {
	s.Imei = &v
	return s
}

func (s *ListCardsResponseBodyCards) SetImsi(v string) *ListCardsResponseBodyCards {
	s.Imsi = &v
	return s
}

func (s *ListCardsResponseBodyCards) SetIpAddress(v string) *ListCardsResponseBodyCards {
	s.IpAddress = &v
	return s
}

func (s *ListCardsResponseBodyCards) SetLock(v bool) *ListCardsResponseBodyCards {
	s.Lock = &v
	return s
}

func (s *ListCardsResponseBodyCards) SetMsisdn(v string) *ListCardsResponseBodyCards {
	s.Msisdn = &v
	return s
}

func (s *ListCardsResponseBodyCards) SetName(v string) *ListCardsResponseBodyCards {
	s.Name = &v
	return s
}

func (s *ListCardsResponseBodyCards) SetNetType(v string) *ListCardsResponseBodyCards {
	s.NetType = &v
	return s
}

func (s *ListCardsResponseBodyCards) SetOrderId(v string) *ListCardsResponseBodyCards {
	s.OrderId = &v
	return s
}

func (s *ListCardsResponseBodyCards) SetSpec(v string) *ListCardsResponseBodyCards {
	s.Spec = &v
	return s
}

func (s *ListCardsResponseBodyCards) SetStatus(v string) *ListCardsResponseBodyCards {
	s.Status = &v
	return s
}

func (s *ListCardsResponseBodyCards) SetUsageDataMonth(v int64) *ListCardsResponseBodyCards {
	s.UsageDataMonth = &v
	return s
}

func (s *ListCardsResponseBodyCards) SetUsageDataTotal(v string) *ListCardsResponseBodyCards {
	s.UsageDataTotal = &v
	return s
}

type ListCardsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCardsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCardsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCardsResponse) GoString() string {
	return s.String()
}

func (s *ListCardsResponse) SetHeaders(v map[string]*string) *ListCardsResponse {
	s.Headers = v
	return s
}

func (s *ListCardsResponse) SetStatusCode(v int32) *ListCardsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCardsResponse) SetBody(v *ListCardsResponseBody) *ListCardsResponse {
	s.Body = v
	return s
}

type ListDataPackagesRequest struct {
	DataPackageIds           []*string `json:"DataPackageIds,omitempty" xml:"DataPackageIds,omitempty" type:"Repeated"`
	MaxResults               *int64    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Names                    []*string `json:"Names,omitempty" xml:"Names,omitempty" type:"Repeated"`
	NextToken                *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Statuses                 []*string `json:"Statuses,omitempty" xml:"Statuses,omitempty" type:"Repeated"`
	WirelessCloudConnectorId *string   `json:"WirelessCloudConnectorId,omitempty" xml:"WirelessCloudConnectorId,omitempty"`
}

func (s ListDataPackagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDataPackagesRequest) GoString() string {
	return s.String()
}

func (s *ListDataPackagesRequest) SetDataPackageIds(v []*string) *ListDataPackagesRequest {
	s.DataPackageIds = v
	return s
}

func (s *ListDataPackagesRequest) SetMaxResults(v int64) *ListDataPackagesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDataPackagesRequest) SetNames(v []*string) *ListDataPackagesRequest {
	s.Names = v
	return s
}

func (s *ListDataPackagesRequest) SetNextToken(v string) *ListDataPackagesRequest {
	s.NextToken = &v
	return s
}

func (s *ListDataPackagesRequest) SetStatuses(v []*string) *ListDataPackagesRequest {
	s.Statuses = v
	return s
}

func (s *ListDataPackagesRequest) SetWirelessCloudConnectorId(v string) *ListDataPackagesRequest {
	s.WirelessCloudConnectorId = &v
	return s
}

type ListDataPackagesResponseBody struct {
	DataPackages []*ListDataPackagesResponseBodyDataPackages `json:"DataPackages,omitempty" xml:"DataPackages,omitempty" type:"Repeated"`
	MaxResults   *string                                     `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string                                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount   *string                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataPackagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDataPackagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataPackagesResponseBody) SetDataPackages(v []*ListDataPackagesResponseBodyDataPackages) *ListDataPackagesResponseBody {
	s.DataPackages = v
	return s
}

func (s *ListDataPackagesResponseBody) SetMaxResults(v string) *ListDataPackagesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDataPackagesResponseBody) SetNextToken(v string) *ListDataPackagesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDataPackagesResponseBody) SetRequestId(v string) *ListDataPackagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataPackagesResponseBody) SetTotalCount(v string) *ListDataPackagesResponseBody {
	s.TotalCount = &v
	return s
}

type ListDataPackagesResponseBodyDataPackages struct {
	CardCount     *string `json:"CardCount,omitempty" xml:"CardCount,omitempty"`
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DataPackageId *string `json:"DataPackageId,omitempty" xml:"DataPackageId,omitempty"`
	ExpiredTime   *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	ISP           *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	// 代表创建时间的资源属性字段
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// 代表资源名称的资源属性字段
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDataPackagesResponseBodyDataPackages) String() string {
	return tea.Prettify(s)
}

func (s ListDataPackagesResponseBodyDataPackages) GoString() string {
	return s.String()
}

func (s *ListDataPackagesResponseBodyDataPackages) SetCardCount(v string) *ListDataPackagesResponseBodyDataPackages {
	s.CardCount = &v
	return s
}

func (s *ListDataPackagesResponseBodyDataPackages) SetCreateTime(v string) *ListDataPackagesResponseBodyDataPackages {
	s.CreateTime = &v
	return s
}

func (s *ListDataPackagesResponseBodyDataPackages) SetDataPackageId(v string) *ListDataPackagesResponseBodyDataPackages {
	s.DataPackageId = &v
	return s
}

func (s *ListDataPackagesResponseBodyDataPackages) SetExpiredTime(v string) *ListDataPackagesResponseBodyDataPackages {
	s.ExpiredTime = &v
	return s
}

func (s *ListDataPackagesResponseBodyDataPackages) SetISP(v string) *ListDataPackagesResponseBodyDataPackages {
	s.ISP = &v
	return s
}

func (s *ListDataPackagesResponseBodyDataPackages) SetName(v string) *ListDataPackagesResponseBodyDataPackages {
	s.Name = &v
	return s
}

func (s *ListDataPackagesResponseBodyDataPackages) SetSize(v string) *ListDataPackagesResponseBodyDataPackages {
	s.Size = &v
	return s
}

func (s *ListDataPackagesResponseBodyDataPackages) SetStatus(v string) *ListDataPackagesResponseBodyDataPackages {
	s.Status = &v
	return s
}

type ListDataPackagesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDataPackagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDataPackagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDataPackagesResponse) GoString() string {
	return s.String()
}

func (s *ListDataPackagesResponse) SetHeaders(v map[string]*string) *ListDataPackagesResponse {
	s.Headers = v
	return s
}

func (s *ListDataPackagesResponse) SetStatusCode(v int32) *ListDataPackagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataPackagesResponse) SetBody(v *ListDataPackagesResponseBody) *ListDataPackagesResponse {
	s.Body = v
	return s
}

type ListDiagnoseInfoForSingleCardRequest struct {
	MaxResults               *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken                *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionNo                 *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	Source                   *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceType               *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Status                   *string `json:"Status,omitempty" xml:"Status,omitempty"`
	WirelessCloudConnectorId *string `json:"WirelessCloudConnectorId,omitempty" xml:"WirelessCloudConnectorId,omitempty"`
}

func (s ListDiagnoseInfoForSingleCardRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDiagnoseInfoForSingleCardRequest) GoString() string {
	return s.String()
}

func (s *ListDiagnoseInfoForSingleCardRequest) SetMaxResults(v int32) *ListDiagnoseInfoForSingleCardRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDiagnoseInfoForSingleCardRequest) SetNextToken(v string) *ListDiagnoseInfoForSingleCardRequest {
	s.NextToken = &v
	return s
}

func (s *ListDiagnoseInfoForSingleCardRequest) SetRegionNo(v string) *ListDiagnoseInfoForSingleCardRequest {
	s.RegionNo = &v
	return s
}

func (s *ListDiagnoseInfoForSingleCardRequest) SetSource(v string) *ListDiagnoseInfoForSingleCardRequest {
	s.Source = &v
	return s
}

func (s *ListDiagnoseInfoForSingleCardRequest) SetSourceType(v string) *ListDiagnoseInfoForSingleCardRequest {
	s.SourceType = &v
	return s
}

func (s *ListDiagnoseInfoForSingleCardRequest) SetStatus(v string) *ListDiagnoseInfoForSingleCardRequest {
	s.Status = &v
	return s
}

func (s *ListDiagnoseInfoForSingleCardRequest) SetWirelessCloudConnectorId(v string) *ListDiagnoseInfoForSingleCardRequest {
	s.WirelessCloudConnectorId = &v
	return s
}

type ListDiagnoseInfoForSingleCardResponseBody struct {
	DiagnoseInfo []*ListDiagnoseInfoForSingleCardResponseBodyDiagnoseInfo `json:"DiagnoseInfo,omitempty" xml:"DiagnoseInfo,omitempty" type:"Repeated"`
	MaxResults   *int64                                                   `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string                                                  `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount   *int64                                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDiagnoseInfoForSingleCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDiagnoseInfoForSingleCardResponseBody) GoString() string {
	return s.String()
}

func (s *ListDiagnoseInfoForSingleCardResponseBody) SetDiagnoseInfo(v []*ListDiagnoseInfoForSingleCardResponseBodyDiagnoseInfo) *ListDiagnoseInfoForSingleCardResponseBody {
	s.DiagnoseInfo = v
	return s
}

func (s *ListDiagnoseInfoForSingleCardResponseBody) SetMaxResults(v int64) *ListDiagnoseInfoForSingleCardResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDiagnoseInfoForSingleCardResponseBody) SetNextToken(v string) *ListDiagnoseInfoForSingleCardResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDiagnoseInfoForSingleCardResponseBody) SetRequestId(v string) *ListDiagnoseInfoForSingleCardResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDiagnoseInfoForSingleCardResponseBody) SetTotalCount(v int64) *ListDiagnoseInfoForSingleCardResponseBody {
	s.TotalCount = &v
	return s
}

type ListDiagnoseInfoForSingleCardResponseBodyDiagnoseInfo struct {
	BeginTime                *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	CardIp                   *string `json:"CardIp,omitempty" xml:"CardIp,omitempty"`
	Destination              *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	DestinationType          *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	DiagnoseTaskId           *string `json:"DiagnoseTaskId,omitempty" xml:"DiagnoseTaskId,omitempty"`
	DiagnoseTime             *int64  `json:"DiagnoseTime,omitempty" xml:"DiagnoseTime,omitempty"`
	EndTime                  *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IccId                    *string `json:"IccId,omitempty" xml:"IccId,omitempty"`
	Source                   *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceType               *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Status                   *string `json:"Status,omitempty" xml:"Status,omitempty"`
	WirelessCloudConnectorId *string `json:"WirelessCloudConnectorId,omitempty" xml:"WirelessCloudConnectorId,omitempty"`
}

func (s ListDiagnoseInfoForSingleCardResponseBodyDiagnoseInfo) String() string {
	return tea.Prettify(s)
}

func (s ListDiagnoseInfoForSingleCardResponseBodyDiagnoseInfo) GoString() string {
	return s.String()
}

func (s *ListDiagnoseInfoForSingleCardResponseBodyDiagnoseInfo) SetBeginTime(v int64) *ListDiagnoseInfoForSingleCardResponseBodyDiagnoseInfo {
	s.BeginTime = &v
	return s
}

func (s *ListDiagnoseInfoForSingleCardResponseBodyDiagnoseInfo) SetCardIp(v string) *ListDiagnoseInfoForSingleCardResponseBodyDiagnoseInfo {
	s.CardIp = &v
	return s
}

func (s *ListDiagnoseInfoForSingleCardResponseBodyDiagnoseInfo) SetDestination(v string) *ListDiagnoseInfoForSingleCardResponseBodyDiagnoseInfo {
	s.Destination = &v
	return s
}

func (s *ListDiagnoseInfoForSingleCardResponseBodyDiagnoseInfo) SetDestinationType(v string) *ListDiagnoseInfoForSingleCardResponseBodyDiagnoseInfo {
	s.DestinationType = &v
	return s
}

func (s *ListDiagnoseInfoForSingleCardResponseBodyDiagnoseInfo) SetDiagnoseTaskId(v string) *ListDiagnoseInfoForSingleCardResponseBodyDiagnoseInfo {
	s.DiagnoseTaskId = &v
	return s
}

func (s *ListDiagnoseInfoForSingleCardResponseBodyDiagnoseInfo) SetDiagnoseTime(v int64) *ListDiagnoseInfoForSingleCardResponseBodyDiagnoseInfo {
	s.DiagnoseTime = &v
	return s
}

func (s *ListDiagnoseInfoForSingleCardResponseBodyDiagnoseInfo) SetEndTime(v int64) *ListDiagnoseInfoForSingleCardResponseBodyDiagnoseInfo {
	s.EndTime = &v
	return s
}

func (s *ListDiagnoseInfoForSingleCardResponseBodyDiagnoseInfo) SetIccId(v string) *ListDiagnoseInfoForSingleCardResponseBodyDiagnoseInfo {
	s.IccId = &v
	return s
}

func (s *ListDiagnoseInfoForSingleCardResponseBodyDiagnoseInfo) SetSource(v string) *ListDiagnoseInfoForSingleCardResponseBodyDiagnoseInfo {
	s.Source = &v
	return s
}

func (s *ListDiagnoseInfoForSingleCardResponseBodyDiagnoseInfo) SetSourceType(v string) *ListDiagnoseInfoForSingleCardResponseBodyDiagnoseInfo {
	s.SourceType = &v
	return s
}

func (s *ListDiagnoseInfoForSingleCardResponseBodyDiagnoseInfo) SetStatus(v string) *ListDiagnoseInfoForSingleCardResponseBodyDiagnoseInfo {
	s.Status = &v
	return s
}

func (s *ListDiagnoseInfoForSingleCardResponseBodyDiagnoseInfo) SetWirelessCloudConnectorId(v string) *ListDiagnoseInfoForSingleCardResponseBodyDiagnoseInfo {
	s.WirelessCloudConnectorId = &v
	return s
}

type ListDiagnoseInfoForSingleCardResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDiagnoseInfoForSingleCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDiagnoseInfoForSingleCardResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDiagnoseInfoForSingleCardResponse) GoString() string {
	return s.String()
}

func (s *ListDiagnoseInfoForSingleCardResponse) SetHeaders(v map[string]*string) *ListDiagnoseInfoForSingleCardResponse {
	s.Headers = v
	return s
}

func (s *ListDiagnoseInfoForSingleCardResponse) SetStatusCode(v int32) *ListDiagnoseInfoForSingleCardResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDiagnoseInfoForSingleCardResponse) SetBody(v *ListDiagnoseInfoForSingleCardResponseBody) *ListDiagnoseInfoForSingleCardResponse {
	s.Body = v
	return s
}

type ListGroupAuthorizationRulesRequest struct {
	AuthorizationRuleIds          []*string `json:"AuthorizationRuleIds,omitempty" xml:"AuthorizationRuleIds,omitempty" type:"Repeated"`
	Destination                   *string   `json:"Destination,omitempty" xml:"Destination,omitempty"`
	DestinationPort               *string   `json:"DestinationPort,omitempty" xml:"DestinationPort,omitempty"`
	DestinationType               *string   `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	Dns                           *bool     `json:"Dns,omitempty" xml:"Dns,omitempty"`
	MaxResults                    *int64    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Names                         []*string `json:"Names,omitempty" xml:"Names,omitempty" type:"Repeated"`
	NextToken                     *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Policy                        *string   `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Protocol                      *string   `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	Statuses                      []*string `json:"Statuses,omitempty" xml:"Statuses,omitempty" type:"Repeated"`
	Type                          *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	WirelessCloudConnectorGroupId *string   `json:"WirelessCloudConnectorGroupId,omitempty" xml:"WirelessCloudConnectorGroupId,omitempty"`
}

func (s ListGroupAuthorizationRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGroupAuthorizationRulesRequest) GoString() string {
	return s.String()
}

func (s *ListGroupAuthorizationRulesRequest) SetAuthorizationRuleIds(v []*string) *ListGroupAuthorizationRulesRequest {
	s.AuthorizationRuleIds = v
	return s
}

func (s *ListGroupAuthorizationRulesRequest) SetDestination(v string) *ListGroupAuthorizationRulesRequest {
	s.Destination = &v
	return s
}

func (s *ListGroupAuthorizationRulesRequest) SetDestinationPort(v string) *ListGroupAuthorizationRulesRequest {
	s.DestinationPort = &v
	return s
}

func (s *ListGroupAuthorizationRulesRequest) SetDestinationType(v string) *ListGroupAuthorizationRulesRequest {
	s.DestinationType = &v
	return s
}

func (s *ListGroupAuthorizationRulesRequest) SetDns(v bool) *ListGroupAuthorizationRulesRequest {
	s.Dns = &v
	return s
}

func (s *ListGroupAuthorizationRulesRequest) SetMaxResults(v int64) *ListGroupAuthorizationRulesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListGroupAuthorizationRulesRequest) SetNames(v []*string) *ListGroupAuthorizationRulesRequest {
	s.Names = v
	return s
}

func (s *ListGroupAuthorizationRulesRequest) SetNextToken(v string) *ListGroupAuthorizationRulesRequest {
	s.NextToken = &v
	return s
}

func (s *ListGroupAuthorizationRulesRequest) SetPolicy(v string) *ListGroupAuthorizationRulesRequest {
	s.Policy = &v
	return s
}

func (s *ListGroupAuthorizationRulesRequest) SetProtocol(v string) *ListGroupAuthorizationRulesRequest {
	s.Protocol = &v
	return s
}

func (s *ListGroupAuthorizationRulesRequest) SetStatuses(v []*string) *ListGroupAuthorizationRulesRequest {
	s.Statuses = v
	return s
}

func (s *ListGroupAuthorizationRulesRequest) SetType(v string) *ListGroupAuthorizationRulesRequest {
	s.Type = &v
	return s
}

func (s *ListGroupAuthorizationRulesRequest) SetWirelessCloudConnectorGroupId(v string) *ListGroupAuthorizationRulesRequest {
	s.WirelessCloudConnectorGroupId = &v
	return s
}

type ListGroupAuthorizationRulesResponseBody struct {
	GroupAuthorizationRules []*ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules `json:"GroupAuthorizationRules,omitempty" xml:"GroupAuthorizationRules,omitempty" type:"Repeated"`
	MaxResults              *string                                                           `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken               *string                                                           `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId               *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount              *string                                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListGroupAuthorizationRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGroupAuthorizationRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupAuthorizationRulesResponseBody) SetGroupAuthorizationRules(v []*ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules) *ListGroupAuthorizationRulesResponseBody {
	s.GroupAuthorizationRules = v
	return s
}

func (s *ListGroupAuthorizationRulesResponseBody) SetMaxResults(v string) *ListGroupAuthorizationRulesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListGroupAuthorizationRulesResponseBody) SetNextToken(v string) *ListGroupAuthorizationRulesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListGroupAuthorizationRulesResponseBody) SetRequestId(v string) *ListGroupAuthorizationRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGroupAuthorizationRulesResponseBody) SetTotalCount(v string) *ListGroupAuthorizationRulesResponseBody {
	s.TotalCount = &v
	return s
}

type ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules struct {
	// 代表资源一级ID的资源属性字段
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Destination         *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	DestinationPort     *string `json:"DestinationPort,omitempty" xml:"DestinationPort,omitempty"`
	DestinationType     *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	Dns                 *bool   `json:"Dns,omitempty" xml:"Dns,omitempty"`
	// 代表创建时间的资源属性字段
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Policy     *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Protocol   *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	SourceCidr *string `json:"SourceCidr,omitempty" xml:"SourceCidr,omitempty"`
	// 代表资源名称的资源属性字段
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules) String() string {
	return tea.Prettify(s)
}

func (s ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules) GoString() string {
	return s.String()
}

func (s *ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules) SetAuthorizationRuleId(v string) *ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules {
	s.AuthorizationRuleId = &v
	return s
}

func (s *ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules) SetCreateTime(v string) *ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules {
	s.CreateTime = &v
	return s
}

func (s *ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules) SetDescription(v string) *ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules {
	s.Description = &v
	return s
}

func (s *ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules) SetDestination(v string) *ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules {
	s.Destination = &v
	return s
}

func (s *ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules) SetDestinationPort(v string) *ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules {
	s.DestinationPort = &v
	return s
}

func (s *ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules) SetDestinationType(v string) *ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules {
	s.DestinationType = &v
	return s
}

func (s *ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules) SetDns(v bool) *ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules {
	s.Dns = &v
	return s
}

func (s *ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules) SetName(v string) *ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules {
	s.Name = &v
	return s
}

func (s *ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules) SetPolicy(v string) *ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules {
	s.Policy = &v
	return s
}

func (s *ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules) SetProtocol(v string) *ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules {
	s.Protocol = &v
	return s
}

func (s *ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules) SetSourceCidr(v string) *ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules {
	s.SourceCidr = &v
	return s
}

func (s *ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules) SetStatus(v string) *ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules {
	s.Status = &v
	return s
}

func (s *ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules) SetType(v string) *ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules {
	s.Type = &v
	return s
}

type ListGroupAuthorizationRulesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListGroupAuthorizationRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListGroupAuthorizationRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGroupAuthorizationRulesResponse) GoString() string {
	return s.String()
}

func (s *ListGroupAuthorizationRulesResponse) SetHeaders(v map[string]*string) *ListGroupAuthorizationRulesResponse {
	s.Headers = v
	return s
}

func (s *ListGroupAuthorizationRulesResponse) SetStatusCode(v int32) *ListGroupAuthorizationRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGroupAuthorizationRulesResponse) SetBody(v *ListGroupAuthorizationRulesResponseBody) *ListGroupAuthorizationRulesResponse {
	s.Body = v
	return s
}

type ListIoTCloudConnectorBackhaulRouteRequest struct {
	NetLinkId                *string `json:"NetLinkId,omitempty" xml:"NetLinkId,omitempty"`
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	WirelessCloudConnectorId *string `json:"WirelessCloudConnectorId,omitempty" xml:"WirelessCloudConnectorId,omitempty"`
}

func (s ListIoTCloudConnectorBackhaulRouteRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIoTCloudConnectorBackhaulRouteRequest) GoString() string {
	return s.String()
}

func (s *ListIoTCloudConnectorBackhaulRouteRequest) SetNetLinkId(v string) *ListIoTCloudConnectorBackhaulRouteRequest {
	s.NetLinkId = &v
	return s
}

func (s *ListIoTCloudConnectorBackhaulRouteRequest) SetRegionId(v string) *ListIoTCloudConnectorBackhaulRouteRequest {
	s.RegionId = &v
	return s
}

func (s *ListIoTCloudConnectorBackhaulRouteRequest) SetWirelessCloudConnectorId(v string) *ListIoTCloudConnectorBackhaulRouteRequest {
	s.WirelessCloudConnectorId = &v
	return s
}

type ListIoTCloudConnectorBackhaulRouteResponseBody struct {
	NetLinkId *string                                                 `json:"NetLinkId,omitempty" xml:"NetLinkId,omitempty"`
	RequestId *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Routes    []*ListIoTCloudConnectorBackhaulRouteResponseBodyRoutes `json:"Routes,omitempty" xml:"Routes,omitempty" type:"Repeated"`
}

func (s ListIoTCloudConnectorBackhaulRouteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIoTCloudConnectorBackhaulRouteResponseBody) GoString() string {
	return s.String()
}

func (s *ListIoTCloudConnectorBackhaulRouteResponseBody) SetNetLinkId(v string) *ListIoTCloudConnectorBackhaulRouteResponseBody {
	s.NetLinkId = &v
	return s
}

func (s *ListIoTCloudConnectorBackhaulRouteResponseBody) SetRequestId(v string) *ListIoTCloudConnectorBackhaulRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIoTCloudConnectorBackhaulRouteResponseBody) SetRoutes(v []*ListIoTCloudConnectorBackhaulRouteResponseBodyRoutes) *ListIoTCloudConnectorBackhaulRouteResponseBody {
	s.Routes = v
	return s
}

type ListIoTCloudConnectorBackhaulRouteResponseBodyRoutes struct {
	// 代表创建时间的资源属性字段
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 代表资源名称的资源属性字段
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// 代表创建时间的资源属性字段
	NextHopId *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	// 代表创建时间的资源属性字段
	NextHopType *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	// 代表创建时间的资源属性字段
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListIoTCloudConnectorBackhaulRouteResponseBodyRoutes) String() string {
	return tea.Prettify(s)
}

func (s ListIoTCloudConnectorBackhaulRouteResponseBodyRoutes) GoString() string {
	return s.String()
}

func (s *ListIoTCloudConnectorBackhaulRouteResponseBodyRoutes) SetDescription(v string) *ListIoTCloudConnectorBackhaulRouteResponseBodyRoutes {
	s.Description = &v
	return s
}

func (s *ListIoTCloudConnectorBackhaulRouteResponseBodyRoutes) SetDestinationCidrBlock(v string) *ListIoTCloudConnectorBackhaulRouteResponseBodyRoutes {
	s.DestinationCidrBlock = &v
	return s
}

func (s *ListIoTCloudConnectorBackhaulRouteResponseBodyRoutes) SetNextHopId(v string) *ListIoTCloudConnectorBackhaulRouteResponseBodyRoutes {
	s.NextHopId = &v
	return s
}

func (s *ListIoTCloudConnectorBackhaulRouteResponseBodyRoutes) SetNextHopType(v string) *ListIoTCloudConnectorBackhaulRouteResponseBodyRoutes {
	s.NextHopType = &v
	return s
}

func (s *ListIoTCloudConnectorBackhaulRouteResponseBodyRoutes) SetStatus(v string) *ListIoTCloudConnectorBackhaulRouteResponseBodyRoutes {
	s.Status = &v
	return s
}

type ListIoTCloudConnectorBackhaulRouteResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListIoTCloudConnectorBackhaulRouteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListIoTCloudConnectorBackhaulRouteResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIoTCloudConnectorBackhaulRouteResponse) GoString() string {
	return s.String()
}

func (s *ListIoTCloudConnectorBackhaulRouteResponse) SetHeaders(v map[string]*string) *ListIoTCloudConnectorBackhaulRouteResponse {
	s.Headers = v
	return s
}

func (s *ListIoTCloudConnectorBackhaulRouteResponse) SetStatusCode(v int32) *ListIoTCloudConnectorBackhaulRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIoTCloudConnectorBackhaulRouteResponse) SetBody(v *ListIoTCloudConnectorBackhaulRouteResponseBody) *ListIoTCloudConnectorBackhaulRouteResponse {
	s.Body = v
	return s
}

type ListOrdersRequest struct {
	MaxResults               *int64    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken                *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OrderAction              *string   `json:"OrderAction,omitempty" xml:"OrderAction,omitempty"`
	OrderIds                 []*string `json:"OrderIds,omitempty" xml:"OrderIds,omitempty" type:"Repeated"`
	Statuses                 []*string `json:"Statuses,omitempty" xml:"Statuses,omitempty" type:"Repeated"`
	WirelessCloudConnectorId *string   `json:"WirelessCloudConnectorId,omitempty" xml:"WirelessCloudConnectorId,omitempty"`
}

func (s ListOrdersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOrdersRequest) GoString() string {
	return s.String()
}

func (s *ListOrdersRequest) SetMaxResults(v int64) *ListOrdersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListOrdersRequest) SetNextToken(v string) *ListOrdersRequest {
	s.NextToken = &v
	return s
}

func (s *ListOrdersRequest) SetOrderAction(v string) *ListOrdersRequest {
	s.OrderAction = &v
	return s
}

func (s *ListOrdersRequest) SetOrderIds(v []*string) *ListOrdersRequest {
	s.OrderIds = v
	return s
}

func (s *ListOrdersRequest) SetStatuses(v []*string) *ListOrdersRequest {
	s.Statuses = v
	return s
}

func (s *ListOrdersRequest) SetWirelessCloudConnectorId(v string) *ListOrdersRequest {
	s.WirelessCloudConnectorId = &v
	return s
}

type ListOrdersResponseBody struct {
	MaxResults *string                         `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                         `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Orders     []*ListOrdersResponseBodyOrders `json:"Orders,omitempty" xml:"Orders,omitempty" type:"Repeated"`
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *string                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOrdersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOrdersResponseBody) GoString() string {
	return s.String()
}

func (s *ListOrdersResponseBody) SetMaxResults(v string) *ListOrdersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListOrdersResponseBody) SetNextToken(v string) *ListOrdersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListOrdersResponseBody) SetOrders(v []*ListOrdersResponseBodyOrders) *ListOrdersResponseBody {
	s.Orders = v
	return s
}

func (s *ListOrdersResponseBody) SetRequestId(v string) *ListOrdersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOrdersResponseBody) SetTotalCount(v string) *ListOrdersResponseBody {
	s.TotalCount = &v
	return s
}

type ListOrdersResponseBodyOrders struct {
	// 代表创建时间的资源属性字段
	Action              *string `json:"Action,omitempty" xml:"Action,omitempty"`
	CardCount           *string `json:"CardCount,omitempty" xml:"CardCount,omitempty"`
	CardNetType         *string `json:"CardNetType,omitempty" xml:"CardNetType,omitempty"`
	CardType            *string `json:"CardType,omitempty" xml:"CardType,omitempty"`
	ContactName         *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	ContactPhone        *string `json:"ContactPhone,omitempty" xml:"ContactPhone,omitempty"`
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	LogisticsId         *string `json:"LogisticsId,omitempty" xml:"LogisticsId,omitempty"`
	LogisticsStatus     *string `json:"LogisticsStatus,omitempty" xml:"LogisticsStatus,omitempty"`
	LogisticsType       *string `json:"LogisticsType,omitempty" xml:"LogisticsType,omitempty"`
	LogisticsUpdateTime *string `json:"LogisticsUpdateTime,omitempty" xml:"LogisticsUpdateTime,omitempty"`
	OrderId             *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	PayTime             *string `json:"PayTime,omitempty" xml:"PayTime,omitempty"`
	PostAddress         *string `json:"PostAddress,omitempty" xml:"PostAddress,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 代表资源名称的资源属性字段
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListOrdersResponseBodyOrders) String() string {
	return tea.Prettify(s)
}

func (s ListOrdersResponseBodyOrders) GoString() string {
	return s.String()
}

func (s *ListOrdersResponseBodyOrders) SetAction(v string) *ListOrdersResponseBodyOrders {
	s.Action = &v
	return s
}

func (s *ListOrdersResponseBodyOrders) SetCardCount(v string) *ListOrdersResponseBodyOrders {
	s.CardCount = &v
	return s
}

func (s *ListOrdersResponseBodyOrders) SetCardNetType(v string) *ListOrdersResponseBodyOrders {
	s.CardNetType = &v
	return s
}

func (s *ListOrdersResponseBodyOrders) SetCardType(v string) *ListOrdersResponseBodyOrders {
	s.CardType = &v
	return s
}

func (s *ListOrdersResponseBodyOrders) SetContactName(v string) *ListOrdersResponseBodyOrders {
	s.ContactName = &v
	return s
}

func (s *ListOrdersResponseBodyOrders) SetContactPhone(v string) *ListOrdersResponseBodyOrders {
	s.ContactPhone = &v
	return s
}

func (s *ListOrdersResponseBodyOrders) SetCreateTime(v string) *ListOrdersResponseBodyOrders {
	s.CreateTime = &v
	return s
}

func (s *ListOrdersResponseBodyOrders) SetDescription(v string) *ListOrdersResponseBodyOrders {
	s.Description = &v
	return s
}

func (s *ListOrdersResponseBodyOrders) SetLogisticsId(v string) *ListOrdersResponseBodyOrders {
	s.LogisticsId = &v
	return s
}

func (s *ListOrdersResponseBodyOrders) SetLogisticsStatus(v string) *ListOrdersResponseBodyOrders {
	s.LogisticsStatus = &v
	return s
}

func (s *ListOrdersResponseBodyOrders) SetLogisticsType(v string) *ListOrdersResponseBodyOrders {
	s.LogisticsType = &v
	return s
}

func (s *ListOrdersResponseBodyOrders) SetLogisticsUpdateTime(v string) *ListOrdersResponseBodyOrders {
	s.LogisticsUpdateTime = &v
	return s
}

func (s *ListOrdersResponseBodyOrders) SetOrderId(v string) *ListOrdersResponseBodyOrders {
	s.OrderId = &v
	return s
}

func (s *ListOrdersResponseBodyOrders) SetPayTime(v string) *ListOrdersResponseBodyOrders {
	s.PayTime = &v
	return s
}

func (s *ListOrdersResponseBodyOrders) SetPostAddress(v string) *ListOrdersResponseBodyOrders {
	s.PostAddress = &v
	return s
}

func (s *ListOrdersResponseBodyOrders) SetRegionId(v string) *ListOrdersResponseBodyOrders {
	s.RegionId = &v
	return s
}

func (s *ListOrdersResponseBodyOrders) SetStatus(v string) *ListOrdersResponseBodyOrders {
	s.Status = &v
	return s
}

type ListOrdersResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListOrdersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOrdersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOrdersResponse) GoString() string {
	return s.String()
}

func (s *ListOrdersResponse) SetHeaders(v map[string]*string) *ListOrdersResponse {
	s.Headers = v
	return s
}

func (s *ListOrdersResponse) SetStatusCode(v int32) *ListOrdersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOrdersResponse) SetBody(v *ListOrdersResponseBody) *ListOrdersResponse {
	s.Body = v
	return s
}

type ListRegionsRequest struct {
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsRequest) GoString() string {
	return s.String()
}

func (s *ListRegionsRequest) SetAcceptLanguage(v string) *ListRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListRegionsRequest) SetRegionId(v string) *ListRegionsRequest {
	s.RegionId = &v
	return s
}

type ListRegionsResponseBody struct {
	Regions   []*ListRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBody) SetRegions(v []*ListRegionsResponseBodyRegions) *ListRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *ListRegionsResponseBody) SetRequestId(v string) *ListRegionsResponseBody {
	s.RequestId = &v
	return s
}

type ListRegionsResponseBodyRegions struct {
	// 代表资源名称的资源属性字段
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// 代表资源一级ID的资源属性字段
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBodyRegions) SetLocalName(v string) *ListRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *ListRegionsResponseBodyRegions) SetRegionId(v string) *ListRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

type ListRegionsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponse) GoString() string {
	return s.String()
}

func (s *ListRegionsResponse) SetHeaders(v map[string]*string) *ListRegionsResponse {
	s.Headers = v
	return s
}

func (s *ListRegionsResponse) SetStatusCode(v int32) *ListRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRegionsResponse) SetBody(v *ListRegionsResponseBody) *ListRegionsResponse {
	s.Body = v
	return s
}

type ListWirelessCloudConnectorGroupsRequest struct {
	MaxResults                        *int64    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken                         *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId                          *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	WirelessCloudConnectorGroupIds    []*string `json:"WirelessCloudConnectorGroupIds,omitempty" xml:"WirelessCloudConnectorGroupIds,omitempty" type:"Repeated"`
	WirelessCloudConnectorGroupNames  []*string `json:"WirelessCloudConnectorGroupNames,omitempty" xml:"WirelessCloudConnectorGroupNames,omitempty" type:"Repeated"`
	WirelessCloudConnectorGroupStatus []*string `json:"WirelessCloudConnectorGroupStatus,omitempty" xml:"WirelessCloudConnectorGroupStatus,omitempty" type:"Repeated"`
}

func (s ListWirelessCloudConnectorGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWirelessCloudConnectorGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListWirelessCloudConnectorGroupsRequest) SetMaxResults(v int64) *ListWirelessCloudConnectorGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWirelessCloudConnectorGroupsRequest) SetNextToken(v string) *ListWirelessCloudConnectorGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListWirelessCloudConnectorGroupsRequest) SetRegionId(v string) *ListWirelessCloudConnectorGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *ListWirelessCloudConnectorGroupsRequest) SetWirelessCloudConnectorGroupIds(v []*string) *ListWirelessCloudConnectorGroupsRequest {
	s.WirelessCloudConnectorGroupIds = v
	return s
}

func (s *ListWirelessCloudConnectorGroupsRequest) SetWirelessCloudConnectorGroupNames(v []*string) *ListWirelessCloudConnectorGroupsRequest {
	s.WirelessCloudConnectorGroupNames = v
	return s
}

func (s *ListWirelessCloudConnectorGroupsRequest) SetWirelessCloudConnectorGroupStatus(v []*string) *ListWirelessCloudConnectorGroupsRequest {
	s.WirelessCloudConnectorGroupStatus = v
	return s
}

type ListWirelessCloudConnectorGroupsResponseBody struct {
	MaxResults                   *string                                                                     `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken                    *string                                                                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId                    *string                                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount                   *string                                                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	WirelessCloudConnectorGroups []*ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroups `json:"WirelessCloudConnectorGroups,omitempty" xml:"WirelessCloudConnectorGroups,omitempty" type:"Repeated"`
}

func (s ListWirelessCloudConnectorGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWirelessCloudConnectorGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListWirelessCloudConnectorGroupsResponseBody) SetMaxResults(v string) *ListWirelessCloudConnectorGroupsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListWirelessCloudConnectorGroupsResponseBody) SetNextToken(v string) *ListWirelessCloudConnectorGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListWirelessCloudConnectorGroupsResponseBody) SetRequestId(v string) *ListWirelessCloudConnectorGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWirelessCloudConnectorGroupsResponseBody) SetTotalCount(v string) *ListWirelessCloudConnectorGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListWirelessCloudConnectorGroupsResponseBody) SetWirelessCloudConnectorGroups(v []*ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroups) *ListWirelessCloudConnectorGroupsResponseBody {
	s.WirelessCloudConnectorGroups = v
	return s
}

type ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroups struct {
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 代表创建时间的资源属性字段
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 代表资源名称的资源属性字段
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 代表资源一级ID的资源属性字段
	WirelessCloudConnectorGroupId *string                                                                                            `json:"WirelessCloudConnectorGroupId,omitempty" xml:"WirelessCloudConnectorGroupId,omitempty"`
	WirelessCloudConnectors       []*ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroupsWirelessCloudConnectors `json:"WirelessCloudConnectors,omitempty" xml:"WirelessCloudConnectors,omitempty" type:"Repeated"`
}

func (s ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroups) String() string {
	return tea.Prettify(s)
}

func (s ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroups) GoString() string {
	return s.String()
}

func (s *ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroups) SetCreateTime(v string) *ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroups {
	s.CreateTime = &v
	return s
}

func (s *ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroups) SetDescription(v string) *ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroups {
	s.Description = &v
	return s
}

func (s *ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroups) SetName(v string) *ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroups {
	s.Name = &v
	return s
}

func (s *ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroups) SetRegionId(v string) *ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroups {
	s.RegionId = &v
	return s
}

func (s *ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroups) SetStatus(v string) *ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroups {
	s.Status = &v
	return s
}

func (s *ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroups) SetWirelessCloudConnectorGroupId(v string) *ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroups {
	s.WirelessCloudConnectorGroupId = &v
	return s
}

func (s *ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroups) SetWirelessCloudConnectors(v []*ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroupsWirelessCloudConnectors) *ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroups {
	s.WirelessCloudConnectors = v
	return s
}

type ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroupsWirelessCloudConnectors struct {
	BusinessType             *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	CardCount                *string `json:"CardCount,omitempty" xml:"CardCount,omitempty"`
	CreateTime               *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DataPackageId            *string `json:"DataPackageId,omitempty" xml:"DataPackageId,omitempty"`
	DataPackageType          *string `json:"DataPackageType,omitempty" xml:"DataPackageType,omitempty"`
	Description              *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name                     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceType              *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	Status                   *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UseCase                  *string `json:"UseCase,omitempty" xml:"UseCase,omitempty"`
	WirelessCloudConnectorId *string `json:"WirelessCloudConnectorId,omitempty" xml:"WirelessCloudConnectorId,omitempty"`
}

func (s ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroupsWirelessCloudConnectors) String() string {
	return tea.Prettify(s)
}

func (s ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroupsWirelessCloudConnectors) GoString() string {
	return s.String()
}

func (s *ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroupsWirelessCloudConnectors) SetBusinessType(v string) *ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroupsWirelessCloudConnectors {
	s.BusinessType = &v
	return s
}

func (s *ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroupsWirelessCloudConnectors) SetCardCount(v string) *ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroupsWirelessCloudConnectors {
	s.CardCount = &v
	return s
}

func (s *ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroupsWirelessCloudConnectors) SetCreateTime(v string) *ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroupsWirelessCloudConnectors {
	s.CreateTime = &v
	return s
}

func (s *ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroupsWirelessCloudConnectors) SetDataPackageId(v string) *ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroupsWirelessCloudConnectors {
	s.DataPackageId = &v
	return s
}

func (s *ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroupsWirelessCloudConnectors) SetDataPackageType(v string) *ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroupsWirelessCloudConnectors {
	s.DataPackageType = &v
	return s
}

func (s *ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroupsWirelessCloudConnectors) SetDescription(v string) *ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroupsWirelessCloudConnectors {
	s.Description = &v
	return s
}

func (s *ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroupsWirelessCloudConnectors) SetName(v string) *ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroupsWirelessCloudConnectors {
	s.Name = &v
	return s
}

func (s *ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroupsWirelessCloudConnectors) SetRegionId(v string) *ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroupsWirelessCloudConnectors {
	s.RegionId = &v
	return s
}

func (s *ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroupsWirelessCloudConnectors) SetServiceType(v string) *ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroupsWirelessCloudConnectors {
	s.ServiceType = &v
	return s
}

func (s *ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroupsWirelessCloudConnectors) SetStatus(v string) *ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroupsWirelessCloudConnectors {
	s.Status = &v
	return s
}

func (s *ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroupsWirelessCloudConnectors) SetUseCase(v string) *ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroupsWirelessCloudConnectors {
	s.UseCase = &v
	return s
}

func (s *ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroupsWirelessCloudConnectors) SetWirelessCloudConnectorId(v string) *ListWirelessCloudConnectorGroupsResponseBodyWirelessCloudConnectorGroupsWirelessCloudConnectors {
	s.WirelessCloudConnectorId = &v
	return s
}

type ListWirelessCloudConnectorGroupsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListWirelessCloudConnectorGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListWirelessCloudConnectorGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWirelessCloudConnectorGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListWirelessCloudConnectorGroupsResponse) SetHeaders(v map[string]*string) *ListWirelessCloudConnectorGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListWirelessCloudConnectorGroupsResponse) SetStatusCode(v int32) *ListWirelessCloudConnectorGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWirelessCloudConnectorGroupsResponse) SetBody(v *ListWirelessCloudConnectorGroupsResponseBody) *ListWirelessCloudConnectorGroupsResponse {
	s.Body = v
	return s
}

type ListWirelessCloudConnectorsRequest struct {
	BusinessType                  *string   `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	IsInGroup                     *string   `json:"IsInGroup,omitempty" xml:"IsInGroup,omitempty"`
	MaxResults                    *int64    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Names                         []*string `json:"Names,omitempty" xml:"Names,omitempty" type:"Repeated"`
	NextToken                     *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId                      *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Statuses                      []*string `json:"Statuses,omitempty" xml:"Statuses,omitempty" type:"Repeated"`
	WirelessCloudConnectorGroupId *string   `json:"WirelessCloudConnectorGroupId,omitempty" xml:"WirelessCloudConnectorGroupId,omitempty"`
	WirelessCloudConnectorIds     []*string `json:"WirelessCloudConnectorIds,omitempty" xml:"WirelessCloudConnectorIds,omitempty" type:"Repeated"`
}

func (s ListWirelessCloudConnectorsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWirelessCloudConnectorsRequest) GoString() string {
	return s.String()
}

func (s *ListWirelessCloudConnectorsRequest) SetBusinessType(v string) *ListWirelessCloudConnectorsRequest {
	s.BusinessType = &v
	return s
}

func (s *ListWirelessCloudConnectorsRequest) SetIsInGroup(v string) *ListWirelessCloudConnectorsRequest {
	s.IsInGroup = &v
	return s
}

func (s *ListWirelessCloudConnectorsRequest) SetMaxResults(v int64) *ListWirelessCloudConnectorsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWirelessCloudConnectorsRequest) SetNames(v []*string) *ListWirelessCloudConnectorsRequest {
	s.Names = v
	return s
}

func (s *ListWirelessCloudConnectorsRequest) SetNextToken(v string) *ListWirelessCloudConnectorsRequest {
	s.NextToken = &v
	return s
}

func (s *ListWirelessCloudConnectorsRequest) SetRegionId(v string) *ListWirelessCloudConnectorsRequest {
	s.RegionId = &v
	return s
}

func (s *ListWirelessCloudConnectorsRequest) SetStatuses(v []*string) *ListWirelessCloudConnectorsRequest {
	s.Statuses = v
	return s
}

func (s *ListWirelessCloudConnectorsRequest) SetWirelessCloudConnectorGroupId(v string) *ListWirelessCloudConnectorsRequest {
	s.WirelessCloudConnectorGroupId = &v
	return s
}

func (s *ListWirelessCloudConnectorsRequest) SetWirelessCloudConnectorIds(v []*string) *ListWirelessCloudConnectorsRequest {
	s.WirelessCloudConnectorIds = v
	return s
}

type ListWirelessCloudConnectorsResponseBody struct {
	MaxResults              *string                                                           `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken               *string                                                           `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId               *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount              *string                                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	WirelessCloudConnectors []*ListWirelessCloudConnectorsResponseBodyWirelessCloudConnectors `json:"WirelessCloudConnectors,omitempty" xml:"WirelessCloudConnectors,omitempty" type:"Repeated"`
}

func (s ListWirelessCloudConnectorsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWirelessCloudConnectorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListWirelessCloudConnectorsResponseBody) SetMaxResults(v string) *ListWirelessCloudConnectorsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListWirelessCloudConnectorsResponseBody) SetNextToken(v string) *ListWirelessCloudConnectorsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListWirelessCloudConnectorsResponseBody) SetRequestId(v string) *ListWirelessCloudConnectorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWirelessCloudConnectorsResponseBody) SetTotalCount(v string) *ListWirelessCloudConnectorsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListWirelessCloudConnectorsResponseBody) SetWirelessCloudConnectors(v []*ListWirelessCloudConnectorsResponseBodyWirelessCloudConnectors) *ListWirelessCloudConnectorsResponseBody {
	s.WirelessCloudConnectors = v
	return s
}

type ListWirelessCloudConnectorsResponseBodyWirelessCloudConnectors struct {
	BusinessType    *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	CardCount       *string `json:"CardCount,omitempty" xml:"CardCount,omitempty"`
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DataPackageId   *string `json:"DataPackageId,omitempty" xml:"DataPackageId,omitempty"`
	DataPackageType *string `json:"DataPackageType,omitempty" xml:"DataPackageType,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 代表创建时间的资源属性字段
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// 代表资源名称的资源属性字段
	Status                        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UseCase                       *string `json:"UseCase,omitempty" xml:"UseCase,omitempty"`
	WirelessCloudConnectorGroupId *string `json:"WirelessCloudConnectorGroupId,omitempty" xml:"WirelessCloudConnectorGroupId,omitempty"`
	// 代表资源一级ID的资源属性字段
	WirelessCloudConnectorId *string `json:"WirelessCloudConnectorId,omitempty" xml:"WirelessCloudConnectorId,omitempty"`
}

func (s ListWirelessCloudConnectorsResponseBodyWirelessCloudConnectors) String() string {
	return tea.Prettify(s)
}

func (s ListWirelessCloudConnectorsResponseBodyWirelessCloudConnectors) GoString() string {
	return s.String()
}

func (s *ListWirelessCloudConnectorsResponseBodyWirelessCloudConnectors) SetBusinessType(v string) *ListWirelessCloudConnectorsResponseBodyWirelessCloudConnectors {
	s.BusinessType = &v
	return s
}

func (s *ListWirelessCloudConnectorsResponseBodyWirelessCloudConnectors) SetCardCount(v string) *ListWirelessCloudConnectorsResponseBodyWirelessCloudConnectors {
	s.CardCount = &v
	return s
}

func (s *ListWirelessCloudConnectorsResponseBodyWirelessCloudConnectors) SetCreateTime(v string) *ListWirelessCloudConnectorsResponseBodyWirelessCloudConnectors {
	s.CreateTime = &v
	return s
}

func (s *ListWirelessCloudConnectorsResponseBodyWirelessCloudConnectors) SetDataPackageId(v string) *ListWirelessCloudConnectorsResponseBodyWirelessCloudConnectors {
	s.DataPackageId = &v
	return s
}

func (s *ListWirelessCloudConnectorsResponseBodyWirelessCloudConnectors) SetDataPackageType(v string) *ListWirelessCloudConnectorsResponseBodyWirelessCloudConnectors {
	s.DataPackageType = &v
	return s
}

func (s *ListWirelessCloudConnectorsResponseBodyWirelessCloudConnectors) SetDescription(v string) *ListWirelessCloudConnectorsResponseBodyWirelessCloudConnectors {
	s.Description = &v
	return s
}

func (s *ListWirelessCloudConnectorsResponseBodyWirelessCloudConnectors) SetName(v string) *ListWirelessCloudConnectorsResponseBodyWirelessCloudConnectors {
	s.Name = &v
	return s
}

func (s *ListWirelessCloudConnectorsResponseBodyWirelessCloudConnectors) SetRegionId(v string) *ListWirelessCloudConnectorsResponseBodyWirelessCloudConnectors {
	s.RegionId = &v
	return s
}

func (s *ListWirelessCloudConnectorsResponseBodyWirelessCloudConnectors) SetServiceType(v string) *ListWirelessCloudConnectorsResponseBodyWirelessCloudConnectors {
	s.ServiceType = &v
	return s
}

func (s *ListWirelessCloudConnectorsResponseBodyWirelessCloudConnectors) SetStatus(v string) *ListWirelessCloudConnectorsResponseBodyWirelessCloudConnectors {
	s.Status = &v
	return s
}

func (s *ListWirelessCloudConnectorsResponseBodyWirelessCloudConnectors) SetUseCase(v string) *ListWirelessCloudConnectorsResponseBodyWirelessCloudConnectors {
	s.UseCase = &v
	return s
}

func (s *ListWirelessCloudConnectorsResponseBodyWirelessCloudConnectors) SetWirelessCloudConnectorGroupId(v string) *ListWirelessCloudConnectorsResponseBodyWirelessCloudConnectors {
	s.WirelessCloudConnectorGroupId = &v
	return s
}

func (s *ListWirelessCloudConnectorsResponseBodyWirelessCloudConnectors) SetWirelessCloudConnectorId(v string) *ListWirelessCloudConnectorsResponseBodyWirelessCloudConnectors {
	s.WirelessCloudConnectorId = &v
	return s
}

type ListWirelessCloudConnectorsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListWirelessCloudConnectorsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListWirelessCloudConnectorsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWirelessCloudConnectorsResponse) GoString() string {
	return s.String()
}

func (s *ListWirelessCloudConnectorsResponse) SetHeaders(v map[string]*string) *ListWirelessCloudConnectorsResponse {
	s.Headers = v
	return s
}

func (s *ListWirelessCloudConnectorsResponse) SetStatusCode(v int32) *ListWirelessCloudConnectorsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWirelessCloudConnectorsResponse) SetBody(v *ListWirelessCloudConnectorsResponseBody) *ListWirelessCloudConnectorsResponse {
	s.Body = v
	return s
}

type ListZonesRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListZonesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListZonesRequest) GoString() string {
	return s.String()
}

func (s *ListZonesRequest) SetRegionId(v string) *ListZonesRequest {
	s.RegionId = &v
	return s
}

type ListZonesResponseBody struct {
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Zones     []*ListZonesResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s ListZonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListZonesResponseBody) GoString() string {
	return s.String()
}

func (s *ListZonesResponseBody) SetRequestId(v string) *ListZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListZonesResponseBody) SetZones(v []*ListZonesResponseBodyZones) *ListZonesResponseBody {
	s.Zones = v
	return s
}

type ListZonesResponseBodyZones struct {
	// 代表创建时间的资源属性字段
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// 代表资源名称的资源属性字段
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListZonesResponseBodyZones) String() string {
	return tea.Prettify(s)
}

func (s ListZonesResponseBodyZones) GoString() string {
	return s.String()
}

func (s *ListZonesResponseBodyZones) SetLocalName(v string) *ListZonesResponseBodyZones {
	s.LocalName = &v
	return s
}

func (s *ListZonesResponseBodyZones) SetZoneId(v string) *ListZonesResponseBodyZones {
	s.ZoneId = &v
	return s
}

type ListZonesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListZonesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListZonesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListZonesResponse) GoString() string {
	return s.String()
}

func (s *ListZonesResponse) SetHeaders(v map[string]*string) *ListZonesResponse {
	s.Headers = v
	return s
}

func (s *ListZonesResponse) SetStatusCode(v int32) *ListZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListZonesResponse) SetBody(v *ListZonesResponseBody) *ListZonesResponse {
	s.Body = v
	return s
}

type LockCardsRequest struct {
	ClientToken *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	Iccids      []*string `json:"Iccids,omitempty" xml:"Iccids,omitempty" type:"Repeated"`
	RegionId    *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s LockCardsRequest) String() string {
	return tea.Prettify(s)
}

func (s LockCardsRequest) GoString() string {
	return s.String()
}

func (s *LockCardsRequest) SetClientToken(v string) *LockCardsRequest {
	s.ClientToken = &v
	return s
}

func (s *LockCardsRequest) SetDryRun(v bool) *LockCardsRequest {
	s.DryRun = &v
	return s
}

func (s *LockCardsRequest) SetIccids(v []*string) *LockCardsRequest {
	s.Iccids = v
	return s
}

func (s *LockCardsRequest) SetRegionId(v string) *LockCardsRequest {
	s.RegionId = &v
	return s
}

type LockCardsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LockCardsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LockCardsResponseBody) GoString() string {
	return s.String()
}

func (s *LockCardsResponseBody) SetRequestId(v string) *LockCardsResponseBody {
	s.RequestId = &v
	return s
}

type LockCardsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *LockCardsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LockCardsResponse) String() string {
	return tea.Prettify(s)
}

func (s LockCardsResponse) GoString() string {
	return s.String()
}

func (s *LockCardsResponse) SetHeaders(v map[string]*string) *LockCardsResponse {
	s.Headers = v
	return s
}

func (s *LockCardsResponse) SetStatusCode(v int32) *LockCardsResponse {
	s.StatusCode = &v
	return s
}

func (s *LockCardsResponse) SetBody(v *LockCardsResponseBody) *LockCardsResponse {
	s.Body = v
	return s
}

type ModifyWirelessCloudConnectorFeatureRequest struct {
	ClientToken              *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                   *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	FeatureName              *string `json:"FeatureName,omitempty" xml:"FeatureName,omitempty"`
	FeatureValue             *string `json:"FeatureValue,omitempty" xml:"FeatureValue,omitempty"`
	WirelessCloudConnectorId *string `json:"WirelessCloudConnectorId,omitempty" xml:"WirelessCloudConnectorId,omitempty"`
}

func (s ModifyWirelessCloudConnectorFeatureRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyWirelessCloudConnectorFeatureRequest) GoString() string {
	return s.String()
}

func (s *ModifyWirelessCloudConnectorFeatureRequest) SetClientToken(v string) *ModifyWirelessCloudConnectorFeatureRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyWirelessCloudConnectorFeatureRequest) SetDryRun(v bool) *ModifyWirelessCloudConnectorFeatureRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyWirelessCloudConnectorFeatureRequest) SetFeatureName(v string) *ModifyWirelessCloudConnectorFeatureRequest {
	s.FeatureName = &v
	return s
}

func (s *ModifyWirelessCloudConnectorFeatureRequest) SetFeatureValue(v string) *ModifyWirelessCloudConnectorFeatureRequest {
	s.FeatureValue = &v
	return s
}

func (s *ModifyWirelessCloudConnectorFeatureRequest) SetWirelessCloudConnectorId(v string) *ModifyWirelessCloudConnectorFeatureRequest {
	s.WirelessCloudConnectorId = &v
	return s
}

type ModifyWirelessCloudConnectorFeatureResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWirelessCloudConnectorFeatureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyWirelessCloudConnectorFeatureResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWirelessCloudConnectorFeatureResponseBody) SetRequestId(v string) *ModifyWirelessCloudConnectorFeatureResponseBody {
	s.RequestId = &v
	return s
}

type ModifyWirelessCloudConnectorFeatureResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyWirelessCloudConnectorFeatureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyWirelessCloudConnectorFeatureResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyWirelessCloudConnectorFeatureResponse) GoString() string {
	return s.String()
}

func (s *ModifyWirelessCloudConnectorFeatureResponse) SetHeaders(v map[string]*string) *ModifyWirelessCloudConnectorFeatureResponse {
	s.Headers = v
	return s
}

func (s *ModifyWirelessCloudConnectorFeatureResponse) SetStatusCode(v int32) *ModifyWirelessCloudConnectorFeatureResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyWirelessCloudConnectorFeatureResponse) SetBody(v *ModifyWirelessCloudConnectorFeatureResponseBody) *ModifyWirelessCloudConnectorFeatureResponse {
	s.Body = v
	return s
}

type OpenCc5gServiceRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s OpenCc5gServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenCc5gServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenCc5gServiceRequest) SetRegionId(v string) *OpenCc5gServiceRequest {
	s.RegionId = &v
	return s
}

type OpenCc5gServiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenCc5gServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenCc5gServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenCc5gServiceResponseBody) SetRequestId(v string) *OpenCc5gServiceResponseBody {
	s.RequestId = &v
	return s
}

type OpenCc5gServiceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OpenCc5gServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpenCc5gServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenCc5gServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenCc5gServiceResponse) SetHeaders(v map[string]*string) *OpenCc5gServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenCc5gServiceResponse) SetStatusCode(v int32) *OpenCc5gServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenCc5gServiceResponse) SetBody(v *OpenCc5gServiceResponseBody) *OpenCc5gServiceResponse {
	s.Body = v
	return s
}

type RebindCardsRequest struct {
	ClientToken *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	Iccids      []*string `json:"Iccids,omitempty" xml:"Iccids,omitempty" type:"Repeated"`
	RegionId    *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RebindCardsRequest) String() string {
	return tea.Prettify(s)
}

func (s RebindCardsRequest) GoString() string {
	return s.String()
}

func (s *RebindCardsRequest) SetClientToken(v string) *RebindCardsRequest {
	s.ClientToken = &v
	return s
}

func (s *RebindCardsRequest) SetDryRun(v bool) *RebindCardsRequest {
	s.DryRun = &v
	return s
}

func (s *RebindCardsRequest) SetIccids(v []*string) *RebindCardsRequest {
	s.Iccids = v
	return s
}

func (s *RebindCardsRequest) SetRegionId(v string) *RebindCardsRequest {
	s.RegionId = &v
	return s
}

type RebindCardsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RebindCardsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RebindCardsResponseBody) GoString() string {
	return s.String()
}

func (s *RebindCardsResponseBody) SetRequestId(v string) *RebindCardsResponseBody {
	s.RequestId = &v
	return s
}

type RebindCardsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RebindCardsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RebindCardsResponse) String() string {
	return tea.Prettify(s)
}

func (s RebindCardsResponse) GoString() string {
	return s.String()
}

func (s *RebindCardsResponse) SetHeaders(v map[string]*string) *RebindCardsResponse {
	s.Headers = v
	return s
}

func (s *RebindCardsResponse) SetStatusCode(v int32) *RebindCardsResponse {
	s.StatusCode = &v
	return s
}

func (s *RebindCardsResponse) SetBody(v *RebindCardsResponseBody) *RebindCardsResponse {
	s.Body = v
	return s
}

type RemoveWirelessCloudConnectorFromGroupRequest struct {
	ClientToken                   *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                        *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	RegionId                      *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	WirelessCloudConnectorGroupId *string   `json:"WirelessCloudConnectorGroupId,omitempty" xml:"WirelessCloudConnectorGroupId,omitempty"`
	WirelessCloudConnectorIds     []*string `json:"WirelessCloudConnectorIds,omitempty" xml:"WirelessCloudConnectorIds,omitempty" type:"Repeated"`
}

func (s RemoveWirelessCloudConnectorFromGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveWirelessCloudConnectorFromGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveWirelessCloudConnectorFromGroupRequest) SetClientToken(v string) *RemoveWirelessCloudConnectorFromGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *RemoveWirelessCloudConnectorFromGroupRequest) SetDryRun(v bool) *RemoveWirelessCloudConnectorFromGroupRequest {
	s.DryRun = &v
	return s
}

func (s *RemoveWirelessCloudConnectorFromGroupRequest) SetRegionId(v string) *RemoveWirelessCloudConnectorFromGroupRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveWirelessCloudConnectorFromGroupRequest) SetWirelessCloudConnectorGroupId(v string) *RemoveWirelessCloudConnectorFromGroupRequest {
	s.WirelessCloudConnectorGroupId = &v
	return s
}

func (s *RemoveWirelessCloudConnectorFromGroupRequest) SetWirelessCloudConnectorIds(v []*string) *RemoveWirelessCloudConnectorFromGroupRequest {
	s.WirelessCloudConnectorIds = v
	return s
}

type RemoveWirelessCloudConnectorFromGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveWirelessCloudConnectorFromGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveWirelessCloudConnectorFromGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveWirelessCloudConnectorFromGroupResponseBody) SetRequestId(v string) *RemoveWirelessCloudConnectorFromGroupResponseBody {
	s.RequestId = &v
	return s
}

type RemoveWirelessCloudConnectorFromGroupResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveWirelessCloudConnectorFromGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveWirelessCloudConnectorFromGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveWirelessCloudConnectorFromGroupResponse) GoString() string {
	return s.String()
}

func (s *RemoveWirelessCloudConnectorFromGroupResponse) SetHeaders(v map[string]*string) *RemoveWirelessCloudConnectorFromGroupResponse {
	s.Headers = v
	return s
}

func (s *RemoveWirelessCloudConnectorFromGroupResponse) SetStatusCode(v int32) *RemoveWirelessCloudConnectorFromGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveWirelessCloudConnectorFromGroupResponse) SetBody(v *RemoveWirelessCloudConnectorFromGroupResponseBody) *RemoveWirelessCloudConnectorFromGroupResponse {
	s.Body = v
	return s
}

type ResumeCardsRequest struct {
	ClientToken *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	Iccids      []*string `json:"Iccids,omitempty" xml:"Iccids,omitempty" type:"Repeated"`
	RegionId    *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ResumeCardsRequest) String() string {
	return tea.Prettify(s)
}

func (s ResumeCardsRequest) GoString() string {
	return s.String()
}

func (s *ResumeCardsRequest) SetClientToken(v string) *ResumeCardsRequest {
	s.ClientToken = &v
	return s
}

func (s *ResumeCardsRequest) SetDryRun(v bool) *ResumeCardsRequest {
	s.DryRun = &v
	return s
}

func (s *ResumeCardsRequest) SetIccids(v []*string) *ResumeCardsRequest {
	s.Iccids = v
	return s
}

func (s *ResumeCardsRequest) SetRegionId(v string) *ResumeCardsRequest {
	s.RegionId = &v
	return s
}

type ResumeCardsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResumeCardsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResumeCardsResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeCardsResponseBody) SetRequestId(v string) *ResumeCardsResponseBody {
	s.RequestId = &v
	return s
}

type ResumeCardsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResumeCardsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResumeCardsResponse) String() string {
	return tea.Prettify(s)
}

func (s ResumeCardsResponse) GoString() string {
	return s.String()
}

func (s *ResumeCardsResponse) SetHeaders(v map[string]*string) *ResumeCardsResponse {
	s.Headers = v
	return s
}

func (s *ResumeCardsResponse) SetStatusCode(v int32) *ResumeCardsResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeCardsResponse) SetBody(v *ResumeCardsResponseBody) *ResumeCardsResponse {
	s.Body = v
	return s
}

type RevokeNetLinkRequest struct {
	ClientToken              *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                   *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	NetLinkId                *string `json:"NetLinkId,omitempty" xml:"NetLinkId,omitempty"`
	WirelessCloudConnectorId *string `json:"WirelessCloudConnectorId,omitempty" xml:"WirelessCloudConnectorId,omitempty"`
}

func (s RevokeNetLinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RevokeNetLinkRequest) GoString() string {
	return s.String()
}

func (s *RevokeNetLinkRequest) SetClientToken(v string) *RevokeNetLinkRequest {
	s.ClientToken = &v
	return s
}

func (s *RevokeNetLinkRequest) SetDryRun(v bool) *RevokeNetLinkRequest {
	s.DryRun = &v
	return s
}

func (s *RevokeNetLinkRequest) SetNetLinkId(v string) *RevokeNetLinkRequest {
	s.NetLinkId = &v
	return s
}

func (s *RevokeNetLinkRequest) SetWirelessCloudConnectorId(v string) *RevokeNetLinkRequest {
	s.WirelessCloudConnectorId = &v
	return s
}

type RevokeNetLinkResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeNetLinkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RevokeNetLinkResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeNetLinkResponseBody) SetRequestId(v string) *RevokeNetLinkResponseBody {
	s.RequestId = &v
	return s
}

type RevokeNetLinkResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RevokeNetLinkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RevokeNetLinkResponse) String() string {
	return tea.Prettify(s)
}

func (s RevokeNetLinkResponse) GoString() string {
	return s.String()
}

func (s *RevokeNetLinkResponse) SetHeaders(v map[string]*string) *RevokeNetLinkResponse {
	s.Headers = v
	return s
}

func (s *RevokeNetLinkResponse) SetStatusCode(v int32) *RevokeNetLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeNetLinkResponse) SetBody(v *RevokeNetLinkResponseBody) *RevokeNetLinkResponse {
	s.Body = v
	return s
}

type StopCardsRequest struct {
	ClientToken *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	Iccids      []*string `json:"Iccids,omitempty" xml:"Iccids,omitempty" type:"Repeated"`
	RegionId    *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StopCardsRequest) String() string {
	return tea.Prettify(s)
}

func (s StopCardsRequest) GoString() string {
	return s.String()
}

func (s *StopCardsRequest) SetClientToken(v string) *StopCardsRequest {
	s.ClientToken = &v
	return s
}

func (s *StopCardsRequest) SetDryRun(v bool) *StopCardsRequest {
	s.DryRun = &v
	return s
}

func (s *StopCardsRequest) SetIccids(v []*string) *StopCardsRequest {
	s.Iccids = v
	return s
}

func (s *StopCardsRequest) SetRegionId(v string) *StopCardsRequest {
	s.RegionId = &v
	return s
}

type StopCardsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopCardsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopCardsResponseBody) GoString() string {
	return s.String()
}

func (s *StopCardsResponseBody) SetRequestId(v string) *StopCardsResponseBody {
	s.RequestId = &v
	return s
}

type StopCardsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopCardsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopCardsResponse) String() string {
	return tea.Prettify(s)
}

func (s StopCardsResponse) GoString() string {
	return s.String()
}

func (s *StopCardsResponse) SetHeaders(v map[string]*string) *StopCardsResponse {
	s.Headers = v
	return s
}

func (s *StopCardsResponse) SetStatusCode(v int32) *StopCardsResponse {
	s.StatusCode = &v
	return s
}

func (s *StopCardsResponse) SetBody(v *StopCardsResponseBody) *StopCardsResponse {
	s.Body = v
	return s
}

type SubmitDiagnoseTaskForSingleCardRequest struct {
	BeginTime                *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	Destination              *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	EndTime                  *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RegionNo                 *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	ResourceUid              *int64  `json:"ResourceUid,omitempty" xml:"ResourceUid,omitempty"`
	Source                   *string `json:"Source,omitempty" xml:"Source,omitempty"`
	WirelessCloudConnectorId *string `json:"WirelessCloudConnectorId,omitempty" xml:"WirelessCloudConnectorId,omitempty"`
}

func (s SubmitDiagnoseTaskForSingleCardRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitDiagnoseTaskForSingleCardRequest) GoString() string {
	return s.String()
}

func (s *SubmitDiagnoseTaskForSingleCardRequest) SetBeginTime(v int64) *SubmitDiagnoseTaskForSingleCardRequest {
	s.BeginTime = &v
	return s
}

func (s *SubmitDiagnoseTaskForSingleCardRequest) SetDestination(v string) *SubmitDiagnoseTaskForSingleCardRequest {
	s.Destination = &v
	return s
}

func (s *SubmitDiagnoseTaskForSingleCardRequest) SetEndTime(v int64) *SubmitDiagnoseTaskForSingleCardRequest {
	s.EndTime = &v
	return s
}

func (s *SubmitDiagnoseTaskForSingleCardRequest) SetRegionNo(v string) *SubmitDiagnoseTaskForSingleCardRequest {
	s.RegionNo = &v
	return s
}

func (s *SubmitDiagnoseTaskForSingleCardRequest) SetResourceUid(v int64) *SubmitDiagnoseTaskForSingleCardRequest {
	s.ResourceUid = &v
	return s
}

func (s *SubmitDiagnoseTaskForSingleCardRequest) SetSource(v string) *SubmitDiagnoseTaskForSingleCardRequest {
	s.Source = &v
	return s
}

func (s *SubmitDiagnoseTaskForSingleCardRequest) SetWirelessCloudConnectorId(v string) *SubmitDiagnoseTaskForSingleCardRequest {
	s.WirelessCloudConnectorId = &v
	return s
}

type SubmitDiagnoseTaskForSingleCardResponseBody struct {
	DiagnoseTaskId *string `json:"DiagnoseTaskId,omitempty" xml:"DiagnoseTaskId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitDiagnoseTaskForSingleCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitDiagnoseTaskForSingleCardResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitDiagnoseTaskForSingleCardResponseBody) SetDiagnoseTaskId(v string) *SubmitDiagnoseTaskForSingleCardResponseBody {
	s.DiagnoseTaskId = &v
	return s
}

func (s *SubmitDiagnoseTaskForSingleCardResponseBody) SetRequestId(v string) *SubmitDiagnoseTaskForSingleCardResponseBody {
	s.RequestId = &v
	return s
}

type SubmitDiagnoseTaskForSingleCardResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitDiagnoseTaskForSingleCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitDiagnoseTaskForSingleCardResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitDiagnoseTaskForSingleCardResponse) GoString() string {
	return s.String()
}

func (s *SubmitDiagnoseTaskForSingleCardResponse) SetHeaders(v map[string]*string) *SubmitDiagnoseTaskForSingleCardResponse {
	s.Headers = v
	return s
}

func (s *SubmitDiagnoseTaskForSingleCardResponse) SetStatusCode(v int32) *SubmitDiagnoseTaskForSingleCardResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitDiagnoseTaskForSingleCardResponse) SetBody(v *SubmitDiagnoseTaskForSingleCardResponseBody) *SubmitDiagnoseTaskForSingleCardResponse {
	s.Body = v
	return s
}

type SwitchWirelessCloudConnectorToBusinessRequest struct {
	ClientToken              *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                   *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	WirelessCloudConnectorId *string `json:"WirelessCloudConnectorId,omitempty" xml:"WirelessCloudConnectorId,omitempty"`
}

func (s SwitchWirelessCloudConnectorToBusinessRequest) String() string {
	return tea.Prettify(s)
}

func (s SwitchWirelessCloudConnectorToBusinessRequest) GoString() string {
	return s.String()
}

func (s *SwitchWirelessCloudConnectorToBusinessRequest) SetClientToken(v string) *SwitchWirelessCloudConnectorToBusinessRequest {
	s.ClientToken = &v
	return s
}

func (s *SwitchWirelessCloudConnectorToBusinessRequest) SetDryRun(v bool) *SwitchWirelessCloudConnectorToBusinessRequest {
	s.DryRun = &v
	return s
}

func (s *SwitchWirelessCloudConnectorToBusinessRequest) SetWirelessCloudConnectorId(v string) *SwitchWirelessCloudConnectorToBusinessRequest {
	s.WirelessCloudConnectorId = &v
	return s
}

type SwitchWirelessCloudConnectorToBusinessResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SwitchWirelessCloudConnectorToBusinessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SwitchWirelessCloudConnectorToBusinessResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchWirelessCloudConnectorToBusinessResponseBody) SetRequestId(v string) *SwitchWirelessCloudConnectorToBusinessResponseBody {
	s.RequestId = &v
	return s
}

type SwitchWirelessCloudConnectorToBusinessResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SwitchWirelessCloudConnectorToBusinessResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SwitchWirelessCloudConnectorToBusinessResponse) String() string {
	return tea.Prettify(s)
}

func (s SwitchWirelessCloudConnectorToBusinessResponse) GoString() string {
	return s.String()
}

func (s *SwitchWirelessCloudConnectorToBusinessResponse) SetHeaders(v map[string]*string) *SwitchWirelessCloudConnectorToBusinessResponse {
	s.Headers = v
	return s
}

func (s *SwitchWirelessCloudConnectorToBusinessResponse) SetStatusCode(v int32) *SwitchWirelessCloudConnectorToBusinessResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchWirelessCloudConnectorToBusinessResponse) SetBody(v *SwitchWirelessCloudConnectorToBusinessResponseBody) *SwitchWirelessCloudConnectorToBusinessResponse {
	s.Body = v
	return s
}

type UnlockCardsRequest struct {
	ClientToken *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	Iccids      []*string `json:"Iccids,omitempty" xml:"Iccids,omitempty" type:"Repeated"`
	RegionId    *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UnlockCardsRequest) String() string {
	return tea.Prettify(s)
}

func (s UnlockCardsRequest) GoString() string {
	return s.String()
}

func (s *UnlockCardsRequest) SetClientToken(v string) *UnlockCardsRequest {
	s.ClientToken = &v
	return s
}

func (s *UnlockCardsRequest) SetDryRun(v bool) *UnlockCardsRequest {
	s.DryRun = &v
	return s
}

func (s *UnlockCardsRequest) SetIccids(v []*string) *UnlockCardsRequest {
	s.Iccids = v
	return s
}

func (s *UnlockCardsRequest) SetRegionId(v string) *UnlockCardsRequest {
	s.RegionId = &v
	return s
}

type UnlockCardsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnlockCardsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnlockCardsResponseBody) GoString() string {
	return s.String()
}

func (s *UnlockCardsResponseBody) SetRequestId(v string) *UnlockCardsResponseBody {
	s.RequestId = &v
	return s
}

type UnlockCardsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnlockCardsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnlockCardsResponse) String() string {
	return tea.Prettify(s)
}

func (s UnlockCardsResponse) GoString() string {
	return s.String()
}

func (s *UnlockCardsResponse) SetHeaders(v map[string]*string) *UnlockCardsResponse {
	s.Headers = v
	return s
}

func (s *UnlockCardsResponse) SetStatusCode(v int32) *UnlockCardsResponse {
	s.StatusCode = &v
	return s
}

func (s *UnlockCardsResponse) SetBody(v *UnlockCardsResponseBody) *UnlockCardsResponse {
	s.Body = v
	return s
}

type UpdateAuthorizationRuleRequest struct {
	AuthorizationRuleId      *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	ClientToken              *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description              *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Destination              *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	DestinationPort          *string `json:"DestinationPort,omitempty" xml:"DestinationPort,omitempty"`
	DryRun                   *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	Name                     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Policy                   *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Protocol                 *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	SourceCidr               *string `json:"SourceCidr,omitempty" xml:"SourceCidr,omitempty"`
	WirelessCloudConnectorId *string `json:"WirelessCloudConnectorId,omitempty" xml:"WirelessCloudConnectorId,omitempty"`
}

func (s UpdateAuthorizationRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAuthorizationRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateAuthorizationRuleRequest) SetAuthorizationRuleId(v string) *UpdateAuthorizationRuleRequest {
	s.AuthorizationRuleId = &v
	return s
}

func (s *UpdateAuthorizationRuleRequest) SetClientToken(v string) *UpdateAuthorizationRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAuthorizationRuleRequest) SetDescription(v string) *UpdateAuthorizationRuleRequest {
	s.Description = &v
	return s
}

func (s *UpdateAuthorizationRuleRequest) SetDestination(v string) *UpdateAuthorizationRuleRequest {
	s.Destination = &v
	return s
}

func (s *UpdateAuthorizationRuleRequest) SetDestinationPort(v string) *UpdateAuthorizationRuleRequest {
	s.DestinationPort = &v
	return s
}

func (s *UpdateAuthorizationRuleRequest) SetDryRun(v bool) *UpdateAuthorizationRuleRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateAuthorizationRuleRequest) SetName(v string) *UpdateAuthorizationRuleRequest {
	s.Name = &v
	return s
}

func (s *UpdateAuthorizationRuleRequest) SetPolicy(v string) *UpdateAuthorizationRuleRequest {
	s.Policy = &v
	return s
}

func (s *UpdateAuthorizationRuleRequest) SetProtocol(v string) *UpdateAuthorizationRuleRequest {
	s.Protocol = &v
	return s
}

func (s *UpdateAuthorizationRuleRequest) SetSourceCidr(v string) *UpdateAuthorizationRuleRequest {
	s.SourceCidr = &v
	return s
}

func (s *UpdateAuthorizationRuleRequest) SetWirelessCloudConnectorId(v string) *UpdateAuthorizationRuleRequest {
	s.WirelessCloudConnectorId = &v
	return s
}

type UpdateAuthorizationRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAuthorizationRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAuthorizationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAuthorizationRuleResponseBody) SetRequestId(v string) *UpdateAuthorizationRuleResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAuthorizationRuleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAuthorizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAuthorizationRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAuthorizationRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateAuthorizationRuleResponse) SetHeaders(v map[string]*string) *UpdateAuthorizationRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateAuthorizationRuleResponse) SetStatusCode(v int32) *UpdateAuthorizationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAuthorizationRuleResponse) SetBody(v *UpdateAuthorizationRuleResponseBody) *UpdateAuthorizationRuleResponse {
	s.Body = v
	return s
}

type UpdateBatchOperateCardsTaskRequest struct {
	BatchOperateCardsTaskId   *string   `json:"BatchOperateCardsTaskId,omitempty" xml:"BatchOperateCardsTaskId,omitempty"`
	ClientToken               *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description               *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	DryRun                    *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EffectType                *string   `json:"EffectType,omitempty" xml:"EffectType,omitempty"`
	Iccids                    []*string `json:"Iccids,omitempty" xml:"Iccids,omitempty" type:"Repeated"`
	IccidsOssFilePath         *string   `json:"IccidsOssFilePath,omitempty" xml:"IccidsOssFilePath,omitempty"`
	Name                      *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	OperateType               *string   `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	RegionId                  *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Threshold                 *int64    `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	WirelessCloudConnectorIds []*string `json:"WirelessCloudConnectorIds,omitempty" xml:"WirelessCloudConnectorIds,omitempty" type:"Repeated"`
}

func (s UpdateBatchOperateCardsTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateBatchOperateCardsTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateBatchOperateCardsTaskRequest) SetBatchOperateCardsTaskId(v string) *UpdateBatchOperateCardsTaskRequest {
	s.BatchOperateCardsTaskId = &v
	return s
}

func (s *UpdateBatchOperateCardsTaskRequest) SetClientToken(v string) *UpdateBatchOperateCardsTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateBatchOperateCardsTaskRequest) SetDescription(v string) *UpdateBatchOperateCardsTaskRequest {
	s.Description = &v
	return s
}

func (s *UpdateBatchOperateCardsTaskRequest) SetDryRun(v bool) *UpdateBatchOperateCardsTaskRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateBatchOperateCardsTaskRequest) SetEffectType(v string) *UpdateBatchOperateCardsTaskRequest {
	s.EffectType = &v
	return s
}

func (s *UpdateBatchOperateCardsTaskRequest) SetIccids(v []*string) *UpdateBatchOperateCardsTaskRequest {
	s.Iccids = v
	return s
}

func (s *UpdateBatchOperateCardsTaskRequest) SetIccidsOssFilePath(v string) *UpdateBatchOperateCardsTaskRequest {
	s.IccidsOssFilePath = &v
	return s
}

func (s *UpdateBatchOperateCardsTaskRequest) SetName(v string) *UpdateBatchOperateCardsTaskRequest {
	s.Name = &v
	return s
}

func (s *UpdateBatchOperateCardsTaskRequest) SetOperateType(v string) *UpdateBatchOperateCardsTaskRequest {
	s.OperateType = &v
	return s
}

func (s *UpdateBatchOperateCardsTaskRequest) SetRegionId(v string) *UpdateBatchOperateCardsTaskRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateBatchOperateCardsTaskRequest) SetThreshold(v int64) *UpdateBatchOperateCardsTaskRequest {
	s.Threshold = &v
	return s
}

func (s *UpdateBatchOperateCardsTaskRequest) SetWirelessCloudConnectorIds(v []*string) *UpdateBatchOperateCardsTaskRequest {
	s.WirelessCloudConnectorIds = v
	return s
}

type UpdateBatchOperateCardsTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateBatchOperateCardsTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateBatchOperateCardsTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBatchOperateCardsTaskResponseBody) SetRequestId(v string) *UpdateBatchOperateCardsTaskResponseBody {
	s.RequestId = &v
	return s
}

type UpdateBatchOperateCardsTaskResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateBatchOperateCardsTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateBatchOperateCardsTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateBatchOperateCardsTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateBatchOperateCardsTaskResponse) SetHeaders(v map[string]*string) *UpdateBatchOperateCardsTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateBatchOperateCardsTaskResponse) SetStatusCode(v int32) *UpdateBatchOperateCardsTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBatchOperateCardsTaskResponse) SetBody(v *UpdateBatchOperateCardsTaskResponseBody) *UpdateBatchOperateCardsTaskResponse {
	s.Body = v
	return s
}

type UpdateCardRequest struct {
	ClientToken              *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description              *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DryRun                   *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	Iccid                    *string `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	Name                     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	WirelessCloudConnectorId *string `json:"WirelessCloudConnectorId,omitempty" xml:"WirelessCloudConnectorId,omitempty"`
}

func (s UpdateCardRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCardRequest) GoString() string {
	return s.String()
}

func (s *UpdateCardRequest) SetClientToken(v string) *UpdateCardRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCardRequest) SetDescription(v string) *UpdateCardRequest {
	s.Description = &v
	return s
}

func (s *UpdateCardRequest) SetDryRun(v bool) *UpdateCardRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateCardRequest) SetIccid(v string) *UpdateCardRequest {
	s.Iccid = &v
	return s
}

func (s *UpdateCardRequest) SetName(v string) *UpdateCardRequest {
	s.Name = &v
	return s
}

func (s *UpdateCardRequest) SetWirelessCloudConnectorId(v string) *UpdateCardRequest {
	s.WirelessCloudConnectorId = &v
	return s
}

type UpdateCardResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCardResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCardResponseBody) SetRequestId(v string) *UpdateCardResponseBody {
	s.RequestId = &v
	return s
}

type UpdateCardResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCardResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCardResponse) GoString() string {
	return s.String()
}

func (s *UpdateCardResponse) SetHeaders(v map[string]*string) *UpdateCardResponse {
	s.Headers = v
	return s
}

func (s *UpdateCardResponse) SetStatusCode(v int32) *UpdateCardResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCardResponse) SetBody(v *UpdateCardResponseBody) *UpdateCardResponse {
	s.Body = v
	return s
}

type UpdateDNSAuthorizationRuleRequest struct {
	AuthorizationRuleId      *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	ClientToken              *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description              *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DestinationIp            *string `json:"DestinationIp,omitempty" xml:"DestinationIp,omitempty"`
	DryRun                   *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	Name                     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SourceDNSIp              *string `json:"SourceDNSIp,omitempty" xml:"SourceDNSIp,omitempty"`
	WirelessCloudConnectorId *string `json:"WirelessCloudConnectorId,omitempty" xml:"WirelessCloudConnectorId,omitempty"`
}

func (s UpdateDNSAuthorizationRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDNSAuthorizationRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateDNSAuthorizationRuleRequest) SetAuthorizationRuleId(v string) *UpdateDNSAuthorizationRuleRequest {
	s.AuthorizationRuleId = &v
	return s
}

func (s *UpdateDNSAuthorizationRuleRequest) SetClientToken(v string) *UpdateDNSAuthorizationRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateDNSAuthorizationRuleRequest) SetDescription(v string) *UpdateDNSAuthorizationRuleRequest {
	s.Description = &v
	return s
}

func (s *UpdateDNSAuthorizationRuleRequest) SetDestinationIp(v string) *UpdateDNSAuthorizationRuleRequest {
	s.DestinationIp = &v
	return s
}

func (s *UpdateDNSAuthorizationRuleRequest) SetDryRun(v bool) *UpdateDNSAuthorizationRuleRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateDNSAuthorizationRuleRequest) SetName(v string) *UpdateDNSAuthorizationRuleRequest {
	s.Name = &v
	return s
}

func (s *UpdateDNSAuthorizationRuleRequest) SetSourceDNSIp(v string) *UpdateDNSAuthorizationRuleRequest {
	s.SourceDNSIp = &v
	return s
}

func (s *UpdateDNSAuthorizationRuleRequest) SetWirelessCloudConnectorId(v string) *UpdateDNSAuthorizationRuleRequest {
	s.WirelessCloudConnectorId = &v
	return s
}

type UpdateDNSAuthorizationRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDNSAuthorizationRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDNSAuthorizationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDNSAuthorizationRuleResponseBody) SetRequestId(v string) *UpdateDNSAuthorizationRuleResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDNSAuthorizationRuleResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateDNSAuthorizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDNSAuthorizationRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDNSAuthorizationRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateDNSAuthorizationRuleResponse) SetHeaders(v map[string]*string) *UpdateDNSAuthorizationRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateDNSAuthorizationRuleResponse) SetStatusCode(v int32) *UpdateDNSAuthorizationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDNSAuthorizationRuleResponse) SetBody(v *UpdateDNSAuthorizationRuleResponseBody) *UpdateDNSAuthorizationRuleResponse {
	s.Body = v
	return s
}

type UpdateGroupAuthorizationRuleRequest struct {
	AuthorizationRuleId           *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	ClientToken                   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description                   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Destination                   *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	DestinationPort               *string `json:"DestinationPort,omitempty" xml:"DestinationPort,omitempty"`
	DryRun                        *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	Name                          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Policy                        *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Protocol                      *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	SourceCidr                    *string `json:"SourceCidr,omitempty" xml:"SourceCidr,omitempty"`
	WirelessCloudConnectorGroupId *string `json:"WirelessCloudConnectorGroupId,omitempty" xml:"WirelessCloudConnectorGroupId,omitempty"`
}

func (s UpdateGroupAuthorizationRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupAuthorizationRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupAuthorizationRuleRequest) SetAuthorizationRuleId(v string) *UpdateGroupAuthorizationRuleRequest {
	s.AuthorizationRuleId = &v
	return s
}

func (s *UpdateGroupAuthorizationRuleRequest) SetClientToken(v string) *UpdateGroupAuthorizationRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateGroupAuthorizationRuleRequest) SetDescription(v string) *UpdateGroupAuthorizationRuleRequest {
	s.Description = &v
	return s
}

func (s *UpdateGroupAuthorizationRuleRequest) SetDestination(v string) *UpdateGroupAuthorizationRuleRequest {
	s.Destination = &v
	return s
}

func (s *UpdateGroupAuthorizationRuleRequest) SetDestinationPort(v string) *UpdateGroupAuthorizationRuleRequest {
	s.DestinationPort = &v
	return s
}

func (s *UpdateGroupAuthorizationRuleRequest) SetDryRun(v bool) *UpdateGroupAuthorizationRuleRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateGroupAuthorizationRuleRequest) SetName(v string) *UpdateGroupAuthorizationRuleRequest {
	s.Name = &v
	return s
}

func (s *UpdateGroupAuthorizationRuleRequest) SetPolicy(v string) *UpdateGroupAuthorizationRuleRequest {
	s.Policy = &v
	return s
}

func (s *UpdateGroupAuthorizationRuleRequest) SetProtocol(v string) *UpdateGroupAuthorizationRuleRequest {
	s.Protocol = &v
	return s
}

func (s *UpdateGroupAuthorizationRuleRequest) SetSourceCidr(v string) *UpdateGroupAuthorizationRuleRequest {
	s.SourceCidr = &v
	return s
}

func (s *UpdateGroupAuthorizationRuleRequest) SetWirelessCloudConnectorGroupId(v string) *UpdateGroupAuthorizationRuleRequest {
	s.WirelessCloudConnectorGroupId = &v
	return s
}

type UpdateGroupAuthorizationRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateGroupAuthorizationRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupAuthorizationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGroupAuthorizationRuleResponseBody) SetRequestId(v string) *UpdateGroupAuthorizationRuleResponseBody {
	s.RequestId = &v
	return s
}

type UpdateGroupAuthorizationRuleResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateGroupAuthorizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateGroupAuthorizationRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupAuthorizationRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateGroupAuthorizationRuleResponse) SetHeaders(v map[string]*string) *UpdateGroupAuthorizationRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateGroupAuthorizationRuleResponse) SetStatusCode(v int32) *UpdateGroupAuthorizationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGroupAuthorizationRuleResponse) SetBody(v *UpdateGroupAuthorizationRuleResponseBody) *UpdateGroupAuthorizationRuleResponse {
	s.Body = v
	return s
}

type UpdateGroupDnsAuthorizationRuleRequest struct {
	AuthorizationRuleId           *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	ClientToken                   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description                   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DestinationIp                 *string `json:"DestinationIp,omitempty" xml:"DestinationIp,omitempty"`
	DryRun                        *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	Name                          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SourceDNSIp                   *string `json:"SourceDNSIp,omitempty" xml:"SourceDNSIp,omitempty"`
	WirelessCloudConnectorGroupId *string `json:"WirelessCloudConnectorGroupId,omitempty" xml:"WirelessCloudConnectorGroupId,omitempty"`
}

func (s UpdateGroupDnsAuthorizationRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupDnsAuthorizationRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupDnsAuthorizationRuleRequest) SetAuthorizationRuleId(v string) *UpdateGroupDnsAuthorizationRuleRequest {
	s.AuthorizationRuleId = &v
	return s
}

func (s *UpdateGroupDnsAuthorizationRuleRequest) SetClientToken(v string) *UpdateGroupDnsAuthorizationRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateGroupDnsAuthorizationRuleRequest) SetDescription(v string) *UpdateGroupDnsAuthorizationRuleRequest {
	s.Description = &v
	return s
}

func (s *UpdateGroupDnsAuthorizationRuleRequest) SetDestinationIp(v string) *UpdateGroupDnsAuthorizationRuleRequest {
	s.DestinationIp = &v
	return s
}

func (s *UpdateGroupDnsAuthorizationRuleRequest) SetDryRun(v bool) *UpdateGroupDnsAuthorizationRuleRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateGroupDnsAuthorizationRuleRequest) SetName(v string) *UpdateGroupDnsAuthorizationRuleRequest {
	s.Name = &v
	return s
}

func (s *UpdateGroupDnsAuthorizationRuleRequest) SetSourceDNSIp(v string) *UpdateGroupDnsAuthorizationRuleRequest {
	s.SourceDNSIp = &v
	return s
}

func (s *UpdateGroupDnsAuthorizationRuleRequest) SetWirelessCloudConnectorGroupId(v string) *UpdateGroupDnsAuthorizationRuleRequest {
	s.WirelessCloudConnectorGroupId = &v
	return s
}

type UpdateGroupDnsAuthorizationRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateGroupDnsAuthorizationRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupDnsAuthorizationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGroupDnsAuthorizationRuleResponseBody) SetRequestId(v string) *UpdateGroupDnsAuthorizationRuleResponseBody {
	s.RequestId = &v
	return s
}

type UpdateGroupDnsAuthorizationRuleResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateGroupDnsAuthorizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateGroupDnsAuthorizationRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupDnsAuthorizationRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateGroupDnsAuthorizationRuleResponse) SetHeaders(v map[string]*string) *UpdateGroupDnsAuthorizationRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateGroupDnsAuthorizationRuleResponse) SetStatusCode(v int32) *UpdateGroupDnsAuthorizationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGroupDnsAuthorizationRuleResponse) SetBody(v *UpdateGroupDnsAuthorizationRuleResponseBody) *UpdateGroupDnsAuthorizationRuleResponse {
	s.Body = v
	return s
}

type UpdateWirelessCloudConnectorRequest struct {
	ClientToken              *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description              *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DryRun                   *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	Name                     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	WirelessCloudConnectorId *string `json:"WirelessCloudConnectorId,omitempty" xml:"WirelessCloudConnectorId,omitempty"`
}

func (s UpdateWirelessCloudConnectorRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWirelessCloudConnectorRequest) GoString() string {
	return s.String()
}

func (s *UpdateWirelessCloudConnectorRequest) SetClientToken(v string) *UpdateWirelessCloudConnectorRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateWirelessCloudConnectorRequest) SetDescription(v string) *UpdateWirelessCloudConnectorRequest {
	s.Description = &v
	return s
}

func (s *UpdateWirelessCloudConnectorRequest) SetDryRun(v bool) *UpdateWirelessCloudConnectorRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateWirelessCloudConnectorRequest) SetName(v string) *UpdateWirelessCloudConnectorRequest {
	s.Name = &v
	return s
}

func (s *UpdateWirelessCloudConnectorRequest) SetWirelessCloudConnectorId(v string) *UpdateWirelessCloudConnectorRequest {
	s.WirelessCloudConnectorId = &v
	return s
}

type UpdateWirelessCloudConnectorResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateWirelessCloudConnectorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateWirelessCloudConnectorResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWirelessCloudConnectorResponseBody) SetRequestId(v string) *UpdateWirelessCloudConnectorResponseBody {
	s.RequestId = &v
	return s
}

type UpdateWirelessCloudConnectorResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateWirelessCloudConnectorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateWirelessCloudConnectorResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWirelessCloudConnectorResponse) GoString() string {
	return s.String()
}

func (s *UpdateWirelessCloudConnectorResponse) SetHeaders(v map[string]*string) *UpdateWirelessCloudConnectorResponse {
	s.Headers = v
	return s
}

func (s *UpdateWirelessCloudConnectorResponse) SetStatusCode(v int32) *UpdateWirelessCloudConnectorResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWirelessCloudConnectorResponse) SetBody(v *UpdateWirelessCloudConnectorResponseBody) *UpdateWirelessCloudConnectorResponse {
	s.Body = v
	return s
}

type UpdateWirelessCloudConnectorGroupRequest struct {
	ClientToken                   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description                   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DryRun                        *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	Name                          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	WirelessCloudConnectorGroupId *string `json:"WirelessCloudConnectorGroupId,omitempty" xml:"WirelessCloudConnectorGroupId,omitempty"`
}

func (s UpdateWirelessCloudConnectorGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWirelessCloudConnectorGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateWirelessCloudConnectorGroupRequest) SetClientToken(v string) *UpdateWirelessCloudConnectorGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateWirelessCloudConnectorGroupRequest) SetDescription(v string) *UpdateWirelessCloudConnectorGroupRequest {
	s.Description = &v
	return s
}

func (s *UpdateWirelessCloudConnectorGroupRequest) SetDryRun(v bool) *UpdateWirelessCloudConnectorGroupRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateWirelessCloudConnectorGroupRequest) SetName(v string) *UpdateWirelessCloudConnectorGroupRequest {
	s.Name = &v
	return s
}

func (s *UpdateWirelessCloudConnectorGroupRequest) SetWirelessCloudConnectorGroupId(v string) *UpdateWirelessCloudConnectorGroupRequest {
	s.WirelessCloudConnectorGroupId = &v
	return s
}

type UpdateWirelessCloudConnectorGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateWirelessCloudConnectorGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateWirelessCloudConnectorGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWirelessCloudConnectorGroupResponseBody) SetRequestId(v string) *UpdateWirelessCloudConnectorGroupResponseBody {
	s.RequestId = &v
	return s
}

type UpdateWirelessCloudConnectorGroupResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateWirelessCloudConnectorGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateWirelessCloudConnectorGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWirelessCloudConnectorGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateWirelessCloudConnectorGroupResponse) SetHeaders(v map[string]*string) *UpdateWirelessCloudConnectorGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateWirelessCloudConnectorGroupResponse) SetStatusCode(v int32) *UpdateWirelessCloudConnectorGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWirelessCloudConnectorGroupResponse) SetBody(v *UpdateWirelessCloudConnectorGroupResponseBody) *UpdateWirelessCloudConnectorGroupResponse {
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
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("cc5g"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddDNSAuthorizationRuleWithOptions(request *AddDNSAuthorizationRuleRequest, runtime *util.RuntimeOptions) (_result *AddDNSAuthorizationRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationIp)) {
		query["DestinationIp"] = request.DestinationIp
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.SourceDNSIp)) {
		query["SourceDNSIp"] = request.SourceDNSIp
	}

	if !tea.BoolValue(util.IsUnset(request.WirelessCloudConnectorId)) {
		query["WirelessCloudConnectorId"] = request.WirelessCloudConnectorId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddDNSAuthorizationRule"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddDNSAuthorizationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddDNSAuthorizationRule(request *AddDNSAuthorizationRuleRequest) (_result *AddDNSAuthorizationRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddDNSAuthorizationRuleResponse{}
	_body, _err := client.AddDNSAuthorizationRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddGroupDnsAuthorizationRuleWithOptions(request *AddGroupDnsAuthorizationRuleRequest, runtime *util.RuntimeOptions) (_result *AddGroupDnsAuthorizationRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationIp)) {
		query["DestinationIp"] = request.DestinationIp
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.SourceDNSIp)) {
		query["SourceDNSIp"] = request.SourceDNSIp
	}

	if !tea.BoolValue(util.IsUnset(request.WirelessCloudConnectorGroupId)) {
		query["WirelessCloudConnectorGroupId"] = request.WirelessCloudConnectorGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddGroupDnsAuthorizationRule"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddGroupDnsAuthorizationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddGroupDnsAuthorizationRule(request *AddGroupDnsAuthorizationRuleRequest) (_result *AddGroupDnsAuthorizationRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddGroupDnsAuthorizationRuleResponse{}
	_body, _err := client.AddGroupDnsAuthorizationRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddWirelessCloudConnectorToGroupWithOptions(request *AddWirelessCloudConnectorToGroupRequest, runtime *util.RuntimeOptions) (_result *AddWirelessCloudConnectorToGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.WirelessCloudConnectorGroupId)) {
		query["WirelessCloudConnectorGroupId"] = request.WirelessCloudConnectorGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.WirelessCloudConnectorIds)) {
		query["WirelessCloudConnectorIds"] = request.WirelessCloudConnectorIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddWirelessCloudConnectorToGroup"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddWirelessCloudConnectorToGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddWirelessCloudConnectorToGroup(request *AddWirelessCloudConnectorToGroupRequest) (_result *AddWirelessCloudConnectorToGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddWirelessCloudConnectorToGroupResponse{}
	_body, _err := client.AddWirelessCloudConnectorToGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AttachVpcToNetLinkWithOptions(request *AttachVpcToNetLinkRequest, runtime *util.RuntimeOptions) (_result *AttachVpcToNetLinkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.NetLinkId)) {
		query["NetLinkId"] = request.NetLinkId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitches)) {
		query["VSwitches"] = request.VSwitches
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.WirelessCloudConnectorId)) {
		query["WirelessCloudConnectorId"] = request.WirelessCloudConnectorId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachVpcToNetLink"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachVpcToNetLinkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttachVpcToNetLink(request *AttachVpcToNetLinkRequest) (_result *AttachVpcToNetLinkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachVpcToNetLinkResponse{}
	_body, _err := client.AttachVpcToNetLinkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAuthorizationRuleWithOptions(request *CreateAuthorizationRuleRequest, runtime *util.RuntimeOptions) (_result *CreateAuthorizationRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Destination)) {
		query["Destination"] = request.Destination
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationPort)) {
		query["DestinationPort"] = request.DestinationPort
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationType)) {
		query["DestinationType"] = request.DestinationType
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Policy)) {
		query["Policy"] = request.Policy
	}

	if !tea.BoolValue(util.IsUnset(request.Protocol)) {
		query["Protocol"] = request.Protocol
	}

	if !tea.BoolValue(util.IsUnset(request.SourceCidr)) {
		query["SourceCidr"] = request.SourceCidr
	}

	if !tea.BoolValue(util.IsUnset(request.WirelessCloudConnectorId)) {
		query["WirelessCloudConnectorId"] = request.WirelessCloudConnectorId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAuthorizationRule"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAuthorizationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAuthorizationRule(request *CreateAuthorizationRuleRequest) (_result *CreateAuthorizationRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAuthorizationRuleResponse{}
	_body, _err := client.CreateAuthorizationRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateBatchOperateCardsTaskWithOptions(request *CreateBatchOperateCardsTaskRequest, runtime *util.RuntimeOptions) (_result *CreateBatchOperateCardsTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EffectType)) {
		query["EffectType"] = request.EffectType
	}

	if !tea.BoolValue(util.IsUnset(request.Iccids)) {
		query["Iccids"] = request.Iccids
	}

	if !tea.BoolValue(util.IsUnset(request.IccidsOssFilePath)) {
		query["IccidsOssFilePath"] = request.IccidsOssFilePath
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OperateType)) {
		query["OperateType"] = request.OperateType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Threshold)) {
		query["Threshold"] = request.Threshold
	}

	if !tea.BoolValue(util.IsUnset(request.WirelessCloudConnectorIds)) {
		query["WirelessCloudConnectorIds"] = request.WirelessCloudConnectorIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateBatchOperateCardsTask"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateBatchOperateCardsTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateBatchOperateCardsTask(request *CreateBatchOperateCardsTaskRequest) (_result *CreateBatchOperateCardsTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateBatchOperateCardsTaskResponse{}
	_body, _err := client.CreateBatchOperateCardsTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGroupAuthorizationRuleWithOptions(request *CreateGroupAuthorizationRuleRequest, runtime *util.RuntimeOptions) (_result *CreateGroupAuthorizationRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Destination)) {
		query["Destination"] = request.Destination
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationPort)) {
		query["DestinationPort"] = request.DestinationPort
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationType)) {
		query["DestinationType"] = request.DestinationType
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Policy)) {
		query["Policy"] = request.Policy
	}

	if !tea.BoolValue(util.IsUnset(request.Protocol)) {
		query["Protocol"] = request.Protocol
	}

	if !tea.BoolValue(util.IsUnset(request.SourceCidr)) {
		query["SourceCidr"] = request.SourceCidr
	}

	if !tea.BoolValue(util.IsUnset(request.WirelessCloudConnectorGroupId)) {
		query["WirelessCloudConnectorGroupId"] = request.WirelessCloudConnectorGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGroupAuthorizationRule"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGroupAuthorizationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGroupAuthorizationRule(request *CreateGroupAuthorizationRuleRequest) (_result *CreateGroupAuthorizationRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGroupAuthorizationRuleResponse{}
	_body, _err := client.CreateGroupAuthorizationRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateIoTCloudConnectorBackhaulRouteWithOptions(request *CreateIoTCloudConnectorBackhaulRouteRequest, runtime *util.RuntimeOptions) (_result *CreateIoTCloudConnectorBackhaulRouteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.NetLinkId)) {
		query["NetLinkId"] = request.NetLinkId
	}

	if !tea.BoolValue(util.IsUnset(request.WirelessCloudConnectorId)) {
		query["WirelessCloudConnectorId"] = request.WirelessCloudConnectorId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateIoTCloudConnectorBackhaulRoute"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateIoTCloudConnectorBackhaulRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateIoTCloudConnectorBackhaulRoute(request *CreateIoTCloudConnectorBackhaulRouteRequest) (_result *CreateIoTCloudConnectorBackhaulRouteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateIoTCloudConnectorBackhaulRouteResponse{}
	_body, _err := client.CreateIoTCloudConnectorBackhaulRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateWirelessCloudConnectorWithOptions(request *CreateWirelessCloudConnectorRequest, runtime *util.RuntimeOptions) (_result *CreateWirelessCloudConnectorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BusinessType)) {
		query["BusinessType"] = request.BusinessType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ISP)) {
		query["ISP"] = request.ISP
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NetLinks)) {
		query["NetLinks"] = request.NetLinks
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UseCase)) {
		query["UseCase"] = request.UseCase
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateWirelessCloudConnector"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateWirelessCloudConnectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateWirelessCloudConnector(request *CreateWirelessCloudConnectorRequest) (_result *CreateWirelessCloudConnectorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateWirelessCloudConnectorResponse{}
	_body, _err := client.CreateWirelessCloudConnectorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateWirelessCloudConnectorGroupWithOptions(request *CreateWirelessCloudConnectorGroupRequest, runtime *util.RuntimeOptions) (_result *CreateWirelessCloudConnectorGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateWirelessCloudConnectorGroup"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateWirelessCloudConnectorGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateWirelessCloudConnectorGroup(request *CreateWirelessCloudConnectorGroupRequest) (_result *CreateWirelessCloudConnectorGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateWirelessCloudConnectorGroupResponse{}
	_body, _err := client.CreateWirelessCloudConnectorGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAuthorizationRuleWithOptions(request *DeleteAuthorizationRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteAuthorizationRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthorizationRuleId)) {
		query["AuthorizationRuleId"] = request.AuthorizationRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.WirelessCloudConnectorId)) {
		query["WirelessCloudConnectorId"] = request.WirelessCloudConnectorId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAuthorizationRule"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAuthorizationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAuthorizationRule(request *DeleteAuthorizationRuleRequest) (_result *DeleteAuthorizationRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAuthorizationRuleResponse{}
	_body, _err := client.DeleteAuthorizationRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBatchOperateCardsTaskWithOptions(request *DeleteBatchOperateCardsTaskRequest, runtime *util.RuntimeOptions) (_result *DeleteBatchOperateCardsTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BatchOperateCardsTaskId)) {
		query["BatchOperateCardsTaskId"] = request.BatchOperateCardsTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteBatchOperateCardsTask"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteBatchOperateCardsTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBatchOperateCardsTask(request *DeleteBatchOperateCardsTaskRequest) (_result *DeleteBatchOperateCardsTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteBatchOperateCardsTaskResponse{}
	_body, _err := client.DeleteBatchOperateCardsTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteGroupAuthorizationRuleWithOptions(request *DeleteGroupAuthorizationRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteGroupAuthorizationRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthorizationRuleId)) {
		query["AuthorizationRuleId"] = request.AuthorizationRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.WirelessCloudConnectorGroupId)) {
		query["WirelessCloudConnectorGroupId"] = request.WirelessCloudConnectorGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGroupAuthorizationRule"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGroupAuthorizationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGroupAuthorizationRule(request *DeleteGroupAuthorizationRuleRequest) (_result *DeleteGroupAuthorizationRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGroupAuthorizationRuleResponse{}
	_body, _err := client.DeleteGroupAuthorizationRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteIoTCloudConnectorBackhaulRouteWithOptions(request *DeleteIoTCloudConnectorBackhaulRouteRequest, runtime *util.RuntimeOptions) (_result *DeleteIoTCloudConnectorBackhaulRouteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.NetLinkId)) {
		query["NetLinkId"] = request.NetLinkId
	}

	if !tea.BoolValue(util.IsUnset(request.WirelessCloudConnectorId)) {
		query["WirelessCloudConnectorId"] = request.WirelessCloudConnectorId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteIoTCloudConnectorBackhaulRoute"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteIoTCloudConnectorBackhaulRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteIoTCloudConnectorBackhaulRoute(request *DeleteIoTCloudConnectorBackhaulRouteRequest) (_result *DeleteIoTCloudConnectorBackhaulRouteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteIoTCloudConnectorBackhaulRouteResponse{}
	_body, _err := client.DeleteIoTCloudConnectorBackhaulRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteWirelessCloudConnectorWithOptions(request *DeleteWirelessCloudConnectorRequest, runtime *util.RuntimeOptions) (_result *DeleteWirelessCloudConnectorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.WirelessCloudConnectorId)) {
		query["WirelessCloudConnectorId"] = request.WirelessCloudConnectorId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteWirelessCloudConnector"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteWirelessCloudConnectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteWirelessCloudConnector(request *DeleteWirelessCloudConnectorRequest) (_result *DeleteWirelessCloudConnectorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteWirelessCloudConnectorResponse{}
	_body, _err := client.DeleteWirelessCloudConnectorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteWirelessCloudConnectorGroupWithOptions(request *DeleteWirelessCloudConnectorGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteWirelessCloudConnectorGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.WirelessCloudConnectorGroupId)) {
		query["WirelessCloudConnectorGroupId"] = request.WirelessCloudConnectorGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteWirelessCloudConnectorGroup"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteWirelessCloudConnectorGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteWirelessCloudConnectorGroup(request *DeleteWirelessCloudConnectorGroupRequest) (_result *DeleteWirelessCloudConnectorGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteWirelessCloudConnectorGroupResponse{}
	_body, _err := client.DeleteWirelessCloudConnectorGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetachVpcFromNetLinkWithOptions(request *DetachVpcFromNetLinkRequest, runtime *util.RuntimeOptions) (_result *DetachVpcFromNetLinkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.NetLinkId)) {
		query["NetLinkId"] = request.NetLinkId
	}

	if !tea.BoolValue(util.IsUnset(request.WirelessCloudConnectorId)) {
		query["WirelessCloudConnectorId"] = request.WirelessCloudConnectorId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachVpcFromNetLink"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachVpcFromNetLinkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetachVpcFromNetLink(request *DetachVpcFromNetLinkRequest) (_result *DetachVpcFromNetLinkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachVpcFromNetLinkResponse{}
	_body, _err := client.DetachVpcFromNetLinkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FailCardsWithOptions(request *FailCardsRequest, runtime *util.RuntimeOptions) (_result *FailCardsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.Iccids)) {
		query["Iccids"] = request.Iccids
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FailCards"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FailCardsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FailCards(request *FailCardsRequest) (_result *FailCardsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FailCardsResponse{}
	_body, _err := client.FailCardsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCardWithOptions(request *GetCardRequest, runtime *util.RuntimeOptions) (_result *GetCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCard"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCard(request *GetCardRequest) (_result *GetCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCardResponse{}
	_body, _err := client.GetCardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCardLockReasonWithOptions(request *GetCardLockReasonRequest, runtime *util.RuntimeOptions) (_result *GetCardLockReasonResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCardLockReason"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCardLockReasonResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCardLockReason(request *GetCardLockReasonRequest) (_result *GetCardLockReasonResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCardLockReasonResponse{}
	_body, _err := client.GetCardLockReasonWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCreateCustomerInformationWithOptions(request *GetCreateCustomerInformationRequest, runtime *util.RuntimeOptions) (_result *GetCreateCustomerInformationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCreateCustomerInformation"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCreateCustomerInformationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCreateCustomerInformation(request *GetCreateCustomerInformationRequest) (_result *GetCreateCustomerInformationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCreateCustomerInformationResponse{}
	_body, _err := client.GetCreateCustomerInformationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDiagnoseResultForSingleCardWithOptions(request *GetDiagnoseResultForSingleCardRequest, runtime *util.RuntimeOptions) (_result *GetDiagnoseResultForSingleCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DiagnoseTaskId)) {
		query["DiagnoseTaskId"] = request.DiagnoseTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionNo)) {
		query["RegionNo"] = request.RegionNo
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDiagnoseResultForSingleCard"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDiagnoseResultForSingleCardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDiagnoseResultForSingleCard(request *GetDiagnoseResultForSingleCardRequest) (_result *GetDiagnoseResultForSingleCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDiagnoseResultForSingleCardResponse{}
	_body, _err := client.GetDiagnoseResultForSingleCardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWirelessCloudConnectorWithOptions(request *GetWirelessCloudConnectorRequest, runtime *util.RuntimeOptions) (_result *GetWirelessCloudConnectorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWirelessCloudConnector"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWirelessCloudConnectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWirelessCloudConnector(request *GetWirelessCloudConnectorRequest) (_result *GetWirelessCloudConnectorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWirelessCloudConnectorResponse{}
	_body, _err := client.GetWirelessCloudConnectorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GrantNetLinkWithOptions(request *GrantNetLinkRequest, runtime *util.RuntimeOptions) (_result *GrantNetLinkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.GrantAliUid)) {
		query["GrantAliUid"] = request.GrantAliUid
	}

	if !tea.BoolValue(util.IsUnset(request.NetLinkId)) {
		query["NetLinkId"] = request.NetLinkId
	}

	if !tea.BoolValue(util.IsUnset(request.WirelessCloudConnectorId)) {
		query["WirelessCloudConnectorId"] = request.WirelessCloudConnectorId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GrantNetLink"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GrantNetLinkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GrantNetLink(request *GrantNetLinkRequest) (_result *GrantNetLinkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GrantNetLinkResponse{}
	_body, _err := client.GrantNetLinkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAPNsWithOptions(request *ListAPNsRequest, runtime *util.RuntimeOptions) (_result *ListAPNsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAPNs"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAPNsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAPNs(request *ListAPNsRequest) (_result *ListAPNsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAPNsResponse{}
	_body, _err := client.ListAPNsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAuthorizationRulesWithOptions(request *ListAuthorizationRulesRequest, runtime *util.RuntimeOptions) (_result *ListAuthorizationRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAuthorizationRules"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAuthorizationRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAuthorizationRules(request *ListAuthorizationRulesRequest) (_result *ListAuthorizationRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAuthorizationRulesResponse{}
	_body, _err := client.ListAuthorizationRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListBatchOperateCardsTasksWithOptions(request *ListBatchOperateCardsTasksRequest, runtime *util.RuntimeOptions) (_result *ListBatchOperateCardsTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListBatchOperateCardsTasks"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListBatchOperateCardsTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListBatchOperateCardsTasks(request *ListBatchOperateCardsTasksRequest) (_result *ListBatchOperateCardsTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListBatchOperateCardsTasksResponse{}
	_body, _err := client.ListBatchOperateCardsTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCardUsagesWithOptions(request *ListCardUsagesRequest, runtime *util.RuntimeOptions) (_result *ListCardUsagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCardUsages"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCardUsagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCardUsages(request *ListCardUsagesRequest) (_result *ListCardUsagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCardUsagesResponse{}
	_body, _err := client.ListCardUsagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCardsWithOptions(request *ListCardsRequest, runtime *util.RuntimeOptions) (_result *ListCardsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCards"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCardsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCards(request *ListCardsRequest) (_result *ListCardsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCardsResponse{}
	_body, _err := client.ListCardsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDataPackagesWithOptions(request *ListDataPackagesRequest, runtime *util.RuntimeOptions) (_result *ListDataPackagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDataPackages"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDataPackagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDataPackages(request *ListDataPackagesRequest) (_result *ListDataPackagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDataPackagesResponse{}
	_body, _err := client.ListDataPackagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDiagnoseInfoForSingleCardWithOptions(request *ListDiagnoseInfoForSingleCardRequest, runtime *util.RuntimeOptions) (_result *ListDiagnoseInfoForSingleCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionNo)) {
		query["RegionNo"] = request.RegionNo
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.WirelessCloudConnectorId)) {
		query["WirelessCloudConnectorId"] = request.WirelessCloudConnectorId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDiagnoseInfoForSingleCard"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDiagnoseInfoForSingleCardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDiagnoseInfoForSingleCard(request *ListDiagnoseInfoForSingleCardRequest) (_result *ListDiagnoseInfoForSingleCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDiagnoseInfoForSingleCardResponse{}
	_body, _err := client.ListDiagnoseInfoForSingleCardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListGroupAuthorizationRulesWithOptions(request *ListGroupAuthorizationRulesRequest, runtime *util.RuntimeOptions) (_result *ListGroupAuthorizationRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGroupAuthorizationRules"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGroupAuthorizationRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListGroupAuthorizationRules(request *ListGroupAuthorizationRulesRequest) (_result *ListGroupAuthorizationRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListGroupAuthorizationRulesResponse{}
	_body, _err := client.ListGroupAuthorizationRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListIoTCloudConnectorBackhaulRouteWithOptions(request *ListIoTCloudConnectorBackhaulRouteRequest, runtime *util.RuntimeOptions) (_result *ListIoTCloudConnectorBackhaulRouteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListIoTCloudConnectorBackhaulRoute"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIoTCloudConnectorBackhaulRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListIoTCloudConnectorBackhaulRoute(request *ListIoTCloudConnectorBackhaulRouteRequest) (_result *ListIoTCloudConnectorBackhaulRouteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListIoTCloudConnectorBackhaulRouteResponse{}
	_body, _err := client.ListIoTCloudConnectorBackhaulRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOrdersWithOptions(request *ListOrdersRequest, runtime *util.RuntimeOptions) (_result *ListOrdersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListOrders"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOrdersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOrders(request *ListOrdersRequest) (_result *ListOrdersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOrdersResponse{}
	_body, _err := client.ListOrdersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRegionsWithOptions(request *ListRegionsRequest, runtime *util.RuntimeOptions) (_result *ListRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRegions"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRegions(request *ListRegionsRequest) (_result *ListRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRegionsResponse{}
	_body, _err := client.ListRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListWirelessCloudConnectorGroupsWithOptions(request *ListWirelessCloudConnectorGroupsRequest, runtime *util.RuntimeOptions) (_result *ListWirelessCloudConnectorGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListWirelessCloudConnectorGroups"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListWirelessCloudConnectorGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListWirelessCloudConnectorGroups(request *ListWirelessCloudConnectorGroupsRequest) (_result *ListWirelessCloudConnectorGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListWirelessCloudConnectorGroupsResponse{}
	_body, _err := client.ListWirelessCloudConnectorGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListWirelessCloudConnectorsWithOptions(request *ListWirelessCloudConnectorsRequest, runtime *util.RuntimeOptions) (_result *ListWirelessCloudConnectorsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListWirelessCloudConnectors"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListWirelessCloudConnectorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListWirelessCloudConnectors(request *ListWirelessCloudConnectorsRequest) (_result *ListWirelessCloudConnectorsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListWirelessCloudConnectorsResponse{}
	_body, _err := client.ListWirelessCloudConnectorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListZonesWithOptions(request *ListZonesRequest, runtime *util.RuntimeOptions) (_result *ListZonesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListZones"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListZones(request *ListZonesRequest) (_result *ListZonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListZonesResponse{}
	_body, _err := client.ListZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LockCardsWithOptions(request *LockCardsRequest, runtime *util.RuntimeOptions) (_result *LockCardsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.Iccids)) {
		query["Iccids"] = request.Iccids
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("LockCards"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &LockCardsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LockCards(request *LockCardsRequest) (_result *LockCardsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LockCardsResponse{}
	_body, _err := client.LockCardsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyWirelessCloudConnectorFeatureWithOptions(request *ModifyWirelessCloudConnectorFeatureRequest, runtime *util.RuntimeOptions) (_result *ModifyWirelessCloudConnectorFeatureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.FeatureName)) {
		query["FeatureName"] = request.FeatureName
	}

	if !tea.BoolValue(util.IsUnset(request.FeatureValue)) {
		query["FeatureValue"] = request.FeatureValue
	}

	if !tea.BoolValue(util.IsUnset(request.WirelessCloudConnectorId)) {
		query["WirelessCloudConnectorId"] = request.WirelessCloudConnectorId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyWirelessCloudConnectorFeature"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyWirelessCloudConnectorFeatureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyWirelessCloudConnectorFeature(request *ModifyWirelessCloudConnectorFeatureRequest) (_result *ModifyWirelessCloudConnectorFeatureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyWirelessCloudConnectorFeatureResponse{}
	_body, _err := client.ModifyWirelessCloudConnectorFeatureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenCc5gServiceWithOptions(request *OpenCc5gServiceRequest, runtime *util.RuntimeOptions) (_result *OpenCc5gServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OpenCc5gService"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenCc5gServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenCc5gService(request *OpenCc5gServiceRequest) (_result *OpenCc5gServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenCc5gServiceResponse{}
	_body, _err := client.OpenCc5gServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RebindCardsWithOptions(request *RebindCardsRequest, runtime *util.RuntimeOptions) (_result *RebindCardsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.Iccids)) {
		query["Iccids"] = request.Iccids
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RebindCards"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RebindCardsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RebindCards(request *RebindCardsRequest) (_result *RebindCardsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RebindCardsResponse{}
	_body, _err := client.RebindCardsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveWirelessCloudConnectorFromGroupWithOptions(request *RemoveWirelessCloudConnectorFromGroupRequest, runtime *util.RuntimeOptions) (_result *RemoveWirelessCloudConnectorFromGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.WirelessCloudConnectorGroupId)) {
		query["WirelessCloudConnectorGroupId"] = request.WirelessCloudConnectorGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.WirelessCloudConnectorIds)) {
		query["WirelessCloudConnectorIds"] = request.WirelessCloudConnectorIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveWirelessCloudConnectorFromGroup"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveWirelessCloudConnectorFromGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveWirelessCloudConnectorFromGroup(request *RemoveWirelessCloudConnectorFromGroupRequest) (_result *RemoveWirelessCloudConnectorFromGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveWirelessCloudConnectorFromGroupResponse{}
	_body, _err := client.RemoveWirelessCloudConnectorFromGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResumeCardsWithOptions(request *ResumeCardsRequest, runtime *util.RuntimeOptions) (_result *ResumeCardsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.Iccids)) {
		query["Iccids"] = request.Iccids
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResumeCards"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResumeCardsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResumeCards(request *ResumeCardsRequest) (_result *ResumeCardsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResumeCardsResponse{}
	_body, _err := client.ResumeCardsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RevokeNetLinkWithOptions(request *RevokeNetLinkRequest, runtime *util.RuntimeOptions) (_result *RevokeNetLinkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.NetLinkId)) {
		query["NetLinkId"] = request.NetLinkId
	}

	if !tea.BoolValue(util.IsUnset(request.WirelessCloudConnectorId)) {
		query["WirelessCloudConnectorId"] = request.WirelessCloudConnectorId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RevokeNetLink"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RevokeNetLinkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RevokeNetLink(request *RevokeNetLinkRequest) (_result *RevokeNetLinkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RevokeNetLinkResponse{}
	_body, _err := client.RevokeNetLinkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopCardsWithOptions(request *StopCardsRequest, runtime *util.RuntimeOptions) (_result *StopCardsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.Iccids)) {
		query["Iccids"] = request.Iccids
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopCards"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopCardsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopCards(request *StopCardsRequest) (_result *StopCardsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopCardsResponse{}
	_body, _err := client.StopCardsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitDiagnoseTaskForSingleCardWithOptions(request *SubmitDiagnoseTaskForSingleCardRequest, runtime *util.RuntimeOptions) (_result *SubmitDiagnoseTaskForSingleCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginTime)) {
		query["BeginTime"] = request.BeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.Destination)) {
		query["Destination"] = request.Destination
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.RegionNo)) {
		query["RegionNo"] = request.RegionNo
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceUid)) {
		query["ResourceUid"] = request.ResourceUid
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.WirelessCloudConnectorId)) {
		query["WirelessCloudConnectorId"] = request.WirelessCloudConnectorId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitDiagnoseTaskForSingleCard"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitDiagnoseTaskForSingleCardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitDiagnoseTaskForSingleCard(request *SubmitDiagnoseTaskForSingleCardRequest) (_result *SubmitDiagnoseTaskForSingleCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitDiagnoseTaskForSingleCardResponse{}
	_body, _err := client.SubmitDiagnoseTaskForSingleCardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SwitchWirelessCloudConnectorToBusinessWithOptions(request *SwitchWirelessCloudConnectorToBusinessRequest, runtime *util.RuntimeOptions) (_result *SwitchWirelessCloudConnectorToBusinessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.WirelessCloudConnectorId)) {
		query["WirelessCloudConnectorId"] = request.WirelessCloudConnectorId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SwitchWirelessCloudConnectorToBusiness"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SwitchWirelessCloudConnectorToBusinessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SwitchWirelessCloudConnectorToBusiness(request *SwitchWirelessCloudConnectorToBusinessRequest) (_result *SwitchWirelessCloudConnectorToBusinessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SwitchWirelessCloudConnectorToBusinessResponse{}
	_body, _err := client.SwitchWirelessCloudConnectorToBusinessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnlockCardsWithOptions(request *UnlockCardsRequest, runtime *util.RuntimeOptions) (_result *UnlockCardsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.Iccids)) {
		query["Iccids"] = request.Iccids
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnlockCards"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnlockCardsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnlockCards(request *UnlockCardsRequest) (_result *UnlockCardsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnlockCardsResponse{}
	_body, _err := client.UnlockCardsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAuthorizationRuleWithOptions(request *UpdateAuthorizationRuleRequest, runtime *util.RuntimeOptions) (_result *UpdateAuthorizationRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthorizationRuleId)) {
		query["AuthorizationRuleId"] = request.AuthorizationRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Destination)) {
		query["Destination"] = request.Destination
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationPort)) {
		query["DestinationPort"] = request.DestinationPort
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Policy)) {
		query["Policy"] = request.Policy
	}

	if !tea.BoolValue(util.IsUnset(request.Protocol)) {
		query["Protocol"] = request.Protocol
	}

	if !tea.BoolValue(util.IsUnset(request.SourceCidr)) {
		query["SourceCidr"] = request.SourceCidr
	}

	if !tea.BoolValue(util.IsUnset(request.WirelessCloudConnectorId)) {
		query["WirelessCloudConnectorId"] = request.WirelessCloudConnectorId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAuthorizationRule"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAuthorizationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAuthorizationRule(request *UpdateAuthorizationRuleRequest) (_result *UpdateAuthorizationRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAuthorizationRuleResponse{}
	_body, _err := client.UpdateAuthorizationRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateBatchOperateCardsTaskWithOptions(request *UpdateBatchOperateCardsTaskRequest, runtime *util.RuntimeOptions) (_result *UpdateBatchOperateCardsTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BatchOperateCardsTaskId)) {
		query["BatchOperateCardsTaskId"] = request.BatchOperateCardsTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EffectType)) {
		query["EffectType"] = request.EffectType
	}

	if !tea.BoolValue(util.IsUnset(request.Iccids)) {
		query["Iccids"] = request.Iccids
	}

	if !tea.BoolValue(util.IsUnset(request.IccidsOssFilePath)) {
		query["IccidsOssFilePath"] = request.IccidsOssFilePath
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OperateType)) {
		query["OperateType"] = request.OperateType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Threshold)) {
		query["Threshold"] = request.Threshold
	}

	if !tea.BoolValue(util.IsUnset(request.WirelessCloudConnectorIds)) {
		query["WirelessCloudConnectorIds"] = request.WirelessCloudConnectorIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateBatchOperateCardsTask"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateBatchOperateCardsTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateBatchOperateCardsTask(request *UpdateBatchOperateCardsTaskRequest) (_result *UpdateBatchOperateCardsTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateBatchOperateCardsTaskResponse{}
	_body, _err := client.UpdateBatchOperateCardsTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCardWithOptions(request *UpdateCardRequest, runtime *util.RuntimeOptions) (_result *UpdateCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.Iccid)) {
		query["Iccid"] = request.Iccid
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.WirelessCloudConnectorId)) {
		query["WirelessCloudConnectorId"] = request.WirelessCloudConnectorId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateCard"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateCardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCard(request *UpdateCardRequest) (_result *UpdateCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCardResponse{}
	_body, _err := client.UpdateCardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDNSAuthorizationRuleWithOptions(request *UpdateDNSAuthorizationRuleRequest, runtime *util.RuntimeOptions) (_result *UpdateDNSAuthorizationRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthorizationRuleId)) {
		query["AuthorizationRuleId"] = request.AuthorizationRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationIp)) {
		query["DestinationIp"] = request.DestinationIp
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.SourceDNSIp)) {
		query["SourceDNSIp"] = request.SourceDNSIp
	}

	if !tea.BoolValue(util.IsUnset(request.WirelessCloudConnectorId)) {
		query["WirelessCloudConnectorId"] = request.WirelessCloudConnectorId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDNSAuthorizationRule"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDNSAuthorizationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDNSAuthorizationRule(request *UpdateDNSAuthorizationRuleRequest) (_result *UpdateDNSAuthorizationRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDNSAuthorizationRuleResponse{}
	_body, _err := client.UpdateDNSAuthorizationRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateGroupAuthorizationRuleWithOptions(request *UpdateGroupAuthorizationRuleRequest, runtime *util.RuntimeOptions) (_result *UpdateGroupAuthorizationRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthorizationRuleId)) {
		query["AuthorizationRuleId"] = request.AuthorizationRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Destination)) {
		query["Destination"] = request.Destination
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationPort)) {
		query["DestinationPort"] = request.DestinationPort
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Policy)) {
		query["Policy"] = request.Policy
	}

	if !tea.BoolValue(util.IsUnset(request.Protocol)) {
		query["Protocol"] = request.Protocol
	}

	if !tea.BoolValue(util.IsUnset(request.SourceCidr)) {
		query["SourceCidr"] = request.SourceCidr
	}

	if !tea.BoolValue(util.IsUnset(request.WirelessCloudConnectorGroupId)) {
		query["WirelessCloudConnectorGroupId"] = request.WirelessCloudConnectorGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateGroupAuthorizationRule"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateGroupAuthorizationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateGroupAuthorizationRule(request *UpdateGroupAuthorizationRuleRequest) (_result *UpdateGroupAuthorizationRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateGroupAuthorizationRuleResponse{}
	_body, _err := client.UpdateGroupAuthorizationRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateGroupDnsAuthorizationRuleWithOptions(request *UpdateGroupDnsAuthorizationRuleRequest, runtime *util.RuntimeOptions) (_result *UpdateGroupDnsAuthorizationRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthorizationRuleId)) {
		query["AuthorizationRuleId"] = request.AuthorizationRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationIp)) {
		query["DestinationIp"] = request.DestinationIp
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.SourceDNSIp)) {
		query["SourceDNSIp"] = request.SourceDNSIp
	}

	if !tea.BoolValue(util.IsUnset(request.WirelessCloudConnectorGroupId)) {
		query["WirelessCloudConnectorGroupId"] = request.WirelessCloudConnectorGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateGroupDnsAuthorizationRule"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateGroupDnsAuthorizationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateGroupDnsAuthorizationRule(request *UpdateGroupDnsAuthorizationRuleRequest) (_result *UpdateGroupDnsAuthorizationRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateGroupDnsAuthorizationRuleResponse{}
	_body, _err := client.UpdateGroupDnsAuthorizationRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateWirelessCloudConnectorWithOptions(request *UpdateWirelessCloudConnectorRequest, runtime *util.RuntimeOptions) (_result *UpdateWirelessCloudConnectorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.WirelessCloudConnectorId)) {
		query["WirelessCloudConnectorId"] = request.WirelessCloudConnectorId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateWirelessCloudConnector"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateWirelessCloudConnectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateWirelessCloudConnector(request *UpdateWirelessCloudConnectorRequest) (_result *UpdateWirelessCloudConnectorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateWirelessCloudConnectorResponse{}
	_body, _err := client.UpdateWirelessCloudConnectorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateWirelessCloudConnectorGroupWithOptions(request *UpdateWirelessCloudConnectorGroupRequest, runtime *util.RuntimeOptions) (_result *UpdateWirelessCloudConnectorGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.WirelessCloudConnectorGroupId)) {
		query["WirelessCloudConnectorGroupId"] = request.WirelessCloudConnectorGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateWirelessCloudConnectorGroup"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateWirelessCloudConnectorGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateWirelessCloudConnectorGroup(request *UpdateWirelessCloudConnectorGroupRequest) (_result *UpdateWirelessCloudConnectorGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateWirelessCloudConnectorGroupResponse{}
	_body, _err := client.UpdateWirelessCloudConnectorGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
