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

type AddCidrToConnectionPoolRequest struct {
	Cidrs               []*string `json:"Cidrs,omitempty" xml:"Cidrs,omitempty" type:"Repeated"`
	ClientToken         *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConnectionPoolId    *string   `json:"ConnectionPoolId,omitempty" xml:"ConnectionPoolId,omitempty"`
	DryRun              *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorId *string   `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	RegionId            *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddCidrToConnectionPoolRequest) String() string {
	return tea.Prettify(s)
}

func (s AddCidrToConnectionPoolRequest) GoString() string {
	return s.String()
}

func (s *AddCidrToConnectionPoolRequest) SetCidrs(v []*string) *AddCidrToConnectionPoolRequest {
	s.Cidrs = v
	return s
}

func (s *AddCidrToConnectionPoolRequest) SetClientToken(v string) *AddCidrToConnectionPoolRequest {
	s.ClientToken = &v
	return s
}

func (s *AddCidrToConnectionPoolRequest) SetConnectionPoolId(v string) *AddCidrToConnectionPoolRequest {
	s.ConnectionPoolId = &v
	return s
}

func (s *AddCidrToConnectionPoolRequest) SetDryRun(v bool) *AddCidrToConnectionPoolRequest {
	s.DryRun = &v
	return s
}

func (s *AddCidrToConnectionPoolRequest) SetIoTCloudConnectorId(v string) *AddCidrToConnectionPoolRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *AddCidrToConnectionPoolRequest) SetRegionId(v string) *AddCidrToConnectionPoolRequest {
	s.RegionId = &v
	return s
}

type AddCidrToConnectionPoolResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddCidrToConnectionPoolResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddCidrToConnectionPoolResponseBody) GoString() string {
	return s.String()
}

func (s *AddCidrToConnectionPoolResponseBody) SetRequestId(v string) *AddCidrToConnectionPoolResponseBody {
	s.RequestId = &v
	return s
}

type AddCidrToConnectionPoolResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddCidrToConnectionPoolResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddCidrToConnectionPoolResponse) String() string {
	return tea.Prettify(s)
}

func (s AddCidrToConnectionPoolResponse) GoString() string {
	return s.String()
}

func (s *AddCidrToConnectionPoolResponse) SetHeaders(v map[string]*string) *AddCidrToConnectionPoolResponse {
	s.Headers = v
	return s
}

func (s *AddCidrToConnectionPoolResponse) SetStatusCode(v int32) *AddCidrToConnectionPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCidrToConnectionPoolResponse) SetBody(v *AddCidrToConnectionPoolResponseBody) *AddCidrToConnectionPoolResponse {
	s.Body = v
	return s
}

type AddIoTCloudConnectorToGroupRequest struct {
	ClientToken              *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                   *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorGroupId *string   `json:"IoTCloudConnectorGroupId,omitempty" xml:"IoTCloudConnectorGroupId,omitempty"`
	IoTCloudConnectorId      []*string `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty" type:"Repeated"`
	RegionId                 *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddIoTCloudConnectorToGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddIoTCloudConnectorToGroupRequest) GoString() string {
	return s.String()
}

func (s *AddIoTCloudConnectorToGroupRequest) SetClientToken(v string) *AddIoTCloudConnectorToGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *AddIoTCloudConnectorToGroupRequest) SetDryRun(v bool) *AddIoTCloudConnectorToGroupRequest {
	s.DryRun = &v
	return s
}

func (s *AddIoTCloudConnectorToGroupRequest) SetIoTCloudConnectorGroupId(v string) *AddIoTCloudConnectorToGroupRequest {
	s.IoTCloudConnectorGroupId = &v
	return s
}

func (s *AddIoTCloudConnectorToGroupRequest) SetIoTCloudConnectorId(v []*string) *AddIoTCloudConnectorToGroupRequest {
	s.IoTCloudConnectorId = v
	return s
}

func (s *AddIoTCloudConnectorToGroupRequest) SetRegionId(v string) *AddIoTCloudConnectorToGroupRequest {
	s.RegionId = &v
	return s
}

type AddIoTCloudConnectorToGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddIoTCloudConnectorToGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddIoTCloudConnectorToGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddIoTCloudConnectorToGroupResponseBody) SetRequestId(v string) *AddIoTCloudConnectorToGroupResponseBody {
	s.RequestId = &v
	return s
}

type AddIoTCloudConnectorToGroupResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddIoTCloudConnectorToGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddIoTCloudConnectorToGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AddIoTCloudConnectorToGroupResponse) GoString() string {
	return s.String()
}

func (s *AddIoTCloudConnectorToGroupResponse) SetHeaders(v map[string]*string) *AddIoTCloudConnectorToGroupResponse {
	s.Headers = v
	return s
}

func (s *AddIoTCloudConnectorToGroupResponse) SetStatusCode(v int32) *AddIoTCloudConnectorToGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddIoTCloudConnectorToGroupResponse) SetBody(v *AddIoTCloudConnectorToGroupResponseBody) *AddIoTCloudConnectorToGroupResponse {
	s.Body = v
	return s
}

type AssociateIpWithConnectionPoolRequest struct {
	ClientToken         *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConnectionPoolId    *string   `json:"ConnectionPoolId,omitempty" xml:"ConnectionPoolId,omitempty"`
	DryRun              *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorId *string   `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	Ips                 []*string `json:"Ips,omitempty" xml:"Ips,omitempty" type:"Repeated"`
	IpsFilePath         *string   `json:"IpsFilePath,omitempty" xml:"IpsFilePath,omitempty"`
	RegionId            *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AssociateIpWithConnectionPoolRequest) String() string {
	return tea.Prettify(s)
}

func (s AssociateIpWithConnectionPoolRequest) GoString() string {
	return s.String()
}

func (s *AssociateIpWithConnectionPoolRequest) SetClientToken(v string) *AssociateIpWithConnectionPoolRequest {
	s.ClientToken = &v
	return s
}

func (s *AssociateIpWithConnectionPoolRequest) SetConnectionPoolId(v string) *AssociateIpWithConnectionPoolRequest {
	s.ConnectionPoolId = &v
	return s
}

func (s *AssociateIpWithConnectionPoolRequest) SetDryRun(v bool) *AssociateIpWithConnectionPoolRequest {
	s.DryRun = &v
	return s
}

func (s *AssociateIpWithConnectionPoolRequest) SetIoTCloudConnectorId(v string) *AssociateIpWithConnectionPoolRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *AssociateIpWithConnectionPoolRequest) SetIps(v []*string) *AssociateIpWithConnectionPoolRequest {
	s.Ips = v
	return s
}

func (s *AssociateIpWithConnectionPoolRequest) SetIpsFilePath(v string) *AssociateIpWithConnectionPoolRequest {
	s.IpsFilePath = &v
	return s
}

func (s *AssociateIpWithConnectionPoolRequest) SetRegionId(v string) *AssociateIpWithConnectionPoolRequest {
	s.RegionId = &v
	return s
}

type AssociateIpWithConnectionPoolResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateIpWithConnectionPoolResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssociateIpWithConnectionPoolResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateIpWithConnectionPoolResponseBody) SetRequestId(v string) *AssociateIpWithConnectionPoolResponseBody {
	s.RequestId = &v
	return s
}

type AssociateIpWithConnectionPoolResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AssociateIpWithConnectionPoolResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AssociateIpWithConnectionPoolResponse) String() string {
	return tea.Prettify(s)
}

func (s AssociateIpWithConnectionPoolResponse) GoString() string {
	return s.String()
}

func (s *AssociateIpWithConnectionPoolResponse) SetHeaders(v map[string]*string) *AssociateIpWithConnectionPoolResponse {
	s.Headers = v
	return s
}

func (s *AssociateIpWithConnectionPoolResponse) SetStatusCode(v int32) *AssociateIpWithConnectionPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateIpWithConnectionPoolResponse) SetBody(v *AssociateIpWithConnectionPoolResponseBody) *AssociateIpWithConnectionPoolResponse {
	s.Body = v
	return s
}

type AssociateVSwitchWithIoTCloudConnectorRequest struct {
	ClientToken         *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun              *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorId *string   `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	RegionId            *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VSwitchList         []*string `json:"VSwitchList,omitempty" xml:"VSwitchList,omitempty" type:"Repeated"`
	VpcId               *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s AssociateVSwitchWithIoTCloudConnectorRequest) String() string {
	return tea.Prettify(s)
}

func (s AssociateVSwitchWithIoTCloudConnectorRequest) GoString() string {
	return s.String()
}

func (s *AssociateVSwitchWithIoTCloudConnectorRequest) SetClientToken(v string) *AssociateVSwitchWithIoTCloudConnectorRequest {
	s.ClientToken = &v
	return s
}

func (s *AssociateVSwitchWithIoTCloudConnectorRequest) SetDryRun(v bool) *AssociateVSwitchWithIoTCloudConnectorRequest {
	s.DryRun = &v
	return s
}

func (s *AssociateVSwitchWithIoTCloudConnectorRequest) SetIoTCloudConnectorId(v string) *AssociateVSwitchWithIoTCloudConnectorRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *AssociateVSwitchWithIoTCloudConnectorRequest) SetRegionId(v string) *AssociateVSwitchWithIoTCloudConnectorRequest {
	s.RegionId = &v
	return s
}

func (s *AssociateVSwitchWithIoTCloudConnectorRequest) SetVSwitchList(v []*string) *AssociateVSwitchWithIoTCloudConnectorRequest {
	s.VSwitchList = v
	return s
}

func (s *AssociateVSwitchWithIoTCloudConnectorRequest) SetVpcId(v string) *AssociateVSwitchWithIoTCloudConnectorRequest {
	s.VpcId = &v
	return s
}

type AssociateVSwitchWithIoTCloudConnectorResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateVSwitchWithIoTCloudConnectorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssociateVSwitchWithIoTCloudConnectorResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateVSwitchWithIoTCloudConnectorResponseBody) SetRequestId(v string) *AssociateVSwitchWithIoTCloudConnectorResponseBody {
	s.RequestId = &v
	return s
}

type AssociateVSwitchWithIoTCloudConnectorResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AssociateVSwitchWithIoTCloudConnectorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AssociateVSwitchWithIoTCloudConnectorResponse) String() string {
	return tea.Prettify(s)
}

func (s AssociateVSwitchWithIoTCloudConnectorResponse) GoString() string {
	return s.String()
}

func (s *AssociateVSwitchWithIoTCloudConnectorResponse) SetHeaders(v map[string]*string) *AssociateVSwitchWithIoTCloudConnectorResponse {
	s.Headers = v
	return s
}

func (s *AssociateVSwitchWithIoTCloudConnectorResponse) SetStatusCode(v int32) *AssociateVSwitchWithIoTCloudConnectorResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateVSwitchWithIoTCloudConnectorResponse) SetBody(v *AssociateVSwitchWithIoTCloudConnectorResponseBody) *AssociateVSwitchWithIoTCloudConnectorResponse {
	s.Body = v
	return s
}

type ConfirmIoTCloudConnectorRequest struct {
	ClientToken         *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConfirmStatus       *string `json:"ConfirmStatus,omitempty" xml:"ConfirmStatus,omitempty"`
	DryRun              *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorId *string `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ConfirmIoTCloudConnectorRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfirmIoTCloudConnectorRequest) GoString() string {
	return s.String()
}

func (s *ConfirmIoTCloudConnectorRequest) SetClientToken(v string) *ConfirmIoTCloudConnectorRequest {
	s.ClientToken = &v
	return s
}

func (s *ConfirmIoTCloudConnectorRequest) SetConfirmStatus(v string) *ConfirmIoTCloudConnectorRequest {
	s.ConfirmStatus = &v
	return s
}

func (s *ConfirmIoTCloudConnectorRequest) SetDryRun(v bool) *ConfirmIoTCloudConnectorRequest {
	s.DryRun = &v
	return s
}

func (s *ConfirmIoTCloudConnectorRequest) SetIoTCloudConnectorId(v string) *ConfirmIoTCloudConnectorRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *ConfirmIoTCloudConnectorRequest) SetRegionId(v string) *ConfirmIoTCloudConnectorRequest {
	s.RegionId = &v
	return s
}

type ConfirmIoTCloudConnectorResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s ConfirmIoTCloudConnectorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfirmIoTCloudConnectorResponseBody) GoString() string {
	return s.String()
}

func (s *ConfirmIoTCloudConnectorResponseBody) SetRequestId(v string) *ConfirmIoTCloudConnectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfirmIoTCloudConnectorResponseBody) SetResourceId(v string) *ConfirmIoTCloudConnectorResponseBody {
	s.ResourceId = &v
	return s
}

type ConfirmIoTCloudConnectorResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ConfirmIoTCloudConnectorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfirmIoTCloudConnectorResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfirmIoTCloudConnectorResponse) GoString() string {
	return s.String()
}

func (s *ConfirmIoTCloudConnectorResponse) SetHeaders(v map[string]*string) *ConfirmIoTCloudConnectorResponse {
	s.Headers = v
	return s
}

func (s *ConfirmIoTCloudConnectorResponse) SetStatusCode(v int32) *ConfirmIoTCloudConnectorResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfirmIoTCloudConnectorResponse) SetBody(v *ConfirmIoTCloudConnectorResponseBody) *ConfirmIoTCloudConnectorResponse {
	s.Body = v
	return s
}

type CreateAuthorizationRuleRequest struct {
	AuthorizationRuleDescription *string   `json:"AuthorizationRuleDescription,omitempty" xml:"AuthorizationRuleDescription,omitempty"`
	AuthorizationRuleName        *string   `json:"AuthorizationRuleName,omitempty" xml:"AuthorizationRuleName,omitempty"`
	ClientToken                  *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Destination                  *string   `json:"Destination,omitempty" xml:"Destination,omitempty"`
	DestinationPort              *string   `json:"DestinationPort,omitempty" xml:"DestinationPort,omitempty"`
	DestinationType              *string   `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	DryRun                       *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorId          *string   `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	Policy                       *string   `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Protocol                     *string   `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	RegionId                     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SourceCidrs                  []*string `json:"SourceCidrs,omitempty" xml:"SourceCidrs,omitempty" type:"Repeated"`
}

func (s CreateAuthorizationRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAuthorizationRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateAuthorizationRuleRequest) SetAuthorizationRuleDescription(v string) *CreateAuthorizationRuleRequest {
	s.AuthorizationRuleDescription = &v
	return s
}

func (s *CreateAuthorizationRuleRequest) SetAuthorizationRuleName(v string) *CreateAuthorizationRuleRequest {
	s.AuthorizationRuleName = &v
	return s
}

func (s *CreateAuthorizationRuleRequest) SetClientToken(v string) *CreateAuthorizationRuleRequest {
	s.ClientToken = &v
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

func (s *CreateAuthorizationRuleRequest) SetIoTCloudConnectorId(v string) *CreateAuthorizationRuleRequest {
	s.IoTCloudConnectorId = &v
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

func (s *CreateAuthorizationRuleRequest) SetRegionId(v string) *CreateAuthorizationRuleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAuthorizationRuleRequest) SetSourceCidrs(v []*string) *CreateAuthorizationRuleRequest {
	s.SourceCidrs = v
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

type CreateAuthorizationRulesRequest struct {
	AuthorizationRules  []*CreateAuthorizationRulesRequestAuthorizationRules `json:"AuthorizationRules,omitempty" xml:"AuthorizationRules,omitempty" type:"Repeated"`
	ClientToken         *string                                              `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun              *bool                                                `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorId *string                                              `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	RegionId            *string                                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateAuthorizationRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAuthorizationRulesRequest) GoString() string {
	return s.String()
}

func (s *CreateAuthorizationRulesRequest) SetAuthorizationRules(v []*CreateAuthorizationRulesRequestAuthorizationRules) *CreateAuthorizationRulesRequest {
	s.AuthorizationRules = v
	return s
}

func (s *CreateAuthorizationRulesRequest) SetClientToken(v string) *CreateAuthorizationRulesRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAuthorizationRulesRequest) SetDryRun(v bool) *CreateAuthorizationRulesRequest {
	s.DryRun = &v
	return s
}

func (s *CreateAuthorizationRulesRequest) SetIoTCloudConnectorId(v string) *CreateAuthorizationRulesRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *CreateAuthorizationRulesRequest) SetRegionId(v string) *CreateAuthorizationRulesRequest {
	s.RegionId = &v
	return s
}

type CreateAuthorizationRulesRequestAuthorizationRules struct {
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Destination     *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	DestinationPort *string `json:"DestinationPort,omitempty" xml:"DestinationPort,omitempty"`
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Policy          *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Protocol        *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	SourceCidr      *string `json:"SourceCidr,omitempty" xml:"SourceCidr,omitempty"`
}

func (s CreateAuthorizationRulesRequestAuthorizationRules) String() string {
	return tea.Prettify(s)
}

func (s CreateAuthorizationRulesRequestAuthorizationRules) GoString() string {
	return s.String()
}

func (s *CreateAuthorizationRulesRequestAuthorizationRules) SetDescription(v string) *CreateAuthorizationRulesRequestAuthorizationRules {
	s.Description = &v
	return s
}

func (s *CreateAuthorizationRulesRequestAuthorizationRules) SetDestination(v string) *CreateAuthorizationRulesRequestAuthorizationRules {
	s.Destination = &v
	return s
}

func (s *CreateAuthorizationRulesRequestAuthorizationRules) SetDestinationPort(v string) *CreateAuthorizationRulesRequestAuthorizationRules {
	s.DestinationPort = &v
	return s
}

func (s *CreateAuthorizationRulesRequestAuthorizationRules) SetDestinationType(v string) *CreateAuthorizationRulesRequestAuthorizationRules {
	s.DestinationType = &v
	return s
}

func (s *CreateAuthorizationRulesRequestAuthorizationRules) SetName(v string) *CreateAuthorizationRulesRequestAuthorizationRules {
	s.Name = &v
	return s
}

func (s *CreateAuthorizationRulesRequestAuthorizationRules) SetPolicy(v string) *CreateAuthorizationRulesRequestAuthorizationRules {
	s.Policy = &v
	return s
}

func (s *CreateAuthorizationRulesRequestAuthorizationRules) SetProtocol(v string) *CreateAuthorizationRulesRequestAuthorizationRules {
	s.Protocol = &v
	return s
}

func (s *CreateAuthorizationRulesRequestAuthorizationRules) SetSourceCidr(v string) *CreateAuthorizationRulesRequestAuthorizationRules {
	s.SourceCidr = &v
	return s
}

type CreateAuthorizationRulesResponseBody struct {
	AuthorizationRuleIds []*string `json:"AuthorizationRuleIds,omitempty" xml:"AuthorizationRuleIds,omitempty" type:"Repeated"`
	RequestId            *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAuthorizationRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAuthorizationRulesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAuthorizationRulesResponseBody) SetAuthorizationRuleIds(v []*string) *CreateAuthorizationRulesResponseBody {
	s.AuthorizationRuleIds = v
	return s
}

func (s *CreateAuthorizationRulesResponseBody) SetRequestId(v string) *CreateAuthorizationRulesResponseBody {
	s.RequestId = &v
	return s
}

type CreateAuthorizationRulesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAuthorizationRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAuthorizationRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAuthorizationRulesResponse) GoString() string {
	return s.String()
}

func (s *CreateAuthorizationRulesResponse) SetHeaders(v map[string]*string) *CreateAuthorizationRulesResponse {
	s.Headers = v
	return s
}

func (s *CreateAuthorizationRulesResponse) SetStatusCode(v int32) *CreateAuthorizationRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAuthorizationRulesResponse) SetBody(v *CreateAuthorizationRulesResponseBody) *CreateAuthorizationRulesResponse {
	s.Body = v
	return s
}

type CreateConnectionPoolRequest struct {
	Cidrs                     []*string `json:"Cidrs,omitempty" xml:"Cidrs,omitempty" type:"Repeated"`
	ClientToken               *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConnectionPoolDescription *string   `json:"ConnectionPoolDescription,omitempty" xml:"ConnectionPoolDescription,omitempty"`
	ConnectionPoolName        *string   `json:"ConnectionPoolName,omitempty" xml:"ConnectionPoolName,omitempty"`
	Count                     *int64    `json:"Count,omitempty" xml:"Count,omitempty"`
	DryRun                    *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorId       *string   `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	RegionId                  *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateConnectionPoolRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateConnectionPoolRequest) GoString() string {
	return s.String()
}

func (s *CreateConnectionPoolRequest) SetCidrs(v []*string) *CreateConnectionPoolRequest {
	s.Cidrs = v
	return s
}

func (s *CreateConnectionPoolRequest) SetClientToken(v string) *CreateConnectionPoolRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateConnectionPoolRequest) SetConnectionPoolDescription(v string) *CreateConnectionPoolRequest {
	s.ConnectionPoolDescription = &v
	return s
}

func (s *CreateConnectionPoolRequest) SetConnectionPoolName(v string) *CreateConnectionPoolRequest {
	s.ConnectionPoolName = &v
	return s
}

func (s *CreateConnectionPoolRequest) SetCount(v int64) *CreateConnectionPoolRequest {
	s.Count = &v
	return s
}

func (s *CreateConnectionPoolRequest) SetDryRun(v bool) *CreateConnectionPoolRequest {
	s.DryRun = &v
	return s
}

func (s *CreateConnectionPoolRequest) SetIoTCloudConnectorId(v string) *CreateConnectionPoolRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *CreateConnectionPoolRequest) SetRegionId(v string) *CreateConnectionPoolRequest {
	s.RegionId = &v
	return s
}

type CreateConnectionPoolResponseBody struct {
	ConnectionPoolId *string `json:"ConnectionPoolId,omitempty" xml:"ConnectionPoolId,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateConnectionPoolResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateConnectionPoolResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConnectionPoolResponseBody) SetConnectionPoolId(v string) *CreateConnectionPoolResponseBody {
	s.ConnectionPoolId = &v
	return s
}

func (s *CreateConnectionPoolResponseBody) SetRequestId(v string) *CreateConnectionPoolResponseBody {
	s.RequestId = &v
	return s
}

type CreateConnectionPoolResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateConnectionPoolResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateConnectionPoolResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateConnectionPoolResponse) GoString() string {
	return s.String()
}

func (s *CreateConnectionPoolResponse) SetHeaders(v map[string]*string) *CreateConnectionPoolResponse {
	s.Headers = v
	return s
}

func (s *CreateConnectionPoolResponse) SetStatusCode(v int32) *CreateConnectionPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateConnectionPoolResponse) SetBody(v *CreateConnectionPoolResponseBody) *CreateConnectionPoolResponse {
	s.Body = v
	return s
}

type CreateDNSServiceRuleRequest struct {
	ClientToken               *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DNSServiceRuleDescription *string `json:"DNSServiceRuleDescription,omitempty" xml:"DNSServiceRuleDescription,omitempty"`
	DNSServiceRuleName        *string `json:"DNSServiceRuleName,omitempty" xml:"DNSServiceRuleName,omitempty"`
	Destination               *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	DryRun                    *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorId       *string `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	RegionId                  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceType               *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	Source                    *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s CreateDNSServiceRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDNSServiceRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateDNSServiceRuleRequest) SetClientToken(v string) *CreateDNSServiceRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDNSServiceRuleRequest) SetDNSServiceRuleDescription(v string) *CreateDNSServiceRuleRequest {
	s.DNSServiceRuleDescription = &v
	return s
}

func (s *CreateDNSServiceRuleRequest) SetDNSServiceRuleName(v string) *CreateDNSServiceRuleRequest {
	s.DNSServiceRuleName = &v
	return s
}

func (s *CreateDNSServiceRuleRequest) SetDestination(v string) *CreateDNSServiceRuleRequest {
	s.Destination = &v
	return s
}

func (s *CreateDNSServiceRuleRequest) SetDryRun(v bool) *CreateDNSServiceRuleRequest {
	s.DryRun = &v
	return s
}

func (s *CreateDNSServiceRuleRequest) SetIoTCloudConnectorId(v string) *CreateDNSServiceRuleRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *CreateDNSServiceRuleRequest) SetRegionId(v string) *CreateDNSServiceRuleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDNSServiceRuleRequest) SetServiceType(v string) *CreateDNSServiceRuleRequest {
	s.ServiceType = &v
	return s
}

func (s *CreateDNSServiceRuleRequest) SetSource(v string) *CreateDNSServiceRuleRequest {
	s.Source = &v
	return s
}

type CreateDNSServiceRuleResponseBody struct {
	DNSServiceRuleId *string `json:"DNSServiceRuleId,omitempty" xml:"DNSServiceRuleId,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDNSServiceRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDNSServiceRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDNSServiceRuleResponseBody) SetDNSServiceRuleId(v string) *CreateDNSServiceRuleResponseBody {
	s.DNSServiceRuleId = &v
	return s
}

func (s *CreateDNSServiceRuleResponseBody) SetRequestId(v string) *CreateDNSServiceRuleResponseBody {
	s.RequestId = &v
	return s
}

type CreateDNSServiceRuleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDNSServiceRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDNSServiceRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDNSServiceRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateDNSServiceRuleResponse) SetHeaders(v map[string]*string) *CreateDNSServiceRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateDNSServiceRuleResponse) SetStatusCode(v int32) *CreateDNSServiceRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDNSServiceRuleResponse) SetBody(v *CreateDNSServiceRuleResponseBody) *CreateDNSServiceRuleResponse {
	s.Body = v
	return s
}

type CreateGroupAuthorizationRuleRequest struct {
	AuthorizationRuleDescription *string   `json:"AuthorizationRuleDescription,omitempty" xml:"AuthorizationRuleDescription,omitempty"`
	AuthorizationRuleName        *string   `json:"AuthorizationRuleName,omitempty" xml:"AuthorizationRuleName,omitempty"`
	ClientToken                  *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Destination                  *string   `json:"Destination,omitempty" xml:"Destination,omitempty"`
	DestinationPort              *string   `json:"DestinationPort,omitempty" xml:"DestinationPort,omitempty"`
	DestinationType              *string   `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	DryRun                       *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorGroupId     *string   `json:"IoTCloudConnectorGroupId,omitempty" xml:"IoTCloudConnectorGroupId,omitempty"`
	Policy                       *string   `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Protocol                     *string   `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	RegionId                     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SourceCidrs                  []*string `json:"SourceCidrs,omitempty" xml:"SourceCidrs,omitempty" type:"Repeated"`
	Type                         *string   `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateGroupAuthorizationRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupAuthorizationRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupAuthorizationRuleRequest) SetAuthorizationRuleDescription(v string) *CreateGroupAuthorizationRuleRequest {
	s.AuthorizationRuleDescription = &v
	return s
}

func (s *CreateGroupAuthorizationRuleRequest) SetAuthorizationRuleName(v string) *CreateGroupAuthorizationRuleRequest {
	s.AuthorizationRuleName = &v
	return s
}

func (s *CreateGroupAuthorizationRuleRequest) SetClientToken(v string) *CreateGroupAuthorizationRuleRequest {
	s.ClientToken = &v
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

func (s *CreateGroupAuthorizationRuleRequest) SetIoTCloudConnectorGroupId(v string) *CreateGroupAuthorizationRuleRequest {
	s.IoTCloudConnectorGroupId = &v
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

func (s *CreateGroupAuthorizationRuleRequest) SetRegionId(v string) *CreateGroupAuthorizationRuleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateGroupAuthorizationRuleRequest) SetSourceCidrs(v []*string) *CreateGroupAuthorizationRuleRequest {
	s.SourceCidrs = v
	return s
}

func (s *CreateGroupAuthorizationRuleRequest) SetType(v string) *CreateGroupAuthorizationRuleRequest {
	s.Type = &v
	return s
}

type CreateGroupAuthorizationRuleResponseBody struct {
	AuthorizationRuleId      *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	IoTCloudConnectorGroupId *string `json:"IoTCloudConnectorGroupId,omitempty" xml:"IoTCloudConnectorGroupId,omitempty"`
	RequestId                *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *CreateGroupAuthorizationRuleResponseBody) SetIoTCloudConnectorGroupId(v string) *CreateGroupAuthorizationRuleResponseBody {
	s.IoTCloudConnectorGroupId = &v
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

type CreateGroupDNSServiceRuleRequest struct {
	ClientToken               *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DNSServiceRuleDescription *string `json:"DNSServiceRuleDescription,omitempty" xml:"DNSServiceRuleDescription,omitempty"`
	DNSServiceRuleName        *string `json:"DNSServiceRuleName,omitempty" xml:"DNSServiceRuleName,omitempty"`
	Destination               *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	DryRun                    *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorGroupId  *string `json:"IoTCloudConnectorGroupId,omitempty" xml:"IoTCloudConnectorGroupId,omitempty"`
	RegionId                  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceType               *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	Source                    *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s CreateGroupDNSServiceRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupDNSServiceRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupDNSServiceRuleRequest) SetClientToken(v string) *CreateGroupDNSServiceRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateGroupDNSServiceRuleRequest) SetDNSServiceRuleDescription(v string) *CreateGroupDNSServiceRuleRequest {
	s.DNSServiceRuleDescription = &v
	return s
}

func (s *CreateGroupDNSServiceRuleRequest) SetDNSServiceRuleName(v string) *CreateGroupDNSServiceRuleRequest {
	s.DNSServiceRuleName = &v
	return s
}

func (s *CreateGroupDNSServiceRuleRequest) SetDestination(v string) *CreateGroupDNSServiceRuleRequest {
	s.Destination = &v
	return s
}

func (s *CreateGroupDNSServiceRuleRequest) SetDryRun(v bool) *CreateGroupDNSServiceRuleRequest {
	s.DryRun = &v
	return s
}

func (s *CreateGroupDNSServiceRuleRequest) SetIoTCloudConnectorGroupId(v string) *CreateGroupDNSServiceRuleRequest {
	s.IoTCloudConnectorGroupId = &v
	return s
}

func (s *CreateGroupDNSServiceRuleRequest) SetRegionId(v string) *CreateGroupDNSServiceRuleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateGroupDNSServiceRuleRequest) SetServiceType(v string) *CreateGroupDNSServiceRuleRequest {
	s.ServiceType = &v
	return s
}

func (s *CreateGroupDNSServiceRuleRequest) SetSource(v string) *CreateGroupDNSServiceRuleRequest {
	s.Source = &v
	return s
}

type CreateGroupDNSServiceRuleResponseBody struct {
	DNSServiceRuleId         *string `json:"DNSServiceRuleId,omitempty" xml:"DNSServiceRuleId,omitempty"`
	IoTCloudConnectorGroupId *string `json:"IoTCloudConnectorGroupId,omitempty" xml:"IoTCloudConnectorGroupId,omitempty"`
	RequestId                *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateGroupDNSServiceRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupDNSServiceRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupDNSServiceRuleResponseBody) SetDNSServiceRuleId(v string) *CreateGroupDNSServiceRuleResponseBody {
	s.DNSServiceRuleId = &v
	return s
}

func (s *CreateGroupDNSServiceRuleResponseBody) SetIoTCloudConnectorGroupId(v string) *CreateGroupDNSServiceRuleResponseBody {
	s.IoTCloudConnectorGroupId = &v
	return s
}

func (s *CreateGroupDNSServiceRuleResponseBody) SetRequestId(v string) *CreateGroupDNSServiceRuleResponseBody {
	s.RequestId = &v
	return s
}

type CreateGroupDNSServiceRuleResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateGroupDNSServiceRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGroupDNSServiceRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupDNSServiceRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupDNSServiceRuleResponse) SetHeaders(v map[string]*string) *CreateGroupDNSServiceRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateGroupDNSServiceRuleResponse) SetStatusCode(v int32) *CreateGroupDNSServiceRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGroupDNSServiceRuleResponse) SetBody(v *CreateGroupDNSServiceRuleResponseBody) *CreateGroupDNSServiceRuleResponse {
	s.Body = v
	return s
}

type CreateGroupIpMappingRuleRequest struct {
	ClientToken              *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DestinationIp            *string `json:"DestinationIp,omitempty" xml:"DestinationIp,omitempty"`
	DryRun                   *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorGroupId *string `json:"IoTCloudConnectorGroupId,omitempty" xml:"IoTCloudConnectorGroupId,omitempty"`
	IpMappingRuleDescription *string `json:"IpMappingRuleDescription,omitempty" xml:"IpMappingRuleDescription,omitempty"`
	IpMappingRuleName        *string `json:"IpMappingRuleName,omitempty" xml:"IpMappingRuleName,omitempty"`
	MappingIp                *string `json:"MappingIp,omitempty" xml:"MappingIp,omitempty"`
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateGroupIpMappingRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupIpMappingRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupIpMappingRuleRequest) SetClientToken(v string) *CreateGroupIpMappingRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateGroupIpMappingRuleRequest) SetDestinationIp(v string) *CreateGroupIpMappingRuleRequest {
	s.DestinationIp = &v
	return s
}

func (s *CreateGroupIpMappingRuleRequest) SetDryRun(v bool) *CreateGroupIpMappingRuleRequest {
	s.DryRun = &v
	return s
}

func (s *CreateGroupIpMappingRuleRequest) SetIoTCloudConnectorGroupId(v string) *CreateGroupIpMappingRuleRequest {
	s.IoTCloudConnectorGroupId = &v
	return s
}

func (s *CreateGroupIpMappingRuleRequest) SetIpMappingRuleDescription(v string) *CreateGroupIpMappingRuleRequest {
	s.IpMappingRuleDescription = &v
	return s
}

func (s *CreateGroupIpMappingRuleRequest) SetIpMappingRuleName(v string) *CreateGroupIpMappingRuleRequest {
	s.IpMappingRuleName = &v
	return s
}

func (s *CreateGroupIpMappingRuleRequest) SetMappingIp(v string) *CreateGroupIpMappingRuleRequest {
	s.MappingIp = &v
	return s
}

func (s *CreateGroupIpMappingRuleRequest) SetRegionId(v string) *CreateGroupIpMappingRuleRequest {
	s.RegionId = &v
	return s
}

type CreateGroupIpMappingRuleResponseBody struct {
	GroupIpMappingRuleId *string `json:"GroupIpMappingRuleId,omitempty" xml:"GroupIpMappingRuleId,omitempty"`
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateGroupIpMappingRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupIpMappingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupIpMappingRuleResponseBody) SetGroupIpMappingRuleId(v string) *CreateGroupIpMappingRuleResponseBody {
	s.GroupIpMappingRuleId = &v
	return s
}

func (s *CreateGroupIpMappingRuleResponseBody) SetRequestId(v string) *CreateGroupIpMappingRuleResponseBody {
	s.RequestId = &v
	return s
}

type CreateGroupIpMappingRuleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateGroupIpMappingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGroupIpMappingRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupIpMappingRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupIpMappingRuleResponse) SetHeaders(v map[string]*string) *CreateGroupIpMappingRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateGroupIpMappingRuleResponse) SetStatusCode(v int32) *CreateGroupIpMappingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGroupIpMappingRuleResponse) SetBody(v *CreateGroupIpMappingRuleResponseBody) *CreateGroupIpMappingRuleResponse {
	s.Body = v
	return s
}

type CreateIoTCloudConnectorRequest struct {
	APN                          *string `json:"APN,omitempty" xml:"APN,omitempty"`
	ClientToken                  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	ISP                          *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	IoTCloudConnectorDescription *string `json:"IoTCloudConnectorDescription,omitempty" xml:"IoTCloudConnectorDescription,omitempty"`
	IoTCloudConnectorName        *string `json:"IoTCloudConnectorName,omitempty" xml:"IoTCloudConnectorName,omitempty"`
	RegionId                     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceUid                  *int64  `json:"ResourceUid,omitempty" xml:"ResourceUid,omitempty"`
	Type                         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Deprecated
	WildcardDomainEnabled *bool `json:"WildcardDomainEnabled,omitempty" xml:"WildcardDomainEnabled,omitempty"`
}

func (s CreateIoTCloudConnectorRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIoTCloudConnectorRequest) GoString() string {
	return s.String()
}

func (s *CreateIoTCloudConnectorRequest) SetAPN(v string) *CreateIoTCloudConnectorRequest {
	s.APN = &v
	return s
}

