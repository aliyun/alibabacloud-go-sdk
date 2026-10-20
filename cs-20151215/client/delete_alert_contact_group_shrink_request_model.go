// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertContactGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactGroupIdsShrink(v string) *DeleteAlertContactGroupShrinkRequest
	GetContactGroupIdsShrink() *string
}

type DeleteAlertContactGroupShrinkRequest struct {
	// The list of alert contact group IDs.
	//
	// This parameter is required.
	ContactGroupIdsShrink *string `json:"contact_group_ids,omitempty" xml:"contact_group_ids,omitempty"`
}

func (s DeleteAlertContactGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertContactGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlertContactGroupShrinkRequest) GetContactGroupIdsShrink() *string {
	return s.ContactGroupIdsShrink
}

func (s *DeleteAlertContactGroupShrinkRequest) SetContactGroupIdsShrink(v string) *DeleteAlertContactGroupShrinkRequest {
	s.ContactGroupIdsShrink = &v
	return s
}

func (s *DeleteAlertContactGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
