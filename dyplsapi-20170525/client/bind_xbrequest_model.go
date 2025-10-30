// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindXBRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallerParentId(v int64) *BindXBRequest
	GetCallerParentId() *int64
	SetCustomerPoolKey(v string) *BindXBRequest
	GetCustomerPoolKey() *string
	SetOwnerId(v int64) *BindXBRequest
	GetOwnerId() *int64
	SetReqId(v string) *BindXBRequest
	GetReqId() *string
	SetResourceOwnerAccount(v string) *BindXBRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *BindXBRequest
	GetResourceOwnerId() *int64
	SetTelB(v string) *BindXBRequest
	GetTelB() *string
	SetTelX(v string) *BindXBRequest
	GetTelX() *string
	SetUserData(v string) *BindXBRequest
	GetUserData() *string
}

type BindXBRequest struct {
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
	// 员工真实号码
	//
	// This parameter is required.
	//
	// example:
	//
	// 18*******22
	TelB *string `json:"TelB,omitempty" xml:"TelB,omitempty"`
	// X号码
	//
	// This parameter is required.
	//
	// example:
	//
	// 17*******22
	TelX *string `json:"TelX,omitempty" xml:"TelX,omitempty"`
	// 客户自定义参数回调带回
	//
	// example:
	//
	// 000
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s BindXBRequest) String() string {
	return dara.Prettify(s)
}

func (s BindXBRequest) GoString() string {
	return s.String()
}

func (s *BindXBRequest) GetCallerParentId() *int64 {
	return s.CallerParentId
}

func (s *BindXBRequest) GetCustomerPoolKey() *string {
	return s.CustomerPoolKey
}

func (s *BindXBRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BindXBRequest) GetReqId() *string {
	return s.ReqId
}

func (s *BindXBRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *BindXBRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *BindXBRequest) GetTelB() *string {
	return s.TelB
}

func (s *BindXBRequest) GetTelX() *string {
	return s.TelX
}

func (s *BindXBRequest) GetUserData() *string {
	return s.UserData
}

func (s *BindXBRequest) SetCallerParentId(v int64) *BindXBRequest {
	s.CallerParentId = &v
	return s
}

func (s *BindXBRequest) SetCustomerPoolKey(v string) *BindXBRequest {
	s.CustomerPoolKey = &v
	return s
}

func (s *BindXBRequest) SetOwnerId(v int64) *BindXBRequest {
	s.OwnerId = &v
	return s
}

func (s *BindXBRequest) SetReqId(v string) *BindXBRequest {
	s.ReqId = &v
	return s
}

func (s *BindXBRequest) SetResourceOwnerAccount(v string) *BindXBRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindXBRequest) SetResourceOwnerId(v int64) *BindXBRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindXBRequest) SetTelB(v string) *BindXBRequest {
	s.TelB = &v
	return s
}

func (s *BindXBRequest) SetTelX(v string) *BindXBRequest {
	s.TelX = &v
	return s
}

func (s *BindXBRequest) SetUserData(v string) *BindXBRequest {
	s.UserData = &v
	return s
}

func (s *BindXBRequest) Validate() error {
	return dara.Validate(s)
}
