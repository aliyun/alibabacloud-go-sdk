// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshDBClusterStorageUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *RefreshDBClusterStorageUsageRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RefreshDBClusterStorageUsageRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *RefreshDBClusterStorageUsageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RefreshDBClusterStorageUsageRequest
	GetResourceOwnerId() *int64
	SetSyncRealTime(v bool) *RefreshDBClusterStorageUsageRequest
	GetSyncRealTime() *bool
}

type RefreshDBClusterStorageUsageRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to enable real-time synchronization. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**: T + 1
	//
	// example:
	//
	// true
	SyncRealTime *bool `json:"SyncRealTime,omitempty" xml:"SyncRealTime,omitempty"`
}

func (s RefreshDBClusterStorageUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s RefreshDBClusterStorageUsageRequest) GoString() string {
	return s.String()
}

func (s *RefreshDBClusterStorageUsageRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RefreshDBClusterStorageUsageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RefreshDBClusterStorageUsageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RefreshDBClusterStorageUsageRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RefreshDBClusterStorageUsageRequest) GetSyncRealTime() *bool {
	return s.SyncRealTime
}

func (s *RefreshDBClusterStorageUsageRequest) SetOwnerAccount(v string) *RefreshDBClusterStorageUsageRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RefreshDBClusterStorageUsageRequest) SetOwnerId(v int64) *RefreshDBClusterStorageUsageRequest {
	s.OwnerId = &v
	return s
}

func (s *RefreshDBClusterStorageUsageRequest) SetResourceOwnerAccount(v string) *RefreshDBClusterStorageUsageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RefreshDBClusterStorageUsageRequest) SetResourceOwnerId(v int64) *RefreshDBClusterStorageUsageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RefreshDBClusterStorageUsageRequest) SetSyncRealTime(v bool) *RefreshDBClusterStorageUsageRequest {
	s.SyncRealTime = &v
	return s
}

func (s *RefreshDBClusterStorageUsageRequest) Validate() error {
	return dara.Validate(s)
}
