// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMOTokenUsageDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *DescribeMOTokenUsageDetailRequest
	GetApiKey() *string
	SetConsumerName(v string) *DescribeMOTokenUsageDetailRequest
	GetConsumerName() *string
	SetCursor(v string) *DescribeMOTokenUsageDetailRequest
	GetCursor() *string
	SetEndTime(v string) *DescribeMOTokenUsageDetailRequest
	GetEndTime() *string
	SetInstanceId(v string) *DescribeMOTokenUsageDetailRequest
	GetInstanceId() *string
	SetModel(v string) *DescribeMOTokenUsageDetailRequest
	GetModel() *string
	SetPage(v int32) *DescribeMOTokenUsageDetailRequest
	GetPage() *int32
	SetPageSize(v int32) *DescribeMOTokenUsageDetailRequest
	GetPageSize() *int32
	SetRegion(v string) *DescribeMOTokenUsageDetailRequest
	GetRegion() *string
	SetStartTime(v string) *DescribeMOTokenUsageDetailRequest
	GetStartTime() *string
	SetUsageType(v string) *DescribeMOTokenUsageDetailRequest
	GetUsageType() *string
}

type DescribeMOTokenUsageDetailRequest struct {
	// The API key used for the request.
	//
	// example:
	//
	// sk-rds-*****
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// The consumer associated with the API key.
	//
	// example:
	//
	// 16******4_rds_copilot****_public_cn-4****02
	ConsumerName *string `json:"ConsumerName,omitempty" xml:"ConsumerName,omitempty"`
	// The cursor-based pagination token. This parameter takes priority over Page. Leave this parameter empty for the first call. For subsequent calls, use the NextCursor value returned in the previous response.
	//
	// example:
	//
	// eyJpZCI6MTIzNDUsInRzIjoiMjAyNi0wNC0wOFQwMDowMDowMFoifQ==
	Cursor *string `json:"Cursor,omitempty" xml:"Cursor,omitempty"`
	// The end time in ISO 8601 format (UTC).
	//
	// example:
	//
	// 2025-12-13T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rds_copilot***_public_cn-*********6
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The model that was called.
	//
	// example:
	//
	// qwen-flash
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The page number. Minimum value: 1. Default value: 1.
	//
	// example:
	//
	// 2
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of records per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region in which the instance resides.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The start time in ISO 8601 format (UTC).
	//
	// example:
	//
	// 2026-01-04T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The type of usage to query.
	//
	// example:
	//
	// text
	UsageType *string `json:"UsageType,omitempty" xml:"UsageType,omitempty"`
}

func (s DescribeMOTokenUsageDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMOTokenUsageDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeMOTokenUsageDetailRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *DescribeMOTokenUsageDetailRequest) GetConsumerName() *string {
	return s.ConsumerName
}

func (s *DescribeMOTokenUsageDetailRequest) GetCursor() *string {
	return s.Cursor
}

func (s *DescribeMOTokenUsageDetailRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeMOTokenUsageDetailRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeMOTokenUsageDetailRequest) GetModel() *string {
	return s.Model
}

func (s *DescribeMOTokenUsageDetailRequest) GetPage() *int32 {
	return s.Page
}

func (s *DescribeMOTokenUsageDetailRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeMOTokenUsageDetailRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeMOTokenUsageDetailRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeMOTokenUsageDetailRequest) GetUsageType() *string {
	return s.UsageType
}

func (s *DescribeMOTokenUsageDetailRequest) SetApiKey(v string) *DescribeMOTokenUsageDetailRequest {
	s.ApiKey = &v
	return s
}

func (s *DescribeMOTokenUsageDetailRequest) SetConsumerName(v string) *DescribeMOTokenUsageDetailRequest {
	s.ConsumerName = &v
	return s
}

func (s *DescribeMOTokenUsageDetailRequest) SetCursor(v string) *DescribeMOTokenUsageDetailRequest {
	s.Cursor = &v
	return s
}

func (s *DescribeMOTokenUsageDetailRequest) SetEndTime(v string) *DescribeMOTokenUsageDetailRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeMOTokenUsageDetailRequest) SetInstanceId(v string) *DescribeMOTokenUsageDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeMOTokenUsageDetailRequest) SetModel(v string) *DescribeMOTokenUsageDetailRequest {
	s.Model = &v
	return s
}

func (s *DescribeMOTokenUsageDetailRequest) SetPage(v int32) *DescribeMOTokenUsageDetailRequest {
	s.Page = &v
	return s
}

func (s *DescribeMOTokenUsageDetailRequest) SetPageSize(v int32) *DescribeMOTokenUsageDetailRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMOTokenUsageDetailRequest) SetRegion(v string) *DescribeMOTokenUsageDetailRequest {
	s.Region = &v
	return s
}

func (s *DescribeMOTokenUsageDetailRequest) SetStartTime(v string) *DescribeMOTokenUsageDetailRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeMOTokenUsageDetailRequest) SetUsageType(v string) *DescribeMOTokenUsageDetailRequest {
	s.UsageType = &v
	return s
}

func (s *DescribeMOTokenUsageDetailRequest) Validate() error {
	return dara.Validate(s)
}
