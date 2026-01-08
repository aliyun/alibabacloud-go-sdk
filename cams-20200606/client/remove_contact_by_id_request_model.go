// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveContactByIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactId(v string) *RemoveContactByIdRequest
	GetContactId() *string
	SetGroupId(v string) *RemoveContactByIdRequest
	GetGroupId() *string
	SetOwnerId(v int64) *RemoveContactByIdRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *RemoveContactByIdRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RemoveContactByIdRequest
	GetResourceOwnerId() *int64
}

type RemoveContactByIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 3456456346**
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 46546546546**
	GroupId              *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RemoveContactByIdRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveContactByIdRequest) GoString() string {
	return s.String()
}

func (s *RemoveContactByIdRequest) GetContactId() *string {
	return s.ContactId
}

func (s *RemoveContactByIdRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *RemoveContactByIdRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RemoveContactByIdRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RemoveContactByIdRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RemoveContactByIdRequest) SetContactId(v string) *RemoveContactByIdRequest {
	s.ContactId = &v
	return s
}

func (s *RemoveContactByIdRequest) SetGroupId(v string) *RemoveContactByIdRequest {
	s.GroupId = &v
	return s
}

func (s *RemoveContactByIdRequest) SetOwnerId(v int64) *RemoveContactByIdRequest {
	s.OwnerId = &v
	return s
}

func (s *RemoveContactByIdRequest) SetResourceOwnerAccount(v string) *RemoveContactByIdRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RemoveContactByIdRequest) SetResourceOwnerId(v int64) *RemoveContactByIdRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RemoveContactByIdRequest) Validate() error {
	return dara.Validate(s)
}
