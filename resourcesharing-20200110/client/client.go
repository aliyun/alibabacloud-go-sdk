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

type AcceptResourceShareInvitationRequest struct {
	ResourceShareInvitationId *string `json:"ResourceShareInvitationId,omitempty" xml:"ResourceShareInvitationId,omitempty"`
}

func (s AcceptResourceShareInvitationRequest) String() string {
	return tea.Prettify(s)
}

func (s AcceptResourceShareInvitationRequest) GoString() string {
	return s.String()
}

func (s *AcceptResourceShareInvitationRequest) SetResourceShareInvitationId(v string) *AcceptResourceShareInvitationRequest {
	s.ResourceShareInvitationId = &v
	return s
}

type AcceptResourceShareInvitationResponseBody struct {
	RequestId               *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceShareInvitation *AcceptResourceShareInvitationResponseBodyResourceShareInvitation `json:"ResourceShareInvitation,omitempty" xml:"ResourceShareInvitation,omitempty" type:"Struct"`
}

func (s AcceptResourceShareInvitationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AcceptResourceShareInvitationResponseBody) GoString() string {
	return s.String()
}

func (s *AcceptResourceShareInvitationResponseBody) SetRequestId(v string) *AcceptResourceShareInvitationResponseBody {
	s.RequestId = &v
	return s
}

func (s *AcceptResourceShareInvitationResponseBody) SetResourceShareInvitation(v *AcceptResourceShareInvitationResponseBodyResourceShareInvitation) *AcceptResourceShareInvitationResponseBody {
	s.ResourceShareInvitation = v
	return s
}

type AcceptResourceShareInvitationResponseBodyResourceShareInvitation struct {
	CreateTime                *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ReceiverAccountId         *string `json:"ReceiverAccountId,omitempty" xml:"ReceiverAccountId,omitempty"`
	ResourceShareId           *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
	ResourceShareInvitationId *string `json:"ResourceShareInvitationId,omitempty" xml:"ResourceShareInvitationId,omitempty"`
	ResourceShareName         *string `json:"ResourceShareName,omitempty" xml:"ResourceShareName,omitempty"`
	SenderAccountId           *string `json:"SenderAccountId,omitempty" xml:"SenderAccountId,omitempty"`
	Status                    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s AcceptResourceShareInvitationResponseBodyResourceShareInvitation) String() string {
	return tea.Prettify(s)
}

func (s AcceptResourceShareInvitationResponseBodyResourceShareInvitation) GoString() string {
	return s.String()
}

func (s *AcceptResourceShareInvitationResponseBodyResourceShareInvitation) SetCreateTime(v string) *AcceptResourceShareInvitationResponseBodyResourceShareInvitation {
	s.CreateTime = &v
	return s
}

func (s *AcceptResourceShareInvitationResponseBodyResourceShareInvitation) SetReceiverAccountId(v string) *AcceptResourceShareInvitationResponseBodyResourceShareInvitation {
	s.ReceiverAccountId = &v
	return s
}

func (s *AcceptResourceShareInvitationResponseBodyResourceShareInvitation) SetResourceShareId(v string) *AcceptResourceShareInvitationResponseBodyResourceShareInvitation {
	s.ResourceShareId = &v
	return s
}

func (s *AcceptResourceShareInvitationResponseBodyResourceShareInvitation) SetResourceShareInvitationId(v string) *AcceptResourceShareInvitationResponseBodyResourceShareInvitation {
	s.ResourceShareInvitationId = &v
	return s
}

func (s *AcceptResourceShareInvitationResponseBodyResourceShareInvitation) SetResourceShareName(v string) *AcceptResourceShareInvitationResponseBodyResourceShareInvitation {
	s.ResourceShareName = &v
	return s
}

func (s *AcceptResourceShareInvitationResponseBodyResourceShareInvitation) SetSenderAccountId(v string) *AcceptResourceShareInvitationResponseBodyResourceShareInvitation {
	s.SenderAccountId = &v
	return s
}

func (s *AcceptResourceShareInvitationResponseBodyResourceShareInvitation) SetStatus(v string) *AcceptResourceShareInvitationResponseBodyResourceShareInvitation {
	s.Status = &v
	return s
}

type AcceptResourceShareInvitationResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AcceptResourceShareInvitationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AcceptResourceShareInvitationResponse) String() string {
	return tea.Prettify(s)
}

func (s AcceptResourceShareInvitationResponse) GoString() string {
	return s.String()
}

func (s *AcceptResourceShareInvitationResponse) SetHeaders(v map[string]*string) *AcceptResourceShareInvitationResponse {
	s.Headers = v
	return s
}

func (s *AcceptResourceShareInvitationResponse) SetStatusCode(v int32) *AcceptResourceShareInvitationResponse {
	s.StatusCode = &v
	return s
}

func (s *AcceptResourceShareInvitationResponse) SetBody(v *AcceptResourceShareInvitationResponseBody) *AcceptResourceShareInvitationResponse {
	s.Body = v
	return s
}

type AssociateResourceShareRequest struct {
	ResourceShareId *string                                   `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
	Resources       []*AssociateResourceShareRequestResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	Targets         []*string                                 `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
}

func (s AssociateResourceShareRequest) String() string {
	return tea.Prettify(s)
}

func (s AssociateResourceShareRequest) GoString() string {
	return s.String()
}

func (s *AssociateResourceShareRequest) SetResourceShareId(v string) *AssociateResourceShareRequest {
	s.ResourceShareId = &v
	return s
}

func (s *AssociateResourceShareRequest) SetResources(v []*AssociateResourceShareRequestResources) *AssociateResourceShareRequest {
	s.Resources = v
	return s
}

func (s *AssociateResourceShareRequest) SetTargets(v []*string) *AssociateResourceShareRequest {
	s.Targets = v
	return s
}

type AssociateResourceShareRequestResources struct {
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s AssociateResourceShareRequestResources) String() string {
	return tea.Prettify(s)
}

func (s AssociateResourceShareRequestResources) GoString() string {
	return s.String()
}

func (s *AssociateResourceShareRequestResources) SetResourceId(v string) *AssociateResourceShareRequestResources {
	s.ResourceId = &v
	return s
}

func (s *AssociateResourceShareRequestResources) SetResourceType(v string) *AssociateResourceShareRequestResources {
	s.ResourceType = &v
	return s
}

type AssociateResourceShareResponseBody struct {
	RequestId                 *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceShareAssociations []*AssociateResourceShareResponseBodyResourceShareAssociations `json:"ResourceShareAssociations,omitempty" xml:"ResourceShareAssociations,omitempty" type:"Repeated"`
}

func (s AssociateResourceShareResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssociateResourceShareResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateResourceShareResponseBody) SetRequestId(v string) *AssociateResourceShareResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateResourceShareResponseBody) SetResourceShareAssociations(v []*AssociateResourceShareResponseBodyResourceShareAssociations) *AssociateResourceShareResponseBody {
	s.ResourceShareAssociations = v
	return s
}

type AssociateResourceShareResponseBodyResourceShareAssociations struct {
	AssociationStatus        *string `json:"AssociationStatus,omitempty" xml:"AssociationStatus,omitempty"`
	AssociationStatusMessage *string `json:"AssociationStatusMessage,omitempty" xml:"AssociationStatusMessage,omitempty"`
	AssociationType          *string `json:"AssociationType,omitempty" xml:"AssociationType,omitempty"`
	CreateTime               *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EntityId                 *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	EntityType               *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	ResourceShareId          *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
	ResourceShareName        *string `json:"ResourceShareName,omitempty" xml:"ResourceShareName,omitempty"`
	UpdateTime               *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s AssociateResourceShareResponseBodyResourceShareAssociations) String() string {
	return tea.Prettify(s)
}

func (s AssociateResourceShareResponseBodyResourceShareAssociations) GoString() string {
	return s.String()
}

func (s *AssociateResourceShareResponseBodyResourceShareAssociations) SetAssociationStatus(v string) *AssociateResourceShareResponseBodyResourceShareAssociations {
	s.AssociationStatus = &v
	return s
}

