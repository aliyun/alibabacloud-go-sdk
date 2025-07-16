// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMultiDimTableRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseId(v string) *UpdateMultiDimTableRecordsRequest
	GetBaseId() *string
	SetRecordIds(v []*UpdateMultiDimTableRecordsRequestRecordIds) *UpdateMultiDimTableRecordsRequest
	GetRecordIds() []*UpdateMultiDimTableRecordsRequestRecordIds
	SetSheetIdOrName(v string) *UpdateMultiDimTableRecordsRequest
	GetSheetIdOrName() *string
	SetTenantContext(v *UpdateMultiDimTableRecordsRequestTenantContext) *UpdateMultiDimTableRecordsRequest
	GetTenantContext() *UpdateMultiDimTableRecordsRequestTenantContext
}

type UpdateMultiDimTableRecordsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// r1R7q3QmWew5lo02fxB7nxxxxxxxx
	BaseId *string `json:"BaseId,omitempty" xml:"BaseId,omitempty"`
	// This parameter is required.
	RecordIds []*UpdateMultiDimTableRecordsRequestRecordIds `json:"RecordIds,omitempty" xml:"RecordIds,omitempty" type:"Repeated"`
	// This parameter is required.
	SheetIdOrName *string                                         `json:"SheetIdOrName,omitempty" xml:"SheetIdOrName,omitempty"`
	TenantContext *UpdateMultiDimTableRecordsRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s UpdateMultiDimTableRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultiDimTableRecordsRequest) GoString() string {
	return s.String()
}

func (s *UpdateMultiDimTableRecordsRequest) GetBaseId() *string {
	return s.BaseId
}

func (s *UpdateMultiDimTableRecordsRequest) GetRecordIds() []*UpdateMultiDimTableRecordsRequestRecordIds {
	return s.RecordIds
}

func (s *UpdateMultiDimTableRecordsRequest) GetSheetIdOrName() *string {
	return s.SheetIdOrName
}

func (s *UpdateMultiDimTableRecordsRequest) GetTenantContext() *UpdateMultiDimTableRecordsRequestTenantContext {
	return s.TenantContext
}

func (s *UpdateMultiDimTableRecordsRequest) SetBaseId(v string) *UpdateMultiDimTableRecordsRequest {
	s.BaseId = &v
	return s
}

func (s *UpdateMultiDimTableRecordsRequest) SetRecordIds(v []*UpdateMultiDimTableRecordsRequestRecordIds) *UpdateMultiDimTableRecordsRequest {
	s.RecordIds = v
	return s
}

func (s *UpdateMultiDimTableRecordsRequest) SetSheetIdOrName(v string) *UpdateMultiDimTableRecordsRequest {
	s.SheetIdOrName = &v
	return s
}

func (s *UpdateMultiDimTableRecordsRequest) SetTenantContext(v *UpdateMultiDimTableRecordsRequestTenantContext) *UpdateMultiDimTableRecordsRequest {
	s.TenantContext = v
	return s
}

func (s *UpdateMultiDimTableRecordsRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateMultiDimTableRecordsRequestRecordIds struct {
	// This parameter is required.
	Fields map[string]interface{} `json:"Fields,omitempty" xml:"Fields,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HyDGtSj
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpdateMultiDimTableRecordsRequestRecordIds) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultiDimTableRecordsRequestRecordIds) GoString() string {
	return s.String()
}

func (s *UpdateMultiDimTableRecordsRequestRecordIds) GetFields() map[string]interface{} {
	return s.Fields
}

func (s *UpdateMultiDimTableRecordsRequestRecordIds) GetId() *string {
	return s.Id
}

func (s *UpdateMultiDimTableRecordsRequestRecordIds) SetFields(v map[string]interface{}) *UpdateMultiDimTableRecordsRequestRecordIds {
	s.Fields = v
	return s
}

func (s *UpdateMultiDimTableRecordsRequestRecordIds) SetId(v string) *UpdateMultiDimTableRecordsRequestRecordIds {
	s.Id = &v
	return s
}

func (s *UpdateMultiDimTableRecordsRequestRecordIds) Validate() error {
	return dara.Validate(s)
}

type UpdateMultiDimTableRecordsRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateMultiDimTableRecordsRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultiDimTableRecordsRequestTenantContext) GoString() string {
	return s.String()
}

func (s *UpdateMultiDimTableRecordsRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateMultiDimTableRecordsRequestTenantContext) SetTenantId(v string) *UpdateMultiDimTableRecordsRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *UpdateMultiDimTableRecordsRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
