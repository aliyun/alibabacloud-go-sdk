// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainPvDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVodDomainPvDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainPvDataRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeVodDomainPvDataRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribeVodDomainPvDataRequest
	GetStartTime() *string
}

type DescribeVodDomainPvDataRequest struct {
	// This parameter is required.
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainPvDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainPvDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainPvDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainPvDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainPvDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodDomainPvDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainPvDataRequest) SetDomainName(v string) *DescribeVodDomainPvDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainPvDataRequest) SetEndTime(v string) *DescribeVodDomainPvDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainPvDataRequest) SetOwnerId(v int64) *DescribeVodDomainPvDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodDomainPvDataRequest) SetStartTime(v string) *DescribeVodDomainPvDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainPvDataRequest) Validate() error {
	return dara.Validate(s)
}