func (s *CreateIoTCloudConnectorRequest) SetClientToken(v string) *CreateIoTCloudConnectorRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateIoTCloudConnectorRequest) SetDryRun(v bool) *CreateIoTCloudConnectorRequest {
	s.DryRun = &v
	return s
}

func (s *CreateIoTCloudConnectorRequest) SetISP(v string) *CreateIoTCloudConnectorRequest {
	s.ISP = &v
	return s
}

func (s *CreateIoTCloudConnectorRequest) SetIoTCloudConnectorDescription(v string) *CreateIoTCloudConnectorRequest {
	s.IoTCloudConnectorDescription = &v
	return s
}

func (s *CreateIoTCloudConnectorRequest) SetIoTCloudConnectorName(v string) *CreateIoTCloudConnectorRequest {
	s.IoTCloudConnectorName = &v
	return s
}

func (s *CreateIoTCloudConnectorRequest) SetRegionId(v string) *CreateIoTCloudConnectorRequest {
	s.RegionId = &v
	return s
}

func (s *CreateIoTCloudConnectorRequest) SetResourceUid(v int64) *CreateIoTCloudConnectorRequest {
	s.ResourceUid = &v
	return s
}

func (s *CreateIoTCloudConnectorRequest) SetType(v string) *CreateIoTCloudConnectorRequest {
	s.Type = &v
	return s
}

func (s *CreateIoTCloudConnectorRequest) SetWildcardDomainEnabled(v bool) *CreateIoTCloudConnectorRequest {
	s.WildcardDomainEnabled = &v
	return s
}

type CreateIoTCloudConnectorResponseBody struct {
	IoTCloudConnectorId *string `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateIoTCloudConnectorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIoTCloudConnectorResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIoTCloudConnectorResponseBody) SetIoTCloudConnectorId(v string) *CreateIoTCloudConnectorResponseBody {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *CreateIoTCloudConnectorResponseBody) SetRequestId(v string) *CreateIoTCloudConnectorResponseBody {
	s.RequestId = &v
	return s
}

type CreateIoTCloudConnectorResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateIoTCloudConnectorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateIoTCloudConnectorResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIoTCloudConnectorResponse) GoString() string {
	return s.String()
}

func (s *CreateIoTCloudConnectorResponse) SetHeaders(v map[string]*string) *CreateIoTCloudConnectorResponse {
	s.Headers = v
	return s
}

func (s *CreateIoTCloudConnectorResponse) SetStatusCode(v int32) *CreateIoTCloudConnectorResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIoTCloudConnectorResponse) SetBody(v *CreateIoTCloudConnectorResponseBody) *CreateIoTCloudConnectorResponse {
	s.Body = v
	return s
}

type CreateIoTCloudConnectorBackhaulRouteRequest struct {
	ClientToken         *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun              *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorId *string `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *CreateIoTCloudConnectorBackhaulRouteRequest) SetIoTCloudConnectorId(v string) *CreateIoTCloudConnectorBackhaulRouteRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *CreateIoTCloudConnectorBackhaulRouteRequest) SetRegionId(v string) *CreateIoTCloudConnectorBackhaulRouteRequest {
	s.RegionId = &v
	return s
}

type CreateIoTCloudConnectorBackhaulRouteResponseBody struct {
	IoTCloudConnectorId *string `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateIoTCloudConnectorBackhaulRouteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIoTCloudConnectorBackhaulRouteResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIoTCloudConnectorBackhaulRouteResponseBody) SetIoTCloudConnectorId(v string) *CreateIoTCloudConnectorBackhaulRouteResponseBody {
	s.IoTCloudConnectorId = &v
	return s
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

type CreateIoTCloudConnectorGroupRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DryRun      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateIoTCloudConnectorGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIoTCloudConnectorGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateIoTCloudConnectorGroupRequest) SetClientToken(v string) *CreateIoTCloudConnectorGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateIoTCloudConnectorGroupRequest) SetDescription(v string) *CreateIoTCloudConnectorGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateIoTCloudConnectorGroupRequest) SetDryRun(v bool) *CreateIoTCloudConnectorGroupRequest {
	s.DryRun = &v
	return s
}

func (s *CreateIoTCloudConnectorGroupRequest) SetName(v string) *CreateIoTCloudConnectorGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateIoTCloudConnectorGroupRequest) SetRegionId(v string) *CreateIoTCloudConnectorGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateIoTCloudConnectorGroupRequest) SetType(v string) *CreateIoTCloudConnectorGroupRequest {
	s.Type = &v
	return s
}

type CreateIoTCloudConnectorGroupResponseBody struct {
	IoTCloudConnectorGroupId *string `json:"IoTCloudConnectorGroupId,omitempty" xml:"IoTCloudConnectorGroupId,omitempty"`
	RequestId                *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateIoTCloudConnectorGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIoTCloudConnectorGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIoTCloudConnectorGroupResponseBody) SetIoTCloudConnectorGroupId(v string) *CreateIoTCloudConnectorGroupResponseBody {
	s.IoTCloudConnectorGroupId = &v
	return s
}

func (s *CreateIoTCloudConnectorGroupResponseBody) SetRequestId(v string) *CreateIoTCloudConnectorGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateIoTCloudConnectorGroupResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateIoTCloudConnectorGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateIoTCloudConnectorGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIoTCloudConnectorGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateIoTCloudConnectorGroupResponse) SetHeaders(v map[string]*string) *CreateIoTCloudConnectorGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateIoTCloudConnectorGroupResponse) SetStatusCode(v int32) *CreateIoTCloudConnectorGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIoTCloudConnectorGroupResponse) SetBody(v *CreateIoTCloudConnectorGroupResponseBody) *CreateIoTCloudConnectorGroupResponse {
	s.Body = v
	return s
}

type CreateIpMappingRuleRequest struct {
	ClientToken              *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DestinationIp            *string `json:"DestinationIp,omitempty" xml:"DestinationIp,omitempty"`
	DryRun                   *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorId      *string `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	IpMappingRuleDescription *string `json:"IpMappingRuleDescription,omitempty" xml:"IpMappingRuleDescription,omitempty"`
	IpMappingRuleName        *string `json:"IpMappingRuleName,omitempty" xml:"IpMappingRuleName,omitempty"`
	MappingIp                *string `json:"MappingIp,omitempty" xml:"MappingIp,omitempty"`
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateIpMappingRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIpMappingRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateIpMappingRuleRequest) SetClientToken(v string) *CreateIpMappingRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateIpMappingRuleRequest) SetDestinationIp(v string) *CreateIpMappingRuleRequest {
	s.DestinationIp = &v
	return s
}

func (s *CreateIpMappingRuleRequest) SetDryRun(v bool) *CreateIpMappingRuleRequest {
	s.DryRun = &v
	return s
}

func (s *CreateIpMappingRuleRequest) SetIoTCloudConnectorId(v string) *CreateIpMappingRuleRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *CreateIpMappingRuleRequest) SetIpMappingRuleDescription(v string) *CreateIpMappingRuleRequest {
	s.IpMappingRuleDescription = &v
	return s
}

func (s *CreateIpMappingRuleRequest) SetIpMappingRuleName(v string) *CreateIpMappingRuleRequest {
	s.IpMappingRuleName = &v
	return s
}

func (s *CreateIpMappingRuleRequest) SetMappingIp(v string) *CreateIpMappingRuleRequest {
	s.MappingIp = &v
	return s
}

func (s *CreateIpMappingRuleRequest) SetRegionId(v string) *CreateIpMappingRuleRequest {
	s.RegionId = &v
	return s
}

type CreateIpMappingRuleResponseBody struct {
	IpMappingRuleId *string `json:"IpMappingRuleId,omitempty" xml:"IpMappingRuleId,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateIpMappingRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIpMappingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIpMappingRuleResponseBody) SetIpMappingRuleId(v string) *CreateIpMappingRuleResponseBody {
	s.IpMappingRuleId = &v
	return s
}

func (s *CreateIpMappingRuleResponseBody) SetRequestId(v string) *CreateIpMappingRuleResponseBody {
	s.RequestId = &v
	return s
}

type CreateIpMappingRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateIpMappingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateIpMappingRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIpMappingRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateIpMappingRuleResponse) SetHeaders(v map[string]*string) *CreateIpMappingRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateIpMappingRuleResponse) SetStatusCode(v int32) *CreateIpMappingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIpMappingRuleResponse) SetBody(v *CreateIpMappingRuleResponseBody) *CreateIpMappingRuleResponse {
	s.Body = v
	return s
}

type CreateServiceRequest struct {
	ClientToken         *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun              *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorId *string `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceDescription  *string `json:"ServiceDescription,omitempty" xml:"ServiceDescription,omitempty"`
	ServiceName         *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s CreateServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceRequest) SetClientToken(v string) *CreateServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServiceRequest) SetDryRun(v bool) *CreateServiceRequest {
	s.DryRun = &v
	return s
}

func (s *CreateServiceRequest) SetIoTCloudConnectorId(v string) *CreateServiceRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *CreateServiceRequest) SetRegionId(v string) *CreateServiceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateServiceRequest) SetServiceDescription(v string) *CreateServiceRequest {
	s.ServiceDescription = &v
	return s
}

func (s *CreateServiceRequest) SetServiceName(v string) *CreateServiceRequest {
	s.ServiceName = &v
	return s
}

type CreateServiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s CreateServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceResponseBody) SetRequestId(v string) *CreateServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceResponseBody) SetServiceId(v string) *CreateServiceResponseBody {
	s.ServiceId = &v
	return s
}

type CreateServiceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceResponse) SetHeaders(v map[string]*string) *CreateServiceResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceResponse) SetStatusCode(v int32) *CreateServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceResponse) SetBody(v *CreateServiceResponseBody) *CreateServiceResponse {
	s.Body = v
	return s
}

type CreateServiceEntryRequest struct {
	ClientToken             *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                  *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorId     *string `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	RegionId                *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceEntryDescription *string `json:"ServiceEntryDescription,omitempty" xml:"ServiceEntryDescription,omitempty"`
	ServiceEntryName        *string `json:"ServiceEntryName,omitempty" xml:"ServiceEntryName,omitempty"`
	ServiceId               *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	Target                  *string `json:"Target,omitempty" xml:"Target,omitempty"`
	TargetType              *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s CreateServiceEntryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceEntryRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceEntryRequest) SetClientToken(v string) *CreateServiceEntryRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServiceEntryRequest) SetDryRun(v bool) *CreateServiceEntryRequest {
	s.DryRun = &v
	return s
}

func (s *CreateServiceEntryRequest) SetIoTCloudConnectorId(v string) *CreateServiceEntryRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *CreateServiceEntryRequest) SetRegionId(v string) *CreateServiceEntryRequest {
	s.RegionId = &v
	return s
}

func (s *CreateServiceEntryRequest) SetServiceEntryDescription(v string) *CreateServiceEntryRequest {
	s.ServiceEntryDescription = &v
	return s
}

func (s *CreateServiceEntryRequest) SetServiceEntryName(v string) *CreateServiceEntryRequest {
	s.ServiceEntryName = &v
	return s
}

func (s *CreateServiceEntryRequest) SetServiceId(v string) *CreateServiceEntryRequest {
	s.ServiceId = &v
	return s
}

func (s *CreateServiceEntryRequest) SetTarget(v string) *CreateServiceEntryRequest {
	s.Target = &v
	return s
}

func (s *CreateServiceEntryRequest) SetTargetType(v string) *CreateServiceEntryRequest {
	s.TargetType = &v
	return s
}

type CreateServiceEntryResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceEntryId *string `json:"ServiceEntryId,omitempty" xml:"ServiceEntryId,omitempty"`
}

func (s CreateServiceEntryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceEntryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceEntryResponseBody) SetRequestId(v string) *CreateServiceEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceEntryResponseBody) SetServiceEntryId(v string) *CreateServiceEntryResponseBody {
	s.ServiceEntryId = &v
	return s
}

type CreateServiceEntryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateServiceEntryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateServiceEntryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceEntryResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceEntryResponse) SetHeaders(v map[string]*string) *CreateServiceEntryResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceEntryResponse) SetStatusCode(v int32) *CreateServiceEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceEntryResponse) SetBody(v *CreateServiceEntryResponseBody) *CreateServiceEntryResponse {
	s.Body = v
	return s
}

type DeleteAuthorizationRuleRequest struct {
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	ClientToken         *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun              *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorId *string `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *DeleteAuthorizationRuleRequest) SetIoTCloudConnectorId(v string) *DeleteAuthorizationRuleRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *DeleteAuthorizationRuleRequest) SetRegionId(v string) *DeleteAuthorizationRuleRequest {
	s.RegionId = &v
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

type DeleteAuthorizationRulesRequest struct {
	AuthorizationRuleIds []*string `json:"AuthorizationRuleIds,omitempty" xml:"AuthorizationRuleIds,omitempty" type:"Repeated"`
	ClientToken          *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun               *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorId  *string   `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteAuthorizationRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAuthorizationRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteAuthorizationRulesRequest) SetAuthorizationRuleIds(v []*string) *DeleteAuthorizationRulesRequest {
	s.AuthorizationRuleIds = v
	return s
}

func (s *DeleteAuthorizationRulesRequest) SetClientToken(v string) *DeleteAuthorizationRulesRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteAuthorizationRulesRequest) SetDryRun(v bool) *DeleteAuthorizationRulesRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteAuthorizationRulesRequest) SetIoTCloudConnectorId(v string) *DeleteAuthorizationRulesRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *DeleteAuthorizationRulesRequest) SetRegionId(v string) *DeleteAuthorizationRulesRequest {
	s.RegionId = &v
	return s
}

type DeleteAuthorizationRulesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAuthorizationRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAuthorizationRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAuthorizationRulesResponseBody) SetRequestId(v string) *DeleteAuthorizationRulesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAuthorizationRulesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAuthorizationRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAuthorizationRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAuthorizationRulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteAuthorizationRulesResponse) SetHeaders(v map[string]*string) *DeleteAuthorizationRulesResponse {
	s.Headers = v
	return s
}

func (s *DeleteAuthorizationRulesResponse) SetStatusCode(v int32) *DeleteAuthorizationRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAuthorizationRulesResponse) SetBody(v *DeleteAuthorizationRulesResponseBody) *DeleteAuthorizationRulesResponse {
	s.Body = v
	return s
}

type DeleteConnectionPoolRequest struct {
	ClientToken         *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConnectionPoolId    *string `json:"ConnectionPoolId,omitempty" xml:"ConnectionPoolId,omitempty"`
	DryRun              *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorId *string `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteConnectionPoolRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteConnectionPoolRequest) GoString() string {
	return s.String()
}

func (s *DeleteConnectionPoolRequest) SetClientToken(v string) *DeleteConnectionPoolRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteConnectionPoolRequest) SetConnectionPoolId(v string) *DeleteConnectionPoolRequest {
	s.ConnectionPoolId = &v
	return s
}

func (s *DeleteConnectionPoolRequest) SetDryRun(v bool) *DeleteConnectionPoolRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteConnectionPoolRequest) SetIoTCloudConnectorId(v string) *DeleteConnectionPoolRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *DeleteConnectionPoolRequest) SetRegionId(v string) *DeleteConnectionPoolRequest {
	s.RegionId = &v
	return s
}

type DeleteConnectionPoolResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteConnectionPoolResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteConnectionPoolResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConnectionPoolResponseBody) SetRequestId(v string) *DeleteConnectionPoolResponseBody {
	s.RequestId = &v
	return s
}

type DeleteConnectionPoolResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteConnectionPoolResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteConnectionPoolResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteConnectionPoolResponse) GoString() string {
	return s.String()
}

func (s *DeleteConnectionPoolResponse) SetHeaders(v map[string]*string) *DeleteConnectionPoolResponse {
	s.Headers = v
	return s
}

func (s *DeleteConnectionPoolResponse) SetStatusCode(v int32) *DeleteConnectionPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteConnectionPoolResponse) SetBody(v *DeleteConnectionPoolResponseBody) *DeleteConnectionPoolResponse {
	s.Body = v
	return s
}

type DeleteDNSServiceRuleRequest struct {
	ClientToken         *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DNSServiceRuleId    *string `json:"DNSServiceRuleId,omitempty" xml:"DNSServiceRuleId,omitempty"`
	DryRun              *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorId *string `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteDNSServiceRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDNSServiceRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteDNSServiceRuleRequest) SetClientToken(v string) *DeleteDNSServiceRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteDNSServiceRuleRequest) SetDNSServiceRuleId(v string) *DeleteDNSServiceRuleRequest {
	s.DNSServiceRuleId = &v
	return s
}

func (s *DeleteDNSServiceRuleRequest) SetDryRun(v bool) *DeleteDNSServiceRuleRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteDNSServiceRuleRequest) SetIoTCloudConnectorId(v string) *DeleteDNSServiceRuleRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *DeleteDNSServiceRuleRequest) SetRegionId(v string) *DeleteDNSServiceRuleRequest {
	s.RegionId = &v
	return s
}

type DeleteDNSServiceRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDNSServiceRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDNSServiceRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDNSServiceRuleResponseBody) SetRequestId(v string) *DeleteDNSServiceRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDNSServiceRuleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDNSServiceRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDNSServiceRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDNSServiceRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteDNSServiceRuleResponse) SetHeaders(v map[string]*string) *DeleteDNSServiceRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteDNSServiceRuleResponse) SetStatusCode(v int32) *DeleteDNSServiceRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDNSServiceRuleResponse) SetBody(v *DeleteDNSServiceRuleResponseBody) *DeleteDNSServiceRuleResponse {
	s.Body = v
	return s
}

type DeleteGroupAuthorizationRuleRequest struct {
	AuthorizationRuleId      *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	ClientToken              *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                   *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorGroupId *string `json:"IoTCloudConnectorGroupId,omitempty" xml:"IoTCloudConnectorGroupId,omitempty"`
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *DeleteGroupAuthorizationRuleRequest) SetIoTCloudConnectorGroupId(v string) *DeleteGroupAuthorizationRuleRequest {
	s.IoTCloudConnectorGroupId = &v
	return s
}

func (s *DeleteGroupAuthorizationRuleRequest) SetRegionId(v string) *DeleteGroupAuthorizationRuleRequest {
	s.RegionId = &v
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

type DeleteGroupDNSServiceRuleRequest struct {
	ClientToken              *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DNSServiceRuleId         *string `json:"DNSServiceRuleId,omitempty" xml:"DNSServiceRuleId,omitempty"`
	DryRun                   *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorGroupId *string `json:"IoTCloudConnectorGroupId,omitempty" xml:"IoTCloudConnectorGroupId,omitempty"`
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteGroupDNSServiceRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupDNSServiceRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteGroupDNSServiceRuleRequest) SetClientToken(v string) *DeleteGroupDNSServiceRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteGroupDNSServiceRuleRequest) SetDNSServiceRuleId(v string) *DeleteGroupDNSServiceRuleRequest {
	s.DNSServiceRuleId = &v
	return s
}

func (s *DeleteGroupDNSServiceRuleRequest) SetDryRun(v bool) *DeleteGroupDNSServiceRuleRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteGroupDNSServiceRuleRequest) SetIoTCloudConnectorGroupId(v string) *DeleteGroupDNSServiceRuleRequest {
	s.IoTCloudConnectorGroupId = &v
	return s
}

func (s *DeleteGroupDNSServiceRuleRequest) SetRegionId(v string) *DeleteGroupDNSServiceRuleRequest {
	s.RegionId = &v
	return s
}

type DeleteGroupDNSServiceRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGroupDNSServiceRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupDNSServiceRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGroupDNSServiceRuleResponseBody) SetRequestId(v string) *DeleteGroupDNSServiceRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteGroupDNSServiceRuleResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteGroupDNSServiceRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteGroupDNSServiceRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupDNSServiceRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteGroupDNSServiceRuleResponse) SetHeaders(v map[string]*string) *DeleteGroupDNSServiceRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteGroupDNSServiceRuleResponse) SetStatusCode(v int32) *DeleteGroupDNSServiceRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGroupDNSServiceRuleResponse) SetBody(v *DeleteGroupDNSServiceRuleResponseBody) *DeleteGroupDNSServiceRuleResponse {
	s.Body = v
	return s
}

type DeleteGroupIpMappingRuleRequest struct {
	ClientToken              *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                   *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	GroupIpMappingRuleId     *string `json:"GroupIpMappingRuleId,omitempty" xml:"GroupIpMappingRuleId,omitempty"`
	IoTCloudConnectorGroupId *string `json:"IoTCloudConnectorGroupId,omitempty" xml:"IoTCloudConnectorGroupId,omitempty"`
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteGroupIpMappingRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupIpMappingRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteGroupIpMappingRuleRequest) SetClientToken(v string) *DeleteGroupIpMappingRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteGroupIpMappingRuleRequest) SetDryRun(v bool) *DeleteGroupIpMappingRuleRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteGroupIpMappingRuleRequest) SetGroupIpMappingRuleId(v string) *DeleteGroupIpMappingRuleRequest {
	s.GroupIpMappingRuleId = &v
	return s
}

func (s *DeleteGroupIpMappingRuleRequest) SetIoTCloudConnectorGroupId(v string) *DeleteGroupIpMappingRuleRequest {
	s.IoTCloudConnectorGroupId = &v
	return s
}

func (s *DeleteGroupIpMappingRuleRequest) SetRegionId(v string) *DeleteGroupIpMappingRuleRequest {
	s.RegionId = &v
	return s
}

type DeleteGroupIpMappingRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGroupIpMappingRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupIpMappingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGroupIpMappingRuleResponseBody) SetRequestId(v string) *DeleteGroupIpMappingRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteGroupIpMappingRuleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteGroupIpMappingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteGroupIpMappingRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupIpMappingRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteGroupIpMappingRuleResponse) SetHeaders(v map[string]*string) *DeleteGroupIpMappingRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteGroupIpMappingRuleResponse) SetStatusCode(v int32) *DeleteGroupIpMappingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGroupIpMappingRuleResponse) SetBody(v *DeleteGroupIpMappingRuleResponseBody) *DeleteGroupIpMappingRuleResponse {
	s.Body = v
	return s
}

type DeleteIoTCloudConnectorRequest struct {
	ClientToken         *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun              *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorId *string `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteIoTCloudConnectorRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIoTCloudConnectorRequest) GoString() string {
	return s.String()
}

func (s *DeleteIoTCloudConnectorRequest) SetClientToken(v string) *DeleteIoTCloudConnectorRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteIoTCloudConnectorRequest) SetDryRun(v bool) *DeleteIoTCloudConnectorRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteIoTCloudConnectorRequest) SetIoTCloudConnectorId(v string) *DeleteIoTCloudConnectorRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *DeleteIoTCloudConnectorRequest) SetRegionId(v string) *DeleteIoTCloudConnectorRequest {
	s.RegionId = &v
	return s
}

type DeleteIoTCloudConnectorResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIoTCloudConnectorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteIoTCloudConnectorResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIoTCloudConnectorResponseBody) SetRequestId(v string) *DeleteIoTCloudConnectorResponseBody {
	s.RequestId = &v
	return s
}

type DeleteIoTCloudConnectorResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteIoTCloudConnectorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteIoTCloudConnectorResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIoTCloudConnectorResponse) GoString() string {
	return s.String()
}

func (s *DeleteIoTCloudConnectorResponse) SetHeaders(v map[string]*string) *DeleteIoTCloudConnectorResponse {
	s.Headers = v
	return s
}

func (s *DeleteIoTCloudConnectorResponse) SetStatusCode(v int32) *DeleteIoTCloudConnectorResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIoTCloudConnectorResponse) SetBody(v *DeleteIoTCloudConnectorResponseBody) *DeleteIoTCloudConnectorResponse {
	s.Body = v
	return s
}

type DeleteIoTCloudConnectorGroupRequest struct {
	ClientToken              *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                   *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorGroupId *string `json:"IoTCloudConnectorGroupId,omitempty" xml:"IoTCloudConnectorGroupId,omitempty"`
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteIoTCloudConnectorGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIoTCloudConnectorGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteIoTCloudConnectorGroupRequest) SetClientToken(v string) *DeleteIoTCloudConnectorGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteIoTCloudConnectorGroupRequest) SetDryRun(v bool) *DeleteIoTCloudConnectorGroupRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteIoTCloudConnectorGroupRequest) SetIoTCloudConnectorGroupId(v string) *DeleteIoTCloudConnectorGroupRequest {
	s.IoTCloudConnectorGroupId = &v
	return s
}

func (s *DeleteIoTCloudConnectorGroupRequest) SetRegionId(v string) *DeleteIoTCloudConnectorGroupRequest {
	s.RegionId = &v
	return s
}

type DeleteIoTCloudConnectorGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIoTCloudConnectorGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteIoTCloudConnectorGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIoTCloudConnectorGroupResponseBody) SetRequestId(v string) *DeleteIoTCloudConnectorGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteIoTCloudConnectorGroupResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteIoTCloudConnectorGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteIoTCloudConnectorGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIoTCloudConnectorGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteIoTCloudConnectorGroupResponse) SetHeaders(v map[string]*string) *DeleteIoTCloudConnectorGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteIoTCloudConnectorGroupResponse) SetStatusCode(v int32) *DeleteIoTCloudConnectorGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIoTCloudConnectorGroupResponse) SetBody(v *DeleteIoTCloudConnectorGroupResponseBody) *DeleteIoTCloudConnectorGroupResponse {
	s.Body = v
	return s
}

type DeleteIoTCloudConnetorBackhaulRouteRequest struct {
	ClientToken         *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun              *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorId *string `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteIoTCloudConnetorBackhaulRouteRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIoTCloudConnetorBackhaulRouteRequest) GoString() string {
	return s.String()
}

func (s *DeleteIoTCloudConnetorBackhaulRouteRequest) SetClientToken(v string) *DeleteIoTCloudConnetorBackhaulRouteRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteIoTCloudConnetorBackhaulRouteRequest) SetDryRun(v bool) *DeleteIoTCloudConnetorBackhaulRouteRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteIoTCloudConnetorBackhaulRouteRequest) SetIoTCloudConnectorId(v string) *DeleteIoTCloudConnetorBackhaulRouteRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *DeleteIoTCloudConnetorBackhaulRouteRequest) SetRegionId(v string) *DeleteIoTCloudConnetorBackhaulRouteRequest {
	s.RegionId = &v
	return s
}

type DeleteIoTCloudConnetorBackhaulRouteResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIoTCloudConnetorBackhaulRouteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteIoTCloudConnetorBackhaulRouteResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIoTCloudConnetorBackhaulRouteResponseBody) SetRequestId(v string) *DeleteIoTCloudConnetorBackhaulRouteResponseBody {
	s.RequestId = &v
	return s
}

type DeleteIoTCloudConnetorBackhaulRouteResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteIoTCloudConnetorBackhaulRouteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteIoTCloudConnetorBackhaulRouteResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIoTCloudConnetorBackhaulRouteResponse) GoString() string {
	return s.String()
}

func (s *DeleteIoTCloudConnetorBackhaulRouteResponse) SetHeaders(v map[string]*string) *DeleteIoTCloudConnetorBackhaulRouteResponse {
	s.Headers = v
	return s
}

func (s *DeleteIoTCloudConnetorBackhaulRouteResponse) SetStatusCode(v int32) *DeleteIoTCloudConnetorBackhaulRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIoTCloudConnetorBackhaulRouteResponse) SetBody(v *DeleteIoTCloudConnetorBackhaulRouteResponseBody) *DeleteIoTCloudConnetorBackhaulRouteResponse {
	s.Body = v
	return s
}

type DeleteIpMappingRuleRequest struct {
	ClientToken         *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun              *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorId *string `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	IpMappingRuleId     *string `json:"IpMappingRuleId,omitempty" xml:"IpMappingRuleId,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteIpMappingRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpMappingRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteIpMappingRuleRequest) SetClientToken(v string) *DeleteIpMappingRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteIpMappingRuleRequest) SetDryRun(v bool) *DeleteIpMappingRuleRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteIpMappingRuleRequest) SetIoTCloudConnectorId(v string) *DeleteIpMappingRuleRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *DeleteIpMappingRuleRequest) SetIpMappingRuleId(v string) *DeleteIpMappingRuleRequest {
	s.IpMappingRuleId = &v
	return s
}

func (s *DeleteIpMappingRuleRequest) SetRegionId(v string) *DeleteIpMappingRuleRequest {
	s.RegionId = &v
	return s
}

type DeleteIpMappingRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIpMappingRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpMappingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIpMappingRuleResponseBody) SetRequestId(v string) *DeleteIpMappingRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteIpMappingRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteIpMappingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteIpMappingRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpMappingRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteIpMappingRuleResponse) SetHeaders(v map[string]*string) *DeleteIpMappingRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteIpMappingRuleResponse) SetStatusCode(v int32) *DeleteIpMappingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIpMappingRuleResponse) SetBody(v *DeleteIpMappingRuleResponseBody) *DeleteIpMappingRuleResponse {
	s.Body = v
	return s
}

type DeleteServiceRequest struct {
	ClientToken         *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun              *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorId *string `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceId           *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s DeleteServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceRequest) SetClientToken(v string) *DeleteServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteServiceRequest) SetDryRun(v bool) *DeleteServiceRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteServiceRequest) SetIoTCloudConnectorId(v string) *DeleteServiceRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *DeleteServiceRequest) SetRegionId(v string) *DeleteServiceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteServiceRequest) SetServiceId(v string) *DeleteServiceRequest {
	s.ServiceId = &v
	return s
}

type DeleteServiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceResponseBody) SetRequestId(v string) *DeleteServiceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteServiceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceResponse) SetHeaders(v map[string]*string) *DeleteServiceResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceResponse) SetStatusCode(v int32) *DeleteServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceResponse) SetBody(v *DeleteServiceResponseBody) *DeleteServiceResponse {
	s.Body = v
	return s
}

type DeleteServiceEntryRequest struct {
	ClientToken         *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun              *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorId *string `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceEntryId      *string `json:"ServiceEntryId,omitempty" xml:"ServiceEntryId,omitempty"`
	ServiceId           *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s DeleteServiceEntryRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceEntryRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceEntryRequest) SetClientToken(v string) *DeleteServiceEntryRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteServiceEntryRequest) SetDryRun(v bool) *DeleteServiceEntryRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteServiceEntryRequest) SetIoTCloudConnectorId(v string) *DeleteServiceEntryRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *DeleteServiceEntryRequest) SetRegionId(v string) *DeleteServiceEntryRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteServiceEntryRequest) SetServiceEntryId(v string) *DeleteServiceEntryRequest {
	s.ServiceEntryId = &v
	return s
}

func (s *DeleteServiceEntryRequest) SetServiceId(v string) *DeleteServiceEntryRequest {
	s.ServiceId = &v
	return s
}

type DeleteServiceEntryResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServiceEntryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceEntryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceEntryResponseBody) SetRequestId(v string) *DeleteServiceEntryResponseBody {
	s.RequestId = &v
	return s
}

type DeleteServiceEntryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteServiceEntryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteServiceEntryResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceEntryResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceEntryResponse) SetHeaders(v map[string]*string) *DeleteServiceEntryResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceEntryResponse) SetStatusCode(v int32) *DeleteServiceEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceEntryResponse) SetBody(v *DeleteServiceEntryResponseBody) *DeleteServiceEntryResponse {
	s.Body = v
	return s
}

type DisableIoTCloudConnectorAccessLogRequest struct {
	ClientToken         *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun              *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorId *string `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DisableIoTCloudConnectorAccessLogRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableIoTCloudConnectorAccessLogRequest) GoString() string {
	return s.String()
}

func (s *DisableIoTCloudConnectorAccessLogRequest) SetClientToken(v string) *DisableIoTCloudConnectorAccessLogRequest {
	s.ClientToken = &v
	return s
}

func (s *DisableIoTCloudConnectorAccessLogRequest) SetDryRun(v bool) *DisableIoTCloudConnectorAccessLogRequest {
	s.DryRun = &v
	return s
}

func (s *DisableIoTCloudConnectorAccessLogRequest) SetIoTCloudConnectorId(v string) *DisableIoTCloudConnectorAccessLogRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *DisableIoTCloudConnectorAccessLogRequest) SetRegionId(v string) *DisableIoTCloudConnectorAccessLogRequest {
	s.RegionId = &v
	return s
}

type DisableIoTCloudConnectorAccessLogResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableIoTCloudConnectorAccessLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableIoTCloudConnectorAccessLogResponseBody) GoString() string {
	return s.String()
}

func (s *DisableIoTCloudConnectorAccessLogResponseBody) SetRequestId(v string) *DisableIoTCloudConnectorAccessLogResponseBody {
	s.RequestId = &v
	return s
}

type DisableIoTCloudConnectorAccessLogResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisableIoTCloudConnectorAccessLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableIoTCloudConnectorAccessLogResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableIoTCloudConnectorAccessLogResponse) GoString() string {
	return s.String()
}

func (s *DisableIoTCloudConnectorAccessLogResponse) SetHeaders(v map[string]*string) *DisableIoTCloudConnectorAccessLogResponse {
	s.Headers = v
	return s
}

func (s *DisableIoTCloudConnectorAccessLogResponse) SetStatusCode(v int32) *DisableIoTCloudConnectorAccessLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableIoTCloudConnectorAccessLogResponse) SetBody(v *DisableIoTCloudConnectorAccessLogResponseBody) *DisableIoTCloudConnectorAccessLogResponse {
	s.Body = v
	return s
}

type DissociateIpFromConnectionPoolRequest struct {
	ClientToken         *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConnectionPoolId    *string   `json:"ConnectionPoolId,omitempty" xml:"ConnectionPoolId,omitempty"`
	DryRun              *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorId *string   `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	Ips                 []*string `json:"Ips,omitempty" xml:"Ips,omitempty" type:"Repeated"`
	IpsFilePath         *string   `json:"IpsFilePath,omitempty" xml:"IpsFilePath,omitempty"`
	RegionId            *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DissociateIpFromConnectionPoolRequest) String() string {
	return tea.Prettify(s)
}

func (s DissociateIpFromConnectionPoolRequest) GoString() string {
	return s.String()
}

func (s *DissociateIpFromConnectionPoolRequest) SetClientToken(v string) *DissociateIpFromConnectionPoolRequest {
	s.ClientToken = &v
	return s
}

func (s *DissociateIpFromConnectionPoolRequest) SetConnectionPoolId(v string) *DissociateIpFromConnectionPoolRequest {
	s.ConnectionPoolId = &v
	return s
}

func (s *DissociateIpFromConnectionPoolRequest) SetDryRun(v bool) *DissociateIpFromConnectionPoolRequest {
	s.DryRun = &v
	return s
}

func (s *DissociateIpFromConnectionPoolRequest) SetIoTCloudConnectorId(v string) *DissociateIpFromConnectionPoolRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *DissociateIpFromConnectionPoolRequest) SetIps(v []*string) *DissociateIpFromConnectionPoolRequest {
	s.Ips = v
	return s
}

func (s *DissociateIpFromConnectionPoolRequest) SetIpsFilePath(v string) *DissociateIpFromConnectionPoolRequest {
	s.IpsFilePath = &v
	return s
}

func (s *DissociateIpFromConnectionPoolRequest) SetRegionId(v string) *DissociateIpFromConnectionPoolRequest {
	s.RegionId = &v
	return s
}

type DissociateIpFromConnectionPoolResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DissociateIpFromConnectionPoolResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DissociateIpFromConnectionPoolResponseBody) GoString() string {
	return s.String()
}

func (s *DissociateIpFromConnectionPoolResponseBody) SetRequestId(v string) *DissociateIpFromConnectionPoolResponseBody {
	s.RequestId = &v
	return s
}

type DissociateIpFromConnectionPoolResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DissociateIpFromConnectionPoolResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DissociateIpFromConnectionPoolResponse) String() string {
	return tea.Prettify(s)
}

func (s DissociateIpFromConnectionPoolResponse) GoString() string {
	return s.String()
}

func (s *DissociateIpFromConnectionPoolResponse) SetHeaders(v map[string]*string) *DissociateIpFromConnectionPoolResponse {
	s.Headers = v
	return s
}

func (s *DissociateIpFromConnectionPoolResponse) SetStatusCode(v int32) *DissociateIpFromConnectionPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *DissociateIpFromConnectionPoolResponse) SetBody(v *DissociateIpFromConnectionPoolResponseBody) *DissociateIpFromConnectionPoolResponse {
	s.Body = v
	return s
}

type DissociateVSwitchFromIoTCloudConnectorRequest struct {
	ClientToken         *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun              *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorId *string `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DissociateVSwitchFromIoTCloudConnectorRequest) String() string {
	return tea.Prettify(s)
}

