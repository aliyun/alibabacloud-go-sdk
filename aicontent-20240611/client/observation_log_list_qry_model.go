// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObservationLogListQry interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyId(v int64) *ObservationLogListQry
	GetApiKeyId() *int64
	SetClientId(v int64) *ObservationLogListQry
	GetClientId() *int64
	SetEndTime(v string) *ObservationLogListQry
	GetEndTime() *string
	SetModelId(v int64) *ObservationLogListQry
	GetModelId() *int64
	SetPage(v int32) *ObservationLogListQry
	GetPage() *int32
	SetPageSize(v int32) *ObservationLogListQry
	GetPageSize() *int32
	SetStartTime(v string) *ObservationLogListQry
	GetStartTime() *string
	SetTimeRange(v string) *ObservationLogListQry
	GetTimeRange() *string
}

type ObservationLogListQry struct {
	// example:
	//
	// 1
	ApiKeyId *int64 `json:"apiKeyId,omitempty" xml:"apiKeyId,omitempty"`
	// example:
	//
	// 1
	ClientId *int64 `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// example:
	//
	// 2024-01-02T00:00:00Z
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 1
	ModelId *int64 `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 2024-01-01T00:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// 24h
	TimeRange *string `json:"timeRange,omitempty" xml:"timeRange,omitempty"`
}

func (s ObservationLogListQry) String() string {
	return dara.Prettify(s)
}

func (s ObservationLogListQry) GoString() string {
	return s.String()
}

func (s *ObservationLogListQry) GetApiKeyId() *int64 {
	return s.ApiKeyId
}

func (s *ObservationLogListQry) GetClientId() *int64 {
	return s.ClientId
}

func (s *ObservationLogListQry) GetEndTime() *string {
	return s.EndTime
}

func (s *ObservationLogListQry) GetModelId() *int64 {
	return s.ModelId
}

func (s *ObservationLogListQry) GetPage() *int32 {
	return s.Page
}

func (s *ObservationLogListQry) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ObservationLogListQry) GetStartTime() *string {
	return s.StartTime
}

func (s *ObservationLogListQry) GetTimeRange() *string {
	return s.TimeRange
}

func (s *ObservationLogListQry) SetApiKeyId(v int64) *ObservationLogListQry {
	s.ApiKeyId = &v
	return s
}

func (s *ObservationLogListQry) SetClientId(v int64) *ObservationLogListQry {
	s.ClientId = &v
	return s
}

func (s *ObservationLogListQry) SetEndTime(v string) *ObservationLogListQry {
	s.EndTime = &v
	return s
}

func (s *ObservationLogListQry) SetModelId(v int64) *ObservationLogListQry {
	s.ModelId = &v
	return s
}

func (s *ObservationLogListQry) SetPage(v int32) *ObservationLogListQry {
	s.Page = &v
	return s
}

func (s *ObservationLogListQry) SetPageSize(v int32) *ObservationLogListQry {
	s.PageSize = &v
	return s
}

func (s *ObservationLogListQry) SetStartTime(v string) *ObservationLogListQry {
	s.StartTime = &v
	return s
}

func (s *ObservationLogListQry) SetTimeRange(v string) *ObservationLogListQry {
	s.TimeRange = &v
	return s
}

func (s *ObservationLogListQry) Validate() error {
	return dara.Validate(s)
}
