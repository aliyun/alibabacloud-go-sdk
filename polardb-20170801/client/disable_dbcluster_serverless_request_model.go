// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDBClusterServerlessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DisableDBClusterServerlessRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DisableDBClusterServerlessRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DisableDBClusterServerlessRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DisableDBClusterServerlessRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DisableDBClusterServerlessRequest
	GetResourceOwnerId() *int64
}

type DisableDBClusterServerlessRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-***************
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DisableDBClusterServerlessRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableDBClusterServerlessRequest) GoString() string {
	return s.String()
}

func (s *DisableDBClusterServerlessRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DisableDBClusterServerlessRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DisableDBClusterServerlessRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DisableDBClusterServerlessRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DisableDBClusterServerlessRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DisableDBClusterServerlessRequest) SetDBClusterId(v string) *DisableDBClusterServerlessRequest {
	s.DBClusterId = &v
	return s
}

func (s *DisableDBClusterServerlessRequest) SetOwnerAccount(v string) *DisableDBClusterServerlessRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DisableDBClusterServerlessRequest) SetOwnerId(v int64) *DisableDBClusterServerlessRequest {
	s.OwnerId = &v
	return s
}

func (s *DisableDBClusterServerlessRequest) SetResourceOwnerAccount(v string) *DisableDBClusterServerlessRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DisableDBClusterServerlessRequest) SetResourceOwnerId(v int64) *DisableDBClusterServerlessRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DisableDBClusterServerlessRequest) Validate() error {
	return dara.Validate(s)
}
