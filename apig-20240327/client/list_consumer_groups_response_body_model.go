// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConsumerGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListConsumerGroupsResponseBody
	GetCode() *string
	SetData(v *ListConsumerGroupsResponseBodyData) *ListConsumerGroupsResponseBody
	GetData() *ListConsumerGroupsResponseBodyData
	SetMessage(v string) *ListConsumerGroupsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListConsumerGroupsResponseBody
	GetRequestId() *string
}

type ListConsumerGroupsResponseBody struct {
	// example:
	//
	// Ok
	Code *string                             `json:"code,omitempty" xml:"code,omitempty"`
	Data *ListConsumerGroupsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListConsumerGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListConsumerGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConsumerGroupsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListConsumerGroupsResponseBody) GetData() *ListConsumerGroupsResponseBodyData {
	return s.Data
}

func (s *ListConsumerGroupsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListConsumerGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListConsumerGroupsResponseBody) SetCode(v string) *ListConsumerGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *ListConsumerGroupsResponseBody) SetData(v *ListConsumerGroupsResponseBodyData) *ListConsumerGroupsResponseBody {
	s.Data = v
	return s
}

func (s *ListConsumerGroupsResponseBody) SetMessage(v string) *ListConsumerGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *ListConsumerGroupsResponseBody) SetRequestId(v string) *ListConsumerGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConsumerGroupsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListConsumerGroupsResponseBodyData struct {
	Items []*ListConsumerGroupsResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 2
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListConsumerGroupsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListConsumerGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListConsumerGroupsResponseBodyData) GetItems() []*ListConsumerGroupsResponseBodyDataItems {
	return s.Items
}

func (s *ListConsumerGroupsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListConsumerGroupsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListConsumerGroupsResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListConsumerGroupsResponseBodyData) SetItems(v []*ListConsumerGroupsResponseBodyDataItems) *ListConsumerGroupsResponseBodyData {
	s.Items = v
	return s
}

func (s *ListConsumerGroupsResponseBodyData) SetPageNumber(v int32) *ListConsumerGroupsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListConsumerGroupsResponseBodyData) SetPageSize(v int32) *ListConsumerGroupsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListConsumerGroupsResponseBodyData) SetTotalSize(v int32) *ListConsumerGroupsResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListConsumerGroupsResponseBodyData) Validate() error {
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

type ListConsumerGroupsResponseBodyDataItems struct {
	// example:
	//
	// 3
	ConsumerCount *int64 `json:"consumerCount,omitempty" xml:"consumerCount,omitempty"`
	// example:
	//
	// csg-8c13d2b4f8a1
	ConsumerGroupId *string `json:"consumerGroupId,omitempty" xml:"consumerGroupId,omitempty"`
	// example:
	//
	// 1715769600000
	CreateTimestamp *int64 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	// example:
	//
	// 用于线上 API 调用方分组
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// API
	GatewayType *string `json:"gatewayType,omitempty" xml:"gatewayType,omitempty"`
	// example:
	//
	// api-consumer-group
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListConsumerGroupsResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s ListConsumerGroupsResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListConsumerGroupsResponseBodyDataItems) GetConsumerCount() *int64 {
	return s.ConsumerCount
}

func (s *ListConsumerGroupsResponseBodyDataItems) GetConsumerGroupId() *string {
	return s.ConsumerGroupId
}

func (s *ListConsumerGroupsResponseBodyDataItems) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *ListConsumerGroupsResponseBodyDataItems) GetDescription() *string {
	return s.Description
}

func (s *ListConsumerGroupsResponseBodyDataItems) GetGatewayType() *string {
	return s.GatewayType
}

func (s *ListConsumerGroupsResponseBodyDataItems) GetName() *string {
	return s.Name
}

func (s *ListConsumerGroupsResponseBodyDataItems) SetConsumerCount(v int64) *ListConsumerGroupsResponseBodyDataItems {
	s.ConsumerCount = &v
	return s
}

func (s *ListConsumerGroupsResponseBodyDataItems) SetConsumerGroupId(v string) *ListConsumerGroupsResponseBodyDataItems {
	s.ConsumerGroupId = &v
	return s
}

func (s *ListConsumerGroupsResponseBodyDataItems) SetCreateTimestamp(v int64) *ListConsumerGroupsResponseBodyDataItems {
	s.CreateTimestamp = &v
	return s
}

func (s *ListConsumerGroupsResponseBodyDataItems) SetDescription(v string) *ListConsumerGroupsResponseBodyDataItems {
	s.Description = &v
	return s
}

func (s *ListConsumerGroupsResponseBodyDataItems) SetGatewayType(v string) *ListConsumerGroupsResponseBodyDataItems {
	s.GatewayType = &v
	return s
}

func (s *ListConsumerGroupsResponseBodyDataItems) SetName(v string) *ListConsumerGroupsResponseBodyDataItems {
	s.Name = &v
	return s
}

func (s *ListConsumerGroupsResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
