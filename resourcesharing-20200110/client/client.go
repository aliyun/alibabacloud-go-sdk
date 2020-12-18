// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

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

type DescribeRegionsResponse struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Regions   []*DescribeRegionsResponseRegions `json:"Regions,omitempty" xml:"Regions,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetRequestId(v string) *DescribeRegionsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponse) SetRegions(v []*DescribeRegionsResponseRegions) *DescribeRegionsResponse {
	s.Regions = v
	return s
}

type DescribeRegionsResponseRegions struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	LocalName      *string `json:"LocalName,omitempty" xml:"LocalName,omitempty" require:"true"`
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty" require:"true"`
}

func (s DescribeRegionsResponseRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseRegions) SetRegionId(v string) *DescribeRegionsResponseRegions {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseRegions) SetLocalName(v string) *DescribeRegionsResponseRegions {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseRegions) SetRegionEndpoint(v string) *DescribeRegionsResponseRegions {
	s.RegionEndpoint = &v
	return s
}

type ListResourceSharesRequest struct {
	ResourceShareIds    []*string `json:"ResourceShareIds,omitempty" xml:"ResourceShareIds,omitempty" type:"Repeated"`
	ResourceOwner       *string   `json:"ResourceOwner,omitempty" xml:"ResourceOwner,omitempty" require:"true"`
	ResourceShareName   *string   `json:"ResourceShareName,omitempty" xml:"ResourceShareName,omitempty"`
	ResourceShareStatus *string   `json:"ResourceShareStatus,omitempty" xml:"ResourceShareStatus,omitempty"`
	MaxResults          *int      `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken           *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListResourceSharesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourceSharesRequest) GoString() string {
	return s.String()
}

func (s *ListResourceSharesRequest) SetResourceShareIds(v []*string) *ListResourceSharesRequest {
	s.ResourceShareIds = v
	return s
}

func (s *ListResourceSharesRequest) SetResourceOwner(v string) *ListResourceSharesRequest {
	s.ResourceOwner = &v
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

func (s *ListResourceSharesRequest) SetMaxResults(v int) *ListResourceSharesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListResourceSharesRequest) SetNextToken(v string) *ListResourceSharesRequest {
	s.NextToken = &v
	return s
}

type ListResourceSharesResponse struct {
	RequestId      *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	NextToken      *string                                     `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	ResourceShares []*ListResourceSharesResponseResourceShares `json:"ResourceShares,omitempty" xml:"ResourceShares,omitempty" require:"true" type:"Repeated"`
}

func (s ListResourceSharesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourceSharesResponse) GoString() string {
	return s.String()
}

func (s *ListResourceSharesResponse) SetRequestId(v string) *ListResourceSharesResponse {
	s.RequestId = &v
	return s
}

func (s *ListResourceSharesResponse) SetNextToken(v string) *ListResourceSharesResponse {
	s.NextToken = &v
	return s
}

func (s *ListResourceSharesResponse) SetResourceShares(v []*ListResourceSharesResponseResourceShares) *ListResourceSharesResponse {
	s.ResourceShares = v
	return s
}

type ListResourceSharesResponseResourceShares struct {
	ResourceShareId     *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty" require:"true"`
	ResourceShareName   *string `json:"ResourceShareName,omitempty" xml:"ResourceShareName,omitempty" require:"true"`
	ResourceShareOwner  *string `json:"ResourceShareOwner,omitempty" xml:"ResourceShareOwner,omitempty" require:"true"`
	ResourceShareStatus *string `json:"ResourceShareStatus,omitempty" xml:"ResourceShareStatus,omitempty" require:"true"`
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	UpdateTime          *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
}

func (s ListResourceSharesResponseResourceShares) String() string {
	return tea.Prettify(s)
}

func (s ListResourceSharesResponseResourceShares) GoString() string {
	return s.String()
}

func (s *ListResourceSharesResponseResourceShares) SetResourceShareId(v string) *ListResourceSharesResponseResourceShares {
	s.ResourceShareId = &v
	return s
}

func (s *ListResourceSharesResponseResourceShares) SetResourceShareName(v string) *ListResourceSharesResponseResourceShares {
	s.ResourceShareName = &v
	return s
}

func (s *ListResourceSharesResponseResourceShares) SetResourceShareOwner(v string) *ListResourceSharesResponseResourceShares {
	s.ResourceShareOwner = &v
	return s
}

