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

type AcceptVpcPeerConnectionRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses **request ID** as the **client token**. The **request ID** may be different for each request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// *   **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	// *   **false** (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the VPC peering connection to be accepted by the accepter VPC.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the resource group.
	//
	// For more information about resource groups, see [What is a resource group?](~~94475~~)
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s AcceptVpcPeerConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s AcceptVpcPeerConnectionRequest) GoString() string {
	return s.String()
}

func (s *AcceptVpcPeerConnectionRequest) SetClientToken(v string) *AcceptVpcPeerConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *AcceptVpcPeerConnectionRequest) SetDryRun(v bool) *AcceptVpcPeerConnectionRequest {
	s.DryRun = &v
	return s
}

func (s *AcceptVpcPeerConnectionRequest) SetInstanceId(v string) *AcceptVpcPeerConnectionRequest {
	s.InstanceId = &v
	return s
}

func (s *AcceptVpcPeerConnectionRequest) SetResourceGroupId(v string) *AcceptVpcPeerConnectionRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *AcceptVpcPeerConnectionRequest) SetResourceOwnerAccount(v string) *AcceptVpcPeerConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

type AcceptVpcPeerConnectionResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AcceptVpcPeerConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AcceptVpcPeerConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *AcceptVpcPeerConnectionResponseBody) SetRequestId(v string) *AcceptVpcPeerConnectionResponseBody {
	s.RequestId = &v
	return s
}

type AcceptVpcPeerConnectionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AcceptVpcPeerConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AcceptVpcPeerConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s AcceptVpcPeerConnectionResponse) GoString() string {
	return s.String()
}

func (s *AcceptVpcPeerConnectionResponse) SetHeaders(v map[string]*string) *AcceptVpcPeerConnectionResponse {
	s.Headers = v
	return s
}

func (s *AcceptVpcPeerConnectionResponse) SetStatusCode(v int32) *AcceptVpcPeerConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *AcceptVpcPeerConnectionResponse) SetBody(v *AcceptVpcPeerConnectionResponseBody) *AcceptVpcPeerConnectionResponse {
	s.Body = v
	return s
}

type CreateVpcPeerConnectionRequest struct {
	// The ID of the Alibaba Cloud account to which the accepter VPC belongs.
	//
	// *   To create a VPC peering connection within your Alibaba Cloud account, enter the ID of your Alibaba Cloud account.
	// *   To create a VPC peering connection between your Alibaba Cloud account and another Alibaba Cloud account, enter the ID of the peer Alibaba Cloud account.
	//
	// >  If the accepter is a RAM user, set **AcceptingAliUid** to the ID of the Alibaba Cloud account that created the RAM user.
	AcceptingAliUid *int64 `json:"AcceptingAliUid,omitempty" xml:"AcceptingAliUid,omitempty"`
	// The region ID of the accepter VPC of the VPC peering connection that you want to create.
	//
	// *   To create an intra-region VPC peering connection, enter a region ID that is the same as that of the requester VPC.
	// *   To create an inter-region VPC peering connection, enter a region ID that is different from that of the requester VPC.
	AcceptingRegionId *string `json:"AcceptingRegionId,omitempty" xml:"AcceptingRegionId,omitempty"`
	// The ID of the accepter VPC.
	AcceptingVpcId *string `json:"AcceptingVpcId,omitempty" xml:"AcceptingVpcId,omitempty"`
	// The bandwidth of the VPC peering connection. Unit: Mbit/s. The value must be an integer greater than 0. Before you specify this parameter, make sure that you create an inter-region VPC peering connection.
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID** as the **client token**. The **request ID** may be different for each request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the VPC peering connection.
	//
	// The description must be 2 to 256 characters in length. The description must start with a letter but cannot start with `http://` or `https://`.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// *   **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	// *   **false** (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The name of the VPC peering connection.
	//
	// The name must be 2 to 128 characters in length, and can contain digits, underscores (\_), and hyphens (-). It must start with a letter.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region where you want to create a VPC peering connection.
	//
	// You can call the [DescribeRegions](~~36063~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// For more information about resource groups, see [Resource groups](~~94475~~).
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the requester VPC.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateVpcPeerConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcPeerConnectionRequest) GoString() string {
	return s.String()
}

func (s *CreateVpcPeerConnectionRequest) SetAcceptingAliUid(v int64) *CreateVpcPeerConnectionRequest {
	s.AcceptingAliUid = &v
	return s
}

func (s *CreateVpcPeerConnectionRequest) SetAcceptingRegionId(v string) *CreateVpcPeerConnectionRequest {
	s.AcceptingRegionId = &v
	return s
}

func (s *CreateVpcPeerConnectionRequest) SetAcceptingVpcId(v string) *CreateVpcPeerConnectionRequest {
	s.AcceptingVpcId = &v
	return s
}

func (s *CreateVpcPeerConnectionRequest) SetBandwidth(v int32) *CreateVpcPeerConnectionRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateVpcPeerConnectionRequest) SetClientToken(v string) *CreateVpcPeerConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateVpcPeerConnectionRequest) SetDescription(v string) *CreateVpcPeerConnectionRequest {
	s.Description = &v
	return s
}

func (s *CreateVpcPeerConnectionRequest) SetDryRun(v bool) *CreateVpcPeerConnectionRequest {
	s.DryRun = &v
	return s
}

func (s *CreateVpcPeerConnectionRequest) SetName(v string) *CreateVpcPeerConnectionRequest {
	s.Name = &v
	return s
}

func (s *CreateVpcPeerConnectionRequest) SetRegionId(v string) *CreateVpcPeerConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *CreateVpcPeerConnectionRequest) SetResourceGroupId(v string) *CreateVpcPeerConnectionRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateVpcPeerConnectionRequest) SetVpcId(v string) *CreateVpcPeerConnectionRequest {
	s.VpcId = &v
	return s
}

type CreateVpcPeerConnectionResponseBody struct {
	// The ID of the instance on which the VPC peering connection is created.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateVpcPeerConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcPeerConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVpcPeerConnectionResponseBody) SetInstanceId(v string) *CreateVpcPeerConnectionResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateVpcPeerConnectionResponseBody) SetRequestId(v string) *CreateVpcPeerConnectionResponseBody {
	s.RequestId = &v
	return s
}

type CreateVpcPeerConnectionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateVpcPeerConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateVpcPeerConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcPeerConnectionResponse) GoString() string {
	return s.String()
}

func (s *CreateVpcPeerConnectionResponse) SetHeaders(v map[string]*string) *CreateVpcPeerConnectionResponse {
	s.Headers = v
	return s
}

func (s *CreateVpcPeerConnectionResponse) SetStatusCode(v int32) *CreateVpcPeerConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVpcPeerConnectionResponse) SetBody(v *CreateVpcPeerConnectionResponseBody) *CreateVpcPeerConnectionResponse {
	s.Body = v
	return s
}

type DeleteVpcPeerConnectionRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not set this parameter, the system uses **RequestId** as **ClientToken**. The value of **RequestId** for each API request may be different.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to check the request without performing the operation. Valid values:
	//
	// *   **true**: checks the request without performing the operation. The system checks the required parameters, request syntax, and limits. If the request fails to pass the check, an error message is returned. If the request passes the check, the `DryRunOperation` error code is returned.
	// *   **false** (default): sends the request. If the request passes the check, an HTTP 2xx status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to forcefully delete the VPC peering connection. Valid values:
	//
	// *   **false** (default): no.
	// *   **true**: yes. If you forcefully delete the VPC peering connection, the system deletes the routes that point to the VPC peering connection from the VPC route table.
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// The ID of the VPC peering connection that you want to delete.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteVpcPeerConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpcPeerConnectionRequest) GoString() string {
	return s.String()
}

