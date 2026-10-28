// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContactGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactGroupIds(v []*string) *DeleteContactGroupsRequest
	GetContactGroupIds() []*string
}

type DeleteContactGroupsRequest struct {
	// The IDs of the contact groups.
	//
	// This parameter is required.
	ContactGroupIds []*string `json:"contactGroupIds,omitempty" xml:"contactGroupIds,omitempty" type:"Repeated"`
}

func (s DeleteContactGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteContactGroupsRequest) GoString() string {
	return s.String()
}

func (s *DeleteContactGroupsRequest) GetContactGroupIds() []*string {
	return s.ContactGroupIds
}

func (s *DeleteContactGroupsRequest) SetContactGroupIds(v []*string) *DeleteContactGroupsRequest {
	s.ContactGroupIds = v
	return s
}

func (s *DeleteContactGroupsRequest) Validate() error {
	return dara.Validate(s)
}
