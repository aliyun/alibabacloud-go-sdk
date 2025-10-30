// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSmsSignRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalledNo(v string) *CreateSmsSignRequest
	GetCalledNo() *string
	SetCallerParentId(v int64) *CreateSmsSignRequest
	GetCallerParentId() *int64
	SetCallingNo(v string) *CreateSmsSignRequest
	GetCallingNo() *string
	SetCustomerPoolKey(v string) *CreateSmsSignRequest
	GetCustomerPoolKey() *string
	SetOwnerId(v int64) *CreateSmsSignRequest
	GetOwnerId() *int64
	SetReqId(v string) *CreateSmsSignRequest
	GetReqId() *string
	SetResourceOwnerAccount(v string) *CreateSmsSignRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateSmsSignRequest
	GetResourceOwnerId() *int64
}

type CreateSmsSignRequest struct {
	// 收信人号码
	//
	// This parameter is required.
	//
	// example:
	//
	// 178****3614
	CalledNo *string `json:"CalledNo,omitempty" xml:"CalledNo,omitempty"`
	// 客户uid
	//
	// example:
	//
	// -
	CallerParentId *int64 `json:"CallerParentId,omitempty" xml:"CallerParentId,omitempty"`
	// 发信人号码
	//
	// This parameter is required.
	//
	// example:
	//
	// 150****6539
	CallingNo *string `json:"CallingNo,omitempty" xml:"CallingNo,omitempty"`
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

func (s CreateSmsSignRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSmsSignRequest) GoString() string {
	return s.String()
}

func (s *CreateSmsSignRequest) GetCalledNo() *string {
	return s.CalledNo
}

func (s *CreateSmsSignRequest) GetCallerParentId() *int64 {
	return s.CallerParentId
}

func (s *CreateSmsSignRequest) GetCallingNo() *string {
	return s.CallingNo
}

func (s *CreateSmsSignRequest) GetCustomerPoolKey() *string {
	return s.CustomerPoolKey
}

func (s *CreateSmsSignRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateSmsSignRequest) GetReqId() *string {
	return s.ReqId
}

func (s *CreateSmsSignRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateSmsSignRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateSmsSignRequest) SetCalledNo(v string) *CreateSmsSignRequest {
	s.CalledNo = &v
	return s
}

func (s *CreateSmsSignRequest) SetCallerParentId(v int64) *CreateSmsSignRequest {
	s.CallerParentId = &v
	return s
}

func (s *CreateSmsSignRequest) SetCallingNo(v string) *CreateSmsSignRequest {
	s.CallingNo = &v
	return s
}

func (s *CreateSmsSignRequest) SetCustomerPoolKey(v string) *CreateSmsSignRequest {
	s.CustomerPoolKey = &v
	return s
}

func (s *CreateSmsSignRequest) SetOwnerId(v int64) *CreateSmsSignRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSmsSignRequest) SetReqId(v string) *CreateSmsSignRequest {
	s.ReqId = &v
	return s
}

func (s *CreateSmsSignRequest) SetResourceOwnerAccount(v string) *CreateSmsSignRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateSmsSignRequest) SetResourceOwnerId(v int64) *CreateSmsSignRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateSmsSignRequest) Validate() error {
	return dara.Validate(s)
}
