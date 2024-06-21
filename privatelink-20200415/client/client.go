// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddUserToVpcEndpointServiceRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The region ID of the endpoint service. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// eu-west-1
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The endpoint service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The whitelist in the format of Aliyun Resource Name (ARN).
	//
	// example:
	//
	// acs:ram:*:<account-id>:*
	UserARN *string `json:"UserARN,omitempty" xml:"UserARN,omitempty"`
	// The account ID that you want to add to the whitelist.
	//
	// example:
	//
	// 132193271328****
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
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

func (s *AddUserToVpcEndpointServiceRequest) SetUserARN(v string) *AddUserToVpcEndpointServiceRequest {
	s.UserARN = &v
	return s
}

func (s *AddUserToVpcEndpointServiceRequest) SetUserId(v int64) *AddUserToVpcEndpointServiceRequest {
	s.UserId = &v
	return s
}

type AddUserToVpcEndpointServiceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddUserToVpcEndpointServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *AddUserToVpcEndpointServiceResponse) SetStatusCode(v int32) *AddUserToVpcEndpointServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *AddUserToVpcEndpointServiceResponse) SetBody(v *AddUserToVpcEndpointServiceResponseBody) *AddUserToVpcEndpointServiceResponse {
	s.Body = v
	return s
}

type AddZoneToVpcEndpointRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the endpoint to which you want to add the zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// ep-hp33b2e43fays7s8****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The region ID of the endpoint.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the vSwitch in the zone that you want to add. The system automatically creates an endpoint ENI in the vSwitch.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-hjkshjvdkdvd****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the zone that you want to add.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The IP address of the endpoint elastic network interface (ENI) in the zone that you want to add.
	//
	// example:
	//
	// 192.XX.XX.32
	Ip *string `json:"ip,omitempty" xml:"ip,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddZoneToVpcEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *AddZoneToVpcEndpointResponse) SetStatusCode(v int32) *AddZoneToVpcEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *AddZoneToVpcEndpointResponse) SetBody(v *AddZoneToVpcEndpointResponseBody) *AddZoneToVpcEndpointResponse {
	s.Body = v
	return s
}

type AttachResourceToVpcEndpointServiceRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The region ID of the endpoint service to which you want to add the service resource.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service resource ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-hp32z1wp5peaoox2q****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the service resource. Valid values:
	//
	// 	- **slb**: a Classic Load Balancer (CLB) instance
	//
	// 	- **alb**: an Application Load Balancer (ALB) instance
	//
	// 	- **nlb**: a Network Load Balancer (NLB) instance
	//
	// This parameter is required.
	//
	// example:
	//
	// slb
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The ID of the endpoint service to which you want to add the service resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The zone ID of the service resource.
	//
	// example:
	//
	// cn-hangzhou-j
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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

func (s *AttachResourceToVpcEndpointServiceRequest) SetZoneId(v string) *AttachResourceToVpcEndpointServiceRequest {
	s.ZoneId = &v
	return s
}

type AttachResourceToVpcEndpointServiceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
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
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachResourceToVpcEndpointServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *AttachResourceToVpcEndpointServiceResponse) SetStatusCode(v int32) *AttachResourceToVpcEndpointServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachResourceToVpcEndpointServiceResponse) SetBody(v *AttachResourceToVpcEndpointServiceResponseBody) *AttachResourceToVpcEndpointServiceResponse {
	s.Body = v
	return s
}

type AttachSecurityGroupToVpcEndpointRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the endpoint with which you want to associate the security group.
	//
	// This parameter is required.
	//
	// example:
	//
	// ep-hp33b2e43fays7s8****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The region ID of the endpoint with which you want to associate with the security group. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// eu-west-1
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the security group with which you want to associate the endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// sg-hp3c8qj1tyct90ej****
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
	// The request ID.
	//
	// example:
	//
	// 8D778FF9-7640-4C13-BCD6-9265CA9A2F81
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
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachSecurityGroupToVpcEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *AttachSecurityGroupToVpcEndpointResponse) SetStatusCode(v int32) *AttachSecurityGroupToVpcEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachSecurityGroupToVpcEndpointResponse) SetBody(v *AttachSecurityGroupToVpcEndpointResponseBody) *AttachSecurityGroupToVpcEndpointResponse {
	s.Body = v
	return s
}

type ChangeResourceGroupRequest struct {
	// The resource group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource IDs. You can specify up to 50 resource IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ep-hp3i05294c2d2d****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of resource. Valid values:
	//
	// 	- **VpcEndpoint**: endpoint
	//
	// 	- **VpcEndpointService**: endpoint service
	//
	// example:
	//
	// VpcEndpoint
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ChangeResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupRequest) SetResourceGroupId(v string) *ChangeResourceGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceId(v string) *ChangeResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceType(v string) *ChangeResourceGroupRequest {
	s.ResourceType = &v
	return s
}

type ChangeResourceGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8D8992C1-6712-423C-BAC5-E5E817484C6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponseBody) SetRequestId(v string) *ChangeResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type ChangeResourceGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponse) SetHeaders(v map[string]*string) *ChangeResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *ChangeResourceGroupResponse) SetStatusCode(v int32) *ChangeResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeResourceGroupResponse) SetBody(v *ChangeResourceGroupResponseBody) *ChangeResourceGroupResponse {
	s.Body = v
	return s
}

type CheckProductOpenResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8D8992C1-6712-423C-BAC5-E5E817484C6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether PrivateLink is activated.
	//
	// Only **true*	- is returned. The value indicates that PrivateLink is activated.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckProductOpenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CheckProductOpenResponse) SetStatusCode(v int32) *CheckProductOpenResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckProductOpenResponse) SetBody(v *CheckProductOpenResponseBody) *CheckProductOpenResponse {
	s.Body = v
	return s
}

type CreateVpcEndpointRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The description of the endpoint.
	//
	// The description must be 2 to 256 characters in length, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// This is my Endpoint.
	EndpointDescription *string `json:"EndpointDescription,omitempty" xml:"EndpointDescription,omitempty"`
	// The name of the endpoint.
	//
	// The name must be 2 to 128 characters in length, and can contain digits, underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// test
	EndpointName *string `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
	// The type of the endpoint.
	//
	// Set the value to **Interface**. Then, you can specify Application Load Balancer (ALB) and Classic Load Balancer (CLB) instances as service resources for the endpoint service.
	//
	// example:
	//
	// Interface
	EndpointType   *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	// Specifies whether to enable user authentication. This parameter is available in Security Token Service (STS) mode. Valid values:
	//
	// 	- **true**: enables user authentication. After user authentication is enabled, only the user who creates the endpoint can modify or delete the endpoint in STS mode.
	//
	// 	- **false*	- (default): disables user authentication.
	//
	// example:
	//
	// false
	ProtectedEnabled *bool `json:"ProtectedEnabled,omitempty" xml:"ProtectedEnabled,omitempty"`
	// The region ID of the endpoint.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The IDs of security groups that are associated with the endpoint elastic network interface (ENI).
	//
	// example:
	//
	// sg-hp33bw6ynvm2yb0e****
	SecurityGroupId []*string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" type:"Repeated"`
	// The ID of the endpoint service with which the endpoint is associated.
	//
	// example:
	//
	// epsrv-hp3xdsq46ael67lo****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The name of the endpoint service with which the endpoint is associated.
	//
	// example:
	//
	// com.aliyuncs.privatelink.cn-huhehaote.epsrv-hp3vpx8yqxblby3i****
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The tags to add to the resource.
	Tag []*CreateVpcEndpointRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC) to which the endpoint belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-hp356stwkxg3fn2xe****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zones where the endpoint is deployed.
	Zone []*CreateVpcEndpointRequestZone `json:"Zone,omitempty" xml:"Zone,omitempty" type:"Repeated"`
	// The number of private IP addresses that are assigned to an elastic network interface (ENI) in each zone. Set the value to **1**.
	//
	// example:
	//
	// 1
	ZonePrivateIpAddressCount *int64 `json:"ZonePrivateIpAddressCount,omitempty" xml:"ZonePrivateIpAddressCount,omitempty"`
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

func (s *CreateVpcEndpointRequest) SetPolicyDocument(v string) *CreateVpcEndpointRequest {
	s.PolicyDocument = &v
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

func (s *CreateVpcEndpointRequest) SetResourceGroupId(v string) *CreateVpcEndpointRequest {
	s.ResourceGroupId = &v
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

func (s *CreateVpcEndpointRequest) SetTag(v []*CreateVpcEndpointRequestTag) *CreateVpcEndpointRequest {
	s.Tag = v
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

type CreateVpcEndpointRequestTag struct {
	// The key of the tag to add to the resource.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag to add to the resource.
	//
	// example:
	//
	// prod
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateVpcEndpointRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcEndpointRequestTag) GoString() string {
	return s.String()
}

func (s *CreateVpcEndpointRequestTag) SetKey(v string) *CreateVpcEndpointRequestTag {
	s.Key = &v
	return s
}

func (s *CreateVpcEndpointRequestTag) SetValue(v string) *CreateVpcEndpointRequestTag {
	s.Value = &v
	return s
}

type CreateVpcEndpointRequestZone struct {
	// The ID of the vSwitch where you want to create the endpoint ENI in the zone. You can specify up to 10 vSwitch IDs.
	//
	// example:
	//
	// vsw-hp3uf6045ljdhd5zr****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the zone in which the endpoint is deployed.
	//
	// You can specify up to 10 zone IDs.
	//
	// example:
	//
	// cn-huhehaote-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The IP address of the zone in which the endpoint is deployed.
	//
	// You can specify up to 10 IP addresses.
	//
	// example:
	//
	// 192.168.XX.XX
	Ip *string `json:"ip,omitempty" xml:"ip,omitempty"`
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
	// The bandwidth of the endpoint connection. Unit: Mbit/s.
	//
	// example:
	//
	// 200
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The state of the endpoint connection. Valid values:
	//
	// 	- **Pending**: The connection is being modified.
	//
	// 	- **Connecting**: The connection is being established.
	//
	// 	- **Connected**: The connection is established.
	//
	// 	- **Disconnecting**: The endpoint is being disconnected from the endpoint service.
	//
	// 	- **Disconnected**: The endpoint is disconnected from the endpoint service.
	//
	// 	- **Deleting**: The connection is being deleted.
	//
	// example:
	//
	// Disconnected
	ConnectionStatus *string `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	// The time when the endpoint was created.
	//
	// example:
	//
	// 2022-01-02T19:11:12Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The service state of the endpoint. Valid values:
	//
	// 	- **Normal**: The endpoint runs as expected.
	//
	// 	- **FinancialLocked**: The endpoint is locked due to overdue payments.
	//
	// example:
	//
	// Normal
	EndpointBusinessStatus *string `json:"EndpointBusinessStatus,omitempty" xml:"EndpointBusinessStatus,omitempty"`
	// The description of the endpoint.
	//
	// example:
	//
	// This is my Endpoint.
	EndpointDescription *string `json:"EndpointDescription,omitempty" xml:"EndpointDescription,omitempty"`
	// The domain name of the endpoint.
	//
	// example:
	//
	// ep-hp33b2e43fays7s8****.epsrv-hp3xdsq46ael67lo****.cn-huhehaote.privatelink.aliyuncs.com
	EndpointDomain *string `json:"EndpointDomain,omitempty" xml:"EndpointDomain,omitempty"`
	// The endpoint ID.
	//
	// example:
	//
	// ep-hp33b2e43fays7s8****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The name of the endpoint.
	//
	// example:
	//
	// test
	EndpointName *string `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
	// The state of the endpoint. Valid values:
	//
	// 	- **Creating**: The endpoint is being created.
	//
	// 	- **Active**: The endpoint is available.
	//
	// 	- **Pending**: The endpoint is being modified.
	//
	// 	- **Deleting**: The endpoint is being deleted.
	//
	// example:
	//
	// Active
	EndpointStatus *string `json:"EndpointStatus,omitempty" xml:"EndpointStatus,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the endpoint service with which the endpoint is associated.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The name of the endpoint service with which the endpoint is associated.
	//
	// example:
	//
	// com.aliyuncs.privatelink.cn-huhehaote.epsrv-hp3xdsq46ael67lo****
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The ID of the VPC to which the endpoint belongs.
	//
	// example:
	//
	// vpc-hp356stwkxg3fn2xe****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVpcEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateVpcEndpointResponse) SetStatusCode(v int32) *CreateVpcEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVpcEndpointResponse) SetBody(v *CreateVpcEndpointResponseBody) *CreateVpcEndpointResponse {
	s.Body = v
	return s
}

type CreateVpcEndpointServiceRequest struct {
	// Specifies whether to automatically accept endpoint connection requests. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	AutoAcceptEnabled *bool `json:"AutoAcceptEnabled,omitempty" xml:"AutoAcceptEnabled,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request.
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The payer of the endpoint service. Valid values:
	//
	// 	- **Endpoint**: the service consumer
	//
	// 	- **EndpointService**: the service provider
	//
	// > By default, the feature of allowing the service provider to pay is unavailable. To use this feature, log on to the [Quota Center console](https://quotas.console.aliyun.com/white-list-products/privatelink/quotas) and click Privileges in the left-side navigation pane. On the **Privileges*	- page, enter the quota ID `privatelink_whitelist/epsvc_payer_mode`, and click Apply in the Actions column.
	//
	// example:
	//
	// Endpoint
	Payer *string `json:"Payer,omitempty" xml:"Payer,omitempty"`
	// The region ID of the endpoint service.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service resources of the endpoint service.
	Resource []*CreateVpcEndpointServiceRequestResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Repeated"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The description of the endpoint service.
	//
	// example:
	//
	// This is my EndpointService.
	ServiceDescription *string `json:"ServiceDescription,omitempty" xml:"ServiceDescription,omitempty"`
	// The type of the service resource. Valid values:
	//
	// 	- **slb**: a Classic Load Balancer (CLB) instance
	//
	// 	- **alb**: an Application Load Balancer (ALB) instance
	//
	// 	- **nlb**: a Network Load Balancer (NLB) instance
	//
	// example:
	//
	// slb
	ServiceResourceType *string `json:"ServiceResourceType,omitempty" xml:"ServiceResourceType,omitempty"`
	// Specifies whether to enable IPv6 for the endpoint service. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	ServiceSupportIPv6 *bool `json:"ServiceSupportIPv6,omitempty" xml:"ServiceSupportIPv6,omitempty"`
	// The tags to add to the resource.
	Tag []*CreateVpcEndpointServiceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Specifies whether to first resolve the domain name of the nearest endpoint that is associated with the endpoint service. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	ZoneAffinityEnabled *bool `json:"ZoneAffinityEnabled,omitempty" xml:"ZoneAffinityEnabled,omitempty"`
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

func (s *CreateVpcEndpointServiceRequest) SetResourceGroupId(v string) *CreateVpcEndpointServiceRequest {
	s.ResourceGroupId = &v
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

func (s *CreateVpcEndpointServiceRequest) SetServiceSupportIPv6(v bool) *CreateVpcEndpointServiceRequest {
	s.ServiceSupportIPv6 = &v
	return s
}

func (s *CreateVpcEndpointServiceRequest) SetTag(v []*CreateVpcEndpointServiceRequestTag) *CreateVpcEndpointServiceRequest {
	s.Tag = v
	return s
}

func (s *CreateVpcEndpointServiceRequest) SetZoneAffinityEnabled(v bool) *CreateVpcEndpointServiceRequest {
	s.ZoneAffinityEnabled = &v
	return s
}

type CreateVpcEndpointServiceRequestResource struct {
	// The ID of the service resource that is added to the endpoint service. You can specify up to 20 service resource IDs.
	//
	// example:
	//
	// lb-hp32z1wp5peaoox2q****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the service resource that is added to the endpoint service. You can add up to 20 service resources to the endpoint service. Valid values:
	//
	// 	- **slb**: Classic Load Balancer (CLB) instance
	//
	// 	- **alb**: Application Load Balancer (ALB) instance
	//
	// 	- **nlb**: Network Load Balancer (NLB) instance
	//
	// >  In regions where PrivateLink is supported, CLB instances deployed in virtual private clouds (VPCs) can serve as the service resources of the endpoint service.
	//
	// example:
	//
	// slb
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-huhehaote-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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

func (s *CreateVpcEndpointServiceRequestResource) SetZoneId(v string) *CreateVpcEndpointServiceRequestResource {
	s.ZoneId = &v
	return s
}

type CreateVpcEndpointServiceRequestTag struct {
	// The key of the tag to add to the resource. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `aliyun` or `acs:`.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag to add to the resource. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag value cannot start with `aliyun` or `acs:`.
	//
	// example:
	//
	// prod
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateVpcEndpointServiceRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcEndpointServiceRequestTag) GoString() string {
	return s.String()
}

func (s *CreateVpcEndpointServiceRequestTag) SetKey(v string) *CreateVpcEndpointServiceRequestTag {
	s.Key = &v
	return s
}

func (s *CreateVpcEndpointServiceRequestTag) SetValue(v string) *CreateVpcEndpointServiceRequestTag {
	s.Value = &v
	return s
}

type CreateVpcEndpointServiceResponseBody struct {
	// Indicates whether the endpoint service automatically accepts endpoint connection requests. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	AutoAcceptEnabled *bool `json:"AutoAcceptEnabled,omitempty" xml:"AutoAcceptEnabled,omitempty"`
	// The time when the endpoint service was created.
	//
	// example:
	//
	// 2022-01-02T19:11:12Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8D8992C1-6712-423C-BAC5-E5E817484C6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The service state of the endpoint service. Valid values:
	//
	// 	- **Normal**: The endpoint service runs as expected.
	//
	// 	- **FinancialLocked**: The endpoint service is locked due to overdue payments.
	//
	// example:
	//
	// Normal
	ServiceBusinessStatus *string `json:"ServiceBusinessStatus,omitempty" xml:"ServiceBusinessStatus,omitempty"`
	// The description of the endpoint service.
	//
	// example:
	//
	// This is my EndpointService.
	ServiceDescription *string `json:"ServiceDescription,omitempty" xml:"ServiceDescription,omitempty"`
	// The domain name of the endpoint service.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****.cn-huhehaote.privatelink.aliyuncs.com
	ServiceDomain *string `json:"ServiceDomain,omitempty" xml:"ServiceDomain,omitempty"`
	// The endpoint service ID.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The name of the endpoint service.
	//
	// example:
	//
	// com.aliyuncs.privatelink.cn-huhehaote.epsrv-hp3vpx8yqxblby3i****
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The state of the endpoint service. Valid values:
	//
	// 	- **Creating**: The endpoint service is being created.
	//
	// 	- **Pending**: The endpoint service is being modified.
	//
	// 	- **Active**: The endpoint service is available.
	//
	// 	- **Deleting**: The endpoint service is being deleted.
	//
	// example:
	//
	// Active
	ServiceStatus *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	// Indicates whether IPv6 was enabled for the endpoint service. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	ServiceSupportIPv6 *bool `json:"ServiceSupportIPv6,omitempty" xml:"ServiceSupportIPv6,omitempty"`
	// Indicates whether the domain name of the nearest endpoint that is associated with the endpoint service is resolved first. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	ZoneAffinityEnabled *bool `json:"ZoneAffinityEnabled,omitempty" xml:"ZoneAffinityEnabled,omitempty"`
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

func (s *CreateVpcEndpointServiceResponseBody) SetResourceGroupId(v string) *CreateVpcEndpointServiceResponseBody {
	s.ResourceGroupId = &v
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

func (s *CreateVpcEndpointServiceResponseBody) SetServiceSupportIPv6(v bool) *CreateVpcEndpointServiceResponseBody {
	s.ServiceSupportIPv6 = &v
	return s
}

func (s *CreateVpcEndpointServiceResponseBody) SetZoneAffinityEnabled(v bool) *CreateVpcEndpointServiceResponseBody {
	s.ZoneAffinityEnabled = &v
	return s
}

type CreateVpcEndpointServiceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVpcEndpointServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateVpcEndpointServiceResponse) SetStatusCode(v int32) *CreateVpcEndpointServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVpcEndpointServiceResponse) SetBody(v *CreateVpcEndpointServiceResponseBody) *CreateVpcEndpointServiceResponse {
	s.Body = v
	return s
}

type DeleteVpcEndpointRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the endpoint that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// ep-hp33b2e43fays7s8****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The region ID of the endpoint that you want to delete. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// eu-west-1
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// 8D8992C1-6712-423C-BAC5-E5E817484C6B
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVpcEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteVpcEndpointResponse) SetStatusCode(v int32) *DeleteVpcEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVpcEndpointResponse) SetBody(v *DeleteVpcEndpointResponseBody) *DeleteVpcEndpointResponse {
	s.Body = v
	return s
}

type DeleteVpcEndpointServiceRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The region ID of the endpoint service. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// eu-west-1
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the endpoint service that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVpcEndpointServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteVpcEndpointServiceResponse) SetStatusCode(v int32) *DeleteVpcEndpointServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVpcEndpointServiceResponse) SetBody(v *DeleteVpcEndpointServiceResponseBody) *DeleteVpcEndpointServiceResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
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
	// The available regions.
	Regions *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 611CB80C-B6A9-43DB-9E38-0B0AC3D9B58F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The name of the region.
	//
	// example:
	//
	// China (Hangzhou)
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The endpoint of the region.
	//
	// example:
	//
	// privatelink.cn-hangzhou.aliyuncs.com
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeZonesRequest struct {
	// The region ID of the zone. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
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
	// The request ID.
	//
	// example:
	//
	// 611CB80C-B6A9-43DB-9E38-0B0AC3D9B58F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned zones.
	Zones *DescribeZonesResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
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
	// The name of the zone.
	//
	// example:
	//
	// Hangzhou Zone B
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeZonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeZonesResponse) SetStatusCode(v int32) *DescribeZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeZonesResponse) SetBody(v *DescribeZonesResponseBody) *DescribeZonesResponse {
	s.Body = v
	return s
}

type DetachResourceFromVpcEndpointServiceRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate a value, but you must make sure that the value is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The region ID of the endpoint.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service resource ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-hp32z1wp5peaoox2q****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the service resource. Valid values:
	//
	// 	- **slb**: a Classic Load Balancer (CLB) instance that supports PrivateLink. In addition, the CLB instance is deployed in a virtual private cloud (VPC).
	//
	// 	- **alb**: an Application Load Balancer (ALB) instance that supports PrivateLink. In addition, the ALB instance is deployed in a VPC.
	//
	// example:
	//
	// slb
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The endpoint service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The ID of the zone that you want to remove.
	//
	// example:
	//
	// cn-hangzhou-c
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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

func (s *DetachResourceFromVpcEndpointServiceRequest) SetZoneId(v string) *DetachResourceFromVpcEndpointServiceRequest {
	s.ZoneId = &v
	return s
}

type DetachResourceFromVpcEndpointServiceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
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
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachResourceFromVpcEndpointServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DetachResourceFromVpcEndpointServiceResponse) SetStatusCode(v int32) *DetachResourceFromVpcEndpointServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachResourceFromVpcEndpointServiceResponse) SetBody(v *DetachResourceFromVpcEndpointServiceResponseBody) *DetachResourceFromVpcEndpointServiceResponse {
	s.Body = v
	return s
}

type DetachSecurityGroupFromVpcEndpointRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the endpoint that you want to disassociate from the security group.
	//
	// This parameter is required.
	//
	// example:
	//
	// ep-hp33b2e43fays7s8****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The region ID of the endpoint that you want to disassociate from the security group. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// eu-west-1
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the security group from which you want to disassociate the endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// sg-hp3c8qj1tyct90ej****
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
	// The request ID.
	//
	// example:
	//
	// 8D778FF9-7640-4C13-BCD6-9265CA9A2F81
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
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachSecurityGroupFromVpcEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DetachSecurityGroupFromVpcEndpointResponse) SetStatusCode(v int32) *DetachSecurityGroupFromVpcEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachSecurityGroupFromVpcEndpointResponse) SetBody(v *DetachSecurityGroupFromVpcEndpointResponseBody) *DetachSecurityGroupFromVpcEndpointResponse {
	s.Body = v
	return s
}

type DisableVpcEndpointConnectionRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The endpoint ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ep-hp33b2e43fays7s8****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The ID of the region where the connection request from the endpoint is rejected. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// eu-west-1
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The endpoint service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableVpcEndpointConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DisableVpcEndpointConnectionResponse) SetStatusCode(v int32) *DisableVpcEndpointConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableVpcEndpointConnectionResponse) SetBody(v *DisableVpcEndpointConnectionResponseBody) *DisableVpcEndpointConnectionResponse {
	s.Body = v
	return s
}

type DisableVpcEndpointZoneConnectionRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The endpoint ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ep-hp33b2e43fays7s8****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The ID of the region where the connection request from the endpoint is rejected.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to disconnect the endpoint from the previous connection after the service resource is smoothly migrated. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// > Set the value to true if you want to disconnect the endpoint from the previous connection in the zone after the service resource is smoothly migrated.
	//
	// example:
	//
	// false
	ReplacedResource *bool `json:"ReplacedResource,omitempty" xml:"ReplacedResource,omitempty"`
	// The endpoint service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The ID of the zone that is associated with the endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DisableVpcEndpointZoneConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableVpcEndpointZoneConnectionRequest) GoString() string {
	return s.String()
}

func (s *DisableVpcEndpointZoneConnectionRequest) SetClientToken(v string) *DisableVpcEndpointZoneConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *DisableVpcEndpointZoneConnectionRequest) SetDryRun(v bool) *DisableVpcEndpointZoneConnectionRequest {
	s.DryRun = &v
	return s
}

func (s *DisableVpcEndpointZoneConnectionRequest) SetEndpointId(v string) *DisableVpcEndpointZoneConnectionRequest {
	s.EndpointId = &v
	return s
}

func (s *DisableVpcEndpointZoneConnectionRequest) SetRegionId(v string) *DisableVpcEndpointZoneConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *DisableVpcEndpointZoneConnectionRequest) SetReplacedResource(v bool) *DisableVpcEndpointZoneConnectionRequest {
	s.ReplacedResource = &v
	return s
}

func (s *DisableVpcEndpointZoneConnectionRequest) SetServiceId(v string) *DisableVpcEndpointZoneConnectionRequest {
	s.ServiceId = &v
	return s
}

func (s *DisableVpcEndpointZoneConnectionRequest) SetZoneId(v string) *DisableVpcEndpointZoneConnectionRequest {
	s.ZoneId = &v
	return s
}

type DisableVpcEndpointZoneConnectionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableVpcEndpointZoneConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableVpcEndpointZoneConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *DisableVpcEndpointZoneConnectionResponseBody) SetRequestId(v string) *DisableVpcEndpointZoneConnectionResponseBody {
	s.RequestId = &v
	return s
}

type DisableVpcEndpointZoneConnectionResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableVpcEndpointZoneConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableVpcEndpointZoneConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableVpcEndpointZoneConnectionResponse) GoString() string {
	return s.String()
}

func (s *DisableVpcEndpointZoneConnectionResponse) SetHeaders(v map[string]*string) *DisableVpcEndpointZoneConnectionResponse {
	s.Headers = v
	return s
}

func (s *DisableVpcEndpointZoneConnectionResponse) SetStatusCode(v int32) *DisableVpcEndpointZoneConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableVpcEndpointZoneConnectionResponse) SetBody(v *DisableVpcEndpointZoneConnectionResponseBody) *DisableVpcEndpointZoneConnectionResponse {
	s.Body = v
	return s
}

type EnableVpcEndpointConnectionRequest struct {
	// The bandwidth of the endpoint connection. Unit: Mbit/s. Valid values: **3072 to 10240**.
	//
	// >  The bandwidth of an endpoint connection is in the range of **100 to 10,240*	- Mbit/s. The default bandwidth is **3,072*	- Mbit/s. When the endpoint is connected to the endpoint service, the default bandwidth is the minimum bandwidth. In this case, the connection bandwidth range is **3,072 to 10,240*	- Mbit/s. If Classic Load Balancer (CLB) instances or Application Load Balancer (ALB) instances are specified as service resources, you can modify the endpoint connection bandwidth based on your business requirements. This parameter is invalid if Network Load Balancer (NLB) instances are specified as service resources.
	//
	// example:
	//
	// 1024
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the check, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The endpoint ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ep-hp33b2e43fays7s8****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The ID of the region where the connection request is accepted.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The endpoint service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableVpcEndpointConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *EnableVpcEndpointConnectionResponse) SetStatusCode(v int32) *EnableVpcEndpointConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableVpcEndpointConnectionResponse) SetBody(v *EnableVpcEndpointConnectionResponseBody) *EnableVpcEndpointConnectionResponse {
	s.Body = v
	return s
}

type EnableVpcEndpointZoneConnectionRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The endpoint ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ep-hp33b2e43fays7s8****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The ID of the region where the endpoint connection request is accepted. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The endpoint service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The ID of the zone that is associated with the endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s EnableVpcEndpointZoneConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableVpcEndpointZoneConnectionRequest) GoString() string {
	return s.String()
}

func (s *EnableVpcEndpointZoneConnectionRequest) SetClientToken(v string) *EnableVpcEndpointZoneConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *EnableVpcEndpointZoneConnectionRequest) SetDryRun(v bool) *EnableVpcEndpointZoneConnectionRequest {
	s.DryRun = &v
	return s
}

func (s *EnableVpcEndpointZoneConnectionRequest) SetEndpointId(v string) *EnableVpcEndpointZoneConnectionRequest {
	s.EndpointId = &v
	return s
}

func (s *EnableVpcEndpointZoneConnectionRequest) SetRegionId(v string) *EnableVpcEndpointZoneConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *EnableVpcEndpointZoneConnectionRequest) SetServiceId(v string) *EnableVpcEndpointZoneConnectionRequest {
	s.ServiceId = &v
	return s
}

func (s *EnableVpcEndpointZoneConnectionRequest) SetZoneId(v string) *EnableVpcEndpointZoneConnectionRequest {
	s.ZoneId = &v
	return s
}

type EnableVpcEndpointZoneConnectionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableVpcEndpointZoneConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableVpcEndpointZoneConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *EnableVpcEndpointZoneConnectionResponseBody) SetRequestId(v string) *EnableVpcEndpointZoneConnectionResponseBody {
	s.RequestId = &v
	return s
}

type EnableVpcEndpointZoneConnectionResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableVpcEndpointZoneConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableVpcEndpointZoneConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableVpcEndpointZoneConnectionResponse) GoString() string {
	return s.String()
}

func (s *EnableVpcEndpointZoneConnectionResponse) SetHeaders(v map[string]*string) *EnableVpcEndpointZoneConnectionResponse {
	s.Headers = v
	return s
}

func (s *EnableVpcEndpointZoneConnectionResponse) SetStatusCode(v int32) *EnableVpcEndpointZoneConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableVpcEndpointZoneConnectionResponse) SetBody(v *EnableVpcEndpointZoneConnectionResponseBody) *EnableVpcEndpointZoneConnectionResponse {
	s.Body = v
	return s
}

