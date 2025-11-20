// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRangeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRangeAddress(v string) *GetRangeRequest
	GetRangeAddress() *string
	SetSelect(v string) *GetRangeRequest
	GetSelect() *string
	SetSheetId(v string) *GetRangeRequest
	GetSheetId() *string
	SetTenantContext(v *GetRangeRequestTenantContext) *GetRangeRequest
	GetTenantContext() *GetRangeRequestTenantContext
	SetWorkbookId(v string) *GetRangeRequest
	GetWorkbookId() *string
}

type GetRangeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// A3:C3
	RangeAddress *string `json:"RangeAddress,omitempty" xml:"RangeAddress,omitempty"`
	// example:
	//
	// values
	Select *string `json:"Select,omitempty" xml:"Select,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Sheet1
	SheetId       *string                       `json:"SheetId,omitempty" xml:"SheetId,omitempty"`
	TenantContext *GetRangeRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// stxxxx
	WorkbookId *string `json:"WorkbookId,omitempty" xml:"WorkbookId,omitempty"`
}

func (s GetRangeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRangeRequest) GoString() string {
	return s.String()
}

func (s *GetRangeRequest) GetRangeAddress() *string {
	return s.RangeAddress
}

func (s *GetRangeRequest) GetSelect() *string {
	return s.Select
}

func (s *GetRangeRequest) GetSheetId() *string {
	return s.SheetId
}

func (s *GetRangeRequest) GetTenantContext() *GetRangeRequestTenantContext {
	return s.TenantContext
}

func (s *GetRangeRequest) GetWorkbookId() *string {
	return s.WorkbookId
}

func (s *GetRangeRequest) SetRangeAddress(v string) *GetRangeRequest {
	s.RangeAddress = &v
	return s
}

func (s *GetRangeRequest) SetSelect(v string) *GetRangeRequest {
	s.Select = &v
	return s
}

func (s *GetRangeRequest) SetSheetId(v string) *GetRangeRequest {
	s.SheetId = &v
	return s
}

func (s *GetRangeRequest) SetTenantContext(v *GetRangeRequestTenantContext) *GetRangeRequest {
	s.TenantContext = v
	return s
}

func (s *GetRangeRequest) SetWorkbookId(v string) *GetRangeRequest {
	s.WorkbookId = &v
	return s
}

func (s *GetRangeRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRangeRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetRangeRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetRangeRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetRangeRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetRangeRequestTenantContext) SetTenantId(v string) *GetRangeRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetRangeRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
