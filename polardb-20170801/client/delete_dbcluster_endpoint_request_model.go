// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBClusterEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteDBClusterEndpointRequest
	GetDBClusterId() *string
	SetDBEndpointId(v string) *DeleteDBClusterEndpointRequest
	GetDBEndpointId() *string
	SetOwnerAccount(v string) *DeleteDBClusterEndpointRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteDBClusterEndpointRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteDBClusterEndpointRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteDBClusterEndpointRequest
	GetResourceOwnerId() *int64
}

type DeleteDBClusterEndpointRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-******************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the custom cluster endpoint.
	//
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

func (s DeleteDBClusterEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBClusterEndpointRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBClusterEndpointRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteDBClusterEndpointRequest) GetDBEndpointId() *string {
	return s.DBEndpointId
}

func (s *DeleteDBClusterEndpointRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteDBClusterEndpointRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteDBClusterEndpointRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteDBClusterEndpointRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteDBClusterEndpointRequest) SetDBClusterId(v string) *DeleteDBClusterEndpointRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteDBClusterEndpointRequest) SetDBEndpointId(v string) *DeleteDBClusterEndpointRequest {
	s.DBEndpointId = &v
	return s
}

func (s *DeleteDBClusterEndpointRequest) SetOwnerAccount(v string) *DeleteDBClusterEndpointRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDBClusterEndpointRequest) SetOwnerId(v int64) *DeleteDBClusterEndpointRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDBClusterEndpointRequest) SetResourceOwnerAccount(v string) *DeleteDBClusterEndpointRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDBClusterEndpointRequest) SetResourceOwnerId(v int64) *DeleteDBClusterEndpointRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteDBClusterEndpointRequest) Validate() error {
	return dara.Validate(s)
}