type GetVpcEndpointAttributeRequest struct {
	// The ID of the endpoint whose attributes you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// ep-hp33b2e43fays7s8****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The region ID of the endpoint whose attributes you want to query.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The bandwidth of the endpoint connection. Unit: Mbit/s.
	//
	// example:
	//
	// 1024
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The state of the endpoint connection. Valid values:
	//
	// 	- **Pending**: The connection is being modified.
	//
	// 	- **Connecting**: The connection is being established.
	//
	// 	- **Connected**: The connection is established.
	//
	// 	- **Disconnecting**: The endpoint is being disconnected from the endpoint service.
	//
	// 	- **Disconnected**: The endpoint is disconnected from the endpoint service.
	//
	// 	- **Deleting**: The connection is being deleted.
	//
	// example:
	//
	// Connected
	ConnectionStatus *string `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	// The time when the endpoint was created.
	//
	// example:
	//
	// 2021-09-24T18:00:07Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The service state of the endpoint. Valid values:
	//
	// 	- **Normal**: The endpoint runs as expected.
	//
	// 	- **FinancialLocked**: The endpoint is locked due to overdue payments.
	//
	// example:
	//
	// Normal
	EndpointBusinessStatus *string `json:"EndpointBusinessStatus,omitempty" xml:"EndpointBusinessStatus,omitempty"`
	// The description of the endpoint.
	//
	// example:
	//
	// This is my Endpoint.
	EndpointDescription *string `json:"EndpointDescription,omitempty" xml:"EndpointDescription,omitempty"`
	// The domain name of the endpoint.
	//
	// example:
	//
	// ep-hp33b2e43fays7s8****.epsrv-hp3xdsq46ael67lo****.cn-huhehaote.privatelink.aliyuncs.com
	EndpointDomain *string `json:"EndpointDomain,omitempty" xml:"EndpointDomain,omitempty"`
	// The endpoint ID.
	//
	// example:
	//
	// ep-hp33b2e43fays7s8****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The name of the endpoint.
	//
	// example:
	//
	// test
	EndpointName *string `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
	// The state of the endpoint. Valid values:
	//
	// 	- **Creating**: The endpoint is being created.
	//
	// 	- **Active**: The endpoint is available.
	//
	// 	- **Pending**: The endpoint is being modified.
	//
	// 	- **Deleting**: The endpoint is being deleted.
	//
	// example:
	//
	// Active
	EndpointStatus *string `json:"EndpointStatus,omitempty" xml:"EndpointStatus,omitempty"`
	// The type of the endpoint.
	//
	// **Interface*	- is returned. The value indicates the interface endpoint with which the Classic Load Balancer (CLB) instances are associated.
	//
	// example:
	//
	// Interface
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The payer. Valid values:
	//
	// 	- **Endpoint**: the service consumer.
	//
	// 	- **EndpointService**: the service provider.
	//
	// example:
	//
	// Endpoint
	Payer          *string `json:"Payer,omitempty" xml:"Payer,omitempty"`
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	// The region ID of the endpoint.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8D8992C1-6712-423C-BAC5-E5E817484C6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmz7nocpei***
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Indicates whether the endpoint and the endpoint service belong to the same Alibaba Cloud account. Valid values:
	//
	// 	- **true**: The endpoint and the endpoint service belong to the same Alibaba Cloud account.
	//
	// 	- **false**: The endpoint and the endpoint service do not belong to the same Alibaba Cloud account.
	//
	// example:
	//
	// true
	ResourceOwner *bool `json:"ResourceOwner,omitempty" xml:"ResourceOwner,omitempty"`
	// The ID of the endpoint service with which the endpoint is associated.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The name of the endpoint service with which the endpoint is associated.
	//
	// example:
	//
	// com.aliyuncs.privatelink.cn-huhehaote.epsrv-hp3xdsq46ael67lo****
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The ID of the virtual private cloud (VPC) to which the endpoint belongs.
	//
	// example:
	//
	// vpc-fdfhkjafhjvcvdjf****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// Indicates whether the domain name of the nearest endpoint that is associated with the endpoint service is resolved first. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	ZoneAffinityEnabled *bool `json:"ZoneAffinityEnabled,omitempty" xml:"ZoneAffinityEnabled,omitempty"`
	// The number of private IP addresses that are assigned to an elastic network interface (ENI) in each zone. Only **1*	- is returned.
	//
	// example:
	//
	// 1
	ZonePrivateIpAddressCount *int64 `json:"ZonePrivateIpAddressCount,omitempty" xml:"ZonePrivateIpAddressCount,omitempty"`
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

func (s *GetVpcEndpointAttributeResponseBody) SetPolicyDocument(v string) *GetVpcEndpointAttributeResponseBody {
	s.PolicyDocument = &v
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

func (s *GetVpcEndpointAttributeResponseBody) SetResourceGroupId(v string) *GetVpcEndpointAttributeResponseBody {
	s.ResourceGroupId = &v
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVpcEndpointAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *GetVpcEndpointAttributeResponse) SetStatusCode(v int32) *GetVpcEndpointAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVpcEndpointAttributeResponse) SetBody(v *GetVpcEndpointAttributeResponseBody) *GetVpcEndpointAttributeResponse {
	s.Body = v
	return s
}

type GetVpcEndpointServiceAttributeRequest struct {
	// The region ID of the endpoint service.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The endpoint service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
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
	// Indicates whether endpoint connection requests are automatically accepted. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	AutoAcceptEnabled *bool `json:"AutoAcceptEnabled,omitempty" xml:"AutoAcceptEnabled,omitempty"`
	// The default maximum bandwidth of the endpoint connection. Unit: Mbit/s. Valid values: **100*	- to 10240.
	//
	// example:
	//
	// 1024
	ConnectBandwidth *int32 `json:"ConnectBandwidth,omitempty" xml:"ConnectBandwidth,omitempty"`
	// The time when the endpoint service was created.
	//
	// example:
	//
	// 2020-01-02T19:11:12Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The maximum bandwidth of the endpoint connection. Unit: Mbit/s.
	//
	// example:
	//
	// 1024
	MaxBandwidth *int32 `json:"MaxBandwidth,omitempty" xml:"MaxBandwidth,omitempty"`
	// The minimum bandwidth of the endpoint connection. Unit: Mbit/s.
	//
	// example:
	//
	// 100
	MinBandwidth *int32 `json:"MinBandwidth,omitempty" xml:"MinBandwidth,omitempty"`
	// The payer of the endpoint service. Valid values:
	//
	// 	- **Endpoint**: the service consumer.
	//
	// 	- **EndpointService**: the service provider.
	//
	// example:
	//
	// Endpoint
	Payer *string `json:"Payer,omitempty" xml:"Payer,omitempty"`
	// The region ID of the endpoint service.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8D8992C1-6712-423C-BAC5-E5E817484C6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The service state of the endpoint service. Valid values:
	//
	// 	- **Normal**: The endpoint service runs as expected.
	//
	// 	- **FinacialLocked**: The endpoint service is locked due to overdue payments.
	//
	// example:
	//
	// Normal
	ServiceBusinessStatus *string `json:"ServiceBusinessStatus,omitempty" xml:"ServiceBusinessStatus,omitempty"`
	// The description of the endpoint service.
	//
	// example:
	//
	// This is my EndpointService.
	ServiceDescription *string `json:"ServiceDescription,omitempty" xml:"ServiceDescription,omitempty"`
	// The domain name of the endpoint service.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****.cn-huhehaote.privatelink.aliyuncs.com
	ServiceDomain *string `json:"ServiceDomain,omitempty" xml:"ServiceDomain,omitempty"`
	// The endpoint service ID.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The name of the endpoint service.
	//
	// example:
	//
	// com.aliyuncs.privatelink.cn-huhehaote.epsrv-hp3vpx8yqxblby3i****
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The type of the service resource. Valid values:
	//
	// 	- **slb**: a CLB instance.
	//
	// 	- **alb**: an ALB instance.
	//
	// example:
	//
	// slb
	ServiceResourceType *string `json:"ServiceResourceType,omitempty" xml:"ServiceResourceType,omitempty"`
	// The state of the endpoint service. Valid values:
	//
	// 	- **Creating**: The endpoint service is being created.
	//
	// 	- **Pending**: The endpoint service is being modified.
	//
	// 	- **Active**: The endpoint service is available.
	//
	// 	- **Deleting**: The endpoint service is being deleted.
	//
	// 	- **Inactive**: The endpoint service is unavailable.
	//
	// example:
	//
	// Active
	ServiceStatus *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	// Indicates whether IPv6 is enabled for the endpoint service. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	ServiceSupportIPv6 *bool `json:"ServiceSupportIPv6,omitempty" xml:"ServiceSupportIPv6,omitempty"`
	// The type of the endpoint.
	//
	// Only **Interface*	- is returned. The value indicates the interface endpoint. Then, you can specify ALB and CLB instances as service resources for the endpoint service.
	//
	// example:
	//
	// Interface
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// Indicates whether the domain name of the nearest endpoint that is associated with the endpoint service is resolved first. Valid values:
	//
	// 	- **true*	- (default)
	//
	// 	- **false**
	//
	// example:
	//
	// true
	ZoneAffinityEnabled *bool `json:"ZoneAffinityEnabled,omitempty" xml:"ZoneAffinityEnabled,omitempty"`
	// The zones to which the service resources belong.
	Zones []*string `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
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

func (s *GetVpcEndpointServiceAttributeResponseBody) SetResourceGroupId(v string) *GetVpcEndpointServiceAttributeResponseBody {
	s.ResourceGroupId = &v
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

func (s *GetVpcEndpointServiceAttributeResponseBody) SetServiceSupportIPv6(v bool) *GetVpcEndpointServiceAttributeResponseBody {
	s.ServiceSupportIPv6 = &v
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
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVpcEndpointServiceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *GetVpcEndpointServiceAttributeResponse) SetStatusCode(v int32) *GetVpcEndpointServiceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponse) SetBody(v *GetVpcEndpointServiceAttributeResponseBody) *GetVpcEndpointServiceAttributeResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// 	- If this is your first request or no next requests are to be sent, you do not need to specify this parameter.
	//
	// 	- If a next request is to be sent, you must specify the token that is obtained from the previous request as the value of **NextToken**.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region where the resource resides. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource IDs. You can specify up to 50 resource IDs.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the resource. Valid values:
	//
	// 	- **vpcendpoint**: endpoint
	//
	// 	- **vpcendpointservice**: endpoint service
	//
	// This parameter is required.
	//
	// example:
	//
	// vpcendpoint
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags to add to the resource.
	Tag []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetClientToken(v string) *ListTagResourcesRequest {
	s.ClientToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
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

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	// The key of tag N to add to the resource. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The tag key must start with a letter but cannot start with `aliyun` or `acs:`. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the resource. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The tag value must start with a letter but cannot start with `aliyun` or `acs:`. The tag value cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
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
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If this is your first request and no next requests are to be performed, you do not need to specify this parameter.
	//
	// 	- If a next request is to be performed, set the parameter to the value of NextToken that is returned from the last call.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8D8992C1-6712-423C-BAC5-E5E817484C6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resources to which the tags are added.
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
	// The resource ID.
	//
	// example:
	//
	// ep-hp3i05294c2d2d******
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource. Valid values:
	//
	// 	- **vpcendpoint**: endpoint
	//
	// 	- **vpcendpointservice**: endpoint service
	//
	// example:
	//
	// vpcendpoint
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The key of tag N added to the resource.
	//
	// example:
	//
	// endpoint
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of tag N added to the resource.
	//
	// example:
	//
	// FinanceJoshua
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ListVpcEndpointConnectionsRequest struct {
	// The state of the endpoint connection. Valid values:
	//
	// 	- **Pending**: The endpoint connection is being modified.
	//
	// 	- **Connecting**: The endpoint connection is being established.
	//
	// 	- **Connected**: The endpoint connection is established.
	//
	// 	- **Disconnecting**: The endpoint is being disconnected from the endpoint service.
	//
	// 	- **Disconnected**: The endpoint is disconnected from the endpoint service.
	//
	// 	- **Deleting**: The connection is being deleted.
	//
	// 	- **ServiceDeleted**: The corresponding endpoint service has been deleted.
	//
	// example:
	//
	// Disconnected
	ConnectionStatus *string `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	// The endpoint ID.
	//
	// example:
	//
	// ep-hp33b2e43fays7s8****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The ID of the Alibaba Cloud account to which the endpoint belongs.
	//
	// example:
	//
	// 25346073170691****
	EndpointOwnerId *int64 `json:"EndpointOwnerId,omitempty" xml:"EndpointOwnerId,omitempty"`
	// The ID of the endpoint elastic network interface (ENI).
	//
	// example:
	//
	// eni-hp32lk0pyv6o94hs****
	EniId *string `json:"EniId,omitempty" xml:"EniId,omitempty"`
	// The number of entries to return on each page. Valid values: **1*	- to **50**. Default value: **50**.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If this is your first request and no next requests are to be performed, you do not need to specify this parameter.
	//
	// 	- If a next request is to be performed, set the value to the value of **NextToken*	- that is returned from the last call.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the endpoint connection.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the replaced service resource in smooth migration scenarios.
	//
	// example:
	//
	// lb-hp32z1wp5peaoox2q****
	ReplacedResourceId *string `json:"ReplacedResourceId,omitempty" xml:"ReplacedResourceId,omitempty"`
	// The ID of the resource group to which the endpoint belongs.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The service resource ID.
	//
	// example:
	//
	// lb-hp32z1wp5peaoox2q****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The endpoint service ID.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
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

func (s *ListVpcEndpointConnectionsRequest) SetReplacedResourceId(v string) *ListVpcEndpointConnectionsRequest {
	s.ReplacedResourceId = &v
	return s
}

func (s *ListVpcEndpointConnectionsRequest) SetResourceGroupId(v string) *ListVpcEndpointConnectionsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVpcEndpointConnectionsRequest) SetResourceId(v string) *ListVpcEndpointConnectionsRequest {
	s.ResourceId = &v
	return s
}

func (s *ListVpcEndpointConnectionsRequest) SetServiceId(v string) *ListVpcEndpointConnectionsRequest {
	s.ServiceId = &v
	return s
}

type ListVpcEndpointConnectionsResponseBody struct {
	// The endpoint connections.
	Connections []*ListVpcEndpointConnectionsResponseBodyConnections `json:"Connections,omitempty" xml:"Connections,omitempty" type:"Repeated"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If no value is returned for **NextToken**, no next requests are performed.
	//
	// 	- If a value is returned for **NextToken**, the value can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *ListVpcEndpointConnectionsResponseBody) SetMaxResults(v int32) *ListVpcEndpointConnectionsResponseBody {
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

func (s *ListVpcEndpointConnectionsResponseBody) SetTotalCount(v string) *ListVpcEndpointConnectionsResponseBody {
	s.TotalCount = &v
	return s
}

type ListVpcEndpointConnectionsResponseBodyConnections struct {
	// The bandwidth of the endpoint connection. Valid values: **1024 to 10240**. Unit: Mbit/s.
	//
	// example:
	//
	// 1024
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The state of the endpoint connection. Valid values:
	//
	// 	- **Pending**: The connection is being modified.
	//
	// 	- **Connecting**: The connection is being established.
	//
	// 	- **Connected**: The connection is established.
	//
	// 	- **Disconnecting**: The endpoint is being disconnected from the endpoint service.
	//
	// 	- **Disconnected**: The endpoint is disconnected from the endpoint service.
	//
	// 	- **Deleting**: The connection is being deleted.
	//
	// 	- **ServiceDeleted**: The corresponding endpoint service has been deleted.
	//
	// example:
	//
	// Disconnected
	ConnectionStatus *string `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	// The endpoint ID.
	//
	// example:
	//
	// ep-hp33b2e43fays7s8****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The ID of the Alibaba Cloud account to which the endpoint belongs.
	//
	// example:
	//
	// 25346073170691****
	EndpointOwnerId *int64 `json:"EndpointOwnerId,omitempty" xml:"EndpointOwnerId,omitempty"`
	// The ID of the virtual private cloud (VPC) to which the endpoint belongs.
	//
	// example:
	//
	// vpc-hp356stwkxg3fn2xe****
	EndpointVpcId *string `json:"EndpointVpcId,omitempty" xml:"EndpointVpcId,omitempty"`
	// The time when the endpoint connection was modified.
	//
	// example:
	//
	// 2021-09-28T10:34:46Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The ID of the resource group to which the endpoint belongs.
	//
	// example:
	//
	// rg-acfmw353z35v***
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Indicates whether the endpoint and the endpoint service belong to the same Alibaba Cloud account. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	ResourceOwner *bool `json:"ResourceOwner,omitempty" xml:"ResourceOwner,omitempty"`
	// The endpoint service ID.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The zones.
	Zones []*ListVpcEndpointConnectionsResponseBodyConnectionsZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
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

func (s *ListVpcEndpointConnectionsResponseBodyConnections) SetResourceGroupId(v string) *ListVpcEndpointConnectionsResponseBodyConnections {
	s.ResourceGroupId = &v
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
	// The endpoint ENI ID.
	//
	// example:
	//
	// eni-hp32lk0pyv6o94hs****
	EniId *string `json:"EniId,omitempty" xml:"EniId,omitempty"`
	// The ID of the replaced endpoint ENI in smooth migration scenarios.
	//
	// example:
	//
	// eni-hp32lk0pyv6o94hs****
	ReplacedEniId *string `json:"ReplacedEniId,omitempty" xml:"ReplacedEniId,omitempty"`
	// The ID of the replaced service resource in smooth migration scenarios.
	//
	// example:
	//
	// lb-hp32z1wp5peaoox2q****
	ReplacedResourceId *string `json:"ReplacedResourceId,omitempty" xml:"ReplacedResourceId,omitempty"`
	// The service resource ID.
	//
	// example:
	//
	// lb-hp32z1wp5peaoox2q****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The ID of the vSwitch to which the endpoint belongs.
	//
	// example:
	//
	// vsw-hp3uf6045ljdhd5zr****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The domain name of the zone.
	//
	// example:
	//
	// ep-hp34jaz8ykl0exwt****-cn-huhehaote.epsrv-hp3vpx8yqxblby3i****.cn-huhehaote-b.privatelink.aliyuncs.com
	ZoneDomain *string `json:"ZoneDomain,omitempty" xml:"ZoneDomain,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-huhehaote-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The state of the zone. Valid values:
	//
	// 	- **Creating**: The zone is being created.
	//
	// 	- **Wait**: The zone is to be connected.
	//
	// 	- **Connected**: The zone is connected.
	//
	// 	- **Deleting**: The zone is being deleted.
	//
	// 	- **Disconnecting**: The zone is being disconnected.
	//
	// 	- **Disconnected**: The zone is disconnected.
	//
	// 	- **Connecting**: The zone is being connected.
	//
	// 	- **Migrating**: The zone is being migrated.
	//
	// 	- **Migrated**: The zone is migrated.
	//
	// example:
	//
	// Connected
	ZoneStatus *string `json:"ZoneStatus,omitempty" xml:"ZoneStatus,omitempty"`
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

func (s *ListVpcEndpointConnectionsResponseBodyConnectionsZones) SetReplacedEniId(v string) *ListVpcEndpointConnectionsResponseBodyConnectionsZones {
	s.ReplacedEniId = &v
	return s
}

func (s *ListVpcEndpointConnectionsResponseBodyConnectionsZones) SetReplacedResourceId(v string) *ListVpcEndpointConnectionsResponseBodyConnectionsZones {
	s.ReplacedResourceId = &v
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

func (s *ListVpcEndpointConnectionsResponseBodyConnectionsZones) SetZoneStatus(v string) *ListVpcEndpointConnectionsResponseBodyConnectionsZones {
	s.ZoneStatus = &v
	return s
}

type ListVpcEndpointConnectionsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVpcEndpointConnectionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ListVpcEndpointConnectionsResponse) SetStatusCode(v int32) *ListVpcEndpointConnectionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVpcEndpointConnectionsResponse) SetBody(v *ListVpcEndpointConnectionsResponseBody) *ListVpcEndpointConnectionsResponse {
	s.Body = v
	return s
}

type ListVpcEndpointSecurityGroupsRequest struct {
	// The ID of the endpoint that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// ep-hp33b2e43fays7s8****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The number of entries to return on each page. Valid values:**1*	- to **50**. Default value: **50**.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If this is your first request and no next requests are to be performed, you do not need to specify this parameter.
	//
	// 	- If a next request is to be performed, set the parameter to the value of **NextToken*	- that is returned from the last call.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the endpoint that you want to query.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The number of entries returned per page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If no value is returned for **NextToken**, no next requests are performed.
	//
	// 	- If a value is returned for **NextToken**, the value can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the security groups.
	SecurityGroups []*ListVpcEndpointSecurityGroupsResponseBodySecurityGroups `json:"SecurityGroups,omitempty" xml:"SecurityGroups,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 12
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListVpcEndpointSecurityGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointSecurityGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointSecurityGroupsResponseBody) SetMaxResults(v int32) *ListVpcEndpointSecurityGroupsResponseBody {
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

func (s *ListVpcEndpointSecurityGroupsResponseBody) SetTotalCount(v int32) *ListVpcEndpointSecurityGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type ListVpcEndpointSecurityGroupsResponseBodySecurityGroups struct {
	// The ID of the security group that is associated with the endpoint.
	//
	// example:
	//
	// sg-hp33bw6ynvm2yb0e****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The associate status of the security group, valid values:
	//
	// - Attaching: The security group is being attached.
	//
	// - Attached: The security group is attached.
	//
	// - Detaching: The security group is being detached.
	//
	// example:
	//
	// Attached
	SecurityGroupStatus *string `json:"SecurityGroupStatus,omitempty" xml:"SecurityGroupStatus,omitempty"`
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

func (s *ListVpcEndpointSecurityGroupsResponseBodySecurityGroups) SetSecurityGroupStatus(v string) *ListVpcEndpointSecurityGroupsResponseBodySecurityGroups {
	s.SecurityGroupStatus = &v
	return s
}

type ListVpcEndpointSecurityGroupsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVpcEndpointSecurityGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ListVpcEndpointSecurityGroupsResponse) SetStatusCode(v int32) *ListVpcEndpointSecurityGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVpcEndpointSecurityGroupsResponse) SetBody(v *ListVpcEndpointSecurityGroupsResponseBody) *ListVpcEndpointSecurityGroupsResponse {
	s.Body = v
	return s
}

type ListVpcEndpointServiceResourcesRequest struct {
	// The number of entries to return on each page. Valid values: **1*	- to **50**. Default value: **50**.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If this is your first request and no next requests are to be performed, you do not need to specify this parameter.
	//
	// 	- If a next request is to be performed, set the parameter to the value of NextToken that is returned from the last call.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the service resource.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The endpoint service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
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
	// The number of entries returned on each page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If no value is returned for **NextToken**, no next requests are performed.
	//
	// 	- If a value is returned for **NextToken**, the value can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The service resources.
	Resources []*ListVpcEndpointServiceResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s ListVpcEndpointServiceResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointServiceResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServiceResourcesResponseBody) SetMaxResults(v int32) *ListVpcEndpointServiceResourcesResponseBody {
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
	// Indicates whether automatic resource allocation is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	AutoAllocatedEnabled *bool `json:"AutoAllocatedEnabled,omitempty" xml:"AutoAllocatedEnabled,omitempty"`
	// The IP address of the service resource.
	//
	// example:
	//
	// 192.168.10.23
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The ID of the region where the service resource is deployed.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of endpoints that are associated with the service resource that is smoothly migrated.
	//
	// example:
	//
	// 10
	RelatedDeprecatedEndpointCount *int64 `json:"RelatedDeprecatedEndpointCount,omitempty" xml:"RelatedDeprecatedEndpointCount,omitempty"`
	// The number of endpoints that are associated with the service resource.
	//
	// example:
	//
	// 10
	RelatedEndpointCount *int64 `json:"RelatedEndpointCount,omitempty" xml:"RelatedEndpointCount,omitempty"`
	// The service resource ID.
	//
	// example:
	//
	// lb-hp32z1wp5peaoox2q****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// Indicates whether IPv6 is enabled for the endpoint service. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	ResourceSupportIPv6 *bool `json:"ResourceSupportIPv6,omitempty" xml:"ResourceSupportIPv6,omitempty"`
	// The type of the service resource.
	//
	// Only **slb*	- is returned. This value indicates a Classic Load Balancer (CLB) instance.
	//
	// example:
	//
	// slb
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The ID of the vSwitch to which the service resource belongs.
	//
	// example:
	//
	// vsw-hp3uf6045ljdhd5zr****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC) to which the service resource belongs.
	//
	// example:
	//
	// vpc-hp356stwkxg3fn2xe****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the zone to which the service resource belongs.
	//
	// example:
	//
	// cn-huhehaote-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListVpcEndpointServiceResourcesResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointServiceResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServiceResourcesResponseBodyResources) SetAutoAllocatedEnabled(v bool) *ListVpcEndpointServiceResourcesResponseBodyResources {
	s.AutoAllocatedEnabled = &v
	return s
}

func (s *ListVpcEndpointServiceResourcesResponseBodyResources) SetIp(v string) *ListVpcEndpointServiceResourcesResponseBodyResources {
	s.Ip = &v
	return s
}

func (s *ListVpcEndpointServiceResourcesResponseBodyResources) SetRegionId(v string) *ListVpcEndpointServiceResourcesResponseBodyResources {
	s.RegionId = &v
	return s
}

func (s *ListVpcEndpointServiceResourcesResponseBodyResources) SetRelatedDeprecatedEndpointCount(v int64) *ListVpcEndpointServiceResourcesResponseBodyResources {
	s.RelatedDeprecatedEndpointCount = &v
	return s
}

func (s *ListVpcEndpointServiceResourcesResponseBodyResources) SetRelatedEndpointCount(v int64) *ListVpcEndpointServiceResourcesResponseBodyResources {
	s.RelatedEndpointCount = &v
	return s
}

func (s *ListVpcEndpointServiceResourcesResponseBodyResources) SetResourceId(v string) *ListVpcEndpointServiceResourcesResponseBodyResources {
	s.ResourceId = &v
	return s
}

func (s *ListVpcEndpointServiceResourcesResponseBodyResources) SetResourceSupportIPv6(v bool) *ListVpcEndpointServiceResourcesResponseBodyResources {
	s.ResourceSupportIPv6 = &v
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
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVpcEndpointServiceResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ListVpcEndpointServiceResourcesResponse) SetStatusCode(v int32) *ListVpcEndpointServiceResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVpcEndpointServiceResourcesResponse) SetBody(v *ListVpcEndpointServiceResourcesResponseBody) *ListVpcEndpointServiceResourcesResponse {
	s.Body = v
	return s
}

type ListVpcEndpointServiceUsersRequest struct {
	// The number of entries to return on each page. Valid values: **1 to 50**. Default value: **50**.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If this is your first request and no next requests are to be performed, you do not need to specify this parameter.
	//
	// 	- If a next request is to be performed, set the value to the value of **NextToken*	- that is returned from the last call.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the endpoint service that you want to query.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The endpoint service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The ID of the Alibaba Cloud account in the whitelist of the endpoint service.
	//
	// example:
	//
	// 12345678
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The type of the user list in the whitelist of the endpoint service.
	//
	// example:
	//
	// Users
	UserListType *string `json:"UserListType,omitempty" xml:"UserListType,omitempty"`
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

func (s *ListVpcEndpointServiceUsersRequest) SetUserListType(v string) *ListVpcEndpointServiceUsersRequest {
	s.UserListType = &v
	return s
}

type ListVpcEndpointServiceUsersResponseBody struct {
	// The number of entries returned on each page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If no value is returned for **NextToken**, no next requests are performed.
	//
	// 	- If a value is returned for **NextToken**, the value can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 3
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The whitelists in the format of Aliyun Resource Name (ARN).
	UserARNs []*ListVpcEndpointServiceUsersResponseBodyUserARNs `json:"UserARNs,omitempty" xml:"UserARNs,omitempty" type:"Repeated"`
	// The Alibaba Cloud accounts in the whitelist of the endpoint service.
	Users []*ListVpcEndpointServiceUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ListVpcEndpointServiceUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointServiceUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServiceUsersResponseBody) SetMaxResults(v int32) *ListVpcEndpointServiceUsersResponseBody {
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

func (s *ListVpcEndpointServiceUsersResponseBody) SetTotalCount(v string) *ListVpcEndpointServiceUsersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVpcEndpointServiceUsersResponseBody) SetUserARNs(v []*ListVpcEndpointServiceUsersResponseBodyUserARNs) *ListVpcEndpointServiceUsersResponseBody {
	s.UserARNs = v
	return s
}

func (s *ListVpcEndpointServiceUsersResponseBody) SetUsers(v []*ListVpcEndpointServiceUsersResponseBodyUsers) *ListVpcEndpointServiceUsersResponseBody {
	s.Users = v
	return s
}

type ListVpcEndpointServiceUsersResponseBodyUserARNs struct {
	// The whitelist in the format of ARN.
	//
	// example:
	//
	// acs:ram:*::*
	UserARN *string `json:"UserARN,omitempty" xml:"UserARN,omitempty"`
}

func (s ListVpcEndpointServiceUsersResponseBodyUserARNs) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointServiceUsersResponseBodyUserARNs) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServiceUsersResponseBodyUserARNs) SetUserARN(v string) *ListVpcEndpointServiceUsersResponseBodyUserARNs {
	s.UserARN = &v
	return s
}

