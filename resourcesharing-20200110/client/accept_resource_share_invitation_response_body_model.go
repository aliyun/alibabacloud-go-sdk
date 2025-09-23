// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAcceptResourceShareInvitationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AcceptResourceShareInvitationResponseBody
	GetRequestId() *string
	SetResourceShareInvitation(v *AcceptResourceShareInvitationResponseBodyResourceShareInvitation) *AcceptResourceShareInvitationResponseBody
	GetResourceShareInvitation() *AcceptResourceShareInvitationResponseBodyResourceShareInvitation
}

type AcceptResourceShareInvitationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 08F18B04-47CB-5C0E-A6D2-37DEF5C2A961
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the resource sharing invitation.
	ResourceShareInvitation *AcceptResourceShareInvitationResponseBodyResourceShareInvitation `json:"ResourceShareInvitation,omitempty" xml:"ResourceShareInvitation,omitempty" type:"Struct"`
}

func (s AcceptResourceShareInvitationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AcceptResourceShareInvitationResponseBody) GoString() string {
	return s.String()
}

func (s *AcceptResourceShareInvitationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AcceptResourceShareInvitationResponseBody) GetResourceShareInvitation() *AcceptResourceShareInvitationResponseBodyResourceShareInvitation {
	return s.ResourceShareInvitation
}

func (s *AcceptResourceShareInvitationResponseBody) SetRequestId(v string) *AcceptResourceShareInvitationResponseBody {
	s.RequestId = &v
	return s
}

func (s *AcceptResourceShareInvitationResponseBody) SetResourceShareInvitation(v *AcceptResourceShareInvitationResponseBodyResourceShareInvitation) *AcceptResourceShareInvitationResponseBody {
	s.ResourceShareInvitation = v
	return s
}

func (s *AcceptResourceShareInvitationResponseBody) Validate() error {
	return dara.Validate(s)
}

type AcceptResourceShareInvitationResponseBodyResourceShareInvitation struct {
	// The information about the failure.
	AcceptInvitationFailedDetails []*AcceptResourceShareInvitationResponseBodyResourceShareInvitationAcceptInvitationFailedDetails `json:"AcceptInvitationFailedDetails,omitempty" xml:"AcceptInvitationFailedDetails,omitempty" type:"Repeated"`
	// The time when the invitation was created. The time is displayed in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-09-02T06:43:12.353Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The Alibaba Cloud account ID of the invitee.
	//
	// This parameter is required.
	//
	// example:
	//
	// 134254031178****
	ReceiverAccountId *string `json:"ReceiverAccountId,omitempty" xml:"ReceiverAccountId,omitempty"`
	// The ID of the resource share.
	//
	// This parameter is required.
	//
	// example:
	//
	// rs-ysGRci9z****
	ResourceShareId *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
	// The ID of the resource sharing invitation.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-pMnItMX19fBJ****
	ResourceShareInvitationId *string `json:"ResourceShareInvitationId,omitempty" xml:"ResourceShareInvitationId,omitempty"`
	// The name of the resource share.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	ResourceShareName *string `json:"ResourceShareName,omitempty" xml:"ResourceShareName,omitempty"`
	// The Alibaba Cloud account ID of the inviter.
	//
	// This parameter is required.
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
	// This parameter is required.
	//
	// example:
	//
	// AcceptFailed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s AcceptResourceShareInvitationResponseBodyResourceShareInvitation) String() string {
	return dara.Prettify(s)
}

func (s AcceptResourceShareInvitationResponseBodyResourceShareInvitation) GoString() string {
	return s.String()
}

func (s *AcceptResourceShareInvitationResponseBodyResourceShareInvitation) GetAcceptInvitationFailedDetails() []*AcceptResourceShareInvitationResponseBodyResourceShareInvitationAcceptInvitationFailedDetails {
	return s.AcceptInvitationFailedDetails
}

func (s *AcceptResourceShareInvitationResponseBodyResourceShareInvitation) GetCreateTime() *string {
	return s.CreateTime
}

func (s *AcceptResourceShareInvitationResponseBodyResourceShareInvitation) GetReceiverAccountId() *string {
	return s.ReceiverAccountId
}

func (s *AcceptResourceShareInvitationResponseBodyResourceShareInvitation) GetResourceShareId() *string {
	return s.ResourceShareId
}

func (s *AcceptResourceShareInvitationResponseBodyResourceShareInvitation) GetResourceShareInvitationId() *string {
	return s.ResourceShareInvitationId
}

