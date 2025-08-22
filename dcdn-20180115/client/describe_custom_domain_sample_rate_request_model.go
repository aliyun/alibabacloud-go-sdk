// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomDomainSampleRateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainNames(v string) *DescribeCustomDomainSampleRateRequest
	GetDomainNames() *string
	SetPageNumber(v int64) *DescribeCustomDomainSampleRateRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeCustomDomainSampleRateRequest
	GetPageSize() *int64
}

type DescribeCustomDomainSampleRateRequest struct {
	DomainNames *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
	PageNumber  *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeCustomDomainSampleRateRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomDomainSampleRateRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomDomainSampleRateRequest) GetDomainNames() *string {
	return s.DomainNames
}

func (s *DescribeCustomDomainSampleRateRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeCustomDomainSampleRateRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeCustomDomainSampleRateRequest) SetDomainNames(v string) *DescribeCustomDomainSampleRateRequest {
	s.DomainNames = &v
	return s
}

func (s *DescribeCustomDomainSampleRateRequest) SetPageNumber(v int64) *DescribeCustomDomainSampleRateRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCustomDomainSampleRateRequest) SetPageSize(v int64) *DescribeCustomDomainSampleRateRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCustomDomainSampleRateRequest) Validate() error {
	return dara.Validate(s)
}