func (s *AssociateResourceShareResponseBodyResourceShareAssociations) SetAssociationStatusMessage(v string) *AssociateResourceShareResponseBodyResourceShareAssociations {
	s.AssociationStatusMessage = &v
	return s
}

func (s *AssociateResourceShareResponseBodyResourceShareAssociations) SetAssociationType(v string) *AssociateResourceShareResponseBodyResourceShareAssociations {
	s.AssociationType = &v
	return s
}

func (s *AssociateResourceShareResponseBodyResourceShareAssociations) SetCreateTime(v string) *AssociateResourceShareResponseBodyResourceShareAssociations {
	s.CreateTime = &v
	return s
}

func (s *AssociateResourceShareResponseBodyResourceShareAssociations) SetEntityId(v string) *AssociateResourceShareResponseBodyResourceShareAssociations {
	s.EntityId = &v
	return s
}

func (s *AssociateResourceShareResponseBodyResourceShareAssociations) SetEntityType(v string) *AssociateResourceShareResponseBodyResourceShareAssociations {
	s.EntityType = &v
	return s
}

func (s *AssociateResourceShareResponseBodyResourceShareAssociations) SetResourceShareId(v string) *AssociateResourceShareResponseBodyResourceShareAssociations {
	s.ResourceShareId = &v
	return s
}

func (s *AssociateResourceShareResponseBodyResourceShareAssociations) SetResourceShareName(v string) *AssociateResourceShareResponseBodyResourceShareAssociations {
	s.ResourceShareName = &v
	return s
}

func (s *AssociateResourceShareResponseBodyResourceShareAssociations) SetUpdateTime(v string) *AssociateResourceShareResponseBodyResourceShareAssociations {
	s.UpdateTime = &v
	return s
}

type AssociateResourceShareResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AssociateResourceShareResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AssociateResourceShareResponse) String() string {
	return tea.Prettify(s)
}

func (s AssociateResourceShareResponse) GoString() string {
	return s.String()
}

func (s *AssociateResourceShareResponse) SetHeaders(v map[string]*string) *AssociateResourceShareResponse {
	s.Headers = v
	return s
}

func (s *AssociateResourceShareResponse) SetStatusCode(v int32) *AssociateResourceShareResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateResourceShareResponse) SetBody(v *AssociateResourceShareResponseBody) *AssociateResourceShareResponse {
	s.Body = v
	return s
}

