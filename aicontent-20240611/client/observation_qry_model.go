// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObservationQry interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyId(v int64) *ObservationQry
	GetApiKeyId() *int64
	SetClientId(v int64) *ObservationQry
	GetClientId() *int64
	SetEndTime(v string) *ObservationQry
	GetEndTime() *string
	SetModelId(v int64) *ObservationQry
	GetModelId() *int64
	SetStartTime(v string) *ObservationQry
	GetStartTime() *string
	SetTimeRange(v string) *ObservationQry
	GetTimeRange() *string
}

type ObservationQry struct {
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
	// 2024-01-01T00:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// 24h
	TimeRange *string `json:"timeRange,omitempty" xml:"timeRange,omitempty"`
}

func (s ObservationQry) String() string {
	return dara.Prettify(s)
}

func (s ObservationQry) GoString() string {
	return s.String()
}

func (s *ObservationQry) GetApiKeyId() *int64 {
	return s.ApiKeyId
}

func (s *ObservationQry) GetClientId() *int64 {
	return s.ClientId
}

func (s *ObservationQry) GetEndTime() *string {
	return s.EndTime
}

func (s *ObservationQry) GetModelId() *int64 {
	return s.ModelId
}

func (s *ObservationQry) GetStartTime() *string {
	return s.StartTime
}

func (s *ObservationQry) GetTimeRange() *string {
	return s.TimeRange
}

func (s *ObservationQry) SetApiKeyId(v int64) *ObservationQry {
	s.ApiKeyId = &v
	return s
}

func (s *ObservationQry) SetClientId(v int64) *ObservationQry {
	s.ClientId = &v
	return s
}

func (s *ObservationQry) SetEndTime(v string) *ObservationQry {
	s.EndTime = &v
	return s
}

func (s *ObservationQry) SetModelId(v int64) *ObservationQry {
	s.ModelId = &v
	return s
}

func (s *ObservationQry) SetStartTime(v string) *ObservationQry {
	s.StartTime = &v
	return s
}

func (s *ObservationQry) SetTimeRange(v string) *ObservationQry {
	s.TimeRange = &v
	return s
}

func (s *ObservationQry) Validate() error {
	return dara.Validate(s)
}
