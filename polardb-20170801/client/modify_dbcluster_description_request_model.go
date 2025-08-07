// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterDescription(v string) *ModifyDBClusterDescriptionRequest
	GetDBClusterDescription() *string
	SetDBClusterId(v string) *ModifyDBClusterDescriptionRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *ModifyDBClusterDescriptionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBClusterDescriptionRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyDBClusterDescriptionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBClusterDescriptionRequest
	GetResourceOwnerId() *int64
}

type ModifyDBClusterDescriptionRequest struct {
	// The name of the cluster. The name must meet the following requirements:
	//
	// 	- The name cannot start with `http://` or `https://`.
	//
	// 	- The name must be 2 to 256 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// ClusterDescriptionTest
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// The ID of the PolarDB cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBClusterDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterDescriptionRequest) GetDBClusterDescription() *string {
	return s.DBClusterDescription
}

func (s *ModifyDBClusterDescriptionRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterDescriptionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBClusterDescriptionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBClusterDescriptionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBClusterDescriptionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBClusterDescriptionRequest) SetDBClusterDescription(v string) *ModifyDBClusterDescriptionRequest {
	s.DBClusterDescription = &v
	return s
}

func (s *ModifyDBClusterDescriptionRequest) SetDBClusterId(v string) *ModifyDBClusterDescriptionRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterDescriptionRequest) SetOwnerAccount(v string) *ModifyDBClusterDescriptionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterDescriptionRequest) SetOwnerId(v int64) *ModifyDBClusterDescriptionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterDescriptionRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterDescriptionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterDescriptionRequest) SetResourceOwnerId(v int64) *ModifyDBClusterDescriptionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterDescriptionRequest) Validate() error {
	return dara.Validate(s)
}
