// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsDomainRecordDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVsDomainRecordDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVsDomainRecordDataRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeVsDomainRecordDataRequest
	GetOwnerId() *int64
	SetRegion(v string) *DescribeVsDomainRecordDataRequest
	GetRegion() *string
	SetStartTime(v string) *DescribeVsDomainRecordDataRequest
	GetStartTime() *string
}

type DescribeVsDomainRecordDataRequest struct {
	// example:
	//
	// example.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// 2021-11-19T15:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// 2021-09-29T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVsDomainRecordDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainRecordDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainRecordDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVsDomainRecordDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVsDomainRecordDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVsDomainRecordDataRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeVsDomainRecordDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVsDomainRecordDataRequest) SetDomainName(v string) *DescribeVsDomainRecordDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainRecordDataRequest) SetEndTime(v string) *DescribeVsDomainRecordDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVsDomainRecordDataRequest) SetOwnerId(v int64) *DescribeVsDomainRecordDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVsDomainRecordDataRequest) SetRegion(v string) *DescribeVsDomainRecordDataRequest {
	s.Region = &v
	return s
}

func (s *DescribeVsDomainRecordDataRequest) SetStartTime(v string) *DescribeVsDomainRecordDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVsDomainRecordDataRequest) Validate() error {
	return dara.Validate(s)
}
