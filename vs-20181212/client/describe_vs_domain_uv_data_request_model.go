// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsDomainUvDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVsDomainUvDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVsDomainUvDataRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeVsDomainUvDataRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribeVsDomainUvDataRequest
	GetStartTime() *string
}

type DescribeVsDomainUvDataRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// example.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// 2021-11-24T00:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 2021-10-12T10:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVsDomainUvDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainUvDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainUvDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVsDomainUvDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVsDomainUvDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVsDomainUvDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVsDomainUvDataRequest) SetDomainName(v string) *DescribeVsDomainUvDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainUvDataRequest) SetEndTime(v string) *DescribeVsDomainUvDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVsDomainUvDataRequest) SetOwnerId(v int64) *DescribeVsDomainUvDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVsDomainUvDataRequest) SetStartTime(v string) *DescribeVsDomainUvDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVsDomainUvDataRequest) Validate() error {
	return dara.Validate(s)
}
