// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDriveSpaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSpaceId(v string) *DeleteDriveSpaceRequest
	GetSpaceId() *string
	SetTenantContext(v *DeleteDriveSpaceRequestTenantContext) *DeleteDriveSpaceRequest
	GetTenantContext() *DeleteDriveSpaceRequestTenantContext
}

type DeleteDriveSpaceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	SpaceId       *string                               `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	TenantContext *DeleteDriveSpaceRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s DeleteDriveSpaceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDriveSpaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteDriveSpaceRequest) GetSpaceId() *string {
	return s.SpaceId
}

func (s *DeleteDriveSpaceRequest) GetTenantContext() *DeleteDriveSpaceRequestTenantContext {
	return s.TenantContext
}

func (s *DeleteDriveSpaceRequest) SetSpaceId(v string) *DeleteDriveSpaceRequest {
	s.SpaceId = &v
	return s
}

func (s *DeleteDriveSpaceRequest) SetTenantContext(v *DeleteDriveSpaceRequestTenantContext) *DeleteDriveSpaceRequest {
	s.TenantContext = v
	return s
}

func (s *DeleteDriveSpaceRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteDriveSpaceRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s DeleteDriveSpaceRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s DeleteDriveSpaceRequestTenantContext) GoString() string {
	return s.String()
}

func (s *DeleteDriveSpaceRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteDriveSpaceRequestTenantContext) SetTenantId(v string) *DeleteDriveSpaceRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *DeleteDriveSpaceRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
