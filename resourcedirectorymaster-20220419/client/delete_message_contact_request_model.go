// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMessageContactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactId(v string) *DeleteMessageContactRequest
	GetContactId() *string
	SetRetainContactInMembers(v bool) *DeleteMessageContactRequest
	GetRetainContactInMembers() *bool
}

type DeleteMessageContactRequest struct {
	// The ID of the contact.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-qL4HqKONzOM7****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// Specifies whether to retain the contact for members. Valid values:
	//
	// 	- true (default): retains the contact for members. In this case, the contact can still receive messages for the members.
	//
	// 	- false: does not retain the contact for members. In this case, the contact can no longer receive messages for the members. If you set this parameter to false, the response is asynchronously returned. You can call [GetMessageContactDeletionStatus](~~GetMessageContactDeletionStatus~~) to obtain the deletion result.
	//
	// example:
	//
	// true
	RetainContactInMembers *bool `json:"RetainContactInMembers,omitempty" xml:"RetainContactInMembers,omitempty"`
}

func (s DeleteMessageContactRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMessageContactRequest) GoString() string {
	return s.String()
}

func (s *DeleteMessageContactRequest) GetContactId() *string {
	return s.ContactId
}

func (s *DeleteMessageContactRequest) GetRetainContactInMembers() *bool {
	return s.RetainContactInMembers
}

func (s *DeleteMessageContactRequest) SetContactId(v string) *DeleteMessageContactRequest {
	s.ContactId = &v
	return s
}

func (s *DeleteMessageContactRequest) SetRetainContactInMembers(v bool) *DeleteMessageContactRequest {
	s.RetainContactInMembers = &v
	return s
}

func (s *DeleteMessageContactRequest) Validate() error {
	return dara.Validate(s)
}
