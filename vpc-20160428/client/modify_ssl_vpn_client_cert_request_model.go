// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySslVpnClientCertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifySslVpnClientCertRequest
	GetClientToken() *string
	SetName(v string) *ModifySslVpnClientCertRequest
	GetName() *string
	SetOwnerAccount(v string) *ModifySslVpnClientCertRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifySslVpnClientCertRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifySslVpnClientCertRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifySslVpnClientCertRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySslVpnClientCertRequest
	GetResourceOwnerId() *int64
	SetSslVpnClientCertId(v string) *ModifySslVpnClientCertRequest
	GetSslVpnClientCertId() *string
}

type ModifySslVpnClientCertRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-0016e04115b
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The new name of the SSL client certificate. This parameter cannot be left empty.
	//
	// The name must be 1 to 100 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// cert2
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	// vsc-bp1n8wcf134yl0osrc****
	SslVpnClientCertId *string `json:"SslVpnClientCertId,omitempty" xml:"SslVpnClientCertId,omitempty"`
}

func (s ModifySslVpnClientCertRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySslVpnClientCertRequest) GoString() string {
	return s.String()
}

func (s *ModifySslVpnClientCertRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifySslVpnClientCertRequest) GetName() *string {
	return s.Name
}

func (s *ModifySslVpnClientCertRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifySslVpnClientCertRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySslVpnClientCertRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySslVpnClientCertRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySslVpnClientCertRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySslVpnClientCertRequest) GetSslVpnClientCertId() *string {
	return s.SslVpnClientCertId
}

func (s *ModifySslVpnClientCertRequest) SetClientToken(v string) *ModifySslVpnClientCertRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifySslVpnClientCertRequest) SetName(v string) *ModifySslVpnClientCertRequest {
	s.Name = &v
	return s
}

func (s *ModifySslVpnClientCertRequest) SetOwnerAccount(v string) *ModifySslVpnClientCertRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySslVpnClientCertRequest) SetOwnerId(v int64) *ModifySslVpnClientCertRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySslVpnClientCertRequest) SetRegionId(v string) *ModifySslVpnClientCertRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySslVpnClientCertRequest) SetResourceOwnerAccount(v string) *ModifySslVpnClientCertRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySslVpnClientCertRequest) SetResourceOwnerId(v int64) *ModifySslVpnClientCertRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySslVpnClientCertRequest) SetSslVpnClientCertId(v string) *ModifySslVpnClientCertRequest {
	s.SslVpnClientCertId = &v
	return s
}

func (s *ModifySslVpnClientCertRequest) Validate() error {
	return dara.Validate(s)
}
