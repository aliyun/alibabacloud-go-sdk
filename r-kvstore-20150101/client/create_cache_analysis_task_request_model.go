// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCacheAnalysisTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateCacheAnalysisTaskRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *CreateCacheAnalysisTaskRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateCacheAnalysisTaskRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateCacheAnalysisTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateCacheAnalysisTaskRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *CreateCacheAnalysisTaskRequest
	GetSecurityToken() *string
}

type CreateCacheAnalysisTaskRequest struct {
	// The ID of the instance. You can call the [DescribeInstances](https://help.aliyun.com/document_detail/473778.html) operation to query the ID of the instance.
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
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s CreateCacheAnalysisTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCacheAnalysisTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateCacheAnalysisTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateCacheAnalysisTaskRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateCacheAnalysisTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateCacheAnalysisTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateCacheAnalysisTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateCacheAnalysisTaskRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreateCacheAnalysisTaskRequest) SetInstanceId(v string) *CreateCacheAnalysisTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCacheAnalysisTaskRequest) SetOwnerAccount(v string) *CreateCacheAnalysisTaskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateCacheAnalysisTaskRequest) SetOwnerId(v int64) *CreateCacheAnalysisTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateCacheAnalysisTaskRequest) SetResourceOwnerAccount(v string) *CreateCacheAnalysisTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateCacheAnalysisTaskRequest) SetResourceOwnerId(v int64) *CreateCacheAnalysisTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateCacheAnalysisTaskRequest) SetSecurityToken(v string) *CreateCacheAnalysisTaskRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateCacheAnalysisTaskRequest) Validate() error {
	return dara.Validate(s)
}
