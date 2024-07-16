// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddDNSAuthorizationRuleRequest struct {
	// example:
	//
	// TF-******-1633255280-43c94bf7-2dd3-4c14-8
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// ar5g-test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1.1.1.1
	DestinationIp *string `json:"DestinationIp,omitempty" xml:"DestinationIp,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.0.0.0
	SourceDNSIp *string `json:"SourceDNSIp,omitempty" xml:"SourceDNSIp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cc5g-xxxx
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
	// example:
	//
	// ar5g-xxxx
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	// example:
	//
	// 1651FBB6-4FBF-49FF-A9F5-DF5D696C7EC6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDNSAuthorizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// TF-******-1633255280-43c94bf7-2dd3-4c14-8
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// ar5gg-test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1.1.1.1
	DestinationIp *string `json:"DestinationIp,omitempty" xml:"DestinationIp,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 114.114.114.114
	SourceDNSIp *string `json:"SourceDNSIp,omitempty" xml:"SourceDNSIp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cc5gg-****
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
	// example:
	//
	// ar5gg-12fnojjtkkre28t08j
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	// example:
	//
	// D3BAFE75-C670-52DB-9739-03133E5FC914
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddGroupDnsAuthorizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// TF-******-1633255280-43c94bf7-2dd3-4c14-8
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cc5gg-****
	WirelessCloudConnectorGroupId *string `json:"WirelessCloudConnectorGroupId,omitempty" xml:"WirelessCloudConnectorGroupId,omitempty"`
	// This parameter is required.
	WirelessCloudConnectorIds []*string `json:"WirelessCloudConnectorIds,omitempty" xml:"WirelessCloudConnectorIds,omitempty" type:"Repeated"`
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
	// example:
	//
	// 1651FBB6-4FBF-49FF-A9F5-DF5D696C7EC6
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
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddWirelessCloudConnectorToGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// TF-******-1633255280-43c94bf7-2dd3-4c14-8
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cciot-xxxx
	NetLinkId *string `json:"NetLinkId,omitempty" xml:"NetLinkId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	VSwitches []*string `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1etmgc4vooo5ahbkzou
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cc5g-xxxx
	WirelessCloudConnectorId *string `json:"WirelessCloudConnectorId,omitempty" xml:"WirelessCloudConnectorId,omitempty"`
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
	// example:
	//
	// 1651FBB6-4FBF-49FF-A9F5-DF5D696C7EC6
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachVpcToNetLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// TF-******-1633255280-43c94bf7-2dd3-4c14-8
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// ar5g-test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10.0.0.1
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// example:
	//
	// -1/-1
	DestinationPort *string `json:"DestinationPort,omitempty" xml:"DestinationPort,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Cidr
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// Test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Permit
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// all
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// 0.0.0.0/0
	SourceCidr *string `json:"SourceCidr,omitempty" xml:"SourceCidr,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cc5g-xxxx
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
	// example:
	//
	// ar5g-12fnojjtkkre28t08j
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	// example:
	//
	// 61202DF5-668B-5EDA-8177-F2D71F4737C1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAuthorizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// TF-******-1633255280-43c94bf7-2dd3-4c14-8
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// task- test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CurrentMonth
	EffectType *string   `json:"EffectType,omitempty" xml:"EffectType,omitempty"`
	Iccids     []*string `json:"Iccids,omitempty" xml:"Iccids,omitempty" type:"Repeated"`
	// example:
	//
	// https://examplebucket.oss-eu-central-1.aliyuncs.com/***
	IccidsOssFilePath *string `json:"IccidsOssFilePath,omitempty" xml:"IccidsOssFilePath,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// BreakNetwork
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
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
	// example:
	//
	// boct-xxxx
	BatchOperateCardsTaskId *string `json:"BatchOperateCardsTaskId,omitempty" xml:"BatchOperateCardsTaskId,omitempty"`
	// example:
	//
	// https://examplebucket.oss-eu-central-1.aliyuncs.com/***
	OperateResultOssFilePath *string `json:"OperateResultOssFilePath,omitempty" xml:"OperateResultOssFilePath,omitempty"`
	// example:
	//
	// 1651FBB6-4FBF-49FF-A9F5-DF5D696C7EC6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBatchOperateCardsTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// TF-******-1633255280-43c94bf7-2dd3-4c14-8
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// ar5gg-test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 106.14.175.4
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// example:
	//
	// -1/-1
	DestinationPort *string `json:"DestinationPort,omitempty" xml:"DestinationPort,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Cidr
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Permit
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// all
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// 0.0.0.0/0
	SourceCidr *string `json:"SourceCidr,omitempty" xml:"SourceCidr,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cc5gg-****
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
	// example:
	//
	// ar5gg-xxxx
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	// example:
	//
	// AA93A076-3D94-51AB-AC70-4C94AF003B03
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGroupAuthorizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// TF-******-1633255280-43c94bf7-2dd3-4c14-8
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// cciot-xxxx
	NetLinkId *string `json:"NetLinkId,omitempty" xml:"NetLinkId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cc5g-xxxx
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
	// example:
	//
	// 1651FBB6-4FBF-49FF-A9F5-DF5D696C7EC6
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
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIoTCloudConnectorBackhaulRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// Business
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// example:
	//
	// TF-******-1633255280-43c94bf7-2dd3-4c14-8
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// unicom
	ISP *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	NetLinks []*CreateWirelessCloudConnectorRequestNetLinks `json:"NetLinks,omitempty" xml:"NetLinks,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// test
	UseCase *string `json:"UseCase,omitempty" xml:"UseCase,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// zjalyjs01s.5gjs.njiot
	APN *string `json:"APN,omitempty" xml:"APN,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	VSwitchs []*string `json:"VSwitchs,omitempty" xml:"VSwitchs,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// vpc-xxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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
	// example:
	//
	// E1E237B1-EED7-55D5-AE5F-671E642DB3B5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// cc5g-xxxx
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWirelessCloudConnectorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// TF-******-1633255280-43c94bf7-2dd3-4c14-8
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// cc5gg-test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	DryRun *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// example:
	//
	// 88C2F32F-B641-5980-B7A5-6907050ABCD1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// cc5gg-****
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
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWirelessCloudConnectorGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// ar5g-xxxx
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	// example:
	//
	// TF-******-1633255280-43c94bf7-2dd3-4c14-8
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cc5g-xxxx
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
	// example:
	//
	// 1651FBB6-4FBF-49FF-A9F5-DF5D696C7EC6
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAuthorizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// boct-xxx
	BatchOperateCardsTaskId *string `json:"BatchOperateCardsTaskId,omitempty" xml:"BatchOperateCardsTaskId,omitempty"`
	// example:
	//
	// TF-******-1633255280-43c94bf7-2dd3-4c14-8
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// example:
	//
	// 1651FBB6-4FBF-49FF-A9F5-DF5D696C7EC6
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBatchOperateCardsTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// ar5gg-12fnojjtkkre28t08j
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	// example:
	//
	// TF-******-1633255280-43c94bf7-2dd3-4c14-8
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cc5gg-****
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
	// example:
	//
	// A175C091-3910-51AF-8D84-D5717B9B8D38
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGroupAuthorizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// TF-******-1633255280-43c94bf7-2dd3-4c14-8
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// cciot-xxxx
	NetLinkId *string `json:"NetLinkId,omitempty" xml:"NetLinkId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cc5g-xxxx
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
	// example:
	//
	// AA93A076-3D94-51AB-AC70-4C94AF003B03
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
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIoTCloudConnectorBackhaulRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// TF-******-1633255280-43c94bf7-2dd3-4c14-8
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cc5g-xxxx
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
	// example:
	//
	// 1651FBB6-4FBF-49FF-A9F5-DF5D696C7EC6
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWirelessCloudConnectorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// TF-******-1633255280-43c94bf7-2dd3-4c14-8
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cc5gg-****
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
	// example:
	//
	// 2E759287-F208-589B-82D8-6D7A30F417E3
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
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWirelessCloudConnectorGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// TF-******-1633255280-43c94bf7-2dd3-4c14-8
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cciot-xxxx
	NetLinkId *string `json:"NetLinkId,omitempty" xml:"NetLinkId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cc5g-xxxx
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
	// example:
	//
	// 1651FBB6-4FBF-49FF-A9F5-DF5D696C7EC6
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachVpcFromNetLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// TF-CreateServerlessDBInstance-1639649079-51d933f6-0251-4057-9276
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FailCards
	Iccids []*string `json:"Iccids,omitempty" xml:"Iccids,omitempty" type:"Repeated"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// example:
	//
	// 186B6E37-D9D1-5E64-ACDE-E89F7E68A37B
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FailCardsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 11111111111
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
	// example:
	//
	// zjalyjs01s.5gjs.njiot
	APN *string `json:"APN,omitempty" xml:"APN,omitempty"`
	// example:
	//
	// 2022-08-25 17:34:57
	ActivatedTime *string `json:"ActivatedTime,omitempty" xml:"ActivatedTime,omitempty"`
	// example:
	//
	// 1000
	AlarmThreshold *int64 `json:"AlarmThreshold,omitempty" xml:"AlarmThreshold,omitempty"`
	// example:
	//
	// iotcc-xxxxxxxx
	CloudConnectorId *string `json:"CloudConnectorId,omitempty" xml:"CloudConnectorId,omitempty"`
	// example:
	//
	// task-test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// Unicom
	ISP *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	// example:
	//
	// 111111111
	Iccid *string `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	// example:
	//
	// 23412313513512
	Imei *string `json:"Imei,omitempty" xml:"Imei,omitempty"`
	// example:
	//
	// 12123123
	Imsi *string `json:"Imsi,omitempty" xml:"Imsi,omitempty"`
	// example:
	//
	// 172.24.2.121
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// example:
	//
	// 50
	LimitThreshold *int64 `json:"LimitThreshold,omitempty" xml:"LimitThreshold,omitempty"`
	// example:
	//
	// true
	Lock *string `json:"Lock,omitempty" xml:"Lock,omitempty"`
	// example:
	//
	// 2312341231235
	Msisdn *string `json:"Msisdn,omitempty" xml:"Msisdn,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 5G
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// example:
	//
	// 12312341234124
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 0B28B72A-BD16-51F7-AAEE-55247FE02729
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Deactivated
	SimStatus *string `json:"SimStatus,omitempty" xml:"SimStatus,omitempty"`
	// example:
	//
	// CONSUMER_THREE_IN_ONE
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 4096
	StopThreshold *int64 `json:"StopThreshold,omitempty" xml:"StopThreshold,omitempty"`
	// example:
	//
	// 50
	UsageDataMonth *int32 `json:"UsageDataMonth,omitempty" xml:"UsageDataMonth,omitempty"`
	// example:
	//
	// 1000
	UsageDataTotal *int64 `json:"UsageDataTotal,omitempty" xml:"UsageDataTotal,omitempty"`
	// example:
	//
	// cc5g-xxxxxxx
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
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// *********11111
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
	// example:
	//
	// ForceTwoWayStop
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// example:
	//
	// F32C0BBE-5025-5561-8976-7BD096BA21FB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCardLockReasonResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type GetCreateCustomerInfomationRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetCreateCustomerInfomationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCreateCustomerInfomationRequest) GoString() string {
	return s.String()
}

func (s *GetCreateCustomerInfomationRequest) SetRegionId(v string) *GetCreateCustomerInfomationRequest {
	s.RegionId = &v
	return s
}

type GetCreateCustomerInfomationResponseBody struct {
	// example:
	//
	// True
	CanBuyCard *string `json:"CanBuyCard,omitempty" xml:"CanBuyCard,omitempty"`
	// example:
	//
	// 2E759287-F208-589B-82D8-6D7A30F417E3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// https://*******
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s GetCreateCustomerInfomationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCreateCustomerInfomationResponseBody) GoString() string {
	return s.String()
}

func (s *GetCreateCustomerInfomationResponseBody) SetCanBuyCard(v string) *GetCreateCustomerInfomationResponseBody {
	s.CanBuyCard = &v
	return s
}

func (s *GetCreateCustomerInfomationResponseBody) SetRequestId(v string) *GetCreateCustomerInfomationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCreateCustomerInfomationResponseBody) SetURL(v string) *GetCreateCustomerInfomationResponseBody {
	s.URL = &v
	return s
}

type GetCreateCustomerInfomationResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCreateCustomerInfomationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCreateCustomerInfomationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCreateCustomerInfomationResponse) GoString() string {
	return s.String()
}

func (s *GetCreateCustomerInfomationResponse) SetHeaders(v map[string]*string) *GetCreateCustomerInfomationResponse {
	s.Headers = v
	return s
}

func (s *GetCreateCustomerInfomationResponse) SetStatusCode(v int32) *GetCreateCustomerInfomationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCreateCustomerInfomationResponse) SetBody(v *GetCreateCustomerInfomationResponseBody) *GetCreateCustomerInfomationResponse {
	s.Body = v
	return s
}

type GetCreateCustomerInformationRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// cc5g-xxx
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
	// example:
	//
	// True
	CanBuyCard *string `json:"CanBuyCard,omitempty" xml:"CanBuyCard,omitempty"`
	// example:
	//
	// AE0BA8E5-1E0C-5171-852F-826301C76BBE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// https://*******
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCreateCustomerInformationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// diagnoseTask-bp2qhwdp2n9x6zkvq****
	DiagnoseTaskId *string `json:"DiagnoseTaskId,omitempty" xml:"DiagnoseTaskId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
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
	// example:
	//
	// 1646100223024
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// example:
	//
	// 10.176.78.61
	CardIp *string `json:"CardIp,omitempty" xml:"CardIp,omitempty"`
	// example:
	//
	// 106.14.175.4
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// example:
	//
	// Domain
	DestinationType *string                                                   `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	DiagnoseItem    []*GetDiagnoseResultForSingleCardResponseBodyDiagnoseItem `json:"DiagnoseItem,omitempty" xml:"DiagnoseItem,omitempty" type:"Repeated"`
	// example:
	//
	// 1646100223264
	EndTime     *int64                                                   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ErrorResult []*GetDiagnoseResultForSingleCardResponseBodyErrorResult `json:"ErrorResult,omitempty" xml:"ErrorResult,omitempty" type:"Repeated"`
	// example:
	//
	// 89860477102170127555
	IccId *string `json:"IccId,omitempty" xml:"IccId,omitempty"`
	// example:
	//
	// 94032572-8758-575E-B306-86F59D1B826E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// failed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// cc5g-xxxx
	WirelessCloudConnectorId *string `json:"WirelessCloudConnectorId,omitempty" xml:"WirelessCloudConnectorId,omitempty"`
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
	// example:
	//
	// Terminal
	Part *string `json:"Part,omitempty" xml:"Part,omitempty"`
	// example:
	//
	// Wrong
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
	// example:
	//
	// CardStatus
	SubItem *string `json:"SubItem,omitempty" xml:"SubItem,omitempty"`
	// example:
	//
	// Card status is real-time
	SubItemInfo *string `json:"SubItemInfo,omitempty" xml:"SubItemInfo,omitempty"`
	// example:
	//
	// Normal
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
	// example:
	//
	// Terminal abnormal
	ErrorDesc *string `json:"ErrorDesc,omitempty" xml:"ErrorDesc,omitempty"`
	// example:
	//
	// CardStatus
	ErrorItem *string `json:"ErrorItem,omitempty" xml:"ErrorItem,omitempty"`
	// example:
	//
	// Warning
	ErrorLevel *string `json:"ErrorLevel,omitempty" xml:"ErrorLevel,omitempty"`
	// example:
	//
	// Terminal
	ErrorPart *string `json:"ErrorPart,omitempty" xml:"ErrorPart,omitempty"`
	// example:
	//
	// Please check whether the terminal and card are compatible or whether the terminal can send messages
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
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDiagnoseResultForSingleCardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cc5g-xxxx
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
	// example:
	//
	// Business
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// example:
	//
	// 20000
	CardCount *string `json:"CardCount,omitempty" xml:"CardCount,omitempty"`
	// example:
	//
	// 2022-3-15 22:20:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// dp-xxxx
	DataPackageId *string `json:"DataPackageId,omitempty" xml:"DataPackageId,omitempty"`
	// example:
	//
	// 200
	DataPackageType *string `json:"DataPackageType,omitempty" xml:"DataPackageType,omitempty"`
	// example:
	//
	// cc5g-test
	Description *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	Features    []*string `json:"Features,omitempty" xml:"Features,omitempty" type:"Repeated"`
	// example:
	//
	// test
	Name     *string                                          `json:"Name,omitempty" xml:"Name,omitempty"`
	NetLinks []*GetWirelessCloudConnectorResponseBodyNetLinks `json:"NetLinks,omitempty" xml:"NetLinks,omitempty" type:"Repeated"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 1651FBB6-4FBF-49FF-A9F5-DF5D696C7EC6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// CC5GCCIoT
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// example:
	//
	// Available
	Status  *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UseCase *string `json:"UseCase,omitempty" xml:"UseCase,omitempty"`
	// example:
	//
	// cc5gg-****
	WirelessCloudConnectorGroupId *string `json:"WirelessCloudConnectorGroupId,omitempty" xml:"WirelessCloudConnectorGroupId,omitempty"`
	// example:
	//
	// cc5g-xxxx
	WirelessCloudConnectorId *string `json:"WirelessCloudConnectorId,omitempty" xml:"WirelessCloudConnectorId,omitempty"`
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
	// example:
	//
	// bjaliyun01s.5gbj.bjiot
	APN         *string `json:"APN,omitempty" xml:"APN,omitempty"`
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 11111111111
	GrantAliUid *string `json:"GrantAliUid,omitempty" xml:"GrantAliUid,omitempty"`
	// example:
	//
	// Telecom
	ISP *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	// 代表创建时间的资源属性字段
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 代表资源名称的资源属性字段
	//
	// example:
	//
	// cciot-xxxx
	NetLinkId *string `json:"NetLinkId,omitempty" xml:"NetLinkId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// Attached
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// Standard
	Type     *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	VSwitchs []*string `json:"VSwitchs,omitempty" xml:"VSwitchs,omitempty" type:"Repeated"`
	// example:
	//
	// vpc-xxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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

func (s *GetWirelessCloudConnectorResponseBodyNetLinks) SetType(v string) *GetWirelessCloudConnectorResponseBodyNetLinks {
	s.Type = &v
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWirelessCloudConnectorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// f0ba469f-de7f-4dde-8390-0966a1a0ef9a
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// False
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 11111111111
	GrantAliUid *int64 `json:"GrantAliUid,omitempty" xml:"GrantAliUid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cciot-xxxx
	NetLinkId *string `json:"NetLinkId,omitempty" xml:"NetLinkId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cc5g-xxxx
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
	// example:
	//
	// E6B58AF9-F1BF-5894-B898-6CE2108463C5
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GrantNetLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type InnerLimitRateCardsRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	Iccids   []*string `json:"Iccids,omitempty" xml:"Iccids,omitempty" type:"Repeated"`
	InnerApi *bool     `json:"InnerApi,omitempty" xml:"InnerApi,omitempty"`
	RegionId *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Task     *bool     `json:"Task,omitempty" xml:"Task,omitempty"`
}

func (s InnerLimitRateCardsRequest) String() string {
	return tea.Prettify(s)
}

func (s InnerLimitRateCardsRequest) GoString() string {
	return s.String()
}

func (s *InnerLimitRateCardsRequest) SetClientToken(v string) *InnerLimitRateCardsRequest {
	s.ClientToken = &v
	return s
}

func (s *InnerLimitRateCardsRequest) SetDryRun(v bool) *InnerLimitRateCardsRequest {
	s.DryRun = &v
	return s
}

func (s *InnerLimitRateCardsRequest) SetIccids(v []*string) *InnerLimitRateCardsRequest {
	s.Iccids = v
	return s
}

func (s *InnerLimitRateCardsRequest) SetInnerApi(v bool) *InnerLimitRateCardsRequest {
	s.InnerApi = &v
	return s
}

func (s *InnerLimitRateCardsRequest) SetRegionId(v string) *InnerLimitRateCardsRequest {
	s.RegionId = &v
	return s
}

func (s *InnerLimitRateCardsRequest) SetTask(v bool) *InnerLimitRateCardsRequest {
	s.Task = &v
	return s
}

type InnerLimitRateCardsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InnerLimitRateCardsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InnerLimitRateCardsResponseBody) GoString() string {
	return s.String()
}

func (s *InnerLimitRateCardsResponseBody) SetRequestId(v string) *InnerLimitRateCardsResponseBody {
	s.RequestId = &v
	return s
}

type InnerLimitRateCardsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InnerLimitRateCardsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InnerLimitRateCardsResponse) String() string {
	return tea.Prettify(s)
}

func (s InnerLimitRateCardsResponse) GoString() string {
	return s.String()
}

func (s *InnerLimitRateCardsResponse) SetHeaders(v map[string]*string) *InnerLimitRateCardsResponse {
	s.Headers = v
	return s
}

func (s *InnerLimitRateCardsResponse) SetStatusCode(v int32) *InnerLimitRateCardsResponse {
	s.StatusCode = &v
	return s
}

func (s *InnerLimitRateCardsResponse) SetBody(v *InnerLimitRateCardsResponseBody) *InnerLimitRateCardsResponse {
	s.Body = v
	return s
}

type InnerStopCardsRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	Iccids   []*string `json:"Iccids,omitempty" xml:"Iccids,omitempty" type:"Repeated"`
	InnerApi *bool     `json:"InnerApi,omitempty" xml:"InnerApi,omitempty"`
	RegionId *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Task     *bool     `json:"Task,omitempty" xml:"Task,omitempty"`
}

func (s InnerStopCardsRequest) String() string {
	return tea.Prettify(s)
}

func (s InnerStopCardsRequest) GoString() string {
	return s.String()
}

func (s *InnerStopCardsRequest) SetClientToken(v string) *InnerStopCardsRequest {
	s.ClientToken = &v
	return s
}

func (s *InnerStopCardsRequest) SetDryRun(v bool) *InnerStopCardsRequest {
	s.DryRun = &v
	return s
}

func (s *InnerStopCardsRequest) SetIccids(v []*string) *InnerStopCardsRequest {
	s.Iccids = v
	return s
}

func (s *InnerStopCardsRequest) SetInnerApi(v bool) *InnerStopCardsRequest {
	s.InnerApi = &v
	return s
}

func (s *InnerStopCardsRequest) SetRegionId(v string) *InnerStopCardsRequest {
	s.RegionId = &v
	return s
}

func (s *InnerStopCardsRequest) SetTask(v bool) *InnerStopCardsRequest {
	s.Task = &v
	return s
}

type InnerStopCardsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InnerStopCardsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InnerStopCardsResponseBody) GoString() string {
	return s.String()
}

func (s *InnerStopCardsResponseBody) SetRequestId(v string) *InnerStopCardsResponseBody {
	s.RequestId = &v
	return s
}

type InnerStopCardsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InnerStopCardsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InnerStopCardsResponse) String() string {
	return tea.Prettify(s)
}

func (s InnerStopCardsResponse) GoString() string {
	return s.String()
}

func (s *InnerStopCardsResponse) SetHeaders(v map[string]*string) *InnerStopCardsResponse {
	s.Headers = v
	return s
}

func (s *InnerStopCardsResponse) SetStatusCode(v int32) *InnerStopCardsResponse {
	s.StatusCode = &v
	return s
}

func (s *InnerStopCardsResponse) SetBody(v *InnerStopCardsResponseBody) *InnerStopCardsResponse {
	s.Body = v
	return s
}

type ListAPNsRequest struct {
	// example:
	//
	// CMWAP
	APN *string `json:"APN,omitempty" xml:"APN,omitempty"`
	// example:
	//
	// unicom
	ISP *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	APNs []*ListAPNsResponseBodyAPNs `json:"APNs,omitempty" xml:"APNs,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// vhNxI+jwQfkLCUChCQ+iOoiRkyUbtAilGSvWuwktFBxF+bar2HpV0efI6n1+GkMv
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// A175C091-3910-51AF-8D84-D5717B9B8D38
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
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
	// 代表资源名称的资源属性字段
	//
	// example:
	//
	// zjhzaliyun01s.5gjs.njiot
	APN         *string `json:"APN,omitempty" xml:"APN,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 代表资源一级ID的资源属性字段
	//
	// example:
	//
	// unicom
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
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAPNsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	AuthorizationRuleIds []*string `json:"AuthorizationRuleIds,omitempty" xml:"AuthorizationRuleIds,omitempty" type:"Repeated"`
	// example:
	//
	// 10.0.0.1
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// example:
	//
	// -1/-1
	DestinationPort *string `json:"DestinationPort,omitempty" xml:"DestinationPort,omitempty"`
	// example:
	//
	// Cidr
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// example:
	//
	// true
	Dns *bool `json:"Dns,omitempty" xml:"Dns,omitempty"`
	// example:
	//
	// 10
	MaxResults *int64    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Names      []*string `json:"Names,omitempty" xml:"Names,omitempty" type:"Repeated"`
	// example:
	//
	// iou001238joojjaere
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// Permit
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// all
	Protocol *string   `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	Statuses []*string `json:"Statuses,omitempty" xml:"Statuses,omitempty" type:"Repeated"`
	// example:
	//
	// System
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cc5g-xxxx
	WirelessCloudConnectorId *string `json:"WirelessCloudConnectorId,omitempty" xml:"WirelessCloudConnectorId,omitempty"`
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
	// example:
	//
	// 10
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// iou001238joojjaere
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 207F9929-806E-5622-8E4F-8484F7000C22
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	//
	// example:
	//
	// ar5g-xxxx
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	// example:
	//
	// 2022-3-15 22:20:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// ar5g-test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 10.0.0.0/24
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// example:
	//
	// 0-65535
	DestinationPort *string `json:"DestinationPort,omitempty" xml:"DestinationPort,omitempty"`
	// example:
	//
	// Cidr
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// example:
	//
	// true
	Dns *bool `json:"Dns,omitempty" xml:"Dns,omitempty"`
	// 代表创建时间的资源属性字段
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Permit
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// all
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// 0.0.0.0/0
	SourceCidr *string `json:"SourceCidr,omitempty" xml:"SourceCidr,omitempty"`
	// 代表资源名称的资源属性字段
	//
	// example:
	//
	// Pending
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// Default
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAuthorizationRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// 10
	MaxResults *int64    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Names      []*string `json:"Names,omitempty" xml:"Names,omitempty" type:"Repeated"`
	// example:
	//
	// iou001238joojjaere
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Statuses []*string `json:"Statuses,omitempty" xml:"Statuses,omitempty" type:"Repeated"`
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
	// example:
	//
	// 10
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// iou001238joojjaere
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1651FBB6-4FBF-49FF-A9F5-DF5D696C7EC6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	//
	// example:
	//
	// boct-xxxx
	BatchOperateCardsTaskId *string `json:"BatchOperateCardsTaskId,omitempty" xml:"BatchOperateCardsTaskId,omitempty"`
	// example:
	//
	// 2022-3-15 22:20:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// task-test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// CurrentMonth
	EffectType *string `json:"EffectType,omitempty" xml:"EffectType,omitempty"`
	// example:
	//
	// https://examplebucket.oss-eu-central-1.aliyuncs.com/***
	IccidsOssFilePath *string `json:"IccidsOssFilePath,omitempty" xml:"IccidsOssFilePath,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// https://examplebucket.oss-eu-central-1.aliyuncs.com/***
	OperateResultOssFilePath *string `json:"OperateResultOssFilePath,omitempty" xml:"OperateResultOssFilePath,omitempty"`
	// 代表创建时间的资源属性字段
	//
	// example:
	//
	// BreakNetwork
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 100
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
	// example:
	//
	// Deleted
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// cc5g-xxxx
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBatchOperateCardsTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ListCardAreaLimitSupportAreaRequest struct {
	// example:
	//
	// en
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	ProvinceName   *string `json:"ProvinceName,omitempty" xml:"ProvinceName,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListCardAreaLimitSupportAreaRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCardAreaLimitSupportAreaRequest) GoString() string {
	return s.String()
}

