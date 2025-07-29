// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertContactShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactIdsShrink(v string) *DeleteAlertContactShrinkRequest
	GetContactIdsShrink() *string
}

type DeleteAlertContactShrinkRequest struct {
	// The list of alert contact IDs.
	//
	// This parameter is required.
	ContactIdsShrink *string `json:"contact_ids,omitempty" xml:"contact_ids,omitempty"`
}

func (s DeleteAlertContactShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertContactShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlertContactShrinkRequest) GetContactIdsShrink() *string {
	return s.ContactIdsShrink
}

func (s *DeleteAlertContactShrinkRequest) SetContactIdsShrink(v string) *DeleteAlertContactShrinkRequest {
	s.ContactIdsShrink = &v
	return s
}

func (s *DeleteAlertContactShrinkRequest) Validate() error {
	return dara.Validate(s)
}
