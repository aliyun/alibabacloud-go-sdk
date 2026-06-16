// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v *GetDomainResponseBodyDomain) *GetDomainResponseBody
	GetDomain() *GetDomainResponseBodyDomain
	SetRequestId(v string) *GetDomainResponseBody
	GetRequestId() *string
}

type GetDomainResponseBody struct {
	// The domain name object.
	Domain *GetDomainResponseBodyDomain `json:"Domain,omitempty" xml:"Domain,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDomainResponseBody) GoString() string {
	return s.String()
}

func (s *GetDomainResponseBody) GetDomain() *GetDomainResponseBodyDomain {
	return s.Domain
}

func (s *GetDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDomainResponseBody) SetDomain(v *GetDomainResponseBodyDomain) *GetDomainResponseBody {
	s.Domain = v
	return s
}

func (s *GetDomainResponseBody) SetRequestId(v string) *GetDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDomainResponseBody) Validate() error {
	if s.Domain != nil {
		if err := s.Domain.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDomainResponseBodyDomain struct {
	// The ID of the brand.
	//
	// example:
	//
	// brand_xxxxx
	BrandId *string `json:"BrandId,omitempty" xml:"BrandId,omitempty"`
	// The time when the domain name was created. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1649830226000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Specifies whether the domain name is the default domain name.
	//
	// example:
	//
	// false
	DefaultDomain *bool `json:"DefaultDomain,omitempty" xml:"DefaultDomain,omitempty"`
	// The domain name.
	//
	// example:
	//
	// login.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ID of the domain name.
	//
	// example:
	//
	// dm_examplexxxxx
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// The type of the domain name. Valid values:
	//
	// - system_init: The initial domain name.
	//
	// - user_custom: A custom domain name.
	//
	// example:
	//
	// system_init
	DomainType *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	// The ICP filing information about the domain name.
	Filing *GetDomainResponseBodyDomainFiling `json:"Filing,omitempty" xml:"Filing,omitempty" type:"Struct"`
	// The ID of the instance.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The lock status of the domain name. Valid values:
	//
	// - unlock: Normal.
	//
	// - lockByLicense: The domain name is unavailable due to license restrictions.
	//
	// example:
	//
	// unlock
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The time when the domain name was last updated. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1649830226000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetDomainResponseBodyDomain) String() string {
	return dara.Prettify(s)
}

func (s GetDomainResponseBodyDomain) GoString() string {
	return s.String()
}

func (s *GetDomainResponseBodyDomain) GetBrandId() *string {
	return s.BrandId
}

func (s *GetDomainResponseBodyDomain) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetDomainResponseBodyDomain) GetDefaultDomain() *bool {
	return s.DefaultDomain
}

func (s *GetDomainResponseBodyDomain) GetDomain() *string {
	return s.Domain
}

func (s *GetDomainResponseBodyDomain) GetDomainId() *string {
	return s.DomainId
}

func (s *GetDomainResponseBodyDomain) GetDomainType() *string {
	return s.DomainType
}

func (s *GetDomainResponseBodyDomain) GetFiling() *GetDomainResponseBodyDomainFiling {
	return s.Filing
}

func (s *GetDomainResponseBodyDomain) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetDomainResponseBodyDomain) GetLockMode() *string {
	return s.LockMode
}

func (s *GetDomainResponseBodyDomain) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetDomainResponseBodyDomain) SetBrandId(v string) *GetDomainResponseBodyDomain {
	s.BrandId = &v
	return s
}

func (s *GetDomainResponseBodyDomain) SetCreateTime(v int64) *GetDomainResponseBodyDomain {
	s.CreateTime = &v
	return s
}

func (s *GetDomainResponseBodyDomain) SetDefaultDomain(v bool) *GetDomainResponseBodyDomain {
	s.DefaultDomain = &v
	return s
}

func (s *GetDomainResponseBodyDomain) SetDomain(v string) *GetDomainResponseBodyDomain {
	s.Domain = &v
	return s
}

func (s *GetDomainResponseBodyDomain) SetDomainId(v string) *GetDomainResponseBodyDomain {
	s.DomainId = &v
	return s
}

func (s *GetDomainResponseBodyDomain) SetDomainType(v string) *GetDomainResponseBodyDomain {
	s.DomainType = &v
	return s
}

func (s *GetDomainResponseBodyDomain) SetFiling(v *GetDomainResponseBodyDomainFiling) *GetDomainResponseBodyDomain {
	s.Filing = v
	return s
}

func (s *GetDomainResponseBodyDomain) SetInstanceId(v string) *GetDomainResponseBodyDomain {
	s.InstanceId = &v
	return s
}

func (s *GetDomainResponseBodyDomain) SetLockMode(v string) *GetDomainResponseBodyDomain {
	s.LockMode = &v
	return s
}

func (s *GetDomainResponseBodyDomain) SetUpdateTime(v int64) *GetDomainResponseBodyDomain {
	s.UpdateTime = &v
	return s
}

func (s *GetDomainResponseBodyDomain) Validate() error {
	if s.Filing != nil {
		if err := s.Filing.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDomainResponseBodyDomainFiling struct {
	// The ICP filing number that is associated with the domain name. The ICP filing number can be for an entity or a website.
	//
	// example:
	//
	// xICPxxxxxx-xx
	IcpNumber *string `json:"IcpNumber,omitempty" xml:"IcpNumber,omitempty"`
}

func (s GetDomainResponseBodyDomainFiling) String() string {
	return dara.Prettify(s)
}

func (s GetDomainResponseBodyDomainFiling) GoString() string {
	return s.String()
}

func (s *GetDomainResponseBodyDomainFiling) GetIcpNumber() *string {
	return s.IcpNumber
}

func (s *GetDomainResponseBodyDomainFiling) SetIcpNumber(v string) *GetDomainResponseBodyDomainFiling {
	s.IcpNumber = &v
	return s
}

func (s *GetDomainResponseBodyDomainFiling) Validate() error {
	return dara.Validate(s)
}
