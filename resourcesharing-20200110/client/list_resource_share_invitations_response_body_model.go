// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceShareInvitationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListResourceShareInvitationsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListResourceShareInvitationsResponseBody
	GetRequestId() *string
	SetResourceShareInvitations(v []*ListResourceShareInvitationsResponseBodyResourceShareInvitations) *ListResourceShareInvitationsResponseBody
	GetResourceShareInvitations() []*ListResourceShareInvitationsResponseBodyResourceShareInvitations
}

type ListResourceShareInvitationsResponseBody struct {
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of `NextToken`.
	//
	// example:
	//
	// TGlzdFJlc291cm****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 30EC8328-1BDE-51D5-BFAB-039508BD91A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the resource sharing invitations.
	ResourceShareInvitations []*ListResourceShareInvitationsResponseBodyResourceShareInvitations `json:"ResourceShareInvitations,omitempty" xml:"ResourceShareInvitations,omitempty" type:"Repeated"`
}

func (s ListResourceShareInvitationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourceShareInvitationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceShareInvitationsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListResourceShareInvitationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourceShareInvitationsResponseBody) GetResourceShareInvitations() []*ListResourceShareInvitationsResponseBodyResourceShareInvitations {
	return s.ResourceShareInvitations
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

func (s *ListResourceShareInvitationsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListResourceShareInvitationsResponseBodyResourceShareInvitations struct {
	// The time when the invitation was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-08-18T05:36:45.024Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The information about the failure.
	InvitationFailedDetails []*ListResourceShareInvitationsResponseBodyResourceShareInvitationsInvitationFailedDetails `json:"InvitationFailedDetails,omitempty" xml:"InvitationFailedDetails,omitempty" type:"Repeated"`
	// The Alibaba Cloud account ID of the invitee.
	//
	// example:
	//
	// 134254031178****
	ReceiverAccountId *string `json:"ReceiverAccountId,omitempty" xml:"ReceiverAccountId,omitempty"`
	// The ID of the resource share.
	//
	// example:
	//
	// rs-ysGRci9z****
	ResourceShareId *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
	// The ID of the invitation.
	//
	// example:
	//
	// i-p6eRytrkjVvM****
	ResourceShareInvitationId *string `json:"ResourceShareInvitationId,omitempty" xml:"ResourceShareInvitationId,omitempty"`
	// The name of the resource share.
	//
	// example:
	//
	// example
	ResourceShareName *string `json:"ResourceShareName,omitempty" xml:"ResourceShareName,omitempty"`
	// The Alibaba Cloud account ID of the inviter.
	//
	// example:
	//
	// 151266687691****
	SenderAccountId *string `json:"SenderAccountId,omitempty" xml:"SenderAccountId,omitempty"`
	// The status of the invitation. Valid values:
	//
	// 	- Pending
	//
	// 	- Accepted
	//
	// 	- Cancelled
	//
	// 	- Rejected
	//
	// 	- Expired
	//
	// 	- AcceptFailed
	//
	// example:
	//
	// Pending
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListResourceShareInvitationsResponseBodyResourceShareInvitations) String() string {
	return dara.Prettify(s)
}

func (s ListResourceShareInvitationsResponseBodyResourceShareInvitations) GoString() string {
	return s.String()
}

func (s *ListResourceShareInvitationsResponseBodyResourceShareInvitations) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListResourceShareInvitationsResponseBodyResourceShareInvitations) GetInvitationFailedDetails() []*ListResourceShareInvitationsResponseBodyResourceShareInvitationsInvitationFailedDetails {
	return s.InvitationFailedDetails
}

func (s *ListResourceShareInvitationsResponseBodyResourceShareInvitations) GetReceiverAccountId() *string {
	return s.ReceiverAccountId
}

func (s *ListResourceShareInvitationsResponseBodyResourceShareInvitations) GetResourceShareId() *string {
	return s.ResourceShareId
}

func (s *ListResourceShareInvitationsResponseBodyResourceShareInvitations) GetResourceShareInvitationId() *string {
	return s.ResourceShareInvitationId
}

func (s *ListResourceShareInvitationsResponseBodyResourceShareInvitations) GetResourceShareName() *string {
	return s.ResourceShareName
}

func (s *ListResourceShareInvitationsResponseBodyResourceShareInvitations) GetSenderAccountId() *string {
	return s.SenderAccountId
}

func (s *ListResourceShareInvitationsResponseBodyResourceShareInvitations) GetStatus() *string {
	return s.Status
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

func (s *ListResourceShareInvitationsResponseBodyResourceShareInvitations) Validate() error {
	return dara.Validate(s)
}

type ListResourceShareInvitationsResponseBodyResourceShareInvitationsInvitationFailedDetails struct {
	// This parameter is deprecated. The OperationType parameter is used instead.
	//
	// example:
	//
	// Associate
	AssociateType *string `json:"AssociateType,omitempty" xml:"AssociateType,omitempty"`
	// The failure description.
	//
	// example:
	//
	// You cannot access the specified resource at this time.
	FailureDescription *string `json:"FailureDescription,omitempty" xml:"FailureDescription,omitempty"`
	// The failure cause. Valid values:
	//
	// 	- Unavailable: The resource cannot be shared.
	//
	// 	- LimitExceeded: The number of shared resources within the Alibaba Cloud account exceeds the upper limit.
	//
	// 	- ZonalResourceInaccessible: The resource is unavailable in this region.
	//
	// 	- InternalError: An internal error occurred during the check.
	//
	// 	- UnsupportedOperation: You cannot perform this operation.
	//
	// example:
	//
	// Unavailable
	FailureReason *string `json:"FailureReason,omitempty" xml:"FailureReason,omitempty"`
	// The operation type. Valid values:
	//
	// 	- Associate
	//
	// 	- Disassociate
	//
	// example:
	//
	// Associate
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	ResourceArn   *string `json:"ResourceArn,omitempty" xml:"ResourceArn,omitempty"`
	// The ID of the shared resource.
	//
	// example:
	//
	// s-7xvh46nx5oqlre0wv***
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the shared resource.
	//
	// For more information about the types of resources that can be shared, see [Services that work with Resource Sharing](https://help.aliyun.com/document_detail/450526.html).
	//
	// example:
	//
	// Snapshot
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// This parameter is deprecated. The FailureReason parameter is used instead.
	//
	// example:
	//
	// Unavailable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// This parameter is deprecated. The FailureDescription parameter is used instead.
	//
	// example:
	//
	// You cannot access the specified resource at this time.
	StatusMessage *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
}

func (s ListResourceShareInvitationsResponseBodyResourceShareInvitationsInvitationFailedDetails) String() string {
	return dara.Prettify(s)
}

func (s ListResourceShareInvitationsResponseBodyResourceShareInvitationsInvitationFailedDetails) GoString() string {
	return s.String()
}

func (s *ListResourceShareInvitationsResponseBodyResourceShareInvitationsInvitationFailedDetails) GetAssociateType() *string {
	return s.AssociateType
}

func (s *ListResourceShareInvitationsResponseBodyResourceShareInvitationsInvitationFailedDetails) GetFailureDescription() *string {
	return s.FailureDescription
}

func (s *ListResourceShareInvitationsResponseBodyResourceShareInvitationsInvitationFailedDetails) GetFailureReason() *string {
	return s.FailureReason
}

func (s *ListResourceShareInvitationsResponseBodyResourceShareInvitationsInvitationFailedDetails) GetOperationType() *string {
	return s.OperationType
}

func (s *ListResourceShareInvitationsResponseBodyResourceShareInvitationsInvitationFailedDetails) GetResourceArn() *string {
	return s.ResourceArn
}

func (s *ListResourceShareInvitationsResponseBodyResourceShareInvitationsInvitationFailedDetails) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListResourceShareInvitationsResponseBodyResourceShareInvitationsInvitationFailedDetails) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListResourceShareInvitationsResponseBodyResourceShareInvitationsInvitationFailedDetails) GetStatus() *string {
	return s.Status
}