type CreateResourceShareRequest struct {
	AllowExternalTargets *bool                                  `json:"AllowExternalTargets,omitempty" xml:"AllowExternalTargets,omitempty"`
	ResourceShareName    *string                                `json:"ResourceShareName,omitempty" xml:"ResourceShareName,omitempty"`
	Resources            []*CreateResourceShareRequestResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	Targets              []*string                              `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
}

func (s CreateResourceShareRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceShareRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceShareRequest) SetAllowExternalTargets(v bool) *CreateResourceShareRequest {
	s.AllowExternalTargets = &v
	return s
}

func (s *CreateResourceShareRequest) SetResourceShareName(v string) *CreateResourceShareRequest {
	s.ResourceShareName = &v
	return s
}

func (s *CreateResourceShareRequest) SetResources(v []*CreateResourceShareRequestResources) *CreateResourceShareRequest {
	s.Resources = v
	return s
}

func (s *CreateResourceShareRequest) SetTargets(v []*string) *CreateResourceShareRequest {
	s.Targets = v
	return s
}

type CreateResourceShareRequestResources struct {
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s CreateResourceShareRequestResources) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceShareRequestResources) GoString() string {
	return s.String()
}

func (s *CreateResourceShareRequestResources) SetResourceId(v string) *CreateResourceShareRequestResources {
	s.ResourceId = &v
	return s
}

func (s *CreateResourceShareRequestResources) SetResourceType(v string) *CreateResourceShareRequestResources {
	s.ResourceType = &v
	return s
}

type CreateResourceShareResponseBody struct {
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceShare *CreateResourceShareResponseBodyResourceShare `json:"ResourceShare,omitempty" xml:"ResourceShare,omitempty" type:"Struct"`
}

func (s CreateResourceShareResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceShareResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceShareResponseBody) SetRequestId(v string) *CreateResourceShareResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateResourceShareResponseBody) SetResourceShare(v *CreateResourceShareResponseBodyResourceShare) *CreateResourceShareResponseBody {
	s.ResourceShare = v
	return s
}

type CreateResourceShareResponseBodyResourceShare struct {
	AllowExternalTargets *bool   `json:"AllowExternalTargets,omitempty" xml:"AllowExternalTargets,omitempty"`
	CreateTime           *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ResourceShareId      *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
	ResourceShareName    *string `json:"ResourceShareName,omitempty" xml:"ResourceShareName,omitempty"`
	ResourceShareOwner   *string `json:"ResourceShareOwner,omitempty" xml:"ResourceShareOwner,omitempty"`
	ResourceShareStatus  *string `json:"ResourceShareStatus,omitempty" xml:"ResourceShareStatus,omitempty"`
	UpdateTime           *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s CreateResourceShareResponseBodyResourceShare) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceShareResponseBodyResourceShare) GoString() string {
	return s.String()
}

func (s *CreateResourceShareResponseBodyResourceShare) SetAllowExternalTargets(v bool) *CreateResourceShareResponseBodyResourceShare {
	s.AllowExternalTargets = &v
	return s
}

func (s *CreateResourceShareResponseBodyResourceShare) SetCreateTime(v string) *CreateResourceShareResponseBodyResourceShare {
	s.CreateTime = &v
	return s
}

func (s *CreateResourceShareResponseBodyResourceShare) SetResourceShareId(v string) *CreateResourceShareResponseBodyResourceShare {
	s.ResourceShareId = &v
	return s
}

func (s *CreateResourceShareResponseBodyResourceShare) SetResourceShareName(v string) *CreateResourceShareResponseBodyResourceShare {
	s.ResourceShareName = &v
	return s
}

func (s *CreateResourceShareResponseBodyResourceShare) SetResourceShareOwner(v string) *CreateResourceShareResponseBodyResourceShare {
	s.ResourceShareOwner = &v
	return s
}

func (s *CreateResourceShareResponseBodyResourceShare) SetResourceShareStatus(v string) *CreateResourceShareResponseBodyResourceShare {
	s.ResourceShareStatus = &v
	return s
}

func (s *CreateResourceShareResponseBodyResourceShare) SetUpdateTime(v string) *CreateResourceShareResponseBodyResourceShare {
	s.UpdateTime = &v
	return s
}

type CreateResourceShareResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateResourceShareResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateResourceShareResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceShareResponse) GoString() string {
	return s.String()
}

func (s *CreateResourceShareResponse) SetHeaders(v map[string]*string) *CreateResourceShareResponse {
	s.Headers = v
	return s
}

func (s *CreateResourceShareResponse) SetStatusCode(v int32) *CreateResourceShareResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateResourceShareResponse) SetBody(v *CreateResourceShareResponseBody) *CreateResourceShareResponse {
	s.Body = v
	return s
}

type DeleteResourceShareRequest struct {
	ResourceShareId *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
}

func (s DeleteResourceShareRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceShareRequest) GoString() string {
	return s.String()
}

func (s *DeleteResourceShareRequest) SetResourceShareId(v string) *DeleteResourceShareRequest {
	s.ResourceShareId = &v
	return s
}

type DeleteResourceShareResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteResourceShareResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceShareResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceShareResponseBody) SetRequestId(v string) *DeleteResourceShareResponseBody {
	s.RequestId = &v
	return s
}

type DeleteResourceShareResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteResourceShareResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteResourceShareResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceShareResponse) GoString() string {
	return s.String()
}

func (s *DeleteResourceShareResponse) SetHeaders(v map[string]*string) *DeleteResourceShareResponse {
	s.Headers = v
	return s
}

func (s *DeleteResourceShareResponse) SetStatusCode(v int32) *DeleteResourceShareResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResourceShareResponse) SetBody(v *DeleteResourceShareResponseBody) *DeleteResourceShareResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetAcceptLanguage(v string) *DescribeRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

type DescribeRegionsResponseBody struct {
	Regions   []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	LocalName      *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetLocalName(v string) *DescribeRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionId(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DisassociateResourceShareRequest struct {
	ResourceOwner   *string                                      `json:"ResourceOwner,omitempty" xml:"ResourceOwner,omitempty"`
	ResourceShareId *string                                      `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
	Resources       []*DisassociateResourceShareRequestResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	Targets         []*string                                    `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
}

func (s DisassociateResourceShareRequest) String() string {
	return tea.Prettify(s)
}

func (s DisassociateResourceShareRequest) GoString() string {
	return s.String()
}

func (s *DisassociateResourceShareRequest) SetResourceOwner(v string) *DisassociateResourceShareRequest {
	s.ResourceOwner = &v
	return s
}

func (s *DisassociateResourceShareRequest) SetResourceShareId(v string) *DisassociateResourceShareRequest {
	s.ResourceShareId = &v
	return s
}

func (s *DisassociateResourceShareRequest) SetResources(v []*DisassociateResourceShareRequestResources) *DisassociateResourceShareRequest {
	s.Resources = v
	return s
}

func (s *DisassociateResourceShareRequest) SetTargets(v []*string) *DisassociateResourceShareRequest {
	s.Targets = v
	return s
}

type DisassociateResourceShareRequestResources struct {
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DisassociateResourceShareRequestResources) String() string {
	return tea.Prettify(s)
}

func (s DisassociateResourceShareRequestResources) GoString() string {
	return s.String()
}

func (s *DisassociateResourceShareRequestResources) SetResourceId(v string) *DisassociateResourceShareRequestResources {
	s.ResourceId = &v
	return s
}

func (s *DisassociateResourceShareRequestResources) SetResourceType(v string) *DisassociateResourceShareRequestResources {
	s.ResourceType = &v
	return s
}

type DisassociateResourceShareResponseBody struct {
	RequestId                 *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceShareAssociations []*DisassociateResourceShareResponseBodyResourceShareAssociations `json:"ResourceShareAssociations,omitempty" xml:"ResourceShareAssociations,omitempty" type:"Repeated"`
}

func (s DisassociateResourceShareResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisassociateResourceShareResponseBody) GoString() string {
	return s.String()
}

func (s *DisassociateResourceShareResponseBody) SetRequestId(v string) *DisassociateResourceShareResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisassociateResourceShareResponseBody) SetResourceShareAssociations(v []*DisassociateResourceShareResponseBodyResourceShareAssociations) *DisassociateResourceShareResponseBody {
	s.ResourceShareAssociations = v
	return s
}

type DisassociateResourceShareResponseBodyResourceShareAssociations struct {
	AssociationStatus        *string `json:"AssociationStatus,omitempty" xml:"AssociationStatus,omitempty"`
	AssociationStatusMessage *string `json:"AssociationStatusMessage,omitempty" xml:"AssociationStatusMessage,omitempty"`
	AssociationType          *string `json:"AssociationType,omitempty" xml:"AssociationType,omitempty"`
	CreateTime               *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EntityId                 *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	EntityType               *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	ResourceShareId          *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
	ResourceShareName        *string `json:"ResourceShareName,omitempty" xml:"ResourceShareName,omitempty"`
	UpdateTime               *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DisassociateResourceShareResponseBodyResourceShareAssociations) String() string {
	return tea.Prettify(s)
}

func (s DisassociateResourceShareResponseBodyResourceShareAssociations) GoString() string {
	return s.String()
}

func (s *DisassociateResourceShareResponseBodyResourceShareAssociations) SetAssociationStatus(v string) *DisassociateResourceShareResponseBodyResourceShareAssociations {
	s.AssociationStatus = &v
	return s
}

func (s *DisassociateResourceShareResponseBodyResourceShareAssociations) SetAssociationStatusMessage(v string) *DisassociateResourceShareResponseBodyResourceShareAssociations {
	s.AssociationStatusMessage = &v
	return s
}

func (s *DisassociateResourceShareResponseBodyResourceShareAssociations) SetAssociationType(v string) *DisassociateResourceShareResponseBodyResourceShareAssociations {
	s.AssociationType = &v
	return s
}

func (s *DisassociateResourceShareResponseBodyResourceShareAssociations) SetCreateTime(v string) *DisassociateResourceShareResponseBodyResourceShareAssociations {
	s.CreateTime = &v
	return s
}

func (s *DisassociateResourceShareResponseBodyResourceShareAssociations) SetEntityId(v string) *DisassociateResourceShareResponseBodyResourceShareAssociations {
	s.EntityId = &v
	return s
}

func (s *DisassociateResourceShareResponseBodyResourceShareAssociations) SetEntityType(v string) *DisassociateResourceShareResponseBodyResourceShareAssociations {
	s.EntityType = &v
	return s
}

func (s *DisassociateResourceShareResponseBodyResourceShareAssociations) SetResourceShareId(v string) *DisassociateResourceShareResponseBodyResourceShareAssociations {
	s.ResourceShareId = &v
	return s
}

func (s *DisassociateResourceShareResponseBodyResourceShareAssociations) SetResourceShareName(v string) *DisassociateResourceShareResponseBodyResourceShareAssociations {
	s.ResourceShareName = &v
	return s
}

func (s *DisassociateResourceShareResponseBodyResourceShareAssociations) SetUpdateTime(v string) *DisassociateResourceShareResponseBodyResourceShareAssociations {
	s.UpdateTime = &v
	return s
}

type DisassociateResourceShareResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisassociateResourceShareResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisassociateResourceShareResponse) String() string {
	return tea.Prettify(s)
}

func (s DisassociateResourceShareResponse) GoString() string {
	return s.String()
}

func (s *DisassociateResourceShareResponse) SetHeaders(v map[string]*string) *DisassociateResourceShareResponse {
	s.Headers = v
	return s
}

func (s *DisassociateResourceShareResponse) SetStatusCode(v int32) *DisassociateResourceShareResponse {
	s.StatusCode = &v
	return s
}

func (s *DisassociateResourceShareResponse) SetBody(v *DisassociateResourceShareResponseBody) *DisassociateResourceShareResponse {
	s.Body = v
	return s
}

type EnableSharingWithResourceDirectoryResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableSharingWithResourceDirectoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableSharingWithResourceDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *EnableSharingWithResourceDirectoryResponseBody) SetRequestId(v string) *EnableSharingWithResourceDirectoryResponseBody {
	s.RequestId = &v
	return s
}

type EnableSharingWithResourceDirectoryResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableSharingWithResourceDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableSharingWithResourceDirectoryResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableSharingWithResourceDirectoryResponse) GoString() string {
	return s.String()
}

func (s *EnableSharingWithResourceDirectoryResponse) SetHeaders(v map[string]*string) *EnableSharingWithResourceDirectoryResponse {
	s.Headers = v
	return s
}

func (s *EnableSharingWithResourceDirectoryResponse) SetStatusCode(v int32) *EnableSharingWithResourceDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableSharingWithResourceDirectoryResponse) SetBody(v *EnableSharingWithResourceDirectoryResponseBody) *EnableSharingWithResourceDirectoryResponse {
	s.Body = v
	return s
}

type ListResourceShareAssociationsRequest struct {
	AssociationStatus *string   `json:"AssociationStatus,omitempty" xml:"AssociationStatus,omitempty"`
	AssociationType   *string   `json:"AssociationType,omitempty" xml:"AssociationType,omitempty"`
	MaxResults        *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken         *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceId        *string   `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceShareIds  []*string `json:"ResourceShareIds,omitempty" xml:"ResourceShareIds,omitempty" type:"Repeated"`
	Target            *string   `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s ListResourceShareAssociationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourceShareAssociationsRequest) GoString() string {
	return s.String()
}

func (s *ListResourceShareAssociationsRequest) SetAssociationStatus(v string) *ListResourceShareAssociationsRequest {
	s.AssociationStatus = &v
	return s
}

func (s *ListResourceShareAssociationsRequest) SetAssociationType(v string) *ListResourceShareAssociationsRequest {
	s.AssociationType = &v
	return s
}

func (s *ListResourceShareAssociationsRequest) SetMaxResults(v int32) *ListResourceShareAssociationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListResourceShareAssociationsRequest) SetNextToken(v string) *ListResourceShareAssociationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListResourceShareAssociationsRequest) SetResourceId(v string) *ListResourceShareAssociationsRequest {
	s.ResourceId = &v
	return s
}

func (s *ListResourceShareAssociationsRequest) SetResourceShareIds(v []*string) *ListResourceShareAssociationsRequest {
	s.ResourceShareIds = v
	return s
}

func (s *ListResourceShareAssociationsRequest) SetTarget(v string) *ListResourceShareAssociationsRequest {
	s.Target = &v
	return s
}

type ListResourceShareAssociationsResponseBody struct {
	NextToken                 *string                                                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId                 *string                                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceShareAssociations []*ListResourceShareAssociationsResponseBodyResourceShareAssociations `json:"ResourceShareAssociations,omitempty" xml:"ResourceShareAssociations,omitempty" type:"Repeated"`
}

func (s ListResourceShareAssociationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourceShareAssociationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceShareAssociationsResponseBody) SetNextToken(v string) *ListResourceShareAssociationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListResourceShareAssociationsResponseBody) SetRequestId(v string) *ListResourceShareAssociationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceShareAssociationsResponseBody) SetResourceShareAssociations(v []*ListResourceShareAssociationsResponseBodyResourceShareAssociations) *ListResourceShareAssociationsResponseBody {
	s.ResourceShareAssociations = v
	return s
}

type ListResourceShareAssociationsResponseBodyResourceShareAssociations struct {
	AssociationStatus        *string `json:"AssociationStatus,omitempty" xml:"AssociationStatus,omitempty"`
	AssociationStatusMessage *string `json:"AssociationStatusMessage,omitempty" xml:"AssociationStatusMessage,omitempty"`
	AssociationType          *string `json:"AssociationType,omitempty" xml:"AssociationType,omitempty"`
	CreateTime               *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EntityId                 *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	EntityType               *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	External                 *bool   `json:"External,omitempty" xml:"External,omitempty"`
	ResourceShareId          *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
	ResourceShareName        *string `json:"ResourceShareName,omitempty" xml:"ResourceShareName,omitempty"`
	UpdateTime               *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListResourceShareAssociationsResponseBodyResourceShareAssociations) String() string {
	return tea.Prettify(s)
}

func (s ListResourceShareAssociationsResponseBodyResourceShareAssociations) GoString() string {
	return s.String()
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociations) SetAssociationStatus(v string) *ListResourceShareAssociationsResponseBodyResourceShareAssociations {
	s.AssociationStatus = &v
	return s
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociations) SetAssociationStatusMessage(v string) *ListResourceShareAssociationsResponseBodyResourceShareAssociations {
	s.AssociationStatusMessage = &v
	return s
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociations) SetAssociationType(v string) *ListResourceShareAssociationsResponseBodyResourceShareAssociations {
	s.AssociationType = &v
	return s
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociations) SetCreateTime(v string) *ListResourceShareAssociationsResponseBodyResourceShareAssociations {
	s.CreateTime = &v
	return s
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociations) SetEntityId(v string) *ListResourceShareAssociationsResponseBodyResourceShareAssociations {
	s.EntityId = &v
	return s
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociations) SetEntityType(v string) *ListResourceShareAssociationsResponseBodyResourceShareAssociations {
	s.EntityType = &v
	return s
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociations) SetExternal(v bool) *ListResourceShareAssociationsResponseBodyResourceShareAssociations {
	s.External = &v
	return s
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociations) SetResourceShareId(v string) *ListResourceShareAssociationsResponseBodyResourceShareAssociations {
	s.ResourceShareId = &v
	return s
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociations) SetResourceShareName(v string) *ListResourceShareAssociationsResponseBodyResourceShareAssociations {
	s.ResourceShareName = &v
	return s
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociations) SetUpdateTime(v string) *ListResourceShareAssociationsResponseBodyResourceShareAssociations {
	s.UpdateTime = &v
	return s
}

type ListResourceShareAssociationsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListResourceShareAssociationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListResourceShareAssociationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourceShareAssociationsResponse) GoString() string {
	return s.String()
}

func (s *ListResourceShareAssociationsResponse) SetHeaders(v map[string]*string) *ListResourceShareAssociationsResponse {
	s.Headers = v
	return s
}

func (s *ListResourceShareAssociationsResponse) SetStatusCode(v int32) *ListResourceShareAssociationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceShareAssociationsResponse) SetBody(v *ListResourceShareAssociationsResponseBody) *ListResourceShareAssociationsResponse {
	s.Body = v
	return s
}

type ListResourceShareInvitationsRequest struct {
	MaxResults                 *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken                  *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceShareIds           []*string `json:"ResourceShareIds,omitempty" xml:"ResourceShareIds,omitempty" type:"Repeated"`
	ResourceShareInvitationIds []*string `json:"ResourceShareInvitationIds,omitempty" xml:"ResourceShareInvitationIds,omitempty" type:"Repeated"`
}

func (s ListResourceShareInvitationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourceShareInvitationsRequest) GoString() string {
	return s.String()
}

func (s *ListResourceShareInvitationsRequest) SetMaxResults(v int32) *ListResourceShareInvitationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListResourceShareInvitationsRequest) SetNextToken(v string) *ListResourceShareInvitationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListResourceShareInvitationsRequest) SetResourceShareIds(v []*string) *ListResourceShareInvitationsRequest {
	s.ResourceShareIds = v
	return s
}

func (s *ListResourceShareInvitationsRequest) SetResourceShareInvitationIds(v []*string) *ListResourceShareInvitationsRequest {
	s.ResourceShareInvitationIds = v
	return s
}

type ListResourceShareInvitationsResponseBody struct {
	NextToken                *string                                                             `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId                *string                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceShareInvitations []*ListResourceShareInvitationsResponseBodyResourceShareInvitations `json:"ResourceShareInvitations,omitempty" xml:"ResourceShareInvitations,omitempty" type:"Repeated"`
}

