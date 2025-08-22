// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainMultiUsageDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnDomainMultiUsageDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainMultiUsageDataRequest
	GetEndTime() *string
	SetStartTime(v string) *DescribeDcdnDomainMultiUsageDataRequest
	GetStartTime() *string
}

type DescribeDcdnDomainMultiUsageDataRequest struct {
	// If this parameter is not set, data of all your accelerated domain names is queried.
	//
	// You can specify multiple domain names and separate them with commas (,).
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end time must be later than the start time.
	//
	// 	- Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// 	- The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// 	- You need to set both the start time and the end time.
	//
	// example:
	//
	// 2017-12-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The beginning of the time range to query.
	//
	// 	- Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// 	- The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// 	- The resolution of the queried data is 5 minutes.
	//
	// 	- If you do not set this parameter, data in the last 24 hours is queried.
	//
	// 	- You need to set both the start time and the end time.
	//
	// example:
	//
	// 2017-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainMultiUsageDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainMultiUsageDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainMultiUsageDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainMultiUsageDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainMultiUsageDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainMultiUsageDataRequest) SetDomainName(v string) *DescribeDcdnDomainMultiUsageDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainMultiUsageDataRequest) SetEndTime(v string) *DescribeDcdnDomainMultiUsageDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainMultiUsageDataRequest) SetStartTime(v string) *DescribeDcdnDomainMultiUsageDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainMultiUsageDataRequest) Validate() error {
	return dara.Validate(s)
}
