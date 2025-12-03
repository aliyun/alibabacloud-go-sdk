// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadServerCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliCloudCertificateId(v string) *UploadServerCertificateRequest
	GetAliCloudCertificateId() *string
	SetAliCloudCertificateName(v string) *UploadServerCertificateRequest
	GetAliCloudCertificateName() *string
	SetAliCloudCertificateRegionId(v string) *UploadServerCertificateRequest
	GetAliCloudCertificateRegionId() *string
	SetOwnerAccount(v string) *UploadServerCertificateRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UploadServerCertificateRequest
	GetOwnerId() *int64
	SetPrivateKey(v string) *UploadServerCertificateRequest
	GetPrivateKey() *string
	SetRegionId(v string) *UploadServerCertificateRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *UploadServerCertificateRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *UploadServerCertificateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UploadServerCertificateRequest
	GetResourceOwnerId() *int64
	SetServerCertificate(v string) *UploadServerCertificateRequest
	GetServerCertificate() *string
	SetServerCertificateName(v string) *UploadServerCertificateRequest
	GetServerCertificateName() *string
	SetTag(v []*UploadServerCertificateRequestTag) *UploadServerCertificateRequest
	GetTag() []*UploadServerCertificateRequestTag
}

type UploadServerCertificateRequest struct {
	// AliCloud certificate ID.
	//
	// example:
	//
	// 775****
	AliCloudCertificateId *string `json:"AliCloudCertificateId,omitempty" xml:"AliCloudCertificateId,omitempty"`
	// AliCloud certificate name.
	//
	// example:
	//
	// cloudcertificate
	AliCloudCertificateName *string `json:"AliCloudCertificateName,omitempty" xml:"AliCloudCertificateName,omitempty"`
	// The region ID of AliCloud certificate.
	//
	// example:
	//
	// cn-hangzhou
	AliCloudCertificateRegionId *string `json:"AliCloudCertificateRegionId,omitempty" xml:"AliCloudCertificateRegionId,omitempty"`
	OwnerAccount                *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The private key of the certificate.
	//
	// example:
	//
	// -----BEGIN RSA PRIVATE KEY----- MIIEogIB*****	- -----END RSA PRIVATE KEY-----
	PrivateKey *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	// The region ID of the CLB instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/27584.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-atstuj3rto****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The server certificate to be uploaded.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE----- MIIGDTCC*****	- -----END CERTIFICATE-----
	ServerCertificate *string `json:"ServerCertificate,omitempty" xml:"ServerCertificate,omitempty"`
	// The name of the server certificate.
	//
	// The name must be 1 to 80 characters in length. It must start with an English letter. It can contain letters, numbers, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// mycert01
	ServerCertificateName *string `json:"ServerCertificateName,omitempty" xml:"ServerCertificateName,omitempty"`
	// The tags.
	Tag []*UploadServerCertificateRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s UploadServerCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadServerCertificateRequest) GoString() string {
	return s.String()
}

func (s *UploadServerCertificateRequest) GetAliCloudCertificateId() *string {
	return s.AliCloudCertificateId
}

func (s *UploadServerCertificateRequest) GetAliCloudCertificateName() *string {
	return s.AliCloudCertificateName
}

func (s *UploadServerCertificateRequest) GetAliCloudCertificateRegionId() *string {
	return s.AliCloudCertificateRegionId
}

func (s *UploadServerCertificateRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UploadServerCertificateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UploadServerCertificateRequest) GetPrivateKey() *string {
	return s.PrivateKey
}

func (s *UploadServerCertificateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UploadServerCertificateRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UploadServerCertificateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UploadServerCertificateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UploadServerCertificateRequest) GetServerCertificate() *string {
	return s.ServerCertificate
}

func (s *UploadServerCertificateRequest) GetServerCertificateName() *string {
	return s.ServerCertificateName
}

func (s *UploadServerCertificateRequest) GetTag() []*UploadServerCertificateRequestTag {
	return s.Tag
}

func (s *UploadServerCertificateRequest) SetAliCloudCertificateId(v string) *UploadServerCertificateRequest {
	s.AliCloudCertificateId = &v
	return s
}

func (s *UploadServerCertificateRequest) SetAliCloudCertificateName(v string) *UploadServerCertificateRequest {
	s.AliCloudCertificateName = &v
	return s
}

func (s *UploadServerCertificateRequest) SetAliCloudCertificateRegionId(v string) *UploadServerCertificateRequest {
	s.AliCloudCertificateRegionId = &v
	return s
}

func (s *UploadServerCertificateRequest) SetOwnerAccount(v string) *UploadServerCertificateRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UploadServerCertificateRequest) SetOwnerId(v int64) *UploadServerCertificateRequest {
	s.OwnerId = &v
	return s
}

func (s *UploadServerCertificateRequest) SetPrivateKey(v string) *UploadServerCertificateRequest {
	s.PrivateKey = &v
	return s
}

func (s *UploadServerCertificateRequest) SetRegionId(v string) *UploadServerCertificateRequest {
	s.RegionId = &v
	return s
}

func (s *UploadServerCertificateRequest) SetResourceGroupId(v string) *UploadServerCertificateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UploadServerCertificateRequest) SetResourceOwnerAccount(v string) *UploadServerCertificateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UploadServerCertificateRequest) SetResourceOwnerId(v int64) *UploadServerCertificateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UploadServerCertificateRequest) SetServerCertificate(v string) *UploadServerCertificateRequest {
	s.ServerCertificate = &v
	return s
}

func (s *UploadServerCertificateRequest) SetServerCertificateName(v string) *UploadServerCertificateRequest {
	s.ServerCertificateName = &v
	return s
}

func (s *UploadServerCertificateRequest) SetTag(v []*UploadServerCertificateRequestTag) *UploadServerCertificateRequest {
	s.Tag = v
	return s
}

func (s *UploadServerCertificateRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UploadServerCertificateRequestTag struct {
	// The key of tag N. Valid values of N: **1 to 20**. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length, and cannot contain `http://` or `https://`. The tag key cannot start with `aliyun` or `acs:`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N. Valid values of N: **1 to 20**. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag value cannot start with `aliyun` and `acs:`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UploadServerCertificateRequestTag) String() string {
	return dara.Prettify(s)
}

func (s UploadServerCertificateRequestTag) GoString() string {
	return s.String()
}

func (s *UploadServerCertificateRequestTag) GetKey() *string {
	return s.Key
}

func (s *UploadServerCertificateRequestTag) GetValue() *string {
	return s.Value
}

func (s *UploadServerCertificateRequestTag) SetKey(v string) *UploadServerCertificateRequestTag {
	s.Key = &v
	return s
}

func (s *UploadServerCertificateRequestTag) SetValue(v string) *UploadServerCertificateRequestTag {
	s.Value = &v
	return s
}

func (s *UploadServerCertificateRequestTag) Validate() error {
	return dara.Validate(s)
}
