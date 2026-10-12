// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHiMarketConsumerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetHiMarketConsumerResponseBody
	GetCode() *string
	SetData(v *GetHiMarketConsumerResponseBodyData) *GetHiMarketConsumerResponseBody
	GetData() *GetHiMarketConsumerResponseBodyData
	SetMessage(v string) *GetHiMarketConsumerResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHiMarketConsumerResponseBody
	GetRequestId() *string
}

type GetHiMarketConsumerResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *GetHiMarketConsumerResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetHiMarketConsumerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHiMarketConsumerResponseBody) GoString() string {
	return s.String()
}

func (s *GetHiMarketConsumerResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetHiMarketConsumerResponseBody) GetData() *GetHiMarketConsumerResponseBodyData {
	return s.Data
}

func (s *GetHiMarketConsumerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHiMarketConsumerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHiMarketConsumerResponseBody) SetCode(v string) *GetHiMarketConsumerResponseBody {
	s.Code = &v
	return s
}

func (s *GetHiMarketConsumerResponseBody) SetData(v *GetHiMarketConsumerResponseBodyData) *GetHiMarketConsumerResponseBody {
	s.Data = v
	return s
}

func (s *GetHiMarketConsumerResponseBody) SetMessage(v string) *GetHiMarketConsumerResponseBody {
	s.Message = &v
	return s
}

func (s *GetHiMarketConsumerResponseBody) SetRequestId(v string) *GetHiMarketConsumerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHiMarketConsumerResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetHiMarketConsumerResponseBodyData struct {
	// The consumer ID.
	//
	// example:
	//
	// cs-xxxxxxxxxxxx
	ConsumerId *string `json:"consumerId,omitempty" xml:"consumerId,omitempty"`
	// The description of the consumer.
	//
	// example:
	//
	// Dedicated test consumer
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Indicates whether the consumer is the primary consumer.
	IsPrimary *bool `json:"isPrimary,omitempty" xml:"isPrimary,omitempty"`
	// The name of the consumer.
	//
	// example:
	//
	// consumer-1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetHiMarketConsumerResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetHiMarketConsumerResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHiMarketConsumerResponseBodyData) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *GetHiMarketConsumerResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetHiMarketConsumerResponseBodyData) GetIsPrimary() *bool {
	return s.IsPrimary
}

func (s *GetHiMarketConsumerResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetHiMarketConsumerResponseBodyData) SetConsumerId(v string) *GetHiMarketConsumerResponseBodyData {
	s.ConsumerId = &v
	return s
}

func (s *GetHiMarketConsumerResponseBodyData) SetDescription(v string) *GetHiMarketConsumerResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetHiMarketConsumerResponseBodyData) SetIsPrimary(v bool) *GetHiMarketConsumerResponseBodyData {
	s.IsPrimary = &v
	return s
}

func (s *GetHiMarketConsumerResponseBodyData) SetName(v string) *GetHiMarketConsumerResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetHiMarketConsumerResponseBodyData) Validate() error {
	return dara.Validate(s)
}