func (s ListResourceShareInvitationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourceShareInvitationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceShareInvitationsResponseBody) SetNextToken(v string) *ListResourceShareInvitationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListResourceShareInvitationsResponseBody) SetRequestId(v string) *ListResourceShareInvitationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceShareInvitationsResponseBody) SetResourceShareInvitations(v []*ListResourceShareInvitationsResponseBodyResourceShareInvitations) *ListResourceShareInvitationsResponseBody {
	s.ResourceShareInvitations = v
	return s
}

type ListResourceShareInvitationsResponseBodyResourceShareInvitations struct {
	CreateTime                *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ReceiverAccountId         *string `json:"ReceiverAccountId,omitempty" xml:"ReceiverAccountId,omitempty"`
	ResourceShareId           *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
	ResourceShareInvitationId *string `json:"ResourceShareInvitationId,omitempty" xml:"ResourceShareInvitationId,omitempty"`
	ResourceShareName         *string `json:"ResourceShareName,omitempty" xml:"ResourceShareName,omitempty"`
	SenderAccountId           *string `json:"SenderAccountId,omitempty" xml:"SenderAccountId,omitempty"`
	Status                    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListResourceShareInvitationsResponseBodyResourceShareInvitations) String() string {
	return tea.Prettify(s)
}

func (s ListResourceShareInvitationsResponseBodyResourceShareInvitations) GoString() string {
	return s.String()
}

func (s *ListResourceShareInvitationsResponseBodyResourceShareInvitations) SetCreateTime(v string) *ListResourceShareInvitationsResponseBodyResourceShareInvitations {
	s.CreateTime = &v
	return s
}

func (s *ListResourceShareInvitationsResponseBodyResourceShareInvitations) SetReceiverAccountId(v string) *ListResourceShareInvitationsResponseBodyResourceShareInvitations {
	s.ReceiverAccountId = &v
	return s
}

func (s *ListResourceShareInvitationsResponseBodyResourceShareInvitations) SetResourceShareId(v string) *ListResourceShareInvitationsResponseBodyResourceShareInvitations {
	s.ResourceShareId = &v
	return s
}