type ListVpcEndpointServiceUsersResponseBodyUsers struct {
	// The ID of the Alibaba Cloud account in the whitelist of the endpoint service.
	//
	// example:
	//
	// 12345678
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVpcEndpointServiceUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ListVpcEndpointServiceUsersResponse) SetStatusCode(v int32) *ListVpcEndpointServiceUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVpcEndpointServiceUsersResponse) SetBody(v *ListVpcEndpointServiceUsersResponseBody) *ListVpcEndpointServiceUsersResponse {
	s.Body = v
	return s
}

type ListVpcEndpointServicesRequest struct {
	// Specifies whether to automatically accept endpoint connection requests. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	AutoAcceptEnabled *bool `json:"AutoAcceptEnabled,omitempty" xml:"AutoAcceptEnabled,omitempty"`
	// The number of entries to return on each page. Valid values: **1*	- to **50**. Default value: **50**.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If this is your first request and no next requests are to be performed, you do not need to specify this parameter.
	//
	// 	- If a next request is to be performed, set the parameter to the value of NextToken that is returned from the last call.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the endpoint service.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The service resource ID.
	//
	// example:
	//
	// lb-hp32z1wp5peaoox2q****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The service state of the endpoint service. Valid values:
	//
	// 	- **Normal**: The endpoint service runs as expected.
	//
	// 	- **FinancialLocked**: The endpoint service is locked due to overdue payments.
	//
	// example:
	//
	// Normal
	ServiceBusinessStatus *string `json:"ServiceBusinessStatus,omitempty" xml:"ServiceBusinessStatus,omitempty"`
	// The endpoint service ID.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The name of the endpoint service.
	//
	// example:
	//
	// com.aliyuncs.privatelink.cn-huhehaote.epsrv-hp3vpx8yqxblby3i****
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The type of the service resource. Valid values:
	//
	// 	- **slb**: a Classic Load Balancer (CLB) instance
	//
	// 	- **alb**: an Application Load Balancer (ALB) instance
	//
	// example:
	//
	// slb
	ServiceResourceType *string `json:"ServiceResourceType,omitempty" xml:"ServiceResourceType,omitempty"`
	// The state of the endpoint service. Valid values:
	//
	// 	- **Creating**: The endpoint service is being created.
	//
	// 	- **Pending**: The endpoint service is being modified.
	//
	// 	- **Active**: The endpoint service is available.
	//
	// 	- **Deleting**: The endpoint service is being deleted
	//
	// example:
	//
	// Active
	ServiceStatus *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	// The list of tags.
	Tag []*ListVpcEndpointServicesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Specifies whether to first resolve the domain name of the nearest endpoint that is associated with the endpoint service. Valid values:
	//
	// 	- **true*	- (default)
	//
	// 	- **false**
	//
	// example:
	//
	// true
	ZoneAffinityEnabled *bool `json:"ZoneAffinityEnabled,omitempty" xml:"ZoneAffinityEnabled,omitempty"`
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

func (s *ListVpcEndpointServicesRequest) SetResourceGroupId(v string) *ListVpcEndpointServicesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVpcEndpointServicesRequest) SetResourceId(v string) *ListVpcEndpointServicesRequest {
	s.ResourceId = &v
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

func (s *ListVpcEndpointServicesRequest) SetTag(v []*ListVpcEndpointServicesRequestTag) *ListVpcEndpointServicesRequest {
	s.Tag = v
	return s
}

func (s *ListVpcEndpointServicesRequest) SetZoneAffinityEnabled(v bool) *ListVpcEndpointServicesRequest {
	s.ZoneAffinityEnabled = &v
	return s
}

type ListVpcEndpointServicesRequestTag struct {
	// The key of the tag. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key must be 1 to 64 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListVpcEndpointServicesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointServicesRequestTag) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServicesRequestTag) SetKey(v string) *ListVpcEndpointServicesRequestTag {
	s.Key = &v
	return s
}

func (s *ListVpcEndpointServicesRequestTag) SetValue(v string) *ListVpcEndpointServicesRequestTag {
	s.Value = &v
	return s
}

type ListVpcEndpointServicesResponseBody struct {
	// The number of entries returned per page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If no value is returned for **NextToken**, no next requests are performed.
	//
	// 	- If a value is returned for **NextToken**, the value can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The endpoint services.
	Services []*ListVpcEndpointServicesResponseBodyServices `json:"Services,omitempty" xml:"Services,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 12
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListVpcEndpointServicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServicesResponseBody) SetMaxResults(v int32) *ListVpcEndpointServicesResponseBody {
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

func (s *ListVpcEndpointServicesResponseBody) SetTotalCount(v int32) *ListVpcEndpointServicesResponseBody {
	s.TotalCount = &v
	return s
}

type ListVpcEndpointServicesResponseBodyServices struct {
	// Indicates whether endpoint connection requests are automatically accepted. Valid values:
	//
	// 	- **true**: Endpoint connection requests are automatically accepted.
	//
	// 	- **false**: Endpoint connection requests are not automatically accepted.
	//
	// example:
	//
	// true
	AutoAcceptEnabled *bool `json:"AutoAcceptEnabled,omitempty" xml:"AutoAcceptEnabled,omitempty"`
	// The default maximum bandwidth of the endpoint connection. Unit: Mbit/s.
	//
	// example:
	//
	// 1024
	ConnectBandwidth *int32 `json:"ConnectBandwidth,omitempty" xml:"ConnectBandwidth,omitempty"`
	// The time when the endpoint service was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-09-24T17:15:10Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The maximum bandwidth of the endpoint connection. Unit: Mbit/s.
	//
	// example:
	//
	// 1024
	MaxBandwidth *int32 `json:"MaxBandwidth,omitempty" xml:"MaxBandwidth,omitempty"`
	// The minimum bandwidth of the endpoint connection. Unit: Mbit/s.
	//
	// example:
	//
	// 100
	MinBandwidth *int32 `json:"MinBandwidth,omitempty" xml:"MinBandwidth,omitempty"`
	// The payer. Valid values:
	//
	// 	- **Endpoint**: service consumer
	//
	// 	- **EndpointService**: service provider
	//
	// example:
	//
	// Endpoint
	Payer *string `json:"Payer,omitempty" xml:"Payer,omitempty"`
	// The region ID of the endpoint service.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The service state of the endpoint service. Valid values:
	//
	// 	- **Normal**: The endpoint service runs as expected.
	//
	// 	- **FinancialLocked**: The endpoint service is locked due to overdue payments.
	//
	// example:
	//
	// Normal
	ServiceBusinessStatus *string `json:"ServiceBusinessStatus,omitempty" xml:"ServiceBusinessStatus,omitempty"`
	// The description of the endpoint service.
	//
	// example:
	//
	// This is my EndpointService.
	ServiceDescription *string `json:"ServiceDescription,omitempty" xml:"ServiceDescription,omitempty"`
	// The domain name of the endpoint service.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****.cn-huhehaote.privatelink.aliyuncs.com
	ServiceDomain *string `json:"ServiceDomain,omitempty" xml:"ServiceDomain,omitempty"`
	// The ID of the endpoint service.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The name of the endpoint service.
	//
	// example:
	//
	// com.aliyuncs.privatelink.cn-huhehaote.epsrv-hp3vpx8yqxblby3i****
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The type of the service resource. Valid values:
	//
	// 	- **slb**: Classic Load Balancer (CLB) instance
	//
	// 	- **alb**: Application Load Balancer (ALB) instance
	//
	// 	- **nlb**: Network Load Balancer (NLB) instance
	//
	// example:
	//
	// slb
	ServiceResourceType *string `json:"ServiceResourceType,omitempty" xml:"ServiceResourceType,omitempty"`
	// The state of the endpoint service. Valid values:
	//
	// 	- **Creating**: The endpoint service is being created.
	//
	// 	- **Pending**: The endpoint service is being modified.
	//
	// 	- **Active**: The endpoint service is available.
	//
	// 	- **Deleting**: The endpoint service is being deleted.
	//
	// example:
	//
	// Active
	ServiceStatus *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	// Indicates whether the endpoint service supports IPv6. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	ServiceSupportIPv6 *bool `json:"ServiceSupportIPv6,omitempty" xml:"ServiceSupportIPv6,omitempty"`
	// The type of the endpoint service.
	//
	// 	- Only **Interface*	- may be returned. You can specify CLB, ALB, and NLB instances as the service resources of the endpoint service.
	//
	// example:
	//
	// Interface
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The tags added to the resource.
	Tags []*ListVpcEndpointServicesResponseBodyServicesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// Indicates whether zone affinity is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	ZoneAffinityEnabled *bool `json:"ZoneAffinityEnabled,omitempty" xml:"ZoneAffinityEnabled,omitempty"`
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

func (s *ListVpcEndpointServicesResponseBodyServices) SetResourceGroupId(v string) *ListVpcEndpointServicesResponseBodyServices {
	s.ResourceGroupId = &v
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

func (s *ListVpcEndpointServicesResponseBodyServices) SetServiceSupportIPv6(v bool) *ListVpcEndpointServicesResponseBodyServices {
	s.ServiceSupportIPv6 = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServices) SetServiceType(v string) *ListVpcEndpointServicesResponseBodyServices {
	s.ServiceType = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServices) SetTags(v []*ListVpcEndpointServicesResponseBodyServicesTags) *ListVpcEndpointServicesResponseBodyServices {
	s.Tags = v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServices) SetZoneAffinityEnabled(v bool) *ListVpcEndpointServicesResponseBodyServices {
	s.ZoneAffinityEnabled = &v
	return s
}

type ListVpcEndpointServicesResponseBodyServicesTags struct {
	// The key of the tag added to the resource.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag added to the resource.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListVpcEndpointServicesResponseBodyServicesTags) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointServicesResponseBodyServicesTags) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServicesResponseBodyServicesTags) SetKey(v string) *ListVpcEndpointServicesResponseBodyServicesTags {
	s.Key = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServicesTags) SetValue(v string) *ListVpcEndpointServicesResponseBodyServicesTags {
	s.Value = &v
	return s
}

type ListVpcEndpointServicesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVpcEndpointServicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ListVpcEndpointServicesResponse) SetStatusCode(v int32) *ListVpcEndpointServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVpcEndpointServicesResponse) SetBody(v *ListVpcEndpointServicesResponseBody) *ListVpcEndpointServicesResponse {
	s.Body = v
	return s
}

type ListVpcEndpointServicesByEndUserRequest struct {
	// The number of entries to return on each page. Valid values: **1*	- to **50**. Default value: **50**.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If this is your first request and no next requests are to be performed, you do not need to specify this parameter.
	//
	// 	- If a next request is to be performed, set the value to the value of **NextToken*	- that is returned from the last call.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the endpoint.
	//
	// You can call the [DescribeRegions](~~DescribeRegions~~) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the endpoint service that you want to query.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The name of the endpoint service that you want to query.
	//
	// example:
	//
	// com.aliyuncs.privatelink.cn-huhehaote.epsrv-hp3xdsq46ael67lo****
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The type of the endpoint service.
	//
	// Set the value to **Interface**. You can specify CLB and ALB instances as service resources for the endpoint service.
	//
	// example:
	//
	// Interface
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The list of tags.
	Tag []*ListVpcEndpointServicesByEndUserRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s *ListVpcEndpointServicesByEndUserRequest) SetResourceGroupId(v string) *ListVpcEndpointServicesByEndUserRequest {
	s.ResourceGroupId = &v
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

func (s *ListVpcEndpointServicesByEndUserRequest) SetTag(v []*ListVpcEndpointServicesByEndUserRequestTag) *ListVpcEndpointServicesByEndUserRequest {
	s.Tag = v
	return s
}

type ListVpcEndpointServicesByEndUserRequestTag struct {
	// The key of the tag. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key must be 1 to 64 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListVpcEndpointServicesByEndUserRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointServicesByEndUserRequestTag) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServicesByEndUserRequestTag) SetKey(v string) *ListVpcEndpointServicesByEndUserRequestTag {
	s.Key = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserRequestTag) SetValue(v string) *ListVpcEndpointServicesByEndUserRequestTag {
	s.Value = &v
	return s
}

type ListVpcEndpointServicesByEndUserResponseBody struct {
	// The number of entries returned per page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If no value is returned for **NextToken**, no next requests are performed.
	//
	// 	- If a value is returned for **NextToken**, the value can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The endpoint services.
	Services []*ListVpcEndpointServicesByEndUserResponseBodyServices `json:"Services,omitempty" xml:"Services,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 29
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListVpcEndpointServicesByEndUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointServicesByEndUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServicesByEndUserResponseBody) SetMaxResults(v int32) *ListVpcEndpointServicesByEndUserResponseBody {
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

func (s *ListVpcEndpointServicesByEndUserResponseBody) SetTotalCount(v string) *ListVpcEndpointServicesByEndUserResponseBody {
	s.TotalCount = &v
	return s
}

type ListVpcEndpointServicesByEndUserResponseBodyServices struct {
	// The payer. Valid values:
	//
	// 	- **Endpoint**: the service consumer
	//
	// 	- **EndpointService**: the service provider
	//
	// example:
	//
	// Endpoint
	Payer *string `json:"Payer,omitempty" xml:"Payer,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The domain name of the endpoint service that can be associated with the endpoint.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****.cn-huhehaote.privatelink.aliyuncs.com
	ServiceDomain *string `json:"ServiceDomain,omitempty" xml:"ServiceDomain,omitempty"`
	// The ID of the endpoint service that can be associated with the endpoint.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The name of the endpoint service that can be associated with the endpoint.
	//
	// example:
	//
	// com.aliyuncs.privatelink.cn-huhehaote.epsrv-hp3vpx8yqxblby3i****
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The type of the service resource. Valid values:
	//
	// 	- **slb**: Classic Load Balancer (CLB) instance
	//
	// 	- **alb**: Application Load Balancer (ALB) instance
	//
	// 	- **nlb**: Network Load Balancer (NLB) instance
	//
	// example:
	//
	// slb
	ServiceResourceType *string `json:"ServiceResourceType,omitempty" xml:"ServiceResourceType,omitempty"`
	// Indicates whether IPv6 is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	ServiceSupportIPv6 *bool `json:"ServiceSupportIPv6,omitempty" xml:"ServiceSupportIPv6,omitempty"`
	// The type of the endpoint service.
	//
	// Only **Interface*	- is returned, which indicates an interface endpoint. You can specify **CLB*	- and **ALB*	- instances as service resources.
	//
	// example:
	//
	// Interface
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The list of tags.
	Tags []*ListVpcEndpointServicesByEndUserResponseBodyServicesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The zones of the endpoint service that can be associated with the endpoint.
	Zones []*string `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
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

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) SetResourceGroupId(v string) *ListVpcEndpointServicesByEndUserResponseBodyServices {
	s.ResourceGroupId = &v
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

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) SetServiceResourceType(v string) *ListVpcEndpointServicesByEndUserResponseBodyServices {
	s.ServiceResourceType = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) SetServiceSupportIPv6(v bool) *ListVpcEndpointServicesByEndUserResponseBodyServices {
	s.ServiceSupportIPv6 = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) SetServiceType(v string) *ListVpcEndpointServicesByEndUserResponseBodyServices {
	s.ServiceType = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) SetTags(v []*ListVpcEndpointServicesByEndUserResponseBodyServicesTags) *ListVpcEndpointServicesByEndUserResponseBodyServices {
	s.Tags = v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) SetZones(v []*string) *ListVpcEndpointServicesByEndUserResponseBodyServices {
	s.Zones = v
	return s
}

type ListVpcEndpointServicesByEndUserResponseBodyServicesTags struct {
	// The key of the tag.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListVpcEndpointServicesByEndUserResponseBodyServicesTags) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointServicesByEndUserResponseBodyServicesTags) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServicesTags) SetKey(v string) *ListVpcEndpointServicesByEndUserResponseBodyServicesTags {
	s.Key = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServicesTags) SetValue(v string) *ListVpcEndpointServicesByEndUserResponseBodyServicesTags {
	s.Value = &v
	return s
}

type ListVpcEndpointServicesByEndUserResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVpcEndpointServicesByEndUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ListVpcEndpointServicesByEndUserResponse) SetStatusCode(v int32) *ListVpcEndpointServicesByEndUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponse) SetBody(v *ListVpcEndpointServicesByEndUserResponseBody) *ListVpcEndpointServicesByEndUserResponse {
	s.Body = v
	return s
}

type ListVpcEndpointZonesRequest struct {
	// The ID of the endpoint for which you want to query zones.
	//
	// After you specify an endpoint ID, the system queries the zones of the specified endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// ep-hp33b2e43fays7s8****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The number of entries to return on each page. Valid values: **1*	- to **50**. Default value: **50**.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If this is your first request and no next requests are to be performed, you do not need to specify this parameter.
	//
	// 	- If a next request is to be performed, set the parameter to the value of **NextToken*	- that is returned from the last call.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the endpoint.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The number of entries returned on each page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If no value is returned for **NextToken**, no next requests are performed.
	//
	// 	- If a value is returned for **NextToken**, the value can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The information about the zones.
	Zones []*ListVpcEndpointZonesResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s ListVpcEndpointZonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointZonesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointZonesResponseBody) SetMaxResults(v int32) *ListVpcEndpointZonesResponseBody {
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

func (s *ListVpcEndpointZonesResponseBody) SetTotalCount(v int32) *ListVpcEndpointZonesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVpcEndpointZonesResponseBody) SetZones(v []*ListVpcEndpointZonesResponseBodyZones) *ListVpcEndpointZonesResponseBody {
	s.Zones = v
	return s
}

