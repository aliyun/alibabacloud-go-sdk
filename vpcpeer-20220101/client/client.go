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
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun               *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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

func (s *AcceptVpcPeerConnectionRequest) SetResourceOwnerAccount(v string) *AcceptVpcPeerConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

type AcceptVpcPeerConnectionResponseBody struct {
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
	AcceptingAliUid   *int64  `json:"AcceptingAliUid,omitempty" xml:"AcceptingAliUid,omitempty"`
	AcceptingRegionId *string `json:"AcceptingRegionId,omitempty" xml:"AcceptingRegionId,omitempty"`
	AcceptingVpcId    *string `json:"AcceptingVpcId,omitempty" xml:"AcceptingVpcId,omitempty"`
	ClientToken       *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DryRun            *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VpcId             *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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

func (s *CreateVpcPeerConnectionRequest) SetVpcId(v string) *CreateVpcPeerConnectionRequest {
	s.VpcId = &v
	return s
}

type CreateVpcPeerConnectionResponseBody struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	Force       *bool   `json:"Force,omitempty" xml:"Force,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	AcceptingOwnerUid *int64                                                 `json:"AcceptingOwnerUid,omitempty" xml:"AcceptingOwnerUid,omitempty"`
	AcceptingRegionId *string                                                `json:"AcceptingRegionId,omitempty" xml:"AcceptingRegionId,omitempty"`
	AcceptingVpc      *GetVpcPeerConnectionAttributeResponseBodyAcceptingVpc `json:"AcceptingVpc,omitempty" xml:"AcceptingVpc,omitempty" type:"Struct"`
	Bandwidth         *int32                                                 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	BizStatus         *string                                                `json:"BizStatus,omitempty" xml:"BizStatus,omitempty"`
	Description       *string                                                `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate         *string                                                `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtExpired        *string                                                `json:"GmtExpired,omitempty" xml:"GmtExpired,omitempty"`
	GmtModified       *string                                                `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	InstanceId        *string                                                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name              *string                                                `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId           *int64                                                 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId          *string                                                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId         *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status            *string                                                `json:"Status,omitempty" xml:"Status,omitempty"`
	Vpc               *GetVpcPeerConnectionAttributeResponseBodyVpc          `json:"Vpc,omitempty" xml:"Vpc,omitempty" type:"Struct"`
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

func (s *GetVpcPeerConnectionAttributeResponseBody) SetStatus(v string) *GetVpcPeerConnectionAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *GetVpcPeerConnectionAttributeResponseBody) SetVpc(v *GetVpcPeerConnectionAttributeResponseBodyVpc) *GetVpcPeerConnectionAttributeResponseBody {
	s.Vpc = v
	return s
}

type GetVpcPeerConnectionAttributeResponseBodyAcceptingVpc struct {
	Ipv4Cidrs []*string `json:"Ipv4Cidrs,omitempty" xml:"Ipv4Cidrs,omitempty" type:"Repeated"`
	Ipv6Cidrs []*string `json:"Ipv6Cidrs,omitempty" xml:"Ipv6Cidrs,omitempty" type:"Repeated"`
	VpcId     *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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

type GetVpcPeerConnectionAttributeResponseBodyVpc struct {
	Ipv4Cidrs []*string `json:"Ipv4Cidrs,omitempty" xml:"Ipv4Cidrs,omitempty" type:"Repeated"`
	Ipv6Cidrs []*string `json:"Ipv6Cidrs,omitempty" xml:"Ipv6Cidrs,omitempty" type:"Repeated"`
	VpcId     *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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

type ListVpcPeerConnectionsRequest struct {
	InstanceId *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxResults *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Name       *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	NextToken  *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VpcId      []*string `json:"VpcId,omitempty" xml:"VpcId,omitempty" type:"Repeated"`
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

func (s *ListVpcPeerConnectionsRequest) SetVpcId(v []*string) *ListVpcPeerConnectionsRequest {
	s.VpcId = v
	return s
}

type ListVpcPeerConnectionsShrinkRequest struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxResults  *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *ListVpcPeerConnectionsShrinkRequest) SetVpcIdShrink(v string) *ListVpcPeerConnectionsShrinkRequest {
	s.VpcIdShrink = &v
	return s
}

type ListVpcPeerConnectionsResponseBody struct {
	MaxResults      *int32                                               `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken       *string                                              `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId       *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount      *int32                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	AcceptingOwnerUid *int32                                                         `json:"AcceptingOwnerUid,omitempty" xml:"AcceptingOwnerUid,omitempty"`
	AcceptingRegionId *string                                                        `json:"AcceptingRegionId,omitempty" xml:"AcceptingRegionId,omitempty"`
	AcceptingVpc      *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsAcceptingVpc `json:"AcceptingVpc,omitempty" xml:"AcceptingVpc,omitempty" type:"Struct"`
	Bandwidth         *int32                                                         `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	BizStatus         *string                                                        `json:"BizStatus,omitempty" xml:"BizStatus,omitempty"`
	Description       *string                                                        `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate         *string                                                        `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtExpired        *string                                                        `json:"GmtExpired,omitempty" xml:"GmtExpired,omitempty"`
	GmtModified       *string                                                        `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	InstanceId        *string                                                        `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name              *string                                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId           *int32                                                         `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId          *string                                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status            *string                                                        `json:"Status,omitempty" xml:"Status,omitempty"`
	Vpc               *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsVpc          `json:"Vpc,omitempty" xml:"Vpc,omitempty" type:"Struct"`
}

func (s ListVpcPeerConnectionsResponseBodyVpcPeerConnects) String() string {
	return tea.Prettify(s)
}

func (s ListVpcPeerConnectionsResponseBodyVpcPeerConnects) GoString() string {
	return s.String()
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) SetAcceptingOwnerUid(v int32) *ListVpcPeerConnectionsResponseBodyVpcPeerConnects {
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

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) SetOwnerId(v int32) *ListVpcPeerConnectionsResponseBodyVpcPeerConnects {
	s.OwnerId = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) SetRegionId(v string) *ListVpcPeerConnectionsResponseBodyVpcPeerConnects {
	s.RegionId = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) SetStatus(v string) *ListVpcPeerConnectionsResponseBodyVpcPeerConnects {
	s.Status = &v
	return s
}

func (s *ListVpcPeerConnectionsResponseBodyVpcPeerConnects) SetVpc(v *ListVpcPeerConnectionsResponseBodyVpcPeerConnectsVpc) *ListVpcPeerConnectionsResponseBodyVpcPeerConnects {
	s.Vpc = v
	return s
}

type ListVpcPeerConnectionsResponseBodyVpcPeerConnectsAcceptingVpc struct {
	Ipv4Cidrs []*string `json:"Ipv4Cidrs,omitempty" xml:"Ipv4Cidrs,omitempty" type:"Repeated"`
	Ipv6Cidrs []*string `json:"Ipv6Cidrs,omitempty" xml:"Ipv6Cidrs,omitempty" type:"Repeated"`
	VpcId     *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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

type ListVpcPeerConnectionsResponseBodyVpcPeerConnectsVpc struct {
	Ipv4Cidrs []*string `json:"Ipv4Cidrs,omitempty" xml:"Ipv4Cidrs,omitempty" type:"Repeated"`
	Ipv6Cidrs []*string `json:"Ipv6Cidrs,omitempty" xml:"Ipv6Cidrs,omitempty" type:"Repeated"`
	VpcId     *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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
	Bandwidth   *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DryRun      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
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

type RejectVpcPeerConnectionRequest struct {
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun               *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
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
		Body: openapiutil.ParseToMap(body),
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
