// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableColumnsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalog(v string) *GetTableColumnsRequest
	GetCatalog() *string
	SetOpTenantId(v int64) *GetTableColumnsRequest
	GetOpTenantId() *int64
	SetTableName(v string) *GetTableColumnsRequest
	GetTableName() *string
}

type GetTableColumnsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// LD_test01_dev
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// t_test01
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetTableColumnsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTableColumnsRequest) GoString() string {
	return s.String()
}

func (s *GetTableColumnsRequest) GetCatalog() *string {
	return s.Catalog
}

func (s *GetTableColumnsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetTableColumnsRequest) GetTableName() *string {
	return s.TableName
}

func (s *GetTableColumnsRequest) SetCatalog(v string) *GetTableColumnsRequest {
	s.Catalog = &v
	return s
}

func (s *GetTableColumnsRequest) SetOpTenantId(v int64) *GetTableColumnsRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetTableColumnsRequest) SetTableName(v string) *GetTableColumnsRequest {
	s.TableName = &v
	return s
}

func (s *GetTableColumnsRequest) Validate() error {
	return dara.Validate(s)
}