func (s *AcceptResourceShareInvitationResponseBodyResourceShareInvitation) GetResourceShareName() *string {
	return s.ResourceShareName
}

func (s *AcceptResourceShareInvitationResponseBodyResourceShareInvitation) GetSenderAccountId() *string {
	return s.SenderAccountId
}

func (s *AcceptResourceShareInvitationResponseBodyResourceShareInvitation) GetStatus() *string {
	return s.Status
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

func (s *AcceptResourceShareInvitationResponseBodyResourceShareInvitation) Validate() error {
	return dara.Validate(s)
}

type AcceptResourceShareInvitationResponseBodyResourceShareInvitationAcceptInvitationFailedDetails struct {
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
	// example:
	//
	// Unavailable
	FailureReason *string `json:"FailureReason,omitempty" xml:"FailureReason,omitempty"`
	// The operation type. Valid values:
	//
	// 	- Associate
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

func (s AcceptResourceShareInvitationResponseBodyResourceShareInvitationAcceptInvitationFailedDetails) String() string {
	return dara.Prettify(s)
}

func (s AcceptResourceShareInvitationResponseBodyResourceShareInvitationAcceptInvitationFailedDetails) GoString() string {
	return s.String()
}

func (s *AcceptResourceShareInvitationResponseBodyResourceShareInvitationAcceptInvitationFailedDetails) GetAssociateType() *string {
	return s.AssociateType
}

func (s *AcceptResourceShareInvitationResponseBodyResourceShareInvitationAcceptInvitationFailedDetails) GetFailureDescription() *string {
	return s.FailureDescription
}

func (s *AcceptResourceShareInvitationResponseBodyResourceShareInvitationAcceptInvitationFailedDetails) GetFailureReason() *string {
	return s.FailureReason
}

func (s *AcceptResourceShareInvitationResponseBodyResourceShareInvitationAcceptInvitationFailedDetails) GetOperationType() *string {
	return s.OperationType
}

func (s *AcceptResourceShareInvitationResponseBodyResourceShareInvitationAcceptInvitationFailedDetails) GetResourceArn() *string {
	return s.ResourceArn
}

func (s *AcceptResourceShareInvitationResponseBodyResourceShareInvitationAcceptInvitationFailedDetails) GetResourceId() *string {
	return s.ResourceId
}

func (s *AcceptResourceShareInvitationResponseBodyResourceShareInvitationAcceptInvitationFailedDetails) GetResourceType() *string {
	return s.ResourceType
}

func (s *AcceptResourceShareInvitationResponseBodyResourceShareInvitationAcceptInvitationFailedDetails) GetStatus() *string {
	return s.Status
}

func (s *AcceptResourceShareInvitationResponseBodyResourceShareInvitationAcceptInvitationFailedDetails) GetStatusMessage() *string {
	return s.StatusMessage
}

func (s *AcceptResourceShareInvitationResponseBodyResourceShareInvitationAcceptInvitationFailedDetails) SetAssociateType(v string) *AcceptResourceShareInvitationResponseBodyResourceShareInvitationAcceptInvitationFailedDetails {
	s.AssociateType = &v
	return s
}

func (s *AcceptResourceShareInvitationResponseBodyResourceShareInvitationAcceptInvitationFailedDetails) SetFailureDescription(v string) *AcceptResourceShareInvitationResponseBodyResourceShareInvitationAcceptInvitationFailedDetails {
	s.FailureDescription = &v
	return s
}

func (s *AcceptResourceShareInvitationResponseBodyResourceShareInvitationAcceptInvitationFailedDetails) SetFailureReason(v string) *AcceptResourceShareInvitationResponseBodyResourceShareInvitationAcceptInvitationFailedDetails {
	s.FailureReason = &v
	return s
}

func (s *AcceptResourceShareInvitationResponseBodyResourceShareInvitationAcceptInvitationFailedDetails) SetOperationType(v string) *AcceptResourceShareInvitationResponseBodyResourceShareInvitationAcceptInvitationFailedDetails {
	s.OperationType = &v
	return s
}

func (s *AcceptResourceShareInvitationResponseBodyResourceShareInvitationAcceptInvitationFailedDetails) SetResourceArn(v string) *AcceptResourceShareInvitationResponseBodyResourceShareInvitationAcceptInvitationFailedDetails {
	s.ResourceArn = &v
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

func (s *AcceptResourceShareInvitationResponseBodyResourceShareInvitationAcceptInvitationFailedDetails) Validate() error {
	return dara.Validate(s)
}
