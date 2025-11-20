// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiDimTableSheetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseId(v string) *GetMultiDimTableSheetRequest
	GetBaseId() *string
	SetSheetIdOrName(v string) *GetMultiDimTableSheetRequest
	GetSheetIdOrName() *string
	SetTenantContext(v *GetMultiDimTableSheetRequestTenantContext) *GetMultiDimTableSheetRequest
	GetTenantContext() *GetMultiDimTableSheetRequestTenantContext
}

type GetMultiDimTableSheetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 169899
	BaseId *string `json:"BaseId,omitempty" xml:"BaseId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxx
	SheetIdOrName *string                                    `json:"SheetIdOrName,omitempty" xml:"SheetIdOrName,omitempty"`
	TenantContext *GetMultiDimTableSheetRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s GetMultiDimTableSheetRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMultiDimTableSheetRequest) GoString() string {
	return s.String()
}

func (s *GetMultiDimTableSheetRequest) GetBaseId() *string {
	return s.BaseId
}

func (s *GetMultiDimTableSheetRequest) GetSheetIdOrName() *string {
	return s.SheetIdOrName
}

func (s *GetMultiDimTableSheetRequest) GetTenantContext() *GetMultiDimTableSheetRequestTenantContext {
	return s.TenantContext
}

func (s *GetMultiDimTableSheetRequest) SetBaseId(v string) *GetMultiDimTableSheetRequest {
	s.BaseId = &v
	return s
}

func (s *GetMultiDimTableSheetRequest) SetSheetIdOrName(v string) *GetMultiDimTableSheetRequest {
	s.SheetIdOrName = &v
	return s
}

func (s *GetMultiDimTableSheetRequest) SetTenantContext(v *GetMultiDimTableSheetRequestTenantContext) *GetMultiDimTableSheetRequest {
	s.TenantContext = v
	return s
}

func (s *GetMultiDimTableSheetRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMultiDimTableSheetRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetMultiDimTableSheetRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetMultiDimTableSheetRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetMultiDimTableSheetRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetMultiDimTableSheetRequestTenantContext) SetTenantId(v string) *GetMultiDimTableSheetRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetMultiDimTableSheetRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