type ListVpcEndpointZonesResponseBodyZones struct {
	// The ID of the endpoint ENI.
	//
	// example:
	//
	// eni-hp3c8qj1tyct8aqy****
	EniId *string `json:"EniId,omitempty" xml:"EniId,omitempty"`
	// The IP address of the endpoint ENI.
	//
	// example:
	//
	// 192.168.2.23
	EniIp *string `json:"EniIp,omitempty" xml:"EniIp,omitempty"`
	// The region ID of the endpoint.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the vSwitch in the zone. The system automatically creates an endpoint elastic network interface (ENI) in the vSwitch.
	//
	// example:
	//
	// vsw-hjkshjvdkdvd****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The domain name of the zone.
	//
	// After the endpoint in the zone is connected to the endpoint service, you can access the service resources of the endpoint service by using the domain name of the zone.
	//
	// example:
	//
	// ep-hp3f033dp24c5yc9****-cn-huhehaote.epsrv-hp3itcpowf37m3d5****.cn-huhehaote-a.privatelink.aliyuncs.com
	ZoneDomain *string `json:"ZoneDomain,omitempty" xml:"ZoneDomain,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-huhehaote-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// Indicates whether the endpoint service supports IPv6. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	ZoneIpv6Address *string `json:"ZoneIpv6Address,omitempty" xml:"ZoneIpv6Address,omitempty"`
	// The state of the zone. Valid values:
	//
	// 	- **Creating**: The zone is being created.
	//
	// 	- **Wait**: The zone is to be connected.
	//
	// 	- **Connected**: The zone is connected.
	//
	// 	- **Deleting**: The zone is being deleted.
	//
	// 	- **Disconnecting**: The zone is being disconnected.
	//
	// 	- **Disconnected**: The zone is disconnected.
	//
	// 	- **Connecting**: The zone is being connected.
	//
	// example:
	//
	// Wait
	ZoneStatus *string `json:"ZoneStatus,omitempty" xml:"ZoneStatus,omitempty"`
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

func (s *ListVpcEndpointZonesResponseBodyZones) SetZoneIpv6Address(v string) *ListVpcEndpointZonesResponseBodyZones {
	s.ZoneIpv6Address = &v
	return s
}

func (s *ListVpcEndpointZonesResponseBodyZones) SetZoneStatus(v string) *ListVpcEndpointZonesResponseBodyZones {
	s.ZoneStatus = &v
	return s
}

type ListVpcEndpointZonesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVpcEndpointZonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ListVpcEndpointZonesResponse) SetStatusCode(v int32) *ListVpcEndpointZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVpcEndpointZonesResponse) SetBody(v *ListVpcEndpointZonesResponseBody) *ListVpcEndpointZonesResponse {
	s.Body = v
	return s
}

type ListVpcEndpointsRequest struct {
	// The state of the endpoint connection. Valid values:
	//
	// 	- **Pending**: The endpoint connection is being modified.
	//
	// 	- **Connecting**: The endpoint connection is being established.
	//
	// 	- **Connected**: The endpoint connection is established.
	//
	// 	- **Disconnecting**: The endpoint is being disconnected from the endpoint service.
	//
	// 	- **Disconnected**: The endpoint is disconnected from the endpoint service.
	//
	// 	- **Deleting**: The connection is being deleted.
	//
	// 	- **ServiceDeleted**: The corresponding endpoint service has been deleted.
	//
	// example:
	//
	// Disconnected
	ConnectionStatus *string `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	// The ID of the endpoint.
	//
	// example:
	//
	// ep-hp33b2e43fays7s8****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The name of the endpoint.
	//
	// example:
	//
	// test
	EndpointName *string `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
	// The state of the endpoint. Valid values:
	//
	// 	- **Creating**: The endpoint is being created.
	//
	// 	- **Active**: The endpoint is available.
	//
	// 	- **Pending**: The endpoint is being modified.
	//
	// 	- **Deleting**: The endpoint is being deleted.
	//
	// example:
	//
	// Active
	EndpointStatus *string `json:"EndpointStatus,omitempty" xml:"EndpointStatus,omitempty"`
	// The type of the endpoint.
	//
	// Set the value to **Interface**. Then, you can specify Application Load Balancer (ALB) and Classic Load Balancer (CLB) instances as service resources for the endpoint service.
	//
	// example:
	//
	// Interface
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If this is your first request and no next requests are to be performed, you do not need to specify this parameter.
	//
	// 	- If a next request is to be performed, set the parameter to the value of **NextToken*	- that is returned from the last call.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the endpoint.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// 1
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The name of the endpoint service with which the endpoint is associated.
	//
	// example:
	//
	// com.aliyuncs.privatelink.cn-huhehaote.epsrv-hp3vpx8yqxblby3i****
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The list of tags.
	Tag []*ListVpcEndpointsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the VPC to which the endpoint belongs.
	//
	// example:
	//
	// vpc-fdjkf789dfdfdfde****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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

func (s *ListVpcEndpointsRequest) SetResourceGroupId(v string) *ListVpcEndpointsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVpcEndpointsRequest) SetServiceName(v string) *ListVpcEndpointsRequest {
	s.ServiceName = &v
	return s
}

func (s *ListVpcEndpointsRequest) SetTag(v []*ListVpcEndpointsRequestTag) *ListVpcEndpointsRequest {
	s.Tag = v
	return s
}

func (s *ListVpcEndpointsRequest) SetVpcId(v string) *ListVpcEndpointsRequest {
	s.VpcId = &v
	return s
}

type ListVpcEndpointsRequestTag struct {
	// The key of the tag. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key must be 1 to 64 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListVpcEndpointsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointsRequestTag) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointsRequestTag) SetKey(v string) *ListVpcEndpointsRequestTag {
	s.Key = &v
	return s
}

func (s *ListVpcEndpointsRequestTag) SetValue(v string) *ListVpcEndpointsRequestTag {
	s.Value = &v
	return s
}

type ListVpcEndpointsResponseBody struct {
	// The information about the endpoints.
	Endpoints []*ListVpcEndpointsResponseBodyEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If this is your first request and no next requests are to be performed, you do not need to specify this parameter.
	//
	// 	- If a next request is to be performed, set the parameter to the value of **NextToken*	- that is returned from the last call.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 12
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *ListVpcEndpointsResponseBody) SetMaxResults(v int32) *ListVpcEndpointsResponseBody {
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

func (s *ListVpcEndpointsResponseBody) SetTotalCount(v int32) *ListVpcEndpointsResponseBody {
	s.TotalCount = &v
	return s
}

type ListVpcEndpointsResponseBodyEndpoints struct {
	// The bandwidth of the endpoint connection. Unit: Mbit/s.
	//
	// example:
	//
	// 1024
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The state of the endpoint connection. Valid values:
	//
	// 	- **Pending**: The endpoint connection is being modified.
	//
	// 	- **Connecting**: The endpoint connection is being established.
	//
	// 	- **Connected**: The endpoint connection is established.
	//
	// 	- **Disconnecting**: The endpoint is being disconnected from the endpoint service.
	//
	// 	- **Disconnected**: The endpoint is disconnected from the endpoint service.
	//
	// 	- **Deleting**: The endpoint connection is being deleted.
	//
	// 	- **ServiceDeleted**: The corresponding service is deleted.
	//
	// example:
	//
	// Disconnected
	ConnectionStatus *string `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	// The time when the endpoint was created.
	//
	// example:
	//
	// 2021-09-24T18:00:07Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The service state of the endpoint. Valid values:
	//
	// 	- **Normal**: The endpoint runs as expected.
	//
	// 	- **FinancialLocked**: The endpoint is locked due to overdue payments.
	//
	// example:
	//
	// Normal
	EndpointBusinessStatus *string `json:"EndpointBusinessStatus,omitempty" xml:"EndpointBusinessStatus,omitempty"`
	// The description of the endpoint.
	//
	// example:
	//
	// This is my Endpoint.
	EndpointDescription *string `json:"EndpointDescription,omitempty" xml:"EndpointDescription,omitempty"`
	// The domain name of the endpoint.
	//
	// example:
	//
	// ep-hp33b2e43fays7s8****.epsrv-hp3xdsq46ael67lo****.cn-huhehaote.privatelink.aliyuncs.com
	EndpointDomain *string `json:"EndpointDomain,omitempty" xml:"EndpointDomain,omitempty"`
	// The ID of the endpoint.
	//
	// example:
	//
	// ep-hp33b2e43fays7s8****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The name of the endpoint.
	//
	// example:
	//
	// test
	EndpointName *string `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
	// The state of the endpoint. Valid values:
	//
	// 	- **Creating**: The endpoint is being created.
	//
	// 	- **Active**: The endpoint is available.
	//
	// 	- **Pending**: The endpoint is being modified.
	//
	// 	- **Deleting**: The endpoint is being deleted.
	//
	// example:
	//
	// Active
	EndpointStatus *string `json:"EndpointStatus,omitempty" xml:"EndpointStatus,omitempty"`
	// The type of the endpoint.
	//
	// Only **Interface*	- may be returned, which indicates an interface endpoint. You can specify Application Load Balancer (ALB) instances, Classic Load Balancer (CLB) instances, and Network Load Balancer (NLB) instances as service resources.
	//
	// example:
	//
	// Interface
	EndpointType   *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	// The region ID of the endpoint.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// 1
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Indicates whether the endpoint and the endpoint service belong to the same Alibaba Cloud account. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	ResourceOwner *bool `json:"ResourceOwner,omitempty" xml:"ResourceOwner,omitempty"`
	// The ID of the endpoint service that is associated with the endpoint.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The name of the endpoint service that is associated with the endpoint.
	//
	// example:
	//
	// com.aliyuncs.privatelink.cn-huhehaote.epsrv-hp3xdsq46ael67lo****
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The tags added to the resource.
	Tags []*ListVpcEndpointsResponseBodyEndpointsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC) to which the endpoint belongs.
	//
	// example:
	//
	// vpc-hp356stwkxg3fn2xe****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// Indicates whether the domain name of the nearest endpoint that is associated with the endpoint service is resolved first. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	ZoneAffinityEnabled *bool `json:"ZoneAffinityEnabled,omitempty" xml:"ZoneAffinityEnabled,omitempty"`
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

func (s *ListVpcEndpointsResponseBodyEndpoints) SetPolicyDocument(v string) *ListVpcEndpointsResponseBodyEndpoints {
	s.PolicyDocument = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyEndpoints) SetRegionId(v string) *ListVpcEndpointsResponseBodyEndpoints {
	s.RegionId = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyEndpoints) SetResourceGroupId(v string) *ListVpcEndpointsResponseBodyEndpoints {
	s.ResourceGroupId = &v
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

func (s *ListVpcEndpointsResponseBodyEndpoints) SetTags(v []*ListVpcEndpointsResponseBodyEndpointsTags) *ListVpcEndpointsResponseBodyEndpoints {
	s.Tags = v
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

type ListVpcEndpointsResponseBodyEndpointsTags struct {
	// The key of the tag added to the resource.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag added to the resource.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListVpcEndpointsResponseBodyEndpointsTags) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointsResponseBodyEndpointsTags) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointsResponseBodyEndpointsTags) SetKey(v string) *ListVpcEndpointsResponseBodyEndpointsTags {
	s.Key = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyEndpointsTags) SetValue(v string) *ListVpcEndpointsResponseBodyEndpointsTags {
	s.Value = &v
	return s
}

type ListVpcEndpointsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVpcEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ListVpcEndpointsResponse) SetStatusCode(v int32) *ListVpcEndpointsResponse {
	s.StatusCode = &v
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
	// The order ID.
	//
	// example:
	//
	// 3245****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 427688B8-ADFB-4C4E-9D45-EF5C1FD6E23D
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenPrivateLinkServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *OpenPrivateLinkServiceResponse) SetStatusCode(v int32) *OpenPrivateLinkServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenPrivateLinkServiceResponse) SetBody(v *OpenPrivateLinkServiceResponseBody) *OpenPrivateLinkServiceResponse {
	s.Body = v
	return s
}

type RemoveUserFromVpcEndpointServiceRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the AccessKey pair, the permissions of the RAM user, and the required parameters. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The region ID of the endpoint service for which you want to remove the account ID from the whitelist. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// eu-west-1
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The endpoint service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The whitelist in the format of Aliyun Resource Name (ARN).
	//
	// example:
	//
	// acs:ram:*:<account-id>:*
	UserARN *string `json:"UserARN,omitempty" xml:"UserARN,omitempty"`
	// The account ID that you want to remove from the whitelist.
	//
	// example:
	//
	// 12345678
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
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

func (s *RemoveUserFromVpcEndpointServiceRequest) SetUserARN(v string) *RemoveUserFromVpcEndpointServiceRequest {
	s.UserARN = &v
	return s
}

func (s *RemoveUserFromVpcEndpointServiceRequest) SetUserId(v int64) *RemoveUserFromVpcEndpointServiceRequest {
	s.UserId = &v
	return s
}

type RemoveUserFromVpcEndpointServiceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
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
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveUserFromVpcEndpointServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *RemoveUserFromVpcEndpointServiceResponse) SetStatusCode(v int32) *RemoveUserFromVpcEndpointServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveUserFromVpcEndpointServiceResponse) SetBody(v *RemoveUserFromVpcEndpointServiceResponseBody) *RemoveUserFromVpcEndpointServiceResponse {
	s.Body = v
	return s
}

type RemoveZoneFromVpcEndpointRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the endpoint for which you want to delete the zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// ep-hp33b2e43fays7s8****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The region ID of the endpoint for which you want to delete the zone. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// eu-west-1
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the zone that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// eu-west-1a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveZoneFromVpcEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *RemoveZoneFromVpcEndpointResponse) SetStatusCode(v int32) *RemoveZoneFromVpcEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveZoneFromVpcEndpointResponse) SetBody(v *RemoveZoneFromVpcEndpointResponseBody) *RemoveZoneFromVpcEndpointResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token.*******	- The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The region ID of the PrivateLink instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource IDs. Up to 50 resource IDs are supported.
	//
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of resource. Valid values:
	//
	// 	- **vpcendpoint**: endpoint
	//
	// 	- **vpcendpointservice**: endpoint service
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags to add to the resources.
	//
	// This parameter is required.
	Tag []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetClientToken(v string) *TagResourcesRequest {
	s.ClientToken = &v
	return s
}

func (s *TagResourcesRequest) SetDryRun(v bool) *TagResourcesRequest {
	s.DryRun = &v
	return s
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
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
	// The key of tag N to add to the resource. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `aliyun` or `acs:`.
	//
	// This parameter is required.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the resource. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag value cannot start with `acs:` or `aliyun`.
	//
	// This parameter is required.
	//
	// example:
	//
	// FinanceJoshua
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
	// The request ID.
	//
	// example:
	//
	// C46FF5A8-C5F0-4024-8262-B16B639225A0
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	// Specifies whether to remove all tags from the resource. Valid values:
	//
	// 	- **true**: removes all tags from the resource.
	//
	// 	- **false**: does not remove all tags from the resource.
	//
	// >  If you specify both this parameter and **TagKey**, this parameter is invalid.
	//
	// example:
	//
	// true
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token.*******	- The request ID may be different for each request.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a `2xx HTTP` status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The region ID of the PrivateLink instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource IDs. You can specify up to 50 resource IDs.
	//
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The resource type. Valid values:
	//
	// 	- **vpcendpoint**: endpoint
	//
	// 	- **vpcendpointservice**: endpoint service
	//
	// This parameter is required.
	//
	// example:
	//
	// slb
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The keys of the tags that you want to remove from the resource. You can specify up to 20 tag keys.
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
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

func (s *UntagResourcesRequest) SetClientToken(v string) *UntagResourcesRequest {
	s.ClientToken = &v
	return s
}

func (s *UntagResourcesRequest) SetDryRun(v bool) *UntagResourcesRequest {
	s.DryRun = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
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
	// The ID of the request.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type UpdateVpcEndpointAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The description of the endpoint.
	//
	// The description must be 2 to 256 characters in length. It cannot start with `http://` or `https://`.
	//
	// example:
	//
	// This is my Endpoint.
	EndpointDescription *string `json:"EndpointDescription,omitempty" xml:"EndpointDescription,omitempty"`
	// The endpoint ID whose attributes you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// ep-hp33b2e43fays7s8****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The name of the endpoint.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// test
	EndpointName   *string `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	// The region ID of the endpoint whose attributes you want to modify. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// eu-west-1
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *UpdateVpcEndpointAttributeRequest) SetPolicyDocument(v string) *UpdateVpcEndpointAttributeRequest {
	s.PolicyDocument = &v
	return s
}

func (s *UpdateVpcEndpointAttributeRequest) SetRegionId(v string) *UpdateVpcEndpointAttributeRequest {
	s.RegionId = &v
	return s
}

type UpdateVpcEndpointAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVpcEndpointAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *UpdateVpcEndpointAttributeResponse) SetStatusCode(v int32) *UpdateVpcEndpointAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVpcEndpointAttributeResponse) SetBody(v *UpdateVpcEndpointAttributeResponseBody) *UpdateVpcEndpointAttributeResponse {
	s.Body = v
	return s
}

type UpdateVpcEndpointConnectionAttributeRequest struct {
	// The bandwidth of the endpoint connection that you want to modify. Unit: Mbit/s. Valid values: **3072*	- to **10240**.
	//
	// >  The bandwidth of an endpoint connection is in the range of **100*	- to **10,240*	- Mbit/s. The default bandwidth is **3,072*	- Mbit/s. When the endpoint is connected to the endpoint service, the default bandwidth is the minimum bandwidth. In this case, the connection bandwidth range is **3,072*	- to **10,240*	- Mbit/s. If Classic Load Balancer (CLB) instances or Application Load Balancer (ALB) instances are specified as service resources, you can modify the endpoint connection bandwidth based on your business requirements. This parameter is invalid if Network Load Balancer (NLB) instances are specified as service resources.
	//
	// example:
	//
	// 1000
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The endpoint ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ep-hp33b2e43fays7s8****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The region ID of the endpoint connection whose bandwidth you want to modify. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// eu-west-1
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The endpoint service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
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
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVpcEndpointConnectionAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *UpdateVpcEndpointConnectionAttributeResponse) SetStatusCode(v int32) *UpdateVpcEndpointConnectionAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVpcEndpointConnectionAttributeResponse) SetBody(v *UpdateVpcEndpointConnectionAttributeResponseBody) *UpdateVpcEndpointConnectionAttributeResponse {
	s.Body = v
	return s
}

type UpdateVpcEndpointServiceAttributeRequest struct {
	// Specifies whether to automatically accept endpoint connection requests. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	AutoAcceptEnabled *bool `json:"AutoAcceptEnabled,omitempty" xml:"AutoAcceptEnabled,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The default maximum bandwidth of the endpoint connection. Unit: Mbit/s. Default value: **3072**.
	//
	// Valid values: **100*	- to **10240**.
	//
	// >  You can specify this parameter only if you specify Classic Load Balancer (CLB) instances or Application Load Balancer (ALB) instances as service resources.
	//
	// example:
	//
	// 200
	ConnectBandwidth *int32 `json:"ConnectBandwidth,omitempty" xml:"ConnectBandwidth,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The region ID of the endpoint service.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The description of the endpoint service.
	//
	// example:
	//
	// This is my EndpointService.
	ServiceDescription *string `json:"ServiceDescription,omitempty" xml:"ServiceDescription,omitempty"`
	// The endpoint service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// Specifies whether to enable IPv6. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	ServiceSupportIPv6 *bool `json:"ServiceSupportIPv6,omitempty" xml:"ServiceSupportIPv6,omitempty"`
	// Specifies whether to first resolve the domain name of the nearest endpoint that is associated with the endpoint service. Valid values:
	//
	// 	- **true*	- (default)
	//
	// 	- **false**
	//
	// example:
	//
	// true
	ZoneAffinityEnabled *bool `json:"ZoneAffinityEnabled,omitempty" xml:"ZoneAffinityEnabled,omitempty"`
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

func (s *UpdateVpcEndpointServiceAttributeRequest) SetServiceSupportIPv6(v bool) *UpdateVpcEndpointServiceAttributeRequest {
	s.ServiceSupportIPv6 = &v
	return s
}

func (s *UpdateVpcEndpointServiceAttributeRequest) SetZoneAffinityEnabled(v bool) *UpdateVpcEndpointServiceAttributeRequest {
	s.ZoneAffinityEnabled = &v
	return s
}

type UpdateVpcEndpointServiceAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8D8992C1-6712-423C-BAC5-E5E817484C6B
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
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVpcEndpointServiceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *UpdateVpcEndpointServiceAttributeResponse) SetStatusCode(v int32) *UpdateVpcEndpointServiceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVpcEndpointServiceAttributeResponse) SetBody(v *UpdateVpcEndpointServiceAttributeResponseBody) *UpdateVpcEndpointServiceAttributeResponse {
	s.Body = v
	return s
}

type UpdateVpcEndpointServiceResourceAttributeRequest struct {
	// Specifies whether to enable automatic resource allocation. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	AutoAllocatedEnabled *bool `json:"AutoAllocatedEnabled,omitempty" xml:"AutoAllocatedEnabled,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// 	- **true**: performs a dry run. The system checks the required parameters, request syntax, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false**: performs a dry run and sends the request. If the request passes the dry run, an HTTP 2xx status code is returned and the operation is performed. This is the default value.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the region where the service resource is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service resource ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-hp32z1wp5peaoox2q****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The endpoint service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The zone ID of the service resource.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s UpdateVpcEndpointServiceResourceAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateVpcEndpointServiceResourceAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateVpcEndpointServiceResourceAttributeRequest) SetAutoAllocatedEnabled(v bool) *UpdateVpcEndpointServiceResourceAttributeRequest {
	s.AutoAllocatedEnabled = &v
	return s
}

