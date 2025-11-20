// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateListByUserIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOffset(v int64) *GetTemplateListByUserIdRequest
	GetOffset() *int64
	SetSize(v int64) *GetTemplateListByUserIdRequest
	GetSize() *int64
	SetTenantContext(v *GetTemplateListByUserIdRequestTenantContext) *GetTemplateListByUserIdRequest
	GetTenantContext() *GetTemplateListByUserIdRequestTenantContext
}

type GetTemplateListByUserIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 10
	Offset *int64 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	Size          *int64                                       `json:"Size,omitempty" xml:"Size,omitempty"`
	TenantContext *GetTemplateListByUserIdRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s GetTemplateListByUserIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateListByUserIdRequest) GoString() string {
	return s.String()
}

func (s *GetTemplateListByUserIdRequest) GetOffset() *int64 {
	return s.Offset
}

func (s *GetTemplateListByUserIdRequest) GetSize() *int64 {
	return s.Size
}

func (s *GetTemplateListByUserIdRequest) GetTenantContext() *GetTemplateListByUserIdRequestTenantContext {
	return s.TenantContext
}

func (s *GetTemplateListByUserIdRequest) SetOffset(v int64) *GetTemplateListByUserIdRequest {
	s.Offset = &v
	return s
}

func (s *GetTemplateListByUserIdRequest) SetSize(v int64) *GetTemplateListByUserIdRequest {
	s.Size = &v
	return s
}

func (s *GetTemplateListByUserIdRequest) SetTenantContext(v *GetTemplateListByUserIdRequestTenantContext) *GetTemplateListByUserIdRequest {
	s.TenantContext = v
	return s
}

func (s *GetTemplateListByUserIdRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTemplateListByUserIdRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetTemplateListByUserIdRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateListByUserIdRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetTemplateListByUserIdRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetTemplateListByUserIdRequestTenantContext) SetTenantId(v string) *GetTemplateListByUserIdRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetTemplateListByUserIdRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
