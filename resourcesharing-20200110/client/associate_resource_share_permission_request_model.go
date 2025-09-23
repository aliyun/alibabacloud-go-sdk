// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateResourceSharePermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPermissionName(v string) *AssociateResourceSharePermissionRequest
	GetPermissionName() *string
	SetReplace(v bool) *AssociateResourceSharePermissionRequest
	GetReplace() *bool
	SetResourceShareId(v string) *AssociateResourceSharePermissionRequest
	GetResourceShareId() *string
}

type AssociateResourceSharePermissionRequest struct {
	// The name of the permission.
	//
	// This parameter is required.
	//
	// example:
	//
	// AliyunRSDefaultPermissionVSwitch
	PermissionName *string `json:"PermissionName,omitempty" xml:"PermissionName,omitempty"`
	// Specifies whether to use the specified permission to replace an existing permission. Valid values:
	//
	// 	- false: does not use the specified permission to replace an existing permission. This is the default value. If you set the value to false for a resource share that does not have associated permissions, the system associates the specified permission with the resource share. In a resource share, one resource type can have only one permission. If you set the value to false for a resource share that already has a permission for the resource type indicated by the specified permission, the system reports an error. This prevents you from replacing the existing permission by mistake.
	//
	// 	- true: uses the specified permission to replace an existing permission of the same resource type.
	//
	// example:
	//
	// false
	Replace *bool `json:"Replace,omitempty" xml:"Replace,omitempty"`
	// The ID of the resource share.
	//
	// This parameter is required.
	//
	// example:
	//
	// rs-6GRmdD3X****
	ResourceShareId *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
}

func (s AssociateResourceSharePermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateResourceSharePermissionRequest) GoString() string {
	return s.String()
}

func (s *AssociateResourceSharePermissionRequest) GetPermissionName() *string {
	return s.PermissionName
}

func (s *AssociateResourceSharePermissionRequest) GetReplace() *bool {
	return s.Replace
}

func (s *AssociateResourceSharePermissionRequest) GetResourceShareId() *string {
	return s.ResourceShareId
}

func (s *AssociateResourceSharePermissionRequest) SetPermissionName(v string) *AssociateResourceSharePermissionRequest {
	s.PermissionName = &v
	return s
}

func (s *AssociateResourceSharePermissionRequest) SetReplace(v bool) *AssociateResourceSharePermissionRequest {
	s.Replace = &v
	return s
}

func (s *AssociateResourceSharePermissionRequest) SetResourceShareId(v string) *AssociateResourceSharePermissionRequest {
	s.ResourceShareId = &v
	return s
}

func (s *AssociateResourceSharePermissionRequest) Validate() error {
	return dara.Validate(s)
}