func (s DissociateVSwitchFromIoTCloudConnectorRequest) GoString() string {
	return s.String()
}

func (s *DissociateVSwitchFromIoTCloudConnectorRequest) SetClientToken(v string) *DissociateVSwitchFromIoTCloudConnectorRequest {
	s.ClientToken = &v
	return s
}

func (s *DissociateVSwitchFromIoTCloudConnectorRequest) SetDryRun(v bool) *DissociateVSwitchFromIoTCloudConnectorRequest {
	s.DryRun = &v
	return s
}

func (s *DissociateVSwitchFromIoTCloudConnectorRequest) SetIoTCloudConnectorId(v string) *DissociateVSwitchFromIoTCloudConnectorRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *DissociateVSwitchFromIoTCloudConnectorRequest) SetRegionId(v string) *DissociateVSwitchFromIoTCloudConnectorRequest {
	s.RegionId = &v
	return s
}

type DissociateVSwitchFromIoTCloudConnectorResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DissociateVSwitchFromIoTCloudConnectorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DissociateVSwitchFromIoTCloudConnectorResponseBody) GoString() string {
	return s.String()
}

func (s *DissociateVSwitchFromIoTCloudConnectorResponseBody) SetRequestId(v string) *DissociateVSwitchFromIoTCloudConnectorResponseBody {
	s.RequestId = &v
	return s
}

type DissociateVSwitchFromIoTCloudConnectorResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DissociateVSwitchFromIoTCloudConnectorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DissociateVSwitchFromIoTCloudConnectorResponse) String() string {
	return tea.Prettify(s)
}

func (s DissociateVSwitchFromIoTCloudConnectorResponse) GoString() string {
	return s.String()
}

func (s *DissociateVSwitchFromIoTCloudConnectorResponse) SetHeaders(v map[string]*string) *DissociateVSwitchFromIoTCloudConnectorResponse {
	s.Headers = v
	return s
}

func (s *DissociateVSwitchFromIoTCloudConnectorResponse) SetStatusCode(v int32) *DissociateVSwitchFromIoTCloudConnectorResponse {
	s.StatusCode = &v
	return s
}

func (s *DissociateVSwitchFromIoTCloudConnectorResponse) SetBody(v *DissociateVSwitchFromIoTCloudConnectorResponseBody) *DissociateVSwitchFromIoTCloudConnectorResponse {
	s.Body = v
	return s
}

type EnableIoTCloudConnectorAccessLogRequest struct {
	AccessLogSlsLogStore *string `json:"AccessLogSlsLogStore,omitempty" xml:"AccessLogSlsLogStore,omitempty"`
	AccessLogSlsProject  *string `json:"AccessLogSlsProject,omitempty" xml:"AccessLogSlsProject,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun               *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorId  *string `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s EnableIoTCloudConnectorAccessLogRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableIoTCloudConnectorAccessLogRequest) GoString() string {
	return s.String()
}

func (s *EnableIoTCloudConnectorAccessLogRequest) SetAccessLogSlsLogStore(v string) *EnableIoTCloudConnectorAccessLogRequest {
	s.AccessLogSlsLogStore = &v
	return s
}

func (s *EnableIoTCloudConnectorAccessLogRequest) SetAccessLogSlsProject(v string) *EnableIoTCloudConnectorAccessLogRequest {
	s.AccessLogSlsProject = &v
	return s
}

func (s *EnableIoTCloudConnectorAccessLogRequest) SetClientToken(v string) *EnableIoTCloudConnectorAccessLogRequest {
	s.ClientToken = &v
	return s
}

func (s *EnableIoTCloudConnectorAccessLogRequest) SetDryRun(v bool) *EnableIoTCloudConnectorAccessLogRequest {
	s.DryRun = &v
	return s
}

func (s *EnableIoTCloudConnectorAccessLogRequest) SetIoTCloudConnectorId(v string) *EnableIoTCloudConnectorAccessLogRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *EnableIoTCloudConnectorAccessLogRequest) SetRegionId(v string) *EnableIoTCloudConnectorAccessLogRequest {
	s.RegionId = &v
	return s
}

type EnableIoTCloudConnectorAccessLogResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableIoTCloudConnectorAccessLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableIoTCloudConnectorAccessLogResponseBody) GoString() string {
	return s.String()
}

func (s *EnableIoTCloudConnectorAccessLogResponseBody) SetRequestId(v string) *EnableIoTCloudConnectorAccessLogResponseBody {
	s.RequestId = &v
	return s
}

type EnableIoTCloudConnectorAccessLogResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableIoTCloudConnectorAccessLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableIoTCloudConnectorAccessLogResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableIoTCloudConnectorAccessLogResponse) GoString() string {
	return s.String()
}

func (s *EnableIoTCloudConnectorAccessLogResponse) SetHeaders(v map[string]*string) *EnableIoTCloudConnectorAccessLogResponse {
	s.Headers = v
	return s
}

func (s *EnableIoTCloudConnectorAccessLogResponse) SetStatusCode(v int32) *EnableIoTCloudConnectorAccessLogResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableIoTCloudConnectorAccessLogResponse) SetBody(v *EnableIoTCloudConnectorAccessLogResponseBody) *EnableIoTCloudConnectorAccessLogResponse {
	s.Body = v
	return s
}

type GetConnectionPoolIpOperationResultRequest struct {
	ConnectionPoolId    *string `json:"ConnectionPoolId,omitempty" xml:"ConnectionPoolId,omitempty"`
	IoTCloudConnectorId *string `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	QueryRequestId      *string `json:"QueryRequestId,omitempty" xml:"QueryRequestId,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetConnectionPoolIpOperationResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionPoolIpOperationResultRequest) GoString() string {
	return s.String()
}

func (s *GetConnectionPoolIpOperationResultRequest) SetConnectionPoolId(v string) *GetConnectionPoolIpOperationResultRequest {
	s.ConnectionPoolId = &v
	return s
}

func (s *GetConnectionPoolIpOperationResultRequest) SetIoTCloudConnectorId(v string) *GetConnectionPoolIpOperationResultRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *GetConnectionPoolIpOperationResultRequest) SetQueryRequestId(v string) *GetConnectionPoolIpOperationResultRequest {
	s.QueryRequestId = &v
	return s
}

func (s *GetConnectionPoolIpOperationResultRequest) SetRegionId(v string) *GetConnectionPoolIpOperationResultRequest {
	s.RegionId = &v
	return s
}

type GetConnectionPoolIpOperationResultResponseBody struct {
	RequestId       *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultFilePaths []*string `json:"ResultFilePaths,omitempty" xml:"ResultFilePaths,omitempty" type:"Repeated"`
}

func (s GetConnectionPoolIpOperationResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionPoolIpOperationResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetConnectionPoolIpOperationResultResponseBody) SetRequestId(v string) *GetConnectionPoolIpOperationResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConnectionPoolIpOperationResultResponseBody) SetResultFilePaths(v []*string) *GetConnectionPoolIpOperationResultResponseBody {
	s.ResultFilePaths = v
	return s
}

type GetConnectionPoolIpOperationResultResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetConnectionPoolIpOperationResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetConnectionPoolIpOperationResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionPoolIpOperationResultResponse) GoString() string {
	return s.String()
}

func (s *GetConnectionPoolIpOperationResultResponse) SetHeaders(v map[string]*string) *GetConnectionPoolIpOperationResultResponse {
	s.Headers = v
	return s
}

func (s *GetConnectionPoolIpOperationResultResponse) SetStatusCode(v int32) *GetConnectionPoolIpOperationResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConnectionPoolIpOperationResultResponse) SetBody(v *GetConnectionPoolIpOperationResultResponseBody) *GetConnectionPoolIpOperationResultResponse {
	s.Body = v
	return s
}

type GetDiagnoseResultForSingleCardRequest struct {
	DiagnoseTaskId *string `json:"DiagnoseTaskId,omitempty" xml:"DiagnoseTaskId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *GetDiagnoseResultForSingleCardRequest) SetRegionId(v string) *GetDiagnoseResultForSingleCardRequest {
	s.RegionId = &v
	return s
}

type GetDiagnoseResultForSingleCardResponseBody struct {
	BeginTime           *int64                                                    `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	CardIp              *string                                                   `json:"CardIp,omitempty" xml:"CardIp,omitempty"`
	Destination         *string                                                   `json:"Destination,omitempty" xml:"Destination,omitempty"`
	DiagnoseItem        []*GetDiagnoseResultForSingleCardResponseBodyDiagnoseItem `json:"DiagnoseItem,omitempty" xml:"DiagnoseItem,omitempty" type:"Repeated"`
	EndTime             *int64                                                    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ErrorResult         []*GetDiagnoseResultForSingleCardResponseBodyErrorResult  `json:"ErrorResult,omitempty" xml:"ErrorResult,omitempty" type:"Repeated"`
	IccId               *string                                                   `json:"IccId,omitempty" xml:"IccId,omitempty"`
	IoTCloudConnectorId *string                                                   `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	RequestId           *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status              *string                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
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

func (s *GetDiagnoseResultForSingleCardResponseBody) SetIoTCloudConnectorId(v string) *GetDiagnoseResultForSingleCardResponseBody {
	s.IoTCloudConnectorId = &v
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

type GetDiagnoseResultForSingleCardResponseBodyDiagnoseItem struct {
	Part   *string `json:"Part,omitempty" xml:"Part,omitempty"`
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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

type GetDiagnoseResultForSingleCardResponseBodyErrorResult struct {
	ErrorDesc       *string `json:"ErrorDesc,omitempty" xml:"ErrorDesc,omitempty"`
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

type GetIoTCloudConnectorAccessLogRequest struct {
	ClientToken         *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun              *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorId *string `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetIoTCloudConnectorAccessLogRequest) String() string {
	return tea.Prettify(s)
}

func (s GetIoTCloudConnectorAccessLogRequest) GoString() string {
	return s.String()
}

func (s *GetIoTCloudConnectorAccessLogRequest) SetClientToken(v string) *GetIoTCloudConnectorAccessLogRequest {
	s.ClientToken = &v
	return s
}

func (s *GetIoTCloudConnectorAccessLogRequest) SetDryRun(v bool) *GetIoTCloudConnectorAccessLogRequest {
	s.DryRun = &v
	return s
}

func (s *GetIoTCloudConnectorAccessLogRequest) SetIoTCloudConnectorId(v string) *GetIoTCloudConnectorAccessLogRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *GetIoTCloudConnectorAccessLogRequest) SetRegionId(v string) *GetIoTCloudConnectorAccessLogRequest {
	s.RegionId = &v
	return s
}

type GetIoTCloudConnectorAccessLogResponseBody struct {
	AccessLogSlsLogStore *string `json:"AccessLogSlsLogStore,omitempty" xml:"AccessLogSlsLogStore,omitempty"`
	AccessLogSlsProject  *string `json:"AccessLogSlsProject,omitempty" xml:"AccessLogSlsProject,omitempty"`
	AccessLogStatus      *string `json:"AccessLogStatus,omitempty" xml:"AccessLogStatus,omitempty"`
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetIoTCloudConnectorAccessLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIoTCloudConnectorAccessLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetIoTCloudConnectorAccessLogResponseBody) SetAccessLogSlsLogStore(v string) *GetIoTCloudConnectorAccessLogResponseBody {
	s.AccessLogSlsLogStore = &v
	return s
}

func (s *GetIoTCloudConnectorAccessLogResponseBody) SetAccessLogSlsProject(v string) *GetIoTCloudConnectorAccessLogResponseBody {
	s.AccessLogSlsProject = &v
	return s
}

func (s *GetIoTCloudConnectorAccessLogResponseBody) SetAccessLogStatus(v string) *GetIoTCloudConnectorAccessLogResponseBody {
	s.AccessLogStatus = &v
	return s
}

func (s *GetIoTCloudConnectorAccessLogResponseBody) SetRequestId(v string) *GetIoTCloudConnectorAccessLogResponseBody {
	s.RequestId = &v
	return s
}

type GetIoTCloudConnectorAccessLogResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetIoTCloudConnectorAccessLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetIoTCloudConnectorAccessLogResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIoTCloudConnectorAccessLogResponse) GoString() string {
	return s.String()
}

func (s *GetIoTCloudConnectorAccessLogResponse) SetHeaders(v map[string]*string) *GetIoTCloudConnectorAccessLogResponse {
	s.Headers = v
	return s
}

func (s *GetIoTCloudConnectorAccessLogResponse) SetStatusCode(v int32) *GetIoTCloudConnectorAccessLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIoTCloudConnectorAccessLogResponse) SetBody(v *GetIoTCloudConnectorAccessLogResponseBody) *GetIoTCloudConnectorAccessLogResponse {
	s.Body = v
	return s
}

type GetStsInfoAndOssPathRequest struct {
	ClientToken         *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConnectionPoolId    *string `json:"ConnectionPoolId,omitempty" xml:"ConnectionPoolId,omitempty"`
	DryRun              *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	FileName            *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	IoTCloudConnectorId *string `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetStsInfoAndOssPathRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStsInfoAndOssPathRequest) GoString() string {
	return s.String()
}

func (s *GetStsInfoAndOssPathRequest) SetClientToken(v string) *GetStsInfoAndOssPathRequest {
	s.ClientToken = &v
	return s
}

func (s *GetStsInfoAndOssPathRequest) SetConnectionPoolId(v string) *GetStsInfoAndOssPathRequest {
	s.ConnectionPoolId = &v
	return s
}

func (s *GetStsInfoAndOssPathRequest) SetDryRun(v bool) *GetStsInfoAndOssPathRequest {
	s.DryRun = &v
	return s
}

func (s *GetStsInfoAndOssPathRequest) SetFileName(v string) *GetStsInfoAndOssPathRequest {
	s.FileName = &v
	return s
}

func (s *GetStsInfoAndOssPathRequest) SetIoTCloudConnectorId(v string) *GetStsInfoAndOssPathRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *GetStsInfoAndOssPathRequest) SetRegionId(v string) *GetStsInfoAndOssPathRequest {
	s.RegionId = &v
	return s
}

type GetStsInfoAndOssPathResponseBody struct {
	AccessKeyId     *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	Expiration      *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	OssPath         *string `json:"OssPath,omitempty" xml:"OssPath,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetStsInfoAndOssPathResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStsInfoAndOssPathResponseBody) GoString() string {
	return s.String()
}

func (s *GetStsInfoAndOssPathResponseBody) SetAccessKeyId(v string) *GetStsInfoAndOssPathResponseBody {
	s.AccessKeyId = &v
	return s
}

func (s *GetStsInfoAndOssPathResponseBody) SetAccessKeySecret(v string) *GetStsInfoAndOssPathResponseBody {
	s.AccessKeySecret = &v
	return s
}

func (s *GetStsInfoAndOssPathResponseBody) SetExpiration(v string) *GetStsInfoAndOssPathResponseBody {
	s.Expiration = &v
	return s
}

func (s *GetStsInfoAndOssPathResponseBody) SetOssPath(v string) *GetStsInfoAndOssPathResponseBody {
	s.OssPath = &v
	return s
}

func (s *GetStsInfoAndOssPathResponseBody) SetRequestId(v string) *GetStsInfoAndOssPathResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStsInfoAndOssPathResponseBody) SetSecurityToken(v string) *GetStsInfoAndOssPathResponseBody {
	s.SecurityToken = &v
	return s
}

type GetStsInfoAndOssPathResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetStsInfoAndOssPathResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetStsInfoAndOssPathResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStsInfoAndOssPathResponse) GoString() string {
	return s.String()
}

func (s *GetStsInfoAndOssPathResponse) SetHeaders(v map[string]*string) *GetStsInfoAndOssPathResponse {
	s.Headers = v
	return s
}

func (s *GetStsInfoAndOssPathResponse) SetStatusCode(v int32) *GetStsInfoAndOssPathResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStsInfoAndOssPathResponse) SetBody(v *GetStsInfoAndOssPathResponseBody) *GetStsInfoAndOssPathResponse {
	s.Body = v
	return s
}

type GrantVirtualBorderRouterRequest struct {
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VirtualBorderRouterId *string `json:"VirtualBorderRouterId,omitempty" xml:"VirtualBorderRouterId,omitempty"`
}

func (s GrantVirtualBorderRouterRequest) String() string {
	return tea.Prettify(s)
}

func (s GrantVirtualBorderRouterRequest) GoString() string {
	return s.String()
}

func (s *GrantVirtualBorderRouterRequest) SetRegionId(v string) *GrantVirtualBorderRouterRequest {
	s.RegionId = &v
	return s
}

func (s *GrantVirtualBorderRouterRequest) SetVirtualBorderRouterId(v string) *GrantVirtualBorderRouterRequest {
	s.VirtualBorderRouterId = &v
	return s
}

type GrantVirtualBorderRouterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GrantVirtualBorderRouterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GrantVirtualBorderRouterResponseBody) GoString() string {
	return s.String()
}

func (s *GrantVirtualBorderRouterResponseBody) SetRequestId(v string) *GrantVirtualBorderRouterResponseBody {
	s.RequestId = &v
	return s
}

type GrantVirtualBorderRouterResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GrantVirtualBorderRouterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GrantVirtualBorderRouterResponse) String() string {
	return tea.Prettify(s)
}

func (s GrantVirtualBorderRouterResponse) GoString() string {
	return s.String()
}

func (s *GrantVirtualBorderRouterResponse) SetHeaders(v map[string]*string) *GrantVirtualBorderRouterResponse {
	s.Headers = v
	return s
}

func (s *GrantVirtualBorderRouterResponse) SetStatusCode(v int32) *GrantVirtualBorderRouterResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantVirtualBorderRouterResponse) SetBody(v *GrantVirtualBorderRouterResponseBody) *GrantVirtualBorderRouterResponse {
	s.Body = v
	return s
}

type ListAPNsRequest struct {
	APN        *string `json:"APN,omitempty" xml:"APN,omitempty"`
	ISP        *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
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

func (s *ListAPNsRequest) SetMaxResults(v int32) *ListAPNsRequest {
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

func (s *ListAPNsRequest) SetType(v string) *ListAPNsRequest {
	s.Type = &v
	return s
}

type ListAPNsResponseBody struct {
	APNs       []*ListAPNsResponseBodyAPNs `json:"APNs,omitempty" xml:"APNs,omitempty" type:"Repeated"`
	MaxResults *int32                      `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *ListAPNsResponseBody) SetMaxResults(v int32) *ListAPNsResponseBody {
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

func (s *ListAPNsResponseBody) SetTotalCount(v int32) *ListAPNsResponseBody {
	s.TotalCount = &v
	return s
}

type ListAPNsResponseBodyAPNs struct {
	APN         *string   `json:"APN,omitempty" xml:"APN,omitempty"`
	Description *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	FeatureList []*string `json:"FeatureList,omitempty" xml:"FeatureList,omitempty" type:"Repeated"`
	ISP         *string   `json:"ISP,omitempty" xml:"ISP,omitempty"`
	Name        *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	ZoneList    []*string `json:"ZoneList,omitempty" xml:"ZoneList,omitempty" type:"Repeated"`
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

func (s *ListAPNsResponseBodyAPNs) SetFeatureList(v []*string) *ListAPNsResponseBodyAPNs {
	s.FeatureList = v
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

func (s *ListAPNsResponseBodyAPNs) SetZoneList(v []*string) *ListAPNsResponseBodyAPNs {
	s.ZoneList = v
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
	AuthorizationRuleIds       []*string `json:"AuthorizationRuleIds,omitempty" xml:"AuthorizationRuleIds,omitempty" type:"Repeated"`
	AuthorizationRuleName      []*string `json:"AuthorizationRuleName,omitempty" xml:"AuthorizationRuleName,omitempty" type:"Repeated"`
	AuthorizationRuleStatus    []*string `json:"AuthorizationRuleStatus,omitempty" xml:"AuthorizationRuleStatus,omitempty" type:"Repeated"`
	AuthorizationRuleType      *string   `json:"AuthorizationRuleType,omitempty" xml:"AuthorizationRuleType,omitempty"`
	Destination                []*string `json:"Destination,omitempty" xml:"Destination,omitempty" type:"Repeated"`
	DestinationPort            []*string `json:"DestinationPort,omitempty" xml:"DestinationPort,omitempty" type:"Repeated"`
	DestinationType            []*string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty" type:"Repeated"`
	FuzzyAuthorizationRuleName *string   `json:"FuzzyAuthorizationRuleName,omitempty" xml:"FuzzyAuthorizationRuleName,omitempty"`
	FuzzyDestination           *string   `json:"FuzzyDestination,omitempty" xml:"FuzzyDestination,omitempty"`
	IoTCloudConnectorId        *string   `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	MaxResults                 *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken                  *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Policy                     []*string `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Repeated"`
	Protocol                   []*string `json:"Protocol,omitempty" xml:"Protocol,omitempty" type:"Repeated"`
	RegionId                   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *ListAuthorizationRulesRequest) SetAuthorizationRuleName(v []*string) *ListAuthorizationRulesRequest {
	s.AuthorizationRuleName = v
	return s
}

func (s *ListAuthorizationRulesRequest) SetAuthorizationRuleStatus(v []*string) *ListAuthorizationRulesRequest {
	s.AuthorizationRuleStatus = v
	return s
}

func (s *ListAuthorizationRulesRequest) SetAuthorizationRuleType(v string) *ListAuthorizationRulesRequest {
	s.AuthorizationRuleType = &v
	return s
}

func (s *ListAuthorizationRulesRequest) SetDestination(v []*string) *ListAuthorizationRulesRequest {
	s.Destination = v
	return s
}

func (s *ListAuthorizationRulesRequest) SetDestinationPort(v []*string) *ListAuthorizationRulesRequest {
	s.DestinationPort = v
	return s
}

func (s *ListAuthorizationRulesRequest) SetDestinationType(v []*string) *ListAuthorizationRulesRequest {
	s.DestinationType = v
	return s
}

func (s *ListAuthorizationRulesRequest) SetFuzzyAuthorizationRuleName(v string) *ListAuthorizationRulesRequest {
	s.FuzzyAuthorizationRuleName = &v
	return s
}

func (s *ListAuthorizationRulesRequest) SetFuzzyDestination(v string) *ListAuthorizationRulesRequest {
	s.FuzzyDestination = &v
	return s
}

func (s *ListAuthorizationRulesRequest) SetIoTCloudConnectorId(v string) *ListAuthorizationRulesRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *ListAuthorizationRulesRequest) SetMaxResults(v int32) *ListAuthorizationRulesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAuthorizationRulesRequest) SetNextToken(v string) *ListAuthorizationRulesRequest {
	s.NextToken = &v
	return s
}

func (s *ListAuthorizationRulesRequest) SetPolicy(v []*string) *ListAuthorizationRulesRequest {
	s.Policy = v
	return s
}

func (s *ListAuthorizationRulesRequest) SetProtocol(v []*string) *ListAuthorizationRulesRequest {
	s.Protocol = v
	return s
}

func (s *ListAuthorizationRulesRequest) SetRegionId(v string) *ListAuthorizationRulesRequest {
	s.RegionId = &v
	return s
}

type ListAuthorizationRulesResponseBody struct {
	AuthorizationRules []*ListAuthorizationRulesResponseBodyAuthorizationRules `json:"AuthorizationRules,omitempty" xml:"AuthorizationRules,omitempty" type:"Repeated"`
	MaxResults         *int32                                                  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken          *string                                                 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId          *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount         *int32                                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *ListAuthorizationRulesResponseBody) SetMaxResults(v int32) *ListAuthorizationRulesResponseBody {
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

func (s *ListAuthorizationRulesResponseBody) SetTotalCount(v int32) *ListAuthorizationRulesResponseBody {
	s.TotalCount = &v
	return s
}

type ListAuthorizationRulesResponseBodyAuthorizationRules struct {
	AuthorizationRuleDescription *string   `json:"AuthorizationRuleDescription,omitempty" xml:"AuthorizationRuleDescription,omitempty"`
	AuthorizationRuleId          *string   `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	AuthorizationRuleName        *string   `json:"AuthorizationRuleName,omitempty" xml:"AuthorizationRuleName,omitempty"`
	AuthorizationRuleStatus      *string   `json:"AuthorizationRuleStatus,omitempty" xml:"AuthorizationRuleStatus,omitempty"`
	AuthorizationRuleType        *string   `json:"AuthorizationRuleType,omitempty" xml:"AuthorizationRuleType,omitempty"`
	Destination                  *string   `json:"Destination,omitempty" xml:"Destination,omitempty"`
	DestinationPort              *string   `json:"DestinationPort,omitempty" xml:"DestinationPort,omitempty"`
	DestinationType              *string   `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	IoTCloudConnectorId          *string   `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	Policy                       *string   `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Protocol                     *string   `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	SourceCidrs                  []*string `json:"SourceCidrs,omitempty" xml:"SourceCidrs,omitempty" type:"Repeated"`
}

func (s ListAuthorizationRulesResponseBodyAuthorizationRules) String() string {
	return tea.Prettify(s)
}

func (s ListAuthorizationRulesResponseBodyAuthorizationRules) GoString() string {
	return s.String()
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) SetAuthorizationRuleDescription(v string) *ListAuthorizationRulesResponseBodyAuthorizationRules {
	s.AuthorizationRuleDescription = &v
	return s
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) SetAuthorizationRuleId(v string) *ListAuthorizationRulesResponseBodyAuthorizationRules {
	s.AuthorizationRuleId = &v
	return s
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) SetAuthorizationRuleName(v string) *ListAuthorizationRulesResponseBodyAuthorizationRules {
	s.AuthorizationRuleName = &v
	return s
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) SetAuthorizationRuleStatus(v string) *ListAuthorizationRulesResponseBodyAuthorizationRules {
	s.AuthorizationRuleStatus = &v
	return s
}

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) SetAuthorizationRuleType(v string) *ListAuthorizationRulesResponseBodyAuthorizationRules {
	s.AuthorizationRuleType = &v
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

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) SetIoTCloudConnectorId(v string) *ListAuthorizationRulesResponseBodyAuthorizationRules {
	s.IoTCloudConnectorId = &v
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

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) SetSourceCidrs(v []*string) *ListAuthorizationRulesResponseBodyAuthorizationRules {
	s.SourceCidrs = v
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

type ListConnectionPoolAllIpsRequest struct {
	ConnectionPoolId    *string `json:"ConnectionPoolId,omitempty" xml:"ConnectionPoolId,omitempty"`
	IoTCloudConnectorId *string `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	Ip                  *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	MaxResults          *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken           *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListConnectionPoolAllIpsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListConnectionPoolAllIpsRequest) GoString() string {
	return s.String()
}

func (s *ListConnectionPoolAllIpsRequest) SetConnectionPoolId(v string) *ListConnectionPoolAllIpsRequest {
	s.ConnectionPoolId = &v
	return s
}

func (s *ListConnectionPoolAllIpsRequest) SetIoTCloudConnectorId(v string) *ListConnectionPoolAllIpsRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *ListConnectionPoolAllIpsRequest) SetIp(v string) *ListConnectionPoolAllIpsRequest {
	s.Ip = &v
	return s
}

func (s *ListConnectionPoolAllIpsRequest) SetMaxResults(v int32) *ListConnectionPoolAllIpsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListConnectionPoolAllIpsRequest) SetNextToken(v string) *ListConnectionPoolAllIpsRequest {
	s.NextToken = &v
	return s
}

func (s *ListConnectionPoolAllIpsRequest) SetRegionId(v string) *ListConnectionPoolAllIpsRequest {
	s.RegionId = &v
	return s
}

func (s *ListConnectionPoolAllIpsRequest) SetType(v string) *ListConnectionPoolAllIpsRequest {
	s.Type = &v
	return s
}

type ListConnectionPoolAllIpsResponseBody struct {
	ConnectionPoolIps []*ListConnectionPoolAllIpsResponseBodyConnectionPoolIps `json:"ConnectionPoolIps,omitempty" xml:"ConnectionPoolIps,omitempty" type:"Repeated"`
	MaxResults        *int32                                                   `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken         *string                                                  `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId         *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalIpsCount     *int32                                                   `json:"TotalIpsCount,omitempty" xml:"TotalIpsCount,omitempty"`
}

func (s ListConnectionPoolAllIpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListConnectionPoolAllIpsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConnectionPoolAllIpsResponseBody) SetConnectionPoolIps(v []*ListConnectionPoolAllIpsResponseBodyConnectionPoolIps) *ListConnectionPoolAllIpsResponseBody {
	s.ConnectionPoolIps = v
	return s
}

func (s *ListConnectionPoolAllIpsResponseBody) SetMaxResults(v int32) *ListConnectionPoolAllIpsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListConnectionPoolAllIpsResponseBody) SetNextToken(v string) *ListConnectionPoolAllIpsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListConnectionPoolAllIpsResponseBody) SetRequestId(v string) *ListConnectionPoolAllIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConnectionPoolAllIpsResponseBody) SetTotalIpsCount(v int32) *ListConnectionPoolAllIpsResponseBody {
	s.TotalIpsCount = &v
	return s
}

type ListConnectionPoolAllIpsResponseBodyConnectionPoolIps struct {
	ConnectionPoolId *string `json:"ConnectionPoolId,omitempty" xml:"ConnectionPoolId,omitempty"`
	Ip               *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	IpNum            *int64  `json:"IpNum,omitempty" xml:"IpNum,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListConnectionPoolAllIpsResponseBodyConnectionPoolIps) String() string {
	return tea.Prettify(s)
}

func (s ListConnectionPoolAllIpsResponseBodyConnectionPoolIps) GoString() string {
	return s.String()
}

func (s *ListConnectionPoolAllIpsResponseBodyConnectionPoolIps) SetConnectionPoolId(v string) *ListConnectionPoolAllIpsResponseBodyConnectionPoolIps {
	s.ConnectionPoolId = &v
	return s
}

func (s *ListConnectionPoolAllIpsResponseBodyConnectionPoolIps) SetIp(v string) *ListConnectionPoolAllIpsResponseBodyConnectionPoolIps {
	s.Ip = &v
	return s
}

func (s *ListConnectionPoolAllIpsResponseBodyConnectionPoolIps) SetIpNum(v int64) *ListConnectionPoolAllIpsResponseBodyConnectionPoolIps {
	s.IpNum = &v
	return s
}

func (s *ListConnectionPoolAllIpsResponseBodyConnectionPoolIps) SetStatus(v string) *ListConnectionPoolAllIpsResponseBodyConnectionPoolIps {
	s.Status = &v
	return s
}

func (s *ListConnectionPoolAllIpsResponseBodyConnectionPoolIps) SetType(v string) *ListConnectionPoolAllIpsResponseBodyConnectionPoolIps {
	s.Type = &v
	return s
}

type ListConnectionPoolAllIpsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListConnectionPoolAllIpsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListConnectionPoolAllIpsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListConnectionPoolAllIpsResponse) GoString() string {
	return s.String()
}

func (s *ListConnectionPoolAllIpsResponse) SetHeaders(v map[string]*string) *ListConnectionPoolAllIpsResponse {
	s.Headers = v
	return s
}

func (s *ListConnectionPoolAllIpsResponse) SetStatusCode(v int32) *ListConnectionPoolAllIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConnectionPoolAllIpsResponse) SetBody(v *ListConnectionPoolAllIpsResponseBody) *ListConnectionPoolAllIpsResponse {
	s.Body = v
	return s
}

type ListConnectionPoolIpsRequest struct {
	ConnectionPoolId    *string `json:"ConnectionPoolId,omitempty" xml:"ConnectionPoolId,omitempty"`
	IoTCloudConnectorId *string `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	Ip                  *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	MaxResults          *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken           *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListConnectionPoolIpsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListConnectionPoolIpsRequest) GoString() string {
	return s.String()
}

func (s *ListConnectionPoolIpsRequest) SetConnectionPoolId(v string) *ListConnectionPoolIpsRequest {
	s.ConnectionPoolId = &v
	return s
}

func (s *ListConnectionPoolIpsRequest) SetIoTCloudConnectorId(v string) *ListConnectionPoolIpsRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *ListConnectionPoolIpsRequest) SetIp(v string) *ListConnectionPoolIpsRequest {
	s.Ip = &v
	return s
}

func (s *ListConnectionPoolIpsRequest) SetMaxResults(v int32) *ListConnectionPoolIpsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListConnectionPoolIpsRequest) SetNextToken(v string) *ListConnectionPoolIpsRequest {
	s.NextToken = &v
	return s
}

func (s *ListConnectionPoolIpsRequest) SetRegionId(v string) *ListConnectionPoolIpsRequest {
	s.RegionId = &v
	return s
}

type ListConnectionPoolIpsResponseBody struct {
	ConnectionPoolIps []*ListConnectionPoolIpsResponseBodyConnectionPoolIps `json:"ConnectionPoolIps,omitempty" xml:"ConnectionPoolIps,omitempty" type:"Repeated"`
	MaxResults        *int32                                                `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken         *string                                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId         *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount        *int32                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListConnectionPoolIpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListConnectionPoolIpsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConnectionPoolIpsResponseBody) SetConnectionPoolIps(v []*ListConnectionPoolIpsResponseBodyConnectionPoolIps) *ListConnectionPoolIpsResponseBody {
	s.ConnectionPoolIps = v
	return s
}

func (s *ListConnectionPoolIpsResponseBody) SetMaxResults(v int32) *ListConnectionPoolIpsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListConnectionPoolIpsResponseBody) SetNextToken(v string) *ListConnectionPoolIpsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListConnectionPoolIpsResponseBody) SetRequestId(v string) *ListConnectionPoolIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConnectionPoolIpsResponseBody) SetTotalCount(v int32) *ListConnectionPoolIpsResponseBody {
	s.TotalCount = &v
	return s
}

type ListConnectionPoolIpsResponseBodyConnectionPoolIps struct {
	ConnectionPoolId *string `json:"ConnectionPoolId,omitempty" xml:"ConnectionPoolId,omitempty"`
	Ip               *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListConnectionPoolIpsResponseBodyConnectionPoolIps) String() string {
	return tea.Prettify(s)
}

func (s ListConnectionPoolIpsResponseBodyConnectionPoolIps) GoString() string {
	return s.String()
}

func (s *ListConnectionPoolIpsResponseBodyConnectionPoolIps) SetConnectionPoolId(v string) *ListConnectionPoolIpsResponseBodyConnectionPoolIps {
	s.ConnectionPoolId = &v
	return s
}

func (s *ListConnectionPoolIpsResponseBodyConnectionPoolIps) SetIp(v string) *ListConnectionPoolIpsResponseBodyConnectionPoolIps {
	s.Ip = &v
	return s
}

func (s *ListConnectionPoolIpsResponseBodyConnectionPoolIps) SetStatus(v string) *ListConnectionPoolIpsResponseBodyConnectionPoolIps {
	s.Status = &v
	return s
}

type ListConnectionPoolIpsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListConnectionPoolIpsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListConnectionPoolIpsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListConnectionPoolIpsResponse) GoString() string {
	return s.String()
}

func (s *ListConnectionPoolIpsResponse) SetHeaders(v map[string]*string) *ListConnectionPoolIpsResponse {
	s.Headers = v
	return s
}

func (s *ListConnectionPoolIpsResponse) SetStatusCode(v int32) *ListConnectionPoolIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConnectionPoolIpsResponse) SetBody(v *ListConnectionPoolIpsResponseBody) *ListConnectionPoolIpsResponse {
	s.Body = v
	return s
}

type ListConnectionPoolsRequest struct {
	ConnectionPoolIds    []*string `json:"ConnectionPoolIds,omitempty" xml:"ConnectionPoolIds,omitempty" type:"Repeated"`
	ConnectionPoolName   []*string `json:"ConnectionPoolName,omitempty" xml:"ConnectionPoolName,omitempty" type:"Repeated"`
	ConnectionPoolStatus []*string `json:"ConnectionPoolStatus,omitempty" xml:"ConnectionPoolStatus,omitempty" type:"Repeated"`
	IoTCloudConnectorId  *string   `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	MaxResults           *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken            *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListConnectionPoolsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListConnectionPoolsRequest) GoString() string {
	return s.String()
}

