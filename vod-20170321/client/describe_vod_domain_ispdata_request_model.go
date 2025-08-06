// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainISPDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVodDomainISPDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainISPDataRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeVodDomainISPDataRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribeVodDomainISPDataRequest
	GetStartTime() *string
}

type DescribeVodDomainISPDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainISPDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainISPDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainISPDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainISPDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainISPDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodDomainISPDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainISPDataRequest) SetDomainName(v string) *DescribeVodDomainISPDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainISPDataRequest) SetEndTime(v string) *DescribeVodDomainISPDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainISPDataRequest) SetOwnerId(v int64) *DescribeVodDomainISPDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodDomainISPDataRequest) SetStartTime(v string) *DescribeVodDomainISPDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainISPDataRequest) Validate() error {
	return dara.Validate(s)
}
