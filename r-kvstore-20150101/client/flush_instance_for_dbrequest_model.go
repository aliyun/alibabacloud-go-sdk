// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlushInstanceForDBRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbIndex(v int32) *FlushInstanceForDBRequest
	GetDbIndex() *int32
	SetInstanceId(v string) *FlushInstanceForDBRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *FlushInstanceForDBRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *FlushInstanceForDBRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *FlushInstanceForDBRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *FlushInstanceForDBRequest
	GetResourceOwnerId() *int64
}

type FlushInstanceForDBRequest struct {
	// The index number of the database. Valid values: 0 to 255.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DbIndex *int32 `json:"DbIndex,omitempty" xml:"DbIndex,omitempty"`
	// The instance ID. You can call the [DescribeInstances](https://help.aliyun.com/document_detail/473778.html) operation to query the ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s FlushInstanceForDBRequest) String() string {
	return dara.Prettify(s)
}

func (s FlushInstanceForDBRequest) GoString() string {
	return s.String()
}

func (s *FlushInstanceForDBRequest) GetDbIndex() *int32 {
	return s.DbIndex
}

func (s *FlushInstanceForDBRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *FlushInstanceForDBRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *FlushInstanceForDBRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *FlushInstanceForDBRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *FlushInstanceForDBRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *FlushInstanceForDBRequest) SetDbIndex(v int32) *FlushInstanceForDBRequest {
	s.DbIndex = &v
	return s
}

func (s *FlushInstanceForDBRequest) SetInstanceId(v string) *FlushInstanceForDBRequest {
	s.InstanceId = &v
	return s
}

func (s *FlushInstanceForDBRequest) SetOwnerAccount(v string) *FlushInstanceForDBRequest {
	s.OwnerAccount = &v
	return s
}

func (s *FlushInstanceForDBRequest) SetOwnerId(v int64) *FlushInstanceForDBRequest {
	s.OwnerId = &v
	return s
}

func (s *FlushInstanceForDBRequest) SetResourceOwnerAccount(v string) *FlushInstanceForDBRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *FlushInstanceForDBRequest) SetResourceOwnerId(v int64) *FlushInstanceForDBRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *FlushInstanceForDBRequest) Validate() error {
	return dara.Validate(s)
}
