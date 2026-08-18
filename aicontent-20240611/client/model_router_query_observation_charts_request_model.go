// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryObservationChartsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyId(v int64) *ModelRouterQueryObservationChartsRequest
	GetApiKeyId() *int64
	SetClientId(v int64) *ModelRouterQueryObservationChartsRequest
	GetClientId() *int64
	SetClientIds(v string) *ModelRouterQueryObservationChartsRequest
	GetClientIds() *string
	SetEndTime(v string) *ModelRouterQueryObservationChartsRequest
	GetEndTime() *string
	SetMemberUserIds(v string) *ModelRouterQueryObservationChartsRequest
	GetMemberUserIds() *string
	SetModelId(v int64) *ModelRouterQueryObservationChartsRequest
	GetModelId() *int64
	SetStartTime(v string) *ModelRouterQueryObservationChartsRequest
	GetStartTime() *string
	SetTimeRange(v string) *ModelRouterQueryObservationChartsRequest
	GetTimeRange() *string
}

type ModelRouterQueryObservationChartsRequest struct {
	// The API key ID used to filter results.
	//
	// example:
	//
	// 1
	ApiKeyId *int64 `json:"apiKeyId,omitempty" xml:"apiKeyId,omitempty"`
	// The client ID used to filter results.
	//
	// example:
	//
	// 1
	ClientId *int64 `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// The list of department IDs, separated by commas. Supports querying data for multiple departments. This parameter is mutually exclusive with client_id.
	//
	// example:
	//
	// 1,2,3
	ClientIds *string `json:"clientIds,omitempty" xml:"clientIds,omitempty"`
	// The custom end time.
	//
	// example:
	//
	// 2024-01-02T00:00:00Z
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// Optional. Filters by member IDs, separated by commas. If not specified, the department and all its members are included. If an empty value is specified, only the department is included without members.
	//
	// example:
	//
	// 30001,30002
	MemberUserIds *string `json:"memberUserIds,omitempty" xml:"memberUserIds,omitempty"`
	// The model ID used to filter results.
	//
	// example:
	//
	// 1
	ModelId *int64 `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// The custom start time.
	//
	// example:
	//
	// 2024-01-01T00:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The time range for the query. Valid values: 1h, 6h, 24h, 7d, and 30d.
	//
	// example:
	//
	// 24h
	TimeRange *string `json:"timeRange,omitempty" xml:"timeRange,omitempty"`
}

func (s ModelRouterQueryObservationChartsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryObservationChartsRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryObservationChartsRequest) GetApiKeyId() *int64 {
	return s.ApiKeyId
}

func (s *ModelRouterQueryObservationChartsRequest) GetClientId() *int64 {
	return s.ClientId
}

func (s *ModelRouterQueryObservationChartsRequest) GetClientIds() *string {
	return s.ClientIds
}

func (s *ModelRouterQueryObservationChartsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ModelRouterQueryObservationChartsRequest) GetMemberUserIds() *string {
	return s.MemberUserIds
}

func (s *ModelRouterQueryObservationChartsRequest) GetModelId() *int64 {
	return s.ModelId
}

func (s *ModelRouterQueryObservationChartsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ModelRouterQueryObservationChartsRequest) GetTimeRange() *string {
	return s.TimeRange
}

func (s *ModelRouterQueryObservationChartsRequest) SetApiKeyId(v int64) *ModelRouterQueryObservationChartsRequest {
	s.ApiKeyId = &v
	return s
}

func (s *ModelRouterQueryObservationChartsRequest) SetClientId(v int64) *ModelRouterQueryObservationChartsRequest {
	s.ClientId = &v
	return s
}

func (s *ModelRouterQueryObservationChartsRequest) SetClientIds(v string) *ModelRouterQueryObservationChartsRequest {
	s.ClientIds = &v
	return s
}

func (s *ModelRouterQueryObservationChartsRequest) SetEndTime(v string) *ModelRouterQueryObservationChartsRequest {
	s.EndTime = &v
	return s
}

func (s *ModelRouterQueryObservationChartsRequest) SetMemberUserIds(v string) *ModelRouterQueryObservationChartsRequest {
	s.MemberUserIds = &v
	return s
}

func (s *ModelRouterQueryObservationChartsRequest) SetModelId(v int64) *ModelRouterQueryObservationChartsRequest {
	s.ModelId = &v
	return s
}

func (s *ModelRouterQueryObservationChartsRequest) SetStartTime(v string) *ModelRouterQueryObservationChartsRequest {
	s.StartTime = &v
	return s
}

func (s *ModelRouterQueryObservationChartsRequest) SetTimeRange(v string) *ModelRouterQueryObservationChartsRequest {
	s.TimeRange = &v
	return s
}

func (s *ModelRouterQueryObservationChartsRequest) Validate() error {
	return dara.Validate(s)
}
