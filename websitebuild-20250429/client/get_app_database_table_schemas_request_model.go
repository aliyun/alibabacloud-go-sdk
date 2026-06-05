// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppDatabaseTableSchemasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *GetAppDatabaseTableSchemasRequest
	GetBizId() *string
	SetTableName(v string) *GetAppDatabaseTableSchemasRequest
	GetTableName() *string
}

type GetAppDatabaseTableSchemasRequest struct {
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// default.ai_advertising_material_rec_train_v1103
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetAppDatabaseTableSchemasRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAppDatabaseTableSchemasRequest) GoString() string {
	return s.String()
}

func (s *GetAppDatabaseTableSchemasRequest) GetBizId() *string {
	return s.BizId
}

func (s *GetAppDatabaseTableSchemasRequest) GetTableName() *string {
	return s.TableName
}

func (s *GetAppDatabaseTableSchemasRequest) SetBizId(v string) *GetAppDatabaseTableSchemasRequest {
	s.BizId = &v
	return s
}

func (s *GetAppDatabaseTableSchemasRequest) SetTableName(v string) *GetAppDatabaseTableSchemasRequest {
	s.TableName = &v
	return s
}

func (s *GetAppDatabaseTableSchemasRequest) Validate() error {
	return dara.Validate(s)
}
