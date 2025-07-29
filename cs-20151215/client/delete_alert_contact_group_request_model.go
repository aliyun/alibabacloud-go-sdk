// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertContactGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactGroupIds(v []*int64) *DeleteAlertContactGroupRequest
	GetContactGroupIds() []*int64
}

type DeleteAlertContactGroupRequest struct {
	// The list of alert contact group IDs.
	//
	// This parameter is required.
	ContactGroupIds []*int64 `json:"contact_group_ids,omitempty" xml:"contact_group_ids,omitempty" type:"Repeated"`
}

func (s DeleteAlertContactGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertContactGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlertContactGroupRequest) GetContactGroupIds() []*int64 {
	return s.ContactGroupIds
}

func (s *DeleteAlertContactGroupRequest) SetContactGroupIds(v []*int64) *DeleteAlertContactGroupRequest {
	s.ContactGroupIds = v
	return s
}

func (s *DeleteAlertContactGroupRequest) Validate() error {
	return dara.Validate(s)
}
