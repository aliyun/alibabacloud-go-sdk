// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeActiveOperationTaskTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsHistory(v int32) *DescribeActiveOperationTaskTypeRequest
	GetIsHistory() *int32
	SetOwnerAccount(v string) *DescribeActiveOperationTaskTypeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeActiveOperationTaskTypeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeActiveOperationTaskTypeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeActiveOperationTaskTypeRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeActiveOperationTaskTypeRequest
	GetSecurityToken() *string
}

type DescribeActiveOperationTaskTypeRequest struct {
	// example:
	//
	// 0
	IsHistory            *int32  `json:"IsHistory,omitempty" xml:"IsHistory,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeActiveOperationTaskTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveOperationTaskTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTaskTypeRequest) GetIsHistory() *int32 {
	return s.IsHistory
}

func (s *DescribeActiveOperationTaskTypeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeActiveOperationTaskTypeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeActiveOperationTaskTypeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeActiveOperationTaskTypeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeActiveOperationTaskTypeRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeActiveOperationTaskTypeRequest) SetIsHistory(v int32) *DescribeActiveOperationTaskTypeRequest {
	s.IsHistory = &v
	return s
}

func (s *DescribeActiveOperationTaskTypeRequest) SetOwnerAccount(v string) *DescribeActiveOperationTaskTypeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeActiveOperationTaskTypeRequest) SetOwnerId(v int64) *DescribeActiveOperationTaskTypeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeActiveOperationTaskTypeRequest) SetResourceOwnerAccount(v string) *DescribeActiveOperationTaskTypeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeActiveOperationTaskTypeRequest) SetResourceOwnerId(v int64) *DescribeActiveOperationTaskTypeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeActiveOperationTaskTypeRequest) SetSecurityToken(v string) *DescribeActiveOperationTaskTypeRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeActiveOperationTaskTypeRequest) Validate() error {
	return dara.Validate(s)
}
