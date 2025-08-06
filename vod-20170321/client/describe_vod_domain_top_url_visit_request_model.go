// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainTopUrlVisitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVodDomainTopUrlVisitRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainTopUrlVisitRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeVodDomainTopUrlVisitRequest
	GetOwnerId() *int64
	SetPercent(v string) *DescribeVodDomainTopUrlVisitRequest
	GetPercent() *string
	SetSortBy(v string) *DescribeVodDomainTopUrlVisitRequest
	GetSortBy() *string
	SetStartTime(v string) *DescribeVodDomainTopUrlVisitRequest
	GetStartTime() *string
}

type DescribeVodDomainTopUrlVisitRequest struct {
	// This parameter is required.
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Percent    *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	SortBy     *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainTopUrlVisitRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainTopUrlVisitRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainTopUrlVisitRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainTopUrlVisitRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainTopUrlVisitRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodDomainTopUrlVisitRequest) GetPercent() *string {
	return s.Percent
}

func (s *DescribeVodDomainTopUrlVisitRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *DescribeVodDomainTopUrlVisitRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainTopUrlVisitRequest) SetDomainName(v string) *DescribeVodDomainTopUrlVisitRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainTopUrlVisitRequest) SetEndTime(v string) *DescribeVodDomainTopUrlVisitRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainTopUrlVisitRequest) SetOwnerId(v int64) *DescribeVodDomainTopUrlVisitRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodDomainTopUrlVisitRequest) SetPercent(v string) *DescribeVodDomainTopUrlVisitRequest {
	s.Percent = &v
	return s
}

func (s *DescribeVodDomainTopUrlVisitRequest) SetSortBy(v string) *DescribeVodDomainTopUrlVisitRequest {
	s.SortBy = &v
	return s
}

func (s *DescribeVodDomainTopUrlVisitRequest) SetStartTime(v string) *DescribeVodDomainTopUrlVisitRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainTopUrlVisitRequest) Validate() error {
	return dara.Validate(s)
}
