// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiDimTableAllSheetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseId(v string) *GetMultiDimTableAllSheetsRequest
	GetBaseId() *string
	SetTenantContext(v *GetMultiDimTableAllSheetsRequestTenantContext) *GetMultiDimTableAllSheetsRequest
	GetTenantContext() *GetMultiDimTableAllSheetsRequestTenantContext
}

type GetMultiDimTableAllSheetsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 144972
	BaseId        *string                                        `json:"BaseId,omitempty" xml:"BaseId,omitempty"`
	TenantContext *GetMultiDimTableAllSheetsRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s GetMultiDimTableAllSheetsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMultiDimTableAllSheetsRequest) GoString() string {
	return s.String()
}

func (s *GetMultiDimTableAllSheetsRequest) GetBaseId() *string {
	return s.BaseId
}

func (s *GetMultiDimTableAllSheetsRequest) GetTenantContext() *GetMultiDimTableAllSheetsRequestTenantContext {
	return s.TenantContext
}

func (s *GetMultiDimTableAllSheetsRequest) SetBaseId(v string) *GetMultiDimTableAllSheetsRequest {
	s.BaseId = &v
	return s
}

func (s *GetMultiDimTableAllSheetsRequest) SetTenantContext(v *GetMultiDimTableAllSheetsRequestTenantContext) *GetMultiDimTableAllSheetsRequest {
	s.TenantContext = v
	return s
}

func (s *GetMultiDimTableAllSheetsRequest) Validate() error {
	return dara.Validate(s)
}

type GetMultiDimTableAllSheetsRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetMultiDimTableAllSheetsRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetMultiDimTableAllSheetsRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetMultiDimTableAllSheetsRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetMultiDimTableAllSheetsRequestTenantContext) SetTenantId(v string) *GetMultiDimTableAllSheetsRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetMultiDimTableAllSheetsRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
