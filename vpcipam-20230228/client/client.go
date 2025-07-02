// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddIpamPoolCidrRequest struct {
	// The CIDR block that you want to provision.
	//
	// >  Only IPv4 CIDR blocks are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.1.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the IPAM pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-pool-6rcq3tobayc20t****
	IpamPoolId *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
	// The ID of the region where the IPAM instance is hosted.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddIpamPoolCidrRequest) String() string {
	return tea.Prettify(s)
}

func (s AddIpamPoolCidrRequest) GoString() string {
	return s.String()
}

func (s *AddIpamPoolCidrRequest) SetCidr(v string) *AddIpamPoolCidrRequest {
	s.Cidr = &v
	return s
}

func (s *AddIpamPoolCidrRequest) SetClientToken(v string) *AddIpamPoolCidrRequest {
	s.ClientToken = &v
	return s
}

func (s *AddIpamPoolCidrRequest) SetDryRun(v bool) *AddIpamPoolCidrRequest {
	s.DryRun = &v
	return s
}

func (s *AddIpamPoolCidrRequest) SetIpamPoolId(v string) *AddIpamPoolCidrRequest {
	s.IpamPoolId = &v
	return s
}

func (s *AddIpamPoolCidrRequest) SetRegionId(v string) *AddIpamPoolCidrRequest {
	s.RegionId = &v
	return s
}

type AddIpamPoolCidrResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 558BC336-8B88-53B0-B4AD-980EE900AB01
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddIpamPoolCidrResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddIpamPoolCidrResponseBody) GoString() string {
	return s.String()
}

func (s *AddIpamPoolCidrResponseBody) SetRequestId(v string) *AddIpamPoolCidrResponseBody {
	s.RequestId = &v
	return s
}

type AddIpamPoolCidrResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddIpamPoolCidrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddIpamPoolCidrResponse) String() string {
	return tea.Prettify(s)
}

func (s AddIpamPoolCidrResponse) GoString() string {
	return s.String()
}

func (s *AddIpamPoolCidrResponse) SetHeaders(v map[string]*string) *AddIpamPoolCidrResponse {
	s.Headers = v
	return s
}

func (s *AddIpamPoolCidrResponse) SetStatusCode(v int32) *AddIpamPoolCidrResponse {
	s.StatusCode = &v
	return s
}

func (s *AddIpamPoolCidrResponse) SetBody(v *AddIpamPoolCidrResponseBody) *AddIpamPoolCidrResponse {
	s.Body = v
	return s
}

type AssociateIpamResourceDiscoveryRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform the dry run. Valid values:
	//
	// 	- **true**: Performs a dry run without associating the resource discovery and IPAM instance. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and sends the request. After the request passes the check, an HTTP 2xx status code is returned and the resource discovery and IPAM instances are associated.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the IPAM.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-ccxbnsbhew0d6t****
	IpamId *string `json:"IpamId,omitempty" xml:"IpamId,omitempty"`
	// The ID of resource discovery instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-res-disco-jt5f2af2u6nk2z321****
	IpamResourceDiscoveryId *string `json:"IpamResourceDiscoveryId,omitempty" xml:"IpamResourceDiscoveryId,omitempty"`
	OwnerAccount            *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                 *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The request region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AssociateIpamResourceDiscoveryRequest) String() string {
	return tea.Prettify(s)
}

func (s AssociateIpamResourceDiscoveryRequest) GoString() string {
	return s.String()
}

func (s *AssociateIpamResourceDiscoveryRequest) SetClientToken(v string) *AssociateIpamResourceDiscoveryRequest {
	s.ClientToken = &v
	return s
}

func (s *AssociateIpamResourceDiscoveryRequest) SetDryRun(v bool) *AssociateIpamResourceDiscoveryRequest {
	s.DryRun = &v
	return s
}

func (s *AssociateIpamResourceDiscoveryRequest) SetIpamId(v string) *AssociateIpamResourceDiscoveryRequest {
	s.IpamId = &v
	return s
}

func (s *AssociateIpamResourceDiscoveryRequest) SetIpamResourceDiscoveryId(v string) *AssociateIpamResourceDiscoveryRequest {
	s.IpamResourceDiscoveryId = &v
	return s
}

func (s *AssociateIpamResourceDiscoveryRequest) SetOwnerAccount(v string) *AssociateIpamResourceDiscoveryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AssociateIpamResourceDiscoveryRequest) SetOwnerId(v int64) *AssociateIpamResourceDiscoveryRequest {
	s.OwnerId = &v
	return s
}

func (s *AssociateIpamResourceDiscoveryRequest) SetRegionId(v string) *AssociateIpamResourceDiscoveryRequest {
	s.RegionId = &v
	return s
}

func (s *AssociateIpamResourceDiscoveryRequest) SetResourceOwnerAccount(v string) *AssociateIpamResourceDiscoveryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AssociateIpamResourceDiscoveryRequest) SetResourceOwnerId(v int64) *AssociateIpamResourceDiscoveryRequest {
	s.ResourceOwnerId = &v
	return s
}

type AssociateIpamResourceDiscoveryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// E897D16A-50EB-543F-B002-C5A26AB818FF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateIpamResourceDiscoveryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssociateIpamResourceDiscoveryResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateIpamResourceDiscoveryResponseBody) SetRequestId(v string) *AssociateIpamResourceDiscoveryResponseBody {
	s.RequestId = &v
	return s
}

type AssociateIpamResourceDiscoveryResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateIpamResourceDiscoveryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateIpamResourceDiscoveryResponse) String() string {
	return tea.Prettify(s)
}

func (s AssociateIpamResourceDiscoveryResponse) GoString() string {
	return s.String()
}

func (s *AssociateIpamResourceDiscoveryResponse) SetHeaders(v map[string]*string) *AssociateIpamResourceDiscoveryResponse {
	s.Headers = v
	return s
}

func (s *AssociateIpamResourceDiscoveryResponse) SetStatusCode(v int32) *AssociateIpamResourceDiscoveryResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateIpamResourceDiscoveryResponse) SetBody(v *AssociateIpamResourceDiscoveryResponseBody) *AssociateIpamResourceDiscoveryResponse {
	s.Body = v
	return s
}

type ChangeResourceGroupRequest struct {
	// The ID of the new resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-aek3ctkufaw****
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the IPAM instance is hosted. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the IPAM resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-uq5dcfc2eqhpf4****
	ResourceId           *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Resource type, with values:
	//
	// - Ipam:IPAM instance
	//
	// - IpamScope:IPAM scope
	//
	// - IpamPool:IPAM address pool
	//
	// - IpamResourceDiscovery:IPAM resource discovery
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ChangeResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupRequest) SetNewResourceGroupId(v string) *ChangeResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetOwnerAccount(v string) *ChangeResourceGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetOwnerId(v int64) *ChangeResourceGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetRegionId(v string) *ChangeResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceId(v string) *ChangeResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceOwnerAccount(v string) *ChangeResourceGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceOwnerId(v int64) *ChangeResourceGroupRequest {
	s.ResourceOwnerId = &v
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
	// BB2C39DE-CEB8-595A-981A-F2EFCBE7324E
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

type CreateIpamRequest struct {
	// The client token used to ensure the idempotence of the request. Use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The description of the IPAM.
	//
	// It must be 1 to 256 characters in length. Start with a letter but cannot start with `http://` or `https://`. If you do not specify a description, the description is empty by default.
	//
	// example:
	//
	// This is my first Ipam
	IpamDescription *string `json:"IpamDescription,omitempty" xml:"IpamDescription,omitempty"`
	// The name of the IPAM.
	//
	// It must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// abc
	IpamName *string `json:"IpamName,omitempty" xml:"IpamName,omitempty"`
	// The effective regions of the IPAM.
	//
	// This parameter is required.
	OperatingRegionList []*string `json:"OperatingRegionList,omitempty" xml:"OperatingRegionList,omitempty" type:"Repeated"`
	OwnerAccount        *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId             *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the IPAM instance is hosted. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID of the IPAM.
	//
	// example:
	//
	// rg-acfmxazb4ph6aiy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tag list.
	Tag []*CreateIpamRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateIpamRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIpamRequest) GoString() string {
	return s.String()
}

func (s *CreateIpamRequest) SetClientToken(v string) *CreateIpamRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateIpamRequest) SetDryRun(v bool) *CreateIpamRequest {
	s.DryRun = &v
	return s
}

func (s *CreateIpamRequest) SetIpamDescription(v string) *CreateIpamRequest {
	s.IpamDescription = &v
	return s
}

func (s *CreateIpamRequest) SetIpamName(v string) *CreateIpamRequest {
	s.IpamName = &v
	return s
}

func (s *CreateIpamRequest) SetOperatingRegionList(v []*string) *CreateIpamRequest {
	s.OperatingRegionList = v
	return s
}

func (s *CreateIpamRequest) SetOwnerAccount(v string) *CreateIpamRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateIpamRequest) SetOwnerId(v int64) *CreateIpamRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateIpamRequest) SetRegionId(v string) *CreateIpamRequest {
	s.RegionId = &v
	return s
}

func (s *CreateIpamRequest) SetResourceGroupId(v string) *CreateIpamRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateIpamRequest) SetResourceOwnerAccount(v string) *CreateIpamRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateIpamRequest) SetResourceOwnerId(v int64) *CreateIpamRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateIpamRequest) SetTag(v []*CreateIpamRequestTag) *CreateIpamRequest {
	s.Tag = v
	return s
}

type CreateIpamRequestTag struct {
	// The tag key of the resource. You can specify at most 20 tag keys. It cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter but cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. You can specify up to 20 tag values. You can specify empty strings as tag values.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateIpamRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateIpamRequestTag) GoString() string {
	return s.String()
}

func (s *CreateIpamRequestTag) SetKey(v string) *CreateIpamRequestTag {
	s.Key = &v
	return s
}

func (s *CreateIpamRequestTag) SetValue(v string) *CreateIpamRequestTag {
	s.Value = &v
	return s
}

type CreateIpamResponseBody struct {
	// The ID of the default resource discovery association.
	//
	// example:
	//
	// ipam-res-disco-assoc-jt5fac8twugdbbgip****
	DefaultResourceDiscoveryAssociationId *string `json:"DefaultResourceDiscoveryAssociationId,omitempty" xml:"DefaultResourceDiscoveryAssociationId,omitempty"`
	// The ID of the default resource discovery instance.
	//
	// example:
	//
	// ipam-res-disco-jt5f2af2u6nk2z321****
	DefaultResourceDiscoveryId *string `json:"DefaultResourceDiscoveryId,omitempty" xml:"DefaultResourceDiscoveryId,omitempty"`
	// The ID of the IPAM.
	//
	// example:
	//
	// ipam-ccxbnsbhew0d6t****
	IpamId *string `json:"IpamId,omitempty" xml:"IpamId,omitempty"`
	// The default private scope created by the system after the IPAM is created.
	//
	// example:
	//
	// ipam-scope-8wiontzmiy6cg0****
	PrivateDefaultScopeId *string `json:"PrivateDefaultScopeId,omitempty" xml:"PrivateDefaultScopeId,omitempty"`
	// The default public scope created by the system after the IPAM is created.
	//
	// example:
	//
	// ipam-scope-r5c5c7bmym1brc****
	PublicDefaultScopeId *string `json:"PublicDefaultScopeId,omitempty" xml:"PublicDefaultScopeId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED39DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of discovered resources.
	//
	// example:
	//
	// 1
	ResourceDiscoveryAssociationCount *int32 `json:"ResourceDiscoveryAssociationCount,omitempty" xml:"ResourceDiscoveryAssociationCount,omitempty"`
}

func (s CreateIpamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIpamResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIpamResponseBody) SetDefaultResourceDiscoveryAssociationId(v string) *CreateIpamResponseBody {
	s.DefaultResourceDiscoveryAssociationId = &v
	return s
}

func (s *CreateIpamResponseBody) SetDefaultResourceDiscoveryId(v string) *CreateIpamResponseBody {
	s.DefaultResourceDiscoveryId = &v
	return s
}

func (s *CreateIpamResponseBody) SetIpamId(v string) *CreateIpamResponseBody {
	s.IpamId = &v
	return s
}

func (s *CreateIpamResponseBody) SetPrivateDefaultScopeId(v string) *CreateIpamResponseBody {
	s.PrivateDefaultScopeId = &v
	return s
}

func (s *CreateIpamResponseBody) SetPublicDefaultScopeId(v string) *CreateIpamResponseBody {
	s.PublicDefaultScopeId = &v
	return s
}

func (s *CreateIpamResponseBody) SetRequestId(v string) *CreateIpamResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIpamResponseBody) SetResourceDiscoveryAssociationCount(v int32) *CreateIpamResponseBody {
	s.ResourceDiscoveryAssociationCount = &v
	return s
}

type CreateIpamResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIpamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIpamResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIpamResponse) GoString() string {
	return s.String()
}

func (s *CreateIpamResponse) SetHeaders(v map[string]*string) *CreateIpamResponse {
	s.Headers = v
	return s
}

func (s *CreateIpamResponse) SetStatusCode(v int32) *CreateIpamResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIpamResponse) SetBody(v *CreateIpamResponseBody) *CreateIpamResponse {
	s.Body = v
	return s
}

type CreateIpamPoolRequest struct {
	// The default network mask assigned to the IPAM pool.
	//
	// An IPv4 mask must be **0 to 32*	- bits in length.
	//
	// example:
	//
	// 28
	AllocationDefaultCidrMask *int32 `json:"AllocationDefaultCidrMask,omitempty" xml:"AllocationDefaultCidrMask,omitempty"`
	// The maximum network mask assigned to the IPAM pool.
	//
	// An IPv4 mask must be **0 to 32*	- bits in length.
	//
	// example:
	//
	// 32
	AllocationMaxCidrMask *int32 `json:"AllocationMaxCidrMask,omitempty" xml:"AllocationMaxCidrMask,omitempty"`
	// The minimum network mask assigned to the IPAM pool.
	//
	// An IPv4 mask must be **0 to 32*	- bits in length.
	//
	// example:
	//
	// 8
	AllocationMinCidrMask *int32 `json:"AllocationMinCidrMask,omitempty" xml:"AllocationMinCidrMask,omitempty"`
	// Whether the pool has the auto-import feature enabled.
	//
	// example:
	//
	// true
	AutoImport *bool `json:"AutoImport,omitempty" xml:"AutoImport,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The IP version. Only **IPv4*	- is supported.
	//
	// example:
	//
	// IPv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The description of the IPAM pool.
	//
	// It must be 2 to 256 characters in length. It must start with a letter, but cannot start with a `http://` or `https://`. This parameter is empty by default.
	//
	// example:
	//
	// test description
	IpamPoolDescription *string `json:"IpamPoolDescription,omitempty" xml:"IpamPoolDescription,omitempty"`
	// The name of the IPAM pool.
	//
	// It must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// abc
	IpamPoolName *string `json:"IpamPoolName,omitempty" xml:"IpamPoolName,omitempty"`
	// The ID of the IPAM scope.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-scope-glfmcyldpm8lsy****
	IpamScopeId  *string `json:"IpamScopeId,omitempty" xml:"IpamScopeId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The effective region of the IPAM pool.
	//
	// example:
	//
	// cn-hangzhou
	PoolRegionId *string `json:"PoolRegionId,omitempty" xml:"PoolRegionId,omitempty"`
	// The ID of the region where the IPAM instance is hosted. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxazb4ph6aiy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the source IPAM pool.
	//
	// >  If you do not specify this parameter, the pool is a parent pool.
	//
	// example:
	//
	// ipam-pool-6rcq3tobayc20t****
	SourceIpamPoolId *string `json:"SourceIpamPoolId,omitempty" xml:"SourceIpamPoolId,omitempty"`
	// The tag list.
	Tag []*CreateIpamPoolRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateIpamPoolRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIpamPoolRequest) GoString() string {
	return s.String()
}

func (s *CreateIpamPoolRequest) SetAllocationDefaultCidrMask(v int32) *CreateIpamPoolRequest {
	s.AllocationDefaultCidrMask = &v
	return s
}

func (s *CreateIpamPoolRequest) SetAllocationMaxCidrMask(v int32) *CreateIpamPoolRequest {
	s.AllocationMaxCidrMask = &v
	return s
}

func (s *CreateIpamPoolRequest) SetAllocationMinCidrMask(v int32) *CreateIpamPoolRequest {
	s.AllocationMinCidrMask = &v
	return s
}

func (s *CreateIpamPoolRequest) SetAutoImport(v bool) *CreateIpamPoolRequest {
	s.AutoImport = &v
	return s
}

func (s *CreateIpamPoolRequest) SetClientToken(v string) *CreateIpamPoolRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateIpamPoolRequest) SetDryRun(v bool) *CreateIpamPoolRequest {
	s.DryRun = &v
	return s
}

func (s *CreateIpamPoolRequest) SetIpVersion(v string) *CreateIpamPoolRequest {
	s.IpVersion = &v
	return s
}

func (s *CreateIpamPoolRequest) SetIpamPoolDescription(v string) *CreateIpamPoolRequest {
	s.IpamPoolDescription = &v
	return s
}

func (s *CreateIpamPoolRequest) SetIpamPoolName(v string) *CreateIpamPoolRequest {
	s.IpamPoolName = &v
	return s
}

func (s *CreateIpamPoolRequest) SetIpamScopeId(v string) *CreateIpamPoolRequest {
	s.IpamScopeId = &v
	return s
}

func (s *CreateIpamPoolRequest) SetOwnerAccount(v string) *CreateIpamPoolRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateIpamPoolRequest) SetOwnerId(v int64) *CreateIpamPoolRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateIpamPoolRequest) SetPoolRegionId(v string) *CreateIpamPoolRequest {
	s.PoolRegionId = &v
	return s
}

func (s *CreateIpamPoolRequest) SetRegionId(v string) *CreateIpamPoolRequest {
	s.RegionId = &v
	return s
}

func (s *CreateIpamPoolRequest) SetResourceGroupId(v string) *CreateIpamPoolRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateIpamPoolRequest) SetResourceOwnerAccount(v string) *CreateIpamPoolRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateIpamPoolRequest) SetResourceOwnerId(v int64) *CreateIpamPoolRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateIpamPoolRequest) SetSourceIpamPoolId(v string) *CreateIpamPoolRequest {
	s.SourceIpamPoolId = &v
	return s
}

func (s *CreateIpamPoolRequest) SetTag(v []*CreateIpamPoolRequestTag) *CreateIpamPoolRequest {
	s.Tag = v
	return s
}

type CreateIpamPoolRequestTag struct {
	// The tag key. You can specify at most 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The tag key must start with a letter but cannot start with `aliyun` or `acs:`. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It cannot start with a `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateIpamPoolRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateIpamPoolRequestTag) GoString() string {
	return s.String()
}

func (s *CreateIpamPoolRequestTag) SetKey(v string) *CreateIpamPoolRequestTag {
	s.Key = &v
	return s
}

func (s *CreateIpamPoolRequestTag) SetValue(v string) *CreateIpamPoolRequestTag {
	s.Value = &v
	return s
}

type CreateIpamPoolResponseBody struct {
	// The ID of the IPAM pool.
	//
	// example:
	//
	// ipam-pool-6rcq3tobayc20t****
	IpamPoolId *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BB2C39DE-CEB8-595A-981A-F2EFCBE7324E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateIpamPoolResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIpamPoolResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIpamPoolResponseBody) SetIpamPoolId(v string) *CreateIpamPoolResponseBody {
	s.IpamPoolId = &v
	return s
}

func (s *CreateIpamPoolResponseBody) SetRequestId(v string) *CreateIpamPoolResponseBody {
	s.RequestId = &v
	return s
}

type CreateIpamPoolResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIpamPoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIpamPoolResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIpamPoolResponse) GoString() string {
	return s.String()
}

func (s *CreateIpamPoolResponse) SetHeaders(v map[string]*string) *CreateIpamPoolResponse {
	s.Headers = v
	return s
}

func (s *CreateIpamPoolResponse) SetStatusCode(v int32) *CreateIpamPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIpamPoolResponse) SetBody(v *CreateIpamPoolResponseBody) *CreateIpamPoolResponse {
	s.Body = v
	return s
}

type CreateIpamPoolAllocationRequest struct {
	// Enter a CIDR block to reserve a custom CIDR block.
	//
	// **Usage notes*	- Specify at least one of **Cidr*	- and **CidrMask*	- .
	//
	// example:
	//
	// 192.168.1.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// Enter a mask to reserve a custom CIDR block.
	//
	// **Usage notes*	- Specify at least one of **Cidr*	- and **CidrMask*	- .
	//
	// example:
	//
	// 24
	CidrMask *int32 `json:"CidrMask,omitempty" xml:"CidrMask,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// **Usage notes*	- If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The description of the allocation.
	//
	// example:
	//
	// test description
	IpamPoolAllocationDescription *string `json:"IpamPoolAllocationDescription,omitempty" xml:"IpamPoolAllocationDescription,omitempty"`
	// The name of the allocation.
	//
	// example:
	//
	// test name
	IpamPoolAllocationName *string `json:"IpamPoolAllocationName,omitempty" xml:"IpamPoolAllocationName,omitempty"`
	// The ID of the IPAM pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-pool-6rcq3tobayc20t****
	IpamPoolId *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
	// The region ID of the custom CIDR block that you want to reserve.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateIpamPoolAllocationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIpamPoolAllocationRequest) GoString() string {
	return s.String()
}

func (s *CreateIpamPoolAllocationRequest) SetCidr(v string) *CreateIpamPoolAllocationRequest {
	s.Cidr = &v
	return s
}

func (s *CreateIpamPoolAllocationRequest) SetCidrMask(v int32) *CreateIpamPoolAllocationRequest {
	s.CidrMask = &v
	return s
}

func (s *CreateIpamPoolAllocationRequest) SetClientToken(v string) *CreateIpamPoolAllocationRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateIpamPoolAllocationRequest) SetDryRun(v bool) *CreateIpamPoolAllocationRequest {
	s.DryRun = &v
	return s
}

func (s *CreateIpamPoolAllocationRequest) SetIpamPoolAllocationDescription(v string) *CreateIpamPoolAllocationRequest {
	s.IpamPoolAllocationDescription = &v
	return s
}

func (s *CreateIpamPoolAllocationRequest) SetIpamPoolAllocationName(v string) *CreateIpamPoolAllocationRequest {
	s.IpamPoolAllocationName = &v
	return s
}

func (s *CreateIpamPoolAllocationRequest) SetIpamPoolId(v string) *CreateIpamPoolAllocationRequest {
	s.IpamPoolId = &v
	return s
}

func (s *CreateIpamPoolAllocationRequest) SetRegionId(v string) *CreateIpamPoolAllocationRequest {
	s.RegionId = &v
	return s
}

type CreateIpamPoolAllocationResponseBody struct {
	// The custom reserved CIDR block.
	//
	// example:
	//
	// 192.168.1.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The ID of the custom reserved CIDR block.
	//
	// example:
	//
	// ipam-pool-alloc-112za33e4****
	IpamPoolAllocationId *string `json:"IpamPoolAllocationId,omitempty" xml:"IpamPoolAllocationId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CE9CDAE5-341E-5D0B-AC8A-2BAC707D3EB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The source CIDR block.
	//
	// example:
	//
	// 192.168.0.0/16
	SourceCidr *string `json:"SourceCidr,omitempty" xml:"SourceCidr,omitempty"`
}

func (s CreateIpamPoolAllocationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIpamPoolAllocationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIpamPoolAllocationResponseBody) SetCidr(v string) *CreateIpamPoolAllocationResponseBody {
	s.Cidr = &v
	return s
}

func (s *CreateIpamPoolAllocationResponseBody) SetIpamPoolAllocationId(v string) *CreateIpamPoolAllocationResponseBody {
	s.IpamPoolAllocationId = &v
	return s
}

func (s *CreateIpamPoolAllocationResponseBody) SetRequestId(v string) *CreateIpamPoolAllocationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIpamPoolAllocationResponseBody) SetSourceCidr(v string) *CreateIpamPoolAllocationResponseBody {
	s.SourceCidr = &v
	return s
}

type CreateIpamPoolAllocationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIpamPoolAllocationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIpamPoolAllocationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIpamPoolAllocationResponse) GoString() string {
	return s.String()
}

func (s *CreateIpamPoolAllocationResponse) SetHeaders(v map[string]*string) *CreateIpamPoolAllocationResponse {
	s.Headers = v
	return s
}

func (s *CreateIpamPoolAllocationResponse) SetStatusCode(v int32) *CreateIpamPoolAllocationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIpamPoolAllocationResponse) SetBody(v *CreateIpamPoolAllocationResponseBody) *CreateIpamPoolAllocationResponse {
	s.Body = v
	return s
}

type CreateIpamResourceDiscoveryRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run, without sending the actual request. Valid value:
	//
	// 	- **true**: Performs the dry run without creating a custom resource discovery instance. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// 	- **false*	- (default): Performs a dry run and the actual request. If the request passes the dry run, an HTTP 2xx status code is returned and a custom resource discovery instance is created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The description of resource discovery.
	//
	// example:
	//
	// description
	IpamResourceDiscoveryDescription *string `json:"IpamResourceDiscoveryDescription,omitempty" xml:"IpamResourceDiscoveryDescription,omitempty"`
	// The name of the resource discovery.
	//
	// example:
	//
	// name
	IpamResourceDiscoveryName *string `json:"IpamResourceDiscoveryName,omitempty" xml:"IpamResourceDiscoveryName,omitempty"`
	// The list of effective regions.
	//
	// This parameter is required.
	OperatingRegionList []*string `json:"OperatingRegionList,omitempty" xml:"OperatingRegionList,omitempty" type:"Repeated"`
	OwnerAccount        *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId             *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The request region.
	//
	// >  The request region is the managed region of the resource discovery instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek2sermdd6****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tag information.
	Tag []*CreateIpamResourceDiscoveryRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateIpamResourceDiscoveryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIpamResourceDiscoveryRequest) GoString() string {
	return s.String()
}

func (s *CreateIpamResourceDiscoveryRequest) SetClientToken(v string) *CreateIpamResourceDiscoveryRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateIpamResourceDiscoveryRequest) SetDryRun(v bool) *CreateIpamResourceDiscoveryRequest {
	s.DryRun = &v
	return s
}

func (s *CreateIpamResourceDiscoveryRequest) SetIpamResourceDiscoveryDescription(v string) *CreateIpamResourceDiscoveryRequest {
	s.IpamResourceDiscoveryDescription = &v
	return s
}

func (s *CreateIpamResourceDiscoveryRequest) SetIpamResourceDiscoveryName(v string) *CreateIpamResourceDiscoveryRequest {
	s.IpamResourceDiscoveryName = &v
	return s
}

func (s *CreateIpamResourceDiscoveryRequest) SetOperatingRegionList(v []*string) *CreateIpamResourceDiscoveryRequest {
	s.OperatingRegionList = v
	return s
}

func (s *CreateIpamResourceDiscoveryRequest) SetOwnerAccount(v string) *CreateIpamResourceDiscoveryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateIpamResourceDiscoveryRequest) SetOwnerId(v int64) *CreateIpamResourceDiscoveryRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateIpamResourceDiscoveryRequest) SetRegionId(v string) *CreateIpamResourceDiscoveryRequest {
	s.RegionId = &v
	return s
}

func (s *CreateIpamResourceDiscoveryRequest) SetResourceGroupId(v string) *CreateIpamResourceDiscoveryRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateIpamResourceDiscoveryRequest) SetResourceOwnerAccount(v string) *CreateIpamResourceDiscoveryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateIpamResourceDiscoveryRequest) SetResourceOwnerId(v int64) *CreateIpamResourceDiscoveryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateIpamResourceDiscoveryRequest) SetTag(v []*CreateIpamResourceDiscoveryRequestTag) *CreateIpamResourceDiscoveryRequest {
	s.Tag = v
	return s
}

