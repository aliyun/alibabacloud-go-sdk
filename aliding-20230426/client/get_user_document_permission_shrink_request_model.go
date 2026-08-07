// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserDocumentPermissionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDentryId(v int64) *GetUserDocumentPermissionShrinkRequest
	GetDentryId() *int64
	SetDentryUuid(v string) *GetUserDocumentPermissionShrinkRequest
	GetDentryUuid() *string
	SetResourceType(v int32) *GetUserDocumentPermissionShrinkRequest
	GetResourceType() *int32
	SetSpaceId(v int64) *GetUserDocumentPermissionShrinkRequest
	GetSpaceId() *int64
	SetTenantContextShrink(v string) *GetUserDocumentPermissionShrinkRequest
	GetTenantContextShrink() *string
}

type GetUserDocumentPermissionShrinkRequest struct {
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
	SpaceId             *int64  `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s GetUserDocumentPermissionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserDocumentPermissionShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetUserDocumentPermissionShrinkRequest) GetDentryId() *int64 {
	return s.DentryId
}

func (s *GetUserDocumentPermissionShrinkRequest) GetDentryUuid() *string {
	return s.DentryUuid
}

func (s *GetUserDocumentPermissionShrinkRequest) GetResourceType() *int32 {
	return s.ResourceType
}

func (s *GetUserDocumentPermissionShrinkRequest) GetSpaceId() *int64 {
	return s.SpaceId
}

func (s *GetUserDocumentPermissionShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetUserDocumentPermissionShrinkRequest) SetDentryId(v int64) *GetUserDocumentPermissionShrinkRequest {
	s.DentryId = &v
	return s
}

func (s *GetUserDocumentPermissionShrinkRequest) SetDentryUuid(v string) *GetUserDocumentPermissionShrinkRequest {
	s.DentryUuid = &v
	return s
}

func (s *GetUserDocumentPermissionShrinkRequest) SetResourceType(v int32) *GetUserDocumentPermissionShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *GetUserDocumentPermissionShrinkRequest) SetSpaceId(v int64) *GetUserDocumentPermissionShrinkRequest {
	s.SpaceId = &v
	return s
}

func (s *GetUserDocumentPermissionShrinkRequest) SetTenantContextShrink(v string) *GetUserDocumentPermissionShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetUserDocumentPermissionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
