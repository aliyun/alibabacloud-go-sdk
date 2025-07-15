// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainWithIntegrityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeDomainWithIntegrityRequest
	GetEndTime() *string
	SetIntegrity(v float32) *DescribeDomainWithIntegrityRequest
	GetIntegrity() *float32
	SetOwnerId(v int64) *DescribeDomainWithIntegrityRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeDomainWithIntegrityRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeDomainWithIntegrityRequest
	GetStartTime() *string
}

type DescribeDomainWithIntegrityRequest struct {
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the ISO 8601 standard. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2017-10-10T20:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The integrity.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.95
	Integrity *float32 `json:"Integrity,omitempty" xml:"Integrity,omitempty"`
	OwnerId   *int64   `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId  *string  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2017-10-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainWithIntegrityRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainWithIntegrityRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainWithIntegrityRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainWithIntegrityRequest) GetIntegrity() *float32 {
	return s.Integrity
}

func (s *DescribeDomainWithIntegrityRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDomainWithIntegrityRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDomainWithIntegrityRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainWithIntegrityRequest) SetEndTime(v string) *DescribeDomainWithIntegrityRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainWithIntegrityRequest) SetIntegrity(v float32) *DescribeDomainWithIntegrityRequest {
	s.Integrity = &v
	return s
}

func (s *DescribeDomainWithIntegrityRequest) SetOwnerId(v int64) *DescribeDomainWithIntegrityRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainWithIntegrityRequest) SetRegionId(v string) *DescribeDomainWithIntegrityRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDomainWithIntegrityRequest) SetStartTime(v string) *DescribeDomainWithIntegrityRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainWithIntegrityRequest) Validate() error {
	return dara.Validate(s)
}