type CreateIpamResourceDiscoveryRequestTag struct {
	// The tag keys. You can specify at most 20 tag keys. It cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter but cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. You can specify up to 20 tag values. You can specify empty strings as tag values.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateIpamResourceDiscoveryRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateIpamResourceDiscoveryRequestTag) GoString() string {
	return s.String()
}

func (s *CreateIpamResourceDiscoveryRequestTag) SetKey(v string) *CreateIpamResourceDiscoveryRequestTag {
	s.Key = &v
	return s
}

func (s *CreateIpamResourceDiscoveryRequestTag) SetValue(v string) *CreateIpamResourceDiscoveryRequestTag {
	s.Value = &v
	return s
}

type CreateIpamResourceDiscoveryResponseBody struct {
	// The ID of the instance for resource discovery.
	//
	// example:
	//
	// ipam-res-disco-jt5f2af2u6nk2z321****
	IpamResourceDiscoveryId *string `json:"IpamResourceDiscoveryId,omitempty" xml:"IpamResourceDiscoveryId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BB2C39DE-CEB8-595A-981A-F2EFCBE7324E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateIpamResourceDiscoveryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIpamResourceDiscoveryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIpamResourceDiscoveryResponseBody) SetIpamResourceDiscoveryId(v string) *CreateIpamResourceDiscoveryResponseBody {
	s.IpamResourceDiscoveryId = &v
	return s
}

func (s *CreateIpamResourceDiscoveryResponseBody) SetRequestId(v string) *CreateIpamResourceDiscoveryResponseBody {
	s.RequestId = &v
	return s
}

type CreateIpamResourceDiscoveryResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIpamResourceDiscoveryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIpamResourceDiscoveryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIpamResourceDiscoveryResponse) GoString() string {
	return s.String()
}

func (s *CreateIpamResourceDiscoveryResponse) SetHeaders(v map[string]*string) *CreateIpamResourceDiscoveryResponse {
	s.Headers = v
	return s
}

func (s *CreateIpamResourceDiscoveryResponse) SetStatusCode(v int32) *CreateIpamResourceDiscoveryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIpamResourceDiscoveryResponse) SetBody(v *CreateIpamResourceDiscoveryResponseBody) *CreateIpamResourceDiscoveryResponse {
	s.Body = v
	return s
}

type CreateIpamScopeRequest struct {
	// The client token used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including invalid AccessKey pairs, unauthorized RAM users, and missing parameter values. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false**: performs a dry run and the actual request. After the request passes the dry run, a 2xx HTTP status code is returned and the IPAM scope is created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the IPAM.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-ccxbnsbhew0d6t****
	IpamId *string `json:"IpamId,omitempty" xml:"IpamId,omitempty"`
	// The description of the IPAM scope.
	//
	// It must be 1 to 256 characters in length. It must start with a letter but cannot start with `http://` or `https://`. This parameter is empty by default.
	//
	// example:
	//
	// test description
	IpamScopeDescription *string `json:"IpamScopeDescription,omitempty" xml:"IpamScopeDescription,omitempty"`
	// The name of the IPAM scope.
	//
	// It must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	IpamScopeName *string `json:"IpamScopeName,omitempty" xml:"IpamScopeName,omitempty"`
	// The type of IPAM scope: **private**
	//
	//
	// **Usage notes*	- You can create only private IPAM scopes.
	//
	// example:
	//
	// private
	IpamScopeType *string `json:"IpamScopeType,omitempty" xml:"IpamScopeType,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the IPAM instance is hosted. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID of the IPAM scope.
	//
	// example:
	//
	// rg-acfmxazb4ph6aiy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tag list.
	Tag []*CreateIpamScopeRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateIpamScopeRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIpamScopeRequest) GoString() string {
	return s.String()
}

func (s *CreateIpamScopeRequest) SetClientToken(v string) *CreateIpamScopeRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateIpamScopeRequest) SetDryRun(v bool) *CreateIpamScopeRequest {
	s.DryRun = &v
	return s
}

func (s *CreateIpamScopeRequest) SetIpamId(v string) *CreateIpamScopeRequest {
	s.IpamId = &v
	return s
}

func (s *CreateIpamScopeRequest) SetIpamScopeDescription(v string) *CreateIpamScopeRequest {
	s.IpamScopeDescription = &v
	return s
}

func (s *CreateIpamScopeRequest) SetIpamScopeName(v string) *CreateIpamScopeRequest {
	s.IpamScopeName = &v
	return s
}

func (s *CreateIpamScopeRequest) SetIpamScopeType(v string) *CreateIpamScopeRequest {
	s.IpamScopeType = &v
	return s
}

func (s *CreateIpamScopeRequest) SetOwnerAccount(v string) *CreateIpamScopeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateIpamScopeRequest) SetOwnerId(v int64) *CreateIpamScopeRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateIpamScopeRequest) SetRegionId(v string) *CreateIpamScopeRequest {
	s.RegionId = &v
	return s
}

func (s *CreateIpamScopeRequest) SetResourceGroupId(v string) *CreateIpamScopeRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateIpamScopeRequest) SetResourceOwnerAccount(v string) *CreateIpamScopeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateIpamScopeRequest) SetResourceOwnerId(v int64) *CreateIpamScopeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateIpamScopeRequest) SetTag(v []*CreateIpamScopeRequestTag) *CreateIpamScopeRequest {
	s.Tag = v
	return s
}

type CreateIpamScopeRequestTag struct {
	// The tag key of the resource. You can specify at most 20 tag keys. It cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter but cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. You can specify up to 20 tag values. You can specify empty strings as tag values.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateIpamScopeRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateIpamScopeRequestTag) GoString() string {
	return s.String()
}

func (s *CreateIpamScopeRequestTag) SetKey(v string) *CreateIpamScopeRequestTag {
	s.Key = &v
	return s
}

func (s *CreateIpamScopeRequestTag) SetValue(v string) *CreateIpamScopeRequestTag {
	s.Value = &v
	return s
}

type CreateIpamScopeResponseBody struct {
	// The ID of the IPAM scope.
	//
	// example:
	//
	// ipam-scope-glfmcyldpm8lsy****
	IpamScopeId *string `json:"IpamScopeId,omitempty" xml:"IpamScopeId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E897D16A-50EB-543F-B002-C5A26AB818FF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateIpamScopeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIpamScopeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIpamScopeResponseBody) SetIpamScopeId(v string) *CreateIpamScopeResponseBody {
	s.IpamScopeId = &v
	return s
}

func (s *CreateIpamScopeResponseBody) SetRequestId(v string) *CreateIpamScopeResponseBody {
	s.RequestId = &v
	return s
}

type CreateIpamScopeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIpamScopeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIpamScopeResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIpamScopeResponse) GoString() string {
	return s.String()
}

func (s *CreateIpamScopeResponse) SetHeaders(v map[string]*string) *CreateIpamScopeResponse {
	s.Headers = v
	return s
}

func (s *CreateIpamScopeResponse) SetStatusCode(v int32) *CreateIpamScopeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIpamScopeResponse) SetBody(v *CreateIpamScopeResponseBody) *CreateIpamScopeResponse {
	s.Body = v
	return s
}

type DeleteIpamRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// **
	//
	// **Usage notes*	- If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the IPAM.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-ccxbnsbhew0d6t****
	IpamId       *string `json:"IpamId,omitempty" xml:"IpamId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the IPAM instance is hosted. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteIpamRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpamRequest) GoString() string {
	return s.String()
}

func (s *DeleteIpamRequest) SetClientToken(v string) *DeleteIpamRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteIpamRequest) SetDryRun(v bool) *DeleteIpamRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteIpamRequest) SetIpamId(v string) *DeleteIpamRequest {
	s.IpamId = &v
	return s
}

func (s *DeleteIpamRequest) SetOwnerAccount(v string) *DeleteIpamRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteIpamRequest) SetOwnerId(v int64) *DeleteIpamRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteIpamRequest) SetRegionId(v string) *DeleteIpamRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteIpamRequest) SetResourceOwnerAccount(v string) *DeleteIpamRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteIpamRequest) SetResourceOwnerId(v int64) *DeleteIpamRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteIpamResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 30A20EE2-6223-5D0F-BF49-D7C78F206063
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIpamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpamResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIpamResponseBody) SetRequestId(v string) *DeleteIpamResponseBody {
	s.RequestId = &v
	return s
}

type DeleteIpamResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIpamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIpamResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpamResponse) GoString() string {
	return s.String()
}

func (s *DeleteIpamResponse) SetHeaders(v map[string]*string) *DeleteIpamResponse {
	s.Headers = v
	return s
}

func (s *DeleteIpamResponse) SetStatusCode(v int32) *DeleteIpamResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIpamResponse) SetBody(v *DeleteIpamResponseBody) *DeleteIpamResponse {
	s.Body = v
	return s
}

type DeleteIpamPoolRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// **
	//
	// **Usage notes*	- If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, DryRunOperation is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the IPAM pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-pool-6rcq3tobayc20t****
	IpamPoolId   *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the IPAM instance is hosted. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteIpamPoolRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpamPoolRequest) GoString() string {
	return s.String()
}

func (s *DeleteIpamPoolRequest) SetClientToken(v string) *DeleteIpamPoolRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteIpamPoolRequest) SetDryRun(v bool) *DeleteIpamPoolRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteIpamPoolRequest) SetIpamPoolId(v string) *DeleteIpamPoolRequest {
	s.IpamPoolId = &v
	return s
}

func (s *DeleteIpamPoolRequest) SetOwnerAccount(v string) *DeleteIpamPoolRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteIpamPoolRequest) SetOwnerId(v int64) *DeleteIpamPoolRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteIpamPoolRequest) SetRegionId(v string) *DeleteIpamPoolRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteIpamPoolRequest) SetResourceOwnerAccount(v string) *DeleteIpamPoolRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteIpamPoolRequest) SetResourceOwnerId(v int64) *DeleteIpamPoolRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteIpamPoolResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 57B7DCCA-F192-5528-8AF3-2FE1413228C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIpamPoolResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpamPoolResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIpamPoolResponseBody) SetRequestId(v string) *DeleteIpamPoolResponseBody {
	s.RequestId = &v
	return s
}

type DeleteIpamPoolResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIpamPoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIpamPoolResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpamPoolResponse) GoString() string {
	return s.String()
}

func (s *DeleteIpamPoolResponse) SetHeaders(v map[string]*string) *DeleteIpamPoolResponse {
	s.Headers = v
	return s
}

func (s *DeleteIpamPoolResponse) SetStatusCode(v int32) *DeleteIpamPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIpamPoolResponse) SetBody(v *DeleteIpamPoolResponseBody) *DeleteIpamPoolResponse {
	s.Body = v
	return s
}

type DeleteIpamPoolAllocationRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// **Usage notes*	- If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// 	- **false*	- (default): sends the API request. If the request passes the check, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the custom reserved CIDR block to be deleted.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-alloc-c4vhvr3b22mmc6****
	IpamPoolAllocationId *string `json:"IpamPoolAllocationId,omitempty" xml:"IpamPoolAllocationId,omitempty"`
	// The region ID of the custom reserved CIDR block.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteIpamPoolAllocationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpamPoolAllocationRequest) GoString() string {
	return s.String()
}

func (s *DeleteIpamPoolAllocationRequest) SetClientToken(v string) *DeleteIpamPoolAllocationRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteIpamPoolAllocationRequest) SetDryRun(v bool) *DeleteIpamPoolAllocationRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteIpamPoolAllocationRequest) SetIpamPoolAllocationId(v string) *DeleteIpamPoolAllocationRequest {
	s.IpamPoolAllocationId = &v
	return s
}

func (s *DeleteIpamPoolAllocationRequest) SetRegionId(v string) *DeleteIpamPoolAllocationRequest {
	s.RegionId = &v
	return s
}

type DeleteIpamPoolAllocationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B90776C8-F703-51D5-893A-AD1CA699D535
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIpamPoolAllocationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpamPoolAllocationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIpamPoolAllocationResponseBody) SetRequestId(v string) *DeleteIpamPoolAllocationResponseBody {
	s.RequestId = &v
	return s
}

type DeleteIpamPoolAllocationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIpamPoolAllocationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIpamPoolAllocationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpamPoolAllocationResponse) GoString() string {
	return s.String()
}

func (s *DeleteIpamPoolAllocationResponse) SetHeaders(v map[string]*string) *DeleteIpamPoolAllocationResponse {
	s.Headers = v
	return s
}

func (s *DeleteIpamPoolAllocationResponse) SetStatusCode(v int32) *DeleteIpamPoolAllocationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIpamPoolAllocationResponse) SetBody(v *DeleteIpamPoolAllocationResponseBody) *DeleteIpamPoolAllocationResponse {
	s.Body = v
	return s
}

type DeleteIpamPoolCidrRequest struct {
	// The provisioned CIDR block to be deleted.
	//
	// >  Only IPv4 CIDR blocks are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.1.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the IPAM pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-pool-6rcq3tobayc20t****
	IpamPoolId *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
	// The ID of the region where the IPAM instance is hosted.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteIpamPoolCidrRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpamPoolCidrRequest) GoString() string {
	return s.String()
}

func (s *DeleteIpamPoolCidrRequest) SetCidr(v string) *DeleteIpamPoolCidrRequest {
	s.Cidr = &v
	return s
}

func (s *DeleteIpamPoolCidrRequest) SetClientToken(v string) *DeleteIpamPoolCidrRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteIpamPoolCidrRequest) SetDryRun(v bool) *DeleteIpamPoolCidrRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteIpamPoolCidrRequest) SetIpamPoolId(v string) *DeleteIpamPoolCidrRequest {
	s.IpamPoolId = &v
	return s
}

func (s *DeleteIpamPoolCidrRequest) SetRegionId(v string) *DeleteIpamPoolCidrRequest {
	s.RegionId = &v
	return s
}

type DeleteIpamPoolCidrResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F28A239E-F88D-500E-ADE7-FA5E8CA3A170
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIpamPoolCidrResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpamPoolCidrResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIpamPoolCidrResponseBody) SetRequestId(v string) *DeleteIpamPoolCidrResponseBody {
	s.RequestId = &v
	return s
}

type DeleteIpamPoolCidrResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIpamPoolCidrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIpamPoolCidrResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpamPoolCidrResponse) GoString() string {
	return s.String()
}

func (s *DeleteIpamPoolCidrResponse) SetHeaders(v map[string]*string) *DeleteIpamPoolCidrResponse {
	s.Headers = v
	return s
}

func (s *DeleteIpamPoolCidrResponse) SetStatusCode(v int32) *DeleteIpamPoolCidrResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIpamPoolCidrResponse) SetBody(v *DeleteIpamPoolCidrResponseBody) *DeleteIpamPoolCidrResponse {
	s.Body = v
	return s
}

type DeleteIpamResourceDiscoveryRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform the dry run. Valid values:
	//
	// 	- **true**: Performs a dry run without deleting the resource discovery instance. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// 	- **false*	- (default): Performs a dry run and the actual request. If the request passes the check, an HTTP 2xx status code is returned and the resource discovery instance is deleted.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of resource discovery instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-res-disco-jt5f2af2u6nk2z321****
	IpamResourceDiscoveryId *string `json:"IpamResourceDiscoveryId,omitempty" xml:"IpamResourceDiscoveryId,omitempty"`
	OwnerAccount            *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                 *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The request region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteIpamResourceDiscoveryRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpamResourceDiscoveryRequest) GoString() string {
	return s.String()
}

func (s *DeleteIpamResourceDiscoveryRequest) SetClientToken(v string) *DeleteIpamResourceDiscoveryRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteIpamResourceDiscoveryRequest) SetDryRun(v bool) *DeleteIpamResourceDiscoveryRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteIpamResourceDiscoveryRequest) SetIpamResourceDiscoveryId(v string) *DeleteIpamResourceDiscoveryRequest {
	s.IpamResourceDiscoveryId = &v
	return s
}

func (s *DeleteIpamResourceDiscoveryRequest) SetOwnerAccount(v string) *DeleteIpamResourceDiscoveryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteIpamResourceDiscoveryRequest) SetOwnerId(v int64) *DeleteIpamResourceDiscoveryRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteIpamResourceDiscoveryRequest) SetRegionId(v string) *DeleteIpamResourceDiscoveryRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteIpamResourceDiscoveryRequest) SetResourceOwnerAccount(v string) *DeleteIpamResourceDiscoveryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteIpamResourceDiscoveryRequest) SetResourceOwnerId(v int64) *DeleteIpamResourceDiscoveryRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteIpamResourceDiscoveryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9F8315CB-560E-5F1E-B069-6E44B440CAF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIpamResourceDiscoveryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpamResourceDiscoveryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIpamResourceDiscoveryResponseBody) SetRequestId(v string) *DeleteIpamResourceDiscoveryResponseBody {
	s.RequestId = &v
	return s
}

type DeleteIpamResourceDiscoveryResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIpamResourceDiscoveryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIpamResourceDiscoveryResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpamResourceDiscoveryResponse) GoString() string {
	return s.String()
}

func (s *DeleteIpamResourceDiscoveryResponse) SetHeaders(v map[string]*string) *DeleteIpamResourceDiscoveryResponse {
	s.Headers = v
	return s
}

func (s *DeleteIpamResourceDiscoveryResponse) SetStatusCode(v int32) *DeleteIpamResourceDiscoveryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIpamResourceDiscoveryResponse) SetBody(v *DeleteIpamResourceDiscoveryResponseBody) *DeleteIpamResourceDiscoveryResponse {
	s.Body = v
	return s
}

type DeleteIpamScopeRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 88144bdb-b190-4813-a3f5-66cc86694162
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the IPAM scope.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-scope-glfmcyldpm8lsy****
	IpamScopeId  *string `json:"IpamScopeId,omitempty" xml:"IpamScopeId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the IPAM instance is hosted. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteIpamScopeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpamScopeRequest) GoString() string {
	return s.String()
}

func (s *DeleteIpamScopeRequest) SetClientToken(v string) *DeleteIpamScopeRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteIpamScopeRequest) SetDryRun(v bool) *DeleteIpamScopeRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteIpamScopeRequest) SetIpamScopeId(v string) *DeleteIpamScopeRequest {
	s.IpamScopeId = &v
	return s
}

func (s *DeleteIpamScopeRequest) SetOwnerAccount(v string) *DeleteIpamScopeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteIpamScopeRequest) SetOwnerId(v int64) *DeleteIpamScopeRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteIpamScopeRequest) SetRegionId(v string) *DeleteIpamScopeRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteIpamScopeRequest) SetResourceOwnerAccount(v string) *DeleteIpamScopeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteIpamScopeRequest) SetResourceOwnerId(v int64) *DeleteIpamScopeRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteIpamScopeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9F8315CB-560E-5F1E-B069-6E44B440CAF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIpamScopeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpamScopeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIpamScopeResponseBody) SetRequestId(v string) *DeleteIpamScopeResponseBody {
	s.RequestId = &v
	return s
}

type DeleteIpamScopeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIpamScopeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIpamScopeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpamScopeResponse) GoString() string {
	return s.String()
}

func (s *DeleteIpamScopeResponse) SetHeaders(v map[string]*string) *DeleteIpamScopeResponse {
	s.Headers = v
	return s
}

func (s *DeleteIpamScopeResponse) SetStatusCode(v int32) *DeleteIpamScopeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIpamScopeResponse) SetBody(v *DeleteIpamScopeResponseBody) *DeleteIpamScopeResponse {
	s.Body = v
	return s
}

type DissociateIpamResourceDiscoveryRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: Performs a dry run without disassociating the resource discovery and IPAM instance. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// 	- **false*	- (default): Performs a dry run and sends the request. After the request passes the check, an HTTP 2xx status code is returned and the resource discovery and IPAM instances are disassociated.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the IPAM.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-ccxbnsbhew0d6t****
	IpamId *string `json:"IpamId,omitempty" xml:"IpamId,omitempty"`
	// The ID of the resource discovery instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-res-disco-jt5f2af2u6nk2z321****
	IpamResourceDiscoveryId *string `json:"IpamResourceDiscoveryId,omitempty" xml:"IpamResourceDiscoveryId,omitempty"`
	OwnerAccount            *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                 *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The request region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DissociateIpamResourceDiscoveryRequest) String() string {
	return tea.Prettify(s)
}

func (s DissociateIpamResourceDiscoveryRequest) GoString() string {
	return s.String()
}

func (s *DissociateIpamResourceDiscoveryRequest) SetClientToken(v string) *DissociateIpamResourceDiscoveryRequest {
	s.ClientToken = &v
	return s
}

func (s *DissociateIpamResourceDiscoveryRequest) SetDryRun(v bool) *DissociateIpamResourceDiscoveryRequest {
	s.DryRun = &v
	return s
}

func (s *DissociateIpamResourceDiscoveryRequest) SetIpamId(v string) *DissociateIpamResourceDiscoveryRequest {
	s.IpamId = &v
	return s
}

func (s *DissociateIpamResourceDiscoveryRequest) SetIpamResourceDiscoveryId(v string) *DissociateIpamResourceDiscoveryRequest {
	s.IpamResourceDiscoveryId = &v
	return s
}

func (s *DissociateIpamResourceDiscoveryRequest) SetOwnerAccount(v string) *DissociateIpamResourceDiscoveryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DissociateIpamResourceDiscoveryRequest) SetOwnerId(v int64) *DissociateIpamResourceDiscoveryRequest {
	s.OwnerId = &v
	return s
}

func (s *DissociateIpamResourceDiscoveryRequest) SetRegionId(v string) *DissociateIpamResourceDiscoveryRequest {
	s.RegionId = &v
	return s
}

func (s *DissociateIpamResourceDiscoveryRequest) SetResourceOwnerAccount(v string) *DissociateIpamResourceDiscoveryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DissociateIpamResourceDiscoveryRequest) SetResourceOwnerId(v int64) *DissociateIpamResourceDiscoveryRequest {
	s.ResourceOwnerId = &v
	return s
}

type DissociateIpamResourceDiscoveryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 86137597-443F-5B66-B9B6-8514E0C50B8F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DissociateIpamResourceDiscoveryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DissociateIpamResourceDiscoveryResponseBody) GoString() string {
	return s.String()
}

func (s *DissociateIpamResourceDiscoveryResponseBody) SetRequestId(v string) *DissociateIpamResourceDiscoveryResponseBody {
	s.RequestId = &v
	return s
}

type DissociateIpamResourceDiscoveryResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DissociateIpamResourceDiscoveryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DissociateIpamResourceDiscoveryResponse) String() string {
	return tea.Prettify(s)
}

func (s DissociateIpamResourceDiscoveryResponse) GoString() string {
	return s.String()
}

func (s *DissociateIpamResourceDiscoveryResponse) SetHeaders(v map[string]*string) *DissociateIpamResourceDiscoveryResponse {
	s.Headers = v
	return s
}

func (s *DissociateIpamResourceDiscoveryResponse) SetStatusCode(v int32) *DissociateIpamResourceDiscoveryResponse {
	s.StatusCode = &v
	return s
}

func (s *DissociateIpamResourceDiscoveryResponse) SetBody(v *DissociateIpamResourceDiscoveryResponseBody) *DissociateIpamResourceDiscoveryResponse {
	s.Body = v
	return s
}

type GetIpamPoolAllocationRequest struct {
	// The ID of the instance to which CIDR blocks are allocated from the IPAM pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-pool-alloc-112za33e4****
	IpamPoolAllocationId *string `json:"IpamPoolAllocationId,omitempty" xml:"IpamPoolAllocationId,omitempty"`
	// The region of the IPAM pool.
	//
	// >  If the IPAM pool to which CIDR allocation belongs has the region attribute, this parameter is the region of the IPAM pool. If not, this parameter is the IPAM managed region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetIpamPoolAllocationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetIpamPoolAllocationRequest) GoString() string {
	return s.String()
}

func (s *GetIpamPoolAllocationRequest) SetIpamPoolAllocationId(v string) *GetIpamPoolAllocationRequest {
	s.IpamPoolAllocationId = &v
	return s
}

func (s *GetIpamPoolAllocationRequest) SetRegionId(v string) *GetIpamPoolAllocationRequest {
	s.RegionId = &v
	return s
}

type GetIpamPoolAllocationResponseBody struct {
	// The allocated CIDR block.
	//
	// example:
	//
	// 192.168.1.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The time when the instance was created.
	//
	// example:
	//
	// 2024-10-15T10:24:19+08:00
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the CIDR allocation of the IPAM pool.
	//
	// The description must be 1 to 256 characters in length and start with a letter, but cannot start with a `http://` or `https://`. This parameter is empty by default.
	//
	// example:
	//
	// ipam pool allocation description
	IpamPoolAllocationDescription *string `json:"IpamPoolAllocationDescription,omitempty" xml:"IpamPoolAllocationDescription,omitempty"`
	// The ID of the instance to which CIDR blocks are allocated from the IPAM pool.
	//
	// example:
	//
	// ipam-pool-alloc-112za33e4****
	IpamPoolAllocationId *string `json:"IpamPoolAllocationId,omitempty" xml:"IpamPoolAllocationId,omitempty"`
	// The name of the CIDR allocation of the IPAM pool.
	//
	// It must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// ipam pool allocation name
	IpamPoolAllocationName *string `json:"IpamPoolAllocationName,omitempty" xml:"IpamPoolAllocationName,omitempty"`
	// The ID of the IPAM pool.
	//
	// example:
	//
	// ipam-pool-6rcq3tobayc20t****
	IpamPoolId *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
	// The region of the IPAM pool.
	//
	// >  If the IPAM pool to which the CIDR block allocation belongs has the region attribute, this parameter is the region of the IPAM pool. If not, this parameter is the IPAM managed region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3748DEFF-68BE-5EED-9937-7C1D0C21BAB4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource to which the CIDR block is allocated.
	//
	// example:
	//
	// vpc-bp16qjewdsunr41m1****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The ID of the Alibaba Cloud account to which the resource belongs.
	//
	// example:
	//
	// 1616080591216318
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The effective region ID of the resource.
	//
	// example:
	//
	// cn-hangzhou
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// The type of the resource to which the CIDR block is allocated. Valid values:
	//
	// 	- **VPC**
	//
	// 	- **IpamPool**
	//
	// 	- **Custom**
	//
	// example:
	//
	// VPC
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The source CIDR block.
	//
	// example:
	//
	// 192.168.0.0/16
	SourceCidr *string `json:"SourceCidr,omitempty" xml:"SourceCidr,omitempty"`
	// The instance state. Valid values:
	//
	// 	- **Created**
	//
	// 	- **Deleted**
	//
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetIpamPoolAllocationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIpamPoolAllocationResponseBody) GoString() string {
	return s.String()
}

func (s *GetIpamPoolAllocationResponseBody) SetCidr(v string) *GetIpamPoolAllocationResponseBody {
	s.Cidr = &v
	return s
}

func (s *GetIpamPoolAllocationResponseBody) SetCreationTime(v string) *GetIpamPoolAllocationResponseBody {
	s.CreationTime = &v
	return s
}

func (s *GetIpamPoolAllocationResponseBody) SetIpamPoolAllocationDescription(v string) *GetIpamPoolAllocationResponseBody {
	s.IpamPoolAllocationDescription = &v
	return s
}

func (s *GetIpamPoolAllocationResponseBody) SetIpamPoolAllocationId(v string) *GetIpamPoolAllocationResponseBody {
	s.IpamPoolAllocationId = &v
	return s
}

func (s *GetIpamPoolAllocationResponseBody) SetIpamPoolAllocationName(v string) *GetIpamPoolAllocationResponseBody {
	s.IpamPoolAllocationName = &v
	return s
}

func (s *GetIpamPoolAllocationResponseBody) SetIpamPoolId(v string) *GetIpamPoolAllocationResponseBody {
	s.IpamPoolId = &v
	return s
}

func (s *GetIpamPoolAllocationResponseBody) SetRegionId(v string) *GetIpamPoolAllocationResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetIpamPoolAllocationResponseBody) SetRequestId(v string) *GetIpamPoolAllocationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIpamPoolAllocationResponseBody) SetResourceId(v string) *GetIpamPoolAllocationResponseBody {
	s.ResourceId = &v
	return s
}