func (s *ListResourceSharesResponseResourceShares) SetResourceShareStatus(v string) *ListResourceSharesResponseResourceShares {
	s.ResourceShareStatus = &v
	return s
}

func (s *ListResourceSharesResponseResourceShares) SetCreateTime(v string) *ListResourceSharesResponseResourceShares {
	s.CreateTime = &v
	return s
}

func (s *ListResourceSharesResponseResourceShares) SetUpdateTime(v string) *ListResourceSharesResponseResourceShares {
	s.UpdateTime = &v
	return s
}

type ListSharedResourcesRequest struct {
	ResourceOwner    *string   `json:"ResourceOwner,omitempty" xml:"ResourceOwner,omitempty" require:"true"`
	ResourceShareIds []*string `json:"ResourceShareIds,omitempty" xml:"ResourceShareIds,omitempty" type:"Repeated"`
	ResourceType     *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ResourceIds      []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	Target           *string   `json:"Target,omitempty" xml:"Target,omitempty"`
	MaxResults       *int      `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken        *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListSharedResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSharedResourcesRequest) GoString() string {
	return s.String()
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

func (s *ListSharedResourcesRequest) SetResourceIds(v []*string) *ListSharedResourcesRequest {
	s.ResourceIds = v
	return s
}

func (s *ListSharedResourcesRequest) SetTarget(v string) *ListSharedResourcesRequest {
	s.Target = &v
	return s
}

func (s *ListSharedResourcesRequest) SetMaxResults(v int) *ListSharedResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSharedResourcesRequest) SetNextToken(v string) *ListSharedResourcesRequest {
	s.NextToken = &v
	return s
}

type ListSharedResourcesResponse struct {
	RequestId       *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	NextToken       *string                                       `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	SharedResources []*ListSharedResourcesResponseSharedResources `json:"SharedResources,omitempty" xml:"SharedResources,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSharedResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListSharedResourcesResponse) SetRequestId(v string) *ListSharedResourcesResponse {
	s.RequestId = &v
	return s
}

func (s *ListSharedResourcesResponse) SetNextToken(v string) *ListSharedResourcesResponse {
	s.NextToken = &v
	return s
}

func (s *ListSharedResourcesResponse) SetSharedResources(v []*ListSharedResourcesResponseSharedResources) *ListSharedResourcesResponse {
	s.SharedResources = v
	return s
}

type ListSharedResourcesResponseSharedResources struct {
	ResourceShareId       *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty" require:"true"`
	ResourceId            *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true"`
	ResourceType          *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" require:"true"`
	ResourceStatus        *string `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty" require:"true"`
	ResourceStatusMessage *string `json:"ResourceStatusMessage,omitempty" xml:"ResourceStatusMessage,omitempty" require:"true"`
	CreateTime            *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	UpdateTime            *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
}

func (s ListSharedResourcesResponseSharedResources) String() string {
	return tea.Prettify(s)
}

func (s ListSharedResourcesResponseSharedResources) GoString() string {
	return s.String()
}

func (s *ListSharedResourcesResponseSharedResources) SetResourceShareId(v string) *ListSharedResourcesResponseSharedResources {
	s.ResourceShareId = &v
	return s
}

func (s *ListSharedResourcesResponseSharedResources) SetResourceId(v string) *ListSharedResourcesResponseSharedResources {
	s.ResourceId = &v
	return s
}

func (s *ListSharedResourcesResponseSharedResources) SetResourceType(v string) *ListSharedResourcesResponseSharedResources {
	s.ResourceType = &v
	return s
}

func (s *ListSharedResourcesResponseSharedResources) SetResourceStatus(v string) *ListSharedResourcesResponseSharedResources {
	s.ResourceStatus = &v
	return s
}

func (s *ListSharedResourcesResponseSharedResources) SetResourceStatusMessage(v string) *ListSharedResourcesResponseSharedResources {
	s.ResourceStatusMessage = &v
	return s
}

func (s *ListSharedResourcesResponseSharedResources) SetCreateTime(v string) *ListSharedResourcesResponseSharedResources {
	s.CreateTime = &v
	return s
}

func (s *ListSharedResourcesResponseSharedResources) SetUpdateTime(v string) *ListSharedResourcesResponseSharedResources {
	s.UpdateTime = &v
	return s
}

