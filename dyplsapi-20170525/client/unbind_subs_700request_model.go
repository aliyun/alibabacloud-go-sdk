// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindSubs700Request interface {
	dara.Model
	String() string
	GoString() string
	SetIndustrialId(v string) *UnbindSubs700Request
	GetIndustrialId() *string
	SetOrderId(v string) *UnbindSubs700Request
	GetOrderId() *string
	SetOwnerId(v int64) *UnbindSubs700Request
	GetOwnerId() *int64
	SetPoolKey(v string) *UnbindSubs700Request
	GetPoolKey() *string
	SetResourceOwnerAccount(v string) *UnbindSubs700Request
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UnbindSubs700Request
	GetResourceOwnerId() *int64
	SetSubsId(v int64) *UnbindSubs700Request
	GetSubsId() *int64
	SetTelX(v string) *UnbindSubs700Request
	GetTelX() *string
}

type UnbindSubs700Request struct {
	// example:
	//
	// 700.100.1/12345678
	IndustrialId *string `json:"IndustrialId,omitempty" xml:"IndustrialId,omitempty"`
	// example:
	//
	// 12345678
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FC10000016848****
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100000****
	SubsId *int64 `json:"SubsId,omitempty" xml:"SubsId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 700********0000
	TelX *string `json:"TelX,omitempty" xml:"TelX,omitempty"`
}

func (s UnbindSubs700Request) String() string {
	return dara.Prettify(s)
}

func (s UnbindSubs700Request) GoString() string {
	return s.String()
}

func (s *UnbindSubs700Request) GetIndustrialId() *string {
	return s.IndustrialId
}

func (s *UnbindSubs700Request) GetOrderId() *string {
	return s.OrderId
}

func (s *UnbindSubs700Request) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UnbindSubs700Request) GetPoolKey() *string {
	return s.PoolKey
}

func (s *UnbindSubs700Request) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UnbindSubs700Request) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UnbindSubs700Request) GetSubsId() *int64 {
	return s.SubsId
}

func (s *UnbindSubs700Request) GetTelX() *string {
	return s.TelX
}

func (s *UnbindSubs700Request) SetIndustrialId(v string) *UnbindSubs700Request {
	s.IndustrialId = &v
	return s
}

func (s *UnbindSubs700Request) SetOrderId(v string) *UnbindSubs700Request {
	s.OrderId = &v
	return s
}

func (s *UnbindSubs700Request) SetOwnerId(v int64) *UnbindSubs700Request {
	s.OwnerId = &v
	return s
}

func (s *UnbindSubs700Request) SetPoolKey(v string) *UnbindSubs700Request {
	s.PoolKey = &v
	return s
}

func (s *UnbindSubs700Request) SetResourceOwnerAccount(v string) *UnbindSubs700Request {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnbindSubs700Request) SetResourceOwnerId(v int64) *UnbindSubs700Request {
	s.ResourceOwnerId = &v
	return s
}

func (s *UnbindSubs700Request) SetSubsId(v int64) *UnbindSubs700Request {
	s.SubsId = &v
	return s
}

func (s *UnbindSubs700Request) SetTelX(v string) *UnbindSubs700Request {
	s.TelX = &v
	return s
}

func (s *UnbindSubs700Request) Validate() error {
	return dara.Validate(s)
}
