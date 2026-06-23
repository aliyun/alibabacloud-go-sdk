// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDoSEventMaxRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeDDoSEventMaxRequest
	GetEndTime() *string
	SetSiteId(v int64) *DescribeDDoSEventMaxRequest
	GetSiteId() *int64
	SetStartTime(v string) *DescribeDDoSEventMaxRequest
	GetStartTime() *string
}

type DescribeDDoSEventMaxRequest struct {
	// The end of the time range to query.
	//
	// The date is in ISO 8601 format. The time is displayed in UTC. The format is yyyy-MM-ddTHH:mm:ssZ. The maximum time span between the start time and end time is 31 days.
	//
	// If you do not set this parameter, the current time is used as the end time.
	//
	// example:
	//
	// 2023-04-10T02:10:10Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The site ID, which can be obtained by calling the [ListSites](~~ListSites~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7096621098****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The beginning of the time range to query.
	//
	// The date is in ISO 8601 format. The time is displayed in UTC. The format is yyyy-MM-ddTHH:mm:ssZ.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-02-18T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDDoSEventMaxRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDoSEventMaxRequest) GoString() string {
	return s.String()
}

func (s *DescribeDDoSEventMaxRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDDoSEventMaxRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DescribeDDoSEventMaxRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDDoSEventMaxRequest) SetEndTime(v string) *DescribeDDoSEventMaxRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDDoSEventMaxRequest) SetSiteId(v int64) *DescribeDDoSEventMaxRequest {
	s.SiteId = &v
	return s
}

func (s *DescribeDDoSEventMaxRequest) SetStartTime(v string) *DescribeDDoSEventMaxRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDDoSEventMaxRequest) Validate() error {
	return dara.Validate(s)
}