type ListSharedTargetsRequest struct {
	ResourceOwner    *string   `json:"ResourceOwner,omitempty" xml:"ResourceOwner,omitempty" require:"true"`
	ResourceShareIds []*string `json:"ResourceShareIds,omitempty" xml:"ResourceShareIds,omitempty" type:"Repeated"`
	Targets          []*string `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
	ResourceType     *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ResourceId       *string   `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	MaxResults       *int      `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken        *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListSharedTargetsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSharedTargetsRequest) GoString() string {
	return s.String()
}

func (s *ListSharedTargetsRequest) SetResourceOwner(v string) *ListSharedTargetsRequest {
	s.ResourceOwner = &v
	return s
}

func (s *ListSharedTargetsRequest) SetResourceShareIds(v []*string) *ListSharedTargetsRequest {
	s.ResourceShareIds = v
	return s
}

func (s *ListSharedTargetsRequest) SetTargets(v []*string) *ListSharedTargetsRequest {
	s.Targets = v
	return s
}

func (s *ListSharedTargetsRequest) SetResourceType(v string) *ListSharedTargetsRequest {
	s.ResourceType = &v
	return s
}

func (s *ListSharedTargetsRequest) SetResourceId(v string) *ListSharedTargetsRequest {
	s.ResourceId = &v
	return s
}

func (s *ListSharedTargetsRequest) SetMaxResults(v int) *ListSharedTargetsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSharedTargetsRequest) SetNextToken(v string) *ListSharedTargetsRequest {
	s.NextToken = &v
	return s
}

type ListSharedTargetsResponse struct {
	RequestId     *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	NextToken     *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	SharedTargets []*ListSharedTargetsResponseSharedTargets `json:"SharedTargets,omitempty" xml:"SharedTargets,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedTargetsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSharedTargetsResponse) GoString() string {
	return s.String()
}

func (s *ListSharedTargetsResponse) SetRequestId(v string) *ListSharedTargetsResponse {
	s.RequestId = &v
	return s
}

func (s *ListSharedTargetsResponse) SetNextToken(v string) *ListSharedTargetsResponse {
	s.NextToken = &v
	return s
}

func (s *ListSharedTargetsResponse) SetSharedTargets(v []*ListSharedTargetsResponseSharedTargets) *ListSharedTargetsResponse {
	s.SharedTargets = v
	return s
}

type ListSharedTargetsResponseSharedTargets struct {
	ResourceShareId *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty" require:"true"`
	TargetId        *string `json:"TargetId,omitempty" xml:"TargetId,omitempty" require:"true"`
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	UpdateTime      *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
}

func (s ListSharedTargetsResponseSharedTargets) String() string {
	return tea.Prettify(s)
}

func (s ListSharedTargetsResponseSharedTargets) GoString() string {
	return s.String()
}

func (s *ListSharedTargetsResponseSharedTargets) SetResourceShareId(v string) *ListSharedTargetsResponseSharedTargets {
	s.ResourceShareId = &v
	return s
}

func (s *ListSharedTargetsResponseSharedTargets) SetTargetId(v string) *ListSharedTargetsResponseSharedTargets {
	s.TargetId = &v
	return s
}

func (s *ListSharedTargetsResponseSharedTargets) SetCreateTime(v string) *ListSharedTargetsResponseSharedTargets {
	s.CreateTime = &v
	return s
}

func (s *ListSharedTargetsResponseSharedTargets) SetUpdateTime(v string) *ListSharedTargetsResponseSharedTargets {
	s.UpdateTime = &v
	return s
}

