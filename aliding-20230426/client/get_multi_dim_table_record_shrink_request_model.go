// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiDimTableRecordShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseId(v string) *GetMultiDimTableRecordShrinkRequest
	GetBaseId() *string
	SetRecordId(v string) *GetMultiDimTableRecordShrinkRequest
	GetRecordId() *string
	SetSheetIdOrName(v string) *GetMultiDimTableRecordShrinkRequest
	GetSheetIdOrName() *string
	SetTenantContextShrink(v string) *GetMultiDimTableRecordShrinkRequest
	GetTenantContextShrink() *string
}

type GetMultiDimTableRecordShrinkRequest struct {
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
	SheetIdOrName       *string `json:"SheetIdOrName,omitempty" xml:"SheetIdOrName,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s GetMultiDimTableRecordShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMultiDimTableRecordShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetMultiDimTableRecordShrinkRequest) GetBaseId() *string {
	return s.BaseId
}

func (s *GetMultiDimTableRecordShrinkRequest) GetRecordId() *string {
	return s.RecordId
}

func (s *GetMultiDimTableRecordShrinkRequest) GetSheetIdOrName() *string {
	return s.SheetIdOrName
}

func (s *GetMultiDimTableRecordShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetMultiDimTableRecordShrinkRequest) SetBaseId(v string) *GetMultiDimTableRecordShrinkRequest {
	s.BaseId = &v
	return s
}

func (s *GetMultiDimTableRecordShrinkRequest) SetRecordId(v string) *GetMultiDimTableRecordShrinkRequest {
	s.RecordId = &v
	return s
}

func (s *GetMultiDimTableRecordShrinkRequest) SetSheetIdOrName(v string) *GetMultiDimTableRecordShrinkRequest {
	s.SheetIdOrName = &v
	return s
}

func (s *GetMultiDimTableRecordShrinkRequest) SetTenantContextShrink(v string) *GetMultiDimTableRecordShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetMultiDimTableRecordShrinkRequest) Validate() error {
	return dara.Validate(s)
}
