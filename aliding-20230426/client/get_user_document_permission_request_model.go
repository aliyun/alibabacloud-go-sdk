// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserDocumentPermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDentryId(v int64) *GetUserDocumentPermissionRequest
	GetDentryId() *int64
	SetDentryUuid(v string) *GetUserDocumentPermissionRequest
	GetDentryUuid() *string
	SetResourceType(v int32) *GetUserDocumentPermissionRequest
	GetResourceType() *int32
	SetSpaceId(v int64) *GetUserDocumentPermissionRequest
	GetSpaceId() *int64
	SetTenantContext(v *GetUserDocumentPermissionRequestTenantContext) *GetUserDocumentPermissionRequest
	GetTenantContext() *GetUserDocumentPermissionRequestTenantContext
}

type GetUserDocumentPermissionRequest struct {
	// example:
	//
	// 87654321
	DentryId *int64 `json:"DentryId,omitempty" xml:"DentryId,omitempty"`
	// example:
	//
	// a9E05BDRVQRkezKGCDOvkbzrJ63zgkYA
	DentryUuid *string `json:"DentryUuid,omitempty" xml:"DentryUuid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	ResourceType *int32 `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// 12345678
	SpaceId       *int64                                         `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	TenantContext *GetUserDocumentPermissionRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s GetUserDocumentPermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserDocumentPermissionRequest) GoString() string {
	return s.String()
}

func (s *GetUserDocumentPermissionRequest) GetDentryId() *int64 {
	return s.DentryId
}

func (s *GetUserDocumentPermissionRequest) GetDentryUuid() *string {
	return s.DentryUuid
}

func (s *GetUserDocumentPermissionRequest) GetResourceType() *int32 {
	return s.ResourceType
}

func (s *GetUserDocumentPermissionRequest) GetSpaceId() *int64 {
	return s.SpaceId
}

func (s *GetUserDocumentPermissionRequest) GetTenantContext() *GetUserDocumentPermissionRequestTenantContext {
	return s.TenantContext
}

func (s *GetUserDocumentPermissionRequest) SetDentryId(v int64) *GetUserDocumentPermissionRequest {
	s.DentryId = &v
	return s
}

func (s *GetUserDocumentPermissionRequest) SetDentryUuid(v string) *GetUserDocumentPermissionRequest {
	s.DentryUuid = &v
	return s
}

func (s *GetUserDocumentPermissionRequest) SetResourceType(v int32) *GetUserDocumentPermissionRequest {
	s.ResourceType = &v
	return s
}

func (s *GetUserDocumentPermissionRequest) SetSpaceId(v int64) *GetUserDocumentPermissionRequest {
	s.SpaceId = &v
	return s
}

func (s *GetUserDocumentPermissionRequest) SetTenantContext(v *GetUserDocumentPermissionRequestTenantContext) *GetUserDocumentPermissionRequest {
	s.TenantContext = v
	return s
}

func (s *GetUserDocumentPermissionRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserDocumentPermissionRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetUserDocumentPermissionRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetUserDocumentPermissionRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetUserDocumentPermissionRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetUserDocumentPermissionRequestTenantContext) SetTenantId(v string) *GetUserDocumentPermissionRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetUserDocumentPermissionRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