func (s *GetIpamPoolAllocationResponseBody) SetResourceOwnerId(v int64) *GetIpamPoolAllocationResponseBody {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetIpamPoolAllocationResponseBody) SetResourceRegionId(v string) *GetIpamPoolAllocationResponseBody {
	s.ResourceRegionId = &v
	return s
}

func (s *GetIpamPoolAllocationResponseBody) SetResourceType(v string) *GetIpamPoolAllocationResponseBody {
	s.ResourceType = &v
	return s
}

func (s *GetIpamPoolAllocationResponseBody) SetSourceCidr(v string) *GetIpamPoolAllocationResponseBody {
	s.SourceCidr = &v
	return s
}

func (s *GetIpamPoolAllocationResponseBody) SetStatus(v string) *GetIpamPoolAllocationResponseBody {
	s.Status = &v
	return s
}

type GetIpamPoolAllocationResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIpamPoolAllocationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIpamPoolAllocationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIpamPoolAllocationResponse) GoString() string {
	return s.String()
}

func (s *GetIpamPoolAllocationResponse) SetHeaders(v map[string]*string) *GetIpamPoolAllocationResponse {
	s.Headers = v
	return s
}

func (s *GetIpamPoolAllocationResponse) SetStatusCode(v int32) *GetIpamPoolAllocationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIpamPoolAllocationResponse) SetBody(v *GetIpamPoolAllocationResponseBody) *GetIpamPoolAllocationResponse {
	s.Body = v
	return s
}

type GetIpamPoolNextAvailableCidrRequest struct {
	// CIDR to be allocated.
	//
	// >  You must enter at least one of the CidrBlock and CidrMask fields.
	//
	// example:
	//
	// 172.68.0.0/26
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The length of the CIDR mask to be allocated.
	//
	// >  You must enter at least one of the CidrBlock and CidrMask fields.
	//
	// example:
	//
	// 26
	CidrMask *int32 `json:"CidrMask,omitempty" xml:"CidrMask,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the IPAM pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-pool-6rcq3tobayc20t****
	IpamPoolId *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
	// The region of the IPAM pool.
	//
	// >  If the IPAM pool has the region attribute, this parameter is set to the effective region of the IPAM pool. If not, this is set to the managed region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetIpamPoolNextAvailableCidrRequest) String() string {
	return tea.Prettify(s)
}

func (s GetIpamPoolNextAvailableCidrRequest) GoString() string {
	return s.String()
}

func (s *GetIpamPoolNextAvailableCidrRequest) SetCidrBlock(v string) *GetIpamPoolNextAvailableCidrRequest {
	s.CidrBlock = &v
	return s
}

func (s *GetIpamPoolNextAvailableCidrRequest) SetCidrMask(v int32) *GetIpamPoolNextAvailableCidrRequest {
	s.CidrMask = &v
	return s
}

func (s *GetIpamPoolNextAvailableCidrRequest) SetClientToken(v string) *GetIpamPoolNextAvailableCidrRequest {
	s.ClientToken = &v
	return s
}

func (s *GetIpamPoolNextAvailableCidrRequest) SetIpamPoolId(v string) *GetIpamPoolNextAvailableCidrRequest {
	s.IpamPoolId = &v
	return s
}

func (s *GetIpamPoolNextAvailableCidrRequest) SetRegionId(v string) *GetIpamPoolNextAvailableCidrRequest {
	s.RegionId = &v
	return s
}

type GetIpamPoolNextAvailableCidrResponseBody struct {
	// Available CIDR.
	//
	// example:
	//
	// 172.68.0.0/26
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 29FC6758-9B7C-5CC7-8CBF-4DD846FE7D82
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetIpamPoolNextAvailableCidrResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIpamPoolNextAvailableCidrResponseBody) GoString() string {
	return s.String()
}

func (s *GetIpamPoolNextAvailableCidrResponseBody) SetCidrBlock(v string) *GetIpamPoolNextAvailableCidrResponseBody {
	s.CidrBlock = &v
	return s
}

func (s *GetIpamPoolNextAvailableCidrResponseBody) SetRequestId(v string) *GetIpamPoolNextAvailableCidrResponseBody {
	s.RequestId = &v
	return s
}

type GetIpamPoolNextAvailableCidrResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIpamPoolNextAvailableCidrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIpamPoolNextAvailableCidrResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIpamPoolNextAvailableCidrResponse) GoString() string {
	return s.String()
}

func (s *GetIpamPoolNextAvailableCidrResponse) SetHeaders(v map[string]*string) *GetIpamPoolNextAvailableCidrResponse {
	s.Headers = v
	return s
}

func (s *GetIpamPoolNextAvailableCidrResponse) SetStatusCode(v int32) *GetIpamPoolNextAvailableCidrResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIpamPoolNextAvailableCidrResponse) SetBody(v *GetIpamPoolNextAvailableCidrResponseBody) *GetIpamPoolNextAvailableCidrResponse {
	s.Body = v
	return s
}

type GetVpcIpamServiceStatusRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the IPAM instance is hosted.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetVpcIpamServiceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVpcIpamServiceStatusRequest) GoString() string {
	return s.String()
}

func (s *GetVpcIpamServiceStatusRequest) SetClientToken(v string) *GetVpcIpamServiceStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *GetVpcIpamServiceStatusRequest) SetOwnerAccount(v string) *GetVpcIpamServiceStatusRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetVpcIpamServiceStatusRequest) SetOwnerId(v int64) *GetVpcIpamServiceStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *GetVpcIpamServiceStatusRequest) SetRegionId(v string) *GetVpcIpamServiceStatusRequest {
	s.RegionId = &v
	return s
}

func (s *GetVpcIpamServiceStatusRequest) SetResourceOwnerAccount(v string) *GetVpcIpamServiceStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetVpcIpamServiceStatusRequest) SetResourceOwnerId(v int64) *GetVpcIpamServiceStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

type GetVpcIpamServiceStatusResponseBody struct {
	// Indicates whether IPAM is activated.
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2FEE9FFF-57EE-5832-BE88-9308352F3B68
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetVpcIpamServiceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVpcIpamServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetVpcIpamServiceStatusResponseBody) SetEnabled(v bool) *GetVpcIpamServiceStatusResponseBody {
	s.Enabled = &v
	return s
}

func (s *GetVpcIpamServiceStatusResponseBody) SetRequestId(v string) *GetVpcIpamServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

type GetVpcIpamServiceStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVpcIpamServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVpcIpamServiceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVpcIpamServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *GetVpcIpamServiceStatusResponse) SetHeaders(v map[string]*string) *GetVpcIpamServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *GetVpcIpamServiceStatusResponse) SetStatusCode(v int32) *GetVpcIpamServiceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVpcIpamServiceStatusResponse) SetBody(v *GetVpcIpamServiceStatusResponseBody) *GetVpcIpamServiceStatusResponse {
	s.Body = v
	return s
}

type ListIpamDiscoveredResourceRequest struct {
	// The ID of the resource discovery instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-res-disco-jt5f2af2u6nk2z321****
	IpamResourceDiscoveryId *string `json:"IpamResourceDiscoveryId,omitempty" xml:"IpamResourceDiscoveryId,omitempty"`
	// The maximum number of entries on each page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- You do not need to specify this parameter for the first request.
	//
	// 	- You must specify the token that is obtained from the previous query as the value of **NextToken**.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the managed region of the IPAM pool.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The region where resource discovery is performed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// The resource type. Valid values:
	//
	// 	- **VPC**
	//
	// 	- **VSwitch**
	//
	// example:
	//
	// VPC
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListIpamDiscoveredResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIpamDiscoveredResourceRequest) GoString() string {
	return s.String()
}

func (s *ListIpamDiscoveredResourceRequest) SetIpamResourceDiscoveryId(v string) *ListIpamDiscoveredResourceRequest {
	s.IpamResourceDiscoveryId = &v
	return s
}

func (s *ListIpamDiscoveredResourceRequest) SetMaxResults(v int32) *ListIpamDiscoveredResourceRequest {
	s.MaxResults = &v
	return s
}

func (s *ListIpamDiscoveredResourceRequest) SetNextToken(v string) *ListIpamDiscoveredResourceRequest {
	s.NextToken = &v
	return s
}

func (s *ListIpamDiscoveredResourceRequest) SetRegionId(v string) *ListIpamDiscoveredResourceRequest {
	s.RegionId = &v
	return s
}

func (s *ListIpamDiscoveredResourceRequest) SetResourceRegionId(v string) *ListIpamDiscoveredResourceRequest {
	s.ResourceRegionId = &v
	return s
}

func (s *ListIpamDiscoveredResourceRequest) SetResourceType(v string) *ListIpamDiscoveredResourceRequest {
	s.ResourceType = &v
	return s
}

type ListIpamDiscoveredResourceResponseBody struct {
	// The maximum number of entries on each page.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The list of resources.
	IpamDiscoveredResources []*ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources `json:"IpamDiscoveredResources,omitempty" xml:"IpamDiscoveredResources,omitempty" type:"Repeated"`
	// The maximum number of entries on each page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, there is no next page.
	//
	// 	- If a value of **NextToken*	- is returned, it indicates the token that is used for the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3748DEFF-68BE-5EED-9937-7C1D0C21BAB4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1000
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListIpamDiscoveredResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIpamDiscoveredResourceResponseBody) GoString() string {
	return s.String()
}

func (s *ListIpamDiscoveredResourceResponseBody) SetCount(v int32) *ListIpamDiscoveredResourceResponseBody {
	s.Count = &v
	return s
}

func (s *ListIpamDiscoveredResourceResponseBody) SetIpamDiscoveredResources(v []*ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources) *ListIpamDiscoveredResourceResponseBody {
	s.IpamDiscoveredResources = v
	return s
}

func (s *ListIpamDiscoveredResourceResponseBody) SetMaxResults(v int32) *ListIpamDiscoveredResourceResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListIpamDiscoveredResourceResponseBody) SetNextToken(v string) *ListIpamDiscoveredResourceResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListIpamDiscoveredResourceResponseBody) SetRequestId(v string) *ListIpamDiscoveredResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIpamDiscoveredResourceResponseBody) SetTotalCount(v int64) *ListIpamDiscoveredResourceResponseBody {
	s.TotalCount = &v
	return s
}

type ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 132193271328****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The CIDR block of the resource.
	//
	// example:
	//
	// 192.168.1.0/32
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The time when the resource was discovered.
	//
	// >  If the resource has not been modified since it was created, the discovery time remains unchanged.
	//
	// example:
	//
	// 2024-01-01 00:00:00
	DiscoveryTime *string                                                                     `json:"DiscoveryTime,omitempty" xml:"DiscoveryTime,omitempty"`
	IpCountDetail *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResourcesIpCountDetail `json:"IpCountDetail,omitempty" xml:"IpCountDetail,omitempty" type:"Struct"`
	// The IP usage in decimal form.
	//
	// example:
	//
	// 0
	IpUsage *string `json:"IpUsage,omitempty" xml:"IpUsage,omitempty"`
	// The ID of resource discovery instance.
	//
	// example:
	//
	// ipam-res-disco-jt5f2af2u6nk2z321****
	IpamResourceDiscoveryId *string `json:"IpamResourceDiscoveryId,omitempty" xml:"IpamResourceDiscoveryId,omitempty"`
	// The ID of the resource.
	//
	// example:
	//
	// vpc-uf611fp465c7dyb4z****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The ID of the Alibaba Cloud account to which the resource belongs.
	//
	// example:
	//
	// 132193271328****
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the region to which the resource belongs.
	//
	// example:
	//
	// cn-hangzhou
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// The resource type. Valid values:
	//
	// 	- **VPC**
	//
	// 	- **VSwitch**
	//
	// example:
	//
	// VPC
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The source CIDR block.
	//
	// example:
	//
	// 192.168.1.0/24
	SourceCidr *string `json:"SourceCidr,omitempty" xml:"SourceCidr,omitempty"`
	// The ID of the VPC to which the resource belongs.
	//
	// example:
	//
	// vpc-uf611fp465c7dyb4z****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources) String() string {
	return tea.Prettify(s)
}

func (s ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources) GoString() string {
	return s.String()
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources) SetAliUid(v int64) *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources {
	s.AliUid = &v
	return s
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources) SetCidr(v string) *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources {
	s.Cidr = &v
	return s
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources) SetDiscoveryTime(v string) *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources {
	s.DiscoveryTime = &v
	return s
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources) SetIpCountDetail(v *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResourcesIpCountDetail) *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources {
	s.IpCountDetail = v
	return s
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources) SetIpUsage(v string) *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources {
	s.IpUsage = &v
	return s
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources) SetIpamResourceDiscoveryId(v string) *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources {
	s.IpamResourceDiscoveryId = &v
	return s
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources) SetResourceId(v string) *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources {
	s.ResourceId = &v
	return s
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources) SetResourceOwnerId(v int64) *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources) SetResourceRegionId(v string) *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources {
	s.ResourceRegionId = &v
	return s
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources) SetResourceType(v string) *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources {
	s.ResourceType = &v
	return s
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources) SetSourceCidr(v string) *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources {
	s.SourceCidr = &v
	return s
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources) SetVpcId(v string) *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources {
	s.VpcId = &v
	return s
}

type ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResourcesIpCountDetail struct {
	FreeIpCount  *string `json:"FreeIpCount,omitempty" xml:"FreeIpCount,omitempty"`
	TotalIpCount *string `json:"TotalIpCount,omitempty" xml:"TotalIpCount,omitempty"`
	UsedIpCount  *string `json:"UsedIpCount,omitempty" xml:"UsedIpCount,omitempty"`
}

func (s ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResourcesIpCountDetail) String() string {
	return tea.Prettify(s)
}

func (s ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResourcesIpCountDetail) GoString() string {
	return s.String()
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResourcesIpCountDetail) SetFreeIpCount(v string) *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResourcesIpCountDetail {
	s.FreeIpCount = &v
	return s
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResourcesIpCountDetail) SetTotalIpCount(v string) *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResourcesIpCountDetail {
	s.TotalIpCount = &v
	return s
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResourcesIpCountDetail) SetUsedIpCount(v string) *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResourcesIpCountDetail {
	s.UsedIpCount = &v
	return s
}

type ListIpamDiscoveredResourceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIpamDiscoveredResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIpamDiscoveredResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIpamDiscoveredResourceResponse) GoString() string {
	return s.String()
}

func (s *ListIpamDiscoveredResourceResponse) SetHeaders(v map[string]*string) *ListIpamDiscoveredResourceResponse {
	s.Headers = v
	return s
}

func (s *ListIpamDiscoveredResourceResponse) SetStatusCode(v int32) *ListIpamDiscoveredResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIpamDiscoveredResourceResponse) SetBody(v *ListIpamDiscoveredResourceResponseBody) *ListIpamDiscoveredResourceResponse {
	s.Body = v
	return s
}

type ListIpamPoolAllocationsRequest struct {
	// Specifies whether to query allocations by specifying allocated CIDR blocks.
	//
	// **
	//
	// **Usage notes*	- Only IPv4 CIDR blocks are supported.
	//
	// example:
	//
	// 192.168.1.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The IDs of the instances to which CIDR blocks are allocated from the IPAM pool.
	IpamPoolAllocationIds []*string `json:"IpamPoolAllocationIds,omitempty" xml:"IpamPoolAllocationIds,omitempty" type:"Repeated"`
	// The name of  allocations.
	//
	// example:
	//
	// test name
	IpamPoolAllocationName *string `json:"IpamPoolAllocationName,omitempty" xml:"IpamPoolAllocationName,omitempty"`
	// The ID of the IPAM pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-pool-6rcq3tobayc20t****
	IpamPoolId *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
	// The number of entries per page. Valid values: **1*	- to **100**. Default value: **10**.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region where you want to perform the operation.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListIpamPoolAllocationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIpamPoolAllocationsRequest) GoString() string {
	return s.String()
}

func (s *ListIpamPoolAllocationsRequest) SetCidr(v string) *ListIpamPoolAllocationsRequest {
	s.Cidr = &v
	return s
}

func (s *ListIpamPoolAllocationsRequest) SetIpamPoolAllocationIds(v []*string) *ListIpamPoolAllocationsRequest {
	s.IpamPoolAllocationIds = v
	return s
}

func (s *ListIpamPoolAllocationsRequest) SetIpamPoolAllocationName(v string) *ListIpamPoolAllocationsRequest {
	s.IpamPoolAllocationName = &v
	return s
}

func (s *ListIpamPoolAllocationsRequest) SetIpamPoolId(v string) *ListIpamPoolAllocationsRequest {
	s.IpamPoolId = &v
	return s
}

func (s *ListIpamPoolAllocationsRequest) SetMaxResults(v int32) *ListIpamPoolAllocationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListIpamPoolAllocationsRequest) SetNextToken(v string) *ListIpamPoolAllocationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListIpamPoolAllocationsRequest) SetRegionId(v string) *ListIpamPoolAllocationsRequest {
	s.RegionId = &v
	return s
}

type ListIpamPoolAllocationsResponseBody struct {
	// The number of entries returned.
	//
	// example:
	//
	// 10
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The IDs of the instances to which CIDR blocks are allocated from the IPAM pool.
	IpamPoolAllocations []*ListIpamPoolAllocationsResponseBodyIpamPoolAllocations `json:"IpamPoolAllocations,omitempty" xml:"IpamPoolAllocations,omitempty" type:"Repeated"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3748DEFF-68BE-5EED-9937-7C1D0C21BAB4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1000
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListIpamPoolAllocationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIpamPoolAllocationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIpamPoolAllocationsResponseBody) SetCount(v int64) *ListIpamPoolAllocationsResponseBody {
	s.Count = &v
	return s
}

func (s *ListIpamPoolAllocationsResponseBody) SetIpamPoolAllocations(v []*ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) *ListIpamPoolAllocationsResponseBody {
	s.IpamPoolAllocations = v
	return s
}

func (s *ListIpamPoolAllocationsResponseBody) SetMaxResults(v int64) *ListIpamPoolAllocationsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListIpamPoolAllocationsResponseBody) SetNextToken(v string) *ListIpamPoolAllocationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListIpamPoolAllocationsResponseBody) SetRequestId(v string) *ListIpamPoolAllocationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIpamPoolAllocationsResponseBody) SetTotalCount(v int64) *ListIpamPoolAllocationsResponseBody {
	s.TotalCount = &v
	return s
}

type ListIpamPoolAllocationsResponseBodyIpamPoolAllocations struct {
	// The allocated CIDR block.
	//
	// example:
	//
	// 192.168.1.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The time when the instance was created.
	//
	// example:
	//
	// 2023-05-19T08:59:18Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the allocation.
	//
	// example:
	//
	// test description
	IpamPoolAllocationDescription *string `json:"IpamPoolAllocationDescription,omitempty" xml:"IpamPoolAllocationDescription,omitempty"`
	// The ID of the instance to which CIDR blocks are allocated from the IPAM pool.
	//
	// example:
	//
	// ipam-pool-alloc-112za33e4****
	IpamPoolAllocationId *string `json:"IpamPoolAllocationId,omitempty" xml:"IpamPoolAllocationId,omitempty"`
	// The name of the allocation.
	//
	// example:
	//
	// test name
	IpamPoolAllocationName *string `json:"IpamPoolAllocationName,omitempty" xml:"IpamPoolAllocationName,omitempty"`
	// The ID of the IPAM pool.
	//
	// example:
	//
	// ipam-pool-6rcq3tobayc20t****
	IpamPoolId *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
	// The region ID of the resource.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource to which the CIDR block is allocated.
	//
	// example:
	//
	// vpc-bp16qjewdsunr41m1****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The ID of the Alibaba Cloud account to which the resource belongs.
	//
	// example:
	//
	// 132193271328****
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The effective region ID of the resource.
	//
	// example:
	//
	// cn-hangzhou
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// The type of the resource to which the CIDR block is allocated. Valid values:
	//
	// 	- **VPC**
	//
	// 	- **IpamPool**
	//
	// 	- **Custom**
	//
	// example:
	//
	// Custom
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The source CIDR block.
	//
	// example:
	//
	// 192.168.0.0/16
	SourceCidr *string `json:"SourceCidr,omitempty" xml:"SourceCidr,omitempty"`
	// The status of the instance. Valid values:
	//
	// 	- **Created**
	//
	// 	- **Deleted**
	//
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) String() string {
	return tea.Prettify(s)
}

func (s ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) GoString() string {
	return s.String()
}

func (s *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) SetCidr(v string) *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations {
	s.Cidr = &v
	return s
}

func (s *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) SetCreationTime(v string) *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations {
	s.CreationTime = &v
	return s
}

func (s *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) SetIpamPoolAllocationDescription(v string) *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations {
	s.IpamPoolAllocationDescription = &v
	return s
}

func (s *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) SetIpamPoolAllocationId(v string) *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations {
	s.IpamPoolAllocationId = &v
	return s
}

func (s *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) SetIpamPoolAllocationName(v string) *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations {
	s.IpamPoolAllocationName = &v
	return s
}

func (s *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) SetIpamPoolId(v string) *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations {
	s.IpamPoolId = &v
	return s
}

func (s *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) SetRegionId(v string) *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations {
	s.RegionId = &v
	return s
}

func (s *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) SetResourceId(v string) *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations {
	s.ResourceId = &v
	return s
}

func (s *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) SetResourceOwnerId(v int64) *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) SetResourceRegionId(v string) *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations {
	s.ResourceRegionId = &v
	return s
}

func (s *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) SetResourceType(v string) *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations {
	s.ResourceType = &v
	return s
}

func (s *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) SetSourceCidr(v string) *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations {
	s.SourceCidr = &v
	return s
}

func (s *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) SetStatus(v string) *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations {
	s.Status = &v
	return s
}

type ListIpamPoolAllocationsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIpamPoolAllocationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIpamPoolAllocationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIpamPoolAllocationsResponse) GoString() string {
	return s.String()
}

func (s *ListIpamPoolAllocationsResponse) SetHeaders(v map[string]*string) *ListIpamPoolAllocationsResponse {
	s.Headers = v
	return s
}

func (s *ListIpamPoolAllocationsResponse) SetStatusCode(v int32) *ListIpamPoolAllocationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIpamPoolAllocationsResponse) SetBody(v *ListIpamPoolAllocationsResponseBody) *ListIpamPoolAllocationsResponse {
	s.Body = v
	return s
}

type ListIpamPoolCidrsRequest struct {
	// The provisioned CIDR block that you want to query.
	//
	// >  Only IPv4 CIDR blocks are supported.
	//
	// example:
	//
	// 192.168.1.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The ID of the IPAM pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-pool-6rcq3tobayc20t****
	IpamPoolId *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
	// The number of entries per page. Valid values: **1*	- to **100**. Default value: **10**.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region where the IPAM instance is hosted.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListIpamPoolCidrsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIpamPoolCidrsRequest) GoString() string {
	return s.String()
}

func (s *ListIpamPoolCidrsRequest) SetCidr(v string) *ListIpamPoolCidrsRequest {
	s.Cidr = &v
	return s
}

func (s *ListIpamPoolCidrsRequest) SetIpamPoolId(v string) *ListIpamPoolCidrsRequest {
	s.IpamPoolId = &v
	return s
}

func (s *ListIpamPoolCidrsRequest) SetMaxResults(v int32) *ListIpamPoolCidrsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListIpamPoolCidrsRequest) SetNextToken(v string) *ListIpamPoolCidrsRequest {
	s.NextToken = &v
	return s
}

func (s *ListIpamPoolCidrsRequest) SetRegionId(v string) *ListIpamPoolCidrsRequest {
	s.RegionId = &v
	return s
}

type ListIpamPoolCidrsResponseBody struct {
	// The number of entries returned.
	//
	// example:
	//
	// 10
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The IDs of IPAM pools.
	IpamPoolCidrs []*ListIpamPoolCidrsResponseBodyIpamPoolCidrs `json:"IpamPoolCidrs,omitempty" xml:"IpamPoolCidrs,omitempty" type:"Repeated"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9E7CCB95-62E0-534D-9B9A-71F27E8B71B1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1000
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListIpamPoolCidrsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIpamPoolCidrsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIpamPoolCidrsResponseBody) SetCount(v int64) *ListIpamPoolCidrsResponseBody {
	s.Count = &v
	return s
}

func (s *ListIpamPoolCidrsResponseBody) SetIpamPoolCidrs(v []*ListIpamPoolCidrsResponseBodyIpamPoolCidrs) *ListIpamPoolCidrsResponseBody {
	s.IpamPoolCidrs = v
	return s
}

func (s *ListIpamPoolCidrsResponseBody) SetMaxResults(v int64) *ListIpamPoolCidrsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListIpamPoolCidrsResponseBody) SetNextToken(v string) *ListIpamPoolCidrsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListIpamPoolCidrsResponseBody) SetRequestId(v string) *ListIpamPoolCidrsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIpamPoolCidrsResponseBody) SetTotalCount(v int64) *ListIpamPoolCidrsResponseBody {
	s.TotalCount = &v
	return s
}

type ListIpamPoolCidrsResponseBodyIpamPoolCidrs struct {
	// The provisioned CIDR block.
	//
	// example:
	//
	// 192.168.1.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The ID of the IPAM pool.
	//
	// example:
	//
	// ipam-pool-6rcq3tobayc20t****
	IpamPoolId *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
	// The status of the CIDR block provisioned to the IPAM pool. Valid values:
	//
	// 	- **Created**
	//
	// 	- **Deleted**
	//
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListIpamPoolCidrsResponseBodyIpamPoolCidrs) String() string {
	return tea.Prettify(s)
}

func (s ListIpamPoolCidrsResponseBodyIpamPoolCidrs) GoString() string {
	return s.String()
}

func (s *ListIpamPoolCidrsResponseBodyIpamPoolCidrs) SetCidr(v string) *ListIpamPoolCidrsResponseBodyIpamPoolCidrs {
	s.Cidr = &v
	return s
}

func (s *ListIpamPoolCidrsResponseBodyIpamPoolCidrs) SetIpamPoolId(v string) *ListIpamPoolCidrsResponseBodyIpamPoolCidrs {
	s.IpamPoolId = &v
	return s
}

func (s *ListIpamPoolCidrsResponseBodyIpamPoolCidrs) SetStatus(v string) *ListIpamPoolCidrsResponseBodyIpamPoolCidrs {
	s.Status = &v
	return s
}

type ListIpamPoolCidrsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIpamPoolCidrsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIpamPoolCidrsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIpamPoolCidrsResponse) GoString() string {
	return s.String()
}

func (s *ListIpamPoolCidrsResponse) SetHeaders(v map[string]*string) *ListIpamPoolCidrsResponse {
	s.Headers = v
	return s
}

func (s *ListIpamPoolCidrsResponse) SetStatusCode(v int32) *ListIpamPoolCidrsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIpamPoolCidrsResponse) SetBody(v *ListIpamPoolCidrsResponseBody) *ListIpamPoolCidrsResponse {
	s.Body = v
	return s
}

type ListIpamPoolsRequest struct {
	// The IDs of IPAM pools. Valid values of N: 1 to 100. A maximum of 100 IPAM pools can be queried at a time.
	IpamPoolIds []*string `json:"IpamPoolIds,omitempty" xml:"IpamPoolIds,omitempty" type:"Repeated"`
	// The name of the IPAM pool. You can enter at most 20 names.
	//
	// It must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	IpamPoolName *string `json:"IpamPoolName,omitempty" xml:"IpamPoolName,omitempty"`
	// The ID of the IPAM scope.
	//
	// example:
	//
	// ipam-scope-glfmcyldpm8lsy****
	IpamScopeId *string `json:"IpamScopeId,omitempty" xml:"IpamScopeId,omitempty"`
	// Whether it is a shared pool.
	//
	// example:
	//
	// true
	IsShared *bool `json:"IsShared,omitempty" xml:"IsShared,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If NextToken is empty, no next page exists.
	//
	// 	- You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The effective region of the IPAM pool.
	//
	// example:
	//
	// cn-hangzhou
	PoolRegionId *string `json:"PoolRegionId,omitempty" xml:"PoolRegionId,omitempty"`
	// The ID of the region where the IPAM instance is hosted. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the IPAM pool belongs.
	//
	// example:
	//
	// rg-aek2sermdd6****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the source IPAM pool.
	//
	// example:
	//
	// ipam-pool-lfnwi4jok1ss0g****
	SourceIpamPoolId *string `json:"SourceIpamPoolId,omitempty" xml:"SourceIpamPoolId,omitempty"`
	// The tag information.
	Tags []*ListIpamPoolsRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListIpamPoolsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIpamPoolsRequest) GoString() string {
	return s.String()
}

