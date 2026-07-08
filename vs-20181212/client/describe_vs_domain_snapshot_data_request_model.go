// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsDomainSnapshotDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVsDomainSnapshotDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVsDomainSnapshotDataRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeVsDomainSnapshotDataRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribeVsDomainSnapshotDataRequest
	GetStartTime() *string
}

type DescribeVsDomainSnapshotDataRequest struct {
	// Visual Edge Computing Service domain name.
	//
	// example:
	//
	// example.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// End time of the data range. Must be later than StartTime. Use ISO 8601 notation and UTC time.<br>Format: YYYY-MM-DDThh:mm:ssZ<br>
	//
	// example:
	//
	// 2021-10-18T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Start time of the data range. Use ISO 8601 notation and UTC time.<br>Format: YYYY-MM-DDThh:mm:ssZ<br>Minimum data granularity is 5 minutes.<br>If you omit this parameter, the API returns data from the last 24 hours.<br><br><br>
	//
	// example:
	//
	// 2021-10-05T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVsDomainSnapshotDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainSnapshotDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainSnapshotDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVsDomainSnapshotDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVsDomainSnapshotDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVsDomainSnapshotDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVsDomainSnapshotDataRequest) SetDomainName(v string) *DescribeVsDomainSnapshotDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainSnapshotDataRequest) SetEndTime(v string) *DescribeVsDomainSnapshotDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVsDomainSnapshotDataRequest) SetOwnerId(v int64) *DescribeVsDomainSnapshotDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVsDomainSnapshotDataRequest) SetStartTime(v string) *DescribeVsDomainSnapshotDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVsDomainSnapshotDataRequest) Validate() error {
	return dara.Validate(s)
}
