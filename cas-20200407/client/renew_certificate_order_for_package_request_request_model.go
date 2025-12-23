// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewCertificateOrderForPackageRequestRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCsr(v string) *RenewCertificateOrderForPackageRequestRequest
	GetCsr() *string
	SetOrderId(v int64) *RenewCertificateOrderForPackageRequestRequest
	GetOrderId() *int64
	SetTags(v []*RenewCertificateOrderForPackageRequestRequestTags) *RenewCertificateOrderForPackageRequestRequest
	GetTags() []*RenewCertificateOrderForPackageRequestRequestTags
}

type RenewCertificateOrderForPackageRequestRequest struct {
	// The content of the certificate signing request (CSR) file that is manually generated for the domain name by using OpenSSL or Keytool. The key algorithm in the CSR file must be Rivest-Shamir-Adleman (RSA) or elliptic-curve cryptography (ECC), and the key length of the RSA algorithm must be greater than or equal to 2,048 characters. For more information about how to create a CSR file, see [How do I create a CSR file?](https://help.aliyun.com/document_detail/42218.html)
	//
	// If you do not specify this parameter, Certificate Management Service automatically generates a CSR file for the domain name in the certificate application order that you want to renew.
	//
	// A CSR file contains the information about your server and company. When you apply for a certificate, you must submit the CSR file to the CA. The CA signs the CSR file by using the private key of the root certificate and generates a public key file to issue your certificate.
	//
	// >  The **CN*	- field in the CSR file specifies the domain name that is bound to the certificate.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE REQUEST----- MIIC1TCCAb0CAQAwgY8xCzAJBgNVBAYTAkNOMRIwEAYDVQQIDAlHdWFuZ3pob3Ux ETAPBgNVBAcMCFNoZW56aGVuMQ8wDQYDVQQKDAZDaGFjdW8xEDAOBgNVBAsMB0lU IERlcHQxFzAVBgNVBAMMDnd3dy5jaGFjdW8ubmV0MR0wGwYJKoZIhvcNAQkBFg44 MjkyNjY5QHFxLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBALo7 atRvQf9tKo1NJ/MQqzHvIjHNhU+0MMerDq+tRlJ+a7Ro1r6IWNF5MB0Z*****	- -----END CERTIFICATE REQUEST-----
	Csr *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	// The ID of the certificate application order that you want to renew.
	//
	// >  After you call the [CreateCertificateForPackageRequest](https://help.aliyun.com/document_detail/455296.html), [CreateCertificateRequest](https://help.aliyun.com/document_detail/455292.html), or [CreateCertificateWithCsrRequest](https://help.aliyun.com/document_detail/455801.html) operation to submit a certificate application, you can obtain the ID of the certificate application order from the **OrderId*	- response parameter. You can also call the [ListUserCertificateOrder](https://help.aliyun.com/document_detail/455804.html) operation to obtain the order ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123451222
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The tags.
	Tags []*RenewCertificateOrderForPackageRequestRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s RenewCertificateOrderForPackageRequestRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewCertificateOrderForPackageRequestRequest) GoString() string {
	return s.String()
}

func (s *RenewCertificateOrderForPackageRequestRequest) GetCsr() *string {
	return s.Csr
}

func (s *RenewCertificateOrderForPackageRequestRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *RenewCertificateOrderForPackageRequestRequest) GetTags() []*RenewCertificateOrderForPackageRequestRequestTags {
	return s.Tags
}

func (s *RenewCertificateOrderForPackageRequestRequest) SetCsr(v string) *RenewCertificateOrderForPackageRequestRequest {
	s.Csr = &v
	return s
}

func (s *RenewCertificateOrderForPackageRequestRequest) SetOrderId(v int64) *RenewCertificateOrderForPackageRequestRequest {
	s.OrderId = &v
	return s
}

func (s *RenewCertificateOrderForPackageRequestRequest) SetTags(v []*RenewCertificateOrderForPackageRequestRequestTags) *RenewCertificateOrderForPackageRequestRequest {
	s.Tags = v
	return s
}

func (s *RenewCertificateOrderForPackageRequestRequest) Validate() error {
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

type RenewCertificateOrderForPackageRequestRequestTags struct {
	// The tag key. You can specify at most 20 tag keys. The tag key cannot be an empty string.
	//
	// A tag key can be up to 128 characters in length. It cannot start with aliyun or acs:, and cannot contain http:// or https://.
	//
	// example:
	//
	// account
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the resource tag. A maximum of 20 tag values can be entered. If this value needs to be passed in, an empty string can be entered.
	//
	// A maximum of 128 characters are supported, it cannot start with \\"aliyun\\" or \\"acs:\\", and it cannot contain \\"http://\\" or \\"https://\\".
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s RenewCertificateOrderForPackageRequestRequestTags) String() string {
	return dara.Prettify(s)
}

func (s RenewCertificateOrderForPackageRequestRequestTags) GoString() string {
	return s.String()
}

func (s *RenewCertificateOrderForPackageRequestRequestTags) GetKey() *string {
	return s.Key
}

func (s *RenewCertificateOrderForPackageRequestRequestTags) GetValue() *string {
	return s.Value
}

func (s *RenewCertificateOrderForPackageRequestRequestTags) SetKey(v string) *RenewCertificateOrderForPackageRequestRequestTags {
	s.Key = &v
	return s
}

func (s *RenewCertificateOrderForPackageRequestRequestTags) SetValue(v string) *RenewCertificateOrderForPackageRequestRequestTags {
	s.Value = &v
	return s
}

func (s *RenewCertificateOrderForPackageRequestRequestTags) Validate() error {
	return dara.Validate(s)
}
