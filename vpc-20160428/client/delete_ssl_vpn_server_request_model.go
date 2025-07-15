// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSslVpnServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteSslVpnServerRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *DeleteSslVpnServerRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteSslVpnServerRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteSslVpnServerRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteSslVpnServerRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteSslVpnServerRequest
	GetResourceOwnerId() *int64
	SetSslVpnServerId(v string) *DeleteSslVpnServerRequest
	GetSslVpnServerId() *string
}

type DeleteSslVpnServerRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-0016e04115b
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the SSL server.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SSL server.
	//
	// This parameter is required.
	//
	// example:
	//
	// vss-bp18q7hzj6largv4v****
	SslVpnServerId *string `json:"SslVpnServerId,omitempty" xml:"SslVpnServerId,omitempty"`
}

func (s DeleteSslVpnServerRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSslVpnServerRequest) GoString() string {
	return s.String()
}

func (s *DeleteSslVpnServerRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteSslVpnServerRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteSslVpnServerRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteSslVpnServerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteSslVpnServerRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteSslVpnServerRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteSslVpnServerRequest) GetSslVpnServerId() *string {
	return s.SslVpnServerId
}

func (s *DeleteSslVpnServerRequest) SetClientToken(v string) *DeleteSslVpnServerRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteSslVpnServerRequest) SetOwnerAccount(v string) *DeleteSslVpnServerRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteSslVpnServerRequest) SetOwnerId(v int64) *DeleteSslVpnServerRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSslVpnServerRequest) SetRegionId(v string) *DeleteSslVpnServerRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSslVpnServerRequest) SetResourceOwnerAccount(v string) *DeleteSslVpnServerRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteSslVpnServerRequest) SetResourceOwnerId(v int64) *DeleteSslVpnServerRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteSslVpnServerRequest) SetSslVpnServerId(v string) *DeleteSslVpnServerRequest {
	s.SslVpnServerId = &v
	return s
}

func (s *DeleteSslVpnServerRequest) Validate() error {
	return dara.Validate(s)
}
