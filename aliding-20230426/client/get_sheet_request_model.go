// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSheetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSheetId(v string) *GetSheetRequest
	GetSheetId() *string
	SetTenantContext(v *GetSheetRequestTenantContext) *GetSheetRequest
	GetTenantContext() *GetSheetRequestTenantContext
	SetWorkbookId(v string) *GetSheetRequest
	GetWorkbookId() *string
}

type GetSheetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Sheet1
	SheetId       *string                       `json:"SheetId,omitempty" xml:"SheetId,omitempty"`
	TenantContext *GetSheetRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// stxxxx
	WorkbookId *string `json:"WorkbookId,omitempty" xml:"WorkbookId,omitempty"`
}

func (s GetSheetRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSheetRequest) GoString() string {
	return s.String()
}

func (s *GetSheetRequest) GetSheetId() *string {
	return s.SheetId
}

func (s *GetSheetRequest) GetTenantContext() *GetSheetRequestTenantContext {
	return s.TenantContext
}

func (s *GetSheetRequest) GetWorkbookId() *string {
	return s.WorkbookId
}

func (s *GetSheetRequest) SetSheetId(v string) *GetSheetRequest {
	s.SheetId = &v
	return s
}

func (s *GetSheetRequest) SetTenantContext(v *GetSheetRequestTenantContext) *GetSheetRequest {
	s.TenantContext = v
	return s
}

func (s *GetSheetRequest) SetWorkbookId(v string) *GetSheetRequest {
	s.WorkbookId = &v
	return s
}

func (s *GetSheetRequest) Validate() error {
	return dara.Validate(s)
}

type GetSheetRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetSheetRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetSheetRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetSheetRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetSheetRequestTenantContext) SetTenantId(v string) *GetSheetRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetSheetRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
