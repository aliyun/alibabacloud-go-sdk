// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataImportSQLTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *ListDataImportSQLTypeRequest
	GetOrderId() *int64
	SetTid(v int64) *ListDataImportSQLTypeRequest
	GetTid() *int64
}

type ListDataImportSQLTypeRequest struct {
	// The ticket ID. You can call the [ListOrders](https://help.aliyun.com/document_detail/144643.html) operation to query the ticket ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 420****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The tenant ID. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to query the tenant ID.
	//
	// example:
	//
	// 1
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListDataImportSQLTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataImportSQLTypeRequest) GoString() string {
	return s.String()
}

func (s *ListDataImportSQLTypeRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *ListDataImportSQLTypeRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListDataImportSQLTypeRequest) SetOrderId(v int64) *ListDataImportSQLTypeRequest {
	s.OrderId = &v
	return s
}

func (s *ListDataImportSQLTypeRequest) SetTid(v int64) *ListDataImportSQLTypeRequest {
	s.Tid = &v
	return s
}

func (s *ListDataImportSQLTypeRequest) Validate() error {
	return dara.Validate(s)
}
