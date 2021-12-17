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

type AddUserToVpcEndpointServiceRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceId   *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	UserId      *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AddUserToVpcEndpointServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s AddUserToVpcEndpointServiceRequest) GoString() string {
	return s.String()
}

func (s *AddUserToVpcEndpointServiceRequest) SetClientToken(v string) *AddUserToVpcEndpointServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *AddUserToVpcEndpointServiceRequest) SetDryRun(v bool) *AddUserToVpcEndpointServiceRequest {
	s.DryRun = &v
	return s
}

func (s *AddUserToVpcEndpointServiceRequest) SetRegionId(v string) *AddUserToVpcEndpointServiceRequest {
	s.RegionId = &v
	return s
}

func (s *AddUserToVpcEndpointServiceRequest) SetServiceId(v string) *AddUserToVpcEndpointServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *AddUserToVpcEndpointServiceRequest) SetUserId(v int64) *AddUserToVpcEndpointServiceRequest {
	s.UserId = &v
	return s
}

type AddUserToVpcEndpointServiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddUserToVpcEndpointServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddUserToVpcEndpointServiceResponseBody) GoString() string {
	return s.String()
}

func (s *AddUserToVpcEndpointServiceResponseBody) SetRequestId(v string) *AddUserToVpcEndpointServiceResponseBody {
	s.RequestId = &v
	return s
}

type AddUserToVpcEndpointServiceResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddUserToVpcEndpointServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddUserToVpcEndpointServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s AddUserToVpcEndpointServiceResponse) GoString() string {
	return s.String()
}

func (s *AddUserToVpcEndpointServiceResponse) SetHeaders(v map[string]*string) *AddUserToVpcEndpointServiceResponse {
	s.Headers = v
	return s
}

func (s *AddUserToVpcEndpointServiceResponse) SetBody(v *AddUserToVpcEndpointServiceResponseBody) *AddUserToVpcEndpointServiceResponse {
	s.Body = v
	return s
}

type AddZoneToVpcEndpointRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EndpointId  *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VSwitchId   *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId      *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	Ip          *string `json:"ip,omitempty" xml:"ip,omitempty"`
}

func (s AddZoneToVpcEndpointRequest) String() string {
	return tea.Prettify(s)
}

func (s AddZoneToVpcEndpointRequest) GoString() string {
	return s.String()
}

func (s *AddZoneToVpcEndpointRequest) SetClientToken(v string) *AddZoneToVpcEndpointRequest {
	s.ClientToken = &v
	return s
}

func (s *AddZoneToVpcEndpointRequest) SetDryRun(v bool) *AddZoneToVpcEndpointRequest {
	s.DryRun = &v
	return s
}

func (s *AddZoneToVpcEndpointRequest) SetEndpointId(v string) *AddZoneToVpcEndpointRequest {
	s.EndpointId = &v
	return s
}

func (s *AddZoneToVpcEndpointRequest) SetRegionId(v string) *AddZoneToVpcEndpointRequest {
	s.RegionId = &v
	return s
}

func (s *AddZoneToVpcEndpointRequest) SetVSwitchId(v string) *AddZoneToVpcEndpointRequest {
	s.VSwitchId = &v
	return s
}

func (s *AddZoneToVpcEndpointRequest) SetZoneId(v string) *AddZoneToVpcEndpointRequest {
	s.ZoneId = &v
	return s
}

func (s *AddZoneToVpcEndpointRequest) SetIp(v string) *AddZoneToVpcEndpointRequest {
	s.Ip = &v
	return s
}

type AddZoneToVpcEndpointResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddZoneToVpcEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddZoneToVpcEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *AddZoneToVpcEndpointResponseBody) SetRequestId(v string) *AddZoneToVpcEndpointResponseBody {
	s.RequestId = &v
	return s
}

type AddZoneToVpcEndpointResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddZoneToVpcEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddZoneToVpcEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s AddZoneToVpcEndpointResponse) GoString() string {
	return s.String()
}

func (s *AddZoneToVpcEndpointResponse) SetHeaders(v map[string]*string) *AddZoneToVpcEndpointResponse {
	s.Headers = v
	return s
}

func (s *AddZoneToVpcEndpointResponse) SetBody(v *AddZoneToVpcEndpointResponseBody) *AddZoneToVpcEndpointResponse {
	s.Body = v
	return s
}

type AttachResourceToVpcEndpointServiceRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ServiceId    *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s AttachResourceToVpcEndpointServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachResourceToVpcEndpointServiceRequest) GoString() string {
	return s.String()
}

func (s *AttachResourceToVpcEndpointServiceRequest) SetClientToken(v string) *AttachResourceToVpcEndpointServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *AttachResourceToVpcEndpointServiceRequest) SetDryRun(v bool) *AttachResourceToVpcEndpointServiceRequest {
	s.DryRun = &v
	return s
}

func (s *AttachResourceToVpcEndpointServiceRequest) SetRegionId(v string) *AttachResourceToVpcEndpointServiceRequest {
	s.RegionId = &v
	return s
}

func (s *AttachResourceToVpcEndpointServiceRequest) SetResourceId(v string) *AttachResourceToVpcEndpointServiceRequest {
	s.ResourceId = &v
	return s
}

func (s *AttachResourceToVpcEndpointServiceRequest) SetResourceType(v string) *AttachResourceToVpcEndpointServiceRequest {
	s.ResourceType = &v
	return s
}

func (s *AttachResourceToVpcEndpointServiceRequest) SetServiceId(v string) *AttachResourceToVpcEndpointServiceRequest {
	s.ServiceId = &v
	return s
}

type AttachResourceToVpcEndpointServiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachResourceToVpcEndpointServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachResourceToVpcEndpointServiceResponseBody) GoString() string {
	return s.String()
}

func (s *AttachResourceToVpcEndpointServiceResponseBody) SetRequestId(v string) *AttachResourceToVpcEndpointServiceResponseBody {
	s.RequestId = &v
	return s
}

type AttachResourceToVpcEndpointServiceResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AttachResourceToVpcEndpointServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttachResourceToVpcEndpointServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachResourceToVpcEndpointServiceResponse) GoString() string {
	return s.String()
}

func (s *AttachResourceToVpcEndpointServiceResponse) SetHeaders(v map[string]*string) *AttachResourceToVpcEndpointServiceResponse {
	s.Headers = v
	return s
}

func (s *AttachResourceToVpcEndpointServiceResponse) SetBody(v *AttachResourceToVpcEndpointServiceResponseBody) *AttachResourceToVpcEndpointServiceResponse {
	s.Body = v
	return s
}

type AttachSecurityGroupToVpcEndpointRequest struct {
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun          *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EndpointId      *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s AttachSecurityGroupToVpcEndpointRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachSecurityGroupToVpcEndpointRequest) GoString() string {
	return s.String()
}

func (s *AttachSecurityGroupToVpcEndpointRequest) SetClientToken(v string) *AttachSecurityGroupToVpcEndpointRequest {
	s.ClientToken = &v
	return s
}

func (s *AttachSecurityGroupToVpcEndpointRequest) SetDryRun(v bool) *AttachSecurityGroupToVpcEndpointRequest {
	s.DryRun = &v
	return s
}

func (s *AttachSecurityGroupToVpcEndpointRequest) SetEndpointId(v string) *AttachSecurityGroupToVpcEndpointRequest {
	s.EndpointId = &v
	return s
}

func (s *AttachSecurityGroupToVpcEndpointRequest) SetRegionId(v string) *AttachSecurityGroupToVpcEndpointRequest {
	s.RegionId = &v
	return s
}

func (s *AttachSecurityGroupToVpcEndpointRequest) SetSecurityGroupId(v string) *AttachSecurityGroupToVpcEndpointRequest {
	s.SecurityGroupId = &v
	return s
}

type AttachSecurityGroupToVpcEndpointResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachSecurityGroupToVpcEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachSecurityGroupToVpcEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *AttachSecurityGroupToVpcEndpointResponseBody) SetRequestId(v string) *AttachSecurityGroupToVpcEndpointResponseBody {
	s.RequestId = &v
	return s
}

type AttachSecurityGroupToVpcEndpointResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AttachSecurityGroupToVpcEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttachSecurityGroupToVpcEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachSecurityGroupToVpcEndpointResponse) GoString() string {
	return s.String()
}

func (s *AttachSecurityGroupToVpcEndpointResponse) SetHeaders(v map[string]*string) *AttachSecurityGroupToVpcEndpointResponse {
	s.Headers = v
	return s
}

func (s *AttachSecurityGroupToVpcEndpointResponse) SetBody(v *AttachSecurityGroupToVpcEndpointResponseBody) *AttachSecurityGroupToVpcEndpointResponse {
	s.Body = v
	return s
}

type CheckProductOpenResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckProductOpenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckProductOpenResponseBody) GoString() string {
	return s.String()
}

func (s *CheckProductOpenResponseBody) SetRequestId(v string) *CheckProductOpenResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckProductOpenResponseBody) SetSuccess(v bool) *CheckProductOpenResponseBody {
	s.Success = &v
	return s
}

type CheckProductOpenResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckProductOpenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckProductOpenResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckProductOpenResponse) GoString() string {
	return s.String()
}

func (s *CheckProductOpenResponse) SetHeaders(v map[string]*string) *CheckProductOpenResponse {
	s.Headers = v
	return s
}

func (s *CheckProductOpenResponse) SetBody(v *CheckProductOpenResponseBody) *CheckProductOpenResponse {
	s.Body = v
	return s
}

type CreateVpcEndpointRequest struct {
	ClientToken               *string                         `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                    *bool                           `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EndpointDescription       *string                         `json:"EndpointDescription,omitempty" xml:"EndpointDescription,omitempty"`
	EndpointName              *string                         `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
	EndpointType              *string                         `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	ProtectedEnabled          *bool                           `json:"ProtectedEnabled,omitempty" xml:"ProtectedEnabled,omitempty"`
	RegionId                  *string                         `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityGroupId           []*string                       `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" type:"Repeated"`
	ServiceId                 *string                         `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceName               *string                         `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	VpcId                     *string                         `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	Zone                      []*CreateVpcEndpointRequestZone `json:"Zone,omitempty" xml:"Zone,omitempty" type:"Repeated"`
	ZonePrivateIpAddressCount *int64                          `json:"ZonePrivateIpAddressCount,omitempty" xml:"ZonePrivateIpAddressCount,omitempty"`
}

func (s CreateVpcEndpointRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcEndpointRequest) GoString() string {
	return s.String()
}

func (s *CreateVpcEndpointRequest) SetClientToken(v string) *CreateVpcEndpointRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateVpcEndpointRequest) SetDryRun(v bool) *CreateVpcEndpointRequest {
	s.DryRun = &v
	return s
}

func (s *CreateVpcEndpointRequest) SetEndpointDescription(v string) *CreateVpcEndpointRequest {
	s.EndpointDescription = &v
	return s
}

func (s *CreateVpcEndpointRequest) SetEndpointName(v string) *CreateVpcEndpointRequest {
	s.EndpointName = &v
	return s
}

func (s *CreateVpcEndpointRequest) SetEndpointType(v string) *CreateVpcEndpointRequest {
	s.EndpointType = &v
	return s
}

func (s *CreateVpcEndpointRequest) SetProtectedEnabled(v bool) *CreateVpcEndpointRequest {
	s.ProtectedEnabled = &v
	return s
}

func (s *CreateVpcEndpointRequest) SetRegionId(v string) *CreateVpcEndpointRequest {
	s.RegionId = &v
	return s
}

func (s *CreateVpcEndpointRequest) SetSecurityGroupId(v []*string) *CreateVpcEndpointRequest {
	s.SecurityGroupId = v
	return s
}

func (s *CreateVpcEndpointRequest) SetServiceId(v string) *CreateVpcEndpointRequest {
	s.ServiceId = &v
	return s
}

func (s *CreateVpcEndpointRequest) SetServiceName(v string) *CreateVpcEndpointRequest {
	s.ServiceName = &v
	return s
}

func (s *CreateVpcEndpointRequest) SetVpcId(v string) *CreateVpcEndpointRequest {
	s.VpcId = &v
	return s
}

func (s *CreateVpcEndpointRequest) SetZone(v []*CreateVpcEndpointRequestZone) *CreateVpcEndpointRequest {
	s.Zone = v
	return s
}

func (s *CreateVpcEndpointRequest) SetZonePrivateIpAddressCount(v int64) *CreateVpcEndpointRequest {
	s.ZonePrivateIpAddressCount = &v
	return s
}

type CreateVpcEndpointRequestZone struct {
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId    *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	Ip        *string `json:"ip,omitempty" xml:"ip,omitempty"`
}

func (s CreateVpcEndpointRequestZone) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcEndpointRequestZone) GoString() string {
	return s.String()
}

func (s *CreateVpcEndpointRequestZone) SetVSwitchId(v string) *CreateVpcEndpointRequestZone {
	s.VSwitchId = &v
	return s
}

func (s *CreateVpcEndpointRequestZone) SetZoneId(v string) *CreateVpcEndpointRequestZone {
	s.ZoneId = &v
	return s
}

func (s *CreateVpcEndpointRequestZone) SetIp(v string) *CreateVpcEndpointRequestZone {
	s.Ip = &v
	return s
}

type CreateVpcEndpointResponseBody struct {
	Bandwidth              *int64  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	ConnectionStatus       *string `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	CreateTime             *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EndpointBusinessStatus *string `json:"EndpointBusinessStatus,omitempty" xml:"EndpointBusinessStatus,omitempty"`
	EndpointDescription    *string `json:"EndpointDescription,omitempty" xml:"EndpointDescription,omitempty"`
	EndpointDomain         *string `json:"EndpointDomain,omitempty" xml:"EndpointDomain,omitempty"`
	EndpointId             *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	EndpointName           *string `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
	EndpointStatus         *string `json:"EndpointStatus,omitempty" xml:"EndpointStatus,omitempty"`
	RequestId              *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceId              *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceName            *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	VpcId                  *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateVpcEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVpcEndpointResponseBody) SetBandwidth(v int64) *CreateVpcEndpointResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *CreateVpcEndpointResponseBody) SetConnectionStatus(v string) *CreateVpcEndpointResponseBody {
	s.ConnectionStatus = &v
	return s
}

func (s *CreateVpcEndpointResponseBody) SetCreateTime(v string) *CreateVpcEndpointResponseBody {
	s.CreateTime = &v
	return s
}

func (s *CreateVpcEndpointResponseBody) SetEndpointBusinessStatus(v string) *CreateVpcEndpointResponseBody {
	s.EndpointBusinessStatus = &v
	return s
}

func (s *CreateVpcEndpointResponseBody) SetEndpointDescription(v string) *CreateVpcEndpointResponseBody {
	s.EndpointDescription = &v
	return s
}

func (s *CreateVpcEndpointResponseBody) SetEndpointDomain(v string) *CreateVpcEndpointResponseBody {
	s.EndpointDomain = &v
	return s
}

func (s *CreateVpcEndpointResponseBody) SetEndpointId(v string) *CreateVpcEndpointResponseBody {
	s.EndpointId = &v
	return s
}

func (s *CreateVpcEndpointResponseBody) SetEndpointName(v string) *CreateVpcEndpointResponseBody {
	s.EndpointName = &v
	return s
}

func (s *CreateVpcEndpointResponseBody) SetEndpointStatus(v string) *CreateVpcEndpointResponseBody {
	s.EndpointStatus = &v
	return s
}

func (s *CreateVpcEndpointResponseBody) SetRequestId(v string) *CreateVpcEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVpcEndpointResponseBody) SetServiceId(v string) *CreateVpcEndpointResponseBody {
	s.ServiceId = &v
	return s
}

func (s *CreateVpcEndpointResponseBody) SetServiceName(v string) *CreateVpcEndpointResponseBody {
	s.ServiceName = &v
	return s
}

func (s *CreateVpcEndpointResponseBody) SetVpcId(v string) *CreateVpcEndpointResponseBody {
	s.VpcId = &v
	return s
}

type CreateVpcEndpointResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateVpcEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateVpcEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcEndpointResponse) GoString() string {
	return s.String()
}

func (s *CreateVpcEndpointResponse) SetHeaders(v map[string]*string) *CreateVpcEndpointResponse {
	s.Headers = v
	return s
}

func (s *CreateVpcEndpointResponse) SetBody(v *CreateVpcEndpointResponseBody) *CreateVpcEndpointResponse {
	s.Body = v
	return s
}

type CreateVpcEndpointServiceRequest struct {
	AutoAcceptEnabled   *bool                                      `json:"AutoAcceptEnabled,omitempty" xml:"AutoAcceptEnabled,omitempty"`
	ClientToken         *string                                    `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun              *bool                                      `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	Payer               *string                                    `json:"Payer,omitempty" xml:"Payer,omitempty"`
	RegionId            *string                                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Resource            []*CreateVpcEndpointServiceRequestResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Repeated"`
	ServiceDescription  *string                                    `json:"ServiceDescription,omitempty" xml:"ServiceDescription,omitempty"`
	ServiceResourceType *string                                    `json:"ServiceResourceType,omitempty" xml:"ServiceResourceType,omitempty"`
	ZoneAffinityEnabled *bool                                      `json:"ZoneAffinityEnabled,omitempty" xml:"ZoneAffinityEnabled,omitempty"`
}

