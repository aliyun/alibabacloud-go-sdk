// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactName(v string) *DeleteContactRequest
	GetContactName() *string
}

type DeleteContactRequest struct {
	// The name of the alert contact.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-01
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
}

func (s DeleteContactRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteContactRequest) GoString() string {
	return s.String()
}

func (s *DeleteContactRequest) GetContactName() *string {
	return s.ContactName
}

func (s *DeleteContactRequest) SetContactName(v string) *DeleteContactRequest {
	s.ContactName = &v
	return s
}

func (s *DeleteContactRequest) Validate() error {
	return dara.Validate(s)
}
