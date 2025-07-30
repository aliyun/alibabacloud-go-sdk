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
}

type DescribeCertificatePrivateKeyRequest struct {
	// The password that is used to encrypt the private key. The password can contain letters, digits, and special characters, such as `, + - _ #`. The password can be up to 32 bytes in length.
	//
	// **Warning*	- You must remember the password that you specify. The password is required to decrypt the encrypted private key. If you forget the password, the encrypted private key that is returned cannot be decrypted. You must call this operation again.
	//
	// This parameter is required.
	//
	// example:
	//
	// !QA@WS3ed
	EncryptedCode *string `json:"EncryptedCode,omitempty" xml:"EncryptedCode,omitempty"`
	// The unique identifier of the client certificate or server certificate that you want to query.
	//
	// >  You can call the [ListClientCertificate](https://help.aliyun.com/document_detail/330884.html) operation to query the unique identifiers of all client certificates and server certificates.
	//
	// This parameter is required.
	//
	// example:
	//
	// bc37133bb7ed68c7938d928fd26d****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
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

func (s *DescribeCertificatePrivateKeyRequest) SetEncryptedCode(v string) *DescribeCertificatePrivateKeyRequest {
	s.EncryptedCode = &v
	return s
}

func (s *DescribeCertificatePrivateKeyRequest) SetIdentifier(v string) *DescribeCertificatePrivateKeyRequest {
	s.Identifier = &v
	return s
}

func (s *DescribeCertificatePrivateKeyRequest) Validate() error {
	return dara.Validate(s)
}
