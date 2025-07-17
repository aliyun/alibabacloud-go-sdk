// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataImportSQLPreCheckDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *ListDataImportSQLPreCheckDetailRequest
	GetOrderId() *int64
	SetPageNumer(v int64) *ListDataImportSQLPreCheckDetailRequest
	GetPageNumer() *int64
	SetPageSize(v int64) *ListDataImportSQLPreCheckDetailRequest
	GetPageSize() *int64
	SetSqlType(v string) *ListDataImportSQLPreCheckDetailRequest
	GetSqlType() *string
	SetStatusCode(v string) *ListDataImportSQLPreCheckDetailRequest
	GetStatusCode() *string
	SetTid(v int64) *ListDataImportSQLPreCheckDetailRequest
	GetTid() *int64
}

type ListDataImportSQLPreCheckDetailRequest struct {
	// The ticket ID. You can call the [ListOrders](https://help.aliyun.com/document_detail/144643.html) operation to query the ticket ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The page number. Pages start from page 1.
	//
	// example:
	//
	// 1
	PageNumer *int64 `json:"PageNumer,omitempty" xml:"PageNumer,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the SQL statement. Valid values:
	//
	// 	- **SELECT**
	//
	// 	- **INSERT**
	//
	// 	- **DELETE**
	//
	// 	- **CREATE_TABLE**
	//
	// > You can log on to the Data Management (DMS) console and choose **Security and Specifications*	- > **Operation Audit*	- in the top navigation bar to view more types of SQL statements.
	//
	// example:
	//
	// INSERT
	SqlType *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	// The state of the ticket. If you leave this parameter empty, all the states are queried by default. Valid values:
	//
	// 	- **INIT**: The ticket is being initialized.
	//
	// 	- **RUNNING**: The ticket is in progress.
	//
	// 	- **SUCCESS**: The ticket is complete.
	//
	// 	- **TIMEOUT**: The ticket is skipped due to timeout.
	//
	// 	- **FAIL**: The ticket fails.
	//
	// example:
	//
	// SUCCESS
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// The tenant ID. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to query the tenant ID.
	//
	// example:
	//
	// 1
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListDataImportSQLPreCheckDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataImportSQLPreCheckDetailRequest) GoString() string {
	return s.String()
}

func (s *ListDataImportSQLPreCheckDetailRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *ListDataImportSQLPreCheckDetailRequest) GetPageNumer() *int64 {
	return s.PageNumer
}

func (s *ListDataImportSQLPreCheckDetailRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListDataImportSQLPreCheckDetailRequest) GetSqlType() *string {
	return s.SqlType
}

func (s *ListDataImportSQLPreCheckDetailRequest) GetStatusCode() *string {
	return s.StatusCode
}

func (s *ListDataImportSQLPreCheckDetailRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListDataImportSQLPreCheckDetailRequest) SetOrderId(v int64) *ListDataImportSQLPreCheckDetailRequest {
	s.OrderId = &v
	return s
}

func (s *ListDataImportSQLPreCheckDetailRequest) SetPageNumer(v int64) *ListDataImportSQLPreCheckDetailRequest {
	s.PageNumer = &v
	return s
}

func (s *ListDataImportSQLPreCheckDetailRequest) SetPageSize(v int64) *ListDataImportSQLPreCheckDetailRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataImportSQLPreCheckDetailRequest) SetSqlType(v string) *ListDataImportSQLPreCheckDetailRequest {
	s.SqlType = &v
	return s
}

func (s *ListDataImportSQLPreCheckDetailRequest) SetStatusCode(v string) *ListDataImportSQLPreCheckDetailRequest {
	s.StatusCode = &v
	return s
}

func (s *ListDataImportSQLPreCheckDetailRequest) SetTid(v int64) *ListDataImportSQLPreCheckDetailRequest {
	s.Tid = &v
	return s
}

func (s *ListDataImportSQLPreCheckDetailRequest) Validate() error {
	return dara.Validate(s)
}
