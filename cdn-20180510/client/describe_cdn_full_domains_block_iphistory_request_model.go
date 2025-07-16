// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnFullDomainsBlockIPHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeCdnFullDomainsBlockIPHistoryRequest
	GetEndTime() *string
	SetIPList(v string) *DescribeCdnFullDomainsBlockIPHistoryRequest
	GetIPList() *string
	SetStartTime(v string) *DescribeCdnFullDomainsBlockIPHistoryRequest
	GetStartTime() *string
}

type DescribeCdnFullDomainsBlockIPHistoryRequest struct {
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The end time must be later than the start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-04-24T19:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The IP address or CIDR blocks to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.XXX.XXX.1,2.XXX.XXX.2
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

func (s DescribeCdnFullDomainsBlockIPHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnFullDomainsBlockIPHistoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnFullDomainsBlockIPHistoryRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeCdnFullDomainsBlockIPHistoryRequest) GetIPList() *string {
	return s.IPList
}

func (s *DescribeCdnFullDomainsBlockIPHistoryRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeCdnFullDomainsBlockIPHistoryRequest) SetEndTime(v string) *DescribeCdnFullDomainsBlockIPHistoryRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeCdnFullDomainsBlockIPHistoryRequest) SetIPList(v string) *DescribeCdnFullDomainsBlockIPHistoryRequest {
	s.IPList = &v
	return s
}

func (s *DescribeCdnFullDomainsBlockIPHistoryRequest) SetStartTime(v string) *DescribeCdnFullDomainsBlockIPHistoryRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeCdnFullDomainsBlockIPHistoryRequest) Validate() error {
	return dara.Validate(s)
}
