// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetRowsVisibilityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRow(v int64) *SetRowsVisibilityRequest
	GetRow() *int64
	SetRowCount(v int64) *SetRowsVisibilityRequest
	GetRowCount() *int64
	SetSheetId(v string) *SetRowsVisibilityRequest
	GetSheetId() *string
	SetTenantContext(v *SetRowsVisibilityRequestTenantContext) *SetRowsVisibilityRequest
	GetTenantContext() *SetRowsVisibilityRequestTenantContext
	SetVisibility(v string) *SetRowsVisibilityRequest
	GetVisibility() *string
	SetWorkbookId(v string) *SetRowsVisibilityRequest
	GetWorkbookId() *string
}

type SetRowsVisibilityRequest struct {
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
	// 20
	RowCount *int64 `json:"RowCount,omitempty" xml:"RowCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Sheet1
	SheetId       *string                                `json:"SheetId,omitempty" xml:"SheetId,omitempty"`
	TenantContext *SetRowsVisibilityRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// hidden
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// stxxxx
	WorkbookId *string `json:"WorkbookId,omitempty" xml:"WorkbookId,omitempty"`
}

func (s SetRowsVisibilityRequest) String() string {
	return dara.Prettify(s)
}

func (s SetRowsVisibilityRequest) GoString() string {
	return s.String()
}

func (s *SetRowsVisibilityRequest) GetRow() *int64 {
	return s.Row
}

func (s *SetRowsVisibilityRequest) GetRowCount() *int64 {
	return s.RowCount
}

func (s *SetRowsVisibilityRequest) GetSheetId() *string {
	return s.SheetId
}

func (s *SetRowsVisibilityRequest) GetTenantContext() *SetRowsVisibilityRequestTenantContext {
	return s.TenantContext
}

func (s *SetRowsVisibilityRequest) GetVisibility() *string {
	return s.Visibility
}

func (s *SetRowsVisibilityRequest) GetWorkbookId() *string {
	return s.WorkbookId
}

func (s *SetRowsVisibilityRequest) SetRow(v int64) *SetRowsVisibilityRequest {
	s.Row = &v
	return s
}

func (s *SetRowsVisibilityRequest) SetRowCount(v int64) *SetRowsVisibilityRequest {
	s.RowCount = &v
	return s
}

func (s *SetRowsVisibilityRequest) SetSheetId(v string) *SetRowsVisibilityRequest {
	s.SheetId = &v
	return s
}

func (s *SetRowsVisibilityRequest) SetTenantContext(v *SetRowsVisibilityRequestTenantContext) *SetRowsVisibilityRequest {
	s.TenantContext = v
	return s
}

func (s *SetRowsVisibilityRequest) SetVisibility(v string) *SetRowsVisibilityRequest {
	s.Visibility = &v
	return s
}

func (s *SetRowsVisibilityRequest) SetWorkbookId(v string) *SetRowsVisibilityRequest {
	s.WorkbookId = &v
	return s
}

func (s *SetRowsVisibilityRequest) Validate() error {
	return dara.Validate(s)
}

type SetRowsVisibilityRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s SetRowsVisibilityRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s SetRowsVisibilityRequestTenantContext) GoString() string {
	return s.String()
}

func (s *SetRowsVisibilityRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *SetRowsVisibilityRequestTenantContext) SetTenantId(v string) *SetRowsVisibilityRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *SetRowsVisibilityRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
