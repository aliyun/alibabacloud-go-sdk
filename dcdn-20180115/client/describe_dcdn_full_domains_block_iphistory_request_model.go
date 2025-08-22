// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnFullDomainsBlockIPHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeDcdnFullDomainsBlockIPHistoryRequest
	GetEndTime() *string
	SetIPList(v string) *DescribeDcdnFullDomainsBlockIPHistoryRequest
	GetIPList() *string
	SetStartTime(v string) *DescribeDcdnFullDomainsBlockIPHistoryRequest
	GetStartTime() *string
}

type DescribeDcdnFullDomainsBlockIPHistoryRequest struct {
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The end time must be later than the start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-04-24T19:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The IP address or CIDR block to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.XX.XX.10/24
	IPList *string `json:"IPList,omitempty" xml:"IPList,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-04-24T17:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnFullDomainsBlockIPHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnFullDomainsBlockIPHistoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnFullDomainsBlockIPHistoryRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnFullDomainsBlockIPHistoryRequest) GetIPList() *string {
	return s.IPList
}

func (s *DescribeDcdnFullDomainsBlockIPHistoryRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnFullDomainsBlockIPHistoryRequest) SetEndTime(v string) *DescribeDcdnFullDomainsBlockIPHistoryRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnFullDomainsBlockIPHistoryRequest) SetIPList(v string) *DescribeDcdnFullDomainsBlockIPHistoryRequest {
	s.IPList = &v
	return s
}

func (s *DescribeDcdnFullDomainsBlockIPHistoryRequest) SetStartTime(v string) *DescribeDcdnFullDomainsBlockIPHistoryRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnFullDomainsBlockIPHistoryRequest) Validate() error {
	return dara.Validate(s)
}
