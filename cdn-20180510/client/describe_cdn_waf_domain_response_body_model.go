// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnWafDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOutPutDomains(v []*DescribeCdnWafDomainResponseBodyOutPutDomains) *DescribeCdnWafDomainResponseBody
	GetOutPutDomains() []*DescribeCdnWafDomainResponseBodyOutPutDomains
	SetRequestId(v string) *DescribeCdnWafDomainResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeCdnWafDomainResponseBody
	GetTotalCount() *int32
}

type DescribeCdnWafDomainResponseBody struct {
	// The information about the accelerated domain name.
	OutPutDomains []*DescribeCdnWafDomainResponseBodyOutPutDomains `json:"OutPutDomains,omitempty" xml:"OutPutDomains,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-802B-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of accelerated domain names.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCdnWafDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnWafDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnWafDomainResponseBody) GetOutPutDomains() []*DescribeCdnWafDomainResponseBodyOutPutDomains {
	return s.OutPutDomains
}

func (s *DescribeCdnWafDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdnWafDomainResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCdnWafDomainResponseBody) SetOutPutDomains(v []*DescribeCdnWafDomainResponseBodyOutPutDomains) *DescribeCdnWafDomainResponseBody {
	s.OutPutDomains = v
	return s
}

func (s *DescribeCdnWafDomainResponseBody) SetRequestId(v string) *DescribeCdnWafDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnWafDomainResponseBody) SetTotalCount(v int32) *DescribeCdnWafDomainResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCdnWafDomainResponseBody) Validate() error {
	if s.OutPutDomains != nil {
		for _, item := range s.OutPutDomains {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCdnWafDomainResponseBodyOutPutDomains struct {
	// The status of the access control list (ACL) feature. Valid values:
	//
	// 	- **0**: disabled
	//
	// 	- **1**: enabled
	//
	// example:
	//
	// 1
	AclStatus *string `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	// The status of protection against HTTP flood attacks. Valid values:
	//
	// 	- **0**: disabled
	//
	// 	- **1**: enabled
	//
	// example:
	//
	// 1
	CcStatus *string `json:"CcStatus,omitempty" xml:"CcStatus,omitempty"`
	// The accelerated domain name.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The WAF status of the domain name. Valid values:
	//
	// 	- **1**: The domain name is added to WAF or valid.
	//
	// 	- **10**: The domain name is being added to WAF.
	//
	// 	- **11**: The domain name failed to be added to WAF.
	//
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The status of WAF. Valid values:
	//
	// 	- **0**: disabled
	//
	// 	- **1**: enabled
	//
	// example:
	//
	// 1
	WafStatus *string `json:"WafStatus,omitempty" xml:"WafStatus,omitempty"`
}

func (s DescribeCdnWafDomainResponseBodyOutPutDomains) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnWafDomainResponseBodyOutPutDomains) GoString() string {
	return s.String()
}

func (s *DescribeCdnWafDomainResponseBodyOutPutDomains) GetAclStatus() *string {
	return s.AclStatus
}

func (s *DescribeCdnWafDomainResponseBodyOutPutDomains) GetCcStatus() *string {
	return s.CcStatus
}

func (s *DescribeCdnWafDomainResponseBodyOutPutDomains) GetDomain() *string {
	return s.Domain
}

func (s *DescribeCdnWafDomainResponseBodyOutPutDomains) GetStatus() *string {
	return s.Status
}

func (s *DescribeCdnWafDomainResponseBodyOutPutDomains) GetWafStatus() *string {
	return s.WafStatus
}

func (s *DescribeCdnWafDomainResponseBodyOutPutDomains) SetAclStatus(v string) *DescribeCdnWafDomainResponseBodyOutPutDomains {
	s.AclStatus = &v
	return s
}

func (s *DescribeCdnWafDomainResponseBodyOutPutDomains) SetCcStatus(v string) *DescribeCdnWafDomainResponseBodyOutPutDomains {
	s.CcStatus = &v
	return s
}

func (s *DescribeCdnWafDomainResponseBodyOutPutDomains) SetDomain(v string) *DescribeCdnWafDomainResponseBodyOutPutDomains {
	s.Domain = &v
	return s
}

func (s *DescribeCdnWafDomainResponseBodyOutPutDomains) SetStatus(v string) *DescribeCdnWafDomainResponseBodyOutPutDomains {
	s.Status = &v
	return s
}

func (s *DescribeCdnWafDomainResponseBodyOutPutDomains) SetWafStatus(v string) *DescribeCdnWafDomainResponseBodyOutPutDomains {
	s.WafStatus = &v
	return s
}

func (s *DescribeCdnWafDomainResponseBodyOutPutDomains) Validate() error {
	return dara.Validate(s)
}