func (s *ListIpamPoolsRequest) SetIpamPoolIds(v []*string) *ListIpamPoolsRequest {
	s.IpamPoolIds = v
	return s
}

func (s *ListIpamPoolsRequest) SetIpamPoolName(v string) *ListIpamPoolsRequest {
	s.IpamPoolName = &v
	return s
}

func (s *ListIpamPoolsRequest) SetIpamScopeId(v string) *ListIpamPoolsRequest {
	s.IpamScopeId = &v
	return s
}

func (s *ListIpamPoolsRequest) SetIsShared(v bool) *ListIpamPoolsRequest {
	s.IsShared = &v
	return s
}

func (s *ListIpamPoolsRequest) SetMaxResults(v int32) *ListIpamPoolsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListIpamPoolsRequest) SetNextToken(v string) *ListIpamPoolsRequest {
	s.NextToken = &v
	return s
}

func (s *ListIpamPoolsRequest) SetOwnerAccount(v string) *ListIpamPoolsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListIpamPoolsRequest) SetOwnerId(v int64) *ListIpamPoolsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListIpamPoolsRequest) SetPoolRegionId(v string) *ListIpamPoolsRequest {
	s.PoolRegionId = &v
	return s
}

func (s *ListIpamPoolsRequest) SetRegionId(v string) *ListIpamPoolsRequest {
	s.RegionId = &v
	return s
}

func (s *ListIpamPoolsRequest) SetResourceGroupId(v string) *ListIpamPoolsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListIpamPoolsRequest) SetResourceOwnerAccount(v string) *ListIpamPoolsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListIpamPoolsRequest) SetResourceOwnerId(v int64) *ListIpamPoolsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListIpamPoolsRequest) SetSourceIpamPoolId(v string) *ListIpamPoolsRequest {
	s.SourceIpamPoolId = &v
	return s
}

func (s *ListIpamPoolsRequest) SetTags(v []*ListIpamPoolsRequestTags) *ListIpamPoolsRequest {
	s.Tags = v
	return s
}

type ListIpamPoolsRequestTags struct {
	// The tag key. You can specify at most 20 tag keys. It cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The tag key must start with a letter but cannot start with `aliyun` or `acs:`. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can specify at most 20 tag values. It can be an empty string.
	//
	// The tag value can be up to 128 characters in length. It must start with a letter and can contain digits, periods (.), underscores (_), and hyphens (-). It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListIpamPoolsRequestTags) String() string {
	return tea.Prettify(s)
}

func (s ListIpamPoolsRequestTags) GoString() string {
	return s.String()
}

func (s *ListIpamPoolsRequestTags) SetKey(v string) *ListIpamPoolsRequestTags {
	s.Key = &v
	return s
}

func (s *ListIpamPoolsRequestTags) SetValue(v string) *ListIpamPoolsRequestTags {
	s.Value = &v
	return s
}

type ListIpamPoolsResponseBody struct {
	// The number of entries returned.
	//
	// example:
	//
	// 10
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The IPAM pools.
	IpamPools []*ListIpamPoolsResponseBodyIpamPools `json:"IpamPools,omitempty" xml:"IpamPools,omitempty" type:"Repeated"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B54867DE-83DC-56B4-A57E-69A03119D0B1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1000
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListIpamPoolsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIpamPoolsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIpamPoolsResponseBody) SetCount(v int64) *ListIpamPoolsResponseBody {
	s.Count = &v
	return s
}

func (s *ListIpamPoolsResponseBody) SetIpamPools(v []*ListIpamPoolsResponseBodyIpamPools) *ListIpamPoolsResponseBody {
	s.IpamPools = v
	return s
}

func (s *ListIpamPoolsResponseBody) SetMaxResults(v int64) *ListIpamPoolsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListIpamPoolsResponseBody) SetNextToken(v string) *ListIpamPoolsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListIpamPoolsResponseBody) SetRequestId(v string) *ListIpamPoolsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIpamPoolsResponseBody) SetTotalCount(v int64) *ListIpamPoolsResponseBody {
	s.TotalCount = &v
	return s
}

type ListIpamPoolsResponseBodyIpamPools struct {
	// The default network mask assigned to the IPAM pool.
	//
	// An IPv4 mask must be **0 to 32*	- bits in length.
	//
	// example:
	//
	// 28
	AllocationDefaultCidrMask *int32 `json:"AllocationDefaultCidrMask,omitempty" xml:"AllocationDefaultCidrMask,omitempty"`
	// The maximum network mask assigned to the IPAM pool.
	//
	// An IPv4 mask must be **0 to 32*	- bits in length.
	//
	// example:
	//
	// 32
	AllocationMaxCidrMask *int32 `json:"AllocationMaxCidrMask,omitempty" xml:"AllocationMaxCidrMask,omitempty"`
	// The minimum network mask assigned to the IPAM pool.
	//
	// An IPv4 mask must be **0 to 32*	- bits in length.
	//
	// example:
	//
	// 8
	AllocationMinCidrMask *int32 `json:"AllocationMinCidrMask,omitempty" xml:"AllocationMinCidrMask,omitempty"`
	// Whether the pool has the auto-import feature enabled.
	//
	// example:
	//
	// true
	AutoImport *bool `json:"AutoImport,omitempty" xml:"AutoImport,omitempty"`
	// The time when the IPAM pool was created.
	//
	// example:
	//
	// 2023-04-19T16:49:01Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the pool is a subpool. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	HasSubPool *bool `json:"HasSubPool,omitempty" xml:"HasSubPool,omitempty"`
	// The IP version. Only **IPv4*	- may be returned.
	//
	// example:
	//
	// IPv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The ID of the IPAM.
	//
	// example:
	//
	// ipam-b5mtlx3q7xcnyr****
	IpamId *string `json:"IpamId,omitempty" xml:"IpamId,omitempty"`
	// The description of the IPAM pool.
	//
	// example:
	//
	// test description
	IpamPoolDescription *string `json:"IpamPoolDescription,omitempty" xml:"IpamPoolDescription,omitempty"`
	// The ID of the IPAM pool.
	//
	// example:
	//
	// ipam-pool-6rcq3tobayc20t****
	IpamPoolId *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
	// The name of the IPAM pool.
	//
	// example:
	//
	// test
	IpamPoolName *string `json:"IpamPoolName,omitempty" xml:"IpamPoolName,omitempty"`
	// The ID of the region where the IPAM to which the IPAM pool belongs is hosted.
	//
	// example:
	//
	// cn-hangzhou
	IpamRegionId *string `json:"IpamRegionId,omitempty" xml:"IpamRegionId,omitempty"`
	// The ID of the IPAM scope.
	//
	// example:
	//
	// ipam-scope-glfmcyldpm8lsy****
	IpamScopeId *string `json:"IpamScopeId,omitempty" xml:"IpamScopeId,omitempty"`
	// The type of the IPAM scope. Valid values:
	//
	// 	- **public**
	//
	// 	- **private**
	//
	// example:
	//
	// private
	IpamScopeType *string `json:"IpamScopeType,omitempty" xml:"IpamScopeType,omitempty"`
	// Whether it is a shared pool.
	//
	// example:
	//
	// true
	IsShared *bool `json:"IsShared,omitempty" xml:"IsShared,omitempty"`
	// The Alibaba Cloud account of the owner for the IPAM pool.
	//
	// example:
	//
	// 1210123456******
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The depth of the IPAM pool. Valid values: **0 to 10**.
	//
	// example:
	//
	// 2
	PoolDepth *int32 `json:"PoolDepth,omitempty" xml:"PoolDepth,omitempty"`
	// The effective region of the IPAM pool. The ID of the effective region for the IPAM pool.
	//
	// example:
	//
	// cn-hangzhou
	PoolRegionId *string `json:"PoolRegionId,omitempty" xml:"PoolRegionId,omitempty"`
	// The ID of the region where the operation is called.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxazb4ph6aiy****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the source IPAM pool.
	//
	// example:
	//
	// ipam-pool-lfnwi4jok1ss0g****
	SourceIpamPoolId *string `json:"SourceIpamPoolId,omitempty" xml:"SourceIpamPoolId,omitempty"`
	// The status of the IPAM pool. Valid values:
	//
	// 	- **Creating**
	//
	// 	- **Created**: indicates that the creation is complete.
	//
	// 	- **Modifying**
	//
	// 	- **Deleting**
	//
	// 	- **Deleted**
	//
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag list.
	Tags []*ListIpamPoolsResponseBodyIpamPoolsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListIpamPoolsResponseBodyIpamPools) String() string {
	return tea.Prettify(s)
}

func (s ListIpamPoolsResponseBodyIpamPools) GoString() string {
	return s.String()
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetAllocationDefaultCidrMask(v int32) *ListIpamPoolsResponseBodyIpamPools {
	s.AllocationDefaultCidrMask = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetAllocationMaxCidrMask(v int32) *ListIpamPoolsResponseBodyIpamPools {
	s.AllocationMaxCidrMask = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetAllocationMinCidrMask(v int32) *ListIpamPoolsResponseBodyIpamPools {
	s.AllocationMinCidrMask = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetAutoImport(v bool) *ListIpamPoolsResponseBodyIpamPools {
	s.AutoImport = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetCreateTime(v string) *ListIpamPoolsResponseBodyIpamPools {
	s.CreateTime = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetHasSubPool(v bool) *ListIpamPoolsResponseBodyIpamPools {
	s.HasSubPool = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetIpVersion(v string) *ListIpamPoolsResponseBodyIpamPools {
	s.IpVersion = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetIpamId(v string) *ListIpamPoolsResponseBodyIpamPools {
	s.IpamId = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetIpamPoolDescription(v string) *ListIpamPoolsResponseBodyIpamPools {
	s.IpamPoolDescription = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetIpamPoolId(v string) *ListIpamPoolsResponseBodyIpamPools {
	s.IpamPoolId = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetIpamPoolName(v string) *ListIpamPoolsResponseBodyIpamPools {
	s.IpamPoolName = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetIpamRegionId(v string) *ListIpamPoolsResponseBodyIpamPools {
	s.IpamRegionId = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetIpamScopeId(v string) *ListIpamPoolsResponseBodyIpamPools {
	s.IpamScopeId = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetIpamScopeType(v string) *ListIpamPoolsResponseBodyIpamPools {
	s.IpamScopeType = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetIsShared(v bool) *ListIpamPoolsResponseBodyIpamPools {
	s.IsShared = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetOwnerId(v int64) *ListIpamPoolsResponseBodyIpamPools {
	s.OwnerId = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetPoolDepth(v int32) *ListIpamPoolsResponseBodyIpamPools {
	s.PoolDepth = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetPoolRegionId(v string) *ListIpamPoolsResponseBodyIpamPools {
	s.PoolRegionId = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetRegionId(v string) *ListIpamPoolsResponseBodyIpamPools {
	s.RegionId = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetResourceGroupId(v string) *ListIpamPoolsResponseBodyIpamPools {
	s.ResourceGroupId = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetSourceIpamPoolId(v string) *ListIpamPoolsResponseBodyIpamPools {
	s.SourceIpamPoolId = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetStatus(v string) *ListIpamPoolsResponseBodyIpamPools {
	s.Status = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetTags(v []*ListIpamPoolsResponseBodyIpamPoolsTags) *ListIpamPoolsResponseBodyIpamPools {
	s.Tags = v
	return s
}

type ListIpamPoolsResponseBodyIpamPoolsTags struct {
	// The tag key.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListIpamPoolsResponseBodyIpamPoolsTags) String() string {
	return tea.Prettify(s)
}

func (s ListIpamPoolsResponseBodyIpamPoolsTags) GoString() string {
	return s.String()
}

func (s *ListIpamPoolsResponseBodyIpamPoolsTags) SetKey(v string) *ListIpamPoolsResponseBodyIpamPoolsTags {
	s.Key = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPoolsTags) SetValue(v string) *ListIpamPoolsResponseBodyIpamPoolsTags {
	s.Value = &v
	return s
}

type ListIpamPoolsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIpamPoolsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIpamPoolsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIpamPoolsResponse) GoString() string {
	return s.String()
}

func (s *ListIpamPoolsResponse) SetHeaders(v map[string]*string) *ListIpamPoolsResponse {
	s.Headers = v
	return s
}

func (s *ListIpamPoolsResponse) SetStatusCode(v int32) *ListIpamPoolsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIpamPoolsResponse) SetBody(v *ListIpamPoolsResponseBody) *ListIpamPoolsResponse {
	s.Body = v
	return s
}

type ListIpamResourceCidrsRequest struct {
	// The ID of the IPAM pool.
	//
	// >  You must specify at least one of **IpamScopeId*	- and **IpamPoolId**.
	//
	// example:
	//
	// ipam-pool-6rcq3tobayc20t****
	IpamPoolId *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
	// The ID of the IPAM scope.
	//
	// >  You must specify at least one of **IpamScopeId*	- and **IpamPoolId**.
	//
	// example:
	//
	// ipam-scope-glfmcyldpm8lsy****
	IpamScopeId *string `json:"IpamScopeId,omitempty" xml:"IpamScopeId,omitempty"`
	// The number of entries per page. Valid values: **1*	- to **100**. Default value: **10**.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- You do not need to specify this parameter for the first request.
	//
	// 	- You must specify the token that is obtained from the previous query as the value of **NextToken**.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region where the IPAM instance is hosted.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// vpc-bp16qjewdsunr41m1****
	ResourceId      *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of resource. Valid values:
	//
	// 	- **VPC**
	//
	// 	- **VSwitch**
	//
	// example:
	//
	// VPC
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-bp1fjfnrg3av6zb9e****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListIpamResourceCidrsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIpamResourceCidrsRequest) GoString() string {
	return s.String()
}

func (s *ListIpamResourceCidrsRequest) SetIpamPoolId(v string) *ListIpamResourceCidrsRequest {
	s.IpamPoolId = &v
	return s
}

func (s *ListIpamResourceCidrsRequest) SetIpamScopeId(v string) *ListIpamResourceCidrsRequest {
	s.IpamScopeId = &v
	return s
}

func (s *ListIpamResourceCidrsRequest) SetMaxResults(v int32) *ListIpamResourceCidrsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListIpamResourceCidrsRequest) SetNextToken(v string) *ListIpamResourceCidrsRequest {
	s.NextToken = &v
	return s
}

func (s *ListIpamResourceCidrsRequest) SetRegionId(v string) *ListIpamResourceCidrsRequest {
	s.RegionId = &v
	return s
}

func (s *ListIpamResourceCidrsRequest) SetResourceId(v string) *ListIpamResourceCidrsRequest {
	s.ResourceId = &v
	return s
}

func (s *ListIpamResourceCidrsRequest) SetResourceOwnerId(v int64) *ListIpamResourceCidrsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListIpamResourceCidrsRequest) SetResourceType(v string) *ListIpamResourceCidrsRequest {
	s.ResourceType = &v
	return s
}

func (s *ListIpamResourceCidrsRequest) SetVpcId(v string) *ListIpamResourceCidrsRequest {
	s.VpcId = &v
	return s
}

type ListIpamResourceCidrsResponseBody struct {
	// The number of entries returned.
	//
	// example:
	//
	// 10
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The list of resources in the IPAM pool.
	IpamResourceCidrs []*ListIpamResourceCidrsResponseBodyIpamResourceCidrs `json:"IpamResourceCidrs,omitempty" xml:"IpamResourceCidrs,omitempty" type:"Repeated"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 49A9DE56-B68C-5FFC-BC06-509D086F287C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1000
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListIpamResourceCidrsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIpamResourceCidrsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIpamResourceCidrsResponseBody) SetCount(v int64) *ListIpamResourceCidrsResponseBody {
	s.Count = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBody) SetIpamResourceCidrs(v []*ListIpamResourceCidrsResponseBodyIpamResourceCidrs) *ListIpamResourceCidrsResponseBody {
	s.IpamResourceCidrs = v
	return s
}

func (s *ListIpamResourceCidrsResponseBody) SetMaxResults(v int64) *ListIpamResourceCidrsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBody) SetNextToken(v string) *ListIpamResourceCidrsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBody) SetRequestId(v string) *ListIpamResourceCidrsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBody) SetTotalCount(v int64) *ListIpamResourceCidrsResponseBody {
	s.TotalCount = &v
	return s
}

