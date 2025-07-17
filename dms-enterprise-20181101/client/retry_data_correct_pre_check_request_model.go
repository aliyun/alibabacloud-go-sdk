// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryDataCorrectPreCheckRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *RetryDataCorrectPreCheckRequest
	GetOrderId() *int64
	SetRealLoginUserUid(v string) *RetryDataCorrectPreCheckRequest
	GetRealLoginUserUid() *string
	SetTid(v int64) *RetryDataCorrectPreCheckRequest
	GetTid() *int64
}

type RetryDataCorrectPreCheckRequest struct {
	// The ID of the data change ticket. You can call the [ListOrders](https://help.aliyun.com/document_detail/144643.html) operation to query the ID of the data change ticket.
	//
	// This parameter is required.
	//
	// example:
	//
	// 414****
	OrderId          *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RealLoginUserUid *string `json:"RealLoginUserUid,omitempty" xml:"RealLoginUserUid,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) operation to query the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s RetryDataCorrectPreCheckRequest) String() string {
	return dara.Prettify(s)
}

func (s RetryDataCorrectPreCheckRequest) GoString() string {
	return s.String()
}

func (s *RetryDataCorrectPreCheckRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *RetryDataCorrectPreCheckRequest) GetRealLoginUserUid() *string {
	return s.RealLoginUserUid
}

func (s *RetryDataCorrectPreCheckRequest) GetTid() *int64 {
	return s.Tid
}

func (s *RetryDataCorrectPreCheckRequest) SetOrderId(v int64) *RetryDataCorrectPreCheckRequest {
	s.OrderId = &v
	return s
}

func (s *RetryDataCorrectPreCheckRequest) SetRealLoginUserUid(v string) *RetryDataCorrectPreCheckRequest {
	s.RealLoginUserUid = &v
	return s
}

func (s *RetryDataCorrectPreCheckRequest) SetTid(v int64) *RetryDataCorrectPreCheckRequest {
	s.Tid = &v
	return s
}

func (s *RetryDataCorrectPreCheckRequest) Validate() error {
	return dara.Validate(s)
}
