// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExternalCACertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiPassthrough(v *CreateExternalCACertificateRequestApiPassthrough) *CreateExternalCACertificateRequest
	GetApiPassthrough() *CreateExternalCACertificateRequestApiPassthrough
	SetCsr(v string) *CreateExternalCACertificateRequest
	GetCsr() *string
	SetInstanceId(v string) *CreateExternalCACertificateRequest
	GetInstanceId() *string
	SetResourceGroupId(v string) *CreateExternalCACertificateRequest
	GetResourceGroupId() *string
	SetTags(v []*CreateExternalCACertificateRequestTags) *CreateExternalCACertificateRequest
	GetTags() []*CreateExternalCACertificateRequestTags
	SetValidity(v string) *CreateExternalCACertificateRequest
	GetValidity() *string
}

type CreateExternalCACertificateRequest struct {
	ApiPassthrough *CreateExternalCACertificateRequestApiPassthrough `json:"ApiPassthrough,omitempty" xml:"ApiPassthrough,omitempty" type:"Struct"`
	// example:
	//
	// -----BEGIN CERTIFICATE REQUEST-----
	//
	// MIIBczCCARgCAQAwgYoxFDASBgNVBAMMC2FsaXl1bi50ZXN0MQ0wCwYDVQQ
	//
	// ...
	//
	// vbIgMQIhAKHDWD6/WAMbtezAt4bysJ/BZIDz1jPWuUR5GV4TJ/mS
	//
	// -----END CERTIFICATE REQUEST-----
	Csr *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	// example:
	//
	// cas_deposit-cn-1234abcd
	InstanceId      *string                                   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ResourceGroupId *string                                   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tags            []*CreateExternalCACertificateRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// 10y
	Validity *string `json:"Validity,omitempty" xml:"Validity,omitempty"`
}

func (s CreateExternalCACertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateExternalCACertificateRequest) GoString() string {
	return s.String()
}

func (s *CreateExternalCACertificateRequest) GetApiPassthrough() *CreateExternalCACertificateRequestApiPassthrough {
	return s.ApiPassthrough
}

func (s *CreateExternalCACertificateRequest) GetCsr() *string {
	return s.Csr
}

func (s *CreateExternalCACertificateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateExternalCACertificateRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateExternalCACertificateRequest) GetTags() []*CreateExternalCACertificateRequestTags {
	return s.Tags
}

func (s *CreateExternalCACertificateRequest) GetValidity() *string {
	return s.Validity
}

func (s *CreateExternalCACertificateRequest) SetApiPassthrough(v *CreateExternalCACertificateRequestApiPassthrough) *CreateExternalCACertificateRequest {
	s.ApiPassthrough = v
	return s
}

func (s *CreateExternalCACertificateRequest) SetCsr(v string) *CreateExternalCACertificateRequest {
	s.Csr = &v
	return s
}

func (s *CreateExternalCACertificateRequest) SetInstanceId(v string) *CreateExternalCACertificateRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateExternalCACertificateRequest) SetResourceGroupId(v string) *CreateExternalCACertificateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateExternalCACertificateRequest) SetTags(v []*CreateExternalCACertificateRequestTags) *CreateExternalCACertificateRequest {
	s.Tags = v
	return s
}

func (s *CreateExternalCACertificateRequest) SetValidity(v string) *CreateExternalCACertificateRequest {
	s.Validity = &v
	return s
}

