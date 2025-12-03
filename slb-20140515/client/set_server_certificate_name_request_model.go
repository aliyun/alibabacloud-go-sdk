// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetServerCertificateNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *SetServerCertificateNameRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SetServerCertificateNameRequest
	GetOwnerId() *int64
	SetRegionId(v string) *SetServerCertificateNameRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *SetServerCertificateNameRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SetServerCertificateNameRequest
	GetResourceOwnerId() *int64
	SetServerCertificateId(v string) *SetServerCertificateNameRequest
	GetServerCertificateId() *string
	SetServerCertificateName(v string) *SetServerCertificateNameRequest
	GetServerCertificateName() *string
}

type SetServerCertificateNameRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the Classic Load Balancer (CLB) instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/2401682.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the server certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 139a0******-cn-east-hangzhou-01
	ServerCertificateId *string `json:"ServerCertificateId,omitempty" xml:"ServerCertificateId,omitempty"`
	// The name of the third-party server certificate that you want to upload. The name must be 1 to 80 characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), underscores (_), and asterisks (\\*).
	//
	// This parameter is required.
	//
	// example:
	//
	// mycert01
	ServerCertificateName *string `json:"ServerCertificateName,omitempty" xml:"ServerCertificateName,omitempty"`
}

func (s SetServerCertificateNameRequest) String() string {
	return dara.Prettify(s)
}

func (s SetServerCertificateNameRequest) GoString() string {
	return s.String()
}

func (s *SetServerCertificateNameRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SetServerCertificateNameRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetServerCertificateNameRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetServerCertificateNameRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SetServerCertificateNameRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SetServerCertificateNameRequest) GetServerCertificateId() *string {
	return s.ServerCertificateId
}

func (s *SetServerCertificateNameRequest) GetServerCertificateName() *string {
	return s.ServerCertificateName
}

func (s *SetServerCertificateNameRequest) SetOwnerAccount(v string) *SetServerCertificateNameRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetServerCertificateNameRequest) SetOwnerId(v int64) *SetServerCertificateNameRequest {
	s.OwnerId = &v
	return s
}

func (s *SetServerCertificateNameRequest) SetRegionId(v string) *SetServerCertificateNameRequest {
	s.RegionId = &v
	return s
}

func (s *SetServerCertificateNameRequest) SetResourceOwnerAccount(v string) *SetServerCertificateNameRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetServerCertificateNameRequest) SetResourceOwnerId(v int64) *SetServerCertificateNameRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetServerCertificateNameRequest) SetServerCertificateId(v string) *SetServerCertificateNameRequest {
	s.ServerCertificateId = &v
	return s
}

func (s *SetServerCertificateNameRequest) SetServerCertificateName(v string) *SetServerCertificateNameRequest {
	s.ServerCertificateName = &v
	return s
}

func (s *SetServerCertificateNameRequest) Validate() error {
	return dara.Validate(s)
}
