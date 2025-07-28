// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeDomainDetailRequest
	GetDomain() *string
	SetDomainId(v string) *DescribeDomainDetailRequest
	GetDomainId() *string
	SetInstanceId(v string) *DescribeDomainDetailRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeDomainDetailRequest
	GetRegionId() *string
}

type DescribeDomainDetailRequest struct {
	// The domain name that you want to query.
	//
	// example:
	//
	// www.aliyundoc.com
	Domain   *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to obtain the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// 	- **cn-hangzhou:*	- the Chinese mainland.
	//
	// 	- **ap-southeast-1:*	- outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDomainDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainDetailRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainDetailRequest) GetDomainId() *string {
	return s.DomainId
}

func (s *DescribeDomainDetailRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDomainDetailRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDomainDetailRequest) SetDomain(v string) *DescribeDomainDetailRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainDetailRequest) SetDomainId(v string) *DescribeDomainDetailRequest {
	s.DomainId = &v
	return s
}

func (s *DescribeDomainDetailRequest) SetInstanceId(v string) *DescribeDomainDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDomainDetailRequest) SetRegionId(v string) *DescribeDomainDetailRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDomainDetailRequest) Validate() error {
	return dara.Validate(s)
}
