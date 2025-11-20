// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMultiDimTableFieldRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseId(v string) *UpdateMultiDimTableFieldRequest
	GetBaseId() *string
	SetFieldIdOrName(v string) *UpdateMultiDimTableFieldRequest
	GetFieldIdOrName() *string
	SetName(v string) *UpdateMultiDimTableFieldRequest
	GetName() *string
	SetProperty(v map[string]interface{}) *UpdateMultiDimTableFieldRequest
	GetProperty() map[string]interface{}
	SetSheetIdOrName(v string) *UpdateMultiDimTableFieldRequest
	GetSheetIdOrName() *string
	SetTenantContext(v *UpdateMultiDimTableFieldRequestTenantContext) *UpdateMultiDimTableFieldRequest
	GetTenantContext() *UpdateMultiDimTableFieldRequestTenantContext
}

type UpdateMultiDimTableFieldRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// r1R7q3QmWew5lo02fxB7noNyJxxxxxx
	BaseId *string `json:"BaseId,omitempty" xml:"BaseId,omitempty"`
	// This parameter is required.
	FieldIdOrName *string `json:"FieldIdOrName,omitempty" xml:"FieldIdOrName,omitempty"`
	// This parameter is required.
	Name     *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	Property map[string]interface{} `json:"Property,omitempty" xml:"Property,omitempty"`
	// This parameter is required.
	SheetIdOrName *string                                       `json:"SheetIdOrName,omitempty" xml:"SheetIdOrName,omitempty"`
	TenantContext *UpdateMultiDimTableFieldRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s UpdateMultiDimTableFieldRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultiDimTableFieldRequest) GoString() string {
	return s.String()
}

func (s *UpdateMultiDimTableFieldRequest) GetBaseId() *string {
	return s.BaseId
}

func (s *UpdateMultiDimTableFieldRequest) GetFieldIdOrName() *string {
	return s.FieldIdOrName
}

func (s *UpdateMultiDimTableFieldRequest) GetName() *string {
	return s.Name
}

func (s *UpdateMultiDimTableFieldRequest) GetProperty() map[string]interface{} {
	return s.Property
}

func (s *UpdateMultiDimTableFieldRequest) GetSheetIdOrName() *string {
	return s.SheetIdOrName
}

func (s *UpdateMultiDimTableFieldRequest) GetTenantContext() *UpdateMultiDimTableFieldRequestTenantContext {
	return s.TenantContext
}

func (s *UpdateMultiDimTableFieldRequest) SetBaseId(v string) *UpdateMultiDimTableFieldRequest {
	s.BaseId = &v
	return s
}

func (s *UpdateMultiDimTableFieldRequest) SetFieldIdOrName(v string) *UpdateMultiDimTableFieldRequest {
	s.FieldIdOrName = &v
	return s
}

func (s *UpdateMultiDimTableFieldRequest) SetName(v string) *UpdateMultiDimTableFieldRequest {
	s.Name = &v
	return s
}

func (s *UpdateMultiDimTableFieldRequest) SetProperty(v map[string]interface{}) *UpdateMultiDimTableFieldRequest {
	s.Property = v
	return s
}

func (s *UpdateMultiDimTableFieldRequest) SetSheetIdOrName(v string) *UpdateMultiDimTableFieldRequest {
	s.SheetIdOrName = &v
	return s
}

func (s *UpdateMultiDimTableFieldRequest) SetTenantContext(v *UpdateMultiDimTableFieldRequestTenantContext) *UpdateMultiDimTableFieldRequest {
	s.TenantContext = v
	return s
}

func (s *UpdateMultiDimTableFieldRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMultiDimTableFieldRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateMultiDimTableFieldRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultiDimTableFieldRequestTenantContext) GoString() string {
	return s.String()
}

func (s *UpdateMultiDimTableFieldRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateMultiDimTableFieldRequestTenantContext) SetTenantId(v string) *UpdateMultiDimTableFieldRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *UpdateMultiDimTableFieldRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
