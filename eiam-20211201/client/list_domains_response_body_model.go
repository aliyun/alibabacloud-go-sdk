// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDomainsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomains(v []*ListDomainsResponseBodyDomains) *ListDomainsResponseBody
	GetDomains() []*ListDomainsResponseBodyDomains
	SetRequestId(v string) *ListDomainsResponseBody
	GetRequestId() *string
}

type ListDomainsResponseBody struct {
	// The information about the domain names.
	Domains []*ListDomainsResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDomainsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDomainsResponseBody) GetDomains() []*ListDomainsResponseBodyDomains {
	return s.Domains
}

func (s *ListDomainsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDomainsResponseBody) SetDomains(v []*ListDomainsResponseBodyDomains) *ListDomainsResponseBody {
	s.Domains = v
	return s
}

func (s *ListDomainsResponseBody) SetRequestId(v string) *ListDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDomainsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDomainsResponseBodyDomains struct {
	// The time when the domain name was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1649830226000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the domain name is the default domain.
	//
	// example:
	//
	// false
	DefaultDomain *bool `json:"DefaultDomain,omitempty" xml:"DefaultDomain,omitempty"`
	// The domain.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The domain ID.
	//
	// example:
	//
	// dm_examplexxxxx
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// The type of the domain name. Valid values:
	//
	// 	- system_init: an initial domain name.
	//
	// 	- user_custom: a custom domain name.
	//
	// example:
	//
	// system_init
	DomainType *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	// The information about the Internet content provider (ICP) filing of the domain name.
	Filing *ListDomainsResponseBodyDomainsFiling `json:"Filing,omitempty" xml:"Filing,omitempty" type:"Struct"`
	// The instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the domain name is locked. Valid values:
	//
	// 	- unlock
	//
	// 	- lockByLicense
	//
	// example:
	//
	// unlock
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The time when the domain name was last updated. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1649830226000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListDomainsResponseBodyDomains) String() string {
	return dara.Prettify(s)
}

func (s ListDomainsResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *ListDomainsResponseBodyDomains) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListDomainsResponseBodyDomains) GetDefaultDomain() *bool {
	return s.DefaultDomain
}

func (s *ListDomainsResponseBodyDomains) GetDomain() *string {
	return s.Domain
}

func (s *ListDomainsResponseBodyDomains) GetDomainId() *string {
	return s.DomainId
}

func (s *ListDomainsResponseBodyDomains) GetDomainType() *string {
	return s.DomainType
}

func (s *ListDomainsResponseBodyDomains) GetFiling() *ListDomainsResponseBodyDomainsFiling {
	return s.Filing
}

func (s *ListDomainsResponseBodyDomains) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDomainsResponseBodyDomains) GetLockMode() *string {
	return s.LockMode
}

func (s *ListDomainsResponseBodyDomains) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListDomainsResponseBodyDomains) SetCreateTime(v int64) *ListDomainsResponseBodyDomains {
	s.CreateTime = &v
	return s
}

func (s *ListDomainsResponseBodyDomains) SetDefaultDomain(v bool) *ListDomainsResponseBodyDomains {
	s.DefaultDomain = &v
	return s
}

func (s *ListDomainsResponseBodyDomains) SetDomain(v string) *ListDomainsResponseBodyDomains {
	s.Domain = &v
	return s
}

func (s *ListDomainsResponseBodyDomains) SetDomainId(v string) *ListDomainsResponseBodyDomains {
	s.DomainId = &v
	return s
}

func (s *ListDomainsResponseBodyDomains) SetDomainType(v string) *ListDomainsResponseBodyDomains {
	s.DomainType = &v
	return s
}

func (s *ListDomainsResponseBodyDomains) SetFiling(v *ListDomainsResponseBodyDomainsFiling) *ListDomainsResponseBodyDomains {
	s.Filing = v
	return s
}

func (s *ListDomainsResponseBodyDomains) SetInstanceId(v string) *ListDomainsResponseBodyDomains {
	s.InstanceId = &v
	return s
}

func (s *ListDomainsResponseBodyDomains) SetLockMode(v string) *ListDomainsResponseBodyDomains {
	s.LockMode = &v
	return s
}

func (s *ListDomainsResponseBodyDomains) SetUpdateTime(v int64) *ListDomainsResponseBodyDomains {
	s.UpdateTime = &v
	return s
}

func (s *ListDomainsResponseBodyDomains) Validate() error {
	return dara.Validate(s)
}

type ListDomainsResponseBodyDomainsFiling struct {
	// The ICP number associated with the domain name. Both the entity ICP number and website ICP number are supported.
	//
	// example:
	//
	// Zhexx-xxxxxx
	IcpNumber *string `json:"IcpNumber,omitempty" xml:"IcpNumber,omitempty"`
}

func (s ListDomainsResponseBodyDomainsFiling) String() string {
	return dara.Prettify(s)
}

func (s ListDomainsResponseBodyDomainsFiling) GoString() string {
	return s.String()
}

func (s *ListDomainsResponseBodyDomainsFiling) GetIcpNumber() *string {
	return s.IcpNumber
}

func (s *ListDomainsResponseBodyDomainsFiling) SetIcpNumber(v string) *ListDomainsResponseBodyDomainsFiling {
	s.IcpNumber = &v
	return s
}

func (s *ListDomainsResponseBodyDomainsFiling) Validate() error {
	return dara.Validate(s)
}
