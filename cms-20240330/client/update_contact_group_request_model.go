// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContactGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContacts(v []*string) *UpdateContactGroupRequest
	GetContacts() []*string
	SetName(v string) *UpdateContactGroupRequest
	GetName() *string
}

type UpdateContactGroupRequest struct {
	// The ID of the contact.
	Contacts []*string `json:"contacts,omitempty" xml:"contacts,omitempty" type:"Repeated"`
	// The name of the contact.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s UpdateContactGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateContactGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateContactGroupRequest) GetContacts() []*string {
	return s.Contacts
}

func (s *UpdateContactGroupRequest) GetName() *string {
	return s.Name
}

func (s *UpdateContactGroupRequest) SetContacts(v []*string) *UpdateContactGroupRequest {
	s.Contacts = v
	return s
}

func (s *UpdateContactGroupRequest) SetName(v string) *UpdateContactGroupRequest {
	s.Name = &v
	return s
}

func (s *UpdateContactGroupRequest) Validate() error {
	return dara.Validate(s)
}
