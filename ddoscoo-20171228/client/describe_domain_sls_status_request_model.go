// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainSlsStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeDomainSlsStatusRequest
	GetDomain() *string
	SetLang(v string) *DescribeDomainSlsStatusRequest
	GetLang() *string
	SetResourceGroupId(v string) *DescribeDomainSlsStatusRequest
	GetResourceGroupId() *string
	SetSourceIp(v string) *DescribeDomainSlsStatusRequest
	GetSourceIp() *string
}

type DescribeDomainSlsStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// xx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeDomainSlsStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSlsStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainSlsStatusRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainSlsStatusRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDomainSlsStatusRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDomainSlsStatusRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeDomainSlsStatusRequest) SetDomain(v string) *DescribeDomainSlsStatusRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainSlsStatusRequest) SetLang(v string) *DescribeDomainSlsStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDomainSlsStatusRequest) SetResourceGroupId(v string) *DescribeDomainSlsStatusRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainSlsStatusRequest) SetSourceIp(v string) *DescribeDomainSlsStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainSlsStatusRequest) Validate() error {
	return dara.Validate(s)
}
