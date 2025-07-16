// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMultiDimTableRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseId(v string) *DeleteMultiDimTableRecordsRequest
	GetBaseId() *string
	SetRecordIds(v []*string) *DeleteMultiDimTableRecordsRequest
	GetRecordIds() []*string
	SetSheetIdOrName(v string) *DeleteMultiDimTableRecordsRequest
	GetSheetIdOrName() *string
	SetTenantContext(v *DeleteMultiDimTableRecordsRequestTenantContext) *DeleteMultiDimTableRecordsRequest
	GetTenantContext() *DeleteMultiDimTableRecordsRequestTenantContext
}

type DeleteMultiDimTableRecordsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// r1R7q3QmWew5lo02fxB7nxxxxxxxx
	BaseId *string `json:"BaseId,omitempty" xml:"BaseId,omitempty"`
	// This parameter is required.
	RecordIds []*string `json:"RecordIds,omitempty" xml:"RecordIds,omitempty" type:"Repeated"`
	// This parameter is required.
	SheetIdOrName *string                                         `json:"SheetIdOrName,omitempty" xml:"SheetIdOrName,omitempty"`
	TenantContext *DeleteMultiDimTableRecordsRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s DeleteMultiDimTableRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMultiDimTableRecordsRequest) GoString() string {
	return s.String()
}

func (s *DeleteMultiDimTableRecordsRequest) GetBaseId() *string {
	return s.BaseId
}

func (s *DeleteMultiDimTableRecordsRequest) GetRecordIds() []*string {
	return s.RecordIds
}

func (s *DeleteMultiDimTableRecordsRequest) GetSheetIdOrName() *string {
	return s.SheetIdOrName
}

func (s *DeleteMultiDimTableRecordsRequest) GetTenantContext() *DeleteMultiDimTableRecordsRequestTenantContext {
	return s.TenantContext
}

func (s *DeleteMultiDimTableRecordsRequest) SetBaseId(v string) *DeleteMultiDimTableRecordsRequest {
	s.BaseId = &v
	return s
}

func (s *DeleteMultiDimTableRecordsRequest) SetRecordIds(v []*string) *DeleteMultiDimTableRecordsRequest {
	s.RecordIds = v
	return s
}

func (s *DeleteMultiDimTableRecordsRequest) SetSheetIdOrName(v string) *DeleteMultiDimTableRecordsRequest {
	s.SheetIdOrName = &v
	return s
}

func (s *DeleteMultiDimTableRecordsRequest) SetTenantContext(v *DeleteMultiDimTableRecordsRequestTenantContext) *DeleteMultiDimTableRecordsRequest {
	s.TenantContext = v
	return s
}

func (s *DeleteMultiDimTableRecordsRequest) Validate() error {
	return dara.Validate(s)
}

type DeleteMultiDimTableRecordsRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s DeleteMultiDimTableRecordsRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s DeleteMultiDimTableRecordsRequestTenantContext) GoString() string {
	return s.String()
}

func (s *DeleteMultiDimTableRecordsRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteMultiDimTableRecordsRequestTenantContext) SetTenantId(v string) *DeleteMultiDimTableRecordsRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *DeleteMultiDimTableRecordsRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
