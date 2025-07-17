// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContactGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactGroupId(v int64) *DeleteContactGroupRequest
	GetContactGroupId() *int64
}

type DeleteContactGroupRequest struct {
	// The ID of the alert contact group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	ContactGroupId *int64 `json:"ContactGroupId,omitempty" xml:"ContactGroupId,omitempty"`
}

func (s DeleteContactGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteContactGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteContactGroupRequest) GetContactGroupId() *int64 {
	return s.ContactGroupId
}

func (s *DeleteContactGroupRequest) SetContactGroupId(v int64) *DeleteContactGroupRequest {
	s.ContactGroupId = &v
	return s
}

func (s *DeleteContactGroupRequest) Validate() error {
	return dara.Validate(s)
}
