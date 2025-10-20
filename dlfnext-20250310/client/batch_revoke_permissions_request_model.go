// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchRevokePermissionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPermissions(v []*Permission) *BatchRevokePermissionsRequest
	GetPermissions() []*Permission
}

type BatchRevokePermissionsRequest struct {
	Permissions []*Permission `json:"permissions,omitempty" xml:"permissions,omitempty" type:"Repeated"`
}

func (s BatchRevokePermissionsRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchRevokePermissionsRequest) GoString() string {
	return s.String()
}

func (s *BatchRevokePermissionsRequest) GetPermissions() []*Permission {
	return s.Permissions
}

func (s *BatchRevokePermissionsRequest) SetPermissions(v []*Permission) *BatchRevokePermissionsRequest {
	s.Permissions = v
	return s
}

func (s *BatchRevokePermissionsRequest) Validate() error {
	if s.Permissions != nil {
		for _, item := range s.Permissions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
