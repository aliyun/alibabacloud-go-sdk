// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataImportSQLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *GetDataImportSQLRequest
	GetOrderId() *int64
	SetSqlId(v int64) *GetDataImportSQLRequest
	GetSqlId() *int64
	SetTid(v int64) *GetDataImportSQLRequest
	GetTid() *int64
}

type GetDataImportSQLRequest struct {
	// The ticket ID. You can call the [ListOrders](https://help.aliyun.com/document_detail/144643.html) operation to query the ticket ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 14****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The SQL ID. You can call the ListDataImportSQLPreCheckDetail operation to query the SQL ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 15****
	SqlId *int64 `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
	// The tenant ID. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to query the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetDataImportSQLRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataImportSQLRequest) GoString() string {
	return s.String()
}

func (s *GetDataImportSQLRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *GetDataImportSQLRequest) GetSqlId() *int64 {
	return s.SqlId
}

func (s *GetDataImportSQLRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetDataImportSQLRequest) SetOrderId(v int64) *GetDataImportSQLRequest {
	s.OrderId = &v
	return s
}

func (s *GetDataImportSQLRequest) SetSqlId(v int64) *GetDataImportSQLRequest {
	s.SqlId = &v
	return s
}

func (s *GetDataImportSQLRequest) SetTid(v int64) *GetDataImportSQLRequest {
	s.Tid = &v
	return s
}

func (s *GetDataImportSQLRequest) Validate() error {
	return dara.Validate(s)
}
