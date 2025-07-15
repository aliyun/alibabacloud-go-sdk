// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamsTotalCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveStreamsTotalCountRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveStreamsTotalCountRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeLiveStreamsTotalCountRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveStreamsTotalCountRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeLiveStreamsTotalCountRequest
	GetStartTime() *string
	SetTyp(v string) *DescribeLiveStreamsTotalCountRequest
	GetTyp() *string
}

type DescribeLiveStreamsTotalCountRequest struct {
	// The ingest domain or streaming domain. This parameter is required if you want to query data based on domain names. You can specify up to 10 domain names. Separate multiple domain names with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// >  The maximum time range for a query is 15 days. The end time must be earlier than the current time. Data of the current day can be queried on the next day.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-07-25T16:00:00Z
	EndTime  *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// >  You can query data in the last 18 months.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-07-24T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The type of data that you want to query. If you leave this parameter empty, data is returned by domain name. If you want to query data by UID, specify the UID for this parameter.
	//
	// example:
	//
	// aliuid
	Typ *string `json:"Typ,omitempty" xml:"Typ,omitempty"`
}

func (s DescribeLiveStreamsTotalCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamsTotalCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsTotalCountRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveStreamsTotalCountRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveStreamsTotalCountRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveStreamsTotalCountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveStreamsTotalCountRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveStreamsTotalCountRequest) GetTyp() *string {
	return s.Typ
}

func (s *DescribeLiveStreamsTotalCountRequest) SetDomainName(v string) *DescribeLiveStreamsTotalCountRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamsTotalCountRequest) SetEndTime(v string) *DescribeLiveStreamsTotalCountRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveStreamsTotalCountRequest) SetOwnerId(v int64) *DescribeLiveStreamsTotalCountRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveStreamsTotalCountRequest) SetRegionId(v string) *DescribeLiveStreamsTotalCountRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveStreamsTotalCountRequest) SetStartTime(v string) *DescribeLiveStreamsTotalCountRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveStreamsTotalCountRequest) SetTyp(v string) *DescribeLiveStreamsTotalCountRequest {
	s.Typ = &v
	return s
}

func (s *DescribeLiveStreamsTotalCountRequest) Validate() error {
	return dara.Validate(s)
}
