// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListXTelephonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallerParentId(v int64) *ListXTelephonesRequest
	GetCallerParentId() *int64
	SetCustomerPoolKey(v string) *ListXTelephonesRequest
	GetCustomerPoolKey() *string
	SetOwnerId(v int64) *ListXTelephonesRequest
	GetOwnerId() *int64
	SetPageNo(v int64) *ListXTelephonesRequest
	GetPageNo() *int64
	SetPageSize(v int64) *ListXTelephonesRequest
	GetPageSize() *int64
	SetReqId(v string) *ListXTelephonesRequest
	GetReqId() *string
	SetResourceOwnerAccount(v string) *ListXTelephonesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListXTelephonesRequest
	GetResourceOwnerId() *int64
}

type ListXTelephonesRequest struct {
	// 客户uid
	//
	// example:
	//
	// -
	CallerParentId *int64 `json:"CallerParentId,omitempty" xml:"CallerParentId,omitempty"`
	// 号码池key
	//
	// example:
	//
	// FC5**********************a1a
	CustomerPoolKey *string `json:"CustomerPoolKey,omitempty" xml:"CustomerPoolKey,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 页码从1开始
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// 每页条数
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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

func (s ListXTelephonesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListXTelephonesRequest) GoString() string {
	return s.String()
}

func (s *ListXTelephonesRequest) GetCallerParentId() *int64 {
	return s.CallerParentId
}

func (s *ListXTelephonesRequest) GetCustomerPoolKey() *string {
	return s.CustomerPoolKey
}

func (s *ListXTelephonesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListXTelephonesRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListXTelephonesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListXTelephonesRequest) GetReqId() *string {
	return s.ReqId
}

func (s *ListXTelephonesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListXTelephonesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListXTelephonesRequest) SetCallerParentId(v int64) *ListXTelephonesRequest {
	s.CallerParentId = &v
	return s
}

func (s *ListXTelephonesRequest) SetCustomerPoolKey(v string) *ListXTelephonesRequest {
	s.CustomerPoolKey = &v
	return s
}

func (s *ListXTelephonesRequest) SetOwnerId(v int64) *ListXTelephonesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListXTelephonesRequest) SetPageNo(v int64) *ListXTelephonesRequest {
	s.PageNo = &v
	return s
}

func (s *ListXTelephonesRequest) SetPageSize(v int64) *ListXTelephonesRequest {
	s.PageSize = &v
	return s
}

func (s *ListXTelephonesRequest) SetReqId(v string) *ListXTelephonesRequest {
	s.ReqId = &v
	return s
}

func (s *ListXTelephonesRequest) SetResourceOwnerAccount(v string) *ListXTelephonesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListXTelephonesRequest) SetResourceOwnerId(v int64) *ListXTelephonesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListXTelephonesRequest) Validate() error {
	return dara.Validate(s)
}
