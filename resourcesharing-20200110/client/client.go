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
	// The ID of the resource sharing invitation.
	//
	// You can call the [ListResourceShareInvitations](~~450564~~) operation to obtain the ID.
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
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the resource sharing invitation.
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
	// The information about the failure.
	AcceptInvitationFailedDetails []*AcceptResourceShareInvitationResponseBodyResourceShareInvitationAcceptInvitationFailedDetails `json:"AcceptInvitationFailedDetails,omitempty" xml:"AcceptInvitationFailedDetails,omitempty" type:"Repeated"`
	// The time when the invitation was created. The time is displayed in UTC.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The Alibaba Cloud account ID of the invitee.
	ReceiverAccountId *string `json:"ReceiverAccountId,omitempty" xml:"ReceiverAccountId,omitempty"`
	// The ID of the resource share.
	ResourceShareId *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
	// The ID of the resource sharing invitation.
	ResourceShareInvitationId *string `json:"ResourceShareInvitationId,omitempty" xml:"ResourceShareInvitationId,omitempty"`
	// The name of the resource share.
	ResourceShareName *string `json:"ResourceShareName,omitempty" xml:"ResourceShareName,omitempty"`
	// The Alibaba Cloud account ID of the inviter.
	SenderAccountId *string `json:"SenderAccountId,omitempty" xml:"SenderAccountId,omitempty"`
	// The status of the invitation. Valid values:
	//
	// *   Pending
	// *   Accepted
	// *   Cancelled
	// *   Rejected
	// *   Expired
	// *   AcceptFailed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s AcceptResourceShareInvitationResponseBodyResourceShareInvitation) String() string {
	return tea.Prettify(s)
}

func (s AcceptResourceShareInvitationResponseBodyResourceShareInvitation) GoString() string {
	return s.String()
}

func (s *AcceptResourceShareInvitationResponseBodyResourceShareInvitation) SetAcceptInvitationFailedDetails(v []*AcceptResourceShareInvitationResponseBodyResourceShareInvitationAcceptInvitationFailedDetails) *AcceptResourceShareInvitationResponseBodyResourceShareInvitation {
	s.AcceptInvitationFailedDetails = v
	return s
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

type AcceptResourceShareInvitationResponseBodyResourceShareInvitationAcceptInvitationFailedDetails struct {
	// The type of the sharing operation. Valid values:
	//
	// *   Associate
	AssociateType *string `json:"AssociateType,omitempty" xml:"AssociateType,omitempty"`
	// The ID of the shared resource.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the shared resource.
	//
	// For more information about the types of resources that can be shared, see [Services that work with Resource Sharing](~~450526~~).
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The failure status. Valid values:
	//
	// *   Unavailable: The resource cannot be shared.
	// *   LimitExceeded: The number of shared resources within the Alibaba Cloud account exceeds the upper limit.
	// *   ZonalResourceInaccessible: The resource is unavailable in this region.
	// *   InternalError: An internal error occurred during the check.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The failure cause.
	StatusMessage *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
}

func (s AcceptResourceShareInvitationResponseBodyResourceShareInvitationAcceptInvitationFailedDetails) String() string {
	return tea.Prettify(s)
}

func (s AcceptResourceShareInvitationResponseBodyResourceShareInvitationAcceptInvitationFailedDetails) GoString() string {
	return s.String()
}

func (s *AcceptResourceShareInvitationResponseBodyResourceShareInvitationAcceptInvitationFailedDetails) SetAssociateType(v string) *AcceptResourceShareInvitationResponseBodyResourceShareInvitationAcceptInvitationFailedDetails {
	s.AssociateType = &v
	return s
}

func (s *AcceptResourceShareInvitationResponseBodyResourceShareInvitationAcceptInvitationFailedDetails) SetResourceId(v string) *AcceptResourceShareInvitationResponseBodyResourceShareInvitationAcceptInvitationFailedDetails {
	s.ResourceId = &v
	return s
}

func (s *AcceptResourceShareInvitationResponseBodyResourceShareInvitationAcceptInvitationFailedDetails) SetResourceType(v string) *AcceptResourceShareInvitationResponseBodyResourceShareInvitationAcceptInvitationFailedDetails {
	s.ResourceType = &v
	return s
}

func (s *AcceptResourceShareInvitationResponseBodyResourceShareInvitationAcceptInvitationFailedDetails) SetStatus(v string) *AcceptResourceShareInvitationResponseBodyResourceShareInvitationAcceptInvitationFailedDetails {
	s.Status = &v
	return s
}

func (s *AcceptResourceShareInvitationResponseBodyResourceShareInvitationAcceptInvitationFailedDetails) SetStatusMessage(v string) *AcceptResourceShareInvitationResponseBodyResourceShareInvitationAcceptInvitationFailedDetails {
	s.StatusMessage = &v
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
	// The information about the permissions. If you do not configure this parameter, the system automatically associates the default permission for the specified resource type with the resource share. For more information, see [Permission library](~~465474~~).
	PermissionNames []*string `json:"PermissionNames,omitempty" xml:"PermissionNames,omitempty" type:"Repeated"`
	// The ID of the resource share.
	ResourceShareId *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
	// The information about the resources.
	Resources []*AssociateResourceShareRequestResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// The information about the principals.
	Targets []*string `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
}

func (s AssociateResourceShareRequest) String() string {
	return tea.Prettify(s)
}

func (s AssociateResourceShareRequest) GoString() string {
	return s.String()
}

func (s *AssociateResourceShareRequest) SetPermissionNames(v []*string) *AssociateResourceShareRequest {
	s.PermissionNames = v
	return s
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
	// The ID of a shared resource.
	//
	// Valid values of N: 1 to 5. This indicates that a maximum of five shared resources can be specified at a time.
	//
	// >  Resources.N.ResourceId and Resources.N.ResourceType must be used in pairs.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of a shared resource.
	//
	// Valid values of N: 1 to 5. This indicates that a maximum of five shared resources can be specified at a time.
	//
	// For more information about the types of resources that can be shared, see [Services that work with Resource Sharing](~~450526~~).
	//
	// >  `Resources.N.ResourceId` and `Resources.N.ResourceType` must be used in pairs.
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
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the entities that are associated with the resource share.
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
	// The association status. Valid values:
	//
	// *   Associating: The entity is being associated.
	// *   Associated: The entity is associated.
	// *   Failed: The entity fails to be associated.
	// *   Disassociating: The entity is being disassociated.
	// *   Disassociated: The entity is disassociated.
	//
	// >  The system deletes the records of entities in the `Failed` or `Disassociated` state within 48 hours to 96 hours.
	AssociationStatus *string `json:"AssociationStatus,omitempty" xml:"AssociationStatus,omitempty"`
	// The cause of the association failure.
	AssociationStatusMessage *string `json:"AssociationStatusMessage,omitempty" xml:"AssociationStatusMessage,omitempty"`
	// The association type. Valid values:
	//
	// *   Resource
	// *   Target
	AssociationType *string `json:"AssociationType,omitempty" xml:"AssociationType,omitempty"`
	// The time when the association of the entity was created. The value of this parameter depends on the value of the AssociationType parameter:
	//
	// *   If the value of `AssociationType` is `Resource`, the value of this parameter is the time when the shared resource was associated with the resource share.
	// *   If the value of `AssociationType` is `Target`, the value of this parameter is the time when the principal was associated with the resource share.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the entity. The value of this parameter depends on the value of the AssociationType parameter:
	//
	// *   If the value of `AssociationType` is `Resource`, the value of this parameter is the ID of the shared resource.
	// *   If the value of `AssociationType` is `Target`, the value of this parameter is the ID of the principal.
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The type of the entity. The value of this parameter depends on the value of the AssociationType parameter:
	//
	// *   If the value of AssociationType is Resource, the value of this parameter is the type of the shared resource. For more information about the types of resources that can be shared, see [Services that work with Resource Sharing](~~450526~~).
	// *   If the value of AssociationType is Target, the value of this parameter is `Account`.
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The ID of the resource share.
	ResourceShareId *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
	// The name of the resource share.
	ResourceShareName *string `json:"ResourceShareName,omitempty" xml:"ResourceShareName,omitempty"`
	TargetProperty    *string `json:"TargetProperty,omitempty" xml:"TargetProperty,omitempty"`
	// The time when the association of the entity was updated. The value of this parameter depends on the value of the AssociationType parameter:
	//
	// *   If the value of `AssociationType` is `Resource`, the value of this parameter is the time when the association of the shared resource was updated.
	// *   If the value of `AssociationType` is `Target`, the value of this parameter is the time when the association of the principal was updated.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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

func (s *AssociateResourceShareResponseBodyResourceShareAssociations) SetTargetProperty(v string) *AssociateResourceShareResponseBodyResourceShareAssociations {
	s.TargetProperty = &v
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

type AssociateResourceSharePermissionRequest struct {
	// The name of the permission.
	PermissionName *string `json:"PermissionName,omitempty" xml:"PermissionName,omitempty"`
	// Specifies whether to use the specified permission to replace an existing permission. Valid values:
	//
	// *   false: does not use the specified permission to replace an existing permission. This is the default value. If you set the value to false for a resource share that does not have associated permissions, the system associates the specified permission with the resource share. In a resource share, one resource type can have only one permission. If you set the value to false for a resource share that already has a permission for the resource type indicated by the specified permission, the system reports an error. This prevents you from replacing the existing permission by mistake.
	// *   true: uses the specified permission to replace an existing permission of the same resource type.
	Replace *bool `json:"Replace,omitempty" xml:"Replace,omitempty"`
	// The ID of the resource share.
	ResourceShareId *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
}

func (s AssociateResourceSharePermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s AssociateResourceSharePermissionRequest) GoString() string {
	return s.String()
}

func (s *AssociateResourceSharePermissionRequest) SetPermissionName(v string) *AssociateResourceSharePermissionRequest {
	s.PermissionName = &v
	return s
}

func (s *AssociateResourceSharePermissionRequest) SetReplace(v bool) *AssociateResourceSharePermissionRequest {
	s.Replace = &v
	return s
}

func (s *AssociateResourceSharePermissionRequest) SetResourceShareId(v string) *AssociateResourceSharePermissionRequest {
	s.ResourceShareId = &v
	return s
}

type AssociateResourceSharePermissionResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateResourceSharePermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssociateResourceSharePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateResourceSharePermissionResponseBody) SetRequestId(v string) *AssociateResourceSharePermissionResponseBody {
	s.RequestId = &v
	return s
}

type AssociateResourceSharePermissionResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AssociateResourceSharePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AssociateResourceSharePermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s AssociateResourceSharePermissionResponse) GoString() string {
	return s.String()
}

func (s *AssociateResourceSharePermissionResponse) SetHeaders(v map[string]*string) *AssociateResourceSharePermissionResponse {
	s.Headers = v
	return s
}

func (s *AssociateResourceSharePermissionResponse) SetStatusCode(v int32) *AssociateResourceSharePermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateResourceSharePermissionResponse) SetBody(v *AssociateResourceSharePermissionResponseBody) *AssociateResourceSharePermissionResponse {
	s.Body = v
	return s
}

type ChangeResourceGroupRequest struct {
	// The ID of the destination resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the resource share.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The region ID of the resource share.
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
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

func (s *ChangeResourceGroupRequest) SetResourceRegionId(v string) *ChangeResourceGroupRequest {
	s.ResourceRegionId = &v
	return s
}

type ChangeResourceGroupResponseBody struct {
	// The request ID.
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ChangeResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type CheckSharingWithResourceDirectoryStatusResponseBody struct {
	// Indicates whether resource sharing within a resource directory is enabled. Valid values:
	//
	// *   false
	// *   true
	EnableSharingWithRd *bool `json:"EnableSharingWithRd,omitempty" xml:"EnableSharingWithRd,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckSharingWithResourceDirectoryStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckSharingWithResourceDirectoryStatusResponseBody) GoString() string {
	return s.String()
}

func (s *CheckSharingWithResourceDirectoryStatusResponseBody) SetEnableSharingWithRd(v bool) *CheckSharingWithResourceDirectoryStatusResponseBody {
	s.EnableSharingWithRd = &v
	return s
}

func (s *CheckSharingWithResourceDirectoryStatusResponseBody) SetRequestId(v string) *CheckSharingWithResourceDirectoryStatusResponseBody {
	s.RequestId = &v
	return s
}

type CheckSharingWithResourceDirectoryStatusResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckSharingWithResourceDirectoryStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckSharingWithResourceDirectoryStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckSharingWithResourceDirectoryStatusResponse) GoString() string {
	return s.String()
}

