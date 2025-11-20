// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultiDimTableFieldRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseId(v string) *CreateMultiDimTableFieldRequest
	GetBaseId() *string
	SetName(v string) *CreateMultiDimTableFieldRequest
	GetName() *string
	SetProperty(v map[string]interface{}) *CreateMultiDimTableFieldRequest
	GetProperty() map[string]interface{}
	SetSheetIdOrName(v string) *CreateMultiDimTableFieldRequest
	GetSheetIdOrName() *string
	SetTenantContext(v *CreateMultiDimTableFieldRequestTenantContext) *CreateMultiDimTableFieldRequest
	GetTenantContext() *CreateMultiDimTableFieldRequestTenantContext
	SetType(v string) *CreateMultiDimTableFieldRequest
	GetType() *string
}

type CreateMultiDimTableFieldRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// r1R7q3QmWew5lo02fxB7noNyJxxxxxx
	BaseId *string `json:"BaseId,omitempty" xml:"BaseId,omitempty"`
	// This parameter is required.
	Name     *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	Property map[string]interface{} `json:"Property,omitempty" xml:"Property,omitempty"`
	// This parameter is required.
	SheetIdOrName *string                                       `json:"SheetIdOrName,omitempty" xml:"SheetIdOrName,omitempty"`
	TenantContext *CreateMultiDimTableFieldRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// text
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateMultiDimTableFieldRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMultiDimTableFieldRequest) GoString() string {
	return s.String()
}

func (s *CreateMultiDimTableFieldRequest) GetBaseId() *string {
	return s.BaseId
}

func (s *CreateMultiDimTableFieldRequest) GetName() *string {
	return s.Name
}

func (s *CreateMultiDimTableFieldRequest) GetProperty() map[string]interface{} {
	return s.Property
}

func (s *CreateMultiDimTableFieldRequest) GetSheetIdOrName() *string {
	return s.SheetIdOrName
}

func (s *CreateMultiDimTableFieldRequest) GetTenantContext() *CreateMultiDimTableFieldRequestTenantContext {
	return s.TenantContext
}

func (s *CreateMultiDimTableFieldRequest) GetType() *string {
	return s.Type
}

func (s *CreateMultiDimTableFieldRequest) SetBaseId(v string) *CreateMultiDimTableFieldRequest {
	s.BaseId = &v
	return s
}

func (s *CreateMultiDimTableFieldRequest) SetName(v string) *CreateMultiDimTableFieldRequest {
	s.Name = &v
	return s
}

func (s *CreateMultiDimTableFieldRequest) SetProperty(v map[string]interface{}) *CreateMultiDimTableFieldRequest {
	s.Property = v
	return s
}

func (s *CreateMultiDimTableFieldRequest) SetSheetIdOrName(v string) *CreateMultiDimTableFieldRequest {
	s.SheetIdOrName = &v
	return s
}

func (s *CreateMultiDimTableFieldRequest) SetTenantContext(v *CreateMultiDimTableFieldRequestTenantContext) *CreateMultiDimTableFieldRequest {
	s.TenantContext = v
	return s
}

func (s *CreateMultiDimTableFieldRequest) SetType(v string) *CreateMultiDimTableFieldRequest {
	s.Type = &v
	return s
}

func (s *CreateMultiDimTableFieldRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMultiDimTableFieldRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateMultiDimTableFieldRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s CreateMultiDimTableFieldRequestTenantContext) GoString() string {
	return s.String()
}

func (s *CreateMultiDimTableFieldRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateMultiDimTableFieldRequestTenantContext) SetTenantId(v string) *CreateMultiDimTableFieldRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *CreateMultiDimTableFieldRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