func (s *ListCardAreaLimitSupportAreaRequest) SetAcceptLanguage(v string) *ListCardAreaLimitSupportAreaRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListCardAreaLimitSupportAreaRequest) SetProvinceName(v string) *ListCardAreaLimitSupportAreaRequest {
	s.ProvinceName = &v
	return s
}

func (s *ListCardAreaLimitSupportAreaRequest) SetRegionId(v string) *ListCardAreaLimitSupportAreaRequest {
	s.RegionId = &v
	return s
}

type ListCardAreaLimitSupportAreaResponseBody struct {
	Provinces []*string `json:"Provinces,omitempty" xml:"Provinces,omitempty" type:"Repeated"`
	// example:
	//
	// 1651FBB6-4FBF-49FF-A9F5-DF5D696C7EC6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCardAreaLimitSupportAreaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCardAreaLimitSupportAreaResponseBody) GoString() string {
	return s.String()
}

func (s *ListCardAreaLimitSupportAreaResponseBody) SetProvinces(v []*string) *ListCardAreaLimitSupportAreaResponseBody {
	s.Provinces = v
	return s
}

func (s *ListCardAreaLimitSupportAreaResponseBody) SetRequestId(v string) *ListCardAreaLimitSupportAreaResponseBody {
	s.RequestId = &v
	return s
}

type ListCardAreaLimitSupportAreaResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCardAreaLimitSupportAreaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCardAreaLimitSupportAreaResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCardAreaLimitSupportAreaResponse) GoString() string {
	return s.String()
}

func (s *ListCardAreaLimitSupportAreaResponse) SetHeaders(v map[string]*string) *ListCardAreaLimitSupportAreaResponse {
	s.Headers = v
	return s
}

func (s *ListCardAreaLimitSupportAreaResponse) SetStatusCode(v int32) *ListCardAreaLimitSupportAreaResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCardAreaLimitSupportAreaResponse) SetBody(v *ListCardAreaLimitSupportAreaResponseBody) *ListCardAreaLimitSupportAreaResponse {
	s.Body = v
	return s
}

type ListCardDayUsagesRequest struct {
	// This parameter is required.
	Iccids []*string `json:"Iccids,omitempty" xml:"Iccids,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	LatestMonthNum *int32 `json:"LatestMonthNum,omitempty" xml:"LatestMonthNum,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cc5g-xxxx
	WirelessCloudConnectorId *string `json:"WirelessCloudConnectorId,omitempty" xml:"WirelessCloudConnectorId,omitempty"`
}

func (s ListCardDayUsagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCardDayUsagesRequest) GoString() string {
	return s.String()
}

func (s *ListCardDayUsagesRequest) SetIccids(v []*string) *ListCardDayUsagesRequest {
	s.Iccids = v
	return s
}

func (s *ListCardDayUsagesRequest) SetLatestMonthNum(v int32) *ListCardDayUsagesRequest {
	s.LatestMonthNum = &v
	return s
}

func (s *ListCardDayUsagesRequest) SetWirelessCloudConnectorId(v string) *ListCardDayUsagesRequest {
	s.WirelessCloudConnectorId = &v
	return s
}

type ListCardDayUsagesResponseBody struct {
	Cards []*ListCardDayUsagesResponseBodyCards `json:"Cards,omitempty" xml:"Cards,omitempty" type:"Repeated"`
	// example:
	//
	// AA93A076-3D94-51AB-AC70-4C94AF003B03
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCardDayUsagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCardDayUsagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCardDayUsagesResponseBody) SetCards(v []*ListCardDayUsagesResponseBodyCards) *ListCardDayUsagesResponseBody {
	s.Cards = v
	return s
}

