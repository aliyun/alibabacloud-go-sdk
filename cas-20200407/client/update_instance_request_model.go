// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoReissue(v string) *UpdateInstanceRequest
	GetAutoReissue() *string
	SetCertificateName(v string) *UpdateInstanceRequest
	GetCertificateName() *string
	SetCity(v string) *UpdateInstanceRequest
	GetCity() *string
	SetCompanyId(v int64) *UpdateInstanceRequest
	GetCompanyId() *int64
	SetContactIdList(v []*int64) *UpdateInstanceRequest
	GetContactIdList() []*int64
	SetCountryCode(v string) *UpdateInstanceRequest
	GetCountryCode() *string
	SetCsr(v string) *UpdateInstanceRequest
	GetCsr() *string
	SetDomain(v string) *UpdateInstanceRequest
	GetDomain() *string
	SetGenerateCsrMethod(v string) *UpdateInstanceRequest
	GetGenerateCsrMethod() *string
	SetInstanceId(v string) *UpdateInstanceRequest
	GetInstanceId() *string
	SetKeyAlgorithm(v string) *UpdateInstanceRequest
	GetKeyAlgorithm() *string
	SetProvince(v string) *UpdateInstanceRequest
	GetProvince() *string
	SetResourceGroupId(v string) *UpdateInstanceRequest
	GetResourceGroupId() *string
	SetTags(v []*UpdateInstanceRequestTags) *UpdateInstanceRequest
	GetTags() []*UpdateInstanceRequestTags
	SetValidationMethod(v string) *UpdateInstanceRequest
	GetValidationMethod() *string
}

type UpdateInstanceRequest struct {
	// example:
	//
	// enable
	AutoReissue *string `json:"AutoReissue,omitempty" xml:"AutoReissue,omitempty"`
	// example:
	//
	// 123
	CertificateName *string `json:"CertificateName,omitempty" xml:"CertificateName,omitempty"`
	// example:
	//
	// Beijing
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// example:
	//
	// 44211
	CompanyId     *int64   `json:"CompanyId,omitempty" xml:"CompanyId,omitempty"`
	ContactIdList []*int64 `json:"ContactIdList,omitempty" xml:"ContactIdList,omitempty" type:"Repeated"`
	// example:
	//
	// CN
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	// example:
	//
	// -----BEGIN CERTIFICATE REQUEST----- ...... -----END CERTIFICATE REQUEST-----
	Csr *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	// example:
	//
	// test.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// online
	GenerateCsrMethod *string `json:"GenerateCsrMethod,omitempty" xml:"GenerateCsrMethod,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cas-cn-68n1mm16****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// RSA_2048
	KeyAlgorithm *string `json:"KeyAlgorithm,omitempty" xml:"KeyAlgorithm,omitempty"`
	// example:
	//
	// Beijing
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// example:
	//
	// rg-ae******4wia
	ResourceGroupId *string                      `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tags            []*UpdateInstanceRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// DNS
	ValidationMethod *string `json:"ValidationMethod,omitempty" xml:"ValidationMethod,omitempty"`
}

func (s UpdateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequest) GetAutoReissue() *string {
	return s.AutoReissue
}

func (s *UpdateInstanceRequest) GetCertificateName() *string {
	return s.CertificateName
}

func (s *UpdateInstanceRequest) GetCity() *string {
	return s.City
}

func (s *UpdateInstanceRequest) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *UpdateInstanceRequest) GetContactIdList() []*int64 {
	return s.ContactIdList
}

func (s *UpdateInstanceRequest) GetCountryCode() *string {
	return s.CountryCode
}

func (s *UpdateInstanceRequest) GetCsr() *string {
	return s.Csr
}

func (s *UpdateInstanceRequest) GetDomain() *string {
	return s.Domain
}

func (s *UpdateInstanceRequest) GetGenerateCsrMethod() *string {
	return s.GenerateCsrMethod
}

func (s *UpdateInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateInstanceRequest) GetKeyAlgorithm() *string {
	return s.KeyAlgorithm
}

func (s *UpdateInstanceRequest) GetProvince() *string {
	return s.Province
}

func (s *UpdateInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateInstanceRequest) GetTags() []*UpdateInstanceRequestTags {
	return s.Tags
}

func (s *UpdateInstanceRequest) GetValidationMethod() *string {
	return s.ValidationMethod
}

func (s *UpdateInstanceRequest) SetAutoReissue(v string) *UpdateInstanceRequest {
	s.AutoReissue = &v
	return s
}

func (s *UpdateInstanceRequest) SetCertificateName(v string) *UpdateInstanceRequest {
	s.CertificateName = &v
	return s
}

func (s *UpdateInstanceRequest) SetCity(v string) *UpdateInstanceRequest {
	s.City = &v
	return s
}

func (s *UpdateInstanceRequest) SetCompanyId(v int64) *UpdateInstanceRequest {
	s.CompanyId = &v
	return s
}

func (s *UpdateInstanceRequest) SetContactIdList(v []*int64) *UpdateInstanceRequest {
	s.ContactIdList = v
	return s
}

func (s *UpdateInstanceRequest) SetCountryCode(v string) *UpdateInstanceRequest {
	s.CountryCode = &v
	return s
}

func (s *UpdateInstanceRequest) SetCsr(v string) *UpdateInstanceRequest {
	s.Csr = &v
	return s
}

func (s *UpdateInstanceRequest) SetDomain(v string) *UpdateInstanceRequest {
	s.Domain = &v
	return s
}

func (s *UpdateInstanceRequest) SetGenerateCsrMethod(v string) *UpdateInstanceRequest {
	s.GenerateCsrMethod = &v
	return s
}

func (s *UpdateInstanceRequest) SetInstanceId(v string) *UpdateInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceRequest) SetKeyAlgorithm(v string) *UpdateInstanceRequest {
	s.KeyAlgorithm = &v
	return s
}

func (s *UpdateInstanceRequest) SetProvince(v string) *UpdateInstanceRequest {
	s.Province = &v
	return s
}

func (s *UpdateInstanceRequest) SetResourceGroupId(v string) *UpdateInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateInstanceRequest) SetTags(v []*UpdateInstanceRequestTags) *UpdateInstanceRequest {
	s.Tags = v
	return s
}

func (s *UpdateInstanceRequest) SetValidationMethod(v string) *UpdateInstanceRequest {
	s.ValidationMethod = &v
	return s
}

func (s *UpdateInstanceRequest) Validate() error {
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

type UpdateInstanceRequestTags struct {
	// example:
	//
	// test
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// example:
	//
	// test
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s UpdateInstanceRequestTags) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceRequestTags) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequestTags) GetTagKey() *string {
	return s.TagKey
}

func (s *UpdateInstanceRequestTags) GetTagValue() *string {
	return s.TagValue
}

func (s *UpdateInstanceRequestTags) SetTagKey(v string) *UpdateInstanceRequestTags {
	s.TagKey = &v
	return s
}

func (s *UpdateInstanceRequestTags) SetTagValue(v string) *UpdateInstanceRequestTags {
	s.TagValue = &v
	return s
}

func (s *UpdateInstanceRequestTags) Validate() error {
	return dara.Validate(s)
}
