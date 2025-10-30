// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetXDefaultConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallerParentId(v int64) *GetXDefaultConfigRequest
	GetCallerParentId() *int64
	SetCustomerPoolKey(v string) *GetXDefaultConfigRequest
	GetCustomerPoolKey() *string
	SetOwnerId(v int64) *GetXDefaultConfigRequest
	GetOwnerId() *int64
	SetReqId(v string) *GetXDefaultConfigRequest
	GetReqId() *string
	SetResourceOwnerAccount(v string) *GetXDefaultConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetXDefaultConfigRequest
	GetResourceOwnerId() *int64
	SetTelX(v string) *GetXDefaultConfigRequest
	GetTelX() *string
}

type GetXDefaultConfigRequest struct {
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

func (s GetXDefaultConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetXDefaultConfigRequest) GoString() string {
	return s.String()
}

func (s *GetXDefaultConfigRequest) GetCallerParentId() *int64 {
	return s.CallerParentId
}

func (s *GetXDefaultConfigRequest) GetCustomerPoolKey() *string {
	return s.CustomerPoolKey
}

func (s *GetXDefaultConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetXDefaultConfigRequest) GetReqId() *string {
	return s.ReqId
}

func (s *GetXDefaultConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetXDefaultConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetXDefaultConfigRequest) GetTelX() *string {
	return s.TelX
}

func (s *GetXDefaultConfigRequest) SetCallerParentId(v int64) *GetXDefaultConfigRequest {
	s.CallerParentId = &v
	return s
}

func (s *GetXDefaultConfigRequest) SetCustomerPoolKey(v string) *GetXDefaultConfigRequest {
	s.CustomerPoolKey = &v
	return s
}

func (s *GetXDefaultConfigRequest) SetOwnerId(v int64) *GetXDefaultConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *GetXDefaultConfigRequest) SetReqId(v string) *GetXDefaultConfigRequest {
	s.ReqId = &v
	return s
}

func (s *GetXDefaultConfigRequest) SetResourceOwnerAccount(v string) *GetXDefaultConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetXDefaultConfigRequest) SetResourceOwnerId(v int64) *GetXDefaultConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetXDefaultConfigRequest) SetTelX(v string) *GetXDefaultConfigRequest {
	s.TelX = &v
	return s
}

func (s *GetXDefaultConfigRequest) Validate() error {
	return dara.Validate(s)
}
