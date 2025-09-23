// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLayer7CCRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeLayer7CCRulesRequest
	GetDomain() *string
	SetOffset(v int32) *DescribeLayer7CCRulesRequest
	GetOffset() *int32
	SetPageSize(v string) *DescribeLayer7CCRulesRequest
	GetPageSize() *string
	SetResourceGroupId(v string) *DescribeLayer7CCRulesRequest
	GetResourceGroupId() *string
	SetSourceIp(v string) *DescribeLayer7CCRulesRequest
	GetSourceIp() *string
}

type DescribeLayer7CCRulesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
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
	// test
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeLayer7CCRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLayer7CCRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeLayer7CCRulesRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeLayer7CCRulesRequest) GetOffset() *int32 {
	return s.Offset
}

func (s *DescribeLayer7CCRulesRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeLayer7CCRulesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeLayer7CCRulesRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeLayer7CCRulesRequest) SetDomain(v string) *DescribeLayer7CCRulesRequest {
	s.Domain = &v
	return s
}

func (s *DescribeLayer7CCRulesRequest) SetOffset(v int32) *DescribeLayer7CCRulesRequest {
	s.Offset = &v
	return s
}

func (s *DescribeLayer7CCRulesRequest) SetPageSize(v string) *DescribeLayer7CCRulesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLayer7CCRulesRequest) SetResourceGroupId(v string) *DescribeLayer7CCRulesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeLayer7CCRulesRequest) SetSourceIp(v string) *DescribeLayer7CCRulesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeLayer7CCRulesRequest) Validate() error {
	return dara.Validate(s)
}
