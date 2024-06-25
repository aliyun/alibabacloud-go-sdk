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
	// This parameter is required.
	//
	// example:
	//
	// 192.168.1.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ipam-pool-6rcq3tobayc20t****
	IpamPoolId *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
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

type ChangeResourceGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// rg-aek3ctkufaw****
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// ipam-uq5dcfc2eqhpf4****
	ResourceId           *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
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
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// This is my first Ipam
	IpamDescription *string `json:"IpamDescription,omitempty" xml:"IpamDescription,omitempty"`
	// example:
	//
	// abc
	IpamName            *string   `json:"IpamName,omitempty" xml:"IpamName,omitempty"`
	OperatingRegionList []*string `json:"OperatingRegionList,omitempty" xml:"OperatingRegionList,omitempty" type:"Repeated"`
	OwnerAccount        *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId             *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfmxazb4ph6aiy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
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

type CreateIpamResponseBody struct {
	// example:
	//
	// ipam-ccxbnsbhew0d6t****
	IpamId *string `json:"IpamId,omitempty" xml:"IpamId,omitempty"`
	// example:
	//
	// ipam-scope-8wiontzmiy6cg0****
	PrivateDefaultScopeId *string `json:"PrivateDefaultScopeId,omitempty" xml:"PrivateDefaultScopeId,omitempty"`
	// example:
	//
	// ipam-scope-r5c5c7bmym1brc****
	PublicDefaultScopeId *string `json:"PublicDefaultScopeId,omitempty" xml:"PublicDefaultScopeId,omitempty"`
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED39DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateIpamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIpamResponseBody) GoString() string {
	return s.String()
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
	// example:
	//
	// 28
	AllocationDefaultCidrMask *int32 `json:"AllocationDefaultCidrMask,omitempty" xml:"AllocationDefaultCidrMask,omitempty"`
	// example:
	//
	// 32
	AllocationMaxCidrMask *int32 `json:"AllocationMaxCidrMask,omitempty" xml:"AllocationMaxCidrMask,omitempty"`
	// example:
	//
	// 8
	AllocationMinCidrMask *int32 `json:"AllocationMinCidrMask,omitempty" xml:"AllocationMinCidrMask,omitempty"`
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// IPv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// example:
	//
	// test description
	IpamPoolDescription *string `json:"IpamPoolDescription,omitempty" xml:"IpamPoolDescription,omitempty"`
	// example:
	//
	// abc
	IpamPoolName *string `json:"IpamPoolName,omitempty" xml:"IpamPoolName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ipam-scope-glfmcyldpm8lsy****
	IpamScopeId  *string `json:"IpamScopeId,omitempty" xml:"IpamScopeId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// cn-hangzhou
	PoolRegionId *string `json:"PoolRegionId,omitempty" xml:"PoolRegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// ipam-pool-6rcq3tobayc20t****
	SourceIpamPoolId *string `json:"SourceIpamPoolId,omitempty" xml:"SourceIpamPoolId,omitempty"`
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

type CreateIpamPoolResponseBody struct {
	// example:
	//
	// ipam-pool-6rcq3tobayc20t****
	IpamPoolId *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
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
	// example:
	//
	// 192.168.1.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// example:
	//
	// 24
	CidrMask *int32 `json:"CidrMask,omitempty" xml:"CidrMask,omitempty"`
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ipam-pool-6rcq3tobayc20t****
	IpamPoolId *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
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

func (s *CreateIpamPoolAllocationRequest) SetIpamPoolId(v string) *CreateIpamPoolAllocationRequest {
	s.IpamPoolId = &v
	return s
}

func (s *CreateIpamPoolAllocationRequest) SetRegionId(v string) *CreateIpamPoolAllocationRequest {
	s.RegionId = &v
	return s
}

type CreateIpamPoolAllocationResponseBody struct {
	// example:
	//
	// 192.168.1.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// example:
	//
	// ipam-pool-alloc-112za33e4****
	IpamPoolAllocationId *string `json:"IpamPoolAllocationId,omitempty" xml:"IpamPoolAllocationId,omitempty"`
	// example:
	//
	// CE9CDAE5-341E-5D0B-AC8A-2BAC707D3EB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

type CreateIpamScopeRequest struct {
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ipam-ccxbnsbhew0d6t****
	IpamId *string `json:"IpamId,omitempty" xml:"IpamId,omitempty"`
	// example:
	//
	// test description
	IpamScopeDescription *string `json:"IpamScopeDescription,omitempty" xml:"IpamScopeDescription,omitempty"`
	// example:
	//
	// test
	IpamScopeName *string `json:"IpamScopeName,omitempty" xml:"IpamScopeName,omitempty"`
	// example:
	//
	// private
	IpamScopeType *string `json:"IpamScopeType,omitempty" xml:"IpamScopeType,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
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

func (s *CreateIpamScopeRequest) SetResourceOwnerAccount(v string) *CreateIpamScopeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateIpamScopeRequest) SetResourceOwnerId(v int64) *CreateIpamScopeRequest {
	s.ResourceOwnerId = &v
	return s
}

