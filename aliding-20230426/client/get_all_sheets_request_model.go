// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAllSheetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContext(v *GetAllSheetsRequestTenantContext) *GetAllSheetsRequest
	GetTenantContext() *GetAllSheetsRequestTenantContext
	SetWorkbookId(v string) *GetAllSheetsRequest
	GetWorkbookId() *string
}

type GetAllSheetsRequest struct {
	TenantContext *GetAllSheetsRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// stxxxx
	WorkbookId *string `json:"WorkbookId,omitempty" xml:"WorkbookId,omitempty"`
}

func (s GetAllSheetsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAllSheetsRequest) GoString() string {
	return s.String()
}

func (s *GetAllSheetsRequest) GetTenantContext() *GetAllSheetsRequestTenantContext {
	return s.TenantContext
}

func (s *GetAllSheetsRequest) GetWorkbookId() *string {
	return s.WorkbookId
}

func (s *GetAllSheetsRequest) SetTenantContext(v *GetAllSheetsRequestTenantContext) *GetAllSheetsRequest {
	s.TenantContext = v
	return s
}

func (s *GetAllSheetsRequest) SetWorkbookId(v string) *GetAllSheetsRequest {
	s.WorkbookId = &v
	return s
}

func (s *GetAllSheetsRequest) Validate() error {
	return dara.Validate(s)
}

type GetAllSheetsRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetAllSheetsRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetAllSheetsRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetAllSheetsRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetAllSheetsRequestTenantContext) SetTenantId(v string) *GetAllSheetsRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetAllSheetsRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
