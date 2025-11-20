// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertMultiDimTableRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseId(v string) *InsertMultiDimTableRecordRequest
	GetBaseId() *string
	SetRecords(v []*InsertMultiDimTableRecordRequestRecords) *InsertMultiDimTableRecordRequest
	GetRecords() []*InsertMultiDimTableRecordRequestRecords
	SetSheetIdOrName(v string) *InsertMultiDimTableRecordRequest
	GetSheetIdOrName() *string
	SetTenantContext(v *InsertMultiDimTableRecordRequestTenantContext) *InsertMultiDimTableRecordRequest
	GetTenantContext() *InsertMultiDimTableRecordRequestTenantContext
}

type InsertMultiDimTableRecordRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// r1R7q3QmWew5lo02fxB7nxxxxxxxx
	BaseId *string `json:"BaseId,omitempty" xml:"BaseId,omitempty"`
	// This parameter is required.
	Records []*InsertMultiDimTableRecordRequestRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// This parameter is required.
	SheetIdOrName *string                                        `json:"SheetIdOrName,omitempty" xml:"SheetIdOrName,omitempty"`
	TenantContext *InsertMultiDimTableRecordRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s InsertMultiDimTableRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s InsertMultiDimTableRecordRequest) GoString() string {
	return s.String()
}

func (s *InsertMultiDimTableRecordRequest) GetBaseId() *string {
	return s.BaseId
}

func (s *InsertMultiDimTableRecordRequest) GetRecords() []*InsertMultiDimTableRecordRequestRecords {
	return s.Records
}

func (s *InsertMultiDimTableRecordRequest) GetSheetIdOrName() *string {
	return s.SheetIdOrName
}

func (s *InsertMultiDimTableRecordRequest) GetTenantContext() *InsertMultiDimTableRecordRequestTenantContext {
	return s.TenantContext
}

func (s *InsertMultiDimTableRecordRequest) SetBaseId(v string) *InsertMultiDimTableRecordRequest {
	s.BaseId = &v
	return s
}

func (s *InsertMultiDimTableRecordRequest) SetRecords(v []*InsertMultiDimTableRecordRequestRecords) *InsertMultiDimTableRecordRequest {
	s.Records = v
	return s
}

func (s *InsertMultiDimTableRecordRequest) SetSheetIdOrName(v string) *InsertMultiDimTableRecordRequest {
	s.SheetIdOrName = &v
	return s
}

func (s *InsertMultiDimTableRecordRequest) SetTenantContext(v *InsertMultiDimTableRecordRequestTenantContext) *InsertMultiDimTableRecordRequest {
	s.TenantContext = v
	return s
}

func (s *InsertMultiDimTableRecordRequest) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InsertMultiDimTableRecordRequestRecords struct {
	// This parameter is required.
	Fields map[string]interface{} `json:"Fields,omitempty" xml:"Fields,omitempty"`
}

func (s InsertMultiDimTableRecordRequestRecords) String() string {
	return dara.Prettify(s)
}

func (s InsertMultiDimTableRecordRequestRecords) GoString() string {
	return s.String()
}

func (s *InsertMultiDimTableRecordRequestRecords) GetFields() map[string]interface{} {
	return s.Fields
}

func (s *InsertMultiDimTableRecordRequestRecords) SetFields(v map[string]interface{}) *InsertMultiDimTableRecordRequestRecords {
	s.Fields = v
	return s
}

func (s *InsertMultiDimTableRecordRequestRecords) Validate() error {
	return dara.Validate(s)
}

type InsertMultiDimTableRecordRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s InsertMultiDimTableRecordRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s InsertMultiDimTableRecordRequestTenantContext) GoString() string {
	return s.String()
}

func (s *InsertMultiDimTableRecordRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *InsertMultiDimTableRecordRequestTenantContext) SetTenantId(v string) *InsertMultiDimTableRecordRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *InsertMultiDimTableRecordRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
