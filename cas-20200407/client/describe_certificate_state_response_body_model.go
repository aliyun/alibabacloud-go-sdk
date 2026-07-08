// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCertificateStateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertId(v string) *DescribeCertificateStateResponseBody
	GetCertId() *string
	SetCertificate(v string) *DescribeCertificateStateResponseBody
	GetCertificate() *string
	SetContent(v string) *DescribeCertificateStateResponseBody
	GetContent() *string
	SetDomain(v string) *DescribeCertificateStateResponseBody
	GetDomain() *string
	SetPrivateKey(v string) *DescribeCertificateStateResponseBody
	GetPrivateKey() *string
	SetRecordDomain(v string) *DescribeCertificateStateResponseBody
	GetRecordDomain() *string
	SetRecordType(v string) *DescribeCertificateStateResponseBody
	GetRecordType() *string
	SetRecordValue(v string) *DescribeCertificateStateResponseBody
	GetRecordValue() *string
	SetRequestId(v string) *DescribeCertificateStateResponseBody
	GetRequestId() *string
	SetType(v string) *DescribeCertificateStateResponseBody
	GetType() *string
	SetUri(v string) *DescribeCertificateStateResponseBody
	GetUri() *string
	SetValidateType(v string) *DescribeCertificateStateResponseBody
	GetValidateType() *string
}

