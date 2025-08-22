// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOutPutDomains(v []*DescribeDcdnWafDomainResponseBodyOutPutDomains) *DescribeDcdnWafDomainResponseBody
	GetOutPutDomains() []*DescribeDcdnWafDomainResponseBodyOutPutDomains
	SetRequestId(v string) *DescribeDcdnWafDomainResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDcdnWafDomainResponseBody
	GetTotalCount() *int32
}

type DescribeDcdnWafDomainResponseBody struct {
	// The accelerated domain name.
	OutPutDomains []*DescribeDcdnWafDomainResponseBodyOutPutDomains `json:"OutPutDomains,omitempty" xml:"OutPutDomains,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-802B-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of accelerated domain names returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDcdnWafDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafDomainResponseBody) GetOutPutDomains() []*DescribeDcdnWafDomainResponseBodyOutPutDomains {
	return s.OutPutDomains
}

func (s *DescribeDcdnWafDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnWafDomainResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDcdnWafDomainResponseBody) SetOutPutDomains(v []*DescribeDcdnWafDomainResponseBodyOutPutDomains) *DescribeDcdnWafDomainResponseBody {
	s.OutPutDomains = v
	return s
}

func (s *DescribeDcdnWafDomainResponseBody) SetRequestId(v string) *DescribeDcdnWafDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnWafDomainResponseBody) SetTotalCount(v int32) *DescribeDcdnWafDomainResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDcdnWafDomainResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnWafDomainResponseBodyOutPutDomains struct {
	// The status of the ACL. Valid values:
	//
	// 	- 0: disabled
	//
	// 	- 1: enabled
	//
	// example:
	//
	// 1
	AclStatus *int32 `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	// The status of protection against HTTP flood attacks. Valid values:
	//
	// 	- 0: disabled
	//
	// 	- 1: enabled
	//
	// example:
	//
	// 1
	CcStatus *int32 `json:"CcStatus,omitempty" xml:"CcStatus,omitempty"`
	// The domain name that has WAF enabled.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The status of the domain name. Valid values:
	//
	// 	- 1: The domain name is added to WAF, or the domain name is valid.
	//
	// 	- 10: The domain name is being added to WAF.
	//
	// 	- 11: The domain name failed to be added to WAF.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The status of WAF. Valid values:
	//
	// 	- 0: disabled
	//
	// 	- 1: enabled
	//
	// example:
	//
	// 1
	WafStatus *int32 `json:"WafStatus,omitempty" xml:"WafStatus,omitempty"`
}

func (s DescribeDcdnWafDomainResponseBodyOutPutDomains) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafDomainResponseBodyOutPutDomains) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafDomainResponseBodyOutPutDomains) GetAclStatus() *int32 {
	return s.AclStatus
}

func (s *DescribeDcdnWafDomainResponseBodyOutPutDomains) GetCcStatus() *int32 {
	return s.CcStatus
}

func (s *DescribeDcdnWafDomainResponseBodyOutPutDomains) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDcdnWafDomainResponseBodyOutPutDomains) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeDcdnWafDomainResponseBodyOutPutDomains) GetWafStatus() *int32 {
	return s.WafStatus
}

func (s *DescribeDcdnWafDomainResponseBodyOutPutDomains) SetAclStatus(v int32) *DescribeDcdnWafDomainResponseBodyOutPutDomains {
	s.AclStatus = &v
	return s
}

func (s *DescribeDcdnWafDomainResponseBodyOutPutDomains) SetCcStatus(v int32) *DescribeDcdnWafDomainResponseBodyOutPutDomains {
	s.CcStatus = &v
	return s
}

func (s *DescribeDcdnWafDomainResponseBodyOutPutDomains) SetDomain(v string) *DescribeDcdnWafDomainResponseBodyOutPutDomains {
	s.Domain = &v
	return s
}

func (s *DescribeDcdnWafDomainResponseBodyOutPutDomains) SetStatus(v int32) *DescribeDcdnWafDomainResponseBodyOutPutDomains {
	s.Status = &v
	return s
}

func (s *DescribeDcdnWafDomainResponseBodyOutPutDomains) SetWafStatus(v int32) *DescribeDcdnWafDomainResponseBodyOutPutDomains {
	s.WafStatus = &v
	return s
}

func (s *DescribeDcdnWafDomainResponseBodyOutPutDomains) Validate() error {
	return dara.Validate(s)
}
