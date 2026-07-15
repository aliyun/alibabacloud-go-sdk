// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConsumerGroupConsumersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListConsumerGroupConsumersResponseBody
	GetCode() *string
	SetData(v *ListConsumerGroupConsumersResponseBodyData) *ListConsumerGroupConsumersResponseBody
	GetData() *ListConsumerGroupConsumersResponseBodyData
	SetMessage(v string) *ListConsumerGroupConsumersResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListConsumerGroupConsumersResponseBody
	GetRequestId() *string
}

type ListConsumerGroupConsumersResponseBody struct {
	// example:
	//
	// Ok
	Code *string                                     `json:"code,omitempty" xml:"code,omitempty"`
	Data *ListConsumerGroupConsumersResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListConsumerGroupConsumersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListConsumerGroupConsumersResponseBody) GoString() string {
	return s.String()
}

func (s *ListConsumerGroupConsumersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListConsumerGroupConsumersResponseBody) GetData() *ListConsumerGroupConsumersResponseBodyData {
	return s.Data
}

func (s *ListConsumerGroupConsumersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListConsumerGroupConsumersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListConsumerGroupConsumersResponseBody) SetCode(v string) *ListConsumerGroupConsumersResponseBody {
	s.Code = &v
	return s
}

func (s *ListConsumerGroupConsumersResponseBody) SetData(v *ListConsumerGroupConsumersResponseBodyData) *ListConsumerGroupConsumersResponseBody {
	s.Data = v
	return s
}

func (s *ListConsumerGroupConsumersResponseBody) SetMessage(v string) *ListConsumerGroupConsumersResponseBody {
	s.Message = &v
	return s
}

func (s *ListConsumerGroupConsumersResponseBody) SetRequestId(v string) *ListConsumerGroupConsumersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConsumerGroupConsumersResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListConsumerGroupConsumersResponseBodyData struct {
	Items []*ListConsumerGroupConsumersResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
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

func (s ListConsumerGroupConsumersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListConsumerGroupConsumersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListConsumerGroupConsumersResponseBodyData) GetItems() []*ListConsumerGroupConsumersResponseBodyDataItems {
	return s.Items
}

func (s *ListConsumerGroupConsumersResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListConsumerGroupConsumersResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListConsumerGroupConsumersResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListConsumerGroupConsumersResponseBodyData) SetItems(v []*ListConsumerGroupConsumersResponseBodyDataItems) *ListConsumerGroupConsumersResponseBodyData {
	s.Items = v
	return s
}

func (s *ListConsumerGroupConsumersResponseBodyData) SetPageNumber(v int32) *ListConsumerGroupConsumersResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListConsumerGroupConsumersResponseBodyData) SetPageSize(v int32) *ListConsumerGroupConsumersResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListConsumerGroupConsumersResponseBodyData) SetTotalSize(v int32) *ListConsumerGroupConsumersResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListConsumerGroupConsumersResponseBodyData) Validate() error {
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

type ListConsumerGroupConsumersResponseBodyDataItems struct {
	// example:
	//
	// cs-8c13d2b4f8a1
	ConsumerId *string `json:"consumerId,omitempty" xml:"consumerId,omitempty"`
	// example:
	//
	// Success
	DeployStatus *string `json:"deployStatus,omitempty" xml:"deployStatus,omitempty"`
	// example:
	//
	// 线上 API 调用方
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// 1715769600000
	JoinTimestamp *int64 `json:"joinTimestamp,omitempty" xml:"joinTimestamp,omitempty"`
	// example:
	//
	// api-consumer
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListConsumerGroupConsumersResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s ListConsumerGroupConsumersResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListConsumerGroupConsumersResponseBodyDataItems) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *ListConsumerGroupConsumersResponseBodyDataItems) GetDeployStatus() *string {
	return s.DeployStatus
}

func (s *ListConsumerGroupConsumersResponseBodyDataItems) GetDescription() *string {
	return s.Description
}

func (s *ListConsumerGroupConsumersResponseBodyDataItems) GetEnable() *bool {
	return s.Enable
}

func (s *ListConsumerGroupConsumersResponseBodyDataItems) GetJoinTimestamp() *int64 {
	return s.JoinTimestamp
}

func (s *ListConsumerGroupConsumersResponseBodyDataItems) GetName() *string {
	return s.Name
}

func (s *ListConsumerGroupConsumersResponseBodyDataItems) SetConsumerId(v string) *ListConsumerGroupConsumersResponseBodyDataItems {
	s.ConsumerId = &v
	return s
}

func (s *ListConsumerGroupConsumersResponseBodyDataItems) SetDeployStatus(v string) *ListConsumerGroupConsumersResponseBodyDataItems {
	s.DeployStatus = &v
	return s
}

func (s *ListConsumerGroupConsumersResponseBodyDataItems) SetDescription(v string) *ListConsumerGroupConsumersResponseBodyDataItems {
	s.Description = &v
	return s
}

func (s *ListConsumerGroupConsumersResponseBodyDataItems) SetEnable(v bool) *ListConsumerGroupConsumersResponseBodyDataItems {
	s.Enable = &v
	return s
}

func (s *ListConsumerGroupConsumersResponseBodyDataItems) SetJoinTimestamp(v int64) *ListConsumerGroupConsumersResponseBodyDataItems {
	s.JoinTimestamp = &v
	return s
}

func (s *ListConsumerGroupConsumersResponseBodyDataItems) SetName(v string) *ListConsumerGroupConsumersResponseBodyDataItems {
	s.Name = &v
	return s
}

func (s *ListConsumerGroupConsumersResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
