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
	SetResourceGroupId(v string) *DescribeActiveOperationTaskTypeRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeActiveOperationTaskTypeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeActiveOperationTaskTypeRequest
	GetResourceOwnerId() *int64
}

type DescribeActiveOperationTaskTypeRequest struct {
	// Specifies whether to return historical O\\&M tasks. Valid values:
	//
	// 	- **0*	- (default): The system returns only pending O\\&M tasks.
	//
	// 	- **1**: The system returns historical O\\&M tasks.
	//
	// example:
	//
	// 0
	IsHistory    *int32  `json:"IsHistory,omitempty" xml:"IsHistory,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The resource group ID. You can call the [DescribeSecurityGroupConfiguration](https://help.aliyun.com/document_detail/146130.html) operation to query the resource group ID.
	//
	// example:
	//
	// rg-acfmyiu4ekp****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
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

func (s *DescribeActiveOperationTaskTypeRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeActiveOperationTaskTypeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeActiveOperationTaskTypeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
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

func (s *DescribeActiveOperationTaskTypeRequest) SetResourceGroupId(v string) *DescribeActiveOperationTaskTypeRequest {
	s.ResourceGroupId = &v
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

func (s *DescribeActiveOperationTaskTypeRequest) Validate() error {
	return dara.Validate(s)
}
