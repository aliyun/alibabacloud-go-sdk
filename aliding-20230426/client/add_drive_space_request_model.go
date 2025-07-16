// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDriveSpaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *AddDriveSpaceRequest
	GetName() *string
	SetTenantContext(v *AddDriveSpaceRequestTenantContext) *AddDriveSpaceRequest
	GetTenantContext() *AddDriveSpaceRequestTenantContext
}

type AddDriveSpaceRequest struct {
	// This parameter is required.
	Name          *string                            `json:"Name,omitempty" xml:"Name,omitempty"`
	TenantContext *AddDriveSpaceRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s AddDriveSpaceRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDriveSpaceRequest) GoString() string {
	return s.String()
}

func (s *AddDriveSpaceRequest) GetName() *string {
	return s.Name
}

func (s *AddDriveSpaceRequest) GetTenantContext() *AddDriveSpaceRequestTenantContext {
	return s.TenantContext
}

func (s *AddDriveSpaceRequest) SetName(v string) *AddDriveSpaceRequest {
	s.Name = &v
	return s
}

func (s *AddDriveSpaceRequest) SetTenantContext(v *AddDriveSpaceRequestTenantContext) *AddDriveSpaceRequest {
	s.TenantContext = v
	return s
}

func (s *AddDriveSpaceRequest) Validate() error {
	return dara.Validate(s)
}

type AddDriveSpaceRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s AddDriveSpaceRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s AddDriveSpaceRequestTenantContext) GoString() string {
	return s.String()
}

func (s *AddDriveSpaceRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *AddDriveSpaceRequestTenantContext) SetTenantId(v string) *AddDriveSpaceRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *AddDriveSpaceRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
