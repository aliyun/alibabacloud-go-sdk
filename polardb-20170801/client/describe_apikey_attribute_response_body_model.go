// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApikeyAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeApikeyAttributeResponseBodyItems) *DescribeApikeyAttributeResponseBody
	GetItems() []*DescribeApikeyAttributeResponseBodyItems
	SetRequestId(v string) *DescribeApikeyAttributeResponseBody
	GetRequestId() *string
}

type DescribeApikeyAttributeResponseBody struct {
	// The list of consumer objects.
	Items []*DescribeApikeyAttributeResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// F3322AFE-083E-4D77-A074-421301******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeApikeyAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApikeyAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApikeyAttributeResponseBody) GetItems() []*DescribeApikeyAttributeResponseBodyItems {
	return s.Items
}

func (s *DescribeApikeyAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApikeyAttributeResponseBody) SetItems(v []*DescribeApikeyAttributeResponseBodyItems) *DescribeApikeyAttributeResponseBody {
	s.Items = v
	return s
}

func (s *DescribeApikeyAttributeResponseBody) SetRequestId(v string) *DescribeApikeyAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApikeyAttributeResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApikeyAttributeResponseBodyItems struct {
	// The consumer information.
	Consumer *DescribeApikeyAttributeResponseBodyItemsConsumer `json:"Consumer,omitempty" xml:"Consumer,omitempty" type:"Struct"`
	// The usage statistics for the consumer.
	UsageStatistics []*DescribeApikeyAttributeResponseBodyItemsUsageStatistics `json:"UsageStatistics,omitempty" xml:"UsageStatistics,omitempty" type:"Repeated"`
}

func (s DescribeApikeyAttributeResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeApikeyAttributeResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeApikeyAttributeResponseBodyItems) GetConsumer() *DescribeApikeyAttributeResponseBodyItemsConsumer {
	return s.Consumer
}

func (s *DescribeApikeyAttributeResponseBodyItems) GetUsageStatistics() []*DescribeApikeyAttributeResponseBodyItemsUsageStatistics {
	return s.UsageStatistics
}

func (s *DescribeApikeyAttributeResponseBodyItems) SetConsumer(v *DescribeApikeyAttributeResponseBodyItemsConsumer) *DescribeApikeyAttributeResponseBodyItems {
	s.Consumer = v
	return s
}

func (s *DescribeApikeyAttributeResponseBodyItems) SetUsageStatistics(v []*DescribeApikeyAttributeResponseBodyItemsUsageStatistics) *DescribeApikeyAttributeResponseBodyItems {
	s.UsageStatistics = v
	return s
}

