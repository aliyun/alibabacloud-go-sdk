// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisassociateResourceSharePermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPermissionName(v string) *DisassociateResourceSharePermissionRequest
	GetPermissionName() *string
	SetResourceShareId(v string) *DisassociateResourceSharePermissionRequest
	GetResourceShareId() *string
}

type DisassociateResourceSharePermissionRequest struct {
	// The name of the permission. For more information, see [Permission library](https://help.aliyun.com/document_detail/465474.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// AliyunRSDefaultPermissionVSwitch
	PermissionName *string `json:"PermissionName,omitempty" xml:"PermissionName,omitempty"`
	// The ID of the resource share.
	//
	// This parameter is required.
	//
	// example:
	//
	// rs-6GRmdD3X****
	ResourceShareId *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
}

func (s DisassociateResourceSharePermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s DisassociateResourceSharePermissionRequest) GoString() string {
	return s.String()
}

func (s *DisassociateResourceSharePermissionRequest) GetPermissionName() *string {
	return s.PermissionName
}

func (s *DisassociateResourceSharePermissionRequest) GetResourceShareId() *string {
	return s.ResourceShareId
}

func (s *DisassociateResourceSharePermissionRequest) SetPermissionName(v string) *DisassociateResourceSharePermissionRequest {
	s.PermissionName = &v
	return s
}

func (s *DisassociateResourceSharePermissionRequest) SetResourceShareId(v string) *DisassociateResourceSharePermissionRequest {
	s.ResourceShareId = &v
	return s
}

func (s *DisassociateResourceSharePermissionRequest) Validate() error {
	return dara.Validate(s)
}
