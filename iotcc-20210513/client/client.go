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
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AssociateIpWithConnectionPoolResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AssociateVSwitchWithIoTCloudConnectorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AssociateVSwitchWithIoTCloudConnectorResponse) SetBody(v *AssociateVSwitchWithIoTCloudConnectorResponseBody) *AssociateVSwitchWithIoTCloudConnectorResponse {
	s.Body = v
	return s
}

type CreateAuthorizationRuleRequest struct {
	AuthorizationRuleDescription *string   `json:"AuthorizationRuleDescription,omitempty" xml:"AuthorizationRuleDescription,omitempty"`
	AuthorizationRuleName        *string   `json:"AuthorizationRuleName,omitempty" xml:"AuthorizationRuleName,omitempty"`
	ClientToken                  *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Destination                  *string   `json:"Destination,omitempty" xml:"Destination,omitempty"`
	DestinationType              *string   `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	DryRun                       *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorId          *string   `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	Policy                       *string   `json:"Policy,omitempty" xml:"Policy,omitempty"`
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
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAuthorizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateAuthorizationRuleResponse) SetBody(v *CreateAuthorizationRuleResponseBody) *CreateAuthorizationRuleResponse {
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
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateConnectionPoolResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateConnectionPoolResponse) SetBody(v *CreateConnectionPoolResponseBody) *CreateConnectionPoolResponse {
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
	WildcardDomainEnabled        *bool   `json:"WildcardDomainEnabled,omitempty" xml:"WildcardDomainEnabled,omitempty"`
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
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateIoTCloudConnectorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateIoTCloudConnectorResponse) SetBody(v *CreateIoTCloudConnectorResponseBody) *CreateIoTCloudConnectorResponse {
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateServiceEntryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAuthorizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteAuthorizationRuleResponse) SetBody(v *DeleteAuthorizationRuleResponseBody) *DeleteAuthorizationRuleResponse {
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
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteConnectionPoolResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteConnectionPoolResponse) SetBody(v *DeleteConnectionPoolResponseBody) *DeleteConnectionPoolResponse {
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
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteIoTCloudConnectorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteIoTCloudConnectorResponse) SetBody(v *DeleteIoTCloudConnectorResponseBody) *DeleteIoTCloudConnectorResponse {
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteServiceEntryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableIoTCloudConnectorAccessLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DissociateIpFromConnectionPoolResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DissociateVSwitchFromIoTCloudConnectorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableIoTCloudConnectorAccessLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// OssPath
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
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetConnectionPoolIpOperationResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetConnectionPoolIpOperationResultResponse) SetBody(v *GetConnectionPoolIpOperationResultResponseBody) *GetConnectionPoolIpOperationResultResponse {
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
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetIoTCloudConnectorAccessLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// Sts info of accessKeyId
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// Sts info of accessKeySecret
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	// Sts info expiration time
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// OssPath
	OssPath *string `json:"OssPath,omitempty" xml:"OssPath,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Sts info of securityToken
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
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
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetStsInfoAndOssPathResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// Id of the request
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
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GrantVirtualBorderRouterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAPNsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListAPNsResponse) SetBody(v *ListAPNsResponseBody) *ListAPNsResponse {
	s.Body = v
	return s
}

type ListAuthorizationRulesRequest struct {
	AuthorizationRuleIds    []*string `json:"AuthorizationRuleIds,omitempty" xml:"AuthorizationRuleIds,omitempty" type:"Repeated"`
	AuthorizationRuleName   []*string `json:"AuthorizationRuleName,omitempty" xml:"AuthorizationRuleName,omitempty" type:"Repeated"`
	AuthorizationRuleStatus []*string `json:"AuthorizationRuleStatus,omitempty" xml:"AuthorizationRuleStatus,omitempty" type:"Repeated"`
	Destination             []*string `json:"Destination,omitempty" xml:"Destination,omitempty" type:"Repeated"`
	DestinationType         []*string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty" type:"Repeated"`
	IoTCloudConnectorId     *string   `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	MaxResults              *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken               *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Policy                  []*string `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Repeated"`
	RegionId                *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *ListAuthorizationRulesRequest) SetDestination(v []*string) *ListAuthorizationRulesRequest {
	s.Destination = v
	return s
}

