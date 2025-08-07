// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartDBLinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *RestartDBLinkRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *RestartDBLinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RestartDBLinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *RestartDBLinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RestartDBLinkRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *RestartDBLinkRequest
	GetSecurityToken() *string
}

type RestartDBLinkRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query the IDs of all clusters in an Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s RestartDBLinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartDBLinkRequest) GoString() string {
	return s.String()
}

func (s *RestartDBLinkRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *RestartDBLinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RestartDBLinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RestartDBLinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RestartDBLinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RestartDBLinkRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *RestartDBLinkRequest) SetDBClusterId(v string) *RestartDBLinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *RestartDBLinkRequest) SetOwnerAccount(v string) *RestartDBLinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RestartDBLinkRequest) SetOwnerId(v int64) *RestartDBLinkRequest {
	s.OwnerId = &v
	return s
}

func (s *RestartDBLinkRequest) SetResourceOwnerAccount(v string) *RestartDBLinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RestartDBLinkRequest) SetResourceOwnerId(v int64) *RestartDBLinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RestartDBLinkRequest) SetSecurityToken(v string) *RestartDBLinkRequest {
	s.SecurityToken = &v
	return s
}

func (s *RestartDBLinkRequest) Validate() error {
	return dara.Validate(s)
}
