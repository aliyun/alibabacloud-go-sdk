// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMessageContactDeletionStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactId(v string) *GetMessageContactDeletionStatusRequest
	GetContactId() *string
}

type GetMessageContactDeletionStatusRequest struct {
	// The ID of the contact.
	//
	// example:
	//
	// c-qL4HqKONzOM7****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
}

func (s GetMessageContactDeletionStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMessageContactDeletionStatusRequest) GoString() string {
	return s.String()
}

func (s *GetMessageContactDeletionStatusRequest) GetContactId() *string {
	return s.ContactId
}

func (s *GetMessageContactDeletionStatusRequest) SetContactId(v string) *GetMessageContactDeletionStatusRequest {
	s.ContactId = &v
	return s
}

func (s *GetMessageContactDeletionStatusRequest) Validate() error {
	return dara.Validate(s)
}