func (s *UpdateVpcEndpointServiceResourceAttributeRequest) SetClientToken(v string) *UpdateVpcEndpointServiceResourceAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateVpcEndpointServiceResourceAttributeRequest) SetDryRun(v bool) *UpdateVpcEndpointServiceResourceAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateVpcEndpointServiceResourceAttributeRequest) SetRegionId(v string) *UpdateVpcEndpointServiceResourceAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateVpcEndpointServiceResourceAttributeRequest) SetResourceId(v string) *UpdateVpcEndpointServiceResourceAttributeRequest {
	s.ResourceId = &v
	return s
}

func (s *UpdateVpcEndpointServiceResourceAttributeRequest) SetServiceId(v string) *UpdateVpcEndpointServiceResourceAttributeRequest {
	s.ServiceId = &v
	return s
}

func (s *UpdateVpcEndpointServiceResourceAttributeRequest) SetZoneId(v string) *UpdateVpcEndpointServiceResourceAttributeRequest {
	s.ZoneId = &v
	return s
}

type UpdateVpcEndpointServiceResourceAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateVpcEndpointServiceResourceAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateVpcEndpointServiceResourceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVpcEndpointServiceResourceAttributeResponseBody) SetRequestId(v string) *UpdateVpcEndpointServiceResourceAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateVpcEndpointServiceResourceAttributeResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVpcEndpointServiceResourceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVpcEndpointServiceResourceAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateVpcEndpointServiceResourceAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateVpcEndpointServiceResourceAttributeResponse) SetHeaders(v map[string]*string) *UpdateVpcEndpointServiceResourceAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateVpcEndpointServiceResourceAttributeResponse) SetStatusCode(v int32) *UpdateVpcEndpointServiceResourceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVpcEndpointServiceResourceAttributeResponse) SetBody(v *UpdateVpcEndpointServiceResourceAttributeResponseBody) *UpdateVpcEndpointServiceResourceAttributeResponse {
	s.Body = v
	return s
}

type UpdateVpcEndpointZoneConnectionResourceAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The endpoint ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ep-hp33b2e43fays7s8****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The region ID of the endpoint connection whose bandwidth you want to modify.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource allocation mode. You can change the resource allocation mode only if the endpoint connection is in the **Disconnected*	- state. Valid values:
	//
	// 	- **Auto**: automatically and randomly allocates service resources. In this mode, the specified service resource is deleted.
	//
	// 	- **Manual**: manually allocates service resources. If you set the value to Manual, you must also specify the **ResourceId*	- and **ResourceType*	- parameters.
	//
	// example:
	//
	// Auto
	ResourceAllocateMode *string `json:"ResourceAllocateMode,omitempty" xml:"ResourceAllocateMode,omitempty"`
	// The ID of the service resource that you want to manually allocate or migrate in the zone where the endpoint connection is deployed.
	//
	// > If **ResourceAllocateMode*	- is set to **Mannual**, or **ResourceReplaceMode*	- is set, this parameter is required.
	//
	// example:
	//
	// lb-hp32z1wp5peaoox2q****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The migration mode of the service resource. Valid values:
	//
	// 	- **Graceful**: smooth migration. Service resources in the zone are smoothly migrated.
	//
	// 	- **Force**: forced migration. Service resources in the zone are forcefully migrated.
	//
	// >  You need to specify this parameter only if you want to migrate service resources and the endpoint connection is in the **Connected*	- state. If you specify this parameter, you must also specify the **ResourceId*	- and **ResourceType*	- parameters.
	//
	// example:
	//
	// Graceful
	ResourceReplaceMode *string `json:"ResourceReplaceMode,omitempty" xml:"ResourceReplaceMode,omitempty"`
	// The type of the service resource. Valid values:
	//
	// 	- **slb**: a CLB instance that supports PrivateLink. In addition, the CLB instance is deployed in a VPC.
	//
	// 	- **alb**: an Application Load Balancer (ALB) instance that supports PrivateLink. In addition, the ALB instance is deployed in a VPC.
	//
	// > If **ResourceAllocateMode*	- is set to **Mannual**, or **ResourceReplaceMode*	- is set, this parameter is required.
	//
	// example:
	//
	// slb
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The endpoint service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The zone ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s UpdateVpcEndpointZoneConnectionResourceAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateVpcEndpointZoneConnectionResourceAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateVpcEndpointZoneConnectionResourceAttributeRequest) SetClientToken(v string) *UpdateVpcEndpointZoneConnectionResourceAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateVpcEndpointZoneConnectionResourceAttributeRequest) SetDryRun(v bool) *UpdateVpcEndpointZoneConnectionResourceAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateVpcEndpointZoneConnectionResourceAttributeRequest) SetEndpointId(v string) *UpdateVpcEndpointZoneConnectionResourceAttributeRequest {
	s.EndpointId = &v
	return s
}

func (s *UpdateVpcEndpointZoneConnectionResourceAttributeRequest) SetRegionId(v string) *UpdateVpcEndpointZoneConnectionResourceAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateVpcEndpointZoneConnectionResourceAttributeRequest) SetResourceAllocateMode(v string) *UpdateVpcEndpointZoneConnectionResourceAttributeRequest {
	s.ResourceAllocateMode = &v
	return s
}

func (s *UpdateVpcEndpointZoneConnectionResourceAttributeRequest) SetResourceId(v string) *UpdateVpcEndpointZoneConnectionResourceAttributeRequest {
	s.ResourceId = &v
	return s
}

func (s *UpdateVpcEndpointZoneConnectionResourceAttributeRequest) SetResourceReplaceMode(v string) *UpdateVpcEndpointZoneConnectionResourceAttributeRequest {
	s.ResourceReplaceMode = &v
	return s
}

func (s *UpdateVpcEndpointZoneConnectionResourceAttributeRequest) SetResourceType(v string) *UpdateVpcEndpointZoneConnectionResourceAttributeRequest {
	s.ResourceType = &v
	return s
}

func (s *UpdateVpcEndpointZoneConnectionResourceAttributeRequest) SetServiceId(v string) *UpdateVpcEndpointZoneConnectionResourceAttributeRequest {
	s.ServiceId = &v
	return s
}

func (s *UpdateVpcEndpointZoneConnectionResourceAttributeRequest) SetZoneId(v string) *UpdateVpcEndpointZoneConnectionResourceAttributeRequest {
	s.ZoneId = &v
	return s
}

type UpdateVpcEndpointZoneConnectionResourceAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateVpcEndpointZoneConnectionResourceAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateVpcEndpointZoneConnectionResourceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVpcEndpointZoneConnectionResourceAttributeResponseBody) SetRequestId(v string) *UpdateVpcEndpointZoneConnectionResourceAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateVpcEndpointZoneConnectionResourceAttributeResponse struct {
	Headers    map[string]*string                                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVpcEndpointZoneConnectionResourceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVpcEndpointZoneConnectionResourceAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateVpcEndpointZoneConnectionResourceAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateVpcEndpointZoneConnectionResourceAttributeResponse) SetHeaders(v map[string]*string) *UpdateVpcEndpointZoneConnectionResourceAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateVpcEndpointZoneConnectionResourceAttributeResponse) SetStatusCode(v int32) *UpdateVpcEndpointZoneConnectionResourceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVpcEndpointZoneConnectionResourceAttributeResponse) SetBody(v *UpdateVpcEndpointZoneConnectionResourceAttributeResponseBody) *UpdateVpcEndpointZoneConnectionResourceAttributeResponse {
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

// Summary:
//
// Adds an account ID to the whitelist of an endpoint service.
//
// Description:
//
//   Before you add an account ID to the whitelist of an endpoint service, make sure that the endpoint service is in the **Active*	- state. You can call the [GetVpcEndpointServiceAttribute](https://help.aliyun.com/document_detail/469330.html) operation to query the status of the endpoint service.
//
// 	- You cannot repeatedly call the **AddUserToVpcEndpointService*	- operation to add the ID of an Alibaba Cloud account to the whitelist of an endpoint service within a specified period of time.
//
// @param request - AddUserToVpcEndpointServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserToVpcEndpointServiceResponse
func (client *Client) AddUserToVpcEndpointServiceWithOptions(request *AddUserToVpcEndpointServiceRequest, runtime *util.RuntimeOptions) (_result *AddUserToVpcEndpointServiceResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.UserARN)) {
		query["UserARN"] = request.UserARN
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

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

// Summary:
//
// Adds an account ID to the whitelist of an endpoint service.
//
// Description:
//
//   Before you add an account ID to the whitelist of an endpoint service, make sure that the endpoint service is in the **Active*	- state. You can call the [GetVpcEndpointServiceAttribute](https://help.aliyun.com/document_detail/469330.html) operation to query the status of the endpoint service.
//
// 	- You cannot repeatedly call the **AddUserToVpcEndpointService*	- operation to add the ID of an Alibaba Cloud account to the whitelist of an endpoint service within a specified period of time.
//
// @param request - AddUserToVpcEndpointServiceRequest
//
// @return AddUserToVpcEndpointServiceResponse
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

// Summary:
//
// Adds a zone to an endpoint.
//
// Description:
//
//   **AddZoneToVpcEndpoint*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [ListVpcEndpointZones](https://help.aliyun.com/document_detail/183560.html) operation to query the state of the zone.
//
//     	- If the zone is in the **Creating*	- state, the zone is being added.
//
//     	- If the zone is in the Wait state, the zone is added.
//
// 	- You cannot repeatedly call the **AddZoneToVpcEndpoint*	- operation to add a zone to an endpoint within a specified period of time.
//
// @param request - AddZoneToVpcEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddZoneToVpcEndpointResponse
func (client *Client) AddZoneToVpcEndpointWithOptions(request *AddZoneToVpcEndpointRequest, runtime *util.RuntimeOptions) (_result *AddZoneToVpcEndpointResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.EndpointId)) {
		query["EndpointId"] = request.EndpointId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["ip"] = request.Ip
	}

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

// Summary:
//
// Adds a zone to an endpoint.
//
// Description:
//
//   **AddZoneToVpcEndpoint*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [ListVpcEndpointZones](https://help.aliyun.com/document_detail/183560.html) operation to query the state of the zone.
//
//     	- If the zone is in the **Creating*	- state, the zone is being added.
//
//     	- If the zone is in the Wait state, the zone is added.
//
// 	- You cannot repeatedly call the **AddZoneToVpcEndpoint*	- operation to add a zone to an endpoint within a specified period of time.
//
// @param request - AddZoneToVpcEndpointRequest
//
// @return AddZoneToVpcEndpointResponse
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

// Summary:
//
// Adds a service resource to an endpoint service.
//
// Description:
//
//   Before you add a service resource to an endpoint service, make sure that the endpoint service is in the **Active*	- state. You can call the [GetVpcEndpointServiceAttribute](https://help.aliyun.com/document_detail/469330.html) operation to query the status of the endpoint service.
//
// 	- You cannot repeatedly call the **AttachResourceToVpcEndpointService*	- operation to add a service resource to an endpoint service within a specified period of time.
//
// @param request - AttachResourceToVpcEndpointServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachResourceToVpcEndpointServiceResponse
func (client *Client) AttachResourceToVpcEndpointServiceWithOptions(request *AttachResourceToVpcEndpointServiceRequest, runtime *util.RuntimeOptions) (_result *AttachResourceToVpcEndpointServiceResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

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

// Summary:
//
// Adds a service resource to an endpoint service.
//
// Description:
//
//   Before you add a service resource to an endpoint service, make sure that the endpoint service is in the **Active*	- state. You can call the [GetVpcEndpointServiceAttribute](https://help.aliyun.com/document_detail/469330.html) operation to query the status of the endpoint service.
//
// 	- You cannot repeatedly call the **AttachResourceToVpcEndpointService*	- operation to add a service resource to an endpoint service within a specified period of time.
//
// @param request - AttachResourceToVpcEndpointServiceRequest
//
// @return AttachResourceToVpcEndpointServiceResponse
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

// Summary:
//
// Associates an endpoint with a security group.
//
// Description:
//
//   **AttachSecurityGroupToVpcEndpoint*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [ListVpcEndpoints](https://help.aliyun.com/document_detail/183558.html) operation to query the state of the endpoint.
//
//     	- If the endpoint is in the **Pending*	- state, the endpoint is being associated with the security group.
//
//     	- If the endpoint is in the **Active*	- state, the endpoint is associated with the security group.
//
// 	- You cannot repeatedly call the **AttachSecurityGroupToVpcEndpoint*	- operation to associate an endpoint with a security group within a specified period of time.
//
// @param request - AttachSecurityGroupToVpcEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachSecurityGroupToVpcEndpointResponse
func (client *Client) AttachSecurityGroupToVpcEndpointWithOptions(request *AttachSecurityGroupToVpcEndpointRequest, runtime *util.RuntimeOptions) (_result *AttachSecurityGroupToVpcEndpointResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.EndpointId)) {
		query["EndpointId"] = request.EndpointId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

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

// Summary:
//
// Associates an endpoint with a security group.
//
// Description:
//
//   **AttachSecurityGroupToVpcEndpoint*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [ListVpcEndpoints](https://help.aliyun.com/document_detail/183558.html) operation to query the state of the endpoint.
//
//     	- If the endpoint is in the **Pending*	- state, the endpoint is being associated with the security group.
//
//     	- If the endpoint is in the **Active*	- state, the endpoint is associated with the security group.
//
// 	- You cannot repeatedly call the **AttachSecurityGroupToVpcEndpoint*	- operation to associate an endpoint with a security group within a specified period of time.
//
// @param request - AttachSecurityGroupToVpcEndpointRequest
//
// @return AttachSecurityGroupToVpcEndpointResponse
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

// Summary:
//
// Modifies a resource group.
//
// @param request - ChangeResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroupWithOptions(request *ChangeResourceGroupRequest, runtime *util.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeResourceGroup"),
		Version:     tea.String("2020-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a resource group.
//
// @param request - ChangeResourceGroupRequest
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroup(request *ChangeResourceGroupRequest) (_result *ChangeResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.ChangeResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries whether PrivateLink is activated.
//
// @param request - CheckProductOpenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckProductOpenResponse
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

// Summary:
//
// Queries whether PrivateLink is activated.
//
// @return CheckProductOpenResponse
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

// Summary:
//
// Creates an endpoint.
//
// Description:
//
// *CreateVpcEndpoint*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [GetVpcEndpointAttribute](https://help.aliyun.com/document_detail/183568.html) operation to check whether the endpoint is created.
//
// 	- If the endpoint is in the **Creating*	- state, the endpoint is being created.
//
// 	- If the endpoint is in the **Active*	- state, the endpoint is created.
//
// @param request - CreateVpcEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVpcEndpointResponse
func (client *Client) CreateVpcEndpointWithOptions(request *CreateVpcEndpointRequest, runtime *util.RuntimeOptions) (_result *CreateVpcEndpointResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.EndpointDescription)) {
		query["EndpointDescription"] = request.EndpointDescription
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointName)) {
		query["EndpointName"] = request.EndpointName
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointType)) {
		query["EndpointType"] = request.EndpointType
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyDocument)) {
		query["PolicyDocument"] = request.PolicyDocument
	}

	if !tea.BoolValue(util.IsUnset(request.ProtectedEnabled)) {
		query["ProtectedEnabled"] = request.ProtectedEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.Zone)) {
		query["Zone"] = request.Zone
	}

	if !tea.BoolValue(util.IsUnset(request.ZonePrivateIpAddressCount)) {
		query["ZonePrivateIpAddressCount"] = request.ZonePrivateIpAddressCount
	}

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

// Summary:
//
// Creates an endpoint.
//
// Description:
//
// *CreateVpcEndpoint*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [GetVpcEndpointAttribute](https://help.aliyun.com/document_detail/183568.html) operation to check whether the endpoint is created.
//
// 	- If the endpoint is in the **Creating*	- state, the endpoint is being created.
//
// 	- If the endpoint is in the **Active*	- state, the endpoint is created.
//
// @param request - CreateVpcEndpointRequest
//
// @return CreateVpcEndpointResponse
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

// Summary:
//
// Creates an endpoint service.
//
// Description:
//
//   Before you create an endpoint service, make sure that you have created a Server Load Balancer (SLB) instance that supports PrivateLink. For more information, see [CreateLoadBalancer](https://help.aliyun.com/document_detail/174064.html).
//
// 	- **CreateVpcEndpointService*	- is an asynchronous operation. After a request is sent, the system returns a request ID and an instance ID and runs the task in the background. You can call the [GetVpcEndpointServiceAttribute](https://help.aliyun.com/document_detail/183542.html) operation to query the status of the endpoint service.
//
//     	- If the endpoint service is in the **Creating*	- state, the endpoint service is being created.
//
//     	- If the endpoint service is in the **Active*	- state, the endpoint service is created.
//
// @param request - CreateVpcEndpointServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVpcEndpointServiceResponse
func (client *Client) CreateVpcEndpointServiceWithOptions(request *CreateVpcEndpointServiceRequest, runtime *util.RuntimeOptions) (_result *CreateVpcEndpointServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoAcceptEnabled)) {
		query["AutoAcceptEnabled"] = request.AutoAcceptEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.Payer)) {
		query["Payer"] = request.Payer
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Resource)) {
		query["Resource"] = request.Resource
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceDescription)) {
		query["ServiceDescription"] = request.ServiceDescription
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceResourceType)) {
		query["ServiceResourceType"] = request.ServiceResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceSupportIPv6)) {
		query["ServiceSupportIPv6"] = request.ServiceSupportIPv6
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneAffinityEnabled)) {
		query["ZoneAffinityEnabled"] = request.ZoneAffinityEnabled
	}

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

