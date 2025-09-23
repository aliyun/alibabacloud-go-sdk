// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRejectResourceShareInvitationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceShareInvitationId(v string) *RejectResourceShareInvitationRequest
	GetResourceShareInvitationId() *string
}

type RejectResourceShareInvitationRequest struct {
	// The ID of the resource sharing invitation.
	//
	// You can call the [ListResourceShareInvitations](https://help.aliyun.com/document_detail/450564.html) operation to obtain the ID of a resource sharing invitation.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-yyTWbkjHArYh****
	ResourceShareInvitationId *string `json:"ResourceShareInvitationId,omitempty" xml:"ResourceShareInvitationId,omitempty"`
}

func (s RejectResourceShareInvitationRequest) String() string {
	return dara.Prettify(s)
}

func (s RejectResourceShareInvitationRequest) GoString() string {
	return s.String()
}

func (s *RejectResourceShareInvitationRequest) GetResourceShareInvitationId() *string {
	return s.ResourceShareInvitationId
}

func (s *RejectResourceShareInvitationRequest) SetResourceShareInvitationId(v string) *RejectResourceShareInvitationRequest {
	s.ResourceShareInvitationId = &v
	return s
}

func (s *RejectResourceShareInvitationRequest) Validate() error {
	return dara.Validate(s)
}
