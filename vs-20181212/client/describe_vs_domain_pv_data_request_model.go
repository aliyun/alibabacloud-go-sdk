// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsDomainPvDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVsDomainPvDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVsDomainPvDataRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeVsDomainPvDataRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribeVsDomainPvDataRequest
	GetStartTime() *string
}

type DescribeVsDomainPvDataRequest struct {
	// The domain name for Visual Edge Computing Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the ISO 8601 standard in the \\`YYYY-MM-DDThh:mm:ssZ\\` format. The time must be in UTC.
	//
	// example:
	//
	// 2021-10-15T15:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the \\`YYYY-MM-DDThh:mm:ssZ\\` format. The time must be in UTC.<br>The minimum data granularity is 1 hour.<br>If you do not set this parameter, the data in the last 24 hours is returned.<br><br>
	//
	// example:
	//
	// 2021-10-10T17:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVsDomainPvDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainPvDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainPvDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVsDomainPvDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVsDomainPvDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVsDomainPvDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVsDomainPvDataRequest) SetDomainName(v string) *DescribeVsDomainPvDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainPvDataRequest) SetEndTime(v string) *DescribeVsDomainPvDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVsDomainPvDataRequest) SetOwnerId(v int64) *DescribeVsDomainPvDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVsDomainPvDataRequest) SetStartTime(v string) *DescribeVsDomainPvDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVsDomainPvDataRequest) Validate() error {
	return dara.Validate(s)
}