func (s *ListConnectionPoolsRequest) SetConnectionPoolIds(v []*string) *ListConnectionPoolsRequest {
	s.ConnectionPoolIds = v
	return s
}

func (s *ListConnectionPoolsRequest) SetConnectionPoolName(v []*string) *ListConnectionPoolsRequest {
	s.ConnectionPoolName = v
	return s
}

func (s *ListConnectionPoolsRequest) SetConnectionPoolStatus(v []*string) *ListConnectionPoolsRequest {
	s.ConnectionPoolStatus = v
	return s
}

func (s *ListConnectionPoolsRequest) SetIoTCloudConnectorId(v string) *ListConnectionPoolsRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *ListConnectionPoolsRequest) SetMaxResults(v int32) *ListConnectionPoolsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListConnectionPoolsRequest) SetNextToken(v string) *ListConnectionPoolsRequest {
	s.NextToken = &v
	return s
}

func (s *ListConnectionPoolsRequest) SetRegionId(v string) *ListConnectionPoolsRequest {
	s.RegionId = &v
	return s
}

type ListConnectionPoolsResponseBody struct {
	ConnectionPools []*ListConnectionPoolsResponseBodyConnectionPools `json:"ConnectionPools,omitempty" xml:"ConnectionPools,omitempty" type:"Repeated"`
	MaxResults      *int32                                            `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken       *string                                           `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId       *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount      *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListConnectionPoolsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListConnectionPoolsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConnectionPoolsResponseBody) SetConnectionPools(v []*ListConnectionPoolsResponseBodyConnectionPools) *ListConnectionPoolsResponseBody {
	s.ConnectionPools = v
	return s
}

func (s *ListConnectionPoolsResponseBody) SetMaxResults(v int32) *ListConnectionPoolsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListConnectionPoolsResponseBody) SetNextToken(v string) *ListConnectionPoolsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListConnectionPoolsResponseBody) SetRequestId(v string) *ListConnectionPoolsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConnectionPoolsResponseBody) SetTotalCount(v int32) *ListConnectionPoolsResponseBody {
	s.TotalCount = &v
	return s
}

type ListConnectionPoolsResponseBodyConnectionPools struct {
	Cidrs                     []*string `json:"Cidrs,omitempty" xml:"Cidrs,omitempty" type:"Repeated"`
	ConnectionPoolDescription *string   `json:"ConnectionPoolDescription,omitempty" xml:"ConnectionPoolDescription,omitempty"`
	ConnectionPoolId          *string   `json:"ConnectionPoolId,omitempty" xml:"ConnectionPoolId,omitempty"`
	ConnectionPoolName        *string   `json:"ConnectionPoolName,omitempty" xml:"ConnectionPoolName,omitempty"`
	ConnectionPoolStatus      *string   `json:"ConnectionPoolStatus,omitempty" xml:"ConnectionPoolStatus,omitempty"`
	OperateResultRequestID    *string   `json:"OperateResultRequestID,omitempty" xml:"OperateResultRequestID,omitempty"`
}

func (s ListConnectionPoolsResponseBodyConnectionPools) String() string {
	return tea.Prettify(s)
}

func (s ListConnectionPoolsResponseBodyConnectionPools) GoString() string {
	return s.String()
}

func (s *ListConnectionPoolsResponseBodyConnectionPools) SetCidrs(v []*string) *ListConnectionPoolsResponseBodyConnectionPools {
	s.Cidrs = v
	return s
}

func (s *ListConnectionPoolsResponseBodyConnectionPools) SetConnectionPoolDescription(v string) *ListConnectionPoolsResponseBodyConnectionPools {
	s.ConnectionPoolDescription = &v
	return s
}

func (s *ListConnectionPoolsResponseBodyConnectionPools) SetConnectionPoolId(v string) *ListConnectionPoolsResponseBodyConnectionPools {
	s.ConnectionPoolId = &v
	return s
}

func (s *ListConnectionPoolsResponseBodyConnectionPools) SetConnectionPoolName(v string) *ListConnectionPoolsResponseBodyConnectionPools {
	s.ConnectionPoolName = &v
	return s
}

func (s *ListConnectionPoolsResponseBodyConnectionPools) SetConnectionPoolStatus(v string) *ListConnectionPoolsResponseBodyConnectionPools {
	s.ConnectionPoolStatus = &v
	return s
}

func (s *ListConnectionPoolsResponseBodyConnectionPools) SetOperateResultRequestID(v string) *ListConnectionPoolsResponseBodyConnectionPools {
	s.OperateResultRequestID = &v
	return s
}

type ListConnectionPoolsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListConnectionPoolsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListConnectionPoolsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListConnectionPoolsResponse) GoString() string {
	return s.String()
}

func (s *ListConnectionPoolsResponse) SetHeaders(v map[string]*string) *ListConnectionPoolsResponse {
	s.Headers = v
	return s
}

func (s *ListConnectionPoolsResponse) SetStatusCode(v int32) *ListConnectionPoolsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConnectionPoolsResponse) SetBody(v *ListConnectionPoolsResponseBody) *ListConnectionPoolsResponse {
	s.Body = v
	return s
}

type ListDNSServiceRulesRequest struct {
	DNSServiceRuleIds    []*string `json:"DNSServiceRuleIds,omitempty" xml:"DNSServiceRuleIds,omitempty" type:"Repeated"`
	DNSServiceRuleName   []*string `json:"DNSServiceRuleName,omitempty" xml:"DNSServiceRuleName,omitempty" type:"Repeated"`
	DNSServiceRuleStatus []*string `json:"DNSServiceRuleStatus,omitempty" xml:"DNSServiceRuleStatus,omitempty" type:"Repeated"`
	Destination          []*string `json:"Destination,omitempty" xml:"Destination,omitempty" type:"Repeated"`
	IoTCloudConnectorId  *string   `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	MaxResults           *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken            *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceType          *string   `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	Source               []*string `json:"Source,omitempty" xml:"Source,omitempty" type:"Repeated"`
}

func (s ListDNSServiceRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDNSServiceRulesRequest) GoString() string {
	return s.String()
}

func (s *ListDNSServiceRulesRequest) SetDNSServiceRuleIds(v []*string) *ListDNSServiceRulesRequest {
	s.DNSServiceRuleIds = v
	return s
}

func (s *ListDNSServiceRulesRequest) SetDNSServiceRuleName(v []*string) *ListDNSServiceRulesRequest {
	s.DNSServiceRuleName = v
	return s
}

func (s *ListDNSServiceRulesRequest) SetDNSServiceRuleStatus(v []*string) *ListDNSServiceRulesRequest {
	s.DNSServiceRuleStatus = v
	return s
}

func (s *ListDNSServiceRulesRequest) SetDestination(v []*string) *ListDNSServiceRulesRequest {
	s.Destination = v
	return s
}

func (s *ListDNSServiceRulesRequest) SetIoTCloudConnectorId(v string) *ListDNSServiceRulesRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *ListDNSServiceRulesRequest) SetMaxResults(v int32) *ListDNSServiceRulesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDNSServiceRulesRequest) SetNextToken(v string) *ListDNSServiceRulesRequest {
	s.NextToken = &v
	return s
}

func (s *ListDNSServiceRulesRequest) SetRegionId(v string) *ListDNSServiceRulesRequest {
	s.RegionId = &v
	return s
}

func (s *ListDNSServiceRulesRequest) SetServiceType(v string) *ListDNSServiceRulesRequest {
	s.ServiceType = &v
	return s
}

func (s *ListDNSServiceRulesRequest) SetSource(v []*string) *ListDNSServiceRulesRequest {
	s.Source = v
	return s
}

type ListDNSServiceRulesResponseBody struct {
	DNSServiceRules []*ListDNSServiceRulesResponseBodyDNSServiceRules `json:"DNSServiceRules,omitempty" xml:"DNSServiceRules,omitempty" type:"Repeated"`
	MaxResults      *int32                                            `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken       *string                                           `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId       *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount      *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDNSServiceRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDNSServiceRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDNSServiceRulesResponseBody) SetDNSServiceRules(v []*ListDNSServiceRulesResponseBodyDNSServiceRules) *ListDNSServiceRulesResponseBody {
	s.DNSServiceRules = v
	return s
}

func (s *ListDNSServiceRulesResponseBody) SetMaxResults(v int32) *ListDNSServiceRulesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDNSServiceRulesResponseBody) SetNextToken(v string) *ListDNSServiceRulesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDNSServiceRulesResponseBody) SetRequestId(v string) *ListDNSServiceRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDNSServiceRulesResponseBody) SetTotalCount(v int32) *ListDNSServiceRulesResponseBody {
	s.TotalCount = &v
	return s
}

type ListDNSServiceRulesResponseBodyDNSServiceRules struct {
	DNSServiceRuleDescription *string `json:"DNSServiceRuleDescription,omitempty" xml:"DNSServiceRuleDescription,omitempty"`
	DNSServiceRuleId          *string `json:"DNSServiceRuleId,omitempty" xml:"DNSServiceRuleId,omitempty"`
	DNSServiceRuleName        *string `json:"DNSServiceRuleName,omitempty" xml:"DNSServiceRuleName,omitempty"`
	DNSServiceRuleStatus      *string `json:"DNSServiceRuleStatus,omitempty" xml:"DNSServiceRuleStatus,omitempty"`
	Destination               *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	IoTCloudConnectorId       *string `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	ServiceType               *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	Source                    *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s ListDNSServiceRulesResponseBodyDNSServiceRules) String() string {
	return tea.Prettify(s)
}

func (s ListDNSServiceRulesResponseBodyDNSServiceRules) GoString() string {
	return s.String()
}

func (s *ListDNSServiceRulesResponseBodyDNSServiceRules) SetDNSServiceRuleDescription(v string) *ListDNSServiceRulesResponseBodyDNSServiceRules {
	s.DNSServiceRuleDescription = &v
	return s
}

func (s *ListDNSServiceRulesResponseBodyDNSServiceRules) SetDNSServiceRuleId(v string) *ListDNSServiceRulesResponseBodyDNSServiceRules {
	s.DNSServiceRuleId = &v
	return s
}

func (s *ListDNSServiceRulesResponseBodyDNSServiceRules) SetDNSServiceRuleName(v string) *ListDNSServiceRulesResponseBodyDNSServiceRules {
	s.DNSServiceRuleName = &v
	return s
}

func (s *ListDNSServiceRulesResponseBodyDNSServiceRules) SetDNSServiceRuleStatus(v string) *ListDNSServiceRulesResponseBodyDNSServiceRules {
	s.DNSServiceRuleStatus = &v
	return s
}

func (s *ListDNSServiceRulesResponseBodyDNSServiceRules) SetDestination(v string) *ListDNSServiceRulesResponseBodyDNSServiceRules {
	s.Destination = &v
	return s
}

func (s *ListDNSServiceRulesResponseBodyDNSServiceRules) SetIoTCloudConnectorId(v string) *ListDNSServiceRulesResponseBodyDNSServiceRules {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *ListDNSServiceRulesResponseBodyDNSServiceRules) SetServiceType(v string) *ListDNSServiceRulesResponseBodyDNSServiceRules {
	s.ServiceType = &v
	return s
}

func (s *ListDNSServiceRulesResponseBodyDNSServiceRules) SetSource(v string) *ListDNSServiceRulesResponseBodyDNSServiceRules {
	s.Source = &v
	return s
}

type ListDNSServiceRulesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDNSServiceRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDNSServiceRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDNSServiceRulesResponse) GoString() string {
	return s.String()
}

func (s *ListDNSServiceRulesResponse) SetHeaders(v map[string]*string) *ListDNSServiceRulesResponse {
	s.Headers = v
	return s
}

func (s *ListDNSServiceRulesResponse) SetStatusCode(v int32) *ListDNSServiceRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDNSServiceRulesResponse) SetBody(v *ListDNSServiceRulesResponseBody) *ListDNSServiceRulesResponse {
	s.Body = v
	return s
}

type ListDiagnoseInfoForSingleCardRequest struct {
	IoTCloudConnectorId *string `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	MaxResults          *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken           *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Source              *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceType          *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s ListDiagnoseInfoForSingleCardRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDiagnoseInfoForSingleCardRequest) GoString() string {
	return s.String()
}

func (s *ListDiagnoseInfoForSingleCardRequest) SetIoTCloudConnectorId(v string) *ListDiagnoseInfoForSingleCardRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *ListDiagnoseInfoForSingleCardRequest) SetMaxResults(v int32) *ListDiagnoseInfoForSingleCardRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDiagnoseInfoForSingleCardRequest) SetNextToken(v string) *ListDiagnoseInfoForSingleCardRequest {
	s.NextToken = &v
	return s
}

func (s *ListDiagnoseInfoForSingleCardRequest) SetRegionId(v string) *ListDiagnoseInfoForSingleCardRequest {
	s.RegionId = &v
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
	BeginTime           *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	CardIp              *string `json:"CardIp,omitempty" xml:"CardIp,omitempty"`
	Destination         *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	DestinationType     *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	DiagnoseTime        *int64  `json:"DiagnoseTime,omitempty" xml:"DiagnoseTime,omitempty"`
	EndTime             *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IccId               *string `json:"IccId,omitempty" xml:"IccId,omitempty"`
	IoTCloudConnectorId *string `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	Source              *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceType          *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId              *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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

func (s *ListDiagnoseInfoForSingleCardResponseBodyDiagnoseInfo) SetIoTCloudConnectorId(v string) *ListDiagnoseInfoForSingleCardResponseBodyDiagnoseInfo {
	s.IoTCloudConnectorId = &v
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

func (s *ListDiagnoseInfoForSingleCardResponseBodyDiagnoseInfo) SetTaskId(v string) *ListDiagnoseInfoForSingleCardResponseBodyDiagnoseInfo {
	s.TaskId = &v
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
	AuthorizationRuleIds     []*string `json:"AuthorizationRuleIds,omitempty" xml:"AuthorizationRuleIds,omitempty" type:"Repeated"`
	AuthorizationRuleName    []*string `json:"AuthorizationRuleName,omitempty" xml:"AuthorizationRuleName,omitempty" type:"Repeated"`
	AuthorizationRuleStatus  []*string `json:"AuthorizationRuleStatus,omitempty" xml:"AuthorizationRuleStatus,omitempty" type:"Repeated"`
	Destination              []*string `json:"Destination,omitempty" xml:"Destination,omitempty" type:"Repeated"`
	DestinationPort          []*string `json:"DestinationPort,omitempty" xml:"DestinationPort,omitempty" type:"Repeated"`
	DestinationType          []*string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty" type:"Repeated"`
	IoTCloudConnectorGroupId *string   `json:"IoTCloudConnectorGroupId,omitempty" xml:"IoTCloudConnectorGroupId,omitempty"`
	MaxResults               *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken                *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Policy                   []*string `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Repeated"`
	Protocol                 []*string `json:"Protocol,omitempty" xml:"Protocol,omitempty" type:"Repeated"`
	RegionId                 *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Type                     *string   `json:"Type,omitempty" xml:"Type,omitempty"`
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

func (s *ListGroupAuthorizationRulesRequest) SetAuthorizationRuleName(v []*string) *ListGroupAuthorizationRulesRequest {
	s.AuthorizationRuleName = v
	return s
}

func (s *ListGroupAuthorizationRulesRequest) SetAuthorizationRuleStatus(v []*string) *ListGroupAuthorizationRulesRequest {
	s.AuthorizationRuleStatus = v
	return s
}

func (s *ListGroupAuthorizationRulesRequest) SetDestination(v []*string) *ListGroupAuthorizationRulesRequest {
	s.Destination = v
	return s
}

func (s *ListGroupAuthorizationRulesRequest) SetDestinationPort(v []*string) *ListGroupAuthorizationRulesRequest {
	s.DestinationPort = v
	return s
}

func (s *ListGroupAuthorizationRulesRequest) SetDestinationType(v []*string) *ListGroupAuthorizationRulesRequest {
	s.DestinationType = v
	return s
}

func (s *ListGroupAuthorizationRulesRequest) SetIoTCloudConnectorGroupId(v string) *ListGroupAuthorizationRulesRequest {
	s.IoTCloudConnectorGroupId = &v
	return s
}

func (s *ListGroupAuthorizationRulesRequest) SetMaxResults(v int32) *ListGroupAuthorizationRulesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListGroupAuthorizationRulesRequest) SetNextToken(v string) *ListGroupAuthorizationRulesRequest {
	s.NextToken = &v
	return s
}

func (s *ListGroupAuthorizationRulesRequest) SetPolicy(v []*string) *ListGroupAuthorizationRulesRequest {
	s.Policy = v
	return s
}

func (s *ListGroupAuthorizationRulesRequest) SetProtocol(v []*string) *ListGroupAuthorizationRulesRequest {
	s.Protocol = v
	return s
}

func (s *ListGroupAuthorizationRulesRequest) SetRegionId(v string) *ListGroupAuthorizationRulesRequest {
	s.RegionId = &v
	return s
}

func (s *ListGroupAuthorizationRulesRequest) SetType(v string) *ListGroupAuthorizationRulesRequest {
	s.Type = &v
	return s
}

type ListGroupAuthorizationRulesResponseBody struct {
	GroupAuthorizationRules []*ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules `json:"GroupAuthorizationRules,omitempty" xml:"GroupAuthorizationRules,omitempty" type:"Repeated"`
	MaxResults              *int32                                                            `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken               *string                                                           `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId               *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount              *int32                                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *ListGroupAuthorizationRulesResponseBody) SetMaxResults(v int32) *ListGroupAuthorizationRulesResponseBody {
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

func (s *ListGroupAuthorizationRulesResponseBody) SetTotalCount(v int32) *ListGroupAuthorizationRulesResponseBody {
	s.TotalCount = &v
	return s
}

type ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules struct {
	AuthorizationRuleDescription *string   `json:"AuthorizationRuleDescription,omitempty" xml:"AuthorizationRuleDescription,omitempty"`
	AuthorizationRuleId          *string   `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	AuthorizationRuleName        *string   `json:"AuthorizationRuleName,omitempty" xml:"AuthorizationRuleName,omitempty"`
	AuthorizationRuleStatus      *string   `json:"AuthorizationRuleStatus,omitempty" xml:"AuthorizationRuleStatus,omitempty"`
	Destination                  *string   `json:"Destination,omitempty" xml:"Destination,omitempty"`
	DestinationPort              *string   `json:"DestinationPort,omitempty" xml:"DestinationPort,omitempty"`
	DestinationType              *string   `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	IoTCloudConnectorGroupId     *string   `json:"IoTCloudConnectorGroupId,omitempty" xml:"IoTCloudConnectorGroupId,omitempty"`
	Policy                       *string   `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Protocol                     *string   `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	SourceCidrs                  []*string `json:"SourceCidrs,omitempty" xml:"SourceCidrs,omitempty" type:"Repeated"`
	Type                         *string   `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules) String() string {
	return tea.Prettify(s)
}

func (s ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules) GoString() string {
	return s.String()
}

func (s *ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules) SetAuthorizationRuleDescription(v string) *ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules {
	s.AuthorizationRuleDescription = &v
	return s
}

func (s *ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules) SetAuthorizationRuleId(v string) *ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules {
	s.AuthorizationRuleId = &v
	return s
}

func (s *ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules) SetAuthorizationRuleName(v string) *ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules {
	s.AuthorizationRuleName = &v
	return s
}

func (s *ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules) SetAuthorizationRuleStatus(v string) *ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules {
	s.AuthorizationRuleStatus = &v
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

func (s *ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules) SetIoTCloudConnectorGroupId(v string) *ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules {
	s.IoTCloudConnectorGroupId = &v
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

func (s *ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules) SetSourceCidrs(v []*string) *ListGroupAuthorizationRulesResponseBodyGroupAuthorizationRules {
	s.SourceCidrs = v
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

type ListGroupDNSServiceRulesRequest struct {
	DNSServiceRuleIds        []*string `json:"DNSServiceRuleIds,omitempty" xml:"DNSServiceRuleIds,omitempty" type:"Repeated"`
	DNSServiceRuleName       []*string `json:"DNSServiceRuleName,omitempty" xml:"DNSServiceRuleName,omitempty" type:"Repeated"`
	DNSServiceRuleStatus     []*string `json:"DNSServiceRuleStatus,omitempty" xml:"DNSServiceRuleStatus,omitempty" type:"Repeated"`
	Destination              []*string `json:"Destination,omitempty" xml:"Destination,omitempty" type:"Repeated"`
	IoTCloudConnectorGroupId *string   `json:"IoTCloudConnectorGroupId,omitempty" xml:"IoTCloudConnectorGroupId,omitempty"`
	MaxResults               *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken                *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId                 *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceType              *string   `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	Source                   []*string `json:"Source,omitempty" xml:"Source,omitempty" type:"Repeated"`
}

func (s ListGroupDNSServiceRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGroupDNSServiceRulesRequest) GoString() string {
	return s.String()
}

func (s *ListGroupDNSServiceRulesRequest) SetDNSServiceRuleIds(v []*string) *ListGroupDNSServiceRulesRequest {
	s.DNSServiceRuleIds = v
	return s
}

func (s *ListGroupDNSServiceRulesRequest) SetDNSServiceRuleName(v []*string) *ListGroupDNSServiceRulesRequest {
	s.DNSServiceRuleName = v
	return s
}

func (s *ListGroupDNSServiceRulesRequest) SetDNSServiceRuleStatus(v []*string) *ListGroupDNSServiceRulesRequest {
	s.DNSServiceRuleStatus = v
	return s
}

func (s *ListGroupDNSServiceRulesRequest) SetDestination(v []*string) *ListGroupDNSServiceRulesRequest {
	s.Destination = v
	return s
}

func (s *ListGroupDNSServiceRulesRequest) SetIoTCloudConnectorGroupId(v string) *ListGroupDNSServiceRulesRequest {
	s.IoTCloudConnectorGroupId = &v
	return s
}

func (s *ListGroupDNSServiceRulesRequest) SetMaxResults(v int32) *ListGroupDNSServiceRulesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListGroupDNSServiceRulesRequest) SetNextToken(v string) *ListGroupDNSServiceRulesRequest {
	s.NextToken = &v
	return s
}

func (s *ListGroupDNSServiceRulesRequest) SetRegionId(v string) *ListGroupDNSServiceRulesRequest {
	s.RegionId = &v
	return s
}

func (s *ListGroupDNSServiceRulesRequest) SetServiceType(v string) *ListGroupDNSServiceRulesRequest {
	s.ServiceType = &v
	return s
}

func (s *ListGroupDNSServiceRulesRequest) SetSource(v []*string) *ListGroupDNSServiceRulesRequest {
	s.Source = v
	return s
}

type ListGroupDNSServiceRulesResponseBody struct {
	DNSServiceRules []*ListGroupDNSServiceRulesResponseBodyDNSServiceRules `json:"DNSServiceRules,omitempty" xml:"DNSServiceRules,omitempty" type:"Repeated"`
	MaxResults      *int32                                                 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken       *string                                                `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId       *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount      *int32                                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListGroupDNSServiceRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGroupDNSServiceRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupDNSServiceRulesResponseBody) SetDNSServiceRules(v []*ListGroupDNSServiceRulesResponseBodyDNSServiceRules) *ListGroupDNSServiceRulesResponseBody {
	s.DNSServiceRules = v
	return s
}

func (s *ListGroupDNSServiceRulesResponseBody) SetMaxResults(v int32) *ListGroupDNSServiceRulesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListGroupDNSServiceRulesResponseBody) SetNextToken(v string) *ListGroupDNSServiceRulesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListGroupDNSServiceRulesResponseBody) SetRequestId(v string) *ListGroupDNSServiceRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGroupDNSServiceRulesResponseBody) SetTotalCount(v int32) *ListGroupDNSServiceRulesResponseBody {
	s.TotalCount = &v
	return s
}

type ListGroupDNSServiceRulesResponseBodyDNSServiceRules struct {
	DNSServiceRuleDescription *string `json:"DNSServiceRuleDescription,omitempty" xml:"DNSServiceRuleDescription,omitempty"`
	DNSServiceRuleId          *string `json:"DNSServiceRuleId,omitempty" xml:"DNSServiceRuleId,omitempty"`
	DNSServiceRuleName        *string `json:"DNSServiceRuleName,omitempty" xml:"DNSServiceRuleName,omitempty"`
	DNSServiceRuleStatus      *string `json:"DNSServiceRuleStatus,omitempty" xml:"DNSServiceRuleStatus,omitempty"`
	Destination               *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	IoTCloudConnectorGroupId  *string `json:"IoTCloudConnectorGroupId,omitempty" xml:"IoTCloudConnectorGroupId,omitempty"`
	ServiceType               *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	Source                    *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s ListGroupDNSServiceRulesResponseBodyDNSServiceRules) String() string {
	return tea.Prettify(s)
}

func (s ListGroupDNSServiceRulesResponseBodyDNSServiceRules) GoString() string {
	return s.String()
}

func (s *ListGroupDNSServiceRulesResponseBodyDNSServiceRules) SetDNSServiceRuleDescription(v string) *ListGroupDNSServiceRulesResponseBodyDNSServiceRules {
	s.DNSServiceRuleDescription = &v
	return s
}

func (s *ListGroupDNSServiceRulesResponseBodyDNSServiceRules) SetDNSServiceRuleId(v string) *ListGroupDNSServiceRulesResponseBodyDNSServiceRules {
	s.DNSServiceRuleId = &v
	return s
}

func (s *ListGroupDNSServiceRulesResponseBodyDNSServiceRules) SetDNSServiceRuleName(v string) *ListGroupDNSServiceRulesResponseBodyDNSServiceRules {
	s.DNSServiceRuleName = &v
	return s
}

func (s *ListGroupDNSServiceRulesResponseBodyDNSServiceRules) SetDNSServiceRuleStatus(v string) *ListGroupDNSServiceRulesResponseBodyDNSServiceRules {
	s.DNSServiceRuleStatus = &v
	return s
}

func (s *ListGroupDNSServiceRulesResponseBodyDNSServiceRules) SetDestination(v string) *ListGroupDNSServiceRulesResponseBodyDNSServiceRules {
	s.Destination = &v
	return s
}

func (s *ListGroupDNSServiceRulesResponseBodyDNSServiceRules) SetIoTCloudConnectorGroupId(v string) *ListGroupDNSServiceRulesResponseBodyDNSServiceRules {
	s.IoTCloudConnectorGroupId = &v
	return s
}

func (s *ListGroupDNSServiceRulesResponseBodyDNSServiceRules) SetServiceType(v string) *ListGroupDNSServiceRulesResponseBodyDNSServiceRules {
	s.ServiceType = &v
	return s
}

func (s *ListGroupDNSServiceRulesResponseBodyDNSServiceRules) SetSource(v string) *ListGroupDNSServiceRulesResponseBodyDNSServiceRules {
	s.Source = &v
	return s
}

type ListGroupDNSServiceRulesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListGroupDNSServiceRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListGroupDNSServiceRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGroupDNSServiceRulesResponse) GoString() string {
	return s.String()
}

func (s *ListGroupDNSServiceRulesResponse) SetHeaders(v map[string]*string) *ListGroupDNSServiceRulesResponse {
	s.Headers = v
	return s
}

func (s *ListGroupDNSServiceRulesResponse) SetStatusCode(v int32) *ListGroupDNSServiceRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGroupDNSServiceRulesResponse) SetBody(v *ListGroupDNSServiceRulesResponseBody) *ListGroupDNSServiceRulesResponse {
	s.Body = v
	return s
}

type ListGroupIpMappingRulesRequest struct {
	DestinationIps           []*string `json:"DestinationIps,omitempty" xml:"DestinationIps,omitempty" type:"Repeated"`
	IoTCloudConnectorGroupId *string   `json:"IoTCloudConnectorGroupId,omitempty" xml:"IoTCloudConnectorGroupId,omitempty"`
	IpMappingRuleIds         []*string `json:"IpMappingRuleIds,omitempty" xml:"IpMappingRuleIds,omitempty" type:"Repeated"`
	IpMappingRuleNames       []*string `json:"IpMappingRuleNames,omitempty" xml:"IpMappingRuleNames,omitempty" type:"Repeated"`
	IpMappingRuleStatuses    []*string `json:"IpMappingRuleStatuses,omitempty" xml:"IpMappingRuleStatuses,omitempty" type:"Repeated"`
	MappingIps               []*string `json:"MappingIps,omitempty" xml:"MappingIps,omitempty" type:"Repeated"`
	MaxResults               *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken                *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId                 *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListGroupIpMappingRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGroupIpMappingRulesRequest) GoString() string {
	return s.String()
}

func (s *ListGroupIpMappingRulesRequest) SetDestinationIps(v []*string) *ListGroupIpMappingRulesRequest {
	s.DestinationIps = v
	return s
}

func (s *ListGroupIpMappingRulesRequest) SetIoTCloudConnectorGroupId(v string) *ListGroupIpMappingRulesRequest {
	s.IoTCloudConnectorGroupId = &v
	return s
}

func (s *ListGroupIpMappingRulesRequest) SetIpMappingRuleIds(v []*string) *ListGroupIpMappingRulesRequest {
	s.IpMappingRuleIds = v
	return s
}

func (s *ListGroupIpMappingRulesRequest) SetIpMappingRuleNames(v []*string) *ListGroupIpMappingRulesRequest {
	s.IpMappingRuleNames = v
	return s
}

func (s *ListGroupIpMappingRulesRequest) SetIpMappingRuleStatuses(v []*string) *ListGroupIpMappingRulesRequest {
	s.IpMappingRuleStatuses = v
	return s
}

func (s *ListGroupIpMappingRulesRequest) SetMappingIps(v []*string) *ListGroupIpMappingRulesRequest {
	s.MappingIps = v
	return s
}

func (s *ListGroupIpMappingRulesRequest) SetMaxResults(v int32) *ListGroupIpMappingRulesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListGroupIpMappingRulesRequest) SetNextToken(v string) *ListGroupIpMappingRulesRequest {
	s.NextToken = &v
	return s
}

func (s *ListGroupIpMappingRulesRequest) SetRegionId(v string) *ListGroupIpMappingRulesRequest {
	s.RegionId = &v
	return s
}

type ListGroupIpMappingRulesResponseBody struct {
	IpMappingRules []*ListGroupIpMappingRulesResponseBodyIpMappingRules `json:"IpMappingRules,omitempty" xml:"IpMappingRules,omitempty" type:"Repeated"`
	MaxResults     *int32                                               `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken      *string                                              `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId      *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount     *int32                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListGroupIpMappingRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGroupIpMappingRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupIpMappingRulesResponseBody) SetIpMappingRules(v []*ListGroupIpMappingRulesResponseBodyIpMappingRules) *ListGroupIpMappingRulesResponseBody {
	s.IpMappingRules = v
	return s
}

func (s *ListGroupIpMappingRulesResponseBody) SetMaxResults(v int32) *ListGroupIpMappingRulesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListGroupIpMappingRulesResponseBody) SetNextToken(v string) *ListGroupIpMappingRulesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListGroupIpMappingRulesResponseBody) SetRequestId(v string) *ListGroupIpMappingRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGroupIpMappingRulesResponseBody) SetTotalCount(v int32) *ListGroupIpMappingRulesResponseBody {
	s.TotalCount = &v
	return s
}

type ListGroupIpMappingRulesResponseBodyIpMappingRules struct {
	DestinationIp            *string `json:"DestinationIp,omitempty" xml:"DestinationIp,omitempty"`
	IoTCloudConnectorGroupId *string `json:"IoTCloudConnectorGroupId,omitempty" xml:"IoTCloudConnectorGroupId,omitempty"`
	IpMappingRuleDescription *string `json:"IpMappingRuleDescription,omitempty" xml:"IpMappingRuleDescription,omitempty"`
	IpMappingRuleId          *string `json:"IpMappingRuleId,omitempty" xml:"IpMappingRuleId,omitempty"`
	IpMappingRuleName        *string `json:"IpMappingRuleName,omitempty" xml:"IpMappingRuleName,omitempty"`
	IpMappingRuleStatus      *string `json:"IpMappingRuleStatus,omitempty" xml:"IpMappingRuleStatus,omitempty"`
	MappingIp                *string `json:"MappingIp,omitempty" xml:"MappingIp,omitempty"`
}

func (s ListGroupIpMappingRulesResponseBodyIpMappingRules) String() string {
	return tea.Prettify(s)
}

func (s ListGroupIpMappingRulesResponseBodyIpMappingRules) GoString() string {
	return s.String()
}

func (s *ListGroupIpMappingRulesResponseBodyIpMappingRules) SetDestinationIp(v string) *ListGroupIpMappingRulesResponseBodyIpMappingRules {
	s.DestinationIp = &v
	return s
}

func (s *ListGroupIpMappingRulesResponseBodyIpMappingRules) SetIoTCloudConnectorGroupId(v string) *ListGroupIpMappingRulesResponseBodyIpMappingRules {
	s.IoTCloudConnectorGroupId = &v
	return s
}

func (s *ListGroupIpMappingRulesResponseBodyIpMappingRules) SetIpMappingRuleDescription(v string) *ListGroupIpMappingRulesResponseBodyIpMappingRules {
	s.IpMappingRuleDescription = &v
	return s
}

func (s *ListGroupIpMappingRulesResponseBodyIpMappingRules) SetIpMappingRuleId(v string) *ListGroupIpMappingRulesResponseBodyIpMappingRules {
	s.IpMappingRuleId = &v
	return s
}

func (s *ListGroupIpMappingRulesResponseBodyIpMappingRules) SetIpMappingRuleName(v string) *ListGroupIpMappingRulesResponseBodyIpMappingRules {
	s.IpMappingRuleName = &v
	return s
}

func (s *ListGroupIpMappingRulesResponseBodyIpMappingRules) SetIpMappingRuleStatus(v string) *ListGroupIpMappingRulesResponseBodyIpMappingRules {
	s.IpMappingRuleStatus = &v
	return s
}

func (s *ListGroupIpMappingRulesResponseBodyIpMappingRules) SetMappingIp(v string) *ListGroupIpMappingRulesResponseBodyIpMappingRules {
	s.MappingIp = &v
	return s
}

type ListGroupIpMappingRulesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListGroupIpMappingRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListGroupIpMappingRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGroupIpMappingRulesResponse) GoString() string {
	return s.String()
}

func (s *ListGroupIpMappingRulesResponse) SetHeaders(v map[string]*string) *ListGroupIpMappingRulesResponse {
	s.Headers = v
	return s
}

func (s *ListGroupIpMappingRulesResponse) SetStatusCode(v int32) *ListGroupIpMappingRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGroupIpMappingRulesResponse) SetBody(v *ListGroupIpMappingRulesResponseBody) *ListGroupIpMappingRulesResponse {
	s.Body = v
	return s
}

type ListIoTCloudConnectorAvailableZonesRequest struct {
	IoTCloudConnectorId *string `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListIoTCloudConnectorAvailableZonesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIoTCloudConnectorAvailableZonesRequest) GoString() string {
	return s.String()
}

func (s *ListIoTCloudConnectorAvailableZonesRequest) SetIoTCloudConnectorId(v string) *ListIoTCloudConnectorAvailableZonesRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *ListIoTCloudConnectorAvailableZonesRequest) SetRegionId(v string) *ListIoTCloudConnectorAvailableZonesRequest {
	s.RegionId = &v
	return s
}

