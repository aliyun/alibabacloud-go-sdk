// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDBClusterFromGDNRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *RemoveDBClusterFromGDNRequest
	GetDBClusterId() *string
	SetForce(v bool) *RemoveDBClusterFromGDNRequest
	GetForce() *bool
	SetGDNId(v string) *RemoveDBClusterFromGDNRequest
	GetGDNId() *string
	SetOwnerAccount(v string) *RemoveDBClusterFromGDNRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RemoveDBClusterFromGDNRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *RemoveDBClusterFromGDNRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RemoveDBClusterFromGDNRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *RemoveDBClusterFromGDNRequest
	GetSecurityToken() *string
}

type RemoveDBClusterFromGDNRequest struct {
	// The ID of the cluster in the GDN.
	//
	// >  You can call the [DescribeGlobalDatabaseNetwork](https://help.aliyun.com/document_detail/264580.html) operation to view the ID of the cluster in the GDN.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-wz9fb5nn44u1d****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Force       *bool   `json:"Force,omitempty" xml:"Force,omitempty"`
	// The ID of the GDN.
	//
	// This parameter is required.
	//
	// example:
	//
	// gdn-bp1fttxsrmv*****
	GDNId                *string `json:"GDNId,omitempty" xml:"GDNId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s RemoveDBClusterFromGDNRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveDBClusterFromGDNRequest) GoString() string {
	return s.String()
}

func (s *RemoveDBClusterFromGDNRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *RemoveDBClusterFromGDNRequest) GetForce() *bool {
	return s.Force
}

func (s *RemoveDBClusterFromGDNRequest) GetGDNId() *string {
	return s.GDNId
}

func (s *RemoveDBClusterFromGDNRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RemoveDBClusterFromGDNRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RemoveDBClusterFromGDNRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RemoveDBClusterFromGDNRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RemoveDBClusterFromGDNRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *RemoveDBClusterFromGDNRequest) SetDBClusterId(v string) *RemoveDBClusterFromGDNRequest {
	s.DBClusterId = &v
	return s
}

func (s *RemoveDBClusterFromGDNRequest) SetForce(v bool) *RemoveDBClusterFromGDNRequest {
	s.Force = &v
	return s
}

func (s *RemoveDBClusterFromGDNRequest) SetGDNId(v string) *RemoveDBClusterFromGDNRequest {
	s.GDNId = &v
	return s
}

func (s *RemoveDBClusterFromGDNRequest) SetOwnerAccount(v string) *RemoveDBClusterFromGDNRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RemoveDBClusterFromGDNRequest) SetOwnerId(v int64) *RemoveDBClusterFromGDNRequest {
	s.OwnerId = &v
	return s
}

func (s *RemoveDBClusterFromGDNRequest) SetResourceOwnerAccount(v string) *RemoveDBClusterFromGDNRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RemoveDBClusterFromGDNRequest) SetResourceOwnerId(v int64) *RemoveDBClusterFromGDNRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RemoveDBClusterFromGDNRequest) SetSecurityToken(v string) *RemoveDBClusterFromGDNRequest {
	s.SecurityToken = &v
	return s
}

func (s *RemoveDBClusterFromGDNRequest) Validate() error {
	return dara.Validate(s)
}
