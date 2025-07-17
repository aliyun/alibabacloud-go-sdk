// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataCorrectPreCheckSQLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbId(v int64) *ListDataCorrectPreCheckSQLRequest
	GetDbId() *int64
	SetOrderId(v int64) *ListDataCorrectPreCheckSQLRequest
	GetOrderId() *int64
	SetPageNumber(v int64) *ListDataCorrectPreCheckSQLRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListDataCorrectPreCheckSQLRequest
	GetPageSize() *int64
	SetTid(v int64) *ListDataCorrectPreCheckSQLRequest
	GetTid() *int64
}

type ListDataCorrectPreCheckSQLRequest struct {
	// The ID of the database. The database can be a physical database or a logical database.
	//
	// 	- To query the ID of a physical database, call the [ListDatabases](https://help.aliyun.com/document_detail/141873.html) or [SearchDatabase](https://help.aliyun.com/document_detail/141876.html) operation.
	//
	// 	- To query the ID of a logical database, call the [ListLogicDatabases](https://help.aliyun.com/document_detail/141874.html) or [SearchDatabase](https://help.aliyun.com/document_detail/141876.html) operation.
	//
	// example:
	//
	// 1930****
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// The ID of the data change ticket. You can call the [ListOrders](https://help.aliyun.com/document_detail/144643.html) operation to query the ID of the data change ticket.
	//
	// This parameter is required.
	//
	// example:
	//
	// 453****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) operation to query the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListDataCorrectPreCheckSQLRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataCorrectPreCheckSQLRequest) GoString() string {
	return s.String()
}

func (s *ListDataCorrectPreCheckSQLRequest) GetDbId() *int64 {
	return s.DbId
}

func (s *ListDataCorrectPreCheckSQLRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *ListDataCorrectPreCheckSQLRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListDataCorrectPreCheckSQLRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListDataCorrectPreCheckSQLRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListDataCorrectPreCheckSQLRequest) SetDbId(v int64) *ListDataCorrectPreCheckSQLRequest {
	s.DbId = &v
	return s
}

func (s *ListDataCorrectPreCheckSQLRequest) SetOrderId(v int64) *ListDataCorrectPreCheckSQLRequest {
	s.OrderId = &v
	return s
}

func (s *ListDataCorrectPreCheckSQLRequest) SetPageNumber(v int64) *ListDataCorrectPreCheckSQLRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataCorrectPreCheckSQLRequest) SetPageSize(v int64) *ListDataCorrectPreCheckSQLRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataCorrectPreCheckSQLRequest) SetTid(v int64) *ListDataCorrectPreCheckSQLRequest {
	s.Tid = &v
	return s
}

func (s *ListDataCorrectPreCheckSQLRequest) Validate() error {
	return dara.Validate(s)
}
