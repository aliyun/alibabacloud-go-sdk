// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertContactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactIds(v []*int64) *DeleteAlertContactRequest
	GetContactIds() []*int64
}

type DeleteAlertContactRequest struct {
	// The list of alert contact IDs.
	//
	// This parameter is required.
	ContactIds []*int64 `json:"contact_ids,omitempty" xml:"contact_ids,omitempty" type:"Repeated"`
}

func (s DeleteAlertContactRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertContactRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlertContactRequest) GetContactIds() []*int64 {
	return s.ContactIds
}

func (s *DeleteAlertContactRequest) SetContactIds(v []*int64) *DeleteAlertContactRequest {
	s.ContactIds = v
	return s
}

func (s *DeleteAlertContactRequest) Validate() error {
	return dara.Validate(s)
}
