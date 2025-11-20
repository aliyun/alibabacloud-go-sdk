// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiDimTableRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseId(v string) *GetMultiDimTableRecordRequest
	GetBaseId() *string
	SetRecordId(v string) *GetMultiDimTableRecordRequest
	GetRecordId() *string
	SetSheetIdOrName(v string) *GetMultiDimTableRecordRequest
	GetSheetIdOrName() *string
	SetTenantContext(v *GetMultiDimTableRecordRequestTenantContext) *GetMultiDimTableRecordRequest
	GetTenantContext() *GetMultiDimTableRecordRequestTenantContext
}

type GetMultiDimTableRecordRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 101114
	BaseId *string `json:"BaseId,omitempty" xml:"BaseId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxx
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxx
	SheetIdOrName *string                                     `json:"SheetIdOrName,omitempty" xml:"SheetIdOrName,omitempty"`
	TenantContext *GetMultiDimTableRecordRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s GetMultiDimTableRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMultiDimTableRecordRequest) GoString() string {
	return s.String()
}

func (s *GetMultiDimTableRecordRequest) GetBaseId() *string {
	return s.BaseId
}

func (s *GetMultiDimTableRecordRequest) GetRecordId() *string {
	return s.RecordId
}

func (s *GetMultiDimTableRecordRequest) GetSheetIdOrName() *string {
	return s.SheetIdOrName
}

func (s *GetMultiDimTableRecordRequest) GetTenantContext() *GetMultiDimTableRecordRequestTenantContext {
	return s.TenantContext
}

func (s *GetMultiDimTableRecordRequest) SetBaseId(v string) *GetMultiDimTableRecordRequest {
	s.BaseId = &v
	return s
}

func (s *GetMultiDimTableRecordRequest) SetRecordId(v string) *GetMultiDimTableRecordRequest {
	s.RecordId = &v
	return s
}

func (s *GetMultiDimTableRecordRequest) SetSheetIdOrName(v string) *GetMultiDimTableRecordRequest {
	s.SheetIdOrName = &v
	return s
}

func (s *GetMultiDimTableRecordRequest) SetTenantContext(v *GetMultiDimTableRecordRequestTenantContext) *GetMultiDimTableRecordRequest {
	s.TenantContext = v
	return s
}

func (s *GetMultiDimTableRecordRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMultiDimTableRecordRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetMultiDimTableRecordRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetMultiDimTableRecordRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetMultiDimTableRecordRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetMultiDimTableRecordRequestTenantContext) SetTenantId(v string) *GetMultiDimTableRecordRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetMultiDimTableRecordRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
