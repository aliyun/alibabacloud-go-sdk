// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityIpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterIPArrayAttribute(v string) *ModifySecurityIpsRequest
	GetDBClusterIPArrayAttribute() *string
	SetDBClusterIPArrayName(v string) *ModifySecurityIpsRequest
	GetDBClusterIPArrayName() *string
	SetDBClusterId(v string) *ModifySecurityIpsRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *ModifySecurityIpsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifySecurityIpsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifySecurityIpsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySecurityIpsRequest
	GetResourceOwnerId() *int64
	SetSecurityIps(v string) *ModifySecurityIpsRequest
	GetSecurityIps() *string
}

type ModifySecurityIpsRequest struct {
	// The property of the IP address whitelist group.
	//
	// example:
	//
	// hidden
	DBClusterIPArrayAttribute *string `json:"DBClusterIPArrayAttribute,omitempty" xml:"DBClusterIPArrayAttribute,omitempty"`
	// The name of the IP address whitelist group.
	//
	// example:
	//
	// default
	DBClusterIPArrayName *string `json:"DBClusterIPArrayName,omitempty" xml:"DBClusterIPArrayName,omitempty"`
	// The ID of the PolarDB cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-******************
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The IP addresses or CIDR blocks in the whitelist. Separate multiple values with commas.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.23.12.24
	SecurityIps *string `json:"SecurityIps,omitempty" xml:"SecurityIps,omitempty"`
}

func (s ModifySecurityIpsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityIpsRequest) GoString() string {
	return s.String()
}

func (s *ModifySecurityIpsRequest) GetDBClusterIPArrayAttribute() *string {
	return s.DBClusterIPArrayAttribute
}

func (s *ModifySecurityIpsRequest) GetDBClusterIPArrayName() *string {
	return s.DBClusterIPArrayName
}

func (s *ModifySecurityIpsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifySecurityIpsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifySecurityIpsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySecurityIpsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySecurityIpsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySecurityIpsRequest) GetSecurityIps() *string {
	return s.SecurityIps
}

func (s *ModifySecurityIpsRequest) SetDBClusterIPArrayAttribute(v string) *ModifySecurityIpsRequest {
	s.DBClusterIPArrayAttribute = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetDBClusterIPArrayName(v string) *ModifySecurityIpsRequest {
	s.DBClusterIPArrayName = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetDBClusterId(v string) *ModifySecurityIpsRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetOwnerAccount(v string) *ModifySecurityIpsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetOwnerId(v int64) *ModifySecurityIpsRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetResourceOwnerAccount(v string) *ModifySecurityIpsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetResourceOwnerId(v int64) *ModifySecurityIpsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetSecurityIps(v string) *ModifySecurityIpsRequest {
	s.SecurityIps = &v
	return s
}

func (s *ModifySecurityIpsRequest) Validate() error {
	return dara.Validate(s)
}
