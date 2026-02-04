// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppDomainDnsRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *DescribeAppDomainDnsRecordRequest
	GetBizId() *string
	SetDomainName(v string) *DescribeAppDomainDnsRecordRequest
	GetDomainName() *string
	SetPurpose(v string) *DescribeAppDomainDnsRecordRequest
	GetPurpose() *string
}

type DescribeAppDomainDnsRecordRequest struct {
	// Business ID
	//
	// example:
	//
	// WD20250821114240000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// Domain name
	//
	// example:
	//
	// rayihealth.cn
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Purpose for querying DNS configuration
	//
	// example:
	//
	// restore
	Purpose *string `json:"Purpose,omitempty" xml:"Purpose,omitempty"`
}

func (s DescribeAppDomainDnsRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppDomainDnsRecordRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppDomainDnsRecordRequest) GetBizId() *string {
	return s.BizId
}

func (s *DescribeAppDomainDnsRecordRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeAppDomainDnsRecordRequest) GetPurpose() *string {
	return s.Purpose
}

func (s *DescribeAppDomainDnsRecordRequest) SetBizId(v string) *DescribeAppDomainDnsRecordRequest {
	s.BizId = &v
	return s
}

func (s *DescribeAppDomainDnsRecordRequest) SetDomainName(v string) *DescribeAppDomainDnsRecordRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeAppDomainDnsRecordRequest) SetPurpose(v string) *DescribeAppDomainDnsRecordRequest {
	s.Purpose = &v
	return s
}

func (s *DescribeAppDomainDnsRecordRequest) Validate() error {
	return dara.Validate(s)
}
