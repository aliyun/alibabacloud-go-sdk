// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCascadeDelete(v string) *DeleteGroupRequest
	GetCascadeDelete() *string
}

type DeleteGroupRequest struct {
	// Specifies whether to perform a cascade delete. If enabled, deleting the service group automatically deletes all services within the service group. This feature is disabled by default.
	//
	// example:
	//
	// false
	CascadeDelete *string `json:"CascadeDelete,omitempty" xml:"CascadeDelete,omitempty"`
}

func (s DeleteGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteGroupRequest) GetCascadeDelete() *string {
	return s.CascadeDelete
}

func (s *DeleteGroupRequest) SetCascadeDelete(v string) *DeleteGroupRequest {
	s.CascadeDelete = &v
	return s
}

func (s *DeleteGroupRequest) Validate() error {
	return dara.Validate(s)
}
