// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyDentryByNodeIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDentryUuid(v string) *CopyDentryByNodeIdRequest
	GetDentryUuid() *string
	SetName(v string) *CopyDentryByNodeIdRequest
	GetName() *string
	SetTenantContext(v *CopyDentryByNodeIdRequestTenantContext) *CopyDentryByNodeIdRequest
	GetTenantContext() *CopyDentryByNodeIdRequestTenantContext
	SetToNextNodeId(v string) *CopyDentryByNodeIdRequest
	GetToNextNodeId() *string
	SetToParentNodeId(v string) *CopyDentryByNodeIdRequest
	GetToParentNodeId() *string
	SetToPrevNodeId(v string) *CopyDentryByNodeIdRequest
	GetToPrevNodeId() *string
}

type CopyDentryByNodeIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// P7QG4Yx2Jpx4OolYC1QPg5BaJ9dEq3XD
	DentryUuid *string `json:"DentryUuid,omitempty" xml:"DentryUuid,omitempty"`
	// This parameter is required.
	Name          *string                                 `json:"Name,omitempty" xml:"Name,omitempty"`
	TenantContext *CopyDentryByNodeIdRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
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

func (s CopyDentryByNodeIdRequest) String() string {
	return dara.Prettify(s)
}

func (s CopyDentryByNodeIdRequest) GoString() string {
	return s.String()
}

func (s *CopyDentryByNodeIdRequest) GetDentryUuid() *string {
	return s.DentryUuid
}

func (s *CopyDentryByNodeIdRequest) GetName() *string {
	return s.Name
}

func (s *CopyDentryByNodeIdRequest) GetTenantContext() *CopyDentryByNodeIdRequestTenantContext {
	return s.TenantContext
}

func (s *CopyDentryByNodeIdRequest) GetToNextNodeId() *string {
	return s.ToNextNodeId
}

func (s *CopyDentryByNodeIdRequest) GetToParentNodeId() *string {
	return s.ToParentNodeId
}

func (s *CopyDentryByNodeIdRequest) GetToPrevNodeId() *string {
	return s.ToPrevNodeId
}

func (s *CopyDentryByNodeIdRequest) SetDentryUuid(v string) *CopyDentryByNodeIdRequest {
	s.DentryUuid = &v
	return s
}

func (s *CopyDentryByNodeIdRequest) SetName(v string) *CopyDentryByNodeIdRequest {
	s.Name = &v
	return s
}

func (s *CopyDentryByNodeIdRequest) SetTenantContext(v *CopyDentryByNodeIdRequestTenantContext) *CopyDentryByNodeIdRequest {
	s.TenantContext = v
	return s
}

func (s *CopyDentryByNodeIdRequest) SetToNextNodeId(v string) *CopyDentryByNodeIdRequest {
	s.ToNextNodeId = &v
	return s
}

func (s *CopyDentryByNodeIdRequest) SetToParentNodeId(v string) *CopyDentryByNodeIdRequest {
	s.ToParentNodeId = &v
	return s
}

func (s *CopyDentryByNodeIdRequest) SetToPrevNodeId(v string) *CopyDentryByNodeIdRequest {
	s.ToPrevNodeId = &v
	return s
}

func (s *CopyDentryByNodeIdRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CopyDentryByNodeIdRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CopyDentryByNodeIdRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s CopyDentryByNodeIdRequestTenantContext) GoString() string {
	return s.String()
}

func (s *CopyDentryByNodeIdRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *CopyDentryByNodeIdRequestTenantContext) SetTenantId(v string) *CopyDentryByNodeIdRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *CopyDentryByNodeIdRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
