// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrderDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMemberUid(v int64) *GetOrderDetailRequest
	GetMemberUid() *int64
	SetOrderId(v string) *GetOrderDetailRequest
	GetOrderId() *string
	SetOwnerId(v int64) *GetOrderDetailRequest
	GetOwnerId() *int64
}

type GetOrderDetailRequest struct {
	// example:
	//
	// 1715322405372273
	MemberUid *int64 `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 233501558440169
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s GetOrderDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOrderDetailRequest) GoString() string {
	return s.String()
}

func (s *GetOrderDetailRequest) GetMemberUid() *int64 {
	return s.MemberUid
}

func (s *GetOrderDetailRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *GetOrderDetailRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetOrderDetailRequest) SetMemberUid(v int64) *GetOrderDetailRequest {
	s.MemberUid = &v
	return s
}

func (s *GetOrderDetailRequest) SetOrderId(v string) *GetOrderDetailRequest {
	s.OrderId = &v
	return s
}

func (s *GetOrderDetailRequest) SetOwnerId(v int64) *GetOrderDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *GetOrderDetailRequest) Validate() error {
	return dara.Validate(s)
}
