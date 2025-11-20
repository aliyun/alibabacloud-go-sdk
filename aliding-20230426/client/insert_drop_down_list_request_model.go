// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertDropDownListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOptions(v []*InsertDropDownListRequestOptions) *InsertDropDownListRequest
	GetOptions() []*InsertDropDownListRequestOptions
	SetRangeAddress(v string) *InsertDropDownListRequest
	GetRangeAddress() *string
	SetSheetId(v string) *InsertDropDownListRequest
	GetSheetId() *string
	SetTenantContext(v *InsertDropDownListRequestTenantContext) *InsertDropDownListRequest
	GetTenantContext() *InsertDropDownListRequestTenantContext
	SetWorkbookId(v string) *InsertDropDownListRequest
	GetWorkbookId() *string
}

type InsertDropDownListRequest struct {
	// This parameter is required.
	Options []*InsertDropDownListRequestOptions `json:"Options,omitempty" xml:"Options,omitempty" type:"Repeated"`
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
	SheetId       *string                                 `json:"SheetId,omitempty" xml:"SheetId,omitempty"`
	TenantContext *InsertDropDownListRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// stxxxx
	WorkbookId *string `json:"WorkbookId,omitempty" xml:"WorkbookId,omitempty"`
}

func (s InsertDropDownListRequest) String() string {
	return dara.Prettify(s)
}

func (s InsertDropDownListRequest) GoString() string {
	return s.String()
}

func (s *InsertDropDownListRequest) GetOptions() []*InsertDropDownListRequestOptions {
	return s.Options
}

func (s *InsertDropDownListRequest) GetRangeAddress() *string {
	return s.RangeAddress
}

func (s *InsertDropDownListRequest) GetSheetId() *string {
	return s.SheetId
}

func (s *InsertDropDownListRequest) GetTenantContext() *InsertDropDownListRequestTenantContext {
	return s.TenantContext
}

func (s *InsertDropDownListRequest) GetWorkbookId() *string {
	return s.WorkbookId
}

func (s *InsertDropDownListRequest) SetOptions(v []*InsertDropDownListRequestOptions) *InsertDropDownListRequest {
	s.Options = v
	return s
}

func (s *InsertDropDownListRequest) SetRangeAddress(v string) *InsertDropDownListRequest {
	s.RangeAddress = &v
	return s
}

func (s *InsertDropDownListRequest) SetSheetId(v string) *InsertDropDownListRequest {
	s.SheetId = &v
	return s
}

func (s *InsertDropDownListRequest) SetTenantContext(v *InsertDropDownListRequestTenantContext) *InsertDropDownListRequest {
	s.TenantContext = v
	return s
}

func (s *InsertDropDownListRequest) SetWorkbookId(v string) *InsertDropDownListRequest {
	s.WorkbookId = &v
	return s
}

func (s *InsertDropDownListRequest) Validate() error {
	if s.Options != nil {
		for _, item := range s.Options {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InsertDropDownListRequestOptions struct {
	// This parameter is required.
	//
	// example:
	//
	// #FF0000
	Color *string `json:"Color,omitempty" xml:"Color,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// A
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s InsertDropDownListRequestOptions) String() string {
	return dara.Prettify(s)
}

func (s InsertDropDownListRequestOptions) GoString() string {
	return s.String()
}

func (s *InsertDropDownListRequestOptions) GetColor() *string {
	return s.Color
}

func (s *InsertDropDownListRequestOptions) GetValue() *string {
	return s.Value
}

func (s *InsertDropDownListRequestOptions) SetColor(v string) *InsertDropDownListRequestOptions {
	s.Color = &v
	return s
}

func (s *InsertDropDownListRequestOptions) SetValue(v string) *InsertDropDownListRequestOptions {
	s.Value = &v
	return s
}

func (s *InsertDropDownListRequestOptions) Validate() error {
	return dara.Validate(s)
}

type InsertDropDownListRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s InsertDropDownListRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s InsertDropDownListRequestTenantContext) GoString() string {
	return s.String()
}

func (s *InsertDropDownListRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *InsertDropDownListRequestTenantContext) SetTenantId(v string) *InsertDropDownListRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *InsertDropDownListRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
