// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRejectResourceShareInvitationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RejectResourceShareInvitationResponseBody
	GetRequestId() *string
	SetResourceShareInvitation(v *RejectResourceShareInvitationResponseBodyResourceShareInvitation) *RejectResourceShareInvitationResponseBody
	GetResourceShareInvitation() *RejectResourceShareInvitationResponseBodyResourceShareInvitation
}

type RejectResourceShareInvitationResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// E446D6DE-BFC8-5F37-A494-33D7B118147D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the resource sharing invitation.
	ResourceShareInvitation *RejectResourceShareInvitationResponseBodyResourceShareInvitation `json:"ResourceShareInvitation,omitempty" xml:"ResourceShareInvitation,omitempty" type:"Struct"`
}

func (s RejectResourceShareInvitationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RejectResourceShareInvitationResponseBody) GoString() string {
	return s.String()
}

func (s *RejectResourceShareInvitationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RejectResourceShareInvitationResponseBody) GetResourceShareInvitation() *RejectResourceShareInvitationResponseBodyResourceShareInvitation {
	return s.ResourceShareInvitation
}

func (s *RejectResourceShareInvitationResponseBody) SetRequestId(v string) *RejectResourceShareInvitationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RejectResourceShareInvitationResponseBody) SetResourceShareInvitation(v *RejectResourceShareInvitationResponseBodyResourceShareInvitation) *RejectResourceShareInvitationResponseBody {
	s.ResourceShareInvitation = v
	return s
}

func (s *RejectResourceShareInvitationResponseBody) Validate() error {
	if s.ResourceShareInvitation != nil {
		if err := s.ResourceShareInvitation.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RejectResourceShareInvitationResponseBodyResourceShareInvitation struct {
	// The time when the invitation was created. The time is displayed in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-09-02T07:07:30.809Z
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
	// rs-JoA1Ayjm****
	ResourceShareId *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
	// The ID of the invitation.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-yyTWbkjHArYh****
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
	// 	- Pending: The invitation is waiting for confirmation.
	//
	// 	- Accepted: The invitation is accepted.
	//
	// 	- Cancelled: The invitation is canceled.
	//
	// 	- Rejected: The invitation is rejected.
	//
	// 	- Expired: The invitation has expired.
	//
	// This parameter is required.
	//
	// example:
	//
	// Rejected
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s RejectResourceShareInvitationResponseBodyResourceShareInvitation) String() string {
	return dara.Prettify(s)
}

func (s RejectResourceShareInvitationResponseBodyResourceShareInvitation) GoString() string {
	return s.String()
}

func (s *RejectResourceShareInvitationResponseBodyResourceShareInvitation) GetCreateTime() *string {
	return s.CreateTime
}

func (s *RejectResourceShareInvitationResponseBodyResourceShareInvitation) GetReceiverAccountId() *string {
	return s.ReceiverAccountId
}

func (s *RejectResourceShareInvitationResponseBodyResourceShareInvitation) GetResourceShareId() *string {
	return s.ResourceShareId
}

func (s *RejectResourceShareInvitationResponseBodyResourceShareInvitation) GetResourceShareInvitationId() *string {
	return s.ResourceShareInvitationId
}

func (s *RejectResourceShareInvitationResponseBodyResourceShareInvitation) GetResourceShareName() *string {
	return s.ResourceShareName
}

func (s *RejectResourceShareInvitationResponseBodyResourceShareInvitation) GetSenderAccountId() *string {
	return s.SenderAccountId
}

func (s *RejectResourceShareInvitationResponseBodyResourceShareInvitation) GetStatus() *string {
	return s.Status
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

func (s *RejectResourceShareInvitationResponseBodyResourceShareInvitation) Validate() error {
	return dara.Validate(s)
}
