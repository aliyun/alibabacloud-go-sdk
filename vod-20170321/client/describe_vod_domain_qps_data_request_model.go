// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainQpsDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVodDomainQpsDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainQpsDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeVodDomainQpsDataRequest
	GetInterval() *string
	SetIspNameEn(v string) *DescribeVodDomainQpsDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeVodDomainQpsDataRequest
	GetLocationNameEn() *string
	SetOwnerId(v int64) *DescribeVodDomainQpsDataRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribeVodDomainQpsDataRequest
	GetStartTime() *string
}

type DescribeVodDomainQpsDataRequest struct {
	// The accelerated domain names. Separate multiple domain names with commas (,). You can specify up to 500 domain names in a single request.
	//
	// By default, all accelerated domain names are queried.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end time. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// example:
	//
	// 2024-05-02T15:59:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity of the queried data. Unit: seconds. Valid values: **300**, **3600**, and **86400**.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The name of the Internet service provider (ISP). Default value: all ISPs.
	//
	// example:
	//
	// unicom
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The name of the region. If you do not specify this parameter, all regions are queried.
	//
	// example:
	//
	// beijing
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The start time. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// example:
	//
	// 2024-05-02T15:50:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainQpsDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainQpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainQpsDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainQpsDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainQpsDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeVodDomainQpsDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeVodDomainQpsDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeVodDomainQpsDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodDomainQpsDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainQpsDataRequest) SetDomainName(v string) *DescribeVodDomainQpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainQpsDataRequest) SetEndTime(v string) *DescribeVodDomainQpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainQpsDataRequest) SetInterval(v string) *DescribeVodDomainQpsDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeVodDomainQpsDataRequest) SetIspNameEn(v string) *DescribeVodDomainQpsDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeVodDomainQpsDataRequest) SetLocationNameEn(v string) *DescribeVodDomainQpsDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeVodDomainQpsDataRequest) SetOwnerId(v int64) *DescribeVodDomainQpsDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodDomainQpsDataRequest) SetStartTime(v string) *DescribeVodDomainQpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainQpsDataRequest) Validate() error {
	return dara.Validate(s)
}
