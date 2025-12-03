// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCACertificateNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCACertificateId(v string) *SetCACertificateNameRequest
	GetCACertificateId() *string
	SetCACertificateName(v string) *SetCACertificateNameRequest
	GetCACertificateName() *string
	SetOwnerAccount(v string) *SetCACertificateNameRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SetCACertificateNameRequest
	GetOwnerId() *int64
	SetRegionId(v string) *SetCACertificateNameRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *SetCACertificateNameRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SetCACertificateNameRequest
	GetResourceOwnerId() *int64
}

type SetCACertificateNameRequest struct {
	// The ID of the CA certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 139a0******-cn-east-hangzhou-01
	CACertificateId *string `json:"CACertificateId,omitempty" xml:"CACertificateId,omitempty"`
	// The CA certificate name.
	//
	// The name must be 1 to 80 character in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
	//
	// This parameter is required.
	//
	// example:
	//
	// mycacert02
	CACertificateName *string `json:"CACertificateName,omitempty" xml:"CACertificateName,omitempty"`
	OwnerAccount      *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region of the CA certificate.
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

func (s SetCACertificateNameRequest) String() string {
	return dara.Prettify(s)
}

func (s SetCACertificateNameRequest) GoString() string {
	return s.String()
}

func (s *SetCACertificateNameRequest) GetCACertificateId() *string {
	return s.CACertificateId
}

func (s *SetCACertificateNameRequest) GetCACertificateName() *string {
	return s.CACertificateName
}

func (s *SetCACertificateNameRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SetCACertificateNameRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetCACertificateNameRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetCACertificateNameRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SetCACertificateNameRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SetCACertificateNameRequest) SetCACertificateId(v string) *SetCACertificateNameRequest {
	s.CACertificateId = &v
	return s
}

func (s *SetCACertificateNameRequest) SetCACertificateName(v string) *SetCACertificateNameRequest {
	s.CACertificateName = &v
	return s
}

func (s *SetCACertificateNameRequest) SetOwnerAccount(v string) *SetCACertificateNameRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetCACertificateNameRequest) SetOwnerId(v int64) *SetCACertificateNameRequest {
	s.OwnerId = &v
	return s
}

func (s *SetCACertificateNameRequest) SetRegionId(v string) *SetCACertificateNameRequest {
	s.RegionId = &v
	return s
}

func (s *SetCACertificateNameRequest) SetResourceOwnerAccount(v string) *SetCACertificateNameRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetCACertificateNameRequest) SetResourceOwnerId(v int64) *SetCACertificateNameRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetCACertificateNameRequest) Validate() error {
	return dara.Validate(s)
}