type ListIpamResourceCidrsResponseBodyIpamResourceCidrs struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 132193271328****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The CIDR block of the resource.
	//
	// example:
	//
	// 192.168.1.0/32
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The compliance status of the resource.
	//
	// 	- **Compliant**
	//
	// 	- **Noncompliant**
	//
	// 	- **Ignored*	- Ignored resources are not monitored.
	//
	// 	- **Unmanaged**: The resource does not have a CIDR block allocated from the IPAM pool. IPAM does not monitor whether the CIDR block of the resource meets the allocation rules of the IP address pool.
	//
	// example:
	//
	// Compliant
	ComplianceStatus *string                                                          `json:"ComplianceStatus,omitempty" xml:"ComplianceStatus,omitempty"`
	IpCountDetail    *ListIpamResourceCidrsResponseBodyIpamResourceCidrsIpCountDetail `json:"IpCountDetail,omitempty" xml:"IpCountDetail,omitempty" type:"Struct"`
	// The IP usage that is displayed in decimal form.
	//
	// example:
	//
	// 0
	IpUsage *string `json:"IpUsage,omitempty" xml:"IpUsage,omitempty"`
	// The ID of the instance to which CIDR blocks are allocated from the IPAM pool.
	//
	// example:
	//
	// ipam-pool-alloc-112za33e4****
	IpamAllocationId *string `json:"IpamAllocationId,omitempty" xml:"IpamAllocationId,omitempty"`
	// The ID of the IPAM.
	//
	// example:
	//
	// ipam-uq5dcfc2eqhpf4****
	IpamId *string `json:"IpamId,omitempty" xml:"IpamId,omitempty"`
	// The ID of the IPAM pool.
	//
	// example:
	//
	// ipam-pool-6rcq3tobayc20t***
	IpamPoolId *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
	// The ID of the IPAM scope.
	//
	// example:
	//
	// ipam-scope-glfmcyldpm8lsy****
	IpamScopeId *string `json:"IpamScopeId,omitempty" xml:"IpamScopeId,omitempty"`
	// The management status of the resource.
	//
	// 	- **Managed**: The resource has a CIDR block allocated from an IPAM pool. IPAM is monitoring whether the allocated CIDR block overlaps with other CIDR blocks and whether the allocated CIDR block meets the allocation rules.
	//
	// 	- **Unmanaged**: The resource does not have a CIDR block allocated from the IPAM pool. IPAM is monitoring whether the resource has CIDR blocks that meet the allocation rules. Monitor whether CIDR blocks overlap with each other.
	//
	// 	- **Ignored**: The resource is not monitored. Ignored resources are not monitored. If you ignore a resource, CIDR blocks allocated to the resource are returned to the IPAM pool and will not be automatically allocated to the resource (if automatic allocation rules are specified).
	//
	// example:
	//
	// Managed
	ManagementStatus *string `json:"ManagementStatus,omitempty" xml:"ManagementStatus,omitempty"`
	// List of resources that overlap with the current resource.
	OverlapDetail []*ListIpamResourceCidrsResponseBodyIpamResourceCidrsOverlapDetail `json:"OverlapDetail,omitempty" xml:"OverlapDetail,omitempty" type:"Repeated"`
	// The overlapping status of the resource.
	//
	// 	- **Nonoverlapping**
	//
	// 	- **Overlapping**
	//
	// 	- **Ignored*	- Ignored resources are not monitored.
	//
	// example:
	//
	// Nonoverlapping
	OverlapStatus *string `json:"OverlapStatus,omitempty" xml:"OverlapStatus,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// vpc-bp16qjewdsunr41m1****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The ID of the Alibaba Cloud account to which the resource belongs.
	//
	// example:
	//
	// 132193271328****
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The effective region ID of the resource.
	//
	// example:
	//
	// cn-hangzhou
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// The type of resource. Valid values:
	//
	// 	- **VPC**
	//
	// 	- **VSwitch**
	//
	// example:
	//
	// VPC
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The source CIDR block.
	//
	// example:
	//
	// 192.168.1.0/24
	SourceCidr *string `json:"SourceCidr,omitempty" xml:"SourceCidr,omitempty"`
	// The status of the resource in the IPAM pool. Valid values:
	//
	// 	- **Created**
	//
	// 	- **Deleted**
	//
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-bp1fjfnrg3av6zb9e****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListIpamResourceCidrsResponseBodyIpamResourceCidrs) String() string {
	return tea.Prettify(s)
}

func (s ListIpamResourceCidrsResponseBodyIpamResourceCidrs) GoString() string {
	return s.String()
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) SetAliUid(v int64) *ListIpamResourceCidrsResponseBodyIpamResourceCidrs {
	s.AliUid = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) SetCidr(v string) *ListIpamResourceCidrsResponseBodyIpamResourceCidrs {
	s.Cidr = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) SetComplianceStatus(v string) *ListIpamResourceCidrsResponseBodyIpamResourceCidrs {
	s.ComplianceStatus = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) SetIpCountDetail(v *ListIpamResourceCidrsResponseBodyIpamResourceCidrsIpCountDetail) *ListIpamResourceCidrsResponseBodyIpamResourceCidrs {
	s.IpCountDetail = v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) SetIpUsage(v string) *ListIpamResourceCidrsResponseBodyIpamResourceCidrs {
	s.IpUsage = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) SetIpamAllocationId(v string) *ListIpamResourceCidrsResponseBodyIpamResourceCidrs {
	s.IpamAllocationId = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) SetIpamId(v string) *ListIpamResourceCidrsResponseBodyIpamResourceCidrs {
	s.IpamId = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) SetIpamPoolId(v string) *ListIpamResourceCidrsResponseBodyIpamResourceCidrs {
	s.IpamPoolId = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) SetIpamScopeId(v string) *ListIpamResourceCidrsResponseBodyIpamResourceCidrs {
	s.IpamScopeId = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) SetManagementStatus(v string) *ListIpamResourceCidrsResponseBodyIpamResourceCidrs {
	s.ManagementStatus = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) SetOverlapDetail(v []*ListIpamResourceCidrsResponseBodyIpamResourceCidrsOverlapDetail) *ListIpamResourceCidrsResponseBodyIpamResourceCidrs {
	s.OverlapDetail = v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) SetOverlapStatus(v string) *ListIpamResourceCidrsResponseBodyIpamResourceCidrs {
	s.OverlapStatus = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) SetResourceId(v string) *ListIpamResourceCidrsResponseBodyIpamResourceCidrs {
	s.ResourceId = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) SetResourceOwnerId(v int64) *ListIpamResourceCidrsResponseBodyIpamResourceCidrs {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) SetResourceRegionId(v string) *ListIpamResourceCidrsResponseBodyIpamResourceCidrs {
	s.ResourceRegionId = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) SetResourceType(v string) *ListIpamResourceCidrsResponseBodyIpamResourceCidrs {
	s.ResourceType = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) SetSourceCidr(v string) *ListIpamResourceCidrsResponseBodyIpamResourceCidrs {
	s.SourceCidr = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) SetStatus(v string) *ListIpamResourceCidrsResponseBodyIpamResourceCidrs {
	s.Status = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) SetVpcId(v string) *ListIpamResourceCidrsResponseBodyIpamResourceCidrs {
	s.VpcId = &v
	return s
}

type ListIpamResourceCidrsResponseBodyIpamResourceCidrsIpCountDetail struct {
	FreeIpCount  *string `json:"FreeIpCount,omitempty" xml:"FreeIpCount,omitempty"`
	TotalIpCount *string `json:"TotalIpCount,omitempty" xml:"TotalIpCount,omitempty"`
	UsedIpCount  *string `json:"UsedIpCount,omitempty" xml:"UsedIpCount,omitempty"`
}

func (s ListIpamResourceCidrsResponseBodyIpamResourceCidrsIpCountDetail) String() string {
	return tea.Prettify(s)
}

func (s ListIpamResourceCidrsResponseBodyIpamResourceCidrsIpCountDetail) GoString() string {
	return s.String()
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrsIpCountDetail) SetFreeIpCount(v string) *ListIpamResourceCidrsResponseBodyIpamResourceCidrsIpCountDetail {
	s.FreeIpCount = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrsIpCountDetail) SetTotalIpCount(v string) *ListIpamResourceCidrsResponseBodyIpamResourceCidrsIpCountDetail {
	s.TotalIpCount = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrsIpCountDetail) SetUsedIpCount(v string) *ListIpamResourceCidrsResponseBodyIpamResourceCidrsIpCountDetail {
	s.UsedIpCount = &v
	return s
}

type ListIpamResourceCidrsResponseBodyIpamResourceCidrsOverlapDetail struct {
	// The CIDR that overlaps with the current resource.
	//
	// example:
	//
	// 192.168.1.0/24
	OverlapResourceCidr *string `json:"OverlapResourceCidr,omitempty" xml:"OverlapResourceCidr,omitempty"`
	// Instance ID that overlaps with the current resource.
	//
	// example:
	//
	// vpc-aq3fjgnig5av6jb8d****
	OverlapResourceId *string `json:"OverlapResourceId,omitempty" xml:"OverlapResourceId,omitempty"`
	// The region of instance that overlaps with the current resource.
	//
	// example:
	//
	// cn-hangzhou
	OverlapResourceRegion *string `json:"OverlapResourceRegion,omitempty" xml:"OverlapResourceRegion,omitempty"`
}

func (s ListIpamResourceCidrsResponseBodyIpamResourceCidrsOverlapDetail) String() string {
	return tea.Prettify(s)
}

func (s ListIpamResourceCidrsResponseBodyIpamResourceCidrsOverlapDetail) GoString() string {
	return s.String()
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrsOverlapDetail) SetOverlapResourceCidr(v string) *ListIpamResourceCidrsResponseBodyIpamResourceCidrsOverlapDetail {
	s.OverlapResourceCidr = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrsOverlapDetail) SetOverlapResourceId(v string) *ListIpamResourceCidrsResponseBodyIpamResourceCidrsOverlapDetail {
	s.OverlapResourceId = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrsOverlapDetail) SetOverlapResourceRegion(v string) *ListIpamResourceCidrsResponseBodyIpamResourceCidrsOverlapDetail {
	s.OverlapResourceRegion = &v
	return s
}

type ListIpamResourceCidrsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIpamResourceCidrsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIpamResourceCidrsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIpamResourceCidrsResponse) GoString() string {
	return s.String()
}

func (s *ListIpamResourceCidrsResponse) SetHeaders(v map[string]*string) *ListIpamResourceCidrsResponse {
	s.Headers = v
	return s
}

func (s *ListIpamResourceCidrsResponse) SetStatusCode(v int32) *ListIpamResourceCidrsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIpamResourceCidrsResponse) SetBody(v *ListIpamResourceCidrsResponseBody) *ListIpamResourceCidrsResponse {
	s.Body = v
	return s
}

type ListIpamResourceDiscoveriesRequest struct {
	// The IDs of resource discovery instances. Valid values of N: 1 to 100. A maximum of 100 resource discoveries can be queried at a time.
	IpamResourceDiscoveryIds []*string `json:"IpamResourceDiscoveryIds,omitempty" xml:"IpamResourceDiscoveryIds,omitempty" type:"Repeated"`
	// The name of the resource discovery.
	//
	// The name must be 1 to 128 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// test
	IpamResourceDiscoveryName *string `json:"IpamResourceDiscoveryName,omitempty" xml:"IpamResourceDiscoveryName,omitempty"`
	// Whether it is a shared resource discovery.
	//
	// example:
	//
	// true
	IsShared *bool `json:"IsShared,omitempty" xml:"IsShared,omitempty"`
	// The maximum number of entries on each page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, there is no next page.
	//
	// 	- If a value of **NextToken*	- is returned, it indicates the token that is used for the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where you want to query resource discovery. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group that resource discovery belongs.
	//
	// example:
	//
	// rg-aek2sermdd6****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tag.
	Tags []*ListIpamResourceDiscoveriesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The type of resource discovery.
	//
	// > Supported types:
	//
	// > - system: default resource discovery created by the system.
	//
	// > - custom: custom resource discovery created by users.
	//
	// example:
	//
	// system
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListIpamResourceDiscoveriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIpamResourceDiscoveriesRequest) GoString() string {
	return s.String()
}

func (s *ListIpamResourceDiscoveriesRequest) SetIpamResourceDiscoveryIds(v []*string) *ListIpamResourceDiscoveriesRequest {
	s.IpamResourceDiscoveryIds = v
	return s
}

func (s *ListIpamResourceDiscoveriesRequest) SetIpamResourceDiscoveryName(v string) *ListIpamResourceDiscoveriesRequest {
	s.IpamResourceDiscoveryName = &v
	return s
}

func (s *ListIpamResourceDiscoveriesRequest) SetIsShared(v bool) *ListIpamResourceDiscoveriesRequest {
	s.IsShared = &v
	return s
}

func (s *ListIpamResourceDiscoveriesRequest) SetMaxResults(v int32) *ListIpamResourceDiscoveriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListIpamResourceDiscoveriesRequest) SetNextToken(v string) *ListIpamResourceDiscoveriesRequest {
	s.NextToken = &v
	return s
}

func (s *ListIpamResourceDiscoveriesRequest) SetOwnerAccount(v string) *ListIpamResourceDiscoveriesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListIpamResourceDiscoveriesRequest) SetOwnerId(v int64) *ListIpamResourceDiscoveriesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListIpamResourceDiscoveriesRequest) SetRegionId(v string) *ListIpamResourceDiscoveriesRequest {
	s.RegionId = &v
	return s
}

func (s *ListIpamResourceDiscoveriesRequest) SetResourceGroupId(v string) *ListIpamResourceDiscoveriesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListIpamResourceDiscoveriesRequest) SetResourceOwnerAccount(v string) *ListIpamResourceDiscoveriesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListIpamResourceDiscoveriesRequest) SetResourceOwnerId(v int64) *ListIpamResourceDiscoveriesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListIpamResourceDiscoveriesRequest) SetTags(v []*ListIpamResourceDiscoveriesRequestTags) *ListIpamResourceDiscoveriesRequest {
	s.Tags = v
	return s
}

func (s *ListIpamResourceDiscoveriesRequest) SetType(v string) *ListIpamResourceDiscoveriesRequest {
	s.Type = &v
	return s
}

type ListIpamResourceDiscoveriesRequestTags struct {
	// The key of the tag. You can specify at most 20 tag keys. It cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The tag key must start with a letter but cannot start with `aliyun` or `acs:`. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag. You can specify at most 20 tag values. The tag value cannot be an empty string.
	//
	// A tag value can be up to 128 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListIpamResourceDiscoveriesRequestTags) String() string {
	return tea.Prettify(s)
}

func (s ListIpamResourceDiscoveriesRequestTags) GoString() string {
	return s.String()
}

func (s *ListIpamResourceDiscoveriesRequestTags) SetKey(v string) *ListIpamResourceDiscoveriesRequestTags {
	s.Key = &v
	return s
}

func (s *ListIpamResourceDiscoveriesRequestTags) SetValue(v string) *ListIpamResourceDiscoveriesRequestTags {
	s.Value = &v
	return s
}

type ListIpamResourceDiscoveriesResponseBody struct {
	// The maximum number of entries on each page.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The list of resource discovery instances.
	IpamResourceDiscoveries []*ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries `json:"IpamResourceDiscoveries,omitempty" xml:"IpamResourceDiscoveries,omitempty" type:"Repeated"`
	// The maximum number of entries on each page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, there is no next page.
	//
	// 	- If a value of **NextToken*	- is returned, it indicates the token that is used for the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 86137597-443F-5B66-B9B6-8514E0C50B8F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListIpamResourceDiscoveriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIpamResourceDiscoveriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListIpamResourceDiscoveriesResponseBody) SetCount(v int32) *ListIpamResourceDiscoveriesResponseBody {
	s.Count = &v
	return s
}

func (s *ListIpamResourceDiscoveriesResponseBody) SetIpamResourceDiscoveries(v []*ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries) *ListIpamResourceDiscoveriesResponseBody {
	s.IpamResourceDiscoveries = v
	return s
}

func (s *ListIpamResourceDiscoveriesResponseBody) SetMaxResults(v int32) *ListIpamResourceDiscoveriesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListIpamResourceDiscoveriesResponseBody) SetNextToken(v string) *ListIpamResourceDiscoveriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListIpamResourceDiscoveriesResponseBody) SetRequestId(v string) *ListIpamResourceDiscoveriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIpamResourceDiscoveriesResponseBody) SetTotalCount(v int64) *ListIpamResourceDiscoveriesResponseBody {
	s.TotalCount = &v
	return s
}

type ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries struct {
	// The time when the resource was discovered.
	//
	// example:
	//
	// 2022-07-01T02:05:23Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the resource discovery.
	//
	// example:
	//
	// test description
	IpamResourceDiscoveryDescription *string `json:"IpamResourceDiscoveryDescription,omitempty" xml:"IpamResourceDiscoveryDescription,omitempty"`
	// The ID of resource discovery instance.
	//
	// example:
	//
	// ipam-res-disco-jt5f2af2u6nk2z321****
	IpamResourceDiscoveryId *string `json:"IpamResourceDiscoveryId,omitempty" xml:"IpamResourceDiscoveryId,omitempty"`
	// The name of the resource discovery.
	//
	// example:
	//
	// test
	IpamResourceDiscoveryName *string `json:"IpamResourceDiscoveryName,omitempty" xml:"IpamResourceDiscoveryName,omitempty"`
	// The status of the resource discovery instance. Valid values:
	//
	// 	- **Creating**
	//
	// 	- **Created**
	//
	// 	- **Modifying**
	//
	// 	- **Deleting**
	//
	// 	- **Deleted**
	//
	// example:
	//
	// Created
	IpamResourceDiscoveryStatus *string `json:"IpamResourceDiscoveryStatus,omitempty" xml:"IpamResourceDiscoveryStatus,omitempty"`
	// The list of resource discovery regions.
	OperatingRegionList []*string `json:"OperatingRegionList,omitempty" xml:"OperatingRegionList,omitempty" type:"Repeated"`
	// The Alibaba Cloud account that owns the resource discovery.
	//
	// example:
	//
	// 1210123456******
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the queried resource discovery instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group that resource discovery belongs.
	//
	// example:
	//
	// rg-aek2sermdd6****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The sharing status of the resource.
	//
	// 	- If the value is empty, the resource is as an average instance.
	//
	// 	- If the value is Shared, the resource discovery comes from a shared source.
	//
	// 	- If the value is Sharing, the resource discovery is being shared.
	//
	// example:
	//
	// Shared
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The tag list.
	Tags []*ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveriesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The type of resource discovery.
	//
	// example:
	//
	// system
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries) String() string {
	return tea.Prettify(s)
}

func (s ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries) GoString() string {
	return s.String()
}

func (s *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries) SetCreateTime(v string) *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries {
	s.CreateTime = &v
	return s
}

func (s *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries) SetIpamResourceDiscoveryDescription(v string) *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries {
	s.IpamResourceDiscoveryDescription = &v
	return s
}

func (s *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries) SetIpamResourceDiscoveryId(v string) *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries {
	s.IpamResourceDiscoveryId = &v
	return s
}

func (s *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries) SetIpamResourceDiscoveryName(v string) *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries {
	s.IpamResourceDiscoveryName = &v
	return s
}

func (s *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries) SetIpamResourceDiscoveryStatus(v string) *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries {
	s.IpamResourceDiscoveryStatus = &v
	return s
}

func (s *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries) SetOperatingRegionList(v []*string) *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries {
	s.OperatingRegionList = v
	return s
}

func (s *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries) SetOwnerId(v int64) *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries {
	s.OwnerId = &v
	return s
}

func (s *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries) SetRegionId(v string) *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries {
	s.RegionId = &v
	return s
}

func (s *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries) SetResourceGroupId(v string) *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries {
	s.ResourceGroupId = &v
	return s
}

func (s *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries) SetShareType(v string) *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries {
	s.ShareType = &v
	return s
}

func (s *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries) SetTags(v []*ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveriesTags) *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries {
	s.Tags = v
	return s
}

func (s *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries) SetType(v string) *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries {
	s.Type = &v
	return s
}

type ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveriesTags struct {
	// The tag key.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveriesTags) String() string {
	return tea.Prettify(s)
}

func (s ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveriesTags) GoString() string {
	return s.String()
}

func (s *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveriesTags) SetKey(v string) *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveriesTags {
	s.Key = &v
	return s
}

func (s *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveriesTags) SetValue(v string) *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveriesTags {
	s.Value = &v
	return s
}

type ListIpamResourceDiscoveriesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIpamResourceDiscoveriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIpamResourceDiscoveriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIpamResourceDiscoveriesResponse) GoString() string {
	return s.String()
}

func (s *ListIpamResourceDiscoveriesResponse) SetHeaders(v map[string]*string) *ListIpamResourceDiscoveriesResponse {
	s.Headers = v
	return s
}

func (s *ListIpamResourceDiscoveriesResponse) SetStatusCode(v int32) *ListIpamResourceDiscoveriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIpamResourceDiscoveriesResponse) SetBody(v *ListIpamResourceDiscoveriesResponseBody) *ListIpamResourceDiscoveriesResponse {
	s.Body = v
	return s
}

type ListIpamResourceDiscoveryAssociationsRequest struct {
	// The ID of the IPAM.
	//
	// example:
	//
	// ipam-ccxbnsbhew0d6t****
	IpamId *string `json:"IpamId,omitempty" xml:"IpamId,omitempty"`
	// The ID of resource discovery instance.
	//
	// example:
	//
	// ipam-res-disco-jt5f2af2u6nk2z321****
	IpamResourceDiscoveryId *string `json:"IpamResourceDiscoveryId,omitempty" xml:"IpamResourceDiscoveryId,omitempty"`
	// The maximum number of entries on each page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If this is your first or only query, this parameter is left empty.
	//
	// 	- If a next query is to be sent, the returned value is the value of NextToken that was returned last time this operation was called.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The request region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListIpamResourceDiscoveryAssociationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIpamResourceDiscoveryAssociationsRequest) GoString() string {
	return s.String()
}

func (s *ListIpamResourceDiscoveryAssociationsRequest) SetIpamId(v string) *ListIpamResourceDiscoveryAssociationsRequest {
	s.IpamId = &v
	return s
}

func (s *ListIpamResourceDiscoveryAssociationsRequest) SetIpamResourceDiscoveryId(v string) *ListIpamResourceDiscoveryAssociationsRequest {
	s.IpamResourceDiscoveryId = &v
	return s
}

func (s *ListIpamResourceDiscoveryAssociationsRequest) SetMaxResults(v int32) *ListIpamResourceDiscoveryAssociationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListIpamResourceDiscoveryAssociationsRequest) SetNextToken(v string) *ListIpamResourceDiscoveryAssociationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListIpamResourceDiscoveryAssociationsRequest) SetOwnerAccount(v string) *ListIpamResourceDiscoveryAssociationsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListIpamResourceDiscoveryAssociationsRequest) SetOwnerId(v int64) *ListIpamResourceDiscoveryAssociationsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListIpamResourceDiscoveryAssociationsRequest) SetRegionId(v string) *ListIpamResourceDiscoveryAssociationsRequest {
	s.RegionId = &v
	return s
}

func (s *ListIpamResourceDiscoveryAssociationsRequest) SetResourceOwnerAccount(v string) *ListIpamResourceDiscoveryAssociationsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListIpamResourceDiscoveryAssociationsRequest) SetResourceOwnerId(v int64) *ListIpamResourceDiscoveryAssociationsRequest {
	s.ResourceOwnerId = &v
	return s
}

type ListIpamResourceDiscoveryAssociationsResponseBody struct {
	// The number of entries on each page.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The list of associations.
	IpamResourceDiscoveryAssociations []*ListIpamResourceDiscoveryAssociationsResponseBodyIpamResourceDiscoveryAssociations `json:"IpamResourceDiscoveryAssociations,omitempty" xml:"IpamResourceDiscoveryAssociations,omitempty" type:"Repeated"`
	// The maximum number of entries on each page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, there is no next page.
	//
	// 	- If a value of **NextToken*	- is returned, it indicates the token that is used for the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F28A239E-F88D-500E-ADE7-FA5E8CA3A170
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 20
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListIpamResourceDiscoveryAssociationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIpamResourceDiscoveryAssociationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIpamResourceDiscoveryAssociationsResponseBody) SetCount(v int32) *ListIpamResourceDiscoveryAssociationsResponseBody {
	s.Count = &v
	return s
}

func (s *ListIpamResourceDiscoveryAssociationsResponseBody) SetIpamResourceDiscoveryAssociations(v []*ListIpamResourceDiscoveryAssociationsResponseBodyIpamResourceDiscoveryAssociations) *ListIpamResourceDiscoveryAssociationsResponseBody {
	s.IpamResourceDiscoveryAssociations = v
	return s
}

func (s *ListIpamResourceDiscoveryAssociationsResponseBody) SetMaxResults(v int32) *ListIpamResourceDiscoveryAssociationsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListIpamResourceDiscoveryAssociationsResponseBody) SetNextToken(v string) *ListIpamResourceDiscoveryAssociationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListIpamResourceDiscoveryAssociationsResponseBody) SetRequestId(v string) *ListIpamResourceDiscoveryAssociationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIpamResourceDiscoveryAssociationsResponseBody) SetTotalCount(v int64) *ListIpamResourceDiscoveryAssociationsResponseBody {
	s.TotalCount = &v
	return s
}

type ListIpamResourceDiscoveryAssociationsResponseBodyIpamResourceDiscoveryAssociations struct {
	// The ID of the IPAM.
	//
	// example:
	//
	// ipam-ccxbnsbhew0d6t****
	IpamId *string `json:"IpamId,omitempty" xml:"IpamId,omitempty"`
	// The ID of resource discovery instance.
	//
	// example:
	//
	// ipam-res-disco-jt5f2af2u6nk2z321****
	IpamResourceDiscoveryId *string `json:"IpamResourceDiscoveryId,omitempty" xml:"IpamResourceDiscoveryId,omitempty"`
	// The ID of the Alibaba Cloud account to which the resource discovery belongs.
	//
	// example:
	//
	// 1210123456******
	IpamResourceDiscoveryOwnerId *string `json:"IpamResourceDiscoveryOwnerId,omitempty" xml:"IpamResourceDiscoveryOwnerId,omitempty"`
	// The status of the resource discovery instance. Valid values:
	//
	// 	- **Creating**
	//
	// 	- **Created**
	//
	// 	- **Modifying**
	//
	// 	- **Deleting**
	//
	// 	- **Deleted**
	//
	// example:
	//
	// Created
	IpamResourceDiscoveryStatus *string `json:"IpamResourceDiscoveryStatus,omitempty" xml:"IpamResourceDiscoveryStatus,omitempty"`
	// The type of resource discovery. Valid values:
	//
	// 	- **system**: default resource discovery created by the system.
	//
	// 	- **custom**: custom resource discovery created by users.
	//
	// example:
	//
	// custom
	IpamResourceDiscoveryType *string `json:"IpamResourceDiscoveryType,omitempty" xml:"IpamResourceDiscoveryType,omitempty"`
	// The status of the associations. Valid values:
	//
	// 	- **Created**
	//
	// 	- **Deleted**
	//
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListIpamResourceDiscoveryAssociationsResponseBodyIpamResourceDiscoveryAssociations) String() string {
	return tea.Prettify(s)
}

func (s ListIpamResourceDiscoveryAssociationsResponseBodyIpamResourceDiscoveryAssociations) GoString() string {
	return s.String()
}

func (s *ListIpamResourceDiscoveryAssociationsResponseBodyIpamResourceDiscoveryAssociations) SetIpamId(v string) *ListIpamResourceDiscoveryAssociationsResponseBodyIpamResourceDiscoveryAssociations {
	s.IpamId = &v
	return s
}

func (s *ListIpamResourceDiscoveryAssociationsResponseBodyIpamResourceDiscoveryAssociations) SetIpamResourceDiscoveryId(v string) *ListIpamResourceDiscoveryAssociationsResponseBodyIpamResourceDiscoveryAssociations {
	s.IpamResourceDiscoveryId = &v
	return s
}

func (s *ListIpamResourceDiscoveryAssociationsResponseBodyIpamResourceDiscoveryAssociations) SetIpamResourceDiscoveryOwnerId(v string) *ListIpamResourceDiscoveryAssociationsResponseBodyIpamResourceDiscoveryAssociations {
	s.IpamResourceDiscoveryOwnerId = &v
	return s
}

func (s *ListIpamResourceDiscoveryAssociationsResponseBodyIpamResourceDiscoveryAssociations) SetIpamResourceDiscoveryStatus(v string) *ListIpamResourceDiscoveryAssociationsResponseBodyIpamResourceDiscoveryAssociations {
	s.IpamResourceDiscoveryStatus = &v
	return s
}

func (s *ListIpamResourceDiscoveryAssociationsResponseBodyIpamResourceDiscoveryAssociations) SetIpamResourceDiscoveryType(v string) *ListIpamResourceDiscoveryAssociationsResponseBodyIpamResourceDiscoveryAssociations {
	s.IpamResourceDiscoveryType = &v
	return s
}

func (s *ListIpamResourceDiscoveryAssociationsResponseBodyIpamResourceDiscoveryAssociations) SetStatus(v string) *ListIpamResourceDiscoveryAssociationsResponseBodyIpamResourceDiscoveryAssociations {
	s.Status = &v
	return s
}

type ListIpamResourceDiscoveryAssociationsResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIpamResourceDiscoveryAssociationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIpamResourceDiscoveryAssociationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIpamResourceDiscoveryAssociationsResponse) GoString() string {
	return s.String()
}

func (s *ListIpamResourceDiscoveryAssociationsResponse) SetHeaders(v map[string]*string) *ListIpamResourceDiscoveryAssociationsResponse {
	s.Headers = v
	return s
}

func (s *ListIpamResourceDiscoveryAssociationsResponse) SetStatusCode(v int32) *ListIpamResourceDiscoveryAssociationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIpamResourceDiscoveryAssociationsResponse) SetBody(v *ListIpamResourceDiscoveryAssociationsResponseBody) *ListIpamResourceDiscoveryAssociationsResponse {
	s.Body = v
	return s
}

type ListIpamScopesRequest struct {
	// The ID of the IPAM.
	//
	// example:
	//
	// ipam-ccxbnsbhew0d6t****
	IpamId *string `json:"IpamId,omitempty" xml:"IpamId,omitempty"`
	// The IDs of IPAM scopes.
	IpamScopeIds []*string `json:"IpamScopeIds,omitempty" xml:"IpamScopeIds,omitempty" type:"Repeated"`
	// The name of the IPAM scope.
	//
	// It must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	IpamScopeName *string `json:"IpamScopeName,omitempty" xml:"IpamScopeName,omitempty"`
	// The type of the IPAM scope. Valid values:
	//
	// 	- **public**
	//
	// 	- **private**
	//
	// example:
	//
	// private
	IpamScopeType *string `json:"IpamScopeType,omitempty" xml:"IpamScopeType,omitempty"`
	// The number of entries per page. Valid values: **1*	- to **100**. Default value: **10**.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- You do not need to specify this parameter for the first request.
	//
	// 	- You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the IPAM instance is hosted. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID of the IPAM scope.
	//
	// example:
	//
	// rg-aek2sermdd6****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tag list.
	Tags []*ListIpamScopesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListIpamScopesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIpamScopesRequest) GoString() string {
	return s.String()
}

func (s *ListIpamScopesRequest) SetIpamId(v string) *ListIpamScopesRequest {
	s.IpamId = &v
	return s
}

func (s *ListIpamScopesRequest) SetIpamScopeIds(v []*string) *ListIpamScopesRequest {
	s.IpamScopeIds = v
	return s
}

func (s *ListIpamScopesRequest) SetIpamScopeName(v string) *ListIpamScopesRequest {
	s.IpamScopeName = &v
	return s
}

func (s *ListIpamScopesRequest) SetIpamScopeType(v string) *ListIpamScopesRequest {
	s.IpamScopeType = &v
	return s
}

func (s *ListIpamScopesRequest) SetMaxResults(v int64) *ListIpamScopesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListIpamScopesRequest) SetNextToken(v string) *ListIpamScopesRequest {
	s.NextToken = &v
	return s
}

func (s *ListIpamScopesRequest) SetOwnerAccount(v string) *ListIpamScopesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListIpamScopesRequest) SetOwnerId(v int64) *ListIpamScopesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListIpamScopesRequest) SetRegionId(v string) *ListIpamScopesRequest {
	s.RegionId = &v
	return s
}

func (s *ListIpamScopesRequest) SetResourceGroupId(v string) *ListIpamScopesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListIpamScopesRequest) SetResourceOwnerAccount(v string) *ListIpamScopesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListIpamScopesRequest) SetResourceOwnerId(v int64) *ListIpamScopesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListIpamScopesRequest) SetTags(v []*ListIpamScopesRequestTags) *ListIpamScopesRequest {
	s.Tags = v
	return s
}

type ListIpamScopesRequestTags struct {
	// The tag key. You can specify at most 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The tag key must start with a letter but cannot start with `aliyun` or `acs:`. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It cannot start with a `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListIpamScopesRequestTags) String() string {
	return tea.Prettify(s)
}

func (s ListIpamScopesRequestTags) GoString() string {
	return s.String()
}

func (s *ListIpamScopesRequestTags) SetKey(v string) *ListIpamScopesRequestTags {
	s.Key = &v
	return s
}

func (s *ListIpamScopesRequestTags) SetValue(v string) *ListIpamScopesRequestTags {
	s.Value = &v
	return s
}

type ListIpamScopesResponseBody struct {
	// The number of entries returned.
	//
	// example:
	//
	// 10
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The IPAM scopes.
	IpamScopes []*ListIpamScopesResponseBodyIpamScopes `json:"IpamScopes,omitempty" xml:"IpamScopes,omitempty" type:"Repeated"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8859C501-97E7-53D4-B94B-2A9E16003B22
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1000
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListIpamScopesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIpamScopesResponseBody) GoString() string {
	return s.String()
}

func (s *ListIpamScopesResponseBody) SetCount(v int64) *ListIpamScopesResponseBody {
	s.Count = &v
	return s
}

func (s *ListIpamScopesResponseBody) SetIpamScopes(v []*ListIpamScopesResponseBodyIpamScopes) *ListIpamScopesResponseBody {
	s.IpamScopes = v
	return s
}

func (s *ListIpamScopesResponseBody) SetMaxResults(v int64) *ListIpamScopesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListIpamScopesResponseBody) SetNextToken(v string) *ListIpamScopesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListIpamScopesResponseBody) SetRequestId(v string) *ListIpamScopesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIpamScopesResponseBody) SetTotalCount(v int64) *ListIpamScopesResponseBody {
	s.TotalCount = &v
	return s
}

type ListIpamScopesResponseBodyIpamScopes struct {
	// The time when the IPAM scope was created.
	//
	// example:
	//
	// 2022-04-18T03:12:37Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the IPAM.
	//
	// example:
	//
	// ipam-ccxbnsbhew0d6t****
	IpamId *string `json:"IpamId,omitempty" xml:"IpamId,omitempty"`
	// The description of the IPAM scope.
	//
	// example:
	//
	// test description
	IpamScopeDescription *string `json:"IpamScopeDescription,omitempty" xml:"IpamScopeDescription,omitempty"`
	// The ID of the IPAM scope.
	//
	// example:
	//
	// ipam-scope-glfmcyldpm8lsy****
	IpamScopeId *string `json:"IpamScopeId,omitempty" xml:"IpamScopeId,omitempty"`
	// The name of the IPAM scope.
	//
	// example:
	//
	// test
	IpamScopeName *string `json:"IpamScopeName,omitempty" xml:"IpamScopeName,omitempty"`
	// The type of the IPAM scope. Valid values:
	//
	// 	- **public**
	//
	// 	- **private**
	//
	// example:
	//
	// private
	IpamScopeType *string `json:"IpamScopeType,omitempty" xml:"IpamScopeType,omitempty"`
	// Indicates whether the scope is the default scope. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The Alibaba Cloud account that owns the IPAM scope.
	//
	// example:
	//
	// 1210123456******
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of pools in the IPAM scope.
	//
	// example:
	//
	// 2
	PoolCount *int32 `json:"PoolCount,omitempty" xml:"PoolCount,omitempty"`
	// The region ID of the IPAM.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxazb4ph6aiy****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the IPAM scope. Valid values:
	//
	// 	- **Creating**
	//
	// 	- **Created**
	//
	// 	- **Deleting**
	//
	// 	- **Deleted**
	//
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag list.
	Tags []*ListIpamScopesResponseBodyIpamScopesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListIpamScopesResponseBodyIpamScopes) String() string {
	return tea.Prettify(s)
}

func (s ListIpamScopesResponseBodyIpamScopes) GoString() string {
	return s.String()
}

func (s *ListIpamScopesResponseBodyIpamScopes) SetCreateTime(v string) *ListIpamScopesResponseBodyIpamScopes {
	s.CreateTime = &v
	return s
}

func (s *ListIpamScopesResponseBodyIpamScopes) SetIpamId(v string) *ListIpamScopesResponseBodyIpamScopes {
	s.IpamId = &v
	return s
}

func (s *ListIpamScopesResponseBodyIpamScopes) SetIpamScopeDescription(v string) *ListIpamScopesResponseBodyIpamScopes {
	s.IpamScopeDescription = &v
	return s
}

func (s *ListIpamScopesResponseBodyIpamScopes) SetIpamScopeId(v string) *ListIpamScopesResponseBodyIpamScopes {
	s.IpamScopeId = &v
	return s
}

func (s *ListIpamScopesResponseBodyIpamScopes) SetIpamScopeName(v string) *ListIpamScopesResponseBodyIpamScopes {
	s.IpamScopeName = &v
	return s
}

func (s *ListIpamScopesResponseBodyIpamScopes) SetIpamScopeType(v string) *ListIpamScopesResponseBodyIpamScopes {
	s.IpamScopeType = &v
	return s
}

func (s *ListIpamScopesResponseBodyIpamScopes) SetIsDefault(v bool) *ListIpamScopesResponseBodyIpamScopes {
	s.IsDefault = &v
	return s
}

func (s *ListIpamScopesResponseBodyIpamScopes) SetOwnerId(v int64) *ListIpamScopesResponseBodyIpamScopes {
	s.OwnerId = &v
	return s
}

func (s *ListIpamScopesResponseBodyIpamScopes) SetPoolCount(v int32) *ListIpamScopesResponseBodyIpamScopes {
	s.PoolCount = &v
	return s
}

func (s *ListIpamScopesResponseBodyIpamScopes) SetRegionId(v string) *ListIpamScopesResponseBodyIpamScopes {
	s.RegionId = &v
	return s
}

func (s *ListIpamScopesResponseBodyIpamScopes) SetResourceGroupId(v string) *ListIpamScopesResponseBodyIpamScopes {
	s.ResourceGroupId = &v
	return s
}

func (s *ListIpamScopesResponseBodyIpamScopes) SetStatus(v string) *ListIpamScopesResponseBodyIpamScopes {
	s.Status = &v
	return s
}

func (s *ListIpamScopesResponseBodyIpamScopes) SetTags(v []*ListIpamScopesResponseBodyIpamScopesTags) *ListIpamScopesResponseBodyIpamScopes {
	s.Tags = v
	return s
}

type ListIpamScopesResponseBodyIpamScopesTags struct {
	// The tag key.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// FinanceDept
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListIpamScopesResponseBodyIpamScopesTags) String() string {
	return tea.Prettify(s)
}

func (s ListIpamScopesResponseBodyIpamScopesTags) GoString() string {
	return s.String()
}

func (s *ListIpamScopesResponseBodyIpamScopesTags) SetKey(v string) *ListIpamScopesResponseBodyIpamScopesTags {
	s.Key = &v
	return s
}

func (s *ListIpamScopesResponseBodyIpamScopesTags) SetValue(v string) *ListIpamScopesResponseBodyIpamScopesTags {
	s.Value = &v
	return s
}

type ListIpamScopesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIpamScopesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIpamScopesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIpamScopesResponse) GoString() string {
	return s.String()
}

func (s *ListIpamScopesResponse) SetHeaders(v map[string]*string) *ListIpamScopesResponse {
	s.Headers = v
	return s
}

func (s *ListIpamScopesResponse) SetStatusCode(v int32) *ListIpamScopesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIpamScopesResponse) SetBody(v *ListIpamScopesResponseBody) *ListIpamScopesResponse {
	s.Body = v
	return s
}

type ListIpamsRequest struct {
	// The IDs of IPAMs. Valid values of N: 1 to 100. A maximum of 100 IPAMs can be queried at a time.
	IpamIds []*string `json:"IpamIds,omitempty" xml:"IpamIds,omitempty" type:"Repeated"`
	// The name of the IPAM.
	//
	// It must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	IpamName *string `json:"IpamName,omitempty" xml:"IpamName,omitempty"`
	// The number of entries per page. Valid values: **1*	- to **100**. Default value: **10**.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- You do not need to specify this parameter for the first request.
	//
	// 	- You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the IPAM instance is hosted. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID of the IPAM.
	//
	// example:
	//
	// rg-aek2sermdd6****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tag information.
	Tags []*ListIpamsRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListIpamsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIpamsRequest) GoString() string {
	return s.String()
}

func (s *ListIpamsRequest) SetIpamIds(v []*string) *ListIpamsRequest {
	s.IpamIds = v
	return s
}

func (s *ListIpamsRequest) SetIpamName(v string) *ListIpamsRequest {
	s.IpamName = &v
	return s
}

func (s *ListIpamsRequest) SetMaxResults(v int64) *ListIpamsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListIpamsRequest) SetNextToken(v string) *ListIpamsRequest {
	s.NextToken = &v
	return s
}

func (s *ListIpamsRequest) SetOwnerAccount(v string) *ListIpamsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListIpamsRequest) SetOwnerId(v int64) *ListIpamsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListIpamsRequest) SetRegionId(v string) *ListIpamsRequest {
	s.RegionId = &v
	return s
}

func (s *ListIpamsRequest) SetResourceGroupId(v string) *ListIpamsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListIpamsRequest) SetResourceOwnerAccount(v string) *ListIpamsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListIpamsRequest) SetResourceOwnerId(v int64) *ListIpamsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListIpamsRequest) SetTags(v []*ListIpamsRequestTags) *ListIpamsRequest {
	s.Tags = v
	return s
}

type ListIpamsRequestTags struct {
	// The tag key. You can specify at most 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The tag key must start with a letter but cannot start with `aliyun` or `acs:`. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can specify at most 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length. It must start with a letter and can contain digits, periods (.), underscores (_), and hyphens (-). It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListIpamsRequestTags) String() string {
	return tea.Prettify(s)
}

func (s ListIpamsRequestTags) GoString() string {
	return s.String()
}

func (s *ListIpamsRequestTags) SetKey(v string) *ListIpamsRequestTags {
	s.Key = &v
	return s
}

func (s *ListIpamsRequestTags) SetValue(v string) *ListIpamsRequestTags {
	s.Value = &v
	return s
}

type ListIpamsResponseBody struct {
	// The number of entries returned.
	//
	// example:
	//
	// 10
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The IPAMs.
	Ipams []*ListIpamsResponseBodyIpams `json:"Ipams,omitempty" xml:"Ipams,omitempty" type:"Repeated"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 23CA0A0B-B0F5-5495-B355-7D9A9203A46B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListIpamsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIpamsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIpamsResponseBody) SetCount(v int64) *ListIpamsResponseBody {
	s.Count = &v
	return s
}