type CreateIpamScopeResponseBody struct {
	// example:
	//
	// ipam-scope-glfmcyldpm8lsy****
	IpamScopeId *string `json:"IpamScopeId,omitempty" xml:"IpamScopeId,omitempty"`
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
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ipam-ccxbnsbhew0d6t****
	IpamId       *string `json:"IpamId,omitempty" xml:"IpamId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ipam-pool-6rcq3tobayc20t****
	IpamPoolId   *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// example:
	//
	// 192.168.1.0/32
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ipam-alloc-c4vhvr3b22mmc6****
	IpamPoolAllocationId *string `json:"IpamPoolAllocationId,omitempty" xml:"IpamPoolAllocationId,omitempty"`
	// example:
	//
	// ipam-pool-6rcq3tobayc20t****
	IpamPoolId *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
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

func (s *DeleteIpamPoolAllocationRequest) SetCidr(v string) *DeleteIpamPoolAllocationRequest {
	s.Cidr = &v
	return s
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

func (s *DeleteIpamPoolAllocationRequest) SetIpamPoolId(v string) *DeleteIpamPoolAllocationRequest {
	s.IpamPoolId = &v
	return s
}

func (s *DeleteIpamPoolAllocationRequest) SetRegionId(v string) *DeleteIpamPoolAllocationRequest {
	s.RegionId = &v
	return s
}

type DeleteIpamPoolAllocationResponseBody struct {
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
	// This parameter is required.
	//
	// example:
	//
	// 192.168.1.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ipam-pool-6rcq3tobayc20t****
	IpamPoolId *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
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

type DeleteIpamScopeRequest struct {
	// example:
	//
	// 88144bdb-b190-4813-a3f5-66cc86694162
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ipam-scope-glfmcyldpm8lsy****
	IpamScopeId  *string `json:"IpamScopeId,omitempty" xml:"IpamScopeId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

type GetVpcIpamServiceStatusRequest struct {
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
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

type ListIpamPoolAllocationsRequest struct {
	// example:
	//
	// 192.168.1.0/24
	Cidr                  *string   `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	IpamPoolAllocationIds []*string `json:"IpamPoolAllocationIds,omitempty" xml:"IpamPoolAllocationIds,omitempty" type:"Repeated"`
	// example:
	//
	// ipam-pool-6rcq3tobayc20t****
	IpamPoolId *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
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
	IpamPoolAllocations []*ListIpamPoolAllocationsResponseBodyIpamPoolAllocations `json:"IpamPoolAllocations,omitempty" xml:"IpamPoolAllocations,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 3748DEFF-68BE-5EED-9937-7C1D0C21BAB4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 192.168.1.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// example:
	//
	// 2023-05-19T08:59:18Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// example:
	//
	// ipam-pool-alloc-112za33e4****
	IpamPoolAllocationId *string `json:"IpamPoolAllocationId,omitempty" xml:"IpamPoolAllocationId,omitempty"`
	// example:
	//
	// ipam-pool-6rcq3tobayc20t****
	IpamPoolId *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// vpc-bp16qjewdsunr41m1****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// 132193271328****
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// cn-hangzhou
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// example:
	//
	// Custom
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// 192.168.0.0/16
	SourceCidr *string `json:"SourceCidr,omitempty" xml:"SourceCidr,omitempty"`
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

func (s *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) SetIpamPoolAllocationId(v string) *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations {
	s.IpamPoolAllocationId = &v
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
	// example:
	//
	// 192.168.1.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ipam-pool-6rcq3tobayc20t****
	IpamPoolId *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
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
	IpamPoolCidrs []*ListIpamPoolCidrsResponseBodyIpamPoolCidrs `json:"IpamPoolCidrs,omitempty" xml:"IpamPoolCidrs,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 9E7CCB95-62E0-534D-9B9A-71F27E8B71B1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 192.168.1.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// example:
	//
	// ipam-pool-6rcq3tobayc20t****
	IpamPoolId *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
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
	IpamPoolIds []*string `json:"IpamPoolIds,omitempty" xml:"IpamPoolIds,omitempty" type:"Repeated"`
	// example:
	//
	// test
	IpamPoolName *string `json:"IpamPoolName,omitempty" xml:"IpamPoolName,omitempty"`
	// example:
	//
	// ipam-scope-glfmcyldpm8lsy****
	IpamScopeId *string `json:"IpamScopeId,omitempty" xml:"IpamScopeId,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// cn-hangzhou
	PoolRegionId *string `json:"PoolRegionId,omitempty" xml:"PoolRegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-aek2sermdd6****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// ipam-pool-lfnwi4jok1ss0g****
	SourceIpamPoolId *string                     `json:"SourceIpamPoolId,omitempty" xml:"SourceIpamPoolId,omitempty"`
	Tags             []*ListIpamPoolsRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	IpamPools []*ListIpamPoolsResponseBodyIpamPools `json:"IpamPools,omitempty" xml:"IpamPools,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// B54867DE-83DC-56B4-A57E-69A03119D0B1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 28
	AllocationDefaultCidrMask *int32 `json:"AllocationDefaultCidrMask,omitempty" xml:"AllocationDefaultCidrMask,omitempty"`
	// example:
	//
	// 32
	AllocationMaxCidrMask *int32 `json:"AllocationMaxCidrMask,omitempty" xml:"AllocationMaxCidrMask,omitempty"`
	// example:
	//
	// 8
	AllocationMinCidrMask *int32 `json:"AllocationMinCidrMask,omitempty" xml:"AllocationMinCidrMask,omitempty"`
	// example:
	//
	// 2023-04-19T16:49:01Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// true
	HasSubPool *bool `json:"HasSubPool,omitempty" xml:"HasSubPool,omitempty"`
	// example:
	//
	// IPv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// example:
	//
	// ipam-b5mtlx3q7xcnyr****
	IpamId *string `json:"IpamId,omitempty" xml:"IpamId,omitempty"`
	// example:
	//
	// test description
	IpamPoolDescription *string `json:"IpamPoolDescription,omitempty" xml:"IpamPoolDescription,omitempty"`
	// example:
	//
	// ipam-pool-6rcq3tobayc20t****
	IpamPoolId *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
	// example:
	//
	// test
	IpamPoolName *string `json:"IpamPoolName,omitempty" xml:"IpamPoolName,omitempty"`
	// example:
	//
	// cn-hangzhou
	IpamRegionId *string `json:"IpamRegionId,omitempty" xml:"IpamRegionId,omitempty"`
	// example:
	//
	// ipam-scope-glfmcyldpm8lsy****
	IpamScopeId *string `json:"IpamScopeId,omitempty" xml:"IpamScopeId,omitempty"`
	// example:
	//
	// private
	IpamScopeType *string `json:"IpamScopeType,omitempty" xml:"IpamScopeType,omitempty"`
	// example:
	//
	// 1210123456******
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 2
	PoolDepth *int32 `json:"PoolDepth,omitempty" xml:"PoolDepth,omitempty"`
	// example:
	//
	// cn-hangzhou
	PoolRegionId *string `json:"PoolRegionId,omitempty" xml:"PoolRegionId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// ipam-pool-lfnwi4jok1ss0g****
	SourceIpamPoolId *string `json:"SourceIpamPoolId,omitempty" xml:"SourceIpamPoolId,omitempty"`
	// example:
	//
	// Created
	Status *string                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   []*ListIpamPoolsResponseBodyIpamPoolsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	// example:
	//
	// ipam-pool-6rcq3tobayc20t****
	IpamPoolId *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
	// example:
	//
	// ipam-scope-glfmcyldpm8lsy****
	IpamScopeId *string `json:"IpamScopeId,omitempty" xml:"IpamScopeId,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
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
	// example:
	//
	// vpc-bp16qjewdsunr41m1****
	ResourceId      *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// VPC
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
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

type ListIpamResourceCidrsResponseBody struct {
	IpamResourceCidrs []*ListIpamResourceCidrsResponseBodyIpamResourceCidrs `json:"IpamResourceCidrs,omitempty" xml:"IpamResourceCidrs,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 49A9DE56-B68C-5FFC-BC06-509D086F287C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 132193271328****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// example:
	//
	// 192.168.1.0/32
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// example:
	//
	// Compliant
	ComplianceStatus *string `json:"ComplianceStatus,omitempty" xml:"ComplianceStatus,omitempty"`
	// example:
	//
	// 0
	IpUsage *string `json:"IpUsage,omitempty" xml:"IpUsage,omitempty"`
	// example:
	//
	// ipam-pool-alloc-112za33e4****
	IpamAllocationId *string `json:"IpamAllocationId,omitempty" xml:"IpamAllocationId,omitempty"`
	// example:
	//
	// ipam-uq5dcfc2eqhpf4****
	IpamId *string `json:"IpamId,omitempty" xml:"IpamId,omitempty"`
	// example:
	//
	// ipam-pool-6rcq3tobayc20t***
	IpamPoolId *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
	// example:
	//
	// ipam-scope-glfmcyldpm8lsy****
	IpamScopeId *string `json:"IpamScopeId,omitempty" xml:"IpamScopeId,omitempty"`
	// example:
	//
	// Managed
	ManagementStatus *string `json:"ManagementStatus,omitempty" xml:"ManagementStatus,omitempty"`
	// example:
	//
	// Nonoverlapping
	OverlapStatus *string `json:"OverlapStatus,omitempty" xml:"OverlapStatus,omitempty"`
	// example:
	//
	// vpc-bp16qjewdsunr41m1****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// 132193271328****
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// cn-hangzhou
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// example:
	//
	// VPC
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// 192.168.1.0/24
	SourceCidr *string `json:"SourceCidr,omitempty" xml:"SourceCidr,omitempty"`
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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

type ListIpamScopesRequest struct {
	// example:
	//
	// ipam-ccxbnsbhew0d6t****
	IpamId       *string   `json:"IpamId,omitempty" xml:"IpamId,omitempty"`
	IpamScopeIds []*string `json:"IpamScopeIds,omitempty" xml:"IpamScopeIds,omitempty" type:"Repeated"`
	// example:
	//
	// test
	IpamScopeName *string `json:"IpamScopeName,omitempty" xml:"IpamScopeName,omitempty"`
	// example:
	//
	// private
	IpamScopeType *string `json:"IpamScopeType,omitempty" xml:"IpamScopeType,omitempty"`
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-aek2sermdd6****
	ResourceGroupId      *string                      `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string                      `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                       `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Tags                 []*ListIpamScopesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	IpamScopes []*ListIpamScopesResponseBodyIpamScopes `json:"IpamScopes,omitempty" xml:"IpamScopes,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 8859C501-97E7-53D4-B94B-2A9E16003B22
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 2022-04-18T03:12:37Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// ipam-ccxbnsbhew0d6t****
	IpamId *string `json:"IpamId,omitempty" xml:"IpamId,omitempty"`
	// example:
	//
	// test description
	IpamScopeDescription *string `json:"IpamScopeDescription,omitempty" xml:"IpamScopeDescription,omitempty"`
	// example:
	//
	// ipam-scope-glfmcyldpm8lsy****
	IpamScopeId *string `json:"IpamScopeId,omitempty" xml:"IpamScopeId,omitempty"`
	// example:
	//
	// test
	IpamScopeName *string `json:"IpamScopeName,omitempty" xml:"IpamScopeName,omitempty"`
	// example:
	//
	// private
	IpamScopeType *string `json:"IpamScopeType,omitempty" xml:"IpamScopeType,omitempty"`
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// example:
	//
	// 1210123456******
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 2
	PoolCount *int32 `json:"PoolCount,omitempty" xml:"PoolCount,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// Created
	Status *string                                     `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   []*ListIpamScopesResponseBodyIpamScopesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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

func (s *ListIpamScopesResponseBodyIpamScopes) SetStatus(v string) *ListIpamScopesResponseBodyIpamScopes {
	s.Status = &v
	return s
}

func (s *ListIpamScopesResponseBodyIpamScopes) SetTags(v []*ListIpamScopesResponseBodyIpamScopesTags) *ListIpamScopesResponseBodyIpamScopes {
	s.Tags = v
	return s
}

type ListIpamScopesResponseBodyIpamScopesTags struct {
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	IpamIds []*string `json:"IpamIds,omitempty" xml:"IpamIds,omitempty" type:"Repeated"`
	// example:
	//
	// test
	IpamName *string `json:"IpamName,omitempty" xml:"IpamName,omitempty"`
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-aek2sermdd6****
	ResourceGroupId      *string                 `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string                 `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Tags                 []*ListIpamsRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	Ipams []*ListIpamsResponseBodyIpams `json:"Ipams,omitempty" xml:"Ipams,omitempty" type:"Repeated"`
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
	// 23CA0A0B-B0F5-5495-B355-7D9A9203A46B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 2022-07-01T02:05:23Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// test description
	IpamDescription *string `json:"IpamDescription,omitempty" xml:"IpamDescription,omitempty"`
	// example:
	//
	// ipam-ccxbnsbhew0d6t****
	IpamId *string `json:"IpamId,omitempty" xml:"IpamId,omitempty"`
	// example:
	//
	// test
	IpamName *string `json:"IpamName,omitempty" xml:"IpamName,omitempty"`
	// example:
	//
	// Created
	IpamStatus          *string   `json:"IpamStatus,omitempty" xml:"IpamStatus,omitempty"`
	OperatingRegionList []*string `json:"OperatingRegionList,omitempty" xml:"OperatingRegionList,omitempty" type:"Repeated"`
	// example:
	//
	// 1210123456******
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// ipam-scope-okoerbco6unqfr****
	PrivateDefaultScopeId *string `json:"PrivateDefaultScopeId,omitempty" xml:"PrivateDefaultScopeId,omitempty"`
	// example:
	//
	// ipam-scope-ovb76p1g1m19dr****
	PublicDefaultScopeId *string `json:"PublicDefaultScopeId,omitempty" xml:"PublicDefaultScopeId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-aek2dbprgpt****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 2
	ScopeCount *int32                            `json:"ScopeCount,omitempty" xml:"ScopeCount,omitempty"`
	Tags       []*ListIpamsResponseBodyIpamsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// IPAM
	ResourceType *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 672053AB-90C9-5693-AB96-458F137A5ED6
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
	// example:
	//
	// ipam-uq5dcfc2eqhpf4****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// IPAM
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// FinanceDept
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
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
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// IPAM
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
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
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	// example:
	//
	// false
	All          *bool   `json:"All,omitempty" xml:"All,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// IPAM
	ResourceType *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
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
	AddOperatingRegion []*string `json:"AddOperatingRegion,omitempty" xml:"AddOperatingRegion,omitempty" type:"Repeated"`
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// test description
	IpamDescription *string `json:"IpamDescription,omitempty" xml:"IpamDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ipam-ccxbnsbhew0d6t****
	IpamId *string `json:"IpamId,omitempty" xml:"IpamId,omitempty"`
	// example:
	//
	// test
	IpamName     *string `json:"IpamName,omitempty" xml:"IpamName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId              *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// example:
	//
	// 28
	AllocationDefaultCidrMask *int32 `json:"AllocationDefaultCidrMask,omitempty" xml:"AllocationDefaultCidrMask,omitempty"`
	// example:
	//
	// 32
	AllocationMaxCidrMask *int32 `json:"AllocationMaxCidrMask,omitempty" xml:"AllocationMaxCidrMask,omitempty"`
	// example:
	//
	// 8
	AllocationMinCidrMask *int32 `json:"AllocationMinCidrMask,omitempty" xml:"AllocationMinCidrMask,omitempty"`
	// example:
	//
	// true
	ClearAllocationDefaultCidrMask *bool `json:"ClearAllocationDefaultCidrMask,omitempty" xml:"ClearAllocationDefaultCidrMask,omitempty"`
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// test description
	IpamPoolDescription *string `json:"IpamPoolDescription,omitempty" xml:"IpamPoolDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ipam-pool-6rcq3tobayc20t****
	IpamPoolId *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
	// example:
	//
	// test
	IpamPoolName *string `json:"IpamPoolName,omitempty" xml:"IpamPoolName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

type UpdateIpamScopeRequest struct {
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// test description
	IpamScopeDescription *string `json:"IpamScopeDescription,omitempty" xml:"IpamScopeDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ipam-scope-glfmcyldpm8lsy****
	IpamScopeId *string `json:"IpamScopeId,omitempty" xml:"IpamScopeId,omitempty"`
	// example:
	//
	// test
	IpamScopeName *string `json:"IpamScopeName,omitempty" xml:"IpamScopeName,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
// 创建IPAM实例。
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
// 创建IPAM实例。
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

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIpamPoolId)) {
		query["SourceIpamPoolId"] = request.SourceIpamPoolId
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
	if !tea.BoolValue(util.IsUnset(request.Cidr)) {
		query["Cidr"] = request.Cidr
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.IpamPoolAllocationId)) {
		query["IpamPoolAllocationId"] = request.IpamPoolAllocationId
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
// 查询IPAM功能的开通状态。
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
// 查询IPAM功能的开通状态。
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
// 查询ipam
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
// 查询ipam
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
// 查询资源标签列表
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
// 查询资源标签列表
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
// 开通IPAM功能。
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
// 开通IPAM功能。
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
// 为资源实例绑定资源标签
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
// 为资源实例绑定资源标签
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
// 为资源解绑资源标签
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
// 为资源解绑资源标签
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
// 更新ipam
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
// 更新ipam
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