func (s *CheckSharingWithResourceDirectoryStatusResponse) SetHeaders(v map[string]*string) *CheckSharingWithResourceDirectoryStatusResponse {
	s.Headers = v
	return s
}

func (s *CheckSharingWithResourceDirectoryStatusResponse) SetStatusCode(v int32) *CheckSharingWithResourceDirectoryStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckSharingWithResourceDirectoryStatusResponse) SetBody(v *CheckSharingWithResourceDirectoryStatusResponseBody) *CheckSharingWithResourceDirectoryStatusResponse {
	s.Body = v
	return s
}

type CreateResourceShareRequest struct {
	// Specifies whether resources in the resource share can be shared with accounts outside the resource directory. Valid values:
	//
	// *   false (default): Resources in the resource share can be shared only with accounts in the resource directory.
	// *   true: Resources in the resource share can be shared with both accounts in the resource directory and accounts outside the resource directory.
	AllowExternalTargets *bool `json:"AllowExternalTargets,omitempty" xml:"AllowExternalTargets,omitempty"`
	// The information about the permissions. If you do not configure this parameter, the system automatically associates the default permission for the specified resource type with the resource share. For more information, see [Permission library](~~465474~~).
	PermissionNames []*string `json:"PermissionNames,omitempty" xml:"PermissionNames,omitempty" type:"Repeated"`
	// The name of the resource share.
	//
	// The name must be 1 to 50 characters in length.
	//
	// The name can contain letters, digits, periods (.), underscores (\_), and hyphens (-).
	ResourceShareName *string `json:"ResourceShareName,omitempty" xml:"ResourceShareName,omitempty"`
	// The information about the shared resources.
	Resources []*CreateResourceShareRequestResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// The information about the principals.
	Targets []*string `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
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

func (s *CreateResourceShareRequest) SetPermissionNames(v []*string) *CreateResourceShareRequest {
	s.PermissionNames = v
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
	// The ID of a shared resource.
	//
	// Valid values of N: 1 to 5. This indicates that a maximum of five shared resources can be specified at a time.
	//
	// >  `Resources.N.ResourceId` and `Resources.N.ResourceType` must be used in pairs.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of a shared resource.
	//
	// Valid values of N: 1 to 5. This indicates that a maximum of five shared resources can be specified at a time.
	//
	// For more information about the types of resources that can be shared, see [Services that work with Resource Sharing](~~450526~~).
	//
	// >  `Resources.N.ResourceId` and `Resources.N.ResourceType` must be used in pairs.
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
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the resource share.
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
	// Indicates whether resources in the resource share can be shared with accounts outside the resource directory. Valid values:
	//
	// *   false: Resources in the resource share can be shared only with accounts in the resource directory.
	// *   true: Resources in the resource share can be shared with both accounts in the resource directory and accounts outside the resource directory.
	AllowExternalTargets *bool `json:"AllowExternalTargets,omitempty" xml:"AllowExternalTargets,omitempty"`
	// The time when the resource share was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the resource share.
	ResourceShareId *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
	// The name of the resource share.
	ResourceShareName *string `json:"ResourceShareName,omitempty" xml:"ResourceShareName,omitempty"`
	// The owner of the resource share.
	ResourceShareOwner *string `json:"ResourceShareOwner,omitempty" xml:"ResourceShareOwner,omitempty"`
	// The status of the resource share. Valid values:
	//
	// *   Active: The resource share is enabled.
	// *   Pending: The resource share is associated with one or more resource sharing invitations that are waiting for confirmation.
	// *   Deleting: The resource share is being deleted.
	// *   Deleted: The resource share is deleted.
	//
	// >  The system automatically deletes the records of resource shares in the Deleted state within 48 hours to 96 hours after you delete the resource shares.
	ResourceShareStatus *string `json:"ResourceShareStatus,omitempty" xml:"ResourceShareStatus,omitempty"`
	// The time when the resource share was updated.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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
	// The ID of the resource share.
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
	// The ID of the request.
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
	// The supported natural language. Valid values:
	//
	// *   zh-CN: Chinese
	// *   en-US: English
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
	// The information of the regions.
	Regions []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The name of the region.
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The endpoint of the Resource Sharing service in the region.
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The owner of the resource share. Valid values:
	//
	// *   Self: The resource share belongs to the current account. This is the default value. If you are the management account or a member of a resource directory and you want to remove resources or principals from a resource share, set this parameter to Self.
	// *   OtherAccounts: The resource share belongs to another account. If you are not the management account or a member of a resource directory and you want to exit a resource share, set this parameter to OtherAccounts.
	ResourceOwner *string `json:"ResourceOwner,omitempty" xml:"ResourceOwner,omitempty"`
	// The ID of the resource share.
	ResourceShareId *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
	// The shared resource.
	Resources []*DisassociateResourceShareRequestResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// The ID of a principal.
	Targets []*string `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
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
	// The ID of a shared resource.
	//
	// Valid values of N: 1 to 5. This indicates that a maximum of five shared resources can be specified at a time.
	//
	// >  Resources.N.ResourceId and Resources.N.ResourceType must be used in pairs.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of a shared resource.
	//
	// Valid values of N: 1 to 5. This indicates that a maximum of five shared resources can be specified at a time.
	//
	// For more information about the types of resources that can be shared, see [Services that work with Resource Sharing](~~450526~~).
	//
	// >  Resources.N.ResourceId and Resources.N.ResourceType must be used in pairs.
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
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the entities that are associated with the resource share.
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
	// The association status. Valid values:
	//
	// *   Associating: The entity is being associated.
	// *   Associated: The entity is associated.
	// *   Failed: The entity fails to be associated.
	// *   Disassociating: The entity is being disassociated.
	// *   Disassociated: The entity is disassociated.
	//
	// >  The system deletes the records of entities in the `Failed` or `Disassociated` state within 48 hours to 96 hours.
	AssociationStatus *string `json:"AssociationStatus,omitempty" xml:"AssociationStatus,omitempty"`
	// The cause of the disassociation failure.
	AssociationStatusMessage *string `json:"AssociationStatusMessage,omitempty" xml:"AssociationStatusMessage,omitempty"`
	// The association type. Valid values:
	//
	// *   Resource
	// *   Target
	AssociationType *string `json:"AssociationType,omitempty" xml:"AssociationType,omitempty"`
	// The time when the disassociation of the entity was performed. The value of this parameter depends on the value of the AssociationType parameter:
	//
	// *   If the value of `AssociationType` is `Resource`, the value of this parameter is the time when the resource was disassociated from the resource share.
	// *   If the value of `AssociationType` is `Target`, the value of this parameter is the time when the principal was disassociated from the resource share.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the entity. The value of this parameter depends on the value of the AssociationType parameter:
	//
	// *   If the value of `AssociationType` is `Resource`, the value of this parameter is the ID of the resource.
	// *   If the value of `AssociationType` is `Target`, the value of this parameter is the ID of the resource directory, folder, or member.
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The type of the entity. The value of this parameter depends on the value of the AssociationType parameter:
	//
	// *   If the value of AssociationType is Resource, the value of this parameter is the type of the resource. For more information about the types of resources that can be shared, see [Services that work with Resource Sharing](~~450526~~).
	// *   If the value of AssociationType is Target, the value of this parameter is Account.
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The ID of the resource share.
	ResourceShareId *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
	// The name of the resource share.
	ResourceShareName *string `json:"ResourceShareName,omitempty" xml:"ResourceShareName,omitempty"`
	TargetProperty    *string `json:"TargetProperty,omitempty" xml:"TargetProperty,omitempty"`
	// The time when the disassociation of the entity was updated. The value of this parameter depends on the value of the AssociationType parameter:
	//
	// *   If the value of `AssociationType` is `Resource`, the value of this parameter is the time when the disassociation of the resource was updated.
	// *   If the value of `AssociationType` is `Target`, the value of this parameter is the time when the disassociation of the principal was updated.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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

func (s *DisassociateResourceShareResponseBodyResourceShareAssociations) SetTargetProperty(v string) *DisassociateResourceShareResponseBodyResourceShareAssociations {
	s.TargetProperty = &v
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

type DisassociateResourceSharePermissionRequest struct {
	// The name of the permission. For more information, see [Permission library](~~465474~~).
	PermissionName *string `json:"PermissionName,omitempty" xml:"PermissionName,omitempty"`
	// The ID of the resource share.
	ResourceShareId *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
}

func (s DisassociateResourceSharePermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s DisassociateResourceSharePermissionRequest) GoString() string {
	return s.String()
}

func (s *DisassociateResourceSharePermissionRequest) SetPermissionName(v string) *DisassociateResourceSharePermissionRequest {
	s.PermissionName = &v
	return s
}

func (s *DisassociateResourceSharePermissionRequest) SetResourceShareId(v string) *DisassociateResourceSharePermissionRequest {
	s.ResourceShareId = &v
	return s
}

type DisassociateResourceSharePermissionResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisassociateResourceSharePermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisassociateResourceSharePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *DisassociateResourceSharePermissionResponseBody) SetRequestId(v string) *DisassociateResourceSharePermissionResponseBody {
	s.RequestId = &v
	return s
}

type DisassociateResourceSharePermissionResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisassociateResourceSharePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisassociateResourceSharePermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s DisassociateResourceSharePermissionResponse) GoString() string {
	return s.String()
}