func (s *ListIpamsResponseBody) SetIpams(v []*ListIpamsResponseBodyIpams) *ListIpamsResponseBody {
	s.Ipams = v
	return s
}

func (s *ListIpamsResponseBody) SetMaxResults(v int64) *ListIpamsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListIpamsResponseBody) SetNextToken(v string) *ListIpamsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListIpamsResponseBody) SetRequestId(v string) *ListIpamsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIpamsResponseBody) SetTotalCount(v int64) *ListIpamsResponseBody {
	s.TotalCount = &v
	return s
}

type ListIpamsResponseBodyIpams struct {
	// The time when the IPAM was created.
	//
	// example:
	//
	// 2022-07-01T02:05:23Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Default resource discovery association ID.
	//
	// example:
	//
	// ipam-res-disco-assoc-jt5fac8twugdbbgip****
	DefaultResourceDiscoveryAssociationId *string `json:"DefaultResourceDiscoveryAssociationId,omitempty" xml:"DefaultResourceDiscoveryAssociationId,omitempty"`
	// Default resource discovery instance ID.
	//
	// example:
	//
	// ipam-res-disco-jt5f2af2u6nk2z321****
	DefaultResourceDiscoveryId *string `json:"DefaultResourceDiscoveryId,omitempty" xml:"DefaultResourceDiscoveryId,omitempty"`
	// The description of the IPAM.
	//
	// example:
	//
	// test description
	IpamDescription *string `json:"IpamDescription,omitempty" xml:"IpamDescription,omitempty"`
	// The ID of the IPAM.
	//
	// example:
	//
	// ipam-ccxbnsbhew0d6t****
	IpamId *string `json:"IpamId,omitempty" xml:"IpamId,omitempty"`
	// The name of the IPAM.
	//
	// example:
	//
	// test
	IpamName *string `json:"IpamName,omitempty" xml:"IpamName,omitempty"`
	// The status of the IPAM. Valid values:
	//
	// 	- **Creating**
	//
	// 	- **Created**
	//
	// 	- **Deleting**
	//
	// 	- **Deleted**
	//
	// example:
	//
	// Created
	IpamStatus *string `json:"IpamStatus,omitempty" xml:"IpamStatus,omitempty"`
	// The effective regions of the IPAM.
	OperatingRegionList []*string `json:"OperatingRegionList,omitempty" xml:"OperatingRegionList,omitempty" type:"Repeated"`
	// The Alibaba Cloud account that owns the IPAM.
	//
	// example:
	//
	// 1210123456******
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The default private scope created by the system after the IPAM is created.
	//
	// example:
	//
	// ipam-scope-okoerbco6unqfr****
	PrivateDefaultScopeId *string `json:"PrivateDefaultScopeId,omitempty" xml:"PrivateDefaultScopeId,omitempty"`
	// The default public scope created by the system after the IPAM is created.
	//
	// example:
	//
	// ipam-scope-ovb76p1g1m19dr****
	PublicDefaultScopeId *string `json:"PublicDefaultScopeId,omitempty" xml:"PublicDefaultScopeId,omitempty"`
	// The region ID of the IPAM.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Number of resource discovery associations.
	//
	// example:
	//
	// 1
	ResourceDiscoveryAssociationCount *int32 `json:"ResourceDiscoveryAssociationCount,omitempty" xml:"ResourceDiscoveryAssociationCount,omitempty"`
	// The resource group ID of the IPAM.
	//
	// example:
	//
	// rg-aek2dbprgpt****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The number of IPAM scopes. Value: **2 to 5**.
	//
	// example:
	//
	// 2
	ScopeCount *int32 `json:"ScopeCount,omitempty" xml:"ScopeCount,omitempty"`
	// The tag list.
	Tags []*ListIpamsResponseBodyIpamsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListIpamsResponseBodyIpams) String() string {
	return tea.Prettify(s)
}

func (s ListIpamsResponseBodyIpams) GoString() string {
	return s.String()
}

func (s *ListIpamsResponseBodyIpams) SetCreateTime(v string) *ListIpamsResponseBodyIpams {
	s.CreateTime = &v
	return s
}

func (s *ListIpamsResponseBodyIpams) SetDefaultResourceDiscoveryAssociationId(v string) *ListIpamsResponseBodyIpams {
	s.DefaultResourceDiscoveryAssociationId = &v
	return s
}

func (s *ListIpamsResponseBodyIpams) SetDefaultResourceDiscoveryId(v string) *ListIpamsResponseBodyIpams {
	s.DefaultResourceDiscoveryId = &v
	return s
}

func (s *ListIpamsResponseBodyIpams) SetIpamDescription(v string) *ListIpamsResponseBodyIpams {
	s.IpamDescription = &v
	return s
}

func (s *ListIpamsResponseBodyIpams) SetIpamId(v string) *ListIpamsResponseBodyIpams {
	s.IpamId = &v
	return s
}

func (s *ListIpamsResponseBodyIpams) SetIpamName(v string) *ListIpamsResponseBodyIpams {
	s.IpamName = &v
	return s
}

func (s *ListIpamsResponseBodyIpams) SetIpamStatus(v string) *ListIpamsResponseBodyIpams {
	s.IpamStatus = &v
	return s
}

func (s *ListIpamsResponseBodyIpams) SetOperatingRegionList(v []*string) *ListIpamsResponseBodyIpams {
	s.OperatingRegionList = v
	return s
}

func (s *ListIpamsResponseBodyIpams) SetOwnerId(v int64) *ListIpamsResponseBodyIpams {
	s.OwnerId = &v
	return s
}

func (s *ListIpamsResponseBodyIpams) SetPrivateDefaultScopeId(v string) *ListIpamsResponseBodyIpams {
	s.PrivateDefaultScopeId = &v
	return s
}

func (s *ListIpamsResponseBodyIpams) SetPublicDefaultScopeId(v string) *ListIpamsResponseBodyIpams {
	s.PublicDefaultScopeId = &v
	return s
}

func (s *ListIpamsResponseBodyIpams) SetRegionId(v string) *ListIpamsResponseBodyIpams {
	s.RegionId = &v
	return s
}

func (s *ListIpamsResponseBodyIpams) SetResourceDiscoveryAssociationCount(v int32) *ListIpamsResponseBodyIpams {
	s.ResourceDiscoveryAssociationCount = &v
	return s
}

func (s *ListIpamsResponseBodyIpams) SetResourceGroupId(v string) *ListIpamsResponseBodyIpams {
	s.ResourceGroupId = &v
	return s
}

func (s *ListIpamsResponseBodyIpams) SetScopeCount(v int32) *ListIpamsResponseBodyIpams {
	s.ScopeCount = &v
	return s
}

func (s *ListIpamsResponseBodyIpams) SetTags(v []*ListIpamsResponseBodyIpamsTags) *ListIpamsResponseBodyIpams {
	s.Tags = v
	return s
}

type ListIpamsResponseBodyIpamsTags struct {
	// The tag key.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListIpamsResponseBodyIpamsTags) String() string {
	return tea.Prettify(s)
}

func (s ListIpamsResponseBodyIpamsTags) GoString() string {
	return s.String()
}

func (s *ListIpamsResponseBodyIpamsTags) SetKey(v string) *ListIpamsResponseBodyIpamsTags {
	s.Key = &v
	return s
}

func (s *ListIpamsResponseBodyIpamsTags) SetValue(v string) *ListIpamsResponseBodyIpamsTags {
	s.Value = &v
	return s
}

type ListIpamsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIpamsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIpamsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIpamsResponse) GoString() string {
	return s.String()
}

func (s *ListIpamsResponse) SetHeaders(v map[string]*string) *ListIpamsResponse {
	s.Headers = v
	return s
}

func (s *ListIpamsResponse) SetStatusCode(v int32) *ListIpamsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIpamsResponse) SetBody(v *ListIpamsResponseBody) *ListIpamsResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// The number of entries per page. Valid values: **1*	- to **50**. Default value: **10**.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- You do not need to specify this parameter for the first request.
	//
	// 	- If a value is returned for NextToken, you must specify the token that is obtained from the previous query as the value of **NextToken**.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource IDs.
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The resource type. Valid values:
	//
	// 	- **IPAM**
	//
	// 	- **IPAMSCOPE**
	//
	// 	- **IPAMPOOL**
	//
	// This parameter is required.
	//
	// example:
	//
	// IPAM
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag list.
	Tag []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetMaxResults(v int32) *ListTagResourcesRequest {
	s.MaxResults = &v
	return s
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

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	// The tag key. You can specify at most 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It cannot start with a `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// >  You must specify **ResourceId.N*	- or **Tag.N*	- that consists of **Tag.N.Key*	- and **Tag.N.Value**.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can specify at most 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It cannot start with a `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// >  You must specify **ResourceId.N*	- or **Tag.N*	- that consists of **Tag.N.Key*	- and **Tag.N.Value**.
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
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 672053AB-90C9-5693-AB96-458F137A5ED6
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
	// ipam-uq5dcfc2eqhpf4****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type. Valid values:
	//
	// 	- **IPAM**
	//
	// 	- **IPAMSCOPE**
	//
	// 	- **IPAMPOOL**
	//
	// example:
	//
	// IPAM
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag key.
	//
	// example:
	//
	// FinanceDept
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
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

type OpenVpcIpamServiceRequest struct {
	// Client token, used to ensure the idempotence of requests.
	//
	// Generate a unique value for this parameter from your client to ensure it is unique across different requests. ClientToken supports only ASCII characters.
	//
	// > If not specified, the system automatically uses the RequestId of the API request as the ClientToken identifier. The RequestId may differ for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the IPAM instance is hosted.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s OpenVpcIpamServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenVpcIpamServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenVpcIpamServiceRequest) SetClientToken(v string) *OpenVpcIpamServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *OpenVpcIpamServiceRequest) SetOwnerAccount(v string) *OpenVpcIpamServiceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *OpenVpcIpamServiceRequest) SetOwnerId(v int64) *OpenVpcIpamServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *OpenVpcIpamServiceRequest) SetRegionId(v string) *OpenVpcIpamServiceRequest {
	s.RegionId = &v
	return s
}

func (s *OpenVpcIpamServiceRequest) SetResourceOwnerAccount(v string) *OpenVpcIpamServiceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *OpenVpcIpamServiceRequest) SetResourceOwnerId(v int64) *OpenVpcIpamServiceRequest {
	s.ResourceOwnerId = &v
	return s
}

type OpenVpcIpamServiceResponseBody struct {
	// Status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Information returned upon successful IPAM activation.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 3F814C37-B032-5477-AF5A-2925D0593CD4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenVpcIpamServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenVpcIpamServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenVpcIpamServiceResponseBody) SetCode(v string) *OpenVpcIpamServiceResponseBody {
	s.Code = &v
	return s
}

func (s *OpenVpcIpamServiceResponseBody) SetMessage(v string) *OpenVpcIpamServiceResponseBody {
	s.Message = &v
	return s
}

func (s *OpenVpcIpamServiceResponseBody) SetRequestId(v string) *OpenVpcIpamServiceResponseBody {
	s.RequestId = &v
	return s
}

type OpenVpcIpamServiceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenVpcIpamServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenVpcIpamServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenVpcIpamServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenVpcIpamServiceResponse) SetHeaders(v map[string]*string) *OpenVpcIpamServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenVpcIpamServiceResponse) SetStatusCode(v int32) *OpenVpcIpamServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenVpcIpamServiceResponse) SetBody(v *OpenVpcIpamServiceResponseBody) *OpenVpcIpamServiceResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the resources.
	//
	// This parameter is required.
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The resource type. Valid values:
	//
	// 	- **IPAM**
	//
	// 	- **IPAMSCOPE**
	//
	// 	- **IPAMPOOL**
	//
	// This parameter is required.
	//
	// example:
	//
	// IPAM
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

func (s *TagResourcesRequest) SetOwnerAccount(v string) *TagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *TagResourcesRequest) SetOwnerId(v int64) *TagResourcesRequest {
	s.OwnerId = &v
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
	// The tag key. You can specify at most 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It cannot start with a `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can specify at most 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It cannot start with a `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
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
	// BF872550-9700-52FD-839C-4D3F05543FA8
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
	// Specifies whether to remove all tags from the specified resource. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	All          *bool   `json:"All,omitempty" xml:"All,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource IDs.
	//
	// This parameter is required.
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The resource type. Valid values:
	//
	// 	- **IPAM**
	//
	// 	- **IPAMSCOPE**
	//
	// 	- **IPAMPOOL**
	//
	// This parameter is required.
	//
	// example:
	//
	// IPAM
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The keys of the tags that you want to remove from the resource.
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

func (s *UntagResourcesRequest) SetOwnerAccount(v string) *UntagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UntagResourcesRequest) SetOwnerId(v int64) *UntagResourcesRequest {
	s.OwnerId = &v
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
	// The request ID.
	//
	// example:
	//
	// 44C884BD-2D56-5637-A523-1FA920A01E7D
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

type UpdateIpamRequest struct {
	// The effective region that you want to add.
	AddOperatingRegion []*string `json:"AddOperatingRegion,omitempty" xml:"AddOperatingRegion,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The description of the IPAM.
	//
	// It must be 2 to 256 characters in length and must start with a letter. It cannot start with `http://` or `https://`. The default value is empty.
	//
	// example:
	//
	// test description
	IpamDescription *string `json:"IpamDescription,omitempty" xml:"IpamDescription,omitempty"`
	// The ID of the IPAM.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-ccxbnsbhew0d6t****
	IpamId *string `json:"IpamId,omitempty" xml:"IpamId,omitempty"`
	// The name of the IPAM.
	//
	// It must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	IpamName     *string `json:"IpamName,omitempty" xml:"IpamName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the IPAM instance is hosted. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The effective region that you want to remove.
	RemoveOperatingRegion []*string `json:"RemoveOperatingRegion,omitempty" xml:"RemoveOperatingRegion,omitempty" type:"Repeated"`
	ResourceOwnerAccount  *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpdateIpamRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateIpamRequest) GoString() string {
	return s.String()
}

func (s *UpdateIpamRequest) SetAddOperatingRegion(v []*string) *UpdateIpamRequest {
	s.AddOperatingRegion = v
	return s
}

func (s *UpdateIpamRequest) SetClientToken(v string) *UpdateIpamRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateIpamRequest) SetDryRun(v bool) *UpdateIpamRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateIpamRequest) SetIpamDescription(v string) *UpdateIpamRequest {
	s.IpamDescription = &v
	return s
}

func (s *UpdateIpamRequest) SetIpamId(v string) *UpdateIpamRequest {
	s.IpamId = &v
	return s
}

func (s *UpdateIpamRequest) SetIpamName(v string) *UpdateIpamRequest {
	s.IpamName = &v
	return s
}

func (s *UpdateIpamRequest) SetOwnerAccount(v string) *UpdateIpamRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateIpamRequest) SetOwnerId(v int64) *UpdateIpamRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateIpamRequest) SetRegionId(v string) *UpdateIpamRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateIpamRequest) SetRemoveOperatingRegion(v []*string) *UpdateIpamRequest {
	s.RemoveOperatingRegion = v
	return s
}

func (s *UpdateIpamRequest) SetResourceOwnerAccount(v string) *UpdateIpamRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateIpamRequest) SetResourceOwnerId(v int64) *UpdateIpamRequest {
	s.ResourceOwnerId = &v
	return s
}

type UpdateIpamResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F4650E33-895C-53F0-9CD5-D1338F322DE8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateIpamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateIpamResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIpamResponseBody) SetRequestId(v string) *UpdateIpamResponseBody {
	s.RequestId = &v
	return s
}

type UpdateIpamResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateIpamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateIpamResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateIpamResponse) GoString() string {
	return s.String()
}

func (s *UpdateIpamResponse) SetHeaders(v map[string]*string) *UpdateIpamResponse {
	s.Headers = v
	return s
}

func (s *UpdateIpamResponse) SetStatusCode(v int32) *UpdateIpamResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIpamResponse) SetBody(v *UpdateIpamResponseBody) *UpdateIpamResponse {
	s.Body = v
	return s
}

type UpdateIpamPoolRequest struct {
	// The new default network mask for the IPAM pool.
	//
	// The mask must be **0 to 32*	- bits in length.
	//
	// example:
	//
	// 28
	AllocationDefaultCidrMask *int32 `json:"AllocationDefaultCidrMask,omitempty" xml:"AllocationDefaultCidrMask,omitempty"`
	// The new maximum network mask for the IPAM pool.
	//
	// The mask must be **0 to 32*	- bits in length.
	//
	// example:
	//
	// 32
	AllocationMaxCidrMask *int32 `json:"AllocationMaxCidrMask,omitempty" xml:"AllocationMaxCidrMask,omitempty"`
	// The new minimum network mask for the IPAM pool.
	//
	// The mask must be **0 to 32*	- bits in length.
	//
	// example:
	//
	// 8
	AllocationMinCidrMask *int32 `json:"AllocationMinCidrMask,omitempty" xml:"AllocationMinCidrMask,omitempty"`
	// Whether the pool has the auto-import feature enabled.
	//
	// example:
	//
	// true
	AutoImport *bool `json:"AutoImport,omitempty" xml:"AutoImport,omitempty"`
	// Specifies whether to delete the default network mask for the IPAM pool. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	ClearAllocationDefaultCidrMask *bool `json:"ClearAllocationDefaultCidrMask,omitempty" xml:"ClearAllocationDefaultCidrMask,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The new description of the IPAM pool.
	//
	// It must be 2 to 268 characters in length. It must start with a letter but cannot start with a `http://` or `https://`. This parameter is empty by default.
	//
	// example:
	//
	// test description
	IpamPoolDescription *string `json:"IpamPoolDescription,omitempty" xml:"IpamPoolDescription,omitempty"`
	// The ID of the IPAM pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-pool-6rcq3tobayc20t****
	IpamPoolId *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
	// The new name of the IPAM pool.
	//
	// It must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	IpamPoolName *string `json:"IpamPoolName,omitempty" xml:"IpamPoolName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the IPAM instance is hosted. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpdateIpamPoolRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateIpamPoolRequest) GoString() string {
	return s.String()
}

func (s *UpdateIpamPoolRequest) SetAllocationDefaultCidrMask(v int32) *UpdateIpamPoolRequest {
	s.AllocationDefaultCidrMask = &v
	return s
}

func (s *UpdateIpamPoolRequest) SetAllocationMaxCidrMask(v int32) *UpdateIpamPoolRequest {
	s.AllocationMaxCidrMask = &v
	return s
}

func (s *UpdateIpamPoolRequest) SetAllocationMinCidrMask(v int32) *UpdateIpamPoolRequest {
	s.AllocationMinCidrMask = &v
	return s
}

func (s *UpdateIpamPoolRequest) SetAutoImport(v bool) *UpdateIpamPoolRequest {
	s.AutoImport = &v
	return s
}

func (s *UpdateIpamPoolRequest) SetClearAllocationDefaultCidrMask(v bool) *UpdateIpamPoolRequest {
	s.ClearAllocationDefaultCidrMask = &v
	return s
}

func (s *UpdateIpamPoolRequest) SetClientToken(v string) *UpdateIpamPoolRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateIpamPoolRequest) SetDryRun(v bool) *UpdateIpamPoolRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateIpamPoolRequest) SetIpamPoolDescription(v string) *UpdateIpamPoolRequest {
	s.IpamPoolDescription = &v
	return s
}

func (s *UpdateIpamPoolRequest) SetIpamPoolId(v string) *UpdateIpamPoolRequest {
	s.IpamPoolId = &v
	return s
}

func (s *UpdateIpamPoolRequest) SetIpamPoolName(v string) *UpdateIpamPoolRequest {
	s.IpamPoolName = &v
	return s
}

func (s *UpdateIpamPoolRequest) SetOwnerAccount(v string) *UpdateIpamPoolRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateIpamPoolRequest) SetOwnerId(v int64) *UpdateIpamPoolRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateIpamPoolRequest) SetRegionId(v string) *UpdateIpamPoolRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateIpamPoolRequest) SetResourceOwnerAccount(v string) *UpdateIpamPoolRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateIpamPoolRequest) SetResourceOwnerId(v int64) *UpdateIpamPoolRequest {
	s.ResourceOwnerId = &v
	return s
}

type UpdateIpamPoolResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9DED57B9-7654-5B6D-93F7-BCA5839FEE38
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateIpamPoolResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateIpamPoolResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIpamPoolResponseBody) SetRequestId(v string) *UpdateIpamPoolResponseBody {
	s.RequestId = &v
	return s
}

type UpdateIpamPoolResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateIpamPoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateIpamPoolResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateIpamPoolResponse) GoString() string {
	return s.String()
}

func (s *UpdateIpamPoolResponse) SetHeaders(v map[string]*string) *UpdateIpamPoolResponse {
	s.Headers = v
	return s
}

func (s *UpdateIpamPoolResponse) SetStatusCode(v int32) *UpdateIpamPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIpamPoolResponse) SetBody(v *UpdateIpamPoolResponseBody) *UpdateIpamPoolResponse {
	s.Body = v
	return s
}

type UpdateIpamPoolAllocationRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to precheck this request. Valid values:
	//
	// 	- **true**: performs a dry run without modifying the CIDR blocks of IPAM pools. The system checks the required parameters, request syntax, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and the actual request. If the request passes the check, an HTTP 2xx status code is returned and the CIDR allocation information of the IPAM address pool is modified.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The description of the CIDR allocation of the IPAM pool.
	//
	// The description must be 1 to 256 characters in length and start with a letter, but cannot start with a `http://` or `https://`. This parameter is empty by default.
	//
	// example:
	//
	// test description
	IpamPoolAllocationDescription *string `json:"IpamPoolAllocationDescription,omitempty" xml:"IpamPoolAllocationDescription,omitempty"`
	// The ID of the instance to which CIDR blocks are allocated from the IPAM pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-pool-alloc-112za33e4****
	IpamPoolAllocationId *string `json:"IpamPoolAllocationId,omitempty" xml:"IpamPoolAllocationId,omitempty"`
	// The name of the CIDR allocation of the IPAM pool.
	//
	// It must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test name
	IpamPoolAllocationName *string `json:"IpamPoolAllocationName,omitempty" xml:"IpamPoolAllocationName,omitempty"`
	// The ID of the region where you want to perform the operation. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateIpamPoolAllocationRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateIpamPoolAllocationRequest) GoString() string {
	return s.String()
}

func (s *UpdateIpamPoolAllocationRequest) SetClientToken(v string) *UpdateIpamPoolAllocationRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateIpamPoolAllocationRequest) SetDryRun(v bool) *UpdateIpamPoolAllocationRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateIpamPoolAllocationRequest) SetIpamPoolAllocationDescription(v string) *UpdateIpamPoolAllocationRequest {
	s.IpamPoolAllocationDescription = &v
	return s
}

func (s *UpdateIpamPoolAllocationRequest) SetIpamPoolAllocationId(v string) *UpdateIpamPoolAllocationRequest {
	s.IpamPoolAllocationId = &v
	return s
}

func (s *UpdateIpamPoolAllocationRequest) SetIpamPoolAllocationName(v string) *UpdateIpamPoolAllocationRequest {
	s.IpamPoolAllocationName = &v
	return s
}