func (s *DescribeApikeyAttributeResponseBodyItems) Validate() error {
	if s.Consumer != nil {
		if err := s.Consumer.Validate(); err != nil {
			return err
		}
	}
	if s.UsageStatistics != nil {
		for _, item := range s.UsageStatistics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApikeyAttributeResponseBodyItemsConsumer struct {
	// The API key.
	//
	// example:
	//
	// ***
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// The user group ID.
	//
	// example:
	//
	// cg-bq6rcdjp02vt
	ConsumerGroupId *string `json:"ConsumerGroupId,omitempty" xml:"ConsumerGroupId,omitempty"`
	// The consumer ID.
	//
	// example:
	//
	// c-71qh3pscbd3i
	ConsumerId *string `json:"ConsumerId,omitempty" xml:"ConsumerId,omitempty"`
	// The consumer tag.
	//
	// example:
	//
	// test
	ConsumerTag *string `json:"ConsumerTag,omitempty" xml:"ConsumerTag,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2021-07-18T07:32:30Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The gateway instance ID.
	//
	// example:
	//
	// pg-2ze5n62ef4s165***
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	// The last modification time.
	//
	// example:
	//
	// 2026-04-10T01:48:25Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The consumer name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The consumer status.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeApikeyAttributeResponseBodyItemsConsumer) String() string {
	return dara.Prettify(s)
}

func (s DescribeApikeyAttributeResponseBodyItemsConsumer) GoString() string {
	return s.String()
}

func (s *DescribeApikeyAttributeResponseBodyItemsConsumer) GetApiKey() *string {
	return s.ApiKey
}

func (s *DescribeApikeyAttributeResponseBodyItemsConsumer) GetConsumerGroupId() *string {
	return s.ConsumerGroupId
}

func (s *DescribeApikeyAttributeResponseBodyItemsConsumer) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *DescribeApikeyAttributeResponseBodyItemsConsumer) GetConsumerTag() *string {
	return s.ConsumerTag
}

func (s *DescribeApikeyAttributeResponseBodyItemsConsumer) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeApikeyAttributeResponseBodyItemsConsumer) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *DescribeApikeyAttributeResponseBodyItemsConsumer) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *DescribeApikeyAttributeResponseBodyItemsConsumer) GetName() *string {
	return s.Name
}

func (s *DescribeApikeyAttributeResponseBodyItemsConsumer) GetStatus() *string {
	return s.Status
}

func (s *DescribeApikeyAttributeResponseBodyItemsConsumer) SetApiKey(v string) *DescribeApikeyAttributeResponseBodyItemsConsumer {
	s.ApiKey = &v
	return s
}

func (s *DescribeApikeyAttributeResponseBodyItemsConsumer) SetConsumerGroupId(v string) *DescribeApikeyAttributeResponseBodyItemsConsumer {
	s.ConsumerGroupId = &v
	return s
}

func (s *DescribeApikeyAttributeResponseBodyItemsConsumer) SetConsumerId(v string) *DescribeApikeyAttributeResponseBodyItemsConsumer {
	s.ConsumerId = &v
	return s
}

func (s *DescribeApikeyAttributeResponseBodyItemsConsumer) SetConsumerTag(v string) *DescribeApikeyAttributeResponseBodyItemsConsumer {
	s.ConsumerTag = &v
	return s
}

func (s *DescribeApikeyAttributeResponseBodyItemsConsumer) SetCreateTime(v string) *DescribeApikeyAttributeResponseBodyItemsConsumer {
	s.CreateTime = &v
	return s
}

func (s *DescribeApikeyAttributeResponseBodyItemsConsumer) SetGwClusterId(v string) *DescribeApikeyAttributeResponseBodyItemsConsumer {
	s.GwClusterId = &v
	return s
}

func (s *DescribeApikeyAttributeResponseBodyItemsConsumer) SetModifyTime(v string) *DescribeApikeyAttributeResponseBodyItemsConsumer {
	s.ModifyTime = &v
	return s
}

func (s *DescribeApikeyAttributeResponseBodyItemsConsumer) SetName(v string) *DescribeApikeyAttributeResponseBodyItemsConsumer {
	s.Name = &v
	return s
}

func (s *DescribeApikeyAttributeResponseBodyItemsConsumer) SetStatus(v string) *DescribeApikeyAttributeResponseBodyItemsConsumer {
	s.Status = &v
	return s
}

func (s *DescribeApikeyAttributeResponseBodyItemsConsumer) Validate() error {
	return dara.Validate(s)
}

type DescribeApikeyAttributeResponseBodyItemsUsageStatistics struct {
	// The dimension reference ID, which is the ConsumerId.
	//
	// example:
	//
	// c-xxxxxx
	DimensionRefId *string `json:"DimensionRefId,omitempty" xml:"DimensionRefId,omitempty"`
	// The statistics dimension. The value is typically Consumer.
	//
	// example:
	//
	// Consumer
	DimensionType *string `json:"DimensionType,omitempty" xml:"DimensionType,omitempty"`
	// The gateway instance ID.
	//
	// example:
	//
	// pg-bp1ln7w98yrhzz7i2
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	// The number of cache tokens for the current month.
	//
	// example:
	//
	// 10
	MonthlyCacheToken *string `json:"MonthlyCacheToken,omitempty" xml:"MonthlyCacheToken,omitempty"`
	// The cost points consumed for the current month.
	//
	// example:
	//
	// 10
	MonthlyCostPoints *string `json:"MonthlyCostPoints,omitempty" xml:"MonthlyCostPoints,omitempty"`
	// The number of input tokens for the current month.
	//
	// example:
	//
	// 10
	MonthlyInputToken *string `json:"MonthlyInputToken,omitempty" xml:"MonthlyInputToken,omitempty"`
	// The number of output tokens for the current month.
	//
	// example:
	//
	// 10
	MonthlyOutputToken *string `json:"MonthlyOutputToken,omitempty" xml:"MonthlyOutputToken,omitempty"`
	// The number of tokens for the current month.
	//
	// example:
	//
	// 10
	MonthlyToken *string `json:"MonthlyToken,omitempty" xml:"MonthlyToken,omitempty"`
	// The cumulative number of cache tokens.
	//
	// example:
	//
	// 10
	TotalCacheToken *string `json:"TotalCacheToken,omitempty" xml:"TotalCacheToken,omitempty"`
	// The cumulative cost points consumed.
	//
	// example:
	//
	// 10
	TotalCostPoints *string `json:"TotalCostPoints,omitempty" xml:"TotalCostPoints,omitempty"`
	// The cumulative number of input tokens.
	//
	// example:
	//
	// 10
	TotalInputToken *string `json:"TotalInputToken,omitempty" xml:"TotalInputToken,omitempty"`
	// The cumulative number of output tokens.
	//
	// example:
	//
	// 10
	TotalOutputToken *string `json:"TotalOutputToken,omitempty" xml:"TotalOutputToken,omitempty"`
	// The cumulative number of tokens.
	//
	// example:
	//
	// 10
	TotalToken *string `json:"TotalToken,omitempty" xml:"TotalToken,omitempty"`
}

func (s DescribeApikeyAttributeResponseBodyItemsUsageStatistics) String() string {
	return dara.Prettify(s)
}

func (s DescribeApikeyAttributeResponseBodyItemsUsageStatistics) GoString() string {
	return s.String()
}

func (s *DescribeApikeyAttributeResponseBodyItemsUsageStatistics) GetDimensionRefId() *string {
	return s.DimensionRefId
}

func (s *DescribeApikeyAttributeResponseBodyItemsUsageStatistics) GetDimensionType() *string {
	return s.DimensionType
}

func (s *DescribeApikeyAttributeResponseBodyItemsUsageStatistics) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *DescribeApikeyAttributeResponseBodyItemsUsageStatistics) GetMonthlyCacheToken() *string {
	return s.MonthlyCacheToken
}

func (s *DescribeApikeyAttributeResponseBodyItemsUsageStatistics) GetMonthlyCostPoints() *string {
	return s.MonthlyCostPoints
}

func (s *DescribeApikeyAttributeResponseBodyItemsUsageStatistics) GetMonthlyInputToken() *string {
	return s.MonthlyInputToken
}

func (s *DescribeApikeyAttributeResponseBodyItemsUsageStatistics) GetMonthlyOutputToken() *string {
	return s.MonthlyOutputToken
}

func (s *DescribeApikeyAttributeResponseBodyItemsUsageStatistics) GetMonthlyToken() *string {
	return s.MonthlyToken
}

func (s *DescribeApikeyAttributeResponseBodyItemsUsageStatistics) GetTotalCacheToken() *string {
	return s.TotalCacheToken
}

func (s *DescribeApikeyAttributeResponseBodyItemsUsageStatistics) GetTotalCostPoints() *string {
	return s.TotalCostPoints
}

func (s *DescribeApikeyAttributeResponseBodyItemsUsageStatistics) GetTotalInputToken() *string {
	return s.TotalInputToken
}

func (s *DescribeApikeyAttributeResponseBodyItemsUsageStatistics) GetTotalOutputToken() *string {
	return s.TotalOutputToken
}

func (s *DescribeApikeyAttributeResponseBodyItemsUsageStatistics) GetTotalToken() *string {
	return s.TotalToken
}

func (s *DescribeApikeyAttributeResponseBodyItemsUsageStatistics) SetDimensionRefId(v string) *DescribeApikeyAttributeResponseBodyItemsUsageStatistics {
	s.DimensionRefId = &v
	return s
}

func (s *DescribeApikeyAttributeResponseBodyItemsUsageStatistics) SetDimensionType(v string) *DescribeApikeyAttributeResponseBodyItemsUsageStatistics {
	s.DimensionType = &v
	return s
}

func (s *DescribeApikeyAttributeResponseBodyItemsUsageStatistics) SetGwClusterId(v string) *DescribeApikeyAttributeResponseBodyItemsUsageStatistics {
	s.GwClusterId = &v
	return s
}

func (s *DescribeApikeyAttributeResponseBodyItemsUsageStatistics) SetMonthlyCacheToken(v string) *DescribeApikeyAttributeResponseBodyItemsUsageStatistics {
	s.MonthlyCacheToken = &v
	return s
}

func (s *DescribeApikeyAttributeResponseBodyItemsUsageStatistics) SetMonthlyCostPoints(v string) *DescribeApikeyAttributeResponseBodyItemsUsageStatistics {
	s.MonthlyCostPoints = &v
	return s
}

func (s *DescribeApikeyAttributeResponseBodyItemsUsageStatistics) SetMonthlyInputToken(v string) *DescribeApikeyAttributeResponseBodyItemsUsageStatistics {
	s.MonthlyInputToken = &v
	return s
}

func (s *DescribeApikeyAttributeResponseBodyItemsUsageStatistics) SetMonthlyOutputToken(v string) *DescribeApikeyAttributeResponseBodyItemsUsageStatistics {
	s.MonthlyOutputToken = &v
	return s
}

func (s *DescribeApikeyAttributeResponseBodyItemsUsageStatistics) SetMonthlyToken(v string) *DescribeApikeyAttributeResponseBodyItemsUsageStatistics {
	s.MonthlyToken = &v
	return s
}

func (s *DescribeApikeyAttributeResponseBodyItemsUsageStatistics) SetTotalCacheToken(v string) *DescribeApikeyAttributeResponseBodyItemsUsageStatistics {
	s.TotalCacheToken = &v
	return s
}

func (s *DescribeApikeyAttributeResponseBodyItemsUsageStatistics) SetTotalCostPoints(v string) *DescribeApikeyAttributeResponseBodyItemsUsageStatistics {
	s.TotalCostPoints = &v
	return s
}

func (s *DescribeApikeyAttributeResponseBodyItemsUsageStatistics) SetTotalInputToken(v string) *DescribeApikeyAttributeResponseBodyItemsUsageStatistics {
	s.TotalInputToken = &v
	return s
}

func (s *DescribeApikeyAttributeResponseBodyItemsUsageStatistics) SetTotalOutputToken(v string) *DescribeApikeyAttributeResponseBodyItemsUsageStatistics {
	s.TotalOutputToken = &v
	return s
}

func (s *DescribeApikeyAttributeResponseBodyItemsUsageStatistics) SetTotalToken(v string) *DescribeApikeyAttributeResponseBodyItemsUsageStatistics {
	s.TotalToken = &v
	return s
}

func (s *DescribeApikeyAttributeResponseBodyItemsUsageStatistics) Validate() error {
	return dara.Validate(s)
}
