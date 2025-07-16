// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMultiDimTableRecordsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseId(v string) *UpdateMultiDimTableRecordsShrinkRequest
	GetBaseId() *string
	SetRecordIdsShrink(v string) *UpdateMultiDimTableRecordsShrinkRequest
	GetRecordIdsShrink() *string
	SetSheetIdOrName(v string) *UpdateMultiDimTableRecordsShrinkRequest
	GetSheetIdOrName() *string
	SetTenantContextShrink(v string) *UpdateMultiDimTableRecordsShrinkRequest
	GetTenantContextShrink() *string
}

type UpdateMultiDimTableRecordsShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// r1R7q3QmWew5lo02fxB7nxxxxxxxx
	BaseId *string `json:"BaseId,omitempty" xml:"BaseId,omitempty"`
	// This parameter is required.
	RecordIdsShrink *string `json:"RecordIds,omitempty" xml:"RecordIds,omitempty"`
	// This parameter is required.
	SheetIdOrName       *string `json:"SheetIdOrName,omitempty" xml:"SheetIdOrName,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s UpdateMultiDimTableRecordsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultiDimTableRecordsShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateMultiDimTableRecordsShrinkRequest) GetBaseId() *string {
	return s.BaseId
}

func (s *UpdateMultiDimTableRecordsShrinkRequest) GetRecordIdsShrink() *string {
	return s.RecordIdsShrink
}

func (s *UpdateMultiDimTableRecordsShrinkRequest) GetSheetIdOrName() *string {
	return s.SheetIdOrName
}

func (s *UpdateMultiDimTableRecordsShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *UpdateMultiDimTableRecordsShrinkRequest) SetBaseId(v string) *UpdateMultiDimTableRecordsShrinkRequest {
	s.BaseId = &v
	return s
}

func (s *UpdateMultiDimTableRecordsShrinkRequest) SetRecordIdsShrink(v string) *UpdateMultiDimTableRecordsShrinkRequest {
	s.RecordIdsShrink = &v
	return s
}

func (s *UpdateMultiDimTableRecordsShrinkRequest) SetSheetIdOrName(v string) *UpdateMultiDimTableRecordsShrinkRequest {
	s.SheetIdOrName = &v
	return s
}

func (s *UpdateMultiDimTableRecordsShrinkRequest) SetTenantContextShrink(v string) *UpdateMultiDimTableRecordsShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *UpdateMultiDimTableRecordsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