func (s *ListResourceShareInvitationsResponseBodyResourceShareInvitations) SetResourceShareInvitationId(v string) *ListResourceShareInvitationsResponseBodyResourceShareInvitations {
	s.ResourceShareInvitationId = &v
	return s
}

func (s *ListResourceShareInvitationsResponseBodyResourceShareInvitations) SetResourceShareName(v string) *ListResourceShareInvitationsResponseBodyResourceShareInvitations {
	s.ResourceShareName = &v
	return s
}

func (s *ListResourceShareInvitationsResponseBodyResourceShareInvitations) SetSenderAccountId(v string) *ListResourceShareInvitationsResponseBodyResourceShareInvitations {
	s.SenderAccountId = &v
	return s
}

func (s *ListResourceShareInvitationsResponseBodyResourceShareInvitations) SetStatus(v string) *ListResourceShareInvitationsResponseBodyResourceShareInvitations {
	s.Status = &v
	return s
}

type ListResourceShareInvitationsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListResourceShareInvitationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListResourceShareInvitationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourceShareInvitationsResponse) GoString() string {
	return s.String()
}

func (s *ListResourceShareInvitationsResponse) SetHeaders(v map[string]*string) *ListResourceShareInvitationsResponse {
	s.Headers = v
	return s
}

func (s *ListResourceShareInvitationsResponse) SetStatusCode(v int32) *ListResourceShareInvitationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceShareInvitationsResponse) SetBody(v *ListResourceShareInvitationsResponseBody) *ListResourceShareInvitationsResponse {
	s.Body = v
	return s
}

type ListResourceSharesRequest struct {
	MaxResults          *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken           *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceOwner       *string   `json:"ResourceOwner,omitempty" xml:"ResourceOwner,omitempty"`
	ResourceShareIds    []*string `json:"ResourceShareIds,omitempty" xml:"ResourceShareIds,omitempty" type:"Repeated"`
	ResourceShareName   *string   `json:"ResourceShareName,omitempty" xml:"ResourceShareName,omitempty"`
	ResourceShareStatus *string   `json:"ResourceShareStatus,omitempty" xml:"ResourceShareStatus,omitempty"`
}

func (s ListResourceSharesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourceSharesRequest) GoString() string {
	return s.String()
}

func (s *ListResourceSharesRequest) SetMaxResults(v int32) *ListResourceSharesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListResourceSharesRequest) SetNextToken(v string) *ListResourceSharesRequest {
	s.NextToken = &v
	return s
}

func (s *ListResourceSharesRequest) SetResourceOwner(v string) *ListResourceSharesRequest {
	s.ResourceOwner = &v
	return s
}

func (s *ListResourceSharesRequest) SetResourceShareIds(v []*string) *ListResourceSharesRequest {
	s.ResourceShareIds = v
	return s
}

func (s *ListResourceSharesRequest) SetResourceShareName(v string) *ListResourceSharesRequest {
	s.ResourceShareName = &v
	return s
}

func (s *ListResourceSharesRequest) SetResourceShareStatus(v string) *ListResourceSharesRequest {
	s.ResourceShareStatus = &v
	return s
}

type ListResourceSharesResponseBody struct {
	NextToken      *string                                         `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceShares []*ListResourceSharesResponseBodyResourceShares `json:"ResourceShares,omitempty" xml:"ResourceShares,omitempty" type:"Repeated"`
}

func (s ListResourceSharesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourceSharesResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceSharesResponseBody) SetNextToken(v string) *ListResourceSharesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListResourceSharesResponseBody) SetRequestId(v string) *ListResourceSharesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceSharesResponseBody) SetResourceShares(v []*ListResourceSharesResponseBodyResourceShares) *ListResourceSharesResponseBody {
	s.ResourceShares = v
	return s
}

type ListResourceSharesResponseBodyResourceShares struct {
	AllowExternalTargets *bool   `json:"AllowExternalTargets,omitempty" xml:"AllowExternalTargets,omitempty"`
	CreateTime           *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ResourceShareId      *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
	ResourceShareName    *string `json:"ResourceShareName,omitempty" xml:"ResourceShareName,omitempty"`
	ResourceShareOwner   *string `json:"ResourceShareOwner,omitempty" xml:"ResourceShareOwner,omitempty"`
	ResourceShareStatus  *string `json:"ResourceShareStatus,omitempty" xml:"ResourceShareStatus,omitempty"`
	UpdateTime           *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListResourceSharesResponseBodyResourceShares) String() string {
	return tea.Prettify(s)
}

func (s ListResourceSharesResponseBodyResourceShares) GoString() string {
	return s.String()
}

func (s *ListResourceSharesResponseBodyResourceShares) SetAllowExternalTargets(v bool) *ListResourceSharesResponseBodyResourceShares {
	s.AllowExternalTargets = &v
	return s
}

func (s *ListResourceSharesResponseBodyResourceShares) SetCreateTime(v string) *ListResourceSharesResponseBodyResourceShares {
	s.CreateTime = &v
	return s
}

func (s *ListResourceSharesResponseBodyResourceShares) SetResourceShareId(v string) *ListResourceSharesResponseBodyResourceShares {
	s.ResourceShareId = &v
	return s
}

func (s *ListResourceSharesResponseBodyResourceShares) SetResourceShareName(v string) *ListResourceSharesResponseBodyResourceShares {
	s.ResourceShareName = &v
	return s
}

func (s *ListResourceSharesResponseBodyResourceShares) SetResourceShareOwner(v string) *ListResourceSharesResponseBodyResourceShares {
	s.ResourceShareOwner = &v
	return s
}

func (s *ListResourceSharesResponseBodyResourceShares) SetResourceShareStatus(v string) *ListResourceSharesResponseBodyResourceShares {
	s.ResourceShareStatus = &v
	return s
}

func (s *ListResourceSharesResponseBodyResourceShares) SetUpdateTime(v string) *ListResourceSharesResponseBodyResourceShares {
	s.UpdateTime = &v
	return s
}

type ListResourceSharesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListResourceSharesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListResourceSharesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourceSharesResponse) GoString() string {
	return s.String()
}

func (s *ListResourceSharesResponse) SetHeaders(v map[string]*string) *ListResourceSharesResponse {
	s.Headers = v
	return s
}

func (s *ListResourceSharesResponse) SetStatusCode(v int32) *ListResourceSharesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceSharesResponse) SetBody(v *ListResourceSharesResponseBody) *ListResourceSharesResponse {
	s.Body = v
	return s
}

type ListSharedResourcesRequest struct {
	MaxResults       *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken        *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceIds      []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	ResourceOwner    *string   `json:"ResourceOwner,omitempty" xml:"ResourceOwner,omitempty"`
	ResourceShareIds []*string `json:"ResourceShareIds,omitempty" xml:"ResourceShareIds,omitempty" type:"Repeated"`
	ResourceType     *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Target           *string   `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s ListSharedResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSharedResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListSharedResourcesRequest) SetMaxResults(v int32) *ListSharedResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSharedResourcesRequest) SetNextToken(v string) *ListSharedResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListSharedResourcesRequest) SetResourceIds(v []*string) *ListSharedResourcesRequest {
	s.ResourceIds = v
	return s
}

func (s *ListSharedResourcesRequest) SetResourceOwner(v string) *ListSharedResourcesRequest {
	s.ResourceOwner = &v
	return s
}

func (s *ListSharedResourcesRequest) SetResourceShareIds(v []*string) *ListSharedResourcesRequest {
	s.ResourceShareIds = v
	return s
}

func (s *ListSharedResourcesRequest) SetResourceType(v string) *ListSharedResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListSharedResourcesRequest) SetTarget(v string) *ListSharedResourcesRequest {
	s.Target = &v
	return s
}