func (s CreateVpcEndpointServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcEndpointServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateVpcEndpointServiceRequest) SetAutoAcceptEnabled(v bool) *CreateVpcEndpointServiceRequest {
	s.AutoAcceptEnabled = &v
	return s
}

func (s *CreateVpcEndpointServiceRequest) SetClientToken(v string) *CreateVpcEndpointServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateVpcEndpointServiceRequest) SetDryRun(v bool) *CreateVpcEndpointServiceRequest {
	s.DryRun = &v
	return s
}

func (s *CreateVpcEndpointServiceRequest) SetPayer(v string) *CreateVpcEndpointServiceRequest {
	s.Payer = &v
	return s
}

func (s *CreateVpcEndpointServiceRequest) SetRegionId(v string) *CreateVpcEndpointServiceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateVpcEndpointServiceRequest) SetResource(v []*CreateVpcEndpointServiceRequestResource) *CreateVpcEndpointServiceRequest {
	s.Resource = v
	return s
}

func (s *CreateVpcEndpointServiceRequest) SetServiceDescription(v string) *CreateVpcEndpointServiceRequest {
	s.ServiceDescription = &v
	return s
}

func (s *CreateVpcEndpointServiceRequest) SetServiceResourceType(v string) *CreateVpcEndpointServiceRequest {
	s.ServiceResourceType = &v
	return s
}

func (s *CreateVpcEndpointServiceRequest) SetZoneAffinityEnabled(v bool) *CreateVpcEndpointServiceRequest {
	s.ZoneAffinityEnabled = &v
	return s
}

type CreateVpcEndpointServiceRequestResource struct {
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s CreateVpcEndpointServiceRequestResource) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcEndpointServiceRequestResource) GoString() string {
	return s.String()
}

func (s *CreateVpcEndpointServiceRequestResource) SetResourceId(v string) *CreateVpcEndpointServiceRequestResource {
	s.ResourceId = &v
	return s
}

func (s *CreateVpcEndpointServiceRequestResource) SetResourceType(v string) *CreateVpcEndpointServiceRequestResource {
	s.ResourceType = &v
	return s
}

type CreateVpcEndpointServiceResponseBody struct {
	AutoAcceptEnabled     *bool   `json:"AutoAcceptEnabled,omitempty" xml:"AutoAcceptEnabled,omitempty"`
	CreateTime            *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	RequestId             *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceBusinessStatus *string `json:"ServiceBusinessStatus,omitempty" xml:"ServiceBusinessStatus,omitempty"`
	ServiceDescription    *string `json:"ServiceDescription,omitempty" xml:"ServiceDescription,omitempty"`
	ServiceDomain         *string `json:"ServiceDomain,omitempty" xml:"ServiceDomain,omitempty"`
	ServiceId             *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceName           *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	ServiceStatus         *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	ZoneAffinityEnabled   *bool   `json:"ZoneAffinityEnabled,omitempty" xml:"ZoneAffinityEnabled,omitempty"`
}

func (s CreateVpcEndpointServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcEndpointServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVpcEndpointServiceResponseBody) SetAutoAcceptEnabled(v bool) *CreateVpcEndpointServiceResponseBody {
	s.AutoAcceptEnabled = &v
	return s
}

func (s *CreateVpcEndpointServiceResponseBody) SetCreateTime(v string) *CreateVpcEndpointServiceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *CreateVpcEndpointServiceResponseBody) SetRequestId(v string) *CreateVpcEndpointServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVpcEndpointServiceResponseBody) SetServiceBusinessStatus(v string) *CreateVpcEndpointServiceResponseBody {
	s.ServiceBusinessStatus = &v
	return s
}

func (s *CreateVpcEndpointServiceResponseBody) SetServiceDescription(v string) *CreateVpcEndpointServiceResponseBody {
	s.ServiceDescription = &v
	return s
}

func (s *CreateVpcEndpointServiceResponseBody) SetServiceDomain(v string) *CreateVpcEndpointServiceResponseBody {
	s.ServiceDomain = &v
	return s
}

func (s *CreateVpcEndpointServiceResponseBody) SetServiceId(v string) *CreateVpcEndpointServiceResponseBody {
	s.ServiceId = &v
	return s
}

func (s *CreateVpcEndpointServiceResponseBody) SetServiceName(v string) *CreateVpcEndpointServiceResponseBody {
	s.ServiceName = &v
	return s
}

func (s *CreateVpcEndpointServiceResponseBody) SetServiceStatus(v string) *CreateVpcEndpointServiceResponseBody {
	s.ServiceStatus = &v
	return s
}

func (s *CreateVpcEndpointServiceResponseBody) SetZoneAffinityEnabled(v bool) *CreateVpcEndpointServiceResponseBody {
	s.ZoneAffinityEnabled = &v
	return s
}

type CreateVpcEndpointServiceResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateVpcEndpointServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateVpcEndpointServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcEndpointServiceResponse) GoString() string {
	return s.String()
}

func (s *CreateVpcEndpointServiceResponse) SetHeaders(v map[string]*string) *CreateVpcEndpointServiceResponse {
	s.Headers = v
	return s
}

func (s *CreateVpcEndpointServiceResponse) SetBody(v *CreateVpcEndpointServiceResponseBody) *CreateVpcEndpointServiceResponse {
	s.Body = v
	return s
}

type DeleteVpcEndpointRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EndpointId  *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteVpcEndpointRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpcEndpointRequest) GoString() string {
	return s.String()
}

func (s *DeleteVpcEndpointRequest) SetClientToken(v string) *DeleteVpcEndpointRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteVpcEndpointRequest) SetDryRun(v bool) *DeleteVpcEndpointRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteVpcEndpointRequest) SetEndpointId(v string) *DeleteVpcEndpointRequest {
	s.EndpointId = &v
	return s
}

func (s *DeleteVpcEndpointRequest) SetRegionId(v string) *DeleteVpcEndpointRequest {
	s.RegionId = &v
	return s
}

type DeleteVpcEndpointResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVpcEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpcEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVpcEndpointResponseBody) SetRequestId(v string) *DeleteVpcEndpointResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVpcEndpointResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteVpcEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVpcEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpcEndpointResponse) GoString() string {
	return s.String()
}

func (s *DeleteVpcEndpointResponse) SetHeaders(v map[string]*string) *DeleteVpcEndpointResponse {
	s.Headers = v
	return s
}

func (s *DeleteVpcEndpointResponse) SetBody(v *DeleteVpcEndpointResponseBody) *DeleteVpcEndpointResponse {
	s.Body = v
	return s
}

type DeleteVpcEndpointServiceRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceId   *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s DeleteVpcEndpointServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpcEndpointServiceRequest) GoString() string {
	return s.String()
}

func (s *DeleteVpcEndpointServiceRequest) SetClientToken(v string) *DeleteVpcEndpointServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteVpcEndpointServiceRequest) SetDryRun(v bool) *DeleteVpcEndpointServiceRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteVpcEndpointServiceRequest) SetRegionId(v string) *DeleteVpcEndpointServiceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteVpcEndpointServiceRequest) SetServiceId(v string) *DeleteVpcEndpointServiceRequest {
	s.ServiceId = &v
	return s
}

type DeleteVpcEndpointServiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVpcEndpointServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpcEndpointServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVpcEndpointServiceResponseBody) SetRequestId(v string) *DeleteVpcEndpointServiceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVpcEndpointServiceResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteVpcEndpointServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVpcEndpointServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpcEndpointServiceResponse) GoString() string {
	return s.String()
}

func (s *DeleteVpcEndpointServiceResponse) SetHeaders(v map[string]*string) *DeleteVpcEndpointServiceResponse {
	s.Headers = v
	return s
}

func (s *DeleteVpcEndpointServiceResponse) SetBody(v *DeleteVpcEndpointServiceResponseBody) *DeleteVpcEndpointServiceResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetRegionId(v string) *DescribeRegionsRequest {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponseBody struct {
	Regions   *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
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
	LocalName      *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionEndpoint = &v
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

type DescribeZonesRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeZonesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeZonesRequest) SetRegionId(v string) *DescribeZonesRequest {
	s.RegionId = &v
	return s
}

type DescribeZonesResponseBody struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Zones     *DescribeZonesResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
}

func (s DescribeZonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBody) SetRequestId(v string) *DescribeZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeZonesResponseBody) SetZones(v *DescribeZonesResponseBodyZones) *DescribeZonesResponseBody {
	s.Zones = v
	return s
}