func (s *DisassociateResourceSharePermissionResponse) SetHeaders(v map[string]*string) *DisassociateResourceSharePermissionResponse {
	s.Headers = v
	return s
}

func (s *DisassociateResourceSharePermissionResponse) SetStatusCode(v int32) *DisassociateResourceSharePermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *DisassociateResourceSharePermissionResponse) SetBody(v *DisassociateResourceSharePermissionResponseBody) *DisassociateResourceSharePermissionResponse {
	s.Body = v
	return s
}

type EnableSharingWithResourceDirectoryResponseBody struct {
	// The request ID.
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

type GetPermissionRequest struct {
	// The name of the permission.
	PermissionName *string `json:"PermissionName,omitempty" xml:"PermissionName,omitempty"`
	// The version of the permission.
	PermissionVersion *string `json:"PermissionVersion,omitempty" xml:"PermissionVersion,omitempty"`
}

func (s GetPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPermissionRequest) GoString() string {
	return s.String()
}

func (s *GetPermissionRequest) SetPermissionName(v string) *GetPermissionRequest {
	s.PermissionName = &v
	return s
}

func (s *GetPermissionRequest) SetPermissionVersion(v string) *GetPermissionRequest {
	s.PermissionVersion = &v
	return s
}

type GetPermissionResponseBody struct {
	// The information about the permission.
	Permission *GetPermissionResponseBodyPermission `json:"Permission,omitempty" xml:"Permission,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *GetPermissionResponseBody) SetPermission(v *GetPermissionResponseBodyPermission) *GetPermissionResponseBody {
	s.Permission = v
	return s
}

func (s *GetPermissionResponseBody) SetRequestId(v string) *GetPermissionResponseBody {
	s.RequestId = &v
	return s
}

type GetPermissionResponseBodyPermission struct {
	// The creation time.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the permission is the default permission. Valid values:
	//
	// *   false: The permission is not the default permission.
	// *   true: The permission is the default permission.
	DefaultPermission *bool `json:"DefaultPermission,omitempty" xml:"DefaultPermission,omitempty"`
	// Indicates whether the version is the default version. Valid values:
	//
	// *   false: The version is not the default version.
	// *   true: The version is the default version.
	DefaultVersion *bool `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	// The document of the policy related to the permission.
	Permission *string `json:"Permission,omitempty" xml:"Permission,omitempty"`
	// The name of the permission.
	PermissionName *string `json:"PermissionName,omitempty" xml:"PermissionName,omitempty"`
	// The version of the permission.
	PermissionVersion *string `json:"PermissionVersion,omitempty" xml:"PermissionVersion,omitempty"`
	// The type of the shared resources.
	//
	// For more information about the types of resources that can be shared, see [Services that work with Resource Sharing](~~450526~~).
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The update time.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetPermissionResponseBodyPermission) String() string {
	return tea.Prettify(s)
}

func (s GetPermissionResponseBodyPermission) GoString() string {
	return s.String()
}

func (s *GetPermissionResponseBodyPermission) SetCreateTime(v string) *GetPermissionResponseBodyPermission {
	s.CreateTime = &v
	return s
}

func (s *GetPermissionResponseBodyPermission) SetDefaultPermission(v bool) *GetPermissionResponseBodyPermission {
	s.DefaultPermission = &v
	return s
}

func (s *GetPermissionResponseBodyPermission) SetDefaultVersion(v bool) *GetPermissionResponseBodyPermission {
	s.DefaultVersion = &v
	return s
}

func (s *GetPermissionResponseBodyPermission) SetPermission(v string) *GetPermissionResponseBodyPermission {
	s.Permission = &v
	return s
}

func (s *GetPermissionResponseBodyPermission) SetPermissionName(v string) *GetPermissionResponseBodyPermission {
	s.PermissionName = &v
	return s
}

func (s *GetPermissionResponseBodyPermission) SetPermissionVersion(v string) *GetPermissionResponseBodyPermission {
	s.PermissionVersion = &v
	return s
}

func (s *GetPermissionResponseBodyPermission) SetResourceType(v string) *GetPermissionResponseBodyPermission {
	s.ResourceType = &v
	return s
}

func (s *GetPermissionResponseBodyPermission) SetUpdateTime(v string) *GetPermissionResponseBodyPermission {
	s.UpdateTime = &v
	return s
}

type GetPermissionResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPermissionResponse) GoString() string {
	return s.String()
}

func (s *GetPermissionResponse) SetHeaders(v map[string]*string) *GetPermissionResponse {
	s.Headers = v
	return s
}

func (s *GetPermissionResponse) SetStatusCode(v int32) *GetPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPermissionResponse) SetBody(v *GetPermissionResponseBody) *GetPermissionResponse {
	s.Body = v
	return s
}

type ListPermissionVersionsRequest struct {
	// The maximum number of entries to return for a single request.
	//
	// Valid values: 1 to 100. Default value: 20.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The `token` that is used to initiate the next request. If the response of the current request is truncated, you can use the token to initiate another request and obtain the remaining records.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The name of the permission.
	PermissionName *string `json:"PermissionName,omitempty" xml:"PermissionName,omitempty"`
}

func (s ListPermissionVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPermissionVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListPermissionVersionsRequest) SetMaxResults(v int32) *ListPermissionVersionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPermissionVersionsRequest) SetNextToken(v string) *ListPermissionVersionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListPermissionVersionsRequest) SetPermissionName(v string) *ListPermissionVersionsRequest {
	s.PermissionName = &v
	return s
}

type ListPermissionVersionsResponseBody struct {
	// The token that is used to initiate the next request. If the response of the current request is truncated, you can use the token to initiate another request and obtain the remaining records.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The information about the permission.
	Permissions []*ListPermissionVersionsResponseBodyPermissions `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPermissionVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPermissionVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPermissionVersionsResponseBody) SetNextToken(v string) *ListPermissionVersionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPermissionVersionsResponseBody) SetPermissions(v []*ListPermissionVersionsResponseBodyPermissions) *ListPermissionVersionsResponseBody {
	s.Permissions = v
	return s
}

func (s *ListPermissionVersionsResponseBody) SetRequestId(v string) *ListPermissionVersionsResponseBody {
	s.RequestId = &v
	return s
}

type ListPermissionVersionsResponseBodyPermissions struct {
	// The creation time.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the permission is the default permission. Valid values:
	//
	// *   false: The permission is not the default permission.
	// *   true: The permission is the default permission.
	DefaultPermission *bool `json:"DefaultPermission,omitempty" xml:"DefaultPermission,omitempty"`
	// Indicates whether the version is the default version. Valid values:
	//
	// *   false: The version is not the default version.
	// *   true: The version is the default version.
	DefaultVersion *bool `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	// The name of the permission.
	PermissionName *string `json:"PermissionName,omitempty" xml:"PermissionName,omitempty"`
	// The version of the permission.
	PermissionVersion *string `json:"PermissionVersion,omitempty" xml:"PermissionVersion,omitempty"`
	// The type of the shared resources.
	//
	// For more information about the types of resources that can be shared, see [Services that work with Resource Sharing](~~450526~~).
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The update time.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListPermissionVersionsResponseBodyPermissions) String() string {
	return tea.Prettify(s)
}

func (s ListPermissionVersionsResponseBodyPermissions) GoString() string {
	return s.String()
}

func (s *ListPermissionVersionsResponseBodyPermissions) SetCreateTime(v string) *ListPermissionVersionsResponseBodyPermissions {
	s.CreateTime = &v
	return s
}

func (s *ListPermissionVersionsResponseBodyPermissions) SetDefaultPermission(v bool) *ListPermissionVersionsResponseBodyPermissions {
	s.DefaultPermission = &v
	return s
}

func (s *ListPermissionVersionsResponseBodyPermissions) SetDefaultVersion(v bool) *ListPermissionVersionsResponseBodyPermissions {
	s.DefaultVersion = &v
	return s
}

func (s *ListPermissionVersionsResponseBodyPermissions) SetPermissionName(v string) *ListPermissionVersionsResponseBodyPermissions {
	s.PermissionName = &v
	return s
}

func (s *ListPermissionVersionsResponseBodyPermissions) SetPermissionVersion(v string) *ListPermissionVersionsResponseBodyPermissions {
	s.PermissionVersion = &v
	return s
}

func (s *ListPermissionVersionsResponseBodyPermissions) SetResourceType(v string) *ListPermissionVersionsResponseBodyPermissions {
	s.ResourceType = &v
	return s
}

