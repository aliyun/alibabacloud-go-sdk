// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSheetContentJobIdShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDentryUuid(v string) *GetSheetContentJobIdShrinkRequest
	GetDentryUuid() *string
	SetExportType(v string) *GetSheetContentJobIdShrinkRequest
	GetExportType() *string
	SetTenantContextShrink(v string) *GetSheetContentJobIdShrinkRequest
	GetTenantContextShrink() *string
}

type GetSheetContentJobIdShrinkRequest struct {
	// example:
	//
	// MNDoBb60VLYDGNPytQe7Gzp4JlemrZQ3
	DentryUuid *string `json:"DentryUuid,omitempty" xml:"DentryUuid,omitempty"`
	// example:
	//
	// dingTalksheetToxlsx
	ExportType          *string `json:"ExportType,omitempty" xml:"ExportType,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s GetSheetContentJobIdShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSheetContentJobIdShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetSheetContentJobIdShrinkRequest) GetDentryUuid() *string {
	return s.DentryUuid
}

func (s *GetSheetContentJobIdShrinkRequest) GetExportType() *string {
	return s.ExportType
}

func (s *GetSheetContentJobIdShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetSheetContentJobIdShrinkRequest) SetDentryUuid(v string) *GetSheetContentJobIdShrinkRequest {
	s.DentryUuid = &v
	return s
}

func (s *GetSheetContentJobIdShrinkRequest) SetExportType(v string) *GetSheetContentJobIdShrinkRequest {
	s.ExportType = &v
	return s
}

func (s *GetSheetContentJobIdShrinkRequest) SetTenantContextShrink(v string) *GetSheetContentJobIdShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetSheetContentJobIdShrinkRequest) Validate() error {
	return dara.Validate(s)
}
