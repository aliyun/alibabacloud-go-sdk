// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSslVpnClientCertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteSslVpnClientCertRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *DeleteSslVpnClientCertRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteSslVpnClientCertRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteSslVpnClientCertRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteSslVpnClientCertRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteSslVpnClientCertRequest
	GetResourceOwnerId() *int64
	SetSslVpnClientCertId(v string) *DeleteSslVpnClientCertRequest
	GetSslVpnClientCertId() *string
}

type DeleteSslVpnClientCertRequest struct {
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
	// The ID of the region where the SSL client certificate is created.
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
	// The ID of the SSL client certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsc-bp1n8wcf134yl0osr****
	SslVpnClientCertId *string `json:"SslVpnClientCertId,omitempty" xml:"SslVpnClientCertId,omitempty"`
}

func (s DeleteSslVpnClientCertRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSslVpnClientCertRequest) GoString() string {
	return s.String()
}

func (s *DeleteSslVpnClientCertRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteSslVpnClientCertRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteSslVpnClientCertRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteSslVpnClientCertRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteSslVpnClientCertRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteSslVpnClientCertRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteSslVpnClientCertRequest) GetSslVpnClientCertId() *string {
	return s.SslVpnClientCertId
}

func (s *DeleteSslVpnClientCertRequest) SetClientToken(v string) *DeleteSslVpnClientCertRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteSslVpnClientCertRequest) SetOwnerAccount(v string) *DeleteSslVpnClientCertRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteSslVpnClientCertRequest) SetOwnerId(v int64) *DeleteSslVpnClientCertRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSslVpnClientCertRequest) SetRegionId(v string) *DeleteSslVpnClientCertRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSslVpnClientCertRequest) SetResourceOwnerAccount(v string) *DeleteSslVpnClientCertRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteSslVpnClientCertRequest) SetResourceOwnerId(v int64) *DeleteSslVpnClientCertRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteSslVpnClientCertRequest) SetSslVpnClientCertId(v string) *DeleteSslVpnClientCertRequest {
	s.SslVpnClientCertId = &v
	return s
}

func (s *DeleteSslVpnClientCertRequest) Validate() error {
	return dara.Validate(s)
}
