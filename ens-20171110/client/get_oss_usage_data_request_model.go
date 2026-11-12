// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssUsageDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucket(v string) *GetOssUsageDataRequest
	GetBucket() *string
	SetEndTime(v string) *GetOssUsageDataRequest
	GetEndTime() *string
	SetPeriod(v string) *GetOssUsageDataRequest
	GetPeriod() *string
	SetStartTime(v string) *GetOssUsageDataRequest
	GetStartTime() *string
}

type GetOssUsageDataRequest struct {
	// The name of the OSS logical bucket.
	//
	// example:
	//
	// ens-sink-bucketzyp1656903494
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The end time in UTC. Format: 2010-01-21T09:50:23Z.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-01-12T00:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time aggregation granularity. Unit: minutes. Default value: 5. Valid values: 5 to 1440.
	//
	// example:
	//
	// 10
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The start time in UTC. Format: 2010-01-21T09:50:23Z.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-01-11T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetOssUsageDataRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOssUsageDataRequest) GoString() string {
	return s.String()
}

func (s *GetOssUsageDataRequest) GetBucket() *string {
	return s.Bucket
}

func (s *GetOssUsageDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetOssUsageDataRequest) GetPeriod() *string {
	return s.Period
}

func (s *GetOssUsageDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetOssUsageDataRequest) SetBucket(v string) *GetOssUsageDataRequest {
	s.Bucket = &v
	return s
}

func (s *GetOssUsageDataRequest) SetEndTime(v string) *GetOssUsageDataRequest {
	s.EndTime = &v
	return s
}

func (s *GetOssUsageDataRequest) SetPeriod(v string) *GetOssUsageDataRequest {
	s.Period = &v
	return s
}

func (s *GetOssUsageDataRequest) SetStartTime(v string) *GetOssUsageDataRequest {
	s.StartTime = &v
	return s
}

func (s *GetOssUsageDataRequest) Validate() error {
	return dara.Validate(s)
}