func (s *DeleteVpcPeerConnectionRequest) SetClientToken(v string) *DeleteVpcPeerConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteVpcPeerConnectionRequest) SetDryRun(v bool) *DeleteVpcPeerConnectionRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteVpcPeerConnectionRequest) SetForce(v bool) *DeleteVpcPeerConnectionRequest {
	s.Force = &v
	return s
}

func (s *DeleteVpcPeerConnectionRequest) SetInstanceId(v string) *DeleteVpcPeerConnectionRequest {
	s.InstanceId = &v
	return s
}

type DeleteVpcPeerConnectionResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVpcPeerConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpcPeerConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVpcPeerConnectionResponseBody) SetRequestId(v string) *DeleteVpcPeerConnectionResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVpcPeerConnectionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteVpcPeerConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVpcPeerConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpcPeerConnectionResponse) GoString() string {
	return s.String()
}

func (s *DeleteVpcPeerConnectionResponse) SetHeaders(v map[string]*string) *DeleteVpcPeerConnectionResponse {
	s.Headers = v
	return s
}

func (s *DeleteVpcPeerConnectionResponse) SetStatusCode(v int32) *DeleteVpcPeerConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVpcPeerConnectionResponse) SetBody(v *DeleteVpcPeerConnectionResponseBody) *DeleteVpcPeerConnectionResponse {
	s.Body = v
	return s
}

type GetVpcPeerConnectionAttributeRequest struct {
	// The ID of the VPC peering connection that you want to query.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s GetVpcPeerConnectionAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVpcPeerConnectionAttributeRequest) GoString() string {
	return s.String()
}

func (s *GetVpcPeerConnectionAttributeRequest) SetInstanceId(v string) *GetVpcPeerConnectionAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *GetVpcPeerConnectionAttributeRequest) SetResourceOwnerAccount(v string) *GetVpcPeerConnectionAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

type GetVpcPeerConnectionAttributeResponseBody struct {
	// The ID of the Alibaba Cloud account to which the accepter VPC belongs.
	AcceptingOwnerUid *int64 `json:"AcceptingOwnerUid,omitempty" xml:"AcceptingOwnerUid,omitempty"`
	// The region ID of the accepter VPC.
	AcceptingRegionId *string `json:"AcceptingRegionId,omitempty" xml:"AcceptingRegionId,omitempty"`
	// The details of the accepter VPC.
	AcceptingVpc *GetVpcPeerConnectionAttributeResponseBodyAcceptingVpc `json:"AcceptingVpc,omitempty" xml:"AcceptingVpc,omitempty" type:"Struct"`
	// The bandwidth of the VPC peering connection. Unit: Mbit/s. The value must be an integer greater than 0.
	//
	// >  If the value is set to **-1**, it indicates that no limit is imposed on the bandwidth.
	//
	// Default value:
	//
	// *   The default bandwidth of an inter-region VPC peering connection is **1024** Mbit/s.
	// *   The default bandwidth of an intra-region VPC peering connection is **-1** Mbit/s.
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The business status of the VPC peering connection. Valid values:
	//
	// *   **Normal**
	// *   **FinancialLocked**
	BizStatus *string `json:"BizStatus,omitempty" xml:"BizStatus,omitempty"`
	// The description of the VPC peering connection.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the VPC peering connection was created. The time is displayed in the `YYYY-MM-DDThh:mm:ssZ` format in UTC.
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The expiration time of the VPC peering connection.
	//
	// The expiration time is returned only when the **Status** of the VPC peering connection is **Accepting** or **Expired**. Otherwise, **null** is returned.
	GmtExpired *string `json:"GmtExpired,omitempty" xml:"GmtExpired,omitempty"`
	// The time when the VPC peering connection was modified. The time is displayed in the `YYYY-MM-DDThh:mm:ssZ` format in UTC.
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the VPC peering connection.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the VPC peering connection.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the Alibaba Cloud account to which the requester VPC belongs.
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the requester VPC.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the VPC peering connection. Valid values:
	//
	// *   **Creating**
	// *   **Accepting**
	// *   **Updating**
	// *   **Rejected**
	// *   **Expired**
	// *   **Activated**
	// *   **Deleting**
	// *   **Deleted**
	//
	// For more information about the status of VPC peering connections, see [Overview of VPC peering connections](~~418507~~).
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag list.
	Tags []*GetVpcPeerConnectionAttributeResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The details of the requester VPC.
	Vpc *GetVpcPeerConnectionAttributeResponseBodyVpc `json:"Vpc,omitempty" xml:"Vpc,omitempty" type:"Struct"`
}

func (s GetVpcPeerConnectionAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVpcPeerConnectionAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetVpcPeerConnectionAttributeResponseBody) SetAcceptingOwnerUid(v int64) *GetVpcPeerConnectionAttributeResponseBody {
	s.AcceptingOwnerUid = &v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBody) SetAcceptingRegionId(v string) *GetVpcPeerConnectionAttributeResponseBody {
	s.AcceptingRegionId = &v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBody) SetAcceptingVpc(v *GetVpcPeerConnectionAttributeResponseBodyAcceptingVpc) *GetVpcPeerConnectionAttributeResponseBody {
	s.AcceptingVpc = v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBody) SetBandwidth(v int32) *GetVpcPeerConnectionAttributeResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBody) SetBizStatus(v string) *GetVpcPeerConnectionAttributeResponseBody {
	s.BizStatus = &v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBody) SetDescription(v string) *GetVpcPeerConnectionAttributeResponseBody {
	s.Description = &v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBody) SetGmtCreate(v string) *GetVpcPeerConnectionAttributeResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBody) SetGmtExpired(v string) *GetVpcPeerConnectionAttributeResponseBody {
	s.GmtExpired = &v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBody) SetGmtModified(v string) *GetVpcPeerConnectionAttributeResponseBody {
	s.GmtModified = &v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBody) SetInstanceId(v string) *GetVpcPeerConnectionAttributeResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBody) SetName(v string) *GetVpcPeerConnectionAttributeResponseBody {
	s.Name = &v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBody) SetOwnerId(v int64) *GetVpcPeerConnectionAttributeResponseBody {
	s.OwnerId = &v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBody) SetRegionId(v string) *GetVpcPeerConnectionAttributeResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBody) SetRequestId(v string) *GetVpcPeerConnectionAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBody) SetResourceGroupId(v string) *GetVpcPeerConnectionAttributeResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBody) SetStatus(v string) *GetVpcPeerConnectionAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBody) SetTags(v []*GetVpcPeerConnectionAttributeResponseBodyTags) *GetVpcPeerConnectionAttributeResponseBody {
	s.Tags = v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBody) SetVpc(v *GetVpcPeerConnectionAttributeResponseBodyVpc) *GetVpcPeerConnectionAttributeResponseBody {
	s.Vpc = v
	return s
}

type GetVpcPeerConnectionAttributeResponseBodyAcceptingVpc struct {
	// The CIDR block of the accepter VPC.
	Ipv4Cidrs []*string `json:"Ipv4Cidrs,omitempty" xml:"Ipv4Cidrs,omitempty" type:"Repeated"`
	// The IPv6 CIDR block of the accepter VPC.
	Ipv6Cidrs []*string `json:"Ipv6Cidrs,omitempty" xml:"Ipv6Cidrs,omitempty" type:"Repeated"`
	// The ID of the accepter VPC.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetVpcPeerConnectionAttributeResponseBodyAcceptingVpc) String() string {
	return tea.Prettify(s)
}

func (s GetVpcPeerConnectionAttributeResponseBodyAcceptingVpc) GoString() string {
	return s.String()
}

