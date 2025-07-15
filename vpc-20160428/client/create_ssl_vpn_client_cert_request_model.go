// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSslVpnClientCertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateSslVpnClientCertRequest
	GetClientToken() *string
	SetName(v string) *CreateSslVpnClientCertRequest
	GetName() *string
	SetOwnerAccount(v string) *CreateSslVpnClientCertRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateSslVpnClientCertRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateSslVpnClientCertRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateSslVpnClientCertRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateSslVpnClientCertRequest
	GetResourceOwnerId() *int64
	SetSslVpnServerId(v string) *CreateSslVpnClientCertRequest
	GetSslVpnServerId() *string
}

type CreateSslVpnClientCertRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// >  If you do not set this parameter, **ClientToken*	- is set to the value of **RequestId**. The value of **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-0016e04115b
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The name of the SSL client certificate.
	//
	// The name must be 1 to 100 characters in length, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// SslVpnClientCert1
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the VPN gateway is created.
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
	// vss-m5et0q3iy1qex328w****
	SslVpnServerId *string `json:"SslVpnServerId,omitempty" xml:"SslVpnServerId,omitempty"`
}

func (s CreateSslVpnClientCertRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSslVpnClientCertRequest) GoString() string {
	return s.String()
}

func (s *CreateSslVpnClientCertRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateSslVpnClientCertRequest) GetName() *string {
	return s.Name
}

func (s *CreateSslVpnClientCertRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateSslVpnClientCertRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateSslVpnClientCertRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSslVpnClientCertRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateSslVpnClientCertRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateSslVpnClientCertRequest) GetSslVpnServerId() *string {
	return s.SslVpnServerId
}

func (s *CreateSslVpnClientCertRequest) SetClientToken(v string) *CreateSslVpnClientCertRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSslVpnClientCertRequest) SetName(v string) *CreateSslVpnClientCertRequest {
	s.Name = &v
	return s
}

func (s *CreateSslVpnClientCertRequest) SetOwnerAccount(v string) *CreateSslVpnClientCertRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateSslVpnClientCertRequest) SetOwnerId(v int64) *CreateSslVpnClientCertRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSslVpnClientCertRequest) SetRegionId(v string) *CreateSslVpnClientCertRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSslVpnClientCertRequest) SetResourceOwnerAccount(v string) *CreateSslVpnClientCertRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateSslVpnClientCertRequest) SetResourceOwnerId(v int64) *CreateSslVpnClientCertRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateSslVpnClientCertRequest) SetSslVpnServerId(v string) *CreateSslVpnClientCertRequest {
	s.SslVpnServerId = &v
	return s
}

func (s *CreateSslVpnClientCertRequest) Validate() error {
	return dara.Validate(s)
}
