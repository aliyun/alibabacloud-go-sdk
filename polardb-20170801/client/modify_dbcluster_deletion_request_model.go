// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterDeletionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyDBClusterDeletionRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *ModifyDBClusterDeletionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBClusterDeletionRequest
	GetOwnerId() *int64
	SetProtection(v bool) *ModifyDBClusterDeletionRequest
	GetProtection() *bool
	SetResourceOwnerAccount(v string) *ModifyDBClusterDeletionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBClusterDeletionRequest
	GetResourceOwnerId() *int64
}

type ModifyDBClusterDeletionRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to.obtain the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-bp1313h70cd5m****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies whether to enable the cluster lock feature. Default value: false. Valid values:
	//
	// 	- **true**: enables the cluster lock feature. If you enable the cluster lock feature, you cannot directly release the cluster. You must disable the cluster lock feature before you can release the cluster.
	//
	// 	- **false**: disables the cluster lock feature.
	//
	// example:
	//
	// true
	Protection           *bool   `json:"Protection,omitempty" xml:"Protection,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBClusterDeletionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterDeletionRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterDeletionRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterDeletionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBClusterDeletionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBClusterDeletionRequest) GetProtection() *bool {
	return s.Protection
}

func (s *ModifyDBClusterDeletionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBClusterDeletionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBClusterDeletionRequest) SetDBClusterId(v string) *ModifyDBClusterDeletionRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterDeletionRequest) SetOwnerAccount(v string) *ModifyDBClusterDeletionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterDeletionRequest) SetOwnerId(v int64) *ModifyDBClusterDeletionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterDeletionRequest) SetProtection(v bool) *ModifyDBClusterDeletionRequest {
	s.Protection = &v
	return s
}

func (s *ModifyDBClusterDeletionRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterDeletionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterDeletionRequest) SetResourceOwnerId(v int64) *ModifyDBClusterDeletionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterDeletionRequest) Validate() error {
	return dara.Validate(s)
}
