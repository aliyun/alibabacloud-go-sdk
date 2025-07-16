// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiDimTableAllFieldsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseId(v string) *GetMultiDimTableAllFieldsRequest
	GetBaseId() *string
	SetSheetIdOrName(v string) *GetMultiDimTableAllFieldsRequest
	GetSheetIdOrName() *string
	SetTenantContext(v *GetMultiDimTableAllFieldsRequestTenantContext) *GetMultiDimTableAllFieldsRequest
	GetTenantContext() *GetMultiDimTableAllFieldsRequestTenantContext
}

type GetMultiDimTableAllFieldsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 338534
	BaseId *string `json:"BaseId,omitempty" xml:"BaseId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxx
	SheetIdOrName *string                                        `json:"SheetIdOrName,omitempty" xml:"SheetIdOrName,omitempty"`
	TenantContext *GetMultiDimTableAllFieldsRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s GetMultiDimTableAllFieldsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMultiDimTableAllFieldsRequest) GoString() string {
	return s.String()
}

func (s *GetMultiDimTableAllFieldsRequest) GetBaseId() *string {
	return s.BaseId
}

func (s *GetMultiDimTableAllFieldsRequest) GetSheetIdOrName() *string {
	return s.SheetIdOrName
}

func (s *GetMultiDimTableAllFieldsRequest) GetTenantContext() *GetMultiDimTableAllFieldsRequestTenantContext {
	return s.TenantContext
}

func (s *GetMultiDimTableAllFieldsRequest) SetBaseId(v string) *GetMultiDimTableAllFieldsRequest {
	s.BaseId = &v
	return s
}

func (s *GetMultiDimTableAllFieldsRequest) SetSheetIdOrName(v string) *GetMultiDimTableAllFieldsRequest {
	s.SheetIdOrName = &v
	return s
}

func (s *GetMultiDimTableAllFieldsRequest) SetTenantContext(v *GetMultiDimTableAllFieldsRequestTenantContext) *GetMultiDimTableAllFieldsRequest {
	s.TenantContext = v
	return s
}

func (s *GetMultiDimTableAllFieldsRequest) Validate() error {
	return dara.Validate(s)
}

type GetMultiDimTableAllFieldsRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetMultiDimTableAllFieldsRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetMultiDimTableAllFieldsRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetMultiDimTableAllFieldsRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetMultiDimTableAllFieldsRequestTenantContext) SetTenantId(v string) *GetMultiDimTableAllFieldsRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetMultiDimTableAllFieldsRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