func (s *ListCardDayUsagesResponseBody) SetRequestId(v string) *ListCardDayUsagesResponseBody {
	s.RequestId = &v
	return s
}

type ListCardDayUsagesResponseBodyCards struct {
	// 代表资源一级ID的资源属性字段
	//
	// example:
	//
	// 89091231231219123
	Iccid           *string                                              `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	UsageDataMonths []*ListCardDayUsagesResponseBodyCardsUsageDataMonths `json:"UsageDataMonths,omitempty" xml:"UsageDataMonths,omitempty" type:"Repeated"`
}

func (s ListCardDayUsagesResponseBodyCards) String() string {
	return tea.Prettify(s)
}

func (s ListCardDayUsagesResponseBodyCards) GoString() string {
	return s.String()
}

func (s *ListCardDayUsagesResponseBodyCards) SetIccid(v string) *ListCardDayUsagesResponseBodyCards {
	s.Iccid = &v
	return s
}

func (s *ListCardDayUsagesResponseBodyCards) SetUsageDataMonths(v []*ListCardDayUsagesResponseBodyCardsUsageDataMonths) *ListCardDayUsagesResponseBodyCards {
	s.UsageDataMonths = v
	return s
}

type ListCardDayUsagesResponseBodyCardsUsageDataMonths struct {
	CardDayUsages []*ListCardDayUsagesResponseBodyCardsUsageDataMonthsCardDayUsages `json:"CardDayUsages,omitempty" xml:"CardDayUsages,omitempty" type:"Repeated"`
	// example:
	//
	// 202207
	Month *string `json:"Month,omitempty" xml:"Month,omitempty"`
	// example:
	//
	// 1030
	UsageDataMonth *string `json:"UsageDataMonth,omitempty" xml:"UsageDataMonth,omitempty"`
}

func (s ListCardDayUsagesResponseBodyCardsUsageDataMonths) String() string {
	return tea.Prettify(s)
}

func (s ListCardDayUsagesResponseBodyCardsUsageDataMonths) GoString() string {
	return s.String()
}

func (s *ListCardDayUsagesResponseBodyCardsUsageDataMonths) SetCardDayUsages(v []*ListCardDayUsagesResponseBodyCardsUsageDataMonthsCardDayUsages) *ListCardDayUsagesResponseBodyCardsUsageDataMonths {
	s.CardDayUsages = v
	return s
}

func (s *ListCardDayUsagesResponseBodyCardsUsageDataMonths) SetMonth(v string) *ListCardDayUsagesResponseBodyCardsUsageDataMonths {
	s.Month = &v
	return s
}

func (s *ListCardDayUsagesResponseBodyCardsUsageDataMonths) SetUsageDataMonth(v string) *ListCardDayUsagesResponseBodyCardsUsageDataMonths {
	s.UsageDataMonth = &v
	return s
}

type ListCardDayUsagesResponseBodyCardsUsageDataMonthsCardDayUsages struct {
	// example:
	//
	// 20220701
	Day *string `json:"Day,omitempty" xml:"Day,omitempty"`
	// example:
	//
	// 40
	UsageData *string `json:"UsageData,omitempty" xml:"UsageData,omitempty"`
}

func (s ListCardDayUsagesResponseBodyCardsUsageDataMonthsCardDayUsages) String() string {
	return tea.Prettify(s)
}

func (s ListCardDayUsagesResponseBodyCardsUsageDataMonthsCardDayUsages) GoString() string {
	return s.String()
}

func (s *ListCardDayUsagesResponseBodyCardsUsageDataMonthsCardDayUsages) SetDay(v string) *ListCardDayUsagesResponseBodyCardsUsageDataMonthsCardDayUsages {
	s.Day = &v
	return s
}

func (s *ListCardDayUsagesResponseBodyCardsUsageDataMonthsCardDayUsages) SetUsageData(v string) *ListCardDayUsagesResponseBodyCardsUsageDataMonthsCardDayUsages {
	s.UsageData = &v
	return s
}

type ListCardDayUsagesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCardDayUsagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCardDayUsagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCardDayUsagesResponse) GoString() string {
	return s.String()
}

func (s *ListCardDayUsagesResponse) SetHeaders(v map[string]*string) *ListCardDayUsagesResponse {
	s.Headers = v
	return s
}

func (s *ListCardDayUsagesResponse) SetStatusCode(v int32) *ListCardDayUsagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCardDayUsagesResponse) SetBody(v *ListCardDayUsagesResponseBody) *ListCardDayUsagesResponse {
	s.Body = v
	return s
}

type ListCardUsagesRequest struct {
	Iccids []*string `json:"Iccids,omitempty" xml:"Iccids,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// cc5g-xxxx
	WirelessCloudConnectorId *string `json:"WirelessCloudConnectorId,omitempty" xml:"WirelessCloudConnectorId,omitempty"`
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
	Cards []*ListCardUsagesResponseBodyCards `json:"Cards,omitempty" xml:"Cards,omitempty" type:"Repeated"`
	// example:
	//
	// 186B6E37-D9D1-5E64-ACDE-E89F7E68A37B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	//
	// example:
	//
	// 89091231231219123
	Iccid *string `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	// example:
	//
	// 20
	UsageDataMonth *int64 `json:"UsageDataMonth,omitempty" xml:"UsageDataMonth,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCardUsagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// bjaliyun01s.5gbj.bjiot
	Apn *string `json:"Apn,omitempty" xml:"Apn,omitempty"`
	// example:
	//
	// 8909123123121912
	Iccid  *string   `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	Iccids []*string `json:"Iccids,omitempty" xml:"Iccids,omitempty" type:"Repeated"`
	// example:
	//
	// 10.0.0.1
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// example:
	//
	// true
	Lock *bool `json:"Lock,omitempty" xml:"Lock,omitempty"`
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 1441204374591
	Msisdn *string `json:"Msisdn,omitempty" xml:"Msisdn,omitempty"`
	// example:
	//
	// cciot-xxxx
	NetLinkId *string `json:"NetLinkId,omitempty" xml:"NetLinkId,omitempty"`
	// example:
	//
	// iou001238joojjaere
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// true
	Online   *bool     `json:"Online,omitempty" xml:"Online,omitempty"`
	Statuses []*string `json:"Statuses,omitempty" xml:"Statuses,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// cc5g-xxxx
	WirelessCloudConnectorId *string `json:"WirelessCloudConnectorId,omitempty" xml:"WirelessCloudConnectorId,omitempty"`
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
	Cards []*ListCardsResponseBodyCards `json:"Cards,omitempty" xml:"Cards,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// iou001238joojjaere
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 88C2F32F-B641-5980-B7A5-6907050ABCD1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	//
	// example:
	//
	// bjaliyun01s.5gbj.bjiot
	APN *string `json:"APN,omitempty" xml:"APN,omitempty"`
	// example:
	//
	// 2022-3-15 22:20:00
	ActivatedTime *string `json:"ActivatedTime,omitempty" xml:"ActivatedTime,omitempty"`
	// example:
	//
	// Available
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// Telecom
	ISP *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	// 代表资源一级ID的资源属性字段
	//
	// example:
	//
	// 89091231231219123
	Iccid *string `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	// example:
	//
	// 7896129300
	Imei *string `json:"Imei,omitempty" xml:"Imei,omitempty"`
	// example:
	//
	// 8912039123
	Imsi *string `json:"Imsi,omitempty" xml:"Imsi,omitempty"`
	// example:
	//
	// 10.0.0.118
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// example:
	//
	// false
	Lock *bool `json:"Lock,omitempty" xml:"Lock,omitempty"`
	// example:
	//
	// 1441204374591
	Msisdn *string `json:"Msisdn,omitempty" xml:"Msisdn,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 代表资源名称的资源属性字段
	//
	// example:
	//
	// 5G
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// example:
	//
	// 89099123012
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// CONSUMER_THREE_IN_ONE
	Spec   *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 20
	UsageDataMonth *int64 `json:"UsageDataMonth,omitempty" xml:"UsageDataMonth,omitempty"`
	// example:
	//
	// 50
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCardsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	DataPackageIds []*string `json:"DataPackageIds,omitempty" xml:"DataPackageIds,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *int64    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Names      []*string `json:"Names,omitempty" xml:"Names,omitempty" type:"Repeated"`
	// example:
	//
	// iou001238joojjaere
	NextToken *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Statuses  []*string `json:"Statuses,omitempty" xml:"Statuses,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// cc5g-xxxx
	WirelessCloudConnectorId *string `json:"WirelessCloudConnectorId,omitempty" xml:"WirelessCloudConnectorId,omitempty"`
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
	// example:
	//
	// 10
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// iou001238joojjaere
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 8734B5B2-FC84-54C4-8038-D40228CF01FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// example:
	//
	// 20000
	CardCount *string `json:"CardCount,omitempty" xml:"CardCount,omitempty"`
	// example:
	//
	// 2022-3-15 22:20:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// dp-xxxx
	DataPackageId *string `json:"DataPackageId,omitempty" xml:"DataPackageId,omitempty"`
	// example:
	//
	// 2099-3-15 22:20:00
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// example:
	//
	// Telecom
	ISP *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	// 代表创建时间的资源属性字段
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 200
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// 代表资源名称的资源属性字段
	//
	// example:
	//
	// Available
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataPackagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// example:
	//
	// 89860477102170127555
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// Iccid
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cc5g-xxxx
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
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// E1E237B1-EED7-55D5-AE5F-671E642DB3B5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 50
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// example:
	//
	// 1662307200000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// example:
	//
	// 10.176.78.61
	CardIp *string `json:"CardIp,omitempty" xml:"CardIp,omitempty"`
	// example:
	//
	// www.xxx.com
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// example:
	//
	// Domain
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// example:
	//
	// diagnoseTask-bp2x7tzkdl655liw3mhbg
	DiagnoseTaskId *string `json:"DiagnoseTaskId,omitempty" xml:"DiagnoseTaskId,omitempty"`
	// example:
	//
	// 1662307202400
	DiagnoseTime *int64 `json:"DiagnoseTime,omitempty" xml:"DiagnoseTime,omitempty"`
	// example:
	//
	// 1662307201200
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 89860477102170127555
	IccId *string `json:"IccId,omitempty" xml:"IccId,omitempty"`
	// example:
	//
	// 89860477102170127555
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// Iccid
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// failed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// cc5g-xxxx
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDiagnoseInfoForSingleCardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	AuthorizationRuleIds []*string `json:"AuthorizationRuleIds,omitempty" xml:"AuthorizationRuleIds,omitempty" type:"Repeated"`
	// example:
	//
	// 10.0.0.1
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// example:
	//
	// -1/-1
	DestinationPort *string `json:"DestinationPort,omitempty" xml:"DestinationPort,omitempty"`
	// example:
	//
	// Cidr
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// example:
	//
	// true
	Dns *bool `json:"Dns,omitempty" xml:"Dns,omitempty"`
	// example:
	//
	// 10
	MaxResults *int64    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Names      []*string `json:"Names,omitempty" xml:"Names,omitempty" type:"Repeated"`
	// example:
	//
	// iou001238joojjaere
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// Permit
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// all
	Protocol *string   `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	Statuses []*string `json:"Statuses,omitempty" xml:"Statuses,omitempty" type:"Repeated"`
	// example:
	//
	// System
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cc5gg-****
	WirelessCloudConnectorGroupId *string `json:"WirelessCloudConnectorGroupId,omitempty" xml:"WirelessCloudConnectorGroupId,omitempty"`
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
	// example:
	//
	// 10
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// iou001238joojjaere
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1651FBB6-4FBF-49FF-A9F5-DF5D696C7EC6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	//
	// example:
	//
	// ar5gg-y5g4rp2w1tl1w7zg59
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	// example:
	//
	// 2022-06-21 15:27:24
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// ar5gg-test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 47.103.184.100/32
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// example:
	//
	// 0-65535
	DestinationPort *string `json:"DestinationPort,omitempty" xml:"DestinationPort,omitempty"`
	// example:
	//
	// Cidr
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// example:
	//
	// true
	Dns *bool `json:"Dns,omitempty" xml:"Dns,omitempty"`
	// 代表创建时间的资源属性字段
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Permit
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// all
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// 0.0.0.0/0
	SourceCidr *string `json:"SourceCidr,omitempty" xml:"SourceCidr,omitempty"`
	// 代表资源名称的资源属性字段
	//
	// example:
	//
	// Creating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// Default
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGroupAuthorizationRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// cciot-xxxx
	NetLinkId *string `json:"NetLinkId,omitempty" xml:"NetLinkId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cc5g-xxxx
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
	// example:
	//
	// cciot-xxxx
	NetLinkId *string `json:"NetLinkId,omitempty" xml:"NetLinkId,omitempty"`
	// example:
	//
	// 94032572-8758-575E-B306-86F59D1B826E
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
	//
	// example:
	//
	// full cidr block route
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 代表资源名称的资源属性字段
	//
	// example:
	//
	// 10.33.190.0/24
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// 代表创建时间的资源属性字段
	//
	// example:
	//
	// eni-uf677iw3xihqxiz2ssir
	NextHopId *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	// 代表创建时间的资源属性字段
	//
	// example:
	//
	// NetworkInterface
	NextHopType *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	// 代表创建时间的资源属性字段
	//
	// example:
	//
	// published
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
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIoTCloudConnectorBackhaulRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// iou001238joojjaere
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// NewBuy
	OrderAction *string   `json:"OrderAction,omitempty" xml:"OrderAction,omitempty"`
	OrderIds    []*string `json:"OrderIds,omitempty" xml:"OrderIds,omitempty" type:"Repeated"`
	Statuses    []*string `json:"Statuses,omitempty" xml:"Statuses,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// cc5g-xxxx
	WirelessCloudConnectorId *string `json:"WirelessCloudConnectorId,omitempty" xml:"WirelessCloudConnectorId,omitempty"`
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
	// example:
	//
	// 10
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// iou001238joojjaere
	NextToken *string                         `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Orders    []*ListOrdersResponseBodyOrders `json:"Orders,omitempty" xml:"Orders,omitempty" type:"Repeated"`
	// example:
	//
	// 88C2F32F-B641-5980-B7A5-6907050ABCD1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	//
	// example:
	//
	// NewBuy
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// 20000
	CardCount *string `json:"CardCount,omitempty" xml:"CardCount,omitempty"`
	// example:
	//
	// 5G
	CardNetType *string `json:"CardNetType,omitempty" xml:"CardNetType,omitempty"`
	// example:
	//
	// CONSUMER_THREE_IN_ONE
	CardType    *string `json:"CardType,omitempty" xml:"CardType,omitempty"`
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// example:
	//
	// 139000*******
	ContactPhone *string `json:"ContactPhone,omitempty" xml:"ContactPhone,omitempty"`
	// example:
	//
	// 2022-3-15 21:20:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// order-test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// SF14*********
	LogisticsId *string `json:"LogisticsId,omitempty" xml:"LogisticsId,omitempty"`
	// example:
	//
	// Signed
	LogisticsStatus *string `json:"LogisticsStatus,omitempty" xml:"LogisticsStatus,omitempty"`
	// example:
	//
	// SF
	LogisticsType *string `json:"LogisticsType,omitempty" xml:"LogisticsType,omitempty"`
	// example:
	//
	// 2022-3-15 22:20:00
	LogisticsUpdateTime *string `json:"LogisticsUpdateTime,omitempty" xml:"LogisticsUpdateTime,omitempty"`
	// example:
	//
	// 89************
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 2022-3-15 21:30:00
	PayTime     *string `json:"PayTime,omitempty" xml:"PayTime,omitempty"`
	PostAddress *string `json:"PostAddress,omitempty" xml:"PostAddress,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 代表资源名称的资源属性字段
	//
	// example:
	//
	// Pending
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOrdersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	Regions []*ListRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// example:
	//
	// 2DCFA69E-A161-512D-99A7-108022580719
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	//
	// example:
	//
	// cn-hangzhou
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// iou001238joojjaere
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// cn-hangzhou
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
	// example:
	//
	// 10
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// iou001238joojjaere
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// D3BAFE75-C670-52DB-9739-03133E5FC914
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
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
	// example:
	//
	// 2022-07-04 19:19:14
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// cc5gg-test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 代表创建时间的资源属性字段
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 代表资源名称的资源属性字段
	//
	// example:
	//
	// Creating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 代表资源一级ID的资源属性字段
	//
	// example:
	//
	// cc5gg-xxxx
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
	// example:
	//
	// Business
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// example:
	//
	// 115
	CardCount *string `json:"CardCount,omitempty" xml:"CardCount,omitempty"`
	// example:
	//
	// 2022-04-28 17:40:53
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// dp-xxxx
	DataPackageId *string `json:"DataPackageId,omitempty" xml:"DataPackageId,omitempty"`
	// example:
	//
	// 200
	DataPackageType *string `json:"DataPackageType,omitempty" xml:"DataPackageType,omitempty"`
	// example:
	//
	// cc5g-test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// CC5G
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// example:
	//
	// Available
	Status  *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UseCase *string `json:"UseCase,omitempty" xml:"UseCase,omitempty"`
	// example:
	//
	// cc5g-62eb02o7lz1w5aurl3
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
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWirelessCloudConnectorGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// Business
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// example:
	//
	// false
	IsInGroup *string `json:"IsInGroup,omitempty" xml:"IsInGroup,omitempty"`
	// example:
	//
	// 10
	MaxResults *int64    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Names      []*string `json:"Names,omitempty" xml:"Names,omitempty" type:"Repeated"`
	// example:
	//
	// iou001238joojjaere
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Statuses []*string `json:"Statuses,omitempty" xml:"Statuses,omitempty" type:"Repeated"`
	// example:
	//
	// cc5gg-****
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
	// example:
	//
	// 10
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// iou001238joojjaere
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1651FBB6-4FBF-49FF-A9F5-DF5D696C7EC6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
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
	// example:
	//
	// Business
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// example:
	//
	// 20000
	CardCount *string `json:"CardCount,omitempty" xml:"CardCount,omitempty"`
	// example:
	//
	// 2022-3-15 22:20:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// dp-xxxx
	DataPackageId *string `json:"DataPackageId,omitempty" xml:"DataPackageId,omitempty"`
	// example:
	//
	// 200
	DataPackageType *string `json:"DataPackageType,omitempty" xml:"DataPackageType,omitempty"`
	// example:
	//
	// cc5g-test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 代表创建时间的资源属性字段
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// CC5G
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// 代表资源名称的资源属性字段
	//
	// example:
	//
	// Available
	Status  *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UseCase *string `json:"UseCase,omitempty" xml:"UseCase,omitempty"`
	// example:
	//
	// cc5gg-****
	WirelessCloudConnectorGroupId *string `json:"WirelessCloudConnectorGroupId,omitempty" xml:"WirelessCloudConnectorGroupId,omitempty"`
	// 代表资源一级ID的资源属性字段
	//
	// example:
	//
	// cc5g-xxxx
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWirelessCloudConnectorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
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
	// example:
	//
	// C2C1DE7C-E14C-53DE-920D-D63ACD7421A4
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
	//
	// example:
	//
	// cn-hangzhou-h
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListZonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// TF-******-1633255280-43c94bf7-2dd3-4c14-8
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	Iccids []*string `json:"Iccids,omitempty" xml:"Iccids,omitempty" type:"Repeated"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// example:
	//
	// 8734B5B2-FC84-54C4-8038-D40228CF01FE
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LockCardsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// TF-******-1633255280-43c94bf7-2dd3-4c14-8
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// DirectMode
	FeatureName *string `json:"FeatureName,omitempty" xml:"FeatureName,omitempty"`
	// example:
	//
	// On
	FeatureValue *string `json:"FeatureValue,omitempty" xml:"FeatureValue,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cc5g-xxxx
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
	// example:
	//
	// 999E0F64-9A10-567B-BB84-B7C46023BA3B
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
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyWirelessCloudConnectorFeatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
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
	// example:
	//
	// 2DCFA69E-A161-512D-99A7-108022580719
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenCc5gServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// TF-******-1633255280-43c94bf7-2dd3-4c14-8
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	Iccids []*string `json:"Iccids,omitempty" xml:"Iccids,omitempty" type:"Repeated"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// example:
	//
	// D2A9864E-BB8D-5A5C-90FC-E2FBD5D0BF9B
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RebindCardsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// TF-******-1633255280-43c94bf7-2dd3-4c14-8
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cc5gg-****
	WirelessCloudConnectorGroupId *string `json:"WirelessCloudConnectorGroupId,omitempty" xml:"WirelessCloudConnectorGroupId,omitempty"`
	// This parameter is required.
	WirelessCloudConnectorIds []*string `json:"WirelessCloudConnectorIds,omitempty" xml:"WirelessCloudConnectorIds,omitempty" type:"Repeated"`
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
	// example:
	//
	// 2DCFA69E-A161-512D-99A7-108022580719
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
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveWirelessCloudConnectorFromGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ResetAreaLimitCardsRequest struct {
	// example:
	//
	// TF-******-1633255280-43c94bf7-2dd3-4c14-8
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	Iccids []*string `json:"Iccids,omitempty" xml:"Iccids,omitempty" type:"Repeated"`
	// This parameter is required.
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ResetAreaLimitCardsRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetAreaLimitCardsRequest) GoString() string {
	return s.String()
}

func (s *ResetAreaLimitCardsRequest) SetClientToken(v string) *ResetAreaLimitCardsRequest {
	s.ClientToken = &v
	return s
}

func (s *ResetAreaLimitCardsRequest) SetDryRun(v bool) *ResetAreaLimitCardsRequest {
	s.DryRun = &v
	return s
}

func (s *ResetAreaLimitCardsRequest) SetIccids(v []*string) *ResetAreaLimitCardsRequest {
	s.Iccids = v
	return s
}

func (s *ResetAreaLimitCardsRequest) SetProvince(v string) *ResetAreaLimitCardsRequest {
	s.Province = &v
	return s
}

func (s *ResetAreaLimitCardsRequest) SetRegionId(v string) *ResetAreaLimitCardsRequest {
	s.RegionId = &v
	return s
}

type ResetAreaLimitCardsResponseBody struct {
	// example:
	//
	// 1651FBB6-4FBF-49FF-A9F5-DF5D696C7EC6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetAreaLimitCardsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetAreaLimitCardsResponseBody) GoString() string {
	return s.String()
}

func (s *ResetAreaLimitCardsResponseBody) SetRequestId(v string) *ResetAreaLimitCardsResponseBody {
	s.RequestId = &v
	return s
}

type ResetAreaLimitCardsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetAreaLimitCardsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetAreaLimitCardsResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetAreaLimitCardsResponse) GoString() string {
	return s.String()
}

func (s *ResetAreaLimitCardsResponse) SetHeaders(v map[string]*string) *ResetAreaLimitCardsResponse {
	s.Headers = v
	return s
}

func (s *ResetAreaLimitCardsResponse) SetStatusCode(v int32) *ResetAreaLimitCardsResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetAreaLimitCardsResponse) SetBody(v *ResetAreaLimitCardsResponseBody) *ResetAreaLimitCardsResponse {
	s.Body = v
	return s
}

type ResumeCardsRequest struct {
	// example:
	//
	// TF-******-1633255280-43c94bf7-2dd3-4c14-8
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	Iccids []*string `json:"Iccids,omitempty" xml:"Iccids,omitempty" type:"Repeated"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// example:
	//
	// 1651FBB6-4FBF-49FF-A9F5-DF5D696C7EC6
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResumeCardsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// 8278082c-0b8c-412c-b0ad-876a3cf0d0f8
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// iotcc-ua6xy4vpmi24y114ws
	NetLinkId *string `json:"NetLinkId,omitempty" xml:"NetLinkId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cc5g-xxxx
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
	// example:
	//
	// 1E781588-69D4-56C3-93BA-DEFECF2596B8
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeNetLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// TF-******-1633255280-43c94bf7-2dd3-4c14-8
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	Iccids []*string `json:"Iccids,omitempty" xml:"Iccids,omitempty" type:"Repeated"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// example:
	//
	// 8734B5B2-FC84-54C4-8038-D40228CF01FE
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopCardsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 1662307201200
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// example:
	//
	// 106.14.175.4
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1662307201200
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// example:
	//
	// 253460731706911258
	ResourceUid *int64 `json:"ResourceUid,omitempty" xml:"ResourceUid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 89860477102170127555
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cc5g-xxxx
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
	// example:
	//
	// diagnoseTask-bp2n6rgaj49qcs34jyzo8
	DiagnoseTaskId *string `json:"DiagnoseTaskId,omitempty" xml:"DiagnoseTaskId,omitempty"`
	// example:
	//
	// 1651FBB6-4FBF-49FF-A9F5-DF5D696C7EC6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitDiagnoseTaskForSingleCardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// TF-******-1633255280-43c94bf7-2dd3-4c14-8
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cc5g-xxxx
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
	//
	// example:
	//
	// 74ADDCF3-E3FE-5200-9290-1320125D71DC
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
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchWirelessCloudConnectorToBusinessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// TF-******-1633255280-43c94bf7-2dd3-4c14-8
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	Iccids []*string `json:"Iccids,omitempty" xml:"Iccids,omitempty" type:"Repeated"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// example:
	//
	// 81CB64DF-E328-50A3-B386-03CAD604C114
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnlockCardsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// ar5g-xxxx
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	// example:
	//
	// TF-******-1633255280-43c94bf7-2dd3-4c14-8
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// ar5g-test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 10.0.0.1
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// example:
	//
	// -1/-1
	DestinationPort *string `json:"DestinationPort,omitempty" xml:"DestinationPort,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// Test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Permit
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// all
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// 0.0.0.0/0
	SourceCidr *string `json:"SourceCidr,omitempty" xml:"SourceCidr,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cc5g-xxxx
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
	// example:
	//
	// D3BAFE75-C670-52DB-9739-03133E5FC914
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAuthorizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// boct-xxxx
	BatchOperateCardsTaskId *string `json:"BatchOperateCardsTaskId,omitempty" xml:"BatchOperateCardsTaskId,omitempty"`
	// example:
	//
	// TF-******-1633255280-43c94bf7-2dd3-4c14-8
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// task-test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// CurrentMonth
	EffectType        *string   `json:"EffectType,omitempty" xml:"EffectType,omitempty"`
	Iccids            []*string `json:"Iccids,omitempty" xml:"Iccids,omitempty" type:"Repeated"`
	IccidsOssFilePath *string   `json:"IccidsOssFilePath,omitempty" xml:"IccidsOssFilePath,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// BreakNetwork
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 100
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
	// example:
	//
	// 1651FBB6-4FBF-49FF-A9F5-DF5D696C7EC6
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateBatchOperateCardsTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// TF-******-1633255280-43c94bf7-2dd3-4c14-8
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// card-test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// 890912389123
	Iccid *string `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cc5g-xxxx
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
	// example:
	//
	// 8734B5B2-FC84-54C4-8038-D40228CF01FE
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// ar5g-xxxx
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	// example:
	//
	// TF-******-1633255280-43c94bf7-2dd3-4c14-8
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// ar5g-test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 10.0.0.1
	DestinationIp *string `json:"DestinationIp,omitempty" xml:"DestinationIp,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 0.0.0.0
	SourceDNSIp *string `json:"SourceDNSIp,omitempty" xml:"SourceDNSIp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cc5g-xxxx
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
	// example:
	//
	// 1651FBB6-4FBF-49FF-A9F5-DF5D696C7EC6
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDNSAuthorizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// ar5gg-xxxx
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	// example:
	//
	// TF-******-1633255280-43c94bf7-2dd3-4c14-8
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 192.168.XX.XX
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// example:
	//
	// 80/80
	DestinationPort *string `json:"DestinationPort,omitempty" xml:"DestinationPort,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Permit
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// all
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// 192.168.0.1
	SourceCidr *string `json:"SourceCidr,omitempty" xml:"SourceCidr,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cc5gg-****
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
	// example:
	//
	// 79B01B1A-7089-552F-9798-DEDFFBAAED5F
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGroupAuthorizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// ar5gg-xxxx
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	// example:
	//
	// TF-******-1633255280-43c94bf7-2dd3-4c14-8
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 192.168.0.2
	DestinationIp *string `json:"DestinationIp,omitempty" xml:"DestinationIp,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 100.100.2.136
	SourceDNSIp *string `json:"SourceDNSIp,omitempty" xml:"SourceDNSIp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cc5gg-****
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
	// example:
	//
	// 207F9929-806E-5622-8E4F-8484F7000C22
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
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGroupDnsAuthorizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// TF-******-1633255280-43c94bf7-2dd3-4c14-8
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// cc5g-test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cc5g-xxxx
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
	// example:
	//
	// 1651FBB6-4FBF-49FF-A9F5-DF5D696C7EC6
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWirelessCloudConnectorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// TF-******-1633255280-43c94bf7-2dd3-4c14-8
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cc5gg-****
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
	// example:
	//
	// 8734B5B2-FC84-54C4-8038-D40228CF01FE
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
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWirelessCloudConnectorGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

// Summary:
//
// 添加5G高速上云服务实例的DNS授权规则
//
// @param request - AddDNSAuthorizationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDNSAuthorizationRuleResponse
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

// Summary:
//
// 添加5G高速上云服务实例的DNS授权规则
//
// @param request - AddDNSAuthorizationRuleRequest
//
// @return AddDNSAuthorizationRuleResponse
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

// Summary:
//
// 添加5G高速上云服务实例组的DNS授权规则
//
// @param request - AddGroupDnsAuthorizationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddGroupDnsAuthorizationRuleResponse
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

// Summary:
//
// 添加5G高速上云服务实例组的DNS授权规则
//
// @param request - AddGroupDnsAuthorizationRuleRequest
//
// @return AddGroupDnsAuthorizationRuleResponse
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

// Summary:
//
// 添加5G高速上云服务实例到组
//
// @param request - AddWirelessCloudConnectorToGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddWirelessCloudConnectorToGroupResponse
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

// Summary:
//
// 添加5G高速上云服务实例到组
//
// @param request - AddWirelessCloudConnectorToGroupRequest
//
// @return AddWirelessCloudConnectorToGroupResponse
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

// Summary:
//
// 创建5G高速上云服务实例下的网络连接
//
// @param request - AttachVpcToNetLinkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachVpcToNetLinkResponse
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

// Summary:
//
// 创建5G高速上云服务实例下的网络连接
//
// @param request - AttachVpcToNetLinkRequest
//
// @return AttachVpcToNetLinkResponse
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

// Summary:
//
// 创建5G高速上云服务实例的授权规则
//
// @param request - CreateAuthorizationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAuthorizationRuleResponse
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

// Summary:
//
// 创建5G高速上云服务实例的授权规则
//
// @param request - CreateAuthorizationRuleRequest
//
// @return CreateAuthorizationRuleResponse
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

// Summary:
//
// 创建批量操作卡任务
//
// @param request - CreateBatchOperateCardsTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBatchOperateCardsTaskResponse
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

// Summary:
//
// 创建批量操作卡任务
//
// @param request - CreateBatchOperateCardsTaskRequest
//
// @return CreateBatchOperateCardsTaskResponse
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

// Summary:
//
// 创建5G高速上云服务实例组的授权规则
//
// @param request - CreateGroupAuthorizationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGroupAuthorizationRuleResponse
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

// Summary:
//
// 创建5G高速上云服务实例组的授权规则
//
// @param request - CreateGroupAuthorizationRuleRequest
//
// @return CreateGroupAuthorizationRuleResponse
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

// Summary:
//
// 下发iotcc实例的回程路由
//
// @param request - CreateIoTCloudConnectorBackhaulRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIoTCloudConnectorBackhaulRouteResponse
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

// Summary:
//
// 下发iotcc实例的回程路由
//
// @param request - CreateIoTCloudConnectorBackhaulRouteRequest
//
// @return CreateIoTCloudConnectorBackhaulRouteResponse
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

// Summary:
//
// 创建5G高速上云服务实例
//
// @param request - CreateWirelessCloudConnectorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWirelessCloudConnectorResponse
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

// Summary:
//
// 创建5G高速上云服务实例
//
// @param request - CreateWirelessCloudConnectorRequest
//
// @return CreateWirelessCloudConnectorResponse
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

// Summary:
//
// 创建5G高速上云服务实例组
//
// @param request - CreateWirelessCloudConnectorGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWirelessCloudConnectorGroupResponse
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

// Summary:
//
// 创建5G高速上云服务实例组
//
// @param request - CreateWirelessCloudConnectorGroupRequest
//
// @return CreateWirelessCloudConnectorGroupResponse
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

// Summary:
//
// 删除5G高速上云服务实例的授权规则
//
// @param request - DeleteAuthorizationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAuthorizationRuleResponse
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

// Summary:
//
// 删除5G高速上云服务实例的授权规则
//
// @param request - DeleteAuthorizationRuleRequest
//
// @return DeleteAuthorizationRuleResponse
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

// Summary:
//
// 删除批量操作卡任务
//
// @param request - DeleteBatchOperateCardsTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBatchOperateCardsTaskResponse
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

// Summary:
//
// 删除批量操作卡任务
//
// @param request - DeleteBatchOperateCardsTaskRequest
//
// @return DeleteBatchOperateCardsTaskResponse
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

// Summary:
//
// 删除5G高速上云服务实例组的授权规则
//
// @param request - DeleteGroupAuthorizationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGroupAuthorizationRuleResponse
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

// Summary:
//
// 删除5G高速上云服务实例组的授权规则
//
// @param request - DeleteGroupAuthorizationRuleRequest
//
// @return DeleteGroupAuthorizationRuleResponse
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

// Summary:
//
// 删除iotcc实例的回程路由
//
// @param request - DeleteIoTCloudConnectorBackhaulRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIoTCloudConnectorBackhaulRouteResponse
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

// Summary:
//
// 删除iotcc实例的回程路由
//
// @param request - DeleteIoTCloudConnectorBackhaulRouteRequest
//
// @return DeleteIoTCloudConnectorBackhaulRouteResponse
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

// Summary:
//
// 删除5G高速上云服务实例
//
// @param request - DeleteWirelessCloudConnectorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWirelessCloudConnectorResponse
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

// Summary:
//
// 删除5G高速上云服务实例
//
// @param request - DeleteWirelessCloudConnectorRequest
//
// @return DeleteWirelessCloudConnectorResponse
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

// Summary:
//
// 删除5G高速上云服务实例分组
//
// @param request - DeleteWirelessCloudConnectorGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWirelessCloudConnectorGroupResponse
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

// Summary:
//
// 删除5G高速上云服务实例分组
//
// @param request - DeleteWirelessCloudConnectorGroupRequest
//
// @return DeleteWirelessCloudConnectorGroupResponse
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

// Summary:
//
// 解除5G高速上云服务实例下的网络连接和VPC的绑定
//
// @param request - DetachVpcFromNetLinkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachVpcFromNetLinkResponse
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

// Summary:
//
// 解除5G高速上云服务实例下的网络连接和VPC的绑定
//
// @param request - DetachVpcFromNetLinkRequest
//
// @return DetachVpcFromNetLinkResponse
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

// Summary:
//
// 卡注销接口
//
// @param request - FailCardsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FailCardsResponse
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

// Summary:
//
// 卡注销接口
//
// @param request - FailCardsRequest
//
// @return FailCardsResponse
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

// Summary:
//
// 查询5G高速上云服务实例的详细连接信息
//
// @param request - GetCardRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCardResponse
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

// Summary:
//
// 查询5G高速上云服务实例的详细连接信息
//
// @param request - GetCardRequest
//
// @return GetCardResponse
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

// Summary:
//
// 查询卡锁定停机的原因
//
// @param request - GetCardLockReasonRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCardLockReasonResponse
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

// Summary:
//
// 查询卡锁定停机的原因
//
// @param request - GetCardLockReasonRequest
//
// @return GetCardLockReasonResponse
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

// @param request - GetCreateCustomerInfomationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCreateCustomerInfomationResponse
func (client *Client) GetCreateCustomerInfomationWithOptions(request *GetCreateCustomerInfomationRequest, runtime *util.RuntimeOptions) (_result *GetCreateCustomerInfomationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCreateCustomerInfomation"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCreateCustomerInfomationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetCreateCustomerInfomationRequest
//
// @return GetCreateCustomerInfomationResponse
func (client *Client) GetCreateCustomerInfomation(request *GetCreateCustomerInfomationRequest) (_result *GetCreateCustomerInfomationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCreateCustomerInfomationResponse{}
	_body, _err := client.GetCreateCustomerInfomationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取录入客户资料的相关信息
//
// @param request - GetCreateCustomerInformationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCreateCustomerInformationResponse
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

// Summary:
//
// 获取录入客户资料的相关信息
//
// @param request - GetCreateCustomerInformationRequest
//
// @return GetCreateCustomerInformationResponse
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

// Summary:
//
// 查询单卡诊断结果
//
// @param request - GetDiagnoseResultForSingleCardRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDiagnoseResultForSingleCardResponse
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

// Summary:
//
// 查询单卡诊断结果
//
// @param request - GetDiagnoseResultForSingleCardRequest
//
// @return GetDiagnoseResultForSingleCardResponse
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

// Summary:
//
// 查询5G高速上云服务实例详情
//
// @param request - GetWirelessCloudConnectorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWirelessCloudConnectorResponse
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

// Summary:
//
// 查询5G高速上云服务实例详情
//
// @param request - GetWirelessCloudConnectorRequest
//
// @return GetWirelessCloudConnectorResponse
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

// Summary:
//
// 分享授权网络连接
//
// @param request - GrantNetLinkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantNetLinkResponse
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

// Summary:
//
// 分享授权网络连接
//
// @param request - GrantNetLinkRequest
//
// @return GrantNetLinkResponse
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

// @param request - InnerLimitRateCardsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InnerLimitRateCardsResponse
func (client *Client) InnerLimitRateCardsWithOptions(request *InnerLimitRateCardsRequest, runtime *util.RuntimeOptions) (_result *InnerLimitRateCardsResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.InnerApi)) {
		query["InnerApi"] = request.InnerApi
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Task)) {
		query["Task"] = request.Task
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InnerLimitRateCards"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InnerLimitRateCardsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - InnerLimitRateCardsRequest
//
// @return InnerLimitRateCardsResponse
func (client *Client) InnerLimitRateCards(request *InnerLimitRateCardsRequest) (_result *InnerLimitRateCardsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InnerLimitRateCardsResponse{}
	_body, _err := client.InnerLimitRateCardsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - InnerStopCardsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InnerStopCardsResponse
func (client *Client) InnerStopCardsWithOptions(request *InnerStopCardsRequest, runtime *util.RuntimeOptions) (_result *InnerStopCardsResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.InnerApi)) {
		query["InnerApi"] = request.InnerApi
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Task)) {
		query["Task"] = request.Task
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InnerStopCards"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InnerStopCardsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - InnerStopCardsRequest
//
// @return InnerStopCardsResponse
func (client *Client) InnerStopCards(request *InnerStopCardsRequest) (_result *InnerStopCardsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InnerStopCardsResponse{}
	_body, _err := client.InnerStopCardsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询5G高速上云服务支持的APN列表
//
// @param request - ListAPNsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAPNsResponse
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

// Summary:
//
// 查询5G高速上云服务支持的APN列表
//
// @param request - ListAPNsRequest
//
// @return ListAPNsResponse
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

// Summary:
//
// 查询5G高速上云服务实例的授权规则列表
//
// @param request - ListAuthorizationRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAuthorizationRulesResponse
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

// Summary:
//
// 查询5G高速上云服务实例的授权规则列表
//
// @param request - ListAuthorizationRulesRequest
//
// @return ListAuthorizationRulesResponse
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

// Summary:
//
// 查询批量操作卡任务列表
//
// @param request - ListBatchOperateCardsTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBatchOperateCardsTasksResponse
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

// Summary:
//
// 查询批量操作卡任务列表
//
// @param request - ListBatchOperateCardsTasksRequest
//
// @return ListBatchOperateCardsTasksResponse
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

// Summary:
//
// 查询5G高速上云服务物联网卡区域解锁支持区域
//
// @param request - ListCardAreaLimitSupportAreaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCardAreaLimitSupportAreaResponse
func (client *Client) ListCardAreaLimitSupportAreaWithOptions(request *ListCardAreaLimitSupportAreaRequest, runtime *util.RuntimeOptions) (_result *ListCardAreaLimitSupportAreaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCardAreaLimitSupportArea"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCardAreaLimitSupportAreaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询5G高速上云服务物联网卡区域解锁支持区域
//
// @param request - ListCardAreaLimitSupportAreaRequest
//
// @return ListCardAreaLimitSupportAreaResponse
func (client *Client) ListCardAreaLimitSupportArea(request *ListCardAreaLimitSupportAreaRequest) (_result *ListCardAreaLimitSupportAreaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCardAreaLimitSupportAreaResponse{}
	_body, _err := client.ListCardAreaLimitSupportAreaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询5G高速上云服务实例的卡单日流量信息列表
//
// @param request - ListCardDayUsagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCardDayUsagesResponse
func (client *Client) ListCardDayUsagesWithOptions(request *ListCardDayUsagesRequest, runtime *util.RuntimeOptions) (_result *ListCardDayUsagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCardDayUsages"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCardDayUsagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询5G高速上云服务实例的卡单日流量信息列表
//
// @param request - ListCardDayUsagesRequest
//
// @return ListCardDayUsagesResponse
func (client *Client) ListCardDayUsages(request *ListCardDayUsagesRequest) (_result *ListCardDayUsagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCardDayUsagesResponse{}
	_body, _err := client.ListCardDayUsagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询5G高速上云服务实例的流量信息列表
//
// @param request - ListCardUsagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCardUsagesResponse
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

// Summary:
//
// 查询5G高速上云服务实例的流量信息列表
//
// @param request - ListCardUsagesRequest
//
// @return ListCardUsagesResponse
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

// Summary:
//
// 查询5G高速上云服务实例的连接列表
//
// @param request - ListCardsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCardsResponse
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

// Summary:
//
// 查询5G高速上云服务实例的连接列表
//
// @param request - ListCardsRequest
//
// @return ListCardsResponse
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

// Summary:
//
// 查询5G高速上云服务实例下的带宽包列表
//
// @param request - ListDataPackagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataPackagesResponse
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

// Summary:
//
// 查询5G高速上云服务实例下的带宽包列表
//
// @param request - ListDataPackagesRequest
//
// @return ListDataPackagesResponse
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

// Summary:
//
// 查询单卡诊断信息列表
//
// @param request - ListDiagnoseInfoForSingleCardRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDiagnoseInfoForSingleCardResponse
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

// Summary:
//
// 查询单卡诊断信息列表
//
// @param request - ListDiagnoseInfoForSingleCardRequest
//
// @return ListDiagnoseInfoForSingleCardResponse
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

// Summary:
//
// 查询5G高速上云服务实例组的授权规则列表
//
// @param request - ListGroupAuthorizationRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGroupAuthorizationRulesResponse
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

// Summary:
//
// 查询5G高速上云服务实例组的授权规则列表
//
// @param request - ListGroupAuthorizationRulesRequest
//
// @return ListGroupAuthorizationRulesResponse
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

// Summary:
//
// 查询iotcc实例回程路由
//
// @param request - ListIoTCloudConnectorBackhaulRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIoTCloudConnectorBackhaulRouteResponse
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

// Summary:
//
// 查询iotcc实例回程路由
//
// @param request - ListIoTCloudConnectorBackhaulRouteRequest
//
// @return ListIoTCloudConnectorBackhaulRouteResponse
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

// Summary:
//
// 查询5G高速上云服务实例下的订单列表
//
// @param request - ListOrdersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOrdersResponse
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

// Summary:
//
// 查询5G高速上云服务实例下的订单列表
//
// @param request - ListOrdersRequest
//
// @return ListOrdersResponse
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

// Summary:
//
// 查询5G高速上云服务支持的REGION
//
// @param request - ListRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRegionsResponse
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

// Summary:
//
// 查询5G高速上云服务支持的REGION
//
// @param request - ListRegionsRequest
//
// @return ListRegionsResponse
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

// Summary:
//
// 查询5G高速上云服务实例分组列表
//
// @param request - ListWirelessCloudConnectorGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWirelessCloudConnectorGroupsResponse
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

// Summary:
//
// 查询5G高速上云服务实例分组列表
//
// @param request - ListWirelessCloudConnectorGroupsRequest
//
// @return ListWirelessCloudConnectorGroupsResponse
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

// Summary:
//
// 查询5G高速上云服务实例列表
//
// @param request - ListWirelessCloudConnectorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWirelessCloudConnectorsResponse
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

// Summary:
//
// 查询5G高速上云服务实例列表
//
// @param request - ListWirelessCloudConnectorsRequest
//
// @return ListWirelessCloudConnectorsResponse
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

// Summary:
//
// 查询5G高速上云服务支持的可用区
//
// @param request - ListZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListZonesResponse
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

// Summary:
//
// 查询5G高速上云服务支持的可用区
//
// @param request - ListZonesRequest
//
// @return ListZonesResponse
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

// Summary:
//
// 运营商侧卡停机
//
// @param request - LockCardsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LockCardsResponse
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

// Summary:
//
// 运营商侧卡停机
//
// @param request - LockCardsRequest
//
// @return LockCardsResponse
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

// Summary:
//
// 修改5G高速上云服务实例属性
//
// @param request - ModifyWirelessCloudConnectorFeatureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyWirelessCloudConnectorFeatureResponse
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

// Summary:
//
// 修改5G高速上云服务实例属性
//
// @param request - ModifyWirelessCloudConnectorFeatureRequest
//
// @return ModifyWirelessCloudConnectorFeatureResponse
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

// Summary:
//
// 开通5G高速上云服务
//
// @param request - OpenCc5gServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenCc5gServiceResponse
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

// Summary:
//
// 开通5G高速上云服务
//
// @param request - OpenCc5gServiceRequest
//
// @return OpenCc5gServiceResponse
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

// Summary:
//
// 物联网卡换绑解锁操作
//
// @param request - RebindCardsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RebindCardsResponse
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

// Summary:
//
// 物联网卡换绑解锁操作
//
// @param request - RebindCardsRequest
//
// @return RebindCardsResponse
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

// Summary:
//
// 从组里移除5G高速上云服务实例
//
// @param request - RemoveWirelessCloudConnectorFromGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveWirelessCloudConnectorFromGroupResponse
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

// Summary:
//
// 从组里移除5G高速上云服务实例
//
// @param request - RemoveWirelessCloudConnectorFromGroupRequest
//
// @return RemoveWirelessCloudConnectorFromGroupResponse
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

// Summary:
//
// 重置单卡区域限制
//
// @param request - ResetAreaLimitCardsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetAreaLimitCardsResponse
func (client *Client) ResetAreaLimitCardsWithOptions(request *ResetAreaLimitCardsRequest, runtime *util.RuntimeOptions) (_result *ResetAreaLimitCardsResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Province)) {
		query["Province"] = request.Province
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetAreaLimitCards"),
		Version:     tea.String("2022-03-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetAreaLimitCardsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重置单卡区域限制
//
// @param request - ResetAreaLimitCardsRequest
//
// @return ResetAreaLimitCardsResponse
func (client *Client) ResetAreaLimitCards(request *ResetAreaLimitCardsRequest) (_result *ResetAreaLimitCardsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetAreaLimitCardsResponse{}
	_body, _err := client.ResetAreaLimitCardsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 激活单卡运营商侧卡状态恢复
//
// @param request - ResumeCardsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeCardsResponse
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

// Summary:
//
// 激活单卡运营商侧卡状态恢复
//
// @param request - ResumeCardsRequest
//
// @return ResumeCardsResponse
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

// Summary:
//
// 撤销分享授权网络连接
//
// @param request - RevokeNetLinkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeNetLinkResponse
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

// Summary:
//
// 撤销分享授权网络连接
//
// @param request - RevokeNetLinkRequest
//
// @return RevokeNetLinkResponse
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

// Summary:
//
// 运营商侧卡停机
//
// @param request - StopCardsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopCardsResponse
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

// Summary:
//
// 运营商侧卡停机
//
// @param request - StopCardsRequest
//
// @return StopCardsResponse
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

// Summary:
//
// 开启cciot单卡一键诊断
//
// @param request - SubmitDiagnoseTaskForSingleCardRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitDiagnoseTaskForSingleCardResponse
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

// Summary:
//
// 开启cciot单卡一键诊断
//
// @param request - SubmitDiagnoseTaskForSingleCardRequest
//
// @return SubmitDiagnoseTaskForSingleCardResponse
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

// Summary:
//
// 切换cc5g实例商业类型到商业版
//
// @param request - SwitchWirelessCloudConnectorToBusinessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchWirelessCloudConnectorToBusinessResponse
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

// Summary:
//
// 切换cc5g实例商业类型到商业版
//
// @param request - SwitchWirelessCloudConnectorToBusinessRequest
//
// @return SwitchWirelessCloudConnectorToBusinessResponse
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

// Summary:
//
// 批量解锁卡
//
// @param request - UnlockCardsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnlockCardsResponse
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

// Summary:
//
// 批量解锁卡
//
// @param request - UnlockCardsRequest
//
// @return UnlockCardsResponse
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

// Summary:
//
// 更新5G高速上云服务实例的授权规则
//
// @param request - UpdateAuthorizationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAuthorizationRuleResponse
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

// Summary:
//
// 更新5G高速上云服务实例的授权规则
//
// @param request - UpdateAuthorizationRuleRequest
//
// @return UpdateAuthorizationRuleResponse
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

// Summary:
//
// 修改批量操作卡任务
//
// @param request - UpdateBatchOperateCardsTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBatchOperateCardsTaskResponse
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

// Summary:
//
// 修改批量操作卡任务
//
// @param request - UpdateBatchOperateCardsTaskRequest
//
// @return UpdateBatchOperateCardsTaskResponse
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

// Summary:
//
// 修改5G高速上云服务实例的连接描述和名称信息
//
// @param request - UpdateCardRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCardResponse
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

// Summary:
//
// 修改5G高速上云服务实例的连接描述和名称信息
//
// @param request - UpdateCardRequest
//
// @return UpdateCardResponse
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

// Summary:
//
// 更新5G高速上云服务实例的DNS授权规则
//
// @param request - UpdateDNSAuthorizationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDNSAuthorizationRuleResponse
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

// Summary:
//
// 更新5G高速上云服务实例的DNS授权规则
//
// @param request - UpdateDNSAuthorizationRuleRequest
//
// @return UpdateDNSAuthorizationRuleResponse
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

// Summary:
//
// 更新5G高速上云服务实例组的授权规则
//
// @param request - UpdateGroupAuthorizationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGroupAuthorizationRuleResponse
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

// Summary:
//
// 更新5G高速上云服务实例组的授权规则
//
// @param request - UpdateGroupAuthorizationRuleRequest
//
// @return UpdateGroupAuthorizationRuleResponse
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

// Summary:
//
// 更新5G高速上云服务实例组的DNS授权规则
//
// @param request - UpdateGroupDnsAuthorizationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGroupDnsAuthorizationRuleResponse
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

// Summary:
//
// 更新5G高速上云服务实例组的DNS授权规则
//
// @param request - UpdateGroupDnsAuthorizationRuleRequest
//
// @return UpdateGroupDnsAuthorizationRuleResponse
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

// Summary:
//
// 修改5G高速上云服务实例
//
// @param request - UpdateWirelessCloudConnectorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWirelessCloudConnectorResponse
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

// Summary:
//
// 修改5G高速上云服务实例
//
// @param request - UpdateWirelessCloudConnectorRequest
//
// @return UpdateWirelessCloudConnectorResponse
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

// Summary:
//
// 修改5G高速上云服务实例分组
//
// @param request - UpdateWirelessCloudConnectorGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWirelessCloudConnectorGroupResponse
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

// Summary:
//
// 修改5G高速上云服务实例分组
//
// @param request - UpdateWirelessCloudConnectorGroupRequest
//
// @return UpdateWirelessCloudConnectorGroupResponse
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
