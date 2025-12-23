// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCertificateWithCsrRequestRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCsr(v string) *CreateCertificateWithCsrRequestRequest
	GetCsr() *string
	SetEmail(v string) *CreateCertificateWithCsrRequestRequest
	GetEmail() *string
	SetPhone(v string) *CreateCertificateWithCsrRequestRequest
	GetPhone() *string
	SetProductCode(v string) *CreateCertificateWithCsrRequestRequest
	GetProductCode() *string
	SetTags(v []*CreateCertificateWithCsrRequestRequestTags) *CreateCertificateWithCsrRequestRequest
	GetTags() []*CreateCertificateWithCsrRequestRequestTags
	SetUsername(v string) *CreateCertificateWithCsrRequestRequest
	GetUsername() *string
	SetValidateType(v string) *CreateCertificateWithCsrRequestRequest
	GetValidateType() *string
}

type CreateCertificateWithCsrRequestRequest struct {
	// The content of the CSR file.\\
	//
	// The key algorithm in the CSR file must be Rivest-Shamir-Adleman (RSA) or elliptic-curve cryptography (ECC), and the key length of the RSA algorithm must be greater than or equal to 2,048 characters. For more information about how to create a CSR file, see [How do I create a CSR file?](https://help.aliyun.com/document_detail/42218.html)\\
	//
	// A CSR file contains the information about your server and company. When you apply for a certificate, you must submit the CSR file to the CA. The CA signs the CSR file by using the private key of the root certificate and generates a public key file to issue your certificate.
	//
	// >  The **CN*	- field in the CSR file specifies the domain name that is bound to the certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE REQUEST----- ...... -----END CERTIFICATE REQUEST-----
	Csr *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	// The contact email address of the applicant.
	//
	// This parameter is required.
	//
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The phone number of the applicant.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1390000****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// The specifications of the certificate that you want to apply for. Valid values:
	//
	// 	- **digicert-free-1-free*	- (default): DigiCert single-domain DV certificate in a three-month free trial, available only on the China site (aliyun.com).
	//
	// 	- **symantec-free-1-free**: DigiCert single-domain DV certificate in a one-year free trial, available only on the China site (aliyun.com).
	//
	// 	- **symantec-dv-1-starter**: DigiCert wildcard DV certificate.
	//
	// 	- **geotrust-dv-1-starter**: GeoTrust single-domain DV certificate.
	//
	// 	- **geotrust-dv-w-starter**: GeoTrust wildcard DV certificate.
	//
	// 	- **globalsign-dv-1-personal**: GlobalSign single-domain DV certificate.
	//
	// 	- **globalsign-dv-w-advanced**: GlobalSign wildcard DV certificate.
	//
	// example:
	//
	// symantec-free-1-free
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The tag list.
	Tags []*CreateCertificateWithCsrRequestRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The name of the applicant.
	//
	// This parameter is required.
	//
	// example:
	//
	// Tom
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// The method to verify the ownership of a domain name. Valid values:
	//
	// 	- **DNS**: DNS verification. If you use this method, you must add a TXT record to the DNS records of the domain name in the management platform of the domain name. You must have operation permissions on domain name resolution to verify the ownership of the domain name.
	//
	// 	- **FILE**: file verification. If you use this method, you must create a specified file on the DNS server. You must have administrative rights on the DNS server to verify the ownership of the domain name.
	//
	// For more information about the verification methods, see [Verify the ownership of a domain name](https://help.aliyun.com/document_detail/48016.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// DNS
	ValidateType *string `json:"ValidateType,omitempty" xml:"ValidateType,omitempty"`
}

func (s CreateCertificateWithCsrRequestRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCertificateWithCsrRequestRequest) GoString() string {
	return s.String()
}

func (s *CreateCertificateWithCsrRequestRequest) GetCsr() *string {
	return s.Csr
}

func (s *CreateCertificateWithCsrRequestRequest) GetEmail() *string {
	return s.Email
}

func (s *CreateCertificateWithCsrRequestRequest) GetPhone() *string {
	return s.Phone
}

func (s *CreateCertificateWithCsrRequestRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *CreateCertificateWithCsrRequestRequest) GetTags() []*CreateCertificateWithCsrRequestRequestTags {
	return s.Tags
}

func (s *CreateCertificateWithCsrRequestRequest) GetUsername() *string {
	return s.Username
}

func (s *CreateCertificateWithCsrRequestRequest) GetValidateType() *string {
	return s.ValidateType
}

func (s *CreateCertificateWithCsrRequestRequest) SetCsr(v string) *CreateCertificateWithCsrRequestRequest {
	s.Csr = &v
	return s
}

func (s *CreateCertificateWithCsrRequestRequest) SetEmail(v string) *CreateCertificateWithCsrRequestRequest {
	s.Email = &v
	return s
}

func (s *CreateCertificateWithCsrRequestRequest) SetPhone(v string) *CreateCertificateWithCsrRequestRequest {
	s.Phone = &v
	return s
}

func (s *CreateCertificateWithCsrRequestRequest) SetProductCode(v string) *CreateCertificateWithCsrRequestRequest {
	s.ProductCode = &v
	return s
}

func (s *CreateCertificateWithCsrRequestRequest) SetTags(v []*CreateCertificateWithCsrRequestRequestTags) *CreateCertificateWithCsrRequestRequest {
	s.Tags = v
	return s
}

func (s *CreateCertificateWithCsrRequestRequest) SetUsername(v string) *CreateCertificateWithCsrRequestRequest {
	s.Username = &v
	return s
}

func (s *CreateCertificateWithCsrRequestRequest) SetValidateType(v string) *CreateCertificateWithCsrRequestRequest {
	s.ValidateType = &v
	return s
}

func (s *CreateCertificateWithCsrRequestRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateCertificateWithCsrRequestRequestTags struct {
	// The tag key. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// database
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateCertificateWithCsrRequestRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateCertificateWithCsrRequestRequestTags) GoString() string {
	return s.String()
}

func (s *CreateCertificateWithCsrRequestRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateCertificateWithCsrRequestRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateCertificateWithCsrRequestRequestTags) SetKey(v string) *CreateCertificateWithCsrRequestRequestTags {
	s.Key = &v
	return s
}

func (s *CreateCertificateWithCsrRequestRequestTags) SetValue(v string) *CreateCertificateWithCsrRequestRequestTags {
	s.Value = &v
	return s
}

func (s *CreateCertificateWithCsrRequestRequestTags) Validate() error {
	return dara.Validate(s)
}
