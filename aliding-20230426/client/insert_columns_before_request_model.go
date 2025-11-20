// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertColumnsBeforeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColumn(v int64) *InsertColumnsBeforeRequest
	GetColumn() *int64
	SetColumnCount(v int64) *InsertColumnsBeforeRequest
	GetColumnCount() *int64
	SetSheetId(v string) *InsertColumnsBeforeRequest
	GetSheetId() *string
	SetTenantContext(v *InsertColumnsBeforeRequestTenantContext) *InsertColumnsBeforeRequest
	GetTenantContext() *InsertColumnsBeforeRequestTenantContext
	SetWorkbookId(v string) *InsertColumnsBeforeRequest
	GetWorkbookId() *string
}

type InsertColumnsBeforeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	Column *int64 `json:"Column,omitempty" xml:"Column,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	ColumnCount *int64 `json:"ColumnCount,omitempty" xml:"ColumnCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Sheet1
	SheetId       *string                                  `json:"SheetId,omitempty" xml:"SheetId,omitempty"`
	TenantContext *InsertColumnsBeforeRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// stxxxx
	WorkbookId *string `json:"WorkbookId,omitempty" xml:"WorkbookId,omitempty"`
}

func (s InsertColumnsBeforeRequest) String() string {
	return dara.Prettify(s)
}

func (s InsertColumnsBeforeRequest) GoString() string {
	return s.String()
}

func (s *InsertColumnsBeforeRequest) GetColumn() *int64 {
	return s.Column
}

func (s *InsertColumnsBeforeRequest) GetColumnCount() *int64 {
	return s.ColumnCount
}

func (s *InsertColumnsBeforeRequest) GetSheetId() *string {
	return s.SheetId
}

func (s *InsertColumnsBeforeRequest) GetTenantContext() *InsertColumnsBeforeRequestTenantContext {
	return s.TenantContext
}

func (s *InsertColumnsBeforeRequest) GetWorkbookId() *string {
	return s.WorkbookId
}

func (s *InsertColumnsBeforeRequest) SetColumn(v int64) *InsertColumnsBeforeRequest {
	s.Column = &v
	return s
}

func (s *InsertColumnsBeforeRequest) SetColumnCount(v int64) *InsertColumnsBeforeRequest {
	s.ColumnCount = &v
	return s
}

func (s *InsertColumnsBeforeRequest) SetSheetId(v string) *InsertColumnsBeforeRequest {
	s.SheetId = &v
	return s
}

func (s *InsertColumnsBeforeRequest) SetTenantContext(v *InsertColumnsBeforeRequestTenantContext) *InsertColumnsBeforeRequest {
	s.TenantContext = v
	return s
}

func (s *InsertColumnsBeforeRequest) SetWorkbookId(v string) *InsertColumnsBeforeRequest {
	s.WorkbookId = &v
	return s
}

func (s *InsertColumnsBeforeRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InsertColumnsBeforeRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s InsertColumnsBeforeRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s InsertColumnsBeforeRequestTenantContext) GoString() string {
	return s.String()
}

func (s *InsertColumnsBeforeRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *InsertColumnsBeforeRequestTenantContext) SetTenantId(v string) *InsertColumnsBeforeRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *InsertColumnsBeforeRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