func (s *ListResourceShareInvitationsResponseBodyResourceShareInvitationsInvitationFailedDetails) GetStatusMessage() *string {
	return s.StatusMessage
}

func (s *ListResourceShareInvitationsResponseBodyResourceShareInvitationsInvitationFailedDetails) SetAssociateType(v string) *ListResourceShareInvitationsResponseBodyResourceShareInvitationsInvitationFailedDetails {
	s.AssociateType = &v
	return s
}

func (s *ListResourceShareInvitationsResponseBodyResourceShareInvitationsInvitationFailedDetails) SetFailureDescription(v string) *ListResourceShareInvitationsResponseBodyResourceShareInvitationsInvitationFailedDetails {
	s.FailureDescription = &v
	return s
}

func (s *ListResourceShareInvitationsResponseBodyResourceShareInvitationsInvitationFailedDetails) SetFailureReason(v string) *ListResourceShareInvitationsResponseBodyResourceShareInvitationsInvitationFailedDetails {
	s.FailureReason = &v
	return s
}

func (s *ListResourceShareInvitationsResponseBodyResourceShareInvitationsInvitationFailedDetails) SetOperationType(v string) *ListResourceShareInvitationsResponseBodyResourceShareInvitationsInvitationFailedDetails {
	s.OperationType = &v
	return s
}

func (s *ListResourceShareInvitationsResponseBodyResourceShareInvitationsInvitationFailedDetails) SetResourceArn(v string) *ListResourceShareInvitationsResponseBodyResourceShareInvitationsInvitationFailedDetails {
	s.ResourceArn = &v
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

func (s *ListResourceShareInvitationsResponseBodyResourceShareInvitationsInvitationFailedDetails) Validate() error {
	return dara.Validate(s)
}