func (s *UpdateIpamPoolAllocationRequest) SetRegionId(v string) *UpdateIpamPoolAllocationRequest {
	s.RegionId = &v
	return s
}

type UpdateIpamPoolAllocationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F4650E33-895C-53F0-9CD5-D1338F322DE8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateIpamPoolAllocationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateIpamPoolAllocationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIpamPoolAllocationResponseBody) SetRequestId(v string) *UpdateIpamPoolAllocationResponseBody {
	s.RequestId = &v
	return s
}

type UpdateIpamPoolAllocationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateIpamPoolAllocationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateIpamPoolAllocationResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateIpamPoolAllocationResponse) GoString() string {
	return s.String()
}

func (s *UpdateIpamPoolAllocationResponse) SetHeaders(v map[string]*string) *UpdateIpamPoolAllocationResponse {
	s.Headers = v
	return s
}

func (s *UpdateIpamPoolAllocationResponse) SetStatusCode(v int32) *UpdateIpamPoolAllocationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIpamPoolAllocationResponse) SetBody(v *UpdateIpamPoolAllocationResponseBody) *UpdateIpamPoolAllocationResponse {
	s.Body = v
	return s
}

type UpdateIpamResourceDiscoveryRequest struct {
	// The list of effective regions to add.
	AddOperatingRegion []*string `json:"AddOperatingRegion,omitempty" xml:"AddOperatingRegion,omitempty" type:"Repeated"`
	// The client token used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform the dry run. Valid values:
	//
	// 	- **true**: Performs a dry run without modifying the resource discovery instance. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// 	- **false*	- (default): Performs a dry run and the actual request. If the request passes the check, an HTTP 2xx status code is returned and the resource discovery instance is modified.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The description of resource discovery.
	//
	// example:
	//
	// test description
	IpamResourceDiscoveryDescription *string `json:"IpamResourceDiscoveryDescription,omitempty" xml:"IpamResourceDiscoveryDescription,omitempty"`
	// The ID of resource discovery instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-res-disco-jt5f2af2u6nk2z321****
	IpamResourceDiscoveryId *string `json:"IpamResourceDiscoveryId,omitempty" xml:"IpamResourceDiscoveryId,omitempty"`
	// The name of the resource discovery.
	//
	// example:
	//
	// test
	IpamResourceDiscoveryName *string `json:"IpamResourceDiscoveryName,omitempty" xml:"IpamResourceDiscoveryName,omitempty"`
	OwnerAccount              *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The request region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of effective regions to remove.
	RemoveOperatingRegion []*string `json:"RemoveOperatingRegion,omitempty" xml:"RemoveOperatingRegion,omitempty" type:"Repeated"`
	ResourceOwnerAccount  *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpdateIpamResourceDiscoveryRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateIpamResourceDiscoveryRequest) GoString() string {
	return s.String()
}

func (s *UpdateIpamResourceDiscoveryRequest) SetAddOperatingRegion(v []*string) *UpdateIpamResourceDiscoveryRequest {
	s.AddOperatingRegion = v
	return s
}

func (s *UpdateIpamResourceDiscoveryRequest) SetClientToken(v string) *UpdateIpamResourceDiscoveryRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateIpamResourceDiscoveryRequest) SetDryRun(v bool) *UpdateIpamResourceDiscoveryRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateIpamResourceDiscoveryRequest) SetIpamResourceDiscoveryDescription(v string) *UpdateIpamResourceDiscoveryRequest {
	s.IpamResourceDiscoveryDescription = &v
	return s
}

func (s *UpdateIpamResourceDiscoveryRequest) SetIpamResourceDiscoveryId(v string) *UpdateIpamResourceDiscoveryRequest {
	s.IpamResourceDiscoveryId = &v
	return s
}

func (s *UpdateIpamResourceDiscoveryRequest) SetIpamResourceDiscoveryName(v string) *UpdateIpamResourceDiscoveryRequest {
	s.IpamResourceDiscoveryName = &v
	return s
}

func (s *UpdateIpamResourceDiscoveryRequest) SetOwnerAccount(v string) *UpdateIpamResourceDiscoveryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateIpamResourceDiscoveryRequest) SetOwnerId(v int64) *UpdateIpamResourceDiscoveryRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateIpamResourceDiscoveryRequest) SetRegionId(v string) *UpdateIpamResourceDiscoveryRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateIpamResourceDiscoveryRequest) SetRemoveOperatingRegion(v []*string) *UpdateIpamResourceDiscoveryRequest {
	s.RemoveOperatingRegion = v
	return s
}

func (s *UpdateIpamResourceDiscoveryRequest) SetResourceOwnerAccount(v string) *UpdateIpamResourceDiscoveryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateIpamResourceDiscoveryRequest) SetResourceOwnerId(v int64) *UpdateIpamResourceDiscoveryRequest {
	s.ResourceOwnerId = &v
	return s
}

type UpdateIpamResourceDiscoveryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BB2C39DE-CEB8-595A-981A-F2EFCBE7324E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateIpamResourceDiscoveryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateIpamResourceDiscoveryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIpamResourceDiscoveryResponseBody) SetRequestId(v string) *UpdateIpamResourceDiscoveryResponseBody {
	s.RequestId = &v
	return s
}

type UpdateIpamResourceDiscoveryResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateIpamResourceDiscoveryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateIpamResourceDiscoveryResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateIpamResourceDiscoveryResponse) GoString() string {
	return s.String()
}

func (s *UpdateIpamResourceDiscoveryResponse) SetHeaders(v map[string]*string) *UpdateIpamResourceDiscoveryResponse {
	s.Headers = v
	return s
}

func (s *UpdateIpamResourceDiscoveryResponse) SetStatusCode(v int32) *UpdateIpamResourceDiscoveryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIpamResourceDiscoveryResponse) SetBody(v *UpdateIpamResourceDiscoveryResponseBody) *UpdateIpamResourceDiscoveryResponse {
	s.Body = v
	return s
}

type UpdateIpamScopeRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, DryRunOperation is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The new description of the IPAM scope.
	//
	// It must be 2 to 256 characters in length. It must start with a letter but cannot start with `http://` or `https://`. This parameter is empty by default.
	//
	// example:
	//
	// test description
	IpamScopeDescription *string `json:"IpamScopeDescription,omitempty" xml:"IpamScopeDescription,omitempty"`
	// The ID of the IPAM scope.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-scope-glfmcyldpm8lsy****
	IpamScopeId *string `json:"IpamScopeId,omitempty" xml:"IpamScopeId,omitempty"`
	// The new name of the IPAM scope.
	//
	// It must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	IpamScopeName *string `json:"IpamScopeName,omitempty" xml:"IpamScopeName,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the IPAM instance is hosted. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpdateIpamScopeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateIpamScopeRequest) GoString() string {
	return s.String()
}

func (s *UpdateIpamScopeRequest) SetClientToken(v string) *UpdateIpamScopeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateIpamScopeRequest) SetDryRun(v bool) *UpdateIpamScopeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateIpamScopeRequest) SetIpamScopeDescription(v string) *UpdateIpamScopeRequest {
	s.IpamScopeDescription = &v
	return s
}

func (s *UpdateIpamScopeRequest) SetIpamScopeId(v string) *UpdateIpamScopeRequest {
	s.IpamScopeId = &v
	return s
}

func (s *UpdateIpamScopeRequest) SetIpamScopeName(v string) *UpdateIpamScopeRequest {
	s.IpamScopeName = &v
	return s
}

func (s *UpdateIpamScopeRequest) SetOwnerAccount(v string) *UpdateIpamScopeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateIpamScopeRequest) SetOwnerId(v int64) *UpdateIpamScopeRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateIpamScopeRequest) SetRegionId(v string) *UpdateIpamScopeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateIpamScopeRequest) SetResourceOwnerAccount(v string) *UpdateIpamScopeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateIpamScopeRequest) SetResourceOwnerId(v int64) *UpdateIpamScopeRequest {
	s.ResourceOwnerId = &v
	return s
}

type UpdateIpamScopeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BA8054F5-852A-570A-ACFF-BCA63E0B02D5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateIpamScopeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateIpamScopeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIpamScopeResponseBody) SetRequestId(v string) *UpdateIpamScopeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateIpamScopeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateIpamScopeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateIpamScopeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateIpamScopeResponse) GoString() string {
	return s.String()
}

func (s *UpdateIpamScopeResponse) SetHeaders(v map[string]*string) *UpdateIpamScopeResponse {
	s.Headers = v
	return s
}

