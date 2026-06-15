// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPushStatByAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v int64) *QueryPushStatByAppRequest
	GetAppKey() *int64
	SetEndTime(v string) *QueryPushStatByAppRequest
	GetEndTime() *string
	SetGranularity(v string) *QueryPushStatByAppRequest
	GetGranularity() *string
	SetStartTime(v string) *QueryPushStatByAppRequest
	GetStartTime() *string
}

type QueryPushStatByAppRequest struct {
	// The AppKey value.
	//
	// This parameter is required.
	//
	// example:
	//
	// 23267207
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// The end time of the query. Specify the time in ISO 8601 format, YYYY-MM-DDThh:mm:ssZ.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2016-07-29T00:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The data granularity of the response. You can only query data for up to 31 days at daily granularity. Valid values:
	//
	// - **DAY**: Query data at daily granularity.
	//
	// This parameter is required.
	//
	// example:
	//
	// DAY
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// The start time of the query. Specify the time in ISO 8601 format, YYYY-MM-DDThh:mm:ssZ.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2016-07-25T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryPushStatByAppRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryPushStatByAppRequest) GoString() string {
	return s.String()
}

func (s *QueryPushStatByAppRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *QueryPushStatByAppRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *QueryPushStatByAppRequest) GetGranularity() *string {
	return s.Granularity
}

func (s *QueryPushStatByAppRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryPushStatByAppRequest) SetAppKey(v int64) *QueryPushStatByAppRequest {
	s.AppKey = &v
	return s
}

func (s *QueryPushStatByAppRequest) SetEndTime(v string) *QueryPushStatByAppRequest {
	s.EndTime = &v
	return s
}

func (s *QueryPushStatByAppRequest) SetGranularity(v string) *QueryPushStatByAppRequest {
	s.Granularity = &v
	return s
}

func (s *QueryPushStatByAppRequest) SetStartTime(v string) *QueryPushStatByAppRequest {
	s.StartTime = &v
	return s
}

func (s *QueryPushStatByAppRequest) Validate() error {
	return dara.Validate(s)
}
