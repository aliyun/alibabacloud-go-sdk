// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainTopReferVisitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVodDomainTopReferVisitRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainTopReferVisitRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeVodDomainTopReferVisitRequest
	GetOwnerId() *int64
	SetPercent(v string) *DescribeVodDomainTopReferVisitRequest
	GetPercent() *string
	SetSortBy(v string) *DescribeVodDomainTopReferVisitRequest
	GetSortBy() *string
	SetStartTime(v string) *DescribeVodDomainTopReferVisitRequest
	GetStartTime() *string
}

type DescribeVodDomainTopReferVisitRequest struct {
	// This parameter is required.
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Percent    *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	SortBy     *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainTopReferVisitRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainTopReferVisitRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainTopReferVisitRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainTopReferVisitRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainTopReferVisitRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodDomainTopReferVisitRequest) GetPercent() *string {
	return s.Percent
}

func (s *DescribeVodDomainTopReferVisitRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *DescribeVodDomainTopReferVisitRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainTopReferVisitRequest) SetDomainName(v string) *DescribeVodDomainTopReferVisitRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainTopReferVisitRequest) SetEndTime(v string) *DescribeVodDomainTopReferVisitRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainTopReferVisitRequest) SetOwnerId(v int64) *DescribeVodDomainTopReferVisitRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodDomainTopReferVisitRequest) SetPercent(v string) *DescribeVodDomainTopReferVisitRequest {
	s.Percent = &v
	return s
}

func (s *DescribeVodDomainTopReferVisitRequest) SetSortBy(v string) *DescribeVodDomainTopReferVisitRequest {
	s.SortBy = &v
	return s
}

func (s *DescribeVodDomainTopReferVisitRequest) SetStartTime(v string) *DescribeVodDomainTopReferVisitRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainTopReferVisitRequest) Validate() error {
	return dara.Validate(s)
}
