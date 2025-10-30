// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnBindAXBRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBindId(v string) *UnBindAXBRequest
	GetBindId() *string
	SetCallerParentId(v int64) *UnBindAXBRequest
	GetCallerParentId() *int64
	SetCustomerPoolKey(v string) *UnBindAXBRequest
	GetCustomerPoolKey() *string
	SetOwnerId(v int64) *UnBindAXBRequest
	GetOwnerId() *int64
	SetReqId(v string) *UnBindAXBRequest
	GetReqId() *string
	SetResourceOwnerAccount(v string) *UnBindAXBRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UnBindAXBRequest
	GetResourceOwnerId() *int64
}

type UnBindAXBRequest struct {
	// bindId绑定关系AXB唯一id
	//
	// This parameter is required.
	//
	// example:
	//
	// 4534543
	BindId *string `json:"BindId,omitempty" xml:"BindId,omitempty"`
	// 客户uid
	//
	// example:
	//
	// -
	CallerParentId *int64 `json:"CallerParentId,omitempty" xml:"CallerParentId,omitempty"`
	// 号码池key
	//
	// This parameter is required.
	//
	// example:
	//
	// FC5**********************a1a
	CustomerPoolKey *string `json:"CustomerPoolKey,omitempty" xml:"CustomerPoolKey,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 请求去重ID, reqId最大长度为20位,接入方需要保持原子性
	//
	// This parameter is required.
	//
	// example:
	//
	// 564**********879
	ReqId                *string `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UnBindAXBRequest) String() string {
	return dara.Prettify(s)
}

func (s UnBindAXBRequest) GoString() string {
	return s.String()
}

func (s *UnBindAXBRequest) GetBindId() *string {
	return s.BindId
}

func (s *UnBindAXBRequest) GetCallerParentId() *int64 {
	return s.CallerParentId
}

func (s *UnBindAXBRequest) GetCustomerPoolKey() *string {
	return s.CustomerPoolKey
}

func (s *UnBindAXBRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UnBindAXBRequest) GetReqId() *string {
	return s.ReqId
}

func (s *UnBindAXBRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UnBindAXBRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UnBindAXBRequest) SetBindId(v string) *UnBindAXBRequest {
	s.BindId = &v
	return s
}

func (s *UnBindAXBRequest) SetCallerParentId(v int64) *UnBindAXBRequest {
	s.CallerParentId = &v
	return s
}

func (s *UnBindAXBRequest) SetCustomerPoolKey(v string) *UnBindAXBRequest {
	s.CustomerPoolKey = &v
	return s
}

func (s *UnBindAXBRequest) SetOwnerId(v int64) *UnBindAXBRequest {
	s.OwnerId = &v
	return s
}

func (s *UnBindAXBRequest) SetReqId(v string) *UnBindAXBRequest {
	s.ReqId = &v
	return s
}

func (s *UnBindAXBRequest) SetResourceOwnerAccount(v string) *UnBindAXBRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnBindAXBRequest) SetResourceOwnerId(v int64) *UnBindAXBRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UnBindAXBRequest) Validate() error {
	return dara.Validate(s)
}
