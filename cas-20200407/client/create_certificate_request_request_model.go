// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCertificateRequestRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *CreateCertificateRequestRequest
	GetDomain() *string
	SetEmail(v string) *CreateCertificateRequestRequest
	GetEmail() *string
	SetPhone(v string) *CreateCertificateRequestRequest
	GetPhone() *string
	SetProductCode(v string) *CreateCertificateRequestRequest
	GetProductCode() *string
	SetTags(v []*CreateCertificateRequestRequestTags) *CreateCertificateRequestRequest
	GetTags() []*CreateCertificateRequestRequestTags
	SetUsername(v string) *CreateCertificateRequestRequest
	GetUsername() *string
	SetValidateType(v string) *CreateCertificateRequestRequest
	GetValidateType() *string
}

type CreateCertificateRequestRequest struct {
	// The domain name that you want to bind to the certificate. You can specify only one domain name.
	//
	// >  The domain name must match the certificate specifications that you specify for the **ProductCode*	- parameter. If you apply for a single-domain certificate, you must specify a single domain name for this parameter. If you apply for a wildcard certificate, you must specify a wildcard domain name such as `*.aliyundoc.com` for this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
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
	// The specifications of the certificate. Valid values:
	//
	// 	- **digicert-free-1-free*	- (default): DigiCert single-domain DV certificate, which is free and valid for 3 months.
	//
	// 	- **symantec-free-1-free**: DigiCert single-domain DV certificate, which is free and valid for 1 year. This value is available only on the China site (aliyun.com).
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
	// The tags.
	Tags []*CreateCertificateRequestRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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

func (s CreateCertificateRequestRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCertificateRequestRequest) GoString() string {
	return s.String()
}

func (s *CreateCertificateRequestRequest) GetDomain() *string {
	return s.Domain
}

func (s *CreateCertificateRequestRequest) GetEmail() *string {
	return s.Email
}

func (s *CreateCertificateRequestRequest) GetPhone() *string {
	return s.Phone
}

func (s *CreateCertificateRequestRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *CreateCertificateRequestRequest) GetTags() []*CreateCertificateRequestRequestTags {
	return s.Tags
}

func (s *CreateCertificateRequestRequest) GetUsername() *string {
	return s.Username
}

func (s *CreateCertificateRequestRequest) GetValidateType() *string {
	return s.ValidateType
}

func (s *CreateCertificateRequestRequest) SetDomain(v string) *CreateCertificateRequestRequest {
	s.Domain = &v
	return s
}

func (s *CreateCertificateRequestRequest) SetEmail(v string) *CreateCertificateRequestRequest {
	s.Email = &v
	return s
}

func (s *CreateCertificateRequestRequest) SetPhone(v string) *CreateCertificateRequestRequest {
	s.Phone = &v
	return s
}

func (s *CreateCertificateRequestRequest) SetProductCode(v string) *CreateCertificateRequestRequest {
	s.ProductCode = &v
	return s
}

func (s *CreateCertificateRequestRequest) SetTags(v []*CreateCertificateRequestRequestTags) *CreateCertificateRequestRequest {
	s.Tags = v
	return s
}

func (s *CreateCertificateRequestRequest) SetUsername(v string) *CreateCertificateRequestRequest {
	s.Username = &v
	return s
}

func (s *CreateCertificateRequestRequest) SetValidateType(v string) *CreateCertificateRequestRequest {
	s.ValidateType = &v
	return s
}

func (s *CreateCertificateRequestRequest) Validate() error {
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

type CreateCertificateRequestRequestTags struct {
	// The tag key of the resource. You can specify up to 20 tag keys. You cannot specify empty strings as tag keys.
	//
	// The key can be up to 64 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The key must start with a letter but cannot start with `aliyun` or `acs:`. The key cannot contain `http://` or `https://`.
	//
	// >  You must specify at least one of **Tag.N*	- (**Tag.N.Key*	- and **Tag.N.Value**).
	//
	// example:
	//
	// acs:rm:rgId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value cannot exceed 128 characters in length, and can contain digits, periods (.), underscores (_), and hyphens (-). The key must start with a letter but cannot start with `aliyun` or `acs:`. The key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateCertificateRequestRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateCertificateRequestRequestTags) GoString() string {
	return s.String()
}

func (s *CreateCertificateRequestRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateCertificateRequestRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateCertificateRequestRequestTags) SetKey(v string) *CreateCertificateRequestRequestTags {
	s.Key = &v
	return s
}

func (s *CreateCertificateRequestRequestTags) SetValue(v string) *CreateCertificateRequestRequestTags {
	s.Value = &v
	return s
}

func (s *CreateCertificateRequestRequestTags) Validate() error {
	return dara.Validate(s)
}
