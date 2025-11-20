// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRangeAddress(v string) *ClearDataRequest
	GetRangeAddress() *string
	SetSheetId(v string) *ClearDataRequest
	GetSheetId() *string
	SetTenantContext(v *ClearDataRequestTenantContext) *ClearDataRequest
	GetTenantContext() *ClearDataRequestTenantContext
	SetWorkbookId(v string) *ClearDataRequest
	GetWorkbookId() *string
}

type ClearDataRequest struct {
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
	SheetId       *string                        `json:"SheetId,omitempty" xml:"SheetId,omitempty"`
	TenantContext *ClearDataRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// stxxxx
	WorkbookId *string `json:"WorkbookId,omitempty" xml:"WorkbookId,omitempty"`
}

func (s ClearDataRequest) String() string {
	return dara.Prettify(s)
}

func (s ClearDataRequest) GoString() string {
	return s.String()
}

func (s *ClearDataRequest) GetRangeAddress() *string {
	return s.RangeAddress
}

func (s *ClearDataRequest) GetSheetId() *string {
	return s.SheetId
}

func (s *ClearDataRequest) GetTenantContext() *ClearDataRequestTenantContext {
	return s.TenantContext
}

func (s *ClearDataRequest) GetWorkbookId() *string {
	return s.WorkbookId
}

func (s *ClearDataRequest) SetRangeAddress(v string) *ClearDataRequest {
	s.RangeAddress = &v
	return s
}

func (s *ClearDataRequest) SetSheetId(v string) *ClearDataRequest {
	s.SheetId = &v
	return s
}

func (s *ClearDataRequest) SetTenantContext(v *ClearDataRequestTenantContext) *ClearDataRequest {
	s.TenantContext = v
	return s
}

func (s *ClearDataRequest) SetWorkbookId(v string) *ClearDataRequest {
	s.WorkbookId = &v
	return s
}

func (s *ClearDataRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClearDataRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ClearDataRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s ClearDataRequestTenantContext) GoString() string {
	return s.String()
}

func (s *ClearDataRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *ClearDataRequestTenantContext) SetTenantId(v string) *ClearDataRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *ClearDataRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
