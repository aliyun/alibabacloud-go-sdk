// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseClusterPublicConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ReleaseClusterPublicConnectionRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *ReleaseClusterPublicConnectionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ReleaseClusterPublicConnectionRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ReleaseClusterPublicConnectionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ReleaseClusterPublicConnectionRequest
	GetResourceOwnerId() *int64
}

type ReleaseClusterPublicConnectionRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ReleaseClusterPublicConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseClusterPublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *ReleaseClusterPublicConnectionRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ReleaseClusterPublicConnectionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ReleaseClusterPublicConnectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ReleaseClusterPublicConnectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ReleaseClusterPublicConnectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ReleaseClusterPublicConnectionRequest) SetDBClusterId(v string) *ReleaseClusterPublicConnectionRequest {
	s.DBClusterId = &v
	return s
}

func (s *ReleaseClusterPublicConnectionRequest) SetOwnerAccount(v string) *ReleaseClusterPublicConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ReleaseClusterPublicConnectionRequest) SetOwnerId(v int64) *ReleaseClusterPublicConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *ReleaseClusterPublicConnectionRequest) SetResourceOwnerAccount(v string) *ReleaseClusterPublicConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReleaseClusterPublicConnectionRequest) SetResourceOwnerId(v int64) *ReleaseClusterPublicConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReleaseClusterPublicConnectionRequest) Validate() error {
	return dara.Validate(s)
}
