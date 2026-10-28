// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContactGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactGroupId(v string) *CreateContactGroupRequest
	GetContactGroupId() *string
	SetContacts(v []*string) *CreateContactGroupRequest
	GetContacts() []*string
	SetName(v string) *CreateContactGroupRequest
	GetName() *string
	SetWorkspace(v string) *CreateContactGroupRequest
	GetWorkspace() *string
}

type CreateContactGroupRequest struct {
	// example:
	//
	// test
	ContactGroupId *string   `json:"contactGroupId,omitempty" xml:"contactGroupId,omitempty"`
	Contacts       []*string `json:"contacts,omitempty" xml:"contacts,omitempty" type:"Repeated"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// default-cms-5884500827389836-cn-hangzhou
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s CreateContactGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateContactGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateContactGroupRequest) GetContactGroupId() *string {
	return s.ContactGroupId
}

func (s *CreateContactGroupRequest) GetContacts() []*string {
	return s.Contacts
}

func (s *CreateContactGroupRequest) GetName() *string {
	return s.Name
}

func (s *CreateContactGroupRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *CreateContactGroupRequest) SetContactGroupId(v string) *CreateContactGroupRequest {
	s.ContactGroupId = &v
	return s
}

func (s *CreateContactGroupRequest) SetContacts(v []*string) *CreateContactGroupRequest {
	s.Contacts = v
	return s
}

func (s *CreateContactGroupRequest) SetName(v string) *CreateContactGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateContactGroupRequest) SetWorkspace(v string) *CreateContactGroupRequest {
	s.Workspace = &v
	return s
}

func (s *CreateContactGroupRequest) Validate() error {
	return dara.Validate(s)
}
