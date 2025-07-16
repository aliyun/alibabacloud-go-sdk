// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMultiDimTableRecordsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseId(v string) *DeleteMultiDimTableRecordsShrinkRequest
	GetBaseId() *string
	SetRecordIdsShrink(v string) *DeleteMultiDimTableRecordsShrinkRequest
	GetRecordIdsShrink() *string
	SetSheetIdOrName(v string) *DeleteMultiDimTableRecordsShrinkRequest
	GetSheetIdOrName() *string
	SetTenantContextShrink(v string) *DeleteMultiDimTableRecordsShrinkRequest
	GetTenantContextShrink() *string
}

type DeleteMultiDimTableRecordsShrinkRequest struct {
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

func (s DeleteMultiDimTableRecordsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMultiDimTableRecordsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteMultiDimTableRecordsShrinkRequest) GetBaseId() *string {
	return s.BaseId
}

func (s *DeleteMultiDimTableRecordsShrinkRequest) GetRecordIdsShrink() *string {
	return s.RecordIdsShrink
}

func (s *DeleteMultiDimTableRecordsShrinkRequest) GetSheetIdOrName() *string {
	return s.SheetIdOrName
}

func (s *DeleteMultiDimTableRecordsShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *DeleteMultiDimTableRecordsShrinkRequest) SetBaseId(v string) *DeleteMultiDimTableRecordsShrinkRequest {
	s.BaseId = &v
	return s
}

func (s *DeleteMultiDimTableRecordsShrinkRequest) SetRecordIdsShrink(v string) *DeleteMultiDimTableRecordsShrinkRequest {
	s.RecordIdsShrink = &v
	return s
}

func (s *DeleteMultiDimTableRecordsShrinkRequest) SetSheetIdOrName(v string) *DeleteMultiDimTableRecordsShrinkRequest {
	s.SheetIdOrName = &v
	return s
}

func (s *DeleteMultiDimTableRecordsShrinkRequest) SetTenantContextShrink(v string) *DeleteMultiDimTableRecordsShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *DeleteMultiDimTableRecordsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
