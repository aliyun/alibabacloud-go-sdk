// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetMemberDeletionPermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetStatus(v string) *SetMemberDeletionPermissionRequest
	GetStatus() *string
}

type SetMemberDeletionPermissionRequest struct {
	// Specifies whether to enable the member deletion feature. Valid values:
	//
	// 	- Enabled: enables the member deletion feature
	//
	// 	- Disabled: disables the member deletion feature
	//
	// This parameter is required.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SetMemberDeletionPermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s SetMemberDeletionPermissionRequest) GoString() string {
	return s.String()
}

func (s *SetMemberDeletionPermissionRequest) GetStatus() *string {
	return s.Status
}

func (s *SetMemberDeletionPermissionRequest) SetStatus(v string) *SetMemberDeletionPermissionRequest {
	s.Status = &v
	return s
}

func (s *SetMemberDeletionPermissionRequest) Validate() error {
	return dara.Validate(s)
}