func (s *ListPermissionVersionsResponseBodyPermissions) SetUpdateTime(v string) *ListPermissionVersionsResponseBodyPermissions {
	s.UpdateTime = &v
	return s
}

type ListPermissionVersionsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPermissionVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPermissionVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPermissionVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListPermissionVersionsResponse) SetHeaders(v map[string]*string) *ListPermissionVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListPermissionVersionsResponse) SetStatusCode(v int32) *ListPermissionVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPermissionVersionsResponse) SetBody(v *ListPermissionVersionsResponseBody) *ListPermissionVersionsResponse {
	s.Body = v
	return s
}

type ListPermissionsRequest struct {
	// The maximum number of entries to return for a single request.
	//
	// Valid values: 1 to 100. Default value: 20.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The `token` that is used to initiate the next request. If the response of the current request is truncated, you can use the token to initiate another request and obtain the remaining records.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The type of the shared resources.
	//
	// For more information about the types of resources that can be shared, see [Services that work with Resource Sharing](~~450526~~).
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListPermissionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPermissionsRequest) GoString() string {
	return s.String()
}

func (s *ListPermissionsRequest) SetMaxResults(v int32) *ListPermissionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPermissionsRequest) SetNextToken(v string) *ListPermissionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListPermissionsRequest) SetResourceType(v string) *ListPermissionsRequest {
	s.ResourceType = &v
	return s
}

type ListPermissionsResponseBody struct {
	// The token that is used to initiate the next request. If the response of the current request is truncated, you can use the token to initiate another request and obtain the remaining records.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The information about the permission.
	Permissions []*ListPermissionsResponseBodyPermissions `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPermissionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPermissionsResponseBody) SetNextToken(v string) *ListPermissionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPermissionsResponseBody) SetPermissions(v []*ListPermissionsResponseBodyPermissions) *ListPermissionsResponseBody {
	s.Permissions = v
	return s
}

func (s *ListPermissionsResponseBody) SetRequestId(v string) *ListPermissionsResponseBody {
	s.RequestId = &v
	return s
}

type ListPermissionsResponseBodyPermissions struct {
	// The creation time.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the permission is the default permission. Valid values:
	//
	// *   false: The permission is not the default permission.
	// *   true: The permission is the default permission.
	DefaultPermission *bool `json:"DefaultPermission,omitempty" xml:"DefaultPermission,omitempty"`
	// Indicates whether the version is the default version. Valid values:
	//
	// *   false: The version is not the default version.
	// *   true: The version is the default version.
	DefaultVersion *bool `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	// The name of the permission.
	PermissionName *string `json:"PermissionName,omitempty" xml:"PermissionName,omitempty"`
	// The version of the permission.
	PermissionVersion *string `json:"PermissionVersion,omitempty" xml:"PermissionVersion,omitempty"`
	// The type of the shared resources.
	//
	// For more information about the types of resources that can be shared, see [Services that work with Resource Sharing](~~450526~~).
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The update time.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListPermissionsResponseBodyPermissions) String() string {
	return tea.Prettify(s)
}

func (s ListPermissionsResponseBodyPermissions) GoString() string {
	return s.String()
}

func (s *ListPermissionsResponseBodyPermissions) SetCreateTime(v string) *ListPermissionsResponseBodyPermissions {
	s.CreateTime = &v
	return s
}

func (s *ListPermissionsResponseBodyPermissions) SetDefaultPermission(v bool) *ListPermissionsResponseBodyPermissions {
	s.DefaultPermission = &v
	return s
}

func (s *ListPermissionsResponseBodyPermissions) SetDefaultVersion(v bool) *ListPermissionsResponseBodyPermissions {
	s.DefaultVersion = &v
	return s
}

func (s *ListPermissionsResponseBodyPermissions) SetPermissionName(v string) *ListPermissionsResponseBodyPermissions {
	s.PermissionName = &v
	return s
}

func (s *ListPermissionsResponseBodyPermissions) SetPermissionVersion(v string) *ListPermissionsResponseBodyPermissions {
	s.PermissionVersion = &v
	return s
}

func (s *ListPermissionsResponseBodyPermissions) SetResourceType(v string) *ListPermissionsResponseBodyPermissions {
	s.ResourceType = &v
	return s
}

func (s *ListPermissionsResponseBodyPermissions) SetUpdateTime(v string) *ListPermissionsResponseBodyPermissions {
	s.UpdateTime = &v
	return s
}

type ListPermissionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPermissionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPermissionsResponse) GoString() string {
	return s.String()
}

func (s *ListPermissionsResponse) SetHeaders(v map[string]*string) *ListPermissionsResponse {
	s.Headers = v
	return s
}

func (s *ListPermissionsResponse) SetStatusCode(v int32) *ListPermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPermissionsResponse) SetBody(v *ListPermissionsResponseBody) *ListPermissionsResponse {
	s.Body = v
	return s
}

type ListResourceShareAssociationsRequest struct {
	// The association status. Valid values:
	//
	// *   Associating: The entity is being associated.
	// *   Associated: The entity is associated.
	// *   Failed: The entity fails to be associated.
	// *   Disassociating: The entity is being disassociated.
	// *   Disassociated: The entity is disassociated.
	//
	// >  The system deletes the records of entities in the `Failed` or `Disassociated` state within 48 hours to 96 hours.
	AssociationStatus *string `json:"AssociationStatus,omitempty" xml:"AssociationStatus,omitempty"`
	// The association type. Valid values:
	//
	// *   Resource
	// *   Target
	AssociationType *string `json:"AssociationType,omitempty" xml:"AssociationType,omitempty"`
	// The maximum number of entries to return for a single request.
	//
	// Valid values: 1 to 100. Default value: 20.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The `token` that is used to initiate the next request. If the response of the current request is truncated, you can use the token to initiate another request and obtain the remaining records.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the resource.
	//
	// >  This parameter is unavailable if you set the `AssociationType` parameter to `Target`.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The IDs of the resource shares.
	ResourceShareIds []*string `json:"ResourceShareIds,omitempty" xml:"ResourceShareIds,omitempty" type:"Repeated"`
	// The ID of the principal.
	//
	// >  This parameter is unavailable if you set the `AssociationType` parameter to `Resource`.
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
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
	// The `token` that is used to initiate the next request. If the response of the current request is truncated, you can use the token to initiate another request and obtain the remaining records.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the entities.
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
	AssociationFailedDetails []*ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails `json:"AssociationFailedDetails,omitempty" xml:"AssociationFailedDetails,omitempty" type:"Repeated"`
	// The association status. Valid values:
	//
	// *   Associating: The entity is being associated.
	// *   Associated: The entity is associated.
	// *   Failed: The entity fails to be associated.
	// *   Disassociating: The entity is being disassociated.
	// *   Disassociated: The entity is disassociated.
	//
	// >  The system deletes the records of entities in the `Failed` or `Disassociated` state within 48 hours to 96 hours.
	AssociationStatus *string `json:"AssociationStatus,omitempty" xml:"AssociationStatus,omitempty"`
	// The cause of the association failure.
	AssociationStatusMessage *string `json:"AssociationStatusMessage,omitempty" xml:"AssociationStatusMessage,omitempty"`
	// The association type. Valid values:
	//
	// *   Resource
	// *   Target
	AssociationType *string `json:"AssociationType,omitempty" xml:"AssociationType,omitempty"`
	// The time when the association of the entity was created. The value of this parameter depends on the value of the AssociationType parameter:
	//
	// *   If the value of `AssociationType` is `Resource`, the value of this parameter is the time when the shared resource was associated with or disassociated from the resource share.
	// *   If the value of `AssociationType` is `Target`, the value of this parameter is the time when the principal was associated with or disassociated from the resource share.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the entity. The value of this parameter depends on the value of the AssociationType parameter:
	//
	// *   If the value of `AssociationType` is `Resource`, the value of this parameter is the ID of the shared resource.
	// *   If the value of `AssociationType` is `Target`, the value of this parameter is the ID of the principal.
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The type of the entity. The value of this parameter depends on the value of the AssociationType parameter:
	//
	// *   If the value of AssociationType is Resource, the value of this parameter is the type of the resource. For more information about the types of resources that can be shared, see [Services that work with Resource Sharing](~~450526~~).
	// *   If the value of AssociationType is Target, the value of this parameter is `Account`.
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// Indicates whether the principal is outside the resource directory. Valid values:
	//
	// *   true: The principal is outside the resource directory.
	// *   false: The principal is in the resource directory.
	External *bool `json:"External,omitempty" xml:"External,omitempty"`
	// The ID of the resource share.
	ResourceShareId *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
	// The name of the resource share.
	ResourceShareName *string `json:"ResourceShareName,omitempty" xml:"ResourceShareName,omitempty"`
	TargetProperty    *string `json:"TargetProperty,omitempty" xml:"TargetProperty,omitempty"`
	// The time when the association of the entity was updated. The value of this parameter depends on the value of the AssociationType parameter:
	//
	// *   If the value of `AssociationType` is `Resource`, the value of this parameter is the time when the association of the shared resource was updated.
	// *   If the value of `AssociationType` is `Target`, the value of this parameter is the time when the association of the principal was updated.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListResourceShareAssociationsResponseBodyResourceShareAssociations) String() string {
	return tea.Prettify(s)
}

func (s ListResourceShareAssociationsResponseBodyResourceShareAssociations) GoString() string {
	return s.String()
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociations) SetAssociationFailedDetails(v []*ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails) *ListResourceShareAssociationsResponseBodyResourceShareAssociations {
	s.AssociationFailedDetails = v
	return s
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

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociations) SetTargetProperty(v string) *ListResourceShareAssociationsResponseBodyResourceShareAssociations {
	s.TargetProperty = &v
	return s
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociations) SetUpdateTime(v string) *ListResourceShareAssociationsResponseBodyResourceShareAssociations {
	s.UpdateTime = &v
	return s
}

type ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails struct {
	AssociateType *string `json:"AssociateType,omitempty" xml:"AssociateType,omitempty"`
	EntityId      *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	EntityType    *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusMessage *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
}

func (s ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails) String() string {
	return tea.Prettify(s)
}

