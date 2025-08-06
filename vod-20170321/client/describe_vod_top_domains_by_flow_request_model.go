// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodTopDomainsByFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeVodTopDomainsByFlowRequest
	GetEndTime() *string
	SetLimit(v int64) *DescribeVodTopDomainsByFlowRequest
	GetLimit() *int64
	SetOwnerId(v int64) *DescribeVodTopDomainsByFlowRequest
	GetOwnerId() *int64
	SetProduct(v string) *DescribeVodTopDomainsByFlowRequest
	GetProduct() *string
	SetStartTime(v string) *DescribeVodTopDomainsByFlowRequest
	GetStartTime() *string
}

type DescribeVodTopDomainsByFlowRequest struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Limit     *int64  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Product   *string `json:"Product,omitempty" xml:"Product,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodTopDomainsByFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodTopDomainsByFlowRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodTopDomainsByFlowRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodTopDomainsByFlowRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *DescribeVodTopDomainsByFlowRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodTopDomainsByFlowRequest) GetProduct() *string {
	return s.Product
}

func (s *DescribeVodTopDomainsByFlowRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodTopDomainsByFlowRequest) SetEndTime(v string) *DescribeVodTopDomainsByFlowRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodTopDomainsByFlowRequest) SetLimit(v int64) *DescribeVodTopDomainsByFlowRequest {
	s.Limit = &v
	return s
}

func (s *DescribeVodTopDomainsByFlowRequest) SetOwnerId(v int64) *DescribeVodTopDomainsByFlowRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodTopDomainsByFlowRequest) SetProduct(v string) *DescribeVodTopDomainsByFlowRequest {
	s.Product = &v
	return s
}

func (s *DescribeVodTopDomainsByFlowRequest) SetStartTime(v string) *DescribeVodTopDomainsByFlowRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodTopDomainsByFlowRequest) Validate() error {
	return dara.Validate(s)
}
