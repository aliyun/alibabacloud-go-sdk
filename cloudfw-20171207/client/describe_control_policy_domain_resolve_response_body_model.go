// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeControlPolicyDomainResolveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeControlPolicyDomainResolveResponseBody
	GetRequestId() *string
	SetResolveResult(v []*DescribeControlPolicyDomainResolveResponseBodyResolveResult) *DescribeControlPolicyDomainResolveResponseBody
	GetResolveResult() []*DescribeControlPolicyDomainResolveResponseBodyResolveResult
}

type DescribeControlPolicyDomainResolveResponseBody struct {
	// example:
	//
	// F0F82705-CFC7-5F83-86C8-A063892F****
	RequestId     *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResolveResult []*DescribeControlPolicyDomainResolveResponseBodyResolveResult `json:"ResolveResult,omitempty" xml:"ResolveResult,omitempty" type:"Repeated"`
}

func (s DescribeControlPolicyDomainResolveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeControlPolicyDomainResolveResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeControlPolicyDomainResolveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeControlPolicyDomainResolveResponseBody) GetResolveResult() []*DescribeControlPolicyDomainResolveResponseBodyResolveResult {
	return s.ResolveResult
}

func (s *DescribeControlPolicyDomainResolveResponseBody) SetRequestId(v string) *DescribeControlPolicyDomainResolveResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeControlPolicyDomainResolveResponseBody) SetResolveResult(v []*DescribeControlPolicyDomainResolveResponseBodyResolveResult) *DescribeControlPolicyDomainResolveResponseBody {
	s.ResolveResult = v
	return s
}

func (s *DescribeControlPolicyDomainResolveResponseBody) Validate() error {
	if s.ResolveResult != nil {
		for _, item := range s.ResolveResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeControlPolicyDomainResolveResponseBodyResolveResult struct {
	// example:
	//
	// example.com
	Domain     *string   `json:"Domain,omitempty" xml:"Domain,omitempty"`
	IpAddrList []*string `json:"IpAddrList,omitempty" xml:"IpAddrList,omitempty" type:"Repeated"`
	// example:
	//
	// 4
	IpVersion *int32 `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// example:
	//
	// 1579091739
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeControlPolicyDomainResolveResponseBodyResolveResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeControlPolicyDomainResolveResponseBodyResolveResult) GoString() string {
	return s.String()
}

func (s *DescribeControlPolicyDomainResolveResponseBodyResolveResult) GetDomain() *string {
	return s.Domain
}

func (s *DescribeControlPolicyDomainResolveResponseBodyResolveResult) GetIpAddrList() []*string {
	return s.IpAddrList
}

func (s *DescribeControlPolicyDomainResolveResponseBodyResolveResult) GetIpVersion() *int32 {
	return s.IpVersion
}

func (s *DescribeControlPolicyDomainResolveResponseBodyResolveResult) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *DescribeControlPolicyDomainResolveResponseBodyResolveResult) SetDomain(v string) *DescribeControlPolicyDomainResolveResponseBodyResolveResult {
	s.Domain = &v
	return s
}

func (s *DescribeControlPolicyDomainResolveResponseBodyResolveResult) SetIpAddrList(v []*string) *DescribeControlPolicyDomainResolveResponseBodyResolveResult {
	s.IpAddrList = v
	return s
}

func (s *DescribeControlPolicyDomainResolveResponseBodyResolveResult) SetIpVersion(v int32) *DescribeControlPolicyDomainResolveResponseBodyResolveResult {
	s.IpVersion = &v
	return s
}

func (s *DescribeControlPolicyDomainResolveResponseBodyResolveResult) SetUpdateTime(v int64) *DescribeControlPolicyDomainResolveResponseBodyResolveResult {
	s.UpdateTime = &v
	return s
}

func (s *DescribeControlPolicyDomainResolveResponseBodyResolveResult) Validate() error {
	return dara.Validate(s)
}
