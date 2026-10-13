// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMOTokenUsageSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *DescribeMOTokenUsageSummaryRequest
	GetApiKey() *string
	SetEndTime(v string) *DescribeMOTokenUsageSummaryRequest
	GetEndTime() *string
	SetInstanceId(v string) *DescribeMOTokenUsageSummaryRequest
	GetInstanceId() *string
	SetModel(v string) *DescribeMOTokenUsageSummaryRequest
	GetModel() *string
	SetStartTime(v string) *DescribeMOTokenUsageSummaryRequest
	GetStartTime() *string
	SetUsageType(v string) *DescribeMOTokenUsageSummaryRequest
	GetUsageType() *string
}

type DescribeMOTokenUsageSummaryRequest struct {
	// The API key used.
	//
	// example:
	//
	// sk-rds-*****
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// The end time of the instance. Format: YYYY-MM-DD. The interval between EndTime and StartTime cannot exceed 31 days.
	//
	// example:
	//
	// 2026-04-30
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance name.
	//
	// This parameter is required.
	//
	// example:
	//
	// rds_copilotpost_public_cn-xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The model called.
	//
	// example:
	//
	// qwen-flash
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The start time of the instance. Format: YYYY-MM-DD.
	//
	// example:
	//
	// 2026-04-21
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The query type.
	//
	// example:
	//
	// text
	UsageType *string `json:"UsageType,omitempty" xml:"UsageType,omitempty"`
}

func (s DescribeMOTokenUsageSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMOTokenUsageSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeMOTokenUsageSummaryRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *DescribeMOTokenUsageSummaryRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeMOTokenUsageSummaryRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeMOTokenUsageSummaryRequest) GetModel() *string {
	return s.Model
}

func (s *DescribeMOTokenUsageSummaryRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeMOTokenUsageSummaryRequest) GetUsageType() *string {
	return s.UsageType
}

func (s *DescribeMOTokenUsageSummaryRequest) SetApiKey(v string) *DescribeMOTokenUsageSummaryRequest {
	s.ApiKey = &v
	return s
}

func (s *DescribeMOTokenUsageSummaryRequest) SetEndTime(v string) *DescribeMOTokenUsageSummaryRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeMOTokenUsageSummaryRequest) SetInstanceId(v string) *DescribeMOTokenUsageSummaryRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeMOTokenUsageSummaryRequest) SetModel(v string) *DescribeMOTokenUsageSummaryRequest {
	s.Model = &v
	return s
}

func (s *DescribeMOTokenUsageSummaryRequest) SetStartTime(v string) *DescribeMOTokenUsageSummaryRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeMOTokenUsageSummaryRequest) SetUsageType(v string) *DescribeMOTokenUsageSummaryRequest {
	s.UsageType = &v
	return s
}

func (s *DescribeMOTokenUsageSummaryRequest) Validate() error {
	return dara.Validate(s)
}