type ListSharedResourcesResponseBody struct {
	NextToken       *string                                           `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId       *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SharedResources []*ListSharedResourcesResponseBodySharedResources `json:"SharedResources,omitempty" xml:"SharedResources,omitempty" type:"Repeated"`
}

func (s ListSharedResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSharedResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSharedResourcesResponseBody) SetNextToken(v string) *ListSharedResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSharedResourcesResponseBody) SetRequestId(v string) *ListSharedResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSharedResourcesResponseBody) SetSharedResources(v []*ListSharedResourcesResponseBodySharedResources) *ListSharedResourcesResponseBody {
	s.SharedResources = v
	return s
}

type ListSharedResourcesResponseBodySharedResources struct {
	CreateTime            *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ResourceId            *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceShareId       *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
	ResourceStatus        *string `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty"`
	ResourceStatusMessage *string `json:"ResourceStatusMessage,omitempty" xml:"ResourceStatusMessage,omitempty"`
	ResourceType          *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	UpdateTime            *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListSharedResourcesResponseBodySharedResources) String() string {
	return tea.Prettify(s)
}

func (s ListSharedResourcesResponseBodySharedResources) GoString() string {
	return s.String()
}

func (s *ListSharedResourcesResponseBodySharedResources) SetCreateTime(v string) *ListSharedResourcesResponseBodySharedResources {
	s.CreateTime = &v
	return s
}

func (s *ListSharedResourcesResponseBodySharedResources) SetResourceId(v string) *ListSharedResourcesResponseBodySharedResources {
	s.ResourceId = &v
	return s
}

func (s *ListSharedResourcesResponseBodySharedResources) SetResourceShareId(v string) *ListSharedResourcesResponseBodySharedResources {
	s.ResourceShareId = &v
	return s
}

func (s *ListSharedResourcesResponseBodySharedResources) SetResourceStatus(v string) *ListSharedResourcesResponseBodySharedResources {
	s.ResourceStatus = &v
	return s
}

func (s *ListSharedResourcesResponseBodySharedResources) SetResourceStatusMessage(v string) *ListSharedResourcesResponseBodySharedResources {
	s.ResourceStatusMessage = &v
	return s
}

func (s *ListSharedResourcesResponseBodySharedResources) SetResourceType(v string) *ListSharedResourcesResponseBodySharedResources {
	s.ResourceType = &v
	return s
}

func (s *ListSharedResourcesResponseBodySharedResources) SetUpdateTime(v string) *ListSharedResourcesResponseBodySharedResources {
	s.UpdateTime = &v
	return s
}

type ListSharedResourcesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSharedResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSharedResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSharedResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListSharedResourcesResponse) SetHeaders(v map[string]*string) *ListSharedResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListSharedResourcesResponse) SetStatusCode(v int32) *ListSharedResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSharedResourcesResponse) SetBody(v *ListSharedResourcesResponseBody) *ListSharedResourcesResponse {
	s.Body = v
	return s
}

type ListSharedTargetsRequest struct {
	MaxResults       *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken        *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceId       *string   `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceOwner    *string   `json:"ResourceOwner,omitempty" xml:"ResourceOwner,omitempty"`
	ResourceShareIds []*string `json:"ResourceShareIds,omitempty" xml:"ResourceShareIds,omitempty" type:"Repeated"`
	ResourceType     *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Targets          []*string `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
}

func (s ListSharedTargetsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSharedTargetsRequest) GoString() string {
	return s.String()
}

func (s *ListSharedTargetsRequest) SetMaxResults(v int32) *ListSharedTargetsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSharedTargetsRequest) SetNextToken(v string) *ListSharedTargetsRequest {
	s.NextToken = &v
	return s
}

func (s *ListSharedTargetsRequest) SetResourceId(v string) *ListSharedTargetsRequest {
	s.ResourceId = &v
	return s
}

func (s *ListSharedTargetsRequest) SetResourceOwner(v string) *ListSharedTargetsRequest {
	s.ResourceOwner = &v
	return s
}

func (s *ListSharedTargetsRequest) SetResourceShareIds(v []*string) *ListSharedTargetsRequest {
	s.ResourceShareIds = v
	return s
}

func (s *ListSharedTargetsRequest) SetResourceType(v string) *ListSharedTargetsRequest {
	s.ResourceType = &v
	return s
}

func (s *ListSharedTargetsRequest) SetTargets(v []*string) *ListSharedTargetsRequest {
	s.Targets = v
	return s
}

type ListSharedTargetsResponseBody struct {
	NextToken     *string                                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SharedTargets []*ListSharedTargetsResponseBodySharedTargets `json:"SharedTargets,omitempty" xml:"SharedTargets,omitempty" type:"Repeated"`
}

func (s ListSharedTargetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSharedTargetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSharedTargetsResponseBody) SetNextToken(v string) *ListSharedTargetsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSharedTargetsResponseBody) SetRequestId(v string) *ListSharedTargetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSharedTargetsResponseBody) SetSharedTargets(v []*ListSharedTargetsResponseBodySharedTargets) *ListSharedTargetsResponseBody {
	s.SharedTargets = v
	return s
}

type ListSharedTargetsResponseBodySharedTargets struct {
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	External        *bool   `json:"External,omitempty" xml:"External,omitempty"`
	ResourceShareId *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
	TargetId        *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	UpdateTime      *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListSharedTargetsResponseBodySharedTargets) String() string {
	return tea.Prettify(s)
}

func (s ListSharedTargetsResponseBodySharedTargets) GoString() string {
	return s.String()
}

func (s *ListSharedTargetsResponseBodySharedTargets) SetCreateTime(v string) *ListSharedTargetsResponseBodySharedTargets {
	s.CreateTime = &v
	return s
}

func (s *ListSharedTargetsResponseBodySharedTargets) SetExternal(v bool) *ListSharedTargetsResponseBodySharedTargets {
	s.External = &v
	return s
}

func (s *ListSharedTargetsResponseBodySharedTargets) SetResourceShareId(v string) *ListSharedTargetsResponseBodySharedTargets {
	s.ResourceShareId = &v
	return s
}

func (s *ListSharedTargetsResponseBodySharedTargets) SetTargetId(v string) *ListSharedTargetsResponseBodySharedTargets {
	s.TargetId = &v
	return s
}

func (s *ListSharedTargetsResponseBodySharedTargets) SetUpdateTime(v string) *ListSharedTargetsResponseBodySharedTargets {
	s.UpdateTime = &v
	return s
}

type ListSharedTargetsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSharedTargetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSharedTargetsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSharedTargetsResponse) GoString() string {
	return s.String()
}

func (s *ListSharedTargetsResponse) SetHeaders(v map[string]*string) *ListSharedTargetsResponse {
	s.Headers = v
	return s
}

func (s *ListSharedTargetsResponse) SetStatusCode(v int32) *ListSharedTargetsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSharedTargetsResponse) SetBody(v *ListSharedTargetsResponseBody) *ListSharedTargetsResponse {
	s.Body = v
	return s
}

type RejectResourceShareInvitationRequest struct {
	ResourceShareInvitationId *string `json:"ResourceShareInvitationId,omitempty" xml:"ResourceShareInvitationId,omitempty"`
}

func (s RejectResourceShareInvitationRequest) String() string {
	return tea.Prettify(s)
}

func (s RejectResourceShareInvitationRequest) GoString() string {
	return s.String()
}

func (s *RejectResourceShareInvitationRequest) SetResourceShareInvitationId(v string) *RejectResourceShareInvitationRequest {
	s.ResourceShareInvitationId = &v
	return s
}

type RejectResourceShareInvitationResponseBody struct {
	RequestId               *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceShareInvitation *RejectResourceShareInvitationResponseBodyResourceShareInvitation `json:"ResourceShareInvitation,omitempty" xml:"ResourceShareInvitation,omitempty" type:"Struct"`
}

func (s RejectResourceShareInvitationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RejectResourceShareInvitationResponseBody) GoString() string {
	return s.String()
}

func (s *RejectResourceShareInvitationResponseBody) SetRequestId(v string) *RejectResourceShareInvitationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RejectResourceShareInvitationResponseBody) SetResourceShareInvitation(v *RejectResourceShareInvitationResponseBodyResourceShareInvitation) *RejectResourceShareInvitationResponseBody {
	s.ResourceShareInvitation = v
	return s
}

