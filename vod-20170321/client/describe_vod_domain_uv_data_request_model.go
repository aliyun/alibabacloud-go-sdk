// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainUvDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVodDomainUvDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainUvDataRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeVodDomainUvDataRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribeVodDomainUvDataRequest
	GetStartTime() *string
}

type DescribeVodDomainUvDataRequest struct {
	// This parameter is required.
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainUvDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainUvDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainUvDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainUvDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainUvDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodDomainUvDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainUvDataRequest) SetDomainName(v string) *DescribeVodDomainUvDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainUvDataRequest) SetEndTime(v string) *DescribeVodDomainUvDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainUvDataRequest) SetOwnerId(v int64) *DescribeVodDomainUvDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodDomainUvDataRequest) SetStartTime(v string) *DescribeVodDomainUvDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainUvDataRequest) Validate() error {
	return dara.Validate(s)
}
