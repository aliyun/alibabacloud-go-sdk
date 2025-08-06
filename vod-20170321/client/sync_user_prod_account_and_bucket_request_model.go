// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncUserProdAccountAndBucketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBindOssNotification(v bool) *SyncUserProdAccountAndBucketRequest
	GetBindOssNotification() *bool
	SetOwnerId(v int64) *SyncUserProdAccountAndBucketRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *SyncUserProdAccountAndBucketRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SyncUserProdAccountAndBucketRequest
	GetResourceOwnerId() *int64
	SetResourceRealOwnerId(v int64) *SyncUserProdAccountAndBucketRequest
	GetResourceRealOwnerId() *int64
}

type SyncUserProdAccountAndBucketRequest struct {
	BindOssNotification  *bool   `json:"BindOssNotification,omitempty" xml:"BindOssNotification,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceRealOwnerId  *int64  `json:"ResourceRealOwnerId,omitempty" xml:"ResourceRealOwnerId,omitempty"`
}

func (s SyncUserProdAccountAndBucketRequest) String() string {
	return dara.Prettify(s)
}

func (s SyncUserProdAccountAndBucketRequest) GoString() string {
	return s.String()
}

func (s *SyncUserProdAccountAndBucketRequest) GetBindOssNotification() *bool {
	return s.BindOssNotification
}

func (s *SyncUserProdAccountAndBucketRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SyncUserProdAccountAndBucketRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SyncUserProdAccountAndBucketRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SyncUserProdAccountAndBucketRequest) GetResourceRealOwnerId() *int64 {
	return s.ResourceRealOwnerId
}

func (s *SyncUserProdAccountAndBucketRequest) SetBindOssNotification(v bool) *SyncUserProdAccountAndBucketRequest {
	s.BindOssNotification = &v
	return s
}

func (s *SyncUserProdAccountAndBucketRequest) SetOwnerId(v int64) *SyncUserProdAccountAndBucketRequest {
	s.OwnerId = &v
	return s
}

func (s *SyncUserProdAccountAndBucketRequest) SetResourceOwnerAccount(v string) *SyncUserProdAccountAndBucketRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SyncUserProdAccountAndBucketRequest) SetResourceOwnerId(v int64) *SyncUserProdAccountAndBucketRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SyncUserProdAccountAndBucketRequest) SetResourceRealOwnerId(v int64) *SyncUserProdAccountAndBucketRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *SyncUserProdAccountAndBucketRequest) Validate() error {
	return dara.Validate(s)
}