// Summary:
//
// Creates an endpoint service.
//
// Description:
//
//   Before you create an endpoint service, make sure that you have created a Server Load Balancer (SLB) instance that supports PrivateLink. For more information, see [CreateLoadBalancer](https://help.aliyun.com/document_detail/174064.html).
//
// 	- **CreateVpcEndpointService*	- is an asynchronous operation. After a request is sent, the system returns a request ID and an instance ID and runs the task in the background. You can call the [GetVpcEndpointServiceAttribute](https://help.aliyun.com/document_detail/183542.html) operation to query the status of the endpoint service.
//
//     	- If the endpoint service is in the **Creating*	- state, the endpoint service is being created.
//
//     	- If the endpoint service is in the **Active*	- state, the endpoint service is created.
//
// @param request - CreateVpcEndpointServiceRequest
//
// @return CreateVpcEndpointServiceResponse
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

// Summary:
//
// Deletes an endpoint.
//
// Description:
//
//   Before you delete an endpoint, you must delete the zones that are added to the endpoint.
//
// 	- **DeleteVpcEndpoint*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [GetVpcEndpointAttribute](https://help.aliyun.com/document_detail/183568.html) operation to check whether the endpoint is deleted.
//
//     	- If the endpoint is in the **Deleting*	- state, the endpoint is being deleted.
//
//     	- If the endpoint cannot be queried, the endpoint is deleted.
//
// @param request - DeleteVpcEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVpcEndpointResponse
func (client *Client) DeleteVpcEndpointWithOptions(request *DeleteVpcEndpointRequest, runtime *util.RuntimeOptions) (_result *DeleteVpcEndpointResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.EndpointId)) {
		query["EndpointId"] = request.EndpointId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

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

// Summary:
//
// Deletes an endpoint.
//
// Description:
//
//   Before you delete an endpoint, you must delete the zones that are added to the endpoint.
//
// 	- **DeleteVpcEndpoint*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [GetVpcEndpointAttribute](https://help.aliyun.com/document_detail/183568.html) operation to check whether the endpoint is deleted.
//
//     	- If the endpoint is in the **Deleting*	- state, the endpoint is being deleted.
//
//     	- If the endpoint cannot be queried, the endpoint is deleted.
//
// @param request - DeleteVpcEndpointRequest
//
// @return DeleteVpcEndpointResponse
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

// Summary:
//
// Deletes an endpoint service.
//
// Description:
//
//   Before you delete an endpoint service, you must disconnect the endpoint from the endpoint service and remove the service resources.
//
// 	- **DeleteVpcEndpointService*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [GetVpcEndpointServiceAttribute](https://help.aliyun.com/document_detail/183542.html) operation to check whether the endpoint service is deleted.
//
//     	- If the endpoint service is in the **Deleting*	- state, the endpoint service is being deleted.
//
//     	- If the endpoint service cannot be queried, the endpoint service is deleted.
//
// 	- You cannot repeatedly call the **DeleteVpcEndpointService*	- operation to delete an endpoint service that belongs to an Alibaba Cloud account within a specified period of time.
//
// @param request - DeleteVpcEndpointServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVpcEndpointServiceResponse
func (client *Client) DeleteVpcEndpointServiceWithOptions(request *DeleteVpcEndpointServiceRequest, runtime *util.RuntimeOptions) (_result *DeleteVpcEndpointServiceResponse, _err error) {
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

// Summary:
//
// Deletes an endpoint service.
//
// Description:
//
//   Before you delete an endpoint service, you must disconnect the endpoint from the endpoint service and remove the service resources.
//
// 	- **DeleteVpcEndpointService*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [GetVpcEndpointServiceAttribute](https://help.aliyun.com/document_detail/183542.html) operation to check whether the endpoint service is deleted.
//
//     	- If the endpoint service is in the **Deleting*	- state, the endpoint service is being deleted.
//
//     	- If the endpoint service cannot be queried, the endpoint service is deleted.
//
// 	- You cannot repeatedly call the **DeleteVpcEndpointService*	- operation to delete an endpoint service that belongs to an Alibaba Cloud account within a specified period of time.
//
// @param request - DeleteVpcEndpointServiceRequest
//
// @return DeleteVpcEndpointServiceResponse
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

// Summary:
//
// Queries available regions and zones.
//
// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

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

// Summary:
//
// Queries available regions and zones.
//
// @param request - DescribeRegionsRequest
//
// @return DescribeRegionsResponse
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

// Summary:
//
// Queries a list of zones in a specified region.
//
// @param request - DescribeZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeZonesResponse
func (client *Client) DescribeZonesWithOptions(request *DescribeZonesRequest, runtime *util.RuntimeOptions) (_result *DescribeZonesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

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

// Summary:
//
// Queries a list of zones in a specified region.
//
// @param request - DescribeZonesRequest
//
// @return DescribeZonesResponse
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

// Summary:
//
// Removes a service resource from an endpoint service.
//
// Description:
//
//   Before you remove a service resource from an endpoint service, make sure that the endpoint service is in the **Active*	- state. You can call the [GetVpcEndpointServiceAttribute](https://help.aliyun.com/document_detail/469330.html) operation to query the status of the endpoint service.
//
// 	- You cannot repeatedly call the **DetachResourceFromVpcEndpointService*	- operation to remove a service resource from an endpoint service within a specified period of time.
//
// @param request - DetachResourceFromVpcEndpointServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachResourceFromVpcEndpointServiceResponse
func (client *Client) DetachResourceFromVpcEndpointServiceWithOptions(request *DetachResourceFromVpcEndpointServiceRequest, runtime *util.RuntimeOptions) (_result *DetachResourceFromVpcEndpointServiceResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

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

// Summary:
//
// Removes a service resource from an endpoint service.
//
// Description:
//
//   Before you remove a service resource from an endpoint service, make sure that the endpoint service is in the **Active*	- state. You can call the [GetVpcEndpointServiceAttribute](https://help.aliyun.com/document_detail/469330.html) operation to query the status of the endpoint service.
//
// 	- You cannot repeatedly call the **DetachResourceFromVpcEndpointService*	- operation to remove a service resource from an endpoint service within a specified period of time.
//
// @param request - DetachResourceFromVpcEndpointServiceRequest
//
// @return DetachResourceFromVpcEndpointServiceResponse
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

// Summary:
//
// Disassociates an endpoint from a security group.
//
// Description:
//
//   **DetachSecurityGroupFromVpcEndpoint*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [ListVpcEndpoints](https://help.aliyun.com/document_detail/183558.html) to check whether the endpoint is disassociated from the security group.
//
//     	- If the endpoint is in the **Pending*	- state, the endpoint is being disassociated from the security group.
//
//     	- If you cannot query the endpoint in the security group, the endpoint is disassociated from the security group.
//
// 	- You cannot repeatedly call the **DetachSecurityGroupFromVpcEndpoint*	- operation to disassociate an endpoint from a security group within a specified period of time.
//
// @param request - DetachSecurityGroupFromVpcEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachSecurityGroupFromVpcEndpointResponse
func (client *Client) DetachSecurityGroupFromVpcEndpointWithOptions(request *DetachSecurityGroupFromVpcEndpointRequest, runtime *util.RuntimeOptions) (_result *DetachSecurityGroupFromVpcEndpointResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.EndpointId)) {
		query["EndpointId"] = request.EndpointId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

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

// Summary:
//
// Disassociates an endpoint from a security group.
//
// Description:
//
//   **DetachSecurityGroupFromVpcEndpoint*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [ListVpcEndpoints](https://help.aliyun.com/document_detail/183558.html) to check whether the endpoint is disassociated from the security group.
//
//     	- If the endpoint is in the **Pending*	- state, the endpoint is being disassociated from the security group.
//
//     	- If you cannot query the endpoint in the security group, the endpoint is disassociated from the security group.
//
// 	- You cannot repeatedly call the **DetachSecurityGroupFromVpcEndpoint*	- operation to disassociate an endpoint from a security group within a specified period of time.
//
// @param request - DetachSecurityGroupFromVpcEndpointRequest
//
// @return DetachSecurityGroupFromVpcEndpointResponse
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

// Summary:
//
// Rejects a connection request from an endpoint.
//
// Description:
//
//   **DisableVpcEndpointConnection*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [GetVpcEndpointAttribute](https://help.aliyun.com/document_detail/183568.html) operation to query the state of the endpoint connection.
//
//     	- If the endpoint connection is in the **Disconnecting*	- state, the endpoint is being disconnected from the endpoint service.
//
//     	- If the endpoint connection is in the **Disconnected*	- state, the endpoint is disconnected from the endpoint service.
//
// 	- You cannot repeatedly call the **DisableVpcEndpointConnection*	- operation to allow an endpoint service to reject a connection request from an endpoint within a specified period of time.
//
// @param request - DisableVpcEndpointConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableVpcEndpointConnectionResponse
func (client *Client) DisableVpcEndpointConnectionWithOptions(request *DisableVpcEndpointConnectionRequest, runtime *util.RuntimeOptions) (_result *DisableVpcEndpointConnectionResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.EndpointId)) {
		query["EndpointId"] = request.EndpointId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
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

// Summary:
//
// Rejects a connection request from an endpoint.
//
// Description:
//
//   **DisableVpcEndpointConnection*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [GetVpcEndpointAttribute](https://help.aliyun.com/document_detail/183568.html) operation to query the state of the endpoint connection.
//
//     	- If the endpoint connection is in the **Disconnecting*	- state, the endpoint is being disconnected from the endpoint service.
//
//     	- If the endpoint connection is in the **Disconnected*	- state, the endpoint is disconnected from the endpoint service.
//
// 	- You cannot repeatedly call the **DisableVpcEndpointConnection*	- operation to allow an endpoint service to reject a connection request from an endpoint within a specified period of time.
//
// @param request - DisableVpcEndpointConnectionRequest
//
// @return DisableVpcEndpointConnectionResponse
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

// Summary:
//
// Disconnects an endpoint from a connection in the specified zone.
//
// Description:
//
//   You can call this operation only when the state of the endpoint is **Connected*	- and the state of the zone associated with the endpoint is **Connected*	- or **Migrated**.
//
// 	- **DisableVpcEndpointZoneConnection*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [ListVpcEndpointZones](https://help.aliyun.com/document_detail/183560.html) operation to query the status of the task.
//
//     	- If the zone is in the **Disconnecting*	- state, the task is running.
//
//     	- If the zone is in the **Disconnected*	- state, the task is successful.
//
// 	- You cannot repeatedly call the **DisableVpcEndpointZoneConnection*	- operation to allow an endpoint service to reject a connection request from the endpoint in the zone within a specified period of time.
//
// @param request - DisableVpcEndpointZoneConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableVpcEndpointZoneConnectionResponse
func (client *Client) DisableVpcEndpointZoneConnectionWithOptions(request *DisableVpcEndpointZoneConnectionRequest, runtime *util.RuntimeOptions) (_result *DisableVpcEndpointZoneConnectionResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.EndpointId)) {
		query["EndpointId"] = request.EndpointId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ReplacedResource)) {
		query["ReplacedResource"] = request.ReplacedResource
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableVpcEndpointZoneConnection"),
		Version:     tea.String("2020-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableVpcEndpointZoneConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disconnects an endpoint from a connection in the specified zone.
//
// Description:
//
//   You can call this operation only when the state of the endpoint is **Connected*	- and the state of the zone associated with the endpoint is **Connected*	- or **Migrated**.
//
// 	- **DisableVpcEndpointZoneConnection*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [ListVpcEndpointZones](https://help.aliyun.com/document_detail/183560.html) operation to query the status of the task.
//
//     	- If the zone is in the **Disconnecting*	- state, the task is running.
//
//     	- If the zone is in the **Disconnected*	- state, the task is successful.
//
// 	- You cannot repeatedly call the **DisableVpcEndpointZoneConnection*	- operation to allow an endpoint service to reject a connection request from the endpoint in the zone within a specified period of time.
//
// @param request - DisableVpcEndpointZoneConnectionRequest
//
// @return DisableVpcEndpointZoneConnectionResponse
func (client *Client) DisableVpcEndpointZoneConnection(request *DisableVpcEndpointZoneConnectionRequest) (_result *DisableVpcEndpointZoneConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableVpcEndpointZoneConnectionResponse{}
	_body, _err := client.DisableVpcEndpointZoneConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Accepts connection requests from an endpoint.
//
// Description:
//
//   **EnableVpcEndpointConnection*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [GetVpcEndpointAttribute](https://help.aliyun.com/document_detail/183568.html) operation to query the state of the endpoint connection.
//
//     	- If the state is **Connecting**, the endpoint connection is being established.
//
//     	- If the state is **Connected**, the endpoint connection is established.
//
// 	- You cannot repeatedly call the **EnableVpcEndpointConnection*	- operation to allow an endpoint service of an Alibaba Cloud account to accept a connection request from an endpoint within a specified period of time.
//
// @param request - EnableVpcEndpointConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableVpcEndpointConnectionResponse
func (client *Client) EnableVpcEndpointConnectionWithOptions(request *EnableVpcEndpointConnectionRequest, runtime *util.RuntimeOptions) (_result *EnableVpcEndpointConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bandwidth)) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointId)) {
		query["EndpointId"] = request.EndpointId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
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

// Summary:
//
// Accepts connection requests from an endpoint.
//
// Description:
//
//   **EnableVpcEndpointConnection*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [GetVpcEndpointAttribute](https://help.aliyun.com/document_detail/183568.html) operation to query the state of the endpoint connection.
//
//     	- If the state is **Connecting**, the endpoint connection is being established.
//
//     	- If the state is **Connected**, the endpoint connection is established.
//
// 	- You cannot repeatedly call the **EnableVpcEndpointConnection*	- operation to allow an endpoint service of an Alibaba Cloud account to accept a connection request from an endpoint within a specified period of time.
//
// @param request - EnableVpcEndpointConnectionRequest
//
// @return EnableVpcEndpointConnectionResponse
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

// Summary:
//
// Accepts a connection request from an endpoint in the specified zone.
//
// Description:
//
//   You can call this operation only when the state of the endpoint is **Connected*	- and the state of the associated zone is **Disconnected**.
//
// 	- **EnableVpcEndpointZoneConnection*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [ListVpcEndpointZones](https://help.aliyun.com/document_detail/183560.html) operation to check whether the endpoint service accepts a connection request from the endpoint in the associated zone.
//
//     	- If the zone is in the **Connecting*	- state, the endpoint service is accepting the connection request from the endpoint.
//
//     	- If the zone is in the **Connected*	- state, the endpoint service has accepted the connection request from the endpoint.
//
// 	- You cannot repeatedly call the **EnableVpcEndpointZoneConnection*	- operation to allow an endpoint service to accept a connection request from an endpoint in the associated zone within a specified period of time.
//
// @param request - EnableVpcEndpointZoneConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableVpcEndpointZoneConnectionResponse
func (client *Client) EnableVpcEndpointZoneConnectionWithOptions(request *EnableVpcEndpointZoneConnectionRequest, runtime *util.RuntimeOptions) (_result *EnableVpcEndpointZoneConnectionResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.EndpointId)) {
		query["EndpointId"] = request.EndpointId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableVpcEndpointZoneConnection"),
		Version:     tea.String("2020-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableVpcEndpointZoneConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Accepts a connection request from an endpoint in the specified zone.
//
// Description:
//
//   You can call this operation only when the state of the endpoint is **Connected*	- and the state of the associated zone is **Disconnected**.
//
// 	- **EnableVpcEndpointZoneConnection*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [ListVpcEndpointZones](https://help.aliyun.com/document_detail/183560.html) operation to check whether the endpoint service accepts a connection request from the endpoint in the associated zone.
//
//     	- If the zone is in the **Connecting*	- state, the endpoint service is accepting the connection request from the endpoint.
//
//     	- If the zone is in the **Connected*	- state, the endpoint service has accepted the connection request from the endpoint.
//
// 	- You cannot repeatedly call the **EnableVpcEndpointZoneConnection*	- operation to allow an endpoint service to accept a connection request from an endpoint in the associated zone within a specified period of time.
//
// @param request - EnableVpcEndpointZoneConnectionRequest
//
// @return EnableVpcEndpointZoneConnectionResponse
func (client *Client) EnableVpcEndpointZoneConnection(request *EnableVpcEndpointZoneConnectionRequest) (_result *EnableVpcEndpointZoneConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableVpcEndpointZoneConnectionResponse{}
	_body, _err := client.EnableVpcEndpointZoneConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the attributes of an endpoint.
//
// @param request - GetVpcEndpointAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVpcEndpointAttributeResponse
func (client *Client) GetVpcEndpointAttributeWithOptions(request *GetVpcEndpointAttributeRequest, runtime *util.RuntimeOptions) (_result *GetVpcEndpointAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndpointId)) {
		query["EndpointId"] = request.EndpointId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

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

// Summary:
//
// Queries the attributes of an endpoint.
//
// @param request - GetVpcEndpointAttributeRequest
//
// @return GetVpcEndpointAttributeResponse
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

// Summary:
//
// Queries the attributes of an endpoint service.
//
// @param request - GetVpcEndpointServiceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVpcEndpointServiceAttributeResponse
func (client *Client) GetVpcEndpointServiceAttributeWithOptions(request *GetVpcEndpointServiceAttributeRequest, runtime *util.RuntimeOptions) (_result *GetVpcEndpointServiceAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
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

// Summary:
//
// Queries the attributes of an endpoint service.
//
// @param request - GetVpcEndpointServiceAttributeRequest
//
// @return GetVpcEndpointServiceAttributeResponse
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

// Summary:
//
// Queries the tags that are added to resources.
//
// Description:
//
//   You must specify **ResourceId.N*	- or **Tag.N*	- in the request to specify the object that you want to query.
//
// 	- **Tag.N*	- is a resource tag that consists of a key-value pair (Tag.N.Key and Tag.N.Value). If you specify only **Tag.N.Key**, all tag values that are associated with the specified key are returned. If you specify only **Tag.N.Value**, an error message is returned.
//
// 	- If you specify **Tag.N*	- and **ResourceId.N*	- to filter tags, **ResourceId.N*	- must match all specified key-value pairs.
//
// 	- If you specify multiple key-value pairs, resources that contain these key-value pairs are returned.
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2020-04-15"),
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

// Summary:
//
// Queries the tags that are added to resources.
//
// Description:
//
//   You must specify **ResourceId.N*	- or **Tag.N*	- in the request to specify the object that you want to query.
//
// 	- **Tag.N*	- is a resource tag that consists of a key-value pair (Tag.N.Key and Tag.N.Value). If you specify only **Tag.N.Key**, all tag values that are associated with the specified key are returned. If you specify only **Tag.N.Value**, an error message is returned.
//
// 	- If you specify **Tag.N*	- and **ResourceId.N*	- to filter tags, **ResourceId.N*	- must match all specified key-value pairs.
//
// 	- If you specify multiple key-value pairs, resources that contain these key-value pairs are returned.
//
// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
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

// Summary:
//
// Queries endpoint connections.
//
// @param request - ListVpcEndpointConnectionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVpcEndpointConnectionsResponse
func (client *Client) ListVpcEndpointConnectionsWithOptions(request *ListVpcEndpointConnectionsRequest, runtime *util.RuntimeOptions) (_result *ListVpcEndpointConnectionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnectionStatus)) {
		query["ConnectionStatus"] = request.ConnectionStatus
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointId)) {
		query["EndpointId"] = request.EndpointId
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointOwnerId)) {
		query["EndpointOwnerId"] = request.EndpointOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.EniId)) {
		query["EniId"] = request.EniId
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ReplacedResourceId)) {
		query["ReplacedResourceId"] = request.ReplacedResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

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

// Summary:
//
// Queries endpoint connections.
//
// @param request - ListVpcEndpointConnectionsRequest
//
// @return ListVpcEndpointConnectionsResponse
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

// Summary:
//
// Queries the security groups that are associated with an endpoint.
//
// @param request - ListVpcEndpointSecurityGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVpcEndpointSecurityGroupsResponse
func (client *Client) ListVpcEndpointSecurityGroupsWithOptions(request *ListVpcEndpointSecurityGroupsRequest, runtime *util.RuntimeOptions) (_result *ListVpcEndpointSecurityGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndpointId)) {
		query["EndpointId"] = request.EndpointId
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

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

// Summary:
//
// Queries the security groups that are associated with an endpoint.
//
// @param request - ListVpcEndpointSecurityGroupsRequest
//
// @return ListVpcEndpointSecurityGroupsResponse
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

// Summary:
//
// Queries the service resources that are added to an endpoint service.
//
// @param request - ListVpcEndpointServiceResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVpcEndpointServiceResourcesResponse
func (client *Client) ListVpcEndpointServiceResourcesWithOptions(request *ListVpcEndpointServiceResourcesRequest, runtime *util.RuntimeOptions) (_result *ListVpcEndpointServiceResourcesResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
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

// Summary:
//
// Queries the service resources that are added to an endpoint service.
//
// @param request - ListVpcEndpointServiceResourcesRequest
//
// @return ListVpcEndpointServiceResourcesResponse
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

// Summary:
//
// Queries the whitelist of an endpoint service.
//
// @param request - ListVpcEndpointServiceUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVpcEndpointServiceUsersResponse
func (client *Client) ListVpcEndpointServiceUsersWithOptions(request *ListVpcEndpointServiceUsersRequest, runtime *util.RuntimeOptions) (_result *ListVpcEndpointServiceUsersResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserListType)) {
		query["UserListType"] = request.UserListType
	}

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

// Summary:
//
// Queries the whitelist of an endpoint service.
//
// @param request - ListVpcEndpointServiceUsersRequest
//
// @return ListVpcEndpointServiceUsersResponse
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

// Summary:
//
// Queries a list of endpoint services.
//
// @param request - ListVpcEndpointServicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVpcEndpointServicesResponse
func (client *Client) ListVpcEndpointServicesWithOptions(request *ListVpcEndpointServicesRequest, runtime *util.RuntimeOptions) (_result *ListVpcEndpointServicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoAcceptEnabled)) {
		query["AutoAcceptEnabled"] = request.AutoAcceptEnabled
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceBusinessStatus)) {
		query["ServiceBusinessStatus"] = request.ServiceBusinessStatus
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceResourceType)) {
		query["ServiceResourceType"] = request.ServiceResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceStatus)) {
		query["ServiceStatus"] = request.ServiceStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneAffinityEnabled)) {
		query["ZoneAffinityEnabled"] = request.ZoneAffinityEnabled
	}

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

