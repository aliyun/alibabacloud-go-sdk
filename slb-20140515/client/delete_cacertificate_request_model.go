// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCACertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCACertificateId(v string) *DeleteCACertificateRequest
	GetCACertificateId() *string
	SetOwnerAccount(v string) *DeleteCACertificateRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteCACertificateRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteCACertificateRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteCACertificateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteCACertificateRequest
	GetResourceOwnerId() *int64
}

type DeleteCACertificateRequest struct {
	// The CA certificate ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123157908xxxxxxx_15c73d77203_-986300114_-2110544xxx
	CACertificateId *string `json:"CACertificateId,omitempty" xml:"CACertificateId,omitempty"`
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region of the CA certificates.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/27584.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteCACertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCACertificateRequest) GoString() string {
	return s.String()
}

func (s *DeleteCACertificateRequest) GetCACertificateId() *string {
	return s.CACertificateId
}

func (s *DeleteCACertificateRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteCACertificateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteCACertificateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteCACertificateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteCACertificateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteCACertificateRequest) SetCACertificateId(v string) *DeleteCACertificateRequest {
	s.CACertificateId = &v
	return s
}

func (s *DeleteCACertificateRequest) SetOwnerAccount(v string) *DeleteCACertificateRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteCACertificateRequest) SetOwnerId(v int64) *DeleteCACertificateRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteCACertificateRequest) SetRegionId(v string) *DeleteCACertificateRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteCACertificateRequest) SetResourceOwnerAccount(v string) *DeleteCACertificateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteCACertificateRequest) SetResourceOwnerId(v int64) *DeleteCACertificateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteCACertificateRequest) Validate() error {
	return dara.Validate(s)
}