func (s *CreateExternalCACertificateRequest) Validate() error {
	if s.ApiPassthrough != nil {
		if err := s.ApiPassthrough.Validate(); err != nil {
			return err
		}
	}
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

type CreateExternalCACertificateRequestApiPassthrough struct {
	Extensions *CreateExternalCACertificateRequestApiPassthroughExtensions `json:"Extensions,omitempty" xml:"Extensions,omitempty" type:"Struct"`
	Subject    *CreateExternalCACertificateRequestApiPassthroughSubject    `json:"Subject,omitempty" xml:"Subject,omitempty" type:"Struct"`
}

func (s CreateExternalCACertificateRequestApiPassthrough) String() string {
	return dara.Prettify(s)
}

func (s CreateExternalCACertificateRequestApiPassthrough) GoString() string {
	return s.String()
}

func (s *CreateExternalCACertificateRequestApiPassthrough) GetExtensions() *CreateExternalCACertificateRequestApiPassthroughExtensions {
	return s.Extensions
}

func (s *CreateExternalCACertificateRequestApiPassthrough) GetSubject() *CreateExternalCACertificateRequestApiPassthroughSubject {
	return s.Subject
}

func (s *CreateExternalCACertificateRequestApiPassthrough) SetExtensions(v *CreateExternalCACertificateRequestApiPassthroughExtensions) *CreateExternalCACertificateRequestApiPassthrough {
	s.Extensions = v
	return s
}

func (s *CreateExternalCACertificateRequestApiPassthrough) SetSubject(v *CreateExternalCACertificateRequestApiPassthroughSubject) *CreateExternalCACertificateRequestApiPassthrough {
	s.Subject = v
	return s
}

func (s *CreateExternalCACertificateRequestApiPassthrough) Validate() error {
	if s.Extensions != nil {
		if err := s.Extensions.Validate(); err != nil {
			return err
		}
	}
	if s.Subject != nil {
		if err := s.Subject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateExternalCACertificateRequestApiPassthroughExtensions struct {
	ExtendedKeyUsages []*string `json:"ExtendedKeyUsages,omitempty" xml:"ExtendedKeyUsages,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	PathLenConstraint *int32 `json:"PathLenConstraint,omitempty" xml:"PathLenConstraint,omitempty"`
}

func (s CreateExternalCACertificateRequestApiPassthroughExtensions) String() string {
	return dara.Prettify(s)
}

func (s CreateExternalCACertificateRequestApiPassthroughExtensions) GoString() string {
	return s.String()
}

func (s *CreateExternalCACertificateRequestApiPassthroughExtensions) GetExtendedKeyUsages() []*string {
	return s.ExtendedKeyUsages
}

func (s *CreateExternalCACertificateRequestApiPassthroughExtensions) GetPathLenConstraint() *int32 {
	return s.PathLenConstraint
}

func (s *CreateExternalCACertificateRequestApiPassthroughExtensions) SetExtendedKeyUsages(v []*string) *CreateExternalCACertificateRequestApiPassthroughExtensions {
	s.ExtendedKeyUsages = v
	return s
}

func (s *CreateExternalCACertificateRequestApiPassthroughExtensions) SetPathLenConstraint(v int32) *CreateExternalCACertificateRequestApiPassthroughExtensions {
	s.PathLenConstraint = &v
	return s
}

func (s *CreateExternalCACertificateRequestApiPassthroughExtensions) Validate() error {
	return dara.Validate(s)
}

type CreateExternalCACertificateRequestApiPassthroughSubject struct {
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// example:
	//
	// CN
	Country          *string `json:"Country,omitempty" xml:"Country,omitempty"`
	Locality         *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	Organization     *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	State            *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s CreateExternalCACertificateRequestApiPassthroughSubject) String() string {
	return dara.Prettify(s)
}

func (s CreateExternalCACertificateRequestApiPassthroughSubject) GoString() string {
	return s.String()
}

func (s *CreateExternalCACertificateRequestApiPassthroughSubject) GetCommonName() *string {
	return s.CommonName
}

func (s *CreateExternalCACertificateRequestApiPassthroughSubject) GetCountry() *string {
	return s.Country
}

func (s *CreateExternalCACertificateRequestApiPassthroughSubject) GetLocality() *string {
	return s.Locality
}

func (s *CreateExternalCACertificateRequestApiPassthroughSubject) GetOrganization() *string {
	return s.Organization
}

func (s *CreateExternalCACertificateRequestApiPassthroughSubject) GetOrganizationUnit() *string {
	return s.OrganizationUnit
}

func (s *CreateExternalCACertificateRequestApiPassthroughSubject) GetState() *string {
	return s.State
}

func (s *CreateExternalCACertificateRequestApiPassthroughSubject) SetCommonName(v string) *CreateExternalCACertificateRequestApiPassthroughSubject {
	s.CommonName = &v
	return s
}

func (s *CreateExternalCACertificateRequestApiPassthroughSubject) SetCountry(v string) *CreateExternalCACertificateRequestApiPassthroughSubject {
	s.Country = &v
	return s
}

func (s *CreateExternalCACertificateRequestApiPassthroughSubject) SetLocality(v string) *CreateExternalCACertificateRequestApiPassthroughSubject {
	s.Locality = &v
	return s
}

func (s *CreateExternalCACertificateRequestApiPassthroughSubject) SetOrganization(v string) *CreateExternalCACertificateRequestApiPassthroughSubject {
	s.Organization = &v
	return s
}

func (s *CreateExternalCACertificateRequestApiPassthroughSubject) SetOrganizationUnit(v string) *CreateExternalCACertificateRequestApiPassthroughSubject {
	s.OrganizationUnit = &v
	return s
}

func (s *CreateExternalCACertificateRequestApiPassthroughSubject) SetState(v string) *CreateExternalCACertificateRequestApiPassthroughSubject {
	s.State = &v
	return s
}

func (s *CreateExternalCACertificateRequestApiPassthroughSubject) Validate() error {
	return dara.Validate(s)
}

type CreateExternalCACertificateRequestTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateExternalCACertificateRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateExternalCACertificateRequestTags) GoString() string {
	return s.String()
}

func (s *CreateExternalCACertificateRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateExternalCACertificateRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateExternalCACertificateRequestTags) SetKey(v string) *CreateExternalCACertificateRequestTags {
	s.Key = &v
	return s
}

func (s *CreateExternalCACertificateRequestTags) SetValue(v string) *CreateExternalCACertificateRequestTags {
	s.Value = &v
	return s
}

func (s *CreateExternalCACertificateRequestTags) Validate() error {
	return dara.Validate(s)
}