func (s *GetVpcPeerConnectionAttributeResponseBodyAcceptingVpc) SetIpv4Cidrs(v []*string) *GetVpcPeerConnectionAttributeResponseBodyAcceptingVpc {
	s.Ipv4Cidrs = v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBodyAcceptingVpc) SetIpv6Cidrs(v []*string) *GetVpcPeerConnectionAttributeResponseBodyAcceptingVpc {
	s.Ipv6Cidrs = v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBodyAcceptingVpc) SetVpcId(v string) *GetVpcPeerConnectionAttributeResponseBodyAcceptingVpc {
	s.VpcId = &v
	return s
}

type GetVpcPeerConnectionAttributeResponseBodyTags struct {
	// The tag key.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetVpcPeerConnectionAttributeResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s GetVpcPeerConnectionAttributeResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetVpcPeerConnectionAttributeResponseBodyTags) SetKey(v string) *GetVpcPeerConnectionAttributeResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBodyTags) SetValue(v string) *GetVpcPeerConnectionAttributeResponseBodyTags {
	s.Value = &v
	return s
}

type GetVpcPeerConnectionAttributeResponseBodyVpc struct {
	// The CIDR block of the requester VPC.
	Ipv4Cidrs []*string `json:"Ipv4Cidrs,omitempty" xml:"Ipv4Cidrs,omitempty" type:"Repeated"`
	// The IPv6 CIDR block of the requester VPC.
	Ipv6Cidrs []*string `json:"Ipv6Cidrs,omitempty" xml:"Ipv6Cidrs,omitempty" type:"Repeated"`
	// The ID of the requester VPC.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetVpcPeerConnectionAttributeResponseBodyVpc) String() string {
	return tea.Prettify(s)
}

func (s GetVpcPeerConnectionAttributeResponseBodyVpc) GoString() string {
	return s.String()
}

func (s *GetVpcPeerConnectionAttributeResponseBodyVpc) SetIpv4Cidrs(v []*string) *GetVpcPeerConnectionAttributeResponseBodyVpc {
	s.Ipv4Cidrs = v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBodyVpc) SetIpv6Cidrs(v []*string) *GetVpcPeerConnectionAttributeResponseBodyVpc {
	s.Ipv6Cidrs = v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBodyVpc) SetVpcId(v string) *GetVpcPeerConnectionAttributeResponseBodyVpc {
	s.VpcId = &v
	return s
}

type GetVpcPeerConnectionAttributeResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetVpcPeerConnectionAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetVpcPeerConnectionAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVpcPeerConnectionAttributeResponse) GoString() string {
	return s.String()
}

func (s *GetVpcPeerConnectionAttributeResponse) SetHeaders(v map[string]*string) *GetVpcPeerConnectionAttributeResponse {
	s.Headers = v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponse) SetStatusCode(v int32) *GetVpcPeerConnectionAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponse) SetBody(v *GetVpcPeerConnectionAttributeResponseBody) *GetVpcPeerConnectionAttributeResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// The number of entries to return on each page. Valid values: **1** to **50**. Default value: **50**.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used for the next query. Valid values:
	//
	// *   If this is your first query or no subsequent query is to be sent, ignore this parameter.
	// *   If a next query is to be sent, set the value to the value of **NextToken** that is returned in the last call.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the resource.
	//
	// You can call the [DescribeRegions](~~36063~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Deprecated
	// The resource ID.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the resource. Set the value to **PeerConnection**, which specifies a VPC peering connection.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// Deprecated
	// The tags.
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
	// The key of the tag that is added to the resource. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The key can be up to 64 characters in length and can contain letters, digits, periods (.), underscores (\_), and hyphens (-). The key must start with a letter but cannot start with `aliyun` or `acs:`. The key cannot contain `http://` or `https://`.
	//
	// >  Specify at least one of the **ResourceId.N** and **Tag.N** parameters (**Tag.N.Key** and **Tag.N.Value**).
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag that is added to the resource. You can specify up to 20 tag values. It can be an empty string.
	//
	// The value can be up to 128 characters in length and can contain letters, digits, periods (.), underscores (\_), and hyphens (-). The value must start with a letter but cannot start with `aliyun` or `acs:`. The value cannot contain `http://` or `https://`.
	//
	// >  Specify at least one of the **ResourceId.N** and **Tag.N** parameters (**Tag.N.Key** and **Tag.N.Value**).
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
	// The number of entries returned per page.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used for the next query. Valid values:
	//
	// *   If **NextToken** is empty, it indicates that no next query is to be sent.
	// *   If **NextToken** is returned, the value indicates the token that is used for the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
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

func (s *ListTagResourcesResponseBody) SetMaxResults(v int32) *ListTagResourcesResponseBody {
	s.MaxResults = &v
	return s
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
	// The region of the requester VPC.
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The ID of the resource.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource. The value is set to **PeerConnection**, which indicates a VPC peering connection.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag key.
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetRegionNo(v string) *ListTagResourcesResponseBodyTagResources {
	s.RegionNo = &v
	return s
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ListVpcPeerConnectionsRequest struct {
	// The ID of the VPC peering connection that you want to query.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of entries to return per page. Valid values: **1** to **100**. Default value: **20**.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The name of the VPC peering connection that you want to query.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The token that is used for the next query. Valid values:
	//
	// *   You do not need to specify this parameter for the first request.
	// *   You must specify the token that is obtained from the previous query as the value of NextToken.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region where you want to query VPC peering connections.
	//
	// You can call the [DescribeRegions](~~36063~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// For more information about resource groups, see [What is a resource group?](~~94475~~)
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	Tags []*ListVpcPeerConnectionsRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the requester VPC or accepter VPC of the VPC peering connection that you want to query.
	VpcId []*string `json:"VpcId,omitempty" xml:"VpcId,omitempty" type:"Repeated"`
}

func (s ListVpcPeerConnectionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVpcPeerConnectionsRequest) GoString() string {
	return s.String()
}

func (s *ListVpcPeerConnectionsRequest) SetInstanceId(v string) *ListVpcPeerConnectionsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListVpcPeerConnectionsRequest) SetMaxResults(v int32) *ListVpcPeerConnectionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListVpcPeerConnectionsRequest) SetName(v string) *ListVpcPeerConnectionsRequest {
	s.Name = &v
	return s
}

func (s *ListVpcPeerConnectionsRequest) SetNextToken(v string) *ListVpcPeerConnectionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListVpcPeerConnectionsRequest) SetRegionId(v string) *ListVpcPeerConnectionsRequest {
	s.RegionId = &v
	return s
}

func (s *ListVpcPeerConnectionsRequest) SetResourceGroupId(v string) *ListVpcPeerConnectionsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVpcPeerConnectionsRequest) SetTags(v []*ListVpcPeerConnectionsRequestTags) *ListVpcPeerConnectionsRequest {
	s.Tags = v
	return s
}

func (s *ListVpcPeerConnectionsRequest) SetVpcId(v []*string) *ListVpcPeerConnectionsRequest {
	s.VpcId = v
	return s
}

type ListVpcPeerConnectionsRequestTags struct {
	// The tag key. You can specify at most 20 tag keys. It cannot be an empty string.
	//
	// The key cannot exceed 64 characters in length, and can contain digits, periods (.), underscores (\_), and hyphens (-). The key must start with a letter but cannot start with `aliyun` or `acs:`. The key cannot contain `http://` or `https://`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can specify at most 20 tag values. The tag key can be an empty string.
	//
	// The tag value cannot exceed 128 characters in length, and can contain digits, periods (.), underscores (\_), and hyphens (-). The key must start with a letter but cannot start with `aliyun` or `acs:`. The key cannot contain `http://` or `https://`.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListVpcPeerConnectionsRequestTags) String() string {
	return tea.Prettify(s)
}