// Summary:
//
// Queries a list of endpoint services.
//
// @param request - ListVpcEndpointServicesRequest
//
// @return ListVpcEndpointServicesResponse
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

// Summary:
//
// Queries a list of endpoint services that can be associated with the endpoint created within the current account.
//
// @param request - ListVpcEndpointServicesByEndUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVpcEndpointServicesByEndUserResponse
func (client *Client) ListVpcEndpointServicesByEndUserWithOptions(request *ListVpcEndpointServicesByEndUserRequest, runtime *util.RuntimeOptions) (_result *ListVpcEndpointServicesByEndUserResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceType)) {
		query["ServiceType"] = request.ServiceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

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

// Summary:
//
// Queries a list of endpoint services that can be associated with the endpoint created within the current account.
//
// @param request - ListVpcEndpointServicesByEndUserRequest
//
// @return ListVpcEndpointServicesByEndUserResponse
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

// Summary:
//
// Queries the zones of an endpoint.
//
// @param request - ListVpcEndpointZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVpcEndpointZonesResponse
func (client *Client) ListVpcEndpointZonesWithOptions(request *ListVpcEndpointZonesRequest, runtime *util.RuntimeOptions) (_result *ListVpcEndpointZonesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndpointId)) {
		query["EndpointId"] = request.EndpointId
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

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

// Summary:
//
// Queries the zones of an endpoint.
//
// @param request - ListVpcEndpointZonesRequest
//
// @return ListVpcEndpointZonesResponse
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

// Summary:
//
// Queries a list of endpoints.
//
// @param request - ListVpcEndpointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVpcEndpointsResponse
func (client *Client) ListVpcEndpointsWithOptions(request *ListVpcEndpointsRequest, runtime *util.RuntimeOptions) (_result *ListVpcEndpointsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnectionStatus)) {
		query["ConnectionStatus"] = request.ConnectionStatus
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointId)) {
		query["EndpointId"] = request.EndpointId
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointName)) {
		query["EndpointName"] = request.EndpointName
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointStatus)) {
		query["EndpointStatus"] = request.EndpointStatus
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointType)) {
		query["EndpointType"] = request.EndpointType
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

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

// Summary:
//
// Queries a list of endpoints.
//
// @param request - ListVpcEndpointsRequest
//
// @return ListVpcEndpointsResponse
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

// Summary:
//
// Activates PrivateLink.
//
// @param request - OpenPrivateLinkServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenPrivateLinkServiceResponse
func (client *Client) OpenPrivateLinkServiceWithOptions(request *OpenPrivateLinkServiceRequest, runtime *util.RuntimeOptions) (_result *OpenPrivateLinkServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

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

// Summary:
//
// Activates PrivateLink.
//
// @param request - OpenPrivateLinkServiceRequest
//
// @return OpenPrivateLinkServiceResponse
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

// Summary:
//
// Removes an account ID from the whitelist of an endpoint service.
//
// Description:
//
//   Before you remove an account ID from the whitelist of an endpoint service, make sure that the endpoint service is in the **Active*	- state. You can call the [GetVpcEndpointServiceAttribute](https://help.aliyun.com/document_detail/469330.html) operation to query the status of the endpoint service.
//
// 	- You cannot repeatedly call the **RemoveUserFromVpcEndpointService*	- operation to remove the ID of an Alibaba Cloud account from the whitelist of an endpoint service within a specified period of time.
//
// @param request - RemoveUserFromVpcEndpointServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveUserFromVpcEndpointServiceResponse
func (client *Client) RemoveUserFromVpcEndpointServiceWithOptions(request *RemoveUserFromVpcEndpointServiceRequest, runtime *util.RuntimeOptions) (_result *RemoveUserFromVpcEndpointServiceResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.UserARN)) {
		query["UserARN"] = request.UserARN
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

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

// Summary:
//
// Removes an account ID from the whitelist of an endpoint service.
//
// Description:
//
//   Before you remove an account ID from the whitelist of an endpoint service, make sure that the endpoint service is in the **Active*	- state. You can call the [GetVpcEndpointServiceAttribute](https://help.aliyun.com/document_detail/469330.html) operation to query the status of the endpoint service.
//
// 	- You cannot repeatedly call the **RemoveUserFromVpcEndpointService*	- operation to remove the ID of an Alibaba Cloud account from the whitelist of an endpoint service within a specified period of time.
//
// @param request - RemoveUserFromVpcEndpointServiceRequest
//
// @return RemoveUserFromVpcEndpointServiceResponse
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

// Summary:
//
// Deletes a zone of an endpoint.
//
// Description:
//
//   **RemoveZoneFromVpcEndpoint*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [ListVpcEndpointZones](https://help.aliyun.com/document_detail/183560.html) operation to check whether the zone of the endpoint is deleted.
//
//     	- If the zone of the endpoint is in the **Deleting*	- state, the zone is being deleted.
//
//     	- If the zone cannot be queried, the zone is deleted.
//
// 	- You cannot repeatedly call the **RemoveZoneFromVpcEndpoint*	- operation to delete a zone of an endpoint within a specified period of time.
//
// @param request - RemoveZoneFromVpcEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveZoneFromVpcEndpointResponse
func (client *Client) RemoveZoneFromVpcEndpointWithOptions(request *RemoveZoneFromVpcEndpointRequest, runtime *util.RuntimeOptions) (_result *RemoveZoneFromVpcEndpointResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.EndpointId)) {
		query["EndpointId"] = request.EndpointId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

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

// Summary:
//
// Deletes a zone of an endpoint.
//
// Description:
//
//   **RemoveZoneFromVpcEndpoint*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [ListVpcEndpointZones](https://help.aliyun.com/document_detail/183560.html) operation to check whether the zone of the endpoint is deleted.
//
//     	- If the zone of the endpoint is in the **Deleting*	- state, the zone is being deleted.
//
//     	- If the zone cannot be queried, the zone is deleted.
//
// 	- You cannot repeatedly call the **RemoveZoneFromVpcEndpoint*	- operation to delete a zone of an endpoint within a specified period of time.
//
// @param request - RemoveZoneFromVpcEndpointRequest
//
// @return RemoveZoneFromVpcEndpointResponse
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

// Summary:
//
// Adds tags to resources. You can call this API operation to add tags to one or more endpoints or endpoint services.
//
// Description:
//
// > You can add up to 20 tags to an instance. Before you add tags to a resource, Alibaba Cloud checks the number of existing tags of the resource. If the maximum number is reached, an error message is returned.
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		bodyFlat["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		bodyFlat["Tag"] = request.Tag
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2020-04-15"),
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

// Summary:
//
// Adds tags to resources. You can call this API operation to add tags to one or more endpoints or endpoint services.
//
// Description:
//
// > You can add up to 20 tags to an instance. Before you add tags to a resource, Alibaba Cloud checks the number of existing tags of the resource. If the maximum number is reached, an error message is returned.
//
// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
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

// Summary:
//
// Removes tags from resources. You can call the UntagResources operation to remove tags from one or more endpoints or endpoint services.
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		body["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		bodyFlat["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		bodyFlat["TagKey"] = request.TagKey
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2020-04-15"),
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

// Summary:
//
// Removes tags from resources. You can call the UntagResources operation to remove tags from one or more endpoints or endpoint services.
//
// @param request - UntagResourcesRequest
//
// @return UntagResourcesResponse
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

// Summary:
//
// Modifies the attributes of an endpoint.
//
// Description:
//
// You cannot repeatedly call the **UpdateVpcEndpointAttribute*	- operation to modify the attributes of an endpoint that belongs to an Alibaba Cloud account within a specified period of time.
//
// @param request - UpdateVpcEndpointAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVpcEndpointAttributeResponse
func (client *Client) UpdateVpcEndpointAttributeWithOptions(request *UpdateVpcEndpointAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateVpcEndpointAttributeResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.EndpointDescription)) {
		query["EndpointDescription"] = request.EndpointDescription
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointId)) {
		query["EndpointId"] = request.EndpointId
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointName)) {
		query["EndpointName"] = request.EndpointName
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyDocument)) {
		query["PolicyDocument"] = request.PolicyDocument
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

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

// Summary:
//
// Modifies the attributes of an endpoint.
//
// Description:
//
// You cannot repeatedly call the **UpdateVpcEndpointAttribute*	- operation to modify the attributes of an endpoint that belongs to an Alibaba Cloud account within a specified period of time.
//
// @param request - UpdateVpcEndpointAttributeRequest
//
// @return UpdateVpcEndpointAttributeResponse
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

// Summary:
//
// Modifies the attributes of an endpoint connection.
//
// Description:
//
// You cannot repeatedly call the **UpdateVpcEndpointConnectionAttribute*	- operation to modify the bandwidth of an endpoint connection that belongs to an Alibaba Cloud account within a specified period of time.
//
// @param request - UpdateVpcEndpointConnectionAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVpcEndpointConnectionAttributeResponse
func (client *Client) UpdateVpcEndpointConnectionAttributeWithOptions(request *UpdateVpcEndpointConnectionAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateVpcEndpointConnectionAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bandwidth)) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointId)) {
		query["EndpointId"] = request.EndpointId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
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

// Summary:
//
// Modifies the attributes of an endpoint connection.
//
// Description:
//
// You cannot repeatedly call the **UpdateVpcEndpointConnectionAttribute*	- operation to modify the bandwidth of an endpoint connection that belongs to an Alibaba Cloud account within a specified period of time.
//
// @param request - UpdateVpcEndpointConnectionAttributeRequest
//
// @return UpdateVpcEndpointConnectionAttributeResponse
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

// Summary:
//
// Modifies the attributes of an endpoint service.
//
// Description:
//
// You cannot repeatedly call the **UpdateVpcEndpointServiceAttribute*	- operation to modify the attributes of an endpoint service that belongs to an Alibaba Cloud account within a specified period of time.
//
// @param request - UpdateVpcEndpointServiceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVpcEndpointServiceAttributeResponse
func (client *Client) UpdateVpcEndpointServiceAttributeWithOptions(request *UpdateVpcEndpointServiceAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateVpcEndpointServiceAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoAcceptEnabled)) {
		query["AutoAcceptEnabled"] = request.AutoAcceptEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectBandwidth)) {
		query["ConnectBandwidth"] = request.ConnectBandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
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

	if !tea.BoolValue(util.IsUnset(request.ServiceSupportIPv6)) {
		query["ServiceSupportIPv6"] = request.ServiceSupportIPv6
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneAffinityEnabled)) {
		query["ZoneAffinityEnabled"] = request.ZoneAffinityEnabled
	}

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

// Summary:
//
// Modifies the attributes of an endpoint service.
//
// Description:
//
// You cannot repeatedly call the **UpdateVpcEndpointServiceAttribute*	- operation to modify the attributes of an endpoint service that belongs to an Alibaba Cloud account within a specified period of time.
//
// @param request - UpdateVpcEndpointServiceAttributeRequest
//
// @return UpdateVpcEndpointServiceAttributeResponse
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

// Summary:
//
// Modifies the attributes of a service resource that is added to an endpoint service.
//
// Description:
//
// You cannot repeatedly call the **UpdateVpcEndpointServiceResourceAttribute*	- operation to modify the attributes of a service resource that is added to an endpoint service within a specified period of time.
//
// @param request - UpdateVpcEndpointServiceResourceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVpcEndpointServiceResourceAttributeResponse
func (client *Client) UpdateVpcEndpointServiceResourceAttributeWithOptions(request *UpdateVpcEndpointServiceResourceAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateVpcEndpointServiceResourceAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoAllocatedEnabled)) {
		query["AutoAllocatedEnabled"] = request.AutoAllocatedEnabled
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateVpcEndpointServiceResourceAttribute"),
		Version:     tea.String("2020-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateVpcEndpointServiceResourceAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the attributes of a service resource that is added to an endpoint service.
//
// Description:
//
// You cannot repeatedly call the **UpdateVpcEndpointServiceResourceAttribute*	- operation to modify the attributes of a service resource that is added to an endpoint service within a specified period of time.
//
// @param request - UpdateVpcEndpointServiceResourceAttributeRequest
//
// @return UpdateVpcEndpointServiceResourceAttributeResponse
func (client *Client) UpdateVpcEndpointServiceResourceAttribute(request *UpdateVpcEndpointServiceResourceAttributeRequest) (_result *UpdateVpcEndpointServiceResourceAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateVpcEndpointServiceResourceAttributeResponse{}
	_body, _err := client.UpdateVpcEndpointServiceResourceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a service resource of a zone to which an endpoint connection belongs.
//
// Description:
//
// ### Prerequisites
//
// By default, the feature of modifying a service resource of a zone to which an endpoint connection belongs is unavailable. To use this feature, log on to the [Quota Center console](https://quotas.console.aliyun.com/white-list-products/privatelink/quotas). Click Whitelist Quotas in the left-side navigation pane and click PrivateLink in the Networking section. On the page that appears, search for `privatelink_whitelist/svc_res_mgt_uat` and click Apply in the Actions column.
//
// ### Usage notes
//
// 	- If the endpoint connection is in the **Disconnected*	- state, you can manually allocate the service resource in the zone.
//
// 	- If the endpoint connection is in the **Connected*	- state, you can manually migrate the service resource in the zone. In this case, the connection might be interrupted.
//
// 	- **UpdateVpcEndpointZoneConnectionResourceAttribute*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [GetVpcEndpointServiceAttribute](https://help.aliyun.com/document_detail/469330.html) operation to check whether the service resource is modified.
//
//     	- If the endpoint service is in the **Pending*	- state, the service resource is being modified.
//
//     	- If the endpoint service is in the **Active*	- state, the service resource is modified.
//
// 	- You cannot repeatedly call the **UpdateVpcEndpointZoneConnectionResourceAttribute*	- operation to modify a service resource in the zone to which an endpoint connection belongs within a specified period of time.
//
// @param request - UpdateVpcEndpointZoneConnectionResourceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVpcEndpointZoneConnectionResourceAttributeResponse
func (client *Client) UpdateVpcEndpointZoneConnectionResourceAttributeWithOptions(request *UpdateVpcEndpointZoneConnectionResourceAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateVpcEndpointZoneConnectionResourceAttributeResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.EndpointId)) {
		query["EndpointId"] = request.EndpointId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceAllocateMode)) {
		query["ResourceAllocateMode"] = request.ResourceAllocateMode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceReplaceMode)) {
		query["ResourceReplaceMode"] = request.ResourceReplaceMode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateVpcEndpointZoneConnectionResourceAttribute"),
		Version:     tea.String("2020-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateVpcEndpointZoneConnectionResourceAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a service resource of a zone to which an endpoint connection belongs.
//
// Description:
//
// ### Prerequisites
//
// By default, the feature of modifying a service resource of a zone to which an endpoint connection belongs is unavailable. To use this feature, log on to the [Quota Center console](https://quotas.console.aliyun.com/white-list-products/privatelink/quotas). Click Whitelist Quotas in the left-side navigation pane and click PrivateLink in the Networking section. On the page that appears, search for `privatelink_whitelist/svc_res_mgt_uat` and click Apply in the Actions column.
//
// ### Usage notes
//
// 	- If the endpoint connection is in the **Disconnected*	- state, you can manually allocate the service resource in the zone.
//
// 	- If the endpoint connection is in the **Connected*	- state, you can manually migrate the service resource in the zone. In this case, the connection might be interrupted.
//
// 	- **UpdateVpcEndpointZoneConnectionResourceAttribute*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [GetVpcEndpointServiceAttribute](https://help.aliyun.com/document_detail/469330.html) operation to check whether the service resource is modified.
//
//     	- If the endpoint service is in the **Pending*	- state, the service resource is being modified.
//
//     	- If the endpoint service is in the **Active*	- state, the service resource is modified.
//
// 	- You cannot repeatedly call the **UpdateVpcEndpointZoneConnectionResourceAttribute*	- operation to modify a service resource in the zone to which an endpoint connection belongs within a specified period of time.
//
// @param request - UpdateVpcEndpointZoneConnectionResourceAttributeRequest
//
// @return UpdateVpcEndpointZoneConnectionResourceAttributeResponse
func (client *Client) UpdateVpcEndpointZoneConnectionResourceAttribute(request *UpdateVpcEndpointZoneConnectionResourceAttributeRequest) (_result *UpdateVpcEndpointZoneConnectionResourceAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateVpcEndpointZoneConnectionResourceAttributeResponse{}
	_body, _err := client.UpdateVpcEndpointZoneConnectionResourceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