type DescribeCertificateStateResponseBody struct {
	// The certificate ID.
	//
	// > This parameter is returned when the certificate is issued.
	//
	// example:
	//
	// 111111
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The certificate content (in PEM format). For more information about the PEM format and how to convert the format of a certificate, see [What are the formats of mainstream digital certificates?](https://help.aliyun.com/document_detail/42214.html).
	//
	// > This parameter is returned only when **Type*	- is set to **certificate*	- (indicating that the certificate has been issued).
	//
	// example:
	//
	// -----BEGIN CERTIFICATE----- …… -----END CERTIFICATE-----
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// The content that you need to write to the newly created file when you use the file validation method for domain validation.
	//
	// > This parameter is returned only when **Type*	- is set to **domain_verify*	- (indicating the domain validation stage) and **ValidateType*	- is set to **FILE*	- (indicating the file validation method).
	//
	// example:
	//
	// http://example.com/.well-known/pki-validation/fileauth.txt
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The domain name to be validated when you use the file validation method for domain validation. You need to connect to the server corresponding to this domain name and create the specified file (i.e., **Uri**) on the server.
	//
	// > This parameter is returned only when **Type*	- is set to **domain_verify*	- (indicating the domain validation stage) and **ValidateType*	- is set to **FILE*	- (indicating the file validation method).
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The content of the certificate private key (in PEM format). For more information about the PEM format and how to convert the format of a certificate, see [What are the formats of mainstream digital certificates?](https://help.aliyun.com/document_detail/42214.html).
	//
	// > This parameter is returned only when **Type*	- is set to **certificate*	- (indicating that the certificate has been issued).
	//
	// example:
	//
	// -----BEGIN RSA PRIVATE KEY-----…… -----END RSA PRIVATE KEY-----
	PrivateKey *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	// The host record that you need to operate when you use the DNS validation method for domain validation.
	//
	// > This parameter is returned only when **Type*	- is set to **domain_verify*	- (indicating the domain validation stage) and **ValidateType*	- is set to **DNS*	- (indicating the DNS validation method).
	//
	// example:
	//
	// _dnsauth
	RecordDomain *string `json:"RecordDomain,omitempty" xml:"RecordDomain,omitempty"`
	// The type of DNS record that you need to add when you use the DNS validation method for domain validation. Valid values:
	//
	// - **TXT**: text record.
	//
	// - **CNAME**: alias record.
	//
	// > This parameter is returned only when **Type*	- is set to **domain_verify*	- (indicating the domain validation stage) and **ValidateType*	- is set to **DNS*	- (indicating the DNS validation method).
	//
	// example:
	//
	// TXT
	RecordType *string `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	// The record value that you need to add when you use the DNS validation method for domain validation.
	//
	// > This parameter is returned only when **Type*	- is set to **domain_verify*	- (indicating the domain validation stage) and **ValidateType*	- is set to **DNS*	- (indicating the DNS validation method).
	//
	// example:
	//
	// 20200420000000223erigacv46uhaubchcm0o7spxi7i2isvjq59mlx9lucnkqcy
	RecordValue *string `json:"RecordValue,omitempty" xml:"RecordValue,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 082FAB35-6AB9-4FD5-8750-D36673548E76
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the certificate request order. Valid values:
	//
	// - **domain_verify**: **Pending validation**, which indicates that you have not completed domain validation after submitting the certificate request.
	//
	//   > After you submit a certificate request, you must manually complete domain ownership validation before the certificate request can enter the review stage. If you have not completed domain validation, you can refer to the response parameters of this operation to complete domain validation.
	//
	// - **process**: **Under review**, which indicates that the certificate request is being reviewed by the CA center.
	//
	// - **verify_fail**: **Review failed**, which indicates that the certificate request failed the review.
	//
	//   > The review may fail because the certificate request information you submitted is incorrect. We recommend that you call [DeleteCertificateRequest](https://help.aliyun.com/document_detail/455294.html) to delete the order that failed the review (deleted orders do not consume certificate resource plan quota) and submit a new certificate request.
	//
	// - **certificate**: **Issued**, which indicates that the certificate has been issued.
	//
	// - **payed**: **Pending request**, which indicates that the certificate is pending request.
	//
	// - **unknow**: **Unknown status**.
	//
	// example:
	//
	// domain_verify
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The file that you need to create on the domain server when you use the file validation method for domain validation. **Uri*	- includes the file path and name.
	//
	// > This parameter is returned only when **Type*	- is set to **domain_verify*	- (indicating the domain validation stage) and **ValidateType*	- is set to **FILE*	- (indicating the file validation method).
	//
	// example:
	//
	// /.well-known/pki-validation/fileauth.txt
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	// The domain validation method selected when submitting the certificate request. Valid values:
	//
	// - **DNS**: DNS validation. This method validates domain ownership by adding the specified DNS record to the domain on the DNS management platform.
	//
	// - **FILE**: file validation. This method validates domain ownership by creating the specified file on the domain server.
	//
	// > This parameter is returned only when **Type*	- is set to **domain_verify*	- (indicating the domain validation stage).
	//
	// example:
	//
	// FILE
	ValidateType *string `json:"ValidateType,omitempty" xml:"ValidateType,omitempty"`
}

func (s DescribeCertificateStateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCertificateStateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCertificateStateResponseBody) GetCertId() *string {
	return s.CertId
}

func (s *DescribeCertificateStateResponseBody) GetCertificate() *string {
	return s.Certificate
}

func (s *DescribeCertificateStateResponseBody) GetContent() *string {
	return s.Content
}

func (s *DescribeCertificateStateResponseBody) GetDomain() *string {
	return s.Domain
}

func (s *DescribeCertificateStateResponseBody) GetPrivateKey() *string {
	return s.PrivateKey
}

func (s *DescribeCertificateStateResponseBody) GetRecordDomain() *string {
	return s.RecordDomain
}

func (s *DescribeCertificateStateResponseBody) GetRecordType() *string {
	return s.RecordType
}

func (s *DescribeCertificateStateResponseBody) GetRecordValue() *string {
	return s.RecordValue
}

func (s *DescribeCertificateStateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCertificateStateResponseBody) GetType() *string {
	return s.Type
}

func (s *DescribeCertificateStateResponseBody) GetUri() *string {
	return s.Uri
}

func (s *DescribeCertificateStateResponseBody) GetValidateType() *string {
	return s.ValidateType
}

func (s *DescribeCertificateStateResponseBody) SetCertId(v string) *DescribeCertificateStateResponseBody {
	s.CertId = &v
	return s
}

func (s *DescribeCertificateStateResponseBody) SetCertificate(v string) *DescribeCertificateStateResponseBody {
	s.Certificate = &v
	return s
}

func (s *DescribeCertificateStateResponseBody) SetContent(v string) *DescribeCertificateStateResponseBody {
	s.Content = &v
	return s
}

func (s *DescribeCertificateStateResponseBody) SetDomain(v string) *DescribeCertificateStateResponseBody {
	s.Domain = &v
	return s
}

func (s *DescribeCertificateStateResponseBody) SetPrivateKey(v string) *DescribeCertificateStateResponseBody {
	s.PrivateKey = &v
	return s
}

func (s *DescribeCertificateStateResponseBody) SetRecordDomain(v string) *DescribeCertificateStateResponseBody {
	s.RecordDomain = &v
	return s
}

func (s *DescribeCertificateStateResponseBody) SetRecordType(v string) *DescribeCertificateStateResponseBody {
	s.RecordType = &v
	return s
}

func (s *DescribeCertificateStateResponseBody) SetRecordValue(v string) *DescribeCertificateStateResponseBody {
	s.RecordValue = &v
	return s
}

func (s *DescribeCertificateStateResponseBody) SetRequestId(v string) *DescribeCertificateStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCertificateStateResponseBody) SetType(v string) *DescribeCertificateStateResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeCertificateStateResponseBody) SetUri(v string) *DescribeCertificateStateResponseBody {
	s.Uri = &v
	return s
}

func (s *DescribeCertificateStateResponseBody) SetValidateType(v string) *DescribeCertificateStateResponseBody {
	s.ValidateType = &v
	return s
}

func (s *DescribeCertificateStateResponseBody) Validate() error {
	return dara.Validate(s)
}
