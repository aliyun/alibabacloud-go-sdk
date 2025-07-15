// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainSnapshotDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveDomainSnapshotDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveDomainSnapshotDataRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeLiveDomainSnapshotDataRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveDomainSnapshotDataRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeLiveDomainSnapshotDataRequest
	GetStartTime() *string
}

type DescribeLiveDomainSnapshotDataRequest struct {
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
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2018-01-02T00:00:00Z
	EndTime  *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// >  You can query data in the last **90*	- days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2018-01-01T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLiveDomainSnapshotDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainSnapshotDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainSnapshotDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDomainSnapshotDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveDomainSnapshotDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveDomainSnapshotDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveDomainSnapshotDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveDomainSnapshotDataRequest) SetDomainName(v string) *DescribeLiveDomainSnapshotDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainSnapshotDataRequest) SetEndTime(v string) *DescribeLiveDomainSnapshotDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDomainSnapshotDataRequest) SetOwnerId(v int64) *DescribeLiveDomainSnapshotDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveDomainSnapshotDataRequest) SetRegionId(v string) *DescribeLiveDomainSnapshotDataRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveDomainSnapshotDataRequest) SetStartTime(v string) *DescribeLiveDomainSnapshotDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainSnapshotDataRequest) Validate() error {
	return dara.Validate(s)
}
