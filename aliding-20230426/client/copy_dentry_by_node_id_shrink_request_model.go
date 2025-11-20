// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyDentryByNodeIdShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDentryUuid(v string) *CopyDentryByNodeIdShrinkRequest
	GetDentryUuid() *string
	SetName(v string) *CopyDentryByNodeIdShrinkRequest
	GetName() *string
	SetTenantContextShrink(v string) *CopyDentryByNodeIdShrinkRequest
	GetTenantContextShrink() *string
	SetToNextNodeId(v string) *CopyDentryByNodeIdShrinkRequest
	GetToNextNodeId() *string
	SetToParentNodeId(v string) *CopyDentryByNodeIdShrinkRequest
	GetToParentNodeId() *string
	SetToPrevNodeId(v string) *CopyDentryByNodeIdShrinkRequest
	GetToPrevNodeId() *string
}

type CopyDentryByNodeIdShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// P7QG4Yx2Jpx4OolYC1QPg5BaJ9dEq3XD
	DentryUuid *string `json:"DentryUuid,omitempty" xml:"DentryUuid,omitempty"`
	// This parameter is required.
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// example:
	//
	// P7QG4Yx2Jpx4OolYC1QPg5BaJ9dEq3XD
	ToNextNodeId *string `json:"ToNextNodeId,omitempty" xml:"ToNextNodeId,omitempty"`
	// example:
	//
	// P7QG4Yx2Jpx4OolYC1QPg5BaJ9dEq3XD
	ToParentNodeId *string `json:"ToParentNodeId,omitempty" xml:"ToParentNodeId,omitempty"`
	// example:
	//
	// P7QG4Yx2Jpx4OolYC1QPg5BaJ9dEq3XD
	ToPrevNodeId *string `json:"ToPrevNodeId,omitempty" xml:"ToPrevNodeId,omitempty"`
}

func (s CopyDentryByNodeIdShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CopyDentryByNodeIdShrinkRequest) GoString() string {
	return s.String()
}

func (s *CopyDentryByNodeIdShrinkRequest) GetDentryUuid() *string {
	return s.DentryUuid
}

func (s *CopyDentryByNodeIdShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CopyDentryByNodeIdShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *CopyDentryByNodeIdShrinkRequest) GetToNextNodeId() *string {
	return s.ToNextNodeId
}

func (s *CopyDentryByNodeIdShrinkRequest) GetToParentNodeId() *string {
	return s.ToParentNodeId
}

func (s *CopyDentryByNodeIdShrinkRequest) GetToPrevNodeId() *string {
	return s.ToPrevNodeId
}

func (s *CopyDentryByNodeIdShrinkRequest) SetDentryUuid(v string) *CopyDentryByNodeIdShrinkRequest {
	s.DentryUuid = &v
	return s
}

func (s *CopyDentryByNodeIdShrinkRequest) SetName(v string) *CopyDentryByNodeIdShrinkRequest {
	s.Name = &v
	return s
}

func (s *CopyDentryByNodeIdShrinkRequest) SetTenantContextShrink(v string) *CopyDentryByNodeIdShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *CopyDentryByNodeIdShrinkRequest) SetToNextNodeId(v string) *CopyDentryByNodeIdShrinkRequest {
	s.ToNextNodeId = &v
	return s
}

func (s *CopyDentryByNodeIdShrinkRequest) SetToParentNodeId(v string) *CopyDentryByNodeIdShrinkRequest {
	s.ToParentNodeId = &v
	return s
}

func (s *CopyDentryByNodeIdShrinkRequest) SetToPrevNodeId(v string) *CopyDentryByNodeIdShrinkRequest {
	s.ToPrevNodeId = &v
	return s
}

func (s *CopyDentryByNodeIdShrinkRequest) Validate() error {
	return dara.Validate(s)
}