type DescribeZonesResponseBodyZones struct {
	Zone []*DescribeZonesResponseBodyZonesZone `json:"Zone,omitempty" xml:"Zone,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBodyZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBodyZones) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZones) SetZone(v []*DescribeZonesResponseBodyZonesZone) *DescribeZonesResponseBodyZones {
	s.Zone = v
	return s
}

type DescribeZonesResponseBodyZonesZone struct {
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	ZoneId    *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeZonesResponseBodyZonesZone) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZone) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZone) SetLocalName(v string) *DescribeZonesResponseBodyZonesZone {
	s.LocalName = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetZoneId(v string) *DescribeZonesResponseBodyZonesZone {
	s.ZoneId = &v
	return s
}

type DescribeZonesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeZonesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeZonesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponse) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponse) SetHeaders(v map[string]*string) *DescribeZonesResponse {
	s.Headers = v
	return s
}

func (s *DescribeZonesResponse) SetBody(v *DescribeZonesResponseBody) *DescribeZonesResponse {
	s.Body = v
	return s
}

type DetachResourceFromVpcEndpointServiceRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ServiceId    *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s DetachResourceFromVpcEndpointServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachResourceFromVpcEndpointServiceRequest) GoString() string {
	return s.String()
}

func (s *DetachResourceFromVpcEndpointServiceRequest) SetClientToken(v string) *DetachResourceFromVpcEndpointServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *DetachResourceFromVpcEndpointServiceRequest) SetDryRun(v bool) *DetachResourceFromVpcEndpointServiceRequest {
	s.DryRun = &v
	return s
}

func (s *DetachResourceFromVpcEndpointServiceRequest) SetRegionId(v string) *DetachResourceFromVpcEndpointServiceRequest {
	s.RegionId = &v
	return s
}

func (s *DetachResourceFromVpcEndpointServiceRequest) SetResourceId(v string) *DetachResourceFromVpcEndpointServiceRequest {
	s.ResourceId = &v
	return s
}

func (s *DetachResourceFromVpcEndpointServiceRequest) SetResourceType(v string) *DetachResourceFromVpcEndpointServiceRequest {
	s.ResourceType = &v
	return s
}

func (s *DetachResourceFromVpcEndpointServiceRequest) SetServiceId(v string) *DetachResourceFromVpcEndpointServiceRequest {
	s.ServiceId = &v
	return s
}

type DetachResourceFromVpcEndpointServiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachResourceFromVpcEndpointServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachResourceFromVpcEndpointServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DetachResourceFromVpcEndpointServiceResponseBody) SetRequestId(v string) *DetachResourceFromVpcEndpointServiceResponseBody {
	s.RequestId = &v
	return s
}

type DetachResourceFromVpcEndpointServiceResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetachResourceFromVpcEndpointServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetachResourceFromVpcEndpointServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachResourceFromVpcEndpointServiceResponse) GoString() string {
	return s.String()
}

func (s *DetachResourceFromVpcEndpointServiceResponse) SetHeaders(v map[string]*string) *DetachResourceFromVpcEndpointServiceResponse {
	s.Headers = v
	return s
}

func (s *DetachResourceFromVpcEndpointServiceResponse) SetBody(v *DetachResourceFromVpcEndpointServiceResponseBody) *DetachResourceFromVpcEndpointServiceResponse {
	s.Body = v
	return s
}

type DetachSecurityGroupFromVpcEndpointRequest struct {
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun          *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EndpointId      *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s DetachSecurityGroupFromVpcEndpointRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachSecurityGroupFromVpcEndpointRequest) GoString() string {
	return s.String()
}

func (s *DetachSecurityGroupFromVpcEndpointRequest) SetClientToken(v string) *DetachSecurityGroupFromVpcEndpointRequest {
	s.ClientToken = &v
	return s
}

func (s *DetachSecurityGroupFromVpcEndpointRequest) SetDryRun(v bool) *DetachSecurityGroupFromVpcEndpointRequest {
	s.DryRun = &v
	return s
}

func (s *DetachSecurityGroupFromVpcEndpointRequest) SetEndpointId(v string) *DetachSecurityGroupFromVpcEndpointRequest {
	s.EndpointId = &v
	return s
}

func (s *DetachSecurityGroupFromVpcEndpointRequest) SetRegionId(v string) *DetachSecurityGroupFromVpcEndpointRequest {
	s.RegionId = &v
	return s
}

func (s *DetachSecurityGroupFromVpcEndpointRequest) SetSecurityGroupId(v string) *DetachSecurityGroupFromVpcEndpointRequest {
	s.SecurityGroupId = &v
	return s
}

type DetachSecurityGroupFromVpcEndpointResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachSecurityGroupFromVpcEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachSecurityGroupFromVpcEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *DetachSecurityGroupFromVpcEndpointResponseBody) SetRequestId(v string) *DetachSecurityGroupFromVpcEndpointResponseBody {
	s.RequestId = &v
	return s
}

type DetachSecurityGroupFromVpcEndpointResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetachSecurityGroupFromVpcEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetachSecurityGroupFromVpcEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachSecurityGroupFromVpcEndpointResponse) GoString() string {
	return s.String()
}

func (s *DetachSecurityGroupFromVpcEndpointResponse) SetHeaders(v map[string]*string) *DetachSecurityGroupFromVpcEndpointResponse {
	s.Headers = v
	return s
}

func (s *DetachSecurityGroupFromVpcEndpointResponse) SetBody(v *DetachSecurityGroupFromVpcEndpointResponseBody) *DetachSecurityGroupFromVpcEndpointResponse {
	s.Body = v
	return s
}

type DisableVpcEndpointConnectionRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EndpointId  *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceId   *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s DisableVpcEndpointConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableVpcEndpointConnectionRequest) GoString() string {
	return s.String()
}

func (s *DisableVpcEndpointConnectionRequest) SetClientToken(v string) *DisableVpcEndpointConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *DisableVpcEndpointConnectionRequest) SetDryRun(v bool) *DisableVpcEndpointConnectionRequest {
	s.DryRun = &v
	return s
}

func (s *DisableVpcEndpointConnectionRequest) SetEndpointId(v string) *DisableVpcEndpointConnectionRequest {
	s.EndpointId = &v
	return s
}

func (s *DisableVpcEndpointConnectionRequest) SetRegionId(v string) *DisableVpcEndpointConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *DisableVpcEndpointConnectionRequest) SetServiceId(v string) *DisableVpcEndpointConnectionRequest {
	s.ServiceId = &v
	return s
}

type DisableVpcEndpointConnectionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableVpcEndpointConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableVpcEndpointConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *DisableVpcEndpointConnectionResponseBody) SetRequestId(v string) *DisableVpcEndpointConnectionResponseBody {
	s.RequestId = &v
	return s
}

type DisableVpcEndpointConnectionResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableVpcEndpointConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableVpcEndpointConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableVpcEndpointConnectionResponse) GoString() string {
	return s.String()
}

func (s *DisableVpcEndpointConnectionResponse) SetHeaders(v map[string]*string) *DisableVpcEndpointConnectionResponse {
	s.Headers = v
	return s
}

func (s *DisableVpcEndpointConnectionResponse) SetBody(v *DisableVpcEndpointConnectionResponseBody) *DisableVpcEndpointConnectionResponse {
	s.Body = v
	return s
}

type EnableVpcEndpointConnectionRequest struct {
	Bandwidth   *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EndpointId  *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceId   *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s EnableVpcEndpointConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableVpcEndpointConnectionRequest) GoString() string {
	return s.String()
}

func (s *EnableVpcEndpointConnectionRequest) SetBandwidth(v int32) *EnableVpcEndpointConnectionRequest {
	s.Bandwidth = &v
	return s
}

func (s *EnableVpcEndpointConnectionRequest) SetClientToken(v string) *EnableVpcEndpointConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *EnableVpcEndpointConnectionRequest) SetDryRun(v bool) *EnableVpcEndpointConnectionRequest {
	s.DryRun = &v
	return s
}

func (s *EnableVpcEndpointConnectionRequest) SetEndpointId(v string) *EnableVpcEndpointConnectionRequest {
	s.EndpointId = &v
	return s
}

func (s *EnableVpcEndpointConnectionRequest) SetRegionId(v string) *EnableVpcEndpointConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *EnableVpcEndpointConnectionRequest) SetServiceId(v string) *EnableVpcEndpointConnectionRequest {
	s.ServiceId = &v
	return s
}

type EnableVpcEndpointConnectionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableVpcEndpointConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableVpcEndpointConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *EnableVpcEndpointConnectionResponseBody) SetRequestId(v string) *EnableVpcEndpointConnectionResponseBody {
	s.RequestId = &v
	return s
}

type EnableVpcEndpointConnectionResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableVpcEndpointConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableVpcEndpointConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableVpcEndpointConnectionResponse) GoString() string {
	return s.String()
}

func (s *EnableVpcEndpointConnectionResponse) SetHeaders(v map[string]*string) *EnableVpcEndpointConnectionResponse {
	s.Headers = v
	return s
}

func (s *EnableVpcEndpointConnectionResponse) SetBody(v *EnableVpcEndpointConnectionResponseBody) *EnableVpcEndpointConnectionResponse {
	s.Body = v
	return s
}

type GetVpcEndpointAttributeRequest struct {
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetVpcEndpointAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVpcEndpointAttributeRequest) GoString() string {
	return s.String()
}

func (s *GetVpcEndpointAttributeRequest) SetEndpointId(v string) *GetVpcEndpointAttributeRequest {
	s.EndpointId = &v
	return s
}

func (s *GetVpcEndpointAttributeRequest) SetRegionId(v string) *GetVpcEndpointAttributeRequest {
	s.RegionId = &v
	return s
}

type GetVpcEndpointAttributeResponseBody struct {
	Bandwidth                 *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	ConnectionStatus          *string `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	CreateTime                *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EndpointBusinessStatus    *string `json:"EndpointBusinessStatus,omitempty" xml:"EndpointBusinessStatus,omitempty"`
	EndpointDescription       *string `json:"EndpointDescription,omitempty" xml:"EndpointDescription,omitempty"`
	EndpointDomain            *string `json:"EndpointDomain,omitempty" xml:"EndpointDomain,omitempty"`
	EndpointId                *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	EndpointName              *string `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
	EndpointStatus            *string `json:"EndpointStatus,omitempty" xml:"EndpointStatus,omitempty"`
	EndpointType              *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	Payer                     *string `json:"Payer,omitempty" xml:"Payer,omitempty"`
	RegionId                  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId                 *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceOwner             *bool   `json:"ResourceOwner,omitempty" xml:"ResourceOwner,omitempty"`
	ServiceId                 *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceName               *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	VpcId                     *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneAffinityEnabled       *bool   `json:"ZoneAffinityEnabled,omitempty" xml:"ZoneAffinityEnabled,omitempty"`
	ZonePrivateIpAddressCount *int64  `json:"ZonePrivateIpAddressCount,omitempty" xml:"ZonePrivateIpAddressCount,omitempty"`
}

func (s GetVpcEndpointAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVpcEndpointAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetVpcEndpointAttributeResponseBody) SetBandwidth(v int32) *GetVpcEndpointAttributeResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *GetVpcEndpointAttributeResponseBody) SetConnectionStatus(v string) *GetVpcEndpointAttributeResponseBody {
	s.ConnectionStatus = &v
	return s
}

func (s *GetVpcEndpointAttributeResponseBody) SetCreateTime(v string) *GetVpcEndpointAttributeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetVpcEndpointAttributeResponseBody) SetEndpointBusinessStatus(v string) *GetVpcEndpointAttributeResponseBody {
	s.EndpointBusinessStatus = &v
	return s
}

func (s *GetVpcEndpointAttributeResponseBody) SetEndpointDescription(v string) *GetVpcEndpointAttributeResponseBody {
	s.EndpointDescription = &v
	return s
}

func (s *GetVpcEndpointAttributeResponseBody) SetEndpointDomain(v string) *GetVpcEndpointAttributeResponseBody {
	s.EndpointDomain = &v
	return s
}

func (s *GetVpcEndpointAttributeResponseBody) SetEndpointId(v string) *GetVpcEndpointAttributeResponseBody {
	s.EndpointId = &v
	return s
}

func (s *GetVpcEndpointAttributeResponseBody) SetEndpointName(v string) *GetVpcEndpointAttributeResponseBody {
	s.EndpointName = &v
	return s
}

func (s *GetVpcEndpointAttributeResponseBody) SetEndpointStatus(v string) *GetVpcEndpointAttributeResponseBody {
	s.EndpointStatus = &v
	return s
}

func (s *GetVpcEndpointAttributeResponseBody) SetEndpointType(v string) *GetVpcEndpointAttributeResponseBody {
	s.EndpointType = &v
	return s
}

func (s *GetVpcEndpointAttributeResponseBody) SetPayer(v string) *GetVpcEndpointAttributeResponseBody {
	s.Payer = &v
	return s
}

func (s *GetVpcEndpointAttributeResponseBody) SetRegionId(v string) *GetVpcEndpointAttributeResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetVpcEndpointAttributeResponseBody) SetRequestId(v string) *GetVpcEndpointAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVpcEndpointAttributeResponseBody) SetResourceOwner(v bool) *GetVpcEndpointAttributeResponseBody {
	s.ResourceOwner = &v
	return s
}

func (s *GetVpcEndpointAttributeResponseBody) SetServiceId(v string) *GetVpcEndpointAttributeResponseBody {
	s.ServiceId = &v
	return s
}

func (s *GetVpcEndpointAttributeResponseBody) SetServiceName(v string) *GetVpcEndpointAttributeResponseBody {
	s.ServiceName = &v
	return s
}

func (s *GetVpcEndpointAttributeResponseBody) SetVpcId(v string) *GetVpcEndpointAttributeResponseBody {
	s.VpcId = &v
	return s
}

func (s *GetVpcEndpointAttributeResponseBody) SetZoneAffinityEnabled(v bool) *GetVpcEndpointAttributeResponseBody {
	s.ZoneAffinityEnabled = &v
	return s
}

func (s *GetVpcEndpointAttributeResponseBody) SetZonePrivateIpAddressCount(v int64) *GetVpcEndpointAttributeResponseBody {
	s.ZonePrivateIpAddressCount = &v
	return s
}

type GetVpcEndpointAttributeResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetVpcEndpointAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetVpcEndpointAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVpcEndpointAttributeResponse) GoString() string {
	return s.String()
}

func (s *GetVpcEndpointAttributeResponse) SetHeaders(v map[string]*string) *GetVpcEndpointAttributeResponse {
	s.Headers = v
	return s
}

func (s *GetVpcEndpointAttributeResponse) SetBody(v *GetVpcEndpointAttributeResponseBody) *GetVpcEndpointAttributeResponse {
	s.Body = v
	return s
}

type GetVpcEndpointServiceAttributeRequest struct {
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s GetVpcEndpointServiceAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVpcEndpointServiceAttributeRequest) GoString() string {
	return s.String()
}

func (s *GetVpcEndpointServiceAttributeRequest) SetRegionId(v string) *GetVpcEndpointServiceAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeRequest) SetServiceId(v string) *GetVpcEndpointServiceAttributeRequest {
	s.ServiceId = &v
	return s
}

type GetVpcEndpointServiceAttributeResponseBody struct {
	AutoAcceptEnabled     *bool     `json:"AutoAcceptEnabled,omitempty" xml:"AutoAcceptEnabled,omitempty"`
	ConnectBandwidth      *int32    `json:"ConnectBandwidth,omitempty" xml:"ConnectBandwidth,omitempty"`
	CreateTime            *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	MaxBandwidth          *int32    `json:"MaxBandwidth,omitempty" xml:"MaxBandwidth,omitempty"`
	MinBandwidth          *int32    `json:"MinBandwidth,omitempty" xml:"MinBandwidth,omitempty"`
	Payer                 *string   `json:"Payer,omitempty" xml:"Payer,omitempty"`
	RegionId              *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId             *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceBusinessStatus *string   `json:"ServiceBusinessStatus,omitempty" xml:"ServiceBusinessStatus,omitempty"`
	ServiceDescription    *string   `json:"ServiceDescription,omitempty" xml:"ServiceDescription,omitempty"`
	ServiceDomain         *string   `json:"ServiceDomain,omitempty" xml:"ServiceDomain,omitempty"`
	ServiceId             *string   `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceName           *string   `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	ServiceResourceType   *string   `json:"ServiceResourceType,omitempty" xml:"ServiceResourceType,omitempty"`
	ServiceStatus         *string   `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	ServiceType           *string   `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	ZoneAffinityEnabled   *bool     `json:"ZoneAffinityEnabled,omitempty" xml:"ZoneAffinityEnabled,omitempty"`
	Zones                 []*string `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s GetVpcEndpointServiceAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVpcEndpointServiceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetVpcEndpointServiceAttributeResponseBody) SetAutoAcceptEnabled(v bool) *GetVpcEndpointServiceAttributeResponseBody {
	s.AutoAcceptEnabled = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponseBody) SetConnectBandwidth(v int32) *GetVpcEndpointServiceAttributeResponseBody {
	s.ConnectBandwidth = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponseBody) SetCreateTime(v string) *GetVpcEndpointServiceAttributeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponseBody) SetMaxBandwidth(v int32) *GetVpcEndpointServiceAttributeResponseBody {
	s.MaxBandwidth = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponseBody) SetMinBandwidth(v int32) *GetVpcEndpointServiceAttributeResponseBody {
	s.MinBandwidth = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponseBody) SetPayer(v string) *GetVpcEndpointServiceAttributeResponseBody {
	s.Payer = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponseBody) SetRegionId(v string) *GetVpcEndpointServiceAttributeResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponseBody) SetRequestId(v string) *GetVpcEndpointServiceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponseBody) SetServiceBusinessStatus(v string) *GetVpcEndpointServiceAttributeResponseBody {
	s.ServiceBusinessStatus = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponseBody) SetServiceDescription(v string) *GetVpcEndpointServiceAttributeResponseBody {
	s.ServiceDescription = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponseBody) SetServiceDomain(v string) *GetVpcEndpointServiceAttributeResponseBody {
	s.ServiceDomain = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponseBody) SetServiceId(v string) *GetVpcEndpointServiceAttributeResponseBody {
	s.ServiceId = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponseBody) SetServiceName(v string) *GetVpcEndpointServiceAttributeResponseBody {
	s.ServiceName = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponseBody) SetServiceResourceType(v string) *GetVpcEndpointServiceAttributeResponseBody {
	s.ServiceResourceType = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponseBody) SetServiceStatus(v string) *GetVpcEndpointServiceAttributeResponseBody {
	s.ServiceStatus = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponseBody) SetServiceType(v string) *GetVpcEndpointServiceAttributeResponseBody {
	s.ServiceType = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponseBody) SetZoneAffinityEnabled(v bool) *GetVpcEndpointServiceAttributeResponseBody {
	s.ZoneAffinityEnabled = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponseBody) SetZones(v []*string) *GetVpcEndpointServiceAttributeResponseBody {
	s.Zones = v
	return s
}

type GetVpcEndpointServiceAttributeResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetVpcEndpointServiceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetVpcEndpointServiceAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVpcEndpointServiceAttributeResponse) GoString() string {
	return s.String()
}

func (s *GetVpcEndpointServiceAttributeResponse) SetHeaders(v map[string]*string) *GetVpcEndpointServiceAttributeResponse {
	s.Headers = v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponse) SetBody(v *GetVpcEndpointServiceAttributeResponseBody) *GetVpcEndpointServiceAttributeResponse {
	s.Body = v
	return s
}

type ListVpcEndpointConnectionsRequest struct {
	ConnectionStatus *string `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	EndpointId       *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	EndpointOwnerId  *int64  `json:"EndpointOwnerId,omitempty" xml:"EndpointOwnerId,omitempty"`
	EniId            *string `json:"EniId,omitempty" xml:"EniId,omitempty"`
	MaxResults       *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken        *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceId        *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s ListVpcEndpointConnectionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointConnectionsRequest) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointConnectionsRequest) SetConnectionStatus(v string) *ListVpcEndpointConnectionsRequest {
	s.ConnectionStatus = &v
	return s
}

func (s *ListVpcEndpointConnectionsRequest) SetEndpointId(v string) *ListVpcEndpointConnectionsRequest {
	s.EndpointId = &v
	return s
}

func (s *ListVpcEndpointConnectionsRequest) SetEndpointOwnerId(v int64) *ListVpcEndpointConnectionsRequest {
	s.EndpointOwnerId = &v
	return s
}

func (s *ListVpcEndpointConnectionsRequest) SetEniId(v string) *ListVpcEndpointConnectionsRequest {
	s.EniId = &v
	return s
}

func (s *ListVpcEndpointConnectionsRequest) SetMaxResults(v int32) *ListVpcEndpointConnectionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListVpcEndpointConnectionsRequest) SetNextToken(v string) *ListVpcEndpointConnectionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListVpcEndpointConnectionsRequest) SetRegionId(v string) *ListVpcEndpointConnectionsRequest {
	s.RegionId = &v
	return s
}

func (s *ListVpcEndpointConnectionsRequest) SetServiceId(v string) *ListVpcEndpointConnectionsRequest {
	s.ServiceId = &v
	return s
}

type ListVpcEndpointConnectionsResponseBody struct {
	Connections []*ListVpcEndpointConnectionsResponseBodyConnections `json:"Connections,omitempty" xml:"Connections,omitempty" type:"Repeated"`
	MaxResults  *string                                              `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string                                              `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId   *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListVpcEndpointConnectionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointConnectionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointConnectionsResponseBody) SetConnections(v []*ListVpcEndpointConnectionsResponseBodyConnections) *ListVpcEndpointConnectionsResponseBody {
	s.Connections = v
	return s
}

func (s *ListVpcEndpointConnectionsResponseBody) SetMaxResults(v string) *ListVpcEndpointConnectionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListVpcEndpointConnectionsResponseBody) SetNextToken(v string) *ListVpcEndpointConnectionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListVpcEndpointConnectionsResponseBody) SetRequestId(v string) *ListVpcEndpointConnectionsResponseBody {
	s.RequestId = &v
	return s
}

type ListVpcEndpointConnectionsResponseBodyConnections struct {
	Bandwidth        *int32                                                    `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	ConnectionStatus *string                                                   `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	EndpointId       *string                                                   `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	EndpointOwnerId  *int64                                                    `json:"EndpointOwnerId,omitempty" xml:"EndpointOwnerId,omitempty"`
	EndpointVpcId    *string                                                   `json:"EndpointVpcId,omitempty" xml:"EndpointVpcId,omitempty"`
	ModifiedTime     *string                                                   `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ResourceOwner    *bool                                                     `json:"ResourceOwner,omitempty" xml:"ResourceOwner,omitempty"`
	ServiceId        *string                                                   `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	Zones            []*ListVpcEndpointConnectionsResponseBodyConnectionsZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s ListVpcEndpointConnectionsResponseBodyConnections) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointConnectionsResponseBodyConnections) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointConnectionsResponseBodyConnections) SetBandwidth(v int32) *ListVpcEndpointConnectionsResponseBodyConnections {
	s.Bandwidth = &v
	return s
}

func (s *ListVpcEndpointConnectionsResponseBodyConnections) SetConnectionStatus(v string) *ListVpcEndpointConnectionsResponseBodyConnections {
	s.ConnectionStatus = &v
	return s
}

func (s *ListVpcEndpointConnectionsResponseBodyConnections) SetEndpointId(v string) *ListVpcEndpointConnectionsResponseBodyConnections {
	s.EndpointId = &v
	return s
}

func (s *ListVpcEndpointConnectionsResponseBodyConnections) SetEndpointOwnerId(v int64) *ListVpcEndpointConnectionsResponseBodyConnections {
	s.EndpointOwnerId = &v
	return s
}

func (s *ListVpcEndpointConnectionsResponseBodyConnections) SetEndpointVpcId(v string) *ListVpcEndpointConnectionsResponseBodyConnections {
	s.EndpointVpcId = &v
	return s
}

func (s *ListVpcEndpointConnectionsResponseBodyConnections) SetModifiedTime(v string) *ListVpcEndpointConnectionsResponseBodyConnections {
	s.ModifiedTime = &v
	return s
}

func (s *ListVpcEndpointConnectionsResponseBodyConnections) SetResourceOwner(v bool) *ListVpcEndpointConnectionsResponseBodyConnections {
	s.ResourceOwner = &v
	return s
}

func (s *ListVpcEndpointConnectionsResponseBodyConnections) SetServiceId(v string) *ListVpcEndpointConnectionsResponseBodyConnections {
	s.ServiceId = &v
	return s
}

func (s *ListVpcEndpointConnectionsResponseBodyConnections) SetZones(v []*ListVpcEndpointConnectionsResponseBodyConnectionsZones) *ListVpcEndpointConnectionsResponseBodyConnections {
	s.Zones = v
	return s
}

type ListVpcEndpointConnectionsResponseBodyConnectionsZones struct {
	EniId      *string `json:"EniId,omitempty" xml:"EniId,omitempty"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	VSwitchId  *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneDomain *string `json:"ZoneDomain,omitempty" xml:"ZoneDomain,omitempty"`
	ZoneId     *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListVpcEndpointConnectionsResponseBodyConnectionsZones) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointConnectionsResponseBodyConnectionsZones) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointConnectionsResponseBodyConnectionsZones) SetEniId(v string) *ListVpcEndpointConnectionsResponseBodyConnectionsZones {
	s.EniId = &v
	return s
}

func (s *ListVpcEndpointConnectionsResponseBodyConnectionsZones) SetResourceId(v string) *ListVpcEndpointConnectionsResponseBodyConnectionsZones {
	s.ResourceId = &v
	return s
}

func (s *ListVpcEndpointConnectionsResponseBodyConnectionsZones) SetVSwitchId(v string) *ListVpcEndpointConnectionsResponseBodyConnectionsZones {
	s.VSwitchId = &v
	return s
}

func (s *ListVpcEndpointConnectionsResponseBodyConnectionsZones) SetZoneDomain(v string) *ListVpcEndpointConnectionsResponseBodyConnectionsZones {
	s.ZoneDomain = &v
	return s
}

func (s *ListVpcEndpointConnectionsResponseBodyConnectionsZones) SetZoneId(v string) *ListVpcEndpointConnectionsResponseBodyConnectionsZones {
	s.ZoneId = &v
	return s
}

type ListVpcEndpointConnectionsResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListVpcEndpointConnectionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVpcEndpointConnectionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointConnectionsResponse) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointConnectionsResponse) SetHeaders(v map[string]*string) *ListVpcEndpointConnectionsResponse {
	s.Headers = v
	return s
}

func (s *ListVpcEndpointConnectionsResponse) SetBody(v *ListVpcEndpointConnectionsResponseBody) *ListVpcEndpointConnectionsResponse {
	s.Body = v
	return s
}

type ListVpcEndpointSecurityGroupsRequest struct {
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListVpcEndpointSecurityGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointSecurityGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointSecurityGroupsRequest) SetEndpointId(v string) *ListVpcEndpointSecurityGroupsRequest {
	s.EndpointId = &v
	return s
}

func (s *ListVpcEndpointSecurityGroupsRequest) SetMaxResults(v int32) *ListVpcEndpointSecurityGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListVpcEndpointSecurityGroupsRequest) SetNextToken(v string) *ListVpcEndpointSecurityGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListVpcEndpointSecurityGroupsRequest) SetRegionId(v string) *ListVpcEndpointSecurityGroupsRequest {
	s.RegionId = &v
	return s
}

type ListVpcEndpointSecurityGroupsResponseBody struct {
	MaxResults     *string                                                    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken      *string                                                    `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId      *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecurityGroups []*ListVpcEndpointSecurityGroupsResponseBodySecurityGroups `json:"SecurityGroups,omitempty" xml:"SecurityGroups,omitempty" type:"Repeated"`
}

func (s ListVpcEndpointSecurityGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointSecurityGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointSecurityGroupsResponseBody) SetMaxResults(v string) *ListVpcEndpointSecurityGroupsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListVpcEndpointSecurityGroupsResponseBody) SetNextToken(v string) *ListVpcEndpointSecurityGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListVpcEndpointSecurityGroupsResponseBody) SetRequestId(v string) *ListVpcEndpointSecurityGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVpcEndpointSecurityGroupsResponseBody) SetSecurityGroups(v []*ListVpcEndpointSecurityGroupsResponseBodySecurityGroups) *ListVpcEndpointSecurityGroupsResponseBody {
	s.SecurityGroups = v
	return s
}

type ListVpcEndpointSecurityGroupsResponseBodySecurityGroups struct {
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s ListVpcEndpointSecurityGroupsResponseBodySecurityGroups) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointSecurityGroupsResponseBodySecurityGroups) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointSecurityGroupsResponseBodySecurityGroups) SetSecurityGroupId(v string) *ListVpcEndpointSecurityGroupsResponseBodySecurityGroups {
	s.SecurityGroupId = &v
	return s
}

type ListVpcEndpointSecurityGroupsResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListVpcEndpointSecurityGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVpcEndpointSecurityGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointSecurityGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointSecurityGroupsResponse) SetHeaders(v map[string]*string) *ListVpcEndpointSecurityGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListVpcEndpointSecurityGroupsResponse) SetBody(v *ListVpcEndpointSecurityGroupsResponseBody) *ListVpcEndpointSecurityGroupsResponse {
	s.Body = v
	return s
}

type ListVpcEndpointServiceResourcesRequest struct {
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceId  *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s ListVpcEndpointServiceResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointServiceResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServiceResourcesRequest) SetMaxResults(v int32) *ListVpcEndpointServiceResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListVpcEndpointServiceResourcesRequest) SetNextToken(v string) *ListVpcEndpointServiceResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListVpcEndpointServiceResourcesRequest) SetRegionId(v string) *ListVpcEndpointServiceResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListVpcEndpointServiceResourcesRequest) SetServiceId(v string) *ListVpcEndpointServiceResourcesRequest {
	s.ServiceId = &v
	return s
}

type ListVpcEndpointServiceResourcesResponseBody struct {
	MaxResults *string                                                 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                                                 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resources  []*ListVpcEndpointServiceResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s ListVpcEndpointServiceResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointServiceResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServiceResourcesResponseBody) SetMaxResults(v string) *ListVpcEndpointServiceResourcesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListVpcEndpointServiceResourcesResponseBody) SetNextToken(v string) *ListVpcEndpointServiceResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListVpcEndpointServiceResourcesResponseBody) SetRequestId(v string) *ListVpcEndpointServiceResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVpcEndpointServiceResourcesResponseBody) SetResources(v []*ListVpcEndpointServiceResourcesResponseBodyResources) *ListVpcEndpointServiceResourcesResponseBody {
	s.Resources = v
	return s
}

type ListVpcEndpointServiceResourcesResponseBodyResources struct {
	Ip           *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	VSwitchId    *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId        *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId       *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListVpcEndpointServiceResourcesResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointServiceResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServiceResourcesResponseBodyResources) SetIp(v string) *ListVpcEndpointServiceResourcesResponseBodyResources {
	s.Ip = &v
	return s
}

func (s *ListVpcEndpointServiceResourcesResponseBodyResources) SetRegionId(v string) *ListVpcEndpointServiceResourcesResponseBodyResources {
	s.RegionId = &v
	return s
}

func (s *ListVpcEndpointServiceResourcesResponseBodyResources) SetResourceId(v string) *ListVpcEndpointServiceResourcesResponseBodyResources {
	s.ResourceId = &v
	return s
}

func (s *ListVpcEndpointServiceResourcesResponseBodyResources) SetResourceType(v string) *ListVpcEndpointServiceResourcesResponseBodyResources {
	s.ResourceType = &v
	return s
}

