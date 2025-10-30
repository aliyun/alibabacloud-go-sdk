// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySoundRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallId(v string) *QuerySoundRecordRequest
	GetCallId() *string
	SetCallerParentId(v int64) *QuerySoundRecordRequest
	GetCallerParentId() *int64
	SetCustomerPoolKey(v string) *QuerySoundRecordRequest
	GetCustomerPoolKey() *string
	SetOwnerId(v int64) *QuerySoundRecordRequest
	GetOwnerId() *int64
	SetReqId(v string) *QuerySoundRecordRequest
	GetReqId() *string
	SetResourceOwnerAccount(v string) *QuerySoundRecordRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QuerySoundRecordRequest
	GetResourceOwnerId() *int64
}

type QuerySoundRecordRequest struct {
	// 本次呼叫唯一id
	//
	// This parameter is required.
	//
	// example:
	//
	// ac445343254
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
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

func (s QuerySoundRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySoundRecordRequest) GoString() string {
	return s.String()
}

func (s *QuerySoundRecordRequest) GetCallId() *string {
	return s.CallId
}

func (s *QuerySoundRecordRequest) GetCallerParentId() *int64 {
	return s.CallerParentId
}

func (s *QuerySoundRecordRequest) GetCustomerPoolKey() *string {
	return s.CustomerPoolKey
}

func (s *QuerySoundRecordRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QuerySoundRecordRequest) GetReqId() *string {
	return s.ReqId
}

func (s *QuerySoundRecordRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QuerySoundRecordRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QuerySoundRecordRequest) SetCallId(v string) *QuerySoundRecordRequest {
	s.CallId = &v
	return s
}

func (s *QuerySoundRecordRequest) SetCallerParentId(v int64) *QuerySoundRecordRequest {
	s.CallerParentId = &v
	return s
}

func (s *QuerySoundRecordRequest) SetCustomerPoolKey(v string) *QuerySoundRecordRequest {
	s.CustomerPoolKey = &v
	return s
}

func (s *QuerySoundRecordRequest) SetOwnerId(v int64) *QuerySoundRecordRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySoundRecordRequest) SetReqId(v string) *QuerySoundRecordRequest {
	s.ReqId = &v
	return s
}

func (s *QuerySoundRecordRequest) SetResourceOwnerAccount(v string) *QuerySoundRecordRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySoundRecordRequest) SetResourceOwnerId(v int64) *QuerySoundRecordRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuerySoundRecordRequest) Validate() error {
	return dara.Validate(s)
}
