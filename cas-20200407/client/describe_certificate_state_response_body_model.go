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
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The content of the certificate in the PEM format. For more information about the PEM format and how to convert certificate formats, see [What formats are used for mainstream digital certificates?](https://help.aliyun.com/document_detail/42214.html)
	//
	// >  This parameter is returned only when the value of the **Type*	- parameter is **certificate**. The value certificate indicates that the certificate is issued.
	//
	// example:
	//
	// ——BEGIN CERTIFICATE—— …… ——END CERTIFICATE——
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// The content that you need to write to the newly created file when you use the file verification method.
	//
	// >  This parameter is returned only when the value of the **Type*	- parameter is **domain_verify*	- and the value of the **ValidateType*	- parameter is **FILE**. The value domain_verify indicates that the verification of the domain name ownership is not complete, and the value FILE indicates that the file verification method is used.
	//
	// example:
	//
	// http://example.com/.well-known/pki-validation/fileauth.txt
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The domain name to be verified when you use the file verification method. You must connect to the DNS server of the domain name and create a file on the server. The file is specified by the **Uri*	- parameter.
	//
	// >  This parameter is returned only when the value of the **Type*	- parameter is **domain_verify*	- and the value of the **ValidateType*	- parameter is **FILE**. The value domain_verify indicates that the verification of the domain name ownership is not complete, and the value FILE indicates that the file verification method is used.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The private key of the certificate in the PEM format. For more information about the PEM format and how to convert certificate formats, see [What formats are used for mainstream digital certificates?](https://help.aliyun.com/document_detail/42214.html)
	//
	// >  This parameter is returned only when the value of the **Type*	- parameter is **certificate**. The value certificate indicates that the certificate is issued.
	//
	// example:
	//
	// ——BEGIN RSA PRIVATE KEY—— …… ——END RSA PRIVATE KEY——
	PrivateKey *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	// The DNS record that you need to manage when you use the DNS verification method.
	//
	// >  This parameter is returned only when the value of the **Type*	- parameter is **domain_verify*	- and the value of the **ValidateType*	- parameter is **DNS**. The value domain_verify indicates that the verification of the domain name ownership is not complete, and the value DNS indicates that the DNS verification method is used.
	//
	// example:
	//
	// _dnsauth
	RecordDomain *string `json:"RecordDomain,omitempty" xml:"RecordDomain,omitempty"`
	// The type of the DNS record that you need to add when you use the DNS verification method. Valid values:
	//
	// 	- **TXT**
	//
	// 	- **CNAME**
	//
	// >  This parameter is returned only when the value of the **Type*	- parameter is **domain_verify*	- and the value of the **ValidateType*	- parameter is **DNS**. The value domain_verify indicates that the verification of the domain name ownership is not complete.
	//
	// example:
	//
	// TXT
	RecordType *string `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	// You need to add a TXT record to the DNS records only when you use the DNS verification method.
	//
	// >  This parameter is returned only when the value of the **Type*	- parameter is **domain_verify*	- and the value of the **ValidateType*	- parameter is **DNS**. The value domain_verify indicates that the verification of the domain name ownership is not complete, and the value DNS indicates that the DNS verification method is used.
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
	// The status of the certificate application order. Valid values:
	//
	// 	- **domain_verify**: **pending review**, which indicates that you have not completed the verification of the domain name ownership after you submit the certificate application.
	//
	//      >After you submit a certificate application, you must manually complete the verification of the domain name ownership. The CA reviews the certificate application only after the verification is complete. If you have not completed the verification of the domain name ownership, you can complete the verification based on the data returned by this operation.
	//
	// 	- **process**: **being reviewed**, which indicates that the certificate application is being reviewed by the CA.
	//
	// 	- **verify_fail**: **review failed**, which indicates that the certificate application failed to be reviewed.
	//
	//     >  If a certificate application fails to be reviewed, the information that you specified in the certificate application may be incorrect. We recommend that you call the [DeleteCertificateRequest](https://help.aliyun.com/document_detail/164109.html) operation to delete the certificate application order and resubmit a certificate application. After the order is deleted, the quota that is consumed for the order is released.
	//
	// 	- **certificate**: **issued**, which indicates that the certificate is issued.
	//
	// 	- **payed**: **pending application**, which indicates that you have not submitted a certificate application.
	//
	// 	- **unknow**: The status is **unknown**.
	//
	// example:
	//
	// domain_verify
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The file that you need to create on the DNS server when you use the file verification method. **The value of this parameter contains the file path and file name.**
	//
	// >  This parameter is returned only when the value of the **Type*	- parameter is **domain_verify*	- and the value of the **ValidateType*	- parameter is **FILE**. The value domain_verify indicates that the verification of the domain name ownership is not complete, and the value FILE indicates that the file verification method is used.
	//
	// example:
	//
	// /.well-known/pki-validation/fileauth.txt
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	// The verification method of the domain name ownership. Valid values:
	//
	// 	- **DNS**: DNS verification. If you use this method, you must add a TXT record to the DNS records of the domain name in the management platform of the domain name.
	//
	// 	- **FILE**: file verification. If you use this method, you must create a specified file on the DNS server.
	//
	// >  This parameter is returned only when the value of the **Type*	- parameter is **domain_verify**. The value domain_verify indicates that the verification of the domain name ownership is not complete.
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