func (s *ListVpcEndpointServiceResourcesResponseBodyResources) SetVSwitchId(v string) *ListVpcEndpointServiceResourcesResponseBodyResources {
	s.VSwitchId = &v
	return s
}

func (s *ListVpcEndpointServiceResourcesResponseBodyResources) SetVpcId(v string) *ListVpcEndpointServiceResourcesResponseBodyResources {
	s.VpcId = &v
	return s
}

func (s *ListVpcEndpointServiceResourcesResponseBodyResources) SetZoneId(v string) *ListVpcEndpointServiceResourcesResponseBodyResources {
	s.ZoneId = &v
	return s
}

type ListVpcEndpointServiceResourcesResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListVpcEndpointServiceResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVpcEndpointServiceResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointServiceResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServiceResourcesResponse) SetHeaders(v map[string]*string) *ListVpcEndpointServiceResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListVpcEndpointServiceResourcesResponse) SetBody(v *ListVpcEndpointServiceResourcesResponseBody) *ListVpcEndpointServiceResourcesResponse {
	s.Body = v
	return s
}

type ListVpcEndpointServiceUsersRequest struct {
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceId  *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	UserId     *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListVpcEndpointServiceUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointServiceUsersRequest) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServiceUsersRequest) SetMaxResults(v int32) *ListVpcEndpointServiceUsersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListVpcEndpointServiceUsersRequest) SetNextToken(v string) *ListVpcEndpointServiceUsersRequest {
	s.NextToken = &v
	return s
}

func (s *ListVpcEndpointServiceUsersRequest) SetRegionId(v string) *ListVpcEndpointServiceUsersRequest {
	s.RegionId = &v
	return s
}

func (s *ListVpcEndpointServiceUsersRequest) SetServiceId(v string) *ListVpcEndpointServiceUsersRequest {
	s.ServiceId = &v
	return s
}

func (s *ListVpcEndpointServiceUsersRequest) SetUserId(v int64) *ListVpcEndpointServiceUsersRequest {
	s.UserId = &v
	return s
}

type ListVpcEndpointServiceUsersResponseBody struct {
	MaxResults *string                                         `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                                         `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Users      []*ListVpcEndpointServiceUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ListVpcEndpointServiceUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointServiceUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServiceUsersResponseBody) SetMaxResults(v string) *ListVpcEndpointServiceUsersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListVpcEndpointServiceUsersResponseBody) SetNextToken(v string) *ListVpcEndpointServiceUsersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListVpcEndpointServiceUsersResponseBody) SetRequestId(v string) *ListVpcEndpointServiceUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVpcEndpointServiceUsersResponseBody) SetUsers(v []*ListVpcEndpointServiceUsersResponseBodyUsers) *ListVpcEndpointServiceUsersResponseBody {
	s.Users = v
	return s
}

type ListVpcEndpointServiceUsersResponseBodyUsers struct {
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListVpcEndpointServiceUsersResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointServiceUsersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServiceUsersResponseBodyUsers) SetUserId(v int64) *ListVpcEndpointServiceUsersResponseBodyUsers {
	s.UserId = &v
	return s
}

type ListVpcEndpointServiceUsersResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListVpcEndpointServiceUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVpcEndpointServiceUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointServiceUsersResponse) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServiceUsersResponse) SetHeaders(v map[string]*string) *ListVpcEndpointServiceUsersResponse {
	s.Headers = v
	return s
}

func (s *ListVpcEndpointServiceUsersResponse) SetBody(v *ListVpcEndpointServiceUsersResponseBody) *ListVpcEndpointServiceUsersResponse {
	s.Body = v
	return s
}

type ListVpcEndpointServicesRequest struct {
	AutoAcceptEnabled     *bool   `json:"AutoAcceptEnabled,omitempty" xml:"AutoAcceptEnabled,omitempty"`
	MaxResults            *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken             *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceBusinessStatus *string `json:"ServiceBusinessStatus,omitempty" xml:"ServiceBusinessStatus,omitempty"`
	ServiceId             *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceName           *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	ServiceResourceType   *string `json:"ServiceResourceType,omitempty" xml:"ServiceResourceType,omitempty"`
	ServiceStatus         *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	ZoneAffinityEnabled   *bool   `json:"ZoneAffinityEnabled,omitempty" xml:"ZoneAffinityEnabled,omitempty"`
}

func (s ListVpcEndpointServicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointServicesRequest) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServicesRequest) SetAutoAcceptEnabled(v bool) *ListVpcEndpointServicesRequest {
	s.AutoAcceptEnabled = &v
	return s
}

func (s *ListVpcEndpointServicesRequest) SetMaxResults(v int32) *ListVpcEndpointServicesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListVpcEndpointServicesRequest) SetNextToken(v string) *ListVpcEndpointServicesRequest {
	s.NextToken = &v
	return s
}

func (s *ListVpcEndpointServicesRequest) SetRegionId(v string) *ListVpcEndpointServicesRequest {
	s.RegionId = &v
	return s
}

func (s *ListVpcEndpointServicesRequest) SetServiceBusinessStatus(v string) *ListVpcEndpointServicesRequest {
	s.ServiceBusinessStatus = &v
	return s
}

func (s *ListVpcEndpointServicesRequest) SetServiceId(v string) *ListVpcEndpointServicesRequest {
	s.ServiceId = &v
	return s
}

func (s *ListVpcEndpointServicesRequest) SetServiceName(v string) *ListVpcEndpointServicesRequest {
	s.ServiceName = &v
	return s
}

func (s *ListVpcEndpointServicesRequest) SetServiceResourceType(v string) *ListVpcEndpointServicesRequest {
	s.ServiceResourceType = &v
	return s
}

func (s *ListVpcEndpointServicesRequest) SetServiceStatus(v string) *ListVpcEndpointServicesRequest {
	s.ServiceStatus = &v
	return s
}

func (s *ListVpcEndpointServicesRequest) SetZoneAffinityEnabled(v bool) *ListVpcEndpointServicesRequest {
	s.ZoneAffinityEnabled = &v
	return s
}

type ListVpcEndpointServicesResponseBody struct {
	MaxResults *string                                        `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                                        `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Services   []*ListVpcEndpointServicesResponseBodyServices `json:"Services,omitempty" xml:"Services,omitempty" type:"Repeated"`
}

func (s ListVpcEndpointServicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServicesResponseBody) SetMaxResults(v string) *ListVpcEndpointServicesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBody) SetNextToken(v string) *ListVpcEndpointServicesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBody) SetRequestId(v string) *ListVpcEndpointServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBody) SetServices(v []*ListVpcEndpointServicesResponseBodyServices) *ListVpcEndpointServicesResponseBody {
	s.Services = v
	return s
}

type ListVpcEndpointServicesResponseBodyServices struct {
	AutoAcceptEnabled     *bool   `json:"AutoAcceptEnabled,omitempty" xml:"AutoAcceptEnabled,omitempty"`
	ConnectBandwidth      *int32  `json:"ConnectBandwidth,omitempty" xml:"ConnectBandwidth,omitempty"`
	CreateTime            *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	MaxBandwidth          *int32  `json:"MaxBandwidth,omitempty" xml:"MaxBandwidth,omitempty"`
	MinBandwidth          *int32  `json:"MinBandwidth,omitempty" xml:"MinBandwidth,omitempty"`
	Payer                 *string `json:"Payer,omitempty" xml:"Payer,omitempty"`
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceBusinessStatus *string `json:"ServiceBusinessStatus,omitempty" xml:"ServiceBusinessStatus,omitempty"`
	ServiceDescription    *string `json:"ServiceDescription,omitempty" xml:"ServiceDescription,omitempty"`
	ServiceDomain         *string `json:"ServiceDomain,omitempty" xml:"ServiceDomain,omitempty"`
	ServiceId             *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceName           *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	ServiceResourceType   *string `json:"ServiceResourceType,omitempty" xml:"ServiceResourceType,omitempty"`
	ServiceStatus         *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	ServiceType           *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	ZoneAffinityEnabled   *bool   `json:"ZoneAffinityEnabled,omitempty" xml:"ZoneAffinityEnabled,omitempty"`
}

func (s ListVpcEndpointServicesResponseBodyServices) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointServicesResponseBodyServices) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServicesResponseBodyServices) SetAutoAcceptEnabled(v bool) *ListVpcEndpointServicesResponseBodyServices {
	s.AutoAcceptEnabled = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServices) SetConnectBandwidth(v int32) *ListVpcEndpointServicesResponseBodyServices {
	s.ConnectBandwidth = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServices) SetCreateTime(v string) *ListVpcEndpointServicesResponseBodyServices {
	s.CreateTime = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServices) SetMaxBandwidth(v int32) *ListVpcEndpointServicesResponseBodyServices {
	s.MaxBandwidth = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServices) SetMinBandwidth(v int32) *ListVpcEndpointServicesResponseBodyServices {
	s.MinBandwidth = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServices) SetPayer(v string) *ListVpcEndpointServicesResponseBodyServices {
	s.Payer = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServices) SetRegionId(v string) *ListVpcEndpointServicesResponseBodyServices {
	s.RegionId = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServices) SetServiceBusinessStatus(v string) *ListVpcEndpointServicesResponseBodyServices {
	s.ServiceBusinessStatus = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServices) SetServiceDescription(v string) *ListVpcEndpointServicesResponseBodyServices {
	s.ServiceDescription = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServices) SetServiceDomain(v string) *ListVpcEndpointServicesResponseBodyServices {
	s.ServiceDomain = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServices) SetServiceId(v string) *ListVpcEndpointServicesResponseBodyServices {
	s.ServiceId = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServices) SetServiceName(v string) *ListVpcEndpointServicesResponseBodyServices {
	s.ServiceName = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServices) SetServiceResourceType(v string) *ListVpcEndpointServicesResponseBodyServices {
	s.ServiceResourceType = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServices) SetServiceStatus(v string) *ListVpcEndpointServicesResponseBodyServices {
	s.ServiceStatus = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServices) SetServiceType(v string) *ListVpcEndpointServicesResponseBodyServices {
	s.ServiceType = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServices) SetZoneAffinityEnabled(v bool) *ListVpcEndpointServicesResponseBodyServices {
	s.ZoneAffinityEnabled = &v
	return s
}

type ListVpcEndpointServicesResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListVpcEndpointServicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVpcEndpointServicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointServicesResponse) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServicesResponse) SetHeaders(v map[string]*string) *ListVpcEndpointServicesResponse {
	s.Headers = v
	return s
}

func (s *ListVpcEndpointServicesResponse) SetBody(v *ListVpcEndpointServicesResponseBody) *ListVpcEndpointServicesResponse {
	s.Body = v
	return s
}

type ListVpcEndpointServicesByEndUserRequest struct {
	MaxResults  *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceId   *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s ListVpcEndpointServicesByEndUserRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointServicesByEndUserRequest) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServicesByEndUserRequest) SetMaxResults(v int32) *ListVpcEndpointServicesByEndUserRequest {
	s.MaxResults = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserRequest) SetNextToken(v string) *ListVpcEndpointServicesByEndUserRequest {
	s.NextToken = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserRequest) SetRegionId(v string) *ListVpcEndpointServicesByEndUserRequest {
	s.RegionId = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserRequest) SetServiceId(v string) *ListVpcEndpointServicesByEndUserRequest {
	s.ServiceId = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserRequest) SetServiceName(v string) *ListVpcEndpointServicesByEndUserRequest {
	s.ServiceName = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserRequest) SetServiceType(v string) *ListVpcEndpointServicesByEndUserRequest {
	s.ServiceType = &v
	return s
}

type ListVpcEndpointServicesByEndUserResponseBody struct {
	MaxResults *string                                                 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                                                 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Services   []*ListVpcEndpointServicesByEndUserResponseBodyServices `json:"Services,omitempty" xml:"Services,omitempty" type:"Repeated"`
}

func (s ListVpcEndpointServicesByEndUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointServicesByEndUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServicesByEndUserResponseBody) SetMaxResults(v string) *ListVpcEndpointServicesByEndUserResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponseBody) SetNextToken(v string) *ListVpcEndpointServicesByEndUserResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponseBody) SetRequestId(v string) *ListVpcEndpointServicesByEndUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponseBody) SetServices(v []*ListVpcEndpointServicesByEndUserResponseBodyServices) *ListVpcEndpointServicesByEndUserResponseBody {
	s.Services = v
	return s
}

type ListVpcEndpointServicesByEndUserResponseBodyServices struct {
	Payer         *string   `json:"Payer,omitempty" xml:"Payer,omitempty"`
	ServiceDomain *string   `json:"ServiceDomain,omitempty" xml:"ServiceDomain,omitempty"`
	ServiceId     *string   `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceName   *string   `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	ServiceType   *string   `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	Zones         []*string `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s ListVpcEndpointServicesByEndUserResponseBodyServices) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointServicesByEndUserResponseBodyServices) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) SetPayer(v string) *ListVpcEndpointServicesByEndUserResponseBodyServices {
	s.Payer = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) SetServiceDomain(v string) *ListVpcEndpointServicesByEndUserResponseBodyServices {
	s.ServiceDomain = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) SetServiceId(v string) *ListVpcEndpointServicesByEndUserResponseBodyServices {
	s.ServiceId = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) SetServiceName(v string) *ListVpcEndpointServicesByEndUserResponseBodyServices {
	s.ServiceName = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) SetServiceType(v string) *ListVpcEndpointServicesByEndUserResponseBodyServices {
	s.ServiceType = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) SetZones(v []*string) *ListVpcEndpointServicesByEndUserResponseBodyServices {
	s.Zones = v
	return s
}

type ListVpcEndpointServicesByEndUserResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListVpcEndpointServicesByEndUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVpcEndpointServicesByEndUserResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointServicesByEndUserResponse) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServicesByEndUserResponse) SetHeaders(v map[string]*string) *ListVpcEndpointServicesByEndUserResponse {
	s.Headers = v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponse) SetBody(v *ListVpcEndpointServicesByEndUserResponseBody) *ListVpcEndpointServicesByEndUserResponse {
	s.Body = v
	return s
}

type ListVpcEndpointZonesRequest struct {
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListVpcEndpointZonesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointZonesRequest) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointZonesRequest) SetEndpointId(v string) *ListVpcEndpointZonesRequest {
	s.EndpointId = &v
	return s
}

func (s *ListVpcEndpointZonesRequest) SetMaxResults(v int32) *ListVpcEndpointZonesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListVpcEndpointZonesRequest) SetNextToken(v string) *ListVpcEndpointZonesRequest {
	s.NextToken = &v
	return s
}

