// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDriveSpaceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSpaceId(v string) *DeleteDriveSpaceShrinkRequest
	GetSpaceId() *string
	SetTenantContextShrink(v string) *DeleteDriveSpaceShrinkRequest
	GetTenantContextShrink() *string
}

type DeleteDriveSpaceShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	SpaceId             *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s DeleteDriveSpaceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDriveSpaceShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteDriveSpaceShrinkRequest) GetSpaceId() *string {
	return s.SpaceId
}

func (s *DeleteDriveSpaceShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *DeleteDriveSpaceShrinkRequest) SetSpaceId(v string) *DeleteDriveSpaceShrinkRequest {
	s.SpaceId = &v
	return s
}

func (s *DeleteDriveSpaceShrinkRequest) SetTenantContextShrink(v string) *DeleteDriveSpaceShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *DeleteDriveSpaceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
