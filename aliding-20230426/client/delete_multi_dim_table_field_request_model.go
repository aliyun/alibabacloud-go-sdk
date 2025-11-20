// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMultiDimTableFieldRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseId(v string) *DeleteMultiDimTableFieldRequest
	GetBaseId() *string
	SetFieldIdOrName(v string) *DeleteMultiDimTableFieldRequest
	GetFieldIdOrName() *string
	SetSheetIdOrName(v string) *DeleteMultiDimTableFieldRequest
	GetSheetIdOrName() *string
	SetTenantContext(v *DeleteMultiDimTableFieldRequestTenantContext) *DeleteMultiDimTableFieldRequest
	GetTenantContext() *DeleteMultiDimTableFieldRequestTenantContext
}

type DeleteMultiDimTableFieldRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// r1R7q3QmWew5lo02fxB7nxxxxxxxx
	BaseId *string `json:"BaseId,omitempty" xml:"BaseId,omitempty"`
	// This parameter is required.
	FieldIdOrName *string `json:"FieldIdOrName,omitempty" xml:"FieldIdOrName,omitempty"`
	// This parameter is required.
	SheetIdOrName *string                                       `json:"SheetIdOrName,omitempty" xml:"SheetIdOrName,omitempty"`
	TenantContext *DeleteMultiDimTableFieldRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s DeleteMultiDimTableFieldRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMultiDimTableFieldRequest) GoString() string {
	return s.String()
}

func (s *DeleteMultiDimTableFieldRequest) GetBaseId() *string {
	return s.BaseId
}

func (s *DeleteMultiDimTableFieldRequest) GetFieldIdOrName() *string {
	return s.FieldIdOrName
}

func (s *DeleteMultiDimTableFieldRequest) GetSheetIdOrName() *string {
	return s.SheetIdOrName
}

func (s *DeleteMultiDimTableFieldRequest) GetTenantContext() *DeleteMultiDimTableFieldRequestTenantContext {
	return s.TenantContext
}

func (s *DeleteMultiDimTableFieldRequest) SetBaseId(v string) *DeleteMultiDimTableFieldRequest {
	s.BaseId = &v
	return s
}

func (s *DeleteMultiDimTableFieldRequest) SetFieldIdOrName(v string) *DeleteMultiDimTableFieldRequest {
	s.FieldIdOrName = &v
	return s
}

func (s *DeleteMultiDimTableFieldRequest) SetSheetIdOrName(v string) *DeleteMultiDimTableFieldRequest {
	s.SheetIdOrName = &v
	return s
}

func (s *DeleteMultiDimTableFieldRequest) SetTenantContext(v *DeleteMultiDimTableFieldRequestTenantContext) *DeleteMultiDimTableFieldRequest {
	s.TenantContext = v
	return s
}

func (s *DeleteMultiDimTableFieldRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteMultiDimTableFieldRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s DeleteMultiDimTableFieldRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s DeleteMultiDimTableFieldRequestTenantContext) GoString() string {
	return s.String()
}

func (s *DeleteMultiDimTableFieldRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteMultiDimTableFieldRequestTenantContext) SetTenantId(v string) *DeleteMultiDimTableFieldRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *DeleteMultiDimTableFieldRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
