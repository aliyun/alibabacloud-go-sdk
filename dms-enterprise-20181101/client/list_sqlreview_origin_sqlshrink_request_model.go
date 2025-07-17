// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSQLReviewOriginSQLShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderActionDetailShrink(v string) *ListSQLReviewOriginSQLShrinkRequest
	GetOrderActionDetailShrink() *string
	SetOrderId(v int64) *ListSQLReviewOriginSQLShrinkRequest
	GetOrderId() *int64
	SetTid(v int64) *ListSQLReviewOriginSQLShrinkRequest
	GetTid() *int64
}

type ListSQLReviewOriginSQLShrinkRequest struct {
	// The parameters that are used to filter SQL statements involved in the ticket.
	OrderActionDetailShrink *string `json:"OrderActionDetail,omitempty" xml:"OrderActionDetail,omitempty"`
	// The ID of the SQL review ticket. You can call the [CreateSQLReviewOrder](https://help.aliyun.com/document_detail/257777.html) operation to query the ticket ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123321
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The tenant ID. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to query the tenant ID.
	//
	// example:
	//
	// 1
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListSQLReviewOriginSQLShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSQLReviewOriginSQLShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListSQLReviewOriginSQLShrinkRequest) GetOrderActionDetailShrink() *string {
	return s.OrderActionDetailShrink
}

func (s *ListSQLReviewOriginSQLShrinkRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *ListSQLReviewOriginSQLShrinkRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListSQLReviewOriginSQLShrinkRequest) SetOrderActionDetailShrink(v string) *ListSQLReviewOriginSQLShrinkRequest {
	s.OrderActionDetailShrink = &v
	return s
}

func (s *ListSQLReviewOriginSQLShrinkRequest) SetOrderId(v int64) *ListSQLReviewOriginSQLShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *ListSQLReviewOriginSQLShrinkRequest) SetTid(v int64) *ListSQLReviewOriginSQLShrinkRequest {
	s.Tid = &v
	return s
}

func (s *ListSQLReviewOriginSQLShrinkRequest) Validate() error {
	return dara.Validate(s)
}
