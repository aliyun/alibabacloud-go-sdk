// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConsumerGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateConsumerGroupResponseBody
	GetCode() *string
	SetData(v *CreateConsumerGroupResponseBodyData) *CreateConsumerGroupResponseBody
	GetData() *CreateConsumerGroupResponseBodyData
	SetMessage(v string) *CreateConsumerGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateConsumerGroupResponseBody
	GetRequestId() *string
}

type CreateConsumerGroupResponseBody struct {
	// example:
	//
	// Ok
	Code *string                              `json:"code,omitempty" xml:"code,omitempty"`
	Data *CreateConsumerGroupResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateConsumerGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateConsumerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConsumerGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateConsumerGroupResponseBody) GetData() *CreateConsumerGroupResponseBodyData {
	return s.Data
}

func (s *CreateConsumerGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateConsumerGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateConsumerGroupResponseBody) SetCode(v string) *CreateConsumerGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CreateConsumerGroupResponseBody) SetData(v *CreateConsumerGroupResponseBodyData) *CreateConsumerGroupResponseBody {
	s.Data = v
	return s
}

func (s *CreateConsumerGroupResponseBody) SetMessage(v string) *CreateConsumerGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateConsumerGroupResponseBody) SetRequestId(v string) *CreateConsumerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateConsumerGroupResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateConsumerGroupResponseBodyData struct {
	// example:
	//
	// csg-8c13d2b4f8a1
	ConsumerGroupId *string `json:"consumerGroupId,omitempty" xml:"consumerGroupId,omitempty"`
	// example:
	//
	// API
	GatewayType *string `json:"gatewayType,omitempty" xml:"gatewayType,omitempty"`
}

func (s CreateConsumerGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateConsumerGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateConsumerGroupResponseBodyData) GetConsumerGroupId() *string {
	return s.ConsumerGroupId
}

func (s *CreateConsumerGroupResponseBodyData) GetGatewayType() *string {
	return s.GatewayType
}

func (s *CreateConsumerGroupResponseBodyData) SetConsumerGroupId(v string) *CreateConsumerGroupResponseBodyData {
	s.ConsumerGroupId = &v
	return s
}

func (s *CreateConsumerGroupResponseBodyData) SetGatewayType(v string) *CreateConsumerGroupResponseBodyData {
	s.GatewayType = &v
	return s
}

func (s *CreateConsumerGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}
