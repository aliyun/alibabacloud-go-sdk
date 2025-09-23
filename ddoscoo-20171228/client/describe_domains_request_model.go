// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeDomainsRequest
	GetDomain() *string
	SetInstanceIds(v []*string) *DescribeDomainsRequest
	GetInstanceIds() []*string
	SetOffset(v int32) *DescribeDomainsRequest
	GetOffset() *int32
	SetPageSize(v string) *DescribeDomainsRequest
	GetPageSize() *string
	SetQueryDomainPattern(v string) *DescribeDomainsRequest
	GetQueryDomainPattern() *string
	SetResourceGroupId(v string) *DescribeDomainsRequest
	GetResourceGroupId() *string
	SetSourceIp(v string) *DescribeDomainsRequest
	GetSourceIp() *string
}

type DescribeDomainsRequest struct {
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// ddoscoo-cn-XXXXXX
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	Offset *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// fuzzy
	QueryDomainPattern *string `json:"QueryDomainPattern,omitempty" xml:"QueryDomainPattern,omitempty"`
	// example:
	//
	// test
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainsRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainsRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeDomainsRequest) GetOffset() *int32 {
	return s.Offset
}

func (s *DescribeDomainsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeDomainsRequest) GetQueryDomainPattern() *string {
	return s.QueryDomainPattern
}

func (s *DescribeDomainsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDomainsRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeDomainsRequest) SetDomain(v string) *DescribeDomainsRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainsRequest) SetInstanceIds(v []*string) *DescribeDomainsRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeDomainsRequest) SetOffset(v int32) *DescribeDomainsRequest {
	s.Offset = &v
	return s
}

func (s *DescribeDomainsRequest) SetPageSize(v string) *DescribeDomainsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainsRequest) SetQueryDomainPattern(v string) *DescribeDomainsRequest {
	s.QueryDomainPattern = &v
	return s
}

func (s *DescribeDomainsRequest) SetResourceGroupId(v string) *DescribeDomainsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainsRequest) SetSourceIp(v string) *DescribeDomainsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainsRequest) Validate() error {
	return dara.Validate(s)
}
