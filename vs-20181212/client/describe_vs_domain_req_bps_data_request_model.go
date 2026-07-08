// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsDomainReqBpsDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVsDomainReqBpsDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVsDomainReqBpsDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeVsDomainReqBpsDataRequest
	GetInterval() *string
	SetIspNameEn(v string) *DescribeVsDomainReqBpsDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeVsDomainReqBpsDataRequest
	GetLocationNameEn() *string
	SetOwnerId(v int64) *DescribeVsDomainReqBpsDataRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribeVsDomainReqBpsDataRequest
	GetStartTime() *string
}

type DescribeVsDomainReqBpsDataRequest struct {
	// Visual Edge Computing Service domain name.
	//
	// example:
	//
	// example.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// End time must be later than start time. Use ISO 8601 notation and UTC time.<br>Format: YYYY-MM-DDThh:mm:ssZ<br>
	//
	// example:
	//
	// 2021-10-16T07:00:46Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Time granularity for the query, in seconds. Valid values: 300, 3600, and 86400. If you omit this parameter or specify an unsupported value, the default value 300 is used.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// ISP name in English. Get this value from the DescribeCdnRegionAndIsp operation. If you omit this parameter, the system queries data for all ISPs.
	//
	// example:
	//
	// telecom
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// Region name in English. Get this value from the DescribeCdnRegionAndIsp operation. If you omit this parameter, the system queries data for all regions.
	//
	// example:
	//
	// beijing
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Start time of the data query. Use ISO 8601 notation and UTC time.<br>Format: YYYY-MM-DDThh:mm:ssZ<br>Minimum data granularity is 5 minutes.<br>If you omit this parameter, the system reads data from the last 24 hours by default.<br><br><br>
	//
	// example:
	//
	// 2022-01-15T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVsDomainReqBpsDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainReqBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainReqBpsDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVsDomainReqBpsDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVsDomainReqBpsDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeVsDomainReqBpsDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeVsDomainReqBpsDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeVsDomainReqBpsDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVsDomainReqBpsDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVsDomainReqBpsDataRequest) SetDomainName(v string) *DescribeVsDomainReqBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainReqBpsDataRequest) SetEndTime(v string) *DescribeVsDomainReqBpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVsDomainReqBpsDataRequest) SetInterval(v string) *DescribeVsDomainReqBpsDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeVsDomainReqBpsDataRequest) SetIspNameEn(v string) *DescribeVsDomainReqBpsDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeVsDomainReqBpsDataRequest) SetLocationNameEn(v string) *DescribeVsDomainReqBpsDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeVsDomainReqBpsDataRequest) SetOwnerId(v int64) *DescribeVsDomainReqBpsDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVsDomainReqBpsDataRequest) SetStartTime(v string) *DescribeVsDomainReqBpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVsDomainReqBpsDataRequest) Validate() error {
	return dara.Validate(s)
}