type AssociateResourceShareRequest struct {
	ResourceShareId *string                                   `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty" require:"true"`
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

type AssociateResourceShareResponse struct {
	RequestId                 *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ResourceShareAssociations []*AssociateResourceShareResponseResourceShareAssociations `json:"ResourceShareAssociations,omitempty" xml:"ResourceShareAssociations,omitempty" require:"true" type:"Repeated"`
}

func (s AssociateResourceShareResponse) String() string {
	return tea.Prettify(s)
}

func (s AssociateResourceShareResponse) GoString() string {
	return s.String()
}

func (s *AssociateResourceShareResponse) SetRequestId(v string) *AssociateResourceShareResponse {
	s.RequestId = &v
	return s
}

func (s *AssociateResourceShareResponse) SetResourceShareAssociations(v []*AssociateResourceShareResponseResourceShareAssociations) *AssociateResourceShareResponse {
	s.ResourceShareAssociations = v
	return s
}

type AssociateResourceShareResponseResourceShareAssociations struct {
	ResourceShareId          *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty" require:"true"`
	ResourceShareName        *string `json:"ResourceShareName,omitempty" xml:"ResourceShareName,omitempty" require:"true"`
	AssociationType          *string `json:"AssociationType,omitempty" xml:"AssociationType,omitempty" require:"true"`
	EntityId                 *string `json:"EntityId,omitempty" xml:"EntityId,omitempty" require:"true"`
	EntityType               *string `json:"EntityType,omitempty" xml:"EntityType,omitempty" require:"true"`
	AssociationStatus        *string `json:"AssociationStatus,omitempty" xml:"AssociationStatus,omitempty" require:"true"`
	AssociationStatusMessage *string `json:"AssociationStatusMessage,omitempty" xml:"AssociationStatusMessage,omitempty" require:"true"`
	CreateTime               *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	UpdateTime               *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
}

func (s AssociateResourceShareResponseResourceShareAssociations) String() string {
	return tea.Prettify(s)
}

func (s AssociateResourceShareResponseResourceShareAssociations) GoString() string {
	return s.String()
}

func (s *AssociateResourceShareResponseResourceShareAssociations) SetResourceShareId(v string) *AssociateResourceShareResponseResourceShareAssociations {
	s.ResourceShareId = &v
	return s
}

func (s *AssociateResourceShareResponseResourceShareAssociations) SetResourceShareName(v string) *AssociateResourceShareResponseResourceShareAssociations {
	s.ResourceShareName = &v
	return s
}

func (s *AssociateResourceShareResponseResourceShareAssociations) SetAssociationType(v string) *AssociateResourceShareResponseResourceShareAssociations {
	s.AssociationType = &v
	return s
}

func (s *AssociateResourceShareResponseResourceShareAssociations) SetEntityId(v string) *AssociateResourceShareResponseResourceShareAssociations {
	s.EntityId = &v
	return s
}

func (s *AssociateResourceShareResponseResourceShareAssociations) SetEntityType(v string) *AssociateResourceShareResponseResourceShareAssociations {
	s.EntityType = &v
	return s
}

func (s *AssociateResourceShareResponseResourceShareAssociations) SetAssociationStatus(v string) *AssociateResourceShareResponseResourceShareAssociations {
	s.AssociationStatus = &v
	return s
}

func (s *AssociateResourceShareResponseResourceShareAssociations) SetAssociationStatusMessage(v string) *AssociateResourceShareResponseResourceShareAssociations {
	s.AssociationStatusMessage = &v
	return s
}

func (s *AssociateResourceShareResponseResourceShareAssociations) SetCreateTime(v string) *AssociateResourceShareResponseResourceShareAssociations {
	s.CreateTime = &v
	return s
}

func (s *AssociateResourceShareResponseResourceShareAssociations) SetUpdateTime(v string) *AssociateResourceShareResponseResourceShareAssociations {
	s.UpdateTime = &v
	return s
}

type UpdateResourceShareRequest struct {
	ResourceShareId   *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty" require:"true"`
	ResourceShareName *string `json:"ResourceShareName,omitempty" xml:"ResourceShareName,omitempty" require:"true"`
}

func (s UpdateResourceShareRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceShareRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceShareRequest) SetResourceShareId(v string) *UpdateResourceShareRequest {
	s.ResourceShareId = &v
	return s
}

func (s *UpdateResourceShareRequest) SetResourceShareName(v string) *UpdateResourceShareRequest {
	s.ResourceShareName = &v
	return s
}

type UpdateResourceShareResponse struct {
	RequestId     *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ResourceShare *UpdateResourceShareResponseResourceShare `json:"ResourceShare,omitempty" xml:"ResourceShare,omitempty" require:"true" type:"Struct"`
}

func (s UpdateResourceShareResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceShareResponse) GoString() string {
	return s.String()
}

func (s *UpdateResourceShareResponse) SetRequestId(v string) *UpdateResourceShareResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateResourceShareResponse) SetResourceShare(v *UpdateResourceShareResponseResourceShare) *UpdateResourceShareResponse {
	s.ResourceShare = v
	return s
}

