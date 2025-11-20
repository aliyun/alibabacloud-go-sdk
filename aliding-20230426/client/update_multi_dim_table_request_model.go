// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMultiDimTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseId(v string) *UpdateMultiDimTableRequest
	GetBaseId() *string
	SetName(v string) *UpdateMultiDimTableRequest
	GetName() *string
	SetSheetIdOrName(v string) *UpdateMultiDimTableRequest
	GetSheetIdOrName() *string
	SetTenantContext(v *UpdateMultiDimTableRequestTenantContext) *UpdateMultiDimTableRequest
	GetTenantContext() *UpdateMultiDimTableRequestTenantContext
}

type UpdateMultiDimTableRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// r1R7q3QmWew5lo02fxB7noNyJxxxxxx
	BaseId *string `json:"BaseId,omitempty" xml:"BaseId,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	SheetIdOrName *string                                  `json:"SheetIdOrName,omitempty" xml:"SheetIdOrName,omitempty"`
	TenantContext *UpdateMultiDimTableRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s UpdateMultiDimTableRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultiDimTableRequest) GoString() string {
	return s.String()
}

func (s *UpdateMultiDimTableRequest) GetBaseId() *string {
	return s.BaseId
}

func (s *UpdateMultiDimTableRequest) GetName() *string {
	return s.Name
}

func (s *UpdateMultiDimTableRequest) GetSheetIdOrName() *string {
	return s.SheetIdOrName
}

func (s *UpdateMultiDimTableRequest) GetTenantContext() *UpdateMultiDimTableRequestTenantContext {
	return s.TenantContext
}

func (s *UpdateMultiDimTableRequest) SetBaseId(v string) *UpdateMultiDimTableRequest {
	s.BaseId = &v
	return s
}

func (s *UpdateMultiDimTableRequest) SetName(v string) *UpdateMultiDimTableRequest {
	s.Name = &v
	return s
}

func (s *UpdateMultiDimTableRequest) SetSheetIdOrName(v string) *UpdateMultiDimTableRequest {
	s.SheetIdOrName = &v
	return s
}

func (s *UpdateMultiDimTableRequest) SetTenantContext(v *UpdateMultiDimTableRequestTenantContext) *UpdateMultiDimTableRequest {
	s.TenantContext = v
	return s
}

func (s *UpdateMultiDimTableRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMultiDimTableRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateMultiDimTableRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultiDimTableRequestTenantContext) GoString() string {
	return s.String()
}

func (s *UpdateMultiDimTableRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateMultiDimTableRequestTenantContext) SetTenantId(v string) *UpdateMultiDimTableRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *UpdateMultiDimTableRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
