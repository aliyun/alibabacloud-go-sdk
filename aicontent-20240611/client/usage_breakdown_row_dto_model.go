// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUsageBreakdownRowDTO interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyId(v int64) *UsageBreakdownRowDTO
	GetApiKeyId() *int64
	SetApiKeyName(v string) *UsageBreakdownRowDTO
	GetApiKeyName() *string
	SetClientId(v int64) *UsageBreakdownRowDTO
	GetClientId() *int64
	SetClientName(v string) *UsageBreakdownRowDTO
	GetClientName() *string
	SetMemberUserId(v int64) *UsageBreakdownRowDTO
	GetMemberUserId() *int64
	SetMemberUserName(v string) *UsageBreakdownRowDTO
	GetMemberUserName() *string
	SetMetrics(v []*MetricKVPairDTO) *UsageBreakdownRowDTO
	GetMetrics() []*MetricKVPairDTO
	SetModelCode(v string) *UsageBreakdownRowDTO
	GetModelCode() *string
	SetModelId(v int64) *UsageBreakdownRowDTO
	GetModelId() *int64
	SetModelName(v string) *UsageBreakdownRowDTO
	GetModelName() *string
	SetModelType(v string) *UsageBreakdownRowDTO
	GetModelType() *string
	SetSummaryTime(v int64) *UsageBreakdownRowDTO
	GetSummaryTime() *int64
}

type UsageBreakdownRowDTO struct {
	// The API key ID. A value of 0 indicates that historical data is not broken down by API key.
	//
	// example:
	//
	// 0
	ApiKeyId *int64 `json:"apiKeyId,omitempty" xml:"apiKeyId,omitempty"`
	// The API key name, corresponding to api_key_id.
	//
	// example:
	//
	// Default key
	ApiKeyName *string `json:"apiKeyName,omitempty" xml:"apiKeyName,omitempty"`
	// The department ID. A value of 0 indicates no affiliated department.
	//
	// example:
	//
	// 0
	ClientId *int64 `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// The department name.
	//
	// example:
	//
	// R&D Department
	ClientName *string `json:"clientName,omitempty" xml:"clientName,omitempty"`
	// The member ID for a member row. The value is 0 for a department row.
	//
	// example:
	//
	// 30001
	MemberUserId *int64 `json:"memberUserId,omitempty" xml:"memberUserId,omitempty"`
	// The member name for a member row. The value is empty for a department row.
	//
	// example:
	//
	// John Smith
	MemberUserName *string `json:"memberUserName,omitempty" xml:"memberUserName,omitempty"`
	// The usage metric array. Only entries with non-zero values are included.
	//
	// example:
	//
	// [{"key": "total_calls", "value": 100}]
	Metrics []*MetricKVPairDTO `json:"metrics,omitempty" xml:"metrics,omitempty" type:"Repeated"`
	// The model identifier.
	//
	// example:
	//
	// qwen-plus
	ModelCode *string `json:"modelCode,omitempty" xml:"modelCode,omitempty"`
	// The model ID.
	//
	// example:
	//
	// 1
	ModelId *int64 `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// The model name.
	//
	// example:
	//
	// Qwen-Plus
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	// The statistical dimension.
	//
	// example:
	//
	// llm
	ModelType *string `json:"modelType,omitempty" xml:"modelType,omitempty"`
	// The statistical time point, in UNIX timestamp (seconds).
	//
	// example:
	//
	// 1700000000
	SummaryTime *int64 `json:"summaryTime,omitempty" xml:"summaryTime,omitempty"`
}

func (s UsageBreakdownRowDTO) String() string {
	return dara.Prettify(s)
}

func (s UsageBreakdownRowDTO) GoString() string {
	return s.String()
}

func (s *UsageBreakdownRowDTO) GetApiKeyId() *int64 {
	return s.ApiKeyId
}

func (s *UsageBreakdownRowDTO) GetApiKeyName() *string {
	return s.ApiKeyName
}

func (s *UsageBreakdownRowDTO) GetClientId() *int64 {
	return s.ClientId
}

func (s *UsageBreakdownRowDTO) GetClientName() *string {
	return s.ClientName
}

func (s *UsageBreakdownRowDTO) GetMemberUserId() *int64 {
	return s.MemberUserId
}

func (s *UsageBreakdownRowDTO) GetMemberUserName() *string {
	return s.MemberUserName
}

func (s *UsageBreakdownRowDTO) GetMetrics() []*MetricKVPairDTO {
	return s.Metrics
}

func (s *UsageBreakdownRowDTO) GetModelCode() *string {
	return s.ModelCode
}

func (s *UsageBreakdownRowDTO) GetModelId() *int64 {
	return s.ModelId
}

func (s *UsageBreakdownRowDTO) GetModelName() *string {
	return s.ModelName
}

func (s *UsageBreakdownRowDTO) GetModelType() *string {
	return s.ModelType
}

func (s *UsageBreakdownRowDTO) GetSummaryTime() *int64 {
	return s.SummaryTime
}

func (s *UsageBreakdownRowDTO) SetApiKeyId(v int64) *UsageBreakdownRowDTO {
	s.ApiKeyId = &v
	return s
}

func (s *UsageBreakdownRowDTO) SetApiKeyName(v string) *UsageBreakdownRowDTO {
	s.ApiKeyName = &v
	return s
}

func (s *UsageBreakdownRowDTO) SetClientId(v int64) *UsageBreakdownRowDTO {
	s.ClientId = &v
	return s
}

func (s *UsageBreakdownRowDTO) SetClientName(v string) *UsageBreakdownRowDTO {
	s.ClientName = &v
	return s
}

func (s *UsageBreakdownRowDTO) SetMemberUserId(v int64) *UsageBreakdownRowDTO {
	s.MemberUserId = &v
	return s
}

func (s *UsageBreakdownRowDTO) SetMemberUserName(v string) *UsageBreakdownRowDTO {
	s.MemberUserName = &v
	return s
}

func (s *UsageBreakdownRowDTO) SetMetrics(v []*MetricKVPairDTO) *UsageBreakdownRowDTO {
	s.Metrics = v
	return s
}

func (s *UsageBreakdownRowDTO) SetModelCode(v string) *UsageBreakdownRowDTO {
	s.ModelCode = &v
	return s
}

func (s *UsageBreakdownRowDTO) SetModelId(v int64) *UsageBreakdownRowDTO {
	s.ModelId = &v
	return s
}

func (s *UsageBreakdownRowDTO) SetModelName(v string) *UsageBreakdownRowDTO {
	s.ModelName = &v
	return s
}

func (s *UsageBreakdownRowDTO) SetModelType(v string) *UsageBreakdownRowDTO {
	s.ModelType = &v
	return s
}

func (s *UsageBreakdownRowDTO) SetSummaryTime(v int64) *UsageBreakdownRowDTO {
	s.SummaryTime = &v
	return s
}

func (s *UsageBreakdownRowDTO) Validate() error {
	if s.Metrics != nil {
		for _, item := range s.Metrics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
