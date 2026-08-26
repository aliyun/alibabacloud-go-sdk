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
	// The streaming domain of the streamer to query.
	//
	// - You can query a single domain name or multiple domain names at a time. Separate multiple domain names with commas (,).
	//
	// - If this parameter is left empty, the merged data of all live streaming domain names is returned by default.
	//
	// - When you specify DomainName, make sure that the specified domain name is a live streaming domain name and that the user calling this operation has the permissions to operate on the specified domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2017-12-10T22:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity of the query. Unit: seconds. Valid values:
	//
	// - **3600**: by hour.
	//
	// - **86400**: by day.
	//
	// > If this parameter is left empty, the default granularity is by hour.
	//
	// example:
	//
	// 3600
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The time precision of the query. Valid values:
	//
	// - **min*	- (default): in minutes.
	//
	// - **sec**: in seconds.
	//
	// example:
	//
	// min
	Precision *string `json:"Precision,omitempty" xml:"Precision,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The grouping key. Valid values:
	//
	// - **domain**: domain name. If the Split (grouping key) parameter is set to domain, the Domain response parameter takes effect.
	//
	// - **region**: live center region. If the Split (grouping key) parameter is set to region, the Region response parameter takes effect.
	//
	// - **transcode_type**: transcoding type. If the Split (grouping key) parameter is set to transcode_type, the TanscodeType response parameter takes effect.
	//
	// - **resolution**: resolution. If the Split (grouping key) parameter is set to resolution, the Resolution response parameter takes effect.
	//
	// - **fps**: frame rate. If the Split (grouping key) parameter is set to fps, the Fps response parameter takes effect.
	//
	// You can specify one or more values. Separate multiple values with commas (,).
	//
	// Default value: `domain,region,transcode_type,resolution,fps`, which means all grouping keys are applied.
	//
	// example:
	//
	// domain
	Split *string `json:"Split,omitempty" xml:"Split,omitempty"`
	// The beginning of the time range to query. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format in UTC.
	//
	// - The minimum data time granularity is 1 hour.
	//
	// - If this parameter is left empty, data from the last 24 hours is read by default.
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