func (s *ListVpcEndpointZonesRequest) SetRegionId(v string) *ListVpcEndpointZonesRequest {
	s.RegionId = &v
	return s
}

type ListVpcEndpointZonesResponseBody struct {
	MaxResults *string                                  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                                  `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Zones      []*ListVpcEndpointZonesResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s ListVpcEndpointZonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointZonesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointZonesResponseBody) SetMaxResults(v string) *ListVpcEndpointZonesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListVpcEndpointZonesResponseBody) SetNextToken(v string) *ListVpcEndpointZonesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListVpcEndpointZonesResponseBody) SetRequestId(v string) *ListVpcEndpointZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVpcEndpointZonesResponseBody) SetZones(v []*ListVpcEndpointZonesResponseBodyZones) *ListVpcEndpointZonesResponseBody {
	s.Zones = v
	return s
}

type ListVpcEndpointZonesResponseBodyZones struct {
	EniId         *string `json:"EniId,omitempty" xml:"EniId,omitempty"`
	EniIp         *string `json:"EniIp,omitempty" xml:"EniIp,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceStatus *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	VSwitchId     *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneDomain    *string `json:"ZoneDomain,omitempty" xml:"ZoneDomain,omitempty"`
	ZoneId        *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ZoneStatus    *string `json:"ZoneStatus,omitempty" xml:"ZoneStatus,omitempty"`
}

func (s ListVpcEndpointZonesResponseBodyZones) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointZonesResponseBodyZones) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointZonesResponseBodyZones) SetEniId(v string) *ListVpcEndpointZonesResponseBodyZones {
	s.EniId = &v
	return s
}

func (s *ListVpcEndpointZonesResponseBodyZones) SetEniIp(v string) *ListVpcEndpointZonesResponseBodyZones {
	s.EniIp = &v
	return s
}

func (s *ListVpcEndpointZonesResponseBodyZones) SetRegionId(v string) *ListVpcEndpointZonesResponseBodyZones {
	s.RegionId = &v
	return s
}

func (s *ListVpcEndpointZonesResponseBodyZones) SetServiceStatus(v string) *ListVpcEndpointZonesResponseBodyZones {
	s.ServiceStatus = &v
	return s
}

func (s *ListVpcEndpointZonesResponseBodyZones) SetVSwitchId(v string) *ListVpcEndpointZonesResponseBodyZones {
	s.VSwitchId = &v
	return s
}

func (s *ListVpcEndpointZonesResponseBodyZones) SetZoneDomain(v string) *ListVpcEndpointZonesResponseBodyZones {
	s.ZoneDomain = &v
	return s
}

func (s *ListVpcEndpointZonesResponseBodyZones) SetZoneId(v string) *ListVpcEndpointZonesResponseBodyZones {
	s.ZoneId = &v
	return s
}

func (s *ListVpcEndpointZonesResponseBodyZones) SetZoneStatus(v string) *ListVpcEndpointZonesResponseBodyZones {
	s.ZoneStatus = &v
	return s
}

type ListVpcEndpointZonesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListVpcEndpointZonesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVpcEndpointZonesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointZonesResponse) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointZonesResponse) SetHeaders(v map[string]*string) *ListVpcEndpointZonesResponse {
	s.Headers = v
	return s
}

func (s *ListVpcEndpointZonesResponse) SetBody(v *ListVpcEndpointZonesResponseBody) *ListVpcEndpointZonesResponse {
	s.Body = v
	return s
}

type ListVpcEndpointsRequest struct {
	ConnectionStatus *string `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	EndpointId       *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	EndpointName     *string `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
	EndpointStatus   *string `json:"EndpointStatus,omitempty" xml:"EndpointStatus,omitempty"`
	EndpointType     *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	MaxResults       *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken        *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceName      *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	VpcId            *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListVpcEndpointsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointsRequest) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointsRequest) SetConnectionStatus(v string) *ListVpcEndpointsRequest {
	s.ConnectionStatus = &v
	return s
}

func (s *ListVpcEndpointsRequest) SetEndpointId(v string) *ListVpcEndpointsRequest {
	s.EndpointId = &v
	return s
}

func (s *ListVpcEndpointsRequest) SetEndpointName(v string) *ListVpcEndpointsRequest {
	s.EndpointName = &v
	return s
}

func (s *ListVpcEndpointsRequest) SetEndpointStatus(v string) *ListVpcEndpointsRequest {
	s.EndpointStatus = &v
	return s
}

func (s *ListVpcEndpointsRequest) SetEndpointType(v string) *ListVpcEndpointsRequest {
	s.EndpointType = &v
	return s
}

func (s *ListVpcEndpointsRequest) SetMaxResults(v int32) *ListVpcEndpointsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListVpcEndpointsRequest) SetNextToken(v string) *ListVpcEndpointsRequest {
	s.NextToken = &v
	return s
}

func (s *ListVpcEndpointsRequest) SetRegionId(v string) *ListVpcEndpointsRequest {
	s.RegionId = &v
	return s
}

func (s *ListVpcEndpointsRequest) SetServiceName(v string) *ListVpcEndpointsRequest {
	s.ServiceName = &v
	return s
}

func (s *ListVpcEndpointsRequest) SetVpcId(v string) *ListVpcEndpointsRequest {
	s.VpcId = &v
	return s
}

type ListVpcEndpointsResponseBody struct {
	Endpoints  []*ListVpcEndpointsResponseBodyEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	MaxResults *string                                  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                                  `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListVpcEndpointsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointsResponseBody) SetEndpoints(v []*ListVpcEndpointsResponseBodyEndpoints) *ListVpcEndpointsResponseBody {
	s.Endpoints = v
	return s
}

func (s *ListVpcEndpointsResponseBody) SetMaxResults(v string) *ListVpcEndpointsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListVpcEndpointsResponseBody) SetNextToken(v string) *ListVpcEndpointsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListVpcEndpointsResponseBody) SetRequestId(v string) *ListVpcEndpointsResponseBody {
	s.RequestId = &v
	return s
}

type ListVpcEndpointsResponseBodyEndpoints struct {
	Bandwidth              *int64  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	ConnectionStatus       *string `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	CreateTime             *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EndpointBusinessStatus *string `json:"EndpointBusinessStatus,omitempty" xml:"EndpointBusinessStatus,omitempty"`
	EndpointDescription    *string `json:"EndpointDescription,omitempty" xml:"EndpointDescription,omitempty"`
	EndpointDomain         *string `json:"EndpointDomain,omitempty" xml:"EndpointDomain,omitempty"`
	EndpointId             *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	EndpointName           *string `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
	EndpointStatus         *string `json:"EndpointStatus,omitempty" xml:"EndpointStatus,omitempty"`
	EndpointType           *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwner          *bool   `json:"ResourceOwner,omitempty" xml:"ResourceOwner,omitempty"`
	ServiceId              *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceName            *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	VpcId                  *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneAffinityEnabled    *bool   `json:"ZoneAffinityEnabled,omitempty" xml:"ZoneAffinityEnabled,omitempty"`
}

func (s ListVpcEndpointsResponseBodyEndpoints) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointsResponseBodyEndpoints) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointsResponseBodyEndpoints) SetBandwidth(v int64) *ListVpcEndpointsResponseBodyEndpoints {
	s.Bandwidth = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyEndpoints) SetConnectionStatus(v string) *ListVpcEndpointsResponseBodyEndpoints {
	s.ConnectionStatus = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyEndpoints) SetCreateTime(v string) *ListVpcEndpointsResponseBodyEndpoints {
	s.CreateTime = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyEndpoints) SetEndpointBusinessStatus(v string) *ListVpcEndpointsResponseBodyEndpoints {
	s.EndpointBusinessStatus = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyEndpoints) SetEndpointDescription(v string) *ListVpcEndpointsResponseBodyEndpoints {
	s.EndpointDescription = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyEndpoints) SetEndpointDomain(v string) *ListVpcEndpointsResponseBodyEndpoints {
	s.EndpointDomain = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyEndpoints) SetEndpointId(v string) *ListVpcEndpointsResponseBodyEndpoints {
	s.EndpointId = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyEndpoints) SetEndpointName(v string) *ListVpcEndpointsResponseBodyEndpoints {
	s.EndpointName = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyEndpoints) SetEndpointStatus(v string) *ListVpcEndpointsResponseBodyEndpoints {
	s.EndpointStatus = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyEndpoints) SetEndpointType(v string) *ListVpcEndpointsResponseBodyEndpoints {
	s.EndpointType = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyEndpoints) SetRegionId(v string) *ListVpcEndpointsResponseBodyEndpoints {
	s.RegionId = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyEndpoints) SetResourceOwner(v bool) *ListVpcEndpointsResponseBodyEndpoints {
	s.ResourceOwner = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyEndpoints) SetServiceId(v string) *ListVpcEndpointsResponseBodyEndpoints {
	s.ServiceId = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyEndpoints) SetServiceName(v string) *ListVpcEndpointsResponseBodyEndpoints {
	s.ServiceName = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyEndpoints) SetVpcId(v string) *ListVpcEndpointsResponseBodyEndpoints {
	s.VpcId = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyEndpoints) SetZoneAffinityEnabled(v bool) *ListVpcEndpointsResponseBodyEndpoints {
	s.ZoneAffinityEnabled = &v
	return s
}

type ListVpcEndpointsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListVpcEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVpcEndpointsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointsResponse) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointsResponse) SetHeaders(v map[string]*string) *ListVpcEndpointsResponse {
	s.Headers = v
	return s
}

func (s *ListVpcEndpointsResponse) SetBody(v *ListVpcEndpointsResponseBody) *ListVpcEndpointsResponse {
	s.Body = v
	return s
}

type OpenPrivateLinkServiceRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s OpenPrivateLinkServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenPrivateLinkServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenPrivateLinkServiceRequest) SetOwnerId(v int64) *OpenPrivateLinkServiceRequest {
	s.OwnerId = &v
	return s
}

type OpenPrivateLinkServiceResponseBody struct {
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenPrivateLinkServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenPrivateLinkServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenPrivateLinkServiceResponseBody) SetOrderId(v string) *OpenPrivateLinkServiceResponseBody {
	s.OrderId = &v
	return s
}

func (s *OpenPrivateLinkServiceResponseBody) SetRequestId(v string) *OpenPrivateLinkServiceResponseBody {
	s.RequestId = &v
	return s
}

type OpenPrivateLinkServiceResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OpenPrivateLinkServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpenPrivateLinkServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenPrivateLinkServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenPrivateLinkServiceResponse) SetHeaders(v map[string]*string) *OpenPrivateLinkServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenPrivateLinkServiceResponse) SetBody(v *OpenPrivateLinkServiceResponseBody) *OpenPrivateLinkServiceResponse {
	s.Body = v
	return s
}

type RemoveUserFromVpcEndpointServiceRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceId   *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	UserId      *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s RemoveUserFromVpcEndpointServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserFromVpcEndpointServiceRequest) GoString() string {
	return s.String()
}

func (s *RemoveUserFromVpcEndpointServiceRequest) SetClientToken(v string) *RemoveUserFromVpcEndpointServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *RemoveUserFromVpcEndpointServiceRequest) SetDryRun(v bool) *RemoveUserFromVpcEndpointServiceRequest {
	s.DryRun = &v
	return s
}

func (s *RemoveUserFromVpcEndpointServiceRequest) SetRegionId(v string) *RemoveUserFromVpcEndpointServiceRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveUserFromVpcEndpointServiceRequest) SetServiceId(v string) *RemoveUserFromVpcEndpointServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *RemoveUserFromVpcEndpointServiceRequest) SetUserId(v int64) *RemoveUserFromVpcEndpointServiceRequest {
	s.UserId = &v
	return s
}

type RemoveUserFromVpcEndpointServiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveUserFromVpcEndpointServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserFromVpcEndpointServiceResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveUserFromVpcEndpointServiceResponseBody) SetRequestId(v string) *RemoveUserFromVpcEndpointServiceResponseBody {
	s.RequestId = &v
	return s
}

type RemoveUserFromVpcEndpointServiceResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveUserFromVpcEndpointServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveUserFromVpcEndpointServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserFromVpcEndpointServiceResponse) GoString() string {
	return s.String()
}

func (s *RemoveUserFromVpcEndpointServiceResponse) SetHeaders(v map[string]*string) *RemoveUserFromVpcEndpointServiceResponse {
	s.Headers = v
	return s
}

func (s *RemoveUserFromVpcEndpointServiceResponse) SetBody(v *RemoveUserFromVpcEndpointServiceResponseBody) *RemoveUserFromVpcEndpointServiceResponse {
	s.Body = v
	return s
}

type RemoveZoneFromVpcEndpointRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EndpointId  *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId      *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s RemoveZoneFromVpcEndpointRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveZoneFromVpcEndpointRequest) GoString() string {
	return s.String()
}

func (s *RemoveZoneFromVpcEndpointRequest) SetClientToken(v string) *RemoveZoneFromVpcEndpointRequest {
	s.ClientToken = &v
	return s
}

func (s *RemoveZoneFromVpcEndpointRequest) SetDryRun(v bool) *RemoveZoneFromVpcEndpointRequest {
	s.DryRun = &v
	return s
}

func (s *RemoveZoneFromVpcEndpointRequest) SetEndpointId(v string) *RemoveZoneFromVpcEndpointRequest {
	s.EndpointId = &v
	return s
}

func (s *RemoveZoneFromVpcEndpointRequest) SetRegionId(v string) *RemoveZoneFromVpcEndpointRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveZoneFromVpcEndpointRequest) SetZoneId(v string) *RemoveZoneFromVpcEndpointRequest {
	s.ZoneId = &v
	return s
}

type RemoveZoneFromVpcEndpointResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveZoneFromVpcEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveZoneFromVpcEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveZoneFromVpcEndpointResponseBody) SetRequestId(v string) *RemoveZoneFromVpcEndpointResponseBody {
	s.RequestId = &v
	return s
}

type RemoveZoneFromVpcEndpointResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveZoneFromVpcEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveZoneFromVpcEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveZoneFromVpcEndpointResponse) GoString() string {
	return s.String()
}

func (s *RemoveZoneFromVpcEndpointResponse) SetHeaders(v map[string]*string) *RemoveZoneFromVpcEndpointResponse {
	s.Headers = v
	return s
}

func (s *RemoveZoneFromVpcEndpointResponse) SetBody(v *RemoveZoneFromVpcEndpointResponseBody) *RemoveZoneFromVpcEndpointResponse {
	s.Body = v
	return s
}

type UpdateVpcEndpointAttributeRequest struct {
	ClientToken         *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun              *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EndpointDescription *string `json:"EndpointDescription,omitempty" xml:"EndpointDescription,omitempty"`
	EndpointId          *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	EndpointName        *string `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateVpcEndpointAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateVpcEndpointAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateVpcEndpointAttributeRequest) SetClientToken(v string) *UpdateVpcEndpointAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateVpcEndpointAttributeRequest) SetDryRun(v bool) *UpdateVpcEndpointAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateVpcEndpointAttributeRequest) SetEndpointDescription(v string) *UpdateVpcEndpointAttributeRequest {
	s.EndpointDescription = &v
	return s
}

