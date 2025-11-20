// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRowsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRow(v int64) *DeleteRowsRequest
	GetRow() *int64
	SetRowCount(v int64) *DeleteRowsRequest
	GetRowCount() *int64
	SetSheetId(v string) *DeleteRowsRequest
	GetSheetId() *string
	SetTenantContext(v *DeleteRowsRequestTenantContext) *DeleteRowsRequest
	GetTenantContext() *DeleteRowsRequestTenantContext
	SetWorkbookId(v string) *DeleteRowsRequest
	GetWorkbookId() *string
}

type DeleteRowsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
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
	SheetId       *string                         `json:"SheetId,omitempty" xml:"SheetId,omitempty"`
	TenantContext *DeleteRowsRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// stxxxx
	WorkbookId *string `json:"WorkbookId,omitempty" xml:"WorkbookId,omitempty"`
}

func (s DeleteRowsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRowsRequest) GoString() string {
	return s.String()
}

func (s *DeleteRowsRequest) GetRow() *int64 {
	return s.Row
}

func (s *DeleteRowsRequest) GetRowCount() *int64 {
	return s.RowCount
}

func (s *DeleteRowsRequest) GetSheetId() *string {
	return s.SheetId
}

func (s *DeleteRowsRequest) GetTenantContext() *DeleteRowsRequestTenantContext {
	return s.TenantContext
}

func (s *DeleteRowsRequest) GetWorkbookId() *string {
	return s.WorkbookId
}

func (s *DeleteRowsRequest) SetRow(v int64) *DeleteRowsRequest {
	s.Row = &v
	return s
}

func (s *DeleteRowsRequest) SetRowCount(v int64) *DeleteRowsRequest {
	s.RowCount = &v
	return s
}

func (s *DeleteRowsRequest) SetSheetId(v string) *DeleteRowsRequest {
	s.SheetId = &v
	return s
}

func (s *DeleteRowsRequest) SetTenantContext(v *DeleteRowsRequestTenantContext) *DeleteRowsRequest {
	s.TenantContext = v
	return s
}

func (s *DeleteRowsRequest) SetWorkbookId(v string) *DeleteRowsRequest {
	s.WorkbookId = &v
	return s
}

func (s *DeleteRowsRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteRowsRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s DeleteRowsRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s DeleteRowsRequestTenantContext) GoString() string {
	return s.String()
}

func (s *DeleteRowsRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteRowsRequestTenantContext) SetTenantId(v string) *DeleteRowsRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *DeleteRowsRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
