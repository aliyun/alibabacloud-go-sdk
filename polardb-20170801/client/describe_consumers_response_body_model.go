// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConsumersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeConsumersResponseBodyItems) *DescribeConsumersResponseBody
	GetItems() []*DescribeConsumersResponseBodyItems
	SetPageNumber(v int32) *DescribeConsumersResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeConsumersResponseBody
	GetPageRecordCount() *int32
	SetPageSize(v int32) *DescribeConsumersResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeConsumersResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeConsumersResponseBody
	GetTotalRecordCount() *int32
}

type DescribeConsumersResponseBody struct {
	Items []*DescribeConsumersResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 1
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 24A1990B-4F6E-482B-B8CB-75C612******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeConsumersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeConsumersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConsumersResponseBody) GetItems() []*DescribeConsumersResponseBodyItems {
	return s.Items
}

func (s *DescribeConsumersResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeConsumersResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeConsumersResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeConsumersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeConsumersResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeConsumersResponseBody) SetItems(v []*DescribeConsumersResponseBodyItems) *DescribeConsumersResponseBody {
	s.Items = v
	return s
}

func (s *DescribeConsumersResponseBody) SetPageNumber(v int32) *DescribeConsumersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeConsumersResponseBody) SetPageRecordCount(v int32) *DescribeConsumersResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeConsumersResponseBody) SetPageSize(v int32) *DescribeConsumersResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeConsumersResponseBody) SetRequestId(v string) *DescribeConsumersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeConsumersResponseBody) SetTotalRecordCount(v int32) *DescribeConsumersResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeConsumersResponseBody) Validate() error {
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

type DescribeConsumersResponseBodyItems struct {
	// example:
	//
	// "[]"
	AllowedModels *string `json:"AllowedModels,omitempty" xml:"AllowedModels,omitempty"`
	// example:
	//
	// xxxxxxxx
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// example:
	//
	// cg-xxxxxx
	ConsumerGroupId *string `json:"ConsumerGroupId,omitempty" xml:"ConsumerGroupId,omitempty"`
	// example:
	//
	// test
	ConsumerGroupName *string `json:"ConsumerGroupName,omitempty" xml:"ConsumerGroupName,omitempty"`
	// example:
	//
	// c-mqveroemc***
	ConsumerId *string `json:"ConsumerId,omitempty" xml:"ConsumerId,omitempty"`
	// example:
	//
	// 2026-01-28T09:56:03+08:00
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// example:
	//
	// 2026-01-04T16:09:29+08:00
	GmtModified        *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	LifetimeCostCount  *int64  `json:"LifetimeCostCount,omitempty" xml:"LifetimeCostCount,omitempty"`
	LifetimeTokenCount *int64  `json:"LifetimeTokenCount,omitempty" xml:"LifetimeTokenCount,omitempty"`
	MtdCostCount       *int64  `json:"MtdCostCount,omitempty" xml:"MtdCostCount,omitempty"`
	MtdTokenCount      *int64  `json:"MtdTokenCount,omitempty" xml:"MtdTokenCount,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// yonghu
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
}

func (s DescribeConsumersResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeConsumersResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeConsumersResponseBodyItems) GetAllowedModels() *string {
	return s.AllowedModels
}

func (s *DescribeConsumersResponseBodyItems) GetApiKey() *string {
	return s.ApiKey
}

func (s *DescribeConsumersResponseBodyItems) GetConsumerGroupId() *string {
	return s.ConsumerGroupId
}

func (s *DescribeConsumersResponseBodyItems) GetConsumerGroupName() *string {
	return s.ConsumerGroupName
}

func (s *DescribeConsumersResponseBodyItems) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *DescribeConsumersResponseBodyItems) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeConsumersResponseBodyItems) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeConsumersResponseBodyItems) GetLifetimeCostCount() *int64 {
	return s.LifetimeCostCount
}

func (s *DescribeConsumersResponseBodyItems) GetLifetimeTokenCount() *int64 {
	return s.LifetimeTokenCount
}

func (s *DescribeConsumersResponseBodyItems) GetMtdCostCount() *int64 {
	return s.MtdCostCount
}

func (s *DescribeConsumersResponseBodyItems) GetMtdTokenCount() *int64 {
	return s.MtdTokenCount
}

func (s *DescribeConsumersResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *DescribeConsumersResponseBodyItems) GetNickName() *string {
	return s.NickName
}

func (s *DescribeConsumersResponseBodyItems) SetAllowedModels(v string) *DescribeConsumersResponseBodyItems {
	s.AllowedModels = &v
	return s
}

func (s *DescribeConsumersResponseBodyItems) SetApiKey(v string) *DescribeConsumersResponseBodyItems {
	s.ApiKey = &v
	return s
}

func (s *DescribeConsumersResponseBodyItems) SetConsumerGroupId(v string) *DescribeConsumersResponseBodyItems {
	s.ConsumerGroupId = &v
	return s
}

func (s *DescribeConsumersResponseBodyItems) SetConsumerGroupName(v string) *DescribeConsumersResponseBodyItems {
	s.ConsumerGroupName = &v
	return s
}

func (s *DescribeConsumersResponseBodyItems) SetConsumerId(v string) *DescribeConsumersResponseBodyItems {
	s.ConsumerId = &v
	return s
}

func (s *DescribeConsumersResponseBodyItems) SetGmtCreated(v string) *DescribeConsumersResponseBodyItems {
	s.GmtCreated = &v
	return s
}

func (s *DescribeConsumersResponseBodyItems) SetGmtModified(v string) *DescribeConsumersResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *DescribeConsumersResponseBodyItems) SetLifetimeCostCount(v int64) *DescribeConsumersResponseBodyItems {
	s.LifetimeCostCount = &v
	return s
}

func (s *DescribeConsumersResponseBodyItems) SetLifetimeTokenCount(v int64) *DescribeConsumersResponseBodyItems {
	s.LifetimeTokenCount = &v
	return s
}

func (s *DescribeConsumersResponseBodyItems) SetMtdCostCount(v int64) *DescribeConsumersResponseBodyItems {
	s.MtdCostCount = &v
	return s
}

func (s *DescribeConsumersResponseBodyItems) SetMtdTokenCount(v int64) *DescribeConsumersResponseBodyItems {
	s.MtdTokenCount = &v
	return s
}

func (s *DescribeConsumersResponseBodyItems) SetName(v string) *DescribeConsumersResponseBodyItems {
	s.Name = &v
	return s
}

func (s *DescribeConsumersResponseBodyItems) SetNickName(v string) *DescribeConsumersResponseBodyItems {
	s.NickName = &v
	return s
}

func (s *DescribeConsumersResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
