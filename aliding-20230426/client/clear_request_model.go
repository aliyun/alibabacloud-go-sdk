// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRangeAddress(v string) *ClearRequest
	GetRangeAddress() *string
	SetSheetId(v string) *ClearRequest
	GetSheetId() *string
	SetTenantContext(v *ClearRequestTenantContext) *ClearRequest
	GetTenantContext() *ClearRequestTenantContext
	SetWorkbookId(v string) *ClearRequest
	GetWorkbookId() *string
}

type ClearRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// A3:C3
	RangeAddress *string `json:"RangeAddress,omitempty" xml:"RangeAddress,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Sheet1
	SheetId       *string                    `json:"SheetId,omitempty" xml:"SheetId,omitempty"`
	TenantContext *ClearRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// stxxxx
	WorkbookId *string `json:"WorkbookId,omitempty" xml:"WorkbookId,omitempty"`
}

func (s ClearRequest) String() string {
	return dara.Prettify(s)
}

func (s ClearRequest) GoString() string {
	return s.String()
}

func (s *ClearRequest) GetRangeAddress() *string {
	return s.RangeAddress
}

func (s *ClearRequest) GetSheetId() *string {
	return s.SheetId
}

func (s *ClearRequest) GetTenantContext() *ClearRequestTenantContext {
	return s.TenantContext
}

func (s *ClearRequest) GetWorkbookId() *string {
	return s.WorkbookId
}

func (s *ClearRequest) SetRangeAddress(v string) *ClearRequest {
	s.RangeAddress = &v
	return s
}

func (s *ClearRequest) SetSheetId(v string) *ClearRequest {
	s.SheetId = &v
	return s
}

func (s *ClearRequest) SetTenantContext(v *ClearRequestTenantContext) *ClearRequest {
	s.TenantContext = v
	return s
}

func (s *ClearRequest) SetWorkbookId(v string) *ClearRequest {
	s.WorkbookId = &v
	return s
}

func (s *ClearRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClearRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ClearRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s ClearRequestTenantContext) GoString() string {
	return s.String()
}

func (s *ClearRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *ClearRequestTenantContext) SetTenantId(v string) *ClearRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *ClearRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