func (s ListVpcPeerConnectionsRequestTags) GoString() string {
	return s.String()
}

func (s *ListVpcPeerConnectionsRequestTags) SetKey(v string) *ListVpcPeerConnectionsRequestTags {
	s.Key = &v
	return s
}

func (s *ListVpcPeerConnectionsRequestTags) SetValue(v string) *ListVpcPeerConnectionsRequestTags {
	s.Value = &v
	return s
}

type ListVpcPeerConnectionsShrinkRequest struct {
	// The ID of the VPC peering connection that you want to query.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of entries to return per page. Valid values: **1** to **100**. Default value: **20**.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The name of the VPC peering connection that you want to query.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The token that is used for the next query. Valid values:
	//
	// *   You do not need to specify this parameter for the first request.
	// *   You must specify the token that is obtained from the previous query as the value of NextToken.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region where you want to query VPC peering connections.
	//
	// You can call the [DescribeRegions](~~36063~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// For more information about resource groups, see [What is a resource group?](~~94475~~)
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	Tags []*ListVpcPeerConnectionsShrinkRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the requester VPC or accepter VPC of the VPC peering connection that you want to query.
	VpcIdShrink *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListVpcPeerConnectionsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVpcPeerConnectionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListVpcPeerConnectionsShrinkRequest) SetInstanceId(v string) *ListVpcPeerConnectionsShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ListVpcPeerConnectionsShrinkRequest) SetMaxResults(v int32) *ListVpcPeerConnectionsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListVpcPeerConnectionsShrinkRequest) SetName(v string) *ListVpcPeerConnectionsShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListVpcPeerConnectionsShrinkRequest) SetNextToken(v string) *ListVpcPeerConnectionsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListVpcPeerConnectionsShrinkRequest) SetRegionId(v string) *ListVpcPeerConnectionsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListVpcPeerConnectionsShrinkRequest) SetResourceGroupId(v string) *ListVpcPeerConnectionsShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVpcPeerConnectionsShrinkRequest) SetTags(v []*ListVpcPeerConnectionsShrinkRequestTags) *ListVpcPeerConnectionsShrinkRequest {
	s.Tags = v
	return s
}

func (s *ListVpcPeerConnectionsShrinkRequest) SetVpcIdShrink(v string) *ListVpcPeerConnectionsShrinkRequest {
	s.VpcIdShrink = &v
	return s
}

type ListVpcPeerConnectionsShrinkRequestTags struct {
	// The tag key. You can specify at most 20 tag keys. It cannot be an empty string.
	//
	// The key cannot exceed 64 characters in length, and can contain digits, periods (.), underscores (\_), and hyphens (-). The key must start with a letter but cannot start with `aliyun` or `acs:`. The key cannot contain `http://` or `https://`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can specify at most 20 tag values. The tag key can be an empty string.
	//
	// The tag value cannot exceed 128 characters in length, and can contain digits, periods (.), underscores (\_), and hyphens (-). The key must start with a letter but cannot start with `aliyun` or `acs:`. The key cannot contain `http://` or `https://`.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListVpcPeerConnectionsShrinkRequestTags) String() string {
	return tea.Prettify(s)
}

func (s ListVpcPeerConnectionsShrinkRequestTags) GoString() string {
	return s.String()
}

func (s *ListVpcPeerConnectionsShrinkRequestTags) SetKey(v string) *ListVpcPeerConnectionsShrinkRequestTags {
	s.Key = &v
	return s
}

func (s *ListVpcPeerConnectionsShrinkRequestTags) SetValue(v string) *ListVpcPeerConnectionsShrinkRequestTags {
	s.Value = &v
	return s
}

type ListVpcPeerConnectionsResponseBody struct {
	// The number of entries per page. Valid values: **1** to **100**. Default value: **20**.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used for the next query. Valid values:
	//
	// *   If no value is returned for **NextToken**, no next queries are sent.
	// *   If the value of **NextToken** is returned, the value indicates the token that is used for the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The details of the VPC peering connections.
	VpcPeerConnects []*ListVpcPeerConnectionsResponseBodyVpcPeerConnects `json:"VpcPeerConnects,omitempty" xml:"VpcPeerConnects,omitempty" type:"Repeated"`
}

func (s ListVpcPeerConnectionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVpcPeerConnectionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpcPeerConnectionsResponseBody) SetMaxResults(v int32) *ListVpcPeerConnectionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBody) SetNextToken(v string) *ListVpcPeerConnectionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBody) SetRequestId(v string) *ListVpcPeerConnectionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBody) SetTotalCount(v int32) *ListVpcPeerConnectionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBody) SetVpcPeerConnects(v []*ListVpcPeerConnectionsResponseBodyVpcPeerConnects) *ListVpcPeerConnectionsResponseBody {
	s.VpcPeerConnects = v
	return s
}

type ListVpcPeerConnectionsResponseBodyVpcPeerConnects struct {
	// The ID of the Alibaba Cloud account to which the accepter VPC belongs.
	AcceptingOwnerUid *int64 `json:"AcceptingOwnerUid,omitempty" xml:"AcceptingOwnerUid,omitempty"`
	// The region ID of the accepter VPC.
	AcceptingRegionId *string `json:"AcceptingRegionId,omitempty" xml:"AcceptingRegionId,omitempty"`
	// The details of the accepter VPC.
	AcceptingVpc *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsAcceptingVpc `json:"AcceptingVpc,omitempty" xml:"AcceptingVpc,omitempty" type:"Struct"`
	// The bandwidth of the VPC peering connection. Unit: Mbit/s. The value is an integer greater than 0.
	//
	// >  If the value is set to -1, it indicates that no limit is imposed on the bandwidth.
	//
	// Default value:
	//
	// *   The default bandwidth of an inter-region VPC peering connection is **1024** Mbit/s.
	// *   The default bandwidth of an intra-region VPC peering connection is **-1** Mbit/s.
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The business status of the VPC peering connection. Valid values:
	//
	// *   **Normal**
	// *   **FinancialLocked**
	BizStatus *string `json:"BizStatus,omitempty" xml:"BizStatus,omitempty"`
	// The description of the VPC peering connection.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the VPC peering connection was created. The time is displayed in the `YYYY-MM-DDThh:mm:ssZ` format in UTC.
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The expiration time of the VPC peering connection. The time is displayed in the `YYYY-MM-DDThh:mm:ssZ` format in UTC.
	//
	// The expiration time is returned only when the **Status** of the VPC peering connection is **Accepting** or **Expired**. Otherwise, **null** is returned.
	GmtExpired *string `json:"GmtExpired,omitempty" xml:"GmtExpired,omitempty"`
	// The time when the VPC peering connection was modified. The time is displayed in the `YYYY-MM-DDThh:mm:ssZ` format in UTC.
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the VPC peering connection.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the VPC peering connection.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the Alibaba Cloud account to which the requester VPC belongs.
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the requester VPC.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the VPC peering connection. Valid values:
	//
	// *   **Creating**
	// *   **Accepting**
	// *   **Updating**
	// *   **Rejected**
	// *   **Expired**
	// *   **Activated**
	// *   **Deleting**
	// *   **Deleted**
	//
	// For more information about the status of VPC peering connections, see [Overview of VPC peering connections](~~418507~~).
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag list.
	Tags []*ListVpcPeerConnectionsResponseBodyVpcPeerConnectsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The details of the requester VPC.
	Vpc *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsVpc `json:"Vpc,omitempty" xml:"Vpc,omitempty" type:"Struct"`
}

func (s ListVpcPeerConnectionsResponseBodyVpcPeerConnects) String() string {
	return tea.Prettify(s)
}

