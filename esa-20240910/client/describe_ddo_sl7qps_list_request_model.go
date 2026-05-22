// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDoSL7QpsListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeDDoSL7QpsListRequest
	GetEndTime() *string
	SetInterval(v int32) *DescribeDDoSL7QpsListRequest
	GetInterval() *int32
	SetRecordId(v int64) *DescribeDDoSL7QpsListRequest
	GetRecordId() *int64
	SetSiteId(v int64) *DescribeDDoSL7QpsListRequest
	GetSiteId() *int64
	SetStartTime(v string) *DescribeDDoSL7QpsListRequest
	GetStartTime() *string
}

type DescribeDDoSL7QpsListRequest struct {
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// This parameter is required.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDDoSL7QpsListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDoSL7QpsListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDDoSL7QpsListRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDDoSL7QpsListRequest) GetInterval() *int32 {
	return s.Interval
}

func (s *DescribeDDoSL7QpsListRequest) GetRecordId() *int64 {
	return s.RecordId
}

func (s *DescribeDDoSL7QpsListRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DescribeDDoSL7QpsListRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDDoSL7QpsListRequest) SetEndTime(v string) *DescribeDDoSL7QpsListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDDoSL7QpsListRequest) SetInterval(v int32) *DescribeDDoSL7QpsListRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDDoSL7QpsListRequest) SetRecordId(v int64) *DescribeDDoSL7QpsListRequest {
	s.RecordId = &v
	return s
}

func (s *DescribeDDoSL7QpsListRequest) SetSiteId(v int64) *DescribeDDoSL7QpsListRequest {
	s.SiteId = &v
	return s
}

func (s *DescribeDDoSL7QpsListRequest) SetStartTime(v string) *DescribeDDoSL7QpsListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDDoSL7QpsListRequest) Validate() error {
	return dara.Validate(s)
}
