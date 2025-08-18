// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDoSL7QpsListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v int32) *DescribeDDoSL7QpsListResponseBody
	GetDataInterval() *int32
	SetDataModule(v []*DescribeDDoSL7QpsListResponseBodyDataModule) *DescribeDDoSL7QpsListResponseBody
	GetDataModule() []*DescribeDDoSL7QpsListResponseBodyDataModule
	SetEndTime(v string) *DescribeDDoSL7QpsListResponseBody
	GetEndTime() *string
	SetRecordId(v int64) *DescribeDDoSL7QpsListResponseBody
	GetRecordId() *int64
	SetRequestId(v string) *DescribeDDoSL7QpsListResponseBody
	GetRequestId() *string
	SetSiteId(v int64) *DescribeDDoSL7QpsListResponseBody
	GetSiteId() *int64
	SetStartTime(v string) *DescribeDDoSL7QpsListResponseBody
	GetStartTime() *string
}

type DescribeDDoSL7QpsListResponseBody struct {
	// The time granularity of the queried data, in seconds.
	//
	// example:
	//
	// 300
	DataInterval *int32 `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	// Application layer time trend data list.
	DataModule []*DescribeDDoSL7QpsListResponseBodyDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
	// The end time of the query.
	//
	// The date format follows ISO8601 notation and uses UTC+0, formatted as yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2023-04-19T19:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Record ID.
	//
	// example:
	//
	// 86510927836942****
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 156A6B-677B1A-4297B7-9187B7-2B44792
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Site ID.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The start time of the query.
	//
	// The date format follows ISO8601 notation and uses UTC+0, formatted as yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2023-04-19T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDDoSL7QpsListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDoSL7QpsListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDDoSL7QpsListResponseBody) GetDataInterval() *int32 {
	return s.DataInterval
}

func (s *DescribeDDoSL7QpsListResponseBody) GetDataModule() []*DescribeDDoSL7QpsListResponseBodyDataModule {
	return s.DataModule
}

func (s *DescribeDDoSL7QpsListResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDDoSL7QpsListResponseBody) GetRecordId() *int64 {
	return s.RecordId
}

func (s *DescribeDDoSL7QpsListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDDoSL7QpsListResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DescribeDDoSL7QpsListResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDDoSL7QpsListResponseBody) SetDataInterval(v int32) *DescribeDDoSL7QpsListResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDDoSL7QpsListResponseBody) SetDataModule(v []*DescribeDDoSL7QpsListResponseBodyDataModule) *DescribeDDoSL7QpsListResponseBody {
	s.DataModule = v
	return s
}

func (s *DescribeDDoSL7QpsListResponseBody) SetEndTime(v string) *DescribeDDoSL7QpsListResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDDoSL7QpsListResponseBody) SetRecordId(v int64) *DescribeDDoSL7QpsListResponseBody {
	s.RecordId = &v
	return s
}

func (s *DescribeDDoSL7QpsListResponseBody) SetRequestId(v string) *DescribeDDoSL7QpsListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDDoSL7QpsListResponseBody) SetSiteId(v int64) *DescribeDDoSL7QpsListResponseBody {
	s.SiteId = &v
	return s
}

func (s *DescribeDDoSL7QpsListResponseBody) SetStartTime(v string) *DescribeDDoSL7QpsListResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDDoSL7QpsListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDDoSL7QpsListResponseBodyDataModule struct {
	// Attack QPS.
	//
	// example:
	//
	// 5
	Attack *int64 `json:"Attack,omitempty" xml:"Attack,omitempty"`
	// Normal QPS.
	//
	// example:
	//
	// 4
	Normal *int64 `json:"Normal,omitempty" xml:"Normal,omitempty"`
	// Data time, following ISO8601 notation and using UTC+0, formatted as yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2023-04-19T16:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// Total QPS.
	//
	// example:
	//
	// 9
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDDoSL7QpsListResponseBodyDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDoSL7QpsListResponseBodyDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDDoSL7QpsListResponseBodyDataModule) GetAttack() *int64 {
	return s.Attack
}

func (s *DescribeDDoSL7QpsListResponseBodyDataModule) GetNormal() *int64 {
	return s.Normal
}

func (s *DescribeDDoSL7QpsListResponseBodyDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDDoSL7QpsListResponseBodyDataModule) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeDDoSL7QpsListResponseBodyDataModule) SetAttack(v int64) *DescribeDDoSL7QpsListResponseBodyDataModule {
	s.Attack = &v
	return s
}

func (s *DescribeDDoSL7QpsListResponseBodyDataModule) SetNormal(v int64) *DescribeDDoSL7QpsListResponseBodyDataModule {
	s.Normal = &v
	return s
}

func (s *DescribeDDoSL7QpsListResponseBodyDataModule) SetTimeStamp(v string) *DescribeDDoSL7QpsListResponseBodyDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDDoSL7QpsListResponseBodyDataModule) SetTotal(v int64) *DescribeDDoSL7QpsListResponseBodyDataModule {
	s.Total = &v
	return s
}

func (s *DescribeDDoSL7QpsListResponseBodyDataModule) Validate() error {
	return dara.Validate(s)
}
