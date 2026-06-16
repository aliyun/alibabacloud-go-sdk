// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataFilter interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DataFilter
	GetEndTime() *int64
	SetMaxRecords(v int32) *DataFilter
	GetMaxRecords() *int32
	SetProvided(v map[string]interface{}) *DataFilter
	GetProvided() map[string]interface{}
	SetQuery(v string) *DataFilter
	GetQuery() *string
	SetSamplingRate(v int32) *DataFilter
	GetSamplingRate() *int32
	SetStartTime(v int64) *DataFilter
	GetStartTime() *int64
}

type DataFilter struct {
	EndTime      *int64                 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	MaxRecords   *int32                 `json:"maxRecords,omitempty" xml:"maxRecords,omitempty"`
	Provided     map[string]interface{} `json:"provided,omitempty" xml:"provided,omitempty"`
	Query        *string                `json:"query,omitempty" xml:"query,omitempty"`
	SamplingRate *int32                 `json:"samplingRate,omitempty" xml:"samplingRate,omitempty"`
	StartTime    *int64                 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s DataFilter) String() string {
	return dara.Prettify(s)
}

func (s DataFilter) GoString() string {
	return s.String()
}

func (s *DataFilter) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DataFilter) GetMaxRecords() *int32 {
	return s.MaxRecords
}

func (s *DataFilter) GetProvided() map[string]interface{} {
	return s.Provided
}

func (s *DataFilter) GetQuery() *string {
	return s.Query
}

func (s *DataFilter) GetSamplingRate() *int32 {
	return s.SamplingRate
}

func (s *DataFilter) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DataFilter) SetEndTime(v int64) *DataFilter {
	s.EndTime = &v
	return s
}

func (s *DataFilter) SetMaxRecords(v int32) *DataFilter {
	s.MaxRecords = &v
	return s
}

func (s *DataFilter) SetProvided(v map[string]interface{}) *DataFilter {
	s.Provided = v
	return s
}

func (s *DataFilter) SetQuery(v string) *DataFilter {
	s.Query = &v
	return s
}

func (s *DataFilter) SetSamplingRate(v int32) *DataFilter {
	s.SamplingRate = &v
	return s
}

func (s *DataFilter) SetStartTime(v int64) *DataFilter {
	s.StartTime = &v
	return s
}

func (s *DataFilter) Validate() error {
	return dara.Validate(s)
}