func (s ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails) GoString() string {
	return s.String()
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails) SetAssociateType(v string) *ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails {
	s.AssociateType = &v
	return s
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails) SetEntityId(v string) *ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails {
	s.EntityId = &v
	return s
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails) SetEntityType(v string) *ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails {
	s.EntityType = &v
	return s
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails) SetStatus(v string) *ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails {
	s.Status = &v
	return s
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails) SetStatusMessage(v string) *ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails {
	s.StatusMessage = &v
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
	// The maximum number of entries to return for a single request.
	//
	// Valid values: 1 to 100. Default value: 20.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of `NextToken`.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The IDs of the resource shares.
	ResourceShareIds []*string `json:"ResourceShareIds,omitempty" xml:"ResourceShareIds,omitempty" type:"Repeated"`
	// The IDs of the resource sharing invitations.
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
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of `NextToken`.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the resource sharing invitations.
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
	// The time when the invitation was created. The time is displayed in UTC.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The information about the failure.
	InvitationFailedDetails []*ListResourceShareInvitationsResponseBodyResourceShareInvitationsInvitationFailedDetails `json:"InvitationFailedDetails,omitempty" xml:"InvitationFailedDetails,omitempty" type:"Repeated"`
	// The Alibaba Cloud account ID of the invitee.
	ReceiverAccountId *string `json:"ReceiverAccountId,omitempty" xml:"ReceiverAccountId,omitempty"`
	// The ID of the resource share.
	ResourceShareId *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
	// The ID of the invitation.
	ResourceShareInvitationId *string `json:"ResourceShareInvitationId,omitempty" xml:"ResourceShareInvitationId,omitempty"`
	// The name of the resource share.
	ResourceShareName *string `json:"ResourceShareName,omitempty" xml:"ResourceShareName,omitempty"`
	// The Alibaba Cloud account ID of the inviter.
	SenderAccountId *string `json:"SenderAccountId,omitempty" xml:"SenderAccountId,omitempty"`
	// The status of the invitation. Valid values:
	//
	// *   Pending
	// *   Accepted
	// *   Cancelled
	// *   Rejected
	// *   Expired
	// *   AcceptFailed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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

func (s *ListResourceShareInvitationsResponseBodyResourceShareInvitations) SetInvitationFailedDetails(v []*ListResourceShareInvitationsResponseBodyResourceShareInvitationsInvitationFailedDetails) *ListResourceShareInvitationsResponseBodyResourceShareInvitations {
	s.InvitationFailedDetails = v
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

type ListResourceShareInvitationsResponseBodyResourceShareInvitationsInvitationFailedDetails struct {
	// The type of the sharing operation. Valid values:
	//
	// *   Associate
	// *   Disassociate
	AssociateType *string `json:"AssociateType,omitempty" xml:"AssociateType,omitempty"`
	// The ID of the shared resource.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the shared resource.
	//
	// For more information about the types of resources that can be shared, see [Services that work with Resource Sharing](~~450526~~).
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The failure status. Valid values:
	//
	// *   Unavailable: The resource cannot be shared.
	// *   LimitExceeded: The number of shared resources within the Alibaba Cloud account exceeds the upper limit.
	// *   ZonalResourceInaccessible: The resource is unavailable in this region.
	// *   UnsupportedOperation: The operation is not allowed because another association exists.
	// *   InternalError: An internal error occurred during the check.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The failure cause.
	StatusMessage *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
}

func (s ListResourceShareInvitationsResponseBodyResourceShareInvitationsInvitationFailedDetails) String() string {
	return tea.Prettify(s)
}

func (s ListResourceShareInvitationsResponseBodyResourceShareInvitationsInvitationFailedDetails) GoString() string {
	return s.String()
}

func (s *ListResourceShareInvitationsResponseBodyResourceShareInvitationsInvitationFailedDetails) SetAssociateType(v string) *ListResourceShareInvitationsResponseBodyResourceShareInvitationsInvitationFailedDetails {
	s.AssociateType = &v
	return s
}

func (s *ListResourceShareInvitationsResponseBodyResourceShareInvitationsInvitationFailedDetails) SetResourceId(v string) *ListResourceShareInvitationsResponseBodyResourceShareInvitationsInvitationFailedDetails {
	s.ResourceId = &v
	return s
}

func (s *ListResourceShareInvitationsResponseBodyResourceShareInvitationsInvitationFailedDetails) SetResourceType(v string) *ListResourceShareInvitationsResponseBodyResourceShareInvitationsInvitationFailedDetails {
	s.ResourceType = &v
	return s
}

func (s *ListResourceShareInvitationsResponseBodyResourceShareInvitationsInvitationFailedDetails) SetStatus(v string) *ListResourceShareInvitationsResponseBodyResourceShareInvitationsInvitationFailedDetails {
	s.Status = &v
	return s
}

func (s *ListResourceShareInvitationsResponseBodyResourceShareInvitationsInvitationFailedDetails) SetStatusMessage(v string) *ListResourceShareInvitationsResponseBodyResourceShareInvitationsInvitationFailedDetails {
	s.StatusMessage = &v
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

type ListResourceSharePermissionsRequest struct {
	// The maximum number of entries to return for a single request.
	//
	// Valid values: 1 to 100. Default value: 20.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The `token` that is used to initiate the next request. If the response of the current request is truncated, you can use the token to initiate another request and obtain the remaining records.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The owner of the resource share. Valid values:
	//
	// *   Self: the current account
	// *   OtherAccounts: an account other than the current account
	ResourceOwner *string `json:"ResourceOwner,omitempty" xml:"ResourceOwner,omitempty"`
	// The ID of the resource share.
	ResourceShareId *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
}

func (s ListResourceSharePermissionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourceSharePermissionsRequest) GoString() string {
	return s.String()
}

func (s *ListResourceSharePermissionsRequest) SetMaxResults(v int32) *ListResourceSharePermissionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListResourceSharePermissionsRequest) SetNextToken(v string) *ListResourceSharePermissionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListResourceSharePermissionsRequest) SetResourceOwner(v string) *ListResourceSharePermissionsRequest {
	s.ResourceOwner = &v
	return s
}

func (s *ListResourceSharePermissionsRequest) SetResourceShareId(v string) *ListResourceSharePermissionsRequest {
	s.ResourceShareId = &v
	return s
}

type ListResourceSharePermissionsResponseBody struct {
	// The `token` that is used to initiate the next request. If the response of the current request is truncated, you can use the token to initiate another request and obtain the remaining records.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The information about the permissions.
	Permissions []*ListResourceSharePermissionsResponseBodyPermissions `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListResourceSharePermissionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourceSharePermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceSharePermissionsResponseBody) SetNextToken(v string) *ListResourceSharePermissionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListResourceSharePermissionsResponseBody) SetPermissions(v []*ListResourceSharePermissionsResponseBodyPermissions) *ListResourceSharePermissionsResponseBody {
	s.Permissions = v
	return s
}

func (s *ListResourceSharePermissionsResponseBody) SetRequestId(v string) *ListResourceSharePermissionsResponseBody {
	s.RequestId = &v
	return s
}

type ListResourceSharePermissionsResponseBodyPermissions struct {
	// The creation time.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the permission is the default permission. Valid values:
	//
	// *   false: The permission is not the default permission.
	// *   true: The permission is the default permission.
	DefaultPermission *bool `json:"DefaultPermission,omitempty" xml:"DefaultPermission,omitempty"`
	// Indicates whether the version is the default version. Valid values:
	//
	// *   false: The version is not the default version.
	// *   true: The version is the default version.
	DefaultVersion *bool `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	// The name of the permission.
	PermissionName *string `json:"PermissionName,omitempty" xml:"PermissionName,omitempty"`
	// The version of the permission.
	PermissionVersion *string `json:"PermissionVersion,omitempty" xml:"PermissionVersion,omitempty"`
	// The type of the shared resources.
	//
	// For more information about the types of resources that can be shared, see [Services that work with Resource Sharing](~~450526~~).
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The update time.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListResourceSharePermissionsResponseBodyPermissions) String() string {
	return tea.Prettify(s)
}

func (s ListResourceSharePermissionsResponseBodyPermissions) GoString() string {
	return s.String()
}

func (s *ListResourceSharePermissionsResponseBodyPermissions) SetCreateTime(v string) *ListResourceSharePermissionsResponseBodyPermissions {
	s.CreateTime = &v
	return s
}

func (s *ListResourceSharePermissionsResponseBodyPermissions) SetDefaultPermission(v bool) *ListResourceSharePermissionsResponseBodyPermissions {
	s.DefaultPermission = &v
	return s
}

func (s *ListResourceSharePermissionsResponseBodyPermissions) SetDefaultVersion(v bool) *ListResourceSharePermissionsResponseBodyPermissions {
	s.DefaultVersion = &v
	return s
}

func (s *ListResourceSharePermissionsResponseBodyPermissions) SetPermissionName(v string) *ListResourceSharePermissionsResponseBodyPermissions {
	s.PermissionName = &v
	return s
}

func (s *ListResourceSharePermissionsResponseBodyPermissions) SetPermissionVersion(v string) *ListResourceSharePermissionsResponseBodyPermissions {
	s.PermissionVersion = &v
	return s
}

func (s *ListResourceSharePermissionsResponseBodyPermissions) SetResourceType(v string) *ListResourceSharePermissionsResponseBodyPermissions {
	s.ResourceType = &v
	return s
}

func (s *ListResourceSharePermissionsResponseBodyPermissions) SetUpdateTime(v string) *ListResourceSharePermissionsResponseBodyPermissions {
	s.UpdateTime = &v
	return s
}

type ListResourceSharePermissionsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListResourceSharePermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListResourceSharePermissionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourceSharePermissionsResponse) GoString() string {
	return s.String()
}

func (s *ListResourceSharePermissionsResponse) SetHeaders(v map[string]*string) *ListResourceSharePermissionsResponse {
	s.Headers = v
	return s
}

func (s *ListResourceSharePermissionsResponse) SetStatusCode(v int32) *ListResourceSharePermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceSharePermissionsResponse) SetBody(v *ListResourceSharePermissionsResponseBody) *ListResourceSharePermissionsResponse {
	s.Body = v
	return s
}

