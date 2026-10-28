// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContactsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactIds(v []*string) *DeleteContactsRequest
	GetContactIds() []*string
}

type DeleteContactsRequest struct {
	// The IDs of the contacts.
	ContactIds []*string `json:"contactIds,omitempty" xml:"contactIds,omitempty" type:"Repeated"`
}

func (s DeleteContactsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteContactsRequest) GoString() string {
	return s.String()
}

func (s *DeleteContactsRequest) GetContactIds() []*string {
	return s.ContactIds
}

func (s *DeleteContactsRequest) SetContactIds(v []*string) *DeleteContactsRequest {
	s.ContactIds = v
	return s
}

func (s *DeleteContactsRequest) Validate() error {
	return dara.Validate(s)
}
