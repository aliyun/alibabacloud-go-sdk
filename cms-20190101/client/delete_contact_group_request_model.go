// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContactGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactGroupName(v string) *DeleteContactGroupRequest
	GetContactGroupName() *string
}

type DeleteContactGroupRequest struct {
	// The name of the alert contact group.
	//
	// This parameter is required.
	//
	// example:
	//
	// contact_group_2019_templatedfkXQ
	ContactGroupName *string `json:"ContactGroupName,omitempty" xml:"ContactGroupName,omitempty"`
}

func (s DeleteContactGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteContactGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteContactGroupRequest) GetContactGroupName() *string {
	return s.ContactGroupName
}

func (s *DeleteContactGroupRequest) SetContactGroupName(v string) *DeleteContactGroupRequest {
	s.ContactGroupName = &v
	return s
}

func (s *DeleteContactGroupRequest) Validate() error {
	return dara.Validate(s)
}
