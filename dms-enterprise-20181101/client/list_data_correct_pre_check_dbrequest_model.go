// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataCorrectPreCheckDBRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *ListDataCorrectPreCheckDBRequest
	GetOrderId() *int64
	SetPageNumber(v int64) *ListDataCorrectPreCheckDBRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListDataCorrectPreCheckDBRequest
	GetPageSize() *int64
	SetTid(v int64) *ListDataCorrectPreCheckDBRequest
	GetTid() *int64
}

type ListDataCorrectPreCheckDBRequest struct {
	// The ID of the ticket for the data change.
	//
	// This parameter is required.
	//
	// example:
	//
	// 432****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The number of the page to return.
	//
	// Valid values: an integer that is greater than 0.
	//
	// Default value: 1.
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
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to query the ID of the tenant.
	//
	// example:
	//
	// 4321****
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListDataCorrectPreCheckDBRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataCorrectPreCheckDBRequest) GoString() string {
	return s.String()
}

func (s *ListDataCorrectPreCheckDBRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *ListDataCorrectPreCheckDBRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListDataCorrectPreCheckDBRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListDataCorrectPreCheckDBRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListDataCorrectPreCheckDBRequest) SetOrderId(v int64) *ListDataCorrectPreCheckDBRequest {
	s.OrderId = &v
	return s
}

func (s *ListDataCorrectPreCheckDBRequest) SetPageNumber(v int64) *ListDataCorrectPreCheckDBRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataCorrectPreCheckDBRequest) SetPageSize(v int64) *ListDataCorrectPreCheckDBRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataCorrectPreCheckDBRequest) SetTid(v int64) *ListDataCorrectPreCheckDBRequest {
	s.Tid = &v
	return s
}

func (s *ListDataCorrectPreCheckDBRequest) Validate() error {
	return dara.Validate(s)
}
