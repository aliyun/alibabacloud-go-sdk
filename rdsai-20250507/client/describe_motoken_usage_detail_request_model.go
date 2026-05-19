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
}

type DescribeMOTokenUsageDetailRequest struct {
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// example:
	//
	// 16******4_rds_copilot****_public_cn-4****02
	ConsumerName *string `json:"ConsumerName,omitempty" xml:"ConsumerName,omitempty"`
	// example:
	//
	// eyJpZCI6MTIzNDUsInRzIjoiMjAyNi0wNC0wOFQwMDowMDowMFoifQ==
	Cursor *string `json:"Cursor,omitempty" xml:"Cursor,omitempty"`
	// example:
	//
	// 2025-12-13T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rds_copilot***_public_cn-*********6
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// qwen-flash
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// 2
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// 2026-01-04T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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

func (s *DescribeMOTokenUsageDetailRequest) Validate() error {
	return dara.Validate(s)
}
