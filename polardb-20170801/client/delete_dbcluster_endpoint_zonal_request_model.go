// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBClusterEndpointZonalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteDBClusterEndpointZonalRequest
	GetDBClusterId() *string
	SetDBEndpointId(v string) *DeleteDBClusterEndpointZonalRequest
	GetDBEndpointId() *string
	SetOwnerAccount(v string) *DeleteDBClusterEndpointZonalRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteDBClusterEndpointZonalRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteDBClusterEndpointZonalRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteDBClusterEndpointZonalRequest
	GetResourceOwnerId() *int64
}

type DeleteDBClusterEndpointZonalRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pe-******************
	DBEndpointId         *string `json:"DBEndpointId,omitempty" xml:"DBEndpointId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteDBClusterEndpointZonalRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBClusterEndpointZonalRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBClusterEndpointZonalRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteDBClusterEndpointZonalRequest) GetDBEndpointId() *string {
	return s.DBEndpointId
}

func (s *DeleteDBClusterEndpointZonalRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteDBClusterEndpointZonalRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteDBClusterEndpointZonalRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteDBClusterEndpointZonalRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteDBClusterEndpointZonalRequest) SetDBClusterId(v string) *DeleteDBClusterEndpointZonalRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteDBClusterEndpointZonalRequest) SetDBEndpointId(v string) *DeleteDBClusterEndpointZonalRequest {
	s.DBEndpointId = &v
	return s
}

func (s *DeleteDBClusterEndpointZonalRequest) SetOwnerAccount(v string) *DeleteDBClusterEndpointZonalRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDBClusterEndpointZonalRequest) SetOwnerId(v int64) *DeleteDBClusterEndpointZonalRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDBClusterEndpointZonalRequest) SetResourceOwnerAccount(v string) *DeleteDBClusterEndpointZonalRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDBClusterEndpointZonalRequest) SetResourceOwnerId(v int64) *DeleteDBClusterEndpointZonalRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteDBClusterEndpointZonalRequest) Validate() error {
	return dara.Validate(s)
}