func (s *UpdateVpcEndpointAttributeRequest) SetEndpointId(v string) *UpdateVpcEndpointAttributeRequest {
	s.EndpointId = &v
	return s
}

func (s *UpdateVpcEndpointAttributeRequest) SetEndpointName(v string) *UpdateVpcEndpointAttributeRequest {
	s.EndpointName = &v
	return s
}

func (s *UpdateVpcEndpointAttributeRequest) SetRegionId(v string) *UpdateVpcEndpointAttributeRequest {
	s.RegionId = &v
	return s
}

type UpdateVpcEndpointAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateVpcEndpointAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateVpcEndpointAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVpcEndpointAttributeResponseBody) SetRequestId(v string) *UpdateVpcEndpointAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateVpcEndpointAttributeResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateVpcEndpointAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateVpcEndpointAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateVpcEndpointAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateVpcEndpointAttributeResponse) SetHeaders(v map[string]*string) *UpdateVpcEndpointAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateVpcEndpointAttributeResponse) SetBody(v *UpdateVpcEndpointAttributeResponseBody) *UpdateVpcEndpointAttributeResponse {
	s.Body = v
	return s
}

type UpdateVpcEndpointConnectionAttributeRequest struct {
	Bandwidth   *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EndpointId  *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceId   *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s UpdateVpcEndpointConnectionAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateVpcEndpointConnectionAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateVpcEndpointConnectionAttributeRequest) SetBandwidth(v int32) *UpdateVpcEndpointConnectionAttributeRequest {
	s.Bandwidth = &v
	return s
}

func (s *UpdateVpcEndpointConnectionAttributeRequest) SetClientToken(v string) *UpdateVpcEndpointConnectionAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateVpcEndpointConnectionAttributeRequest) SetDryRun(v bool) *UpdateVpcEndpointConnectionAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateVpcEndpointConnectionAttributeRequest) SetEndpointId(v string) *UpdateVpcEndpointConnectionAttributeRequest {
	s.EndpointId = &v
	return s
}

func (s *UpdateVpcEndpointConnectionAttributeRequest) SetRegionId(v string) *UpdateVpcEndpointConnectionAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateVpcEndpointConnectionAttributeRequest) SetServiceId(v string) *UpdateVpcEndpointConnectionAttributeRequest {
	s.ServiceId = &v
	return s
}

type UpdateVpcEndpointConnectionAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateVpcEndpointConnectionAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateVpcEndpointConnectionAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVpcEndpointConnectionAttributeResponseBody) SetRequestId(v string) *UpdateVpcEndpointConnectionAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateVpcEndpointConnectionAttributeResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateVpcEndpointConnectionAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateVpcEndpointConnectionAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateVpcEndpointConnectionAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateVpcEndpointConnectionAttributeResponse) SetHeaders(v map[string]*string) *UpdateVpcEndpointConnectionAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateVpcEndpointConnectionAttributeResponse) SetBody(v *UpdateVpcEndpointConnectionAttributeResponseBody) *UpdateVpcEndpointConnectionAttributeResponse {
	s.Body = v
	return s
}

type UpdateVpcEndpointServiceAttributeRequest struct {
	AutoAcceptEnabled   *bool   `json:"AutoAcceptEnabled,omitempty" xml:"AutoAcceptEnabled,omitempty"`
	ClientToken         *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConnectBandwidth    *int32  `json:"ConnectBandwidth,omitempty" xml:"ConnectBandwidth,omitempty"`
	DryRun              *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceDescription  *string `json:"ServiceDescription,omitempty" xml:"ServiceDescription,omitempty"`
	ServiceId           *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ZoneAffinityEnabled *bool   `json:"ZoneAffinityEnabled,omitempty" xml:"ZoneAffinityEnabled,omitempty"`
}

func (s UpdateVpcEndpointServiceAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateVpcEndpointServiceAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateVpcEndpointServiceAttributeRequest) SetAutoAcceptEnabled(v bool) *UpdateVpcEndpointServiceAttributeRequest {
	s.AutoAcceptEnabled = &v
	return s
}

func (s *UpdateVpcEndpointServiceAttributeRequest) SetClientToken(v string) *UpdateVpcEndpointServiceAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateVpcEndpointServiceAttributeRequest) SetConnectBandwidth(v int32) *UpdateVpcEndpointServiceAttributeRequest {
	s.ConnectBandwidth = &v
	return s
}

func (s *UpdateVpcEndpointServiceAttributeRequest) SetDryRun(v bool) *UpdateVpcEndpointServiceAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateVpcEndpointServiceAttributeRequest) SetRegionId(v string) *UpdateVpcEndpointServiceAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateVpcEndpointServiceAttributeRequest) SetServiceDescription(v string) *UpdateVpcEndpointServiceAttributeRequest {
	s.ServiceDescription = &v
	return s
}

func (s *UpdateVpcEndpointServiceAttributeRequest) SetServiceId(v string) *UpdateVpcEndpointServiceAttributeRequest {
	s.ServiceId = &v
	return s
}

func (s *UpdateVpcEndpointServiceAttributeRequest) SetZoneAffinityEnabled(v bool) *UpdateVpcEndpointServiceAttributeRequest {
	s.ZoneAffinityEnabled = &v
	return s
}

type UpdateVpcEndpointServiceAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateVpcEndpointServiceAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateVpcEndpointServiceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVpcEndpointServiceAttributeResponseBody) SetRequestId(v string) *UpdateVpcEndpointServiceAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateVpcEndpointServiceAttributeResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateVpcEndpointServiceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateVpcEndpointServiceAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateVpcEndpointServiceAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateVpcEndpointServiceAttributeResponse) SetHeaders(v map[string]*string) *UpdateVpcEndpointServiceAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateVpcEndpointServiceAttributeResponse) SetBody(v *UpdateVpcEndpointServiceAttributeResponseBody) *UpdateVpcEndpointServiceAttributeResponse {
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
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("privatelink"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddUserToVpcEndpointServiceWithOptions(request *AddUserToVpcEndpointServiceRequest, runtime *util.RuntimeOptions) (_result *AddUserToVpcEndpointServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["RegionId"] = request.RegionId
	query["RegionId"] = request.RegionId
	query["ServiceId"] = request.ServiceId
	query["UserId"] = request.UserId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddUserToVpcEndpointService"),
		Version:     tea.String("2020-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddUserToVpcEndpointServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddUserToVpcEndpointService(request *AddUserToVpcEndpointServiceRequest) (_result *AddUserToVpcEndpointServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddUserToVpcEndpointServiceResponse{}
	_body, _err := client.AddUserToVpcEndpointServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddZoneToVpcEndpointWithOptions(request *AddZoneToVpcEndpointRequest, runtime *util.RuntimeOptions) (_result *AddZoneToVpcEndpointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["EndpointId"] = request.EndpointId
	query["RegionId"] = request.RegionId
	query["RegionId"] = request.RegionId
	query["VSwitchId"] = request.VSwitchId
	query["ZoneId"] = request.ZoneId
	query["ip"] = request.Ip
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddZoneToVpcEndpoint"),
		Version:     tea.String("2020-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddZoneToVpcEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddZoneToVpcEndpoint(request *AddZoneToVpcEndpointRequest) (_result *AddZoneToVpcEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddZoneToVpcEndpointResponse{}
	_body, _err := client.AddZoneToVpcEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AttachResourceToVpcEndpointServiceWithOptions(request *AttachResourceToVpcEndpointServiceRequest, runtime *util.RuntimeOptions) (_result *AttachResourceToVpcEndpointServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["RegionId"] = request.RegionId
	query["RegionId"] = request.RegionId
	query["ResourceId"] = request.ResourceId
	query["ResourceType"] = request.ResourceType
	query["ServiceId"] = request.ServiceId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachResourceToVpcEndpointService"),
		Version:     tea.String("2020-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachResourceToVpcEndpointServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttachResourceToVpcEndpointService(request *AttachResourceToVpcEndpointServiceRequest) (_result *AttachResourceToVpcEndpointServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachResourceToVpcEndpointServiceResponse{}
	_body, _err := client.AttachResourceToVpcEndpointServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AttachSecurityGroupToVpcEndpointWithOptions(request *AttachSecurityGroupToVpcEndpointRequest, runtime *util.RuntimeOptions) (_result *AttachSecurityGroupToVpcEndpointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["EndpointId"] = request.EndpointId
	query["RegionId"] = request.RegionId
	query["RegionId"] = request.RegionId
	query["SecurityGroupId"] = request.SecurityGroupId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachSecurityGroupToVpcEndpoint"),
		Version:     tea.String("2020-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachSecurityGroupToVpcEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttachSecurityGroupToVpcEndpoint(request *AttachSecurityGroupToVpcEndpointRequest) (_result *AttachSecurityGroupToVpcEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachSecurityGroupToVpcEndpointResponse{}
	_body, _err := client.AttachSecurityGroupToVpcEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckProductOpenWithOptions(runtime *util.RuntimeOptions) (_result *CheckProductOpenResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("CheckProductOpen"),
		Version:     tea.String("2020-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckProductOpenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckProductOpen() (_result *CheckProductOpenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckProductOpenResponse{}
	_body, _err := client.CheckProductOpenWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVpcEndpointWithOptions(request *CreateVpcEndpointRequest, runtime *util.RuntimeOptions) (_result *CreateVpcEndpointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["EndpointDescription"] = request.EndpointDescription
	query["EndpointName"] = request.EndpointName
	query["EndpointType"] = request.EndpointType
	query["ProtectedEnabled"] = request.ProtectedEnabled
	query["RegionId"] = request.RegionId
	query["RegionId"] = request.RegionId
	query["SecurityGroupId"] = request.SecurityGroupId
	query["ServiceId"] = request.ServiceId
	query["ServiceName"] = request.ServiceName
	query["VpcId"] = request.VpcId
	query["Zone"] = request.Zone
	query["ZonePrivateIpAddressCount"] = request.ZonePrivateIpAddressCount
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateVpcEndpoint"),
		Version:     tea.String("2020-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateVpcEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVpcEndpoint(request *CreateVpcEndpointRequest) (_result *CreateVpcEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVpcEndpointResponse{}
	_body, _err := client.CreateVpcEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVpcEndpointServiceWithOptions(request *CreateVpcEndpointServiceRequest, runtime *util.RuntimeOptions) (_result *CreateVpcEndpointServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AutoAcceptEnabled"] = request.AutoAcceptEnabled
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["Payer"] = request.Payer
	query["RegionId"] = request.RegionId
	query["RegionId"] = request.RegionId
	query["Resource"] = request.Resource
	query["ServiceDescription"] = request.ServiceDescription
	query["ServiceResourceType"] = request.ServiceResourceType
	query["ZoneAffinityEnabled"] = request.ZoneAffinityEnabled
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateVpcEndpointService"),
		Version:     tea.String("2020-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateVpcEndpointServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVpcEndpointService(request *CreateVpcEndpointServiceRequest) (_result *CreateVpcEndpointServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVpcEndpointServiceResponse{}
	_body, _err := client.CreateVpcEndpointServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVpcEndpointWithOptions(request *DeleteVpcEndpointRequest, runtime *util.RuntimeOptions) (_result *DeleteVpcEndpointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["EndpointId"] = request.EndpointId
	query["RegionId"] = request.RegionId
	query["RegionId"] = request.RegionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVpcEndpoint"),
		Version:     tea.String("2020-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVpcEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVpcEndpoint(request *DeleteVpcEndpointRequest) (_result *DeleteVpcEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVpcEndpointResponse{}
	_body, _err := client.DeleteVpcEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVpcEndpointServiceWithOptions(request *DeleteVpcEndpointServiceRequest, runtime *util.RuntimeOptions) (_result *DeleteVpcEndpointServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["RegionId"] = request.RegionId
	query["RegionId"] = request.RegionId
	query["ServiceId"] = request.ServiceId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVpcEndpointService"),
		Version:     tea.String("2020-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVpcEndpointServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVpcEndpointService(request *DeleteVpcEndpointServiceRequest) (_result *DeleteVpcEndpointServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVpcEndpointServiceResponse{}
	_body, _err := client.DeleteVpcEndpointServiceWithOptions(request, runtime)
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
	query["RegionId"] = request.RegionId
	query["RegionId"] = request.RegionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2020-04-15"),
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

func (client *Client) DescribeZonesWithOptions(request *DescribeZonesRequest, runtime *util.RuntimeOptions) (_result *DescribeZonesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["RegionId"] = request.RegionId
	query["RegionId"] = request.RegionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeZones"),
		Version:     tea.String("2020-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeZones(request *DescribeZonesRequest) (_result *DescribeZonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeZonesResponse{}
	_body, _err := client.DescribeZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetachResourceFromVpcEndpointServiceWithOptions(request *DetachResourceFromVpcEndpointServiceRequest, runtime *util.RuntimeOptions) (_result *DetachResourceFromVpcEndpointServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["RegionId"] = request.RegionId
	query["RegionId"] = request.RegionId
	query["ResourceId"] = request.ResourceId
	query["ResourceType"] = request.ResourceType
	query["ServiceId"] = request.ServiceId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachResourceFromVpcEndpointService"),
		Version:     tea.String("2020-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachResourceFromVpcEndpointServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetachResourceFromVpcEndpointService(request *DetachResourceFromVpcEndpointServiceRequest) (_result *DetachResourceFromVpcEndpointServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachResourceFromVpcEndpointServiceResponse{}
	_body, _err := client.DetachResourceFromVpcEndpointServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetachSecurityGroupFromVpcEndpointWithOptions(request *DetachSecurityGroupFromVpcEndpointRequest, runtime *util.RuntimeOptions) (_result *DetachSecurityGroupFromVpcEndpointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["EndpointId"] = request.EndpointId
	query["RegionId"] = request.RegionId
	query["RegionId"] = request.RegionId
	query["SecurityGroupId"] = request.SecurityGroupId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachSecurityGroupFromVpcEndpoint"),
		Version:     tea.String("2020-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachSecurityGroupFromVpcEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetachSecurityGroupFromVpcEndpoint(request *DetachSecurityGroupFromVpcEndpointRequest) (_result *DetachSecurityGroupFromVpcEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachSecurityGroupFromVpcEndpointResponse{}
	_body, _err := client.DetachSecurityGroupFromVpcEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableVpcEndpointConnectionWithOptions(request *DisableVpcEndpointConnectionRequest, runtime *util.RuntimeOptions) (_result *DisableVpcEndpointConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["EndpointId"] = request.EndpointId
	query["RegionId"] = request.RegionId
	query["RegionId"] = request.RegionId
	query["ServiceId"] = request.ServiceId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableVpcEndpointConnection"),
		Version:     tea.String("2020-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableVpcEndpointConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableVpcEndpointConnection(request *DisableVpcEndpointConnectionRequest) (_result *DisableVpcEndpointConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableVpcEndpointConnectionResponse{}
	_body, _err := client.DisableVpcEndpointConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableVpcEndpointConnectionWithOptions(request *EnableVpcEndpointConnectionRequest, runtime *util.RuntimeOptions) (_result *EnableVpcEndpointConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Bandwidth"] = request.Bandwidth
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["EndpointId"] = request.EndpointId
	query["RegionId"] = request.RegionId
	query["RegionId"] = request.RegionId
	query["ServiceId"] = request.ServiceId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableVpcEndpointConnection"),
		Version:     tea.String("2020-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableVpcEndpointConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableVpcEndpointConnection(request *EnableVpcEndpointConnectionRequest) (_result *EnableVpcEndpointConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableVpcEndpointConnectionResponse{}
	_body, _err := client.EnableVpcEndpointConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetVpcEndpointAttributeWithOptions(request *GetVpcEndpointAttributeRequest, runtime *util.RuntimeOptions) (_result *GetVpcEndpointAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["EndpointId"] = request.EndpointId
	query["RegionId"] = request.RegionId
	query["RegionId"] = request.RegionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetVpcEndpointAttribute"),
		Version:     tea.String("2020-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetVpcEndpointAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetVpcEndpointAttribute(request *GetVpcEndpointAttributeRequest) (_result *GetVpcEndpointAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVpcEndpointAttributeResponse{}
	_body, _err := client.GetVpcEndpointAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetVpcEndpointServiceAttributeWithOptions(request *GetVpcEndpointServiceAttributeRequest, runtime *util.RuntimeOptions) (_result *GetVpcEndpointServiceAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["RegionId"] = request.RegionId
	query["RegionId"] = request.RegionId
	query["ServiceId"] = request.ServiceId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetVpcEndpointServiceAttribute"),
		Version:     tea.String("2020-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetVpcEndpointServiceAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetVpcEndpointServiceAttribute(request *GetVpcEndpointServiceAttributeRequest) (_result *GetVpcEndpointServiceAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVpcEndpointServiceAttributeResponse{}
	_body, _err := client.GetVpcEndpointServiceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVpcEndpointConnectionsWithOptions(request *ListVpcEndpointConnectionsRequest, runtime *util.RuntimeOptions) (_result *ListVpcEndpointConnectionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ConnectionStatus"] = request.ConnectionStatus
	query["EndpointId"] = request.EndpointId
	query["EndpointOwnerId"] = request.EndpointOwnerId
	query["EniId"] = request.EniId
	query["MaxResults"] = request.MaxResults
	query["NextToken"] = request.NextToken
	query["RegionId"] = request.RegionId
	query["RegionId"] = request.RegionId
	query["ServiceId"] = request.ServiceId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListVpcEndpointConnections"),
		Version:     tea.String("2020-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListVpcEndpointConnectionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVpcEndpointConnections(request *ListVpcEndpointConnectionsRequest) (_result *ListVpcEndpointConnectionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVpcEndpointConnectionsResponse{}
	_body, _err := client.ListVpcEndpointConnectionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVpcEndpointSecurityGroupsWithOptions(request *ListVpcEndpointSecurityGroupsRequest, runtime *util.RuntimeOptions) (_result *ListVpcEndpointSecurityGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["EndpointId"] = request.EndpointId
	query["MaxResults"] = request.MaxResults
	query["NextToken"] = request.NextToken
	query["RegionId"] = request.RegionId
	query["RegionId"] = request.RegionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListVpcEndpointSecurityGroups"),
		Version:     tea.String("2020-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListVpcEndpointSecurityGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVpcEndpointSecurityGroups(request *ListVpcEndpointSecurityGroupsRequest) (_result *ListVpcEndpointSecurityGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVpcEndpointSecurityGroupsResponse{}
	_body, _err := client.ListVpcEndpointSecurityGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVpcEndpointServiceResourcesWithOptions(request *ListVpcEndpointServiceResourcesRequest, runtime *util.RuntimeOptions) (_result *ListVpcEndpointServiceResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["MaxResults"] = request.MaxResults
	query["NextToken"] = request.NextToken
	query["RegionId"] = request.RegionId
	query["RegionId"] = request.RegionId
	query["ServiceId"] = request.ServiceId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListVpcEndpointServiceResources"),
		Version:     tea.String("2020-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListVpcEndpointServiceResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVpcEndpointServiceResources(request *ListVpcEndpointServiceResourcesRequest) (_result *ListVpcEndpointServiceResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVpcEndpointServiceResourcesResponse{}
	_body, _err := client.ListVpcEndpointServiceResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVpcEndpointServiceUsersWithOptions(request *ListVpcEndpointServiceUsersRequest, runtime *util.RuntimeOptions) (_result *ListVpcEndpointServiceUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["MaxResults"] = request.MaxResults
	query["NextToken"] = request.NextToken
	query["RegionId"] = request.RegionId
	query["RegionId"] = request.RegionId
	query["ServiceId"] = request.ServiceId
	query["UserId"] = request.UserId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListVpcEndpointServiceUsers"),
		Version:     tea.String("2020-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListVpcEndpointServiceUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVpcEndpointServiceUsers(request *ListVpcEndpointServiceUsersRequest) (_result *ListVpcEndpointServiceUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVpcEndpointServiceUsersResponse{}
	_body, _err := client.ListVpcEndpointServiceUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVpcEndpointServicesWithOptions(request *ListVpcEndpointServicesRequest, runtime *util.RuntimeOptions) (_result *ListVpcEndpointServicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AutoAcceptEnabled"] = request.AutoAcceptEnabled
	query["MaxResults"] = request.MaxResults
	query["NextToken"] = request.NextToken
	query["RegionId"] = request.RegionId
	query["RegionId"] = request.RegionId
	query["ServiceBusinessStatus"] = request.ServiceBusinessStatus
	query["ServiceId"] = request.ServiceId
	query["ServiceName"] = request.ServiceName
	query["ServiceResourceType"] = request.ServiceResourceType
	query["ServiceStatus"] = request.ServiceStatus
	query["ZoneAffinityEnabled"] = request.ZoneAffinityEnabled
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListVpcEndpointServices"),
		Version:     tea.String("2020-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListVpcEndpointServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVpcEndpointServices(request *ListVpcEndpointServicesRequest) (_result *ListVpcEndpointServicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVpcEndpointServicesResponse{}
	_body, _err := client.ListVpcEndpointServicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVpcEndpointServicesByEndUserWithOptions(request *ListVpcEndpointServicesByEndUserRequest, runtime *util.RuntimeOptions) (_result *ListVpcEndpointServicesByEndUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["MaxResults"] = request.MaxResults
	query["NextToken"] = request.NextToken
	query["RegionId"] = request.RegionId
	query["RegionId"] = request.RegionId
	query["ServiceId"] = request.ServiceId
	query["ServiceName"] = request.ServiceName
	query["ServiceType"] = request.ServiceType
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListVpcEndpointServicesByEndUser"),
		Version:     tea.String("2020-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListVpcEndpointServicesByEndUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVpcEndpointServicesByEndUser(request *ListVpcEndpointServicesByEndUserRequest) (_result *ListVpcEndpointServicesByEndUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVpcEndpointServicesByEndUserResponse{}
	_body, _err := client.ListVpcEndpointServicesByEndUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVpcEndpointZonesWithOptions(request *ListVpcEndpointZonesRequest, runtime *util.RuntimeOptions) (_result *ListVpcEndpointZonesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["EndpointId"] = request.EndpointId
	query["MaxResults"] = request.MaxResults
	query["NextToken"] = request.NextToken
	query["RegionId"] = request.RegionId
	query["RegionId"] = request.RegionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListVpcEndpointZones"),
		Version:     tea.String("2020-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListVpcEndpointZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVpcEndpointZones(request *ListVpcEndpointZonesRequest) (_result *ListVpcEndpointZonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVpcEndpointZonesResponse{}
	_body, _err := client.ListVpcEndpointZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVpcEndpointsWithOptions(request *ListVpcEndpointsRequest, runtime *util.RuntimeOptions) (_result *ListVpcEndpointsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ConnectionStatus"] = request.ConnectionStatus
	query["EndpointId"] = request.EndpointId
	query["EndpointName"] = request.EndpointName
	query["EndpointStatus"] = request.EndpointStatus
	query["EndpointType"] = request.EndpointType
	query["MaxResults"] = request.MaxResults
	query["NextToken"] = request.NextToken
	query["RegionId"] = request.RegionId
	query["RegionId"] = request.RegionId
	query["ServiceName"] = request.ServiceName
	query["VpcId"] = request.VpcId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListVpcEndpoints"),
		Version:     tea.String("2020-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListVpcEndpointsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVpcEndpoints(request *ListVpcEndpointsRequest) (_result *ListVpcEndpointsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVpcEndpointsResponse{}
	_body, _err := client.ListVpcEndpointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenPrivateLinkServiceWithOptions(request *OpenPrivateLinkServiceRequest, runtime *util.RuntimeOptions) (_result *OpenPrivateLinkServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["OwnerId"] = request.OwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OpenPrivateLinkService"),
		Version:     tea.String("2020-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenPrivateLinkServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenPrivateLinkService(request *OpenPrivateLinkServiceRequest) (_result *OpenPrivateLinkServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenPrivateLinkServiceResponse{}
	_body, _err := client.OpenPrivateLinkServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveUserFromVpcEndpointServiceWithOptions(request *RemoveUserFromVpcEndpointServiceRequest, runtime *util.RuntimeOptions) (_result *RemoveUserFromVpcEndpointServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["RegionId"] = request.RegionId
	query["RegionId"] = request.RegionId
	query["ServiceId"] = request.ServiceId
	query["UserId"] = request.UserId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveUserFromVpcEndpointService"),
		Version:     tea.String("2020-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveUserFromVpcEndpointServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveUserFromVpcEndpointService(request *RemoveUserFromVpcEndpointServiceRequest) (_result *RemoveUserFromVpcEndpointServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveUserFromVpcEndpointServiceResponse{}
	_body, _err := client.RemoveUserFromVpcEndpointServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveZoneFromVpcEndpointWithOptions(request *RemoveZoneFromVpcEndpointRequest, runtime *util.RuntimeOptions) (_result *RemoveZoneFromVpcEndpointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["EndpointId"] = request.EndpointId
	query["RegionId"] = request.RegionId
	query["RegionId"] = request.RegionId
	query["ZoneId"] = request.ZoneId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveZoneFromVpcEndpoint"),
		Version:     tea.String("2020-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveZoneFromVpcEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveZoneFromVpcEndpoint(request *RemoveZoneFromVpcEndpointRequest) (_result *RemoveZoneFromVpcEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveZoneFromVpcEndpointResponse{}
	_body, _err := client.RemoveZoneFromVpcEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateVpcEndpointAttributeWithOptions(request *UpdateVpcEndpointAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateVpcEndpointAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["EndpointDescription"] = request.EndpointDescription
	query["EndpointId"] = request.EndpointId
	query["EndpointName"] = request.EndpointName
	query["RegionId"] = request.RegionId
	query["RegionId"] = request.RegionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateVpcEndpointAttribute"),
		Version:     tea.String("2020-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateVpcEndpointAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateVpcEndpointAttribute(request *UpdateVpcEndpointAttributeRequest) (_result *UpdateVpcEndpointAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateVpcEndpointAttributeResponse{}
	_body, _err := client.UpdateVpcEndpointAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateVpcEndpointConnectionAttributeWithOptions(request *UpdateVpcEndpointConnectionAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateVpcEndpointConnectionAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Bandwidth"] = request.Bandwidth
	query["ClientToken"] = request.ClientToken
	query["DryRun"] = request.DryRun
	query["EndpointId"] = request.EndpointId
	query["RegionId"] = request.RegionId
	query["RegionId"] = request.RegionId
	query["ServiceId"] = request.ServiceId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateVpcEndpointConnectionAttribute"),
		Version:     tea.String("2020-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateVpcEndpointConnectionAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateVpcEndpointConnectionAttribute(request *UpdateVpcEndpointConnectionAttributeRequest) (_result *UpdateVpcEndpointConnectionAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateVpcEndpointConnectionAttributeResponse{}
	_body, _err := client.UpdateVpcEndpointConnectionAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateVpcEndpointServiceAttributeWithOptions(request *UpdateVpcEndpointServiceAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateVpcEndpointServiceAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AutoAcceptEnabled"] = request.AutoAcceptEnabled
	query["ClientToken"] = request.ClientToken
	query["ConnectBandwidth"] = request.ConnectBandwidth
	query["DryRun"] = request.DryRun
	query["RegionId"] = request.RegionId
	query["RegionId"] = request.RegionId
	query["ServiceDescription"] = request.ServiceDescription
	query["ServiceId"] = request.ServiceId
	query["ZoneAffinityEnabled"] = request.ZoneAffinityEnabled
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateVpcEndpointServiceAttribute"),
		Version:     tea.String("2020-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateVpcEndpointServiceAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateVpcEndpointServiceAttribute(request *UpdateVpcEndpointServiceAttributeRequest) (_result *UpdateVpcEndpointServiceAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateVpcEndpointServiceAttributeResponse{}
	_body, _err := client.UpdateVpcEndpointServiceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