func (s *UpdateIpamScopeResponse) SetStatusCode(v int32) *UpdateIpamScopeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIpamScopeResponse) SetBody(v *UpdateIpamScopeResponseBody) *UpdateIpamScopeResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("vpcipam"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Provisions a CIDR block to an IP Address Manager (IPAM) pool.
//
// Description:
//
//	  Before you provision a CIDR block, make sure that an IPAM pool is created. You can call the **CreateIpamPool*	- operation to create an IPAM pool.
//
//		- If no CIDR block is provisioned to a parent pool, you cannot provision CIDR blocks to its subpools.
//
//		- If a CIDR block is provisioned to a parent pool, you can provision CIDR blocks to its subpools and the CIDR blocks must be subsets of the CIDR block provisioned to the parent pool.
//
//		- If a CIDR block is provisioned to a parent pool and allocations are created, CIDR blocks provisioned to its subpools cannot overlap with existing allocated CIDR blocks.
//
//		- You can provision CIDR blocks to a pool only in the region where the IPAM is hosted.
//
//		- CIDR blocks provisioned to an IPAM pool cannot overlap with the CIDR blocks provisioned to other pools in the same scope.
//
//		- You can provision at most 50 CIDR blocks to each pool.
//
// @param request - AddIpamPoolCidrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddIpamPoolCidrResponse
func (client *Client) AddIpamPoolCidrWithOptions(request *AddIpamPoolCidrRequest, runtime *util.RuntimeOptions) (_result *AddIpamPoolCidrResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cidr)) {
		query["Cidr"] = request.Cidr
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.IpamPoolId)) {
		query["IpamPoolId"] = request.IpamPoolId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddIpamPoolCidr"),
		Version:     tea.String("2023-02-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddIpamPoolCidrResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Provisions a CIDR block to an IP Address Manager (IPAM) pool.
//
// Description:
//
//	  Before you provision a CIDR block, make sure that an IPAM pool is created. You can call the **CreateIpamPool*	- operation to create an IPAM pool.
//
//		- If no CIDR block is provisioned to a parent pool, you cannot provision CIDR blocks to its subpools.
//
//		- If a CIDR block is provisioned to a parent pool, you can provision CIDR blocks to its subpools and the CIDR blocks must be subsets of the CIDR block provisioned to the parent pool.
//
//		- If a CIDR block is provisioned to a parent pool and allocations are created, CIDR blocks provisioned to its subpools cannot overlap with existing allocated CIDR blocks.
//
//		- You can provision CIDR blocks to a pool only in the region where the IPAM is hosted.
//
//		- CIDR blocks provisioned to an IPAM pool cannot overlap with the CIDR blocks provisioned to other pools in the same scope.
//
//		- You can provision at most 50 CIDR blocks to each pool.
//
// @param request - AddIpamPoolCidrRequest
//
// @return AddIpamPoolCidrResponse
func (client *Client) AddIpamPoolCidr(request *AddIpamPoolCidrRequest) (_result *AddIpamPoolCidrResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddIpamPoolCidrResponse{}
	_body, _err := client.AddIpamPoolCidrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates resource discovery with an IPAM instance.
//
// Description:
//
//	The specified resource discovery instance can only be associated with one IPAM instance and associations cannot be duplicated.
//
// @param request - AssociateIpamResourceDiscoveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateIpamResourceDiscoveryResponse
func (client *Client) AssociateIpamResourceDiscoveryWithOptions(request *AssociateIpamResourceDiscoveryRequest, runtime *util.RuntimeOptions) (_result *AssociateIpamResourceDiscoveryResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.IpamId)) {
		query["IpamId"] = request.IpamId
	}

	if !tea.BoolValue(util.IsUnset(request.IpamResourceDiscoveryId)) {
		query["IpamResourceDiscoveryId"] = request.IpamResourceDiscoveryId
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AssociateIpamResourceDiscovery"),
		Version:     tea.String("2023-02-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AssociateIpamResourceDiscoveryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates resource discovery with an IPAM instance.
//
// Description:
//
//	The specified resource discovery instance can only be associated with one IPAM instance and associations cannot be duplicated.
//
// @param request - AssociateIpamResourceDiscoveryRequest
//
// @return AssociateIpamResourceDiscoveryResponse
func (client *Client) AssociateIpamResourceDiscovery(request *AssociateIpamResourceDiscoveryRequest) (_result *AssociateIpamResourceDiscoveryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssociateIpamResourceDiscoveryResponse{}
	_body, _err := client.AssociateIpamResourceDiscoveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the resource group of an IPAM resource.
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
	if !tea.BoolValue(util.IsUnset(request.NewResourceGroupId)) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeResourceGroup"),
		Version:     tea.String("2023-02-28"),
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
// Changes the resource group of an IPAM resource.
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
// Creates an IP Address Manager (IPAM).
//
// Description:
//
// - You can create only one IPAM with each Alibaba Cloud account in each region.
//
// - Only IPv4 IP addresses can be allocated.
//
// - When you create an IPAM instance:
//
//   - If there is no custom resource discovery in the region, the system creates a default resource discovery associated with the IPAM instance.
//
//   - If there is a custom resource discovery in the region, the system converts it to a default resource discovery and associates it with the IPAM instance.
//
// @param request - CreateIpamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIpamResponse
func (client *Client) CreateIpamWithOptions(request *CreateIpamRequest, runtime *util.RuntimeOptions) (_result *CreateIpamResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.IpamDescription)) {
		query["IpamDescription"] = request.IpamDescription
	}

	if !tea.BoolValue(util.IsUnset(request.IpamName)) {
		query["IpamName"] = request.IpamName
	}

	if !tea.BoolValue(util.IsUnset(request.OperatingRegionList)) {
		query["OperatingRegionList"] = request.OperatingRegionList
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

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateIpam"),
		Version:     tea.String("2023-02-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateIpamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an IP Address Manager (IPAM).
//
// Description:
//
// - You can create only one IPAM with each Alibaba Cloud account in each region.
//
// - Only IPv4 IP addresses can be allocated.
//
// - When you create an IPAM instance:
//
//   - If there is no custom resource discovery in the region, the system creates a default resource discovery associated with the IPAM instance.
//
//   - If there is a custom resource discovery in the region, the system converts it to a default resource discovery and associates it with the IPAM instance.
//
// @param request - CreateIpamRequest
//
// @return CreateIpamResponse
func (client *Client) CreateIpam(request *CreateIpamRequest) (_result *CreateIpamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateIpamResponse{}
	_body, _err := client.CreateIpamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an IP Address Manager (IPAM) pool.
//
// @param request - CreateIpamPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIpamPoolResponse
func (client *Client) CreateIpamPoolWithOptions(request *CreateIpamPoolRequest, runtime *util.RuntimeOptions) (_result *CreateIpamPoolResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllocationDefaultCidrMask)) {
		query["AllocationDefaultCidrMask"] = request.AllocationDefaultCidrMask
	}

	if !tea.BoolValue(util.IsUnset(request.AllocationMaxCidrMask)) {
		query["AllocationMaxCidrMask"] = request.AllocationMaxCidrMask
	}

	if !tea.BoolValue(util.IsUnset(request.AllocationMinCidrMask)) {
		query["AllocationMinCidrMask"] = request.AllocationMinCidrMask
	}

	if !tea.BoolValue(util.IsUnset(request.AutoImport)) {
		query["AutoImport"] = request.AutoImport
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.IpVersion)) {
		query["IpVersion"] = request.IpVersion
	}

	if !tea.BoolValue(util.IsUnset(request.IpamPoolDescription)) {
		query["IpamPoolDescription"] = request.IpamPoolDescription
	}

	if !tea.BoolValue(util.IsUnset(request.IpamPoolName)) {
		query["IpamPoolName"] = request.IpamPoolName
	}

	if !tea.BoolValue(util.IsUnset(request.IpamScopeId)) {
		query["IpamScopeId"] = request.IpamScopeId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PoolRegionId)) {
		query["PoolRegionId"] = request.PoolRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIpamPoolId)) {
		query["SourceIpamPoolId"] = request.SourceIpamPoolId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateIpamPool"),
		Version:     tea.String("2023-02-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateIpamPoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an IP Address Manager (IPAM) pool.
//
// @param request - CreateIpamPoolRequest
//
// @return CreateIpamPoolResponse
func (client *Client) CreateIpamPool(request *CreateIpamPoolRequest) (_result *CreateIpamPoolResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateIpamPoolResponse{}
	_body, _err := client.CreateIpamPoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Reserves a custom CIDR block from an IP Address Manager (IPAM) pool.
//
// Description:
//
//	  Before you reserve a custom CIDR block, make sure that an IPAM pool is created and CIDR blocks are added to the pool. You can call **CreateIpamPool*	- to create an IPAM pool and call **AddIpamPoolCidr*	- to add CIDR blocks to the pool.
//
//		- When you specify Cidr or CidrMask to reserve a custom CIDR block, the mask must fall within the range specified by the IPAM pool.
//
//		- If the IPAM pool has the region attribute, you must reserve a custom CIDR block in the region to which the IPAM pool belongs.
//
//		- The custom CIDR block that you want to reserve cannot overlap with existing CIDR blocks created from the IPAM pool.
//
// @param request - CreateIpamPoolAllocationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIpamPoolAllocationResponse
func (client *Client) CreateIpamPoolAllocationWithOptions(request *CreateIpamPoolAllocationRequest, runtime *util.RuntimeOptions) (_result *CreateIpamPoolAllocationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cidr)) {
		query["Cidr"] = request.Cidr
	}

	if !tea.BoolValue(util.IsUnset(request.CidrMask)) {
		query["CidrMask"] = request.CidrMask
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.IpamPoolAllocationDescription)) {
		query["IpamPoolAllocationDescription"] = request.IpamPoolAllocationDescription
	}

	if !tea.BoolValue(util.IsUnset(request.IpamPoolAllocationName)) {
		query["IpamPoolAllocationName"] = request.IpamPoolAllocationName
	}

	if !tea.BoolValue(util.IsUnset(request.IpamPoolId)) {
		query["IpamPoolId"] = request.IpamPoolId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateIpamPoolAllocation"),
		Version:     tea.String("2023-02-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateIpamPoolAllocationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reserves a custom CIDR block from an IP Address Manager (IPAM) pool.
//
// Description:
//
//	  Before you reserve a custom CIDR block, make sure that an IPAM pool is created and CIDR blocks are added to the pool. You can call **CreateIpamPool*	- to create an IPAM pool and call **AddIpamPoolCidr*	- to add CIDR blocks to the pool.
//
//		- When you specify Cidr or CidrMask to reserve a custom CIDR block, the mask must fall within the range specified by the IPAM pool.
//
//		- If the IPAM pool has the region attribute, you must reserve a custom CIDR block in the region to which the IPAM pool belongs.
//
//		- The custom CIDR block that you want to reserve cannot overlap with existing CIDR blocks created from the IPAM pool.
//
// @param request - CreateIpamPoolAllocationRequest
//
// @return CreateIpamPoolAllocationResponse
func (client *Client) CreateIpamPoolAllocation(request *CreateIpamPoolAllocationRequest) (_result *CreateIpamPoolAllocationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateIpamPoolAllocationResponse{}
	_body, _err := client.CreateIpamPoolAllocationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a custom resource discovery instance.
//
// Description:
//
//	  Each Alibaba Cloud account can create only one resource discovery instance in each region.
//
//		- You can create only custom resource discovery instances.
//
// @param request - CreateIpamResourceDiscoveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIpamResourceDiscoveryResponse
func (client *Client) CreateIpamResourceDiscoveryWithOptions(request *CreateIpamResourceDiscoveryRequest, runtime *util.RuntimeOptions) (_result *CreateIpamResourceDiscoveryResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.IpamResourceDiscoveryDescription)) {
		query["IpamResourceDiscoveryDescription"] = request.IpamResourceDiscoveryDescription
	}

	if !tea.BoolValue(util.IsUnset(request.IpamResourceDiscoveryName)) {
		query["IpamResourceDiscoveryName"] = request.IpamResourceDiscoveryName
	}

	if !tea.BoolValue(util.IsUnset(request.OperatingRegionList)) {
		query["OperatingRegionList"] = request.OperatingRegionList
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

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateIpamResourceDiscovery"),
		Version:     tea.String("2023-02-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateIpamResourceDiscoveryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom resource discovery instance.
//
// Description:
//
//	  Each Alibaba Cloud account can create only one resource discovery instance in each region.
//
//		- You can create only custom resource discovery instances.
//
// @param request - CreateIpamResourceDiscoveryRequest
//
// @return CreateIpamResourceDiscoveryResponse
func (client *Client) CreateIpamResourceDiscovery(request *CreateIpamResourceDiscoveryRequest) (_result *CreateIpamResourceDiscoveryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateIpamResourceDiscoveryResponse{}
	_body, _err := client.CreateIpamResourceDiscoveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a public scope and private scope to respectively manage public and private IP addresses.
//
// @param request - CreateIpamScopeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIpamScopeResponse
func (client *Client) CreateIpamScopeWithOptions(request *CreateIpamScopeRequest, runtime *util.RuntimeOptions) (_result *CreateIpamScopeResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.IpamId)) {
		query["IpamId"] = request.IpamId
	}

	if !tea.BoolValue(util.IsUnset(request.IpamScopeDescription)) {
		query["IpamScopeDescription"] = request.IpamScopeDescription
	}

	if !tea.BoolValue(util.IsUnset(request.IpamScopeName)) {
		query["IpamScopeName"] = request.IpamScopeName
	}

	if !tea.BoolValue(util.IsUnset(request.IpamScopeType)) {
		query["IpamScopeType"] = request.IpamScopeType
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

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateIpamScope"),
		Version:     tea.String("2023-02-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateIpamScopeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a public scope and private scope to respectively manage public and private IP addresses.
//
// @param request - CreateIpamScopeRequest
//
// @return CreateIpamScopeResponse
func (client *Client) CreateIpamScope(request *CreateIpamScopeRequest) (_result *CreateIpamScopeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateIpamScopeResponse{}
	_body, _err := client.CreateIpamScopeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an IP Address Manager (IPAM).
//
// Description:
//
// ## [](#)Prerequisites
//
//   - Before you delete an IPAM, make sure that all IPAM pools of the IPAM are deleted. You can call **DeleteIpamPool*	- to delete IPAM pools.
//
//   - Before you delete an IPAM, make sure that all IPAM scopes of the IPAM are deleted. You can call **DeleteIpamScope*	- to delete IPAM scopes.
//
// @param request - DeleteIpamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIpamResponse
func (client *Client) DeleteIpamWithOptions(request *DeleteIpamRequest, runtime *util.RuntimeOptions) (_result *DeleteIpamResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.IpamId)) {
		query["IpamId"] = request.IpamId
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteIpam"),
		Version:     tea.String("2023-02-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteIpamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an IP Address Manager (IPAM).
//
// Description:
//
// ## [](#)Prerequisites
//
//   - Before you delete an IPAM, make sure that all IPAM pools of the IPAM are deleted. You can call **DeleteIpamPool*	- to delete IPAM pools.
//
//   - Before you delete an IPAM, make sure that all IPAM scopes of the IPAM are deleted. You can call **DeleteIpamScope*	- to delete IPAM scopes.
//
// @param request - DeleteIpamRequest
//
// @return DeleteIpamResponse
func (client *Client) DeleteIpam(request *DeleteIpamRequest) (_result *DeleteIpamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteIpamResponse{}
	_body, _err := client.DeleteIpamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an IP Address Manager (IPAM) scope.
//
// Description:
//
// ### [](#)Usage notes
//
//   - Before you delete a parent pool, make sure that all subpools of the parent pool are deleted.
//
//   - If an effective region is specified for a parent pool and IP addresses are allocated from the parent pool, you cannot delete the parent pool.
//
//   - If an effective region is specified for a subpool and IP addresses are allocated from the subpool, you cannot delete the subpool.
//
// @param request - DeleteIpamPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIpamPoolResponse
func (client *Client) DeleteIpamPoolWithOptions(request *DeleteIpamPoolRequest, runtime *util.RuntimeOptions) (_result *DeleteIpamPoolResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.IpamPoolId)) {
		query["IpamPoolId"] = request.IpamPoolId
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteIpamPool"),
		Version:     tea.String("2023-02-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteIpamPoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an IP Address Manager (IPAM) scope.
//
// Description:
//
// ### [](#)Usage notes
//
//   - Before you delete a parent pool, make sure that all subpools of the parent pool are deleted.
//
//   - If an effective region is specified for a parent pool and IP addresses are allocated from the parent pool, you cannot delete the parent pool.
//
//   - If an effective region is specified for a subpool and IP addresses are allocated from the subpool, you cannot delete the subpool.
//
// @param request - DeleteIpamPoolRequest
//
// @return DeleteIpamPoolResponse
func (client *Client) DeleteIpamPool(request *DeleteIpamPoolRequest) (_result *DeleteIpamPoolResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteIpamPoolResponse{}
	_body, _err := client.DeleteIpamPoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a custom reserved CIDR block from an IP Address Manager (IPAM) pool.
//
// @param request - DeleteIpamPoolAllocationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIpamPoolAllocationResponse
func (client *Client) DeleteIpamPoolAllocationWithOptions(request *DeleteIpamPoolAllocationRequest, runtime *util.RuntimeOptions) (_result *DeleteIpamPoolAllocationResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.IpamPoolAllocationId)) {
		query["IpamPoolAllocationId"] = request.IpamPoolAllocationId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteIpamPoolAllocation"),
		Version:     tea.String("2023-02-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteIpamPoolAllocationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom reserved CIDR block from an IP Address Manager (IPAM) pool.
//
// @param request - DeleteIpamPoolAllocationRequest
//
// @return DeleteIpamPoolAllocationResponse
func (client *Client) DeleteIpamPoolAllocation(request *DeleteIpamPoolAllocationRequest) (_result *DeleteIpamPoolAllocationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteIpamPoolAllocationResponse{}
	_body, _err := client.DeleteIpamPoolAllocationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a CIDR block provisioned to an IP Address Manager (IPAM) pool.
//
// Description:
//
//	  If CIDR blocks are provisioned to a parent pool and its subpools, you must first delete the CIDR blocks provisioned to the subpools before you delete the ones provisioned to the parent pool.
//
//		- If CIDR blocks are provisioned only to the parent pool, directly delete them.
//
//		- If CIDR blocks are allocated from provisioned ones, you must first delete the allocated CIDR blocks before you delete the provisioned ones.
//
//		- You can delete CIDR blocks provisioned to an IPAM pool only in the region where the IPAM is hosted.
//
// @param request - DeleteIpamPoolCidrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIpamPoolCidrResponse
func (client *Client) DeleteIpamPoolCidrWithOptions(request *DeleteIpamPoolCidrRequest, runtime *util.RuntimeOptions) (_result *DeleteIpamPoolCidrResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cidr)) {
		query["Cidr"] = request.Cidr
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.IpamPoolId)) {
		query["IpamPoolId"] = request.IpamPoolId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteIpamPoolCidr"),
		Version:     tea.String("2023-02-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteIpamPoolCidrResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a CIDR block provisioned to an IP Address Manager (IPAM) pool.
//
// Description:
//
//	  If CIDR blocks are provisioned to a parent pool and its subpools, you must first delete the CIDR blocks provisioned to the subpools before you delete the ones provisioned to the parent pool.
//
//		- If CIDR blocks are provisioned only to the parent pool, directly delete them.
//
//		- If CIDR blocks are allocated from provisioned ones, you must first delete the allocated CIDR blocks before you delete the provisioned ones.
//
//		- You can delete CIDR blocks provisioned to an IPAM pool only in the region where the IPAM is hosted.
//
// @param request - DeleteIpamPoolCidrRequest
//
// @return DeleteIpamPoolCidrResponse
func (client *Client) DeleteIpamPoolCidr(request *DeleteIpamPoolCidrRequest) (_result *DeleteIpamPoolCidrResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteIpamPoolCidrResponse{}
	_body, _err := client.DeleteIpamPoolCidrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a custom resource discovery instance.
//
// Description:
//
//	If a resource discovery instance is shared, it cannot be deleted.
//
// @param request - DeleteIpamResourceDiscoveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIpamResourceDiscoveryResponse
func (client *Client) DeleteIpamResourceDiscoveryWithOptions(request *DeleteIpamResourceDiscoveryRequest, runtime *util.RuntimeOptions) (_result *DeleteIpamResourceDiscoveryResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.IpamResourceDiscoveryId)) {
		query["IpamResourceDiscoveryId"] = request.IpamResourceDiscoveryId
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteIpamResourceDiscovery"),
		Version:     tea.String("2023-02-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteIpamResourceDiscoveryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom resource discovery instance.
//
// Description:
//
//	If a resource discovery instance is shared, it cannot be deleted.
//
// @param request - DeleteIpamResourceDiscoveryRequest
//
// @return DeleteIpamResourceDiscoveryResponse
func (client *Client) DeleteIpamResourceDiscovery(request *DeleteIpamResourceDiscoveryRequest) (_result *DeleteIpamResourceDiscoveryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteIpamResourceDiscoveryResponse{}
	_body, _err := client.DeleteIpamResourceDiscoveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an IP Address Manager (IPAM) scope.
//
// Description:
//
// ### [](#)Usage notes
//
//   - You cannot delete the private scope and public scope created by the system.
//
//   - Before you delete an IPAM scope, make sure that all pools within the scope are deleted. You can call **DeleteIpamPool*	- to delete IPAM pools.
//
// @param request - DeleteIpamScopeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIpamScopeResponse
func (client *Client) DeleteIpamScopeWithOptions(request *DeleteIpamScopeRequest, runtime *util.RuntimeOptions) (_result *DeleteIpamScopeResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.IpamScopeId)) {
		query["IpamScopeId"] = request.IpamScopeId
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteIpamScope"),
		Version:     tea.String("2023-02-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteIpamScopeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an IP Address Manager (IPAM) scope.
//
// Description:
//
// ### [](#)Usage notes
//
//   - You cannot delete the private scope and public scope created by the system.
//
//   - Before you delete an IPAM scope, make sure that all pools within the scope are deleted. You can call **DeleteIpamPool*	- to delete IPAM pools.
//
// @param request - DeleteIpamScopeRequest
//
// @return DeleteIpamScopeResponse
func (client *Client) DeleteIpamScope(request *DeleteIpamScopeRequest) (_result *DeleteIpamScopeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteIpamScopeResponse{}
	_body, _err := client.DeleteIpamScopeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disassociates resource discovery and IPAM instances.
//
// @param request - DissociateIpamResourceDiscoveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DissociateIpamResourceDiscoveryResponse
func (client *Client) DissociateIpamResourceDiscoveryWithOptions(request *DissociateIpamResourceDiscoveryRequest, runtime *util.RuntimeOptions) (_result *DissociateIpamResourceDiscoveryResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.IpamId)) {
		query["IpamId"] = request.IpamId
	}

	if !tea.BoolValue(util.IsUnset(request.IpamResourceDiscoveryId)) {
		query["IpamResourceDiscoveryId"] = request.IpamResourceDiscoveryId
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DissociateIpamResourceDiscovery"),
		Version:     tea.String("2023-02-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DissociateIpamResourceDiscoveryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates resource discovery and IPAM instances.
//
// @param request - DissociateIpamResourceDiscoveryRequest
//
// @return DissociateIpamResourceDiscoveryResponse
func (client *Client) DissociateIpamResourceDiscovery(request *DissociateIpamResourceDiscoveryRequest) (_result *DissociateIpamResourceDiscoveryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DissociateIpamResourceDiscoveryResponse{}
	_body, _err := client.DissociateIpamResourceDiscoveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries CIDR block allocations of an IP Address Manager (IPAM) pool.
//
// @param request - GetIpamPoolAllocationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIpamPoolAllocationResponse
func (client *Client) GetIpamPoolAllocationWithOptions(request *GetIpamPoolAllocationRequest, runtime *util.RuntimeOptions) (_result *GetIpamPoolAllocationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetIpamPoolAllocation"),
		Version:     tea.String("2023-02-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetIpamPoolAllocationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries CIDR block allocations of an IP Address Manager (IPAM) pool.
//
// @param request - GetIpamPoolAllocationRequest
//
// @return GetIpamPoolAllocationResponse
func (client *Client) GetIpamPoolAllocation(request *GetIpamPoolAllocationRequest) (_result *GetIpamPoolAllocationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetIpamPoolAllocationResponse{}
	_body, _err := client.GetIpamPoolAllocationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Gets the available CIDR blocks of the IPAM pool.
//
// @param request - GetIpamPoolNextAvailableCidrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIpamPoolNextAvailableCidrResponse
func (client *Client) GetIpamPoolNextAvailableCidrWithOptions(request *GetIpamPoolNextAvailableCidrRequest, runtime *util.RuntimeOptions) (_result *GetIpamPoolNextAvailableCidrResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetIpamPoolNextAvailableCidr"),
		Version:     tea.String("2023-02-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetIpamPoolNextAvailableCidrResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets the available CIDR blocks of the IPAM pool.
//
// @param request - GetIpamPoolNextAvailableCidrRequest
//
// @return GetIpamPoolNextAvailableCidrResponse
func (client *Client) GetIpamPoolNextAvailableCidr(request *GetIpamPoolNextAvailableCidrRequest) (_result *GetIpamPoolNextAvailableCidrResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetIpamPoolNextAvailableCidrResponse{}
	_body, _err := client.GetIpamPoolNextAvailableCidrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries whether IP Address Manager (IPAM) is activated.
//
// @param request - GetVpcIpamServiceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVpcIpamServiceStatusResponse
func (client *Client) GetVpcIpamServiceStatusWithOptions(request *GetVpcIpamServiceStatusRequest, runtime *util.RuntimeOptions) (_result *GetVpcIpamServiceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetVpcIpamServiceStatus"),
		Version:     tea.String("2023-02-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetVpcIpamServiceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether IP Address Manager (IPAM) is activated.
//
// @param request - GetVpcIpamServiceStatusRequest
//
// @return GetVpcIpamServiceStatusResponse
func (client *Client) GetVpcIpamServiceStatus(request *GetVpcIpamServiceStatusRequest) (_result *GetVpcIpamServiceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVpcIpamServiceStatusResponse{}
	_body, _err := client.GetVpcIpamServiceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries discovered resources.
//
// @param request - ListIpamDiscoveredResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIpamDiscoveredResourceResponse
func (client *Client) ListIpamDiscoveredResourceWithOptions(request *ListIpamDiscoveredResourceRequest, runtime *util.RuntimeOptions) (_result *ListIpamDiscoveredResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IpamResourceDiscoveryId)) {
		query["IpamResourceDiscoveryId"] = request.IpamResourceDiscoveryId
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

	if !tea.BoolValue(util.IsUnset(request.ResourceRegionId)) {
		query["ResourceRegionId"] = request.ResourceRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListIpamDiscoveredResource"),
		Version:     tea.String("2023-02-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIpamDiscoveredResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries discovered resources.
//
// @param request - ListIpamDiscoveredResourceRequest
//
// @return ListIpamDiscoveredResourceResponse
func (client *Client) ListIpamDiscoveredResource(request *ListIpamDiscoveredResourceRequest) (_result *ListIpamDiscoveredResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListIpamDiscoveredResourceResponse{}
	_body, _err := client.ListIpamDiscoveredResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries CIDR block allocations of an IP Address Manager (IPAM) pool.
//
// @param request - ListIpamPoolAllocationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIpamPoolAllocationsResponse
func (client *Client) ListIpamPoolAllocationsWithOptions(request *ListIpamPoolAllocationsRequest, runtime *util.RuntimeOptions) (_result *ListIpamPoolAllocationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cidr)) {
		query["Cidr"] = request.Cidr
	}

	if !tea.BoolValue(util.IsUnset(request.IpamPoolAllocationIds)) {
		query["IpamPoolAllocationIds"] = request.IpamPoolAllocationIds
	}

	if !tea.BoolValue(util.IsUnset(request.IpamPoolAllocationName)) {
		query["IpamPoolAllocationName"] = request.IpamPoolAllocationName
	}

	if !tea.BoolValue(util.IsUnset(request.IpamPoolId)) {
		query["IpamPoolId"] = request.IpamPoolId
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
		Action:      tea.String("ListIpamPoolAllocations"),
		Version:     tea.String("2023-02-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIpamPoolAllocationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries CIDR block allocations of an IP Address Manager (IPAM) pool.
//
// @param request - ListIpamPoolAllocationsRequest
//
// @return ListIpamPoolAllocationsResponse
func (client *Client) ListIpamPoolAllocations(request *ListIpamPoolAllocationsRequest) (_result *ListIpamPoolAllocationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListIpamPoolAllocationsResponse{}
	_body, _err := client.ListIpamPoolAllocationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries CIDR blocks provisioned to an IP Address Manager (IPAM) pool.
//
// @param request - ListIpamPoolCidrsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIpamPoolCidrsResponse
func (client *Client) ListIpamPoolCidrsWithOptions(request *ListIpamPoolCidrsRequest, runtime *util.RuntimeOptions) (_result *ListIpamPoolCidrsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cidr)) {
		query["Cidr"] = request.Cidr
	}

	if !tea.BoolValue(util.IsUnset(request.IpamPoolId)) {
		query["IpamPoolId"] = request.IpamPoolId
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
		Action:      tea.String("ListIpamPoolCidrs"),
		Version:     tea.String("2023-02-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIpamPoolCidrsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries CIDR blocks provisioned to an IP Address Manager (IPAM) pool.
//
// @param request - ListIpamPoolCidrsRequest
//
// @return ListIpamPoolCidrsResponse
func (client *Client) ListIpamPoolCidrs(request *ListIpamPoolCidrsRequest) (_result *ListIpamPoolCidrsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListIpamPoolCidrsResponse{}
	_body, _err := client.ListIpamPoolCidrsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries IP Address Manager (IPAM) pools.
//
// @param request - ListIpamPoolsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIpamPoolsResponse
func (client *Client) ListIpamPoolsWithOptions(request *ListIpamPoolsRequest, runtime *util.RuntimeOptions) (_result *ListIpamPoolsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IpamPoolIds)) {
		query["IpamPoolIds"] = request.IpamPoolIds
	}

	if !tea.BoolValue(util.IsUnset(request.IpamPoolName)) {
		query["IpamPoolName"] = request.IpamPoolName
	}

	if !tea.BoolValue(util.IsUnset(request.IpamScopeId)) {
		query["IpamScopeId"] = request.IpamScopeId
	}

	if !tea.BoolValue(util.IsUnset(request.IsShared)) {
		query["IsShared"] = request.IsShared
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PoolRegionId)) {
		query["PoolRegionId"] = request.PoolRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIpamPoolId)) {
		query["SourceIpamPoolId"] = request.SourceIpamPoolId
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListIpamPools"),
		Version:     tea.String("2023-02-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIpamPoolsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries IP Address Manager (IPAM) pools.
//
// @param request - ListIpamPoolsRequest
//
// @return ListIpamPoolsResponse
func (client *Client) ListIpamPools(request *ListIpamPoolsRequest) (_result *ListIpamPoolsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListIpamPoolsResponse{}
	_body, _err := client.ListIpamPoolsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries resources in an IP Address Manager (IPAM) pool.
//
// @param request - ListIpamResourceCidrsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIpamResourceCidrsResponse
func (client *Client) ListIpamResourceCidrsWithOptions(request *ListIpamResourceCidrsRequest, runtime *util.RuntimeOptions) (_result *ListIpamResourceCidrsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IpamPoolId)) {
		query["IpamPoolId"] = request.IpamPoolId
	}

	if !tea.BoolValue(util.IsUnset(request.IpamScopeId)) {
		query["IpamScopeId"] = request.IpamScopeId
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

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListIpamResourceCidrs"),
		Version:     tea.String("2023-02-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIpamResourceCidrsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries resources in an IP Address Manager (IPAM) pool.
//
// @param request - ListIpamResourceCidrsRequest
//
// @return ListIpamResourceCidrsResponse
func (client *Client) ListIpamResourceCidrs(request *ListIpamResourceCidrsRequest) (_result *ListIpamResourceCidrsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListIpamResourceCidrsResponse{}
	_body, _err := client.ListIpamResourceCidrsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries IPAM resource discovery instances.
//
// @param request - ListIpamResourceDiscoveriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIpamResourceDiscoveriesResponse
func (client *Client) ListIpamResourceDiscoveriesWithOptions(request *ListIpamResourceDiscoveriesRequest, runtime *util.RuntimeOptions) (_result *ListIpamResourceDiscoveriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IpamResourceDiscoveryIds)) {
		query["IpamResourceDiscoveryIds"] = request.IpamResourceDiscoveryIds
	}

	if !tea.BoolValue(util.IsUnset(request.IpamResourceDiscoveryName)) {
		query["IpamResourceDiscoveryName"] = request.IpamResourceDiscoveryName
	}

	if !tea.BoolValue(util.IsUnset(request.IsShared)) {
		query["IsShared"] = request.IsShared
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

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

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListIpamResourceDiscoveries"),
		Version:     tea.String("2023-02-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIpamResourceDiscoveriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries IPAM resource discovery instances.
//
// @param request - ListIpamResourceDiscoveriesRequest
//
// @return ListIpamResourceDiscoveriesResponse
func (client *Client) ListIpamResourceDiscoveries(request *ListIpamResourceDiscoveriesRequest) (_result *ListIpamResourceDiscoveriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListIpamResourceDiscoveriesResponse{}
	_body, _err := client.ListIpamResourceDiscoveriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the association between resource discovery and IPAM.
//
// @param request - ListIpamResourceDiscoveryAssociationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIpamResourceDiscoveryAssociationsResponse
func (client *Client) ListIpamResourceDiscoveryAssociationsWithOptions(request *ListIpamResourceDiscoveryAssociationsRequest, runtime *util.RuntimeOptions) (_result *ListIpamResourceDiscoveryAssociationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IpamId)) {
		query["IpamId"] = request.IpamId
	}

	if !tea.BoolValue(util.IsUnset(request.IpamResourceDiscoveryId)) {
		query["IpamResourceDiscoveryId"] = request.IpamResourceDiscoveryId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

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

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListIpamResourceDiscoveryAssociations"),
		Version:     tea.String("2023-02-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIpamResourceDiscoveryAssociationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the association between resource discovery and IPAM.
//
// @param request - ListIpamResourceDiscoveryAssociationsRequest
//
// @return ListIpamResourceDiscoveryAssociationsResponse
func (client *Client) ListIpamResourceDiscoveryAssociations(request *ListIpamResourceDiscoveryAssociationsRequest) (_result *ListIpamResourceDiscoveryAssociationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListIpamResourceDiscoveryAssociationsResponse{}
	_body, _err := client.ListIpamResourceDiscoveryAssociationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries IP Address Manager (IPAM) scopes.
//
// @param request - ListIpamScopesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIpamScopesResponse
func (client *Client) ListIpamScopesWithOptions(request *ListIpamScopesRequest, runtime *util.RuntimeOptions) (_result *ListIpamScopesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IpamId)) {
		query["IpamId"] = request.IpamId
	}

	if !tea.BoolValue(util.IsUnset(request.IpamScopeIds)) {
		query["IpamScopeIds"] = request.IpamScopeIds
	}

	if !tea.BoolValue(util.IsUnset(request.IpamScopeName)) {
		query["IpamScopeName"] = request.IpamScopeName
	}

	if !tea.BoolValue(util.IsUnset(request.IpamScopeType)) {
		query["IpamScopeType"] = request.IpamScopeType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

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

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListIpamScopes"),
		Version:     tea.String("2023-02-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIpamScopesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries IP Address Manager (IPAM) scopes.
//
// @param request - ListIpamScopesRequest
//
// @return ListIpamScopesResponse
func (client *Client) ListIpamScopes(request *ListIpamScopesRequest) (_result *ListIpamScopesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListIpamScopesResponse{}
	_body, _err := client.ListIpamScopesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries IP Address Managers (IPAMs).
//
// @param request - ListIpamsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIpamsResponse
func (client *Client) ListIpamsWithOptions(request *ListIpamsRequest, runtime *util.RuntimeOptions) (_result *ListIpamsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IpamIds)) {
		query["IpamIds"] = request.IpamIds
	}

	if !tea.BoolValue(util.IsUnset(request.IpamName)) {
		query["IpamName"] = request.IpamName
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

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

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListIpams"),
		Version:     tea.String("2023-02-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIpamsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries IP Address Managers (IPAMs).
//
// @param request - ListIpamsRequest
//
// @return ListIpamsResponse
func (client *Client) ListIpams(request *ListIpamsRequest) (_result *ListIpamsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListIpamsResponse{}
	_body, _err := client.ListIpamsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of resource tags.
//
// Description:
//
// ### [](#)Usage notes
//
//   - You must specify **ResourceId.N*	- or **Tag.N*	- that consists of **Tag.N.Key*	- and **Tag.N.Value*	- in the request to specify the object that you want to query.
//
//   - **Tag.N*	- is a resource tag that consists of a key-value pair. If you specify only **Tag.N.Key**, all tag values that are associated with the specified key are returned. If you specify only **Tag.N.Value**, an error message is returned.
//
//   - If you specify **Tag.N*	- and **ResourceId.N*	- to filter tags, **ResourceId.N*	- must match all specified key-value pairs.
//
//   - If you specify multiple key-value pairs, resources that contain these key-value pairs are returned.
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
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

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

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2023-02-28"),
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
// Queries a list of resource tags.
//
// Description:
//
// ### [](#)Usage notes
//
//   - You must specify **ResourceId.N*	- or **Tag.N*	- that consists of **Tag.N.Key*	- and **Tag.N.Value*	- in the request to specify the object that you want to query.
//
//   - **Tag.N*	- is a resource tag that consists of a key-value pair. If you specify only **Tag.N.Key**, all tag values that are associated with the specified key are returned. If you specify only **Tag.N.Value**, an error message is returned.
//
//   - If you specify **Tag.N*	- and **ResourceId.N*	- to filter tags, **ResourceId.N*	- must match all specified key-value pairs.
//
//   - If you specify multiple key-value pairs, resources that contain these key-value pairs are returned.
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
// Activates IP Address Manager (IPAM).
//
// @param request - OpenVpcIpamServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenVpcIpamServiceResponse
func (client *Client) OpenVpcIpamServiceWithOptions(request *OpenVpcIpamServiceRequest, runtime *util.RuntimeOptions) (_result *OpenVpcIpamServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OpenVpcIpamService"),
		Version:     tea.String("2023-02-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenVpcIpamServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Activates IP Address Manager (IPAM).
//
// @param request - OpenVpcIpamServiceRequest
//
// @return OpenVpcIpamServiceResponse
func (client *Client) OpenVpcIpamService(request *OpenVpcIpamServiceRequest) (_result *OpenVpcIpamServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenVpcIpamServiceResponse{}
	_body, _err := client.OpenVpcIpamServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a tag to a resource.
//
// Description:
//
// ### [](#)Usage notes
//
// Tags are used to classify instances. Each tag consists of a key-value pair. Before you use tags, take note of the following items:
//
//   - Each tag key that is added to an instance must be unique.
//
//   - You cannot create tags without adding them to instances. All tags must be added to instances.
//
//   - You can add at most 20 tags to each instance. Before you add a tag to an instance, the system automatically checks the number of existing tags. An error message is returned if the maximum number of tags is reached.
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
	query := map[string]interface{}{}
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

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2023-02-28"),
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
// Adds a tag to a resource.
//
// Description:
//
// ### [](#)Usage notes
//
// Tags are used to classify instances. Each tag consists of a key-value pair. Before you use tags, take note of the following items:
//
//   - Each tag key that is added to an instance must be unique.
//
//   - You cannot create tags without adding them to instances. All tags must be added to instances.
//
//   - You can add at most 20 tags to each instance. Before you add a tag to an instance, the system automatically checks the number of existing tags. An error message is returned if the maximum number of tags is reached.
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
// Removes a tag from a resource.
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

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2023-02-28"),
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
// Removes a tag from a resource.
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
// Updates an IP Address Manager (IPAM).
//
// @param request - UpdateIpamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIpamResponse
func (client *Client) UpdateIpamWithOptions(request *UpdateIpamRequest, runtime *util.RuntimeOptions) (_result *UpdateIpamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddOperatingRegion)) {
		query["AddOperatingRegion"] = request.AddOperatingRegion
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.IpamDescription)) {
		query["IpamDescription"] = request.IpamDescription
	}

	if !tea.BoolValue(util.IsUnset(request.IpamId)) {
		query["IpamId"] = request.IpamId
	}

	if !tea.BoolValue(util.IsUnset(request.IpamName)) {
		query["IpamName"] = request.IpamName
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

	if !tea.BoolValue(util.IsUnset(request.RemoveOperatingRegion)) {
		query["RemoveOperatingRegion"] = request.RemoveOperatingRegion
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateIpam"),
		Version:     tea.String("2023-02-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateIpamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an IP Address Manager (IPAM).
//
// @param request - UpdateIpamRequest
//
// @return UpdateIpamResponse
func (client *Client) UpdateIpam(request *UpdateIpamRequest) (_result *UpdateIpamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateIpamResponse{}
	_body, _err := client.UpdateIpamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the basic information about an IP Address Manager (IPAM) pool.
//
// @param request - UpdateIpamPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIpamPoolResponse
func (client *Client) UpdateIpamPoolWithOptions(request *UpdateIpamPoolRequest, runtime *util.RuntimeOptions) (_result *UpdateIpamPoolResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllocationDefaultCidrMask)) {
		query["AllocationDefaultCidrMask"] = request.AllocationDefaultCidrMask
	}

	if !tea.BoolValue(util.IsUnset(request.AllocationMaxCidrMask)) {
		query["AllocationMaxCidrMask"] = request.AllocationMaxCidrMask
	}

	if !tea.BoolValue(util.IsUnset(request.AllocationMinCidrMask)) {
		query["AllocationMinCidrMask"] = request.AllocationMinCidrMask
	}

	if !tea.BoolValue(util.IsUnset(request.AutoImport)) {
		query["AutoImport"] = request.AutoImport
	}

	if !tea.BoolValue(util.IsUnset(request.ClearAllocationDefaultCidrMask)) {
		query["ClearAllocationDefaultCidrMask"] = request.ClearAllocationDefaultCidrMask
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.IpamPoolDescription)) {
		query["IpamPoolDescription"] = request.IpamPoolDescription
	}

	if !tea.BoolValue(util.IsUnset(request.IpamPoolId)) {
		query["IpamPoolId"] = request.IpamPoolId
	}

	if !tea.BoolValue(util.IsUnset(request.IpamPoolName)) {
		query["IpamPoolName"] = request.IpamPoolName
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateIpamPool"),
		Version:     tea.String("2023-02-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateIpamPoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the basic information about an IP Address Manager (IPAM) pool.
//
// @param request - UpdateIpamPoolRequest
//
// @return UpdateIpamPoolResponse
func (client *Client) UpdateIpamPool(request *UpdateIpamPoolRequest) (_result *UpdateIpamPoolResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateIpamPoolResponse{}
	_body, _err := client.UpdateIpamPoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies CIDR block allocations of an IP Address Manager (IPAM) pool.
//
// @param request - UpdateIpamPoolAllocationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIpamPoolAllocationResponse
func (client *Client) UpdateIpamPoolAllocationWithOptions(request *UpdateIpamPoolAllocationRequest, runtime *util.RuntimeOptions) (_result *UpdateIpamPoolAllocationResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.IpamPoolAllocationDescription)) {
		query["IpamPoolAllocationDescription"] = request.IpamPoolAllocationDescription
	}

	if !tea.BoolValue(util.IsUnset(request.IpamPoolAllocationId)) {
		query["IpamPoolAllocationId"] = request.IpamPoolAllocationId
	}

	if !tea.BoolValue(util.IsUnset(request.IpamPoolAllocationName)) {
		query["IpamPoolAllocationName"] = request.IpamPoolAllocationName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateIpamPoolAllocation"),
		Version:     tea.String("2023-02-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateIpamPoolAllocationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies CIDR block allocations of an IP Address Manager (IPAM) pool.
//
// @param request - UpdateIpamPoolAllocationRequest
//
// @return UpdateIpamPoolAllocationResponse
func (client *Client) UpdateIpamPoolAllocation(request *UpdateIpamPoolAllocationRequest) (_result *UpdateIpamPoolAllocationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateIpamPoolAllocationResponse{}
	_body, _err := client.UpdateIpamPoolAllocationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a resource discovery instance.
//
// Description:
//
//	  You can add or remove effective regions only for custom resource discovery instances.
//
//		- When removing effective regions from a resource discovery instance, the managed region cannot be included.
//
// @param request - UpdateIpamResourceDiscoveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIpamResourceDiscoveryResponse
func (client *Client) UpdateIpamResourceDiscoveryWithOptions(request *UpdateIpamResourceDiscoveryRequest, runtime *util.RuntimeOptions) (_result *UpdateIpamResourceDiscoveryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddOperatingRegion)) {
		query["AddOperatingRegion"] = request.AddOperatingRegion
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.IpamResourceDiscoveryDescription)) {
		query["IpamResourceDiscoveryDescription"] = request.IpamResourceDiscoveryDescription
	}

	if !tea.BoolValue(util.IsUnset(request.IpamResourceDiscoveryId)) {
		query["IpamResourceDiscoveryId"] = request.IpamResourceDiscoveryId
	}

	if !tea.BoolValue(util.IsUnset(request.IpamResourceDiscoveryName)) {
		query["IpamResourceDiscoveryName"] = request.IpamResourceDiscoveryName
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

	if !tea.BoolValue(util.IsUnset(request.RemoveOperatingRegion)) {
		query["RemoveOperatingRegion"] = request.RemoveOperatingRegion
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateIpamResourceDiscovery"),
		Version:     tea.String("2023-02-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateIpamResourceDiscoveryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a resource discovery instance.
//
// Description:
//
//	  You can add or remove effective regions only for custom resource discovery instances.
//
//		- When removing effective regions from a resource discovery instance, the managed region cannot be included.
//
// @param request - UpdateIpamResourceDiscoveryRequest
//
// @return UpdateIpamResourceDiscoveryResponse
func (client *Client) UpdateIpamResourceDiscovery(request *UpdateIpamResourceDiscoveryRequest) (_result *UpdateIpamResourceDiscoveryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateIpamResourceDiscoveryResponse{}
	_body, _err := client.UpdateIpamResourceDiscoveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the basic information about an IP Address Manager (IPAM) scope.
//
// @param request - UpdateIpamScopeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIpamScopeResponse
func (client *Client) UpdateIpamScopeWithOptions(request *UpdateIpamScopeRequest, runtime *util.RuntimeOptions) (_result *UpdateIpamScopeResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.IpamScopeDescription)) {
		query["IpamScopeDescription"] = request.IpamScopeDescription
	}

	if !tea.BoolValue(util.IsUnset(request.IpamScopeId)) {
		query["IpamScopeId"] = request.IpamScopeId
	}

	if !tea.BoolValue(util.IsUnset(request.IpamScopeName)) {
		query["IpamScopeName"] = request.IpamScopeName
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateIpamScope"),
		Version:     tea.String("2023-02-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateIpamScopeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the basic information about an IP Address Manager (IPAM) scope.
//
// @param request - UpdateIpamScopeRequest
//
// @return UpdateIpamScopeResponse
func (client *Client) UpdateIpamScope(request *UpdateIpamScopeRequest) (_result *UpdateIpamScopeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateIpamScopeResponse{}
	_body, _err := client.UpdateIpamScopeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
