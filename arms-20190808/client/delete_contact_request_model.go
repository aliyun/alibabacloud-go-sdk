// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactId(v int64) *DeleteContactRequest
	GetContactId() *int64
}

type DeleteContactRequest struct {
	// The ID of the alert contact.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	ContactId *int64 `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
}

func (s DeleteContactRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteContactRequest) GoString() string {
	return s.String()
}

func (s *DeleteContactRequest) GetContactId() *int64 {
	return s.ContactId
}

func (s *DeleteContactRequest) SetContactId(v int64) *DeleteContactRequest {
	s.ContactId = &v
	return s
}

func (s *DeleteContactRequest) Validate() error {
	return dara.Validate(s)
}
