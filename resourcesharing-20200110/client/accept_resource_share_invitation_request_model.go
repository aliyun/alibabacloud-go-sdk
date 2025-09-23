// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAcceptResourceShareInvitationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceShareInvitationId(v string) *AcceptResourceShareInvitationRequest
	GetResourceShareInvitationId() *string
}

type AcceptResourceShareInvitationRequest struct {
	// The ID of the resource sharing invitation.
	//
	// You can call the [ListResourceShareInvitations](https://help.aliyun.com/document_detail/450564.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-pMnItMX19fBJ****
	ResourceShareInvitationId *string `json:"ResourceShareInvitationId,omitempty" xml:"ResourceShareInvitationId,omitempty"`
}

func (s AcceptResourceShareInvitationRequest) String() string {
	return dara.Prettify(s)
}

func (s AcceptResourceShareInvitationRequest) GoString() string {
	return s.String()
}

func (s *AcceptResourceShareInvitationRequest) GetResourceShareInvitationId() *string {
	return s.ResourceShareInvitationId
}

func (s *AcceptResourceShareInvitationRequest) SetResourceShareInvitationId(v string) *AcceptResourceShareInvitationRequest {
	s.ResourceShareInvitationId = &v
	return s
}

func (s *AcceptResourceShareInvitationRequest) Validate() error {
	return dara.Validate(s)
}