type UpdateResourceShareResponseResourceShare struct {
	ResourceShareId     *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty" require:"true"`
	ResourceShareName   *string `json:"ResourceShareName,omitempty" xml:"ResourceShareName,omitempty" require:"true"`
	ResourceShareOwner  *string `json:"ResourceShareOwner,omitempty" xml:"ResourceShareOwner,omitempty" require:"true"`
	ResourceShareStatus *string `json:"ResourceShareStatus,omitempty" xml:"ResourceShareStatus,omitempty" require:"true"`
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	UpdateTime          *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
}

func (s UpdateResourceShareResponseResourceShare) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceShareResponseResourceShare) GoString() string {
	return s.String()
}

func (s *UpdateResourceShareResponseResourceShare) SetResourceShareId(v string) *UpdateResourceShareResponseResourceShare {
	s.ResourceShareId = &v
	return s
}

func (s *UpdateResourceShareResponseResourceShare) SetResourceShareName(v string) *UpdateResourceShareResponseResourceShare {
	s.ResourceShareName = &v
	return s
}

func (s *UpdateResourceShareResponseResourceShare) SetResourceShareOwner(v string) *UpdateResourceShareResponseResourceShare {
	s.ResourceShareOwner = &v
	return s
}

func (s *UpdateResourceShareResponseResourceShare) SetResourceShareStatus(v string) *UpdateResourceShareResponseResourceShare {
	s.ResourceShareStatus = &v
	return s
}

func (s *UpdateResourceShareResponseResourceShare) SetCreateTime(v string) *UpdateResourceShareResponseResourceShare {
	s.CreateTime = &v
	return s
}

func (s *UpdateResourceShareResponseResourceShare) SetUpdateTime(v string) *UpdateResourceShareResponseResourceShare {
	s.UpdateTime = &v
	return s
}

type DeleteResourceShareRequest struct {
	ResourceShareId *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty" require:"true"`
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

type DeleteResourceShareResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteResourceShareResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceShareResponse) GoString() string {
	return s.String()
}

func (s *DeleteResourceShareResponse) SetRequestId(v string) *DeleteResourceShareResponse {
	s.RequestId = &v
	return s
}

type DisassociateResourceShareRequest struct {
	ResourceShareId *string                                      `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty" require:"true"`
	Resources       []*DisassociateResourceShareRequestResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	Targets         []*string                                    `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
}

func (s DisassociateResourceShareRequest) String() string {
	return tea.Prettify(s)
}

func (s DisassociateResourceShareRequest) GoString() string {
	return s.String()
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

type DisassociateResourceShareResponse struct {
	RequestId                 *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ResourceShareAssociations []*DisassociateResourceShareResponseResourceShareAssociations `json:"ResourceShareAssociations,omitempty" xml:"ResourceShareAssociations,omitempty" require:"true" type:"Repeated"`
}

func (s DisassociateResourceShareResponse) String() string {
	return tea.Prettify(s)
}

func (s DisassociateResourceShareResponse) GoString() string {
	return s.String()
}

func (s *DisassociateResourceShareResponse) SetRequestId(v string) *DisassociateResourceShareResponse {
	s.RequestId = &v
	return s
}

func (s *DisassociateResourceShareResponse) SetResourceShareAssociations(v []*DisassociateResourceShareResponseResourceShareAssociations) *DisassociateResourceShareResponse {
	s.ResourceShareAssociations = v
	return s
}

type DisassociateResourceShareResponseResourceShareAssociations struct {
	ResourceShareId          *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty" require:"true"`
	ResourceShareName        *string `json:"ResourceShareName,omitempty" xml:"ResourceShareName,omitempty" require:"true"`
	AssociationType          *string `json:"AssociationType,omitempty" xml:"AssociationType,omitempty" require:"true"`
	EntityId                 *string `json:"EntityId,omitempty" xml:"EntityId,omitempty" require:"true"`
	EntityType               *string `json:"EntityType,omitempty" xml:"EntityType,omitempty" require:"true"`
	AssociationStatus        *string `json:"AssociationStatus,omitempty" xml:"AssociationStatus,omitempty" require:"true"`
	AssociationStatusMessage *string `json:"AssociationStatusMessage,omitempty" xml:"AssociationStatusMessage,omitempty" require:"true"`
	CreateTime               *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	UpdateTime               *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
}

func (s DisassociateResourceShareResponseResourceShareAssociations) String() string {
	return tea.Prettify(s)
}

