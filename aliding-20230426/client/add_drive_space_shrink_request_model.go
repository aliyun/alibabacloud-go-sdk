// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDriveSpaceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *AddDriveSpaceShrinkRequest
	GetName() *string
	SetTenantContextShrink(v string) *AddDriveSpaceShrinkRequest
	GetTenantContextShrink() *string
}

type AddDriveSpaceShrinkRequest struct {
	// This parameter is required.
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s AddDriveSpaceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDriveSpaceShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddDriveSpaceShrinkRequest) GetName() *string {
	return s.Name
}

func (s *AddDriveSpaceShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *AddDriveSpaceShrinkRequest) SetName(v string) *AddDriveSpaceShrinkRequest {
	s.Name = &v
	return s
}

func (s *AddDriveSpaceShrinkRequest) SetTenantContextShrink(v string) *AddDriveSpaceShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *AddDriveSpaceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
