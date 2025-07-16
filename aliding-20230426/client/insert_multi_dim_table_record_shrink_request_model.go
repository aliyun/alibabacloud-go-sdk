// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertMultiDimTableRecordShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseId(v string) *InsertMultiDimTableRecordShrinkRequest
	GetBaseId() *string
	SetRecordsShrink(v string) *InsertMultiDimTableRecordShrinkRequest
	GetRecordsShrink() *string
	SetSheetIdOrName(v string) *InsertMultiDimTableRecordShrinkRequest
	GetSheetIdOrName() *string
	SetTenantContextShrink(v string) *InsertMultiDimTableRecordShrinkRequest
	GetTenantContextShrink() *string
}

type InsertMultiDimTableRecordShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// r1R7q3QmWew5lo02fxB7nxxxxxxxx
	BaseId *string `json:"BaseId,omitempty" xml:"BaseId,omitempty"`
	// This parameter is required.
	RecordsShrink *string `json:"Records,omitempty" xml:"Records,omitempty"`
	// This parameter is required.
	SheetIdOrName       *string `json:"SheetIdOrName,omitempty" xml:"SheetIdOrName,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s InsertMultiDimTableRecordShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s InsertMultiDimTableRecordShrinkRequest) GoString() string {
	return s.String()
}

func (s *InsertMultiDimTableRecordShrinkRequest) GetBaseId() *string {
	return s.BaseId
}

func (s *InsertMultiDimTableRecordShrinkRequest) GetRecordsShrink() *string {
	return s.RecordsShrink
}

func (s *InsertMultiDimTableRecordShrinkRequest) GetSheetIdOrName() *string {
	return s.SheetIdOrName
}

func (s *InsertMultiDimTableRecordShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *InsertMultiDimTableRecordShrinkRequest) SetBaseId(v string) *InsertMultiDimTableRecordShrinkRequest {
	s.BaseId = &v
	return s
}

func (s *InsertMultiDimTableRecordShrinkRequest) SetRecordsShrink(v string) *InsertMultiDimTableRecordShrinkRequest {
	s.RecordsShrink = &v
	return s
}

func (s *InsertMultiDimTableRecordShrinkRequest) SetSheetIdOrName(v string) *InsertMultiDimTableRecordShrinkRequest {
	s.SheetIdOrName = &v
	return s
}

func (s *InsertMultiDimTableRecordShrinkRequest) SetTenantContextShrink(v string) *InsertMultiDimTableRecordShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *InsertMultiDimTableRecordShrinkRequest) Validate() error {
	return dara.Validate(s)
}