func (s *ListAuthorizationRulesRequest) SetDestinationType(v []*string) *ListAuthorizationRulesRequest {
	s.DestinationType = v
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
	Destination                  *string   `json:"Destination,omitempty" xml:"Destination,omitempty"`
	DestinationType              *string   `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	IoTCloudConnectorId          *string   `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	Policy                       *string   `json:"Policy,omitempty" xml:"Policy,omitempty"`
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

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) SetDestination(v string) *ListAuthorizationRulesResponseBodyAuthorizationRules {
	s.Destination = &v
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

func (s *ListAuthorizationRulesResponseBodyAuthorizationRules) SetSourceCidrs(v []*string) *ListAuthorizationRulesResponseBodyAuthorizationRules {
	s.SourceCidrs = v
	return s
}

type ListAuthorizationRulesResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAuthorizationRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListAuthorizationRulesResponse) SetBody(v *ListAuthorizationRulesResponseBody) *ListAuthorizationRulesResponse {
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
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListConnectionPoolIpsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListConnectionPoolsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListConnectionPoolsResponse) SetBody(v *ListConnectionPoolsResponseBody) *ListConnectionPoolsResponse {
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
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListIoTCloudConnectorAvailableZonesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListIoTCloudConnectorAvailableZonesResponse) SetBody(v *ListIoTCloudConnectorAvailableZonesResponseBody) *ListIoTCloudConnectorAvailableZonesResponse {
	s.Body = v
	return s
}

