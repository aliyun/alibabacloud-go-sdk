// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnBindXBRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthId(v string) *UnBindXBRequest
	GetAuthId() *string
	SetCallerParentId(v int64) *UnBindXBRequest
	GetCallerParentId() *int64
	SetCustomerPoolKey(v string) *UnBindXBRequest
	GetCustomerPoolKey() *string
	SetOwnerId(v int64) *UnBindXBRequest
	GetOwnerId() *int64
	SetReqId(v string) *UnBindXBRequest
	GetReqId() *string
	SetResourceOwnerAccount(v string) *UnBindXBRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UnBindXBRequest
	GetResourceOwnerId() *int64
	SetTelX(v string) *UnBindXBRequest
	GetTelX() *string
}

type UnBindXBRequest struct {
	// authId绑定关系BX唯一id
	//
	// This parameter is required.
	//
	// example:
	//
	// 34*****46
	AuthId *string `json:"AuthId,omitempty" xml:"AuthId,omitempty"`
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
	// X号码
	//
	// This parameter is required.
	//
	// example:
	//
	// 17*******22
	TelX *string `json:"TelX,omitempty" xml:"TelX,omitempty"`
}

func (s UnBindXBRequest) String() string {
	return dara.Prettify(s)
}

func (s UnBindXBRequest) GoString() string {
	return s.String()
}

func (s *UnBindXBRequest) GetAuthId() *string {
	return s.AuthId
}

func (s *UnBindXBRequest) GetCallerParentId() *int64 {
	return s.CallerParentId
}

func (s *UnBindXBRequest) GetCustomerPoolKey() *string {
	return s.CustomerPoolKey
}

func (s *UnBindXBRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UnBindXBRequest) GetReqId() *string {
	return s.ReqId
}

func (s *UnBindXBRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UnBindXBRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UnBindXBRequest) GetTelX() *string {
	return s.TelX
}

func (s *UnBindXBRequest) SetAuthId(v string) *UnBindXBRequest {
	s.AuthId = &v
	return s
}

func (s *UnBindXBRequest) SetCallerParentId(v int64) *UnBindXBRequest {
	s.CallerParentId = &v
	return s
}

func (s *UnBindXBRequest) SetCustomerPoolKey(v string) *UnBindXBRequest {
	s.CustomerPoolKey = &v
	return s
}

func (s *UnBindXBRequest) SetOwnerId(v int64) *UnBindXBRequest {
	s.OwnerId = &v
	return s
}

func (s *UnBindXBRequest) SetReqId(v string) *UnBindXBRequest {
	s.ReqId = &v
	return s
}

func (s *UnBindXBRequest) SetResourceOwnerAccount(v string) *UnBindXBRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnBindXBRequest) SetResourceOwnerId(v int64) *UnBindXBRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UnBindXBRequest) SetTelX(v string) *UnBindXBRequest {
	s.TelX = &v
	return s
}

func (s *UnBindXBRequest) Validate() error {
	return dara.Validate(s)
}
