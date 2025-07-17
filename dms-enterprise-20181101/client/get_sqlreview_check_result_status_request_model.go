// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSQLReviewCheckResultStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *GetSQLReviewCheckResultStatusRequest
	GetOrderId() *int64
	SetTid(v int64) *GetSQLReviewCheckResultStatusRequest
	GetTid() *int64
}

type GetSQLReviewCheckResultStatusRequest struct {
	// The ID of the ticket. You can obtain the ticket ID from the response parameters of the [CreateSQLReviewOrder](https://help.aliyun.com/document_detail/257777.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123321
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to obtain the ID of the tenant.
	//
	// example:
	//
	// 1
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetSQLReviewCheckResultStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSQLReviewCheckResultStatusRequest) GoString() string {
	return s.String()
}

func (s *GetSQLReviewCheckResultStatusRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *GetSQLReviewCheckResultStatusRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetSQLReviewCheckResultStatusRequest) SetOrderId(v int64) *GetSQLReviewCheckResultStatusRequest {
	s.OrderId = &v
	return s
}

func (s *GetSQLReviewCheckResultStatusRequest) SetTid(v int64) *GetSQLReviewCheckResultStatusRequest {
	s.Tid = &v
	return s
}

func (s *GetSQLReviewCheckResultStatusRequest) Validate() error {
	return dara.Validate(s)
}
