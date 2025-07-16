// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGrantPermissionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPermissions(v []*Permission) *BatchGrantPermissionsRequest
	GetPermissions() []*Permission
}

type BatchGrantPermissionsRequest struct {
	Permissions []*Permission `json:"permissions,omitempty" xml:"permissions,omitempty" type:"Repeated"`
}

func (s BatchGrantPermissionsRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchGrantPermissionsRequest) GoString() string {
	return s.String()
}

func (s *BatchGrantPermissionsRequest) GetPermissions() []*Permission {
	return s.Permissions
}

func (s *BatchGrantPermissionsRequest) SetPermissions(v []*Permission) *BatchGrantPermissionsRequest {
	s.Permissions = v
	return s
}

func (s *BatchGrantPermissionsRequest) Validate() error {
	return dara.Validate(s)
}
