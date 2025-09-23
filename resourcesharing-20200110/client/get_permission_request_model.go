// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPermissionName(v string) *GetPermissionRequest
	GetPermissionName() *string
	SetPermissionVersion(v string) *GetPermissionRequest
	GetPermissionVersion() *string
}

type GetPermissionRequest struct {
	// The name of the permission.
	//
	// This parameter is required.
	//
	// example:
	//
	// AliyunRSDefaultPermissionVSwitch
	PermissionName *string `json:"PermissionName,omitempty" xml:"PermissionName,omitempty"`
	// The version of the permission.
	//
	// example:
	//
	// v1
	PermissionVersion *string `json:"PermissionVersion,omitempty" xml:"PermissionVersion,omitempty"`
}

func (s GetPermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPermissionRequest) GoString() string {
	return s.String()
}

func (s *GetPermissionRequest) GetPermissionName() *string {
	return s.PermissionName
}

func (s *GetPermissionRequest) GetPermissionVersion() *string {
	return s.PermissionVersion
}

func (s *GetPermissionRequest) SetPermissionName(v string) *GetPermissionRequest {
	s.PermissionName = &v
	return s
}

func (s *GetPermissionRequest) SetPermissionVersion(v string) *GetPermissionRequest {
	s.PermissionVersion = &v
	return s
}

func (s *GetPermissionRequest) Validate() error {
	return dara.Validate(s)
}
