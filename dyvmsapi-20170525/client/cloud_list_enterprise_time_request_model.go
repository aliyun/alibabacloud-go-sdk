// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudListEnterpriseTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *CloudListEnterpriseTimeRequest
	GetEnterpriseId() *int64
	SetName(v string) *CloudListEnterpriseTimeRequest
	GetName() *string
	SetOwnerId(v int64) *CloudListEnterpriseTimeRequest
	GetOwnerId() *int64
	SetPriority(v int64) *CloudListEnterpriseTimeRequest
	GetPriority() *int64
	SetResourceOwnerAccount(v string) *CloudListEnterpriseTimeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudListEnterpriseTimeRequest
	GetResourceOwnerId() *int64
	SetType(v int64) *CloudListEnterpriseTimeRequest
	GetType() *int64
}

type CloudListEnterpriseTimeRequest struct {
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 时间条件名称；同一呼叫中心下不能重复
	//
	// example:
	//
	// test-name2
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 时间条件优先级；同一呼叫中心下不能重复、 值从1开始，值越小优先级越高
	//
	// example:
	//
	// 1
	Priority             *int64  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 时间条件类型；1:按照星期配置 2:按照固定时间配置
	//
	// example:
	//
	// 1
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CloudListEnterpriseTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudListEnterpriseTimeRequest) GoString() string {
	return s.String()
}

func (s *CloudListEnterpriseTimeRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudListEnterpriseTimeRequest) GetName() *string {
	return s.Name
}

func (s *CloudListEnterpriseTimeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudListEnterpriseTimeRequest) GetPriority() *int64 {
	return s.Priority
}

func (s *CloudListEnterpriseTimeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudListEnterpriseTimeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudListEnterpriseTimeRequest) GetType() *int64 {
	return s.Type
}

func (s *CloudListEnterpriseTimeRequest) SetEnterpriseId(v int64) *CloudListEnterpriseTimeRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudListEnterpriseTimeRequest) SetName(v string) *CloudListEnterpriseTimeRequest {
	s.Name = &v
	return s
}

func (s *CloudListEnterpriseTimeRequest) SetOwnerId(v int64) *CloudListEnterpriseTimeRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudListEnterpriseTimeRequest) SetPriority(v int64) *CloudListEnterpriseTimeRequest {
	s.Priority = &v
	return s
}

func (s *CloudListEnterpriseTimeRequest) SetResourceOwnerAccount(v string) *CloudListEnterpriseTimeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudListEnterpriseTimeRequest) SetResourceOwnerId(v int64) *CloudListEnterpriseTimeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudListEnterpriseTimeRequest) SetType(v int64) *CloudListEnterpriseTimeRequest {
	s.Type = &v
	return s
}

func (s *CloudListEnterpriseTimeRequest) Validate() error {
	return dara.Validate(s)
}