func (s ListVpcPeerConnectionsResponseBodyVpcPeerConnects) GoString() string {
	return s.String()
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) SetAcceptingOwnerUid(v int64) *ListVpcPeerConnectionsResponseBodyVpcPeerConnects {
	s.AcceptingOwnerUid = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) SetAcceptingRegionId(v string) *ListVpcPeerConnectionsResponseBodyVpcPeerConnects {
	s.AcceptingRegionId = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) SetAcceptingVpc(v *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsAcceptingVpc) *ListVpcPeerConnectionsResponseBodyVpcPeerConnects {
	s.AcceptingVpc = v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) SetBandwidth(v int32) *ListVpcPeerConnectionsResponseBodyVpcPeerConnects {
	s.Bandwidth = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) SetBizStatus(v string) *ListVpcPeerConnectionsResponseBodyVpcPeerConnects {
	s.BizStatus = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) SetDescription(v string) *ListVpcPeerConnectionsResponseBodyVpcPeerConnects {
	s.Description = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) SetGmtCreate(v string) *ListVpcPeerConnectionsResponseBodyVpcPeerConnects {
	s.GmtCreate = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) SetGmtExpired(v string) *ListVpcPeerConnectionsResponseBodyVpcPeerConnects {
	s.GmtExpired = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) SetGmtModified(v string) *ListVpcPeerConnectionsResponseBodyVpcPeerConnects {
	s.GmtModified = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) SetInstanceId(v string) *ListVpcPeerConnectionsResponseBodyVpcPeerConnects {
	s.InstanceId = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) SetName(v string) *ListVpcPeerConnectionsResponseBodyVpcPeerConnects {
	s.Name = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) SetOwnerId(v int64) *ListVpcPeerConnectionsResponseBodyVpcPeerConnects {
	s.OwnerId = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) SetRegionId(v string) *ListVpcPeerConnectionsResponseBodyVpcPeerConnects {
	s.RegionId = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) SetResourceGroupId(v string) *ListVpcPeerConnectionsResponseBodyVpcPeerConnects {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) SetStatus(v string) *ListVpcPeerConnectionsResponseBodyVpcPeerConnects {
	s.Status = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) SetTags(v []*ListVpcPeerConnectionsResponseBodyVpcPeerConnectsTags) *ListVpcPeerConnectionsResponseBodyVpcPeerConnects {
	s.Tags = v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) SetVpc(v *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsVpc) *ListVpcPeerConnectionsResponseBodyVpcPeerConnects {
	s.Vpc = v
	return s
}

type ListVpcPeerConnectionsResponseBodyVpcPeerConnectsAcceptingVpc struct {
	// The CIDR block of the accepter VPC.
	Ipv4Cidrs []*string `json:"Ipv4Cidrs,omitempty" xml:"Ipv4Cidrs,omitempty" type:"Repeated"`
	// The IPv6 CIDR block of the accepter VPC.
	Ipv6Cidrs []*string `json:"Ipv6Cidrs,omitempty" xml:"Ipv6Cidrs,omitempty" type:"Repeated"`
	// The ID of the accepter VPC.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListVpcPeerConnectionsResponseBodyVpcPeerConnectsAcceptingVpc) String() string {
	return tea.Prettify(s)
}

func (s ListVpcPeerConnectionsResponseBodyVpcPeerConnectsAcceptingVpc) GoString() string {
	return s.String()
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsAcceptingVpc) SetIpv4Cidrs(v []*string) *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsAcceptingVpc {
	s.Ipv4Cidrs = v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsAcceptingVpc) SetIpv6Cidrs(v []*string) *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsAcceptingVpc {
	s.Ipv6Cidrs = v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsAcceptingVpc) SetVpcId(v string) *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsAcceptingVpc {
	s.VpcId = &v
	return s
}

type ListVpcPeerConnectionsResponseBodyVpcPeerConnectsTags struct {
	// The tag key.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListVpcPeerConnectionsResponseBodyVpcPeerConnectsTags) String() string {
	return tea.Prettify(s)
}

func (s ListVpcPeerConnectionsResponseBodyVpcPeerConnectsTags) GoString() string {
	return s.String()
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsTags) SetKey(v string) *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsTags {
	s.Key = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsTags) SetValue(v string) *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsTags {
	s.Value = &v
	return s
}

type ListVpcPeerConnectionsResponseBodyVpcPeerConnectsVpc struct {
	// The CIDR block of the requester VPC.
	Ipv4Cidrs []*string `json:"Ipv4Cidrs,omitempty" xml:"Ipv4Cidrs,omitempty" type:"Repeated"`
	// The IPv6 CIDR block of the requester VPC.
	Ipv6Cidrs []*string `json:"Ipv6Cidrs,omitempty" xml:"Ipv6Cidrs,omitempty" type:"Repeated"`
	// The ID of the requester VPC.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListVpcPeerConnectionsResponseBodyVpcPeerConnectsVpc) String() string {
	return tea.Prettify(s)
}

func (s ListVpcPeerConnectionsResponseBodyVpcPeerConnectsVpc) GoString() string {
	return s.String()
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsVpc) SetIpv4Cidrs(v []*string) *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsVpc {
	s.Ipv4Cidrs = v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsVpc) SetIpv6Cidrs(v []*string) *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsVpc {
	s.Ipv6Cidrs = v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsVpc) SetVpcId(v string) *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsVpc {
	s.VpcId = &v
	return s
}

type ListVpcPeerConnectionsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListVpcPeerConnectionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVpcPeerConnectionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVpcPeerConnectionsResponse) GoString() string {
	return s.String()
}

func (s *ListVpcPeerConnectionsResponse) SetHeaders(v map[string]*string) *ListVpcPeerConnectionsResponse {
	s.Headers = v
	return s
}

func (s *ListVpcPeerConnectionsResponse) SetStatusCode(v int32) *ListVpcPeerConnectionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVpcPeerConnectionsResponse) SetBody(v *ListVpcPeerConnectionsResponseBody) *ListVpcPeerConnectionsResponse {
	s.Body = v
	return s
}

type ModifyVpcPeerConnectionRequest struct {
	// The new bandwidth of the VPC peering connection. Unit: Mbit/s. The value must be an integer greater than 0.
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system uses **RequestId** as **ClientToken**. **RequestId** may be different for each API request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The new description of the VPC peering connection.
	//
	// The description must be 1 to 256 characters in length, and cannot start with `http://` or `https://`.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to only precheck the request. Valid values:
	//
	// *   **true**: checks the request without performing the operation. The system prechecks the required parameters, request syntax, and limits. If the request fails the precheck, an error message is returned. If the request passes the check, the `DryRunOperation` error code is returned.
	// *   **false** (default): sends the request. If the request passes the precheck, an HTTP 2xx status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the VPC peering connection whose name or description you want to modify.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The new name of the VPC peering connection.
	//
	// The name must be 1 to 128 characters in length, and cannot start with `http://` or `https://`.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ModifyVpcPeerConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcPeerConnectionRequest) GoString() string {
	return s.String()
}

func (s *ModifyVpcPeerConnectionRequest) SetBandwidth(v int32) *ModifyVpcPeerConnectionRequest {
	s.Bandwidth = &v
	return s
}

func (s *ModifyVpcPeerConnectionRequest) SetClientToken(v string) *ModifyVpcPeerConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyVpcPeerConnectionRequest) SetDescription(v string) *ModifyVpcPeerConnectionRequest {
	s.Description = &v
	return s
}

func (s *ModifyVpcPeerConnectionRequest) SetDryRun(v bool) *ModifyVpcPeerConnectionRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyVpcPeerConnectionRequest) SetInstanceId(v string) *ModifyVpcPeerConnectionRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyVpcPeerConnectionRequest) SetName(v string) *ModifyVpcPeerConnectionRequest {
	s.Name = &v
	return s
}

type ModifyVpcPeerConnectionResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVpcPeerConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcPeerConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVpcPeerConnectionResponseBody) SetRequestId(v string) *ModifyVpcPeerConnectionResponseBody {
	s.RequestId = &v
	return s
}

type ModifyVpcPeerConnectionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyVpcPeerConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyVpcPeerConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcPeerConnectionResponse) GoString() string {
	return s.String()
}

func (s *ModifyVpcPeerConnectionResponse) SetHeaders(v map[string]*string) *ModifyVpcPeerConnectionResponse {
	s.Headers = v
	return s
}

func (s *ModifyVpcPeerConnectionResponse) SetStatusCode(v int32) *ModifyVpcPeerConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVpcPeerConnectionResponse) SetBody(v *ModifyVpcPeerConnectionResponseBody) *ModifyVpcPeerConnectionResponse {
	s.Body = v
	return s
}

type MoveResourceGroupRequest struct {
	// The ID of the new resource group.
	//
	// >  You can use resource groups to manage resources owned by your Alibaba Cloud account. Resource groups simplify the resource and permission management of your Alibaba Cloud account. For more information, see [What is resource management?](~~94475~~).
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	// The ID of the region to which the resource belongs.
	//
	// You can call the [DescribeRegions](~~36063~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the VPC peering connection.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type. Set the value to **PeerConnection**, which specifies a VPC peering connection.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s MoveResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupRequest) SetNewResourceGroupId(v string) *MoveResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetRegionId(v string) *MoveResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceId(v string) *MoveResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceType(v string) *MoveResourceGroupRequest {
	s.ResourceType = &v
	return s
}

type MoveResourceGroupResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. Valid values:
	//
	// *   **true**
	// *   **false**
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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

func (s *MoveResourceGroupResponseBody) SetSuccess(v bool) *MoveResourceGroupResponseBody {
	s.Success = &v
	return s
}

type MoveResourceGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *MoveResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *MoveResourceGroupResponse) SetStatusCode(v int32) *MoveResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveResourceGroupResponse) SetBody(v *MoveResourceGroupResponseBody) *MoveResourceGroupResponse {
	s.Body = v
	return s
}

type RejectVpcPeerConnectionRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not set this parameter, the system uses **RequestId** as **ClientToken**. The value of **RequestId** for each API request may be different.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to check the request without performing the operation. Valid values:
	//
	// *   **true**: checks the request without performing the operation. The system checks the required parameters, request syntax, and limits. If the request fails to pass the check, an error message is returned. If the request passes the check, the `DryRunOperation` error code is returned.
	// *   **false** (default): sends the request. If the request passes the check, an HTTP 2xx status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the VPC peering connection to be rejected by the acceptor VPC.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s RejectVpcPeerConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s RejectVpcPeerConnectionRequest) GoString() string {
	return s.String()
}

func (s *RejectVpcPeerConnectionRequest) SetClientToken(v string) *RejectVpcPeerConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *RejectVpcPeerConnectionRequest) SetDryRun(v bool) *RejectVpcPeerConnectionRequest {
	s.DryRun = &v
	return s
}

func (s *RejectVpcPeerConnectionRequest) SetInstanceId(v string) *RejectVpcPeerConnectionRequest {
	s.InstanceId = &v
	return s
}

func (s *RejectVpcPeerConnectionRequest) SetResourceOwnerAccount(v string) *RejectVpcPeerConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

type RejectVpcPeerConnectionResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RejectVpcPeerConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RejectVpcPeerConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *RejectVpcPeerConnectionResponseBody) SetRequestId(v string) *RejectVpcPeerConnectionResponseBody {
	s.RequestId = &v
	return s
}

type RejectVpcPeerConnectionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RejectVpcPeerConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RejectVpcPeerConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s RejectVpcPeerConnectionResponse) GoString() string {
	return s.String()
}

func (s *RejectVpcPeerConnectionResponse) SetHeaders(v map[string]*string) *RejectVpcPeerConnectionResponse {
	s.Headers = v
	return s
}

func (s *RejectVpcPeerConnectionResponse) SetStatusCode(v int32) *RejectVpcPeerConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *RejectVpcPeerConnectionResponse) SetBody(v *RejectVpcPeerConnectionResponseBody) *RejectVpcPeerConnectionResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among all requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, **ClientToken** is set to the value of **RequestId**. The value of **RequestId** for each API request may be different.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID of the resource to which you want to create and add tags.
	//
	// You can call the [DescribeRegions](~~36063~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of resources.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the resource. Set the value to **PeerConnection**, which specifies a VPC peering connection.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags.
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
	// The key of the tag. You must enter at least one tag key and at most 20 tag keys. The tag key cannot be an empty string.
	//
	// The key cannot exceed 64 characters in length, and can contain digits, periods (.), underscores (\_), and hyphens (-). The key must start with a letter but cannot start with `aliyun` or `acs:`. The key cannot contain `http://` or `https://`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag. You must enter at least one tag value and at most 20 tag values. It can be an empty string.
	//
	// The tag value cannot exceed 128 characters in length, and can contain digits, periods (.), underscores (\_), and hyphens (-). It must start with a letter but cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
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
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. Valid values:
	//
	// *   **true**: yes
	// *   **false**: no
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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

func (s *TagResourcesResponseBody) SetSuccess(v bool) *TagResourcesResponseBody {
	s.Success = &v
	return s
}

type TagResourcesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type UnTagResourcesRequest struct {
	// Specifies whether to remove all tags from the specified resource. Valid values:
	//
	// *   **true**: removes all tags from the specified resource.
	// *   **false**: does not remove all tags from the specified resource. This is the default value.
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among all requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, **ClientToken** is set to the value of **RequestId**. The value of **RequestId** for each API request may be different.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID of the tag.
	//
	// You can call the [DescribeRegions](~~36063~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of resources.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the resource. Set the value to **PeerConnection**, which specifies a VPC peering connection.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags.
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UnTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UnTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UnTagResourcesRequest) SetAll(v bool) *UnTagResourcesRequest {
	s.All = &v
	return s
}

func (s *UnTagResourcesRequest) SetClientToken(v string) *UnTagResourcesRequest {
	s.ClientToken = &v
	return s
}

func (s *UnTagResourcesRequest) SetRegionId(v string) *UnTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UnTagResourcesRequest) SetResourceId(v []*string) *UnTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UnTagResourcesRequest) SetResourceType(v string) *UnTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UnTagResourcesRequest) SetTagKey(v []*string) *UnTagResourcesRequest {
	s.TagKey = v
	return s
}

type UnTagResourcesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the tags are removed. Valid values:
	//
	// *   **true**: yes
	// *   **false**: no
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UnTagResourcesResponseBody) SetRequestId(v string) *UnTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnTagResourcesResponseBody) SetSuccess(v bool) *UnTagResourcesResponseBody {
	s.Success = &v
	return s
}

type UnTagResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UnTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UnTagResourcesResponse) SetHeaders(v map[string]*string) *UnTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UnTagResourcesResponse) SetStatusCode(v int32) *UnTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UnTagResourcesResponse) SetBody(v *UnTagResourcesResponseBody) *UnTagResourcesResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("vpcpeer"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

/**
 * *   For a cross-account VPC peering connection, the connection is activated only after the accepter VPC accepts the connection request.
 * *   **AcceptVpcPeerConnection** is an asynchronous operation. After a request is sent, the system returns a **request ID** and runs the operation in the background. You can call the [GetVpcPeerConnectionAttribute](~~426100~~) operation to query the status of the task.
 *     *   If a VPC peering connection is in the **Updating** state, the VPC peering connection is being activated.
 *     *   If a VPC peering connection is in the **Activated** state, the VPC peering connection is activated.
 * *   You cannot repeatedly call the **AcceptVpcPeerConnection** operation within a specific period of time.
 *
 * @param request AcceptVpcPeerConnectionRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return AcceptVpcPeerConnectionResponse
 */
func (client *Client) AcceptVpcPeerConnectionWithOptions(request *AcceptVpcPeerConnectionRequest, runtime *util.RuntimeOptions) (_result *AcceptVpcPeerConnectionResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		body["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AcceptVpcPeerConnection"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AcceptVpcPeerConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   For a cross-account VPC peering connection, the connection is activated only after the accepter VPC accepts the connection request.
 * *   **AcceptVpcPeerConnection** is an asynchronous operation. After a request is sent, the system returns a **request ID** and runs the operation in the background. You can call the [GetVpcPeerConnectionAttribute](~~426100~~) operation to query the status of the task.
 *     *   If a VPC peering connection is in the **Updating** state, the VPC peering connection is being activated.
 *     *   If a VPC peering connection is in the **Activated** state, the VPC peering connection is activated.
 * *   You cannot repeatedly call the **AcceptVpcPeerConnection** operation within a specific period of time.
 *
 * @param request AcceptVpcPeerConnectionRequest
 * @return AcceptVpcPeerConnectionResponse
 */
func (client *Client) AcceptVpcPeerConnection(request *AcceptVpcPeerConnectionRequest) (_result *AcceptVpcPeerConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AcceptVpcPeerConnectionResponse{}
	_body, _err := client.AcceptVpcPeerConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Before you create a VPC peering connection, make sure that the following requirements are met:
 * *   Cloud Data Transfer (CDT) is activated to manage the billing of intra-border data transfers. To activate CDT, call the [OpenCdtService](~~337842~~) operation.
 * *   **CreateVpcPeerConnection** is an asynchronous operation. After a request is sent, the system returns a request ID and an **instance ID** and runs the task in the background. You can call the [GetVpcPeerConnectionAttribute](~~426095~~) operation to query the status of the task.
 *     *   If a VPC peering connection is in the **Creating** state, the VPC peering connection is being created.
 *     *   If a VPC peering connection is in the **Activated** state, the VPC peering connection is created.
 *     *   If a VPC peering connection is in the **Accepting** state, the VPC peering connection is created across accounts and the accepter is accepting the VPC peering connection.
 * *   You cannot repeatedly call the **CreateVpcPeerConnection** operation within a specific period of time.
 *
 * @param request CreateVpcPeerConnectionRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateVpcPeerConnectionResponse
 */
func (client *Client) CreateVpcPeerConnectionWithOptions(request *CreateVpcPeerConnectionRequest, runtime *util.RuntimeOptions) (_result *CreateVpcPeerConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptingAliUid)) {
		body["AcceptingAliUid"] = request.AcceptingAliUid
	}

	if !tea.BoolValue(util.IsUnset(request.AcceptingRegionId)) {
		body["AcceptingRegionId"] = request.AcceptingRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.AcceptingVpcId)) {
		body["AcceptingVpcId"] = request.AcceptingVpcId
	}

	if !tea.BoolValue(util.IsUnset(request.Bandwidth)) {
		body["Bandwidth"] = request.Bandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		body["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateVpcPeerConnection"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateVpcPeerConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Before you create a VPC peering connection, make sure that the following requirements are met:
 * *   Cloud Data Transfer (CDT) is activated to manage the billing of intra-border data transfers. To activate CDT, call the [OpenCdtService](~~337842~~) operation.
 * *   **CreateVpcPeerConnection** is an asynchronous operation. After a request is sent, the system returns a request ID and an **instance ID** and runs the task in the background. You can call the [GetVpcPeerConnectionAttribute](~~426095~~) operation to query the status of the task.
 *     *   If a VPC peering connection is in the **Creating** state, the VPC peering connection is being created.
 *     *   If a VPC peering connection is in the **Activated** state, the VPC peering connection is created.
 *     *   If a VPC peering connection is in the **Accepting** state, the VPC peering connection is created across accounts and the accepter is accepting the VPC peering connection.
 * *   You cannot repeatedly call the **CreateVpcPeerConnection** operation within a specific period of time.
 *
 * @param request CreateVpcPeerConnectionRequest
 * @return CreateVpcPeerConnectionResponse
 */
func (client *Client) CreateVpcPeerConnection(request *CreateVpcPeerConnectionRequest) (_result *CreateVpcPeerConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVpcPeerConnectionResponse{}
	_body, _err := client.CreateVpcPeerConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   You can delete VPC peering connections. After you delete a VPC peering connection, your service is affected. Proceed with caution.
 *     *   If you forcefully delete a VPC peering connection, the system deletes the routes that point to the VPC peering connection from the VPC route table.
 *     *   If you do not forcefully delete a VPC peering connection, the system does not delete these routes. You must manually delete them.
 * *   The **DeleteVpcPeerConnection** operation is asynchronous. After you send a request, the system returns **RequestId**, but the operation is still being performed in the background. You can call the [GetVpcPeerConnectionAttribute](~~426100~~) operation to query the status of a VPC peering connection.
 *     *   If a VPC peering connection is in the **Deleting** state, it is being deleted.
 *     *   If a VPC peering connection is in the **Deleted** state, it is deleted.
 * *   You cannot repeatedly call the **DeleteVpcPeerConnection** operation for the same VPC peering connection within the specified period of time.
 *
 * @param request DeleteVpcPeerConnectionRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteVpcPeerConnectionResponse
 */
func (client *Client) DeleteVpcPeerConnectionWithOptions(request *DeleteVpcPeerConnectionRequest, runtime *util.RuntimeOptions) (_result *DeleteVpcPeerConnectionResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Force)) {
		body["Force"] = request.Force
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVpcPeerConnection"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVpcPeerConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   You can delete VPC peering connections. After you delete a VPC peering connection, your service is affected. Proceed with caution.
 *     *   If you forcefully delete a VPC peering connection, the system deletes the routes that point to the VPC peering connection from the VPC route table.
 *     *   If you do not forcefully delete a VPC peering connection, the system does not delete these routes. You must manually delete them.
 * *   The **DeleteVpcPeerConnection** operation is asynchronous. After you send a request, the system returns **RequestId**, but the operation is still being performed in the background. You can call the [GetVpcPeerConnectionAttribute](~~426100~~) operation to query the status of a VPC peering connection.
 *     *   If a VPC peering connection is in the **Deleting** state, it is being deleted.
 *     *   If a VPC peering connection is in the **Deleted** state, it is deleted.
 * *   You cannot repeatedly call the **DeleteVpcPeerConnection** operation for the same VPC peering connection within the specified period of time.
 *
 * @param request DeleteVpcPeerConnectionRequest
 * @return DeleteVpcPeerConnectionResponse
 */
func (client *Client) DeleteVpcPeerConnection(request *DeleteVpcPeerConnectionRequest) (_result *DeleteVpcPeerConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVpcPeerConnectionResponse{}
	_body, _err := client.DeleteVpcPeerConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetVpcPeerConnectionAttributeWithOptions(request *GetVpcPeerConnectionAttributeRequest, runtime *util.RuntimeOptions) (_result *GetVpcPeerConnectionAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		body["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetVpcPeerConnectionAttribute"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetVpcPeerConnectionAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetVpcPeerConnectionAttribute(request *GetVpcPeerConnectionAttributeRequest) (_result *GetVpcPeerConnectionAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVpcPeerConnectionAttributeResponse{}
	_body, _err := client.GetVpcPeerConnectionAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   Set **ResourceId.N** or **Tag.N** that consists of **Tag.N.Key** and **Tag.N.Value** in the request to specify the object to be queried.
 * *   **Tag.N** is a resource tag that consists of a key-value pair. If you set only **Tag.N.Key**, all tag values that are associated with the specified key are returned. If you set only **Tag.N.Value**, an error message is returned.
 * *   If you set **Tag.N** and **ResourceId.N** to filter tags, **ResourceId.N** must match all specified key-value pairs.
 * *   If you specify multiple key-value pairs, resources that contain these key-value pairs are returned.
 *
 * @param request ListTagResourcesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListTagResourcesResponse
 */
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
		Version:     tea.String("2022-01-01"),
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

/**
 * *   Set **ResourceId.N** or **Tag.N** that consists of **Tag.N.Key** and **Tag.N.Value** in the request to specify the object to be queried.
 * *   **Tag.N** is a resource tag that consists of a key-value pair. If you set only **Tag.N.Key**, all tag values that are associated with the specified key are returned. If you set only **Tag.N.Value**, an error message is returned.
 * *   If you set **Tag.N** and **ResourceId.N** to filter tags, **ResourceId.N** must match all specified key-value pairs.
 * *   If you specify multiple key-value pairs, resources that contain these key-value pairs are returned.
 *
 * @param request ListTagResourcesRequest
 * @return ListTagResourcesResponse
 */
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

func (client *Client) ListVpcPeerConnectionsWithOptions(tmpReq *ListVpcPeerConnectionsRequest, runtime *util.RuntimeOptions) (_result *ListVpcPeerConnectionsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListVpcPeerConnectionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.VpcId)) {
		request.VpcIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VpcId, tea.String("VpcId"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcIdShrink)) {
		body["VpcId"] = request.VpcIdShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListVpcPeerConnections"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListVpcPeerConnectionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVpcPeerConnections(request *ListVpcPeerConnectionsRequest) (_result *ListVpcPeerConnectionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVpcPeerConnectionsResponse{}
	_body, _err := client.ListVpcPeerConnectionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   The **ModifyVpcPeerConnection** operation is asynchronous. After you send a request, the system returns **RequestId**, but the operation is still being performed in the background. You can call the [GetVpcPeerConnectionAttribute](~~426100~~) operation to query the status of a VPC peering connection.
 *     *   If a VPC peering connection is in the **Updating** state, the VPC peering connection is being modified.
 *     *   If a VPC peering connection is in the **Activated** state, the VPC peering connection is modified.
 * *   You cannot repeatedly call the **ModifyVpcPeerConnection** operation for the same VPC peering connection within the specified period of time.
 *
 * @param request ModifyVpcPeerConnectionRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyVpcPeerConnectionResponse
 */
func (client *Client) ModifyVpcPeerConnectionWithOptions(request *ModifyVpcPeerConnectionRequest, runtime *util.RuntimeOptions) (_result *ModifyVpcPeerConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bandwidth)) {
		body["Bandwidth"] = request.Bandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyVpcPeerConnection"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyVpcPeerConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   The **ModifyVpcPeerConnection** operation is asynchronous. After you send a request, the system returns **RequestId**, but the operation is still being performed in the background. You can call the [GetVpcPeerConnectionAttribute](~~426100~~) operation to query the status of a VPC peering connection.
 *     *   If a VPC peering connection is in the **Updating** state, the VPC peering connection is being modified.
 *     *   If a VPC peering connection is in the **Activated** state, the VPC peering connection is modified.
 * *   You cannot repeatedly call the **ModifyVpcPeerConnection** operation for the same VPC peering connection within the specified period of time.
 *
 * @param request ModifyVpcPeerConnectionRequest
 * @return ModifyVpcPeerConnectionResponse
 */
func (client *Client) ModifyVpcPeerConnection(request *ModifyVpcPeerConnectionRequest) (_result *ModifyVpcPeerConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyVpcPeerConnectionResponse{}
	_body, _err := client.ModifyVpcPeerConnectionWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.NewResourceGroupId)) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("MoveResourceGroup"),
		Version:     tea.String("2022-01-01"),
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

/**
 * *   An acceptor VPC can reject a connection request from the requester VPC of a cross-account VPC peering connection. After the connection request is rejected, the VPC peering connection enters the **Rejected** state.
 * *   You cannot repeatedly call the **RejectVpcPeerConnection** operation for the same VPC peering connection within the specified period of time.
 *
 * @param request RejectVpcPeerConnectionRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return RejectVpcPeerConnectionResponse
 */
func (client *Client) RejectVpcPeerConnectionWithOptions(request *RejectVpcPeerConnectionRequest, runtime *util.RuntimeOptions) (_result *RejectVpcPeerConnectionResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		body["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RejectVpcPeerConnection"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RejectVpcPeerConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   An acceptor VPC can reject a connection request from the requester VPC of a cross-account VPC peering connection. After the connection request is rejected, the VPC peering connection enters the **Rejected** state.
 * *   You cannot repeatedly call the **RejectVpcPeerConnection** operation for the same VPC peering connection within the specified period of time.
 *
 * @param request RejectVpcPeerConnectionRequest
 * @return RejectVpcPeerConnectionResponse
 */
func (client *Client) RejectVpcPeerConnection(request *RejectVpcPeerConnectionRequest) (_result *RejectVpcPeerConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RejectVpcPeerConnectionResponse{}
	_body, _err := client.RejectVpcPeerConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Tags are used to classify instances. Each tag consists of a key-value pair. Before you use tags, take note of the following limits:
 * *   The keys of tags that are added to the same instance must be unique.
 * *   You cannot create tags without adding them to instances. All tags must be added to instances.
 * *   Tag information is not shared across regions.
 *     For example, you cannot view the tags that are created in the China (Hangzhou) region from the China (Shanghai) region.
 * *   For the same account and region, tags added to different VPC peering connections are shared.
 *     For example, if a tag is added to a VPC peering connection, the tag can also be added to other VPC peering connections within the same account and region. You can modify the key and the value of a tag or remove a tag from an instance. After you delete an instance, all tags that are added to the instance are deleted.
 * *   You can add up to 20 tags to each instance. Before you add a tag to an instance, the system automatically checks the number of existing tags. An error message is returned if the maximum number of tags is reached.
 *
 * @param request TagResourcesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return TagResourcesResponse
 */
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
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
		Action:      tea.String("TagResources"),
		Version:     tea.String("2022-01-01"),
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

/**
 * Tags are used to classify instances. Each tag consists of a key-value pair. Before you use tags, take note of the following limits:
 * *   The keys of tags that are added to the same instance must be unique.
 * *   You cannot create tags without adding them to instances. All tags must be added to instances.
 * *   Tag information is not shared across regions.
 *     For example, you cannot view the tags that are created in the China (Hangzhou) region from the China (Shanghai) region.
 * *   For the same account and region, tags added to different VPC peering connections are shared.
 *     For example, if a tag is added to a VPC peering connection, the tag can also be added to other VPC peering connections within the same account and region. You can modify the key and the value of a tag or remove a tag from an instance. After you delete an instance, all tags that are added to the instance are deleted.
 * *   You can add up to 20 tags to each instance. Before you add a tag to an instance, the system automatically checks the number of existing tags. An error message is returned if the maximum number of tags is reached.
 *
 * @param request TagResourcesRequest
 * @return TagResourcesResponse
 */
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

func (client *Client) UnTagResourcesWithOptions(request *UnTagResourcesRequest, runtime *util.RuntimeOptions) (_result *UnTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
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

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnTagResources"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnTagResources(request *UnTagResourcesRequest) (_result *UnTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnTagResourcesResponse{}
	_body, _err := client.UnTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
