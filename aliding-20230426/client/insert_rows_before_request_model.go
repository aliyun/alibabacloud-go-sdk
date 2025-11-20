// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertRowsBeforeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRow(v int64) *InsertRowsBeforeRequest
	GetRow() *int64
	SetRowCount(v int64) *InsertRowsBeforeRequest
	GetRowCount() *int64
	SetSheetId(v string) *InsertRowsBeforeRequest
	GetSheetId() *string
	SetTenantContext(v *InsertRowsBeforeRequestTenantContext) *InsertRowsBeforeRequest
	GetTenantContext() *InsertRowsBeforeRequestTenantContext
	SetWorkbookId(v string) *InsertRowsBeforeRequest
	GetWorkbookId() *string
}

type InsertRowsBeforeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2
	Row *int64 `json:"Row,omitempty" xml:"Row,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	RowCount *int64 `json:"RowCount,omitempty" xml:"RowCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Sheet1
	SheetId       *string                               `json:"SheetId,omitempty" xml:"SheetId,omitempty"`
	TenantContext *InsertRowsBeforeRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// stxxxx
	WorkbookId *string `json:"WorkbookId,omitempty" xml:"WorkbookId,omitempty"`
}

func (s InsertRowsBeforeRequest) String() string {
	return dara.Prettify(s)
}

func (s InsertRowsBeforeRequest) GoString() string {
	return s.String()
}

func (s *InsertRowsBeforeRequest) GetRow() *int64 {
	return s.Row
}

func (s *InsertRowsBeforeRequest) GetRowCount() *int64 {
	return s.RowCount
}

func (s *InsertRowsBeforeRequest) GetSheetId() *string {
	return s.SheetId
}

func (s *InsertRowsBeforeRequest) GetTenantContext() *InsertRowsBeforeRequestTenantContext {
	return s.TenantContext
}

func (s *InsertRowsBeforeRequest) GetWorkbookId() *string {
	return s.WorkbookId
}

func (s *InsertRowsBeforeRequest) SetRow(v int64) *InsertRowsBeforeRequest {
	s.Row = &v
	return s
}

func (s *InsertRowsBeforeRequest) SetRowCount(v int64) *InsertRowsBeforeRequest {
	s.RowCount = &v
	return s
}

func (s *InsertRowsBeforeRequest) SetSheetId(v string) *InsertRowsBeforeRequest {
	s.SheetId = &v
	return s
}

func (s *InsertRowsBeforeRequest) SetTenantContext(v *InsertRowsBeforeRequestTenantContext) *InsertRowsBeforeRequest {
	s.TenantContext = v
	return s
}

func (s *InsertRowsBeforeRequest) SetWorkbookId(v string) *InsertRowsBeforeRequest {
	s.WorkbookId = &v
	return s
}

func (s *InsertRowsBeforeRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InsertRowsBeforeRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s InsertRowsBeforeRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s InsertRowsBeforeRequestTenantContext) GoString() string {
	return s.String()
}

func (s *InsertRowsBeforeRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *InsertRowsBeforeRequestTenantContext) SetTenantId(v string) *InsertRowsBeforeRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *InsertRowsBeforeRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
