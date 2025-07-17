// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStructSyncExecSqlDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *GetStructSyncExecSqlDetailRequest
	GetOrderId() *int64
	SetPageNumber(v int64) *GetStructSyncExecSqlDetailRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *GetStructSyncExecSqlDetailRequest
	GetPageSize() *int64
	SetTid(v int64) *GetStructSyncExecSqlDetailRequest
	GetTid() *int64
}

type GetStructSyncExecSqlDetailRequest struct {
	// The ID of the ticket.
	//
	// This parameter is required.
	//
	// example:
	//
	// 342153
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
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the tenant.
	//
	// > To view the tenant ID, move the pointer over the profile picture in the upper-right corner of the Data Management (DMS) console. For more information, see [Manage DMS tenants](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 1
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetStructSyncExecSqlDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStructSyncExecSqlDetailRequest) GoString() string {
	return s.String()
}

func (s *GetStructSyncExecSqlDetailRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *GetStructSyncExecSqlDetailRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *GetStructSyncExecSqlDetailRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetStructSyncExecSqlDetailRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetStructSyncExecSqlDetailRequest) SetOrderId(v int64) *GetStructSyncExecSqlDetailRequest {
	s.OrderId = &v
	return s
}

func (s *GetStructSyncExecSqlDetailRequest) SetPageNumber(v int64) *GetStructSyncExecSqlDetailRequest {
	s.PageNumber = &v
	return s
}

func (s *GetStructSyncExecSqlDetailRequest) SetPageSize(v int64) *GetStructSyncExecSqlDetailRequest {
	s.PageSize = &v
	return s
}

func (s *GetStructSyncExecSqlDetailRequest) SetTid(v int64) *GetStructSyncExecSqlDetailRequest {
	s.Tid = &v
	return s
}

func (s *GetStructSyncExecSqlDetailRequest) Validate() error {
	return dara.Validate(s)
}