func (s DisassociateResourceShareResponseResourceShareAssociations) GoString() string {
	return s.String()
}

func (s *DisassociateResourceShareResponseResourceShareAssociations) SetResourceShareId(v string) *DisassociateResourceShareResponseResourceShareAssociations {
	s.ResourceShareId = &v
	return s
}

func (s *DisassociateResourceShareResponseResourceShareAssociations) SetResourceShareName(v string) *DisassociateResourceShareResponseResourceShareAssociations {
	s.ResourceShareName = &v
	return s
}

func (s *DisassociateResourceShareResponseResourceShareAssociations) SetAssociationType(v string) *DisassociateResourceShareResponseResourceShareAssociations {
	s.AssociationType = &v
	return s
}

func (s *DisassociateResourceShareResponseResourceShareAssociations) SetEntityId(v string) *DisassociateResourceShareResponseResourceShareAssociations {
	s.EntityId = &v
	return s
}

func (s *DisassociateResourceShareResponseResourceShareAssociations) SetEntityType(v string) *DisassociateResourceShareResponseResourceShareAssociations {
	s.EntityType = &v
	return s
}

func (s *DisassociateResourceShareResponseResourceShareAssociations) SetAssociationStatus(v string) *DisassociateResourceShareResponseResourceShareAssociations {
	s.AssociationStatus = &v
	return s
}

func (s *DisassociateResourceShareResponseResourceShareAssociations) SetAssociationStatusMessage(v string) *DisassociateResourceShareResponseResourceShareAssociations {
	s.AssociationStatusMessage = &v
	return s
}

func (s *DisassociateResourceShareResponseResourceShareAssociations) SetCreateTime(v string) *DisassociateResourceShareResponseResourceShareAssociations {
	s.CreateTime = &v
	return s
}

func (s *DisassociateResourceShareResponseResourceShareAssociations) SetUpdateTime(v string) *DisassociateResourceShareResponseResourceShareAssociations {
	s.UpdateTime = &v
	return s
}

type ListResourceShareAssociationsRequest struct {
	ResourceShareIds  []*string `json:"ResourceShareIds,omitempty" xml:"ResourceShareIds,omitempty" type:"Repeated"`
	ResourceId        *string   `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Target            *string   `json:"Target,omitempty" xml:"Target,omitempty"`
	AssociationType   *string   `json:"AssociationType,omitempty" xml:"AssociationType,omitempty" require:"true"`
	AssociationStatus *string   `json:"AssociationStatus,omitempty" xml:"AssociationStatus,omitempty"`
	MaxResults        *int      `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken         *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListResourceShareAssociationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourceShareAssociationsRequest) GoString() string {
	return s.String()
}

func (s *ListResourceShareAssociationsRequest) SetResourceShareIds(v []*string) *ListResourceShareAssociationsRequest {
	s.ResourceShareIds = v
	return s
}

func (s *ListResourceShareAssociationsRequest) SetResourceId(v string) *ListResourceShareAssociationsRequest {
	s.ResourceId = &v
	return s
}

func (s *ListResourceShareAssociationsRequest) SetTarget(v string) *ListResourceShareAssociationsRequest {
	s.Target = &v
	return s
}

func (s *ListResourceShareAssociationsRequest) SetAssociationType(v string) *ListResourceShareAssociationsRequest {
	s.AssociationType = &v
	return s
}

func (s *ListResourceShareAssociationsRequest) SetAssociationStatus(v string) *ListResourceShareAssociationsRequest {
	s.AssociationStatus = &v
	return s
}

func (s *ListResourceShareAssociationsRequest) SetMaxResults(v int) *ListResourceShareAssociationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListResourceShareAssociationsRequest) SetNextToken(v string) *ListResourceShareAssociationsRequest {
	s.NextToken = &v
	return s
}

