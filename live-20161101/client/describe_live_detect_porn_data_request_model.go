// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDetectPornDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApp(v string) *DescribeLiveDetectPornDataRequest
	GetApp() *string
	SetDomainName(v string) *DescribeLiveDetectPornDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveDetectPornDataRequest
	GetEndTime() *string
	SetFee(v string) *DescribeLiveDetectPornDataRequest
	GetFee() *string
	SetOwnerId(v int64) *DescribeLiveDetectPornDataRequest
	GetOwnerId() *int64
	SetRegion(v string) *DescribeLiveDetectPornDataRequest
	GetRegion() *string
	SetRegionId(v string) *DescribeLiveDetectPornDataRequest
	GetRegionId() *string
	SetScene(v string) *DescribeLiveDetectPornDataRequest
	GetScene() *string
	SetSplitBy(v string) *DescribeLiveDetectPornDataRequest
	GetSplitBy() *string
	SetStartTime(v string) *DescribeLiveDetectPornDataRequest
	GetStartTime() *string
	SetStream(v string) *DescribeLiveDetectPornDataRequest
	GetStream() *string
}

type DescribeLiveDetectPornDataRequest struct {
	// The name of the application to which the live stream belongs.
	//
	// example:
	//
	// liveApp****
	App *string `json:"App,omitempty" xml:"App,omitempty"`
	// The main streaming domain to query.
	//
	// 	- You can query one or more domain names. If you specify multiple domain names, separate them with commas (,).
	//
	// 	- If you do not specify this parameter, the data of all domain names within your Alibaba Cloud account is returned.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2017-12-10T09:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Specifies whether a quota of free image scanning is available. Valid values:
	//
	// 	- **free**: specifies that a quota of free image scanning is available.
	//
	// 	- **charge**: specifies that a quota of free image scanning is not available and fees are charged.
	//
	// example:
	//
	// free
	Fee     *string `json:"Fee,omitempty" xml:"Fee,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the domain name resides.
	//
	// example:
	//
	// cn-shanghai
	Region   *string `json:"Region,omitempty" xml:"Region,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The moderation scenario. Valid values:
	//
	// 	- **porn**: pornography detection. This is the default value.
	//
	// 	- **terrorism**: terrorism detection
	//
	// 	- **ad**: ad violation detection
	//
	// 	- **live**: undesirable scene detection
	//
	// 	- **logo**: logo detection
	//
	// example:
	//
	// porn
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The fields based on which data is grouped. Separate multiple fields with commas (,).
	//
	// > If you leave the **SplitBy*	- parameter empty, only the **TimeStamp*	- and **Count*	- parameters are returned.
	//
	// example:
	//
	// liveApp****,liveStream****
	SplitBy *string `json:"SplitBy,omitempty" xml:"SplitBy,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// >
	//
	// 	- You can query data in the last 90 days.
	//
	// 	- The minimum data granularity is 5 minutes. If you do not specify this parameter, data in the last 24 hours is queried.
	//
	// example:
	//
	// 2017-12-10T08:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the live stream.
	//
	// example:
	//
	// liveStream****
	Stream *string `json:"Stream,omitempty" xml:"Stream,omitempty"`
}

func (s DescribeLiveDetectPornDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDetectPornDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDetectPornDataRequest) GetApp() *string {
	return s.App
}

func (s *DescribeLiveDetectPornDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDetectPornDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveDetectPornDataRequest) GetFee() *string {
	return s.Fee
}

func (s *DescribeLiveDetectPornDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveDetectPornDataRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeLiveDetectPornDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveDetectPornDataRequest) GetScene() *string {
	return s.Scene
}

func (s *DescribeLiveDetectPornDataRequest) GetSplitBy() *string {
	return s.SplitBy
}

func (s *DescribeLiveDetectPornDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveDetectPornDataRequest) GetStream() *string {
	return s.Stream
}

func (s *DescribeLiveDetectPornDataRequest) SetApp(v string) *DescribeLiveDetectPornDataRequest {
	s.App = &v
	return s
}

func (s *DescribeLiveDetectPornDataRequest) SetDomainName(v string) *DescribeLiveDetectPornDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDetectPornDataRequest) SetEndTime(v string) *DescribeLiveDetectPornDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDetectPornDataRequest) SetFee(v string) *DescribeLiveDetectPornDataRequest {
	s.Fee = &v
	return s
}

func (s *DescribeLiveDetectPornDataRequest) SetOwnerId(v int64) *DescribeLiveDetectPornDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveDetectPornDataRequest) SetRegion(v string) *DescribeLiveDetectPornDataRequest {
	s.Region = &v
	return s
}

func (s *DescribeLiveDetectPornDataRequest) SetRegionId(v string) *DescribeLiveDetectPornDataRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveDetectPornDataRequest) SetScene(v string) *DescribeLiveDetectPornDataRequest {
	s.Scene = &v
	return s
}

func (s *DescribeLiveDetectPornDataRequest) SetSplitBy(v string) *DescribeLiveDetectPornDataRequest {
	s.SplitBy = &v
	return s
}

func (s *DescribeLiveDetectPornDataRequest) SetStartTime(v string) *DescribeLiveDetectPornDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDetectPornDataRequest) SetStream(v string) *DescribeLiveDetectPornDataRequest {
	s.Stream = &v
	return s
}

func (s *DescribeLiveDetectPornDataRequest) Validate() error {
	return dara.Validate(s)
}