type ListIoTCloudConnectorsRequest struct {
	APN                     []*string `json:"APN,omitempty" xml:"APN,omitempty" type:"Repeated"`
	ISP                     []*string `json:"ISP,omitempty" xml:"ISP,omitempty" type:"Repeated"`
	IoTCloudConnectorIds    []*string `json:"IoTCloudConnectorIds,omitempty" xml:"IoTCloudConnectorIds,omitempty" type:"Repeated"`
	IoTCloudConnectorName   []*string `json:"IoTCloudConnectorName,omitempty" xml:"IoTCloudConnectorName,omitempty" type:"Repeated"`
	IoTCloudConnectorStatus []*string `json:"IoTCloudConnectorStatus,omitempty" xml:"IoTCloudConnectorStatus,omitempty" type:"Repeated"`
	MaxResults              *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken               *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId                *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VpcId                   []*string `json:"VpcId,omitempty" xml:"VpcId,omitempty" type:"Repeated"`
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
	ISP                             *string   `json:"ISP,omitempty" xml:"ISP,omitempty"`
	IoTCloudConnectorBusinessStatus *string   `json:"IoTCloudConnectorBusinessStatus,omitempty" xml:"IoTCloudConnectorBusinessStatus,omitempty"`
	IoTCloudConnectorDescription    *string   `json:"IoTCloudConnectorDescription,omitempty" xml:"IoTCloudConnectorDescription,omitempty"`
	IoTCloudConnectorId             *string   `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	IoTCloudConnectorName           *string   `json:"IoTCloudConnectorName,omitempty" xml:"IoTCloudConnectorName,omitempty"`
	IoTCloudConnectorStatus         *string   `json:"IoTCloudConnectorStatus,omitempty" xml:"IoTCloudConnectorStatus,omitempty"`
	ModifyTime                      *int64    `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	RateLimit                       *int64    `json:"RateLimit,omitempty" xml:"RateLimit,omitempty"`
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

func (s *ListIoTCloudConnectorsResponseBodyIoTCloudConnectors) SetModifyTime(v int64) *ListIoTCloudConnectorsResponseBodyIoTCloudConnectors {
	s.ModifyTime = &v
	return s
}

func (s *ListIoTCloudConnectorsResponseBodyIoTCloudConnectors) SetRateLimit(v int64) *ListIoTCloudConnectorsResponseBodyIoTCloudConnectors {
	s.RateLimit = &v
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
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListIoTCloudConnectorsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListIoTCloudConnectorsResponse) SetBody(v *ListIoTCloudConnectorsResponseBody) *ListIoTCloudConnectorsResponse {
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
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListServiceEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListServiceEntriesResponse) SetBody(v *ListServiceEntriesResponseBody) *ListServiceEntriesResponse {
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
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OpenIoTCloudConnectorServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *OpenIoTCloudConnectorServiceResponse) SetBody(v *OpenIoTCloudConnectorServiceResponseBody) *OpenIoTCloudConnectorServiceResponse {
	s.Body = v
	return s
}

type UpdateAuthorizationRuleAttributeRequest struct {
	AuthorizationRuleDescription *string   `json:"AuthorizationRuleDescription,omitempty" xml:"AuthorizationRuleDescription,omitempty"`
	AuthorizationRuleId          *string   `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	AuthorizationRuleName        *string   `json:"AuthorizationRuleName,omitempty" xml:"AuthorizationRuleName,omitempty"`
	ClientToken                  *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Destination                  *string   `json:"Destination,omitempty" xml:"Destination,omitempty"`
	DestinationType              *string   `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	DryRun                       *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorId          *string   `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	Policy                       *string   `json:"Policy,omitempty" xml:"Policy,omitempty"`
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
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAuthorizationRuleAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateConnectionPoolAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateConnectionPoolAttributeResponse) SetBody(v *UpdateConnectionPoolAttributeResponseBody) *UpdateConnectionPoolAttributeResponse {
	s.Body = v
	return s
}

type UpdateIoTCloudConnectorAttributeRequest struct {
	ClientToken                  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IoTCloudConnectorDescription *string `json:"IoTCloudConnectorDescription,omitempty" xml:"IoTCloudConnectorDescription,omitempty"`
	IoTCloudConnectorId          *string `json:"IoTCloudConnectorId,omitempty" xml:"IoTCloudConnectorId,omitempty"`
	IoTCloudConnectorName        *string `json:"IoTCloudConnectorName,omitempty" xml:"IoTCloudConnectorName,omitempty"`
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
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateIoTCloudConnectorAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateIoTCloudConnectorAttributeResponse) SetBody(v *UpdateIoTCloudConnectorAttributeResponseBody) *UpdateIoTCloudConnectorAttributeResponse {
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
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateServiceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateServiceEntryAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (client *Client) AssociateIpWithConnectionPoolWithOptions(request *AssociateIpWithConnectionPoolRequest, runtime *util.RuntimeOptions) (_result *AssociateIpWithConnectionPoolResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AssociateIpWithConnectionPoolResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AssociateIpWithConnectionPool"), tea.String("2021-05-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AssociateVSwitchWithIoTCloudConnectorResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AssociateVSwitchWithIoTCloudConnector"), tea.String("2021-05-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) CreateAuthorizationRuleWithOptions(request *CreateAuthorizationRuleRequest, runtime *util.RuntimeOptions) (_result *CreateAuthorizationRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAuthorizationRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAuthorizationRule"), tea.String("2021-05-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) CreateConnectionPoolWithOptions(request *CreateConnectionPoolRequest, runtime *util.RuntimeOptions) (_result *CreateConnectionPoolResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateConnectionPoolResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateConnectionPool"), tea.String("2021-05-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) CreateIoTCloudConnectorWithOptions(request *CreateIoTCloudConnectorRequest, runtime *util.RuntimeOptions) (_result *CreateIoTCloudConnectorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateIoTCloudConnectorResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateIoTCloudConnector"), tea.String("2021-05-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) CreateServiceWithOptions(request *CreateServiceRequest, runtime *util.RuntimeOptions) (_result *CreateServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateService"), tea.String("2021-05-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateServiceEntryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateServiceEntry"), tea.String("2021-05-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAuthorizationRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteAuthorizationRule"), tea.String("2021-05-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DeleteConnectionPoolWithOptions(request *DeleteConnectionPoolRequest, runtime *util.RuntimeOptions) (_result *DeleteConnectionPoolResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteConnectionPoolResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteConnectionPool"), tea.String("2021-05-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DeleteIoTCloudConnectorWithOptions(request *DeleteIoTCloudConnectorRequest, runtime *util.RuntimeOptions) (_result *DeleteIoTCloudConnectorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteIoTCloudConnectorResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteIoTCloudConnector"), tea.String("2021-05-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DeleteServiceWithOptions(request *DeleteServiceRequest, runtime *util.RuntimeOptions) (_result *DeleteServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteService"), tea.String("2021-05-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteServiceEntryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteServiceEntry"), tea.String("2021-05-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DisableIoTCloudConnectorAccessLogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DisableIoTCloudConnectorAccessLog"), tea.String("2021-05-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DissociateIpFromConnectionPoolResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DissociateIpFromConnectionPool"), tea.String("2021-05-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DissociateVSwitchFromIoTCloudConnectorResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DissociateVSwitchFromIoTCloudConnector"), tea.String("2021-05-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnableIoTCloudConnectorAccessLogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableIoTCloudConnectorAccessLog"), tea.String("2021-05-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetConnectionPoolIpOperationResultResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetConnectionPoolIpOperationResult"), tea.String("2021-05-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetIoTCloudConnectorAccessLogWithOptions(request *GetIoTCloudConnectorAccessLogRequest, runtime *util.RuntimeOptions) (_result *GetIoTCloudConnectorAccessLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetIoTCloudConnectorAccessLogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetIoTCloudConnectorAccessLog"), tea.String("2021-05-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetStsInfoAndOssPathResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetStsInfoAndOssPath"), tea.String("2021-05-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GrantVirtualBorderRouterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GrantVirtualBorderRouter"), tea.String("2021-05-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListAPNsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAPNs"), tea.String("2021-05-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListAuthorizationRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAuthorizationRules"), tea.String("2021-05-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListConnectionPoolIpsWithOptions(request *ListConnectionPoolIpsRequest, runtime *util.RuntimeOptions) (_result *ListConnectionPoolIpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListConnectionPoolIpsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListConnectionPoolIps"), tea.String("2021-05-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListConnectionPoolsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListConnectionPools"), tea.String("2021-05-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListIoTCloudConnectorAvailableZonesWithOptions(request *ListIoTCloudConnectorAvailableZonesRequest, runtime *util.RuntimeOptions) (_result *ListIoTCloudConnectorAvailableZonesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListIoTCloudConnectorAvailableZonesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListIoTCloudConnectorAvailableZones"), tea.String("2021-05-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListIoTCloudConnectorsWithOptions(request *ListIoTCloudConnectorsRequest, runtime *util.RuntimeOptions) (_result *ListIoTCloudConnectorsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListIoTCloudConnectorsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListIoTCloudConnectors"), tea.String("2021-05-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListRegionsWithOptions(request *ListRegionsRequest, runtime *util.RuntimeOptions) (_result *ListRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListRegionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListRegions"), tea.String("2021-05-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListService"), tea.String("2021-05-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListServiceEntriesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListServiceEntries"), tea.String("2021-05-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) OpenIoTCloudConnectorServiceWithOptions(request *OpenIoTCloudConnectorServiceRequest, runtime *util.RuntimeOptions) (_result *OpenIoTCloudConnectorServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OpenIoTCloudConnectorServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OpenIoTCloudConnectorService"), tea.String("2021-05-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) UpdateAuthorizationRuleAttributeWithOptions(request *UpdateAuthorizationRuleAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateAuthorizationRuleAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateAuthorizationRuleAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateAuthorizationRuleAttribute"), tea.String("2021-05-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateConnectionPoolAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateConnectionPoolAttribute"), tea.String("2021-05-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) UpdateIoTCloudConnectorAttributeWithOptions(request *UpdateIoTCloudConnectorAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateIoTCloudConnectorAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateIoTCloudConnectorAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateIoTCloudConnectorAttribute"), tea.String("2021-05-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) UpdateServiceAttributeWithOptions(request *UpdateServiceAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateServiceAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateServiceAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateServiceAttribute"), tea.String("2021-05-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateServiceEntryAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateServiceEntryAttribute"), tea.String("2021-05-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
