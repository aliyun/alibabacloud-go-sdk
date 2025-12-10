// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceGroupAdminSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreatorAsAdmin(v bool) *UpdateResourceGroupAdminSettingRequest
	GetCreatorAsAdmin() *bool
}

type UpdateResourceGroupAdminSettingRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	CreatorAsAdmin *bool `json:"CreatorAsAdmin,omitempty" xml:"CreatorAsAdmin,omitempty"`
}

func (s UpdateResourceGroupAdminSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceGroupAdminSettingRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceGroupAdminSettingRequest) GetCreatorAsAdmin() *bool {
	return s.CreatorAsAdmin
}

func (s *UpdateResourceGroupAdminSettingRequest) SetCreatorAsAdmin(v bool) *UpdateResourceGroupAdminSettingRequest {
	s.CreatorAsAdmin = &v
	return s
}

func (s *UpdateResourceGroupAdminSettingRequest) Validate() error {
	return dara.Validate(s)
}