type RejectResourceShareInvitationResponseBodyResourceShareInvitation struct {
	CreateTime                *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ReceiverAccountId         *string `json:"ReceiverAccountId,omitempty" xml:"ReceiverAccountId,omitempty"`
	ResourceShareId           *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
	ResourceShareInvitationId *string `json:"ResourceShareInvitationId,omitempty" xml:"ResourceShareInvitationId,omitempty"`
	ResourceShareName         *string `json:"ResourceShareName,omitempty" xml:"ResourceShareName,omitempty"`
	SenderAccountId           *string `json:"SenderAccountId,omitempty" xml:"SenderAccountId,omitempty"`
	Status                    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s RejectResourceShareInvitationResponseBodyResourceShareInvitation) String() string {
	return tea.Prettify(s)
}

func (s RejectResourceShareInvitationResponseBodyResourceShareInvitation) GoString() string {
	return s.String()
}

func (s *RejectResourceShareInvitationResponseBodyResourceShareInvitation) SetCreateTime(v string) *RejectResourceShareInvitationResponseBodyResourceShareInvitation {
	s.CreateTime = &v
	return s
}

func (s *RejectResourceShareInvitationResponseBodyResourceShareInvitation) SetReceiverAccountId(v string) *RejectResourceShareInvitationResponseBodyResourceShareInvitation {
	s.ReceiverAccountId = &v
	return s
}

func (s *RejectResourceShareInvitationResponseBodyResourceShareInvitation) SetResourceShareId(v string) *RejectResourceShareInvitationResponseBodyResourceShareInvitation {
	s.ResourceShareId = &v
	return s
}

func (s *RejectResourceShareInvitationResponseBodyResourceShareInvitation) SetResourceShareInvitationId(v string) *RejectResourceShareInvitationResponseBodyResourceShareInvitation {
	s.ResourceShareInvitationId = &v
	return s
}

func (s *RejectResourceShareInvitationResponseBodyResourceShareInvitation) SetResourceShareName(v string) *RejectResourceShareInvitationResponseBodyResourceShareInvitation {
	s.ResourceShareName = &v
	return s
}

func (s *RejectResourceShareInvitationResponseBodyResourceShareInvitation) SetSenderAccountId(v string) *RejectResourceShareInvitationResponseBodyResourceShareInvitation {
	s.SenderAccountId = &v
	return s
}

func (s *RejectResourceShareInvitationResponseBodyResourceShareInvitation) SetStatus(v string) *RejectResourceShareInvitationResponseBodyResourceShareInvitation {
	s.Status = &v
	return s
}

type RejectResourceShareInvitationResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RejectResourceShareInvitationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RejectResourceShareInvitationResponse) String() string {
	return tea.Prettify(s)
}

func (s RejectResourceShareInvitationResponse) GoString() string {
	return s.String()
}

func (s *RejectResourceShareInvitationResponse) SetHeaders(v map[string]*string) *RejectResourceShareInvitationResponse {
	s.Headers = v
	return s
}

func (s *RejectResourceShareInvitationResponse) SetStatusCode(v int32) *RejectResourceShareInvitationResponse {
	s.StatusCode = &v
	return s
}

func (s *RejectResourceShareInvitationResponse) SetBody(v *RejectResourceShareInvitationResponseBody) *RejectResourceShareInvitationResponse {
	s.Body = v
	return s
}

type UpdateResourceShareRequest struct {
	AllowExternalTargets *bool   `json:"AllowExternalTargets,omitempty" xml:"AllowExternalTargets,omitempty"`
	ResourceShareId      *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
	ResourceShareName    *string `json:"ResourceShareName,omitempty" xml:"ResourceShareName,omitempty"`
}

func (s UpdateResourceShareRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceShareRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceShareRequest) SetAllowExternalTargets(v bool) *UpdateResourceShareRequest {
	s.AllowExternalTargets = &v
	return s
}

func (s *UpdateResourceShareRequest) SetResourceShareId(v string) *UpdateResourceShareRequest {
	s.ResourceShareId = &v
	return s
}

func (s *UpdateResourceShareRequest) SetResourceShareName(v string) *UpdateResourceShareRequest {
	s.ResourceShareName = &v
	return s
}

type UpdateResourceShareResponseBody struct {
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceShare *UpdateResourceShareResponseBodyResourceShare `json:"ResourceShare,omitempty" xml:"ResourceShare,omitempty" type:"Struct"`
}

func (s UpdateResourceShareResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceShareResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourceShareResponseBody) SetRequestId(v string) *UpdateResourceShareResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateResourceShareResponseBody) SetResourceShare(v *UpdateResourceShareResponseBodyResourceShare) *UpdateResourceShareResponseBody {
	s.ResourceShare = v
	return s
}

type UpdateResourceShareResponseBodyResourceShare struct {
	AllowExternalTargets *bool   `json:"AllowExternalTargets,omitempty" xml:"AllowExternalTargets,omitempty"`
	CreateTime           *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ResourceShareId      *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
	ResourceShareName    *string `json:"ResourceShareName,omitempty" xml:"ResourceShareName,omitempty"`
	ResourceShareOwner   *string `json:"ResourceShareOwner,omitempty" xml:"ResourceShareOwner,omitempty"`
	ResourceShareStatus  *string `json:"ResourceShareStatus,omitempty" xml:"ResourceShareStatus,omitempty"`
	UpdateTime           *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s UpdateResourceShareResponseBodyResourceShare) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceShareResponseBodyResourceShare) GoString() string {
	return s.String()
}

func (s *UpdateResourceShareResponseBodyResourceShare) SetAllowExternalTargets(v bool) *UpdateResourceShareResponseBodyResourceShare {
	s.AllowExternalTargets = &v
	return s
}

func (s *UpdateResourceShareResponseBodyResourceShare) SetCreateTime(v string) *UpdateResourceShareResponseBodyResourceShare {
	s.CreateTime = &v
	return s
}

func (s *UpdateResourceShareResponseBodyResourceShare) SetResourceShareId(v string) *UpdateResourceShareResponseBodyResourceShare {
	s.ResourceShareId = &v
	return s
}

func (s *UpdateResourceShareResponseBodyResourceShare) SetResourceShareName(v string) *UpdateResourceShareResponseBodyResourceShare {
	s.ResourceShareName = &v
	return s
}

func (s *UpdateResourceShareResponseBodyResourceShare) SetResourceShareOwner(v string) *UpdateResourceShareResponseBodyResourceShare {
	s.ResourceShareOwner = &v
	return s
}

func (s *UpdateResourceShareResponseBodyResourceShare) SetResourceShareStatus(v string) *UpdateResourceShareResponseBodyResourceShare {
	s.ResourceShareStatus = &v
	return s
}

func (s *UpdateResourceShareResponseBodyResourceShare) SetUpdateTime(v string) *UpdateResourceShareResponseBodyResourceShare {
	s.UpdateTime = &v
	return s
}

type UpdateResourceShareResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateResourceShareResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateResourceShareResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceShareResponse) GoString() string {
	return s.String()
}

func (s *UpdateResourceShareResponse) SetHeaders(v map[string]*string) *UpdateResourceShareResponse {
	s.Headers = v
	return s
}

