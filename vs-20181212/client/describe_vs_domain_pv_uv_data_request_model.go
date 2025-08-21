// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsDomainPvUvDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVsDomainPvUvDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVsDomainPvUvDataRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeVsDomainPvUvDataRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribeVsDomainPvUvDataRequest
	GetStartTime() *string
}

type DescribeVsDomainPvUvDataRequest struct {
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
	// 2021-10-14T23:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVsDomainPvUvDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainPvUvDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainPvUvDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVsDomainPvUvDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVsDomainPvUvDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVsDomainPvUvDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVsDomainPvUvDataRequest) SetDomainName(v string) *DescribeVsDomainPvUvDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainPvUvDataRequest) SetEndTime(v string) *DescribeVsDomainPvUvDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVsDomainPvUvDataRequest) SetOwnerId(v int64) *DescribeVsDomainPvUvDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVsDomainPvUvDataRequest) SetStartTime(v string) *DescribeVsDomainPvUvDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVsDomainPvUvDataRequest) Validate() error {
	return dara.Validate(s)
}
