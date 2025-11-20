// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetColumnsVisibilityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColumn(v int64) *SetColumnsVisibilityRequest
	GetColumn() *int64
	SetColumnCount(v int64) *SetColumnsVisibilityRequest
	GetColumnCount() *int64
	SetSheetId(v string) *SetColumnsVisibilityRequest
	GetSheetId() *string
	SetTenantContext(v *SetColumnsVisibilityRequestTenantContext) *SetColumnsVisibilityRequest
	GetTenantContext() *SetColumnsVisibilityRequestTenantContext
	SetVisibility(v string) *SetColumnsVisibilityRequest
	GetVisibility() *string
	SetWorkbookId(v string) *SetColumnsVisibilityRequest
	GetWorkbookId() *string
}

type SetColumnsVisibilityRequest struct {
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
	// 20
	ColumnCount *int64 `json:"ColumnCount,omitempty" xml:"ColumnCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Sheet1
	SheetId       *string                                   `json:"SheetId,omitempty" xml:"SheetId,omitempty"`
	TenantContext *SetColumnsVisibilityRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
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

func (s SetColumnsVisibilityRequest) String() string {
	return dara.Prettify(s)
}

func (s SetColumnsVisibilityRequest) GoString() string {
	return s.String()
}

func (s *SetColumnsVisibilityRequest) GetColumn() *int64 {
	return s.Column
}

func (s *SetColumnsVisibilityRequest) GetColumnCount() *int64 {
	return s.ColumnCount
}

func (s *SetColumnsVisibilityRequest) GetSheetId() *string {
	return s.SheetId
}

func (s *SetColumnsVisibilityRequest) GetTenantContext() *SetColumnsVisibilityRequestTenantContext {
	return s.TenantContext
}

func (s *SetColumnsVisibilityRequest) GetVisibility() *string {
	return s.Visibility
}

func (s *SetColumnsVisibilityRequest) GetWorkbookId() *string {
	return s.WorkbookId
}

func (s *SetColumnsVisibilityRequest) SetColumn(v int64) *SetColumnsVisibilityRequest {
	s.Column = &v
	return s
}

func (s *SetColumnsVisibilityRequest) SetColumnCount(v int64) *SetColumnsVisibilityRequest {
	s.ColumnCount = &v
	return s
}

func (s *SetColumnsVisibilityRequest) SetSheetId(v string) *SetColumnsVisibilityRequest {
	s.SheetId = &v
	return s
}

func (s *SetColumnsVisibilityRequest) SetTenantContext(v *SetColumnsVisibilityRequestTenantContext) *SetColumnsVisibilityRequest {
	s.TenantContext = v
	return s
}

func (s *SetColumnsVisibilityRequest) SetVisibility(v string) *SetColumnsVisibilityRequest {
	s.Visibility = &v
	return s
}

func (s *SetColumnsVisibilityRequest) SetWorkbookId(v string) *SetColumnsVisibilityRequest {
	s.WorkbookId = &v
	return s
}

func (s *SetColumnsVisibilityRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SetColumnsVisibilityRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s SetColumnsVisibilityRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s SetColumnsVisibilityRequestTenantContext) GoString() string {
	return s.String()
}

func (s *SetColumnsVisibilityRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *SetColumnsVisibilityRequestTenantContext) SetTenantId(v string) *SetColumnsVisibilityRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *SetColumnsVisibilityRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
