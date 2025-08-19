// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMessageContactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactId(v string) *GetMessageContactRequest
	GetContactId() *string
}

type GetMessageContactRequest struct {
	// The ID of the contact.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-qL4HqKONzOM7****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
}

func (s GetMessageContactRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMessageContactRequest) GoString() string {
	return s.String()
}

func (s *GetMessageContactRequest) GetContactId() *string {
	return s.ContactId
}

func (s *GetMessageContactRequest) SetContactId(v string) *GetMessageContactRequest {
	s.ContactId = &v
	return s
}

func (s *GetMessageContactRequest) Validate() error {
	return dara.Validate(s)
}
