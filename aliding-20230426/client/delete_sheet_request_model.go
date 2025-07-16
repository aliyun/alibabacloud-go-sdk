// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSheetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSheetId(v string) *DeleteSheetRequest
	GetSheetId() *string
	SetTenantContext(v *DeleteSheetRequestTenantContext) *DeleteSheetRequest
	GetTenantContext() *DeleteSheetRequestTenantContext
	SetWorkbookId(v string) *DeleteSheetRequest
	GetWorkbookId() *string
}

type DeleteSheetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Sheet1
	SheetId       *string                          `json:"SheetId,omitempty" xml:"SheetId,omitempty"`
	TenantContext *DeleteSheetRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// stxxxx
	WorkbookId *string `json:"WorkbookId,omitempty" xml:"WorkbookId,omitempty"`
}

func (s DeleteSheetRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSheetRequest) GoString() string {
	return s.String()
}

func (s *DeleteSheetRequest) GetSheetId() *string {
	return s.SheetId
}

func (s *DeleteSheetRequest) GetTenantContext() *DeleteSheetRequestTenantContext {
	return s.TenantContext
}

func (s *DeleteSheetRequest) GetWorkbookId() *string {
	return s.WorkbookId
}

func (s *DeleteSheetRequest) SetSheetId(v string) *DeleteSheetRequest {
	s.SheetId = &v
	return s
}

func (s *DeleteSheetRequest) SetTenantContext(v *DeleteSheetRequestTenantContext) *DeleteSheetRequest {
	s.TenantContext = v
	return s
}

func (s *DeleteSheetRequest) SetWorkbookId(v string) *DeleteSheetRequest {
	s.WorkbookId = &v
	return s
}

func (s *DeleteSheetRequest) Validate() error {
	return dara.Validate(s)
}

type DeleteSheetRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s DeleteSheetRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s DeleteSheetRequestTenantContext) GoString() string {
	return s.String()
}

func (s *DeleteSheetRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteSheetRequestTenantContext) SetTenantId(v string) *DeleteSheetRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *DeleteSheetRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
