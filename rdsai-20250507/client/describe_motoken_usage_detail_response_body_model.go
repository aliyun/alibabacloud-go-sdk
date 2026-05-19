// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMOTokenUsageDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextCursor(v string) *DescribeMOTokenUsageDetailResponseBody
	GetNextCursor() *string
	SetPage(v int32) *DescribeMOTokenUsageDetailResponseBody
	GetPage() *int32
	SetPageSize(v int32) *DescribeMOTokenUsageDetailResponseBody
	GetPageSize() *int32
	SetRecords(v []*DescribeMOTokenUsageDetailResponseBodyRecords) *DescribeMOTokenUsageDetailResponseBody
	GetRecords() []*DescribeMOTokenUsageDetailResponseBodyRecords
	SetRequestId(v string) *DescribeMOTokenUsageDetailResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeMOTokenUsageDetailResponseBody
	GetTotalCount() *int32
}

type DescribeMOTokenUsageDetailResponseBody struct {
	// example:
	//
	// eyJpZCI6MTIzNDUsInRzIjoiMjAyNi0wNC0wOFQwMDowMDowMFoifQ==
	NextCursor *string `json:"NextCursor,omitempty" xml:"NextCursor,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// example:
	//
	// 10
	PageSize *int32                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Records  []*DescribeMOTokenUsageDetailResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329241C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeMOTokenUsageDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMOTokenUsageDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMOTokenUsageDetailResponseBody) GetNextCursor() *string {
	return s.NextCursor
}

func (s *DescribeMOTokenUsageDetailResponseBody) GetPage() *int32 {
	return s.Page
}

func (s *DescribeMOTokenUsageDetailResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeMOTokenUsageDetailResponseBody) GetRecords() []*DescribeMOTokenUsageDetailResponseBodyRecords {
	return s.Records
}

func (s *DescribeMOTokenUsageDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMOTokenUsageDetailResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeMOTokenUsageDetailResponseBody) SetNextCursor(v string) *DescribeMOTokenUsageDetailResponseBody {
	s.NextCursor = &v
	return s
}

func (s *DescribeMOTokenUsageDetailResponseBody) SetPage(v int32) *DescribeMOTokenUsageDetailResponseBody {
	s.Page = &v
	return s
}

func (s *DescribeMOTokenUsageDetailResponseBody) SetPageSize(v int32) *DescribeMOTokenUsageDetailResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeMOTokenUsageDetailResponseBody) SetRecords(v []*DescribeMOTokenUsageDetailResponseBodyRecords) *DescribeMOTokenUsageDetailResponseBody {
	s.Records = v
	return s
}

func (s *DescribeMOTokenUsageDetailResponseBody) SetRequestId(v string) *DescribeMOTokenUsageDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMOTokenUsageDetailResponseBody) SetTotalCount(v int32) *DescribeMOTokenUsageDetailResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeMOTokenUsageDetailResponseBody) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMOTokenUsageDetailResponseBodyRecords struct {
	// example:
	//
	// 16******4_rds_copilot****_public_cn-4****02
	ConsumerName *string `json:"ConsumerName,omitempty" xml:"ConsumerName,omitempty"`
	// example:
	//
	// 10000
	InputTokens *float64 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
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
	// 10000
	OutputTokens *float64 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// 2026-04-01T01:00:00Z
	RequestTime *string `json:"RequestTime,omitempty" xml:"RequestTime,omitempty"`
	// example:
	//
	// 100000
	TotalTokens *float64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s DescribeMOTokenUsageDetailResponseBodyRecords) String() string {
	return dara.Prettify(s)
}

func (s DescribeMOTokenUsageDetailResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *DescribeMOTokenUsageDetailResponseBodyRecords) GetConsumerName() *string {
	return s.ConsumerName
}

func (s *DescribeMOTokenUsageDetailResponseBodyRecords) GetInputTokens() *float64 {
	return s.InputTokens
}

func (s *DescribeMOTokenUsageDetailResponseBodyRecords) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeMOTokenUsageDetailResponseBodyRecords) GetModel() *string {
	return s.Model
}

func (s *DescribeMOTokenUsageDetailResponseBodyRecords) GetOutputTokens() *float64 {
	return s.OutputTokens
}

func (s *DescribeMOTokenUsageDetailResponseBodyRecords) GetRegion() *string {
	return s.Region
}

func (s *DescribeMOTokenUsageDetailResponseBodyRecords) GetRequestTime() *string {
	return s.RequestTime
}

func (s *DescribeMOTokenUsageDetailResponseBodyRecords) GetTotalTokens() *float64 {
	return s.TotalTokens
}

func (s *DescribeMOTokenUsageDetailResponseBodyRecords) SetConsumerName(v string) *DescribeMOTokenUsageDetailResponseBodyRecords {
	s.ConsumerName = &v
	return s
}

func (s *DescribeMOTokenUsageDetailResponseBodyRecords) SetInputTokens(v float64) *DescribeMOTokenUsageDetailResponseBodyRecords {
	s.InputTokens = &v
	return s
}

func (s *DescribeMOTokenUsageDetailResponseBodyRecords) SetInstanceId(v string) *DescribeMOTokenUsageDetailResponseBodyRecords {
	s.InstanceId = &v
	return s
}

func (s *DescribeMOTokenUsageDetailResponseBodyRecords) SetModel(v string) *DescribeMOTokenUsageDetailResponseBodyRecords {
	s.Model = &v
	return s
}

func (s *DescribeMOTokenUsageDetailResponseBodyRecords) SetOutputTokens(v float64) *DescribeMOTokenUsageDetailResponseBodyRecords {
	s.OutputTokens = &v
	return s
}

func (s *DescribeMOTokenUsageDetailResponseBodyRecords) SetRegion(v string) *DescribeMOTokenUsageDetailResponseBodyRecords {
	s.Region = &v
	return s
}

func (s *DescribeMOTokenUsageDetailResponseBodyRecords) SetRequestTime(v string) *DescribeMOTokenUsageDetailResponseBodyRecords {
	s.RequestTime = &v
	return s
}

func (s *DescribeMOTokenUsageDetailResponseBodyRecords) SetTotalTokens(v float64) *DescribeMOTokenUsageDetailResponseBodyRecords {
	s.TotalTokens = &v
	return s
}

func (s *DescribeMOTokenUsageDetailResponseBodyRecords) Validate() error {
	return dara.Validate(s)
}
