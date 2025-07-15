// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainStreamTranscodeDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveDomainStreamTranscodeDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveDomainStreamTranscodeDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeLiveDomainStreamTranscodeDataRequest
	GetInterval() *string
	SetOwnerId(v int64) *DescribeLiveDomainStreamTranscodeDataRequest
	GetOwnerId() *int64
	SetPrecision(v string) *DescribeLiveDomainStreamTranscodeDataRequest
	GetPrecision() *string
	SetRegionId(v string) *DescribeLiveDomainStreamTranscodeDataRequest
	GetRegionId() *string
	SetSplit(v string) *DescribeLiveDomainStreamTranscodeDataRequest
	GetSplit() *string
	SetStartTime(v string) *DescribeLiveDomainStreamTranscodeDataRequest
	GetStartTime() *string
}

type DescribeLiveDomainStreamTranscodeDataRequest struct {
	// The main streaming domain to query.
	//
	// 	- You can query one or more domain names. If you specify multiple domain names, separate them with commas (,).
	//
	// 	- If you leave this parameter empty, the data of all domain names within your Alibaba Cloud account is returned.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2017-12-10T22:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity of the query. Unit: seconds. Valid values:
	//
	// 	- **3600**: 1 hour
	//
	// 	- **86400**: 1 day
	//
	// >  If you do not specify this parameter, the time granularity of 1 hour is used by default.
	//
	// example:
	//
	// 3600
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The time precision of the query. Valid values:
	//
	// 	- **min*	- (default): in minutes.
	//
	// 	- **sec**: in seconds.
	//
	// example:
	//
	// min
	Precision *string `json:"Precision,omitempty" xml:"Precision,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The key that is used to group data. Valid values:
	//
	// 	- **domain**: The DomainName parameter is available in the response only if Split is set to domain.
	//
	// 	- **region**: The Region parameter is available in the response only if Split is set to region.
	//
	// 	- **transcode_type**: The TanscodeType parameter is available in the response only if Split is set to transcode_type.
	//
	// 	- **resolution**: The Resolution parameter is available in the response only if Split is set to resolution.
	//
	// 	- **fps**: The Fps parameter is available in the response only if Split is set to fps.
	//
	// You can specify one or more keys. If you specify multiple keys, separate them with commas (,).
	//
	// Default value: `domain,region,transcode_type,resolution,fps`.
	//
	// example:
	//
	// domain
	Split *string `json:"Split,omitempty" xml:"Split,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// 	- The minimum query interval is 1 hour.
	//
	// 	- If you do not set this parameter, the transcoding length for the last 24 hours is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2017-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLiveDomainStreamTranscodeDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainStreamTranscodeDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainStreamTranscodeDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDomainStreamTranscodeDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveDomainStreamTranscodeDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeLiveDomainStreamTranscodeDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveDomainStreamTranscodeDataRequest) GetPrecision() *string {
	return s.Precision
}

func (s *DescribeLiveDomainStreamTranscodeDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveDomainStreamTranscodeDataRequest) GetSplit() *string {
	return s.Split
}

func (s *DescribeLiveDomainStreamTranscodeDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveDomainStreamTranscodeDataRequest) SetDomainName(v string) *DescribeLiveDomainStreamTranscodeDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainStreamTranscodeDataRequest) SetEndTime(v string) *DescribeLiveDomainStreamTranscodeDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDomainStreamTranscodeDataRequest) SetInterval(v string) *DescribeLiveDomainStreamTranscodeDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeLiveDomainStreamTranscodeDataRequest) SetOwnerId(v int64) *DescribeLiveDomainStreamTranscodeDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveDomainStreamTranscodeDataRequest) SetPrecision(v string) *DescribeLiveDomainStreamTranscodeDataRequest {
	s.Precision = &v
	return s
}

func (s *DescribeLiveDomainStreamTranscodeDataRequest) SetRegionId(v string) *DescribeLiveDomainStreamTranscodeDataRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveDomainStreamTranscodeDataRequest) SetSplit(v string) *DescribeLiveDomainStreamTranscodeDataRequest {
	s.Split = &v
	return s
}

func (s *DescribeLiveDomainStreamTranscodeDataRequest) SetStartTime(v string) *DescribeLiveDomainStreamTranscodeDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainStreamTranscodeDataRequest) Validate() error {
	return dara.Validate(s)
}