type ListResourceShareAssociationsResponse struct {
	RequestId                 *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	NextToken                 *string                                                           `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	ResourceShareAssociations []*ListResourceShareAssociationsResponseResourceShareAssociations `json:"ResourceShareAssociations,omitempty" xml:"ResourceShareAssociations,omitempty" require:"true" type:"Repeated"`
}

func (s ListResourceShareAssociationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourceShareAssociationsResponse) GoString() string {
	return s.String()
}

func (s *ListResourceShareAssociationsResponse) SetRequestId(v string) *ListResourceShareAssociationsResponse {
	s.RequestId = &v
	return s
}

func (s *ListResourceShareAssociationsResponse) SetNextToken(v string) *ListResourceShareAssociationsResponse {
	s.NextToken = &v
	return s
}

func (s *ListResourceShareAssociationsResponse) SetResourceShareAssociations(v []*ListResourceShareAssociationsResponseResourceShareAssociations) *ListResourceShareAssociationsResponse {
	s.ResourceShareAssociations = v
	return s
}

type ListResourceShareAssociationsResponseResourceShareAssociations struct {
	ResourceShareId          *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty" require:"true"`
	ResourceShareName        *string `json:"ResourceShareName,omitempty" xml:"ResourceShareName,omitempty" require:"true"`
	AssociationType          *string `json:"AssociationType,omitempty" xml:"AssociationType,omitempty" require:"true"`
	EntityId                 *string `json:"EntityId,omitempty" xml:"EntityId,omitempty" require:"true"`
	EntityType               *string `json:"EntityType,omitempty" xml:"EntityType,omitempty" require:"true"`
	AssociationStatus        *string `json:"AssociationStatus,omitempty" xml:"AssociationStatus,omitempty" require:"true"`
	AssociationStatusMessage *string `json:"AssociationStatusMessage,omitempty" xml:"AssociationStatusMessage,omitempty" require:"true"`
	CreateTime               *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	UpdateTime               *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
}

func (s ListResourceShareAssociationsResponseResourceShareAssociations) String() string {
	return tea.Prettify(s)
}

func (s ListResourceShareAssociationsResponseResourceShareAssociations) GoString() string {
	return s.String()
}

func (s *ListResourceShareAssociationsResponseResourceShareAssociations) SetResourceShareId(v string) *ListResourceShareAssociationsResponseResourceShareAssociations {
	s.ResourceShareId = &v
	return s
}

func (s *ListResourceShareAssociationsResponseResourceShareAssociations) SetResourceShareName(v string) *ListResourceShareAssociationsResponseResourceShareAssociations {
	s.ResourceShareName = &v
	return s
}

func (s *ListResourceShareAssociationsResponseResourceShareAssociations) SetAssociationType(v string) *ListResourceShareAssociationsResponseResourceShareAssociations {
	s.AssociationType = &v
	return s
}

func (s *ListResourceShareAssociationsResponseResourceShareAssociations) SetEntityId(v string) *ListResourceShareAssociationsResponseResourceShareAssociations {
	s.EntityId = &v
	return s
}

func (s *ListResourceShareAssociationsResponseResourceShareAssociations) SetEntityType(v string) *ListResourceShareAssociationsResponseResourceShareAssociations {
	s.EntityType = &v
	return s
}

func (s *ListResourceShareAssociationsResponseResourceShareAssociations) SetAssociationStatus(v string) *ListResourceShareAssociationsResponseResourceShareAssociations {
	s.AssociationStatus = &v
	return s
}

func (s *ListResourceShareAssociationsResponseResourceShareAssociations) SetAssociationStatusMessage(v string) *ListResourceShareAssociationsResponseResourceShareAssociations {
	s.AssociationStatusMessage = &v
	return s
}

func (s *ListResourceShareAssociationsResponseResourceShareAssociations) SetCreateTime(v string) *ListResourceShareAssociationsResponseResourceShareAssociations {
	s.CreateTime = &v
	return s
}

func (s *ListResourceShareAssociationsResponseResourceShareAssociations) SetUpdateTime(v string) *ListResourceShareAssociationsResponseResourceShareAssociations {
	s.UpdateTime = &v
	return s
}

