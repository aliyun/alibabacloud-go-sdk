// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSheetContentJobIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDentryUuid(v string) *GetSheetContentJobIdRequest
	GetDentryUuid() *string
	SetExportType(v string) *GetSheetContentJobIdRequest
	GetExportType() *string
	SetTenantContext(v *GetSheetContentJobIdRequestTenantContext) *GetSheetContentJobIdRequest
	GetTenantContext() *GetSheetContentJobIdRequestTenantContext
}

type GetSheetContentJobIdRequest struct {
	// example:
	//
	// MNDoBb60VLYDGNPytQe7Gzp4JlemrZQ3
	DentryUuid *string `json:"DentryUuid,omitempty" xml:"DentryUuid,omitempty"`
	// example:
	//
	// dingTalksheetToxlsx
	ExportType    *string                                   `json:"ExportType,omitempty" xml:"ExportType,omitempty"`
	TenantContext *GetSheetContentJobIdRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s GetSheetContentJobIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSheetContentJobIdRequest) GoString() string {
	return s.String()
}

func (s *GetSheetContentJobIdRequest) GetDentryUuid() *string {
	return s.DentryUuid
}

func (s *GetSheetContentJobIdRequest) GetExportType() *string {
	return s.ExportType
}

func (s *GetSheetContentJobIdRequest) GetTenantContext() *GetSheetContentJobIdRequestTenantContext {
	return s.TenantContext
}

func (s *GetSheetContentJobIdRequest) SetDentryUuid(v string) *GetSheetContentJobIdRequest {
	s.DentryUuid = &v
	return s
}

func (s *GetSheetContentJobIdRequest) SetExportType(v string) *GetSheetContentJobIdRequest {
	s.ExportType = &v
	return s
}

func (s *GetSheetContentJobIdRequest) SetTenantContext(v *GetSheetContentJobIdRequestTenantContext) *GetSheetContentJobIdRequest {
	s.TenantContext = v
	return s
}

func (s *GetSheetContentJobIdRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSheetContentJobIdRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetSheetContentJobIdRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetSheetContentJobIdRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetSheetContentJobIdRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetSheetContentJobIdRequestTenantContext) SetTenantId(v string) *GetSheetContentJobIdRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetSheetContentJobIdRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