type ListIoTCloudConnectorAvailableZonesResponseBody struct {
	AvailableZoneList   []*string `json:"AvailableZoneList,omitempty" xml:"AvailableZoneList,omitempty" type:"Repeated"`
	IoTCloudConnectorId *string   `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	RequestId           *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListIoTCloudConnectorAvailableZonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIoTCloudConnectorAvailableZonesResponseBody) GoString() string {
	return s.String()
}

func (s *ListIoTCloudConnectorAvailableZonesResponseBody) SetAvailableZoneList(v []*string) *ListIoTCloudConnectorAvailableZonesResponseBody {
	s.AvailableZoneList = v
	return s
}

func (s *ListIoTCloudConnectorAvailableZonesResponseBody) SetIoTCloudConnectorId(v string) *ListIoTCloudConnectorAvailableZonesResponseBody {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *ListIoTCloudConnectorAvailableZonesResponseBody) SetRequestId(v string) *ListIoTCloudConnectorAvailableZonesResponseBody {
	s.RequestId = &v
	return s
}

type ListIoTCloudConnectorAvailableZonesResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListIoTCloudConnectorAvailableZonesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListIoTCloudConnectorAvailableZonesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIoTCloudConnectorAvailableZonesResponse) GoString() string {
	return s.String()
}

func (s *ListIoTCloudConnectorAvailableZonesResponse) SetHeaders(v map[string]*string) *ListIoTCloudConnectorAvailableZonesResponse {
	s.Headers = v
	return s
}

func (s *ListIoTCloudConnectorAvailableZonesResponse) SetStatusCode(v int32) *ListIoTCloudConnectorAvailableZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIoTCloudConnectorAvailableZonesResponse) SetBody(v *ListIoTCloudConnectorAvailableZonesResponseBody) *ListIoTCloudConnectorAvailableZonesResponse {
	s.Body = v
	return s
}

type ListIoTCloudConnectorEIPsRequest struct {
	IoTCloudConnectorId *string `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	MaxResults          *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken           *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListIoTCloudConnectorEIPsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIoTCloudConnectorEIPsRequest) GoString() string {
	return s.String()
}

func (s *ListIoTCloudConnectorEIPsRequest) SetIoTCloudConnectorId(v string) *ListIoTCloudConnectorEIPsRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *ListIoTCloudConnectorEIPsRequest) SetMaxResults(v int32) *ListIoTCloudConnectorEIPsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListIoTCloudConnectorEIPsRequest) SetNextToken(v string) *ListIoTCloudConnectorEIPsRequest {
	s.NextToken = &v
	return s
}

func (s *ListIoTCloudConnectorEIPsRequest) SetRegionId(v string) *ListIoTCloudConnectorEIPsRequest {
	s.RegionId = &v
	return s
}

type ListIoTCloudConnectorEIPsResponseBody struct {
	EIPs       []*string `json:"EIPs,omitempty" xml:"EIPs,omitempty" type:"Repeated"`
	MaxResults *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListIoTCloudConnectorEIPsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIoTCloudConnectorEIPsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIoTCloudConnectorEIPsResponseBody) SetEIPs(v []*string) *ListIoTCloudConnectorEIPsResponseBody {
	s.EIPs = v
	return s
}

func (s *ListIoTCloudConnectorEIPsResponseBody) SetMaxResults(v int32) *ListIoTCloudConnectorEIPsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListIoTCloudConnectorEIPsResponseBody) SetNextToken(v string) *ListIoTCloudConnectorEIPsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListIoTCloudConnectorEIPsResponseBody) SetRequestId(v string) *ListIoTCloudConnectorEIPsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIoTCloudConnectorEIPsResponseBody) SetTotalCount(v int32) *ListIoTCloudConnectorEIPsResponseBody {
	s.TotalCount = &v
	return s
}

type ListIoTCloudConnectorEIPsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListIoTCloudConnectorEIPsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListIoTCloudConnectorEIPsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIoTCloudConnectorEIPsResponse) GoString() string {
	return s.String()
}

func (s *ListIoTCloudConnectorEIPsResponse) SetHeaders(v map[string]*string) *ListIoTCloudConnectorEIPsResponse {
	s.Headers = v
	return s
}

func (s *ListIoTCloudConnectorEIPsResponse) SetStatusCode(v int32) *ListIoTCloudConnectorEIPsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIoTCloudConnectorEIPsResponse) SetBody(v *ListIoTCloudConnectorEIPsResponseBody) *ListIoTCloudConnectorEIPsResponse {
	s.Body = v
	return s
}

type ListIoTCloudConnectorGroupsRequest struct {
	IoTCloudConnectorGroupIds    []*string `json:"IoTCloudConnectorGroupIds,omitempty" xml:"IoTCloudConnectorGroupIds,omitempty" type:"Repeated"`
	IoTCloudConnectorGroupName   []*string `json:"IoTCloudConnectorGroupName,omitempty" xml:"IoTCloudConnectorGroupName,omitempty" type:"Repeated"`
	IoTCloudConnectorGroupStatus []*string `json:"IoTCloudConnectorGroupStatus,omitempty" xml:"IoTCloudConnectorGroupStatus,omitempty" type:"Repeated"`
	MaxResults                   *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken                    *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId                     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Type                         *string   `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListIoTCloudConnectorGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIoTCloudConnectorGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListIoTCloudConnectorGroupsRequest) SetIoTCloudConnectorGroupIds(v []*string) *ListIoTCloudConnectorGroupsRequest {
	s.IoTCloudConnectorGroupIds = v
	return s
}

func (s *ListIoTCloudConnectorGroupsRequest) SetIoTCloudConnectorGroupName(v []*string) *ListIoTCloudConnectorGroupsRequest {
	s.IoTCloudConnectorGroupName = v
	return s
}

func (s *ListIoTCloudConnectorGroupsRequest) SetIoTCloudConnectorGroupStatus(v []*string) *ListIoTCloudConnectorGroupsRequest {
	s.IoTCloudConnectorGroupStatus = v
	return s
}

func (s *ListIoTCloudConnectorGroupsRequest) SetMaxResults(v int32) *ListIoTCloudConnectorGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListIoTCloudConnectorGroupsRequest) SetNextToken(v string) *ListIoTCloudConnectorGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListIoTCloudConnectorGroupsRequest) SetRegionId(v string) *ListIoTCloudConnectorGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *ListIoTCloudConnectorGroupsRequest) SetType(v string) *ListIoTCloudConnectorGroupsRequest {
	s.Type = &v
	return s
}

type ListIoTCloudConnectorGroupsResponseBody struct {
	IoTCloudConnectorGroups []*ListIoTCloudConnectorGroupsResponseBodyIoTCloudConnectorGroups `json:"IoTCloudConnectorGroups,omitempty" xml:"IoTCloudConnectorGroups,omitempty" type:"Repeated"`
	MaxResults              *int32                                                            `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken               *string                                                           `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId               *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount              *int32                                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListIoTCloudConnectorGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIoTCloudConnectorGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIoTCloudConnectorGroupsResponseBody) SetIoTCloudConnectorGroups(v []*ListIoTCloudConnectorGroupsResponseBodyIoTCloudConnectorGroups) *ListIoTCloudConnectorGroupsResponseBody {
	s.IoTCloudConnectorGroups = v
	return s
}

func (s *ListIoTCloudConnectorGroupsResponseBody) SetMaxResults(v int32) *ListIoTCloudConnectorGroupsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListIoTCloudConnectorGroupsResponseBody) SetNextToken(v string) *ListIoTCloudConnectorGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListIoTCloudConnectorGroupsResponseBody) SetRequestId(v string) *ListIoTCloudConnectorGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIoTCloudConnectorGroupsResponseBody) SetTotalCount(v int32) *ListIoTCloudConnectorGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type ListIoTCloudConnectorGroupsResponseBodyIoTCloudConnectorGroups struct {
	CreateTime                   *int64                                                                              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description                  *string                                                                             `json:"Description,omitempty" xml:"Description,omitempty"`
	IoTCloudConnectorGroupId     *string                                                                             `json:"IoTCloudConnectorGroupId,omitempty" xml:"IoTCloudConnectorGroupId,omitempty"`
	IoTCloudConnectorGroupStatus *string                                                                             `json:"IoTCloudConnectorGroupStatus,omitempty" xml:"IoTCloudConnectorGroupStatus,omitempty"`
	IoTCloudConnectors           []*ListIoTCloudConnectorGroupsResponseBodyIoTCloudConnectorGroupsIoTCloudConnectors `json:"IoTCloudConnectors,omitempty" xml:"IoTCloudConnectors,omitempty" type:"Repeated"`
	Name                         *string                                                                             `json:"Name,omitempty" xml:"Name,omitempty"`
	ServiceType                  *string                                                                             `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	Type                         *string                                                                             `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListIoTCloudConnectorGroupsResponseBodyIoTCloudConnectorGroups) String() string {
	return tea.Prettify(s)
}

func (s ListIoTCloudConnectorGroupsResponseBodyIoTCloudConnectorGroups) GoString() string {
	return s.String()
}

func (s *ListIoTCloudConnectorGroupsResponseBodyIoTCloudConnectorGroups) SetCreateTime(v int64) *ListIoTCloudConnectorGroupsResponseBodyIoTCloudConnectorGroups {
	s.CreateTime = &v
	return s
}

func (s *ListIoTCloudConnectorGroupsResponseBodyIoTCloudConnectorGroups) SetDescription(v string) *ListIoTCloudConnectorGroupsResponseBodyIoTCloudConnectorGroups {
	s.Description = &v
	return s
}

func (s *ListIoTCloudConnectorGroupsResponseBodyIoTCloudConnectorGroups) SetIoTCloudConnectorGroupId(v string) *ListIoTCloudConnectorGroupsResponseBodyIoTCloudConnectorGroups {
	s.IoTCloudConnectorGroupId = &v
	return s
}

func (s *ListIoTCloudConnectorGroupsResponseBodyIoTCloudConnectorGroups) SetIoTCloudConnectorGroupStatus(v string) *ListIoTCloudConnectorGroupsResponseBodyIoTCloudConnectorGroups {
	s.IoTCloudConnectorGroupStatus = &v
	return s
}

func (s *ListIoTCloudConnectorGroupsResponseBodyIoTCloudConnectorGroups) SetIoTCloudConnectors(v []*ListIoTCloudConnectorGroupsResponseBodyIoTCloudConnectorGroupsIoTCloudConnectors) *ListIoTCloudConnectorGroupsResponseBodyIoTCloudConnectorGroups {
	s.IoTCloudConnectors = v
	return s
}

func (s *ListIoTCloudConnectorGroupsResponseBodyIoTCloudConnectorGroups) SetName(v string) *ListIoTCloudConnectorGroupsResponseBodyIoTCloudConnectorGroups {
	s.Name = &v
	return s
}

func (s *ListIoTCloudConnectorGroupsResponseBodyIoTCloudConnectorGroups) SetServiceType(v string) *ListIoTCloudConnectorGroupsResponseBodyIoTCloudConnectorGroups {
	s.ServiceType = &v
	return s
}

func (s *ListIoTCloudConnectorGroupsResponseBodyIoTCloudConnectorGroups) SetType(v string) *ListIoTCloudConnectorGroupsResponseBodyIoTCloudConnectorGroups {
	s.Type = &v
	return s
}

type ListIoTCloudConnectorGroupsResponseBodyIoTCloudConnectorGroupsIoTCloudConnectors struct {
	APN                          *string `json:"APN,omitempty" xml:"APN,omitempty"`
	CreateTime                   *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ISP                          *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	IoTCloudConnectorDescription *string `json:"IoTCloudConnectorDescription,omitempty" xml:"IoTCloudConnectorDescription,omitempty"`
	IoTCloudConnectorId          *string `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	IoTCloudConnectorName        *string `json:"IoTCloudConnectorName,omitempty" xml:"IoTCloudConnectorName,omitempty"`
	IoTCloudConnectorStatus      *string `json:"IoTCloudConnectorStatus,omitempty" xml:"IoTCloudConnectorStatus,omitempty"`
	ServiceType                  *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s ListIoTCloudConnectorGroupsResponseBodyIoTCloudConnectorGroupsIoTCloudConnectors) String() string {
	return tea.Prettify(s)
}

func (s ListIoTCloudConnectorGroupsResponseBodyIoTCloudConnectorGroupsIoTCloudConnectors) GoString() string {
	return s.String()
}

func (s *ListIoTCloudConnectorGroupsResponseBodyIoTCloudConnectorGroupsIoTCloudConnectors) SetAPN(v string) *ListIoTCloudConnectorGroupsResponseBodyIoTCloudConnectorGroupsIoTCloudConnectors {
	s.APN = &v
	return s
}

func (s *ListIoTCloudConnectorGroupsResponseBodyIoTCloudConnectorGroupsIoTCloudConnectors) SetCreateTime(v int64) *ListIoTCloudConnectorGroupsResponseBodyIoTCloudConnectorGroupsIoTCloudConnectors {
	s.CreateTime = &v
	return s
}

func (s *ListIoTCloudConnectorGroupsResponseBodyIoTCloudConnectorGroupsIoTCloudConnectors) SetISP(v string) *ListIoTCloudConnectorGroupsResponseBodyIoTCloudConnectorGroupsIoTCloudConnectors {
	s.ISP = &v
	return s
}

func (s *ListIoTCloudConnectorGroupsResponseBodyIoTCloudConnectorGroupsIoTCloudConnectors) SetIoTCloudConnectorDescription(v string) *ListIoTCloudConnectorGroupsResponseBodyIoTCloudConnectorGroupsIoTCloudConnectors {
	s.IoTCloudConnectorDescription = &v
	return s
}

func (s *ListIoTCloudConnectorGroupsResponseBodyIoTCloudConnectorGroupsIoTCloudConnectors) SetIoTCloudConnectorId(v string) *ListIoTCloudConnectorGroupsResponseBodyIoTCloudConnectorGroupsIoTCloudConnectors {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *ListIoTCloudConnectorGroupsResponseBodyIoTCloudConnectorGroupsIoTCloudConnectors) SetIoTCloudConnectorName(v string) *ListIoTCloudConnectorGroupsResponseBodyIoTCloudConnectorGroupsIoTCloudConnectors {
	s.IoTCloudConnectorName = &v
	return s
}

func (s *ListIoTCloudConnectorGroupsResponseBodyIoTCloudConnectorGroupsIoTCloudConnectors) SetIoTCloudConnectorStatus(v string) *ListIoTCloudConnectorGroupsResponseBodyIoTCloudConnectorGroupsIoTCloudConnectors {
	s.IoTCloudConnectorStatus = &v
	return s
}

func (s *ListIoTCloudConnectorGroupsResponseBodyIoTCloudConnectorGroupsIoTCloudConnectors) SetServiceType(v string) *ListIoTCloudConnectorGroupsResponseBodyIoTCloudConnectorGroupsIoTCloudConnectors {
	s.ServiceType = &v
	return s
}

type ListIoTCloudConnectorGroupsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListIoTCloudConnectorGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListIoTCloudConnectorGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIoTCloudConnectorGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListIoTCloudConnectorGroupsResponse) SetHeaders(v map[string]*string) *ListIoTCloudConnectorGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListIoTCloudConnectorGroupsResponse) SetStatusCode(v int32) *ListIoTCloudConnectorGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIoTCloudConnectorGroupsResponse) SetBody(v *ListIoTCloudConnectorGroupsResponseBody) *ListIoTCloudConnectorGroupsResponse {
	s.Body = v
	return s
}

type ListIoTCloudConnectorsRequest struct {
	APN                      []*string `json:"APN,omitempty" xml:"APN,omitempty" type:"Repeated"`
	ISP                      []*string `json:"ISP,omitempty" xml:"ISP,omitempty" type:"Repeated"`
	IoTCloudConnectorGroupId *string   `json:"IoTCloudConnectorGroupId,omitempty" xml:"IoTCloudConnectorGroupId,omitempty"`
	IoTCloudConnectorIds     []*string `json:"IoTCloudConnectorIds,omitempty" xml:"IoTCloudConnectorIds,omitempty" type:"Repeated"`
	IoTCloudConnectorName    []*string `json:"IoTCloudConnectorName,omitempty" xml:"IoTCloudConnectorName,omitempty" type:"Repeated"`
	IoTCloudConnectorStatus  []*string `json:"IoTCloudConnectorStatus,omitempty" xml:"IoTCloudConnectorStatus,omitempty" type:"Repeated"`
	IsInGroup                *bool     `json:"IsInGroup,omitempty" xml:"IsInGroup,omitempty"`
	MaxResults               *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken                *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId                 *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VpcId                    []*string `json:"VpcId,omitempty" xml:"VpcId,omitempty" type:"Repeated"`
}

func (s ListIoTCloudConnectorsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIoTCloudConnectorsRequest) GoString() string {
	return s.String()
}

func (s *ListIoTCloudConnectorsRequest) SetAPN(v []*string) *ListIoTCloudConnectorsRequest {
	s.APN = v
	return s
}

func (s *ListIoTCloudConnectorsRequest) SetISP(v []*string) *ListIoTCloudConnectorsRequest {
	s.ISP = v
	return s
}

func (s *ListIoTCloudConnectorsRequest) SetIoTCloudConnectorGroupId(v string) *ListIoTCloudConnectorsRequest {
	s.IoTCloudConnectorGroupId = &v
	return s
}

func (s *ListIoTCloudConnectorsRequest) SetIoTCloudConnectorIds(v []*string) *ListIoTCloudConnectorsRequest {
	s.IoTCloudConnectorIds = v
	return s
}

func (s *ListIoTCloudConnectorsRequest) SetIoTCloudConnectorName(v []*string) *ListIoTCloudConnectorsRequest {
	s.IoTCloudConnectorName = v
	return s
}

func (s *ListIoTCloudConnectorsRequest) SetIoTCloudConnectorStatus(v []*string) *ListIoTCloudConnectorsRequest {
	s.IoTCloudConnectorStatus = v
	return s
}

func (s *ListIoTCloudConnectorsRequest) SetIsInGroup(v bool) *ListIoTCloudConnectorsRequest {
	s.IsInGroup = &v
	return s
}

func (s *ListIoTCloudConnectorsRequest) SetMaxResults(v int32) *ListIoTCloudConnectorsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListIoTCloudConnectorsRequest) SetNextToken(v string) *ListIoTCloudConnectorsRequest {
	s.NextToken = &v
	return s
}

func (s *ListIoTCloudConnectorsRequest) SetRegionId(v string) *ListIoTCloudConnectorsRequest {
	s.RegionId = &v
	return s
}

func (s *ListIoTCloudConnectorsRequest) SetVpcId(v []*string) *ListIoTCloudConnectorsRequest {
	s.VpcId = v
	return s
}

type ListIoTCloudConnectorsResponseBody struct {
	IoTCloudConnectors []*ListIoTCloudConnectorsResponseBodyIoTCloudConnectors `json:"IoTCloudConnectors,omitempty" xml:"IoTCloudConnectors,omitempty" type:"Repeated"`
	MaxResults         *int32                                                  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken          *string                                                 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId          *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount         *int32                                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListIoTCloudConnectorsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIoTCloudConnectorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIoTCloudConnectorsResponseBody) SetIoTCloudConnectors(v []*ListIoTCloudConnectorsResponseBodyIoTCloudConnectors) *ListIoTCloudConnectorsResponseBody {
	s.IoTCloudConnectors = v
	return s
}

func (s *ListIoTCloudConnectorsResponseBody) SetMaxResults(v int32) *ListIoTCloudConnectorsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListIoTCloudConnectorsResponseBody) SetNextToken(v string) *ListIoTCloudConnectorsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListIoTCloudConnectorsResponseBody) SetRequestId(v string) *ListIoTCloudConnectorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIoTCloudConnectorsResponseBody) SetTotalCount(v int32) *ListIoTCloudConnectorsResponseBody {
	s.TotalCount = &v
	return s
}

type ListIoTCloudConnectorsResponseBodyIoTCloudConnectors struct {
	APN                             *string   `json:"APN,omitempty" xml:"APN,omitempty"`
	CreateTime                      *int64    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	GrantAliUid                     *string   `json:"GrantAliUid,omitempty" xml:"GrantAliUid,omitempty"`
	ISP                             *string   `json:"ISP,omitempty" xml:"ISP,omitempty"`
	IoTCloudConnectorBusinessStatus *string   `json:"IoTCloudConnectorBusinessStatus,omitempty" xml:"IoTCloudConnectorBusinessStatus,omitempty"`
	IoTCloudConnectorDescription    *string   `json:"IoTCloudConnectorDescription,omitempty" xml:"IoTCloudConnectorDescription,omitempty"`
	IoTCloudConnectorGroupId        *string   `json:"IoTCloudConnectorGroupId,omitempty" xml:"IoTCloudConnectorGroupId,omitempty"`
	IoTCloudConnectorId             *string   `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	IoTCloudConnectorName           *string   `json:"IoTCloudConnectorName,omitempty" xml:"IoTCloudConnectorName,omitempty"`
	IoTCloudConnectorStatus         *string   `json:"IoTCloudConnectorStatus,omitempty" xml:"IoTCloudConnectorStatus,omitempty"`
	IpFeature                       *string   `json:"IpFeature,omitempty" xml:"IpFeature,omitempty"`
	Mode                            *string   `json:"Mode,omitempty" xml:"Mode,omitempty"`
	ModifyTime                      *int64    `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	RateLimit                       *int64    `json:"RateLimit,omitempty" xml:"RateLimit,omitempty"`
	ServiceType                     *string   `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	Type                            *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	VSwitchList                     []*string `json:"VSwitchList,omitempty" xml:"VSwitchList,omitempty" type:"Repeated"`
	VpcId                           *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	WildcardDomainEnabled           *bool     `json:"WildcardDomainEnabled,omitempty" xml:"WildcardDomainEnabled,omitempty"`
}

func (s ListIoTCloudConnectorsResponseBodyIoTCloudConnectors) String() string {
	return tea.Prettify(s)
}

func (s ListIoTCloudConnectorsResponseBodyIoTCloudConnectors) GoString() string {
	return s.String()
}

func (s *ListIoTCloudConnectorsResponseBodyIoTCloudConnectors) SetAPN(v string) *ListIoTCloudConnectorsResponseBodyIoTCloudConnectors {
	s.APN = &v
	return s
}

func (s *ListIoTCloudConnectorsResponseBodyIoTCloudConnectors) SetCreateTime(v int64) *ListIoTCloudConnectorsResponseBodyIoTCloudConnectors {
	s.CreateTime = &v
	return s
}

func (s *ListIoTCloudConnectorsResponseBodyIoTCloudConnectors) SetGrantAliUid(v string) *ListIoTCloudConnectorsResponseBodyIoTCloudConnectors {
	s.GrantAliUid = &v
	return s
}

func (s *ListIoTCloudConnectorsResponseBodyIoTCloudConnectors) SetISP(v string) *ListIoTCloudConnectorsResponseBodyIoTCloudConnectors {
	s.ISP = &v
	return s
}

func (s *ListIoTCloudConnectorsResponseBodyIoTCloudConnectors) SetIoTCloudConnectorBusinessStatus(v string) *ListIoTCloudConnectorsResponseBodyIoTCloudConnectors {
	s.IoTCloudConnectorBusinessStatus = &v
	return s
}

func (s *ListIoTCloudConnectorsResponseBodyIoTCloudConnectors) SetIoTCloudConnectorDescription(v string) *ListIoTCloudConnectorsResponseBodyIoTCloudConnectors {
	s.IoTCloudConnectorDescription = &v
	return s
}

func (s *ListIoTCloudConnectorsResponseBodyIoTCloudConnectors) SetIoTCloudConnectorGroupId(v string) *ListIoTCloudConnectorsResponseBodyIoTCloudConnectors {
	s.IoTCloudConnectorGroupId = &v
	return s
}

func (s *ListIoTCloudConnectorsResponseBodyIoTCloudConnectors) SetIoTCloudConnectorId(v string) *ListIoTCloudConnectorsResponseBodyIoTCloudConnectors {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *ListIoTCloudConnectorsResponseBodyIoTCloudConnectors) SetIoTCloudConnectorName(v string) *ListIoTCloudConnectorsResponseBodyIoTCloudConnectors {
	s.IoTCloudConnectorName = &v
	return s
}

func (s *ListIoTCloudConnectorsResponseBodyIoTCloudConnectors) SetIoTCloudConnectorStatus(v string) *ListIoTCloudConnectorsResponseBodyIoTCloudConnectors {
	s.IoTCloudConnectorStatus = &v
	return s
}

func (s *ListIoTCloudConnectorsResponseBodyIoTCloudConnectors) SetIpFeature(v string) *ListIoTCloudConnectorsResponseBodyIoTCloudConnectors {
	s.IpFeature = &v
	return s
}

func (s *ListIoTCloudConnectorsResponseBodyIoTCloudConnectors) SetMode(v string) *ListIoTCloudConnectorsResponseBodyIoTCloudConnectors {
	s.Mode = &v
	return s
}

func (s *ListIoTCloudConnectorsResponseBodyIoTCloudConnectors) SetModifyTime(v int64) *ListIoTCloudConnectorsResponseBodyIoTCloudConnectors {
	s.ModifyTime = &v
	return s
}

func (s *ListIoTCloudConnectorsResponseBodyIoTCloudConnectors) SetRateLimit(v int64) *ListIoTCloudConnectorsResponseBodyIoTCloudConnectors {
	s.RateLimit = &v
	return s
}

func (s *ListIoTCloudConnectorsResponseBodyIoTCloudConnectors) SetServiceType(v string) *ListIoTCloudConnectorsResponseBodyIoTCloudConnectors {
	s.ServiceType = &v
	return s
}

func (s *ListIoTCloudConnectorsResponseBodyIoTCloudConnectors) SetType(v string) *ListIoTCloudConnectorsResponseBodyIoTCloudConnectors {
	s.Type = &v
	return s
}

func (s *ListIoTCloudConnectorsResponseBodyIoTCloudConnectors) SetVSwitchList(v []*string) *ListIoTCloudConnectorsResponseBodyIoTCloudConnectors {
	s.VSwitchList = v
	return s
}

func (s *ListIoTCloudConnectorsResponseBodyIoTCloudConnectors) SetVpcId(v string) *ListIoTCloudConnectorsResponseBodyIoTCloudConnectors {
	s.VpcId = &v
	return s
}

func (s *ListIoTCloudConnectorsResponseBodyIoTCloudConnectors) SetWildcardDomainEnabled(v bool) *ListIoTCloudConnectorsResponseBodyIoTCloudConnectors {
	s.WildcardDomainEnabled = &v
	return s
}

type ListIoTCloudConnectorsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListIoTCloudConnectorsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListIoTCloudConnectorsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIoTCloudConnectorsResponse) GoString() string {
	return s.String()
}

func (s *ListIoTCloudConnectorsResponse) SetHeaders(v map[string]*string) *ListIoTCloudConnectorsResponse {
	s.Headers = v
	return s
}

func (s *ListIoTCloudConnectorsResponse) SetStatusCode(v int32) *ListIoTCloudConnectorsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIoTCloudConnectorsResponse) SetBody(v *ListIoTCloudConnectorsResponseBody) *ListIoTCloudConnectorsResponse {
	s.Body = v
	return s
}

type ListIoTCoudConnectorBackhaulRouteRequest struct {
	IoTCloudConnectorId *string `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	MaxResults          *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken           *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListIoTCoudConnectorBackhaulRouteRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIoTCoudConnectorBackhaulRouteRequest) GoString() string {
	return s.String()
}

func (s *ListIoTCoudConnectorBackhaulRouteRequest) SetIoTCloudConnectorId(v string) *ListIoTCoudConnectorBackhaulRouteRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *ListIoTCoudConnectorBackhaulRouteRequest) SetMaxResults(v int32) *ListIoTCoudConnectorBackhaulRouteRequest {
	s.MaxResults = &v
	return s
}

func (s *ListIoTCoudConnectorBackhaulRouteRequest) SetNextToken(v string) *ListIoTCoudConnectorBackhaulRouteRequest {
	s.NextToken = &v
	return s
}

func (s *ListIoTCoudConnectorBackhaulRouteRequest) SetRegionId(v string) *ListIoTCoudConnectorBackhaulRouteRequest {
	s.RegionId = &v
	return s
}

type ListIoTCoudConnectorBackhaulRouteResponseBody struct {
	MaxResults *int32                                                 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                                                `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Routes     []*ListIoTCoudConnectorBackhaulRouteResponseBodyRoutes `json:"Routes,omitempty" xml:"Routes,omitempty" type:"Repeated"`
	TotalCount *int32                                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListIoTCoudConnectorBackhaulRouteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIoTCoudConnectorBackhaulRouteResponseBody) GoString() string {
	return s.String()
}

func (s *ListIoTCoudConnectorBackhaulRouteResponseBody) SetMaxResults(v int32) *ListIoTCoudConnectorBackhaulRouteResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListIoTCoudConnectorBackhaulRouteResponseBody) SetNextToken(v string) *ListIoTCoudConnectorBackhaulRouteResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListIoTCoudConnectorBackhaulRouteResponseBody) SetRequestId(v string) *ListIoTCoudConnectorBackhaulRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIoTCoudConnectorBackhaulRouteResponseBody) SetRoutes(v []*ListIoTCoudConnectorBackhaulRouteResponseBodyRoutes) *ListIoTCoudConnectorBackhaulRouteResponseBody {
	s.Routes = v
	return s
}

func (s *ListIoTCoudConnectorBackhaulRouteResponseBody) SetTotalCount(v int32) *ListIoTCoudConnectorBackhaulRouteResponseBody {
	s.TotalCount = &v
	return s
}

type ListIoTCoudConnectorBackhaulRouteResponseBodyRoutes struct {
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	NextHopId            *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	NextHopType          *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListIoTCoudConnectorBackhaulRouteResponseBodyRoutes) String() string {
	return tea.Prettify(s)
}

func (s ListIoTCoudConnectorBackhaulRouteResponseBodyRoutes) GoString() string {
	return s.String()
}

func (s *ListIoTCoudConnectorBackhaulRouteResponseBodyRoutes) SetDescription(v string) *ListIoTCoudConnectorBackhaulRouteResponseBodyRoutes {
	s.Description = &v
	return s
}

func (s *ListIoTCoudConnectorBackhaulRouteResponseBodyRoutes) SetDestinationCidrBlock(v string) *ListIoTCoudConnectorBackhaulRouteResponseBodyRoutes {
	s.DestinationCidrBlock = &v
	return s
}

func (s *ListIoTCoudConnectorBackhaulRouteResponseBodyRoutes) SetNextHopId(v string) *ListIoTCoudConnectorBackhaulRouteResponseBodyRoutes {
	s.NextHopId = &v
	return s
}

func (s *ListIoTCoudConnectorBackhaulRouteResponseBodyRoutes) SetNextHopType(v string) *ListIoTCoudConnectorBackhaulRouteResponseBodyRoutes {
	s.NextHopType = &v
	return s
}

func (s *ListIoTCoudConnectorBackhaulRouteResponseBodyRoutes) SetStatus(v string) *ListIoTCoudConnectorBackhaulRouteResponseBodyRoutes {
	s.Status = &v
	return s
}

type ListIoTCoudConnectorBackhaulRouteResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListIoTCoudConnectorBackhaulRouteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListIoTCoudConnectorBackhaulRouteResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIoTCoudConnectorBackhaulRouteResponse) GoString() string {
	return s.String()
}

func (s *ListIoTCoudConnectorBackhaulRouteResponse) SetHeaders(v map[string]*string) *ListIoTCoudConnectorBackhaulRouteResponse {
	s.Headers = v
	return s
}

func (s *ListIoTCoudConnectorBackhaulRouteResponse) SetStatusCode(v int32) *ListIoTCoudConnectorBackhaulRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIoTCoudConnectorBackhaulRouteResponse) SetBody(v *ListIoTCoudConnectorBackhaulRouteResponseBody) *ListIoTCoudConnectorBackhaulRouteResponse {
	s.Body = v
	return s
}

type ListIpMappingRulesRequest struct {
	DestinationIps        []*string `json:"DestinationIps,omitempty" xml:"DestinationIps,omitempty" type:"Repeated"`
	IoTCloudConnectorId   *string   `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	IpMappingRuleIds      []*string `json:"IpMappingRuleIds,omitempty" xml:"IpMappingRuleIds,omitempty" type:"Repeated"`
	IpMappingRuleNames    []*string `json:"IpMappingRuleNames,omitempty" xml:"IpMappingRuleNames,omitempty" type:"Repeated"`
	IpMappingRuleStatuses []*string `json:"IpMappingRuleStatuses,omitempty" xml:"IpMappingRuleStatuses,omitempty" type:"Repeated"`
	MappingIps            []*string `json:"MappingIps,omitempty" xml:"MappingIps,omitempty" type:"Repeated"`
	MaxResults            *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken             *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId              *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListIpMappingRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIpMappingRulesRequest) GoString() string {
	return s.String()
}

func (s *ListIpMappingRulesRequest) SetDestinationIps(v []*string) *ListIpMappingRulesRequest {
	s.DestinationIps = v
	return s
}

func (s *ListIpMappingRulesRequest) SetIoTCloudConnectorId(v string) *ListIpMappingRulesRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *ListIpMappingRulesRequest) SetIpMappingRuleIds(v []*string) *ListIpMappingRulesRequest {
	s.IpMappingRuleIds = v
	return s
}

func (s *ListIpMappingRulesRequest) SetIpMappingRuleNames(v []*string) *ListIpMappingRulesRequest {
	s.IpMappingRuleNames = v
	return s
}

func (s *ListIpMappingRulesRequest) SetIpMappingRuleStatuses(v []*string) *ListIpMappingRulesRequest {
	s.IpMappingRuleStatuses = v
	return s
}

func (s *ListIpMappingRulesRequest) SetMappingIps(v []*string) *ListIpMappingRulesRequest {
	s.MappingIps = v
	return s
}

func (s *ListIpMappingRulesRequest) SetMaxResults(v int32) *ListIpMappingRulesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListIpMappingRulesRequest) SetNextToken(v string) *ListIpMappingRulesRequest {
	s.NextToken = &v
	return s
}

func (s *ListIpMappingRulesRequest) SetRegionId(v string) *ListIpMappingRulesRequest {
	s.RegionId = &v
	return s
}

type ListIpMappingRulesResponseBody struct {
	IpMappingRules []*ListIpMappingRulesResponseBodyIpMappingRules `json:"IpMappingRules,omitempty" xml:"IpMappingRules,omitempty" type:"Repeated"`
	MaxResults     *int32                                          `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken      *string                                         `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount     *int32                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListIpMappingRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIpMappingRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListIpMappingRulesResponseBody) SetIpMappingRules(v []*ListIpMappingRulesResponseBodyIpMappingRules) *ListIpMappingRulesResponseBody {
	s.IpMappingRules = v
	return s
}

func (s *ListIpMappingRulesResponseBody) SetMaxResults(v int32) *ListIpMappingRulesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListIpMappingRulesResponseBody) SetNextToken(v string) *ListIpMappingRulesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListIpMappingRulesResponseBody) SetRequestId(v string) *ListIpMappingRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIpMappingRulesResponseBody) SetTotalCount(v int32) *ListIpMappingRulesResponseBody {
	s.TotalCount = &v
	return s
}

