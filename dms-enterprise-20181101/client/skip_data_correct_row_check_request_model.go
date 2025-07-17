// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSkipDataCorrectRowCheckRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *SkipDataCorrectRowCheckRequest
	GetOrderId() *int64
	SetRealLoginUserUid(v string) *SkipDataCorrectRowCheckRequest
	GetRealLoginUserUid() *string
	SetReason(v string) *SkipDataCorrectRowCheckRequest
	GetReason() *string
	SetTid(v int64) *SkipDataCorrectRowCheckRequest
	GetTid() *int64
}

type SkipDataCorrectRowCheckRequest struct {
	// The ticket ID. You can call the [ListOrders](https://help.aliyun.com/document_detail/144643.html) operation to obtain the ticket ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 420****
	OrderId          *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RealLoginUserUid *string `json:"RealLoginUserUid,omitempty" xml:"RealLoginUserUid,omitempty"`
	// The reason for skipping the verification on the number of rows in the precheck for data change.
	//
	// This parameter is required.
	//
	// example:
	//
	// save test time
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The tenant ID. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s SkipDataCorrectRowCheckRequest) String() string {
	return dara.Prettify(s)
}

func (s SkipDataCorrectRowCheckRequest) GoString() string {
	return s.String()
}

func (s *SkipDataCorrectRowCheckRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *SkipDataCorrectRowCheckRequest) GetRealLoginUserUid() *string {
	return s.RealLoginUserUid
}

func (s *SkipDataCorrectRowCheckRequest) GetReason() *string {
	return s.Reason
}

func (s *SkipDataCorrectRowCheckRequest) GetTid() *int64 {
	return s.Tid
}

func (s *SkipDataCorrectRowCheckRequest) SetOrderId(v int64) *SkipDataCorrectRowCheckRequest {
	s.OrderId = &v
	return s
}

func (s *SkipDataCorrectRowCheckRequest) SetRealLoginUserUid(v string) *SkipDataCorrectRowCheckRequest {
	s.RealLoginUserUid = &v
	return s
}

func (s *SkipDataCorrectRowCheckRequest) SetReason(v string) *SkipDataCorrectRowCheckRequest {
	s.Reason = &v
	return s
}

func (s *SkipDataCorrectRowCheckRequest) SetTid(v int64) *SkipDataCorrectRowCheckRequest {
	s.Tid = &v
	return s
}

func (s *SkipDataCorrectRowCheckRequest) Validate() error {
	return dara.Validate(s)
}
