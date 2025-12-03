// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServerCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DeleteServerCertificateRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteServerCertificateRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteServerCertificateRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteServerCertificateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteServerCertificateRequest
	GetResourceOwnerId() *int64
	SetServerCertificateId(v string) *DeleteServerCertificateRequest
	GetServerCertificateId() *string
}

type DeleteServerCertificateRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the Classic Load Balancer (CLB) instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/27584.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The server certificate ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123157xxxxxxx_166f8204689_1714763408_709981430
	ServerCertificateId *string `json:"ServerCertificateId,omitempty" xml:"ServerCertificateId,omitempty"`
}

func (s DeleteServerCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteServerCertificateRequest) GoString() string {
	return s.String()
}

func (s *DeleteServerCertificateRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteServerCertificateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteServerCertificateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteServerCertificateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteServerCertificateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteServerCertificateRequest) GetServerCertificateId() *string {
	return s.ServerCertificateId
}

func (s *DeleteServerCertificateRequest) SetOwnerAccount(v string) *DeleteServerCertificateRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteServerCertificateRequest) SetOwnerId(v int64) *DeleteServerCertificateRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteServerCertificateRequest) SetRegionId(v string) *DeleteServerCertificateRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteServerCertificateRequest) SetResourceOwnerAccount(v string) *DeleteServerCertificateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteServerCertificateRequest) SetResourceOwnerId(v int64) *DeleteServerCertificateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteServerCertificateRequest) SetServerCertificateId(v string) *DeleteServerCertificateRequest {
	s.ServerCertificateId = &v
	return s
}

func (s *DeleteServerCertificateRequest) Validate() error {
	return dara.Validate(s)
}