type ListIpMappingRulesResponseBodyIpMappingRules struct {
	DestinationIp            *string `json:"DestinationIp,omitempty" xml:"DestinationIp,omitempty"`
	IoTCloudConnectorId      *string `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	IpMappingRuleDescription *string `json:"IpMappingRuleDescription,omitempty" xml:"IpMappingRuleDescription,omitempty"`
	IpMappingRuleId          *string `json:"IpMappingRuleId,omitempty" xml:"IpMappingRuleId,omitempty"`
	IpMappingRuleName        *string `json:"IpMappingRuleName,omitempty" xml:"IpMappingRuleName,omitempty"`
	IpMappingRuleStatus      *string `json:"IpMappingRuleStatus,omitempty" xml:"IpMappingRuleStatus,omitempty"`
	MappingIp                *string `json:"MappingIp,omitempty" xml:"MappingIp,omitempty"`
}

func (s ListIpMappingRulesResponseBodyIpMappingRules) String() string {
	return tea.Prettify(s)
}

func (s ListIpMappingRulesResponseBodyIpMappingRules) GoString() string {
	return s.String()
}

func (s *ListIpMappingRulesResponseBodyIpMappingRules) SetDestinationIp(v string) *ListIpMappingRulesResponseBodyIpMappingRules {
	s.DestinationIp = &v
	return s
}

func (s *ListIpMappingRulesResponseBodyIpMappingRules) SetIoTCloudConnectorId(v string) *ListIpMappingRulesResponseBodyIpMappingRules {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *ListIpMappingRulesResponseBodyIpMappingRules) SetIpMappingRuleDescription(v string) *ListIpMappingRulesResponseBodyIpMappingRules {
	s.IpMappingRuleDescription = &v
	return s
}

func (s *ListIpMappingRulesResponseBodyIpMappingRules) SetIpMappingRuleId(v string) *ListIpMappingRulesResponseBodyIpMappingRules {
	s.IpMappingRuleId = &v
	return s
}

func (s *ListIpMappingRulesResponseBodyIpMappingRules) SetIpMappingRuleName(v string) *ListIpMappingRulesResponseBodyIpMappingRules {
	s.IpMappingRuleName = &v
	return s
}

func (s *ListIpMappingRulesResponseBodyIpMappingRules) SetIpMappingRuleStatus(v string) *ListIpMappingRulesResponseBodyIpMappingRules {
	s.IpMappingRuleStatus = &v
	return s
}

func (s *ListIpMappingRulesResponseBodyIpMappingRules) SetMappingIp(v string) *ListIpMappingRulesResponseBodyIpMappingRules {
	s.MappingIp = &v
	return s
}

type ListIpMappingRulesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListIpMappingRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListIpMappingRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIpMappingRulesResponse) GoString() string {
	return s.String()
}

func (s *ListIpMappingRulesResponse) SetHeaders(v map[string]*string) *ListIpMappingRulesResponse {
	s.Headers = v
	return s
}

func (s *ListIpMappingRulesResponse) SetStatusCode(v int32) *ListIpMappingRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIpMappingRulesResponse) SetBody(v *ListIpMappingRulesResponseBody) *ListIpMappingRulesResponse {
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
	LocalName      *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *ListRegionsResponseBodyRegions) SetRegionEndpoint(v string) *ListRegionsResponseBodyRegions {
	s.RegionEndpoint = &v
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

type ListServiceRequest struct {
	IoTCloudConnectorId *string   `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	MaxResults          *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken           *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId            *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceStatuses    []*string `json:"ResourceStatuses,omitempty" xml:"ResourceStatuses,omitempty" type:"Repeated"`
	ServiceIds          []*string `json:"ServiceIds,omitempty" xml:"ServiceIds,omitempty" type:"Repeated"`
	ServiceNames        []*string `json:"ServiceNames,omitempty" xml:"ServiceNames,omitempty" type:"Repeated"`
}

func (s ListServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceRequest) GoString() string {
	return s.String()
}

func (s *ListServiceRequest) SetIoTCloudConnectorId(v string) *ListServiceRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *ListServiceRequest) SetMaxResults(v int32) *ListServiceRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceRequest) SetNextToken(v string) *ListServiceRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceRequest) SetRegionId(v string) *ListServiceRequest {
	s.RegionId = &v
	return s
}

func (s *ListServiceRequest) SetResourceStatuses(v []*string) *ListServiceRequest {
	s.ResourceStatuses = v
	return s
}

func (s *ListServiceRequest) SetServiceIds(v []*string) *ListServiceRequest {
	s.ServiceIds = v
	return s
}

func (s *ListServiceRequest) SetServiceNames(v []*string) *ListServiceRequest {
	s.ServiceNames = v
	return s
}

type ListServiceResponseBody struct {
	MaxResults *int32                             `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                            `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Services   []*ListServiceResponseBodyServices `json:"Services,omitempty" xml:"Services,omitempty" type:"Repeated"`
	TotalCount *int32                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceResponseBody) SetMaxResults(v int32) *ListServiceResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceResponseBody) SetNextToken(v string) *ListServiceResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceResponseBody) SetRequestId(v string) *ListServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceResponseBody) SetServices(v []*ListServiceResponseBodyServices) *ListServiceResponseBody {
	s.Services = v
	return s
}

func (s *ListServiceResponseBody) SetTotalCount(v int32) *ListServiceResponseBody {
	s.TotalCount = &v
	return s
}

type ListServiceResponseBodyServices struct {
	IoTCloudConnectorId *string `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	ServiceDescription  *string `json:"ServiceDescription,omitempty" xml:"ServiceDescription,omitempty"`
	ServiceId           *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceName         *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	ServiceStatus       *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
}

func (s ListServiceResponseBodyServices) String() string {
	return tea.Prettify(s)
}

func (s ListServiceResponseBodyServices) GoString() string {
	return s.String()
}

func (s *ListServiceResponseBodyServices) SetIoTCloudConnectorId(v string) *ListServiceResponseBodyServices {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *ListServiceResponseBodyServices) SetServiceDescription(v string) *ListServiceResponseBodyServices {
	s.ServiceDescription = &v
	return s
}

func (s *ListServiceResponseBodyServices) SetServiceId(v string) *ListServiceResponseBodyServices {
	s.ServiceId = &v
	return s
}

func (s *ListServiceResponseBodyServices) SetServiceName(v string) *ListServiceResponseBodyServices {
	s.ServiceName = &v
	return s
}

func (s *ListServiceResponseBodyServices) SetServiceStatus(v string) *ListServiceResponseBodyServices {
	s.ServiceStatus = &v
	return s
}

type ListServiceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceResponse) GoString() string {
	return s.String()
}

func (s *ListServiceResponse) SetHeaders(v map[string]*string) *ListServiceResponse {
	s.Headers = v
	return s
}

func (s *ListServiceResponse) SetStatusCode(v int32) *ListServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceResponse) SetBody(v *ListServiceResponseBody) *ListServiceResponse {
	s.Body = v
	return s
}

type ListServiceEntriesRequest struct {
	IoTCloudConnectorId *string   `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	MaxResults          *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken           *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId            *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceEntryIds     []*string `json:"ServiceEntryIds,omitempty" xml:"ServiceEntryIds,omitempty" type:"Repeated"`
	ServiceEntryName    []*string `json:"ServiceEntryName,omitempty" xml:"ServiceEntryName,omitempty" type:"Repeated"`
	ServiceEntryStatus  []*string `json:"ServiceEntryStatus,omitempty" xml:"ServiceEntryStatus,omitempty" type:"Repeated"`
	ServiceId           *string   `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	Target              []*string `json:"Target,omitempty" xml:"Target,omitempty" type:"Repeated"`
	TargetType          []*string `json:"TargetType,omitempty" xml:"TargetType,omitempty" type:"Repeated"`
}

func (s ListServiceEntriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceEntriesRequest) GoString() string {
	return s.String()
}

func (s *ListServiceEntriesRequest) SetIoTCloudConnectorId(v string) *ListServiceEntriesRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *ListServiceEntriesRequest) SetMaxResults(v int32) *ListServiceEntriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceEntriesRequest) SetNextToken(v string) *ListServiceEntriesRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceEntriesRequest) SetRegionId(v string) *ListServiceEntriesRequest {
	s.RegionId = &v
	return s
}

func (s *ListServiceEntriesRequest) SetServiceEntryIds(v []*string) *ListServiceEntriesRequest {
	s.ServiceEntryIds = v
	return s
}

func (s *ListServiceEntriesRequest) SetServiceEntryName(v []*string) *ListServiceEntriesRequest {
	s.ServiceEntryName = v
	return s
}

func (s *ListServiceEntriesRequest) SetServiceEntryStatus(v []*string) *ListServiceEntriesRequest {
	s.ServiceEntryStatus = v
	return s
}

func (s *ListServiceEntriesRequest) SetServiceId(v string) *ListServiceEntriesRequest {
	s.ServiceId = &v
	return s
}

func (s *ListServiceEntriesRequest) SetTarget(v []*string) *ListServiceEntriesRequest {
	s.Target = v
	return s
}

func (s *ListServiceEntriesRequest) SetTargetType(v []*string) *ListServiceEntriesRequest {
	s.TargetType = v
	return s
}

type ListServiceEntriesResponseBody struct {
	MaxResults     *int32                                          `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken      *string                                         `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceEntries []*ListServiceEntriesResponseBodyServiceEntries `json:"ServiceEntries,omitempty" xml:"ServiceEntries,omitempty" type:"Repeated"`
	TotalCount     *int32                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServiceEntriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceEntriesResponseBody) SetMaxResults(v int32) *ListServiceEntriesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceEntriesResponseBody) SetNextToken(v string) *ListServiceEntriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceEntriesResponseBody) SetRequestId(v string) *ListServiceEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceEntriesResponseBody) SetServiceEntries(v []*ListServiceEntriesResponseBodyServiceEntries) *ListServiceEntriesResponseBody {
	s.ServiceEntries = v
	return s
}

func (s *ListServiceEntriesResponseBody) SetTotalCount(v int32) *ListServiceEntriesResponseBody {
	s.TotalCount = &v
	return s
}

type ListServiceEntriesResponseBodyServiceEntries struct {
	ServiceEntryDescription *string `json:"ServiceEntryDescription,omitempty" xml:"ServiceEntryDescription,omitempty"`
	ServiceEntryId          *string `json:"ServiceEntryId,omitempty" xml:"ServiceEntryId,omitempty"`
	ServiceEntryName        *string `json:"ServiceEntryName,omitempty" xml:"ServiceEntryName,omitempty"`
	ServiceEntryStatus      *string `json:"ServiceEntryStatus,omitempty" xml:"ServiceEntryStatus,omitempty"`
	ServiceId               *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	Target                  *string `json:"Target,omitempty" xml:"Target,omitempty"`
	TargetType              *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ListServiceEntriesResponseBodyServiceEntries) String() string {
	return tea.Prettify(s)
}

func (s ListServiceEntriesResponseBodyServiceEntries) GoString() string {
	return s.String()
}

func (s *ListServiceEntriesResponseBodyServiceEntries) SetServiceEntryDescription(v string) *ListServiceEntriesResponseBodyServiceEntries {
	s.ServiceEntryDescription = &v
	return s
}

func (s *ListServiceEntriesResponseBodyServiceEntries) SetServiceEntryId(v string) *ListServiceEntriesResponseBodyServiceEntries {
	s.ServiceEntryId = &v
	return s
}

func (s *ListServiceEntriesResponseBodyServiceEntries) SetServiceEntryName(v string) *ListServiceEntriesResponseBodyServiceEntries {
	s.ServiceEntryName = &v
	return s
}

func (s *ListServiceEntriesResponseBodyServiceEntries) SetServiceEntryStatus(v string) *ListServiceEntriesResponseBodyServiceEntries {
	s.ServiceEntryStatus = &v
	return s
}

func (s *ListServiceEntriesResponseBodyServiceEntries) SetServiceId(v string) *ListServiceEntriesResponseBodyServiceEntries {
	s.ServiceId = &v
	return s
}

func (s *ListServiceEntriesResponseBodyServiceEntries) SetTarget(v string) *ListServiceEntriesResponseBodyServiceEntries {
	s.Target = &v
	return s
}

func (s *ListServiceEntriesResponseBodyServiceEntries) SetTargetType(v string) *ListServiceEntriesResponseBodyServiceEntries {
	s.TargetType = &v
	return s
}

type ListServiceEntriesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListServiceEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListServiceEntriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceEntriesResponse) GoString() string {
	return s.String()
}

func (s *ListServiceEntriesResponse) SetHeaders(v map[string]*string) *ListServiceEntriesResponse {
	s.Headers = v
	return s
}

func (s *ListServiceEntriesResponse) SetStatusCode(v int32) *ListServiceEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceEntriesResponse) SetBody(v *ListServiceEntriesResponseBody) *ListServiceEntriesResponse {
	s.Body = v
	return s
}

type MoveAuthorizationRuleToDNSServiceRequest struct {
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	ClientToken         *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun              *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorId *string `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s MoveAuthorizationRuleToDNSServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveAuthorizationRuleToDNSServiceRequest) GoString() string {
	return s.String()
}

func (s *MoveAuthorizationRuleToDNSServiceRequest) SetAuthorizationRuleId(v string) *MoveAuthorizationRuleToDNSServiceRequest {
	s.AuthorizationRuleId = &v
	return s
}

func (s *MoveAuthorizationRuleToDNSServiceRequest) SetClientToken(v string) *MoveAuthorizationRuleToDNSServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *MoveAuthorizationRuleToDNSServiceRequest) SetDryRun(v bool) *MoveAuthorizationRuleToDNSServiceRequest {
	s.DryRun = &v
	return s
}

func (s *MoveAuthorizationRuleToDNSServiceRequest) SetIoTCloudConnectorId(v string) *MoveAuthorizationRuleToDNSServiceRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *MoveAuthorizationRuleToDNSServiceRequest) SetRegionId(v string) *MoveAuthorizationRuleToDNSServiceRequest {
	s.RegionId = &v
	return s
}

type MoveAuthorizationRuleToDNSServiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MoveAuthorizationRuleToDNSServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveAuthorizationRuleToDNSServiceResponseBody) GoString() string {
	return s.String()
}

func (s *MoveAuthorizationRuleToDNSServiceResponseBody) SetRequestId(v string) *MoveAuthorizationRuleToDNSServiceResponseBody {
	s.RequestId = &v
	return s
}

type MoveAuthorizationRuleToDNSServiceResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *MoveAuthorizationRuleToDNSServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MoveAuthorizationRuleToDNSServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveAuthorizationRuleToDNSServiceResponse) GoString() string {
	return s.String()
}

func (s *MoveAuthorizationRuleToDNSServiceResponse) SetHeaders(v map[string]*string) *MoveAuthorizationRuleToDNSServiceResponse {
	s.Headers = v
	return s
}

func (s *MoveAuthorizationRuleToDNSServiceResponse) SetStatusCode(v int32) *MoveAuthorizationRuleToDNSServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveAuthorizationRuleToDNSServiceResponse) SetBody(v *MoveAuthorizationRuleToDNSServiceResponseBody) *MoveAuthorizationRuleToDNSServiceResponse {
	s.Body = v
	return s
}

type MoveGroupAuthorizationRuleToDNSServiceRequest struct {
	AuthorizationRuleId      *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	ClientToken              *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                   *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorGroupId *string `json:"IoTCloudConnectorGroupId,omitempty" xml:"IoTCloudConnectorGroupId,omitempty"`
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s MoveGroupAuthorizationRuleToDNSServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveGroupAuthorizationRuleToDNSServiceRequest) GoString() string {
	return s.String()
}

func (s *MoveGroupAuthorizationRuleToDNSServiceRequest) SetAuthorizationRuleId(v string) *MoveGroupAuthorizationRuleToDNSServiceRequest {
	s.AuthorizationRuleId = &v
	return s
}

func (s *MoveGroupAuthorizationRuleToDNSServiceRequest) SetClientToken(v string) *MoveGroupAuthorizationRuleToDNSServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *MoveGroupAuthorizationRuleToDNSServiceRequest) SetDryRun(v bool) *MoveGroupAuthorizationRuleToDNSServiceRequest {
	s.DryRun = &v
	return s
}

func (s *MoveGroupAuthorizationRuleToDNSServiceRequest) SetIoTCloudConnectorGroupId(v string) *MoveGroupAuthorizationRuleToDNSServiceRequest {
	s.IoTCloudConnectorGroupId = &v
	return s
}

func (s *MoveGroupAuthorizationRuleToDNSServiceRequest) SetRegionId(v string) *MoveGroupAuthorizationRuleToDNSServiceRequest {
	s.RegionId = &v
	return s
}

type MoveGroupAuthorizationRuleToDNSServiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MoveGroupAuthorizationRuleToDNSServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveGroupAuthorizationRuleToDNSServiceResponseBody) GoString() string {
	return s.String()
}

func (s *MoveGroupAuthorizationRuleToDNSServiceResponseBody) SetRequestId(v string) *MoveGroupAuthorizationRuleToDNSServiceResponseBody {
	s.RequestId = &v
	return s
}

type MoveGroupAuthorizationRuleToDNSServiceResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *MoveGroupAuthorizationRuleToDNSServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MoveGroupAuthorizationRuleToDNSServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveGroupAuthorizationRuleToDNSServiceResponse) GoString() string {
	return s.String()
}

func (s *MoveGroupAuthorizationRuleToDNSServiceResponse) SetHeaders(v map[string]*string) *MoveGroupAuthorizationRuleToDNSServiceResponse {
	s.Headers = v
	return s
}

func (s *MoveGroupAuthorizationRuleToDNSServiceResponse) SetStatusCode(v int32) *MoveGroupAuthorizationRuleToDNSServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveGroupAuthorizationRuleToDNSServiceResponse) SetBody(v *MoveGroupAuthorizationRuleToDNSServiceResponseBody) *MoveGroupAuthorizationRuleToDNSServiceResponse {
	s.Body = v
	return s
}

type OpenIoTCloudConnectorServiceRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s OpenIoTCloudConnectorServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenIoTCloudConnectorServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenIoTCloudConnectorServiceRequest) SetRegionId(v string) *OpenIoTCloudConnectorServiceRequest {
	s.RegionId = &v
	return s
}

type OpenIoTCloudConnectorServiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenIoTCloudConnectorServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenIoTCloudConnectorServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenIoTCloudConnectorServiceResponseBody) SetRequestId(v string) *OpenIoTCloudConnectorServiceResponseBody {
	s.RequestId = &v
	return s
}

type OpenIoTCloudConnectorServiceResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OpenIoTCloudConnectorServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpenIoTCloudConnectorServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenIoTCloudConnectorServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenIoTCloudConnectorServiceResponse) SetHeaders(v map[string]*string) *OpenIoTCloudConnectorServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenIoTCloudConnectorServiceResponse) SetStatusCode(v int32) *OpenIoTCloudConnectorServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenIoTCloudConnectorServiceResponse) SetBody(v *OpenIoTCloudConnectorServiceResponseBody) *OpenIoTCloudConnectorServiceResponse {
	s.Body = v
	return s
}

type RemoveIoTCloudConnectorFromGroupRequest struct {
	ClientToken              *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                   *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorGroupId *string   `json:"IoTCloudConnectorGroupId,omitempty" xml:"IoTCloudConnectorGroupId,omitempty"`
	IoTCloudConnectorId      []*string `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty" type:"Repeated"`
	RegionId                 *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RemoveIoTCloudConnectorFromGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveIoTCloudConnectorFromGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveIoTCloudConnectorFromGroupRequest) SetClientToken(v string) *RemoveIoTCloudConnectorFromGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *RemoveIoTCloudConnectorFromGroupRequest) SetDryRun(v bool) *RemoveIoTCloudConnectorFromGroupRequest {
	s.DryRun = &v
	return s
}

func (s *RemoveIoTCloudConnectorFromGroupRequest) SetIoTCloudConnectorGroupId(v string) *RemoveIoTCloudConnectorFromGroupRequest {
	s.IoTCloudConnectorGroupId = &v
	return s
}

func (s *RemoveIoTCloudConnectorFromGroupRequest) SetIoTCloudConnectorId(v []*string) *RemoveIoTCloudConnectorFromGroupRequest {
	s.IoTCloudConnectorId = v
	return s
}

func (s *RemoveIoTCloudConnectorFromGroupRequest) SetRegionId(v string) *RemoveIoTCloudConnectorFromGroupRequest {
	s.RegionId = &v
	return s
}

type RemoveIoTCloudConnectorFromGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveIoTCloudConnectorFromGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveIoTCloudConnectorFromGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveIoTCloudConnectorFromGroupResponseBody) SetRequestId(v string) *RemoveIoTCloudConnectorFromGroupResponseBody {
	s.RequestId = &v
	return s
}

type RemoveIoTCloudConnectorFromGroupResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveIoTCloudConnectorFromGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveIoTCloudConnectorFromGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveIoTCloudConnectorFromGroupResponse) GoString() string {
	return s.String()
}

func (s *RemoveIoTCloudConnectorFromGroupResponse) SetHeaders(v map[string]*string) *RemoveIoTCloudConnectorFromGroupResponse {
	s.Headers = v
	return s
}

func (s *RemoveIoTCloudConnectorFromGroupResponse) SetStatusCode(v int32) *RemoveIoTCloudConnectorFromGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveIoTCloudConnectorFromGroupResponse) SetBody(v *RemoveIoTCloudConnectorFromGroupResponseBody) *RemoveIoTCloudConnectorFromGroupResponse {
	s.Body = v
	return s
}

type RevertIoTCloudConnectorRequest struct {
	ClientToken         *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun              *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorId *string `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RevertIoTCloudConnectorRequest) String() string {
	return tea.Prettify(s)
}

func (s RevertIoTCloudConnectorRequest) GoString() string {
	return s.String()
}

func (s *RevertIoTCloudConnectorRequest) SetClientToken(v string) *RevertIoTCloudConnectorRequest {
	s.ClientToken = &v
	return s
}

func (s *RevertIoTCloudConnectorRequest) SetDryRun(v bool) *RevertIoTCloudConnectorRequest {
	s.DryRun = &v
	return s
}

func (s *RevertIoTCloudConnectorRequest) SetIoTCloudConnectorId(v string) *RevertIoTCloudConnectorRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *RevertIoTCloudConnectorRequest) SetRegionId(v string) *RevertIoTCloudConnectorRequest {
	s.RegionId = &v
	return s
}

type RevertIoTCloudConnectorResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s RevertIoTCloudConnectorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RevertIoTCloudConnectorResponseBody) GoString() string {
	return s.String()
}

func (s *RevertIoTCloudConnectorResponseBody) SetRequestId(v string) *RevertIoTCloudConnectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevertIoTCloudConnectorResponseBody) SetResourceId(v string) *RevertIoTCloudConnectorResponseBody {
	s.ResourceId = &v
	return s
}

type RevertIoTCloudConnectorResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RevertIoTCloudConnectorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RevertIoTCloudConnectorResponse) String() string {
	return tea.Prettify(s)
}

func (s RevertIoTCloudConnectorResponse) GoString() string {
	return s.String()
}

func (s *RevertIoTCloudConnectorResponse) SetHeaders(v map[string]*string) *RevertIoTCloudConnectorResponse {
	s.Headers = v
	return s
}

func (s *RevertIoTCloudConnectorResponse) SetStatusCode(v int32) *RevertIoTCloudConnectorResponse {
	s.StatusCode = &v
	return s
}

func (s *RevertIoTCloudConnectorResponse) SetBody(v *RevertIoTCloudConnectorResponseBody) *RevertIoTCloudConnectorResponse {
	s.Body = v
	return s
}

type SubmitDiagnoseTaskForSingleCardRequest struct {
	BeginTime           *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	Destination         *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	DestinationType     *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	EndTime             *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IoTCloudConnectorId *string `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceUid         *int64  `json:"ResourceUid,omitempty" xml:"ResourceUid,omitempty"`
	Source              *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceType          *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
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

func (s *SubmitDiagnoseTaskForSingleCardRequest) SetDestinationType(v string) *SubmitDiagnoseTaskForSingleCardRequest {
	s.DestinationType = &v
	return s
}

func (s *SubmitDiagnoseTaskForSingleCardRequest) SetEndTime(v int64) *SubmitDiagnoseTaskForSingleCardRequest {
	s.EndTime = &v
	return s
}

func (s *SubmitDiagnoseTaskForSingleCardRequest) SetIoTCloudConnectorId(v string) *SubmitDiagnoseTaskForSingleCardRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *SubmitDiagnoseTaskForSingleCardRequest) SetRegionId(v string) *SubmitDiagnoseTaskForSingleCardRequest {
	s.RegionId = &v
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

func (s *SubmitDiagnoseTaskForSingleCardRequest) SetSourceType(v string) *SubmitDiagnoseTaskForSingleCardRequest {
	s.SourceType = &v
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

type UpdateAuthorizationRuleAttributeRequest struct {
	AuthorizationRuleDescription *string   `json:"AuthorizationRuleDescription,omitempty" xml:"AuthorizationRuleDescription,omitempty"`
	AuthorizationRuleId          *string   `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	AuthorizationRuleName        *string   `json:"AuthorizationRuleName,omitempty" xml:"AuthorizationRuleName,omitempty"`
	ClientToken                  *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Destination                  *string   `json:"Destination,omitempty" xml:"Destination,omitempty"`
	DestinationPort              *string   `json:"DestinationPort,omitempty" xml:"DestinationPort,omitempty"`
	DestinationType              *string   `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	DryRun                       *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorId          *string   `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	Policy                       *string   `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Protocol                     *string   `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	RegionId                     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SourceCidrs                  []*string `json:"SourceCidrs,omitempty" xml:"SourceCidrs,omitempty" type:"Repeated"`
}

func (s UpdateAuthorizationRuleAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAuthorizationRuleAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateAuthorizationRuleAttributeRequest) SetAuthorizationRuleDescription(v string) *UpdateAuthorizationRuleAttributeRequest {
	s.AuthorizationRuleDescription = &v
	return s
}

func (s *UpdateAuthorizationRuleAttributeRequest) SetAuthorizationRuleId(v string) *UpdateAuthorizationRuleAttributeRequest {
	s.AuthorizationRuleId = &v
	return s
}

func (s *UpdateAuthorizationRuleAttributeRequest) SetAuthorizationRuleName(v string) *UpdateAuthorizationRuleAttributeRequest {
	s.AuthorizationRuleName = &v
	return s
}

func (s *UpdateAuthorizationRuleAttributeRequest) SetClientToken(v string) *UpdateAuthorizationRuleAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAuthorizationRuleAttributeRequest) SetDestination(v string) *UpdateAuthorizationRuleAttributeRequest {
	s.Destination = &v
	return s
}

func (s *UpdateAuthorizationRuleAttributeRequest) SetDestinationPort(v string) *UpdateAuthorizationRuleAttributeRequest {
	s.DestinationPort = &v
	return s
}

func (s *UpdateAuthorizationRuleAttributeRequest) SetDestinationType(v string) *UpdateAuthorizationRuleAttributeRequest {
	s.DestinationType = &v
	return s
}

func (s *UpdateAuthorizationRuleAttributeRequest) SetDryRun(v bool) *UpdateAuthorizationRuleAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateAuthorizationRuleAttributeRequest) SetIoTCloudConnectorId(v string) *UpdateAuthorizationRuleAttributeRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *UpdateAuthorizationRuleAttributeRequest) SetPolicy(v string) *UpdateAuthorizationRuleAttributeRequest {
	s.Policy = &v
	return s
}

func (s *UpdateAuthorizationRuleAttributeRequest) SetProtocol(v string) *UpdateAuthorizationRuleAttributeRequest {
	s.Protocol = &v
	return s
}

func (s *UpdateAuthorizationRuleAttributeRequest) SetRegionId(v string) *UpdateAuthorizationRuleAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateAuthorizationRuleAttributeRequest) SetSourceCidrs(v []*string) *UpdateAuthorizationRuleAttributeRequest {
	s.SourceCidrs = v
	return s
}

type UpdateAuthorizationRuleAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAuthorizationRuleAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAuthorizationRuleAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAuthorizationRuleAttributeResponseBody) SetRequestId(v string) *UpdateAuthorizationRuleAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAuthorizationRuleAttributeResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAuthorizationRuleAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAuthorizationRuleAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAuthorizationRuleAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateAuthorizationRuleAttributeResponse) SetHeaders(v map[string]*string) *UpdateAuthorizationRuleAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateAuthorizationRuleAttributeResponse) SetStatusCode(v int32) *UpdateAuthorizationRuleAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAuthorizationRuleAttributeResponse) SetBody(v *UpdateAuthorizationRuleAttributeResponseBody) *UpdateAuthorizationRuleAttributeResponse {
	s.Body = v
	return s
}

type UpdateConnectionPoolAttributeRequest struct {
	Cidrs                     []*string `json:"Cidrs,omitempty" xml:"Cidrs,omitempty" type:"Repeated"`
	ClientToken               *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConnectionPoolDescription *string   `json:"ConnectionPoolDescription,omitempty" xml:"ConnectionPoolDescription,omitempty"`
	ConnectionPoolId          *string   `json:"ConnectionPoolId,omitempty" xml:"ConnectionPoolId,omitempty"`
	ConnectionPoolName        *string   `json:"ConnectionPoolName,omitempty" xml:"ConnectionPoolName,omitempty"`
	Count                     *int64    `json:"Count,omitempty" xml:"Count,omitempty"`
	DryRun                    *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorId       *string   `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	RegionId                  *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateConnectionPoolAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateConnectionPoolAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateConnectionPoolAttributeRequest) SetCidrs(v []*string) *UpdateConnectionPoolAttributeRequest {
	s.Cidrs = v
	return s
}

func (s *UpdateConnectionPoolAttributeRequest) SetClientToken(v string) *UpdateConnectionPoolAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateConnectionPoolAttributeRequest) SetConnectionPoolDescription(v string) *UpdateConnectionPoolAttributeRequest {
	s.ConnectionPoolDescription = &v
	return s
}

func (s *UpdateConnectionPoolAttributeRequest) SetConnectionPoolId(v string) *UpdateConnectionPoolAttributeRequest {
	s.ConnectionPoolId = &v
	return s
}

func (s *UpdateConnectionPoolAttributeRequest) SetConnectionPoolName(v string) *UpdateConnectionPoolAttributeRequest {
	s.ConnectionPoolName = &v
	return s
}

func (s *UpdateConnectionPoolAttributeRequest) SetCount(v int64) *UpdateConnectionPoolAttributeRequest {
	s.Count = &v
	return s
}

func (s *UpdateConnectionPoolAttributeRequest) SetDryRun(v bool) *UpdateConnectionPoolAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateConnectionPoolAttributeRequest) SetIoTCloudConnectorId(v string) *UpdateConnectionPoolAttributeRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *UpdateConnectionPoolAttributeRequest) SetRegionId(v string) *UpdateConnectionPoolAttributeRequest {
	s.RegionId = &v
	return s
}

type UpdateConnectionPoolAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateConnectionPoolAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateConnectionPoolAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConnectionPoolAttributeResponseBody) SetRequestId(v string) *UpdateConnectionPoolAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateConnectionPoolAttributeResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateConnectionPoolAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateConnectionPoolAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateConnectionPoolAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateConnectionPoolAttributeResponse) SetHeaders(v map[string]*string) *UpdateConnectionPoolAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateConnectionPoolAttributeResponse) SetStatusCode(v int32) *UpdateConnectionPoolAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateConnectionPoolAttributeResponse) SetBody(v *UpdateConnectionPoolAttributeResponseBody) *UpdateConnectionPoolAttributeResponse {
	s.Body = v
	return s
}

