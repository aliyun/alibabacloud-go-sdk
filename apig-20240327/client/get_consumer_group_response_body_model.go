// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConsumerGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetConsumerGroupResponseBody
	GetCode() *string
	SetData(v *GetConsumerGroupResponseBodyData) *GetConsumerGroupResponseBody
	GetData() *GetConsumerGroupResponseBodyData
	SetMessage(v string) *GetConsumerGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetConsumerGroupResponseBody
	GetRequestId() *string
}

type GetConsumerGroupResponseBody struct {
	// example:
	//
	// Ok
	Code *string                           `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetConsumerGroupResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetConsumerGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetConsumerGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetConsumerGroupResponseBody) GetData() *GetConsumerGroupResponseBodyData {
	return s.Data
}

func (s *GetConsumerGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetConsumerGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetConsumerGroupResponseBody) SetCode(v string) *GetConsumerGroupResponseBody {
	s.Code = &v
	return s
}

func (s *GetConsumerGroupResponseBody) SetData(v *GetConsumerGroupResponseBodyData) *GetConsumerGroupResponseBody {
	s.Data = v
	return s
}

func (s *GetConsumerGroupResponseBody) SetMessage(v string) *GetConsumerGroupResponseBody {
	s.Message = &v
	return s
}

func (s *GetConsumerGroupResponseBody) SetRequestId(v string) *GetConsumerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConsumerGroupResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetConsumerGroupResponseBodyData struct {
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
	// example:
	//
	// 1715769600000
	UpdateTimestamp *int64 `json:"updateTimestamp,omitempty" xml:"updateTimestamp,omitempty"`
}

func (s GetConsumerGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetConsumerGroupResponseBodyData) GetConsumerCount() *int64 {
	return s.ConsumerCount
}

func (s *GetConsumerGroupResponseBodyData) GetConsumerGroupId() *string {
	return s.ConsumerGroupId
}

func (s *GetConsumerGroupResponseBodyData) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *GetConsumerGroupResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetConsumerGroupResponseBodyData) GetGatewayType() *string {
	return s.GatewayType
}

func (s *GetConsumerGroupResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetConsumerGroupResponseBodyData) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *GetConsumerGroupResponseBodyData) SetConsumerCount(v int64) *GetConsumerGroupResponseBodyData {
	s.ConsumerCount = &v
	return s
}

func (s *GetConsumerGroupResponseBodyData) SetConsumerGroupId(v string) *GetConsumerGroupResponseBodyData {
	s.ConsumerGroupId = &v
	return s
}

func (s *GetConsumerGroupResponseBodyData) SetCreateTimestamp(v int64) *GetConsumerGroupResponseBodyData {
	s.CreateTimestamp = &v
	return s
}

func (s *GetConsumerGroupResponseBodyData) SetDescription(v string) *GetConsumerGroupResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetConsumerGroupResponseBodyData) SetGatewayType(v string) *GetConsumerGroupResponseBodyData {
	s.GatewayType = &v
	return s
}

func (s *GetConsumerGroupResponseBodyData) SetName(v string) *GetConsumerGroupResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetConsumerGroupResponseBodyData) SetUpdateTimestamp(v int64) *GetConsumerGroupResponseBodyData {
	s.UpdateTimestamp = &v
	return s
}

func (s *GetConsumerGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}
