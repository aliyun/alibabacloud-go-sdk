// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAXBCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthId(v string) *BindAXBCallRequest
	GetAuthId() *string
	SetCallerParentId(v int64) *BindAXBCallRequest
	GetCallerParentId() *int64
	SetCustomerPoolKey(v string) *BindAXBCallRequest
	GetCustomerPoolKey() *string
	SetExpiration(v int64) *BindAXBCallRequest
	GetExpiration() *int64
	SetOwnerId(v int64) *BindAXBCallRequest
	GetOwnerId() *int64
	SetReqId(v string) *BindAXBCallRequest
	GetReqId() *string
	SetResourceOwnerAccount(v string) *BindAXBCallRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *BindAXBCallRequest
	GetResourceOwnerId() *int64
	SetTelA(v string) *BindAXBCallRequest
	GetTelA() *string
	SetUserData(v string) *BindAXBCallRequest
	GetUserData() *string
}

type BindAXBCallRequest struct {
	// authId绑定关系BX唯一id
	//
	// This parameter is required.
	//
	// example:
	//
	// 12353
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
	// 绑定关系过期失效时间： 取值必须大于0； 单位：秒，ct_time + expiration = 自动解绑时间
	//
	// This parameter is required.
	//
	// example:
	//
	// 7200
	Expiration *int64 `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	OwnerId    *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// 客户A号码
	//
	// This parameter is required.
	//
	// example:
	//
	// 13*******43
	TelA *string `json:"TelA,omitempty" xml:"TelA,omitempty"`
	// 客户自定义参数回调带回
	//
	// example:
	//
	// 000
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s BindAXBCallRequest) String() string {
	return dara.Prettify(s)
}

func (s BindAXBCallRequest) GoString() string {
	return s.String()
}

func (s *BindAXBCallRequest) GetAuthId() *string {
	return s.AuthId
}

func (s *BindAXBCallRequest) GetCallerParentId() *int64 {
	return s.CallerParentId
}

func (s *BindAXBCallRequest) GetCustomerPoolKey() *string {
	return s.CustomerPoolKey
}

func (s *BindAXBCallRequest) GetExpiration() *int64 {
	return s.Expiration
}

func (s *BindAXBCallRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BindAXBCallRequest) GetReqId() *string {
	return s.ReqId
}

func (s *BindAXBCallRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *BindAXBCallRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *BindAXBCallRequest) GetTelA() *string {
	return s.TelA
}

func (s *BindAXBCallRequest) GetUserData() *string {
	return s.UserData
}

func (s *BindAXBCallRequest) SetAuthId(v string) *BindAXBCallRequest {
	s.AuthId = &v
	return s
}

func (s *BindAXBCallRequest) SetCallerParentId(v int64) *BindAXBCallRequest {
	s.CallerParentId = &v
	return s
}

func (s *BindAXBCallRequest) SetCustomerPoolKey(v string) *BindAXBCallRequest {
	s.CustomerPoolKey = &v
	return s
}

func (s *BindAXBCallRequest) SetExpiration(v int64) *BindAXBCallRequest {
	s.Expiration = &v
	return s
}

func (s *BindAXBCallRequest) SetOwnerId(v int64) *BindAXBCallRequest {
	s.OwnerId = &v
	return s
}

func (s *BindAXBCallRequest) SetReqId(v string) *BindAXBCallRequest {
	s.ReqId = &v
	return s
}

func (s *BindAXBCallRequest) SetResourceOwnerAccount(v string) *BindAXBCallRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindAXBCallRequest) SetResourceOwnerId(v int64) *BindAXBCallRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindAXBCallRequest) SetTelA(v string) *BindAXBCallRequest {
	s.TelA = &v
	return s
}

func (s *BindAXBCallRequest) SetUserData(v string) *BindAXBCallRequest {
	s.UserData = &v
	return s
}

func (s *BindAXBCallRequest) Validate() error {
	return dara.Validate(s)
}