type UpdateDNSServiceRuleAttributeRequest struct {
	ClientToken               *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DNSServiceRuleDescription *string `json:"DNSServiceRuleDescription,omitempty" xml:"DNSServiceRuleDescription,omitempty"`
	DNSServiceRuleId          *string `json:"DNSServiceRuleId,omitempty" xml:"DNSServiceRuleId,omitempty"`
	DNSServiceRuleName        *string `json:"DNSServiceRuleName,omitempty" xml:"DNSServiceRuleName,omitempty"`
	Destination               *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	DryRun                    *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorId       *string `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	RegionId                  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceType               *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	Source                    *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s UpdateDNSServiceRuleAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDNSServiceRuleAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateDNSServiceRuleAttributeRequest) SetClientToken(v string) *UpdateDNSServiceRuleAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateDNSServiceRuleAttributeRequest) SetDNSServiceRuleDescription(v string) *UpdateDNSServiceRuleAttributeRequest {
	s.DNSServiceRuleDescription = &v
	return s
}

func (s *UpdateDNSServiceRuleAttributeRequest) SetDNSServiceRuleId(v string) *UpdateDNSServiceRuleAttributeRequest {
	s.DNSServiceRuleId = &v
	return s
}

func (s *UpdateDNSServiceRuleAttributeRequest) SetDNSServiceRuleName(v string) *UpdateDNSServiceRuleAttributeRequest {
	s.DNSServiceRuleName = &v
	return s
}

func (s *UpdateDNSServiceRuleAttributeRequest) SetDestination(v string) *UpdateDNSServiceRuleAttributeRequest {
	s.Destination = &v
	return s
}

func (s *UpdateDNSServiceRuleAttributeRequest) SetDryRun(v bool) *UpdateDNSServiceRuleAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateDNSServiceRuleAttributeRequest) SetIoTCloudConnectorId(v string) *UpdateDNSServiceRuleAttributeRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *UpdateDNSServiceRuleAttributeRequest) SetRegionId(v string) *UpdateDNSServiceRuleAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateDNSServiceRuleAttributeRequest) SetServiceType(v string) *UpdateDNSServiceRuleAttributeRequest {
	s.ServiceType = &v
	return s
}

func (s *UpdateDNSServiceRuleAttributeRequest) SetSource(v string) *UpdateDNSServiceRuleAttributeRequest {
	s.Source = &v
	return s
}

type UpdateDNSServiceRuleAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDNSServiceRuleAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDNSServiceRuleAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDNSServiceRuleAttributeResponseBody) SetRequestId(v string) *UpdateDNSServiceRuleAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDNSServiceRuleAttributeResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateDNSServiceRuleAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDNSServiceRuleAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDNSServiceRuleAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateDNSServiceRuleAttributeResponse) SetHeaders(v map[string]*string) *UpdateDNSServiceRuleAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateDNSServiceRuleAttributeResponse) SetStatusCode(v int32) *UpdateDNSServiceRuleAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDNSServiceRuleAttributeResponse) SetBody(v *UpdateDNSServiceRuleAttributeResponseBody) *UpdateDNSServiceRuleAttributeResponse {
	s.Body = v
	return s
}

type UpdateGroupAuthorizationRuleAttributeRequest struct {
	AuthorizationRuleDescription *string   `json:"AuthorizationRuleDescription,omitempty" xml:"AuthorizationRuleDescription,omitempty"`
	AuthorizationRuleId          *string   `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	AuthorizationRuleName        *string   `json:"AuthorizationRuleName,omitempty" xml:"AuthorizationRuleName,omitempty"`
	ClientToken                  *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Destination                  *string   `json:"Destination,omitempty" xml:"Destination,omitempty"`
	DestinationPort              *string   `json:"DestinationPort,omitempty" xml:"DestinationPort,omitempty"`
	DestinationType              *string   `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	DryRun                       *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorGroupId     *string   `json:"IoTCloudConnectorGroupId,omitempty" xml:"IoTCloudConnectorGroupId,omitempty"`
	Policy                       *string   `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Protocol                     *string   `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	RegionId                     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SourceCidrs                  []*string `json:"SourceCidrs,omitempty" xml:"SourceCidrs,omitempty" type:"Repeated"`
}

func (s UpdateGroupAuthorizationRuleAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupAuthorizationRuleAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupAuthorizationRuleAttributeRequest) SetAuthorizationRuleDescription(v string) *UpdateGroupAuthorizationRuleAttributeRequest {
	s.AuthorizationRuleDescription = &v
	return s
}

func (s *UpdateGroupAuthorizationRuleAttributeRequest) SetAuthorizationRuleId(v string) *UpdateGroupAuthorizationRuleAttributeRequest {
	s.AuthorizationRuleId = &v
	return s
}

func (s *UpdateGroupAuthorizationRuleAttributeRequest) SetAuthorizationRuleName(v string) *UpdateGroupAuthorizationRuleAttributeRequest {
	s.AuthorizationRuleName = &v
	return s
}

func (s *UpdateGroupAuthorizationRuleAttributeRequest) SetClientToken(v string) *UpdateGroupAuthorizationRuleAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateGroupAuthorizationRuleAttributeRequest) SetDestination(v string) *UpdateGroupAuthorizationRuleAttributeRequest {
	s.Destination = &v
	return s
}

func (s *UpdateGroupAuthorizationRuleAttributeRequest) SetDestinationPort(v string) *UpdateGroupAuthorizationRuleAttributeRequest {
	s.DestinationPort = &v
	return s
}

func (s *UpdateGroupAuthorizationRuleAttributeRequest) SetDestinationType(v string) *UpdateGroupAuthorizationRuleAttributeRequest {
	s.DestinationType = &v
	return s
}

func (s *UpdateGroupAuthorizationRuleAttributeRequest) SetDryRun(v bool) *UpdateGroupAuthorizationRuleAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateGroupAuthorizationRuleAttributeRequest) SetIoTCloudConnectorGroupId(v string) *UpdateGroupAuthorizationRuleAttributeRequest {
	s.IoTCloudConnectorGroupId = &v
	return s
}

func (s *UpdateGroupAuthorizationRuleAttributeRequest) SetPolicy(v string) *UpdateGroupAuthorizationRuleAttributeRequest {
	s.Policy = &v
	return s
}

func (s *UpdateGroupAuthorizationRuleAttributeRequest) SetProtocol(v string) *UpdateGroupAuthorizationRuleAttributeRequest {
	s.Protocol = &v
	return s
}

func (s *UpdateGroupAuthorizationRuleAttributeRequest) SetRegionId(v string) *UpdateGroupAuthorizationRuleAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateGroupAuthorizationRuleAttributeRequest) SetSourceCidrs(v []*string) *UpdateGroupAuthorizationRuleAttributeRequest {
	s.SourceCidrs = v
	return s
}

type UpdateGroupAuthorizationRuleAttributeResponseBody struct {
	AuthorizationRuleId      *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	IoTCloudConnectorGroupId *string `json:"IoTCloudConnectorGroupId,omitempty" xml:"IoTCloudConnectorGroupId,omitempty"`
	RequestId                *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateGroupAuthorizationRuleAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupAuthorizationRuleAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGroupAuthorizationRuleAttributeResponseBody) SetAuthorizationRuleId(v string) *UpdateGroupAuthorizationRuleAttributeResponseBody {
	s.AuthorizationRuleId = &v
	return s
}

func (s *UpdateGroupAuthorizationRuleAttributeResponseBody) SetIoTCloudConnectorGroupId(v string) *UpdateGroupAuthorizationRuleAttributeResponseBody {
	s.IoTCloudConnectorGroupId = &v
	return s
}

func (s *UpdateGroupAuthorizationRuleAttributeResponseBody) SetRequestId(v string) *UpdateGroupAuthorizationRuleAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateGroupAuthorizationRuleAttributeResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateGroupAuthorizationRuleAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateGroupAuthorizationRuleAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupAuthorizationRuleAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateGroupAuthorizationRuleAttributeResponse) SetHeaders(v map[string]*string) *UpdateGroupAuthorizationRuleAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateGroupAuthorizationRuleAttributeResponse) SetStatusCode(v int32) *UpdateGroupAuthorizationRuleAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGroupAuthorizationRuleAttributeResponse) SetBody(v *UpdateGroupAuthorizationRuleAttributeResponseBody) *UpdateGroupAuthorizationRuleAttributeResponse {
	s.Body = v
	return s
}

type UpdateGroupDNSServiceRuleAttributeRequest struct {
	ClientToken               *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DNSServiceRuleDescription *string `json:"DNSServiceRuleDescription,omitempty" xml:"DNSServiceRuleDescription,omitempty"`
	DNSServiceRuleId          *string `json:"DNSServiceRuleId,omitempty" xml:"DNSServiceRuleId,omitempty"`
	DNSServiceRuleName        *string `json:"DNSServiceRuleName,omitempty" xml:"DNSServiceRuleName,omitempty"`
	Destination               *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	DryRun                    *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorGroupId  *string `json:"IoTCloudConnectorGroupId,omitempty" xml:"IoTCloudConnectorGroupId,omitempty"`
	RegionId                  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceType               *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	Source                    *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s UpdateGroupDNSServiceRuleAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupDNSServiceRuleAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupDNSServiceRuleAttributeRequest) SetClientToken(v string) *UpdateGroupDNSServiceRuleAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateGroupDNSServiceRuleAttributeRequest) SetDNSServiceRuleDescription(v string) *UpdateGroupDNSServiceRuleAttributeRequest {
	s.DNSServiceRuleDescription = &v
	return s
}

func (s *UpdateGroupDNSServiceRuleAttributeRequest) SetDNSServiceRuleId(v string) *UpdateGroupDNSServiceRuleAttributeRequest {
	s.DNSServiceRuleId = &v
	return s
}

func (s *UpdateGroupDNSServiceRuleAttributeRequest) SetDNSServiceRuleName(v string) *UpdateGroupDNSServiceRuleAttributeRequest {
	s.DNSServiceRuleName = &v
	return s
}

func (s *UpdateGroupDNSServiceRuleAttributeRequest) SetDestination(v string) *UpdateGroupDNSServiceRuleAttributeRequest {
	s.Destination = &v
	return s
}

func (s *UpdateGroupDNSServiceRuleAttributeRequest) SetDryRun(v bool) *UpdateGroupDNSServiceRuleAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateGroupDNSServiceRuleAttributeRequest) SetIoTCloudConnectorGroupId(v string) *UpdateGroupDNSServiceRuleAttributeRequest {
	s.IoTCloudConnectorGroupId = &v
	return s
}

func (s *UpdateGroupDNSServiceRuleAttributeRequest) SetRegionId(v string) *UpdateGroupDNSServiceRuleAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateGroupDNSServiceRuleAttributeRequest) SetServiceType(v string) *UpdateGroupDNSServiceRuleAttributeRequest {
	s.ServiceType = &v
	return s
}

func (s *UpdateGroupDNSServiceRuleAttributeRequest) SetSource(v string) *UpdateGroupDNSServiceRuleAttributeRequest {
	s.Source = &v
	return s
}

type UpdateGroupDNSServiceRuleAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateGroupDNSServiceRuleAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupDNSServiceRuleAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGroupDNSServiceRuleAttributeResponseBody) SetRequestId(v string) *UpdateGroupDNSServiceRuleAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateGroupDNSServiceRuleAttributeResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateGroupDNSServiceRuleAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateGroupDNSServiceRuleAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupDNSServiceRuleAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateGroupDNSServiceRuleAttributeResponse) SetHeaders(v map[string]*string) *UpdateGroupDNSServiceRuleAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateGroupDNSServiceRuleAttributeResponse) SetStatusCode(v int32) *UpdateGroupDNSServiceRuleAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGroupDNSServiceRuleAttributeResponse) SetBody(v *UpdateGroupDNSServiceRuleAttributeResponseBody) *UpdateGroupDNSServiceRuleAttributeResponse {
	s.Body = v
	return s
}

type UpdateGroupIpMappingRuleRequest struct {
	ClientToken              *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DestinationIp            *string `json:"DestinationIp,omitempty" xml:"DestinationIp,omitempty"`
	DryRun                   *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	GroupIpMappingRuleId     *string `json:"GroupIpMappingRuleId,omitempty" xml:"GroupIpMappingRuleId,omitempty"`
	IoTCloudConnectorGroupId *string `json:"IoTCloudConnectorGroupId,omitempty" xml:"IoTCloudConnectorGroupId,omitempty"`
	IpMappingRuleDescription *string `json:"IpMappingRuleDescription,omitempty" xml:"IpMappingRuleDescription,omitempty"`
	IpMappingRuleName        *string `json:"IpMappingRuleName,omitempty" xml:"IpMappingRuleName,omitempty"`
	MappingIp                *string `json:"MappingIp,omitempty" xml:"MappingIp,omitempty"`
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateGroupIpMappingRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupIpMappingRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupIpMappingRuleRequest) SetClientToken(v string) *UpdateGroupIpMappingRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateGroupIpMappingRuleRequest) SetDestinationIp(v string) *UpdateGroupIpMappingRuleRequest {
	s.DestinationIp = &v
	return s
}

func (s *UpdateGroupIpMappingRuleRequest) SetDryRun(v bool) *UpdateGroupIpMappingRuleRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateGroupIpMappingRuleRequest) SetGroupIpMappingRuleId(v string) *UpdateGroupIpMappingRuleRequest {
	s.GroupIpMappingRuleId = &v
	return s
}

func (s *UpdateGroupIpMappingRuleRequest) SetIoTCloudConnectorGroupId(v string) *UpdateGroupIpMappingRuleRequest {
	s.IoTCloudConnectorGroupId = &v
	return s
}

func (s *UpdateGroupIpMappingRuleRequest) SetIpMappingRuleDescription(v string) *UpdateGroupIpMappingRuleRequest {
	s.IpMappingRuleDescription = &v
	return s
}

func (s *UpdateGroupIpMappingRuleRequest) SetIpMappingRuleName(v string) *UpdateGroupIpMappingRuleRequest {
	s.IpMappingRuleName = &v
	return s
}

func (s *UpdateGroupIpMappingRuleRequest) SetMappingIp(v string) *UpdateGroupIpMappingRuleRequest {
	s.MappingIp = &v
	return s
}

func (s *UpdateGroupIpMappingRuleRequest) SetRegionId(v string) *UpdateGroupIpMappingRuleRequest {
	s.RegionId = &v
	return s
}

type UpdateGroupIpMappingRuleResponseBody struct {
	AuthorizationRuleId      *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	IoTCloudConnectorGroupId *string `json:"IoTCloudConnectorGroupId,omitempty" xml:"IoTCloudConnectorGroupId,omitempty"`
	RequestId                *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateGroupIpMappingRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupIpMappingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGroupIpMappingRuleResponseBody) SetAuthorizationRuleId(v string) *UpdateGroupIpMappingRuleResponseBody {
	s.AuthorizationRuleId = &v
	return s
}

func (s *UpdateGroupIpMappingRuleResponseBody) SetIoTCloudConnectorGroupId(v string) *UpdateGroupIpMappingRuleResponseBody {
	s.IoTCloudConnectorGroupId = &v
	return s
}

func (s *UpdateGroupIpMappingRuleResponseBody) SetRequestId(v string) *UpdateGroupIpMappingRuleResponseBody {
	s.RequestId = &v
	return s
}

type UpdateGroupIpMappingRuleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateGroupIpMappingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateGroupIpMappingRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupIpMappingRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateGroupIpMappingRuleResponse) SetHeaders(v map[string]*string) *UpdateGroupIpMappingRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateGroupIpMappingRuleResponse) SetStatusCode(v int32) *UpdateGroupIpMappingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGroupIpMappingRuleResponse) SetBody(v *UpdateGroupIpMappingRuleResponseBody) *UpdateGroupIpMappingRuleResponse {
	s.Body = v
	return s
}

type UpdateIoTCloudConnectorAttributeRequest struct {
	ClientToken                  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorDescription *string `json:"IoTCloudConnectorDescription,omitempty" xml:"IoTCloudConnectorDescription,omitempty"`
	IoTCloudConnectorId          *string `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	IoTCloudConnectorName        *string `json:"IoTCloudConnectorName,omitempty" xml:"IoTCloudConnectorName,omitempty"`
	Mode                         *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	RegionId                     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	WildcardDomainEnabled        *bool   `json:"WildcardDomainEnabled,omitempty" xml:"WildcardDomainEnabled,omitempty"`
}

func (s UpdateIoTCloudConnectorAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateIoTCloudConnectorAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateIoTCloudConnectorAttributeRequest) SetClientToken(v string) *UpdateIoTCloudConnectorAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateIoTCloudConnectorAttributeRequest) SetDryRun(v bool) *UpdateIoTCloudConnectorAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateIoTCloudConnectorAttributeRequest) SetIoTCloudConnectorDescription(v string) *UpdateIoTCloudConnectorAttributeRequest {
	s.IoTCloudConnectorDescription = &v
	return s
}

func (s *UpdateIoTCloudConnectorAttributeRequest) SetIoTCloudConnectorId(v string) *UpdateIoTCloudConnectorAttributeRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *UpdateIoTCloudConnectorAttributeRequest) SetIoTCloudConnectorName(v string) *UpdateIoTCloudConnectorAttributeRequest {
	s.IoTCloudConnectorName = &v
	return s
}

func (s *UpdateIoTCloudConnectorAttributeRequest) SetMode(v string) *UpdateIoTCloudConnectorAttributeRequest {
	s.Mode = &v
	return s
}

func (s *UpdateIoTCloudConnectorAttributeRequest) SetRegionId(v string) *UpdateIoTCloudConnectorAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateIoTCloudConnectorAttributeRequest) SetWildcardDomainEnabled(v bool) *UpdateIoTCloudConnectorAttributeRequest {
	s.WildcardDomainEnabled = &v
	return s
}

type UpdateIoTCloudConnectorAttributeResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s UpdateIoTCloudConnectorAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateIoTCloudConnectorAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIoTCloudConnectorAttributeResponseBody) SetRequestId(v string) *UpdateIoTCloudConnectorAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateIoTCloudConnectorAttributeResponseBody) SetResourceId(v string) *UpdateIoTCloudConnectorAttributeResponseBody {
	s.ResourceId = &v
	return s
}

type UpdateIoTCloudConnectorAttributeResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateIoTCloudConnectorAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateIoTCloudConnectorAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateIoTCloudConnectorAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateIoTCloudConnectorAttributeResponse) SetHeaders(v map[string]*string) *UpdateIoTCloudConnectorAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateIoTCloudConnectorAttributeResponse) SetStatusCode(v int32) *UpdateIoTCloudConnectorAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIoTCloudConnectorAttributeResponse) SetBody(v *UpdateIoTCloudConnectorAttributeResponseBody) *UpdateIoTCloudConnectorAttributeResponse {
	s.Body = v
	return s
}

type UpdateIoTCloudConnectorGroupAttributeRequest struct {
	ClientToken              *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description              *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DryRun                   *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorGroupId *string `json:"IoTCloudConnectorGroupId,omitempty" xml:"IoTCloudConnectorGroupId,omitempty"`
	Name                     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateIoTCloudConnectorGroupAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateIoTCloudConnectorGroupAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateIoTCloudConnectorGroupAttributeRequest) SetClientToken(v string) *UpdateIoTCloudConnectorGroupAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateIoTCloudConnectorGroupAttributeRequest) SetDescription(v string) *UpdateIoTCloudConnectorGroupAttributeRequest {
	s.Description = &v
	return s
}

func (s *UpdateIoTCloudConnectorGroupAttributeRequest) SetDryRun(v bool) *UpdateIoTCloudConnectorGroupAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateIoTCloudConnectorGroupAttributeRequest) SetIoTCloudConnectorGroupId(v string) *UpdateIoTCloudConnectorGroupAttributeRequest {
	s.IoTCloudConnectorGroupId = &v
	return s
}

func (s *UpdateIoTCloudConnectorGroupAttributeRequest) SetName(v string) *UpdateIoTCloudConnectorGroupAttributeRequest {
	s.Name = &v
	return s
}

func (s *UpdateIoTCloudConnectorGroupAttributeRequest) SetRegionId(v string) *UpdateIoTCloudConnectorGroupAttributeRequest {
	s.RegionId = &v
	return s
}

type UpdateIoTCloudConnectorGroupAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateIoTCloudConnectorGroupAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateIoTCloudConnectorGroupAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIoTCloudConnectorGroupAttributeResponseBody) SetRequestId(v string) *UpdateIoTCloudConnectorGroupAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateIoTCloudConnectorGroupAttributeResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateIoTCloudConnectorGroupAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateIoTCloudConnectorGroupAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateIoTCloudConnectorGroupAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateIoTCloudConnectorGroupAttributeResponse) SetHeaders(v map[string]*string) *UpdateIoTCloudConnectorGroupAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateIoTCloudConnectorGroupAttributeResponse) SetStatusCode(v int32) *UpdateIoTCloudConnectorGroupAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIoTCloudConnectorGroupAttributeResponse) SetBody(v *UpdateIoTCloudConnectorGroupAttributeResponseBody) *UpdateIoTCloudConnectorGroupAttributeResponse {
	s.Body = v
	return s
}

type UpdateIpMappingRuleRequest struct {
	ClientToken              *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DestinationIp            *string `json:"DestinationIp,omitempty" xml:"DestinationIp,omitempty"`
	DryRun                   *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorId      *string `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	IpMappingRuleDescription *string `json:"IpMappingRuleDescription,omitempty" xml:"IpMappingRuleDescription,omitempty"`
	IpMappingRuleId          *string `json:"IpMappingRuleId,omitempty" xml:"IpMappingRuleId,omitempty"`
	IpMappingRuleName        *string `json:"IpMappingRuleName,omitempty" xml:"IpMappingRuleName,omitempty"`
	MappingIp                *string `json:"MappingIp,omitempty" xml:"MappingIp,omitempty"`
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateIpMappingRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateIpMappingRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateIpMappingRuleRequest) SetClientToken(v string) *UpdateIpMappingRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateIpMappingRuleRequest) SetDestinationIp(v string) *UpdateIpMappingRuleRequest {
	s.DestinationIp = &v
	return s
}

func (s *UpdateIpMappingRuleRequest) SetDryRun(v bool) *UpdateIpMappingRuleRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateIpMappingRuleRequest) SetIoTCloudConnectorId(v string) *UpdateIpMappingRuleRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *UpdateIpMappingRuleRequest) SetIpMappingRuleDescription(v string) *UpdateIpMappingRuleRequest {
	s.IpMappingRuleDescription = &v
	return s
}

func (s *UpdateIpMappingRuleRequest) SetIpMappingRuleId(v string) *UpdateIpMappingRuleRequest {
	s.IpMappingRuleId = &v
	return s
}

func (s *UpdateIpMappingRuleRequest) SetIpMappingRuleName(v string) *UpdateIpMappingRuleRequest {
	s.IpMappingRuleName = &v
	return s
}

func (s *UpdateIpMappingRuleRequest) SetMappingIp(v string) *UpdateIpMappingRuleRequest {
	s.MappingIp = &v
	return s
}

func (s *UpdateIpMappingRuleRequest) SetRegionId(v string) *UpdateIpMappingRuleRequest {
	s.RegionId = &v
	return s
}

type UpdateIpMappingRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateIpMappingRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateIpMappingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIpMappingRuleResponseBody) SetRequestId(v string) *UpdateIpMappingRuleResponseBody {
	s.RequestId = &v
	return s
}

type UpdateIpMappingRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateIpMappingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateIpMappingRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateIpMappingRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateIpMappingRuleResponse) SetHeaders(v map[string]*string) *UpdateIpMappingRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateIpMappingRuleResponse) SetStatusCode(v int32) *UpdateIpMappingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIpMappingRuleResponse) SetBody(v *UpdateIpMappingRuleResponseBody) *UpdateIpMappingRuleResponse {
	s.Body = v
	return s
}

type UpdateServiceAttributeRequest struct {
	ClientToken         *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun              *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorId *string `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceDescription  *string `json:"ServiceDescription,omitempty" xml:"ServiceDescription,omitempty"`
	ServiceId           *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceName         *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s UpdateServiceAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceAttributeRequest) SetClientToken(v string) *UpdateServiceAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateServiceAttributeRequest) SetDryRun(v bool) *UpdateServiceAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateServiceAttributeRequest) SetIoTCloudConnectorId(v string) *UpdateServiceAttributeRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *UpdateServiceAttributeRequest) SetRegionId(v string) *UpdateServiceAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateServiceAttributeRequest) SetServiceDescription(v string) *UpdateServiceAttributeRequest {
	s.ServiceDescription = &v
	return s
}

func (s *UpdateServiceAttributeRequest) SetServiceId(v string) *UpdateServiceAttributeRequest {
	s.ServiceId = &v
	return s
}

func (s *UpdateServiceAttributeRequest) SetServiceName(v string) *UpdateServiceAttributeRequest {
	s.ServiceName = &v
	return s
}

type UpdateServiceAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServiceAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceAttributeResponseBody) SetRequestId(v string) *UpdateServiceAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateServiceAttributeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateServiceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateServiceAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceAttributeResponse) SetHeaders(v map[string]*string) *UpdateServiceAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceAttributeResponse) SetStatusCode(v int32) *UpdateServiceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceAttributeResponse) SetBody(v *UpdateServiceAttributeResponseBody) *UpdateServiceAttributeResponse {
	s.Body = v
	return s
}