func (s *UpdateResourceShareResponse) SetStatusCode(v int32) *UpdateResourceShareResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResourceShareResponse) SetBody(v *UpdateResourceShareResponseBody) *UpdateResourceShareResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("resourcesharing"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AcceptResourceShareInvitationWithOptions(request *AcceptResourceShareInvitationRequest, runtime *util.RuntimeOptions) (_result *AcceptResourceShareInvitationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceShareInvitationId)) {
		query["ResourceShareInvitationId"] = request.ResourceShareInvitationId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AcceptResourceShareInvitation"),
		Version:     tea.String("2020-01-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AcceptResourceShareInvitationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AcceptResourceShareInvitation(request *AcceptResourceShareInvitationRequest) (_result *AcceptResourceShareInvitationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AcceptResourceShareInvitationResponse{}
	_body, _err := client.AcceptResourceShareInvitationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AssociateResourceShareWithOptions(request *AssociateResourceShareRequest, runtime *util.RuntimeOptions) (_result *AssociateResourceShareResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceShareId)) {
		query["ResourceShareId"] = request.ResourceShareId
	}

	if !tea.BoolValue(util.IsUnset(request.Resources)) {
		query["Resources"] = request.Resources
	}

	if !tea.BoolValue(util.IsUnset(request.Targets)) {
		query["Targets"] = request.Targets
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AssociateResourceShare"),
		Version:     tea.String("2020-01-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AssociateResourceShareResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AssociateResourceShare(request *AssociateResourceShareRequest) (_result *AssociateResourceShareResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssociateResourceShareResponse{}
	_body, _err := client.AssociateResourceShareWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateResourceShareWithOptions(request *CreateResourceShareRequest, runtime *util.RuntimeOptions) (_result *CreateResourceShareResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllowExternalTargets)) {
		query["AllowExternalTargets"] = request.AllowExternalTargets
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceShareName)) {
		query["ResourceShareName"] = request.ResourceShareName
	}

	if !tea.BoolValue(util.IsUnset(request.Resources)) {
		query["Resources"] = request.Resources
	}

	if !tea.BoolValue(util.IsUnset(request.Targets)) {
		query["Targets"] = request.Targets
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateResourceShare"),
		Version:     tea.String("2020-01-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateResourceShareResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateResourceShare(request *CreateResourceShareRequest) (_result *CreateResourceShareResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateResourceShareResponse{}
	_body, _err := client.CreateResourceShareWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteResourceShareWithOptions(request *DeleteResourceShareRequest, runtime *util.RuntimeOptions) (_result *DeleteResourceShareResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceShareId)) {
		query["ResourceShareId"] = request.ResourceShareId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteResourceShare"),
		Version:     tea.String("2020-01-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteResourceShareResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteResourceShare(request *DeleteResourceShareRequest) (_result *DeleteResourceShareResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteResourceShareResponse{}
	_body, _err := client.DeleteResourceShareWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2020-01-10"),
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

func (client *Client) DisassociateResourceShareWithOptions(request *DisassociateResourceShareRequest, runtime *util.RuntimeOptions) (_result *DisassociateResourceShareResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceOwner)) {
		query["ResourceOwner"] = request.ResourceOwner
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceShareId)) {
		query["ResourceShareId"] = request.ResourceShareId
	}

	if !tea.BoolValue(util.IsUnset(request.Resources)) {
		query["Resources"] = request.Resources
	}

	if !tea.BoolValue(util.IsUnset(request.Targets)) {
		query["Targets"] = request.Targets
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisassociateResourceShare"),
		Version:     tea.String("2020-01-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisassociateResourceShareResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisassociateResourceShare(request *DisassociateResourceShareRequest) (_result *DisassociateResourceShareResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisassociateResourceShareResponse{}
	_body, _err := client.DisassociateResourceShareWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableSharingWithResourceDirectoryWithOptions(runtime *util.RuntimeOptions) (_result *EnableSharingWithResourceDirectoryResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("EnableSharingWithResourceDirectory"),
		Version:     tea.String("2020-01-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableSharingWithResourceDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableSharingWithResourceDirectory() (_result *EnableSharingWithResourceDirectoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableSharingWithResourceDirectoryResponse{}
	_body, _err := client.EnableSharingWithResourceDirectoryWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListResourceShareAssociationsWithOptions(request *ListResourceShareAssociationsRequest, runtime *util.RuntimeOptions) (_result *ListResourceShareAssociationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssociationStatus)) {
		query["AssociationStatus"] = request.AssociationStatus
	}

	if !tea.BoolValue(util.IsUnset(request.AssociationType)) {
		query["AssociationType"] = request.AssociationType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceShareIds)) {
		query["ResourceShareIds"] = request.ResourceShareIds
	}

	if !tea.BoolValue(util.IsUnset(request.Target)) {
		query["Target"] = request.Target
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResourceShareAssociations"),
		Version:     tea.String("2020-01-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResourceShareAssociationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListResourceShareAssociations(request *ListResourceShareAssociationsRequest) (_result *ListResourceShareAssociationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListResourceShareAssociationsResponse{}
	_body, _err := client.ListResourceShareAssociationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListResourceShareInvitationsWithOptions(request *ListResourceShareInvitationsRequest, runtime *util.RuntimeOptions) (_result *ListResourceShareInvitationsResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceShareIds)) {
		query["ResourceShareIds"] = request.ResourceShareIds
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceShareInvitationIds)) {
		query["ResourceShareInvitationIds"] = request.ResourceShareInvitationIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResourceShareInvitations"),
		Version:     tea.String("2020-01-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResourceShareInvitationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListResourceShareInvitations(request *ListResourceShareInvitationsRequest) (_result *ListResourceShareInvitationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListResourceShareInvitationsResponse{}
	_body, _err := client.ListResourceShareInvitationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListResourceSharesWithOptions(request *ListResourceSharesRequest, runtime *util.RuntimeOptions) (_result *ListResourceSharesResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceOwner)) {
		query["ResourceOwner"] = request.ResourceOwner
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceShareIds)) {
		query["ResourceShareIds"] = request.ResourceShareIds
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceShareName)) {
		query["ResourceShareName"] = request.ResourceShareName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceShareStatus)) {
		query["ResourceShareStatus"] = request.ResourceShareStatus
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResourceShares"),
		Version:     tea.String("2020-01-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResourceSharesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListResourceShares(request *ListResourceSharesRequest) (_result *ListResourceSharesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListResourceSharesResponse{}
	_body, _err := client.ListResourceSharesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSharedResourcesWithOptions(request *ListSharedResourcesRequest, runtime *util.RuntimeOptions) (_result *ListSharedResourcesResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceIds)) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwner)) {
		query["ResourceOwner"] = request.ResourceOwner
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceShareIds)) {
		query["ResourceShareIds"] = request.ResourceShareIds
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Target)) {
		query["Target"] = request.Target
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSharedResources"),
		Version:     tea.String("2020-01-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSharedResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSharedResources(request *ListSharedResourcesRequest) (_result *ListSharedResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSharedResourcesResponse{}
	_body, _err := client.ListSharedResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSharedTargetsWithOptions(request *ListSharedTargetsRequest, runtime *util.RuntimeOptions) (_result *ListSharedTargetsResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwner)) {
		query["ResourceOwner"] = request.ResourceOwner
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceShareIds)) {
		query["ResourceShareIds"] = request.ResourceShareIds
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Targets)) {
		query["Targets"] = request.Targets
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSharedTargets"),
		Version:     tea.String("2020-01-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSharedTargetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSharedTargets(request *ListSharedTargetsRequest) (_result *ListSharedTargetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSharedTargetsResponse{}
	_body, _err := client.ListSharedTargetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RejectResourceShareInvitationWithOptions(request *RejectResourceShareInvitationRequest, runtime *util.RuntimeOptions) (_result *RejectResourceShareInvitationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceShareInvitationId)) {
		query["ResourceShareInvitationId"] = request.ResourceShareInvitationId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RejectResourceShareInvitation"),
		Version:     tea.String("2020-01-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RejectResourceShareInvitationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RejectResourceShareInvitation(request *RejectResourceShareInvitationRequest) (_result *RejectResourceShareInvitationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RejectResourceShareInvitationResponse{}
	_body, _err := client.RejectResourceShareInvitationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateResourceShareWithOptions(request *UpdateResourceShareRequest, runtime *util.RuntimeOptions) (_result *UpdateResourceShareResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllowExternalTargets)) {
		query["AllowExternalTargets"] = request.AllowExternalTargets
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceShareId)) {
		query["ResourceShareId"] = request.ResourceShareId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceShareName)) {
		query["ResourceShareName"] = request.ResourceShareName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateResourceShare"),
		Version:     tea.String("2020-01-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateResourceShareResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateResourceShare(request *UpdateResourceShareRequest) (_result *UpdateResourceShareResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateResourceShareResponse{}
	_body, _err := client.UpdateResourceShareWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