type ListResourceSharesRequest struct {
	// The maximum number of entries to return for a single request.
	//
	// Valid values: 1 to 100. Default value: 20.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The `token` that is used to initiate the next request. If the response of the current request is truncated, you can use the token to initiate another request and obtain the remaining records.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The name of the permission. For more information, see [Permission library](~~465474~~).
	PermissionName *string `json:"PermissionName,omitempty" xml:"PermissionName,omitempty"`
	// The owner of the resource shares. Valid values:
	//
	// *   Self: the current account
	// *   OtherAccounts: an account other than the current account
	ResourceOwner *string `json:"ResourceOwner,omitempty" xml:"ResourceOwner,omitempty"`
	// The ID of a resource share.
	ResourceShareIds []*string `json:"ResourceShareIds,omitempty" xml:"ResourceShareIds,omitempty" type:"Repeated"`
	// The name of the resource share.
	ResourceShareName *string `json:"ResourceShareName,omitempty" xml:"ResourceShareName,omitempty"`
	// The status of the resource share. Valid values:
	//
	// *   Active: The resource share is enabled.
	// *   Pending: The resource share is associated with one or more resource sharing invitations that are waiting for confirmation.
	// *   Deleting: The resource share is being deleted.
	// *   Deleted: The resource share is deleted.
	//
	// >  The system deletes the records of resource shares in the Deleted state within 48 hours to 96 hours after you delete the resource shares.
	ResourceShareStatus *string `json:"ResourceShareStatus,omitempty" xml:"ResourceShareStatus,omitempty"`
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

func (s *ListResourceSharesRequest) SetPermissionName(v string) *ListResourceSharesRequest {
	s.PermissionName = &v
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
	// The `token` that is used to initiate the next request. If the response of the current request is truncated, you can use the token to initiate another request and obtain the remaining records.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the resource shares.
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
	// Indicates whether resources in the resource share can be shared with accounts outside the resource directory. Valid values:
	//
	// *   false: Resources in the resource share can be shared only with accounts in the resource directory.
	// *   true: Resources in the resource share can be shared with both accounts in the resource directory and accounts outside the resource directory.
	AllowExternalTargets *bool `json:"AllowExternalTargets,omitempty" xml:"AllowExternalTargets,omitempty"`
	// The time when the resource share was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the resource share.
	ResourceShareId *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
	// The name of the resource share.
	ResourceShareName *string `json:"ResourceShareName,omitempty" xml:"ResourceShareName,omitempty"`
	// The owner of the resource share.
	ResourceShareOwner *string `json:"ResourceShareOwner,omitempty" xml:"ResourceShareOwner,omitempty"`
	// The status of the resource share. Valid values:
	//
	// *   Active: The resource share is enabled.
	// *   Pending: The resource share is associated with one or more resource sharing invitations that are waiting for confirmation.
	// *   Deleting: The resource share is being deleted.
	// *   Deleted: The resource share is deleted.
	//
	// >  The system deletes the records of resource shares in the Deleted state within 48 hours to 96 hours after you delete the resource shares.
	ResourceShareStatus *string `json:"ResourceShareStatus,omitempty" xml:"ResourceShareStatus,omitempty"`
	// The time when the resource share was updated.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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
	// The maximum number of entries to return for a single request.
	//
	// Valid values: 1 to 100. Default value: 20.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The `token` that is used to initiate the next request. If the response of the current request is truncated, you can use the token to initiate another request and obtain the remaining records.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of a shared resource.
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// The owner of the resource shares. Valid values:
	//
	// *   Self: your account. If you set the value to Self, the resources you share with other accounts are queried.
	// *   OtherAccounts: another account. If you set the value to OtherAccounts, the resources other accounts share with you are queried.
	ResourceOwner *string `json:"ResourceOwner,omitempty" xml:"ResourceOwner,omitempty"`
	// The ID of a resource share.
	ResourceShareIds []*string `json:"ResourceShareIds,omitempty" xml:"ResourceShareIds,omitempty" type:"Repeated"`
	// The type of the shared resources.
	//
	// For more information about the types of resources that can be shared, see [Services that work with Resource Sharing](~~450526~~).
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The ID of the principal or resource owner.
	//
	// *   If the value of `ResourceOwner` is `Self`, set this parameter to the ID of a principal.
	// *   If the value of `ResourceOwner` is `OtherAccounts`, set this parameter to the ID of a resource owner.
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
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
	// The token that is used to initiate the next request. If the response of the current request is truncated, you can use the token to initiate another request and obtain the remaining records.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the shared resources.
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
	// The time when the shared resource was associated with the resource share.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the shared resource.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The ID of the resource share.
	ResourceShareId *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
	// The status of the shared resource. This parameter is returned only when you query the resources that other accounts share with you.
	//
	// Valid values:
	//
	// *   Available: The resource is available.
	// *   ZonalResourceInaccessible: The resource is unavailable in the current zone.
	// *   LimitExceeded: The resource is unavailable because the maximum number of resources that other accounts can share with you exceeds the upper limit.
	// *   Unavailable: The resource is unavailable.
	ResourceStatus *string `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty"`
	// The cause of the association failure.
	ResourceStatusMessage *string `json:"ResourceStatusMessage,omitempty" xml:"ResourceStatusMessage,omitempty"`
	// The type of the shared resource.
	//
	// For more information about the types of resources that can be shared, see [Services that work with Resource Sharing](~~450526~~).
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The time when the association of the shared resource was updated.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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
	// The maximum number of entries to return for a single request.
	//
	// Valid values: 1 to 100. Default value: 20.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The `token` that is used to initiate the next request. If the response of the current request is truncated, you can use the token to initiate another request and obtain the remaining records.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the shared resource.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The owner of the resource share.
	//
	// *   Self: your account. If you set the value to Self, the principals that are associated with your resource shares are queried.
	// *   OtherAccounts: another account. If you set the value to OtherAccounts, the resource shares with which your account is associated and the owners of the resource shares are queried.
	ResourceOwner *string `json:"ResourceOwner,omitempty" xml:"ResourceOwner,omitempty"`
	// The ID of a resource share.
	ResourceShareIds []*string `json:"ResourceShareIds,omitempty" xml:"ResourceShareIds,omitempty" type:"Repeated"`
	// The type of the shared resources.
	//
	// For more information about the types of resources that can be shared, see [Services that work with Resource Sharing](~~450526~~).
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The ID of a principal.
	Targets []*string `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
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
	// The `token` that is used to initiate the next request. If the response of the current request is truncated, you can use the token to initiate another request and obtain the remaining records.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the principals.
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
	// The time when the principal was associated with the resource share.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the principal is outside the resource directory. Valid values:
	//
	// *   true: The principal is outside the resource directory.
	// *   false: The principal is in the resource directory.
	External *bool `json:"External,omitempty" xml:"External,omitempty"`
	// The ID of the resource share.
	ResourceShareId *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
	// The ID of the principal or resource owner.
	//
	// *   If the value of `ResourceOwner` is `Self`, the value of this parameter is the ID of a principal.
	// *   If the value of `ResourceOwner` is `OtherAccounts`, the value of this parameter is the ID of a resource owner.
	TargetId       *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	TargetProperty *string `json:"TargetProperty,omitempty" xml:"TargetProperty,omitempty"`
	// The time when the association of the principal was updated.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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

func (s *ListSharedTargetsResponseBodySharedTargets) SetTargetProperty(v string) *ListSharedTargetsResponseBodySharedTargets {
	s.TargetProperty = &v
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
	// The ID of the resource sharing invitation.
	//
	// You can call the [ListResourceShareInvitations](~~450564~~) operation to obtain the ID of a resource sharing invitation.
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
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the resource sharing invitation.
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
	// The time when the invitation was created. The time is displayed in UTC.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The Alibaba Cloud account ID of the invitee.
	ReceiverAccountId *string `json:"ReceiverAccountId,omitempty" xml:"ReceiverAccountId,omitempty"`
	// The ID of the resource share.
	ResourceShareId *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
	// The ID of the invitation.
	ResourceShareInvitationId *string `json:"ResourceShareInvitationId,omitempty" xml:"ResourceShareInvitationId,omitempty"`
	// The name of the resource share.
	ResourceShareName *string `json:"ResourceShareName,omitempty" xml:"ResourceShareName,omitempty"`
	// The Alibaba Cloud account ID of the inviter.
	SenderAccountId *string `json:"SenderAccountId,omitempty" xml:"SenderAccountId,omitempty"`
	// The status of the invitation. Valid values:
	//
	// *   Pending: The invitation is waiting for confirmation.
	// *   Accepted: The invitation is accepted.
	// *   Cancelled: The invitation is canceled.
	// *   Rejected: The invitation is rejected.
	// *   Expired: The invitation has expired.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	// Specifies whether resources in the resource share can be shared with accounts outside the resource directory. Valid values:
	//
	// *   false: Resources in the resource share can be shared only with accounts in the resource directory.
	// *   true: Resources in the resource share can be shared with both accounts in the resource directory and accounts outside the resource directory.
	AllowExternalTargets *bool `json:"AllowExternalTargets,omitempty" xml:"AllowExternalTargets,omitempty"`
	// The ID of the resource share.
	ResourceShareId *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
	// The new name of the resource share.
	//
	// The name must be 1 to 50 characters in length.
	//
	// The name can contain letters, digits, periods (.), underscores (\_), and hyphens (-).
	ResourceShareName *string `json:"ResourceShareName,omitempty" xml:"ResourceShareName,omitempty"`
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
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the resource share.
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
	// Indicates whether resources in the resource share can be shared with accounts outside the resource directory. Valid values:
	//
	// *   false: Resources in the resource share can be shared only with accounts in the resource directory.
	// *   true: Resources in the resource share can be shared with both accounts in the resource directory and accounts outside the resource directory.
	AllowExternalTargets *bool `json:"AllowExternalTargets,omitempty" xml:"AllowExternalTargets,omitempty"`
	// The time when the resource share was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the resource share.
	ResourceShareId *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
	// The name of the resource share.
	ResourceShareName *string `json:"ResourceShareName,omitempty" xml:"ResourceShareName,omitempty"`
	// The owner of the resource share.
	ResourceShareOwner *string `json:"ResourceShareOwner,omitempty" xml:"ResourceShareOwner,omitempty"`
	// The status of the resource share. Valid values:
	//
	// *   Active: The resource share is enabled.
	// *   Pending: The resource share is associated with one or more resource sharing invitations that are waiting for confirmation.
	// *   Deleting: The resource share is being deleted.
	// *   Deleted: The resource share is deleted.
	//
	// >  The system deletes the records of resource shares in the Deleted state within 48 hours to 96 hours after you delete the resource shares.
	ResourceShareStatus *string `json:"ResourceShareStatus,omitempty" xml:"ResourceShareStatus,omitempty"`
	// The time when the resource share was updated.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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

/**
 * ### [](#)
 * *   A principal needs to accept or reject a resource sharing invitation only if the principal is not the management account or a member of a resource directory. If you share resources with an object in a resource directory, the system automatically accepts the resource sharing invitation for the object.
 * *   A resource sharing invitation is valid for seven days. A principal must accept or reject the invitation within the validity period.
 * This topic provides an example on how to call the API operation to accept the resource sharing invitation whose ID is `i-pMnItMX19fBJ****` in the `cn-hangzhou` region.
 *
 * @param request AcceptResourceShareInvitationRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return AcceptResourceShareInvitationResponse
 */
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

/**
 * ### [](#)
 * *   A principal needs to accept or reject a resource sharing invitation only if the principal is not the management account or a member of a resource directory. If you share resources with an object in a resource directory, the system automatically accepts the resource sharing invitation for the object.
 * *   A resource sharing invitation is valid for seven days. A principal must accept or reject the invitation within the validity period.
 * This topic provides an example on how to call the API operation to accept the resource sharing invitation whose ID is `i-pMnItMX19fBJ****` in the `cn-hangzhou` region.
 *
 * @param request AcceptResourceShareInvitationRequest
 * @return AcceptResourceShareInvitationResponse
 */
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

/**
 * This topic provides an example on how to call the API operation to associate the vSwitch `vsw-bp183p93qs667muql****` and the member `172050525300****` with the resource share `rs-6GRmdD3X****` in the `cn-hangzhou` region. After the association, the vSwitch is shared with the member.
 *
 * @param request AssociateResourceShareRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return AssociateResourceShareResponse
 */
func (client *Client) AssociateResourceShareWithOptions(request *AssociateResourceShareRequest, runtime *util.RuntimeOptions) (_result *AssociateResourceShareResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PermissionNames)) {
		query["PermissionNames"] = request.PermissionNames
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

/**
 * This topic provides an example on how to call the API operation to associate the vSwitch `vsw-bp183p93qs667muql****` and the member `172050525300****` with the resource share `rs-6GRmdD3X****` in the `cn-hangzhou` region. After the association, the vSwitch is shared with the member.
 *
 * @param request AssociateResourceShareRequest
 * @return AssociateResourceShareResponse
 */
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

/**
 * This topic provides an example on how to call the API operation to associate the `AliyunRSDefaultPermissionVSwitch` permission with the `rs-6GRmdD3X****` resource share in the `cn-hangzhou` region.
 *
 * @param request AssociateResourceSharePermissionRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return AssociateResourceSharePermissionResponse
 */
func (client *Client) AssociateResourceSharePermissionWithOptions(request *AssociateResourceSharePermissionRequest, runtime *util.RuntimeOptions) (_result *AssociateResourceSharePermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PermissionName)) {
		query["PermissionName"] = request.PermissionName
	}

	if !tea.BoolValue(util.IsUnset(request.Replace)) {
		query["Replace"] = request.Replace
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceShareId)) {
		query["ResourceShareId"] = request.ResourceShareId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AssociateResourceSharePermission"),
		Version:     tea.String("2020-01-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AssociateResourceSharePermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to associate the `AliyunRSDefaultPermissionVSwitch` permission with the `rs-6GRmdD3X****` resource share in the `cn-hangzhou` region.
 *
 * @param request AssociateResourceSharePermissionRequest
 * @return AssociateResourceSharePermissionResponse
 */
func (client *Client) AssociateResourceSharePermission(request *AssociateResourceSharePermissionRequest) (_result *AssociateResourceSharePermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssociateResourceSharePermissionResponse{}
	_body, _err := client.AssociateResourceSharePermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

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

	if !tea.BoolValue(util.IsUnset(request.ResourceRegionId)) {
		query["ResourceRegionId"] = request.ResourceRegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeResourceGroup"),
		Version:     tea.String("2020-01-10"),
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

func (client *Client) CheckSharingWithResourceDirectoryStatusWithOptions(runtime *util.RuntimeOptions) (_result *CheckSharingWithResourceDirectoryStatusResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("CheckSharingWithResourceDirectoryStatus"),
		Version:     tea.String("2020-01-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckSharingWithResourceDirectoryStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckSharingWithResourceDirectoryStatus() (_result *CheckSharingWithResourceDirectoryStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckSharingWithResourceDirectoryStatusResponse{}
	_body, _err := client.CheckSharingWithResourceDirectoryStatusWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Resource Sharing allows you to share your resources with one or more accounts and access the resources shared by other accounts. For more information, see [Resource Sharing overview](~~160622~~).
 * This topic provides an example on how to call the API operation to create a resource share named `test` in the `cn-hangzhou` region to share the vSwitch `vsw-bp183p93qs667muql****` with the member `172050525300****`. In this example, the management account of a resource directory is used to call this API operation.
 *
 * @param request CreateResourceShareRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateResourceShareResponse
 */
func (client *Client) CreateResourceShareWithOptions(request *CreateResourceShareRequest, runtime *util.RuntimeOptions) (_result *CreateResourceShareResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllowExternalTargets)) {
		query["AllowExternalTargets"] = request.AllowExternalTargets
	}

	if !tea.BoolValue(util.IsUnset(request.PermissionNames)) {
		query["PermissionNames"] = request.PermissionNames
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

/**
 * Resource Sharing allows you to share your resources with one or more accounts and access the resources shared by other accounts. For more information, see [Resource Sharing overview](~~160622~~).
 * This topic provides an example on how to call the API operation to create a resource share named `test` in the `cn-hangzhou` region to share the vSwitch `vsw-bp183p93qs667muql****` with the member `172050525300****`. In this example, the management account of a resource directory is used to call this API operation.
 *
 * @param request CreateResourceShareRequest
 * @return CreateResourceShareResponse
 */
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

/**
 * After a resource share is deleted, all principals in the resource share can no longer access the resources in the resource share. However, the resources are not deleted with the resource share.
 * A resource share that is deleted is in the `Deleted` state. The system deletes the record of the resource share within 48 hours to 96 hours.
 * This topic provides an example on how to call the API operation to delete the resource share `rs-qSkW1HBY****` in the `cn-hangzhou` region.
 *
 * @param request DeleteResourceShareRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteResourceShareResponse
 */
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

/**
 * After a resource share is deleted, all principals in the resource share can no longer access the resources in the resource share. However, the resources are not deleted with the resource share.
 * A resource share that is deleted is in the `Deleted` state. The system deletes the record of the resource share within 48 hours to 96 hours.
 * This topic provides an example on how to call the API operation to delete the resource share `rs-qSkW1HBY****` in the `cn-hangzhou` region.
 *
 * @param request DeleteResourceShareRequest
 * @return DeleteResourceShareResponse
 */
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

/**
 * *   A resource owner can call this API operation to remove shared resources or principals from a resource share.
 * *   If an Alibaba Cloud account that is not the management account or a member of a resource directory is added to a resource share as a principal, you can use the Alibaba Cloud account to call this API operation to exit the resource share. For more information, see [Exit a resource share](~~440614~~).
 * This topic provides an example on how to use the management account of a resource directory to call the API operation to remove the member `172050525300****` from the resource share `rs-6GRmdD3X****` in the `cn-hangzhou` region. After the member is removed from the resource share, the member cannot share the resources in the resource share.
 *
 * @param request DisassociateResourceShareRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DisassociateResourceShareResponse
 */
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

/**
 * *   A resource owner can call this API operation to remove shared resources or principals from a resource share.
 * *   If an Alibaba Cloud account that is not the management account or a member of a resource directory is added to a resource share as a principal, you can use the Alibaba Cloud account to call this API operation to exit the resource share. For more information, see [Exit a resource share](~~440614~~).
 * This topic provides an example on how to use the management account of a resource directory to call the API operation to remove the member `172050525300****` from the resource share `rs-6GRmdD3X****` in the `cn-hangzhou` region. After the member is removed from the resource share, the member cannot share the resources in the resource share.
 *
 * @param request DisassociateResourceShareRequest
 * @return DisassociateResourceShareResponse
 */
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

/**
 * This topic provides an example on how to call the API operation to disassociate the `AliyunRSDefaultPermissionVSwitch` permission from the `rs-6GRmdD3X****` resource share in the `cn-hangzhou` region.
 *
 * @param request DisassociateResourceSharePermissionRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DisassociateResourceSharePermissionResponse
 */
func (client *Client) DisassociateResourceSharePermissionWithOptions(request *DisassociateResourceSharePermissionRequest, runtime *util.RuntimeOptions) (_result *DisassociateResourceSharePermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PermissionName)) {
		query["PermissionName"] = request.PermissionName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceShareId)) {
		query["ResourceShareId"] = request.ResourceShareId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisassociateResourceSharePermission"),
		Version:     tea.String("2020-01-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisassociateResourceSharePermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to disassociate the `AliyunRSDefaultPermissionVSwitch` permission from the `rs-6GRmdD3X****` resource share in the `cn-hangzhou` region.
 *
 * @param request DisassociateResourceSharePermissionRequest
 * @return DisassociateResourceSharePermissionResponse
 */
func (client *Client) DisassociateResourceSharePermission(request *DisassociateResourceSharePermissionRequest) (_result *DisassociateResourceSharePermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisassociateResourceSharePermissionResponse{}
	_body, _err := client.DisassociateResourceSharePermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can share your resources with all members in your resource directory, all members in a specific folder in your resource directory, or a specific member in your resource directory as a resource owner only after you enable resource sharing for your resource directory.
 * You can call this API operation only by using the management account of your resource directory or a RAM user or RAM role to which the required permissions are granted within the management account.
 *
 * @param request EnableSharingWithResourceDirectoryRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return EnableSharingWithResourceDirectoryResponse
 */
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

/**
 * You can share your resources with all members in your resource directory, all members in a specific folder in your resource directory, or a specific member in your resource directory as a resource owner only after you enable resource sharing for your resource directory.
 * You can call this API operation only by using the management account of your resource directory or a RAM user or RAM role to which the required permissions are granted within the management account.
 *
 * @return EnableSharingWithResourceDirectoryResponse
 */
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

/**
 * This topic provides an example on how to call the API operation to query the information about the `AliyunRSDefaultPermissionVSwitch` permission whose version is `v1` in the `cn-hangzhou` region.
 *
 * @param request GetPermissionRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetPermissionResponse
 */
func (client *Client) GetPermissionWithOptions(request *GetPermissionRequest, runtime *util.RuntimeOptions) (_result *GetPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PermissionName)) {
		query["PermissionName"] = request.PermissionName
	}

	if !tea.BoolValue(util.IsUnset(request.PermissionVersion)) {
		query["PermissionVersion"] = request.PermissionVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPermission"),
		Version:     tea.String("2020-01-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to query the information about the `AliyunRSDefaultPermissionVSwitch` permission whose version is `v1` in the `cn-hangzhou` region.
 *
 * @param request GetPermissionRequest
 * @return GetPermissionResponse
 */
func (client *Client) GetPermission(request *GetPermissionRequest) (_result *GetPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPermissionResponse{}
	_body, _err := client.GetPermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to query the versions of the `AliyunRSDefaultPermissionVSwitch` permission in the `cn-hangzhou` region.
 *
 * @param request ListPermissionVersionsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListPermissionVersionsResponse
 */
func (client *Client) ListPermissionVersionsWithOptions(request *ListPermissionVersionsRequest, runtime *util.RuntimeOptions) (_result *ListPermissionVersionsResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.PermissionName)) {
		query["PermissionName"] = request.PermissionName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPermissionVersions"),
		Version:     tea.String("2020-01-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPermissionVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to query the versions of the `AliyunRSDefaultPermissionVSwitch` permission in the `cn-hangzhou` region.
 *
 * @param request ListPermissionVersionsRequest
 * @return ListPermissionVersionsResponse
 */
func (client *Client) ListPermissionVersions(request *ListPermissionVersionsRequest) (_result *ListPermissionVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPermissionVersionsResponse{}
	_body, _err := client.ListPermissionVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to query the information about the default permission for the `VSwitch` resource type in the `cn-hangzhou` region.
 *
 * @param request ListPermissionsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListPermissionsResponse
 */
func (client *Client) ListPermissionsWithOptions(request *ListPermissionsRequest, runtime *util.RuntimeOptions) (_result *ListPermissionsResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPermissions"),
		Version:     tea.String("2020-01-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPermissionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to query the information about the default permission for the `VSwitch` resource type in the `cn-hangzhou` region.
 *
 * @param request ListPermissionsRequest
 * @return ListPermissionsResponse
 */
func (client *Client) ListPermissions(request *ListPermissionsRequest) (_result *ListPermissionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPermissionsResponse{}
	_body, _err := client.ListPermissionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to query the association records of the resource shares that are created by using the current Alibaba Cloud account in the `cn-hangzhou` region. The response shows the following records:
 * *   The resource `vsw-bp1upw03qyz8n7us9****` of the `VSwitch` type has been associated with the resource share `rs-6GRmdD3X****`. The resource is in the `Associated` state. This indicates that the resource is being shared.
 * *   The resource `vsw-bp183p93qs667muql****` of the `VSwitch` type has been disassociated from the resource share `rs-6GRmdD3X****`. The resource is in the `Disassociated` state. This indicates that the sharing of the resource is stopped.
 *
 * @param request ListResourceShareAssociationsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListResourceShareAssociationsResponse
 */
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

/**
 * This topic provides an example on how to call the API operation to query the association records of the resource shares that are created by using the current Alibaba Cloud account in the `cn-hangzhou` region. The response shows the following records:
 * *   The resource `vsw-bp1upw03qyz8n7us9****` of the `VSwitch` type has been associated with the resource share `rs-6GRmdD3X****`. The resource is in the `Associated` state. This indicates that the resource is being shared.
 * *   The resource `vsw-bp183p93qs667muql****` of the `VSwitch` type has been disassociated from the resource share `rs-6GRmdD3X****`. The resource is in the `Disassociated` state. This indicates that the sharing of the resource is stopped.
 *
 * @param request ListResourceShareAssociationsRequest
 * @return ListResourceShareAssociationsResponse
 */
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

/**
 * ### [](#)
 * This topic provides an example on how to call the API operation to query the resource sharing invitations that are received by the current account in the `cn-hangzhou` region. The response shows that one invitation is received by the current account and is waiting for confirmation.
 *
 * @param request ListResourceShareInvitationsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListResourceShareInvitationsResponse
 */
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

/**
 * ### [](#)
 * This topic provides an example on how to call the API operation to query the resource sharing invitations that are received by the current account in the `cn-hangzhou` region. The response shows that one invitation is received by the current account and is waiting for confirmation.
 *
 * @param request ListResourceShareInvitationsRequest
 * @return ListResourceShareInvitationsResponse
 */
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

/**
 * This topic provides an example on how to call the API operation to query the permissions that are associated with the resource share created by using the current Alibaba Cloud account in the `cn-hangzhou` region.
 *
 * @param request ListResourceSharePermissionsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListResourceSharePermissionsResponse
 */
func (client *Client) ListResourceSharePermissionsWithOptions(request *ListResourceSharePermissionsRequest, runtime *util.RuntimeOptions) (_result *ListResourceSharePermissionsResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceShareId)) {
		query["ResourceShareId"] = request.ResourceShareId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResourceSharePermissions"),
		Version:     tea.String("2020-01-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResourceSharePermissionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to query the permissions that are associated with the resource share created by using the current Alibaba Cloud account in the `cn-hangzhou` region.
 *
 * @param request ListResourceSharePermissionsRequest
 * @return ListResourceSharePermissionsResponse
 */
func (client *Client) ListResourceSharePermissions(request *ListResourceSharePermissionsRequest) (_result *ListResourceSharePermissionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListResourceSharePermissionsResponse{}
	_body, _err := client.ListResourceSharePermissionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to query the resource shares that are created by using the current Alibaba Cloud account in the `cn-hangzhou` region. The response shows that the following resource shares are created by using the account whose ID is `151266687691****`:
 * *   `rs-hX9wC5jO****`, which is in the `Deleted` state
 * *   `rs-PqysnzIj****`, which is in the `Active` state
 *
 * @param request ListResourceSharesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListResourceSharesResponse
 */
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

	if !tea.BoolValue(util.IsUnset(request.PermissionName)) {
		query["PermissionName"] = request.PermissionName
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

/**
 * This topic provides an example on how to call the API operation to query the resource shares that are created by using the current Alibaba Cloud account in the `cn-hangzhou` region. The response shows that the following resource shares are created by using the account whose ID is `151266687691****`:
 * *   `rs-hX9wC5jO****`, which is in the `Deleted` state
 * *   `rs-PqysnzIj****`, which is in the `Active` state
 *
 * @param request ListResourceSharesRequest
 * @return ListResourceSharesResponse
 */
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

/**
 * This topic provides an example on how to call the API operation to query the resources that you share with other accounts in the `cn-hangzhou` region. The response shows that in the resource share `rs-6GRmdD3X****`, you share the `vsw-bp1upw03qyz8n7us9****` resource of the `VSwitch` type with other accounts.
 *
 * @param request ListSharedResourcesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListSharedResourcesResponse
 */
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

/**
 * This topic provides an example on how to call the API operation to query the resources that you share with other accounts in the `cn-hangzhou` region. The response shows that in the resource share `rs-6GRmdD3X****`, you share the `vsw-bp1upw03qyz8n7us9****` resource of the `VSwitch` type with other accounts.
 *
 * @param request ListSharedResourcesRequest
 * @return ListSharedResourcesResponse
 */
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

/**
 * If you are a resource owner, you can query the principals with which you share your resources.
 * If you are a principal, you can query the resources that are shared with you.
 * This topic provides an example on how to call the API operation to query the principals with which you share your resources in the `cn-hangzhou` region. The response shows that you share your resources with the principals `114240524784****` and `172050525300****`.
 *
 * @param request ListSharedTargetsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListSharedTargetsResponse
 */
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

/**
 * If you are a resource owner, you can query the principals with which you share your resources.
 * If you are a principal, you can query the resources that are shared with you.
 * This topic provides an example on how to call the API operation to query the principals with which you share your resources in the `cn-hangzhou` region. The response shows that you share your resources with the principals `114240524784****` and `172050525300****`.
 *
 * @param request ListSharedTargetsRequest
 * @return ListSharedTargetsResponse
 */
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

/**
 * This topic provides an example on how to call the API operation to reject the resource sharing invitation `i-yyTWbkjHArYh****` in the `cn-hangzhou` region.
 *
 * @param request RejectResourceShareInvitationRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return RejectResourceShareInvitationResponse
 */
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

/**
 * This topic provides an example on how to call the API operation to reject the resource sharing invitation `i-yyTWbkjHArYh****` in the `cn-hangzhou` region.
 *
 * @param request RejectResourceShareInvitationRequest
 * @return RejectResourceShareInvitationResponse
 */
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

/**
 * You can call this API operation to change the name or resource sharing scope of a resource share.
 * This topic provides an example on how to call the API operation to change the name of the resource share `rs-qSkW1HBY****` in the `cn-hangzhou` region from `test` to `new`.
 *
 * @param request UpdateResourceShareRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateResourceShareResponse
 */
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

/**
 * You can call this API operation to change the name or resource sharing scope of a resource share.
 * This topic provides an example on how to call the API operation to change the name of the resource share `rs-qSkW1HBY****` in the `cn-hangzhou` region from `test` to `new`.
 *
 * @param request UpdateResourceShareRequest
 * @return UpdateResourceShareResponse
 */
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