type UpdateServiceEntryAttributeRequest struct {
	ClientToken             *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                  *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorId     *string `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	RegionId                *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceEntryDescription *string `json:"ServiceEntryDescription,omitempty" xml:"ServiceEntryDescription,omitempty"`
	ServiceEntryId          *string `json:"ServiceEntryId,omitempty" xml:"ServiceEntryId,omitempty"`
	ServiceEntryName        *string `json:"ServiceEntryName,omitempty" xml:"ServiceEntryName,omitempty"`
	ServiceId               *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s UpdateServiceEntryAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceEntryAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceEntryAttributeRequest) SetClientToken(v string) *UpdateServiceEntryAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateServiceEntryAttributeRequest) SetDryRun(v bool) *UpdateServiceEntryAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateServiceEntryAttributeRequest) SetIoTCloudConnectorId(v string) *UpdateServiceEntryAttributeRequest {
	s.IoTCloudConnectorId = &v
	return s
}

func (s *UpdateServiceEntryAttributeRequest) SetRegionId(v string) *UpdateServiceEntryAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateServiceEntryAttributeRequest) SetServiceEntryDescription(v string) *UpdateServiceEntryAttributeRequest {
	s.ServiceEntryDescription = &v
	return s
}

func (s *UpdateServiceEntryAttributeRequest) SetServiceEntryId(v string) *UpdateServiceEntryAttributeRequest {
	s.ServiceEntryId = &v
	return s
}

func (s *UpdateServiceEntryAttributeRequest) SetServiceEntryName(v string) *UpdateServiceEntryAttributeRequest {
	s.ServiceEntryName = &v
	return s
}

func (s *UpdateServiceEntryAttributeRequest) SetServiceId(v string) *UpdateServiceEntryAttributeRequest {
	s.ServiceId = &v
	return s
}

type UpdateServiceEntryAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServiceEntryAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceEntryAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceEntryAttributeResponseBody) SetRequestId(v string) *UpdateServiceEntryAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateServiceEntryAttributeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateServiceEntryAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateServiceEntryAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceEntryAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceEntryAttributeResponse) SetHeaders(v map[string]*string) *UpdateServiceEntryAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceEntryAttributeResponse) SetStatusCode(v int32) *UpdateServiceEntryAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceEntryAttributeResponse) SetBody(v *UpdateServiceEntryAttributeResponseBody) *UpdateServiceEntryAttributeResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("iotcc"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddCidrToConnectionPoolWithOptions(request *AddCidrToConnectionPoolRequest, runtime *util.RuntimeOptions) (_result *AddCidrToConnectionPoolResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cidrs)) {
		query["Cidrs"] = request.Cidrs
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionPoolId)) {
		query["ConnectionPoolId"] = request.ConnectionPoolId
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddCidrToConnectionPool"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddCidrToConnectionPoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddCidrToConnectionPool(request *AddCidrToConnectionPoolRequest) (_result *AddCidrToConnectionPoolResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddCidrToConnectionPoolResponse{}
	_body, _err := client.AddCidrToConnectionPoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddIoTCloudConnectorToGroupWithOptions(request *AddIoTCloudConnectorToGroupRequest, runtime *util.RuntimeOptions) (_result *AddIoTCloudConnectorToGroupResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorGroupId)) {
		query["IoTCloudConnectorGroupId"] = request.IoTCloudConnectorGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddIoTCloudConnectorToGroup"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddIoTCloudConnectorToGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddIoTCloudConnectorToGroup(request *AddIoTCloudConnectorToGroupRequest) (_result *AddIoTCloudConnectorToGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddIoTCloudConnectorToGroupResponse{}
	_body, _err := client.AddIoTCloudConnectorToGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AssociateIpWithConnectionPoolWithOptions(request *AssociateIpWithConnectionPoolRequest, runtime *util.RuntimeOptions) (_result *AssociateIpWithConnectionPoolResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionPoolId)) {
		query["ConnectionPoolId"] = request.ConnectionPoolId
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.Ips)) {
		query["Ips"] = request.Ips
	}

	if !tea.BoolValue(util.IsUnset(request.IpsFilePath)) {
		query["IpsFilePath"] = request.IpsFilePath
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AssociateIpWithConnectionPool"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AssociateIpWithConnectionPoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AssociateIpWithConnectionPool(request *AssociateIpWithConnectionPoolRequest) (_result *AssociateIpWithConnectionPoolResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssociateIpWithConnectionPoolResponse{}
	_body, _err := client.AssociateIpWithConnectionPoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AssociateVSwitchWithIoTCloudConnectorWithOptions(request *AssociateVSwitchWithIoTCloudConnectorRequest, runtime *util.RuntimeOptions) (_result *AssociateVSwitchWithIoTCloudConnectorResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchList)) {
		query["VSwitchList"] = request.VSwitchList
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AssociateVSwitchWithIoTCloudConnector"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AssociateVSwitchWithIoTCloudConnectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AssociateVSwitchWithIoTCloudConnector(request *AssociateVSwitchWithIoTCloudConnectorRequest) (_result *AssociateVSwitchWithIoTCloudConnectorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssociateVSwitchWithIoTCloudConnectorResponse{}
	_body, _err := client.AssociateVSwitchWithIoTCloudConnectorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfirmIoTCloudConnectorWithOptions(request *ConfirmIoTCloudConnectorRequest, runtime *util.RuntimeOptions) (_result *ConfirmIoTCloudConnectorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ConfirmStatus)) {
		query["ConfirmStatus"] = request.ConfirmStatus
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfirmIoTCloudConnector"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConfirmIoTCloudConnectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfirmIoTCloudConnector(request *ConfirmIoTCloudConnectorRequest) (_result *ConfirmIoTCloudConnectorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfirmIoTCloudConnectorResponse{}
	_body, _err := client.ConfirmIoTCloudConnectorWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.AuthorizationRuleDescription)) {
		query["AuthorizationRuleDescription"] = request.AuthorizationRuleDescription
	}

	if !tea.BoolValue(util.IsUnset(request.AuthorizationRuleName)) {
		query["AuthorizationRuleName"] = request.AuthorizationRuleName
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
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

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.Policy)) {
		query["Policy"] = request.Policy
	}

	if !tea.BoolValue(util.IsUnset(request.Protocol)) {
		query["Protocol"] = request.Protocol
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceCidrs)) {
		query["SourceCidrs"] = request.SourceCidrs
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAuthorizationRule"),
		Version:     tea.String("2021-05-13"),
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

func (client *Client) CreateAuthorizationRulesWithOptions(request *CreateAuthorizationRulesRequest, runtime *util.RuntimeOptions) (_result *CreateAuthorizationRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthorizationRules)) {
		query["AuthorizationRules"] = request.AuthorizationRules
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAuthorizationRules"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAuthorizationRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAuthorizationRules(request *CreateAuthorizationRulesRequest) (_result *CreateAuthorizationRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAuthorizationRulesResponse{}
	_body, _err := client.CreateAuthorizationRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateConnectionPoolWithOptions(request *CreateConnectionPoolRequest, runtime *util.RuntimeOptions) (_result *CreateConnectionPoolResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cidrs)) {
		query["Cidrs"] = request.Cidrs
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionPoolDescription)) {
		query["ConnectionPoolDescription"] = request.ConnectionPoolDescription
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionPoolName)) {
		query["ConnectionPoolName"] = request.ConnectionPoolName
	}

	if !tea.BoolValue(util.IsUnset(request.Count)) {
		query["Count"] = request.Count
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateConnectionPool"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateConnectionPoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateConnectionPool(request *CreateConnectionPoolRequest) (_result *CreateConnectionPoolResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateConnectionPoolResponse{}
	_body, _err := client.CreateConnectionPoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDNSServiceRuleWithOptions(request *CreateDNSServiceRuleRequest, runtime *util.RuntimeOptions) (_result *CreateDNSServiceRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DNSServiceRuleDescription)) {
		query["DNSServiceRuleDescription"] = request.DNSServiceRuleDescription
	}

	if !tea.BoolValue(util.IsUnset(request.DNSServiceRuleName)) {
		query["DNSServiceRuleName"] = request.DNSServiceRuleName
	}

	if !tea.BoolValue(util.IsUnset(request.Destination)) {
		query["Destination"] = request.Destination
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceType)) {
		query["ServiceType"] = request.ServiceType
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDNSServiceRule"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDNSServiceRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDNSServiceRule(request *CreateDNSServiceRuleRequest) (_result *CreateDNSServiceRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDNSServiceRuleResponse{}
	_body, _err := client.CreateDNSServiceRuleWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.AuthorizationRuleDescription)) {
		query["AuthorizationRuleDescription"] = request.AuthorizationRuleDescription
	}

	if !tea.BoolValue(util.IsUnset(request.AuthorizationRuleName)) {
		query["AuthorizationRuleName"] = request.AuthorizationRuleName
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
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

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorGroupId)) {
		query["IoTCloudConnectorGroupId"] = request.IoTCloudConnectorGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Policy)) {
		query["Policy"] = request.Policy
	}

	if !tea.BoolValue(util.IsUnset(request.Protocol)) {
		query["Protocol"] = request.Protocol
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceCidrs)) {
		query["SourceCidrs"] = request.SourceCidrs
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGroupAuthorizationRule"),
		Version:     tea.String("2021-05-13"),
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

func (client *Client) CreateGroupDNSServiceRuleWithOptions(request *CreateGroupDNSServiceRuleRequest, runtime *util.RuntimeOptions) (_result *CreateGroupDNSServiceRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DNSServiceRuleDescription)) {
		query["DNSServiceRuleDescription"] = request.DNSServiceRuleDescription
	}

	if !tea.BoolValue(util.IsUnset(request.DNSServiceRuleName)) {
		query["DNSServiceRuleName"] = request.DNSServiceRuleName
	}

	if !tea.BoolValue(util.IsUnset(request.Destination)) {
		query["Destination"] = request.Destination
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorGroupId)) {
		query["IoTCloudConnectorGroupId"] = request.IoTCloudConnectorGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceType)) {
		query["ServiceType"] = request.ServiceType
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGroupDNSServiceRule"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGroupDNSServiceRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGroupDNSServiceRule(request *CreateGroupDNSServiceRuleRequest) (_result *CreateGroupDNSServiceRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGroupDNSServiceRuleResponse{}
	_body, _err := client.CreateGroupDNSServiceRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGroupIpMappingRuleWithOptions(request *CreateGroupIpMappingRuleRequest, runtime *util.RuntimeOptions) (_result *CreateGroupIpMappingRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationIp)) {
		query["DestinationIp"] = request.DestinationIp
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorGroupId)) {
		query["IoTCloudConnectorGroupId"] = request.IoTCloudConnectorGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.IpMappingRuleDescription)) {
		query["IpMappingRuleDescription"] = request.IpMappingRuleDescription
	}

	if !tea.BoolValue(util.IsUnset(request.IpMappingRuleName)) {
		query["IpMappingRuleName"] = request.IpMappingRuleName
	}

	if !tea.BoolValue(util.IsUnset(request.MappingIp)) {
		query["MappingIp"] = request.MappingIp
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGroupIpMappingRule"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGroupIpMappingRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGroupIpMappingRule(request *CreateGroupIpMappingRuleRequest) (_result *CreateGroupIpMappingRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGroupIpMappingRuleResponse{}
	_body, _err := client.CreateGroupIpMappingRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateIoTCloudConnectorWithOptions(request *CreateIoTCloudConnectorRequest, runtime *util.RuntimeOptions) (_result *CreateIoTCloudConnectorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.APN)) {
		query["APN"] = request.APN
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ISP)) {
		query["ISP"] = request.ISP
	}

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorDescription)) {
		query["IoTCloudConnectorDescription"] = request.IoTCloudConnectorDescription
	}

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorName)) {
		query["IoTCloudConnectorName"] = request.IoTCloudConnectorName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceUid)) {
		query["ResourceUid"] = request.ResourceUid
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.WildcardDomainEnabled)) {
		query["WildcardDomainEnabled"] = request.WildcardDomainEnabled
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateIoTCloudConnector"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateIoTCloudConnectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateIoTCloudConnector(request *CreateIoTCloudConnectorRequest) (_result *CreateIoTCloudConnectorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateIoTCloudConnectorResponse{}
	_body, _err := client.CreateIoTCloudConnectorWithOptions(request, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateIoTCloudConnectorBackhaulRoute"),
		Version:     tea.String("2021-05-13"),
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

func (client *Client) CreateIoTCloudConnectorGroupWithOptions(request *CreateIoTCloudConnectorGroupRequest, runtime *util.RuntimeOptions) (_result *CreateIoTCloudConnectorGroupResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateIoTCloudConnectorGroup"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateIoTCloudConnectorGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateIoTCloudConnectorGroup(request *CreateIoTCloudConnectorGroupRequest) (_result *CreateIoTCloudConnectorGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateIoTCloudConnectorGroupResponse{}
	_body, _err := client.CreateIoTCloudConnectorGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateIpMappingRuleWithOptions(request *CreateIpMappingRuleRequest, runtime *util.RuntimeOptions) (_result *CreateIpMappingRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationIp)) {
		query["DestinationIp"] = request.DestinationIp
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.IpMappingRuleDescription)) {
		query["IpMappingRuleDescription"] = request.IpMappingRuleDescription
	}

	if !tea.BoolValue(util.IsUnset(request.IpMappingRuleName)) {
		query["IpMappingRuleName"] = request.IpMappingRuleName
	}

	if !tea.BoolValue(util.IsUnset(request.MappingIp)) {
		query["MappingIp"] = request.MappingIp
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateIpMappingRule"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateIpMappingRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateIpMappingRule(request *CreateIpMappingRuleRequest) (_result *CreateIpMappingRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateIpMappingRuleResponse{}
	_body, _err := client.CreateIpMappingRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateServiceWithOptions(request *CreateServiceRequest, runtime *util.RuntimeOptions) (_result *CreateServiceResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceDescription)) {
		query["ServiceDescription"] = request.ServiceDescription
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateService"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateService(request *CreateServiceRequest) (_result *CreateServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateServiceResponse{}
	_body, _err := client.CreateServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateServiceEntryWithOptions(request *CreateServiceEntryRequest, runtime *util.RuntimeOptions) (_result *CreateServiceEntryResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceEntryDescription)) {
		query["ServiceEntryDescription"] = request.ServiceEntryDescription
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceEntryName)) {
		query["ServiceEntryName"] = request.ServiceEntryName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.Target)) {
		query["Target"] = request.Target
	}

	if !tea.BoolValue(util.IsUnset(request.TargetType)) {
		query["TargetType"] = request.TargetType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateServiceEntry"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServiceEntryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateServiceEntry(request *CreateServiceEntryRequest) (_result *CreateServiceEntryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateServiceEntryResponse{}
	_body, _err := client.CreateServiceEntryWithOptions(request, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAuthorizationRule"),
		Version:     tea.String("2021-05-13"),
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

func (client *Client) DeleteAuthorizationRulesWithOptions(request *DeleteAuthorizationRulesRequest, runtime *util.RuntimeOptions) (_result *DeleteAuthorizationRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthorizationRuleIds)) {
		query["AuthorizationRuleIds"] = request.AuthorizationRuleIds
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAuthorizationRules"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAuthorizationRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAuthorizationRules(request *DeleteAuthorizationRulesRequest) (_result *DeleteAuthorizationRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAuthorizationRulesResponse{}
	_body, _err := client.DeleteAuthorizationRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteConnectionPoolWithOptions(request *DeleteConnectionPoolRequest, runtime *util.RuntimeOptions) (_result *DeleteConnectionPoolResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionPoolId)) {
		query["ConnectionPoolId"] = request.ConnectionPoolId
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteConnectionPool"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteConnectionPoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteConnectionPool(request *DeleteConnectionPoolRequest) (_result *DeleteConnectionPoolResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteConnectionPoolResponse{}
	_body, _err := client.DeleteConnectionPoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDNSServiceRuleWithOptions(request *DeleteDNSServiceRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteDNSServiceRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DNSServiceRuleId)) {
		query["DNSServiceRuleId"] = request.DNSServiceRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDNSServiceRule"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDNSServiceRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDNSServiceRule(request *DeleteDNSServiceRuleRequest) (_result *DeleteDNSServiceRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDNSServiceRuleResponse{}
	_body, _err := client.DeleteDNSServiceRuleWithOptions(request, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorGroupId)) {
		query["IoTCloudConnectorGroupId"] = request.IoTCloudConnectorGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGroupAuthorizationRule"),
		Version:     tea.String("2021-05-13"),
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

func (client *Client) DeleteGroupDNSServiceRuleWithOptions(request *DeleteGroupDNSServiceRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteGroupDNSServiceRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DNSServiceRuleId)) {
		query["DNSServiceRuleId"] = request.DNSServiceRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorGroupId)) {
		query["IoTCloudConnectorGroupId"] = request.IoTCloudConnectorGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGroupDNSServiceRule"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGroupDNSServiceRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGroupDNSServiceRule(request *DeleteGroupDNSServiceRuleRequest) (_result *DeleteGroupDNSServiceRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGroupDNSServiceRuleResponse{}
	_body, _err := client.DeleteGroupDNSServiceRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteGroupIpMappingRuleWithOptions(request *DeleteGroupIpMappingRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteGroupIpMappingRuleResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.GroupIpMappingRuleId)) {
		query["GroupIpMappingRuleId"] = request.GroupIpMappingRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorGroupId)) {
		query["IoTCloudConnectorGroupId"] = request.IoTCloudConnectorGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGroupIpMappingRule"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGroupIpMappingRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGroupIpMappingRule(request *DeleteGroupIpMappingRuleRequest) (_result *DeleteGroupIpMappingRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGroupIpMappingRuleResponse{}
	_body, _err := client.DeleteGroupIpMappingRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteIoTCloudConnectorWithOptions(request *DeleteIoTCloudConnectorRequest, runtime *util.RuntimeOptions) (_result *DeleteIoTCloudConnectorResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteIoTCloudConnector"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteIoTCloudConnectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteIoTCloudConnector(request *DeleteIoTCloudConnectorRequest) (_result *DeleteIoTCloudConnectorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteIoTCloudConnectorResponse{}
	_body, _err := client.DeleteIoTCloudConnectorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteIoTCloudConnectorGroupWithOptions(request *DeleteIoTCloudConnectorGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteIoTCloudConnectorGroupResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorGroupId)) {
		query["IoTCloudConnectorGroupId"] = request.IoTCloudConnectorGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteIoTCloudConnectorGroup"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteIoTCloudConnectorGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteIoTCloudConnectorGroup(request *DeleteIoTCloudConnectorGroupRequest) (_result *DeleteIoTCloudConnectorGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteIoTCloudConnectorGroupResponse{}
	_body, _err := client.DeleteIoTCloudConnectorGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteIoTCloudConnetorBackhaulRouteWithOptions(request *DeleteIoTCloudConnetorBackhaulRouteRequest, runtime *util.RuntimeOptions) (_result *DeleteIoTCloudConnetorBackhaulRouteResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteIoTCloudConnetorBackhaulRoute"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteIoTCloudConnetorBackhaulRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteIoTCloudConnetorBackhaulRoute(request *DeleteIoTCloudConnetorBackhaulRouteRequest) (_result *DeleteIoTCloudConnetorBackhaulRouteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteIoTCloudConnetorBackhaulRouteResponse{}
	_body, _err := client.DeleteIoTCloudConnetorBackhaulRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteIpMappingRuleWithOptions(request *DeleteIpMappingRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteIpMappingRuleResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.IpMappingRuleId)) {
		query["IpMappingRuleId"] = request.IpMappingRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteIpMappingRule"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteIpMappingRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteIpMappingRule(request *DeleteIpMappingRuleRequest) (_result *DeleteIpMappingRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteIpMappingRuleResponse{}
	_body, _err := client.DeleteIpMappingRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteServiceWithOptions(request *DeleteServiceRequest, runtime *util.RuntimeOptions) (_result *DeleteServiceResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteService"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteService(request *DeleteServiceRequest) (_result *DeleteServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteServiceResponse{}
	_body, _err := client.DeleteServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteServiceEntryWithOptions(request *DeleteServiceEntryRequest, runtime *util.RuntimeOptions) (_result *DeleteServiceEntryResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceEntryId)) {
		query["ServiceEntryId"] = request.ServiceEntryId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteServiceEntry"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteServiceEntryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteServiceEntry(request *DeleteServiceEntryRequest) (_result *DeleteServiceEntryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteServiceEntryResponse{}
	_body, _err := client.DeleteServiceEntryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableIoTCloudConnectorAccessLogWithOptions(request *DisableIoTCloudConnectorAccessLogRequest, runtime *util.RuntimeOptions) (_result *DisableIoTCloudConnectorAccessLogResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableIoTCloudConnectorAccessLog"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableIoTCloudConnectorAccessLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableIoTCloudConnectorAccessLog(request *DisableIoTCloudConnectorAccessLogRequest) (_result *DisableIoTCloudConnectorAccessLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableIoTCloudConnectorAccessLogResponse{}
	_body, _err := client.DisableIoTCloudConnectorAccessLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DissociateIpFromConnectionPoolWithOptions(request *DissociateIpFromConnectionPoolRequest, runtime *util.RuntimeOptions) (_result *DissociateIpFromConnectionPoolResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionPoolId)) {
		query["ConnectionPoolId"] = request.ConnectionPoolId
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.Ips)) {
		query["Ips"] = request.Ips
	}

	if !tea.BoolValue(util.IsUnset(request.IpsFilePath)) {
		query["IpsFilePath"] = request.IpsFilePath
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DissociateIpFromConnectionPool"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DissociateIpFromConnectionPoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DissociateIpFromConnectionPool(request *DissociateIpFromConnectionPoolRequest) (_result *DissociateIpFromConnectionPoolResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DissociateIpFromConnectionPoolResponse{}
	_body, _err := client.DissociateIpFromConnectionPoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DissociateVSwitchFromIoTCloudConnectorWithOptions(request *DissociateVSwitchFromIoTCloudConnectorRequest, runtime *util.RuntimeOptions) (_result *DissociateVSwitchFromIoTCloudConnectorResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DissociateVSwitchFromIoTCloudConnector"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DissociateVSwitchFromIoTCloudConnectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DissociateVSwitchFromIoTCloudConnector(request *DissociateVSwitchFromIoTCloudConnectorRequest) (_result *DissociateVSwitchFromIoTCloudConnectorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DissociateVSwitchFromIoTCloudConnectorResponse{}
	_body, _err := client.DissociateVSwitchFromIoTCloudConnectorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableIoTCloudConnectorAccessLogWithOptions(request *EnableIoTCloudConnectorAccessLogRequest, runtime *util.RuntimeOptions) (_result *EnableIoTCloudConnectorAccessLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessLogSlsLogStore)) {
		query["AccessLogSlsLogStore"] = request.AccessLogSlsLogStore
	}

	if !tea.BoolValue(util.IsUnset(request.AccessLogSlsProject)) {
		query["AccessLogSlsProject"] = request.AccessLogSlsProject
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableIoTCloudConnectorAccessLog"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableIoTCloudConnectorAccessLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableIoTCloudConnectorAccessLog(request *EnableIoTCloudConnectorAccessLogRequest) (_result *EnableIoTCloudConnectorAccessLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableIoTCloudConnectorAccessLogResponse{}
	_body, _err := client.EnableIoTCloudConnectorAccessLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetConnectionPoolIpOperationResultWithOptions(request *GetConnectionPoolIpOperationResultRequest, runtime *util.RuntimeOptions) (_result *GetConnectionPoolIpOperationResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnectionPoolId)) {
		query["ConnectionPoolId"] = request.ConnectionPoolId
	}

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.QueryRequestId)) {
		query["QueryRequestId"] = request.QueryRequestId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetConnectionPoolIpOperationResult"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetConnectionPoolIpOperationResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetConnectionPoolIpOperationResult(request *GetConnectionPoolIpOperationResultRequest) (_result *GetConnectionPoolIpOperationResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetConnectionPoolIpOperationResultResponse{}
	_body, _err := client.GetConnectionPoolIpOperationResultWithOptions(request, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDiagnoseResultForSingleCard"),
		Version:     tea.String("2021-05-13"),
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

func (client *Client) GetIoTCloudConnectorAccessLogWithOptions(request *GetIoTCloudConnectorAccessLogRequest, runtime *util.RuntimeOptions) (_result *GetIoTCloudConnectorAccessLogResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetIoTCloudConnectorAccessLog"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetIoTCloudConnectorAccessLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetIoTCloudConnectorAccessLog(request *GetIoTCloudConnectorAccessLogRequest) (_result *GetIoTCloudConnectorAccessLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetIoTCloudConnectorAccessLogResponse{}
	_body, _err := client.GetIoTCloudConnectorAccessLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetStsInfoAndOssPathWithOptions(request *GetStsInfoAndOssPathRequest, runtime *util.RuntimeOptions) (_result *GetStsInfoAndOssPathResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionPoolId)) {
		query["ConnectionPoolId"] = request.ConnectionPoolId
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["FileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetStsInfoAndOssPath"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetStsInfoAndOssPathResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetStsInfoAndOssPath(request *GetStsInfoAndOssPathRequest) (_result *GetStsInfoAndOssPathResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetStsInfoAndOssPathResponse{}
	_body, _err := client.GetStsInfoAndOssPathWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GrantVirtualBorderRouterWithOptions(request *GrantVirtualBorderRouterRequest, runtime *util.RuntimeOptions) (_result *GrantVirtualBorderRouterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.VirtualBorderRouterId)) {
		query["VirtualBorderRouterId"] = request.VirtualBorderRouterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GrantVirtualBorderRouter"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GrantVirtualBorderRouterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GrantVirtualBorderRouter(request *GrantVirtualBorderRouterRequest) (_result *GrantVirtualBorderRouterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GrantVirtualBorderRouterResponse{}
	_body, _err := client.GrantVirtualBorderRouterWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.APN)) {
		query["APN"] = request.APN
	}

	if !tea.BoolValue(util.IsUnset(request.ISP)) {
		query["ISP"] = request.ISP
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAPNs"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthorizationRuleIds)) {
		query["AuthorizationRuleIds"] = request.AuthorizationRuleIds
	}

	if !tea.BoolValue(util.IsUnset(request.AuthorizationRuleName)) {
		query["AuthorizationRuleName"] = request.AuthorizationRuleName
	}

	if !tea.BoolValue(util.IsUnset(request.AuthorizationRuleStatus)) {
		query["AuthorizationRuleStatus"] = request.AuthorizationRuleStatus
	}

	if !tea.BoolValue(util.IsUnset(request.AuthorizationRuleType)) {
		query["AuthorizationRuleType"] = request.AuthorizationRuleType
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

	if !tea.BoolValue(util.IsUnset(request.FuzzyAuthorizationRuleName)) {
		query["FuzzyAuthorizationRuleName"] = request.FuzzyAuthorizationRuleName
	}

	if !tea.BoolValue(util.IsUnset(request.FuzzyDestination)) {
		query["FuzzyDestination"] = request.FuzzyDestination
	}

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Policy)) {
		query["Policy"] = request.Policy
	}

	if !tea.BoolValue(util.IsUnset(request.Protocol)) {
		query["Protocol"] = request.Protocol
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAuthorizationRules"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
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

func (client *Client) ListConnectionPoolAllIpsWithOptions(request *ListConnectionPoolAllIpsRequest, runtime *util.RuntimeOptions) (_result *ListConnectionPoolAllIpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnectionPoolId)) {
		query["ConnectionPoolId"] = request.ConnectionPoolId
	}

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListConnectionPoolAllIps"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListConnectionPoolAllIpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListConnectionPoolAllIps(request *ListConnectionPoolAllIpsRequest) (_result *ListConnectionPoolAllIpsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListConnectionPoolAllIpsResponse{}
	_body, _err := client.ListConnectionPoolAllIpsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListConnectionPoolIpsWithOptions(request *ListConnectionPoolIpsRequest, runtime *util.RuntimeOptions) (_result *ListConnectionPoolIpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnectionPoolId)) {
		query["ConnectionPoolId"] = request.ConnectionPoolId
	}

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListConnectionPoolIps"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListConnectionPoolIpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListConnectionPoolIps(request *ListConnectionPoolIpsRequest) (_result *ListConnectionPoolIpsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListConnectionPoolIpsResponse{}
	_body, _err := client.ListConnectionPoolIpsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListConnectionPoolsWithOptions(request *ListConnectionPoolsRequest, runtime *util.RuntimeOptions) (_result *ListConnectionPoolsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnectionPoolIds)) {
		query["ConnectionPoolIds"] = request.ConnectionPoolIds
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionPoolName)) {
		query["ConnectionPoolName"] = request.ConnectionPoolName
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionPoolStatus)) {
		query["ConnectionPoolStatus"] = request.ConnectionPoolStatus
	}

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListConnectionPools"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListConnectionPoolsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListConnectionPools(request *ListConnectionPoolsRequest) (_result *ListConnectionPoolsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListConnectionPoolsResponse{}
	_body, _err := client.ListConnectionPoolsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDNSServiceRulesWithOptions(request *ListDNSServiceRulesRequest, runtime *util.RuntimeOptions) (_result *ListDNSServiceRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DNSServiceRuleIds)) {
		query["DNSServiceRuleIds"] = request.DNSServiceRuleIds
	}

	if !tea.BoolValue(util.IsUnset(request.DNSServiceRuleName)) {
		query["DNSServiceRuleName"] = request.DNSServiceRuleName
	}

	if !tea.BoolValue(util.IsUnset(request.DNSServiceRuleStatus)) {
		query["DNSServiceRuleStatus"] = request.DNSServiceRuleStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Destination)) {
		query["Destination"] = request.Destination
	}

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceType)) {
		query["ServiceType"] = request.ServiceType
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDNSServiceRules"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDNSServiceRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDNSServiceRules(request *ListDNSServiceRulesRequest) (_result *ListDNSServiceRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDNSServiceRulesResponse{}
	_body, _err := client.ListDNSServiceRulesWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDiagnoseInfoForSingleCard"),
		Version:     tea.String("2021-05-13"),
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthorizationRuleIds)) {
		query["AuthorizationRuleIds"] = request.AuthorizationRuleIds
	}

	if !tea.BoolValue(util.IsUnset(request.AuthorizationRuleName)) {
		query["AuthorizationRuleName"] = request.AuthorizationRuleName
	}

	if !tea.BoolValue(util.IsUnset(request.AuthorizationRuleStatus)) {
		query["AuthorizationRuleStatus"] = request.AuthorizationRuleStatus
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

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorGroupId)) {
		query["IoTCloudConnectorGroupId"] = request.IoTCloudConnectorGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Policy)) {
		query["Policy"] = request.Policy
	}

	if !tea.BoolValue(util.IsUnset(request.Protocol)) {
		query["Protocol"] = request.Protocol
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGroupAuthorizationRules"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
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

func (client *Client) ListGroupDNSServiceRulesWithOptions(request *ListGroupDNSServiceRulesRequest, runtime *util.RuntimeOptions) (_result *ListGroupDNSServiceRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DNSServiceRuleIds)) {
		query["DNSServiceRuleIds"] = request.DNSServiceRuleIds
	}

	if !tea.BoolValue(util.IsUnset(request.DNSServiceRuleName)) {
		query["DNSServiceRuleName"] = request.DNSServiceRuleName
	}

	if !tea.BoolValue(util.IsUnset(request.DNSServiceRuleStatus)) {
		query["DNSServiceRuleStatus"] = request.DNSServiceRuleStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Destination)) {
		query["Destination"] = request.Destination
	}

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorGroupId)) {
		query["IoTCloudConnectorGroupId"] = request.IoTCloudConnectorGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceType)) {
		query["ServiceType"] = request.ServiceType
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGroupDNSServiceRules"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGroupDNSServiceRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListGroupDNSServiceRules(request *ListGroupDNSServiceRulesRequest) (_result *ListGroupDNSServiceRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListGroupDNSServiceRulesResponse{}
	_body, _err := client.ListGroupDNSServiceRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListGroupIpMappingRulesWithOptions(request *ListGroupIpMappingRulesRequest, runtime *util.RuntimeOptions) (_result *ListGroupIpMappingRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DestinationIps)) {
		query["DestinationIps"] = request.DestinationIps
	}

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorGroupId)) {
		query["IoTCloudConnectorGroupId"] = request.IoTCloudConnectorGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.IpMappingRuleIds)) {
		query["IpMappingRuleIds"] = request.IpMappingRuleIds
	}

	if !tea.BoolValue(util.IsUnset(request.IpMappingRuleNames)) {
		query["IpMappingRuleNames"] = request.IpMappingRuleNames
	}

	if !tea.BoolValue(util.IsUnset(request.IpMappingRuleStatuses)) {
		query["IpMappingRuleStatuses"] = request.IpMappingRuleStatuses
	}

	if !tea.BoolValue(util.IsUnset(request.MappingIps)) {
		query["MappingIps"] = request.MappingIps
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGroupIpMappingRules"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGroupIpMappingRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListGroupIpMappingRules(request *ListGroupIpMappingRulesRequest) (_result *ListGroupIpMappingRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListGroupIpMappingRulesResponse{}
	_body, _err := client.ListGroupIpMappingRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListIoTCloudConnectorAvailableZonesWithOptions(request *ListIoTCloudConnectorAvailableZonesRequest, runtime *util.RuntimeOptions) (_result *ListIoTCloudConnectorAvailableZonesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListIoTCloudConnectorAvailableZones"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIoTCloudConnectorAvailableZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListIoTCloudConnectorAvailableZones(request *ListIoTCloudConnectorAvailableZonesRequest) (_result *ListIoTCloudConnectorAvailableZonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListIoTCloudConnectorAvailableZonesResponse{}
	_body, _err := client.ListIoTCloudConnectorAvailableZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListIoTCloudConnectorEIPsWithOptions(request *ListIoTCloudConnectorEIPsRequest, runtime *util.RuntimeOptions) (_result *ListIoTCloudConnectorEIPsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListIoTCloudConnectorEIPs"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIoTCloudConnectorEIPsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListIoTCloudConnectorEIPs(request *ListIoTCloudConnectorEIPsRequest) (_result *ListIoTCloudConnectorEIPsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListIoTCloudConnectorEIPsResponse{}
	_body, _err := client.ListIoTCloudConnectorEIPsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListIoTCloudConnectorGroupsWithOptions(request *ListIoTCloudConnectorGroupsRequest, runtime *util.RuntimeOptions) (_result *ListIoTCloudConnectorGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorGroupIds)) {
		query["IoTCloudConnectorGroupIds"] = request.IoTCloudConnectorGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorGroupName)) {
		query["IoTCloudConnectorGroupName"] = request.IoTCloudConnectorGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorGroupStatus)) {
		query["IoTCloudConnectorGroupStatus"] = request.IoTCloudConnectorGroupStatus
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListIoTCloudConnectorGroups"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIoTCloudConnectorGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListIoTCloudConnectorGroups(request *ListIoTCloudConnectorGroupsRequest) (_result *ListIoTCloudConnectorGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListIoTCloudConnectorGroupsResponse{}
	_body, _err := client.ListIoTCloudConnectorGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListIoTCloudConnectorsWithOptions(request *ListIoTCloudConnectorsRequest, runtime *util.RuntimeOptions) (_result *ListIoTCloudConnectorsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.APN)) {
		query["APN"] = request.APN
	}

	if !tea.BoolValue(util.IsUnset(request.ISP)) {
		query["ISP"] = request.ISP
	}

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorGroupId)) {
		query["IoTCloudConnectorGroupId"] = request.IoTCloudConnectorGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorIds)) {
		query["IoTCloudConnectorIds"] = request.IoTCloudConnectorIds
	}

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorName)) {
		query["IoTCloudConnectorName"] = request.IoTCloudConnectorName
	}

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorStatus)) {
		query["IoTCloudConnectorStatus"] = request.IoTCloudConnectorStatus
	}

	if !tea.BoolValue(util.IsUnset(request.IsInGroup)) {
		query["IsInGroup"] = request.IsInGroup
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListIoTCloudConnectors"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIoTCloudConnectorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListIoTCloudConnectors(request *ListIoTCloudConnectorsRequest) (_result *ListIoTCloudConnectorsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListIoTCloudConnectorsResponse{}
	_body, _err := client.ListIoTCloudConnectorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListIoTCoudConnectorBackhaulRouteWithOptions(request *ListIoTCoudConnectorBackhaulRouteRequest, runtime *util.RuntimeOptions) (_result *ListIoTCoudConnectorBackhaulRouteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListIoTCoudConnectorBackhaulRoute"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIoTCoudConnectorBackhaulRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListIoTCoudConnectorBackhaulRoute(request *ListIoTCoudConnectorBackhaulRouteRequest) (_result *ListIoTCoudConnectorBackhaulRouteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListIoTCoudConnectorBackhaulRouteResponse{}
	_body, _err := client.ListIoTCoudConnectorBackhaulRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListIpMappingRulesWithOptions(request *ListIpMappingRulesRequest, runtime *util.RuntimeOptions) (_result *ListIpMappingRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DestinationIps)) {
		query["DestinationIps"] = request.DestinationIps
	}

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.IpMappingRuleIds)) {
		query["IpMappingRuleIds"] = request.IpMappingRuleIds
	}

	if !tea.BoolValue(util.IsUnset(request.IpMappingRuleNames)) {
		query["IpMappingRuleNames"] = request.IpMappingRuleNames
	}

	if !tea.BoolValue(util.IsUnset(request.IpMappingRuleStatuses)) {
		query["IpMappingRuleStatuses"] = request.IpMappingRuleStatuses
	}

	if !tea.BoolValue(util.IsUnset(request.MappingIps)) {
		query["MappingIps"] = request.MappingIps
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListIpMappingRules"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIpMappingRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListIpMappingRules(request *ListIpMappingRulesRequest) (_result *ListIpMappingRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListIpMappingRulesResponse{}
	_body, _err := client.ListIpMappingRulesWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRegions"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
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

func (client *Client) ListServiceWithOptions(request *ListServiceRequest, runtime *util.RuntimeOptions) (_result *ListServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceStatuses)) {
		query["ResourceStatuses"] = request.ResourceStatuses
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceIds)) {
		query["ServiceIds"] = request.ServiceIds
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceNames)) {
		query["ServiceNames"] = request.ServiceNames
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListService"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListService(request *ListServiceRequest) (_result *ListServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListServiceResponse{}
	_body, _err := client.ListServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListServiceEntriesWithOptions(request *ListServiceEntriesRequest, runtime *util.RuntimeOptions) (_result *ListServiceEntriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceEntryIds)) {
		query["ServiceEntryIds"] = request.ServiceEntryIds
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceEntryName)) {
		query["ServiceEntryName"] = request.ServiceEntryName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceEntryStatus)) {
		query["ServiceEntryStatus"] = request.ServiceEntryStatus
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.Target)) {
		query["Target"] = request.Target
	}

	if !tea.BoolValue(util.IsUnset(request.TargetType)) {
		query["TargetType"] = request.TargetType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServiceEntries"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceEntriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListServiceEntries(request *ListServiceEntriesRequest) (_result *ListServiceEntriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListServiceEntriesResponse{}
	_body, _err := client.ListServiceEntriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MoveAuthorizationRuleToDNSServiceWithOptions(request *MoveAuthorizationRuleToDNSServiceRequest, runtime *util.RuntimeOptions) (_result *MoveAuthorizationRuleToDNSServiceResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("MoveAuthorizationRuleToDNSService"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MoveAuthorizationRuleToDNSServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MoveAuthorizationRuleToDNSService(request *MoveAuthorizationRuleToDNSServiceRequest) (_result *MoveAuthorizationRuleToDNSServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MoveAuthorizationRuleToDNSServiceResponse{}
	_body, _err := client.MoveAuthorizationRuleToDNSServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MoveGroupAuthorizationRuleToDNSServiceWithOptions(request *MoveGroupAuthorizationRuleToDNSServiceRequest, runtime *util.RuntimeOptions) (_result *MoveGroupAuthorizationRuleToDNSServiceResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorGroupId)) {
		query["IoTCloudConnectorGroupId"] = request.IoTCloudConnectorGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("MoveGroupAuthorizationRuleToDNSService"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MoveGroupAuthorizationRuleToDNSServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MoveGroupAuthorizationRuleToDNSService(request *MoveGroupAuthorizationRuleToDNSServiceRequest) (_result *MoveGroupAuthorizationRuleToDNSServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MoveGroupAuthorizationRuleToDNSServiceResponse{}
	_body, _err := client.MoveGroupAuthorizationRuleToDNSServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenIoTCloudConnectorServiceWithOptions(request *OpenIoTCloudConnectorServiceRequest, runtime *util.RuntimeOptions) (_result *OpenIoTCloudConnectorServiceResponse, _err error) {
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
		Action:      tea.String("OpenIoTCloudConnectorService"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenIoTCloudConnectorServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenIoTCloudConnectorService(request *OpenIoTCloudConnectorServiceRequest) (_result *OpenIoTCloudConnectorServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenIoTCloudConnectorServiceResponse{}
	_body, _err := client.OpenIoTCloudConnectorServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveIoTCloudConnectorFromGroupWithOptions(request *RemoveIoTCloudConnectorFromGroupRequest, runtime *util.RuntimeOptions) (_result *RemoveIoTCloudConnectorFromGroupResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorGroupId)) {
		query["IoTCloudConnectorGroupId"] = request.IoTCloudConnectorGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveIoTCloudConnectorFromGroup"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveIoTCloudConnectorFromGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveIoTCloudConnectorFromGroup(request *RemoveIoTCloudConnectorFromGroupRequest) (_result *RemoveIoTCloudConnectorFromGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveIoTCloudConnectorFromGroupResponse{}
	_body, _err := client.RemoveIoTCloudConnectorFromGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RevertIoTCloudConnectorWithOptions(request *RevertIoTCloudConnectorRequest, runtime *util.RuntimeOptions) (_result *RevertIoTCloudConnectorResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RevertIoTCloudConnector"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RevertIoTCloudConnectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RevertIoTCloudConnector(request *RevertIoTCloudConnectorRequest) (_result *RevertIoTCloudConnectorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RevertIoTCloudConnectorResponse{}
	_body, _err := client.RevertIoTCloudConnectorWithOptions(request, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.DestinationType)) {
		query["DestinationType"] = request.DestinationType
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceUid)) {
		query["ResourceUid"] = request.ResourceUid
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitDiagnoseTaskForSingleCard"),
		Version:     tea.String("2021-05-13"),
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

func (client *Client) UpdateAuthorizationRuleAttributeWithOptions(request *UpdateAuthorizationRuleAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateAuthorizationRuleAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthorizationRuleDescription)) {
		query["AuthorizationRuleDescription"] = request.AuthorizationRuleDescription
	}

	if !tea.BoolValue(util.IsUnset(request.AuthorizationRuleId)) {
		query["AuthorizationRuleId"] = request.AuthorizationRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.AuthorizationRuleName)) {
		query["AuthorizationRuleName"] = request.AuthorizationRuleName
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
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

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.Policy)) {
		query["Policy"] = request.Policy
	}

	if !tea.BoolValue(util.IsUnset(request.Protocol)) {
		query["Protocol"] = request.Protocol
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceCidrs)) {
		query["SourceCidrs"] = request.SourceCidrs
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAuthorizationRuleAttribute"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAuthorizationRuleAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAuthorizationRuleAttribute(request *UpdateAuthorizationRuleAttributeRequest) (_result *UpdateAuthorizationRuleAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAuthorizationRuleAttributeResponse{}
	_body, _err := client.UpdateAuthorizationRuleAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateConnectionPoolAttributeWithOptions(request *UpdateConnectionPoolAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateConnectionPoolAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cidrs)) {
		query["Cidrs"] = request.Cidrs
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionPoolDescription)) {
		query["ConnectionPoolDescription"] = request.ConnectionPoolDescription
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionPoolId)) {
		query["ConnectionPoolId"] = request.ConnectionPoolId
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionPoolName)) {
		query["ConnectionPoolName"] = request.ConnectionPoolName
	}

	if !tea.BoolValue(util.IsUnset(request.Count)) {
		query["Count"] = request.Count
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateConnectionPoolAttribute"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateConnectionPoolAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateConnectionPoolAttribute(request *UpdateConnectionPoolAttributeRequest) (_result *UpdateConnectionPoolAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateConnectionPoolAttributeResponse{}
	_body, _err := client.UpdateConnectionPoolAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDNSServiceRuleAttributeWithOptions(request *UpdateDNSServiceRuleAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateDNSServiceRuleAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DNSServiceRuleDescription)) {
		query["DNSServiceRuleDescription"] = request.DNSServiceRuleDescription
	}

	if !tea.BoolValue(util.IsUnset(request.DNSServiceRuleId)) {
		query["DNSServiceRuleId"] = request.DNSServiceRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.DNSServiceRuleName)) {
		query["DNSServiceRuleName"] = request.DNSServiceRuleName
	}

	if !tea.BoolValue(util.IsUnset(request.Destination)) {
		query["Destination"] = request.Destination
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceType)) {
		query["ServiceType"] = request.ServiceType
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDNSServiceRuleAttribute"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDNSServiceRuleAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDNSServiceRuleAttribute(request *UpdateDNSServiceRuleAttributeRequest) (_result *UpdateDNSServiceRuleAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDNSServiceRuleAttributeResponse{}
	_body, _err := client.UpdateDNSServiceRuleAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateGroupAuthorizationRuleAttributeWithOptions(request *UpdateGroupAuthorizationRuleAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateGroupAuthorizationRuleAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthorizationRuleDescription)) {
		query["AuthorizationRuleDescription"] = request.AuthorizationRuleDescription
	}

	if !tea.BoolValue(util.IsUnset(request.AuthorizationRuleId)) {
		query["AuthorizationRuleId"] = request.AuthorizationRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.AuthorizationRuleName)) {
		query["AuthorizationRuleName"] = request.AuthorizationRuleName
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
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

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorGroupId)) {
		query["IoTCloudConnectorGroupId"] = request.IoTCloudConnectorGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Policy)) {
		query["Policy"] = request.Policy
	}

	if !tea.BoolValue(util.IsUnset(request.Protocol)) {
		query["Protocol"] = request.Protocol
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceCidrs)) {
		query["SourceCidrs"] = request.SourceCidrs
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateGroupAuthorizationRuleAttribute"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateGroupAuthorizationRuleAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateGroupAuthorizationRuleAttribute(request *UpdateGroupAuthorizationRuleAttributeRequest) (_result *UpdateGroupAuthorizationRuleAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateGroupAuthorizationRuleAttributeResponse{}
	_body, _err := client.UpdateGroupAuthorizationRuleAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateGroupDNSServiceRuleAttributeWithOptions(request *UpdateGroupDNSServiceRuleAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateGroupDNSServiceRuleAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DNSServiceRuleDescription)) {
		query["DNSServiceRuleDescription"] = request.DNSServiceRuleDescription
	}

	if !tea.BoolValue(util.IsUnset(request.DNSServiceRuleId)) {
		query["DNSServiceRuleId"] = request.DNSServiceRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.DNSServiceRuleName)) {
		query["DNSServiceRuleName"] = request.DNSServiceRuleName
	}

	if !tea.BoolValue(util.IsUnset(request.Destination)) {
		query["Destination"] = request.Destination
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorGroupId)) {
		query["IoTCloudConnectorGroupId"] = request.IoTCloudConnectorGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceType)) {
		query["ServiceType"] = request.ServiceType
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateGroupDNSServiceRuleAttribute"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateGroupDNSServiceRuleAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateGroupDNSServiceRuleAttribute(request *UpdateGroupDNSServiceRuleAttributeRequest) (_result *UpdateGroupDNSServiceRuleAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateGroupDNSServiceRuleAttributeResponse{}
	_body, _err := client.UpdateGroupDNSServiceRuleAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateGroupIpMappingRuleWithOptions(request *UpdateGroupIpMappingRuleRequest, runtime *util.RuntimeOptions) (_result *UpdateGroupIpMappingRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationIp)) {
		query["DestinationIp"] = request.DestinationIp
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.GroupIpMappingRuleId)) {
		query["GroupIpMappingRuleId"] = request.GroupIpMappingRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorGroupId)) {
		query["IoTCloudConnectorGroupId"] = request.IoTCloudConnectorGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.IpMappingRuleDescription)) {
		query["IpMappingRuleDescription"] = request.IpMappingRuleDescription
	}

	if !tea.BoolValue(util.IsUnset(request.IpMappingRuleName)) {
		query["IpMappingRuleName"] = request.IpMappingRuleName
	}

	if !tea.BoolValue(util.IsUnset(request.MappingIp)) {
		query["MappingIp"] = request.MappingIp
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateGroupIpMappingRule"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateGroupIpMappingRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateGroupIpMappingRule(request *UpdateGroupIpMappingRuleRequest) (_result *UpdateGroupIpMappingRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateGroupIpMappingRuleResponse{}
	_body, _err := client.UpdateGroupIpMappingRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateIoTCloudConnectorAttributeWithOptions(request *UpdateIoTCloudConnectorAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateIoTCloudConnectorAttributeResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorDescription)) {
		query["IoTCloudConnectorDescription"] = request.IoTCloudConnectorDescription
	}

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorName)) {
		query["IoTCloudConnectorName"] = request.IoTCloudConnectorName
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		query["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.WildcardDomainEnabled)) {
		query["WildcardDomainEnabled"] = request.WildcardDomainEnabled
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateIoTCloudConnectorAttribute"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateIoTCloudConnectorAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateIoTCloudConnectorAttribute(request *UpdateIoTCloudConnectorAttributeRequest) (_result *UpdateIoTCloudConnectorAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateIoTCloudConnectorAttributeResponse{}
	_body, _err := client.UpdateIoTCloudConnectorAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateIoTCloudConnectorGroupAttributeWithOptions(request *UpdateIoTCloudConnectorGroupAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateIoTCloudConnectorGroupAttributeResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorGroupId)) {
		query["IoTCloudConnectorGroupId"] = request.IoTCloudConnectorGroupId
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
		Action:      tea.String("UpdateIoTCloudConnectorGroupAttribute"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateIoTCloudConnectorGroupAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateIoTCloudConnectorGroupAttribute(request *UpdateIoTCloudConnectorGroupAttributeRequest) (_result *UpdateIoTCloudConnectorGroupAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateIoTCloudConnectorGroupAttributeResponse{}
	_body, _err := client.UpdateIoTCloudConnectorGroupAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateIpMappingRuleWithOptions(request *UpdateIpMappingRuleRequest, runtime *util.RuntimeOptions) (_result *UpdateIpMappingRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationIp)) {
		query["DestinationIp"] = request.DestinationIp
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.IpMappingRuleDescription)) {
		query["IpMappingRuleDescription"] = request.IpMappingRuleDescription
	}

	if !tea.BoolValue(util.IsUnset(request.IpMappingRuleId)) {
		query["IpMappingRuleId"] = request.IpMappingRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.IpMappingRuleName)) {
		query["IpMappingRuleName"] = request.IpMappingRuleName
	}

	if !tea.BoolValue(util.IsUnset(request.MappingIp)) {
		query["MappingIp"] = request.MappingIp
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateIpMappingRule"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateIpMappingRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateIpMappingRule(request *UpdateIpMappingRuleRequest) (_result *UpdateIpMappingRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateIpMappingRuleResponse{}
	_body, _err := client.UpdateIpMappingRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateServiceAttributeWithOptions(request *UpdateServiceAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateServiceAttributeResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceDescription)) {
		query["ServiceDescription"] = request.ServiceDescription
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateServiceAttribute"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateServiceAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateServiceAttribute(request *UpdateServiceAttributeRequest) (_result *UpdateServiceAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateServiceAttributeResponse{}
	_body, _err := client.UpdateServiceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateServiceEntryAttributeWithOptions(request *UpdateServiceEntryAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateServiceEntryAttributeResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.IoTCloudConnectorId)) {
		query["IoTCloudConnectorId"] = request.IoTCloudConnectorId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceEntryDescription)) {
		query["ServiceEntryDescription"] = request.ServiceEntryDescription
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceEntryId)) {
		query["ServiceEntryId"] = request.ServiceEntryId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceEntryName)) {
		query["ServiceEntryName"] = request.ServiceEntryName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateServiceEntryAttribute"),
		Version:     tea.String("2021-05-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateServiceEntryAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateServiceEntryAttribute(request *UpdateServiceEntryAttributeRequest) (_result *UpdateServiceEntryAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateServiceEntryAttributeResponse{}
	_body, _err := client.UpdateServiceEntryAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
