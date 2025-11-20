// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteColumnsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColumn(v int64) *DeleteColumnsRequest
	GetColumn() *int64
	SetColumnCount(v int64) *DeleteColumnsRequest
	GetColumnCount() *int64
	SetSheetId(v string) *DeleteColumnsRequest
	GetSheetId() *string
	SetTenantContext(v *DeleteColumnsRequestTenantContext) *DeleteColumnsRequest
	GetTenantContext() *DeleteColumnsRequestTenantContext
	SetWorkbookId(v string) *DeleteColumnsRequest
	GetWorkbookId() *string
}

type DeleteColumnsRequest struct {
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
	// 10
	ColumnCount *int64 `json:"ColumnCount,omitempty" xml:"ColumnCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Sheet1
	SheetId       *string                            `json:"SheetId,omitempty" xml:"SheetId,omitempty"`
	TenantContext *DeleteColumnsRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// stxxxx
	WorkbookId *string `json:"WorkbookId,omitempty" xml:"WorkbookId,omitempty"`
}

func (s DeleteColumnsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteColumnsRequest) GoString() string {
	return s.String()
}

func (s *DeleteColumnsRequest) GetColumn() *int64 {
	return s.Column
}

func (s *DeleteColumnsRequest) GetColumnCount() *int64 {
	return s.ColumnCount
}

func (s *DeleteColumnsRequest) GetSheetId() *string {
	return s.SheetId
}

func (s *DeleteColumnsRequest) GetTenantContext() *DeleteColumnsRequestTenantContext {
	return s.TenantContext
}

func (s *DeleteColumnsRequest) GetWorkbookId() *string {
	return s.WorkbookId
}

func (s *DeleteColumnsRequest) SetColumn(v int64) *DeleteColumnsRequest {
	s.Column = &v
	return s
}

func (s *DeleteColumnsRequest) SetColumnCount(v int64) *DeleteColumnsRequest {
	s.ColumnCount = &v
	return s
}

func (s *DeleteColumnsRequest) SetSheetId(v string) *DeleteColumnsRequest {
	s.SheetId = &v
	return s
}

func (s *DeleteColumnsRequest) SetTenantContext(v *DeleteColumnsRequestTenantContext) *DeleteColumnsRequest {
	s.TenantContext = v
	return s
}

func (s *DeleteColumnsRequest) SetWorkbookId(v string) *DeleteColumnsRequest {
	s.WorkbookId = &v
	return s
}

func (s *DeleteColumnsRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteColumnsRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s DeleteColumnsRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s DeleteColumnsRequestTenantContext) GoString() string {
	return s.String()
}

func (s *DeleteColumnsRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteColumnsRequestTenantContext) SetTenantId(v string) *DeleteColumnsRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *DeleteColumnsRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