type CreateResourceShareRequest struct {
	ResourceShareName *string                                `json:"ResourceShareName,omitempty" xml:"ResourceShareName,omitempty" require:"true"`
	Resources         []*CreateResourceShareRequestResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	Targets           []*string                              `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
}

func (s CreateResourceShareRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceShareRequest) GoString() string {
	return s.String()
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

type CreateResourceShareResponse struct {
	RequestId     *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ResourceShare *CreateResourceShareResponseResourceShare `json:"ResourceShare,omitempty" xml:"ResourceShare,omitempty" require:"true" type:"Struct"`
}

func (s CreateResourceShareResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceShareResponse) GoString() string {
	return s.String()
}

func (s *CreateResourceShareResponse) SetRequestId(v string) *CreateResourceShareResponse {
	s.RequestId = &v
	return s
}

func (s *CreateResourceShareResponse) SetResourceShare(v *CreateResourceShareResponseResourceShare) *CreateResourceShareResponse {
	s.ResourceShare = v
	return s
}

type CreateResourceShareResponseResourceShare struct {
	ResourceShareId     *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty" require:"true"`
	ResourceShareName   *string `json:"ResourceShareName,omitempty" xml:"ResourceShareName,omitempty" require:"true"`
	ResourceShareOwner  *string `json:"ResourceShareOwner,omitempty" xml:"ResourceShareOwner,omitempty" require:"true"`
	ResourceShareStatus *string `json:"ResourceShareStatus,omitempty" xml:"ResourceShareStatus,omitempty" require:"true"`
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	UpdateTime          *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
}

func (s CreateResourceShareResponseResourceShare) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceShareResponseResourceShare) GoString() string {
	return s.String()
}

func (s *CreateResourceShareResponseResourceShare) SetResourceShareId(v string) *CreateResourceShareResponseResourceShare {
	s.ResourceShareId = &v
	return s
}

func (s *CreateResourceShareResponseResourceShare) SetResourceShareName(v string) *CreateResourceShareResponseResourceShare {
	s.ResourceShareName = &v
	return s
}

func (s *CreateResourceShareResponseResourceShare) SetResourceShareOwner(v string) *CreateResourceShareResponseResourceShare {
	s.ResourceShareOwner = &v
	return s
}

func (s *CreateResourceShareResponseResourceShare) SetResourceShareStatus(v string) *CreateResourceShareResponseResourceShare {
	s.ResourceShareStatus = &v
	return s
}

func (s *CreateResourceShareResponseResourceShare) SetCreateTime(v string) *CreateResourceShareResponseResourceShare {
	s.CreateTime = &v
	return s
}

func (s *CreateResourceShareResponseResourceShare) SetUpdateTime(v string) *CreateResourceShareResponseResourceShare {
	s.UpdateTime = &v
	return s
}

type Client struct {
	rpc.Client
}

func NewClient(config *rpc.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *rpc.Config) (_err error) {
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

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeRegions"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-01-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) ListResourceSharesWithOptions(request *ListResourceSharesRequest, runtime *util.RuntimeOptions) (_result *ListResourceSharesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListResourceSharesResponse{}
	_body, _err := client.DoRequest(tea.String("ListResourceShares"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-01-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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
	_result = &ListSharedResourcesResponse{}
	_body, _err := client.DoRequest(tea.String("ListSharedResources"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-01-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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
	_result = &ListSharedTargetsResponse{}
	_body, _err := client.DoRequest(tea.String("ListSharedTargets"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-01-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) AssociateResourceShareWithOptions(request *AssociateResourceShareRequest, runtime *util.RuntimeOptions) (_result *AssociateResourceShareResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AssociateResourceShareResponse{}
	_body, _err := client.DoRequest(tea.String("AssociateResourceShare"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-01-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) UpdateResourceShareWithOptions(request *UpdateResourceShareRequest, runtime *util.RuntimeOptions) (_result *UpdateResourceShareResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateResourceShareResponse{}
	_body, _err := client.DoRequest(tea.String("UpdateResourceShare"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-01-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) DeleteResourceShareWithOptions(request *DeleteResourceShareRequest, runtime *util.RuntimeOptions) (_result *DeleteResourceShareResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteResourceShareResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteResourceShare"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-01-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) DisassociateResourceShareWithOptions(request *DisassociateResourceShareRequest, runtime *util.RuntimeOptions) (_result *DisassociateResourceShareResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DisassociateResourceShareResponse{}
	_body, _err := client.DoRequest(tea.String("DisassociateResourceShare"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-01-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) ListResourceShareAssociationsWithOptions(request *ListResourceShareAssociationsRequest, runtime *util.RuntimeOptions) (_result *ListResourceShareAssociationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListResourceShareAssociationsResponse{}
	_body, _err := client.DoRequest(tea.String("ListResourceShareAssociations"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-01-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) CreateResourceShareWithOptions(request *CreateResourceShareRequest, runtime *util.RuntimeOptions) (_result *CreateResourceShareResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateResourceShareResponse{}
	_body, _err := client.DoRequest(tea.String("CreateResourceShare"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-01-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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
