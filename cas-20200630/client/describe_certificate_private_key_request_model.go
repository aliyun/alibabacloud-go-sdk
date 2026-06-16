// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCertificatePrivateKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEncryptedCode(v string) *DescribeCertificatePrivateKeyRequest
	GetEncryptedCode() *string
	SetIdentifier(v string) *DescribeCertificatePrivateKeyRequest
	GetIdentifier() *string
	SetResourceGroupId(v string) *DescribeCertificatePrivateKeyRequest
	GetResourceGroupId() *string
}

type DescribeCertificatePrivateKeyRequest struct {
	// The password to encrypt the private key. The password can contain uppercase letters, lowercase letters, digits, and special characters, such as `,.+-_#`. The maximum length is 32 bytes.
	//
	// 	Warning:
	//
	// Remember the password you set. You need this password to decrypt the encrypted private key. If you forget the password, you cannot decrypt the private key that you get from this API call. You must call this API again to get a new encrypted key.
	//
	// This parameter is required.
	//
	// example:
	//
	// !Demo@WS3ed
	EncryptedCode *string `json:"EncryptedCode,omitempty" xml:"EncryptedCode,omitempty"`
	// The unique identifier of the client or server-side certificate for which you want to get the private key.
	//
	// > Call [ListClientCertificate](https://help.aliyun.com/document_detail/465990.html) to query the unique identifiers of all client and server-side certificates.
	//
	// This parameter is required.
	//
	// example:
	//
	// bc37133bb7ed68c7938d928fd26d****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The ID of the resource group to which the certificate belongs.
	//
	// example:
	//
	// test
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeCertificatePrivateKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCertificatePrivateKeyRequest) GoString() string {
	return s.String()
}

func (s *DescribeCertificatePrivateKeyRequest) GetEncryptedCode() *string {
	return s.EncryptedCode
}

func (s *DescribeCertificatePrivateKeyRequest) GetIdentifier() *string {
	return s.Identifier
}

func (s *DescribeCertificatePrivateKeyRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeCertificatePrivateKeyRequest) SetEncryptedCode(v string) *DescribeCertificatePrivateKeyRequest {
	s.EncryptedCode = &v
	return s
}

func (s *DescribeCertificatePrivateKeyRequest) SetIdentifier(v string) *DescribeCertificatePrivateKeyRequest {
	s.Identifier = &v
	return s
}

func (s *DescribeCertificatePrivateKeyRequest) SetResourceGroupId(v string) *DescribeCertificatePrivateKeyRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeCertificatePrivateKeyRequest) Validate() error {
	return dara.Validate(s)
}
